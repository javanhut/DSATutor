package practice

// BinarySearchProblems contains all binary-search category problems
var BinarySearchProblems = []*Problem{
	{
		ID:         "binary-search",
		Number:     20,
		Title:      "Binary Search",
		Difficulty: "Easy",
		Category:   "binary-search",
		Tags:       []string{"Array", "Binary Search"},
		RelatedChapters: []int{6},
		Description: `Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.`,
		Constraints: []string{
			"1 <= nums.length <= 10^4",
			"-10^4 < nums[i], target < 10^4",
			"All the integers in nums are unique",
			"nums is sorted in ascending order",
		},
		Examples: []Example{
			{Input: "nums = [-1,0,3,5,9,12], target = 9", Output: "4"},
			{Input: "nums = [-1,0,3,5,9,12], target = 2", Output: "-1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{-1, 0, 3, 5, 9, 12}, "target": 9}, Expected: 4},
			{Input: map[string]interface{}{"nums": []int{-1, 0, 3, 5, 9, 12}, "target": 2}, Expected: -1},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def search(nums, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Compare target with middle element to eliminate half the search space."},
		},
		Solution: Solution{
			Code: `def search(nums, target):
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return -1`,
			Explanation: "Standard binary search. Compare mid element, adjust left or right pointer.",
		},
	},
	{
		ID:              "search-2d-matrix",
		Number:          46,
		Title:           "Search a 2D Matrix",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search", "Matrix"},
		RelatedChapters: []int{1, 6},
		Description: `You are given an m x n integer matrix with the following two properties:

- Each row is sorted in non-decreasing order.
- The first integer of each row is greater than the last integer of the previous row.

Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.`,
		Constraints: []string{
			"m == matrix.length",
			"n == matrix[i].length",
			"1 <= m, n <= 100",
			"-10^4 <= matrix[i][j], target <= 10^4",
		},
		Examples: []Example{
			{Input: "matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3", Output: "true"},
			{Input: "matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, "target": 3}, Expected: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, "target": 13}, Expected: false},
		},
		TimeComplexity:  "O(log(m * n))",
		SpaceComplexity: "O(1)",
		StarterCode:     "def searchMatrix(matrix, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Treat the 2D matrix as a 1D sorted array. Use binary search on indices 0 to m*n-1.", ChapterRef: 1},
			{Level: 2, Type: "algorithm", Content: "Convert 1D index to 2D: row = mid // cols, col = mid % cols. Then apply standard binary search."},
			{Level: 3, Type: "code", Content: "left, right = 0, m*n-1. mid = (left+right)//2. row, col = mid//n, mid%n. Compare matrix[row][col] with target."},
		},
		Solution: Solution{
			Code: `def searchMatrix(matrix, target):
    if not matrix or not matrix[0]:
        return False

    m, n = len(matrix), len(matrix[0])
    left, right = 0, m * n - 1

    while left <= right:
        mid = (left + right) // 2
        row, col = mid // n, mid % n
        val = matrix[row][col]

        if val == target:
            return True
        elif val < target:
            left = mid + 1
        else:
            right = mid - 1

    return False`,
			Explanation:     "Treat 2D matrix as flattened 1D array. Convert mid index to row/col coordinates. Standard binary search.",
			TimeComplexity:  "O(log(m * n))",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Set search range", Explanation: "Search indices 0 to m*n-1", CodeSnippet: "left, right = 0, m * n - 1", LineStart: 6, LineEnd: 6},
				{Title: "Convert to 2D", Explanation: "row = mid // cols, col = mid % cols", CodeSnippet: "row, col = mid // n, mid % n", LineStart: 10, LineEnd: 10},
				{Title: "Binary search", Explanation: "Compare and narrow search space", CodeSnippet: "if val == target:", LineStart: 13, LineEnd: 18},
			},
		},
	},
	{
		ID:              "koko-eating-bananas",
		Number:          47,
		Title:           "Koko Eating Bananas",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search"},
		RelatedChapters: []int{1, 6},
		Description: `Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.`,
		Constraints: []string{
			"1 <= piles.length <= 10^4",
			"piles.length <= h <= 10^9",
			"1 <= piles[i] <= 10^9",
		},
		Examples: []Example{
			{Input: "piles = [3,6,7,11], h = 8", Output: "4"},
			{Input: "piles = [30,11,23,4,20], h = 5", Output: "30"},
			{Input: "piles = [30,11,23,4,20], h = 6", Output: "23"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"piles": []int{3, 6, 7, 11}, "h": 8}, Expected: 4},
			{Input: map[string]interface{}{"piles": []int{30, 11, 23, 4, 20}, "h": 5}, Expected: 30},
			{Input: map[string]interface{}{"piles": []int{30, 11, 23, 4, 20}, "h": 6}, Expected: 23},
		},
		TimeComplexity:  "O(n * log(max(piles)))",
		SpaceComplexity: "O(1)",
		StarterCode:     "def minEatingSpeed(piles, h):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Binary search on the answer. The speed k ranges from 1 to max(piles). Find minimum k where total hours <= h."},
			{Level: 2, Type: "algorithm", Content: "For each speed k, calculate hours needed: sum of ceil(pile/k) for each pile. If hours <= h, try smaller k; else try larger."},
			{Level: 3, Type: "code", Content: "Hours for pile = (pile + k - 1) // k (ceiling division). Binary search: if hours <= h, right = mid; else left = mid + 1."},
		},
		Solution: Solution{
			Code: `def minEatingSpeed(piles, h):
    def hours_needed(k):
        return sum((pile + k - 1) // k for pile in piles)

    left, right = 1, max(piles)

    while left < right:
        mid = (left + right) // 2
        if hours_needed(mid) <= h:
            right = mid
        else:
            left = mid + 1

    return left`,
			Explanation:     "Binary search on eating speed. Find minimum speed where total hours to eat all bananas <= h.",
			TimeComplexity:  "O(n * log(max(piles)))",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Calculate hours", Explanation: "Ceiling division: (pile + k - 1) // k", CodeSnippet: "def hours_needed(k):", LineStart: 2, LineEnd: 3},
				{Title: "Binary search range", Explanation: "Speed from 1 to max pile size", CodeSnippet: "left, right = 1, max(piles)", LineStart: 5, LineEnd: 5},
				{Title: "Find minimum speed", Explanation: "If can finish in time, try slower speed", CodeSnippet: "if hours_needed(mid) <= h:\n    right = mid", LineStart: 9, LineEnd: 10},
			},
		},
	},
	{
		ID:              "find-minimum-in-rotated-sorted-array",
		Number:          48,
		Title:           "Find Minimum in Rotated Sorted Array",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search"},
		RelatedChapters: []int{1, 4, 6},
		Description: `Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

- [4,5,6,7,0,1,2] if it was rotated 4 times.
- [0,1,2,4,5,6,7] if it was rotated 7 times.

Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.`,
		Constraints: []string{
			"n == nums.length",
			"1 <= n <= 5000",
			"-5000 <= nums[i] <= 5000",
			"All the integers of nums are unique",
			"nums is sorted and rotated between 1 and n times",
		},
		Examples: []Example{
			{Input: "nums = [3,4,5,1,2]", Output: "1", Explanation: "Original array was [1,2,3,4,5] rotated 3 times."},
			{Input: "nums = [4,5,6,7,0,1,2]", Output: "0"},
			{Input: "nums = [11,13,15,17]", Output: "11"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{3, 4, 5, 1, 2}}, Expected: 1},
			{Input: map[string]interface{}{"nums": []int{4, 5, 6, 7, 0, 1, 2}}, Expected: 0},
			{Input: map[string]interface{}{"nums": []int{11, 13, 15, 17}}, Expected: 11},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def findMin(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "The minimum is at the rotation pivot. Use binary search comparing mid with right to find the unsorted half.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "If nums[mid] > nums[right], minimum is in right half. Otherwise, it's in left half (including mid)."},
			{Level: 3, Type: "code", Content: "while left < right: if nums[mid] > nums[right]: left = mid + 1; else: right = mid. Return nums[left]."},
		},
		Solution: Solution{
			Code: `def findMin(nums):
    left, right = 0, len(nums) - 1

    while left < right:
        mid = (left + right) // 2
        if nums[mid] > nums[right]:
            # Minimum is in right half
            left = mid + 1
        else:
            # Minimum is in left half (including mid)
            right = mid

    return nums[left]`,
			Explanation:     "Binary search for rotation point. If mid > right, rotation point (minimum) is in right half. Otherwise, it's in left half.",
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Compare mid with right", Explanation: "Determines which half contains the minimum", CodeSnippet: "if nums[mid] > nums[right]:", LineStart: 6, LineEnd: 8},
				{Title: "Narrow to unsorted half", Explanation: "Minimum is in the unsorted portion", CodeSnippet: "left = mid + 1", LineStart: 8, LineEnd: 8},
				{Title: "Return minimum", Explanation: "When left == right, we found the minimum", CodeSnippet: "return nums[left]", LineStart: 13, LineEnd: 13},
			},
		},
	},
	{
		ID:              "search-in-rotated-sorted-array",
		Number:          49,
		Title:           "Search in Rotated Sorted Array",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search"},
		RelatedChapters: []int{1, 4, 6},
		Description: `There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.`,
		Constraints: []string{
			"1 <= nums.length <= 5000",
			"-10^4 <= nums[i] <= 10^4",
			"All values of nums are unique",
			"nums is an ascending array that is possibly rotated",
			"-10^4 <= target <= 10^4",
		},
		Examples: []Example{
			{Input: "nums = [4,5,6,7,0,1,2], target = 0", Output: "4"},
			{Input: "nums = [4,5,6,7,0,1,2], target = 3", Output: "-1"},
			{Input: "nums = [1], target = 0", Output: "-1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{4, 5, 6, 7, 0, 1, 2}, "target": 0}, Expected: 4},
			{Input: map[string]interface{}{"nums": []int{4, 5, 6, 7, 0, 1, 2}, "target": 3}, Expected: -1},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 0}, Expected: -1},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def search(nums, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "One half is always sorted. Determine which half is sorted, then check if target is in that range.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "If left half sorted (nums[left] <= nums[mid]): check if target in [left, mid). If right half sorted: check if target in (mid, right]."},
			{Level: 3, Type: "code", Content: "if nums[left] <= nums[mid]: if nums[left] <= target < nums[mid]: search left; else search right."},
		},
		Solution: Solution{
			Code: `def search(nums, target):
    left, right = 0, len(nums) - 1

    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid

        # Left half is sorted
        if nums[left] <= nums[mid]:
            if nums[left] <= target < nums[mid]:
                right = mid - 1
            else:
                left = mid + 1
        # Right half is sorted
        else:
            if nums[mid] < target <= nums[right]:
                left = mid + 1
            else:
                right = mid - 1

    return -1`,
			Explanation:     "Determine sorted half, check if target is in that range. If yes, search there; otherwise search the other half.",
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Check if found", Explanation: "Return immediately if mid equals target", CodeSnippet: "if nums[mid] == target:", LineStart: 6, LineEnd: 7},
				{Title: "Identify sorted half", Explanation: "Compare left with mid to determine sorted portion", CodeSnippet: "if nums[left] <= nums[mid]:", LineStart: 10, LineEnd: 10},
				{Title: "Search appropriate half", Explanation: "Check if target is in sorted range, search accordingly", CodeSnippet: "if nums[left] <= target < nums[mid]:", LineStart: 11, LineEnd: 14},
			},
		},
	},
	{
		ID:              "time-based-key-value-store",
		Number:          50,
		Title:           "Time Based Key-Value Store",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Hash Table", "String", "Binary Search", "Design"},
		RelatedChapters: []int{1, 5, 6},
		Description: `Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

Implement the TimeMap class:
- TimeMap() Initializes the object of the data structure.
- void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
- String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp. If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".`,
		Constraints: []string{
			"1 <= key.length, value.length <= 100",
			"key and value consist of lowercase English letters and digits",
			"1 <= timestamp <= 10^7",
			"All the timestamps timestamp of set are strictly increasing",
			"At most 2 * 10^5 calls will be made to set and get",
		},
		Examples: []Example{
			{Input: `["TimeMap","set","get","get","set","get","get"]
[[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]`, Output: `[null,null,"bar","bar",null,"bar2","bar2"]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"TimeMap", "set", "get", "get", "set", "get", "get"}, "values": [][]interface{}{{}, {"foo", "bar", 1}, {"foo", 1}, {"foo", 3}, {"foo", "bar2", 4}, {"foo", 4}, {"foo", 5}}}, Expected: []interface{}{nil, nil, "bar", "bar", nil, "bar2", "bar2"}},
		},
		TimeComplexity:  "O(log n) for get, O(1) for set",
		SpaceComplexity: "O(n)",
		StarterCode: `class TimeMap:
    def __init__(self):
        pass

    def set(self, key: str, value: str, timestamp: int) -> None:
        pass

    def get(self, key: str, timestamp: int) -> str:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a dictionary mapping keys to lists of (timestamp, value) pairs. Since timestamps are increasing, lists are sorted."},
			{Level: 2, Type: "algorithm", Content: "For get: binary search for largest timestamp <= query timestamp. Use bisect_right and check index-1."},
			{Level: 3, Type: "code", Content: "from bisect import bisect_right. idx = bisect_right(timestamps, timestamp) - 1. Return values[idx] if idx >= 0 else ''."},
		},
		Solution: Solution{
			Code: `from collections import defaultdict
from bisect import bisect_right

class TimeMap:
    def __init__(self):
        self.store = defaultdict(list)  # key -> [(timestamp, value), ...]

    def set(self, key: str, value: str, timestamp: int) -> None:
        self.store[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        if key not in self.store:
            return ""

        pairs = self.store[key]
        # Binary search for rightmost timestamp <= query
        idx = bisect_right(pairs, (timestamp, chr(127))) - 1

        if idx >= 0:
            return pairs[idx][1]
        return ""`,
			Explanation:     "Store (timestamp, value) pairs per key. Binary search for largest timestamp <= query. Use chr(127) as upper bound for value comparison.",
			TimeComplexity:  "O(1) for set, O(log n) for get",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Store pairs", Explanation: "Append (timestamp, value) to key's list", CodeSnippet: "self.store[key].append((timestamp, value))", LineStart: 9, LineEnd: 9},
				{Title: "Binary search", Explanation: "Find insertion point and go back one", CodeSnippet: "idx = bisect_right(pairs, (timestamp, chr(127))) - 1", LineStart: 17, LineEnd: 17},
				{Title: "Return value", Explanation: "Return value at found index, or empty string", CodeSnippet: "return pairs[idx][1]", LineStart: 20, LineEnd: 20},
			},
		},
	},
	{
		ID:              "median-of-two-sorted-arrays",
		Number:          51,
		Title:           "Median of Two Sorted Arrays",
		Difficulty:      "Hard",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search", "Divide and Conquer"},
		RelatedChapters: []int{1, 4, 6, 12},
		Description: `Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).`,
		Constraints: []string{
			"nums1.length == m",
			"nums2.length == n",
			"0 <= m <= 1000",
			"0 <= n <= 1000",
			"1 <= m + n <= 2000",
			"-10^6 <= nums1[i], nums2[i] <= 10^6",
		},
		Examples: []Example{
			{Input: "nums1 = [1,3], nums2 = [2]", Output: "2.00000", Explanation: "Merged array = [1,2,3] and median is 2."},
			{Input: "nums1 = [1,2], nums2 = [3,4]", Output: "2.50000", Explanation: "Merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums1": []int{1, 3}, "nums2": []int{2}}, Expected: 2.0},
			{Input: map[string]interface{}{"nums1": []int{1, 2}, "nums2": []int{3, 4}}, Expected: 2.5},
		},
		TimeComplexity:  "O(log(min(m, n)))",
		SpaceComplexity: "O(1)",
		StarterCode:     "def findMedianSortedArrays(nums1, nums2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Binary search on the smaller array to find a partition. Left partition should have (m+n+1)//2 elements total.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "Partition arrays so max(left1, left2) <= min(right1, right2). Then median involves these boundary elements."},
			{Level: 3, Type: "code", Content: "Binary search on partition of smaller array. Check if maxLeft1 <= minRight2 and maxLeft2 <= minRight1. Handle edge cases with infinity."},
		},
		Solution: Solution{
			Code: `def findMedianSortedArrays(nums1, nums2):
    # Ensure nums1 is smaller for binary search efficiency
    if len(nums1) > len(nums2):
        nums1, nums2 = nums2, nums1

    m, n = len(nums1), len(nums2)
    left, right = 0, m
    half = (m + n + 1) // 2

    while left <= right:
        i = (left + right) // 2  # Partition in nums1
        j = half - i              # Partition in nums2

        left1 = nums1[i - 1] if i > 0 else float('-inf')
        right1 = nums1[i] if i < m else float('inf')
        left2 = nums2[j - 1] if j > 0 else float('-inf')
        right2 = nums2[j] if j < n else float('inf')

        if left1 <= right2 and left2 <= right1:
            # Found correct partition
            if (m + n) % 2 == 1:
                return max(left1, left2)
            return (max(left1, left2) + min(right1, right2)) / 2
        elif left1 > right2:
            right = i - 1
        else:
            left = i + 1

    return 0`,
			Explanation:     "Binary search on smaller array's partition. Find where max(left elements) <= min(right elements). Median from boundary elements.",
			TimeComplexity:  "O(log(min(m, n)))",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Ensure smaller first", Explanation: "Binary search on smaller array for efficiency", CodeSnippet: "if len(nums1) > len(nums2):", LineStart: 3, LineEnd: 4},
				{Title: "Find partitions", Explanation: "i partitions nums1, j = half - i partitions nums2", CodeSnippet: "i = (left + right) // 2\nj = half - i", LineStart: 11, LineEnd: 12},
				{Title: "Check partition validity", Explanation: "Valid if all left elements <= all right elements", CodeSnippet: "if left1 <= right2 and left2 <= right1:", LineStart: 19, LineEnd: 23},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, BinarySearchProblems...)
}
