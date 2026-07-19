package practice

// TwoPointersProblems contains all two-pointers category problems
var TwoPointersProblems = []*Problem{
	{
		ID:              "valid-palindrome",
		Number:          10,
		Title:           "Valid Palindrome",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Two Pointers", "String"},
		RelatedChapters: []int{4},
		Description: `A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward.

Given a string s, return true if it is a palindrome, or false otherwise.`,
		Constraints: []string{
			"1 <= s.length <= 2 * 10^5",
			"s consists only of printable ASCII characters",
		},
		Examples: []Example{
			{Input: `s = "A man, a plan, a canal: Panama"`, Output: "true", Explanation: `"amanaplanacanalpanama" is a palindrome.`},
			{Input: `s = "race a car"`, Output: "false", Explanation: `"raceacar" is not a palindrome.`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "A man, a plan, a canal: Panama"}, Expected: true},
			{Input: map[string]interface{}{"s": "race a car"}, Expected: false},
			{Input: map[string]interface{}{"s": " "}, Expected: true},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def isPalindrome(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use two pointers from both ends."},
			{Level: 2, Type: "algorithm", Content: "Skip non-alphanumeric chars, compare lowercase versions."},
		},
		Solution: Solution{
			Code: `def isPalindrome(s):
    left, right = 0, len(s) - 1
    while left < right:
        while left < right and not s[left].isalnum():
            left += 1
        while left < right and not s[right].isalnum():
            right -= 1
        if s[left].lower() != s[right].lower():
            return False
        left += 1
        right -= 1
    return True`,
			Explanation: "Two pointers from ends, skip non-alphanumeric, compare lowercase.",
		},
	},
	{
		ID:              "3sum",
		Number:          12,
		Title:           "3Sum",
		Difficulty:      "Medium",
		Category:        "two-pointers",
		Tags:            []string{"Array", "Two Pointers", "Sorting"},
		RelatedChapters: []int{4},
		Description: `Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.`,
		Constraints: []string{
			"3 <= nums.length <= 3000",
			"-10^5 <= nums[i] <= 10^5",
		},
		Examples: []Example{
			{Input: "nums = [-1,0,1,2,-1,-4]", Output: "[[-1,-1,2],[-1,0,1]]"},
			{Input: "nums = [0,1,1]", Output: "[]"},
			{Input: "nums = [0,0,0]", Output: "[[0,0,0]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{-1, 0, 1, 2, -1, -4}}, Expected: [][]int{{-1, -1, 2}, {-1, 0, 1}}, OrderIndependent: true},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0}}, Expected: [][]int{{0, 0, 0}}, OrderIndependent: true},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def threeSum(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Sort the array first, then use two pointers for each fixed element."},
			{Level: 2, Type: "algorithm", Content: "Fix one element, use two pointers on the rest to find pairs that sum to its negative."},
		},
		Solution: Solution{
			Code: `def threeSum(nums):
    nums.sort()
    result = []
    for i in range(len(nums) - 2):
        if i > 0 and nums[i] == nums[i-1]:
            continue
        left, right = i + 1, len(nums) - 1
        while left < right:
            total = nums[i] + nums[left] + nums[right]
            if total < 0:
                left += 1
            elif total > 0:
                right -= 1
            else:
                result.append([nums[i], nums[left], nums[right]])
                while left < right and nums[left] == nums[left+1]:
                    left += 1
                while left < right and nums[right] == nums[right-1]:
                    right -= 1
                left += 1
                right -= 1
    return result`,
			Explanation: "Sort array, fix first element, use two pointers for remaining pair. Skip duplicates.",
		},
	},
	{
		ID:              "container-with-most-water",
		Number:          13,
		Title:           "Container With Most Water",
		Difficulty:      "Medium",
		Category:        "two-pointers",
		Tags:            []string{"Array", "Two Pointers", "Greedy"},
		RelatedChapters: []int{4},
		Description: `You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.`,
		Constraints: []string{
			"n == height.length",
			"2 <= n <= 10^5",
			"0 <= height[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "height = [1,8,6,2,5,4,8,3,7]", Output: "49"},
			{Input: "height = [1,1]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"height": []int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, Expected: 49},
			{Input: map[string]interface{}{"height": []int{1, 1}}, Expected: 1},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def maxArea(height):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Start with widest container (first and last lines)."},
			{Level: 2, Type: "algorithm", Content: "Move the pointer pointing to the shorter line inward."},
		},
		Solution: Solution{
			Code: `def maxArea(height):
    left, right = 0, len(height) - 1
    max_area = 0
    while left < right:
        area = min(height[left], height[right]) * (right - left)
        max_area = max(max_area, area)
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1
    return max_area`,
			Explanation: "Two pointers from ends. Area = min(heights) * width. Move shorter pointer inward.",
		},
	},
	{
		ID:         "two-sum-ii",
		Number:     11,
		Title:      "Two Sum II - Input Array Is Sorted",
		Difficulty: "Medium",
		Category:   "two-pointers",
		Tags:       []string{"Array", "Two Pointers", "Binary Search"},
		Description: `Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.`,
		Constraints: []string{
			"2 <= numbers.length <= 3 * 10^4",
			"-1000 <= numbers[i] <= 1000",
			"numbers is sorted in non-decreasing order",
			"-1000 <= target <= 1000",
			"The tests are generated such that there is exactly one solution",
		},
		Examples: []Example{
			{Input: "numbers = [2,7,11,15], target = 9", Output: "[1,2]"},
			{Input: "numbers = [2,3,4], target = 6", Output: "[1,3]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"numbers": []int{2, 7, 11, 15}, "target": 9}, Expected: []int{1, 2}},
			{Input: map[string]interface{}{"numbers": []int{2, 3, 4}, "target": 6}, Expected: []int{1, 3}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def twoSum(numbers, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use two pointers since array is sorted."},
			{Level: 2, Type: "algorithm", Content: "If sum too small, move left pointer right. If too big, move right pointer left."},
		},
		Solution: Solution{
			Code: `def twoSum(numbers, target):
    left, right = 0, len(numbers) - 1
    while left < right:
        total = numbers[left] + numbers[right]
        if total == target:
            return [left + 1, right + 1]
        elif total < target:
            left += 1
        else:
            right -= 1
    return []`,
			Explanation: "Two pointers from ends. Adjust based on sum comparison with target.",
		},
	},
	{
		ID:          "trapping-rain-water",
		Number:      14,
		Title:       "Trapping Rain Water",
		Difficulty:  "Hard",
		Category:    "two-pointers",
		Tags:        []string{"Array", "Two Pointers", "Dynamic Programming", "Stack"},
		Description: `Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.`,
		Constraints: []string{
			"n == height.length",
			"1 <= n <= 2 * 10^4",
			"0 <= height[i] <= 10^5",
		},
		Examples: []Example{
			{Input: "height = [0,1,0,2,1,0,1,3,2,1,2,1]", Output: "6"},
			{Input: "height = [4,2,0,3,2,5]", Output: "9"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"height": []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, Expected: 6},
			{Input: map[string]interface{}{"height": []int{4, 2, 0, 3, 2, 5}}, Expected: 9},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def trap(height):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Water at each position depends on min of max heights on left and right."},
			{Level: 2, Type: "algorithm", Content: "Use two pointers tracking max heights from each side."},
		},
		Solution: Solution{
			Code: `def trap(height):
    if not height:
        return 0

    left, right = 0, len(height) - 1
    left_max, right_max = height[left], height[right]
    water = 0

    while left < right:
        if left_max < right_max:
            left += 1
            left_max = max(left_max, height[left])
            water += left_max - height[left]
        else:
            right -= 1
            right_max = max(right_max, height[right])
            water += right_max - height[right]

    return water`,
			Explanation: "Two pointers with max heights. Water at position = min(left_max, right_max) - height.",
		},
	},
	{
		ID:              "sort-colors",
		Number:          159,
		Title:           "Sort Colors",
		Difficulty:      "Medium",
		Category:        "two-pointers",
		Tags:            []string{"Array", "Two Pointers", "Sorting"},
		RelatedChapters: []int{2, 4, 12},
		Description: `Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function. Can you do it in one pass with O(1) extra space?`,
		Constraints: []string{
			"n == nums.length",
			"1 <= n <= 300",
			"nums[i] is either 0, 1, or 2",
		},
		Examples: []Example{
			{Input: "nums = [2,0,2,1,1,0]", Output: "[0,0,1,1,2,2]"},
			{Input: "nums = [2,0,1]", Output: "[0,1,2]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 0, 2, 1, 1, 0}}, Expected: []int{0, 0, 1, 1, 2, 2}},
			{Input: map[string]interface{}{"nums": []int{2, 0, 1}}, Expected: []int{0, 1, 2}},
			{Input: map[string]interface{}{"nums": []int{0}}, Expected: []int{0}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def sortColors(nums):\n    # Modify nums in-place, do not return anything\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "This is the Dutch National Flag problem. Counting sort works in two passes; three pointers do it in one.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "Keep low, mid, high pointers. Everything before low is 0, everything after high is 2, and mid scans the unknown middle: swap 0s down to low, swap 2s up to high, and step over 1s."},
			{Level: 3, Type: "code", Content: "while mid <= high: if nums[mid] == 0: swap(low, mid); low += 1; mid += 1. elif nums[mid] == 2: swap(mid, high); high -= 1 (do not advance mid). else: mid += 1."},
		},
		Solution: Solution{
			Code: `def sortColors(nums):
    low, mid, high = 0, 0, len(nums) - 1
    while mid <= high:
        if nums[mid] == 0:
            nums[low], nums[mid] = nums[mid], nums[low]
            low += 1
            mid += 1
        elif nums[mid] == 2:
            nums[mid], nums[high] = nums[high], nums[mid]
            high -= 1
        else:
            mid += 1`,
			Explanation:     "Dutch National Flag partitioning: three pointers maintain regions of 0s, 1s, unknowns, and 2s, shrinking the unknown region to nothing in one pass.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Zero goes low", Explanation: "Swap into the 0-region and grow it; the swapped-back value is always a known 1", CodeSnippet: "nums[low], nums[mid] = nums[mid], nums[low]"},
				{Title: "Two goes high", Explanation: "Swap into the 2-region but re-examine what came back — it is still unknown", CodeSnippet: "nums[mid], nums[high] = nums[high], nums[mid]"},
			},
		},
	},
	{
		ID:              "move-zeroes",
		Number:          160,
		Title:           "Move Zeroes",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Array", "Two Pointers"},
		RelatedChapters: []int{2, 12},
		Description: `Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.`,
		Constraints: []string{
			"1 <= nums.length <= 10^4",
			"-2^31 <= nums[i] <= 2^31 - 1",
		},
		Examples: []Example{
			{Input: "nums = [0,1,0,3,12]", Output: "[1,3,12,0,0]"},
			{Input: "nums = [0]", Output: "[0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{0, 1, 0, 3, 12}}, Expected: []int{1, 3, 12, 0, 0}},
			{Input: map[string]interface{}{"nums": []int{0}}, Expected: []int{0}},
			{Input: map[string]interface{}{"nums": []int{1, 0, 1}}, Expected: []int{1, 1, 0}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def moveZeroes(nums):\n    # Modify nums in-place, do not return anything\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Think of it as compacting the non-zero elements to the front; the zeros end up at the back automatically.", ChapterRef: 2},
			{Level: 2, Type: "algorithm", Content: "A slow pointer marks where the next non-zero belongs. A fast pointer scans; on each non-zero, swap it into the slow slot and advance slow."},
			{Level: 3, Type: "code", Content: "slow = 0; for fast in range(len(nums)): if nums[fast] != 0: nums[slow], nums[fast] = nums[fast], nums[slow]; slow += 1."},
		},
		Solution: Solution{
			Code: `def moveZeroes(nums):
    slow = 0
    for fast in range(len(nums)):
        if nums[fast] != 0:
            nums[slow], nums[fast] = nums[fast], nums[slow]
            slow += 1`,
			Explanation:     "Slow/fast pointer compaction: every non-zero is swapped into the next open slot at the front, preserving order; zeros drift to the back.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Scan with fast", Explanation: "fast visits every element once", CodeSnippet: "for fast in range(len(nums)):"},
				{Title: "Compact with slow", Explanation: "slow always points at the leftmost zero (or the next free slot)", CodeSnippet: "nums[slow], nums[fast] = nums[fast], nums[slow]"},
			},
		},
	},
	{
		ID:              "merge-sorted-array",
		Number:          161,
		Title:           "Merge Sorted Array",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Array", "Two Pointers", "Sorting"},
		RelatedChapters: []int{2, 12},
		Description: `You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored.`,
		Constraints: []string{
			"nums1.length == m + n",
			"nums2.length == n",
			"0 <= m, n <= 200",
			"-10^9 <= nums1[i], nums2[j] <= 10^9",
		},
		Examples: []Example{
			{Input: "nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3", Output: "[1,2,2,3,5,6]"},
			{Input: "nums1 = [1], m = 1, nums2 = [], n = 0", Output: "[1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums1": []int{1, 2, 3, 0, 0, 0}, "m": 3, "nums2": []int{2, 5, 6}, "n": 3}, Expected: []int{1, 2, 2, 3, 5, 6}},
			{Input: map[string]interface{}{"nums1": []int{1}, "m": 1, "nums2": []int{}, "n": 0}, Expected: []int{1}},
			{Input: map[string]interface{}{"nums1": []int{0}, "m": 0, "nums2": []int{1}, "n": 1}, Expected: []int{1}},
		},
		TimeComplexity:  "O(m + n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def merge(nums1, m, nums2, n):\n    # Modify nums1 in-place, do not return anything\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Merging from the front would overwrite nums1 values you still need. Where in nums1 is there guaranteed free space?", ChapterRef: 2},
			{Level: 2, Type: "algorithm", Content: "Merge backwards: compare the largest remaining element of each array and write it to the end of nums1, walking all three pointers down."},
			{Level: 3, Type: "code", Content: "i, j, w = m - 1, n - 1, m + n - 1. While j >= 0: if i >= 0 and nums1[i] > nums2[j]: nums1[w] = nums1[i]; i -= 1 else: nums1[w] = nums2[j]; j -= 1; w -= 1."},
		},
		Solution: Solution{
			Code: `def merge(nums1, m, nums2, n):
    i, j, w = m - 1, n - 1, m + n - 1
    while j >= 0:
        if i >= 0 and nums1[i] > nums2[j]:
            nums1[w] = nums1[i]
            i -= 1
        else:
            nums1[w] = nums2[j]
            j -= 1
        w -= 1`,
			Explanation:     "Merging from the back uses nums1's spare tail as the write area, so no element is overwritten before it is read. Once nums2 is exhausted, the rest of nums1 is already in place.",
			TimeComplexity:  "O(m + n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Three pointers at the ends", Explanation: "Read pointers at each array's last real element, write pointer at nums1's very end", CodeSnippet: "i, j, w = m - 1, n - 1, m + n - 1"},
				{Title: "Write the larger", Explanation: "The larger of the two candidates belongs at the current write slot", CodeSnippet: "if i >= 0 and nums1[i] > nums2[j]:"},
				{Title: "Stop when nums2 empties", Explanation: "Leftover nums1 elements are already sorted in place", CodeSnippet: "while j >= 0:"},
			},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, TwoPointersProblems...)
}
