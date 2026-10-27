package sandbox

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// Oxigen support: learner code in OxigenLang (https://www.oxigenlang.org)
// runs under the oxigen binary instead of python3. There's no tracer, so an
// Oxigen run has output and a judged answer but no steps to visualize.

//go:embed harness_lib.oxi
var oxigenHarnessLib string

//go:embed validate.py
var validatorRunner []byte

// OxigenBinary is the interpreter Oxigen code runs under.
var OxigenBinary = "oxigen"

// resultMarker starts the line harness_lib.oxi's emit() prints.
const resultMarker = "@@dsatutor-result@@"

// AllowedOxigenModules are the standard library modules learner code may
// introduce. Everything else reaches files, processes, the network or C.
var AllowedOxigenModules = []string{
	"array", "datetime", "encoding", "hash", "math", "random", "regex", "result", "strings", "time",
}

var (
	oxigenImportRe  = regexp.MustCompile(`\b(?:introduce|intro)\b[ \t]*([^\n]*)`)
	oxigenCommentRe = regexp.MustCompile(`//[^\n]*`)
	// Native functions such as __exec and __read_file are globals in every
	// Oxigen module, so no identifier may start with two underscores.
	oxigenNativeRe = regexp.MustCompile(`__[A-Za-z0-9_]`)
	oxigenStructRe = regexp.MustCompile(`(?m)^\s*struct\s+([A-Za-z_][A-Za-z0-9_]*)`)
)

// OxigenAvailable reports whether the oxigen binary is on PATH.
func OxigenAvailable() bool {
	_, err := exec.LookPath(OxigenBinary)
	return err == nil
}

// ValidateOxigenCode rejects imports outside AllowedOxigenModules, local file
// imports and the double-underscore natives.
func ValidateOxigenCode(code string) error {
	if loc := oxigenNativeRe.FindStringIndex(code); loc != nil {
		return fmt.Errorf("security error: names starting with __ are reserved (line %d)", lineOf(code, loc[0]))
	}
	allowed := map[string]bool{}
	for _, m := range AllowedOxigenModules {
		allowed[m] = true
	}
	// Blank out comments, keeping offsets, so "// introduce ..." in a note
	// isn't read as an import.
	scan := oxigenCommentRe.ReplaceAllStringFunc(code, func(c string) string { return strings.Repeat(" ", len(c)) })
	for _, m := range oxigenImportRe.FindAllStringSubmatchIndex(scan, -1) {
		target := strings.TrimSpace(scan[m[2]:m[3]])
		if i := strings.Index(target, "from "); strings.HasPrefix(target, "{") && i >= 0 {
			target = strings.TrimSpace(target[i+len("from "):])
		}
		if fields := strings.Fields(target); len(fields) > 0 {
			target = fields[0]
		}
		if !allowed[target] {
			return fmt.Errorf("security error: introduce %s is not allowed (line %d); available modules: %s",
				target, lineOf(code, m[0]), strings.Join(AllowedOxigenModules, ", "))
		}
	}
	return nil
}

func lineOf(s string, offset int) int {
	return strings.Count(s[:offset], "\n") + 1
}

func executeOxigen(req ExecuteRequest, timeout int) (*ExecuteResponse, error) {
	if !OxigenAvailable() {
		return failed("Oxigen isn't installed on the server: run `imlazy setup`, or see https://www.oxigenlang.org/download/"), nil
	}
	if err := ValidateOxigenCode(req.Code); err != nil {
		return failed(err.Error()), nil
	}

	tmpDir, err := os.MkdirTemp("", "dsatutor-oxigen-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	if req.Harness == nil {
		if err := os.WriteFile(filepath.Join(tmpDir, "main.oxi"), []byte(req.Code), 0600); err != nil {
			return nil, fmt.Errorf("failed to write code: %w", err)
		}
		stdout, stderr, err := runOxigen(tmpDir, "main.oxi", timeout)
		return oxigenResponse(stdout, stderr, err, timeout, false).ExecuteResponse, nil
	}

	spec := req.Harness
	solution := req.Code + "\n" + oxigenNodeStructs(spec, req.Code)
	if err := os.WriteFile(filepath.Join(tmpDir, "solution.oxi"), []byte(solution), 0600); err != nil {
		return nil, fmt.Errorf("failed to write code: %w", err)
	}
	inputPath := filepath.Join(tmpDir, "input.json")
	input, err := json.Marshal(map[string]interface{}{"inputs": spec.Inputs})
	if err != nil {
		return nil, fmt.Errorf("failed to encode inputs: %w", err)
	}
	if err := os.WriteFile(inputPath, input, 0600); err != nil {
		return nil, fmt.Errorf("failed to write inputs: %w", err)
	}
	harness, err := oxigenHarness(spec, req.Code, inputPath)
	if err != nil {
		return failed(err.Error()), nil
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "harness.oxi"), []byte(harness), 0600); err != nil {
		return nil, fmt.Errorf("failed to write harness: %w", err)
	}

	stdout, stderr, err := runOxigen(tmpDir, "harness.oxi", timeout)
	resp := oxigenResponse(stdout, stderr, err, timeout, true)
	if !resp.HasResult || resp.Error != "" {
		return resp.ExecuteResponse, nil
	}

	switch {
	case spec.ReturnType == "GraphNode" || spec.ReturnType == "RandomList":
		// The Python validators for the clone problems inspect node objects;
		// the harness answers their question (does the copy reuse nodes?) and
		// the judge compares the copy's shape.
		if shared, _ := resp.extra["shares"].(bool); shared {
			resp.Verdict = &Verdict{Message: "the copy reuses nodes from the original; build new nodes instead"}
		}
	case spec.Validator != "":
		verdict, err := runValidator(spec, resp.Result, timeout)
		if err != nil {
			resp.Error = err.Error()
			resp.Success = false
			return resp.ExecuteResponse, nil
		}
		resp.Verdict = verdict
	}
	return resp.ExecuteResponse, nil
}

func failed(msg string) *ExecuteResponse {
	return &ExecuteResponse{Success: false, Steps: []Step{}, Error: msg}
}

func runOxigen(dir, file string, timeout int) ([]byte, []byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, OxigenBinary, file)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "NO_COLOR=1")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		err = context.DeadlineExceeded
	}
	return stdout.Bytes(), stderr.Bytes(), err
}

// oxigenResult is an ExecuteResponse plus the harness's side channel.
type oxigenResult struct {
	*ExecuteResponse
	extra map[string]interface{}
}

func oxigenResponse(stdout, stderr []byte, runErr error, timeout int, judged bool) oxigenResult {
	res := oxigenResult{ExecuteResponse: &ExecuteResponse{Steps: []Step{}}}
	if errors.Is(runErr, context.DeadlineExceeded) {
		res.Error = fmt.Sprintf("Execution timed out after %d ms", timeout)
		return res
	}

	var printed []string
	var payload string
	scanner := bufio.NewScanner(bytes.NewReader(stdout))
	scanner.Buffer(make([]byte, 64*1024), 16*1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, resultMarker) {
			payload = strings.TrimPrefix(line, resultMarker)
			continue
		}
		printed = append(printed, line)
	}
	res.Output = strings.Join(printed, "\n")
	if res.Output != "" {
		res.Output += "\n"
	}

	if runErr != nil {
		res.Error = cleanOxigenError(string(stderr))
		if res.Error == "" {
			res.Error = "Execution failed: " + runErr.Error()
		}
		return res
	}
	res.Success = true
	if !judged {
		return res
	}
	if payload == "" {
		res.Success = false
		res.Error = "no result returned"
		return res
	}
	var out struct {
		Result interface{}            `json:"result"`
		Extra  map[string]interface{} `json:"extra"`
	}
	if err := json.Unmarshal([]byte(payload), &out); err != nil {
		res.Success = false
		res.Error = "could not read the result: " + err.Error()
		return res
	}
	if msg, _ := out.Extra["error"].(string); msg != "" {
		res.Success = false
		res.Error = msg
		return res
	}
	res.HasResult = true
	res.Result = out.Result
	res.extra = out.Extra
	return res
}

var (
	missingFunctionRe = regexp.MustCompile(`module '(?:\./solution|your code)' has no function '([A-Za-z_][A-Za-z0-9_]*)'`)
	missingMethodRe   = regexp.MustCompile(`method '([A-Za-z_][A-Za-z0-9_]*)' not found on ([A-Za-z_][A-Za-z0-9_]*)`)
	missingStructRe   = regexp.MustCompile(`undefined variable: ([A-Za-z_][A-Za-z0-9_]*)`)
)

// cleanOxigenError rewrites errors for the learner. An error located in the
// judge's harness file means the code doesn't define what the judge calls,
// and the harness's own lines and stack frames mean nothing to them.
func cleanOxigenError(stderr string) string {
	if m := missingFunctionRe.FindStringSubmatch(stderr); m != nil {
		return fmt.Sprintf("Your code has no function %s. The judge calls it by the name in the starter code, so keep that name.", m[1])
	}
	lines := strings.Split(strings.TrimRight(stderr, "\n"), "\n")
	if len(lines) > 1 && strings.Contains(lines[1], "--> harness.oxi") {
		if m := missingMethodRe.FindStringSubmatch(lines[0]); m != nil {
			return fmt.Sprintf("%s has no method %s. The judge calls each method by the name in the starter code, so keep those names.", m[2], m[1])
		}
		if m := missingStructRe.FindStringSubmatch(lines[0]); m != nil {
			return fmt.Sprintf("Your code has no struct %s. The judge builds it by the name in the starter code, so keep that name.", m[1])
		}
		return strings.ReplaceAll(lines[0], "./solution", "your code")
	}
	var keep []string
	for _, line := range strings.Split(strings.TrimRight(stderr, "\n"), "\n") {
		if strings.Contains(line, "harness.oxi") && strings.Contains(line, " at `") {
			continue
		}
		keep = append(keep, line)
	}
	msg := strings.Join(keep, "\n")
	msg = strings.ReplaceAll(msg, "./solution", "your code")
	return msg
}

// oxigenNodeStructs returns the node structs the problem uses, minus any
// the learner defined themselves. They go after the learner's code so error
// line numbers still match the editor.
func oxigenNodeStructs(spec *HarnessSpec, code string) string {
	defined := map[string]bool{}
	for _, m := range oxigenStructRe.FindAllStringSubmatch(code, -1) {
		defined[m[1]] = true
	}
	var b strings.Builder
	add := func(name, body string) {
		if !defined[name] {
			b.WriteString("\nstruct " + name + " {\n" + body + "}\n")
			defined[name] = true
		}
	}
	add("ListNode", "    val <generic>\n    next <ListNode> || <None>\n")
	add("TreeNode", "    val <generic>\n    left <TreeNode> || <None>\n    right <TreeNode> || <None>\n")
	switch {
	case usesType(spec, "GraphNode"):
		add("Node", "    val <generic>\n    neighbors <array>\n")
	case usesType(spec, "RandomList"):
		add("Node", "    val <generic>\n    next <Node> || <None>\n    random <Node> || <None>\n")
	}
	return b.String()
}

func usesType(spec *HarnessSpec, kind string) bool {
	if spec.ReturnType == kind {
		return true
	}
	for _, t := range spec.ParamTypes {
		if t == kind {
			return true
		}
	}
	return false
}

// oxigenHarness writes harness.oxi for one test case: imports, the library,
// and a main block that builds the arguments, calls the learner's code and
// emits the answer as JSON.
func oxigenHarness(spec *HarnessSpec, code, inputPath string) (string, error) {
	structs := map[string]bool{"ListNode": true, "TreeNode": true}
	for _, m := range oxigenStructRe.FindAllStringSubmatch(code, -1) {
		structs[m[1]] = true
	}
	if usesType(spec, "GraphNode") || usesType(spec, "RandomList") {
		structs["Node"] = true
	}
	var names []string
	for name := range structs {
		names = append(names, name)
	}
	sort.Strings(names)

	var b strings.Builder
	b.WriteString("introduce {parse, stringify, read} from json\n")
	b.WriteString("introduce .solution\n")
	b.WriteString("introduce {" + strings.Join(names, ", ") + "} from .solution\n\n")
	b.WriteString(oxigenHarnessLib)
	b.WriteString("\nmain {\n")
	fmt.Fprintf(&b, "    raw := read(%s)[\"inputs\"]\n", oxigenString(inputPath))

	switch spec.Mode {
	case "design":
		if err := writeOxigenDesign(&b, spec); err != nil {
			return "", err
		}
	case "roundtrip":
		writeOxigenRoundtrip(&b, spec)
	default:
		writeOxigenFunction(&b, spec)
	}
	b.WriteString("}\n")
	return b.String(), nil
}

// oxigenString quotes s as an Oxigen string literal. Braces are escaped
// because a bare { starts an interpolation.
func oxigenString(s string) string {
	r := strings.NewReplacer(`\`, `\\`, `"`, `\"`, "{", `\{`, "}", `\}`, "\n", `\n`)
	return `"` + r.Replace(s) + `"`
}

// writeOxigenArgs builds args, a map of the converted inputs.
func writeOxigenArgs(b *strings.Builder, spec *HarnessSpec) {
	b.WriteString("    args := {}\n")
	var keys []string
	for k := range spec.Inputs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, name := range keys {
		kind := spec.ParamTypes[name]
		value := fmt.Sprintf("raw[%s]", oxigenString(name))
		switch {
		case kind == "ListNode":
			value = "build_list(" + value + ")"
		case kind == "ListNode[]":
			fmt.Fprintf(b, "    lists := []\n    each v in %s {\n        lists := push(lists, build_list(v))\n    }\n", value)
			value = "lists"
		case kind == "TreeNode":
			value = "build_tree(" + value + ")"
		case kind == "GraphNode":
			value = "build_graph(" + value + ")"
		case kind == "RandomList":
			value = "build_random_list(" + value + ")"
		case strings.HasPrefix(kind, "CycleList:"):
			value = fmt.Sprintf("build_cycle_list(%s, raw[%s])", value, oxigenString(strings.TrimPrefix(kind, "CycleList:")))
		case strings.HasPrefix(kind, "TreeNodeRef:"):
			continue // second pass, once the tree exists
		}
		fmt.Fprintf(b, "    args[%s] = %s\n", oxigenString(name), value)
	}
	for _, name := range keys {
		if kind := spec.ParamTypes[name]; strings.HasPrefix(kind, "TreeNodeRef:") {
			fmt.Fprintf(b, "    args[%s] = find_tree_node(args[%s], raw[%s])\n",
				oxigenString(name), oxigenString(strings.TrimPrefix(kind, "TreeNodeRef:")), oxigenString(name))
		}
	}
}

func writeOxigenFunction(b *strings.Builder, spec *HarnessSpec) {
	writeOxigenArgs(b, spec)
	var call []string
	for _, p := range spec.Params {
		call = append(call, fmt.Sprintf("args[%s]", oxigenString(p)))
	}
	fmt.Fprintf(b, "    result := solution.%s(%s)\n", spec.Function, strings.Join(call, ", "))
	if spec.InPlace != "" {
		fmt.Fprintf(b, "    result := args[%s]\n", oxigenString(spec.InPlace))
	}
	b.WriteString("    extra := {}\n")
	for name, kind := range spec.ParamTypes {
		var fields string
		switch kind {
		case "GraphNode":
			fields = `["neighbors"]`
		case "RandomList":
			fields = `["next", "random"]`
		default:
			continue
		}
		fmt.Fprintf(b, "    extra[\"shares\"] = shares_nodes(args[%s], result, %s)\n", oxigenString(name), fields)
	}
	fmt.Fprintf(b, "    emit(to_json(result, %s), extra)\n", oxigenString(spec.ReturnType))
}

func writeOxigenRoundtrip(b *strings.Builder, spec *HarnessSpec) {
	writeOxigenArgs(b, spec)
	names := strings.SplitN(spec.Function, ",", 2)
	enc, dec := strings.TrimSpace(names[0]), strings.TrimSpace(names[1])
	owner := "solution"
	if spec.ClassName != "" {
		fmt.Fprintf(b, "    codec <%s>\n", spec.ClassName)
		owner = "codec"
	}
	fmt.Fprintf(b, "    encoded := %s.%s(args[%s])\n", owner, enc, oxigenString(spec.Params[0]))
	b.WriteString("    option {\n")
	fmt.Fprintf(b, "        type(encoded) != \"STRING\" -> {\n            emit(None, {\"error\": \"%s must return a string, got \" + type(encoded)})\n            give None\n        }\n", enc)
	b.WriteString("    }\n")
	fmt.Fprintf(b, "    emit(to_json(%s.%s(encoded), %s), {})\n", owner, dec, oxigenString(spec.ReturnType))
}

// writeOxigenDesign unrolls the operations of a design test case into
// direct method calls. Oxigen functions return their last expression, so an
// operation whose expected answer is null is treated as returning nothing.
func writeOxigenDesign(b *strings.Builder, spec *HarnessSpec) error {
	ops, _ := spec.Inputs["operations"].([]interface{})
	valuesKey := "values"
	values, _ := spec.Inputs["values"].([]interface{})
	if values == nil {
		valuesKey = "args"
		values, _ = spec.Inputs["args"].([]interface{})
	}
	if len(ops) == 0 {
		return errors.New("design test case has no operations")
	}
	expected, _ := spec.Expected.([]interface{})
	arity := func(i int) int {
		if i < len(values) {
			if v, ok := values[i].([]interface{}); ok {
				return len(v)
			}
		}
		return 0
	}
	callArgs := func(i int) string {
		var parts []string
		for j := 0; j < arity(i); j++ {
			parts = append(parts, fmt.Sprintf("vals[%d][%d]", i, j))
		}
		return strings.Join(parts, ", ")
	}

	fmt.Fprintf(b, "    vals := raw[%s]\n", oxigenString(valuesKey))
	fmt.Fprintf(b, "    obj <%s>\n", ops[0])
	fmt.Fprintf(b, "    obj.init(%s)\n", callArgs(0))
	b.WriteString("    out := [None]\n")
	for i := 1; i < len(ops); i++ {
		op, _ := ops[i].(string)
		call := fmt.Sprintf("obj.%s(%s)", op, callArgs(i))
		if i < len(expected) && expected[i] == nil {
			fmt.Fprintf(b, "    %s\n    out := push(out, None)\n", call)
		} else {
			fmt.Fprintf(b, "    out := push(out, to_json(%s, \"\"))\n", call)
		}
	}
	b.WriteString("    emit(out, {})\n")
	return nil
}

// runValidator runs a problem's trusted Python check() on the Oxigen
// answer. Oxigen answers arrive as JSON, so args is the raw input.
func runValidator(spec *HarnessSpec, result interface{}, timeout int) (*Verdict, error) {
	tmpDir, err := os.MkdirTemp("", "dsatutor-validate-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)
	script := filepath.Join(tmpDir, "validate.py")
	if err := os.WriteFile(script, validatorRunner, 0600); err != nil {
		return nil, fmt.Errorf("failed to write validator: %w", err)
	}
	payload, err := json.Marshal(map[string]interface{}{
		"validator": spec.Validator,
		"raw":       spec.Inputs,
		"result":    result,
		"expected":  spec.Expected,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to encode validator input: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, "python3", "-I", script)
	cmd.Dir = tmpDir
	cmd.Stdin = bytes.NewReader(payload)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("validator failed: %v", err)
	}
	var verdict Verdict
	if err := json.Unmarshal(out, &verdict); err != nil {
		return nil, fmt.Errorf("validator output unreadable: %v", err)
	}
	return &verdict, nil
}
