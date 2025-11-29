// Package practice provides LeetCode-style problem management and validation.
package practice

import (
	"time"
)

// Problem represents a LeetCode-style coding problem.
type Problem struct {
	ID              string     `json:"id"`
	Number          int        `json:"number"`
	Title           string     `json:"title"`
	Slug            string     `json:"slug"`
	Difficulty      string     `json:"difficulty"` // Easy, Medium, Hard
	Category        string     `json:"category"`
	Tags            []string   `json:"tags"`
	RelatedChapters []int      `json:"relatedChapters"`
	Description     string     `json:"description"` // Markdown
	Constraints     []string   `json:"constraints"`
	Examples        []Example  `json:"examples"`
	TestCases       []TestCase `json:"testCases"`
	TimeComplexity  string     `json:"timeComplexity"`
	SpaceComplexity string     `json:"spaceComplexity"`
	StarterCode     string     `json:"starterCode"`
	Hints           []Hint     `json:"hints"`
	Solution        Solution   `json:"solution"`
	CreatedAt       time.Time  `json:"createdAt,omitempty"`
	UpdatedAt       time.Time  `json:"updatedAt,omitempty"`
}

// Example shows a sample input/output for the problem description.
type Example struct {
	ID          string `json:"id"`
	Input       string `json:"input"`
	Output      string `json:"output"`
	Explanation string `json:"explanation,omitempty"`
}

// TestCase is a machine-readable test for validation.
type TestCase struct {
	ID               string      `json:"id,omitempty"`
	Input            interface{} `json:"input"`
	Expected         interface{} `json:"expected"`
	Hidden           bool        `json:"hidden"`
	Timeout          int         `json:"timeout,omitempty"` // milliseconds
	OrderIndependent bool        `json:"orderIndependent"`  // Result can be in any order (e.g., group anagrams)
}

// Hint represents a progressive hint for the problem.
type Hint struct {
	Level      int    `json:"level"` // 1, 2, 3 (increasing detail)
	Type       string `json:"type"`  // approach, algorithm, code
	Content    string `json:"content"`
	ChapterRef int    `json:"chapterRef,omitempty"`
}

// Solution contains the full solution and walkthrough.
type Solution struct {
	Code            string            `json:"code"`
	Explanation     string            `json:"explanation"`
	TimeComplexity  string            `json:"timeComplexity"`
	SpaceComplexity string            `json:"spaceComplexity"`
	Walkthrough     []WalkthroughStep `json:"walkthrough"`
	Alternatives    []AltSolution     `json:"alternatives,omitempty"`
}

// WalkthroughStep explains one part of the solution.
type WalkthroughStep struct {
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
	CodeSnippet string `json:"codeSnippet"`
	LineStart   int    `json:"lineStart"`
	LineEnd     int    `json:"lineEnd"`
}

// AltSolution represents an alternative approach.
type AltSolution struct {
	Name            string `json:"name"`
	Code            string `json:"code"`
	TimeComplexity  string `json:"timeComplexity"`
	SpaceComplexity string `json:"spaceComplexity"`
	Explanation     string `json:"explanation"`
}

// ProblemIndex represents the master index of all problems.
type ProblemIndex struct {
	Version       string         `json:"version"`
	TotalProblems int            `json:"totalProblems"`
	LastUpdated   time.Time      `json:"lastUpdated"`
	Problems      []ProblemEntry `json:"problems"`
}

// ProblemEntry is a summary entry in the problem index.
type ProblemEntry struct {
	ID         string   `json:"id"`
	Number     int      `json:"number"`
	Title      string   `json:"title"`
	Difficulty string   `json:"difficulty"`
	Category   string   `json:"category"`
	Tags       []string `json:"tags"`
	Path       string   `json:"path"`
}

// GetVisibleTestCases returns only non-hidden test cases.
func (p *Problem) GetVisibleTestCases() []TestCase {
	visible := make([]TestCase, 0)
	for _, tc := range p.TestCases {
		if !tc.Hidden {
			visible = append(visible, tc)
		}
	}
	return visible
}

// GetHintByLevel returns the hint at the specified level.
func (p *Problem) GetHintByLevel(level int) *Hint {
	for _, h := range p.Hints {
		if h.Level == level {
			return &h
		}
	}
	return nil
}

// MaxHintLevel returns the highest hint level available.
func (p *Problem) MaxHintLevel() int {
	max := 0
	for _, h := range p.Hints {
		if h.Level > max {
			max = h.Level
		}
	}
	return max
}
