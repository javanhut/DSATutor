package tutor

import (
	"strings"
	"testing"

	"dsatutor/internal/sandbox"
)

func call(name string, args map[string]any) ToolCall {
	var c ToolCall
	c.Function.Name = name
	c.Function.Arguments = args
	return c
}

func problemWorkspace() *Workspace {
	return NewWorkspace(Context{
		View:      "problem",
		Code:      "a\nb\nc",
		TestCount: 3,
		Targets: []Target{
			{ID: "code", Label: "editor", Lines: true, Writable: true},
			{ID: "example-1", Label: "Example 1"},
			{ID: "example-1", Label: "duplicate is dropped"},
			{ID: "", Label: "no id is dropped"},
		},
	})
}

func TestApplyEdit(t *testing.T) {
	cases := []struct {
		mode       string
		start, end int
		text, want string
	}{
		{"replace_lines", 2, 2, "B", "a\nB\nc"},
		{"replace_lines", 2, 9, "X\nY", "a\nX\nY"}, // end clamps to the last line
		{"replace_lines", 1, 0, "Z", "Z\nb\nc"},    // end < start means one line
		{"insert_after_line", 0, 0, "top", "top\na\nb\nc"},
		{"insert_after_line", 3, 0, "end\n", "a\nb\nc\nend"},
		{"replace_all", 0, 0, "new", "new"},
	}
	for _, c := range cases {
		got, err := applyEdit("a\nb\nc", c.mode, c.start, c.end, c.text)
		if err != nil || got != c.want {
			t.Errorf("%s %d-%d %q: got %q, %v; want %q", c.mode, c.start, c.end, c.text, got, err, c.want)
		}
	}
	for _, bad := range []struct {
		mode  string
		start int
	}{{"replace_lines", 0}, {"replace_lines", 4}, {"insert_after_line", 4}, {"delete", 1}} {
		if _, err := applyEdit("a\nb\nc", bad.mode, bad.start, bad.start, "x"); err == nil {
			t.Errorf("%s at %d: want an error", bad.mode, bad.start)
		}
	}
}

func TestStripLineNumbers(t *testing.T) {
	numbered := "  1 | def f(a, b):\n  2 |\n  3 |     return a | b"
	if got := stripLineNumbers(numbered); got != "def f(a, b):\n\n    return a | b" {
		t.Errorf("got %q", got)
	}
	plain := "x = a | b\ny = 2"
	if got := stripLineNumbers(plain); got != plain {
		t.Errorf("plain code changed: %q", got)
	}
}

func TestWorkspaceToolsAndTargets(t *testing.T) {
	ws := problemWorkspace()
	if len(ws.targets) != 2 {
		t.Fatalf("targets = %+v, want code and example-1", ws.targets)
	}
	names := map[string]bool{}
	for _, tool := range ws.Tools() {
		names[tool.Function.Name] = true
	}
	for _, want := range []string{"highlight", "write_code", "open_tab", "select_test"} {
		if !names[want] {
			t.Errorf("missing tool %s", want)
		}
	}
	if !strings.Contains(ws.Describe(), "`example-1`: Example 1") {
		t.Errorf("describe = %s", ws.Describe())
	}

	// A chapter page: the scratchpad and highlight, nothing to edit or switch.
	ch := NewWorkspace(Context{View: "chapter", Targets: []Target{{ID: "section-1", Label: "Overview"}}})
	got := map[string]bool{}
	for _, tool := range ch.Tools() {
		got[tool.Function.Name] = true
	}
	for _, want := range []string{"highlight", "scratchpad_run_code", "scratchpad_visualize", "ask_learner"} {
		if !got[want] {
			t.Errorf("chapter page lacks %s", want)
		}
	}
	if got["write_code"] || got["open_tab"] {
		t.Errorf("chapter page got problem tools: %v", got)
	}
	// With nothing to point at, only the scratchpad remains.
	if tools := NewWorkspace(Context{View: "home"}).Tools(); len(tools) != 4 {
		t.Errorf("home page tools = %d, want the 4 scratchpad tools", len(tools))
	}
}

func TestExecute(t *testing.T) {
	ws := problemWorkspace()

	res, act := ws.Execute(call("highlight", map[string]any{"target": "code", "start_line": 2.0, "end_line": "9", "note": "here"}))
	if act == nil || act["start_line"] != 2 || act["end_line"] != 3 || act["note"] != "here" {
		t.Errorf("highlight: %s %v", res, act)
	}
	if _, act := ws.Execute(call("highlight", map[string]any{"target": "nowhere"})); act != nil {
		t.Error("unknown target accepted")
	}
	if _, act := ws.Execute(call("highlight", map[string]any{"target": "code", "start_line": 7.0})); act != nil {
		t.Error("line past the end accepted")
	}

	res, act = ws.Execute(call("write_code", map[string]any{"mode": "replace_lines", "start_line": 2.0, "code": "  2 | B"}))
	if act == nil || act["code"] != "a\nB\nc" || !strings.Contains(res, "  2 | B") {
		t.Errorf("write_code: %s %v", res, act)
	}
	// Later calls see the edited code.
	_, act = ws.Execute(call("write_code", map[string]any{"mode": "insert_after_line", "start_line": 3.0, "code": "d"}))
	if act == nil || act["code"] != "a\nB\nc\nd" {
		t.Errorf("second edit: %v", act)
	}

	if _, act := ws.Execute(call("open_tab", map[string]any{"tab": "solution"})); act != nil {
		t.Error("solution tab opened")
	}
	if _, act := ws.Execute(call("select_test", map[string]any{"case": 4.0})); act != nil {
		t.Error("test case 4 of 3 accepted")
	}
	if _, act := ws.Execute(call("select_test", map[string]any{"case": 2.0})); act == nil || act["case"] != 2 {
		t.Errorf("select_test: %v", act)
	}
}

func TestScratchpadTools(t *testing.T) {
	ws := NewWorkspace(Context{View: "chapter"})

	res, act := ws.Execute(call("scratchpad_visualize", map[string]any{
		"title": "window",
		"frames": []any{map[string]any{"caption": "start", "structures": []any{
			map[string]any{"name": "nums", "type": "array", "data": []any{1.0, 2.0}},
		}}},
	}))
	if act == nil || act["name"] != "scratchpad_visualize" || !strings.Contains(res, "1 frames") {
		t.Errorf("visualize: %s %v", res, act)
	}
	if _, act := ws.Execute(call("scratchpad_visualize", map[string]any{"title": "x", "frames": []any{
		map[string]any{"caption": "c", "structures": []any{map[string]any{"name": "n", "type": "hologram", "data": 1.0}}},
	}})); act != nil {
		t.Error("unknown structure type accepted")
	}

	res, act = ws.Execute(call("scratchpad_sketch", map[string]any{
		"title":    "window",
		"captions": []any{"start", "slide"},
		"elements": []any{
			map[string]any{"type": "cells", "x": 10.0, "y": 10.0, "values": []any{2.0, 1.0, 5.0}},
			map[string]any{"type": "arrow", "step": 1.0, "x": 10.0, "y": 30.0, "x2": 20.0, "y2": 15.0, "text": "left"},
		},
	}))
	if act == nil || act["name"] != "scratchpad_sketch" || !strings.Contains(res, "2 steps") {
		t.Errorf("sketch: %s %v", res, act)
	}
	if _, act := ws.Execute(call("scratchpad_sketch", map[string]any{"title": "x", "elements": []any{map[string]any{"type": "laser"}}})); act != nil {
		t.Error("unknown sketch element accepted")
	}

	_, act = ws.Execute(call("ask_learner", map[string]any{"question": "Which pointer moves?", "choices": []any{"left", "right", ""}}))
	if act == nil || len(act["choices"].([]string)) != 2 {
		t.Errorf("ask_learner: %v", act)
	}
	// Lists sent as JSON strings, as some models do, still work.
	_, act = ws.Execute(call("ask_learner", map[string]any{"question": "Which?", "choices": `["a", "b"]`}))
	if act == nil || len(act["choices"].([]string)) != 2 {
		t.Errorf("stringified choices: %v", act)
	}
	if _, act := ws.Execute(call("ask_learner", map[string]any{"question": "Q", "choices": []any{"only one"}})); act != nil {
		t.Error("a question with one choice accepted")
	}

	// An interviewer has none of these.
	iv := NewWorkspace(Context{View: "problem", Interview: &Interview{}})
	if len(iv.Tools()) != 0 {
		t.Errorf("interviewer without targets got tools: %v", iv.Tools())
	}
}

func TestScratchpadRunsCode(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python")
	}
	ws := NewWorkspace(Context{View: "chapter"})
	res, act := ws.Execute(call("scratchpad_run_code", map[string]any{"code": "nums = [3, 1, 2]\nnums.sort()\nprint(nums)"}))
	if act == nil || !strings.Contains(res, "[1, 2, 3]") || len(act["steps"].([]sandbox.Step)) == 0 {
		t.Errorf("run: %s %v", res, act)
	}
}

func TestWeakSketch(t *testing.T) {
	el := func(typ string) any { return map[string]any{"type": typ} }
	cases := []struct {
		els  []any
		weak bool
	}{
		{[]any{el("line")}, true},
		{[]any{el("text"), el("text"), el("line")}, true},
		{[]any{el("cells"), el("text")}, true},
		{[]any{el("cells"), el("text"), el("arrow")}, false},
	}
	for i, c := range cases {
		if got := WeakSketch(Action{"elements": c.els}); got != c.weak {
			t.Errorf("case %d: weak = %v, want %v", i, got, c.weak)
		}
	}
}

// ornith-1.5 sent text elements without text and cells without values,
// which drew a blank board; those are dropped, and an all-empty sketch is
// rejected with a reason so the model can redraw.
func TestSketchContentAndDrawingCap(t *testing.T) {
	ws := NewWorkspace(Context{View: "chapter"})
	empty := map[string]any{"title": "t", "elements": []any{
		map[string]any{"type": "text", "x": 5.0, "y": 5.0, "text": nil},
		map[string]any{"type": "cells", "x": 5.0, "y": 20.0, "values": nil},
	}}
	if res, act := ws.Execute(call("scratchpad_sketch", empty)); act != nil || !strings.Contains(res, "blank") {
		t.Errorf("empty sketch: %s %v", res, act)
	}
	mixed := map[string]any{"title": "t", "elements": []any{
		map[string]any{"type": "text", "x": 5.0, "y": 5.0},
		map[string]any{"type": "cells", "x": 5.0, "y": 20.0, "values": []any{1.0, 2.0}},
	}}
	res, act := ws.Execute(call("scratchpad_sketch", mixed))
	if act == nil || len(act["elements"].([]any)) != 1 || !strings.Contains(res, "left out") {
		t.Errorf("mixed sketch: %s %v", res, act)
	}
	ws.Execute(call("scratchpad_sketch", mixed))
	if res, act := ws.Execute(call("scratchpad_sketch", mixed)); act != nil || !strings.Contains(res, "already drawn") {
		t.Errorf("third drawing allowed: %s", res)
	}
	ws.ResetDrawings()
	if _, act := ws.Execute(call("scratchpad_sketch", mixed)); act == nil {
		t.Error("drawing blocked after ResetDrawings")
	}
}
