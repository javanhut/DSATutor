package tutor

import (
	"fmt"
	"strings"

	"dsatutor/internal/chapter"
	"dsatutor/internal/practice"
)

// Limits on what the browser can put in a prompt.
const (
	maxCodeChars    = 20000
	maxRunChars     = 4000
	maxMessageChars = 8000
	maxHistory      = 24
)

// Context is what the learner has on screen, sent with every chat request.
// The server looks the problem and chapter up itself, so the prompt always
// carries the real statement, hints and reference notes.
type Context struct {
	View             string   `json:"view"` // "problem", "chapter", "sandbox" or "home"
	ProblemID        string   `json:"problemId,omitempty"`
	ChapterNumber    int      `json:"chapterNumber,omitempty"`
	Language         string   `json:"language,omitempty"` // python or oxigen
	Code             string   `json:"code,omitempty"`
	LastRun          string   `json:"lastRun,omitempty"` // the last run/submit result as shown
	HintsRevealed    int      `json:"hintsRevealed,omitempty"`
	SolutionUnlocked bool     `json:"solutionUnlocked,omitempty"`
	Voice            bool     `json:"voice,omitempty"`     // replies are read aloud
	Targets          []Target `json:"targets,omitempty"`   // regions the tutor can point at
	TestCount        int      `json:"testCount,omitempty"` // visible test cases
	// Interview is set during a mock interview: the tutor plays the
	// interviewer instead of a coach.
	Interview *Interview `json:"interview,omitempty"`
	// Teach turns the chapter chat into Teacher Mode: the tutor leads the
	// lesson instead of answering questions.
	Teach      bool            `json:"teach,omitempty"`
	Lesson     *Lesson         `json:"lesson,omitempty"` // Teacher Mode's own lesson plan
	Scratchpad *Scratchpad     `json:"scratchpad,omitempty"`
	Learner    *LearnerProfile `json:"learner,omitempty"`
}

// LearnerProfile is durable evidence from retrieval, completion, practice,
// hints and spaced review. It lets the tutor adapt support to demonstrated
// knowledge instead of inferring ability from the latest chat message.
type LearnerProfile struct {
	Level              string   `json:"level,omitempty"`
	Mastery            int      `json:"mastery,omitempty"`
	ChecksCorrect      int      `json:"checksCorrect,omitempty"`
	ChecksTotal        int      `json:"checksTotal,omitempty"`
	CompletionsCorrect int      `json:"completionsCorrect,omitempty"`
	CompletionsTotal   int      `json:"completionsTotal,omitempty"`
	PracticeSolved     int      `json:"practiceSolved,omitempty"`
	PracticeAttempted  int      `json:"practiceAttempted,omitempty"`
	HintsUsed          int      `json:"hintsUsed,omitempty"`
	ReviewReps         int      `json:"reviewReps,omitempty"`
	ReviewsDue         int      `json:"reviewsDue,omitempty"`
	Misconceptions     []string `json:"misconceptions,omitempty"`
	NextStep           string   `json:"nextStep,omitempty"`
	ProblemAttempts    int      `json:"problemAttempts,omitempty"`
	ProblemSolved      bool     `json:"problemSolved,omitempty"`
	ProblemHintsUsed   int      `json:"problemHintsUsed,omitempty"`
	SolutionViewed     bool     `json:"solutionViewed,omitempty"`
}

const persona = `You are the DSA Tutor's built-in coach: a patient, sharp teacher of data structures and algorithms helping one learner prepare for coding interviews.

How you teach:
- Coach, don't solve. Ask a guiding question or give the smallest hint that unblocks the learner. Escalate (nudge -> key insight -> plan -> code) only when they are still stuck or ask for more.
- Don't volunteer the full solution. But the learner is in charge: when they explicitly ask you to fix, write or finish their code, do it, then explain what changed and why. Short snippets that illustrate an idea or fix one line are always fine.
- When reviewing their code, point to specific line numbers, explain WHY the logic fails, and give a concrete input that breaks it. Trace small examples by hand.
- Bring up time and space complexity: ask them to state it, then confirm or correct it with reasoning. Compare brute force with the better approach.
- Name the pattern (two pointers, sliding window, BFS, DP...) once they are close, and connect to the chapter that teaches it.
- Be honest and precise. If their code is correct, say so and suggest what to check next (edge cases, complexity, cleaner code).
- Keep replies focused: a few short paragraphs or a short list. Use Markdown, with fenced code blocks for code. Never use LaTeX or $...$ math; write math as plain text, like 6 - 3 = 3 or O(n log n).`

const toolNote = `

You can act on the learner's screen with tools:
- highlight: whenever you refer to something on screen, point at it. When you talk about a bug, highlight the exact lines in their code with a one-sentence note. When you mention an example, a constraint or a lesson section, highlight it. One or two highlights per reply is plenty.
- write_code: edit their editor only when they ask you to write, fix, or scaffold code, or say yes when you offer. Make the smallest edit that does the job. When coaching, a skeleton with TODO comments beats a finished answer.
- open_tab and select_test: show them a tab or a test case when it helps.
- scratchpad_sketch: a hand-drawn whiteboard in the scratchpad for intuition: notes, arrows, circled cells, built up step by step.
- scratchpad_visualize and scratchpad_run_code: the teacher's scratchpad pops up beside the page. Draw a diagram or run a tiny program there whenever a picture or a concrete trace would make an idea click, rather than describing it in words.
- ask_learner: check understanding with clickable answers, then wait.
Call the tools; never describe a tool call in text. After using them, still answer in words.`

const voiceNote = `
The learner is talking to you by voice and your reply is read aloud: answer in two to five short spoken sentences, avoid code blocks and symbols-heavy notation, and say identifiers as plain words.`

// SystemPrompt builds the system message for one request. ws is nil when
// the model can't use tools.
func SystemPrompt(ctx Context, problem *practice.Problem, ch *chapter.Chapter, ws *Workspace) string {
	var b strings.Builder
	b.WriteString(persona)
	if ws != nil && len(ws.Tools()) > 0 {
		b.WriteString(toolNote)
	}
	if ctx.Voice {
		b.WriteString(voiceNote)
	}
	b.WriteString("\n\n# What the learner is looking at\n")

	switch {
	case problem != nil:
		writeProblem(&b, ctx, problem)
	case ch != nil:
		writeChapter(&b, ch)
	case ctx.View == "sandbox":
		b.WriteString("The code sandbox, where they run their own code and watch it execute step by step.\n")
	default:
		b.WriteString("The course home page. Help them choose what to study or answer general DSA questions.\n")
	}

	if code := strings.TrimSpace(ctx.Code); code != "" {
		lang := ctx.Language
		if lang == "" {
			lang = "python"
		}
		fmt.Fprintf(&b, "\n## The learner's current code (%s, with line numbers)\n```\n%s\n```\n", lang, numberLines(truncate(code, maxCodeChars)))
		b.WriteString("The `N | ` prefixes are line numbers for you to cite (\"line 5\"); they are not part of the code, so leave them out when you quote it.\n")
		if lang == "oxigen" {
			b.WriteString("Oxigen is a small Python-like teaching language; treat its syntax charitably.\n")
		}
	}
	if run := strings.TrimSpace(ctx.LastRun); run != "" {
		fmt.Fprintf(&b, "\n## Their latest run or submission\n%s\n", truncate(run, maxRunChars))
	}
	writeScratchpad(&b, ctx.Scratchpad)
	writeLearnerProfile(&b, ctx.Learner)
	if ws != nil {
		b.WriteString(ws.Describe())
	}
	return b.String()
}

func writeProblem(b *strings.Builder, ctx Context, p *practice.Problem) {
	fmt.Fprintf(b, "Practice problem: **%s** (%s, category: %s)\n\n%s\n", p.Title, p.Difficulty, p.Category, strings.TrimSpace(p.Description))
	if len(p.Constraints) > 0 {
		b.WriteString("\nConstraints:\n")
		for _, c := range p.Constraints {
			fmt.Fprintf(b, "- %s\n", c)
		}
	}
	for i, ex := range p.Examples {
		fmt.Fprintf(b, "\nExample %d: input %s -> output %s", i+1, ex.Input, ex.Output)
		if ex.Explanation != "" {
			fmt.Fprintf(b, " (%s)", ex.Explanation)
		}
	}
	b.WriteString("\n")

	if ctx.HintsRevealed > 0 {
		b.WriteString("\nHints the learner has already revealed:\n")
		for i, h := range p.Hints {
			if i >= ctx.HintsRevealed {
				break
			}
			fmt.Fprintf(b, "- (%s) %s\n", h.Type, h.Content)
		}
	}

	// Reference notes make the tutor accurate; the persona keeps it from
	// reading them out. The code itself is only shared once the learner has
	// opened the solution, so a small model can't leak it early.
	s := p.Solution
	b.WriteString("\n## Reference notes (for you only; do not reveal them wholesale)\n")
	if s.BruteForce != "" {
		fmt.Fprintf(b, "- Brute force: %s\n", s.BruteForce)
	}
	if s.KeyInsight != "" {
		fmt.Fprintf(b, "- Key insight: %s\n", s.KeyInsight)
	}
	if s.WhyItWorks != "" {
		fmt.Fprintf(b, "- Why it works: %s\n", s.WhyItWorks)
	}
	tc, sc := firstNonEmpty(s.TimeComplexity, p.TimeComplexity), firstNonEmpty(s.SpaceComplexity, p.SpaceComplexity)
	if tc != "" || sc != "" {
		fmt.Fprintf(b, "- Optimal complexity: time %s, space %s\n", tc, sc)
	}
	for _, m := range s.CommonMistakes {
		fmt.Fprintf(b, "- Common mistake: %s\n", m)
	}
	if ctx.SolutionUnlocked && s.Code != "" {
		fmt.Fprintf(b, "\nThe learner has unlocked the model solution, so you may discuss it freely:\n```python\n%s\n```\n", strings.TrimSpace(s.Code))
	}
}

func writeChapter(b *strings.Builder, ch *chapter.Chapter) {
	fmt.Fprintf(b, "Chapter %d: **%s**. %s\n", ch.Number, ch.Title, ch.Summary)
	if ch.Motivation != "" {
		fmt.Fprintf(b, "\nWhy it matters: %s\n", ch.Motivation)
	}
	if len(ch.Objectives) > 0 {
		b.WriteString("\nObjectives:\n")
		for _, o := range ch.Objectives {
			fmt.Fprintf(b, "- %s\n", o)
		}
	}
	if len(ch.WhenToUse) > 0 {
		b.WriteString("\nWhen to use this pattern:\n")
		for _, w := range ch.WhenToUse {
			fmt.Fprintf(b, "- %s\n", w)
		}
	}
	for _, c := range ch.Concepts {
		fmt.Fprintf(b, "\nConcept: %s. %s", c.Name, c.Description)
		if c.Intuition != "" {
			fmt.Fprintf(b, " Intuition: %s", c.Intuition)
		}
		if len(c.CommonMistakes) > 0 {
			fmt.Fprintf(b, " Common mistakes: %s", strings.Join(c.CommonMistakes, "; "))
		}
		b.WriteString("\n")
	}
	b.WriteString("\nHelp them understand this chapter: check their understanding with questions, give fresh small examples, and relate ideas to interview problems.\n")
}

// Conversation assembles the messages sent to the model: the system prompt
// and the most recent turns, each capped in size.
func Conversation(system string, history []Message) []Message {
	var turns []Message
	for _, m := range history {
		if m.Role != "user" && m.Role != "assistant" {
			continue // the browser never gets to set the system prompt
		}
		if c := strings.TrimSpace(m.Content); c != "" {
			turns = append(turns, Message{Role: m.Role, Content: truncate(c, maxMessageChars)})
		}
	}
	if len(turns) > maxHistory {
		turns = turns[len(turns)-maxHistory:]
	}
	return append([]Message{{Role: "system", Content: system}}, turns...)
}

func numberLines(code string) string {
	lines := strings.Split(code, "\n")
	for i, l := range lines {
		lines[i] = fmt.Sprintf("%3d | %s", i+1, l)
	}
	return strings.Join(lines, "\n")
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "\n... (truncated)"
}

func firstNonEmpty(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func writeScratchpad(b *strings.Builder, sp *Scratchpad) {
	if sp == nil {
		return
	}
	title, caption, code := strings.TrimSpace(sp.Title), strings.TrimSpace(sp.Caption), strings.TrimSpace(sp.Code)
	if title == "" && caption == "" && code == "" {
		return
	}
	b.WriteString("\n## What the learner currently sees in the scratchpad\n")
	if title != "" {
		fmt.Fprintf(b, "%s", truncate(title, 200))
		if sp.Steps > 0 {
			fmt.Fprintf(b, " (step %d of %d)", min(max(sp.Step, 1), sp.Steps), sp.Steps)
		}
		b.WriteString("\n")
	}
	if caption != "" {
		fmt.Fprintf(b, "Current caption: %s\n", truncate(caption, 500))
	}
	if code != "" {
		fmt.Fprintf(b, "The learner may have edited this code:\n```python\n%s\n```\n", numberLines(truncate(code, 6000)))
	}
	if out := strings.TrimSpace(sp.Output); out != "" {
		fmt.Fprintf(b, "Its last output:\n%s\n", truncate(out, 1500))
	}
}

func writeLearnerProfile(b *strings.Builder, profile *LearnerProfile) {
	if profile == nil {
		return
	}
	fmt.Fprintf(b, "\n## Learner model (use this to choose support, not to praise a score)\nLevel: %s; mastery evidence: %d%%\n", profile.Level, profile.Mastery)
	fmt.Fprintf(b, "Retrieval: %d/%d; completions: %d/%d; practice: %d solved from %d attempted; hints used: %d; spaced reviews: %d (%d due).\n",
		profile.ChecksCorrect, profile.ChecksTotal, profile.CompletionsCorrect, profile.CompletionsTotal,
		profile.PracticeSolved, profile.PracticeAttempted, profile.HintsUsed, profile.ReviewReps, profile.ReviewsDue)
	if profile.ProblemAttempts > 0 || profile.ProblemSolved {
		fmt.Fprintf(b, "Current problem: %d attempts; %d hints used; solved=%t; solution viewed=%t.\n", profile.ProblemAttempts, profile.ProblemHintsUsed, profile.ProblemSolved, profile.SolutionViewed)
	}
	if len(profile.Misconceptions) > 0 {
		b.WriteString("Recent misconceptions or support needs:\n")
		for _, misconception := range profile.Misconceptions {
			fmt.Fprintf(b, "- %s\n", truncate(misconception, 300))
		}
	}
	if profile.NextStep != "" {
		fmt.Fprintf(b, "Best next evidence: %s\n", truncate(profile.NextStep, 400))
	}
}

const teacherPersona = `You are the DSA Tutor's teacher, running a live, one-on-one lesson with one learner. You lead; the learner follows and answers. You write this lesson yourself, from your own knowledge, following your lesson plan below. The chapter's prewritten page is not part of this lesson: the learner sees a blank lesson page that you fill in section by section with write_lesson, as their reference. The teacher's scratchpad beside it is only for examples: sketches, diagrams and code you trace.

How you teach:
- One section of your plan at a time, in order. Starting a section, first write it onto the lesson page with write_lesson: your own short explanation of the idea with one tiny concrete example. Then teach it in a few short turns before moving on.
- Every turn is one coordinated teaching beat: point (highlight) at the words on the lesson page you are teaching, show the example on the scratchpad when a visual helps, and end with one question. Keep all of it on the same single idea. Then stop and wait.
- Only point at what you have written on the lesson page (the lesson-N targets); quote its words verbatim.
- Do not introduce a second complexity claim in the checkpoint; save each new claim for its own turn.
- Keep your chat words short: two to five sentences per turn. The lesson page, the scratchpad and the question do the heavy lifting; don't repeat the lesson text in chat.
- The scratchpad must show exactly the example your words and the lesson page use: the same array, the same numbers, the same pointer positions.
- Prefer scratchpad_sketch for movement and structure (array shifts, pointer motion, windows, nodes). Use scratchpad_run_code when execution itself is the lesson or the learner needs to edit/run an exercise; do not replace a simple visual trace with a large code listing.
- Never use LaTeX or $...$ math; write math as plain text, like 1 + 11 = 12 or O(n).
- Adapt support from the learner model. Starting: model one small worked example explicitly. Building: use a completion step and ask for a prediction. Practicing: let the learner explain or perform the next step before showing anything. Secure: skip routine explanation and use transfer, contrast, edge cases, or mixed-pattern questions. If a misconception is listed, diagnose it directly before advancing.
- Use at most one scratchpad tool per turn (write_lesson doesn't count). Do not redraw an idea the learner has already demonstrated unless they ask or their answer exposes confusion. Never write "here's the code:" with nothing after it.
- Ask exactly one question with ask_learner whenever tools are available. Make it a prediction about the example now on the scratchpad or lesson page, not a generic recall question. Never repeat that question in your text and never write stage directions.
- Use small concrete inputs (5 to 8 elements) and the same input across a section so the learner can follow it.
- Adapt: if they answer right, confirm briefly and, once the section's goal is met, start the next section by writing it; if they're wrong, find the misconception, show it on the scratchpad, and ask again a different way. If they say they're confused, slow down and use a smaller example. If they say "next", start the next section.
- Near the end, give a short exercise on the scratchpad (code with a blank to fill or a function to finish) and review their attempt.
- Never dump the whole lesson at once and never ask more than one question per turn.`

// TeacherPrompt builds the system message for Teacher Mode on a chapter.
// The chapter contributes only its topic and scope; the lesson itself is
// the model's own, following the plan. lastUser is the learner's latest
// message (it may ask to move on to the next section).
func TeacherPrompt(ctx Context, ch *chapter.Chapter, ws *Workspace, lastUser string) string {
	var b strings.Builder
	b.WriteString(teacherPersona)
	if ctx.Voice {
		b.WriteString(voiceNote)
	}
	b.WriteString("\n\n# The topic you are teaching\n")
	writeLessonTopic(&b, ch)
	writeLesson(&b, ctx.Lesson)
	if due := ctx.Lesson.SectionDue(lastUser); due >= 0 {
		fmt.Fprintf(&b, "\nThis turn: start section %d (%s). Write it onto the lesson page with write_lesson first.\n", due, ctx.Lesson.Sections[due].Heading)
	}
	writeScratchpad(&b, ctx.Scratchpad)
	writeLearnerProfile(&b, ctx.Learner)
	if ws != nil {
		b.WriteString(ws.Describe())
	}
	return truncate(b.String(), 60000)
}
