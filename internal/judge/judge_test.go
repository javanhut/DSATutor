package judge

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"dsatutor/internal/practice"
)

// TestReferenceSolutions runs every problem's solution, and every alternative
// that ships code, through the same judge a learner's Submit uses. A failure
// here means a learner could write the model answer and still be told it's
// wrong.
//
// Filter with -run, e.g. -run 'TestReferenceSolutions/graphs/'.
// Set DSATUTOR_JUDGE_VERBOSE=1 to print each failing case.
func TestReferenceSolutions(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	verbose := os.Getenv("DSATUTOR_JUDGE_VERBOSE") != ""
	for _, p := range practice.EmbeddedProblems {
		p := p
		t.Run(p.Category+"/"+p.ID, func(t *testing.T) {
			t.Parallel()
			codes := []struct{ name, code string }{{"solution", p.Solution.Code}}
			for _, alt := range p.Solution.Alternatives {
				if strings.TrimSpace(alt.Code) != "" {
					codes = append(codes, struct{ name, code string }{"alternative " + alt.Name, alt.Code})
				}
			}
			for _, c := range codes {
				for i := range p.TestCases {
					tc := &p.TestCases[i]
					res, err := Run(p, c.code, tc, Options{})
					if err != nil {
						t.Fatalf("%s, case %d: %v", c.name, i+1, err)
					}
					if !res.Passed {
						msg := fmt.Sprintf("%s fails case %d (%s): expected %v, got %v", c.name, i+1, tc.Note, jsonish(tc.Expected), jsonish(res.Actual))
						if res.Error != "" {
							msg += "; error: " + res.Error
						}
						if res.Message != "" {
							msg += "; " + res.Message
						}
						if verbose && res.Output != "" {
							msg += "\nprinted: " + res.Output
						}
						t.Error(msg)
					}
				}
			}
		})
	}
}

// TestJudgeRejectsWrongAnswers makes sure the comparison isn't so loose that
// anything passes.
func TestJudgeRejectsWrongAnswers(t *testing.T) {
	cases := []struct {
		actual, expected interface{}
		mode             string
		want             bool
	}{
		{[]interface{}{1.0, 2.0}, []int{1, 2}, "", true},
		{[]interface{}{2.0, 1.0}, []int{1, 2}, "", false},
		{[]interface{}{2.0, 1.0}, []int{1, 2}, "unordered", true},
		{[]interface{}{[]interface{}{1.0, 0.0}}, [][]int{{0, 1}}, "unordered", false},
		{[]interface{}{[]interface{}{1.0, 0.0}}, [][]int{{0, 1}}, "unordered-nested", true},
		{2.0, 2.00000001, "", true},
		{"(())", "(())", "", true},
		{"bab", "aba", "", false},
		{nil, []int{}, "", false},
		{true, false, "", false},
	}
	for _, c := range cases {
		if got := Compare(c.actual, c.expected, c.mode); got != c.want {
			t.Errorf("Compare(%v, %v, %q) = %v, want %v", c.actual, c.expected, c.mode, got, c.want)
		}
	}
}

func TestSignature(t *testing.T) {
	name, params := Signature("from typing import List\n\ndef twoSum(nums: List[int], target: int) -> List[int]:\n    pass")
	if name != "twoSum" || strings.Join(params, ",") != "nums,target" {
		t.Fatalf("got %s %v", name, params)
	}
	name, params = Signature("class Solution:\n    def helper(self):\n        pass\n\ndef f(grid: Dict[str, int], k=3):\n    pass")
	if name != "f" || strings.Join(params, ",") != "grid,k" {
		t.Fatalf("got %s %v", name, params)
	}
}

func jsonish(v interface{}) string {
	return fmt.Sprintf("%v", normalize(v))
}
