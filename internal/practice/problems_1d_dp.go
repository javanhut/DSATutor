package practice

// DP1DProblems contains all 1d-dp category problems
var DP1DProblems = []*Problem{
	{
		ID:         "climbing-stairs",
		Number:     36,
		Title:      "Climbing Stairs",
		Difficulty: "Easy",
		Category:   "1d-dp",
		Tags:       []string{"Dynamic Programming", "Math", "Memoization"},
		RelatedChapters: []int{12},
		Description: `You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?`,
		Constraints: []string{
			"1 <= n <= 45",
		},
		Examples: []Example{
			{Input: "n = 2", Output: "2", Explanation: "1+1 or 2"},
			{Input: "n = 3", Output: "3", Explanation: "1+1+1, 1+2, or 2+1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 2}, Expected: 2},
			{Input: map[string]interface{}{"n": 3}, Expected: 3},
			{Input: map[string]interface{}{"n": 5}, Expected: 8},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def climbStairs(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "This is the Fibonacci sequence in disguise."},
			{Level: 2, Type: "algorithm", Content: "dp[i] = dp[i-1] + dp[i-2]"},
		},
		Solution: Solution{
			Code: `def climbStairs(n):
    if n <= 2:
        return n
    prev, curr = 1, 2
    for i in range(3, n + 1):
        prev, curr = curr, prev + curr
    return curr`,
			Explanation: "Fibonacci: ways to reach step n = ways to reach n-1 + ways to reach n-2.",
		},
	},
	{
		ID:         "house-robber",
		Number:     37,
		Title:      "House Robber",
		Difficulty: "Medium",
		Category:   "1d-dp",
		Tags:       []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{12},
		Description: `You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.`,
		Constraints: []string{
			"1 <= nums.length <= 100",
			"0 <= nums[i] <= 400",
		},
		Examples: []Example{
			{Input: "nums = [1,2,3,1]", Output: "4", Explanation: "Rob house 1 (1) + house 3 (3) = 4"},
			{Input: "nums = [2,7,9,3,1]", Output: "12", Explanation: "Rob house 1 (2) + house 3 (9) + house 5 (1) = 12"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 1}}, Expected: 4},
			{Input: map[string]interface{}{"nums": []int{2, 7, 9, 3, 1}}, Expected: 12},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def rob(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "At each house, decide: rob it (add to max from 2 houses ago) or skip (keep max from previous)."},
			{Level: 2, Type: "algorithm", Content: "dp[i] = max(dp[i-1], dp[i-2] + nums[i])"},
		},
		Solution: Solution{
			Code: `def rob(nums):
    if len(nums) == 1:
        return nums[0]
    prev2, prev1 = 0, 0
    for num in nums:
        prev2, prev1 = prev1, max(prev1, prev2 + num)
    return prev1`,
			Explanation: "Track max if we rob this house vs skip. Only need two previous values.",
		},
	},
	{
		ID:              "house-robber-ii",
		Number:          83,
		Title:           "House Robber II",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.`,
		Constraints: []string{
			"1 <= nums.length <= 100",
			"0 <= nums[i] <= 1000",
		},
		Examples: []Example{
			{Input: "nums = [2,3,2]", Output: "3", Explanation: "Rob house 2 (money = 3)."},
			{Input: "nums = [1,2,3,1]", Output: "4", Explanation: "Rob house 1 and house 3."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 3, 2}}, Expected: 3},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 1}}, Expected: 4},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def rob(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Since houses form a circle, can't rob both first and last. Run House Robber I twice: exclude first OR exclude last.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "Helper function for linear house robber. Return max of robbing houses[0:n-1] and houses[1:n]."},
			{Level: 3, Type: "code", Content: "def rob_linear(houses): two pointers prev1, prev2. return max(rob_linear(nums[:-1]), rob_linear(nums[1:]))."},
		},
		Solution: Solution{
			Code: `def rob(nums):
    if len(nums) == 1:
        return nums[0]

    def rob_linear(houses):
        prev1, prev2 = 0, 0
        for money in houses:
            prev1, prev2 = max(prev1, prev2 + money), prev1
        return prev1

    return max(rob_linear(nums[:-1]), rob_linear(nums[1:]))`,
			Explanation:     "Can't rob both first and last house (circular). Solve linear House Robber twice: exclude first OR exclude last.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle single house", Explanation: "Edge case: only one house", CodeSnippet: "if len(nums) == 1:\n    return nums[0]", LineStart: 2, LineEnd: 3},
				{Title: "Linear robber", Explanation: "Standard house robber DP", CodeSnippet: "prev1, prev2 = max(prev1, prev2 + money), prev1", LineStart: 8, LineEnd: 8},
				{Title: "Two cases", Explanation: "Max of excluding first vs excluding last", CodeSnippet: "return max(rob_linear(nums[:-1]), rob_linear(nums[1:]))", LineStart: 11, LineEnd: 11},
			},
		},
	},
	{
		ID:              "longest-palindromic-substring",
		Number:          84,
		Title:           "Longest Palindromic Substring",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `Given a string s, return the longest palindromic substring in s.`,
		Constraints: []string{
			"1 <= s.length <= 1000",
			"s consist of only digits and English letters",
		},
		Examples: []Example{
			{Input: `s = "babad"`, Output: `"bab"`, Explanation: `"aba" is also valid.`},
			{Input: `s = "cbbd"`, Output: `"bb"`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "babad"}, Expected: "bab"},
			{Input: map[string]interface{}{"s": "cbbd"}, Expected: "bb"},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def longestPalindrome(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Expand around center. For each position, expand outward while characters match. Try both odd and even length centers."},
			{Level: 2, Type: "algorithm", Content: "For each index i, expand from (i, i) for odd length, (i, i+1) for even length. Track longest found."},
			{Level: 3, Type: "code", Content: "def expand(l, r): while l >= 0 and r < n and s[l] == s[r]: expand. Return length. Track start and max_len."},
		},
		Solution: Solution{
			Code: `def longestPalindrome(s):
    n = len(s)
    start, max_len = 0, 1

    def expand(l, r):
        while l >= 0 and r < n and s[l] == s[r]:
            l -= 1
            r += 1
        return r - l - 1  # Length of palindrome

    for i in range(n):
        len1 = expand(i, i)      # Odd length
        len2 = expand(i, i + 1)  # Even length
        length = max(len1, len2)

        if length > max_len:
            max_len = length
            start = i - (length - 1) // 2

    return s[start:start + max_len]`,
			Explanation:     "Expand around each center. Try both odd (single center) and even (double center) palindromes. Track longest.",
			TimeComplexity:  "O(n^2)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Expand helper", Explanation: "Expand while characters match", CodeSnippet: "while l >= 0 and r < n and s[l] == s[r]:", LineStart: 6, LineEnd: 8},
				{Title: "Two center types", Explanation: "Odd (i, i) and even (i, i+1)", CodeSnippet: "len1 = expand(i, i)\nlen2 = expand(i, i + 1)", LineStart: 12, LineEnd: 13},
				{Title: "Calculate start", Explanation: "Convert center and length to start index", CodeSnippet: "start = i - (length - 1) // 2", LineStart: 18, LineEnd: 18},
			},
		},
	},
	{
		ID:              "coin-change",
		Number:          85,
		Title:           "Coin Change",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Dynamic Programming", "BFS"},
		RelatedChapters: []int{9, 12},
		Description: `You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.`,
		Constraints: []string{
			"1 <= coins.length <= 12",
			"1 <= coins[i] <= 2^31 - 1",
			"0 <= amount <= 10^4",
		},
		Examples: []Example{
			{Input: "coins = [1,2,5], amount = 11", Output: "3", Explanation: "11 = 5 + 5 + 1"},
			{Input: "coins = [2], amount = 3", Output: "-1"},
			{Input: "coins = [1], amount = 0", Output: "0"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"coins": []int{1, 2, 5}, "amount": 11}, Expected: 3},
			{Input: map[string]interface{}{"coins": []int{2}, "amount": 3}, Expected: -1},
		},
		TimeComplexity:  "O(amount * n)",
		SpaceComplexity: "O(amount)",
		StarterCode:     "def coinChange(coins, amount):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "DP where dp[i] = min coins for amount i. For each amount, try each coin: dp[i] = min(dp[i], dp[i-coin] + 1).", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "Initialize dp with infinity (impossible). dp[0] = 0. For amount 1 to target, try all coins."},
			{Level: 3, Type: "code", Content: "dp = [inf] * (amount + 1); dp[0] = 0. for i in 1..amount: for coin: if coin <= i: dp[i] = min(dp[i], dp[i-coin]+1)."},
		},
		Solution: Solution{
			Code: `def coinChange(coins, amount):
    dp = [float('inf')] * (amount + 1)
    dp[0] = 0

    for i in range(1, amount + 1):
        for coin in coins:
            if coin <= i:
                dp[i] = min(dp[i], dp[i - coin] + 1)

    return dp[amount] if dp[amount] != float('inf') else -1`,
			Explanation:     "DP: dp[i] = minimum coins for amount i. Try each coin and take minimum.",
			TimeComplexity:  "O(amount * n)",
			SpaceComplexity: "O(amount)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize", Explanation: "Infinity for impossible, 0 for base case", CodeSnippet: "dp = [float('inf')] * (amount + 1)\ndp[0] = 0", LineStart: 2, LineEnd: 3},
				{Title: "Try each coin", Explanation: "Only if coin <= current amount", CodeSnippet: "if coin <= i:\n    dp[i] = min(dp[i], dp[i - coin] + 1)", LineStart: 7, LineEnd: 8},
				{Title: "Return result", Explanation: "-1 if still infinity (impossible)", CodeSnippet: "return dp[amount] if dp[amount] != float('inf') else -1", LineStart: 10, LineEnd: 10},
			},
		},
	},
	{
		ID:              "longest-increasing-subsequence",
		Number:          86,
		Title:           "Longest Increasing Subsequence",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Binary Search", "Dynamic Programming"},
		RelatedChapters: []int{1, 9, 12},
		Description: `Given an integer array nums, return the length of the longest strictly increasing subsequence.`,
		Constraints: []string{
			"1 <= nums.length <= 2500",
			"-10^4 <= nums[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "nums = [10,9,2,5,3,7,101,18]", Output: "4", Explanation: "The LIS is [2,3,7,101]."},
			{Input: "nums = [0,1,0,3,2,3]", Output: "4"},
			{Input: "nums = [7,7,7,7,7,7,7]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{10, 9, 2, 5, 3, 7, 101, 18}}, Expected: 4},
			{Input: map[string]interface{}{"nums": []int{0, 1, 0, 3, 2, 3}}, Expected: 4},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def lengthOfLIS(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "O(n^2): dp[i] = LIS ending at i. O(n log n): maintain array where arr[i] = smallest ending element of LIS of length i+1.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "Binary search approach: maintain sorted array of smallest endings. For each num, binary search for position to replace/extend."},
			{Level: 3, Type: "code", Content: "tails = []. For each num: binary search for position. If at end, append. Else replace tails[pos]. Return len(tails)."},
		},
		Solution: Solution{
			Code: `from bisect import bisect_left

def lengthOfLIS(nums):
    tails = []

    for num in nums:
        pos = bisect_left(tails, num)
        if pos == len(tails):
            tails.append(num)
        else:
            tails[pos] = num

    return len(tails)`,
			Explanation:     "Maintain array where tails[i] is smallest ending of LIS of length i+1. Binary search for position. Length of tails is LIS length.",
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Binary search position", Explanation: "Find where num would be inserted", CodeSnippet: "pos = bisect_left(tails, num)", LineStart: 7, LineEnd: 7},
				{Title: "Extend or replace", Explanation: "Extend if at end, else replace for smaller ending", CodeSnippet: "if pos == len(tails):\n    tails.append(num)\nelse:\n    tails[pos] = num", LineStart: 8, LineEnd: 11},
				{Title: "Result", Explanation: "Length of tails = LIS length", CodeSnippet: "return len(tails)", LineStart: 13, LineEnd: 13},
			},
		},
	},
	{
		ID:              "word-break",
		Number:          87,
		Title:           "Word Break",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Hash Table", "String", "Dynamic Programming", "Trie", "Memoization"},
		RelatedChapters: []int{9, 11, 12},
		Description: `Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.`,
		Constraints: []string{
			"1 <= s.length <= 300",
			"1 <= wordDict.length <= 1000",
			"1 <= wordDict[i].length <= 20",
			"s and wordDict[i] consist of only lowercase English letters",
			"All the strings of wordDict are unique",
		},
		Examples: []Example{
			{Input: `s = "leetcode", wordDict = ["leet","code"]`, Output: "true"},
			{Input: `s = "applepenapple", wordDict = ["apple","pen"]`, Output: "true"},
			{Input: `s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]`, Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "leetcode", "wordDict": []string{"leet", "code"}}, Expected: true},
			{Input: map[string]interface{}{"s": "applepenapple", "wordDict": []string{"apple", "pen"}}, Expected: true},
		},
		TimeComplexity:  "O(n^2 * m)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def wordBreak(s, wordDict):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "DP where dp[i] = True if s[0:i] can be segmented. Check all possible last words ending at position i.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "dp[0] = True (empty string). For each i, check if any word ends at i where dp[i-len(word)] is True."},
			{Level: 3, Type: "code", Content: "word_set = set(wordDict). dp[0] = True. for i in 1..n: for j in 0..i: if dp[j] and s[j:i] in word_set: dp[i] = True."},
		},
		Solution: Solution{
			Code: `def wordBreak(s, wordDict):
    word_set = set(wordDict)
    n = len(s)
    dp = [False] * (n + 1)
    dp[0] = True  # Empty string

    for i in range(1, n + 1):
        for j in range(i):
            if dp[j] and s[j:i] in word_set:
                dp[i] = True
                break

    return dp[n]`,
			Explanation:     "DP: dp[i] = can segment s[0:i]. For each position, check if any valid word ends there with valid prefix.",
			TimeComplexity:  "O(n^2 * m) where m is avg word length",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Base case", Explanation: "Empty string can always be segmented", CodeSnippet: "dp[0] = True", LineStart: 5, LineEnd: 5},
				{Title: "Check all splits", Explanation: "For each j, check if s[j:i] is a word", CodeSnippet: "if dp[j] and s[j:i] in word_set:", LineStart: 9, LineEnd: 11},
				{Title: "Early break", Explanation: "Once we find valid segmentation, no need to continue", CodeSnippet: "break", LineStart: 11, LineEnd: 11},
			},
		},
	},
	{
		ID:              "palindromic-substrings",
		Number:          100,
		Title:           "Palindromic Substrings",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.`,
		Constraints: []string{
			"1 <= s.length <= 1000",
			"s consists of lowercase English letters",
		},
		Examples: []Example{
			{Input: `s = "abc"`, Output: "3", Explanation: "Three palindromic strings: 'a', 'b', 'c'."},
			{Input: `s = "aaa"`, Output: "6", Explanation: "Six palindromic strings: 'a', 'a', 'a', 'aa', 'aa', 'aaa'."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "abc"}, Expected: 3},
			{Input: map[string]interface{}{"s": "aaa"}, Expected: 6},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def countSubstrings(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Expand around center like longest palindromic substring. Count all palindromes found during expansion."},
			{Level: 2, Type: "algorithm", Content: "For each center (single char and between chars), expand outward counting palindromes."},
			{Level: 3, Type: "code", Content: "def expand(l, r): count = 0. while valid and s[l] == s[r]: count += 1, expand. Return count."},
		},
		Solution: Solution{
			Code: `def countSubstrings(s):
    n = len(s)
    count = 0

    def expand(l, r):
        nonlocal count
        while l >= 0 and r < n and s[l] == s[r]:
            count += 1
            l -= 1
            r += 1

    for i in range(n):
        expand(i, i)      # Odd length
        expand(i, i + 1)  # Even length

    return count`,
			Explanation:     "Expand around each center, counting palindromes. Try both odd (single center) and even (double center) palindromes.",
			TimeComplexity:  "O(n^2)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Expand and count", Explanation: "Each valid expansion is a palindrome", CodeSnippet: "while l >= 0 and r < n and s[l] == s[r]:\n    count += 1", LineStart: 7, LineEnd: 8},
				{Title: "Two center types", Explanation: "Odd (single) and even (pair) centers", CodeSnippet: "expand(i, i)\nexpand(i, i + 1)", LineStart: 13, LineEnd: 14},
			},
		},
	},
	{
		ID:              "decode-ways",
		Number:          119,
		Title:           "Decode Ways",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `A message containing letters from A-Z can be encoded into numbers using the following mapping:

'A' -> "1"
'B' -> "2"
...
'Z' -> "26"

To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways).

Given a string s containing only digits, return the number of ways to decode it.`,
		Constraints: []string{
			"1 <= s.length <= 100",
			"s contains only digits and may contain leading zero(s)",
		},
		Examples: []Example{
			{Input: `s = "12"`, Output: "2", Explanation: "'AB' (1 2) or 'L' (12)."},
			{Input: `s = "226"`, Output: "3", Explanation: "'BZ' (2 26), 'VF' (22 6), or 'BBF' (2 2 6)."},
			{Input: `s = "06"`, Output: "0", Explanation: "'06' cannot be mapped to 'F' due to leading zero."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "12"}, Expected: 2},
			{Input: map[string]interface{}{"s": "226"}, Expected: 3},
			{Input: map[string]interface{}{"s": "06"}, Expected: 0},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def numDecodings(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "DP where dp[i] = ways to decode s[0:i]. Single digit (1-9) or two digits (10-26).", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "dp[i] = dp[i-1] if s[i-1] != '0', plus dp[i-2] if s[i-2:i] is valid (10-26)."},
			{Level: 3, Type: "code", Content: "Use two variables prev1, prev2. Update based on current digit and two-digit number."},
		},
		Solution: Solution{
			Code: `def numDecodings(s):
    if not s or s[0] == '0':
        return 0

    prev2 = 1  # dp[i-2]
    prev1 = 1  # dp[i-1]

    for i in range(1, len(s)):
        current = 0

        # Single digit decode
        if s[i] != '0':
            current += prev1

        # Two digit decode
        two_digit = int(s[i-1:i+1])
        if 10 <= two_digit <= 26:
            current += prev2

        prev2 = prev1
        prev1 = current

    return prev1`,
			Explanation:     "DP with space optimization. Single digit if 1-9, two digits if 10-26. Sum both possibilities.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Single digit", Explanation: "Valid if not '0'", CodeSnippet: "if s[i] != '0':\n    current += prev1", LineStart: 12, LineEnd: 13},
				{Title: "Two digits", Explanation: "Valid if 10-26", CodeSnippet: "if 10 <= two_digit <= 26:\n    current += prev2", LineStart: 16, LineEnd: 18},
			},
		},
	},
	{
		ID:              "maximum-product-subarray",
		Number:          120,
		Title:           "Maximum Product Subarray",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `Given an integer array nums, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.`,
		Constraints: []string{
			"1 <= nums.length <= 2 * 10^4",
			"-10 <= nums[i] <= 10",
			"The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer",
		},
		Examples: []Example{
			{Input: "nums = [2,3,-2,4]", Output: "6", Explanation: "[2,3] has the largest product."},
			{Input: "nums = [-2,0,-1]", Output: "0"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 3, -2, 4}}, Expected: 6},
			{Input: map[string]interface{}{"nums": []int{-2, 0, -1}}, Expected: 0},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def maxProduct(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Track both max and min product ending at current position. Negative * negative = positive.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "cur_max = max(num, cur_max * num, cur_min * num). cur_min similarly. Negative can flip max/min."},
			{Level: 3, Type: "code", Content: "cur_max = cur_min = result = nums[0]. for num in nums[1:]: candidates = (num, cur_max*num, cur_min*num). Update max, min, result."},
		},
		Solution: Solution{
			Code: `def maxProduct(nums):
    cur_max = cur_min = result = nums[0]

    for num in nums[1:]:
        candidates = (num, cur_max * num, cur_min * num)
        cur_max = max(candidates)
        cur_min = min(candidates)
        result = max(result, cur_max)

    return result`,
			Explanation:     "Track max and min products. Negative number can turn min into max. Consider starting fresh or extending.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Three candidates", Explanation: "Start fresh, extend max, or extend min", CodeSnippet: "candidates = (num, cur_max * num, cur_min * num)", LineStart: 5, LineEnd: 5},
				{Title: "Track both", Explanation: "Need min because negative * negative = positive", CodeSnippet: "cur_max = max(candidates)\ncur_min = min(candidates)", LineStart: 6, LineEnd: 7},
			},
		},
	},
	{
		ID:              "partition-equal-subset-sum",
		Number:          121,
		Title:           "Partition Equal Subset Sum",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or false otherwise.`,
		Constraints: []string{
			"1 <= nums.length <= 200",
			"1 <= nums[i] <= 100",
		},
		Examples: []Example{
			{Input: "nums = [1,5,11,5]", Output: "true", Explanation: "[1, 5, 5] and [11]"},
			{Input: "nums = [1,2,3,5]", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 5, 11, 5}}, Expected: true},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 5}}, Expected: false},
		},
		TimeComplexity:  "O(n * sum)",
		SpaceComplexity: "O(sum)",
		StarterCode:     "def canPartition(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "If total is odd, impossible. Otherwise, find subset summing to total/2. This is subset sum problem.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "DP: dp[s] = can we make sum s? For each num, update from right to left."},
			{Level: 3, Type: "code", Content: "target = sum // 2. dp = set([0]). for num: dp |= {s + num for s in dp if s + num <= target}. return target in dp."},
		},
		Solution: Solution{
			Code: `def canPartition(nums):
    total = sum(nums)
    if total % 2 != 0:
        return False

    target = total // 2
    dp = [False] * (target + 1)
    dp[0] = True

    for num in nums:
        for s in range(target, num - 1, -1):
            dp[s] = dp[s] or dp[s - num]

    return dp[target]`,
			Explanation:     "Find subset summing to half of total. If found, other subset also sums to half.",
			TimeComplexity:  "O(n * sum)",
			SpaceComplexity: "O(sum)",
			Walkthrough: []WalkthroughStep{
				{Title: "Check parity", Explanation: "Odd total can't be split equally", CodeSnippet: "if total % 2 != 0:\n    return False", LineStart: 3, LineEnd: 4},
				{Title: "Subset sum DP", Explanation: "Can we make each sum?", CodeSnippet: "dp[s] = dp[s] or dp[s - num]", LineStart: 12, LineEnd: 12},
				{Title: "Right to left", Explanation: "Prevent using same element twice", CodeSnippet: "for s in range(target, num - 1, -1):", LineStart: 11, LineEnd: 11},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, DP1DProblems...)
}
