package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"dsatutor/internal/practice"
	"dsatutor/internal/sandbox"
)

// PracticeServer handles practice-related API endpoints.
type PracticeServer struct {
	loader *practice.ProblemLoader
}

// NewPracticeServer creates a new practice server with loaded problems.
func NewPracticeServer(loader *practice.ProblemLoader) *PracticeServer {
	return &PracticeServer{
		loader: loader,
	}
}

// RegisterRoutes adds practice routes to the provided mux.
func (ps *PracticeServer) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/practice/problems", ps.handleListProblems)
	mux.HandleFunc("/api/practice/problems/", ps.handleGetProblem)
	mux.HandleFunc("/api/practice/categories", ps.handleListCategories)
	mux.HandleFunc("/api/practice/run", ps.handleRunCode)
	mux.HandleFunc("/api/practice/submit", ps.handleSubmitCode)
	mux.HandleFunc("/api/practice/hints/", ps.handleGetHint)
	mux.HandleFunc("/api/practice/solution/", ps.handleGetSolution)
	mux.HandleFunc("/api/practice/solution-viz/", ps.handleVisualizeSolution)
	mux.HandleFunc("/api/practice/progress/export", ps.handleExportProgress)
	mux.HandleFunc("/api/practice/progress/import", ps.handleImportProgress)
}

// handleListProblems returns all problems or filtered results.
func (ps *PracticeServer) handleListProblems(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get query parameters for filtering
	category := r.URL.Query().Get("category")
	difficulty := r.URL.Query().Get("difficulty")
	tag := r.URL.Query().Get("tag")
	search := r.URL.Query().Get("search")

	var problems []*practice.Problem
	if category != "" || difficulty != "" || tag != "" || search != "" {
		problems = ps.loader.FilterProblems(category, difficulty, tag, search)
	} else {
		problems = ps.loader.GetAllProblems()
	}

	// Return problem summaries (not full content)
	type ProblemSummary struct {
		ID              string   `json:"id"`
		Number          int      `json:"number"`
		Title           string   `json:"title"`
		Slug            string   `json:"slug"`
		Difficulty      string   `json:"difficulty"`
		Category        string   `json:"category"`
		Tags            []string `json:"tags"`
		RelatedChapters []int    `json:"relatedChapters"`
	}

	summaries := make([]ProblemSummary, len(problems))
	for i, p := range problems {
		summaries[i] = ProblemSummary{
			ID:              p.ID,
			Number:          p.Number,
			Title:           p.Title,
			Slug:            p.Slug,
			Difficulty:      p.Difficulty,
			Category:        p.Category,
			Tags:            p.Tags,
			RelatedChapters: p.RelatedChapters,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"problems": summaries,
		"total":    len(summaries),
		"stats":    ps.loader.GetStats(),
	})
}

// handleGetProblem returns a single problem by ID.
func (ps *PracticeServer) handleGetProblem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract problem ID from path: /api/practice/problems/{id}
	path := strings.TrimPrefix(r.URL.Path, "/api/practice/problems/")
	problemID := strings.TrimSuffix(path, "/")

	if problemID == "" {
		http.Error(w, "Problem ID required", http.StatusBadRequest)
		return
	}

	problem, err := ps.loader.GetProblem(problemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Return problem with visible test cases only (hide hidden test cases' expected values)
	type SafeTestCase struct {
		ID       string      `json:"id,omitempty"`
		Input    interface{} `json:"input"`
		Expected interface{} `json:"expected,omitempty"`
		Hidden   bool        `json:"hidden"`
		Timeout  int         `json:"timeout,omitempty"`
	}

	safeTestCases := make([]SafeTestCase, len(problem.TestCases))
	for i, tc := range problem.TestCases {
		safeTestCases[i] = SafeTestCase{
			ID:      tc.ID,
			Input:   tc.Input,
			Hidden:  tc.Hidden,
			Timeout: tc.Timeout,
		}
		if !tc.Hidden {
			safeTestCases[i].Expected = tc.Expected
		}
	}

	// Build response without full solution (user must explicitly request it)
	response := map[string]interface{}{
		"id":              problem.ID,
		"number":          problem.Number,
		"title":           problem.Title,
		"slug":            problem.Slug,
		"difficulty":      problem.Difficulty,
		"category":        problem.Category,
		"tags":            problem.Tags,
		"relatedChapters": problem.RelatedChapters,
		"description":     problem.Description,
		"constraints":     problem.Constraints,
		"examples":        problem.Examples,
		"testCases":       safeTestCases,
		"timeComplexity":  problem.TimeComplexity,
		"spaceComplexity": problem.SpaceComplexity,
		"starterCode":     problem.StarterCode,
		"hints":           problem.Hints,
		"hintCount":       len(problem.Hints),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleListCategories returns all problem categories.
func (ps *PracticeServer) handleListCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	categories := ps.loader.GetCategories()

	// Add problem counts to each category
	type CategoryWithCount struct {
		practice.Category
		ProblemCount int `json:"problemCount"`
	}

	result := make([]CategoryWithCount, len(categories))
	for i, cat := range categories {
		problems := ps.loader.GetProblemsByCategory(cat.ID)
		result[i] = CategoryWithCount{
			Category:     cat,
			ProblemCount: len(problems),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// handleRunCode executes code on a single test case with visualization.
func (ps *PracticeServer) handleRunCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req practice.RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(practice.RunResponse{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}

	// Get the problem
	problem, err := ps.loader.GetProblem(req.ProblemID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(practice.RunResponse{
			Success: false,
			Error:   "Problem not found: " + err.Error(),
		})
		return
	}

	// Find test case
	var testCase *practice.TestCase
	if req.TestCaseID != "" {
		for _, tc := range problem.TestCases {
			if tc.ID == req.TestCaseID {
				testCase = &tc
				break
			}
		}
	} else if req.TestCaseIndex != nil && *req.TestCaseIndex >= 0 && *req.TestCaseIndex < len(problem.TestCases) {
		testCase = &problem.TestCases[*req.TestCaseIndex]
	} else if len(problem.TestCases) > 0 {
		// Use first visible test case by default
		for _, tc := range problem.TestCases {
			if !tc.Hidden {
				testCase = &tc
				break
			}
		}
		if testCase == nil {
			testCase = &problem.TestCases[0]
		}
	}

	if testCase == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(practice.RunResponse{
			Success: false,
			Error:   "No test case found",
		})
		return
	}

	// Generate code that runs with test input
	fullCode := generateTestWrapper(req.Code, problem, testCase)

	// Execute in sandbox
	execReq := sandbox.ExecuteRequest{
		Code:    fullCode,
		Timeout: testCase.Timeout,
	}

	execResp, err := sandbox.Execute(execReq)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(practice.RunResponse{
			Success: false,
			Error:   "Execution error: " + err.Error(),
		})
		return
	}

	// Build response
	response := practice.RunResponse{
		Success:    execResp.Success,
		Output:     stripSentinel(execResp.Output),
		Error:      execResp.Error,
		Steps:      execResp.Steps,
		Structures: nil, // Will be populated from steps if available
	}

	// Compare output with expected value
	if execResp.Success && testCase.Expected != nil {
		response.Passed = compareOutput(execResp.Output, testCase.Expected, testCase.OrderIndependent)
	}

	// Extract structures from last step if available
	if len(execResp.Steps) > 0 {
		lastStep := execResp.Steps[len(execResp.Steps)-1]
		response.Structures = lastStep.Structures
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleSubmitCode validates code against all test cases.
func (ps *PracticeServer) handleSubmitCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req practice.SubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(practice.SubmitResponse{
			Status: practice.StatusCompileError,
		})
		return
	}

	// Get the problem
	problem, err := ps.loader.GetProblem(req.ProblemID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(practice.SubmitResponse{
			Status: practice.StatusCompileError,
		})
		return
	}

	response := practice.NewSubmitResponse()
	response.TotalCount = len(problem.TestCases)

	// Run against all test cases
	for _, tc := range problem.TestCases {
		fullCode := generateTestWrapper(req.Code, problem, &tc)

		execReq := sandbox.ExecuteRequest{
			Code:    fullCode,
			Timeout: tc.Timeout,
		}

		execResp, err := sandbox.Execute(execReq)

		result := practice.TestResult{
			TestCaseID: tc.ID,
			Hidden:     tc.Hidden,
		}

		if err != nil {
			result.Passed = false
			result.Error = err.Error()
		} else if !execResp.Success {
			result.Passed = false
			result.Error = execResp.Error
		} else {
			// Parse output and compare with expected
			result.Passed = compareOutput(execResp.Output, tc.Expected, tc.OrderIndependent)
			if payload, ok := extractResultPayload(execResp.Output); ok {
				result.Actual = json.RawMessage(payload)
			} else {
				result.Actual = stripSentinel(execResp.Output)
			}
			if !tc.Hidden {
				result.Input = tc.Input
				result.Expected = tc.Expected
			}
		}

		if result.Passed {
			response.PassedCount++
		}

		response.Results = append(response.Results, result)
	}

	response.CalculateStatus()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleGetHint returns a specific hint for a problem.
func (ps *PracticeServer) handleGetHint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract problem ID and hint level from path: /api/practice/hints/{problemId}/{level}
	path := strings.TrimPrefix(r.URL.Path, "/api/practice/hints/")
	parts := strings.Split(strings.TrimSuffix(path, "/"), "/")

	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Problem ID required", http.StatusBadRequest)
		return
	}

	problemID := parts[0]
	level := 1
	if len(parts) > 1 {
		if lvl, err := strconv.Atoi(parts[1]); err == nil {
			level = lvl
		}
	}

	problem, err := ps.loader.GetProblem(problemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	hint := problem.GetHintByLevel(level)
	if hint == nil {
		http.Error(w, "Hint not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"hint":     hint,
		"maxLevel": problem.MaxHintLevel(),
	})
}

// handleGetSolution returns the full solution for a problem.
func (ps *PracticeServer) handleGetSolution(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract problem ID from path: /api/practice/solution/{problemId}
	path := strings.TrimPrefix(r.URL.Path, "/api/practice/solution/")
	problemID := strings.TrimSuffix(path, "/")

	if problemID == "" {
		http.Error(w, "Problem ID required", http.StatusBadRequest)
		return
	}

	problem, err := ps.loader.GetProblem(problemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(problem.Solution)
}

// handleVisualizeSolution runs the solution code through the sandbox for visualization.
func (ps *PracticeServer) handleVisualizeSolution(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract problem ID from path: /api/practice/solution-viz/{problemId}
	path := strings.TrimPrefix(r.URL.Path, "/api/practice/solution-viz/")
	problemID := strings.TrimSuffix(path, "/")

	if problemID == "" {
		http.Error(w, "Problem ID required", http.StatusBadRequest)
		return
	}

	// Parse request body for optional test case selection
	var req struct {
		TestCaseIndex int `json:"testCaseIndex"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	problem, err := ps.loader.GetProblem(problemID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(practice.RunResponse{
			Success: false,
			Error:   "Problem not found: " + err.Error(),
		})
		return
	}

	// Check if solution code exists
	if problem.Solution.Code == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(practice.RunResponse{
			Success: false,
			Error:   "No solution code available for this problem",
		})
		return
	}

	// Find test case
	var testCase *practice.TestCase
	if req.TestCaseIndex >= 0 && req.TestCaseIndex < len(problem.TestCases) {
		testCase = &problem.TestCases[req.TestCaseIndex]
	} else if len(problem.TestCases) > 0 {
		// Use first visible test case by default
		for _, tc := range problem.TestCases {
			if !tc.Hidden {
				testCase = &tc
				break
			}
		}
		if testCase == nil {
			testCase = &problem.TestCases[0]
		}
	}

	if testCase == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(practice.RunResponse{
			Success: false,
			Error:   "No test case found",
		})
		return
	}

	// Generate code that runs the solution with test input
	fullCode := generateTestWrapper(problem.Solution.Code, problem, testCase)

	// Execute in sandbox
	execReq := sandbox.ExecuteRequest{
		Code:    fullCode,
		Timeout: testCase.Timeout,
	}

	execResp, err := sandbox.Execute(execReq)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(practice.RunResponse{
			Success: false,
			Error:   "Execution error: " + err.Error(),
		})
		return
	}

	// Build response with visualization steps
	response := practice.RunResponse{
		Success:    execResp.Success,
		Output:     stripSentinel(execResp.Output),
		Error:      execResp.Error,
		Steps:      execResp.Steps,
		Structures: nil,
	}

	// Compare output with expected value
	if execResp.Success && testCase.Expected != nil {
		response.Passed = compareOutput(execResp.Output, testCase.Expected, testCase.OrderIndependent)
	}

	// Extract structures from last step if available
	if len(execResp.Steps) > 0 {
		lastStep := execResp.Steps[len(execResp.Steps)-1]
		response.Structures = lastStep.Structures
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleExportProgress exports user progress as JSON.
func (ps *PracticeServer) handleExportProgress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Progress is stored client-side, this endpoint just validates the format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Export progress from browser localStorage",
		"key":     "dsa-practice-progress",
	})
}

// handleImportProgress validates and echoes back imported progress.
func (ps *PracticeServer) handleImportProgress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var progress practice.UserProgress
	if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid progress format: " + err.Error(),
		})
		return
	}

	// Validate problem IDs exist
	for problemID := range progress.SolvedProblems {
		if _, err := ps.loader.GetProblem(problemID); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Unknown problem ID: " + problemID,
			})
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"valid":    true,
		"progress": progress,
	})
}

// resultSentinel prefixes the JSON-serialized return value in sandbox output,
// so user print() calls can't corrupt grading.
const resultSentinel = "__DSA_RESULT__:"

// pyHarness is appended AFTER user code (so editor line numbers stay valid).
// It builds linked lists / trees / graphs from JSON test inputs and serializes
// results back to JSON. All names start with "_" so the tracer hides them.
// It must not use builtins blocked by sandbox.ValidateCode (no hasattr etc.).
const pyHarness = `
# --- dsatutor test harness ---
import json

def __dsa_build_list__(vals):
    head = None
    tail = None
    for v in vals or []:
        n = ListNode(v)
        if head is None:
            head = n
        else:
            tail.next = n
        tail = n
    return head

def __dsa_build_random_list__(pairs):
    ns = [ListNode(p[0]) for p in pairs]
    for i in range(len(ns)):
        if i + 1 < len(ns):
            ns[i].next = ns[i + 1]
        ns[i].random = None if pairs[i][1] is None else ns[pairs[i][1]]
    return ns[0] if ns else None

def __dsa_make_cycle__(head, pos):
    if head is None or pos is None or pos < 0:
        return head
    ns = []
    cur = head
    while cur:
        ns.append(cur)
        cur = cur.next
    if pos < len(ns):
        ns[-1].next = ns[pos]
    return head

def __dsa_build_tree__(vals):
    if not vals or vals[0] is None:
        return None
    root = TreeNode(vals[0])
    q = [root]
    qi = 0
    i = 1
    while qi < len(q) and i < len(vals):
        node = q[qi]
        qi += 1
        if i < len(vals):
            if vals[i] is not None:
                node.left = TreeNode(vals[i])
                q.append(node.left)
            i += 1
        if i < len(vals):
            if vals[i] is not None:
                node.right = TreeNode(vals[i])
                q.append(node.right)
            i += 1
    return root

def __dsa_find_node__(root, val):
    if root is None:
        return None
    if root.val == val:
        return root
    found = __dsa_find_node__(root.left, val)
    return found if found is not None else __dsa_find_node__(root.right, val)

def __dsa_build_graph__(adj):
    ns = [GraphNode(i + 1) for i in range(len(adj))]
    for i, nbrs in enumerate(adj):
        ns[i].neighbors = [ns[j - 1] for j in nbrs]
    return ns[0] if ns else None

def __dsa_ser__(x):
    if x is None or isinstance(x, (bool, int, float, str)):
        return x
    if isinstance(x, (list, tuple)):
        return [__dsa_ser__(v) for v in x]
    if isinstance(x, set):
        return sorted([__dsa_ser__(v) for v in x], key=lambda v: str(v))
    if isinstance(x, dict):
        return {str(k): __dsa_ser__(v) for k, v in x.items()}
    try:
        x.next
        seen = set()
        out = []
        cur = x
        while cur is not None and id(cur) not in seen:
            seen.add(id(cur))
            out.append(__dsa_ser__(cur.val))
            cur = cur.next
        return out
    except AttributeError:
        pass
    try:
        x.left
        out = []
        q = [x]
        qi = 0
        while qi < len(q):
            n = q[qi]
            qi += 1
            if n is None:
                out.append(None)
            else:
                out.append(__dsa_ser__(n.val))
                q.append(n.left)
                q.append(n.right)
        while out and out[-1] is None:
            out.pop()
        return out
    except AttributeError:
        pass
    try:
        x.neighbors
        seen = set()
        order = []
        stack = [x]
        while stack:
            n = stack.pop()
            if id(n) in seen:
                continue
            seen.add(id(n))
            order.append(n)
            for nb in n.neighbors:
                stack.append(nb)
        adj = {}
        for n in order:
            adj[n.val] = [nb.val for nb in n.neighbors]
        return [adj[k] for k in sorted(adj)]
    except AttributeError:
        pass
    return repr(x)

def __dsa_ser_random__(head):
    ns = []
    idx = {}
    cur = head
    while cur is not None:
        idx[id(cur)] = len(ns)
        ns.append(cur)
        cur = cur.next
    out = []
    for n in ns:
        try:
            r = n.random
        except AttributeError:
            r = None
        out.append([n.val, None if r is None else idx[id(r)]])
    return out
# --- end harness ---
`

var pyIdentRe = regexp.MustCompile(`^[A-Za-z_]\w*$`)

var pyClassRe = regexp.MustCompile(`(?m)^class\s+([A-Za-z_]\w*)`)

// extractClassName returns the first top-level class defined in the code, if any.
func extractClassName(code string) string {
	if m := pyClassRe.FindStringSubmatch(code); m != nil {
		return m[1]
	}
	return ""
}

// linkedListParams are parameter names converted from JSON arrays to ListNode chains.
var linkedListParams = map[string]bool{"head": true, "list1": true, "list2": true, "l1": true, "l2": true}

// pyLoads emits a Python expression that parses the value from JSON,
// so null/true/false round-trip correctly (they are not Python literals).
func pyLoads(value interface{}) string {
	b, _ := json.Marshal(value)
	return "json.loads(" + strconv.Quote(string(b)) + ")"
}

// isEmptySlice reports whether v is a zero-length slice/array.
func isEmptySlice(v interface{}) bool {
	if v == nil {
		return false
	}
	rv := reflect.ValueOf(v)
	return (rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array) && rv.Len() == 0
}

// isScalar reports whether expected is a plain value (not a collection).
func isScalar(v interface{}) bool {
	if v == nil {
		return false
	}
	switch reflect.ValueOf(v).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return false
	}
	return true
}

// generateTestWrapper creates code that runs user code against one test case
// and prints the JSON-serialized result after resultSentinel.
func generateTestWrapper(userCode string, problem *practice.Problem, tc *practice.TestCase) string {
	input, _ := tc.Input.(map[string]interface{})

	var b strings.Builder
	b.WriteString(userCode)
	b.WriteString("\n")
	b.WriteString(pyHarness)

	// Class-based design problems: operations/values format.
	if ops, ok := opsFromInput(input); ok {
		writeOpsRunner(&b, input, ops, tc.Expected)
		return b.String()
	}

	funcSignatures := extractAllFunctionSignatures(problem.StarterCode)
	if len(funcSignatures) == 0 {
		return userCode + "\n# Could not extract function name\n"
	}
	funcName := funcSignatures[0].Name
	funcParams := funcSignatures[0].Params

	// Assign inputs in sorted key order (deterministic).
	keys := make([]string, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if !pyIdentRe.MatchString(k) {
			continue
		}
		b.WriteString(k + " = " + pyLoads(input[k]) + "\n")
	}

	// Structure conversions, driven by category and parameter names.
	switch problem.Category {
	case "linked-list":
		for _, k := range keys {
			if !linkedListParams[k] {
				continue
			}
			if problem.ID == "copy-list-with-random-pointer" {
				b.WriteString(k + " = __dsa_build_random_list__(" + k + ")\n")
			} else {
				b.WriteString(k + " = __dsa_build_list__(" + k + ")\n")
			}
		}
		if _, ok := input["lists"]; ok {
			b.WriteString("lists = [__dsa_build_list__(x) for x in lists]\n")
		}
		if _, ok := input["pos"]; ok {
			b.WriteString("head = __dsa_make_cycle__(head, pos)\n")
		}
	case "trees":
		// Convert roots first: scalar p/q lookups below need the built tree.
		for _, k := range keys {
			if k == "root" || k == "subRoot" {
				b.WriteString(k + " = __dsa_build_tree__(" + k + ")\n")
			}
		}
		for _, k := range keys {
			if k != "p" && k != "q" {
				continue
			}
			if !isScalar(input[k]) {
				b.WriteString(k + " = __dsa_build_tree__(" + k + ")\n")
			} else {
				// e.g. lowest-common-ancestor: p/q given as values, solution needs nodes
				b.WriteString(k + " = __dsa_find_node__(root, " + k + ")\n")
			}
		}
	}
	// LeetCode-style solutions reference a `Node` class for these problems.
	switch problem.ID {
	case "clone-graph":
		b.WriteString("Node = GraphNode\n")
		b.WriteString("node = __dsa_build_graph__(adjList)\n")
	case "copy-list-with-random-pointer":
		b.WriteString("Node = ListNode\n")
	}

	callArgs := funcParams
	if len(callArgs) == 0 {
		callArgs = keys
	}

	if len(funcSignatures) == 2 && isEncodeDecodePair(funcSignatures) {
		encodeFn, decodeFn := getEncodeDecodeFunctions(funcSignatures)
		// Codec-style problems define the pair as methods on a class.
		if class := extractClassName(problem.StarterCode); class != "" {
			b.WriteString("__dsa_codec__ = " + class + "()\n")
			encodeFn = "__dsa_codec__." + encodeFn
			decodeFn = "__dsa_codec__." + decodeFn
		}
		b.WriteString("__dsa_enc__ = " + encodeFn + "(" + strings.Join(callArgs, ", ") + ")\n")
		b.WriteString("__dsa_result__ = " + decodeFn + "(__dsa_enc__)\n")
	} else {
		b.WriteString("__dsa_result__ = " + funcName + "(" + strings.Join(callArgs, ", ") + ")\n")
		if isEmptySlice(tc.Expected) {
			// e.g. removing the only node of a linked list: None means empty
			b.WriteString("if __dsa_result__ is None:\n    __dsa_result__ = []\n")
		} else if tc.Expected != nil && len(callArgs) > 0 {
			// In-place problems (reorder-list, rotate-image, ...) return None;
			// grade the mutated first argument instead.
			b.WriteString("if __dsa_result__ is None:\n    __dsa_result__ = " + callArgs[0] + "\n")
		}
	}

	if isScalar(tc.Expected) {
		// e.g. LCA returns a node but the expected answer is its value
		b.WriteString("if isinstance(__dsa_result__, (ListNode, TreeNode, GraphNode)):\n    __dsa_result__ = __dsa_result__.val\n")
	}

	serFn := "__dsa_ser__"
	if problem.ID == "copy-list-with-random-pointer" {
		serFn = "__dsa_ser_random__"
	}
	b.WriteString("print(" + strconv.Quote(resultSentinel) + " + json.dumps(" + serFn + "(__dsa_result__)))\n")
	return b.String()
}

// opsFromInput extracts the operations list for design problems.
func opsFromInput(input map[string]interface{}) ([]string, bool) {
	raw, ok := input["operations"]
	if !ok {
		return nil, false
	}
	b, _ := json.Marshal(raw)
	var ops []string
	if err := json.Unmarshal(b, &ops); err != nil || len(ops) == 0 {
		return nil, false
	}
	for _, op := range ops {
		if !pyIdentRe.MatchString(op) {
			return nil, false
		}
	}
	return ops, true
}

// writeOpsRunner emits a statically-unrolled runner for class-based problems:
// ops[0] is the constructor, subsequent ops are method calls. Where the
// expected output is null the op is treated as void and its return value is
// discarded (LeetCode semantics: a pop() that returns a value still passes).
func writeOpsRunner(b *strings.Builder, input map[string]interface{}, ops []string, expected interface{}) {
	values, ok := input["values"]
	if !ok {
		values = input["args"]
	}

	var expectedList []interface{}
	if eb, err := json.Marshal(expected); err == nil {
		json.Unmarshal(eb, &expectedList)
	}

	b.WriteString("__dsa_vals__ = " + pyLoads(values) + "\n")
	b.WriteString("__dsa_vals__ = [v if isinstance(v, list) else [v] for v in __dsa_vals__]\n")
	b.WriteString("__dsa_obj__ = " + ops[0] + "(*__dsa_vals__[0])\n")
	b.WriteString("__dsa_out__ = [None]\n")
	for i, op := range ops[1:] {
		call := fmt.Sprintf("__dsa_obj__.%s(*__dsa_vals__[%d])", op, i+1)
		if i+1 < len(expectedList) && expectedList[i+1] == nil {
			b.WriteString(call + "\n__dsa_out__.append(None)\n")
		} else {
			b.WriteString("__dsa_out__.append(__dsa_ser__(" + call + "))\n")
		}
	}
	b.WriteString("print(" + strconv.Quote(resultSentinel) + " + json.dumps(__dsa_out__))\n")
}

// FunctionSignature holds extracted function name and parameters
type FunctionSignature struct {
	Name   string
	Params []string
}

// isEncodeDecodePair checks if the two functions form an encode/decode pair
func isEncodeDecodePair(sigs []FunctionSignature) bool {
	if len(sigs) != 2 {
		return false
	}
	names := []string{strings.ToLower(sigs[0].Name), strings.ToLower(sigs[1].Name)}
	hasEncode := false
	hasDecode := false
	for _, n := range names {
		if isDecodeName(n) {
			hasDecode = true
		} else if strings.Contains(n, "encode") || strings.Contains(n, "serialize") {
			hasEncode = true
		}
	}
	return hasEncode && hasDecode
}

// isDecodeName must be checked before the encode test: "deserialize"
// contains "serialize" as a substring.
func isDecodeName(name string) bool {
	return strings.Contains(name, "decode") || strings.Contains(name, "deserialize")
}

// getEncodeDecodeFunctions returns the encode and decode function names in order
func getEncodeDecodeFunctions(sigs []FunctionSignature) (string, string) {
	var encodeFn, decodeFn string
	for _, sig := range sigs {
		if isDecodeName(strings.ToLower(sig.Name)) {
			decodeFn = sig.Name
		} else {
			encodeFn = sig.Name
		}
	}
	return encodeFn, decodeFn
}

// extractAllFunctionSignatures extracts all function signatures from starter code
func extractAllFunctionSignatures(starterCode string) []FunctionSignature {
	signatures := make([]FunctionSignature, 0)
	lines := strings.Split(starterCode, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "def ") {
			// Extract: def funcName(param1, param2, ...):
			line = strings.TrimPrefix(line, "def ")
			parenStart := strings.Index(line, "(")
			parenEnd := strings.Index(line, ")")
			if parenStart == -1 || parenEnd == -1 {
				continue
			}
			funcName := strings.TrimSpace(line[:parenStart])
			paramsStr := line[parenStart+1 : parenEnd]

			// Parse parameters
			params := make([]string, 0)
			if paramsStr != "" {
				for _, p := range strings.Split(paramsStr, ",") {
					p = strings.TrimSpace(p)
					// Remove type hints if present
					if colonIdx := strings.Index(p, ":"); colonIdx != -1 {
						p = strings.TrimSpace(p[:colonIdx])
					}
					// Remove default values if present
					if eqIdx := strings.Index(p, "="); eqIdx != -1 {
						p = strings.TrimSpace(p[:eqIdx])
					}
					if p != "" && p != "self" {
						params = append(params, p)
					}
				}
			}
			signatures = append(signatures, FunctionSignature{Name: funcName, Params: params})
		}
	}
	return signatures
}

// extractResultPayload pulls the JSON result emitted by the test wrapper out
// of the sandbox stdout (last sentinel line wins). Returns ok=false if the
// wrapper never printed a result (e.g. the code crashed before finishing).
func extractResultPayload(output string) (string, bool) {
	payload := ""
	found := false
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, resultSentinel) {
			payload = strings.TrimPrefix(line, resultSentinel)
			found = true
		}
	}
	return payload, found
}

// stripSentinel makes sandbox output user-presentable: the harness result line
// is shown as a plain value instead of the internal sentinel prefix.
func stripSentinel(output string) string {
	return strings.ReplaceAll(output, resultSentinel, "")
}

// compareOutput compares the wrapper-emitted JSON result with the expected value.
func compareOutput(output string, expected interface{}, orderIndependent bool) bool {
	payload, ok := extractResultPayload(output)
	if !ok {
		return false
	}

	var actualValue interface{}
	if err := json.Unmarshal([]byte(payload), &actualValue); err != nil {
		return false
	}

	if orderIndependent {
		return compareOrderIndependent(actualValue, expected)
	}

	// Round-trip expected through JSON so Go types ([]int vs []interface{})
	// and numeric representations (2 vs 2.0) compare canonically.
	var expectedValue interface{}
	expectedJSON, _ := json.Marshal(expected)
	if err := json.Unmarshal(expectedJSON, &expectedValue); err != nil {
		return false
	}
	actualJSON, _ := json.Marshal(actualValue)
	canonicalExpected, _ := json.Marshal(expectedValue)
	return string(actualJSON) == string(canonicalExpected)
}

// compareOrderIndependent compares two values regardless of order.
// Works for arrays of arrays (like group anagrams) and simple arrays.
func compareOrderIndependent(actual, expected interface{}) bool {
	// Convert expected to interface{} via JSON round-trip to normalize types
	// This handles Go types like [][]string -> []interface{} containing []interface{}
	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		return false
	}
	var expectedNormalized interface{}
	if err := json.Unmarshal(expectedJSON, &expectedNormalized); err != nil {
		return false
	}

	// Convert both to normalized form
	actualNorm := normalizeForComparison(actual)
	expectedNorm := normalizeForComparison(expectedNormalized)

	return actualNorm == expectedNorm
}

// normalizeForComparison converts a value to a canonical string form for comparison.
// For arrays: sorts elements. For arrays of arrays: sorts inner arrays, then sorts outer.
func normalizeForComparison(v interface{}) string {
	switch val := v.(type) {
	case []interface{}:
		// Check if it's an array of arrays
		if len(val) > 0 {
			if _, ok := val[0].([]interface{}); ok {
				// Array of arrays - sort each inner array, then sort outer by canonical form
				innerStrs := make([]string, len(val))
				for i, inner := range val {
					innerStrs[i] = normalizeForComparison(inner)
				}
				sort.Strings(innerStrs)
				return "[" + strings.Join(innerStrs, ",") + "]"
			}
		}
		// Simple array - convert elements to strings and sort
		strs := make([]string, len(val))
		for i, elem := range val {
			strs[i] = normalizeElement(elem)
		}
		sort.Strings(strs)
		return "[" + strings.Join(strs, ",") + "]"
	default:
		return normalizeElement(v)
	}
}

// normalizeElement converts a single element to its canonical string form.
func normalizeElement(v interface{}) string {
	switch val := v.(type) {
	case string:
		return fmt.Sprintf("%q", val)
	case float64:
		// JSON numbers are float64
		if val == float64(int(val)) {
			return fmt.Sprintf("%d", int(val))
		}
		return fmt.Sprintf("%g", val)
	case bool:
		return fmt.Sprintf("%t", val)
	case nil:
		return "null"
	case []interface{}:
		return normalizeForComparison(val)
	default:
		b, _ := json.Marshal(val)
		return string(b)
	}
}

