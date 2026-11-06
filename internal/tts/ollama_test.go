package tts

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

// fakeOllama serves just enough of the Ollama API for the tests: /api/tags
// with the given model names, and /v1/audio/speech returning fake WAV bytes.
func fakeOllama(t *testing.T, models []string, speechCalls *atomic.Int64) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/tags", func(w http.ResponseWriter, r *http.Request) {
		type m struct {
			Name string `json:"name"`
		}
		list := []m{}
		for _, name := range models {
			list = append(list, m{name})
		}
		json.NewEncoder(w).Encode(map[string]any{"models": list})
	})
	mux.HandleFunc("/v1/audio/speech", func(w http.ResponseWriter, r *http.Request) {
		if speechCalls != nil {
			speechCalls.Add(1)
		}
		var req struct {
			Model string `json:"model"`
			Input string `json:"input"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Model == "" || req.Input == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "audio/wav")
		w.Write([]byte("RIFFfake-wav-for-" + req.Model))
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

func TestOllamaDetectsSpeechModels(t *testing.T) {
	srv := fakeOllama(t, []string{"llama3:8b", "lfm3-audio:latest", "phonon:latest"}, nil)
	t.Setenv("DSATUTOR_OLLAMA_URL", srv.URL)
	o := NewOllama()
	o.prefer = "" // name order, whatever OS the tests run on
	models, reason := o.SpeechModels()
	if reason != "" {
		t.Fatalf("unexpected reason: %s", reason)
	}
	want := []string{"lfm3-audio:latest", "phonon:latest"}
	if len(models) != len(want) || models[0] != want[0] || models[1] != want[1] {
		t.Errorf("SpeechModels = %v, want %v (llama3 must be excluded)", models, want)
	}
}

// On macOS NewOllama prefers phonon; the preferred model must lead the list,
// because the first available engine is the default voice.
func TestOllamaPreferredModelFirst(t *testing.T) {
	srv := fakeOllama(t, []string{"lfm3-audio:latest", "phonon:latest"}, nil)
	t.Setenv("DSATUTOR_OLLAMA_URL", srv.URL)
	o := NewOllama()
	o.prefer = "phonon"
	models, _ := o.SpeechModels()
	want := []string{"phonon:latest", "lfm3-audio:latest"}
	if len(models) != len(want) || models[0] != want[0] || models[1] != want[1] {
		t.Errorf("SpeechModels = %v, want %v", models, want)
	}
}

func TestOllamaNoSpeechModels(t *testing.T) {
	srv := fakeOllama(t, []string{"llama3:8b"}, nil)
	t.Setenv("DSATUTOR_OLLAMA_URL", srv.URL)
	models, reason := NewOllama().SpeechModels()
	if len(models) != 0 {
		t.Errorf("SpeechModels = %v, want none", models)
	}
	if !strings.Contains(reason, "no speech models") {
		t.Errorf("reason = %q, want a pull hint", reason)
	}
}

func TestOllamaUnreachable(t *testing.T) {
	t.Setenv("DSATUTOR_OLLAMA_URL", "http://127.0.0.1:1")
	models, reason := NewOllama().SpeechModels()
	if len(models) != 0 || !strings.Contains(reason, "unreachable") {
		t.Errorf("SpeechModels = %v, reason = %q; want unreachable", models, reason)
	}
}

func TestOllamaSynthesize(t *testing.T) {
	srv := fakeOllama(t, []string{"phonon:latest"}, nil)
	t.Setenv("DSATUTOR_OLLAMA_URL", srv.URL)
	wav, err := NewOllama().Synthesize(context.Background(), "phonon:latest", "hello")
	if err != nil {
		t.Fatalf("Synthesize: %v", err)
	}
	if !strings.HasPrefix(string(wav), "RIFF") {
		t.Errorf("unexpected audio payload: %q", wav)
	}
}
