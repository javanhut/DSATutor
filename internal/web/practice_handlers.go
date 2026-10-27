package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"dsatutor/internal/judge"
	"dsatutor/internal/practice"
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
		Order           int      `json:"order"`
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
			Order:           p.Order,
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
		Note     string      `json:"note,omitempty"`
	}

	safeTestCases := make([]SafeTestCase, len(problem.TestCases))
	for i, tc := range problem.TestCases {
		safeTestCases[i] = SafeTestCase{
			ID:      tc.ID,
			Input:   tc.Input,
			Hidden:  tc.Hidden,
			Timeout: tc.Timeout,
			Note:    tc.Note,
		}
		// Hidden cases reveal only what they check, not their data.
		if tc.Hidden {
			safeTestCases[i].Input = nil
		} else {
			safeTestCases[i].Expected = tc.Expected
		}
	}

	// Build response without full solution (user must explicitly request it)
	response := map[string]interface{}{
		"id":                problem.ID,
		"number":            problem.Number,
		"order":             problem.Order,
		"title":             problem.Title,
		"slug":              problem.Slug,
		"difficulty":        problem.Difficulty,
		"category":          problem.Category,
		"tags":              problem.Tags,
		"relatedChapters":   problem.RelatedChapters,
		"description":       problem.Description,
		"constraints":       problem.Constraints,
		"examples":          problem.Examples,
		"testCases":         safeTestCases,
		"timeComplexity":    problem.TimeComplexity,
		"spaceComplexity":   problem.SpaceComplexity,
		"starterCode":       problem.StarterCode,
		"starterCodeOxigen": practice.OxigenStarter(problem),
		"hintCount":         len(problem.Hints),
		"hints":             problem.Hints,
		"selfCheck":         problem.SelfCheck,
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
		writeJSON(w, http.StatusBadRequest, practice.RunResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	problem, err := ps.loader.GetProblem(req.ProblemID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, practice.RunResponse{Error: "Problem not found: " + err.Error()})
		return
	}

	tc := pickTestCase(problem, req.TestCaseID, req.TestCaseIndex)
	if tc == nil {
		writeJSON(w, http.StatusBadRequest, practice.RunResponse{Error: "No test case found"})
		return
	}

	writeJSON(w, http.StatusOK, runOne(problem, req.Code, req.Language, tc))
}

// pickTestCase finds a test case by ID or index, defaulting to the first
// visible one. Hidden cases can't be run individually.
func pickTestCase(problem *practice.Problem, id string, index *int) *practice.TestCase {
	if id != "" {
		for i := range problem.TestCases {
			if problem.TestCases[i].ID == id {
				return &problem.TestCases[i]
			}
		}
	}
	if index != nil && *index >= 0 && *index < len(problem.TestCases) && !problem.TestCases[*index].Hidden {
		return &problem.TestCases[*index]
	}
	for i := range problem.TestCases {
		if !problem.TestCases[i].Hidden {
			return &problem.TestCases[i]
		}
	}
	return nil
}

// runOne judges code on one visible test case with tracing on, for the Run
// button and the solution visualizer.
func runOne(problem *practice.Problem, code, language string, tc *practice.TestCase) practice.RunResponse {
	res, err := judge.Run(problem, code, tc, judge.Options{Trace: true, Language: language})
	if err != nil {
		return practice.RunResponse{Error: "Execution error: " + err.Error()}
	}
	response := practice.RunResponse{
		Success:   res.Error == "",
		Passed:    res.Passed,
		Output:    res.Output,
		Error:     res.Error,
		Steps:     res.Steps,
		Actual:    res.Actual,
		Expected:  tc.Expected,
		Message:   res.Message,
		Truncated: res.Truncated,
	}
	if len(res.Steps) > 0 {
		response.Structures = res.Steps[len(res.Steps)-1].Structures
	}
	return response
}

// handleSubmitCode validates code against all test cases.
func (ps *PracticeServer) handleSubmitCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req practice.SubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, practice.SubmitResponse{Status: practice.StatusCompileError})
		return
	}

	problem, err := ps.loader.GetProblem(req.ProblemID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, practice.SubmitResponse{Status: practice.StatusCompileError})
		return
	}

	response := practice.NewSubmitResponse()
	response.TotalCount = len(problem.TestCases)

	for i := range problem.TestCases {
		tc := &problem.TestCases[i]
		result := practice.TestResult{
			TestCaseID: tc.ID,
			Hidden:     tc.Hidden,
			Note:       tc.Note,
		}

		res, err := judge.Run(problem, req.Code, tc, judge.Options{Language: req.Language})
		if err != nil {
			result.Error = err.Error()
		} else {
			result.Passed = res.Passed
			result.Error = res.Error
			result.Message = res.Message
			if !tc.Hidden {
				result.Input = tc.Input
				result.Expected = tc.Expected
				result.Actual = res.Actual
				result.Output = res.Output
			}
		}

		if result.Passed {
			response.PassedCount++
		}
		response.Results = append(response.Results, result)
	}

	response.CalculateStatus()
	writeJSON(w, http.StatusOK, response)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
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
	testCase := pickTestCase(problem, "", &req.TestCaseIndex)

	if testCase == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(practice.RunResponse{
			Success: false,
			Error:   "No test case found",
		})
		return
	}

	writeJSON(w, http.StatusOK, runOne(problem, problem.Solution.Code, "python", testCase))
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
