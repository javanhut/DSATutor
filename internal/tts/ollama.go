// Ollama is the optional high-quality narration backend: speech models such
// as lfm3-audio and phonon served by an Ollama instance (Ollama >= 0.40.13)
// through its OpenAI-compatible /v1/audio/speech endpoint. Installed speech
// models are discovered from /api/tags, so pulling a new model makes it show
// up as a voice choice without reconfiguring anything.
package tts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	// ollamaSynthTimeout is generous: lfm3-audio is a 2B model and can run
	// CPU-only inside the compose ollama service.
	ollamaSynthTimeout = 120 * time.Second
	// modelsTTL bounds how stale the detected model list may be, so a model
	// pulled while the server runs appears without a restart.
	modelsTTL = 30 * time.Second
)

// speechModelPrefixes are the Ollama models known to accept /v1/audio/speech.
// DSATUTOR_TTS_OLLAMA_MODELS extends the list for future models.
var speechModelPrefixes = []string{"lfm3-audio", "phonon"}

// Ollama talks to one Ollama server. The zero value is not usable; create
// one with NewOllama.
type Ollama struct {
	baseURL string
	voice   string // optional, e.g. "Zoey" for phonon; empty lets the model pick
	client  *http.Client

	mu        sync.Mutex
	lastCheck time.Time
	models    []string // detected speech models, sorted
	reason    string   // why empty, surfaced by /api/tts/status
}

// NewOllama reads its configuration from the environment:
// DSATUTOR_OLLAMA_URL (default http://localhost:11434, the compose file sets
// http://ollama:11434) and DSATUTOR_TTS_OLLAMA_VOICE.
func NewOllama() *Ollama {
	url := os.Getenv("DSATUTOR_OLLAMA_URL")
	if url == "" {
		url = "http://localhost:11434"
	}
	return &Ollama{
		baseURL: strings.TrimRight(url, "/"),
		voice:   os.Getenv("DSATUTOR_TTS_OLLAMA_VOICE"),
		client:  &http.Client{Timeout: ollamaSynthTimeout},
	}
}

// SpeechModels returns the installed speech-capable models, or a reason why
// there are none. The result is cached briefly so lesson pages polling
// /api/tts/status don't hammer Ollama.
func (o *Ollama) SpeechModels() (models []string, reason string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if time.Since(o.lastCheck) < modelsTTL {
		return o.models, o.reason
	}
	o.lastCheck = time.Now()
	o.models, o.reason = o.fetchModels()
	return o.models, o.reason
}

func (o *Ollama) fetchModels() ([]string, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, o.baseURL+"/api/tags", nil)
	if err != nil {
		return nil, err.Error()
	}
	resp, err := o.client.Do(req)
	if err != nil {
		return nil, fmt.Sprintf("ollama unreachable at %s: %v", o.baseURL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Sprintf("ollama at %s returned %s for /api/tags", o.baseURL, resp.Status)
	}
	var tags struct {
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return nil, fmt.Sprintf("ollama sent malformed /api/tags response: %v", err)
	}

	prefixes := speechModelPrefixes
	if extra := os.Getenv("DSATUTOR_TTS_OLLAMA_MODELS"); extra != "" {
		for _, name := range strings.Split(extra, ",") {
			if name = strings.TrimSpace(name); name != "" {
				prefixes = append(prefixes, name)
			}
		}
	}
	var models []string
	for _, m := range tags.Models {
		for _, prefix := range prefixes {
			if m.Name == prefix || strings.HasPrefix(m.Name, prefix+":") {
				models = append(models, m.Name)
				break
			}
		}
	}
	if len(models) == 0 {
		return nil, fmt.Sprintf("no speech models installed on ollama at %s (try: ollama pull lfm3-audio, or ollama pull phonon)", o.baseURL)
	}
	sort.Strings(models)
	return models, ""
}

// Synthesize asks Ollama's OpenAI-compatible speech endpoint for a WAV
// rendering of text with the given model.
func (o *Ollama) Synthesize(ctx context.Context, model, text string) ([]byte, error) {
	body := map[string]string{
		"model":           model,
		"input":           text,
		"response_format": "wav",
	}
	if o.voice != "" {
		body["voice"] = o.voice
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, ollamaSynthTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+"/v1/audio/speech", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := o.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ollama speech request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 400))
		return nil, fmt.Errorf("ollama %s: %s", resp.Status, strings.TrimSpace(string(msg)))
	}
	wav, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ollama speech response truncated: %w", err)
	}
	if len(wav) == 0 {
		return nil, fmt.Errorf("ollama returned empty audio for model %s", model)
	}
	return wav, nil
}
