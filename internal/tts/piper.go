// Piper is the built-in narration engine: free, CPU-only, bundled in the
// Docker image, so narration always works even with no Ollama around. The
// voice model is loaded once in a persistent python3 worker (speaker.py);
// after that, synthesis on a CPU is faster than real time.
package tts

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

//go:embed speaker.py
var speakerScript []byte

const (
	// startTimeout covers loading the voice model (~60MB) on first use.
	startTimeout = 60 * time.Second
	synthTimeout = 60 * time.Second
)

// voiceDirs are searched for a *.onnx voice when DSATUTOR_TTS_VOICE is not
// set; /app/voices is where the Dockerfile puts the downloaded voice.
var voiceDirs = []string{"/app/voices", "voices"}

// errSynthesis marks a failure reported by a healthy worker, so the retry
// loop doesn't restart the process for text it can't speak anyway.
var errSynthesis = errors.New("piper could not synthesize this text")

// Piper owns the worker process. The zero value is not usable; create one
// with NewPiper.
type Piper struct {
	mu        sync.Mutex // guards the worker process and the availability check
	checked   bool
	available bool
	reason    string // why unavailable, surfaced by /api/tts/status
	voice     string // path to the .onnx voice model

	scriptDir string
	cmd       *exec.Cmd
	stdin     io.WriteCloser
	stdout    *bufio.Reader
	stderr    *bytes.Buffer
}

func NewPiper() *Piper {
	return &Piper{}
}

// Status reports whether synthesis can work here and, if not, why. The check
// (a voice model on disk, python3 able to import piper) runs once and is
// remembered; images built from the Dockerfile always pass it, while local
// checkouts without Piper just hide the engine in the UI.
func (p *Piper) Status() (available bool, voice, reason string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.checkLocked()
	return p.available, voiceName(p.voice), p.reason
}

// Synthesize returns a WAV rendering of text. Requests are serialized: Piper
// is CPU-bound, so running them one at a time is as fast as interleaving and
// keeps the worker protocol trivial.
func (p *Piper) Synthesize(ctx context.Context, text string) ([]byte, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.checkLocked()
	if !p.available {
		return nil, errors.New(p.reason)
	}
	ctx, cancel := context.WithTimeout(ctx, synthTimeout)
	defer cancel()

	// One retry after a restart, in case the worker died since its last use.
	var wav []byte
	var err error
	for attempt := 0; attempt < 2; attempt++ {
		if err = p.startLocked(); err != nil {
			return nil, err
		}
		wav, err = p.roundTripLocked(ctx, text)
		if err == nil || errors.Is(err, errSynthesis) {
			break
		}
		p.stopLocked()
		if ctx.Err() != nil {
			break
		}
	}
	return wav, err
}

func (p *Piper) checkLocked() {
	if p.checked {
		return
	}
	p.checked = true
	voice, err := findVoice()
	if err != nil {
		p.reason = err.Error()
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if out, err := exec.CommandContext(ctx, "python3", "-c", "import piper").CombinedOutput(); err != nil {
		p.reason = fmt.Sprintf("python3 cannot import piper (pip install piper-tts): %s", lastLine(out, err))
		return
	}
	p.voice = voice
	p.available = true
}

func findVoice() (string, error) {
	if path := os.Getenv("DSATUTOR_TTS_VOICE"); path != "" {
		if _, err := os.Stat(path); err != nil {
			return "", fmt.Errorf("DSATUTOR_TTS_VOICE: %w", err)
		}
		return path, nil
	}
	for _, dir := range voiceDirs {
		matches, _ := filepath.Glob(filepath.Join(dir, "*.onnx"))
		if len(matches) > 0 {
			sort.Strings(matches)
			return matches[0], nil
		}
	}
	return "", errors.New("no Piper voice model found: set DSATUTOR_TTS_VOICE to a .onnx voice, or download one with: python3 -m piper.download_voices en_US-lessac-medium --data-dir voices")
}

func voiceName(path string) string {
	if path == "" {
		return ""
	}
	return strings.TrimSuffix(filepath.Base(path), ".onnx")
}

func (p *Piper) startLocked() error {
	if p.cmd != nil {
		return nil
	}
	if p.scriptDir == "" {
		dir, err := os.MkdirTemp("", "dsatutor-tts-*")
		if err != nil {
			return fmt.Errorf("failed to create tts script directory: %w", err)
		}
		if err := os.WriteFile(filepath.Join(dir, "speaker.py"), speakerScript, 0600); err != nil {
			return fmt.Errorf("failed to write tts worker script: %w", err)
		}
		p.scriptDir = dir
	}

	cmd := exec.Command("python3", filepath.Join(p.scriptDir, "speaker.py"), p.voice)
	p.stderr = &bytes.Buffer{}
	cmd.Stderr = p.stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to open tts worker stdin: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to open tts worker stdout: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start tts worker: %w", err)
	}
	p.cmd, p.stdin, p.stdout = cmd, stdin, bufio.NewReader(stdout)

	// Wait for the ready line so a model that fails to load is reported now
	// rather than as a confusing mid-request error.
	ctx, cancel := context.WithTimeout(context.Background(), startTimeout)
	defer cancel()
	line, err := p.readLineLocked(ctx)
	if err != nil {
		tail := p.stderrTail()
		p.stopLocked()
		return fmt.Errorf("tts worker failed to load voice %s: %v%s", voiceName(p.voice), err, tail)
	}
	var ready struct {
		Ready bool `json:"ready"`
	}
	if err := json.Unmarshal(line, &ready); err != nil || !ready.Ready {
		p.stopLocked()
		return fmt.Errorf("tts worker sent unexpected startup output: %q", strings.TrimSpace(string(line)))
	}
	return nil
}

func (p *Piper) stopLocked() {
	if p.cmd == nil {
		return
	}
	p.stdin.Close()
	p.cmd.Process.Kill()
	p.cmd.Wait()
	p.cmd, p.stdin, p.stdout = nil, nil, nil
}

func (p *Piper) roundTripLocked(ctx context.Context, text string) ([]byte, error) {
	req, err := json.Marshal(map[string]string{"text": text})
	if err != nil {
		return nil, err
	}
	if _, err := p.stdin.Write(append(req, '\n')); err != nil {
		return nil, fmt.Errorf("tts worker write failed: %w%s", err, p.stderrTail())
	}
	line, err := p.readLineLocked(ctx)
	if err != nil {
		return nil, fmt.Errorf("tts worker read failed: %w%s", err, p.stderrTail())
	}
	var resp struct {
		Wav   string `json:"wav"`
		Error string `json:"error"`
	}
	if err := json.Unmarshal(line, &resp); err != nil {
		return nil, fmt.Errorf("tts worker sent a malformed response: %w", err)
	}
	if resp.Error != "" {
		return nil, fmt.Errorf("%w: %s", errSynthesis, resp.Error)
	}
	wav, err := base64.StdEncoding.DecodeString(resp.Wav)
	if err != nil || len(wav) == 0 {
		return nil, fmt.Errorf("tts worker sent unusable audio: %v", err)
	}
	return wav, nil
}

// readLineLocked reads one response line, killing the worker if ctx expires:
// once a request is abandoned mid-synthesis the stream is out of step, so the
// process must be replaced anyway.
func (p *Piper) readLineLocked(ctx context.Context) ([]byte, error) {
	type result struct {
		line []byte
		err  error
	}
	ch := make(chan result, 1)
	stdout := p.stdout
	go func() {
		line, err := stdout.ReadBytes('\n')
		ch <- result{line, err}
	}()
	select {
	case <-ctx.Done():
		p.cmd.Process.Kill()
		<-ch
		return nil, ctx.Err()
	case r := <-ch:
		return r.line, r.err
	}
}

func (p *Piper) stderrTail() string {
	s := strings.TrimSpace(p.stderr.String())
	if s == "" {
		return ""
	}
	if len(s) > 400 {
		s = "..." + s[len(s)-400:]
	}
	return " (stderr: " + s + ")"
}

// lastLine favours the tool's own message (e.g. the final line of a Python
// traceback) over Go's generic exit-status error.
func lastLine(out []byte, err error) string {
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if line := strings.TrimSpace(lines[len(lines)-1]); line != "" {
		return line
	}
	return err.Error()
}
