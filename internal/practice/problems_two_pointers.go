package practice

// TwoPointersProblems is the two-pointers ladder (chapter 3): compare the two
// ends -> use sorted order to decide which end to move -> fix one value and
// run the pair search -> move the pointer that can't be part of a better
// answer -> carry a running maximum from each side.
var TwoPointersProblems = []*Problem{
	{
		ID:              "valid-palindrome",
		Number:          125,
		Order:           1,
		Title:           "Valid Palindrome",
		Slug:            "valid-palindrome",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Two Pointers", "String"},
		RelatedChapters: []int{3},
		Description: `A string counts as a palindrome here if it reads the same forwards and backwards after two clean-up steps: turn every uppercase letter into lowercase, and drop every character that is not a letter or a digit (spaces, punctuation, underscores and so on).

Given a string s, return true if it is a palindrome under this rule, and false otherwise. If nothing is left after cleaning, the empty string reads the same both ways, so return true.`,
		Constraints: []string{
			"1 <= s.length <= 2 * 10^5",
			"s contains only printable ASCII characters",
		},
		Examples: []Example{
			{Input: `s = "A man, a plan, a canal: Panama"`, Output: "true", Explanation: "After cleaning it becomes \"amanaplanacanalpanama\", which is the same reversed."},
			{Input: `s = "race a car"`, Output: "false", Explanation: "After cleaning it becomes \"raceacar\"; reversed that is \"racaecar\"."},
			{Input: `s = " "`, Output: "true", Explanation: "Cleaning leaves the empty string."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "A man, a plan, a canal: Panama"}, Expected: true, Note: "letters only after cleaning; mixed case"},
			{Input: map[string]interface{}{"s": "race a car"}, Expected: false, Note: "not a palindrome after cleaning"},
			{Input: map[string]interface{}{"s": " "}, Expected: true, Note: "nothing left after cleaning counts as a palindrome"},
			{Input: map[string]interface{}{"s": "0P"}, Expected: false, Hidden: true, Note: "digits count; 0 and p differ"},
			{Input: map[string]interface{}{"s": "ab_a"}, Expected: true, Hidden: true, Note: "underscore is not alphanumeric"},
			{Input: map[string]interface{}{"s": "a."}, Expected: true, Hidden: true, Note: "punctuation at the end"},
			{Input: map[string]interface{}{"s": ".,"}, Expected: true, Hidden: true, Note: "only punctuation"},
			{Input: map[string]interface{}{"s": "No 'x' in Nixon"}, Expected: true, Hidden: true, Note: "apostrophes and capitals"},
			{Input: map[string]interface{}{"s": "12321"}, Expected: true, Hidden: true, Note: "digits only, odd length"},
			{Input: map[string]interface{}{"s": "abca"}, Expected: false, Hidden: true, Note: "mismatch in the middle pair"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def isPalindrome(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "A palindrome's first character must match its last, its second must match its second-to-last, and so on. How could you check those pairs without building a cleaned copy of the string?"},
			{Level: 2, Type: "insight", Content: "Put one pointer at each end and walk them toward each other, comparing the characters they point at. Before comparing, step each pointer past characters that aren't letters or digits. This is the opposite-ends two-pointer pattern from Chapter 3.", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Place the pointers: left at index 0, right at the last index.\n2. Skip what doesn't count: move left forward and right backward past non-alphanumeric characters (without crossing).\n3. Compare and decide: if the lowercased characters differ, return False.\n4. Move both pointers inward and repeat.\n5. Stop when they meet: every pair matched, so return True."},
			{Level: 4, Type: "skeleton", Content: `def isPalindrome(s):
    left, right = 0, len(s) - 1
    while left < right:
        while left < right and ____:
            left += 1
        while left < right and ____:
            right -= 1
        if ____:
            return False
        left += 1
        right -= 1
    return True`},
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
			Explanation: "Walk one pointer in from each end, skipping characters that don't count and comparing the rest case-insensitively. The plainest form of Chapter 3's opposite-ends pattern: each comparison settles one mirror pair for good, so both pointers move inward. Compare Valid Anagram (Chapter 2), where order didn't matter and counting was enough; here order is the whole point.",
			BruteForce:  "Build the cleaned string, reverse it, and compare: O(n) time, but O(n) extra memory for two new strings. A slower brute force compares s with its reverse character by character after each removal. The waste in the first approach is memory: we copy the whole string just to read it from both ends, which two indices can do in place.",
			KeyInsight:  "A palindrome check only ever compares position i from the left with position i from the right. Two indices moving inward visit exactly those pairs, with no copy.",
			WhyItWorks:  "Invariant: every alphanumeric character outside the window [left, right] has already been matched with its mirror partner. The inner loops move each pointer to the next character that counts (never crossing the other), so s[left] and s[right] are mirror partners in the cleaned string. If they differ, the cleaned string isn't a palindrome. If they match, the window shrinks and the invariant still holds. When left >= right, at most one character (the middle) is left unmatched, and a middle character always equals itself.",
			CommonMistakes: []string{
				"Skipping only spaces, not punctuation or underscores. Use str.isalnum() so \"ab_a\" is cleaned to \"aba\".",
				"Forgetting that digits count. \"0P\" must return false: 0 and p are both alphanumeric and differ. isalpha() would drop the 0 and wrongly return true.",
				"Letting the skip loops run without the left < right guard, which runs off the end of the string on input like \".,\" and raises IndexError.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Place the pointers", Explanation: "One index at each end of the original string. No cleaned copy is built.", CodeSnippet: "left, right = 0, len(s) - 1", LineStart: 2, LineEnd: 2},
				{Title: "Skip what doesn't count", Explanation: "Move each pointer past spaces and punctuation, but never past the other pointer.", CodeSnippet: "while left < right and not s[left].isalnum():\n    left += 1\nwhile left < right and not s[right].isalnum():\n    right -= 1", LineStart: 4, LineEnd: 7},
				{Title: "Compare and decide", Explanation: "The two characters are mirror partners. Lowercase both before comparing so 'A' matches 'a'.", CodeSnippet: "if s[left].lower() != s[right].lower():\n    return False", LineStart: 8, LineEnd: 9},
				{Title: "Move both pointers inward", Explanation: "This pair is settled, so both pointers move toward the middle. When they meet, every pair matched.", CodeSnippet: "left += 1\nright -= 1\nreturn True", LineStart: 10, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Clean, then compare with the reverse",
					Code: `def isPalindrome(s):
    cleaned = [ch.lower() for ch in s if ch.isalnum()]
    return cleaned == cleaned[::-1]`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Short and clear: build the cleaned list and compare it with its reverse. Uses O(n) extra memory for the copies.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the loop stop at left < right instead of left <= right?", Answer: "When left == right both pointers sit on the same middle character, which trivially matches itself. When left > right every pair has been compared. Either way there is nothing left to check."},
			{Prompt: "What if you were allowed to delete at most one character to make it a palindrome (LeetCode 680)?", Answer: "Run the same two pointers. At the first mismatch, try both options: skip the left character or skip the right one, and check whether the remaining window is a palindrome with a plain two-pointer pass. That's still O(n)."},
		},
	},
	{
		ID:              "two-sum-ii",
		Number:          167,
		Order:           2,
		Title:           "Two Sum II - Input Array Is Sorted",
		Slug:            "two-sum-ii-input-array-is-sorted",
		Difficulty:      "Medium",
		Category:        "two-pointers",
		Tags:            []string{"Array", "Two Pointers", "Binary Search"},
		RelatedChapters: []int{3, 2},
		Description: `You are given a list of integers numbers sorted in non-decreasing order (each value is at least the one before it) and an integer target. Exactly one pair of different positions has values that add up to target.

Return the two positions as [index1, index2] with index1 < index2, counting positions from 1 (the first element is position 1, not 0).

Your solution must use only O(1) extra memory, so a hash map of seen values (as in Two Sum) is not allowed.`,
		Constraints: []string{
			"2 <= numbers.length <= 3 * 10^4",
			"-1000 <= numbers[i] <= 1000",
			"numbers is sorted in non-decreasing order",
			"-1000 <= target <= 1000",
			"Exactly one valid pair exists",
		},
		Examples: []Example{
			{Input: "numbers = [2,7,11,15], target = 9", Output: "[1,2]", Explanation: "2 + 7 = 9. They are at positions 1 and 2 when counting from 1."},
			{Input: "numbers = [2,3,4], target = 6", Output: "[1,3]"},
			{Input: "numbers = [-1,0], target = -1", Output: "[1,2]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"numbers": []int{2, 7, 11, 15}, "target": 9}, Expected: []int{1, 2}, Note: "answer is 1-indexed"},
			{Input: map[string]interface{}{"numbers": []int{2, 3, 4}, "target": 6}, Expected: []int{1, 3}, Note: "pair is the two ends"},
			{Input: map[string]interface{}{"numbers": []int{-1, 0}, "target": -1}, Expected: []int{1, 2}, Note: "negative numbers, length 2"},
			{Input: map[string]interface{}{"numbers": []int{1, 1, 3, 5}, "target": 2}, Expected: []int{1, 2}, Hidden: true, Note: "two equal values"},
			{Input: map[string]interface{}{"numbers": []int{-10, -3, 0, 4, 9, 12}, "target": 1}, Expected: []int{2, 4}, Hidden: true, Note: "pair in the middle after several moves"},
			{Input: map[string]interface{}{"numbers": []int{0, 0, 3, 4}, "target": 0}, Expected: []int{1, 2}, Hidden: true, Note: "target 0"},
			{Input: map[string]interface{}{"numbers": []int{1, 2, 3, 4, 4, 9, 56, 90}, "target": 8}, Expected: []int{4, 5}, Hidden: true, Note: "duplicate pair inside, big values on the right"},
			{Input: map[string]interface{}{"numbers": []int{-1000, -1, 0, 1, 1000}, "target": 999}, Expected: []int{2, 5}, Hidden: true, Note: "extreme values from the constraints"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def twoSum(numbers, target):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Add the smallest and the largest number. If that sum is too small, can the smallest number be part of the answer with any partner at all?"},
			{Level: 2, Type: "insight", Content: "Start with left at the smallest value and right at the largest. If their sum is too small, numbers[left] can't work even with the biggest partner, so discard it (left += 1). If the sum is too big, numbers[right] can't work even with the smallest partner, so discard it (right -= 1). Sorted order replaces Two Sum's hash map. This is the opposite-ends pattern from Chapter 3.", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Place the pointers: left = 0, right = last index.\n2. Compare and decide: compute the sum of the two values.\n3. Record the answer: if it equals target, return [left + 1, right + 1].\n4. Move the pointer that can't help: sum too small -> move left right; sum too big -> move right left.\n5. Stop when they meet (never happens here, because a pair is guaranteed)."},
			{Level: 4, Type: "skeleton", Content: `def twoSum(numbers, target):
    left, right = 0, len(numbers) - 1
    while left < right:
        total = numbers[left] + numbers[right]
        if total == target:
            return ____
        if ____:
            left += 1
        else:
            ____
    return []`},
		},
		Solution: Solution{
			Code: `def twoSum(numbers, target):
    left, right = 0, len(numbers) - 1
    while left < right:
        total = numbers[left] + numbers[right]
        if total == target:
            return [left + 1, right + 1]
        if total < target:
            left += 1  # even the largest partner is too small for numbers[left]
        else:
            right -= 1  # even the smallest partner is too big for numbers[right]
    return []`,
			Explanation: "Pointers at both ends: if the sum is too small, the left value is useless and moves right; if too big, the right value is useless and moves left. Builds on Valid Palindrome's opposite-ends walk, but now the comparison decides which single pointer moves. Contrast with Two Sum (arrays-hashing), which solves the unsorted version in O(n) time but needs O(n) memory for its hash map; sorted input lets us get O(n) time with O(1) memory.",
			BruteForce:  "Try every pair: O(n^2), about 450 million pair checks for n = 30,000. Two Sum's hash map would be O(n) time but breaks the O(1)-memory rule. Binary search for target - numbers[i] for each i is O(n log n) with O(1) memory. The waste in all of these: they ignore that one comparison of the two extremes can rule out an entire value at once.",
			KeyInsight:  "With sorted input, numbers[left] + numbers[right] is the largest sum numbers[left] can make and the smallest sum numbers[right] can make. So one comparison with target proves one of the two values is in no valid pair.",
			WhyItWorks:  "Invariant: if the answer pair is (a, b), then left <= a and b <= right, i.e. no pointer has stepped past an answer index. Suppose total < target. Every partner for numbers[left] inside the window is at most numbers[right] (sorted), so every sum with numbers[left] is at most total < target: numbers[left] belongs to no pair, so left can't be a, and moving it keeps the invariant. The case total > target is symmetric for right. Each step shrinks the window by one, so after at most n - 1 steps the pointers reach exactly a and b.",
			CommonMistakes: []string{
				"Returning 0-based indices. The problem counts from 1, so return [left + 1, right + 1].",
				"Moving both pointers after a miss, or moving the wrong one. Only the value proven useless may be discarded: too small -> left += 1, too big -> right -= 1.",
				"Reaching for a dict as in Two Sum. It passes the tests but violates the O(1)-memory requirement, which is the whole point of this problem.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Place the pointers", Explanation: "Smallest value on the left, largest on the right.", CodeSnippet: "left, right = 0, len(numbers) - 1", LineStart: 2, LineEnd: 2},
				{Title: "Compare and decide", Explanation: "One sum tells us whether we are on target, below or above.", CodeSnippet: "total = numbers[left] + numbers[right]", LineStart: 4, LineEnd: 4},
				{Title: "Record the answer", Explanation: "The sum hits target: convert both indices to 1-based positions and return.", CodeSnippet: "if total == target:\n    return [left + 1, right + 1]", LineStart: 5, LineEnd: 6},
				{Title: "Move the pointer that can't help", Explanation: "For [2,7,11,15], target 9: 2 + 15 = 17 is too big, so 15 is useless; 2 + 11 = 13 too big, drop 11; 2 + 7 = 9 found.", CodeSnippet: "if total < target:\n    left += 1  # even the largest partner is too small for numbers[left]\nelse:\n    right -= 1  # even the smallest partner is too big for numbers[right]", LineStart: 7, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: try every pair",
					Code: `def twoSum(numbers, target):
    for i in range(len(numbers)):
        for j in range(i + 1, len(numbers)):
            if numbers[i] + numbers[j] == target:
                return [i + 1, j + 1]
    return []`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Meets the memory rule but ignores the sorted order, so it is far too slow for 30,000 numbers.",
				},
				{
					Name: "Binary search for each partner",
					Code: `import bisect

def twoSum(numbers, target):
    for i in range(len(numbers)):
        need = target - numbers[i]
        j = bisect.bisect_left(numbers, need, i + 1)
        if j < len(numbers) and numbers[j] == need:
            return [i + 1, j + 1]
    return []`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Uses the sorted order to find each partner by binary search (Chapter 6). Correct with O(1) memory, but a log factor slower than two pointers.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to discard numbers[left] when the sum is too small?", Answer: "numbers[right] is the largest value still in play, so numbers[left] + numbers[right] is the biggest sum numbers[left] can form with any remaining partner. If even that is below target, every pair using numbers[left] is below target, so it can't be in the answer."},
			{Prompt: "Two Sum (arrays-hashing) and this problem ask almost the same question. When would you pick each technique?", Answer: "Unsorted input, indices needed: the hash map (O(n) time, O(n) memory); sorting would scramble the indices. Sorted input or a memory limit: two pointers (O(n) time, O(1) memory). If the input is unsorted but you only need the values, you could sort first and use two pointers for O(n log n) time."},
			{Prompt: "What if there could be several valid pairs and you had to count them?", Answer: "On a match, don't return; count it and move inward. With duplicates you must count carefully: if the values differ, count the run of equal values on each side and add the product; if they are equal, all k copies pair with each other, k(k-1)/2 pairs."},
		},
	},
	{
		ID:              "3sum",
		Number:          15,
		Order:           3,
		Title:           "3Sum",
		Slug:            "3sum",
		Difficulty:      "Medium",
		Category:        "two-pointers",
		Tags:            []string{"Array", "Two Pointers", "Sorting"},
		RelatedChapters: []int{3, 2},
		Description: `You are given a list of integers nums. Find every triplet of values [a, b, c] taken from three different positions such that a + b + c == 0.

Return each distinct triplet once. Two triplets count as the same if they contain the same values, even if they come from different positions (so [-1, 0, 1] and [0, 1, -1] are the same triplet). Return the triplets in any order, and the values inside a triplet in any order. If no triplet exists, return an empty list.`,
		Constraints: []string{
			"3 <= nums.length <= 3000",
			"-10^5 <= nums[i] <= 10^5",
		},
		Examples: []Example{
			{Input: "nums = [-1,0,1,2,-1,-4]", Output: "[[-1,-1,2],[-1,0,1]]", Explanation: "-1 + -1 + 2 = 0 and -1 + 0 + 1 = 0. The triplet [-1,0,1] can be built from either -1, but it is listed only once."},
			{Input: "nums = [0,1,1]", Output: "[]", Explanation: "The only triplet sums to 2."},
			{Input: "nums = [0,0,0]", Output: "[[0,0,0]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{-1, 0, 1, 2, -1, -4}}, Expected: [][]int{{-1, -1, 2}, {-1, 0, 1}}, Note: "two triplets; the repeated -1 must not create a duplicate triplet"},
			{Input: map[string]interface{}{"nums": []int{0, 1, 1}}, Expected: [][]int{}, Note: "no triplet sums to zero"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0}}, Expected: [][]int{{0, 0, 0}}, Note: "three zeros"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0, 0}}, Expected: [][]int{{0, 0, 0}}, Hidden: true, Note: "four zeros still give one triplet"},
			{Input: map[string]interface{}{"nums": []int{-2, 0, 1, 1, 2}}, Expected: [][]int{{-2, 0, 2}, {-2, 1, 1}}, Hidden: true, Note: "triplet using a repeated value once each"},
			{Input: map[string]interface{}{"nums": []int{-1, 0, 1, 0}}, Expected: [][]int{{-1, 0, 1}}, Hidden: true, Note: "duplicate middle values"},
			{Input: map[string]interface{}{"nums": []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}}, Expected: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}, Hidden: true, Note: "many duplicates on both sides"},
			{Input: map[string]interface{}{"nums": []int{1, 2, -2, -1}}, Expected: [][]int{}, Hidden: true, Note: "no zero-sum triplet among mixed signs"},
			{Input: map[string]interface{}{"nums": []int{3, -2, -1, 0}}, Expected: [][]int{{-2, -1, 3}}, Hidden: true, Note: "unsorted input, exactly one triplet"},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(1) extra (besides sorting and the output)",
		StarterCode:     "def threeSum(nums):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered-nested"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you fix the first number a, what problem is left for the other two numbers? Have you solved that problem before, and what did it need from the input?"},
			{Level: 2, Type: "insight", Content: "Sort the list. For each index i, you need two numbers to the right of i that sum to -nums[i]: that's Two Sum II, solved with opposite-end pointers in O(n). Sorting also puts equal values next to each other, so you can skip a value equal to the previous one to avoid duplicate triplets. Chapter 3 covers sorting first so two pointers can work.", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Sort the input.\n2. Fix the first value: loop i over the indices; skip i if nums[i] equals nums[i - 1] (same first value, same triplets).\n3. Place the pointers: left = i + 1, right = last index.\n4. Compare and decide: if the three sum below 0 move left right; above 0 move right left.\n5. Record the answer: on 0, save the triplet, move both pointers, and skip duplicates: advance left past values equal to the one just used."},
			{Level: 4, Type: "skeleton", Content: `def threeSum(nums):
    nums.sort()
    triplets = []
    for i in range(len(nums) - 2):
        if ____:
            continue
        left, right = i + 1, len(nums) - 1
        while left < right:
            total = nums[i] + nums[left] + nums[right]
            if total < 0:
                ____
            elif total > 0:
                ____
            else:
                triplets.append([nums[i], nums[left], nums[right]])
                left += 1
                right -= 1
                while left < right and ____:
                    left += 1
    return triplets`},
		},
		Solution: Solution{
			Code: `def threeSum(nums):
    nums.sort()
    triplets = []
    for i in range(len(nums) - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue  # same first value -> same triplets as last time
        left, right = i + 1, len(nums) - 1
        while left < right:
            total = nums[i] + nums[left] + nums[right]
            if total < 0:
                left += 1
            elif total > 0:
                right -= 1
            else:
                triplets.append([nums[i], nums[left], nums[right]])
                left += 1
                right -= 1
                while left < right and nums[left] == nums[left - 1]:
                    left += 1
    return triplets`,
			Explanation: "Sort, fix the first value, and run Two Sum II's opposite-end search on the rest for the pair that cancels it. Builds on Two Sum II and adds one idea: skipping equal neighbours so each distinct triplet is produced once, which sorting makes easy.",
			BruteForce:  "Try all triples i < j < k, sort each zero-sum triple and put it in a set to remove duplicates: n(n-1)(n-2)/6 triples, about 4.5 billion for n = 3000, so O(n^3). The waste: once a and b are fixed, the third value is determined (-(a + b)), yet the innermost loop scans for it; and the same triplet is built many times from duplicate values.",
			KeyInsight:  "Fixing a turns 3Sum into \"find pairs summing to -a in a sorted range\", which two pointers solve in O(n). Sorting also groups duplicates, so skipping an element equal to its predecessor removes duplicate triplets without a set.",
			WhyItWorks:  "Completeness: for a fixed i, the inner loop is Two Sum II on nums[i+1..], whose discard rule only removes values that can't reach -nums[i] (see Two Sum II), and after a match moving both pointers is safe because, with nums[left] fixed, only one value of nums[right] completes it. So every pair is found. No duplicates: we only use the first index of each run of equal values as the first element, and after a match we advance left past copies of the value just used, so the same (b, c) isn't emitted twice for the same a. Since triplets are built in sorted order a <= b <= c, two equal triplets would have to share all three values, which these skips rule out. Time: n outer iterations x O(n) inner = O(n^2).",
			CommonMistakes: []string{
				"Getting duplicate triplets for input like [-1,-1,0,1]. Skip a first value equal to the previous one (i > 0 and nums[i] == nums[i - 1]) and, after a match, skip left past equal values.",
				"Skipping when nums[i] == nums[i + 1] instead of nums[i - 1]. That skips the first copy, losing triplets such as [-1,-1,2] that need two copies of the same value.",
				"Forgetting to move the pointers after recording a match, which loops forever.",
			},
			TimeComplexity:  "O(n^2)",
			SpaceComplexity: "O(1) extra (besides sorting and the output)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort the input", Explanation: "Sorting enables the pointer moves and puts duplicate values side by side.", CodeSnippet: "nums.sort()", LineStart: 2, LineEnd: 2},
				{Title: "Skip duplicates", Explanation: "A first value equal to the previous one would rediscover the same triplets, so skip it. Comparing with i - 1 keeps the first copy.", CodeSnippet: "if i > 0 and nums[i] == nums[i - 1]:\n    continue  # same first value -> same triplets as last time", LineStart: 5, LineEnd: 6},
				{Title: "Place the pointers", Explanation: "Search the part to the right of i for a pair summing to -nums[i].", CodeSnippet: "left, right = i + 1, len(nums) - 1", LineStart: 7, LineEnd: 7},
				{Title: "Move the pointer that can't help", Explanation: "Too small: the left value can't reach zero even with the biggest partner. Too big: the right value is too large even with the smallest.", CodeSnippet: "if total < 0:\n    left += 1\nelif total > 0:\n    right -= 1", LineStart: 10, LineEnd: 13},
				{Title: "Record the answer", Explanation: "Save the triplet, move both pointers inward, then skip left past copies of the value just used.", CodeSnippet: "triplets.append([nums[i], nums[left], nums[right]])\nleft += 1\nright -= 1\nwhile left < right and nums[left] == nums[left - 1]:\n    left += 1", LineStart: 15, LineEnd: 19},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: all triples into a set",
					Code: `def threeSum(nums):
    found = set()
    n = len(nums)
    for i in range(n):
        for j in range(i + 1, n):
            for k in range(j + 1, n):
                if nums[i] + nums[j] + nums[k] == 0:
                    found.add(tuple(sorted((nums[i], nums[j], nums[k]))))
    return [list(t) for t in found]`,
					TimeComplexity:  "O(n^3)",
					SpaceComplexity: "O(n^3) in the worst case for the set",
					Explanation:     "Checks every triple and lets a set of sorted tuples remove duplicates. Easy to trust, far too slow for n = 3000.",
				},
				{
					Name: "Fix one value, hash set for the pair",
					Code: `def threeSum(nums):
    nums = sorted(nums)
    found = set()
    for i in range(len(nums)):
        seen = set()
        for j in range(i + 1, len(nums)):
            need = -nums[i] - nums[j]
            if need in seen:
                found.add((nums[i], need, nums[j]))
            seen.add(nums[j])
    return [list(t) for t in found]`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n)",
					Explanation:     "Uses Two Sum's seen-so-far set for the inner pair instead of two pointers. Same time, more memory, and a set is needed to remove duplicate triplets.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why compare nums[i] with nums[i - 1] rather than with nums[i + 1] when skipping?", Answer: "Comparing with i - 1 skips the later copies and keeps the first one, so the first copy still gets to search all positions to its right, including the other copies. Comparing with i + 1 would skip the first copy, and a triplet like [-1,-1,2] that needs both -1s would be lost."},
			{Prompt: "Why is it enough to skip duplicates only on the left pointer after a match?", Answer: "After a match with values (a, b, c), any later pair with the same b must also have the same c = -a - b, which would repeat the triplet. Skipping left past all copies of b prevents that. The right pointer then moves on its own as the sums change, so it can't produce the same triplet again without the same b."},
			{Prompt: "What if the target were some value t instead of 0?", Answer: "Nothing structural changes: compare nums[i] + nums[left] + nums[right] with t instead of 0. The pointer moves and duplicate skips work the same way."},
		},
	},
	{
		ID:              "container-with-most-water",
		Number:          11,
		Order:           4,
		Title:           "Container With Most Water",
		Slug:            "container-with-most-water",
		Difficulty:      "Medium",
		Category:        "two-pointers",
		Tags:            []string{"Array", "Two Pointers", "Greedy"},
		RelatedChapters: []int{3},
		Description: `You are given a list height of n non-negative integers. Picture n vertical lines on the x-axis: line i stands at x = i and is height[i] tall.

Choose two different lines. Together with the x-axis they form a container. The water it holds is a rectangle whose width is the distance between the two lines and whose height is the shorter of the two lines (water above the shorter line spills over). Lines between the two chosen ones don't get in the way.

Return the largest amount of water any container can hold. The container can't be tilted.`,
		Constraints: []string{
			"2 <= height.length <= 10^5",
			"0 <= height[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "height = [1,8,6,2,5,4,8,3,7]", Output: "49", Explanation: "Lines at index 1 (height 8) and index 8 (height 7): width 8 - 1 = 7, height min(8, 7) = 7, area 49."},
			{Input: "height = [1,1]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"height": []int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, Expected: 49, Note: "best pair is inside, not the two ends"},
			{Input: map[string]interface{}{"height": []int{1, 1}}, Expected: 1, Note: "smallest input"},
			{Input: map[string]interface{}{"height": []int{4, 3, 2, 1, 4}}, Expected: 16, Hidden: true, Note: "the two ends win"},
			{Input: map[string]interface{}{"height": []int{1, 2, 1}}, Expected: 2, Hidden: true, Note: "width beats height here"},
			{Input: map[string]interface{}{"height": []int{1, 2, 3, 4, 5, 6}}, Expected: 9, Hidden: true, Note: "strictly increasing"},
			{Input: map[string]interface{}{"height": []int{0, 0, 0}}, Expected: 0, Hidden: true, Note: "all zero heights hold nothing"},
			{Input: map[string]interface{}{"height": []int{2, 3, 10, 5, 7, 8, 9}}, Expected: 36, Hidden: true, Note: "tall lines in the middle"},
			{Input: map[string]interface{}{"height": []int{1, 3, 2, 5, 25, 24, 5}}, Expected: 24, Hidden: true, Note: "two tall neighbours beat wide short pairs"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def maxArea(height):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Start with the widest container: the two outer lines. Any other container is narrower. If you move the taller line inward, can the water level ever rise above the shorter line you kept?"},
			{Level: 2, Type: "insight", Content: "Pointers at both ends. The area is (right - left) * min(height[left], height[right]). The shorter line limits this container, and every narrower container that still uses it is limited by it too, so the shorter line can't be part of anything better. Discard it by moving that pointer inward. This is the \"move the pointer that can't help\" step from Chapter 3.", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Place the pointers: left = 0, right = last index.\n2. Record the answer: compute the current area and keep the best.\n3. Move the pointer that can't help: move whichever pointer is at the shorter line (either one on a tie).\n4. Stop when they meet and return the best area."},
			{Level: 4, Type: "skeleton", Content: `def maxArea(height):
    left, right = 0, len(height) - 1
    best = 0
    while left < right:
        width = ____
        best = max(best, ____)
        if ____:
            left += 1
        else:
            right -= 1
    return best`},
		},
		Solution: Solution{
			Code: `def maxArea(height):
    left, right = 0, len(height) - 1
    best = 0
    while left < right:
        width = right - left
        best = max(best, width * min(height[left], height[right]))
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1
    return best`,
			Explanation: "Start with the widest container and repeatedly discard the shorter of its two lines, recording the area each time. Builds on Two Sum II: the same opposite-ends walk where each step proves one end can't be part of a better answer, but the proof now comes from the area formula instead of sorted order.",
			BruteForce:  "Compute the area of every pair: n(n-1)/2 pairs, about 5 billion for n = 10^5, so O(n^2). The waste: once a line is the shorter side of a container, every narrower container using it is guaranteed to be smaller, yet the brute force still measures them all.",
			KeyInsight:  "The shorter line caps the height of every container that uses it. Pairing it with any line further in makes the container narrower without raising that cap, so the shorter line is done after one measurement.",
			WhyItWorks:  "Why moving the shorter line is safe: say height[left] <= height[right]. Any other container using left pairs it with some right' with left < right' < right. Its width is right' - left < right - left, and its height is min(height[left], height[right']) <= height[left] = the current height. So its area is strictly smaller than the area we just recorded, and line left can be discarded without missing the optimum. The same argument works for right when it is shorter. Each step discards one line and the best pair is measured before either of its lines is discarded, so the maximum is found in n - 1 steps.",
			CommonMistakes: []string{
				"Moving the taller line. That can discard a line that is part of the best container: with [1, 8, 6, 2, 5, 4, 8, 3, 7] moving the right pointer first throws away the 7 before it is paired with the 8 at index 1.",
				"Using max instead of min for the height. Water spills over the shorter line, so the height is min(height[left], height[right]).",
				"Using right - left + 1 for the width. Lines at x = 1 and x = 8 are 7 apart, not 8.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Place the pointers", Explanation: "Start with the widest possible container: the two outermost lines.", CodeSnippet: "left, right = 0, len(height) - 1", LineStart: 2, LineEnd: 2},
				{Title: "Record the answer", Explanation: "Measure the current container: width times the shorter line.", CodeSnippet: "width = right - left\nbest = max(best, width * min(height[left], height[right]))", LineStart: 5, LineEnd: 6},
				{Title: "Move the pointer that can't help", Explanation: "The shorter line can't be in a bigger container (narrower, and no taller), so drop it. For [1,8,6,...,7] the 1 goes first.", CodeSnippet: "if height[left] < height[right]:\n    left += 1\nelse:\n    right -= 1", LineStart: 7, LineEnd: 10},
				{Title: "Stop when they meet", Explanation: "Every line has been discarded only after it was proven useless, so best is the maximum.", CodeSnippet: "return best", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: every pair",
					Code: `def maxArea(height):
    best = 0
    for i in range(len(height)):
        for j in range(i + 1, len(height)):
            best = max(best, (j - i) * min(height[i], height[j]))
    return best`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Measures all n(n-1)/2 containers. Correct, but about 5 billion steps at the maximum size.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to move the shorter line, but not the taller one?", Answer: "Every narrower container with the shorter line is at most as tall as now (the short line still caps it) and narrower, so it holds less: the shorter line has nothing better to offer. The taller line, paired with some line further in that is taller than the current short one, could hold more, so discarding it might lose the answer."},
			{Prompt: "When the two lines are equal, which pointer should move, and why doesn't it matter?", Answer: "Either. If height[left] == height[right], any better container must be taller than both, so it can't use either line (each caps the height at the current value while being narrower). Both lines are useless, so discarding either first is safe."},
			{Prompt: "How is this different from Trapping Rain Water, which also has bars and water?", Answer: "Here you choose two lines and ignore everything between them; the answer is one rectangle. In Trapping Rain Water every bar is a wall and you add up the water standing on top of each bar, which depends on the tallest bar on each side of it."},
		},
	},
	{
		ID:              "trapping-rain-water",
		Number:          42,
		Order:           5,
		Title:           "Trapping Rain Water",
		Slug:            "trapping-rain-water",
		Difficulty:      "Hard",
		Category:        "two-pointers",
		Tags:            []string{"Array", "Two Pointers", "Dynamic Programming", "Stack"},
		RelatedChapters: []int{3, 2, 5},
		Description: `You are given a list height of n non-negative integers describing a landscape of bars, each 1 unit wide: bar i is height[i] units tall. Rain falls and fills every dip it can without spilling over the edges.

Water above bar i rises to the level of the lower of two walls: the tallest bar at or to the left of i, and the tallest bar at or to the right of i. So the water standing on bar i is min(tallest on the left, tallest on the right) - height[i], which is 0 when bar i is itself the tallest on one side. Water runs off both ends of the list.

Return the total units of water trapped.`,
		Constraints: []string{
			"1 <= height.length <= 2 * 10^4",
			"0 <= height[i] <= 10^5",
		},
		Examples: []Example{
			{Input: "height = [0,1,0,2,1,0,1,3,2,1,2,1]", Output: "6", Explanation: "Water stands 1 unit on index 2, 1 on index 4, 2 on index 5, 1 on index 6 and 1 on index 9: 6 in total."},
			{Input: "height = [4,2,0,3,2,5]", Output: "9", Explanation: "The walls are 4 and 5, so the level is 4: water 2 + 4 + 1 + 2 = 9 on the four inner bars."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"height": []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, Expected: 6, Note: "several pools of different depths"},
			{Input: map[string]interface{}{"height": []int{4, 2, 0, 3, 2, 5}}, Expected: 9, Note: "one wide pool bounded by 4 and 5"},
			{Input: map[string]interface{}{"height": []int{5}}, Expected: 0, Hidden: true, Note: "a single bar holds nothing"},
			{Input: map[string]interface{}{"height": []int{1, 2, 3, 4}}, Expected: 0, Hidden: true, Note: "rising bars: water runs off the left"},
			{Input: map[string]interface{}{"height": []int{4, 3, 2, 1}}, Expected: 0, Hidden: true, Note: "falling bars: water runs off the right"},
			{Input: map[string]interface{}{"height": []int{5, 0, 5}}, Expected: 5, Hidden: true, Note: "one deep valley"},
			{Input: map[string]interface{}{"height": []int{3, 0, 0, 2, 0, 4}}, Expected: 10, Hidden: true, Note: "level set by the shorter wall (3), not the taller (4)"},
			{Input: map[string]interface{}{"height": []int{2, 2, 2}}, Expected: 0, Hidden: true, Note: "flat ground"},
			{Input: map[string]interface{}{"height": []int{0, 7, 1, 4, 6}}, Expected: 7, Hidden: true, Note: "taller wall on the left inside the array"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def trap(height):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The water on bar i depends on the tallest bar on its left and the tallest bar on its right, but only on the smaller of the two. If you already know one side's tallest bar is the smaller one, do you need to know the other side's exact tallest bar?"},
			{Level: 2, Type: "insight", Content: "Keep pointers at both ends with running maxima left_max and right_max. Process whichever side has the lower current bar. If height[left] < height[right], the right side already has a bar taller than height[left], and in fact taller than left_max, so the level at left is set by left_max alone: add left_max - height[left] and move left. Symmetric for the right. Chapter 3's \"move the pointer that can't help\", combined with running maxima.", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Place the pointers: left = 0, right = last index; left_max = right_max = 0; water = 0.\n2. Compare and decide: which side has the lower bar?\n3. Lower on the left: update left_max, add left_max - height[left], move left right.\n4. Otherwise: update right_max, add right_max - height[right], move right left.\n5. Stop when they meet and return water."},
			{Level: 4, Type: "skeleton", Content: `def trap(height):
    left, right = 0, len(height) - 1
    left_max = right_max = 0
    water = 0
    while left < right:
        if height[left] < height[right]:
            left_max = ____
            water += ____
            left += 1
        else:
            right_max = ____
            water += ____
            right -= 1
    return water`},
		},
		Solution: Solution{
			Code: `def trap(height):
    left, right = 0, len(height) - 1
    left_max = right_max = 0
    water = 0
    while left < right:
        if height[left] < height[right]:
            left_max = max(left_max, height[left])
            water += left_max - height[left]
            left += 1
        else:
            right_max = max(right_max, height[right])
            water += right_max - height[right]
            right -= 1
    return water`,
			Explanation: "Walk in from both ends, always processing the side with the lower bar, and keep the tallest bar seen so far on each side. The water on the processed bar is that side's running maximum minus its height. Builds on Container With Most Water (move the side with the shorter bar) and on Product of Array Except Self (information from the left and from the right), compressing the two prefix tables into two running values.",
			BruteForce:  "For each bar, scan left for the tallest bar and scan right for the tallest bar, then add min(left, right) - height[i]: O(n) per bar, O(n^2) total, 400 million steps for n = 20,000. The waste: neighbouring bars rescan almost the same ranges. Storing prefix and suffix maxima in two arrays fixes the time (O(n)) but spends O(n) memory.",
			KeyInsight:  "You don't need both maxima exactly, only the smaller one. If the current left bar is lower than the current right bar, the right side is guaranteed to have a wall at least as tall as left_max, so left_max alone decides the water at left.",
			WhyItWorks:  "Why the smaller side's max is enough: claim that whenever we process left (height[left] < height[right]), the tallest bar at or right of position right is at least left_max. left_max is the height of some bar p we already passed with the left pointer; we passed it because it was lower than the right pointer's bar at that time. The right pointer only moves past a bar when the left bar is at least as tall, so if it had moved past every bar taller than height[p], then the left pointer would have been on a bar at least that tall, contradicting that height[p] is the left maximum. So a wall >= left_max remains on the right, and the true level at left, min(max left, max right), equals left_max. The same argument holds for the right side. Each step settles one bar, so the sum is exact after n - 1 steps.",
			CommonMistakes: []string{
				"Using max(left_max, right_max) as the water level. Water spills over the lower wall: for [3,0,0,2,0,4] the level over the dips is 3, not 4.",
				"Adding water before updating the running maximum, which gives negative water when the current bar is a new maximum. Update left_max first, then add left_max - height[left] (0 for a new maximum).",
				"Moving the pointer on the taller side. Then the other side's maximum is not guaranteed to be the limit, and water can be over- or under-counted.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Place the pointers", Explanation: "Pointers at both ends, a running maximum for each side, and the total.", CodeSnippet: "left, right = 0, len(height) - 1\nleft_max = right_max = 0\nwater = 0", LineStart: 2, LineEnd: 4},
				{Title: "Compare and decide", Explanation: "The lower current bar tells us which side's maximum is the limiting wall.", CodeSnippet: "if height[left] < height[right]:", LineStart: 6, LineEnd: 6},
				{Title: "Record the answer", Explanation: "On the lower side the level is that side's running maximum. Update it first so a new maximum adds 0.", CodeSnippet: "left_max = max(left_max, height[left])\nwater += left_max - height[left]\nleft += 1", LineStart: 7, LineEnd: 9},
				{Title: "Move the pointer that can't help", Explanation: "The mirror case for the right side when its bar is lower or equal.", CodeSnippet: "right_max = max(right_max, height[right])\nwater += right_max - height[right]\nright -= 1", LineStart: 11, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: scan both sides for every bar",
					Code: `def trap(height):
    total = 0
    for i in range(len(height)):
        tallest_left = max(height[:i + 1])
        tallest_right = max(height[i:])
        total += min(tallest_left, tallest_right) - height[i]
    return total`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Applies the definition directly. Easy to trust; each bar rescans the whole list.",
				},
				{
					Name: "Prefix and suffix maxima",
					Code: `def trap(height):
    n = len(height)
    left_max = [0] * n
    right_max = [0] * n
    for i in range(n):
        left_max[i] = max(height[i], left_max[i - 1] if i > 0 else 0)
    for i in range(n - 1, -1, -1):
        right_max[i] = max(height[i], right_max[i + 1] if i < n - 1 else 0)
    return sum(min(left_max[i], right_max[i]) - height[i] for i in range(n))`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Precompute the tallest bar to the left and to the right of every position, as Product of Array Except Self does with products. O(n) time, and the easiest O(n) version to get right; the two-pointer solution removes the arrays.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the smaller side's running maximum enough to know the water on that bar?", Answer: "The water level is min(tallest left, tallest right). When height[left] < height[right], the right side already contains a bar at least as tall as left_max (it is what kept the left pointer moving), so the minimum is left_max. We don't need the right side's exact maximum, only the fact that it is at least left_max."},
			{Prompt: "What if you processed the side with the taller bar instead?", Answer: "Then the other side's running maximum might be smaller than the true limiting wall still to come, or larger than it, so left_max - height[left] could be wrong in either direction. Processing the lower side is what guarantees the opposite side has a wall tall enough."},
			{Prompt: "How does this problem combine ideas from earlier in the course?", Answer: "From Product of Array Except Self: the answer at i depends on everything to its left and to its right, which can be precomputed as prefix and suffix values. From Container With Most Water: move the pointer at the shorter bar because it is the one whose answer is already determined. Together they turn the O(n)-memory prefix solution into O(1) memory."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, TwoPointersProblems...)
}
