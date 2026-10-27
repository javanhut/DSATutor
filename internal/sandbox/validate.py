"""Runs a problem's trusted check() on an answer that came from Oxigen code.

Python answers are validated inside harness.py; Oxigen answers arrive as
JSON, so this reads {validator, raw, result, expected} on stdin, passes the
raw input as args, and prints {passed, message}.
"""

import json
import sys


def main():
    data = json.load(sys.stdin)
    scope = {}
    exec(compile(data["validator"], "<validator>", "exec"), scope)
    check = scope.get("check")
    if check is None:
        verdict = {"passed": False, "message": "validator has no check() function"}
    else:
        try:
            out = check(data["raw"], data["raw"], data["result"], data.get("expected"))
        except Exception as e:  # a crash in the checker means the answer has the wrong shape
            out = (False, f"answer has the wrong shape: {type(e).__name__}: {e}")
        if isinstance(out, tuple):
            verdict = {"passed": bool(out[0]), "message": str(out[1]) if len(out) > 1 else ""}
        else:
            verdict = {"passed": bool(out), "message": ""}
    json.dump(verdict, sys.stdout)


main()
