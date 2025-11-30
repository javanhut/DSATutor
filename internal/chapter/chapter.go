package chapter

import (
	"encoding/json"
	"time"
)

// Chapter represents a reusable template for a tutorial-heavy, animated lesson.
// It focuses on structure over content so you can plug in visuals, code,
// and interactive prompts later.
type Chapter struct {
	Number      int           `json:"number"`
	Slug        string        `json:"slug"`
	Title       string        `json:"title"`
	Summary     string        `json:"summary"`
	Objectives  []string      `json:"objectives"`
	Concepts    []Concept     `json:"concepts"`
	Animations  []Storyboard  `json:"animations"`
	Tutorials   []Tutorial    `json:"tutorials"`
	Exercises   []Exercise    `json:"exercises"`
	Visualizers []Visualizer  `json:"visualizers"`
	Examples    []CodeExample `json:"examples"`
}

// Concept is a key idea that you want the learner to remember.
type Concept struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	WhyItWorks  string        `json:"whyItWorks"`  // Detailed explanation of WHY this pattern works
	Intuition   string        `json:"intuition"`   // Easy-to-understand mental model
	CoreIdeas   []string      `json:"coreIdeas"`
	CommonMistakes []string   `json:"commonMistakes"` // What beginners often get wrong
	Examples    []CodeExample `json:"examples"`
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
