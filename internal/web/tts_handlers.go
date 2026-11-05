package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"dsatutor/internal/tts"
)

// handleTTSStatus lets the UI decide whether to show Listen buttons at all;
// a checkout without Piper installed simply hides the feature.
func (s *Server) handleTTSStatus(w http.ResponseWriter, r *http.Request) {
	available, voice, reason := s.tts.Status()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"available": available,
		"voice":     voice,
		"reason":    reason,
	})
}

// handleTTS renders a chunk of lesson text as WAV audio.
func (s *Server) handleTTS(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}
	wav, err := s.tts.Synthesize(r.Context(), req.Text)
	switch {
	case errors.Is(err, tts.ErrInvalidText):
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	case errors.Is(err, tts.ErrUnavailable):
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "audio/wav")
	w.Write(wav)
}
