package web

import (
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
