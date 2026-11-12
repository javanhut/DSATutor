package tutor

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"dsatutor/internal/practice"
)

// Mock interviews: the tutor plays a Google interviewer on one practice
// problem. The candidate can't run code, hints and the solution are hidden,
// and at the end the model writes a scorecard on Google's four coding
// criteria from the transcript, the final code and the hidden test results.

// Interview is the clock of a running mock interview, sent with every chat
// request so the interviewer can manage time.
type Interview struct {
	ElapsedSeconds  int    `json:"elapsedSeconds"`
	DurationSeconds int    `json:"durationSeconds"`
	Level           string `json:"level,omitempty"` // L3, L4 or L5
}

func (iv *Interview) minutes() (elapsed, left int) {
	elapsed = iv.ElapsedSeconds / 60
	left = (iv.DurationSeconds - iv.ElapsedSeconds + 59) / 60
	if left < 0 {
		left = 0
	}
	return elapsed, left
}

func (iv *Interview) level() string {
	switch iv.Level {
	case "L3", "L4", "L5":
		return iv.Level
	}
	return "L4"
}

const interviewerPersona = `You are a software engineer at Google conducting a live coding interview with a candidate for an %s software engineer role. The interview runs in a shared document: the candidate types code but cannot run it.

How you behave, like a real interviewer:
- Opening: greet the candidate in one sentence, then present the problem in your own words with one small example. Never name the pattern, the intended approach, the difficulty or a LeetCode number.
- Clarifying questions: answer them precisely and briefly from the problem statement and constraints. If something isn't specified, pick a reasonable answer and say it.
- Let the candidate drive. Do not teach or lecture. If they are silent or going in circles, ask one probing question ("What would a brute force look like?", "What's the bottleneck there?").
- Hints: give one only when asked or when the candidate is clearly stuck, keep it as small as possible, and remember that it counts against them.
- Before coding: make sure they described an approach and its time and space complexity. If they start coding without that, ask for it.
- While they code: stay quiet unless asked. You may ask about a specific line. Never point out a bug directly; ask them to test instead.
- After coding: ask them to walk through an example by hand and to think about edge cases, then for the final complexity. Let them do the tracing; never trace it for them. If time remains, ask one follow-up (a variation, or how it scales).
- There is no Run or Submit button and no test suite in this interview. When you are done, tell them they can press End interview.
- Time: you are told how many minutes are left. With about 10 minutes left, nudge them to finish coding; at 0, thank them and close the interview.
- Never write the solution or more than a line of code. Never say the code is correct before they have tested it. Don't grade or give feedback during the interview.
- Speak like a person: one to four sentences per reply, no headings or lists.`

// InterviewerPrompt builds the system message for a mock interview turn.
func InterviewerPrompt(ctx Context, p *practice.Problem, ws *Workspace) string {
	iv := ctx.Interview
	var b strings.Builder
	fmt.Fprintf(&b, interviewerPersona, iv.level())
	if ctx.Voice {
		b.WriteString("\nThe candidate is talking to you by voice and your replies are read aloud.")
	}
	elapsed, left := iv.minutes()
	fmt.Fprintf(&b, "\n\n# Clock\n%d minutes elapsed, %d minutes left of %d.\n", elapsed, left, iv.DurationSeconds/60)

	fmt.Fprintf(&b, "\n# The problem you are asking\n**%s**\n\n%s\n", p.Title, strings.TrimSpace(p.Description))
	if len(p.Constraints) > 0 {
		b.WriteString("\nConstraints:\n")
		for _, c := range p.Constraints {
			fmt.Fprintf(&b, "- %s\n", c)
		}
	}
	for i, ex := range p.Examples {
		fmt.Fprintf(&b, "\nExample %d: input %s -> output %s", i+1, ex.Input, ex.Output)
		if ex.Explanation != "" {
			fmt.Fprintf(&b, " (%s)", ex.Explanation)
		}
	}

	s := p.Solution
	b.WriteString("\n\n# Interviewer notes (never reveal them)\n")
	if s.BruteForce != "" {
		fmt.Fprintf(&b, "- Brute force: %s\n", s.BruteForce)
	}
	if s.KeyInsight != "" {
		fmt.Fprintf(&b, "- Key insight: %s\n", s.KeyInsight)
	}
	if tc, sc := firstNonEmpty(s.TimeComplexity, p.TimeComplexity), firstNonEmpty(s.SpaceComplexity, p.SpaceComplexity); tc != "" {
		fmt.Fprintf(&b, "- Expected optimal complexity: time %s, space %s\n", tc, sc)
	}
	for _, m := range s.CommonMistakes {
		fmt.Fprintf(&b, "- Watch for: %s\n", m)
	}
	for _, q := range p.SelfCheck {
		fmt.Fprintf(&b, "- Possible follow-up: %s\n", q.Prompt)
	}

	if code := strings.TrimSpace(ctx.Code); code != "" {
		fmt.Fprintf(&b, "\n# The candidate's document right now (%s, with line numbers)\n```\n%s\n```\n", firstNonEmpty(ctx.Language, "python"), numberLines(truncate(code, maxCodeChars)))
		b.WriteString("The `N | ` prefixes are line numbers for you to cite; they are not part of the code.\n")
	} else {
		b.WriteString("\nThe candidate hasn't written any code yet.\n")
	}
	if ws != nil {
		b.WriteString(ws.Describe())
		b.WriteString("You may use highlight to point at a line of their code when you ask about it. You have no other tools.\n")
	}
	return b.String()
}

// Scorecard is the model's written feedback after a mock interview.
type Scorecard struct {
	Verdict        string            `json:"verdict"`
	Summary        string            `json:"summary"`
	Ratings        map[string]Rating `json:"ratings"`
	Complexity     string            `json:"complexity"`
	Strengths      []string          `json:"strengths"`
	Improvements   []string          `json:"improvements"`
	HintsGiven     int               `json:"hints_given"`
	ReachedOptimal bool              `json:"reached_optimal"`
}

// Rating is one criterion's score, 1 (poor) to 4 (excellent), with the
// evidence it rests on.
type Rating struct {
	Score    int    `json:"score"`
	Evidence string `json:"evidence"`
}

// Criteria are Google's coding-interview attributes, in display order.
var Criteria = []string{"algorithms", "coding", "communication", "problem_solving"}

var verdicts = []string{"Strong No Hire", "No Hire", "Lean No Hire", "Lean Hire", "Hire", "Strong Hire"}

func scorecardSchema() map[string]any {
	rating := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"score":    map[string]any{"type": "integer", "minimum": 1, "maximum": 4},
			"evidence": map[string]any{"type": "string"},
		},
		"required": []string{"score", "evidence"},
	}
	ratings := map[string]any{}
	for _, c := range Criteria {
		ratings[c] = rating
	}
	strs := map[string]any{"type": "array", "items": map[string]any{"type": "string"}}
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"verdict":         map[string]any{"type": "string", "enum": verdicts},
			"summary":         map[string]any{"type": "string"},
			"ratings":         map[string]any{"type": "object", "properties": ratings, "required": Criteria},
			"complexity":      map[string]any{"type": "string"},
			"strengths":       strs,
			"improvements":    strs,
			"hints_given":     map[string]any{"type": "integer", "minimum": 0},
			"reached_optimal": map[string]any{"type": "boolean"},
		},
		"required": []string{"verdict", "summary", "ratings", "complexity", "strengths", "improvements", "hints_given", "reached_optimal"},
	}
}

const scorerPrompt = `You are a calibrated Google interviewer writing feedback for the hiring committee about a %s coding interview you just ran. Be honest and specific; a kind but inflated score does the candidate no favours.

Rate each criterion from 1 to 4 (1 = poor, 2 = mixed, 3 = good, 4 = excellent), citing concrete moments from the transcript or lines of the final code as evidence:
- algorithms: chose a correct, efficient approach; knew the relevant data structures; stated time and space complexity correctly.
- coding: the final code is correct, clean and idiomatic, handles edge cases, and would compile/run. Weigh the automated test results heavily.
- communication: asked clarifying questions, explained the plan before coding, thought out loud, responded well to questions.
- problem_solving: broke the problem down, started from a brute force and improved it, tested by hand, found and fixed their own bugs, needed few hints.

verdict: the overall recommendation for an %s role. hints_given: how many hints you gave. reached_optimal: whether they reached the expected optimal complexity. complexity: one sentence comparing what they stated with the actual complexity of their final code. strengths and improvements: two to four short, specific bullet sentences each; improvements should say what to do differently next time. summary: two or three sentences.`

// ScoreInput is everything the scorer sees.
type ScoreInput struct {
	Problem    *practice.Problem
	Code       string
	Language   string
	Transcript []Message
	Elapsed    time.Duration
	Duration   time.Duration
	Level      string
	Tests      *practice.SubmitResponse
}

// Score asks model for a scorecard on a finished interview.
func (c *Client) Score(ctx context.Context, model string, in ScoreInput) (*Scorecard, string, error) {
	level := (&Interview{Level: in.Level}).level()
	p := in.Problem
	var b strings.Builder
	fmt.Fprintf(&b, "# Problem: %s\n%s\n", p.Title, strings.TrimSpace(p.Description))
	s := p.Solution
	fmt.Fprintf(&b, "\nExpected optimal: time %s, space %s. Key insight: %s\n", firstNonEmpty(s.TimeComplexity, p.TimeComplexity), firstNonEmpty(s.SpaceComplexity, p.SpaceComplexity), s.KeyInsight)

	fmt.Fprintf(&b, "\n# Time\nUsed %d of %d minutes.\n", int(in.Elapsed.Minutes()), int(in.Duration.Minutes()))

	b.WriteString("\n# Transcript\n")
	turns := 0
	for _, m := range in.Transcript {
		if m.Role != "user" && m.Role != "assistant" {
			continue
		}
		who := "Candidate"
		if m.Role == "assistant" {
			who = "Interviewer"
		}
		fmt.Fprintf(&b, "%s: %s\n", who, truncate(strings.TrimSpace(m.Content), 2000))
		turns++
	}
	if turns == 0 {
		b.WriteString("(the candidate said nothing)\n")
	}

	code := strings.TrimSpace(in.Code)
	if code == "" {
		b.WriteString("\n# Final code\n(none written)\n")
	} else {
		fmt.Fprintf(&b, "\n# Final code (%s)\n```\n%s\n```\n", firstNonEmpty(in.Language, "python"), numberLines(truncate(code, maxCodeChars)))
	}

	if t := in.Tests; t != nil {
		fmt.Fprintf(&b, "\n# Automated tests (the candidate never saw these)\n%s: passed %d of %d.\n", t.Status, t.PassedCount, t.TotalCount)
		for _, r := range t.Results {
			if r.Passed {
				continue
			}
			line := "- failed"
			if r.Note != "" {
				line += ": " + r.Note
			}
			if r.Error != "" {
				line += " (error: " + truncate(r.Error, 200) + ")"
			}
			b.WriteString(line + "\n")
		}
	}

	messages := []Message{
		{Role: "system", Content: fmt.Sprintf(scorerPrompt, level, level)},
		{Role: "user", Content: b.String()},
	}
	raw, used, err := c.ChatJSON(ctx, model, messages, scorecardSchema())
	if err != nil {
		return nil, used, err
	}
	var card Scorecard
	if err := json.Unmarshal(raw, &card); err != nil {
		return nil, used, fmt.Errorf("the model's scorecard wasn't valid JSON: %w", err)
	}
	card.normalize()
	return &card, used, nil
}

// normalize clamps what a small model may get wrong.
func (s *Scorecard) normalize() {
	if s.Ratings == nil {
		s.Ratings = map[string]Rating{}
	}
	for _, c := range Criteria {
		r := s.Ratings[c]
		r.Score = max(1, min(4, r.Score))
		s.Ratings[c] = r
	}
	valid := false
	for _, v := range verdicts {
		if strings.EqualFold(s.Verdict, v) {
			s.Verdict, valid = v, true
		}
	}
	if !valid {
		// Derive it from the average rating.
		sum := 0
		for _, c := range Criteria {
			sum += s.Ratings[c].Score
		}
		s.Verdict = verdicts[min(len(verdicts)-1, max(0, (sum*len(verdicts))/(4*len(Criteria)+1)))]
	}
	if s.HintsGiven < 0 {
		s.HintsGiven = 0
	}
}

// ChatJSON asks model for one non-streamed reply that follows schema
// (Ollama structured outputs) and returns the raw JSON.
func (c *Client) ChatJSON(ctx context.Context, model string, messages []Message, schema map[string]any) (json.RawMessage, string, error) {
	m, err := c.resolve(model)
	if err != nil {
		return nil, "", err
	}
	body := map[string]any{
		"model":    m.Name,
		"messages": messages,
		"stream":   false,
		"format":   schema,
		"options":  map[string]any{"temperature": 0.2},
	}
	if m.Thinking {
		body["think"] = false
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, m.Name, err
	}
	ctx, cancel := context.WithTimeout(ctx, chatTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/api/chat", bytes.NewReader(payload))
	if err != nil {
		return nil, m.Name, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, m.Name, fmt.Errorf("%w: ollama request failed: %v", ErrUnavailable, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 400))
		return nil, m.Name, fmt.Errorf("ollama %s: %s", resp.Status, strings.TrimSpace(string(msg)))
	}
	var out struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
		Error string `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, m.Name, fmt.Errorf("malformed ollama response: %w", err)
	}
	if out.Error != "" {
		return nil, m.Name, errors.New(out.Error)
	}
	return json.RawMessage(strings.TrimSpace(out.Message.Content)), m.Name, nil
}
