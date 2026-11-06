package tts

import (
	"bytes"
	"context"
	"errors"
	"path/filepath"
	"sync/atomic"
	"testing"
)

// disablePiper makes Piper deterministically unavailable regardless of what
// is installed on the machine running the tests.
func disablePiper(t *testing.T) {
	t.Helper()
	t.Setenv("DSATUTOR_TTS_VOICE", filepath.Join(t.TempDir(), "missing.onnx"))
}

func TestServiceRejectsInvalidText(t *testing.T) {
	disablePiper(t)
	t.Setenv("DSATUTOR_OLLAMA_URL", "http://127.0.0.1:1")
	s := NewService()
	for _, text := range []string{"", "   ", string(bytes.Repeat([]byte("a"), MaxTextLen+1))} {
		if _, _, err := s.Synthesize(context.Background(), text, "", ""); !errors.Is(err, ErrInvalidText) {
			t.Errorf("Synthesize(%d chars) error = %v, want ErrInvalidText", len(text), err)
		}
	}
}

func TestServiceUnavailableWithoutEngines(t *testing.T) {
	disablePiper(t)
	t.Setenv("DSATUTOR_OLLAMA_URL", "http://127.0.0.1:1")
	s := NewService()
	st := s.Status()
	if st.Available || st.Default != "" {
		t.Errorf("Status = %+v, want unavailable with no default", st)
	}
	if _, _, err := s.Synthesize(context.Background(), "hello", "", ""); !errors.Is(err, ErrUnavailable) {
		t.Errorf("Synthesize error = %v, want ErrUnavailable", err)
	}
}

func TestServiceUsesOllamaAndCaches(t *testing.T) {
	disablePiper(t)
	var calls atomic.Int64
	srv := fakeOllama(t, []string{"lfm3-audio:latest"}, &calls)
	t.Setenv("DSATUTOR_OLLAMA_URL", srv.URL)
	s := NewService()

	st := s.Status()
	if !st.Available || st.Default != "ollama:lfm3-audio:latest" {
		t.Fatalf("Status = %+v, want default ollama:lfm3-audio:latest", st)
	}

	wav, engine, err := s.Synthesize(context.Background(), "hello", "", "")
	if err != nil || engine != "ollama:lfm3-audio:latest" {
		t.Fatalf("Synthesize = engine %q, err %v", engine, err)
	}
	if !bytes.HasPrefix(wav, []byte("RIFF")) {
		t.Errorf("unexpected audio: %q", wav)
	}
	if _, _, err := s.Synthesize(context.Background(), "hello", "", ""); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 1 {
		t.Errorf("speech endpoint called %d times, want 1 (second hit must come from cache)", calls.Load())
	}
}

func TestServiceRejectsUnknownEngine(t *testing.T) {
	disablePiper(t)
	srv := fakeOllama(t, []string{"phonon:latest"}, nil)
	t.Setenv("DSATUTOR_OLLAMA_URL", srv.URL)
	s := NewService()
	if _, _, err := s.Synthesize(context.Background(), "hello", "ollama:gone:latest", ""); !errors.Is(err, ErrUnknownEngine) {
		t.Errorf("Synthesize error = %v, want ErrUnknownEngine", err)
	}
}

func TestServiceVoiceSelection(t *testing.T) {
	disablePiper(t)
	srv := fakeOllama(t, []string{"phonon:latest"}, nil)
	t.Setenv("DSATUTOR_OLLAMA_URL", srv.URL)
	s := NewService()

	if got := s.Status().Engines[0].Voices; len(got) != 2 || got[0] != "Alfie" || got[1] != "Zoey" {
		t.Fatalf("status voices = %v, want [Alfie Zoey]", got)
	}
	// The same text in each voice: a chosen voice reaches Ollama, an unknown
	// one falls back to the default, and no voice is served another's audio
	// from the cache.
	for voice, want := range map[string]string{
		"Zoey":   "RIFFfake-wav-for-phonon:latest#Zoey",
		"Alfie":  "RIFFfake-wav-for-phonon:latest#Alfie",
		"":       "RIFFfake-wav-for-phonon:latest",
		"Nobody": "RIFFfake-wav-for-phonon:latest",
	} {
		wav, engine, err := s.Synthesize(context.Background(), "hello", "ollama:phonon:latest", voice)
		if err != nil || engine != "ollama:phonon:latest" {
			t.Fatalf("voice %q: engine %q, err %v", voice, engine, err)
		}
		if string(wav) != want {
			t.Errorf("voice %q: audio = %q, want %q", voice, wav, want)
		}
	}
}

func TestServiceCacheEvictsOldestFirst(t *testing.T) {
	s := &Service{cache: map[string][]byte{}, cacheLimit: 10}
	s.cachePut("a", []byte("12345"))
	s.cachePut("a", []byte("12345")) // duplicate put must not double-count
	s.cachePut("b", []byte("12345"))
	s.cachePut("c", []byte("12345")) // pushes the cache over the limit
	if _, ok := s.cacheGet("a"); ok {
		t.Error("oldest entry survived eviction")
	}
	for _, key := range []string{"b", "c"} {
		if _, ok := s.cacheGet(key); !ok {
			t.Errorf("entry %q was evicted, want it kept", key)
		}
	}
}
