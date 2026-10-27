# Chapter System

Chapters are the lessons of the curriculum, numbered in learning order. Each
chapter teaches one pattern and, when it has a `Category`, leads straight into
that category's practice ladder. What a chapter must contain, and the research
behind it, is in [pedagogy.md](pedagogy.md). This page covers the mechanics.

## Where chapters live

- `internal/chapter/chapter.go`: the data model
- `internal/chapter/templates.go`: `DefaultChapters()`, the curriculum order
- `internal/chapter/chNN_<slug>.go`: one file per chapter
- `internal/contentcheck/content_test.go`: the checks every chapter must pass

## Chapter page layout

The page reads top to bottom as one lesson:

1. **Overview**: summary, the motivating problem ("why you need this"),
   objectives, and prerequisite chapters.
2. **Understanding the Algorithm**: each concept's *why it works*, mental
   model, key points and common mistakes.
3. **When to Use It**: cues in a problem statement that point to the pattern.
4. **Worked Example**: revealed one subgoal-labelled step at a time, with a
   prompt to predict each step first.
5. **Visualization** and **Steps to Remember**: the animated storyboard and
   its memory tips (hidden when a chapter has no storyboard).
6. **Your Turn: Complete the Code**: the faded example, checked in the browser.
7. **Check Your Understanding**: retrieval questions, answered before the
   answer is shown. Answering all of them marks the chapter done (✓ in the
   sidebar).
8. **Practice Ladder**: the category's problems in order, with progress.
9. **Code Examples**: every concept's code in one place.

## Data model

```go
type Chapter struct {
    Number, Slug, Title string
    Category       string              // practice category it teaches
    Prerequisites  []int               // earlier chapters
    Summary        string
    Motivation     string              // concrete problem solved the slow way
    Objectives     []string
    WhenToUse      []string            // recognition cues
    Concepts       []Concept           // WhyItWorks, Intuition, CoreIdeas, CommonMistakes, Examples
    Animations     []Storyboard        // ID must have an animator in app.js
    WorkedExamples []WorkedExample     // steps: Subgoal, Explanation, Code, State
    Completions    []CompletionProblem // Code with ____ blanks + Blanks answers
    CheckQuestions []CheckQuestion     // multiple choice (Choices + Answer) or free recall
}
```

`Exercises`, `Tutorials` and `Visualizers` are older fields that the page no
longer shows.

## Adding a chapter

1. Create `internal/chapter/chNN_<slug>.go` with a `chapterNN()` function, and
   add it to `DefaultChapters()` in order.
2. If it teaches a practice category, set `Category` and make that category's
   `ChapterRefs[0]` (in `internal/practice/category.go`) point at it.
3. Storyboard IDs must match an animator registered in
   `internal/web/static/app.js` (`animationRegistry['id'] = ...`); see
   [visualizer-system.md](visualizer-system.md) to add one.
4. Run `imlazy test`. The content checks report anything missing.
