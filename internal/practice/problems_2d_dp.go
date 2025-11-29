package practice

// DP2DProblems contains all 2d-dp category problems
var DP2DProblems = []*Problem{
	{
		ID:              "longest-common-subsequence",
		Number:          102,
		Title:           "Longest Common Subsequence",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

A common subsequence of two strings is a subsequence that is common to both strings.`,
		Constraints: []string{
			"1 <= text1.length, text2.length <= 1000",
			"text1 and text2 consist of only lowercase English characters",
		},
		Examples: []Example{
			{Input: `text1 = "abcde", text2 = "ace"`, Output: "3", Explanation: "The longest common subsequence is 'ace'."},
			{Input: `text1 = "abc", text2 = "abc"`, Output: "3"},
			{Input: `text1 = "abc", text2 = "def"`, Output: "0"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"text1": "abcde", "text2": "ace"}, Expected: 3},
			{Input: map[string]interface{}{"text1": "abc", "text2": "def"}, Expected: 0},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def longestCommonSubsequence(text1, text2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "2D DP where dp[i][j] = LCS of text1[0:i] and text2[0:j].", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "If chars match: dp[i][j] = dp[i-1][j-1] + 1. Else: dp[i][j] = max(dp[i-1][j], dp[i][j-1])."},
			{Level: 3, Type: "code", Content: "dp = (m+1) x (n+1) zeros. for i, j: if match, diagonal + 1. else max of left, up."},
		},
		Solution: Solution{
			Code: `def longestCommonSubsequence(text1, text2):
    m, n = len(text1), len(text2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]

    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if text1[i - 1] == text2[j - 1]:
                dp[i][j] = dp[i - 1][j - 1] + 1
            else:
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])

    return dp[m][n]`,
			Explanation:     "2D DP: if characters match, extend LCS from diagonal. Otherwise, take max from excluding either character.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Characters match", Explanation: "Extend LCS by 1 from diagonal", CodeSnippet: "dp[i][j] = dp[i - 1][j - 1] + 1", LineStart: 8, LineEnd: 8},
				{Title: "Characters differ", Explanation: "Take max of excluding either char", CodeSnippet: "dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])", LineStart: 10, LineEnd: 10},
			},
		},
	},
	{
		ID:              "edit-distance",
		Number:          103,
		Title:           "Edit Distance",
		Difficulty:      "Hard",
		Category:        "2d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:
- Insert a character
- Delete a character
- Replace a character`,
		Constraints: []string{
			"0 <= word1.length, word2.length <= 500",
			"word1 and word2 consist of lowercase English letters",
		},
		Examples: []Example{
			{Input: `word1 = "horse", word2 = "ros"`, Output: "3", Explanation: "horse -> rorse -> rose -> ros"},
			{Input: `word1 = "intention", word2 = "execution"`, Output: "5"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"word1": "horse", "word2": "ros"}, Expected: 3},
			{Input: map[string]interface{}{"word1": "intention", "word2": "execution"}, Expected: 5},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def minDistance(word1, word2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "2D DP where dp[i][j] = min ops to convert word1[0:i] to word2[0:j].", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "If chars match: dp[i][j] = dp[i-1][j-1]. Else: 1 + min(insert, delete, replace)."},
			{Level: 3, Type: "code", Content: "Base: dp[i][0] = i, dp[0][j] = j. If match, diagonal. Else 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])."},
		},
		Solution: Solution{
			Code: `def minDistance(word1, word2):
    m, n = len(word1), len(word2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]

    # Base cases
    for i in range(m + 1):
        dp[i][0] = i
    for j in range(n + 1):
        dp[0][j] = j

    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if word1[i - 1] == word2[j - 1]:
                dp[i][j] = dp[i - 1][j - 1]
            else:
                dp[i][j] = 1 + min(
                    dp[i - 1][j],      # Delete
                    dp[i][j - 1],      # Insert
                    dp[i - 1][j - 1]   # Replace
                )

    return dp[m][n]`,
			Explanation:     "2D DP: if match, no operation needed. Otherwise, try all three operations and take minimum.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Base cases", Explanation: "Converting to/from empty string", CodeSnippet: "dp[i][0] = i\ndp[0][j] = j", LineStart: 6, LineEnd: 9},
				{Title: "Characters match", Explanation: "No operation needed", CodeSnippet: "dp[i][j] = dp[i - 1][j - 1]", LineStart: 14, LineEnd: 14},
				{Title: "Try all operations", Explanation: "Delete, insert, or replace", CodeSnippet: "dp[i][j] = 1 + min(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1])", LineStart: 16, LineEnd: 20},
			},
		},
	},
	{
		ID:              "target-sum",
		Number:          104,
		Title:           "Target Sum",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Array", "Dynamic Programming", "Backtracking"},
		RelatedChapters: []int{4, 9, 12},
		Description: `You are given an integer array nums and an integer target.

You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

Return the number of different expressions that you can build, which evaluates to target.`,
		Constraints: []string{
			"1 <= nums.length <= 20",
			"0 <= nums[i] <= 1000",
			"0 <= sum(nums[i]) <= 1000",
			"-1000 <= target <= 1000",
		},
		Examples: []Example{
			{Input: "nums = [1,1,1,1,1], target = 3", Output: "5", Explanation: "5 ways: -1+1+1+1+1, +1-1+1+1+1, +1+1-1+1+1, +1+1+1-1+1, +1+1+1+1-1"},
			{Input: "nums = [1], target = 1", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 1, 1}, "target": 3}, Expected: 5},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 1}, Expected: 1},
		},
		TimeComplexity:  "O(n * sum)",
		SpaceComplexity: "O(sum)",
		StarterCode:     "def findTargetSumWays(nums, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Convert to subset sum: find subset P such that sum(P) - sum(rest) = target. This means sum(P) = (total + target) / 2.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "DP: count ways to reach each sum. dp[s] = number of ways to achieve sum s."},
			{Level: 3, Type: "code", Content: "target_sum = (total + target) // 2. dp[0] = 1. for num: for s from target_sum down to num: dp[s] += dp[s - num]."},
		},
		Solution: Solution{
			Code: `def findTargetSumWays(nums, target):
    total = sum(nums)

    # Check if solution is possible
    if (total + target) % 2 != 0 or abs(target) > total:
        return 0

    target_sum = (total + target) // 2
    dp = [0] * (target_sum + 1)
    dp[0] = 1

    for num in nums:
        for s in range(target_sum, num - 1, -1):
            dp[s] += dp[s - num]

    return dp[target_sum]`,
			Explanation:     "Convert to subset sum problem. Find subsets that sum to (total + target) / 2. Use 1D DP counting ways.",
			TimeComplexity:  "O(n * sum)",
			SpaceComplexity: "O(sum)",
			Walkthrough: []WalkthroughStep{
				{Title: "Convert problem", Explanation: "Find subset summing to (total + target) / 2", CodeSnippet: "target_sum = (total + target) // 2", LineStart: 8, LineEnd: 8},
				{Title: "Count ways", Explanation: "Each number can be included or not", CodeSnippet: "dp[s] += dp[s - num]", LineStart: 14, LineEnd: 14},
			},
		},
	},
	{
		ID:              "coin-change-ii",
		Number:          105,
		Title:           "Coin Change II",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the number of combinations that make up that amount. If that amount of money cannot be made up by any combination of the coins, return 0.

You may assume that you have an infinite number of each kind of coin.`,
		Constraints: []string{
			"1 <= coins.length <= 300",
			"1 <= coins[i] <= 5000",
			"All the values of coins are unique",
			"0 <= amount <= 5000",
		},
		Examples: []Example{
			{Input: "amount = 5, coins = [1,2,5]", Output: "4", Explanation: "4 ways: 5, 2+2+1, 2+1+1+1, 1+1+1+1+1"},
			{Input: "amount = 3, coins = [2]", Output: "0"},
			{Input: "amount = 10, coins = [10]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"amount": 5, "coins": []int{1, 2, 5}}, Expected: 4},
			{Input: map[string]interface{}{"amount": 3, "coins": []int{2}}, Expected: 0},
		},
		TimeComplexity:  "O(n * amount)",
		SpaceComplexity: "O(amount)",
		StarterCode:     "def change(amount, coins):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "DP where dp[i] = number of ways to make amount i. Process coins one by one to avoid counting permutations.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "For each coin, update dp[amount] by adding dp[amount - coin]. Order of loops matters for combinations vs permutations."},
			{Level: 3, Type: "code", Content: "dp[0] = 1. for coin in coins: for amount from coin to target: dp[amount] += dp[amount - coin]."},
		},
		Solution: Solution{
			Code: `def change(amount, coins):
    dp = [0] * (amount + 1)
    dp[0] = 1  # One way to make amount 0

    for coin in coins:
        for a in range(coin, amount + 1):
            dp[a] += dp[a - coin]

    return dp[amount]`,
			Explanation:     "DP counts combinations (not permutations) by processing one coin at a time. dp[a] = ways to make amount a.",
			TimeComplexity:  "O(n * amount)",
			SpaceComplexity: "O(amount)",
			Walkthrough: []WalkthroughStep{
				{Title: "Base case", Explanation: "One way to make amount 0 (use no coins)", CodeSnippet: "dp[0] = 1", LineStart: 3, LineEnd: 3},
				{Title: "Process each coin", Explanation: "Outer loop on coins avoids counting permutations", CodeSnippet: "for coin in coins:", LineStart: 5, LineEnd: 5},
				{Title: "Count ways", Explanation: "Add ways using this coin", CodeSnippet: "dp[a] += dp[a - coin]", LineStart: 7, LineEnd: 7},
			},
		},
	},
	{
		ID:              "interleaving-string",
		Number:          127,
		Title:           "Interleaving String",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.

An interleaving of two strings s and t is a configuration where s and t are divided into n and m substrings respectively, such that the interleaving is s1 + t1 + s2 + t2 + ... or t1 + s1 + t2 + s2 + ...`,
		Constraints: []string{
			"0 <= s1.length, s2.length <= 100",
			"0 <= s3.length <= 200",
			"s1, s2, and s3 consist of lowercase English letters",
		},
		Examples: []Example{
			{Input: `s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"`, Output: "true"},
			{Input: `s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"`, Output: "false"},
			{Input: `s1 = "", s2 = "", s3 = ""`, Output: "true"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s1": "aabcc", "s2": "dbbca", "s3": "aadbbcbcac"}, Expected: true},
			{Input: map[string]interface{}{"s1": "aabcc", "s2": "dbbca", "s3": "aadbbbaccc"}, Expected: false},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def isInterleave(s1, s2, s3):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "DP where dp[i][j] = can first i chars of s1 and first j chars of s2 form first i+j chars of s3.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "dp[i][j] = (dp[i-1][j] and s1[i-1]==s3[i+j-1]) or (dp[i][j-1] and s2[j-1]==s3[i+j-1])"},
			{Level: 3, Type: "code", Content: "Can optimize to 1D: dp[j] = dp[j] and s1[i-1]==s3[i+j-1] or dp[j-1] and s2[j-1]==s3[i+j-1]"},
		},
		Solution: Solution{
			Code: `def isInterleave(s1, s2, s3):
    m, n = len(s1), len(s2)
    if m + n != len(s3):
        return False

    dp = [False] * (n + 1)
    dp[0] = True

    for j in range(1, n + 1):
        dp[j] = dp[j - 1] and s2[j - 1] == s3[j - 1]

    for i in range(1, m + 1):
        dp[0] = dp[0] and s1[i - 1] == s3[i - 1]
        for j in range(1, n + 1):
            dp[j] = (dp[j] and s1[i - 1] == s3[i + j - 1]) or \
                    (dp[j - 1] and s2[j - 1] == s3[i + j - 1])

    return dp[n]`,
			Explanation:     "DP tracks if prefixes of s1 and s2 can interleave to form prefix of s3.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Length check", Explanation: "Must have m + n = len(s3)", CodeSnippet: "if m + n != len(s3):\n    return False", LineStart: 3, LineEnd: 4},
				{Title: "Two choices", Explanation: "Take from s1 (dp[j]) or s2 (dp[j-1])", CodeSnippet: "(dp[j] and s1[i - 1] == s3[i + j - 1]) or (dp[j - 1] and s2[j - 1] == s3[i + j - 1])", LineStart: 15, LineEnd: 16},
			},
		},
	},
	{
		ID:              "longest-increasing-path-in-matrix",
		Number:          128,
		Title:           "Longest Increasing Path in a Matrix",
		Difficulty:      "Hard",
		Category:        "2d-dp",
		Tags:            []string{"Matrix", "DFS", "Dynamic Programming", "Memoization"},
		RelatedChapters: []int{6, 9, 12},
		Description: `Given an m x n integers matrix, return the length of the longest increasing path in matrix.

From each cell, you can either move in four directions: left, right, up, or down. You may not move diagonally or move outside the boundary.`,
		Constraints: []string{
			"m == matrix.length",
			"n == matrix[i].length",
			"1 <= m, n <= 200",
			"0 <= matrix[i][j] <= 2^31 - 1",
		},
		Examples: []Example{
			{Input: "matrix = [[9,9,4],[6,6,8],[2,1,1]]", Output: "4", Explanation: "Longest path is [1, 2, 6, 9]."},
			{Input: "matrix = [[3,4,5],[3,2,6],[2,2,1]]", Output: "4", Explanation: "Longest path is [3, 4, 5, 6]."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}}, Expected: 4},
			{Input: map[string]interface{}{"matrix": [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}}, Expected: 4},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def longestIncreasingPath(matrix):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "DFS with memoization. Each cell's answer depends only on larger neighbors.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "memo[i][j] = 1 + max(dfs(neighbor) for larger neighbors). No cycles due to strictly increasing."},
			{Level: 3, Type: "code", Content: "Use @lru_cache. dfs(r,c): return 1 + max(dfs(nr,nc) for valid larger neighbors, default 0)."},
		},
		Solution: Solution{
			Code: `def longestIncreasingPath(matrix):
    from functools import lru_cache

    m, n = len(matrix), len(matrix[0])

    @lru_cache(maxsize=None)
    def dfs(r, c):
        val = matrix[r][c]
        result = 1
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nr, nc = r + dr, c + dc
            if 0 <= nr < m and 0 <= nc < n and matrix[nr][nc] > val:
                result = max(result, 1 + dfs(nr, nc))
        return result

    return max(dfs(r, c) for r in range(m) for c in range(n))`,
			Explanation:     "DFS from each cell, memoize results. No cycle possible since path is strictly increasing.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Memoization", Explanation: "Cache results to avoid recomputation", CodeSnippet: "@lru_cache(maxsize=None)", LineStart: 6, LineEnd: 6},
				{Title: "Only larger neighbors", Explanation: "Path must be strictly increasing", CodeSnippet: "if 0 <= nr < m and 0 <= nc < n and matrix[nr][nc] > val:", LineStart: 12, LineEnd: 12},
				{Title: "Try all starts", Explanation: "Longest path could start anywhere", CodeSnippet: "max(dfs(r, c) for r in range(m) for c in range(n))", LineStart: 16, LineEnd: 16},
			},
		},
	},
	{
		ID:              "distinct-subsequences",
		Number:          129,
		Title:           "Distinct Subsequences",
		Difficulty:      "Hard",
		Category:        "2d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `Given two strings s and t, return the number of distinct subsequences of s which equals t.

The test cases are generated so that the answer fits on a 32-bit signed integer.`,
		Constraints: []string{
			"1 <= s.length, t.length <= 1000",
			"s and t consist of English letters",
		},
		Examples: []Example{
			{Input: `s = "rabbbit", t = "rabbit"`, Output: "3"},
			{Input: `s = "babgbag", t = "bag"`, Output: "5"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "rabbbit", "t": "rabbit"}, Expected: 3},
			{Input: map[string]interface{}{"s": "babgbag", "t": "bag"}, Expected: 5},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def numDistinct(s, t):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "DP where dp[i][j] = number of ways to form t[0:j] from s[0:i].", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "If s[i]==t[j]: dp[i][j] = dp[i-1][j-1] + dp[i-1][j]. Else: dp[i][j] = dp[i-1][j]."},
			{Level: 3, Type: "code", Content: "Optimize to 1D going right to left. dp[j] += dp[j-1] if s[i]==t[j]."},
		},
		Solution: Solution{
			Code: `def numDistinct(s, t):
    m, n = len(s), len(t)
    dp = [0] * (n + 1)
    dp[0] = 1  # Empty t can be formed from any prefix

    for i in range(1, m + 1):
        for j in range(min(i, n), 0, -1):
            if s[i - 1] == t[j - 1]:
                dp[j] += dp[j - 1]

    return dp[n]`,
			Explanation:     "DP counts ways to select characters from s to form t. Update right to left to use previous row values.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Base case", Explanation: "Empty t formed in one way (select nothing)", CodeSnippet: "dp[0] = 1", LineStart: 4, LineEnd: 4},
				{Title: "Match case", Explanation: "Use char (dp[j-1]) or skip (dp[j])", CodeSnippet: "if s[i - 1] == t[j - 1]:\n    dp[j] += dp[j - 1]", LineStart: 8, LineEnd: 9},
				{Title: "Right to left", Explanation: "Preserve previous row values", CodeSnippet: "for j in range(min(i, n), 0, -1):", LineStart: 7, LineEnd: 7},
			},
		},
	},
	{
		ID:              "burst-balloons",
		Number:          130,
		Title:           "Burst Balloons",
		Difficulty:      "Hard",
		Category:        "2d-dp",
		Tags:            []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `You are given n balloons, indexed from 0 to n - 1. Each balloon is painted with a number on it represented by an array nums. You are asked to burst all the balloons.

If you burst the ith balloon, you will get nums[i - 1] * nums[i] * nums[i + 1] coins. If i - 1 or i + 1 goes out of bounds, treat it as having a balloon with value 1.

Return the maximum coins you can collect by bursting the balloons wisely.`,
		Constraints: []string{
			"n == nums.length",
			"1 <= n <= 300",
			"0 <= nums[i] <= 100",
		},
		Examples: []Example{
			{Input: "nums = [3,1,5,8]", Output: "167", Explanation: "Burst 1, then 5, then 3, then 8. Coins = 3*1*5 + 3*5*8 + 1*3*8 + 1*8*1 = 167."},
			{Input: "nums = [1,5]", Output: "10"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{3, 1, 5, 8}}, Expected: 167},
			{Input: map[string]interface{}{"nums": []int{1, 5}}, Expected: 10},
		},
		TimeComplexity:  "O(n^3)",
		SpaceComplexity: "O(n^2)",
		StarterCode:     "def maxCoins(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Think in reverse: which balloon to burst LAST in a range. Interval DP.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "dp[l][r] = max coins for bursting all balloons in range (l,r). Last balloon k: dp[l][r] = dp[l][k] + nums[l]*nums[k]*nums[r] + dp[k][r]."},
			{Level: 3, Type: "code", Content: "Add virtual 1s at ends. for length 2 to n+2: for left: for k in (left+1, right): dp[l][r] = max over k."},
		},
		Solution: Solution{
			Code: `def maxCoins(nums):
    nums = [1] + nums + [1]
    n = len(nums)
    dp = [[0] * n for _ in range(n)]

    for length in range(2, n):
        for left in range(n - length):
            right = left + length
            for k in range(left + 1, right):
                dp[left][right] = max(
                    dp[left][right],
                    dp[left][k] + nums[left] * nums[k] * nums[right] + dp[k][right]
                )

    return dp[0][n - 1]`,
			Explanation:     "Interval DP. Think of k as last balloon to burst in range (left, right).",
			TimeComplexity:  "O(n^3)",
			SpaceComplexity: "O(n^2)",
			Walkthrough: []WalkthroughStep{
				{Title: "Add boundaries", Explanation: "Virtual balloons with value 1", CodeSnippet: "nums = [1] + nums + [1]", LineStart: 2, LineEnd: 2},
				{Title: "Last balloon k", Explanation: "k is last to burst, neighbors are left and right", CodeSnippet: "nums[left] * nums[k] * nums[right]", LineStart: 12, LineEnd: 12},
				{Title: "Combine subproblems", Explanation: "Solve left and right independently", CodeSnippet: "dp[left][k] + ... + dp[k][right]", LineStart: 11, LineEnd: 12},
			},
		},
	},
	{
		ID:              "regular-expression-matching",
		Number:          131,
		Title:           "Regular Expression Matching",
		Difficulty:      "Hard",
		Category:        "2d-dp",
		Tags:            []string{"String", "Dynamic Programming", "Recursion"},
		RelatedChapters: []int{9, 12},
		Description: `Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.
'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).`,
		Constraints: []string{
			"1 <= s.length <= 20",
			"1 <= p.length <= 20",
			"s contains only lowercase English letters",
			"p contains only lowercase English letters, '.', and '*'",
			"It is guaranteed for each '*', there is a previous valid character to match",
		},
		Examples: []Example{
			{Input: `s = "aa", p = "a"`, Output: "false"},
			{Input: `s = "aa", p = "a*"`, Output: "true"},
			{Input: `s = "ab", p = ".*"`, Output: "true"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "aa", "p": "a"}, Expected: false},
			{Input: map[string]interface{}{"s": "aa", "p": "a*"}, Expected: true},
			{Input: map[string]interface{}{"s": "ab", "p": ".*"}, Expected: true},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def isMatch(s, p):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "DP where dp[i][j] = does s[0:i] match p[0:j]. Handle '*' specially.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "If p[j-1]=='*': dp[i][j] = dp[i][j-2] (skip x*) or (match and dp[i-1][j]). Else: dp[i][j] = match and dp[i-1][j-1]."},
			{Level: 3, Type: "code", Content: "match = p[j-1]=='.' or p[j-1]==s[i-1]. Handle '*' case: zero matches (dp[i][j-2]) or one+ matches."},
		},
		Solution: Solution{
			Code: `def isMatch(s, p):
    m, n = len(s), len(p)
    dp = [[False] * (n + 1) for _ in range(m + 1)]
    dp[0][0] = True

    # Handle patterns like a*, a*b*, a*b*c* matching empty string
    for j in range(2, n + 1):
        if p[j - 1] == '*':
            dp[0][j] = dp[0][j - 2]

    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if p[j - 1] == '*':
                # Zero matches or one+ matches
                dp[i][j] = dp[i][j - 2] or \
                           (dp[i - 1][j] and (p[j - 2] == '.' or p[j - 2] == s[i - 1]))
            else:
                # Direct match
                dp[i][j] = dp[i - 1][j - 1] and (p[j - 1] == '.' or p[j - 1] == s[i - 1])

    return dp[m][n]`,
			Explanation:     "DP with special handling for '*'. Either skip pattern (zero matches) or consume string char.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Empty string patterns", Explanation: "a* can match empty string", CodeSnippet: "if p[j - 1] == '*':\n    dp[0][j] = dp[0][j - 2]", LineStart: 8, LineEnd: 9},
				{Title: "Star cases", Explanation: "Zero matches (dp[i][j-2]) or extend match", CodeSnippet: "dp[i][j] = dp[i][j - 2] or (dp[i - 1][j] and ...)", LineStart: 15, LineEnd: 16},
				{Title: "Dot wildcard", Explanation: "'.' matches any character", CodeSnippet: "p[j - 1] == '.' or p[j - 1] == s[i - 1]", LineStart: 19, LineEnd: 19},
			},
		},
	},
	{
		ID:              "best-time-to-buy-sell-stock-cooldown",
		Number:          142,
		Title:           "Best Time to Buy and Sell Stock with Cooldown",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{9, 12},
		Description: `You are given an array prices where prices[i] is the price of a given stock on the ith day.

Find the maximum profit you can achieve. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:

After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).`,
		Constraints: []string{
			"1 <= prices.length <= 5000",
			"0 <= prices[i] <= 1000",
		},
		Examples: []Example{
			{Input: "prices = [1,2,3,0,2]", Output: "3", Explanation: "Buy, sell, cooldown, buy, sell"},
			{Input: "prices = [1]", Output: "0"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"prices": []int{1, 2, 3, 0, 2}}, Expected: 3},
			{Input: map[string]interface{}{"prices": []int{1}}, Expected: 0},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def maxProfit(prices):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "State machine DP: held, sold, cooldown states.", ChapterRef: 9},
			{Level: 2, Type: "algorithm", Content: "held[i] = max(held[i-1], cooldown[i-1]-price). sold[i] = held[i-1]+price. cooldown[i] = max(cooldown[i-1], sold[i-1])."},
			{Level: 3, Type: "code", Content: "held, sold, cooldown = -inf, 0, 0. for price: held, sold, cooldown = max(held, cooldown-price), held+price, max(cooldown, sold)."},
		},
		Solution: Solution{
			Code: `def maxProfit(prices):
    if not prices:
        return 0

    held = float('-inf')
    sold = 0
    cooldown = 0

    for price in prices:
        prev_held = held
        held = max(held, cooldown - price)
        cooldown = max(cooldown, sold)
        sold = prev_held + price

    return max(sold, cooldown)`,
			Explanation:     "Three states: holding stock, just sold, cooldown. Track max profit in each state.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Held state", Explanation: "Keep holding or buy from cooldown", CodeSnippet: "held = max(held, cooldown - price)", LineStart: 11, LineEnd: 11},
				{Title: "Sold state", Explanation: "Sell what we're holding", CodeSnippet: "sold = prev_held + price", LineStart: 13, LineEnd: 13},
				{Title: "Cooldown state", Explanation: "Stay in cooldown or transition from sold", CodeSnippet: "cooldown = max(cooldown, sold)", LineStart: 12, LineEnd: 12},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, DP2DProblems...)
}
