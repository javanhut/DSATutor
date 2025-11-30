package practice

// SlidingWindowProblems contains all sliding-window category problems
var SlidingWindowProblems = []*Problem{
	{
		ID:         "longest-substring-without-repeating",
		Number:     15,
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
	{
		ID:         "longest-repeating-character-replacement",
		Number:     16,
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
		Number:     17,
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
		Number:     18,
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
		Number:     19,
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

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, SlidingWindowProblems...)
}
