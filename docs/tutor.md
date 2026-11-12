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

## Pointing at the page and writing code

With a model that supports tool calling, the tutor acts on the screen
instead of only describing it:

| Tool | What you see |
|------|--------------|
| `highlight` | the region scrolls into view and gets outlined, with the tutor's note pinned under it. In your code, the lines themselves light up with the note right below them. |
| `write_code` | your editor changes and the changed lines flash green. The chat shows **Edited your code** with an **Undo** button (one undo step; Ctrl/Cmd+Z works too). |
| `open_tab` | the problem panel switches to Description, Hints or Visualize (never Solution) |
| `select_test` | a test case is selected, ready for Run |

Each action leaves a chip in the conversation; clicking a *Pointed at…*
chip shows that highlight again. The tutor edits your code only when you ask
it to write, fix or scaffold code, or accept its offer.

What it can point at depends on the view:

- **Problem**: your code (by line), the statement, each example, the
  constraints, the test panel, your last run and submission, revealed hints,
  the Run and Submit buttons
- **Chapter**: each lesson section, each revealed worked-example step, each
  check question, and code blocks
- **Sandbox**: your code (by line) and its output

The page sends this list (ids and labels) with every message. The server
offers tools that only accept those ids and runs the tool loop itself: each
call is streamed to the page as an `action` event, and the result goes back
to the model. Edits are applied to the server's copy of the code too, so the
model's next step sees the code as it now is. A reply can go through at most
six tool rounds before the model has to answer in words.

**Which model**: tool use depends on the model. In testing, `ornith`
reliably highlighted lines and wrote fixes, while `gemma4` mostly answered
in text without using the tools. Pick the model in the panel (it's
remembered), or set `DSATUTOR_TUTOR_MODEL=ornith`.

## Teacher Mode and the Teacher's Scratchpad

On a chapter page, press **Teacher mode**. Instead of waiting for
questions, the tutor teaches the chapter one idea at a time. Each turn it:

1. points at the section it's on (`highlight`),
2. makes the idea visible in the **Teacher's Scratchpad**, either as a
   step-by-step diagram or as a small program it runs and traces, and
3. asks you one question, often with clickable answers (`ask_learner`), and
   waits.

It adapts: a right answer moves on; a wrong one gets the misconception shown
in the scratchpad and the question asked another way. The quick prompts
**Next**, **Show me**, **I'm confused**, **Quiz me** and **Let me try** steer
it. Each chapter keeps its own lesson, so you can leave and resume. The
teacher works from the whole chapter: concepts with why they work, worked
examples with their traces, and the check questions.

The **scratchpad** is a floating window that you can drag, resize and minimize:

- **Programs** (`scratchpad_run_code`) run on the server through the same
  tracer as the Sandbox. You step through them line by line, with the
  current line highlighted, the variables, the data structures (lists,
  dicts, sets, linked lists, trees, grids) and the output.
- **Diagrams** (`scratchpad_visualize`) are frames the tutor describes
  (array, bar chart, grid, linked list, tree, hash map, set, graph or a
  single value, with pointers and windows highlighted). They use the same
  renderer as the visualizer.
- You can edit the code and press **Run** yourself; **Code** opens the
  editor over a diagram. The tutor sees the scratchpad's code and output
  with every message, so "Let me try" exercises get reviewed.
- Everything shown stays in the scratchpad's history (‹ 2 of 3 ›), and each
  chip in the chat reopens its diagram or reruns its program.

The scratchpad tools are available to the regular tutor too (on problems,
chapters and the sandbox); only the interviewer goes without them.

## Mock interviews

**Mock interview** in the problem list picks a Google-reported problem at
the level you choose (L3, L4 or L5; optionally only unsolved ones), and
**Interview mode** on a problem runs one on that problem. Interviews are 30,
45 or 60 minutes.

During the interview the page becomes a shared document. Run, Submit, hints,
the solution, the visualizer, related chapters and the LeetCode number are
hidden, and your code starts from the function signature (your draft comes
back afterwards). The tutor panel is the interviewer. It presents the
problem, answers clarifying questions, asks for your approach and its
complexity before you code, stays quiet while you code, asks you to test by
hand, and asks a follow-up if there's time. It knows the clock and nudges
you with about 10 minutes left. Use the microphone to think out loud.
Leaving the problem mid-interview asks first.

**End interview** (or running out of time) runs your final code against
every test, hidden ones included, and asks the model for a scorecard:

- a verdict from Strong No Hire to Strong Hire,
- a 1–4 rating with cited evidence for each of Google's coding criteria
  (algorithms, coding, communication, problem solving),
- your stated versus actual complexity, hints used, and whether you reached
  the optimal solution,
- strengths, what to do differently, the failed tests' notes, your code and
  the transcript.

The scorecard uses Ollama's structured outputs (a JSON schema), so any chat
model returns a well-formed card. The server clamps out-of-range values. In
testing, a clean interview scored Strong Hire and a weak one (no clarifying
questions, a hint, a bug) No Hire, with the right tests blamed. Past
interviews are kept in the browser (the last 25) and listed in the interview
dialogs. Treat the verdict as a practice signal, not a hiring prediction.

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
| `/api/tutor/chat` | POST | `{"model", "think", "messages": [{"role", "content"}], "context": {"view", "problemId", "chapterNumber", "language", "code", "lastRun", "hintsRevealed", "solutionUnlocked", "voice", "targets": [{"id", "label", "lines", "writable"}], "testCount", "interview": {"elapsedSeconds", "durationSeconds", "level"}, "teach", "scratchpad": {"code", "output"}}}` → NDJSON stream of `{"type":"thinking"\|"delta","text"}` and `{"type":"action","name":"highlight"\|"write_code"\|"open_tab"\|"select_test"\|"scratchpad_code"\|"scratchpad_visualize"\|"ask_learner", ...}`, then `{"type":"done","model"}` or `{"type":"error","error"}` |
| `/api/tutor/interview/score` | POST | `{"model", "problemId", "language", "code", "messages", "elapsedSeconds", "durationSeconds", "level"}` → `{"scorecard", "tests": {"status", "passed", "total", "failures"}, "model"}` (`error` instead of `scorecard` if the write-up failed) |
| `/api/tutor/transcribe` | POST | body: a WAV recording, `?model=` names the chat model in use → `{"text"}` |

The server builds the system prompt itself (`internal/tutor/prompt.go`).
System messages from the browser are dropped, and history is capped at the
last 24 turns.
