package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

// Pointing DSATUTOR_TTS_VOICE at a missing file makes availability fail
// deterministically, whatever is installed on the machine running the tests.
func newServerWithoutTTS(t *testing.T) *Server {
	t.Setenv("DSATUTOR_TTS_VOICE", filepath.Join(t.TempDir(), "missing.onnx"))
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
		Available bool   `json:"available"`
		Reason    string `json:"reason"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if body.Available {
		t.Error("available = true with a nonexistent voice model")
	}
	if body.Reason == "" {
		t.Error("reason is empty; the UI and operators need to know why")
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
