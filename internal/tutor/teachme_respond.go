package tutor

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// Teach Me answers whatever the learner types (a warm-up guess, a
// say-it-back after a part, an answer about a line of code, the final
// explain-it-back) by confirming or correcting each claim in it, and keeps
// asking until the answer shows real understanding.

// TeachReply is the response to something the learner typed.
type TeachReply struct {
	Verdict  string       `json:"verdict"` // solid, partial or missing
	Reply    string       `json:"reply"`
	Claims   []TeachClaim `json:"claims"`
	FollowUp string       `json:"follow_up,omitempty"`
	Missing  []string     `json:"missing"`
}

// TeachClaim is one assertion the learner made and whether it holds.
type TeachClaim struct {
	Claim   string `json:"claim"`
	Verdict string `json:"verdict"` // correct, partly or wrong
	Note    string `json:"note"`
}

// TeachRespondInput is what the learner was asked and what they typed.
type TeachRespondInput struct {
	Topic     string // the lesson topic
	Kind      string // warmup, chunk, line or explain
	Context   string // what was just taught (a part, a line of code and its explanation)
	Prompt    string // the question the learner answered
	Reference string // what a good answer contains
	Answer    string // the learner's words
	Attempt   int    // 1 for a first answer
}

const respondSystem = `You are a patient, exacting tutor of data structures and algorithms. The learner just typed a response during a lesson. Respond to what they actually said.

- "claims": split the learner's response into the separate assertions it makes (up to 5), each quoted or closely paraphrased in "claim". For each, "verdict" is correct, partly or wrong, and "note" is a full sentence that confirms it with the reason ("Yes: ... because ...") or corrects it with the true statement and the reason ("No: ... because ..."), never a bare verdict word. Never let a wrong or vague assertion pass, but never mark one wrong over wording when its substance is right: a claim that says the same thing as the true statement in other words is correct. A question the learner asks is not a claim; answer it in "reply".
- "reply": 2 to 5 sentences to the learner as "you". Confirm what they got right, then fix the most important gap or misconception, explaining the reason behind it (why the code or idea must be that way), with a tiny concrete example when it helps. If they say they are lost or ask to have it explained again, re-explain from a different angle than the lesson used (a concrete trace, an analogy, or what breaks without it), not by repeating it.
- "verdict": solid only when the response shows they understand both what happens and why; partial when the what is right but the why is missing or shaky; missing when the key idea is absent or wrong, or they only said they don't know.
- "follow_up": when the verdict is not solid, one new, specific question (ending in "?") that targets the gap, worded differently from the question they just answered, so they must think again. Empty when solid.
- "missing": up to 3 short points a complete answer would include that they left out.
Judge understanding, not wording, spelling or grammar. Plain text, no LaTeX (write O(log n), not math markup); inline code in backticks. Respond with JSON only.`

// Respond replies to the learner's typed response.
func (c *Client) Respond(ctx context.Context, model string, in TeachRespondInput) (*TeachReply, string, error) {
	answer := truncate(strings.TrimSpace(in.Answer), 3000)
	if answer == "" {
		return nil, "", fmt.Errorf("empty answer")
	}
	var user strings.Builder
	fmt.Fprintf(&user, "Topic: %s\n", truncate(in.Topic, 300))
	switch in.Kind {
	case "warmup":
		user.WriteString("This is the warm-up, asked before any teaching: be encouraging, but still confirm or correct each assertion.\n")
	case "line":
		user.WriteString("The learner is being taken through the algorithm's code one line at a time and must understand this line, and why it is written this way, before moving on.\n")
	}
	if s := strings.TrimSpace(in.Context); s != "" {
		fmt.Fprintf(&user, "What was just taught:\n%s\n", truncate(s, 2500))
	}
	if s := strings.TrimSpace(in.Prompt); s != "" {
		fmt.Fprintf(&user, "Question asked: %s\n", truncate(s, 400))
	}
	if s := strings.TrimSpace(in.Reference); s != "" {
		fmt.Fprintf(&user, "A good answer covers: %s\n", truncate(s, 800))
	}
	if in.Attempt > 1 {
		fmt.Fprintf(&user, "This is their attempt number %d at this point, so explain it a new way.\n", in.Attempt)
	}
	fmt.Fprintf(&user, "\nThe learner wrote:\n%s", answer)

	str := map[string]any{"type": "string"}
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"claims": map[string]any{"type": "array", "maxItems": 5, "items": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"claim":   map[string]any{"type": "string", "minLength": 3, "description": "the learner's assertion, quoted or closely paraphrased"},
					"verdict": map[string]any{"type": "string", "enum": []string{"correct", "partly", "wrong"}},
					"note":    map[string]any{"type": "string", "minLength": 20, "description": "a full sentence: why the claim is right, or the true statement and the reason it is"},
				},
				"required": []string{"claim", "verdict", "note"},
			}},
			"reply":     str,
			"verdict":   map[string]any{"type": "string", "enum": []string{"solid", "partial", "missing"}},
			"follow_up": str,
			"missing":   map[string]any{"type": "array", "items": str, "maxItems": 3},
		},
		"required": []string{"claims", "reply", "verdict", "follow_up", "missing"},
	}
	raw, name, err := c.ChatJSON(ctx, model, []Message{{Role: "system", Content: respondSystem}, {Role: "user", Content: user.String()}}, schema)
	if err != nil {
		return nil, name, err
	}
	var r TeachReply
	if err := json.Unmarshal(raw, &r); err != nil {
		return nil, name, fmt.Errorf("reply was malformed JSON: %w", err)
	}
	r.clean()
	return &r, name, nil
}

func (r *TeachReply) clean() {
	t := func(s string, n int) string { return truncate(strings.TrimSpace(s), n) }
	switch r.Verdict {
	case "solid", "partial", "missing":
	default:
		r.Verdict = "partial"
	}
	r.Reply = t(r.Reply, 1500)
	r.FollowUp = t(r.FollowUp, 400)
	if r.Verdict == "solid" {
		r.FollowUp = ""
	}
	r.Missing = cleanList(r.Missing, 3, 200)
	var claims []TeachClaim
	for _, c := range r.Claims {
		c.Claim, c.Note = t(c.Claim, 300), t(c.Note, 500)
		if len(strings.Fields(c.Note)) < 3 {
			c.Note = "" // a bare "wrong" says nothing the mark doesn't
		}
		switch c.Verdict {
		case "correct", "partly", "wrong":
		default:
			c.Verdict = "partly"
		}
		if c.Claim != "" && len(claims) < 5 {
			claims = append(claims, c)
		}
	}
	r.Claims = claims
}

// TeachWalkStep is one group of lines in a line-by-line walkthrough.
type TeachWalkStep struct {
	From     int    `json:"from"` // first line, 1-based
	To       int    `json:"to"`   // last line, inclusive
	Code     string `json:"code"`
	What     string `json:"what"`
	Why      string `json:"why"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

const walkSystem = `You teach an algorithm by walking a learner through its code one line at a time. You get the code with numbered lines. Cover every line in order, in steps. A step is usually a single line; group consecutive lines only when they make no sense apart (a def and its docstring, a two-line swap). The def line and plain return lines get their own short steps too.

For each step:
- "from", "to": its first and last line numbers.
- "what": what the line does, concretely, in 1 to 3 sentences, naming the variables.
- "why": the rationale: why it is written exactly this way, what invariant it keeps or what would break if it were different (an off-by-one, < vs <=, the order of two lines). 1 to 3 sentences. This is the heart of the lesson.
- "question": one question (ending in "?") that checks the learner understood the why, not just the what: predict a value, say what breaks if it changed, explain a choice. Not answerable by copying "what".
- "answer": the correct answer to the question, stated plainly (not another question), with the reason.
Plain text; inline code in backticks. Respond with JSON only: {"steps": [{"from", "to", "what", "why", "question", "answer"}]}`

// WalkCode builds a line-by-line walkthrough of code. Models are poor at
// quoting their own code, so they only give line numbers; the server cuts
// the lines and repairs gaps and overlaps so every line is covered once.
func (c *Client) WalkCode(ctx context.Context, model, topic, code string) ([]TeachWalkStep, string, error) {
	code = stripFences(truncate(strings.TrimSpace(code), 2500))
	lines := strings.Split(code, "\n")
	if strings.TrimSpace(code) == "" || len(lines) < 2 {
		return nil, "", fmt.Errorf("no code to walk through")
	}
	var numbered strings.Builder
	for i, line := range lines {
		fmt.Fprintf(&numbered, "%d: %s\n", i+1, line)
	}
	str := map[string]any{"type": "string", "minLength": 1}
	line := map[string]any{"type": "integer", "minimum": 1, "maximum": len(lines)}
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"steps": map[string]any{"type": "array", "minItems": 1, "maxItems": 30, "items": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"from": line, "to": line, "what": str, "why": str, "question": str, "answer": str,
				},
				"required": []string{"from", "to", "what", "why", "question", "answer"},
			}},
		},
		"required": []string{"steps"},
	}
	raw, name, err := c.ChatJSON(ctx, model, []Message{
		{Role: "system", Content: walkSystem},
		{Role: "user", Content: fmt.Sprintf("Topic: %s\n\n%s", truncate(topic, 300), numbered.String())},
	}, schema)
	if err != nil {
		return nil, name, err
	}
	var out struct {
		Steps []TeachWalkStep `json:"steps"`
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, name, fmt.Errorf("walkthrough was malformed JSON: %w", err)
	}
	steps := fitWalk(out.Steps, lines)
	if len(steps) == 0 {
		return nil, name, fmt.Errorf("the model returned an empty walkthrough")
	}
	return steps, name, nil
}

// fitWalk orders the steps, clamps their ranges, drops ones that add
// nothing, closes gaps by widening the step before, and fills each step's
// code from lines. Lines that are blank are never a step of their own.
func fitWalk(in []TeachWalkStep, lines []string) []TeachWalkStep {
	t := func(s string, n int) string { return truncate(strings.TrimSpace(s), n) }
	n := len(lines)
	var steps []TeachWalkStep
	for _, s := range in {
		if s.From < 1 {
			s.From = 1
		}
		if s.To < s.From {
			s.To = s.From
		}
		if s.From > n {
			continue
		}
		if s.To > n {
			s.To = n
		}
		s.What, s.Why = t(s.What, 600), t(s.Why, 600)
		s.Question, s.Answer = t(s.Question, 400), t(s.Answer, 600)
		if s.What == "" && s.Why == "" {
			continue
		}
		steps = append(steps, s)
	}
	sort.SliceStable(steps, func(i, j int) bool { return steps[i].From < steps[j].From })
	var out []TeachWalkStep
	next := 1 // first line not yet covered
	for _, s := range steps {
		if s.To < next {
			continue // wholly inside an earlier step
		}
		if s.From > next {
			if len(out) > 0 {
				out[len(out)-1].To = s.From - 1 // a gap: the step before takes it
			} else {
				s.From = 1
			}
		} else {
			s.From = next
		}
		out = append(out, s)
		next = s.To + 1
	}
	if len(out) > 0 && next <= n {
		out[len(out)-1].To = n
	}
	for i := range out {
		s := &out[i]
		// Trim blank lines off the edges of the step.
		for s.From < s.To && strings.TrimSpace(lines[s.From-1]) == "" {
			s.From++
		}
		for s.To > s.From && strings.TrimSpace(lines[s.To-1]) == "" {
			s.To--
		}
		s.Code = strings.Join(lines[s.From-1:s.To], "\n")
	}
	return out
}
