package tutor

import (
	"encoding/json"
	"fmt"
	"strings"
)

// The tutor can act on the page through tools: point at a region (an
// example, a lesson section, lines of the learner's code), edit the code
// editor, switch a problem's tab and pick a test case. The browser sends
// the regions on screen as Targets; tool schemas only accept those ids, so
// the model can't invent places to point at. Tools run here, on the server:
// each call becomes an Action the browser carries out, and the model gets a
// result describing what changed. Code edits are applied to the server's
// copy of the code too, so later steps see the code as it now is.

// Target is one region of the page the tutor may point at.
type Target struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Lines    bool   `json:"lines,omitempty"`    // a code editor: line ranges can be highlighted
	Writable bool   `json:"writable,omitempty"` // a code editor the tutor may edit
}

// Tool is an Ollama tool definition.
type Tool struct {
	Type     string       `json:"type"`
	Function ToolFunction `json:"function"`
}

type ToolFunction struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters"`
}

// ToolCall is a call the model made.
type ToolCall struct {
	Function struct {
		Name      string         `json:"name"`
		Arguments map[string]any `json:"arguments"`
	} `json:"function"`
}

// Action is what the browser does for one tool call; it is streamed as
// {"type":"action", ...}.
type Action map[string]any

const (
	maxTargets = 60
	// More than one board in a teaching beat produces duplicate history
	// items and makes the learner decide which version to follow.
	maxDrawingsPerTurn = 1
)

// Workspace is the page as one request sees it: the regions and the
// learner's code, updated as the tutor edits it.
type Workspace struct {
	targets   []Target
	byID      map[string]Target
	code      string
	codeID    string // the target holding Code
	tests     int    // test cases on screen; 0 when no problem is open
	isProblem bool
	interview bool    // an interviewer only points; it never edits or switches tabs
	teach     bool    // Teacher Mode: writes its own lesson articles on the pad
	lesson    *Lesson // Teacher Mode's plan, updated as articles are written
	drawn     int     // scratchpad drawings this turn; capped so a model can't spam the board
}

// NewWorkspace validates the browser's targets.
func NewWorkspace(ctx Context) *Workspace {
	w := &Workspace{byID: map[string]Target{}, code: ctx.Code, tests: ctx.TestCount, isProblem: ctx.View == "problem", interview: ctx.Interview != nil}
	w.teach = ctx.Teach && !w.interview
	w.lesson = ctx.Lesson
	for _, t := range ctx.Targets {
		t.ID, t.Label = strings.TrimSpace(t.ID), strings.TrimSpace(t.Label)
		if t.ID == "" || len(t.ID) > 64 || w.byID[t.ID].ID != "" || len(w.targets) >= maxTargets {
			continue
		}
		if len(t.Label) > 160 {
			t.Label = t.Label[:160]
		}
		w.targets = append(w.targets, t)
		w.byID[t.ID] = t
		if t.Writable && w.codeID == "" {
			w.codeID = t.ID
		}
	}
	return w
}

// Describe lists the targets for the system prompt.
func (w *Workspace) Describe() string {
	if len(w.targets) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("\n## Parts of the page you can point at (tool target ids)\n")
	for _, t := range w.targets {
		extra := ""
		switch {
		case t.Writable:
			extra = " (the learner's code editor: highlight line ranges, and edit it with write_code)"
		case t.Lines:
			extra = " (line ranges can be highlighted)"
		}
		fmt.Fprintf(&b, "- `%s`: %s%s\n", t.ID, t.Label, extra)
	}
	return b.String()
}

// Tools returns the tool definitions for this page.
func (w *Workspace) Tools() []Tool {
	var tools []Tool
	if !w.interview {
		tools = append(tools, scratchpadTools()...)
	}
	if w.teach {
		n := 0
		if w.lesson.Valid() {
			n = len(w.lesson.Sections)
		}
		tools = append(tools, articleTool(n))
	}
	if len(w.targets) == 0 {
		return tools
	}
	ids := make([]string, len(w.targets))
	for i, t := range w.targets {
		ids[i] = t.ID
	}
	tools = append(tools, fn("highlight",
		"Point at the precise words that directly support the claim you are teaching now: it scrolls there, highlights them, and a hand or arrow points at them, with your note beside it. Give quote (the exact words on the page) or start_line/end_line (for code), not just a section. Never jump to a merely related sentence or a later common mistake because it shares a keyword.",
		map[string]any{
			"target":     map[string]any{"type": "string", "enum": ids, "description": "id of the region that contains the spot"},
			"quote":      map[string]any{"type": "string", "description": "the exact words from the page to highlight, copied verbatim (3 to 15 words), e.g. a sentence of the lesson or a line of an example"},
			"start_line": map[string]any{"type": "integer", "description": "first line to highlight (code editors and code blocks)"},
			"end_line":   map[string]any{"type": "integer", "description": "last line to highlight (defaults to start_line)"},
			"pointer":    map[string]any{"type": "string", "enum": []string{"hand", "arrow", "none"}, "description": "hand: you're reading or walking through this sentence or line with the learner; arrow: 'look here' at a spot; none: just highlight"},
			"note":       map[string]any{"type": "string", "description": "a short note shown beside the highlight (one sentence)"},
		}, "target"))

	if w.interview {
		return tools
	}
	if w.codeID != "" {
		tools = append(tools, fn("write_code",
			"Edit the learner's code editor. Only use it when the learner asks you to write, fix, or scaffold code, or agrees to it. Prefer small edits. The learner can undo it.",
			map[string]any{
				"mode": map[string]any{"type": "string", "enum": []string{"replace_lines", "insert_after_line", "replace_all"},
					"description": "replace_lines replaces start_line..end_line; insert_after_line inserts after start_line (0 inserts at the top); replace_all replaces the whole editor"},
				"start_line": map[string]any{"type": "integer"},
				"end_line":   map[string]any{"type": "integer"},
				"code":       map[string]any{"type": "string", "description": "the new code, without line-number prefixes, indented to fit"},
				"summary":    map[string]any{"type": "string", "description": "a few words on what changed, e.g. 'check the complement before storing'"},
			}, "mode", "code"))
	}
	if w.isProblem {
		tools = append(tools, fn("open_tab",
			"Switch the problem panel to another tab. The solution tab is not available to you.",
			map[string]any{"tab": map[string]any{"type": "string", "enum": []string{"description", "hints", "visualization"}}}, "tab"))
		if w.tests > 0 {
			tools = append(tools, fn("select_test",
				"Select one of the visible test cases so the learner can run their code on it.",
				map[string]any{"case": map[string]any{"type": "integer", "description": fmt.Sprintf("1 to %d", w.tests)}}, "case"))
		}
	}
	return tools
}

func fn(name, desc string, props map[string]any, required ...string) Tool {
	return Tool{Type: "function", Function: ToolFunction{
		Name:        name,
		Description: desc,
		Parameters:  map[string]any{"type": "object", "properties": props, "required": required},
	}}
}

// Execute runs one tool call. It returns the result for the model, and the
// action for the browser (nil when the call was rejected).
func (w *Workspace) Execute(call ToolCall) (string, Action) {
	args := call.Function.Arguments
	switch call.Function.Name {
	case "highlight":
		t, ok := w.byID[str(args, "target")]
		if !ok {
			return "Error: unknown target. Use one of the listed ids.", nil
		}
		act := Action{"name": "highlight", "target": t.ID}
		if note := strings.TrimSpace(str(args, "note")); note != "" {
			act["note"] = truncate(note, 300)
		}
		if q := strings.TrimSpace(str(args, "quote")); q != "" {
			act["quote"] = truncate(q, 300)
		}
		switch p := str(args, "pointer"); p {
		case "hand", "arrow", "none":
			act["pointer"] = p
		}
		where := t.Label
		if t.Lines && t.ID != w.codeID {
			// A code block in a lesson: the server doesn't hold its text, so
			// the page clamps the range to the block.
			if start := intArg(args, "start_line"); start > 0 {
				end := max(start, intArg(args, "end_line"))
				act["start_line"], act["end_line"] = start, end
				where = fmt.Sprintf("lines %d-%d of %s", start, end, t.Label)
			}
		} else if t.Lines {
			n := lineCount(w.codeFor(t))
			start, end := intArg(args, "start_line"), intArg(args, "end_line")
			if start > 0 {
				if end < start {
					end = start
				}
				if start > n {
					return fmt.Sprintf("Error: the code has only %d lines.", n), nil
				}
				end = min(end, n)
				act["start_line"], act["end_line"] = start, end
				where = fmt.Sprintf("lines %d-%d of %s", start, end, t.Label)
			}
		}
		if q, ok := act["quote"].(string); ok {
			where = fmt.Sprintf("%q in %s", q, t.Label)
		}
		return "Highlighted " + where + "; the learner can see it now.", act

	case "write_code":
		if w.codeID == "" || w.interview {
			return "Error: there is no editable code on this page.", nil
		}
		code, err := applyEdit(w.code, str(args, "mode"), intArg(args, "start_line"), intArg(args, "end_line"), str(args, "code"))
		if err != nil {
			return "Error: " + err.Error(), nil
		}
		w.code = code
		act := Action{"name": "write_code", "target": w.codeID, "code": code}
		if s := strings.TrimSpace(str(args, "summary")); s != "" {
			act["summary"] = truncate(s, 200)
		}
		return fmt.Sprintf("Done; the learner can undo it. The editor now holds:\n```\n%s\n```", numberLines(code)), act

	case "scratchpad_run_code", "scratchpad_visualize", "scratchpad_sketch":
		if w.drawn >= maxDrawingsPerTurn {
			return fmt.Sprintf("You have already drawn %d times this turn; the scratchpad keeps them. Don't draw again now: continue in words.", w.drawn), nil
		}
		res, act := executeScratchpad(call)
		if act != nil {
			w.drawn++
		}
		return res, act

	case "ask_learner":
		return executeScratchpad(call)

	case "write_lesson":
		if !w.teach {
			return "Error: lesson articles are only for Teacher Mode.", nil
		}
		res, act := executeArticle(args, w.lesson)
		if act != nil {
			// The section is on the page now: the teacher may point at it
			// for the rest of this turn.
			id := fmt.Sprintf("lesson-%d", act["section"])
			if _, ok := w.byID[id]; !ok && len(w.targets) < maxTargets {
				t := Target{ID: id, Label: fmt.Sprintf("Lesson page, section %d: %s", act["section"], act["heading"])}
				w.targets = append(w.targets, t)
				w.byID[id] = t
			}
		}
		return res, act

	case "open_tab":
		tab := str(args, "tab")
		if !w.isProblem || w.interview || (tab != "description" && tab != "hints" && tab != "visualization") {
			return "Error: tab must be description, hints or visualization.", nil
		}
		return "Opened the " + tab + " tab.", Action{"name": "open_tab", "tab": tab}

	case "select_test":
		n := intArg(args, "case")
		if !w.isProblem || w.interview || n < 1 || n > w.tests {
			return fmt.Sprintf("Error: case must be between 1 and %d.", w.tests), nil
		}
		return fmt.Sprintf("Selected test case %d.", n), Action{"name": "select_test", "case": n}
	}
	return "Error: unknown tool " + call.Function.Name, nil
}

func (w *Workspace) codeFor(t Target) string {
	if t.ID == w.codeID {
		return w.code
	}
	return ""
}

// applyEdit edits code by 1-based line numbers.
func applyEdit(code, mode string, start, end int, text string) (string, error) {
	text = strings.TrimRight(stripLineNumbers(text), "\n")
	if mode == "replace_all" {
		return text, nil
	}
	var lines []string
	if code != "" {
		lines = strings.Split(code, "\n")
	}
	newLines := strings.Split(text, "\n")
	switch mode {
	case "replace_lines":
		if end < start {
			end = start
		}
		if start < 1 || start > len(lines) {
			return "", fmt.Errorf("start_line must be between 1 and %d", len(lines))
		}
		end = min(end, len(lines))
		out := append(append(append([]string{}, lines[:start-1]...), newLines...), lines[end:]...)
		return strings.Join(out, "\n"), nil
	case "insert_after_line":
		if start < 0 || start > len(lines) {
			return "", fmt.Errorf("start_line must be between 0 and %d", len(lines))
		}
		out := append(append(append([]string{}, lines[:start]...), newLines...), lines[start:]...)
		return strings.Join(out, "\n"), nil
	}
	return "", fmt.Errorf("mode must be replace_lines, insert_after_line or replace_all")
}

// stripLineNumbers undoes the "  4 | " prefixes the prompt shows, in case
// the model copies them into code it writes.
func stripLineNumbers(text string) string {
	lines := strings.Split(text, "\n")
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		if !lineNumbered(l) {
			return text
		}
	}
	for i, l := range lines {
		if lineNumbered(l) {
			rest := l[strings.Index(l, "|")+1:]
			lines[i] = strings.TrimPrefix(rest, " ")
		}
	}
	return strings.Join(lines, "\n")
}

func lineNumbered(l string) bool {
	i := strings.Index(l, "|")
	if i <= 0 {
		return false
	}
	num := strings.TrimSpace(l[:i])
	if num == "" {
		return false
	}
	for _, r := range num {
		if r < '0' || r > '9' {
			return false
		}
	}
	return len(l) == i+1 || l[i+1] == ' '
}

func lineCount(code string) int {
	if code == "" {
		return 0
	}
	return strings.Count(code, "\n") + 1
}

func str(args map[string]any, key string) string {
	s, _ := args[key].(string)
	return s
}

// intArg accepts numbers and numeric strings: small models send both.
func intArg(args map[string]any, key string) int {
	switch v := args[key].(type) {
	case float64:
		return int(v)
	case int:
		return v
	case json.Number:
		n, _ := v.Int64()
		return int(n)
	case string:
		var n int
		fmt.Sscanf(strings.TrimSpace(v), "%d", &n)
		return n
	}
	return 0
}

// listArg reads a list argument. Small models sometimes send a list as a
// JSON-encoded string ("[\"a\", \"b\"]"), so that is accepted too.
func listArg(args map[string]any, key string) []any {
	switch v := args[key].(type) {
	case []any:
		return v
	case string:
		var out []any
		if json.Unmarshal([]byte(strings.TrimSpace(v)), &out) == nil {
			return out
		}
	}
	return nil
}

// ResetDrawings lifts the per-turn drawing cap, for the director's sketch
// after the model's own attempts.
func (w *Workspace) ResetDrawings() { w.drawn = 0 }
