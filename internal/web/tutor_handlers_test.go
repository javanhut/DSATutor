package web

import (
	"dsatutor/internal/chapter"
	"dsatutor/internal/tutor"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTutorStatusReportsUnavailable(t *testing.T) {
	s := newServerWithoutTTS(t) // also points Ollama at a closed port
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/tutor/status", nil))
	var body struct {
		Available bool   `json:"available"`
		Reason    string `json:"reason"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if body.Available || body.Reason == "" {
		t.Errorf("status = %+v, want unavailable with a reason", body)
	}
}

func TestTutorChatWithoutOllamaIs503(t *testing.T) {
	s := newServerWithoutTTS(t)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/tutor/chat",
		strings.NewReader(`{"messages":[{"role":"user","content":"help"}],"context":{"view":"home"}}`))
	s.ServeHTTP(rec, req)
	if rec.Code != http.StatusServiceUnavailable {
		t.Errorf("status = %d, want 503", rec.Code)
	}
}

func TestTutorChatRejectsBadRequests(t *testing.T) {
	s := newServerWithoutTTS(t)
	for _, body := range []string{`not json`, `{"messages":[]}`, `{"messages":[{"role":"assistant","content":"hi"}]}`} {
		rec := httptest.NewRecorder()
		s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/tutor/chat", strings.NewReader(body)))
		if rec.Code != http.StatusBadRequest {
			t.Errorf("%s: status = %d, want 400", body, rec.Code)
		}
	}
}

func TestTutorTranscribeRejectsNonWAV(t *testing.T) {
	s := newServerWithoutTTS(t)
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/tutor/transcribe", strings.NewReader("not audio")))
	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", rec.Code)
	}
}

// A fake Ollama whose model first calls highlight and write_code, then
// answers in words once it has the tool results.
func TestTutorChatRunsToolLoop(t *testing.T) {
	var rounds int
	var lastMessages []map[string]any
	ollama := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/tags":
			w.Write([]byte(`{"models":[{"name":"m:latest"}]}`))
		case "/api/show":
			w.Write([]byte(`{"capabilities":["completion","tools"]}`))
		case "/api/chat":
			var body struct {
				Messages []map[string]any `json:"messages"`
				Tools    []any            `json:"tools"`
			}
			json.NewDecoder(r.Body).Decode(&body)
			lastMessages = body.Messages
			rounds++
			if rounds == 1 {
				if len(body.Tools) == 0 {
					t.Error("no tools offered")
				}
				w.Write([]byte(`{"message":{"content":"Look:","tool_calls":[{"function":{"name":"highlight","arguments":{"target":"code","start_line":1,"note":"bug"}}},{"function":{"name":"write_code","arguments":{"mode":"replace_all","code":"fixed()"}}}]},"done":true}` + "\n"))
				return
			}
			w.Write([]byte(`{"message":{"content":"Fixed it."},"done":true}` + "\n"))
		}
	}))
	defer ollama.Close()
	t.Setenv("DSATUTOR_OLLAMA_URL", ollama.URL)
	t.Setenv("DSATUTOR_TUTOR_MODEL", "")
	s := NewServer(nil)

	req := `{"messages":[{"role":"user","content":"fix it"}],"context":{"view":"sandbox","code":"broken()","targets":[{"id":"code","label":"editor","lines":true,"writable":true}]}}`
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/tutor/chat", strings.NewReader(req)))

	var events []map[string]any
	for _, line := range strings.Split(strings.TrimSpace(rec.Body.String()), "\n") {
		var ev map[string]any
		if err := json.Unmarshal([]byte(line), &ev); err != nil {
			t.Fatalf("bad line %q", line)
		}
		events = append(events, ev)
	}
	var kinds []string
	for _, ev := range events {
		k := ev["type"].(string)
		if k == "action" {
			k += ":" + ev["name"].(string)
		}
		kinds = append(kinds, k)
	}
	if got := strings.Join(kinds, ","); got != "delta,action:highlight,action:write_code,delta,delta,done" {
		t.Errorf("events = %s", got)
	}
	if events[2]["code"] != "fixed()" {
		t.Errorf("write_code action = %v", events[2])
	}
	if rounds != 2 {
		t.Errorf("rounds = %d, want 2", rounds)
	}
	// The second round got both tool results, the edit showing the new code.
	var results []string
	for _, m := range lastMessages {
		if m["role"] == "tool" {
			results = append(results, m["content"].(string))
		}
	}
	if len(results) != 2 || !strings.Contains(results[1], "fixed()") {
		t.Errorf("tool results = %q", results)
	}
}

// A model that calls a tool on every turn, even when no tools are offered,
// must not keep the request running forever.
func TestTutorChatToolLoopIsBounded(t *testing.T) {
	rounds := 0
	ollama := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/tags":
			w.Write([]byte(`{"models":[{"name":"m:latest"}]}`))
		case "/api/show":
			w.Write([]byte(`{"capabilities":["completion","tools"]}`))
		case "/api/chat":
			rounds++
			w.Write([]byte(`{"message":{"content":"","tool_calls":[{"function":{"name":"ask_learner","arguments":{"question":"q","choices":"not a list"}}}]},"done":true}` + "\n"))
		}
	}))
	defer ollama.Close()
	t.Setenv("DSATUTOR_OLLAMA_URL", ollama.URL)
	t.Setenv("DSATUTOR_TUTOR_MODEL", "")
	s := NewServer(nil)

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/tutor/chat",
		strings.NewReader(`{"messages":[{"role":"user","content":"teach me"}],"context":{"view":"home"}}`)))
	if rounds != maxToolRounds+1 {
		t.Errorf("model called %d times, want %d", rounds, maxToolRounds+1)
	}
	if !strings.Contains(rec.Body.String(), `"type":"done"`) {
		t.Errorf("reply never finished: %s", rec.Body.String())
	}
}

// In Teacher Mode, a new lesson is planned first, then a model that only
// talks gets its turn staged by the director: one structured call returns
// the section's lesson-page text, a sketch and a checkpoint, which are
// carried out like tool calls.
func TestTeacherDirectorStagesTextOnlyTurns(t *testing.T) {
	var formats, plans int
	thin, goalMet := false, false
	ollama := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/tags":
			w.Write([]byte(`{"models":[{"name":"m:latest"}]}`))
		case "/api/show":
			w.Write([]byte(`{"capabilities":["completion","tools"]}`))
		case "/api/chat":
			var body map[string]any
			json.NewDecoder(r.Body).Decode(&body)
			if f, ok := body["format"].(map[string]any); ok {
				props, _ := f["properties"].(map[string]any)
				if props["sections"] != nil {
					plans++
					plan := `{"title":"Two Pointers","sections":[{"heading":"Why","goal":"see the slow way"},{"heading":"Walk inward","goal":"trace the pointers"},{"heading":"Invariant","goal":"explain why"},{"heading":"Practice","goal":"solve one"}]}`
					json.NewEncoder(w).Encode(map[string]any{"message": map[string]string{"content": plan}, "done": true})
					return
				}
				if props["goal_met"] != nil {
					json.NewEncoder(w).Encode(map[string]any{"message": map[string]string{"content": fmt.Sprintf(`{"evidence":"e","goal_met":%t}`, goalMet)}, "done": true})
					return
				}
				formats++
				transition := ""
				if art, _ := props["article"].(map[string]any); art != nil {
					if ap, _ := art["properties"].(map[string]any); ap["transition"] != nil {
						transition = `"transition":"That settles the slow way; now the pointers do the work.",`
					}
				}
				article := `"article":{` + transition + `"heading":"Walk inward","body":"Put one pointer at each end of the sorted array and move the one that brings the sum closer to the target. Each step discards one element for good."},`
				card := `{` + article + `"sketch":{"title":"t","captions":["a"],"rows":[{"step":0,"x":10,"y":10,"label":"nums","values":[1,2,3],"pointers":{"left":0}}],"notes":[{"step":0,"x":10,"y":40,"text":"start","color":"ink"}],"arrows":[{"step":1,"x":10,"y":50,"x2":20,"y2":30,"label":"move"}],"boxes":[]},"check":{"question":"Which pointer moves next?","choices":["left","right"]}}`
				if thin {
					card = `{` + article + `"sketch":{"title":"t","captions":["a"],"rows":[],"notes":[{"step":0,"x":10,"y":10,"text":"idea","color":"ink"}],"arrows":[],"boxes":[]},"check":{"question":"Which pointer moves next?","choices":["left","right"]}}`
				}
				json.NewEncoder(w).Encode(map[string]any{"message": map[string]string{"content": card}, "done": true})
				return
			}
			w.Write([]byte(`{"message":{"content":"Two pointers walk inward."},"done":true}` + "\n"))
		}
	}))
	defer ollama.Close()
	t.Setenv("DSATUTOR_OLLAMA_URL", ollama.URL)
	t.Setenv("DSATUTOR_TUTOR_MODEL", "")
	s := NewServer([]chapter.Chapter{{Number: 3, Title: "Two Pointers", Summary: "Two pointers walk inward."}})

	events := func(body string) ([]string, map[string]map[string]any) {
		rec := httptest.NewRecorder()
		s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/tutor/chat", strings.NewReader(body)))
		var kinds []string
		byName := map[string]map[string]any{"delta": {"text": ""}}
		for _, line := range strings.Split(strings.TrimSpace(rec.Body.String()), "\n") {
			var ev map[string]any
			json.Unmarshal([]byte(line), &ev)
			k, _ := ev["type"].(string)
			if k == "action" {
				k += ":" + ev["name"].(string)
				byName[ev["name"].(string)] = ev
			}
			if k == "lesson" {
				byName["lesson"] = ev
			}
			if k == "delta" {
				byName["delta"]["text"] = byName["delta"]["text"].(string) + ev["text"].(string)
			}
			kinds = append(kinds, k)
		}
		return kinds, byName
	}

	// A fresh lesson: plan, then the first section is written.
	kickoff := `{"messages":[{"role":"user","content":"(Start the lesson)"}],"context":{"view":"chapter","chapterNumber":3,"teach":true}}`
	kinds, got := events(kickoff)
	if j := strings.Join(kinds, ","); j != "status,lesson,status,action:write_lesson,delta,action:scratchpad_sketch,action:ask_learner,done" {
		t.Errorf("events = %s", j)
	}
	if plans != 1 || formats != 1 {
		t.Errorf("plans = %d, director calls = %d; want 1 and 1", plans, formats)
	}
	if got["write_lesson"]["section"] != float64(0) || got["write_lesson"]["total"] != float64(4) {
		t.Errorf("write_lesson = %v", got["write_lesson"])
	}
	// Director rows become cells elements with real values and pointers.
	el := got["scratchpad_sketch"]["elements"].([]any)[0].(map[string]any)
	if _, ok := el["values"].([]any); !ok {
		t.Errorf("values not decoded: %#v", el["values"])
	}
	if _, ok := el["pointers"].(map[string]any); !ok {
		t.Errorf("pointers not decoded: %#v", el["pointers"])
	}

	// A thin sketch from the director gets exactly one retry.
	formats, thin = 0, true
	events(kickoff)
	if formats != 2 {
		t.Errorf("thin sketch: director called %d times, want 2 (one retry)", formats)
	}
	thin = false

	// Moving on with a plan in hand writes the next section, without
	// replanning; a secure learner gets no forced sketch.
	formats, plans = 0, 0
	lesson := `{"title":"Two Pointers","sections":[{"heading":"Why","goal":"g"},{"heading":"Invariant","goal":"g"}],"current":0,"written":[0]}`
	next := `{"messages":[{"role":"user","content":"(Start)"},{"role":"assistant","content":"ok"},{"role":"user","content":"Got it. Let's move on to the next idea."}],"context":{"view":"chapter","chapterNumber":3,"teach":true,"lesson":` + lesson + `,"learner":{"level":"secure","mastery":82}}}`
	kinds, got = events(next)
	if j := strings.Join(kinds, ","); j != "status,delta,action:write_lesson,action:ask_learner,done" {
		t.Errorf("move-on events = %s", j)
	}
	if text := got["delta"]["text"]; text != "Two pointers walk inward.\n\nThat settles the slow way; now the pointers do the work." {
		t.Errorf("move-on reply = %q, want the transition after the reply", text)
	}
	if plans != 0 || got["write_lesson"]["section"] != float64(1) {
		t.Errorf("plans = %d, write_lesson = %v", plans, got["write_lesson"])
	}

	// Mid-section, nothing new is written; only the checkpoint is staged.
	mid := strings.Replace(next, "Got it. Let's move on to the next idea.", "Is it O(n)?", 1)
	if kinds, _ = events(mid); strings.Join(kinds, ",") != "status,status,delta,action:ask_learner,done" {
		t.Errorf("mid-section events = %s", strings.Join(kinds, ","))
	}

	// Once the learner has met the section's goal, the next section is
	// written even though the teacher's reply stayed put.
	goalMet = true
	kinds, got = events(mid)
	if j := strings.Join(kinds, ","); j != "status,status,delta,action:write_lesson,action:ask_learner,done" || got["write_lesson"]["section"] != float64(1) {
		t.Errorf("goal met: events = %s, write_lesson = %v", j, got["write_lesson"])
	}
	goalMet = false

	// Outside Teacher Mode there is no planning and no director.
	formats, plans = 0, 0
	events(`{"messages":[{"role":"user","content":"hi"}],"context":{"view":"chapter","chapterNumber":3,"targets":[{"id":"section-1","label":"Overview"}]}}`)
	if formats != 0 || plans != 0 {
		t.Error("director or planner ran outside Teacher Mode")
	}
}

// A thin sketch from the model gives way to the director's, so the
// scratchpad gets one drawing per turn; a real one from the model stays.
func TestTeachTurnKeepsOneDrawing(t *testing.T) {
	thin := tutor.Action{"type": "action", "name": "scratchpad_sketch", "elements": []any{map[string]any{"type": "text", "text": "idea"}}}
	full := tutor.Action{"type": "action", "name": "scratchpad_sketch", "title": "director", "elements": []any{
		map[string]any{"type": "cells", "values": []any{1, 2}}, map[string]any{"type": "text", "text": "a"}, map[string]any{"type": "text", "text": "b"}}}
	section := tutor.Action{"type": "action", "name": "write_lesson", "section": 0}

	turn := &teachTurn{}
	turn.add(thin, false)
	turn.add(section, true)
	turn.add(full, true)
	var names []string
	for _, ev := range turn.events("Words.") {
		if m, ok := ev.(tutor.Action); ok {
			names = append(names, m["name"].(string))
		} else {
			names = append(names, "delta")
		}
	}
	if got := strings.Join(names, ","); got != "write_lesson,delta,scratchpad_sketch" {
		t.Errorf("order = %s, want the section read first, then words, then one drawing", got)
	}
	if turn.drawing["title"] != "director" {
		t.Errorf("kept %v, want the director's sketch", turn.drawing)
	}

	turn = &teachTurn{}
	model := tutor.Action{"name": "scratchpad_code", "code": "x = 1"}
	turn.add(model, false)
	turn.add(full, true)
	if turn.drawing["name"] != "scratchpad_code" {
		t.Errorf("a real model drawing was replaced by %v", turn.drawing["name"])
	}
}
