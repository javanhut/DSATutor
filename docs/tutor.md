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

A reply runs at most six rounds of tool calls before the model must answer in words; the server stops it there even if the model keeps calling tools, and it accepts list arguments that small models send as JSON strings. Each action leaves a chip in the conversation; clicking a *Pointed at…*
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
questions, the tutor teaches the chapter's topic one idea at a time, from a
lesson it plans and writes itself. The chapter's prewritten page is put away
while Teacher Mode is on (turn it off to get it back); the normal chapter
view and the regular tutor still use the prewritten pages.

**A lesson written live.** When a lesson starts (or restarts with Clear), the
server asks the model for a fresh plan of 4 to 8 sections (`PlanLesson` in
`internal/tutor/lesson.go`, streamed to the page as `{"type":"lesson"}`). The
chapter contributes only the topic and its scope (title, summary,
objectives, concept names), never its prewritten explanations, worked
examples or check questions. The lesson page starts blank, with the plan as
an outline, and the teacher writes each section onto it when it starts that
section (`write_lesson`). A section is written in parts so the concept is
always stated, not just shown: `concept` (the idea in general terms: what it
is, how it works, what it costs), `why` (the reason or invariant), `example`
(one tiny illustration), `key_points`, and optional `code`; the server
composes them into the section's text and sends back a concept too thin to
teach from. Questions check the concept (why, what changes, cost, when to
use), not just details of the example. That text is the learner's reference for the
lesson. The plan and written sections are kept per chapter in localStorage
(`dsatutor_teacher_lesson:<n>`) and sent back with each turn, so the server
knows which section is current and which are written. Each turn the teacher:

1. writes the section on the lesson page when it starts one, and points at
   the words it's teaching there (`highlight`, targets `lesson-N`),
2. shows the example in the **Teacher's Scratchpad**, which holds only
   examples: a hand-drawn whiteboard sketch, an exact step-by-step diagram,
   or a small program it runs and traces, and
3. lets the example unfold automatically, then asks one prediction question
   directly on the scratchpad (`ask_learner`) and waits there for your answer.

**Precise pointing.** `highlight` takes a `quote` (the exact words on the
page) or a line range (the editor, a lesson code block, an exercise, or the
animation's code panel). The page finds the words — ignoring case,
whitespace, punctuation and Markdown, searching the whole lesson if the
model named the wrong region, and using the longest matching start if it
paraphrased the end — highlights exactly that text in place, scrolls it into
view, pins the note under it, and a **hand** (reading through this line) or
an **arrow** (look here) glides into the margin beside it and follows it
while the page scrolls. The scratchpad moves out of the way of whatever is
being pointed at, unless you dragged it somewhere yourself.

**Drawings must have content.** Sketch elements need their text (boxes,
nodes and notes without a label are dropped) and cell rows need values; a
sketch with nothing visible is rejected with the reason so the model redraws,
and a model may draw at most twice per turn. The director's own sketch uses a
schema where every row, note, arrow and box has its content as a required
field, and it reuses the turn's exact conversation so Ollama's prompt cache
skips re-reading the prompt (about 25 s per turn instead of 90 s with
ornith-1.5).

**Keeping the teacher on track.** Each request carries a record of what
earlier replies showed: sections written, drawings, and the question asked
with its choices, prefixed to the learner's next message as
`[Your last reply also showed: …]`. Without it the model can't tell what an
answer answers. The prompt also carries the text already on the lesson page,
and a written section can't be rewritten. In Teacher Mode each round's text
is held until the round ends so narration of the tools ("Let me write…",
"I'm waiting for your answer") is dropped (`tutor.StripStageDirections`),
and the model's own question is held to the end of the turn. When the
teacher stays on a section, one small structured call (`GoalMet`) asks
whether the learner has already shown its goal; if so the held question is
dropped and the next section is written and taught instead of quizzing the
same idea again. **Lesson order.** A Teacher Mode turn is held on the server and sent in
lesson order (`teachTurn` in `tutor_handlers.go`): the section written on the
lesson page, the teacher's words, any pointer, one drawing, then the
question. A model sketch too thin to teach from is replaced by the
director's rather than shown beside it, so the scratchpad gets one drawing
per turn. With speech on, the page reads a newly written section aloud first
(marked while it's read), and the drawing and question appear once that
reading is done (`speech.until` in `tutor.js`); stopping the voice shows them
at once. The teacher then talks through the example: drawings carry an
`explanation` (said as the drawing appears) and per-step captions (sketch
`captions`, diagram frame captions), and the scratchpad steps along with the
narration instead of playing on a timer (`manual`, `dsaScratchpad.step`).
Finally the question appears and is read aloud with its choices. When the server writes a later section, the same call
also produces a one-sentence transition from the previous section, appended
to the reply (and spoken) before the section appears.

**Every turn is staged.** Whether a model writes the section, draws, or
creates a checkpoint varies by model and run, so after each Teacher Mode turn
the server checks: if a section was due and not written (the current one
until it has text, or the next one when you asked to move on), or it drew
nothing substantial (a sketch
with fewer than three elements or no cells, nodes, boxes, circles or
arrows), or omitted the question, a "director" asks the same model for the missing parts with one
structured-output call (JSON schema, which every model follows), retrying
once if its sketch is still thin. The chat shows "Writing up the lesson…"
meanwhile.

**Voice.** Teacher Mode reads the teacher's replies aloud by default (the
regular tutor doesn't; **Speak replies** is remembered separately for each).
Every reply has a 🔊 button to hear it again. Browsers block sound until
you've clicked on the page, so when a lesson starts on its own (after a
reload) the first reply shows **🔊 Play voice**; one click plays it, and
later replies play by themselves.

**Starting the lesson.** A chapter's lesson starts by itself whenever the
tutor panel shows that chapter in Teacher Mode with nothing said yet: when
you turn Teacher Mode on, open the panel, press Clear, reload the page, or
reach the chapter any way at all (the chapter title is watched, not just
sidebar clicks). If the start fails (Ollama not running, a model still
loading), the panel says so and offers **Start the lesson** instead of
retrying in a loop; a failed start leaves nothing behind in the
conversation.

It adapts: a right answer moves on; a wrong one gets the misconception shown
in the scratchpad and the question asked another way. The quick prompts
**Next**, **Show me**, **I'm confused**, **Quiz me** and **Let me try** steer
it. Each chapter keeps its own lesson, so you can leave and resume.

**Learner model.** Every tutor request carries durable learning evidence for
the current chapter: retrieval questions, completion exercises, practice
attempts and solves, hint use, completed spaced reviews, and recent
misconceptions. Teacher Mode fades support from a worked example for a
starting learner, to completion and prediction while building, to learner-led
steps and transfer questions while practising. A chapter can be completed
only after its recall and completion work plus at least one transfer problem;
the stronger **secure** level for categorized pattern chapters additionally
requires a successful delayed review. When a learner later demonstrates a previously missed idea, that
misconception is removed from the model.

The **scratchpad** is a floating window that you can drag, resize and minimize:

- **Programs** (`scratchpad_run_code`) run on the server through the same
  tracer as the Sandbox. You step through them line by line, with the
  current line highlighted, the variables, the data structures (lists,
  dicts, sets, linked lists, trees, grids) and the output.
- **Diagrams** (`scratchpad_visualize`) are frames the tutor describes
  (array, bar chart, grid, linked list, tree, hash map, set, graph or a
  single value, with pointers and windows highlighted). They use the same
  renderer as the visualizer.
- Short diagrams and sketches play automatically. A **Your turn** checkpoint
  stays disabled until the final step, so the question and the visual form
  one teaching beat instead of competing for attention. Answering there
  continues the tutor conversation immediately.
- **Whiteboard sketches** (`scratchpad_sketch`) are hand-drawn: rough.js
  draws boxes, circles, arrows and rows of array cells in a pen style, with
  text in the Caveat handwriting font on grid paper. Each step's strokes
  draw themselves in order like a pen, in about 2.5 seconds. Pointers are
  given by index (`"pointers": {"left": 0, "right": 4}`) and drawn under
  the right cells, and the board spaces out anything the model placed
  overlapping. Without the rough.js CDN it falls back to its own wobbly
  lines.
- The **✎** button gives every other diagram and traced run a handwritten
  look too (handwriting font, pen-wobble lines, grid paper). It's
  remembered per browser.
- You can edit the code and press **Run** yourself; **Code** opens the
  editor over a diagram. With every message the tutor sees the current board
  title, step and caption, plus editable code and output, so it can continue
  from exactly what you are looking at.
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
| `/api/tutor/chat` | POST | `{"model", "think", "messages": [{"role", "content"}], "context": {"view", "problemId", "chapterNumber", "language", "code", "lastRun", "hintsRevealed", "solutionUnlocked", "voice", "targets": [{"id", "label", "lines", "writable"}], "testCount", "interview": {"elapsedSeconds", "durationSeconds", "level"}, "teach", "lesson": {"title", "sections": [{"heading", "goal"}], "current", "written"}, "scratchpad": {"code", "output"}}}` → NDJSON stream of `{"type":"thinking"\|"delta","text"}` and `{"type":"action","name":"highlight"\|"write_code"\|"open_tab"\|"select_test"\|"scratchpad_code"\|"scratchpad_visualize"\|"scratchpad_sketch"\|"write_lesson"\|"ask_learner", ...}`, `{"type":"lesson","lesson"}` (Teacher Mode's new plan), `{"type":"status","text"}`, then `{"type":"done","model"}` or `{"type":"error","error"}` |
| `/api/tutor/interview/score` | POST | `{"model", "problemId", "language", "code", "messages", "elapsedSeconds", "durationSeconds", "level"}` → `{"scorecard", "tests": {"status", "passed", "total", "failures"}, "model"}` (`error` instead of `scorecard` if the write-up failed) |
| `/api/tutor/transcribe` | POST | body: a WAV recording, `?model=` names the chat model in use → `{"text"}` |

The server builds the system prompt itself (`internal/tutor/prompt.go`).
System messages from the browser are dropped, and history is capped at the
last 24 turns.
