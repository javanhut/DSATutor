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
// a precise pointer at the page, a whiteboard sketch — is requested in one
// structured-output call, which every model follows, and then executed
// like ordinary tool calls.

// DirectorNeeds says which parts a turn is missing.
type DirectorNeeds struct {
	Highlight bool
	Sketch    bool
}

// Direct returns tool calls that stage the teacher's reply: a highlight with
// a verbatim quote and a pointer, and/or a whiteboard sketch.
func (c *Client) Direct(ctx context.Context, model string, tctx Context, ch *chapter.Chapter, ws *Workspace, reply string, need DirectorNeeds) ([]ToolCall, error) {
	if !need.Highlight && !need.Sketch {
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
		props["sketch"] = map[string]any{
			"type": "object",
			"properties": map[string]any{
				"title":    map[string]any{"type": "string"},
				"captions": map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
				"elements": map[string]any{"type": "array", "items": sketchElementSchema(), "minItems": 4},
			},
			"required": []string{"title", "captions", "elements"},
		}
		required = append(required, "sketch")
		asks = append(asks, `"sketch": a hand-drawn whiteboard that shows the reply's idea on a concrete small input, built up over 2 to 4 steps (element.step 0, 1, 2...; captions[i] says what step i shows in one short sentence). Coordinates run 0-100 across and 0-60 down. Draw, don't write: a row of "cells" for an array (values, an optional text label, "pointers" like {"left": 0, "right": 4} to put labelled arrows under cells, "marks" to circle cells), "node" circles and "arrow"/"line" for trees and graphs, "box" for variables. Redraw the same cells row (same x and y) in a later step with new pointers or marks to show movement. Keep text to short labels (six words or fewer), placed below or to the right of the drawing. Use the reply's own example numbers when it has them; when the reply has no concrete example, take one from the chapter material (its worked example or overview numbers). Always include at least one structural element (a "cells" row, "node"s, or "box"es) — never only text.`)
	}
	if len(required) == 0 {
		return nil, nil
	}

	var sys strings.Builder
	sys.WriteString("You stage a teacher's lesson on screen. The teacher just said the reply below to a learner, while teaching the chapter below. Produce JSON with:\n")
	for _, a := range asks {
		sys.WriteString("- " + a + "\n")
	}
	sys.WriteString("\nBe faithful to the reply: illustrate exactly what it explains, nothing else.\n\n# Chapter material\n")
	writeChapterMaterial(&sys, ch)
	if ws != nil {
		sys.WriteString(ws.Describe())
	}

	schema := map[string]any{"type": "object", "properties": props, "required": required}
	raw, _, err := c.ChatJSON(ctx, model, []Message{
		{Role: "system", Content: truncate(sys.String(), 60000)},
		{Role: "user", Content: "The teacher's reply:\n\n" + truncate(reply, 6000)},
	}, schema)
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
		if WeakSketch(Action{"elements": sk["elements"]}) && !strings.Contains(reply, retryMarker) {
			// One retry, asking only for the sketch, with the reason.
			if more, err := c.Direct(ctx, model, tctx, ch, ws, reply+retryMarker, DirectorNeeds{Sketch: true}); err == nil && len(more) > 0 {
				sk = more[0].Function.Arguments
			}
		}
		calls = append(calls, newCall("scratchpad_sketch", sk))
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
