package practice

import "strings"

// ArraysHashingDrills are the arrays and hash maps foundations: a frequency
// map, the two-pass count-then-find shape (first unique character), prefix
// sums that answer range queries in O(1), prefix sums combined with a map of
// counts (subarray sum equals k), a map of earlier counts that counts pairs
// as it goes, and an in-place array rotation by reversal. They build up to
// Valid Anagram, Top K Frequent Elements, Two Sum, Product of Array Except
// Self and the rest of the ladder without repeating them. Cases with large
// inputs carry a 5 second limit so that rescanning (quadratic) solutions fail.
// They follow docs/pedagogy.md like any problem.
var ArraysHashingDrills = []*Problem{
	{
		ID:              "drill-count-frequencies",
		Number:          0,
		Kind:            KindDrill,
		Order:           1,
		Title:           "Count How Often Each Value Appears",
		Slug:            "drill-count-frequencies",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Hash Table", "Counting", "Foundations"},
		RelatedChapters: []int{2, 1},
		Description:     "Given an array `nums`, count how many times each distinct value appears. Return a list of `[value, count]` pairs, one pair per distinct value, in **any order**.\n\nThe tool for this is a **frequency map**: a dictionary from value to the number of times you have seen it so far. Walk the array once. For each value, look up its current count (0 if it has never been seen) and store that count plus one. A dictionary lookup and store each take O(1) on average, so the whole walk is O(n), where counting each value by rescanning the array would be O(n * distinct).\n\nThis map is the building block under Valid Anagram and Top K Frequent Elements: those problems count first and then compare or rank the counts.",
		Constraints:     []string{"0 <= nums.length <= 100000", "-1000000000 <= nums[i] <= 1000000000", "Return one [value, count] pair per distinct value, in any order", "An empty array has no values, so return []"},
		Examples: []Example{
			{Input: "nums = [1,2,2,3,3,3]", Output: "[[1,1],[2,2],[3,3]]", Explanation: "1 appears once, 2 twice and 3 three times. The pairs may be returned in any order."},
			{Input: "nums = [5,5,5]", Output: "[[5,3]]", Explanation: "One distinct value, seen three times."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 2, 3, 3, 3}}, Expected: [][]int{{1, 1}, {2, 2}, {3, 3}}, Note: "example 1: three values with different counts"},
			{Input: map[string]interface{}{"nums": []int{5, 5, 5}}, Expected: [][]int{{5, 3}}, Note: "example 2: all equal, one pair"},
			{Input: map[string]interface{}{"nums": []int{}}, Expected: []int{}, Hidden: true, Note: "empty input returns an empty list"},
			{Input: map[string]interface{}{"nums": []int{7}}, Expected: [][]int{{7, 1}}, Hidden: true, Note: "a single element"},
			{Input: map[string]interface{}{"nums": []int{0, -1, 0, -1, -1, 0, 4}}, Expected: [][]int{{0, 3}, {-1, 3}, {4, 1}}, Hidden: true, Note: "negatives and zero are ordinary keys; equal counts for different values must stay separate pairs"},
			{Input: map[string]interface{}{"nums": []int{3, 1, 2, 1, 3, 1}}, Expected: [][]int{{3, 2}, {1, 3}, {2, 1}}, Hidden: true, Note: "repeats are not adjacent: a count that only looks at the previous element fails"},
			{Input: map[string]interface{}{"nums": []int{1000000000, -1000000000, 1000000000}}, Expected: [][]int{{1000000000, 2}, {-1000000000, 1}}, Hidden: true, Note: "extreme values from the constraints"},
			{Input: map[string]interface{}{"nums": ahFreqBig()}, Expected: ahFreqExpected(ahFreqBig()), Hidden: true, Timeout: 5000, Note: "100000 values with about 50000 distinct: re-counting each value with a scan is quadratic and runs out of time"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def countFrequencies(nums):\n    # Return [[value, count], ...] with one pair per distinct value, in any order\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You are at some value in the array. What single fact about the earlier values would let you update that value's count without looking back at them?"},
			{Level: 2, Type: "insight", Content: "Keep a dictionary from value to the count so far. Each step reads the old count (0 when the value is new) and stores it plus one; the dict.get(key, 0) call does the read and the default in one go. This frequency map is the core of the hash-map chapter (Chapter 2).", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Start the map: an empty dictionary, counts.\n2. Count: for each value, set counts[value] to its old count (0 if missing) plus 1.\n3. Report: turn each (value, count) item of the map into a [value, count] pair and return the list."},
			{Level: 4, Type: "skeleton", Content: "def countFrequencies(nums):\n    counts = {}\n    for value in nums:\n        counts[value] = counts.get(value, ____) + ____\n    return [[value, count] for value, count in ____]"},
		},
		Solution: Solution{
			Code:            "def countFrequencies(nums):\n    counts = {}\n    for value in nums:\n        counts[value] = counts.get(value, 0) + 1\n    return [[value, count] for value, count in counts.items()]",
			Explanation:     "This is the plainest use of a hash map: the map holds, for every value passed so far, how many times it has appeared. One pass builds it and the items are the answer. Valid Anagram and Top K Frequent Elements in the ladder both start with exactly this map and then compare two of them or sort by count. The alternatives show the two standard library shortcuts.",
			BruteForce:      "For each distinct value, scan the whole array and count its matches with nums.count(value). That is O(n) per distinct value, O(n * distinct) in all, about 5 billion steps for the big case. The waste is that the array is re-read from the start for every value, although one pass could update all the counts together.",
			KeyInsight:      "A dictionary lookup costs O(1), so keep a running count per value and update it as you walk, instead of searching for the value's matches later.",
			WhyItWorks:      "Invariant: after the loop has processed nums[0..i], counts[v] equals the number of times v occurs in nums[0..i] for every v, and v is a key exactly when it occurred. It holds before the loop (empty map, empty prefix). Processing one more element x changes only the count of x, by exactly one, which is what the update does, so the invariant carries on. At the end the prefix is the whole array.",
			CommonMistakes:  []string{"Writing `counts[value] = 1` inside the loop. Every value then ends with count 1, because the store overwrites the earlier count; read the old count first and add one.", "Writing `counts[value] += 1` on a value that is not yet a key. The first time a value appears it raises KeyError; use counts.get(value, 0) + 1, a defaultdict(int) or a Counter.", "Returning one pair per element (`[[v, nums.count(v)] for v in nums]`) instead of one per distinct value: repeated values appear several times, and it is O(n^2) as well.", "Swapping the pair to [count, value]: the statement wants the value first."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Start the map", Explanation: "counts maps each value seen so far to how many times it has appeared. Before any element is read it is empty, which is correct for an empty prefix.", CodeSnippet: "counts = {}", LineStart: 2, LineEnd: 2},
				{Title: "Read the old count", Explanation: "counts.get(value, 0) gives the number of earlier copies of this value, and 0 when it is new. This avoids the KeyError that counts[value] would raise.", CodeSnippet: "counts.get(value, 0) + 1", LineStart: 4, LineEnd: 4},
				{Title: "Store the new count", Explanation: "Write the count back with this copy included. For [1,2,2], the map goes {1:1}, then {1:1, 2:1}, then {1:1, 2:2}.", CodeSnippet: "counts[value] = counts.get(value, 0) + 1", LineStart: 4, LineEnd: 4},
				{Title: "Report the pairs", Explanation: "Each item of the map is one distinct value with its final count, which is exactly one answer pair.", CodeSnippet: "return [[value, count] for value, count in counts.items()]", LineStart: 5, LineEnd: 5},
			},
			Alternatives: []AltSolution{
				{Name: "Counter", Code: "from collections import Counter\n\n\ndef countFrequencies(nums):\n    return [[value, count] for value, count in Counter(nums).items()]", TimeComplexity: "O(n)", SpaceComplexity: "O(n)", Explanation: "Counter builds the same frequency map in one call. Know the loop above first: it is what Counter does, and interviews expect you to be able to write it."},
				{Name: "defaultdict", Code: "from collections import defaultdict\n\n\ndef countFrequencies(nums):\n    counts = defaultdict(int)\n    for value in nums:\n        counts[value] += 1\n    return [[value, count] for value, count in counts.items()]", TimeComplexity: "O(n)", SpaceComplexity: "O(n)", Explanation: "A defaultdict(int) supplies 0 for a missing key, so counts[value] += 1 is safe on the first sight of a value."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is a dictionary the right choice here and not a list indexed by value?", Answer: "A list indexed by value needs a slot for every possible value; here values reach plus or minus a billion (and could be negative), so that list would be enormous and mostly empty. A dictionary stores only the values that actually appear and still finds each in O(1) on average."},
			{Prompt: "What changes if you must return the pairs sorted by value?", Answer: "Counting stays O(n); then sort the distinct pairs, which costs O(d log d) for d distinct values. The counting pass does not need the order, which is why the problem allows any order."},
		},
	},
	{
		ID:              "drill-first-unique-character",
		Number:          387,
		Kind:            KindDrill,
		Order:           2,
		Title:           "First Unique Character",
		Slug:            "drill-first-unique-character",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Hash Table", "String", "Counting", "Foundations"},
		RelatedChapters: []int{2, 1},
		Description:     "Given a string `s` of lowercase letters, return the **index** of the first character that appears exactly once in the whole string. If every character repeats, return `-1`.\n\nA character is unique if it occurs once in the entire string, not just once so far. That is why one pass is not enough: when you reach the first `l` of `leetcode` you cannot yet know whether another `l` comes later. Use two passes that share a map:\n\n1. **Count**: build a frequency map of every character (the previous drill).\n2. **Find**: walk the string again from the left, and return the index of the first character whose count is 1.\n\nThe second pass must walk the **string**, not the map's keys, because the answer is the earliest position, and that is what the string's order gives you.",
		Constraints:     []string{"1 <= s.length <= 100000", "s consists of lowercase English letters", "Return an index, not the character", "Return -1 when no character is unique"},
		Examples: []Example{
			{Input: "s = \"leetcode\"", Output: "0", Explanation: "l appears once, and it is the first character."},
			{Input: "s = \"loveleetcode\"", Output: "2", Explanation: "l, o and e all repeat; v at index 2 is the first letter that appears once."},
			{Input: "s = \"aabb\"", Output: "-1", Explanation: "Every letter repeats, so there is no answer."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "leetcode"}, Expected: 0, Note: "example 1: the first character is already unique"},
			{Input: map[string]interface{}{"s": "loveleetcode"}, Expected: 2, Note: "example 2: the answer is after some repeated characters"},
			{Input: map[string]interface{}{"s": "aabb"}, Expected: -1, Hidden: true, Note: "no unique character returns -1"},
			{Input: map[string]interface{}{"s": "z"}, Expected: 0, Hidden: true, Note: "a single character is unique"},
			{Input: map[string]interface{}{"s": "aabbc"}, Expected: 4, Hidden: true, Note: "the only unique character is the last one"},
			{Input: map[string]interface{}{"s": "abcabcd"}, Expected: 6, Hidden: true, Note: "early characters repeat later in the string, so a character is not unique just because it is new when first seen"},
			{Input: map[string]interface{}{"s": "dddccdbba"}, Expected: 8, Hidden: true, Note: "the unique character is last and the repeated letters are not adjacent: any early return before counting finishes is wrong"},
			{Input: map[string]interface{}{"s": "cbbcdae"}, Expected: 4, Hidden: true, Note: "two unique characters (d, a, e): the answer is the first of them, not the last"},
			{Input: map[string]interface{}{"s": ahUniqBig()}, Expected: ahFirstUniq(ahUniqBig()), Hidden: true, Timeout: 5000, Note: "99001 characters in three long runs of a, b and c followed by one d: comparing each character against every other one in a pure Python double loop is billions of steps and runs out of time"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def firstUniqChar(s):\n    # Return the index of the first character that appears exactly once, or -1\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When you stand on the first character, what do you need to know about the whole rest of the string before you can say it is unique? What if you had gathered that before choosing?"},
			{Level: 2, Type: "insight", Content: "Do it in two passes. Pass one counts every character in a frequency map (Chapter 2). Pass two walks the string in order and stops at the first character whose count is 1. The map answers 'how many times does this character appear anywhere?' in O(1).", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Count: one pass over s, building counts[char].\n2. Find: a second pass over s with the index; return the index of the first character with counts[char] == 1.\n3. Fallback: if the second pass ends without a hit, return -1."},
			{Level: 4, Type: "skeleton", Content: "def firstUniqChar(s):\n    counts = {}\n    for ch in s:\n        counts[ch] = counts.get(ch, 0) + 1\n    for i, ch in enumerate(s):\n        if counts[ch] == ____:\n            return ____\n    return ____"},
		},
		Solution: Solution{
			Code:            "def firstUniqChar(s):\n    counts = {}\n    for ch in s:\n        counts[ch] = counts.get(ch, 0) + 1\n    for i, ch in enumerate(s):\n        if counts[ch] == 1:\n            return i\n    return -1",
			Explanation:     "This builds on the frequency map of the previous drill and adds the idea that a map can answer a question about the whole input after one full pass: the counts are final before the second pass reads them. Walking the string, and not the map, in the second pass is what makes the answer the earliest index. The same count-then-decide shape appears in Valid Anagram and Top K Frequent Elements.",
			BruteForce:      "For each index i, compare s[i] with every other character (or call s.count(s[i])) and stop at the first one with no match. That is O(n^2) in all: the string is rescanned for every position, and a double loop in pure Python runs out of time on a long string. The waste is that the same letter is recounted at each of its positions.",
			KeyInsight:      "Uniqueness is a fact about the whole string, so finish counting first; then a second walk in string order finds the earliest character whose final count is 1.",
			WhyItWorks:      "After pass one, counts[c] is the number of occurrences of c in all of s (the invariant of the frequency map at the end of the prefix = whole string). A character is unique exactly when its count is 1. Pass two visits the indices in increasing order, so the first index it accepts is the smallest index of a unique character. If it accepts none, no character has count 1, so -1 is right. With only 26 letters the map has at most 26 keys, so the extra space is O(1).",
			CommonMistakes:  []string{"Returning the last unique character instead of the first, for example by walking the string from the right or by overwriting a result variable; stop at the first hit when walking left to right.", "Returning the character (`return ch`) instead of its index.", "Checking uniqueness while counting, for example returning at the first character not seen before. The first character is new when you reach it, but may repeat later; finish counting first.", "Returning 0 or None instead of -1 when no character is unique.", "Comparing every character with every other one in a double loop. It is correct, but on a 99001-character string it runs out of time."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Count every character", Explanation: "The first pass makes counts final: counts[c] is the number of times c appears in the whole string.", CodeSnippet: "for ch in s:\n    counts[ch] = counts.get(ch, 0) + 1", LineStart: 3, LineEnd: 4},
				{Title: "Walk in string order", Explanation: "The second pass reads the string, not the map, so the first hit is the earliest index. enumerate supplies the index.", CodeSnippet: "for i, ch in enumerate(s):", LineStart: 5, LineEnd: 5},
				{Title: "Test the count", Explanation: "A final count of 1 means the character appears nowhere else. Return its index right away.", CodeSnippet: "if counts[ch] == 1:\n    return i", LineStart: 6, LineEnd: 7},
				{Title: "Handle no answer", Explanation: "Reaching the end means every character repeats.", CodeSnippet: "return -1", LineStart: 8, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{Name: "Index map", Code: "def firstUniqChar(s):\n    first_index = {}\n    for i, ch in enumerate(s):\n        if ch in first_index:\n            first_index[ch] = -1\n        else:\n            first_index[ch] = i\n    unique = [i for i in first_index.values() if i != -1]\n    return min(unique) if unique else -1", TimeComplexity: "O(n)", SpaceComplexity: "O(1)", Explanation: "One pass over s, with a different map: it stores each character's first index and replaces it with -1 once the character repeats. The surviving indices belong to unique characters; the smallest is the answer."},
				{Name: "Brute force", Code: "def firstUniqChar(s):\n    for i in range(len(s)):\n        if s.count(s[i]) == 1:\n            return i\n    return -1", TimeComplexity: "O(n^2)", SpaceComplexity: "O(1)", Explanation: "Ask the string how often s[i] occurs, for each i. The count runs in C so it passes here, but it still rescans the string for each position and grows quadratically."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can't the first pass also decide the answer?", Answer: "A character's count is only final after the whole string is read: the first l of leetcode looks new, but a later l makes it repeat. So the decision has to wait for a second pass, or for a map that can undo a guess (the index-map alternative)."},
			{Prompt: "Why is the space O(1) when we use a dictionary?", Answer: "The key set is the lowercase alphabet, at most 26 letters, however long the string. A map whose size is bounded by a constant does not grow with n. If the characters were arbitrary Unicode, it would be O(distinct characters) instead."},
		},
	},
	{
		ID:              "drill-range-sum-queries",
		Number:          0,
		Kind:            KindDrill,
		Order:           3,
		Title:           "Many Range Sum Queries",
		Slug:            "drill-range-sum-queries",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Prefix Sum", "Foundations"},
		RelatedChapters: []int{2, 1},
		Description:     "Given an integer array `nums` and a list of `queries`, where each query is `[left, right]` (both **inclusive**, 0-indexed), return a list with the sum of `nums[left..right]` for every query, in the same order.\n\nRe-adding the range for each query costs O(length of range), up to O(n) per query. With 100,000 queries on a 100,000-element array that is billions of additions. Instead, do the adding once, up front, with a **prefix-sum array**: `prefix[i]` is the sum of the first `i` elements, so `prefix[0] = 0` and `prefix[i + 1] = prefix[i] + nums[i]`. Then a query is a subtraction: `sum(nums[left..right]) = prefix[right + 1] - prefix[left]`.\n\nThe extra leading 0 is the point: it lets a range starting at index 0 use the same formula with no special case. This is the idea inside the Range Sum Query problem, and the next drill puts a hash map on top of it.",
		Constraints:     []string{"1 <= nums.length <= 100000", "-1000 <= nums[i] <= 1000", "0 <= queries.length <= 100000", "0 <= left <= right < nums.length, both ends inclusive", "Return the answers in the order of the queries"},
		Examples: []Example{
			{Input: "nums = [-2,0,3,-5,2,-1], queries = [[0,2],[2,5],[0,5]]", Output: "[1,-1,-3]", Explanation: "-2+0+3 = 1; 3-5+2-1 = -1; the whole array sums to -3."},
			{Input: "nums = [4], queries = [[0,0]]", Output: "[4]", Explanation: "A one-element range is that element."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{-2, 0, 3, -5, 2, -1}, "queries": [][]int{{0, 2}, {2, 5}, {0, 5}}}, Expected: []int{1, -1, -3}, Note: "example 1: ranges that start at index 0, in the middle and cover everything"},
			{Input: map[string]interface{}{"nums": []int{4}, "queries": [][]int{{0, 0}}}, Expected: []int{4}, Note: "example 2: a single element"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5}, "queries": []int{}}, Expected: []int{}, Hidden: true, Note: "no queries returns an empty list"},
			{Input: map[string]interface{}{"nums": []int{5, -3, 8, 1}, "queries": [][]int{{0, 0}, {1, 1}, {3, 3}}}, Expected: []int{5, -3, 1}, Hidden: true, Note: "single-element ranges, including at index 0 and at the last index"},
			{Input: map[string]interface{}{"nums": []int{3, 1, 4, 1, 5, 9, 2, 6}, "queries": [][]int{{2, 4}, {0, 7}, {5, 6}, {1, 1}, {0, 3}}}, Expected: []int{10, 31, 11, 1, 9}, Hidden: true, Note: "several queries on one array, answered in input order"},
			{Input: map[string]interface{}{"nums": []int{-1, -1, -1, -1}, "queries": [][]int{{1, 3}, {0, 1}}}, Expected: []int{-3, -2}, Hidden: true, Note: "all negative values"},
			{Input: map[string]interface{}{"nums": []int{1000, 1000, 1000, 1000, 1000}, "queries": [][]int{{0, 4}}}, Expected: []int{5000}, Hidden: true, Note: "large values add up past the element size"},
			{Input: map[string]interface{}{"nums": ahRangeNums(), "queries": ahRangeQueries()}, Expected: ahRangeSums(ahRangeNums(), ahRangeQueries()), Hidden: true, Timeout: 5000, Note: "100000 elements and 100000 wide queries: adding every range from scratch is about 5 * 10^9 additions and runs out of time"},
		},
		TimeComplexity:  "O(n + q)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def rangeSums(nums, queries):\n    # Return the sum of nums[left..right] (inclusive) for each [left, right] in queries\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you already knew the sum of the first right+1 elements and the sum of the first left elements, how would you get the sum between them with one operation?"},
			{Level: 2, Type: "insight", Content: "Precompute running totals once: prefix[i] is the sum of the first i elements. The range sum is then the difference of two totals, prefix[right + 1] - prefix[left]. Starting the array with prefix[0] = 0 means left = 0 needs no special case. This is the prefix-sum idea from Chapter 2.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Prepare: build prefix with prefix[0] = 0 and prefix[i + 1] = prefix[i] + nums[i].\n2. Answer: for each [left, right], compute prefix[right + 1] - prefix[left].\n3. Collect: return the answers in query order."},
			{Level: 4, Type: "skeleton", Content: "def rangeSums(nums, queries):\n    prefix = [0]\n    for value in nums:\n        prefix.append(prefix[-1] + ____)\n    return [prefix[____] - prefix[____] for left, right in queries]"},
		},
		Solution: Solution{
			Code:            "def rangeSums(nums, queries):\n    prefix = [0]\n    for value in nums:\n        prefix.append(prefix[-1] + value)\n    return [prefix[right + 1] - prefix[left] for left, right in queries]",
			Explanation:     "This is the first drill where precomputation pays for repeated queries: one pass builds an array of running totals, and then every range sum is two lookups and a subtraction. It builds on the walk-once habit of the counting drills, and the next drill (Subarray Sum Equals K) combines these prefix totals with a hash map, which Range Sum Query-style problems and the ladder's Product of Array Except Self (prefix and suffix) also lean on.",
			BruteForce:      "For each query, add up nums[left..right] with a loop or sum(nums[left:right + 1]). That is O(n) per query and O(n * q) in all, around 5 * 10^9 additions for the big case. The waste is that overlapping ranges re-add the same elements.",
			KeyInsight:      "Any range sum is the difference of two running totals, so compute the running totals once and every query becomes O(1).",
			WhyItWorks:      "Invariant: prefix[i] equals nums[0] + ... + nums[i - 1], with prefix[0] = 0 (the empty sum). The sum of nums[left..right] is (nums[0] + ... + nums[right]) minus (nums[0] + ... + nums[left - 1]), that is prefix[right + 1] - prefix[left]. The elements before left cancel, leaving exactly the elements from left to right. Building prefix is n additions and each query is one subtraction.",
			CommonMistakes:  []string{"Leaving out the leading 0, so prefix[i] means the sum up to and including i. Then the formula needs prefix[left - 1], which for left = 0 reads prefix[-1], the last element in Python, and silently gives a wrong answer.", "Using prefix[right] - prefix[left] with the leading 0: that excludes nums[right], which is off by one. The range is inclusive, so use prefix[right + 1].", "Re-summing the slice for every query. It is right but quadratic and runs out of time on the large case.", "Returning the answers out of order or as a single total; return one sum per query, in query order."},
			TimeComplexity:  "O(n + q)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Seed the prefix array", Explanation: "prefix[0] = 0 is the sum of no elements. It is what lets a range that starts at index 0 subtract nothing.", CodeSnippet: "prefix = [0]", LineStart: 2, LineEnd: 2},
				{Title: "Build the running totals", Explanation: "Each new total adds one more element to the previous one, so prefix[i] is the sum of the first i elements. For [-2, 0, 3] the array is [0, -2, -2, 1].", CodeSnippet: "for value in nums:\n    prefix.append(prefix[-1] + value)", LineStart: 3, LineEnd: 4},
				{Title: "Answer a query by subtraction", Explanation: "The total up to right, minus the total before left, leaves the elements from left to right. The +1 is because prefix[right + 1] includes nums[right].", CodeSnippet: "return [prefix[right + 1] - prefix[left] for left, right in queries]", LineStart: 5, LineEnd: 5},
			},
			Alternatives: []AltSolution{
				{Name: "itertools.accumulate", Code: "from itertools import accumulate\n\n\ndef rangeSums(nums, queries):\n    prefix = [0] + list(accumulate(nums))\n    return [prefix[right + 1] - prefix[left] for left, right in queries]", TimeComplexity: "O(n + q)", SpaceComplexity: "O(n)", Explanation: "accumulate yields the running totals, so putting a 0 in front gives the same prefix array."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the prefix array have n + 1 entries and not n?", Answer: "The extra first entry, prefix[0] = 0, stands for the empty prefix. Without it a range that starts at index 0 would need a special case or would read prefix[-1]. With it, every range uses the same formula."},
			{Prompt: "What changes if the array can be modified between queries?", Answer: "Prefix sums become stale: changing nums[i] changes every prefix after i, which is O(n) to repair. For many updates and queries together you need a structure such as a Fenwick tree or segment tree, which gives O(log n) for both."},
		},
	},
	{
		ID:              "drill-subarray-sum-equals-k",
		Number:          560,
		Kind:            KindDrill,
		Order:           4,
		Title:           "Subarray Sum Equals K",
		Slug:            "drill-subarray-sum-equals-k",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Hash Table", "Prefix Sum", "Array", "Foundations"},
		RelatedChapters: []int{2, 1},
		Description:     "Given an integer array `nums` and an integer `k`, return the number of **contiguous, non-empty subarrays** whose elements sum to `k`. The array can contain negative numbers and zeros, so different subarrays can have the same sum, and a window that only grows or shrinks cannot track the sum.\n\nCombine the last two drills. Let `prefix` be the running sum of the elements read so far. The subarray ending at the current element and starting just after an earlier point sums to `prefix - earlier_prefix`, so it equals `k` exactly when `earlier_prefix == prefix - k`. Keep a map from each earlier prefix value to **how many times it has occurred**, and add that count to the answer at every step.\n\nSeed the map with `{0: 1}`: the empty prefix before the first element has sum 0, and without it a subarray that starts at index 0 is never counted.",
		Constraints:     []string{"1 <= nums.length <= 50000", "-1000 <= nums[i] <= 1000", "-10000000 <= k <= 10000000", "The subarray must be contiguous and contain at least one element", "Count subarrays by position: equal subarrays at different positions count separately"},
		Examples: []Example{
			{Input: "nums = [1,1,1], k = 2", Output: "2", Explanation: "The subarrays [1,1] at indexes 0-1 and at indexes 1-2."},
			{Input: "nums = [1,2,3], k = 3", Output: "2", Explanation: "[1,2] and [3]."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 1, 1}, "k": 2}, Expected: 2, Note: "example 1: overlapping subarrays with equal values count twice"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}, "k": 3}, Expected: 2, Note: "example 2: one subarray of two elements and one of a single element"},
			{Input: map[string]interface{}{"nums": []int{5}, "k": 5}, Expected: 1, Hidden: true, Note: "a single element equal to k: counted only if the map is seeded with {0: 1}"},
			{Input: map[string]interface{}{"nums": []int{1}, "k": 0}, Expected: 0, Hidden: true, Note: "a single element and k = 0: nothing sums to 0, and the empty subarray must not count"},
			{Input: map[string]interface{}{"nums": []int{1, -1, 0}, "k": 0}, Expected: 3, Hidden: true, Note: "negatives and zero: [1,-1], [0] and [1,-1,0] all sum to 0"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0}, "k": 0}, Expected: 6, Hidden: true, Note: "k = 0 with all zeros: every non-empty subarray counts (1 + 2 + 3), so the map must hold counts, not just yes or no"},
			{Input: map[string]interface{}{"nums": []int{3, 4, 7, 2, -3, 1, 4, 2}, "k": 7}, Expected: 4, Hidden: true, Note: "a mixed array: [3,4], [7], [7,2,-3,1] and [1,4,2]; a window that shrinks when the sum is too big misses the one with a negative"},
			{Input: map[string]interface{}{"nums": []int{-1, -1, 1}, "k": 0}, Expected: 1, Hidden: true, Note: "negatives: only [-1,1] works, so a sliding window that assumes positives fails"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 1, 2, 1}, "k": 3}, Expected: 4, Hidden: true, Note: "the same sum appears from several starting points"},
			{Input: map[string]interface{}{"nums": ahZeros(50000), "k": 0}, Expected: 1250025000, Hidden: true, Timeout: 5000, Note: "50000 zeros with k = 0: every one of the 1250025000 non-empty subarrays counts, and checking all the subarrays one by one runs out of time"},
			{Input: map[string]interface{}{"nums": ahSubBig(), "k": 7}, Expected: ahSubCount(ahSubBig(), 7), Hidden: true, Timeout: 5000, Note: "50000 random small values in -5..5: checking each start and end pair is over 10^9 steps"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def subarraySum(nums, k):\n    # Return the number of contiguous non-empty subarrays whose sum is k\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "At the current element, the subarray you want ends here. What would you need to know about the running total at the point where it starts?"},
			{Level: 2, Type: "insight", Content: "With running total prefix, a subarray that ends here and starts after an earlier prefix value e has sum prefix - e, so it is k when e = prefix - k. Keep a map from each earlier prefix value to how many times it has occurred and add that count. Seed the map with {0: 1} for the empty prefix. This is the prefix-sum plus hash-map idea from Chapter 2.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Seed: seen = {0: 1}, prefix = 0, count = 0.\n2. Look up first: add prefix += value, then count += seen.get(prefix - k, 0).\n3. Record after: add this prefix to the map, seen[prefix] = seen.get(prefix, 0) + 1.\n4. Return count."},
			{Level: 4, Type: "skeleton", Content: "def subarraySum(nums, k):\n    seen = {0: ____}\n    prefix = 0\n    count = 0\n    for value in nums:\n        prefix += value\n        count += seen.get(____, 0)\n        seen[prefix] = seen.get(prefix, 0) + 1\n    return count"},
		},
		Solution: Solution{
			Code:            "def subarraySum(nums, k):\n    seen = {0: 1}\n    prefix = 0\n    count = 0\n    for value in nums:\n        prefix += value\n        count += seen.get(prefix - k, 0)\n        seen[prefix] = seen.get(prefix, 0) + 1\n    return count",
			Explanation:     "This is Many Range Sum Queries run in reverse: instead of asking for the sum between two chosen points, ask how many earlier points make the sum equal to k, and let a frequency map (the first drill) answer in O(1). Counts, not yes or no, are stored because several earlier points can share a prefix value, and each of them is a different subarray. Three details carry the whole problem: the {0: 1} seed, look up before you record, and keep counts.",
			BruteForce:      "Try every start and every end and keep a running sum, or sum each pair afresh. With a running sum that is O(n^2) pairs, over 10^9 for the big cases. The waste is that for a fixed end, the starts that work are exactly those with a particular earlier prefix value, which a map can count in O(1) instead of testing every start.",
			KeyInsight:      "A subarray sums to k exactly when two prefix totals differ by k, so while walking, count how many earlier prefix totals equal prefix - k.",
			WhyItWorks:      "Invariant: before processing element i, seen[p] is the number of indices j in 0..i (counting the empty prefix j = 0) whose prefix sum is p. The subarrays ending at element i are those starting after one of these points j, with sum prefix_i - prefix_j. It equals k exactly when prefix_j = prefix_i - k, and there are seen[prefix_i - k] such j, so that is how many subarrays end at i. Looking up before recording means the current prefix is not counted against itself (that would be the empty subarray, which only matters when k = 0). Summing over i counts every subarray once, by its end.",
			CommonMistakes:  []string{"Starting with an empty map instead of {0: 1}. A subarray that begins at index 0 uses the empty prefix, which has sum 0; without the seed it is never counted, so [5] with k = 5 gives 0.", "Storing only whether a prefix was seen (a set, or a map to True) and adding 1. A prefix value can occur many times (zeros, or sums that return to the same value), and each occurrence is a different start; store a count and add the count.", "Recording the current prefix before looking up prefix - k. When k = 0 the lookup then finds the current prefix itself and counts an empty subarray at every element.", "Using a sliding window that shrinks when the sum exceeds k. That relies on all values being positive; with negatives a longer window can bring the sum back to k.", "Checking every start and end pair. It is correct but takes over 10^9 steps on the large cases and runs out of time."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Seed the empty prefix", Explanation: "The sum before the first element is 0, and it has happened once. This is what lets a subarray that starts at index 0 be counted.", CodeSnippet: "seen = {0: 1}", LineStart: 2, LineEnd: 2},
				{Title: "Extend the running total", Explanation: "prefix is the sum of nums[0..i]. It is the right-hand end of every subarray this step can complete.", CodeSnippet: "prefix += value", LineStart: 6, LineEnd: 6},
				{Title: "Count the subarrays ending here", Explanation: "Each earlier prefix equal to prefix - k marks a start whose subarray sums to k. The map gives how many in one lookup. For [1,1,1] with k = 2 the lookup at the third element finds prefix 1 once.", CodeSnippet: "count += seen.get(prefix - k, 0)", LineStart: 7, LineEnd: 7},
				{Title: "Record this prefix", Explanation: "Only now add the current prefix to the map, so later elements can use it as a start. Add one to its count.", CodeSnippet: "seen[prefix] = seen.get(prefix, 0) + 1", LineStart: 8, LineEnd: 8},
			},
			Alternatives: []AltSolution{},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why do we store counts and not just which prefix sums have appeared?", Answer: "Different earlier points can have the same prefix sum (for example, around zeros, or when a stretch sums to 0). Each of them is a separate start and gives a separate subarray, so the answer has to add the number of them, not 1."},
			{Prompt: "Why does the lookup go before the update when k = 0?", Answer: "With k = 0 we look up prefix - 0 = prefix. If the current prefix were already in the map, the lookup would find it and count a subarray from this position to itself with no elements. Looking up first sees only earlier prefixes."},
		},
	},
	{
		ID:              "drill-count-pairs-with-difference",
		Number:          2006,
		Kind:            KindDrill,
		Order:           5,
		Title:           "Count Pairs with Absolute Difference K",
		Slug:            "drill-count-pairs-with-difference",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Hash Table", "Counting", "Array", "Foundations"},
		RelatedChapters: []int{2, 1},
		Description:     "Given an integer array `nums` and a positive integer `k`, return the number of pairs of positions `(i, j)` with `i < j` such that `|nums[i] - nums[j]| == k`. Pairs are counted by **position**, so if a value appears twice, each copy forms its own pair.\n\nCounting pairs is where a hash map has to hold **counts**, and where you must not count a pair twice. Walk the array left to right and keep a map from each value to how many times it has appeared **before** the current position. For the current value `x`, a partner that is earlier and differs by `k` is either `x - k` or `x + k`; add the map's count for each. Then record `x`. Every pair `(i, j)` with `i < j` is counted once, when `j` is reached.\n\nThis is Two Sum's seen-so-far map with two lookups per step and with counts instead of indexes.",
		Constraints:     []string{"1 <= nums.length <= 100000", "-1000000000 <= nums[i] <= 1000000000", "1 <= k <= 2000000000", "Pairs are by position (i < j); equal values at different positions are different pairs"},
		Examples: []Example{
			{Input: "nums = [1,2,2,1], k = 1", Output: "4", Explanation: "The pairs of positions (0,1), (0,2), (1,3), (2,3): each pairs a 1 with a 2."},
			{Input: "nums = [1,3], k = 3", Output: "0", Explanation: "1 and 3 differ by 2, not 3."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 2, 1}, "k": 1}, Expected: 4, Note: "example 1: duplicates, each copy forms its own pair"},
			{Input: map[string]interface{}{"nums": []int{1, 3}, "k": 3}, Expected: 0, Note: "example 2: no pair has the difference"},
			{Input: map[string]interface{}{"nums": []int{7}, "k": 1}, Expected: 0, Hidden: true, Note: "a single element has no pair"},
			{Input: map[string]interface{}{"nums": []int{3, 2, 1, 5, 4}, "k": 2}, Expected: 3, Hidden: true, Note: "pairs (3,1), (3,5) and (2,4): some partners are smaller and some larger than the value being read"},
			{Input: map[string]interface{}{"nums": []int{5, 3}, "k": 2}, Expected: 1, Hidden: true, Note: "the partner is larger than the value read second: a lookup of only x - k misses it"},
			{Input: map[string]interface{}{"nums": []int{3, 5}, "k": 2}, Expected: 1, Hidden: true, Note: "the partner is smaller than the value read second: a lookup of only x + k misses it"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 2}, "k": 1}, Expected: 0, Hidden: true, Note: "equal values never differ by k when k >= 1"},
			{Input: map[string]interface{}{"nums": []int{4, 4, 6, 6, 4}, "k": 2}, Expected: 6, Hidden: true, Note: "repeats on both sides: a map of yes/no instead of counts gives 4, and counting each pair from both ends gives 12"},
			{Input: map[string]interface{}{"nums": []int{-5, 0, 5, -10, 10}, "k": 5}, Expected: 4, Hidden: true, Note: "negative values and a pair across zero"},
			{Input: map[string]interface{}{"nums": []int{-1000000000, 1000000000}, "k": 2000000000}, Expected: 1, Hidden: true, Note: "the largest difference the constraints allow"},
			{Input: map[string]interface{}{"nums": ahDiffBig(), "k": 3}, Expected: ahDiffCount(ahDiffBig(), 3), Hidden: true, Timeout: 5000, Note: "100000 values in 0..999: checking every pair is 5 * 10^9 steps and runs out of time"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def countKDifference(nums, k):\n    # Return the number of pairs (i < j) with |nums[i] - nums[j]| == k\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You are reading nums[j] and want its partners among the earlier numbers. Which two values could a partner have, and what would you need stored to count them instantly?"},
			{Level: 2, Type: "insight", Content: "A partner differs from x by k, so it is x - k or x + k (these are different values because k is at least 1). Keep a map from each earlier value to how many times it appeared, add the counts for both partners, then record x. Reading left to right with 'look up, then record' counts each pair once. This is the seen-so-far map of Chapter 2.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Map: seen = {} from value to its count so far.\n2. Look up: for each x, add seen.get(x - k, 0) + seen.get(x + k, 0) to the total.\n3. Record: increase seen[x] by one after the lookup.\n4. Return the total."},
			{Level: 4, Type: "skeleton", Content: "def countKDifference(nums, k):\n    seen = {}\n    pairs = 0\n    for value in nums:\n        pairs += seen.get(value - k, 0) + seen.get(____, 0)\n        seen[value] = seen.get(value, 0) + ____\n    return pairs"},
		},
		Solution: Solution{
			Code:            "def countKDifference(nums, k):\n    seen = {}\n    pairs = 0\n    for value in nums:\n        pairs += seen.get(value - k, 0) + seen.get(value + k, 0)\n        seen[value] = seen.get(value, 0) + 1\n    return pairs",
			Explanation:     "This is Two Sum's idea (Chapter 2: remember what you have passed and look up the partner) with the two changes that make counting problems tricky: the map stores counts, because duplicates form separate pairs, and there are two possible partner values. Reading left to right and looking up before recording counts every pair (i, j) exactly once, at j, which avoids dividing a double count by two.",
			BruteForce:      "Test every pair (i, j) with i < j: n(n - 1) / 2 differences, about 5 * 10^9 for the big case, which runs out of time. The waste is that for a fixed value only two other values can pair with it, but the double loop looks at everything.",
			KeyInsight:      "For any value x, its partners can only be x - k and x + k, so a map of counts answers 'how many earlier partners?' in two lookups.",
			WhyItWorks:      "Invariant: before processing position j, seen[v] is the number of positions i < j with nums[i] = v. A pair (i, j) with i < j satisfies |nums[i] - nums[j]| = k exactly when nums[i] is nums[j] - k or nums[j] + k; those are different values because k >= 1, so no earlier position is counted twice. The number of such i is the sum of the two counts. Summing over j counts each pair once, at its larger position.",
			CommonMistakes:  []string{"Looking up only `value - k`. That finds partners that are smaller than the current value, so a pair like [5, 3] with k = 2 (5 first, 3 second) is missed.", "Storing only whether a value was seen (a set) and adding 1. A value that appeared three times forms three pairs with each partner; store and add counts.", "Building the full count map first and then adding both lookups for every element. Each pair is then counted from both of its ends and the total is doubled; look up while reading left to right, or divide by two.", "Checking every pair with a double loop. It is correct but runs out of time at 100000 elements."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Keep earlier counts", Explanation: "seen maps each value to how many times it has appeared at an earlier position. It starts empty.", CodeSnippet: "seen = {}", LineStart: 2, LineEnd: 2},
				{Title: "Count partners below and above", Explanation: "The earlier partners of value are the earlier copies of value - k and of value + k. For [1,2,2,1] with k = 1 and value = 1 at the end, seen has 2 twice, and value - k = 0 is absent.", CodeSnippet: "pairs += seen.get(value - k, 0) + seen.get(value + k, 0)", LineStart: 5, LineEnd: 5},
				{Title: "Record the value after the lookups", Explanation: "Adding it afterwards means a value is never paired with itself, and later elements can find it.", CodeSnippet: "seen[value] = seen.get(value, 0) + 1", LineStart: 6, LineEnd: 6},
			},
			Alternatives: []AltSolution{
				{Name: "Count first, then multiply", Code: "from collections import Counter\n\n\ndef countKDifference(nums, k):\n    counts = Counter(nums)\n    return sum(count * counts.get(value + k, 0) for value, count in counts.items())", TimeComplexity: "O(n)", SpaceComplexity: "O(n)", Explanation: "Count every value, then each distinct value v pairs with each copy of v + k: count(v) * count(v + k) pairs. Only the larger partner is looked up, so each pair is counted once."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why are x - k and x + k different values, and why does it matter?", Answer: "They differ by 2k, which is positive because k >= 1. If k were 0 they would be the same value and the code would add its count twice, so k = 0 needs a single lookup (and counts pairs of equal values)."},
			{Prompt: "How does this differ from Two Sum, which stores an index?", Answer: "Two Sum needs one partner, so the index is enough and it can stop at the first match. Here every earlier partner is a separate pair, and the answer is a number, so the map has to count occurrences and the loop runs to the end."},
		},
	},
	{
		ID:              "drill-rotate-array",
		Number:          189,
		Kind:            KindDrill,
		Order:           6,
		Title:           "Rotate Array in Place",
		Slug:            "drill-rotate-array",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Two Pointers", "Math", "Foundations"},
		RelatedChapters: []int{2, 1},
		Description:     "Rotate the array `nums` to the right by `k` steps, **in place**: modify `nums` itself and return nothing. A right rotation by 1 moves the last element to the front and shifts the others one place right; rotating by `k` does that `k` times. For example, `[1,2,3,4,5]` rotated by 2 becomes `[4,5,1,2,3]`.\n\nTwo facts do the work. First, rotating by `n` (the length) brings the array back to the start, so only `k % n` matters, and `k` can be far larger than `n`. Second, **reversing** does the rotation with O(1) extra space: reverse the whole array, then reverse the first `k % n` elements, then reverse the rest. For `[1,2,3,4,5]` with `k = 2`: reverse everything to `[5,4,3,2,1]`, reverse the first 2 to `[4,5,3,2,1]`, reverse the last 3 to `[4,5,1,2,3]`.\n\nThis is the in-place array handling (swap by index, no new list) that the array problems of the ladder, such as Product of Array Except Self, assume.",
		Constraints:     []string{"1 <= nums.length <= 100000", "-1000000000 <= nums[i] <= 1000000000", "0 <= k <= 1000000000 (k may be larger than nums.length)", "Modify nums in place and return nothing"},
		Examples: []Example{
			{Input: "nums = [1,2,3,4,5,6,7], k = 3", Output: "[5,6,7,1,2,3,4]", Explanation: "The last three elements move to the front in the same order."},
			{Input: "nums = [-1,-100,3,99], k = 2", Output: "[3,99,-1,-100]", Explanation: "The last two elements, 3 and 99, move to the front."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5, 6, 7}, "k": 3}, Expected: []int{5, 6, 7, 1, 2, 3, 4}, Note: "example 1: k smaller than the length"},
			{Input: map[string]interface{}{"nums": []int{-1, -100, 3, 99}, "k": 2}, Expected: []int{3, 99, -1, -100}, Note: "example 2: negative values, k is half the length"},
			{Input: map[string]interface{}{"nums": []int{1}, "k": 5}, Expected: []int{1}, Hidden: true, Note: "one element: any rotation leaves it unchanged, and k % 1 = 0"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}, "k": 0}, Expected: []int{1, 2, 3, 4}, Hidden: true, Note: "k = 0: nothing moves"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}, "k": 4}, Expected: []int{1, 2, 3, 4}, Hidden: true, Note: "k equal to the length is a full turn"},
			{Input: map[string]interface{}{"nums": []int{1, 2}, "k": 3}, Expected: []int{2, 1}, Hidden: true, Note: "k larger than the length: only k % n = 1 counts"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5}, "k": 7}, Expected: []int{4, 5, 1, 2, 3}, Hidden: true, Note: "k = 7 on 5 elements acts as 2; without the modulo it rotates 7 places or crashes"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5, 6}, "k": 1}, Expected: []int{6, 1, 2, 3, 4, 5}, Hidden: true, Note: "k = 1: the last element moves to the front (a left rotation would give [2,3,4,5,6,1])"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 2, 2}, "k": 2}, Expected: []int{2, 2, 1, 1}, Hidden: true, Note: "repeated values"},
			{Input: map[string]interface{}{"nums": ahRotateNums(), "k": 1000000000}, Expected: ahRotate(ahRotateNums(), 1000000000), Hidden: true, Timeout: 5000, Note: "100000 elements and k = 10^9: a loop that shifts one place k times is far too slow, and without k % n the reversal indexes are out of range"},
			{Input: map[string]interface{}{"nums": ahRotateNums(), "k": 50000}, Expected: ahRotate(ahRotateNums(), 50000), Hidden: true, Timeout: 5000, Note: "100000 elements with k = 50000: half a turn, so both blocks are large"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def rotate(nums, k):\n    # Rotate nums to the right by k steps, in place. Return nothing.\n    pass",
		Harness:         Harness{InPlace: "nums"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "After rotating a 5-element array by 7, where does the last element end up? Is it the same place as after rotating by 2, and what does that tell you about large k?"},
			{Level: 2, Type: "insight", Content: "Only k % n matters. For the rotation itself, notice that reversing the whole array puts the last k elements at the front but backwards, and the other elements at the back, also backwards. Reversing each of the two blocks fixes their order. Swapping elements by index keeps the extra space O(1). This is in-place array handling from Chapter 2.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Reduce: k = k % len(nums).\n2. Reverse all: swap nums[left] and nums[right] while left < right, over the whole array.\n3. Reverse the front block: indexes 0..k - 1.\n4. Reverse the back block: indexes k..n - 1."},
			{Level: 4, Type: "skeleton", Content: "def rotate(nums, k):\n    size = len(nums)\n    k %= ____\n\n    def reverse(left, right):\n        while left < right:\n            nums[left], nums[right] = nums[right], nums[left]\n            left += 1\n            right -= 1\n\n    reverse(0, size - 1)\n    reverse(0, ____)\n    reverse(____, size - 1)"},
		},
		Solution: Solution{
			Code:            "def rotate(nums, k):\n    size = len(nums)\n    k %= size\n\n    def reverse(left, right):\n        while left < right:\n            nums[left], nums[right] = nums[right], nums[left]\n            left += 1\n            right -= 1\n\n    reverse(0, size - 1)\n    reverse(0, k - 1)\n    reverse(k, size - 1)",
			Explanation:     "This drill is about changing an array in place rather than building a new one: swap by index, and reduce k first so a huge k costs nothing. Reverse-all then reverse-each-block moves every element to its rotated place with about 2n swaps and no extra list. It is not a hash-map idea; it is the array-manipulation skill that the later ladder problems (such as Product of Array Except Self, which fills the output array in place) expect.",
			BruteForce:      "Repeat k times: pop the last element and insert it at the front. Each insert shifts every element, so it is O(n * k); with k up to 10^9 and no reduction it never finishes, and even with k reduced it moves about n * (k % n) elements. The waste is moving every element one place at a time, when each one has a known final position.",
			KeyInsight:      "A right rotation by k is: reverse everything, then reverse the first k and the last n - k. Each element is moved a constant number of times, and k only matters modulo n.",
			WhyItWorks:      "Write the array as A B where B is the last k elements and A the first n - k. The rotation should give B A. Reversing the whole array gives rev(B) rev(A). Reversing the first k elements (the block rev(B)) turns it into B, and reversing the remaining n - k elements turns rev(A) into A, so the array is B A. Each reversal swaps pairs without needing extra space. k % n is enough because n steps of rotation return every element to its start; when k % n = 0, reverse(0, -1) does nothing and the two full reversals cancel.",
			CommonMistakes:  []string{"Skipping `k %= len(nums)`. With k larger than the array, the block reversals use indexes outside the array (or rotate a wrong number of places), and for k = 10^9 any loop over k is hopeless.", "Rotating to the left: reversing the first k elements and the rest in the wrong order, or slicing the wrong end. Check [1,2,3,4,5,6] with k = 1: the answer is [6,1,2,3,4,5].", "Writing `nums = nums[-k:] + nums[:-k]`. It rebinds the local name and leaves the caller's list unchanged; the problem asks for an in-place change (assign to nums[:] or swap elements).", "Shifting one place at a time k times without reducing k: it is O(n * k) and runs out of time when k is 10^9."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Reduce k", Explanation: "n rotations bring every element back, so only the remainder matters. It also makes k = 0 safe and keeps all indexes inside the array.", CodeSnippet: "size = len(nums)\nk %= size", LineStart: 2, LineEnd: 3},
				{Title: "Write a reversal by index", Explanation: "Swap the two ends and move inward until the pointers meet. It changes the list itself and uses no extra space.", CodeSnippet: "def reverse(left, right):\n    while left < right:\n        nums[left], nums[right] = nums[right], nums[left]\n        left += 1\n        right -= 1", LineStart: 5, LineEnd: 9},
				{Title: "Reverse everything", Explanation: "The last k elements are now at the front, but in reverse order, and the first n - k elements are at the back, also reversed.", CodeSnippet: "reverse(0, size - 1)", LineStart: 11, LineEnd: 11},
				{Title: "Fix each block", Explanation: "Reversing the front k elements and the back n - k elements restores the order inside each block. For [1,2,3,4,5] with k = 2: [5,4,3,2,1] becomes [4,5,3,2,1], then [4,5,1,2,3].", CodeSnippet: "reverse(0, k - 1)\nreverse(k, size - 1)", LineStart: 12, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{Name: "Slice assignment", Code: "def rotate(nums, k):\n    k %= len(nums)\n    if k:\n        nums[:] = nums[-k:] + nums[:-k]", TimeComplexity: "O(n)", SpaceComplexity: "O(n)", Explanation: "Builds the rotated list and writes it back into the same list with nums[:] (so it is still in place from the caller's view). Reduce k first so that it is below the length; the `if k` check just skips the work for a rotation of 0."},
				{Name: "Move one at a time", Code: "def rotate(nums, k):\n    for _ in range(k % len(nums)):\n        nums.insert(0, nums.pop())", TimeComplexity: "O(n * k)", SpaceComplexity: "O(1)", Explanation: "The definition of rotating by k, done literally. With k reduced it passes here because list.insert is a fast memory move, but each insert still shifts the whole list, so the work is n times k."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it enough to use k % n?", Answer: "Rotating by n places moves every element all the way around and back to where it started, so rotating by k and by k % n give the same array. Taking the remainder also keeps k below n, so every index the code uses is valid."},
			{Prompt: "What goes wrong if you build a new list and assign it to the parameter name?", Answer: "Assigning to the name nums only rebinds the function's local variable; the caller's list is untouched. To change the caller's list you must mutate it: swap elements, or write through nums[:] = ... ."},
		},
	},
}

// ahLCG is a small deterministic generator so the large cases are the same on
// every run.
func ahLCG(seed uint64) func() int {
	state := seed
	return func() int {
		state = state*6364136223846793005 + 1442695040888963407
		return int(state >> 33)
	}
}

func ahFreqBig() []int {
	next := ahLCG(1)
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = next()%50000 - 25000
	}
	return nums
}

func ahFreqExpected(nums []int) [][]int {
	counts := map[int]int{}
	var order []int
	for _, v := range nums {
		if counts[v] == 0 {
			order = append(order, v)
		}
		counts[v]++
	}
	pairs := make([][]int, 0, len(order))
	for _, v := range order {
		pairs = append(pairs, []int{v, counts[v]})
	}
	return pairs
}

// ahUniqBig is three long runs (a, b, c) and one d, the only unique letter.
func ahUniqBig() string {
	return strings.Repeat("a", 33000) + strings.Repeat("b", 33000) + strings.Repeat("c", 33000) + "d"
}

func ahFirstUniq(s string) int {
	var counts [256]int
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if counts[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func ahRangeNums() []int {
	next := ahLCG(2)
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = next()%2001 - 1000
	}
	return nums
}

// ahRangeQueries returns wide ranges (from near the start to near the end).
func ahRangeQueries() [][]int {
	next := ahLCG(3)
	queries := make([][]int, 100000)
	for i := range queries {
		queries[i] = []int{next() % 1000, 99999 - next()%1000}
	}
	return queries
}

func ahRangeSums(nums []int, queries [][]int) []int {
	prefix := make([]int, len(nums)+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + v
	}
	out := make([]int, len(queries))
	for i, q := range queries {
		out[i] = prefix[q[1]+1] - prefix[q[0]]
	}
	return out
}

func ahZeros(n int) []int { return make([]int, n) }

func ahSubBig() []int {
	next := ahLCG(4)
	nums := make([]int, 50000)
	for i := range nums {
		nums[i] = next()%11 - 5
	}
	return nums
}

func ahSubCount(nums []int, k int) int {
	seen := map[int]int{0: 1}
	prefix, count := 0, 0
	for _, v := range nums {
		prefix += v
		count += seen[prefix-k]
		seen[prefix]++
	}
	return count
}

func ahDiffBig() []int {
	next := ahLCG(5)
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = next() % 1000
	}
	return nums
}

func ahDiffCount(nums []int, k int) int {
	seen := map[int]int{}
	pairs := 0
	for _, v := range nums {
		pairs += seen[v-k] + seen[v+k]
		seen[v]++
	}
	return pairs
}

func ahRotateNums() []int {
	next := ahLCG(6)
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = next()%2000001 - 1000000
	}
	return nums
}

func ahRotate(nums []int, k int) []int {
	n := len(nums)
	k %= n
	out := make([]int, n)
	for i, v := range nums {
		out[(i+k)%n] = v
	}
	return out
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, ArraysHashingDrills...)
}
