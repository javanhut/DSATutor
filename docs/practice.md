# Practice Mode

Practice mode is where learners solve LeetCode-style problems in Python, or
in Oxigen (see [oxigen.md](oxigen.md)), after each chapter. Problems are grouped into categories, and each category is
a ladder ordered easy → hard. Every Blind 75 problem is included, plus 64
NeetCode 150 problems and 32 more that Google candidates reported in
2024–2026, for 171 in total (see [Google interview prep](#google-interview-prep)). How problems are written, and why, is
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

## Foundations

Before a ladder applies a structure, its **Foundations** block (at the top of
the category in the problem list, and above the ladder on the chapter page)
drills the basics on their own:

- **Trace by hand** (`static/trace-drills.js`): a freshly generated structure
  each time, and the learner clicks nodes to show what an algorithm does.
  - Trees: preorder, inorder, postorder, level order, the nodes a BST search
    or insert compares against, and naming which order produced a sequence.
  - Graphs: BFS and DFS visit order (neighbours in increasing order), any
    valid topological order, and the order Dijkstra settles nodes.
  - Linked lists: where prev and curr are after k iterations of an in-place
    reverse, where slow ends (middle), where Floyd's pointers meet in a cycle,
    and where trail stops when removing the nth node from the end.
  - Heaps (drawn as a tree and as its array): where a pushed value settles
    after sifting up, where the last value settles after a pop, which array
    bottom-up heapify produces (against sorting, pushing one at a time and
    sifting from the root forwards), and which value breaks the heap rule.
  - Tries: the nodes a search visits, what it returns (True, falls off the
    trie, or a prefix that is not a word), how many nodes an insert creates,
    and the order autocomplete returns words under a prefix.
  - Stacks: the order pops return values, which bracket closes a given
    opener, what a monotonic (next greater element) stack holds after index
    k, and the order it pops values.
  - Binary search: the indexes mid lands on, where lo and hi are after k
    iterations, what lower bound returns (with duplicates and the end), and
    the probes of a first-True search.
  - Sliding window: the fixed-size window with the best sum, where left is
    after shrinking (shortest sum >= target), and where left is and which
    window is recorded in longest-substring-without-repeats.
  - Two pointers: where l and r are after k moves of pair sum and container
    with most water, where the write pointer is when removing duplicates in
    place, and where a palindrome check fails.
  - Backtracking: the m-th result that subsets, combinations (start index)
    and permutations (used[]) record, and the first N-queens solution with
    row 0 fixed, with the dead ends undone in the replay.
  - 1-D DP: the houses in the best house-robber plan (dp table in the
    replay), the maximum subarray (Kadane), how many ways coin change has
    (with the loop-order trap as a wrong option), and any longest increasing
    subsequence.
  - Intervals (closed): which intervals merge into one block, a largest set
    open at once (sweep line), where the arrows go in minimum arrows, and
    which intervals are covered by another.
  - Arrays & hashing: which pair one-pass Two Sum returns, which two prefix
    sums give a range sum, which numbers start a consecutive run, and how
    many subarrays sum to k (missing the {0: 1} seed is a wrong option).
  - 2-D DP: a cheapest right/down path (the dp table fills in the replay),
    paths around blocked cells, any longest common subsequence, and 0/1
    knapsack (the unbounded and ratio-greedy answers are wrong options).
  - Greedy: a fewest-jumps route, the gas station start, partition labels,
    and the fewest coins, where largest-coin-first sometimes fails.
  - Bit manipulation: a number in binary, AND/OR/XOR, n & (n - 1) and
    n & -n, and shifts, clicking the 1 bits of the answer.
  - Math & geometry: gcd by Euclid, one pass of the sieve (from p²),
    spiral order, where a cell goes in a 90° rotation, and which side of a
    line a point is on (cross product).
  - Advanced graphs (distinct weights, so the tree is unique): the order
    Prim's adds nodes, the edges Kruskal's adds (clicking edge weights; the
    replay shows the components), the order Dijkstra settles nodes, and the
    cheapest trip with at most k stops (Bellman-Ford rounds on a copy; the
    unlimited and update-in-place answers are wrong options), the cells one
    Floyd-Warshall pass makes smaller, and the smallest-first Eulerian path
    (Reconstruct Itinerary style; the replay shows Hierholzer's call stack
    and the route built backwards). Choice questions with a `steps()` get a
    replay once answered, and a question with `repeats` lets a node be
    clicked more than once.

  A wrong answer says which step went wrong and why. **Show me step by
  step** replays the algorithm with its code line, call stack, queue, heap
  or pointers. Three first-try answers in a row master an exercise. Results
  are stored in the practice progress under `trace:<category>:<exercise>`,
  so Export, Import and Reset all include them. Each category's exercises
  are a "kit" in that file; the core runs any kit.
- **Drills**: judged problems with `Kind: "drill"` that build the structure or
  code one core algorithm on its own:
  - Trees (8, `problems_trees_drills.go`): build a tree from level order,
    preorder, inorder, postorder, level order, inorder without recursion,
    BST insert, BST delete.
  - Linked List (8, `problems_linked_list_drills.go`): build a list (writing
    `ListNode` too), list to array, index of a value, middle node, insert at
    an index, remove elements, Design Linked List, a doubly linked deque.
  - Graphs (`problems_graphs_drills.go`): adjacency list and matrix, BFS and
    DFS order, DFS without recursion, unweighted shortest path, flood fill,
    topological sort, a union-find class.
  - Advanced Graphs (7, `problems_advanced_graphs_drills.go`): Dijkstra
    costs, Dijkstra with the path rebuilt from parents (every test has a
    unique cheapest path), Prim's MST weight, Kruskal's accepted edges in order
    (ties by input order, a spanning forest when disconnected), Bellman-Ford
    with negative weights (`"inf"` for unreachable, `[]` for a reachable
    negative cycle, an unreachable one is ignored), Floyd-Warshall all pairs
    (-1 for unreachable) and an Eulerian path by Hierholzer, checked by a
    validator because many paths are valid. Only the Prim drill has a large
    case (5000 nodes, 100000 edges, 3 second limit) that a per-step scan of
    every edge fails; the rest check answers, not speed.
  - Heap (7, `problems_heap_drills.go`): build a min-heap with sift-up
    (design), pop with sift-down (design), bottom-up heapify, is this a
    min-heap, heap sort, a max-heap with `heapq` by negation, and a k-way
    merge with `heapq` tuples. The array drills compare exact arrays, so the
    tie rule (smaller child, left on a tie) is part of each statement; heap
    sort has a 50,000-item case with a 5 second limit, so an O(n^2) sort
    fails.
  - Stack (7, `problems_stack_drills.go`): build a stack (design), reverse
    each word, remove adjacent duplicates, queue using two stacks (design),
    next greater element, previous smaller element, evaluate an expression
    with parentheses.
  - Tries (6, `problems_tries_drills.go`): build nodes and count them, the
    end-of-word flag (contains and longest stored prefix), Implement Trie II
    (counts on nodes), autocomplete in alphabetical order, delete with
    pruning, longest common prefix.
  - Arrays & Hashing (6, `problems_arrays_hashing_drills.go`): a frequency
    map (pairs in any order), first unique character (count, then walk the
    string), many range sums with a prefix array, Subarray Sum Equals K
    (prefix sums plus a map of counts seeded with `{0: 1}`), count pairs with
    absolute difference k (a map of earlier counts, looked up before
    recording) and rotate array in place by three reversals. The large
    cases carry a 5 second limit and were measured to fail the naive
    versions (re-counting, re-summing each range, a double loop over pairs
    or subarrays, a pure Python double loop for first unique). The
    `s.count` version of first unique and the repeated `insert(0, pop())`
    rotation pass, so those are not claimed as caught.
  - Two Pointers (8, `problems_two_pointers_drills.go`): the three shapes on
    their own. Opposite ends: reverse a list in place, squares of a sorted
    array (fill from the back). Read/write pointers: remove duplicates from a
    sorted array (a validator checks the returned k and the kept front of the
    list), move zeroes (in place, stable). One pointer per sequence: merge
    two sorted arrays, is subsequence, intersection of two sorted arrays. Last,
    count pairs with a target sum in a sorted array with duplicates, the
    run-skipping that 3Sum needs. Move zeroes has a 100,000-item case and
    count pairs a 50,000-equal-items case, each with a 5 second limit; both
    were measured to fail the O(n^2) versions.
  - Binary Search (7, `problems_binary_search_drills.go`): one half-open
    boundary template (`while lo < hi`, `hi = mid` or `lo = mid + 1`) on search
    insert position, counting occurrences with a lower and an upper bound,
    first and last position, first True in a yes/no array, integer square
    root (searching the answer), peak element (compare with the neighbour)
    and minimum ship capacity (a feasibility check, the bridge to Koko). The
    cases check answers, so a linear scan over the array is not rejected;
    only the capacity drill has 50,000-package cases that fail a scan over
    every capacity.
  - Sliding Window (6, `problems_sliding_window_drills.go`): fixed window
    sums (add what enters, subtract what leaves), maximum average, distinct
    values per window (a count dictionary, deleting keys at zero), minimum
    size subarray sum (shrink while valid), longest subarray with sum at
    most k (shrink until valid), and subarray product less than k (counting
    `right - left + 1` windows per right edge). Recomputing a window from
    scratch fails the 5 second limit on the large cases.
  - Backtracking (6, `problems_backtracking_drills.go`): the choose, explore,
    un-choose skeleton on all binary strings (in counting order), letter case
    permutation (options per position), combinations (the start index),
    arrangements of k of n letters (a `used` list restored after each branch),
    subsets with a target sum (pruning with `break` on sorted positives) and
    counting routes through a grid (mark and unmark a visited cell). The
    40-number target-sum case has a 5 second limit that only a pruned search
    meets.
  - Intervals (6, `problems_intervals_drills.go`): all intervals are closed
    (`[1,3]` and `[3,5]` touch) unless a statement says otherwise. The overlap
    test `max(starts) <= min(ends)`, intersection of two sorted interval
    lists (two pointers, advance the one that ends first), total covered
    length (sort by start, one running block with `max` on the end), remove
    covered intervals (start ascending, end descending), most intervals at one
    point (sweep line, starts before ends on ties; the closed twin of Meeting
    Rooms II), and minimum arrows to burst balloons (sort by end). The
    cases check answers, not speed.
  - 1-D DP (6, `problems_1d_dp_drills.go`): each statement names the state
    and the recurrence. Fibonacci (memo, table or two variables), climb with
    1, 2 or 3 steps (write the recurrence from the last move), min cost
    climbing stairs (min over two choices, the answer at the top), count
    tilings of a 2 by n strip (a coefficient on one term, modulo
    1_000_000_007), maximum subarray (best sum ending at i, answer is the
    table's maximum), and Coin Change II (coins in the outer loop count
    combinations; swapped loops count orderings). Plain recursion fails the
    5 second limit on the large cases, and the tiling drill has n up to
    100000, where a recursive solution (memoised or not) hits the 20000
    recursion limit and only the bottom-up table passes.
  - 2-D DP (6, `problems_2d_dp_drills.go`): each statement names what
    `table[i][j]` means, the recurrence and the base row and column. Unique
    paths with obstacles (an obstacle in the first row blocks the rest of it),
    minimum path sum (a min of two neighbours; a zero border undercounts),
    triangle (filled from the bottom; the one-row version must sweep columns
    left to right), the longest repeated subarray (a mismatch resets to 0 and
    the answer is the largest cell, where Longest Common Subsequence carries
    values forward), the 0/1 knapsack (the one-row version sweeps capacity
    downwards; upwards is the unbounded knapsack) and counting all-ones square
    submatrices (min of above, left and diagonal, summed over corners). Plain
    recursion fails the 5 second limit on the large cases, and the repeated
    subarray drill has 1000-element arrays where checking every start pair
    also runs out of time.
  - Greedy (6, `problems_greedy_drills.go`): each explanation gives the exchange
    argument. Assign cookies (sort both, smallest cookie that fits), lemonade
    change (spend a 10 before three 5s), most meetings in one room (half-open
    intervals, sort by end; tests defeat earliest start and shortest first),
    a coin set checked against the true minimum (returns `[greedy, best]`, so
    greedy fails on `[1,3,4]`), boats to save people (pair heaviest with
    lightest) and maximum units on a truck (sort by units per box, not pile
    size or total). The cases check answers, not speed.
  - Bit Manipulation (8, `problems_bit_manipulation_drills.go`): a ladder
    from representation to tricks. Write a number in binary (`% 2`, `// 2`),
    read a binary string (shift and add), get, set, clear and toggle one bit,
    Power of Two (`n & (n - 1)`), the positions of the set bits (`n & -n`),
    Hamming Distance (XOR, then count), Single Number III (split on the lowest
    bit where the two answers differ) and subsets in bitmask order. Each
    explanation says why its identity holds. `bin`, `int(x, 2)` and `format`
    are not blocked in the sandbox, so the drills do not claim to forbid them,
    and the cases check answers, not speed.
  - Math & Geometry (8, `problems_math_drills.go`): reverse the digits of an
    integer (`% 10`, `// 10`, sign handled explicitly), GCD and LCM by Euclid,
    is-prime with the square-root bound, Count Primes by sieve, fast
    exponentiation with a modulus, transpose, diagonal traverse (cells grouped
    by `r + c`) and the cross-product orientation of three points (exact
    integers; the tests include slopes that floats cannot tell apart). Each
    explanation says why its identity holds. Large cases separate the real
    algorithm from counting down for a gcd, a prime test per number against a
    sieve (about 15 s locally) and an `exp`-step power loop; the largest
    exponent is 10^15 because the judge passes integers as JSON doubles.

  The "without recursion" drills include a hidden input deeper than Python's
  recursion limit, so only an explicit stack passes.

Drills have their own `Order` sequence in front of the ladder and may have
`Number` 0. They come back in Review like any solved problem but stay out of
Mixed Practice, mock interviews, the solved count and the chapter's mastery
score. After a drill is accepted, **Next drill** leads to the next one and
then to the first rung of the ladder. `judge/drills*_test.go` check that each
drill rejects the mistake it is meant to catch.

Every accepted submission is also saved, under
`dsatutor_accepted_solutions`: up to 10 per problem, newest first, identical
code not repeated, including solutions from mock interviews. The **My
Solutions** tab lists them with Load into editor (one undoable change),
Copy and Delete. Reset all keeps them, and Export includes them (format
version 2; Import still reads older files and merges solutions instead of
overwriting).

Progress lives in the browser's localStorage, under
`dsatutor_practice_progress` and `dsatutor_chapter_progress`, and can be
exported and imported as JSON.

## Categories

| # | Category | Problems | Taught in chapter |
|---|---|---|---|
| 1 | Arrays & Hashing | 14 | 2 |
| 2 | Two Pointers | 5 | 3 |
| 3 | Sliding Window | 9 | 4 |
| 4 | Stack | 9 | 5 |
| 5 | Binary Search | 9 | 6 |
| 6 | Linked List | 10 | 7 |
| 7 | Trees | 15 | 9 |
| 8 | Tries | 3 | 10 |
| 9 | Heap / Priority Queue | 11 | 11 |
| 10 | Backtracking | 10 | 12 |
| 11 | Graphs | 16 | 13 |
| 12 | Advanced Graphs | 6 | 14 |
| 13 | 1-D Dynamic Programming | 13 | 15 |
| 14 | 2-D Dynamic Programming | 11 | 16 |
| 15 | Greedy | 10 | 17 |
| 16 | Intervals | 5 | 18 |
| 17 | Math & Geometry | 9 | 19 |
| 18 | Bit Manipulation | 6 | 20 |

## Google interview prep

74 problems carry a `Google` tag: questions that candidates reported from
Google phone screens, onsites and online assessments in 2024–2026. Pick
**Topic → Google** (or a level or round) in the problem list to see only
those; they stay in their technique ladders, so each still links to the
chapter that teaches its pattern.

| Tag | Meaning |
|---|---|
| `Google` | reported at Google |
| `Google L3`, `Google L4`, `Google L5` | the level to aim it at. New problems use the level they were reported for; the 42 tagged existing problems follow difficulty (Easy → L3, Medium → L4, Hard → L5). |
| `Google Phone Screen`, `Google Onsite`, `Google OA` | the round it was reported in |

32 problems were added for this set: Longest Palindrome, Number of Unequal
Triplets, Logger Rate Limiter, Insert Delete GetRandom O(1), First Missing
Positive, Longest Continuous Increasing Subsequence, Longest Substring with
At Most K Distinct Characters, Design Hit Counter, Minimum String Length
After Removing Substrings, Number of Visible People in a Queue, Binary Tree
Right Side View, Zigzag Level Order Traversal, Delete Leaves With a Given
Value, Sum of Distances in Tree, Add Two Numbers II, Top K Frequent Words,
Find K Pairs with Smallest Sums, Minimum Cost to Connect Sticks, Meeting
Rooms III, The Skyline Problem, Number of Closed Islands, Satisfiability of
Equality Equations, Validate Binary Tree Nodes, Number of Islands II, Optimal
Account Balancing, Monotone Increasing Digits, Find Original Array From
Doubled Array, Game of Life, House Robber III, New 21 Game, Russian Doll
Envelopes and Longest Duplicate Substring.

Sources: public candidate reports compiled by
[lets-code.co.in](https://www.lets-code.co.in/previousyearcodingquestion/google-previous-year-coding-questions/),
[landthelab.com](https://landthelab.com/companies/google/rounds/phone-screen),
the [Hello Interview](https://www.hellointerview.com/guides/google/l4) level
guides and [onsites.fyi](https://www.onsites.fyi/blog/article/google-L4-software-engineer-interview-questions).
These are self-reported, so treat them as likely topics, not a guaranteed
question bank. Statements are written in our own words, as for every
problem.

## Architecture

```
internal/practice/
  problem.go            Problem, Harness, Hint, Solution, SelfCheck, TestCase
  category.go           categories in curriculum order, with their chapters
  problems_*.go         one file per category; each appends to EmbeddedProblems in init()
  problems_*_drills.go  a category's foundations drills (Kind "drill")
  loader.go             loads embedded problems plus optional JSON files, sorted by ladder
  problems/             optional custom *.json problems (embedded at build time)
internal/judge/         runs code for one test case and decides pass/fail
internal/sandbox/       sandbox.go (Python subprocess), tracer.py (steps), harness.py (calling convention)
internal/web/           practice_handlers.go (API); static/app.js and static/learn.js (UI);
                        static/trace-drills.js (Trace by hand)
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
