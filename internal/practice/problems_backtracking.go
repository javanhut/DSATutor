package practice

// BacktrackingProblems contains all backtracking category problems
var BacktrackingProblems = []*Problem{
	{
		ID:              "subsets",
		Number:          76,
		Title:           "Subsets",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking", "Bit Manipulation"},
		RelatedChapters: []int{4, 9, 12},
		Description: `Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.`,
		Constraints: []string{
			"1 <= nums.length <= 10",
			"-10 <= nums[i] <= 10",
			"All the numbers of nums are unique",
		},
		Examples: []Example{
			{Input: "nums = [1,2,3]", Output: "[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]"},
			{Input: "nums = [0]", Output: "[[],[0]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}}, Expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		},
		TimeComplexity:  "O(n * 2^n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def subsets(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Backtracking: for each element, choose to include it or not. Or iteratively: for each new element, add it to all existing subsets.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "Backtrack with index. At each position, choose to include nums[i] or skip it. When index reaches end, add current subset to result."},
			{Level: 3, Type: "code", Content: "def backtrack(start, path): result.append(path[:]). for i in range(start, n): path.append(nums[i]); backtrack(i+1, path); path.pop()."},
		},
		Solution: Solution{
			Code: `def subsets(nums):
    result = []

    def backtrack(start, path):
        result.append(path[:])  # Add copy of current subset

        for i in range(start, len(nums)):
            path.append(nums[i])
            backtrack(i + 1, path)
            path.pop()

    backtrack(0, [])
    return result`,
			Explanation:     "Backtracking builds subsets by including or excluding each element. Start from each index to avoid duplicates.",
			TimeComplexity:  "O(n * 2^n)",
			SpaceComplexity: "O(n) for recursion",
			Walkthrough: []WalkthroughStep{
				{Title: "Add current subset", Explanation: "Every path is a valid subset", CodeSnippet: "result.append(path[:])", LineStart: 5, LineEnd: 5},
				{Title: "Try adding each element", Explanation: "Start from current index to avoid duplicates", CodeSnippet: "for i in range(start, len(nums)):", LineStart: 7, LineEnd: 7},
				{Title: "Backtrack", Explanation: "Remove element to try other combinations", CodeSnippet: "path.pop()", LineStart: 10, LineEnd: 10},
			},
		},
	},
	{
		ID:              "word-search",
		Number:          77,
		Title:           "Word Search",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking", "Matrix"},
		RelatedChapters: []int{4, 6, 12},
		Description: `Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.`,
		Constraints: []string{
			"m == board.length",
			"n == board[i].length",
			"1 <= m, n <= 6",
			"1 <= word.length <= 15",
			"board and word consists of only lowercase and uppercase English letters",
		},
		Examples: []Example{
			{Input: `board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"`, Output: "true"},
			{Input: `board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"`, Output: "true"},
			{Input: `board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"`, Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"board": [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}, "word": "ABCCED"}, Expected: true},
			{Input: map[string]interface{}{"board": [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}, "word": "SEE"}, Expected: true},
		},
		TimeComplexity:  "O(m * n * 4^L)",
		SpaceComplexity: "O(L)",
		StarterCode:     "def exist(board, word):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "DFS from each cell. Mark cells as visited during recursion. Backtrack by unmarking.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "For each cell matching word[0], DFS to find rest of word. Mark visited by modifying cell temporarily."},
			{Level: 3, Type: "code", Content: "def dfs(r, c, i): if i == len(word): return True. if out of bounds or board[r][c] != word[i]: return False. Mark, try 4 dirs, unmark."},
		},
		Solution: Solution{
			Code: `def exist(board, word):
    m, n = len(board), len(board[0])

    def dfs(r, c, i):
        if i == len(word):
            return True
        if r < 0 or r >= m or c < 0 or c >= n or board[r][c] != word[i]:
            return False

        # Mark as visited
        temp = board[r][c]
        board[r][c] = '#'

        # Try all 4 directions
        found = (dfs(r + 1, c, i + 1) or
                 dfs(r - 1, c, i + 1) or
                 dfs(r, c + 1, i + 1) or
                 dfs(r, c - 1, i + 1))

        # Restore
        board[r][c] = temp
        return found

    for r in range(m):
        for c in range(n):
            if dfs(r, c, 0):
                return True
    return False`,
			Explanation:     "DFS from each starting cell. Mark visited cells temporarily. Backtrack by restoring. Try all 4 directions.",
			TimeComplexity:  "O(m * n * 4^L) where L is word length",
			SpaceComplexity: "O(L) for recursion stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Check bounds and match", Explanation: "Early return if invalid or no match", CodeSnippet: "if r < 0 or r >= m or c < 0 or c >= n or board[r][c] != word[i]:", LineStart: 7, LineEnd: 8},
				{Title: "Mark visited", Explanation: "Prevent revisiting same cell", CodeSnippet: "board[r][c] = '#'", LineStart: 12, LineEnd: 12},
				{Title: "Restore after backtrack", Explanation: "Allow cell to be used in other paths", CodeSnippet: "board[r][c] = temp", LineStart: 21, LineEnd: 21},
			},
		},
	},
	{
		ID:              "n-queens",
		Number:          78,
		Title:           "N-Queens",
		Difficulty:      "Hard",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking"},
		RelatedChapters: []int{4, 12},
		Description: `The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.`,
		Constraints: []string{
			"1 <= n <= 9",
		},
		Examples: []Example{
			{Input: "n = 4", Output: `[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]`},
			{Input: "n = 1", Output: `[["Q"]]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 4}, Expected: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
		},
		TimeComplexity:  "O(n!)",
		SpaceComplexity: "O(n^2)",
		StarterCode:     "def solveNQueens(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Backtracking row by row. Track which columns and diagonals are under attack.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "Place queen in each valid column of current row. Use sets for columns, main diagonals (r-c), anti-diagonals (r+c)."},
			{Level: 3, Type: "code", Content: "def backtrack(row): if row == n: add solution. for col: if col, r-c, r+c not attacked: place queen, recurse, remove."},
		},
		Solution: Solution{
			Code: `def solveNQueens(n):
    result = []
    board = [['.' for _ in range(n)] for _ in range(n)]
    cols = set()
    diag1 = set()  # r - c
    diag2 = set()  # r + c

    def backtrack(row):
        if row == n:
            result.append([''.join(r) for r in board])
            return

        for col in range(n):
            if col in cols or (row - col) in diag1 or (row + col) in diag2:
                continue

            # Place queen
            board[row][col] = 'Q'
            cols.add(col)
            diag1.add(row - col)
            diag2.add(row + col)

            backtrack(row + 1)

            # Remove queen
            board[row][col] = '.'
            cols.remove(col)
            diag1.remove(row - col)
            diag2.remove(row + col)

    backtrack(0)
    return result`,
			Explanation:     "Place queens row by row. Track attacked columns and diagonals with sets. Backtrack when no valid position.",
			TimeComplexity:  "O(n!)",
			SpaceComplexity: "O(n^2) for board",
			Walkthrough: []WalkthroughStep{
				{Title: "Track attacks", Explanation: "Sets for columns and both diagonals", CodeSnippet: "cols = set()\ndiag1 = set()  # r - c\ndiag2 = set()  # r + c", LineStart: 4, LineEnd: 6},
				{Title: "Check if safe", Explanation: "Column and diagonals not under attack", CodeSnippet: "if col in cols or (row - col) in diag1 or (row + col) in diag2:", LineStart: 14, LineEnd: 15},
				{Title: "Backtrack", Explanation: "Remove queen and attack markers", CodeSnippet: "board[row][col] = '.'\ncols.remove(col)", LineStart: 26, LineEnd: 29},
			},
		},
	},
	{
		ID:              "subsets-ii",
		Number:          113,
		Title:           "Subsets II",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking", "Bit Manipulation"},
		RelatedChapters: []int{4, 12},
		Description: `Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.`,
		Constraints: []string{
			"1 <= nums.length <= 10",
			"-10 <= nums[i] <= 10",
		},
		Examples: []Example{
			{Input: "nums = [1,2,2]", Output: "[[],[1],[1,2],[1,2,2],[2],[2,2]]"},
			{Input: "nums = [0]", Output: "[[],[0]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 2}}, Expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		},
		TimeComplexity:  "O(n * 2^n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def subsetsWithDup(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Sort array to group duplicates. Skip consecutive duplicates at same recursion level.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "Backtrack like Subsets I, but skip nums[i] if nums[i] == nums[i-1] and we didn't include nums[i-1]."},
			{Level: 3, Type: "code", Content: "Sort nums. In backtrack loop: if i > start and nums[i] == nums[i-1]: continue. This skips duplicates at same level."},
		},
		Solution: Solution{
			Code: `def subsetsWithDup(nums):
    nums.sort()
    result = []

    def backtrack(start, path):
        result.append(path[:])

        for i in range(start, len(nums)):
            # Skip duplicates at same level
            if i > start and nums[i] == nums[i - 1]:
                continue
            path.append(nums[i])
            backtrack(i + 1, path)
            path.pop()

    backtrack(0, [])
    return result`,
			Explanation:     "Sort to group duplicates. Skip duplicate at same recursion level to avoid duplicate subsets.",
			TimeComplexity:  "O(n * 2^n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort first", Explanation: "Groups duplicates together", CodeSnippet: "nums.sort()", LineStart: 2, LineEnd: 2},
				{Title: "Skip duplicates", Explanation: "At same level, skip if same as previous", CodeSnippet: "if i > start and nums[i] == nums[i - 1]:\n    continue", LineStart: 10, LineEnd: 11},
			},
		},
	},
	{
		ID:              "combination-sum-ii",
		Number:          114,
		Title:           "Combination Sum II",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Array", "Backtracking"},
		RelatedChapters: []int{4, 12},
		Description: `Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.`,
		Constraints: []string{
			"1 <= candidates.length <= 100",
			"1 <= candidates[i] <= 50",
			"1 <= target <= 30",
		},
		Examples: []Example{
			{Input: "candidates = [10,1,2,7,6,1,5], target = 8", Output: "[[1,1,6],[1,2,5],[1,7],[2,6]]"},
			{Input: "candidates = [2,5,2,1,2], target = 5", Output: "[[1,2,2],[5]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"candidates": []int{10, 1, 2, 7, 6, 1, 5}, "target": 8}, Expected: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		},
		TimeComplexity:  "O(2^n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def combinationSum2(candidates, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Sort and skip duplicates like Subsets II. Each element used at most once.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "Backtrack with remaining target. Skip duplicates at same level. Stop if target < 0."},
			{Level: 3, Type: "code", Content: "Sort. backtrack(start, target, path). if target == 0: add. for i from start: skip dups, recurse with i+1, target-num."},
		},
		Solution: Solution{
			Code: `def combinationSum2(candidates, target):
    candidates.sort()
    result = []

    def backtrack(start, target, path):
        if target == 0:
            result.append(path[:])
            return
        if target < 0:
            return

        for i in range(start, len(candidates)):
            if i > start and candidates[i] == candidates[i - 1]:
                continue
            if candidates[i] > target:
                break
            path.append(candidates[i])
            backtrack(i + 1, target - candidates[i], path)
            path.pop()

    backtrack(0, target, [])
    return result`,
			Explanation:     "Sort and skip duplicates. Each element used once (i+1 in recursion). Early termination when sorted value exceeds target.",
			TimeComplexity:  "O(2^n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Skip duplicates", Explanation: "Avoid duplicate combinations", CodeSnippet: "if i > start and candidates[i] == candidates[i - 1]:\n    continue", LineStart: 13, LineEnd: 14},
				{Title: "Early termination", Explanation: "Sorted array allows pruning", CodeSnippet: "if candidates[i] > target:\n    break", LineStart: 15, LineEnd: 16},
				{Title: "Use once", Explanation: "i+1 ensures element used at most once", CodeSnippet: "backtrack(i + 1, target - candidates[i], path)", LineStart: 18, LineEnd: 18},
			},
		},
	},
	{
		ID:              "letter-combinations-of-a-phone-number",
		Number:          115,
		Title:           "Letter Combinations of a Phone Number",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Hash Table", "String", "Backtracking"},
		RelatedChapters: []int{4, 5},
		Description: `Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.`,
		Constraints: []string{
			"0 <= digits.length <= 4",
			"digits[i] is a digit in the range ['2', '9']",
		},
		Examples: []Example{
			{Input: `digits = "23"`, Output: `["ad","ae","af","bd","be","bf","cd","ce","cf"]`},
			{Input: `digits = ""`, Output: `[]`},
			{Input: `digits = "2"`, Output: `["a","b","c"]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"digits": "23"}, Expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
			{Input: map[string]interface{}{"digits": ""}, Expected: []string{}},
		},
		TimeComplexity:  "O(4^n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def letterCombinations(digits):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Map each digit to its letters. Backtrack through each digit, trying all its letters.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "For each position, try all letters mapped to that digit. When we've processed all digits, add combination."},
			{Level: 3, Type: "code", Content: "phone = {'2': 'abc', ...}. backtrack(index, path). for letter in phone[digits[index]]: recurse with index + 1."},
		},
		Solution: Solution{
			Code: `def letterCombinations(digits):
    if not digits:
        return []

    phone = {
        '2': 'abc', '3': 'def', '4': 'ghi', '5': 'jkl',
        '6': 'mno', '7': 'pqrs', '8': 'tuv', '9': 'wxyz'
    }
    result = []

    def backtrack(index, path):
        if index == len(digits):
            result.append(''.join(path))
            return

        for letter in phone[digits[index]]:
            path.append(letter)
            backtrack(index + 1, path)
            path.pop()

    backtrack(0, [])
    return result`,
			Explanation:     "Map digits to letters. Backtrack through each digit position, trying all mapped letters.",
			TimeComplexity:  "O(4^n) - some digits have 4 letters",
			SpaceComplexity: "O(n) for recursion",
			Walkthrough: []WalkthroughStep{
				{Title: "Phone mapping", Explanation: "Digit to letters mapping", CodeSnippet: "phone = {'2': 'abc', '3': 'def', ...}", LineStart: 5, LineEnd: 8},
				{Title: "Try all letters", Explanation: "For current digit, try each mapped letter", CodeSnippet: "for letter in phone[digits[index]]:", LineStart: 16, LineEnd: 16},
				{Title: "Build combination", Explanation: "When all digits processed, join path", CodeSnippet: "result.append(''.join(path))", LineStart: 13, LineEnd: 13},
			},
		},
	},
	{
		ID:              "palindrome-partitioning",
		Number:          137,
		Title:           "Palindrome Partitioning",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"String", "Backtracking", "Dynamic Programming"},
		RelatedChapters: []int{4, 9, 12},
		Description: `Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.`,
		Constraints: []string{
			"1 <= s.length <= 16",
			"s contains only lowercase English letters",
		},
		Examples: []Example{
			{Input: `s = "aab"`, Output: `[["a","a","b"],["aa","b"]]`},
			{Input: `s = "a"`, Output: `[["a"]]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "aab"}, Expected: [][]string{{"a", "a", "b"}, {"aa", "b"}}},
			{Input: map[string]interface{}{"s": "a"}, Expected: [][]string{{"a"}}},
		},
		TimeComplexity:  "O(n * 2^n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def partition(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Backtracking: try all possible first palindrome prefixes, then recurse on rest.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "For each position, try all palindrome substrings starting there. Backtrack when reaching end."},
			{Level: 3, Type: "code", Content: "def backtrack(start, path): if start == len(s): result.append(path[:]). for end in range(start, n): if isPalindrome: path.append; backtrack; path.pop."},
		},
		Solution: Solution{
			Code: `def partition(s):
    def is_palindrome(start, end):
        while start < end:
            if s[start] != s[end]:
                return False
            start += 1
            end -= 1
        return True

    def backtrack(start, path):
        if start == len(s):
            result.append(path[:])
            return

        for end in range(start, len(s)):
            if is_palindrome(start, end):
                path.append(s[start:end + 1])
                backtrack(end + 1, path)
                path.pop()

    result = []
    backtrack(0, [])
    return result`,
			Explanation:     "Try all palindrome prefixes at each position, backtrack to explore all partitions.",
			TimeComplexity:  "O(n * 2^n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Base case", Explanation: "Reached end, found valid partition", CodeSnippet: "if start == len(s):\n    result.append(path[:])", LineStart: 11, LineEnd: 12},
				{Title: "Try prefixes", Explanation: "All palindrome substrings from start", CodeSnippet: "for end in range(start, len(s)):\n    if is_palindrome(start, end):", LineStart: 15, LineEnd: 16},
				{Title: "Backtrack", Explanation: "Remove and try next option", CodeSnippet: "path.pop()", LineStart: 19, LineEnd: 19},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, BacktrackingProblems...)
}
