package chapter

// chapter20 is a primer on bit manipulation for learners who have never
// worked with binary: representation, the six operators, two's complement and
// Python's unbounded integers, and the core tricks (x & 1, x & (x - 1), XOR
// cancellation) that the bit-manipulation ladder relies on.
// Storyboards to include (IDs registered in app.js): none.
func chapter20() Chapter {
	return Chapter{
		Number:        20,
		Slug:          "bit-manipulation",
		Title:         "Bit Manipulation",
		Category:      "bit-manipulation",
		Prerequisites: []int{1},
		Summary:       "A from-scratch primer on how integers are stored as bits: reading binary, the operators & | ^ ~ << >>, two's complement for negative numbers and why Python needs a 32-bit mask, plus the three tricks behind most interview problems: x & 1, x & (x - 1) and XOR cancellation.",
		Motivation: `Here is a list of 30,000 numbers in which every value appears exactly twice, except one. Find the odd one out. The natural plan is a dictionary of counts (Chapter 2): it works, but it stores about 15,000 entries. Sorting first avoids the dictionary but costs O(n log n). Now suppose you are told to use only a constant amount of memory and a single pass.

It sounds impossible until you look at how the computer stores the numbers. Every integer is a row of bits (0s and 1s), and the operator ^ (XOR) combines two rows column by column so that equal bits give 0. A value XOR-ed with itself is therefore 0, and XOR-ing all 30,000 numbers makes every pair disappear: one variable and one loop find the answer.

Bits also explain things that otherwise look like magic: why 12 & 11 equals 8, why -1 >> 1 is still -1 in Python, and how a CPU adds numbers with no + operator at all. This chapter starts from "what is binary" and builds up to those tricks, one operator at a time.`,
		Objectives: []string{
			"Convert small numbers between decimal and binary and read a bit pattern such as 1011 as 8 + 2 + 1",
			"Predict the result of &, |, ^, ~, << and >> on small numbers by working column by column",
			"Explain why x & (x - 1) clears exactly the lowest set bit and why XOR makes pairs of equal values cancel",
			"Explain two's complement, why Python integers behave as if they had infinitely many sign bits, and how to mask to 32 bits and recover a negative result",
			"Implement bit-by-bit loops (peel off one bit, make room and drop in a bit) such as counting and reversing bits",
		},
		WhenToUse: []string{
			"\"Every element appears twice except one\" together with O(1) extra space: XOR cancellation",
			"The statement talks about bits, binary representation, 1 bits, 32-bit (unsigned) integers or powers of two",
			"Arithmetic is restricted, e.g. \"without using + or -\": build it from XOR (sum without carry) and AND (carries)",
			"You need a set of up to 32 or 64 small items (letters, cities) packed into one integer as a bitmask",
			"A counting table where each number's answer follows from a smaller number with one bit removed (i >> 1 or i & (i - 1))",
		},
		Concepts: []Concept{
			{
				Name:        "Binary representation and the six operators",
				Description: "An integer is stored as bits, each worth a power of two. Bitwise operators work on those bits column by column: & (AND), | (OR), ^ (XOR) and ~ (NOT) combine or flip them, and << and >> slide the whole row left or right.",
				WhyItWorks:  `In binary the rightmost bit is worth 1, the next 2, then 4, 8, 16 and so on, each place worth twice the one to its right. So 1011 = 8 + 0 + 2 + 1 = 11 and 1100 = 8 + 4 = 12. The operators &, | and ^ look at one column at a time and ignore the others: AND gives 1 only when both bits are 1, OR gives 1 when at least one is 1, XOR gives 1 when exactly one is 1 (the bits differ). Shifting left by k moves every bit k places up, which multiplies by 2^k (5 << 1 = 1010 = 10); shifting right by k drops the lowest k bits, which is floor division by 2^k (13 >> 1 = 110 = 6). Because each column is independent for &, | and ^, you can always reason about one bit position at a time and then put the columns back together.`,
				Intuition:   `Picture a row of light switches, the highest on the left. A number is a pattern of on and off switches. AND keeps a light on only where both patterns have it on (an overlap), OR turns on everything either pattern has on (a union), XOR turns on the switches where the patterns disagree. Shifting slides the whole pattern along the row; lights pushed off the right end are lost.`,
				CoreIdeas: []string{
					"Place values are 1, 2, 4, 8, ...: 13 = 1101 = 8 + 4 + 1. In Python, bin(13) == '0b1101', int('1101', 2) == 13 and format(13, '08b') == '00001101'",
					"12 = 1100 and 10 = 1010: 12 & 10 = 1000 = 8, 12 | 10 = 1110 = 14, 12 ^ 10 = 0110 = 6",
					"x << k == x * 2^k and x >> k == x // 2^k: 3 << 2 = 12, 12 >> 2 = 3",
					"~x flips every bit; for Python integers ~x == -x - 1, so ~12 == -13",
					"Read bit i with (x >> i) & 1, set it with x | (1 << i), clear it with x & ~(1 << i), toggle it with x ^ (1 << i)",
				},
				CommonMistakes: []string{
					"Using ^ as 'to the power of'. In Python 2 ^ 3 is 1 (10 XOR 11 = 01), not 8. Powers are written 2 ** 3.",
					"Forgetting operator precedence: + and - bind more tightly than << >> and &, so a + b & 1 means (a + b) & 1 and 1 << n - 1 means 1 << (n - 1). Add parentheses whenever you mix them.",
					"Treating the decimal digits as bits, e.g. reading 11 as '1011' digit by digit or testing n % 10 == 1. Bits come from powers of two; use n & 1 and n >> 1.",
				},
				Examples: []CodeExample{
					{
						ID:       "operator-table",
						Title:    "The six operators on 12 and 10",
						Language: "python",
						Snippet: `a, b = 12, 10
print(format(a, "04b"), format(b, "04b"))  # 1100 1010
print(a & b)    # 8   (1000): 1 where both are 1
print(a | b)    # 14  (1110): 1 where either is 1
print(a ^ b)    # 6   (0110): 1 where they differ
print(~a)       # -13 (~x == -x - 1 in Python)
print(a << 1)   # 24  (11000): times 2
print(a >> 2)   # 3   (11): floor-divide by 4`,
						Notes: "Write the bits in columns, one number above the other, and combine each column on its own.",
					},
				},
			},
			{
				Name:        "Two's complement and Python's unbounded integers",
				Description: "Fixed-width integers store negative numbers in two's complement: the top bit counts as a large negative value. Python integers have no fixed width, so a negative number acts as if it had infinitely many leading 1s. To imitate 32-bit arithmetic you mask with 0xFFFFFFFF and convert back at the end.",
				WhyItWorks: `In a 32-bit two's complement integer, bit 31 is worth -2^31 instead of +2^31 and the others keep their usual values. So 0xFFFFFFFF (all 32 bits set) is -2^31 + (2^31 - 1) = -1, and -x == ~x + 1, because x + ~x is all ones, which is -1. The benefit is that the CPU adds signed and unsigned numbers with the same circuit, simply dropping any carry out of bit 31.

Python's integers grow as large as needed, so there is no bit 31 to drop a carry from. Instead Python behaves as if a negative number had infinitely many 1 bits to the left: -1 is ...11111, -5 is ...11011. That is why -1 >> 1 is still -1 and why a loop like 'while n: n >>= 1' never ends for a negative n. To work in exactly 32 bits, AND with MASK = 0xFFFFFFFF: -5 & MASK = 4294967291 = 0xFFFFFFFB, the 32-bit pattern of -5. To read such a pattern back as signed: if it is at most 0x7FFFFFFF (bit 31 clear) it is already the right value; otherwise its value is x - 2^32, which ~(x ^ MASK) computes using only bit operations (x ^ MASK flips the low 32 bits, and ~ flips everything back while turning the infinitely many upper bits into 1s).`,
				Intuition: `A three-digit odometer that rolls backwards: one step below 000 it shows 999. Two's complement uses that roll-over to mean -1, and 998 means -2, and so on. Python's integers are an odometer with infinitely many wheels, so rolling below 0 shows ...999 forever; the mask is you covering all but the last 32 wheels with tape.`,
				CoreIdeas: []string{
					"In 32 bits, bit 31 is worth -2^31: 0x7FFFFFFF = 2147483647 is the largest int, 0x80000000 is -2147483648, 0xFFFFFFFF is -1",
					"Negation is 'flip all bits and add 1': -x == ~x + 1",
					"Python ints are unbounded; negatives behave as if they had infinitely many leading 1s, so -1 >> 1 == -1",
					"Mask to 32 bits: x & 0xFFFFFFFF gives the unsigned 32-bit pattern (-5 -> 4294967291)",
					"Restore the sign: if x > 0x7FFFFFFF, the signed value is ~(x ^ 0xFFFFFFFF), which equals x - 2^32",
				},
				CommonMistakes: []string{
					"Porting a Java or C bit loop to Python unchanged. Shifting a negative number's carry left never makes it fall off, so the loop can run forever; mask every intermediate value with 0xFFFFFFFF.",
					"Masking but forgetting to convert back, so -5 + -7 returns 4294967284 instead of -12. Check bit 31 at the end and convert.",
					"Expecting bin(-5) to show the two's complement pattern. Python prints '-0b101' (a minus sign and the magnitude); use format(-5 & 0xFFFFFFFF, '032b') to see the 32-bit pattern.",
				},
				Examples: []CodeExample{
					{
						ID:       "mask-and-restore",
						Title:    "Converting between Python ints and 32-bit patterns",
						Language: "python",
						Snippet: `MASK = 0xFFFFFFFF
MAX_INT = 0x7FFFFFFF

def to_unsigned32(x):
    return x & MASK

def to_signed32(pattern):
    if pattern <= MAX_INT:
        return pattern
    return ~(pattern ^ MASK)

print(to_unsigned32(-5))           # 4294967291 (0xFFFFFFFB)
print(to_signed32(4294967291))     # -5
print(to_signed32(to_unsigned32(2**31)))  # -2147483648: 32-bit overflow`,
						Notes: "Sum of Two Integers uses exactly these two steps around its loop.",
					},
				},
			},
			{
				Name:        "Core tricks: x & 1, x & (x - 1) and XOR cancellation",
				Description: "Three facts do most of the work in interview problems: x & 1 reads the lowest bit (odd or even), x & (x - 1) clears the lowest 1 bit, and x ^ x == 0 lets equal values cancel no matter where they are.",
				WhyItWorks: `x & 1 keeps only bit 0, which is 1 exactly for odd numbers; together with x >> 1 (drop that bit) it lets you peel a number apart one bit at a time.

x & (x - 1): write x as (high bits) 1 followed by t zeros, where that 1 is its lowest set bit. Subtracting 1 must borrow from that 1, so x - 1 is (same high bits) 0 followed by t ones. AND-ing them keeps the high bits (they are equal) and zeroes the rest. Worked example: x = 12 = 1100, x - 1 = 11 = 1011, 1100 & 1011 = 1000 = 8: only the lowest 1 (worth 4) disappeared. Repeating it removes one 1 bit per step, so it counts 1 bits, and x & (x - 1) == 0 means x had a single 1 bit (a power of two).

XOR cancellation: x ^ x == 0 (every column agrees), x ^ 0 == x, and ^ is commutative and associative, so a long chain of XORs can be reordered freely. In a ^ b ^ a, regroup to (a ^ a) ^ b = 0 ^ b = b. Every value that appears an even number of times vanishes and the odd one out remains.`,
				Intuition: `XOR is a light switch you flip: flipping the same switch twice leaves the light as it was, whatever you did in between. x & (x - 1) is like an odometer rolling back by one: only the lowest non-zero wheel and the wheels below it change, and AND-ing with the old reading erases exactly that lowest non-zero wheel.`,
				CoreIdeas: []string{
					"x & 1 == 1 means x is odd; x >> 1 drops the lowest bit (the same as x // 2)",
					"x & (x - 1) clears the lowest 1 bit: 12 (1100) -> 8 (1000) -> 0",
					"x > 0 and x & (x - 1) == 0 exactly when x is a power of two",
					"x ^ x == 0, x ^ 0 == x, and order does not matter, so pairs cancel: 4 ^ 1 ^ 2 ^ 1 ^ 2 == 4",
					"Build a number bit by bit with result = (result << 1) | bit: make room, then drop the bit in",
				},
				CommonMistakes: []string{
					"Starting an XOR fold at the first element and then XOR-ing it again in the loop, which cancels it. Start the fold at 0, the value XOR leaves unchanged.",
					"Using 'while n:' to loop over a fixed-width value (e.g. 32 bits to reverse), which stops early and ignores leading zeros. Loop exactly 32 times when the width matters.",
					"Assuming x & (x - 1) clears the highest bit or the last bit position. It clears the lowest bit that is 1; for 6 (110) the result is 4 (100), not 2 or 6 >> 1.",
				},
				Examples: []CodeExample{
					{
						ID:       "core-tricks",
						Title:    "The three tricks in action",
						Language: "python",
						Snippet: `print(7 & 1, 8 & 1)          # 1 0  -> 7 is odd, 8 is even
x = 12                       # 1100
print(x & (x - 1))           # 8 (1000): lowest 1 bit cleared
print(64 & 63 == 0)          # True: 64 is a power of two
nums = [4, 1, 2, 1, 2]
acc = 0
for v in nums:
    acc ^= v                 # 4, 5, 7, 6, 4
print(acc)                   # 4: the pairs cancelled`,
						Notes: "Single Number, Number of 1 Bits, Counting Bits and Missing Number are each one of these lines plus a loop.",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Hamming distance: count the bit positions where two numbers differ",
				Problem: "Given non-negative integers x and y, return how many bit positions have different bits. For x = 1 and y = 4 the answer is 2 (0001 vs 0100). This combines XOR with the x & (x - 1) trick, the same subgoals used in Single Number and Number of 1 Bits.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Write it in binary",
						Explanation: "Before choosing an operator, write the small example in columns. Seeing the bits makes the question concrete: we are counting columns where the two rows disagree.",
						Code:        "# x = 1 -> 0001\n# y = 4 -> 0100",
						State:       "columns (8 4 2 1): x = 0 0 0 1, y = 0 1 0 0; they differ in the 4s and 1s columns",
					},
					{
						Subgoal:     "Pick the bit trick",
						Explanation: "XOR gives 1 exactly in the columns where the bits differ, so x ^ y turns 'count differing columns' into 'count the 1 bits of one number'.",
						Code:        "diff = x ^ y",
						State:       "diff = 0001 ^ 0100 = 0101 = 5",
					},
					{
						Subgoal:     "Peel off one bit",
						Explanation: "diff & (diff - 1) removes the lowest 1 bit. Repeat until diff is 0; each round removes exactly one 1, so the loop runs once per differing column.",
						Code: `count = 0
while diff:
    diff &= diff - 1
    count += 1`,
						State: "diff: 0101 -> 0100 (count 1) -> 0000 (count 2)",
					},
					{
						Subgoal:     "Count the peel",
						Explanation: "The number of rounds equals the number of 1 bits in x ^ y, which is the number of differing positions.",
						Code:        "return count",
						State:       "returns 2",
					},
				},
				Code: `def hamming_distance(x, y):
    diff = x ^ y
    count = 0
    while diff:
        diff &= diff - 1
        count += 1
    return count`,
			},
			{
				Title:   "Simulate 32-bit integer overflow in Python",
				Problem: "Write to_int32(x) that returns what a 32-bit signed integer would hold after computing x, e.g. to_int32(2147483647 + 1) == -2147483648 and to_int32(-5) == -5. Sum of Two Integers wraps its loop in exactly these subgoals.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Write it in binary",
						Explanation: "2147483647 + 1 = 2^31 is a 1 followed by 31 zeros. In 32 bits that 1 lands in bit 31, the sign bit, so a 32-bit int reads it as negative. Python just sees a positive number.",
						Code:        "# 2**31 = 1000 0000 0000 0000 0000 0000 0000 0000",
						State:       "x = 2147483648; bit 31 is set",
					},
					{
						Subgoal:     "Mask to 32 bits",
						Explanation: "AND with 0xFFFFFFFF keeps the low 32 bits and drops everything above, which is what a 32-bit register does. For a negative x this also cuts the infinite run of leading 1s down to 32 bits.",
						Code: `MASK = 0xFFFFFFFF
x &= MASK`,
						State: "2147483648 stays 2147483648; -5 becomes 4294967291 (0xFFFFFFFB)",
					},
					{
						Subgoal:     "Check bit 31",
						Explanation: "A pattern at most 0x7FFFFFFF has bit 31 clear, so it is a non-negative number and already correct.",
						Code: `if x <= 0x7FFFFFFF:
    return x`,
						State: "2147483648 > 2147483647, so bit 31 is set: keep going",
					},
					{
						Subgoal:     "Restore the sign",
						Explanation: "With bit 31 set the pattern means x - 2^32. x ^ MASK flips the low 32 bits and ~ flips them back while setting all the upper bits, producing that negative Python int without using minus.",
						Code:        "return ~(x ^ MASK)",
						State:       "~(2147483648 ^ 4294967295) = ~2147483647 = -2147483648; for -5: ~(4294967291 ^ 4294967295) = ~4 = -5",
					},
				},
				Code: `def to_int32(x):
    MASK = 0xFFFFFFFF
    x &= MASK
    if x <= 0x7FFFFFFF:
        return x
    return ~(x ^ MASK)`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Power of Two",
				Prompt: "Return True when n is a power of two (1, 2, 4, 8, ...). Use the subgoal Pick the bit trick: a power of two has exactly one 1 bit. Remember that 0 and negative numbers are not powers of two.",
				Code: `def is_power_of_two(n):
    return n > 0 and ____`,
				Blanks:      []string{"n & (n - 1) == 0"},
				Explanation: "Clearing the lowest 1 bit of a number with exactly one 1 bit leaves 0; any other positive number still has a 1 bit left. The n > 0 check comes first because 0 & -1 == 0 would otherwise count 0 as a power of two. In Python & binds more tightly than ==, so n & (n - 1) == 0 is (n & (n - 1)) == 0.",
			},
			{
				Title:  "Get, set and clear one bit",
				Prompt: "Fill in the three helpers. Bit i is the bit worth 2^i, counting from 0 at the right. Each one moves a single 1 to position i with 1 << i or moves bit i down to position 0 with >> i.",
				Code: `def get_bit(x, i):
    return ____

def set_bit(x, i):
    return ____

def clear_bit(x, i):
    return ____`,
				Blanks:      []string{"(x >> i) & 1", "x | (1 << i)", "x & ~(1 << i)"},
				Explanation: "Shifting x right by i brings bit i to position 0, where & 1 isolates it. OR with a single 1 at position i turns that bit on and leaves the others alone (b | 0 == b). ~(1 << i) is all ones except position i, so AND with it turns only bit i off. Example: x = 10 (1010): get_bit(x, 1) == 1, set_bit(x, 0) == 11, clear_bit(x, 3) == 2.",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Trace it: what does 13 & (13 - 1) evaluate to?",
				Choices: []string{
					"12",
					"1",
					"8",
					"0",
				},
				Answer:      0,
				Explanation: "13 = 1101 and 12 = 1100. AND-ing column by column gives 1100 = 12: the lowest 1 bit of 13 (worth 1) was cleared and the others kept. 8 would be the answer for 12 & 11, and 0 only happens for powers of two.",
			},
			{
				Prompt: "Which pattern fits? \"Every value in a list of 30,000 integers appears exactly twice except one. Find it using O(1) extra memory.\"",
				Choices: []string{
					"A hash map of counts (Chapter 2)",
					"Sort, then compare neighbours with two pointers (Chapter 3)",
					"XOR all the values (this chapter)",
					"A sliding window over the list (Chapter 4)",
				},
				Answer:      2,
				Explanation: "XOR folds the list in one pass with one variable, and the pairs cancel. A hash map gives the right answer but uses O(n) memory. Sorting works in O(1) extra memory only with an in-place sort, and costs O(n log n). A sliding window needs a contiguous range with a condition, which this problem does not have.",
			},
			{
				Prompt: "In Python, what is -1 >> 1, and what does that mean for the loop 'while n: n >>= 1' when n = -8?",
				Choices: []string{
					"0; the loop ends after a few rounds",
					"-1; the loop never ends, because n gets stuck at -1",
					"2147483647; the loop ends after 31 rounds",
					"It raises an error, because negative numbers cannot be shifted",
				},
				Answer:      1,
				Explanation: "Python integers behave as if negatives had infinitely many leading 1 bits, and >> copies that sign bit in from the left (floor division by 2: -1 // 2 == -1). -8 becomes -4, -2, -1, and then stays -1 forever, which is never 0. Masking with 0xFFFFFFFF first turns n into a non-negative 32-bit pattern and makes the loop end.",
			},
			{
				Prompt:      "Explain, using x = 40 (101000), why x & (x - 1) clears exactly the lowest 1 bit and no other bit.",
				Explanation: "40 - 1 = 39 = 100111: subtracting 1 borrows from the lowest 1 bit (worth 8), turning it into 0 and the three zeros below it into 1s, while the bits above it stay the same. AND-ing 101000 with 100111: the high bit 1 & 1 stays 1, the borrowed bit is 1 & 0 = 0, and the three low bits are 0 & 1 = 0. The result is 100000 = 32, exactly 40 with its lowest 1 bit removed.",
			},
			{
				Prompt:      "Sum of Two Integers adds with a ^ b and (a & b) << 1 in a loop. Why does the Python version need the 0xFFFFFFFF mask, and how is a negative result recovered at the end?",
				Explanation: "In a fixed 32-bit integer the carry eventually shifts past bit 31 and disappears, so the loop ends. Python integers are unbounded and negative numbers have infinitely many 1 bits, so for -1 + 1 the carry keeps moving left forever. Masking every value with 0xFFFFFFFF imitates the 32-bit register. The final masked pattern is non-negative in Python; if bit 31 is set (the value is above 0x7FFFFFFF) it represents value - 2^32, which ~(a ^ 0xFFFFFFFF) produces: for -5 + -7 the pattern 4294967284 becomes ~11 = -12.",
			},
			{
				Prompt: "Chapter 1 review: a loop does n >>= 1 until n is 0. For n = 1,000,000, about how many rounds does it run, and what is its Big O in terms of the value n?",
				Choices: []string{
					"1,000,000 rounds: O(n)",
					"About 20 rounds: O(log n), one per binary digit",
					"About 1,000 rounds: O(sqrt(n))",
					"1 round: O(1)",
				},
				Answer:      1,
				Explanation: "Each shift halves n, and 1,000,000 has 20 binary digits (2^19 = 524,288 <= 1,000,000 < 2^20), so the loop runs 20 times. Halving each round is the same growth as binary search: O(log n). For values bounded by 32 bits it is at most 32 rounds, which is why such loops are often called O(1).",
			},
		},
	}
}
