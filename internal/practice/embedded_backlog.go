package practice

// backlogMeta captures metadata for a placeholder Blind 75 problem.
type backlogMeta struct {
	id         string
	title      string
	category   string
	difficulty string
	tags       []string
}

// Additional Blind 75 style problems added as lightweight placeholders so the full list is available in practice mode.
var backlogMetas = []backlogMeta{
	// Stack
	{id: "min-stack", title: "Min Stack", category: "stack", difficulty: "Medium", tags: []string{"Stack", "Design"}},
	{id: "evaluate-reverse-polish-notation", title: "Evaluate Reverse Polish Notation", category: "stack", difficulty: "Medium", tags: []string{"Stack"}},
	{id: "generate-parentheses", title: "Generate Parentheses", category: "stack", difficulty: "Medium", tags: []string{"Stack", "Backtracking"}},
	{id: "daily-temperatures", title: "Daily Temperatures", category: "stack", difficulty: "Medium", tags: []string{"Monotonic Stack"}},
	{id: "car-fleet", title: "Car Fleet", category: "stack", difficulty: "Medium", tags: []string{"Monotonic Stack"}},
	{id: "largest-rectangle-in-histogram", title: "Largest Rectangle in Histogram", category: "stack", difficulty: "Hard", tags: []string{"Monotonic Stack"}},
	// Binary Search
	{id: "search-a-2d-matrix", title: "Search a 2D Matrix", category: "binary-search", difficulty: "Medium"},
	{id: "koko-eating-bananas", title: "Koko Eating Bananas", category: "binary-search", difficulty: "Medium"},
	{id: "find-minimum-in-rotated-sorted-array", title: "Find Minimum in Rotated Sorted Array", category: "binary-search", difficulty: "Medium"},
	{id: "search-in-rotated-sorted-array", title: "Search in Rotated Sorted Array", category: "binary-search", difficulty: "Medium"},
	{id: "time-based-key-value-store", title: "Time Based Key-Value Store", category: "binary-search", difficulty: "Medium", tags: []string{"Design"}},
	{id: "median-of-two-sorted-arrays", title: "Median of Two Sorted Arrays", category: "binary-search", difficulty: "Hard"},
	// Linked List
	{id: "merge-two-sorted-lists", title: "Merge Two Sorted Lists", category: "linked-list", difficulty: "Easy"},
	{id: "reorder-list", title: "Reorder List", category: "linked-list", difficulty: "Medium"},
	{id: "remove-nth-node-from-end", title: "Remove Nth Node From End", category: "linked-list", difficulty: "Medium"},
	{id: "copy-list-with-random-pointer", title: "Copy List with Random Pointer", category: "linked-list", difficulty: "Medium"},
	{id: "add-two-numbers", title: "Add Two Numbers", category: "linked-list", difficulty: "Medium"},
	{id: "merge-k-sorted-lists", title: "Merge K Sorted Lists", category: "linked-list", difficulty: "Hard"},
	{id: "lru-cache", title: "LRU Cache", category: "linked-list", difficulty: "Medium", tags: []string{"Design"}},
	// Trees
	{id: "same-tree", title: "Same Tree", category: "trees", difficulty: "Easy"},
	{id: "subtree-of-another-tree", title: "Subtree of Another Tree", category: "trees", difficulty: "Easy"},
	{id: "lowest-common-ancestor-of-a-bst", title: "Lowest Common Ancestor of a BST", category: "trees", difficulty: "Medium"},
	{id: "binary-tree-level-order-traversal", title: "Binary Tree Level Order Traversal", category: "trees", difficulty: "Medium"},
	{id: "validate-bst", title: "Validate BST", category: "trees", difficulty: "Medium"},
	{id: "kth-smallest-in-bst", title: "Kth Smallest in BST", category: "trees", difficulty: "Medium"},
	{id: "construct-binary-tree-from-preorder-and-inorder", title: "Construct Binary Tree from Preorder and Inorder Traversal", category: "trees", difficulty: "Medium"},
	{id: "serialize-and-deserialize-binary-tree", title: "Serialize and Deserialize Binary Tree", category: "trees", difficulty: "Hard", tags: []string{"Design"}},
	// Tries
	{id: "implement-trie", title: "Implement Trie", category: "tries", difficulty: "Medium"},
	{id: "design-add-and-search-words", title: "Design Add and Search Words", category: "tries", difficulty: "Medium", tags: []string{"Design"}},
	{id: "word-search-ii", title: "Word Search II", category: "tries", difficulty: "Hard", tags: []string{"Backtracking"}},
	// Heap / Priority Queue
	{id: "kth-largest-element-in-an-array", title: "Kth Largest Element in an Array", category: "heap-priority-queue", difficulty: "Medium"},
	{id: "last-stone-weight", title: "Last Stone Weight", category: "heap-priority-queue", difficulty: "Easy"},
	{id: "k-closest-points-to-origin", title: "K Closest Points to Origin", category: "heap-priority-queue", difficulty: "Medium"},
	{id: "task-scheduler", title: "Task Scheduler", category: "heap-priority-queue", difficulty: "Medium"},
	{id: "design-twitter", title: "Design Twitter", category: "heap-priority-queue", difficulty: "Medium", tags: []string{"Design"}},
	{id: "find-median-from-data-stream", title: "Find Median from Data Stream", category: "heap-priority-queue", difficulty: "Hard"},
	// Backtracking
	{id: "subsets", title: "Subsets", category: "backtracking", difficulty: "Medium"},
	{id: "combination-sum", title: "Combination Sum", category: "backtracking", difficulty: "Medium"},
	{id: "permutations", title: "Permutations", category: "backtracking", difficulty: "Medium"},
	{id: "subsets-ii", title: "Subsets II", category: "backtracking", difficulty: "Medium"},
	{id: "combination-sum-ii", title: "Combination Sum II", category: "backtracking", difficulty: "Medium"},
	{id: "word-search", title: "Word Search", category: "backtracking", difficulty: "Medium"},
	{id: "palindrome-partitioning", title: "Palindrome Partitioning", category: "backtracking", difficulty: "Medium"},
	{id: "letter-combinations-of-a-phone-number", title: "Letter Combinations of a Phone Number", category: "backtracking", difficulty: "Medium"},
	{id: "n-queens", title: "N-Queens", category: "backtracking", difficulty: "Hard"},
	// Graphs
	{id: "clone-graph", title: "Clone Graph", category: "graphs", difficulty: "Medium"},
	{id: "pacific-atlantic-water-flow", title: "Pacific Atlantic Water Flow", category: "graphs", difficulty: "Medium"},
	{id: "surrounded-regions", title: "Surrounded Regions", category: "graphs", difficulty: "Medium"},
	{id: "rotting-oranges", title: "Rotting Oranges", category: "graphs", difficulty: "Medium"},
	{id: "walls-and-gates", title: "Walls and Gates", category: "graphs", difficulty: "Medium"},
	{id: "course-schedule-ii", title: "Course Schedule II", category: "graphs", difficulty: "Medium"},
	{id: "redundant-connection", title: "Redundant Connection", category: "graphs", difficulty: "Medium"},
	{id: "word-ladder", title: "Word Ladder", category: "graphs", difficulty: "Hard"},
	{id: "graph-valid-tree", title: "Graph Valid Tree", category: "graphs", difficulty: "Medium"},
	{id: "number-of-connected-components", title: "Number of Connected Components", category: "graphs", difficulty: "Medium"},
	{id: "alien-dictionary", title: "Alien Dictionary", category: "graphs", difficulty: "Hard"},
	// Advanced Graphs (mapped to graphs category)
	{id: "reconstruct-itinerary", title: "Reconstruct Itinerary", category: "graphs", difficulty: "Hard"},
	{id: "min-cost-to-connect-all-points", title: "Min Cost to Connect All Points", category: "graphs", difficulty: "Medium"},
	{id: "network-delay-time", title: "Network Delay Time", category: "graphs", difficulty: "Medium"},
	{id: "swim-in-rising-water", title: "Swim in Rising Water", category: "graphs", difficulty: "Hard"},
	{id: "cheapest-flights-within-k-stops", title: "Cheapest Flights Within K Stops", category: "graphs", difficulty: "Medium"},
	// 1D DP
	{id: "house-robber-ii", title: "House Robber II", category: "1d-dp", difficulty: "Medium"},
	{id: "longest-palindromic-substring", title: "Longest Palindromic Substring", category: "1d-dp", difficulty: "Medium"},
	{id: "palindromic-substrings", title: "Palindromic Substrings", category: "1d-dp", difficulty: "Medium"},
	{id: "decode-ways", title: "Decode Ways", category: "1d-dp", difficulty: "Medium"},
	{id: "coin-change", title: "Coin Change", category: "1d-dp", difficulty: "Medium"},
	{id: "maximum-product-subarray", title: "Maximum Product Subarray", category: "1d-dp", difficulty: "Medium"},
	{id: "word-break", title: "Word Break", category: "1d-dp", difficulty: "Medium"},
	{id: "longest-increasing-subsequence", title: "Longest Increasing Subsequence", category: "1d-dp", difficulty: "Medium"},
	{id: "partition-equal-subset-sum", title: "Partition Equal Subset Sum", category: "1d-dp", difficulty: "Medium"},
	// 2D DP
	{id: "unique-paths", title: "Unique Paths", category: "2d-dp", difficulty: "Medium"},
	{id: "longest-common-subsequence", title: "Longest Common Subsequence", category: "2d-dp", difficulty: "Medium"},
	{id: "best-time-to-buy-and-sell-stock-with-cooldown", title: "Best Time to Buy and Sell Stock with Cooldown", category: "2d-dp", difficulty: "Medium"},
	{id: "coin-change-ii", title: "Coin Change II", category: "2d-dp", difficulty: "Medium"},
	{id: "target-sum", title: "Target Sum", category: "2d-dp", difficulty: "Medium"},
	{id: "interleaving-string", title: "Interleaving String", category: "2d-dp", difficulty: "Medium"},
	{id: "longest-increasing-path-in-matrix", title: "Longest Increasing Path in Matrix", category: "2d-dp", difficulty: "Hard"},
	{id: "distinct-subsequences", title: "Distinct Subsequences", category: "2d-dp", difficulty: "Hard"},
	{id: "edit-distance", title: "Edit Distance", category: "2d-dp", difficulty: "Medium"},
	{id: "burst-balloons", title: "Burst Balloons", category: "2d-dp", difficulty: "Hard"},
	{id: "regular-expression-matching", title: "Regular Expression Matching", category: "2d-dp", difficulty: "Hard"},
	// Greedy
	{id: "maximum-subarray", title: "Maximum Subarray", category: "greedy", difficulty: "Easy"},
	{id: "jump-game", title: "Jump Game", category: "greedy", difficulty: "Medium"},
	{id: "jump-game-ii", title: "Jump Game II", category: "greedy", difficulty: "Medium"},
	{id: "gas-station", title: "Gas Station", category: "greedy", difficulty: "Medium"},
	{id: "hand-of-straights", title: "Hand of Straights", category: "greedy", difficulty: "Medium"},
	{id: "merge-triplets-to-form-target", title: "Merge Triplets to Form Target", category: "greedy", difficulty: "Medium"},
	{id: "partition-labels", title: "Partition Labels", category: "greedy", difficulty: "Medium"},
	{id: "valid-parenthesis-string", title: "Valid Parenthesis String", category: "greedy", difficulty: "Medium"},
	// Intervals
	{id: "insert-interval", title: "Insert Interval", category: "intervals", difficulty: "Medium"},
	{id: "merge-intervals", title: "Merge Intervals", category: "intervals", difficulty: "Medium"},
	{id: "non-overlapping-intervals", title: "Non-overlapping Intervals", category: "intervals", difficulty: "Medium"},
	{id: "meeting-rooms", title: "Meeting Rooms", category: "intervals", difficulty: "Easy"},
	{id: "meeting-rooms-ii", title: "Meeting Rooms II", category: "intervals", difficulty: "Medium"},
	// Math & Geometry
	{id: "rotate-image", title: "Rotate Image", category: "math-geometry", difficulty: "Medium"},
	{id: "spiral-matrix", title: "Spiral Matrix", category: "math-geometry", difficulty: "Medium"},
	{id: "set-matrix-zeroes", title: "Set Matrix Zeroes", category: "math-geometry", difficulty: "Medium"},
	{id: "happy-number", title: "Happy Number", category: "math-geometry", difficulty: "Easy"},
	{id: "plus-one", title: "Plus One", category: "math-geometry", difficulty: "Easy"},
	{id: "powx-n", title: "Pow(x, n)", category: "math-geometry", difficulty: "Medium"},
	{id: "multiply-strings", title: "Multiply Strings", category: "math-geometry", difficulty: "Medium"},
	{id: "detect-squares", title: "Detect Squares", category: "math-geometry", difficulty: "Medium"},
	// Bit Manipulation
	{id: "single-number", title: "Single Number", category: "bit-manipulation", difficulty: "Easy"},
	{id: "number-of-1-bits", title: "Number of 1 Bits", category: "bit-manipulation", difficulty: "Easy"},
	{id: "counting-bits", title: "Counting Bits", category: "bit-manipulation", difficulty: "Easy"},
	{id: "reverse-bits", title: "Reverse Bits", category: "bit-manipulation", difficulty: "Easy"},
	{id: "missing-number", title: "Missing Number", category: "bit-manipulation", difficulty: "Easy"},
	{id: "sum-of-two-integers", title: "Sum of Two Integers", category: "bit-manipulation", difficulty: "Medium"},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, buildBacklogProblems(backlogMetas, 100)...)
}

func buildBacklogProblems(metas []backlogMeta, startNumber int) []*Problem {
	problems := make([]*Problem, 0, len(metas))
	for i, meta := range metas {
		number := startNumber + i
		tags := meta.tags
		if len(tags) == 0 {
			tags = []string{meta.category}
		}
		problems = append(problems, &Problem{
			ID:              meta.id,
			Number:          number,
			Title:           meta.title,
			Difficulty:      meta.difficulty,
			Category:        meta.category,
			Tags:            tags,
			RelatedChapters: defaultChaptersForCategory(meta.category),
			Description:     meta.title + " problem statement placeholder.",
			Constraints:     []string{"See canonical statement for constraints."},
			Examples:        []Example{},
			TestCases:       []TestCase{},
			StarterCode:     backlogStarterCode(),
			Hints: []Hint{
				{Level: 1, Type: "approach", Content: "TODO: add detailed hints for " + meta.title},
			},
			Solution: Solution{
				Code:        "",
				Explanation: "Solution placeholder.",
			},
		})
	}
	return problems
}

func defaultChaptersForCategory(category string) []int {
	switch category {
	case "stack":
		return []int{3}
	case "binary-search":
		return []int{1, 4}
	case "linked-list":
		return []int{2}
	case "trees", "tries":
		return []int{11}
	case "heap-priority-queue":
		return []int{7}
	case "backtracking":
		return []int{3, 4}
	case "graphs":
		return []int{6, 7}
	case "1d-dp", "2d-dp":
		return []int{9}
	case "greedy", "intervals":
		return []int{8}
	case "math-geometry", "bit-manipulation":
		return []int{1}
	default:
		return []int{}
	}
}

func backlogStarterCode() string {
	return "def solve(input_data):\n    # TODO: implement\n    pass"
}
