// Package judge runs learner code against a problem's test cases and decides
// whether each answer is correct. The web handlers and the content tests both
// go through it, so the reference solutions are checked by the same code path
// a learner's Submit uses.
package judge

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"

	"dsatutor/internal/practice"
	"dsatutor/internal/sandbox"
)

// Result is the outcome of one test case.
type Result struct {
	Passed    bool           `json:"passed"`
	Actual    interface{}    `json:"actual,omitempty"`
	Output    string         `json:"output,omitempty"` // anything the code printed
	Error     string         `json:"error,omitempty"`
	Message   string         `json:"message,omitempty"` // validator explanation, if any
	Steps     []sandbox.Step `json:"-"`
	Truncated bool           `json:"truncated,omitempty"`
}

// Options control how a case is executed.
type Options struct {
	Trace bool // record line-by-line steps for the visualizer
}

// Run executes code against one test case of p.
func Run(p *practice.Problem, code string, tc *practice.TestCase, opts Options) (*Result, error) {
	spec, err := Spec(p, tc)
	if err != nil {
		return &Result{Error: err.Error()}, nil
	}
	resp, err := sandbox.Execute(sandbox.ExecuteRequest{
		Code:    code,
		Timeout: tc.Timeout,
		NoTrace: !opts.Trace,
		Harness: spec,
	})
	if err != nil {
		return nil, err
	}

	res := &Result{
		Output:    resp.Output,
		Error:     resp.Error,
		Steps:     resp.Steps,
		Truncated: resp.Truncated,
	}
	if !resp.Success || !resp.HasResult {
		if res.Error == "" {
			res.Error = "no result returned"
		}
		return res, nil
	}

	res.Actual = resp.Result
	if resp.Verdict != nil {
		res.Passed = resp.Verdict.Passed
		res.Message = resp.Verdict.Message
		return res, nil
	}
	mode := p.Harness.Compare
	if tc.OrderIndependent && mode == "" {
		mode = "unordered-nested"
	}
	res.Passed = Compare(resp.Result, tc.Expected, mode)
	return res, nil
}

// Spec builds the sandbox harness description for one test case.
func Spec(p *practice.Problem, tc *practice.TestCase) (*sandbox.HarnessSpec, error) {
	inputs, err := asObject(tc.Input)
	if err != nil {
		return nil, err
	}
	h := p.Harness
	spec := &sandbox.HarnessSpec{
		Mode:       h.Mode,
		Function:   h.Function,
		ClassName:  h.ClassName,
		Inputs:     inputs,
		ParamTypes: h.ParamTypes,
		ReturnType: h.ReturnType,
		InPlace:    h.InPlace,
		Prelude:    h.Prelude,
		Validator:  h.Validator,
		Expected:   normalize(tc.Expected),
	}

	switch h.Mode {
	case "design":
		return spec, nil
	case "roundtrip":
		if !strings.Contains(h.Function, ",") {
			return nil, fmt.Errorf("roundtrip problem %s must name \"encode,decode\" in Harness.Function", p.ID)
		}
		if len(inputs) != 1 {
			return nil, fmt.Errorf("roundtrip problem %s takes exactly one input", p.ID)
		}
		for k := range inputs {
			spec.Params = []string{k}
		}
		return spec, nil
	case "", "function":
		spec.Mode = "function"
		name, params := Signature(p.StarterCode)
		if h.Function != "" {
			name = h.Function
		}
		if name == "" {
			return nil, fmt.Errorf("could not find a function in the starter code of %s", p.ID)
		}
		spec.Function = name
		spec.Params = params
		for _, param := range params {
			if _, ok := inputs[param]; !ok {
				return nil, fmt.Errorf("test case of %s has no input named %q", p.ID, param)
			}
		}
		return spec, nil
	default:
		return nil, fmt.Errorf("unknown harness mode %q", h.Mode)
	}
}

// Signature returns the name and parameters (without self) of the first
// top-level function in Python source.
func Signature(src string) (string, []string) {
	for _, line := range strings.Split(src, "\n") {
		if !strings.HasPrefix(line, "def ") {
			continue
		}
		rest := strings.TrimPrefix(line, "def ")
		open := strings.Index(rest, "(")
		closeIdx := strings.LastIndex(rest, ")")
		if open < 0 || closeIdx < open {
			continue
		}
		name := strings.TrimSpace(rest[:open])
		var params []string
		for _, raw := range splitTopLevel(rest[open+1 : closeIdx]) {
			p := strings.TrimSpace(raw)
			if i := strings.IndexAny(p, ":="); i >= 0 {
				p = strings.TrimSpace(p[:i])
			}
			p = strings.TrimLeft(p, "*")
			if p != "" && p != "self" {
				params = append(params, p)
			}
		}
		return name, params
	}
	return "", nil
}

// splitTopLevel splits on commas that aren't inside brackets, so type hints
// like Dict[str, int] stay in one piece.
func splitTopLevel(s string) []string {
	var parts []string
	depth, start := 0, 0
	for i, r := range s {
		switch r {
		case '[', '(', '{':
			depth++
		case ']', ')', '}':
			depth--
		case ',':
			if depth == 0 {
				parts = append(parts, s[start:i])
				start = i + 1
			}
		}
	}
	return append(parts, s[start:])
}

func asObject(input interface{}) (map[string]interface{}, error) {
	obj, ok := normalize(input).(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("test input must be an object of named arguments, got %T", input)
	}
	return obj, nil
}

// normalize converts typed Go values ([]int, [][]string, ...) into the
// generic form JSON decoding produces, so both sides compare alike.
func normalize(v interface{}) interface{} {
	b, err := json.Marshal(v)
	if err != nil {
		return v
	}
	var out interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return v
	}
	return out
}

// Compare reports whether actual matches expected under the given mode.
func Compare(actual, expected interface{}, mode string) bool {
	a, e := normalize(actual), normalize(expected)
	switch mode {
	case "unordered", "unordered-nested":
		al, aok := a.([]interface{})
		el, eok := e.([]interface{})
		if !aok || !eok {
			return equal(a, e)
		}
		return equal(canonicalList(al, mode == "unordered-nested"), canonicalList(el, mode == "unordered-nested"))
	default:
		return equal(a, e)
	}
}

// canonicalList sorts a list by each element's JSON form; nested also sorts
// every inner list first.
func canonicalList(list []interface{}, nested bool) []interface{} {
	items := make([]interface{}, len(list))
	keys := make([]string, len(list))
	for i, v := range list {
		if inner, ok := v.([]interface{}); ok && nested {
			v = canonicalList(inner, false)
		}
		items[i] = v
		b, _ := json.Marshal(v)
		keys[i] = string(b)
	}
	idx := make([]int, len(items))
	for i := range idx {
		idx[i] = i
	}
	sort.SliceStable(idx, func(x, y int) bool { return keys[idx[x]] < keys[idx[y]] })
	out := make([]interface{}, len(items))
	for i, j := range idx {
		out[i] = items[j]
	}
	return out
}

// equal is deep equality where numbers match within 1e-5, so 2 equals 2.0
// and floating-point answers like 2.50000 pass.
func equal(a, b interface{}) bool {
	switch av := a.(type) {
	case float64:
		bv, ok := b.(float64)
		if !ok {
			return false
		}
		diff := math.Abs(av - bv)
		return diff <= 1e-5 || diff <= 1e-5*math.Max(math.Abs(av), math.Abs(bv))
	case []interface{}:
		bv, ok := b.([]interface{})
		if !ok || len(av) != len(bv) {
			return false
		}
		for i := range av {
			if !equal(av[i], bv[i]) {
				return false
			}
		}
		return true
	case map[string]interface{}:
		bv, ok := b.(map[string]interface{})
		if !ok || len(av) != len(bv) {
			return false
		}
		for k, v := range av {
			if !equal(v, bv[k]) {
				return false
			}
		}
		return true
	default:
		return a == b
	}
}
