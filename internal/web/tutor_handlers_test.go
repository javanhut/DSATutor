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
