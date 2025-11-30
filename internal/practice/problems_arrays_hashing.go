package practice

// ArraysHashingProblems contains all arrays-hashing category problems
var ArraysHashingProblems = []*Problem{
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
		Description: `Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.`,
		Constraints: []string{
			"1 <= strs.length <= 10^4",
			"0 <= strs[i].length <= 100",
			"strs[i] consists of lowercase English letters",
		},
		Examples: []Example{
			{Input: `strs = ["eat","tea","tan","ate","nat","bat"]`, Output: `[["bat"],["nat","tan"],["ate","eat","tea"]]`, Explanation: "The groups are anagrams of each other. Order within groups and order of groups doesn't matter."},
			{Input: `strs = [""]`, Output: `[[""]]`},
			{Input: `strs = ["a"]`, Output: `[["a"]]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"strs": []string{"eat", "tea", "tan", "ate", "nat", "bat"}}, Expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}, OrderIndependent: true},
			{Input: map[string]interface{}{"strs": []string{""}}, Expected: [][]string{{""}}, OrderIndependent: true},
			{Input: map[string]interface{}{"strs": []string{"a"}}, Expected: [][]string{{"a"}}, OrderIndependent: true},
			{Input: map[string]interface{}{"strs": []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}}, Expected: [][]string{{"cab"}, {"tin"}, {"pew"}, {"duh"}, {"may"}, {"ill"}, {"buy"}, {"bar"}, {"max"}, {"doc"}}, OrderIndependent: true},
			{Input: map[string]interface{}{"strs": []string{"listen", "silent", "enlist", "hello", "world"}}, Expected: [][]string{{"listen", "silent", "enlist"}, {"hello"}, {"world"}}, OrderIndependent: true},
		},
		TimeComplexity:  "O(n * k log k)",
		SpaceComplexity: "O(n * k)",
		StarterCode:     "def groupAnagrams(strs):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Anagrams have the same sorted characters."},
			{Level: 2, Type: "algorithm", Content: "Use sorted string as key in a hash map. All words that are anagrams will have the same sorted key."},
			{Level: 3, Type: "code", Content: "Create a defaultdict(list), iterate through strings, use ''.join(sorted(s)) as key, append s to groups[key]."},
		},
		Solution: Solution{
			Code: `from collections import defaultdict

def groupAnagrams(strs):
    groups = defaultdict(list)
    for s in strs:
        key = ''.join(sorted(s))
        groups[key].append(s)
    return list(groups.values())`,
			Explanation: "Sort each string to create a key. Group strings with the same key. Time: O(n * k log k) where n is number of strings and k is max string length.",
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
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 2, 2, 3}, "k": 2}, Expected: []int{1, 2}, OrderIndependent: true},
			{Input: map[string]interface{}{"nums": []int{1}, "k": 1}, Expected: []int{1}, OrderIndependent: true},
			{Input: map[string]interface{}{"nums": []int{4, 1, -1, 2, -1, 2, 3}, "k": 2}, Expected: []int{-1, 2}, OrderIndependent: true},
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
			{Input: map[string]interface{}{"board": [][]string{
				{"5", "3", ".", ".", "7", ".", ".", ".", "."},
				{"6", ".", ".", "1", "9", "5", ".", ".", "."},
				{".", "9", "8", ".", ".", ".", ".", "6", "."},
				{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
				{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", "6", ".", ".", ".", ".", "2", "8", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "5"},
				{".", ".", ".", ".", "8", ".", ".", "7", "9"},
			}}, Expected: true},
			{Input: map[string]interface{}{"board": [][]string{
				{"8", "3", ".", ".", "7", ".", ".", ".", "."},
				{"6", ".", ".", "1", "9", "5", ".", ".", "."},
				{".", "9", "8", ".", ".", ".", ".", "6", "."},
				{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
				{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", "6", ".", ".", ".", ".", "2", "8", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "5"},
				{".", ".", ".", ".", "8", ".", ".", "7", "9"},
			}}, Expected: false},
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
			// Basic case
			{Input: map[string]interface{}{"strs": []string{"hello", "world"}}, Expected: []string{"hello", "world"}},
			// Empty string in list
			{Input: map[string]interface{}{"strs": []string{""}}, Expected: []string{""}},
			// Multiple empty strings
			{Input: map[string]interface{}{"strs": []string{"", "", ""}}, Expected: []string{"", "", ""}},
			// Empty list
			{Input: map[string]interface{}{"strs": []string{}}, Expected: []string{}},
			// Strings containing common delimiters - will fail naive delimiter solutions
			{Input: map[string]interface{}{"strs": []string{"hello~world", "foo~bar"}}, Expected: []string{"hello~world", "foo~bar"}},
			{Input: map[string]interface{}{"strs": []string{"a,b,c", "d,e,f"}}, Expected: []string{"a,b,c", "d,e,f"}},
			{Input: map[string]interface{}{"strs": []string{"split|by|pipe", "more|pipes"}}, Expected: []string{"split|by|pipe", "more|pipes"}},
			// Strings containing the # character (used in length-prefix solution)
			{Input: map[string]interface{}{"strs": []string{"test#with#hash", "another#one"}}, Expected: []string{"test#with#hash", "another#one"}},
			// Strings with numbers (could confuse length-prefix parsing)
			{Input: map[string]interface{}{"strs": []string{"123", "456#789", "10#hello"}}, Expected: []string{"123", "456#789", "10#hello"}},
			// Strings with special characters
			{Input: map[string]interface{}{"strs": []string{"hello\nworld", "tab\there"}}, Expected: []string{"hello\nworld", "tab\there"}},
			{Input: map[string]interface{}{"strs": []string{"space in middle", "  leading", "trailing  "}}, Expected: []string{"space in middle", "  leading", "trailing  "}},
			// Mixed edge cases
			{Input: map[string]interface{}{"strs": []string{"", "nonempty", ""}}, Expected: []string{"", "nonempty", ""}},
			// Single string
			{Input: map[string]interface{}{"strs": []string{"onlyone"}}, Expected: []string{"onlyone"}},
			// Unicode characters
			{Input: map[string]interface{}{"strs": []string{"cafe", "naive", "resume"}}, Expected: []string{"cafe", "naive", "resume"}},
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

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, ArraysHashingProblems...)
}
