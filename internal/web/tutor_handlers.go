package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"dsatutor/internal/chapter"
	"dsatutor/internal/practice"
	"dsatutor/internal/tutor"
)

// maxRecordingBytes bounds one spoken question: about five minutes of
// 16 kHz mono WAV.
const maxRecordingBytes = 10 << 20

// maxToolRounds bounds how many times one reply may call tools before the
// model has to answer in words.
const maxToolRounds = 6

// handleTutorStatus tells the UI whether the AI tutor works (an Ollama with
// a chat model), which models it can pick from and whether the microphone
// button can transcribe. With no chat model the Tutor button stays hidden.
func (s *Server) handleTutorStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.tutor.Status())
}

// handleTutorChat streams the tutor's reply as newline-delimited JSON:
// {"type":"thinking","text":...} and {"type":"delta","text":...} pieces,
// then one {"type":"done","model":...} or {"type":"error","error":...}.
func (s *Server) handleTutorChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Model    string          `json:"model"`
		Think    bool            `json:"think"`
		Messages []tutor.Message `json:"messages"`
		Context  tutor.Context   `json:"context"`
	}
	if err := json.NewDecoder(io.LimitReader(r.Body, 1<<20)).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(req.Messages) == 0 || req.Messages[len(req.Messages)-1].Role != "user" {
		http.Error(w, "the conversation must end with a user message", http.StatusBadRequest)
		return
	}

	// A mock interview needs its problem; without one it's a normal chat.
	problem := s.tutorProblem(req.Context.ProblemID)
	if problem == nil {
		req.Context.Interview = nil
	}
	ch := s.tutorChapter(req.Context.ChapterNumber)
	teach := req.Context.Teach && ch != nil && problem == nil && req.Context.Interview == nil
	if !teach {
		req.Context.Teach, req.Context.Lesson = false, nil
	} else if req.Context.Lesson != nil {
		req.Context.Lesson.Clean()
	}
	lastUser := req.Messages[len(req.Messages)-1].Content

	w.Header().Set("Content-Type", "application/x-ndjson")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("X-Accel-Buffering", "no")
	enc := json.NewEncoder(w)
	flusher, _ := w.(http.Flusher)
	send := func(v any) error {
		if err := enc.Encode(v); err != nil {
			return err
		}
		if flusher != nil {
			flusher.Flush()
		}
		return nil
	}
	emit := func(d tutor.Delta) error {
		if d.Thinking != "" {
			if err := send(map[string]string{"type": "thinking", "text": d.Thinking}); err != nil {
				return err
			}
		}
		if d.Content != "" {
			return send(map[string]string{"type": "delta", "text": d.Content})
		}
		return nil
	}

	// Teacher Mode writes its own lesson: a new lesson (or one restarted with
	// Clear, which leaves only the kickoff cue) gets a fresh plan first. The
	// browser keeps the plan and sends it back with every turn.
	if teach && (!req.Context.Lesson.Valid() || len(req.Messages) == 1) {
		send(map[string]string{"type": "status", "text": "Planning the lesson…"})
		lesson, err := s.tutor.PlanLesson(r.Context(), req.Model, ch, req.Context.Learner)
		if err != nil {
			if r.Context().Err() != nil {
				return
			}
			send(map[string]string{"type": "error", "error": "couldn't plan the lesson: " + err.Error()})
			return
		}
		req.Context.Lesson = lesson
		send(map[string]any{"type": "lesson", "lesson": lesson})
	}

	// The page regions become tools only for models that can call them.
	var ws *tutor.Workspace
	if s.tutor.SupportsTools(req.Model) {
		ws = tutor.NewWorkspace(req.Context)
	}
	var system string
	switch {
	case req.Context.Interview != nil:
		system = tutor.InterviewerPrompt(req.Context, problem, ws)
	case teach:
		system = tutor.TeacherPrompt(req.Context, ch, ws, lastUser)
	default:
		system = tutor.SystemPrompt(req.Context, problem, ch, ws)
	}
	messages := tutor.Conversation(system, req.Messages)
	var tools []tutor.Tool
	if ws != nil {
		tools = ws.Tools()
	}

	// The tool loop: each round streams the model's text; tool calls are
	// carried out (sent to the page as actions) and their results fed back,
	// until the model answers without calling a tool.
	var model string
	var said strings.Builder // everything the model said this turn (as shown)
	var raw strings.Builder  // the same, before stage directions were removed
	// Teacher Mode holds each round's text until the round ends, so
	// narration of its own tool use can be dropped before the learner sees
	// or hears it.
	var held strings.Builder
	turn := &teachTurn{}
	chatEmit := emit
	if teach {
		chatEmit = func(d tutor.Delta) error {
			held.WriteString(d.Content)
			return emit(tutor.Delta{Thinking: d.Thinking})
		}
	}
	did := map[string]bool{} // actions carried out this turn
	for round := 0; ; round++ {
		if round == maxToolRounds {
			tools = nil // last round: make it answer in words
		} else if ws != nil && round > 0 {
			tools = ws.Tools() // a section written this turn can be pointed at now
		}
		held.Reset()
		reply, err := s.tutor.Chat(r.Context(), req.Model, messages, req.Think, tools, chatEmit)
		if err != nil {
			if r.Context().Err() != nil {
				return // the learner pressed Stop
			}
			if round == 0 {
				// Nothing has been streamed yet, so a real status code
				// still reaches the browser.
				switch {
				case errors.Is(err, tutor.ErrUnknownModel):
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case errors.Is(err, tutor.ErrUnavailable):
					http.Error(w, err.Error(), http.StatusServiceUnavailable)
					return
				}
			}
			send(map[string]string{"type": "error", "error": err.Error()})
			return
		}
		model = reply.Model
		raw.WriteString(reply.Content)
		shown := reply.Content
		if teach {
			// Teacher Mode sends its turn in lesson order at the end (see
			// teachTurn); the text is held with everything else.
			shown = tutor.StripStageDirections(held.String())
			if shown != "" && said.Len() > 0 {
				shown = "\n\n" + shown
			}
		}
		said.WriteString(shown)
		// Stop when the model answers in words, and always after the last
		// round: some models keep emitting tool calls even when no tools are
		// offered, which would otherwise loop forever.
		if len(reply.ToolCalls) == 0 || ws == nil || round >= maxToolRounds {
			break
		}
		messages = append(messages, tutor.Message{Role: "assistant", Content: reply.Content, ToolCalls: reply.ToolCalls})
		for _, call := range reply.ToolCalls {
			result, action := ws.Execute(call)
			if action == nil {
				// A rejected call is fed back to the model; log it so prompt and
				// schema problems are visible.
				args, _ := json.Marshal(call.Function.Arguments)
				log.Printf("tutor: %s rejected: %s (args %.300s)", call.Function.Name, result, args)
			}
			if action != nil {
				// A near-empty sketch doesn't count as the turn's visual, so
				// the director still draws a real one.
				if !(action["name"] == "scratchpad_sketch" && tutor.WeakSketch(action)) {
					did[action["name"].(string)] = true
				}
				action["type"] = "action"
				if teach {
					turn.add(action, false)
				} else if err := send(action); err != nil {
					return
				}
			}
			messages = append(messages, tutor.Message{Role: "tool", ToolName: call.Function.Name, Content: result})
		}
		if reply.Content != "" && !teach {
			// Keep text from successive rounds apart in the page.
			send(map[string]string{"type": "delta", "text": "\n\n"})
		}
	}
	// Teacher Mode stages what a turn left out: the section due on the
	// lesson page, a checkpoint, and a scratchpad example for learners who
	// still need explicit support (faded once their evidence reaches the
	// practising level).
	if teach && ws != nil && (strings.TrimSpace(raw.String()) != "" || !did["ask_learner"]) {
		forceVisual := req.Context.Learner == nil || req.Context.Learner.Level == "starting" || req.Context.Learner.Level == "building" || len(req.Context.Learner.Misconceptions) > 0
		due := -1
		if !did["write_lesson"] {
			due = req.Context.Lesson.SectionDue(lastUser, raw.String())
		}
		advance := false
		if l := req.Context.Lesson; due < 0 && !did["write_lesson"] && l.Valid() && l.HasArticle(l.Current) && l.Current+1 < len(l.Sections) && !l.HasArticle(l.Current+1) && len(req.Messages) > 1 {
			// The teacher stayed on the section; move on if the learner
			// has already shown its goal, rather than quizzing it again.
			send(map[string]string{"type": "status", "text": "Checking your progress…"})
			if s.tutor.GoalMet(r.Context(), model, messages, l, raw.String()) {
				due, advance = l.Current+1, true
			}
		}
		if advance {
			turn.ask = nil // a question about the section the learner already has
		}
		need := tutor.DirectorNeeds{
			Article: due >= 0,
			Section: due,
			Advance: advance,
			Sketch:  forceVisual && (advance || !did["scratchpad_sketch"] && !did["scratchpad_visualize"] && !did["scratchpad_code"]),
			Check:   advance || !did["ask_learner"],
		}
		if need.Article || need.Sketch || need.Check {
			send(map[string]string{"type": "status", "text": "Writing up the lesson…"})
			ws.ResetDrawings()
			reply := said.String()
			if strings.TrimSpace(reply) == "" {
				reply = raw.String() // everything was narration; stage what it meant
			}
			if strings.TrimSpace(reply) == "" {
				reply = "(Continue the lesson from where it stands.)"
			}
			calls, err := s.tutor.Direct(r.Context(), model, messages, ch, req.Context.Lesson, reply, need)
			if err != nil {
				if r.Context().Err() != nil {
					return
				}
				log.Printf("tutor: director failed: %v", err)
			}
			for _, call := range calls {
				if call.Function.Name == "write_lesson" {
					// Lead in to the new section in the reply itself, so it
					// is shown and spoken before the section appears.
					t, _ := call.Function.Arguments["transition"].(string)
					turn.transition = tutor.StripStageDirections(t)
					delete(call.Function.Arguments, "transition")
				}
				result, action := ws.Execute(call)
				if action == nil {
					args, _ := json.Marshal(call.Function.Arguments)
					log.Printf("tutor: director %s rejected: %s (args %.300s)", call.Function.Name, result, args)
					continue
				}
				action["type"] = "action"
				turn.add(action, true)
			}
		}
	}
	if teach {
		for _, ev := range turn.events(said.String()) {
			if err := send(ev); err != nil {
				return
			}
		}
	}
	send(map[string]string{"type": "done", "model": model})
}

// handleInterviewScore ends a mock interview: it runs the candidate's final
// code against every test (hidden ones included) and asks the model for a
// scorecard on Google's four coding criteria.
func (s *Server) handleInterviewScore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Model           string          `json:"model"`
		ProblemID       string          `json:"problemId"`
		Language        string          `json:"language"`
		Code            string          `json:"code"`
		Messages        []tutor.Message `json:"messages"`
		ElapsedSeconds  int             `json:"elapsedSeconds"`
		DurationSeconds int             `json:"durationSeconds"`
		Level           string          `json:"level"`
	}
	if err := json.NewDecoder(io.LimitReader(r.Body, 2<<20)).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}
	problem := s.tutorProblem(req.ProblemID)
	if problem == nil {
		http.Error(w, "unknown problem", http.StatusNotFound)
		return
	}
	language := req.Language
	if language != "oxigen" {
		language = "python"
	}

	tests := judgeAll(problem, req.Code, language)
	card, model, err := s.tutor.Score(r.Context(), req.Model, tutor.ScoreInput{
		Problem:    problem,
		Code:       req.Code,
		Language:   language,
		Transcript: req.Messages,
		Elapsed:    time.Duration(req.ElapsedSeconds) * time.Second,
		Duration:   time.Duration(req.DurationSeconds) * time.Second,
		Level:      req.Level,
		Tests:      tests,
	})
	out := map[string]any{"model": model, "tests": map[string]any{
		"status": tests.Status, "passed": tests.PassedCount, "total": tests.TotalCount, "failures": failedNotes(tests),
	}}
	if err != nil {
		// The test results are still worth showing without the write-up.
		out["error"] = err.Error()
	} else {
		out["scorecard"] = card
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

// teachTurn collects a Teacher Mode turn (the model's rounds and the
// director's staging) and sends it in lesson order: the lesson page first,
// read before anything else, then the teacher's words, then one example on
// the scratchpad, then the question. A model sketch too thin to teach from
// gives way to the director's, so the scratchpad gets one drawing per turn.
type teachTurn struct {
	transition string // bridge from the previous section, when moving on
	sections   []tutor.Action
	highlights []tutor.Action
	drawing    tutor.Action
	weak       bool // drawing is a thin model sketch the director may replace
	ask        tutor.Action
}

func (t *teachTurn) add(a tutor.Action, fromDirector bool) {
	switch a["name"] {
	case "write_lesson":
		t.sections = append(t.sections, a)
	case "highlight":
		t.highlights = append(t.highlights, a)
	case "ask_learner":
		// The director's question only exists when the model asked none, or
		// when moving on replaced the model's.
		t.ask = a
	case "scratchpad_sketch", "scratchpad_visualize", "scratchpad_code":
		weak := a["name"] == "scratchpad_sketch" && tutor.WeakSketch(a)
		if t.drawing == nil || (t.weak && (fromDirector || !weak)) {
			t.drawing, t.weak = a, weak && !fromDirector
		}
	}
}

func (t *teachTurn) events(said string) []any {
	var out []any
	text := func(s string) {
		if s = strings.TrimSpace(s); s != "" {
			out = append(out, map[string]string{"type": "delta", "text": s})
		}
	}
	if t.transition != "" {
		// Closing the previous section, then the bridge, then the new one.
		text(said + "\n\n" + t.transition + "\n\n")
		for _, a := range t.sections {
			out = append(out, a)
		}
	} else {
		for _, a := range t.sections {
			out = append(out, a)
		}
		text(said)
	}
	for _, a := range t.highlights {
		out = append(out, a)
	}
	if t.drawing != nil {
		out = append(out, t.drawing)
	}
	if t.ask != nil {
		out = append(out, t.ask)
	}
	return out
}

func failedNotes(t *practice.SubmitResponse) []string {
	notes := []string{}
	for i, r := range t.Results {
		if r.Passed {
			continue
		}
		note := r.Note
		if note == "" {
			note = fmt.Sprintf("test %d", i+1)
		}
		if r.Error != "" {
			note += " (" + strings.SplitN(r.Error, "\n", 2)[0] + ")"
		}
		notes = append(notes, note)
	}
	return notes
}

// handleTutorTranscribe turns a recorded question (the request body, a WAV
// file) into text. ?model= names the chat model in use, preferred for
// transcription when it also takes audio.
func (s *Server) handleTutorTranscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	wav, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxRecordingBytes))
	if err != nil {
		http.Error(w, "recording too large or unreadable: "+err.Error(), http.StatusRequestEntityTooLarge)
		return
	}
	if len(wav) < 44 || !strings.HasPrefix(string(wav[:4]), "RIFF") {
		http.Error(w, "expected a WAV recording", http.StatusBadRequest)
		return
	}
	text, err := s.tutor.Transcribe(r.Context(), wav, r.URL.Query().Get("model"))
	switch {
	case errors.Is(err, tutor.ErrUnavailable):
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"text": text})
}

func (s *Server) tutorProblem(id string) *practice.Problem {
	if id == "" || s.practiceServer == nil {
		return nil
	}
	p, err := s.practiceServer.loader.GetProblem(id)
	if err != nil {
		return nil
	}
	return p
}

func (s *Server) tutorChapter(number int) *chapter.Chapter {
	for i := range s.chapters {
		if s.chapters[i].Number == number {
			return &s.chapters[i]
		}
	}
	return nil
}
