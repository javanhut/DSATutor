package tutor

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"dsatutor/internal/chapter"
)

// The director makes Teacher Mode dependable. Whether a model actually calls
// highlight or draws in the scratchpad varies by model and by run (gemma4
// often only talks about it). So after a teacher turn, whatever is missing —
// a precise pointer at the page, a whiteboard sketch, or a checkpoint — is requested in one
// structured-output call, which every model follows, and then executed
// like ordinary tool calls.

// DirectorNeeds says which parts a turn is missing.
type DirectorNeeds struct {
	Highlight bool
	Sketch    bool
	Check     bool
}

// Direct returns tool calls that stage the teacher's reply: a highlight with
// a verbatim quote and a pointer, a whiteboard sketch, and/or a checkpoint. base is the
// conversation the teacher's turn was generated from (system prompt with
// the chapter, history, tool calls); the request reuses it word for word
// and only appends the staging instructions, so Ollama's prompt cache
// covers the chapter instead of re-reading it.
func (c *Client) Direct(ctx context.Context, model string, base []Message, ch *chapter.Chapter, ws *Workspace, reply string, need DirectorNeeds) ([]ToolCall, error) {
	if !need.Highlight && !need.Sketch && !need.Check {
		return nil, nil
	}
	props := map[string]any{}
	var required []string
	var asks []string

	if need.Highlight && ws != nil && len(ws.targets) > 0 {
		ids := make([]string, len(ws.targets))
		for i, t := range ws.targets {
			ids[i] = t.ID
		}
		props["highlight"] = map[string]any{
			"type": "object",
			"properties": map[string]any{
				"target":     map[string]any{"type": "string", "enum": ids},
				"quote":      map[string]any{"type": "string"},
				"start_line": map[string]any{"type": "integer"},
				"end_line":   map[string]any{"type": "integer"},
				"pointer":    map[string]any{"type": "string", "enum": []string{"hand", "arrow"}},
				"note":       map[string]any{"type": "string"},
			},
			"required": []string{"target", "quote", "pointer", "note"},
		}
		required = append(required, "highlight")
		asks = append(asks, `"highlight": the exact spot on the page the reply is about. Pick the target id of the region that contains it, and set quote to 3 to 15 words copied VERBATIM from the chapter material below (the sentence or code line the reply discusses), so it can be found on the page. For a code block target you may also give start_line/end_line. pointer: "hand" if the reply walks through that sentence or line, "arrow" if it just says look here. note: one short sentence shown beside the highlight.`)
	}
	if need.Sketch {
		// Content is required field by field (a row's values, a note's
		// text), so constrained decoding can't produce empty elements; the
		// parts are turned into whiteboard elements below.
		num := map[string]any{"type": "number"}
		step := map[string]any{"type": "integer", "minimum": 0, "maximum": 5}
		str := map[string]any{"type": "string", "minLength": 1}
		props["sketch"] = map[string]any{
			"type": "object",
			"properties": map[string]any{
				"title":    str,
				"captions": map[string]any{"type": "array", "items": str, "minItems": 1, "maxItems": 6},
				"rows": map[string]any{"type": "array", "minItems": 1, "maxItems": 6, "items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"step": step, "x": num, "y": num, "label": map[string]any{"type": "string"},
						"values":   map[string]any{"type": "array", "items": map[string]any{"type": "number"}, "minItems": 2, "maxItems": 10},
						"pointers": map[string]any{"type": "object", "additionalProperties": map[string]any{"type": "integer"}},
						"marks":    map[string]any{"type": "array", "items": map[string]any{"type": "integer"}},
					},
					"required": []string{"step", "x", "y", "label", "values"},
				}},
				"notes": map[string]any{"type": "array", "maxItems": 8, "items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"step": step, "x": num, "y": num, "text": str,
						"color": map[string]any{"type": "string", "enum": sketchColors},
					},
					"required": []string{"step", "x", "y", "text", "color"},
				}},
				"arrows": map[string]any{"type": "array", "maxItems": 6, "items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"step": step, "x": num, "y": num, "x2": num, "y2": num, "label": map[string]any{"type": "string"},
					},
					"required": []string{"step", "x", "y", "x2", "y2", "label"},
				}},
				"boxes": map[string]any{"type": "array", "maxItems": 6, "items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"step": step, "x": num, "y": num, "w": num, "h": num, "text": str,
					},
					"required": []string{"step", "x", "y", "w", "h", "text"},
				}},
			},
			"required": []string{"title", "captions", "rows", "notes", "arrows", "boxes"},
		}
		required = append(required, "sketch")
		asks = append(asks, `"sketch": a hand-drawn whiteboard that shows the reply's idea on a concrete small input, built up over 2 to 4 steps (element.step 0, 1, 2...; captions[i] says what step i shows in one short sentence). Coordinates run 0-100 across and 0-60 down. Draw, don't write: "rows" are arrays drawn as cells (values, a label, "pointers" like {"left": 0, "right": 4} for labelled arrows under cells, "marks" to circle cells); "boxes" hold a variable and its value ("sum = 12"); "arrows" connect things; "notes" are short handwritten labels. Rows are 8 units wide per value, so keep x + 8 * len(values) under 100; leave 20 units below a row for its pointers. Redraw the same row (same x and y) in a later step with new pointers or marks to show movement. For growth or counting ideas, show the counts as rows or boxes (e.g. a row of pair counts for n = 1..5). Keep text to short labels (six words or fewer), placed below or to the right of the drawing. Use the reply's own example numbers when it has them; when the reply has no concrete example, take one from the chapter material (its worked example or overview numbers). Always include at least one structural element (a "cells" row, "node"s, or "box"es) — never only text.`)
	}
	if need.Check {
		props["check"] = map[string]any{
			"type": "object",
			"properties": map[string]any{
				"question": map[string]any{"type": "string", "minLength": 1},
				"choices":  map[string]any{"type": "array", "items": map[string]any{"type": "string", "minLength": 1}, "minItems": 2, "maxItems": 4},
			},
			"required": []string{"question", "choices"},
		}
		required = append(required, "check")
		asks = append(asks, `"check": one prediction question about the concrete example in the reply or sketch, with 2 to 4 short answer choices. Wrong choices should be plausible misconceptions. Do not ask a generic definition question.`)
	}
	if len(required) == 0 {
		return nil, nil
	}

	var ask strings.Builder
	ask.WriteString("(Stage direction, not from the learner.) Now put your last reply on screen. Respond with JSON only, containing:\n")
	for _, a := range asks {
		ask.WriteString("- " + a + "\n")
	}
	ask.WriteString("Be faithful to your reply: illustrate exactly what it explains. Quotes must be copied verbatim from the chapter material in the system prompt.\n\nYour reply was:\n" + truncate(reply, 6000))

	messages := append(append([]Message{}, base...), Message{Role: "user", Content: ask.String()})
	if len(base) == 0 {
		// No conversation to build on (tests, odd callers): carry the
		// chapter material ourselves.
		var sys strings.Builder
		writeChapterMaterial(&sys, ch)
		if ws != nil {
			sys.WriteString(ws.Describe())
		}
		messages = []Message{{Role: "system", Content: truncate(sys.String(), 60000)}, messages[0]}
	}

	schema := map[string]any{"type": "object", "properties": props, "required": required}
	raw, _, err := c.ChatJSON(ctx, model, messages, schema)
	if err != nil {
		return nil, err
	}
	var out map[string]map[string]any
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, fmt.Errorf("director sent malformed JSON: %w", err)
	}
	var calls []ToolCall
	if h := out["highlight"]; h != nil && need.Highlight {
		calls = append(calls, newCall("highlight", h))
	}
	if sk := out["sketch"]; sk != nil && need.Sketch {
		sk = sketchFromParts(sk)
		if WeakSketch(Action{"elements": sk["elements"]}) && !strings.Contains(reply, retryMarker) {
			// One retry, asking only for the sketch, with the reason.
			if more, err := c.Direct(ctx, model, base, ch, ws, reply+retryMarker, DirectorNeeds{Sketch: true}); err == nil && len(more) > 0 {
				sk = more[0].Function.Arguments
			}
		}
		calls = append(calls, newCall("scratchpad_sketch", sk))
	}
	if check := out["check"]; check != nil && need.Check {
		calls = append(calls, newCall("ask_learner", check))
	}
	return calls, nil
}

// retryMarker tells the director its last sketch was too thin.
const retryMarker = "\n\n(Your previous sketch for this reply was nearly empty. Draw a concrete example this time: a cells row with real numbers and pointers, plus short labels.)"

func newCall(name string, args map[string]any) ToolCall {
	var c ToolCall
	c.Function.Name = name
	c.Function.Arguments = args
	return c
}

// WeakSketch reports whether a sketch action is too thin to teach from:
// fewer than three elements, or nothing but loose text and lines (no cells,
// nodes, boxes, circles or arrows).
func WeakSketch(act Action) bool {
	els, _ := act["elements"].([]any)
	structural := 0
	for _, e := range els {
		if m, ok := e.(map[string]any); ok {
			switch m["type"] {
			case "cells", "node", "box", "circle", "arrow":
				structural++
			}
		}
	}
	return len(els) < 3 || structural == 0
}

// sketchFromParts turns the director's rows/notes/arrows/boxes into the
// sketch tool's element list.
func sketchFromParts(sk map[string]any) map[string]any {
	var els []any
	add := func(kind string, items any, conv func(m map[string]any) map[string]any) {
		list, _ := items.([]any)
		for _, it := range list {
			if m, ok := it.(map[string]any); ok {
				e := conv(m)
				e["type"] = kind
				els = append(els, e)
			}
		}
	}
	add("cells", sk["rows"], func(m map[string]any) map[string]any {
		e := map[string]any{"step": m["step"], "x": m["x"], "y": m["y"], "values": m["values"], "text": m["label"]}
		if p, ok := m["pointers"].(map[string]any); ok && len(p) > 0 {
			e["pointers"] = p
		}
		if mk, ok := m["marks"].([]any); ok && len(mk) > 0 {
			e["marks"] = mk
		}
		return e
	})
	add("box", sk["boxes"], func(m map[string]any) map[string]any {
		return map[string]any{"step": m["step"], "x": m["x"], "y": m["y"], "w": m["w"], "h": m["h"], "text": m["text"]}
	})
	add("arrow", sk["arrows"], func(m map[string]any) map[string]any {
		return map[string]any{"step": m["step"], "x": m["x"], "y": m["y"], "x2": m["x2"], "y2": m["y2"], "text": m["label"]}
	})
	add("text", sk["notes"], func(m map[string]any) map[string]any {
		return map[string]any{"step": m["step"], "x": m["x"], "y": m["y"], "text": m["text"], "color": m["color"], "size": "m"}
	})
	return map[string]any{"title": sk["title"], "captions": sk["captions"], "elements": els}
}
