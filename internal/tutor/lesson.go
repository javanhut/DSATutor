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
	// Pages is the text already on the lesson page, so the teacher can
	// build on (and point at) what it wrote in earlier turns.
	Pages []LessonPage `json:"pages,omitempty"`
}

// LessonPage is one written section of the lesson page.
type LessonPage struct {
	Section int    `json:"section"`
	Heading string `json:"heading"`
	Body    string `json:"body"`
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
	var pages []LessonPage
	for _, p := range l.Pages {
		if p.Section >= 0 && p.Section < len(secs) && len(pages) < maxLessonSections {
			p.Heading = truncate(strings.TrimSpace(p.Heading), 120)
			p.Body = truncate(strings.TrimSpace(p.Body), 1200)
			pages = append(pages, p)
		}
	}
	l.Pages = pages
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

// SectionDue returns the section this turn should write on the lesson page,
// or -1 when none is due: the current section until it is written, then the
// next one once the learner asks to move on or the teacher's reply moves on
// (models often announce the next section without writing it).
func (l *Lesson) SectionDue(lastUser, reply string) int {
	if !l.Valid() {
		return -1
	}
	if !l.HasArticle(l.Current) {
		return l.Current
	}
	next := l.Current + 1
	if next < len(l.Sections) && !l.HasArticle(next) && (wantsToMoveOn(lastUser) || l.replyMovesOn(reply, next)) {
		return next
	}
	return -1
}

func (l *Lesson) replyMovesOn(reply string, next int) bool {
	r := strings.ToLower(reply)
	if r == "" {
		return false
	}
	if strings.Contains(r, fmt.Sprintf("section %d", next)) || strings.Contains(r, "next section") ||
		strings.Contains(r, "move on") || strings.Contains(r, "moving on") {
		return true
	}
	h := strings.ToLower(l.Sections[next].Heading)
	return len(h) >= 8 && strings.Contains(r, h)
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
	if len(l.Pages) > 0 {
		b.WriteString("\n## Already written on the lesson page (the learner can read it; point at its words with highlight, target lesson-N; do not write these sections again)\n")
		for _, p := range l.Pages {
			fmt.Fprintf(b, "### lesson-%d: %s\n%s\n", p.Section, p.Heading, p.Body)
		}
	}
}

// PlanLesson asks the model for a fresh lesson plan on the chapter's topic.
func (c *Client) PlanLesson(ctx context.Context, model string, ch *chapter.Chapter, learner *LearnerProfile) (*Lesson, error) {
	var sys strings.Builder
	sys.WriteString("You are an expert teacher of data structures and algorithms designing a live, one-on-one lesson for a learner preparing for coding interviews. You write the lesson yourself, from your own knowledge; there is no textbook page.\n\n# Topic\n")
	writeLessonTopic(&sys, ch)
	writeLearnerProfile(&sys, learner)

	ask := fmt.Sprintf(`Plan the lesson as %d to %d short sections taught in order, one idea each. Each section teaches a concept, not just an example: its goal says what the learner will understand (an idea, a reason, a trade-off), not which example they will trace. A good arc: the problem this solves and why the slow way is slow, the core idea and its intuition, how it works step by step, why it works (the invariant), its cost and trade-offs, common mistakes, and when to reach for it in interviews, ending with a short exercise. Adapt the depth to the learner model. Respond with JSON only: "title" (the lesson's name) and "sections", each with "heading" (a few words) and "goal" (one sentence: what the learner can do after it).`, minLessonSections, maxLessonSections)

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

// articleFields describe a lesson-page section as separate parts, so the
// concept itself (not just an example of it) is always written down: models
// asked for free text tend to walk through an example and never state the
// idea in general terms.
func articleFields() map[string]any {
	return map[string]any{
		"heading":    map[string]any{"type": "string", "description": "the section's heading"},
		"concept":    map[string]any{"type": "string", "description": "the idea itself in general terms, not tied to any example: what it is, how it works, and what it costs (complexity in plain text like O(n)), 2 to 4 sentences"},
		"why":        map[string]any{"type": "string", "description": "why it works or why it matters: the reason, rule or invariant behind it, 1 to 3 sentences"},
		"example":    map[string]any{"type": "string", "description": "one tiny concrete example that illustrates the concept, 1 to 3 sentences, using the same numbers you will show on the scratchpad"},
		"key_points": map[string]any{"type": "array", "items": map[string]any{"type": "string"}, "description": "2 to 4 short takeaways worth remembering"},
		"code":       map[string]any{"type": "string", "description": "optional: a short Python snippet (under 20 lines) shown in the section"},
	}
}

func articleTool(sections int) Tool {
	props := articleFields()
	section := map[string]any{"type": "integer", "description": "index of the plan section this article teaches"}
	if sections > 0 {
		section["minimum"], section["maximum"] = 0, sections-1
	}
	props["section"] = section
	return fn("write_lesson",
		"Write one section of your plan onto the lesson page, the learner's reference text for this lesson: you write it yourself, there is no textbook. The page starts blank and fills up section by section as you teach; each section is written once. Write it when you start a section, before any sketch, then point at it with highlight while you teach. Teach the concept first: state the idea in general terms (concept), give the reason it works (why), and only then illustrate it with one tiny example. No question in it; ask with ask_learner. Examples you draw or run go on the scratchpad.",
		props, "section", "heading", "concept", "why", "example")
}

// articleBody composes the section's Markdown: concept, reason, example,
// takeaways. A model that sent a free-text body instead keeps it.
func articleBody(args map[string]any) string {
	concept := strings.TrimSpace(str(args, "concept"))
	if concept == "" {
		return strings.TrimSpace(str(args, "body"))
	}
	parts := []string{concept}
	if why := strings.TrimSpace(str(args, "why")); why != "" {
		parts = append(parts, "**Why it works.** "+why)
	}
	if ex := strings.TrimSpace(str(args, "example")); ex != "" {
		parts = append(parts, "**Example.** "+ex)
	}
	var points []string
	for _, p := range listArg(args, "key_points") {
		if t, ok := p.(string); ok && strings.TrimSpace(t) != "" && len(points) < 5 {
			points = append(points, "- "+strings.TrimSpace(t))
		}
	}
	if len(points) > 0 {
		parts = append(parts, "**Remember:**\n"+strings.Join(points, "\n"))
	}
	return strings.Join(parts, "\n\n")
}

func executeArticle(args map[string]any, l *Lesson) (string, Action) {
	heading := truncate(strings.TrimSpace(str(args, "heading")), 120)
	body := articleBody(args)
	if heading == "" || len(strings.Fields(body)) < 15 {
		return "Error: write_lesson needs a heading and the concept explained in a few sentences (concept, why, example).", nil
	}
	if str(args, "concept") != "" && len(strings.Fields(str(args, "concept"))) < 12 {
		return "Error: concept is too thin. State the idea in general terms: what it is, how it works, and what it costs, in 2 to 4 sentences, before any example.", nil
	}
	section := intArg(args, "section")
	if l.Valid() && (section < 0 || section >= len(l.Sections)) {
		return fmt.Sprintf("Error: section must be between 0 and %d.", len(l.Sections)-1), nil
	}
	if l.Valid() && l.HasArticle(section) {
		return fmt.Sprintf("Error: section %d is already on the lesson page. Don't write it again: point at its words with highlight (target lesson-%d) or show a new example on the scratchpad.", section, section), nil
	}
	act := Action{"name": "write_lesson", "section": section, "heading": heading, "body": truncate(body, maxArticleChars)}
	if l.Valid() {
		act["total"] = len(l.Sections)
		l.Current = section
		l.Written = append(l.Written, section)
	}
	if code := stripLineNumbers(strings.TrimSpace(str(args, "code"))); code != "" {
		act["code"] = truncate(code, maxArticleCode)
	}
	return fmt.Sprintf("Section %d is now on the lesson page (target lesson-%d). Teach it without describing what you did: the learner can see it.", section, section), act
}
