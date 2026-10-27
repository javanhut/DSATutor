// Package practice provides LeetCode-style problem management and validation.
package practice

import (
	"time"
)

// Problem represents a LeetCode-style coding problem.
//
// Content follows docs/pedagogy.md: problems sit in a category ladder (Order),
// hints fade from a question to a code skeleton, and the solution explains the
// brute force, the insight that improves it, and why the result is correct.
type Problem struct {
	ID              string     `json:"id"`
	Number          int        `json:"number"` // LeetCode problem number
	Order           int        `json:"order"`  // position in its category's ladder, from 1
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
	// StarterCodeOxigen overrides the Oxigen starter OxigenStarter derives
	// from StarterCode.
	StarterCodeOxigen string      `json:"starterCodeOxigen,omitempty"`
	Harness           Harness     `json:"harness"`
	Hints             []Hint      `json:"hints"`
	Solution          Solution    `json:"solution"`
	SelfCheck         []SelfCheck `json:"selfCheck,omitempty"`
	CreatedAt         time.Time   `json:"createdAt,omitempty"`
	UpdatedAt         time.Time   `json:"updatedAt,omitempty"`
}

// Harness tells the judge how to call the learner's code. The zero value calls
// the first top-level function in StarterCode with the test inputs as
// arguments and compares the return value exactly.
type Harness struct {
	// Mode is "" (plain function), "design" (inputs are operations/values and
	// the first operation names the class) or "roundtrip" (Function names
	// "encode,decode" and the answer is decode(encode(input))).
	Mode      string `json:"mode,omitempty"`
	Function  string `json:"function,omitempty"`  // overrides the name taken from StarterCode
	ClassName string `json:"className,omitempty"` // roundtrip methods live on this class

	// ParamTypes converts JSON inputs before the call: ListNode, ListNode[],
	// TreeNode, GraphNode (1-indexed adjacency list), RandomList ([[val,
	// randomIndex]]), CycleList:<posParam>, TreeNodeRef:<treeParam>.
	ParamTypes map[string]string `json:"paramTypes,omitempty"`
	// ReturnType converts the answer: ListNode or TreeNode (so that returning
	// None reads as the empty list/tree []), GraphNode, RandomList, or
	// TreeNodeVal (a node, compared by its value). Node answers are turned
	// into lists automatically, but set ReturnType whenever one is returned.
	ReturnType string `json:"returnType,omitempty"`
	// InPlace names the parameter whose final value is the answer, for
	// functions that modify their input and return nothing.
	InPlace string `json:"inPlace,omitempty"`
	// Prelude is trusted Python run before the learner's code, e.g. the Node
	// class LeetCode provides for Clone Graph.
	Prelude string `json:"prelude,omitempty"`

	// Compare is "exact" (default), "unordered" (outer list in any order) or
	// "unordered-nested" (inner lists in any order as well).
	Compare string `json:"compare,omitempty"`
	// Validator is trusted Python defining check(raw, args, result, expected)
	// for problems with many correct answers. It returns a bool or a
	// (bool, message) tuple and replaces Compare when set.
	Validator string `json:"validator,omitempty"`
}

// SelfCheck is a question the learner answers after solving, to explain the
// solution in their own words before reading the model answer.
type SelfCheck struct {
	Prompt string `json:"prompt"`
	Answer string `json:"answer"`
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
	OrderIndependent bool        `json:"orderIndependent"`  // legacy: same as Harness.Compare "unordered-nested"
	Note             string      `json:"note,omitempty"`    // what this case checks, e.g. "empty input"
}

// Hint represents a progressive hint for the problem. Hints fade support step
// by step: 1 "nudge" (a question that points at the key observation),
// 2 "insight" (the observation itself), 3 "plan" (subgoal-labelled steps),
// 4 "skeleton" (code with the key lines left blank).
type Hint struct {
	Level      int    `json:"level"`
	Type       string `json:"type"` // nudge, insight, plan, skeleton
	Content    string `json:"content"`
	ChapterRef int    `json:"chapterRef,omitempty"`
}

// Solution contains the full solution and walkthrough.
type Solution struct {
	Code        string `json:"code"`
	Explanation string `json:"explanation"` // one-paragraph summary of the approach

	BruteForce     string   `json:"bruteForce"`     // the obvious approach and why it is too slow
	KeyInsight     string   `json:"keyInsight"`     // the observation that gets from brute force to optimal
	WhyItWorks     string   `json:"whyItWorks"`     // invariant or correctness argument
	CommonMistakes []string `json:"commonMistakes"` // bugs learners actually make here

	TimeComplexity  string            `json:"timeComplexity"`
	SpaceComplexity string            `json:"spaceComplexity"`
	Walkthrough     []WalkthroughStep `json:"walkthrough"`
	Alternatives    []AltSolution     `json:"alternatives,omitempty"`
}

// WalkthroughStep explains one part of the solution. Title is a subgoal label
// ("Grow the window", "Record the answer") that names what the lines achieve.
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
