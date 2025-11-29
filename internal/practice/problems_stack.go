package practice

// StackProblems contains all stack category problems
var StackProblems = []*Problem{
	{
		ID:         "valid-parentheses",
		Number:     19,
		Title:      "Valid Parentheses",
		Difficulty: "Easy",
		Category:   "stack",
		Tags:       []string{"String", "Stack"},
		RelatedChapters: []int{8},
		Description: `Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.`,
		Constraints: []string{
			"1 <= s.length <= 10^4",
			"s consists of parentheses only '()[]{}'",
		},
		Examples: []Example{
			{Input: `s = "()"`, Output: "true"},
			{Input: `s = "()[]{}"`, Output: "true"},
			{Input: `s = "(]"`, Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "()"}, Expected: true},
			{Input: map[string]interface{}{"s": "()[]{}"}, Expected: true},
			{Input: map[string]interface{}{"s": "(]"}, Expected: false},
			{Input: map[string]interface{}{"s": "([)]"}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def isValid(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a stack to track open brackets."},
			{Level: 2, Type: "algorithm", Content: "Push opens, pop and match for closes."},
		},
		Solution: Solution{
			Code: `def isValid(s):
    stack = []
    pairs = {')': '(', '}': '{', ']': '['}
    for char in s:
        if char in pairs:
            if not stack or stack[-1] != pairs[char]:
                return False
            stack.pop()
        else:
            stack.append(char)
    return len(stack) == 0`,
			Explanation: "Stack for opens. For closes, check if matches top of stack.",
		},
	},
	{
		ID:              "min-stack",
		Number:          40,
		Title:           "Min Stack",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Stack", "Design"},
		RelatedChapters: []int{3, 8, 12},
		Description: `Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:
- MinStack() initializes the stack object.
- void push(int val) pushes the element val onto the stack.
- void pop() removes the element on the top of the stack.
- int top() gets the top element of the stack.
- int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function.`,
		Constraints: []string{
			"-2^31 <= val <= 2^31 - 1",
			"Methods pop, top and getMin operations will always be called on non-empty stacks",
			"At most 3 * 10^4 calls will be made to push, pop, top, and getMin",
		},
		Examples: []Example{
			{Input: `["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]`, Output: "[null,null,null,null,-3,null,0,-2]", Explanation: "MinStack minStack = new MinStack(); minStack.push(-2); minStack.push(0); minStack.push(-3); minStack.getMin(); // return -3; minStack.pop(); minStack.top(); // return 0; minStack.getMin(); // return -2"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}, "values": [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, -3, nil, 0, -2}},
		},
		TimeComplexity:  "O(1)",
		SpaceComplexity: "O(n)",
		StarterCode: `class MinStack:
    def __init__(self):
        # Initialize your data structure here
        pass

    def push(self, val: int) -> None:
        pass

    def pop(self) -> None:
        pass

    def top(self) -> int:
        pass

    def getMin(self) -> int:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Store additional information with each element to track the minimum at that point in the stack.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "Use a stack of tuples (value, current_min). Each push stores the value along with the minimum of all elements at or below it."},
			{Level: 3, Type: "code", Content: "On push: new_min = min(val, self.stack[-1][1]) if stack else val. Store (val, new_min). getMin returns stack[-1][1]."},
		},
		Solution: Solution{
			Code: `class MinStack:
    def __init__(self):
        self.stack = []  # Each element: (value, min_so_far)

    def push(self, val: int) -> None:
        if not self.stack:
            self.stack.append((val, val))
        else:
            current_min = min(val, self.stack[-1][1])
            self.stack.append((val, current_min))

    def pop(self) -> None:
        self.stack.pop()

    def top(self) -> int:
        return self.stack[-1][0]

    def getMin(self) -> int:
        return self.stack[-1][1]`,
			Explanation:     "Store (value, minimum) pairs. Each element knows the minimum of all elements at or below it in the stack.",
			TimeComplexity:  "O(1) for all operations",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize stack", Explanation: "Create empty list to store (value, min) tuples", CodeSnippet: "self.stack = []", LineStart: 3, LineEnd: 3},
				{Title: "Push with min tracking", Explanation: "Calculate new minimum by comparing with previous min", CodeSnippet: "current_min = min(val, self.stack[-1][1])", LineStart: 8, LineEnd: 9},
				{Title: "O(1) getMin", Explanation: "Minimum is always at top of stack's second element", CodeSnippet: "return self.stack[-1][1]", LineStart: 16, LineEnd: 16},
			},
		},
	},
	{
		ID:              "evaluate-reverse-polish-notation",
		Number:          41,
		Title:           "Evaluate Reverse Polish Notation",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Array", "Math", "Stack"},
		RelatedChapters: []int{3, 8},
		Description: `You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:
- The valid operators are '+', '-', '*', and '/'.
- Each operand may be an integer or another expression.
- The division between two integers always truncates toward zero.
- There will not be any division by zero.
- The input represents a valid arithmetic expression in reverse polish notation.
- The answer and all the intermediate calculations can be represented in a 32-bit integer.`,
		Constraints: []string{
			"1 <= tokens.length <= 10^4",
			"tokens[i] is either an operator: '+', '-', '*', or '/', or an integer in the range [-200, 200]",
		},
		Examples: []Example{
			{Input: `tokens = ["2","1","+","3","*"]`, Output: "9", Explanation: "((2 + 1) * 3) = 9"},
			{Input: `tokens = ["4","13","5","/","+"]`, Output: "6", Explanation: "(4 + (13 / 5)) = 6"},
			{Input: `tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]`, Output: "22"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"tokens": []string{"2", "1", "+", "3", "*"}}, Expected: 9},
			{Input: map[string]interface{}{"tokens": []string{"4", "13", "5", "/", "+"}}, Expected: 6},
			{Input: map[string]interface{}{"tokens": []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}}, Expected: 22},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def evalRPN(tokens):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a stack. Numbers get pushed; operators pop two operands, compute, and push the result."},
			{Level: 2, Type: "algorithm", Content: "For each token: if it's a number, push it. If it's an operator, pop two numbers, apply the operator, push the result."},
			{Level: 3, Type: "code", Content: "Watch out for order: b = stack.pop(), a = stack.pop(), then compute a op b. For division, use int(a/b) to truncate toward zero."},
		},
		Solution: Solution{
			Code: `def evalRPN(tokens):
    stack = []
    operators = {'+', '-', '*', '/'}

    for token in tokens:
        if token in operators:
            b = stack.pop()
            a = stack.pop()
            if token == '+':
                stack.append(a + b)
            elif token == '-':
                stack.append(a - b)
            elif token == '*':
                stack.append(a * b)
            else:  # division
                stack.append(int(a / b))  # truncate toward zero
        else:
            stack.append(int(token))

    return stack[0]`,
			Explanation:     "Process tokens left to right. Numbers go on stack. Operators pop two operands, compute, push result.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize stack", Explanation: "Stack will hold intermediate results", CodeSnippet: "stack = []", LineStart: 2, LineEnd: 2},
				{Title: "Process each token", Explanation: "Check if token is operator or number", CodeSnippet: "if token in operators:", LineStart: 6, LineEnd: 6},
				{Title: "Apply operator", Explanation: "Pop two operands (order matters!), apply operator, push result", CodeSnippet: "b = stack.pop()\na = stack.pop()", LineStart: 7, LineEnd: 8},
				{Title: "Handle division", Explanation: "Truncate toward zero using int(a/b)", CodeSnippet: "stack.append(int(a / b))", LineStart: 16, LineEnd: 16},
			},
		},
	},
	{
		ID:              "generate-parentheses",
		Number:          42,
		Title:           "Generate Parentheses",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"String", "Dynamic Programming", "Backtracking"},
		RelatedChapters: []int{3, 4, 12},
		Description: `Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.`,
		Constraints: []string{
			"1 <= n <= 8",
		},
		Examples: []Example{
			{Input: "n = 3", Output: `["((()))","(()())","(())()","()(())","()()()"]`},
			{Input: "n = 1", Output: `["()"]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 3}, Expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
			{Input: map[string]interface{}{"n": 1}, Expected: []string{"()"}},
			{Input: map[string]interface{}{"n": 2}, Expected: []string{"(())", "()()"}},
		},
		TimeComplexity:  "O(4^n / sqrt(n))",
		SpaceComplexity: "O(n)",
		StarterCode:     "def generateParenthesis(n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use backtracking. At each step, you can add '(' if open count < n, or ')' if close count < open count.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "Track open and close counts. Only add ')' when close < open to ensure validity. Base case: both counts equal n."},
			{Level: 3, Type: "code", Content: "def backtrack(s, open, close): if open == close == n: result.append(s). Add '(' if open < n, add ')' if close < open."},
		},
		Solution: Solution{
			Code: `def generateParenthesis(n):
    result = []

    def backtrack(current, open_count, close_count):
        if len(current) == 2 * n:
            result.append(current)
            return

        if open_count < n:
            backtrack(current + '(', open_count + 1, close_count)
        if close_count < open_count:
            backtrack(current + ')', open_count, close_count + 1)

    backtrack('', 0, 0)
    return result`,
			Explanation:     "Backtracking with constraints: add '(' if open < n, add ')' only if close < open to maintain validity.",
			TimeComplexity:  "O(4^n / sqrt(n)) - nth Catalan number",
			SpaceComplexity: "O(n) - recursion depth",
			Walkthrough: []WalkthroughStep{
				{Title: "Base case", Explanation: "When string length is 2n, we have a valid combination", CodeSnippet: "if len(current) == 2 * n:", LineStart: 5, LineEnd: 7},
				{Title: "Add open paren", Explanation: "Can add '(' if we haven't used all n opens", CodeSnippet: "if open_count < n:", LineStart: 9, LineEnd: 10},
				{Title: "Add close paren", Explanation: "Can add ')' only if close < open (ensures validity)", CodeSnippet: "if close_count < open_count:", LineStart: 11, LineEnd: 12},
			},
		},
	},
	{
		ID:              "daily-temperatures",
		Number:          43,
		Title:           "Daily Temperatures",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Monotonic Stack"},
		RelatedChapters: []int{3, 8},
		Description: `Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.`,
		Constraints: []string{
			"1 <= temperatures.length <= 10^5",
			"30 <= temperatures[i] <= 100",
		},
		Examples: []Example{
			{Input: "temperatures = [73,74,75,71,69,72,76,73]", Output: "[1,1,4,2,1,1,0,0]"},
			{Input: "temperatures = [30,40,50,60]", Output: "[1,1,1,0]"},
			{Input: "temperatures = [30,60,90]", Output: "[1,1,0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"temperatures": []int{73, 74, 75, 71, 69, 72, 76, 73}}, Expected: []int{1, 1, 4, 2, 1, 1, 0, 0}},
			{Input: map[string]interface{}{"temperatures": []int{30, 40, 50, 60}}, Expected: []int{1, 1, 1, 0}},
			{Input: map[string]interface{}{"temperatures": []int{30, 60, 90}}, Expected: []int{1, 1, 0}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def dailyTemperatures(temperatures):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a monotonic decreasing stack to track indices of temperatures waiting for a warmer day."},
			{Level: 2, Type: "algorithm", Content: "Stack stores indices. When current temp > stack top's temp, pop and calculate the wait time (current index - popped index)."},
			{Level: 3, Type: "code", Content: "while stack and temps[i] > temps[stack[-1]]: idx = stack.pop(); result[idx] = i - idx. Always push current index."},
		},
		Solution: Solution{
			Code: `def dailyTemperatures(temperatures):
    n = len(temperatures)
    result = [0] * n
    stack = []  # Store indices

    for i in range(n):
        while stack and temperatures[i] > temperatures[stack[-1]]:
            idx = stack.pop()
            result[idx] = i - idx
        stack.append(i)

    return result`,
			Explanation:     "Monotonic decreasing stack of indices. When we find a warmer day, pop all colder days and record their wait times.",
			TimeComplexity:  "O(n) - each index pushed and popped at most once",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize result", Explanation: "Default 0 (no warmer day found)", CodeSnippet: "result = [0] * n", LineStart: 3, LineEnd: 3},
				{Title: "Process warmer days", Explanation: "Pop indices with lower temps and calculate wait time", CodeSnippet: "while stack and temperatures[i] > temperatures[stack[-1]]:", LineStart: 7, LineEnd: 9},
				{Title: "Always push current", Explanation: "Current index may need a future warmer day", CodeSnippet: "stack.append(i)", LineStart: 10, LineEnd: 10},
			},
		},
	},
	{
		ID:              "car-fleet",
		Number:          44,
		Title:           "Car Fleet",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Sorting", "Monotonic Stack"},
		RelatedChapters: []int{3, 8, 12},
		Description: `There are n cars going to the same destination along a one-lane road. The destination is target miles away.

You are given two integer arrays position and speed, both of length n, where position[i] is the position of the ith car and speed[i] is the speed of the ith car (in miles per hour).

A car can never pass another car ahead of it, but it can catch up to it and drive bumper to bumper at the same speed. The faster car will slow down to match the slower car's speed. The distance between these two cars is ignored (i.e., they are assumed to have the same position).

A car fleet is some non-empty set of cars driving at the same position and same speed. Note that a single car is also a car fleet.

If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.

Return the number of car fleets that will arrive at the destination.`,
		Constraints: []string{
			"n == position.length == speed.length",
			"1 <= n <= 10^5",
			"0 < target <= 10^6",
			"0 <= position[i] < target",
			"All the values of position are unique",
			"0 < speed[i] <= 10^6",
		},
		Examples: []Example{
			{Input: "target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]", Output: "3", Explanation: "Cars at 10,8,0,5,3 form 3 fleets: [10], [8], [0,5,3] merge at different times."},
			{Input: "target = 10, position = [3], speed = [3]", Output: "1"},
			{Input: "target = 100, position = [0,2,4], speed = [4,2,1]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"target": 12, "position": []int{10, 8, 0, 5, 3}, "speed": []int{2, 4, 1, 1, 3}}, Expected: 3},
			{Input: map[string]interface{}{"target": 10, "position": []int{3}, "speed": []int{3}}, Expected: 1},
			{Input: map[string]interface{}{"target": 100, "position": []int{0, 2, 4}, "speed": []int{4, 2, 1}}, Expected: 1},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def carFleet(target, position, speed):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Sort cars by position (closest to target first). Calculate time to reach target for each car.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "If a car behind takes less time than one ahead, they merge into one fleet (the slower car's time). Count remaining distinct times."},
			{Level: 3, Type: "code", Content: "Sort by position descending. Use a stack of arrival times. If current time <= stack top, it merges. Otherwise, push new fleet."},
		},
		Solution: Solution{
			Code: `def carFleet(target, position, speed):
    # Pair position and speed, sort by position descending
    cars = sorted(zip(position, speed), reverse=True)
    stack = []  # arrival times of fleets

    for pos, spd in cars:
        time = (target - pos) / spd
        # If this car takes longer, it's a new fleet
        if not stack or time > stack[-1]:
            stack.append(time)
        # Otherwise, it merges with the fleet ahead (don't push)

    return len(stack)`,
			Explanation:     "Sort by position (closest first). Track arrival times with stack. Slower cars (longer time) start new fleets; faster cars merge.",
			TimeComplexity:  "O(n log n) for sorting",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort by position", Explanation: "Process cars from closest to target to farthest", CodeSnippet: "cars = sorted(zip(position, speed), reverse=True)", LineStart: 3, LineEnd: 3},
				{Title: "Calculate arrival time", Explanation: "Time = distance / speed", CodeSnippet: "time = (target - pos) / spd", LineStart: 7, LineEnd: 7},
				{Title: "New fleet or merge", Explanation: "If slower than fleet ahead, start new fleet; otherwise merge", CodeSnippet: "if not stack or time > stack[-1]:", LineStart: 9, LineEnd: 10},
			},
		},
	},
	{
		ID:              "largest-rectangle-in-histogram",
		Number:          45,
		Title:           "Largest Rectangle in Histogram",
		Difficulty:      "Hard",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Monotonic Stack"},
		RelatedChapters: []int{3, 8, 12},
		Description: `Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.`,
		Constraints: []string{
			"1 <= heights.length <= 10^5",
			"0 <= heights[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "heights = [2,1,5,6,2,3]", Output: "10", Explanation: "The largest rectangle has area = 10 units (bars at index 2-3 with height 5)."},
			{Input: "heights = [2,4]", Output: "4"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"heights": []int{2, 1, 5, 6, 2, 3}}, Expected: 10},
			{Input: map[string]interface{}{"heights": []int{2, 4}}, Expected: 4},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def largestRectangleArea(heights):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a monotonic increasing stack. When we see a smaller bar, calculate areas for all taller bars that can no longer extend right."},
			{Level: 2, Type: "algorithm", Content: "Stack stores indices. For each bar popped, its right boundary is current index, left boundary is new stack top. Width = right - left - 1."},
			{Level: 3, Type: "code", Content: "Add a sentinel height 0 at end to flush stack. When popping: width = i - stack[-1] - 1, area = heights[idx] * width."},
		},
		Solution: Solution{
			Code: `def largestRectangleArea(heights):
    stack = []  # indices
    max_area = 0
    heights.append(0)  # sentinel to flush stack at end

    for i, h in enumerate(heights):
        while stack and heights[stack[-1]] > h:
            height = heights[stack.pop()]
            width = i if not stack else i - stack[-1] - 1
            max_area = max(max_area, height * width)
        stack.append(i)

    heights.pop()  # remove sentinel
    return max_area`,
			Explanation:     "Monotonic increasing stack. When a shorter bar appears, pop taller bars and calculate their max possible rectangle area.",
			TimeComplexity:  "O(n) - each bar pushed and popped once",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Add sentinel", Explanation: "Ensures all bars are processed by adding height 0 at end", CodeSnippet: "heights.append(0)", LineStart: 4, LineEnd: 4},
				{Title: "Pop taller bars", Explanation: "Calculate area when we can't extend right anymore", CodeSnippet: "while stack and heights[stack[-1]] > h:", LineStart: 7, LineEnd: 10},
				{Title: "Calculate width", Explanation: "Width extends from previous stack element to current index", CodeSnippet: "width = i if not stack else i - stack[-1] - 1", LineStart: 9, LineEnd: 9},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, StackProblems...)
}
