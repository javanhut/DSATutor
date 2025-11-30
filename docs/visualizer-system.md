# Visualizer and Animation System

The DSA Tutor provides two visualization systems:

1. **Visualizers** - Static/interactive visualizations tied to visualizer metadata (e.g., Big-O curves)
2. **Algorithm Animators** - Step-by-step algorithm animations with code highlighting tied to storyboards

## Architecture Overview

Both systems use a registry pattern to map IDs (defined in the backend) to JavaScript handler classes.

### Full-Page Topics

Some topics render as completely standalone pages that replace the standard chapter content. When a full-page topic is selected from the topic dropdown, the standard chapter sections (overview, tutorial, examples, etc.) are hidden and the topic renders its own complete educational experience.

**Full-page topics are defined in:**
```javascript
const fullPageTopics = ['runtime-shapes'];
```

**Key functions:**
- `isFullPageTopic(storyboardId)` - Check if a topic is full-page
- `showFullPageTopic(sb)` - Mount full-page content into `#topic-full-page`
- `hideFullPageTopic()` - Unmount and restore standard content
- `handleTopicChange(sb)` - Main handler that routes between full-page and standard

**HTML Structure:**
```html
<div id="topic-full-page" class="topic-full-page hidden">
  <!-- Full-page topic content -->
</div>
<div id="standard-chapter-content">
  <!-- Standard chapter sections -->
</div>
```

### Key Components

1. **Visualizer Registry** (`visualizerRegistry` in `app.js`)
   - Maps visualizer IDs to handler classes
   - Located at the top of `app.js`

2. **Handler Classes** (e.g., `RuntimeShapesVisualizer`)
   - Implement the visualization rendering and interaction logic
   - Must implement the standard interface methods

3. **Hook System**
   - Hooks defined in backend metadata (`onStep`, `onReset`, `onCurveSelect`)
   - Dispatched to handler methods when storyboard events occur

## Handler Interface

All visualizer handlers must implement these methods:

```javascript
class MyVisualizer {
  constructor(container, config) {
    // container: DOM element to render into
    // config: visualizer metadata from backend
  }

  mount() {
    // Called when visualizer is attached
    // Build DOM elements and render initial state
  }

  unmount() {
    // Called when visualizer is detached
    // Clean up animations, event listeners, etc.
  }

  // Hook handlers (optional, based on metadata)
  onStep(payload) {
    // payload: { step: number, cue: string }
  }

  onReset(payload) {
    // payload: {}
  }

  onCurveSelect(payload) {
    // payload: { curve: string, visible: boolean }
  }
}
```

## Existing Visualizers

### RuntimeShapesVisualizer (`timeline-big-o`)

A comprehensive educational visualizer for understanding asymptotic complexity (Big O notation). When accessed via the "Runtime Shapes" topic in the Introduction to Algorithms chapter, it provides a complete learning experience about algorithm complexity.

**Features:**
- SVG rendering with logarithmic Y-axis scaling
- Toggle buttons to show/hide individual curves
- Interactive N-value slider for scrubbing
- Animated transitions during storyboard playback
- Comprehensive asymptotic analysis explanation
- Input scaling comparison table showing concrete operation counts
- Time vs Space complexity reference table
- Real-world analogies for each complexity class
- Code pattern examples for each complexity

**Educational Sections:**

1. **Asymptotic Analysis Intro**
   - What is Big O: Description of growth rate notation
   - Why Ignore Constants: Explanation of why constants don't matter at scale
   - Time vs Space Complexity: Covers the difference and tradeoffs

2. **Interactive Complexity Graph**
   - Curves for O(1), O(log n), O(n), O(n log n), O(n^2)
   - Operations table showing concrete numbers at current N
   - Time estimates (if 1ms per operation)
   - Click curves for detailed explanations

3. **Input Scaling Table**
   - Shows operation counts at n=10, 100, 1000, 10000, 1000000
   - Color-coded cells: green (fast), orange (medium), red (slow)
   - Demonstrates how algorithms diverge at scale

4. **Time vs Space Reference**
   - Each complexity with its name, code pattern, space example
   - Real-world analogies for intuitive understanding
   - Tip cards for time-space tradeoff and practical considerations

**Curves displayed:**
- O(1) - Constant (green) - Array lookup, hash access
- O(log n) - Logarithmic (cyan) - Binary search
- O(n) - Linear (orange) - Array traversal
- O(n log n) - Linearithmic (purple) - Merge sort
- O(n^2) - Quadratic (red) - Nested loops

**Configuration (`curveConfig`):**
Each curve includes:
- `color`: Display color
- `label`: Short label
- `name`: Full name (e.g., "Constant")
- `examples`: Array of algorithm examples
- `why`: Explanation of why this complexity occurs
- `spaceExample`: Typical space usage at this complexity
- `codePattern`: Representative code pattern
- `realWorld`: Real-world analogy

**Hooks:**
- `onStep`: Animates to target N value based on step index
- `onReset`: Resets visualization to n=1
- `onCurveSelect`: Toggles curve visibility

**CSS Classes:**
| Class | Description |
|-------|-------------|
| `.runtime-shapes-wrapper` | Main container for all sections |
| `.asymptotic-intro` | Intro section with gradient background |
| `.asymptotic-sections` | Grid of explanation cards |
| `.input-scaling-section` | Scaling comparison table section |
| `.scaling-table` | Table showing ops at different N |
| `.time-fast` / `.time-medium` / `.time-slow` | Color-coded cells |
| `.time-space-section` | Reference table section |
| `.time-space-table` | Complexity reference table |
| `.tip-card` | Educational tip cards |

## Adding a New Visualizer

1. **Define metadata in backend** (`internal/chapter/templates.go`):
```go
Visualizers: []Visualizer{
  {
    ID:           "my-visualizer-id",
    Title:        "My Visualizer",
    Goal:         "Description of what it visualizes",
    DataModel:    "Description of data structure",
    Interactions: []string{"interaction-1", "interaction-2"},
    Hooks:        []string{"onStep", "onReset"},
  },
}
```

2. **Create handler class** in `app.js`:
```javascript
class MyVisualizer {
  constructor(container, config) {
    this.container = container;
    this.config = config;
  }

  mount() {
    // Render your visualization
  }

  unmount() {
    // Cleanup
  }

  onStep(payload) {
    // Handle step changes
  }

  onReset() {
    // Reset state
  }
}
```

3. **Register in visualizerRegistry**:
```javascript
const visualizerRegistry = {
  'timeline-big-o': RuntimeShapesVisualizer,
  'my-visualizer-id': MyVisualizer,  // Add your visualizer
};
```

4. **Add CSS styles** in `styles.css` as needed.

## CSS Classes

| Class | Description |
|-------|-------------|
| `.viz-canvas` | SVG container element |
| `.viz-controls` | Controls container (toggles + slider) |
| `.curve-toggles` | Container for toggle buttons |
| `.curve-toggle` | Individual toggle button |
| `.curve-toggle.active` | Active toggle state |
| `.toggle-dot` | Colored dot in toggle button |
| `.slider-wrap` | Slider container |
| `.n-slider` | Range input for N value |
| `.n-label` | Label showing current N value |

## Hook Event Flow

```
Storyboard Play
    |
    v
playStoryboard()
    |
    v
setTimeout for each step
    |
    v
emitVisualizerHook("onStep", { step, cue })
    |
    v
Check if handler exists and has method
    |
    v
handler.onStep(payload)
    |
    v
Visualization animates
```

## Storyboard Step Mapping

The `stepToNMap` object defines which N value to animate to for each storyboard step:

```javascript
const stepToNMap = { 0: 10, 1: 25, 2: 50, 3: 75, 4: 100 };
```

Customize this mapping based on your storyboard content.

---

# Algorithm Animation System

The animation system provides detailed step-by-step algorithm walkthroughs with synchronized code highlighting.

## Animation Registry

```javascript
const animationRegistry = {
  'binary-search': BinarySearchAnimator,
  'selection-sort': SelectionSortAnimator,
  'bfs': BFSAnimator,
  'dfs': DFSAnimator,
  'call-stack': CallStackAnimator,
  'linked-list': LinkedListAnimator,
  'two-pointers': TwoPointersAnimator,
  'quicksort': QuicksortAnimator,
  'hash-table': HashTableAnimator,
  'dijkstra': DijkstraAnimator,
  'memoization': MemoizationAnimator,
  'topsort': TopSortAnimator,
  'greedy': GreedyAnimator,
  'string-match': StringMatchAnimator,
  'backtracking': BacktrackingAnimator,
  'permutations': PermutationsAnimator,
  'iterative-dp': IterativeDPAnimator,
  'rotated-array': RotatedArrayAnimator,
  'merge-sort': MergeSortAnimator,
  'sliding-window': SlidingWindowAnimator,
  'binary-tree': BinaryTreeAnimator,
  'heap': HeapAnimator,
  'knapsack': KnapsackAnimator,
};
```

## AlgorithmAnimator Base Class

All algorithm animators extend this base class:

```javascript
class AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) { }

  setCode(code)           // Set the code to display
  highlightLine(lineNum)  // Highlight a single line
  highlightLines(nums)    // Highlight multiple lines
  buildSteps()            // Generate animation steps (override)
  render()                // Render current state (override)
  goToStep(idx)           // Jump to specific step
  stepForward()           // Advance one step
  stepBack()              // Go back one step
  play()                  // Start auto-play
  pause()                 // Pause auto-play
  reset()                 // Reset to initial state
  mount()                 // Initialize and render
  unmount()               // Cleanup
  getVariables()          // Return current variable state (override)
  getInputData()          // Return input data for display (override)
  updateVariables(vars)   // Update the variables panel
  updateInputDisplay(data) // Update the input data panel
}
```

## Step Structure

Each animation step contains:

```javascript
{
  lineNum: 5,                    // Single line to highlight
  lineNums: [6, 7],              // Multiple lines to highlight
  state: "Description text",     // State description shown below viz
  apply: () => { /* mutate */ }, // Function to apply state changes
  // ... custom properties for rendering
}
```

## Available Animators

### BinarySearchAnimator (`binary-search`)
- Shows array with low/mid/high pointers
- Highlights discarded portions
- Demonstrates halving search space

### SelectionSortAnimator (`selection-sort`)
- Bar chart visualization with heights
- Shows scanning pointer, minimum tracking
- Animates swaps between elements

### BFSAnimator (`bfs`)
- Graph visualization with nodes and edges
- Queue display showing frontier
- Visit order tracking

### DFSAnimator (`dfs`)
- Graph visualization with stack-based traversal
- Stack display (LIFO) showing current path
- Highlights visited vs in-stack nodes

### CallStackAnimator (`call-stack`)
- Stack frame visualization
- Shows push/pop phases
- Demonstrates recursion unwinding

### LinkedListAnimator (`linked-list`)
- Node chain visualization
- Current pointer tracking
- Search operation with comparison highlighting

### TwoPointersAnimator (`two-pointers`)
- Floyd's cycle detection (tortoise and hare)
- Slow pointer (1 step) vs fast pointer (2 steps)
- Cycle detection visualization

### QuicksortAnimator (`quicksort`)
- Bar chart visualization
- Pivot selection and partitioning
- Recursive subdivision tracking

### HashTableAnimator (`hash-table`)
- Bucket array visualization
- Hash calculation display
- Collision handling with chaining

### DijkstraAnimator (`dijkstra`)
- Weighted graph with edge labels
- Distance table showing shortest paths
- Edge relaxation animation

### MemoizationAnimator (`memoization`)
- Fibonacci with memoization
- Memo table showing cached results
- Call stack visualization
- Memo hits highlighted

### TopSortAnimator (`topsort`)
- DAG visualization with in-degrees
- Queue of zero in-degree nodes
- Result order building

### GreedyAnimator (`greedy`)
- Activity selection / interval scheduling
- Timeline visualization with activities
- Current end marker tracking
- Compatible vs overlapping activity detection
- Selected activities summary

### StringMatchAnimator (`string-match`)
- Text and pattern character display
- Character-by-character comparison
- Match/mismatch highlighting
- Pattern sliding visualization

### BacktrackingAnimator (`backtracking`)
- N-Queens chessboard visualization
- Safe position checking
- Queen placement and removal
- Conflict detection highlighting

### PermutationsAnimator (`permutations`)
- Path and remaining elements display
- Element selection visualization
- Backtracking animation
- Generated permutations list

### IterativeDPAnimator (`iterative-dp`)
- Bottom-up DP table filling
- Base case initialization
- Dependency highlighting
- Formula visualization

### RotatedArrayAnimator (`rotated-array`)
- Rotated sorted array display
- Pivot point identification
- Sorted half detection
- Binary search pointer tracking

### MergeSortAnimator (`merge-sort`)
- Divide phase tree visualization
- Array splitting animation
- Merge phase comparison
- Sorted subarrays combining

### SlidingWindowAnimator (`sliding-window`)
- Array with window boundaries
- Element add/remove animation
- Window sum tracking
- Maximum sum display

### BinaryTreeAnimator (`binary-tree`)
- SVG tree structure
- Node traversal animation
- Inorder result tracking
- Current node highlighting

### HeapAnimator (`heap`)
- Tree and array representation
- Bubble-up animation
- Parent-child comparisons
- Min-heap property visualization

### KnapsackAnimator (`knapsack`)
- 0/1 Knapsack dynamic programming
- 2D DP table visualization
- Item display with weight and value
- Cell highlighting for current, lookup, and result states
- Include vs exclude decision display
- Final result with selected items traceback

## Adding a New Algorithm Animation

1. Create a class extending `AlgorithmAnimator`:

```javascript
class MyAlgorithmAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    // Initialize state
    this.data = [...];

    // Set the code to display
    this.setCode(`def my_algorithm(data):
    # algorithm code here
    pass`);
  }

  buildSteps() {
    this.steps = [];

    // Generate steps by simulating the algorithm
    this.steps.push({
      lineNum: 1,
      state: "Starting algorithm",
      myData: [...this.data],
      apply: () => { /* update this.data */ }
    });

    // ... more steps
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const data = step.myData || this.data;

    // Render visualization to this.canvas
    this.canvas.innerHTML = `<div>...</div>`;
  }
}
```

2. Register in `animationRegistry`:

```javascript
animationRegistry['my-algorithm'] = MyAlgorithmAnimator;
```

3. Use matching storyboard ID in `templates.go`:

```go
Animations: []Storyboard{
  {
    ID:    "my-algorithm",
    Title: "My Algorithm Walkthrough",
    // ...
  },
}
```

## UI Controls

The animation workspace provides:
- **Step Forward/Back** - Manual stepping
- **Play/Pause** - Auto-advance through steps
- **Reset** - Return to initial state
- **Speed slider** - Adjust auto-play speed
- **Step counter** - Shows current position

## CSS Classes Reference

| Class | Description |
|-------|-------------|
| `.animation-workspace` | Main container |
| `.animation-controls` | Control buttons bar |
| `.animation-split` | Side-by-side layout |
| `.code-panel` | Code display area |
| `.code-line` | Individual code line |
| `.code-line.active` | Highlighted line |
| `.viz-panel` | Visualization area |
| `.state-display` | Current state text |
| `.array-viz` | Array visualization |
| `.array-cell` | Array element |
| `.matrix-viz` | Matrix (2D array) visualization |
| `.matrix-table` | Matrix table element |
| `.matrix-cell` | Matrix cell element |
| `.graph-viz` | Graph SVG |
| `.graph-node` | Graph node circle |
| `.stack-viz` | Stack visualization |
| `.stack-frame` | Stack frame element |

---

# Page Layout

The DSA Tutor uses a full-width layout with a collapsible sidebar for navigation.

## Layout Structure

```
+------------------+----------------------------------------+
|                  |                                        |
|     Sidebar      |            Main Content                |
|   (Chapters)     |                                        |
|                  |  +----------+----------+----------+    |
|                  |  |   Code   |   Viz    |   Vars   |    |
|                  |  |  Panel   |  Panel   |  Panel   |    |
|                  |  +----------+----------+----------+    |
|                  |                                        |
+------------------+----------------------------------------+
```

## Sidebar Behavior

- **Collapsed state**: 60px width, shows chapter numbers only
- **Expanded state**: 260px width, shows full chapter info
- **Hover-triggered**: Expands smoothly on mouse hover
- **Transition**: 200ms ease animation

## Three-Panel Workspace

The animation workspace displays three panels:

1. **Code Panel** (left): Shows the algorithm code with line highlighting
2. **Visualization Panel** (center): Renders the algorithm visualization
3. **Variables Panel** (right): Displays current variable state and input data

## Variable Tracking

Each animator implements `getVariables()` and `getInputData()` methods:

```javascript
class MyAnimator extends AlgorithmAnimator {
  getVariables() {
    return {
      'variable_name': value,
      'another_var': another_value,
    };
  }

  getInputData() {
    return this.inputArray; // or object with input info
  }
}
```

The variables panel automatically updates on each step to show:
- Current variable values
- Input data being processed
- Value changes as the algorithm progresses

## CSS Variables

Key CSS variables for theming:

| Variable | Default | Description |
|----------|---------|-------------|
| `--sidebar-width` | 60px | Collapsed sidebar width |
| `--sidebar-expanded` | 260px | Expanded sidebar width |
| `--bg` | #0a0a0f | Background color |
| `--panel` | #12121a | Panel background |
| `--accent` | #38bdf8 | Primary accent color |
| `--accent-2` | #22c55e | Secondary accent (success) |
| `--warning` | #f59e0b | Warning/number color |

## Responsive Breakpoints

- **1200px**: Variables panel hidden, 2-column layout
- **900px**: Single column layout, code panel height reduced
