// Package contentcheck holds tests that hold the curriculum to the rules in
// docs/pedagogy.md. It has no code of its own; it exists so the tests can see
// both the chapters and the practice problems.
package contentcheck

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"dsatutor/internal/chapter"
	"dsatutor/internal/practice"
)

// Optional filters so one category can be checked while others are in progress:
// DSATUTOR_CATEGORY=graphs,advanced-graphs  DSATUTOR_CHAPTER=13,14
func categoryWanted(id string) bool {
	return wanted(os.Getenv("DSATUTOR_CATEGORY"), id)
}

func chapterWanted(n int) bool {
	return wanted(os.Getenv("DSATUTOR_CHAPTER"), fmt.Sprint(n))
}

func wanted(filter, v string) bool {
	if filter == "" {
		return true
	}
	for _, f := range strings.Split(filter, ",") {
		if strings.TrimSpace(f) == v {
			return true
		}
	}
	return false
}

var (
	hintTypes      = []string{"nudge", "insight", "plan", "skeleton"}
	constBigO      = regexp.MustCompile(`O\((\d+)\)`)
	difficulties   = map[string]bool{"Easy": true, "Medium": true, "Hard": true}
	compareModes   = map[string]bool{"": true, "exact": true, "unordered": true, "unordered-nested": true}
	harnessModes   = map[string]bool{"": true, "function": true, "design": true, "roundtrip": true}
	paramTypeNames = regexp.MustCompile(`^(ListNode|ListNode\[\]|TreeNode|GraphNode|RandomList|CycleList:\w+|TreeNodeRef:\w+)$`)
)

func chapters() map[int]chapter.Chapter {
	out := map[int]chapter.Chapter{}
	for _, ch := range chapter.DefaultChapters() {
		out[ch.Number] = ch
	}
	return out
}

func TestProblemContent(t *testing.T) {
	chs := chapters()
	ids := map[string]bool{}
	numbers := map[int]string{}
	orders := map[string]map[int]string{}
	categories := map[string]practice.Category{}
	for _, c := range practice.Blind75Categories {
		categories[c.ID] = c
	}

	for _, p := range practice.EmbeddedProblems {
		// Uniqueness is checked across everything, even when filtering.
		if ids[p.ID] {
			t.Errorf("%s: duplicate ID", p.ID)
		}
		ids[p.ID] = true
		if other, ok := numbers[p.Number]; ok {
			t.Errorf("%s: LeetCode number %d also used by %s", p.ID, p.Number, other)
		}
		numbers[p.Number] = p.ID
		if orders[p.Category] == nil {
			orders[p.Category] = map[int]string{}
		}
		if other, ok := orders[p.Category][p.Order]; ok {
			t.Errorf("%s: Order %d in %s also used by %s", p.ID, p.Order, p.Category, other)
		}
		orders[p.Category][p.Order] = p.ID

		if !categoryWanted(p.Category) {
			continue
		}
		p := p
		t.Run(p.Category+"/"+p.ID, func(t *testing.T) {
			checkProblem(t, p, categories, chs)
		})
	}

	for cat, byOrder := range orders {
		if !categoryWanted(cat) {
			continue
		}
		for i := 1; i <= len(byOrder); i++ {
			if _, ok := byOrder[i]; !ok {
				t.Errorf("category %s: Order values must run 1..%d without gaps (missing %d)", cat, len(byOrder), i)
			}
		}
	}
}

func checkProblem(t *testing.T, p *practice.Problem, categories map[string]practice.Category, chs map[int]chapter.Chapter) {
	cat, ok := categories[p.Category]
	if !ok {
		t.Fatalf("unknown category %q", p.Category)
	}
	if p.Number <= 0 || p.Order <= 0 {
		t.Errorf("Number (LeetCode #) and Order must be set")
	}
	if !difficulties[p.Difficulty] {
		t.Errorf("difficulty %q", p.Difficulty)
	}
	if p.Slug == "" || p.Title == "" || len(p.Tags) == 0 {
		t.Errorf("Slug, Title and Tags are required")
	}

	// Chapter links must exist and include the chapter that teaches the pattern.
	if len(p.RelatedChapters) == 0 {
		t.Errorf("RelatedChapters is empty")
	}
	for _, n := range p.RelatedChapters {
		if _, ok := chs[n]; !ok {
			t.Errorf("RelatedChapters has %d, which doesn't exist", n)
		}
	}
	if len(cat.ChapterRefs) > 0 && !containsInt(p.RelatedChapters, cat.ChapterRefs[0]) {
		t.Errorf("RelatedChapters %v must include chapter %d, which teaches %s", p.RelatedChapters, cat.ChapterRefs[0], cat.Name)
	}
	for _, h := range p.Hints {
		if h.ChapterRef != 0 {
			if _, ok := chs[h.ChapterRef]; !ok {
				t.Errorf("hint %d ChapterRef %d doesn't exist", h.Level, h.ChapterRef)
			}
		}
	}

	if len(strings.TrimSpace(p.Description)) < 80 {
		t.Errorf("Description is too short to be self-contained")
	}
	if len(p.Constraints) == 0 || len(p.Examples) == 0 {
		t.Errorf("needs Constraints and at least one Example")
	}
	for i, ex := range p.Examples {
		if ex.Input == "" || ex.Output == "" {
			t.Errorf("example %d needs Input and Output", i+1)
		}
	}

	checkTests(t, p)
	checkHarness(t, p)
	checkHints(t, p)
	checkSolution(t, p)
	checkComplexity(t, "TimeComplexity", p.TimeComplexity)
	checkComplexity(t, "SpaceComplexity", p.SpaceComplexity)

	if len(p.SelfCheck) < 2 {
		t.Errorf("needs at least 2 SelfCheck questions, got %d", len(p.SelfCheck))
	}
	for i, q := range p.SelfCheck {
		if !strings.Contains(q.Prompt, "?") || len(q.Answer) < 40 {
			t.Errorf("SelfCheck %d needs a question and a full answer", i+1)
		}
	}
}

func checkTests(t *testing.T, p *practice.Problem) {
	visible, hidden := 0, 0
	for i, tc := range p.TestCases {
		if _, ok := tc.Input.(map[string]interface{}); !ok {
			t.Errorf("test %d: Input must be map[string]interface{} of named arguments", i+1)
		}
		if tc.Hidden {
			hidden++
		} else {
			visible++
		}
		if tc.Note == "" {
			t.Errorf("test %d: Note must say what the case checks", i+1)
		}
	}
	if len(p.TestCases) < 6 || visible < 2 || hidden < 2 {
		t.Errorf("needs at least 6 tests with at least 2 visible and 2 hidden (has %d: %d visible, %d hidden)", len(p.TestCases), visible, hidden)
	}
	// If every expected answer is the same, `return <constant>` passes.
	if p.Harness.Mode != "design" && len(p.TestCases) > 1 {
		first := fmt.Sprint(p.TestCases[0].Expected)
		same := true
		for _, tc := range p.TestCases[1:] {
			if fmt.Sprint(tc.Expected) != first {
				same = false
				break
			}
		}
		if same {
			t.Errorf("every test expects %s, so returning a constant passes", first)
		}
	}
}

func checkHarness(t *testing.T, p *practice.Problem) {
	h := p.Harness
	if !harnessModes[h.Mode] {
		t.Errorf("Harness.Mode %q", h.Mode)
	}
	if !compareModes[h.Compare] {
		t.Errorf("Harness.Compare %q", h.Compare)
	}
	for name, kind := range h.ParamTypes {
		if !paramTypeNames.MatchString(kind) {
			t.Errorf("Harness.ParamTypes[%s] = %q is not a known conversion", name, kind)
		}
	}
	if h.Validator != "" && !strings.Contains(h.Validator, "def check(raw, args, result, expected)") {
		t.Errorf("Validator must define check(raw, args, result, expected)")
	}
	if h.Mode == "design" {
		for i, tc := range p.TestCases {
			in, _ := tc.Input.(map[string]interface{})
			if _, ok := in["operations"]; !ok {
				t.Errorf("design test %d needs operations and values", i+1)
			}
			if _, ok := in["values"]; !ok {
				t.Errorf("design test %d needs operations and values", i+1)
			}
		}
	}
	// Problems on linked lists and trees must convert their inputs.
	starter := strings.ToLower(p.StarterCode)
	needsTypes := strings.Contains(starter, "listnode") || strings.Contains(starter, "treenode") ||
		strings.Contains(starter, "(head") || strings.Contains(starter, "(root")
	if needsTypes && len(h.ParamTypes) == 0 && h.ReturnType == "" && h.Mode != "design" {
		t.Errorf("starter code uses nodes but Harness has neither ParamTypes nor ReturnType")
	}
}

func checkHints(t *testing.T, p *practice.Problem) {
	if len(p.Hints) != len(hintTypes) {
		t.Errorf("needs exactly %d hints (nudge, insight, plan, skeleton), got %d", len(hintTypes), len(p.Hints))
		return
	}
	for i, h := range p.Hints {
		if h.Level != i+1 || h.Type != hintTypes[i] {
			t.Errorf("hint %d must be level %d type %q, got level %d type %q", i+1, i+1, hintTypes[i], h.Level, h.Type)
		}
		if strings.TrimSpace(h.Content) == "" {
			t.Errorf("hint %d is empty", i+1)
		}
	}
	if !strings.Contains(p.Hints[0].Content, "?") {
		t.Errorf("the nudge hint should be a question, so the learner does the thinking")
	}
	if !strings.Contains(p.Hints[3].Content, "____") {
		t.Errorf("the skeleton hint should be code with ____ blanks for the key lines")
	}
	sol := normalizeCode(p.Solution.Code)
	if normalizeCode(p.Hints[3].Content) == sol {
		t.Errorf("the skeleton hint must not be the full solution")
	}
}

func checkSolution(t *testing.T, p *practice.Problem) {
	s := p.Solution
	if strings.TrimSpace(s.Code) == "" {
		t.Fatalf("Solution.Code is empty")
	}
	minLen := map[string]int{"Explanation": 60, "BruteForce": 60, "KeyInsight": 60, "WhyItWorks": 150}
	got := map[string]string{"Explanation": s.Explanation, "BruteForce": s.BruteForce, "KeyInsight": s.KeyInsight, "WhyItWorks": s.WhyItWorks}
	for field, n := range minLen {
		if len(strings.TrimSpace(got[field])) < n {
			t.Errorf("Solution.%s needs at least %d characters of real explanation (has %d)", field, n, len(strings.TrimSpace(got[field])))
		}
	}
	if len(s.CommonMistakes) < 2 {
		t.Errorf("needs at least 2 CommonMistakes")
	}
	checkComplexity(t, "Solution.TimeComplexity", s.TimeComplexity)
	checkComplexity(t, "Solution.SpaceComplexity", s.SpaceComplexity)

	hasCode := false
	for _, alt := range s.Alternatives {
		if strings.TrimSpace(alt.Code) != "" {
			hasCode = true
		}
		if alt.Name == "" || alt.Explanation == "" || alt.TimeComplexity == "" {
			t.Errorf("alternative %q needs Name, Explanation and complexity", alt.Name)
		}
	}
	if !hasCode {
		t.Errorf("needs at least one Alternative with runnable Code (usually the brute force)")
	}

	lines := strings.Split(s.Code, "\n")
	if len(s.Walkthrough) < 3 {
		t.Errorf("needs a walkthrough of at least 3 subgoal-labelled steps, got %d", len(s.Walkthrough))
	}
	for i, w := range s.Walkthrough {
		if w.Title == "" || len(w.Explanation) < 30 {
			t.Errorf("walkthrough step %d needs a subgoal Title and an Explanation", i+1)
			continue
		}
		if w.LineStart < 1 || w.LineEnd < w.LineStart || w.LineEnd > len(lines) {
			t.Errorf("walkthrough step %d (%s): lines %d-%d are outside the %d-line solution", i+1, w.Title, w.LineStart, w.LineEnd, len(lines))
			continue
		}
		span := normalizeCode(strings.Join(lines[w.LineStart-1:w.LineEnd], "\n"))
		for _, snippetLine := range strings.Split(w.CodeSnippet, "\n") {
			sl := normalizeCode(snippetLine)
			if sl != "" && !strings.Contains(span, sl) {
				t.Errorf("walkthrough step %d (%s): snippet line %q is not within lines %d-%d", i+1, w.Title, strings.TrimSpace(snippetLine), w.LineStart, w.LineEnd)
				break
			}
		}
	}
}

func checkComplexity(t *testing.T, field, v string) {
	if strings.TrimSpace(v) == "" {
		t.Errorf("%s is empty", field)
		return
	}
	for _, m := range constBigO.FindAllStringSubmatch(v, -1) {
		if m[1] != "1" {
			t.Errorf("%s %q: constants aren't written inside Big O; write O(1) and explain the bound in prose", field, v)
		}
	}
}

func normalizeCode(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func containsInt(list []int, v int) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}

// TestChapterContent checks every chapter has the full learning sequence.
func TestChapterContent(t *testing.T) {
	chs := chapter.DefaultChapters()
	animations := registeredAnimations(t)
	categories := map[string]bool{}
	for _, c := range practice.Blind75Categories {
		categories[c.ID] = true
	}
	seenCategory := map[string]int{}

	for i, ch := range chs {
		if ch.Number != i+1 {
			t.Errorf("chapter %q has Number %d, want %d (chapters are numbered in order)", ch.Slug, ch.Number, i+1)
		}
		if ch.Category != "" {
			if !categories[ch.Category] {
				t.Errorf("chapter %d: unknown category %q", ch.Number, ch.Category)
			}
			if prev, ok := seenCategory[ch.Category]; ok {
				t.Errorf("chapters %d and %d both claim category %s", prev, ch.Number, ch.Category)
			}
			seenCategory[ch.Category] = ch.Number
		}
		if !chapterWanted(ch.Number) {
			continue
		}
		ch := ch
		t.Run(fmt.Sprintf("%02d-%s", ch.Number, ch.Slug), func(t *testing.T) {
			checkChapter(t, ch, animations)
		})
	}
	for _, c := range practice.Blind75Categories {
		if _, ok := seenCategory[c.ID]; !ok {
			t.Errorf("category %s has no chapter teaching it", c.ID)
		}
		if len(c.ChapterRefs) == 0 || seenCategory[c.ID] != c.ChapterRefs[0] {
			t.Errorf("category %s: ChapterRefs[0] must be chapter %d", c.ID, seenCategory[c.ID])
		}
	}
}

func checkChapter(t *testing.T, ch chapter.Chapter, animations map[string]bool) {
	for _, p := range ch.Prerequisites {
		if p >= ch.Number || p < 1 {
			t.Errorf("prerequisite %d must be an earlier chapter", p)
		}
	}
	if len(ch.Summary) < 80 || len(ch.Motivation) < 200 {
		t.Errorf("needs a Summary (80+ chars) and a concrete Motivation (200+ chars)")
	}
	if len(ch.Objectives) < 3 {
		t.Errorf("needs at least 3 Objectives")
	}
	if ch.Category != "" && len(ch.WhenToUse) < 3 {
		t.Errorf("needs at least 3 WhenToUse recognition cues")
	}
	if len(ch.Concepts) < 2 {
		t.Errorf("needs at least 2 Concepts, got %d", len(ch.Concepts))
	}
	for _, c := range ch.Concepts {
		if c.Description == "" || len(c.WhyItWorks) < 150 || len(c.Intuition) < 80 || len(c.CoreIdeas) < 3 || len(c.CommonMistakes) < 2 || len(c.Examples) < 1 {
			t.Errorf("concept %q needs Description, WhyItWorks (150+), Intuition (80+), 3+ CoreIdeas, 2+ CommonMistakes and an Example", c.Name)
		}
	}
	if len(ch.WorkedExamples) < 1 {
		t.Errorf("needs at least one WorkedExample")
	}
	for _, w := range ch.WorkedExamples {
		if w.Problem == "" || w.Code == "" || len(w.Steps) < 4 {
			t.Errorf("worked example %q needs Problem, Code and 4+ steps", w.Title)
		}
		for i, s := range w.Steps {
			if s.Subgoal == "" || len(s.Explanation) < 40 {
				t.Errorf("worked example %q step %d needs a Subgoal label and an Explanation", w.Title, i+1)
			}
		}
	}
	if len(ch.Completions) < 1 {
		t.Errorf("needs at least one Completion (faded example)")
	}
	for _, c := range ch.Completions {
		if n := strings.Count(c.Code, "____"); n == 0 || n != len(c.Blanks) {
			t.Errorf("completion %q has %d blanks in Code but %d answers", c.Title, n, len(c.Blanks))
		}
		if c.Prompt == "" || c.Explanation == "" {
			t.Errorf("completion %q needs Prompt and Explanation", c.Title)
		}
	}
	mc, recall := 0, 0
	for i, q := range ch.CheckQuestions {
		if q.Prompt == "" || len(q.Explanation) < 40 {
			t.Errorf("check question %d needs Prompt and an Explanation", i+1)
		}
		if len(q.Choices) > 0 {
			mc++
			if len(q.Choices) < 3 || q.Answer < 0 || q.Answer >= len(q.Choices) {
				t.Errorf("check question %d: needs 3+ choices and a valid Answer index", i+1)
			}
		} else {
			recall++
		}
	}
	if len(ch.CheckQuestions) < 5 || mc < 2 || recall < 1 {
		t.Errorf("needs 5+ CheckQuestions with 2+ multiple choice and 1+ free recall (has %d: %d mc, %d recall)", len(ch.CheckQuestions), mc, recall)
	}
	for _, sb := range ch.Animations {
		if !animations[sb.ID] {
			t.Errorf("storyboard %q has no animator registered in app.js", sb.ID)
		}
		if len(sb.MemoryTips) < 3 {
			t.Errorf("storyboard %q needs 3+ MemoryTips", sb.ID)
		}
	}
	// Exercises/Tutorials/Visualizers are legacy fields; new content uses
	// CheckQuestions, Completions and WorkedExamples instead.
	for _, f := range []string{"Exercises", "Tutorials", "Visualizers"} {
		if reflect.ValueOf(ch).FieldByName(f).Len() > 0 {
			t.Errorf("%s is legacy; use CheckQuestions, Completions and WorkedExamples", f)
		}
	}
}

var registryRe = regexp.MustCompile(`animationRegistry\['([^']+)'\]`)
var fullPageRe = regexp.MustCompile(`const fullPageTopics = \[([^\]]*)\]`)

func registeredAnimations(t *testing.T) map[string]bool {
	_, file, _, _ := runtime.Caller(0)
	src, err := os.ReadFile(filepath.Join(filepath.Dir(file), "..", "web", "static", "app.js"))
	if err != nil {
		t.Fatalf("reading app.js: %v", err)
	}
	out := map[string]bool{}
	for _, m := range registryRe.FindAllStringSubmatch(string(src), -1) {
		out[m[1]] = true
	}
	if m := fullPageRe.FindStringSubmatch(string(src)); m != nil {
		for _, id := range strings.Split(m[1], ",") {
			out[strings.Trim(strings.TrimSpace(id), `'"`)] = true
		}
	}
	return out
}
