# DSA Tutor

An interactive Data Structures and Algorithms course: 21 chapters that take a
beginner from Big O to dynamic programming, each leading into a ladder of
LeetCode-style practice problems that run and are checked in the browser.

## How it teaches

The course is built on well-replicated learning research; see
[docs/pedagogy.md](docs/pedagogy.md) for the evidence and the authoring rules.

- **Worked examples, then fading**: each chapter solves a problem step by step
  with subgoal labels, then gives a completion problem with the key lines
  blanked out, then the practice ladder.
- **Retrieval practice**: chapters end with questions answered from memory
  before the answer is shown.
- **Hints that fade**: every problem has a guiding question, the key insight,
  a subgoal plan, and a code skeleton, in that order. The solution stays behind
  a soft gate until you've made an attempt.
- **Explanations that say why**: each solution covers the brute force, the
  insight that improves it, a correctness argument, and common mistakes,
  followed by self-explanation questions.
- **Spaced review**: solved problems come back on a schedule set by how hard
  they felt.
- **Interleaved practice**: Mixed Practice draws from every topic you've
  started and asks you to name the pattern before coding.
- **Specific feedback**: every test case says what it checks; failures show
  expected vs actual, and the visualizer replays your run line by line.

## Features

- **Algorithm Animations**: step-by-step walkthroughs with synchronized code highlighting
- **Practice judge**: Python solutions run against visible and hidden tests,
  including linked-list, tree, graph and class-design problems
- **Execution visualization**: any run can be replayed line by line with variables and data structures
- **Sandbox**: free-form Python with the same visualizer

## Available Animators

| Algorithm | Description |
|-----------|-------------|
| Binary Search | Array visualization with low/mid/high pointers |
| Selection Sort | Bar chart with scanning, minimum tracking, and swaps |
| BFS | Graph traversal with queue visualization |
| DFS | Depth-first traversal with stack visualization |
| Call Stack | Recursion visualization with stack frames |
| Linked List | Node traversal and search operations |
| Two Pointers | Floyd's cycle detection with slow/fast pointers |
| Quicksort | Pivot selection and partitioning with bar chart |
| Hash Table | Bucket visualization with collision chaining |
| Dijkstra | Weighted graph with distance table and relaxation |
| Memoization | Fibonacci with memo table and call stack |
| Topological Sort | DAG ordering with in-degree tracking |
| Greedy | Activity selection with interval scheduling |
| String Match | Pattern matching with character comparison |
| Backtracking | N-Queens problem with conflict detection |
| Permutations | Generate permutations with backtracking |
| Iterative DP | Bottom-up Fibonacci with table filling |
| Rotated Array | Binary search on rotated sorted array |
| Merge Sort | Divide and conquer with merge visualization |
| Sliding Window | Fixed-size window maximum sum |
| Binary Tree | Inorder/preorder/postorder traversals |
| Heap | Min-heap insertion with bubble-up |
| Knapsack | 0/1 Knapsack DP with 2D table filling |

## Quick Start

Builds and tests go through ImLazy (`lazy.toml`); `imlazy list` shows the tasks.

```sh
imlazy setup    # installs Go and Python 3 with rvn if missing
imlazy serve    # http://localhost:8080
imlazy test     # unit tests, content checks, and every reference solution through the judge
```

### Command Line Options

```bash
# Start web server on custom port
go run ./cmd/tutor -serve -addr :3000

# Export chapters to JSON
go run ./cmd/tutor -export chapters.json

# Load chapters from JSON file
go run ./cmd/tutor -load chapters.json -serve

# Play a storyboard in the terminal
go run ./cmd/tutor -play 1
```

## Project Structure

```
DSATutor/
├── cmd/tutor/              # Main application entry point
├── internal/
│   ├── chapter/            # Curriculum: chapter.go (model), templates.go (order), chNN_*.go (one per chapter)
│   ├── practice/           # Problems (problems_*.go), categories, loader
│   ├── judge/              # Runs code against a problem's tests and decides pass/fail
│   ├── sandbox/            # Python execution: tracer.py, harness.py, security checks
│   ├── contentcheck/       # Tests that hold content to docs/pedagogy.md
│   ├── storage/            # JSON import/export
│   ├── ui/                 # Console renderer
│   └── web/                # HTTP server and static assets (app.js, learn.js, editor.js)
├── docs/                   # pedagogy.md, chapters.md, practice.md, sandbox.md, ...
├── scripts/system-deps.sh  # Installs Go and Python 3 with rvn
└── lazy.toml               # Build, serve and test tasks
```

## Web Interface

### Layout

The web interface features a three-panel workspace:

- **Code Panel** (left) - Algorithm code with line-by-line highlighting
- **Visualization Panel** (center) - Interactive algorithm visualization
- **Variables Panel** (right) - Real-time variable state and input data

### Controls

| Control | Action |
|---------|--------|
| Step | Advance one step |
| Back | Go back one step |
| Play/Pause | Auto-advance through steps |
| Reset | Return to initial state |
| Speed | Adjust animation speed |

### Sidebar

The sidebar displays available chapters and collapses when not in use. Hover to expand and select a chapter.

## Adding New Animations

1. Create an animator class extending `AlgorithmAnimator`:

```javascript
class MyAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.data = [...];
    this.setCode(`def my_algorithm(data):
    # code here`);
  }

  buildSteps() {
    this.steps = [];
    // Generate animation steps
  }

  render() {
    // Render visualization
  }

  getVariables() {
    return { /* current state */ };
  }

  getInputData() {
    return this.data;
  }
}
```

2. Register in `animationRegistry`:

```javascript
animationRegistry['my-algorithm'] = MyAnimator;
```

3. Add matching storyboard in `templates.go`:

```go
Animations: []Storyboard{
  {
    ID:    "my-algorithm",
    Title: "My Algorithm Walkthrough",
    // ...
  },
}
```

See [docs/visualizer-system.md](docs/visualizer-system.md) for detailed documentation.

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Web interface |
| `/api/chapters` | GET | List all chapters as JSON |
| `/api/practice/*` | GET/POST | Practice problems, run and submit; see [docs/practice.md](docs/practice.md) |
| `/static/*` | GET | Static assets |

## Development

```sh
imlazy build   # go build -o dsatutor ./cmd/tutor
imlazy test    # go vet + go test ./...
```

`imlazy test` includes two content gates:

- `internal/judge`: every problem's reference solution and alternatives pass
  every test through the same judge a learner's Submit uses.
- `internal/contentcheck`: chapters and problems have the structure
  [docs/pedagogy.md](docs/pedagogy.md) requires: fading hints, edge-case tests,
  correct walkthrough line ranges, valid chapter links, and so on.

## License

MIT
