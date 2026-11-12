package practice

// ArraysHashingProblems is the arrays-hashing ladder (chapter 2). Order runs
// from the plainest use of a set to problems that combine several ideas:
// set -> counting -> turning counts into pairs -> turning counts into
// triplets -> seen-so-far map -> a map kept between calls -> grouping by a
// key -> bucketing by count -> several sets at once -> a self-describing
// format -> a map and an array kept in sync -> prefix and suffix products ->
// "only start at the start of a run" -> the array as its own hash table.
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
		ID:              "longest-palindrome",
		Number:          409,
		Order:           3,
		Title:           "Longest Palindrome",
		Slug:            "longest-palindrome",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Hash Table", "String", "Greedy", "Google", "Google L3", "Google OA"},
		RelatedChapters: []int{2, 17},
		Description: `A palindrome is a string that reads the same forwards and backwards, such as "racecar" or "abba".

You are given a string s made of English letters. You may pick any of its letters (each position at most once) and arrange the picked letters in any order you like. Return the length of the longest palindrome you can build this way.

Letters are case-sensitive: "A" and "a" are different letters, so "Aa" is not a palindrome. You don't have to return the palindrome itself, only its length.`,
		Constraints: []string{
			"1 <= s.length <= 2000",
			"s consists of lowercase and/or uppercase English letters",
		},
		Examples: []Example{
			{Input: `s = "abccccdd"`, Output: "7", Explanation: "Use the four c's and the two d's as mirrored pairs and put one leftover letter (a or b) in the middle, for example \"dccaccd\"."},
			{Input: `s = "z"`, Output: "1", Explanation: "A single letter is already a palindrome."},
			{Input: `s = "aaabbb"`, Output: "5", Explanation: "Each letter appears 3 times. Two a's and two b's form pairs, and one of the leftovers goes in the middle, for example \"ababa\"."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "abccccdd"}, Expected: 7, Note: "pairs plus one leftover letter in the middle"},
			{Input: map[string]interface{}{"s": "z"}, Expected: 1, Note: "single letter"},
			{Input: map[string]interface{}{"s": "aaabbb"}, Expected: 5, Note: "two odd counts: only one odd letter may sit in the middle"},
			{Input: map[string]interface{}{"s": "Aa"}, Expected: 1, Hidden: true, Note: "case-sensitive: A and a don't pair"},
			{Input: map[string]interface{}{"s": "abc"}, Expected: 1, Hidden: true, Note: "no pairs at all, but one letter can still be the centre"},
			{Input: map[string]interface{}{"s": "bbbb"}, Expected: 4, Hidden: true, Note: "all counts even: no centre letter to add"},
			{Input: map[string]interface{}{"s": "ccc"}, Expected: 3, Hidden: true, Note: "one odd count used in full"},
			{Input: map[string]interface{}{"s": "AAaaBBbbx"}, Expected: 9, Hidden: true, Note: "mixed case pairs plus a centre"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def longestPalindrome(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Picture the palindrome you're building. Every letter on the left half has a mirror twin on the right half. Which letters, if any, are allowed to have no twin?"},
			{Level: 2, Type: "insight", Content: "Only how many of each letter you have matters, not where they were, so count the letters with a dictionary (counting, from Chapter 2, as in Valid Anagram). Every letter contributes its pairs: count // 2 * 2 characters. On top of that, exactly one letter with a leftover (an odd count) may sit alone in the centre, adding 1.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Count each item: build letter -> how many times it appears.\n2. Use the pairs: add count // 2 * 2 for every letter.\n3. Remember a leftover: note whether any count is odd.\n4. Return the answer: add 1 for the centre if a leftover exists."},
			{Level: 4, Type: "skeleton", Content: `def longestPalindrome(s):
    counts = {}
    for ch in s:
        counts[ch] = counts.get(ch, 0) + 1
    length = 0
    has_odd = False
    for count in counts.values():
        length += ____
        if ____:
            has_odd = True
    if has_odd:
        ____
    return length`},
		},
		Solution: Solution{
			Code: `def longestPalindrome(s):
    counts = {}
    for ch in s:
        counts[ch] = counts.get(ch, 0) + 1
    length = 0
    has_odd = False
    for count in counts.values():
        length += count // 2 * 2
        if count % 2 == 1:
            has_odd = True
    if has_odd:
        length += 1
    return length`,
			Explanation: "Count every letter, take all of its complete pairs, and add one extra letter for the centre if any letter had a leftover. Builds on Valid Anagram: the same letter -> count dictionary, but instead of comparing two sets of counts we turn one set of counts into an answer (pairs mirror each other, one odd letter sits in the middle).",
			BruteForce:  "Try every subset of the letters and every arrangement of each subset, keep the ones that read the same both ways, and report the longest. With n letters that is far more than 2^n candidates, hopeless even for n = 30. The waste: positions and arrangements are irrelevant; whether a multiset of letters can form a palindrome depends only on how many of each letter it holds.",
			KeyInsight:  "In a palindrome, positions pair up around the centre, so every letter appears an even number of times, except at most one letter that takes the single middle spot. So the answer is the sum of the even parts of all counts, plus 1 if anything is left over.",
			WhyItWorks:  "Upper bound: a palindrome of length L pairs position i with position L-1-i, so each letter is used an even number of times, apart from the middle position when L is odd. A letter with count c can therefore contribute at most c // 2 * 2 paired characters, and at most one extra character overall can be unpaired. Achievability: place the pairs symmetrically from the outside in, then put one leftover letter (if any) in the middle. Both bounds meet, so the formula is exactly the maximum.",
			CommonMistakes: []string{
				"Adding every odd count in full, as if every odd letter could be the centre. For \"aaabbb\" that gives 6, but only one letter may sit unpaired. Add count // 2 * 2 for every letter and a single +1 at the end.",
				"Forgetting the centre letter entirely, so \"abc\" returns 0 instead of 1. If any count is odd, one leftover can go in the middle.",
				"Lower-casing the string first. The problem is case-sensitive, so \"Aa\" can't pair up and the answer is 1, not 2.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Count each item", Explanation: "One pass builds letter -> count. There are at most 52 letters, so the dictionary has constant size.", CodeSnippet: "for ch in s:\n    counts[ch] = counts.get(ch, 0) + 1", LineStart: 3, LineEnd: 4},
				{Title: "Use the pairs", Explanation: "count // 2 * 2 is the largest even number not above count: every full pair can be mirrored. For \"abccccdd\" that is 0 + 0 + 4 + 2 = 6.", CodeSnippet: "for count in counts.values():\n    length += count // 2 * 2", LineStart: 7, LineEnd: 8},
				{Title: "Remember a leftover", Explanation: "An odd count leaves one letter unpaired. We only need to know whether at least one such letter exists, not how many.", CodeSnippet: "if count % 2 == 1:\n    has_odd = True", LineStart: 9, LineEnd: 10},
				{Title: "Return the answer", Explanation: "One leftover can sit in the centre, so add 1 if there is any. For \"abccccdd\" that gives 6 + 1 = 7.", CodeSnippet: "if has_odd:\n    length += 1\nreturn length", LineStart: 11, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Pair letters with a set as you go",
					Code: `def longestPalindrome(s):
    unpaired = set()
    length = 0
    for ch in s:
        if ch in unpaired:
            unpaired.remove(ch)
            length += 2
        else:
            unpaired.add(ch)
    if unpaired:
        return length + 1
    return length`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Contains Duplicate's set, used to match each letter with its twin: a letter waiting in the set finds its partner and the pair adds 2. Whatever is left unpaired at the end can supply one centre letter.",
				},
				{
					Name: "Sort, then pair neighbours",
					Code: `def longestPalindrome(s):
    letters = sorted(s)
    length = 0
    i = 0
    while i + 1 < len(letters):
        if letters[i] == letters[i + 1]:
            length += 2
            i += 2
        else:
            i += 1
    if length < len(letters):
        length += 1
    return length`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "After sorting, equal letters sit together, so pairs are adjacent. If some letter was not used in a pair, one of them can be the centre. Correct but slower than counting because of the sort.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can at most one letter with an odd count contribute its leftover?", Answer: "In a palindrome every position except the exact middle has a mirror partner holding the same letter. Only the middle position (present when the length is odd) can hold a letter without a twin, and there is just one middle, so at most one leftover letter fits."},
			{Prompt: "What if you had to return the palindrome itself, not its length?", Answer: "Build the left half from the counts: for each letter add count // 2 copies. Pick any odd-count letter as the middle (or none). The answer is left + middle + reversed(left). The counting step is unchanged; only the output step grows."},
			{Prompt: "Why is the space O(1) even though we store a dictionary?", Answer: "The keys are English letters, upper and lower case, so the dictionary never holds more than 52 entries however long s is. A bound that doesn't grow with n is O(1)."},
		},
	},
	{
		ID:              "number-of-unequal-triplets-in-array",
		Number:          2475,
		Order:           4,
		Title:           "Number of Unequal Triplets in Array",
		Slug:            "number-of-unequal-triplets-in-array",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "Sorting", "Google", "Google L3", "Google Phone Screen"},
		RelatedChapters: []int{2, 1},
		Description: `You are given a list of positive integers nums. A triplet of positions (i, j, k) with i < j < k is called unequal when the three values nums[i], nums[j] and nums[k] are all different from each other: nums[i] != nums[j], nums[j] != nums[k] and nums[i] != nums[k].

Return how many unequal triplets of positions exist. Triplets are counted by position, so two triplets that use the same values at different positions are counted separately.`,
		Constraints: []string{
			"3 <= nums.length <= 100",
			"1 <= nums[i] <= 1000",
		},
		Examples: []Example{
			{Input: "nums = [4,4,2,4,3]", Output: "3", Explanation: "The only values are 4, 2 and 3. A valid triplet needs the 2, the 3 and one of the three 4s, which gives 3 choices of positions: (0,2,4), (1,2,4) and (2,3,4)."},
			{Input: "nums = [1,1,1,1,1]", Output: "0", Explanation: "Every value is the same, so no triplet has three different values."},
			{Input: "nums = [1,2,3]", Output: "1", Explanation: "The single triplet (0,1,2) uses three different values."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{4, 4, 2, 4, 3}}, Expected: 3, Note: "one repeated value pairs with two singletons"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 1, 1}}, Expected: 0, Note: "all equal values"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}}, Expected: 1, Note: "minimum length, all different"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 1}}, Expected: 0, Hidden: true, Note: "first and last are equal: nums[i] != nums[k] must be checked too"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}}, Expected: 4, Hidden: true, Note: "all distinct: every one of the 4 position triplets counts"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 2, 3}}, Expected: 2, Hidden: true, Note: "count position triplets, not combinations of distinct values"},
			{Input: map[string]interface{}{"nums": []int{5, 1, 5, 1, 5, 2}}, Expected: 6, Hidden: true, Note: "groups of sizes 3, 2 and 1 interleaved: 3 * 2 * 1"},
			{Input: map[string]interface{}{"nums": []int{7, 7, 8}}, Expected: 0, Hidden: true, Note: "only two distinct values can never form a triplet"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def unequalTriplets(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Does the order of the positions really matter here? If you knew only how many times each value appears, could you count the triplets without looking at positions at all?"},
			{Level: 2, Type: "insight", Content: "Any three positions holding three different values form exactly one triplet with i < j < k, so only the counts matter (counting with a dictionary, Chapter 2). Process the value groups one at a time. For the current group of size c, let before be the number of elements in groups already processed and after the number in groups still to come. Choosing one element from before, one from this group and one from after gives before * c * after triplets, and each triplet is counted exactly once: when its middle group is the current one.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Count each item: value -> how many times it appears.\n2. Split around the current group: before = elements in earlier groups, after = n - before - c.\n3. Combine the counts: add before * c * after.\n4. Move the group to the left side: before += c.\n5. Return the answer: the running total."},
			{Level: 4, Type: "skeleton", Content: `def unequalTriplets(nums):
    counts = {}
    for num in nums:
        counts[num] = counts.get(num, 0) + 1
    total = 0
    before = 0
    n = len(nums)
    for count in counts.values():
        after = ____
        total += ____
        ____
    return total`},
		},
		Solution: Solution{
			Code: `def unequalTriplets(nums):
    counts = {}
    for num in nums:
        counts[num] = counts.get(num, 0) + 1
    total = 0
    before = 0  # elements in the groups already processed
    n = len(nums)
    for count in counts.values():
        after = n - before - count
        total += before * count * after
        before += count
    return total`,
			Explanation: "Count how often each value appears, then walk the value groups, multiplying the number of elements before the current group, inside it and after it. Builds on Longest Palindrome: there the letter counts were turned into pairs; here the value counts are turned into triplets by multiplication, so we never look at individual positions.",
			BruteForce:  "Three nested loops over i < j < k, checking that the three values are pairwise different: n(n-1)(n-2)/6 checks, O(n^3). With n = 100 that is about 160,000 checks, which is fine here, but at n = 10^5 it would be over 10^14. The waste: it examines every triplet one by one, even though all the triplets drawn from the same three value groups behave identically.",
			KeyInsight:  "Positions with equal values are interchangeable, so a triplet is determined by choosing three different groups and one element from each. Splitting the groups into those before, the current one and those after counts each choice exactly once.",
			WhyItWorks:  "Fix any order of the value groups (the dictionary's order). A set of three positions with three different values touches exactly three groups, and exactly one of them is the middle one in that order. When the loop is on that middle group, the triplet is one of the before * count * after choices; on every other iteration it is not counted, because then the triplet doesn't have exactly one element on each side. Each set of three positions corresponds to exactly one triple i < j < k, so the sum is exactly the number of unequal triplets.",
			CommonMistakes: []string{
				"Checking only neighbouring pairs, nums[i] != nums[j] and nums[j] != nums[k], and forgetting nums[i] != nums[k]. Then [1,2,1] counts a triplet. All three pairs must differ.",
				"Counting combinations of distinct values instead of positions. [4,4,2,4,3] has only one set of three distinct values but 3 position triplets. Multiply by the group sizes.",
				"Counting each triplet several times, for example adding left * count * right for every possible ordering of the three groups, which gives 6 times the answer. Count each triplet only when its middle group is the current one.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Count each item", Explanation: "One pass builds value -> group size. For [4,4,2,4,3] the groups are 4:3, 2:1, 3:1.", CodeSnippet: "for num in nums:\n    counts[num] = counts.get(num, 0) + 1", LineStart: 3, LineEnd: 4},
				{Title: "Split around the current group", Explanation: "before counts the elements of earlier groups; after is everything that is neither before nor in this group.", CodeSnippet: "for count in counts.values():\n    after = n - before - count", LineStart: 8, LineEnd: 9},
				{Title: "Combine the counts", Explanation: "One element from each side and one from this group: before * count * after triplets. For [4,4,2,4,3]: 0*3*2 = 0, then 3*1*1 = 3, then 4*1*0 = 0, total 3.", CodeSnippet: "total += before * count * after", LineStart: 10, LineEnd: 10},
				{Title: "Move the group to the left side", Explanation: "The current group is now an earlier group for every group still to come.", CodeSnippet: "before += count\nreturn total", LineStart: 11, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: check every triplet",
					Code: `def unequalTriplets(nums):
    n = len(nums)
    total = 0
    for i in range(n):
        for j in range(i + 1, n):
            for k in range(j + 1, n):
                if nums[i] != nums[j] and nums[j] != nums[k] and nums[i] != nums[k]:
                    total += 1
    return total`,
					TimeComplexity:  "O(n^3)",
					SpaceComplexity: "O(1)",
					Explanation:     "Checks every triplet directly. With n <= 100 it is fast enough here, which is why this is an Easy problem, but it is the version an interviewer will ask you to improve.",
				},
				{
					Name: "Sort, then split around each run",
					Code: `def unequalTriplets(nums):
    ordered = sorted(nums)
    n = len(ordered)
    total = 0
    start = 0
    while start < n:
        end = start
        while end < n and ordered[end] == ordered[start]:
            end += 1
        total += start * (end - start) * (n - end)
        start = end
    return total`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Sorting makes each value group a contiguous run, so start is the number of elements before the run and n - end the number after it. Same counting argument, with a sort instead of a dictionary.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why may we ignore the condition i < j < k and just count sets of three positions?", Answer: "Any three different positions can be listed in increasing order in exactly one way, so each set of three positions is exactly one triple with i < j < k. Counting sets is the same as counting ordered triples."},
			{Prompt: "Why does the order in which the dictionary visits the groups not matter?", Answer: "The argument only needs some fixed order of the groups: each triplet is counted when its middle group (in that order) is processed. Any order gives each triplet exactly one middle group, so every order yields the same total."},
			{Prompt: "What if you had to count pairs (i, j) with nums[i] != nums[j] instead?", Answer: "The same trick with one side: for each group of size c, add before * c and then before += c. Equivalently, total pairs n(n-1)/2 minus, for each group, c(c-1)/2 equal pairs."},
		},
	},
	{
		ID:              "two-sum",
		Number:          1,
		Order:           5,
		Title:           "Two Sum",
		Slug:            "two-sum",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "Google", "Google L3", "Google Phone Screen"},
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
		ID:              "logger-rate-limiter",
		Number:          359,
		Order:           6,
		Title:           "Logger Rate Limiter",
		Slug:            "logger-rate-limiter",
		Difficulty:      "Easy",
		Category:        "arrays-hashing",
		Tags:            []string{"Hash Table", "Design", "Data Stream", "Google", "Google L3", "Google Phone Screen"},
		RelatedChapters: []int{2},
		Description: `A logging system receives a stream of messages. Each message is a string and arrives with a timestamp, a whole number of seconds. To avoid flooding the log, the same message may be printed at most once every 10 seconds: if a message was last printed at time t, it may be printed again only at time t + 10 or later. Different messages don't affect each other.

Design a class Logger:

- Logger() creates the logger with nothing printed yet.
- shouldPrintMessage(timestamp, message) returns true if the message should be printed at this timestamp, and false otherwise. When it returns true, the message counts as printed at that timestamp. When it returns false, nothing is recorded: a rejected message does not restart its 10-second wait.

Calls arrive in chronological order: each timestamp is greater than or equal to the previous one. Several messages may share the same timestamp.`,
		Constraints: []string{
			"0 <= timestamp <= 10^9",
			"Timestamps are passed in non-decreasing order",
			"1 <= message.length <= 30",
			"At most 10^4 calls are made to shouldPrintMessage",
		},
		Examples: []Example{
			{Input: `["Logger","shouldPrintMessage","shouldPrintMessage","shouldPrintMessage","shouldPrintMessage","shouldPrintMessage","shouldPrintMessage","shouldPrintMessage"]
[[],[1,"ping"],[2,"pong"],[3,"ping"],[8,"pong"],[10,"ping"],[11,"ping"],[12,"pong"]]`, Output: "[null,true,true,false,false,false,true,true]", Explanation: "\"ping\" prints at 1, so it is blocked until 11: the calls at 3 and 10 return false and the call at 11 returns true. \"pong\" prints at 2, so it is blocked at 8 and allowed again at 12."},
			{Input: `["Logger","shouldPrintMessage","shouldPrintMessage","shouldPrintMessage"]
[[],[5,"a"],[14,"a"],[15,"a"]]`, Output: "[null,true,false,true]", Explanation: "After printing at 5, \"a\" may print again at exactly 15, but not at 14."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"Logger", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage"}, "values": [][]interface{}{{}, {1, "ping"}, {2, "pong"}, {3, "ping"}, {8, "pong"}, {10, "ping"}, {11, "ping"}, {12, "pong"}}}, Expected: []interface{}{nil, true, true, false, false, false, true, true}, Note: "two messages with independent 10-second windows"},
			{Input: map[string]interface{}{"operations": []string{"Logger", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage"}, "values": [][]interface{}{{}, {5, "a"}, {14, "a"}, {15, "a"}}}, Expected: []interface{}{nil, true, false, true}, Note: "exactly 10 seconds later is allowed, 9 is not"},
			{Input: map[string]interface{}{"operations": []string{"Logger", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage"}, "values": [][]interface{}{{}, {1, "x"}, {5, "x"}, {9, "x"}, {11, "x"}}}, Expected: []interface{}{nil, true, false, false, true}, Hidden: true, Note: "rejected calls must not restart the wait"},
			{Input: map[string]interface{}{"operations": []string{"Logger", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage"}, "values": [][]interface{}{{}, {0, "a"}, {0, "b"}, {0, "a"}, {9, "b"}}}, Expected: []interface{}{nil, true, true, false, false}, Hidden: true, Note: "timestamp 0 is a real print time (a stored 0 is not \"never printed\")"},
			{Input: map[string]interface{}{"operations": []string{"Logger", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage"}, "values": [][]interface{}{{}, {100, "m"}, {1000000000, "m"}, {1000000005, "m"}}}, Expected: []interface{}{nil, true, true, false}, Hidden: true, Note: "huge gap, then a repeat inside the window at the maximum timestamp"},
			{Input: map[string]interface{}{"operations": []string{"Logger", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage"}, "values": [][]interface{}{{}, {1, "a"}, {2, "b"}, {3, "c"}, {11, "a"}, {11, "b"}, {12, "b"}, {12, "c"}, {13, "c"}}}, Expected: []interface{}{nil, true, true, true, true, false, true, false, true}, Hidden: true, Note: "three messages whose windows end one second apart"},
			{Input: map[string]interface{}{"operations": []string{"Logger", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage"}, "values": [][]interface{}{{}, {7, "same"}, {7, "same"}, {17, "same"}, {17, "same"}}}, Expected: []interface{}{nil, true, false, true, false}, Hidden: true, Note: "the same message twice in one second is printed once"},
		},
		TimeComplexity:  "O(1)",
		SpaceComplexity: "O(m)",
		StarterCode: `class Logger:
    def __init__(self):
        pass

    def shouldPrintMessage(self, timestamp: int, message: str) -> bool:
        pass`,
		Harness: Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "To decide about a message right now, what is the one fact about that message's past that you need, and how could you look it up instantly by the message text?"},
			{Level: 2, Type: "insight", Content: "Keep a dictionary from message to the earliest time it may print again (last print time + 10). A call is allowed exactly when timestamp >= that time; only then update the entry. This is the seen-so-far map from Chapter 2: Two Sum stored value -> index, here we store message -> next allowed time.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Choose what to remember: message -> earliest timestamp it may print again.\n2. Look up before you store: if the current timestamp is earlier than that time, return False and change nothing.\n3. Record the current item: otherwise set the entry to timestamp + 10 and return True."},
			{Level: 4, Type: "skeleton", Content: `class Logger:
    def __init__(self):
        self.next_allowed = ____

    def shouldPrintMessage(self, timestamp: int, message: str) -> bool:
        if ____:
            return False
        ____
        return True`},
		},
		Solution: Solution{
			Code: `class Logger:
    def __init__(self):
        # message -> earliest timestamp at which it may print again
        self.next_allowed = {}

    def shouldPrintMessage(self, timestamp: int, message: str) -> bool:
        if timestamp < self.next_allowed.get(message, 0):
            return False
        self.next_allowed[message] = timestamp + 10
        return True`,
			Explanation: "Store, for each message, the earliest time it may print again. Each call is one dictionary lookup and, if the message prints, one update. Builds on Two Sum: the same \"look up, then store\" loop over a stream, but the dictionary maps each message to a time instead of an index, and it lives inside a class because the calls arrive one at a time.",
			BruteForce:  "Keep a list of every printed (timestamp, message) and, for each call, scan the list for the same message within the last 10 seconds. After k prints that is O(k) per call, O(k^2) over a long stream. The waste: only the most recent print of this one message can block it, yet the scan looks at every print of every message.",
			KeyInsight:  "Whether a message may print depends on a single number: when it was last printed. A dictionary keyed by message text returns that number in O(1).",
			WhyItWorks:  "Invariant: next_allowed[m] is (last time m was printed) + 10 for every message printed so far, and messages never printed have no entry (read as 0, and timestamps are at least 0, so they always pass). A call prints exactly when timestamp >= last print + 10, which is the rule. Rejected calls change nothing, so the invariant still holds after them, and an accepted call updates exactly m's entry to the new print time + 10.",
			CommonMistakes: []string{
				"Updating the stored time on every call, including rejected ones. Then a message sent at 1, 5 and 11 is blocked at 11 because the clock restarted at 5. Only record the time when the message is printed.",
				"Using > 10 instead of >= 10 for the gap. A message printed at 5 may print again at exactly 15. Compare timestamp < last + 10 to reject.",
				"Writing if self.last.get(message): to test whether a message was printed. A message printed at time 0 stores 0, which is falsy, so it is printed again at once. Test with 'in' or compare against a default time.",
			},
			TimeComplexity:  "O(1) per call",
			SpaceComplexity: "O(m)",
			Walkthrough: []WalkthroughStep{
				{Title: "Choose what to remember", Explanation: "For each message we only need one number: the earliest time it may print again. Storing last + 10 makes the check a single comparison.", CodeSnippet: "self.next_allowed = {}", LineStart: 3, LineEnd: 4},
				{Title: "Look up before you store", Explanation: "A message never printed reads as 0, which every timestamp passes. If the timestamp is earlier than the stored time, reject without changing anything. With \"ping\" stored as 11, the calls at 3 and 10 are rejected.", CodeSnippet: "if timestamp < self.next_allowed.get(message, 0):\n    return False", LineStart: 7, LineEnd: 8},
				{Title: "Record the current item", Explanation: "The message prints now, so its next allowed time moves to timestamp + 10.", CodeSnippet: "self.next_allowed[message] = timestamp + 10\nreturn True", LineStart: 9, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: scan every printed line",
					Code: `class Logger:
    def __init__(self):
        self.printed = []  # (timestamp, message) for every printed line

    def shouldPrintMessage(self, timestamp: int, message: str) -> bool:
        for when, text in self.printed:
            if text == message and timestamp - when < 10:
                return False
        self.printed.append((timestamp, message))
        return True`,
					TimeComplexity:  "O(k) per call",
					SpaceComplexity: "O(k)",
					Explanation:     "Correct, and a natural first draft, but every call rescans all k lines printed so far.",
				},
				{
					Name: "Queue of the last 10 seconds",
					Code: `from collections import deque

class Logger:
    def __init__(self):
        self.recent = deque()  # (timestamp, message) printed in the last 10 seconds
        self.recent_messages = set()

    def shouldPrintMessage(self, timestamp: int, message: str) -> bool:
        while self.recent and timestamp - self.recent[0][0] >= 10:
            _, old = self.recent.popleft()
            self.recent_messages.discard(old)
        if message in self.recent_messages:
            return False
        self.recent.append((timestamp, message))
        self.recent_messages.add(message)
        return True`,
					TimeComplexity:  "O(1) amortized per call",
					SpaceComplexity: "O(r)",
					Explanation:     "Because timestamps only grow, prints older than 10 seconds can be dropped from the front of a queue. Memory then holds only the r messages printed in the last 10 seconds, instead of every message ever seen.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must a rejected call leave the dictionary unchanged?", Answer: "The rule is about the last time the message was printed, not the last time it was sent. A rejected message was not printed, so recording its time would push its next allowed time later than the rule says, and a message sent at 1, 5 and 11 would wrongly be blocked at 11."},
			{Prompt: "The dictionary keeps every message forever. What if there are millions of distinct messages and memory matters?", Answer: "Since timestamps never decrease, any entry whose next allowed time is already in the past can be deleted without changing any answer. The queue alternative does this automatically: it drops prints older than 10 seconds, so memory is bounded by the messages printed in the last 10 seconds."},
			{Prompt: "What changes if calls could arrive out of time order?", Answer: "Storing only the last print time is no longer enough, because an earlier timestamp could arrive after a later one and both the past and the future around it matter. You would need to keep the sorted print times per message (for example with bisect) and check the nearest print on each side."},
		},
	},
	{
		ID:              "group-anagrams",
		Number:          49,
		Order:           7,
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
		Order:           8,
		Title:           "Top K Frequent Elements",
		Slug:            "top-k-frequent-elements",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "Bucket Sort", "Heap (Priority Queue)", "Counting", "Google", "Google L4", "Google Phone Screen"},
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
		Order:           9,
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
		Order:           10,
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
		ID:              "insert-delete-getrandom-o1",
		Number:          380,
		Order:           11,
		Title:           "Insert Delete GetRandom O(1)",
		Slug:            "insert-delete-getrandom-o1",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "Math", "Design", "Randomized", "Google", "Google L4", "Google Phone Screen"},
		RelatedChapters: []int{2},
		Description: `Design a class RandomizedSet that stores a set of distinct integers and supports three operations, each in O(1) average time:

- RandomizedSet() creates an empty set.
- insert(val) adds val if it is not already present. Returns true if val was added, false if it was already there.
- remove(val) deletes val if it is present. Returns true if val was removed, false if it was not there.
- getRandom() returns one of the values currently in the set, chosen at random so that every value has the same chance of being picked.

getRandom is only called when the set holds at least one value.`,
		Constraints: []string{
			"-2^31 <= val <= 2^31 - 1",
			"At most 2 * 10^5 calls are made in total",
			"getRandom is only called when the set is not empty",
		},
		Examples: []Example{
			{Input: `["RandomizedSet","insert","remove","insert","remove","insert","getRandom"]
[[],[1],[2],[2],[1],[2],[]]`, Output: "[null,true,false,true,true,false,2]", Explanation: "1 is added, removing 2 fails because it is absent, 2 is added, 1 is removed, adding 2 again fails because it is present. The set is now {2}, so getRandom must return 2."},
			{Input: `["RandomizedSet","insert","insert","remove","remove","insert","getRandom"]
[[],[5],[5],[5],[5],[5],[]]`, Output: "[null,true,false,true,false,true,5]", Explanation: "A value can be inserted again after it has been removed."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"RandomizedSet", "insert", "remove", "insert", "remove", "insert", "getRandom"}, "values": [][]int{{}, {1}, {2}, {2}, {1}, {2}, {}}}, Expected: []interface{}{nil, true, false, true, true, false, 2}, Note: "return values of insert and remove, then the one remaining value"},
			{Input: map[string]interface{}{"operations": []string{"RandomizedSet", "insert", "insert", "remove", "remove", "insert", "getRandom"}, "values": [][]int{{}, {5}, {5}, {5}, {5}, {5}, {}}}, Expected: []interface{}{nil, true, false, true, false, true, 5}, Note: "duplicate insert, double remove, re-insert"},
			{Input: map[string]interface{}{"operations": []string{"RandomizedSet", "insert", "insert", "insert", "remove", "remove", "remove", "insert", "getRandom"}, "values": [][]int{{}, {1}, {2}, {3}, {1}, {3}, {2}, {3}, {}}}, Expected: []interface{}{nil, true, true, true, true, true, true, true, 3}, Hidden: true, Note: "the element moved into a hole must get its new position recorded"},
			{Input: map[string]interface{}{"operations": []string{"RandomizedSet", "insert", "insert", "remove", "getRandom", "insert", "remove", "getRandom"}, "values": [][]int{{}, {7}, {8}, {8}, {}, {8}, {7}, {}}}, Expected: []interface{}{nil, true, true, true, 7, true, true, 8}, Hidden: true, Note: "removing the value that is already last in the array"},
			{Input: map[string]interface{}{"operations": []string{"RandomizedSet", "insert", "insert", "insert", "insert", "insert", "insert", "remove", "remove", "remove", "remove", "remove", "getRandom", "getRandom", "getRandom"}, "values": [][]int{{}, {1}, {2}, {3}, {4}, {5}, {6}, {1}, {2}, {3}, {4}, {5}, {}, {}, {}}}, Expected: []interface{}{nil, true, true, true, true, true, true, true, true, true, true, true, 6, 6, 6}, Hidden: true, Note: "after five removals getRandom may only ever return the survivor"},
			{Input: map[string]interface{}{"operations": []string{"RandomizedSet", "insert", "insert", "remove", "getRandom", "remove"}, "values": [][]int{{}, {-2147483648}, {2147483647}, {-2147483648}, {}, {0}}}, Expected: []interface{}{nil, true, true, true, 2147483647, false}, Hidden: true, Note: "extreme 32-bit values; removing an absent 0"},
			{Input: map[string]interface{}{"operations": []string{"RandomizedSet", "remove", "insert", "getRandom"}, "values": [][]int{{}, {1}, {1}, {}}}, Expected: []interface{}{nil, false, true, 1}, Hidden: true, Note: "remove on an empty set"},
		},
		TimeComplexity:  "O(1)",
		SpaceComplexity: "O(n)",
		StarterCode: `class RandomizedSet:
    def __init__(self):
        pass

    def insert(self, val: int) -> bool:
        pass

    def remove(self, val: int) -> bool:
        pass

    def getRandom(self) -> int:
        pass`,
		Harness: Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "A set gives O(1) insert and remove but can't pick a random element quickly; a list can pick random.choice in O(1) but deleting from the middle is slow. What if you used both, and what would the list need to look like so random.choice is fair?"},
			{Level: 2, Type: "insight", Content: "Keep the values packed in a list (so random.choice is uniform and O(1)) and a dictionary value -> index into that list (the value -> index map from Two Sum, Chapter 2). To delete from the middle in O(1), copy the last value into the hole, update that value's index, and pop the end. Order inside the list doesn't matter for a set.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Choose what to remember: a list of values with no gaps, and a dict value -> its index.\n2. Insert: if val is new, record its index as len(list), then append it.\n3. Remove: if val is present, move the last value into val's slot, update the last value's index, pop the list's end, and delete val from the dict.\n4. Pick uniformly: random.choice on the list."},
			{Level: 4, Type: "skeleton", Content: `import random

class RandomizedSet:
    def __init__(self):
        self.values = []
        self.index_of = {}

    def insert(self, val: int) -> bool:
        if val in self.index_of:
            return False
        self.index_of[val] = ____
        self.values.append(val)
        return True

    def remove(self, val: int) -> bool:
        if val not in self.index_of:
            return False
        position = self.index_of[val]
        last = self.values[-1]
        ____
        ____
        self.values.pop()
        del self.index_of[val]
        return True

    def getRandom(self) -> int:
        return random.choice(self.values)`},
		},
		Solution: Solution{
			Code: `import random

class RandomizedSet:
    def __init__(self):
        self.values = []    # the members, packed with no gaps
        self.index_of = {}  # value -> its position in self.values

    def insert(self, val: int) -> bool:
        if val in self.index_of:
            return False
        self.index_of[val] = len(self.values)
        self.values.append(val)
        return True

    def remove(self, val: int) -> bool:
        if val not in self.index_of:
            return False
        position = self.index_of[val]
        last = self.values[-1]
        self.values[position] = last
        self.index_of[last] = position
        self.values.pop()
        del self.index_of[val]
        return True

    def getRandom(self) -> int:
        return random.choice(self.values)`,
			Explanation: "Pair a packed list of the values (for uniform random picks) with a dictionary from each value to its list index (for O(1) membership and lookup). Removal fills the hole with the last element, so the list never has gaps. Builds on Logger Rate Limiter, the first class in this ladder that keeps a dictionary between calls, and on Two Sum's value -> index map; the new idea is keeping two structures in sync so that each covers the other's weakness.",
			BruteForce:  "Keep only a list: insert checks 'val in list' (O(n)), remove calls list.remove (O(n), it searches and then shifts every later element), getRandom uses random.choice (O(1)). Or keep only a set: insert and remove are O(1) but getRandom must copy the set into a list first, O(n). The waste: the list scans for a value whose position could be remembered, and shifting elements after a delete preserves an order that a set doesn't care about.",
			KeyInsight:  "A set has no order, so the list may be rearranged freely. Moving the last element into the deleted slot fills the hole in O(1), and the dictionary tells you exactly where the hole is.",
			WhyItWorks:  "Invariant: values holds each member exactly once with no gaps, and index_of[v] == i exactly when values[i] == v. Insert appends at index len(values) and records that index, so the invariant holds. Remove writes last into position and records index_of[last] = position, pops the now-duplicate end, then deletes val's entry; if val was itself the last element, the write and the index update change nothing and the delete still removes it, which is why the delete must come after the update. Because values has no gaps and contains each member once, random.choice picks each member with probability 1/size.",
			CommonMistakes: []string{
				"Moving the last value into the hole but not updating its index in the dictionary. Its stored index then points past the end of the list (insert 1, 2, 3; remove 1; remove 3 crashes). Always set index_of[last] = position.",
				"Deleting index_of[val] before writing index_of[last] = position. When val is the last element, the second line puts val straight back into the dictionary, so a later insert(val) returns false. Update first, delete last.",
				"Storing the values only in a set and calling random.choice(list(self.items)). It is correct but copies the whole set on every getRandom, which is O(n). Keep the packed list alongside the dictionary.",
			},
			TimeComplexity:  "O(1) average per operation",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Choose what to remember", Explanation: "The list is what makes a uniform random pick O(1); the dictionary is what makes finding a value O(1). Each answers a question the other can't.", CodeSnippet: "self.values = []    # the members, packed with no gaps\nself.index_of = {}  # value -> its position in self.values", LineStart: 5, LineEnd: 6},
				{Title: "Insert: append and record the position", Explanation: "A new value goes at the end, at index len(values), and the dictionary remembers that index. Inserting 1, 2, 3 gives values = [1, 2, 3] and index_of = {1: 0, 2: 1, 3: 2}.", CodeSnippet: "if val in self.index_of:\n    return False\nself.index_of[val] = len(self.values)\nself.values.append(val)", LineStart: 9, LineEnd: 12},
				{Title: "Remove: fill the hole with the last item", Explanation: "Removing 1 from [1, 2, 3]: position = 0, last = 3. Write 3 into slot 0 and record index_of[3] = 0, giving [3, 2, 3]; pop the end to get [3, 2]; delete 1 from the dictionary.", CodeSnippet: "position = self.index_of[val]\nlast = self.values[-1]\nself.values[position] = last\nself.index_of[last] = position\nself.values.pop()\ndel self.index_of[val]", LineStart: 18, LineEnd: 23},
				{Title: "Pick uniformly", Explanation: "The list holds every member exactly once with no gaps, so random.choice gives each the same chance.", CodeSnippet: "return random.choice(self.values)", LineStart: 26, LineEnd: 27},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: a plain list",
					Code: `import random

class RandomizedSet:
    def __init__(self):
        self.values = []

    def insert(self, val: int) -> bool:
        if val in self.values:
            return False
        self.values.append(val)
        return True

    def remove(self, val: int) -> bool:
        if val not in self.values:
            return False
        self.values.remove(val)
        return True

    def getRandom(self) -> int:
        return random.choice(self.values)`,
					TimeComplexity:  "O(n) for insert and remove, O(1) for getRandom",
					SpaceComplexity: "O(n)",
					Explanation:     "Correct and simple, but every membership test scans the list and list.remove shifts the elements after the removed one.",
				},
				{
					Name: "A set, copied on every getRandom",
					Code: `import random

class RandomizedSet:
    def __init__(self):
        self.items = set()

    def insert(self, val: int) -> bool:
        if val in self.items:
            return False
        self.items.add(val)
        return True

    def remove(self, val: int) -> bool:
        if val not in self.items:
            return False
        self.items.remove(val)
        return True

    def getRandom(self) -> int:
        return random.choice(list(self.items))`,
					TimeComplexity:  "O(1) average for insert and remove, O(n) for getRandom",
					SpaceComplexity: "O(n)",
					Explanation:     "Moves the cost to the other side: membership is fast, but a set can't be indexed, so each random pick copies it into a list first.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it allowed to move the last element into the removed element's slot?", Answer: "The structure represents a set, which has no order. Only which values are in the list matters (for membership and for a fair random pick), not where they sit, so swapping positions changes nothing a caller can observe. The dictionary is updated so it still knows where every value lives."},
			{Prompt: "What changes if duplicates are allowed (LeetCode 381) and getRandom must be weighted by how many copies each value has?", Answer: "The list holds every copy, so random.choice is automatically weighted. The dictionary must map each value to a set of indices instead of one index. Remove takes any index from val's set, moves the last element there, and updates the last element's index set by removing its old index and adding the new one."},
			{Prompt: "Why must del self.index_of[val] come after self.index_of[last] = position?", Answer: "When val is the last element, last == val. Writing index_of[last] = position re-creates val's entry. If the delete ran first, that write would leave val in the dictionary although it is no longer in the list. Running the delete last removes it in every case."},
		},
	},
	{
		ID:              "product-of-array-except-self",
		Number:          238,
		Order:           12,
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
		Order:           13,
		Title:           "Longest Consecutive Sequence",
		Slug:            "longest-consecutive-sequence",
		Difficulty:      "Medium",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "Union Find", "Google", "Google L4", "Google Onsite"},
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
	{
		ID:              "first-missing-positive",
		Number:          41,
		Order:           14,
		Title:           "First Missing Positive",
		Slug:            "first-missing-positive",
		Difficulty:      "Hard",
		Category:        "arrays-hashing",
		Tags:            []string{"Array", "Hash Table", "Google", "Google L5", "Google Phone Screen"},
		RelatedChapters: []int{2, 1},
		Description: `You are given an unsorted list of integers nums. It may contain negative numbers, zeros and repeated values. Return the smallest positive integer (1, 2, 3, ...) that does not appear anywhere in nums.

Your algorithm must run in O(n) time and use only O(1) extra memory. You are allowed to modify nums.`,
		Constraints: []string{
			"1 <= nums.length <= 10^5",
			"-2^31 <= nums[i] <= 2^31 - 1",
		},
		Examples: []Example{
			{Input: "nums = [1,2,0]", Output: "3", Explanation: "1 and 2 are present, so the smallest missing positive is 3."},
			{Input: "nums = [3,4,-1,1]", Output: "2", Explanation: "1 is present but 2 is not."},
			{Input: "nums = [7,8,9,11,12]", Output: "1", Explanation: "Not even 1 is present."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 0}}, Expected: 3, Note: "1 and 2 present, answer just past them"},
			{Input: map[string]interface{}{"nums": []int{3, 4, -1, 1}}, Expected: 2, Note: "gap in the middle, with a negative value"},
			{Input: map[string]interface{}{"nums": []int{7, 8, 9, 11, 12}}, Expected: 1, Note: "every value is larger than n"},
			{Input: map[string]interface{}{"nums": []int{1}}, Expected: 2, Hidden: true, Note: "single element present: the answer is n + 1"},
			{Input: map[string]interface{}{"nums": []int{2}}, Expected: 1, Hidden: true, Note: "single element, 1 missing"},
			{Input: map[string]interface{}{"nums": []int{1, 1}}, Expected: 2, Hidden: true, Note: "duplicates: a swap loop that ignores them never stops"},
			{Input: map[string]interface{}{"nums": []int{2, 1}}, Expected: 3, Hidden: true, Note: "needs a swap; Python's a[i], a[a[i]-1] = ... pitfall loops forever here"},
			{Input: map[string]interface{}{"nums": []int{-5, -3, 0}}, Expected: 1, Hidden: true, Note: "no positive values at all"},
			{Input: map[string]interface{}{"nums": []int{5, 4, 3, 2, 1}}, Expected: 6, Hidden: true, Note: "a reversed permutation of 1..n"},
			{Input: map[string]interface{}{"nums": []int{2147483647, 1, 3}}, Expected: 2, Hidden: true, Note: "an extreme value far beyond n is ignored"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def firstMissingPositive(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "With n numbers, how large can the answer possibly be? And if a set of seen values is not allowed, is there an array of size n you already own that could record which of 1..n you have seen?"},
			{Level: 2, Type: "insight", Content: "The answer is always between 1 and n + 1, because n numbers can cover at most 1..n. So only values 1..n matter, and each has a natural home: value v belongs at index v - 1. Use nums itself as the hash table (Chapter 2's \"set of seen values\", stored in place): swap every value in 1..n into its home, then the first index i with nums[i] != i + 1 gives the answer i + 1.", ChapterRef: 2},
			{Level: 3, Type: "plan", Content: "1. Bound the answer: it lies in 1..n + 1, so values outside 1..n can be ignored.\n2. Send each value home: while nums[i] is in 1..n and its home slot doesn't already hold it, swap it into slot nums[i] - 1.\n3. Find the first slot that's wrong: scan for the first i where nums[i] != i + 1.\n4. Return the answer: that i + 1, or n + 1 if every slot is right."},
			{Level: 4, Type: "skeleton", Content: `def firstMissingPositive(nums):
    n = len(nums)
    for i in range(n):
        while ____ and ____:
            home = nums[i] - 1
            ____
    for i in range(n):
        if nums[i] != i + 1:
            return ____
    return ____`},
		},
		Solution: Solution{
			Code: `def firstMissingPositive(nums):
    n = len(nums)
    for i in range(n):
        # Keep swapping nums[i] into its home slot nums[i] - 1
        while 1 <= nums[i] <= n and nums[nums[i] - 1] != nums[i]:
            home = nums[i] - 1
            nums[i], nums[home] = nums[home], nums[i]
    for i in range(n):
        if nums[i] != i + 1:
            return i + 1
    return n + 1`,
			Explanation: "The answer lies in 1..n + 1, so every value in 1..n gets a home slot (value v at index v - 1). Swap values home in place, then the first slot not holding its own value names the missing number. Builds on Longest Consecutive Sequence: there a set answered \"is v present?\" for a run of values; here the same membership question is answered without a set, by turning the input array itself into the lookup table.",
			BruteForce:  "Try 1, 2, 3, ... and scan the whole list for each candidate: up to n + 1 candidates times n comparisons, O(n^2). Putting the values in a set first makes it O(n) time but O(n) extra memory, which the problem forbids. The waste: the scan forgets what it learned about the array between candidates, and the set duplicates storage that the input array already provides.",
			KeyInsight:  "Only the values 1..n can affect the answer, and there are exactly n slots. So the array can serve as its own presence table: slot v - 1 records whether v was seen, by holding v.",
			WhyItWorks:  "Bound: n numbers can contain at most n distinct values from 1..n, so some value in 1..n + 1 is missing. Placement: each swap puts at least one value v into its home slot v - 1, where it then stays (the loop condition refuses to move a value already home), so there are at most n useful swaps in total and the whole first loop is O(n) despite the nested while. Afterwards, for every v in 1..n that occurs in nums, slot v - 1 holds v: if it didn't, the while loop at the index where v ended up would have moved it. So the first i with nums[i] != i + 1 is exactly the smallest missing positive, and if none exists, 1..n are all present and the answer is n + 1.",
			CommonMistakes: []string{
				"Writing the swap as nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]. Python assigns left to right, so nums[i] changes before nums[nums[i] - 1] is evaluated on the left, and the value lands in the wrong slot; [2,1] loops forever. Compute home = nums[i] - 1 first.",
				"Using nums[i] != i + 1 as the loop condition. With duplicates such as [1,1], the second 1 is never at index 1 and its home already holds a 1, so it swaps 1 with 1 forever. Stop when the home slot already holds the value: nums[nums[i] - 1] != nums[i].",
				"Returning n instead of n + 1 when every slot is correct. For [1] the answer is 2, and for [5,4,3,2,1] it is 6.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Bound the answer", Explanation: "With n slots, the answer is between 1 and n + 1, so only values 1..n need a place. Negatives, zero and anything above n are left wherever they are.", CodeSnippet: "n = len(nums)", LineStart: 2, LineEnd: 2},
				{Title: "Send each value home", Explanation: "Value v belongs at index v - 1. For [3,4,-1,1]: i = 0 swaps 3 home ([-1,4,3,1]); i = 1 swaps 4 home ([-1,1,3,4]), then 1 home ([1,-1,3,4]). Each swap settles one value for good, so the total work is O(n).", CodeSnippet: "while 1 <= nums[i] <= n and nums[nums[i] - 1] != nums[i]:\n    home = nums[i] - 1\n    nums[i], nums[home] = nums[home], nums[i]", LineStart: 5, LineEnd: 7},
				{Title: "Find the first slot that's wrong", Explanation: "Slot i should hold i + 1. In [1,-1,3,4], slot 1 holds -1, so 2 is missing.", CodeSnippet: "for i in range(n):\n    if nums[i] != i + 1:\n        return i + 1", LineStart: 8, LineEnd: 10},
				{Title: "Return the answer", Explanation: "Every slot was right, so 1..n are all present and the first missing positive is n + 1.", CodeSnippet: "return n + 1", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: try each candidate",
					Code: `def firstMissingPositive(nums):
    candidate = 1
    while candidate in nums:
        candidate += 1
    return candidate`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Scans the whole list for 1, then for 2, and so on. Correct and memory-free, but quadratic.",
				},
				{
					Name: "Set of values",
					Code: `def firstMissingPositive(nums):
    present = set(nums)
    candidate = 1
    while candidate in present:
        candidate += 1
    return candidate`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Contains Duplicate's set turns each check into O(1), and at most n + 1 candidates are tried. Optimal time, but it breaks the O(1) memory requirement; this is the answer to give first in an interview before optimizing.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "The first loop has a while loop inside a for loop. Why is it O(n) and not O(n^2)?", Answer: "Every swap that happens puts some value v into slot v - 1, and the loop condition never moves a value that is already home. Each of the n slots can be filled with its own value at most once, so there are at most n swaps over the whole run, plus one failed check per index."},
			{Prompt: "Why can every value greater than n be ignored?", Answer: "The answer is at most n + 1, since n numbers can't cover all of 1..n + 1. A value above n can't be the answer and can't be needed to rule out a smaller candidate, so it doesn't matter where it ends up."},
			{Prompt: "What if you were not allowed to modify nums?", Answer: "Then the array can't be your presence table. Use a set (O(n) memory) or a boolean list of size n + 1 marking which of 1..n appear, then scan it for the first unmarked value. Both are O(n) time and O(n) memory; the O(1) memory bound relies on reusing the input."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, ArraysHashingProblems...)
}
