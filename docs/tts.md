# Lesson Narration (Text-to-Speech)

Every lesson section (Overview, Understanding the Algorithm, When to Use It,
Worked Example, Steps to Remember, Check Your Understanding) has a **Listen**
button that reads the section aloud. Speech is synthesized on the server by
[Piper](https://github.com/OHF-Voice/piper1-gpl), a free neural TTS engine
that runs on CPU, so the voice is identical wherever the container runs and
nothing depends on the browser's or OS's voices.

## How it works

- `internal/tts` keeps one persistent `python3` worker (`speaker.py`) alive.
  The voice model loads once; after that, synthesis on a CPU is faster than
  real time. Synthesized WAVs are cached in memory (lesson text is static),
  so replays are instant.
- The UI (`static/tts.js`) splits section text into sentence-sized chunks and
  prefetches the next chunk while the current one plays: audio starts within
  a couple of seconds and plays without gaps.
- If no voice is installed, `/api/tts/status` reports why and the Listen
  buttons simply don't appear. The feature is fully optional.

## API

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/tts` | POST | `{"text": "..."}` (≤ 2000 chars) → `audio/wav` |
| `/api/tts/status` | GET | `{"available": bool, "voice": "...", "reason": "..."}` |

## Docker

The image includes Piper and the `en_US-lessac-medium` voice by default
(roughly 150MB extra: onnxruntime plus a 63MB voice). Pick a different voice
from [rhasspy/piper-voices](https://huggingface.co/rhasspy/piper-voices) at
build time:

```bash
docker build --build-arg PIPER_VOICE=en_GB-alba-medium -t dsatutor .
```

## Local development

The server looks for a voice at `$DSATUTOR_TTS_VOICE` (a `.onnx` path), then
in `./voices` and `/app/voices`. To enable narration on a local checkout:

```bash
pip install piper-tts
python3 -m piper.download_voices en_US-lessac-medium --data-dir voices
imlazy serve
```

Without Piper the server runs as before; the buttons just stay hidden.

**Known issue**: the Piper 1.8 wheel for Python 3.14 mis-resolves espeak-ng's
data directory on macOS (`Error processing file '.../pip/phontab'`). Use
Python ≤ 3.13 locally; the Docker image uses Python 3.12 and is unaffected.
