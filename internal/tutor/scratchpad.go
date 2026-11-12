package tutor

import (
	"encoding/json"
	"fmt"
	"strings"

	"dsatutor/internal/sandbox"
)

// The teacher's scratchpad is a pop-up on the page where the tutor writes
// and runs code and draws data-structure diagrams, and where the learner can
// edit and rerun what it wrote. scratchpad_run_code runs Python through the
// same tracer as the sandbox, so its step-by-step visualization shows what
// really happened. scratchpad_visualize draws frames the model describes,
// in the format the tracer produces (the page has one renderer for both).
// ask_learner puts clickable answers under a question.

// Scratchpad is what the learner's scratchpad holds, sent with each message.
type Scratchpad struct {
	Code   string `json:"code,omitempty"`
	Output string `json:"output,omitempty"`
}

const (
	maxScratchSteps  = 400
	maxFrames        = 30
	maxFrameBytes    = 60000
	scratchTimeoutMs = 5000
)

var (
	sketchTypes  = []string{"text", "box", "circle", "arrow", "line", "cells", "node", "underline"}
	sketchColors = []string{"ink", "red", "blue", "green", "orange"}
)

const maxSketchElements = 150

var structureTypes = []string{"value", "array", "bars", "matrix", "linked_list", "tree", "hash_table", "set", "graph"}

func scratchpadTools() []Tool {
	structure := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"name":       map[string]any{"type": "string", "description": "label, e.g. 'nums' or 'seen'"},
			"type":       map[string]any{"type": "string", "enum": structureTypes},
			"data":       map[string]any{"description": "value: a single number or string (e.g. a running sum); array/bars/set: a list of values; matrix: a list of rows; linked_list: {\"nodes\": [values]}; tree: nested {\"val\", \"left\", \"right\"} (null for none); hash_table: an object of key -> value; graph: {\"nodes\": [ids], \"edges\": [[from, to], ...]}"},
			"highlights": map[string]any{"type": "object", "description": "optional. array/bars: {\"pointers\": {\"left\": 0, \"right\": 4}, \"window\": {\"start\": 1, \"end\": 3}}; matrix: {\"cell\": [[row, col]]}"},
		},
		"required": []string{"name", "type", "data"},
	}
	return []Tool{
		fn("scratchpad_run_code",
			"Write a short Python program in the teacher's scratchpad and run it. The learner sees the code, its output and a line-by-line visualization of the variables and data structures (lists, dicts, sets, linked lists, trees, grids). Use it to demonstrate an idea on a concrete input; keep programs small (under 25 lines) and print the interesting values.",
			map[string]any{
				"code":        map[string]any{"type": "string", "description": "complete Python code; it runs as a script"},
				"explanation": map[string]any{"type": "string", "description": "one sentence shown above the code: what to watch for"},
			}, "code"),
		fn("scratchpad_visualize",
			"Draw a step-by-step diagram in the teacher's scratchpad: a list of frames the learner steps through, each with a caption and one or more data structures. Use it to show how an algorithm changes a structure (pointers moving, a window sliding, a tree being built) without running code.",
			map[string]any{
				"title": map[string]any{"type": "string"},
				"frames": map[string]any{
					"type": "array",
					"items": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"caption":    map[string]any{"type": "string", "description": "what happens in this frame, one sentence"},
							"structures": map[string]any{"type": "array", "items": structure},
						},
						"required": []string{"caption", "structures"},
					},
				},
			}, "title", "frames"),
		fn("scratchpad_sketch",
			"Draw on the scratchpad's whiteboard in a hand-drawn style, like a teacher at a board: handwritten notes, boxes, circles, arrows and rows of array cells, drawn stroke by stroke. Coordinates run 0-100 across and 0-60 down. Give elements a step number to build the drawing up over several steps the learner clicks through (step 0 first); captions[i] explains step i. Use it for intuition: why an idea works, a pointer moving, what a recurrence means. Draw, don't just write: a row of cells for an array, arrows for pointers, circles or marks on the cells that matter, boxes for variables; keep text to short labels (six words or fewer). To move pointers over steps, redraw the same cells with new pointers in the next step. Example: [{\"type\":\"cells\",\"x\":15,\"y\":12,\"values\":[2,1,5,1,3],\"text\":\"nums\",\"pointers\":{\"left\":0,\"right\":2}}, {\"type\":\"cells\",\"step\":1,\"x\":15,\"y\":12,\"values\":[2,1,5,1,3],\"marks\":[0],\"pointers\":{\"left\":1,\"right\":3}}, {\"type\":\"text\",\"step\":1,\"x\":15,\"y\":48,\"text\":\"drop 2, add 1\",\"color\":\"red\"}]. Cells are 8 wide, so values[i] sits at x + 8*i; leave 20 units below a row of cells for its pointers, and put notes below that or to the right of the row.",
			map[string]any{
				"title":    map[string]any{"type": "string"},
				"captions": map[string]any{"type": "array", "items": map[string]any{"type": "string"}, "description": "one short caption per step"},
				"elements": map[string]any{
					"type": "array",
					"items": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"type":     map[string]any{"type": "string", "enum": sketchTypes},
							"step":     map[string]any{"type": "integer", "description": "which step reveals it (0 = first)"},
							"x":        map[string]any{"type": "number"},
							"y":        map[string]any{"type": "number"},
							"x2":       map[string]any{"type": "number", "description": "arrow/line end"},
							"y2":       map[string]any{"type": "number", "description": "arrow/line end"},
							"w":        map[string]any{"type": "number", "description": "box width"},
							"h":        map[string]any{"type": "number", "description": "box height"},
							"r":        map[string]any{"type": "number", "description": "circle/node radius"},
							"text":     map[string]any{"type": "string", "description": "text, or the label of a box/node/arrow"},
							"size":     map[string]any{"type": "string", "enum": []string{"s", "m", "l"}},
							"color":    map[string]any{"type": "string", "enum": sketchColors},
							"values":   map[string]any{"type": "array", "description": "cells: the values in the row"},
							"marks":    map[string]any{"type": "array", "items": map[string]any{"type": "integer"}, "description": "cells: indexes to circle"},
							"pointers": map[string]any{"type": "object", "description": "cells: pointer name -> index, e.g. {\"left\": 0, \"right\": 4}; draws a labelled arrow under each of those cells. Prefer this to placing arrows by hand."},
						},
						"required": []string{"type"},
					},
				},
			}, "title", "elements"),
		fn("ask_learner",
			"Ask the learner a question with clickable answers to check understanding, then stop and wait for their reply. Make wrong choices real misconceptions.",
			map[string]any{
				"question": map[string]any{"type": "string"},
				"choices":  map[string]any{"type": "array", "items": map[string]any{"type": "string"}, "description": "2 to 5 short answers"},
			}, "question", "choices"),
	}
}

func executeScratchpad(call ToolCall) (string, Action) {
	args := call.Function.Arguments
	switch call.Function.Name {
	case "scratchpad_run_code":
		code := stripLineNumbers(strings.TrimSpace(str(args, "code")))
		if code == "" {
			return "Error: code is empty.", nil
		}
		resp, err := sandbox.Execute(sandbox.ExecuteRequest{Code: code, Timeout: scratchTimeoutMs})
		if err != nil {
			return "Error: the scratchpad could not run the code: " + err.Error(), nil
		}
		steps := resp.Steps
		if len(steps) > maxScratchSteps {
			steps = steps[:maxScratchSteps]
		}
		act := Action{"name": "scratchpad_code", "code": code, "steps": steps, "output": truncate(resp.Output, 4000), "error": resp.Error}
		if e := strings.TrimSpace(str(args, "explanation")); e != "" {
			act["explanation"] = truncate(e, 300)
		}
		var b strings.Builder
		fmt.Fprintf(&b, "Ran it in the scratchpad; the learner can step through %d recorded steps.", len(resp.Steps))
		if resp.Error != "" {
			fmt.Fprintf(&b, "\nError: %s", truncate(resp.Error, 800))
		}
		if out := strings.TrimSpace(resp.Output); out != "" {
			fmt.Fprintf(&b, "\nOutput:\n%s", truncate(out, 1500))
		}
		if n := len(resp.Steps); n > 0 {
			if locals, err := json.Marshal(resp.Steps[n-1].Locals); err == nil {
				fmt.Fprintf(&b, "\nVariables at the last step: %s", truncate(string(locals), 1000))
			}
		}
		return b.String(), act

	case "scratchpad_visualize":
		frames := listArg(args, "frames")
		if len(frames) == 0 {
			return "Error: frames is empty.", nil
		}
		if len(frames) > maxFrames {
			frames = frames[:maxFrames]
		}
		clean := make([]any, 0, len(frames))
		for i, f := range frames {
			fm, ok := f.(map[string]any)
			if !ok {
				return fmt.Sprintf("Error: frame %d is not an object.", i+1), nil
			}
			structs := listArg(fm, "structures")
			for j, st := range structs {
				sm, ok := st.(map[string]any)
				if !ok || !validType(str(sm, "type")) {
					return fmt.Sprintf("Error: frame %d structure %d needs a type from %v.", i+1, j+1, structureTypes), nil
				}
			}
			clean = append(clean, map[string]any{"caption": str(fm, "caption"), "structures": structs})
		}
		if raw, _ := json.Marshal(clean); len(raw) > maxFrameBytes {
			return "Error: the diagram is too large; use smaller inputs or fewer frames.", nil
		}
		return fmt.Sprintf("Drew %d frames in the scratchpad; the learner can step through them.", len(clean)),
			Action{"name": "scratchpad_visualize", "title": truncate(str(args, "title"), 120), "frames": clean}

	case "scratchpad_sketch":
		raw := listArg(args, "elements")
		if len(raw) == 0 {
			return "Error: elements must be a non-empty JSON array of element objects.", nil
		}
		if len(raw) > maxSketchElements {
			raw = raw[:maxSketchElements]
		}
		steps := 0
		for i, e := range raw {
			em, ok := e.(map[string]any)
			if !ok || !contains(sketchTypes, str(em, "type")) {
				return fmt.Sprintf("Error: element %d needs a type from %v.", i+1, sketchTypes), nil
			}
			for _, k := range []string{"values", "marks"} {
				if _, isStr := em[k].(string); isStr {
					em[k] = listArg(em, k)
				}
			}
			steps = max(steps, intArg(em, "step")+1)
		}
		if steps > 20 {
			return "Error: use at most 20 steps.", nil
		}
		var captions []string
		if cs := listArg(args, "captions"); cs != nil {
			for _, c := range cs {
				if s, ok := c.(string); ok {
					captions = append(captions, truncate(s, 300))
				}
			}
		}
		if b, _ := json.Marshal(raw); len(b) > maxFrameBytes {
			return "Error: the sketch is too large; draw fewer elements.", nil
		}
		return fmt.Sprintf("Sketched %d elements over %d steps on the whiteboard; the learner can step through them.", len(raw), steps),
			Action{"name": "scratchpad_sketch", "title": truncate(str(args, "title"), 120), "elements": raw, "captions": captions}

	case "ask_learner":
		q := strings.TrimSpace(str(args, "question"))
		raw := listArg(args, "choices")
		var choices []string
		for _, c := range raw {
			if s, ok := c.(string); ok && strings.TrimSpace(s) != "" && len(choices) < 5 {
				choices = append(choices, truncate(strings.TrimSpace(s), 200))
			}
		}
		if q == "" || len(choices) < 2 {
			return "Error: ask_learner needs a question and choices as a JSON array of at least two strings, e.g. [\"yes\", \"no\"]. Don't retry more than once; ask in plain text instead.", nil
		}
		return "Asked. Stop here and wait for the learner's answer.", Action{"name": "ask_learner", "question": truncate(q, 400), "choices": choices}
	}
	return "Error: unknown tool", nil
}

func validType(t string) bool { return contains(structureTypes, t) }

func contains(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}
