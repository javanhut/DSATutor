package ui

import (
	"fmt"
	"io"
	"os"
	"time"

	"dsatutor/internal/chapter"
)

// StoryboardRenderer describes how to play a storyboard; swap this for
// real animation engines later.
type StoryboardRenderer interface {
	Render(sb chapter.Storyboard) error
}

// ConsoleRenderer prints storyboard cues and durations to the terminal.
// It is meant as a placeholder until a graphical renderer is wired up.
type ConsoleRenderer struct {
	Out io.Writer
}

func (r ConsoleRenderer) Render(sb chapter.Storyboard) error {
	out := r.Out
	if out == nil {
		out = os.Stdout
	}
	fmt.Fprintf(out, "Storyboard: %s (%s)\n", sb.Title, sb.Goal)
	for idx, step := range sb.Steps {
		fmt.Fprintf(out, " %02d) cue=%s | %s | hint=%s | duration=%s\n",
			idx+1, step.Cue, step.Narration, step.VisualHint, time.Duration(step.Duration))
	}
	return nil
}

// VisualizerMount lets you plug visualizers into a container (UI, web canvas, etc).
type VisualizerMount interface {
	Mount(v chapter.Visualizer) error
}

// ConsoleVisualizerMount is a no-op placeholder that just logs mount requests.
type ConsoleVisualizerMount struct {
	Out io.Writer
}

func (v ConsoleVisualizerMount) Mount(vis chapter.Visualizer) error {
	out := v.Out
	if out == nil {
		out = os.Stdout
	}
	fmt.Fprintf(out, "Mount visualizer: %s | goal: %s | hooks: %v\n", vis.Title, vis.Goal, vis.Hooks)
	return nil
}
