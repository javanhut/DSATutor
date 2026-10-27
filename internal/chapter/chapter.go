package chapter

import (
	"encoding/json"
	"time"
)

// Chapter is one lesson of the curriculum. Its sections follow the order in
// docs/pedagogy.md: motivate with a concrete problem, explain why the idea
// works, study a worked example, complete a faded one, recall it from memory,
// then practise the category's problem ladder.
type Chapter struct {
	Number        int    `json:"number"`
	Slug          string `json:"slug"`
	Title         string `json:"title"`
	Category      string `json:"category,omitempty"`      // practice category this chapter prepares for
	Prerequisites []int  `json:"prerequisites,omitempty"` // chapters to finish first
	Summary       string `json:"summary"`
	// Motivation is the concrete problem that opens the chapter, solved the
	// slow way first so the learner feels the need for the new idea.
	Motivation  string        `json:"motivation,omitempty"`
	Objectives  []string      `json:"objectives"`
	WhenToUse   []string      `json:"whenToUse,omitempty"` // cues in a problem statement that point to this pattern
	Concepts    []Concept     `json:"concepts"`
	Animations  []Storyboard  `json:"animations"`
	Tutorials   []Tutorial    `json:"tutorials"`
	Exercises   []Exercise    `json:"exercises"`
	Visualizers []Visualizer  `json:"visualizers"`
	Examples    []CodeExample `json:"examples"`

	WorkedExamples []WorkedExample     `json:"workedExamples,omitempty"`
	Completions    []CompletionProblem `json:"completions,omitempty"`
	CheckQuestions []CheckQuestion     `json:"checkQuestions,omitempty"`
}

// WorkedExample solves one problem step by step. Each step is labelled with
// the subgoal it achieves, so learners can reuse the plan on new problems.
type WorkedExample struct {
	Title   string       `json:"title"`
	Problem string       `json:"problem"`
	Steps   []WorkedStep `json:"steps"`
	Code    string       `json:"code"` // the finished solution
}

// WorkedStep is one subgoal of a worked example.
type WorkedStep struct {
	Subgoal     string `json:"subgoal"`         // short label, e.g. "Shrink the window"
	Explanation string `json:"explanation"`     // what to do and why
	Code        string `json:"code,omitempty"`  // the lines that do it
	State       string `json:"state,omitempty"` // a concrete trace, e.g. "left=1 right=3 seen={b,c}"
}

// CompletionProblem is a faded worked example: the solution with its key
// lines blanked out as ____ for the learner to fill in.
type CompletionProblem struct {
	Title       string   `json:"title"`
	Prompt      string   `json:"prompt"`
	Code        string   `json:"code"`        // contains one ____ per blank
	Blanks      []string `json:"blanks"`      // accepted answer for each blank, in order
	Explanation string   `json:"explanation"` // why those lines are the right ones
}

// CheckQuestion is a retrieval-practice question shown with its answer
// hidden. With Choices it is multiple choice and Answer indexes the correct
// choice; without Choices it is free recall and Explanation is the model
// answer the learner compares against.
type CheckQuestion struct {
	Prompt      string   `json:"prompt"`
	Choices     []string `json:"choices,omitempty"`
	Answer      int      `json:"answer"`
	Explanation string   `json:"explanation"`
}

// Concept is a key idea that you want the learner to remember.
type Concept struct {
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	WhyItWorks     string        `json:"whyItWorks"` // Detailed explanation of WHY this pattern works
	Intuition      string        `json:"intuition"`  // Easy-to-understand mental model
	CoreIdeas      []string      `json:"coreIdeas"`
	CommonMistakes []string      `json:"commonMistakes"` // What beginners often get wrong
	Examples       []CodeExample `json:"examples"`
}

// Storyboard is a high-level animation script that can be fed to a renderer.
type Storyboard struct {
	ID         string           `json:"id"`
	Title      string           `json:"title"`
	Goal       string           `json:"goal"`
	MemoryTips []string         `json:"memoryTips"` // Easy-to-remember steps for solving this type of problem
	Steps      []StoryboardStep `json:"steps"`
}

// StoryboardStep is a single beat in an animation or visualization.
type StoryboardStep struct {
	Cue        string   `json:"cue"`
	Narration  string   `json:"narration"`
	VisualHint string   `json:"visualHint"`
	Duration   Duration `json:"duration"`
	CodeRef    string   `json:"codeRef"` // optional: maps this beat to a code example/line
}

// Tutorial is a guided, semi-interactive lesson with prompts.
type Tutorial struct {
	ID      string         `json:"id"`
	Title   string         `json:"title"`
	Outcome string         `json:"outcome"`
	Steps   []TutorialStep `json:"steps"`
}

// TutorialStep outlines a single prompt in the guided flow.
type TutorialStep struct {
	Prompt     string `json:"prompt"`
	Guidance   string `json:"guidance"`
	CodeFocus  string `json:"codeFocus"`
	Checkpoint string `json:"checkpoint"`
}

// Exercise is for self-check questions; solution text is optional.
type Exercise struct {
	ID             string `json:"id"`
	Prompt         string `json:"prompt"`
	Difficulty     string `json:"difficulty"`
	Outcome        string `json:"outcome"`
	SolutionSketch string `json:"solutionSketch"`
}

// Visualizer describes an interactive widget you can attach to a chapter.
type Visualizer struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Goal         string   `json:"goal"`
	DataModel    string   `json:"dataModel"`
	Interactions []string `json:"interactions"`
	Hooks        []string `json:"hooks"` // e.g., "onStep", "onReset", "onSpeedChange"
}

// CodeExample holds sample code for live demos.
type CodeExample struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Language string `json:"language"`
	Snippet  string `json:"snippet"`
	Notes    string `json:"notes"`
}

// Duration wraps time.Duration to serialize cleanly as strings in JSON.
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var text string
	if err := json.Unmarshal(b, &text); err == nil {
		parsed, err := time.ParseDuration(text)
		if err != nil {
			return err
		}
		*d = Duration(parsed)
		return nil
	}

	// Fallback: accept milliseconds as a number.
	var ms float64
	if err := json.Unmarshal(b, &ms); err != nil {
		return err
	}
	*d = Duration(time.Duration(ms) * time.Millisecond)
	return nil
}
