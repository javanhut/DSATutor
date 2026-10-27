# Practice Mode

Practice mode is where learners solve LeetCode-style problems in Python, or
in Oxigen (see [oxigen.md](oxigen.md)), after each chapter. Problems are grouped into categories, and each category is
a ladder ordered easy → hard. Every Blind 75 problem is included, plus 64
NeetCode 150 problems, for 139 in total. How problems are written, and why, is
in [pedagogy.md](pedagogy.md).

## What a learner gets

- **Ladders**: each category lists its problems in `Order`, and each step
  adds one idea to the last. A link at the top of each category leads to the
  chapter that teaches it.
- **Fading hints**: a guiding question, then the key insight, then a subgoal
  plan, then a code skeleton with blanks.
- **Run** a visible test case with a line-by-line visualization, or **Submit**
  against every case, hidden ones included. Every case says what it checks.
- **Solutions**:
  - The solution unlocks after an attempt or all the hints; it can be
    overridden.
  - It covers the brute force, the key insight, why it works, the code with a
    subgoal-labelled walkthrough, common mistakes, alternatives with runnable
    code, and self-check questions.
- **After solving**: self-explanation questions and a difficulty rating that
  schedules the problem for spaced review.
- **Review**: a solved problem comes back after 3 days, and each interval
  then grows about 2.5×. A Hard rating grows it 1.2× and an Easy rating 3.5×.
  Viewing the solution before solving brings a problem back the next day.
- **Mixed Practice**: picks problems from every category the learner has
  started and hides the category. The learner names the pattern before
  coding.

Progress lives in the browser's localStorage, under
`dsatutor_practice_progress` and `dsatutor_chapter_progress`, and can be
exported and imported as JSON.

## Categories

| # | Category | Problems | Taught in chapter |
|---|---|---|---|
| 1 | Arrays & Hashing | 9 | 2 |
| 2 | Two Pointers | 5 | 3 |
| 3 | Sliding Window | 6 | 4 |
| 4 | Stack | 7 | 5 |
| 5 | Binary Search | 7 | 6 |
| 6 | Linked List | 9 | 7 |
| 7 | Trees | 11 | 9 |
| 8 | Tries | 3 | 10 |
| 9 | Heap / Priority Queue | 6 | 11 |
| 10 | Backtracking | 9 | 12 |
| 11 | Graphs | 12 | 13 |
| 12 | Advanced Graphs | 6 | 14 |
| 13 | 1-D Dynamic Programming | 11 | 15 |
| 14 | 2-D Dynamic Programming | 11 | 16 |
| 15 | Greedy | 8 | 17 |
| 16 | Intervals | 5 | 18 |
| 17 | Math & Geometry | 8 | 19 |
| 18 | Bit Manipulation | 6 | 20 |

## Architecture

```
internal/practice/
  problem.go            Problem, Harness, Hint, Solution, SelfCheck, TestCase
  category.go           categories in curriculum order, with their chapters
  problems_*.go         one file per category; each appends to EmbeddedProblems in init()
  loader.go             loads embedded problems plus optional JSON files, sorted by ladder
  problems/             optional custom *.json problems (embedded at build time)
internal/judge/         runs code for one test case and decides pass/fail
internal/sandbox/       sandbox.go (Python subprocess), tracer.py (steps), harness.py (calling convention)
internal/web/           practice_handlers.go (API); static/app.js and static/learn.js (UI)
```

### How a test case is judged

1. `judge.Spec` turns the problem's `Harness` and the test's input map into a
   harness spec: function name and parameters (from the starter code), node
   conversions, design or round-trip mode, an in-place parameter, a prelude,
   and a validator.
2. The sandbox runs the learner's code, then `harness.py` builds the
   arguments (for example a `ListNode` chain from `[1,2,3]`), calls the code,
   and converts the answer back to JSON.
3. The judge compares the result with `Expected`. The comparison is `exact` by
   default, `unordered`, or `unordered-nested`, with numbers matching within
   1e-5. When the problem has a `Validator`, the validator's verdict is used
   instead.

Submit runs without tracing, which keeps it fast. Run and the solution
visualizer trace the first 1000 steps; the code still finishes, and its
answer is still checked.

The harness options are listed in [pedagogy.md](pedagogy.md#harness-how-the-judge-calls-the-code).

## API

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/practice/problems` | GET | Problem summaries in ladder order (`{problems, total, stats}`) |
| `/api/practice/problems/{id}` | GET | Problem with hints and self-check. Hidden tests show only their `note`. |
| `/api/practice/categories` | GET | Categories in curriculum order, with problem counts |
| `/api/practice/run` | POST | `{problemId, code, testCaseIndex, language}`: one visible case, traced when the language is Python |
| `/api/practice/submit` | POST | `{problemId, code, language}`: every case. Returns `{status, passedCount, totalCount, results}`. |
| `/api/practice/hints/{id}/{level}` | GET | A single hint |
| `/api/practice/solution/{id}` | GET | The full solution |
| `/api/practice/solution-viz/{id}` | POST | The reference solution traced on a visible case |
| `/api/practice/progress/export`, `/import` | GET/POST | Progress JSON |

## Adding a problem

Follow [pedagogy.md](pedagogy.md#writing-a-practice-problem). Add the problem
to its category's `problems_*.go` file with the next `Order`, then run:

```sh
DSATUTOR_CATEGORY=<category> go test ./internal/contentcheck/
go test ./internal/judge/ -run 'TestReferenceSolutions/<category>/'
```

Custom problems can also go in `internal/practice/problems/*.json` using the
JSON form of `practice.Problem`. These aren't held to the content checks.
