package practice

// EmbeddedProblems contains the core Blind 75 problems defined in pure Go.
// Additional problems can be added via JSON files.
var EmbeddedProblems = []*Problem{
	// Arrays & Hashing
	{
		ID:         "two-sum",
		Number:     1,
		Title:      "Two Sum",
		Difficulty: "Easy",
		Category:   "arrays-hashing",
		Tags:       []string{"Array", "Hash Table"},
		RelatedChapters: []int{2, 3},
		Description: `Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.`,
		Constraints: []string{
			"2 <= nums.length <= 10^4",
			"-10^9 <= nums[i] <= 10^9",
			"-10^9 <= target <= 10^9",
			"Only one valid answer exists",
		},
		Examples: []Example{
			{Input: "nums = [2,7,11,15], target = 9", Output: "[0,1]", Explanation: "Because nums[0] + nums[1] == 9, we return [0, 1]."},
			{Input: "nums = [3,2,4], target = 6", Output: "[1,2]"},
			{Input: "nums = [3,3], target = 6", Output: "[0,1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 7, 11, 15}, "target": 9}, Expected: []int{0, 1}},
			{Input: map[string]interface{}{"nums": []int{3, 2, 4}, "target": 6}, Expected: []int{1, 2}},
			{Input: map[string]interface{}{"nums": []int{3, 3}, "target": 6}, Expected: []int{0, 1}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def twoSum(nums, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "A brute force approach would check every pair, but that's O(n^2). Can you do better with a hash table?"},
			{Level: 2, Type: "algorithm", Content: "For each number, check if (target - number) exists in a hash map. If not, store the current number and its index."},
			{Level: 3, Type: "code", Content: "Use a dictionary to store {value: index}. For each num, check if target-num is in the dict."},
		},
		Solution: Solution{
			Code: `def twoSum(nums, target):
    seen = {}
    for i, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            return [seen[complement], i]
        seen[num] = i
    return []`,
			Explanation: "Use a hash map to store each number's index. For each number, check if its complement exists.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize hash map", Explanation: "Create empty dict to store seen numbers", CodeSnippet: "seen = {}"},
				{Title: "Iterate through array", Explanation: "For each number, calculate complement and check if seen", CodeSnippet: "complement = target - num"},
				{Title: "Return result", Explanation: "If complement found, return both indices", CodeSnippet: "return [seen[complement], i]"},
			},
		},
	},
	{
		ID:         "contains-duplicate",
		Number:     2,
		Title:      "Contains Duplicate",
		Difficulty: "Easy",
		Category:   "arrays-hashing",
		Tags:       []string{"Array", "Hash Table", "Sorting"},
		RelatedChapters: []int{2, 3},
		Description: `Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.`,
		Constraints: []string{
			"1 <= nums.length <= 10^5",
			"-10^9 <= nums[i] <= 10^9",
		},
		Examples: []Example{
			{Input: "nums = [1,2,3,1]", Output: "true"},
			{Input: "nums = [1,2,3,4]", Output: "false"},
			{Input: "nums = [1,1,1,3,3,4,3,2,4,2]", Output: "true"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 1}}, Expected: true},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}}, Expected: false},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, Expected: true},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def containsDuplicate(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a set to track seen numbers."},
			{Level: 2, Type: "algorithm", Content: "If a number is already in the set, return true. Otherwise add it."},
		},
		Solution: Solution{
			Code: `def containsDuplicate(nums):
    seen = set()
    for num in nums:
        if num in seen:
            return True
        seen.add(num)
    return False`,
			Explanation: "Use a set to track numbers we've seen. Return true if we see a duplicate.",
		},
	},
	{
		ID:         "valid-anagram",
		Number:     3,
		Title:      "Valid Anagram",
		Difficulty: "Easy",
		Category:   "arrays-hashing",
		Tags:       []string{"Hash Table", "String", "Sorting"},
		RelatedChapters: []int{2, 3},
		Description: `Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.`,
		Constraints: []string{
			"1 <= s.length, t.length <= 5 * 10^4",
			"s and t consist of lowercase English letters",
		},
		Examples: []Example{
			{Input: `s = "anagram", t = "nagaram"`, Output: "true"},
			{Input: `s = "rat", t = "car"`, Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "anagram", "t": "nagaram"}, Expected: true},
			{Input: map[string]interface{}{"s": "rat", "t": "car"}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def isAnagram(s, t):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Count character frequencies in both strings."},
			{Level: 2, Type: "algorithm", Content: "Use a hash map or array of size 26 to count characters."},
		},
		Solution: Solution{
			Code: `def isAnagram(s, t):
    if len(s) != len(t):
        return False
    count = {}
    for c in s:
        count[c] = count.get(c, 0) + 1
    for c in t:
        count[c] = count.get(c, 0) - 1
        if count[c] < 0:
            return False
    return True`,
			Explanation: "Count characters in s, then decrement for t. If any count goes negative, not an anagram.",
		},
	},
	{
		ID:         "group-anagrams",
		Number:     4,
		Title:      "Group Anagrams",
		Difficulty: "Medium",
		Category:   "arrays-hashing",
		Tags:       []string{"Array", "Hash Table", "String", "Sorting"},
		RelatedChapters: []int{2, 3},
		Description: `Given an array of strings strs, group the anagrams together. You can return the answer in any order.`,
		Constraints: []string{
			"1 <= strs.length <= 10^4",
			"0 <= strs[i].length <= 100",
			"strs[i] consists of lowercase English letters",
		},
		Examples: []Example{
			{Input: `strs = ["eat","tea","tan","ate","nat","bat"]`, Output: `[["bat"],["nat","tan"],["ate","eat","tea"]]`},
			{Input: `strs = [""]`, Output: `[[""]]`},
			{Input: `strs = ["a"]`, Output: `[["a"]]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"strs": []string{"eat", "tea", "tan", "ate", "nat", "bat"}}, Expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		},
		TimeComplexity:  "O(n * k log k)",
		SpaceComplexity: "O(n * k)",
		StarterCode:     "def groupAnagrams(strs):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Anagrams have the same sorted characters."},
			{Level: 2, Type: "algorithm", Content: "Use sorted string as key in a hash map."},
		},
		Solution: Solution{
			Code: `from collections import defaultdict

def groupAnagrams(strs):
    groups = defaultdict(list)
    for s in strs:
        key = ''.join(sorted(s))
        groups[key].append(s)
    return list(groups.values())`,
			Explanation: "Sort each string to create a key. Group strings with the same key.",
		},
	},
	{
		ID:         "top-k-frequent-elements",
		Number:     5,
		Title:      "Top K Frequent Elements",
		Difficulty: "Medium",
		Category:   "arrays-hashing",
		Tags:       []string{"Array", "Hash Table", "Heap", "Bucket Sort"},
		RelatedChapters: []int{2, 3, 8},
		Description: `Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.`,
		Constraints: []string{
			"1 <= nums.length <= 10^5",
			"-10^4 <= nums[i] <= 10^4",
			"k is in the range [1, the number of unique elements in the array]",
			"It is guaranteed that the answer is unique",
		},
		Examples: []Example{
			{Input: "nums = [1,1,1,2,2,3], k = 2", Output: "[1,2]"},
			{Input: "nums = [1], k = 1", Output: "[1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 2, 2, 3}, "k": 2}, Expected: []int{1, 2}},
			{Input: map[string]interface{}{"nums": []int{1}, "k": 1}, Expected: []int{1}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def topKFrequent(nums, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Count frequencies, then find top k."},
			{Level: 2, Type: "algorithm", Content: "Use bucket sort: index = frequency, value = list of numbers with that frequency."},
		},
		Solution: Solution{
			Code: `def topKFrequent(nums, k):
    count = {}
    for num in nums:
        count[num] = count.get(num, 0) + 1

    # Bucket sort
    buckets = [[] for _ in range(len(nums) + 1)]
    for num, freq in count.items():
        buckets[freq].append(num)

    result = []
    for i in range(len(buckets) - 1, -1, -1):
        for num in buckets[i]:
            result.append(num)
            if len(result) == k:
                return result
    return result`,
			Explanation: "Count frequencies, use bucket sort with frequency as index, then collect from highest buckets.",
		},
	},
	// Two Pointers
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
	// Sliding Window
	{
		ID:         "best-time-to-buy-and-sell-stock",
		Number:     15,
		Title:      "Best Time to Buy and Sell Stock",
		Difficulty: "Easy",
		Category:   "sliding-window",
		Tags:       []string{"Array", "Dynamic Programming"},
		RelatedChapters: []int{5},
		Description: `You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.`,
		Constraints: []string{
			"1 <= prices.length <= 10^5",
			"0 <= prices[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "prices = [7,1,5,3,6,4]", Output: "5", Explanation: "Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5."},
			{Input: "prices = [7,6,4,3,1]", Output: "0", Explanation: "No profit possible."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"prices": []int{7, 1, 5, 3, 6, 4}}, Expected: 5},
			{Input: map[string]interface{}{"prices": []int{7, 6, 4, 3, 1}}, Expected: 0},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def maxProfit(prices):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Track the minimum price seen so far."},
			{Level: 2, Type: "algorithm", Content: "For each price, calculate profit if selling today, update max profit."},
		},
		Solution: Solution{
			Code: `def maxProfit(prices):
    min_price = float('inf')
    max_profit = 0
    for price in prices:
        min_price = min(min_price, price)
        max_profit = max(max_profit, price - min_price)
    return max_profit`,
			Explanation: "Track minimum price seen. At each step, calculate potential profit and update max.",
		},
	},
	{
		ID:         "longest-substring-without-repeating",
		Number:     16,
		Title:      "Longest Substring Without Repeating Characters",
		Difficulty: "Medium",
		Category:   "sliding-window",
		Tags:       []string{"Hash Table", "String", "Sliding Window"},
		RelatedChapters: []int{5},
		Description: `Given a string s, find the length of the longest substring without repeating characters.`,
		Constraints: []string{
			"0 <= s.length <= 5 * 10^4",
			"s consists of English letters, digits, symbols and spaces",
		},
		Examples: []Example{
			{Input: `s = "abcabcbb"`, Output: "3", Explanation: `The answer is "abc", with the length of 3.`},
			{Input: `s = "bbbbb"`, Output: "1", Explanation: `The answer is "b", with the length of 1.`},
			{Input: `s = "pwwkew"`, Output: "3", Explanation: `The answer is "wke", with the length of 3.`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "abcabcbb"}, Expected: 3},
			{Input: map[string]interface{}{"s": "bbbbb"}, Expected: 1},
			{Input: map[string]interface{}{"s": "pwwkew"}, Expected: 3},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(min(n, m))",
		StarterCode:     "def lengthOfLongestSubstring(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use sliding window with a set to track characters in current window."},
			{Level: 2, Type: "algorithm", Content: "Expand right, shrink left when duplicate found."},
		},
		Solution: Solution{
			Code: `def lengthOfLongestSubstring(s):
    char_set = set()
    left = 0
    max_len = 0
    for right in range(len(s)):
        while s[right] in char_set:
            char_set.remove(s[left])
            left += 1
        char_set.add(s[right])
        max_len = max(max_len, right - left + 1)
    return max_len`,
			Explanation: "Sliding window with set. Remove from left until no duplicate, then add right char.",
		},
	},
	// Binary Search
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
	// Linked List
	{
		ID:         "reverse-linked-list",
		Number:     23,
		Title:      "Reverse Linked List",
		Difficulty: "Easy",
		Category:   "linked-list",
		Tags:       []string{"Linked List", "Recursion"},
		RelatedChapters: []int{7},
		Description: `Given the head of a singly linked list, reverse the list, and return the reversed list.`,
		Constraints: []string{
			"The number of nodes in the list is the range [0, 5000]",
			"-5000 <= Node.val <= 5000",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4,5]", Output: "[5,4,3,2,1]"},
			{Input: "head = [1,2]", Output: "[2,1]"},
			{Input: "head = []", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}}, Expected: []int{5, 4, 3, 2, 1}},
			{Input: map[string]interface{}{"head": []int{1, 2}}, Expected: []int{2, 1}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def reverseList(head):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use three pointers: prev, curr, next."},
			{Level: 2, Type: "algorithm", Content: "At each step, reverse the link and move all pointers forward."},
		},
		Solution: Solution{
			Code: `def reverseList(head):
    prev = None
    curr = head
    while curr:
        next_temp = curr.next
        curr.next = prev
        prev = curr
        curr = next_temp
    return prev`,
			Explanation: "Iteratively reverse each link. prev tracks new head.",
		},
	},
	{
		ID:         "linked-list-cycle",
		Number:     25,
		Title:      "Linked List Cycle",
		Difficulty: "Easy",
		Category:   "linked-list",
		Tags:       []string{"Hash Table", "Linked List", "Two Pointers"},
		RelatedChapters: []int{7},
		Description: `Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.

Return true if there is a cycle in the linked list. Otherwise, return false.`,
		Constraints: []string{
			"The number of the nodes in the list is in the range [0, 10^4]",
			"-10^5 <= Node.val <= 10^5",
		},
		Examples: []Example{
			{Input: "head = [3,2,0,-4], pos = 1", Output: "true", Explanation: "Tail connects to node index 1."},
			{Input: "head = [1,2], pos = 0", Output: "true"},
			{Input: "head = [1], pos = -1", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{3, 2, 0, -4}, "pos": 1}, Expected: true},
			{Input: map[string]interface{}{"head": []int{1}, "pos": -1}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def hasCycle(head):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use Floyd's cycle detection (slow and fast pointers)."},
			{Level: 2, Type: "algorithm", Content: "If fast catches up to slow, there's a cycle."},
		},
		Solution: Solution{
			Code: `def hasCycle(head):
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            return True
    return False`,
			Explanation: "Floyd's algorithm: slow moves 1 step, fast moves 2 steps. They meet if cycle exists.",
		},
	},
	// Trees
	{
		ID:         "invert-binary-tree",
		Number:     26,
		Title:      "Invert Binary Tree",
		Difficulty: "Easy",
		Category:   "trees",
		Tags:       []string{"Tree", "Binary Tree", "DFS", "BFS"},
		RelatedChapters: []int{9, 10},
		Description: `Given the root of a binary tree, invert the tree, and return its root.`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [0, 100]",
			"-100 <= Node.val <= 100",
		},
		Examples: []Example{
			{Input: "root = [4,2,7,1,3,6,9]", Output: "[4,7,2,9,6,3,1]"},
			{Input: "root = [2,1,3]", Output: "[2,3,1]"},
			{Input: "root = []", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{4, 2, 7, 1, 3, 6, 9}}, Expected: []interface{}{4, 7, 2, 9, 6, 3, 1}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def invertTree(root):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Recursively swap left and right children."},
		},
		Solution: Solution{
			Code: `def invertTree(root):
    if not root:
        return None
    root.left, root.right = root.right, root.left
    invertTree(root.left)
    invertTree(root.right)
    return root`,
			Explanation: "Swap left and right children, then recursively invert subtrees.",
		},
	},
	{
		ID:         "maximum-depth-of-binary-tree",
		Number:     27,
		Title:      "Maximum Depth of Binary Tree",
		Difficulty: "Easy",
		Category:   "trees",
		Tags:       []string{"Tree", "DFS", "BFS", "Binary Tree"},
		RelatedChapters: []int{9, 10},
		Description: `Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [0, 10^4]",
			"-100 <= Node.val <= 100",
		},
		Examples: []Example{
			{Input: "root = [3,9,20,null,null,15,7]", Output: "3"},
			{Input: "root = [1,null,2]", Output: "2"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 9, 20, nil, nil, 15, 7}}, Expected: 3},
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2}}, Expected: 2},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def maxDepth(root):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Recursively find max depth of left and right subtrees."},
		},
		Solution: Solution{
			Code: `def maxDepth(root):
    if not root:
        return 0
    return 1 + max(maxDepth(root.left), maxDepth(root.right))`,
			Explanation: "Base case: empty tree has depth 0. Otherwise, 1 + max of subtree depths.",
		},
	},
	// Graphs
	{
		ID:         "number-of-islands",
		Number:     32,
		Title:      "Number of Islands",
		Difficulty: "Medium",
		Category:   "graphs",
		Tags:       []string{"Graph", "BFS", "DFS", "Matrix"},
		RelatedChapters: []int{11},
		Description: `Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.`,
		Constraints: []string{
			"m == grid.length",
			"n == grid[i].length",
			"1 <= m, n <= 300",
			"grid[i][j] is '0' or '1'",
		},
		Examples: []Example{
			{Input: `grid = [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]`, Output: "1"},
			{Input: `grid = [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]`, Output: "3"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"grid": [][]string{{"1", "1", "0"}, {"0", "1", "0"}, {"0", "0", "1"}}}, Expected: 2},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def numIslands(grid):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use DFS/BFS to explore and mark each island."},
			{Level: 2, Type: "algorithm", Content: "For each '1' found, increment count and flood-fill to mark entire island."},
		},
		Solution: Solution{
			Code: `def numIslands(grid):
    if not grid:
        return 0

    rows, cols = len(grid), len(grid[0])
    count = 0

    def dfs(r, c):
        if r < 0 or r >= rows or c < 0 or c >= cols or grid[r][c] != '1':
            return
        grid[r][c] = '0'  # Mark as visited
        dfs(r+1, c)
        dfs(r-1, c)
        dfs(r, c+1)
        dfs(r, c-1)

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '1':
                count += 1
                dfs(r, c)

    return count`,
			Explanation: "DFS from each unvisited land cell, marking visited cells. Count DFS starts.",
		},
	},
	{
		ID:         "course-schedule",
		Number:     35,
		Title:      "Course Schedule",
		Difficulty: "Medium",
		Category:   "graphs",
		Tags:       []string{"DFS", "BFS", "Graph", "Topological Sort"},
		RelatedChapters: []int{11},
		Description: `There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

Return true if you can finish all courses. Otherwise, return false.`,
		Constraints: []string{
			"1 <= numCourses <= 2000",
			"0 <= prerequisites.length <= 5000",
			"prerequisites[i].length == 2",
			"0 <= ai, bi < numCourses",
			"All the pairs prerequisites[i] are unique",
		},
		Examples: []Example{
			{Input: "numCourses = 2, prerequisites = [[1,0]]", Output: "true", Explanation: "Take course 0, then course 1."},
			{Input: "numCourses = 2, prerequisites = [[1,0],[0,1]]", Output: "false", Explanation: "Circular dependency."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"numCourses": 2, "prerequisites": [][]int{{1, 0}}}, Expected: true},
			{Input: map[string]interface{}{"numCourses": 2, "prerequisites": [][]int{{1, 0}, {0, 1}}}, Expected: false},
		},
		TimeComplexity:  "O(V + E)",
		SpaceComplexity: "O(V + E)",
		StarterCode:     "def canFinish(numCourses, prerequisites):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "This is cycle detection in a directed graph."},
			{Level: 2, Type: "algorithm", Content: "Use DFS with three states: unvisited, visiting, visited. Cycle if we hit 'visiting'."},
		},
		Solution: Solution{
			Code: `def canFinish(numCourses, prerequisites):
    graph = [[] for _ in range(numCourses)]
    for course, prereq in prerequisites:
        graph[course].append(prereq)

    # 0 = unvisited, 1 = visiting, 2 = visited
    state = [0] * numCourses

    def has_cycle(course):
        if state[course] == 1:
            return True
        if state[course] == 2:
            return False

        state[course] = 1
        for prereq in graph[course]:
            if has_cycle(prereq):
                return True
        state[course] = 2
        return False

    for course in range(numCourses):
        if has_cycle(course):
            return False
    return True`,
			Explanation: "DFS cycle detection. Three states track if node is currently being processed.",
		},
	},
	// Dynamic Programming
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
	// Arrays & Hashing (continued)
	{
		ID:         "product-of-array-except-self",
		Number:     6,
		Title:      "Product of Array Except Self",
		Difficulty: "Medium",
		Category:   "arrays-hashing",
		Tags:       []string{"Array", "Prefix Sum"},
		Description: `Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.`,
		Constraints: []string{
			"2 <= nums.length <= 10^5",
			"-30 <= nums[i] <= 30",
			"The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer",
		},
		Examples: []Example{
			{Input: "nums = [1,2,3,4]", Output: "[24,12,8,6]"},
			{Input: "nums = [-1,1,0,-3,3]", Output: "[0,0,9,0,0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}}, Expected: []int{24, 12, 8, 6}},
			{Input: map[string]interface{}{"nums": []int{-1, 1, 0, -3, 3}}, Expected: []int{0, 0, 9, 0, 0}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def productExceptSelf(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Think about prefix and suffix products."},
			{Level: 2, Type: "algorithm", Content: "For each position, multiply prefix product (left) by suffix product (right)."},
		},
		Solution: Solution{
			Code: `def productExceptSelf(nums):
    n = len(nums)
    result = [1] * n

    # Left pass - prefix products
    prefix = 1
    for i in range(n):
        result[i] = prefix
        prefix *= nums[i]

    # Right pass - suffix products
    suffix = 1
    for i in range(n - 1, -1, -1):
        result[i] *= suffix
        suffix *= nums[i]

    return result`,
			Explanation: "Two passes: first stores prefix products, second multiplies by suffix products.",
		},
	},
	{
		ID:         "valid-sudoku",
		Number:     7,
		Title:      "Valid Sudoku",
		Difficulty: "Medium",
		Category:   "arrays-hashing",
		Tags:       []string{"Array", "Hash Table", "Matrix"},
		Description: `Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

1. Each row must contain the digits 1-9 without repetition.
2. Each column must contain the digits 1-9 without repetition.
3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.`,
		Constraints: []string{
			"board.length == 9",
			"board[i].length == 9",
			"board[i][j] is a digit 1-9 or '.'",
		},
		Examples: []Example{
			{Input: "board = [[\"5\",\"3\",\".\",\".\",\"7\",\".\",\".\",\".\",\".\"],[\"6\",\".\",\".\",\"1\",\"9\",\"5\",\".\",\".\",\".\"],[\".\",\"9\",\"8\",\".\",\".\",\".\",\".\",\"6\",\".\"],[\"8\",\".\",\".\",\".\",\"6\",\".\",\".\",\".\",\"3\"],[\"4\",\".\",\".\",\"8\",\".\",\"3\",\".\",\".\",\"1\"],[\"7\",\".\",\".\",\".\",\"2\",\".\",\".\",\".\",\"6\"],[\".\",\"6\",\".\",\".\",\".\",\".\",\"2\",\"8\",\".\"],[\".\",\".\",\".\",\"4\",\"1\",\"9\",\".\",\".\",\"5\"],[\".\",\".\",\".\",\".\",\"8\",\".\",\".\",\"7\",\"9\"]]", Output: "true"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"board": [][]string{{"5", "3", "."}, {"6", ".", "."}, {".", "9", "8"}}}, Expected: true},
		},
		TimeComplexity:  "O(81)",
		SpaceComplexity: "O(81)",
		StarterCode:     "def isValidSudoku(board):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use sets to track seen numbers in each row, column, and 3x3 box."},
			{Level: 2, Type: "algorithm", Content: "Box index can be calculated as (row // 3) * 3 + (col // 3)."},
		},
		Solution: Solution{
			Code: `def isValidSudoku(board):
    rows = [set() for _ in range(9)]
    cols = [set() for _ in range(9)]
    boxes = [set() for _ in range(9)]

    for r in range(9):
        for c in range(9):
            if board[r][c] == '.':
                continue
            num = board[r][c]
            box_idx = (r // 3) * 3 + (c // 3)

            if num in rows[r] or num in cols[c] or num in boxes[box_idx]:
                return False

            rows[r].add(num)
            cols[c].add(num)
            boxes[box_idx].add(num)

    return True`,
			Explanation: "Track seen numbers in each row, column, and 3x3 box using sets.",
		},
	},
	{
		ID:         "encode-and-decode-strings",
		Number:     8,
		Title:      "Encode and Decode Strings",
		Difficulty: "Medium",
		Category:   "arrays-hashing",
		Tags:       []string{"Array", "String", "Design"},
		Description: `Design an algorithm to encode a list of strings to a string. The encoded string is then sent over the network and is decoded back to the original list of strings.

Implement encode and decode methods.`,
		Constraints: []string{
			"0 <= strs.length < 100",
			"0 <= strs[i].length < 200",
			"strs[i] contains any possible characters out of 256 valid ASCII characters",
		},
		Examples: []Example{
			{Input: `strs = ["hello","world"]`, Output: `["hello","world"]`},
			{Input: `strs = [""]`, Output: `[""]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"strs": []string{"hello", "world"}}, Expected: []string{"hello", "world"}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def encode(strs):\n    # Write your solution here\n    pass\n\ndef decode(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use length prefix to handle any characters including delimiters."},
			{Level: 2, Type: "algorithm", Content: "Format: length + delimiter + string for each word."},
		},
		Solution: Solution{
			Code: `def encode(strs):
    result = ""
    for s in strs:
        result += str(len(s)) + "#" + s
    return result

def decode(s):
    result = []
    i = 0
    while i < len(s):
        j = i
        while s[j] != '#':
            j += 1
        length = int(s[i:j])
        result.append(s[j + 1:j + 1 + length])
        i = j + 1 + length
    return result`,
			Explanation: "Encode with length prefix and # delimiter. Decode by reading length then extracting string.",
		},
	},
	{
		ID:         "longest-consecutive-sequence",
		Number:     9,
		Title:      "Longest Consecutive Sequence",
		Difficulty: "Medium",
		Category:   "arrays-hashing",
		Tags:       []string{"Array", "Hash Table", "Union Find"},
		Description: `Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.`,
		Constraints: []string{
			"0 <= nums.length <= 10^5",
			"-10^9 <= nums[i] <= 10^9",
		},
		Examples: []Example{
			{Input: "nums = [100,4,200,1,3,2]", Output: "4", Explanation: "The longest consecutive sequence is [1, 2, 3, 4]."},
			{Input: "nums = [0,3,7,2,5,8,4,6,0,1]", Output: "9"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{100, 4, 200, 1, 3, 2}}, Expected: 4},
			{Input: map[string]interface{}{"nums": []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}}, Expected: 9},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def longestConsecutive(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a set for O(1) lookups."},
			{Level: 2, Type: "algorithm", Content: "Only start counting from numbers that don't have num-1 in the set (sequence starts)."},
		},
		Solution: Solution{
			Code: `def longestConsecutive(nums):
    num_set = set(nums)
    longest = 0

    for num in num_set:
        if num - 1 not in num_set:  # Start of sequence
            length = 1
            while num + length in num_set:
                length += 1
            longest = max(longest, length)

    return longest`,
			Explanation: "Use set for O(1) lookup. Only start counting from sequence beginnings (no num-1).",
		},
	},
	// Two Pointers (continued)
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
	// Sliding Window (continued)
	{
		ID:         "longest-repeating-character-replacement",
		Number:     17,
		Title:      "Longest Repeating Character Replacement",
		Difficulty: "Medium",
		Category:   "sliding-window",
		Tags:       []string{"Hash Table", "String", "Sliding Window"},
		Description: `You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.`,
		Constraints: []string{
			"1 <= s.length <= 10^5",
			"s consists of only uppercase English letters",
			"0 <= k <= s.length",
		},
		Examples: []Example{
			{Input: `s = "ABAB", k = 2`, Output: "4"},
			{Input: `s = "AABABBA", k = 1`, Output: "4"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "ABAB", "k": 2}, Expected: 4},
			{Input: map[string]interface{}{"s": "AABABBA", "k": 1}, Expected: 4},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(26)",
		StarterCode:     "def characterReplacement(s, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Sliding window where window size - max freq <= k."},
			{Level: 2, Type: "algorithm", Content: "Track character frequencies in window. Shrink if replacements needed > k."},
		},
		Solution: Solution{
			Code: `def characterReplacement(s, k):
    count = {}
    max_freq = 0
    left = 0
    result = 0

    for right in range(len(s)):
        count[s[right]] = count.get(s[right], 0) + 1
        max_freq = max(max_freq, count[s[right]])

        while (right - left + 1) - max_freq > k:
            count[s[left]] -= 1
            left += 1

        result = max(result, right - left + 1)

    return result`,
			Explanation: "Window is valid if (size - max_freq) <= k. Track max frequency to minimize replacements.",
		},
	},
	{
		ID:         "permutation-in-string",
		Number:     18,
		Title:      "Permutation in String",
		Difficulty: "Medium",
		Category:   "sliding-window",
		Tags:       []string{"Hash Table", "Two Pointers", "String", "Sliding Window"},
		Description: `Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.`,
		Constraints: []string{
			"1 <= s1.length, s2.length <= 10^4",
			"s1 and s2 consist of lowercase English letters",
		},
		Examples: []Example{
			{Input: `s1 = "ab", s2 = "eidbaooo"`, Output: "true"},
			{Input: `s1 = "ab", s2 = "eidboaoo"`, Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s1": "ab", "s2": "eidbaooo"}, Expected: true},
			{Input: map[string]interface{}{"s1": "ab", "s2": "eidboaoo"}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(26)",
		StarterCode:     "def checkInclusion(s1, s2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Fixed-size sliding window of length s1."},
			{Level: 2, Type: "algorithm", Content: "Compare character counts in window with s1's counts."},
		},
		Solution: Solution{
			Code: `def checkInclusion(s1, s2):
    if len(s1) > len(s2):
        return False

    s1_count = {}
    window_count = {}

    for c in s1:
        s1_count[c] = s1_count.get(c, 0) + 1

    for i, c in enumerate(s2):
        window_count[c] = window_count.get(c, 0) + 1

        if i >= len(s1):
            left_char = s2[i - len(s1)]
            window_count[left_char] -= 1
            if window_count[left_char] == 0:
                del window_count[left_char]

        if window_count == s1_count:
            return True

    return False`,
			Explanation: "Fixed sliding window of s1's length. Compare character counts.",
		},
	},
	{
		ID:         "minimum-window-substring",
		Number:     19,
		Title:      "Minimum Window Substring",
		Difficulty: "Hard",
		Category:   "sliding-window",
		Tags:       []string{"Hash Table", "String", "Sliding Window"},
		Description: `Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".`,
		Constraints: []string{
			"m == s.length",
			"n == t.length",
			"1 <= m, n <= 10^5",
			"s and t consist of uppercase and lowercase English letters",
		},
		Examples: []Example{
			{Input: `s = "ADOBECODEBANC", t = "ABC"`, Output: `"BANC"`},
			{Input: `s = "a", t = "a"`, Output: `"a"`},
			{Input: `s = "a", t = "aa"`, Output: `""`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "ADOBECODEBANC", "t": "ABC"}, Expected: "BANC"},
			{Input: map[string]interface{}{"s": "a", "t": "a"}, Expected: "a"},
		},
		TimeComplexity:  "O(m + n)",
		SpaceComplexity: "O(m + n)",
		StarterCode:     "def minWindow(s, t):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Expand window until valid, then contract to find minimum."},
			{Level: 2, Type: "algorithm", Content: "Track how many required characters are satisfied. Shrink when all satisfied."},
		},
		Solution: Solution{
			Code: `def minWindow(s, t):
    if not t or not s:
        return ""

    t_count = {}
    for c in t:
        t_count[c] = t_count.get(c, 0) + 1

    required = len(t_count)
    formed = 0
    window_count = {}

    left = 0
    min_len = float('inf')
    min_left = 0

    for right in range(len(s)):
        c = s[right]
        window_count[c] = window_count.get(c, 0) + 1

        if c in t_count and window_count[c] == t_count[c]:
            formed += 1

        while formed == required:
            if right - left + 1 < min_len:
                min_len = right - left + 1
                min_left = left

            left_c = s[left]
            window_count[left_c] -= 1
            if left_c in t_count and window_count[left_c] < t_count[left_c]:
                formed -= 1
            left += 1

    return "" if min_len == float('inf') else s[min_left:min_left + min_len]`,
			Explanation: "Expand until all t chars found, contract to minimize while valid.",
		},
	},
	{
		ID:         "sliding-window-maximum",
		Number:     20,
		Title:      "Sliding Window Maximum",
		Difficulty: "Hard",
		Category:   "sliding-window",
		Tags:       []string{"Array", "Queue", "Sliding Window", "Heap", "Monotonic Queue"},
		Description: `You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.`,
		Constraints: []string{
			"1 <= nums.length <= 10^5",
			"-10^4 <= nums[i] <= 10^4",
			"1 <= k <= nums.length",
		},
		Examples: []Example{
			{Input: "nums = [1,3,-1,-3,5,3,6,7], k = 3", Output: "[3,3,5,5,6,7]"},
			{Input: "nums = [1], k = 1", Output: "[1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 3, -1, -3, 5, 3, 6, 7}, "k": 3}, Expected: []int{3, 3, 5, 5, 6, 7}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(k)",
		StarterCode:     "def maxSlidingWindow(nums, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a monotonic decreasing deque to track potential maximums."},
			{Level: 2, Type: "algorithm", Content: "Remove smaller elements from back, remove out-of-window from front."},
		},
		Solution: Solution{
			Code: `from collections import deque

def maxSlidingWindow(nums, k):
    dq = deque()  # Store indices
    result = []

    for i in range(len(nums)):
        # Remove indices outside window
        while dq and dq[0] < i - k + 1:
            dq.popleft()

        # Remove smaller elements
        while dq and nums[dq[-1]] < nums[i]:
            dq.pop()

        dq.append(i)

        if i >= k - 1:
            result.append(nums[dq[0]])

    return result`,
			Explanation: "Monotonic decreasing deque. Front always has max for current window.",
		},
	},
	// Stack (continued)
	{
		ID:         "valid-parentheses",
		Number:     19,
		Title:      "Valid Parentheses",
		Difficulty: "Easy",
		Category:   "stack",
		Tags:       []string{"String", "Stack"},
		RelatedChapters: []int{8},
		Description: `Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.`,
		Constraints: []string{
			"1 <= s.length <= 10^4",
			"s consists of parentheses only '()[]{}'",
		},
		Examples: []Example{
			{Input: `s = "()"`, Output: "true"},
			{Input: `s = "()[]{}"`, Output: "true"},
			{Input: `s = "(]"`, Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "()"}, Expected: true},
			{Input: map[string]interface{}{"s": "()[]{}"}, Expected: true},
			{Input: map[string]interface{}{"s": "(]"}, Expected: false},
			{Input: map[string]interface{}{"s": "([)]"}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def isValid(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a stack to track open brackets."},
			{Level: 2, Type: "algorithm", Content: "Push opens, pop and match for closes."},
		},
		Solution: Solution{
			Code: `def isValid(s):
    stack = []
    pairs = {')': '(', '}': '{', ']': '['}
    for char in s:
        if char in pairs:
            if not stack or stack[-1] != pairs[char]:
                return False
            stack.pop()
        else:
            stack.append(char)
    return len(stack) == 0`,
			Explanation: "Stack for opens. For closes, check if matches top of stack.",
		},
	},
	// Stack (continued)
	{
		ID:              "min-stack",
		Number:          40,
		Title:           "Min Stack",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Stack", "Design"},
		RelatedChapters: []int{3, 8, 12},
		Description: `Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:
- MinStack() initializes the stack object.
- void push(int val) pushes the element val onto the stack.
- void pop() removes the element on the top of the stack.
- int top() gets the top element of the stack.
- int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function.`,
		Constraints: []string{
			"-2^31 <= val <= 2^31 - 1",
			"Methods pop, top and getMin operations will always be called on non-empty stacks",
			"At most 3 * 10^4 calls will be made to push, pop, top, and getMin",
		},
		Examples: []Example{
			{Input: `["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]`, Output: "[null,null,null,null,-3,null,0,-2]", Explanation: "MinStack minStack = new MinStack(); minStack.push(-2); minStack.push(0); minStack.push(-3); minStack.getMin(); // return -3; minStack.pop(); minStack.top(); // return 0; minStack.getMin(); // return -2"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}, "values": [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, -3, nil, 0, -2}},
		},
		TimeComplexity:  "O(1)",
		SpaceComplexity: "O(n)",
		StarterCode: `class MinStack:
    def __init__(self):
        # Initialize your data structure here
        pass

    def push(self, val: int) -> None:
        pass

    def pop(self) -> None:
        pass

    def top(self) -> int:
        pass

    def getMin(self) -> int:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Store additional information with each element to track the minimum at that point in the stack.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "Use a stack of tuples (value, current_min). Each push stores the value along with the minimum of all elements at or below it."},
			{Level: 3, Type: "code", Content: "On push: new_min = min(val, self.stack[-1][1]) if stack else val. Store (val, new_min). getMin returns stack[-1][1]."},
		},
		Solution: Solution{
			Code: `class MinStack:
    def __init__(self):
        self.stack = []  # Each element: (value, min_so_far)

    def push(self, val: int) -> None:
        if not self.stack:
            self.stack.append((val, val))
        else:
            current_min = min(val, self.stack[-1][1])
            self.stack.append((val, current_min))

    def pop(self) -> None:
        self.stack.pop()

    def top(self) -> int:
        return self.stack[-1][0]

    def getMin(self) -> int:
        return self.stack[-1][1]`,
			Explanation:     "Store (value, minimum) pairs. Each element knows the minimum of all elements at or below it in the stack.",
			TimeComplexity:  "O(1) for all operations",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize stack", Explanation: "Create empty list to store (value, min) tuples", CodeSnippet: "self.stack = []", LineStart: 3, LineEnd: 3},
				{Title: "Push with min tracking", Explanation: "Calculate new minimum by comparing with previous min", CodeSnippet: "current_min = min(val, self.stack[-1][1])", LineStart: 8, LineEnd: 9},
				{Title: "O(1) getMin", Explanation: "Minimum is always at top of stack's second element", CodeSnippet: "return self.stack[-1][1]", LineStart: 16, LineEnd: 16},
			},
		},
	},
	{
		ID:              "evaluate-reverse-polish-notation",
		Number:          41,
		Title:           "Evaluate Reverse Polish Notation",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Array", "Math", "Stack"},
		RelatedChapters: []int{3, 8},
		Description: `You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:
- The valid operators are '+', '-', '*', and '/'.
- Each operand may be an integer or another expression.
- The division between two integers always truncates toward zero.
- There will not be any division by zero.
- The input represents a valid arithmetic expression in reverse polish notation.
- The answer and all the intermediate calculations can be represented in a 32-bit integer.`,
		Constraints: []string{
			"1 <= tokens.length <= 10^4",
			"tokens[i] is either an operator: '+', '-', '*', or '/', or an integer in the range [-200, 200]",
		},
		Examples: []Example{
			{Input: `tokens = ["2","1","+","3","*"]`, Output: "9", Explanation: "((2 + 1) * 3) = 9"},
			{Input: `tokens = ["4","13","5","/","+"]`, Output: "6", Explanation: "(4 + (13 / 5)) = 6"},
			{Input: `tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]`, Output: "22"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"tokens": []string{"2", "1", "+", "3", "*"}}, Expected: 9},
			{Input: map[string]interface{}{"tokens": []string{"4", "13", "5", "/", "+"}}, Expected: 6},
			{Input: map[string]interface{}{"tokens": []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}}, Expected: 22},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def evalRPN(tokens):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a stack. Numbers get pushed; operators pop two operands, compute, and push the result."},
			{Level: 2, Type: "algorithm", Content: "For each token: if it's a number, push it. If it's an operator, pop two numbers, apply the operator, push the result."},
			{Level: 3, Type: "code", Content: "Watch out for order: b = stack.pop(), a = stack.pop(), then compute a op b. For division, use int(a/b) to truncate toward zero."},
		},
		Solution: Solution{
			Code: `def evalRPN(tokens):
    stack = []
    operators = {'+', '-', '*', '/'}

    for token in tokens:
        if token in operators:
            b = stack.pop()
            a = stack.pop()
            if token == '+':
                stack.append(a + b)
            elif token == '-':
                stack.append(a - b)
            elif token == '*':
                stack.append(a * b)
            else:  # division
                stack.append(int(a / b))  # truncate toward zero
        else:
            stack.append(int(token))

    return stack[0]`,
			Explanation:     "Process tokens left to right. Numbers go on stack. Operators pop two operands, compute, push result.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize stack", Explanation: "Stack will hold intermediate results", CodeSnippet: "stack = []", LineStart: 2, LineEnd: 2},
				{Title: "Process each token", Explanation: "Check if token is operator or number", CodeSnippet: "if token in operators:", LineStart: 6, LineEnd: 6},
				{Title: "Apply operator", Explanation: "Pop two operands (order matters!), apply operator, push result", CodeSnippet: "b = stack.pop()\na = stack.pop()", LineStart: 7, LineEnd: 8},
				{Title: "Handle division", Explanation: "Truncate toward zero using int(a/b)", CodeSnippet: "stack.append(int(a / b))", LineStart: 16, LineEnd: 16},
			},
		},
	},
	{
		ID:              "generate-parentheses",
		Number:          42,
		Title:           "Generate Parentheses",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"String", "Dynamic Programming", "Backtracking"},
		RelatedChapters: []int{3, 4, 12},
		Description: `Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.`,
		Constraints: []string{
			"1 <= n <= 8",
		},
		Examples: []Example{
			{Input: "n = 3", Output: `["((()))","(()())","(())()","()(())","()()()"]`},
			{Input: "n = 1", Output: `["()"]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 3}, Expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
			{Input: map[string]interface{}{"n": 1}, Expected: []string{"()"}},
			{Input: map[string]interface{}{"n": 2}, Expected: []string{"(())", "()()"}},
		},
		TimeComplexity:  "O(4^n / sqrt(n))",
		SpaceComplexity: "O(n)",
		StarterCode:     "def generateParenthesis(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use backtracking. At each step, you can add '(' if open count < n, or ')' if close count < open count.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "Track open and close counts. Only add ')' when close < open to ensure validity. Base case: both counts equal n."},
			{Level: 3, Type: "code", Content: "def backtrack(s, open, close): if open == close == n: result.append(s). Add '(' if open < n, add ')' if close < open."},
		},
		Solution: Solution{
			Code: `def generateParenthesis(n):
    result = []

    def backtrack(current, open_count, close_count):
        if len(current) == 2 * n:
            result.append(current)
            return

        if open_count < n:
            backtrack(current + '(', open_count + 1, close_count)
        if close_count < open_count:
            backtrack(current + ')', open_count, close_count + 1)

    backtrack('', 0, 0)
    return result`,
			Explanation:     "Backtracking with constraints: add '(' if open < n, add ')' only if close < open to maintain validity.",
			TimeComplexity:  "O(4^n / sqrt(n)) - nth Catalan number",
			SpaceComplexity: "O(n) - recursion depth",
			Walkthrough: []WalkthroughStep{
				{Title: "Base case", Explanation: "When string length is 2n, we have a valid combination", CodeSnippet: "if len(current) == 2 * n:", LineStart: 5, LineEnd: 7},
				{Title: "Add open paren", Explanation: "Can add '(' if we haven't used all n opens", CodeSnippet: "if open_count < n:", LineStart: 9, LineEnd: 10},
				{Title: "Add close paren", Explanation: "Can add ')' only if close < open (ensures validity)", CodeSnippet: "if close_count < open_count:", LineStart: 11, LineEnd: 12},
			},
		},
	},
	{
		ID:              "daily-temperatures",
		Number:          43,
		Title:           "Daily Temperatures",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Monotonic Stack"},
		RelatedChapters: []int{3, 8},
		Description: `Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.`,
		Constraints: []string{
			"1 <= temperatures.length <= 10^5",
			"30 <= temperatures[i] <= 100",
		},
		Examples: []Example{
			{Input: "temperatures = [73,74,75,71,69,72,76,73]", Output: "[1,1,4,2,1,1,0,0]"},
			{Input: "temperatures = [30,40,50,60]", Output: "[1,1,1,0]"},
			{Input: "temperatures = [30,60,90]", Output: "[1,1,0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"temperatures": []int{73, 74, 75, 71, 69, 72, 76, 73}}, Expected: []int{1, 1, 4, 2, 1, 1, 0, 0}},
			{Input: map[string]interface{}{"temperatures": []int{30, 40, 50, 60}}, Expected: []int{1, 1, 1, 0}},
			{Input: map[string]interface{}{"temperatures": []int{30, 60, 90}}, Expected: []int{1, 1, 0}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def dailyTemperatures(temperatures):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a monotonic decreasing stack to track indices of temperatures waiting for a warmer day."},
			{Level: 2, Type: "algorithm", Content: "Stack stores indices. When current temp > stack top's temp, pop and calculate the wait time (current index - popped index)."},
			{Level: 3, Type: "code", Content: "while stack and temps[i] > temps[stack[-1]]: idx = stack.pop(); result[idx] = i - idx. Always push current index."},
		},
		Solution: Solution{
			Code: `def dailyTemperatures(temperatures):
    n = len(temperatures)
    result = [0] * n
    stack = []  # Store indices

    for i in range(n):
        while stack and temperatures[i] > temperatures[stack[-1]]:
            idx = stack.pop()
            result[idx] = i - idx
        stack.append(i)

    return result`,
			Explanation:     "Monotonic decreasing stack of indices. When we find a warmer day, pop all colder days and record their wait times.",
			TimeComplexity:  "O(n) - each index pushed and popped at most once",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize result", Explanation: "Default 0 (no warmer day found)", CodeSnippet: "result = [0] * n", LineStart: 3, LineEnd: 3},
				{Title: "Process warmer days", Explanation: "Pop indices with lower temps and calculate wait time", CodeSnippet: "while stack and temperatures[i] > temperatures[stack[-1]]:", LineStart: 7, LineEnd: 9},
				{Title: "Always push current", Explanation: "Current index may need a future warmer day", CodeSnippet: "stack.append(i)", LineStart: 10, LineEnd: 10},
			},
		},
	},
	{
		ID:              "car-fleet",
		Number:          44,
		Title:           "Car Fleet",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Sorting", "Monotonic Stack"},
		RelatedChapters: []int{3, 8, 12},
		Description: `There are n cars going to the same destination along a one-lane road. The destination is target miles away.

You are given two integer arrays position and speed, both of length n, where position[i] is the position of the ith car and speed[i] is the speed of the ith car (in miles per hour).

A car can never pass another car ahead of it, but it can catch up to it and drive bumper to bumper at the same speed. The faster car will slow down to match the slower car's speed. The distance between these two cars is ignored (i.e., they are assumed to have the same position).

A car fleet is some non-empty set of cars driving at the same position and same speed. Note that a single car is also a car fleet.

If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.

Return the number of car fleets that will arrive at the destination.`,
		Constraints: []string{
			"n == position.length == speed.length",
			"1 <= n <= 10^5",
			"0 < target <= 10^6",
			"0 <= position[i] < target",
			"All the values of position are unique",
			"0 < speed[i] <= 10^6",
		},
		Examples: []Example{
			{Input: "target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]", Output: "3", Explanation: "Cars at 10,8,0,5,3 form 3 fleets: [10], [8], [0,5,3] merge at different times."},
			{Input: "target = 10, position = [3], speed = [3]", Output: "1"},
			{Input: "target = 100, position = [0,2,4], speed = [4,2,1]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"target": 12, "position": []int{10, 8, 0, 5, 3}, "speed": []int{2, 4, 1, 1, 3}}, Expected: 3},
			{Input: map[string]interface{}{"target": 10, "position": []int{3}, "speed": []int{3}}, Expected: 1},
			{Input: map[string]interface{}{"target": 100, "position": []int{0, 2, 4}, "speed": []int{4, 2, 1}}, Expected: 1},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def carFleet(target, position, speed):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Sort cars by position (closest to target first). Calculate time to reach target for each car.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "If a car behind takes less time than one ahead, they merge into one fleet (the slower car's time). Count remaining distinct times."},
			{Level: 3, Type: "code", Content: "Sort by position descending. Use a stack of arrival times. If current time <= stack top, it merges. Otherwise, push new fleet."},
		},
		Solution: Solution{
			Code: `def carFleet(target, position, speed):
    # Pair position and speed, sort by position descending
    cars = sorted(zip(position, speed), reverse=True)
    stack = []  # arrival times of fleets

    for pos, spd in cars:
        time = (target - pos) / spd
        # If this car takes longer, it's a new fleet
        if not stack or time > stack[-1]:
            stack.append(time)
        # Otherwise, it merges with the fleet ahead (don't push)

    return len(stack)`,
			Explanation:     "Sort by position (closest first). Track arrival times with stack. Slower cars (longer time) start new fleets; faster cars merge.",
			TimeComplexity:  "O(n log n) for sorting",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort by position", Explanation: "Process cars from closest to target to farthest", CodeSnippet: "cars = sorted(zip(position, speed), reverse=True)", LineStart: 3, LineEnd: 3},
				{Title: "Calculate arrival time", Explanation: "Time = distance / speed", CodeSnippet: "time = (target - pos) / spd", LineStart: 7, LineEnd: 7},
				{Title: "New fleet or merge", Explanation: "If slower than fleet ahead, start new fleet; otherwise merge", CodeSnippet: "if not stack or time > stack[-1]:", LineStart: 9, LineEnd: 10},
			},
		},
	},
	{
		ID:              "largest-rectangle-in-histogram",
		Number:          45,
		Title:           "Largest Rectangle in Histogram",
		Difficulty:      "Hard",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Monotonic Stack"},
		RelatedChapters: []int{3, 8, 12},
		Description: `Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.`,
		Constraints: []string{
			"1 <= heights.length <= 10^5",
			"0 <= heights[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "heights = [2,1,5,6,2,3]", Output: "10", Explanation: "The largest rectangle has area = 10 units (bars at index 2-3 with height 5)."},
			{Input: "heights = [2,4]", Output: "4"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"heights": []int{2, 1, 5, 6, 2, 3}}, Expected: 10},
			{Input: map[string]interface{}{"heights": []int{2, 4}}, Expected: 4},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def largestRectangleArea(heights):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a monotonic increasing stack. When we see a smaller bar, calculate areas for all taller bars that can no longer extend right."},
			{Level: 2, Type: "algorithm", Content: "Stack stores indices. For each bar popped, its right boundary is current index, left boundary is new stack top. Width = right - left - 1."},
			{Level: 3, Type: "code", Content: "Add a sentinel height 0 at end to flush stack. When popping: width = i - stack[-1] - 1, area = heights[idx] * width."},
		},
		Solution: Solution{
			Code: `def largestRectangleArea(heights):
    stack = []  # indices
    max_area = 0
    heights.append(0)  # sentinel to flush stack at end

    for i, h in enumerate(heights):
        while stack and heights[stack[-1]] > h:
            height = heights[stack.pop()]
            width = i if not stack else i - stack[-1] - 1
            max_area = max(max_area, height * width)
        stack.append(i)

    heights.pop()  # remove sentinel
    return max_area`,
			Explanation:     "Monotonic increasing stack. When a shorter bar appears, pop taller bars and calculate their max possible rectangle area.",
			TimeComplexity:  "O(n) - each bar pushed and popped once",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Add sentinel", Explanation: "Ensures all bars are processed by adding height 0 at end", CodeSnippet: "heights.append(0)", LineStart: 4, LineEnd: 4},
				{Title: "Pop taller bars", Explanation: "Calculate area when we can't extend right anymore", CodeSnippet: "while stack and heights[stack[-1]] > h:", LineStart: 7, LineEnd: 10},
				{Title: "Calculate width", Explanation: "Width extends from previous stack element to current index", CodeSnippet: "width = i if not stack else i - stack[-1] - 1", LineStart: 9, LineEnd: 9},
			},
		},
	},
	// Binary Search (continued)
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
	// Linked List (continued)
	{
		ID:              "merge-two-sorted-lists",
		Number:          52,
		Title:           "Merge Two Sorted Lists",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Recursion"},
		RelatedChapters: []int{2, 7},
		Description: `You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.`,
		Constraints: []string{
			"The number of nodes in both lists is in the range [0, 50]",
			"-100 <= Node.val <= 100",
			"Both list1 and list2 are sorted in non-decreasing order",
		},
		Examples: []Example{
			{Input: "list1 = [1,2,4], list2 = [1,3,4]", Output: "[1,1,2,3,4,4]"},
			{Input: "list1 = [], list2 = []", Output: "[]"},
			{Input: "list1 = [], list2 = [0]", Output: "[0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"list1": []int{1, 2, 4}, "list2": []int{1, 3, 4}}, Expected: []int{1, 1, 2, 3, 4, 4}},
			{Input: map[string]interface{}{"list1": []int{}, "list2": []int{}}, Expected: []int{}},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def mergeTwoLists(list1, list2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a dummy head to simplify edge cases. Compare nodes and link the smaller one.", ChapterRef: 2},
			{Level: 2, Type: "algorithm", Content: "Maintain a tail pointer. At each step, link tail.next to the smaller of the two current nodes, advance that list."},
			{Level: 3, Type: "code", Content: "dummy = ListNode(0); tail = dummy. while l1 and l2: compare and link. Append remaining nodes at end."},
		},
		Solution: Solution{
			Code: `def mergeTwoLists(list1, list2):
    dummy = ListNode(0)
    tail = dummy

    while list1 and list2:
        if list1.val <= list2.val:
            tail.next = list1
            list1 = list1.next
        else:
            tail.next = list2
            list2 = list2.next
        tail = tail.next

    # Append remaining nodes
    tail.next = list1 if list1 else list2

    return dummy.next`,
			Explanation:     "Use dummy head and tail pointer. Compare and link smaller node, advance that list. Append remaining at end.",
			TimeComplexity:  "O(n + m)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Dummy head", Explanation: "Simplifies handling of head node", CodeSnippet: "dummy = ListNode(0)\ntail = dummy", LineStart: 2, LineEnd: 3},
				{Title: "Compare and link", Explanation: "Link smaller node and advance that list", CodeSnippet: "if list1.val <= list2.val:\n    tail.next = list1", LineStart: 6, LineEnd: 8},
				{Title: "Append remaining", Explanation: "One list may have remaining nodes", CodeSnippet: "tail.next = list1 if list1 else list2", LineStart: 15, LineEnd: 15},
			},
		},
	},
	{
		ID:              "reorder-list",
		Number:          53,
		Title:           "Reorder List",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Two Pointers", "Stack", "Recursion"},
		RelatedChapters: []int{2, 7},
		Description: `You are given the head of a singly linked-list. The list can be represented as:

L0 -> L1 -> ... -> Ln - 1 -> Ln

Reorder the list to be on the following form:

L0 -> Ln -> L1 -> Ln - 1 -> L2 -> Ln - 2 -> ...

You may not modify the values in the list's nodes. Only nodes themselves may be changed.`,
		Constraints: []string{
			"The number of nodes in the list is in the range [1, 5 * 10^4]",
			"1 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4]", Output: "[1,4,2,3]"},
			{Input: "head = [1,2,3,4,5]", Output: "[1,5,2,4,3]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4}}, Expected: []int{1, 4, 2, 3}},
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}}, Expected: []int{1, 5, 2, 4, 3}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def reorderList(head):\n    # Write your solution here (modify in place)\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Three steps: find middle, reverse second half, merge two halves alternating.", ChapterRef: 2},
			{Level: 2, Type: "algorithm", Content: "Use slow/fast pointers to find middle. Reverse from middle to end. Interleave first and reversed second halves."},
			{Level: 3, Type: "code", Content: "Find mid with slow/fast. Reverse(slow.next). Merge by alternating: first.next = second, then second.next = first.next."},
		},
		Solution: Solution{
			Code: `def reorderList(head):
    if not head or not head.next:
        return

    # Find middle
    slow, fast = head, head.next
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next

    # Reverse second half
    second = slow.next
    slow.next = None
    prev = None
    while second:
        tmp = second.next
        second.next = prev
        prev = second
        second = tmp

    # Merge two halves
    first, second = head, prev
    while second:
        tmp1, tmp2 = first.next, second.next
        first.next = second
        second.next = tmp1
        first, second = tmp1, tmp2`,
			Explanation:     "Find middle (slow/fast), reverse second half, merge alternating nodes from both halves.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Find middle", Explanation: "Slow pointer ends at middle after fast reaches end", CodeSnippet: "while fast and fast.next:\n    slow = slow.next", LineStart: 6, LineEnd: 9},
				{Title: "Reverse second half", Explanation: "Standard linked list reversal", CodeSnippet: "prev = None\nwhile second:", LineStart: 13, LineEnd: 18},
				{Title: "Merge halves", Explanation: "Interleave nodes from first and second halves", CodeSnippet: "first.next = second\nsecond.next = tmp1", LineStart: 22, LineEnd: 26},
			},
		},
	},
	{
		ID:              "remove-nth-node-from-end",
		Number:          54,
		Title:           "Remove Nth Node From End of List",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Two Pointers"},
		RelatedChapters: []int{2, 7},
		Description: `Given the head of a linked list, remove the nth node from the end of the list and return its head.`,
		Constraints: []string{
			"The number of nodes in the list is sz",
			"1 <= sz <= 30",
			"0 <= Node.val <= 100",
			"1 <= n <= sz",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4,5], n = 2", Output: "[1,2,3,5]"},
			{Input: "head = [1], n = 1", Output: "[]"},
			{Input: "head = [1,2], n = 1", Output: "[1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}, "n": 2}, Expected: []int{1, 2, 3, 5}},
			{Input: map[string]interface{}{"head": []int{1}, "n": 1}, Expected: []int{}},
			{Input: map[string]interface{}{"head": []int{1, 2}, "n": 1}, Expected: []int{1}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def removeNthFromEnd(head, n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use two pointers with n nodes gap. When fast reaches end, slow is at the node before the one to remove."},
			{Level: 2, Type: "algorithm", Content: "Use dummy node to handle edge case of removing head. Advance fast n+1 steps, then move both until fast is null."},
			{Level: 3, Type: "code", Content: "dummy.next = head. Advance fast n+1 times. while fast: move both. slow.next = slow.next.next."},
		},
		Solution: Solution{
			Code: `def removeNthFromEnd(head, n):
    dummy = ListNode(0, head)
    slow = fast = dummy

    # Advance fast n+1 steps ahead
    for _ in range(n + 1):
        fast = fast.next

    # Move both until fast reaches end
    while fast:
        slow = slow.next
        fast = fast.next

    # Remove the node
    slow.next = slow.next.next

    return dummy.next`,
			Explanation:     "Two pointers n+1 apart. When fast reaches null, slow is just before the node to remove. Use dummy to handle head removal.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Use dummy node", Explanation: "Handles edge case of removing head", CodeSnippet: "dummy = ListNode(0, head)", LineStart: 2, LineEnd: 2},
				{Title: "Create gap", Explanation: "Fast is n+1 steps ahead of slow", CodeSnippet: "for _ in range(n + 1):\n    fast = fast.next", LineStart: 6, LineEnd: 7},
				{Title: "Remove node", Explanation: "Slow points to node before target", CodeSnippet: "slow.next = slow.next.next", LineStart: 15, LineEnd: 15},
			},
		},
	},
	{
		ID:              "copy-list-with-random-pointer",
		Number:          55,
		Title:           "Copy List with Random Pointer",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Hash Table", "Linked List"},
		RelatedChapters: []int{2, 5, 7},
		Description: `A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state.

Return the head of the copied linked list.`,
		Constraints: []string{
			"0 <= n <= 1000",
			"-10^4 <= Node.val <= 10^4",
			"Node.random is null or is pointing to some node in the linked list",
		},
		Examples: []Example{
			{Input: "head = [[7,null],[13,0],[11,4],[10,2],[1,0]]", Output: "[[7,null],[13,0],[11,4],[10,2],[1,0]]"},
			{Input: "head = [[1,1],[2,1]]", Output: "[[1,1],[2,1]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}}, Expected: [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def copyRandomList(head):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a hash map to map original nodes to their copies. Two passes: create all copies, then set pointers.", ChapterRef: 5},
			{Level: 2, Type: "algorithm", Content: "First pass: create copy nodes and store in map. Second pass: set next and random pointers using the map."},
			{Level: 3, Type: "code", Content: "old_to_new = {None: None}. Create copies. For each old node: copy.next = map[old.next], copy.random = map[old.random]."},
		},
		Solution: Solution{
			Code: `def copyRandomList(head):
    if not head:
        return None

    # Map old nodes to new nodes
    old_to_new = {None: None}

    # First pass: create all copy nodes
    curr = head
    while curr:
        old_to_new[curr] = Node(curr.val)
        curr = curr.next

    # Second pass: set next and random pointers
    curr = head
    while curr:
        copy = old_to_new[curr]
        copy.next = old_to_new[curr.next]
        copy.random = old_to_new[curr.random]
        curr = curr.next

    return old_to_new[head]`,
			Explanation:     "Hash map maps original to copy. First create all copies, then set next/random pointers using the map.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n) for the hash map",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize map", Explanation: "Map None to None for handling null pointers", CodeSnippet: "old_to_new = {None: None}", LineStart: 6, LineEnd: 6},
				{Title: "Create copies", Explanation: "First pass creates all new nodes", CodeSnippet: "old_to_new[curr] = Node(curr.val)", LineStart: 10, LineEnd: 11},
				{Title: "Set pointers", Explanation: "Second pass links next and random using map", CodeSnippet: "copy.next = old_to_new[curr.next]\ncopy.random = old_to_new[curr.random]", LineStart: 17, LineEnd: 19},
			},
		},
	},
	{
		ID:              "add-two-numbers",
		Number:          56,
		Title:           "Add Two Numbers",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Math", "Recursion"},
		RelatedChapters: []int{2, 7},
		Description: `You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.`,
		Constraints: []string{
			"The number of nodes in each linked list is in the range [1, 100]",
			"0 <= Node.val <= 9",
			"It is guaranteed that the list represents a number that does not have leading zeros",
		},
		Examples: []Example{
			{Input: "l1 = [2,4,3], l2 = [5,6,4]", Output: "[7,0,8]", Explanation: "342 + 465 = 807"},
			{Input: "l1 = [0], l2 = [0]", Output: "[0]"},
			{Input: "l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]", Output: "[8,9,9,9,0,0,0,1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"l1": []int{2, 4, 3}, "l2": []int{5, 6, 4}}, Expected: []int{7, 0, 8}},
			{Input: map[string]interface{}{"l1": []int{0}, "l2": []int{0}}, Expected: []int{0}},
			{Input: map[string]interface{}{"l1": []int{9, 9, 9, 9, 9, 9, 9}, "l2": []int{9, 9, 9, 9}}, Expected: []int{8, 9, 9, 9, 0, 0, 0, 1}},
		},
		TimeComplexity:  "O(max(m, n))",
		SpaceComplexity: "O(max(m, n))",
		StarterCode:     "def addTwoNumbers(l1, l2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Simulate addition digit by digit, tracking carry. Handle different lengths and final carry."},
			{Level: 2, Type: "algorithm", Content: "At each step: sum = l1.val + l2.val + carry. New digit = sum % 10. New carry = sum // 10."},
			{Level: 3, Type: "code", Content: "Use dummy head. while l1 or l2 or carry: sum and create new node. Don't forget final carry!"},
		},
		Solution: Solution{
			Code: `def addTwoNumbers(l1, l2):
    dummy = ListNode(0)
    curr = dummy
    carry = 0

    while l1 or l2 or carry:
        val1 = l1.val if l1 else 0
        val2 = l2.val if l2 else 0

        total = val1 + val2 + carry
        carry = total // 10
        curr.next = ListNode(total % 10)
        curr = curr.next

        l1 = l1.next if l1 else None
        l2 = l2.next if l2 else None

    return dummy.next`,
			Explanation:     "Add digit by digit with carry. Handle different lengths by treating missing digits as 0. Don't forget final carry.",
			TimeComplexity:  "O(max(m, n))",
			SpaceComplexity: "O(max(m, n))",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle different lengths", Explanation: "Use 0 if list is exhausted", CodeSnippet: "val1 = l1.val if l1 else 0", LineStart: 7, LineEnd: 8},
				{Title: "Calculate digit and carry", Explanation: "digit = total % 10, carry = total // 10", CodeSnippet: "carry = total // 10\ncurr.next = ListNode(total % 10)", LineStart: 11, LineEnd: 12},
				{Title: "Continue while carry", Explanation: "Loop continues if there's a carry to process", CodeSnippet: "while l1 or l2 or carry:", LineStart: 6, LineEnd: 6},
			},
		},
	},
	{
		ID:              "merge-k-sorted-lists",
		Number:          57,
		Title:           "Merge k Sorted Lists",
		Difficulty:      "Hard",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Divide and Conquer", "Heap", "Merge Sort"},
		RelatedChapters: []int{2, 7, 11},
		Description: `You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.`,
		Constraints: []string{
			"k == lists.length",
			"0 <= k <= 10^4",
			"0 <= lists[i].length <= 500",
			"-10^4 <= lists[i][j] <= 10^4",
			"lists[i] is sorted in ascending order",
			"The sum of lists[i].length will not exceed 10^4",
		},
		Examples: []Example{
			{Input: "lists = [[1,4,5],[1,3,4],[2,6]]", Output: "[1,1,2,3,4,4,5,6]"},
			{Input: "lists = []", Output: "[]"},
			{Input: "lists = [[]]", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"lists": [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}}, Expected: []int{1, 1, 2, 3, 4, 4, 5, 6}},
			{Input: map[string]interface{}{"lists": [][]int{}}, Expected: []int{}},
		},
		TimeComplexity:  "O(N log k)",
		SpaceComplexity: "O(k)",
		StarterCode:     "def mergeKLists(lists):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a min-heap to efficiently get the smallest element among k list heads. Or use divide and conquer with pairwise merge.", ChapterRef: 11},
			{Level: 2, Type: "algorithm", Content: "Heap approach: push all heads to heap. Pop smallest, add to result, push its next. Repeat until heap empty."},
			{Level: 3, Type: "code", Content: "heapq with (node.val, index, node) tuples. Index breaks ties. Pop min, push next if exists."},
		},
		Solution: Solution{
			Code: `import heapq

def mergeKLists(lists):
    dummy = ListNode(0)
    curr = dummy
    heap = []

    # Initialize heap with heads of all lists
    for i, lst in enumerate(lists):
        if lst:
            heapq.heappush(heap, (lst.val, i, lst))

    while heap:
        val, i, node = heapq.heappop(heap)
        curr.next = node
        curr = curr.next

        if node.next:
            heapq.heappush(heap, (node.next.val, i, node.next))

    return dummy.next`,
			Explanation:     "Min-heap keeps track of smallest among k list heads. Pop smallest, add to result, push its next node.",
			TimeComplexity:  "O(N log k) - N total nodes, log k for each heap operation",
			SpaceComplexity: "O(k) for the heap",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize heap", Explanation: "Add all non-null list heads", CodeSnippet: "for i, lst in enumerate(lists):\n    if lst:\n        heapq.heappush(heap, (lst.val, i, lst))", LineStart: 9, LineEnd: 11},
				{Title: "Pop and link", Explanation: "Get smallest, link to result", CodeSnippet: "val, i, node = heapq.heappop(heap)\ncurr.next = node", LineStart: 14, LineEnd: 16},
				{Title: "Push next", Explanation: "If popped node has next, push it", CodeSnippet: "if node.next:\n    heapq.heappush(heap, (node.next.val, i, node.next))", LineStart: 18, LineEnd: 19},
			},
		},
	},
	{
		ID:              "lru-cache",
		Number:          58,
		Title:           "LRU Cache",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Hash Table", "Linked List", "Design", "Doubly-Linked List"},
		RelatedChapters: []int{2, 5, 7, 12},
		Description: `Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:
- LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
- int get(int key) Return the value of the key if the key exists, otherwise return -1.
- void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.

The functions get and put must each run in O(1) average time complexity.`,
		Constraints: []string{
			"1 <= capacity <= 3000",
			"0 <= key <= 10^4",
			"0 <= value <= 10^5",
			"At most 2 * 10^5 calls will be made to get and put",
		},
		Examples: []Example{
			{Input: `["LRUCache","put","put","get","put","get","put","get","get","get"]
[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]`, Output: "[null,null,null,1,null,-1,null,-1,3,4]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}, "values": [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}}, Expected: []interface{}{nil, nil, nil, 1, nil, -1, nil, -1, 3, 4}},
		},
		TimeComplexity:  "O(1)",
		SpaceComplexity: "O(capacity)",
		StarterCode: `class LRUCache:
    def __init__(self, capacity: int):
        pass

    def get(self, key: int) -> int:
        pass

    def put(self, key: int, value: int) -> None:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use hash map + doubly linked list. Map gives O(1) lookup. DLL gives O(1) remove/add at ends.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "DLL maintains usage order (MRU at head, LRU at tail). On access, move node to head. On capacity overflow, remove tail."},
			{Level: 3, Type: "code", Content: "Use dummy head and tail nodes. Map stores key -> node. get: move to head. put: add/update and move to head, evict if needed."},
		},
		Solution: Solution{
			Code: `class Node:
    def __init__(self, key=0, val=0):
        self.key = key
        self.val = val
        self.prev = None
        self.next = None

class LRUCache:
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.cache = {}  # key -> Node
        # Dummy head and tail
        self.head = Node()
        self.tail = Node()
        self.head.next = self.tail
        self.tail.prev = self.head

    def _remove(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev

    def _add_to_head(self, node):
        node.next = self.head.next
        node.prev = self.head
        self.head.next.prev = node
        self.head.next = node

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        node = self.cache[key]
        self._remove(node)
        self._add_to_head(node)
        return node.val

    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            node = self.cache[key]
            node.val = value
            self._remove(node)
            self._add_to_head(node)
        else:
            node = Node(key, value)
            self.cache[key] = node
            self._add_to_head(node)
            if len(self.cache) > self.capacity:
                # Remove LRU (tail.prev)
                lru = self.tail.prev
                self._remove(lru)
                del self.cache[lru.key]`,
			Explanation:     "Hash map for O(1) lookup, doubly linked list for O(1) add/remove. MRU at head, LRU at tail. Move to head on access.",
			TimeComplexity:  "O(1) for both get and put",
			SpaceComplexity: "O(capacity)",
			Walkthrough: []WalkthroughStep{
				{Title: "Data structures", Explanation: "HashMap + Doubly linked list with dummy nodes", CodeSnippet: "self.cache = {}\nself.head = Node()\nself.tail = Node()", LineStart: 11, LineEnd: 14},
				{Title: "Move to head on access", Explanation: "Remove from current position, add to head", CodeSnippet: "self._remove(node)\nself._add_to_head(node)", LineStart: 31, LineEnd: 32},
				{Title: "Evict LRU", Explanation: "Remove tail.prev when over capacity", CodeSnippet: "lru = self.tail.prev\nself._remove(lru)\ndel self.cache[lru.key]", LineStart: 46, LineEnd: 49},
			},
		},
	},
	// Trees (continued)
	{
		ID:              "same-tree",
		Number:          59,
		Title:           "Same Tree",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "BFS", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11},
		Description: `Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.`,
		Constraints: []string{
			"The number of nodes in both trees is in the range [0, 100]",
			"-10^4 <= Node.val <= 10^4",
		},
		Examples: []Example{
			{Input: "p = [1,2,3], q = [1,2,3]", Output: "true"},
			{Input: "p = [1,2], q = [1,null,2]", Output: "false"},
			{Input: "p = [1,2,1], q = [1,1,2]", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"p": []interface{}{1, 2, 3}, "q": []interface{}{1, 2, 3}}, Expected: true},
			{Input: map[string]interface{}{"p": []interface{}{1, 2}, "q": []interface{}{1, nil, 2}}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def isSameTree(p, q):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Recursively compare nodes: both null (same), one null (different), values differ (different), else check children."},
			{Level: 2, Type: "algorithm", Content: "Base cases: both None -> True, one None -> False, different values -> False. Recurse on left and right."},
			{Level: 3, Type: "code", Content: "if not p and not q: return True. if not p or not q: return False. return p.val == q.val and recurse."},
		},
		Solution: Solution{
			Code: `def isSameTree(p, q):
    if not p and not q:
        return True
    if not p or not q:
        return False
    if p.val != q.val:
        return False
    return isSameTree(p.left, q.left) and isSameTree(p.right, q.right)`,
			Explanation:     "Recursively compare: both null = same, one null = different, different values = different, else check both children.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(h) - recursion stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Both null", Explanation: "Empty trees are the same", CodeSnippet: "if not p and not q:\n    return True", LineStart: 2, LineEnd: 3},
				{Title: "One null", Explanation: "Structure differs", CodeSnippet: "if not p or not q:\n    return False", LineStart: 4, LineEnd: 5},
				{Title: "Check children", Explanation: "Both subtrees must match", CodeSnippet: "return isSameTree(p.left, q.left) and isSameTree(p.right, q.right)", LineStart: 8, LineEnd: 8},
			},
		},
	},
	{
		ID:              "subtree-of-another-tree",
		Number:          60,
		Title:           "Subtree of Another Tree",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "String Matching", "Binary Tree", "Hash Function"},
		RelatedChapters: []int{9, 10, 11},
		Description: `Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.`,
		Constraints: []string{
			"The number of nodes in the root tree is in the range [1, 2000]",
			"The number of nodes in the subRoot tree is in the range [1, 1000]",
			"-10^4 <= root.val <= 10^4",
			"-10^4 <= subRoot.val <= 10^4",
		},
		Examples: []Example{
			{Input: "root = [3,4,5,1,2], subRoot = [4,1,2]", Output: "true"},
			{Input: "root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 4, 5, 1, 2}, "subRoot": []interface{}{4, 1, 2}}, Expected: true},
			{Input: map[string]interface{}{"root": []interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0}, "subRoot": []interface{}{4, 1, 2}}, Expected: false},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def isSubtree(root, subRoot):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "For each node in root, check if subtree starting there equals subRoot. Use isSameTree helper."},
			{Level: 2, Type: "algorithm", Content: "Base: if root null, return false (unless subRoot also null). Check if trees are same, or recurse on children."},
			{Level: 3, Type: "code", Content: "def isSubtree(root, sub): return isSameTree(root, sub) or isSubtree(root.left, sub) or isSubtree(root.right, sub)."},
		},
		Solution: Solution{
			Code: `def isSubtree(root, subRoot):
    def isSameTree(p, q):
        if not p and not q:
            return True
        if not p or not q:
            return False
        return p.val == q.val and isSameTree(p.left, q.left) and isSameTree(p.right, q.right)

    if not root:
        return False
    if isSameTree(root, subRoot):
        return True
    return isSubtree(root.left, subRoot) or isSubtree(root.right, subRoot)`,
			Explanation:     "At each node, check if tree matches subRoot. If not, recursively check left and right children.",
			TimeComplexity:  "O(m * n) - for each of m nodes, potentially compare n nodes",
			SpaceComplexity: "O(h) - recursion depth",
			Walkthrough: []WalkthroughStep{
				{Title: "isSameTree helper", Explanation: "Checks if two trees are identical", CodeSnippet: "def isSameTree(p, q):", LineStart: 2, LineEnd: 7},
				{Title: "Check current node", Explanation: "See if subtree rooted here matches", CodeSnippet: "if isSameTree(root, subRoot):\n    return True", LineStart: 11, LineEnd: 12},
				{Title: "Recurse on children", Explanation: "Check if subRoot exists in left or right subtree", CodeSnippet: "return isSubtree(root.left, subRoot) or isSubtree(root.right, subRoot)", LineStart: 13, LineEnd: 13},
			},
		},
	},
	{
		ID:              "lowest-common-ancestor-bst",
		Number:          61,
		Title:           "Lowest Common Ancestor of a BST",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Search Tree", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11},
		Description: `Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself)."`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [2, 10^5]",
			"-10^9 <= Node.val <= 10^9",
			"All Node.val are unique",
			"p != q",
			"p and q will exist in the BST",
		},
		Examples: []Example{
			{Input: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8", Output: "6"},
			{Input: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4", Output: "2"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, "p": 2, "q": 8}, Expected: 6},
			{Input: map[string]interface{}{"root": []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, "p": 2, "q": 4}, Expected: 2},
		},
		TimeComplexity:  "O(h)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def lowestCommonAncestor(root, p, q):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use BST property: if both p and q are smaller, go left. If both larger, go right. Otherwise, current is LCA."},
			{Level: 2, Type: "algorithm", Content: "LCA is where p and q split (one goes left, one goes right, or one equals current node)."},
			{Level: 3, Type: "code", Content: "while root: if p.val < root.val and q.val < root.val: go left. elif both > root.val: go right. else: return root."},
		},
		Solution: Solution{
			Code: `def lowestCommonAncestor(root, p, q):
    while root:
        if p.val < root.val and q.val < root.val:
            root = root.left
        elif p.val > root.val and q.val > root.val:
            root = root.right
        else:
            return root`,
			Explanation:     "Use BST property. If both values smaller, go left. Both larger, go right. Otherwise, we're at the LCA (split point).",
			TimeComplexity:  "O(h) - height of tree",
			SpaceComplexity: "O(1) - iterative",
			Walkthrough: []WalkthroughStep{
				{Title: "Both smaller", Explanation: "p and q both in left subtree", CodeSnippet: "if p.val < root.val and q.val < root.val:\n    root = root.left", LineStart: 3, LineEnd: 4},
				{Title: "Both larger", Explanation: "p and q both in right subtree", CodeSnippet: "elif p.val > root.val and q.val > root.val:\n    root = root.right", LineStart: 5, LineEnd: 6},
				{Title: "Split point", Explanation: "p and q are on different sides (or one equals root)", CodeSnippet: "else:\n    return root", LineStart: 7, LineEnd: 8},
			},
		},
	},
	{
		ID:              "binary-tree-level-order-traversal",
		Number:          62,
		Title:           "Binary Tree Level Order Traversal",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "BFS", "Binary Tree"},
		RelatedChapters: []int{6, 9, 10, 11},
		Description: `Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [0, 2000]",
			"-1000 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "root = [3,9,20,null,null,15,7]", Output: "[[3],[9,20],[15,7]]"},
			{Input: "root = [1]", Output: "[[1]]"},
			{Input: "root = []", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 9, 20, nil, nil, 15, 7}}, Expected: [][]int{{3}, {9, 20}, {15, 7}}},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: [][]int{{1}}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def levelOrder(root):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use BFS with a queue. Process all nodes at current level before moving to next level.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Track level size. Process exactly that many nodes, adding their children. All processed nodes form one level."},
			{Level: 3, Type: "code", Content: "while queue: level_size = len(queue). for i in range(level_size): pop, add to level list, push children."},
		},
		Solution: Solution{
			Code: `from collections import deque

def levelOrder(root):
    if not root:
        return []

    result = []
    queue = deque([root])

    while queue:
        level = []
        level_size = len(queue)

        for _ in range(level_size):
            node = queue.popleft()
            level.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)

        result.append(level)

    return result`,
			Explanation:     "BFS with level tracking. Process exactly level_size nodes per iteration, adding children for next level.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n) - queue can hold entire level",
			Walkthrough: []WalkthroughStep{
				{Title: "Track level size", Explanation: "Know how many nodes belong to current level", CodeSnippet: "level_size = len(queue)", LineStart: 12, LineEnd: 12},
				{Title: "Process level", Explanation: "Pop exactly level_size nodes", CodeSnippet: "for _ in range(level_size):\n    node = queue.popleft()", LineStart: 14, LineEnd: 15},
				{Title: "Add children", Explanation: "Children form next level", CodeSnippet: "if node.left:\n    queue.append(node.left)", LineStart: 17, LineEnd: 20},
			},
		},
	},
	{
		ID:              "validate-binary-search-tree",
		Number:          63,
		Title:           "Validate Binary Search Tree",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Search Tree", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11},
		Description: `Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [1, 10^4]",
			"-2^31 <= Node.val <= 2^31 - 1",
		},
		Examples: []Example{
			{Input: "root = [2,1,3]", Output: "true"},
			{Input: "root = [5,1,4,null,null,3,6]", Output: "false", Explanation: "Root is 5 but right child is 4."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{2, 1, 3}}, Expected: true},
			{Input: map[string]interface{}{"root": []interface{}{5, 1, 4, nil, nil, 3, 6}}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def isValidBST(root):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Track valid range [min, max] for each node. Going left updates max, going right updates min."},
			{Level: 2, Type: "algorithm", Content: "Initially range is (-inf, +inf). Left child: (min, parent). Right child: (parent, max). Node must be in range."},
			{Level: 3, Type: "code", Content: "def valid(node, min_val, max_val): if not node: return True. check range, recurse with updated bounds."},
		},
		Solution: Solution{
			Code: `def isValidBST(root):
    def valid(node, min_val, max_val):
        if not node:
            return True
        if node.val <= min_val or node.val >= max_val:
            return False
        return valid(node.left, min_val, node.val) and valid(node.right, node.val, max_val)

    return valid(root, float('-inf'), float('inf'))`,
			Explanation:     "Track valid range for each node. Going left: upper bound becomes parent. Going right: lower bound becomes parent.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(h) - recursion depth",
			Walkthrough: []WalkthroughStep{
				{Title: "Base case", Explanation: "Empty node is valid", CodeSnippet: "if not node:\n    return True", LineStart: 3, LineEnd: 4},
				{Title: "Check range", Explanation: "Node value must be in (min, max)", CodeSnippet: "if node.val <= min_val or node.val >= max_val:\n    return False", LineStart: 5, LineEnd: 6},
				{Title: "Recurse with bounds", Explanation: "Left: update max. Right: update min", CodeSnippet: "return valid(node.left, min_val, node.val) and valid(node.right, node.val, max_val)", LineStart: 7, LineEnd: 7},
			},
		},
	},
	{
		ID:              "kth-smallest-element-in-bst",
		Number:          64,
		Title:           "Kth Smallest Element in a BST",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Search Tree", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11},
		Description: `Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.`,
		Constraints: []string{
			"The number of nodes in the tree is n",
			"1 <= k <= n <= 10^4",
			"0 <= Node.val <= 10^4",
		},
		Examples: []Example{
			{Input: "root = [3,1,4,null,2], k = 1", Output: "1"},
			{Input: "root = [5,3,6,2,4,null,null,1], k = 3", Output: "3"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 1, 4, nil, 2}, "k": 1}, Expected: 1},
			{Input: map[string]interface{}{"root": []interface{}{5, 3, 6, 2, 4, nil, nil, 1}, "k": 3}, Expected: 3},
		},
		TimeComplexity:  "O(H + k)",
		SpaceComplexity: "O(H)",
		StarterCode:     "def kthSmallest(root, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Inorder traversal of BST visits nodes in sorted order. Stop at kth node."},
			{Level: 2, Type: "algorithm", Content: "Use iterative inorder with stack. Push all left children, pop, count, push right."},
			{Level: 3, Type: "code", Content: "stack = []. Go left as far as possible. Pop, decrement k. If k == 0, return. Go right."},
		},
		Solution: Solution{
			Code: `def kthSmallest(root, k):
    stack = []
    curr = root

    while stack or curr:
        # Go left as far as possible
        while curr:
            stack.append(curr)
            curr = curr.left

        curr = stack.pop()
        k -= 1
        if k == 0:
            return curr.val

        curr = curr.right

    return -1`,
			Explanation:     "Iterative inorder traversal. BST inorder gives sorted values. Stop when we've visited k nodes.",
			TimeComplexity:  "O(H + k) - go down H levels, then visit k nodes",
			SpaceComplexity: "O(H) for stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Go left", Explanation: "Push all left children to stack", CodeSnippet: "while curr:\n    stack.append(curr)\n    curr = curr.left", LineStart: 7, LineEnd: 9},
				{Title: "Visit node", Explanation: "Pop and count", CodeSnippet: "curr = stack.pop()\nk -= 1", LineStart: 11, LineEnd: 12},
				{Title: "Check k", Explanation: "Return when kth node visited", CodeSnippet: "if k == 0:\n    return curr.val", LineStart: 13, LineEnd: 14},
			},
		},
	},
	{
		ID:              "construct-binary-tree-from-preorder-inorder",
		Number:          65,
		Title:           "Construct Binary Tree from Preorder and Inorder Traversal",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Array", "Hash Table", "Divide and Conquer", "Tree", "Binary Tree"},
		RelatedChapters: []int{4, 9, 10, 11},
		Description: `Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.`,
		Constraints: []string{
			"1 <= preorder.length <= 3000",
			"inorder.length == preorder.length",
			"-3000 <= preorder[i], inorder[i] <= 3000",
			"preorder and inorder consist of unique values",
			"Each value of inorder also appears in preorder",
			"preorder is guaranteed to be the preorder traversal of the tree",
			"inorder is guaranteed to be the inorder traversal of the tree",
		},
		Examples: []Example{
			{Input: "preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]", Output: "[3,9,20,null,null,15,7]"},
			{Input: "preorder = [-1], inorder = [-1]", Output: "[-1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"preorder": []int{3, 9, 20, 15, 7}, "inorder": []int{9, 3, 15, 20, 7}}, Expected: []interface{}{3, 9, 20, nil, nil, 15, 7}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def buildTree(preorder, inorder):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "First element of preorder is root. Find root in inorder - elements to left are left subtree, to right are right subtree.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "Use hash map for O(1) lookup in inorder. Recursively build left subtree then right subtree."},
			{Level: 3, Type: "code", Content: "Build index map for inorder. Pop from preorder for root. Partition inorder by root index. Recurse on left, then right."},
		},
		Solution: Solution{
			Code: `def buildTree(preorder, inorder):
    inorder_map = {val: idx for idx, val in enumerate(inorder)}
    pre_idx = [0]  # Use list for mutable reference

    def build(left, right):
        if left > right:
            return None

        root_val = preorder[pre_idx[0]]
        pre_idx[0] += 1
        root = TreeNode(root_val)

        mid = inorder_map[root_val]
        root.left = build(left, mid - 1)
        root.right = build(mid + 1, right)

        return root

    return build(0, len(inorder) - 1)`,
			Explanation:     "Preorder gives roots. Inorder partitions into left/right subtrees. Hash map speeds up finding root position.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n) for hash map and recursion",
			Walkthrough: []WalkthroughStep{
				{Title: "Build index map", Explanation: "O(1) lookup for root position in inorder", CodeSnippet: "inorder_map = {val: idx for idx, val in enumerate(inorder)}", LineStart: 2, LineEnd: 2},
				{Title: "Get root from preorder", Explanation: "First unprocessed preorder element is current root", CodeSnippet: "root_val = preorder[pre_idx[0]]\npre_idx[0] += 1", LineStart: 9, LineEnd: 10},
				{Title: "Partition and recurse", Explanation: "Build left subtree before right (preorder order)", CodeSnippet: "root.left = build(left, mid - 1)\nroot.right = build(mid + 1, right)", LineStart: 14, LineEnd: 15},
			},
		},
	},
	{
		ID:              "binary-tree-max-path-sum",
		Number:          66,
		Title:           "Binary Tree Maximum Path Sum",
		Difficulty:      "Hard",
		Category:        "trees",
		Tags:            []string{"Dynamic Programming", "Tree", "DFS", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11, 12},
		Description: `A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path.`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [1, 3 * 10^4]",
			"-1000 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "root = [1,2,3]", Output: "6", Explanation: "Optimal path is 2 -> 1 -> 3 with sum 6."},
			{Input: "root = [-10,9,20,null,null,15,7]", Output: "42", Explanation: "Optimal path is 15 -> 20 -> 7 with sum 42."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3}}, Expected: 6},
			{Input: map[string]interface{}{"root": []interface{}{-10, 9, 20, nil, nil, 15, 7}}, Expected: 42},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def maxPathSum(root):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "At each node, consider: path through node (left + node + right) vs. path continuing upward (node + max(left, right)).", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "DFS returns max gain if path continues upward. Track global max for paths that end at current node (including both children)."},
			{Level: 3, Type: "code", Content: "def dfs(node): returns max(0, node.val + max(left, right)). Update global max with node.val + left + right."},
		},
		Solution: Solution{
			Code: `def maxPathSum(root):
    max_sum = [float('-inf')]

    def dfs(node):
        if not node:
            return 0

        # Max gain from left/right children (ignore if negative)
        left_gain = max(0, dfs(node.left))
        right_gain = max(0, dfs(node.right))

        # Path through current node (using both children)
        path_sum = node.val + left_gain + right_gain
        max_sum[0] = max(max_sum[0], path_sum)

        # Return max gain for path continuing upward
        return node.val + max(left_gain, right_gain)

    dfs(root)
    return max_sum[0]`,
			Explanation:     "DFS computes max path continuing upward. At each node, also check if path through node (both children) is best. Track global max.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(h) - recursion depth",
			Walkthrough: []WalkthroughStep{
				{Title: "Ignore negative paths", Explanation: "Only use child paths if they add positive value", CodeSnippet: "left_gain = max(0, dfs(node.left))", LineStart: 9, LineEnd: 10},
				{Title: "Check path through node", Explanation: "Path using both children ends here", CodeSnippet: "path_sum = node.val + left_gain + right_gain", LineStart: 13, LineEnd: 14},
				{Title: "Return upward gain", Explanation: "Can only use one child if path continues up", CodeSnippet: "return node.val + max(left_gain, right_gain)", LineStart: 17, LineEnd: 17},
			},
		},
	},
	{
		ID:              "serialize-and-deserialize-binary-tree",
		Number:          67,
		Title:           "Serialize and Deserialize Binary Tree",
		Difficulty:      "Hard",
		Category:        "trees",
		Tags:            []string{"String", "Tree", "DFS", "BFS", "Design", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11, 12},
		Description: `Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later.

Design an algorithm to serialize and deserialize a binary tree.`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [0, 10^4]",
			"-1000 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "root = [1,2,3,null,null,4,5]", Output: "[1,2,3,null,null,4,5]"},
			{Input: "root = []", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, nil, nil, 4, 5}}, Expected: []interface{}{1, 2, 3, nil, nil, 4, 5}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `class Codec:
    def serialize(self, root):
        # Write your solution here
        pass

    def deserialize(self, data):
        # Write your solution here
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use preorder traversal with null markers. Serialize: node value or 'N' for null. Deserialize: recursively build from string.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "Preorder with nulls: '1,2,N,N,3,4,N,N,5,N,N'. Deserialize by reading values in order and recursively building."},
			{Level: 3, Type: "code", Content: "serialize: dfs, append val or 'N'. deserialize: use iterator, pop value, if 'N' return None, else create node and recurse."},
		},
		Solution: Solution{
			Code: `class Codec:
    def serialize(self, root):
        result = []

        def dfs(node):
            if not node:
                result.append('N')
                return
            result.append(str(node.val))
            dfs(node.left)
            dfs(node.right)

        dfs(root)
        return ','.join(result)

    def deserialize(self, data):
        values = iter(data.split(','))

        def dfs():
            val = next(values)
            if val == 'N':
                return None
            node = TreeNode(int(val))
            node.left = dfs()
            node.right = dfs()
            return node

        return dfs()`,
			Explanation:     "Preorder traversal with null markers. Serialize to comma-separated string. Deserialize by consuming values in preorder.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Mark nulls", Explanation: "Use 'N' to mark empty nodes", CodeSnippet: "if not node:\n    result.append('N')", LineStart: 6, LineEnd: 8},
				{Title: "Preorder serialize", Explanation: "Root, then left, then right", CodeSnippet: "result.append(str(node.val))\ndfs(node.left)\ndfs(node.right)", LineStart: 9, LineEnd: 11},
				{Title: "Iterator deserialize", Explanation: "Consume values in same order", CodeSnippet: "val = next(values)\nif val == 'N': return None", LineStart: 20, LineEnd: 22},
			},
		},
	},
	// Tries
	{
		ID:              "implement-trie",
		Number:          68,
		Title:           "Implement Trie (Prefix Tree)",
		Difficulty:      "Medium",
		Category:        "tries",
		Tags:            []string{"Hash Table", "String", "Design", "Trie"},
		RelatedChapters: []int{5, 11, 12},
		Description: `A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings.

Implement the Trie class:
- Trie() Initializes the trie object.
- void insert(String word) Inserts the string word into the trie.
- boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
- boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.`,
		Constraints: []string{
			"1 <= word.length, prefix.length <= 2000",
			"word and prefix consist only of lowercase English letters",
			"At most 3 * 10^4 calls in total will be made to insert, search, and startsWith",
		},
		Examples: []Example{
			{Input: `["Trie","insert","search","search","startsWith","insert","search"]
[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]`, Output: "[null,null,true,false,true,null,true]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}, "values": [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}}, Expected: []interface{}{nil, nil, true, false, true, nil, true}},
		},
		TimeComplexity:  "O(m) per operation",
		SpaceComplexity: "O(n * m)",
		StarterCode: `class Trie:
    def __init__(self):
        pass

    def insert(self, word: str) -> None:
        pass

    def search(self, word: str) -> bool:
        pass

    def startsWith(self, prefix: str) -> bool:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Each node has children (dict/array) and an end-of-word flag. Traverse character by character.", ChapterRef: 11},
			{Level: 2, Type: "algorithm", Content: "Insert: create nodes as needed. Search: traverse and check end flag. StartsWith: traverse only, no end check."},
			{Level: 3, Type: "code", Content: "class TrieNode: children = {}, is_end = False. Insert: for char, create if not exists. Search: traverse, check is_end."},
		},
		Solution: Solution{
			Code: `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        node = self.root
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        node.is_end = True

    def search(self, word: str) -> bool:
        node = self._find_node(word)
        return node is not None and node.is_end

    def startsWith(self, prefix: str) -> bool:
        return self._find_node(prefix) is not None

    def _find_node(self, prefix: str):
        node = self.root
        for char in prefix:
            if char not in node.children:
                return None
            node = node.children[char]
        return node`,
			Explanation:     "Trie nodes have children dict and end flag. Insert creates path. Search/startsWith traverse, search also checks end flag.",
			TimeComplexity:  "O(m) where m is word/prefix length",
			SpaceComplexity: "O(n * m) for all inserted characters",
			Walkthrough: []WalkthroughStep{
				{Title: "TrieNode structure", Explanation: "Children dict and end-of-word flag", CodeSnippet: "self.children = {}\nself.is_end = False", LineStart: 3, LineEnd: 4},
				{Title: "Insert word", Explanation: "Create nodes along path, mark end", CodeSnippet: "if char not in node.children:\n    node.children[char] = TrieNode()\nnode.is_end = True", LineStart: 13, LineEnd: 16},
				{Title: "Search vs startsWith", Explanation: "Search checks is_end, startsWith doesn't", CodeSnippet: "return node is not None and node.is_end", LineStart: 20, LineEnd: 20},
			},
		},
	},
	{
		ID:              "design-add-and-search-words",
		Number:          69,
		Title:           "Design Add and Search Words Data Structure",
		Difficulty:      "Medium",
		Category:        "tries",
		Tags:            []string{"String", "DFS", "Design", "Trie"},
		RelatedChapters: []int{5, 11, 12},
		Description: `Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:
- WordDictionary() Initializes the object.
- void addWord(word) Adds word to the data structure, it can be matched later.
- bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.`,
		Constraints: []string{
			"1 <= word.length <= 25",
			"word in addWord consists of lowercase English letters",
			"word in search consist of '.' or lowercase English letters",
			"There will be at most 2 dots in word for search queries",
			"At most 10^4 calls will be made to addWord and search",
		},
		Examples: []Example{
			{Input: `["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]`, Output: "[null,null,null,null,false,true,true,true]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"}, "values": [][]string{{}, {"bad"}, {"dad"}, {"mad"}, {"pad"}, {"bad"}, {".ad"}, {"b.."}}}, Expected: []interface{}{nil, nil, nil, nil, false, true, true, true}},
		},
		TimeComplexity:  "O(m) add, O(26^d * m) search",
		SpaceComplexity: "O(n * m)",
		StarterCode: `class WordDictionary:
    def __init__(self):
        pass

    def addWord(self, word: str) -> None:
        pass

    def search(self, word: str) -> bool:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a Trie. For '.', try all children recursively instead of specific character.", ChapterRef: 11},
			{Level: 2, Type: "algorithm", Content: "addWord is standard Trie insert. search uses DFS: for '.', branch to all children; for letter, follow that child."},
			{Level: 3, Type: "code", Content: "def dfs(node, i): if i == len(word): return node.is_end. if word[i] == '.': try all children. else follow char."},
		},
		Solution: Solution{
			Code: `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class WordDictionary:
    def __init__(self):
        self.root = TrieNode()

    def addWord(self, word: str) -> None:
        node = self.root
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        node.is_end = True

    def search(self, word: str) -> bool:
        def dfs(node, i):
            if i == len(word):
                return node.is_end

            char = word[i]
            if char == '.':
                # Try all children
                for child in node.children.values():
                    if dfs(child, i + 1):
                        return True
                return False
            else:
                if char not in node.children:
                    return False
                return dfs(node.children[char], i + 1)

        return dfs(self.root, 0)`,
			Explanation:     "Trie with wildcard search. For '.', try all children (backtracking). For letters, follow specific path.",
			TimeComplexity:  "O(m) for add, O(26^d * m) for search with d dots",
			SpaceComplexity: "O(n * m) for Trie storage",
			Walkthrough: []WalkthroughStep{
				{Title: "Standard Trie add", Explanation: "Same as regular Trie insert", CodeSnippet: "for char in word:\n    if char not in node.children:", LineStart: 11, LineEnd: 14},
				{Title: "Handle wildcard", Explanation: "Try all children for '.'", CodeSnippet: "if char == '.':\n    for child in node.children.values():", LineStart: 24, LineEnd: 28},
				{Title: "Regular search", Explanation: "Follow specific character path", CodeSnippet: "return dfs(node.children[char], i + 1)", LineStart: 33, LineEnd: 33},
			},
		},
	},
	{
		ID:              "word-search-ii",
		Number:          70,
		Title:           "Word Search II",
		Difficulty:      "Hard",
		Category:        "tries",
		Tags:            []string{"Array", "String", "Backtracking", "Trie", "Matrix"},
		RelatedChapters: []int{4, 5, 11, 12},
		Description: `Given an m x n board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.`,
		Constraints: []string{
			"m == board.length",
			"n == board[i].length",
			"1 <= m, n <= 12",
			"board[i][j] is a lowercase English letter",
			"1 <= words.length <= 3 * 10^4",
			"1 <= words[i].length <= 10",
			"words[i] consists of lowercase English letters",
			"All the strings of words are unique",
		},
		Examples: []Example{
			{Input: `board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]`, Output: `["eat","oath"]`},
			{Input: `board = [["a","b"],["c","d"]], words = ["abcb"]`, Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"board": [][]string{{"o", "a", "a", "n"}, {"e", "t", "a", "e"}, {"i", "h", "k", "r"}, {"i", "f", "l", "v"}}, "words": []string{"oath", "pea", "eat", "rain"}}, Expected: []string{"eat", "oath"}},
		},
		TimeComplexity:  "O(m * n * 4^L)",
		SpaceComplexity: "O(N)",
		StarterCode:     "def findWords(board, words):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Build Trie from words. DFS from each cell, pruning paths not in Trie. More efficient than searching each word separately.", ChapterRef: 11},
			{Level: 2, Type: "algorithm", Content: "Insert all words into Trie. DFS from each cell following Trie. Mark visited cells. When reach end-of-word, add to result."},
			{Level: 3, Type: "code", Content: "Build Trie. For each cell, dfs with Trie node. If char not in children, return. If is_end, add word. Mark visited, explore 4 dirs, unmark."},
		},
		Solution: Solution{
			Code: `def findWords(board, words):
    class TrieNode:
        def __init__(self):
            self.children = {}
            self.word = None

    # Build Trie
    root = TrieNode()
    for word in words:
        node = root
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        node.word = word

    result = []
    m, n = len(board), len(board[0])

    def dfs(r, c, node):
        char = board[r][c]
        if char not in node.children:
            return

        next_node = node.children[char]
        if next_node.word:
            result.append(next_node.word)
            next_node.word = None  # Avoid duplicates

        # Mark visited
        board[r][c] = '#'

        # Explore 4 directions
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nr, nc = r + dr, c + dc
            if 0 <= nr < m and 0 <= nc < n and board[nr][nc] != '#':
                dfs(nr, nc, next_node)

        # Restore
        board[r][c] = char

    for r in range(m):
        for c in range(n):
            dfs(r, c, root)

    return result`,
			Explanation:     "Build Trie from words. DFS from each cell following Trie paths. Much faster than checking each word separately.",
			TimeComplexity:  "O(m * n * 4^L) where L is max word length",
			SpaceComplexity: "O(N) for Trie where N is total characters in words",
			Walkthrough: []WalkthroughStep{
				{Title: "Store word at end", Explanation: "Store full word instead of just flag for easy retrieval", CodeSnippet: "node.word = word", LineStart: 15, LineEnd: 15},
				{Title: "Prune with Trie", Explanation: "Stop DFS if path not in Trie", CodeSnippet: "if char not in node.children:\n    return", LineStart: 22, LineEnd: 23},
				{Title: "Mark and restore", Explanation: "Prevent revisiting same cell in path", CodeSnippet: "board[r][c] = '#'\n# explore...\nboard[r][c] = char", LineStart: 31, LineEnd: 40},
			},
		},
	},
	// Heap / Priority Queue
	{
		ID:              "kth-largest-element-in-stream",
		Number:          71,
		Title:           "Kth Largest Element in a Stream",
		Difficulty:      "Easy",
		Category:        "heap-priority-queue",
		Tags:            []string{"Tree", "Design", "Binary Search Tree", "Heap", "Binary Tree", "Data Stream"},
		RelatedChapters: []int{7, 11, 12},
		Description: `Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:
- KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
- int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.`,
		Constraints: []string{
			"1 <= k <= 10^4",
			"0 <= nums.length <= 10^4",
			"-10^4 <= nums[i] <= 10^4",
			"-10^4 <= val <= 10^4",
			"At most 10^4 calls will be made to add",
			"It is guaranteed that there will be at least k elements when add is called",
		},
		Examples: []Example{
			{Input: `["KthLargest","add","add","add","add","add"]
[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]`, Output: "[null,4,5,5,8,8]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"KthLargest", "add", "add", "add", "add", "add"}, "values": []interface{}{[]int{3, 4, 5, 8, 2}, 3, 5, 10, 9, 4}}, Expected: []interface{}{nil, 4, 5, 5, 8, 8}},
		},
		TimeComplexity:  "O(log k) per add",
		SpaceComplexity: "O(k)",
		StarterCode: `class KthLargest:
    def __init__(self, k: int, nums: List[int]):
        pass

    def add(self, val: int) -> int:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a min-heap of size k. The smallest element in heap is the kth largest overall.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Keep only k largest elements in min-heap. On add: push value, if size > k, pop smallest. Top is kth largest."},
			{Level: 3, Type: "code", Content: "heapq.heappush(heap, val). if len(heap) > k: heapq.heappop(heap). return heap[0]."},
		},
		Solution: Solution{
			Code: `import heapq

class KthLargest:
    def __init__(self, k: int, nums):
        self.k = k
        self.heap = []
        for num in nums:
            self.add(num)

    def add(self, val: int) -> int:
        heapq.heappush(self.heap, val)
        if len(self.heap) > self.k:
            heapq.heappop(self.heap)
        return self.heap[0]`,
			Explanation:     "Min-heap of size k keeps k largest elements. Smallest of these (heap top) is the kth largest.",
			TimeComplexity:  "O(log k) per add",
			SpaceComplexity: "O(k)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize with adds", Explanation: "Use add method to build initial heap", CodeSnippet: "for num in nums:\n    self.add(num)", LineStart: 7, LineEnd: 8},
				{Title: "Maintain size k", Explanation: "Pop smallest when over k elements", CodeSnippet: "if len(self.heap) > self.k:\n    heapq.heappop(self.heap)", LineStart: 12, LineEnd: 13},
				{Title: "Return kth largest", Explanation: "Min of k largest = kth largest", CodeSnippet: "return self.heap[0]", LineStart: 14, LineEnd: 14},
			},
		},
	},
	{
		ID:              "last-stone-weight",
		Number:          72,
		Title:           "Last Stone Weight",
		Difficulty:      "Easy",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Heap"},
		RelatedChapters: []int{7, 11},
		Description: `You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:
- If x == y, both stones are destroyed, and
- If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.

At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left, return 0.`,
		Constraints: []string{
			"1 <= stones.length <= 30",
			"1 <= stones[i] <= 1000",
		},
		Examples: []Example{
			{Input: "stones = [2,7,4,1,8,1]", Output: "1"},
			{Input: "stones = [1]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"stones": []int{2, 7, 4, 1, 8, 1}}, Expected: 1},
			{Input: map[string]interface{}{"stones": []int{1}}, Expected: 1},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def lastStoneWeight(stones):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use max-heap to efficiently get two heaviest stones. Python has min-heap, so negate values.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Pop two largest, if different, push difference back. Continue until 0 or 1 stone remains."},
			{Level: 3, Type: "code", Content: "heap = [-s for s in stones]. heapify. while len > 1: pop two, push -diff if nonzero. return -heap[0] or 0."},
		},
		Solution: Solution{
			Code: `import heapq

def lastStoneWeight(stones):
    # Use negative values for max-heap behavior
    heap = [-s for s in stones]
    heapq.heapify(heap)

    while len(heap) > 1:
        y = -heapq.heappop(heap)  # Heaviest
        x = -heapq.heappop(heap)  # Second heaviest
        if x != y:
            heapq.heappush(heap, -(y - x))

    return -heap[0] if heap else 0`,
			Explanation:     "Max-heap (via negation) for heaviest stones. Pop two, push difference if nonzero. Repeat until done.",
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Max-heap via negation", Explanation: "Python only has min-heap, negate for max", CodeSnippet: "heap = [-s for s in stones]", LineStart: 5, LineEnd: 5},
				{Title: "Smash stones", Explanation: "Pop two heaviest, push difference if nonzero", CodeSnippet: "if x != y:\n    heapq.heappush(heap, -(y - x))", LineStart: 11, LineEnd: 12},
				{Title: "Return result", Explanation: "Remaining stone weight or 0", CodeSnippet: "return -heap[0] if heap else 0", LineStart: 14, LineEnd: 14},
			},
		},
	},
	{
		ID:              "k-closest-points-to-origin",
		Number:          73,
		Title:           "K Closest Points to Origin",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Math", "Divide and Conquer", "Geometry", "Sorting", "Heap", "Quickselect"},
		RelatedChapters: []int{7, 11, 12},
		Description: `Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., sqrt((x1 - x2)^2 + (y1 - y2)^2)).

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).`,
		Constraints: []string{
			"1 <= k <= points.length <= 10^4",
			"-10^4 <= xi, yi <= 10^4",
		},
		Examples: []Example{
			{Input: "points = [[1,3],[-2,2]], k = 1", Output: "[[-2,2]]"},
			{Input: "points = [[3,3],[5,-1],[-2,4]], k = 2", Output: "[[3,3],[-2,4]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"points": [][]int{{1, 3}, {-2, 2}}, "k": 1}, Expected: [][]int{{-2, 2}}},
			{Input: map[string]interface{}{"points": [][]int{{3, 3}, {5, -1}, {-2, 4}}, "k": 2}, Expected: [][]int{{3, 3}, {-2, 4}}},
		},
		TimeComplexity:  "O(n log k)",
		SpaceComplexity: "O(k)",
		StarterCode:     "def kClosest(points, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use max-heap of size k. Compare by squared distance (no need for sqrt). Keep k smallest distances.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "For each point, push (-dist, point) to heap. If size > k, pop largest. Result is k closest points."},
			{Level: 3, Type: "code", Content: "dist = x*x + y*y. heappush with -dist for max-heap. if len > k: heappop. Return points from heap."},
		},
		Solution: Solution{
			Code: `import heapq

def kClosest(points, k):
    heap = []

    for x, y in points:
        dist = x * x + y * y
        heapq.heappush(heap, (-dist, [x, y]))
        if len(heap) > k:
            heapq.heappop(heap)

    return [point for _, point in heap]`,
			Explanation:     "Max-heap of size k (negated distance). Keeps k points with smallest distances. No sqrt needed for comparison.",
			TimeComplexity:  "O(n log k)",
			SpaceComplexity: "O(k)",
			Walkthrough: []WalkthroughStep{
				{Title: "Squared distance", Explanation: "Skip sqrt, relative order unchanged", CodeSnippet: "dist = x * x + y * y", LineStart: 7, LineEnd: 7},
				{Title: "Max-heap for k smallest", Explanation: "Negate distance, pop largest when over k", CodeSnippet: "heapq.heappush(heap, (-dist, [x, y]))\nif len(heap) > k:", LineStart: 8, LineEnd: 10},
				{Title: "Extract points", Explanation: "Return points from heap", CodeSnippet: "return [point for _, point in heap]", LineStart: 12, LineEnd: 12},
			},
		},
	},
	{
		ID:              "task-scheduler",
		Number:          74,
		Title:           "Task Scheduler",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Hash Table", "Greedy", "Sorting", "Heap", "Counting"},
		RelatedChapters: []int{7, 8, 11, 12},
		Description: `You are given an array of CPU tasks, each represented by letters A to Z, and a cooling interval n. Each cycle or interval allows the completion of one task. Tasks can be completed in any order, but there's a constraint: identical tasks must be separated by at least n intervals due to cooling time.

Return the minimum number of intervals required to complete all tasks.`,
		Constraints: []string{
			"1 <= tasks.length <= 10^4",
			"tasks[i] is an uppercase English letter",
			"0 <= n <= 100",
		},
		Examples: []Example{
			{Input: `tasks = ["A","A","A","B","B","B"], n = 2`, Output: "8", Explanation: "A -> B -> idle -> A -> B -> idle -> A -> B"},
			{Input: `tasks = ["A","A","A","B","B","B"], n = 0`, Output: "6"},
			{Input: `tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2`, Output: "16"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"tasks": []string{"A", "A", "A", "B", "B", "B"}, "n": 2}, Expected: 8},
			{Input: map[string]interface{}{"tasks": []string{"A", "A", "A", "B", "B", "B"}, "n": 0}, Expected: 6},
		},
		TimeComplexity:  "O(n * m)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def leastInterval(tasks, n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Max-heap for task counts. Each cycle: pop up to n+1 tasks, decrement, track time for cooldown.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Greedy: always pick task with most remaining. After cooldown, re-add tasks with remaining count."},
			{Level: 3, Type: "code", Content: "heap with -count. Each round: pop n+1 tasks, add to cooldown queue. After n+1 cycles, re-add from queue."},
		},
		Solution: Solution{
			Code: `from collections import Counter, deque
import heapq

def leastInterval(tasks, n):
    count = Counter(tasks)
    heap = [-cnt for cnt in count.values()]
    heapq.heapify(heap)

    time = 0
    queue = deque()  # (time_available, count)

    while heap or queue:
        time += 1

        if heap:
            cnt = heapq.heappop(heap) + 1  # Decrement (negative)
            if cnt != 0:
                queue.append((time + n, cnt))

        if queue and queue[0][0] == time:
            heapq.heappush(heap, queue.popleft()[1])

    return time`,
			Explanation:     "Max-heap tracks remaining counts. Queue tracks cooling tasks. Each tick: run a task, re-add after cooldown.",
			TimeComplexity:  "O(total_time * log 26) = O(n)",
			SpaceComplexity: "O(26) = O(1) for task counts",
			Walkthrough: []WalkthroughStep{
				{Title: "Count frequencies", Explanation: "Max-heap of task counts", CodeSnippet: "heap = [-cnt for cnt in count.values()]", LineStart: 6, LineEnd: 6},
				{Title: "Process each tick", Explanation: "Decrement count, add to cooldown queue", CodeSnippet: "if cnt != 0:\n    queue.append((time + n, cnt))", LineStart: 17, LineEnd: 18},
				{Title: "Re-add after cooldown", Explanation: "Task available again", CodeSnippet: "if queue and queue[0][0] == time:\n    heapq.heappush(heap, queue.popleft()[1])", LineStart: 20, LineEnd: 21},
			},
		},
	},
	{
		ID:              "find-median-from-data-stream",
		Number:          75,
		Title:           "Find Median from Data Stream",
		Difficulty:      "Hard",
		Category:        "heap-priority-queue",
		Tags:            []string{"Two Pointers", "Design", "Sorting", "Heap", "Data Stream"},
		RelatedChapters: []int{7, 11, 12},
		Description: `The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

Implement the MedianFinder class:
- MedianFinder() initializes the MedianFinder object.
- void addNum(int num) adds the integer num from the data stream to the data structure.
- double findMedian() returns the median of all elements so far.`,
		Constraints: []string{
			"-10^5 <= num <= 10^5",
			"There will be at least one element in the data structure before calling findMedian",
			"At most 5 * 10^4 calls will be made to addNum and findMedian",
		},
		Examples: []Example{
			{Input: `["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]`, Output: "[null,null,null,1.5,null,2.0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"}, "values": [][]int{{}, {1}, {2}, {}, {3}, {}}}, Expected: []interface{}{nil, nil, nil, 1.5, nil, 2.0}},
		},
		TimeComplexity:  "O(log n) add, O(1) median",
		SpaceComplexity: "O(n)",
		StarterCode: `class MedianFinder:
    def __init__(self):
        pass

    def addNum(self, num: int) -> None:
        pass

    def findMedian(self) -> float:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Two heaps: max-heap for smaller half, min-heap for larger half. Median is from top of heaps.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "Keep heaps balanced (size differs by at most 1). Add to small, move max to large, rebalance if needed."},
			{Level: 3, Type: "code", Content: "small = max-heap, large = min-heap. Add to small, push -heappop(small) to large. If large bigger, push -heappop(large) to small."},
		},
		Solution: Solution{
			Code: `import heapq

class MedianFinder:
    def __init__(self):
        self.small = []  # Max-heap (negated)
        self.large = []  # Min-heap

    def addNum(self, num: int) -> None:
        # Add to max-heap (small)
        heapq.heappush(self.small, -num)

        # Move largest from small to large
        heapq.heappush(self.large, -heapq.heappop(self.small))

        # Balance: small should have >= elements than large
        if len(self.large) > len(self.small):
            heapq.heappush(self.small, -heapq.heappop(self.large))

    def findMedian(self) -> float:
        if len(self.small) > len(self.large):
            return -self.small[0]
        return (-self.small[0] + self.large[0]) / 2`,
			Explanation:     "Two heaps partition data. Small (max-heap) has smaller half, large (min-heap) has larger half. Median from tops.",
			TimeComplexity:  "O(log n) for addNum, O(1) for findMedian",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Two heaps", Explanation: "Max-heap for small half, min-heap for large half", CodeSnippet: "self.small = []  # Max-heap (negated)\nself.large = []  # Min-heap", LineStart: 5, LineEnd: 6},
				{Title: "Add and rebalance", Explanation: "Always add to small, move to large, rebalance", CodeSnippet: "heapq.heappush(self.small, -num)\nheapq.heappush(self.large, -heapq.heappop(self.small))", LineStart: 10, LineEnd: 13},
				{Title: "Get median", Explanation: "Middle from small's max or average of both tops", CodeSnippet: "return (-self.small[0] + self.large[0]) / 2", LineStart: 22, LineEnd: 22},
			},
		},
	},
	// Backtracking (continued)
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
	// Graphs (continued)
	{
		ID:              "clone-graph",
		Number:          79,
		Title:           "Clone Graph",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Hash Table", "DFS", "BFS", "Graph"},
		RelatedChapters: []int{6, 10},
		Description: `Given a reference of a node in a connected undirected graph, return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.`,
		Constraints: []string{
			"The number of nodes in the graph is in the range [0, 100]",
			"1 <= Node.val <= 100",
			"Node.val is unique for each node",
			"There are no repeated edges and no self-loops in the graph",
			"The Graph is connected and all nodes can be visited starting from the given node",
		},
		Examples: []Example{
			{Input: "adjList = [[2,4],[1,3],[2,4],[1,3]]", Output: "[[2,4],[1,3],[2,4],[1,3]]"},
			{Input: "adjList = [[]]", Output: "[[]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"adjList": [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}}, Expected: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}},
		},
		TimeComplexity:  "O(V + E)",
		SpaceComplexity: "O(V)",
		StarterCode:     "def cloneGraph(node):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use hash map to track original->clone mapping. DFS/BFS to traverse. Create clones and connect neighbors.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "DFS: if node already cloned, return clone. Create new node, add to map, recursively clone neighbors."},
			{Level: 3, Type: "code", Content: "def dfs(node): if node in old_to_new: return it. Create clone, add to map. For each neighbor, append dfs(neighbor) to clone.neighbors."},
		},
		Solution: Solution{
			Code: `def cloneGraph(node):
    if not node:
        return None

    old_to_new = {}

    def dfs(node):
        if node in old_to_new:
            return old_to_new[node]

        clone = Node(node.val)
        old_to_new[node] = clone

        for neighbor in node.neighbors:
            clone.neighbors.append(dfs(neighbor))

        return clone

    return dfs(node)`,
			Explanation:     "DFS with memoization. Map original nodes to clones. Recursively clone neighbors.",
			TimeComplexity:  "O(V + E)",
			SpaceComplexity: "O(V) for hash map",
			Walkthrough: []WalkthroughStep{
				{Title: "Check if already cloned", Explanation: "Return existing clone to avoid infinite loop", CodeSnippet: "if node in old_to_new:\n    return old_to_new[node]", LineStart: 8, LineEnd: 9},
				{Title: "Create clone first", Explanation: "Add to map before recursing to handle cycles", CodeSnippet: "clone = Node(node.val)\nold_to_new[node] = clone", LineStart: 11, LineEnd: 12},
				{Title: "Clone neighbors", Explanation: "Recursively clone and connect", CodeSnippet: "clone.neighbors.append(dfs(neighbor))", LineStart: 15, LineEnd: 15},
			},
		},
	},
	{
		ID:              "pacific-atlantic-water-flow",
		Number:          80,
		Title:           "Pacific Atlantic Water Flow",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Array", "DFS", "BFS", "Matrix"},
		RelatedChapters: []int{6, 10},
		Description: `There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.`,
		Constraints: []string{
			"m == heights.length",
			"n == heights[r].length",
			"1 <= m, n <= 200",
			"0 <= heights[r][c] <= 10^5",
		},
		Examples: []Example{
			{Input: "heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]", Output: "[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"heights": [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}}, Expected: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def pacificAtlantic(heights):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Reverse the flow: DFS from ocean borders going uphill. Find cells reachable from both oceans.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "DFS from Pacific borders (top, left) and Atlantic borders (bottom, right). Return intersection of reachable cells."},
			{Level: 3, Type: "code", Content: "pacific, atlantic = set(), set(). DFS from borders with condition heights[nr][nc] >= heights[r][c]. Return intersection."},
		},
		Solution: Solution{
			Code: `def pacificAtlantic(heights):
    if not heights:
        return []

    m, n = len(heights), len(heights[0])
    pacific = set()
    atlantic = set()

    def dfs(r, c, reachable, prev_height):
        if (r, c) in reachable or r < 0 or r >= m or c < 0 or c >= n:
            return
        if heights[r][c] < prev_height:
            return

        reachable.add((r, c))
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            dfs(r + dr, c + dc, reachable, heights[r][c])

    # DFS from Pacific (top and left borders)
    for c in range(n):
        dfs(0, c, pacific, 0)
    for r in range(m):
        dfs(r, 0, pacific, 0)

    # DFS from Atlantic (bottom and right borders)
    for c in range(n):
        dfs(m - 1, c, atlantic, 0)
    for r in range(m):
        dfs(r, n - 1, atlantic, 0)

    return list(pacific & atlantic)`,
			Explanation:     "Reverse flow: DFS uphill from ocean borders. Cells in both sets can reach both oceans.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Reverse flow", Explanation: "Go uphill: heights[nr][nc] >= heights[r][c]", CodeSnippet: "if heights[r][c] < prev_height:\n    return", LineStart: 11, LineEnd: 12},
				{Title: "DFS from Pacific", Explanation: "Top row and left column", CodeSnippet: "for c in range(n):\n    dfs(0, c, pacific, 0)", LineStart: 20, LineEnd: 23},
				{Title: "Intersection", Explanation: "Cells reachable from both oceans", CodeSnippet: "return list(pacific & atlantic)", LineStart: 31, LineEnd: 31},
			},
		},
	},
	{
		ID:              "rotting-oranges",
		Number:          81,
		Title:           "Rotting Oranges",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Array", "BFS", "Matrix"},
		RelatedChapters: []int{6, 10},
		Description: `You are given an m x n grid where each cell can have one of three values:
- 0 representing an empty cell,
- 1 representing a fresh orange, or
- 2 representing a rotten orange.

Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.`,
		Constraints: []string{
			"m == grid.length",
			"n == grid[i].length",
			"1 <= m, n <= 10",
			"grid[i][j] is 0, 1, or 2",
		},
		Examples: []Example{
			{Input: "grid = [[2,1,1],[1,1,0],[0,1,1]]", Output: "4"},
			{Input: "grid = [[2,1,1],[0,1,1],[1,0,1]]", Output: "-1", Explanation: "Bottom left orange can never rot."},
			{Input: "grid = [[0,2]]", Output: "0"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"grid": [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}}, Expected: 4},
			{Input: map[string]interface{}{"grid": [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}}, Expected: -1},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def orangesRotting(grid):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Multi-source BFS from all rotten oranges simultaneously. Count fresh oranges. Track minutes.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Add all rotten to queue, count fresh. BFS level by level (each level = 1 minute). Decrement fresh count."},
			{Level: 3, Type: "code", Content: "queue = all rotten. while queue: process level, rot adjacent fresh. if fresh > 0: return -1. return minutes."},
		},
		Solution: Solution{
			Code: `from collections import deque

def orangesRotting(grid):
    m, n = len(grid), len(grid[0])
    queue = deque()
    fresh = 0

    # Find all rotten and count fresh
    for r in range(m):
        for c in range(n):
            if grid[r][c] == 2:
                queue.append((r, c))
            elif grid[r][c] == 1:
                fresh += 1

    if fresh == 0:
        return 0

    minutes = 0
    while queue:
        minutes += 1
        for _ in range(len(queue)):
            r, c = queue.popleft()
            for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                nr, nc = r + dr, c + dc
                if 0 <= nr < m and 0 <= nc < n and grid[nr][nc] == 1:
                    grid[nr][nc] = 2
                    fresh -= 1
                    queue.append((nr, nc))

    return minutes - 1 if fresh == 0 else -1`,
			Explanation:     "Multi-source BFS from all rotten oranges. Each BFS level is one minute. Check if all fresh oranges rotted.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n) for queue",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize", Explanation: "Queue all rotten, count fresh", CodeSnippet: "if grid[r][c] == 2:\n    queue.append((r, c))\nelif grid[r][c] == 1:\n    fresh += 1", LineStart: 11, LineEnd: 14},
				{Title: "BFS by level", Explanation: "Each level is one minute", CodeSnippet: "for _ in range(len(queue)):", LineStart: 22, LineEnd: 22},
				{Title: "Rot adjacent", Explanation: "Mark as rotten and add to queue", CodeSnippet: "grid[nr][nc] = 2\nfresh -= 1\nqueue.append((nr, nc))", LineStart: 27, LineEnd: 29},
			},
		},
	},
	{
		ID:              "course-schedule-ii",
		Number:          82,
		Title:           "Course Schedule II",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"DFS", "BFS", "Graph", "Topological Sort"},
		RelatedChapters: []int{6, 10},
		Description: `There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.`,
		Constraints: []string{
			"1 <= numCourses <= 2000",
			"0 <= prerequisites.length <= numCourses * (numCourses - 1)",
			"prerequisites[i].length == 2",
			"0 <= ai, bi < numCourses",
			"ai != bi",
			"All the pairs [ai, bi] are distinct",
		},
		Examples: []Example{
			{Input: "numCourses = 2, prerequisites = [[1,0]]", Output: "[0,1]"},
			{Input: "numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]", Output: "[0,2,1,3]"},
			{Input: "numCourses = 1, prerequisites = []", Output: "[0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"numCourses": 2, "prerequisites": [][]int{{1, 0}}}, Expected: []int{0, 1}},
			{Input: map[string]interface{}{"numCourses": 4, "prerequisites": [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, Expected: []int{0, 1, 2, 3}},
		},
		TimeComplexity:  "O(V + E)",
		SpaceComplexity: "O(V + E)",
		StarterCode:     "def findOrder(numCourses, prerequisites):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Topological sort using Kahn's algorithm (BFS) or DFS. Return empty array if cycle detected.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Kahn's: Track in-degrees. Add courses with in-degree 0 to queue. Process, reduce in-degrees, add new zeros."},
			{Level: 3, Type: "code", Content: "Build graph and in-degree array. Queue courses with in-degree 0. While queue: pop, add to result, decrement neighbors' in-degrees."},
		},
		Solution: Solution{
			Code: `from collections import deque, defaultdict

def findOrder(numCourses, prerequisites):
    graph = defaultdict(list)
    in_degree = [0] * numCourses

    # Build graph
    for course, prereq in prerequisites:
        graph[prereq].append(course)
        in_degree[course] += 1

    # Start with courses having no prerequisites
    queue = deque([i for i in range(numCourses) if in_degree[i] == 0])
    result = []

    while queue:
        course = queue.popleft()
        result.append(course)

        for next_course in graph[course]:
            in_degree[next_course] -= 1
            if in_degree[next_course] == 0:
                queue.append(next_course)

    return result if len(result) == numCourses else []`,
			Explanation:     "Kahn's algorithm for topological sort. Track in-degrees. Process nodes with in-degree 0. Return empty if cycle.",
			TimeComplexity:  "O(V + E)",
			SpaceComplexity: "O(V + E)",
			Walkthrough: []WalkthroughStep{
				{Title: "Build graph", Explanation: "Adjacency list and in-degree count", CodeSnippet: "graph[prereq].append(course)\nin_degree[course] += 1", LineStart: 9, LineEnd: 10},
				{Title: "Start with zeros", Explanation: "Courses with no prerequisites", CodeSnippet: "queue = deque([i for i in range(numCourses) if in_degree[i] == 0])", LineStart: 13, LineEnd: 13},
				{Title: "Process and update", Explanation: "Add to result, decrement in-degrees of neighbors", CodeSnippet: "in_degree[next_course] -= 1\nif in_degree[next_course] == 0:", LineStart: 21, LineEnd: 23},
			},
		},
	},
	// 1D Dynamic Programming (continued)
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
	// Greedy
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
	// Intervals
	{
		ID:              "merge-intervals",
		Number:          90,
		Title:           "Merge Intervals",
		Difficulty:      "Medium",
		Category:        "intervals",
		Tags:            []string{"Array", "Sorting"},
		RelatedChapters: []int{8},
		Description: `Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.`,
		Constraints: []string{
			"1 <= intervals.length <= 10^4",
			"intervals[i].length == 2",
			"0 <= starti <= endi <= 10^4",
		},
		Examples: []Example{
			{Input: "intervals = [[1,3],[2,6],[8,10],[15,18]]", Output: "[[1,6],[8,10],[15,18]]"},
			{Input: "intervals = [[1,4],[4,5]]", Output: "[[1,5]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, Expected: [][]int{{1, 6}, {8, 10}, {15, 18}}},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 4}, {4, 5}}}, Expected: [][]int{{1, 5}}},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def merge(intervals):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Sort by start time. Merge consecutive overlapping intervals by extending end time.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "Sort intervals. For each interval: if overlaps with last in result, extend end. Else add as new interval."},
			{Level: 3, Type: "code", Content: "Sort by start. result = [intervals[0]]. for interval: if interval[0] <= result[-1][1]: merge. else append."},
		},
		Solution: Solution{
			Code: `def merge(intervals):
    intervals.sort(key=lambda x: x[0])
    result = [intervals[0]]

    for start, end in intervals[1:]:
        if start <= result[-1][1]:
            result[-1][1] = max(result[-1][1], end)
        else:
            result.append([start, end])

    return result`,
			Explanation:     "Sort by start. If current overlaps with last merged, extend end. Otherwise, add as new interval.",
			TimeComplexity:  "O(n log n) for sorting",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort by start", Explanation: "Process intervals in order", CodeSnippet: "intervals.sort(key=lambda x: x[0])", LineStart: 2, LineEnd: 2},
				{Title: "Check overlap", Explanation: "Overlaps if start <= previous end", CodeSnippet: "if start <= result[-1][1]:", LineStart: 6, LineEnd: 6},
				{Title: "Merge or add", Explanation: "Extend end or add new interval", CodeSnippet: "result[-1][1] = max(result[-1][1], end)", LineStart: 7, LineEnd: 7},
			},
		},
	},
	{
		ID:              "non-overlapping-intervals",
		Number:          91,
		Title:           "Non-overlapping Intervals",
		Difficulty:      "Medium",
		Category:        "intervals",
		Tags:            []string{"Array", "Dynamic Programming", "Greedy", "Sorting"},
		RelatedChapters: []int{8},
		Description: `Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.`,
		Constraints: []string{
			"1 <= intervals.length <= 10^5",
			"intervals[i].length == 2",
			"-5 * 10^4 <= starti < endi <= 5 * 10^4",
		},
		Examples: []Example{
			{Input: "intervals = [[1,2],[2,3],[3,4],[1,3]]", Output: "1", Explanation: "Remove [1,3]."},
			{Input: "intervals = [[1,2],[1,2],[1,2]]", Output: "2"},
			{Input: "intervals = [[1,2],[2,3]]", Output: "0"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}}, Expected: 1},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {1, 2}, {1, 2}}}, Expected: 2},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def eraseOverlapIntervals(intervals):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Greedy: sort by end time. Keep intervals that end earliest. Count removals when overlap detected.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "Sort by end. Track prev_end. If start < prev_end: overlap, remove one (increment count). Else update prev_end."},
			{Level: 3, Type: "code", Content: "Sort by end. prev_end = intervals[0][1], count = 0. for interval: if start < prev_end: count++. else: prev_end = end."},
		},
		Solution: Solution{
			Code: `def eraseOverlapIntervals(intervals):
    intervals.sort(key=lambda x: x[1])  # Sort by end time
    count = 0
    prev_end = float('-inf')

    for start, end in intervals:
        if start >= prev_end:
            prev_end = end
        else:
            count += 1  # Remove this interval

    return count`,
			Explanation:     "Greedy: sort by end time. Keep intervals ending earliest. Count overlaps (which need removal).",
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort by end", Explanation: "Greedy choice: keep early-ending intervals", CodeSnippet: "intervals.sort(key=lambda x: x[1])", LineStart: 2, LineEnd: 2},
				{Title: "No overlap", Explanation: "Can keep this interval", CodeSnippet: "if start >= prev_end:\n    prev_end = end", LineStart: 7, LineEnd: 8},
				{Title: "Overlap", Explanation: "Must remove this interval", CodeSnippet: "count += 1", LineStart: 10, LineEnd: 10},
			},
		},
	},
	// Math & Geometry
	{
		ID:              "rotate-image",
		Number:          92,
		Title:           "Rotate Image",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Math", "Matrix"},
		RelatedChapters: []int{1},
		Description: `You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.`,
		Constraints: []string{
			"n == matrix.length == matrix[i].length",
			"1 <= n <= 20",
			"-1000 <= matrix[i][j] <= 1000",
		},
		Examples: []Example{
			{Input: "matrix = [[1,2,3],[4,5,6],[7,8,9]]", Output: "[[7,4,1],[8,5,2],[9,6,3]]"},
			{Input: "matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]", Output: "[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, Expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def rotate(matrix):\n    # Write your solution here (modify in place)\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "90-degree clockwise rotation = transpose + reverse each row. Or reverse columns + transpose."},
			{Level: 2, Type: "algorithm", Content: "Transpose: swap matrix[i][j] with matrix[j][i]. Then reverse each row."},
			{Level: 3, Type: "code", Content: "for i in range(n): for j in range(i+1, n): swap [i][j], [j][i]. Then for each row: reverse."},
		},
		Solution: Solution{
			Code: `def rotate(matrix):
    n = len(matrix)

    # Transpose
    for i in range(n):
        for j in range(i + 1, n):
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

    # Reverse each row
    for row in matrix:
        row.reverse()`,
			Explanation:     "90-degree clockwise = transpose + reverse rows. Both operations in-place.",
			TimeComplexity:  "O(n^2)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Transpose", Explanation: "Swap across diagonal", CodeSnippet: "matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]", LineStart: 7, LineEnd: 7},
				{Title: "Reverse rows", Explanation: "Complete the rotation", CodeSnippet: "row.reverse()", LineStart: 11, LineEnd: 11},
			},
		},
	},
	{
		ID:              "spiral-matrix",
		Number:          93,
		Title:           "Spiral Matrix",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Matrix", "Simulation"},
		RelatedChapters: []int{1},
		Description: `Given an m x n matrix, return all elements of the matrix in spiral order.`,
		Constraints: []string{
			"m == matrix.length",
			"n == matrix[i].length",
			"1 <= m, n <= 10",
			"-100 <= matrix[i][j] <= 100",
		},
		Examples: []Example{
			{Input: "matrix = [[1,2,3],[4,5,6],[7,8,9]]", Output: "[1,2,3,6,9,8,7,4,5]"},
			{Input: "matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]", Output: "[1,2,3,4,8,12,11,10,9,5,6,7]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, Expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def spiralOrder(matrix):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Track four boundaries: top, bottom, left, right. Traverse and shrink boundaries after each direction."},
			{Level: 2, Type: "algorithm", Content: "Go right (top row), down (right col), left (bottom row), up (left col). Update boundaries after each."},
			{Level: 3, Type: "code", Content: "while top <= bottom and left <= right: traverse right, top++; traverse down, right--; check, traverse left, bottom--; check, traverse up, left++."},
		},
		Solution: Solution{
			Code: `def spiralOrder(matrix):
    result = []
    top, bottom = 0, len(matrix) - 1
    left, right = 0, len(matrix[0]) - 1

    while top <= bottom and left <= right:
        # Right
        for c in range(left, right + 1):
            result.append(matrix[top][c])
        top += 1

        # Down
        for r in range(top, bottom + 1):
            result.append(matrix[r][right])
        right -= 1

        if top <= bottom:
            # Left
            for c in range(right, left - 1, -1):
                result.append(matrix[bottom][c])
            bottom -= 1

        if left <= right:
            # Up
            for r in range(bottom, top - 1, -1):
                result.append(matrix[r][left])
            left += 1

    return result`,
			Explanation:     "Four boundaries shrink as we traverse. Go right, down, left, up. Check boundaries before left/up.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(1) extra space",
			Walkthrough: []WalkthroughStep{
				{Title: "Four boundaries", Explanation: "Track traversal limits", CodeSnippet: "top, bottom = 0, len(matrix) - 1\nleft, right = 0, len(matrix[0]) - 1", LineStart: 3, LineEnd: 4},
				{Title: "Traverse and shrink", Explanation: "Go right then increment top", CodeSnippet: "for c in range(left, right + 1):\n    result.append(matrix[top][c])\ntop += 1", LineStart: 8, LineEnd: 10},
				{Title: "Check before left/up", Explanation: "Boundaries may have crossed", CodeSnippet: "if top <= bottom:", LineStart: 17, LineEnd: 17},
			},
		},
	},
	// Bit Manipulation
	{
		ID:              "single-number",
		Number:          94,
		Title:           "Single Number",
		Difficulty:      "Easy",
		Category:        "bit-manipulation",
		Tags:            []string{"Array", "Bit Manipulation"},
		RelatedChapters: []int{1},
		Description: `Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.`,
		Constraints: []string{
			"1 <= nums.length <= 3 * 10^4",
			"-3 * 10^4 <= nums[i] <= 3 * 10^4",
			"Each element in the array appears twice except for one element which appears only once",
		},
		Examples: []Example{
			{Input: "nums = [2,2,1]", Output: "1"},
			{Input: "nums = [4,1,2,1,2]", Output: "4"},
			{Input: "nums = [1]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 2, 1}}, Expected: 1},
			{Input: map[string]interface{}{"nums": []int{4, 1, 2, 1, 2}}, Expected: 4},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def singleNumber(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "XOR has property: a ^ a = 0 and a ^ 0 = a. XOR all numbers - pairs cancel out, leaving single one."},
			{Level: 2, Type: "algorithm", Content: "result = 0. For each num: result ^= num. Return result."},
			{Level: 3, Type: "code", Content: "result = 0. for num in nums: result ^= num. return result."},
		},
		Solution: Solution{
			Code: `def singleNumber(nums):
    result = 0
    for num in nums:
        result ^= num
    return result`,
			Explanation:     "XOR all numbers. Duplicates cancel out (a ^ a = 0), leaving the single number.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "XOR property", Explanation: "a ^ a = 0, a ^ 0 = a", CodeSnippet: "result ^= num", LineStart: 4, LineEnd: 4},
				{Title: "Pairs cancel", Explanation: "Only single number remains", CodeSnippet: "return result", LineStart: 5, LineEnd: 5},
			},
		},
	},
	{
		ID:              "counting-bits",
		Number:          95,
		Title:           "Counting Bits",
		Difficulty:      "Easy",
		Category:        "bit-manipulation",
		Tags:            []string{"Dynamic Programming", "Bit Manipulation"},
		RelatedChapters: []int{1, 9},
		Description: `Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.`,
		Constraints: []string{
			"0 <= n <= 10^5",
		},
		Examples: []Example{
			{Input: "n = 2", Output: "[0,1,1]"},
			{Input: "n = 5", Output: "[0,1,1,2,1,2]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 2}, Expected: []int{0, 1, 1}},
			{Input: map[string]interface{}{"n": 5}, Expected: []int{0, 1, 1, 2, 1, 2}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def countBits(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "DP using relation: i's bit count = (i >> 1)'s count + (i & 1). i >> 1 is already computed."},
			{Level: 2, Type: "algorithm", Content: "dp[i] = dp[i >> 1] + (i & 1). Right shift removes last bit, & 1 checks if last bit is 1."},
			{Level: 3, Type: "code", Content: "dp = [0] * (n + 1). for i in 1..n: dp[i] = dp[i >> 1] + (i & 1). return dp."},
		},
		Solution: Solution{
			Code: `def countBits(n):
    dp = [0] * (n + 1)

    for i in range(1, n + 1):
        dp[i] = dp[i >> 1] + (i & 1)

    return dp`,
			Explanation:     "DP: dp[i] = dp[i//2] + last bit. i >> 1 drops last bit (already computed). i & 1 is the last bit.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "DP relation", Explanation: "Use previously computed value", CodeSnippet: "dp[i] = dp[i >> 1] + (i & 1)", LineStart: 5, LineEnd: 5},
				{Title: "i >> 1", Explanation: "Right shift = divide by 2, drop last bit", CodeSnippet: "dp[i >> 1]", LineStart: 5, LineEnd: 5},
				{Title: "i & 1", Explanation: "Check if last bit is 1", CodeSnippet: "(i & 1)", LineStart: 5, LineEnd: 5},
			},
		},
	},
	{
		ID:              "missing-number",
		Number:          96,
		Title:           "Missing Number",
		Difficulty:      "Easy",
		Category:        "bit-manipulation",
		Tags:            []string{"Array", "Hash Table", "Math", "Binary Search", "Bit Manipulation", "Sorting"},
		RelatedChapters: []int{1},
		Description: `Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.`,
		Constraints: []string{
			"n == nums.length",
			"1 <= n <= 10^4",
			"0 <= nums[i] <= n",
			"All the numbers of nums are unique",
		},
		Examples: []Example{
			{Input: "nums = [3,0,1]", Output: "2"},
			{Input: "nums = [0,1]", Output: "2"},
			{Input: "nums = [9,6,4,2,3,5,7,0,1]", Output: "8"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{3, 0, 1}}, Expected: 2},
			{Input: map[string]interface{}{"nums": []int{0, 1}}, Expected: 2},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def missingNumber(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "XOR: XOR all numbers 0 to n with all array elements. Pairs cancel, leaving missing number. Or use math: sum formula."},
			{Level: 2, Type: "algorithm", Content: "expected_sum = n * (n + 1) / 2. actual_sum = sum(nums). missing = expected - actual."},
			{Level: 3, Type: "code", Content: "n = len(nums). return n * (n + 1) // 2 - sum(nums)."},
		},
		Solution: Solution{
			Code: `def missingNumber(nums):
    n = len(nums)
    expected_sum = n * (n + 1) // 2
    actual_sum = sum(nums)
    return expected_sum - actual_sum`,
			Explanation:     "Sum of 0 to n is n*(n+1)/2. Missing number = expected sum - actual sum.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Expected sum", Explanation: "Sum of 0 to n using formula", CodeSnippet: "expected_sum = n * (n + 1) // 2", LineStart: 3, LineEnd: 3},
				{Title: "Missing = difference", Explanation: "Expected - actual gives missing", CodeSnippet: "return expected_sum - actual_sum", LineStart: 5, LineEnd: 5},
			},
		},
	},
}

// Note: Blind75Categories and Category type are defined in category.go
