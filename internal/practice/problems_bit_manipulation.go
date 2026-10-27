package practice

// BitManipulationProblems contains all bit-manipulation category problems, in
// ladder order: XOR cancellation (Single Number), clearing the lowest set bit
// (Number of 1 Bits), reusing smaller answers (Counting Bits), XOR against an
// index range (Missing Number), building a number bit by bit (Reverse Bits)
// and finally fixed-width arithmetic with masks (Sum of Two Integers).
// Chapter 20 is the primer they rely on.
var BitManipulationProblems = []*Problem{
	{
		ID:              "single-number",
		Number:          136,
		Order:           1,
		Title:           "Single Number",
		Slug:            "single-number",
		Difficulty:      "Easy",
		Category:        "bit-manipulation",
		Tags:            []string{"Array", "Bit Manipulation"},
		RelatedChapters: []int{20, 2},
		Description: `You are given a non-empty list of integers nums. Every value appears exactly twice, except for one value that appears exactly once. Return that single value.

Your solution should run in linear time and use only a constant amount of extra memory, so a hash map of counts does not meet the target (although it gives the right answer).`,
		Constraints: []string{
			"1 <= nums.length <= 3 * 10^4",
			"-3 * 10^4 <= nums[i] <= 3 * 10^4",
			"Every element appears exactly twice except for one element, which appears once",
		},
		Examples: []Example{
			{ID: "1", Input: "nums = [2,2,1]", Output: "1", Explanation: "2 ^ 2 ^ 1 = 0 ^ 1 = 1: the two 2s cancel."},
			{ID: "2", Input: "nums = [4,1,2,1,2]", Output: "4"},
			{ID: "3", Input: "nums = [1]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 2, 1}}, Expected: 1, Note: "single value at the end"},
			{Input: map[string]interface{}{"nums": []int{4, 1, 2, 1, 2}}, Expected: 4, Note: "single value first, pairs not adjacent"},
			{Input: map[string]interface{}{"nums": []int{1}}, Expected: 1, Note: "one element", Hidden: true},
			{Input: map[string]interface{}{"nums": []int{-1, -1, -2}}, Expected: -2, Note: "negative numbers", Hidden: true},
			{Input: map[string]interface{}{"nums": []int{7, 3, 0, 3, 7}}, Expected: 0, Note: "the single value is 0, so a result of 0 is not 'not found'", Hidden: true},
			{Input: map[string]interface{}{"nums": []int{30000, -30000, 30000}}, Expected: -30000, Note: "constraint extremes, positive and negative", Hidden: true},
			{Input: map[string]interface{}{"nums": []int{5, 9, 9, 6, 6, 5, 8}}, Expected: 8, Note: "several pairs in mixed order", Hidden: true},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def singleNumber(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Is there an operation that combines two equal values into 'nothing', leaves a value unchanged when combined with 'nothing', and does not care about the order you apply it in?"},
			{Level: 2, Type: "insight", Content: "XOR cancellation (Chapter 20). x ^ x = 0, x ^ 0 = x, and XOR is commutative and associative, so XOR-ing every number lets each pair cancel no matter where its two copies are. What remains is the single value.", ChapterRef: 20},
			{Level: 3, Type: "plan", Content: "1. Start from zero: result = 0, which XOR leaves unchanged.\n2. Fold with XOR: for each number, result ^= number.\n3. Read off the survivor: every pair has cancelled, so result is the single value."},
			{Level: 4, Type: "skeleton", Content: "def singleNumber(nums):\n    result = ____\n    for num in nums:\n        ____\n    return result"},
		},
		Solution: Solution{
			Code: `def singleNumber(nums):
    result = 0
    for num in nums:
        result ^= num
    return result`,
			Explanation: "The first rung uses the plainest bit trick from Chapter 20: XOR cancellation. XOR-ing all values makes every pair vanish because x ^ x = 0, in any order, so the running XOR ends as the one value without a partner.",
			BruteForce:  "Count occurrences in a hash map (Chapter 2) and return the key with count 1. It is O(n) time but O(n) extra memory, about 15,000 entries for the largest input. Sorting and scanning neighbours avoids the map but costs O(n log n).",
			KeyInsight:  "XOR works bit by bit as 'addition without carry': each bit position of the result is 1 exactly when an odd number of inputs have a 1 there. Paired values contribute an even number of 1s to every position, so only the single value's bits survive.",
			WhyItWorks:  "Look at one bit position. Each value that appears twice puts the same bit there twice, contributing 0 or 2 ones, which is even. The single value contributes 0 or 1. XOR of a column of bits is 1 exactly when the number of 1s is odd, so the result's bit equals the single value's bit in every position, which means result equals the single value. Order never matters because XOR is commutative and associative, so the two copies do not need to be next to each other. In Python this also holds for negative numbers, which behave as if they had infinitely many leading 1 bits: those leading bits are paired too.",
			CommonMistakes: []string{
				"Starting result at nums[0] and then also XOR-ing nums[0] in the loop, which cancels it. Fix: start from 0, the value that XOR leaves unchanged.",
				"Using a dictionary or set, which works but uses O(n) memory and misses the constant-space requirement. Fix: fold the list with XOR.",
				"Using OR or AND instead of XOR. OR never removes a bit (1 | 1 = 1), so pairs do not cancel. Fix: use ^.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Start from zero", Explanation: "0 is the identity for XOR (x ^ 0 = x), so it does not disturb the fold.", CodeSnippet: "result = 0", LineStart: 2, LineEnd: 2},
				{Title: "Fold with XOR", Explanation: "Each number is XOR-ed in. For [4,1,2,1,2]: 0^4 = 4, 4^1 = 5, 5^2 = 7, 7^1 = 6, 6^2 = 4.", CodeSnippet: "for num in nums:\n    result ^= num", LineStart: 3, LineEnd: 4},
				{Title: "Read off the survivor", Explanation: "All pairs have cancelled to 0, so what is left is the single value.", CodeSnippet: "return result", LineStart: 5, LineEnd: 5},
			},
			Alternatives: []AltSolution{
				{
					Name: "Count with a dictionary",
					Code: `def singleNumber(nums):
    counts = {}
    for num in nums:
        counts[num] = counts.get(num, 0) + 1
    for num, count in counts.items():
        if count == 1:
            return num`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The Chapter 2 approach: count everything, then find the value with count 1. Correct and general (it also works if values could appear three times), but it uses memory proportional to the input.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does XOR still work when the two copies of a value are far apart in the list, like in [4,1,2,1,2]?", Answer: "XOR is commutative and associative, so the order of the operations can be rearranged freely: 4^1^2^1^2 = 4^(1^1)^(2^2) = 4^0^0 = 4. The running value in between looks random, but the final value only depends on how many times each bit was set, not on the order."},
			{Prompt: "What if every value appeared three times except one? Does XOR still work?", Answer: "No: x ^ x ^ x = x, so triples do not cancel. Instead count, for each of the 32 bit positions, how many numbers have a 1 there and take that count modulo 3; the positions with remainder 1 are the bits of the single value (with care for the sign bit in Python). This is LeetCode 137, Single Number II."},
		},
	},
	{
		ID:              "number-of-1-bits",
		Number:          191,
		Order:           2,
		Title:           "Number of 1 Bits",
		Slug:            "number-of-1-bits",
		Difficulty:      "Easy",
		Category:        "bit-manipulation",
		Tags:            []string{"Divide and Conquer", "Bit Manipulation"},
		RelatedChapters: []int{20},
		Description: `You are given a positive integer n. Return how many 1s appear in its binary representation. This count is also called the Hamming weight or population count.

For example, 11 is 1011 in binary, which has three 1s, so the answer is 3.

The input is an ordinary integer, not a string of 0s and 1s.`,
		Constraints: []string{
			"1 <= n <= 2^31 - 1",
		},
		Examples: []Example{
			{ID: "1", Input: "n = 11", Output: "3", Explanation: "11 = 1011 in binary: three bits are 1."},
			{ID: "2", Input: "n = 128", Output: "1", Explanation: "128 = 10000000: a power of two has exactly one 1."},
			{ID: "3", Input: "n = 2147483645", Output: "30", Explanation: "2147483645 = 1111111111111111111111111111101 (31 bits, one of them 0)."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 11}, Expected: 3, Note: "small number, three bits set"},
			{Input: map[string]interface{}{"n": 128}, Expected: 1, Note: "power of two: one bit set"},
			{Input: map[string]interface{}{"n": 2147483645}, Expected: 30, Note: "almost all bits set", Hidden: true},
			{Input: map[string]interface{}{"n": 1}, Expected: 1, Note: "smallest input", Hidden: true},
			{Input: map[string]interface{}{"n": 2147483647}, Expected: 31, Note: "largest input: all 31 low bits set", Hidden: true},
			{Input: map[string]interface{}{"n": 1073741824}, Expected: 1, Note: "only the highest allowed bit (2^30) is set: catches checking too few positions", Hidden: true},
			{Input: map[string]interface{}{"n": 255}, Expected: 8, Note: "a full byte", Hidden: true},
			{Input: map[string]interface{}{"n": 43261596}, Expected: 12, Note: "scattered bits", Hidden: true},
		},
		TimeComplexity:  "O(k), where k is the number of 1 bits (at most 31)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def hammingWeight(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Write 12 and 11 in binary, one above the other. What does subtracting 1 do to the lowest 1 bit of 12 and to the bits below it, and what happens if you AND the two numbers?"},
			{Level: 2, Type: "insight", Content: "n & (n - 1) clears the lowest set bit of n (Chapter 20). Subtracting 1 turns the lowest 1 into 0 and the 0s below it into 1s, leaving the higher bits alone; AND-ing with n keeps the higher bits and zeroes everything from the lowest 1 down. So each application removes exactly one 1, and the number of applications until n is 0 is the answer.", ChapterRef: 20},
			{Level: 3, Type: "plan", Content: "1. Start a counter at 0.\n2. Loop while bits remain: while n is not 0.\n3. Peel off one bit: n = n & (n - 1) removes the lowest 1.\n4. Count the peel: add 1 to the counter each time, and return it when n reaches 0."},
			{Level: 4, Type: "skeleton", Content: "def hammingWeight(n):\n    count = 0\n    while ____:\n        n = ____\n        count += 1\n    return count"},
		},
		Solution: Solution{
			Code: `def hammingWeight(n):
    count = 0
    while n:
        n &= n - 1
        count += 1
    return count`,
			Explanation: "Single Number used XOR on whole values; this rung looks inside one value. The trick n & (n - 1) from Chapter 20 removes exactly the lowest 1 bit, so counting how many removals it takes to reach 0 counts the 1 bits, and the loop runs once per 1 bit rather than once per bit position.",
			BruteForce:  "Check all 32 positions: count += (n >> i) & 1 for i in 0..31. That is also constant time for 32-bit inputs, but it always does 32 rounds, even for 128 which has a single 1 bit. Converting with bin(n).count('1') works too but hides the bit reasoning this chapter is about.",
			KeyInsight:  "n - 1 flips the lowest 1 bit of n to 0 and every 0 below it to 1, and changes nothing above it. So n & (n - 1) equals n with only its lowest 1 bit cleared. Example: 12 = 1100, 11 = 1011, 12 & 11 = 1000 = 8.",
			WhyItWorks:  "Write n as (higher bits) 1 followed by t zeros, where that 1 is the lowest set bit. Subtracting 1 has to borrow from that 1: the result is (same higher bits) 0 followed by t ones. AND-ing the two: the higher bits are equal, so they stay; the lowest-set position is 1 & 0 = 0; the t positions below are 0 & 1 = 0. So exactly one 1 disappears per round and no new 1 appears. Since n > 0 starts with k ones, after exactly k rounds n == 0 and the loop stops with count == k. For n = 11 (1011): 1011 -> 1010 -> 1000 -> 0000, three rounds.",
			CommonMistakes: []string{
				"Using n - 1 or n >> 1 alone, which changes many bits or loses track of which bits were 1. Fix: n & (n - 1) removes exactly one 1 bit.",
				"Checking n % 10 == 1 or treating the input as a string of digits. The input is an integer and its decimal digits have nothing to do with its bits. Fix: use n & 1 or n & (n - 1).",
				"Looping over a fixed small number of positions (say 16), which misses high bits like 2^30. Fix: loop until n becomes 0, or check all 32 positions.",
			},
			TimeComplexity:  "O(k), where k is the number of 1 bits (at most 31)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Loop while bits remain", Explanation: "A non-zero integer still has at least one 1 bit, so the loop continues exactly as long as there is something to count.", CodeSnippet: "while n:", LineStart: 3, LineEnd: 3},
				{Title: "Peel off one bit", Explanation: "n & (n - 1) clears the lowest 1. For 11: 1011 & 1010 = 1010, then 1010 & 1001 = 1000, then 1000 & 0111 = 0.", CodeSnippet: "n &= n - 1", LineStart: 4, LineEnd: 4},
				{Title: "Count the peel", Explanation: "Each round removed exactly one 1, so the number of rounds is the number of 1 bits.", CodeSnippet: "count += 1\nreturn count", LineStart: 5, LineEnd: 6},
			},
			Alternatives: []AltSolution{
				{
					Name: "Check every bit position",
					Code: `def hammingWeight(n):
    count = 0
    for i in range(32):
        count += (n >> i) & 1
    return count`,
					TimeComplexity:  "O(1) (always 32 rounds)",
					SpaceComplexity: "O(1)",
					Explanation:     "Shift bit i down to position 0 and test it with & 1. Simple and always 32 rounds; the n & (n - 1) version is faster when few bits are set.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Trace n & (n - 1) on n = 40 (101000). Which bit disappears and why do the bits below it not turn into 1s?", Answer: "40 - 1 = 39 = 100111. The lowest 1 of 40 (value 8) became 0 and the three 0s below it became 1s. AND-ing 101000 with 100111 gives 100000 = 32: the high bit survives because both have it, the 8-bit is 1 & 0 = 0, and the low three bits are 0 & 1 = 0. So only the lowest 1 vanished."},
			{Prompt: "How can n & (n - 1) tell you in one step whether a positive n is a power of two?", Answer: "A power of two has exactly one 1 bit. Clearing it leaves 0, so n > 0 and n & (n - 1) == 0 is true exactly for powers of two. Any other positive number still has a 1 bit left after the lowest one is cleared."},
		},
	},
	{
		ID:              "counting-bits",
		Number:          338,
		Order:           3,
		Title:           "Counting Bits",
		Slug:            "counting-bits",
		Difficulty:      "Easy",
		Category:        "bit-manipulation",
		Tags:            []string{"Dynamic Programming", "Bit Manipulation"},
		RelatedChapters: []int{20, 15},
		Description: `Given a non-negative integer n, return a list ans of length n + 1 where ans[i] is the number of 1 bits in the binary representation of i, for every i from 0 to n.

For example, for n = 5 the numbers 0..5 are 0, 1, 10, 11, 100, 101 in binary, so the answer is [0, 1, 1, 2, 1, 2].

Counting the bits of each number separately takes O(n log n) time. Can you do it in O(n) by reusing answers you have already computed?`,
		Constraints: []string{
			"0 <= n <= 10^5",
		},
		Examples: []Example{
			{ID: "1", Input: "n = 2", Output: "[0,1,1]", Explanation: "0 = 0, 1 = 1, 2 = 10."},
			{ID: "2", Input: "n = 5", Output: "[0,1,1,2,1,2]", Explanation: "5 = 101 is 2 = 10 with a 1 appended, so ans[5] = ans[2] + 1 = 2."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 2}, Expected: []int{0, 1, 1}, Note: "small n"},
			{Input: map[string]interface{}{"n": 5}, Expected: []int{0, 1, 1, 2, 1, 2}, Note: "mix of even and odd"},
			{Input: map[string]interface{}{"n": 0}, Expected: []int{0}, Note: "n = 0: list with one entry", Hidden: true},
			{Input: map[string]interface{}{"n": 1}, Expected: []int{0, 1}, Note: "n = 1", Hidden: true},
			{Input: map[string]interface{}{"n": 8}, Expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1}, Note: "ends on a power of two", Hidden: true},
			{Input: map[string]interface{}{"n": 16}, Expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1}, Note: "15 = 1111 then 16 = 10000", Hidden: true},
			{Input: map[string]interface{}{"n": 20}, Expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2}, Note: "past a power of two", Hidden: true},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n) for the output",
		StarterCode:     "def countBits(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Compare the binary forms of 5 (101) and 2 (10), or of 6 (110) and 3 (11). How is the bigger number built from the smaller one, and what does that mean for their counts of 1s?"},
			{Level: 2, Type: "insight", Content: "Peel off one bit and reuse the smaller answer (Chapter 20, plus the reuse idea from Chapter 15). i >> 1 is i with its lowest bit removed, it is smaller than i so its count is already known, and i & 1 is the removed bit. So ones[i] = ones[i >> 1] + (i & 1). Using Number of 1 Bits' trick instead gives ones[i] = ones[i & (i - 1)] + 1.", ChapterRef: 20},
			{Level: 3, Type: "plan", Content: "1. Set the base case: a list of n + 1 zeros, since 0 has no 1 bits.\n2. Peel off one bit: for i from 1 to n, split i into i >> 1 (already solved) and its lowest bit i & 1.\n3. Reuse the smaller answer: ones[i] = ones[i >> 1] + (i & 1).\n4. Return the list."},
			{Level: 4, Type: "skeleton", Content: "def countBits(n):\n    ones = [0] * (n + 1)\n    for i in range(1, n + 1):\n        ones[i] = ones[____] + ____\n    return ones"},
		},
		Solution: Solution{
			Code: `def countBits(n):
    ones = [0] * (n + 1)
    for i in range(1, n + 1):
        ones[i] = ones[i >> 1] + (i & 1)
    return ones`,
			Explanation: "Builds on Number of 1 Bits: instead of counting each number from scratch, notice that removing one bit from i gives a smaller number whose count is already in the table. Shifting right removes the lowest bit (i >> 1) and i & 1 says whether that bit was a 1, so each entry costs O(1). This is the first time a bit trick feeds a dynamic-programming table (Chapter 15).",
			BruteForce:  "Run the Number of 1 Bits loop for every i from 0 to n. Each number up to 10^5 has at most 17 bits, so it is O(n log n), about 1.7 million steps for the largest n, and it recounts bits it has already counted: the count for 10 (1010) repeats the work for 5 (101).",
			KeyInsight:  "i and i >> 1 have the same bits except the lowest one of i, which i >> 1 drops. Because i >> 1 < i for every i >= 1, its count is already in the table when i is reached.",
			WhyItWorks:  "Subproblem: ones[i] is the number of 1 bits in i. Recurrence: shifting i right by one drops exactly its lowest bit and keeps every other bit (moved one place down, which does not change how many there are), so ones[i] = ones[i >> 1] + (i & 1). Base case: ones[0] = 0. Order: i >> 1 < i for all i >= 1, so filling the table in increasing i always reads an entry that is already final. The table covers every i from 0 to n, so it is complete. Example: ones[6] = ones[3] + 0 = 2 because 110 is 11 shifted left; ones[7] = ones[3] + 1 = 3.",
			CommonMistakes: []string{
				"Writing ones[i >> 1] + 1 for every i, which overcounts even numbers: ones[2] would be 2 instead of 1. Fix: add the dropped bit, i & 1, which is 0 for even i.",
				"Writing ones[i - 1] + 1, which assumes adding 1 always adds a 1 bit; 3 (11) + 1 = 4 (100) loses a bit. Fix: relate i to i >> 1 or to i & (i - 1).",
				"Precedence slips in other languages: in C or Java, ones[i >> 1] + i & 1 parses as (ones[i >> 1] + i) & 1. Fix: always parenthesize (i & 1).",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n) for the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set the base case", Explanation: "0 has no 1 bits, and starting every entry at 0 gives ones[0] for free.", CodeSnippet: "ones = [0] * (n + 1)", LineStart: 2, LineEnd: 2},
				{Title: "Peel off one bit", Explanation: "i >> 1 is i without its lowest bit and i & 1 is that bit. For i = 5 (101): 5 >> 1 = 2 (10) and 5 & 1 = 1.", CodeSnippet: "for i in range(1, n + 1):", LineStart: 3, LineEnd: 3},
				{Title: "Reuse the smaller answer", Explanation: "ones[i >> 1] was filled in earlier because i >> 1 < i, so ones[5] = ones[2] + 1 = 2.", CodeSnippet: "ones[i] = ones[i >> 1] + (i & 1)", LineStart: 4, LineEnd: 4},
				{Title: "Return the table", Explanation: "Every entry from 0 to n is final.", CodeSnippet: "return ones", LineStart: 5, LineEnd: 5},
			},
			Alternatives: []AltSolution{
				{
					Name: "Reuse via the lowest set bit",
					Code: `def countBits(n):
    ones = [0] * (n + 1)
    for i in range(1, n + 1):
        ones[i] = ones[i & (i - 1)] + 1
    return ones`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n) for the output",
					Explanation:     "i & (i - 1) is i with its lowest 1 bit cleared, the trick from Number of 1 Bits. It is smaller than i and has exactly one fewer 1 bit, so ones[i] = ones[i & (i - 1)] + 1.",
				},
				{
					Name: "Count each number separately",
					Code: `def countBits(n):
    def count_ones(x):
        count = 0
        while x:
            x &= x - 1
            count += 1
        return count
    return [count_ones(i) for i in range(n + 1)]`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n) for the output",
					Explanation:     "Calls the Number of 1 Bits loop on every i. Correct, but it does not reuse anything between numbers.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to read ones[i >> 1] while filling ones[i]?", Answer: "Because i >> 1 is i // 2, which is strictly less than i for every i >= 1, and the loop fills entries in increasing order. So ones[i >> 1] was already computed (or is ones[0] = 0) and is final when we read it."},
			{Prompt: "Both recurrences ones[i >> 1] + (i & 1) and ones[i & (i - 1)] + 1 are correct. What different 'smaller number' does each one use?", Answer: "The first removes the lowest bit position (whether it is 0 or 1) by shifting, so the smaller number is about half of i and the bit removed may or may not be a 1. The second removes the lowest 1 bit wherever it is, so the smaller number is i minus its lowest power-of-two part and always has exactly one fewer 1."},
		},
	},
	{
		ID:              "missing-number",
		Number:          268,
		Order:           4,
		Title:           "Missing Number",
		Slug:            "missing-number",
		Difficulty:      "Easy",
		Category:        "bit-manipulation",
		Tags:            []string{"Array", "Hash Table", "Math", "Bit Manipulation"},
		RelatedChapters: []int{20, 2},
		Description: `You are given a list nums of n distinct integers, all taken from the range 0 to n inclusive. That range has n + 1 values, so exactly one of them is missing from the list. Return the missing value.

Aim for O(n) time and O(1) extra memory.`,
		Constraints: []string{
			"n == nums.length",
			"1 <= n <= 10^4",
			"0 <= nums[i] <= n",
			"All the numbers in nums are distinct",
		},
		Examples: []Example{
			{ID: "1", Input: "nums = [3,0,1]", Output: "2", Explanation: "n = 3, so the range is 0..3. 0, 1 and 3 are present; 2 is missing."},
			{ID: "2", Input: "nums = [0,1]", Output: "2", Explanation: "n = 2, so the range is 0..2 and the missing value is n itself."},
			{ID: "3", Input: "nums = [9,6,4,2,3,5,7,0,1]", Output: "8"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{3, 0, 1}}, Expected: 2, Note: "missing value in the middle"},
			{Input: map[string]interface{}{"nums": []int{0, 1}}, Expected: 2, Note: "missing value is n itself"},
			{Input: map[string]interface{}{"nums": []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, Expected: 8, Note: "unsorted, longer list"},
			{Input: map[string]interface{}{"nums": []int{0}}, Expected: 1, Note: "n = 1, missing 1", Hidden: true},
			{Input: map[string]interface{}{"nums": []int{1}}, Expected: 0, Note: "n = 1, missing 0", Hidden: true},
			{Input: map[string]interface{}{"nums": []int{1, 2}}, Expected: 0, Note: "missing 0: catches starting the XOR without n", Hidden: true},
			{Input: map[string]interface{}{"nums": []int{0, 2}}, Expected: 1, Note: "missing value between present ones", Hidden: true},
			{Input: map[string]interface{}{"nums": []int{5, 4, 3, 2, 1, 0}}, Expected: 6, Note: "reverse order, missing the top value", Hidden: true},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def missingNumber(nums):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In Single Number every value had a partner except one. Can you give every value in nums a partner from somewhere else, so that only the missing value is left without one?"},
			{Level: 2, Type: "insight", Content: "XOR cancellation again (Chapter 20). XOR together all the indices 0..n and all the values in nums. Every present value appears twice (once as an index or as n, once as a value) and cancels; the missing value appears only once, as an index, so it is what remains.", ChapterRef: 20},
			{Level: 3, Type: "plan", Content: "1. Start with n: the index range 0..n-1 does not include n, so put n into the running XOR first.\n2. Fold with XOR: for each position i, XOR in both i and nums[i].\n3. Read off the survivor: every value except the missing one appeared twice and cancelled."},
			{Level: 4, Type: "skeleton", Content: "def missingNumber(nums):\n    missing = ____\n    for i, num in enumerate(nums):\n        missing ^= ____\n    return missing"},
		},
		Solution: Solution{
			Code: `def missingNumber(nums):
    missing = len(nums)
    for i, num in enumerate(nums):
        missing ^= i ^ num
    return missing`,
			Explanation: "Builds on Single Number. There the pairs were already in the list; here you create them by XOR-ing in every index 0..n-1 together with n. Each value that is present meets its twin and cancels, and the missing value is the only one left unpaired.",
			BruteForce:  "Put nums in a set (Chapter 2) and test 0, 1, ..., n until one is missing: O(n) time but O(n) extra memory. Sorting and looking for the first i with nums[i] != i needs no set but costs O(n log n).",
			KeyInsight:  "XOR of the full range 0..n with all values in nums leaves exactly the missing value, because XOR cancels every value that appears an even number of times and the full range lists each present value once more.",
			WhyItWorks:  "The running value is the XOR of the multiset {0, 1, ..., n} together with the multiset nums. Every value v that is in nums appears in that combined collection exactly twice (once in the range, once in nums), and v ^ v = 0. The missing value m appears exactly once, in the range. XOR is commutative and associative, so the order of the terms does not matter and the result is m ^ 0 ^ 0 ^ ... = m. Starting from n is what puts the last value of the range into the XOR, because enumerate only produces the indices 0..n-1. Example [3,0,1]: 3 ^ (0^3) ^ (1^0) ^ (2^1) = 2.",
			CommonMistakes: []string{
				"Starting the XOR at 0 instead of n, so the value n is never paired. [1, 2] then returns 2 (missing ^ n) instead of 0. Fix: initialise with len(nums).",
				"Assuming the list is sorted and returning the first index where nums[i] != i, which fails for [3,0,1]. Fix: XOR (or sum) does not care about order.",
				"Returning n when no gap is found in a sorted scan but forgetting that 0 can be the missing value. Fix: XOR covers every case in one pass.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Start with n", Explanation: "The indices only cover 0..n-1, so the top value n of the range is folded in at the start.", CodeSnippet: "missing = len(nums)", LineStart: 2, LineEnd: 2},
				{Title: "Fold with XOR", Explanation: "XOR in each index and each value. For [3,0,1]: 3 ^ 0 ^ 3 = 0, then ^ 1 ^ 0 = 1, then ^ 2 ^ 1 = 2.", CodeSnippet: "for i, num in enumerate(nums):\n    missing ^= i ^ num", LineStart: 3, LineEnd: 4},
				{Title: "Read off the survivor", Explanation: "Each present value appeared once as an index (or as n) and once as a value, so only the missing value survives.", CodeSnippet: "return missing", LineStart: 5, LineEnd: 5},
			},
			Alternatives: []AltSolution{
				{
					Name: "Sum formula",
					Code: `def missingNumber(nums):
    n = len(nums)
    return n * (n + 1) // 2 - sum(nums)`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "0 + 1 + ... + n = n(n + 1)/2, and subtracting the actual sum leaves the missing value. Equally fast and very readable. In languages with fixed-size integers the sum can overflow for large n, which XOR never does; in Python it is fine.",
				},
				{
					Name: "Set lookup",
					Code: `def missingNumber(nums):
    present = set(nums)
    for candidate in range(len(nums) + 1):
        if candidate not in present:
            return candidate`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The straightforward Chapter 2 approach: store what is present and check each candidate. Correct, but uses O(n) memory.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the XOR start at len(nums) instead of 0?", Answer: "The pairing needs every value of the range 0..n to appear once on the 'range' side. enumerate gives only the indices 0..n-1, so n would be missing from that side. Starting at n adds it. Without it, n would be unpaired when present, and the answer would come out as missing ^ n."},
			{Prompt: "What if the list could contain one duplicate as well as one missing value (values 1..n, one repeated, one missing)? Is a single XOR still enough?", Answer: "No. XOR of the range and the list then gives missing ^ duplicate, a mix of two unknowns. You need a second fact: split the values by a bit where missing and duplicate differ (any 1 bit of that XOR) and XOR each group separately, or use the sum and the sum of squares to set up two equations."},
		},
	},
	{
		ID:              "reverse-bits",
		Number:          190,
		Order:           5,
		Title:           "Reverse Bits",
		Slug:            "reverse-bits",
		Difficulty:      "Easy",
		Category:        "bit-manipulation",
		Tags:            []string{"Divide and Conquer", "Bit Manipulation"},
		RelatedChapters: []int{20},
		Description: `You are given a 32-bit unsigned integer n. Write it as exactly 32 binary digits (with leading zeros), reverse the order of those 32 digits, and return the resulting 32-bit unsigned integer.

For example, 1 is 00000000000000000000000000000001; reversed it is 10000000000000000000000000000000, which is 2147483648.

The input is an ordinary integer between 0 and 2^32 - 1, not a string.`,
		Constraints: []string{
			"0 <= n <= 2^32 - 1",
		},
		Examples: []Example{
			{ID: "1", Input: "n = 43261596", Output: "964176192", Explanation: "43261596 = 00000010100101000001111010011100 and 964176192 = 00111001011110000010100101000000, the same 32 digits in reverse order."},
			{ID: "2", Input: "n = 4294967293", Output: "3221225471", Explanation: "11111111111111111111111111111101 reversed is 10111111111111111111111111111111."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 43261596}, Expected: 964176192, Note: "the first example"},
			{Input: map[string]interface{}{"n": 4294967293}, Expected: 3221225471, Note: "high bits set"},
			{Input: map[string]interface{}{"n": 0}, Expected: 0, Note: "zero stays zero", Hidden: true},
			{Input: map[string]interface{}{"n": 1}, Expected: 2147483648, Note: "lowest bit moves to the top: catches stopping when n becomes 0", Hidden: true},
			{Input: map[string]interface{}{"n": 2147483648}, Expected: 1, Note: "top bit moves to the bottom", Hidden: true},
			{Input: map[string]interface{}{"n": 4294967295}, Expected: 4294967295, Note: "all 32 bits set: a palindrome", Hidden: true},
			{Input: map[string]interface{}{"n": 6}, Expected: 1610612736, Note: "small number: leading zeros become trailing zeros", Hidden: true},
			{Input: map[string]interface{}{"n": 2863311530}, Expected: 1431655765, Note: "alternating 1010...10 becomes 0101...01", Hidden: true},
		},
		TimeComplexity:  "O(1) (always 32 rounds)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def reverseBits(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Reversing the decimal digits of 123 can be done by repeatedly taking the last digit of the input and appending it to the end of the output. What are 'take the last digit' and 'append a digit' in binary?"},
			{Level: 2, Type: "insight", Content: "Peel off one bit at a time (Chapter 20): n & 1 is the lowest bit and n >>= 1 drops it. Append a bit to the result with result = (result << 1) | bit. Do it exactly 32 times, because the leading zeros of n become trailing zeros of the answer and must still be shifted in.", ChapterRef: 20},
			{Level: 3, Type: "plan", Content: "1. Start result at 0.\n2. Loop exactly 32 times, not 'while n', so leading zeros are counted.\n3. Make room and drop in the bit: result = (result << 1) | (n & 1).\n4. Peel off one bit: n >>= 1.\n5. Return result."},
			{Level: 4, Type: "skeleton", Content: "def reverseBits(n):\n    result = 0\n    for _ in range(____):\n        result = ____ | ____\n        n >>= 1\n    return result"},
		},
		Solution: Solution{
			Code: `def reverseBits(n):
    result = 0
    for _ in range(32):
        result = (result << 1) | (n & 1)
        n >>= 1
    return result`,
			Explanation: "Number of 1 Bits read the bits of n; this rung also writes bits into a new number. Peeling the lowest bit of n (n & 1, then n >>= 1) and pushing it onto the low end of result (result << 1 | bit) reverses the order, exactly like reversing decimal digits with % 10 and * 10. The fixed width of 32 is the new detail: leading zeros matter.",
			BruteForce:  "Format n as a 32-character binary string, reverse it and parse it back: int(format(n, '032b')[::-1], 2). It is correct and also O(1) for fixed width, but it builds strings, and in a language without easy formatting you need the bit loop anyway.",
			KeyInsight:  "The bit you peel first (the lowest bit of n) must end up highest in the result. Shifting result left before adding each new bit pushes earlier bits towards the top, so after 32 rounds the first bit peeled sits at position 31.",
			WhyItWorks:  "Invariant: after k rounds, result holds bits 0..k-1 of n in reverse order, as a k-bit number (bit 0 of n is now at position k-1 of result, bit k-1 at position 0), and n has been shifted right k times so its current lowest bit is original bit k. In the next round result << 1 moves every stored bit up one place and | (n & 1) puts original bit k at position 0, which extends the invariant to k + 1. After exactly 32 rounds, original bit i sits at position 31 - i for every i, which is the reversal. Running a fixed 32 rounds matters: if the loop stopped when n became 0, the zeros above n's highest 1 bit would never be shifted in and result would be too small (1 would give 1 instead of 2147483648).",
			CommonMistakes: []string{
				"Looping 'while n:' instead of exactly 32 times, so leading zeros are skipped: reverseBits(1) returns 1 instead of 2147483648. Fix: always do 32 rounds.",
				"Adding the bit before shifting (result = result | bit; result <<= 1), which leaves an extra 0 at the bottom and pushes the top bit to position 32. Fix: shift first, then OR the bit in.",
				"Shifting n left or result right, which reverses nothing. Fix: read from the low end of n (n & 1, n >>= 1) and write to the low end of result (result << 1).",
			},
			TimeComplexity:  "O(1) (always 32 rounds)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Loop exactly 32 times", Explanation: "The width is fixed at 32 bits, so leading zeros count. Stopping early would drop them.", CodeSnippet: "for _ in range(32):", LineStart: 3, LineEnd: 3},
				{Title: "Make room and drop in the bit", Explanation: "result << 1 moves the bits collected so far up one place; | (n & 1) writes the current lowest bit of n into the freed position 0.", CodeSnippet: "result = (result << 1) | (n & 1)", LineStart: 4, LineEnd: 4},
				{Title: "Peel off one bit", Explanation: "n >>= 1 discards the bit just copied so the next round sees the next one. For n = 6 (110): bits come out 0, 1, 1, then 29 zeros.", CodeSnippet: "n >>= 1", LineStart: 5, LineEnd: 6},
			},
			Alternatives: []AltSolution{
				{
					Name: "Reverse a 32-character binary string",
					Code: `def reverseBits(n):
    bits = format(n, "032b")
    return int(bits[::-1], 2)`,
					TimeComplexity:  "O(1) (fixed 32 characters)",
					SpaceComplexity: "O(1) (fixed 32 characters)",
					Explanation:     "format(n, '032b') pads to 32 digits with leading zeros, so the reversal includes them. Short and handy as a cross-check, but it avoids the bit operations rather than practising them.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does 'while n:' give the wrong answer for n = 1, even though it processes every 1 bit?", Answer: "After the first round n is 0, so the loop stops with result = 1. The answer needs the single 1 bit at position 31, which only happens if result is shifted left 31 more times for the 31 leading zeros of n. The zeros carry information about position, so all 32 rounds are needed."},
			{Prompt: "If reverseBits is called millions of times, how could you speed it up?", Answer: "Precompute the reversal of every 8-bit value in a 256-entry table. Then split n into 4 bytes, look each one up, and assemble them in the opposite byte order with shifts: four lookups instead of 32 rounds. Another option is the divide-and-conquer swap: swap the 16-bit halves, then 8-bit groups, 4, 2 and 1 using masks, which is 5 steps."},
		},
	},
	{
		ID:              "sum-of-two-integers",
		Number:          371,
		Order:           6,
		Title:           "Sum of Two Integers",
		Slug:            "sum-of-two-integers",
		Difficulty:      "Medium",
		Category:        "bit-manipulation",
		Tags:            []string{"Math", "Bit Manipulation"},
		RelatedChapters: []int{20},
		Description: `Given two integers a and b, return a + b without using the + or - operators (and without sum() or similar helpers that add for you). Build the addition from bitwise operations.

Either number can be negative. Treat them as 32-bit two's complement integers, which is how most languages store an int; all inputs and the true sum fit comfortably in that range.`,
		Constraints: []string{
			"-1000 <= a, b <= 1000",
		},
		Examples: []Example{
			{ID: "1", Input: "a = 1, b = 2", Output: "3", Explanation: "01 ^ 10 = 11 and 01 & 10 = 0, so there is no carry and the answer is 11 = 3."},
			{ID: "2", Input: "a = 2, b = 3", Output: "5"},
			{ID: "3", Input: "a = -5, b = -7", Output: "-12"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"a": 1, "b": 2}, Expected: 3, Note: "no carries"},
			{Input: map[string]interface{}{"a": 2, "b": 3}, Expected: 5, Note: "one carry"},
			{Input: map[string]interface{}{"a": -5, "b": -7}, Expected: -12, Note: "two negatives: result must come back negative"},
			{Input: map[string]interface{}{"a": -1, "b": 1}, Expected: 0, Note: "opposites: carry ripples through all 32 bits; unmasked Python loops forever", Hidden: true},
			{Input: map[string]interface{}{"a": 0, "b": 0}, Expected: 0, Note: "both zero", Hidden: true},
			{Input: map[string]interface{}{"a": -12, "b": 5}, Expected: -7, Note: "mixed signs, negative result", Hidden: true},
			{Input: map[string]interface{}{"a": 1000, "b": 1000}, Expected: 2000, Note: "largest positive inputs", Hidden: true},
			{Input: map[string]interface{}{"a": 7, "b": -3}, Expected: 4, Note: "mixed signs, positive result", Hidden: true},
			{Input: map[string]interface{}{"a": -1000, "b": -1000}, Expected: -2000, Note: "smallest inputs", Hidden: true},
		},
		TimeComplexity:  "O(1) (at most 32 carry rounds)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def getSum(a, b):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Add two single bits by hand: 0+0, 0+1, 1+0, 1+1. Which bitwise operator gives the 'sum digit' of each case, and which gives the carry? And what happens to a negative number in Python if you keep shifting its carry left forever?"},
			{Level: 2, Type: "insight", Content: "a ^ b adds every column without carries and (a & b) << 1 is the carries, moved to the column they go into. Repeat with those two numbers until the carry is 0. Python integers are unbounded, so a negative number has infinitely many 1 bits and the carry never dies out: mask every value with 0xFFFFFFFF to keep 32 bits, and at the end turn a result above 0x7FFFFFFF back into a negative number with ~(a ^ 0xFFFFFFFF). Chapter 20 explains both steps.", ChapterRef: 20},
			{Level: 3, Type: "plan", Content: "1. Mask to 32 bits: a &= 0xFFFFFFFF and b &= 0xFFFFFFFF.\n2. Add without carry, then carry: while b is not 0, compute carry = ((a & b) << 1) & MASK, set a = a ^ b, then b = carry.\n3. Restore the sign: if a <= 0x7FFFFFFF it is a non-negative 32-bit value, return it; otherwise bit 31 is set, so return ~(a ^ MASK)."},
			{Level: 4, Type: "skeleton", Content: "def getSum(a, b):\n    MASK = 0xFFFFFFFF\n    MAX_INT = 0x7FFFFFFF\n\n    a &= MASK\n    b &= MASK\n    while b != 0:\n        carry = ____\n        a = ____\n        b = carry\n\n    if a <= MAX_INT:\n        return a\n    return ____"},
		},
		Solution: Solution{
			Code: `def getSum(a, b):
    MASK = 0xFFFFFFFF
    MAX_INT = 0x7FFFFFFF

    a &= MASK
    b &= MASK
    while b != 0:
        carry = ((a & b) << 1) & MASK
        a = a ^ b
        b = carry

    if a <= MAX_INT:
        return a
    return ~(a ^ MASK)`,
			Explanation: "The top rung combines everything in Chapter 20. XOR is addition without carry (the same fact that made Single Number work) and AND finds the columns that carry. The new difficulty is Python's unbounded integers: the problem's two's complement arithmetic only works inside a fixed width, so you mask to 32 bits during the loop and convert the 32-bit pattern back to a signed Python int at the end.",
			BruteForce:  "Simulate a ripple-carry adder: for each of the 32 bit positions compute the sum bit (a_i ^ b_i ^ carry) and the new carry, building the result one bit at a time. That always takes 32 rounds and needs the same masking and sign recovery; the XOR/AND loop handles all columns at once and usually stops after a few rounds.",
			KeyInsight:  "a + b == (a ^ b) + ((a & b) << 1): XOR adds each column ignoring carries, and AND shifted left is exactly the carries placed in the next column. Repeating the identity moves the carries left each round, so within 32 rounds the carry is gone from a 32-bit window.",
			WhyItWorks:  "Invariant: a + b (modulo 2^32) stays the same every round, because a + b == (a ^ b) + ((a & b) << 1) and masking only removes multiples of 2^32. Each round the lowest possible position of a carry moves at least one place left, so after at most 32 rounds the masked carry is 0 and a holds the sum modulo 2^32. Without the mask this fails for negatives in Python: -1 is ...111 with infinitely many 1s, so (-1 & 1) << 1 keeps producing a carry that never leaves the number. Masking makes every value a 32-bit pattern. To read the pattern as a signed number: if bit 31 is 0 (a <= 0x7FFFFFFF) it is already the right non-negative value; if bit 31 is 1 the true value is a - 2^32, and ~(a ^ MASK) computes exactly that without using minus: a ^ MASK flips the low 32 bits, and ~x == -x - 1 flips all bits including the infinite upper ones, which sets the upper bits to 1 and restores the low 32 bits. Example -5 + -7: the masked loop ends with a = 4294967284, which is above 0x7FFFFFFF, and ~(4294967284 ^ 0xFFFFFFFF) = ~11 = -12.",
			CommonMistakes: []string{
				"Porting the Java/C loop 'while b: a, b = a ^ b, (a & b) << 1' to Python unchanged, which loops forever for a = -1, b = 1 because Python integers do not overflow. Fix: mask with 0xFFFFFFFF inside the loop.",
				"Masking during the loop but returning the masked value directly, so -5 + -7 returns 4294967284 instead of -12. Fix: if bit 31 is set, convert back with ~(a ^ MASK).",
				"Updating a before computing the carry from the old a, so the carry is computed from the wrong value. Fix: compute carry from the current a and b first (or use tuple assignment), then update a.",
			},
			TimeComplexity:  "O(1) (at most 32 carry rounds)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Mask to 32 bits", Explanation: "Turn both inputs into their 32-bit two's complement patterns: -5 & 0xFFFFFFFF = 4294967291. From now on every value is a non-negative Python int below 2^32.", CodeSnippet: "MASK = 0xFFFFFFFF\na &= MASK\nb &= MASK", LineStart: 2, LineEnd: 6},
				{Title: "Add without carry, then carry", Explanation: "a ^ b is the column sums without carries; (a & b) << 1 is the carries moved one column left, masked so they fall off past bit 31 instead of growing forever. Repeat until there is nothing left to carry. For 2 + 3: a = 10 ^ 11 = 01, carry = 100; then a = 101 = 5, carry = 0.", CodeSnippet: "while b != 0:\n    carry = ((a & b) << 1) & MASK\n    a = a ^ b\n    b = carry", LineStart: 7, LineEnd: 10},
				{Title: "Restore the sign", Explanation: "A pattern with bit 31 clear is a non-negative number already. A pattern with bit 31 set represents a - 2^32; ~(a ^ MASK) produces that negative Python int.", CodeSnippet: "if a <= MAX_INT:\n    return a\nreturn ~(a ^ MASK)", LineStart: 12, LineEnd: 14},
			},
			Alternatives: []AltSolution{
				{
					Name: "Ripple-carry adder, one bit at a time",
					Code: `def getSum(a, b):
    MASK = 0xFFFFFFFF
    a &= MASK
    b &= MASK
    result = 0
    carry = 0
    for i in range(32):
        bit_a = (a >> i) & 1
        bit_b = (b >> i) & 1
        result |= (bit_a ^ bit_b ^ carry) << i
        carry = (bit_a & bit_b) | (bit_a & carry) | (bit_b & carry)
    if (result >> 31) & 1:
        return ~(result ^ MASK)
    return result`,
					TimeComplexity:  "O(1) (always 32 rounds)",
					SpaceComplexity: "O(1)",
					Explanation:     "How a hardware full adder works: each sum bit is the XOR of the two input bits and the incoming carry, and the outgoing carry is 1 when at least two of those three bits are 1. It shows the column-by-column logic explicitly, with the same mask and sign recovery.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the unmasked loop run forever in Python for a = -1, b = 1, but not in Java?", Answer: "In Python -1 has infinitely many 1 bits, so a & b keeps finding a carry and shifting it left never pushes it out: the carry just moves to higher and higher positions. In Java an int has 32 bits, so after 32 shifts the carry falls off the top and becomes 0. The mask 0xFFFFFFFF imitates that fixed width in Python."},
			{Prompt: "Why is ~(a ^ MASK) the right way to turn a 32-bit pattern with bit 31 set into a negative Python int?", Answer: "The pattern a stands for the value a - 2^32. a ^ MASK flips the low 32 bits, giving MASK - a = 2^32 - 1 - a. Then ~x = -x - 1 gives -(2^32 - 1 - a) - 1 = a - 2^32. In bit terms, the XOR flips the low bits and ~ flips them back while turning all the infinitely many upper bits into 1s, which is how Python stores negative numbers."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, BitManipulationProblems...)
}
