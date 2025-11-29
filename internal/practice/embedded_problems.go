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
	// Stack
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
}

// Note: Blind75Categories and Category type are defined in category.go
