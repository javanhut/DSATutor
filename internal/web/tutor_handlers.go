package web

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"dsatutor/internal/chapter"
	"dsatutor/internal/practice"
	"dsatutor/internal/tutor"
)

// maxRecordingBytes bounds one spoken question: about five minutes of
// 16 kHz mono WAV.
const maxRecordingBytes = 10 << 20

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

	system := tutor.SystemPrompt(req.Context, s.tutorProblem(req.Context.ProblemID), s.tutorChapter(req.Context.ChapterNumber))
	messages := tutor.Conversation(system, req.Messages)

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

	model, err := s.tutor.Chat(r.Context(), req.Model, messages, req.Think, func(d tutor.Delta) error {
		if d.Thinking != "" {
			if err := send(map[string]string{"type": "thinking", "text": d.Thinking}); err != nil {
				return err
			}
		}
		if d.Content != "" {
			return send(map[string]string{"type": "delta", "text": d.Content})
		}
		return nil
	})
	if err != nil {
		if r.Context().Err() != nil {
			return // the learner pressed Stop
		}
		// Nothing has been streamed yet for these, so a real status code
		// still reaches the browser.
		switch {
		case errors.Is(err, tutor.ErrUnknownModel):
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case errors.Is(err, tutor.ErrUnavailable):
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		send(map[string]string{"type": "error", "error": err.Error()})
		return
	}
	send(map[string]string{"type": "done", "model": model})
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
