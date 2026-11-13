package web

import (
	"dsatutor/internal/chapter"
	"encoding/json"
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

// In Teacher Mode, a model that only talks gets its turn staged by the
// director: one structured call returns a highlight and a sketch, which are
// carried out like tool calls.
func TestTeacherDirectorStagesTextOnlyTurns(t *testing.T) {
	var formats int
	thin := false
	ollama := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/tags":
			w.Write([]byte(`{"models":[{"name":"m:latest"}]}`))
		case "/api/show":
			w.Write([]byte(`{"capabilities":["completion","tools"]}`))
		case "/api/chat":
			var body map[string]any
			json.NewDecoder(r.Body).Decode(&body)
			if body["format"] != nil {
				formats++
				card := `{"highlight":{"target":"section-1","quote":"two pointers","pointer":"hand","note":"here"},"sketch":{"title":"t","captions":["a"],"rows":[{"step":0,"x":10,"y":10,"label":"nums","values":[1,2,3],"pointers":{"left":0}}],"notes":[{"step":0,"x":10,"y":40,"text":"start","color":"ink"}],"arrows":[{"step":1,"x":10,"y":50,"x2":20,"y2":30,"label":"move"}],"boxes":[]},"check":{"question":"Which pointer moves next?","choices":["left","right"]}}`
				if thin {
					card = `{"highlight":{"target":"section-1","quote":"two pointers","pointer":"hand","note":"here"},"sketch":{"title":"t","captions":["a"],"rows":[],"notes":[{"step":0,"x":10,"y":10,"text":"idea","color":"ink"}],"arrows":[],"boxes":[]},"check":{"question":"Which pointer moves next?","choices":["left","right"]}}`
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
	s := NewServer([]chapter.Chapter{{Number: 3, Title: "Two Pointers", Summary: "s"}})

	req := `{"messages":[{"role":"user","content":"next"}],"context":{"view":"chapter","chapterNumber":3,"teach":true,"targets":[{"id":"section-1","label":"Overview"}]}}`
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/tutor/chat", strings.NewReader(req)))

	var kinds []string
	var sketch map[string]any
	for _, line := range strings.Split(strings.TrimSpace(rec.Body.String()), "\n") {
		var ev map[string]any
		json.Unmarshal([]byte(line), &ev)
		k := ev["type"].(string)
		if k == "action" {
			k += ":" + ev["name"].(string)
			if ev["name"] == "scratchpad_sketch" {
				sketch = ev
			}
		}
		kinds = append(kinds, k)
	}
	if got := strings.Join(kinds, ","); got != "delta,status,action:highlight,action:scratchpad_sketch,action:ask_learner,done" {
		t.Errorf("events = %s", got)
	}
	if formats != 1 {
		t.Errorf("director called %d times, want 1", formats)
	}
	// Director rows become cells elements with real values and pointers.
	el := sketch["elements"].([]any)[0].(map[string]any)
	if _, ok := el["values"].([]any); !ok {
		t.Errorf("values not decoded: %#v", el["values"])
	}
	if _, ok := el["pointers"].(map[string]any); !ok {
		t.Errorf("pointers not decoded: %#v", el["pointers"])
	}

	// A thin sketch from the director gets exactly one retry.
	formats, thin = 0, true
	rec = httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/tutor/chat", strings.NewReader(req)))
	if formats != 2 {
		t.Errorf("thin sketch: director called %d times, want 2 (one retry)", formats)
	}
	thin = false

	// A learner with secure evidence still gets a precise pointer and a
	// checkpoint, but the director no longer forces another explanatory sketch.
	formats = 0
	secureReq := `{"messages":[{"role":"user","content":"next"}],"context":{"view":"chapter","chapterNumber":3,"teach":true,"learner":{"level":"secure","mastery":82},"targets":[{"id":"section-1","label":"Overview"}]}}`
	rec = httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/tutor/chat", strings.NewReader(secureReq)))
	kinds = nil
	for _, line := range strings.Split(strings.TrimSpace(rec.Body.String()), "\n") {
		var ev map[string]any
		json.Unmarshal([]byte(line), &ev)
		kind := ev["type"].(string)
		if kind == "action" {
			kind += ":" + ev["name"].(string)
		}
		kinds = append(kinds, kind)
	}
	if got := strings.Join(kinds, ","); got != "delta,status,action:highlight,action:ask_learner,done" {
		t.Errorf("secure learner events = %s", got)
	}

	// Outside Teacher Mode there is no director.
	formats = 0
	rec = httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/tutor/chat",
		strings.NewReader(`{"messages":[{"role":"user","content":"hi"}],"context":{"view":"chapter","chapterNumber":3,"targets":[{"id":"section-1","label":"Overview"}]}}`)))
	if formats != 0 {
		t.Error("director ran outside Teacher Mode")
	}
}
