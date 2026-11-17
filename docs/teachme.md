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
| Warm up | Guess an answer before being taught | Activates prior knowledge; trying first (even wrongly) improves what you learn next |
| Big idea | Read the technique in a sentence or two | A frame to hang the parts on |
| Parts (2–4) | Read one small idea, then answer a check | Small chunks keep cognitive load down; retrieval right after teaching beats rereading |
| Worked example | Reveal one step at a time, predicting each first | Worked examples help novices; predicting the step makes it active |
| Fill in the blank | Write the key line of a working function | Faded practice: the scaffold is removed exactly where the idea lives |
| Explain it back | Explain it in your own words; the model grades it | Self-explanation with targeted feedback on the gap |
| Lock it in | Mistakes people make, practice problems, cards scheduled | Practice problems from the bank; flashcards for spaced review |

Checks give feedback for the choice you picked (the misconception behind
that wrong answer), plus why the right one is right. A fill-in-the-blank
miss shows a hint; after two misses you can reveal the answer.

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
- `POST /api/teachme/feedback` `{topic, prompt, reference, answer, model}`
  returns `{feedback: {verdict: solid|partial|missing, feedback, missing}}`.

A lesson takes about 35 s with gemma4 and 50 s with ornith-1.5:9b on the
development machine.
