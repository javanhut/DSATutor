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

// TestPiperEndToEnd runs only where Piper and a voice are installed (e.g.
// inside the Docker image); elsewhere it skips.
func TestPiperEndToEnd(t *testing.T) {
	p := NewPiper()
	available, voice, reason := p.Status()
	if !available {
		t.Skip("piper not available: " + reason)
	}
	t.Logf("using voice %s", voice)
	wav, err := p.Synthesize(context.Background(), "Binary search halves the search range at every step.")
	if err != nil {
		t.Fatalf("Synthesize: %v", err)
	}
	if !bytes.HasPrefix(wav, []byte("RIFF")) {
		t.Errorf("output does not start with a RIFF header: % x", wav[:min(len(wav), 8)])
	}
}
