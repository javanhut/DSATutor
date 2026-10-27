package practice

// DP2DProblems is the 2d-dp ladder taught by Chapter 16. Every solution is
// built with the Chapter 15 recipe (Define the state, Write the recurrence,
// Set the base cases, Choose the evaluation order, Locate the answer,
// Optimise space); here the state has two indices, so the table is a grid.
// The main solutions keep the full 2D table so learners can see it; the
// rolling-array and top-down versions are in Alternatives.
var DP2DProblems = []*Problem{
	{
		ID:              "unique-paths",
		Number:          62,
		Order:           1,
		Title:           "Unique Paths",
		Slug:            "unique-paths",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Math", "Dynamic Programming", "Combinatorics", "Grid"},
		RelatedChapters: []int{16, 15},
		Description: `A robot stands in the top-left cell of a grid with m rows and n columns. It wants to reach the bottom-right cell. On every move it goes exactly one cell **right** or one cell **down**; it can never move up, left or diagonally, and it can't leave the grid.

Return the number of different routes the robot can take from the top-left cell to the bottom-right cell. Two routes are different if their sequences of moves differ.

If the grid is a single cell (m = n = 1), the robot is already at the finish, and the empty route counts as one route.`,
		Constraints: []string{
			"1 <= m, n <= 100",
			"The answer is at most 2 * 10^9 for every test",
		},
		Examples: []Example{
			{Input: "m = 3, n = 7", Output: "28", Explanation: "Every route makes 2 down moves and 6 right moves in some order. The table below shows how many routes reach each cell; the bottom-right cell holds 28."},
			{Input: "m = 3, n = 2", Output: "3", Explanation: "Right-Down-Down, Down-Right-Down and Down-Down-Right."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"m": 3, "n": 7}, Expected: 28, Note: "example: 3 x 7 grid"},
			{Input: map[string]interface{}{"m": 3, "n": 2}, Expected: 3, Note: "example: 3 x 2 grid"},
			{Input: map[string]interface{}{"m": 1, "n": 1}, Expected: 1, Note: "1 x 1: start is the finish, one (empty) path"},
			{Input: map[string]interface{}{"m": 1, "n": 10}, Expected: 1, Hidden: true, Note: "single row: only moves right"},
			{Input: map[string]interface{}{"m": 10, "n": 1}, Expected: 1, Hidden: true, Note: "single column: only moves down"},
			{Input: map[string]interface{}{"m": 2, "n": 2}, Expected: 2, Hidden: true, Note: "2 x 2: right-down or down-right"},
			{Input: map[string]interface{}{"m": 7, "n": 3}, Expected: 28, Hidden: true, Note: "7 x 3 gives the same count as 3 x 7"},
			{Input: map[string]interface{}{"m": 3, "n": 3}, Expected: 6, Hidden: true, Note: "3 x 3 square"},
			{Input: map[string]interface{}{"m": 23, "n": 12}, Expected: 193536720, Hidden: true, Note: "large grid: recursion without a table times out"},
			{Input: map[string]interface{}{"m": 16, "n": 16}, Expected: 155117520, Hidden: true, Note: "16 x 16: 155 million paths, too many to walk one by one"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def uniquePaths(m, n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Pick any cell in the middle of the grid. The robot's last move into that cell came from somewhere. Which cells could it have come from, and how does knowing their route counts help?"},
			{Level: 2, Type: "insight", Content: "Every route into cell (r, c) arrives either from the cell above, (r - 1, c), or from the cell on the left, (r, c - 1), and those two groups of routes never overlap. So routes(r, c) = routes(r - 1, c) + routes(r, c - 1). That is a 2D dynamic programming table: one number per cell, filled row by row. Chapter 16 draws this exact table.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Define the state: paths[r][c] = number of routes from (0, 0) to (r, c).\n2. Set the base cases: every cell of the top row and of the left column has exactly 1 route (a straight line).\n3. Write the recurrence: paths[r][c] = paths[r - 1][c] + paths[r][c - 1].\n4. Choose the evaluation order: row by row, left to right, so the cell above and the cell to the left are always done first.\n5. Locate the answer: the bottom-right cell, paths[m - 1][n - 1]."},
			{Level: 4, Type: "skeleton", Content: "def uniquePaths(m, n):\n    paths = [[0] * n for _ in range(m)]\n    for c in range(n):\n        paths[0][c] = ____\n    for r in range(m):\n        paths[r][0] = ____\n    for r in range(1, m):\n        for c in range(1, n):\n            paths[r][c] = ____\n    return ____"},
		},
		Solution: Solution{
			Code: `def uniquePaths(m, n):
    # paths[r][c] = number of right/down routes from (0, 0) to (r, c)
    paths = [[0] * n for _ in range(m)]
    for c in range(n):
        paths[0][c] = 1   # top row: the only way in is from the left
    for r in range(m):
        paths[r][0] = 1   # left column: the only way in is from above
    for r in range(1, m):
        for c in range(1, n):
            paths[r][c] = paths[r - 1][c] + paths[r][c - 1]
    return paths[m - 1][n - 1]`,
			Explanation: "This is the gentlest grid DP and the first rung of the Chapter 16 ladder. It is Climbing Stairs from Chapter 15 with a second index: there, ways(n) added the counts for the two possible last steps; here paths[r][c] adds the counts for the two possible last moves (from above, from the left). We fill an m x n table row by row and read the answer from its bottom-right corner.",
			BruteForce:  "Recursively try both moves from every cell: count(r, c) = count(r - 1, c) + count(r, c - 1), with no table. The work equals the number of routes, which is C(m + n - 2, m - 1): 28 for 3 x 7, but 155,117,520 for 16 x 16 and about 2.3 x 10^58 for 100 x 100. It recomputes the same cell over and over: in a 16 x 16 grid the cell (1, 1) is reached by millions of different call chains, and each one recounts it from scratch.",
			KeyInsight:  "The number of routes into a cell doesn't depend on how the robot later leaves it, so each cell's count can be computed once and reused by the two cells that depend on it (the one below and the one to the right). There are only m * n cells, so the table does m * n additions instead of one step per route.",
			WhyItWorks:  "State: paths[r][c] = number of right/down routes from (0, 0) to (r, c). Recurrence: a route's last move into (r, c) is either 'down from (r - 1, c)' or 'right from (r, c - 1)'. Every route has exactly one last move, so these two groups cover all routes (complete) and no route is in both (disjoint); adding their sizes counts each route exactly once. Base cases: a cell in row 0 can only be entered from the left and a cell in column 0 only from above, so each has exactly one route (straight along the edge). Order: row by row, left to right, guarantees that (r - 1, c) and (r, c - 1) are final before (r, c) reads them. For m = 3, n = 3 the table is\n\n    1 1 1\n    1 2 3\n    1 3 6\n\nand the answer 6 is in the corner. (It also equals C(4, 2) = 6: choose which 2 of the 4 moves are 'down'.)",
			CommonMistakes: []string{
				"Starting the double loop at r = 0 and c = 0 and reading paths[r - 1][c]. In Python paths[-1][c] silently reads the LAST row, so there is no error, just a wrong count. Fill the top row and left column first (the base cases) and start the loops at 1.",
				"Initialising the whole table to 0 and setting only paths[0][0] = 1, then skipping row 0 and column 0 in the loop. The edges stay 0 and every answer comes out 0. Either set both edges to 1 or let the loop handle them with bounds checks.",
				"Mixing up m and n: m is the number of rows and n the number of columns. Building [[0] * m for _ in range(n)] gives an n x m table and paths[m - 1][n - 1] goes out of range for 3 x 7. The 7 x 3 test case gives the same count, so check a non-square grid while debugging.",
			},
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "One number per cell: how many routes reach it. The table has m rows and n columns, the same shape as the grid.", CodeSnippet: "paths = [[0] * n for _ in range(m)]", LineStart: 2, LineEnd: 3},
				{Title: "Set the base cases", Explanation: "Cells on the top edge can only be entered from the left, cells on the left edge only from above, so each has exactly one route. These are the cells the recurrence can't compute because a neighbour is missing.", CodeSnippet: "paths[0][c] = 1\npaths[r][0] = 1", LineStart: 4, LineEnd: 7},
				{Title: "Choose the evaluation order", Explanation: "Row by row, left to right, starting at (1, 1). When we reach (r, c), the row above is finished and the cell on the left was just computed.", CodeSnippet: "for r in range(1, m):\n    for c in range(1, n):", LineStart: 8, LineEnd: 9},
				{Title: "Write the recurrence", Explanation: "Last move from above or from the left; the two groups of routes are disjoint, so we add. For 3 x 7 the second row becomes 1 2 3 4 5 6 7 and the third 1 3 6 10 15 21 28.", CodeSnippet: "paths[r][c] = paths[r - 1][c] + paths[r][c - 1]", LineStart: 10, LineEnd: 10},
				{Title: "Locate the answer", Explanation: "The finish is the bottom-right cell, so its count is the answer. For a 1 x 1 grid that cell is also the start, already set to 1 by the base cases.", CodeSnippet: "return paths[m - 1][n - 1]", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Optimise space: one rolling row",
					Code: `def uniquePaths(m, n):
    row = [1] * n               # row 0: every cell has one path
    for _ in range(1, m):
        for c in range(1, n):
            # before the update row[c] is the cell above; row[c - 1] is already this row
            row[c] = row[c] + row[c - 1]
    return row[n - 1]`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Each row only reads the row above, so keep one row and update it in place. Before row[c] is overwritten it still holds the cell above; row[c - 1] has already been updated, so it is the cell to the left in the current row. Same additions as the table, n numbers of memory instead of m * n. Write the 2D version first and shrink it afterwards.",
				},
				{
					Name: "Top-down: the brute-force recursion plus a cache",
					Code: `from functools import lru_cache

def uniquePaths(m, n):
    @lru_cache(maxsize=None)    # delete this line and it is the exponential brute force
    def count(r, c):
        if r == 0 or c == 0:
            return 1
        return count(r - 1, c) + count(r, c - 1)

    return count(m - 1, n - 1)`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "count(r, c) is the brute-force recursion: routes into (r, c) = routes into the cell above + routes into the cell on the left. Without @lru_cache it makes one call per route (155 million for 16 x 16, which times out). With the cache each (r, c) is computed once, so it does the same m * n work as the table, in whatever order the recursion asks for cells.",
				},
				{
					Name:            "Combinatorics (no DP)",
					Code:            "import math\n\ndef uniquePaths(m, n):\n    # choose which m - 1 of the m + n - 2 moves go down\n    return math.comb(m + n - 2, m - 1)",
					TimeComplexity:  "O(min(m, n))",
					SpaceComplexity: "O(1)",
					Explanation:     "Every route is a sequence of m - 1 downs and n - 1 rights, and every such sequence is a route, so the count is the binomial coefficient C(m + n - 2, m - 1). Neat, but it only works because this grid has no obstacles or costs; the table version adapts to Unique Paths II (blocked cells) or Minimum Path Sum by changing one line.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it correct to ADD the two neighbours' counts rather than take their maximum?", Answer: "We are counting routes, and every route into (r, c) has exactly one last move: from above or from the left. The routes split into two disjoint groups, so the total is the sum of the group sizes. Max would be right for an optimisation question (the most coins collected along a route), where we keep only the best option, not count all of them."},
			{Prompt: "What changes if some cells are blocked (the robot can't enter them)?", Answer: "Set paths[r][c] = 0 for a blocked cell, because no route ends there, and keep the same recurrence for the others. The base cases also change: a top-row cell has 1 route only if no blocked cell lies to its left, so it is easier to start from paths[0][0] = 1 and let the loop add 'above' and 'left' with bounds checks. The combinatorics formula no longer works, which is why the table is the more useful tool."},
			{Prompt: "The rolling-row version writes row[c] = row[c] + row[c - 1]. Why does it read the right values without a second array?", Answer: "When we are about to update row[c] in row r, row[c] still holds the value from row r - 1 (the cell above), because nothing has overwritten it yet in this row. row[c - 1] was updated a moment ago, so it already holds row r's value (the cell on the left). Those are exactly the two cells the recurrence needs, so updating left to right in place is safe."},
		},
	},
	{
		ID:              "longest-common-subsequence",
		Number:          1143,
		Order:           2,
		Title:           "Longest Common Subsequence",
		Slug:            "longest-common-subsequence",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{16, 15},
		Description: `You are given two strings, text1 and text2. Return the length of the longest string that is a **subsequence of both**. If they share no letter at all, return 0.

A subsequence of a string is what remains after deleting any number of its characters (possibly none) without changing the order of the rest. For example, "ace" is a subsequence of "abcde" (delete b and d), but "aec" is not, because e comes after c in "abcde".

The common subsequence does not have to be contiguous in either string.`,
		Constraints: []string{
			"1 <= text1.length, text2.length <= 1000",
			"text1 and text2 contain only lowercase English letters",
		},
		Examples: []Example{
			{Input: `text1 = "abcde", text2 = "ace"`, Output: "3", Explanation: "\"ace\" is a subsequence of both strings, and no common subsequence is longer."},
			{Input: `text1 = "abc", text2 = "abc"`, Output: "3", Explanation: "The strings are equal, so the whole string is common."},
			{Input: `text1 = "abc", text2 = "def"`, Output: "0", Explanation: "No letter appears in both."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"text1": "abcde", "text2": "ace"}, Expected: 3, Note: "example: 'ace' is common"},
			{Input: map[string]interface{}{"text1": "abc", "text2": "abc"}, Expected: 3, Note: "identical strings: the whole string"},
			{Input: map[string]interface{}{"text1": "abc", "text2": "def"}, Expected: 0, Note: "no letter in common"},
			{Input: map[string]interface{}{"text1": "a", "text2": "a"}, Expected: 1, Hidden: true, Note: "single equal letters"},
			{Input: map[string]interface{}{"text1": "a", "text2": "b"}, Expected: 0, Hidden: true, Note: "single different letters"},
			{Input: map[string]interface{}{"text1": "abc", "text2": "cab"}, Expected: 2, Hidden: true, Note: "order matters: 'ab' (not 'abc') is common"},
			{Input: map[string]interface{}{"text1": "cab", "text2": "abc"}, Expected: 2, Hidden: true, Note: "greedy first-match pairs the c's and finds only 1"},
			{Input: map[string]interface{}{"text1": "bl", "text2": "yby"}, Expected: 1, Hidden: true, Note: "length mismatch, one shared letter"},
			{Input: map[string]interface{}{"text1": "ezupkr", "text2": "ubmrapg"}, Expected: 2, Hidden: true, Note: "scattered matches 'ur' or 'up'"},
			{Input: map[string]interface{}{"text1": "abcba", "text2": "abcbcba"}, Expected: 5, Hidden: true, Note: "the shorter string is a subsequence of the longer"},
			{Input: map[string]interface{}{"text1": "pmjghexybyrgzczy", "text2": "hafcdqbgncrcbihkd"}, Expected: 4, Hidden: true, Note: "unrelated words: greedy first-match finds 2, the answer is 4"},
			{Input: map[string]interface{}{"text1": dp2dLetters(300, "abcd", 1), "text2": dp2dLetters(250, "abcd", 2)}, Expected: 175, Hidden: true, Note: "300 x 250 letters: needs the table, not plain recursion"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def longestCommonSubsequence(text1, text2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Look only at the LAST letter of each string. If the two last letters are equal, should that letter be part of the answer? And if they are different, at least one of them can't be the end of a common subsequence. Which one?"},
			{Level: 2, Type: "insight", Content: "Answer the question for every pair of prefixes: lcs[i][j] = LCS length of text1[:i] and text2[:j]. If text1[i - 1] == text2[j - 1], keep both letters: 1 + lcs[i - 1][j - 1]. Otherwise one of the two letters is useless, but you don't know which, so try skipping each side and keep the better: max(lcs[i - 1][j], lcs[i][j - 1]). This is the two-string alignment pattern from Chapter 16: match, or skip which side.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Define the state: lcs[i][j] = LCS length of the first i letters of text1 and the first j letters of text2; the table is (m + 1) x (n + 1).\n2. Set the base cases: row 0 and column 0 are 0, because an empty prefix has no common subsequence.\n3. Write the recurrence: match -> diagonal + 1; mismatch -> max(up, left).\n4. Choose the evaluation order: i from 1 to m, j from 1 to n, so up, left and diagonal are always ready.\n5. Locate the answer: lcs[m][n], the full strings."},
			{Level: 4, Type: "skeleton", Content: "def longestCommonSubsequence(text1, text2):\n    m, n = len(text1), len(text2)\n    lcs = [[0] * (n + 1) for _ in range(m + 1)]\n    for i in range(1, m + 1):\n        for j in range(1, n + 1):\n            if text1[i - 1] == text2[j - 1]:\n                lcs[i][j] = ____\n            else:\n                lcs[i][j] = ____\n    return ____"},
		},
		Solution: Solution{
			Code: `def longestCommonSubsequence(text1, text2):
    m, n = len(text1), len(text2)
    # lcs[i][j] = LCS length of the prefixes text1[:i] and text2[:j]
    lcs = [[0] * (n + 1) for _ in range(m + 1)]   # row 0 and column 0 stay 0
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if text1[i - 1] == text2[j - 1]:
                lcs[i][j] = lcs[i - 1][j - 1] + 1          # match: keep both
            else:
                lcs[i][j] = max(lcs[i - 1][j], lcs[i][j - 1])  # skip one side
    return lcs[m][n]`,
			Explanation: "The first two-string problem: the state is a pair of prefix lengths (i, j), one from each string. It builds on Unique Paths (problem 1), which also filled a 2D table row by row from the cell above and the cell on the left; here a cell can also read its upper-left diagonal neighbour, and which neighbours it reads depends on whether the two current letters match.",
			BruteForce:  "List every subsequence of text1 (2^m of them), check each against text2 with two pointers, and keep the longest that fits. For m = 1000 that is 2^1000 candidates. The plain recursion on (i, j) without a table is also exponential: on a mismatch it branches twice, and the same (i, j) pair is reached along many different paths of skips.",
			KeyInsight:  "Everything the future needs to know is how many letters of each string are still unused, i.e. the pair (i, j). There are only (m + 1) x (n + 1) such pairs, about one million for two 1000-letter strings, and each is solved once from three neighbours.",
			WhyItWorks:  "State: lcs[i][j] = LCS length of text1[:i] and text2[:j]. Recurrence, by looking at the last letters a = text1[i - 1] and b = text2[j - 1]. If a == b, some longest common subsequence ends by pairing a with b (if an optimal one doesn't use them, swap its last pair for (a, b) without losing length), so lcs[i][j] = lcs[i - 1][j - 1] + 1. If a != b, they can't both be the last letter of the common subsequence, so at least one is unused: either the answer lives in text1[:i - 1] and text2[:j] (up) or in text1[:i] and text2[:j - 1] (left); max of the two covers both cases. Base cases: an empty prefix has LCS 0 (row 0 and column 0). Order: increasing i, then j, makes up, left and diagonal final before they are read. For \"abcde\" vs \"ace\" the last row of the table is 0 1 2 3, giving 3.",
			CommonMistakes: []string{
				"Indexing the strings with the table indices: comparing text1[i] with text2[j] while the table is (m + 1) x (n + 1). Row i of the table means 'the first i letters', whose last letter is text1[i - 1]. The off-by-one either crashes at i = m or compares the wrong letters.",
				"On a mismatch, writing lcs[i - 1][j - 1] (skip both letters) instead of max(up, left). Skipping both throws away a letter that might still match later: for \"abc\" vs \"cab\" it never pairs anything and returns 0 instead of 2.",
				"Greedy matching: for each letter of text1, grab its first unused occurrence in text2. For text1 = \"cab\", text2 = \"abc\" it pairs the c's first, leaving nothing after them, and returns 1 instead of 2. The DP tries both 'skip' choices instead of committing.",
			},
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "One cell per pair of prefix lengths. The extra row and column (index 0) stand for empty prefixes, which keeps the recurrence free of special cases.", CodeSnippet: "lcs = [[0] * (n + 1) for _ in range(m + 1)]", LineStart: 3, LineEnd: 4},
				{Title: "Set the base cases", Explanation: "Row 0 and column 0 are already 0 from the initialisation: nothing is common with an empty string. The loops start at 1 so they never overwrite them.", CodeSnippet: "lcs = [[0] * (n + 1) for _ in range(m + 1)]", LineStart: 4, LineEnd: 4},
				{Title: "Choose the evaluation order", Explanation: "Row by row, left to right. Cell (i, j) reads (i - 1, j - 1), (i - 1, j) and (i, j - 1), which are all above it or to its left.", CodeSnippet: "for i in range(1, m + 1):\n    for j in range(1, n + 1):", LineStart: 5, LineEnd: 6},
				{Title: "Write the recurrence", Explanation: "Match: both letters join the subsequence, so take the diagonal plus one. Mismatch: skip text1's letter (up) or text2's letter (left) and keep the better result.", CodeSnippet: "if text1[i - 1] == text2[j - 1]:\n    lcs[i][j] = lcs[i - 1][j - 1] + 1\nelse:\n    lcs[i][j] = max(lcs[i - 1][j], lcs[i][j - 1])", LineStart: 7, LineEnd: 10},
				{Title: "Locate the answer", Explanation: "The bottom-right cell covers all of both strings.", CodeSnippet: "return lcs[m][n]", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Optimise space: two rows",
					Code: `def longestCommonSubsequence(text1, text2):
    n = len(text2)
    prev = [0] * (n + 1)              # row i - 1 of the table
    for ch in text1:
        cur = [0] * (n + 1)           # row i
        for j in range(1, n + 1):
            if ch == text2[j - 1]:
                cur[j] = prev[j - 1] + 1
            else:
                cur[j] = max(prev[j], cur[j - 1])
        prev = cur
    return prev[n]`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Row i only reads row i - 1 (up and diagonal) and itself (left), so keep prev and cur. A single array would also work but needs one saved variable for the diagonal, because by the time you reach column j the old row's j - 1 has been overwritten. Two rows are easier to get right.",
				},
				{
					Name: "Top-down: the brute-force recursion plus a cache",
					Code: `from functools import lru_cache

def longestCommonSubsequence(text1, text2):
    @lru_cache(maxsize=None)    # without the cache this is the exponential brute force
    def best(i, j):             # LCS of the suffixes text1[i:] and text2[j:]
        if i == len(text1) or j == len(text2):
            return 0
        if text1[i] == text2[j]:
            return 1 + best(i + 1, j + 1)
        return max(best(i + 1, j), best(i, j + 1))

    return best(0, 0)`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "The same recurrence on suffixes instead of prefixes: best(i, j) looks at text1[i] and text2[j]. Without the cache, every mismatch branches twice and the 300 x 250 test times out; with it, each (i, j) is solved once. The recursion can go m + n calls deep, so for two 1000-letter strings the table is the safer choice in Python.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "When the last letters match, why is it safe to take the diagonal + 1 without also checking up and left?", Answer: "Up and left can never beat diagonal + 1. Dropping one letter from one string can lose at most one from the LCS, so lcs[i - 1][j] <= lcs[i - 1][j - 1] + 1, and the same for left. And pairing the two equal last letters is always allowed: take any optimal common subsequence; if it doesn't use both last letters, replace its final pair with (text1[i - 1], text2[j - 1]) and the length doesn't drop."},
			{Prompt: "How would you return the subsequence itself, not just its length?", Answer: "Keep the full table and walk back from (m, n): if the letters at i - 1 and j - 1 match, that letter is in the answer, so record it and move diagonally; otherwise move to whichever of (i - 1, j) and (i, j - 1) holds the larger value (that's the choice the max made). Stop at row 0 or column 0 and reverse the recorded letters. This walk is why the 2D table is worth keeping when the question asks for the actual string."},
			{Prompt: "Edit Distance (later in this ladder) uses the same table shape. What will change in the cell rule?", Answer: "The cell will hold a minimum cost instead of a maximum length. A match still copies the diagonal (free), but a mismatch has three options, each costing 1: delete (up), insert (left) or replace (diagonal). The base cases change too: turning a prefix of length i into the empty string costs i deletions, so row 0 and column 0 are 0, 1, 2, ... instead of all 0."},
		},
	},
	{
		ID:              "best-time-to-buy-sell-stock-cooldown",
		Number:          309,
		Order:           3,
		Title:           "Best Time to Buy and Sell Stock with Cooldown",
		Slug:            "best-time-to-buy-and-sell-stock-with-cooldown",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Array", "Dynamic Programming", "State Machine"},
		RelatedChapters: []int{16, 15},
		Description: `prices[i] is the price of one share of a stock on day i. You may trade as many times as you like, under these rules:

- You hold at most one share at a time: you must sell it before you buy again.
- Each day you may buy (if you hold nothing), sell (if you hold a share) or do nothing.
- **Cooldown:** on the day right after a sale you may not buy.

Profit is the total of all sale prices minus all purchase prices. Return the maximum profit you can make. If no trade is profitable, return 0 (just never buy). A share you still hold at the end earns nothing, so it never helps to end holding one.`,
		Constraints: []string{
			"1 <= prices.length <= 5000",
			"0 <= prices[i] <= 1000",
		},
		Examples: []Example{
			{Input: "prices = [1,2,3,0,2]", Output: "3", Explanation: "Buy on day 0 (price 1), sell on day 1 (price 2): +1. Day 2 is the cooldown. Buy on day 3 (price 0), sell on day 4 (price 2): +2. Total 3. Holding until day 2 (+2) would make day 3 the cooldown and waste the price of 0."},
			{Input: "prices = [1]", Output: "0", Explanation: "With one day you can't buy and sell."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"prices": []int{1, 2, 3, 0, 2}}, Expected: 3, Note: "example: buy, sell, cooldown, buy, sell"},
			{Input: map[string]interface{}{"prices": []int{1}}, Expected: 0, Note: "one day: no trade possible"},
			{Input: map[string]interface{}{"prices": []int{5, 4, 3, 2, 1}}, Expected: 0, Hidden: true, Note: "falling prices: never trade"},
			{Input: map[string]interface{}{"prices": []int{3, 3, 3}}, Expected: 0, Hidden: true, Note: "flat prices: profit 0"},
			{Input: map[string]interface{}{"prices": []int{1, 2}}, Expected: 1, Hidden: true, Note: "two days: one trade"},
			{Input: map[string]interface{}{"prices": []int{1, 3, 1, 3}}, Expected: 2, Hidden: true, Note: "the cooldown blocks buying right after selling"},
			{Input: map[string]interface{}{"prices": []int{2, 1, 4}}, Expected: 3, Hidden: true, Note: "buy on the dip, not on day 0"},
			{Input: map[string]interface{}{"prices": []int{6, 1, 6, 4, 3, 0, 2}}, Expected: 7, Hidden: true, Note: "two trades separated by a cooldown"},
			{Input: map[string]interface{}{"prices": []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}}, Expected: 11, Hidden: true, Note: "choose which small trades to give up"},
			{Input: map[string]interface{}{"prices": dp2dSeq(2000, 1001, 3)}, Expected: 264101, Hidden: true, Note: "2000 days: exponential search times out"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def maxProfit(prices):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Knowing only the day number isn't enough to decide what you may do next. What else do you need to know about yesterday to say whether buying today is allowed?"},
			{Level: 2, Type: "insight", Content: "Give the state a second index for the situation you are in at the end of the day: HOLD (own a share), SOLD (sold today, so tomorrow is a cooldown), REST (own nothing and free to buy). Then best[i][state] = most cash after day i in that state, and each state has only one or two ways to be reached from yesterday. Chapter 16 calls this a state machine DP: the table has one row per day and one column per state.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Define the state: best[i][HOLD], best[i][SOLD], best[i][REST] = most cash at the end of day i in that situation.\n2. Write the recurrence: HOLD = keep holding, or buy today from yesterday's REST; SOLD = sell today what yesterday's HOLD owned; REST = stay resting, or yesterday was the sale day (the cooldown).\n3. Set the base cases: day 0: HOLD = -prices[0], SOLD impossible (-infinity), REST = 0.\n4. Choose the evaluation order: day by day; each row reads only the row before.\n5. Locate the answer: the last day, in SOLD or REST (never end holding).\n6. Optimise space: three variables are enough."},
			{Level: 4, Type: "skeleton", Content: "def maxProfit(prices):\n    n = len(prices)\n    HOLD, SOLD, REST = 0, 1, 2\n    best = [[0] * 3 for _ in range(n)]\n    best[0][HOLD] = ____\n    best[0][SOLD] = float('-inf')\n    for i in range(1, n):\n        price = prices[i]\n        best[i][HOLD] = max(best[i - 1][HOLD], ____)\n        best[i][SOLD] = ____\n        best[i][REST] = max(best[i - 1][REST], ____)\n    return ____"},
		},
		Solution: Solution{
			Code: `def maxProfit(prices):
    n = len(prices)
    HOLD, SOLD, REST = 0, 1, 2
    # best[i][state] = most cash we can have at the end of day i in that state
    best = [[0] * 3 for _ in range(n)]
    best[0][HOLD] = -prices[0]          # bought on day 0
    best[0][SOLD] = float('-inf')       # can't sell before buying
    best[0][REST] = 0                   # did nothing
    for i in range(1, n):
        price = prices[i]
        best[i][HOLD] = max(best[i - 1][HOLD], best[i - 1][REST] - price)
        best[i][SOLD] = best[i - 1][HOLD] + price
        best[i][REST] = max(best[i - 1][REST], best[i - 1][SOLD])
    return max(best[n - 1][SOLD], best[n - 1][REST])`,
			Explanation: "The second index here is not a second string or a grid column but a small set of situations (hold, just sold, resting): a state machine. It builds on Unique Paths and LCS (problems 1 and 2), whose cells also read only the previous row; the new idea is that the columns are named situations with rules for moving between them. The table has n rows and 3 columns, and we fill it day by day.",
			BruteForce:  "Try every sequence of daily actions (buy, sell or wait, subject to the rules) and keep the best total. That is up to 3^n sequences, or 2^n if you only branch on 'act or wait'. For n = 2000 that is hopeless. The waste: many different histories end day i in the same situation, say 'resting', with different cash, and only the one with the most cash can lead to the best future; the brute force continues all of them.",
			KeyInsight:  "What you may do tomorrow depends only on today's situation (holding, just sold, resting), not on the full history. So for each day and each situation, keep only the most cash any history can reach. That is 3 numbers per day instead of exponentially many histories.",
			WhyItWorks:  "State: best[i][s] = the most cash (sales minus purchases) after day i, over all legal histories that end day i in situation s. Recurrence: you HOLD at the end of day i if you already held yesterday, or you were in REST yesterday and buy today (SOLD yesterday is excluded: that is the cooldown). You are in SOLD on day i only by selling today what you held yesterday. You are in REST if you rested yesterday too, or yesterday was a sale day (today is the forced cooldown). These cases cover every legal transition, so taking the max over them is correct; and because the future depends only on the situation, the best cash per situation is all we need (optimal substructure). Base: on day 0 you can hold only by buying (-prices[0]), you can't have sold (-infinity marks 'impossible'), or you rest with 0. Answer: the best of SOLD and REST on the last day; ending in HOLD means a share was bought and never sold. For [1,2,3,0,2] the rows are (HOLD, SOLD, REST): (-1, -inf, 0), (-1, 1, 0), (-1, 2, 1), (1, -1, 2), (1, 3, 2), so the answer is max(3, 2) = 3.",
			CommonMistakes: []string{
				"Letting HOLD buy from yesterday's SOLD: best[i][HOLD] = max(best[i - 1][HOLD], best[i - 1][SOLD] - price). That ignores the cooldown. For [1,3,1,3] it buys again the day after selling and returns 4 instead of 2.",
				"Updating three variables one after another (hold = ..., then sold = hold + price, then rest = max(rest, sold)) so that later lines read TODAY's values instead of yesterday's. Then rest picks up a sale made today, which skips the cooldown: [1,2,3,0,2] returns 4 instead of 3. Compute all three from the previous row (the table does this automatically; with variables, use one simultaneous assignment).",
				"Starting SOLD on day 0 at 0 instead of -infinity and not thinking about it. Here it happens not to change the answer, but in variants (a fee per sale, at most k transactions) a fake 'sold on day 0 for free' state can produce profits that no real history achieves. Mark impossible states as impossible.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "One row per day, one column per situation at the end of that day. Naming the columns keeps the recurrence readable.", CodeSnippet: "HOLD, SOLD, REST = 0, 1, 2\nbest = [[0] * 3 for _ in range(n)]", LineStart: 3, LineEnd: 5},
				{Title: "Set the base cases", Explanation: "Day 0: holding means you bought at prices[0]; 'just sold' is impossible, marked -infinity so a max never picks it; resting means you did nothing.", CodeSnippet: "best[0][HOLD] = -prices[0]\nbest[0][SOLD] = float('-inf')\nbest[0][REST] = 0", LineStart: 6, LineEnd: 8},
				{Title: "Choose the evaluation order", Explanation: "Day by day. Every right-hand side reads row i - 1 only, so the three cells of row i can be filled in any order.", CodeSnippet: "for i in range(1, n):", LineStart: 9, LineEnd: 10},
				{Title: "Write the recurrence", Explanation: "Each line lists every way to arrive in that situation. HOLD: keep, or buy from REST (not from SOLD: cooldown). SOLD: sell yesterday's share. REST: keep resting, or this is the cooldown day after a sale.", CodeSnippet: "best[i][HOLD] = max(best[i - 1][HOLD], best[i - 1][REST] - price)\nbest[i][SOLD] = best[i - 1][HOLD] + price\nbest[i][REST] = max(best[i - 1][REST], best[i - 1][SOLD])", LineStart: 11, LineEnd: 13},
				{Title: "Locate the answer", Explanation: "End the last day without a share: either you sold that day or you are resting. If no trade helps, REST stays 0.", CodeSnippet: "return max(best[n - 1][SOLD], best[n - 1][REST])", LineStart: 14, LineEnd: 14},
			},
			Alternatives: []AltSolution{
				{
					Name: "Optimise space: three variables",
					Code: `def maxProfit(prices):
    hold, sold, rest = -prices[0], float('-inf'), 0
    for price in prices[1:]:
        # the right-hand side uses only yesterday's values
        hold, sold, rest = max(hold, rest - price), hold + price, max(rest, sold)
    return max(sold, rest)`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Each row reads only the row before, so three variables are enough. The tuple assignment evaluates the whole right-hand side with yesterday's values before assigning, which is exactly what the table did.",
				},
				{
					Name: "Top-down: the brute-force recursion plus a cache",
					Code: `from functools import lru_cache

def maxProfit(prices):
    n = len(prices)

    @lru_cache(maxsize=None)    # without the cache: 2^n choices
    def best(day, holding):     # most profit from day onwards
        if day >= n:
            return 0
        wait = best(day + 1, holding)
        if holding:
            act = prices[day] + best(day + 2, False)   # sell, skip the cooldown day
        else:
            act = -prices[day] + best(day + 1, True)   # buy
        return max(wait, act)

    return best(0, False)`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "best(day, holding) is the brute-force decision tree: wait, or act (buy, or sell and jump over the cooldown day). Without the cache it branches twice per day; with it there are only 2n states. The cooldown appears as 'day + 2' after a sale instead of as a separate SOLD state, which is another way to encode the same rule.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why do we need a separate SOLD state instead of just HOLD and NOT-HOLD?", Answer: "Because 'not holding' covers two situations with different rules for tomorrow: if you sold today you may not buy tomorrow, but if you were already resting you may. The state must contain everything the future depends on; merging them would let you buy during the cooldown. With SOLD separate, HOLD can be reached only from REST, which encodes the rule."},
			{Prompt: "What changes if there is no cooldown but every sale costs a fee?", Answer: "Drop the SOLD state: two situations, HOLD and FREE. hold = max(hold, free - price) and free = max(free, hold + price - fee). The table shape (days x situations) and the recipe stay the same; only the list of situations and transitions changes. That is the strength of the state-machine view."},
			{Prompt: "Why is it safe to start best[0][SOLD] at -infinity rather than 0?", Answer: "-infinity says 'this situation cannot happen', so any max that considers it ignores it, and any sum built on it stays -infinity. 0 would claim you sold on day 0 for free, an impossible history. Here it doesn't change the answer, but encoding impossible states honestly prevents subtle bugs in variants."},
		},
	},
	{
		ID:              "coin-change-ii",
		Number:          518,
		Order:           4,
		Title:           "Coin Change II",
		Slug:            "coin-change-ii",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Array", "Dynamic Programming", "Knapsack"},
		RelatedChapters: []int{16, 15},
		Description: `coins lists the values of the coin types you have; each value appears once, and you have an **unlimited supply** of every type. Return the number of different **combinations** of coins that add up to exactly amount.

A combination is a multiset: only how many coins of each type you use matters, not their order. So 1 + 2 + 2 and 2 + 1 + 2 are the same combination.

If amount is 0 there is exactly one combination (use no coins). If amount can't be made, return 0. The answer fits in a signed 32-bit integer.`,
		Constraints: []string{
			"1 <= coins.length <= 300",
			"1 <= coins[i] <= 5000",
			"All values in coins are different",
			"0 <= amount <= 5000",
		},
		Examples: []Example{
			{Input: "amount = 5, coins = [1,2,5]", Output: "4", Explanation: "5 = 5, 5 = 2 + 2 + 1, 5 = 2 + 1 + 1 + 1, 5 = 1 + 1 + 1 + 1 + 1."},
			{Input: "amount = 3, coins = [2]", Output: "0", Explanation: "Coins of 2 only make even amounts."},
			{Input: "amount = 10, coins = [10]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"amount": 5, "coins": []int{1, 2, 5}}, Expected: 4, Note: "example: four combinations"},
			{Input: map[string]interface{}{"amount": 3, "coins": []int{2}}, Expected: 0, Note: "odd amount, only even coins: 0"},
			{Input: map[string]interface{}{"amount": 10, "coins": []int{10}}, Expected: 1, Note: "one coin equals the amount"},
			{Input: map[string]interface{}{"amount": 0, "coins": []int{7}}, Expected: 1, Hidden: true, Note: "amount 0: one way (take nothing)"},
			{Input: map[string]interface{}{"amount": 4, "coins": []int{1, 2}}, Expected: 3, Hidden: true, Note: "1+1+2 and 2+1+1 are the same combination: 3, not 5"},
			{Input: map[string]interface{}{"amount": 3, "coins": []int{5, 10}}, Expected: 0, Hidden: true, Note: "every coin is bigger than the amount"},
			{Input: map[string]interface{}{"amount": 11, "coins": []int{5, 2, 1}}, Expected: 11, Hidden: true, Note: "coins are not sorted"},
			{Input: map[string]interface{}{"amount": 100, "coins": []int{1, 5, 10, 25, 50}}, Expected: 292, Hidden: true, Note: "US coins, one dollar"},
			{Input: map[string]interface{}{"amount": 500, "coins": []int{3, 5, 7, 8, 9, 10, 11}}, Expected: 35502874, Hidden: true, Note: "large count from many small coins"},
			{Input: map[string]interface{}{"amount": 5000, "coins": []int{11, 24, 37, 50, 63, 76, 89, 102}}, Expected: 992951208, Hidden: true, Note: "amount 5000: needs the table"},
		},
		TimeComplexity:  "O(n * amount)",
		SpaceComplexity: "O(n * amount)",
		StarterCode:     "def change(amount, coins):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you build sums by 'pick any coin as the next coin', you'll count 1 + 2 and 2 + 1 separately. How could you force every combination to be built in exactly one order?"},
			{Level: 2, Type: "insight", Content: "Decide the coin types one at a time: first how many 1s, then how many 2s, and so on. Let ways[i][a] = combinations that use only the first i coin types and sum to a. For coin type i, either you never use it (ways[i - 1][a]) or you use it at least once more (ways[i][a - coin], same row, because you may use it again). This is the unbounded knapsack from Chapter 16: item index x budget.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Define the state: ways[i][a] = combinations of the first i coin types that sum to a; the table is (n + 1) x (amount + 1).\n2. Set the base cases: ways[i][0] = 1 for every i (take nothing); ways[0][a] = 0 for a > 0 (no coin types, no way).\n3. Write the recurrence: ways[i][a] = ways[i - 1][a] + (ways[i][a - coin] if a >= coin).\n4. Choose the evaluation order: coin types in the outer loop, amounts increasing in the inner loop, so ways[i][a - coin] is already final.\n5. Locate the answer: ways[n][amount].\n6. Optimise space: one row, amounts going forwards."},
			{Level: 4, Type: "skeleton", Content: "def change(amount, coins):\n    n = len(coins)\n    ways = [[0] * (amount + 1) for _ in range(n + 1)]\n    for i in range(n + 1):\n        ways[i][0] = ____\n    for i in range(1, n + 1):\n        coin = coins[i - 1]\n        for a in range(1, amount + 1):\n            ways[i][a] = ____\n            if a >= coin:\n                ways[i][a] += ____\n    return ways[n][amount]"},
		},
		Solution: Solution{
			Code: `def change(amount, coins):
    n = len(coins)
    # ways[i][a] = combinations that use only the first i coin types and sum to a
    ways = [[0] * (amount + 1) for _ in range(n + 1)]
    for i in range(n + 1):
        ways[i][0] = 1                      # amount 0: take no coins
    for i in range(1, n + 1):
        coin = coins[i - 1]
        for a in range(1, amount + 1):
            ways[i][a] = ways[i - 1][a]     # never use this coin type
            if a >= coin:
                ways[i][a] += ways[i][a - coin]   # use it once more: SAME row
    return ways[n][amount]`,
			Explanation: "The first knapsack-style problem: rows are items (coin types), columns are a budget (the amount). It builds on Chapter 15's Coin Change, which used a 1D table over amounts to find the fewest coins; counting combinations needs the coin index as a second dimension so that each combination is counted once, in the order 'all 1s, then all 2s, then all 5s'. The rows look like Unique Paths (problem 1): each cell adds the cell above and a cell to its left in the same row.",
			BruteForce:  "Recursively choose how many coins of each type to use: for coin i try 0, 1, 2, ... copies and recurse on the rest. The number of calls grows with the answer itself (35 million combinations for amount 500 with coins [3,5,7,8,9,10,11]), and the same question 'how many ways to make 37 from coins[4:]' is asked again for every different way of spending the first four coin types.",
			KeyInsight:  "Fix the order in which coin types are decided. Then a combination is a path through the table that moves from row to row (done with a type) or stays in the same row and moves left by coin (one more of this type). The pair (coin index, remaining amount) is all the future depends on, so there are only (n + 1) x (amount + 1) questions.",
			WhyItWorks:  "State: ways[i][a] = number of multisets of the first i coin types summing to a. Recurrence: split those multisets by whether they contain coin i at all. Those that don't are counted by ways[i - 1][a]. Those that do: remove one copy of coin i, and you get a multiset of the first i types summing to a - coin (it may still contain coin i), and this is a one-to-one match, so they are counted by ways[i][a - coin]. The two groups are disjoint and cover everything, so we add. Base: ways[i][0] = 1 (the empty multiset), ways[0][a > 0] = 0. Order: ways[i][a] reads row i - 1 (finished) and row i at a smaller amount (already filled because a increases). For coins [1,2,5], amount 5, the rows are: no coins 1 0 0 0 0 0; {1}: 1 1 1 1 1 1; {1,2}: 1 1 2 2 3 3; {1,2,5}: 1 1 2 2 3 4. Answer 4.",
			CommonMistakes: []string{
				"Putting the amount loop outside and the coin loop inside (for a: for coin: ways[a] += ways[a - coin]). That counts ordered sequences: 1 + 2 and 2 + 1 become two answers. For amount 4, coins [1, 2] it returns 5 instead of 3. (That version solves a different problem, Combination Sum IV.)",
				"Reading the previous row for the 'use it' case: ways[i][a] += ways[i - 1][a - coin]. That allows each coin type at most once (0/1 knapsack). For amount 5, coins [1, 2, 5] it returns 1 (only {5}) instead of 4.",
				"Setting ways[0] (or ways[i][0]) to 0. Every combination is eventually built on top of 'amount 0, no coins', so if that base case is 0 everything stays 0. Amount 0 has exactly one combination: the empty one.",
			},
			TimeComplexity:  "O(n * amount)",
			SpaceComplexity: "O(n * amount)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "Rows: how many coin types are allowed (0..n). Columns: the amount to make (0..amount).", CodeSnippet: "ways = [[0] * (amount + 1) for _ in range(n + 1)]", LineStart: 3, LineEnd: 4},
				{Title: "Set the base cases", Explanation: "Amount 0 can always be made in exactly one way, with no coins. Row 0 (no coin types) stays 0 for every positive amount.", CodeSnippet: "for i in range(n + 1):\n    ways[i][0] = 1", LineStart: 5, LineEnd: 6},
				{Title: "Choose the evaluation order", Explanation: "Coin types outside, amounts increasing inside. Coin types outside is what makes these combinations (each one is built in the fixed order of the coin list); increasing amounts make ways[i][a - coin] ready.", CodeSnippet: "for i in range(1, n + 1):\n    coin = coins[i - 1]\n    for a in range(1, amount + 1):", LineStart: 7, LineEnd: 9},
				{Title: "Write the recurrence", Explanation: "Never use this coin type (the cell above), plus use it at least once (the cell coin columns to the left in the SAME row, because one more copy is still allowed).", CodeSnippet: "ways[i][a] = ways[i - 1][a]\nif a >= coin:\n    ways[i][a] += ways[i][a - coin]", LineStart: 10, LineEnd: 12},
				{Title: "Locate the answer", Explanation: "All coin types allowed, full amount: the bottom-right cell.", CodeSnippet: "return ways[n][amount]", LineStart: 13, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Optimise space: one row, amounts forwards",
					Code: `def change(amount, coins):
    ways = [1] + [0] * amount       # ways[a] for the coin types seen so far
    for coin in coins:
        for a in range(coin, amount + 1):   # forwards: ways[a - coin] may already use this coin
            ways[a] += ways[a - coin]
    return ways[amount]`,
					TimeComplexity:  "O(n * amount)",
					SpaceComplexity: "O(amount)",
					Explanation:     "Keep one row. Before ways[a] is updated it still holds the row above (ways[i - 1][a]). Going FORWARDS, ways[a - coin] has already been updated for this coin, so it is ways[i][a - coin]: exactly the same-row read the unbounded recurrence wants. Compare Target Sum (problem 5), where each number may be used once and the loop must run backwards.",
				},
				{
					Name: "Top-down: the brute-force recursion plus a cache",
					Code: `from functools import lru_cache

def change(amount, coins):
    @lru_cache(maxsize=None)    # without the cache this is the exponential brute force
    def count(i, remaining):    # combinations of coins[i:] that sum to remaining
        if remaining == 0:
            return 1
        if i == len(coins):
            return 0
        total = count(i + 1, remaining)          # done with coin type i
        if coins[i] <= remaining:
            total += count(i, remaining - coins[i])   # take one more of coin i
        return total

    return count(0, amount)`,
					TimeComplexity:  "O(n * amount)",
					SpaceComplexity: "O(n * amount)",
					Explanation:     "count(i, remaining) is the brute force 'finish with coin i, or take one more of it'. Because it always moves forward through the coin list, each combination is built once. Without the cache the large tests time out; with it there are n * amount states. The recursion can be about amount / smallest coin calls deep, so for amount 5000 with a coin of 1 the table is safer.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the 'use it' case read ways[i][a - coin] (same row) and not ways[i - 1][a - coin]?", Answer: "After using one copy of coin i we may still use more copies, so the remaining amount a - coin is still made from the first i coin types, which is row i. Reading row i - 1 would forbid a second copy: that is the 0/1 knapsack, where each item is used at most once."},
			{Prompt: "In the one-row version, what goes wrong if the amount loop runs backwards (from amount down to coin)?", Answer: "Going backwards, ways[a - coin] has NOT been updated for this coin yet, so it still holds the previous row. The update becomes ways[a] += ways_old[a - coin], which is the 0/1 rule: each coin type used at most once. For amount 5 and coins [1, 2, 5] it returns 1 instead of 4. Direction decides whether the same-row (unbounded) or previous-row (0/1) value is read."},
			{Prompt: "How would you count ordered sequences instead (1 + 2 and 2 + 1 counted separately)?", Answer: "Drop the coin index: seq[a] = sum of seq[a - c] over all coins c <= a, with amounts in the outer loop. Now the state asks only 'what was the last coin?', and every coin can be last, so different orders are different sequences. For amount 4 and coins [1, 2] that gives 5. That is LeetCode 377, Combination Sum IV."},
		},
	},
	{
		ID:              "target-sum",
		Number:          494,
		Order:           5,
		Title:           "Target Sum",
		Slug:            "target-sum",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Array", "Dynamic Programming", "Knapsack", "Backtracking"},
		RelatedChapters: []int{16, 15, 12},
		Description: `You are given a list of non-negative integers nums and an integer target. Put a sign, + or -, in front of **every** number and add them all up. For example, with nums = [2, 1] you could build +2 - 1 = 1.

Return how many of the possible sign assignments give exactly target. Two assignments are different if at least one position has a different sign, so for nums = [0], target = 0 the answer is 2 (+0 and -0 both count).`,
		Constraints: []string{
			"1 <= nums.length <= 20",
			"0 <= nums[i] <= 1000",
			"0 <= sum(nums) <= 1000",
			"-1000 <= target <= 1000",
		},
		Examples: []Example{
			{Input: "nums = [1,1,1,1,1], target = 3", Output: "5", Explanation: "Exactly one number gets a minus sign (4 - 1 = 3), and it can be any of the five: -1+1+1+1+1, +1-1+1+1+1, +1+1-1+1+1, +1+1+1-1+1, +1+1+1+1-1."},
			{Input: "nums = [1], target = 1", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 1, 1}, "target": 3}, Expected: 5, Note: "example: five places for the single minus"},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 1}, Expected: 1, Note: "one number, one way"},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 2}, Expected: 0, Hidden: true, Note: "target bigger than the sum: 0"},
			{Input: map[string]interface{}{"nums": []int{1, 2}, "target": 2}, Expected: 0, Hidden: true, Note: "total + target is odd: 0"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 1, 1}, "target": -3}, Expected: 5, Hidden: true, Note: "negative target"},
			{Input: map[string]interface{}{"nums": []int{0}, "target": 0}, Expected: 2, Hidden: true, Note: "+0 and -0 count as two expressions"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0, 0, 0, 0, 0, 0, 1}, "target": 1}, Expected: 256, Hidden: true, Note: "every zero doubles the count"},
			{Input: map[string]interface{}{"nums": []int{2, 3, 5}, "target": 0}, Expected: 2, Hidden: true, Note: "target 0 with distinct numbers"},
			{Input: map[string]interface{}{"nums": []int{1000}, "target": -1000}, Expected: 1, Hidden: true, Note: "largest value, negative target"},
			{Input: map[string]interface{}{"nums": []int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9}, "target": 1}, Expected: 40, Hidden: true, Note: "mixed values with a zero"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, "target": 0}, Expected: 184756, Hidden: true, Note: "20 numbers: one million sign patterns"},
		},
		TimeComplexity:  "O(n * total)",
		SpaceComplexity: "O(n * total)",
		StarterCode:     "def findTargetSumWays(nums, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Call the numbers that get a + sign the group P, and the rest the group N. What two equations do sum(P) and sum(N) satisfy? Can you solve them for sum(P)?"},
			{Level: 2, Type: "insight", Content: "sum(P) - sum(N) = target and sum(P) + sum(N) = total, so sum(P) = (total + target) / 2. Each sign assignment is exactly one choice of the plus-group, so you are counting subsets with a given sum. That is the 0/1 knapsack from Chapter 16: ways[i][s] = subsets of the first i numbers summing to s, where each number is used at most once.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Turn signs into a subset: if (total + target) is odd or |target| > total, return 0; otherwise goal = (total + target) // 2.\n2. Define the state: ways[i][s] = subsets of the first i numbers that sum to s, for s in 0..goal.\n3. Set the base cases: ways[0][0] = 1 (the empty subset); every other ways[0][s] = 0.\n4. Write the recurrence: ways[i][s] = ways[i - 1][s] (num gets -) + ways[i - 1][s - num] if s >= num (num gets +).\n5. Choose the evaluation order: number by number; every cell reads only the previous row.\n6. Locate the answer: ways[n][goal].\n7. Optimise space: one row, s going backwards."},
			{Level: 4, Type: "skeleton", Content: "def findTargetSumWays(nums, target):\n    total = sum(nums)\n    if abs(target) > total or (total + target) % 2 == 1:\n        return 0\n    goal = ____\n    n = len(nums)\n    ways = [[0] * (goal + 1) for _ in range(n + 1)]\n    ways[0][0] = ____\n    for i in range(1, n + 1):\n        num = nums[i - 1]\n        for s in range(goal + 1):\n            ways[i][s] = ____\n            if s >= num:\n                ways[i][s] += ____\n    return ways[n][goal]"},
		},
		Solution: Solution{
			Code: `def findTargetSumWays(nums, target):
    total = sum(nums)
    # plus-group P and minus-group N: P - N = target and P + N = total,
    # so P = (total + target) / 2 must be a whole number in 0..total
    if abs(target) > total or (total + target) % 2 == 1:
        return 0
    goal = (total + target) // 2
    n = len(nums)
    # ways[i][s] = subsets of the first i numbers that sum to s
    ways = [[0] * (goal + 1) for _ in range(n + 1)]
    ways[0][0] = 1                                # the empty subset
    for i in range(1, n + 1):
        num = nums[i - 1]
        for s in range(goal + 1):
            ways[i][s] = ways[i - 1][s]           # num goes in the minus group
            if s >= num:
                ways[i][s] += ways[i - 1][s - num]    # num goes in the plus group
    return ways[n][goal]`,
			Explanation: "Same table shape as Coin Change II (problem 4), items x budget, with one change: each number is used exactly once, so the 'take it' case reads the PREVIOUS row. That is the 0/1 knapsack. The extra step is a short piece of algebra that turns 'choose a sign for each number' into 'choose which numbers are in the plus-group', so that the budget is a fixed subset sum. The top-down alternative skips the algebra and memoises on (index, running sum) instead.",
			BruteForce:  "Try all 2^n sign assignments (recursion that branches on + and - for each number) and count those that hit target. For n = 20 that is 1,048,576 assignments, which is still fine, but the waste is clear: after the first i numbers only the running sum matters, and there are at most 2 * total + 1 different running sums, yet the brute force continues separately from every assignment that reached the same sum.",
			KeyInsight:  "Let P be the sum of the numbers with +, and N the sum of those with -. Then P - N = target and P + N = total. Adding the equations: 2P = total + target, so P = (total + target) / 2. Signs are fixed once you choose which numbers are in the plus-group, so the answer is the number of subsets with sum P. If total + target is odd, or |target| > total, P is not a whole number in 0..total and the answer is 0.",
			WhyItWorks:  "Step by step for [1,1,1,1,1], target 3: total = 5, P = (5 + 3) / 2 = 4, so we count subsets summing to 4 (choose which four 1s get +): 5 of them. In general the map 'sign assignment -> its plus-group' is one-to-one, and an assignment hits target exactly when its plus-group sums to P, so counting subsets is counting assignments. State: ways[i][s] = number of subsets of the first i numbers with sum s. Recurrence: a subset either leaves out num = nums[i - 1] (ways[i - 1][s]) or includes it once (ways[i - 1][s - num]); both read row i - 1 because num can't be used twice. The two cases are disjoint and complete, so we add. A zero doubles the count (ways[i][s] = 2 * ways[i - 1][s]), matching +0 and -0 being different assignments. Base: before any number only the empty subset exists, sum 0. Answer: ways[n][P].",
			CommonMistakes: []string{
				"Forgetting the feasibility check and computing goal = (total + target) // 2 anyway. For nums = [1, 2], target = 2: total + target = 5 is odd, the floor division gives goal 2, and the subset {2} is counted, returning 1 instead of 0. Also check |target| > total, otherwise a very negative target makes goal negative.",
				"In the one-row version, looping s forwards: for s in range(num, goal + 1). Then ways[s - num] already includes num, so a number can be used twice. For [1, 1, 1, 1, 1], target 3 it counts 'use some 1s several times' and returns 70 instead of 5. Loop s backwards so ways[s - num] is still the previous row.",
				"Treating zeros as 'no choice'. A zero can take + or -, and both count, so each zero doubles the answer: [0, 0, 0, 0, 0, 0, 0, 0, 1] with target 1 gives 2^8 = 256. The recurrence handles this automatically as long as the inner loop includes s = 0 when num = 0.",
			},
			TimeComplexity:  "O(n * total)",
			SpaceComplexity: "O(n * total)",
			Walkthrough: []WalkthroughStep{
				{Title: "Turn signs into a subset", Explanation: "P - N = target and P + N = total give P = (total + target) / 2. If that isn't a whole number between 0 and total, no assignment works.", CodeSnippet: "if abs(target) > total or (total + target) % 2 == 1:\n    return 0\ngoal = (total + target) // 2", LineStart: 2, LineEnd: 7},
				{Title: "Define the state", Explanation: "Rows: how many numbers have been decided (0..n). Columns: the plus-group sum so far (0..goal). Sums above goal can never come back down (all numbers are non-negative), so they are not stored.", CodeSnippet: "ways = [[0] * (goal + 1) for _ in range(n + 1)]", LineStart: 9, LineEnd: 10},
				{Title: "Set the base cases", Explanation: "With no numbers decided, the only subset is the empty one, with sum 0.", CodeSnippet: "ways[0][0] = 1", LineStart: 11, LineEnd: 11},
				{Title: "Write the recurrence", Explanation: "num gets a minus sign (it is not in the plus-group: same sum as the row above) or a plus sign (the row above at s - num). Both read row i - 1, so num is used at most once.", CodeSnippet: "ways[i][s] = ways[i - 1][s]\nif s >= num:\n    ways[i][s] += ways[i - 1][s - num]", LineStart: 15, LineEnd: 17},
				{Title: "Choose the evaluation order", Explanation: "One row per number, top to bottom. Within a row the order of s doesn't matter in the 2D table, because nothing reads the current row.", CodeSnippet: "for i in range(1, n + 1):\n    num = nums[i - 1]\n    for s in range(goal + 1):", LineStart: 12, LineEnd: 14},
				{Title: "Locate the answer", Explanation: "All numbers decided, plus-group sum exactly goal.", CodeSnippet: "return ways[n][goal]", LineStart: 18, LineEnd: 18},
			},
			Alternatives: []AltSolution{
				{
					Name: "Optimise space: one row, sums backwards",
					Code: `def findTargetSumWays(nums, target):
    total = sum(nums)
    if abs(target) > total or (total + target) % 2 == 1:
        return 0
    goal = (total + target) // 2
    ways = [1] + [0] * goal
    for num in nums:
        for s in range(goal, num - 1, -1):    # backwards: ways[s - num] is still the old row
            ways[s] += ways[s - num]
    return ways[goal]`,
					TimeComplexity:  "O(n * total)",
					SpaceComplexity: "O(total)",
					Explanation:     "Same recurrence with one row. Going BACKWARDS from goal, ways[s - num] is a smaller index that hasn't been updated for this number yet, so it still holds the previous row: each number is used at most once. This is the mirror image of Coin Change II, where forwards was right because coins could be reused.",
				},
				{
					Name: "Top-down: memo on (index, running sum)",
					Code: `def findTargetSumWays(nums, target):
    memo = {}

    def count(i, running):      # sign choices for nums[i:] that finish on target
        if i == len(nums):
            return 1 if running == target else 0
        key = (i, running)
        if key not in memo:     # without the memo: 2^n sign patterns
            memo[key] = count(i + 1, running + nums[i]) + count(i + 1, running - nums[i])
        return memo[key]

    return count(0, 0)`,
					TimeComplexity:  "O(n * total)",
					SpaceComplexity: "O(n * total)",
					Explanation:     "No algebra: count(i, running) is the brute-force sign search, and the memo dictionary stores each (i, running) pair. The running sum lies in -total..total, so there are at most n * (2 * total + 1) states. Without the memo it is the 2^n brute force. A dictionary is used because running can be negative.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the one-row loop over s run backwards here, but forwards in Coin Change II?", Answer: "The 'take it' case reads index s - num, which is to the left of s. Going backwards, the left part of the row hasn't been touched for this number yet, so it is still the previous row: the number is used at most once (0/1). Going forwards, the left part has already been updated, so it is the current row: the item may be used again (unbounded). Coin Change II allows unlimited coins, so forwards; Target Sum uses each number once, so backwards."},
			{Prompt: "Why does a zero double the answer, and where does the table do that?", Answer: "+0 and -0 are different assignments with the same value, so every working assignment of the other numbers becomes two. In the table, num = 0 makes both cases read the same cell: ways[i][s] = ways[i - 1][s] + ways[i - 1][s - 0] = 2 * ways[i - 1][s]. For nums with eight zeros and a 1, target 1, that gives 2^8 = 256."},
			{Prompt: "Why is the column range only 0..goal and not -total..total?", Answer: "After the algebra we count subsets of non-negative numbers, and a subset sum only grows as numbers are added, so a partial sum above goal can never come back to goal. Negative sums can't occur at all. The top-down alternative works with signed running sums instead, so it needs the full -total..total range (a dictionary handles the negative keys)."},
		},
	},
	{
		ID:              "interleaving-string",
		Number:          97,
		Order:           6,
		Title:           "Interleaving String",
		Slug:            "interleaving-string",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{16, 15},
		Description: `Given three strings s1, s2 and s3, return true if s3 can be made by **interleaving** s1 and s2, and false otherwise.

Interleaving means shuffling the two strings together like two decks of cards: every letter of s1 and every letter of s2 is used exactly once, the letters of s1 keep their order, the letters of s2 keep their order, and nothing else is added. For example, "aXbY" is an interleaving of "ab" and "XY", but "baXY" is not (b comes before a).

Any of the strings may be empty. If len(s1) + len(s2) != len(s3) the answer is false.`,
		Constraints: []string{
			"0 <= s1.length, s2.length <= 100",
			"0 <= s3.length <= 200",
			"s1, s2 and s3 contain only lowercase English letters",
		},
		Examples: []Example{
			{Input: `s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"`, Output: "true", Explanation: "Split s1 as \"aa\" + \"bc\" + \"c\" and s2 as \"dbbc\" + \"a\", then alternate: \"aa\" + \"dbbc\" + \"bc\" + \"a\" + \"c\" = \"aadbbcbcac\"."},
			{Input: `s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"`, Output: "false", Explanation: "s3 has the right letters, but no order-preserving shuffle produces it."},
			{Input: `s1 = "", s2 = "", s3 = ""`, Output: "true"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s1": "aabcc", "s2": "dbbca", "s3": "aadbbcbcac"}, Expected: true, Note: "example: valid interleaving"},
			{Input: map[string]interface{}{"s1": "aabcc", "s2": "dbbca", "s3": "aadbbbaccc"}, Expected: false, Note: "example: same letters, impossible order"},
			{Input: map[string]interface{}{"s1": "", "s2": "", "s3": ""}, Expected: true, Note: "all empty: true"},
			{Input: map[string]interface{}{"s1": "", "s2": "abc", "s3": "abc"}, Expected: true, Hidden: true, Note: "s1 empty: s3 must equal s2"},
			{Input: map[string]interface{}{"s1": "abc", "s2": "", "s3": "abd"}, Expected: false, Hidden: true, Note: "s2 empty and s3 differs from s1"},
			{Input: map[string]interface{}{"s1": "a", "s2": "", "s3": ""}, Expected: false, Hidden: true, Note: "length mismatch: s3 too short"},
			{Input: map[string]interface{}{"s1": "a", "s2": "b", "s3": "ba"}, Expected: true, Hidden: true, Note: "s2's letter can come first"},
			{Input: map[string]interface{}{"s1": "a", "s2": "b", "s3": "aa"}, Expected: false, Hidden: true, Note: "right length, wrong letters"},
			{Input: map[string]interface{}{"s1": "ab", "s2": "ac", "s3": "acab"}, Expected: true, Hidden: true, Note: "greedy 'take from s1 when it matches' fails here"},
			{Input: map[string]interface{}{"s1": "db", "s2": "b", "s3": "cbb"}, Expected: false, Hidden: true, Note: "first letter matches neither string"},
			{Input: map[string]interface{}{"s1": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", "s2": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "s3": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaac"}, Expected: false, Hidden: true, Note: "many equal letters: plain recursion times out"},
			{Input: map[string]interface{}{"s1": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", "s2": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaac", "s3": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaacaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"}, Expected: true, Hidden: true, Note: "long true case"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def isInterleave(s1, s2, s3):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose you have used the first i letters of s1 and the first j letters of s2. How many letters of s3 have you produced, and which two letters could the next letter of s3 be?"},
			{Level: 2, Type: "insight", Content: "After using i letters of s1 and j of s2 you have produced exactly s3[:i + j], so the pair (i, j) is the whole state. ok[i][j] is true if s3[:i + j] can be built; the last letter came either from s1 (needs ok[i - 1][j] and s1[i - 1] == s3[i + j - 1]) or from s2 (needs ok[i][j - 1] and s2[j - 1] == s3[i + j - 1]). It is the two-string table from Chapter 16, filled with true/false.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Check the lengths: if len(s1) + len(s2) != len(s3), return False.\n2. Define the state: ok[i][j] = can s1[:i] and s2[:j] interleave into s3[:i + j]?\n3. Set the base cases: ok[0][0] = True; the rest of row 0 and column 0 use only one string, and the same rule handles them if you guard i > 0 and j > 0.\n4. Write the recurrence: ok[i][j] = (came from s1) or (came from s2).\n5. Choose the evaluation order: rows top to bottom, columns left to right.\n6. Locate the answer: ok[m][n]."},
			{Level: 4, Type: "skeleton", Content: "def isInterleave(s1, s2, s3):\n    m, n = len(s1), len(s2)\n    if m + n != len(s3):\n        return False\n    ok = [[False] * (n + 1) for _ in range(m + 1)]\n    ok[0][0] = ____\n    for i in range(m + 1):\n        for j in range(n + 1):\n            if i > 0 and ____ and s1[i - 1] == s3[i + j - 1]:\n                ok[i][j] = True\n            if j > 0 and ____ and ____:\n                ok[i][j] = True\n    return ok[m][n]"},
		},
		Solution: Solution{
			Code: `def isInterleave(s1, s2, s3):
    m, n = len(s1), len(s2)
    if m + n != len(s3):
        return False
    # ok[i][j] = can s1[:i] and s2[:j] interleave into s3[:i + j]?
    ok = [[False] * (n + 1) for _ in range(m + 1)]
    ok[0][0] = True                       # nothing from either makes the empty string
    for i in range(m + 1):
        for j in range(n + 1):
            if i > 0 and ok[i - 1][j] and s1[i - 1] == s3[i + j - 1]:
                ok[i][j] = True           # last letter came from s1
            if j > 0 and ok[i][j - 1] and s2[j - 1] == s3[i + j - 1]:
                ok[i][j] = True           # last letter came from s2
    return ok[m][n]`,
			Explanation: "A two-string table again, like LCS (problem 2), but now the cell is a yes/no and both prefixes are being matched against a third string. The neat fact is that the position in s3 is not a third index: it is always i + j. Each cell asks 'did the last letter come from s1 or from s2?', which is the same 'which side moved last' question as LCS's skip-up or skip-left.",
			BruteForce:  "Walk s3 from the left; whenever the next letter matches the next unused letter of s1 or of s2, try each option recursively. When both match (as with long runs of the same letter) the search branches, and it can revisit the same (i, j) along exponentially many paths: for s1 = 'a' * 40 + 'b', s2 = 'a' * 40, s3 = 'a' * 80 + 'c' there are about C(80, 40), roughly 10^23, ways to split the a's before every one of them fails at the end.",
			KeyInsight:  "Whether the rest can be finished depends only on how many letters of each string are used, (i, j), not on the order you used them in. There are (m + 1) x (n + 1) such pairs, at most 101 x 101 here, and each cell needs two lookups.",
			WhyItWorks:  "State: ok[i][j] = True if s1[:i] and s2[:j] can be interleaved into s3[:i + j]. Recurrence: in any such interleaving, the last letter s3[i + j - 1] is either s1's last letter s1[i - 1] (then the rest is an interleaving of s1[:i - 1] and s2[:j], i.e. ok[i - 1][j]) or s2's last letter s2[j - 1] (then ok[i][j - 1]). These two cases cover every interleaving, so 'or' over them is exact. Base: ok[0][0] = True (empty from empty). The guards i > 0 and j > 0 make row 0 and column 0 follow the same rule with only one possible source. Order: ok[i - 1][j] is in the previous row and ok[i][j - 1] to the left, both already filled. The length check up front guarantees that s3[i + j - 1] is always a valid index and that ok[m][n] means 'all of s3'.",
			CommonMistakes: []string{
				"Greedy: take the next letter from s1 whenever it matches, otherwise from s2. When both match, the choice matters. For s1 = 'ab', s2 = 'ac', s3 = 'acab', greedy takes s1's 'a' first and then can't place the 'c', returning False, but taking s2's 'a' first works. The table tries both.",
				"Skipping the length check. Without it, s3[i + j - 1] can be out of range when s3 is shorter, and when s3 is longer, ok[m][n] can be True even though the end of s3 was never matched. For s1 = 'a', s2 = '', s3 = '' the answer must be False.",
				"Using a third index for the position in s3 (a 3D table). It is always i + j, so a third dimension only adds memory and chances for inconsistent states.",
			},
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Check the lengths", Explanation: "Every letter of s1 and s2 is used exactly once and nothing is added, so the lengths must add up. This also keeps s3[i + j - 1] in range below.", CodeSnippet: "if m + n != len(s3):\n    return False", LineStart: 3, LineEnd: 4},
				{Title: "Define the state", Explanation: "ok[i][j] answers the question for the prefixes s1[:i] and s2[:j]; the matching part of s3 is always s3[:i + j].", CodeSnippet: "ok = [[False] * (n + 1) for _ in range(m + 1)]", LineStart: 5, LineEnd: 6},
				{Title: "Set the base cases", Explanation: "Two empty prefixes make the empty prefix of s3. The rest of row 0 and column 0 is filled by the same loop, where only one source is possible.", CodeSnippet: "ok[0][0] = True", LineStart: 7, LineEnd: 7},
				{Title: "Choose the evaluation order", Explanation: "Rows top to bottom, columns left to right: the cell above and the cell to the left are ready when (i, j) is computed. Starting both loops at 0 fills the edges too.", CodeSnippet: "for i in range(m + 1):\n    for j in range(n + 1):", LineStart: 8, LineEnd: 9},
				{Title: "Write the recurrence", Explanation: "The last letter of s3[:i + j] came from s1 (look up, compare s1[i - 1]) or from s2 (look left, compare s2[j - 1]). Either one is enough.", CodeSnippet: "if i > 0 and ok[i - 1][j] and s1[i - 1] == s3[i + j - 1]:\n    ok[i][j] = True\nif j > 0 and ok[i][j - 1] and s2[j - 1] == s3[i + j - 1]:\n    ok[i][j] = True", LineStart: 10, LineEnd: 13},
				{Title: "Locate the answer", Explanation: "All of s1 and all of s2 used, which (after the length check) is all of s3.", CodeSnippet: "return ok[m][n]", LineStart: 14, LineEnd: 14},
			},
			Alternatives: []AltSolution{
				{
					Name: "Optimise space: one row",
					Code: `def isInterleave(s1, s2, s3):
    m, n = len(s1), len(s2)
    if m + n != len(s3):
        return False
    ok = [False] * (n + 1)                # one row of the table, reused
    for i in range(m + 1):
        for j in range(n + 1):
            if i == 0 and j == 0:
                ok[j] = True
                continue
            from_s1 = i > 0 and ok[j] and s1[i - 1] == s3[i + j - 1]       # ok[j] is still row i - 1
            from_s2 = j > 0 and ok[j - 1] and s2[j - 1] == s3[i + j - 1]   # ok[j - 1] is row i
            ok[j] = from_s1 or from_s2
    return ok[n]`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(n)",
					Explanation:     "A cell reads the cell above and the cell to the left, so one row updated left to right is enough: before ok[j] is overwritten it still holds row i - 1 (up), and ok[j - 1] was just written for row i (left).",
				},
				{
					Name: "Top-down: the brute-force recursion plus a cache",
					Code: `from functools import lru_cache

def isInterleave(s1, s2, s3):
    if len(s1) + len(s2) != len(s3):
        return False

    @lru_cache(maxsize=None)    # without the cache this branches exponentially
    def can(i, j):              # can s1[i:] and s2[j:] interleave into s3[i + j:]?
        if i == len(s1) and j == len(s2):
            return True
        k = i + j
        if i < len(s1) and s1[i] == s3[k] and can(i + 1, j):
            return True
        return j < len(s2) and s2[j] == s3[k] and can(i, j + 1)

    return can(0, 0)`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "can(i, j) is the branching search on suffixes: take s3's next letter from s1 or from s2. Without the cache the long run of a's makes it exponential and the hidden test times out; with it each (i, j) is answered once.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why don't we need a third index for the position in s3?", Answer: "Every letter of s3 produced so far came from exactly one of the two strings, so after using i letters of s1 and j letters of s2 we have produced exactly i + j letters of s3. The position k = i + j is determined by (i, j); a third index would only allow inconsistent combinations that we would have to rule out."},
			{Prompt: "For s1 = 'ab', s2 = 'ac', s3 = 'acab', which cells of the table are True?", Answer: "Row 0 uses only s2: ok[0][0] = T, ok[0][1] = T ('a'), ok[0][2] = T ('ac'). Column 0 uses only s1: ok[1][0] = T ('a'), ok[2][0] = F ('ab' is not 'ac'). ok[1][1] (s3 prefix 'ac', one letter from each) = F: neither 'a' can be the final 'c'. ok[1][2] ('aca') = T: up is ok[0][2] = T and s1[0] = 'a' matches the third letter. ok[2][1] = F (both neighbours it could come from are F). ok[2][2] ('acab') = T: up is ok[1][2] = T and s1[1] = 'b' matches the last letter. The True path is s2's a, s2's c, s1's a, s1's b; greedy would have taken s1's a first and got stuck at ok[1][1]."},
		},
	},
	{
		ID:              "edit-distance",
		Number:          72,
		Order:           7,
		Title:           "Edit Distance",
		Slug:            "edit-distance",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{16, 15},
		Description: `Given two strings word1 and word2, return the smallest number of single-letter edits that turn word1 into word2. One edit is one of:

- **Insert** one letter anywhere in the word.
- **Delete** one letter.
- **Replace** one letter with a different letter.

Either word may be empty: turning "" into "abc" takes 3 inserts, and turning two equal words into each other takes 0 edits.`,
		Constraints: []string{
			"0 <= word1.length, word2.length <= 500",
			"word1 and word2 contain only lowercase English letters",
		},
		Examples: []Example{
			{Input: `word1 = "horse", word2 = "ros"`, Output: "3", Explanation: "horse -> rorse (replace h with r) -> rose (delete r) -> ros (delete e)."},
			{Input: `word1 = "intention", word2 = "execution"`, Output: "5", Explanation: "intention -> inention (delete t) -> enention (replace i with e) -> exention (replace n with x) -> exection (replace n with c) -> execution (insert u)."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"word1": "horse", "word2": "ros"}, Expected: 3, Note: "example: replace, delete, delete"},
			{Input: map[string]interface{}{"word1": "intention", "word2": "execution"}, Expected: 5, Note: "example: five edits"},
			{Input: map[string]interface{}{"word1": "", "word2": ""}, Expected: 0, Hidden: true, Note: "both empty: 0"},
			{Input: map[string]interface{}{"word1": "", "word2": "abc"}, Expected: 3, Hidden: true, Note: "word1 empty: insert everything"},
			{Input: map[string]interface{}{"word1": "abc", "word2": ""}, Expected: 3, Hidden: true, Note: "word2 empty: delete everything"},
			{Input: map[string]interface{}{"word1": "abc", "word2": "abc"}, Expected: 0, Hidden: true, Note: "already equal: 0"},
			{Input: map[string]interface{}{"word1": "a", "word2": "b"}, Expected: 1, Hidden: true, Note: "one replace"},
			{Input: map[string]interface{}{"word1": "ab", "word2": "ba"}, Expected: 2, Hidden: true, Note: "swap costs 2 (no swap operation)"},
			{Input: map[string]interface{}{"word1": "kitten", "word2": "sitting"}, Expected: 3, Hidden: true, Note: "classic: 3 edits"},
			{Input: map[string]interface{}{"word1": "sea", "word2": "eat"}, Expected: 2, Hidden: true, Note: "delete at the front, insert at the back"},
			{Input: map[string]interface{}{"word1": dp2dLetters(200, "abc", 4), "word2": dp2dLetters(180, "abc", 5)}, Expected: 79, Hidden: true, Note: "200 x 180 letters: needs the table"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def minDistance(word1, word2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Think about the last letter of word1 and the last letter of word2. If they are equal, do you need to touch them? If they differ, what could the last edit that fixes the end of the word have been?"},
			{Level: 2, Type: "insight", Content: "Let dist[i][j] = fewest edits turning word1[:i] into word2[:j]. If the last letters match, they cost nothing: dist[i - 1][j - 1]. Otherwise pay 1 for one of three moves and solve the rest: delete word1's letter (dist[i - 1][j]), insert word2's letter (dist[i][j - 1]) or replace one with the other (dist[i - 1][j - 1]). It is the LCS table from problem 2 with costs, the 'match or skip which side' pattern from Chapter 16.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Define the state: dist[i][j] = fewest edits from word1[:i] to word2[:j]; the table is (m + 1) x (n + 1).\n2. Set the base cases: dist[i][0] = i (delete everything), dist[0][j] = j (insert everything).\n3. Write the recurrence: match -> diagonal; mismatch -> 1 + min(up = delete, left = insert, diagonal = replace).\n4. Choose the evaluation order: rows top to bottom, columns left to right.\n5. Locate the answer: dist[m][n]."},
			{Level: 4, Type: "skeleton", Content: "def minDistance(word1, word2):\n    m, n = len(word1), len(word2)\n    dist = [[0] * (n + 1) for _ in range(m + 1)]\n    for i in range(m + 1):\n        dist[i][0] = ____\n    for j in range(n + 1):\n        dist[0][j] = ____\n    for i in range(1, m + 1):\n        for j in range(1, n + 1):\n            if word1[i - 1] == word2[j - 1]:\n                dist[i][j] = ____\n            else:\n                dist[i][j] = 1 + min(____, ____, ____)\n    return dist[m][n]"},
		},
		Solution: Solution{
			Code: `def minDistance(word1, word2):
    m, n = len(word1), len(word2)
    # dist[i][j] = fewest edits turning word1[:i] into word2[:j]
    dist = [[0] * (n + 1) for _ in range(m + 1)]
    for i in range(m + 1):
        dist[i][0] = i                    # delete all i letters
    for j in range(n + 1):
        dist[0][j] = j                    # insert all j letters
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if word1[i - 1] == word2[j - 1]:
                dist[i][j] = dist[i - 1][j - 1]       # match: free
            else:
                dist[i][j] = 1 + min(
                    dist[i - 1][j],       # delete word1[i - 1]
                    dist[i][j - 1],       # insert word2[j - 1]
                    dist[i - 1][j - 1],   # replace one with the other
                )
    return dist[m][n]`,
			Explanation: "Edit distance is LCS (problem 2) turned into a cost: the same (m + 1) x (n + 1) table over pairs of prefixes, the same three neighbours, but a minimum instead of a maximum, a third option on a mismatch (replace = the diagonal), and edges that count edits instead of being 0. It also uses both empty-string edges in a way LCS didn't: word1 or word2 can be empty here.",
			BruteForce:  "Recursively try all three edits at every mismatch. Each call can branch three ways, and the depth is up to m + n, so the work is up to about 3^(m + n); for two 10-letter words that is already billions of calls. The same pair (i, j) of 'letters left to fix' is reached by many different edit sequences (delete-then-insert reaches the same cell as insert-then-delete), and the brute force solves it again each time.",
			KeyInsight:  "Only the remaining prefixes matter for the rest of the work, and a prefix pair is just (i, j). With (m + 1) x (n + 1) cells and three lookups each, two 500-letter words need about 250,000 cells instead of an astronomical search.",
			WhyItWorks:  "State: dist[i][j] = fewest edits turning word1[:i] into word2[:j]. Look at the last letters a = word1[i - 1] and b = word2[j - 1] in an optimal edit script. If a == b, leaving them alone is never worse, so dist[i][j] = dist[i - 1][j - 1]. If a != b, the end of the word must be fixed by one of: deleting a (then turn word1[:i - 1] into word2[:j]: up), inserting b at the end (then turn word1[:i] into word2[:j - 1]: left), or replacing a with b (then word1[:i - 1] into word2[:j - 1]: diagonal). Every optimal script ends in one of these, so 1 + min of the three is exact. Base: turning a prefix of length i into the empty string takes i deletes, and the empty string into a prefix of length j takes j inserts. For 'horse' -> 'ros' the last row of the table is 5 4 4 3, so the answer is 3.",
			CommonMistakes: []string{
				"Leaving the edges at 0 (copying LCS's base cases). dist[i][0] must be i: turning 'abc' into '' takes 3 deletes. With zero edges, 'abc' -> '' returns 0 and 'horse' -> 'ros' returns too small a number.",
				"Mixing up which neighbour is insert and which is delete. It doesn't change the number (both cost 1), but it matters when you rebuild the edit script: up (i - 1, j) consumed a letter of word1 without producing one of word2, which is a delete; left consumed a letter of word2, which is an insert.",
				"Charging for a match: using 1 + min(up, left, diagonal) for every cell, whether or not the letters are equal. Equal letters need no edit, so the cell must copy the diagonal for free. With the charge, 'abc' -> 'abc' returns 3 instead of 0 and 'horse' -> 'ros' returns 5 instead of 3.",
			},
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "One cell per pair of prefixes, with an extra row and column for the empty prefix.", CodeSnippet: "dist = [[0] * (n + 1) for _ in range(m + 1)]", LineStart: 3, LineEnd: 4},
				{Title: "Set the base cases", Explanation: "Column 0: delete all i letters. Row 0: insert all j letters. These edges are what make empty words work.", CodeSnippet: "dist[i][0] = i\ndist[0][j] = j", LineStart: 5, LineEnd: 8},
				{Title: "Choose the evaluation order", Explanation: "Up, left and diagonal are all above or to the left, so row by row, left to right.", CodeSnippet: "for i in range(1, m + 1):\n    for j in range(1, n + 1):", LineStart: 9, LineEnd: 10},
				{Title: "Write the recurrence", Explanation: "Equal last letters: free, take the diagonal. Otherwise pay 1 and take the cheapest of delete (up), insert (left) and replace (diagonal).", CodeSnippet: "if word1[i - 1] == word2[j - 1]:\n    dist[i][j] = dist[i - 1][j - 1]\nelse:\n    dist[i][j] = 1 + min(", LineStart: 11, LineEnd: 18},
				{Title: "Locate the answer", Explanation: "Both full words: the bottom-right cell. For two empty words the loops never run and dist[0][0] = 0.", CodeSnippet: "return dist[m][n]", LineStart: 19, LineEnd: 19},
			},
			Alternatives: []AltSolution{
				{
					Name: "Optimise space: two rows",
					Code: `def minDistance(word1, word2):
    n = len(word2)
    prev = list(range(n + 1))             # row 0: j inserts
    for i in range(1, len(word1) + 1):
        cur = [i] + [0] * n               # column 0: i deletes
        for j in range(1, n + 1):
            if word1[i - 1] == word2[j - 1]:
                cur[j] = prev[j - 1]
            else:
                cur[j] = 1 + min(prev[j], cur[j - 1], prev[j - 1])
        prev = cur
    return prev[n]`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Row i reads only row i - 1 and itself. prev starts as row 0 (0, 1, ..., n) and each new row starts with its column-0 value i. Keep the 2D table when you need to rebuild the actual edits.",
				},
				{
					Name: "Top-down: the brute-force recursion plus a cache",
					Code: `from functools import lru_cache

def minDistance(word1, word2):
    @lru_cache(maxsize=None)    # without the cache: up to 3 branches per call
    def edits(i, j):            # fewest edits turning word1[i:] into word2[j:]
        if i == len(word1):
            return len(word2) - j         # insert the rest
        if j == len(word2):
            return len(word1) - i         # delete the rest
        if word1[i] == word2[j]:
            return edits(i + 1, j + 1)
        return 1 + min(edits(i + 1, j), edits(i, j + 1), edits(i + 1, j + 1))

    return edits(0, 0)`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "edits(i, j) works on suffixes: when one word runs out, the rest of the other must be inserted or deleted. Without the cache it branches three ways on every mismatch and the 200-letter test times out; with it each (i, j) is solved once.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does 'ab' -> 'ba' cost 2 and not 1?", Answer: "There is no swap operation. Each edit changes one position: you can replace both letters (2 edits), or delete the a and insert it at the end (2 edits). No single insert, delete or replace turns 'ab' into 'ba', so the minimum is 2. The table finds this because the last letters b and a differ and every neighbour costs at least 1."},
			{Prompt: "How is edit distance related to LCS when only inserts and deletes are allowed?", Answer: "Without replace, every letter that is not part of a common subsequence must be deleted from word1 or inserted from word2, so the cost is m + n - 2 * LCS(word1, word2). Replace is what makes edit distance different: it fixes a mismatched pair for 1 instead of 2."},
			{Prompt: "How would you list the actual edits, not just count them?", Answer: "Keep the full table and walk back from (m, n). If the letters match and dist[i][j] == dist[i - 1][j - 1], move diagonally with no edit. Otherwise find which neighbour satisfies dist[i][j] == 1 + neighbour: up means delete word1[i - 1], left means insert word2[j - 1], diagonal means replace. Stop at (0, 0); the edges give the remaining deletes or inserts."},
		},
	},
	{
		ID:              "longest-increasing-path-in-matrix",
		Number:          329,
		Order:           8,
		Title:           "Longest Increasing Path in a Matrix",
		Slug:            "longest-increasing-path-in-a-matrix",
		Difficulty:      "Hard",
		Category:        "2d-dp",
		Tags:            []string{"Matrix", "Depth-First Search", "Memoization", "Dynamic Programming", "Topological Sort"},
		RelatedChapters: []int{16, 13, 15},
		Description: `You are given an m x n grid of integers, matrix. A path moves from a cell to one of its four neighbours (up, down, left or right; no diagonals, and never outside the grid). A path is **increasing** if every step goes to a strictly larger value.

Return the number of cells on the longest increasing path. A path may start at any cell, and a single cell is a path of length 1, so the answer is at least 1.`,
		Constraints: []string{
			"m == matrix.length",
			"n == matrix[i].length",
			"1 <= m, n <= 200",
			"0 <= matrix[i][j] <= 2^31 - 1",
		},
		Examples: []Example{
			{Input: "matrix = [[9,9,4],[6,6,8],[2,1,1]]", Output: "4", Explanation: "The path 1 -> 2 -> 6 -> 9 (bottom middle, bottom left, middle left, top left) has 4 cells."},
			{Input: "matrix = [[3,4,5],[3,2,6],[2,2,1]]", Output: "4", Explanation: "3 -> 4 -> 5 -> 6 along the top row and down the right side. Moving to an equal value (2 -> 2) is not allowed."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}}, Expected: 4, Note: "example: 1 -> 2 -> 6 -> 9"},
			{Input: map[string]interface{}{"matrix": [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}}, Expected: 4, Note: "example: 3 -> 4 -> 5 -> 6"},
			{Input: map[string]interface{}{"matrix": [][]int{{1}}}, Expected: 1, Hidden: true, Note: "single cell"},
			{Input: map[string]interface{}{"matrix": [][]int{{7, 7}, {7, 7}}}, Expected: 1, Hidden: true, Note: "all equal: strictly increasing means length 1"},
			{Input: map[string]interface{}{"matrix": [][]int{{3, 2, 1, 2, 3}}}, Expected: 3, Hidden: true, Note: "single row, valley in the middle"},
			{Input: map[string]interface{}{"matrix": [][]int{{1}, {2}, {3}, {4}}}, Expected: 4, Hidden: true, Note: "single column, all increasing"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}}, Expected: 9, Hidden: true, Note: "spiral: the path visits every cell"},
			{Input: map[string]interface{}{"matrix": dp2dDiagonal(20)}, Expected: 39, Hidden: true, Note: "20 x 20 with value r + c: plain DFS explores billions of paths"},
			{Input: map[string]interface{}{"matrix": dp2dSnake(30)}, Expected: 900, Hidden: true, Note: "30 x 30 snake: one path of 900 cells"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def longestIncreasingPath(matrix):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you already knew the longest increasing path that starts at each of a cell's larger neighbours, what would the longest path starting at the cell itself be? And can a path ever come back to a cell it has visited?"},
			{Level: 2, Type: "insight", Content: "longest[r][c] = 1 + max(longest of each strictly larger neighbour), or 1 if there is none. Because every step goes to a larger value, a path can never cycle back, so this recursion always ends. Rows and columns don't give a fill order (a cell can depend on the cell above it OR below it), so let a DFS find the order and memoise each cell's answer: Chapter 16's memoised DFS, the top-down form of the recipe.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Define the state: longest[r][c] = length of the longest increasing path that starts at (r, c).\n2. Write the recurrence: 1 + max over the in-bounds neighbours with a strictly larger value.\n3. Set the base cases: a cell with no larger neighbour has longest = 1 (a local maximum).\n4. Choose the evaluation order: let the recursion decide; store each result in a memo table so each cell is computed once.\n5. Locate the answer: the maximum of longest[r][c] over all cells."},
			{Level: 4, Type: "skeleton", Content: "def longestIncreasingPath(matrix):\n    rows, cols = len(matrix), len(matrix[0])\n    longest = [[0] * cols for _ in range(rows)]\n\n    def dfs(r, c):\n        if longest[r][c]:\n            return ____\n        best = 1\n        for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):\n            nr, nc = r + dr, c + dc\n            if 0 <= nr < rows and 0 <= nc < cols and ____:\n                best = max(best, ____)\n        longest[r][c] = best\n        return best\n\n    return ____"},
		},
		Solution: Solution{
			Code: `def longestIncreasingPath(matrix):
    rows, cols = len(matrix), len(matrix[0])
    # longest[r][c] = length of the longest increasing path that STARTS at (r, c)
    longest = [[0] * cols for _ in range(rows)]   # 0 = not computed yet

    def dfs(r, c):
        if longest[r][c]:
            return longest[r][c]                  # memo hit: already final
        best = 1                                  # the path that is just this cell
        for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and matrix[nr][nc] > matrix[r][c]:
                best = max(best, 1 + dfs(nr, nc))
        longest[r][c] = best
        return best

    return max(dfs(r, c) for r in range(rows) for c in range(cols))`,
			Explanation: "The table is still one cell per grid cell, as in Unique Paths (problem 1), but now a cell can depend on neighbours in all four directions, so there is no row-by-row order. The new idea: when the dependencies form no cycles, a depth-first search with a memo computes every cell after the cells it needs, whatever their position. It joins Chapter 13's DFS on a grid with Chapter 15's memoization.",
			BruteForce:  "Start a plain DFS from every cell and explore every increasing path, keeping the longest. The number of increasing paths can be exponential: in a 20 x 20 grid whose value is r + c, every right/down route from the top-left is increasing, and there are C(38, 19), about 3.5 x 10^10 of them. The DFS recomputes 'longest path from this cell' every time it arrives there along a different route.",
			KeyInsight:  "The longest increasing path that starts at a cell does not depend on how you reached the cell: the path can never revisit earlier cells anyway, because they are smaller. So the answer for each cell is a fixed number that can be cached, and there are only m * n cells.",
			WhyItWorks:  "State: longest[r][c] = number of cells on the longest increasing path starting at (r, c). Recurrence: such a path is either just (r, c), or (r, c) followed by a longest path starting at some neighbour with a larger value, so longest[r][c] = 1 + max over larger neighbours (1 if none). No cycles: along any chain of dependencies the values strictly increase, so a cell can never (even indirectly) depend on itself, and the recursion always reaches a local maximum (a base case). This is a DAG: edges go from smaller to larger values. Order: the DFS computes a cell only after the calls for all its larger neighbours have returned, which is a valid order (a reverse topological order of the DAG). Memo: once longest[r][c] is stored it is final, so every later visit is O(1); each cell does at most 4 neighbour checks, giving O(m * n) overall. The answer is the maximum over all starting cells.",
			CommonMistakes: []string{
				"Allowing equal neighbours (>= instead of >). The path must be strictly increasing, and with >= two equal neighbours call each other forever (a cycle), so the recursion never ends or the memo returns half-computed values. [[7, 7], [7, 7]] must return 1.",
				"Adding a 'visited' set as in Chapter 13's DFS and clearing it between starts. It isn't needed (strictly increasing paths can't revisit a cell) and it breaks the memo: a cell's answer would then depend on which cells are 'visited' at the time, so it could no longer be cached.",
				"Only starting from one cell (say the smallest value) or only from the corners. The longest path can start anywhere; take the max over all cells. The memo makes the extra starts cheap.",
			},
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "One number per cell: the longest increasing path that starts there. 0 marks 'not computed yet', which is safe because every real answer is at least 1.", CodeSnippet: "longest = [[0] * cols for _ in range(rows)]", LineStart: 3, LineEnd: 4},
				{Title: "Choose the evaluation order", Explanation: "No row order works, so the recursion decides: dfs(r, c) first asks for its larger neighbours, then fills its own cell. The memo check makes each cell's work happen once.", CodeSnippet: "def dfs(r, c):\n    if longest[r][c]:\n        return longest[r][c]", LineStart: 6, LineEnd: 8},
				{Title: "Set the base cases", Explanation: "best starts at 1, the path that is just this cell. If no neighbour is larger (a local maximum), that is the answer.", CodeSnippet: "best = 1", LineStart: 9, LineEnd: 9},
				{Title: "Write the recurrence", Explanation: "For each in-bounds neighbour with a strictly larger value, extend its best path by this cell. Store the result before returning so no one computes it again.", CodeSnippet: "if 0 <= nr < rows and 0 <= nc < cols and matrix[nr][nc] > matrix[r][c]:\n    best = max(best, 1 + dfs(nr, nc))\nlongest[r][c] = best", LineStart: 10, LineEnd: 15},
				{Title: "Locate the answer", Explanation: "A path may start anywhere, so take the maximum over all cells. For the first example the memo ends as [[1,1,2],[2,2,1],[3,4,2]] and the maximum is 4, at the cell holding 1.", CodeSnippet: "return max(dfs(r, c) for r in range(rows) for c in range(cols))", LineStart: 17, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{
					Name: "Bottom-up: fill cells from the largest value down",
					Code: `def longestIncreasingPath(matrix):
    rows, cols = len(matrix), len(matrix[0])
    # fill order = biggest value first, so every larger neighbour is already final
    cells = sorted(((matrix[r][c], r, c) for r in range(rows) for c in range(cols)), reverse=True)
    longest = [[1] * cols for _ in range(rows)]
    for value, r, c in cells:
        for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and matrix[nr][nc] > value:
                longest[r][c] = max(longest[r][c], 1 + longest[nr][nc])
    return max(max(row) for row in longest)`,
					TimeComplexity:  "O(m * n * log(m * n))",
					SpaceComplexity: "O(m * n)",
					Explanation:     "The evaluation order made explicit: sort the cells by value, largest first. When a cell is processed, every strictly larger neighbour was processed earlier, so its longest value is final. No recursion, so no depth limit (the memoised DFS can go m * n calls deep on a snake-shaped grid), at the cost of the sort.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can't you fill this table row by row like Unique Paths?", Answer: "In Unique Paths a cell depends only on the cell above and the cell to its left, so row-major order always has them ready. Here a cell depends on whichever neighbours are larger, which may be below or to the right; row-major order would read cells that haven't been computed. The correct order is 'larger values first', which the DFS discovers on its own and the sorted alternative makes explicit."},
			{Prompt: "Why does the memo stay correct even though different starting cells reach the same cell along different routes?", Answer: "longest[r][c] depends only on the values in the grid, not on the route used to get there: the path from (r, c) onwards can only visit larger values, so it can never run into the cells that came before (they are smaller). The same question always has the same answer, which is exactly the condition for caching."},
			{Prompt: "On a 200 x 200 snake-shaped grid the memoised DFS can recurse 40,000 calls deep. What can you do about that in Python?", Answer: "Use the bottom-up alternative (sort cells by value, largest first), which has no recursion at all. Alternatively, run Kahn's topological sort from Chapter 13 by peeling local maxima layer by layer and counting layers. Raising the recursion limit is fragile because the interpreter's C stack can still overflow."},
		},
	},
	{
		ID:              "distinct-subsequences",
		Number:          115,
		Order:           9,
		Title:           "Distinct Subsequences",
		Slug:            "distinct-subsequences",
		Difficulty:      "Hard",
		Category:        "2d-dp",
		Tags:            []string{"String", "Dynamic Programming"},
		RelatedChapters: []int{16, 15},
		Description: `Given two strings s and t, count the number of ways to pick letters of s, keeping their order, that spell exactly t. Two ways are different if they use a different set of **positions** in s, even if the letters are the same.

For example, s = "aab", t = "ab" has 2 ways: positions (0, 2) and (1, 2).

Letters are case-sensitive ("A" and "a" are different). If t is longer than s, or can't be spelled, return 0. The answer fits in a signed 32-bit integer.`,
		Constraints: []string{
			"1 <= s.length, t.length <= 1000",
			"s and t contain only English letters (upper or lower case)",
			"The answer fits in a signed 32-bit integer",
		},
		Examples: []Example{
			{Input: `s = "rabbbit", t = "rabbit"`, Output: "3", Explanation: "s has three b's and t needs two of them; there are 3 ways to choose which two, and the rest of t has one choice each."},
			{Input: `s = "babgbag", t = "bag"`, Output: "5", Explanation: "Using positions (0,1,3), (0,1,6), (0,5,6), (2,5,6) and (4,5,6) of \"babgbag\"."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "rabbbit", "t": "rabbit"}, Expected: 3, Note: "example: pick 2 of the 3 b's"},
			{Input: map[string]interface{}{"s": "babgbag", "t": "bag"}, Expected: 5, Note: "example: five ways"},
			{Input: map[string]interface{}{"s": "a", "t": "aa"}, Expected: 0, Hidden: true, Note: "t longer than s: 0"},
			{Input: map[string]interface{}{"s": "abc", "t": "abc"}, Expected: 1, Hidden: true, Note: "equal strings: one way"},
			{Input: map[string]interface{}{"s": "abc", "t": "d"}, Expected: 0, Hidden: true, Note: "letter not in s: 0"},
			{Input: map[string]interface{}{"s": "aaa", "t": "a"}, Expected: 3, Hidden: true, Note: "choose 1 of 3"},
			{Input: map[string]interface{}{"s": "aaaa", "t": "aa"}, Expected: 6, Hidden: true, Note: "choose 2 of 4 = 6"},
			{Input: map[string]interface{}{"s": "Aa", "t": "a"}, Expected: 1, Hidden: true, Note: "upper and lower case are different letters"},
			{Input: map[string]interface{}{"s": "abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc", "t": "abcabcabc"}, Expected: 3124550, Hidden: true, Note: "60 letters: too many subsets to try"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def numDistinct(s, t):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Look at the last letter of s. Either it is used to spell t, or it isn't. When can it be used, and what is left to spell in each case?"},
			{Level: 2, Type: "insight", Content: "ways[i][j] = number of ways to spell t[:j] from s[:i]. The last letter s[i - 1] is either skipped (ways[i - 1][j]) or, if it equals t[j - 1], used as the last letter of t (ways[i - 1][j - 1]). Unlike LCS, a match doesn't force you to use it, so on a match you ADD both options. Chapter 16's two-string table, counting instead of maximising.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Define the state: ways[i][j] = ways to spell t[:j] using positions from s[:i].\n2. Set the base cases: ways[i][0] = 1 for every i (spell the empty string by picking nothing); ways[0][j] = 0 for j > 0.\n3. Write the recurrence: ways[i][j] = ways[i - 1][j] + (ways[i - 1][j - 1] if s[i - 1] == t[j - 1]).\n4. Choose the evaluation order: rows top to bottom; each cell reads only the row above.\n5. Locate the answer: ways[m][n].\n6. Optimise space: one row, j going backwards."},
			{Level: 4, Type: "skeleton", Content: "def numDistinct(s, t):\n    m, n = len(s), len(t)\n    ways = [[0] * (n + 1) for _ in range(m + 1)]\n    for i in range(m + 1):\n        ways[i][0] = ____\n    for i in range(1, m + 1):\n        for j in range(1, n + 1):\n            ways[i][j] = ____\n            if s[i - 1] == t[j - 1]:\n                ways[i][j] += ____\n    return ways[m][n]"},
		},
		Solution: Solution{
			Code: `def numDistinct(s, t):
    m, n = len(s), len(t)
    # ways[i][j] = ways to pick t[:j] as a subsequence of s[:i]
    ways = [[0] * (n + 1) for _ in range(m + 1)]
    for i in range(m + 1):
        ways[i][0] = 1                    # empty t: pick nothing
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            ways[i][j] = ways[i - 1][j]   # skip s[i - 1]
            if s[i - 1] == t[j - 1]:
                ways[i][j] += ways[i - 1][j - 1]   # or use it for t[j - 1]
    return ways[m][n]`,
			Explanation: "The counting version of the two-string table. It builds on LCS (problem 2) and Edit Distance (problem 7), which also asked 'what happens to the last letters?', with two differences: only s can skip letters (t must be spelled completely), and on a match both options are valid, so we add them instead of picking the better one. The 0/1-style 'read the previous row' also links it to Target Sum (problem 5): each letter of s is used at most once.",
			BruteForce:  "Try every way to choose len(t) positions of s (C(m, n) of them) and check whether they spell t, or recurse 'use this letter or skip it'. For s = 'abcabc' * 10 and t = 'abcabcabc' there are C(60, 9), about 15 billion, position sets. The recursion re-solves 'spell t[j:] from s[i:]' for every way of having spelled t[:j] from the letters before i.",
			KeyInsight:  "The number of ways to finish depends only on how much of s is left and how much of t is left, the pair (i, j). Counting per pair, instead of listing every choice of positions, needs (m + 1) x (n + 1) cells.",
			WhyItWorks:  "State: ways[i][j] = number of position sets in s[:i] that spell t[:j]. Split these sets by whether they use position i - 1. Sets that don't use it are exactly the sets for (i - 1, j). Sets that do use it must pair s[i - 1] with the last letter t[j - 1] (positions keep their order, and i - 1 is the last available one), which requires s[i - 1] == t[j - 1]; removing that position leaves a set for (i - 1, j - 1). The two groups are disjoint and complete, so we add. Base: the empty t is spelled once (the empty set) from any prefix; a non-empty t can't be spelled from an empty s. For 'aab', 'ab': rows are i=0: 1 0 0; i=1 ('a'): 1 1 0; i=2 ('aa'): 1 2 0; i=3 ('aab'): 1 2 2. Answer 2.",
			CommonMistakes: []string{
				"Copying LCS: on a match, taking only the diagonal (forcing the use of s[i - 1]). That loses every way that uses an earlier copy of the same letter: 'aaa', 'a' returns 1 instead of 3. A match is an option, not an obligation, so add the skip option too.",
				"Setting ways[0][0] = 1 but leaving ways[i][0] = 0 for i > 0. The empty target can be spelled from ANY prefix of s, and every complete count is built on those cells: with them at 0, only the first letter of s can start a spelling, so 'babgbag', 'bag' returns 3 instead of 5 and 'aaa', 'a' returns 1 instead of 3.",
				"In the one-row version, looping j forwards. Then ways[j - 1] has already been updated for this letter of s, so one letter of s can be used for two letters of t: 'a', 'aa' would count 1 instead of 0. Loop j backwards (the 0/1 rule from Chapter 16).",
			},
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "Rows: how much of s is available. Columns: how much of t must be spelled.", CodeSnippet: "ways = [[0] * (n + 1) for _ in range(m + 1)]", LineStart: 3, LineEnd: 4},
				{Title: "Set the base cases", Explanation: "Column 0 is all 1: the empty target is spelled by choosing no positions. Row 0 beyond column 0 stays 0: nothing can be spelled from nothing.", CodeSnippet: "for i in range(m + 1):\n    ways[i][0] = 1", LineStart: 5, LineEnd: 6},
				{Title: "Choose the evaluation order", Explanation: "Each cell reads only the row above (straight up and up-left), so rows top to bottom works.", CodeSnippet: "for i in range(1, m + 1):\n    for j in range(1, n + 1):", LineStart: 7, LineEnd: 8},
				{Title: "Write the recurrence", Explanation: "Skip s[i - 1] (always allowed), plus use it for t[j - 1] when the letters are equal. Counting, so add.", CodeSnippet: "ways[i][j] = ways[i - 1][j]\nif s[i - 1] == t[j - 1]:\n    ways[i][j] += ways[i - 1][j - 1]", LineStart: 9, LineEnd: 11},
				{Title: "Locate the answer", Explanation: "All of s available, all of t spelled. If t is longer than s the cell is simply 0.", CodeSnippet: "return ways[m][n]", LineStart: 12, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Optimise space: one row, j backwards",
					Code: `def numDistinct(s, t):
    n = len(t)
    ways = [1] + [0] * n              # ways[j] for the prefix of s seen so far
    for ch in s:
        for j in range(n, 0, -1):     # backwards: ways[j - 1] is still the old row
            if ch == t[j - 1]:
                ways[j] += ways[j - 1]
    return ways[n]`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Both reads are from the previous row, at j and j - 1. Going backwards over j, ways[j - 1] hasn't been updated for the current letter yet, so it still holds the previous row: each letter of s is used at most once. Same reason as Target Sum's backwards loop.",
				},
				{
					Name: "Top-down: the brute-force recursion plus a cache",
					Code: `from functools import lru_cache

def numDistinct(s, t):
    @lru_cache(maxsize=None)    # without the cache this is the exponential brute force
    def count(i, j):            # ways to pick t[j:] from s[i:]
        if j == len(t):
            return 1
        if len(s) - i < len(t) - j:
            return 0                      # not enough letters left
        total = count(i + 1, j)           # skip s[i]
        if s[i] == t[j]:
            total += count(i + 1, j + 1)  # use s[i]
        return total

    return count(0, 0)`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "count(i, j) is the brute-force 'skip s[i] or use it'. The early return when fewer letters remain in s than in t prunes hopeless branches. Without the cache it explores every position set; with it each (i, j) is solved once.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "On a match, LCS takes only the diagonal but this problem adds diagonal + up. Why the difference?", Answer: "LCS asks for the best length, and using a matching pair is never worse, so the other options can be dropped. Here we count every way, and using s[i - 1] for t[j - 1] and skipping s[i - 1] (to use an earlier equal letter instead) are different position sets that both spell t. Counting needs every disjoint case, so we add."},
			{Prompt: "Why is the answer 0, not an error, when t is longer than s?", Answer: "The table handles it: row 0 is 0 beyond column 0, and each row can increase the number of spelled letters by at most one, so cell (i, j) is 0 whenever j > i. ways[m][n] with n > m is therefore 0 without any special case."},
		},
	},
	{
		ID:              "burst-balloons",
		Number:          312,
		Order:           10,
		Title:           "Burst Balloons",
		Slug:            "burst-balloons",
		Difficulty:      "Hard",
		Category:        "2d-dp",
		Tags:            []string{"Array", "Dynamic Programming", "Interval DP"},
		RelatedChapters: []int{16, 15, 8},
		Description: `A row of n balloons has a number painted on each: nums[i] for balloon i. You must burst all of them, one at a time, in any order you choose.

When you burst a balloon you earn (value of its current left neighbour) x (its own value) x (value of its current right neighbour), where "current neighbours" are the nearest balloons that have **not** been burst yet. If there is no balloon on a side, use 1 for that side. After a burst the balloon is gone and its two neighbours become adjacent.

Return the largest total number of coins you can earn.`,
		Constraints: []string{
			"n == nums.length",
			"1 <= n <= 300",
			"0 <= nums[i] <= 100",
		},
		Examples: []Example{
			{Input: "nums = [3,1,5,8]", Output: "167", Explanation: "Burst 1: 3*1*5 = 15, row [3,5,8]. Burst 5: 3*5*8 = 120, row [3,8]. Burst 3: 1*3*8 = 24, row [8]. Burst 8: 1*8*1 = 8. Total 15 + 120 + 24 + 8 = 167."},
			{Input: "nums = [1,5]", Output: "10", Explanation: "Burst 1 first (1*1*5 = 5), then 5 (1*5*1 = 5)."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{3, 1, 5, 8}}, Expected: 167, Note: "example: 167"},
			{Input: map[string]interface{}{"nums": []int{1, 5}}, Expected: 10, Note: "two balloons"},
			{Input: map[string]interface{}{"nums": []int{7}}, Expected: 7, Hidden: true, Note: "one balloon: 1 * 7 * 1"},
			{Input: map[string]interface{}{"nums": []int{0}}, Expected: 0, Hidden: true, Note: "a zero balloon earns nothing"},
			{Input: map[string]interface{}{"nums": []int{5, 0, 5}}, Expected: 30, Hidden: true, Note: "burst the zero first"},
			{Input: map[string]interface{}{"nums": []int{9, 76, 64, 21}}, Expected: 116718, Hidden: true, Note: "bursting the biggest first is not best"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 2, 2, 2}}, Expected: 30, Hidden: true, Note: "all equal values"},
			{Input: map[string]interface{}{"nums": dp2dSeq(40, 101, 6)}, Expected: 8382118, Hidden: true, Note: "40 balloons: every burst order is 40! orders"},
		},
		TimeComplexity:  "O(n^3)",
		SpaceComplexity: "O(n^2)",
		StarterCode:     "def maxCoins(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you choose which balloon to burst FIRST, its neighbours change afterwards and the two sides get tangled together. What if you instead choose which balloon in a range is burst LAST? Who are its neighbours at that moment?"},
			{Level: 2, Type: "insight", Content: "Pad the row with a 1 at each end. For two balloons left and right that are still standing, let best[left][right] = most coins from bursting every balloon strictly between them. If k is the last one burst in that gap, its neighbours at that moment are exactly left and right, and the balloons on each side of k were burst independently: best[left][right] = max over k of best[left][k] + vals[left] * vals[k] * vals[right] + best[k][right]. This is interval DP from Chapter 16: the state is a range, filled from short ranges to long ones.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Define the state: vals = [1] + nums + [1]; best[left][right] = most coins from bursting every balloon strictly between positions left and right.\n2. Set the base cases: if right = left + 1 there is nothing between, so best = 0.\n3. Write the recurrence: try every k strictly between as the LAST balloon burst: best[left][k] + vals[left] * vals[k] * vals[right] + best[k][right].\n4. Choose the evaluation order: by gap = right - left, from 2 up to len(vals) - 1, so both halves are shorter and already done.\n5. Locate the answer: best[0][len(vals) - 1], the whole padded row."},
			{Level: 4, Type: "skeleton", Content: "def maxCoins(nums):\n    vals = [1] + nums + [1]\n    n = len(vals)\n    best = [[0] * n for _ in range(n)]\n    for gap in range(2, n):\n        for left in range(n - gap):\n            right = ____\n            for last in range(____, ____):\n                coins = ____\n                if coins > best[left][right]:\n                    best[left][right] = coins\n    return ____"},
		},
		Solution: Solution{
			Code: `def maxCoins(nums):
    vals = [1] + nums + [1]               # walls worth 1 that are never burst
    n = len(vals)
    # best[left][right] = most coins from bursting every balloon strictly between left and right
    best = [[0] * n for _ in range(n)]    # gap 1: nothing between, 0 coins
    for gap in range(2, n):               # short intervals before long ones
        for left in range(n - gap):
            right = left + gap
            for last in range(left + 1, right):
                # last is burst LAST, so its neighbours then are left and right
                coins = best[left][last] + vals[left] * vals[last] * vals[right] + best[last][right]
                if coins > best[left][right]:
                    best[left][right] = coins
    return best[0][n - 1]`,
			Explanation: "Interval DP: the state is a range (left, right) of the row instead of two prefixes, and the table is filled by range length rather than row by row. It builds on the earlier two-index problems (LCS, Edit Distance) for the table and the recipe, and adds the key modelling move of Chapter 16: choose the LAST action inside the range, so that the two sides become independent subproblems.",
			BruteForce:  "Try every order of bursting: n! orders, each simulated in O(n). For the 40-balloon test that is 40!, about 8 x 10^47 orders (10 balloons is already 3.6 million). Choosing the FIRST balloon recursively doesn't split the problem either, because after it is gone its two neighbours touch and the left and right parts influence each other's coins, so the subproblems are whole remaining sets of balloons (2^n of them).",
			KeyInsight:  "Think about the balloon burst LAST in a range. While everything else between left and right is being burst, k is still standing, so it acts as a wall: the balloons on its left only ever see left and k as outer neighbours, and those on its right only see k and right. The two sides are independent intervals, and when k finally bursts its neighbours are exactly left and right.",
			WhyItWorks:  "Pad vals = [1] + nums + [1]; the pads are never burst, so they model 'no neighbour = 1'. State: best[left][right] = most coins from bursting all balloons strictly between left and right, given that left and right are still standing. Recurrence: in any order, some balloon k (left < k < right) is burst last in the gap. Before that, k separates the gap into (left, k) and (k, right); every burst in one part sees only balloons of that part plus its two walls, so the parts are independent and each contributes at most best[left][k] and best[k][right]. Then k earns vals[left] * vals[k] * vals[right]. Taking the max over every k covers every possible order (every order has a last balloon). Base: an empty gap (right = left + 1) earns 0. Order: both subintervals have a smaller gap, so filling by increasing gap has them ready. For [3,1,5,8]: vals = [1,3,1,5,8,1] and the best last balloon for the whole row is 8 (k = 4): best[0][4] + 1*8*1 + best[4][5] = 159 + 8 + 0 = 167.",
			CommonMistakes: []string{
				"Choosing the FIRST balloon to burst in the recurrence: best(l, r) = max over k of nums[k-1] * nums[k] * nums[k+1] + best(l, k-1) + best(k+1, r). After k is burst, its neighbours become adjacent, so the two sides are not independent and the coins in that formula are wrong. Choose the LAST balloon instead.",
				"Filling the table row by row (for left: for right:) instead of by gap. best[left][right] needs best[k][right] for k > left, which is in a LATER row; it would still be 0 and the answer comes out too small. Fill by increasing gap (or let left run downwards from the end).",
				"Greedy: always burst the smallest balloon first (or the biggest last). For [9, 76, 64, 21] no simple greedy rule finds 116718; the interval DP compares every last-balloon choice in every range.",
			},
			TimeComplexity:  "O(n^3)",
			SpaceComplexity: "O(n^2)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "Pad with 1s so every real balloon always has two neighbours. best[left][right] covers the balloons strictly between two standing walls.", CodeSnippet: "vals = [1] + nums + [1]\nbest = [[0] * n for _ in range(n)]", LineStart: 2, LineEnd: 5},
				{Title: "Set the base cases", Explanation: "Adjacent walls (gap 1) have nothing between them, so 0 coins. The zero-filled table already says that; the loop starts at gap 2.", CodeSnippet: "best = [[0] * n for _ in range(n)]", LineStart: 5, LineEnd: 5},
				{Title: "Choose the evaluation order", Explanation: "Short gaps before long ones: every subinterval (left, last) and (last, right) has a smaller gap, so it is already final.", CodeSnippet: "for gap in range(2, n):\n    for left in range(n - gap):\n        right = left + gap", LineStart: 6, LineEnd: 8},
				{Title: "Write the recurrence", Explanation: "Try each balloon as the last one burst between the walls: the two sides plus its own burst with left and right as neighbours. Keep the maximum.", CodeSnippet: "for last in range(left + 1, right):\n    coins = best[left][last] + vals[left] * vals[last] * vals[right] + best[last][right]\n    if coins > best[left][right]:\n        best[left][right] = coins", LineStart: 9, LineEnd: 13},
				{Title: "Locate the answer", Explanation: "The two pads enclose every real balloon.", CodeSnippet: "return best[0][n - 1]", LineStart: 14, LineEnd: 14},
			},
			Alternatives: []AltSolution{
				{
					Name: "Top-down: the same recurrence with a cache",
					Code: `from functools import lru_cache

def maxCoins(nums):
    vals = [1] + nums + [1]

    @lru_cache(maxsize=None)    # without the cache this tries every order
    def best(left, right):      # balloons strictly between left and right
        top = 0
        for last in range(left + 1, right):
            top = max(top, best(left, last) + vals[left] * vals[last] * vals[right] + best(last, right))
        return top

    return best(0, len(vals) - 1)`,
					TimeComplexity:  "O(n^3)",
					SpaceComplexity: "O(n^2)",
					Explanation:     "best(left, right) tries every last balloon and recurses on the two sides; the cache makes it O(n^2) states x O(n) choices. It avoids thinking about the fill order, which is the trickiest part of interval DP. Note there is no cheap 'rolling array' here: a range needs results from many shorter ranges.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does choosing the last balloon make the two sides independent, while choosing the first doesn't?", Answer: "The last balloon k stays standing until everything else in the range is gone, so it is a permanent wall: a balloon on its left can never have a neighbour on its right side, and vice versa. Both sides' coins depend only on their own balloons and their two walls. The first balloon disappears immediately, after which its neighbours touch, so the coins on one side can depend on balloons from the other side."},
			{Prompt: "Why must the table be filled by gap length rather than row by row?", Answer: "best[left][right] reads best[left][k] (same row, smaller right: fine in row order) and best[k][right] with k > left, which belongs to a LATER row in row order and would not be filled yet. Both reads have a smaller gap than (left, right), so increasing gap is a valid order. Running left from n - 1 down to 0 with right increasing also works."},
			{Prompt: "What does a balloon with value 0 do to the best order?", Answer: "Bursting it earns 0 whatever its neighbours are, and while it stands it multiplies any neighbour's burst by 0. So it should usually go early, which is what the DP finds for [5, 0, 5]: burst the 0 first (0 coins), then 5 * 5 * 1 = 25, then 1 * 5 * 1 = 5, total 30. You don't need a special rule; the recurrence compares every last choice."},
		},
	},
	{
		ID:              "regular-expression-matching",
		Number:          10,
		Order:           11,
		Title:           "Regular Expression Matching",
		Slug:            "regular-expression-matching",
		Difficulty:      "Hard",
		Category:        "2d-dp",
		Tags:            []string{"String", "Dynamic Programming", "Recursion"},
		RelatedChapters: []int{16, 15, 8},
		Description: `Implement a tiny pattern matcher. Given a text s and a pattern p, return true if the pattern matches the **whole** of s (not just a part of it), and false otherwise.

The pattern contains lowercase letters and two special characters:

- '.' matches any single letter.
- '*' means "zero or more copies of the character just before it". So "a*" matches "", "a", "aa", ...; ".*" matches any string, including the empty one.

A letter in the pattern matches only that same letter. Every '*' in the pattern has a letter or '.' right before it (a pattern never starts with '*' and never has "**").`,
		Constraints: []string{
			"1 <= s.length <= 20",
			"1 <= p.length <= 20",
			"s contains only lowercase English letters",
			"p contains only lowercase English letters, '.' and '*'",
			"Every '*' is preceded by a letter or '.'",
		},
		Examples: []Example{
			{Input: `s = "aa", p = "a"`, Output: "false", Explanation: "\"a\" matches only the first letter; the whole string must match."},
			{Input: `s = "aa", p = "a*"`, Output: "true", Explanation: "\"a*\" can be two a's."},
			{Input: `s = "ab", p = ".*"`, Output: "true", Explanation: "\".*\" is any number of any letters."},
			{Input: `s = "aab", p = "c*a*b"`, Output: "true", Explanation: "\"c*\" is used zero times, \"a*\" twice, then b matches b."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "aa", "p": "a"}, Expected: false, Note: "example: pattern too short"},
			{Input: map[string]interface{}{"s": "aa", "p": "a*"}, Expected: true, Note: "example: star repeats 'a'"},
			{Input: map[string]interface{}{"s": "ab", "p": ".*"}, Expected: true, Note: "example: '.*' matches anything"},
			{Input: map[string]interface{}{"s": "aab", "p": "c*a*b"}, Expected: true, Hidden: true, Note: "'c*' used zero times"},
			{Input: map[string]interface{}{"s": "mississippi", "p": "mis*is*p*."}, Expected: false, Hidden: true, Note: "near miss: must match the whole string"},
			{Input: map[string]interface{}{"s": "mississippi", "p": "mis*is*ip*."}, Expected: true, Hidden: true, Note: "the same text with a pattern that fits"},
			{Input: map[string]interface{}{"s": "ab", "p": ".*c"}, Expected: false, Hidden: true, Note: "'.*' can't make the final 'c' appear"},
			{Input: map[string]interface{}{"s": "aaa", "p": "a*a"}, Expected: true, Hidden: true, Note: "star must leave one 'a' for the end"},
			{Input: map[string]interface{}{"s": "aaa", "p": "ab*a*c*a"}, Expected: true, Hidden: true, Note: "several stars used zero times"},
			{Input: map[string]interface{}{"s": "a", "p": "ab*"}, Expected: true, Hidden: true, Note: "trailing 'b*' matches nothing"},
			{Input: map[string]interface{}{"s": "ab", "p": "a"}, Expected: false, Hidden: true, Note: "pattern matches only a prefix: false"},
			{Input: map[string]interface{}{"s": "bbbba", "p": ".*a*a"}, Expected: true, Hidden: true, Note: "'.*' must give back the last 'a'"},
			{Input: map[string]interface{}{"s": "aaaaaaaaaaaaaaaaaaaa", "p": "a*a*a*a*a*a*a*a*a*b"}, Expected: false, Hidden: true, Note: "many stars, no match: plain recursion tries millions of splits"},
			{Input: map[string]interface{}{"s": "aaaaaaaaaaaaaaaaaaab", "p": "a*a*a*a*a*a*a*a*a*a*"}, Expected: false, Hidden: true, Note: "ten stars and a letter none of them can match"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def isMatch(s, p):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Treat 'x*' as one unit. When the end of the pattern is 'x*', what are the only two things it can do with the end of the text: match nothing, or match what?"},
			{Level: 2, Type: "insight", Content: "match[i][j] = does s[:i] match p[:j]? If p[j - 1] is a letter or '.', it must match s[i - 1] and the rest is match[i - 1][j - 1], like a diagonal step in LCS. If p[j - 1] is '*' (with x = p[j - 2] before it), either use 'x*' zero times (match[i][j - 2]) or, if x matches s[i - 1], let it eat that letter and stay on the same pattern position (match[i - 1][j]). Chapter 16's two-string table with one special column rule.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. Define the state: match[i][j] = True if the text prefix s[:i] matches the pattern prefix p[:j] completely.\n2. Set the base cases: match[0][0] = True; row 0 (empty text) is True only where the pattern so far is made of 'x*' pairs, which the '*' rule produces by itself if the loop includes i = 0.\n3. Write the recurrence: '*' -> zero copies (two columns left) or one more copy (one row up, same column, if x matches); otherwise -> letters match and the diagonal is True.\n4. Choose the evaluation order: i from 0 to m, j from 1 to n.\n5. Locate the answer: match[m][n]."},
			{Level: 4, Type: "skeleton", Content: "def isMatch(s, p):\n    m, n = len(s), len(p)\n    match = [[False] * (n + 1) for _ in range(m + 1)]\n    match[0][0] = True\n    for i in range(m + 1):\n        for j in range(1, n + 1):\n            if p[j - 1] == '*':\n                zero = ____\n                more = i > 0 and p[j - 2] in (s[i - 1], '.') and ____\n                match[i][j] = zero or more\n            else:\n                same = i > 0 and ____\n                match[i][j] = same and ____\n    return match[m][n]"},
		},
		Solution: Solution{
			Code: `def isMatch(s, p):
    m, n = len(s), len(p)
    # match[i][j] = does s[:i] match the pattern p[:j] completely?
    match = [[False] * (n + 1) for _ in range(m + 1)]
    match[0][0] = True                    # empty matches empty
    for i in range(m + 1):                # row 0 lets "a*b*" match ""
        for j in range(1, n + 1):
            if p[j - 1] == '*':
                zero = match[i][j - 2]    # use "x*" zero times
                more = i > 0 and p[j - 2] in (s[i - 1], '.') and match[i - 1][j]
                match[i][j] = zero or more
            else:
                same = i > 0 and p[j - 1] in (s[i - 1], '.')
                match[i][j] = same and match[i - 1][j - 1]
    return match[m][n]`,
			Explanation: "The last rung combines everything: a two-string table over prefixes (LCS, Edit Distance), a boolean cell (Interleaving String), and a cell rule that depends on what kind of pattern character comes last. A plain character is a diagonal step. A star is the new idea: 'zero copies' jumps two columns left, and 'one more copy' stays in the same column and moves up a row, which is the same 'same row means reuse' trick as Coin Change II's unbounded coins.",
			BruteForce:  "Plain recursion: at each 'x*', try matching zero letters, one letter, two letters, ... and recurse on the rest. Patterns with several stars make this exponential: for s = 20 a's and p = 'a*a*a*a*a*a*a*a*a*b' the recursion tries every way of sharing the 20 a's among 9 stars (C(28, 8), over 3 million splits) before failing on the b. The same (text position, pattern position) pair is checked over and over.",
			KeyInsight:  "Whether the rest matches depends only on how much of the text and how much of the pattern are left: the pair (i, j). There are at most 21 x 21 pairs here, and a star only needs two cases if you let 'one more copy' stay on the same pattern position instead of guessing how many copies to take.",
			WhyItWorks:  "State: match[i][j] = s[:i] matches p[:j] completely. Case 1, p[j - 1] is a letter or '.': the last pattern character must match exactly the last text letter s[i - 1] (letter equal, or '.'), and the prefixes before them must match: match[i - 1][j - 1]. Case 2, p[j - 1] == '*' with x = p[j - 2]: the unit 'x*' matches some number c >= 0 of letters at the end of s[:i]. If c = 0, the rest of the pattern p[:j - 2] must match all of s[:i]: match[i][j - 2]. If c >= 1, the last letter s[i - 1] must match x, and the remaining c - 1 copies are again 'x*' matched against s[:i - 1]: match[i - 1][j]. The two cases cover every c, so 'or' is exact. Base: the empty pattern matches only the empty text (match[0][0] = True, match[i > 0][0] = False). Including i = 0 in the loop fills row 0 by the same rule: 'a*b*' matches '' because each star can take zero copies. Order: every read is to the left in the same row or in the row above.",
			CommonMistakes: []string{
				"Treating '*' as its own wildcard ('any string') instead of as a modifier of the character before it. 'a*' does not match 'b'; '*' only repeats the previous character. Always look at p[j - 2] together with the star.",
				"Forgetting row 0. If the loop starts at i = 1, patterns like 'c*a*b' can never use their stars zero times at the start, and s = 'aab', p = 'c*a*b' returns False instead of True. Start the outer loop at 0 (with the i > 0 guards), or fill row 0 separately.",
				"Accepting a partial match: returning True as soon as the pattern is used up, or checking only that the text is used up. Both prefixes must end together, which is why the answer is match[m][n] and not 'any True cell in the last row or column'. 'mississippi' with 'mis*is*p*.' is False for this reason.",
			},
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "One cell per pair (text prefix, pattern prefix), with row 0 for the empty text and column 0 for the empty pattern.", CodeSnippet: "match = [[False] * (n + 1) for _ in range(m + 1)]", LineStart: 3, LineEnd: 4},
				{Title: "Set the base cases", Explanation: "The empty pattern matches the empty text. Column 0 stays False below that: an empty pattern can't match letters.", CodeSnippet: "match[0][0] = True", LineStart: 5, LineEnd: 5},
				{Title: "Choose the evaluation order", Explanation: "The outer loop starts at i = 0 so that row 0 (empty text) is filled by the star rule: 'x*' pairs can match nothing. Every read is left in the same row or in the row above.", CodeSnippet: "for i in range(m + 1):\n    for j in range(1, n + 1):", LineStart: 6, LineEnd: 7},
				{Title: "Write the recurrence for '*'", Explanation: "Zero copies: drop 'x*' from the pattern (two columns left). One more copy: x must match the last text letter, and then 'x*' may still match more, so stay in column j and go up one row.", CodeSnippet: "zero = match[i][j - 2]\nmore = i > 0 and p[j - 2] in (s[i - 1], '.') and match[i - 1][j]\nmatch[i][j] = zero or more", LineStart: 8, LineEnd: 11},
				{Title: "Write the recurrence for a letter or '.'", Explanation: "A single pattern character consumes exactly one text letter: the letters must agree (or the pattern has '.'), and the shorter prefixes must match (diagonal).", CodeSnippet: "same = i > 0 and p[j - 1] in (s[i - 1], '.')\nmatch[i][j] = same and match[i - 1][j - 1]", LineStart: 12, LineEnd: 14},
				{Title: "Locate the answer", Explanation: "The whole text against the whole pattern. A True somewhere else in the table only means a prefix matched.", CodeSnippet: "return match[m][n]", LineStart: 15, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{
					Name: "Top-down: the brute-force recursion plus a cache",
					Code: `from functools import lru_cache

def isMatch(s, p):
    @lru_cache(maxsize=None)    # without the cache "a*a*a*..." patterns explode
    def fits(i, j):             # does s[i:] match p[j:]?
        if j == len(p):
            return i == len(s)
        first = i < len(s) and p[j] in (s[i], '.')
        if j + 1 < len(p) and p[j + 1] == '*':
            return fits(i, j + 2) or (first and fits(i + 1, j))
        return first and fits(i + 1, j + 1)

    return fits(0, 0)`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "fits(i, j) reads the pattern from the front: if the next unit is 'x*', skip it (j + 2) or, if x matches s[i], consume one letter and stay on 'x*' (i + 1, j). Many people find this order easier to derive. Without the cache the many-stars test is exponential; with it there are (m + 1) x (n + 1) states.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does 'one more copy' read match[i - 1][j] and not match[i - 1][j - 2]?", Answer: "After 'x*' eats the letter s[i - 1], it may still eat more copies before it, so the pattern stays at the same position j (the star is still available). Reading match[i - 1][j - 2] would mean 'exactly one copy', and you would need a separate case for two, three, ... copies. Staying in column j covers every count with one rule, the same way Coin Change II's same-row read allowed any number of copies of a coin."},
			{Prompt: "Why is 'mississippi' / 'mis*is*p*.' false but 'mississippi' / 'mis*is*ip*.' true?", Answer: "In the first pattern, after 'mis*' takes 'miss' and 'is*' takes 'iss', the remaining text is 'ippi', but the pattern has only 'p*.' left, which can match some p's and then one letter: it can't produce the leading 'i'. In the second pattern 'ip*.' matches 'i', 'pp' and the final 'i'. The whole text must be consumed, and the table's last cell captures exactly that."},
			{Prompt: "What would change if '+' (one or more copies) were added?", Answer: "'x+' is 'x' followed by 'x*', so you could rewrite the pattern first. In the table: 'x+' at the end of p[:j] requires at least one copy, so match[i][j] = x matches s[i - 1] and (match[i - 1][j] or match[i - 1][j - 2]): one more copy staying on 'x+', or this was the first copy and the rest is p[:j - 2]. There is no zero-copy case."},
		},
	},
}

// dp2dSeq returns n pseudo-random values in 0..mod-1 from a fixed linear
// congruential generator, so large test inputs stay short in the source.
func dp2dSeq(n, mod, seed int) []int {
	out := make([]int, n)
	x := seed
	for i := range out {
		x = (x*1103515245 + 12345) % 2147483648
		out[i] = (x >> 16) % mod
	}
	return out
}

// dp2dLetters returns n letters drawn from alphabet with dp2dSeq.
func dp2dLetters(n int, alphabet string, seed int) string {
	b := make([]byte, n)
	for i, v := range dp2dSeq(n, len(alphabet), seed) {
		b[i] = alphabet[v]
	}
	return string(b)
}

// dp2dDiagonal returns an n x n grid whose cell (r, c) holds r + c, so every
// right/down route from the top-left corner is increasing.
func dp2dDiagonal(n int) [][]int {
	grid := make([][]int, n)
	for r := range grid {
		grid[r] = make([]int, n)
		for c := range grid[r] {
			grid[r][c] = r + c
		}
	}
	return grid
}

// dp2dSnake returns an n x n grid numbered 0..n*n-1 along a back-and-forth
// path, so the longest increasing path visits every cell.
func dp2dSnake(n int) [][]int {
	grid := make([][]int, n)
	v := 0
	for r := range grid {
		grid[r] = make([]int, n)
		for k := 0; k < n; k++ {
			c := k
			if r%2 == 1 {
				c = n - 1 - k
			}
			grid[r][c] = v
			v++
		}
	}
	return grid
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, DP2DProblems...)
}
