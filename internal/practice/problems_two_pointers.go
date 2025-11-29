package practice

// TwoPointersProblems contains all two-pointers category problems
var TwoPointersProblems = []*Problem{
	{
		ID:         "valid-palindrome",
		Number:     10,
		Title:      "Valid Palindrome",
		Difficulty: "Easy",
		Category:   "two-pointers",
		Tags:       []string{"Two Pointers", "String"},
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
		ID:         "3sum",
		Number:     12,
		Title:      "3Sum",
		Difficulty: "Medium",
		Category:   "two-pointers",
		Tags:       []string{"Array", "Two Pointers", "Sorting"},
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
			{Input: map[string]interface{}{"nums": []int{-1, 0, 1, 2, -1, -4}}, Expected: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0}}, Expected: [][]int{{0, 0, 0}}},
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
		ID:         "container-with-most-water",
		Number:     13,
		Title:      "Container With Most Water",
		Difficulty: "Medium",
		Category:   "two-pointers",
		Tags:       []string{"Array", "Two Pointers", "Greedy"},
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
		ID:         "trapping-rain-water",
		Number:     14,
		Title:      "Trapping Rain Water",
		Difficulty: "Hard",
		Category:   "two-pointers",
		Tags:       []string{"Array", "Two Pointers", "Dynamic Programming", "Stack"},
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

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, TwoPointersProblems...)
}
