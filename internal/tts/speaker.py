"""Persistent Piper text-to-speech worker.

Reads one JSON request per line on stdin ({"text": ...}) and writes one JSON
response per line on stdout ({"wav": "<base64>"} or {"error": ...}). The
voice model is loaded once at startup and a ready line is printed, so tts.go
can keep this process alive and each request only pays for synthesis.
"""
import base64
import io
import json
import sys
import wave


def load_voice(model_path):
    try:
        from piper import PiperVoice  # piper1-gpl, the current piper-tts on PyPI
    except ImportError:
        from piper.voice import PiperVoice  # legacy rhasspy piper-tts
    return PiperVoice.load(model_path)


def synthesize(voice, text):
    buf = io.BytesIO()
    with wave.open(buf, "wb") as wav_file:
        if hasattr(voice, "synthesize_wav"):
            voice.synthesize_wav(text, wav_file)
        else:  # legacy API
            voice.synthesize(text, wav_file)
    return buf.getvalue()


def main():
    voice = load_voice(sys.argv[1])
    print(json.dumps({"ready": True}), flush=True)
    for line in sys.stdin:
        line = line.strip()
        if not line:
            continue
        try:
            text = json.loads(line)["text"]
            reply = {"wav": base64.b64encode(synthesize(voice, text)).decode("ascii")}
        except Exception as exc:  # report the failure but keep serving
            reply = {"error": str(exc)}
        print(json.dumps(reply), flush=True)


if __name__ == "__main__":
    main()
