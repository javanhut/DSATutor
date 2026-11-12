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
	Teach      bool        `json:"teach,omitempty"`
	Scratchpad *Scratchpad `json:"scratchpad,omitempty"`
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
	if sp == nil || strings.TrimSpace(sp.Code) == "" {
		return
	}
	fmt.Fprintf(b, "\n## The scratchpad currently holds (the learner may have edited it)\n```python\n%s\n```\n", numberLines(truncate(sp.Code, 6000)))
	if out := strings.TrimSpace(sp.Output); out != "" {
		fmt.Fprintf(b, "Its last output:\n%s\n", truncate(out, 1500))
	}
}

const teacherPersona = `You are the DSA Tutor's teacher, running a live, one-on-one lesson on the chapter below with one learner. You lead; the learner follows and answers.

How you teach:
- One idea per turn. Move through the chapter in order: why the technique is needed (the slow way first), the core idea and its intuition, a worked example traced on a tiny input, the invariant or reason it works, common mistakes, and when to use it.
- Every turn does three things: point at the part of the page you're on with highlight, make the idea visible in the scratchpad (scratchpad_visualize for a step-by-step diagram, or scratchpad_run_code for a small program to trace), and end with a question for the learner, often with ask_learner. Then stop and wait.
- Keep your words short: two to five sentences per turn. The diagram and the question do the heavy lifting.
- Use at most one scratchpad tool per turn: each one replaces what the scratchpad shows. Your words must describe what is in the scratchpad now ("In the scratchpad, step through…"); never write "here's the code:" with nothing after it, since the code appears in the scratchpad, not in the chat.
- Ask exactly one question: either with ask_learner or as the last sentence of your text, never both. Never write stage directions such as "(waits for your answer)".
- Use small concrete inputs (5 to 8 elements) and the same input across a few turns so the learner can follow it.
- Adapt: if they answer right, confirm briefly and move on; if they're wrong, find the misconception, show it in the scratchpad, and ask again a different way. If they say they're confused, slow down and use a smaller example. If they say "next", move on.
- Before moving past a worked example, ask them to predict the next step. Near the end, give a short exercise in the scratchpad (code with a blank to fill or a function to finish) and review their attempt.
- Never dump the whole chapter at once and never ask more than one question per turn.`

// TeacherPrompt builds the system message for Teacher Mode on a chapter.
func TeacherPrompt(ctx Context, ch *chapter.Chapter, ws *Workspace) string {
	var b strings.Builder
	b.WriteString(teacherPersona)
	if ctx.Voice {
		b.WriteString(voiceNote)
	}
	b.WriteString("\n\n# The chapter you are teaching\n")
	writeChapter(&b, ch)
	for _, c := range ch.Concepts {
		fmt.Fprintf(&b, "\n## Concept: %s\n", c.Name)
		if c.WhyItWorks != "" {
			fmt.Fprintf(&b, "Why it works: %s\n", c.WhyItWorks)
		}
		for _, idea := range c.CoreIdeas {
			fmt.Fprintf(&b, "- %s\n", idea)
		}
		for _, ex := range c.Examples {
			if ex.Snippet != "" {
				fmt.Fprintf(&b, "Example code:\n```\n%s\n```\n", truncate(strings.TrimSpace(ex.Snippet), 1500))
			}
		}
	}
	for _, we := range ch.WorkedExamples {
		fmt.Fprintf(&b, "\n## Worked example: %s\n%s\n", we.Title, we.Problem)
		for i, st := range we.Steps {
			fmt.Fprintf(&b, "%d. %s: %s", i+1, st.Subgoal, st.Explanation)
			if st.State != "" {
				fmt.Fprintf(&b, " (trace: %s)", st.State)
			}
			b.WriteString("\n")
		}
	}
	if len(ch.CheckQuestions) > 0 {
		b.WriteString("\n## Check questions you can draw on (answers are for you)\n")
		for _, q := range ch.CheckQuestions {
			fmt.Fprintf(&b, "- %s", q.Prompt)
			if len(q.Choices) > 0 && q.Answer >= 0 && q.Answer < len(q.Choices) {
				fmt.Fprintf(&b, " Answer: %s.", q.Choices[q.Answer])
			}
			if q.Explanation != "" {
				fmt.Fprintf(&b, " %s", truncate(q.Explanation, 400))
			}
			b.WriteString("\n")
		}
	}
	writeScratchpad(&b, ctx.Scratchpad)
	if ws != nil {
		b.WriteString(ws.Describe())
	}
	return truncate(b.String(), 60000)
}
