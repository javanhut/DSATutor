#!/usr/bin/env bash
# Run the web app with lesson narration wired up: the Go server finds Piper
# through python3 on PATH, so the project's .venv (created by
# scripts/tts-setup.sh) is prepended when it exists.
set -euo pipefail
cd "$(dirname "$0")/.."

[[ -x .venv/bin/python3 ]] && export PATH="$PWD/.venv/bin:$PATH"
exec go run ./cmd/tutor -serve -addr="${DSATUTOR_ADDR:-:8080}"
