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
	// The page regions become tools only for models that can call them.
	var ws *tutor.Workspace
	if s.tutor.SupportsTools(req.Model) {
		ws = tutor.NewWorkspace(req.Context)
	}
	var system string
	ch := s.tutorChapter(req.Context.ChapterNumber)
	if req.Context.Interview != nil {
		system = tutor.InterviewerPrompt(req.Context, problem, ws)
	} else if req.Context.Teach && ch != nil && problem == nil {
		system = tutor.TeacherPrompt(req.Context, ch, ws)
	} else {
		system = tutor.SystemPrompt(req.Context, problem, ch, ws)
	}
	messages := tutor.Conversation(system, req.Messages)
	var tools []tutor.Tool
	if ws != nil {
		tools = ws.Tools()
	}

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

	// The tool loop: each round streams the model's text; tool calls are
	// carried out (sent to the page as actions) and their results fed back,
	// until the model answers without calling a tool.
	var model string
	for round := 0; ; round++ {
		if round == maxToolRounds {
			tools = nil // last round: make it answer in words
		}
		reply, err := s.tutor.Chat(r.Context(), req.Model, messages, req.Think, tools, emit)
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
				action["type"] = "action"
				if err := send(action); err != nil {
					return
				}
			}
			messages = append(messages, tutor.Message{Role: "tool", ToolName: call.Function.Name, Content: result})
		}
		if reply.Content != "" {
			// Keep text from successive rounds apart in the page.
			send(map[string]string{"type": "delta", "text": "\n\n"})
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
