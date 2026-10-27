package chapter

// chapter19 teaches the math-geometry category: matrix index manipulation
// (transpose + reverse, layer-by-layer boundaries), using the input as
// storage, digit arithmetic with carries, and fast exponentiation.
// Storyboards to include (IDs registered in app.js): none.
func chapter19() Chapter {
	return Chapter{
		Number:        19,
		Slug:          "math-geometry",
		Title:         "Math and Matrix Manipulation",
		Category:      "math-geometry",
		Prerequisites: []int{1},
		Summary:       "Four small toolkits that keep showing up in interviews: mapping matrix indices (rotate, transpose, walk layer by layer), storing flags inside the input to save memory, doing arithmetic digit by digit with carries, and computing huge powers with O(log n) multiplications by reading the exponent in binary.",
		Motivation: `Suppose a phone app must rotate a 4000 x 4000 photo by 90 degrees. The obvious code builds a second 4000 x 4000 grid and copies every pixel into its new place. That is 16 million pixels, so the copy needs 64 MB of extra memory (at 4 bytes per pixel), more than a small device may have to spare. Yet a rotation only moves values around: if you know where cell (r, c) must go, you can do it with swaps and no second grid.

Or suppose you need 3 raised to the power 1,000,000,000, modulo 1,000,000,007 (a common step in hashing and cryptography). Multiplying by 3 one billion times takes minutes even in a fast language. Reading the exponent in binary, it has only 30 digits, and fast exponentiation needs about 30 squarings plus at most 30 extra multiplications: around 60 operations instead of a billion.

Neither trick needs a new data structure. Both come from looking at the numbers more carefully: where an index goes, what the digits (decimal or binary) of a number say. This chapter teaches those habits with the problems in the math-geometry ladder.`,
		Objectives: []string{
			"Derive where cell (r, c) goes under a rotation or reflection and implement it in place with a transpose and a reverse",
			"Trace a layer-by-layer (spiral) walk using four boundaries and explain why the boundary checks prevent double visits on non-square matrices",
			"Explain how the first row and column of a matrix can store flags, and why they must be protected and processed last",
			"Implement addition and multiplication digit by digit with carries, and say which output position each digit product lands in",
			"Trace fast exponentiation on a small exponent and explain why it needs only about log2(n) rounds",
		},
		WhenToUse: []string{
			"The input is a 2D grid and the task is a geometric transformation (rotate, transpose, reflect) or a traversal order (spiral, diagonal, layer by layer)",
			"The statement demands that a matrix be modified in place or with O(1) extra space, which suggests storing markers inside the matrix itself",
			"Numbers arrive as digit lists or strings, possibly too long for a 64-bit integer, or you are told not to convert them: do the arithmetic digit by digit with a carry",
			"A huge exponent appears (x^n with n up to 2^31, or a^b mod m): use fast exponentiation",
			"Points on a grid and axis-aligned shapes (squares, rectangles): fix one or two corners, derive the others from coordinates, and count with a hash map",
		},
		Concepts: []Concept{
			{
				Name:        "Matrix index mapping",
				Description: "Before writing any loop, work out where a cell (r, c) must end up. Rotations and reflections are simple formulas on r and c, and every rotation is two reflections done one after the other, each of which is a set of swaps.",
				WhyItWorks:  `A reflection exchanges cells in pairs, so it can be done in place with swaps. Transposing reflects across the main diagonal: (r, c) <-> (c, r). Reversing each row reflects left-right: (r, c) <-> (r, n - 1 - c). Doing the transpose and then the row reversal sends (r, c) -> (c, r) -> (c, n - 1 - r), which is exactly a 90-degree clockwise turn. Any formula you can write as a chain of reflections can be done in place the same way. For traversal orders, the unvisited part of a spiral walk is always a rectangle, so four integers (top, bottom, left, right) describe it completely; finishing one side removes one row or column and shrinks one of them.`,
				Intuition:   `Think of a photo printed on transparent film. Flipping it over along its diagonal is the transpose; flipping it left-to-right is reversing each row. Two flips, one after the other, leave the photo facing up again but turned a quarter: the mapping is (r, c) -> (c, r) for the first flip and (r, c) -> (r, n - 1 - c) for the second.`,
				CoreIdeas: []string{
					"Clockwise: (r, c) -> (c, n - 1 - r) = transpose, then reverse each row",
					"Counterclockwise: (r, c) -> (n - 1 - c, r) = transpose, then reverse the order of the rows",
					"180 degrees: (r, c) -> (n - 1 - r, n - 1 - c) = reverse the order of the rows and reverse each row",
					"Transpose by swapping only pairs with c > r, so each pair is swapped exactly once",
					"Layer by layer: track top, bottom, left, right; walk a side, then shrink that boundary; recheck before walking back",
				},
				CommonMistakes: []string{
					"Transposing with both loops over the full range, which swaps every pair twice and changes nothing. Start the inner loop at r + 1.",
					"Mixing up 'reverse each row' (left-right flip) with 'reverse the row order' (top-bottom flip), which turns a clockwise rotation into a counterclockwise one. Check one cell against the formula.",
					"Assuming a square matrix in a spiral walk. For a 3 x 4 or 4 x 2 matrix the last layer is a single row or column; check top <= bottom and left <= right again before the return trip.",
				},
				Examples: []CodeExample{
					{
						ID:       "transpose-and-reflect",
						Title:    "Transpose and reflect in place",
						Language: "python",
						Snippet: `def transpose(matrix):
    n = len(matrix)
    for r in range(n):
        for c in range(r + 1, n):      # upper triangle only
            matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]

def reflect_left_right(matrix):
    for row in matrix:
        row.reverse()

# clockwise rotation = transpose, then reflect_left_right`,
						Notes: "Each function is a reflection, so each is a set of swaps that needs no second matrix.",
					},
				},
			},
			{
				Name:        "Using the input as storage",
				Description: "When the problem forbids extra memory, look for space inside the input that you can borrow, for example the first row and column of a matrix as flags, as long as you protect what was there before you overwrite it.",
				WhyItWorks:  `Set Matrix Zeroes needs one flag per column and one per row: m + n bits. Row 0 has exactly one cell per column and column 0 has exactly one cell per row, so they have room for those flags. Writing a 0 there is safe because a cell is only marked when its row or column will be cleared anyway, so no information that is still needed is lost, except whether row 0 and column 0 had zeros of their own. Two booleans saved before marking keep that, and processing row 0 and column 0 last means the markers are read before they are erased. The general pattern: record, then modify, and modify the storage area last.`,
				Intuition:   `It is like writing your to-do notes in the margin of a worksheet you will hand in anyway. The margin (row 0 and column 0) is part of the sheet, but before writing in it you note on a sticky label (the two booleans) whether the margin already had marks of its own, and you clean the margin up last.`,
				CoreIdeas: []string{
					"Two phases: first record what must change (markers), then apply the changes",
					"Pick storage that has exactly the right shape: one cell per row in column 0, one cell per column in row 0",
					"Save the original state of the borrowed cells (first_row_has_zero, first_col_has_zero) before overwriting them",
					"Process the borrowed storage last, after every reader of the markers is done",
				},
				CommonMistakes: []string{
					"Changing cells while still scanning, so changed cells are mistaken for original data (a new 0 spreads again). Finish recording before modifying.",
					"Clearing the marker row or column first, which wipes out every marker. Clear the storage area last.",
					"Using matrix[0][0] for both 'row 0' and 'column 0'. One cell cannot hold two flags; use a separate boolean for one of them.",
				},
				Examples: []CodeExample{
					{
						ID:       "markers-in-first-row",
						Title:    "Marking rows and columns inside the matrix",
						Language: "python",
						Snippet: `first_row_has_zero = any(v == 0 for v in matrix[0])
first_col_has_zero = any(row[0] == 0 for row in matrix)
for r in range(1, len(matrix)):
    for c in range(1, len(matrix[0])):
        if matrix[r][c] == 0:
            matrix[r][0] = 0    # flag row r
            matrix[0][c] = 0    # flag column c`,
						Notes: "Row 0 and column 0 become a free bit array of size m + n.",
					},
				},
			},
			{
				Name:        "Digit arithmetic with carries",
				Description: "Numbers stored as digit lists or strings are added and multiplied the way you did in school: start at the lowest digit, combine, write total % 10 and carry total // 10 to the next place. For an integer, n % 10 peels off its lowest digit and n // 10 removes it.",
				WhyItWorks:  `A digit at position i from the right is worth 10^i. When a column total reaches 10 or more, 10 units of 10^i are the same as one unit of 10^(i+1), so writing total % 10 and moving total // 10 one place left keeps the value of the whole number unchanged. When adding two numbers the column total is at most 9 + 9 + 1 = 19, so the carry is always 0 or 1. For multiplication, digit i of an m-digit number times digit j of an n-digit number has place value 10^((m - 1 - i) + (n - 1 - j)), which is index i + j + 1 in an answer array of length m + n, so every partial product has a fixed destination.`,
				Intuition:   `A car's odometer: when the ones wheel goes from 9 to 0 it nudges the tens wheel forward by one, and that nudge can ripple further (1999 -> 2000). The wheels are the digits, the nudge is the carry, and it always travels from the lowest wheel towards the highest.`,
				CoreIdeas: []string{
					"Start at the lowest digit: the last element of a digit list, or n % 10 of an integer",
					"Add with carry: total = digit_a + digit_b + carry",
					"Write the digit, keep the carry: write total % 10, carry = total // 10",
					"Flush the final carry: a leftover carry becomes a new leading digit",
					"In multiplication, digits i and j meet at answer index i + j + 1, with the carry going to i + j",
				},
				CommonMistakes: []string{
					"Walking a digit list from the front. The carry starts at the least significant digit, which is the last element.",
					"Forgetting the final carry, so 95 + 78 comes out as 73 instead of 173. After the loop, append any leftover carry.",
					"Using / instead of // in Python 3, which produces floats like 1.9 instead of the integer quotient 1.",
				},
				Examples: []CodeExample{
					{
						ID:       "digit-peeling",
						Title:    "Peel off digits of an integer",
						Language: "python",
						Snippet: `def digit_square_sum(n):
    total = 0
    while n > 0:
        digit = n % 10      # lowest digit
        total += digit * digit
        n //= 10            # drop it
    return total

print(digit_square_sum(19))  # 1 + 81 = 82`,
						Notes: "The same loop with % 2 and // 2 reads the binary digits instead, which is exactly what fast exponentiation does.",
					},
				},
			},
			{
				Name:        "Fast exponentiation",
				Description: "Compute x^n with about log2(n) multiplications by reading n in binary: repeatedly square x to get x, x^2, x^4, x^8, ..., and multiply into the answer the powers whose binary digit of n is 1.",
				WhyItWorks:  `Every n is a sum of distinct powers of two (its binary digits), and x^(a + b) = x^a * x^b. For n = 13 = 1101 in binary = 8 + 4 + 1, x^13 = x^8 * x^4 * x^1. Squaring turns x^(2^k) into x^(2^(k+1)) with one multiplication, so all the needed powers cost one multiplication per binary digit. In loop form the invariant is result * x^n == X^N (the original base to the original exponent): when n is odd, move one factor x into result; then square x and halve n. Each round halves n, so the loop runs floor(log2 N) + 1 times: 31 rounds for n around 2 billion instead of 2 billion multiplications. A negative exponent is handled first with x^(-n) = (1/x)^n.`,
				Intuition:   `Paying 13 euros with coins worth 1, 2, 4 and 8: you never need more than one coin of each value, and the coins you use are exactly the 1 digits of 13 in binary (8 + 4 + 1). The squarings make the coins (x, x^2, x^4, x^8) and the 1 bits say which coins go into the payment.`,
				CoreIdeas: []string{
					"Handle a negative exponent: if n < 0, use 1/x and -n",
					"Read the lowest bit: n % 2 (written n & 1 in Chapter 20)",
					"Multiply in when the bit is 1: result *= x",
					"Square the base, shift the exponent: x *= x and n //= 2 (n >>= 1)",
					"With a modulus, reduce after every multiplication: (a * b) % m == ((a % m) * (b % m)) % m",
				},
				CommonMistakes: []string{
					"Looping n times. For n near 2^31 that is billions of multiplications; halve n each round instead.",
					"Squaring x before multiplying it into result in the same round, which uses the wrong power. Check the bit, multiply, then square.",
					"In the recursive version, calling power(x, n // 2) twice instead of once and squaring the result, which brings the cost back to O(n).",
				},
				Examples: []CodeExample{
					{
						ID:       "fast-power",
						Title:    "Iterative fast power",
						Language: "python",
						Snippet: `def fast_power(x, n):          # n >= 0
    result = 1
    while n > 0:
        if n % 2 == 1:         # lowest binary digit is 1
            result *= x
        x *= x                 # x, x^2, x^4, x^8, ...
        n //= 2                # drop that binary digit
    return result

print(fast_power(3, 13))       # 1594323`,
						Notes: "13 = 1101 in binary, so the loop multiplies in 3, 81 and 6561 and skips 9.",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Rotate a square matrix 90 degrees counterclockwise, in place",
				Problem: "Given an n x n matrix, rotate it a quarter turn counterclockwise without allocating a second matrix. For [[1,2,3],[4,5,6],[7,8,9]] the answer is [[3,6,9],[2,5,8],[1,4,7]]. (The practice problem Rotate Image asks for the clockwise turn with the same subgoals.)",
				Steps: []WorkedStep{
					{
						Subgoal:     "Map the indices",
						Explanation: "Work out the target of one cell before coding. Counterclockwise, the right column becomes the top row, so (r, c) goes to (n - 1 - c, r). Check with a corner: 3 at (0, 2) must go to (0, 0).",
						Code:        "# (r, c) -> (n - 1 - c, r)",
						State:       "n=3: 1 (0,0)->(2,0); 3 (0,2)->(0,0); 9 (2,2)->(0,2)",
					},
					{
						Subgoal:     "Transpose across the diagonal",
						Explanation: "The transpose sends (r, c) to (c, r). Swap each pair above the diagonal with its mirror below. The inner loop starts at r + 1 so every pair is swapped once and the diagonal stays.",
						Code: `n = len(matrix)
for r in range(n):
    for c in range(r + 1, n):
        matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]`,
						State: "[[1,2,3],[4,5,6],[7,8,9]] -> [[1,4,7],[2,5,8],[3,6,9]]",
					},
					{
						Subgoal:     "Reverse to finish",
						Explanation: "We are at (c, r) and need (n - 1 - c, r): only the row index must be mirrored, which means reversing the order of the rows (not each row). Reversing each row instead would give the clockwise rotation.",
						Code:        "matrix.reverse()",
						State:       "[[1,4,7],[2,5,8],[3,6,9]] -> [[3,6,9],[2,5,8],[1,4,7]]",
					},
					{
						Subgoal:     "Check one cell",
						Explanation: "Follow 2, which starts at (0, 1). The formula says it should end at (n - 1 - 1, 0) = (1, 0). The transpose moved it to (1, 0) and reversing the row order maps row 1 to row 1, so it stays at (1, 0). It matches, and the whole matrix used only swaps.",
						Code:        "assert matrix == [[3, 6, 9], [2, 5, 8], [1, 4, 7]]",
						State:       "2: (0,1) -transpose-> (1,0) -reverse rows-> (1,0) = expected (1,0)",
					},
				},
				Code: `def rotate_counterclockwise(matrix):
    n = len(matrix)
    for r in range(n):
        for c in range(r + 1, n):
            matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
    matrix.reverse()`,
			},
			{
				Title:   "Add two numbers stored as digit lists",
				Problem: "Two non-negative numbers are given as digit lists, most significant digit first, e.g. a = [9, 5] (95) and b = [7, 8] (78). Return their sum as a digit list ([1, 7, 3]) without converting the lists to integers. (Plus One and Multiply Strings reuse these subgoals.)",
				Steps: []WorkedStep{
					{
						Subgoal:     "Start at the lowest digit",
						Explanation: "The ones digits are the last elements, so put one index at the end of each list. The lists may have different lengths, so keep going while either index is valid.",
						Code: `i, j = len(a) - 1, len(b) - 1
carry = 0
out = []
while i >= 0 or j >= 0:`,
						State: "a=[9,5] b=[7,8]: i=1 (digit 5), j=1 (digit 8), carry=0",
					},
					{
						Subgoal:     "Add with carry",
						Explanation: "The column total is the carry from the column to the right plus each digit that exists. A shorter list simply contributes nothing once its index runs out.",
						Code: `    total = carry
    if i >= 0:
        total += a[i]
        i -= 1
    if j >= 0:
        total += b[j]
        j -= 1`,
						State: "ones column: total = 0 + 5 + 8 = 13; tens column: total = 1 + 9 + 7 = 17",
					},
					{
						Subgoal:     "Write the digit, keep the carry",
						Explanation: "Only one digit fits in a column: keep total % 10 and send total // 10 (0 or 1, since a total is at most 19) to the next column.",
						Code: `    out.append(total % 10)
    carry = total // 10`,
						State: "ones: write 3, carry 1 -> out=[3]; tens: write 7, carry 1 -> out=[3,7]",
					},
					{
						Subgoal:     "Flush the final carry",
						Explanation: "When both lists are used up a carry may remain; it becomes a new leading digit. The digits were produced lowest first, so reverse at the end.",
						Code: `if carry:
    out.append(carry)
out.reverse()
return out`,
						State: "carry=1 -> out=[3,7,1] -> reversed [1,7,3] = 173",
					},
				},
				Code: `def add_digit_lists(a, b):
    i, j = len(a) - 1, len(b) - 1
    carry = 0
    out = []
    while i >= 0 or j >= 0:
        total = carry
        if i >= 0:
            total += a[i]
            i -= 1
        if j >= 0:
            total += b[j]
            j -= 1
        out.append(total % 10)
        carry = total // 10
    if carry:
        out.append(carry)
    out.reverse()
    return out`,
			},
			{
				Title:   "Compute 3^13 with fast exponentiation",
				Problem: "Compute x^n for x = 3, n = 13 using O(log n) multiplications. 13 is 1101 in binary. (Pow(x, n) adds negative exponents and floats to the same loop.)",
				Steps: []WorkedStep{
					{
						Subgoal:     "Handle a negative exponent",
						Explanation: "If n were negative we would replace x with 1/x and n with -n, because x^(-n) = (1/x)^n. Here n = 13 is positive, so nothing changes. Start result at 1, the empty product.",
						Code: `if n < 0:
    x, n = 1 / x, -n
result = 1`,
						State: "x=3, n=13 (1101), result=1",
					},
					{
						Subgoal:     "Read the lowest bit",
						Explanation: "n % 2 is the lowest binary digit of n. At round k, x holds 3^(2^k), the power that this bit stands for.",
						Code: `while n > 0:
    if n % 2 == 1:`,
						State: "round 1: n=13 (1101), bit=1, x=3^1",
					},
					{
						Subgoal:     "Multiply in when the bit is 1",
						Explanation: "A 1 bit means that power of two is part of n, so its power of x is part of the answer.",
						Code:        "        result *= x",
						State:       "round 1: result=3; round 2: bit 0, skip; round 3: result=3*81=243; round 4: result=243*6561=1594323",
					},
					{
						Subgoal:     "Square the base, shift the exponent",
						Explanation: "Squaring moves x to the next power 3^(2^(k+1)), and n //= 2 drops the bit just used. The invariant result * x^n == 3^13 holds after every round, so when n reaches 0, result is the answer.",
						Code: `    x *= x
    n //= 2
return result`,
						State: "x: 3 -> 9 -> 81 -> 6561 -> 43046721; n: 13 -> 6 -> 3 -> 1 -> 0; 4 rounds instead of 12 multiplications",
					},
				},
				Code: `def fast_power(x, n):
    if n < 0:
        x, n = 1 / x, -n
    result = 1
    while n > 0:
        if n % 2 == 1:
            result *= x
        x *= x
        n //= 2
    return result`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Fast modular power",
				Prompt: "Fill in the blanks so mod_pow(base, exp, mod) returns (base^exp) % mod for exp up to 10^18 using the subgoals Read the lowest bit, Multiply in when the bit is 1, and Square the base, shift the exponent. Keep every intermediate value below mod.",
				Code: `def mod_pow(base, exp, mod):
    result = 1
    base %= mod
    while exp > 0:
        if ____:
            result = ____
        base = ____
        exp //= 2
    return result`,
				Blanks:      []string{"exp % 2 == 1", "result * base % mod", "base * base % mod"},
				Explanation: "The structure is the same as fast_power. The only change is reducing modulo mod after each multiplication, which is allowed because (a * b) % m == ((a % m) * (b % m)) % m, and which keeps the numbers small. mod_pow(3, 13, 1000) gives 323, the last three digits of 1594323.",
			},
			{
				Title:  "Rotate a matrix 180 degrees in place",
				Prompt: "A half turn sends (r, c) to (n - 1 - r, n - 1 - c). Using the subgoal Reverse to finish twice, fill in the two reflections: one mirrors the row index, the other mirrors the column index.",
				Code: `def rotate_180(matrix):
    matrix.____         # mirror the row index: row r -> row n - 1 - r
    for row in matrix:
        row.____        # mirror the column index: col c -> col n - 1 - c`,
				Blanks:      []string{"reverse()", "reverse()"},
				Explanation: "Reversing the list of rows maps row r to row n - 1 - r, and reversing each row maps column c to column n - 1 - c. Together they give (n - 1 - r, n - 1 - c), the half turn, with no transpose needed. [[1,2],[3,4]] becomes [[4,3],[2,1]].",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Which pattern fits best? \"Given an m x n matrix, return its elements in diagonal zig-zag order: (0,0), then (0,1),(1,0), then (2,0),(1,1),(0,2), and so on.\"",
				Choices: []string{
					"Sliding window (Chapter 4), because the diagonals are contiguous",
					"Matrix index mapping (this chapter): cells on one diagonal share r + c, so walk s = r + c from 0 to m + n - 2 and alternate the direction",
					"Binary search (Chapter 6), because the rows are in order",
					"A hash map from value to position (Chapter 2)",
				},
				Answer:      1,
				Explanation: "The task is a traversal order over a grid, which is this chapter's index reasoning: every cell with the same r + c lies on one anti-diagonal, so the order is fully described by a formula on indices. A sliding window needs a contiguous range of a 1D sequence with a condition to maintain, binary search needs a monotonic property, and a hash map does not help produce an order.",
			},
			{
				Prompt: "You must compute 7^(10^18) mod 1,000,000,007. Which approach finishes in well under a second?",
				Choices: []string{
					"Multiply by 7 in a loop 10^18 times, reducing mod each time",
					"Fast exponentiation: about 60 rounds, one per binary digit of 10^18",
					"Dynamic programming over a table of size 10^18 (Chapter 15)",
					"Compute 7^(10^18) exactly with Python's big integers, then take the remainder",
				},
				Answer:      1,
				Explanation: "10^18 has about 60 binary digits (2^60 is about 1.15 * 10^18), so fast power needs about 60 squarings and at most 60 extra multiplications, each reduced mod 1,000,000,007. The loop would take thousands of years, a table of size 10^18 cannot be stored, and the exact power has about 845 million billion digits.",
			},
			{
				Prompt: "Trace this code for x = 3, n = 5. How many times does 'result *= x' run, and what is returned?\n\nresult = 1\nwhile n > 0:\n    if n % 2 == 1:\n        result *= x\n    x *= x\n    n //= 2\nreturn result",
				Choices: []string{
					"5 times; returns 243",
					"2 times; returns 243",
					"3 times; returns 243",
					"2 times; returns 729",
				},
				Answer:      1,
				Explanation: "5 = 101 in binary. Round 1: n=5 is odd, result = 3, then x = 9, n = 2. Round 2: n=2 is even, skip, x = 81, n = 1. Round 3: n=1 is odd, result = 3 * 81 = 243, x = 6561, n = 0. Two multiplications into result (one per 1 bit) and the answer 3^5 = 243.",
			},
			{
				Prompt:      "In Set Matrix Zeroes with O(1) extra space, why do you need two separate booleans, and why must the first row and first column be cleared last?",
				Explanation: "Row 0 and column 0 are used as the flags for the other columns and rows, so their original contents are overwritten by markers. The booleans record, before marking, whether row 0 and column 0 had zeros of their own; one cell (matrix[0][0]) could not hold both facts. They are cleared last because the inner cells still need to read the markers stored in them; clearing row 0 first would flag every column and zero the whole matrix.",
			},
			{
				Prompt:      "In Multiply Strings, where does the product num1[i] * num2[j] go in the result array of length m + n, and why can every column total be handled with total % 10 and total // 10?",
				Explanation: "Digit i of num1 has place value 10^(m - 1 - i) and digit j of num2 has 10^(n - 1 - j), so their product has place value 10^(m + n - 2 - i - j), which is index i + j + 1 counting from the left of an (m + n)-digit answer. Add the product to result[i + j + 1], keep total % 10 there and add total // 10 to result[i + j]. Moving 10 units of one place into 1 unit of the next place keeps the value unchanged, which is all a carry is.",
			},
			{
				Prompt: "Chapter 1 review: rotating an n x n matrix with transpose + reverse touches each cell a constant number of times. For n = 1000, about how many cell operations is that, and what is the Big O in terms of n?",
				Choices: []string{
					"About 1,000 operations: O(n)",
					"About 1,000,000 operations: O(n^2), which is linear in the number of cells",
					"About 10,000 operations: O(n log n)",
					"About 1,000,000,000 operations: O(n^3)",
				},
				Answer:      1,
				Explanation: "There are n^2 = 1,000,000 cells and each is swapped a constant number of times (once in the transpose, once in the reversal), so the work is proportional to n^2. Since the input itself has n^2 values, no algorithm can do better than reading each one; O(n^2) here means linear in the input size.",
			},
		},
	}
}
