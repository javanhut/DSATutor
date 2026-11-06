// Service picks a narration engine per request: Ollama speech models when an
// Ollama server has some installed (better voices), Piper as the always-there
// fallback. The UI selects an engine at runtime; if the chosen one fails the
// remaining engines are tried in order, so narration degrades instead of
// breaking.
package tts

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"sync"
)

// MaxTextLen bounds one synthesis request; the UI sends lesson text in
// sentence-sized chunks well under this.
const MaxTextLen = 2000

// maxCacheBytes caps the in-memory WAV cache. Lesson text is static, so
// replays hit the cache instead of re-synthesizing.
const maxCacheBytes = 64 << 20

var (
	ErrUnavailable   = errors.New("text-to-speech is not available on this server")
	ErrInvalidText   = errors.New("text must be non-empty and under 2000 characters")
	ErrUnknownEngine = errors.New("unknown narration engine")
)

// EngineStatus describes one selectable narration engine.
type EngineStatus struct {
	ID        string `json:"id"`    // e.g. "ollama:lfm3-audio" or "piper"
	Label     string `json:"label"` // shown in the UI's voice picker
	Available bool   `json:"available"`
	Reason    string `json:"reason,omitempty"`
}

// Status is the /api/tts/status payload.
type Status struct {
	Available bool           `json:"available"` // any engine usable
	Default   string         `json:"default"`   // engine used when none is requested
	Engines   []EngineStatus `json:"engines"`
}

// Service owns the engines and the audio cache. The zero value is not
// usable; create one with NewService.
type Service struct {
	piper  *Piper
	ollama *Ollama

	cacheMu    sync.Mutex
	cache      map[string][]byte
	cacheOrder []string
	cacheSize  int
	cacheLimit int
}

func NewService() *Service {
	return &Service{
		piper:      NewPiper(),
		ollama:     NewOllama(),
		cache:      map[string][]byte{},
		cacheLimit: maxCacheBytes,
	}
}

// Status lists every engine in preference order: Ollama models first (they
// are only listed when installed, and sound better), Piper last.
func (s *Service) Status() Status {
	var st Status
	models, reason := s.ollama.SpeechModels()
	if len(models) == 0 {
		st.Engines = append(st.Engines, EngineStatus{
			ID:     "ollama",
			Label:  "Ollama",
			Reason: reason,
		})
	}
	for _, model := range models {
		st.Engines = append(st.Engines, EngineStatus{
			ID:        "ollama:" + model,
			Label:     "Ollama · " + strings.TrimSuffix(model, ":latest"),
			Available: true,
		})
	}
	available, voice, reason := s.piper.Status()
	label := "Piper"
	if voice != "" {
		label += " · " + voice
	}
	st.Engines = append(st.Engines, EngineStatus{
		ID:        "piper",
		Label:     label,
		Available: available,
		Reason:    reason,
	})
	for _, e := range st.Engines {
		if e.Available {
			st.Available = true
			if st.Default == "" {
				st.Default = e.ID
			}
		}
	}
	return st
}

// Synthesize renders text with the requested engine ("" means the default),
// falling back to the remaining engines in preference order if it fails.
// It returns the audio and the id of the engine that produced it.
func (s *Service) Synthesize(ctx context.Context, text, engine string) ([]byte, string, error) {
	text = strings.TrimSpace(text)
	if text == "" || len(text) > MaxTextLen {
		return nil, "", ErrInvalidText
	}

	candidates, err := s.candidates(engine)
	if err != nil {
		return nil, "", err
	}
	if len(candidates) == 0 {
		return nil, "", ErrUnavailable
	}

	var lastErr error
	for _, id := range candidates {
		key := fmt.Sprintf("%s|%x", id, sha256.Sum256([]byte(text)))
		if wav, ok := s.cacheGet(key); ok {
			return wav, id, nil
		}
		var wav []byte
		if model, ok := strings.CutPrefix(id, "ollama:"); ok {
			wav, err = s.ollama.Synthesize(ctx, model, text)
		} else {
			wav, err = s.piper.Synthesize(ctx, text)
		}
		if err == nil {
			s.cachePut(key, wav)
			return wav, id, nil
		}
		lastErr = fmt.Errorf("%s: %w", id, err)
		if ctx.Err() != nil {
			break
		}
	}
	return nil, "", lastErr
}

// candidates resolves the requested engine to a try-order: the request
// first, then every other available engine as fallback. A known engine that
// is currently down falls back silently; only an id the server has never
// heard of (e.g. a model that was removed) errors, which the frontend
// recovers from by refetching /api/tts/status.
func (s *Service) candidates(engine string) ([]string, error) {
	known := engine == ""
	var order []string
	for _, e := range s.Status().Engines {
		if e.ID == engine {
			known = true
		}
		if !e.Available {
			continue
		}
		if e.ID == engine {
			order = append([]string{e.ID}, order...)
		} else {
			order = append(order, e.ID)
		}
	}
	if !known {
		return nil, fmt.Errorf("%w: %s", ErrUnknownEngine, engine)
	}
	return order, nil
}

func (s *Service) cacheGet(key string) ([]byte, bool) {
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()
	wav, ok := s.cache[key]
	return wav, ok
}

// cachePut evicts oldest-first; lesson narration is replayed far more often
// than it is varied, so simple FIFO is enough.
func (s *Service) cachePut(key string, wav []byte) {
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()
	if _, ok := s.cache[key]; ok {
		return
	}
	s.cache[key] = wav
	s.cacheOrder = append(s.cacheOrder, key)
	s.cacheSize += len(wav)
	for s.cacheSize > s.cacheLimit && len(s.cacheOrder) > 0 {
		oldest := s.cacheOrder[0]
		s.cacheOrder = s.cacheOrder[1:]
		s.cacheSize -= len(s.cache[oldest])
		delete(s.cache, oldest)
	}
}
