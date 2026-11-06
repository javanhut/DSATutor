package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

// Pointing DSATUTOR_TTS_VOICE at a missing file and the Ollama URL at a
// closed port makes every engine fail deterministically, whatever is
// installed on the machine running the tests.
func newServerWithoutTTS(t *testing.T) *Server {
	t.Setenv("DSATUTOR_TTS_VOICE", filepath.Join(t.TempDir(), "missing.onnx"))
	t.Setenv("DSATUTOR_OLLAMA_URL", "http://127.0.0.1:1")
	return NewServer(nil)
}

func TestTTSStatusReportsUnavailable(t *testing.T) {
	s := newServerWithoutTTS(t)
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/tts/status", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var body struct {
		Available bool `json:"available"`
		Engines   []struct {
			ID        string `json:"id"`
			Available bool   `json:"available"`
			Reason    string `json:"reason"`
		} `json:"engines"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if body.Available {
		t.Error("available = true with no working engine")
	}
	if len(body.Engines) == 0 {
		t.Fatal("engines list is empty; the UI and operators need the per-engine reasons")
	}
	for _, e := range body.Engines {
		if e.Available || e.Reason == "" {
			t.Errorf("engine %s: available=%v reason=%q, want unavailable with a reason", e.ID, e.Available, e.Reason)
		}
	}
}

func TestTTSSynthesizeErrors(t *testing.T) {
	s := newServerWithoutTTS(t)
	cases := []struct {
		name string
		body string
		want int
	}{
		{"unavailable", `{"text":"hello"}`, http.StatusServiceUnavailable},
		{"empty text", `{"text":""}`, http.StatusBadRequest},
		{"unknown engine", `{"text":"hello","engine":"nope"}`, http.StatusBadRequest},
		{"malformed JSON", `{"text":`, http.StatusBadRequest},
	}
	for _, tc := range cases {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/tts", strings.NewReader(tc.body))
		s.ServeHTTP(rec, req)
		if rec.Code != tc.want {
			t.Errorf("%s: status = %d, want %d", tc.name, rec.Code, tc.want)
		}
	}

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/tts", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("GET /api/tts status = %d, want 405", rec.Code)
	}
}
