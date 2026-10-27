package practice

// MathProblems contains all math-geometry category problems, in ladder order:
// digit arithmetic first (Plus One, Happy Number), then matrix index
// manipulation (Rotate Image, Set Matrix Zeroes, Spiral Matrix), fast power,
// long multiplication and finally coordinate geometry with counting
// (Detect Squares). Chapter 19 teaches the subgoals they share.
var MathProblems = []*Problem{
	{
		ID:              "plus-one",
		Number:          66,
		Order:           1,
		Title:           "Plus One",
		Slug:            "plus-one",
		Difficulty:      "Easy",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Math"},
		RelatedChapters: []int{19, 1},
		Description: `A non-negative whole number is stored as a list of its decimal digits, most significant digit first. For example, the number 1299 is stored as [1, 2, 9, 9]. The number has no leading zeros, except that the number zero itself is stored as [0].

Add one to the number and return the digits of the result in the same format.

The list can hold up to 100 digits, which is far more than fits in a 64-bit integer in most languages, so solve it digit by digit the way you would on paper: add 1 to the last digit and carry into the next digit to the left when a digit goes past 9.`,
		Constraints: []string{
			"1 <= digits.length <= 100",
			"0 <= digits[i] <= 9",
			"digits has no leading zeros, except for the number 0 itself ([0])",
		},
		Examples: []Example{
			{ID: "1", Input: "digits = [1,2,3]", Output: "[1,2,4]", Explanation: "123 + 1 = 124. The last digit is 3, which becomes 4 with no carry, so nothing else changes."},
			{ID: "2", Input: "digits = [1,2,9,9]", Output: "[1,3,0,0]", Explanation: "1299 + 1 = 1300. Both 9s roll over to 0 and the carry finally lands on the 2."},
			{ID: "3", Input: "digits = [9]", Output: "[1,0]", Explanation: "9 + 1 = 10: the answer has one more digit than the input."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"digits": []int{1, 2, 3}}, Expected: []int{1, 2, 4}, Note: "no carry: only the last digit changes"},
			{Input: map[string]interface{}{"digits": []int{1, 2, 9, 9}}, Expected: []int{1, 3, 0, 0}, Note: "carry ripples through two 9s"},
			{Input: map[string]interface{}{"digits": []int{9}}, Expected: []int{1, 0}, Note: "single 9: the result gains a digit"},
			{Input: map[string]interface{}{"digits": []int{4, 3, 2, 1}}, Expected: []int{4, 3, 2, 2}, Note: "longer number, no carry", Hidden: true},
			{Input: map[string]interface{}{"digits": []int{0}}, Expected: []int{1}, Note: "the number zero", Hidden: true},
			{Input: map[string]interface{}{"digits": []int{9, 9, 9}}, Expected: []int{1, 0, 0, 0}, Note: "all 9s: carry flushes out of the front", Hidden: true},
			{Input: map[string]interface{}{"digits": []int{8, 9, 9, 9}}, Expected: []int{9, 0, 0, 0}, Note: "carry stops at the very first digit, no new digit needed", Hidden: true},
			{Input: map[string]interface{}{"digits": []int{1, 0, 0, 9}}, Expected: []int{1, 0, 1, 0}, Note: "carry stops at a 0 in the middle", Hidden: true},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1) extra, except O(n) for the new list when every digit is 9",
		StarterCode:     "def plusOne(digits):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When you add 1 to 1299 on paper, which digit do you look at first, and what makes you move on to the next digit instead of stopping?"},
			{Level: 2, Type: "insight", Content: "This is digit arithmetic with a carry. Start at the lowest (last) digit. A digit below 9 just goes up by one and you are finished. A 9 becomes 0 and the carry moves one digit to the left. If the carry falls off the front, every digit was 9 and the answer is a 1 followed by zeros. See the 'Add with carry' worked example in Chapter 19.", ChapterRef: 19},
			{Level: 3, Type: "plan", Content: "1. Start at the lowest digit: loop i from the last index down to 0.\n2. Add with carry: if digits[i] is less than 9, increase it and return the list, because the carry is used up.\n3. Write the digit, keep the carry: otherwise set digits[i] to 0 and keep looping.\n4. Flush the final carry: if the loop finishes, every digit was 9, so put a 1 in front."},
			{Level: 4, Type: "skeleton", Content: "def plusOne(digits):\n    for i in range(____, -1, -1):\n        if digits[i] < 9:\n            ____\n            return digits\n        ____\n    return ____"},
		},
		Solution: Solution{
			Code: `def plusOne(digits):
    for i in range(len(digits) - 1, -1, -1):
        if digits[i] < 9:
            digits[i] += 1
            return digits
        digits[i] = 0
    return [1] + digits`,
			Explanation: "This is the first rung of the math ladder and it uses only the digit-arithmetic subgoals from Chapter 19: start at the lowest digit, add with carry, write the digit and keep the carry, and flush the final carry. Adding 1 is a special case of adding two numbers in which the only carry ever moving left is 0 or 1, so the loop can stop as soon as a digit absorbs the carry.",
			BruteForce:  "Turn the digits into one integer (value = value * 10 + d for each digit), add 1, and split it back into digits. That is O(n) and works in Python because its integers are unbounded, but in Java or C++ a 100-digit number overflows a 64-bit integer, and it also does work on every digit even when only the last one changes.",
			KeyInsight:  "A carry only continues past a digit that was 9. The first digit (from the right) that is below 9 absorbs the +1 and every digit to its left stays the same, so you can return right there.",
			WhyItWorks:  "Invariant: when the loop reaches index i, every digit to the right of i has already been written correctly (they were all 9s that became 0s) and exactly one unit of carry is waiting to be added at position i. If digits[i] < 9, adding the carry gives a digit from 1 to 9 with no new carry, so the whole number is now correct and nothing to the left changes. If digits[i] == 9, 9 + 1 = 10 writes 0 and passes the carry left, which keeps the invariant. If the carry passes index 0, the input was 99...9 = 10^n - 1, so the answer is 10^n: a 1 followed by n zeros, which is exactly [1] + digits after the zeros were written.",
			CommonMistakes: []string{
				"Forgetting the all-9s case and returning [0, 0, 0] for [9, 9, 9]. Fix: after the loop ends, prepend a 1.",
				"Adding 1 to every digit, or continuing the loop after a digit below 9 absorbed the carry. Fix: return as soon as no carry remains.",
				"Walking from the front of the list. The carry starts at the least significant digit, which is the last element. Fix: loop from len(digits) - 1 down to 0.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1) extra, except O(n) for the new list when every digit is 9",
			Walkthrough: []WalkthroughStep{
				{Title: "Start at the lowest digit", Explanation: "The +1 enters at the ones place, which is the last element, so the loop walks from the end towards the front.", CodeSnippet: "for i in range(len(digits) - 1, -1, -1):", LineStart: 2, LineEnd: 2},
				{Title: "Add with carry", Explanation: "A digit from 0 to 8 can take the carry without overflowing. Once it does, no carry remains, so the number is complete and we return immediately.", CodeSnippet: "if digits[i] < 9:\n    digits[i] += 1\n    return digits", LineStart: 3, LineEnd: 5},
				{Title: "Write the digit, keep the carry", Explanation: "9 + 1 = 10: write the 0 here and carry the 1 to the next digit on the left by continuing the loop.", CodeSnippet: "digits[i] = 0", LineStart: 6, LineEnd: 6},
				{Title: "Flush the final carry", Explanation: "Reaching this line means every digit was 9 and is now 0. The carry becomes a new leading digit.", CodeSnippet: "return [1] + digits", LineStart: 7, LineEnd: 7},
			},
			Alternatives: []AltSolution{
				{
					Name: "Convert to an integer and back",
					Code: `def plusOne(digits):
    value = 0
    for d in digits:
        value = value * 10 + d
    value += 1
    return [int(ch) for ch in str(value)]`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Builds the whole number, adds 1 and splits it into digits again. It is correct in Python because integers never overflow, but it does not port to languages with fixed-size integers and it touches every digit even when only the last one changes.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can the loop return as soon as it finds a digit below 9?", Answer: "Because that digit absorbs the carry: it becomes at most 9, so it produces no new carry, and the digits to its left are never affected by a +1 that has already been used up. Everything to its right has already been set to 0."},
			{Prompt: "What changes if you must add an arbitrary single digit k (0 to 9) instead of 1?", Answer: "The carry can still only be 0 or 1, but you can no longer stop at the first digit below 9. Keep a carry variable: total = digits[i] + carry, write total % 10, set carry = total // 10, and stop when carry becomes 0. If carry is still 1 after the front, prepend it."},
		},
	},
	{
		ID:              "happy-number",
		Number:          202,
		Order:           2,
		Title:           "Happy Number",
		Slug:            "happy-number",
		Difficulty:      "Easy",
		Category:        "math-geometry",
		Tags:            []string{"Hash Table", "Math", "Two Pointers"},
		RelatedChapters: []int{19, 2, 7},
		Description: `Start with a positive integer n and repeatedly replace it with the sum of the squares of its decimal digits. For example, 19 becomes 1*1 + 9*9 = 82.

If this process eventually reaches 1, the number is called happy (once at 1 it stays at 1, since 1*1 = 1). If it never reaches 1, it must fall into a loop of numbers that repeats forever without containing 1, and the number is not happy.

Return true if n is happy and false otherwise.`,
		Constraints: []string{
			"1 <= n <= 2^31 - 1",
		},
		Examples: []Example{
			{ID: "1", Input: "n = 19", Output: "true", Explanation: "19 -> 1+81 = 82 -> 64+4 = 68 -> 36+64 = 100 -> 1+0+0 = 1. The process reaches 1."},
			{ID: "2", Input: "n = 2", Output: "false", Explanation: "2 -> 4 -> 16 -> 37 -> 58 -> 89 -> 145 -> 42 -> 20 -> 4, and 4 has been seen before, so the loop repeats forever without 1."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 19}, Expected: true, Note: "happy: reaches 1 in four steps"},
			{Input: map[string]interface{}{"n": 2}, Expected: false, Note: "unhappy: falls into the 4 -> 16 -> ... -> 4 loop"},
			{Input: map[string]interface{}{"n": 1}, Expected: true, Note: "already 1 before any step"},
			{Input: map[string]interface{}{"n": 7}, Expected: true, Note: "happy, but needs several steps (7 -> 49 -> 97 -> 130 -> 10 -> 1)", Hidden: true},
			{Input: map[string]interface{}{"n": 4}, Expected: false, Note: "starts inside the unhappy loop", Hidden: true},
			{Input: map[string]interface{}{"n": 100}, Expected: true, Note: "zeros contribute nothing: 100 -> 1", Hidden: true},
			{Input: map[string]interface{}{"n": 2147483647}, Expected: false, Note: "largest allowed input", Hidden: true},
			{Input: map[string]interface{}{"n": 1111111}, Expected: true, Note: "seven 1s -> 7, which is happy", Hidden: true},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(log n)",
		StarterCode:     "def isHappy(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If the sequence never reaches 1, it keeps producing numbers forever. What is the only way an infinite sequence of numbers that quickly become small can avoid reaching 1, and how could you notice that it has happened?"},
			{Level: 2, Type: "insight", Content: "Two separate jobs. First, peel off the lowest digit with num % 10 and drop it with num //= 10 to compute the next number (Chapter 19, digit arithmetic). Second, after one step any number below 2^31 becomes at most 10 * 81 = 810, so an unhappy number must eventually repeat a value. Remember every value in a set (Chapter 2); seeing a value twice means you are in a loop without 1.", ChapterRef: 19},
			{Level: 3, Type: "plan", Content: "1. Peel off the lowest digit: write a helper that sums digit * digit while num > 0, using num % 10 and num //= 10.\n2. Remember what you've seen: keep a set of values already visited.\n3. Step to the next number: while n is not 1 and not in the set, add n to the set and replace n with the helper's result.\n4. Decide: the loop stops either at 1 (happy) or at a repeat (unhappy), so return n == 1."},
			{Level: 4, Type: "skeleton", Content: "def isHappy(n):\n    def next_number(num):\n        total = 0\n        while num > 0:\n            digit = ____\n            total += digit * digit\n            num = ____\n        return total\n\n    seen = set()\n    while n != 1 and ____:\n        seen.add(n)\n        n = ____\n    return n == 1"},
		},
		Solution: Solution{
			Code: `def isHappy(n):
    def next_number(num):
        total = 0
        while num > 0:
            digit = num % 10
            total += digit * digit
            num //= 10
        return total

    seen = set()
    while n != 1 and n not in seen:
        seen.add(n)
        n = next_number(n)
    return n == 1`,
			Explanation: "Builds on Plus One: there you walked a digit list from the lowest digit; here the digits are inside an integer, so you peel off the lowest digit with % 10 and drop it with // 10. The new idea is termination: a hash set (Chapter 2) of values already seen detects the loop that unhappy numbers fall into.",
			BruteForce:  "Iterate a fixed large number of times (say 1,000) and answer 'happy' if 1 ever appears. It usually gives the right answer, but the bound is a guess with no argument behind it, and it wastes hundreds of steps on unhappy numbers whose loop is only 8 numbers long.",
			KeyInsight:  "The values shrink fast: a number with d digits maps to at most 81 * d, so after one step everything is at most 810, and every number up to 810 maps to at most 243, so the sequence never leaves 1..810. A sequence trapped in a finite range must repeat a value, and the first repeat proves an endless loop. A set notices that repeat the moment it happens.",
			WhyItWorks:  "Every number below 2^31 has at most 10 digits, so its next value is at most 10 * 9 * 9 = 810, and any number up to 999 maps to at most 3 * 81 = 243. So after one step the sequence lives inside a finite set of values, and a sequence of infinitely many terms in a finite set must repeat some value. Because each value determines the next one, a repeat means the sequence cycles forever from then on. The loop therefore stops in one of two ways: n == 1 (happy) or n already in seen (a cycle; if 1 were in that cycle we would have stopped at 1 first). The set holds at most a few hundred values, and each step costs one pass over the digits.",
			CommonMistakes: []string{
				"Looping 'while n != 1' with no cycle check, which never ends for an unhappy number like 2. Fix: also stop when n has been seen before.",
				"Computing the digit sum with num / 10, which gives a float in Python 3 and never reaches 0 cleanly. Fix: use integer division num //= 10.",
				"Adding n to the set after computing the next value, so the starting number is never recorded. Fix: add n to the set first, then step.",
			},
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(log n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Peel off the lowest digit", Explanation: "num % 10 is the ones digit and num //= 10 removes it, so the loop visits every digit from right to left and adds its square.", CodeSnippet: "digit = num % 10\ntotal += digit * digit\nnum //= 10", LineStart: 4, LineEnd: 7},
				{Title: "Remember what you've seen", Explanation: "Each value visited goes into a set, so a repeat is detected in O(1).", CodeSnippet: "seen = set()\nwhile n != 1 and n not in seen:\n    seen.add(n)", LineStart: 10, LineEnd: 12},
				{Title: "Step to the next number", Explanation: "Replace n with the sum of the squares of its digits.", CodeSnippet: "n = next_number(n)", LineStart: 13, LineEnd: 13},
				{Title: "Decide", Explanation: "The loop ended either at 1 or at a repeated value, and only the first means happy.", CodeSnippet: "return n == 1", LineStart: 14, LineEnd: 14},
			},
			Alternatives: []AltSolution{
				{
					Name: "Floyd's cycle detection (constant space)",
					Code: `def isHappy(n):
    def next_number(num):
        return sum(int(ch) ** 2 for ch in str(num))

    slow, fast = n, next_number(n)
    while fast != 1 and slow != fast:
        slow = next_number(slow)
        fast = next_number(next_number(fast))
    return fast == 1`,
					TimeComplexity:  "O(log n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Treat the sequence as a linked list where next_number is the next pointer and use the tortoise and hare from Chapter 7: slow moves one step, fast moves two. If there is a loop without 1 they must meet inside it; if the sequence reaches 1, fast gets there first. No set is needed.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why are we sure the sequence for an unhappy number repeats instead of growing forever?", Answer: "Because the next value is bounded by 81 times the number of digits. A 10-digit number maps to at most 810, and every number up to 999 maps to at most 243, so after the first step the sequence never leaves the range 1..810. Infinitely many steps inside a finite range must revisit some value, and from then on it cycles."},
			{Prompt: "How does the Floyd alternative relate to Linked List Cycle, and what does it save?", Answer: "The sequence n, f(n), f(f(n)), ... is a linked list whose next pointer is f. Detecting the unhappy loop is exactly detecting a cycle in that list, so the fast and slow pointers meet inside the loop. It saves the set, so extra space drops from O(log n) to O(1)."},
		},
	},
	{
		ID:              "rotate-image",
		Number:          48,
		Order:           3,
		Title:           "Rotate Image",
		Slug:            "rotate-image",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Math", "Matrix"},
		RelatedChapters: []int{19},
		Description: `You are given an n x n grid of integers (a square matrix, as a list of rows) that represents an image. Rotate the image 90 degrees clockwise.

The rotation must happen in place: change the given matrix directly and return nothing. Do not build a second n x n matrix and copy it back.

After rotating, the first row of the result is the first column of the original read from bottom to top, the second row is the second column read from bottom to top, and so on.`,
		Constraints: []string{
			"n == matrix.length == matrix[i].length",
			"1 <= n <= 20",
			"-1000 <= matrix[i][j] <= 1000",
		},
		Examples: []Example{
			{ID: "1", Input: "matrix = [[1,2,3],[4,5,6],[7,8,9]]", Output: "[[7,4,1],[8,5,2],[9,6,3]]", Explanation: "The left column 1,4,7 read bottom to top (7,4,1) becomes the top row."},
			{ID: "2", Input: "matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]", Output: "[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, Expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, Note: "3x3 example"},
			{Input: map[string]interface{}{"matrix": [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}}, Expected: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}, Note: "4x4: even size, two layers"},
			{Input: map[string]interface{}{"matrix": [][]int{{1}}}, Expected: [][]int{{1}}, Note: "1x1: nothing moves", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2}, {3, 4}}}, Expected: [][]int{{3, 1}, {4, 2}}, Note: "2x2: catches a counterclockwise rotation (which would give [[2,4],[1,3]])", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}}, Expected: [][]int{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}}, Note: "5x5: odd size with a fixed centre; catches swapping each pair twice", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{-1, -2}, {-1000, 1000}}}, Expected: [][]int{{-1000, -1}, {1000, -2}}, Note: "negative values and the constraint extremes", Hidden: true},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def rotate(matrix):\n    # Modify matrix in place; do not return anything\n    pass",
		Harness:         Harness{InPlace: "matrix"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Take the cell at row r, column c. After a clockwise quarter turn, which row and column does it land in? Try it on the 3x3 example with the value 2, which starts at (0, 1)."},
			{Level: 2, Type: "insight", Content: "Map the indices: (r, c) moves to (c, n - 1 - r). That move is two simpler in-place moves done one after the other. A transpose sends (r, c) to (c, r), and reversing each row then sends column r to column n - 1 - r. Both are swaps, so no second matrix is needed. Chapter 19 works the counterclockwise version step by step.", ChapterRef: 19},
			{Level: 3, Type: "plan", Content: "1. Map the indices: confirm on paper that (r, c) must end up at (c, n - 1 - r).\n2. Transpose across the diagonal: for each pair with r < c, swap matrix[r][c] and matrix[c][r]. Only visit the upper triangle, or every swap gets undone.\n3. Reverse to finish: reverse every row."},
			{Level: 4, Type: "skeleton", Content: "def rotate(matrix):\n    n = len(matrix)\n    for r in range(n):\n        for c in range(____, n):\n            matrix[r][c], matrix[c][r] = ____\n    for row in matrix:\n        ____"},
		},
		Solution: Solution{
			Code: `def rotate(matrix):
    n = len(matrix)
    for r in range(n):
        for c in range(r + 1, n):
            matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
    for row in matrix:
        row.reverse()`,
			Explanation: "The first matrix problem on the ladder. Instead of the digits of Plus One and Happy Number, you now reason about where each index goes. A clockwise rotation maps (r, c) to (c, n - 1 - r), and you get there with two moves that are easy to do in place: transpose across the main diagonal, then reverse each row. These are the 'Map the indices', 'Transpose across the diagonal' and 'Reverse to finish' subgoals from Chapter 19.",
			BruteForce:  "Build a new n x n matrix with rotated[c][n - 1 - r] = matrix[r][c] and copy it back. That is O(n^2) time, which is already optimal, but it uses O(n^2) extra memory, which the problem forbids.",
			KeyInsight:  "A rotation is a composition of two reflections. Transposing reflects across the main diagonal ((r, c) -> (c, r)), and reversing each row reflects left-right ((r, c) -> (r, n - 1 - c)). Done in that order they give (r, c) -> (c, n - 1 - r), which is the clockwise rotation, and each reflection is a series of swaps.",
			WhyItWorks:  "Follow one value starting at (r, c). The transpose swaps it to (c, r); because the loop only visits pairs with c > r, each off-diagonal pair is swapped exactly once and the diagonal stays put. Reversing row c then moves column r to column n - 1 - r, so the value ends at (c, n - 1 - r), which is exactly where a clockwise quarter turn puts it: the old row index becomes the distance from the right edge, and the old column index becomes the new row. Every value follows the same two moves, so the whole matrix is rotated, and only a constant number of temporaries is ever used.",
			CommonMistakes: []string{
				"Looping c over the full range 0..n-1 during the transpose, so each pair is swapped twice and the matrix ends up unchanged. Fix: start the inner loop at r + 1.",
				"Reversing the order of the rows instead of reversing each row, which gives a counterclockwise rotation. Fix: transpose then reverse each row for clockwise. Reversing the row order after the transpose is the counterclockwise version.",
				"Writing matrix = rotated_copy inside the function, which only rebinds the local name and leaves the caller's matrix unchanged. Fix: change the cells (or use matrix[:] = ...) so the original object is modified.",
			},
			TimeComplexity:  "O(n^2)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Map the indices", Explanation: "Work out the target before touching code: a clockwise turn sends (r, c) to (c, n - 1 - r). For n = 3, the value 2 at (0, 1) must land at (1, 2).", CodeSnippet: "n = len(matrix)", LineStart: 1, LineEnd: 2},
				{Title: "Transpose across the diagonal", Explanation: "Swap each pair above the diagonal with its mirror image below it. Starting c at r + 1 swaps every pair exactly once. After this, the value 2 sits at (1, 0).", CodeSnippet: "for r in range(n):\n    for c in range(r + 1, n):\n        matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]", LineStart: 3, LineEnd: 5},
				{Title: "Reverse to finish", Explanation: "Reversing each row sends column j to column n - 1 - j, so 2 moves from (1, 0) to (1, 2), its rotated position.", CodeSnippet: "for row in matrix:\n    row.reverse()", LineStart: 6, LineEnd: 7},
			},
			Alternatives: []AltSolution{
				{
					Name: "Copy into a rotated matrix (uses O(n^2) extra space)",
					Code: `def rotate(matrix):
    n = len(matrix)
    rotated = [[0] * n for _ in range(n)]
    for r in range(n):
        for c in range(n):
            rotated[c][n - 1 - r] = matrix[r][c]
    for r in range(n):
        for c in range(n):
            matrix[r][c] = rotated[r][c]`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n^2)",
					Explanation:     "Applies the index map directly into a second matrix and copies it back. It is the easiest way to check the formula (r, c) -> (c, n - 1 - r), but it breaks the in-place rule.",
				},
				{
					Name: "Rotate four cells at a time, layer by layer",
					Code: `def rotate(matrix):
    n = len(matrix)
    for layer in range(n // 2):
        first, last = layer, n - 1 - layer
        for i in range(first, last):
            offset = i - first
            top = matrix[first][i]
            matrix[first][i] = matrix[last - offset][first]
            matrix[last - offset][first] = matrix[last][last - offset]
            matrix[last][last - offset] = matrix[i][last]
            matrix[i][last] = top`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Treat the matrix as rings (layers). Within a ring, the four cells that rotate into each other form a cycle, so save the top one and move left -> top, bottom -> left, right -> bottom, saved top -> right. Each cell moves exactly once, but the indices are easy to get wrong.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the transpose loop start c at r + 1 rather than 0?", Answer: "Because swapping (r, c) with (c, r) and later swapping (c, r) with (r, c) undoes the first swap. Starting at r + 1 visits each off-diagonal pair once, from the upper triangle, and leaves the diagonal (where r == c) alone since those cells do not move."},
			{Prompt: "How would you rotate 90 degrees counterclockwise in place, and how would you rotate 180 degrees?", Answer: "Counterclockwise sends (r, c) to (n - 1 - c, r): transpose, then reverse the order of the rows (equivalently reverse each row first, then transpose). A 180-degree turn sends (r, c) to (n - 1 - r, n - 1 - c): reverse the order of the rows and also reverse each row, with no transpose."},
		},
	},
	{
		ID:              "set-matrix-zeroes",
		Number:          73,
		Order:           4,
		Title:           "Set Matrix Zeroes",
		Slug:            "set-matrix-zeroes",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Hash Table", "Matrix"},
		RelatedChapters: []int{19, 2},
		Description: `You are given an m x n matrix of integers. For every cell that contains 0 in the original matrix, set every cell in its row and every cell in its column to 0.

Do it in place: modify the given matrix and return nothing. Only zeros that were in the original matrix spread; a cell that becomes 0 because of this process does not wipe out its own row and column.

Follow-up: an O(m * n) extra-space copy is easy, and O(m + n) extra space is a good first answer, but can you use only O(1) extra space?`,
		Constraints: []string{
			"m == matrix.length",
			"n == matrix[0].length",
			"1 <= m, n <= 200",
			"-2^31 <= matrix[i][j] <= 2^31 - 1",
		},
		Examples: []Example{
			{ID: "1", Input: "matrix = [[1,1,1],[1,0,1],[1,1,1]]", Output: "[[1,0,1],[0,0,0],[1,0,1]]", Explanation: "The only zero is at row 1, column 1, so row 1 and column 1 become all zeros."},
			{ID: "2", Input: "matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]", Output: "[[0,0,0,0],[0,4,5,0],[0,3,1,0]]", Explanation: "Zeros in columns 0 and 3 of row 0 clear row 0 and both of those columns."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}, Expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, Note: "one zero in the middle"},
			{Input: map[string]interface{}{"matrix": [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}}, Expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}, Note: "zeros in the first row, which is also the marker row"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2}, {3, 4}}}, Expected: [][]int{{1, 2}, {3, 4}}, Note: "no zeros: nothing changes", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{0}}}, Expected: [][]int{{0}}, Note: "1x1 zero", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 0, 3}}}, Expected: [][]int{{0, 0, 0}}, Note: "single row", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1}, {0}, {3}}}, Expected: [][]int{{0}, {0}, {0}}, Note: "single column: the first column itself must be cleared", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}}, Expected: [][]int{{0, 0, 3, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}, Note: "zero in the first column but not the first row: the 3 must survive", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}}, Expected: [][]int{{1, 2, 0}, {4, 5, 0}, {0, 0, 0}}, Note: "zero in a corner away from the marker row and column", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{-1, 2, 3}, {4, 0, -6}, {2147483647, 8, -2147483648}}}, Expected: [][]int{{-1, 0, 3}, {0, 0, 0}, {2147483647, 0, -2147483648}}, Note: "negative and extreme values are not zeros", Hidden: true},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def setZeroes(matrix):\n    # Modify matrix in place; do not return anything\n    pass",
		Harness:         Harness{InPlace: "matrix"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you zero a row the moment you see a 0, later cells in that row look like original zeros and spread further. What is the smallest amount you need to remember, per row and per column, before you change anything?"},
			{Level: 2, Type: "insight", Content: "You only need one yes/no flag per row and one per column. A set of row indices and a set of column indices works in O(m + n) space. For O(1) space, use the matrix itself as storage: row 0 holds the column flags and column 0 holds the row flags. Because row 0 and column 0 are overwritten by markers, record first whether they had a zero of their own. See 'Using the input as storage' in Chapter 19.", ChapterRef: 19},
			{Level: 3, Type: "plan", Content: "1. Save the first row and column's own state: two booleans saying whether row 0 or column 0 contains a 0.\n2. Mark in the first row and column: for every inner cell (r >= 1, c >= 1) that is 0, set matrix[r][0] and matrix[0][c] to 0.\n3. Apply the markers: for every inner cell, set it to 0 if its row marker or column marker is 0.\n4. Fix the first row and column last: clear row 0 and/or column 0 if their saved booleans say so."},
			{Level: 4, Type: "skeleton", Content: "def setZeroes(matrix):\n    rows, cols = len(matrix), len(matrix[0])\n    first_row_has_zero = any(matrix[0][c] == 0 for c in range(cols))\n    first_col_has_zero = ____\n\n    for r in range(1, rows):\n        for c in range(1, cols):\n            if matrix[r][c] == 0:\n                ____\n                ____\n\n    for r in range(1, rows):\n        for c in range(1, cols):\n            if ____:\n                matrix[r][c] = 0\n\n    if first_row_has_zero:\n        for c in range(cols):\n            matrix[0][c] = 0\n    if first_col_has_zero:\n        ____"},
		},
		Solution: Solution{
			Code: `def setZeroes(matrix):
    rows, cols = len(matrix), len(matrix[0])
    first_row_has_zero = any(matrix[0][c] == 0 for c in range(cols))
    first_col_has_zero = any(matrix[r][0] == 0 for r in range(rows))

    for r in range(1, rows):
        for c in range(1, cols):
            if matrix[r][c] == 0:
                matrix[r][0] = 0
                matrix[0][c] = 0

    for r in range(1, rows):
        for c in range(1, cols):
            if matrix[r][0] == 0 or matrix[0][c] == 0:
                matrix[r][c] = 0

    if first_row_has_zero:
        for c in range(cols):
            matrix[0][c] = 0
    if first_col_has_zero:
        for r in range(rows):
            matrix[r][0] = 0`,
			Explanation: "Builds on Rotate Image, which also modified a matrix in place by careful index work. The new idea is using the input as storage: the first row and first column become the 'which columns / which rows must be cleared' flags, so the O(m + n) sets of the simple solution disappear. Two booleans protect the information that the markers overwrite.",
			BruteForce:  "Copy the matrix, then for every 0 in the copy clear its row and column in the original: O(m * n * (m + n)) time and O(m * n) space. Recording zero rows and zero columns in two sets first fixes the time (O(m * n)) but still needs O(m + n) memory.",
			KeyInsight:  "Each row needs one bit ('clear me') and so does each column. Row 0 has exactly one cell per column and column 0 has exactly one cell per row, so they can hold those bits, and writing 0 into them is harmless because those cells will be cleared anyway if their marker is set.",
			WhyItWorks:  "After step 2, for every inner row r >= 1, matrix[r][0] == 0 exactly when row r originally had a zero (in column 0 itself, or in an inner cell that wrote the marker). The same holds for every column c >= 1 and matrix[0][c]. Writing a marker never destroys needed information: matrix[r][0] is set to 0 only when row r must be cleared anyway. Step 3 clears precisely the inner cells whose row or column had an original zero. Row 0 and column 0 are processed last, using booleans saved before any marker was written, because their cells now mix marker bits with original values; if we cleared them first, every column (or row) would look marked.",
			CommonMistakes: []string{
				"Zeroing a row and column immediately when you find a 0, so the new zeros are later mistaken for original ones and the whole matrix collapses to zeros. Fix: record first, modify afterwards.",
				"Using row 0 and column 0 as markers without first recording whether they contained a 0 themselves, which either misses clearing them or clears them wrongly. Fix: save first_row_has_zero and first_col_has_zero before marking.",
				"Clearing row 0 or column 0 before applying the markers to the inner cells, which erases the markers. Fix: fix the first row and column last.",
			},
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Save the first row and column's own state", Explanation: "Row 0 and column 0 are about to be overwritten with markers, so first record whether they contained an original zero.", CodeSnippet: "first_row_has_zero = any(matrix[0][c] == 0 for c in range(cols))\nfirst_col_has_zero = any(matrix[r][0] == 0 for r in range(rows))", LineStart: 3, LineEnd: 4},
				{Title: "Mark in the first row and column", Explanation: "An inner zero at (r, c) flags its row in column 0 and its column in row 0.", CodeSnippet: "if matrix[r][c] == 0:\n    matrix[r][0] = 0\n    matrix[0][c] = 0", LineStart: 6, LineEnd: 10},
				{Title: "Apply the markers", Explanation: "An inner cell is cleared when its row or its column is flagged. Row 0 and column 0 are skipped because they still hold the flags.", CodeSnippet: "if matrix[r][0] == 0 or matrix[0][c] == 0:\n    matrix[r][c] = 0", LineStart: 12, LineEnd: 15},
				{Title: "Fix the first row and column last", Explanation: "Now the markers are no longer needed, so clear row 0 and column 0 according to the booleans saved at the start.", CodeSnippet: "if first_row_has_zero:\nif first_col_has_zero:", LineStart: 17, LineEnd: 22},
			},
			Alternatives: []AltSolution{
				{
					Name: "Record zero rows and columns in sets",
					Code: `def setZeroes(matrix):
    zero_rows, zero_cols = set(), set()
    for r, row in enumerate(matrix):
        for c, value in enumerate(row):
            if value == 0:
                zero_rows.add(r)
                zero_cols.add(c)
    for r, row in enumerate(matrix):
        for c in range(len(row)):
            if r in zero_rows or c in zero_cols:
                row[c] = 0`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(m + n)",
					Explanation:     "The same two phases (record, then modify) with the flags kept in hash sets (Chapter 2) instead of inside the matrix. It is the clearest version and a fine interview answer before the O(1) follow-up.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must row 0 and column 0 be cleared last rather than first?", Answer: "Because they hold the markers for every other row and column. Clearing row 0 first would set every column marker to 0, and the next phase would wipe the whole matrix. They can be cleared only after the inner cells have read their markers."},
			{Prompt: "matrix[0][0] sits in both the first row and the first column. Why does the solution use two separate booleans instead of relying on matrix[0][0]?", Answer: "One cell cannot answer two questions. If matrix[0][0] were used as a marker it could not tell 'row 0 must be cleared' apart from 'column 0 must be cleared'. The two booleans record each fact separately before any marker is written, and the marker loops skip index 0 so matrix[0][0] is never used as a flag."},
		},
	},
	{
		ID:              "spiral-matrix",
		Number:          54,
		Order:           5,
		Title:           "Spiral Matrix",
		Slug:            "spiral-matrix",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Matrix", "Simulation"},
		RelatedChapters: []int{19},
		Description: `You are given an m x n matrix (m rows, n columns; it does not have to be square). Return all of its elements in clockwise spiral order.

Spiral order starts at the top-left cell, goes right along the top row, then down the right column, then left along the bottom row, then up the left column, and then repeats the same pattern on the remaining inner rectangle until every cell has been listed exactly once.`,
		Constraints: []string{
			"m == matrix.length",
			"n == matrix[i].length",
			"1 <= m, n <= 10",
			"-100 <= matrix[i][j] <= 100",
		},
		Examples: []Example{
			{ID: "1", Input: "matrix = [[1,2,3],[4,5,6],[7,8,9]]", Output: "[1,2,3,6,9,8,7,4,5]", Explanation: "Outer ring: right 1,2,3; down 6,9; left 8,7; up 4. Then the inner 1x1 rectangle: 5."},
			{ID: "2", Input: "matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]", Output: "[1,2,3,4,8,12,11,10,9,5,6,7]", Explanation: "After the outer ring, only the middle row 6,7 is left; it is walked rightwards once and must not be walked back."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, Expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, Note: "square, odd size: ends on the centre cell"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}, Expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}, Note: "non-square (3x4): a single middle row is left over"},
			{Input: map[string]interface{}{"matrix": [][]int{{7}}}, Expected: []int{7}, Note: "1x1", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3, 4}}}, Expected: []int{1, 2, 3, 4}, Note: "single row: catches walking it back leftwards", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1}, {2}, {3}}}, Expected: []int{1, 2, 3}, Note: "single column: catches walking it back upwards", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}}, Expected: []int{1, 2, 4, 6, 8, 7, 5, 3}, Note: "tall and narrow (4x2)", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}}, Expected: []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8}, Note: "tall (4x3): a single middle column is left over", Hidden: true},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}}, Expected: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}, Note: "4x4: two full rings", Hidden: true},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def spiralOrder(matrix):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "After you walk the whole top row, that row is used up. What simple numbers could describe 'the part of the matrix not visited yet', and how does each of them change after each side?"},
			{Level: 2, Type: "insight", Content: "The unvisited part is always a rectangle, described by four boundaries: top, bottom, left, right. Walk one side of it and then shrink the matching boundary (after the top row, top += 1). Before walking the bottom row leftwards or the left column upwards, check that the rectangle still has a row or column left, otherwise a single leftover row or column is listed twice. This is the 'layer by layer boundaries' idea from Chapter 19.", ChapterRef: 19},
			{Level: 3, Type: "plan", Content: "1. Set the boundaries: top = 0, bottom = m - 1, left = 0, right = n - 1.\n2. Walk one side, then shrink: while top <= bottom and left <= right, walk the top row left to right and do top += 1, then walk the right column top to bottom and do right -= 1.\n3. Check before the return trip: if top <= bottom, walk the bottom row right to left and do bottom -= 1; if left <= right, walk the left column bottom to top and do left += 1.\n4. Return the collected list when the boundaries cross."},
			{Level: 4, Type: "skeleton", Content: "def spiralOrder(matrix):\n    result = []\n    top, bottom = 0, len(matrix) - 1\n    left, right = 0, len(matrix[0]) - 1\n\n    while top <= bottom and left <= right:\n        for c in range(left, right + 1):\n            result.append(matrix[top][c])\n        ____\n        for r in range(top, bottom + 1):\n            result.append(matrix[r][right])\n        ____\n        if ____:\n            for c in range(right, left - 1, -1):\n                result.append(matrix[bottom][c])\n            bottom -= 1\n        if ____:\n            for r in range(bottom, top - 1, -1):\n                result.append(matrix[r][left])\n            left += 1\n    return result"},
		},
		Solution: Solution{
			Code: `def spiralOrder(matrix):
    result = []
    top, bottom = 0, len(matrix) - 1
    left, right = 0, len(matrix[0]) - 1

    while top <= bottom and left <= right:
        for c in range(left, right + 1):
            result.append(matrix[top][c])
        top += 1
        for r in range(top, bottom + 1):
            result.append(matrix[r][right])
        right -= 1
        if top <= bottom:
            for c in range(right, left - 1, -1):
                result.append(matrix[bottom][c])
            bottom -= 1
        if left <= right:
            for r in range(bottom, top - 1, -1):
                result.append(matrix[r][left])
            left += 1
    return result`,
			Explanation: "Builds on Rotate Image: there the four-way-swap alternative already treated the matrix as rings (layers). Here you walk those rings and keep four boundaries for the unvisited rectangle. The new twist is that the matrix need not be square, so the last layer can be a single row or column, and the 'Check before the return trip' subgoal stops it from being listed twice.",
			BruteForce:  "Simulate a walker with a direction and a visited set: move forward, and turn right when the next cell is outside the matrix or already visited. It is O(m * n) time like the boundary version but needs O(m * n) extra space for the visited set and a bounds-plus-visited check on every step.",
			KeyInsight:  "The unvisited cells always form a rectangle. Four integers describe it, and each side you finish removes exactly one row or column from it, so you never need to remember individual cells.",
			WhyItWorks:  "Invariant: at the top of each loop iteration, result holds exactly the cells outside the rectangle rows top..bottom, columns left..right, in spiral order, and every cell inside that rectangle is unvisited. Walking the top row and doing top += 1 removes that row from the rectangle and keeps the invariant; the same holds for the right column. After those two moves the rectangle may have no rows left (top > bottom) or no columns left (left > right); the two guards skip the return trips in that case, which is exactly when the remaining part was a single row or column that has already been listed. Every iteration shrinks the rectangle, so the loop ends after visiting each of the m * n cells exactly once.",
			CommonMistakes: []string{
				"Leaving out the 'if top <= bottom' and 'if left <= right' checks, which lists the middle row of a 3x4 matrix twice ([..., 6, 7, 6] instead of [..., 6, 7]). Fix: check the boundaries again before walking back.",
				"Shrinking a boundary before walking its side, or walking the right column starting from the old top, which repeats the corner cell. Fix: walk a side, then shrink, and start the next side from the updated boundary.",
				"Assuming a square matrix and looping n // 2 layers, which fails for 3x4 or 4x2 inputs. Fix: loop while top <= bottom and left <= right.",
			},
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(1) extra besides the output list",
			Walkthrough: []WalkthroughStep{
				{Title: "Set the boundaries", Explanation: "The unvisited rectangle starts as the whole matrix: rows 0..m-1, columns 0..n-1.", CodeSnippet: "top, bottom = 0, len(matrix) - 1\nleft, right = 0, len(matrix[0]) - 1", LineStart: 3, LineEnd: 4},
				{Title: "Walk one side, then shrink", Explanation: "List the top row, then drop it with top += 1. List the right column (starting at the new top so the corner is not repeated), then drop it with right -= 1.", CodeSnippet: "for c in range(left, right + 1):\n    result.append(matrix[top][c])\ntop += 1\nright -= 1", LineStart: 6, LineEnd: 12},
				{Title: "Check before the return trip", Explanation: "The rectangle may already be empty in one direction. Only walk the bottom row back if a row remains, and only walk the left column up if a column remains.", CodeSnippet: "if top <= bottom:\nbottom -= 1\nif left <= right:\nleft += 1", LineStart: 13, LineEnd: 20},
				{Title: "Return when the boundaries cross", Explanation: "When top > bottom or left > right, every cell has been listed once.", CodeSnippet: "return result", LineStart: 21, LineEnd: 21},
			},
			Alternatives: []AltSolution{
				{
					Name: "Simulate a walker with a visited set",
					Code: `def spiralOrder(matrix):
    rows, cols = len(matrix), len(matrix[0])
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    visited = set()
    r = c = d = 0
    result = []
    for _ in range(rows * cols):
        result.append(matrix[r][c])
        visited.add((r, c))
        nr, nc = r + directions[d][0], c + directions[d][1]
        if not (0 <= nr < rows and 0 <= nc < cols) or (nr, nc) in visited:
            d = (d + 1) % 4
            nr, nc = r + directions[d][0], c + directions[d][1]
        r, c = nr, nc
    return result`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "Walks cell by cell in the current direction and turns right (right, down, left, up, repeat) when the next cell is off the grid or already visited. It needs no special cases for leftover rows, but pays for a visited set.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "For matrix [[1,2,3,4]] (one row), what goes wrong without the 'if top <= bottom' check?", Answer: "After walking the top row, top becomes 1 while bottom is 0, and right becomes 2. Without the check the loop would walk 'the bottom row' (row 0 again) from column 2 down to 0 and add 3, 2, 1 a second time, giving [1,2,3,4,3,2,1]."},
			{Prompt: "How would you produce the spiral in counterclockwise order (down the left column first)?", Answer: "Keep the same four boundaries but change the order of sides: walk the left column top to bottom (left += 1), the bottom row left to right (bottom -= 1), then, with the same guards, the right column bottom to top (right -= 1) and the top row right to left (top += 1)."},
		},
	},
	{
		ID:              "powx-n",
		Number:          50,
		Order:           6,
		Title:           "Pow(x, n)",
		Slug:            "powx-n",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Math", "Recursion"},
		RelatedChapters: []int{19, 8, 20},
		Description: `Compute x raised to the power n (x^n), where x is a floating-point number and n is an integer that may be zero or negative. Remember that x^0 = 1 and x^(-n) = 1 / x^n.

n can be as large as about 2 billion in absolute value, so multiplying x by itself n times is far too slow. Do not use the built-in power operator or math.pow.

Answers within 10^-5 of the exact value are accepted.`,
		Constraints: []string{
			"-100.0 < x < 100.0",
			"-2^31 <= n <= 2^31 - 1",
			"n is an integer",
			"Either x is not zero or n > 0",
			"-10^4 <= x^n <= 10^4",
		},
		Examples: []Example{
			{ID: "1", Input: "x = 2.00000, n = 10", Output: "1024.00000", Explanation: "10 is 1010 in binary, so 2^10 = 2^8 * 2^2 = 256 * 4 = 1024."},
			{ID: "2", Input: "x = 2.10000, n = 3", Output: "9.26100"},
			{ID: "3", Input: "x = 2.00000, n = -2", Output: "0.25000", Explanation: "2^-2 = 1 / 2^2 = 1/4."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"x": 2.0, "n": 10}, Expected: 1024.0, Note: "even exponent"},
			{Input: map[string]interface{}{"x": 2.1, "n": 3}, Expected: 9.261000000000001, Note: "fractional base, odd exponent"},
			{Input: map[string]interface{}{"x": 2.0, "n": -2}, Expected: 0.25, Note: "negative exponent"},
			{Input: map[string]interface{}{"x": 5.0, "n": 0}, Expected: 1.0, Note: "n = 0 gives 1", Hidden: true},
			{Input: map[string]interface{}{"x": 2.0, "n": -3}, Expected: 0.125, Note: "odd negative exponent", Hidden: true},
			{Input: map[string]interface{}{"x": -2.0, "n": 3}, Expected: -8.0, Note: "negative base, odd exponent keeps the sign", Hidden: true},
			{Input: map[string]interface{}{"x": 1.0, "n": 2147483647}, Expected: 1.0, Note: "largest n: a linear loop times out", Hidden: true},
			{Input: map[string]interface{}{"x": -1.0, "n": -2147483648}, Expected: 1.0, Note: "smallest n: -n does not fit in 32 bits in other languages", Hidden: true},
			{Input: map[string]interface{}{"x": 0.5, "n": 13}, Expected: 0.0001220703125, Note: "13 = 1101 in binary: three bits set", Hidden: true},
			{Input: map[string]interface{}{"x": 1.00001, "n": 123456}, Expected: 3.436844752098875, Note: "large exponent with a base close to 1", Hidden: true},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def myPow(x, n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you already know x^5, how many multiplications does it take to get x^10? And knowing x, x^2, x^4 and x^8, how would you build x^13 from them?"},
			{Level: 2, Type: "insight", Content: "Fast exponentiation (Chapter 19). Write n in binary: 13 = 1101 = 8 + 4 + 1, so x^13 = x^8 * x^4 * x^1. Repeatedly squaring x gives x, x^2, x^4, x^8, ..., and you multiply the ones whose bit is 1 into the answer. n % 2 reads the lowest bit and n //= 2 drops it, so there are only about log2(n) rounds. Handle a negative n first by using 1/x and -n.", ChapterRef: 19},
			{Level: 3, Type: "plan", Content: "1. Handle a negative exponent: if n < 0, replace x with 1 / x and n with -n.\n2. Start result at 1.0.\n3. Read the lowest bit: while n > 0, look at n % 2.\n4. Multiply in when the bit is 1: if it is 1, multiply result by the current x.\n5. Square the base, shift the exponent: x becomes x * x (the next power x^(2^k)) and n becomes n // 2.\n6. Return result."},
			{Level: 4, Type: "skeleton", Content: "def myPow(x, n):\n    if n < 0:\n        x = ____\n        n = -n\n\n    result = 1.0\n    while n > 0:\n        if ____:\n            result *= x\n        x = ____\n        n = ____\n    return result"},
		},
		Solution: Solution{
			Code: `def myPow(x, n):
    if n < 0:
        x = 1 / x
        n = -n

    result = 1.0
    while n > 0:
        if n % 2 == 1:
            result *= x
        x *= x
        n //= 2
    return result`,
			Explanation: "Builds on Happy Number, which processed the decimal digits of an integer with % 10 and // 10. Here you process the binary digits of the exponent with % 2 and // 2: each binary digit of n says whether one of the powers x, x^2, x^4, x^8, ... belongs in the product. That turns up to 2 billion multiplications into about 31 rounds. (Chapter 20 writes n % 2 as n & 1 and n // 2 as n >> 1.)",
			BruteForce:  "Multiply result by x, |n| times. That is O(|n|): for n = 2,147,483,647 it is over 2 billion multiplications, which takes minutes in Python. The waste is recomputing the same partial products: x^8 is built from scratch out of single multiplications instead of by squaring x^4.",
			KeyInsight:  "Every exponent is a sum of powers of two (its binary digits), and x^(a + b) = x^a * x^b. The powers x^1, x^2, x^4, x^8, ... come from repeated squaring, one multiplication each, so you only need one squaring per binary digit of n plus one multiplication per 1 bit.",
			WhyItWorks:  "Let N be the original (non-negative) exponent and X the original base. Invariant at the top of each loop round: result * x^n == X^N. Initially result = 1, x = X and n = N, so it holds. If n is odd, n = 2k + 1 and x^n = x * (x^2)^k, so after result *= x, x = x * x and n = k the product result * x^n is unchanged. If n is even, n = 2k and x^n = (x^2)^k, so squaring x and halving n keeps it too. Each round halves n, so after floor(log2 N) + 1 rounds n == 0 and the invariant says result == X^N. A negative exponent is handled first because x^(-N) = (1/x)^N. Example: 3^13 with 13 = 1101: bits 1,0,1,1 from the lowest, so result picks up 3, skips 9, picks up 81 and 6561: 3 * 81 * 6561 = 1,594,323.",
			CommonMistakes: []string{
				"Using a loop that runs n times, which times out for n near 2^31. Fix: halve n each round.",
				"Forgetting negative exponents (the loop never runs and returns 1.0 for n = -2). Fix: when n < 0, invert x and negate n before the loop.",
				"Squaring x before multiplying it into result in the same round, so the wrong power is used. Fix: check the bit and multiply first, then square x and halve n.",
				"In Java or C++, writing n = -n on a 32-bit int when n = -2^31, which overflows. Fix: copy n into a 64-bit variable first (Python integers do not overflow, so this only matters when you port the code).",
			},
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle a negative exponent", Explanation: "x^(-n) = (1/x)^n, so flip the base and make the exponent positive once, before the loop.", CodeSnippet: "if n < 0:\n    x = 1 / x\n    n = -n", LineStart: 2, LineEnd: 4},
				{Title: "Read the lowest bit", Explanation: "n % 2 is the lowest binary digit of n. The loop runs once per binary digit.", CodeSnippet: "while n > 0:\n    if n % 2 == 1:", LineStart: 7, LineEnd: 8},
				{Title: "Multiply in when the bit is 1", Explanation: "At this point x holds X^(2^k) for the k-th bit. A 1 bit means that power is part of the answer.", CodeSnippet: "result *= x", LineStart: 9, LineEnd: 9},
				{Title: "Square the base, shift the exponent", Explanation: "Squaring moves x from X^(2^k) to X^(2^(k+1)), and n //= 2 drops the bit just used, so the next round reads the next bit.", CodeSnippet: "x *= x\nn //= 2", LineStart: 10, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Recursive halving",
					Code: `def myPow(x, n):
    def power(base, exp):
        if exp == 0:
            return 1.0
        half = power(base, exp // 2)
        if exp % 2 == 0:
            return half * half
        return half * half * base

    if n < 0:
        return 1 / power(x, -n)
    return power(x, n)`,
					TimeComplexity:  "O(log n)",
					SpaceComplexity: "O(log n)",
					Explanation:     "Divide and conquer (Chapter 8): x^n = (x^(n//2))^2, times one extra x when n is odd. Computing the half once and squaring it is what keeps it logarithmic; calling power(base, exp // 2) twice would bring back O(n). The recursion uses O(log n) stack frames, about 32 here.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "How many times does the loop body run for n = 1,000,000, and why?", Answer: "20 times. Each round halves n, and 1,000,000 has 20 binary digits (2^19 = 524,288 <= 1,000,000 < 2^20), so after 20 halvings n reaches 0. Compare that to one million multiplications for the naive loop."},
			{Prompt: "The same loop computes (base^exp) % m for huge exponents if you add '% m' after each multiplication. Why is it valid to take the remainder at every step?", Answer: "Because (a * b) % m == ((a % m) * (b % m)) % m. Every value in the loop is a product of earlier values, so reducing each one modulo m never changes the final remainder, and it keeps the numbers small (below m^2) instead of growing to millions of digits."},
		},
	},
	{
		ID:              "multiply-strings",
		Number:          43,
		Order:           7,
		Title:           "Multiply Strings",
		Slug:            "multiply-strings",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Math", "String", "Simulation"},
		RelatedChapters: []int{19},
		Description: `You are given two non-negative integers num1 and num2 written as strings of decimal digits. Return their product, also as a string of decimal digits.

You may not convert either whole string to a number (no int(num1)) and may not use any big-integer library. Work digit by digit, as in long multiplication on paper.

Neither input has leading zeros, except the number zero itself, which is written "0". The output must not have leading zeros either.`,
		Constraints: []string{
			"1 <= num1.length, num2.length <= 200",
			"num1 and num2 consist of digits only",
			"num1 and num2 have no leading zeros, except the number 0 itself",
		},
		Examples: []Example{
			{ID: "1", Input: `num1 = "2", num2 = "3"`, Output: `"6"`},
			{ID: "2", Input: `num1 = "123", num2 = "456"`, Output: `"56088"`, Explanation: "123 * 456 = 56088. The ones digits 3 * 6 = 18 write 8 in the ones place and carry 1 into the tens place."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"num1": "2", "num2": "3"}, Expected: "6", Note: "single digits, no carry"},
			{Input: map[string]interface{}{"num1": "123", "num2": "456"}, Expected: "56088", Note: "multi-digit with carries"},
			{Input: map[string]interface{}{"num1": "0", "num2": "52"}, Expected: "0", Note: "zero times anything is \"0\", not \"00\"", Hidden: true},
			{Input: map[string]interface{}{"num1": "9", "num2": "9"}, Expected: "81", Note: "single digits whose product carries", Hidden: true},
			{Input: map[string]interface{}{"num1": "999", "num2": "999"}, Expected: "998001", Note: "carries in every column", Hidden: true},
			{Input: map[string]interface{}{"num1": "1", "num2": "12345"}, Expected: "12345", Note: "result shorter than m + n: leading zero must be stripped", Hidden: true},
			{Input: map[string]interface{}{"num1": "123456789", "num2": "987654321"}, Expected: "121932631112635269", Note: "18-digit product", Hidden: true},
			{Input: map[string]interface{}{"num1": "100", "num2": "1000"}, Expected: "100000", Note: "trailing zeros must be kept", Hidden: true},
			{Input: map[string]interface{}{"num1": "498828660196", "num2": "840477629533"}, Expected: "419254329864656431168468", Note: "beyond 64-bit range", Hidden: true},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m + n)",
		StarterCode:     "def multiply(num1, num2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In long multiplication, digit i of num1 times digit j of num2 contributes to one column of the answer. Counting columns from the left of an answer with len(num1) + len(num2) digits, which column is it?"},
			{Level: 2, Type: "insight", Content: "Combine two Chapter 19 ideas: map the indices and add with carry. The product of num1[i] and num2[j] lands in answer positions i + j (its tens) and i + j + 1 (its ones), in an array of length m + n, which is the most digits a product can have. Go from the lowest digits, add each product to position i + j + 1, keep total % 10 there and push total // 10 into position i + j.", ChapterRef: 19},
			{Level: 3, Type: "plan", Content: "1. Handle zero: if either input is \"0\", return \"0\".\n2. Start at the lowest digit: loop i over num1 and j over num2, both from the right.\n3. Map the indices: the product goes to low = i + j + 1, its carry to high = i + j.\n4. Add with carry: total = product + result[low].\n5. Write the digit, keep the carry: result[low] = total % 10, result[high] += total // 10.\n6. Strip leading zeros and join the digits into a string."},
			{Level: 4, Type: "skeleton", Content: "def multiply(num1, num2):\n    if num1 == \"0\" or num2 == \"0\":\n        return \"0\"\n\n    m, n = len(num1), len(num2)\n    result = [0] * (m + n)\n    for i in range(m - 1, -1, -1):\n        for j in range(n - 1, -1, -1):\n            product = (ord(num1[i]) - ord(\"0\")) * (ord(num2[j]) - ord(\"0\"))\n            low, high = ____, ____\n            total = ____\n            result[low] = ____\n            result[high] += ____\n\n    start = 0\n    while start < len(result) - 1 and result[start] == 0:\n        start += 1\n    return \"\".join(str(d) for d in result[start:])"},
		},
		Solution: Solution{
			Code: `def multiply(num1, num2):
    if num1 == "0" or num2 == "0":
        return "0"

    m, n = len(num1), len(num2)
    result = [0] * (m + n)
    for i in range(m - 1, -1, -1):
        for j in range(n - 1, -1, -1):
            product = (ord(num1[i]) - ord("0")) * (ord(num2[j]) - ord("0"))
            low, high = i + j + 1, i + j
            total = product + result[low]
            result[low] = total % 10
            result[high] += total // 10

    start = 0
    while start < len(result) - 1 and result[start] == 0:
        start += 1
    return "".join(str(d) for d in result[start:])`,
			Explanation: "Builds on Plus One (add with carry, starting from the lowest digit) and on the index mapping of Rotate Image. Each pair of digits num1[i] * num2[j] has a fixed place in the answer, i + j + 1, with its carry going to i + j. Accumulating all m * n digit products into one array of m + n slots does long multiplication without ever building big numbers.",
			BruteForce:  "Paper-style: multiply num1 by each digit of num2 to get a partial product string, pad it with zeros, and add all partial products with string addition. It is also O(m * n), but it creates n intermediate strings of length up to m + n and adds each one, so it does about twice the work and is longer to write.",
			KeyInsight:  "The digit at index i (from the left) of an m-digit number has place value 10^(m - 1 - i). Multiplying place values 10^(m - 1 - i) and 10^(n - 1 - j) gives 10^(m + n - 2 - i - j), which is index i + j + 1 in an (m + n)-digit answer. So every digit product has a known destination and the carries can be settled as you go.",
			WhyItWorks:  "The product num1 * num2 equals the sum over all i, j of d1[i] * d2[j] * 10^(m + n - 2 - i - j), so adding each digit product into position i + j + 1 builds the correct total. The loops go from the right (low place values first), and each step adds the product to result[low], keeps total % 10 there and moves total // 10 into result[high], one place higher, so the value of the array (sum of result[k] * 10^(m + n - 1 - k)) always equals the sum of the products added so far. Every slot k >= 1 is the low position of some pair, and no carry is added to it after the last such pair has written total % 10, so it ends as a single digit; slot 0 only receives carries and is a single digit because the whole product is below 10^(m + n). The product of an m-digit and an n-digit number is below 10^(m + n), so m + n slots are always enough; at most one leading zero needs stripping.",
			CommonMistakes: []string{
				"Calling int(num1) * int(num2), which the problem forbids and which does not exist in languages without big integers. Fix: only convert single characters (ord(ch) - ord('0')).",
				"Writing the product to result[i + j] and the carry to result[i + j + 1], which reverses place values. Fix: the ones part goes to i + j + 1 (further right) and the carry to i + j.",
				"Returning the array with its leading zero, e.g. \"012345\" for 1 * 12345, or returning \"0000\" for a zero input. Fix: return \"0\" early for zero inputs and skip leading zeros before joining.",
			},
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m + n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Start at the lowest digit", Explanation: "Both loops run from the last character (the ones digit) towards the first, so carries always move towards positions not yet finished.", CodeSnippet: "for i in range(m - 1, -1, -1):\n    for j in range(n - 1, -1, -1):", LineStart: 7, LineEnd: 9},
				{Title: "Map the indices", Explanation: "num1[i] * num2[j] has place value at index i + j + 1 of the answer; its carry belongs one place to the left at i + j.", CodeSnippet: "low, high = i + j + 1, i + j", LineStart: 10, LineEnd: 10},
				{Title: "Add with carry", Explanation: "The slot may already hold a digit from an earlier pair, so add the product to it.", CodeSnippet: "total = product + result[low]", LineStart: 11, LineEnd: 11},
				{Title: "Write the digit, keep the carry", Explanation: "Keep the ones digit here and push the rest one place left. result[high] can exceed 9 temporarily; the next pair that writes there settles it.", CodeSnippet: "result[low] = total % 10\nresult[high] += total // 10", LineStart: 12, LineEnd: 13},
				{Title: "Strip leading zeros", Explanation: "The answer has m + n or m + n - 1 digits, so at most a leading 0 is skipped before joining.", CodeSnippet: "while start < len(result) - 1 and result[start] == 0:", LineStart: 15, LineEnd: 18},
			},
			Alternatives: []AltSolution{
				{
					Name: "Partial products plus string addition",
					Code: `def multiply(num1, num2):
    def add(a, b):
        out, carry = [], 0
        i, j = len(a) - 1, len(b) - 1
        while i >= 0 or j >= 0 or carry:
            total = carry
            if i >= 0:
                total += ord(a[i]) - 48
                i -= 1
            if j >= 0:
                total += ord(b[j]) - 48
                j -= 1
            out.append(chr(48 + total % 10))
            carry = total // 10
        return "".join(reversed(out))

    def times_digit(a, d):
        out, carry = [], 0
        for ch in reversed(a):
            total = (ord(ch) - 48) * d + carry
            out.append(chr(48 + total % 10))
            carry = total // 10
        if carry:
            out.append(chr(48 + carry))
        return "".join(reversed(out))

    if num1 == "0" or num2 == "0":
        return "0"
    total = "0"
    for shift, ch in enumerate(reversed(num2)):
        partial = times_digit(num1, ord(ch) - 48) + "0" * shift
        total = add(total, partial)
    return total.lstrip("0") or "0"`,
					TimeComplexity:  "O(n * (m + n))",
					SpaceComplexity: "O(m + n)",
					Explanation:     "Exactly the method from school: multiply num1 by one digit of num2 at a time, shift by appending zeros, and add the partial products using the add-with-carry loop. It is easy to trust but builds a new string for every digit of num2.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is an array of length m + n always big enough for the product?", Answer: "num1 < 10^m and num2 < 10^n, so num1 * num2 < 10^(m + n), which has at most m + n digits. It has at least m + n - 1 digits (10^(m-1) * 10^(n-1) = 10^(m+n-2)), so at most one leading zero is stripped."},
			{Prompt: "result[high] += total // 10 can leave a value of 10 or more in result[high]. Why is the final array still all single digits?", Answer: "Position high = i + j is the low position of a pair processed later: (i, j - 1) when j > 0, or (i - 1, 0) when j = 0. That pair adds result[high] into its total and writes back only total % 10, and no carry reaches the slot after it. Index 0 only ever receives carries, and the total is below 10^(m + n), so it ends as a single digit."},
		},
	},
	{
		ID:              "detect-squares",
		Number:          2013,
		Order:           8,
		Title:           "Detect Squares",
		Slug:            "detect-squares",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Hash Table", "Design", "Counting"},
		RelatedChapters: []int{19, 2},
		Description: `Design a class DetectSquares that receives points on a grid and answers square-counting queries.

- DetectSquares() creates an empty collection.
- add(point) adds point = [x, y] to the collection. The same point can be added more than once, and each copy counts as a separate point.
- count(point) is given a query point [x, y]. Return the number of ways to choose three points from the collection that, together with the query point, are the four corners of an axis-aligned square with positive area. Axis-aligned means the sides are horizontal and vertical. Choosing a different copy of a duplicated point counts as a different way.

The query point itself is not added to the collection.`,
		Constraints: []string{
			"point.length == 2",
			"0 <= x, y <= 1000",
			"At most 3000 calls in total are made to add and count",
		},
		Examples: []Example{
			{ID: "1", Input: `["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
[[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]`, Output: "[null, null, null, null, 1, 0, null, 2]", Explanation: "With (3,10), (11,2) and (3,2) stored, query (11,10) completes the square with side 8: one way. (14,8) completes nothing. After a second copy of (11,2) is added, either copy can be used, so the answer becomes 2."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"DetectSquares", "add", "add", "add", "count", "count", "add", "count"}, "values": []interface{}{[]interface{}{}, [][]int{{3, 10}}, [][]int{{11, 2}}, [][]int{{3, 2}}, [][]int{{11, 10}}, [][]int{{14, 8}}, [][]int{{11, 2}}, [][]int{{11, 10}}}}, Expected: []interface{}{nil, nil, nil, nil, 1, 0, nil, 2}, Note: "the example: duplicates multiply the count"},
			{Input: map[string]interface{}{"operations": []string{"DetectSquares", "add", "add", "add", "add", "count"}, "values": []interface{}{[]interface{}{}, [][]int{{3, 10}}, [][]int{{11, 2}}, [][]int{{3, 2}}, [][]int{{11, 2}}, [][]int{{11, 10}}}}, Expected: []interface{}{nil, nil, nil, nil, nil, 2}, Note: "duplicate on the query's column: looping over a list with repeats counts 4 instead of 2"},
			{Input: map[string]interface{}{"operations": []string{"DetectSquares", "count", "add", "count"}, "values": []interface{}{[]interface{}{}, [][]int{{5, 5}}, [][]int{{5, 5}}, [][]int{{5, 5}}}}, Expected: []interface{}{nil, 0, nil, 0}, Note: "empty collection, then a point equal to the query", Hidden: true},
			{Input: map[string]interface{}{"operations": []string{"DetectSquares", "add", "add", "add", "count", "add", "add", "count"}, "values": []interface{}{[]interface{}{}, [][]int{{1, 1}}, [][]int{{1, 1}}, [][]int{{1, 1}}, [][]int{{1, 1}}, [][]int{{1, 2}}, [][]int{{2, 1}}, [][]int{{1, 1}}}}, Expected: []interface{}{nil, nil, nil, nil, 0, nil, nil, 0}, Note: "zero-area squares and a missing fourth corner do not count", Hidden: true},
			{Input: map[string]interface{}{"operations": []string{"DetectSquares", "add", "add", "add", "add", "add", "add", "add", "count", "count", "count"}, "values": []interface{}{[]interface{}{}, [][]int{{1, 0}}, [][]int{{0, 1}}, [][]int{{1, 1}}, [][]int{{2, 1}}, [][]int{{1, 2}}, [][]int{{2, 2}}, [][]int{{0, 2}}, [][]int{{0, 0}}, [][]int{{2, 0}}, [][]int{{0, 1}}}}, Expected: []interface{}{nil, nil, nil, nil, nil, nil, nil, nil, 1, 1, 1}, Note: "squares to the upper-left, upper-right and right of the query", Hidden: true},
			{Input: map[string]interface{}{"operations": []string{"DetectSquares", "add", "add", "add", "add", "add", "add", "count", "count"}, "values": []interface{}{[]interface{}{}, [][]int{{5, 0}}, [][]int{{0, 5}}, [][]int{{5, 5}}, [][]int{{10, 0}}, [][]int{{10, 5}}, [][]int{{0, 0}}, [][]int{{0, 0}}, [][]int{{5, 0}}}}, Expected: []interface{}{nil, nil, nil, nil, nil, nil, nil, 1, 2}, Note: "query between two squares: both sides count", Hidden: true},
			{Input: map[string]interface{}{"operations": []string{"DetectSquares", "add", "add", "add", "add", "add", "count", "count", "add", "count"}, "values": []interface{}{[]interface{}{}, [][]int{{2, 2}}, [][]int{{2, 4}}, [][]int{{4, 2}}, [][]int{{2, 3}}, [][]int{{3, 2}}, [][]int{{4, 4}}, [][]int{{3, 3}}, [][]int{{4, 4}}, [][]int{{3, 3}}}}, Expected: []interface{}{nil, nil, nil, nil, nil, nil, 1, 1, nil, 1}, Note: "squares of different sizes sharing a corner", Hidden: true},
			{Input: map[string]interface{}{"operations": []string{"DetectSquares", "add", "add", "add", "add", "add", "count", "count"}, "values": []interface{}{[]interface{}{}, [][]int{{0, 0}}, [][]int{{0, 0}}, [][]int{{0, 1}}, [][]int{{0, 1}}, [][]int{{1, 0}}, [][]int{{1, 1}}, [][]int{{1, 0}}}}, Expected: []interface{}{nil, nil, nil, nil, nil, nil, 4, 0}, Note: "duplicates at two corners multiply: 2 * 2 * 1 = 4", Hidden: true},
		},
		TimeComplexity:  "O(1) per add, O(p) per count, where p is the number of distinct points",
		SpaceComplexity: "O(p)",
		StarterCode:     "class DetectSquares:\n    def __init__(self):\n        pass\n\n    def add(self, point):\n        pass\n\n    def count(self, point):\n        pass",
		Harness:         Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The query point is one corner. If you also knew the corner diagonally opposite it, would the other two corners be fixed? How many copies of each would you need to know about?"},
			{Level: 2, Type: "insight", Content: "Pick the diagonal corner. A stored point (px, py) is the opposite corner of an axis-aligned square exactly when |px - qx| == |py - qy| and px != qx. The other two corners are then forced: (px, qy) and (qx, py). Keep a hash map from point to how many copies were added (Chapter 2); the number of squares using that diagonal is the product of the three corner counts. Loop over distinct points, so each copy is counted through the multiplication and never through the loop.", ChapterRef: 19},
			{Level: 3, Type: "plan", Content: "1. Count each point: add() increments a counter keyed by (x, y).\n2. Pick the diagonal corner: in count(), loop over each distinct stored point (px, py) with its count; skip it unless |px - qx| == |py - qy| and px != qx.\n3. Multiply the corner counts: add count(px, py) * count(px, qy) * count(qx, py) to the answer.\n4. Return the total."},
			{Level: 4, Type: "skeleton", Content: "from collections import Counter\n\nclass DetectSquares:\n    def __init__(self):\n        self.count_at = Counter()\n\n    def add(self, point):\n        ____\n\n    def count(self, point):\n        qx, qy = point\n        ways = 0\n        for (px, py), diagonal_count in self.count_at.items():\n            if ____:\n                continue\n            ways += ____\n        return ways"},
		},
		Solution: Solution{
			Code: `from collections import Counter

class DetectSquares:
    def __init__(self):
        self.count_at = Counter()

    def add(self, point):
        self.count_at[(point[0], point[1])] += 1

    def count(self, point):
        qx, qy = point
        ways = 0
        for (px, py), diagonal_count in self.count_at.items():
            if abs(px - qx) != abs(py - qy) or px == qx:
                continue
            ways += diagonal_count * self.count_at[(px, qy)] * self.count_at[(qx, py)]
        return ways`,
			Explanation: "The last rung combines coordinate reasoning (the index work of Rotate Image and Spiral Matrix, now on x/y points) with hash-map counting from Chapter 2. Fixing the diagonal corner determines the whole square, and a Counter of points turns 'how many ways' into a product of three counts.",
			BruteForce:  "Store every added point in a list and, for each query, try every triple of stored points and check whether the four points form a square: O(k^3) per query for k stored points, which is about 2.7 * 10^10 checks when k = 3000. Even fixing the diagonal and scanning the list for the other two corners is O(k^2) per query.",
			KeyInsight:  "Two opposite corners of an axis-aligned square determine the other two. So the only free choice is the diagonal corner, and the number of ways for that choice is just the product of how many copies sit at each of the three required positions.",
			WhyItWorks:  "Every valid square with the query at one corner has exactly one corner diagonally opposite the query, at (qx + s, qy + t) with |s| == |t| > 0, and then its other corners must be (qx + s, qy) and (qx, qy + t). So squares are in one-to-one correspondence with choices of a diagonal position plus one copy at each of the three corner positions, and for a fixed diagonal position the number of copy choices is the product of the three counts (zero if any corner is missing). Looping over distinct positions visits each diagonal position once, so nothing is double counted; the px != qx test rules out zero-area 'squares'. This fixes the bug of looping over a list that contains duplicates and also multiplying by the duplicate's count, which counts each square count^2 times instead of count times.",
			CommonMistakes: []string{
				"Looping over a list of stored points that contains duplicates and also multiplying by that point's count, so each duplicate is counted twice: with (3,10), (11,2), (3,2), (11,2) stored, count([11,10]) returns 4 instead of 2. Fix: loop over distinct points (the Counter's keys) and let the multiplication account for copies.",
				"Forgetting the positive-area rule, so a stored copy of the query point (side 0) contributes squares. Fix: skip candidates with px == qx (which, with equal side lengths, also means py == qy).",
				"Only checking squares on one side of the query (for example only up and to the right). Fix: the diagonal can be in any of the four directions, which the |px - qx| == |py - qy| test covers.",
				"Indexing a plain dict with missing corners, which raises KeyError, or using a defaultdict and inserting keys while iterating over it. Fix: use a Counter, whose lookups of missing keys return 0 without inserting.",
			},
			TimeComplexity:  "O(1) per add, O(p) per count, where p is the number of distinct points",
			SpaceComplexity: "O(p)",
			Walkthrough: []WalkthroughStep{
				{Title: "Count each point", Explanation: "Duplicates are allowed and each copy is a separate choice, so store a count per position rather than a set.", CodeSnippet: "self.count_at[(point[0], point[1])] += 1", LineStart: 7, LineEnd: 8},
				{Title: "Pick the diagonal corner", Explanation: "Each distinct stored position is tried once as the corner opposite the query. Equal horizontal and vertical distance makes it a square; px != qx makes the area positive.", CodeSnippet: "for (px, py), diagonal_count in self.count_at.items():\n    if abs(px - qx) != abs(py - qy) or px == qx:\n        continue", LineStart: 13, LineEnd: 15},
				{Title: "Multiply the corner counts", Explanation: "With the diagonal fixed, the other corners are (px, qy) and (qx, py). Any copy at each of the three positions can be chosen, so the counts multiply. A missing corner has count 0 and contributes nothing.", CodeSnippet: "ways += diagonal_count * self.count_at[(px, qy)] * self.count_at[(qx, py)]", LineStart: 16, LineEnd: 16},
			},
			Alternatives: []AltSolution{
				{
					Name: "Scan the query's column (distinct y values)",
					Code: `from collections import Counter, defaultdict

class DetectSquares:
    def __init__(self):
        self.count_at = Counter()
        self.ys_in_column = defaultdict(Counter)

    def add(self, point):
        x, y = point
        self.count_at[(x, y)] += 1
        self.ys_in_column[x][y] += 1

    def count(self, point):
        qx, qy = point
        ways = 0
        for py, same_column_count in self.ys_in_column[qx].items():
            side = py - qy
            if side == 0:
                continue
            for px in (qx + side, qx - side):
                ways += same_column_count * self.count_at[(px, qy)] * self.count_at[(px, py)]
        return ways`,
					TimeComplexity:  "O(1) per add, O(c) per count, where c is the number of distinct y values in the query's column",
					SpaceComplexity: "O(p)",
					Explanation:     "Fix the corner directly above or below the query instead of the diagonal: it sets the side length, and the square can extend left or right. Iterating over the distinct y values in that column (a Counter, not a list with repeats) is what avoids double counting duplicates. It is faster when columns are sparse.",
				},
				{
					Name: "Brute force over a list of points",
					Code: `class DetectSquares:
    def __init__(self):
        self.points = []

    def add(self, point):
        self.points.append((point[0], point[1]))

    def count(self, point):
        qx, qy = point
        ways = 0
        for dx, dy in self.points:
            if dx == qx or abs(dx - qx) != abs(dy - qy):
                continue
            ways += self.points.count((qx, dy)) * self.points.count((dx, qy))
        return ways`,
					TimeComplexity:  "O(k^2) per count, where k is the number of added points",
					SpaceComplexity: "O(k)",
					Explanation:     "Keeps every copy in a list and tries each copy as the diagonal corner, counting the other two corners by scanning the list. Each copy is visited once by the loop and is not multiplied again, so it is correct, just slow.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "With (3,10), (11,2), (3,2), (11,2) stored, why is count([11,10]) equal to 2, and how does a loop that over-counts reach 4?", Answer: "The only square uses diagonal corner (3,2) and corners (3,10) and (11,2). There is 1 copy of (3,2), 1 of (3,10) and 2 of (11,2), so 1 * 1 * 2 = 2 ways. A loop over the list of y values in column 11 sees y = 2 twice and multiplies by count(11,2) = 2 each time, giving 2 + 2 = 4: the duplicates are counted once by the loop and again by the multiplication."},
			{Prompt: "What would change if the squares did not have to be axis-aligned (they could be tilted)?", Answer: "The diagonal no longer fixes the square through equal x and y distances. Instead, for each stored point taken as an adjacent corner, rotate the side vector (dx, dy) by 90 degrees to (-dy, dx) to get the other two corners, check both rotation directions, and multiply the counts. Each square would then be found from both of its neighbours of the query, so the total has to be divided by 2."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, MathProblems...)
}
