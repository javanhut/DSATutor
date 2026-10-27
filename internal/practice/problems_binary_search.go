package practice

// BinarySearchProblems contains all binary-search category problems.
//
// Ladder (chapter 6): exact-value template on a plain array, the same template
// on a matrix through index arithmetic, the boundary ("first true") template
// on an answer range, the boundary template on a rotated array, a rotated
// array with an exact target, a boundary search inside a design problem, and
// finally a binary search over partitions of two arrays.
var BinarySearchProblems = []*Problem{
	{
		ID:              "binary-search",
		Number:          704,
		Order:           1,
		Title:           "Binary Search",
		Slug:            "binary-search",
		Difficulty:      "Easy",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search"},
		RelatedChapters: []int{6, 1},
		Description: `You are given an array nums of distinct integers sorted in ascending order, and an integer target.

Return the index of target in nums, or -1 if target does not appear.

Your algorithm must run in O(log n) time, so checking every element one by one is not allowed.`,
		Constraints: []string{
			"1 <= nums.length <= 10^4",
			"-10^4 < nums[i], target < 10^4",
			"All the integers in nums are unique",
			"nums is sorted in ascending order",
		},
		Examples: []Example{
			{Input: "nums = [-1,0,3,5,9,12], target = 9", Output: "4", Explanation: "nums[4] is 9. Binary search looks at index 2 (value 3, too small), then index 4 (value 9, found): two looks instead of five."},
			{Input: "nums = [-1,0,3,5,9,12], target = 2", Output: "-1", Explanation: "2 would sit between 0 and 3, but it isn't there. The search range becomes empty, so the answer is -1."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{-1, 0, 3, 5, 9, 12}, "target": 9}, Expected: 4, Note: "target in the right half"},
			{Input: map[string]interface{}{"nums": []int{-1, 0, 3, 5, 9, 12}, "target": 2}, Expected: -1, Note: "target missing, would fall between two elements"},
			{Input: map[string]interface{}{"nums": []int{5}, "target": 5}, Expected: 0, Note: "single element that matches (needs lo <= hi, not lo < hi)"},
			{Input: map[string]interface{}{"nums": []int{5}, "target": -5}, Expected: -1, Note: "single element that doesn't match"},
			{Input: map[string]interface{}{"nums": []int{-1, 0, 3, 5, 9, 12}, "target": -1}, Expected: 0, Note: "target is the first element"},
			{Input: map[string]interface{}{"nums": []int{-1, 0, 3, 5, 9, 12}, "target": 12}, Expected: 5, Hidden: true, Note: "target is the last element (catches hi = len(nums) - 2 style off-by-one)"},
			{Input: map[string]interface{}{"nums": []int{-1, 0, 3, 5, 9, 12}, "target": -7}, Expected: -1, Hidden: true, Note: "target smaller than everything"},
			{Input: map[string]interface{}{"nums": []int{-1, 0, 3, 5, 9, 12}, "target": 20}, Expected: -1, Hidden: true, Note: "target larger than everything (hi must not run past the end)"},
			{Input: map[string]interface{}{"nums": []int{2, 5}, "target": 5}, Expected: 1, Hidden: true, Note: "two elements: lo = mid would loop forever here, lo = mid + 1 is required"},
			{Input: map[string]interface{}{"nums": []int{1, 4, 6, 8, 10}, "target": 6}, Expected: 2, Hidden: true, Note: "target found on the very first middle"},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def search(nums, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Look at the middle element and compare it with target. If the middle is smaller than target, which elements can you rule out without ever looking at them?"},
			{Level: 2, Type: "insight", Content: "Because nums is sorted, one comparison with the middle element rules out a whole half. Keep a range [lo, hi] that still could hold target and halve it every step. This is the exact-value template of binary search: loop while lo <= hi and move lo to mid + 1 or hi to mid - 1, because mid itself has been checked.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Define the search space: lo = 0, hi = last index (both ends included).\n2. Pick the middle: while the range is not empty (lo <= hi), mid = (lo + hi) // 2.\n3. Write the yes/no test: if nums[mid] equals target, you're done.\n4. Discard half: if nums[mid] is too small, everything up to mid is too small, so lo = mid + 1; otherwise hi = mid - 1.\n5. Read the answer: if the range empties, target isn't there; return -1."},
			{Level: 4, Type: "skeleton", Content: `def search(nums, target):
    lo, hi = 0, len(nums) - 1
    while ____:
        mid = (lo + hi) // 2
        if nums[mid] == target:
            return mid
        if nums[mid] < target:
            lo = ____
        else:
            hi = ____
    return -1`},
		},
		Solution: Solution{
			Code: `def search(nums, target):
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] == target:
            return mid
        if nums[mid] < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return -1`,
			Explanation: "This is the plainest form of binary search and the first rung of the chapter 6 ladder. Keep a closed range [lo, hi] of indices where target could still be, look at its middle, and throw away the half that can't contain target. Each step halves the range, so an array of 10,000 elements needs at most 14 looks.",
			BruteForce:  "Scan from left to right and return the index where nums[i] == target. That is O(n): for n = 10,000 it may do 10,000 comparisons. It wastes the sorted order: once you know nums[2] = 3 is smaller than 9, indices 0, 1 and 2 are all ruled out at once, yet a scan still visits each of them.",
			KeyInsight:  "In a sorted array, comparing target with one element tells you about all elements on one side of it. If nums[mid] < target, then nums[0..mid] are all smaller too, so a single comparison discards about half of the remaining candidates.",
			WhyItWorks:  "Invariant: if target is in nums, its index lies in [lo, hi]. It holds at the start because [0, n-1] is the whole array. When nums[mid] < target, every index <= mid holds a value <= nums[mid] < target, so setting lo = mid + 1 keeps the invariant; the case nums[mid] > target is symmetric with hi = mid - 1. The range shrinks by at least one every iteration (mid is excluded), so the loop ends. If it ends because lo > hi, the range is empty, and by the invariant target is absent, so -1 is correct.",
			CommonMistakes: []string{
				"Writing while lo < hi with this template: a one-element range (lo == hi) is never checked, so [5] with target 5 returns -1. With a closed range [lo, hi], loop while lo <= hi.",
				"Writing lo = mid or hi = mid instead of mid + 1 / mid - 1: when hi = lo + 1, mid equals lo and the range never shrinks, so the loop runs forever. mid has already been compared, so exclude it.",
				"Starting with hi = len(nums) and then reading nums[hi] or nums[mid] when mid can reach len(nums): an IndexError for a target larger than everything. With the closed template, hi starts at len(nums) - 1.",
			},
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the search space", Explanation: "lo and hi are the first and last indices that could still hold target. Both ends are included, which is why this template is called a closed range.", CodeSnippet: "lo, hi = 0, len(nums) - 1", LineStart: 2, LineEnd: 2},
				{Title: "Pick the middle", Explanation: "While the range [lo, hi] has at least one element, take the middle index. Integer division rounds down, so mid is always inside the range.", CodeSnippet: "while lo <= hi:\nmid = (lo + hi) // 2", LineStart: 3, LineEnd: 4},
				{Title: "Write the yes/no test", Explanation: "The exact-value question: is this the element? If so, return its index straight away.", CodeSnippet: "if nums[mid] == target:\nreturn mid", LineStart: 5, LineEnd: 6},
				{Title: "Discard half", Explanation: "If nums[mid] is too small, mid and everything left of it are too small, so the new range starts at mid + 1. Otherwise mid and everything right of it are too big, so it ends at mid - 1. Excluding mid is what guarantees progress.", CodeSnippet: "lo = mid + 1\nhi = mid - 1", LineStart: 7, LineEnd: 10},
				{Title: "Read the answer", Explanation: "The loop only exits without returning when lo > hi, meaning no candidate index is left, so target is not in the array.", CodeSnippet: "return -1", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Linear scan (brute force)",
					Code: `def search(nums, target):
    for i, value in enumerate(nums):
        if value == target:
            return i
    return -1`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Correct but ignores the sorted order, so it may look at every element. Useful as a reference to test the binary search against.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the loop condition have to be lo <= hi rather than lo < hi in this template?", Answer: "lo and hi are both candidate indices, so when lo == hi there is still exactly one element left that might be target. With lo < hi the loop would stop without checking it, and nums = [5], target = 5 would wrongly return -1."},
			{Prompt: "At most how many times does the loop run for n = 10,000, and why?", Answer: "About 14. Each iteration removes mid and at least half of the rest, so the range size goes 10,000 -> 5,000 -> 2,500 -> ... -> 1, which takes floor(log2 10,000) + 1 = 14 steps. That is the O(log n) bound from chapter 1."},
			{Prompt: "What if the array could contain duplicates and you had to return the FIRST index of target?", Answer: "Stopping at the first match is no longer enough, because an equal element may sit further left. Use the boundary template instead: find the first index where nums[i] >= target with while lo < hi and hi = mid, then check whether nums[lo] == target."},
		},
	},
	{
		ID:              "search-2d-matrix",
		Number:          74,
		Order:           2,
		Title:           "Search a 2D Matrix",
		Slug:            "search-a-2d-matrix",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search", "Matrix"},
		RelatedChapters: []int{6, 2},
		Description: `You are given an m x n integer grid matrix with two properties:

- Each row is sorted in non-decreasing order (left to right).
- The first number of each row is greater than the last number of the row above it.

Given an integer target, return true if target appears anywhere in matrix and false otherwise.

Your solution must run in O(log(m * n)) time.`,
		Constraints: []string{
			"m == matrix.length",
			"n == matrix[i].length",
			"1 <= m, n <= 100",
			"-10^4 <= matrix[i][j], target <= 10^4",
		},
		Examples: []Example{
			{Input: "matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3", Output: "true", Explanation: "Read row by row, the grid is one sorted list of 12 numbers: 1,3,5,7,10,...,60. Position 1 in that list is matrix[0][1] = 3."},
			{Input: "matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13", Output: "false", Explanation: "13 would fall between 11 and 16 but isn't there."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, "target": 3}, Expected: true, Note: "target inside the first row"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, "target": 13}, Expected: false, Note: "target missing, falls inside a row's range"},
			{Input: map[string]interface{}{"matrix": [][]int{{1}}, "target": 1}, Expected: true, Note: "1x1 grid that matches"},
			{Input: map[string]interface{}{"matrix": [][]int{{1}}, "target": 2}, Expected: false, Note: "1x1 grid that doesn't match"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 3, 5}}, "target": 5}, Expected: true, Hidden: true, Note: "single row, last element"},
			{Input: map[string]interface{}{"matrix": [][]int{{1}, {3}}, "target": 3}, Expected: true, Hidden: true, Note: "single column (catches swapping rows and cols in mid // cols, mid % cols)"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, "target": 60}, Expected: true, Hidden: true, Note: "very last cell (catches hi = m * n - 2)"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, "target": 8}, Expected: false, Hidden: true, Note: "target falls in the gap between two rows"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, "target": 0}, Expected: false, Hidden: true, Note: "target smaller than everything"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, "target": 100}, Expected: false, Hidden: true, Note: "target larger than everything"},
		},
		TimeComplexity:  "O(log(m * n))",
		SpaceComplexity: "O(1)",
		StarterCode:     "def searchMatrix(matrix, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you read the grid row after row and wrote the numbers down in one line, what would that line look like?"},
			{Level: 2, Type: "insight", Content: "Row by row, the grid is a single sorted list of m * n numbers, so the exact-value binary search from Binary Search (704) works unchanged. You only need to translate a position k in that imaginary list into a cell: row k // n, column k % n, where n is the number of columns.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Define the search space: positions 0 to m * n - 1 of the flattened list.\n2. Pick the middle: mid = (lo + hi) // 2, then turn it into a cell with row = mid // cols and col = mid % cols.\n3. Write the yes/no test: is that cell equal to target?\n4. Discard half: too small means lo = mid + 1, too big means hi = mid - 1.\n5. Read the answer: an empty range means false."},
			{Level: 4, Type: "skeleton", Content: `def searchMatrix(matrix, target):
    rows, cols = len(matrix), len(matrix[0])
    lo, hi = 0, ____
    while lo <= hi:
        mid = (lo + hi) // 2
        value = matrix[____][____]
        if value == target:
            return True
        if value < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return False`},
		},
		Solution: Solution{
			Code: `def searchMatrix(matrix, target):
    rows, cols = len(matrix), len(matrix[0])
    lo, hi = 0, rows * cols - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        value = matrix[mid // cols][mid % cols]
        if value == target:
            return True
        if value < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return False`,
			Explanation: "Builds on Binary Search (704): the loop is the same exact-value template. The one new idea is that a sorted grid is a sorted list in disguise. Search over positions 0..m*n-1 and convert each middle position into a row and a column with integer division and remainder, so no copy of the grid is ever made.",
			BruteForce:  "Check every cell, or test target in row for each row. That is O(m * n): 10,000 comparisons for a 100 x 100 grid. It wastes both sorted properties: a single comparison with a middle cell could rule out half the grid at once.",
			KeyInsight:  "Because each row starts above where the previous one ended, reading the rows in order gives one sorted sequence. Position k of that sequence is cell (k // cols, k % cols), so you can binary-search the sequence without building it.",
			WhyItWorks:  "Let flat be the rows concatenated. The two grid properties make flat sorted, and flat[k] == matrix[k // cols][k % cols] because each row holds exactly cols numbers: k // cols counts how many full rows come before position k and k % cols is the offset inside that row. The loop is therefore exactly the 704 search on flat, with the same invariant (if target exists, its position is in [lo, hi]) and the same halving, giving O(log(m * n)) steps.",
			CommonMistakes: []string{
				"Dividing by the number of rows instead of columns (mid // rows): it happens to work on square grids and breaks on others, like the single-column [[1],[3]]. Each row holds cols numbers, so divide by cols.",
				"Setting hi = rows * cols instead of rows * cols - 1 with the closed template: a target bigger than everything makes mid reach rows * cols and the row index runs off the grid. The last valid position is rows * cols - 1.",
				"Finding the right row with binary search and then scanning that row linearly: that's O(log m + n), not O(log(m * n)). Search the row with binary search too, or use the flattened index.",
			},
			TimeComplexity:  "O(log(m * n))",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the search space", Explanation: "The candidates are positions 0 through rows * cols - 1 of the flattened grid. For a 3 x 4 grid that's 0..11.", CodeSnippet: "rows, cols = len(matrix), len(matrix[0])\nlo, hi = 0, rows * cols - 1", LineStart: 2, LineEnd: 3},
				{Title: "Pick the middle", Explanation: "Take the middle position and turn it into a cell. With cols = 4, position 5 is row 5 // 4 = 1, column 5 % 4 = 1, which holds 11.", CodeSnippet: "mid = (lo + hi) // 2\nvalue = matrix[mid // cols][mid % cols]", LineStart: 4, LineEnd: 6},
				{Title: "Write the yes/no test", Explanation: "Same exact-value question as in 704: is this cell the target?", CodeSnippet: "if value == target:\nreturn True", LineStart: 7, LineEnd: 8},
				{Title: "Discard half", Explanation: "Too small: every earlier position is smaller too, so skip past mid. Too big: skip everything from mid on.", CodeSnippet: "lo = mid + 1\nhi = mid - 1", LineStart: 9, LineEnd: 12},
				{Title: "Read the answer", Explanation: "The range emptied without a match, so target isn't in the grid.", CodeSnippet: "return False", LineStart: 13, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Scan every row (brute force)",
					Code: `def searchMatrix(matrix, target):
    for row in matrix:
        for value in row:
            if value == target:
                return True
    return False`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Looks at every cell. Fine for tiny grids and a good reference for testing, but it ignores both sorted properties.",
				},
				{
					Name: "Two binary searches: row, then column",
					Code: `def searchMatrix(matrix, target):
    top, bottom = 0, len(matrix) - 1
    while top <= bottom:
        mid_row = (top + bottom) // 2
        row = matrix[mid_row]
        if target < row[0]:
            bottom = mid_row - 1
        elif target > row[-1]:
            top = mid_row + 1
        else:
            lo, hi = 0, len(row) - 1
            while lo <= hi:
                mid = (lo + hi) // 2
                if row[mid] == target:
                    return True
                if row[mid] < target:
                    lo = mid + 1
                else:
                    hi = mid - 1
            return False
    return False`,
					TimeComplexity:  "O(log m + log n)",
					SpaceComplexity: "O(1)",
					Explanation:     "First binary-search the rows for the one whose range [row[0], row[-1]] covers target, then binary-search inside it. log m + log n equals log(m * n), so this is just as fast; it is longer but avoids the index arithmetic.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is position k of the flattened grid at matrix[k // cols][k % cols]?", Answer: "Every row holds exactly cols numbers. Before position k there are k numbers, which fill k // cols complete rows, so k is in row k // cols. Within that row, the k % cols numbers left over come before it, so it sits at column k % cols. For cols = 4, k = 9 is row 2, column 1."},
			{Prompt: "What if rows were sorted but a row could start lower than the previous row ended (LeetCode 240, Search a 2D Matrix II)?", Answer: "The flattened list is no longer sorted, so a single binary search doesn't work. Start at the top-right corner instead: if the value is too big, move left (the whole column below is bigger); if too small, move down. Each step discards a row or a column, for O(m + n)."},
		},
	},
	{
		ID:              "koko-eating-bananas",
		Number:          875,
		Order:           3,
		Title:           "Koko Eating Bananas",
		Slug:            "koko-eating-bananas",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search"},
		RelatedChapters: []int{6},
		Description: `Koko has n piles of bananas; pile i holds piles[i] bananas. The guards will be back in h hours.

Koko picks an eating speed of k bananas per hour (a positive integer). Every hour she chooses one pile and eats k bananas from it. If that pile has fewer than k bananas left, she finishes the pile and does nothing else for the rest of that hour. So a pile of p bananas takes ceil(p / k) hours.

Return the smallest integer speed k that lets her finish all the piles within h hours.`,
		Constraints: []string{
			"1 <= piles.length <= 10^4",
			"piles.length <= h <= 10^9",
			"1 <= piles[i] <= 10^9",
		},
		Examples: []Example{
			{Input: "piles = [3,6,7,11], h = 8", Output: "4", Explanation: "At speed 4 the piles take 1 + 2 + 2 + 3 = 8 hours, which fits. At speed 3 they take 1 + 2 + 3 + 4 = 10 hours, which doesn't. So 4 is the smallest speed that works."},
			{Input: "piles = [30,11,23,4,20], h = 5", Output: "30", Explanation: "Five piles in five hours means one pile per hour, so the speed must cover the largest pile, 30."},
			{Input: "piles = [30,11,23,4,20], h = 6", Output: "23"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"piles": []int{3, 6, 7, 11}, "h": 8}, Expected: 4, Note: "basic case, answer in the middle of the range"},
			{Input: map[string]interface{}{"piles": []int{30, 11, 23, 4, 20}, "h": 5}, Expected: 30, Note: "h equals the number of piles, so the answer is max(piles) (catches hi = max(piles) - 1)"},
			{Input: map[string]interface{}{"piles": []int{30, 11, 23, 4, 20}, "h": 6}, Expected: 23, Note: "one spare hour"},
			{Input: map[string]interface{}{"piles": []int{1}, "h": 1}, Expected: 1, Note: "smallest input; the answer is the lower end of the range"},
			{Input: map[string]interface{}{"piles": []int{1000000000}, "h": 2}, Expected: 500000000, Hidden: true, Note: "huge pile: trying every speed from 1 is far too slow"},
			{Input: map[string]interface{}{"piles": []int{312884470}, "h": 968709470}, Expected: 1, Hidden: true, Note: "more hours than bananas, so speed 1 is enough (catches starting lo at 0 or above 1)"},
			{Input: map[string]interface{}{"piles": []int{5, 5, 5, 5}, "h": 4}, Expected: 5, Hidden: true, Note: "all piles equal, no spare hours"},
			{Input: map[string]interface{}{"piles": []int{5, 5, 5, 5}, "h": 8}, Expected: 3, Hidden: true, Note: "ceil matters: at speed 2 each pile of 5 takes 3 hours, not 2.5 (catches floor division)"},
			{Input: map[string]interface{}{"piles": []int{2, 2}, "h": 3}, Expected: 2, Hidden: true, Note: "speed 1 needs 4 hours, speed 2 needs 2: the answer is the first feasible speed"},
		},
		TimeComplexity:  "O(n log m), where m = max(piles)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def minEatingSpeed(piles, h):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If Koko can finish at speed 10, can she also finish at speed 11? What does that tell you about the list of all speeds, marked 'too slow' or 'fast enough'?"},
			{Level: 2, Type: "insight", Content: "Speeds line up as no, no, ..., no, yes, yes, ...: once a speed is fast enough, every faster one is too. The answer is the first yes. That's binary search on the answer: search over speeds 1..max(piles) (never over the array) with the boundary template from chapter 6: while lo < hi, and hi = mid when mid is fast enough.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Write the yes/no test: a helper that adds up ceil(pile / speed) over all piles and checks it is <= h.\n2. Define the search space: the answer lies in [1, max(piles)]; speed max(piles) always works because it finishes one pile per hour.\n3. Pick the middle: while lo < hi, mid = (lo + hi) // 2.\n4. Discard half: if mid is fast enough, the answer is mid or slower, so hi = mid; otherwise it is faster, so lo = mid + 1.\n5. Read the answer: when lo == hi, that speed is the first one that works."},
			{Level: 4, Type: "skeleton", Content: `def minEatingSpeed(piles, h):
    def hours_at(speed):
        return sum(____ for pile in piles)

    lo, hi = 1, ____
    while lo < hi:
        mid = (lo + hi) // 2
        if ____:
            hi = mid
        else:
            lo = ____
    return lo`},
		},
		Solution: Solution{
			Code: `def minEatingSpeed(piles, h):
    def hours_at(speed):
        return sum((pile + speed - 1) // speed for pile in piles)

    lo, hi = 1, max(piles)
    while lo < hi:
        mid = (lo + hi) // 2
        if hours_at(mid) <= h:
            hi = mid
        else:
            lo = mid + 1
    return lo`,
			Explanation: "The first problem where we binary-search something other than an array. Problems 704 and 74 searched positions in sorted data with the exact-value template; here there's no sorted array at all. What is sorted is the yes/no answer to 'can she finish at speed k?', which flips from no to yes exactly once as k grows. We search the speeds 1..max(piles) for the first yes with the boundary template (lo < hi, hi = mid).",
			BruteForce:  "Try k = 1, 2, 3, ... and return the first speed whose total hours fit in h. Each try costs O(n) and there can be up to max(piles) = 10^9 tries, so O(n * max(piles)): hopeless for piles = [1000000000], h = 2, where the answer is 500,000,000. It wastes work by testing speeds one at a time when one failed test at speed k already proves that every speed below k fails too.",
			KeyInsight:  "Feasibility is monotonic: eating faster never takes more hours. So the predicate 'hours_at(k) <= h' is false for all k below the answer and true from the answer on. A monotonic yes/no predicate can be binary-searched for its first true, which takes about log2(10^9) = 30 checks instead of up to 10^9.",
			WhyItWorks:  "Predicate: ok(k) = hours_at(k) <= h. Each term ceil(pile / k) can only drop or stay as k grows, so ok is monotonic: once true, true forever. ok(max(piles)) is true because each pile then takes one hour and h >= n. Invariant: the answer (the first k with ok(k)) lies in [lo, hi]. If ok(mid), the answer is mid or smaller, so hi = mid keeps it in range (mid itself might be the answer, so it is not excluded). If not ok(mid), every k <= mid fails, so lo = mid + 1. Because mid = (lo + hi) // 2 rounds down, mid < hi whenever lo < hi, so both branches shrink the range and the loop ends with lo == hi, the answer.",
			CommonMistakes: []string{
				"Using floor division pile // speed for the hours: a pile of 5 at speed 2 takes 3 hours, not 2. Use ceiling division, (pile + speed - 1) // speed or -(-pile // speed).",
				"Mixing the templates: while lo <= hi together with hi = mid loops forever once lo == hi and ok(mid) is true. For a first-true search use while lo < hi with hi = mid and lo = mid + 1.",
				"Starting lo at 0: the helper divides by the speed, so mid = 0 crashes with ZeroDivisionError. The slowest legal speed is 1.",
				"Starting hi at max(piles) - 1 or at sum(piles) // h: the true answer can be exactly max(piles) (h == len(piles)), and sum // h can be too small. Choose hi so that ok(hi) is guaranteed true.",
			},
			TimeComplexity:  "O(n log m), where m = max(piles)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Write the yes/no test", Explanation: "hours_at(speed) adds ceil(pile / speed) over all piles using integer ceiling division. The question 'is hours_at(speed) <= h?' is the monotonic predicate we search on.", CodeSnippet: "def hours_at(speed):\nreturn sum((pile + speed - 1) // speed for pile in piles)", LineStart: 2, LineEnd: 3},
				{Title: "Define the search space", Explanation: "The answer is at least 1 (speeds are positive) and at most max(piles), which always works because it clears one pile per hour and h >= len(piles).", CodeSnippet: "lo, hi = 1, max(piles)", LineStart: 5, LineEnd: 5},
				{Title: "Pick the middle", Explanation: "Boundary template: loop while the range has more than one candidate. mid rounds down, so mid < hi and setting hi = mid always shrinks the range.", CodeSnippet: "while lo < hi:\nmid = (lo + hi) // 2", LineStart: 6, LineEnd: 7},
				{Title: "Discard half", Explanation: "For [3,6,7,11], h = 8: lo=1, hi=11, mid=6 takes 6 hours (yes) so hi=6; mid=3 takes 10 (no) so lo=4; mid=5 takes 8 (yes) so hi=5; mid=4 takes 8 (yes) so hi=4. Now lo == hi == 4.", CodeSnippet: "if hours_at(mid) <= h:\nhi = mid\nlo = mid + 1", LineStart: 8, LineEnd: 11},
				{Title: "Read the answer", Explanation: "lo == hi is the only speed left in the range, and by the invariant it is the slowest speed that fits.", CodeSnippet: "return lo", LineStart: 12, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Linear search upward from a lower bound",
					Code: `def minEatingSpeed(piles, h):
    total = sum(piles)
    speed = max(1, -(-total // h))
    while True:
        hours = 0
        for pile in piles:
            hours += -(-pile // speed)
        if hours <= h:
            return speed
        speed += 1`,
					TimeComplexity:  "O(n * m) in the worst case, where m = max(piles)",
					SpaceComplexity: "O(1)",
					Explanation:     "The brute force tries speeds one at a time. Starting from 1 would take up to 10^9 tries, so this version starts at ceil(sum(piles) / h): she must eat at least that many bananas per hour on average, so nothing slower can work. It's fast on most inputs but still linear in the worst case, which is why binary search wins.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is hi = mid (not mid - 1) when mid is fast enough, and why doesn't that loop forever?", Answer: "mid might be the slowest speed that works, so throwing it away could lose the answer; hi = mid keeps it. It can't loop forever because mid = (lo + hi) // 2 rounds down: while lo < hi, mid is strictly less than hi, so hi = mid always shrinks the range. The other branch, lo = mid + 1, shrinks it too."},
			{Prompt: "What makes a problem a candidate for binary search on the answer?", Answer: "You're asked for the minimum (or maximum) value that satisfies a condition, the candidate values have a known range, and the condition is monotonic: if value x works, every larger x works too (or every smaller one). Then you can test the middle value with a check that is usually O(n), and halve the range. Examples: ship capacity within D days, split an array into k parts with the smallest largest sum."},
			{Prompt: "What if Koko had to find the LARGEST speed at which she still needs at least h hours?", Answer: "The predicate hours_at(k) >= h is true for small k and false for large k, so you want the last true. Find the first k where it is false with the same lo < hi template (hi = mid on false, lo = mid + 1 on true) and subtract one, or flip the predicate and keep the first-true code."},
		},
	},
	{
		ID:              "find-minimum-in-rotated-sorted-array",
		Number:          153,
		Order:           4,
		Title:           "Find Minimum in Rotated Sorted Array",
		Slug:            "find-minimum-in-rotated-sorted-array",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search"},
		RelatedChapters: []int{6},
		Description: `An array of n distinct integers was sorted in ascending order and then rotated: its last element was moved to the front some number of times (between 1 and n). For example, [0,1,2,4,5,6,7] rotated 4 times becomes [4,5,6,7,0,1,2], and rotated 7 times it is back to [0,1,2,4,5,6,7].

Given the rotated array nums, return its smallest element.

Your algorithm must run in O(log n) time.`,
		Constraints: []string{
			"n == nums.length",
			"1 <= n <= 5000",
			"-5000 <= nums[i] <= 5000",
			"All the integers of nums are unique",
			"nums is sorted and rotated between 1 and n times",
		},
		Examples: []Example{
			{Input: "nums = [3,4,5,1,2]", Output: "1", Explanation: "The original array was [1,2,3,4,5], rotated 3 times. The array is two ascending runs, [3,4,5] and [1,2]; the minimum is where the second run starts."},
			{Input: "nums = [4,5,6,7,0,1,2]", Output: "0"},
			{Input: "nums = [11,13,15,17]", Output: "11", Explanation: "Rotated n times, the array is in its original order, so the minimum is the first element."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{3, 4, 5, 1, 2}}, Expected: 1, Note: "minimum in the right half"},
			{Input: map[string]interface{}{"nums": []int{4, 5, 6, 7, 0, 1, 2}}, Expected: 0, Note: "classic rotation"},
			{Input: map[string]interface{}{"nums": []int{11, 13, 15, 17}}, Expected: 11, Note: "not rotated at all (catches comparing with nums[0] instead of the last element)"},
			{Input: map[string]interface{}{"nums": []int{1}}, Expected: 1, Note: "single element"},
			{Input: map[string]interface{}{"nums": []int{2, 1}}, Expected: 1, Hidden: true, Note: "two elements, rotated (catches lo <= hi with hi = mid looping forever)"},
			{Input: map[string]interface{}{"nums": []int{5, 1, 2, 3, 4}}, Expected: 1, Hidden: true, Note: "minimum at index 1"},
			{Input: map[string]interface{}{"nums": []int{2, 3, 4, 5, 1}}, Expected: 1, Hidden: true, Note: "minimum is the last element"},
			{Input: map[string]interface{}{"nums": []int{0, 2, -5, -3}}, Expected: -5, Hidden: true, Note: "negative values"},
			{Input: map[string]interface{}{"nums": []int{3, 1, 2}}, Expected: 1, Hidden: true, Note: "the first true is mid itself (catches hi = mid - 1)"},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def findMin(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Compare each element of [4,5,6,7,0,1,2] with the last element, 2. Which elements are bigger than it, which are not, and in what pattern do the answers come?"},
			{Level: 2, Type: "insight", Content: "Elements before the minimum are all greater than nums[-1]; the minimum and everything after it are <= nums[-1]. So the question 'is nums[i] <= nums[-1]?' gives no, no, ..., yes, yes, and the minimum is at the first yes. Use the same first-true boundary template as Koko Eating Bananas (875): while lo < hi, hi = mid on yes, lo = mid + 1 on no.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Write the yes/no test: is nums[mid] <= the last element?\n2. Define the search space: indices 0 to n - 1; the minimum is always somewhere in there.\n3. Pick the middle: while lo < hi, mid = (lo + hi) // 2.\n4. Discard half: yes means mid is in the low run, so the minimum is at mid or to its left (hi = mid); no means mid is in the high run, so the minimum is to its right (lo = mid + 1).\n5. Read the answer: return nums[lo]."},
			{Level: 4, Type: "skeleton", Content: `def findMin(nums):
    last = nums[-1]
    lo, hi = 0, len(nums) - 1
    while lo < hi:
        mid = (lo + hi) // 2
        if ____:
            hi = ____
        else:
            lo = ____
    return nums[lo]`},
		},
		Solution: Solution{
			Code: `def findMin(nums):
    last = nums[-1]
    lo, hi = 0, len(nums) - 1
    while lo < hi:
        mid = (lo + hi) // 2
        if nums[mid] <= last:
            hi = mid
        else:
            lo = mid + 1
    return nums[lo]`,
			Explanation: "Builds on Koko Eating Bananas (875): it's the same first-true boundary search, but now the predicate is about the array itself. A rotated sorted array is two ascending runs, and every element of the first run is bigger than the last element of the array. 'nums[i] <= nums[-1]' is therefore false on the first run and true on the second, and the minimum is the first index where it turns true.",
			BruteForce:  "Return min(nums), or scan for the one place where nums[i] < nums[i - 1]. Both are O(n): up to 5,000 comparisons. They ignore the fact that the array is made of two sorted runs, which lets one comparison at the middle tell you which run you're in and discard half.",
			KeyInsight:  "Comparing an element with the last element tells you which run it belongs to: bigger means the first (high) run, not bigger means the second (low) run, which starts with the minimum. That turns the problem into finding the first index of the low run, a monotonic yes/no search.",
			WhyItWorks:  "Write nums as A + B, where A and B are ascending, and every value in A exceeds every value in B (A may be empty when the array isn't rotated). The last element belongs to B, so every element of A is > last and every element of B is <= last. The predicate p(i) = nums[i] <= last is false on A and true on B, so it's monotonic, and its first true index is the start of B, which holds the minimum (or index 0 when A is empty). The boundary template keeps the invariant 'the first true index is in [lo, hi]': p(mid) true means it's at mid or earlier (hi = mid), false means later (lo = mid + 1). mid rounds down, so the range shrinks every step and stops at lo == hi.",
			CommonMistakes: []string{
				"Comparing with nums[0] instead of the last element: for an unrotated array like [11,13,15,17] every element is >= nums[0], so 'nums[mid] < nums[0]' is never true and the search drifts to the end and returns 17. The last element always belongs to the run that contains the minimum.",
				"Using while lo <= hi with hi = mid: once lo == hi and the predicate is true, hi = mid changes nothing and the loop never ends. Boundary searches use while lo < hi.",
				"Setting hi = mid - 1 when the predicate is true: mid itself may be the minimum. In [3,1,2], mid = 1 holds the minimum 1 (1 <= 2 is true), hi = mid - 1 = 0 throws it away, and the search returns 3. On a true answer keep mid with hi = mid.",
			},
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Write the yes/no test", Explanation: "Remember the last element. 'nums[i] <= last' marks the low run, which starts with the minimum.", CodeSnippet: "last = nums[-1]", LineStart: 2, LineEnd: 2},
				{Title: "Define the search space", Explanation: "The minimum's index is somewhere in 0..n-1, and unlike the exact-value template we never need to report 'not found'.", CodeSnippet: "lo, hi = 0, len(nums) - 1", LineStart: 3, LineEnd: 3},
				{Title: "Pick the middle", Explanation: "Loop while more than one candidate remains. Because mid rounds down, mid < hi.", CodeSnippet: "while lo < hi:\nmid = (lo + hi) // 2", LineStart: 4, LineEnd: 5},
				{Title: "Discard half", Explanation: "For [4,5,6,7,0,1,2] (last = 2): lo=0, hi=6, mid=3 holds 7 > 2 so lo = 4; mid=5 holds 1 <= 2 so hi = 5; mid=4 holds 0 <= 2 so hi = 4. Now lo == hi == 4.", CodeSnippet: "if nums[mid] <= last:\nhi = mid\nlo = mid + 1", LineStart: 6, LineEnd: 9},
				{Title: "Read the answer", Explanation: "lo is the first index of the low run, which holds the minimum: nums[4] = 0.", CodeSnippet: "return nums[lo]", LineStart: 10, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{
					Name: "Find the drop by scanning (brute force)",
					Code: `def findMin(nums):
    for i in range(1, len(nums)):
        if nums[i] < nums[i - 1]:
            return nums[i]
    return nums[0]`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "The minimum is the only element smaller than its left neighbour; if no such drop exists the array isn't rotated and the minimum is nums[0]. Correct, but it may look at every element.",
				},
				{
					Name:            "Compare with nums[hi] instead of the fixed last element",
					Code:            "def findMin(nums):\n    lo, hi = 0, len(nums) - 1\n    while lo < hi:\n        mid = (lo + hi) // 2\n        if nums[mid] > nums[hi]:\n            lo = mid + 1\n        else:\n            hi = mid\n    return nums[lo]",
					TimeComplexity:  "O(log n)",
					SpaceComplexity: "O(1)",
					Explanation:     "The version you'll see most often. It works because nums[lo..hi] always still contains the minimum and is itself a rotated sorted array, so its own last element nums[hi] plays the role of nums[-1].",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why compare with the last element rather than the first?", Answer: "The last element always belongs to the run that contains the minimum, so 'nums[i] <= last' is true exactly on that run, even when the array isn't rotated at all. The first element belongs to the other run, and when there's no rotation that run is the whole array: 'nums[i] < nums[0]' is then never true and the first-true search returns the wrong end."},
			{Prompt: "What if duplicates were allowed, as in [2,2,2,0,2] (LeetCode 154)?", Answer: "The predicate breaks: nums[mid] == last no longer tells you which run mid is in (in [2,2,2,0,2], index 1 equals the last element but is before the minimum). The usual fix compares with nums[hi] and, on equality, just does hi -= 1, which is safe because nums[hi] has a copy at mid. The worst case becomes O(n), e.g. all elements equal."},
		},
	},
	{
		ID:              "search-in-rotated-sorted-array",
		Number:          33,
		Order:           5,
		Title:           "Search in Rotated Sorted Array",
		Slug:            "search-in-rotated-sorted-array",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search"},
		RelatedChapters: []int{6},
		Description: `An ascending array of distinct integers may have been rotated at some unknown index k: [a0, a1, ..., an-1] becomes [ak, ak+1, ..., an-1, a0, a1, ..., ak-1]. For example, [0,1,2,4,5,6,7] rotated at index 3 becomes [4,5,6,7,0,1,2]. The array might also not be rotated at all.

Given the (possibly rotated) array nums and an integer target, return the index of target in nums, or -1 if it isn't there.

Your algorithm must run in O(log n) time.`,
		Constraints: []string{
			"1 <= nums.length <= 5000",
			"-10^4 <= nums[i] <= 10^4",
			"All values of nums are unique",
			"nums is an ascending array that is possibly rotated",
			"-10^4 <= target <= 10^4",
		},
		Examples: []Example{
			{Input: "nums = [4,5,6,7,0,1,2], target = 0", Output: "4", Explanation: "The middle is 7 at index 3. The left half [4,5,6,7] is sorted and 0 isn't in [4,7], so search the right half. Its middle is 1 at index 5; the half [0,1] is sorted and contains 0, so go left and find 0 at index 4."},
			{Input: "nums = [4,5,6,7,0,1,2], target = 3", Output: "-1"},
			{Input: "nums = [1], target = 0", Output: "-1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{4, 5, 6, 7, 0, 1, 2}, "target": 0}, Expected: 4, Note: "target in the rotated (low) part"},
			{Input: map[string]interface{}{"nums": []int{4, 5, 6, 7, 0, 1, 2}, "target": 3}, Expected: -1, Note: "target missing"},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 0}, Expected: -1, Note: "single element, no match"},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 1}, Expected: 0, Note: "single element, match"},
			{Input: map[string]interface{}{"nums": []int{3, 1}, "target": 1}, Expected: 1, Hidden: true, Note: "lo == mid: catches nums[lo] < nums[mid] (strict) choosing the wrong half"},
			{Input: map[string]interface{}{"nums": []int{3, 1}, "target": 3}, Expected: 0, Hidden: true, Note: "two elements, target first"},
			{Input: map[string]interface{}{"nums": []int{5, 1, 3}, "target": 5}, Expected: 0, Hidden: true, Note: "right half sorted but target is in the left"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5, 6}, "target": 4}, Expected: 3, Hidden: true, Note: "not rotated"},
			{Input: map[string]interface{}{"nums": []int{4, 5, 6, 7, 0, 1, 2}, "target": 5}, Expected: 1, Hidden: true, Note: "target in the high part"},
			{Input: map[string]interface{}{"nums": []int{4, 5, 6, 7, 0, 1, 2}, "target": 2}, Expected: 6, Hidden: true, Note: "target is the last element (needs target <= nums[hi], not <)"},
			{Input: map[string]interface{}{"nums": []int{6, 7, 1, 2, 3, 4, 5}, "target": 6}, Expected: 0, Hidden: true, Note: "target is the first element (needs nums[lo] <= target, not <)"},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def search(nums, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Cut [4,5,6,7,0,1,2] at its middle element 7. Is the left piece [4,5,6,7] sorted? Is the right piece [7,0,1,2]? Will one of the two pieces always be sorted, and what can you check easily about a sorted piece?"},
			{Level: 2, Type: "insight", Content: "Whatever mid you pick, at least one of nums[lo..mid] and nums[mid..hi] is in sorted order. For the sorted half you can tell in O(1) whether target is inside it, by checking its two ends. If it is, search there; otherwise search the other half. It's the exact-value template from Binary Search (704), with a different rule for which half to discard. See 'Find the sorted half' in chapter 6.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Define the search space: lo = 0, hi = n - 1, loop while lo <= hi.\n2. Pick the middle, and write the yes/no test: if nums[mid] == target, return mid.\n3. Find the sorted half: if nums[lo] <= nums[mid], the left half is sorted; otherwise the right half is.\n4. Discard half: if target lies between the ends of the sorted half, keep that half; otherwise keep the other one.\n5. Read the answer: an empty range means -1."},
			{Level: 4, Type: "skeleton", Content: `def search(nums, target):
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] == target:
            return mid
        if ____:
            if ____:
                hi = mid - 1
            else:
                lo = mid + 1
        else:
            if ____:
                lo = mid + 1
            else:
                hi = mid - 1
    return -1`},
		},
		Solution: Solution{
			Code: `def search(nums, target):
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] == target:
            return mid
        if nums[lo] <= nums[mid]:
            if nums[lo] <= target < nums[mid]:
                hi = mid - 1
            else:
                lo = mid + 1
        else:
            if nums[mid] < target <= nums[hi]:
                lo = mid + 1
            else:
                hi = mid - 1
    return -1`,
			Explanation: "Combines the two previous ideas: the exact-value template from Binary Search (704) and the 'which run am I in?' reasoning from Find Minimum (153). The rotation breaks the rule 'too small means go right', so instead we look at whichever half is sorted, ask whether target falls between its two ends, and discard the half that can't hold it.",
			BruteForce:  "Scan every element and return the index of target: O(n), up to 5,000 comparisons, ignoring that the array is two sorted runs. Another slow-ish route is to sort a copy with original indices, O(n log n), which throws away the order we were given.",
			KeyInsight:  "A rotated sorted array has a single drop. Any range [lo, hi] split at mid puts that drop in at most one half, so the other half is sorted. For a sorted half, 'is target in it?' is just two comparisons with its ends, so you always know which half to discard.",
			WhyItWorks:  "Invariant: if target is in nums, its index is in [lo, hi]. The drop sits in at most one of nums[lo..mid] and nums[mid..hi], so one is sorted, and nums[lo] <= nums[mid] detects that the left one is (with <=, because when lo == mid the one-element left half is trivially sorted). If the left half is sorted and nums[lo] <= target < nums[mid], target can only be left of mid, since every value in the other half lies outside that interval; otherwise it can only be right of mid. The right-sorted case is symmetric. mid is excluded after checking it, so the range shrinks every step and an empty range proves target is absent.",
			CommonMistakes: []string{
				"Testing nums[lo] < nums[mid] (strict): when lo == mid, as in [3,1] with target 1, the left 'half' is the single element 3, but the strict test says the right half is sorted, checks 3 < 1 <= 1, fails, and discards the half that holds 1. Use <=.",
				"Using open comparisons at the ends, such as nums[lo] < target or target < nums[hi]: target equal to the end value (the first or last element) is wrongly sent to the other half. The sorted half's own ends are inclusive; only mid is excluded, because it was already checked.",
				"Trying to compare target with nums[mid] alone, as in plain binary search: in [4,5,6,7,0,1,2], target 0 is smaller than 7, but it is to the right. You must first know which half is sorted.",
			},
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the search space", Explanation: "Closed range over all indices, exact-value template, because we're looking for one specific value that may be absent.", CodeSnippet: "lo, hi = 0, len(nums) - 1", LineStart: 2, LineEnd: 2},
				{Title: "Pick the middle", Explanation: "Loop while at least one candidate remains and take the middle.", CodeSnippet: "while lo <= hi:\nmid = (lo + hi) // 2", LineStart: 3, LineEnd: 4},
				{Title: "Write the yes/no test", Explanation: "If the middle is the target, we're done.", CodeSnippet: "if nums[mid] == target:\nreturn mid", LineStart: 5, LineEnd: 6},
				{Title: "Find the sorted half", Explanation: "If nums[lo] <= nums[mid], there's no drop between lo and mid, so the left half is sorted; otherwise the drop is there and the right half is sorted. For [4,5,6,7,0,1,2]: lo=0, mid=3, 4 <= 7, so the left half is sorted.", CodeSnippet: "if nums[lo] <= nums[mid]:", LineStart: 7, LineEnd: 7},
				{Title: "Discard half", Explanation: "Target 0 isn't in [4, 7), so we drop the left half: lo = 4. Next mid = 5 (value 1); nums[4] = 0 <= 1, so the left half [0,1] is sorted and 0 <= 0 < 1 holds, so hi = 4. Then mid = 4 matches.", CodeSnippet: "if nums[lo] <= target < nums[mid]:\nhi = mid - 1\nlo = mid + 1\nif nums[mid] < target <= nums[hi]:", LineStart: 8, LineEnd: 16},
				{Title: "Read the answer", Explanation: "An empty range means target was never in the array.", CodeSnippet: "return -1", LineStart: 17, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{
					Name: "Linear scan (brute force)",
					Code: `def search(nums, target):
    for i, value in enumerate(nums):
        if value == target:
            return i
    return -1`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Checks every element. Correct and simple, but ignores the structure.",
				},
				{
					Name: "Find the rotation point, then search one run",
					Code: `def search(nums, target):
    last = nums[-1]
    lo, hi = 0, len(nums) - 1
    while lo < hi:
        mid = (lo + hi) // 2
        if nums[mid] <= last:
            hi = mid
        else:
            lo = mid + 1
    start = lo
    if target <= last:
        lo, hi = start, len(nums) - 1
    else:
        lo, hi = 0, start - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] == target:
            return mid
        if nums[mid] < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return -1`,
					TimeComplexity:  "O(log n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Reuses Find Minimum (153) directly: a first-true search finds where the low run starts, then target <= nums[-1] tells you which run it belongs to, and a plain 704 search finishes the job. Two simple binary searches instead of one clever one; many people find this easier to get right.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is at least one half always sorted?", Answer: "The rotated array has at most one place where a value is followed by a smaller value (the drop). Splitting [lo, hi] at mid gives two pieces, and a single drop can't be in both, so at least one piece has no drop and is therefore in ascending order."},
			{Prompt: "Why does the test use nums[lo] <= nums[mid] with <=, not <?", Answer: "When the range has two elements, lo == mid, so the left half is just nums[mid]. It is sorted, and <= says so because nums[lo] == nums[mid]. With <, the code would treat the right half as sorted even when the drop is there, e.g. [3,1] with target 1 goes to the wrong side and returns -1."},
			{Prompt: "What would change with duplicates allowed (LeetCode 81)?", Answer: "When nums[lo] == nums[mid] == nums[hi] you can't tell which half is sorted, e.g. [1,1,1,3,1]. The usual fix is to shrink both ends by one (lo += 1, hi -= 1) in that case, which keeps correctness but makes the worst case O(n)."},
		},
	},
	{
		ID:              "time-based-key-value-store",
		Number:          981,
		Order:           6,
		Title:           "Time Based Key-Value Store",
		Slug:            "time-based-key-value-store",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Hash Table", "String", "Binary Search", "Design"},
		RelatedChapters: []int{6, 2},
		Description: `Build a key-value store that remembers every value a key has had, each tagged with the time it was set, and can answer "what was the value of this key at time t?".

Implement the class TimeMap:
- TimeMap() creates an empty store.
- set(key, value, timestamp) records that key had value starting at time timestamp.
- get(key, timestamp) returns the value from the set call for key with the largest timestamp that is <= timestamp. If key was never set at or before that time (or never set at all), return the empty string "".

The timestamps passed to set are strictly increasing over all calls.

How the tests call your class: each test lists operations and their arguments. The first operation constructs TimeMap; the expected output has null for the constructor and for every set, and the returned string for every get.`,
		Constraints: []string{
			"1 <= key.length, value.length <= 100",
			"key and value consist of lowercase English letters and digits",
			"1 <= timestamp <= 10^7",
			"All the timestamps of set calls are strictly increasing",
			"At most 2 * 10^5 calls will be made to set and get",
		},
		Examples: []Example{
			{Input: `["TimeMap","set","get","get","set","get","get"]
[[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]`, Output: `[null,null,"bar","bar",null,"bar2","bar2"]`, Explanation: "At time 1 foo is set to bar. get(foo, 3): the latest set at or before 3 is the one at time 1, so bar. After foo is set to bar2 at time 4, get(foo, 4) and get(foo, 5) both see bar2."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"TimeMap", "set", "get", "get", "set", "get", "get"}, "values": []interface{}{[]interface{}{}, []interface{}{"foo", "bar", 1}, []interface{}{"foo", 1}, []interface{}{"foo", 3}, []interface{}{"foo", "bar2", 4}, []interface{}{"foo", 4}, []interface{}{"foo", 5}}}, Expected: []interface{}{nil, nil, "bar", "bar", nil, "bar2", "bar2"}, Note: "the example: exact match, a time between sets, and a time after the last set"},
			{Input: map[string]interface{}{"operations": []string{"TimeMap", "get", "set", "get", "get"}, "values": []interface{}{[]interface{}{}, []interface{}{"a", 1}, []interface{}{"a", "x", 5}, []interface{}{"a", 4}, []interface{}{"a", 5}}}, Expected: []interface{}{nil, "", nil, "", "x"}, Note: "key never set, and a time before the first set: both return the empty string"},
			{Input: map[string]interface{}{"operations": []string{"TimeMap", "set", "set", "get", "get", "get", "get", "get"}, "values": []interface{}{[]interface{}{}, []interface{}{"love", "high", 10}, []interface{}{"love", "low", 20}, []interface{}{"love", 5}, []interface{}{"love", 10}, []interface{}{"love", 15}, []interface{}{"love", 20}, []interface{}{"love", 25}}}, Expected: []interface{}{nil, nil, nil, "", "high", "high", "low", "low"}, Note: "before, at, between and after two timestamps (catches returning the first timestamp >= t)"},
			{Input: map[string]interface{}{"operations": []string{"TimeMap", "set", "set", "set", "get", "get", "get", "get", "get"}, "values": []interface{}{[]interface{}{}, []interface{}{"a", "1", 1}, []interface{}{"b", "2", 2}, []interface{}{"a", "3", 3}, []interface{}{"a", 2}, []interface{}{"b", 1}, []interface{}{"b", 3}, []interface{}{"a", 3}, []interface{}{"c", 3}}}, Expected: []interface{}{nil, nil, nil, nil, "1", "", "2", "3", ""}, Hidden: true, Note: "two keys interleaved: each key keeps its own history"},
			{Input: map[string]interface{}{"operations": []string{"TimeMap", "set", "set", "set", "set", "set", "set", "set", "set", "get", "get", "get", "get", "get", "get"}, "values": []interface{}{[]interface{}{}, []interface{}{"k", "v1", 2}, []interface{}{"k", "v2", 4}, []interface{}{"k", "v3", 6}, []interface{}{"k", "v4", 8}, []interface{}{"k", "v5", 10}, []interface{}{"k", "v6", 12}, []interface{}{"k", "v7", 14}, []interface{}{"k", "v8", 16}, []interface{}{"k", 1}, []interface{}{"k", 2}, []interface{}{"k", 9}, []interface{}{"k", 13}, []interface{}{"k", 16}, []interface{}{"k", 17}}}, Expected: []interface{}{nil, nil, nil, nil, nil, nil, nil, nil, nil, "", "v1", "v4", "v6", "v8", "v8"}, Hidden: true, Note: "longer history: off-by-one errors in the search show up between timestamps"},
			{Input: map[string]interface{}{"operations": []string{"TimeMap", "set", "get", "get", "get"}, "values": []interface{}{[]interface{}{}, []interface{}{"x", "only", 100}, []interface{}{"x", 99}, []interface{}{"x", 100}, []interface{}{"x", 10000000}}}, Expected: []interface{}{nil, nil, "", "only", "only"}, Hidden: true, Note: "single entry: just before, exactly at, and long after"},
		},
		TimeComplexity:  "O(1) per set, O(log n) per get",
		SpaceComplexity: "O(n)",
		StarterCode: `class TimeMap:
    def __init__(self):
        pass

    def set(self, key, value, timestamp):
        pass

    def get(self, key, timestamp):
        pass`,
		Harness: Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Because set timestamps only ever increase, what order is each key's list of (timestamp, value) entries in if you just append them? What are you then looking for in that list on get?"},
			{Level: 2, Type: "insight", Content: "Map each key to a list of [timestamp, value] entries; appending keeps it sorted by timestamp for free. get wants the LAST entry with timestamp <= t, which is one before the FIRST entry with timestamp > t. That's a first-true boundary search (chapter 6), on the predicate 'entry timestamp > t'.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. set: append [timestamp, value] to the key's list (create the list if needed).\n2. Define the search space: for get, lo = 0 and hi = len(entries), so 'no entry is > t' is a possible answer (index len).\n3. Write the yes/no test: is entries[mid]'s timestamp > t?\n4. Discard half: yes means hi = mid, no means lo = mid + 1, while lo < hi.\n5. Read the answer: lo is the first entry after t; if lo == 0 nothing is at or before t, so return \"\"; otherwise return the value at lo - 1."},
			{Level: 4, Type: "skeleton", Content: `class TimeMap:
    def __init__(self):
        self.history = {}

    def set(self, key, value, timestamp):
        if key not in self.history:
            self.history[key] = []
        self.history[key].append([timestamp, value])

    def get(self, key, timestamp):
        entries = self.history.get(key, [])
        lo, hi = 0, ____
        while lo < hi:
            mid = (lo + hi) // 2
            if ____:
                hi = mid
            else:
                lo = mid + 1
        if lo == 0:
            return ""
        return ____`},
		},
		Solution: Solution{
			Code: `class TimeMap:
    def __init__(self):
        self.history = {}

    def set(self, key, value, timestamp):
        if key not in self.history:
            self.history[key] = []
        self.history[key].append([timestamp, value])

    def get(self, key, timestamp):
        entries = self.history.get(key, [])
        lo, hi = 0, len(entries)
        while lo < hi:
            mid = (lo + hi) // 2
            if entries[mid][0] > timestamp:
                hi = mid
            else:
                lo = mid + 1
        if lo == 0:
            return ""
        return entries[lo - 1][1]`,
			Explanation: "A hash map (chapter 2) from key to that key's history, plus the boundary search from Koko (875) and Find Minimum (153). The new twist is that the answer can be 'nothing qualifies', so the search space includes one extra slot past the end, and we look for the first entry that is too late and step back by one.",
			BruteForce:  "On get, walk the key's history backwards from the newest entry and return the first one with timestamp <= t. That is O(n) per get for a key with n entries; with 2 * 10^5 calls on one key that's up to 10^10 steps. It wastes the fact that the history is already sorted by timestamp.",
			KeyInsight:  "Appending with strictly increasing timestamps keeps each history sorted without any extra work, so 'latest entry at or before t' is a boundary in sorted data: entries with timestamp <= t form a prefix, and we want the last element of that prefix.",
			WhyItWorks:  "Each key's entries are sorted by timestamp because set calls arrive in increasing time order and we only append. The predicate p(i) = entries[i][0] > t is false on a prefix and true on the rest, so it's monotonic. With lo = 0 and hi = len(entries) (the index 'past the end' stands for 'no entry is after t'), the first-true search keeps the invariant 'the first true index is in [lo, hi]' and ends with lo equal to it. Every entry before lo has timestamp <= t, and entries[lo - 1] is the latest of them. If lo == 0, no entry is at or before t and the answer is the empty string. The comparison only ever looks at entries[mid][0], the timestamp, so values are never compared.",
			CommonMistakes: []string{
				"Searching for the first timestamp >= t and returning it: for history {10: high, 20: low}, get at 15 would return low, a value from the future. You want the last timestamp <= t, which is one before the first timestamp > t.",
				"Starting with hi = len(entries) - 1: when every entry is <= t (get after the last set), the true boundary is len(entries), which is now outside the range, so the search returns the second-to-last value. Include the past-the-end slot.",
				"Forgetting the lo == 0 case: entries[lo - 1] is then entries[-1], which in Python silently returns the NEWEST value instead of the empty string.",
				"Using bisect on a list of (timestamp, value) tuples with (t, '') or (t, chr(127)) as the probe: tuples compare value strings on ties, so the result depends on the sentinel. Keep timestamps in their own list, or compare entries[mid][0] directly as here.",
			},
			TimeComplexity:  "O(1) per set, O(log n) per get",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Store the history", Explanation: "A dict maps each key to its list of [timestamp, value] entries. Appending keeps the list sorted by time, because timestamps only increase.", CodeSnippet: "self.history = {}\nself.history[key].append([timestamp, value])", LineStart: 3, LineEnd: 8},
				{Title: "Define the search space", Explanation: "Indices 0..len(entries), one more than the last index. The extra slot means 'no entry is after t'. A key that was never set gets an empty list, so lo = hi = 0.", CodeSnippet: "entries = self.history.get(key, [])\nlo, hi = 0, len(entries)", LineStart: 11, LineEnd: 12},
				{Title: "Write the yes/no test", Explanation: "Is this entry later than t? That's false for a prefix of the list and true afterwards.", CodeSnippet: "if entries[mid][0] > timestamp:", LineStart: 15, LineEnd: 15},
				{Title: "Discard half", Explanation: "Too late: the first late entry is mid or before, so hi = mid. Not late: it's after mid, so lo = mid + 1. With history at times [2,4,...,16] and t = 9: lo=0,hi=8, mid=4 (10 > 9) hi=4; mid=2 (6) lo=3; mid=3 (8) lo=4. First late entry is index 4.", CodeSnippet: "while lo < hi:\nmid = (lo + hi) // 2\nhi = mid\nlo = mid + 1", LineStart: 13, LineEnd: 18},
				{Title: "Read the answer", Explanation: "lo - 1 is the last entry at or before t (index 3, value v4 in the trace). If lo is 0, nothing qualifies, so return the empty string.", CodeSnippet: "if lo == 0:\nreturn \"\"\nreturn entries[lo - 1][1]", LineStart: 19, LineEnd: 21},
			},
			Alternatives: []AltSolution{
				{
					Name: "Scan backwards (brute force)",
					Code: `class TimeMap:
    def __init__(self):
        self.history = {}

    def set(self, key, value, timestamp):
        self.history.setdefault(key, []).append((timestamp, value))

    def get(self, key, timestamp):
        for time, value in reversed(self.history.get(key, [])):
            if time <= timestamp:
                return value
        return ""`,
					TimeComplexity:  "O(1) per set, O(n) per get",
					SpaceComplexity: "O(n)",
					Explanation:     "Walks from the newest entry back until it finds one at or before t. Correct and short, but a long history makes every get slow.",
				},
				{
					Name: "Separate timestamp list with bisect",
					Code: `import bisect

class TimeMap:
    def __init__(self):
        self.times = {}
        self.values = {}

    def set(self, key, value, timestamp):
        self.times.setdefault(key, []).append(timestamp)
        self.values.setdefault(key, []).append(value)

    def get(self, key, timestamp):
        times = self.times.get(key, [])
        index = bisect.bisect_right(times, timestamp)
        if index == 0:
            return ""
        return self.values[key][index - 1]`,
					TimeComplexity:  "O(1) per set, O(log n) per get",
					SpaceComplexity: "O(n)",
					Explanation:     "bisect_right(times, t) returns exactly the first index whose timestamp is > t, the same boundary the hand-written loop finds. Keeping timestamps in their own list means bisect compares plain integers, so no tuple sentinel tricks are needed.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is hi set to len(entries) rather than len(entries) - 1?", Answer: "We search for the first entry with timestamp > t, and it's possible that no entry qualifies (every set happened at or before t). Index len(entries) stands for that outcome. With hi = len(entries) - 1 the search could never return it, and a get after the last set would return the second-newest value."},
			{Prompt: "Why are each key's entries sorted without ever calling sort?", Answer: "The problem guarantees that set timestamps are strictly increasing over all calls, so any timestamp appended to a key's list is larger than every timestamp already in it. Appending to the end therefore keeps the list in ascending order."},
			{Prompt: "What if set calls could arrive out of time order?", Answer: "Appending would no longer keep the list sorted. You'd need to insert in sorted position (bisect.insort, O(n) per set because of shifting) or use a balanced tree or sorted container. Binary search on get only works on sorted data."},
		},
	},
	{
		ID:              "median-of-two-sorted-arrays",
		Number:          4,
		Order:           7,
		Title:           "Median of Two Sorted Arrays",
		Slug:            "median-of-two-sorted-arrays",
		Difficulty:      "Hard",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search", "Divide and Conquer"},
		RelatedChapters: []int{6, 3},
		Description: `You are given two arrays nums1 (length m) and nums2 (length n), each sorted in ascending order. Either one may be empty, but not both.

Return the median of all m + n numbers together. The median is the middle value of the combined sorted list when m + n is odd, and the average of the two middle values when m + n is even. Return it as a number (for example 2.0 or 2.5).

Your algorithm must run in O(log(m + n)) time, so merging the arrays is not allowed.`,
		Constraints: []string{
			"nums1.length == m",
			"nums2.length == n",
			"0 <= m <= 1000",
			"0 <= n <= 1000",
			"1 <= m + n <= 2000",
			"-10^6 <= nums1[i], nums2[i] <= 10^6",
		},
		Examples: []Example{
			{Input: "nums1 = [1,3], nums2 = [2]", Output: "2.0", Explanation: "Together the numbers are [1,2,3]; the middle one is 2."},
			{Input: "nums1 = [1,2], nums2 = [3,4]", Output: "2.5", Explanation: "Together they are [1,2,3,4]; the two middle values are 2 and 3, and (2 + 3) / 2 = 2.5."},
			{Input: "nums1 = [], nums2 = [2,3]", Output: "2.5", Explanation: "One array can be empty; the median then comes from the other alone."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums1": []int{1, 3}, "nums2": []int{2}}, Expected: 2.0, Note: "odd total"},
			{Input: map[string]interface{}{"nums1": []int{1, 2}, "nums2": []int{3, 4}}, Expected: 2.5, Note: "even total, every element of nums1 before nums2"},
			{Input: map[string]interface{}{"nums1": []int{}, "nums2": []int{1}}, Expected: 1.0, Note: "first array empty"},
			{Input: map[string]interface{}{"nums1": []int{2}, "nums2": []int{}}, Expected: 2.0, Note: "second array empty (catches not swapping so the shorter one is searched)"},
			{Input: map[string]interface{}{"nums1": []int{}, "nums2": []int{2, 3}}, Expected: 2.5, Hidden: true, Note: "empty array, even total"},
			{Input: map[string]interface{}{"nums1": []int{1, 2}, "nums2": []int{-1, 3}}, Expected: 1.5, Hidden: true, Note: "interleaved values with a negative"},
			{Input: map[string]interface{}{"nums1": []int{1, 1}, "nums2": []int{1, 1}}, Expected: 1.0, Hidden: true, Note: "all equal: the partition test must use <=, not <"},
			{Input: map[string]interface{}{"nums1": []int{1, 2, 3}, "nums2": []int{4, 5, 6, 7, 8}}, Expected: 4.5, Hidden: true, Note: "arrays don't overlap; the partition sits at an end of the shorter array (needs the infinity sentinels)"},
			{Input: map[string]interface{}{"nums1": []int{3}, "nums2": []int{-2, -1}}, Expected: -1.0, Hidden: true, Note: "negative median from the longer array"},
			{Input: map[string]interface{}{"nums1": []int{0, 0}, "nums2": []int{0, 0}}, Expected: 0.0, Hidden: true, Note: "all zeros"},
		},
		TimeComplexity:  "O(log(min(m, n)))",
		SpaceComplexity: "O(1)",
		StarterCode:     "def findMedianSortedArrays(nums1, nums2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The median splits all m + n numbers into a left half and a right half of (almost) equal size. If you decide that the left half takes i numbers from the front of nums1, how many must it take from the front of nums2?"},
			{Level: 2, Type: "insight", Content: "Choose i, the number of elements nums1 contributes to the left half; then j = (m + n + 1) // 2 - i is forced. The split is correct exactly when nums1[i-1] <= nums2[j] and nums2[j-1] <= nums1[i]. If nums1[i-1] is too big, i must shrink; if nums2[j-1] is too big, i must grow. That direction is monotonic, so binary-search i over 0..m on the shorter array (chapter 6, binary search over partitions).", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Define the search space: make nums1 the shorter array; i ranges over 0..m, and half = (m + n + 1) // 2.\n2. Place the partition: i = (lo + hi) // 2 and j = half - i. Read the four border values, using -infinity for an empty left side and +infinity for an empty right side.\n3. Check the partition: valid when left1 <= right2 and left2 <= right1.\n4. Move the partition: if left1 > right2, nums1 gave too many, so hi = i - 1; otherwise it gave too few, so lo = i + 1.\n5. Compute the median: odd total means max(left1, left2); even means the average of that and min(right1, right2)."},
			{Level: 4, Type: "skeleton", Content: `def findMedianSortedArrays(nums1, nums2):
    if len(nums1) > len(nums2):
        nums1, nums2 = nums2, nums1
    m, n = len(nums1), len(nums2)
    half = (m + n + 1) // 2
    lo, hi = 0, m
    while lo <= hi:
        i = (lo + hi) // 2
        j = ____
        left1 = nums1[i - 1] if i > 0 else float("-inf")
        right1 = nums1[i] if i < m else float("inf")
        left2 = nums2[j - 1] if j > 0 else float("-inf")
        right2 = nums2[j] if j < n else float("inf")
        if ____:
            if (m + n) % 2 == 1:
                return float(max(left1, left2))
            return ____
        if left1 > right2:
            hi = ____
        else:
            lo = ____
    return 0.0`},
		},
		Solution: Solution{
			Code: `def findMedianSortedArrays(nums1, nums2):
    if len(nums1) > len(nums2):
        nums1, nums2 = nums2, nums1
    m, n = len(nums1), len(nums2)
    half = (m + n + 1) // 2
    lo, hi = 0, m
    while lo <= hi:
        i = (lo + hi) // 2
        j = half - i
        left1 = nums1[i - 1] if i > 0 else float("-inf")
        right1 = nums1[i] if i < m else float("inf")
        left2 = nums2[j - 1] if j > 0 else float("-inf")
        right2 = nums2[j] if j < n else float("inf")
        if left1 <= right2 and left2 <= right1:
            if (m + n) % 2 == 1:
                return float(max(left1, left2))
            return (max(left1, left2) + min(right1, right2)) / 2
        if left1 > right2:
            hi = i - 1
        else:
            lo = i + 1
    return 0.0`,
			Explanation: "The top of the ladder. Like Koko (875) it searches a space of choices rather than positions of a value: here the choice is how many elements nums1 puts in the left half of the combined list. Like Search in Rotated Sorted Array (33) it uses the closed template and decides direction with a comparison that isn't simply 'mid vs target'. Searching the shorter array keeps it O(log(min(m, n))), and the infinity sentinels handle empty arrays and partitions at the ends.",
			BruteForce:  "Merge the two arrays (the two-pointer merge from chapter 3) or concatenate and sort, then read the middle. Merging is O(m + n) time and space: 2,000 steps at the maximum size. It wastes work because it builds the whole combined list when only the one or two elements at the middle matter.",
			KeyInsight:  "The median is determined by a split, not by the full merge: you need a left part of half = (m + n + 1) // 2 elements in which every element is <= every element of the right part. Taking i elements from nums1 forces j = half - i from nums2, so there's only one number to search for, i, and checking a candidate split takes four comparisons.",
			WhyItWorks:  "For any i in 0..m, the left part is nums1[:i] + nums2[:j] with j = half - i (j stays in 0..n because nums1 is the shorter array). Both arrays are sorted, so the split is correct exactly when the two cross conditions hold: left1 = nums1[i-1] <= right2 = nums2[j] and left2 = nums2[j-1] <= right1 = nums1[i]. If left1 > right2, nums1 put an element that's too big on the left, and so would any larger i, so the answer is below i (hi = i - 1). If left2 > right1, nums2's left side is too big, which only gets worse as i shrinks, so the answer is above i (lo = i + 1). A correct split always exists, so the loop returns. The left part then holds the half smallest numbers: its largest, max(left1, left2), is the median for an odd total, and the even median averages it with the smallest of the right part, min(right1, right2). Missing neighbours become -inf/+inf so they never fail a comparison.",
			CommonMistakes: []string{
				"Binary-searching the longer array: j = half - i can then go negative or exceed n, and indexing breaks (try nums1 = [2], nums2 = [] without the swap). Always swap so that nums1 is the shorter one.",
				"Indexing nums1[i - 1] when i == 0 (or nums1[i] when i == m): Python's nums1[-1] silently reads the LAST element instead of failing, giving wrong answers on inputs like [1,2,3] and [4,5,6,7,8]. Use -infinity and +infinity for missing borders.",
				"Using < instead of <= in the partition test: with equal values such as [1,1] and [1,1], a correct split has left1 == right2 and is rejected, and the search moves past it.",
				"Returning an integer for an odd total or doing integer division // for the even case: (2 + 3) // 2 is 2, not 2.5. Divide with / and return a float.",
			},
			TimeComplexity:  "O(log(min(m, n)))",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the search space", Explanation: "Swap so that nums1 is the shorter array; then i, the count nums1 contributes to the left part, ranges over 0..m, and half is the size of the left part (it gets the extra element when the total is odd).", CodeSnippet: "if len(nums1) > len(nums2):\nnums1, nums2 = nums2, nums1\nhalf = (m + n + 1) // 2\nlo, hi = 0, m", LineStart: 2, LineEnd: 6},
				{Title: "Place the partition", Explanation: "Try the middle i; nums2 must then contribute j = half - i. For nums1 = [1,3,8], nums2 = [7,9,10,11]: half = 4, first i = 1, j = 3.", CodeSnippet: "i = (lo + hi) // 2\nj = half - i", LineStart: 7, LineEnd: 9},
				{Title: "Read the borders", Explanation: "The four numbers next to the cut. A side with no element becomes -infinity (left) or +infinity (right), so it never breaks a comparison. For i = 1, j = 3: left1 = 1, right1 = 3, left2 = 10, right2 = 11.", CodeSnippet: "left1 = nums1[i - 1] if i > 0 else float(\"-inf\")\nright2 = nums2[j] if j < n else float(\"inf\")", LineStart: 10, LineEnd: 13},
				{Title: "Check the partition", Explanation: "Correct when each left border is <= the opposite right border. Then the left part is the half smallest numbers and the median comes from the borders.", CodeSnippet: "if left1 <= right2 and left2 <= right1:\nreturn float(max(left1, left2))\nreturn (max(left1, left2) + min(right1, right2)) / 2", LineStart: 14, LineEnd: 17},
				{Title: "Move the partition", Explanation: "left1 > right2 means nums1 gave too many, so search smaller i. Otherwise left2 > right1 and nums1 gave too few. In the trace, 10 > 3 sends i to 2, then 9 > 8 sends it to 3, where 8 <= 9 and 7 <= inf hold: the median is max(8, 7) = 8.", CodeSnippet: "if left1 > right2:\nhi = i - 1\nlo = i + 1", LineStart: 18, LineEnd: 21},
				{Title: "Read the answer", Explanation: "This line is never reached for valid input, because some i in 0..m always gives a correct split; it's there so the function always returns a number.", CodeSnippet: "return 0.0", LineStart: 22, LineEnd: 22},
			},
			Alternatives: []AltSolution{
				{
					Name: "Combine and sort (brute force)",
					Code: `def findMedianSortedArrays(nums1, nums2):
    combined = sorted(nums1 + nums2)
    total = len(combined)
    middle = total // 2
    if total % 2 == 1:
        return float(combined[middle])
    return (combined[middle - 1] + combined[middle]) / 2`,
					TimeComplexity:  "O((m + n) log(m + n))",
					SpaceComplexity: "O(m + n)",
					Explanation:     "Ignores that the inputs are already sorted. Simplest possible correct answer, useful as a reference.",
				},
				{
					Name: "Two-pointer merge up to the middle",
					Code: `def findMedianSortedArrays(nums1, nums2):
    total = len(nums1) + len(nums2)
    a = b = 0
    previous = current = 0
    for _ in range(total // 2 + 1):
        previous = current
        if b >= len(nums2) or (a < len(nums1) and nums1[a] <= nums2[b]):
            current = nums1[a]
            a += 1
        else:
            current = nums2[b]
            b += 1
    if total % 2 == 1:
        return float(current)
    return (previous + current) / 2`,
					TimeComplexity:  "O(m + n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Merges like in chapter 3 but stops at the middle and remembers only the last two values taken. Linear time and constant space; good enough in practice, but not the O(log) the problem asks for.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the binary search run over the shorter array?", Answer: "j = half - i must be a valid cut of nums2, between 0 and n. If m <= n, then for every i in 0..m, half - i is between half - m >= 0 and half <= n, so j is always valid. If nums1 were longer, some i would give a negative j or one past n. It also makes the search O(log(min(m, n)))."},
			{Prompt: "Why is it safe to treat a missing border as -infinity or +infinity?", Answer: "If i == 0, nums1 puts nothing in the left part, so there is no left1 that could be too big; -infinity makes 'left1 <= right2' automatically true. Likewise when i == m, there is no right1, and +infinity makes 'left2 <= right1' true. The same goes for j. It removes four special cases."},
			{Prompt: "Why does half use (m + n + 1) // 2?", Answer: "It makes the left part the bigger one when the total is odd: for 7 numbers the left part holds 4, so the median is simply the largest left element, max(left1, left2). For an even total it equals (m + n) / 2 and the median averages the largest left and smallest right elements."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, BinarySearchProblems...)
}
