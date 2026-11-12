# AI Tutor

An optional coach you can talk to while you read a chapter or work a
problem. It runs on [Ollama](https://ollama.com), the same server that
provides the narration voices, so nothing leaves your machine unless you pick
an Ollama cloud model.

## Using it

- Press **Tutor** in the sidebar, or **Ask tutor** in a practice problem's
  header. It never opens by itself. Esc closes it.
- Type a question (Enter sends, Shift+Enter adds a line), or press one of
  the quick prompts: *Where does my logic break?*, *Give me a nudge*, *Check
  my complexity*, *Better approach?* on a problem, or *Quiz me* and *Explain it
  simpler* on a chapter.
- **Microphone**: click it, ask your question, click again to send. The
  recording is transcribed on the server and sent. A question you ask by
  voice gets a short spoken answer, so you can go back and forth hands-free.
  Esc cancels a recording.
- **Speak replies** reads every answer aloud with the narration voice you
  picked for lessons (phonon, lfm3-audio or Piper). Speech starts on the
  first finished sentence, while the rest of the answer is still being
  written. Code blocks are skipped.
- **Think deeper** (models that support it) lets the model reason before
  answering. It's slower, but better at finding subtle bugs.
- Each problem, chapter and the sandbox keeps its own conversation, saved in
  the browser. **Clear** starts it over.

## What the tutor sees

Every message carries what is on screen, and the server fills in the rest
from the real course data:

| View | Context |
|------|---------|
| Practice problem | statement, constraints, examples, the hints you've revealed, your code (with line numbers), your last Run or Submit result, plus reference notes: brute force, key insight, optimal complexity, common mistakes |
| Chapter | title, summary, objectives, when to use the pattern, the concepts with their intuition and common mistakes |
| Sandbox | your code and its output |

It's told to coach, not solve: ask a guiding question, give the smallest
hint that unblocks you, cite line numbers, give an input that breaks your
code, and make you state the complexity before confirming it. It writes a
full solution only if you ask for one. The model solution's code goes into
the prompt only after you've solved the problem or opened the solution, so
a small model can't leak it early.

## Models

Any installed Ollama model that can chat is offered in the panel's model
picker. Speech-only models (phonon) and narration models (lfm3-audio) are
left out. Local models come first, then cloud ones (`*:cloud`). The default
is `DSATUTOR_TUTOR_MODEL` when that's set and installed, otherwise the first
local model by name.

```bash
ollama pull gemma4        # 8B, good all-rounder, also transcribes speech
```

Transcription for the microphone uses Ollama's `/v1/audio/transcriptions`
with an audio-capable local model. It prefers the chat model in use when that
model takes audio (gemma4 does), so Ollama doesn't have to load a second
model. Otherwise it uses lfm3-audio, which `scripts/tts-setup.sh` already
installs for narration. `DSATUTOR_STT_MODEL` forces a specific model.

| Variable | Default | Purpose |
|----------|---------|---------|
| `DSATUTOR_OLLAMA_URL` | `http://localhost:11434` | the Ollama server (shared with narration) |
| `DSATUTOR_TUTOR_MODEL` | first local chat model | default chat model |
| `DSATUTOR_STT_MODEL` | chat model if it takes audio, else first audio model | transcription model |

With no Ollama or no chat model, the Tutor buttons don't appear.
`/api/tutor/status` says why.

The microphone needs a secure page: `http://localhost` works, but a plain
`http://` LAN address doesn't (the browser blocks `getUserMedia`).

## API

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/tutor/status` | GET | `{"available", "default", "models": [{"name", "thinking", "cloud"}], "transcribe", "reason"}` |
| `/api/tutor/chat` | POST | `{"model", "think", "messages": [{"role", "content"}], "context": {"view", "problemId", "chapterNumber", "language", "code", "lastRun", "hintsRevealed", "solutionUnlocked", "voice"}}` → NDJSON stream of `{"type":"thinking"\|"delta","text"}` then `{"type":"done","model"}` or `{"type":"error","error"}` |
| `/api/tutor/transcribe` | POST | body: a WAV recording, `?model=` names the chat model in use → `{"text"}` |

The server builds the system prompt itself (`internal/tutor/prompt.go`).
System messages from the browser are dropped, and history is capped at the
last 24 turns.
