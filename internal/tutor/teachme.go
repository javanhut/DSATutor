package tutor

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"strings"
)

// Teach Me builds a short lesson on any DSA topic the learner names, not
// tied to a chapter. Its shape follows the research summarised in
// docs/teachme.md: a prior-knowledge prompt, direct instruction in small
// chunks each followed by a retrieval check, a worked example revealed one
// step at a time, faded (fill-in-the-blank) practice, an explain-it-back
// prompt graded with feedback, and flashcards the page schedules for spaced
// review.

// TeachLesson is one generated Teach Me lesson.
type TeachLesson struct {
	Title      string         `json:"title"`
	Topic      string         `json:"topic"`
	Warmup     string         `json:"warmup"`
	BigIdea    string         `json:"big_idea"`
	Chunks     []TeachChunk   `json:"chunks"`
	Worked     TeachWorked    `json:"worked_example"`
	Code       string         `json:"code,omitempty"`
	Faded      TeachFaded     `json:"faded"`
	Pitfalls   []string       `json:"pitfalls"`
	Explain    string         `json:"explain_prompt"`
	Reference  string         `json:"explain_reference"`
	Cards      []TeachCard    `json:"flashcards"`
	Keywords   []string       `json:"practice_keywords"`
	Complexity string         `json:"complexity,omitempty"`
	Problems   []TeachProblem `json:"problems,omitempty"`
}

// TeachChunk is one piece of direct instruction and the check that follows it.
type TeachChunk struct {
	Heading     string     `json:"heading"`
	Explanation string     `json:"explanation"`
	Check       TeachCheck `json:"check"`
}

// TeachCheck is a multiple-choice retrieval question. The model sends the
// correct choice as text and each wrong one with the misconception behind it
// (models misnumber an answer index, and a shared explanation that says "the
// first choice" breaks once choices are shuffled); Clean turns that into
// shuffled Choices, the Answer index and per-choice Feedback.
type TeachCheck struct {
	Question string       `json:"question"`
	Correct  string       `json:"correct,omitempty"`
	Why      string       `json:"why"`
	Wrong    []TeachWrong `json:"wrong,omitempty"`
	Choices  []string     `json:"choices,omitempty"`
	Answer   int          `json:"answer"`
	Feedback []string     `json:"feedback,omitempty"`
}

// TeachWrong is a wrong choice and the misconception it reflects.
type TeachWrong struct {
	Choice        string `json:"choice"`
	Misconception string `json:"misconception"`
}

// TeachWorked is a worked example the page reveals one step at a time.
type TeachWorked struct {
	Problem string      `json:"problem"`
	Steps   []TeachStep `json:"steps"`
}

// TeachStep is one step of the worked example.
type TeachStep struct {
	Do  string `json:"do"`
	Why string `json:"why"`
}

// TeachFaded is code with one blank for the learner to fill in. The model
// sends complete code and names the piece to hide (Blank); Clean cuts it out,
// so the answer always matches the gap.
type TeachFaded struct {
	Prompt string `json:"prompt"`
	Code   string `json:"code"`
	Blank  string `json:"blank,omitempty"`
	Answer string `json:"answer"`
	Hint   string `json:"hint"`
}

// TeachCard is a flashcard for spaced review.
type TeachCard struct {
	Front string `json:"front"`
	Back  string `json:"back"`
}

// TeachProblem is a practice-bank problem that uses the topic.
type TeachProblem struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	Category   string `json:"category"`
}

// Cloud models served through Ollama don't enforce the format schema, and
// some rename keys (deepseek sends chunk "text", wrong-choice "answer",
// "faded_practice", "explain_it_back", worked-example "input"). These
// decoders accept those aliases so one renamed key doesn't empty the lesson.

func (l *TeachLesson) UnmarshalJSON(b []byte) error {
	type plain TeachLesson
	var v struct {
		plain
		FadedPractice *TeachFaded `json:"faded_practice"`
		ExplainItBack string      `json:"explain_it_back"`
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*l = TeachLesson(v.plain)
	if l.Faded.Code == "" && v.FadedPractice != nil {
		l.Faded = *v.FadedPractice
	}
	l.Explain = orAlias(l.Explain, v.ExplainItBack)
	return nil
}

func (c *TeachChunk) UnmarshalJSON(b []byte) error {
	type plain TeachChunk
	var v struct {
		plain
		Text    string `json:"text"`
		Content string `json:"content"`
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*c = TeachChunk(v.plain)
	c.Explanation = orAlias(c.Explanation, v.Text, v.Content)
	return nil
}

func (w *TeachWrong) UnmarshalJSON(b []byte) error {
	type plain TeachWrong
	var v struct {
		plain
		Answer string `json:"answer"`
		Text   string `json:"text"`
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*w = TeachWrong(v.plain)
	w.Choice = orAlias(w.Choice, v.Answer, v.Text)
	return nil
}

func (w *TeachWorked) UnmarshalJSON(b []byte) error {
	type plain TeachWorked
	var v struct {
		plain
		Input    string `json:"input"`
		Question string `json:"question"`
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*w = TeachWorked(v.plain)
	w.Problem = orAlias(w.Problem, v.Input, v.Question)
	return nil
}

func orAlias(ss ...string) string {
	for _, s := range ss {
		if strings.TrimSpace(s) != "" {
			return s
		}
	}
	return ""
}

// FadedBlank marks the gap in faded practice code.
const FadedBlank = "____"

func teachSchema() map[string]any {
	str := map[string]any{"type": "string", "minLength": 1}
	strs := func(min, max int) map[string]any {
		return map[string]any{"type": "array", "items": str, "minItems": min, "maxItems": max}
	}
	obj := func(props map[string]any, req ...string) map[string]any {
		return map[string]any{"type": "object", "properties": props, "required": req}
	}
	check := obj(map[string]any{
		"question": str,
		"correct":  str,
		"why":      str,
		"wrong": map[string]any{"type": "array", "minItems": 2, "maxItems": 3, "items": obj(map[string]any{
			"choice": str, "misconception": str,
		}, "choice", "misconception")},
	}, "question", "correct", "why", "wrong")
	return obj(map[string]any{
		"title":    str,
		"warmup":   str,
		"big_idea": str,
		"chunks": map[string]any{"type": "array", "minItems": 2, "maxItems": 4, "items": obj(map[string]any{
			"heading": str, "explanation": str, "check": check,
		}, "heading", "explanation", "check")},
		"worked_example": obj(map[string]any{
			"problem": str,
			"steps": map[string]any{"type": "array", "minItems": 3, "maxItems": 7, "items": obj(map[string]any{
				"do": str, "why": str,
			}, "do", "why")},
		}, "problem", "steps"),
		"code":              map[string]any{"type": "string"},
		"complexity":        str,
		"faded":             obj(map[string]any{"prompt": str, "code": str, "blank": str, "hint": str}, "prompt", "code", "blank", "hint"),
		"pitfalls":          strs(2, 4),
		"explain_prompt":    str,
		"explain_reference": str,
		"flashcards": map[string]any{"type": "array", "minItems": 3, "maxItems": 6, "items": obj(map[string]any{
			"front": str, "back": str,
		}, "front", "back")},
		"practice_keywords": strs(1, 6),
	}, "title", "warmup", "big_idea", "chunks", "worked_example", "complexity", "faded", "pitfalls",
		"explain_prompt", "explain_reference", "flashcards", "practice_keywords")
}

const teachSystem = `You are an expert teacher of data structures and algorithms. A learner preparing for coding interviews asked you to teach them one topic, quickly and so it sticks. You write the lesson from your own knowledge.

Teach the way learning research says works:
- Direct instruction first: explain the idea clearly and concretely before asking the learner to use it. Never leave a novice to discover it alone.
- "warmup": a question (ending in "?") the learner tries to answer before being taught, to activate what they already know. Not a statement and not the answer.
- "big_idea": the whole technique in one or two sentences.
- Small chunks: each chunk has a "heading" of 2 to 6 words and is one idea in 2 to 5 sentences of plain language, with complexity written like O(n). Inline code in backticks.
- Retrieval after every chunk: a multiple-choice check that makes the learner recall or apply the idea (predict, choose, spot the bug), not recognise a phrase you just used. "correct" is the one right answer to the question and "why" says why it is right. "wrong" holds 2 or 3 plausible wrong answers to the same question, each a real misconception, with "misconception" telling the learner, as "you", what that choice gets wrong. Never refer to choices by letter or position.
- A worked example: a tiny concrete input, solved in 3 to 7 steps. Each step says what to do ("do") and why ("why").
- Faded practice: a short, complete, correct Python function (under 15 lines) using the idea. "blank" is copied exactly from that code: the one key expression or line the learner must supply (the heart of the technique, not boilerplate). The page hides it; you do not write the blank yourself.
- An explain-it-back prompt asking the learner to explain the idea in their own words, plus a model answer ("explain_reference") of 2 to 4 sentences.
- Flashcards for spaced review: a question on the front that forces recall (when to use it, the invariant, the cost, the key step), a short answer on the back.
- "practice_keywords": short words a practice problem about this would carry in its title, tags or category (for example "sliding window", "heap", "union find").
- "code": optionally, a clean reference implementation in Python under 25 lines, no Markdown fences.

If the request is not about data structures, algorithms or coding interviews, teach the closest DSA topic it relates to. Keep it tight: the learner wants to practise quickly.

Respond with JSON only, using exactly these keys and no others:
{"title", "warmup", "big_idea",
 "chunks": [{"heading", "explanation", "check": {"question", "correct", "why", "wrong": [{"choice", "misconception"}]}}],
 "worked_example": {"problem", "steps": [{"do", "why"}]},
 "code", "complexity",
 "faded": {"prompt", "code", "blank", "hint"},
 "pitfalls": [string], "explain_prompt", "explain_reference",
 "flashcards": [{"front", "back"}], "practice_keywords": [string]}`

// TeachTopic generates a Teach Me lesson on topic.
func (c *Client) TeachTopic(ctx context.Context, model, topic, level string) (*TeachLesson, string, error) {
	topic = truncate(strings.TrimSpace(topic), 300)
	if topic == "" {
		return nil, "", fmt.Errorf("no topic")
	}
	switch level {
	case "beginner", "intermediate", "advanced":
	default:
		level = "intermediate"
	}
	ask := fmt.Sprintf("Teach me: %s\nMy level: %s.", topic, level)
	raw, name, err := c.ChatJSON(ctx, model, []Message{
		{Role: "system", Content: teachSystem},
		{Role: "user", Content: ask},
	}, teachSchema())
	if err != nil {
		return nil, name, err
	}
	var l TeachLesson
	if err := json.Unmarshal(raw, &l); err != nil {
		return nil, name, fmt.Errorf("lesson was malformed JSON: %w", err)
	}
	l.Topic = topic
	faded := l.Faded
	l.Clean()
	if l.Faded.Code == "" && strings.TrimSpace(faded.Code) != "" {
		// The model's blank wasn't in its code: have it pick the line.
		if f, ok := c.pickBlankLine(ctx, model, topic, faded); ok {
			l.Faded = f
		}
	}
	if len(l.Chunks) == 0 {
		return nil, name, fmt.Errorf("the model returned a lesson with nothing to teach; try again or pick another model")
	}
	return &l, name, nil
}

// Clean bounds what the model sent and builds what the page relies on:
// shuffled check choices with the answer index and per-choice feedback, a
// faded blank cut from complete code, non-empty cards.
func (l *TeachLesson) Clean() {
	t := func(s string, n int) string { return truncate(strings.TrimSpace(s), n) }
	l.Title = t(l.Title, 120)
	if l.Title == "" {
		l.Title = l.Topic
	}
	l.Warmup, l.BigIdea = t(l.Warmup, 400), t(l.BigIdea, 400)
	l.Complexity = t(l.Complexity, 200)
	l.Code = stripFences(t(l.Code, 2500))

	var chunks []TeachChunk
	for _, ch := range l.Chunks {
		ch.Heading, ch.Explanation = t(ch.Heading, 120), t(ch.Explanation, 1500)
		if ch.Heading == "" || ch.Explanation == "" {
			continue
		}
		ch.Check = cleanCheck(ch.Check)
		chunks = append(chunks, ch)
		if len(chunks) == 4 {
			break
		}
	}
	l.Chunks = chunks

	l.Worked.Problem = t(l.Worked.Problem, 600)
	var steps []TeachStep
	for _, s := range l.Worked.Steps {
		s.Do, s.Why = t(s.Do, 400), t(s.Why, 400)
		if s.Do != "" && len(steps) < 7 {
			steps = append(steps, s)
		}
	}
	l.Worked.Steps = steps

	f := &l.Faded
	f.Prompt, f.Hint = t(f.Prompt, 400), t(f.Hint, 300)
	f.Code = stripFences(t(f.Code, 1500))
	if blank := strings.TrimSpace(strings.Trim(f.Blank, "`")); blank != "" && !strings.Contains(f.Code, FadedBlank) {
		if code, ok := cutBlank(f.Code, blank); ok {
			f.Code, f.Answer = code, blank
		}
	}
	f.Blank = ""
	if f.Answer == "" || strings.Count(f.Code, FadedBlank) != 1 {
		l.Faded = TeachFaded{}
	} else if f.Prompt == "" {
		f.Prompt = "Fill in the missing piece."
	}

	l.Pitfalls = cleanList(l.Pitfalls, 4, 300)
	l.Keywords = cleanList(l.Keywords, 6, 40)
	l.Explain, l.Reference = t(l.Explain, 300), t(l.Reference, 800)

	var cards []TeachCard
	for _, c := range l.Cards {
		c.Front, c.Back = t(c.Front, 300), t(c.Back, 500)
		if c.Front != "" && c.Back != "" && len(cards) < 6 {
			cards = append(cards, c)
		}
	}
	l.Cards = cards
}

// pickBlankLine asks the model which numbered line of its faded-practice
// code is the key one, and blanks that line.
func (c *Client) pickBlankLine(ctx context.Context, model, topic string, f TeachFaded) (TeachFaded, bool) {
	code := stripFences(truncate(strings.TrimSpace(f.Code), 1500))
	lines := strings.Split(code, "\n")
	var numbered strings.Builder
	for i, line := range lines {
		fmt.Fprintf(&numbered, "%d: %s\n", i+1, line)
	}
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"line":   map[string]any{"type": "integer", "minimum": 1, "maximum": len(lines)},
			"prompt": map[string]any{"type": "string", "description": "one sentence telling the learner what the hidden line must do, without giving it away"},
			"hint":   map[string]any{"type": "string", "description": "a nudge toward that line, without the code"},
		},
		"required": []string{"line", "prompt", "hint"},
	}
	raw, _, err := c.ChatJSON(ctx, model, []Message{{Role: "user", Content: fmt.Sprintf(
		"A learner studying %s will fill in one hidden line of this Python code. Pick the single line that is the heart of the technique (not a def, return of a plain variable, or setup boilerplate). Answer as JSON with its line number, a prompt saying what that hidden line must do (without giving it away) and a hint.\n\n%s", topic, numbered.String())}}, schema)
	if err != nil {
		return f, false
	}
	var pick struct {
		Line         int
		Prompt, Hint string
	}
	if json.Unmarshal(raw, &pick) != nil || pick.Line < 1 || pick.Line > len(lines) {
		return f, false
	}
	line := lines[pick.Line-1]
	answer := strings.TrimSpace(line)
	if answer == "" || strings.HasPrefix(answer, "def ") || strings.HasPrefix(answer, "#") {
		return f, false
	}
	lines[pick.Line-1] = line[:len(line)-len(strings.TrimLeft(line, " \t"))] + FadedBlank
	f.Code, f.Answer, f.Blank = strings.Join(lines, "\n"), answer, ""
	f.Prompt, f.Hint = truncate(strings.TrimSpace(pick.Prompt), 400), truncate(strings.TrimSpace(pick.Hint), 300)
	if f.Prompt == "" {
		f.Prompt = "Fill in the missing line."
	}
	return f, true
}

// cutBlank replaces blank in code with FadedBlank. The last occurrence is
// used, since that is usually where the idea is applied after setup that
// mentions the same names. Models often copy the line with different
// spacing, so a line matching it with whitespace ignored is cut whole.
func cutBlank(code, blank string) (string, bool) {
	if i := strings.LastIndex(code, blank); i >= 0 {
		return code[:i] + FadedBlank + code[i+len(blank):], true
	}
	squash := func(s string) string { return strings.Join(strings.Fields(s), "") }
	want := squash(blank)
	lines := strings.Split(code, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		if squash(lines[i]) == want {
			indent := lines[i][:len(lines[i])-len(strings.TrimLeft(lines[i], " \t"))]
			lines[i] = indent + FadedBlank
			return strings.Join(lines, "\n"), true
		}
	}
	return code, false
}

func cleanCheck(c TeachCheck) TeachCheck {
	t := func(s string, n int) string { return truncate(strings.TrimSpace(s), n) }
	q, correct, why := t(c.Question, 400), t(c.Correct, 200), t(c.Why, 600)
	if q == "" || correct == "" {
		return TeachCheck{}
	}
	type choice struct{ text, feedback string }
	choices := []choice{{correct, ""}}
	seen := map[string]bool{strings.ToLower(correct): true}
	for _, w := range c.Wrong {
		text := t(w.Choice, 200)
		if text == "" || seen[strings.ToLower(text)] || len(choices) == 4 {
			continue
		}
		seen[strings.ToLower(text)] = true
		choices = append(choices, choice{text, t(w.Misconception, 400)})
	}
	if len(choices) < 2 {
		return TeachCheck{}
	}
	rand.Shuffle(len(choices), func(i, j int) { choices[i], choices[j] = choices[j], choices[i] })
	out := TeachCheck{Question: q, Why: why}
	for i, ch := range choices {
		if ch.text == correct {
			out.Answer = i
		}
		out.Choices = append(out.Choices, ch.text)
		out.Feedback = append(out.Feedback, ch.feedback)
	}
	return out
}

func cleanList(in []string, n, size int) []string {
	var out []string
	for _, s := range in {
		if s = truncate(strings.TrimSpace(s), size); s != "" && len(out) < n {
			out = append(out, s)
		}
	}
	return out
}

func stripFences(s string) string {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "```") {
		return s
	}
	if i := strings.Index(s, "\n"); i >= 0 {
		s = s[i+1:]
	}
	return strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(s), "```"))
}

// TeachFeedback is the grade for an explain-it-back answer.
type TeachFeedback struct {
	Verdict  string   `json:"verdict"` // solid, partial or missing
	Feedback string   `json:"feedback"`
	Missing  []string `json:"missing"`
}

// GradeExplanation gives targeted feedback on the learner's own explanation.
func (c *Client) GradeExplanation(ctx context.Context, model, topic, prompt, reference, answer string) (*TeachFeedback, string, error) {
	answer = truncate(strings.TrimSpace(answer), 3000)
	if answer == "" {
		return nil, "", fmt.Errorf("empty answer")
	}
	sys := "You give a learner quick, specific feedback on their explanation of a data structures and algorithms idea. Judge understanding, not wording or grammar. Say what they got right in one sentence, then the single most important gap or misconception and how to fix it. 2 to 4 sentences total, speaking to the learner as \"you\". \"verdict\" is solid (the key idea and why it works are there), partial (the idea is there but the reason or cost is missing or shaky) or missing (the key idea is absent or wrong). \"missing\" lists up to 3 short points they left out. Respond with JSON only."
	user := fmt.Sprintf("Topic: %s\nQuestion: %s\nA model answer: %s\n\nThe learner's answer:\n%s",
		truncate(topic, 300), truncate(prompt, 300), truncate(reference, 800), answer)
	str := map[string]any{"type": "string"}
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"verdict":  map[string]any{"type": "string", "enum": []string{"solid", "partial", "missing"}},
			"feedback": str,
			"missing":  map[string]any{"type": "array", "items": str, "maxItems": 3},
		},
		"required": []string{"verdict", "feedback", "missing"},
	}
	raw, name, err := c.ChatJSON(ctx, model, []Message{{Role: "system", Content: sys}, {Role: "user", Content: user}}, schema)
	if err != nil {
		return nil, name, err
	}
	var fb TeachFeedback
	if err := json.Unmarshal(raw, &fb); err != nil {
		return nil, name, fmt.Errorf("feedback was malformed JSON: %w", err)
	}
	switch fb.Verdict {
	case "solid", "partial", "missing":
	default:
		fb.Verdict = "partial"
	}
	fb.Feedback = truncate(strings.TrimSpace(fb.Feedback), 1000)
	fb.Missing = cleanList(fb.Missing, 3, 200)
	return &fb, name, nil
}
