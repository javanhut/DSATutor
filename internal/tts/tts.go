// Package tts synthesizes speech for lesson text with Piper
// (https://github.com/OHF-Voice/piper1-gpl), so narration is rendered
// server-side and sounds the same wherever the container runs, with no
// dependence on the client's OS voices. The voice model is loaded once in a
// persistent python3 worker (speaker.py); after that, synthesis on a CPU is
// faster than real time.
package tts

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
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

// MaxTextLen bounds one synthesis request; the UI sends lesson text in
// sentence-sized chunks well under this.
const MaxTextLen = 2000

const (
	// startTimeout covers loading the voice model (~60MB) on first use.
	startTimeout = 60 * time.Second
	synthTimeout = 60 * time.Second
	// maxCacheBytes caps the in-memory WAV cache. Lesson text is static, so
	// replays hit the cache instead of re-synthesizing.
	maxCacheBytes = 64 << 20
)

// voiceDirs are searched for a *.onnx voice when DSATUTOR_TTS_VOICE is not
// set; /app/voices is where the Dockerfile puts the downloaded voice.
var voiceDirs = []string{"/app/voices", "voices"}

var (
	ErrUnavailable = errors.New("text-to-speech is not available on this server")
	ErrInvalidText = errors.New("text must be non-empty and under 2000 characters")

	// errSynthesis marks a failure reported by a healthy worker, so the
	// retry loop doesn't restart the process for text it can't speak anyway.
	errSynthesis = errors.New("piper could not synthesize this text")
)

// Engine owns the worker process and the audio cache. The zero value is not
// usable; create one with NewEngine.
type Engine struct {
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

	cacheMu    sync.Mutex
	cache      map[string][]byte
	cacheOrder []string
	cacheSize  int
	cacheLimit int
}

func NewEngine() *Engine {
	return &Engine{cache: map[string][]byte{}, cacheLimit: maxCacheBytes}
}

// Status reports whether synthesis can work here and, if not, why. The check
// (a voice model on disk, python3 able to import piper) runs once and is
// remembered; images built from the Dockerfile always pass it, while local
// checkouts without Piper just hide the feature in the UI.
func (e *Engine) Status() (available bool, voice, reason string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.checkLocked()
	return e.available, voiceName(e.voice), e.reason
}

// Synthesize returns a WAV rendering of text, serving repeats from cache.
// Requests are serialized: Piper is CPU-bound, so running them one at a time
// is as fast as interleaving and keeps the worker protocol trivial.
func (e *Engine) Synthesize(ctx context.Context, text string) ([]byte, error) {
	text = strings.TrimSpace(text)
	if text == "" || len(text) > MaxTextLen {
		return nil, ErrInvalidText
	}
	key := fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
	if wav, ok := e.cacheGet(key); ok {
		return wav, nil
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	e.checkLocked()
	if !e.available {
		return nil, ErrUnavailable
	}
	ctx, cancel := context.WithTimeout(ctx, synthTimeout)
	defer cancel()

	// One retry after a restart, in case the worker died since its last use.
	var wav []byte
	var err error
	for attempt := 0; attempt < 2; attempt++ {
		if err = e.startLocked(); err != nil {
			return nil, err
		}
		wav, err = e.roundTripLocked(ctx, text)
		if err == nil || errors.Is(err, errSynthesis) {
			break
		}
		e.stopLocked()
		if ctx.Err() != nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	e.cachePut(key, wav)
	return wav, nil
}

func (e *Engine) checkLocked() {
	if e.checked {
		return
	}
	e.checked = true
	voice, err := findVoice()
	if err != nil {
		e.reason = err.Error()
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if out, err := exec.CommandContext(ctx, "python3", "-c", "import piper").CombinedOutput(); err != nil {
		e.reason = fmt.Sprintf("python3 cannot import piper (pip install piper-tts): %s", lastLine(out, err))
		return
	}
	e.voice = voice
	e.available = true
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

func (e *Engine) startLocked() error {
	if e.cmd != nil {
		return nil
	}
	if e.scriptDir == "" {
		dir, err := os.MkdirTemp("", "dsatutor-tts-*")
		if err != nil {
			return fmt.Errorf("failed to create tts script directory: %w", err)
		}
		if err := os.WriteFile(filepath.Join(dir, "speaker.py"), speakerScript, 0600); err != nil {
			return fmt.Errorf("failed to write tts worker script: %w", err)
		}
		e.scriptDir = dir
	}

	cmd := exec.Command("python3", filepath.Join(e.scriptDir, "speaker.py"), e.voice)
	e.stderr = &bytes.Buffer{}
	cmd.Stderr = e.stderr
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
	e.cmd, e.stdin, e.stdout = cmd, stdin, bufio.NewReader(stdout)

	// Wait for the ready line so a model that fails to load is reported now
	// rather than as a confusing mid-request error.
	ctx, cancel := context.WithTimeout(context.Background(), startTimeout)
	defer cancel()
	line, err := e.readLineLocked(ctx)
	if err != nil {
		tail := e.stderrTail()
		e.stopLocked()
		return fmt.Errorf("tts worker failed to load voice %s: %v%s", voiceName(e.voice), err, tail)
	}
	var ready struct {
		Ready bool `json:"ready"`
	}
	if err := json.Unmarshal(line, &ready); err != nil || !ready.Ready {
		e.stopLocked()
		return fmt.Errorf("tts worker sent unexpected startup output: %q", strings.TrimSpace(string(line)))
	}
	return nil
}

func (e *Engine) stopLocked() {
	if e.cmd == nil {
		return
	}
	e.stdin.Close()
	e.cmd.Process.Kill()
	e.cmd.Wait()
	e.cmd, e.stdin, e.stdout = nil, nil, nil
}

func (e *Engine) roundTripLocked(ctx context.Context, text string) ([]byte, error) {
	req, err := json.Marshal(map[string]string{"text": text})
	if err != nil {
		return nil, err
	}
	if _, err := e.stdin.Write(append(req, '\n')); err != nil {
		return nil, fmt.Errorf("tts worker write failed: %w%s", err, e.stderrTail())
	}
	line, err := e.readLineLocked(ctx)
	if err != nil {
		return nil, fmt.Errorf("tts worker read failed: %w%s", err, e.stderrTail())
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
func (e *Engine) readLineLocked(ctx context.Context) ([]byte, error) {
	type result struct {
		line []byte
		err  error
	}
	ch := make(chan result, 1)
	stdout := e.stdout
	go func() {
		line, err := stdout.ReadBytes('\n')
		ch <- result{line, err}
	}()
	select {
	case <-ctx.Done():
		e.cmd.Process.Kill()
		<-ch
		return nil, ctx.Err()
	case r := <-ch:
		return r.line, r.err
	}
}

func (e *Engine) stderrTail() string {
	s := strings.TrimSpace(e.stderr.String())
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

func (e *Engine) cacheGet(key string) ([]byte, bool) {
	e.cacheMu.Lock()
	defer e.cacheMu.Unlock()
	wav, ok := e.cache[key]
	return wav, ok
}

// cachePut evicts oldest-first; lesson narration is replayed far more often
// than it is varied, so simple FIFO is enough.
func (e *Engine) cachePut(key string, wav []byte) {
	e.cacheMu.Lock()
	defer e.cacheMu.Unlock()
	if _, ok := e.cache[key]; ok {
		return
	}
	e.cache[key] = wav
	e.cacheOrder = append(e.cacheOrder, key)
	e.cacheSize += len(wav)
	for e.cacheSize > e.cacheLimit && len(e.cacheOrder) > 0 {
		oldest := e.cacheOrder[0]
		e.cacheOrder = e.cacheOrder[1:]
		e.cacheSize -= len(e.cache[oldest])
		delete(e.cache, oldest)
	}
}
