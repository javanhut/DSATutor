package tutor

import "testing"

func TestTeachLessonCleanKeepsCheckAnswerThroughShuffle(t *testing.T) {
	for range 20 {
		l := TeachLesson{Topic: "heaps", Chunks: []TeachChunk{{
			Heading: "Push", Explanation: "Sift up.",
			Check: TeachCheck{Question: "Cost of push?", Correct: "O(log n)", Why: "height", Wrong: []TeachWrong{{"O(1)", "you forgot sifting"}, {"O(n)", "you scanned everything"}}},
		}}}
		l.Clean()
		c := l.Chunks[0].Check
		if c.Choices[c.Answer] != "O(log n)" || c.Feedback[c.Answer] != "" {
			t.Fatalf("answer %d points at %q", c.Answer, c.Choices[c.Answer])
		}
		for i, ch := range c.Choices {
			if ch == "O(n)" && c.Feedback[i] != "you scanned everything" {
				t.Fatalf("feedback for O(n) = %q", c.Feedback[i])
			}
		}
		if c.Correct != "" || c.Wrong != nil {
			t.Fatalf("model fields leaked: %+v", c)
		}
	}
}

func TestTeachLessonCleanDropsBadCheckAndRepairsBlank(t *testing.T) {
	l := TeachLesson{
		Chunks: []TeachChunk{{Heading: "A", Explanation: "B", Check: TeachCheck{Question: "Q", Correct: "x", Wrong: []TeachWrong{{Choice: "x"}}}}},
		Faded:  TeachFaded{Prompt: "p", Code: "```python\nlo = mid + 1\n```", Blank: "mid + 1"},
	}
	l.Clean()
	if l.Title != "" && l.Title != l.Topic {
		t.Errorf("title = %q", l.Title)
	}
	if len(l.Chunks[0].Check.Choices) != 0 {
		t.Errorf("check with no distinct wrong choice kept: %+v", l.Chunks[0].Check)
	}
	if l.Faded.Code != "lo = "+FadedBlank || l.Faded.Answer != "mid + 1" {
		t.Errorf("faded code = %q", l.Faded.Code)
	}

	l = TeachLesson{Faded: TeachFaded{Code: "a = 1", Blank: "zzz"}}
	l.Clean()
	if l.Faded.Code != "" {
		t.Errorf("faded without a blank kept: %+v", l.Faded)
	}
}

func TestCutBlankToleratesSpacing(t *testing.T) {
	code := "def f(a):\n    if a:\n        parent[x]=y\n    return a"
	got, ok := cutBlank(code, "parent[x] = y")
	if !ok || got != "def f(a):\n    if a:\n        ____\n    return a" {
		t.Errorf("cutBlank = %q, %v", got, ok)
	}
	if _, ok := cutBlank(code, "nothing like it"); ok {
		t.Error("cut a blank that isn't there")
	}
}
