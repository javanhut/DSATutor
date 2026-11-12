package tutor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"dsatutor/internal/practice"
)

func TestInterviewerPromptHidesSolutionAndKeepsClock(t *testing.T) {
	p := &practice.Problem{
		Title:       "Two Sum",
		Description: "Find two indices.",
		Solution:    practice.Solution{Code: "def SECRET(): pass", KeyInsight: "complement lookup", TimeComplexity: "O(n)"},
		SelfCheck:   []practice.SelfCheck{{Prompt: "What if the input is sorted?"}},
	}
	ctx := Context{View: "problem", Code: "x = 1", Interview: &Interview{ElapsedSeconds: 30 * 60, DurationSeconds: 45 * 60, Level: "L5"}}
	got := InterviewerPrompt(ctx, p, nil)
	for _, want := range []string{"L5 software engineer", "30 minutes elapsed, 15 minutes left of 45", "complement lookup", "What if the input is sorted?", "  1 | x = 1"} {
		if !strings.Contains(got, want) {
			t.Errorf("prompt lacks %q", want)
		}
	}
	if strings.Contains(got, "SECRET") {
		t.Error("interviewer prompt contains the solution code")
	}
}

func TestInterviewWorkspaceOnlyHighlights(t *testing.T) {
	ws := NewWorkspace(Context{
		View: "problem", Code: "a", TestCount: 3, Interview: &Interview{},
		Targets: []Target{{ID: "code", Label: "doc", Lines: true, Writable: true}},
	})
	tools := ws.Tools()
	if len(tools) != 1 || tools[0].Function.Name != "highlight" {
		t.Fatalf("interview tools = %+v, want only highlight", tools)
	}
	if _, act := ws.Execute(call("write_code", map[string]any{"mode": "replace_all", "code": "x"})); act != nil {
		t.Error("an interviewer edited the candidate's code")
	}
}

func TestScorecardNormalize(t *testing.T) {
	s := Scorecard{Verdict: "maybe", Ratings: map[string]Rating{"algorithms": {Score: 9}, "coding": {Score: 0}}, HintsGiven: -2}
	s.normalize()
	if s.Ratings["algorithms"].Score != 4 || s.Ratings["coding"].Score != 1 || s.Ratings["communication"].Score != 1 {
		t.Errorf("ratings not clamped: %+v", s.Ratings)
	}
	if s.HintsGiven != 0 {
		t.Errorf("hints = %d", s.HintsGiven)
	}
	found := false
	for _, v := range verdicts {
		found = found || v == s.Verdict
	}
	if !found {
		t.Errorf("verdict %q is not a valid verdict", s.Verdict)
	}
	s = Scorecard{Verdict: "lean hire"}
	s.normalize()
	if s.Verdict != "Lean Hire" {
		t.Errorf("verdict = %q, want Lean Hire", s.Verdict)
	}
}

func TestScoreSendsSchemaAndParses(t *testing.T) {
	var body map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/tags":
			w.Write([]byte(`{"models":[{"name":"m:latest"}]}`))
		case "/api/show":
			w.Write([]byte(`{"capabilities":["completion","thinking"]}`))
		case "/api/chat":
			json.NewDecoder(r.Body).Decode(&body)
			card := `{"verdict":"Hire","summary":"ok","ratings":{"algorithms":{"score":3,"evidence":"a"},"coding":{"score":4,"evidence":"b"},"communication":{"score":3,"evidence":"c"},"problem_solving":{"score":3,"evidence":"d"}},"complexity":"O(n)","strengths":["s"],"improvements":["i"],"hints_given":1,"reached_optimal":true}`
			json.NewEncoder(w).Encode(map[string]any{"message": map[string]string{"content": card}, "done": true})
		}
	}))
	defer srv.Close()
	t.Setenv("DSATUTOR_OLLAMA_URL", srv.URL)
	t.Setenv("DSATUTOR_TUTOR_MODEL", "")

	p := &practice.Problem{Title: "Two Sum", Description: "d"}
	tests := &practice.SubmitResponse{Status: practice.StatusWrongAnswer, PassedCount: 3, TotalCount: 4,
		Results: []practice.TestResult{{Passed: false, Note: "duplicates"}}}
	card, model, err := New().Score(context.Background(), "", ScoreInput{
		Problem: p, Code: "pass", Transcript: []Message{{Role: "user", Content: "hello"}},
		Elapsed: 20 * time.Minute, Duration: 45 * time.Minute, Level: "L4", Tests: tests,
	})
	if err != nil {
		t.Fatal(err)
	}
	if model != "m:latest" || card.Verdict != "Hire" || card.Ratings["coding"].Score != 4 || card.HintsGiven != 1 {
		t.Errorf("card = %+v, model %s", card, model)
	}
	if body["format"] == nil || body["stream"] != false || body["think"] != false {
		t.Errorf("request = %v, want a schema, no streaming, no thinking", body)
	}
	msgs := body["messages"].([]any)
	user := msgs[1].(map[string]any)["content"].(string)
	for _, want := range []string{"Used 20 of 45 minutes", "Candidate: hello", "passed 3 of 4", "failed: duplicates"} {
		if !strings.Contains(user, want) {
			t.Errorf("scorer input lacks %q", want)
		}
	}
}
