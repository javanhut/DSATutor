# DSA Tutor

An interactive Data Structures and Algorithms tutor with step-by-step visualizations, code highlighting, and algorithm animations.

## Features

- **Algorithm Animations** - Step-by-step walkthroughs with synchronized code highlighting
- **Interactive Visualizations** - SVG-based visualizations for Big-O complexity curves
- **Variable Tracking** - Real-time display of algorithm state and variable values
- **Multiple Chapters** - Covers fundamental DSA topics from binary search to graph algorithms
- **Web Interface** - Modern, responsive UI with collapsible sidebar navigation

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

### Prerequisites

- Go 1.22 or later

### Running the Web Server

```bash
go run ./cmd/tutor -serve
```

Then open http://localhost:8080 in your browser.

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
├── cmd/tutor/           # Main application entry point
│   └── main.go
├── internal/
│   ├── chapter/         # Chapter definitions and templates
│   │   ├── chapter.go   # Chapter data structures
│   │   ├── registry.go  # Chapter registry
│   │   └── templates.go # Built-in chapter content
│   ├── storage/         # JSON import/export
│   ├── ui/              # Console renderer
│   └── web/             # Web server and static assets
│       ├── server.go
│       └── static/
│           ├── index.html
│           ├── app.js
│           └── styles.css
└── docs/                # Documentation
    └── visualizer-system.md
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
| `/static/*` | GET | Static assets |

## Development

### Building

```bash
go build -o dsatutor ./cmd/tutor
```

### Testing

```bash
go test ./...
```

## License

MIT
