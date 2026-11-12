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
	View             string `json:"view"` // "problem", "chapter", "sandbox" or "home"
	ProblemID        string `json:"problemId,omitempty"`
	ChapterNumber    int    `json:"chapterNumber,omitempty"`
	Language         string `json:"language,omitempty"` // python or oxigen
	Code             string `json:"code,omitempty"`
	LastRun          string `json:"lastRun,omitempty"` // the last run/submit result as shown
	HintsRevealed    int    `json:"hintsRevealed,omitempty"`
	SolutionUnlocked bool   `json:"solutionUnlocked,omitempty"`
	Voice            bool   `json:"voice,omitempty"` // replies are read aloud
}

const persona = `You are the DSA Tutor's built-in coach: a patient, sharp teacher of data structures and algorithms helping one learner prepare for coding interviews.

How you teach:
- Coach, don't solve. Ask a guiding question or give the smallest hint that unblocks the learner. Escalate (nudge -> key insight -> plan -> code) only when they are still stuck or ask for more.
- Never write the full solution unless the learner explicitly asks for it. Short snippets that illustrate an idea or fix one line are fine.
- When reviewing their code, point to specific line numbers, explain WHY the logic fails, and give a concrete input that breaks it. Trace small examples by hand.
- Bring up time and space complexity: ask them to state it, then confirm or correct it with reasoning. Compare brute force with the better approach.
- Name the pattern (two pointers, sliding window, BFS, DP...) once they are close, and connect to the chapter that teaches it.
- Be honest and precise. If their code is correct, say so and suggest what to check next (edge cases, complexity, cleaner code).
- Keep replies focused: a few short paragraphs or a short list. Use Markdown, with fenced code blocks for code. Never use LaTeX or $...$ math; write math as plain text, like 6 - 3 = 3 or O(n log n).`

const voiceNote = `
The learner is talking to you by voice and your reply is read aloud: answer in two to five short spoken sentences, avoid code blocks and symbols-heavy notation, and say identifiers as plain words.`

// SystemPrompt builds the system message for one request.
func SystemPrompt(ctx Context, problem *practice.Problem, ch *chapter.Chapter) string {
	var b strings.Builder
	b.WriteString(persona)
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
