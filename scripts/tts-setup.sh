#!/usr/bin/env bash
# Set up everything lesson narration needs on this machine, detecting the OS
# and architecture and downloading only what is missing:
#
#   - Piper (always): a Python 3.12 venv with piper-tts plus a voice in
#     ./voices, so narration works offline on any machine.
#   - Ollama voices (best effort): a running Ollama >= 0.40.13 with
#     lfm3-audio pulled, and phonon on Apple Silicon (phonon needs MLX, so it
#     only runs in a native macOS Ollama, never in a Linux container).
#
# Narration is optional, so nothing here is fatal: a step that cannot finish
# prints what happened and the server simply offers fewer voices. Idempotent:
# run it as often as you like; it exits quickly once everything is in place.
# `imlazy serve` runs it automatically (lazy.toml); see docs/tts.md.
set -uo pipefail
cd "$(dirname "$0")/.."

OS=$(uname -s)   # Darwin or Linux
ARCH=$(uname -m) # arm64 / aarch64 / x86_64
note() { printf '[narration] %s\n' "$*"; }

# ver_ge A B: is version A >= B? (numeric per component, BSD and GNU sort)
ver_ge() {
    [ "$(printf '%s\n%s\n' "$2" "$1" | sort -t. -k1,1n -k2,2n -k3,3n | head -1)" = "$2" ]
}

# --- Piper: python environment ----------------------------------------------
export PATH="$HOME/.local/bin:$PATH"
if ! .venv/bin/python3 -c 'import piper' >/dev/null 2>&1; then
    if ! command -v uv >/dev/null 2>&1; then
        note "installing uv (manages the Python side of narration)"
        curl -LsSf https://astral.sh/uv/install.sh | sh >/dev/null 2>&1 \
            || note "could not install uv; skipping the Piper voice"
    fi
    if command -v uv >/dev/null 2>&1; then
        note "installing piper-tts into .venv (pyproject pins Python 3.12: the 3.14 wheels are broken on macOS)"
        uv sync >/dev/null || note "uv sync failed; skipping the Piper voice"
    fi
fi

# --- Piper: voice model ------------------------------------------------------
# Downloaded with curl rather than piper's own downloader: curl uses the
# system trust store, which avoids broken-certificate Python installs.
PIPER_VOICE=${DSATUTOR_PIPER_VOICE:-en_US-lessac-medium}
if ! ls voices/*.onnx >/dev/null 2>&1; then
    # en_US-lessac-medium -> en/en_US/lessac/medium on Hugging Face
    lang=${PIPER_VOICE%%-*}
    rest=${PIPER_VOICE#*-}
    speaker=${rest%%-*}
    quality=${rest#*-}
    base="https://huggingface.co/rhasspy/piper-voices/resolve/main/${lang%%_*}/${lang}/${speaker}/${quality}"
    note "downloading Piper voice ${PIPER_VOICE} (~60MB)"
    mkdir -p voices
    if ! curl -fsSL -o "voices/${PIPER_VOICE}.onnx" "${base}/${PIPER_VOICE}.onnx" ||
        ! curl -fsSL -o "voices/${PIPER_VOICE}.onnx.json" "${base}/${PIPER_VOICE}.onnx.json"; then
        rm -f "voices/${PIPER_VOICE}.onnx" "voices/${PIPER_VOICE}.onnx.json"
        note "voice download failed; Piper narration stays off until it succeeds"
    fi
fi

# --- Ollama: find or start a server ------------------------------------------
OLLAMA_URL=${DSATUTOR_OLLAMA_URL:-http://localhost:11434}
ollama_up() { curl -fsS --max-time 2 "$OLLAMA_URL/api/version" >/dev/null 2>&1; }

if ! ollama_up; then
    if command -v ollama >/dev/null 2>&1; then
        note "starting the native Ollama server (log: .ollama-serve.log)"
        nohup ollama serve >.ollama-serve.log 2>&1 &
        for _ in $(seq 1 20); do ollama_up && break; sleep 0.5; done
    elif command -v docker-compose >/dev/null 2>&1; then
        note "no native Ollama; starting the compose ollama service"
        docker-compose up -d ollama >/dev/null 2>&1
        for _ in $(seq 1 60); do ollama_up && break; sleep 1; done
    fi
fi

if ! ollama_up; then
    note "no Ollama available; narration uses the Piper voice (see docs/tts.md for the natural voices)"
    exit 0
fi

version=$(curl -fsS --max-time 5 "$OLLAMA_URL/api/version" | sed -n 's/.*"version":"\([^"]*\)".*/\1/p')
note "found Ollama ${version:-unknown} at ${OLLAMA_URL}"
if ! ver_ge "${version:-0}" 0.40.13; then
    note "speech models need Ollama >= 0.40.13; upgrade it for the natural voices"
    exit 0
fi

have_model() { curl -fsS --max-time 5 "$OLLAMA_URL/api/tags" | grep -q "\"name\":\"$1"; }
pull_model() {
    note "pulling $1 (one-time download, can take minutes)"
    curl -fsS -X POST "$OLLAMA_URL/api/pull" -d "{\"name\":\"$1\"}" -o /dev/null \
        || note "pull of $1 did not finish; narration falls back meanwhile"
}

have_model lfm3-audio || pull_model lfm3-audio
# phonon needs MLX: Apple Silicon only, and only in a native Ollama. Trying
# it against a container on a Mac just fails the pull, which is tolerated.
if [[ $OS == Darwin && $ARCH == arm64 ]]; then
    have_model phonon || pull_model phonon
fi

if have_model lfm3-audio || have_model phonon; then
    note "natural voices ready"
else
    note "no speech model installed yet; narration uses Piper until a pull succeeds"
fi
