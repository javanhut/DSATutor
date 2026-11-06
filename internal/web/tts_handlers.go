package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"dsatutor/internal/tts"
)

// handleTTSStatus tells the UI which narration engines exist (Ollama speech
// models, Piper) so it can render the voice picker; with no engine available
// the Listen buttons are hidden entirely.
func (s *Server) handleTTSStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.tts.Status())
}

// handleTTS renders a chunk of lesson text as WAV audio with the requested
// engine (or the default), falling back to the other engines on failure. The
// X-TTS-Engine header names the engine that actually produced the audio.
func (s *Server) handleTTS(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Text   string `json:"text"`
		Engine string `json:"engine"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}
	wav, engine, err := s.tts.Synthesize(r.Context(), req.Text, req.Engine)
	switch {
	case errors.Is(err, tts.ErrInvalidText), errors.Is(err, tts.ErrUnknownEngine):
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
	w.Header().Set("X-TTS-Engine", engine)
	w.Write(wav)
}
