package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"dsatutor/internal/practice"
)

func TestTeachLessonFromFakeOllama(t *testing.T) {
	lesson := `{"title":"Monotonic stack","warmup":"How would you find the next greater element?","big_idea":"Keep a stack in decreasing order.",
"chunks":[{"heading":"The invariant","explanation":"Values on the stack decrease.","check":{"question":"What pops?","correct":"Smaller values","why":"They found their answer.","wrong":[{"choice":"Larger values","misconception":"You flipped it."},{"choice":"Nothing","misconception":"Something must pop."}]}},
{"heading":"Cost","explanation":"Each value is pushed and popped once.","check":{"question":"Total cost?","correct":"O(n)","why":"Amortized.","wrong":[{"choice":"O(n^2)","misconception":"Pops are paid for."}]}}],
"worked_example":{"problem":"[2,1,3]","steps":[{"do":"push 2","why":"empty"},{"do":"push 1","why":"smaller"},{"do":"3 pops both","why":"greater"}]},
"complexity":"O(n) time","faded":{"prompt":"Fill the loop condition","code":"while stack and stack[-1] < x:\n    stack.pop()","blank":"stack[-1] < x","hint":"compare with the top"},
"pitfalls":["Storing values instead of indices","Forgetting leftovers"],"explain_prompt":"Why is it O(n)?","explain_reference":"Each index is pushed and popped at most once.",
"flashcards":[{"front":"Invariant?","back":"Decreasing"},{"front":"Cost?","back":"O(n)"},{"front":"Store?","back":"Indices"}],"practice_keywords":["stack"]}`
	ollama := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/tags":
			w.Write([]byte(`{"models":[{"name":"m:latest"}]}`))
		case "/api/show":
			w.Write([]byte(`{"capabilities":["completion"]}`))
		case "/api/chat":
			var body struct {
				Format map[string]any `json:"format"`
			}
			json.NewDecoder(r.Body).Decode(&body)
			if body.Format == nil {
				t.Error("no schema sent")
			}
			out, _ := json.Marshal(map[string]any{"message": map[string]any{"content": lesson}, "done": true})
			w.Write(out)
		}
	}))
	defer ollama.Close()
	t.Setenv("DSATUTOR_OLLAMA_URL", ollama.URL)
	t.Setenv("DSATUTOR_TUTOR_MODEL", "")
	s := NewServer(nil)

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/teachme/lesson", strings.NewReader(`{"topic":"monotonic stack"}`)))
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d: %s", rec.Code, rec.Body)
	}
	var got struct {
		Lesson struct {
			Title string `json:"title"`
			Faded struct {
				Code string `json:"code"`
			} `json:"faded"`
			Chunks []any `json:"chunks"`
		} `json:"lesson"`
	}
	json.Unmarshal(rec.Body.Bytes(), &got)
	if got.Lesson.Title != "Monotonic stack" || len(got.Lesson.Chunks) != 2 {
		t.Errorf("lesson = %+v", got.Lesson)
	}
	if !strings.Contains(got.Lesson.Faded.Code, "____") {
		t.Errorf("faded code has no blank: %q", got.Lesson.Faded.Code)
	}
}

func TestTeachRejectsEmptyTopic(t *testing.T) {
	s := NewServer(nil)
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/api/teachme/lesson", strings.NewReader(`{"topic":" "}`)))
	if rec.Code != http.StatusBadRequest {
		t.Errorf("status %d", rec.Code)
	}
}

func TestTeachProblemsPrefersTheLearnersTopic(t *testing.T) {
	loader := practice.NewProblemLoader(practice.ProblemsFS, practice.ProblemsBasePath)
	if err := loader.Load(); err != nil {
		t.Fatal(err)
	}
	s := NewServerWithPractice(nil, loader)
	got := s.teachProblems("binary search on the answer", []string{"two pointers", "monotonic"}, 5)
	if len(got) == 0 || got[0].Category != "binary-search" {
		t.Errorf("first match = %+v", got)
	}
}
