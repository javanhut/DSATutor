package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
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
		Output:     execResp.Output,
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
			result.Actual = execResp.Output
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
	parts := strings.Split(path, "/")

	if len(parts) < 2 {
		http.Error(w, "Problem ID and hint level required", http.StatusBadRequest)
		return
	}

	problemID := parts[0]
	level := 1
	if len(parts) > 1 {
		if _, err := json.Number(parts[1]).Int64(); err == nil {
			lvl, _ := json.Number(parts[1]).Int64()
			level = int(lvl)
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
		Output:     execResp.Output,
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

// generateTestWrapper creates code that runs with test input.
func generateTestWrapper(userCode string, problem *practice.Problem, tc *practice.TestCase) string {
	// Extract function name and parameters from starter code
	funcName, funcParams := extractFunctionSignature(problem.StarterCode)
	if funcName == "" {
		return userCode + "\n# Could not extract function name\n"
	}

	// Build input assignment from test case
	inputAssignments := ""
	callArgs := make([]string, 0)

	// Handle both map[string]interface{} and string input formats
	switch input := tc.Input.(type) {
	case map[string]interface{}:
		// Assign each input variable
		for key, value := range input {
			valueJSON, _ := json.Marshal(value)
			inputAssignments += key + " = " + string(valueJSON) + "\n"
		}
		// Use function parameters in correct order
		if len(funcParams) > 0 {
			callArgs = funcParams
		} else {
			// Fallback: use map keys
			for key := range input {
				callArgs = append(callArgs, key)
			}
		}
	case string:
		// String input format - parse and use directly
		inputAssignments = "__input__ = " + input + "\n"
		// Check if it's a single value or multiple
		if len(funcParams) == 1 {
			callArgs = []string{"__input__"}
		} else {
			// Multiple params - unpack the input
			callArgs = []string{"*__input__"}
		}
	default:
		// Fallback: marshal the input as JSON
		valueJSON, _ := json.Marshal(tc.Input)
		inputAssignments = "__input__ = " + string(valueJSON) + "\n"
		if len(funcParams) == 1 {
			callArgs = []string{"__input__"}
		} else {
			callArgs = []string{"*__input__"}
		}
	}

	// Generate wrapper code
	wrapper := userCode + "\n\n# Test execution\n" + inputAssignments
	wrapper += "\n__result__ = " + funcName + "(" + strings.Join(callArgs, ", ") + ")\n"
	wrapper += "print(__result__)\n"

	return wrapper
}

// extractFunctionSignature extracts the function name and parameter names from starter code.
func extractFunctionSignature(starterCode string) (string, []string) {
	lines := strings.Split(starterCode, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "def ") {
			// Extract: def funcName(param1, param2, ...):
			// Remove "def " prefix
			line = strings.TrimPrefix(line, "def ")
			// Find the parentheses
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
					// Remove type hints if present (e.g., "nums: List[int]" -> "nums")
					if colonIdx := strings.Index(p, ":"); colonIdx != -1 {
						p = strings.TrimSpace(p[:colonIdx])
					}
					// Remove default values if present (e.g., "x=5" -> "x")
					if eqIdx := strings.Index(p, "="); eqIdx != -1 {
						p = strings.TrimSpace(p[:eqIdx])
					}
					if p != "" && p != "self" {
						params = append(params, p)
					}
				}
			}
			return funcName, params
		}
	}
	return "", nil
}

// extractFunctionName extracts the main function name from starter code.
func extractFunctionName(starterCode string) string {
	lines := strings.Split(starterCode, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "def ") {
			// Extract function name
			parts := strings.Split(line, "(")
			if len(parts) > 0 {
				name := strings.TrimPrefix(parts[0], "def ")
				return strings.TrimSpace(name)
			}
		}
	}
	return ""
}

// compareOutput compares execution output with expected value.
func compareOutput(output string, expected interface{}, orderIndependent bool) bool {
	output = strings.TrimSpace(output)

	// Normalize Python output to JSON format
	// Python uses single quotes, True/False, None - convert to JSON equivalents
	normalized := normalizePythonOutput(output)

	// Try to parse normalized output as JSON
	var actualValue interface{}
	if err := json.Unmarshal([]byte(normalized), &actualValue); err != nil {
		// Output is not valid JSON, try direct comparison
		expectedJSON, _ := json.Marshal(expected)
		outputNoSpace := strings.ReplaceAll(strings.ReplaceAll(output, " ", ""), "\t", "")
		expectedNoSpace := strings.ReplaceAll(string(expectedJSON), " ", "")
		return outputNoSpace == expectedNoSpace
	}

	// For order-independent comparison (e.g., group anagrams, top k elements)
	if orderIndependent {
		return compareOrderIndependent(actualValue, expected)
	}

	// Compare JSON values (this normalizes spacing)
	actualJSON, _ := json.Marshal(actualValue)
	expectedJSON, _ := json.Marshal(expected)
	return string(actualJSON) == string(expectedJSON)
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

// normalizePythonOutput converts Python repr output to JSON format
func normalizePythonOutput(output string) string {
	// Replace Python booleans with JSON booleans
	output = strings.ReplaceAll(output, "True", "true")
	output = strings.ReplaceAll(output, "False", "false")
	output = strings.ReplaceAll(output, "None", "null")

	// Replace single quotes with double quotes (for strings)
	// This is a simple replacement - might need more sophisticated parsing for edge cases
	output = strings.ReplaceAll(output, "'", "\"")

	// Replace Python tuples () with arrays []
	output = strings.ReplaceAll(output, "(", "[")
	output = strings.ReplaceAll(output, ")", "]")

	return output
}
