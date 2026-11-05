package tts

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestFindVoiceFromEnv(t *testing.T) {
	voice := filepath.Join(t.TempDir(), "en_US-test-medium.onnx")
	if err := os.WriteFile(voice, []byte("not a real model"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("DSATUTOR_TTS_VOICE", voice)
	got, err := findVoice()
	if err != nil {
		t.Fatalf("findVoice: %v", err)
	}
	if got != voice {
		t.Errorf("findVoice = %q, want %q", got, voice)
	}
	if name := voiceName(got); name != "en_US-test-medium" {
		t.Errorf("voiceName = %q, want en_US-test-medium", name)
	}
}

func TestFindVoiceMissingEnvPath(t *testing.T) {
	t.Setenv("DSATUTOR_TTS_VOICE", filepath.Join(t.TempDir(), "missing.onnx"))
	if _, err := findVoice(); err == nil {
		t.Error("findVoice accepted a nonexistent DSATUTOR_TTS_VOICE path")
	}
}

func TestSynthesizeRejectsInvalidText(t *testing.T) {
	e := NewEngine()
	for _, text := range []string{"", "   ", string(bytes.Repeat([]byte("a"), MaxTextLen+1))} {
		if _, err := e.Synthesize(context.Background(), text); err != ErrInvalidText {
			t.Errorf("Synthesize(%d chars) error = %v, want ErrInvalidText", len(text), err)
		}
	}
}

func TestCacheEvictsOldestFirst(t *testing.T) {
	e := NewEngine()
	e.cacheLimit = 10
	e.cachePut("a", []byte("12345"))
	e.cachePut("a", []byte("12345")) // duplicate put must not double-count
	e.cachePut("b", []byte("12345"))
	e.cachePut("c", []byte("12345")) // pushes the cache over the limit
	if _, ok := e.cacheGet("a"); ok {
		t.Error("oldest entry survived eviction")
	}
	for _, key := range []string{"b", "c"} {
		if _, ok := e.cacheGet(key); !ok {
			t.Errorf("entry %q was evicted, want it kept", key)
		}
	}
}

// TestSynthesizeEndToEnd runs only where Piper and a voice are installed
// (e.g. inside the Docker image); elsewhere it skips.
func TestSynthesizeEndToEnd(t *testing.T) {
	e := NewEngine()
	available, voice, reason := e.Status()
	if !available {
		t.Skip("piper not available: " + reason)
	}
	t.Logf("using voice %s", voice)
	wav, err := e.Synthesize(context.Background(), "Binary search halves the search range at every step.")
	if err != nil {
		t.Fatalf("Synthesize: %v", err)
	}
	if !bytes.HasPrefix(wav, []byte("RIFF")) {
		t.Errorf("output does not start with a RIFF header: % x", wav[:min(len(wav), 8)])
	}
}
