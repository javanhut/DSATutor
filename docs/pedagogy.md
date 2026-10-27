# How DSA Tutor Teaches

This is the authoring guide for chapters and practice problems. Each rule
below comes from learning-science findings that hold up across many studies.
`internal/contentcheck` enforces the checkable parts, and
`internal/judge` runs every reference solution through the same judge a
learner's Submit uses.

## The evidence, and what we do with it

| Finding | What it means here |
|---|---|
| **Worked examples beat unguided problem solving for novices** (Sweller & Cooper 1985; Atkinson et al. 2000) | Every chapter teaches its pattern through at least one fully worked example before asking the learner to solve anything. |
| **Fade the guidance as skill grows** (Renkl & Atkinson 2003; expertise reversal, Kalyuga et al. 2003) | Worked example → completion problem with blanks → practice ladder. Within a problem, hints fade from a question to a code skeleton. |
| **Subgoal labels** (Catrambone 1998; Margulieux, Morrison & Guzdial 2012, 2016, in programming) | Worked-example steps, walkthrough steps and the "plan" hint are labelled with the subgoal they achieve ("Grow the window", "Shrink until valid"). The same labels recur across a chapter, so the learner gets a reusable plan. |
| **Self-explanation** (Chi et al. 1989, 1994; Bisra et al. 2018 meta-analysis) and **elaborative interrogation** (Dunlosky et al. 2013) | Hint 1 is a question. Solutions explain *why*, with an invariant or correctness argument, not just *what*. Each problem ends with SelfCheck "why/what if" questions answered before the model answer is shown. |
| **Retrieval practice** (Roediger & Karpicke 2006; Dunlosky et al. 2013, "high utility") | Chapters end with CheckQuestions answered from memory, with the answer revealed only after an attempt. |
| **Spaced practice** (Cepeda et al. 2006; Dunlosky et al. 2013) | Solved problems come back in the Review queue. The first review is after 3 days, and each later interval grows about 2.5× (1.2× if rated Hard, 3.5× if Easy, back to 1 day if the solution was needed). CheckQuestions include questions about earlier chapters. |
| **Interleaving for telling similar problem types apart** (Rohrer & Taylor 2007; Brunmair & Richter 2019 meta-analysis) | Mixed Practice draws from every unlocked category and hides the category. The learner names the pattern before coding. Each chapter has at least one "which pattern fits?" question that contrasts it with earlier ones. |
| **Productive struggle / desirable difficulties** (Bjork 1994; Kapur 2016) | After the chapter, the learner attempts a problem first. The solution stays behind a soft gate until the learner has submitted once or used all hints. |
| **Trace and explain before writing** (Lister et al. 2004, 2009; PRIMM, Sentance et al. 2019) | Worked steps carry a concrete `State` trace. The visualizer replays any run line by line. CheckQuestions include "what does this code return / what is `left` after…" items. |
| **Specific, elaborated feedback** (Shute 2008; Hattie & Timperley 2007) | Every test case has a `Note` saying what it checks ("empty input", "all negative"). A failure shows expected vs actual and the validator's reason. |
| **Mastery and prerequisites** (Bloom 1968; Kulik et al. 1990) | Chapters list prerequisites. Each category is a ladder ordered easy → hard, where each step adds one new idea. |
| **Narrated, segmented animation** (Mayer 2009) | Storyboards are learner-paced (Step/Back), with narration per step and code highlighting synchronized to the visual. |

## Curriculum

Chapters are numbered in learning order. A chapter with a `Category` teaches
that practice category, and its problems link back to it.

| # | Chapter | Category | Prerequisites |
|---|---|---|---|
| 1 | Start Here: How to Learn Algorithms and Big O | — | — |
| 2 | Arrays and Hash Maps | arrays-hashing | 1 |
| 3 | Two Pointers | two-pointers | 2 |
| 4 | Sliding Window | sliding-window | 3 |
| 5 | Stacks and Monotonic Stacks | stack | 2 |
| 6 | Binary Search | binary-search | 1 |
| 7 | Linked Lists | linked-list | 2, 3 |
| 8 | Recursion and Divide and Conquer | — | 1 |
| 9 | Binary Trees and BSTs | trees | 7, 8 |
| 10 | Tries | tries | 9, 2 |
| 11 | Heaps and Priority Queues | heap-priority-queue | 9 |
| 12 | Backtracking | backtracking | 8 |
| 13 | Graphs: BFS, DFS, Topological Sort, Union-Find | graphs | 9, 8 |
| 14 | Shortest Paths and Spanning Trees | advanced-graphs | 13, 11 |
| 15 | Dynamic Programming in One Dimension | 1d-dp | 8 |
| 16 | Dynamic Programming on Grids and Strings | 2d-dp | 15 |
| 17 | Greedy Algorithms | greedy | 15 |
| 18 | Intervals | intervals | 17 |
| 19 | Math and Matrix Manipulation | math-geometry | 1 |
| 20 | Bit Manipulation | bit-manipulation | 1 |
| 21 | Beyond Interviews: KNN and MapReduce | — | 1 |

## Writing a chapter

A chapter page reads top to bottom as one lesson:

1. **Motivation**: a concrete problem solved the slow way, so the learner
   feels why the new idea is needed. Use real numbers, e.g. "with n = 100,000
   the nested loop does 5 billion comparisons".
2. **Objectives**: 3+ observable outcomes ("Trace…", "Explain why…",
   "Implement…").
3. **When to use it** (`WhenToUse`): 3+ cues in a problem statement that
   point to the pattern ("contiguous subarray/substring", "sorted input",
   "minimum/maximum over all windows").
4. **Concepts**: 2+. Each needs:
   - `WhyItWorks`: the invariant or argument, not a restatement.
   - `Intuition`: one analogy, with the mapping made explicit.
   - 3+ `CoreIdeas`.
   - 2+ `CommonMistakes`, each phrased as the mistake plus how to avoid it.
   - A code `Example`.
5. **Worked example(s)**: a full solution in 4+ steps. Each step has:
   - a short `Subgoal` label, reused across the chapter
   - an `Explanation` that says what to do and why
   - the `Code` lines
   - a concrete `State` trace on a small input
6. **Completion problem(s)**: a *different* problem solved with the same
   subgoals, with the key lines replaced by `____`. `Blanks` lists the answers
   in order; the UI accepts answers that differ only in whitespace.
7. **Storyboards**: keep the IDs listed in the chapter file's comment, since
   they have animators in `app.js`. Each needs 3+ `MemoryTips`, written as the
   subgoal checklist.
8. **CheckQuestions**: 5+ in total, mixing:
   - 2+ multiple choice. The distractors must be real misconceptions, and
     `Answer` is the index of the correct choice.
   - 1+ free recall, where `Explanation` is the model answer.
   - At least one "which pattern fits this problem?" question that contrasts
     this chapter with earlier ones.
   - At least one question about a prerequisite chapter.
   - At least one "trace this code" question.

Don't use the legacy `Exercises`, `Tutorials` or `Visualizers` fields.

## Writing a practice problem

Each category is a **ladder**: `Order` 1..n with no gaps. Order 1 is the
gentlest problem that uses the pattern in its plainest form. Each later
problem adds one new twist, and its `Explanation` names the earlier problem
it builds on.

### Statement

- Describe the task in your own words; don't copy LeetCode's text. Keep the
  same task, inputs, outputs and constraints, and keep `Number` as the
  LeetCode number so learners can find the original.
- The statement must be self-contained (premium problems included): define
  every term, say whether a graph is directed, and state what to return in
  edge cases.
- `Examples` include an explanation for at least the first one.

### Starter code

- A plain top-level function with LeetCode's function name and parameter
  names (`def twoSum(nums, target):`). Don't use `class Solution`.
- Design problems use the LeetCode class (`class MinStack:` with its
  methods).
- When nodes are involved, say what is provided:
  `# ListNode(val, next) is provided` / `# TreeNode(val, left, right) is provided`.
  For a custom `Node`, show its definition in a comment and put the real
  definition in `Harness.Prelude`.
- Allowed imports are `math`, `collections`, `heapq`, `bisect`, `itertools`,
  `functools`, `typing`, `string`, `operator`, and `random`. `List`, `Optional`,
  `Dict`, `Tuple`, `Set` and `Deque` are predefined.

### Harness (how the judge calls the code)

| Situation | Harness |
|---|---|
| Plain function | zero value; inputs map to parameters by name |
| Linked list in / out | `ParamTypes: {"head": "ListNode"}`, `ReturnType: "ListNode"` |
| List of linked lists | `ParamTypes: {"lists": "ListNode[]"}` |
| Cycle given by `pos` | `ParamTypes: {"head": "CycleList:pos"}` (pos stays an input key and is not passed) |
| Tree in / out | `ParamTypes: {"root": "TreeNode"}` (level order, `nil` for gaps), `ReturnType: "TreeNode"` |
| Node arguments given by value | `ParamTypes: {"p": "TreeNodeRef:root"}`; `ReturnType: "TreeNodeVal"` to compare a returned node by value |
| Graph of `Node`s | `ParamTypes: {"node": "GraphNode"}`, `ReturnType: "GraphNode"`, `Prelude` defines `Node` |
| Random-pointer list | `"RandomList"` both ways, `Prelude` defines `Node` |
| Modifies input, returns nothing | `InPlace: "matrix"` |
| Class with operations | `Mode: "design"`; inputs `{"operations": [...], "values": [[...], ...]}` where the first operation is the class name and its values are the constructor args; expected has `nil` for calls that return nothing |
| Encode/decode pair | `Mode: "roundtrip", Function: "encode,decode"` (+ `ClassName: "Codec"` when they're methods) |
| Answer in any order | `Compare: "unordered"` (list of items) or `"unordered-nested"` (also each inner list, e.g. group anagrams, subsets) |
| Many valid answers | `Validator`: trusted Python `def check(raw, args, result, expected):` returning a bool or `(bool, "reason")`. `raw` holds the original JSON inputs, `args` the converted arguments, `result` the JSON answer and `result_obj` the raw returned object. Copy problems must also check that no returned node is an input node. |

### Tests

- 6+ cases, with 2+ visible (the examples) and 2+ hidden, and every case has a
  `Note`.
- Cover each real edge case: empty input, single element, duplicates,
  negatives, all-equal values, and minimum/maximum constraint values when
  small enough.
- Add a test that catches each mistake listed in `CommonMistakes`.
- Answers must vary: if every case expects `true`, then `return True` passes.
- Inputs are `map[string]interface{}` keyed by parameter name.
- Expected values are verified by running both the solution and the
  brute-force alternative; the judge test does this automatically.

### Hints (exactly four, fading)

1. `nudge`: a question that points at the key observation without naming
   the technique ("What would you need to remember about each number you've
   already passed?").
2. `insight`: the observation itself, plus the pattern name and a pointer to
   the chapter (`ChapterRef`).
3. `plan`: numbered, subgoal-labelled steps in plain words, with no code.
4. `skeleton`: code with the key lines replaced by `____`. It must not be the
   full solution.

### Solution

- `Explanation`: one paragraph summarising the approach and what it builds on.
- `BruteForce`: the obvious approach, its complexity, and exactly where it
  wastes work.
- `KeyInsight`: the one observation that removes the waste.
- `WhyItWorks`: the invariant, or a short correctness argument. Greedy
  problems need an exchange argument; DP problems state the subproblem, the
  recurrence and why it's complete.
- `CommonMistakes`: 2+ bugs learners actually write, each with its fix.
- `Walkthrough`: 3+ steps with subgoal-label titles. `LineStart`/`LineEnd` are
  1-based lines of `Code`, and `CodeSnippet` lines must appear in that range.
- `Alternatives`: at least one with runnable `Code`, usually the brute force,
  so learners can compare. All alternatives must pass every test.
- Code is idiomatic Python with descriptive names, usually under 30 lines.
  Complexity is written without constants inside O(), e.g. O(1) and not O(26).

### SelfCheck

2+ questions answered after solving. Use why-questions ("Why is it safe to
move the shorter line?") and what-if transfers ("What changes if the array
isn't sorted?"), each with a full model answer.

## Checking your work

```sh
imlazy test                                    # everything
DSATUTOR_CATEGORY=graphs DSATUTOR_CHAPTER=13 \
  go test ./internal/contentcheck/             # content rules for one area
go test ./internal/judge/ -run 'TestReferenceSolutions/graphs/'
```
