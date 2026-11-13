package tutor

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"dsatutor/internal/chapter"
)

// Teacher Mode does not teach from the chapter's prewritten lesson. It plans
// its own lesson for the chapter's topic (PlanLesson); the page starts blank
// and the teacher writes each section onto it as it goes (write_lesson),
// pointing at what it wrote as reference. The scratchpad holds only the
// examples: sketches, diagrams and traced code. The chapter supplies just
// the topic and its scope: title, summary, objectives and concept names. The
// prewritten pages stay what the normal chapter view and regular tutor use.

// Lesson is Teacher Mode's own plan for one chapter. The browser keeps it
// with the chapter's lesson conversation and sends it back with each turn.
type Lesson struct {
	Title    string          `json:"title"`
	Sections []LessonSection `json:"sections"`
	// Current is the section being taught; Written lists the sections already
	// written on the lesson page.
	Current int   `json:"current"`
	Written []int `json:"written,omitempty"`
}

// LessonSection is one part of the plan: what it covers and what the
// learner should be able to do after it.
type LessonSection struct {
	Heading string `json:"heading"`
	Goal    string `json:"goal"`
}

const (
	minLessonSections = 4
	maxLessonSections = 8
	maxArticleChars   = 4000
	maxArticleCode    = 2500
)

// Valid reports whether the plan has sections to teach.
func (l *Lesson) Valid() bool { return l != nil && len(l.Sections) > 0 }

// Clean bounds a plan that came from the browser or the model.
func (l *Lesson) Clean() {
	l.Title = truncate(strings.TrimSpace(l.Title), 120)
	var secs []LessonSection
	for _, s := range l.Sections {
		s.Heading = truncate(strings.TrimSpace(s.Heading), 120)
		s.Goal = truncate(strings.TrimSpace(s.Goal), 300)
		if s.Heading != "" && len(secs) < maxLessonSections {
			secs = append(secs, s)
		}
	}
	l.Sections = secs
	l.Current = min(max(l.Current, 0), max(len(secs)-1, 0))
	var written []int
	seen := map[int]bool{}
	for _, i := range l.Written {
		if i >= 0 && i < len(secs) && !seen[i] {
			seen[i] = true
			written = append(written, i)
		}
	}
	l.Written = written
}

// HasArticle reports whether section i is written on the lesson page.
func (l *Lesson) HasArticle(i int) bool {
	for _, w := range l.Written {
		if w == i {
			return true
		}
	}
	return false
}

// SectionDue returns the section whose article this turn should write, or
// -1 when none is due: the current section until it has an article, then
// the next one once the learner asks to move on.
func (l *Lesson) SectionDue(lastUser string) int {
	if !l.Valid() {
		return -1
	}
	if !l.HasArticle(l.Current) {
		return l.Current
	}
	if next := l.Current + 1; next < len(l.Sections) && wantsToMoveOn(lastUser) {
		return next
	}
	return -1
}

func wantsToMoveOn(msg string) bool {
	m := strings.ToLower(strings.TrimSpace(msg))
	if m == "next" || m == "next." {
		return true
	}
	for _, cue := range []string{"move on", "next idea", "next section", "next part", "next topic", "keep going", "continue the lesson"} {
		if strings.Contains(m, cue) {
			return true
		}
	}
	return false
}

// writeLessonTopic writes what Teacher Mode may take from the chapter: its
// topic and scope, never its prewritten explanations or examples.
func writeLessonTopic(b *strings.Builder, ch *chapter.Chapter) {
	fmt.Fprintf(b, "Chapter %d: %s.", ch.Number, ch.Title)
	if s := strings.TrimSpace(ch.Summary); s != "" {
		fmt.Fprintf(b, " Scope: %s", truncate(s, 400))
	}
	b.WriteString("\n")
	if len(ch.Objectives) > 0 {
		b.WriteString("By the end the learner should be able to:\n")
		for _, o := range ch.Objectives {
			fmt.Fprintf(b, "- %s\n", truncate(o, 200))
		}
	}
	if len(ch.Concepts) > 0 {
		names := make([]string, 0, len(ch.Concepts))
		for _, c := range ch.Concepts {
			names = append(names, c.Name)
		}
		fmt.Fprintf(b, "Concepts to cover: %s.\n", strings.Join(names, "; "))
	}
}

// writeLesson writes the plan into a prompt, marking where the lesson is.
func writeLesson(b *strings.Builder, l *Lesson) {
	if !l.Valid() {
		b.WriteString("\n## Your lesson plan\nNo plan yet: teach the topic in a sensible order, one section at a time, writing each section on the lesson page with write_lesson.\n")
		return
	}
	fmt.Fprintf(b, "\n## Your lesson plan: %s\n", l.Title)
	for i, s := range l.Sections {
		status := "not started"
		switch {
		case i == l.Current && l.HasArticle(i):
			status = "CURRENT, written on the page"
		case i == l.Current:
			status = "CURRENT, not written yet"
		case l.HasArticle(i):
			status = "done"
		}
		fmt.Fprintf(b, "%d. %s: %s [%s]\n", i, s.Heading, s.Goal, status)
	}
}

// PlanLesson asks the model for a fresh lesson plan on the chapter's topic.
func (c *Client) PlanLesson(ctx context.Context, model string, ch *chapter.Chapter, learner *LearnerProfile) (*Lesson, error) {
	var sys strings.Builder
	sys.WriteString("You are an expert teacher of data structures and algorithms designing a live, one-on-one lesson for a learner preparing for coding interviews. You write the lesson yourself, from your own knowledge; there is no textbook page.\n\n# Topic\n")
	writeLessonTopic(&sys, ch)
	writeLearnerProfile(&sys, learner)

	ask := fmt.Sprintf(`Plan the lesson as %d to %d short sections taught in order, one idea each. A good arc: the problem this solves and the slow way first, the core idea and its intuition, a worked example on a tiny input, why it works (the invariant), common mistakes, and when to reach for it in interviews, ending with a short exercise. Adapt the depth to the learner model. Respond with JSON only: "title" (the lesson's name) and "sections", each with "heading" (a few words) and "goal" (one sentence: what the learner can do after it).`, minLessonSections, maxLessonSections)

	str := map[string]any{"type": "string", "minLength": 1}
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": str,
			"sections": map[string]any{
				"type": "array", "minItems": minLessonSections, "maxItems": maxLessonSections,
				"items": map[string]any{
					"type":       "object",
					"properties": map[string]any{"heading": str, "goal": str},
					"required":   []string{"heading", "goal"},
				},
			},
		},
		"required": []string{"title", "sections"},
	}
	raw, _, err := c.ChatJSON(ctx, model, []Message{
		{Role: "system", Content: sys.String()},
		{Role: "user", Content: ask},
	}, schema)
	if err != nil {
		return nil, err
	}
	var l Lesson
	if err := json.Unmarshal(raw, &l); err != nil {
		return nil, fmt.Errorf("lesson plan was malformed JSON: %w", err)
	}
	l.Clean()
	if !l.Valid() {
		return nil, fmt.Errorf("lesson plan had no sections")
	}
	if l.Title == "" {
		l.Title = ch.Title
	}
	return &l, nil
}

func articleTool(sections int) Tool {
	section := map[string]any{"type": "integer", "description": "index of the plan section this article teaches"}
	if sections > 0 {
		section["minimum"], section["maximum"] = 0, sections-1
	}
	return fn("write_lesson",
		"Write one section of your plan onto the lesson page, the learner's reference text for this lesson: you write it yourself, there is no textbook. The page starts blank and fills up section by section as you teach; rewriting a section replaces it. Write it when you start a section, before any sketch, then point at it with highlight while you teach. 80 to 200 words of Markdown: a plain-language explanation with one tiny concrete example (the same numbers you will draw or trace on the scratchpad), short bullet points for the key facts, and the complexity in plain text like O(n). No question in it; ask with ask_learner. Examples you draw or run go on the scratchpad, not here.",
		map[string]any{
			"section": section,
			"heading": map[string]any{"type": "string", "description": "the section's heading"},
			"body":    map[string]any{"type": "string", "description": "the article in Markdown (paragraphs, bullet lists, `inline code`); no headings, no code fences"},
			"code":    map[string]any{"type": "string", "description": "optional: a short Python snippet (under 20 lines) shown in the section"},
		}, "section", "heading", "body")
}

func executeArticle(args map[string]any, l *Lesson) (string, Action) {
	heading := truncate(strings.TrimSpace(str(args, "heading")), 120)
	body := strings.TrimSpace(str(args, "body"))
	if heading == "" || len(strings.Fields(body)) < 15 {
		return "Error: write_lesson needs a heading and a body of at least a few sentences.", nil
	}
	section := intArg(args, "section")
	if l.Valid() && (section < 0 || section >= len(l.Sections)) {
		return fmt.Sprintf("Error: section must be between 0 and %d.", len(l.Sections)-1), nil
	}
	act := Action{"name": "write_lesson", "section": section, "heading": heading, "body": truncate(body, maxArticleChars)}
	if l.Valid() {
		act["total"] = len(l.Sections)
		l.Current = section
		if !l.HasArticle(section) {
			l.Written = append(l.Written, section)
		}
	}
	if code := stripLineNumbers(strings.TrimSpace(str(args, "code"))); code != "" {
		act["code"] = truncate(code, maxArticleCode)
	}
	return fmt.Sprintf("Section %d is written on the lesson page (target id lesson-%d; highlight its words as you teach). Now teach it: one scratchpad sketch or traced run if it helps, then one ask_learner question.", section, section), act
}
