package practice

// BacktrackingProblems contains all backtracking category problems, in ladder
// order. Every solution follows the same plan from chapter 12: record a
// complete answer, prune bad choices, choose, explore, un-choose.
var BacktrackingProblems = []*Problem{
	{
		ID:              "subsets",
		Number:          78,
		Order:           1,
		Title:           "Subsets",
		Slug:            "subsets",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking", "Bit Manipulation"},
		RelatedChapters: []int{12, 8},
		Description: `You are given an array ` + "`nums`" + ` of **distinct** integers. Return every possible subset of ` + "`nums`" + ` (the *power set*).

A subset is any selection of the elements, including the empty selection and the whole array. The order of elements inside a subset does not matter, so ` + "`[1,2]`" + ` and ` + "`[2,1]`" + ` are the same subset and must not both appear. Return the subsets in any order.`,
		Constraints: []string{
			"1 <= nums.length <= 10",
			"-10 <= nums[i] <= 10",
			"All the numbers in nums are distinct",
		},
		Examples: []Example{
			{Input: "nums = [1,2,3]", Output: "[[],[1],[1,2],[1,2,3],[1,3],[2],[2,3],[3]]", Explanation: "Each of the 3 elements is either in or out, so there are 2^3 = 8 subsets."},
			{Input: "nums = [0]", Output: "[[],[0]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}}, Expected: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}, Hidden: false, Note: "three elements give 2^3 = 8 subsets"},
			{Input: map[string]interface{}{"nums": []int{0}}, Expected: [][]int{{}, {0}}, Hidden: false, Note: "single element: the empty set and itself"},
			{Input: map[string]interface{}{"nums": []int{5, -3}}, Expected: [][]int{{}, {5}, {5, -3}, {-3}}, Hidden: false, Note: "negative numbers are ordinary elements"},
			{Input: map[string]interface{}{"nums": []int{4, 1, 0, 9}}, Expected: [][]int{{}, {4}, {4, 1}, {4, 1, 0}, {4, 1, 0, 9}, {4, 1, 9}, {4, 0}, {4, 0, 9}, {4, 9}, {1}, {1, 0}, {1, 0, 9}, {1, 9}, {0}, {0, 9}, {9}}, Hidden: true, Note: "unsorted input, 16 subsets"},
			{Input: map[string]interface{}{"nums": []int{-1, 2, -3, 4, -5}}, Expected: [][]int{{}, {-1}, {-1, 2}, {-1, 2, -3}, {-1, 2, -3, 4}, {-1, 2, -3, 4, -5}, {-1, 2, -3, -5}, {-1, 2, 4}, {-1, 2, 4, -5}, {-1, 2, -5}, {-1, -3}, {-1, -3, 4}, {-1, -3, 4, -5}, {-1, -3, -5}, {-1, 4}, {-1, 4, -5}, {-1, -5}, {2}, {2, -3}, {2, -3, 4}, {2, -3, 4, -5}, {2, -3, -5}, {2, 4}, {2, 4, -5}, {2, -5}, {-3}, {-3, 4}, {-3, 4, -5}, {-3, -5}, {4}, {4, -5}, {-5}}, Hidden: true, Note: "five elements, 32 subsets: forgetting to copy path gives empty lists"},
			{Input: map[string]interface{}{"nums": []int{10, -10, 7, 3, 1, 2}}, Expected: [][]int{{}, {10}, {10, -10}, {10, -10, 7}, {10, -10, 7, 3}, {10, -10, 7, 3, 1}, {10, -10, 7, 3, 1, 2}, {10, -10, 7, 3, 2}, {10, -10, 7, 1}, {10, -10, 7, 1, 2}, {10, -10, 7, 2}, {10, -10, 3}, {10, -10, 3, 1}, {10, -10, 3, 1, 2}, {10, -10, 3, 2}, {10, -10, 1}, {10, -10, 1, 2}, {10, -10, 2}, {10, 7}, {10, 7, 3}, {10, 7, 3, 1}, {10, 7, 3, 1, 2}, {10, 7, 3, 2}, {10, 7, 1}, {10, 7, 1, 2}, {10, 7, 2}, {10, 3}, {10, 3, 1}, {10, 3, 1, 2}, {10, 3, 2}, {10, 1}, {10, 1, 2}, {10, 2}, {-10}, {-10, 7}, {-10, 7, 3}, {-10, 7, 3, 1}, {-10, 7, 3, 1, 2}, {-10, 7, 3, 2}, {-10, 7, 1}, {-10, 7, 1, 2}, {-10, 7, 2}, {-10, 3}, {-10, 3, 1}, {-10, 3, 1, 2}, {-10, 3, 2}, {-10, 1}, {-10, 1, 2}, {-10, 2}, {7}, {7, 3}, {7, 3, 1}, {7, 3, 1, 2}, {7, 3, 2}, {7, 1}, {7, 1, 2}, {7, 2}, {3}, {3, 1}, {3, 1, 2}, {3, 2}, {1}, {1, 2}, {2}}, Hidden: true, Note: "six elements, 64 subsets: catches path reused instead of copied"},
		},
		TimeComplexity:  "O(n * 2^n)",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def subsets(nums):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered-nested"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you wrote the subsets of [1,2,3] down by hand without missing any or writing one twice, what rule would you follow for which element may come next after you've written 1, then 2?"},
			{Level: 2, Type: "insight", Content: "Build each subset one element at a time and only ever add elements that come *later* in the array than the last one you added. That rule makes every subset appear exactly once, in increasing index order. Every partial path is itself a valid subset, so you record it as soon as you reach it. This is backtracking: choose an element, explore everything that can follow it, then un-choose it and try the next.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: a result list, a path list, and a recursive function backtrack(start).\n2. Record a complete answer: every call appends a copy of path (every node of the decision tree is a subset).\n3. Choose: for each index i from start to the end, append nums[i] to path.\n4. Explore: call backtrack(i + 1) so only later elements can follow.\n5. Un-choose: pop nums[i] off path before trying i + 1."},
			{Level: 4, Type: "skeleton", Content: "def subsets(nums):\n    result = []\n    path = []\n\n    def backtrack(start):\n        result.append(____)          # record a copy, not path itself\n        for i in range(____, len(nums)):\n            path.append(nums[i])     # choose\n            backtrack(____)          # explore\n            ____                     # un-choose\n\n    backtrack(0)\n    return result"},
		},
		Solution: Solution{
			Code: `def subsets(nums):
    result = []
    path = []

    def backtrack(start):
        result.append(path[:])          # record: every node of the tree is a subset
        for i in range(start, len(nums)):
            path.append(nums[i])        # choose nums[i] as the next element
            backtrack(i + 1)            # explore: only later elements may follow
            path.pop()                  # un-choose: restore path for the next i

    backtrack(0)
    return result`,
			Explanation: "This is the first problem of the backtracking ladder and the plainest use of the chapter 12 template. We grow a path one element at a time; the start index means an element can only be followed by later elements, so each subset is generated exactly once, in index order. Every node of this decision tree (not just the leaves) is a subset, so we record the path on entry to every call. Later problems change only what counts as complete, which choices are allowed, and where the next choice may start.",
			BruteForce:  "Generate every ordering of every length (all arrangements), sort each one and drop duplicates with a set. For n = 10 that builds roughly 9.8 million arrangements to keep just 1,024 subsets. The waste is producing [2,1] after [1,2] already exists: order inside a subset doesn't matter, so arrangements are the wrong thing to enumerate.",
			KeyInsight:  "Fix an order: only allow a subset to be built in increasing index order. Then each subset has exactly one path in the tree, and the tree has exactly 2^n nodes, one per subset. No duplicates to filter, no wasted branches.",
			WhyItWorks:  "Invariant: when backtrack(start) is called, path holds a subset of nums[0..start-1] listed in increasing index order, and the call will record every subset that extends path with elements from nums[start..]. Proof sketch: any subset, written in index order as nums[i1], nums[i2], ..., is reached by choosing i1 at depth 1, i2 at depth 2 and so on, because each i(k+1) > i(k) is inside the loop range; and that is the only way to reach it, because the loop never goes back to a smaller index. Each of the 2^n subsets is copied once in O(n), so the total is O(n * 2^n). The include/exclude alternative below is the same tree drawn differently: it decides in/out for nums[0], then nums[1], ..., and records only at depth n.",
			CommonMistakes: []string{
				"Appending path instead of path[:]. Every entry then points at the same list, which is empty when the recursion finishes, so you return [[], [], ...]. Always record a copy.",
				"Recursing with backtrack(start + 1) instead of backtrack(i + 1). After choosing nums[i], the next element must come after i; start + 1 lets earlier skipped elements back in and produces duplicates like [2,1].",
				"Forgetting path.pop(). Without the un-choose step, elements pile up and later subsets contain leftovers from earlier branches.",
			},
			TimeComplexity:  "O(n * 2^n): 2^n subsets, each copied in O(n)",
			SpaceComplexity: "O(n) for the path and recursion depth, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the state", Explanation: "path holds the subset being built; result collects copies. The nested function can read and modify both.", CodeSnippet: "result = []\npath = []", LineStart: 2, LineEnd: 3},
				{Title: "Record a complete answer", Explanation: "Every node of the tree is a valid subset, including the empty one at the root, so we record on entry. path[:] makes a copy that later pops can't change.", CodeSnippet: "result.append(path[:])", LineStart: 6, LineEnd: 6},
				{Title: "Choose", Explanation: "Try each element from start onwards as the next one. Starting at start (not 0) is what keeps subsets in index order and stops duplicates.", CodeSnippet: "for i in range(start, len(nums)):\n    path.append(nums[i])", LineStart: 7, LineEnd: 8},
				{Title: "Explore", Explanation: "Recurse with i + 1: whatever we add next must come after the element we just chose.", CodeSnippet: "backtrack(i + 1)", LineStart: 9, LineEnd: 9},
				{Title: "Un-choose", Explanation: "Pop the element so path is exactly as it was before this choice; the next loop iteration starts from a clean state.", CodeSnippet: "path.pop()", LineStart: 10, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{Name: "Include / exclude decision tree", Code: `def subsets(nums):
    result = []
    path = []

    def decide(i):
        if i == len(nums):              # every element has been decided
            result.append(path[:])
            return
        path.append(nums[i])            # branch 1: include nums[i]
        decide(i + 1)
        path.pop()
        decide(i + 1)                   # branch 2: exclude nums[i]

    decide(0)
    return result`, TimeComplexity: "O(n * 2^n)", SpaceComplexity: "O(n)", Explanation: "The other common picture: at depth i decide whether nums[i] is in or out. The tree is binary with 2^n leaves, and only leaves are recorded. It generates the same subsets as the start-index loop; the loop version generalises more easily to Combination Sum and Subsets II, which is why the ladder uses it."},
				{Name: "Bitmask enumeration", Code: `def subsets(nums):
    n = len(nums)
    result = []
    for mask in range(1 << n):          # bit i of mask says "take nums[i]"
        result.append([nums[i] for i in range(n) if mask >> i & 1])
    return result`, TimeComplexity: "O(n * 2^n)", SpaceComplexity: "O(1) besides the output", Explanation: "Count from 0 to 2^n - 1; bit i of the counter says whether nums[i] is taken. No recursion, but it only works when every choice is a plain yes/no, so it doesn't extend to the later problems."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the start-index loop never produce both [1,2] and [2,1]?", Answer: "After choosing an element at index i, the recursion only offers indices greater than i. So 2 can never be followed by 1 (index 0 < index 1). Each subset can only be built in increasing index order, which gives it exactly one path."},
			{Prompt: "What would you change if you only wanted subsets of exactly size k (LeetCode 77, Combinations)?", Answer: "Record only when len(path) == k and return straight after recording, since nothing deeper can be size k. You can also prune: if the elements left (len(nums) - i) can't fill the remaining k - len(path) slots, stop the loop."},
			{Prompt: "Why is recording on entry to every call correct here, while Permutations records only at the leaves?", Answer: "Here a partial path is already a complete answer: any selection is a subset. In Permutations an answer must use every element, so only paths of full length are complete. 'What counts as complete?' is the first question to ask in every backtracking problem."},
		},
	},
	{
		ID:              "combination-sum",
		Number:          39,
		Order:           2,
		Title:           "Combination Sum",
		Slug:            "combination-sum",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking"},
		RelatedChapters: []int{12, 8},
		Description: `You are given an array ` + "`candidates`" + ` of **distinct** positive integers and a positive integer ` + "`target`" + `. Return every combination of candidates whose numbers add up to ` + "`target`" + `.

You may use the same candidate **any number of times**. Two combinations are the same if they use each number the same number of times (order doesn't matter), so return each distinct combination once. Return the combinations in any order; if none exist, return an empty list. The test data guarantees fewer than 150 combinations.`,
		Constraints: []string{
			"1 <= candidates.length <= 30",
			"2 <= candidates[i] <= 40",
			"All elements of candidates are distinct",
			"1 <= target <= 40",
		},
		Examples: []Example{
			{Input: "candidates = [2,3,6,7], target = 7", Output: "[[2,2,3],[7]]", Explanation: "2 + 2 + 3 = 7 uses 2 twice, which is allowed. 7 on its own also works. No other combination sums to 7."},
			{Input: "candidates = [2,3,5], target = 8", Output: "[[2,2,2,2],[2,3,3],[3,5]]"},
			{Input: "candidates = [2], target = 1", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"candidates": []int{2, 3, 6, 7}, "target": 7}, Expected: [][]int{{2, 2, 3}, {7}}, Hidden: false, Note: "2 is reused: [2,2,3] and [7]"},
			{Input: map[string]interface{}{"candidates": []int{2, 3, 5}, "target": 8}, Expected: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, Hidden: false, Note: "three combinations, one uses 2 four times"},
			{Input: map[string]interface{}{"candidates": []int{2}, "target": 1}, Expected: []interface{}{}, Hidden: false, Note: "no combination reaches the target"},
			{Input: map[string]interface{}{"candidates": []int{7, 3, 2}, "target": 18}, Expected: [][]int{{2, 2, 2, 2, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 3, 3}, {2, 2, 2, 2, 3, 7}, {2, 2, 2, 3, 3, 3, 3}, {2, 2, 7, 7}, {2, 3, 3, 3, 7}, {3, 3, 3, 3, 3, 3}}, Hidden: true, Note: "unsorted candidates: duplicates appear if you restart at 0"},
			{Input: map[string]interface{}{"candidates": []int{8, 7, 4, 3}, "target": 11}, Expected: [][]int{{3, 4, 4}, {3, 8}, {4, 7}}, Hidden: true, Note: "pruning must not cut off valid larger candidates"},
			{Input: map[string]interface{}{"candidates": []int{5}, "target": 5}, Expected: [][]int{{5}}, Hidden: true, Note: "a single candidate equal to the target"},
			{Input: map[string]interface{}{"candidates": []int{2, 4}, "target": 9}, Expected: []interface{}{}, Hidden: true, Note: "even candidates can never make an odd target"},
		},
		TimeComplexity:  "O(n^(t/m)) where t is the target and m the smallest candidate",
		SpaceComplexity: "O(t/m) recursion depth",
		StarterCode:     "def combinationSum(candidates, target):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered-nested"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In Subsets you recursed with i + 1 after choosing nums[i]. What single change to that call would let the same number be chosen again?"},
			{Level: 2, Type: "insight", Content: "Use the Subsets template with two changes. Recurse with i instead of i + 1, so a candidate may repeat but you still never go back to an earlier one (which would create [3,2,2] after [2,2,3]). And track the remaining target: a path is complete when remaining hits 0. Sort first, so the moment a candidate is bigger than what remains, every later one is too and you can stop the loop.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: sort candidates; path, result, and backtrack(start, remaining).\n2. Record a complete answer: if remaining == 0, append a copy of path and return.\n3. Prune: in the loop over i from start, break as soon as candidates[i] > remaining.\n4. Choose: append candidates[i].\n5. Explore: backtrack(i, remaining - candidates[i]); i, not i + 1, allows reuse.\n6. Un-choose: pop."},
			{Level: 4, Type: "skeleton", Content: "def combinationSum(candidates, target):\n    candidates = sorted(candidates)\n    result = []\n    path = []\n\n    def backtrack(start, remaining):\n        if ____:\n            result.append(path[:])\n            return\n        for i in range(start, len(candidates)):\n            if ____:\n                break\n            path.append(candidates[i])\n            backtrack(____, remaining - candidates[i])\n            path.pop()\n\n    backtrack(0, target)\n    return result"},
		},
		Solution: Solution{
			Code: `def combinationSum(candidates, target):
    candidates = sorted(candidates)
    result = []
    path = []

    def backtrack(start, remaining):
        if remaining == 0:                       # record a complete answer
            result.append(path[:])
            return
        for i in range(start, len(candidates)):
            if candidates[i] > remaining:        # prune: sorted, so the rest are too big too
                break
            path.append(candidates[i])           # choose
            backtrack(i, remaining - candidates[i])  # explore: i, not i + 1, allows reuse
            path.pop()                           # un-choose

    backtrack(0, target)
    return result`,
			Explanation: "Builds directly on Subsets (problem 1). The decision tree is the same start-index loop; the new idea is a numeric goal. We carry the remaining target down the tree, record when it reaches exactly 0, and prune any branch that would overshoot. Recursing with i (not i + 1) lets a candidate repeat while still keeping each combination in non-decreasing order, so it is generated once.",
			BruteForce:  "Try every multiset of candidates up to the longest possible length (target / smallest candidate), for example with combinations_with_replacement, and keep those whose sum is the target. For candidates [2,3,6,7] and target 7 it checks 4 + 10 + 20 = 34 multisets to find 2. The waste is continuing to extend [6, 6] or [7, 2] after the sum has already passed the target.",
			KeyInsight:  "Carry the remaining amount down the recursion. Because every candidate is positive, once a partial sum overshoots, no extension can come back down, so the whole subtree can be cut. With sorted candidates the cut becomes a single break: everything after a too-big candidate is bigger still.",
			WhyItWorks:  "Invariant: backtrack(start, remaining) is called with path in non-decreasing order, sum(path) + remaining == target, and it finds every way to finish path using candidates[start:] (with repeats). Each combination, sorted, has exactly one path: choose its smallest number first, then the next, never moving left, so nothing is produced twice. The break is safe because candidates are sorted and positive: if candidates[i] > remaining, every candidates[j] with j > i is too, and no number can make the remainder grow back. Depth is at most target / min(candidates), since each step subtracts at least the smallest candidate.",
			CommonMistakes: []string{
				"Recursing with i + 1. That forbids reuse, so [2,2,3] is never found. Use i: 'this candidate or a later one'.",
				"Looping from 0 instead of start. Then both [2,2,3] and [3,2,2] are built. The start index forces non-decreasing order and makes each combination unique.",
				"Using break without sorting first. On unsorted [7,3,2] the loop would break at 7 and never try 3 or 2. Either sort, or use continue instead of break.",
			},
			TimeComplexity:  "O(n^(t/m)) in the worst case: up to n choices at each of t/m levels (t = target, m = smallest candidate)",
			SpaceComplexity: "O(t/m) for the path and recursion, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the state", Explanation: "Sorting makes the break-prune valid. The remaining target replaces 'compute sum(path)' at every node.", CodeSnippet: "candidates = sorted(candidates)", LineStart: 2, LineEnd: 4},
				{Title: "Record a complete answer", Explanation: "remaining == 0 means path sums to exactly target. Return: adding more positive numbers could only overshoot.", CodeSnippet: "if remaining == 0:\n    result.append(path[:])\n    return", LineStart: 7, LineEnd: 9},
				{Title: "Prune", Explanation: "This candidate already overshoots, and all later ones are larger, so stop the loop entirely.", CodeSnippet: "if candidates[i] > remaining:\n    break", LineStart: 11, LineEnd: 12},
				{Title: "Choose, explore, un-choose", Explanation: "Add the candidate, recurse with the same i so it can be used again, then pop it to restore the path.", CodeSnippet: "path.append(candidates[i])\nbacktrack(i, remaining - candidates[i])\npath.pop()", LineStart: 13, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: every multiset", Code: `from itertools import combinations_with_replacement


def combinationSum(candidates, target):
    result = []
    longest = target // min(candidates)          # most numbers a sum can use
    for size in range(1, longest + 1):
        for combo in combinations_with_replacement(sorted(candidates), size):
            if sum(combo) == target:
                result.append(list(combo))
    return result`, TimeComplexity: "O(C(n + L, L) * L) where L = target / smallest candidate", SpaceComplexity: "O(L)", Explanation: "Enumerates every multiset of each size and sums it. Correct and simple, but it keeps building multisets long after their sum has passed the target."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does recursing with i (instead of i + 1) allow reuse without creating duplicate combinations?", Answer: "With i, the next choice may be the same candidate or a later one, never an earlier one. So every path is non-decreasing. A multiset like {2,2,3} has exactly one non-decreasing ordering, so it is built once. Reuse comes from 'same index allowed'; uniqueness comes from 'never go left'."},
			{Prompt: "What if candidates could include 0 or negative numbers?", Answer: "The search could become infinite: with 0 you can add it forever without changing the sum, and with negatives an overshoot can come back down, so neither the remaining == 0 return nor the break-prune is valid. You'd need a length limit or a different formulation. The problem forbids this by making every candidate at least 2."},
			{Prompt: "How is this different from the coin change problems in chapter 15?", Answer: "Coin change asks for a count or a minimum, which overlapping subproblems make suitable for dynamic programming. Here we must list every combination, so the output itself can be exponential and backtracking that builds each one is the natural fit."},
		},
	},
	{
		ID:              "permutations",
		Number:          46,
		Order:           3,
		Title:           "Permutations",
		Slug:            "permutations",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking"},
		RelatedChapters: []int{12, 8},
		Description: `You are given an array ` + "`nums`" + ` of **distinct** integers. Return every permutation of it: every way to arrange all of its elements in a row.

Here order matters, so ` + "`[1,2,3]`" + ` and ` + "`[3,2,1]`" + ` are different permutations. Each permutation uses every element exactly once. Return the permutations in any order.`,
		Constraints: []string{
			"1 <= nums.length <= 6",
			"-10 <= nums[i] <= 10",
			"All the integers of nums are distinct",
		},
		Examples: []Example{
			{Input: "nums = [1,2,3]", Output: "[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]", Explanation: "3 choices for the first slot, 2 for the second, 1 for the last: 3! = 6 arrangements."},
			{Input: "nums = [0,1]", Output: "[[0,1],[1,0]]"},
			{Input: "nums = [1]", Output: "[[1]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}}, Expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, Hidden: false, Note: "three elements give 3! = 6 orderings"},
			{Input: map[string]interface{}{"nums": []int{0, 1}}, Expected: [][]int{{0, 1}, {1, 0}}, Hidden: false, Note: "two elements, two orders"},
			{Input: map[string]interface{}{"nums": []int{1}}, Expected: [][]int{{1}}, Hidden: false, Note: "single element"},
			{Input: map[string]interface{}{"nums": []int{-1, 5, 0}}, Expected: [][]int{{-1, 5, 0}, {-1, 0, 5}, {5, -1, 0}, {5, 0, -1}, {0, -1, 5}, {0, 5, -1}}, Hidden: true, Note: "negatives and zero"},
			{Input: map[string]interface{}{"nums": []int{4, 3, 2, 1}}, Expected: [][]int{{4, 3, 2, 1}, {4, 3, 1, 2}, {4, 2, 3, 1}, {4, 2, 1, 3}, {4, 1, 3, 2}, {4, 1, 2, 3}, {3, 4, 2, 1}, {3, 4, 1, 2}, {3, 2, 4, 1}, {3, 2, 1, 4}, {3, 1, 4, 2}, {3, 1, 2, 4}, {2, 4, 3, 1}, {2, 4, 1, 3}, {2, 3, 4, 1}, {2, 3, 1, 4}, {2, 1, 4, 3}, {2, 1, 3, 4}, {1, 4, 3, 2}, {1, 4, 2, 3}, {1, 3, 4, 2}, {1, 3, 2, 4}, {1, 2, 4, 3}, {1, 2, 3, 4}}, Hidden: true, Note: "four elements, 24 orderings: a start index would miss most of them"},
			{Input: map[string]interface{}{"nums": []int{6, -2, 9, 0, 1}}, Expected: [][]int{{6, -2, 9, 0, 1}, {6, -2, 9, 1, 0}, {6, -2, 0, 9, 1}, {6, -2, 0, 1, 9}, {6, -2, 1, 9, 0}, {6, -2, 1, 0, 9}, {6, 9, -2, 0, 1}, {6, 9, -2, 1, 0}, {6, 9, 0, -2, 1}, {6, 9, 0, 1, -2}, {6, 9, 1, -2, 0}, {6, 9, 1, 0, -2}, {6, 0, -2, 9, 1}, {6, 0, -2, 1, 9}, {6, 0, 9, -2, 1}, {6, 0, 9, 1, -2}, {6, 0, 1, -2, 9}, {6, 0, 1, 9, -2}, {6, 1, -2, 9, 0}, {6, 1, -2, 0, 9}, {6, 1, 9, -2, 0}, {6, 1, 9, 0, -2}, {6, 1, 0, -2, 9}, {6, 1, 0, 9, -2}, {-2, 6, 9, 0, 1}, {-2, 6, 9, 1, 0}, {-2, 6, 0, 9, 1}, {-2, 6, 0, 1, 9}, {-2, 6, 1, 9, 0}, {-2, 6, 1, 0, 9}, {-2, 9, 6, 0, 1}, {-2, 9, 6, 1, 0}, {-2, 9, 0, 6, 1}, {-2, 9, 0, 1, 6}, {-2, 9, 1, 6, 0}, {-2, 9, 1, 0, 6}, {-2, 0, 6, 9, 1}, {-2, 0, 6, 1, 9}, {-2, 0, 9, 6, 1}, {-2, 0, 9, 1, 6}, {-2, 0, 1, 6, 9}, {-2, 0, 1, 9, 6}, {-2, 1, 6, 9, 0}, {-2, 1, 6, 0, 9}, {-2, 1, 9, 6, 0}, {-2, 1, 9, 0, 6}, {-2, 1, 0, 6, 9}, {-2, 1, 0, 9, 6}, {9, 6, -2, 0, 1}, {9, 6, -2, 1, 0}, {9, 6, 0, -2, 1}, {9, 6, 0, 1, -2}, {9, 6, 1, -2, 0}, {9, 6, 1, 0, -2}, {9, -2, 6, 0, 1}, {9, -2, 6, 1, 0}, {9, -2, 0, 6, 1}, {9, -2, 0, 1, 6}, {9, -2, 1, 6, 0}, {9, -2, 1, 0, 6}, {9, 0, 6, -2, 1}, {9, 0, 6, 1, -2}, {9, 0, -2, 6, 1}, {9, 0, -2, 1, 6}, {9, 0, 1, 6, -2}, {9, 0, 1, -2, 6}, {9, 1, 6, -2, 0}, {9, 1, 6, 0, -2}, {9, 1, -2, 6, 0}, {9, 1, -2, 0, 6}, {9, 1, 0, 6, -2}, {9, 1, 0, -2, 6}, {0, 6, -2, 9, 1}, {0, 6, -2, 1, 9}, {0, 6, 9, -2, 1}, {0, 6, 9, 1, -2}, {0, 6, 1, -2, 9}, {0, 6, 1, 9, -2}, {0, -2, 6, 9, 1}, {0, -2, 6, 1, 9}, {0, -2, 9, 6, 1}, {0, -2, 9, 1, 6}, {0, -2, 1, 6, 9}, {0, -2, 1, 9, 6}, {0, 9, 6, -2, 1}, {0, 9, 6, 1, -2}, {0, 9, -2, 6, 1}, {0, 9, -2, 1, 6}, {0, 9, 1, 6, -2}, {0, 9, 1, -2, 6}, {0, 1, 6, -2, 9}, {0, 1, 6, 9, -2}, {0, 1, -2, 6, 9}, {0, 1, -2, 9, 6}, {0, 1, 9, 6, -2}, {0, 1, 9, -2, 6}, {1, 6, -2, 9, 0}, {1, 6, -2, 0, 9}, {1, 6, 9, -2, 0}, {1, 6, 9, 0, -2}, {1, 6, 0, -2, 9}, {1, 6, 0, 9, -2}, {1, -2, 6, 9, 0}, {1, -2, 6, 0, 9}, {1, -2, 9, 6, 0}, {1, -2, 9, 0, 6}, {1, -2, 0, 6, 9}, {1, -2, 0, 9, 6}, {1, 9, 6, -2, 0}, {1, 9, 6, 0, -2}, {1, 9, -2, 6, 0}, {1, 9, -2, 0, 6}, {1, 9, 0, 6, -2}, {1, 9, 0, -2, 6}, {1, 0, 6, -2, 9}, {1, 0, 6, 9, -2}, {1, 0, -2, 6, 9}, {1, 0, -2, 9, 6}, {1, 0, 9, 6, -2}, {1, 0, 9, -2, 6}}, Hidden: true, Note: "five elements, 120 orderings"},
		},
		TimeComplexity:  "O(n * n!)",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def permute(nums):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In Subsets, after choosing nums[2] you could never choose nums[0]. For permutations, [3,1,2] is a valid answer. What must you track instead of a start index so that any unused element can come next?"},
			{Level: 2, Type: "insight", Content: "Order matters now, so every unused element is a legal next choice, including ones to the left. Replace the start index with a used[] array (or set). The loop runs over all indices and skips used ones. A path is complete only when it holds all n elements.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: path, result, used = [False] * n.\n2. Record a complete answer: when len(path) == n, append a copy and return.\n3. Prune: skip index i if used[i].\n4. Choose: mark used[i] and append nums[i].\n5. Explore: recurse.\n6. Un-choose: pop and clear used[i], in reverse order of choosing."},
			{Level: 4, Type: "skeleton", Content: "def permute(nums):\n    result = []\n    path = []\n    used = [False] * len(nums)\n\n    def backtrack():\n        if ____:\n            result.append(path[:])\n            return\n        for i in range(____):\n            if used[i]:\n                continue\n            ____\n            path.append(nums[i])\n            backtrack()\n            path.pop()\n            ____\n\n    backtrack()\n    return result"},
		},
		Solution: Solution{
			Code: `def permute(nums):
    result = []
    path = []
    used = [False] * len(nums)

    def backtrack():
        if len(path) == len(nums):      # record: every element is placed
            result.append(path[:])
            return
        for i in range(len(nums)):
            if used[i]:                 # prune: each element appears once
                continue
            used[i] = True              # choose
            path.append(nums[i])
            backtrack()                 # explore
            path.pop()                  # un-choose
            used[i] = False

    backtrack()
    return result`,
			Explanation: "Builds on Subsets (problem 1) with one change in which choices are legal. Subsets uses a start index because order doesn't matter; permutations care about order, so any element not yet in the path may come next. A used[] array records that, and both the path and used[] are restored on un-choose. Only full-length paths are recorded.",
			BruteForce:  "Generate all n^n sequences of length n (each slot any element), then keep the ones with no repeated element. For n = 6 that is 46,656 sequences to keep 720. The waste is filling later slots after a repeat has already happened in an earlier one.",
			KeyInsight:  "Check 'already used?' at the moment you choose, not at the end. A used[] array makes that check O(1), and it cuts every branch that would repeat an element, so the tree has exactly n! leaves.",
			WhyItWorks:  "Invariant: on entry to backtrack(), path holds distinct elements and used[i] is True exactly for the indices in path. The loop tries every unused index as the next slot, so every arrangement of the remaining elements is explored; un-choose restores both path and used[] so the invariant holds for the next iteration. Every permutation is reached by exactly one sequence of choices (its own order), so none is produced twice. There are n! leaves, each copied in O(n).",
			CommonMistakes: []string{
				"Copying the Subsets loop with range(start, n). That only builds increasing-index sequences, giving one arrangement instead of n!. Permutations need range(n) and a used[] check.",
				"Forgetting used[i] = False when un-choosing. The element stays 'used' for sibling branches, so most permutations are never built.",
				"Recording before the path is full. Only len(path) == n is a complete permutation; recording earlier adds partial arrangements to the answer.",
			},
			TimeComplexity:  "O(n * n!): n! permutations, each copied in O(n)",
			SpaceComplexity: "O(n) for path, used and the recursion, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the state", Explanation: "used[i] says whether nums[i] is already in the path. It replaces the start index from Subsets.", CodeSnippet: "used = [False] * len(nums)", LineStart: 4, LineEnd: 4},
				{Title: "Record a complete answer", Explanation: "A permutation must contain all n elements, so only full paths count.", CodeSnippet: "if len(path) == len(nums):\n    result.append(path[:])", LineStart: 7, LineEnd: 8},
				{Title: "Prune", Explanation: "Every index is a candidate, but one already in the path can't be used again.", CodeSnippet: "for i in range(len(nums)):\n    if used[i]:\n        continue", LineStart: 10, LineEnd: 12},
				{Title: "Choose, explore, un-choose", Explanation: "Mark and append, recurse to fill the next slot, then undo both in reverse order so the sibling iterations see the same state.", CodeSnippet: "used[i] = True\npath.append(nums[i])\nbacktrack()\npath.pop()\nused[i] = False", LineStart: 13, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{Name: "Library: itertools.permutations", Code: `from itertools import permutations


def permute(nums):
    return [list(p) for p in permutations(nums)]`, TimeComplexity: "O(n * n!)", SpaceComplexity: "O(n) besides the output", Explanation: "Python's standard library does exactly this enumeration. Fine in real code, but interviews expect you to write the recursion, and the library can't help once you need pruning (as in N-Queens)."},
				{Name: "Insert into every gap", Code: `def permute(nums):
    result = [[]]
    for num in nums:                    # insert num into every gap of every partial
        result = [p[:gap] + [num] + p[gap:] for p in result for gap in range(len(p) + 1)]
    return result`, TimeComplexity: "O(n * n!)", SpaceComplexity: "O(n * n!)", Explanation: "Iterative: the permutations of the first k elements, with element k+1 inserted into each of the k + 1 gaps, are exactly the permutations of the first k + 1. It keeps every partial result in memory at once."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does Subsets use a start index while Permutations uses a used[] array?", Answer: "In Subsets, order inside an answer doesn't matter, so we force one canonical order (increasing index) and never look left. In Permutations, order is the whole point, so every unused element is a legal next choice regardless of position, and we need to remember which ones are taken: that is used[]."},
			{Prompt: "What changes if nums may contain duplicates (LeetCode 47, Permutations II)?", Answer: "Sort nums, then skip index i if nums[i] == nums[i-1] and used[i-1] is False. That makes equal values always be used left to right, so swapping two equal values never creates a 'new' permutation. It's the same 'skip a repeated value at this level' idea as Subsets II."},
			{Prompt: "Why is the comparison mode for this problem 'unordered' rather than 'unordered-nested'?", Answer: "The list of permutations may come in any order, but the order inside each permutation is the answer itself: [1,2,3] and [3,2,1] are different permutations. Sorting the inner lists would make them look identical."},
		},
	},
	{
		ID:              "subsets-ii",
		Number:          90,
		Order:           4,
		Title:           "Subsets II",
		Slug:            "subsets-ii",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking", "Bit Manipulation"},
		RelatedChapters: []int{12, 8},
		Description: `You are given an integer array ` + "`nums`" + ` that **may contain duplicates**. Return every possible subset of it.

Equal values are interchangeable: choosing "the first 2" or "the second 2" gives the same subset, so each distinct subset (as a multiset of values) must appear only once. The order of elements inside a subset doesn't matter. Return the subsets in any order.`,
		Constraints: []string{
			"1 <= nums.length <= 10",
			"-10 <= nums[i] <= 10",
		},
		Examples: []Example{
			{Input: "nums = [1,2,2]", Output: "[[],[1],[1,2],[1,2,2],[2],[2,2]]", Explanation: "The subset [1,2] can be made with either 2, but it is listed once. With two equal 2s there are only 3 choices for 'how many 2s' (0, 1, 2), so 2 * 3 = 6 subsets, not 8."},
			{Input: "nums = [0]", Output: "[[],[0]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 2}}, Expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}, Hidden: false, Note: "one duplicate pair: 6 subsets, not 8"},
			{Input: map[string]interface{}{"nums": []int{0}}, Expected: [][]int{{}, {0}}, Hidden: false, Note: "single element"},
			{Input: map[string]interface{}{"nums": []int{2, 1, 2}}, Expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}, Hidden: true, Note: "duplicates not adjacent until sorted"},
			{Input: map[string]interface{}{"nums": []int{4, 4, 4, 1, 4}}, Expected: [][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}, {4, 4, 4, 4}}, Hidden: true, Note: "four copies of 4: only 5 choices for how many 4s"},
			{Input: map[string]interface{}{"nums": []int{3, 3}}, Expected: [][]int{{}, {3}, {3, 3}}, Hidden: true, Note: "all values equal"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}}, Expected: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}, Hidden: true, Note: "no duplicates: behaves like Subsets"},
			{Input: map[string]interface{}{"nums": []int{-1, 0, -1, 2}}, Expected: [][]int{{}, {-1}, {-1, -1}, {-1, -1, 0}, {-1, -1, 0, 2}, {-1, -1, 2}, {-1, 0}, {-1, 0, 2}, {-1, 2}, {0}, {0, 2}, {2}}, Hidden: true, Note: "negative duplicates"},
		},
		TimeComplexity:  "O(n * 2^n)",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def subsetsWithDup(nums):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered-nested"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Run your Subsets code on [1,2,2]. Which subsets come out twice, and at what point in the loop did the second copy start?"},
			{Level: 2, Type: "insight", Content: "Duplicates start when, at the same level of the tree, you choose the same value twice as 'the next element'. Sort nums so equal values sit together; then in the loop skip nums[i] when i > start and nums[i] == nums[i - 1]. The first copy of a value at each level explores every subset that uses it; later copies at the same level would only repeat that work. Deeper levels may still take the second copy, which is how [2,2] is still built.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: sort nums; path, result, backtrack(start).\n2. Record a complete answer: append a copy of path on every call.\n3. Prune: in the loop, skip i if i > start and nums[i] == nums[i - 1].\n4. Choose: append nums[i].\n5. Explore: backtrack(i + 1).\n6. Un-choose: pop."},
			{Level: 4, Type: "skeleton", Content: "def subsetsWithDup(nums):\n    nums = ____\n    result = []\n    path = []\n\n    def backtrack(start):\n        result.append(path[:])\n        for i in range(start, len(nums)):\n            if ____ and nums[i] == nums[i - 1]:\n                continue\n            path.append(nums[i])\n            backtrack(i + 1)\n            path.pop()\n\n    backtrack(0)\n    return result"},
		},
		Solution: Solution{
			Code: `def subsetsWithDup(nums):
    nums = sorted(nums)                 # equal values become neighbours
    result = []
    path = []

    def backtrack(start):
        result.append(path[:])          # record
        for i in range(start, len(nums)):
            if i > start and nums[i] == nums[i - 1]:
                continue                # prune: same value already tried at this level
            path.append(nums[i])        # choose
            backtrack(i + 1)            # explore
            path.pop()                  # un-choose

    backtrack(0)
    return result`,
			Explanation: "This is Subsets (problem 1) plus one new idea: removing duplicate branches. Sorting puts equal values next to each other; then, within one loop (one level of the decision tree), only the first copy of each value is used as the next element. Deeper levels can still pick the following copies, so [2,2] is generated, but [2] is generated only once.",
			BruteForce:  "Generate all 2^n index subsets, sort each one, and put it in a set of tuples to remove repeats. That is the bitmask alternative below. It is O(n * 2^n) like the optimal solution, but it still builds every duplicate (for [4,4,4,4] that is 16 subsets to keep 5) and needs extra memory for the set.",
			KeyInsight:  "Duplicates are born at one place: the same value chosen as the next element twice in the same loop. After sorting, 'same value' means 'same as the previous index', so i > start and nums[i] == nums[i - 1] identifies exactly those repeated branches.",
			WhyItWorks:  "Think of the subsets that start (at this level) with value v. Choosing the first copy of v at index i leads to backtrack(i + 1), which can still take every later copy of v, so it explores every subset whose next value is v. Choosing a later copy at index j > i leads to backtrack(j + 1), which sees a strict subset of those options. Its results are all already produced by the first copy's branch, so skipping it loses nothing. The condition i > start matters: at i == start this is the first option at this level, so taking it can't repeat a sibling, even if it equals nums[start - 1] (which was chosen one level up).",
			CommonMistakes: []string{
				"Skipping when nums[i] == nums[i - 1] without the i > start check. That also forbids taking a second copy deeper in the tree, so [2,2] and [1,2,2] are lost.",
				"Forgetting to sort. On [2,1,2] the two 2s aren't adjacent, the comparison with nums[i - 1] never fires, and duplicates come back.",
				"Removing duplicates with a set at the end and calling it done. It works, but it still generates every duplicate first; the skip removes the duplicate branches themselves.",
			},
			TimeComplexity:  "O(n * 2^n) in the worst case (all values distinct)",
			SpaceComplexity: "O(n) for the path and recursion, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the state", Explanation: "Sorting groups equal values so 'same value as the previous index' can be checked in O(1).", CodeSnippet: "nums = sorted(nums)", LineStart: 2, LineEnd: 2},
				{Title: "Record a complete answer", Explanation: "As in Subsets, every node is a subset, so record a copy on entry.", CodeSnippet: "result.append(path[:])", LineStart: 7, LineEnd: 7},
				{Title: "Prune", Explanation: "At this level, a value equal to the one just tried would only repeat that branch. i > start keeps the first option of each level available.", CodeSnippet: "if i > start and nums[i] == nums[i - 1]:\n    continue", LineStart: 9, LineEnd: 10},
				{Title: "Choose, explore, un-choose", Explanation: "Unchanged from Subsets: append, recurse with i + 1, pop.", CodeSnippet: "path.append(nums[i])\nbacktrack(i + 1)\npath.pop()", LineStart: 11, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: bitmask plus a set", Code: `def subsetsWithDup(nums):
    n = len(nums)
    seen = set()
    result = []
    for mask in range(1 << n):
        subset = tuple(sorted(nums[i] for i in range(n) if mask >> i & 1))
        if subset not in seen:          # a set of sorted tuples removes repeats
            seen.add(subset)
            result.append(list(subset))
    return result`, TimeComplexity: "O(n * 2^n)", SpaceComplexity: "O(n * 2^n) for the set", Explanation: "Enumerate all 2^n index subsets, sort each, and keep the first occurrence of each sorted tuple. It wastes work and memory on duplicates but is a good independent check."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Trace the tree for sorted [1,2,2] at the root level: which indices does the loop actually choose, and why?", Answer: "It chooses index 0 (value 1) and index 1 (the first 2), and skips index 2 because i > start (2 > 0) and nums[2] == nums[1]. Inside the branch that chose index 1, start is 2, so index 2 is the first option there and is taken, building [2,2]."},
			{Prompt: "Why is i > start the right condition, and not i > 0?", Answer: "The skip is about siblings: two choices in the same loop. i == start is the first choice of this loop, so it has no earlier sibling to duplicate; its equal neighbour at start - 1 was chosen one level up, as an ancestor. With i > 0 you would forbid that, and lose every subset with two equal values."},
		},
	},
	{
		ID:              "combination-sum-ii",
		Number:          40,
		Order:           5,
		Title:           "Combination Sum II",
		Slug:            "combination-sum-ii",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking"},
		RelatedChapters: []int{12, 8},
		Description: `You are given an array ` + "`candidates`" + ` of positive integers (which **may contain duplicates**) and a positive integer ` + "`target`" + `. Return every distinct combination of candidates that adds up to ` + "`target`" + `.

Each element of ` + "`candidates`" + ` may be used **at most once** in a combination (two equal values in the array may both be used, since they are different elements). Combinations with the same values are the same combination, whichever copies were used, and must appear only once. Return them in any order; return an empty list if there are none.`,
		Constraints: []string{
			"1 <= candidates.length <= 100",
			"1 <= candidates[i] <= 50",
			"1 <= target <= 30",
		},
		Examples: []Example{
			{Input: "candidates = [10,1,2,7,6,1,5], target = 8", Output: "[[1,1,6],[1,2,5],[1,7],[2,6]]", Explanation: "[1,1,6] uses both 1s. [1,7] can be made with either 1, but appears once."},
			{Input: "candidates = [2,5,2,1,2], target = 5", Output: "[[1,2,2],[5]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"candidates": []int{10, 1, 2, 7, 6, 1, 5}, "target": 8}, Expected: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}, Hidden: false, Note: "two 1s may both be used, but [1,7] appears once"},
			{Input: map[string]interface{}{"candidates": []int{2, 5, 2, 1, 2}, "target": 5}, Expected: [][]int{{1, 2, 2}, {5}}, Hidden: false, Note: "three 2s, used at most twice here"},
			{Input: map[string]interface{}{"candidates": []int{2}, "target": 1}, Expected: []interface{}{}, Hidden: false, Note: "no combination"},
			{Input: map[string]interface{}{"candidates": []int{1, 1, 1, 1}, "target": 2}, Expected: [][]int{{1, 1}}, Hidden: true, Note: "all equal: exactly one answer"},
			{Input: map[string]interface{}{"candidates": []int{3, 1, 3, 5, 1, 1}, "target": 8}, Expected: [][]int{{1, 1, 1, 5}, {1, 1, 3, 3}, {3, 5}}, Hidden: true, Note: "several duplicate values"},
			{Input: map[string]interface{}{"candidates": []int{4, 2}, "target": 6}, Expected: [][]int{{2, 4}}, Hidden: true, Note: "each element used once: [2,2,2] is not allowed"},
			{Input: map[string]interface{}{"candidates": []int{1, 2, 3, 4, 5}, "target": 10}, Expected: [][]int{{1, 2, 3, 4}, {1, 4, 5}, {2, 3, 5}}, Hidden: true, Note: "distinct values, many answers"},
		},
		TimeComplexity:  "O(n * 2^n)",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def combinationSum2(candidates, target):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered-nested"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You've solved Combination Sum (reuse allowed) and Subsets II (duplicates in the input). Which line from each do you need here?"},
			{Level: 2, Type: "insight", Content: "Combine them. From Combination Sum: carry the remaining target, record at 0, and break once a sorted candidate exceeds it. From Subsets II: sort, then skip candidates[i] when i > start and it equals candidates[i - 1]. And because each element is used at most once, recurse with i + 1, not i.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: sort; path, result, backtrack(start, remaining).\n2. Record a complete answer: remaining == 0, append a copy and return.\n3. Prune: skip a repeated value at this level (i > start and equal to previous); break when candidates[i] > remaining.\n4. Choose: append candidates[i].\n5. Explore: backtrack(i + 1, remaining - candidates[i]).\n6. Un-choose: pop."},
			{Level: 4, Type: "skeleton", Content: "def combinationSum2(candidates, target):\n    candidates = sorted(candidates)\n    result = []\n    path = []\n\n    def backtrack(start, remaining):\n        if remaining == 0:\n            result.append(path[:])\n            return\n        for i in range(start, len(candidates)):\n            if ____:\n                continue\n            if ____:\n                break\n            path.append(candidates[i])\n            backtrack(____, remaining - candidates[i])\n            path.pop()\n\n    backtrack(0, target)\n    return result"},
		},
		Solution: Solution{
			Code: `def combinationSum2(candidates, target):
    candidates = sorted(candidates)
    result = []
    path = []

    def backtrack(start, remaining):
        if remaining == 0:                          # record
            result.append(path[:])
            return
        for i in range(start, len(candidates)):
            if i > start and candidates[i] == candidates[i - 1]:
                continue                            # prune: skip a repeated value at this level
            if candidates[i] > remaining:
                break                               # prune: the rest are larger still
            path.append(candidates[i])              # choose
            backtrack(i + 1, remaining - candidates[i])  # explore: i + 1, each item once
            path.pop()                              # un-choose

    backtrack(0, target)
    return result`,
			Explanation: "This problem adds no new technique: it combines two earlier rungs. The remaining-target bookkeeping and the sorted break come from Combination Sum (problem 2); the 'skip a repeated value at this level' rule comes from Subsets II (problem 4). The one difference from Combination Sum is recursing with i + 1, because each element may be used only once.",
			BruteForce:  "Try all 2^n subsets of the elements, keep those that sum to target, and remove repeats with a set of sorted tuples. With 100 candidates that is 2^100 subsets, which is hopeless. It wastes work twice: it keeps extending sums that have passed the target, and it rebuilds the same combination from different copies of equal values.",
			KeyInsight:  "Both kinds of waste can be cut at the moment of choosing. A sorted candidate bigger than the remainder ends the whole loop (break). A value equal to the previous sibling would only rebuild that sibling's combinations (continue).",
			WhyItWorks:  "Invariant: backtrack(start, remaining) has path non-decreasing, sum(path) + remaining == target, and finds every distinct way to finish using each of candidates[start:] at most once. Using i + 1 ensures each element appears at most once. The duplicate skip is safe for the same reason as in Subsets II: the first copy of a value at this level reaches every combination that the later copies could, because its subtree still has those later copies available. The break is safe because the list is sorted and every value is positive. Together each distinct combination is produced exactly once.",
			CommonMistakes: []string{
				"Recursing with i (copied from Combination Sum). That reuses a single element, so with candidates [4,2] and target 6 you'd wrongly return [2,2,2].",
				"Putting the duplicate skip as nums[i] == nums[i - 1] without i > start. Then [1,1,6] is lost because the second 1 can never be taken under the first.",
				"Writing break where continue is needed for the duplicate check. The next different value must still be tried, so a duplicate only skips this iteration.",
			},
			TimeComplexity:  "O(n * 2^n) in the worst case; pruning by target usually cuts far more",
			SpaceComplexity: "O(n) for the path and recursion, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the state", Explanation: "Sorting enables both prunes: equal values become neighbours and the break becomes valid.", CodeSnippet: "candidates = sorted(candidates)", LineStart: 2, LineEnd: 2},
				{Title: "Record a complete answer", Explanation: "The path sums to exactly target. Return, since adding any positive value would overshoot.", CodeSnippet: "if remaining == 0:\n    result.append(path[:])", LineStart: 7, LineEnd: 8},
				{Title: "Prune", Explanation: "Skip a value already tried at this level (continue); stop the loop when the value overshoots (break).", CodeSnippet: "if i > start and candidates[i] == candidates[i - 1]:\n    continue\nif candidates[i] > remaining:\n    break", LineStart: 11, LineEnd: 14},
				{Title: "Choose, explore, un-choose", Explanation: "i + 1 means this element is used up; equal values later in the array can still be chosen deeper.", CodeSnippet: "path.append(candidates[i])\nbacktrack(i + 1, remaining - candidates[i])\npath.pop()", LineStart: 15, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: every subset plus a set", Code: `def combinationSum2(candidates, target):
    n = len(candidates)
    seen = set()
    result = []
    for mask in range(1 << n):
        picked = [candidates[i] for i in range(n) if mask >> i & 1]
        if sum(picked) == target:
            combo = tuple(sorted(picked))
            if combo not in seen:
                seen.add(combo)
                result.append(list(combo))
    return result`, TimeComplexity: "O(n * 2^n)", SpaceComplexity: "O(n * 2^n)", Explanation: "Enumerates every subset of positions with a bitmask, keeps those with the right sum and removes repeats with a set of sorted tuples. Fine for the small tests here; with no pruning it can't handle 100 candidates."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Compare the recursive call here with Combination Sum and Subsets II. What does each choice of i or i + 1 mean?", Answer: "Combination Sum calls backtrack(i, ...): the same element may be chosen again (unlimited reuse). Subsets II and Combination Sum II call backtrack(i + 1, ...): the chosen element is used up, only later ones may follow. The duplicate skip is a separate concern: it prevents equal values from starting the same branch twice."},
			{Prompt: "Why must the duplicate check come before the break check, or does the order not matter?", Answer: "The order doesn't change the result: if candidates[i] equals candidates[i - 1] and the previous one didn't overshoot, this one doesn't either; if it does overshoot, both checks end the useful work. Putting the duplicate check first just reads naturally: 'is this a new option? is it small enough?'"},
		},
	},
	{
		ID:              "letter-combinations-of-a-phone-number",
		Number:          17,
		Order:           6,
		Title:           "Letter Combinations of a Phone Number",
		Slug:            "letter-combinations-of-a-phone-number",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Hash Table", "String", "Backtracking"},
		RelatedChapters: []int{12, 8, 2},
		Description: `On a phone keypad each digit from 2 to 9 stands for a few letters:

| digit | letters | digit | letters |
|---|---|---|---|
| 2 | abc | 6 | mno |
| 3 | def | 7 | pqrs |
| 4 | ghi | 8 | tuv |
| 5 | jkl | 9 | wxyz |

You are given a string ` + "`digits`" + ` containing only the characters 2-9. Return every string you can make by picking one letter for each digit, keeping the digits' order. Return them in any order. If ` + "`digits`" + ` is empty, return an empty list.`,
		Constraints: []string{
			"0 <= digits.length <= 4",
			"digits[i] is a digit in the range '2' to '9'",
		},
		Examples: []Example{
			{Input: `digits = "23"`, Output: `["ad","ae","af","bd","be","bf","cd","ce","cf"]`, Explanation: "3 letters for '2' times 3 letters for '3' gives 9 strings."},
			{Input: `digits = ""`, Output: "[]"},
			{Input: `digits = "2"`, Output: `["a","b","c"]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"digits": "23"}, Expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, Hidden: false, Note: "two digits with three letters each: 9 strings"},
			{Input: map[string]interface{}{"digits": ""}, Expected: []interface{}{}, Hidden: false, Note: "empty input returns [], not [\"\"]"},
			{Input: map[string]interface{}{"digits": "2"}, Expected: []string{"a", "b", "c"}, Hidden: false, Note: "a single digit"},
			{Input: map[string]interface{}{"digits": "79"}, Expected: []string{"pw", "px", "py", "pz", "qw", "qx", "qy", "qz", "rw", "rx", "ry", "rz", "sw", "sx", "sy", "sz"}, Hidden: true, Note: "7 and 9 have four letters each: 16 strings"},
			{Input: map[string]interface{}{"digits": "234"}, Expected: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}, Hidden: true, Note: "three digits: 27 strings"},
			{Input: map[string]interface{}{"digits": "9999"}, Expected: []string{"wwww", "wwwx", "wwwy", "wwwz", "wwxw", "wwxx", "wwxy", "wwxz", "wwyw", "wwyx", "wwyy", "wwyz", "wwzw", "wwzx", "wwzy", "wwzz", "wxww", "wxwx", "wxwy", "wxwz", "wxxw", "wxxx", "wxxy", "wxxz", "wxyw", "wxyx", "wxyy", "wxyz", "wxzw", "wxzx", "wxzy", "wxzz", "wyww", "wywx", "wywy", "wywz", "wyxw", "wyxx", "wyxy", "wyxz", "wyyw", "wyyx", "wyyy", "wyyz", "wyzw", "wyzx", "wyzy", "wyzz", "wzww", "wzwx", "wzwy", "wzwz", "wzxw", "wzxx", "wzxy", "wzxz", "wzyw", "wzyx", "wzyy", "wzyz", "wzzw", "wzzx", "wzzy", "wzzz", "xwww", "xwwx", "xwwy", "xwwz", "xwxw", "xwxx", "xwxy", "xwxz", "xwyw", "xwyx", "xwyy", "xwyz", "xwzw", "xwzx", "xwzy", "xwzz", "xxww", "xxwx", "xxwy", "xxwz", "xxxw", "xxxx", "xxxy", "xxxz", "xxyw", "xxyx", "xxyy", "xxyz", "xxzw", "xxzx", "xxzy", "xxzz", "xyww", "xywx", "xywy", "xywz", "xyxw", "xyxx", "xyxy", "xyxz", "xyyw", "xyyx", "xyyy", "xyyz", "xyzw", "xyzx", "xyzy", "xyzz", "xzww", "xzwx", "xzwy", "xzwz", "xzxw", "xzxx", "xzxy", "xzxz", "xzyw", "xzyx", "xzyy", "xzyz", "xzzw", "xzzx", "xzzy", "xzzz", "ywww", "ywwx", "ywwy", "ywwz", "ywxw", "ywxx", "ywxy", "ywxz", "ywyw", "ywyx", "ywyy", "ywyz", "ywzw", "ywzx", "ywzy", "ywzz", "yxww", "yxwx", "yxwy", "yxwz", "yxxw", "yxxx", "yxxy", "yxxz", "yxyw", "yxyx", "yxyy", "yxyz", "yxzw", "yxzx", "yxzy", "yxzz", "yyww", "yywx", "yywy", "yywz", "yyxw", "yyxx", "yyxy", "yyxz", "yyyw", "yyyx", "yyyy", "yyyz", "yyzw", "yyzx", "yyzy", "yyzz", "yzww", "yzwx", "yzwy", "yzwz", "yzxw", "yzxx", "yzxy", "yzxz", "yzyw", "yzyx", "yzyy", "yzyz", "yzzw", "yzzx", "yzzy", "yzzz", "zwww", "zwwx", "zwwy", "zwwz", "zwxw", "zwxx", "zwxy", "zwxz", "zwyw", "zwyx", "zwyy", "zwyz", "zwzw", "zwzx", "zwzy", "zwzz", "zxww", "zxwx", "zxwy", "zxwz", "zxxw", "zxxx", "zxxy", "zxxz", "zxyw", "zxyx", "zxyy", "zxyz", "zxzw", "zxzx", "zxzy", "zxzz", "zyww", "zywx", "zywy", "zywz", "zyxw", "zyxx", "zyxy", "zyxz", "zyyw", "zyyx", "zyyy", "zyyz", "zyzw", "zyzx", "zyzy", "zyzz", "zzww", "zzwx", "zzwy", "zzwz", "zzxw", "zzxx", "zzxy", "zzxz", "zzyw", "zzyx", "zzyy", "zzyz", "zzzw", "zzzx", "zzzy", "zzzz"}, Hidden: true, Note: "maximum length with repeated four-letter digit: 256 strings"},
		},
		TimeComplexity:  "O(n * 4^n)",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def letterCombinations(digits):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If digits always had exactly two characters, you'd write two nested loops. What do you do when the number of nested loops depends on the input?"},
			{Level: 2, Type: "insight", Content: "Let recursion be the nested loops: level i of the decision tree picks a letter for digits[i]. The choices at each level are just the letters for that digit, and a path is complete when it has one letter per digit. There is nothing to prune: every path is valid. Handle the empty input separately, because the recursion would otherwise record one empty string.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: return [] for empty input; a digit-to-letters map; path, result, backtrack(index).\n2. Record a complete answer: when index == len(digits), append ''.join(path) and return.\n3. Choose: for each letter of digits[index], append it.\n4. Explore: backtrack(index + 1).\n5. Un-choose: pop."},
			{Level: 4, Type: "skeleton", Content: "def letterCombinations(digits):\n    if not digits:\n        return ____\n    letters = {'2': 'abc', '3': 'def', '4': 'ghi', '5': 'jkl',\n               '6': 'mno', '7': 'pqrs', '8': 'tuv', '9': 'wxyz'}\n    result = []\n    path = []\n\n    def backtrack(index):\n        if ____:\n            result.append(''.join(path))\n            return\n        for letter in ____:\n            path.append(letter)\n            backtrack(index + 1)\n            path.pop()\n\n    backtrack(0)\n    return result"},
		},
		Solution: Solution{
			Code: `def letterCombinations(digits):
    if not digits:
        return []
    letters = {'2': 'abc', '3': 'def', '4': 'ghi', '5': 'jkl',
               '6': 'mno', '7': 'pqrs', '8': 'tuv', '9': 'wxyz'}
    result = []
    path = []

    def backtrack(index):
        if index == len(digits):           # record: one letter per digit
            result.append(''.join(path))
            return
        for letter in letters[digits[index]]:
            path.append(letter)            # choose
            backtrack(index + 1)           # explore the next digit
            path.pop()                     # un-choose

    backtrack(0)
    return result`,
			Explanation: "The tree has one level per digit, like Permutations (problem 3) has one level per slot, but the choices at a level come from a lookup table instead of the unused elements. Every root-to-leaf path is valid, so this is backtracking with no pruning: the template with the Prune step left out. It is easier than the duplicate-handling rungs; it sits here to show the template working on strings and a fixed-depth tree.",
			BruteForce:  "Write a fixed number of nested loops, one per digit. That works only for a known length; for lengths 0 to 4 you'd need five different versions. There's no wasted work to remove, since every combination is part of the answer. The problem is that the shape of the loops depends on the input.",
			KeyInsight:  "A recursion depth equal to len(digits) acts as a variable number of nested loops. Each call is one loop level; the path holds the current loop variables.",
			WhyItWorks:  "Invariant: backtrack(index) is called with path holding one letter for each of digits[0..index-1], and it produces every completion of path. At each level the loop tries every letter of digits[index], so all prod(len(letters[d])) strings are produced, and each exactly once because each is one distinct sequence of choices. The output has at most 4^n strings of length n, so the work is O(n * 4^n).",
			CommonMistakes: []string{
				"Returning [\"\"] for empty input. Without the early return, the base case fires immediately and records the empty string. The problem asks for [].",
				"Building strings with path + letter and also appending to path, so letters get added twice. Pick one: either pass a new string down (no un-choose needed) or append/pop a list.",
			},
			TimeComplexity:  "O(n * 4^n): up to 4^n strings of length n",
			SpaceComplexity: "O(n) for the path and recursion, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the state", Explanation: "The empty case is handled first. The map lists the choices available at each level.", CodeSnippet: "if not digits:\n    return []", LineStart: 2, LineEnd: 7},
				{Title: "Record a complete answer", Explanation: "One letter has been picked per digit; join the path into a string.", CodeSnippet: "if index == len(digits):\n    result.append(''.join(path))", LineStart: 10, LineEnd: 11},
				{Title: "Choose, explore, un-choose", Explanation: "Try every letter for this digit, recurse to the next digit, then pop to try the next letter.", CodeSnippet: "for letter in letters[digits[index]]:\n    path.append(letter)\n    backtrack(index + 1)\n    path.pop()", LineStart: 13, LineEnd: 16},
			},
			Alternatives: []AltSolution{
				{Name: "Cartesian product", Code: `from itertools import product


def letterCombinations(digits):
    if not digits:
        return []
    letters = {'2': 'abc', '3': 'def', '4': 'ghi', '5': 'jkl',
               '6': 'mno', '7': 'pqrs', '8': 'tuv', '9': 'wxyz'}
    return [''.join(combo) for combo in product(*(letters[d] for d in digits))]`, TimeComplexity: "O(n * 4^n)", SpaceComplexity: "O(n) besides the output", Explanation: "itertools.product computes exactly this cross product. Equivalent and shorter, but hides the recursion that the rest of the chapter relies on."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does this problem need no Prune step, when Combination Sum did?", Answer: "Every choice at every level leads to a valid answer: any letter for any digit is fine, and there's no target or conflict that can fail later. Pruning only helps when some partial paths can't possibly be completed."},
			{Prompt: "What if you passed path + letter as a new string to the recursive call instead of appending and popping?", Answer: "Then each call has its own string, so there is nothing to undo and no pop is needed. It costs O(n) to build each new string, which is the same order as the join. The append/pop version is the general template; the copy version is fine when the state is small."},
		},
	},
	{
		ID:              "palindrome-partitioning",
		Number:          131,
		Order:           7,
		Title:           "Palindrome Partitioning",
		Slug:            "palindrome-partitioning",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"String", "Backtracking", "Dynamic Programming"},
		RelatedChapters: []int{12, 8, 3},
		Description: `You are given a string ` + "`s`" + `. Cut it into pieces (keeping the letters in order, every letter in exactly one piece) so that **every piece is a palindrome**: it reads the same forwards and backwards. Return every way to do this.

Each answer is a list of the pieces in order. Return the answers in any order.`,
		Constraints: []string{
			"1 <= s.length <= 16",
			"s contains only lowercase English letters",
		},
		Examples: []Example{
			{Input: `s = "aab"`, Output: `[["a","a","b"],["aa","b"]]`, Explanation: `"a|a|b" and "aa|b" work. "a|ab" doesn't, because "ab" is not a palindrome.`},
			{Input: `s = "a"`, Output: `[["a"]]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "aab"}, Expected: [][]string{{"a", "a", "b"}, {"aa", "b"}}, Hidden: false, Note: "two partitions: a|a|b and aa|b"},
			{Input: map[string]interface{}{"s": "a"}, Expected: [][]string{{"a"}}, Hidden: false, Note: "single character"},
			{Input: map[string]interface{}{"s": "ab"}, Expected: [][]string{{"a", "b"}}, Hidden: false, Note: "no palindrome longer than one letter"},
			{Input: map[string]interface{}{"s": "aaa"}, Expected: [][]string{{"a", "a", "a"}, {"a", "aa"}, {"aa", "a"}, {"aaa"}}, Hidden: true, Note: "all equal: every cut pattern works"},
			{Input: map[string]interface{}{"s": "abba"}, Expected: [][]string{{"a", "b", "b", "a"}, {"a", "bb", "a"}, {"abba"}}, Hidden: true, Note: "even-length palindrome in the middle and as a whole"},
			{Input: map[string]interface{}{"s": "racecar"}, Expected: [][]string{{"r", "a", "c", "e", "c", "a", "r"}, {"r", "a", "cec", "a", "r"}, {"r", "aceca", "r"}, {"racecar"}}, Hidden: true, Note: "odd-length nested palindromes"},
			{Input: map[string]interface{}{"s": "abcbaba"}, Expected: [][]string{{"a", "b", "c", "b", "a", "b", "a"}, {"a", "b", "c", "b", "aba"}, {"a", "b", "c", "bab", "a"}, {"a", "bcb", "a", "b", "a"}, {"a", "bcb", "aba"}, {"abcba", "b", "a"}}, Hidden: true, Note: "overlapping palindromes"},
		},
		TimeComplexity:  "O(n * 2^n)",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def partition(s):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Every answer has a first piece. What are the possible first pieces of \"aab\", and which of them are allowed?"},
			{Level: 2, Type: "insight", Content: "The choice at each level is where the next piece ends. From position start, try every end; if s[start:end] is a palindrome, choose it and solve the rest of the string from end. When start reaches len(s), the whole string is cut into palindromes. Checking the palindrome before recursing is the pruning: a non-palindromic piece kills its entire subtree.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: path of pieces, result, backtrack(start).\n2. Record a complete answer: if start == len(s), append a copy of path.\n3. Prune: for each end from start + 1 to len(s), skip the piece s[start:end] if it isn't a palindrome.\n4. Choose: append the piece.\n5. Explore: backtrack(end).\n6. Un-choose: pop."},
			{Level: 4, Type: "skeleton", Content: "def partition(s):\n    result = []\n    path = []\n\n    def backtrack(start):\n        if ____:\n            result.append(path[:])\n            return\n        for end in range(start + 1, len(s) + 1):\n            piece = s[start:end]\n            if ____:\n                continue\n            path.append(piece)\n            backtrack(____)\n            path.pop()\n\n    backtrack(0)\n    return result"},
		},
		Solution: Solution{
			Code: `def partition(s):
    result = []
    path = []

    def backtrack(start):
        if start == len(s):                    # record: the whole string is cut up
            result.append(path[:])
            return
        for end in range(start + 1, len(s) + 1):
            piece = s[start:end]
            if piece != piece[::-1]:           # prune: this piece is not a palindrome
                continue
            path.append(piece)                 # choose the next piece
            backtrack(end)                     # explore the rest of the string
            path.pop()                         # un-choose

    backtrack(0)
    return result`,
			Explanation: "Like Subsets (problem 1), the loop picks where the next choice starts, but now a choice is a whole substring rather than one element: from position start, the next piece ends somewhere in start+1..n. The new idea is a validity check on each choice (is the piece a palindrome?) that prunes whole subtrees, preparing for the constraint checks in Word Search and N-Queens.",
			BruteForce:  "There are n - 1 gaps between letters, and each can be cut or not, so 2^(n-1) ways to cut the string. Try them all (the bitmask alternative) and keep those where every piece is a palindrome. For \"abcdefghijklmnop\" (16 distinct letters) that checks 32,768 cuttings to find one. The waste: once the first piece \"ab\" fails, every cutting that starts with it is still built and checked.",
			KeyInsight:  "Check each piece as soon as it is chosen. A non-palindromic first piece rules out every cutting that begins with it, so the search only ever extends prefixes that are fully valid.",
			WhyItWorks:  "Invariant: on entry to backtrack(start), path is a list of palindromes that together spell s[:start], and the call produces every valid way to cut s[start:]. The loop tries every possible next piece s[start:end], so no valid cutting is missed; each cutting corresponds to one sequence of end positions, so none is produced twice. When start == len(s), path spells all of s with palindromic pieces, which is exactly a complete answer. The worst case (all letters equal) makes every piece valid: 2^(n-1) answers, each costing O(n) to check and copy.",
			CommonMistakes: []string{
				"Using range(start, len(s)) for end with slicing s[start:end]. The first piece is then the empty string and the recursion never advances. Pieces are s[start:end] with end from start + 1 to len(s) inclusive.",
				"Recording when the path looks complete but start < len(s), for example when no further palindrome exists. Only start == len(s) means every letter has been used.",
			},
			TimeComplexity:  "O(n * 2^n): up to 2^(n-1) partitions, each checked and copied in O(n)",
			SpaceComplexity: "O(n) for the path and recursion, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Record a complete answer", Explanation: "start == len(s) means the pieces in path cover the whole string.", CodeSnippet: "if start == len(s):\n    result.append(path[:])", LineStart: 6, LineEnd: 7},
				{Title: "Prune", Explanation: "Try every possible end for the next piece; reject pieces that aren't palindromes before recursing.", CodeSnippet: "for end in range(start + 1, len(s) + 1):\n    piece = s[start:end]\n    if piece != piece[::-1]:\n        continue", LineStart: 9, LineEnd: 12},
				{Title: "Choose, explore, un-choose", Explanation: "Append the piece, solve the rest of the string starting at end, then pop it to try a longer piece.", CodeSnippet: "path.append(piece)\nbacktrack(end)\npath.pop()", LineStart: 13, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: every cutting", Code: `def partition(s):
    n = len(s)
    result = []
    for mask in range(1 << (n - 1)):           # bit i set: cut between s[i] and s[i + 1]
        pieces = []
        start = 0
        for i in range(n - 1):
            if mask >> i & 1:
                pieces.append(s[start:i + 1])
                start = i + 1
        pieces.append(s[start:])
        if all(p == p[::-1] for p in pieces):
            result.append(pieces)
    return result`, TimeComplexity: "O(n * 2^n)", SpaceComplexity: "O(n)", Explanation: "Each of the n - 1 gaps is cut or not, so a bitmask enumerates all 2^(n-1) cuttings. Keep the ones whose pieces are all palindromes. Same worst case as backtracking, but no pruning on typical inputs."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "How could you avoid re-checking the same substring for being a palindrome many times?", Answer: "Precompute a table is_pal[i][j] with dynamic programming: s[i..j] is a palindrome if s[i] == s[j] and s[i+1..j-1] is a palindrome (or the inner part is empty). That costs O(n^2) once, and every check during backtracking becomes O(1). Chapter 16 covers this kind of table."},
			{Prompt: "What is the 'start' parameter here, compared with the start index in Subsets?", Answer: "In both it marks where the next choice begins. In Subsets the choice is one element at index i, and the next starts at i + 1. Here the choice is a piece s[start:end], and the next starts at end. The pattern is the same: the parameter remembers how much of the input has been decided."},
		},
	},
	{
		ID:              "word-search",
		Number:          79,
		Order:           8,
		Title:           "Word Search",
		Slug:            "word-search",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Array", "String", "Backtracking", "Matrix"},
		RelatedChapters: []int{12, 8},
		Description: `You are given an ` + "`m x n`" + ` grid of letters ` + "`board`" + ` and a string ` + "`word`" + `. Return ` + "`true`" + ` if ` + "`word`" + ` can be traced on the board, and ` + "`false`" + ` otherwise.

To trace a word, start on a cell holding its first letter, then move step by step to a cell directly above, below, left or right (not diagonally) holding the next letter. A cell may be used **at most once** in the same word. Letters are case-sensitive: ` + "`'a'`" + ` and ` + "`'A'`" + ` are different.`,
		Constraints: []string{
			"m == board.length, n == board[i].length",
			"1 <= m, n <= 6",
			"1 <= word.length <= 15",
			"board and word consist of only lowercase and uppercase English letters",
		},
		Examples: []Example{
			{Input: `board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"`, Output: "true", Explanation: "Start at A (top left), go right to B and C, down to C, down to E, left to D."},
			{Input: `board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"`, Output: "true"},
			{Input: `board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"`, Output: "false", Explanation: "After A, B, C the only B nearby is the one already used."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"board": [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}, "word": "ABCCED"}, Expected: true, Hidden: false, Note: "path that turns corners"},
			{Input: map[string]interface{}{"board": [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}, "word": "SEE"}, Expected: true, Hidden: false, Note: "path going right then down"},
			{Input: map[string]interface{}{"board": [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}, "word": "ABCB"}, Expected: false, Hidden: false, Note: "would need to reuse the B cell"},
			{Input: map[string]interface{}{"board": [][]string{{"a"}}, "word": "a"}, Expected: true, Hidden: true, Note: "single cell that matches"},
			{Input: map[string]interface{}{"board": [][]string{{"a"}}, "word": "b"}, Expected: false, Hidden: true, Note: "single cell that does not match"},
			{Input: map[string]interface{}{"board": [][]string{{"a", "b"}, {"c", "d"}}, "word": "abdca"}, Expected: false, Hidden: true, Note: "word longer than a simple path allows"},
			{Input: map[string]interface{}{"board": [][]string{{"C", "A", "A"}, {"A", "A", "A"}, {"B", "C", "D"}}, "word": "AAB"}, Expected: true, Hidden: true, Note: "first path fails, so cells must be restored for the next start"},
			{Input: map[string]interface{}{"board": [][]string{{"a", "A"}}, "word": "aa"}, Expected: false, Hidden: true, Note: "letters are case-sensitive: 'A' does not match 'a'"},
			{Input: map[string]interface{}{"board": [][]string{{"a", "b"}, {"c", "d"}}, "word": "ad"}, Expected: false, Hidden: true, Note: "diagonal cells are not adjacent"},
		},
		TimeComplexity:  "O(m * n * 3^L) where L = len(word)",
		SpaceComplexity: "O(L) recursion depth",
		StarterCode:     "def exist(board, word):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose you're standing on a cell that matches word[i]. What do you need to know to decide whether the rest of the word can be traced from here, and what must you remember so you don't step back onto a cell you've already used?"},
			{Level: 2, Type: "insight", Content: "Backtracking on a grid: the state is (row, column, index into word). Choose = mark the current cell as used (overwrite it with '#'), explore = try the four neighbours with index + 1, un-choose = put the letter back. Prune as early as possible: off the board, already used, or wrong letter ends that branch. Restoring the letter matters because a failed path must not block cells for the next attempt.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: a dfs(r, c, i) that answers 'can word[i:] be traced starting at (r, c)?'.\n2. Record a complete answer: if i == len(word), return True.\n3. Prune: return False if (r, c) is off the board or board[r][c] != word[i] (a used cell holds '#', so it never matches).\n4. Choose: save the letter and overwrite it with '#'.\n5. Explore: try the four neighbours with i + 1.\n6. Un-choose: restore the letter, then return what you found.\nFinally try dfs from every cell."},
			{Level: 4, Type: "skeleton", Content: "def exist(board, word):\n    rows, cols = len(board), len(board[0])\n\n    def dfs(r, c, i):\n        if i == len(word):\n            return ____\n        if not (0 <= r < rows and 0 <= c < cols) or ____:\n            return False\n        saved = board[r][c]\n        board[r][c] = ____\n        found = (dfs(r + 1, c, i + 1) or dfs(r - 1, c, i + 1) or\n                 dfs(r, c + 1, i + 1) or dfs(r, c - 1, i + 1))\n        ____\n        return found\n\n    return any(dfs(r, c, 0) for r in range(rows) for c in range(cols))"},
		},
		Solution: Solution{
			Code: `def exist(board, word):
    rows, cols = len(board), len(board[0])

    def dfs(r, c, i):
        if i == len(word):                         # record: every letter matched
            return True
        if not (0 <= r < rows and 0 <= c < cols) or board[r][c] != word[i]:
            return False                           # prune: off the board, used, or wrong letter
        saved = board[r][c]
        board[r][c] = '#'                          # choose: mark the cell as used
        found = (dfs(r + 1, c, i + 1) or dfs(r - 1, c, i + 1) or
                 dfs(r, c + 1, i + 1) or dfs(r, c - 1, i + 1))  # explore
        board[r][c] = saved                        # un-choose: free the cell again
        return found

    for r in range(rows):
        for c in range(cols):
            if dfs(r, c, 0):
                return True
    return False`,
			Explanation: "The same choose / explore / un-choose template as the earlier rungs, moved onto a grid. The used[] array from Permutations (problem 3) becomes 'overwrite the cell with #', and the choices at each step are the four neighbours rather than array indices. The new idea is that we only need one answer: the search returns True as soon as one complete path is found, and the or-chain stops exploring the remaining directions.",
			BruteForce:  "List every simple path of length L on the grid (from every start cell, in every direction, without revisiting) and compare its letters with word. On a 6 x 6 board with L = 15 that is up to 36 * 4 * 3^14, about 690 million paths. The waste is extending a path after its letters already differ from the word.",
			KeyInsight:  "Compare letters while walking. A path can only continue while its letters match word[0..i], so a mismatch prunes the branch at its first wrong letter. Marking the current path's cells in place gives an O(1) 'already used?' check with no extra memory.",
			WhyItWorks:  "Invariant: when dfs(r, c, i) is called, the cells currently marked '#' are exactly the path that spelled word[0..i-1], ending next to (r, c). The call returns True iff word[i:] can be traced from (r, c) without using a marked cell. Marking before recursing prevents reuse within the path; restoring after recursing makes the cell available to other paths, so a failed attempt never hides a valid one. Each step has at most 3 new directions (the fourth is where we came from, which is marked), giving O(m * n * 3^L) in the worst case.",
			CommonMistakes: []string{
				"Not restoring the cell after exploring. A failed branch then leaves '#' marks behind, and a later start (or another branch) can't use those cells, giving false where the answer is true.",
				"Restoring only on the failure path, e.g. 'if dfs(...): return True' before the restore line. That leaves the board modified; it's harmless for this boolean answer but a bug the moment the board is reused (as in Word Search II). Compute found, restore, then return.",
				"Checking the bounds after reading board[r][c]. Negative indices wrap around in Python, so board[-1][c] silently reads the last row. Check 0 <= r < rows first.",
			},
			TimeComplexity:  "O(m * n * 3^L): every start cell, then at most 3 new directions per letter",
			SpaceComplexity: "O(L) for the recursion; the board is marked in place",
			Walkthrough: []WalkthroughStep{
				{Title: "Record a complete answer", Explanation: "All L letters have been matched along a path of distinct cells.", CodeSnippet: "if i == len(word):\n    return True", LineStart: 5, LineEnd: 6},
				{Title: "Prune", Explanation: "Off the board, already on the path (it holds '#'), or the wrong letter: this branch can't spell the word.", CodeSnippet: "if not (0 <= r < rows and 0 <= c < cols) or board[r][c] != word[i]:\n    return False", LineStart: 7, LineEnd: 8},
				{Title: "Choose", Explanation: "Mark the cell as part of the current path by overwriting it; remember the letter to restore it.", CodeSnippet: "saved = board[r][c]\nboard[r][c] = '#'", LineStart: 9, LineEnd: 10},
				{Title: "Explore", Explanation: "Try the four neighbours for the next letter. 'or' stops at the first direction that succeeds.", CodeSnippet: "found = (dfs(r + 1, c, i + 1) or dfs(r - 1, c, i + 1) or", LineStart: 11, LineEnd: 12},
				{Title: "Un-choose", Explanation: "Put the letter back before returning, whatever the result, so other paths can use this cell.", CodeSnippet: "board[r][c] = saved\nreturn found", LineStart: 13, LineEnd: 14},
			},
			Alternatives: []AltSolution{
				{Name: "Explicit path list", Code: `def exist(board, word):
    rows, cols = len(board), len(board[0])

    def walk(path):                                # path: list of (row, col) used so far
        if len(path) == len(word):
            return True
        if not path:
            options = [(r, c) for r in range(rows) for c in range(cols)]
        else:
            r, c = path[-1]
            options = [(r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)]
        for r, c in options:
            if 0 <= r < rows and 0 <= c < cols and (r, c) not in path \
                    and board[r][c] == word[len(path)]:
                if walk(path + [(r, c)]):
                    return True
        return False

    return walk([])`, TimeComplexity: "O(m * n * 3^L * L)", SpaceComplexity: "O(L^2)", Explanation: "Keeps the path as a list of coordinates and checks membership with 'in', copying the path at each step. No board mutation, so nothing to restore, but every check and copy costs O(L)."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does overwriting the cell with '#' work as a 'visited' check, and why must it be undone?", Answer: "No letter in word is '#', so a marked cell fails the board[r][c] != word[i] test: it acts like a wall for the rest of this path. But it should only be a wall for this path. Other paths, from other starts or via other directions, may legitimately use it, so the letter must be restored when this branch finishes."},
			{Prompt: "How could you reject impossible words before searching at all?", Answer: "Count letters: if word needs more of some letter than the board has, return False immediately. Another trick: if the last letter of word is rarer on the board than the first, search for the reversed word instead, which prunes earlier. Both keep the same worst case but help a lot in practice."},
			{Prompt: "How does Word Search II (chapter 10) extend this when you have thousands of words?", Answer: "Instead of running this search once per word, it puts all words in a trie and runs one DFS from each cell, walking the trie alongside the board. A branch is pruned as soon as no word has the current path as a prefix, and all words sharing a prefix are found in the same walk."},
		},
	},
	{
		ID:              "n-queens",
		Number:          51,
		Order:           9,
		Title:           "N-Queens",
		Slug:            "n-queens",
		Difficulty:      "Hard",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking"},
		RelatedChapters: []int{12, 8},
		Description: `Place ` + "`n`" + ` chess queens on an ` + "`n x n`" + ` board so that no two queens attack each other. A queen attacks every square in its row, its column, and both of its diagonals. Return **every** distinct placement.

Each placement is a list of ` + "`n`" + ` strings, one per row from top to bottom; in each string ` + "`'Q'`" + ` marks a queen and ` + "`'.'`" + ` an empty square. Return the placements in any order; if none exist, return an empty list.`,
		Constraints: []string{
			"1 <= n <= 9",
		},
		Examples: []Example{
			{Input: "n = 4", Output: `[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]`, Explanation: "There are exactly two ways to place 4 queens on a 4 x 4 board; they are mirror images."},
			{Input: "n = 1", Output: `[["Q"]]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 4}, Expected: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}, Hidden: false, Note: "two solutions, mirror images"},
			{Input: map[string]interface{}{"n": 1}, Expected: [][]string{{"Q"}}, Hidden: false, Note: "one queen on a 1x1 board"},
			{Input: map[string]interface{}{"n": 2}, Expected: []interface{}{}, Hidden: true, Note: "no solution exists"},
			{Input: map[string]interface{}{"n": 3}, Expected: []interface{}{}, Hidden: true, Note: "no solution exists: the diagonals always clash"},
			{Input: map[string]interface{}{"n": 5}, Expected: [][]string{{"Q....", "..Q..", "....Q", ".Q...", "...Q."}, {"Q....", "...Q.", ".Q...", "....Q", "..Q.."}, {".Q...", "...Q.", "Q....", "..Q..", "....Q"}, {".Q...", "....Q", "..Q..", "Q....", "...Q."}, {"..Q..", "Q....", "...Q.", ".Q...", "....Q"}, {"..Q..", "....Q", ".Q...", "...Q.", "Q...."}, {"...Q.", "Q....", "..Q..", "....Q", ".Q..."}, {"...Q.", ".Q...", "....Q", "..Q..", "Q...."}, {"....Q", ".Q...", "...Q.", "Q....", "..Q.."}, {"....Q", "..Q..", "Q....", "...Q.", ".Q..."}}, Hidden: true, Note: "ten solutions"},
			{Input: map[string]interface{}{"n": 6}, Expected: [][]string{{".Q....", "...Q..", ".....Q", "Q.....", "..Q...", "....Q."}, {"..Q...", ".....Q", ".Q....", "....Q.", "Q.....", "...Q.."}, {"...Q..", "Q.....", "....Q.", ".Q....", ".....Q", "..Q..."}, {"....Q.", "..Q...", "Q.....", ".....Q", "...Q..", ".Q...."}}, Hidden: true, Note: "four solutions: catches checking only one diagonal direction"},
		},
		TimeComplexity:  "O(n!)",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def solveNQueens(n):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Every row must hold exactly one queen. If you place queens one row at a time from the top, what do you need to know about the queens already placed to decide whether a square in the current row is safe?"},
			{Level: 2, Type: "insight", Content: "Place one queen per row, so rows never clash. A square (r, c) is attacked if its column is taken, or its diagonals are. All squares on one down-right diagonal share the same r - c; all squares on one down-left (anti) diagonal share the same r + c. Keep three sets (columns, r - c, r + c) for O(1) checks. Choose = add to all three; un-choose = remove from all three.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: queens (column per placed row), and sets cols, diag (r - c), anti (r + c).\n2. Record a complete answer: when row == n, turn queens into strings and append.\n3. Prune: skip a column if it, row - col or row + col is in its set.\n4. Choose: record the column and add to the three sets.\n5. Explore: backtrack(row + 1).\n6. Un-choose: pop the column and remove from the three sets."},
			{Level: 4, Type: "skeleton", Content: "def solveNQueens(n):\n    result = []\n    queens = []\n    cols, diag, anti = set(), set(), set()\n\n    def backtrack(row):\n        if row == n:\n            result.append(['.' * c + 'Q' + '.' * (n - c - 1) for c in queens])\n            return\n        for col in range(n):\n            if col in cols or ____ in diag or ____ in anti:\n                continue\n            queens.append(col)\n            cols.add(col); diag.add(____); anti.add(____)\n            backtrack(row + 1)\n            queens.pop()\n            cols.remove(col); diag.remove(____); anti.remove(____)\n\n    backtrack(0)\n    return result"},
		},
		Solution: Solution{
			Code: `def solveNQueens(n):
    result = []
    queens = []                       # queens[row] = column of the queen in that row
    cols, diag, anti = set(), set(), set()

    def backtrack(row):
        if row == n:                  # record: one queen in every row
            result.append(['.' * c + 'Q' + '.' * (n - c - 1) for c in queens])
            return
        for col in range(n):
            if col in cols or row - col in diag or row + col in anti:
                continue              # prune: this square is attacked
            queens.append(col)        # choose
            cols.add(col)
            diag.add(row - col)
            anti.add(row + col)
            backtrack(row + 1)        # explore the next row
            queens.pop()              # un-choose
            cols.remove(col)
            diag.remove(row - col)
            anti.remove(row + col)

    backtrack(0)
    return result`,
			Explanation: "The capstone of the ladder: the Permutations (problem 3) tree with a stronger Prune step. Placing one queen per row with distinct columns is exactly choosing a permutation of columns; the new idea is checking both diagonals in O(1) with the r - c and r + c keys, which cuts most of the n! branches long before they are complete.",
			BruteForce:  "Choose any n of the n^2 squares and test whether the queens attack each other: for n = 8 that is C(64, 8), about 4.4 billion placements. Even placing one queen per row and column (a permutation) leaves 8! = 40,320 full boards to test, most of which fail at the second or third row. The waste is completing boards whose first few queens already attack each other.",
			KeyInsight:  "Check attacks as each queen is placed, and make the check O(1). Squares on one down-right diagonal all share r - c (for example (0,1), (1,2), (2,3) all have r - c = -1); squares on one anti-diagonal share r + c ((0,2), (1,1), (2,0) all have 2). So 'is this diagonal taken?' is a set lookup.",
			WhyItWorks:  "Invariant: on entry to backtrack(row), rows 0..row-1 each hold one queen, no two attack each other, and cols/diag/anti contain exactly their columns, r - c values and r + c values. A new queen at (row, col) is safe iff col, row - col and row + col are all absent: rows can't clash (one per row), columns are checked directly, and two squares are on the same diagonal exactly when their r - c (or r + c) values are equal. Un-choose removes the same three keys, restoring the invariant for the next column. Every valid board has one queen per row, so placing row by row and trying every column at each row finds all of them, each once.",
			CommonMistakes: []string{
				"Checking only one diagonal direction (only r - c or only r + c). Boards with queens on the other diagonal slip through; for n = 6 you'd return far more than the 4 real solutions.",
				"Sharing one board list between solutions and appending it directly. Later changes overwrite every recorded solution. Build fresh strings when recording.",
				"Forgetting to remove a key from one of the sets on un-choose. The stale key blocks squares in sibling branches and solutions go missing.",
			},
			TimeComplexity:  "O(n!): row r has at most n - r safe columns, so at most n! placements are tried (pruning removes most); each real solution costs O(n^2) to build",
			SpaceComplexity: "O(n) for the sets, queens list and recursion, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the state", Explanation: "queens[row] is the column of that row's queen. The three sets answer 'is this column / diagonal / anti-diagonal taken?' in O(1).", CodeSnippet: "queens = []\ncols, diag, anti = set(), set(), set()", LineStart: 3, LineEnd: 4},
				{Title: "Record a complete answer", Explanation: "All n rows have a queen: build the strings for this board.", CodeSnippet: "if row == n:\n    result.append(['.' * c + 'Q' + '.' * (n - c - 1) for c in queens])", LineStart: 7, LineEnd: 8},
				{Title: "Prune", Explanation: "Skip a square whose column, r - c diagonal or r + c anti-diagonal already holds a queen.", CodeSnippet: "if col in cols or row - col in diag or row + col in anti:\n    continue", LineStart: 11, LineEnd: 12},
				{Title: "Choose", Explanation: "Place the queen: record its column and claim its column and both diagonals.", CodeSnippet: "queens.append(col)\ncols.add(col)\ndiag.add(row - col)\nanti.add(row + col)", LineStart: 13, LineEnd: 16},
				{Title: "Explore, then un-choose", Explanation: "Fill the next row; afterwards remove the queen and release all three keys so the next column starts clean.", CodeSnippet: "backtrack(row + 1)\nqueens.pop()\ncols.remove(col)\ndiag.remove(row - col)\nanti.remove(row + col)", LineStart: 17, LineEnd: 21},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: test every permutation", Code: `from itertools import permutations


def solveNQueens(n):
    result = []
    for placement in permutations(range(n)):   # one queen per row and column
        if len({r - c for r, c in enumerate(placement)}) == n and \
                len({r + c for r, c in enumerate(placement)}) == n:
            result.append(['.' * c + 'Q' + '.' * (n - c - 1) for c in placement])
    return result`, TimeComplexity: "O(n * n!)", SpaceComplexity: "O(n)", Explanation: "One queen per row and per column is a permutation of columns. Generate all n! of them and keep those with n distinct r - c values and n distinct r + c values. No pruning, so it always builds all n! boards; backtracking abandons most of them after two or three rows."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why do all squares on a down-right diagonal share the same r - c?", Answer: "Moving one step down-right adds 1 to both r and c, so r - c doesn't change. Moving down-left adds 1 to r and subtracts 1 from c, so r + c doesn't change. That's why one number identifies each diagonal, and a set of those numbers tells you which diagonals are taken."},
			{Prompt: "What changes for N-Queens II, which asks only for the number of solutions?", Answer: "Replace 'build the strings and append' with 'count += 1' (or return 1 from the base case and sum the recursive results). The queens list isn't needed at all; the three sets are enough."},
			{Prompt: "Why is there no need to check rows?", Answer: "The recursion places exactly one queen in each row (row is the recursion depth), so two queens can never share a row. Choosing the right decision for each level can remove whole constraints from the check."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, BacktrackingProblems...)
}
