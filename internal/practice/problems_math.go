package practice

// MathProblems contains all math-geometry category problems
var MathProblems = []*Problem{
	{
		ID:              "rotate-image",
		Number:          92,
		Title:           "Rotate Image",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Math", "Matrix"},
		RelatedChapters: []int{1},
		Description: `You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.`,
		Constraints: []string{
			"n == matrix.length == matrix[i].length",
			"1 <= n <= 20",
			"-1000 <= matrix[i][j] <= 1000",
		},
		Examples: []Example{
			{Input: "matrix = [[1,2,3],[4,5,6],[7,8,9]]", Output: "[[7,4,1],[8,5,2],[9,6,3]]"},
			{Input: "matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]", Output: "[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, Expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def rotate(matrix):\n    # Write your solution here (modify in place)\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "90-degree clockwise rotation = transpose + reverse each row. Or reverse columns + transpose."},
			{Level: 2, Type: "algorithm", Content: "Transpose: swap matrix[i][j] with matrix[j][i]. Then reverse each row."},
			{Level: 3, Type: "code", Content: "for i in range(n): for j in range(i+1, n): swap [i][j], [j][i]. Then for each row: reverse."},
		},
		Solution: Solution{
			Code: `def rotate(matrix):
    n = len(matrix)

    # Transpose
    for i in range(n):
        for j in range(i + 1, n):
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

    # Reverse each row
    for row in matrix:
        row.reverse()`,
			Explanation:     "90-degree clockwise = transpose + reverse rows. Both operations in-place.",
			TimeComplexity:  "O(n^2)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Transpose", Explanation: "Swap across diagonal", CodeSnippet: "matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]", LineStart: 7, LineEnd: 7},
				{Title: "Reverse rows", Explanation: "Complete the rotation", CodeSnippet: "row.reverse()", LineStart: 11, LineEnd: 11},
			},
		},
	},
	{
		ID:              "spiral-matrix",
		Number:          93,
		Title:           "Spiral Matrix",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Matrix", "Simulation"},
		RelatedChapters: []int{1},
		Description: `Given an m x n matrix, return all elements of the matrix in spiral order.`,
		Constraints: []string{
			"m == matrix.length",
			"n == matrix[i].length",
			"1 <= m, n <= 10",
			"-100 <= matrix[i][j] <= 100",
		},
		Examples: []Example{
			{Input: "matrix = [[1,2,3],[4,5,6],[7,8,9]]", Output: "[1,2,3,6,9,8,7,4,5]"},
			{Input: "matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]", Output: "[1,2,3,4,8,12,11,10,9,5,6,7]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, Expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def spiralOrder(matrix):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Track four boundaries: top, bottom, left, right. Traverse and shrink boundaries after each direction."},
			{Level: 2, Type: "algorithm", Content: "Go right (top row), down (right col), left (bottom row), up (left col). Update boundaries after each."},
			{Level: 3, Type: "code", Content: "while top <= bottom and left <= right: traverse right, top++; traverse down, right--; check, traverse left, bottom--; check, traverse up, left++."},
		},
		Solution: Solution{
			Code: `def spiralOrder(matrix):
    result = []
    top, bottom = 0, len(matrix) - 1
    left, right = 0, len(matrix[0]) - 1

    while top <= bottom and left <= right:
        # Right
        for c in range(left, right + 1):
            result.append(matrix[top][c])
        top += 1

        # Down
        for r in range(top, bottom + 1):
            result.append(matrix[r][right])
        right -= 1

        if top <= bottom:
            # Left
            for c in range(right, left - 1, -1):
                result.append(matrix[bottom][c])
            bottom -= 1

        if left <= right:
            # Up
            for r in range(bottom, top - 1, -1):
                result.append(matrix[r][left])
            left += 1

    return result`,
			Explanation:     "Four boundaries shrink as we traverse. Go right, down, left, up. Check boundaries before left/up.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(1) extra space",
			Walkthrough: []WalkthroughStep{
				{Title: "Four boundaries", Explanation: "Track traversal limits", CodeSnippet: "top, bottom = 0, len(matrix) - 1\nleft, right = 0, len(matrix[0]) - 1", LineStart: 3, LineEnd: 4},
				{Title: "Traverse and shrink", Explanation: "Go right then increment top", CodeSnippet: "for c in range(left, right + 1):\n    result.append(matrix[top][c])\ntop += 1", LineStart: 8, LineEnd: 10},
				{Title: "Check before left/up", Explanation: "Boundaries may have crossed", CodeSnippet: "if top <= bottom:", LineStart: 17, LineEnd: 17},
			},
		},
	},
	{
		ID:              "set-matrix-zeroes",
		Number:          101,
		Title:           "Set Matrix Zeroes",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Hash Table", "Matrix"},
		RelatedChapters: []int{1, 5},
		Description: `Given an m x n integer matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.`,
		Constraints: []string{
			"m == matrix.length",
			"n == matrix[0].length",
			"1 <= m, n <= 200",
			"-2^31 <= matrix[i][j] <= 2^31 - 1",
		},
		Examples: []Example{
			{Input: "matrix = [[1,1,1],[1,0,1],[1,1,1]]", Output: "[[1,0,1],[0,0,0],[1,0,1]]"},
			{Input: "matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]", Output: "[[0,0,0,0],[0,4,5,0],[0,3,1,0]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}, Expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def setZeroes(matrix):\n    # Write your solution here (modify in place)\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use first row and column as markers. Track if first row/col themselves need zeroing."},
			{Level: 2, Type: "algorithm", Content: "Mark zeros in first row/col. Then iterate and zero based on markers. Finally handle first row/col."},
			{Level: 3, Type: "code", Content: "first_row_zero, first_col_zero = check first row/col. Mark in first row/col. Zero inner. Zero first row/col if needed."},
		},
		Solution: Solution{
			Code: `def setZeroes(matrix):
    m, n = len(matrix), len(matrix[0])
    first_row_zero = any(matrix[0][j] == 0 for j in range(n))
    first_col_zero = any(matrix[i][0] == 0 for i in range(m))

    # Mark zeros in first row/column
    for i in range(1, m):
        for j in range(1, n):
            if matrix[i][j] == 0:
                matrix[i][0] = 0
                matrix[0][j] = 0

    # Zero cells based on markers
    for i in range(1, m):
        for j in range(1, n):
            if matrix[i][0] == 0 or matrix[0][j] == 0:
                matrix[i][j] = 0

    # Zero first row if needed
    if first_row_zero:
        for j in range(n):
            matrix[0][j] = 0

    # Zero first column if needed
    if first_col_zero:
        for i in range(m):
            matrix[i][0] = 0`,
			Explanation:     "Use first row/column as markers. Track if they need zeroing separately. Zero inner cells, then first row/col.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Check first row/col", Explanation: "Track if they have zeros before using as markers", CodeSnippet: "first_row_zero = any(matrix[0][j] == 0 for j in range(n))", LineStart: 3, LineEnd: 4},
				{Title: "Mark zeros", Explanation: "Use first row/col as markers", CodeSnippet: "matrix[i][0] = 0\nmatrix[0][j] = 0", LineStart: 10, LineEnd: 11},
				{Title: "Zero based on markers", Explanation: "Check row and column markers", CodeSnippet: "if matrix[i][0] == 0 or matrix[0][j] == 0:", LineStart: 15, LineEnd: 16},
			},
		},
	},
	{
		ID:              "happy-number",
		Number:          116,
		Title:           "Happy Number",
		Difficulty:      "Easy",
		Category:        "math-geometry",
		Tags:            []string{"Hash Table", "Math", "Two Pointers"},
		RelatedChapters: []int{1, 5},
		Description: `Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

- Starting with any positive integer, replace the number by the sum of the squares of its digits.
- Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
- Those numbers for which this process ends in 1 are happy.

Return true if n is a happy number, and false if not.`,
		Constraints: []string{
			"1 <= n <= 2^31 - 1",
		},
		Examples: []Example{
			{Input: "n = 19", Output: "true", Explanation: "1^2 + 9^2 = 82, 8^2 + 2^2 = 68, 6^2 + 8^2 = 100, 1^2 + 0^2 + 0^2 = 1"},
			{Input: "n = 2", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 19}, Expected: true},
			{Input: map[string]interface{}{"n": 2}, Expected: false},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def isHappy(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Detect cycle using Floyd's tortoise and hare, or use a hash set to track seen numbers."},
			{Level: 2, Type: "algorithm", Content: "Slow and fast pointers: slow = next(n), fast = next(next(n)). If they meet and != 1, cycle without 1."},
			{Level: 3, Type: "code", Content: "def get_next(n): sum of digit squares. slow, fast = n, get_next(n). while fast != 1 and slow != fast: advance. return fast == 1."},
		},
		Solution: Solution{
			Code: `def isHappy(n):
    def get_next(num):
        total = 0
        while num > 0:
            digit = num % 10
            total += digit * digit
            num //= 10
        return total

    slow = n
    fast = get_next(n)

    while fast != 1 and slow != fast:
        slow = get_next(slow)
        fast = get_next(get_next(fast))

    return fast == 1`,
			Explanation:     "Floyd's cycle detection. If we reach 1, happy. If cycle detected (slow == fast), not happy.",
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Digit sum of squares", Explanation: "Calculate next number in sequence", CodeSnippet: "total += digit * digit", LineStart: 6, LineEnd: 6},
				{Title: "Floyd's algorithm", Explanation: "Slow moves 1 step, fast moves 2", CodeSnippet: "slow = get_next(slow)\nfast = get_next(get_next(fast))", LineStart: 14, LineEnd: 15},
			},
		},
	},
	{
		ID:              "plus-one",
		Number:          117,
		Title:           "Plus One",
		Difficulty:      "Easy",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Math"},
		RelatedChapters: []int{1},
		Description: `You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.`,
		Constraints: []string{
			"1 <= digits.length <= 100",
			"0 <= digits[i] <= 9",
			"digits does not contain any leading 0's",
		},
		Examples: []Example{
			{Input: "digits = [1,2,3]", Output: "[1,2,4]"},
			{Input: "digits = [4,3,2,1]", Output: "[4,3,2,2]"},
			{Input: "digits = [9]", Output: "[1,0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"digits": []int{1, 2, 3}}, Expected: []int{1, 2, 4}},
			{Input: map[string]interface{}{"digits": []int{9, 9, 9}}, Expected: []int{1, 0, 0, 0}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def plusOne(digits):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Add 1 to last digit. If 9, set to 0 and carry. If all 9s, prepend 1."},
			{Level: 2, Type: "algorithm", Content: "Iterate from right. If digit < 9, increment and return. If 9, set to 0. If we exit loop, prepend 1."},
			{Level: 3, Type: "code", Content: "for i from n-1 to 0: if digits[i] < 9: digits[i] += 1, return. else: digits[i] = 0. return [1] + digits."},
		},
		Solution: Solution{
			Code: `def plusOne(digits):
    for i in range(len(digits) - 1, -1, -1):
        if digits[i] < 9:
            digits[i] += 1
            return digits
        digits[i] = 0

    return [1] + digits`,
			Explanation:     "Add 1 from right. If digit < 9, done. If 9, set to 0 and continue. If all 9s, prepend 1.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1) or O(n) if all 9s",
			Walkthrough: []WalkthroughStep{
				{Title: "No carry needed", Explanation: "Digit < 9, just increment", CodeSnippet: "if digits[i] < 9:\n    digits[i] += 1\n    return digits", LineStart: 3, LineEnd: 5},
				{Title: "Carry needed", Explanation: "Digit is 9, set to 0 and continue", CodeSnippet: "digits[i] = 0", LineStart: 6, LineEnd: 6},
				{Title: "All 9s case", Explanation: "Need new leading 1", CodeSnippet: "return [1] + digits", LineStart: 8, LineEnd: 8},
			},
		},
	},
	{
		ID:              "powx-n",
		Number:          118,
		Title:           "Pow(x, n)",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Math", "Recursion"},
		RelatedChapters: []int{1, 4},
		Description: `Implement pow(x, n), which calculates x raised to the power n (i.e., x^n).`,
		Constraints: []string{
			"-100.0 < x < 100.0",
			"-2^31 <= n <= 2^31 - 1",
			"n is an integer",
			"Either x is not zero or n > 0",
			"-10^4 <= x^n <= 10^4",
		},
		Examples: []Example{
			{Input: "x = 2.00000, n = 10", Output: "1024.00000"},
			{Input: "x = 2.10000, n = 3", Output: "9.26100"},
			{Input: "x = 2.00000, n = -2", Output: "0.25000"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"x": 2.0, "n": 10}, Expected: 1024.0},
			{Input: map[string]interface{}{"x": 2.0, "n": -2}, Expected: 0.25},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def myPow(x, n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Binary exponentiation: x^n = (x^2)^(n/2). Handle negative exponent by using 1/x.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "If n is even: x^n = (x*x)^(n/2). If n is odd: x^n = x * x^(n-1). Handle n < 0."},
			{Level: 3, Type: "code", Content: "if n < 0: x = 1/x, n = -n. result = 1. while n: if n odd: result *= x. x *= x, n //= 2. return result."},
		},
		Solution: Solution{
			Code: `def myPow(x, n):
    if n < 0:
        x = 1 / x
        n = -n

    result = 1
    while n:
        if n % 2 == 1:
            result *= x
        x *= x
        n //= 2

    return result`,
			Explanation:     "Binary exponentiation. Square x each iteration. If odd power, multiply into result. O(log n) iterations.",
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle negative", Explanation: "x^(-n) = (1/x)^n", CodeSnippet: "if n < 0:\n    x = 1 / x\n    n = -n", LineStart: 2, LineEnd: 4},
				{Title: "Odd power", Explanation: "Multiply current x into result", CodeSnippet: "if n % 2 == 1:\n    result *= x", LineStart: 8, LineEnd: 9},
				{Title: "Square and halve", Explanation: "x^n = (x^2)^(n/2)", CodeSnippet: "x *= x\nn //= 2", LineStart: 10, LineEnd: 11},
			},
		},
	},
	{
		ID:              "multiply-strings",
		Number:          140,
		Title:           "Multiply Strings",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Math", "String", "Simulation"},
		RelatedChapters: []int{1, 12},
		Description: `Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.`,
		Constraints: []string{
			"1 <= num1.length, num2.length <= 200",
			"num1 and num2 consist of digits only",
			"Both num1 and num2 do not contain any leading zero, except the number 0 itself",
		},
		Examples: []Example{
			{Input: `num1 = "2", num2 = "3"`, Output: `"6"`},
			{Input: `num1 = "123", num2 = "456"`, Output: `"56088"`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"num1": "2", "num2": "3"}, Expected: "6"},
			{Input: map[string]interface{}{"num1": "123", "num2": "456"}, Expected: "56088"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m + n)",
		StarterCode:     "def multiply(num1, num2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Simulate grade-school multiplication. Position i*j contributes to position i+j and i+j+1.", ChapterRef: 1},
			{Level: 2, Type: "algorithm", Content: "Create result array of length m+n. Multiply each pair of digits, accumulate at correct positions."},
			{Level: 3, Type: "code", Content: "result = [0]*(m+n). for i,j: mul = d1*d2. result[i+j+1] += mul. Handle carries. Convert to string."},
		},
		Solution: Solution{
			Code: `def multiply(num1, num2):
    if num1 == "0" or num2 == "0":
        return "0"

    m, n = len(num1), len(num2)
    result = [0] * (m + n)

    for i in range(m - 1, -1, -1):
        for j in range(n - 1, -1, -1):
            mul = int(num1[i]) * int(num2[j])
            p1, p2 = i + j, i + j + 1
            total = mul + result[p2]

            result[p2] = total % 10
            result[p1] += total // 10

    # Remove leading zeros
    result_str = ''.join(map(str, result))
    return result_str.lstrip('0') or '0'`,
			Explanation:     "Simulate multiplication. digit[i] * digit[j] contributes to positions i+j and i+j+1.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m + n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Position mapping", Explanation: "i*j affects positions i+j and i+j+1", CodeSnippet: "p1, p2 = i + j, i + j + 1", LineStart: 11, LineEnd: 11},
				{Title: "Handle carry", Explanation: "Digit goes to p2, carry to p1", CodeSnippet: "result[p2] = total % 10\nresult[p1] += total // 10", LineStart: 14, LineEnd: 15},
			},
		},
	},
	{
		ID:              "detect-squares",
		Number:          141,
		Title:           "Detect Squares",
		Difficulty:      "Medium",
		Category:        "math-geometry",
		Tags:            []string{"Array", "Hash Table", "Design", "Counting"},
		RelatedChapters: []int{5, 12},
		Description: `You are given a stream of points on the X-Y plane. Design an algorithm that:

- Adds new points from the stream into a data structure. Duplicate points are allowed.
- Given a query point, counts the number of ways to choose three points from the data structure such that the three points and the query point form an axis-aligned square with positive area.

An axis-aligned square is a square whose edges are all the same length and are either parallel or perpendicular to the x-axis and y-axis.`,
		Constraints: []string{
			"point.length == 2",
			"0 <= x, y <= 1000",
			"At most 3000 calls in total will be made to add and count",
		},
		Examples: []Example{
			{Input: `["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
[[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]`, Output: "[null, null, null, null, 1, 0, null, 2]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"DetectSquares", "add", "add", "add", "count"}, "args": [][]int{{}, {3, 10}, {11, 2}, {3, 2}, {11, 10}}}, Expected: []interface{}{nil, nil, nil, nil, 1}},
		},
		TimeComplexity:  "O(n) for count",
		SpaceComplexity: "O(n)",
		StarterCode:     "class DetectSquares:\n    def __init__(self):\n        pass\n\n    def add(self, point):\n        pass\n\n    def count(self, point):\n        pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Store point counts. For count query, iterate diagonal points.", ChapterRef: 5},
			{Level: 2, Type: "algorithm", Content: "For query (x,y), find all points (px,py) where |px-x|==|py-y|!=0. Check if other two corners exist."},
			{Level: 3, Type: "code", Content: "cnt = Counter of points. For diagonal (px,py): count += cnt[(px,py)] * cnt[(x,py)] * cnt[(px,y)]."},
		},
		Solution: Solution{
			Code: `class DetectSquares:
    def __init__(self):
        from collections import Counter, defaultdict
        self.points = Counter()
        self.x_coords = defaultdict(list)  # x -> list of y values

    def add(self, point):
        x, y = point
        self.points[(x, y)] += 1
        self.x_coords[x].append(y)

    def count(self, point):
        x, y = point
        result = 0

        for py in self.x_coords[x]:
            side = py - y
            if side == 0:
                continue

            # Check two possible squares (left and right)
            for px in [x + side, x - side]:
                result += self.points[(x, py)] * self.points[(px, y)] * self.points[(px, py)]

        return result`,
			Explanation:     "For query point, find diagonal candidates sharing same x. Check if square corners exist.",
			TimeComplexity:  "O(n) for count",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track by x", Explanation: "Store y values for each x coordinate", CodeSnippet: "self.x_coords[x].append(y)", LineStart: 10, LineEnd: 10},
				{Title: "Find diagonals", Explanation: "Same x, different y gives side length", CodeSnippet: "side = py - y", LineStart: 17, LineEnd: 17},
				{Title: "Multiply counts", Explanation: "Product of point counts at three corners", CodeSnippet: "self.points[(x, py)] * self.points[(px, y)] * self.points[(px, py)]", LineStart: 23, LineEnd: 23},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, MathProblems...)
}
