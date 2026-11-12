// Package tutor is the optional AI tutor: a chat with a local (or Ollama
// cloud) language model that sees the lesson or problem on screen, the
// learner's code and their last test run, and coaches rather than hands out
// answers. Speech-to-text for the microphone button also goes through
// Ollama, using any installed audio-capable model (lfm3-audio, gemma4).
//
// Like narration, the whole feature is optional: with no Ollama, or no chat
// model installed, /api/tutor/status says so and the UI hides the button.
package tutor

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	// modelsTTL bounds how stale the detected model list may be, so a model
	// pulled while the server runs appears without a restart.
	modelsTTL = 30 * time.Second
	// chatTimeout covers a whole streamed reply, including loading a large
	// model into memory on first use.
	chatTimeout = 10 * time.Minute
	// transcribeTimeout covers one recorded question.
	transcribeTimeout = 2 * time.Minute
)

var (
	ErrUnavailable  = errors.New("tutor unavailable")
	ErrUnknownModel = errors.New("unknown tutor model")
)

// Model is one installed Ollama model and what it can do.
type Model struct {
	Name     string `json:"name"`
	Thinking bool   `json:"thinking"` // supports a separate reasoning phase
	Audio    bool   `json:"-"`        // accepts audio input (transcription)
	Cloud    bool   `json:"cloud"`    // runs on ollama.com, not this machine
	Tools    bool   `json:"tools"`    // can call tools: point at the page, edit code
}

// Message is one chat turn. ToolCalls and ToolName only appear in the
// server's tool loop, never in what the browser sends.
type Message struct {
	Role      string     `json:"role"` // system, user, assistant or tool
	Content   string     `json:"content"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	ToolName  string     `json:"tool_name,omitempty"`
}

// Client talks to one Ollama server. Create it with New.
type Client struct {
	baseURL   string
	prefer    string // DSATUTOR_TUTOR_MODEL: the default chat model when installed
	sttPrefer string // DSATUTOR_STT_MODEL: the transcription model when installed
	http      *http.Client

	mu        sync.Mutex
	lastCheck time.Time
	chat      []Model // chat-capable models, default first
	audio     []Model // audio-input models usable for transcription
	reason    string  // why chat is empty
}

// New reads DSATUTOR_OLLAMA_URL (default http://localhost:11434), the same
// setting narration uses, plus the optional DSATUTOR_TUTOR_MODEL and
// DSATUTOR_STT_MODEL overrides.
func New() *Client {
	url := os.Getenv("DSATUTOR_OLLAMA_URL")
	if url == "" {
		url = "http://localhost:11434"
	}
	return &Client{
		baseURL:   strings.TrimRight(url, "/"),
		prefer:    strings.TrimSpace(os.Getenv("DSATUTOR_TUTOR_MODEL")),
		sttPrefer: strings.TrimSpace(os.Getenv("DSATUTOR_STT_MODEL")),
		// No client-wide timeout: replies stream for as long as the model
		// talks. Each call bounds itself with a context instead.
		http: &http.Client{},
	}
}

// Status is the /api/tutor/status payload.
type Status struct {
	Available  bool    `json:"available"`
	Default    string  `json:"default,omitempty"`
	Models     []Model `json:"models"`
	Transcribe bool    `json:"transcribe"` // the microphone button works
	Reason     string  `json:"reason,omitempty"`
}

func (c *Client) Status() Status {
	chat, audio, reason := c.models()
	st := Status{Available: len(chat) > 0, Models: chat, Transcribe: len(audio) > 0, Reason: reason}
	if st.Models == nil {
		st.Models = []Model{}
	}
	if len(chat) > 0 {
		st.Default = chat[0].Name
	}
	return st
}

func (c *Client) models() (chat, audio []Model, reason string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if time.Since(c.lastCheck) < modelsTTL {
		return c.chat, c.audio, c.reason
	}
	c.lastCheck = time.Now()
	c.chat, c.audio, c.reason = c.fetchModels()
	return c.chat, c.audio, c.reason
}

func isModel(name, base string) bool {
	return name == base || strings.HasPrefix(name, base+":")
}

func (c *Client) fetchModels() (chat, audio []Model, reason string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var tags struct {
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}
	if err := c.getJSON(ctx, "/api/tags", &tags); err != nil {
		return nil, nil, fmt.Sprintf("ollama unreachable at %s: %v", c.baseURL, err)
	}
	for _, t := range tags.Models {
		var show struct {
			Capabilities []string `json:"capabilities"`
		}
		if err := c.postJSON(ctx, "/api/show", map[string]string{"model": t.Name}, &show); err != nil {
			continue
		}
		caps := map[string]bool{}
		for _, cp := range show.Capabilities {
			caps[cp] = true
		}
		m := Model{
			Name:     t.Name,
			Thinking: caps["thinking"],
			Audio:    caps["audio"],
			Cloud:    strings.HasSuffix(t.Name, ":cloud") || strings.HasSuffix(t.Name, "-cloud"),
			Tools:    caps["tools"],
		}
		// Speech models (lfm3-audio, phonon) are narration voices; they can
		// technically complete text but make poor tutors.
		if caps["completion"] && !caps["speech"] && !caps["embedding"] {
			chat = append(chat, m)
		}
		if m.Audio && !m.Cloud {
			audio = append(audio, m)
		}
	}
	// Default order: the configured model, then local models, then cloud
	// ones, each by name.
	sort.SliceStable(chat, func(i, j int) bool {
		pi, pj := c.prefer != "" && isModel(chat[i].Name, c.prefer), c.prefer != "" && isModel(chat[j].Name, c.prefer)
		if pi != pj {
			return pi
		}
		if chat[i].Cloud != chat[j].Cloud {
			return !chat[i].Cloud
		}
		return chat[i].Name < chat[j].Name
	})
	sort.SliceStable(audio, func(i, j int) bool {
		pi, pj := c.sttPrefer != "" && isModel(audio[i].Name, c.sttPrefer), c.sttPrefer != "" && isModel(audio[j].Name, c.sttPrefer)
		if pi != pj {
			return pi
		}
		return audio[i].Name < audio[j].Name
	})
	if len(chat) == 0 {
		reason = fmt.Sprintf("no chat model installed on ollama at %s (try: ollama pull gemma4)", c.baseURL)
	}
	return chat, audio, reason
}

// resolve picks the chat model to use: the requested one when installed,
// the default when none was requested.
func (c *Client) resolve(name string) (Model, error) {
	chat, _, reason := c.models()
	if len(chat) == 0 {
		return Model{}, fmt.Errorf("%w: %s", ErrUnavailable, reason)
	}
	if name == "" {
		return chat[0], nil
	}
	for _, m := range chat {
		if m.Name == name {
			return m, nil
		}
	}
	return Model{}, fmt.Errorf("%w: %q", ErrUnknownModel, name)
}

// Delta is one streamed piece of a reply. Thinking pieces are the model's
// reasoning (only when think is on); the rest is the answer itself.
type Delta struct {
	Thinking string
	Content  string
}

// Reply is one model turn: its text and any tool calls it made.
type Reply struct {
	Model     string
	Content   string
	ToolCalls []ToolCall
}

// SupportsTools reports whether model ("" for the default) can call tools.
func (c *Client) SupportsTools(model string) bool {
	m, err := c.resolve(model)
	return err == nil && m.Tools
}

// Chat streams one reply from model ("" for the default) to messages,
// calling emit for each piece of text as it arrives. think asks
// thinking-capable models to reason first; it is ignored by the others.
// tools are offered only to models that support them.
func (c *Client) Chat(ctx context.Context, model string, messages []Message, think bool, tools []Tool, emit func(Delta) error) (Reply, error) {
	m, err := c.resolve(model)
	if err != nil {
		return Reply{}, err
	}
	reply := Reply{Model: m.Name}
	body := map[string]any{
		"model":    m.Name,
		"messages": messages,
		"stream":   true,
		"options":  map[string]any{"temperature": 0.4},
	}
	if m.Tools && len(tools) > 0 {
		body["tools"] = tools
	}
	// Thinking models reason by default; say explicitly what we want, but
	// never send the flag to a model that rejects it.
	if m.Thinking {
		body["think"] = think
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return reply, err
	}
	ctx, cancel := context.WithTimeout(ctx, chatTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/api/chat", bytes.NewReader(payload))
	if err != nil {
		return reply, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return reply, fmt.Errorf("%w: ollama chat request failed: %v", ErrUnavailable, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 400))
		return reply, fmt.Errorf("ollama %s: %s", resp.Status, strings.TrimSpace(string(msg)))
	}
	var content strings.Builder
	sc := bufio.NewScanner(resp.Body)
	sc.Buffer(make([]byte, 64*1024), 4*1024*1024)
	for sc.Scan() {
		var chunk struct {
			Message struct {
				Content   string     `json:"content"`
				Thinking  string     `json:"thinking"`
				ToolCalls []ToolCall `json:"tool_calls"`
			} `json:"message"`
			Done  bool   `json:"done"`
			Error string `json:"error"`
		}
		if err := json.Unmarshal(sc.Bytes(), &chunk); err != nil {
			continue
		}
		if chunk.Error != "" {
			return reply, errors.New(chunk.Error)
		}
		reply.ToolCalls = append(reply.ToolCalls, chunk.Message.ToolCalls...)
		content.WriteString(chunk.Message.Content)
		if chunk.Message.Content != "" || chunk.Message.Thinking != "" {
			if err := emit(Delta{Thinking: chunk.Message.Thinking, Content: chunk.Message.Content}); err != nil {
				return reply, err
			}
		}
		if chunk.Done {
			break
		}
	}
	reply.Content = content.String()
	if err := sc.Err(); err != nil {
		return reply, fmt.Errorf("ollama chat stream broke: %w", err)
	}
	return reply, nil
}

// Transcribe turns a recorded question (a WAV file) into text with an
// installed audio-capable model. hint names a model to prefer, typically
// the chat model in use: when it also takes audio, Ollama doesn't have to
// load a second model.
func (c *Client) Transcribe(ctx context.Context, wav []byte, hint string) (string, error) {
	_, audio, _ := c.models()
	if len(audio) == 0 {
		return "", fmt.Errorf("%w: no audio-capable model installed (try: ollama pull lfm3-audio)", ErrUnavailable)
	}
	model := audio[0].Name
	if c.sttPrefer == "" {
		for _, m := range audio {
			if m.Name == hint {
				model = m.Name
			}
		}
	}

	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	if err := mw.WriteField("model", model); err != nil {
		return "", err
	}
	fw, err := mw.CreateFormFile("file", "question.wav")
	if err != nil {
		return "", err
	}
	if _, err := fw.Write(wav); err != nil {
		return "", err
	}
	if err := mw.Close(); err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(ctx, transcribeTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/v1/audio/transcriptions", &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: ollama transcription request failed: %v", ErrUnavailable, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 400))
		return "", fmt.Errorf("ollama %s: %s", resp.Status, strings.TrimSpace(string(msg)))
	}
	var out struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", fmt.Errorf("ollama sent a malformed transcription: %w", err)
	}
	return strings.TrimSpace(out.Text), nil
}

func (c *Client) getJSON(ctx context.Context, path string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return err
	}
	return c.do(req, out)
}

func (c *Client) postJSON(ctx context.Context, path string, in, out any) error {
	payload, err := json.Marshal(in)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(req, out)
}

func (c *Client) do(req *http.Request, out any) error {
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s returned %s", req.URL.Path, resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
