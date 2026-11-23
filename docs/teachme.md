# Teach Me

A page where you ask to be taught any data-structures-and-algorithms topic,
not tied to a chapter, and practise it right away. Press **Teach Me** in the
sidebar, type a topic (or pick a suggestion), choose a level, and the
[Ollama](tutor.md) model writes a short lesson for you. Without a chat model
the page explains what's missing.

## How a lesson is built

The lesson pairs direct instruction (explain the idea clearly first) with
active learning (make you use it straight away), the blend that beats either
one alone. Each stage unlocks only after you've done the one before it.

| Stage | What you do | Why |
|-------|-------------|-----|
| Warm up | Guess an answer before being taught; the tutor answers your guess | Activates prior knowledge; trying first (even wrongly) improves what you learn next |
| Big idea | Read the technique in a sentence or two | A frame to hang the parts on |
| Parts (2–4) | Read one small idea, answer a check until you get it right, optionally say it back or ask about it | Small chunks keep cognitive load down; retrieval right after teaching beats rereading |
| Worked example | Reveal one step at a time, predicting each first | Worked examples help novices; predicting the step makes it active |
| Line by line (algorithms) | Go through the reference code a line or two at a time: what it does, why it's written that way, then answer a question about it | You can't move on until your answer shows you understand the line and its reason |
| Fill in the blank | Write the key line of a working function | Faded practice: the scaffold is removed exactly where the idea lives |
| Explain it back | Explain it in your own words until the explanation is solid | Self-explanation with targeted feedback on the gap |
| Lock it in | Mistakes people make, practice problems, cards scheduled | Practice problems from the bank; flashcards for spaced review |

Checks give feedback for the choice you picked (the misconception behind
that wrong answer) and strike it out, but don't give the answer away: you
pick again until you get it right, and only then see why. The wrap-up counts
checks right first time. A fill-in-the-blank miss shows a hint; after two
misses you can reveal the answer.

## Replies to what you type

Anything you type (the warm-up guess, say-it-back after a part, an answer
about a line of code, the explain-it-back) gets a reply from
`POST /api/teachme/respond`. The reply splits what you wrote into its
claims and marks each one ✓ correct, ~ partly or ✗ wrong, with a sentence
that confirms it with the reason or corrects it. Then comes a short reply
that fixes the most important gap and, until your answer is **solid**
(what happens *and* why), a new follow-up question in different words.

The line-by-line steps and explain-it-back are *gated*: the lesson waits
until an answer is solid, and there's no Skip. **I don't get it yet** /
**Explain it differently** asks for a new angle (a trace, an analogy, what
breaks without it) and a fresh question instead. Each attempt tells the
model which try it is, so it doesn't repeat itself. After two tries at
explain-it-back the model answer is shown, but you still have to explain it
in your own words to finish.

A line-by-line step has one way out: after 3 real attempts ("Explain it
differently" doesn't count) **Show me the answer and move on** appears. It
shows the answer and the reason, and adds a flashcard for that line (the
code and the question on the front) to the review deck, due tomorrow, so
the line still comes back. The wrap-up shows how many lines were explained
back and how many were left to review.

## Line by line

When the model calls the topic an `algorithm` (`kind`; a lesson with code
and no `kind` counts as one), its reference code is required and the lesson
gains a line-by-line stage after the worked example. `POST
/api/teachme/walk` `{topic, code, model}` shows the model the numbered code
and asks only for line ranges with what/why/question/answer: models are
poor at quoting their own code, so the server cuts the lines itself
(`fitWalk` sorts, clamps, drops duplicates and gives any skipped line to
the step before, so every line is covered once). The page builds it in the
background once you reach the last part, not at the start: Ollama answers
one request at a time, and a 30 s walkthrough would hold up the warm-up
reply. It's saved with the lesson.

The page shows the whole code with the current lines lit, earlier ones
explained (collapsible) and later ones dimmed.

## Voice

Teach Me uses the same narration as the chapters ([tts.md](tts.md)): the
voice and speaker pickers (phonon, Piper, ...) move into the lesson bar,
and every stage has Listen and Skip. With **Read aloud** on (the default,
`dsatutor_teachme_voice`), each new stage, check feedback and tutor reply
is read as it appears. Code in the text is said as words (speech-text.js);
code blocks, your own answers and anything inside `<details>` aren't read.

## Spaced review

Each lesson's flashcards join one deck in the browser
(`localStorage` key `dsatutor_teachme_deck`). A new card is first due
tomorrow, then at growing gaps (1, 3, 7, 16, 35, 80 days) as you recall it:
**Good** moves it up a step, **Easy** two, **Hard** keeps its step at a
shorter gap, **Again** sends it back to the start and shows it again in the
same session. Due cards from every topic are shuffled together (interleaved
practice). The sidebar button shows how many are due. Keys during review:
Space shows the answer, 1–4 rate it.

Lessons and your progress through them are kept (last 15) under
`dsatutor_teachme_history`, so you can reopen one from the Teach Me page.

## Server

- `POST /api/teachme/lesson` `{topic, level, model}` returns
  `{model, lesson}`. One structured-output call (`tutor.TeachTopic`, schema
  in `teachSchema`) writes the lesson; `TeachLesson.Clean` then shapes it:
  - Checks come from the model as the correct answer's text plus wrong
    answers with their misconceptions, not as choices with an index. Models
    misnumbered the answer index and wrote explanations like "the first
    choice", which break once choices are shuffled. The server shuffles and
    builds `choices`, `answer` and per-choice `feedback`.
  - Fill-in-the-blank: the model writes complete code and names the piece to
    hide (`blank`); the server cuts it out (last occurrence, whitespace
    tolerant), so the answer always matches the gap. When the named piece
    isn't in the code (ornith often invented one), one small follow-up call
    shows the code with line numbers and asks for the key line plus a
    matching prompt and hint (`pickBlankLine`).
  - Practice problems are matched against tags, category and title
    (`teachProblems`): the whole topic counts most, then each two-word phrase
    in it, then the model's `practice_keywords`; easiest first.
- `POST /api/teachme/respond` `{topic, kind: warmup|chunk|line|explain,
  context, prompt, reference, answer, attempt, model}` returns
  `{feedback: {verdict: solid|partial|missing, reply, claims: [{claim,
  verdict: correct|partly|wrong, note}], follow_up, missing}}`
  (`tutor.Respond`). `/api/teachme/feedback` is an alias for it.
- `POST /api/teachme/walk` `{topic, code, model}` returns `{steps: [{from,
  to, code, what, why, question, answer}]}` (`tutor.WalkCode`).

A reply takes about 5 s and a walkthrough about 30 s with gemma4. A lesson takes about 35 s with gemma4 and 50 s with ornith-1.5:9b on the
development machine.
