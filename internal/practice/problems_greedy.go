package practice

// GreedyProblems contains all greedy category problems
var GreedyProblems = []*Problem{
	{
		ID:              "maximum-subarray",
		Number:          88,
		Title:           "Maximum Subarray",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Divide and Conquer", "Dynamic Programming"},
		RelatedChapters: []int{8, 9},
		Description: `Given an integer array nums, find the subarray with the largest sum, and return its sum.`,
		Constraints: []string{
			"1 <= nums.length <= 10^5",
			"-10^4 <= nums[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "nums = [-2,1,-3,4,-1,2,1,-5,4]", Output: "6", Explanation: "The subarray [4,-1,2,1] has the largest sum 6."},
			{Input: "nums = [1]", Output: "1"},
			{Input: "nums = [5,4,-1,7,8]", Output: "23"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, Expected: 6},
			{Input: map[string]interface{}{"nums": []int{1}}, Expected: 1},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def maxSubArray(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Kadane's algorithm: track current sum, reset to 0 if negative. Track max sum seen.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "For each element: current_sum = max(num, current_sum + num). max_sum = max(max_sum, current_sum)."},
			{Level: 3, Type: "code", Content: "current_sum = 0, max_sum = nums[0]. for num: current_sum = max(num, current_sum + num). max_sum = max(max_sum, current_sum)."},
		},
		Solution: Solution{
			Code: `def maxSubArray(nums):
    max_sum = nums[0]
    current_sum = 0

    for num in nums:
        current_sum = max(num, current_sum + num)
        max_sum = max(max_sum, current_sum)

    return max_sum`,
			Explanation:     "Kadane's algorithm: extend current subarray or start new one. Track maximum sum seen.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Extend or restart", Explanation: "Start new subarray if current sum is negative", CodeSnippet: "current_sum = max(num, current_sum + num)", LineStart: 6, LineEnd: 6},
				{Title: "Track maximum", Explanation: "Update max after each step", CodeSnippet: "max_sum = max(max_sum, current_sum)", LineStart: 7, LineEnd: 7},
			},
		},
	},
	{
		ID:              "jump-game-ii",
		Number:          89,
		Title:           "Jump Game II",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Dynamic Programming", "Greedy"},
		RelatedChapters: []int{8, 9},
		Description: `You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i.

Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].`,
		Constraints: []string{
			"1 <= nums.length <= 10^4",
			"0 <= nums[i] <= 1000",
			"It's guaranteed that you can reach nums[n - 1]",
		},
		Examples: []Example{
			{Input: "nums = [2,3,1,1,4]", Output: "2", Explanation: "Jump from 0 to 1, then from 1 to 4."},
			{Input: "nums = [2,3,0,1,4]", Output: "2"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 3, 1, 1, 4}}, Expected: 2},
			{Input: map[string]interface{}{"nums": []int{2, 3, 0, 1, 4}}, Expected: 2},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def jump(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Greedy: track current jump's end and farthest reachable. When reaching end, must jump and update end to farthest.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "Track end of current jump range and farthest point. When i reaches end, increment jumps and set end = farthest."},
			{Level: 3, Type: "code", Content: "jumps = 0, end = 0, farthest = 0. for i in 0..n-1: farthest = max(farthest, i + nums[i]). if i == end: jumps++, end = farthest."},
		},
		Solution: Solution{
			Code: `def jump(nums):
    n = len(nums)
    if n <= 1:
        return 0

    jumps = 0
    current_end = 0
    farthest = 0

    for i in range(n - 1):
        farthest = max(farthest, i + nums[i])

        if i == current_end:
            jumps += 1
            current_end = farthest

    return jumps`,
			Explanation:     "Greedy: track farthest reachable within current jump. When reaching current boundary, must jump.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track farthest", Explanation: "Maximum reach from any position in current range", CodeSnippet: "farthest = max(farthest, i + nums[i])", LineStart: 11, LineEnd: 11},
				{Title: "Jump when needed", Explanation: "When reaching end of current jump range", CodeSnippet: "if i == current_end:\n    jumps += 1\n    current_end = farthest", LineStart: 13, LineEnd: 15},
			},
		},
	},
	{
		ID:              "gas-station",
		Number:          106,
		Title:           "Gas Station",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Greedy"},
		RelatedChapters: []int{8},
		Description: `There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique.`,
		Constraints: []string{
			"n == gas.length == cost.length",
			"1 <= n <= 10^5",
			"0 <= gas[i], cost[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "gas = [1,2,3,4,5], cost = [3,4,5,1,2]", Output: "3"},
			{Input: "gas = [2,3,4], cost = [3,4,3]", Output: "-1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"gas": []int{1, 2, 3, 4, 5}, "cost": []int{3, 4, 5, 1, 2}}, Expected: 3},
			{Input: map[string]interface{}{"gas": []int{2, 3, 4}, "cost": []int{3, 4, 3}}, Expected: -1},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def canCompleteCircuit(gas, cost):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "If total gas >= total cost, solution exists. Find starting point where tank never goes negative.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "Track current tank. If it goes negative, start from next station. If total >= 0, answer is the last starting point."},
			{Level: 3, Type: "code", Content: "total = 0, tank = 0, start = 0. for i: diff = gas[i] - cost[i]. total += diff, tank += diff. if tank < 0: start = i + 1, tank = 0."},
		},
		Solution: Solution{
			Code: `def canCompleteCircuit(gas, cost):
    total = 0
    tank = 0
    start = 0

    for i in range(len(gas)):
        diff = gas[i] - cost[i]
        total += diff
        tank += diff

        if tank < 0:
            start = i + 1
            tank = 0

    return start if total >= 0 else -1`,
			Explanation:     "If total gas >= total cost, solution exists. Start from station after any point where tank would go negative.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track total and current", Explanation: "Total determines if solution exists", CodeSnippet: "total += diff\ntank += diff", LineStart: 8, LineEnd: 9},
				{Title: "Reset on negative", Explanation: "Can't start from any previous station", CodeSnippet: "if tank < 0:\n    start = i + 1\n    tank = 0", LineStart: 11, LineEnd: 13},
			},
		},
	},
	{
		ID:              "partition-labels",
		Number:          107,
		Title:           "Partition Labels",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Hash Table", "Two Pointers", "String", "Greedy"},
		RelatedChapters: []int{5, 8},
		Description: `You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.

Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.

Return a list of integers representing the size of these parts.`,
		Constraints: []string{
			"1 <= s.length <= 500",
			"s consists of lowercase English letters",
		},
		Examples: []Example{
			{Input: `s = "ababcbacadefegdehijhklij"`, Output: "[9,7,8]"},
			{Input: `s = "eccbbbbdec"`, Output: "[10]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "ababcbacadefegdehijhklij"}, Expected: []int{9, 7, 8}},
			{Input: map[string]interface{}{"s": "eccbbbbdec"}, Expected: []int{10}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def partitionLabels(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Find last occurrence of each character. Extend partition to include all last occurrences of characters seen so far.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "Track end = max last index of chars seen. When i reaches end, we found a partition."},
			{Level: 3, Type: "code", Content: "last = {c: i for each char}. start = end = 0. for i, c: end = max(end, last[c]). if i == end: add partition, start = i + 1."},
		},
		Solution: Solution{
			Code: `def partitionLabels(s):
    last = {c: i for i, c in enumerate(s)}
    result = []
    start = end = 0

    for i, c in enumerate(s):
        end = max(end, last[c])
        if i == end:
            result.append(end - start + 1)
            start = i + 1

    return result`,
			Explanation:     "Extend partition to include last occurrence of each character seen. Create partition when current index reaches end.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1) - 26 letters max",
			Walkthrough: []WalkthroughStep{
				{Title: "Find last occurrences", Explanation: "Map each char to its last index", CodeSnippet: "last = {c: i for i, c in enumerate(s)}", LineStart: 2, LineEnd: 2},
				{Title: "Extend partition", Explanation: "Must include all chars' last occurrences", CodeSnippet: "end = max(end, last[c])", LineStart: 7, LineEnd: 7},
				{Title: "Create partition", Explanation: "When we reach end, partition is complete", CodeSnippet: "if i == end:\n    result.append(end - start + 1)", LineStart: 8, LineEnd: 9},
			},
		},
	},
	{
		ID:              "valid-parenthesis-string",
		Number:          108,
		Title:           "Valid Parenthesis String",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"String", "Dynamic Programming", "Stack", "Greedy"},
		RelatedChapters: []int{3, 8},
		Description: `Given a string s containing only three types of characters: '(', ')' and '*', return true if s is valid.

The following rules define a valid string:
- Any left parenthesis '(' must have a corresponding right parenthesis ')'.
- Any right parenthesis ')' must have a corresponding left parenthesis '('.
- Left parenthesis '(' must go before the corresponding right parenthesis ')'.
- '*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string "".`,
		Constraints: []string{
			"1 <= s.length <= 100",
			"s[i] is '(', ')' or '*'",
		},
		Examples: []Example{
			{Input: `s = "()"`, Output: "true"},
			{Input: `s = "(*)"`, Output: "true"},
			{Input: `s = "(*))"`, Output: "true"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "()"}, Expected: true},
			{Input: map[string]interface{}{"s": "(*)"}, Expected: true},
			{Input: map[string]interface{}{"s": "(*))"}, Expected: true},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def checkValidString(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Track range of possible open parentheses [low, high]. '*' can be '(', ')', or empty.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "'(' increases both. ')' decreases both. '*' decreases low (as ')') and increases high (as '(')."},
			{Level: 3, Type: "code", Content: "low = high = 0. for c: update low, high. if high < 0: return False. low = max(0, low). return low == 0."},
		},
		Solution: Solution{
			Code: `def checkValidString(s):
    low = high = 0

    for c in s:
        if c == '(':
            low += 1
            high += 1
        elif c == ')':
            low -= 1
            high -= 1
        else:  # '*'
            low -= 1   # Treat as ')'
            high += 1  # Treat as '('

        if high < 0:
            return False
        low = max(0, low)  # Can't have negative open count

    return low == 0`,
			Explanation:     "Track range [low, high] of possible open parentheses. Valid if 0 is in range at end.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track range", Explanation: "low = min opens possible, high = max opens possible", CodeSnippet: "low = high = 0", LineStart: 2, LineEnd: 2},
				{Title: "Handle '*'", Explanation: "Could be '(', ')', or empty", CodeSnippet: "low -= 1\nhigh += 1", LineStart: 12, LineEnd: 13},
				{Title: "Validate", Explanation: "High < 0 means too many ')'. Low must be 0 at end.", CodeSnippet: "return low == 0", LineStart: 19, LineEnd: 19},
			},
		},
	},
	{
		ID:              "hand-of-straights",
		Number:          143,
		Title:           "Hand of Straights",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Hash Table", "Greedy", "Sorting"},
		RelatedChapters: []int{8, 12},
		Description: `Alice has some number of cards and she wants to rearrange the cards into groups so that each group is of size groupSize, and consists of groupSize consecutive cards.

Given an integer array hand where hand[i] is the value written on the ith card and an integer groupSize, return true if she can rearrange the cards, or false otherwise.`,
		Constraints: []string{
			"1 <= hand.length <= 10^4",
			"0 <= hand[i] <= 10^9",
			"1 <= groupSize <= hand.length",
		},
		Examples: []Example{
			{Input: "hand = [1,2,3,6,2,3,4,7,8], groupSize = 3", Output: "true", Explanation: "Groups: [1,2,3], [2,3,4], [6,7,8]"},
			{Input: "hand = [1,2,3,4,5], groupSize = 4", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"hand": []int{1, 2, 3, 6, 2, 3, 4, 7, 8}, "groupSize": 3}, Expected: true},
			{Input: map[string]interface{}{"hand": []int{1, 2, 3, 4, 5}, "groupSize": 4}, Expected: false},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def isNStraightHand(hand, groupSize):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Greedy: always start group from smallest available card.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "Count cards. Sort unique values. For each smallest, try to form consecutive group."},
			{Level: 3, Type: "code", Content: "cnt = Counter. for card in sorted(cnt): while cnt[card]: for i in range(groupSize): if not cnt[card+i]: return False; cnt[card+i]-=1."},
		},
		Solution: Solution{
			Code: `def isNStraightHand(hand, groupSize):
    from collections import Counter

    if len(hand) % groupSize != 0:
        return False

    cnt = Counter(hand)

    for card in sorted(cnt):
        while cnt[card] > 0:
            for i in range(groupSize):
                if cnt[card + i] <= 0:
                    return False
                cnt[card + i] -= 1

    return True`,
			Explanation:     "Greedy: start from smallest card, form consecutive groups.",
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Check divisibility", Explanation: "Must divide evenly into groups", CodeSnippet: "if len(hand) % groupSize != 0:", LineStart: 4, LineEnd: 4},
				{Title: "Start smallest", Explanation: "Greedy: use smallest available first", CodeSnippet: "for card in sorted(cnt):", LineStart: 9, LineEnd: 9},
				{Title: "Form group", Explanation: "Need groupSize consecutive cards", CodeSnippet: "for i in range(groupSize):", LineStart: 11, LineEnd: 11},
			},
		},
	},
	{
		ID:              "merge-triplets-to-form-target",
		Number:          144,
		Title:           "Merge Triplets to Form Target Triplet",
		Difficulty:      "Medium",
		Category:        "greedy",
		Tags:            []string{"Array", "Greedy"},
		RelatedChapters: []int{8, 12},
		Description: `A triplet is an array of three integers. You are given a 2D integer array triplets, where triplets[i] = [ai, bi, ci] describes the ith triplet. You are also given an integer array target = [x, y, z] that describes the triplet you want to obtain.

To obtain target, you may apply the following operation on triplets any number of times (possibly zero):

Choose two indices (0-indexed) i and j (i != j) and update triplets[j] to become [max(ai, aj), max(bi, bj), max(ci, cj)].

Return true if it is possible to obtain the target triplet [x, y, z] as an element of triplets, or false otherwise.`,
		Constraints: []string{
			"1 <= triplets.length <= 10^5",
			"triplets[i].length == target.length == 3",
			"1 <= ai, bi, ci, x, y, z <= 1000",
		},
		Examples: []Example{
			{Input: "triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]", Output: "true"},
			{Input: "triplets = [[3,4,5],[4,5,6]], target = [3,2,5]", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"triplets": [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, "target": []int{2, 7, 5}}, Expected: true},
			{Input: map[string]interface{}{"triplets": [][]int{{3, 4, 5}, {4, 5, 6}}, "target": []int{3, 2, 5}}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def mergeTriplets(triplets, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "A triplet is usable only if no element exceeds target. Greedy: track which target positions achieved.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "Skip triplets with any value > target. From valid triplets, check if we can achieve each target value."},
			{Level: 3, Type: "code", Content: "good = set(). for t in triplets: if all(t[i]<=target[i]): for i: if t[i]==target[i]: good.add(i). return len(good)==3."},
		},
		Solution: Solution{
			Code: `def mergeTriplets(triplets, target):
    good = set()

    for t in triplets:
        if t[0] > target[0] or t[1] > target[1] or t[2] > target[2]:
            continue

        for i in range(3):
            if t[i] == target[i]:
                good.add(i)

    return len(good) == 3`,
			Explanation:     "Only use triplets that don't exceed target. Track which positions we can achieve exactly.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Filter invalid", Explanation: "Any value > target makes triplet unusable", CodeSnippet: "if t[0] > target[0] or t[1] > target[1] or t[2] > target[2]:", LineStart: 5, LineEnd: 5},
				{Title: "Track matches", Explanation: "Record positions where we hit target exactly", CodeSnippet: "if t[i] == target[i]:\n    good.add(i)", LineStart: 9, LineEnd: 10},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, GreedyProblems...)
}
