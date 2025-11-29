# Practice Mode

Practice mode provides a LeetCode-style problem-solving experience with integrated visualizations, progressive hints, and progress tracking. This feature is designed to complement the chapter-based learning with hands-on coding challenges.

## Features

- **Blind 75 Problems**: Core problems embedded directly in Go code
- **Progressive Hints**: Approach, algorithm, and code hints revealed incrementally
- **Test Case Validation**: Run code against individual test cases or submit for full validation
- **Chapter Links**: Each problem links to related learning chapters
- **Execution Visualization**: Step-by-step visualization of code execution
- **Progress Tracking**: LocalStorage-based progress with export/import capability
- **Extensible**: Add custom problems via JSON files

## Architecture

### Backend Components

```
internal/practice/
  problem.go           - Core data structures (Problem, TestCase, Hint, Solution)
  embedded_problems.go - Core Blind 75 problems defined in pure Go
  submission.go        - Submission handling (RunRequest, SubmitRequest, Results)
  category.go          - Problem categories (Blind75Categories)
  loader.go            - Problem loading (embedded Go + optional JSON files)
  embed.go             - Embedded filesystem for custom problems
  problems/            - Directory for custom JSON problems (optional)
```

### Problem Loading Strategy

1. **Embedded Problems** (primary): Core Blind 75 problems are defined in `embedded_problems.go` as pure Go structs. These are always available and don't require any external files.

2. **JSON Problems** (optional): Additional problems can be added via JSON files in `internal/practice/problems/`. These are loaded on startup and merged with the embedded problems.

### Frontend Components

The practice UI is integrated into the main `app.js` with:

- `practiceState` - Global state for practice mode
- Problem list view with category grouping
- Problem workspace with 3-column layout
- Code editor with run/submit functionality
- Test case panel with pass/fail indicators
- Hints panel with progressive reveal
- Solution panel with walkthrough

### API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/practice/problems` | GET | List all problems (returns `{problems: [], total, stats}`) |
| `/api/practice/problems/{id}` | GET | Get problem details |
| `/api/practice/categories` | GET | List all categories |
| `/api/practice/run` | POST | Run code against single test case |
| `/api/practice/submit` | POST | Run code against all test cases |
| `/api/practice/hint/{id}/{index}` | GET | Get specific hint |
| `/api/practice/solution/{id}` | GET | Get full solution |
| `/api/practice/progress/export` | GET | Export progress JSON |
| `/api/practice/progress/import` | POST | Import progress JSON |

## Embedded Problems

The core Blind 75 problems are defined directly in Go code in `internal/practice/embedded_problems.go`. This provides:

- **No external dependencies**: Problems are compiled into the binary
- **Type safety**: Full Go compiler checks on problem definitions
- **Easy maintenance**: Edit problems directly in Go code
- **Fast loading**: No file I/O or JSON parsing at startup

Example embedded problem:

```go
{
    ID:         "two-sum",
    Number:     1,
    Title:      "Two Sum",
    Difficulty: "Easy",
    Category:   "arrays-hashing",
    Tags:       []string{"Array", "Hash Table"},
    RelatedChapters: []int{2, 3},
    Description: `Given an array of integers nums and an integer target...`,
    Constraints: []string{
        "2 <= nums.length <= 10^4",
        "-10^9 <= nums[i] <= 10^9",
    },
    Examples: []Example{
        {Input: "nums = [2,7,11,15], target = 9", Output: "[0,1]"},
    },
    TestCases: []TestCase{
        {Input: map[string]interface{}{"nums": []int{2, 7, 11, 15}, "target": 9}, Expected: []int{0, 1}},
    },
    StarterCode: "def twoSum(nums, target):\n    pass",
    Hints: []Hint{
        {Level: 1, Type: "approach", Content: "Use a hash table..."},
    },
    Solution: Solution{
        Code: `def twoSum(nums, target):
    seen = {}
    for i, num in enumerate(nums):
        if target - num in seen:
            return [seen[target - num], i]
        seen[num] = i`,
    },
}
```

## Adding Custom Problems

To add problems beyond the embedded Blind 75, create JSON files in `internal/practice/problems/`:

### JSON Problem Format

```json
{
  "id": "custom-problem",
  "number": 100,
  "title": "Custom Problem",
  "difficulty": "Medium",
  "category": "arrays-hashing",
  "tags": ["Array", "Custom"],
  "relatedChapters": [2],
  "description": "Problem description...",
  "constraints": ["constraint 1"],
  "examples": [
    {
      "input": "nums = [1,2,3]",
      "output": "6",
      "explanation": "Sum of all elements"
    }
  ],
  "testCases": [
    {
      "input": {"nums": [1, 2, 3]},
      "expected": 6,
      "hidden": false
    }
  ],
  "timeComplexity": "O(n)",
  "spaceComplexity": "O(1)",
  "starterCode": "def solve(nums):\n    pass",
  "hints": [
    {"type": "approach", "content": "Hint 1"},
    {"type": "algorithm", "content": "Hint 2"}
  ],
  "solution": {
    "approach": "Simple sum",
    "code": "def solve(nums):\n    return sum(nums)",
    "walkthrough": []
  }
}
```

### Adding a Custom Problem

1. Create a JSON file in `internal/practice/problems/`:
   ```
   internal/practice/problems/custom-problem.json
   ```

2. Rebuild the application to embed the new file:
   ```bash
   go build ./cmd/tutor
   ```

3. The problem will be loaded on startup and merged with embedded problems.

## Problem Categories

The Blind 75 problems are organized into these categories (defined in `category.go`):

| ID | Name | Description |
|----|------|-------------|
| arrays-hashing | Arrays & Hashing | Array manipulation and hash table problems |
| two-pointers | Two Pointers | Problems using the two-pointer technique |
| sliding-window | Sliding Window | Sliding window pattern problems |
| stack | Stack | Stack-based problems |
| binary-search | Binary Search | Binary search and variants |
| linked-list | Linked List | Linked list manipulation problems |
| trees | Trees | Binary tree and BST problems |
| tries | Tries | Trie (prefix tree) problems |
| heap-priority-queue | Heap / Priority Queue | Heap and priority queue problems |
| backtracking | Backtracking | Backtracking and recursive enumeration |
| graphs | Graphs | Graph traversal and algorithms |
| 1d-dp | 1-D Dynamic Programming | One-dimensional DP problems |
| 2d-dp | 2-D Dynamic Programming | Two-dimensional DP problems |
| greedy | Greedy | Greedy algorithm problems |
| intervals | Intervals | Interval scheduling and merging |
| math-geometry | Math & Geometry | Mathematical and geometric problems |
| bit-manipulation | Bit Manipulation | Bitwise operation problems |

## Progress Tracking

Progress is stored in browser LocalStorage under the key `dsatutor_practice_progress`:

```json
{
  "two-sum": {
    "attempts": 3,
    "solved": true,
    "hintsUsed": 1,
    "lastAttempt": "2024-01-15T10:30:00.000Z"
  }
}
```

### Export/Import

Users can export their progress to a JSON file and import it on another device:

- **Export**: Click "Export Progress" button to download JSON file
- **Import**: Click "Import Progress" to upload a previously exported file

## Code Execution

Practice mode uses the existing sandbox system for code execution:

1. User code is wrapped with test case inputs
2. Code runs in Python sandbox with tracing enabled
3. Output is compared against expected result
4. Execution steps are captured for visualization

### Test Input Formats

The loader supports two input formats:

1. **Object format** (preferred):
   ```json
   {"input": {"nums": [1, 2, 3], "target": 5}}
   ```

2. **String format** (legacy):
   ```json
   {"input": "[1, 2, 3], 5"}
   ```

## UI Components

### Problem List View

- Category sections with progress counters
- Problem cards showing difficulty and solved status
- Filter controls for difficulty, category, status, and search

### Problem Workspace

Three-column layout:

1. **Left Panel**: Problem description, examples, constraints
2. **Center Panel**: Code editor with run/submit buttons
3. **Right Panel**: Tabs for test cases, output, hints, visualization, solution

### Visualization

Code execution visualization reuses the existing sandbox visualization system:
- Step-by-step execution trace
- Variable state at each step
- Data structure visualizations (arrays, linked lists, trees)

## Usage

1. Click "Practice" in the sidebar to open practice mode
2. Select a problem from the list
3. Read the description and examples
4. Write your solution in the editor
5. Click "Run" to test against the current test case
6. Click "Submit" to validate against all test cases
7. Use hints if stuck (revealed progressively)
8. View solution after solving or giving up

## Configuration

Practice mode is enabled by default when problems are loaded. The server logs the number of loaded problems at startup:

```
Loaded 20 embedded problems
Total problems loaded: 20
```

If no problems are found, the server continues without practice support, and the Practice button is hidden.
