package practice

// ArraysHashingProblems is the arrays-hashing ladder (chapter 2). Order runs
// from the plainest use of a set to problems that combine several ideas:
// set -> counting -> seen-so-far map -> grouping by a key -> bucketing by
// count -> several sets at once -> a self-describing format -> prefix and
// suffix products -> "only start at the start of a run".
var ArraysHashingProblems = []*Problem{
	{
		ID:              "contains-duplicate",
		Number:          217,
		Order:           1,
		Title:           "Contains Duplicate",
		Slug:            "contains-duplicate",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "Sorting"},
		RelatedChapters: []int{2, 1},
		Description: `You are given a list of integers nums. Return true if at least one value appears two or more times in the list. Return false if every value in the list is different from every other value.

Positions don't matter: the two equal values can be next to each other or far apart.`,
		Constraints: []string{
			"1 <= nums.length <= 10^5",
			"-10^9 <= nums[i] <= 10^9",
		},
		Examples: []Example{
			{Input: "nums = [1,2,3,1]", Output: "true", Explanation: "The value 1 appears at index 0 and again at index 3."},
			{Input: "nums = [1,2,3,4]", Output: "false", Explanation: "All four values are different."},
			{Input: "nums = [1,1,1,3,3,4,3,2,4,2]", Output: "true"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 1}}, Expected: true, Note: "the first value comes back at the end"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}}, Expected: false, Note: "all values distinct"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, Expected: true, Note: "many repeats"},
			{Input: map[string]interface{}{"nums": []int{7}}, Expected: false, Hidden: true, Note: "single element can't repeat"},
			{Input: map[string]interface{}{"nums": []int{-1, -1}}, Expected: true, Hidden: true, Note: "negative duplicate, smallest array with a repeat"},
			{Input: map[string]interface{}{"nums": []int{0, -1, 1, -2, 2}}, Expected: false, Hidden: true, Note: "a value and its negative are different values"},
			{Input: map[string]interface{}{"nums": []int{5, 4, 3, 2, 1, 5}}, Expected: true, Hidden: true, Note: "repeat far apart (not adjacent)"},
			{Input: map[string]interface{}{"nums": []int{1000000000, -1000000000, 999999999, 1000000000}}, Expected: true, Hidden: true, Note: "extreme values from the constraints"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def containsDuplicate(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When you reach a number, what would you need to know about the numbers you have already walked past to answer \"have I seen this before?\" instantly?"},
			{Level: 2, Type: "insight", Content: "Keep a set of every value seen so far. Checking \"x in seen\" on a set takes O(1) time on average because the set hashes x straight to where it would be stored, instead of scanning. This is the seen-so-far pattern from Chapter 2.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Choose what to remember: the set of values already passed.\n2. Look up before you store: for each number, ask whether it is already in the set; if so, you've found a duplicate.\n3. Record the current item: otherwise add it to the set.\n4. Return the answer: if the loop finishes, nothing repeated."},
			{Level: 4, Type: "skeleton", Content: `def containsDuplicate(nums):
    seen = ____
    for num in nums:
        if ____:
            return True
        ____
    return False`},
		},
		Solution: Solution{
			Code: `def containsDuplicate(nums):
    seen = set()
    for num in nums:
        if num in seen:
            return True
        seen.add(num)
    return False`,
			Explanation: "Walk the list once while keeping a set of the values already passed. The moment a value is already in the set, it is a repeat. This is the plainest form of Chapter 2's seen-so-far pattern, and every later problem in this ladder builds on it.",
			BruteForce:  "Compare every pair: for each i, check every j > i for nums[i] == nums[j]. That is n(n-1)/2 comparisons, about 5 billion when n = 100,000, so O(n^2) time. The waste: when you reach index j, you compare it against every earlier value one by one, even though all you need to know is whether its value is among them.",
			KeyInsight:  "The question \"is this value among the earlier ones?\" is a membership test, and a set answers membership in O(1) on average. Store the past in a set and each number costs one lookup instead of a scan.",
			WhyItWorks:  "Invariant: before processing index i, seen holds exactly the values nums[0..i-1]. If nums[i] repeats some earlier value, that value is in seen, so we return True at the first repeat. If we finish the loop, every number was checked against all numbers before it and none matched, so all values are distinct. Every pair (i, j) with i < j is covered, because pair (i, j) is checked when we reach j.",
			CommonMistakes: []string{
				"Adding the number to the set before checking it, so every number \"finds itself\" and the function always returns True. Check first, then add.",
				"Using a list instead of a set for seen. \"x in some_list\" scans the list, so the whole thing is O(n^2) again. Use set() for O(1) average lookups.",
				"Sorting and comparing only nums[i] with nums[i+1] is correct, but forgetting to sort first misses repeats that aren't adjacent, like [5,4,3,2,1,5].",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Choose what to remember", Explanation: "We only need to know which values we've passed, not where they were, so a set (not a dict) is enough.", CodeSnippet: "seen = set()", LineStart: 2, LineEnd: 2},
				{Title: "Look up before you store", Explanation: "Check the current number against the past before adding it. Checking after adding would always succeed.", CodeSnippet: "if num in seen:\n    return True", LineStart: 4, LineEnd: 5},
				{Title: "Record the current item", Explanation: "The number wasn't a repeat, so it joins the past for the numbers that come after it.", CodeSnippet: "seen.add(num)", LineStart: 6, LineEnd: 6},
				{Title: "Return the answer", Explanation: "Reaching the end means no number matched anything before it, so every value is distinct.", CodeSnippet: "return False", LineStart: 7, LineEnd: 7},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: compare every pair",
					Code: `def containsDuplicate(nums):
    for i in range(len(nums)):
        for j in range(i + 1, len(nums)):
            if nums[i] == nums[j]:
                return True
    return False`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Checks all n(n-1)/2 pairs. Correct, and uses no extra memory, but too slow for n = 10^5.",
				},
				{
					Name: "Sort, then compare neighbours",
					Code: `def containsDuplicate(nums):
    ordered = sorted(nums)
    for i in range(1, len(ordered)):
        if ordered[i] == ordered[i - 1]:
            return True
    return False`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "After sorting, equal values sit next to each other, so one pass over neighbours finds them. A good choice when memory for a set is tight and you may sort in place.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the membership check come before seen.add(num)?", Answer: "If you add first, the set already contains num when you check, so every number looks like a repeat and the function returns True on the first element. The set must describe only the numbers before the current one."},
			{Prompt: "What if the question asked you to return the first value whose second copy appears earliest? What would change?", Answer: "Nothing about the data structure: the first time the check \"num in seen\" succeeds, num is exactly the value whose second copy appears earliest, because we scan left to right. Return num instead of True."},
			{Prompt: "When would you prefer the sorting approach over the set?", Answer: "When memory matters more than speed: sorting in place needs O(1) extra space (besides the sort itself) and costs O(n log n) time, while the set needs O(n) extra memory for O(n) time."},
		},
	},
	{
		ID:              "valid-anagram",
		Number:          242,
		Order:           2,
		Title:           "Valid Anagram",
		Slug:            "valid-anagram",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Hash Table", "String", "Sorting"},
		RelatedChapters: []int{2},
		Description: `Two strings are anagrams of each other when one can be turned into the other by rearranging its letters, using every letter exactly as many times as it appears. For example "listen" and "silent" are anagrams, but "aab" and "abb" are not (one has two a's, the other has one).

Given two strings s and t, return true if t is an anagram of s, and false otherwise.`,
		Constraints: []string{
			"1 <= s.length, t.length <= 5 * 10^4",
			"s and t consist of lowercase English letters",
		},
		Examples: []Example{
			{Input: `s = "anagram", t = "nagaram"`, Output: "true", Explanation: "Both strings contain three a's, one n, one g, one r and one m."},
			{Input: `s = "rat", t = "car"`, Output: "false", Explanation: "s has a t where t has a c."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "anagram", "t": "nagaram"}, Expected: true, Note: "same letters rearranged"},
			{Input: map[string]interface{}{"s": "rat", "t": "car"}, Expected: false, Note: "different letters"},
			{Input: map[string]interface{}{"s": "aab", "t": "abb"}, Expected: false, Hidden: true, Note: "same set of letters but different counts"},
			{Input: map[string]interface{}{"s": "a", "t": "ab"}, Expected: false, Hidden: true, Note: "different lengths"},
			{Input: map[string]interface{}{"s": "ab", "t": "a"}, Expected: false, Hidden: true, Note: "t shorter than s"},
			{Input: map[string]interface{}{"s": "listen", "t": "silent"}, Expected: true, Hidden: true, Note: "classic anagram pair"},
			{Input: map[string]interface{}{"s": "a", "t": "a"}, Expected: true, Hidden: true, Note: "single identical letter"},
			{Input: map[string]interface{}{"s": "aaaaaaaaab", "t": "aaaaaaaaba"}, Expected: true, Hidden: true, Note: "many repeats of one letter"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def isAnagram(s, t):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you were given two shuffled piles of letter tiles, how would you check they hold exactly the same tiles without trying every arrangement?"},
			{Level: 2, Type: "insight", Content: "Two strings are anagrams exactly when every letter appears the same number of times in both. Count the letters of s in a dictionary, then use up those counts with the letters of t. This is counting with a dict, from Chapter 2. Contains Duplicate asked \"seen or not?\"; here you need \"how many times?\".", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Return the answer early if the lengths differ: then they can't be anagrams.\n2. Count each item: build letter -> count for s.\n3. Compare the counts: for each letter of t, if its count is already 0 there is an extra letter, so return False; otherwise take one away.\n4. Return the answer: True if every letter of t was matched."},
			{Level: 4, Type: "skeleton", Content: `def isAnagram(s, t):
    if ____:
        return False
    counts = {}
    for ch in s:
        counts[ch] = ____
    for ch in t:
        if ____:
            return False
        counts[ch] -= 1
    return True`},
		},
		Solution: Solution{
			Code: `def isAnagram(s, t):
    if len(s) != len(t):
        return False
    counts = {}
    for ch in s:
        counts[ch] = counts.get(ch, 0) + 1
    for ch in t:
        if counts.get(ch, 0) == 0:
            return False
        counts[ch] -= 1
    return True`,
			Explanation: "Count every letter of s in a dictionary, then spend those counts on the letters of t. Builds on Contains Duplicate: instead of remembering only whether a value was seen (a set), we remember how many times (a dict from value to count).",
			BruteForce:  "Sort both strings and compare them: O(n log n). An even slower idea is, for each letter of t, find and delete a matching letter in s, which rescans s every time: O(n^2). The waste is that the order of letters doesn't matter at all, yet both approaches spend work arranging or searching through positions.",
			KeyInsight:  "Anagrams are about how many of each letter, not where they are. A count per letter captures exactly that and takes one pass to build.",
			WhyItWorks:  "After the first loop, counts[x] is the number of times x appears in s. The second loop removes one x per occurrence in t and stops if t ever needs an x that s doesn't have left. If it finishes, every letter of t was matched to a distinct letter of s, so t uses no letter more often than s. Since the lengths are equal, t can't use any letter less often either (the total would fall short), so the counts match exactly.",
			CommonMistakes: []string{
				"Comparing set(s) == set(t). Sets drop the counts, so \"aab\" and \"abb\" look equal. Count occurrences, don't just record presence.",
				"Skipping the length check and only testing that t's letters are available in s. Then \"ab\" vs \"a\" returns True. Either check lengths first or verify all counts end at 0.",
				"Writing counts[ch] += 1 for a key that doesn't exist yet, which raises KeyError. Use counts.get(ch, 0) + 1 or collections.Counter.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Return the answer early", Explanation: "Different lengths can never be anagrams. This check also makes the final step correct.", CodeSnippet: "if len(s) != len(t):\n    return False", LineStart: 2, LineEnd: 3},
				{Title: "Count each item", Explanation: "One pass over s builds letter -> how many times it appears. At most 26 keys exist, so the dictionary has constant size.", CodeSnippet: "for ch in s:\n    counts[ch] = counts.get(ch, 0) + 1", LineStart: 5, LineEnd: 6},
				{Title: "Compare the counts", Explanation: "Each letter of t uses up one copy from s. A letter with nothing left means t has more of it than s does.", CodeSnippet: "if counts.get(ch, 0) == 0:\n    return False\ncounts[ch] -= 1", LineStart: 8, LineEnd: 10},
				{Title: "Return the answer", Explanation: "Every letter of t was matched and the lengths are equal, so the counts are identical.", CodeSnippet: "return True", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Sort both strings",
					Code: `def isAnagram(s, t):
    return sorted(s) == sorted(t)`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Sorting puts equal letters together, so anagrams sort to the same string. Short and clear, but slower than counting.",
				},
				{
					Name: "Counter",
					Code: `from collections import Counter

def isAnagram(s, t):
    return Counter(s) == Counter(t)`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Counter builds the same letter -> count dictionary in one call; two Counters are equal when every count matches.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the space O(1) and not O(n), even though we build a dictionary?", Answer: "The keys are lowercase English letters, so the dictionary never holds more than 26 entries no matter how long the strings are. A bound that doesn't grow with n is O(1)."},
			{Prompt: "What if the strings could contain any Unicode character? What changes?", Answer: "The dictionary approach works unchanged because it only uses characters as keys; only the space bound changes to O(k) for k distinct characters. A fixed array of 26 counters would no longer work, which is one reason to prefer a dict."},
			{Prompt: "Why does the length check make it safe to stop counting once t's letters are all matched?", Answer: "Matching every letter of t means t uses each letter at most as often as s. If some letter were used less often, t would be shorter than s, which the length check ruled out. So the counts must be exactly equal."},
		},
	},
	{
		ID:              "two-sum",
		Number:          1,
		Order:           3,
		Title:           "Two Sum",
		Slug:            "two-sum",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table"},
		RelatedChapters: []int{2, 3},
		Description: `You are given a list of integers nums and an integer target. Exactly one pair of different positions i and j has nums[i] + nums[j] == target. Return those two positions as a list [i, j].

You may not use the same position twice (but two different positions may hold the same value). The two indices may be returned in either order.`,
		Constraints: []string{
			"2 <= nums.length <= 10^4",
			"-10^9 <= nums[i] <= 10^9",
			"-10^9 <= target <= 10^9",
			"Exactly one valid pair exists",
		},
		Examples: []Example{
			{Input: "nums = [2,7,11,15], target = 9", Output: "[0,1]", Explanation: "nums[0] + nums[1] = 2 + 7 = 9."},
			{Input: "nums = [3,2,4], target = 6", Output: "[1,2]", Explanation: "2 + 4 = 6. Using the 3 at index 0 twice would also make 6, but a position can't be used twice."},
			{Input: "nums = [3,3], target = 6", Output: "[0,1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 7, 11, 15}, "target": 9}, Expected: []int{0, 1}, Note: "pair at the start"},
			{Input: map[string]interface{}{"nums": []int{3, 2, 4}, "target": 6}, Expected: []int{1, 2}, Note: "3 + 3 would be 6 but there is only one 3: don't use an element twice"},
			{Input: map[string]interface{}{"nums": []int{3, 3}, "target": 6}, Expected: []int{0, 1}, Note: "two equal values form the pair"},
			{Input: map[string]interface{}{"nums": []int{-1, -2, -3, -4, -5}, "target": -8}, Expected: []int{2, 4}, Hidden: true, Note: "all negative numbers"},
			{Input: map[string]interface{}{"nums": []int{0, 4, 3, 0}, "target": 0}, Expected: []int{0, 3}, Hidden: true, Note: "target 0 made from two zeros"},
			{Input: map[string]interface{}{"nums": []int{1, 5, 9, 13, 2, 8}, "target": 21}, Expected: []int{3, 5}, Hidden: true, Note: "pair far apart, several near misses"},
			{Input: map[string]interface{}{"nums": []int{-3, 4, 3, 90}, "target": 0}, Expected: []int{0, 2}, Hidden: true, Note: "a value plus its negative"},
			{Input: map[string]interface{}{"nums": []int{1000000000, -1000000000, 5}, "target": 0}, Expected: []int{0, 1}, Hidden: true, Note: "extreme values from the constraints"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def twoSum(nums, target):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When you stand on a number x, which exact value would complete the pair? What would you need to remember about the numbers you've already walked past to find it at once?"},
			{Level: 2, Type: "insight", Content: "For each x, the partner must be target - x (its complement). Keep a dictionary from each value you've passed to its index; then \"has the partner appeared yet?\" is one O(1) lookup. This is the seen-so-far pattern from Chapter 2, upgraded from Contains Duplicate's set to a dict because now you need the partner's position too.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Choose what to remember: a dict value -> index for the numbers already passed.\n2. Look up before you store: compute complement = target - x and check whether it is in the dict; if so, return both indices.\n3. Record the current item: otherwise store x -> its index.\n4. Return the answer: the problem guarantees a pair, so the loop always returns."},
			{Level: 4, Type: "skeleton", Content: `def twoSum(nums, target):
    index_of = {}
    for i, num in enumerate(nums):
        complement = ____
        if ____:
            return [____, i]
        index_of[num] = i
    return []`},
		},
		Solution: Solution{
			Code: `def twoSum(nums, target):
    index_of = {}  # value -> index where we saw it
    for i, num in enumerate(nums):
        complement = target - num
        if complement in index_of:
            return [index_of[complement], i]
        index_of[num] = i
    return []`,
			Explanation: "Walk the list once. For each number, the value that would complete the pair is target - number; look it up in a dictionary of values already passed. Builds on Contains Duplicate: the same \"look up, then store\" loop, but the dictionary stores each value's index so we can return positions, and we look up the complement instead of the number itself.",
			BruteForce:  "Try every pair (i, j) with i < j and test nums[i] + nums[j] == target. That is n(n-1)/2 pairs, about 50 million for n = 10,000: O(n^2). The waste: for a fixed j we already know the one value that could work (target - nums[j]), yet we scan all earlier numbers to look for it.",
			KeyInsight:  "Once you fix one number, its partner is fully determined: target - x. Searching for a known value is a lookup, and a hash map does lookups in O(1) on average.",
			WhyItWorks:  "Invariant: when we reach index i, index_of maps every value in nums[0..i-1] to an index where it occurs. Let the answer be (a, b) with a < b. When the loop reaches b, nums[a] = target - nums[b] is already stored, so we find it and return [a, b]. We never return a wrong pair because we only return when the stored value plus num equals target. Checking before storing guarantees the two indices differ, which handles [3,2,4] with target 6 correctly.",
			CommonMistakes: []string{
				"Storing all numbers in the dict first, then looking up target - x. For [3,2,4], target 6, the 3 finds itself and returns [0,0]. Either look up before storing or check that the indices differ.",
				"Returning the values [complement, num] instead of their indices. The problem asks for positions; that's why the map stores value -> index.",
				"Sorting the array first to use two pointers. That destroys the original indices; you'd have to sort (value, index) pairs. The hash map avoids this (see Two Sum II for when sorting is given).",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Choose what to remember", Explanation: "We need to find a value and report where it was, so the dictionary maps value -> index.", CodeSnippet: "index_of = {}  # value -> index where we saw it", LineStart: 2, LineEnd: 2},
				{Title: "Look up before you store", Explanation: "The partner of num is fixed: target - num. If it was seen earlier, we have the pair. Looking up before storing prevents pairing a number with itself.", CodeSnippet: "complement = target - num\nif complement in index_of:\n    return [index_of[complement], i]", LineStart: 4, LineEnd: 6},
				{Title: "Record the current item", Explanation: "num might be the partner of a later number, so store where it is.", CodeSnippet: "index_of[num] = i", LineStart: 7, LineEnd: 7},
				{Title: "Return the answer", Explanation: "With exactly one valid pair guaranteed the loop always returns; the final line only keeps the function total.", CodeSnippet: "return []", LineStart: 8, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: try every pair",
					Code: `def twoSum(nums, target):
    for i in range(len(nums)):
        for j in range(i + 1, len(nums)):
            if nums[i] + nums[j] == target:
                return [i, j]
    return []`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Checks all pairs. Fine for tiny inputs and a good first draft, but each j rescans everything before it.",
				},
				{
					Name: "Sort (value, index) pairs, then two pointers",
					Code: `def twoSum(nums, target):
    pairs = sorted((value, index) for index, value in enumerate(nums))
    left, right = 0, len(pairs) - 1
    while left < right:
        total = pairs[left][0] + pairs[right][0]
        if total == target:
            return [pairs[left][1], pairs[right][1]]
        if total < target:
            left += 1
        else:
            right -= 1
    return []`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Sorting lets the two-pointer method from Chapter 3 work, but we must carry the original indices along. Slower than the hash map because of the sort.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why do we look up the complement before storing the current number?", Answer: "So a number can never pair with itself. In [3,2,4] with target 6, storing 3 first and then looking up 6 - 3 = 3 would find index 0 and return [0,0]. Looking up first means the dict only holds earlier positions."},
			{Prompt: "What if the input were already sorted and you had to use O(1) extra space? Which pattern would you use?", Answer: "Two pointers (Two Sum II, Chapter 3): start at both ends; if the sum is too small move the left pointer right, if too big move the right pointer left. Sorted order tells you which pointer to move, so no dictionary is needed."},
			{Prompt: "What if duplicates exist and the same value is stored twice: does overwriting index_of[num] break anything?", Answer: "No. If two equal values form the answer (like [3,3]), the second one finds the first before overwriting. If they don't form the answer, any stored index for that value is a valid partner position, so overwriting is harmless."},
		},
	},
	{
		ID:              "group-anagrams",
		Number:          49,
		Order:           4,
		Title:           "Group Anagrams",
		Slug:            "group-anagrams",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "String", "Sorting"},
		RelatedChapters: []int{2},
		Description: `Two words are anagrams when one is a rearrangement of the other's letters, using each letter the same number of times (see Valid Anagram).

Given a list of strings strs, put the words into groups so that two words share a group exactly when they are anagrams of each other. Return the list of groups. The groups may be in any order, and the words inside a group may be in any order. Keep duplicates: if a word appears twice it appears twice in its group.`,
		Constraints: []string{
			"1 <= strs.length <= 10^4",
			"0 <= strs[i].length <= 100",
			"strs[i] consists of lowercase English letters",
		},
		Examples: []Example{
			{Input: `strs = ["eat","tea","tan","ate","nat","bat"]`, Output: `[["bat"],["nat","tan"],["ate","eat","tea"]]`, Explanation: "eat, tea and ate use the letters a, e, t once each; tan and nat use a, n, t; bat has no partner."},
			{Input: `strs = [""]`, Output: `[[""]]`},
			{Input: `strs = ["a"]`, Output: `[["a"]]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"strs": []string{"eat", "tea", "tan", "ate", "nat", "bat"}}, Expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}, Note: "three groups of different sizes"},
			{Input: map[string]interface{}{"strs": []string{""}}, Expected: [][]string{{""}}, Note: "one empty string"},
			{Input: map[string]interface{}{"strs": []string{"a"}}, Expected: [][]string{{"a"}}, Note: "one single-letter string"},
			{Input: map[string]interface{}{"strs": []string{"", ""}}, Expected: [][]string{{"", ""}}, Hidden: true, Note: "two empty strings are anagrams of each other"},
			{Input: map[string]interface{}{"strs": []string{"abc", "bca", "cab", "xyz"}}, Expected: [][]string{{"abc", "bca", "cab"}, {"xyz"}}, Hidden: true, Note: "one big group and one singleton"},
			{Input: map[string]interface{}{"strs": []string{"aab", "aba", "abb", "bab"}}, Expected: [][]string{{"aab", "aba"}, {"abb", "bab"}}, Hidden: true, Note: "same letters, different counts must not be grouped"},
			{Input: map[string]interface{}{"strs": []string{"ab", "ba", "ab"}}, Expected: [][]string{{"ab", "ba", "ab"}}, Hidden: true, Note: "duplicate strings stay in the group twice"},
			{Input: map[string]interface{}{"strs": []string{"bdddddddddd", "bbbbbbbbbbc"}}, Expected: [][]string{{"bdddddddddd"}, {"bbbbbbbbbbc"}}, Hidden: true, Note: "letter counts above 9 (a joined-digit key without separators would collide)"},
		},
		TimeComplexity:  "O(n * k)",
		SpaceComplexity: "O(n * k)",
		StarterCode:     "def groupAnagrams(strs):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered-nested"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Valid Anagram compared two words. With thousands of words, comparing every pair is slow. Could you compute something from a single word that is the same for all of its anagrams and different for everything else?"},
			{Level: 2, Type: "insight", Content: "Give each word a signature that ignores letter order: its 26 letter counts (or its letters sorted). Anagrams have identical signatures, so use the signature as a dictionary key and append each word to the list under its key. This is grouping by a key with a hash map (Chapter 2).", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Choose what to remember: a dict from signature -> list of words with that signature.\n2. Count each item: for each word, build its 26 letter counts and turn them into a tuple (lists can't be dict keys).\n3. Record the current item: append the word to the list under that key, creating the list the first time.\n4. Return the answer: the dict's values are the groups."},
			{Level: 4, Type: "skeleton", Content: `def groupAnagrams(strs):
    groups = {}
    for word in strs:
        counts = [0] * 26
        for ch in word:
            counts[____] += 1
        key = ____
        if key not in groups:
            groups[key] = []
        ____
    return list(groups.values())`},
		},
		Solution: Solution{
			Code: `def groupAnagrams(strs):
    groups = {}  # letter-count signature -> words with it
    for word in strs:
        counts = [0] * 26
        for ch in word:
            counts[ord(ch) - ord('a')] += 1
        key = tuple(counts)
        if key not in groups:
            groups[key] = []
        groups[key].append(word)
    return list(groups.values())`,
			Explanation: "Turn each word into a signature that is identical for anagrams (its letter counts as a tuple) and use it as a dictionary key; each key collects its group. Builds on Valid Anagram (counting letters) and adds one idea: a computed key lets a hash map group items, instead of comparing items pairwise.",
			BruteForce:  "For each word, compare it against one representative of every existing group with an anagram check, and start a new group if none match. With n words and up to n groups that is O(n^2) comparisons of O(k) each, where k is the word length: O(n^2 * k). The waste: we re-derive \"which letters, how many\" for the same words over and over, and search all groups linearly to find the matching one.",
			KeyInsight:  "Whether two words are anagrams depends only on their letter counts. Compute those counts once per word, make them hashable, and the hash map jumps straight to the right group.",
			WhyItWorks:  "Two words are anagrams if and only if their 26 letter counts are equal, so the tuple of counts is a perfect signature: equal for every pair of anagrams, different for every pair of non-anagrams. Grouping by an exact signature therefore puts each word with precisely its anagrams. Each word costs O(k) to count plus O(1) average for the dict lookup (the key has fixed length 26), so the total is O(n * k).",
			CommonMistakes: []string{
				"Using the counts list itself as a key: lists are mutable and unhashable, so Python raises TypeError. Convert with tuple(counts).",
				"Building the key by joining counts without separators, like \"\".join(map(str, counts)). Counts 1 and 11 can then run together and collide (\"bdddddddddd\" vs \"bbbbbbbbbbc\"). Use a tuple, or put a separator such as '#' between counts.",
				"Grouping by set(word) or sorted(set(word)), which throws away how many times each letter appears, so \"aab\" and \"abb\" land together.",
			},
			TimeComplexity:  "O(n * k)",
			SpaceComplexity: "O(n * k)",
			Walkthrough: []WalkthroughStep{
				{Title: "Choose what to remember", Explanation: "A dictionary from signature to the list of words sharing it. Each list becomes one group.", CodeSnippet: "groups = {}  # letter-count signature -> words with it", LineStart: 2, LineEnd: 2},
				{Title: "Count each item", Explanation: "Count letters into 26 slots; ord(ch) - ord('a') maps 'a'..'z' to 0..25. The tuple is hashable and equal exactly for anagrams.", CodeSnippet: "counts = [0] * 26\nfor ch in word:\n    counts[ord(ch) - ord('a')] += 1\nkey = tuple(counts)", LineStart: 4, LineEnd: 7},
				{Title: "Record the current item", Explanation: "Create the group the first time a signature appears, then add the word to it.", CodeSnippet: "if key not in groups:\n    groups[key] = []\ngroups[key].append(word)", LineStart: 8, LineEnd: 10},
				{Title: "Return the answer", Explanation: "The dictionary's values are exactly the groups; their order doesn't matter.", CodeSnippet: "return list(groups.values())", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Sorted word as the key",
					Code: `def groupAnagrams(strs):
    groups = {}
    for word in strs:
        key = "".join(sorted(word))
        groups.setdefault(key, []).append(word)
    return list(groups.values())`,
					TimeComplexity:  "O(n * k log k)",
					SpaceComplexity: "O(n * k)",
					Explanation:     "Sorting a word's letters also gives an order-free signature (\"eat\", \"tea\" -> \"aet\"). Simpler to write; slightly slower for long words because of the sort.",
				},
				{
					Name: "Brute force: compare with each group",
					Code: `def groupAnagrams(strs):
    groups = []
    for word in strs:
        for group in groups:
            if sorted(group[0]) == sorted(word):
                group.append(word)
                break
        else:
            groups.append([word])
    return groups`,
					TimeComplexity:  "O(n^2 * k log k)",
					SpaceComplexity: "O(n * k)",
					Explanation:     "Checks each word against one member of every group. Correct but quadratic in the number of words.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the key be a tuple and not a list?", Answer: "Dictionary keys must be hashable, and a key's hash must never change. Lists are mutable, so Python refuses to hash them. A tuple is immutable, so it can be hashed and used as a key."},
			{Prompt: "What if the words could contain any characters, not just a to z? Which signature would you use?", Answer: "The 26-slot count array no longer fits. Use the sorted word (\"\".join(sorted(word))) or a frozenset-free count such as tuple(sorted(Counter(word).items())); both are order-free and keep counts."},
			{Prompt: "How is this the same idea as Two Sum's dictionary, and how is it different?", Answer: "Both use a hash map to jump straight to the related items instead of scanning. In Two Sum the key is the value itself and we look up a partner; here the key is a computed signature and we collect every item that shares it."},
		},
	},
	{
		ID:              "top-k-frequent-elements",
		Number:          347,
		Order:           5,
		Title:           "Top K Frequent Elements",
		Slug:            "top-k-frequent-elements",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "Bucket Sort", "Heap (Priority Queue)", "Counting"},
		RelatedChapters: []int{2, 11},
		Description: `You are given a list of integers nums and an integer k. The frequency of a value is how many times it appears in nums. Return the k distinct values with the highest frequencies.

The input is chosen so that the answer is unique (there is never a tie at the cut-off). Return the k values in any order.

Follow-up: aim for better than O(n log n) time.`,
		Constraints: []string{
			"1 <= nums.length <= 10^5",
			"-10^4 <= nums[i] <= 10^4",
			"1 <= k <= number of distinct values in nums",
			"The answer is unique",
		},
		Examples: []Example{
			{Input: "nums = [1,1,1,2,2,3], k = 2", Output: "[1,2]", Explanation: "1 appears 3 times, 2 appears twice and 3 once, so the two most frequent values are 1 and 2."},
			{Input: "nums = [1], k = 1", Output: "[1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 2, 2, 3}, "k": 2}, Expected: []int{1, 2}, Note: "two most frequent of three values"},
			{Input: map[string]interface{}{"nums": []int{1}, "k": 1}, Expected: []int{1}, Note: "single element"},
			{Input: map[string]interface{}{"nums": []int{4, 4, 4, 4}, "k": 1}, Expected: []int{4}, Hidden: true, Note: "only one distinct value"},
			{Input: map[string]interface{}{"nums": []int{5, 3, 1, 1, 1, 3, 73, 1}, "k": 1}, Expected: []int{1}, Hidden: true, Note: "most frequent value is not the largest or first"},
			{Input: map[string]interface{}{"nums": []int{-1, -1, 2, 2, 2, 3}, "k": 2}, Expected: []int{2, -1}, Hidden: true, Note: "negative numbers"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5, 6, 6}, "k": 1}, Expected: []int{6}, Hidden: true, Note: "winner appears only twice"},
			{Input: map[string]interface{}{"nums": []int{7, 7, 8, 8, 8, 9, 9, 9, 9}, "k": 3}, Expected: []int{9, 8, 7}, Hidden: true, Note: "k equals the number of distinct values"},
			{Input: map[string]interface{}{"nums": []int{3, 0, 1, 0}, "k": 1}, Expected: []int{0}, Hidden: true, Note: "zero is a valid value"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def topKFrequent(nums, k):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Once you know how often each value appears, you need the k biggest frequencies. A frequency can't be larger than n. How could you use that limit to avoid a full sort?"},
			{Level: 2, Type: "insight", Content: "Count with a dictionary (Chapter 2), then do a bucket sort: make a list of n + 1 buckets where bucket f holds the values that appear exactly f times. Walk the buckets from f = n down to 1 and collect values until you have k. Every frequency is an integer between 1 and n, so the buckets cover all of them with no comparisons.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Count each item: value -> frequency.\n2. Choose what to remember: buckets[f] = list of values with frequency f, for f from 0 to n.\n3. Record the current item: put every value in the bucket for its frequency.\n4. Return the answer: scan buckets from highest frequency down, appending values, and stop once you have k."},
			{Level: 4, Type: "skeleton", Content: `def topKFrequent(nums, k):
    counts = {}
    for num in nums:
        counts[num] = counts.get(num, 0) + 1
    buckets = [[] for _ in range(____)]
    for num, freq in counts.items():
        ____
    result = []
    for freq in range(____, 0, -1):
        for num in buckets[freq]:
            result.append(num)
            if len(result) == k:
                return result
    return result`},
		},
		Solution: Solution{
			Code: `def topKFrequent(nums, k):
    counts = {}
    for num in nums:
        counts[num] = counts.get(num, 0) + 1
    # buckets[f] holds every value that appears exactly f times
    buckets = [[] for _ in range(len(nums) + 1)]
    for num, freq in counts.items():
        buckets[freq].append(num)
    result = []
    for freq in range(len(nums), 0, -1):
        for num in buckets[freq]:
            result.append(num)
            if len(result) == k:
                return result
    return result`,
			Explanation: "Count each value's frequency with a dictionary, then index values by frequency in an array of buckets and read the buckets from the highest frequency down. Builds on Valid Anagram's counting and adds one idea: when the keys you'd sort by are small integers (frequencies 1..n), bucket them instead of sorting.",
			BruteForce:  "Count frequencies, then sort the distinct values by frequency and take the first k: O(n log n). The waste: a full sort orders every value relative to every other, but we only need the top k, and the sort ignores that frequencies are whole numbers from 1 to n.",
			KeyInsight:  "Frequencies are bounded integers (1..n). An array indexed by frequency puts each value straight into its place, so \"sorting\" by frequency costs O(n) with no comparisons.",
			WhyItWorks:  "After counting, counts[v] is exactly v's frequency, which lies between 1 and n, so every value lands in some bucket 1..n. Reading buckets from n down to 1 visits values in non-increasing order of frequency. The first k values visited are therefore the k most frequent; the uniqueness guarantee means ties at the cut-off can't happen. Building buckets touches each distinct value once and the scan touches each bucket and value at most once: O(n).",
			CommonMistakes: []string{
				"Making only len(nums) buckets. A value that fills the whole list has frequency n, so you need indices 0..n: len(nums) + 1 buckets.",
				"Scanning buckets from low to high frequency and returning the least frequent values. Walk from len(nums) down to 1.",
				"Returning after collecting a whole bucket rather than exactly k values; stop the moment len(result) == k.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Count each item", Explanation: "One pass builds value -> frequency, as in Valid Anagram.", CodeSnippet: "for num in nums:\n    counts[num] = counts.get(num, 0) + 1", LineStart: 3, LineEnd: 4},
				{Title: "Choose what to remember", Explanation: "Frequencies run from 1 to n, so a list of n + 1 buckets indexed by frequency replaces sorting.", CodeSnippet: "buckets = [[] for _ in range(len(nums) + 1)]", LineStart: 6, LineEnd: 6},
				{Title: "Record the current item", Explanation: "Drop each distinct value into the bucket for its frequency.", CodeSnippet: "for num, freq in counts.items():\n    buckets[freq].append(num)", LineStart: 7, LineEnd: 8},
				{Title: "Return the answer", Explanation: "Read from the highest frequency down and stop at exactly k values.", CodeSnippet: "for freq in range(len(nums), 0, -1):\n    for num in buckets[freq]:\n        result.append(num)\n        if len(result) == k:\n            return result", LineStart: 10, LineEnd: 14},
			},
			Alternatives: []AltSolution{
				{
					Name: "Count, then sort by frequency",
					Code: `def topKFrequent(nums, k):
    counts = {}
    for num in nums:
        counts[num] = counts.get(num, 0) + 1
    by_frequency = sorted(counts, key=lambda value: counts[value], reverse=True)
    return by_frequency[:k]`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The obvious approach: sort distinct values by count. Easy to write; the sort is the bottleneck.",
				},
				{
					Name: "Heap of size k",
					Code: `import heapq
from collections import Counter

def topKFrequent(nums, k):
    counts = Counter(nums)
    return heapq.nlargest(k, counts.keys(), key=counts.get)`,
					TimeComplexity:  "O(n log k)",
					SpaceComplexity: "O(n)",
					Explanation:     "A heap keeps only the k best candidates, so each value costs O(log k). Useful when k is much smaller than n or the data streams in. Heaps are taught in Chapter 11.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is bucket sort possible here but not for sorting arbitrary numbers?", Answer: "Bucket sort needs the sort keys to be small integers in a known range so they can be array indices. Frequencies are always integers between 1 and n. Arbitrary numbers, like values up to 10^9, would need a huge array."},
			{Prompt: "What if ties were allowed and you had to return values with equal frequency in increasing order of value?", Answer: "Sort each bucket's values before reading it (or read buckets with sorted(...)). The overall approach stays the same; the extra sorting only touches values that share a frequency."},
			{Prompt: "When would the heap version beat the bucket version in practice?", Answer: "When the data arrives as a stream or doesn't fit in memory at once: a heap of size k can keep the current top k as counts arrive, while bucket sort needs all counts and an array sized by n."},
		},
	},
	{
		ID:              "valid-sudoku",
		Number:          36,
		Order:           6,
		Title:           "Valid Sudoku",
		Slug:            "valid-sudoku",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "Matrix"},
		RelatedChapters: []int{2},
		Description: `A Sudoku board is a 9 x 9 grid given as a list of 9 rows, each a list of 9 one-character strings. A cell holds a digit "1" to "9" or "." when it is empty. The grid is split into nine 3 x 3 boxes: rows 0-2, 3-5, 6-8 crossed with columns 0-2, 3-5, 6-8.

Return true if the filled cells break no rule, and false otherwise. The rules: no digit may appear twice in the same row, twice in the same column, or twice in the same 3 x 3 box.

You only check the cells that are filled. You do not need to decide whether the puzzle can actually be solved, and empty cells never break a rule.`,
		Constraints: []string{
			"board.length == 9",
			"board[i].length == 9",
			"board[i][j] is a digit \"1\"-\"9\" or \".\"",
		},
		Examples: []Example{
			{Input: `board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]`, Output: "true", Explanation: "No row, column or 3 x 3 box contains the same digit twice."},
			{Input: `board = [["8","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]`, Output: "false", Explanation: "The same board with 8 instead of 5 in the top-left cell. Column 0 now contains 8 twice (rows 0 and 3), and so does the top-left box (rows 0 and 2)."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"board": [][]string{{"5", "3", ".", ".", "7", ".", ".", ".", "."}, {"6", ".", ".", "1", "9", "5", ".", ".", "."}, {".", "9", "8", ".", ".", ".", ".", "6", "."}, {"8", ".", ".", ".", "6", ".", ".", ".", "3"}, {"4", ".", ".", "8", ".", "3", ".", ".", "1"}, {"7", ".", ".", ".", "2", ".", ".", ".", "6"}, {".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"}}}, Expected: true, Note: "partly filled valid board"},
			{Input: map[string]interface{}{"board": [][]string{{"8", "3", ".", ".", "7", ".", ".", ".", "."}, {"6", ".", ".", "1", "9", "5", ".", ".", "."}, {".", "9", "8", ".", ".", ".", ".", "6", "."}, {"8", ".", ".", ".", "6", ".", ".", ".", "3"}, {"4", ".", ".", "8", ".", "3", ".", ".", "1"}, {"7", ".", ".", ".", "2", ".", ".", ".", "6"}, {".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"}}}, Expected: false, Note: "two 8s in column 0 (and in the top-left box)"},
			{Input: map[string]interface{}{"board": [][]string{{".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}}}, Expected: true, Hidden: true, Note: "empty board: dots are not digits"},
			{Input: map[string]interface{}{"board": [][]string{{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, {"4", "5", "6", "7", "8", "9", "1", "2", "3"}, {"7", "8", "9", "1", "2", "3", "4", "5", "6"}, {"2", "3", "4", "5", "6", "7", "8", "9", "1"}, {"5", "6", "7", "8", "9", "1", "2", "3", "4"}, {"8", "9", "1", "2", "3", "4", "5", "6", "7"}, {"3", "4", "5", "6", "7", "8", "9", "1", "2"}, {"6", "7", "8", "9", "1", "2", "3", "4", "5"}, {"9", "1", "2", "3", "4", "5", "6", "7", "8"}}}, Expected: true, Hidden: true, Note: "completely filled valid board"},
			{Input: map[string]interface{}{"board": [][]string{{"5", ".", ".", ".", ".", ".", ".", ".", "5"}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}}}, Expected: false, Hidden: true, Note: "repeat only in a row"},
			{Input: map[string]interface{}{"board": [][]string{{".", ".", ".", ".", "7", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", "7", ".", ".", ".", "."}}}, Expected: false, Hidden: true, Note: "repeat only in a column"},
			{Input: map[string]interface{}{"board": [][]string{{"9", ".", ".", ".", ".", ".", ".", ".", "."}, {".", "9", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}}}, Expected: false, Hidden: true, Note: "repeat only inside a 3 x 3 box"},
			{Input: map[string]interface{}{"board": [][]string{{".", ".", ".", "4", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {"4", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}}}, Expected: true, Hidden: true, Note: "valid: cells in boxes 1 and 3 (box = r//3 + c//3 would wrongly merge them)"},
		},
		TimeComplexity:  "O(1) (the board is always 9 x 9)",
		SpaceComplexity: "O(1) (the board is always 9 x 9)",
		StarterCode:     "def isValidSudoku(board):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Contains Duplicate answered \"has this value appeared before?\" for one list. Here every cell belongs to three lists at once: its row, its column and its box. What would you keep for each of them?"},
			{Level: 2, Type: "insight", Content: "Keep one seen-set per row, per column and per box: 27 sets in all. For each filled cell, check its digit against its three sets, then add it. The only puzzle is naming the box: box = (r // 3) * 3 + c // 3 numbers the boxes 0..8 left to right, top to bottom. It's the seen-so-far pattern (Chapter 2), run on 27 lists at the same time.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Choose what to remember: 9 row sets, 9 column sets, 9 box sets.\n2. Skip empty cells.\n3. Look up before you store: if the digit is already in its row set, column set or box set, return False.\n4. Record the current item: add the digit to all three sets.\n5. Return the answer: True if no cell clashed."},
			{Level: 4, Type: "skeleton", Content: `def isValidSudoku(board):
    rows = [set() for _ in range(9)]
    cols = [set() for _ in range(9)]
    boxes = [set() for _ in range(9)]
    for r in range(9):
        for c in range(9):
            digit = board[r][c]
            if digit == '.':
                continue
            box = ____
            if ____:
                return False
            rows[r].add(digit)
            ____
            ____
    return True`},
		},
		Solution: Solution{
			Code: `def isValidSudoku(board):
    rows = [set() for _ in range(9)]
    cols = [set() for _ in range(9)]
    boxes = [set() for _ in range(9)]
    for r in range(9):
        for c in range(9):
            digit = board[r][c]
            if digit == '.':
                continue
            box = (r // 3) * 3 + c // 3
            if digit in rows[r] or digit in cols[c] or digit in boxes[box]:
                return False
            rows[r].add(digit)
            cols[c].add(digit)
            boxes[box].add(digit)
    return True`,
			Explanation: "One pass over the grid with a seen-set for every row, column and box. Builds directly on Contains Duplicate: each of the 27 units is a list in which no value may repeat, and every cell is checked against the three units it belongs to.",
			BruteForce:  "For each filled cell, scan its whole row, whole column and whole box looking for the same digit: about 81 cells x 27 other cells. For an n x n board that is O(n^3) work. The waste: the same row is rescanned once for every cell in it, although what we learn (which digits the row holds) could be remembered.",
			KeyInsight:  "Remember, for each row, column and box, the set of digits placed so far. Then each cell needs three O(1) set lookups instead of three scans.",
			WhyItWorks:  "Invariant: before processing cell (r, c) in row-major order, rows[r], cols[c] and boxes[box] contain exactly the digits of the cells already processed in that row, column and box. A rule is broken iff some later cell repeats a digit already in one of its units, which is exactly what the lookup detects at the moment the second copy is reached. The box formula is correct because r // 3 and c // 3 give the box's row and column (0..2), and band * 3 + stack numbers the nine boxes 0..8 without collisions.",
			CommonMistakes: []string{
				"Computing the box as r // 3 + c // 3. Different boxes then share a number (box row 0, column 1 and box row 1, column 0 both give 1). Use (r // 3) * 3 + c // 3.",
				"Treating \".\" as a value, so the second empty cell in a row looks like a repeat. Skip empty cells before checking.",
				"Trying to check whether the puzzle is solvable. The task only asks whether the filled cells already break a rule.",
			},
			TimeComplexity:  "O(1) (the board is always 9 x 9)",
			SpaceComplexity: "O(1) (the board is always 9 x 9)",
			Walkthrough: []WalkthroughStep{
				{Title: "Choose what to remember", Explanation: "One set per unit: rows[r], cols[c] and boxes[b] each hold the digits already placed there.", CodeSnippet: "rows = [set() for _ in range(9)]\ncols = [set() for _ in range(9)]\nboxes = [set() for _ in range(9)]", LineStart: 2, LineEnd: 4},
				{Title: "Skip empty cells", Explanation: "Dots aren't digits, so they never break a rule and must not be stored.", CodeSnippet: "if digit == '.':\n    continue", LineStart: 8, LineEnd: 9},
				{Title: "Look up before you store", Explanation: "Name the box, then check the digit against all three units the cell belongs to.", CodeSnippet: "box = (r // 3) * 3 + c // 3\nif digit in rows[r] or digit in cols[c] or digit in boxes[box]:\n    return False", LineStart: 10, LineEnd: 12},
				{Title: "Record the current item", Explanation: "The digit is now present in its row, column and box for every later cell.", CodeSnippet: "rows[r].add(digit)\ncols[c].add(digit)\nboxes[box].add(digit)", LineStart: 13, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{
					Name: "Check each of the 27 units separately",
					Code: `def isValidSudoku(board):
    units = []
    for i in range(9):
        units.append([board[i][c] for c in range(9)])
        units.append([board[r][i] for r in range(9)])
    for top in range(0, 9, 3):
        for left in range(0, 9, 3):
            units.append([board[r][c] for r in range(top, top + 3) for c in range(left, left + 3)])
    for unit in units:
        digits = [cell for cell in unit if cell != '.']
        if len(digits) != len(set(digits)):
            return False
    return True`,
					TimeComplexity:  "O(1) (the board is always 9 x 9)",
					SpaceComplexity: "O(1) (the board is always 9 x 9)",
					Explanation:     "Build each row, column and box as a list and apply the Contains Duplicate idea: a list has a repeat when turning it into a set shrinks it. Reads each cell three times, which is fine for 81 cells.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is (r // 3) * 3 + c // 3 a correct box number, and what goes wrong with r // 3 + c // 3?", Answer: "r // 3 is the box's band (0..2) and c // 3 its stack (0..2). Multiplying the band by 3 gives each band its own range of three numbers, like digits in base 3, so the nine boxes get 0..8 with no repeats. Adding without multiplying maps (0,1) and (1,0) both to 1, merging two different boxes."},
			{Prompt: "What if the board were n x n with sqrt(n) x sqrt(n) boxes? What would the complexity be?", Answer: "There are n^2 cells and each needs three O(1) set operations, so O(n^2) time and O(n^2) space for the sets. For n = 9 that is a constant, which is why we call the 9 x 9 version O(1)."},
			{Prompt: "How could you use a single set instead of 27?", Answer: "Store tagged entries such as ('row', r, digit), ('col', c, digit) and ('box', box, digit) in one set. A clash is any tag already present. It's the same idea: the tag plays the role of choosing which of the 27 sets to use."},
		},
	},
	{
		ID:              "encode-and-decode-strings",
		Number:          271,
		Order:           7,
		Title:           "Encode and Decode Strings",
		Slug:            "encode-and-decode-strings",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "String", "Design"},
		RelatedChapters: []int{2},
		Description: `Write two functions. encode(strs) takes a list of strings and returns a single string. decode(s) takes a string produced by encode and returns the original list.

For every possible list, decode(encode(strs)) must equal strs exactly: same number of strings, same order, same characters. The strings may contain any of the 256 ASCII characters, including spaces, digits, punctuation, newlines and whatever delimiter you might be tempted to use. Empty strings and an empty list must survive the round trip too.

Do not use built-in serialisers such as eval or repr tricks; design the format yourself.`,
		Constraints: []string{
			"0 <= strs.length < 100",
			"0 <= strs[i].length < 200",
			"strs[i] contains any of the 256 ASCII characters",
		},
		Examples: []Example{
			{Input: `strs = ["neet","code","love","you"]`, Output: `["neet","code","love","you"]`, Explanation: "With a length prefix the encoded string is \"4#neet4#code4#love3#you\", and decoding reads each length then exactly that many characters."},
			{Input: `strs = ["we","say",":","yes"]`, Output: `["we","say",":","yes"]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"strs": []string{"neet", "code", "love", "you"}}, Expected: []string{"neet", "code", "love", "you"}, Note: "ordinary words"},
			{Input: map[string]interface{}{"strs": []string{"we", "say", ":", "yes"}}, Expected: []string{"we", "say", ":", "yes"}, Note: "a string that is only punctuation"},
			{Input: map[string]interface{}{"strs": []string{}}, Expected: []string{}, Hidden: true, Note: "empty list must decode to an empty list, not [\"\"]"},
			{Input: map[string]interface{}{"strs": []string{""}}, Expected: []string{""}, Hidden: true, Note: "one empty string must survive"},
			{Input: map[string]interface{}{"strs": []string{"", "", ""}}, Expected: []string{"", "", ""}, Hidden: true, Note: "several empty strings"},
			{Input: map[string]interface{}{"strs": []string{"a,b", "c;d", "e|f", "g~h"}}, Expected: []string{"a,b", "c;d", "e|f", "g~h"}, Hidden: true, Note: "common delimiter characters inside the strings"},
			{Input: map[string]interface{}{"strs": []string{"4#ab", "#", "12#"}}, Expected: []string{"4#ab", "#", "12#"}, Hidden: true, Note: "strings that look like length prefixes"},
			{Input: map[string]interface{}{"strs": []string{"abcdefghijklmnop", "x"}}, Expected: []string{"abcdefghijklmnop", "x"}, Hidden: true, Note: "length of 16 has two digits"},
			{Input: map[string]interface{}{"strs": []string{"line\nbreak", "tab\there", "  spaces  "}}, Expected: []string{"line\nbreak", "tab\there", "  spaces  "}, Hidden: true, Note: "whitespace and control characters"},
			{Input: map[string]interface{}{"strs": []string{"//", "/:", "a/b"}}, Expected: []string{"//", "/:", "a/b"}, Hidden: true, Note: "slashes and escape-like sequences"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def encode(strs):\n    # Write your solution here\n    pass\n\n\ndef decode(s):\n    # Write your solution here\n    pass",
		Harness:         Harness{Mode: "roundtrip", Function: "encode,decode"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose you join the strings with a comma. What happens when one of the strings contains a comma? Is there any character you could pick that can never appear inside a string?"},
			{Level: 2, Type: "insight", Content: "No delimiter is safe on its own, because any character may appear in the data. Instead, write each string's length before it: \"5#hello\". The decoder reads digits up to the first '#', and then takes exactly that many characters without looking at them, so a '#' or digit inside the string can't confuse it. Like an array, the format records sizes so you can jump straight to the next item (Chapter 2).", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "Encode:\n1. For each string, write its length, a '#', then the string itself.\n2. Join everything into one string.\nDecode:\n1. Read the length: scan from position i to the next '#' and convert those digits to an int.\n2. Take the string: copy exactly that many characters after the '#'.\n3. Jump to the next item: set i just past those characters and repeat until the end."},
			{Level: 4, Type: "skeleton", Content: `def encode(strs):
    parts = []
    for s in strs:
        parts.append(____)
    return ''.join(parts)


def decode(s):
    result = []
    i = 0
    while i < len(s):
        j = i
        while s[j] != '#':
            j += 1
        length = ____
        start = j + 1
        result.append(____)
        i = ____
    return result`},
		},
		Solution: Solution{
			Code: `def encode(strs):
    parts = []
    for s in strs:
        parts.append(str(len(s)) + '#' + s)
    return ''.join(parts)


def decode(s):
    result = []
    i = 0
    while i < len(s):
        j = i
        while s[j] != '#':
            j += 1
        length = int(s[i:j])
        start = j + 1
        result.append(s[start:start + length])
        i = start + length
    return result`,
			Explanation: "Prefix every string with its length and a '#'. The decoder reads the length, then copies exactly that many characters, so the content of a string is never parsed. Builds on the ladder's array thinking: knowing an item's size lets you jump straight to the next one, just as an array index jumps straight to a slot.",
			BruteForce:  "Join with a delimiter such as ',' and split on it. It is O(n) but wrong: [\"a,b\"] decodes as [\"a\", \"b\"], and [] and [\"\"] both encode to \"\". Picking a rarer delimiter only moves the problem, because the strings may contain any character. The flaw isn't speed; it's that the decoder must guess where a string ends by looking at its content.",
			KeyInsight:  "Store the boundary as data instead of as a special character: a length tells the decoder where the string ends without reading the string, so the string can contain anything.",
			WhyItWorks:  "Invariant: at the top of each decode iteration, i points at the first digit of a length prefix. The digits run up to the first '#' after i, and that '#' must be the separator because a length is written only with digits. We then take exactly length characters, whatever they are, which is precisely the original string, and the next prefix starts right after them. By induction every string is recovered in order. An empty list encodes to \"\" (the loop never runs), and an empty string encodes to \"0#\", so the two are distinguished.",
			CommonMistakes: []string{
				"Joining with a delimiter and splitting on it. Any string containing the delimiter splits into pieces; the length prefix avoids looking at the content.",
				"Assuming the length is one digit (reading s[i] only). Strings of 10 or more characters have multi-digit lengths; read all digits up to the '#'.",
				"Decoding with s.split('#'): a '#' inside a string is then taken as a separator. Find only the first '#' after the current position, then skip over the content by its length without reading it.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Write the length prefix", Explanation: "Each string becomes length + '#' + string, so [\"4#ab\"] encodes to \"4#4#ab\" and still decodes correctly.", CodeSnippet: "parts.append(str(len(s)) + '#' + s)", LineStart: 4, LineEnd: 4},
				{Title: "Read the length", Explanation: "Digits run from i to the first '#'. That '#' is always the separator, because the prefix contains only digits.", CodeSnippet: "j = i\nwhile s[j] != '#':\n    j += 1\nlength = int(s[i:j])", LineStart: 12, LineEnd: 15},
				{Title: "Take the string", Explanation: "Copy exactly length characters after the '#'. Their content is never inspected.", CodeSnippet: "start = j + 1\nresult.append(s[start:start + length])", LineStart: 16, LineEnd: 17},
				{Title: "Jump to the next item", Explanation: "The next prefix begins right after the string we just copied.", CodeSnippet: "i = start + length", LineStart: 18, LineEnd: 18},
			},
			Alternatives: []AltSolution{
				{
					Name: "Escape the separator",
					Code: `def encode(strs):
    out = []
    for s in strs:
        out.append(s.replace('/', '//') + '/:')
    return ''.join(out)


def decode(s):
    result, current, i = [], [], 0
    while i < len(s):
        if s[i] == '/':
            if s[i + 1] == '/':
                current.append('/')
            else:
                result.append(''.join(current))
                current = []
            i += 2
        else:
            current.append(s[i])
            i += 1
    return result`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Another safe format: double every '/' inside a string and end each string with '/:'. The decoder reads '/' pairs: '//' is a real slash, '/:' ends a string. This is how CSV quoting and C string escapes work.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can the decoder trust that the first '#' after position i is the separator, even though strings may contain '#'?", Answer: "Position i is always the start of a length prefix, and a prefix contains only digits. So between i and the separator there can be no '#'. Any '#' inside the string comes after the separator, and the decoder skips over the string by its length without looking."},
			{Prompt: "Why do [] and [\"\"] need different encodings, and are they different here?", Answer: "decode must return exactly what was encoded, so it needs to tell \"no strings\" from \"one empty string\". Here [] encodes to \"\" and [\"\"] encodes to \"0#\", so they differ. A plain join would encode both as \"\"."},
			{Prompt: "What if you had to encode a list of lists of strings? How would you extend the format?", Answer: "Apply the same idea recursively: encode each inner list into one string with the length-prefix format, then encode the list of those strings with length prefixes again. Lengths compose because they never depend on the content."},
		},
	},
	{
		ID:              "product-of-array-except-self",
		Number:          238,
		Order:           8,
		Title:           "Product of Array Except Self",
		Slug:            "product-of-array-except-self",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Prefix Sum"},
		RelatedChapters: []int{2},
		Description: `You are given a list of integers nums. Return a list answer of the same length where answer[i] is the product of every element of nums except nums[i].

You must not use division (so you can't compute the total product and divide by nums[i]), and the solution should run in O(n) time. The product of any prefix or suffix of nums fits in a 32-bit integer.

Follow-up: use only O(1) extra space, not counting the output list.`,
		Constraints: []string{
			"2 <= nums.length <= 10^5",
			"-30 <= nums[i] <= 30",
			"Every prefix and suffix product fits in a 32-bit integer",
		},
		Examples: []Example{
			{Input: "nums = [1,2,3,4]", Output: "[24,12,8,6]", Explanation: "answer[0] = 2*3*4 = 24, answer[1] = 1*3*4 = 12, answer[2] = 1*2*4 = 8, answer[3] = 1*2*3 = 6."},
			{Input: "nums = [-1,1,0,-3,3]", Output: "[0,0,9,0,0]", Explanation: "Every product except the one that leaves out the 0 includes the 0."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}}, Expected: []int{24, 12, 8, 6}, Note: "small positive numbers"},
			{Input: map[string]interface{}{"nums": []int{-1, 1, 0, -3, 3}}, Expected: []int{0, 0, 9, 0, 0}, Note: "one zero: only its own position is nonzero"},
			{Input: map[string]interface{}{"nums": []int{2, 3}}, Expected: []int{3, 2}, Hidden: true, Note: "smallest allowed length"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 5}}, Expected: []int{0, 0, 0}, Hidden: true, Note: "two zeros: every product is 0"},
			{Input: map[string]interface{}{"nums": []int{0, 4, 5}}, Expected: []int{20, 0, 0}, Hidden: true, Note: "zero at the very start"},
			{Input: map[string]interface{}{"nums": []int{-2, -3, -4}}, Expected: []int{12, 8, 6}, Hidden: true, Note: "all negative"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 1, 1}}, Expected: []int{1, 1, 1, 1, 1}, Hidden: true, Note: "all ones"},
			{Input: map[string]interface{}{"nums": []int{5, -1, 2, 3, -2, 1}}, Expected: []int{12, -60, 30, 20, -30, 60}, Hidden: true, Note: "mixed signs, longer array"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1) extra (the output list doesn't count)",
		StarterCode:     "def productExceptSelf(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The product of everything except nums[i] splits into two parts around i. What are they, and could you compute each part for every i without starting over each time?"},
			{Level: 2, Type: "insight", Content: "answer[i] = (product of everything left of i) * (product of everything right of i). A left-to-right pass can keep a running prefix product, and a right-to-left pass a running suffix product. Each running product is updated with one multiplication per step, so nothing is recomputed. This is precomputation over an array, the arrays half of Chapter 2.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Choose what to remember: one running product, starting at 1 (the product of nothing).\n2. Left pass: for i from left to right, store the running prefix in answer[i], then multiply nums[i] into it.\n3. Right pass: for i from right to left, multiply answer[i] by the running suffix, then multiply nums[i] into it.\n4. Return the answer."},
			{Level: 4, Type: "skeleton", Content: `def productExceptSelf(nums):
    n = len(nums)
    answer = [1] * n
    prefix = 1
    for i in range(n):
        answer[i] = ____
        prefix *= ____
    suffix = 1
    for i in range(____):
        answer[i] *= ____
        suffix *= nums[i]
    return answer`},
		},
		Solution: Solution{
			Code: `def productExceptSelf(nums):
    n = len(nums)
    answer = [1] * n
    prefix = 1  # product of nums[0..i-1]
    for i in range(n):
        answer[i] = prefix
        prefix *= nums[i]
    suffix = 1  # product of nums[i+1..n-1]
    for i in range(n - 1, -1, -1):
        answer[i] *= suffix
        suffix *= nums[i]
    return answer`,
			Explanation: "Split each answer into the product of the elements to its left and the product of those to its right. One forward pass writes the left products into the answer, and one backward pass multiplies in the right products. Builds on the ladder's \"remember what you've already computed\" theme, but the memory is a running product over array positions instead of a hash map.",
			BruteForce:  "For each i, multiply all the other n - 1 elements: n * (n - 1) multiplications, about 10^10 for n = 10^5, so O(n^2). The waste: the products for i and i + 1 share almost all their factors, yet each is rebuilt from scratch.",
			KeyInsight:  "Everything except i is \"left of i\" times \"right of i\", and both of those can be grown one element at a time as you sweep across the array.",
			WhyItWorks:  "Invariant of the first loop: at the start of iteration i, prefix = nums[0] * ... * nums[i-1] (1 when i = 0), so answer[i] becomes exactly the left product. Invariant of the second loop: at the start of iteration i, suffix = nums[i+1] * ... * nums[n-1] (1 when i = n - 1), so answer[i] becomes left * right, the product of everything except nums[i]. No division is used, so zeros need no special case.",
			CommonMistakes: []string{
				"Updating prefix before writing answer[i], which includes nums[i] in its own product. Write first, then multiply nums[i] in.",
				"Computing the total product and dividing by nums[i]. It's forbidden, and it breaks when nums[i] is 0 (division by zero) or when there are two zeros.",
				"Starting the running products at 0 instead of 1. The product of no elements is 1, and anything times 0 stays 0.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1) extra (the output list doesn't count)",
			Walkthrough: []WalkthroughStep{
				{Title: "Choose what to remember", Explanation: "The output starts as all 1s; a single running product, starting at 1, sweeps across.", CodeSnippet: "answer = [1] * n\nprefix = 1  # product of nums[0..i-1]", LineStart: 3, LineEnd: 4},
				{Title: "Sweep left to right", Explanation: "Write the product of everything left of i, then fold nums[i] in for the next position. For [1,2,3,4] answer becomes [1,1,2,6].", CodeSnippet: "answer[i] = prefix\nprefix *= nums[i]", LineStart: 6, LineEnd: 7},
				{Title: "Sweep right to left", Explanation: "Multiply in the product of everything right of i. The suffixes are 24,12,4,1, giving [24,12,8,6].", CodeSnippet: "answer[i] *= suffix\nsuffix *= nums[i]", LineStart: 10, LineEnd: 11},
				{Title: "Return the answer", Explanation: "Each position now holds left product times right product.", CodeSnippet: "return answer", LineStart: 12, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: multiply the others",
					Code: `def productExceptSelf(nums):
    answer = []
    for i in range(len(nums)):
        product = 1
        for j in range(len(nums)):
            if j != i:
                product *= nums[j]
        answer.append(product)
    return answer`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1) extra",
					Explanation:     "Recomputes each product from scratch; correct but too slow for n = 10^5.",
				},
				{
					Name: "Separate prefix and suffix arrays",
					Code: `def productExceptSelf(nums):
    n = len(nums)
    left = [1] * n
    right = [1] * n
    for i in range(1, n):
        left[i] = left[i - 1] * nums[i - 1]
    for i in range(n - 2, -1, -1):
        right[i] = right[i + 1] * nums[i + 1]
    return [left[i] * right[i] for i in range(n)]`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The same idea with both tables stored explicitly, which is easier to trace; the main solution folds them into the output list to save memory.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is prefix updated after answer[i] is written, not before?", Answer: "answer[i] must exclude nums[i]. Writing first stores the product of nums[0..i-1]; multiplying nums[i] in afterwards prepares the value for position i + 1. Swapping the two lines would include nums[i] in its own answer."},
			{Prompt: "With division allowed, how would you handle zeros correctly?", Answer: "Count the zeros. With two or more, every answer is 0. With exactly one, only the zero's position gets the product of the other numbers; all others are 0. With none, answer[i] = total // nums[i]. The prefix/suffix method needs none of these cases."},
			{Prompt: "What if you needed sums instead of products: the sum of everything except nums[i]?", Answer: "Then subtraction is safe: answer[i] = total - nums[i]. The prefix/suffix idea still works (prefix sum + suffix sum), and it is the basis of prefix sums, which answer range-sum questions in O(1)."},
		},
	},
	{
		ID:              "longest-consecutive-sequence",
		Number:          128,
		Order:           9,
		Title:           "Longest Consecutive Sequence",
		Slug:            "longest-consecutive-sequence",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "Union Find"},
		RelatedChapters: []int{2},
		Description: `You are given an unsorted list of integers nums. A consecutive run is a set of values x, x+1, x+2, ..., x+m-1 that are all present in nums (their positions in the list don't matter). Return the length m of the longest such run. Duplicate values count once. Return 0 for an empty list.

Your algorithm must run in O(n) time, so sorting is not allowed.`,
		Constraints: []string{
			"0 <= nums.length <= 10^5",
			"-10^9 <= nums[i] <= 10^9",
		},
		Examples: []Example{
			{Input: "nums = [100,4,200,1,3,2]", Output: "4", Explanation: "The values 1, 2, 3 and 4 are all present, so the longest run has length 4."},
			{Input: "nums = [0,3,7,2,5,8,4,6,0,1]", Output: "9", Explanation: "0 through 8 are all present; the second 0 doesn't lengthen the run."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{100, 4, 200, 1, 3, 2}}, Expected: 4, Note: "run 1,2,3,4 hidden among other values"},
			{Input: map[string]interface{}{"nums": []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}}, Expected: 9, Note: "long run with a duplicate 0"},
			{Input: map[string]interface{}{"nums": []int{}}, Expected: 0, Hidden: true, Note: "empty array"},
			{Input: map[string]interface{}{"nums": []int{5}}, Expected: 1, Hidden: true, Note: "single element"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 0, 1}}, Expected: 3, Hidden: true, Note: "duplicates must not lengthen the run"},
			{Input: map[string]interface{}{"nums": []int{-3, -1, -2, 10}}, Expected: 3, Hidden: true, Note: "negative run"},
			{Input: map[string]interface{}{"nums": []int{10, 30, 20}}, Expected: 1, Hidden: true, Note: "no two consecutive values"},
			{Input: map[string]interface{}{"nums": []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}}, Expected: 7, Hidden: true, Note: "two runs; the longer one wins"},
			{Input: map[string]interface{}{"nums": []int{7, 7, 7}}, Expected: 1, Hidden: true, Note: "all equal values form a run of 1"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def longestConsecutive(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you could ask \"is x + 1 present?\" in O(1), you could walk a run from any value. But walking from every value repeats a lot of work. Which values are the only ones worth starting a walk from?"},
			{Level: 2, Type: "insight", Content: "Put all values in a set (Chapter 2). A value x starts a run exactly when x - 1 is not in the set. Only from those starts, count upward while x + 1, x + 2, ... are present. Each value is then stepped over by at most one walk, so the total work is O(n).", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Choose what to remember: a set of all values (this also removes duplicates).\n2. Find the starts: skip any value whose predecessor x - 1 is in the set.\n3. Walk the run: from a start, step end = end + 1 while it is in the set.\n4. Record the answer: keep the longest end - start + 1."},
			{Level: 4, Type: "skeleton", Content: `def longestConsecutive(nums):
    values = set(nums)
    longest = 0
    for start in values:
        if ____:
            continue
        end = start
        while ____:
            end += 1
        longest = max(longest, ____)
    return longest`},
		},
		Solution: Solution{
			Code: `def longestConsecutive(nums):
    values = set(nums)
    longest = 0
    for start in values:
        if start - 1 in values:
            continue  # not the first number of its run
        end = start
        while end + 1 in values:
            end += 1
        longest = max(longest, end - start + 1)
    return longest`,
			Explanation: "Store the values in a set, then grow a run only from values that have no predecessor. Builds on Contains Duplicate's set (O(1) membership) and adds one idea: choose where to start so that no work is repeated, which is what turns an O(n^2) walk into O(n).",
			BruteForce:  "Sort and scan for the longest stretch where each value is one more than the previous (skipping duplicates): O(n log n), which the problem forbids. A set-based brute force walks upward from every value: for the run 1..n that is n + (n-1) + ... + 1 = O(n^2) steps. The waste: a walk from 2 retraces exactly the steps of the walk from 1.",
			KeyInsight:  "Every run has exactly one smallest value, the only one whose predecessor is missing. Starting walks only there means every value is walked over once.",
			WhyItWorks:  "Correctness: every run x..x+m-1 has a start x with x - 1 absent, so the loop walks it fully and records m; no longer run is missed, and walks never jump gaps, so no run is overcounted. Duplicates vanish in the set. Time: the outer loop does one O(1) check per distinct value, and each while-step moves through a value that belongs to the run of the current start; since each value belongs to exactly one run and each run is walked once, all while-steps together are at most n. So the total is O(n).",
			CommonMistakes: []string{
				"Walking from every value, not only from run starts. It gives the right answer but is O(n^2) on input like [1, 2, ..., 100000] and times out.",
				"Looping over nums instead of the set. With many duplicates of a run start, such as [1, 1, 1, ..., 2, 3, ...], the same run is walked once per copy. Loop over the set.",
				"Returning 1 for an empty list because longest starts at 1. Start at 0 so [] returns 0.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Choose what to remember", Explanation: "A set gives O(1) \"is v present?\" and drops duplicates.", CodeSnippet: "values = set(nums)", LineStart: 2, LineEnd: 2},
				{Title: "Find the starts", Explanation: "If start - 1 exists, start sits in the middle of a run that will be walked from its real beginning, so skip it.", CodeSnippet: "if start - 1 in values:\n    continue  # not the first number of its run", LineStart: 5, LineEnd: 6},
				{Title: "Walk the run", Explanation: "Step upward while the next value exists. For [100,4,200,1,3,2] the walk from 1 visits 2, 3, 4 and stops at 5.", CodeSnippet: "end = start\nwhile end + 1 in values:\n    end += 1", LineStart: 7, LineEnd: 9},
				{Title: "Record the answer", Explanation: "The run start..end has end - start + 1 values; keep the longest.", CodeSnippet: "longest = max(longest, end - start + 1)", LineStart: 10, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{
					Name: "Sort, then scan",
					Code: `def longestConsecutive(nums):
    if not nums:
        return 0
    ordered = sorted(set(nums))
    longest = current = 1
    for i in range(1, len(ordered)):
        if ordered[i] == ordered[i - 1] + 1:
            current += 1
        else:
            current = 1
        longest = max(longest, current)
    return longest`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "After sorting the distinct values, a run is a stretch where each value is one more than the previous. Simple, but the sort breaks the O(n) requirement.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "The code has a while loop inside a for loop. Why is it O(n) and not O(n^2)?", Answer: "The while loop only runs from run starts, and it steps through the values of that run. Each value belongs to exactly one run, and each run is walked exactly once, so all while-steps added together are at most n. Nested loops multiply only when the inner loop can do a lot of work for many outer iterations."},
			{Prompt: "Why loop over the set rather than over nums?", Answer: "nums can contain many copies of a run start. Each copy passes the \"start - 1 not present\" test, so looping over nums would walk the same run again for every copy, which can be O(n^2). The set holds each value once."},
			{Prompt: "What if you also had to return the run itself, not just its length?", Answer: "Remember the best start along with the best length; at the end return list(range(best_start, best_start + longest)). The walk already knows start and end, so no extra passes are needed."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, ArraysHashingProblems...)
}
