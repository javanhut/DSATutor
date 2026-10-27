package practice

// DP1DProblems contains all 1d-dp category problems.
//
// Every problem is taught with the six-step DP recipe from Chapter 15, and the
// walkthroughs and "plan" hints reuse its subgoal labels:
//
//  1. Define the state
//  2. Write the recurrence
//  3. Set the base cases
//  4. Choose the evaluation order
//  5. Locate the answer
//  6. Optimise space
//
// The main solution always keeps the full table (so the learner can see the
// state); the brute-force recursion and the space-optimised version are
// Alternatives.
var DP1DProblems = []*Problem{
	{
		ID:              "climbing-stairs",
		Number:          70,
		Order:           1,
		Title:           "Climbing Stairs",
		Slug:            "climbing-stairs",
		Difficulty:      "Easy",
		Category:        "1d-dp",
		Tags:            []string{"Dynamic Programming", "Math", "Memoization"},
		RelatedChapters: []int{15, 8},
		Description: `A staircase has **n** steps. You start on the ground (step 0) and want to stand on step n. Each move climbs either **1 step or 2 steps**.

Return the number of different move sequences that take you from the ground to step n exactly. Two sequences are different if they differ in at least one move, so 1+2 and 2+1 count as two ways.`,
		Constraints: []string{
			"1 <= n <= 45",
		},
		Examples: []Example{
			{Input: "n = 2", Output: "2", Explanation: "Two sequences: 1+1, or a single 2."},
			{Input: "n = 3", Output: "3", Explanation: "1+1+1, 1+2 and 2+1."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 2}, Expected: 2, Note: "example 1: two ways"},
			{Input: map[string]interface{}{"n": 3}, Expected: 3, Note: "example 2: order matters, 1+2 and 2+1 both count"},
			{Input: map[string]interface{}{"n": 1}, Expected: 1, Note: "n = 1, the smallest input: a table of size n+1 must not index ways[2]"},
			{Input: map[string]interface{}{"n": 4}, Expected: 5, Hidden: true, Note: "n = 4: 3 + 2 ways"},
			{Input: map[string]interface{}{"n": 5}, Expected: 8, Hidden: true, Note: "n = 5: the Fibonacci pattern continues"},
			{Input: map[string]interface{}{"n": 10}, Expected: 89, Hidden: true, Note: "n = 10: base cases off by one give 55 or 144"},
			{Input: map[string]interface{}{"n": 20}, Expected: 10946, Hidden: true, Note: "n = 20"},
			{Input: map[string]interface{}{"n": 25}, Expected: 121393, Hidden: true, Note: "n = 25: plain recursion already makes about 240,000 calls here"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def climbStairs(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Imagine you are already standing on step n. What was your very last move, and which step must you have been standing on just before it?"},
			{Level: 2, Type: "insight", Content: "Every way to reach step i ends with either a 1-step (from step i-1) or a 2-step (from step i-2), and no sequence ends both ways. So ways(i) = ways(i-1) + ways(i-2), the Fibonacci recurrence. Plain recursion recomputes the same steps billions of times for n = 45; dynamic programming stores each ways(i) once. Chapter 15 builds this exact table.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Define the state: ways[i] = number of move sequences that end on step i.\n2. Write the recurrence: look at the last move; ways[i] = ways[i-1] + ways[i-2].\n3. Set the base cases: ways[0] = 1 (stand still) and ways[1] = 1 (one 1-step).\n4. Choose the evaluation order: i from 2 up to n, so both entries you read are already filled.\n5. Locate the answer: ways[n].\n6. Optimise space: only the last two entries are ever read, so two variables are enough."},
			{Level: 4, Type: "skeleton", Content: "def climbStairs(n):\n    ways = [0] * (n + 1)\n    ways[0] = ____\n    ways[1] = ____\n    for i in range(2, n + 1):\n        ways[i] = ____\n    return ____"},
		},
		Solution: Solution{
			Code: `def climbStairs(n):
    # State: ways[i] = number of move sequences that end on step i
    ways = [0] * (n + 1)
    # Base cases: one way to be on the ground (do nothing), one way to reach step 1
    ways[0] = 1
    ways[1] = 1
    # Evaluation order: low steps first, so ways[i-1] and ways[i-2] are ready
    for i in range(2, n + 1):
        # Recurrence: the last move onto step i was a 1-step or a 2-step
        ways[i] = ways[i - 1] + ways[i - 2]
    # Answer: the entry for the top step
    return ways[n]`,
			Explanation: "This is the first problem of the ladder and the plainest dynamic programming problem there is: we follow Chapter 15's recipe step by step. The state is one number per step (how many sequences end there), the recurrence comes from asking what the last move was, and a table filled from the bottom up computes each entry exactly once.",
			BruteForce:  "Write the recurrence as a recursive function: climb(n) = climb(n-1) + climb(n-2), with climb(0) = climb(1) = 1. It is correct but makes about 2 x ways(n) calls: 15 calls for n = 5, 21,891 for n = 20 and 3,672,623,805 for n = 45. The waste is that the same question is asked again and again: inside climb(45), climb(2) alone is recomputed 701,408,733 times, although there are only 46 different questions (n = 0..45).",
			KeyInsight:  "There are only n+1 different subproblems, ways(0) to ways(n), and each one depends only on the two below it. Compute each one once, smallest first, and store it in a table; every later lookup is O(1).",
			WhyItWorks:  "Subproblem: ways[i] is the number of move sequences from the ground that end exactly on step i. Recurrence: split those sequences by their last move. If it was a 1-step, removing it leaves a sequence ending on step i-1, and every such sequence can be extended; there are ways[i-1] of them. If it was a 2-step, there are ways[i-2]. A sequence has exactly one last move, so the two groups do not overlap and together cover everything: ways[i] = ways[i-1] + ways[i-2]. The base cases ways[0] = 1 (the empty sequence) and ways[1] = 1 start the induction, and filling i in increasing order guarantees both inputs are already final.",
			CommonMistakes: []string{
				"Setting ways[0] = 0 because 'there are zero steps to climb'. Then ways[2] = ways[1] + ways[0] = 1 instead of 2. The empty sequence is one way to stand on the ground; use ways[0] = 1 (or start from ways[1] = 1, ways[2] = 2 and special-case n = 1).",
				"Allocating the table for n+1 entries but writing ways[2] = 2 unconditionally. For n = 1 that is an IndexError; set only ways[0] and ways[1] before the loop, or return early when n is small.",
				"Writing the plain recursion and stopping there. It is correct but exponential; n = 45 makes over 3.6 billion calls. Add a memo dictionary or fill a table.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "One number per step: ways[i] counts the move sequences that finish on step i. The table has n+1 cells because steps run from 0 (the ground) to n (the top).", CodeSnippet: "ways = [0] * (n + 1)", LineStart: 2, LineEnd: 3},
				{Title: "Write the recurrence", Explanation: "Look at the last move. It came from step i-1 (a 1-step) or from step i-2 (a 2-step), and those two groups of sequences never overlap, so we add them.", CodeSnippet: "ways[i] = ways[i - 1] + ways[i - 2]", LineStart: 9, LineEnd: 10},
				{Title: "Set the base cases", Explanation: "The recurrence reads two cells back, so the first two cells must be given directly. There is exactly one way to be on the ground (make no moves) and one way to reach step 1.", CodeSnippet: "ways[0] = 1\nways[1] = 1", LineStart: 4, LineEnd: 6},
				{Title: "Choose the evaluation order", Explanation: "ways[i] reads smaller indices only, so filling i = 2, 3, ..., n guarantees both inputs are final. For n = 5 the table fills as [1, 1, 2, 3, 5, 8].", CodeSnippet: "for i in range(2, n + 1):", LineStart: 7, LineEnd: 8},
				{Title: "Locate the answer", Explanation: "The question asks about step n, which is the last cell of the table.", CodeSnippet: "return ways[n]", LineStart: 11, LineEnd: 12},
				{Title: "Optimise space", Explanation: "Line 10 only ever reads the two previous cells, so the whole table can be replaced by two variables that roll forward. See the 'Rolling variables' alternative: same recurrence, O(1) extra space.", CodeSnippet: "ways[i] = ways[i - 1] + ways[i - 2]", LineStart: 10, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{
					Name: "Plain recursion (brute force)",
					Code: `def climbStairs(n):
    if n <= 1:          # step 0 or step 1: exactly one way
        return 1
    return climbStairs(n - 1) + climbStairs(n - 2)`,
					TimeComplexity:  "O(2^n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The recurrence written directly as code. Correct, and the best place to start, but every call branches in two and the same arguments are recomputed: 21,891 calls for n = 20 and about 3.7 billion for n = 45, which times out. The recursion depth (and so the stack space) is n.",
				},
				{
					Name: "Top-down memoization",
					Code: `def climbStairs(n):
    memo = {0: 1, 1: 1}            # base cases go straight into the memo
    def ways(i):
        if i not in memo:          # compute each state once
            memo[i] = ways(i - 1) + ways(i - 2)
        return memo[i]
    return ways(n)`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The plain recursion plus a dictionary. The first call for each i does the work and stores it; every later call is a lookup. For n = 45 that is 89 calls instead of 3.7 billion. The recursion decides the evaluation order for you, at the cost of a call stack n deep.",
				},
				{
					Name: "Rolling variables (optimise space)",
					Code: `def climbStairs(n):
    two_back, one_back = 1, 1          # ways[0], ways[1]
    for _ in range(2, n + 1):
        two_back, one_back = one_back, one_back + two_back
    return one_back`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "The table version with only the two cells the recurrence reads. After each loop pass, one_back holds ways[i] and two_back holds ways[i-1]. The simultaneous assignment matters: it computes the new sum from the old values before either variable changes.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is ways[0] = 1 and not 0, when there are no steps to climb?", Answer: "ways[i] counts move sequences, and the empty sequence (make no moves) is exactly one sequence that leaves you on step 0. The recurrence needs this: the sequence '2' reaches step 2 by extending the empty sequence, so ways[2] = ways[1] + ways[0] must be 1 + 1 = 2."},
			{Prompt: "What changes if each move may climb 1, 2 or 3 steps?", Answer: "Only the recurrence and the base cases. The last move can now come from i-1, i-2 or i-3, so ways[i] = ways[i-1] + ways[i-2] + ways[i-3], reading 0 for negative indices. The state, order and answer location stay the same, and the rolling version needs three variables instead of two."},
			{Prompt: "Plain recursion for n = 5 makes 15 calls. Which arguments are repeated, and how does a table avoid that?", Answer: "climb(3) is called 2 times, climb(2) 3 times, climb(1) 5 times and climb(0) 3 times, although there are only 6 different arguments (0 to 5). The table computes each ways[i] once, in increasing order, and later reads it in O(1), so the work drops from exponential to n steps."},
		},
	},
	{
		ID:              "house-robber",
		Number:          198,
		Order:           2,
		Title:           "House Robber",
		Slug:            "house-robber",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{15, 8},
		Description: `Houses stand in a row, and **nums[i]** is the amount of money in house i. You may take the money from any set of houses, with one rule: you may **never take from two neighbouring houses** (houses i and i+1), because their alarms are linked.

Return the largest total you can collect.`,
		Constraints: []string{
			"1 <= nums.length <= 100",
			"0 <= nums[i] <= 400",
		},
		Examples: []Example{
			{Input: "nums = [1,2,3,1]", Output: "4", Explanation: "Take houses 0 and 2: 1 + 3 = 4. Taking houses 1 and 3 gives only 2 + 1 = 3."},
			{Input: "nums = [2,7,9,3,1]", Output: "12", Explanation: "Take houses 0, 2 and 4: 2 + 9 + 1 = 12."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 1}}, Expected: 4, Note: "example 1"},
			{Input: map[string]interface{}{"nums": []int{2, 7, 9, 3, 1}}, Expected: 12, Note: "example 2: three non-adjacent houses"},
			{Input: map[string]interface{}{"nums": []int{5}}, Expected: 5, Note: "a single house: take it"},
			{Input: map[string]interface{}{"nums": []int{2, 1}}, Expected: 2, Hidden: true, Note: "two houses: take the larger one"},
			{Input: map[string]interface{}{"nums": []int{2, 1, 1, 2}}, Expected: 4, Hidden: true, Note: "skipping two houses in a row is best (2 + 2); 'every other house' gives only 3"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0}}, Expected: 0, Hidden: true, Note: "all zeros"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 1, 3, 100}}, Expected: 103, Hidden: true, Note: "a big last house changes which earlier houses to take"},
			{Input: map[string]interface{}{"nums": []int{400, 1, 1, 400, 1, 400}}, Expected: 1200, Hidden: true, Note: "maximum values with gaps of different lengths"},
			{Input: map[string]interface{}{"nums": []int{183, 219, 57, 193, 94, 233, 202, 154, 65, 240, 97, 234, 100, 249, 186, 66, 90, 238, 168, 128}}, Expected: 1954, Hidden: true, Note: "20 houses"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def rob(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Look only at the last house. You either take its money or you don't. In each case, which houses are still allowed, and what is the best you could have collected from them?"},
			{Level: 2, Type: "insight", Content: "If you skip the last house, the best total is the best for the houses before it. If you take it, its neighbour is forbidden, so you add its money to the best for the houses two back. best(i) = max(best(i-1), best(i-2) + money of house i). This is Climbing Stairs' 'look at the last move' idea with a max instead of a sum (Chapter 15).", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Define the state: best[i] = most money from the first i houses (nums[0..i-1]).\n2. Write the recurrence: the last of those houses is skipped (best[i-1]) or taken (best[i-2] + nums[i-1]); keep the larger.\n3. Set the base cases: best[0] = 0 (no houses), best[1] = nums[0].\n4. Choose the evaluation order: i from 2 up to n.\n5. Locate the answer: best[n].\n6. Optimise space: only best[i-1] and best[i-2] are read, so two variables suffice."},
			{Level: 4, Type: "skeleton", Content: "def rob(nums):\n    n = len(nums)\n    best = [0] * (n + 1)\n    best[1] = ____\n    for i in range(2, n + 1):\n        skip_it = ____\n        rob_it = ____\n        best[i] = max(skip_it, rob_it)\n    return best[n]"},
		},
		Solution: Solution{
			Code: `def rob(nums):
    n = len(nums)
    # State: best[i] = most money from the first i houses (nums[0..i-1])
    best = [0] * (n + 1)
    # Base cases: no houses -> 0; one house -> take it
    best[0] = 0
    best[1] = nums[0]
    # Evaluation order: left to right, so best[i-1] and best[i-2] are ready
    for i in range(2, n + 1):
        # Recurrence: the last house (nums[i-1]) is either skipped or robbed
        skip_it = best[i - 1]
        rob_it = best[i - 2] + nums[i - 1]
        best[i] = max(skip_it, rob_it)
    # Answer: all n houses considered
    return best[n]`,
			Explanation: "This builds on Climbing Stairs (problem 1): the state is again one number per prefix and the recurrence again looks at the last item. The new idea is optimisation instead of counting: the last house is either taken or skipped, and we keep the better of the two choices with max.",
			BruteForce:  "Try every subset of houses that has no two neighbours and keep the largest sum. Recursively: best_from(i) = max(best_from(i+1), nums[i] + best_from(i+2)). That explores about 1.6^n branches (over 10^20 for n = 100), and best_from(k) for the same k is recomputed along many different paths.",
			KeyInsight:  "The best total for the first i houses does not depend on how you got there, only on i. So there are just n+1 different subproblems, and each is a max over two smaller ones.",
			WhyItWorks:  "Subproblem: best[i] is the most money collectable from houses 0..i-1 without taking two neighbours. Recurrence: an optimal plan for those houses either skips house i-1, in which case it is an optimal plan for the first i-1 houses (best[i-1]), or takes it, in which case house i-2 is forbidden and the rest must be an optimal plan for the first i-2 houses (best[i-2] + nums[i-1]). If the rest were not optimal, swapping in a better plan would improve the total, so taking the max of the two cases is complete and correct (optimal substructure). Base cases best[0] = 0 and best[1] = nums[0] are direct, and the left-to-right loop fills each entry after the two it reads.",
			CommonMistakes: []string{
				"Taking every other house: max(sum of even positions, sum of odd positions). For [2,1,1,2] that gives 3, but taking the first and last house gives 4. Sometimes the best plan skips two houses in a row; only the DP considers that.",
				"Using best[i-1] + nums[i-1] in the 'rob it' case. Taking house i-1 forbids its neighbour, so you must build on best[i-2], two houses back.",
				"Writing best[1] = nums[0] before checking the length, or reading nums[1] unconditionally. With one house the table has two cells and the loop never runs; make sure your base cases work for n = 1.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "best[i] is the best total using only the first i houses. Indexing by 'how many houses' rather than 'which house' gives a natural empty case, best[0].", CodeSnippet: "best = [0] * (n + 1)", LineStart: 3, LineEnd: 4},
				{Title: "Write the recurrence", Explanation: "Decide about the last house in the prefix, nums[i-1]. Skip it and keep best[i-1]; or take it, which forbids its neighbour, so add it to best[i-2]. Keep the larger.", CodeSnippet: "skip_it = best[i - 1]\nrob_it = best[i - 2] + nums[i - 1]\nbest[i] = max(skip_it, rob_it)", LineStart: 10, LineEnd: 13},
				{Title: "Set the base cases", Explanation: "With no houses you collect 0; with one house you take it. The loop starts at i = 2, so these two cells must be set first.", CodeSnippet: "best[0] = 0\nbest[1] = nums[0]", LineStart: 5, LineEnd: 7},
				{Title: "Choose the evaluation order", Explanation: "Each entry reads the two before it, so fill left to right. For [2,7,9,3,1] the table becomes [0, 2, 7, 11, 11, 12].", CodeSnippet: "for i in range(2, n + 1):", LineStart: 8, LineEnd: 9},
				{Title: "Locate the answer", Explanation: "The question is about all the houses, which is the prefix of length n.", CodeSnippet: "return best[n]", LineStart: 14, LineEnd: 15},
				{Title: "Optimise space", Explanation: "Only best[i-1] and best[i-2] are read, so two rolling variables replace the table; see the 'Rolling variables' alternative.", CodeSnippet: "best[i] = max(skip_it, rob_it)", LineStart: 13, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Plain recursion (brute force)",
					Code: `def rob(nums):
    def best_from(i):
        # most money from houses i, i+1, ..., end
        if i >= len(nums):
            return 0
        skip = best_from(i + 1)
        take = nums[i] + best_from(i + 2)
        return max(skip, take)
    return best_from(0)`,
					TimeComplexity:  "O(2^n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The same take-or-skip decision written as a recursion over suffixes. Correct, but best_from(k) is recomputed on many paths, so the number of calls grows like the Fibonacci numbers. Adding a memo dictionary keyed by i makes it O(n).",
				},
				{
					Name: "Rolling variables (optimise space)",
					Code: `def rob(nums):
    two_back, one_back = 0, 0        # best[i-2], best[i-1]
    for money in nums:
        two_back, one_back = one_back, max(one_back, two_back + money)
    return one_back`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "The table version keeping only the two cells the recurrence reads. Starting both at 0 (best[-1] and best[0]) makes the first house come out as max(0, 0 + nums[0]), so no special case is needed for n = 1.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does 'take every other house' fail, and how does the DP avoid that mistake?", Answer: "Alternating forces a gap of exactly one house, but an optimal plan may leave a gap of two (for [2,1,1,2] take houses 0 and 3 for 4). The DP never fixes the gap length: at each house it compares skipping (best[i-1]) with taking (best[i-2] + money), so a skip can follow a skip whenever that pays more."},
			{Prompt: "What if the rule were 'never take from two houses within distance 2' (so i and i+2 are also forbidden)?", Answer: "Only the recurrence changes: taking house i-1 now forbids i-2 and i-3, so best[i] = max(best[i-1], best[i-3] + nums[i-1]), with best of a negative index treated as 0. You would also need one more base case, and three rolling variables instead of two."},
			{Prompt: "best[i] tells you the best total, but not which houses. How could you recover the chosen houses?", Answer: "Walk backwards from i = n: if best[i] == best[i-1], house i-1 was skipped, so go to i-1; otherwise it was taken, record it and jump to i-2. This works because the table keeps every intermediate answer, which the rolling-variable version throws away."},
		},
	},
	{
		ID:              "house-robber-ii",
		Number:          213,
		Order:           3,
		Title:           "House Robber II",
		Slug:            "house-robber-ii",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{15},
		Description: `This is House Robber with the houses arranged in a **circle**: nums[i] is the money in house i, and house 0 and house n-1 are neighbours as well as every pair i, i+1. You may never take from two neighbouring houses.

Return the largest total you can collect. With a single house there are no neighbours, so you simply take it.`,
		Constraints: []string{
			"1 <= nums.length <= 100",
			"0 <= nums[i] <= 1000",
		},
		Examples: []Example{
			{Input: "nums = [2,3,2]", Output: "3", Explanation: "Houses 0 and 2 are neighbours on the circle, so you cannot take both 2s. The best is house 1 alone."},
			{Input: "nums = [1,2,3,1]", Output: "4", Explanation: "Take houses 0 and 2: 1 + 3 = 4."},
			{Input: "nums = [1,2,3]", Output: "3"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 3, 2}}, Expected: 3, Note: "example 1: first and last house are neighbours"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 1}}, Expected: 4, Note: "example 2"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}}, Expected: 3, Note: "example 3"},
			{Input: map[string]interface{}{"nums": []int{5}}, Expected: 5, Hidden: true, Note: "a single house: both slices nums[1:] and nums[:-1] are empty, so it needs its own case"},
			{Input: map[string]interface{}{"nums": []int{1, 2}}, Expected: 2, Hidden: true, Note: "two houses are neighbours both ways round"},
			{Input: map[string]interface{}{"nums": []int{2, 1, 1, 2}}, Expected: 3, Hidden: true, Note: "the straight-line answer (4) takes both ends, which the circle forbids"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 1, 3, 100}}, Expected: 103, Hidden: true, Note: "the big last house rules out house 0"},
			{Input: map[string]interface{}{"nums": []int{200, 3, 140, 20, 10}}, Expected: 340, Hidden: true, Note: "the big first house rules out the last one"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0, 0}}, Expected: 0, Hidden: true, Note: "all zeros"},
			{Input: map[string]interface{}{"nums": []int{94, 40, 49, 65, 21, 21, 106, 80, 92, 81, 679, 4, 61, 6, 237, 12}}, Expected: 1339, Hidden: true, Note: "16 houses"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def rob(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The circle only adds one new conflict: house 0 with house n-1. In any valid plan, at least one of those two houses is left alone. What problem do you get if you decide in advance which one it is?"},
			{Level: 2, Type: "insight", Content: "Split into two straight-line problems. Either house n-1 is not taken (solve House Robber on nums[0..n-2]) or house 0 is not taken (solve it on nums[1..n-1]). Every valid circular plan is a valid plan in one of these lines, so the answer is the larger of the two. Reuse your House Robber DP (Chapter 15) as a helper.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Handle n = 1 first: return nums[0].\n2. Define the state (for one line of houses): best[i] = most money from the first i houses of that line.\n3. Write the recurrence: best[i] = max(best[i-1], best[i-2] + houses[i-1]).\n4. Set the base cases: best[0] = 0, best[1] = houses[0].\n5. Choose the evaluation order: left to right.\n6. Locate the answer: run the line DP on nums without its last house and on nums without its first house; return the larger result.\n7. Optimise space: the helper can use two rolling variables."},
			{Level: 4, Type: "skeleton", Content: "def rob(nums):\n    if len(nums) == 1:\n        return ____\n\n    def rob_line(houses):\n        best = [0] * (len(houses) + 1)\n        best[1] = houses[0]\n        for i in range(2, len(houses) + 1):\n            best[i] = max(best[i - 1], best[i - 2] + houses[i - 1])\n        return best[len(houses)]\n\n    return max(rob_line(____), rob_line(____))"},
		},
		Solution: Solution{
			Code: `def rob(nums):
    if len(nums) == 1:
        return nums[0]

    def rob_line(houses):
        # State: best[i] = most money from the first i houses of this line
        best = [0] * (len(houses) + 1)
        best[1] = houses[0]
        for i in range(2, len(houses) + 1):
            best[i] = max(best[i - 1], best[i - 2] + houses[i - 1])
        return best[len(houses)]

    # The first and last house touch, so at least one of them is skipped:
    skip_last = rob_line(nums[:-1])    # house 0 allowed, last house never
    skip_first = rob_line(nums[1:])    # last house allowed, house 0 never
    return max(skip_last, skip_first)`,
			Explanation: "This builds directly on House Robber (problem 2): the same DP runs unchanged as a helper. The new idea is handling a constraint that breaks the straight line, the wrap-around between the first and last house, by splitting into cases that are each a straight line again.",
			BruteForce:  "Enumerate every subset of houses, reject those with two neighbours (including house 0 with house n-1), and keep the best sum: 2^n subsets, about 10^30 for n = 100. Adding a 'did I take house 0?' flag to the House Robber recursion also works, but the exponential recursion still repeats states unless memoized.",
			KeyInsight:  "A circle has no natural place to start a DP, but any valid plan leaves house 0 or house n-1 (or both) untaken. Removing one of them cuts the circle into a line, and the line is the problem you already solved.",
			WhyItWorks:  "Case split: every valid circular plan skips house n-1 or skips house 0, because taking both would take two neighbours. A plan that skips house n-1 is exactly a valid straight-line plan for nums[0..n-2] (the only extra circular conflict involved house n-1), and vice versa, so rob_line(nums[:-1]) is the best plan in that case; symmetrically for nums[1:]. Plans that skip both are counted in both cases, which is harmless for a maximum. Inside each line, the House Robber recurrence best[i] = max(best[i-1], best[i-2] + houses[i-1]) is correct by the same last-house argument as before.",
			CommonMistakes: []string{
				"Forgetting n = 1. Then nums[:-1] and nums[1:] are both empty and houses[0] raises an IndexError (or the answer comes out 0 instead of nums[0]). Return nums[0] first.",
				"Running House Robber once on the whole array. For [2,1,1,2] it takes both 2s for 4, but on a circle they are neighbours; the answer is 3.",
				"Splitting into 'take house 0' and 'take house n-1' instead of 'skip one of them'. Forcing a take is a different, harder constraint; skipping a house is what turns the circle into a line.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle the single house", Explanation: "With one house there is nothing to conflict with. The two slices below would both be empty, so answer this case directly.", CodeSnippet: "if len(nums) == 1:\n    return nums[0]", LineStart: 2, LineEnd: 3},
				{Title: "Define the state", Explanation: "The helper is House Robber on a straight line: best[i] is the most money from the first i houses of that line.", CodeSnippet: "best = [0] * (len(houses) + 1)", LineStart: 5, LineEnd: 7},
				{Title: "Write the recurrence", Explanation: "Skip the last house of the prefix (best[i-1]) or take it on top of best[i-2]; keep the larger. The base case best[1] = houses[0] is set just before the loop, and the loop runs left to right.", CodeSnippet: "best[1] = houses[0]\nbest[i] = max(best[i - 1], best[i - 2] + houses[i - 1])", LineStart: 8, LineEnd: 10},
				{Title: "Locate the answer", Explanation: "Run the line DP twice: once without the last house, once without the first. For [2,1,1,2] that is rob_line([2,1,1]) = 3 and rob_line([1,1,2]) = 3, so the answer is 3, not the straight-line 4.", CodeSnippet: "skip_last = rob_line(nums[:-1])\nskip_first = rob_line(nums[1:])\nreturn max(skip_last, skip_first)", LineStart: 13, LineEnd: 16},
				{Title: "Optimise space", Explanation: "The helper only reads best[i-1] and best[i-2], so it can roll two variables, as in the alternative. Passing start and end indices instead of slicing also avoids copying the list.", CodeSnippet: "best[i] = max(best[i - 1], best[i - 2] + houses[i - 1])", LineStart: 10, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{
					Name: "Recursion with a 'took house 0' flag (brute force)",
					Code: `def rob(nums):
    n = len(nums)
    def best_from(i, took_first):
        # most money from houses i..n-1, knowing whether house 0 was taken
        if i >= n:
            return 0
        skip = best_from(i + 1, took_first)
        if i == n - 1 and took_first and n > 1:
            return skip                      # the last house touches house 0
        take = nums[i] + best_from(i + 2, took_first or i == 0)
        return max(skip, take)
    return best_from(0, False)`,
					TimeComplexity:  "O(2^n)",
					SpaceComplexity: "O(n)",
					Explanation:     "A different way to handle the circle: carry the one extra fact the last house needs (was house 0 taken?) as part of the state. Without a memo it is exponential; memoizing on (i, took_first) gives 2n states and O(n) time. Adding a flag to the state is a general DP trick you will meet again.",
				},
				{
					Name: "Rolling variables (optimise space)",
					Code: `def rob(nums):
    if len(nums) == 1:
        return nums[0]

    def rob_line(houses):
        two_back, one_back = 0, 0
        for money in houses:
            two_back, one_back = one_back, max(one_back, two_back + money)
        return one_back

    return max(rob_line(nums[:-1]), rob_line(nums[1:]))`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The same case split with the O(1)-space House Robber helper. The slices still copy up to n-1 elements; pass index ranges instead to make the extra space O(1).",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "A plan that skips both house 0 and house n-1 appears in both cases. Why doesn't that double counting cause a wrong answer?", Answer: "We take the maximum of the two cases, not the sum. A plan appearing in both just means both cases can reach its total; the max is still the best total over all valid circular plans. Double counting would only matter if we were counting plans, as in Climbing Stairs."},
			{Prompt: "Why does skipping one end turn the circle into an ordinary line?", Answer: "The circle differs from a line only by the single extra edge between house 0 and house n-1. If house n-1 is never taken, that edge can never be violated, so the remaining constraints are exactly those of the line nums[0..n-2]. The same holds for house 0 and nums[1..n-1]."},
		},
	},
	{
		ID:              "decode-ways",
		Number:          91,
		Order:           4,
		Title:           "Decode Ways",
		Slug:            "decode-ways",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{15},
		Description: `A message of capital letters was encoded with A -> "1", B -> "2", ..., Z -> "26", and the codes were written one after another with no separators. Given the digit string **s**, return how many different letter messages could have produced it.

A code is valid only if it is "1" to "9" or "10" to "26". In particular a code may not start with "0": "06" is not the same as "6" and is invalid, and a lone "0" decodes to nothing. If s cannot be decoded at all, return 0. The answer fits in a 32-bit integer.`,
		Constraints: []string{
			"1 <= s.length <= 100",
			"s contains only digits and may contain leading zeros",
		},
		Examples: []Example{
			{Input: `s = "12"`, Output: "2", Explanation: `"AB" (1 2) or "L" (12).`},
			{Input: `s = "226"`, Output: "3", Explanation: `"BZ" (2 26), "VF" (22 6) or "BBF" (2 2 6).`},
			{Input: `s = "06"`, Output: "0", Explanation: `"06" is not a valid code and "0" alone is not a letter, so nothing decodes.`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "12"}, Expected: 2, Note: "example 1"},
			{Input: map[string]interface{}{"s": "226"}, Expected: 3, Note: "example 2"},
			{Input: map[string]interface{}{"s": "06"}, Expected: 0, Note: "example 3: leading zero"},
			{Input: map[string]interface{}{"s": "1"}, Expected: 1, Hidden: true, Note: "a single digit"},
			{Input: map[string]interface{}{"s": "0"}, Expected: 0, Hidden: true, Note: "a single zero decodes to nothing"},
			{Input: map[string]interface{}{"s": "10"}, Expected: 1, Hidden: true, Note: "a zero is only valid as the second digit of 10 or 20"},
			{Input: map[string]interface{}{"s": "100"}, Expected: 0, Hidden: true, Note: "the second zero has no partner: '00' is not a code"},
			{Input: map[string]interface{}{"s": "2101"}, Expected: 1, Hidden: true, Note: "the zero forces 2|10|1"},
			{Input: map[string]interface{}{"s": "27"}, Expected: 1, Hidden: true, Note: "27 is above 26, so only 2|7"},
			{Input: map[string]interface{}{"s": "230"}, Expected: 0, Hidden: true, Note: "'30' is not a code and '0' alone is not a letter"},
			{Input: map[string]interface{}{"s": "11106"}, Expected: 2, Hidden: true, Note: "1|1|10|6 and 11|10|6"},
			{Input: map[string]interface{}{"s": "2611055971756562"}, Expected: 4, Hidden: true, Note: "longer mixed string"},
			{Input: map[string]interface{}{"s": "111111111111111111111111"}, Expected: 75025, Hidden: true, Note: "24 ones: exactly the Climbing Stairs (Fibonacci) count"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def numDecodings(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Think about the last letter of a decoded message. How many digits at the end of s could it have used, and when is each choice allowed?"},
			{Level: 2, Type: "insight", Content: "The last letter used either the last 1 digit or the last 2 digits, just like the last move in Climbing Stairs was 1 or 2 steps. The difference is that each move is only allowed when its digits form a valid code: one digit must be 1-9, two digits must be 10-26. ways(i) = (ways(i-1) if one-digit code is valid) + (ways(i-2) if two-digit code is valid). See Chapter 15.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Define the state: ways[i] = number of decodings of the prefix s[:i].\n2. Write the recurrence: add ways[i-1] if s[i-1] is not '0'; add ways[i-2] if s[i-2:i] is between 10 and 26.\n3. Set the base cases: ways[0] = 1 (the empty prefix has one decoding: nothing).\n4. Choose the evaluation order: i from 1 to n.\n5. Locate the answer: ways[n].\n6. Optimise space: only the two previous entries are read."},
			{Level: 4, Type: "skeleton", Content: "def numDecodings(s):\n    n = len(s)\n    ways = [0] * (n + 1)\n    ways[0] = ____\n    for i in range(1, n + 1):\n        if ____:\n            ways[i] += ways[i - 1]\n        if i >= 2 and ____:\n            ways[i] += ways[i - 2]\n    return ways[n]"},
		},
		Solution: Solution{
			Code: `def numDecodings(s):
    n = len(s)
    # State: ways[i] = number of ways to decode the prefix s[:i]
    ways = [0] * (n + 1)
    # Base case: the empty prefix has one decoding (decode nothing)
    ways[0] = 1
    # Evaluation order: shorter prefixes first
    for i in range(1, n + 1):
        # Recurrence: the last letter used either 1 digit or 2 digits
        if s[i - 1] != "0":                          # one digit, 1-9
            ways[i] += ways[i - 1]
        if i >= 2 and 10 <= int(s[i - 2:i]) <= 26:   # two digits, 10-26
            ways[i] += ways[i - 2]
    # Answer: the whole string
    return ways[n]`,
			Explanation: "A natural follow-on to Climbing Stairs (problem 1): there the last move was 1 or 2 steps, here the last letter uses 1 or 2 digits. The one new idea is that a transition is only allowed when a condition holds, so each term of the sum is guarded by a validity check, and the count can drop to 0.",
			BruteForce:  "Recursively try to cut a 1-digit code or a 2-digit code off the front and count the complete decodings. For a string of 24 ones that makes 196,417 calls to find 75,025 decodings, and the count of 100 ones is about 5.7 x 10^20, so listing or recursing over every decoding cannot finish.",
			KeyInsight:  "How many ways the prefix s[:i] can be decoded does not depend on how the rest of the string is decoded. So the answer for each prefix can be computed once from the answers for the two shorter prefixes.",
			WhyItWorks:  "Subproblem: ways[i] counts decodings of s[:i]. Recurrence: classify each decoding by its last code. If the last code is one digit, it must be s[i-1] in 1-9, and the rest is any decoding of s[:i-1] (ways[i-1] options). If it is two digits, s[i-2:i] must be 10-26 (which excludes a leading zero, since '05' is 5), and the rest is a decoding of s[:i-2]. Every decoding has exactly one last code, so the cases are disjoint and complete; when a case is invalid it contributes 0. Base case: ways[0] = 1 so that a valid first code (or a valid first pair) counts as one decoding. Filling i upward makes ways[i-1] and ways[i-2] final before use.",
			CommonMistakes: []string{
				"Checking the two-digit code with int(s[i-2:i]) <= 26 only. '06' converts to 6, which passes, but a code cannot start with 0. Require 10 <= value <= 26 (or s[i-2] != '0').",
				"Returning 1 for an empty or zero-only prefix, or treating '0' as a letter. A zero is only valid as the second digit of 10 or 20; '0', '100' and '230' all decode to 0 ways.",
				"Setting ways[0] = 0 because the empty string 'has no letters'. Then ways[1] = 0 for every input and everything stays 0. The empty prefix has exactly one decoding: the empty message.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "ways[i] counts the decodings of the first i characters. Using prefix lengths 0..n gives an empty prefix to start from.", CodeSnippet: "ways = [0] * (n + 1)", LineStart: 3, LineEnd: 4},
				{Title: "Write the recurrence", Explanation: "Look at the last code. A one-digit code is fine unless the digit is 0; a two-digit code must be 10-26. Each valid choice adds the number of ways to decode what is left in front of it.", CodeSnippet: "if s[i - 1] != \"0\":\n    ways[i] += ways[i - 1]\nif i >= 2 and 10 <= int(s[i - 2:i]) <= 26:\n    ways[i] += ways[i - 2]", LineStart: 9, LineEnd: 13},
				{Title: "Set the base cases", Explanation: "ways[0] = 1: the empty prefix has exactly one decoding. Without it the first valid code would add nothing.", CodeSnippet: "ways[0] = 1", LineStart: 5, LineEnd: 6},
				{Title: "Choose the evaluation order", Explanation: "Each entry reads the two before it, so fill i = 1..n. For '226' the table is [1, 1, 2, 3]: '2' -> 1, '22' -> 2 (2|2, 22), '226' -> 2 + 1 = 3.", CodeSnippet: "for i in range(1, n + 1):", LineStart: 7, LineEnd: 8},
				{Title: "Locate the answer", Explanation: "The whole string is the prefix of length n.", CodeSnippet: "return ways[n]", LineStart: 14, LineEnd: 15},
				{Title: "Optimise space", Explanation: "Only ways[i-1] and ways[i-2] are read, so two variables are enough; see the alternative.", CodeSnippet: "ways[i] += ways[i - 2]", LineStart: 13, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Plain recursion (brute force)",
					Code: `def numDecodings(s):
    def count_from(i):
        # number of decodings of the suffix s[i:]
        if i == len(s):
            return 1                      # decoded everything
        if s[i] == "0":
            return 0                      # no code starts with 0
        total = count_from(i + 1)         # one-digit code
        if i + 1 < len(s) and int(s[i:i + 2]) <= 26:
            total += count_from(i + 2)    # two-digit code
        return total
    return count_from(0)`,
					TimeComplexity:  "O(2^n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Cuts codes off the front instead of the back; checking s[i] != '0' first covers the leading-zero rule for both code lengths. Correct, but count_from(k) is recomputed along many paths: for '1' * 24 it makes 196,417 calls. A memo keyed by i makes it O(n).",
				},
				{
					Name: "Rolling variables (optimise space)",
					Code: `def numDecodings(s):
    two_back, one_back = 0, 1         # ways[i-2], ways[i-1], starting at ways[0] = 1
    for i in range(1, len(s) + 1):
        current = 0
        if s[i - 1] != "0":
            current += one_back
        if i >= 2 and 10 <= int(s[i - 2:i]) <= 26:
            current += two_back
        two_back, one_back = one_back, current
    return one_back`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "The table version with only the last two entries kept. current must start at 0 on every pass, because a step with no valid code contributes nothing.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "For a string of only '1's and '2's (no zeros, no pairs above 26), why is the answer a Fibonacci number?", Answer: "Then every one-digit code and every two-digit code is valid, so both guards always pass and the recurrence becomes ways[i] = ways[i-1] + ways[i-2] with ways[0] = ways[1] = 1, exactly Climbing Stairs. A string of 24 ones therefore has 75,025 decodings, the same as climbing 24 stairs."},
			{Prompt: "Why must a zero always be paired with the digit before it, and what does the DP do when that is impossible?", Answer: "No letter has code '0', so a zero can only be the second digit of '10' or '20'. At such a position the one-digit term is skipped; if the pair before it is not 10 or 20 the two-digit term is skipped too, ways[i] becomes 0, and every later entry built on it inherits that 0, so the final answer is 0."},
		},
	},
	{
		ID:              "coin-change",
		Number:          322,
		Order:           5,
		Title:           "Coin Change",
		Slug:            "coin-change",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Dynamic Programming", "Breadth-First Search"},
		RelatedChapters: []int{15, 17},
		Description: `You are given coin values **coins** (each value can be used as many times as you like) and a target **amount**. Return the smallest number of coins whose values add up to exactly amount.

If no combination of coins adds up to amount, return -1. An amount of 0 needs 0 coins.`,
		Constraints: []string{
			"1 <= coins.length <= 12",
			"1 <= coins[i] <= 2^31 - 1",
			"0 <= amount <= 10^4",
		},
		Examples: []Example{
			{Input: "coins = [1,2,5], amount = 11", Output: "3", Explanation: "5 + 5 + 1 = 11 uses 3 coins; no 2-coin combination reaches 11."},
			{Input: "coins = [2], amount = 3", Output: "-1", Explanation: "Only even totals can be made from 2s."},
			{Input: "coins = [1], amount = 0", Output: "0"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"coins": []int{1, 2, 5}, "amount": 11}, Expected: 3, Note: "example 1"},
			{Input: map[string]interface{}{"coins": []int{2}, "amount": 3}, Expected: -1, Note: "example 2: impossible amount"},
			{Input: map[string]interface{}{"coins": []int{1}, "amount": 0}, Expected: 0, Note: "example 3: amount = 0 needs no coins"},
			{Input: map[string]interface{}{"coins": []int{1, 3, 4}, "amount": 6}, Expected: 2, Hidden: true, Note: "greedy takes 4+1+1 (3 coins); the best is 3+3"},
			{Input: map[string]interface{}{"coins": []int{5, 10}, "amount": 3}, Expected: -1, Hidden: true, Note: "every coin is larger than the amount"},
			{Input: map[string]interface{}{"coins": []int{1}, "amount": 7}, Expected: 7, Hidden: true, Note: "a single coin type of value 1"},
			{Input: map[string]interface{}{"coins": []int{7}, "amount": 7}, Expected: 1, Hidden: true, Note: "one coin exactly matches"},
			{Input: map[string]interface{}{"coins": []int{2, 4}, "amount": 7}, Expected: -1, Hidden: true, Note: "unreachable even though smaller amounts are reachable: INF + 1 must stay INF"},
			{Input: map[string]interface{}{"coins": []int{3, 7}, "amount": 13}, Expected: 3, Hidden: true, Note: "3 + 3 + 7, with unreachable amounts (1, 2, 4, 5, ...) along the way"},
			{Input: map[string]interface{}{"coins": []int{2, 5, 10, 1}, "amount": 27}, Expected: 4, Hidden: true, Note: "unsorted coins: 10 + 10 + 5 + 2"},
		},
		TimeComplexity:  "O(amount * k)",
		SpaceComplexity: "O(amount)",
		StarterCode:     "def coinChange(coins, amount):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose someone hands you an optimal pile of coins for the amount. Remove one coin from it. What can you say about the pile that is left?"},
			{Level: 2, Type: "insight", Content: "The remaining pile must be an optimal pile for amount - coin, or you could swap in a smaller one. You don't know which coin was last, so try every coin: fewest(a) = 1 + min over coins c <= a of fewest(a - c). Greedy (always take the biggest coin) fails: for coins [1,3,4] and amount 6 it gives 4+1+1, but 3+3 is better. Chapter 15 walks through this table.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Define the state: fewest[a] = fewest coins that make exactly a, or infinity if impossible.\n2. Write the recurrence: the last coin is some c <= a; fewest[a] = min(fewest[a - c] + 1).\n3. Set the base cases: fewest[0] = 0; every other entry starts at infinity.\n4. Choose the evaluation order: a from 1 up to amount, so every a - c is already done.\n5. Locate the answer: fewest[amount], turned into -1 if it is still infinity.\n6. Optimise space: fewest[a] can read as far back as the largest coin, so keep the 1-D table."},
			{Level: 4, Type: "skeleton", Content: "def coinChange(coins, amount):\n    INF = float(\"inf\")\n    fewest = [INF] * (amount + 1)\n    fewest[0] = ____\n    for a in range(1, amount + 1):\n        for coin in coins:\n            if coin <= a:\n                fewest[a] = min(fewest[a], ____)\n    return ____"},
		},
		Solution: Solution{
			Code: `def coinChange(coins, amount):
    INF = float("inf")
    # State: fewest[a] = fewest coins that add up to exactly a (INF = impossible)
    fewest = [INF] * (amount + 1)
    # Base case: zero coins make amount 0
    fewest[0] = 0
    # Evaluation order: smaller amounts first
    for a in range(1, amount + 1):
        # Recurrence: try every coin as the last coin used
        for coin in coins:
            if coin <= a and fewest[a - coin] + 1 < fewest[a]:
                fewest[a] = fewest[a - coin] + 1
    # Answer: fewest[amount], or -1 if it was never reached
    return fewest[amount] if fewest[amount] != INF else -1`,
			Explanation: "Climbing Stairs (problem 1) looked back a fixed 1 or 2 steps; here the 'steps' are the coin values, there can be many of them, and we minimise instead of count. The new ideas are a recurrence with a min over a variable list of choices, and marking unreachable states with infinity so they never win a min.",
			BruteForce:  "Recursively try every coin as the last coin and recurse on amount - coin, keeping the smallest result. The recursion tree has up to k branches per level and depth up to amount (with a 1-coin), so it is O(k^amount): already about 0.2 seconds of Python for coins [1,2,5,10] and amount 27, and hopeless for amount 10,000. The same remaining amount is solved over and over along different paths. Greedy (biggest coin first) is fast but wrong: [1,3,4], 6 gives 3 coins instead of 2.",
			KeyInsight:  "The fewest coins for an amount depends only on that amount, not on which coins were used to get there. With amount + 1 possible states and k choices for the last coin, a table filled from 0 upward costs O(amount * k).",
			WhyItWorks:  "Subproblem: fewest[a] is the minimum number of coins summing to exactly a. Recurrence: any non-empty optimal pile for a has some last coin c; the rest is a pile for a - c, and it must be optimal for a - c, otherwise replacing it would make a smaller pile for a (optimal substructure). Trying every coin c <= a therefore includes the optimal choice, so fewest[a] = min over c of fewest[a - c] + 1 is complete. Base case fewest[0] = 0. Unreachable amounts stay at infinity, and infinity + 1 is still infinity, so they never produce a fake answer. Filling a in increasing order means every a - c (which is smaller) is final when read.",
			CommonMistakes: []string{
				"Using a greedy 'largest coin first' loop. It is optimal for some coin systems (like 1, 5, 10, 25) but not in general: [1,3,4] with amount 6 gives 4+1+1 instead of 3+3.",
				"Initialising the table with 0 or -1 instead of infinity. A 0 looks like a great answer to min, and -1 breaks the + 1. Use infinity (or amount + 1, which is larger than any real answer) and convert to -1 only at the end.",
				"Forgetting amount = 0: the answer is 0 coins, not -1. With fewest[0] = 0 and the loop starting at 1 this falls out naturally.",
			},
			TimeComplexity:  "O(amount * k)",
			SpaceComplexity: "O(amount)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "fewest[a] is the answer to the same question for a smaller amount a. Infinity means 'cannot be made (yet)'.", CodeSnippet: "fewest = [INF] * (amount + 1)", LineStart: 2, LineEnd: 4},
				{Title: "Write the recurrence", Explanation: "The last coin in an optimal pile is one of the coin values. Try each one that fits: what is left needs fewest[a - coin] coins, plus this one. Keep the smallest.", CodeSnippet: "for coin in coins:\n    if coin <= a and fewest[a - coin] + 1 < fewest[a]:\n        fewest[a] = fewest[a - coin] + 1", LineStart: 9, LineEnd: 12},
				{Title: "Set the base cases", Explanation: "Amount 0 needs no coins. Every other cell starts at infinity, so an amount nobody can reach stays infinite.", CodeSnippet: "fewest[0] = 0", LineStart: 5, LineEnd: 6},
				{Title: "Choose the evaluation order", Explanation: "a - coin is always smaller than a, so filling a = 1, 2, ..., amount means every cell we read is final. For coins [1,3,4] the table up to 6 is [0, 1, 2, 1, 1, 2, 2].", CodeSnippet: "for a in range(1, amount + 1):", LineStart: 7, LineEnd: 8},
				{Title: "Locate the answer", Explanation: "The answer is the cell for the full amount; if it is still infinity, no combination exists and the problem asks for -1.", CodeSnippet: "return fewest[amount] if fewest[amount] != INF else -1", LineStart: 13, LineEnd: 14},
				{Title: "Optimise space", Explanation: "Unlike Climbing Stairs, the recurrence reads up to max(coins) cells back, and coin values are not small, so the full 1-D table of amount + 1 cells is the practical choice. Time O(amount * k) for k coins.", CodeSnippet: "fewest[a] = fewest[a - coin] + 1", LineStart: 12, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Plain recursion (brute force)",
					Code: `def coinChange(coins, amount):
    def fewest(a):
        # fewest coins for amount a, or -1 if impossible
        if a == 0:
            return 0
        best = -1
        for coin in coins:
            if coin <= a:
                rest = fewest(a - coin)
                if rest != -1 and (best == -1 or rest + 1 < best):
                    best = rest + 1
        return best
    return fewest(amount)`,
					TimeComplexity:  "O(k^amount)",
					SpaceComplexity: "O(amount)",
					Explanation:     "The recurrence as a recursive function, using -1 for 'impossible'. It re-solves the same remaining amounts many times, so it only finishes for small amounts. Add a dictionary memo keyed by a to turn it into the top-down version of the table solution.",
				},
				{
					Name: "Top-down memoization",
					Code: `def coinChange(coins, amount):
    memo = {0: 0}
    def fewest(a):
        if a in memo:
            return memo[a]
        best = float("inf")
        for coin in coins:
            if coin <= a:
                best = min(best, fewest(a - coin) + 1)
        memo[a] = best
        return best
    answer = fewest(amount)
    return answer if answer != float("inf") else -1`,
					TimeComplexity:  "O(amount * k)",
					SpaceComplexity: "O(amount)",
					Explanation:     "Same states and recurrence as the table, computed on demand. It only visits amounts reachable by subtracting coins from the target, but the recursion can go amount levels deep (with a 1-coin), so for amount = 10,000 the bottom-up table is safer than relying on a large recursion limit.",
				},
				{
					Name: "Breadth-first search over amounts",
					Code: `from collections import deque

def coinChange(coins, amount):
    coins_used = {0: 0}               # amount -> fewest coins, found in BFS order
    queue = deque([0])
    while queue:
        a = queue.popleft()
        if a == amount:
            return coins_used[a]
        for coin in coins:
            nxt = a + coin
            if nxt <= amount and nxt not in coins_used:
                coins_used[nxt] = coins_used[a] + 1
                queue.append(nxt)
    return -1`,
					TimeComplexity:  "O(amount * k)",
					SpaceComplexity: "O(amount)",
					Explanation:     "Treat each amount as a graph node and each coin as an edge of length 1. BFS (Chapter 13) reaches each amount first by the fewest edges, which is the fewest coins. It explores the same states as the DP in a different order, a good reminder that 'fewest steps' problems are often both shortest paths and DP.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does greedy fail for coins [1,3,4] and amount 6, and why can the DP not make the same mistake?", Answer: "Greedy commits to the biggest coin, 4, and is then stuck making 2 from 1s: 4+1+1 = 3 coins. The DP never commits: for amount 6 it compares the last coin 1 (fewest[5] + 1 = 3), 3 (fewest[3] + 1 = 2) and 4 (fewest[2] + 1 = 3), and keeps 2. It considers every possible last coin, so it cannot miss the optimum."},
			{Prompt: "What changes if the question asks for the NUMBER of combinations that make the amount (order not mattering) instead of the fewest coins?", Answer: "The state becomes a count and the recurrence a sum, but to avoid counting 1+2 and 2+1 separately the loops must be swapped: for each coin, for each amount from coin upward, ways[a] += ways[a - coin], with ways[0] = 1. Processing one coin type at a time fixes an order on the coins. That is Coin Change II, in the 2-D DP chapter."},
		},
	},
	{
		ID:              "word-break",
		Number:          139,
		Order:           6,
		Title:           "Word Break",
		Slug:            "word-break",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"String", "Dynamic Programming", "Hash Table", "Trie", "Memoization"},
		RelatedChapters: []int{15, 10},
		Description: `Given a string **s** and a list of distinct words **wordDict**, return true if s can be cut into a sequence of one or more pieces where every piece is a word from wordDict, and false otherwise.

The same dictionary word may be used any number of times, and not every word has to be used.`,
		Constraints: []string{
			"1 <= s.length <= 300",
			"1 <= wordDict.length <= 1000",
			"1 <= wordDict[i].length <= 20",
			"s and wordDict[i] contain only lowercase English letters",
			"All words in wordDict are distinct",
		},
		Examples: []Example{
			{Input: `s = "leetcode", wordDict = ["leet","code"]`, Output: "true", Explanation: `"leet" + "code".`},
			{Input: `s = "applepenapple", wordDict = ["apple","pen"]`, Output: "true", Explanation: `"apple" + "pen" + "apple": a word may be reused.`},
			{Input: `s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]`, Output: "false", Explanation: `"cats" + "and" leaves "og", and "cat" + "sand" also leaves "og"; no cut covers it.`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "leetcode", "wordDict": []string{"leet", "code"}}, Expected: true, Note: "example 1"},
			{Input: map[string]interface{}{"s": "applepenapple", "wordDict": []string{"apple", "pen"}}, Expected: true, Note: "example 2: a word reused"},
			{Input: map[string]interface{}{"s": "catsandog", "wordDict": []string{"cats", "dog", "sand", "and", "cat"}}, Expected: false, Note: "example 3: a valid start leads to a dead end"},
			{Input: map[string]interface{}{"s": "a", "wordDict": []string{"a"}}, Expected: true, Hidden: true, Note: "one character that is a word"},
			{Input: map[string]interface{}{"s": "a", "wordDict": []string{"b"}}, Expected: false, Hidden: true, Note: "one character that is not a word"},
			{Input: map[string]interface{}{"s": "ab", "wordDict": []string{"a"}}, Expected: false, Hidden: true, Note: "a prefix matches but the rest does not"},
			{Input: map[string]interface{}{"s": "cars", "wordDict": []string{"car", "ca", "rs"}}, Expected: true, Hidden: true, Note: "taking the longest word first (car) fails; ca + rs works"},
			{Input: map[string]interface{}{"s": "aaaaaaa", "wordDict": []string{"aaaa", "aaa"}}, Expected: true, Hidden: true, Note: "3 + 4: needs more than one word length"},
			{Input: map[string]interface{}{"s": "goalspecial", "wordDict": []string{"go", "goal", "goals", "special"}}, Expected: true, Hidden: true, Note: "only one of three prefixes (goal) leads to a full split"},
			{Input: map[string]interface{}{"s": "aaaaaaaaaaaaaaab", "wordDict": []string{"a", "aa", "aaa", "aaaa"}}, Expected: false, Hidden: true, Note: "many ways to cover the a's, but no word ends in b"},
			{Input: map[string]interface{}{"s": "bb", "wordDict": []string{"a", "b", "bbb", "bbbb"}}, Expected: true, Hidden: true, Note: "dictionary words longer than s"},
		},
		TimeComplexity:  "O(n^3)",
		SpaceComplexity: "O(n + m)",
		StarterCode:     "def wordBreak(s, wordDict):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If s can be cut into words, the last piece is some dictionary word ending at the end of s. What must then be true of the part of s in front of that last word?"},
			{Level: 2, Type: "insight", Content: "The prefix in front of the last word must itself be breakable. So 'can the prefix of length i be broken?' depends only on shorter prefixes: can[i] is true if there is a split point j < i with can[j] true and s[j:i] in the dictionary. This is Climbing Stairs where the allowed step lengths are the dictionary words (Chapter 15).", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Define the state: can[i] = True if the prefix s[:i] can be cut into dictionary words.\n2. Write the recurrence: can[i] is True if some j < i has can[j] True and s[j:i] in the word set.\n3. Set the base cases: can[0] = True (the empty prefix needs no words).\n4. Choose the evaluation order: i from 1 to n.\n5. Locate the answer: can[n].\n6. Optimise space: every earlier entry may be read, so keep the table; you can speed it up by only trying j >= i - (longest word length)."},
			{Level: 4, Type: "skeleton", Content: "def wordBreak(s, wordDict):\n    words = set(wordDict)\n    n = len(s)\n    can = [False] * (n + 1)\n    can[0] = ____\n    for i in range(1, n + 1):\n        for j in range(i):\n            if ____:\n                can[i] = True\n                break\n    return ____"},
		},
		Solution: Solution{
			Code: `def wordBreak(s, wordDict):
    words = set(wordDict)
    n = len(s)
    # State: can[i] = True if the prefix s[:i] splits into dictionary words
    can = [False] * (n + 1)
    # Base case: the empty prefix needs no words
    can[0] = True
    # Evaluation order: shorter prefixes first
    for i in range(1, n + 1):
        # Recurrence: the last word is s[j:i] for some split point j
        for j in range(i):
            if can[j] and s[j:i] in words:
                can[i] = True
                break
    # Answer: the whole string
    return can[n]`,
			Explanation: "Decode Ways (problem 4) let the last piece be 1 or 2 characters long, subject to a validity check. Here the last piece can be any length, and the check is 'is it a dictionary word?'. The state also changes type: a yes/no (boolean) instead of a count, combined with 'or' instead of '+'.",
			BruteForce:  "Recursively try every dictionary word that matches the front of s and recurse on the rest. When a cut fails, the same suffix is retried from many different earlier cuts: for 'aaaa...ab' with words a, aa, aaa, aaaa, the number of attempts grows roughly like 1.9^n, which is far too slow for n = 300.",
			KeyInsight:  "Whether the rest of the string can be broken depends only on where it starts, not on how the front was cut. So there are only n + 1 subproblems (one per prefix length), each answered once.",
			WhyItWorks:  "Subproblem: can[i] says whether s[:i] can be cut into dictionary words. Recurrence: if s[:i] can be cut, its last piece is s[j:i] for some j < i, that piece is a word, and s[:j] is cut into words, so can[j] is true. Conversely, if can[j] is true and s[j:i] is a word, appending that word gives a cut of s[:i]. Trying every j therefore decides can[i] exactly. Base case can[0] = True lets a first word s[0:i] count. The loop fills i in increasing order and only reads can[j] for j < i, which is already final.",
			CommonMistakes: []string{
				"Greedily cutting the longest (or shortest) matching word and never backtracking. For 'cars' with ['car','ca','rs'], taking 'car' leaves 's' and fails, but 'ca' + 'rs' works. The DP keeps every reachable cut point.",
				"Checking 's[j:i] in wordDict' on the list. Each check is then a scan of up to 1000 words; convert to a set once.",
				"Starting with can[0] = False. Then nothing can ever become true and every test returns False; the empty prefix is trivially breakable.",
			},
			TimeComplexity:  "O(n^3)",
			SpaceComplexity: "O(n + m)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "can[i] answers the yes/no question for the prefix of length i. A set of words makes each dictionary lookup O(1) on average (plus the cost of hashing the substring).", CodeSnippet: "words = set(wordDict)\ncan = [False] * (n + 1)", LineStart: 2, LineEnd: 5},
				{Title: "Write the recurrence", Explanation: "Try every position j where the last word could start. If the part before it is breakable and s[j:i] is a word, the prefix of length i is breakable; one witness is enough, so break.", CodeSnippet: "for j in range(i):\n    if can[j] and s[j:i] in words:\n        can[i] = True\n        break", LineStart: 10, LineEnd: 14},
				{Title: "Set the base cases", Explanation: "The empty prefix is breakable (zero words), which lets the first word s[0:i] count.", CodeSnippet: "can[0] = True", LineStart: 6, LineEnd: 7},
				{Title: "Choose the evaluation order", Explanation: "can[i] reads can[j] for j < i only, so fill left to right. For 'leetcode' the True cells are 0, 4 ('leet') and 8 ('leet' + 'code').", CodeSnippet: "for i in range(1, n + 1):", LineStart: 8, LineEnd: 9},
				{Title: "Locate the answer", Explanation: "The whole string is the prefix of length n.", CodeSnippet: "return can[n]", LineStart: 15, LineEnd: 16},
				{Title: "Optimise space", Explanation: "can[i] may read any earlier cell, so the table stays. What can be cut is time: a word is at most 20 letters, so only j >= i - 20 can work, making the loops O(n * L) slicing steps for word length L.", CodeSnippet: "for j in range(i):", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Plain recursion (brute force)",
					Code: `def wordBreak(s, wordDict):
    def can_from(start):
        # can the suffix s[start:] be cut into words?
        if start == len(s):
            return True
        for word in wordDict:
            if s.startswith(word, start) and can_from(start + len(word)):
                return True
        return False
    return can_from(0)`,
					TimeComplexity:  "O(m^n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Tries each word at the front, then recurses on what is left. On strings like 'aaaa...ab' the same suffix is retried from many different earlier cuts, so it is exponential; it only survives the small tests here.",
				},
				{
					Name: "Top-down memoization over start positions",
					Code: `def wordBreak(s, wordDict):
    memo = {}
    def can_from(start):
        if start == len(s):
            return True
        if start not in memo:
            memo[start] = any(
                s.startswith(word, start) and can_from(start + len(word))
                for word in wordDict
            )
        return memo[start]
    return can_from(0)`,
					TimeComplexity:  "O(n * m * L)",
					SpaceComplexity: "O(n)",
					Explanation:     "The recursion plus a memo keyed by the start position: each suffix is decided once. This loops over words instead of split points, which is faster when the dictionary is small, with m words of length up to L.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to break out of the j loop as soon as one split works?", Answer: "can[i] is a yes/no question: one valid way to cut s[:i] is enough to make it True, and no later j can make it 'more true'. If the problem asked for the number of ways to cut s, you would have to add up can-counts for every valid j instead of stopping at the first."},
			{Prompt: "What changes if you must return every possible sentence (Word Break II) instead of true or false?", Answer: "The output itself can be exponential (for 'aaaa...' with words a and aa there are Fibonacci-many sentences), so no DP can be polynomial in the output. You would backtrack (Chapter 12) to list the sentences, and use this DP or a memo of 'is the suffix breakable?' to prune dead ends early."},
		},
	},
	{
		ID:              "maximum-product-subarray",
		Number:          152,
		Order:           7,
		Title:           "Maximum Product Subarray",
		Slug:            "maximum-product-subarray",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{15},
		Description: `Given an integer array **nums**, find a non-empty **contiguous** subarray (one or more neighbouring elements) whose product is as large as possible, and return that product.

The numbers can be negative or zero. The answer is guaranteed to fit in a 32-bit integer.`,
		Constraints: []string{
			"1 <= nums.length <= 2 * 10^4",
			"-10 <= nums[i] <= 10",
			"The product of any subarray fits in a 32-bit integer",
		},
		Examples: []Example{
			{Input: "nums = [2,3,-2,4]", Output: "6", Explanation: "[2,3] has product 6. Including -2 makes the product negative."},
			{Input: "nums = [-2,0,-1]", Output: "0", Explanation: "[-2,-1] is not contiguous (the 0 is between them), so the best is the subarray [0]."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 3, -2, 4}}, Expected: 6, Note: "example 1"},
			{Input: map[string]interface{}{"nums": []int{-2, 0, -1}}, Expected: 0, Note: "example 2: zero splits the array"},
			{Input: map[string]interface{}{"nums": []int{-2}}, Expected: -2, Note: "a single negative number: the answer can be negative"},
			{Input: map[string]interface{}{"nums": []int{-2, 3, -4}}, Expected: 24, Hidden: true, Note: "two negatives cancel: tracking only the max gives 3"},
			{Input: map[string]interface{}{"nums": []int{0, 2}}, Expected: 2, Hidden: true, Note: "restart after a zero"},
			{Input: map[string]interface{}{"nums": []int{-1, -2, -3}}, Expected: 6, Hidden: true, Note: "all negative: the best pair is -2 * -3"},
			{Input: map[string]interface{}{"nums": []int{2, -5, -2, -4, 3}}, Expected: 24, Hidden: true, Note: "best is -2 * -4 * 3, not starting at the first element"},
			{Input: map[string]interface{}{"nums": []int{-2, 0}}, Expected: 0, Hidden: true, Note: "zero beats a lone negative"},
			{Input: map[string]interface{}{"nums": []int{3, -1, 4}}, Expected: 4, Hidden: true, Note: "a single element beats the whole product (-12)"},
			{Input: map[string]interface{}{"nums": []int{-4, -3, -2}}, Expected: 12, Hidden: true, Note: "an odd count of negatives: drop one end"},
			{Input: map[string]interface{}{"nums": []int{1, -2, 3, -4, 5, 0, 6}}, Expected: 120, Hidden: true, Note: "the best run lies before a zero"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def maxProduct(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose you know the largest product of a subarray ending just before position i, and nums[i] is negative. Is extending that largest product still the best move? Which earlier product would you rather have?"},
			{Level: 2, Type: "insight", Content: "Multiplying by a negative number turns the biggest product into the smallest and the smallest (most negative) into the biggest. So one number per position is not enough: keep both the largest and the smallest product of a subarray ending at i. Each is the best of three candidates: nums[i] alone, previous max * nums[i], previous min * nums[i]. The state-design idea is from Chapter 15.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Define the state: high[i] and low[i] = largest and smallest product of a subarray that ends exactly at i.\n2. Write the recurrence: the subarray ending at i either starts at i or extends one ending at i-1; candidates are x, high[i-1] * x, low[i-1] * x. high[i] is their max, low[i] their min.\n3. Set the base cases: high[0] = low[0] = nums[0].\n4. Choose the evaluation order: left to right.\n5. Locate the answer: the best subarray can end anywhere, so return max(high).\n6. Optimise space: keep only the previous high and low, plus a running best."},
			{Level: 4, Type: "skeleton", Content: "def maxProduct(nums):\n    n = len(nums)\n    high = [0] * n\n    low = [0] * n\n    high[0] = low[0] = nums[0]\n    for i in range(1, n):\n        x = nums[i]\n        candidates = (x, ____, ____)\n        high[i] = ____\n        low[i] = ____\n    return ____"},
		},
		Solution: Solution{
			Code: `def maxProduct(nums):
    n = len(nums)
    # State: two numbers per index, for subarrays that END at i:
    #   high[i] = largest product, low[i] = smallest (most negative) product
    high = [0] * n
    low = [0] * n
    # Base case: the only subarray ending at 0 is [nums[0]]
    high[0] = low[0] = nums[0]
    # Evaluation order: left to right
    for i in range(1, n):
        x = nums[i]
        # Recurrence: start fresh at x, or extend the best/worst run ending at i-1
        candidates = (x, high[i - 1] * x, low[i - 1] * x)
        high[i] = max(candidates)
        low[i] = min(candidates)
    # Answer: the best subarray can end anywhere
    return max(high)`,
			Explanation: "Earlier problems needed one number per position. Here the new idea is that the state sometimes needs more than one value: a negative number swaps the roles of the largest and smallest products, so we carry both. It also introduces 'ends exactly at i' states, where the answer is the best over all positions rather than the last cell (you will use that again in Longest Increasing Subsequence).",
			BruteForce:  "Try every start i and extend the end j one step at a time, multiplying as you go and keeping the best: O(n^2) products, about 2 x 10^8 for n = 20,000. The waste is that the product for (i, j) and (i', j) share all the elements after max(i, i'), yet each start recomputes them.",
			KeyInsight:  "The best product ending at i only depends on the extreme products ending at i-1: the largest (if x is positive) or the smallest (if x is negative), or neither (start fresh at x, e.g. after a zero). Keeping max and min per position captures all three cases.",
			WhyItWorks:  "Subproblem: high[i] and low[i] are the largest and smallest products over subarrays ending exactly at i. Recurrence: such a subarray is either [x] alone or a subarray ending at i-1 extended by x. Multiplying by x is monotonic: for x >= 0 it preserves order, for x < 0 it reverses it. So the largest extended product is either high[i-1] * x or low[i-1] * x, and the same two values bound the smallest; comparing them with x alone covers every case, including zeros. Base case: index 0 has one subarray. Every subarray ends somewhere, so max(high) over all i is the answer.",
			CommonMistakes: []string{
				"Tracking only the maximum, as in maximum subarray sum (Kadane). For [-2,3,-4] the max-only version gives 3, but -2 * 3 * -4 = 24: the most negative product becomes the largest after another negative.",
				"Updating high before computing low from the old high: high = max(x, high * x, low * x) followed by low = min(x, high * x, low * x) uses the new high. Compute the candidates once from the previous values, as in the solution.",
				"Starting the answer at 0 instead of nums[0]. For [-2] the answer is -2; there is no empty subarray with product 0 to fall back on.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "For each index, store two numbers about subarrays that end there: the largest and the smallest product. The smallest matters because a later negative number can flip it into the largest.", CodeSnippet: "high = [0] * n\nlow = [0] * n", LineStart: 3, LineEnd: 6},
				{Title: "Write the recurrence", Explanation: "A subarray ending at i is either just [x] or one ending at i-1 extended by x. Of the extended ones, the extremes come from high[i-1] * x and low[i-1] * x. Take the max and min of the three candidates.", CodeSnippet: "candidates = (x, high[i - 1] * x, low[i - 1] * x)\nhigh[i] = max(candidates)\nlow[i] = min(candidates)", LineStart: 12, LineEnd: 15},
				{Title: "Set the base cases", Explanation: "At index 0 the only subarray is [nums[0]], so it is both the largest and smallest.", CodeSnippet: "high[0] = low[0] = nums[0]", LineStart: 7, LineEnd: 8},
				{Title: "Choose the evaluation order", Explanation: "Each index reads index i-1 only, so go left to right. For [-2,3,-4]: high = [-2, 3, 24], low = [-2, -6, -12]. At i = 2, low[1] * -4 = 24 becomes the new high.", CodeSnippet: "for i in range(1, n):", LineStart: 9, LineEnd: 11},
				{Title: "Locate the answer", Explanation: "The states are about subarrays ending at i, and the best subarray can end anywhere, so the answer is the maximum over the whole high table, not high[n-1].", CodeSnippet: "return max(high)", LineStart: 16, LineEnd: 17},
				{Title: "Optimise space", Explanation: "Only high[i-1] and low[i-1] are read, so two variables plus a running best are enough; see the 'Rolling variables' alternative.", CodeSnippet: "candidates = (x, high[i - 1] * x, low[i - 1] * x)", LineStart: 13, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Every subarray (brute force)",
					Code: `def maxProduct(nums):
    best = nums[0]
    for start in range(len(nums)):
        product = 1
        for end in range(start, len(nums)):
            product *= nums[end]          # product of nums[start..end]
            best = max(best, product)
    return best`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Extends each start one element at a time so each product costs O(1) to update. Simple and clearly correct, a good checker, but about 2 x 10^8 multiplications at the maximum size.",
				},
				{
					Name: "Rolling variables (optimise space)",
					Code: `def maxProduct(nums):
    high = low = best = nums[0]
    for x in nums[1:]:
        candidates = (x, high * x, low * x)
        high, low = max(candidates), min(candidates)
        best = max(best, high)
    return best`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "The table version keeping only the previous high and low. Because the answer is a max over all positions, a separate running best replaces max(high). Building the candidates tuple before assigning keeps the old values intact.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the state include the smallest product, when the problem only asks for the largest?", Answer: "Because a negative factor reverses order. If the smallest product ending at i-1 is -12 and nums[i] is -2, extending it gives 24, larger than anything built from the largest product. Without the minimum, the DP has no way to know that value existed, so it would violate optimal substructure: the best answer at i would not be built from the stored best at i-1."},
			{Prompt: "What role does the candidate 'x alone' play, and which test fails without it?", Answer: "It lets a subarray start fresh at i. After a zero both high and low become 0, and without 'x alone' every later product would stay 0. For [0,2] the answer 2 needs it, and for [-2,0,-1] it lets -1 start a new run."},
		},
	},
	{
		ID:              "longest-increasing-subsequence",
		Number:          300,
		Order:           8,
		Title:           "Longest Increasing Subsequence",
		Slug:            "longest-increasing-subsequence",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Dynamic Programming", "Binary Search"},
		RelatedChapters: []int{15, 6},
		Description: `Given an integer array **nums**, return the length of its longest **strictly increasing subsequence**.

A subsequence keeps some of the elements (at least one) in their original order, possibly skipping others; the kept elements do not have to be neighbours. Strictly increasing means every kept element is larger than the one before it, so equal values cannot both be used.`,
		Constraints: []string{
			"1 <= nums.length <= 2500",
			"-10^4 <= nums[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "nums = [10,9,2,5,3,7,101,18]", Output: "4", Explanation: "One longest choice is 2, 3, 7, 101 (another is 2, 5, 7, 18)."},
			{Input: "nums = [0,1,0,3,2,3]", Output: "4", Explanation: "0, 1, 2, 3."},
			{Input: "nums = [7,7,7,7,7,7,7]", Output: "1", Explanation: "Equal values are not increasing, so only one 7 can be used."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{10, 9, 2, 5, 3, 7, 101, 18}}, Expected: 4, Note: "example 1"},
			{Input: map[string]interface{}{"nums": []int{0, 1, 0, 3, 2, 3}}, Expected: 4, Note: "example 2"},
			{Input: map[string]interface{}{"nums": []int{7, 7, 7, 7, 7, 7, 7}}, Expected: 1, Note: "example 3: strictly increasing, so <= instead of < gives 7"},
			{Input: map[string]interface{}{"nums": []int{5}}, Expected: 1, Hidden: true, Note: "a single element"},
			{Input: map[string]interface{}{"nums": []int{5, 4, 3, 2, 1}}, Expected: 1, Hidden: true, Note: "strictly decreasing"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5}}, Expected: 5, Hidden: true, Note: "already increasing: the answer ends at the last element"},
			{Input: map[string]interface{}{"nums": []int{4, 10, 4, 3, 8, 9}}, Expected: 3, Hidden: true, Note: "the LIS does not end at the largest value (10)"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 6, 7, 9, 4, 10, 5, 6}}, Expected: 6, Hidden: true, Note: "the answer is max(ends_at), not ends_at[n-1] (which is 5)"},
			{Input: map[string]interface{}{"nums": []int{-2, -1}}, Expected: 2, Hidden: true, Note: "negative values"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 3, 3, 4, 4}}, Expected: 3, Hidden: true, Note: "duplicates mixed with increases"},
			{Input: map[string]interface{}{"nums": []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}}, Expected: 6, Hidden: true, Note: "a later, lower start (2, 4, 5, 6, 7, 12) beats the early run 3, 5, 6, 19"},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def lengthOfLIS(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "'Longest increasing subsequence of the first i elements' is awkward, because you don't know what it ends with. What if you fix the last element? If a subsequence must end with nums[i], what could the element just before it be?"},
			{Level: 2, Type: "insight", Content: "Let ends_at[i] be the length of the longest increasing subsequence that ends with nums[i]. The element before nums[i] is some nums[j] with j < i and nums[j] < nums[i], so ends_at[i] = 1 + max(ends_at[j]) over those j (or 1 if there are none). The overall answer can end anywhere: max(ends_at). Chapter 15 covers 'ends at i' states.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Define the state: ends_at[i] = length of the longest strictly increasing subsequence ending with nums[i].\n2. Write the recurrence: for every j < i with nums[j] < nums[i], ends_at[i] = max(ends_at[i], ends_at[j] + 1).\n3. Set the base cases: every ends_at[i] starts at 1 (the element alone).\n4. Choose the evaluation order: i from left to right; all j < i are final.\n5. Locate the answer: max(ends_at), since the subsequence can end anywhere.\n6. Optimise space: the table is O(n) already; a faster O(n log n) method keeps smallest tails and binary-searches them."},
			{Level: 4, Type: "skeleton", Content: "def lengthOfLIS(nums):\n    n = len(nums)\n    ends_at = [____] * n\n    for i in range(n):\n        for j in range(i):\n            if ____:\n                ends_at[i] = max(ends_at[i], ____)\n    return ____"},
		},
		Solution: Solution{
			Code: `def lengthOfLIS(nums):
    n = len(nums)
    # State: ends_at[i] = length of the longest strictly increasing
    #        subsequence that ENDS with nums[i]
    # Base case: every element alone is a subsequence of length 1
    ends_at = [1] * n
    # Evaluation order: left to right, so every earlier j is final
    for i in range(n):
        # Recurrence: the element before nums[i] is some smaller nums[j], j < i
        for j in range(i):
            if nums[j] < nums[i]:
                ends_at[i] = max(ends_at[i], ends_at[j] + 1)
    # Answer: the LIS can end anywhere, so take the maximum
    return max(ends_at)`,
			Explanation: "This builds on Maximum Product Subarray (problem 7), which introduced states that 'end exactly at i' with the answer taken as a max over all i. The new idea is that the previous element is not necessarily the neighbour: a subsequence can jump, so the recurrence looks back at every earlier index j, which makes this the first O(n^2) DP of the ladder.",
			BruteForce:  "Try every subsequence (keep or skip each element) and keep the longest increasing one: 2^n subsets, about 10^752 for n = 2500. Recursively, 'longest from i given the previous kept value' recomputes the same suffix for the same previous value many times.",
			KeyInsight:  "Fixing the last element makes the subproblem well defined: the best subsequence ending with nums[i] only depends on the best ones ending at earlier, smaller elements. There are n such subproblems, each combining at most n earlier answers.",
			WhyItWorks:  "Subproblem: ends_at[i] is the length of the longest strictly increasing subsequence whose last element is nums[i]. Recurrence: such a subsequence is either nums[i] alone (length 1) or a strictly increasing subsequence ending at some j < i with nums[j] < nums[i], followed by nums[i]. The prefix ending at j must itself be as long as possible, or we could lengthen the whole thing, so ends_at[i] = 1 + max ends_at[j] over valid j. Every j is tried, so the max is complete. Evaluation in increasing i means each ends_at[j] is final when read. Every subsequence ends at some index, so the answer is max(ends_at).",
			CommonMistakes: []string{
				"Using nums[j] <= nums[i]. The problem says strictly increasing, so for [7,7,7,7,7,7,7] the answer is 1, not 7.",
				"Returning ends_at[n-1]. That is the longest subsequence that must end with the last element; for [1,3,6,7,9,4,10,5,6] it is 5 (1,3,4,5,6), but the answer is 6 (1,3,6,7,9,10, ending at index 6).",
				"In the binary-search version, believing the tails list is an actual increasing subsequence. It is not: only its length is meaningful. For [3,4,1] the final tails is [1,4], which is not a subsequence of the input in order.",
			},
			TimeComplexity:  "O(n^2)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "Fix the last element: ends_at[i] is the best length of a strictly increasing subsequence that finishes with nums[i]. Knowing the last element is what makes the next comparison possible.", CodeSnippet: "ends_at = [1] * n", LineStart: 3, LineEnd: 6},
				{Title: "Write the recurrence", Explanation: "The element before nums[i] can be any earlier, strictly smaller element. Extend the best subsequence ending there by one.", CodeSnippet: "for j in range(i):\n    if nums[j] < nums[i]:\n        ends_at[i] = max(ends_at[i], ends_at[j] + 1)", LineStart: 9, LineEnd: 12},
				{Title: "Set the base cases", Explanation: "Every element alone is an increasing subsequence of length 1, so every cell starts at 1. That also handles elements with no smaller element before them.", CodeSnippet: "ends_at = [1] * n", LineStart: 5, LineEnd: 6},
				{Title: "Choose the evaluation order", Explanation: "ends_at[i] reads only j < i, so go left to right. For [10,9,2,5,3,7,101,18] the table is [1, 1, 1, 2, 2, 3, 4, 4].", CodeSnippet: "for i in range(n):", LineStart: 7, LineEnd: 8},
				{Title: "Locate the answer", Explanation: "The longest subsequence can end at any index, so take the maximum over the table rather than the last cell.", CodeSnippet: "return max(ends_at)", LineStart: 13, LineEnd: 14},
				{Title: "Optimise space", Explanation: "The table is already O(n) space; the cost is O(n^2) time from the inner loop. The 'patience sorting' alternative changes the state to 'smallest tail for each length' and binary-searches it, for O(n log n) time.", CodeSnippet: "for j in range(i):", LineStart: 10, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{
					Name: "Take-or-skip recursion (brute force)",
					Code: `def lengthOfLIS(nums):
    def longest_from(i, prev):
        # longest increasing subsequence of nums[i:] using only values > prev
        if i == len(nums):
            return 0
        skip = longest_from(i + 1, prev)
        if nums[i] > prev:
            return max(skip, 1 + longest_from(i + 1, nums[i]))
        return skip
    return longest_from(0, float("-inf"))`,
					TimeComplexity:  "O(2^n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Decide for each element whether to keep it, remembering the last kept value. Explores every subsequence, so it is exponential; memoizing on (i, index of prev) gives an O(n^2) top-down DP equivalent to the table.",
				},
				{
					Name: "Smallest tails with binary search (patience sorting)",
					Code: `from bisect import bisect_left

def lengthOfLIS(nums):
    # tails[k] = smallest last value of any increasing subsequence of length k + 1
    tails = []
    for x in nums:
        k = bisect_left(tails, x)     # first tail >= x; equal values replace (strict)
        if k == len(tails):
            tails.append(x)           # x extends the longest subsequence so far
        else:
            tails[k] = x              # x is a smaller ending for length k + 1
    return len(tails)`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "A different state: for each length L, the smallest value an increasing subsequence of length L can end with. A smaller ending is always at least as good, because more future values can extend it. tails stays sorted, so binary search (Chapter 6) finds where x fits: append it if it beats every tail, otherwise lower the first tail >= x. bisect_left (not bisect_right) makes equal values replace instead of extend, which enforces 'strictly'. Important: tails is not itself an LIS. For [3,4,1] it ends as [1,4], which does not appear in that order in the input; only len(tails) is the answer.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why define the state as 'ending at i' instead of 'the best within the first i elements'?", Answer: "To extend a subsequence you must know its last value, since the next element must be larger. 'Best within the first i' loses that information: two subsequences of the same length with different last values behave differently later. Fixing the last element to nums[i] makes every state carry exactly what the recurrence needs, at the price of taking max(ends_at) at the end."},
			{Prompt: "In the patience-sorting version, why does replacing tails[k] with a smaller x never make the answer wrong, even though tails stops being a real subsequence?", Answer: "tails[k] only claims 'some increasing subsequence of length k+1 ends with this value'. When x replaces it, that claim stays true: the subsequence of length k ending at tails[k-1] < x, followed by x, has length k+1. The lengths stored never change except by appending, and appending happens exactly when a longer subsequence exists, so len(tails) is the LIS length even though the list mixes values from different subsequences."},
		},
	},
	{
		ID:              "partition-equal-subset-sum",
		Number:          416,
		Order:           9,
		Title:           "Partition Equal Subset Sum",
		Slug:            "partition-equal-subset-sum",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{15, 16},
		Description: `Given an array **nums** of positive integers, decide whether it can be split into two groups with equal sums. Every element must go into exactly one of the two groups (the groups do not have to be contiguous or the same size).

Return true if such a split exists, and false otherwise.`,
		Constraints: []string{
			"1 <= nums.length <= 200",
			"1 <= nums[i] <= 100",
		},
		Examples: []Example{
			{Input: "nums = [1,5,11,5]", Output: "true", Explanation: "[1,5,5] and [11] both sum to 11."},
			{Input: "nums = [1,2,3,5]", Output: "false", Explanation: "The total is 11, which is odd, so two equal integer halves are impossible."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 5, 11, 5}}, Expected: true, Note: "example 1"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 5}}, Expected: false, Note: "example 2: odd total"},
			{Input: map[string]interface{}{"nums": []int{1}}, Expected: false, Hidden: true, Note: "a single element can never be split"},
			{Input: map[string]interface{}{"nums": []int{2, 2}}, Expected: true, Hidden: true, Note: "two equal elements"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 5}}, Expected: false, Hidden: true, Note: "even total 8 but no subset makes 4; reusing the 2 (2+2) wrongly says true"},
			{Input: map[string]interface{}{"nums": []int{3, 3, 3, 4, 5}}, Expected: true, Hidden: true, Note: "4 + 5 = 3 + 3 + 3"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 3, 5}}, Expected: false, Hidden: true, Note: "even total 12, half 6 unreachable"},
			{Input: map[string]interface{}{"nums": []int{1, 1}}, Expected: true, Hidden: true, Note: "smallest true case"},
			{Input: map[string]interface{}{"nums": []int{100, 100, 100, 100, 100, 99, 97, 2}}, Expected: false, Hidden: true, Note: "large values, even total 698, half 349 unreachable"},
			{Input: map[string]interface{}{"nums": []int{14, 9, 8, 4, 3, 2}}, Expected: true, Hidden: true, Note: "14 + 6 = 9 + 8 + 3 (half is 20)"},
		},
		TimeComplexity:  "O(n * target)",
		SpaceComplexity: "O(target)",
		StarterCode:     "def canPartition(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If the two groups have equal sums, what must each group sum to? Once you know that number, what simpler question about subsets are you really asking?"},
			{Level: 2, Type: "insight", Content: "Each group must sum to half the total (so an odd total is impossible). The question becomes: is there a subset that sums to exactly target = total / 2? Process the numbers one at a time and keep, for every sum s from 0 to target, whether some subset of the numbers so far reaches s. Each number is used at most once: this is 0/1 subset sum, a DP indexed by a sum instead of a position (Chapter 15).", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Check the total: if it is odd, return False; otherwise target = total // 2.\n2. Define the state: can[s] = True if some subset of the numbers processed so far sums to s.\n3. Write the recurrence: after adding num, s is reachable if it was already reachable (skip num) or s - num was reachable before num (take num).\n4. Set the base cases: can[0] = True (the empty subset).\n5. Choose the evaluation order: one number at a time, building the new row from the old one so num is used at most once.\n6. Locate the answer: can[target].\n7. Optimise space: with a single array, loop s from target down to num so you only read values from before num."},
			{Level: 4, Type: "skeleton", Content: "def canPartition(nums):\n    total = sum(nums)\n    if ____:\n        return False\n    target = total // 2\n    can = [False] * (target + 1)\n    can[0] = True\n    for num in nums:\n        new_can = can[:]\n        for s in range(num, target + 1):\n            if ____:\n                new_can[s] = True\n        can = new_can\n    return ____"},
		},
		Solution: Solution{
			Code: `def canPartition(nums):
    total = sum(nums)
    if total % 2 == 1:
        return False                 # an odd total can't split into equal halves
    target = total // 2
    # State: can[s] = True if some subset of the numbers seen so far sums to s
    can = [False] * (target + 1)
    # Base case: the empty subset makes 0
    can[0] = True
    # Evaluation order: one number at a time; each number builds a new row
    for num in nums:
        new_can = can[:]             # skipping num keeps every sum we had
        # Recurrence: s is reachable if it was, or if s - num was (take num)
        for s in range(num, target + 1):
            if can[s - num]:
                new_can[s] = True
        can = new_can
    # Answer: can we hit exactly half?
    return can[target]`,
			Explanation: "So far the state was indexed by a position in the input. The new idea here is a state indexed by a value (a sum), with the input processed one item at a time and each item used at most once: the 0/1 knapsack pattern. It builds on Coin Change (problem 5), which also had amounts as states, but there coins could be reused; here each number can be taken only once, which is why each number builds a fresh row from the previous one.",
			BruteForce:  "Try all 2^n ways to choose a subset (take or skip each number) and check whether any sums to total / 2. For n = 200 that is 2^200, about 1.6 x 10^60 subsets. Many different subsets reach the same partial sum after the same numbers, and the brute force explores what follows each of them separately.",
			KeyInsight:  "After processing some numbers, all that matters for the future is which sums are reachable, not which subsets reach them. Sums are at most target = total / 2 <= 10,000, so there are only target + 1 facts to keep per step.",
			WhyItWorks:  "Subproblem: after processing the first i numbers, can[s] says whether some subset of them sums to s (think of a 2-D table with one row per i; we keep only the current row). Recurrence: a subset of the first i numbers either skips num = nums[i-1] (then s was reachable in the previous row) or takes it (then s - num was reachable in the previous row). Both cases are covered: new_can starts as a copy of the old row, and the loop adds the 'take' case reading only the old row, so num is used at most once. Base case: before any numbers, only sum 0 is reachable. After all n numbers, can[target] says whether a subset sums to half the total; the remaining numbers then sum to the other half.",
			CommonMistakes: []string{
				"Using a single array and looping s upward: if can[s - num]: can[s] = True. Then can[s - num] may already include num, so a number is reused. For [1,2,5] that marks 4 as 2 + 2 and wrongly returns True. Loop s downward, or build a new row as here.",
				"Forgetting the odd-total check, then using total // 2 anyway: for [1,2,3,5] (total 11) the target becomes 5, which {5} reaches, so the answer comes out True instead of False.",
				"Sorting the numbers and greedily putting each one into the group with the smaller sum. For [4,5,6,7,8] (half 15), largest first: 8 and 7 go to different groups, 6 joins the 7 (13), 5 joins the 8 (13), and 4 makes 17 against 13, yet 7+8 = 4+5+6 works. Greedy commits too early; the DP keeps every reachable sum.",
			},
			TimeComplexity:  "O(n * target)",
			SpaceComplexity: "O(target)",
			Walkthrough: []WalkthroughStep{
				{Title: "Check the total", Explanation: "Two equal integer groups need an even total; each must sum to exactly half. This turns the problem into subset sum with target = total // 2.", CodeSnippet: "if total % 2 == 1:\n    return False\ntarget = total // 2", LineStart: 2, LineEnd: 5},
				{Title: "Define the state", Explanation: "One yes/no per possible sum 0..target: can some subset of the numbers seen so far make it? The state is indexed by a value, not by a position.", CodeSnippet: "can = [False] * (target + 1)", LineStart: 6, LineEnd: 7},
				{Title: "Set the base cases", Explanation: "Before any number is processed, only the empty subset exists, and it sums to 0.", CodeSnippet: "can[0] = True", LineStart: 8, LineEnd: 9},
				{Title: "Write the recurrence", Explanation: "For the new number num, every old sum stays reachable (skip num), and every s with s - num reachable in the OLD row becomes reachable (take num). Reading can (old) and writing new_can keeps the two apart.", CodeSnippet: "new_can = can[:]\nfor s in range(num, target + 1):\n    if can[s - num]:\n        new_can[s] = True", LineStart: 12, LineEnd: 16},
				{Title: "Choose the evaluation order", Explanation: "Process the numbers one at a time, each producing a new row from the previous one. For [1,5,11,5] (target 11) the reachable sums grow {0} -> {0,1} -> {0,1,5,6} -> {0,1,5,6,11} -> {0,1,5,6,10,11}, so can[11] is True.", CodeSnippet: "for num in nums:\n    can = new_can", LineStart: 11, LineEnd: 17},
				{Title: "Locate the answer", Explanation: "After every number has been considered, can[target] says whether half the total is reachable.", CodeSnippet: "return can[target]", LineStart: 18, LineEnd: 19},
				{Title: "Optimise space", Explanation: "The copy can be avoided: update a single array with s running from target down to num. Going downward, can[s - num] has not been updated for this num yet, so it still holds the old row's value. See the alternative.", CodeSnippet: "for s in range(num, target + 1):", LineStart: 14, LineEnd: 14},
			},
			Alternatives: []AltSolution{
				{
					Name: "Take-or-skip recursion (brute force)",
					Code: `def canPartition(nums):
    total = sum(nums)
    if total % 2 == 1:
        return False
    def reach(i, remaining):
        # can some subset of nums[i:] sum to exactly remaining?
        if remaining == 0:
            return True
        if i == len(nums) or remaining < 0:
            return False
        return reach(i + 1, remaining - nums[i]) or reach(i + 1, remaining)
    return reach(0, total // 2)`,
					TimeComplexity:  "O(2^n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Each number is taken or skipped, so every subset is explored in the worst case. Memoizing on (i, remaining) gives the top-down version of the table, with at most n * (target + 1) states.",
				},
				{
					Name: "One array, right to left (optimise space)",
					Code: `def canPartition(nums):
    total = sum(nums)
    if total % 2 == 1:
        return False
    target = total // 2
    can = [False] * (target + 1)
    can[0] = True
    for num in nums:
        # downward: can[s - num] still means 'reachable before num'
        for s in range(target, num - 1, -1):
            if can[s - num]:
                can[s] = True
    return can[target]`,
					TimeComplexity:  "O(n * target)",
					SpaceComplexity: "O(target)",
					Explanation:     "The standard 0/1 knapsack space trick. Updating in place from high s to low s means each read of can[s - num] (a smaller index) happens before that cell is updated for the current num, so num is used at most once without copying the row. Looping upward instead would turn this into the unbounded (Coin Change style) version and reuse numbers.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does Coin Change loop amounts upward, while the in-place version here must loop sums downward?", Answer: "In Coin Change a coin may be used any number of times, so reading fewest[a - coin] that already includes this coin is exactly right. Here each number may be used once. Looping downward means can[s - num] is read before it is updated for num, so it reflects only the earlier numbers; looping upward would let num be added again and again (for [1,2,5] it would reach 4 as 2 + 2)."},
			{Prompt: "The state here is indexed by a sum rather than a position. What does that do to the running time, and when would this approach be a bad idea?", Answer: "The work is O(n * target), proportional to the numeric size of the values, not just their count (pseudo-polynomial). Here target <= 200 * 100 / 2 = 10,000, so it is fast. If the numbers could be as large as 10^9, the table would be far too large and you would need a different method, such as meet-in-the-middle for small n."},
		},
	},
	{
		ID:              "palindromic-substrings",
		Number:          647,
		Order:           10,
		Title:           "Palindromic Substrings",
		Slug:            "palindromic-substrings",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"String", "Dynamic Programming", "Two Pointers"},
		RelatedChapters: []int{15, 16, 3},
		Description: `Given a string **s**, count how many of its substrings are palindromes. A palindrome reads the same forwards and backwards, and a substring is a contiguous, non-empty run of characters.

Substrings at different positions count separately even if they contain the same letters: in "aaa", each of the three single "a"s counts.`,
		Constraints: []string{
			"1 <= s.length <= 1000",
			"s consists of lowercase English letters",
		},
		Examples: []Example{
			{Input: `s = "abc"`, Output: "3", Explanation: `Only the single letters: "a", "b", "c".`},
			{Input: `s = "aaa"`, Output: "6", Explanation: `Three "a", two "aa" (positions 0-1 and 1-2) and one "aaa".`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "abc"}, Expected: 3, Note: "example 1: only single letters"},
			{Input: map[string]interface{}{"s": "aaa"}, Expected: 6, Note: "example 2: overlapping palindromes count separately"},
			{Input: map[string]interface{}{"s": "a"}, Expected: 1, Hidden: true, Note: "a single character"},
			{Input: map[string]interface{}{"s": "ab"}, Expected: 2, Hidden: true, Note: "two different letters"},
			{Input: map[string]interface{}{"s": "zz"}, Expected: 3, Hidden: true, Note: "an even-length palindrome: odd centres alone give 2"},
			{Input: map[string]interface{}{"s": "abba"}, Expected: 6, Hidden: true, Note: "even centre grows into abba"},
			{Input: map[string]interface{}{"s": "racecar"}, Expected: 10, Hidden: true, Note: "odd palindromes nested around one centre"},
			{Input: map[string]interface{}{"s": "aabaa"}, Expected: 9, Hidden: true, Note: "odd and even palindromes overlap"},
			{Input: map[string]interface{}{"s": "abacdfgdcaba"}, Expected: 14, Hidden: true, Note: "a reversed copy inside the string is not a palindrome"},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(n^2)",
		StarterCode:     "def countSubstrings(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you already know whether s[i+1..j-1] is a palindrome, what single extra check tells you whether s[i..j] is one?"},
			{Level: 2, Type: "insight", Content: "s[i..j] is a palindrome exactly when s[i] == s[j] and the inside s[i+1..j-1] is a palindrome (or the inside is empty or one letter). That is a recurrence from a smaller subproblem, so a table is_pal[i][j] filled from short substrings to long ones answers each substring in O(1). The state has two indices, a preview of Chapter 16, but the recipe is Chapter 15's.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Define the state: is_pal[i][j] = True if s[i..j] (inclusive) is a palindrome.\n2. Write the recurrence: is_pal[i][j] = (s[i] == s[j]) and is_pal[i+1][j-1].\n3. Set the base cases: length 1 is always a palindrome; length 2 is one if both letters match.\n4. Choose the evaluation order: by increasing length, so the inside is done first.\n5. Locate the answer: count the True cells.\n6. Optimise space: expanding around each of the 2n - 1 centres finds the same palindromes with O(1) extra space."},
			{Level: 4, Type: "skeleton", Content: "def countSubstrings(s):\n    n = len(s)\n    is_pal = [[False] * n for _ in range(n)]\n    count = 0\n    for length in range(1, n + 1):\n        for i in range(n - length + 1):\n            j = i + length - 1\n            if length <= 2:\n                is_pal[i][j] = ____\n            else:\n                is_pal[i][j] = ____\n            if is_pal[i][j]:\n                count += 1\n    return count"},
		},
		Solution: Solution{
			Code: `def countSubstrings(s):
    n = len(s)
    # State: is_pal[i][j] = True if s[i..j] (inclusive) reads the same both ways
    is_pal = [[False] * n for _ in range(n)]
    count = 0
    # Evaluation order: shorter substrings first, so is_pal[i+1][j-1] is ready
    for length in range(1, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            # Base cases: length 1 always; length 2 if both letters match
            if length <= 2:
                is_pal[i][j] = s[i] == s[j]
            # Recurrence: outer letters match and the inside is a palindrome
            else:
                is_pal[i][j] = s[i] == s[j] and is_pal[i + 1][j - 1]
            # Answer: count every True cell
            if is_pal[i][j]:
                count += 1
    return count`,
			Explanation: "Up to now every state was a single index (a prefix, an end position or a sum). The new idea is a state made of two indices, a substring s[i..j], whose subproblem is the substring with both ends removed. This is the bridge to Chapter 16. To be explicit about the alternatives: the popular 'expand around centres' solution is not a DP (it keeps no table of subproblems); it is what you get after the recipe's last step, when you notice the table is only ever read along one diagonal at a time and replace it with two pointers.",
			BruteForce:  "Check every one of the n(n+1)/2 substrings by comparing it with its reverse: O(n) per check, O(n^3) in total, about 1.7 x 10^8 character comparisons for n = 1000. The waste: checking 'racecar' re-checks 'aceca', 'cec' and 'e', which were already checked as substrings of their own.",
			KeyInsight:  "Whether s[i..j] is a palindrome depends only on its two end letters and on s[i+1..j-1], which is a shorter substring. Storing the answer for every substring makes each new check O(1).",
			WhyItWorks:  "Subproblem: is_pal[i][j] says whether s[i..j] is a palindrome. Recurrence: a string of length >= 3 reads the same backwards exactly when its first and last letters are equal and its middle part also reads the same backwards, so is_pal[i][j] = s[i] == s[j] and is_pal[i+1][j-1]. Base cases: a single letter is a palindrome, and a two-letter string is one iff the letters match (its 'inside' is empty). Evaluation order: the inside is two letters shorter, so filling by increasing length guarantees it is final. Each True cell is a distinct substring position, so counting them gives the answer.",
			CommonMistakes: []string{
				"Filling the table row by row with i increasing. is_pal[i][j] needs is_pal[i+1][j-1], which is in a later row and not yet computed. Fill by increasing length (or loop i from n-1 down to 0).",
				"In the centre-expansion version, using only odd centres (i, i). Even-length palindromes like 'zz' or 'abba' have their centre between two letters; for 'zz' odd centres alone give 2 instead of 3.",
				"Counting distinct palindromic strings instead of positions. 'aaa' has 6 palindromic substrings (three 'a', two 'aa', one 'aaa'), not 3.",
			},
			TimeComplexity:  "O(n^2)",
			SpaceComplexity: "O(n^2)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "One yes/no per substring, identified by its start i and end j. For n = 1000 that is a million cells, which is fine.", CodeSnippet: "is_pal = [[False] * n for _ in range(n)]", LineStart: 3, LineEnd: 4},
				{Title: "Write the recurrence", Explanation: "Peel off the two end letters: they must match, and what is left inside must already be a palindrome.", CodeSnippet: "is_pal[i][j] = s[i] == s[j] and is_pal[i + 1][j - 1]", LineStart: 13, LineEnd: 15},
				{Title: "Set the base cases", Explanation: "Substrings of length 1 and 2 have no 'inside' to look up: length 1 is always a palindrome (s[i] == s[i]), length 2 is one when its two letters match.", CodeSnippet: "if length <= 2:\n    is_pal[i][j] = s[i] == s[j]", LineStart: 10, LineEnd: 12},
				{Title: "Choose the evaluation order", Explanation: "The inside substring is 2 shorter, so fill all substrings of length 1, then 2, then 3, and so on. For 'aaa': lengths 1 give 3 True cells, length 2 gives 2, length 3 gives 1.", CodeSnippet: "for length in range(1, n + 1):\n    for i in range(n - length + 1):\n        j = i + length - 1", LineStart: 6, LineEnd: 9},
				{Title: "Locate the answer", Explanation: "Every True cell is one palindromic substring at one position, so count them as they are filled.", CodeSnippet: "if is_pal[i][j]:\n    count += 1", LineStart: 16, LineEnd: 19},
				{Title: "Optimise space", Explanation: "Each cell only reads the cell one step further in on the same 'centre'. Following one centre outwards with two pointers (the 'Expand around centres' alternative) visits the same cells in the same dependency order with O(1) extra space. That version is no longer a DP; it is the DP's table thrown away.", CodeSnippet: "is_pal[i][j] = s[i] == s[j] and is_pal[i + 1][j - 1]", LineStart: 15, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{
					Name: "Check every substring (brute force)",
					Code: `def countSubstrings(s):
    count = 0
    for i in range(len(s)):
        for j in range(i + 1, len(s) + 1):
            piece = s[i:j]
            if piece == piece[::-1]:
                count += 1
    return count`,
					TimeComplexity:  "O(n^3)",
					SpaceComplexity: "O(n)",
					Explanation:     "Compares each substring with its reverse. Obviously correct, which makes it a good checker, but each check redoes the work of the shorter checks inside it.",
				},
				{
					Name: "Expand around centres (optimise space; not DP)",
					Code: `def countSubstrings(s):
    count = 0
    def expand(left, right):
        found = 0
        while left >= 0 and right < len(s) and s[left] == s[right]:
            found += 1                  # s[left..right] is a palindrome
            left -= 1
            right += 1
        return found
    for centre in range(len(s)):
        count += expand(centre, centre)        # odd lengths
        count += expand(centre, centre + 1)    # even lengths
    return count`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Every palindrome has a centre: a letter (odd length) or the gap between two letters (even length), 2n - 1 centres in all. Grow two pointers outwards from each centre while the letters match; each step outwards is exactly the DP recurrence is_pal[i][j] = s[i] == s[j] and is_pal[i+1][j-1], with is_pal[i+1][j-1] known to be True because we just came from it. The first mismatch ends that centre, because every longer string around it contains the mismatch. Same time as the table, O(1) space, and it is the usual interview answer, but note it stores no subproblem answers, so strictly it is a two-pointer technique (Chapter 3), not dynamic programming.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can't the table be filled with the plain double loop 'for i in range(n): for j in range(i, n)'?", Answer: "is_pal[i][j] reads is_pal[i+1][j-1], which belongs to row i+1. With i increasing, row i+1 has not been filled yet, so the recurrence reads a default False and misses palindromes like 'aba'. Filling by increasing length (or with i decreasing from n-1) guarantees the inside substring is finished first: the evaluation order must follow the dependencies."},
			{Prompt: "Expand-around-centres stops at the first mismatch. Why can't a longer palindrome exist around the same centre after that?", Answer: "Any longer substring around the same centre contains the mismatched pair s[left], s[right] at mirrored positions. A palindrome needs every mirrored pair equal, so all of them fail. This is the recurrence used in reverse: if is_pal[i][j] is False, every is_pal[i-k][j+k] is False too."},
		},
	},
	{
		ID:              "longest-palindromic-substring",
		Number:          5,
		Order:           11,
		Title:           "Longest Palindromic Substring",
		Slug:            "longest-palindromic-substring",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"String", "Dynamic Programming", "Two Pointers"},
		RelatedChapters: []int{15, 16, 3},
		Description: `Given a string **s**, return its longest substring that is a palindrome (reads the same forwards and backwards). A substring is a contiguous, non-empty run of characters.

If several palindromic substrings share the maximum length, return any one of them: for "babad" both "bab" and "aba" are accepted.`,
		Constraints: []string{
			"1 <= s.length <= 1000",
			"s consists of only digits and English letters",
		},
		Examples: []Example{
			{Input: `s = "babad"`, Output: `"bab"`, Explanation: `"aba" is also a correct answer; both have length 3, and no palindrome of length 4 or 5 exists.`},
			{Input: `s = "cbbd"`, Output: `"bb"`, Explanation: "An even-length palindrome."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "babad"}, Expected: "bab", Note: "example 1: \"aba\" is accepted too"},
			{Input: map[string]interface{}{"s": "cbbd"}, Expected: "bb", Note: "example 2: even length"},
			{Input: map[string]interface{}{"s": "a"}, Expected: "a", Hidden: true, Note: "a single character"},
			{Input: map[string]interface{}{"s": "ac"}, Expected: "a", Hidden: true, Note: "no repeated letter: any single letter is accepted"},
			{Input: map[string]interface{}{"s": "racecar"}, Expected: "racecar", Hidden: true, Note: "the whole string"},
			{Input: map[string]interface{}{"s": "forgeeksskeegfor"}, Expected: "geeksskeeg", Hidden: true, Note: "long even-length palindrome in the middle"},
			{Input: map[string]interface{}{"s": "abacdfgdcaba"}, Expected: "aba", Hidden: true, Note: "s reversed contains 'abacd', but that is not a palindrome"},
			{Input: map[string]interface{}{"s": "aaaa"}, Expected: "aaaa", Hidden: true, Note: "all letters equal, even length"},
			{Input: map[string]interface{}{"s": "abb"}, Expected: "bb", Hidden: true, Note: "the palindrome is at the end"},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(n^2)",
		StarterCode:     "def longestPalindrome(s):\n    # Write your solution here\n    pass",
		Harness: Harness{
			Validator: `def check(raw, args, result, expected):
    s = raw["s"]
    if not isinstance(result, str):
        return False, "return the palindrome itself as a string, not its length or indices"
    if len(result) == 0 or result not in s:
        return False, repr(result) + " is not a non-empty substring of s"
    if result != result[::-1]:
        return False, repr(result) + " is not a palindrome"
    if len(result) != len(expected):
        return False, repr(result) + " has length " + str(len(result)) + ", but the longest palindrome has length " + str(len(expected))
    return True, ""`,
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You already know how to decide, for every substring, whether it is a palindrome (problem 10). What extra information would you need to keep while doing that, to return the longest one?"},
			{Level: 2, Type: "insight", Content: "Fill the same table is_pal[i][j] = s[i] == s[j] and is_pal[i+1][j-1], by increasing length. Whenever a cell is True and longer than the best so far, remember its start and length. Only the answer step of the recipe changes (Chapter 15).", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. Define the state: is_pal[i][j] = True if s[i..j] is a palindrome.\n2. Write the recurrence: s[i] == s[j] and (length <= 2 or is_pal[i+1][j-1]).\n3. Set the base cases: length 1 always, length 2 when both letters match (folded into the recurrence by the length <= 2 test).\n4. Choose the evaluation order: by increasing length.\n5. Locate the answer: track the start and length of the longest True cell; return that slice of s.\n6. Optimise space: expanding around the 2n - 1 centres finds the same palindromes with O(1) extra space."},
			{Level: 4, Type: "skeleton", Content: "def longestPalindrome(s):\n    n = len(s)\n    is_pal = [[False] * n for _ in range(n)]\n    best_start, best_len = 0, 1\n    for length in range(1, n + 1):\n        for i in range(n - length + 1):\n            j = i + length - 1\n            if s[i] == s[j] and (____):\n                is_pal[i][j] = True\n                if ____:\n                    best_start, best_len = i, length\n    return ____"},
		},
		Solution: Solution{
			Code: `def longestPalindrome(s):
    n = len(s)
    # State: is_pal[i][j] = True if s[i..j] (inclusive) is a palindrome
    is_pal = [[False] * n for _ in range(n)]
    best_start, best_len = 0, 1
    # Evaluation order: shorter substrings first
    for length in range(1, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            # Base cases (length 1 or 2) and recurrence (outer pair + inside)
            if s[i] == s[j] and (length <= 2 or is_pal[i + 1][j - 1]):
                is_pal[i][j] = True
                # Answer: remember the longest True cell seen so far
                if length > best_len:
                    best_start, best_len = i, length
    return s[best_start:best_start + best_len]`,
			Explanation: "This is the same table as Palindromic Substrings (problem 10); only step 5 of the recipe changes: instead of counting True cells we keep the longest one and return the substring itself. Being explicit, as in problem 10: the best-known solution, expand around centres, is not DP; it is this table with the space optimised away, and it is shown as an alternative.",
			BruteForce:  "Check all n(n+1)/2 substrings against their reverse and keep the longest palindrome: O(n^3), about 1.7 x 10^8 character comparisons for n = 1000. Each check repeats the checks of the substrings nested inside it.",
			KeyInsight:  "Whether s[i..j] is a palindrome is decided in O(1) from s[i], s[j] and the answer for s[i+1..j-1]. With every answer stored, all n^2 substrings are classified in O(n^2), and the longest True one is the answer.",
			WhyItWorks:  "Subproblem: is_pal[i][j] says whether s[i..j] is a palindrome. Recurrence: for length >= 3, s[i..j] is a palindrome iff s[i] == s[j] and s[i+1..j-1] is one; for length 1 or 2 the matching end letters are enough, which the condition 'length <= 2 or is_pal[i+1][j-1]' covers in one line. Filling by increasing length makes the inside final before it is read. Every substring is classified, so the longest True cell is a longest palindromic substring. Because cells are visited by increasing length and we only replace the best on a strictly longer length, the first longest palindrome (leftmost) is returned; the judge accepts any palindrome of the maximum length.",
			CommonMistakes: []string{
				"Reversing s and looking for the longest common substring of s and its reverse. For 'abacdfgdcaba' that finds 'abacd' (it appears reversed at the other end), which is not a palindrome. A common substring must also line up at mirrored positions.",
				"Handling only odd-length palindromes. 'cbbd' has answer 'bb', which has no middle letter; the length-2 base case (or the even centre in expansion) is needed.",
				"Returning the length instead of the substring, or slicing s[best_start:best_len] instead of s[best_start:best_start + best_len].",
			},
			TimeComplexity:  "O(n^2)",
			SpaceComplexity: "O(n^2)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "Exactly as in Palindromic Substrings: one yes/no per substring s[i..j]. Two extra variables hold the best palindrome found so far, starting with the first letter (length 1 always exists).", CodeSnippet: "is_pal = [[False] * n for _ in range(n)]\nbest_start, best_len = 0, 1", LineStart: 3, LineEnd: 5},
				{Title: "Write the recurrence", Explanation: "The end letters must match, and for length 3 or more the inside must be a palindrome. The 'length <= 2' test folds the two base cases into the same line.", CodeSnippet: "if s[i] == s[j] and (length <= 2 or is_pal[i + 1][j - 1]):\n    is_pal[i][j] = True", LineStart: 10, LineEnd: 12},
				{Title: "Choose the evaluation order", Explanation: "By increasing length, so is_pal[i+1][j-1] (two shorter) is always ready. For 'cbbd': length 1 marks all four letters, length 2 marks 'bb' at i = 1, lengths 3 and 4 find nothing.", CodeSnippet: "for length in range(1, n + 1):\n    for i in range(n - length + 1):\n        j = i + length - 1", LineStart: 6, LineEnd: 9},
				{Title: "Locate the answer", Explanation: "Keep the start and length of the longest True cell, and slice it out at the end. The slice end is best_start + best_len, not best_len.", CodeSnippet: "if length > best_len:\n    best_start, best_len = i, length\nreturn s[best_start:best_start + best_len]", LineStart: 13, LineEnd: 16},
				{Title: "Optimise space", Explanation: "The table is n^2 booleans, which the 'Expand around centres' alternative replaces with two pointers per centre: same O(n^2) time, O(1) extra space, and no longer a DP.", CodeSnippet: "is_pal = [[False] * n for _ in range(n)]", LineStart: 4, LineEnd: 4},
			},
			Alternatives: []AltSolution{
				{
					Name: "Check every substring (brute force)",
					Code: `def longestPalindrome(s):
    best = s[0]
    for i in range(len(s)):
        for j in range(i + len(best) + 1, len(s) + 1):   # only longer candidates
            piece = s[i:j]
            if piece == piece[::-1]:
                best = piece
    return best`,
					TimeComplexity:  "O(n^3)",
					SpaceComplexity: "O(n)",
					Explanation:     "Tries every substring longer than the current best and compares it with its reverse. Clearly correct and a good checker, but O(n^3) in the worst case.",
				},
				{
					Name: "Expand around centres (optimise space; not DP)",
					Code: `def longestPalindrome(s):
    def expand(left, right):
        # grow while the ends match; return the last palindrome's bounds
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
        return left + 1, right          # s[left+1:right] is the palindrome
    best_lo, best_hi = 0, 1
    for centre in range(len(s)):
        for lo, hi in (expand(centre, centre), expand(centre, centre + 1)):
            if hi - lo > best_hi - best_lo:
                best_lo, best_hi = lo, hi
    return s[best_lo:best_hi]`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "From each of the 2n - 1 centres (each letter for odd lengths, each gap for even lengths), move two pointers outwards while the letters match. Each step outwards applies the DP recurrence with the inside already known to be a palindrome, so no table is needed. This is the usual interview solution; it is a two-pointer technique (Chapter 3) derived from the DP, not a DP itself, since it stores no subproblem answers. (Manacher's algorithm reaches O(n) but is rarely expected.)",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the judge use a validator here instead of comparing with one expected string?", Answer: "The problem allows any longest palindrome. For 'babad', the DP returns 'bab' (the first one found at length 3) while an expansion that prefers the last centre might return 'aba'; both are correct. The validator checks that the answer is a substring of s, is a palindrome, and has the maximum length, instead of insisting on one particular string."},
			{Prompt: "Palindromic Substrings and this problem share the same table. Which steps of the DP recipe are identical, and which one differs?", Answer: "The state (is_pal[i][j]), recurrence (end letters match and inside is a palindrome), base cases (length 1 and 2), evaluation order (by increasing length) and space optimisation (expand around centres) are all identical. Only 'Locate the answer' differs: count every True cell there, keep the longest True cell here. Recognising that is the point of naming the steps."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, DP1DProblems...)
}
