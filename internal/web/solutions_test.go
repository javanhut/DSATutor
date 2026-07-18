package web

import (
	"testing"

	"dsatutor/internal/practice"
	"dsatutor/internal/sandbox"
)

// TestReferenceSolutionsPass runs every problem's reference solution through
// the real wrapper + sandbox + comparison path. If a reference solution can't
// pass its own visible test cases, users can't either.
func TestReferenceSolutionsPass(t *testing.T) {
	loader := practice.NewProblemLoader(practice.ProblemsFS, practice.ProblemsBasePath)
	if err := loader.Load(); err != nil {
		t.Fatalf("load problems: %v", err)
	}

	for _, p := range loader.GetAllProblems() {
		p := p
		t.Run(p.ID, func(t *testing.T) {
			t.Parallel()
			for i, tc := range p.TestCases {
				code := generateTestWrapper(p.Solution.Code, p, &tc)
				resp, err := sandbox.Execute(sandbox.ExecuteRequest{Code: code, Timeout: tc.Timeout})
				if err != nil {
					t.Fatalf("tc %d: execute error: %v", i, err)
				}
				if !resp.Success {
					t.Fatalf("tc %d: run failed: %s", i, resp.Error)
				}
				if tc.Expected != nil && !compareOutput(resp.Output, tc.Expected, tc.OrderIndependent) {
					payload, _ := extractResultPayload(resp.Output)
					t.Fatalf("tc %d: wrong answer\n  got:  %s\n  want: %v", i, payload, tc.Expected)
				}
			}
		})
	}
}
