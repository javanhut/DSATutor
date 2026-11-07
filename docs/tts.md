# Lesson Narration (Text-to-Speech)

Every lesson section (Overview, Understanding the Algorithm, When to Use It,
Worked Example, Steps to Remember, Check Your Understanding) has a **Listen**
button that reads the section aloud. Speech is synthesized on the server, so
the voice is identical wherever the container runs and nothing depends on the
browser's or OS's voices.

## Engines

Narration supports multiple engines, selected at runtime with the voice
picker in the chapter header (it appears when more than one engine works):

| Engine | Quality | Where it runs | Needs |
|--------|---------|---------------|-------|
| Ollama speech models — [lfm3-audio](https://ollama.com/library/lfm3-audio), [phonon](https://ollama.com/library/phonon) | natural | any Ollama ≥ 0.40.13 (lfm3-audio on Linux needs 0.40.13; phonon's best home is Apple Silicon) | `ollama pull lfm3-audio` and/or `ollama pull phonon` |
| [Piper](https://github.com/OHF-Voice/piper1-gpl) | robotic-but-clear | inside the app container, CPU | nothing — bundled in the image |

The server discovers installed Ollama speech models from `/api/tags` (checked
every 30s, so a newly pulled model shows up without a restart) and prefers
them; Piper is the always-there fallback. The first available engine is the
default voice: phonon when the server runs on macOS and phonon is installed,
otherwise the Ollama models in name order. If the engine a learner picked
fails mid-session, the server silently falls back to the next engine and
names the one actually used in the `X-TTS-Engine` response header.

## How it works

- `internal/tts` has three parts: `piper.go` keeps one persistent `python3`
  worker alive (the voice model loads once; after that CPU synthesis is
  faster than real time), `ollama.go` calls Ollama's OpenAI-compatible
  `POST /v1/audio/speech`, and `service.go` picks the engine per request,
  falls back, and caches WAVs in memory (lesson text is static, so replays
  are instant).
- The UI (`static/tts.js`) splits section text into sentence-sized chunks and
  prefetches the next chunk while the current one plays: audio starts within
  a couple of seconds and plays without gaps.
- Lessons get a Listen button per section. A practice problem gets one in the
  header of its problem panel, which reads the open tab: the Description
  (statement and examples), the revealed Hints, or the Solution. Code blocks,
  constraints and collapsed answers are skipped. The voice pickers move into
  the header of whichever view is on screen.
- If no engine works, `/api/tts/status` reports why per engine and the Listen
  buttons simply don't appear. The feature is fully optional.

## API

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/tts` | POST | `{"text": "...", "engine": "ollama:phonon:latest", "voice": "Zoey"}` (engine and voice optional, ≤ 2000 chars) → `audio/wav`, `X-TTS-Engine` header |
| `/api/tts/status` | GET | `{"available": bool, "default": "...", "engines": [{"id", "label", "available", "reason", "voices"}]}` |

### Named voices

Ollama speech models come with several named voices (phonon has Zoey, Alfie
and some 35 more; lfm3-audio has one male and one female voice per language).
The server reads them from Ollama's `/api/show` and lists them per engine in
`voices`; the page then shows a second picker beside the voice picker, and
remembers the choice per engine. Omitting `voice` selects the engine's
default, which is `DSATUTOR_TTS_OLLAMA_VOICE` when that is set. A voice the
engine does not have is ignored rather than rejected, so a fallback engine
or a remembered voice that no longer exists still produces audio.

## Docker Compose (recommended for the Ollama voices)

`docker-compose.yml` runs the app plus an `ollama` service. Pull the speech
models once (they persist in the `ollama-models` volume):

```bash
docker compose up -d --build
docker compose exec ollama ollama pull lfm3-audio   # 2.1 GB, 18 languages
docker compose exec ollama ollama pull phonon       # small, English voices
```

Reload the lesson page and the voice picker lists the new voices.

**macOS note**: containers get no GPU on a Mac, so the compose `ollama`
service runs the 2B lfm3-audio model on CPU. A natively installed
[Ollama.app](https://ollama.com/download) uses Metal and is much faster —
pull the models there instead and point the app at the host:

```bash
DSATUTOR_OLLAMA_URL=http://host.docker.internal:11434 docker compose up -d --build
```

(with podman, use `http://host.containers.internal:11434`). Phonon on Ollama
began as Apple-Silicon-only, so native Ollama is also the reliable way to get
the phonon voices on a Mac.

## Environment variables

| Variable | Default | Meaning |
|----------|---------|---------|
| `DSATUTOR_OLLAMA_URL` | `http://localhost:11434` | Ollama server for the speech models (compose sets `http://ollama:11434`) |
| `DSATUTOR_TTS_OLLAMA_VOICE` | *(model default)* | Voice name passed to Ollama when the learner has not picked one, e.g. `Zoey` for phonon |
| `DSATUTOR_TTS_OLLAMA_MODELS` | *(empty)* | Extra comma-separated model names to treat as speech models |
| `DSATUTOR_TTS_VOICE` | *(auto: `./voices`, `/app/voices`)* | Path to the Piper `.onnx` voice |

## Piper in the image

The image bundles Piper and the `en_US-lessac-medium` voice (roughly 150MB:
onnxruntime plus a 63MB voice). Pick a different voice from
[rhasspy/piper-voices](https://huggingface.co/rhasspy/piper-voices) at build
time:

```bash
docker build --build-arg PIPER_VOICE=en_GB-alba-medium -t dsatutor .
```

## Local development

One command sets everything up and runs the app:

```bash
imlazy serve      # or: make serve
```

Behind it, `scripts/tts-setup.sh` detects the OS and architecture and
downloads only what is missing, never failing the launch:

1. **Piper**: installs [uv](https://docs.astral.sh/uv/) if needed, `uv sync`s
   the pinned Python 3.12 venv with `piper-tts`, and curls the
   `en_US-lessac-medium` voice into `./voices` (override with
   `DSATUTOR_PIPER_VOICE`). curl is used instead of piper's downloader so a
   Python install with broken certificates can't block it.
2. **Ollama voices**: uses an Ollama already running on `localhost:11434`;
   otherwise starts the native one if installed, or the compose `ollama`
   service if `docker-compose` exists. It then pulls `lfm3-audio` (any
   platform) and `phonon` (Apple Silicon only — it needs MLX, so a pull into
   a Linux container is rejected by Ollama and tolerated by the script).
   An Ollama older than 0.40.13 gets a clear upgrade notice and is skipped.

`scripts/serve.sh` then runs the app with `.venv/bin` on PATH so the server
finds Piper. Run `scripts/tts-setup.sh` again any time; it is idempotent and
exits quickly once everything is in place. Without any engine the server
runs as before; the Listen buttons just stay hidden.

**Known issue**: the Piper 1.8 wheel for Python 3.14 mis-resolves espeak-ng's
data directory on macOS (`Error processing file '.../pip/phontab'`), which is
why `pyproject.toml` pins `requires-python = ">=3.12,<3.14"`. The Docker
image uses Python 3.12 and is unaffected.
