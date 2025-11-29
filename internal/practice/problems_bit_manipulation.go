package practice

// BitManipulationProblems contains all bit-manipulation category problems
var BitManipulationProblems = []*Problem{
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
	{
		ID:              "number-of-1-bits",
		Number:          97,
		Title:           "Number of 1 Bits",
		Difficulty:      "Easy",
		Category:        "bit-manipulation",
		Tags:            []string{"Divide and Conquer", "Bit Manipulation"},
		RelatedChapters: []int{1},
		Description: `Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).`,
		Constraints: []string{
			"The input must be a binary string of length 32",
		},
		Examples: []Example{
			{Input: "n = 00000000000000000000000000001011", Output: "3"},
			{Input: "n = 00000000000000000000000010000000", Output: "1"},
			{Input: "n = 11111111111111111111111111111101", Output: "31"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 11}, Expected: 3},
			{Input: map[string]interface{}{"n": 128}, Expected: 1},
		},
		TimeComplexity:  "O(1)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def hammingWeight(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Count bits by checking each bit position, or use n & (n-1) trick to clear rightmost 1 bit."},
			{Level: 2, Type: "algorithm", Content: "n & (n-1) clears the least significant 1 bit. Count how many times until n becomes 0."},
			{Level: 3, Type: "code", Content: "count = 0. while n: n &= (n - 1); count += 1. return count."},
		},
		Solution: Solution{
			Code: `def hammingWeight(n):
    count = 0
    while n:
        n &= (n - 1)  # Clear rightmost 1 bit
        count += 1
    return count`,
			Explanation:     "n & (n-1) clears the least significant 1 bit. Count iterations until n becomes 0.",
			TimeComplexity:  "O(1) - at most 32 iterations",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Clear rightmost 1", Explanation: "n & (n-1) removes least significant 1 bit", CodeSnippet: "n &= (n - 1)", LineStart: 4, LineEnd: 4},
				{Title: "Count iterations", Explanation: "Each iteration removes one 1 bit", CodeSnippet: "count += 1", LineStart: 5, LineEnd: 5},
			},
		},
	},
	{
		ID:              "reverse-bits",
		Number:          98,
		Title:           "Reverse Bits",
		Difficulty:      "Easy",
		Category:        "bit-manipulation",
		Tags:            []string{"Divide and Conquer", "Bit Manipulation"},
		RelatedChapters: []int{1},
		Description: `Reverse bits of a given 32 bits unsigned integer.`,
		Constraints: []string{
			"The input must be a binary string of length 32",
		},
		Examples: []Example{
			{Input: "n = 00000010100101000001111010011100", Output: "964176192 (00111001011110000010100101000000)"},
			{Input: "n = 11111111111111111111111111111101", Output: "3221225471 (10111111111111111111111111111111)"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 43261596}, Expected: 964176192},
		},
		TimeComplexity:  "O(1)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def reverseBits(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Process each bit from right to left, building result from left to right."},
			{Level: 2, Type: "algorithm", Content: "For each of 32 bits: shift result left, add current bit (n & 1), shift n right."},
			{Level: 3, Type: "code", Content: "result = 0. for i in range(32): result = (result << 1) | (n & 1); n >>= 1. return result."},
		},
		Solution: Solution{
			Code: `def reverseBits(n):
    result = 0
    for i in range(32):
        result = (result << 1) | (n & 1)
        n >>= 1
    return result`,
			Explanation:     "Extract rightmost bit of n, add to result (shifted left). Repeat 32 times.",
			TimeComplexity:  "O(1) - always 32 iterations",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Extract and add bit", Explanation: "Get rightmost bit, add to result", CodeSnippet: "result = (result << 1) | (n & 1)", LineStart: 4, LineEnd: 4},
				{Title: "Move to next bit", Explanation: "Shift n right to process next bit", CodeSnippet: "n >>= 1", LineStart: 5, LineEnd: 5},
			},
		},
	},
	{
		ID:              "sum-of-two-integers",
		Number:          99,
		Title:           "Sum of Two Integers",
		Difficulty:      "Medium",
		Category:        "bit-manipulation",
		Tags:            []string{"Math", "Bit Manipulation"},
		RelatedChapters: []int{1},
		Description: `Given two integers a and b, return the sum of the two integers without using the operators + and -.`,
		Constraints: []string{
			"-1000 <= a, b <= 1000",
		},
		Examples: []Example{
			{Input: "a = 1, b = 2", Output: "3"},
			{Input: "a = 2, b = 3", Output: "5"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"a": 1, "b": 2}, Expected: 3},
			{Input: map[string]interface{}{"a": 2, "b": 3}, Expected: 5},
		},
		TimeComplexity:  "O(1)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def getSum(a, b):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "XOR gives sum without carry. AND shifted left gives carry. Repeat until no carry."},
			{Level: 2, Type: "algorithm", Content: "sum = a ^ b (no carry). carry = (a & b) << 1. Repeat with sum and carry until carry is 0."},
			{Level: 3, Type: "code", Content: "while b: a, b = a ^ b, (a & b) << 1. Handle negative numbers with mask in Python."},
		},
		Solution: Solution{
			Code: `def getSum(a, b):
    # 32-bit mask for handling negative numbers in Python
    MASK = 0xFFFFFFFF
    MAX_INT = 0x7FFFFFFF

    while b != 0:
        carry = ((a & b) << 1) & MASK
        a = (a ^ b) & MASK
        b = carry

    # Handle negative result
    return a if a <= MAX_INT else ~(a ^ MASK)`,
			Explanation:     "XOR gives sum without carry, AND gives carry positions. Repeat until no carry. Handle Python's arbitrary precision integers.",
			TimeComplexity:  "O(1) - at most 32 iterations",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Calculate carry", Explanation: "AND finds carry bits, shift left for position", CodeSnippet: "carry = ((a & b) << 1) & MASK", LineStart: 7, LineEnd: 7},
				{Title: "Calculate sum", Explanation: "XOR gives sum without considering carry", CodeSnippet: "a = (a ^ b) & MASK", LineStart: 8, LineEnd: 8},
				{Title: "Handle negative", Explanation: "Convert back to signed if needed", CodeSnippet: "return a if a <= MAX_INT else ~(a ^ MASK)", LineStart: 12, LineEnd: 12},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, BitManipulationProblems...)
}
