# Chapter System

The DSA Tutor organizes learning content into chapters, each focusing on a specific algorithm or data structure concept. The chapter system is designed to help users understand not just HOW algorithms work, but WHY they work.

## Chapter Page Layout

Each chapter page follows a structured layout designed for optimal learning:

### 1. Overview (Top)
- Brief chapter summary
- Learning objectives displayed as cards
- Sets context for what will be covered

### 2. Understanding the Algorithm (Tutorial Section)
- Detailed concept explanations
- **Why It Works**: Deep explanation of the underlying principles
- **Mental Model**: Easy-to-understand analogies and intuitions
- **Key Points**: Core ideas to remember
- **Common Mistakes**: Pitfalls to avoid

### 3. Visualization
- Interactive animation controls (Play, Step, Reset)
- Main visualization canvas
- Code panel with line highlighting
- Variables panel showing current state

### 4. Steps to Remember
- Numbered memory tips for solving this type of problem
- Detailed storyboard walkthrough
- Step-by-step narration of the algorithm

### 5. Code Examples
- Complete, runnable code implementations
- Annotated with explanatory notes

## Chapter Data Model

Chapters are defined in `internal/chapter/templates.go` using these structures:

### Chapter
```go
type Chapter struct {
    Number      int           // Chapter number (1-12)
    Slug        string        // URL-friendly identifier
    Title       string        // Display title
    Summary     string        // Brief overview
    Objectives  []string      // Learning goals
    Concepts    []Concept     // Key concepts with explanations
    Animations  []Storyboard  // Step-by-step visualizations
    Tutorials   []Tutorial    // Guided lessons
    Exercises   []Exercise    // Practice problems
    Visualizers []Visualizer  // Interactive widgets
    Examples    []CodeExample // Code samples
}
```

### Concept
```go
type Concept struct {
    Name           string        // Concept name
    Description    string        // Brief description
    WhyItWorks     string        // Detailed explanation of WHY
    Intuition      string        // Mental model/analogy
    CoreIdeas      []string      // Key points
    CommonMistakes []string      // Pitfalls to avoid
    Examples       []CodeExample // Related code examples
}
```

### Storyboard
```go
type Storyboard struct {
    ID         string           // Identifier matching animation registry
    Title      string           // Display title
    Goal       string           // What the animation demonstrates
    MemoryTips []string         // Steps to remember for solving
    Steps      []StoryboardStep // Individual animation steps
}
```

## Adding a New Chapter

1. **Define chapter structure** in `internal/chapter/templates.go`:

```go
{
    Number:  13,
    Slug:    "my-algorithm",
    Title:   "My Algorithm",
    Summary: "Learn about my algorithm and when to use it.",
    Objectives: []string{
        "Understand the core principle",
        "Implement the algorithm",
        "Know when to apply it",
    },
    Concepts: []Concept{
        {
            Name:        "Core Concept",
            Description: "Brief description.",
            WhyItWorks:  "Detailed explanation of WHY this works...",
            Intuition:   "Think of it like...",
            CoreIdeas:   []string{"Point 1", "Point 2"},
            CommonMistakes: []string{
                "Mistake 1 - explanation",
                "Mistake 2 - explanation",
            },
            Examples: []CodeExample{
                {
                    ID:       "example_id",
                    Title:    "Example Title",
                    Language: "python",
                    Snippet:  `def my_function():\n    pass`,
                    Notes:    "Explanation of the code.",
                },
            },
        },
    },
    Animations: []Storyboard{
        {
            ID:    "my-animation",
            Title: "Algorithm Walkthrough",
            Goal:  "Show how the algorithm works step by step.",
            MemoryTips: []string{
                "Step 1: Initialize",
                "Step 2: Process",
                "Step 3: Return result",
            },
            Steps: []StoryboardStep{
                {
                    Cue:        "init",
                    Narration:  "Initialize the data structure.",
                    VisualHint: "Highlight initial state",
                    Duration:   Duration(2 * time.Second),
                    CodeRef:    "example_id: line 1",
                },
            },
        },
    },
}
```

2. **Create animation handler** if needed (see visualizer-system.md)

3. **Test the chapter** by running the application and navigating to the chapter

## Writing Effective WHY Explanations

Good "Why It Works" explanations should:

1. **Start with the core insight**: What property makes this algorithm possible?
2. **Connect to concrete outcomes**: Why does this property lead to the result we want?
3. **Quantify when possible**: How does this translate to time/space complexity?

### Example (Binary Search):
> "Binary search works because the data is sorted. When you check the middle element, you instantly know which half contains the target (if it exists). Since you eliminate half the remaining elements with each comparison, you can search through a billion items in just 30 steps instead of a billion steps."

## Writing Effective Intuitions

Good intuition/mental model explanations should:

1. **Use familiar scenarios**: Everyday analogies the reader already understands
2. **Highlight the key mapping**: What corresponds to what between the analogy and the algorithm?
3. **Keep it simple**: One clear analogy is better than three confusing ones

### Example (Hash Tables):
> "Think of a library that assigns books by the first letter of the author's last name. To find 'Smith', you don't search every shelf - you go directly to the 'S' section. The hash function is like that first-letter rule, but with more buckets and better distribution."

## Writing Memory Tips

Memory tips should be:

1. **Actionable**: Start with a verb (Check, Calculate, Compare)
2. **Ordered**: Follow the actual algorithm execution order
3. **Complete**: Cover all key steps, not just the tricky ones
4. **Memorable**: Use patterns like acronyms or visual cues when helpful

### Example (Binary Search):
1. Check if the array is SORTED
2. Set LOW to 0 and HIGH to length-1
3. Calculate MID as (low + high) / 2
4. Compare: if target equals mid value, you found it!
5. If target is GREATER, search RIGHT half (low = mid + 1)
6. If target is SMALLER, search LEFT half (high = mid - 1)
7. Repeat until low > high (not found) or element found

## JavaScript Rendering

The chapter content is rendered by these functions in `app.js`:

- `selectChapter(ch)` - Main function that loads all chapter content
- `renderTutorialContent(ch)` - Renders concepts with WHY explanations
- `renderMemoryTips(sb)` - Renders memory tips from storyboard
- `renderStoryboard(sb)` - Renders detailed storyboard steps
- `renderExamples(ch)` - Renders code examples

## CSS Styling

Key CSS classes for chapter content:

| Class | Description |
|-------|-------------|
| `.chapter-content` | Main content container |
| `.overview-card` | Chapter overview section |
| `.tutorial-card` | Algorithm tutorial section |
| `.viz-card` | Visualization section |
| `.memory-card` | Steps to remember section |
| `.examples-card` | Code examples section |
| `.concept-block` | Individual concept container |
| `.why-section` | Why It Works explanation |
| `.intuition-section` | Mental model explanation |
| `.core-ideas` | Key points list |
| `.mistakes-section` | Common mistakes list |
| `.memory-tip` | Individual memory tip |
