package practice

import (
	"time"
)

// SubmitStatus indicates submission outcome.
type SubmitStatus string

const (
	StatusAccepted          SubmitStatus = "Accepted"
	StatusWrongAnswer       SubmitStatus = "Wrong Answer"
	StatusTimeLimitExceeded SubmitStatus = "Time Limit Exceeded"
	StatusRuntimeError      SubmitStatus = "Runtime Error"
	StatusCompileError      SubmitStatus = "Compile Error"
	StatusPending           SubmitStatus = "Pending"
)

// Submission represents a user's code submission.
type Submission struct {
	ID          string           `json:"id"`
	ProblemID   string           `json:"problemId"`
	Code        string           `json:"code"`
	Language    string           `json:"language"`
	Status      SubmitStatus     `json:"status"`
	Results     []TestResult     `json:"results"`
	Complexity  *ComplexityResult `json:"complexity,omitempty"`
	SubmittedAt time.Time        `json:"submittedAt"`
	Runtime     int              `json:"runtime"` // Total execution time (ms)
}

// TestResult shows outcome for a single test case.
type TestResult struct {
	TestCaseID string      `json:"testCaseId"`
	Passed     bool        `json:"passed"`
	Input      interface{} `json:"input,omitempty"`
	Expected   interface{} `json:"expected,omitempty"`
	Actual     interface{} `json:"actual,omitempty"`
	Runtime    int         `json:"runtime"` // ms
	Error      string      `json:"error,omitempty"`
	Hidden     bool        `json:"hidden"`
}

// ComplexityResult from complexity analysis.
type ComplexityResult struct {
	EstimatedTime   string `json:"estimatedTime"`
	EstimatedSpace  string `json:"estimatedSpace"`
	MeetsRequired   bool   `json:"meetsRequired"`
	Analysis        string `json:"analysis"`
}

// ProgressStatus indicates user progress on a problem.
type ProgressStatus string

const (
	ProgressNotStarted ProgressStatus = "not_started"
	ProgressAttempted  ProgressStatus = "attempted"
	ProgressSolved     ProgressStatus = "solved"
)

// UserProgress tracks overall progress across all problems.
type UserProgress struct {
	SolvedProblems map[string]ProblemProgress `json:"solvedProblems"`
	TotalSolved    int                        `json:"totalSolved"`
	EasySolved     int                        `json:"easySolved"`
	MediumSolved   int                        `json:"mediumSolved"`
	HardSolved     int                        `json:"hardSolved"`
	LastActivity   time.Time                  `json:"lastActivity"`
}

// ProblemProgress tracks progress on a single problem.
type ProblemProgress struct {
	ProblemID      string         `json:"problemId"`
	Status         ProgressStatus `json:"status"`
	HintsUsed      int            `json:"hintsUsed"`
	SolutionViewed bool           `json:"solutionViewed"`
	BestRuntime    int            `json:"bestRuntime"`
	Attempts       int            `json:"attempts"`
	FirstSolvedAt  time.Time      `json:"firstSolvedAt,omitempty"`
	LastAttemptAt  time.Time      `json:"lastAttemptAt"`
}

// NewUserProgress creates an empty progress tracker.
func NewUserProgress() *UserProgress {
	return &UserProgress{
		SolvedProblems: make(map[string]ProblemProgress),
	}
}

// GetProblemProgress returns progress for a specific problem.
func (up *UserProgress) GetProblemProgress(problemID string) ProblemProgress {
	if prog, ok := up.SolvedProblems[problemID]; ok {
		return prog
	}
	return ProblemProgress{
		ProblemID: problemID,
		Status:    ProgressNotStarted,
	}
}

// UpdateProgress updates progress for a problem.
func (up *UserProgress) UpdateProgress(progress ProblemProgress) {
	up.SolvedProblems[progress.ProblemID] = progress
	up.LastActivity = time.Now()
	up.recalculateTotals()
}

// recalculateTotals updates the solved counts.
func (up *UserProgress) recalculateTotals() {
	up.TotalSolved = 0
	up.EasySolved = 0
	up.MediumSolved = 0
	up.HardSolved = 0

	for _, prog := range up.SolvedProblems {
		if prog.Status == ProgressSolved {
			up.TotalSolved++
		}
	}
}

// RunRequest represents a request to run code on a test case.
type RunRequest struct {
	ProblemID  string `json:"problemId"`
	Code       string `json:"code"`
	TestCaseID string `json:"testCaseId,omitempty"`
}

// RunResponse includes execution trace for visualization.
type RunResponse struct {
	Success    bool         `json:"success"`
	Passed     bool         `json:"passed"`
	Output     string       `json:"output"`
	Error      string       `json:"error,omitempty"`
	Steps      interface{}  `json:"steps"`
	Structures interface{}  `json:"structures"`
	Result     *TestResult  `json:"result,omitempty"`
	Runtime    int          `json:"runtime"`
}

// SubmitRequest represents a full submission for validation.
type SubmitRequest struct {
	ProblemID string `json:"problemId"`
	Code      string `json:"code"`
}

// SubmitResponse includes all test results.
type SubmitResponse struct {
	Status       SubmitStatus      `json:"status"`
	PassedCount  int               `json:"passedCount"`
	TotalCount   int               `json:"totalCount"`
	Results      []TestResult      `json:"results"`
	TotalRuntime int               `json:"totalRuntime"`
	Complexity   *ComplexityResult `json:"complexity,omitempty"`
}

// NewSubmitResponse creates a response with initial values.
func NewSubmitResponse() *SubmitResponse {
	return &SubmitResponse{
		Status:  StatusPending,
		Results: make([]TestResult, 0),
	}
}

// CalculateStatus determines the final status based on results.
func (sr *SubmitResponse) CalculateStatus() {
	if sr.PassedCount == sr.TotalCount && sr.TotalCount > 0 {
		sr.Status = StatusAccepted
	} else {
		// Check for specific error types
		for _, r := range sr.Results {
			if !r.Passed {
				if r.Error != "" {
					sr.Status = StatusRuntimeError
					return
				}
			}
		}
		sr.Status = StatusWrongAnswer
	}
}
