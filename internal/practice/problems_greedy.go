package practice

// GreedyProblems contains all greedy category problems, in ladder order.
//
// Subgoal labels shared with chapter 17: "Choose the greedy rule",
// "Order the input", "Track the greedy state", "Commit or skip",
// "Check for failure", "Return the answer", "Justify the choice".
var GreedyProblems = []*Problem{
	{
		ID:              "jump-game",
		Number:          55,
		Order:           1,
		Title:           "Jump Game",
		Slug:            "jump-game",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Greedy", "Dynamic Programming"},
		RelatedChapters: []int{17, 15},
		Description: `You stand on index 0 of an integer array nums. The value nums[i] is the **longest** jump you may make from index i: from i you can land on any index from i + 1 up to i + nums[i] (a value of 0 means you are stuck there).

Return true if some sequence of jumps lands you on the last index, and false otherwise. An array of length 1 counts as already at the last index, so the answer is true.`,
		Constraints: []string{
			"1 <= nums.length <= 10^4",
			"0 <= nums[i] <= 10^5",
		},
		Examples: []Example{
			{Input: "nums = [2,3,1,1,4]", Output: "true", Explanation: "From index 0 jump 1 step to index 1 (value 3), then 3 steps to index 4, the last index."},
			{Input: "nums = [3,2,1,0,4]", Output: "false", Explanation: "Every route ends on index 3, whose value is 0. Indices 0, 1 and 2 can reach at most index 3, so index 4 is never reached."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 3, 1, 1, 4}}, Expected: true, Note: "example 1: reachable"},
			{Input: map[string]interface{}{"nums": []int{3, 2, 1, 0, 4}}, Expected: false, Note: "example 2: every path is trapped on a zero"},
			{Input: map[string]interface{}{"nums": []int{0}}, Expected: true, Hidden: true, Note: "single element: already on the last index"},
			{Input: map[string]interface{}{"nums": []int{0, 1}}, Expected: false, Hidden: true, Note: "stuck at the very start"},
			{Input: map[string]interface{}{"nums": []int{2, 0, 1}}, Expected: true, Hidden: true, Note: "a zero you can jump over is harmless"},
			{Input: map[string]interface{}{"nums": []int{1, 0, 1, 0}}, Expected: false, Hidden: true, Note: "a later big value can't help if you never reach it"},
			{Input: map[string]interface{}{"nums": []int{4, 0, 0, 0, 0, 1}}, Expected: false, Hidden: true, Note: "reach stops exactly one short of the end"},
			{Input: map[string]interface{}{"nums": []int{3, 0, 0, 1, 0}}, Expected: true, Hidden: true, Note: "must land exactly on the one useful index"},
			{Input: map[string]interface{}{"nums": []int{2, 3, 0, 0, 0}}, Expected: true, Hidden: true, Note: "always taking the full jump gets stuck; a shorter jump works"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def canJump(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Instead of asking which jump to take from each index, what single number about the indices you have already passed would tell you whether index i can be reached at all?"},
			{Level: 2, Type: "insight", Content: "The reachable indices always form one unbroken block [0, farthest]: if you can land on index k from j, you can land on anything between j and k with a shorter jump. So keep only farthest, the furthest index any reachable index can jump to. If you ever stand on an index greater than farthest, you can't be there, and nothing after it is reachable either. This is greedy (always keep the best reach) rather than DP over every jump; see chapter 17.", ChapterRef: 17},
			{Level: 3, Type: "plan", Content: "1. Track the greedy state: farthest = 0.\n2. Order the input: walk the indices left to right.\n3. Check for failure: if the current index is beyond farthest, return false.\n4. Commit or skip: extend farthest with i + nums[i] if that is further.\n5. Return the answer: finishing the loop means every index, including the last, was reachable.\n6. Justify the choice: why is the single number farthest enough, without remembering which jumps you took?"},
			{Level: 4, Type: "skeleton", Content: "def canJump(nums):\n    farthest = 0\n    for i, jump in enumerate(nums):\n        if ____:\n            return False\n        farthest = ____\n    return True"},
		},
		Solution: Solution{
			Code: `def canJump(nums):
    farthest = 0  # furthest index reachable so far
    for i, jump in enumerate(nums):
        if i > farthest:
            return False  # i can't be reached, so nothing after it can
        farthest = max(farthest, i + jump)
    return True`,
			Explanation: "This is the gentlest greedy in the ladder: one left-to-right scan that keeps a single summary number, farthest, instead of exploring every route. It is the 1D DP reachability question from chapter 15 with the table collapsed to one value, which is possible because reachable indices form a prefix of the array.",
			BruteForce:  "Mark index 0 reachable, then for every reachable index mark each of the nums[i] indices it can land on. That fills a reachable[] table correctly but costs O(n * max jump), up to 10^4 * 10^5 = 10^9 marks. The waste: it marks index 7 as reachable from index 2, again from index 3, again from index 4, though one mark was enough.",
			KeyInsight:  "Reachable indices are always a contiguous prefix [0, farthest]. So the whole table can be replaced by one number, and checking i > farthest tells you instantly whether index i is reachable.",
			WhyItWorks: `Invariant (greedy stays ahead): when the loop reaches index i, every index in [0, farthest] is reachable and no reachable index lies beyond farthest.
Why it holds: if some reachable j jumps to k = j + nums[j], then every index between j and k is reachable too, because a shorter jump is allowed. So the reachable set has no holes, and the largest j + nums[j] over the reachable j seen so far is exactly its right edge.
Consequence: if i > farthest, index i is not reachable. Any index after it could only be reached from an index before it, and those all stop at farthest < i, so the answer is false. If the loop finishes, the last index was <= farthest, so it is reachable. No route can ever beat farthest, which is why we never need to remember which jumps were taken.`,
			CommonMistakes: []string{
				"Returning false as soon as you see a 0. A zero only traps you if farthest has not already passed it, e.g. [2,0,1] is true. Compare i with farthest instead of looking at the value.",
				"Following one path by always taking the full jump. On [2,3,0,0,0] the full jump from index 0 lands on a 0 and gets stuck, while a 1-step jump reaches the 3 and then the end. Tracking farthest over all reachable indices means you never pick a single path.",
				"Updating farthest before the reachability check. Then an unreachable index can extend farthest, e.g. [0,1] would wrongly return true.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track the greedy state", Explanation: "farthest is the right edge of the block of reachable indices. Index 0 is where you start, so the block begins as [0, 0].", CodeSnippet: "farthest = 0  # furthest index reachable so far", LineStart: 2, LineEnd: 2},
				{Title: "Check for failure", Explanation: "If index i is beyond the block, no earlier index can jump to it, and therefore nothing to its right is reachable. On [3,2,1,0,4], farthest is 3 when i = 4, so we return false.", CodeSnippet: "if i > farthest:\n    return False", LineStart: 4, LineEnd: 5},
				{Title: "Commit or skip", Explanation: "From a reachable index we can extend the block to i + nums[i]. Keeping the max means we commit only to improvements; a short jump never shrinks the reach. On [2,3,1,1,4]: farthest goes 2, 4, 4, ...", CodeSnippet: "farthest = max(farthest, i + jump)", LineStart: 6, LineEnd: 6},
				{Title: "Return the answer", Explanation: "Finishing the loop means every index, including the last one, passed the reachability check.", CodeSnippet: "return True", LineStart: 7, LineEnd: 7},
			},
			Alternatives: []AltSolution{
				{
					Name: "Reachability table (brute force)",
					Code: `def canJump(nums):
    n = len(nums)
    reachable = [False] * n
    reachable[0] = True
    for i in range(n):
        if not reachable[i]:
            continue
        for step in range(1, nums[i] + 1):
            if i + step < n:
                reachable[i + step] = True
    return reachable[n - 1]`,
					TimeComplexity:  "O(n * k), where k is the largest value in nums",
					SpaceComplexity: "O(n)",
					Explanation:     "Marks every index each reachable index can land on, like a 1D DP table. Correct, but it re-marks the same indices many times; the greedy version notices that the marked indices always form a prefix and keeps only its end.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why are the reachable indices always a contiguous block starting at 0, and why does the solution depend on that?", Answer: "A jump of length up to nums[j] lets you land on any index from j+1 to j+nums[j], so whenever k is reachable from j, everything between j and k is too. With no holes, the set of reachable indices is fully described by its right end, farthest, and checking i > farthest is an exact test of reachability."},
			{Prompt: "What changes if each nums[i] were an exact jump length (you must jump exactly nums[i])?", Answer: "The reachable set can now have holes: from index 0 with nums[0] = 3 you reach index 3 but not 1 or 2. One number no longer summarises the past, so the greedy breaks and you need a reachability table or a BFS/DFS over indices, which is the DP or graph approach."},
		},
	},
	{
		ID:              "maximum-subarray",
		Number:          53,
		Order:           2,
		Title:           "Maximum Subarray",
		Slug:            "maximum-subarray",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Greedy", "Dynamic Programming", "Divide and Conquer"},
		RelatedChapters: []int{17, 15},
		Description: `Given an integer array nums, choose a **subarray**: a contiguous, non-empty run of elements nums[i..j]. Return the largest sum any subarray can have.

The subarray must contain at least one element, so if every number is negative the answer is the largest (least negative) single element.`,
		Constraints: []string{
			"1 <= nums.length <= 10^5",
			"-10^4 <= nums[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "nums = [-2,1,-3,4,-1,2,1,-5,4]", Output: "6", Explanation: "The subarray [4,-1,2,1] sums to 6. Including the -1 is worth it because the 2 and 1 after it more than pay it back."},
			{Input: "nums = [5,4,-1,7,8]", Output: "23", Explanation: "The whole array is the best subarray."},
			{Input: "nums = [-3,-1,-2]", Output: "-1", Explanation: "All numbers are negative, so the best non-empty subarray is the single element -1."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, Expected: 6, Note: "example 1: best run crosses a negative number"},
			{Input: map[string]interface{}{"nums": []int{5, 4, -1, 7, 8}}, Expected: 23, Note: "example 2: whole array"},
			{Input: map[string]interface{}{"nums": []int{-3, -1, -2}}, Expected: -1, Note: "all negative: catches best = 0 initialisation"},
			{Input: map[string]interface{}{"nums": []int{1}}, Expected: 1, Hidden: true, Note: "single element"},
			{Input: map[string]interface{}{"nums": []int{-7}}, Expected: -7, Hidden: true, Note: "single negative element"},
			{Input: map[string]interface{}{"nums": []int{2, -1, 2}}, Expected: 3, Hidden: true, Note: "a small dip worth crossing"},
			{Input: map[string]interface{}{"nums": []int{-1, 3, -5, 4}}, Expected: 4, Hidden: true, Note: "a dip too deep to cross: restart"},
			{Input: map[string]interface{}{"nums": []int{0, -1, 0}}, Expected: 0, Hidden: true, Note: "zeros and a negative"},
			{Input: map[string]interface{}{"nums": []int{8, -19, 5, -4, 20}}, Expected: 21, Hidden: true, Note: "the best run starts after a big drop"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def maxSubArray(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose you are building a run that will end at the current number. If the run you have carried so far adds up to a negative amount, would you ever want to keep it attached?"},
			{Level: 2, Type: "insight", Content: "A prefix with a negative sum only drags down whatever comes after it, so drop it and restart at the current number. Keep current = the best sum of a run ending here, and best = the largest current ever seen. This is Kadane's algorithm: greedy in that it discards a negative prefix the moment it appears (chapter 17), and also the 1D DP 'best ending at i' from chapter 15.", ChapterRef: 17},
			{Level: 3, Type: "plan", Content: "1. Track the greedy state: current (best sum of a run ending at the previous index) and best (best answer so far, start it at nums[0], not 0).\n2. Order the input: scan left to right.\n3. Commit or skip: if current is negative, drop it (reset to 0) before adding the new number.\n4. Add the number and update best.\n5. Return the answer: best.\n6. Justify the choice: why can a negative prefix never be part of an optimal run that extends past it?"},
			{Level: 4, Type: "skeleton", Content: "def maxSubArray(nums):\n    best = nums[0]\n    current = 0\n    for num in nums:\n        if ____:\n            current = 0\n        current += num\n        best = ____\n    return best"},
		},
		Solution: Solution{
			Code: `def maxSubArray(nums):
    best = nums[0]  # a subarray can't be empty
    current = 0     # best sum of a run ending at the previous index
    for num in nums:
        if current < 0:
            current = 0  # a negative prefix only hurts: drop it
        current += num
        best = max(best, current)
    return best`,
			Explanation: "Kadane's algorithm. It builds on Jump Game: again one scan and a tiny greedy state, but now the state is a running sum, and the greedy decision is whether to keep the run you are carrying or throw it away. It is the same algorithm as the 1D DP 'best subarray ending at i' from chapter 15, with the table collapsed to one variable.",
			BruteForce:  "Try every start i and every end j, adding as you extend j: O(n^2), about 5 * 10^9 additions for n = 10^5. The waste: when the run nums[i..j] already has a negative sum, every subarray that starts at i and goes past j is beaten by the same subarray started at j+1, yet the brute force still checks them all.",
			KeyInsight:  "The best run ending at index i is either nums[i] alone or nums[i] glued onto the best run ending at i-1, and you glue only when that earlier run has a non-negative sum. One number, current, carries everything you need.",
			WhyItWorks: `Invariant: after processing index i, current is the largest sum of any subarray that ends exactly at i.
Exchange argument for the reset: take any subarray nums[s..i] whose prefix nums[s..k] (k < i) has a negative sum. Swap it for nums[k+1..i]: you removed a negative amount, so the sum went up. So an optimal run never starts with a negative prefix, and when the best run ending at i-1 is negative, the best run ending at i is nums[i] alone. When it is non-negative, extending it is at least as good as starting fresh.
Every subarray ends at some index, so best = max over i of current is the true maximum. best starts at nums[0] rather than 0 because the subarray must be non-empty: on [-3,-1,-2] the answer is -1, not 0.`,
			CommonMistakes: []string{
				"Initialising best = 0. On an all-negative array such as [-3,-1,-2] it returns 0, the sum of an empty subarray, which isn't allowed. Start best at nums[0] (or -infinity).",
				"Restarting the run whenever you meet a negative number instead of when the running sum turns negative. On [2,-1,2] the best run crosses the -1 and sums to 3; restarting at the -1 finds only 2.",
				"Updating best only when you reset. The best run can end anywhere, so compare after every addition.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track the greedy state", Explanation: "best must start at a real subarray sum (nums[0]) so all-negative input works; current is the best sum of a run ending just before the next number.", CodeSnippet: "best = nums[0]  # a subarray can't be empty\ncurrent = 0     # best sum of a run ending at the previous index", LineStart: 2, LineEnd: 3},
				{Title: "Commit or skip", Explanation: "If the carried run is negative it can only lower any sum it is attached to, so drop it. On [-2,1,-3,4,...]: current is -2 after the first number, so it resets before adding 1.", CodeSnippet: "if current < 0:\n    current = 0  # a negative prefix only hurts: drop it", LineStart: 5, LineEnd: 6},
				{Title: "Track the greedy state", Explanation: "Extend the run with the current number and record the best seen. On example 1, current runs -2, 1, -2, 4, 3, 5, 6, 1, 5 and best ends at 6.", CodeSnippet: "current += num\nbest = max(best, current)", LineStart: 7, LineEnd: 8},
				{Title: "Return the answer", Explanation: "best is the maximum over every possible ending index, which covers every subarray.", CodeSnippet: "return best", LineStart: 9, LineEnd: 9},
			},
			Alternatives: []AltSolution{
				{
					Name: "Every start and end (brute force)",
					Code: `def maxSubArray(nums):
    best = nums[0]
    for start in range(len(nums)):
        total = 0
        for end in range(start, len(nums)):
            total += nums[end]
            best = max(best, total)
    return best`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Checks every subarray with a running total per start. Fine for small inputs; too slow at n = 10^5.",
				},
				{
					Name:            "Divide and conquer",
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(log n)",
					Explanation:     "Split the array in half; the best subarray is in the left half, the right half, or crosses the middle (best suffix of the left plus best prefix of the right). A good recursion exercise (chapter 8), but slower than Kadane.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to throw away the running sum as soon as it goes negative, even though later numbers might be large?", Answer: "Any run that keeps that negative prefix has a larger sum if you remove the prefix (exchange argument): the later large numbers are still available to the run that starts fresh. So the prefix can never be part of an optimal run that continues past it."},
			{Prompt: "How would you return the subarray itself (its start and end) instead of only the sum?", Answer: "Remember the index where the current run started: set run_start = i whenever you reset. Whenever current beats best, save best_start = run_start and best_end = i. The scan is unchanged, still O(n) time and O(1) space."},
		},
	},
	{
		ID:              "jump-game-ii",
		Number:          45,
		Order:           3,
		Title:           "Jump Game II",
		Slug:            "jump-game-ii",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Greedy", "Dynamic Programming"},
		RelatedChapters: []int{17, 13, 15},
		Description: `You start on index 0 of an integer array nums of length n. From index i you may jump forward to any index from i + 1 up to i + nums[i].

Return the **minimum number of jumps** needed to land on the last index (n - 1). The input always allows reaching the last index. If n = 1 you are already there and the answer is 0.`,
		Constraints: []string{
			"1 <= nums.length <= 10^4",
			"0 <= nums[i] <= 1000",
			"The last index is always reachable from index 0",
		},
		Examples: []Example{
			{Input: "nums = [2,3,1,1,4]", Output: "2", Explanation: "Jump 1 step from index 0 to index 1, then 3 steps to index 4. One jump can reach only index 1 or 2, neither of which is the end, so 2 is the minimum."},
			{Input: "nums = [2,3,0,1,4]", Output: "2", Explanation: "Again index 0 -> index 1 -> index 4."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 3, 1, 1, 4}}, Expected: 2, Note: "example 1"},
			{Input: map[string]interface{}{"nums": []int{2, 3, 0, 1, 4}}, Expected: 2, Note: "example 2: a zero inside a level"},
			{Input: map[string]interface{}{"nums": []int{0}}, Expected: 0, Hidden: true, Note: "single element: no jumps needed (catches looping to the last index)"},
			{Input: map[string]interface{}{"nums": []int{1, 2}}, Expected: 1, Hidden: true, Note: "one jump"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 1}}, Expected: 3, Hidden: true, Note: "every jump is forced"},
			{Input: map[string]interface{}{"nums": []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}}, Expected: 1, Hidden: true, Note: "first jump reaches the end"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 1, 1, 1}}, Expected: 3, Hidden: true, Note: "best path is not 'always take the biggest value'"},
			{Input: map[string]interface{}{"nums": []int{3, 1, 1, 1, 1, 1, 1}}, Expected: 4, Hidden: true, Note: "long tail after a big first jump"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def jump(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Which indices can you reach with exactly 1 jump? With at most 2 jumps? What shape do those sets have, and how could you get the 2-jump set from the 1-jump set?"},
			{Level: 2, Type: "insight", Content: "It's BFS by levels (chapter 13) on an array: level k is the block of indices whose minimum jump count is k, and like in Jump Game every level is a contiguous range. The next level ends at the farthest point any index in the current level can reach. Count how many times you have to move to a new level before the last index is inside one. Chapter 17 explains why taking the farthest reach is safe.", ChapterRef: 17},
			{Level: 3, Type: "plan", Content: "1. Track the greedy state: jumps = 0, current_end = 0 (end of the current level), farthest = 0 (end of the next level so far).\n2. Order the input: scan indices 0 to n-2. You never jump from the last index.\n3. Commit or skip: extend farthest with i + nums[i].\n4. When i reaches current_end, the current level is used up: count one more jump and set current_end = farthest.\n5. Return the answer: jumps.\n6. Justify the choice: why can no path with k jumps get further than the greedy's k-th level end?"},
			{Level: 4, Type: "skeleton", Content: "def jump(nums):\n    jumps = 0\n    current_end = 0\n    farthest = 0\n    for i in range(____):\n        farthest = max(farthest, i + nums[i])\n        if ____:\n            jumps += 1\n            current_end = ____\n    return jumps"},
		},
		Solution: Solution{
			Code: `def jump(nums):
    jumps = 0
    current_end = 0  # last index reachable with 'jumps' jumps
    farthest = 0     # last index reachable with one more jump
    for i in range(len(nums) - 1):
        farthest = max(farthest, i + nums[i])
        if i == current_end:
            jumps += 1   # this level is used up: jump to the next
            current_end = farthest
    return jumps`,
			Explanation: "Builds on Jump Game, which tracked one block of reachable indices. Here we track blocks level by level: the indices reachable in 0 jumps, in 1 jump, in 2, and so on, exactly like BFS on a graph (chapter 13) but without a queue, because each level is a contiguous range described by its end.",
			BruteForce:  "A DP table min_jumps[j] = min(min_jumps[i] + 1) over every i that can reach j. That is O(n * max jump), up to 10^4 * 1000 = 10^7 updates. The waste: it computes an exact jump count for every index, while all indices in one contiguous block share the same count, so only the block boundaries matter.",
			KeyInsight:  "Indices that need exactly k jumps form a contiguous range, and the next range ends at the farthest point reachable from the current one. So one scan that tracks current_end and farthest does the whole BFS.",
			WhyItWorks: `Greedy stays ahead: let E_k be current_end after the k-th increment. Claim: E_k is the furthest index reachable with at most k jumps, and every index up to E_k is reachable with at most k jumps.
Base: E_0 = 0. Step: any path with k+1 jumps makes its last jump from some index i <= E_k, so it lands at most at i + nums[i] <= max over i <= E_k of (i + nums[i]) = E_{k+1}. And every index up to E_{k+1} is reachable by a shorter or equal jump from the index that achieves the max. So no path with k jumps beats the greedy's k-th level, and the first k with E_k >= n-1 is the minimum.
The loop stops at n-2 because you never jump from the last index: running it to n-1 counts an extra jump whenever a level ends exactly at the last index (on [0] it would return 1).`,
			CommonMistakes: []string{
				"Looping over all n indices. When i reaches the last index and it equals current_end you add a jump you never make; [0] returns 1 and [2,3,1,1,4] returns 3. Loop to len(nums) - 1.",
				"Always making the full jump nums[i] from where you stand. On [2,3,1,1,4] that goes 0 -> 2 -> 3 -> 4 (3 jumps), but 0 -> 1 -> 4 takes 2. The right measure of a take-off point is its reach i + nums[i], taken over the whole level.",
				"Incrementing jumps every time farthest grows. Jumps count levels, not improvements; only count when you cross the end of the current level.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track the greedy state", Explanation: "Level 0 is just index 0, so current_end = 0. farthest collects the end of the next level as we scan this one.", CodeSnippet: "jumps = 0\ncurrent_end = 0  # last index reachable with 'jumps' jumps\nfarthest = 0     # last index reachable with one more jump", LineStart: 2, LineEnd: 4},
				{Title: "Order the input", Explanation: "Scan left to right up to n-2. Standing on the last index means we are done, so its value is never used.", CodeSnippet: "for i in range(len(nums) - 1):", LineStart: 5, LineEnd: 5},
				{Title: "Commit or skip", Explanation: "Every index in the current level can be the take-off point, so the next level reaches the max of i + nums[i]. On [2,3,1,1,4]: level 1 is indices 1..2, and from them farthest becomes 4.", CodeSnippet: "farthest = max(farthest, i + nums[i])", LineStart: 6, LineEnd: 6},
				{Title: "Commit or skip", Explanation: "At the end of the current level we must jump; the best we can do is the next level's end. On example 1: i=0 sets current_end=2 (jumps=1), i=2 sets current_end=4 (jumps=2).", CodeSnippet: "if i == current_end:\n    jumps += 1   # this level is used up: jump to the next\n    current_end = farthest", LineStart: 7, LineEnd: 9},
				{Title: "Return the answer", Explanation: "jumps is the number of levels we had to cross before the last index was inside one.", CodeSnippet: "return jumps", LineStart: 10, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{
					Name: "Minimum-jumps table (DP)",
					Code: `def jump(nums):
    n = len(nums)
    min_jumps = [float('inf')] * n
    min_jumps[0] = 0
    for i in range(n):
        for step in range(1, nums[i] + 1):
            if i + step < n:
                min_jumps[i + step] = min(min_jumps[i + step], min_jumps[i] + 1)
    return min_jumps[n - 1]`,
					TimeComplexity:  "O(n * k), where k is the largest value in nums",
					SpaceComplexity: "O(n)",
					Explanation:     "The chapter 15 style: min_jumps[j] is the fewest jumps to reach j, relaxed from every index that can reach it. Correct and simple, but does work for every index instead of every level.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is this the same as BFS, and why don't we need a queue?", Answer: "BFS from index 0 visits nodes level by level, where level k holds the nodes at distance k. Here the nodes at distance k always form a contiguous range of indices, so the 'queue' for a level is fully described by its end, current_end. Scanning i up to current_end is popping that level; farthest is the end of the next level being built."},
			{Prompt: "What would you do if the last index were not guaranteed to be reachable?", Answer: "Detect a stall: if after processing i == current_end the new current_end (farthest) is still i, no index in the level can move forward, so return -1. This is the same failure check as Jump Game (index beyond farthest)."},
		},
	},
	{
		ID:              "gas-station",
		Number:          134,
		Order:           4,
		Title:           "Gas Station",
		Slug:            "gas-station",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Greedy"},
		RelatedChapters: []int{17},
		Description: `There are n gas stations arranged in a circle, numbered 0 to n - 1. Station i has gas[i] units of fuel, and driving from station i to the next station (i + 1, or station 0 after station n - 1) uses cost[i] units.

Your car's tank is unlimited and starts empty. You pick a starting station, fill up there, drive to the next station, fill up, and so on. The tank may hit exactly 0 on arrival, but it may never be negative.

Return the index of a starting station from which you can drive all the way around once and return to it, or -1 if no such station exists. If an answer exists it is guaranteed to be unique.`,
		Constraints: []string{
			"n == gas.length == cost.length",
			"1 <= n <= 10^5",
			"0 <= gas[i], cost[i] <= 10^4",
			"If a valid start exists, it is unique",
		},
		Examples: []Example{
			{Input: "gas = [1,2,3,4,5], cost = [3,4,5,1,2]", Output: "3", Explanation: "Start at 3 with 4 gas. Tank after each leg: 4-1=3, 3+5-2=6, 6+1-3=4, 4+2-4=2, 2+3-5=0. It never goes negative, so 3 works. Starting at 0, 1 or 2 runs dry on the first leg, and from 4 the tank is 5-2=3 at station 0, 3+1-3=1 at station 1, then 1+2-4=-1 on the way to station 2."},
			{Input: "gas = [2,3,4], cost = [3,4,3]", Output: "-1", Explanation: "There are 9 units of gas in total but the loop costs 10, so no start can work."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"gas": []int{1, 2, 3, 4, 5}, "cost": []int{3, 4, 5, 1, 2}}, Expected: 3, Note: "example 1"},
			{Input: map[string]interface{}{"gas": []int{2, 3, 4}, "cost": []int{3, 4, 3}}, Expected: -1, Note: "example 2: total gas < total cost"},
			{Input: map[string]interface{}{"gas": []int{5}, "cost": []int{4}}, Expected: 0, Hidden: true, Note: "single station that works"},
			{Input: map[string]interface{}{"gas": []int{4}, "cost": []int{5}}, Expected: -1, Hidden: true, Note: "single station that fails"},
			{Input: map[string]interface{}{"gas": []int{5, 1, 2, 3, 4}, "cost": []int{4, 4, 1, 5, 1}}, Expected: 4, Hidden: true, Note: "answer is the last station; route wraps around"},
			{Input: map[string]interface{}{"gas": []int{3, 1, 1}, "cost": []int{1, 2, 2}}, Expected: 0, Hidden: true, Note: "tank ends at exactly 0: allowed"},
			{Input: map[string]interface{}{"gas": []int{4, 5, 2, 6, 5, 3}, "cost": []int{3, 2, 7, 3, 2, 9}}, Expected: -1, Hidden: true, Note: "totals fall short by one"},
			{Input: map[string]interface{}{"gas": []int{0, 0, 9}, "cost": []int{3, 3, 3}}, Expected: 2, Hidden: true, Note: "two resets before the real start"},
			{Input: map[string]interface{}{"gas": []int{1, 2, 3, 4, 5, 5, 70}, "cost": []int{2, 3, 4, 3, 9, 6, 2}}, Expected: 6, Hidden: true, Note: "one huge station rescues the loop"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def canCompleteCircuit(gas, cost):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Say you start at station s and first run out of gas on the way out of station i. Could starting at any station between s and i do better? Think about how much fuel you had when you passed that station."},
			{Level: 2, Type: "insight", Content: "Lemma: if the tank goes negative on the leg out of station i, no start in [s, i] works. When you passed any station p in that range you arrived with a tank >= 0, so starting fresh at p (tank 0) is no better, and you would still fail at i. So jump the candidate start straight to i + 1 and reset the tank. Separately, a start exists exactly when sum(gas) >= sum(cost). It's the same 'drop a negative running sum' move as Kadane's algorithm; see chapter 17.", ChapterRef: 17},
			{Level: 3, Type: "plan", Content: "1. Check for failure: if total gas < total cost, return -1.\n2. Track the greedy state: start = 0, tank = 0.\n3. Order the input: scan stations 0..n-1 once (no need to wrap).\n4. Commit or skip: add gas[i] - cost[i]; if the tank is negative, every start up to i is ruled out, so set start = i + 1 and tank = 0.\n5. Return the answer: start.\n6. Justify the choice: prove the lemma, and explain why the totals check makes the final start work around the wrap."},
			{Level: 4, Type: "skeleton", Content: "def canCompleteCircuit(gas, cost):\n    if ____:\n        return -1\n    start = 0\n    tank = 0\n    for i in range(len(gas)):\n        tank += gas[i] - cost[i]\n        if tank < 0:\n            start = ____\n            tank = ____\n    return start"},
		},
		Solution: Solution{
			Code: `def canCompleteCircuit(gas, cost):
    if sum(gas) < sum(cost):
        return -1  # not enough fuel in the whole loop
    start = 0
    tank = 0
    for i in range(len(gas)):
        tank += gas[i] - cost[i]
        if tank < 0:
            start = i + 1  # no start in [start, i] can work
            tank = 0
    return start`,
			Explanation: "Builds on Maximum Subarray: a running sum of gas[i] - cost[i] that you reset when it goes negative. In Kadane the reset meant 'drop this prefix'; here it means 'every start up to here is ruled out'. A totals check settles whether any start exists.",
			BruteForce:  "Simulate the full loop from every start: O(n^2), 10^10 steps for n = 10^5. The waste: after a simulation from s fails at i, the brute force tries s+1, s+2, ..., i, each of which is doomed to fail at i or earlier.",
			KeyInsight:  "A failure at i rules out the whole block of starts [s, i] at once, so the candidate start only moves forward and one pass suffices.",
			WhyItWorks: `Let diff[i] = gas[i] - cost[i].
Lemma: if you start at s and the tank first goes negative after station i, then no start p in [s, i] works. Proof: for p in (s, i], the tank on arriving at p (starting from s) was >= 0, because i was the first failure. The tank at i from start p equals (tank at i from s) minus (tank on arrival at p from s), which is <= the tank from s, so it is also negative. So we can skip straight to i + 1.
Why the final start works when total >= 0: the resets happen exactly at the lowest points of the prefix sums P(k) = diff[0] + ... + diff[k-1], so the final start sits right after the minimum prefix sum. Driving from it, the tank at any later station is P(now) - P(min) >= 0, and after wrapping past station n-1 it is total + P(now) - P(min) >= 0 because total >= 0. So the tank never goes negative anywhere on the loop.`,
			CommonMistakes: []string{
				"Skipping the totals check and returning start anyway. On [2,3,4] / [3,4,3] the scan ends with start = 2 but the loop is impossible; you need sum(gas) >= sum(cost).",
				"Setting start = i instead of i + 1 after a failure. Station i itself is in the failed block, since the failing leg is the one leaving station i.",
				"Resetting when tank <= 0 instead of tank < 0. Arriving with exactly 0 is allowed: on gas [3,1,1], cost [1,2,2] the tank ends at 0 from start 0, and the <= version moves start to 3, which isn't even a station.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Check for failure", Explanation: "If the whole loop uses more fuel than exists, no start works. Example 2 has 9 gas and costs 10.", CodeSnippet: "if sum(gas) < sum(cost):\n    return -1  # not enough fuel in the whole loop", LineStart: 2, LineEnd: 3},
				{Title: "Track the greedy state", Explanation: "start is the only candidate not yet ruled out; tank is the fuel left when driving from that candidate.", CodeSnippet: "start = 0\ntank = 0", LineStart: 4, LineEnd: 5},
				{Title: "Commit or skip", Explanation: "Drive one leg. If the tank is negative, the lemma rules out every start from 'start' to i, so move the candidate to i + 1 with an empty tank. Example 1: diffs are -2,-2,-2,3,3; the tank fails at i = 0, 1, 2, then stays at 3 and 6 from start 3.", CodeSnippet: "tank += gas[i] - cost[i]\nif tank < 0:\n    start = i + 1  # no start in [start, i] can work\n    tank = 0", LineStart: 7, LineEnd: 10},
				{Title: "Return the answer", Explanation: "The totals check guarantees the surviving candidate also makes it around the wrap.", CodeSnippet: "return start", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Simulate from every start (brute force)",
					Code: `def canCompleteCircuit(gas, cost):
    n = len(gas)
    for start in range(n):
        tank = 0
        completed = True
        for step in range(n):
            station = (start + step) % n
            tank += gas[station] - cost[station]
            if tank < 0:
                completed = False
                break
        if completed:
            return start
    return -1`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Tries each start and drives the full loop. Clear and correct, but repeats the doomed starts that the lemma skips in one move.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does a failure at station i rule out starting at a station p between start and i, and not just at start itself?", Answer: "Driving from start, you arrived at p with a tank >= 0 (i was the first failure). Starting at p means arriving there with 0 instead, which is no more fuel, and the legs from p to i are the same, so the tank at i is no higher and still negative."},
			{Prompt: "Why don't we need to simulate the wrap-around from the final start?", Answer: "The final start is right after the lowest prefix sum. Every later tank value is the prefix sum minus that minimum, so it is >= 0 before the wrap; after the wrap the tank is total plus a prefix sum minus the minimum, and total >= 0 keeps it >= 0. The totals check is exactly what makes the wrap safe."},
		},
	},
	{
		ID:              "hand-of-straights",
		Number:          846,
		Order:           5,
		Title:           "Hand of Straights",
		Slug:            "hand-of-straights",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Hash Table", "Greedy", "Sorting"},
		RelatedChapters: []int{17, 2, 11},
		Description: `You hold some cards, where hand[i] is the number written on the i-th card (numbers can repeat). You want to split all the cards into groups so that:

- every group has exactly groupSize cards, and
- the numbers in each group are consecutive, such as 4, 5, 6.

Every card must be used in exactly one group. Return true if such a split exists, and false otherwise.`,
		Constraints: []string{
			"1 <= hand.length <= 10^4",
			"0 <= hand[i] <= 10^9",
			"1 <= groupSize <= hand.length",
		},
		Examples: []Example{
			{Input: "hand = [1,2,3,6,2,3,4,7,8], groupSize = 3", Output: "true", Explanation: "The smallest card, 1, must start a group, so [1,2,3] is forced. The smallest left is 2, forcing [2,3,4]. Then 6 forces [6,7,8], and all cards are used."},
			{Input: "hand = [1,2,3,4,5], groupSize = 4", Output: "false", Explanation: "5 cards can't be split into groups of 4."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"hand": []int{1, 2, 3, 6, 2, 3, 4, 7, 8}, "groupSize": 3}, Expected: true, Note: "example 1"},
			{Input: map[string]interface{}{"hand": []int{1, 2, 3, 4, 5}, "groupSize": 4}, Expected: false, Note: "example 2: length not divisible by groupSize"},
			{Input: map[string]interface{}{"hand": []int{1, 2, 3, 4}, "groupSize": 3}, Expected: false, Hidden: true, Note: "length not divisible, even though 1,2,3 is a valid group"},
			{Input: map[string]interface{}{"hand": []int{1, 1, 2, 2, 3, 3}, "groupSize": 3}, Expected: true, Hidden: true, Note: "duplicates: two copies of each group"},
			{Input: map[string]interface{}{"hand": []int{1, 1, 2, 2, 3, 4}, "groupSize": 3}, Expected: false, Hidden: true, Note: "two groups start at 1 but only one 3"},
			{Input: map[string]interface{}{"hand": []int{5, 3, 9}, "groupSize": 1}, Expected: true, Hidden: true, Note: "groupSize 1: always possible"},
			{Input: map[string]interface{}{"hand": []int{8, 10, 12}, "groupSize": 3}, Expected: false, Hidden: true, Note: "gap inside the only group"},
			{Input: map[string]interface{}{"hand": []int{2, 1}, "groupSize": 2}, Expected: true, Hidden: true, Note: "unsorted input"},
			{Input: map[string]interface{}{"hand": []int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, "groupSize": 3}, Expected: true, Hidden: true, Note: "overlapping runs share values"},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def isNStraightHand(hand, groupSize):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Look at the smallest card in the hand. Which group could it possibly belong to, and where in that group must it sit?"},
			{Level: 2, Type: "insight", Content: "The smallest remaining card can't be in the middle of a group, because nothing smaller is left to come before it. So its group is forced: c, c+1, ..., c+groupSize-1. When the choice is forced there is nothing to prove wrong; just take it, remove those cards, and repeat. Use a count map (chapter 2) and visit values in sorted order. See chapter 17 on forced choices.", ChapterRef: 17},
			{Level: 3, Type: "plan", Content: "1. Check for failure: if len(hand) is not a multiple of groupSize, return false.\n2. Track the greedy state: count how many copies of each value you hold.\n3. Order the input: visit the distinct values from smallest to largest.\n4. Commit or skip: if value c still has k copies, they all start groups, so c, c+1, ..., c+groupSize-1 each need at least k copies; subtract k from each, or return false if one is short.\n5. Return the answer: true if every value was used up.\n6. Justify the choice: why must the smallest card start its group?"},
			{Level: 4, Type: "skeleton", Content: "from collections import Counter\n\ndef isNStraightHand(hand, groupSize):\n    if ____:\n        return False\n    count = Counter(hand)\n    for card in sorted(count):\n        groups = count[card]\n        if groups == 0:\n            continue\n        for value in range(____):\n            if count[value] < groups:\n                return False\n            count[value] -= ____\n    return True"},
		},
		Solution: Solution{
			Code: `from collections import Counter

def isNStraightHand(hand, groupSize):
    if len(hand) % groupSize != 0:
        return False
    count = Counter(hand)
    for card in sorted(count):
        groups = count[card]  # copies left = groups that must start here
        if groups == 0:
            continue
        for value in range(card, card + groupSize):
            if count[value] < groups:
                return False
            count[value] -= groups
    return True`,
			Explanation: "Builds on the earlier problems by adding sorting as the 'Order the input' step: once values are visited from smallest up, the greedy choice is not just good but forced. A hash map of counts (chapter 2) lets us remove whole groups without searching the hand.",
			BruteForce:  "Try every way to partition the hand into groups (backtracking), which is exponential. A gentler brute force is greedy on a sorted list: take the smallest card and remove c+1, ..., c+groupSize-1 one at a time with list.remove, which costs O(n) per removal, O(n^2) overall. The waste: scanning the list to find each card instead of looking up a count.",
			KeyInsight:  "The smallest remaining card must be the lowest card of its group, so its group is completely determined. Process values in sorted order and subtract whole groups from a count map.",
			WhyItWorks: `Exchange argument, in its strongest form: let c be the smallest remaining value. In any valid split, the group holding a copy of c contains no smaller card (none is left), so c is its lowest card and the group must be exactly c, c+1, ..., c+groupSize-1. Every valid split therefore makes the same choice greedy makes, so committing to it can never lose a solution. Removing those cards leaves a smaller hand with the same question, and induction finishes the proof.
All count[c] copies of c are smallest at once, so all of them start groups together; that is why we subtract 'groups' from each value in one step. If any value in the run is short, no split exists, because the forced groups can't be formed.`,
			CommonMistakes: []string{
				"Using a plain dict and reading count[value] for a card you don't hold, which raises KeyError (e.g. [8,10,12] asks for 9). Use Counter, which returns 0 for missing keys, or count.get(value, 0).",
				"Starting groups at an arbitrary card instead of the smallest, e.g. iterating over the hash map in insertion order. A group started at 2 can strand a 1 forever; always iterate sorted(count).",
				"Subtracting 1 instead of all copies of the smallest card and then not revisiting it. If you subtract one group at a time, loop while count[card] > 0.",
			},
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Check for failure", Explanation: "Groups have exactly groupSize cards and all cards are used, so the total must be a multiple of groupSize. Example 2 has 5 cards and groups of 4.", CodeSnippet: "if len(hand) % groupSize != 0:\n    return False", LineStart: 4, LineEnd: 5},
				{Title: "Track the greedy state", Explanation: "A count per value lets us remove a whole run in O(groupSize) time. Example 1: {1:1, 2:2, 3:2, 4:1, 6:1, 7:1, 8:1}.", CodeSnippet: "count = Counter(hand)", LineStart: 6, LineEnd: 6},
				{Title: "Order the input", Explanation: "Visit values smallest first. Any copies left of the current value have nothing smaller to attach to, so each one starts a group.", CodeSnippet: "for card in sorted(count):\n    groups = count[card]  # copies left = groups that must start here", LineStart: 7, LineEnd: 8},
				{Title: "Commit or skip", Explanation: "Value already used up by earlier groups? Skip it. Otherwise take 'groups' copies of each value in the run, failing if one is short. Example 1: at 1 take [1,2,3]; at 2 one copy is left, take [2,3,4]; 3 and 4 are now 0; at 6 take [6,7,8].", CodeSnippet: "if groups == 0:\n    continue\nfor value in range(card, card + groupSize):\n    if count[value] < groups:\n        return False\n    count[value] -= groups", LineStart: 9, LineEnd: 14},
				{Title: "Return the answer", Explanation: "Every value was absorbed into a forced group, so the split exists.", CodeSnippet: "return True", LineStart: 15, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{
					Name: "Remove from a sorted list (simple, O(n^2))",
					Code: `def isNStraightHand(hand, groupSize):
    cards = sorted(hand)
    while cards:
        smallest = cards[0]
        for value in range(smallest, smallest + groupSize):
            if value in cards:
                cards.remove(value)
            else:
                return False
    return True`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n)",
					Explanation:     "The same forced choice, but each lookup and removal scans the list. It also covers the length check implicitly: a short last group fails the 'value in cards' test.",
				},
				{
					Name:            "Min-heap of distinct values",
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Keep a min-heap (chapter 11) of distinct values and a count map. Build each group from heap[0]; when a value's count hits 0 it must be the heap's smallest, otherwise a gap would be left behind and the answer is false.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the greedy choice here 'forced' rather than just 'good', and why does that make the proof easy?", Answer: "In every valid split the smallest card is the lowest card of its group, so its group is exactly c..c+groupSize-1. Every solution contains that group, so taking it can't rule out any solution. An exchange argument normally swaps an optimal choice for the greedy one; here there is nothing to swap."},
			{Prompt: "What if you processed the largest card first instead?", Answer: "That works too, by the mirror-image argument: the largest card must be the top card of its group, so its group is c-groupSize+1..c. What doesn't work is starting from an arbitrary card, because a card in the middle could belong to several different groups."},
		},
	},
	{
		ID:              "merge-triplets-to-form-target",
		Number:          1899,
		Order:           6,
		Title:           "Merge Triplets to Form Target Triplet",
		Slug:            "merge-triplets-to-form-target-triplet",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Greedy"},
		RelatedChapters: []int{17},
		Description: `You are given a list of triplets, each a list of three integers [a, b, c], and a target triplet [x, y, z].

You may repeat this operation any number of times: pick two triplets i and j and replace triplet j with their position-wise maximum, [max(a_i, a_j), max(b_i, b_j), max(c_i, c_j)].

Return true if, after some sequence of operations (possibly none), the target appears as one of the triplets, and false otherwise. Equivalently: is there a non-empty set of the given triplets whose position-wise maximum is exactly the target?`,
		Constraints: []string{
			"1 <= triplets.length <= 10^5",
			"triplets[i].length == target.length == 3",
			"1 <= every value <= 1000",
		},
		Examples: []Example{
			{Input: "triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]", Output: "true", Explanation: "[1,8,4] has 8 > 7, so using it would push the middle value past the target forever; ignore it. The max of [2,5,3] and [1,7,5] is [2,7,5]."},
			{Input: "triplets = [[3,4,5],[4,5,6]], target = [3,2,5]", Output: "false", Explanation: "Both triplets have a middle value above 2, so neither can be used."},
			{Input: "triplets = [[2,5,3],[2,3,4],[1,2,5],[5,2,3]], target = [5,5,5]", Output: "true", Explanation: "All four are safe; [2,5,3] gives the 5 in the middle, [1,2,5] the last one and [5,2,3] the first."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"triplets": [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, "target": []int{2, 7, 5}}, Expected: true, Note: "example 1: one triplet must be ignored"},
			{Input: map[string]interface{}{"triplets": [][]int{{3, 4, 5}, {4, 5, 6}}, "target": []int{3, 2, 5}}, Expected: false, Note: "example 2: every triplet overshoots"},
			{Input: map[string]interface{}{"triplets": [][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, "target": []int{5, 5, 5}}, Expected: true, Hidden: true, Note: "example 3: three triplets each supply one position"},
			{Input: map[string]interface{}{"triplets": [][]int{{1, 2, 3}}, "target": []int{1, 2, 3}}, Expected: true, Hidden: true, Note: "single triplet equal to target"},
			{Input: map[string]interface{}{"triplets": [][]int{{5, 5, 6}, {1, 1, 1}}, "target": []int{5, 5, 5}}, Expected: false, Hidden: true, Note: "matches come from an overshooting triplet (catches skipping the filter)"},
			{Input: map[string]interface{}{"triplets": [][]int{{1, 1, 1}}, "target": []int{2, 2, 2}}, Expected: false, Hidden: true, Note: "everything too small"},
			{Input: map[string]interface{}{"triplets": [][]int{{1, 5, 3}, {5, 1, 1}}, "target": []int{5, 5, 3}}, Expected: true, Hidden: true, Note: "two triplets combine"},
			{Input: map[string]interface{}{"triplets": [][]int{{1, 2, 3}, {1, 2, 3}}, "target": []int{1, 2, 4}}, Expected: false, Hidden: true, Note: "one position never reached exactly"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def mergeTriplets(triplets, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The merge only takes maximums, so values can only go up. What does that tell you about a triplet that has even one value larger than the target's?"},
			{Level: 2, Type: "insight", Content: "A triplet with any value above the target is poison: using it makes that position too big forever. Every other triplet is safe: merging it can't push any position past the target. So use all safe triplets at once and check whether, together, they hit each target value exactly. Chapter 17 calls this taking every safe choice.", ChapterRef: 17},
			{Level: 3, Type: "plan", Content: "1. Track the greedy state: the set of target positions (0, 1, 2) matched so far.\n2. Order the input: any order works, so scan the triplets once.\n3. Commit or skip: skip a triplet if any of its values exceeds the target's value at that position; otherwise record every position where it equals the target.\n4. Return the answer: true if all three positions were matched.\n5. Justify the choice: why can adding a safe triplet never hurt?"},
			{Level: 4, Type: "skeleton", Content: "def mergeTriplets(triplets, target):\n    matched = set()\n    for triplet in triplets:\n        if ____:\n            continue\n        for k in range(3):\n            if ____:\n                matched.add(k)\n    return ____"},
		},
		Solution: Solution{
			Code: `def mergeTriplets(triplets, target):
    matched = set()  # target positions some safe triplet hits exactly
    for triplet in triplets:
        if any(triplet[k] > target[k] for k in range(3)):
            continue  # would overshoot forever: skip it
        for k in range(3):
            if triplet[k] == target[k]:
                matched.add(k)
    return len(matched) == 3`,
			Explanation: "Builds on Hand of Straights, where the greedy choice was forced; here the greedy rule is 'take everything that is safe'. Merging only takes maximums, so a triplet either can never be used (it overshoots somewhere) or can always be added without harm. That turns a question about subsets into a single filter-and-check pass.",
			BruteForce:  "Try every non-empty subset of triplets and compare its position-wise max with the target: O(2^n) subsets, hopeless beyond about 25 triplets. The waste: most subsets differ only by safe triplets that can't change whether the answer is reachable.",
			KeyInsight:  "Max only grows. A triplet with a value above the target is unusable, and every other triplet is free to include, so the best subset is simply 'all safe triplets'.",
			WhyItWorks: `Split the triplets into poison (some value above the target) and safe (every value <= the target).
Any subset containing a poison triplet has a max above the target in that position, so poison is never usable.
Exchange argument for the safe ones: take any subset S whose max equals the target. Add every other safe triplet to it. Each added value is <= the target in its position, so the max of S does not change: the bigger set still works. So if any working subset exists, 'all safe triplets' works too. And the max of all safe triplets equals the target exactly when, for each position k, some safe triplet has value target[k] there (all are <= target[k], so the max is target[k] iff one of them hits it). That is exactly what matched records.`,
			CommonMistakes: []string{
				"Recording matches from every triplet without filtering first. On [[5,5,6],[1,1,1]] with target [5,5,5], the first triplet matches positions 0 and 1 but also overshoots position 2, so the true answer is false.",
				"Trying to merge triplets step by step and checking equality after each merge. The order doesn't matter, and the incremental version tends to include poison triplets; filter, then check positions.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track the greedy state", Explanation: "matched holds the target positions that some safe triplet hits exactly. At most 3 entries, so constant space.", CodeSnippet: "matched = set()  # target positions some safe triplet hits exactly", LineStart: 2, LineEnd: 2},
				{Title: "Commit or skip", Explanation: "Skip poison: a value above the target can never come back down. In example 1, [1,8,4] is skipped because 8 > 7.", CodeSnippet: "if any(triplet[k] > target[k] for k in range(3)):\n    continue  # would overshoot forever: skip it", LineStart: 4, LineEnd: 5},
				{Title: "Commit or skip", Explanation: "Every safe triplet is committed. Record which positions it supplies exactly. In example 1, [2,5,3] supplies position 0 and [1,7,5] supplies positions 1 and 2.", CodeSnippet: "for k in range(3):\n    if triplet[k] == target[k]:\n        matched.add(k)", LineStart: 6, LineEnd: 8},
				{Title: "Return the answer", Explanation: "The max of all safe triplets equals the target exactly when all three positions are supplied.", CodeSnippet: "return len(matched) == 3", LineStart: 9, LineEnd: 9},
			},
			Alternatives: []AltSolution{
				{
					Name: "Try every subset (brute force)",
					Code: `from itertools import combinations

def mergeTriplets(triplets, target):
    for size in range(1, len(triplets) + 1):
        for chosen in combinations(triplets, size):
            merged = [max(t[k] for t in chosen) for k in range(3)]
            if merged == target:
                return True
    return False`,
					TimeComplexity:  "O(2^n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Checks every non-empty subset's position-wise maximum. Only usable for a handful of triplets, but it makes a good independent check of the greedy argument.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can you include every safe triplet without checking whether it helps?", Answer: "Each of its values is <= the target in the same position, so adding it to a set whose max is already the target leaves that max unchanged. Including it can only help (it might supply an exact match) and never hurt."},
			{Prompt: "What if the operation were position-wise minimum instead of maximum?", Answer: "Mirror the argument: values can only go down, so a triplet with any value below the target is poison, and every triplet with all values >= the target is safe. Keep the safe ones and check that each position is hit exactly by at least one of them."},
		},
	},
	{
		ID:              "partition-labels",
		Number:          763,
		Order:           7,
		Title:           "Partition Labels",
		Slug:            "partition-labels",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"String", "Hash Table", "Greedy", "Two Pointers"},
		RelatedChapters: []int{17, 2, 18},
		Description: `Given a string s of lowercase letters, cut it into as many pieces as possible so that each letter appears in at most one piece. (All copies of a letter must land in the same piece.) Keeping the pieces in order and joining them must give back s.

Return a list of the piece lengths, in order from left to right.`,
		Constraints: []string{
			"1 <= s.length <= 500",
			"s consists of lowercase English letters",
		},
		Examples: []Example{
			{Input: `s = "ababcbacadefegdehijhklij"`, Output: "[9,7,8]", Explanation: "The pieces are \"ababcbaca\", \"defegde\" and \"hijhklij\". The last 'a' is at index 8, and no letter in the first 9 characters appears later, so the first cut goes after index 8. Cutting earlier would split the a's."},
			{Input: `s = "eccbbbbdec"`, Output: "[10]", Explanation: "The first 'e' and the last 'e' are at the two ends, so the whole string is one piece."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "ababcbacadefegdehijhklij"}, Expected: []int{9, 7, 8}, Note: "example 1"},
			{Input: map[string]interface{}{"s": "eccbbbbdec"}, Expected: []int{10}, Note: "example 2: one piece"},
			{Input: map[string]interface{}{"s": "a"}, Expected: []int{1}, Hidden: true, Note: "single character"},
			{Input: map[string]interface{}{"s": "abc"}, Expected: []int{1, 1, 1}, Hidden: true, Note: "all distinct: every cut is valid"},
			{Input: map[string]interface{}{"s": "aaaa"}, Expected: []int{4}, Hidden: true, Note: "all equal"},
			{Input: map[string]interface{}{"s": "abac"}, Expected: []int{3, 1}, Hidden: true, Note: "inner letter ends before the outer one (catches end = last[ch] without max)"},
			{Input: map[string]interface{}{"s": "caedbdedda"}, Expected: []int{1, 9}, Hidden: true, Note: "piece end keeps growing"},
			{Input: map[string]interface{}{"s": "qiejxqfnqceocmy"}, Expected: []int{13, 1, 1}, Hidden: true, Note: "long first piece, short tail pieces"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def partitionLabels(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The first piece must contain s[0], so it must reach at least to the last copy of s[0]. As you walk through that stretch you meet other letters; what does each of them do to where the piece can end?"},
			{Level: 2, Type: "insight", Content: "Precompute the last index of every letter. Walk left to right keeping end = the furthest last-index of any letter in the current piece. When i reaches end, no letter in the piece appears later, so cut there: cutting at the earliest legal point gives the most pieces. It is the same 'extend the frontier' move as Jump Game; see chapter 17.", ChapterRef: 17},
			{Level: 3, Type: "plan", Content: "1. Track the greedy state: last[letter] for every letter, plus start and end of the current piece.\n2. Order the input: scan the string left to right.\n3. Commit or skip: extend end to max(end, last[s[i]]).\n4. When i == end, commit the cut: record end - start + 1 and start the next piece at i + 1.\n5. Return the answer: the list of lengths.\n6. Justify the choice: why is cutting as early as possible never worse?"},
			{Level: 4, Type: "skeleton", Content: "def partitionLabels(s):\n    last = ____\n    sizes = []\n    start = 0\n    end = 0\n    for i, ch in enumerate(s):\n        end = ____\n        if ____:\n            sizes.append(end - start + 1)\n            start = i + 1\n    return sizes"},
		},
		Solution: Solution{
			Code: `def partitionLabels(s):
    last = {ch: i for i, ch in enumerate(s)}  # last index of each letter
    sizes = []
    start = 0
    end = 0  # the current piece must reach at least here
    for i, ch in enumerate(s):
        end = max(end, last[ch])
        if i == end:  # nothing in this piece appears later
            sizes.append(end - start + 1)
            start = i + 1
    return sizes`,
			Explanation: "Builds on Jump Game's frontier (farthest) and Jump Game II's 'close a level when i reaches its end': here end is pushed out by each letter's last occurrence, and when i catches up with it we close a piece. Each letter is really an interval [first, last], so this is also a preview of merging intervals in chapter 18.",
			BruteForce:  "For each piece, try ending it at start, start+1, ... and check whether any letter in s[start..end] appears after end, rescanning the rest of the string each time: O(n^2) or worse. The waste: re-checking the rest of the string, when the last index of every letter can be looked up.",
			KeyInsight:  "With last[letter] precomputed, 'can I cut after index i?' becomes 'is i equal to the largest last index of the letters seen in this piece?', which a running max answers in O(1).",
			WhyItWorks: `A cut right after index i is legal exactly when no letter appears both at or before i and after i, i.e. every letter in s[0..i] has its last index <= i. Note that this depends only on i, not on where the other cuts are.
So the legal cut points can all be used at once, and using all of them gives the most pieces: any partition's cuts are a subset of the legal ones. The greedy cuts exactly at the legal points, because within the current piece, end is the largest last index seen, and i == end means every letter so far ends by i (letters from earlier pieces ended even sooner).
Greedy stays ahead: its k-th cut is the k-th legal cut point, which is no later than the k-th cut of any other valid partition.`,
			CommonMistakes: []string{
				"Setting end = last[ch] instead of max(end, last[ch]). On \"abac\", seeing 'b' at index 1 would set end = 1 and cut \"ab\", splitting the a's; the expected answer is [3,1].",
				"Recording piece lengths as end - start instead of end - start + 1, which is off by one for every piece.",
				"Building the last-index map with the first index (e.g. only setting it if the letter isn't there yet). Then every piece is cut too early.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track the greedy state", Explanation: "last maps each letter to its final position; later writes overwrite earlier ones, so the dict comprehension keeps the last index. At most 26 keys, so O(1) extra space.", CodeSnippet: "last = {ch: i for i, ch in enumerate(s)}  # last index of each letter", LineStart: 2, LineEnd: 2},
				{Title: "Track the greedy state", Explanation: "start is where the current piece begins; end is the furthest it is forced to reach.", CodeSnippet: "start = 0\nend = 0  # the current piece must reach at least here", LineStart: 4, LineEnd: 5},
				{Title: "Commit or skip", Explanation: "Each letter forces the piece to include its last copy. In example 1, 'a' forces end = 8; b (last 5) and c (last 7) don't push it further.", CodeSnippet: "end = max(end, last[ch])", LineStart: 7, LineEnd: 7},
				{Title: "Commit or skip", Explanation: "When i reaches end, every letter in the piece is finished, so this is the earliest legal cut. Commit to it and start the next piece.", CodeSnippet: "if i == end:  # nothing in this piece appears later\n    sizes.append(end - start + 1)\n    start = i + 1", LineStart: 8, LineEnd: 10},
				{Title: "Return the answer", Explanation: "The last character always closes a piece, since its index is the largest possible end.", CodeSnippet: "return sizes", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Grow each piece and rescan (brute force)",
					Code: `def partitionLabels(s):
    sizes = []
    start = 0
    while start < len(s):
        end = start
        while any(ch in set(s[start:end + 1]) for ch in s[end + 1:]):
            end += 1
        sizes.append(end - start + 1)
        start = end + 1
    return sizes`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n)",
					Explanation:     "Grows the piece one character at a time until no letter in it appears later, rescanning the rest of the string each time. Same cuts, much more work.",
				},
				{
					Name:            "Merge letter intervals",
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Turn each letter into the interval [first index, last index], sort by start, and merge overlapping intervals (chapter 18). The merged blocks are the pieces.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does cutting at the earliest legal point give the maximum number of pieces?", Answer: "Whether you may cut after index i depends only on i (no letter may straddle it), not on the other cuts. So the set of legal cuts is fixed and you can use all of them; the greedy uses exactly those, and any other partition uses a subset of them, so it can't have more pieces."},
			{Prompt: "How is this problem an interval problem in disguise?", Answer: "Each letter occupies the interval from its first to its last index. Two letters must share a piece when their intervals overlap. Merging overlapping intervals gives blocks, and those blocks are exactly the pieces. The running max of last indices is the merge's running end."},
		},
	},
	{
		ID:              "valid-parenthesis-string",
		Number:          678,
		Order:           8,
		Title:           "Valid Parenthesis String",
		Slug:            "valid-parenthesis-string",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"String", "Greedy", "Stack", "Dynamic Programming"},
		RelatedChapters: []int{17, 5, 15},
		Description: `A string s contains only the characters '(', ')' and '*'. Each '*' is a wildcard that you may turn into '(', into ')', or into nothing (the empty string).

The string is **valid** if you can choose the wildcards so that the result is a balanced parenthesis string: every '(' has a matching ')' after it, and every ')' has a matching '(' before it.

Return true if s can be made valid, and false otherwise.`,
		Constraints: []string{
			"1 <= s.length <= 100",
			"s[i] is '(', ')' or '*'",
		},
		Examples: []Example{
			{Input: `s = "(*)"`, Output: "true", Explanation: "Turn '*' into nothing: \"()\" is balanced."},
			{Input: `s = "(*))"`, Output: "true", Explanation: "Turn '*' into '(': \"(())\" is balanced."},
			{Input: `s = ")("`, Output: "false", Explanation: "The first ')' has nothing before it to match."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "(*)"}, Expected: true, Note: "example 1: star as empty (catches missing low = max(0, low))"},
			{Input: map[string]interface{}{"s": "(*))"}, Expected: true, Note: "example 2: star as '('"},
			{Input: map[string]interface{}{"s": ")("}, Expected: false, Note: "example 3: close before open"},
			{Input: map[string]interface{}{"s": "(()"}, Expected: false, Hidden: true, Note: "an open paren is never closed"},
			{Input: map[string]interface{}{"s": "*"}, Expected: true, Hidden: true, Note: "lone star becomes empty"},
			{Input: map[string]interface{}{"s": "("}, Expected: false, Hidden: true, Note: "single open paren"},
			{Input: map[string]interface{}{"s": "*))"}, Expected: false, Hidden: true, Note: "too many ')' even with the star as '(' (catches a missing high < 0 check)"},
			{Input: map[string]interface{}{"s": "**(("}, Expected: false, Hidden: true, Note: "stars before the opens can't close them"},
			{Input: map[string]interface{}{"s": "(*()"}, Expected: true, Hidden: true, Note: "star as ')'"},
			{Input: map[string]interface{}{"s": "*(*)*))((**"}, Expected: true, Hidden: true, Note: "longer mix"},
			{Input: map[string]interface{}{"s": "((*)(*))((*"}, Expected: false, Hidden: true, Note: "longer mix that fails at the end"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def checkValidString(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Without stars you'd track one number: how many '(' are still open. A star makes that number uncertain. Instead of guessing, what could you track about all the values it might be?"},
			{Level: 2, Type: "insight", Content: "Track the range [low, high] of possible open counts. '(' adds 1 to both; ')' subtracts 1 from both; '*' does low - 1 (use it as ')') and high + 1 (use it as '('). If high < 0, even the most generous choice has too many ')': fail. A negative count is a dead choice, so clamp low = max(0, low). At the end the string is valid iff low == 0. Rather than committing to one choice, the greedy keeps the best and worst cases; see chapter 17.", ChapterRef: 17},
			{Level: 3, Type: "plan", Content: "1. Track the greedy state: low = high = 0, the smallest and largest possible number of unmatched '('.\n2. Order the input: scan left to right.\n3. Commit or skip: update low and high for '(', ')' or '*'.\n4. Check for failure: if high < 0, return false. Then clamp low to 0, dropping choices that already went negative.\n5. Return the answer: low == 0, i.e. some choice ends with nothing open.\n6. Justify the choice: why is every count between low and high achievable?"},
			{Level: 4, Type: "skeleton", Content: "def checkValidString(s):\n    low = 0\n    high = 0\n    for ch in s:\n        if ch == '(':\n            low += 1\n            high += 1\n        elif ch == ')':\n            low -= 1\n            high -= 1\n        else:\n            low ____\n            high ____\n        if ____:\n            return False\n        low = ____\n    return low == 0"},
		},
		Solution: Solution{
			Code: `def checkValidString(s):
    low = 0   # fewest '(' that can be open now
    high = 0  # most '(' that can be open now
    for ch in s:
        if ch == '(':
            low += 1
            high += 1
        elif ch == ')':
            low -= 1
            high -= 1
        else:  # '*' could be ')', nothing, or '('
            low -= 1
            high += 1
        if high < 0:
            return False  # too many ')' whatever the stars are
        low = max(low, 0)  # a negative count is a dead choice
    return low == 0`,
			Explanation: "The last rung, and the subtlest greedy state: instead of one number (Jump Game's farthest, Kadane's current) it tracks two, the smallest and largest possible open-paren counts across every way of choosing the stars. It builds on the balance counter from the stack chapter (5), and replaces the O(n^2) DP over (index, open count) from chapter 15 with its two extreme values.",
			BruteForce:  "Try all 3^k ways to replace k stars and check each string with a counter: 3^100 is far too many. A DP over (index, open count) brings it to O(n^2). The waste in the DP: it tracks every possible count, but the possible counts always form one unbroken range, so its two ends are enough.",
			KeyInsight:  "The set of open counts you could have after reading a prefix is a contiguous range [low, high]. Each character shifts the ends predictably, so two numbers replace the whole DP table.",
			WhyItWorks: `Invariant: after each character, the achievable open counts (over choices that never went negative) are exactly the integers in [low, high].
Why it's a range: '(' and ')' shift every count by the same amount, and '*' lets each count c become c-1, c or c+1, which turns a range [a, b] into [a-1, b+1] with no holes.
Why high < 0 means false: high is the count when every star so far became '('. If even that is negative, some ')' has no partner under every choice, and nothing later can fix a prefix that already closed too much.
Why low = max(0, low): a count of -1 means that choice already closed a paren that was never opened, which is invalid forever, so it is removed from the range. Without the clamp, "(*)" ends with low = -1 and you'd wrongly answer false, because low would describe a dead choice.
At the end the string is valid iff 0 is achievable, i.e. low == 0 (high >= 0 is already guaranteed).`,
			CommonMistakes: []string{
				"Forgetting low = max(0, low). On \"(*)\" low goes 1, 0, -1 and the final check fails, though \"()\" works. The clamp removes choices that already went negative.",
				"Forgetting the high < 0 check. On \"*))\" low keeps being clamped to 0 and the function returns true, but even with the star as '(' there are too many ')'.",
				"Greedily deciding each star on the spot (for example, always using it as '(' when there are unmatched ')'). A local choice can be undone by later characters; keeping the whole range avoids committing too early.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track the greedy state", Explanation: "Before reading anything, exactly 0 parens are open, so the range is [0, 0].", CodeSnippet: "low = 0   # fewest '(' that can be open now\nhigh = 0  # most '(' that can be open now", LineStart: 2, LineEnd: 3},
				{Title: "Commit or skip", Explanation: "Fixed characters shift the range; a star widens it by one on each side. On \"(*))\": after '(' the range is [1,1], after '*' [0,2], after ')' [-1,1].", CodeSnippet: "if ch == '(':\n    low += 1\n    high += 1\nelif ch == ')':\n    low -= 1\n    high -= 1\nelse:  # '*' could be ')', nothing, or '('\n    low -= 1\n    high += 1", LineStart: 5, LineEnd: 13},
				{Title: "Check for failure", Explanation: "If even the most '('-heavy choice has a negative count, the prefix has an unmatched ')'. On \"*))\" high goes 1, 0, -1 and we return false.", CodeSnippet: "if high < 0:\n    return False  # too many ')' whatever the stars are", LineStart: 14, LineEnd: 15},
				{Title: "Track the greedy state", Explanation: "Negative counts are dead choices, so drop them from the range. On \"(*))\" [-1,1] becomes [0,1], then the last ')' gives [-1,0] -> [0,0].", CodeSnippet: "low = max(low, 0)  # a negative count is a dead choice", LineStart: 16, LineEnd: 16},
				{Title: "Return the answer", Explanation: "Valid iff some choice ends with no open parens, i.e. 0 is in the final range.", CodeSnippet: "return low == 0", LineStart: 17, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{
					Name: "DP over (index, open count)",
					Code: `from functools import lru_cache

def checkValidString(s):
    @lru_cache(maxsize=None)
    def can_finish(i, open_count):
        if open_count < 0:
            return False
        if i == len(s):
            return open_count == 0
        if s[i] == '(':
            return can_finish(i + 1, open_count + 1)
        if s[i] == ')':
            return can_finish(i + 1, open_count - 1)
        return (can_finish(i + 1, open_count + 1)
                or can_finish(i + 1, open_count - 1)
                or can_finish(i + 1, open_count))

    return can_finish(0, 0)`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n^2)",
					Explanation:     "Memoised recursion (chapter 15): can the rest of the string be finished from position i with open_count unmatched '('? It tries all three meanings of each star. Correct, and a good check on the greedy, but it stores every count where the greedy keeps only the two ends.",
				},
				{
					Name:            "Two stacks of indices",
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Push indices of '(' and '*' on separate stacks (chapter 5). Match each ')' with an open paren if possible, else a star. At the end, pair leftover '(' with stars that come after them; any '(' left over means false.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it correct to clamp low to 0 but never to raise high back up when it goes negative?", Answer: "low going negative only kills the most ')'-heavy choices; other choices in the range are still alive, so we drop the dead part and keep the rest. high going negative means even the most '('-heavy choice is dead, so the whole range is empty and the answer is false. There is nothing left to clamp."},
			{Prompt: "Would the [low, high] trick still work with two bracket types, '(' ')' and '[' ']', plus wildcards?", Answer: "No. The range only counts how many brackets are open, and counting is enough for one type. With two types, \"([)]\" has balanced counts but is invalid because the order of the open brackets matters. You'd need a stack of open types (chapter 5), and with wildcards a search or DP over stack states."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, GreedyProblems...)
}
