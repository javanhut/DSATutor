#!/usr/bin/env bash
# Make sure the tools DSA Tutor needs are installed, and install whatever is
# missing with rvn. Every imlazy task that builds, runs or tests DSA Tutor
# depends on this, so a fresh checkout needs nothing done by hand.
#
#   scripts/system-deps.sh           install what is missing
#   scripts/system-deps.sh --check   only report; exit 1 if anything is missing
set -euo pipefail

check_only=false
[[ ${1:-} == --check ]] && check_only=true

# package | command that proves it is there
requirements=(
    "go|go"
    "python|python3"
)

missing=()
for requirement in "${requirements[@]}"; do
    package=${requirement%%|*}
    command -v "${requirement#*|}" >/dev/null 2>&1 || missing+=("$package")
done

if ((${#missing[@]} == 0)); then
    exit 0
fi

echo "DSA Tutor needs: ${missing[*]}"
if $check_only; then
    exit 1
fi

if ! command -v rvn >/dev/null 2>&1; then
    echo "rvn is not available; install these packages with your package manager." >&2
    exit 1
fi
sudo rvn install "${missing[@]}"
