package practice

// StackProblems contains all stack category problems, in ladder order:
// plain stack matching, a stack that carries extra state, a stack as an
// evaluator, then monotonic stacks (next warmer day, fleets, histogram), and
// finally Generate Parentheses, where the stack shrinks to a counter.
var StackProblems = []*Problem{
	{
		ID:              "valid-parentheses",
		Number:          20,
		Order:           1,
		Title:           "Valid Parentheses",
		Slug:            "valid-parentheses",
		Difficulty:      "Easy",
		Category:        "stack",
		Tags:            []string{"String", "Stack"},
		RelatedChapters: []int{5, 2},
		Description: `You are given a string s made only of the six bracket characters '(', ')', '[', ']', '{' and '}'. Decide whether the brackets are balanced and return true or false.

The string is balanced when all three rules hold:
1. Every closing bracket has an opening bracket of the same type before it.
2. Brackets close in the reverse order they were opened, so "([)]" is not balanced even though the counts match: the ')' arrives while '[' is the most recent unclosed bracket.
3. No opening bracket is left unclosed at the end.`,
		Constraints: []string{
			"1 <= s.length <= 10^4",
			"s contains only the characters ()[]{}",
		},
		Examples: []Example{
			{ID: "1", Input: `s = "()[]{}"`, Output: "true", Explanation: "Each opener is closed immediately by its own type, so nothing is left waiting."},
			{ID: "2", Input: `s = "([)]"`, Output: "false", Explanation: "When ')' arrives, the most recent unclosed bracket is '[', not '('. The counts match, but the order doesn't."},
			{ID: "3", Input: `s = "{[]}"`, Output: "true", Explanation: "'[' is opened last, so it must close first; then '}' closes '{'."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "()[]{}"}, Expected: true, Note: "three pairs side by side"},
			{Input: map[string]interface{}{"s": "([)]"}, Expected: false, Note: "right counts, wrong closing order: counting brackets is not enough"},
			{Input: map[string]interface{}{"s": "{[]}"}, Expected: true, Note: "nested pairs close innermost first"},
			{Input: map[string]interface{}{"s": "(]"}, Expected: false, Note: "closer of the wrong type"},
			{Input: map[string]interface{}{"s": "("}, Expected: false, Hidden: true, Note: "single opener never closed: check the stack is empty at the end"},
			{Input: map[string]interface{}{"s": "]"}, Expected: false, Hidden: true, Note: "closer with an empty stack: check before popping"},
			{Input: map[string]interface{}{"s": "(("}, Expected: false, Hidden: true, Note: "only openers"},
			{Input: map[string]interface{}{"s": "){"}, Expected: false, Hidden: true, Note: "closer first, then an opener"},
			{Input: map[string]interface{}{"s": "([]{()})"}, Expected: true, Hidden: true, Note: "deeper mixed nesting"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def isValid(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When a closing bracket arrives, which earlier opening bracket is the only one it is allowed to close?"},
			{Level: 2, Type: "insight", Content: "A closer must match the most recently opened bracket that is still unclosed. \"Most recent first\" is last-in-first-out, which is exactly what a stack gives you: push openers, and on a closer compare it with the top. See Chapter 5, Stacks and Monotonic Stacks.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Decide what the stack holds: the opening brackets still waiting to be closed. Keep a map from each closer to its opener.\n2. Pop what the new item resolves: on a closer, fail if the stack is empty or its top is a different opener; otherwise pop it.\n3. Push the new item: on an opener, push it.\n4. Handle what is left on the stack: the string is balanced only if the stack is empty at the end."},
			{Level: 4, Type: "skeleton", Content: `def isValid(s):
    closer_to_opener = {')': '(', ']': '[', '}': '{'}
    stack = []
    for ch in s:
        if ch in closer_to_opener:
            if ____:
                return False
            stack.pop()
        else:
            ____
    return ____`},
		},
		Solution: Solution{
			Code: `def isValid(s):
    closer_to_opener = {')': '(', ']': '[', '}': '{'}
    stack = []  # openers still waiting for their closer
    for ch in s:
        if ch in closer_to_opener:
            # Pop what the new item resolves: it must close the latest opener
            if not stack or stack[-1] != closer_to_opener[ch]:
                return False
            stack.pop()
        else:
            # Push the new item: this opener now waits for its closer
            stack.append(ch)
    # Handle what is left: any opener still here was never closed
    return not stack`,
			Explanation:     "Scan left to right, keeping a stack of the opening brackets that are still unclosed. Each closer must match the top of the stack, which is the most recently opened bracket; if it does, pop it. At the end the stack must be empty. This is the plainest use of a stack in the ladder: every later problem keeps the same push / pop-on-resolve / check-leftovers rhythm.",
			BruteForce:      "Repeatedly delete every adjacent pair \"()\", \"[]\" or \"{}\" until nothing changes, then check whether the string is empty. Each pass is O(n) and can remove as little as one pair (think \"((((...))))\"), so it takes up to n/2 passes: O(n^2). The waste is rescanning the whole string to find the one pair that just became adjacent, when the stack already knows it: it is the top.",
			KeyInsight:      "A closer can only ever match the most recent unclosed opener. Keeping unclosed openers on a stack puts that opener on top, so each character is handled with one O(1) push or pop.",
			WhyItWorks:      "Invariant: after reading a prefix of s, the stack holds exactly the openers in that prefix that are still unclosed, in the order they were opened (oldest at the bottom). A closer in a balanced string must close the newest unclosed opener, because closing an older one would leave a newer one crossing it, like \"([)]\". So comparing with the top is both necessary and sufficient, and popping keeps the invariant. At the end, a non-empty stack means some opener was never closed.",
			CommonMistakes:  []string{"Counting each bracket type instead of using a stack. Counts accept \"([)]\". Fix: the order matters, so compare each closer with the most recent opener.", "Reading stack[-1] when the stack is empty, which crashes on \"]\". Fix: check \"not stack\" before looking at the top.", "Returning True after the loop without checking leftovers, which accepts \"((\". Fix: return \"not stack\"."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "The stack holds openers that are still waiting for their closer. The dictionary maps each closer to the opener it needs, so one lookup tells us both whether a character is a closer and what it must match.", CodeSnippet: "closer_to_opener = {')': '(', ']': '[', '}': '{'}\nstack = []  # openers still waiting for their closer", LineStart: 2, LineEnd: 3},
				{Title: "Pop what the new item resolves", Explanation: "A closer resolves the newest waiting opener. If there is none, or it is the wrong type, the string is invalid right away. For \"([)]\", when ')' arrives the top is '[', so we return False.", CodeSnippet: "if not stack or stack[-1] != closer_to_opener[ch]:\n    return False\nstack.pop()", LineStart: 5, LineEnd: 9},
				{Title: "Push the new item", Explanation: "An opener can't be checked yet; it waits on the stack until its closer arrives.", CodeSnippet: "stack.append(ch)", LineStart: 10, LineEnd: 12},
				{Title: "Handle what is left on the stack", Explanation: "Anything still on the stack was opened and never closed, so the answer is True only for an empty stack. This catches \"((\".", CodeSnippet: "return not stack", LineStart: 13, LineEnd: 14},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: delete adjacent pairs until nothing changes",
					Code: `def isValid(s):
    previous = None
    while previous != s:
        previous = s
        s = s.replace('()', '').replace('[]', '').replace('{}', '')
    return s == ''`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n)",
					Explanation:     "A balanced string always contains an adjacent matching pair, and deleting it leaves a balanced string, so repeated deletion empties exactly the balanced strings. It is correct but rescans the whole string on every pass.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does \"([)]\" fail even though it has one of each bracket type opened and closed?", Answer: "When ')' arrives, the most recent unclosed opener is '[', so ')' would have to close '(' while '[' is still open inside it. The two pairs cross instead of nesting. The stack catches this because its top is '[', not '('."},
			{Prompt: "What if the string had only one bracket type, '(' and ')'? Could you drop the stack?", Answer: "Yes. With one type, every stack entry is the same character, so only the stack's height matters. Keep a counter: +1 for '(', -1 for ')', fail if it goes below 0, and require 0 at the end. That is O(1) space. Generate Parentheses at the end of this ladder uses exactly this counter."},
		},
	},
	{
		ID:              "min-stack",
		Number:          155,
		Order:           2,
		Title:           "Min Stack",
		Slug:            "min-stack",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Stack", "Design"},
		RelatedChapters: []int{5},
		Description: `Build a class MinStack that behaves like a normal stack of integers but can also report its smallest element. Every method must run in O(1) time.

- MinStack() creates an empty stack.
- push(val) puts val on top.
- pop() removes the top element (it returns nothing).
- top() returns the top element without removing it.
- getMin() returns the smallest element currently in the stack.

pop, top and getMin are only ever called when the stack is not empty.

Input format for the tests: "operations" lists the calls in order, starting with the constructor, and "values" lists the arguments for each call. The expected output has one entry per call, with null for calls that return nothing.`,
		Constraints: []string{
			"-2^31 <= val <= 2^31 - 1",
			"pop, top and getMin are called only on a non-empty stack",
			"At most 3 * 10^4 calls in total",
		},
		Examples: []Example{
			{ID: "1", Input: `operations = ["MinStack","push","push","push","getMin","pop","top","getMin"]
values = [[],[-2],[0],[-3],[],[],[],[]]`, Output: "[null,null,null,null,-3,null,0,-2]", Explanation: "After pushing -2, 0, -3 the minimum is -3. Popping removes -3, the top becomes 0, and the minimum goes back to -2, the smallest of what remains."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}, "values": [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, -3, nil, 0, -2}, Note: "the example: popping the minimum restores the previous one"},
			{Input: map[string]interface{}{"operations": []string{"MinStack", "push", "push", "getMin", "push", "push", "getMin", "pop", "getMin", "pop", "top"}, "values": [][]int{{}, {3}, {5}, {}, {2}, {1}, {}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, 3, nil, nil, 1, nil, 2, nil, 5}, Note: "minimum drops twice, then climbs back one pop at a time"},
			{Input: map[string]interface{}{"operations": []string{"MinStack", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"}, "values": [][]int{{}, {0}, {1}, {0}, {}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, 0, nil, 0, nil, 0}, Hidden: true, Note: "the minimum value appears twice: popping one copy must keep the other"},
			{Input: map[string]interface{}{"operations": []string{"MinStack", "push", "push", "pop", "getMin", "top"}, "values": [][]int{{}, {2}, {2}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, 2, 2}, Hidden: true, Note: "equal pushes: a min tracker that only records strictly smaller values breaks here"},
			{Input: map[string]interface{}{"operations": []string{"MinStack", "push", "top", "getMin"}, "values": [][]int{{}, {7}, {}, {}}}, Expected: []interface{}{nil, nil, 7, 7}, Note: "a single element is both the top and the minimum"},
			{Input: map[string]interface{}{"operations": []string{"MinStack", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"}, "values": [][]int{{}, {-2147483648}, {2147483647}, {-1}, {}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, -2147483648, nil, -2147483648, nil, -2147483648}, Hidden: true, Note: "extreme 32-bit values; the minimum is at the bottom the whole time"},
			{Input: map[string]interface{}{"operations": []string{"MinStack", "push", "push", "push", "getMin", "pop", "getMin", "push", "getMin", "top"}, "values": [][]int{{}, {5}, {4}, {3}, {}, {}, {}, {6}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, 3, nil, 4, nil, 4, 6}, Hidden: true, Note: "pushing a larger value after a pop keeps the current minimum"},
		},
		TimeComplexity:  "O(1)",
		SpaceComplexity: "O(n)",
		StarterCode: `class MinStack:
    def __init__(self):
        pass

    def push(self, val: int) -> None:
        pass

    def pop(self) -> None:
        pass

    def top(self) -> int:
        pass

    def getMin(self) -> int:
        pass`,
		Harness: Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When you pop the current minimum, what was the minimum just before it was pushed, and where could you have written that down?"},
			{Level: 2, Type: "insight", Content: "A stack only changes at the top, so the minimum of the elements at or below any entry never changes while that entry is on the stack. Store that minimum alongside each entry when you push it. Popping then automatically reveals the previous minimum. See Chapter 5 (\"store the answer with each entry\").", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Decide what the stack holds: pairs (value, smallest value at or below this entry).\n2. Push the new item: its pair minimum is the smaller of val and the current top's minimum (or val itself if the stack is empty).\n3. Pop: just remove the top pair; the pair below already carries the right minimum.\n4. top and getMin read the first and second parts of the top pair."},
			{Level: 4, Type: "skeleton", Content: `class MinStack:
    def __init__(self):
        self.stack = []  # (value, smallest at or below)

    def push(self, val: int) -> None:
        if self.stack:
            smallest = ____
        else:
            smallest = val
        self.stack.append(____)

    def pop(self) -> None:
        self.stack.pop()

    def top(self) -> int:
        return self.stack[-1][0]

    def getMin(self) -> int:
        return ____`},
		},
		Solution: Solution{
			Code: `class MinStack:
    def __init__(self):
        # Each entry is a pair: (value, smallest value at or below this entry)
        self.stack = []

    def push(self, val: int) -> None:
        if self.stack:
            smallest = min(val, self.stack[-1][1])
        else:
            smallest = val
        self.stack.append((val, smallest))

    def pop(self) -> None:
        self.stack.pop()

    def top(self) -> int:
        return self.stack[-1][0]

    def getMin(self) -> int:
        return self.stack[-1][1]`,
			Explanation:     "Valid Parentheses used the stack only for its order. Here each entry also carries a snapshot: the minimum of everything at or below it. Because a stack changes only at its top, that snapshot stays correct for as long as the entry is on the stack, so getMin is a single read and pop needs no recomputation.",
			BruteForce:      "Keep a plain list and compute min(self.stack) on every getMin call. push, pop and top are O(1), but getMin scans all n elements, O(n). The waste: the minimum of the lower part of the stack is recomputed from scratch every time, although that part hasn't changed since the last call.",
			KeyInsight:      "The elements below the top never change until the top is popped, so the minimum of \"this entry and everything below\" can be computed once at push time and stored with the entry.",
			WhyItWorks:      "Invariant: for every entry i, stack[i][1] == min of the values stack[0][0] .. stack[i][0]. Push keeps it: the new entry's minimum is min(val, previous top's minimum), and min over a prefix extended by one element is exactly that. Pop keeps it: removing the top doesn't change any lower entry or the prefix it covers. So stack[-1][1] is always the minimum of the whole stack. Each method does a constant amount of work.",
			CommonMistakes:  []string{"Keeping a single self.min variable. After popping the minimum you don't know the previous one. Fix: store a minimum per entry (or a second stack of minimums).", "With a separate stack of minimums, pushing only when val < current min (strictly). After push 2, push 2, pop, the min stack is emptied while a 2 remains. Fix: push when val <= current min, or store a minimum with every entry.", "Using min(val, self.stack[-1][1]) without handling the empty stack, which crashes on the first push. Fix: use val itself when the stack is empty."},
			TimeComplexity:  "O(1) for every operation",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "Each entry is a pair (value, smallest at or below). The second part is a snapshot of the answer to getMin for the stack as it was right after this push.", CodeSnippet: "self.stack = []", LineStart: 3, LineEnd: 4},
				{Title: "Push the new item", Explanation: "The new snapshot is the smaller of val and the snapshot below it. Pushing 3, 5, 2 stores (3,3), (5,3), (2,2). An empty stack has nothing below, so the snapshot is val itself.", CodeSnippet: "smallest = min(val, self.stack[-1][1])\nsmallest = val\nself.stack.append((val, smallest))", LineStart: 6, LineEnd: 11},
				{Title: "Pop restores the previous minimum", Explanation: "Removing the top pair exposes the pair below, whose snapshot was correct when it was on top and still is. Popping (2,2) from the stack above makes getMin return 3 again with no work.", CodeSnippet: "self.stack.pop()", LineStart: 13, LineEnd: 14},
				{Title: "Read the answer from the top", Explanation: "top and getMin are single index reads of the top pair, so both are O(1).", CodeSnippet: "return self.stack[-1][0]\nreturn self.stack[-1][1]", LineStart: 16, LineEnd: 20},
			},
			Alternatives: []AltSolution{
				{
					Name: "Two stacks: values and minimums",
					Code: `class MinStack:
    def __init__(self):
        self.values = []
        self.mins = []  # mins[-1] is the current minimum

    def push(self, val: int) -> None:
        self.values.append(val)
        if not self.mins or val <= self.mins[-1]:  # <= keeps duplicate minimums
            self.mins.append(val)

    def pop(self) -> None:
        if self.values.pop() == self.mins[-1]:
            self.mins.pop()

    def top(self) -> int:
        return self.values[-1]

    def getMin(self) -> int:
        return self.mins[-1]`,
					TimeComplexity:  "O(1) for every operation",
					SpaceComplexity: "O(n)",
					Explanation:     "Only values that are a new minimum (or tie it) are recorded on the second stack, which saves space when minimums are rare. The <= is essential: with <, pushing the same minimum twice records it once and the first pop loses it.",
				},
				{
					Name: "Brute force: scan for the minimum",
					Code: `class MinStack:
    def __init__(self):
        self.items = []

    def push(self, val: int) -> None:
        self.items.append(val)

    def pop(self) -> None:
        del self.items[-1]

    def top(self) -> int:
        return self.items[-1]

    def getMin(self) -> int:
        smallest = self.items[0]
        for value in self.items:
            if value < smallest:
                smallest = value
        return smallest`,
					TimeComplexity:  "O(n) for getMin, O(1) otherwise",
					SpaceComplexity: "O(n)",
					Explanation:     "Correct but getMin rescans the whole stack each call, which breaks the O(1) requirement.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to compute the minimum once at push time and never update it, even as other elements are pushed and popped later?", Answer: "A snapshot covers only the entry itself and the entries below it. Later pushes go above it and later pops remove entries above it, so the set of elements the snapshot describes never changes while the entry is on the stack. When the entry itself is popped, its snapshot leaves with it."},
			{Prompt: "What if you also needed getMax in O(1)? What would you change?", Answer: "Store a third field, the largest value at or below each entry, computed the same way: max(val, top's max). Any summary that can be updated from (new value, summary below) works this way, for example the sum or the count of negatives."},
		},
	},
	{
		ID:              "evaluate-reverse-polish-notation",
		Number:          150,
		Order:           3,
		Title:           "Evaluate Reverse Polish Notation",
		Slug:            "evaluate-reverse-polish-notation",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Array", "Math", "Stack"},
		RelatedChapters: []int{5},
		Description: `In Reverse Polish Notation (postfix), an operator is written after its two operands, so "3 4 +" means 3 + 4 and "2 1 + 3 *" means (2 + 1) * 3. No parentheses are needed.

You are given the expression as a list of string tokens. Each token is either one of the operators "+", "-", "*", "/" or an integer such as "13" or "-11" (a minus sign followed by digits is a negative number, not an operator). Return the integer value of the expression.

- Each operator applies to the two values just before it: the earlier one is the left operand, the later one the right operand.
- Division truncates toward zero, so 13 / 5 = 2 and -7 / 2 = -3.
- The expression is always valid, never divides by zero, and every intermediate result fits in a 32-bit signed integer.`,
		Constraints: []string{
			"1 <= tokens.length <= 10^4",
			"tokens[i] is \"+\", \"-\", \"*\", \"/\", or an integer in the range [-200, 200]",
		},
		Examples: []Example{
			{ID: "1", Input: `tokens = ["2","1","+","3","*"]`, Output: "9", Explanation: "\"2 1 +\" gives 3, then \"3 3 *\" gives 9: ((2 + 1) * 3)."},
			{ID: "2", Input: `tokens = ["4","13","5","/","+"]`, Output: "6", Explanation: "13 / 5 truncates to 2, then 4 + 2 = 6."},
			{ID: "3", Input: `tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]`, Output: "22", Explanation: "9 + 3 = 12; 12 * -11 = -132; 6 / -132 truncates to 0; 10 * 0 = 0; 0 + 17 + 5 = 22."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"tokens": []string{"2", "1", "+", "3", "*"}}, Expected: 9, Note: "example 1: two operators"},
			{Input: map[string]interface{}{"tokens": []string{"4", "13", "5", "/", "+"}}, Expected: 6, Note: "example 2: division truncates"},
			{Input: map[string]interface{}{"tokens": []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}}, Expected: 22, Note: "example 3: long expression with a negative number"},
			{Input: map[string]interface{}{"tokens": []string{"18"}}, Expected: 18, Hidden: true, Note: "a single number and no operators"},
			{Input: map[string]interface{}{"tokens": []string{"3", "4", "-"}}, Expected: -1, Hidden: true, Note: "operand order: the first popped value is the RIGHT operand"},
			{Input: map[string]interface{}{"tokens": []string{"-7", "2", "/"}}, Expected: -3, Hidden: true, Note: "negative division truncates toward zero (floor division // would give -4)"},
			{Input: map[string]interface{}{"tokens": []string{"7", "-2", "/"}}, Expected: -3, Hidden: true, Note: "negative divisor truncates toward zero"},
			{Input: map[string]interface{}{"tokens": []string{"2", "-3", "*", "4", "-"}}, Expected: -10, Hidden: true, Note: "\"-3\" is a number, not the minus operator"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def evalRPN(tokens):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When you reach an operator, which two values does it combine, and how recently did you see them?"},
			{Level: 2, Type: "insight", Content: "An operator always uses the two most recent values that haven't been consumed yet, so keep values on a stack: numbers are pushed, an operator pops two, computes, and pushes the result back, where it becomes an operand for a later operator. See Chapter 5.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Decide what the stack holds: values waiting for an operator.\n2. Push the new item: a number token is converted with int() and pushed.\n3. Pop what the new item resolves: an operator pops the right operand first, then the left.\n4. Record the answer: compute left op right (truncating division toward zero) and push the result.\n5. Handle what is left on the stack: one value remains, the answer."},
			{Level: 4, Type: "skeleton", Content: `def evalRPN(tokens):
    stack = []
    for token in tokens:
        if token in ('+', '-', '*', '/'):
            right = ____
            left = ____
            if token == '+':
                stack.append(left + right)
            elif token == '-':
                stack.append(left - right)
            elif token == '*':
                stack.append(left * right)
            else:
                stack.append(____)  # truncate toward zero
        else:
            stack.append(int(token))
    return stack[-1]`},
		},
		Solution: Solution{
			Code: `def evalRPN(tokens):
    stack = []  # values waiting for an operator
    for token in tokens:
        if token in ('+', '-', '*', '/'):
            right = stack.pop()  # pushed last, so it is the right operand
            left = stack.pop()
            if token == '+':
                stack.append(left + right)
            elif token == '-':
                stack.append(left - right)
            elif token == '*':
                stack.append(left * right)
            else:
                stack.append(int(left / right))  # truncate toward zero
        else:
            stack.append(int(token))
    return stack[-1]`,
			Explanation:     "Builds on Valid Parentheses: there an arriving closer resolved the top opener; here an arriving operator resolves the top two values. Numbers wait on the stack; each operator pops its two operands, computes, and pushes the result, which then waits like any other value. After the last token, one value remains: the answer.",
			BruteForce:      "Repeatedly find the first operator in the list, replace it and the two numbers before it with the result, and start over. Each replacement scans and rebuilds the list, O(n), and there are up to n/2 operators, so O(n^2). The waste is searching from the start every time for the operands, which are always the two most recent unconsumed values.",
			KeyInsight:      "In postfix, an operator's operands are exactly the two most recent values not yet consumed, so a stack hands them over in O(1) and the result goes back on top for the next operator.",
			WhyItWorks:      "Invariant: after processing a prefix of tokens, the stack holds the values of the complete sub-expressions in that prefix that haven't been used yet, left to right from bottom to top. A number is a complete sub-expression, so it is pushed. In a valid postfix expression an operator combines the two sub-expressions immediately before it, which are the top two entries, with the deeper one on the left. Replacing them with their value keeps the invariant. A valid expression is one complete sub-expression, so one value remains at the end.",
			CommonMistakes:  []string{"Popping into left first: left = pop(), right = pop() computes 4 - 3 for [\"3\",\"4\",\"-\"]. Fix: the first pop is the right operand.", "Using // for division. Python's // rounds toward negative infinity, so -7 // 2 is -4, but the task wants -3. Fix: int(left / right), which truncates; it's exact here because values fit in 32 bits. For arbitrary big integers, divide the absolute values with // and restore the sign.", "Detecting operators with token[0] in '+-*/', which treats \"-3\" as the minus operator. Fix: compare the whole token with the four operator strings."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Push the new item", Explanation: "Any token that isn't exactly one of the four operators is a number, including \"-11\". It waits on the stack until an operator uses it.", CodeSnippet: "stack.append(int(token))", LineStart: 15, LineEnd: 16},
				{Title: "Pop what the new item resolves", Explanation: "An operator resolves the top two values. The value pushed last is the right operand, so pop it first. For [\"3\",\"4\",\"-\"], right = 4 and left = 3.", CodeSnippet: "right = stack.pop()  # pushed last, so it is the right operand\nleft = stack.pop()", LineStart: 4, LineEnd: 6},
				{Title: "Record the answer at pop time", Explanation: "Compute left op right and push the result, which becomes an operand for a later operator. int(left / right) truncates toward zero: int(-7 / 2) is -3.", CodeSnippet: "stack.append(left - right)\nstack.append(int(left / right))  # truncate toward zero", LineStart: 7, LineEnd: 14},
				{Title: "Handle what is left on the stack", Explanation: "A valid expression reduces to one value, which is the answer. For [\"18\"] that's the only number pushed.", CodeSnippet: "return stack[-1]", LineStart: 17, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: reduce the first operator repeatedly",
					Code: `def evalRPN(tokens):
    items = list(tokens)
    while len(items) > 1:
        i = 0
        while items[i] not in ('+', '-', '*', '/'):
            i += 1
        a, b = int(items[i - 2]), int(items[i - 1])
        if items[i] == '+':
            value = a + b
        elif items[i] == '-':
            value = a - b
        elif items[i] == '*':
            value = a * b
        else:
            value = abs(a) // abs(b)
            if (a < 0) != (b < 0):
                value = -value
        items[i - 2:i + 1] = [str(value)]
    return int(items[0])`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n)",
					Explanation:     "The first operator in a postfix list always has two plain numbers right before it, so it can be evaluated and spliced back as a number. Repeating this shrinks the list to one value. It is correct but rescans from the start each time. Its division divides absolute values and restores the sign, an exact way to truncate toward zero.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the first value you pop be the right operand?", Answer: "The right operand is written later in the token list, so it was pushed later and sits above the left operand. A stack returns the most recent push first. This matters for - and /, which aren't symmetric: [\"3\",\"4\",\"-\"] is 3 - 4 = -1, not 1."},
			{Prompt: "What if the input could also contain a unary operator such as \"neg\" (negate one value)? How would the loop change?", Answer: "A unary operator resolves only the top value: pop one, push its negation. The stack idea is unchanged; each operator pops as many values as it takes and pushes one result."},
		},
	},
	{
		ID:              "daily-temperatures",
		Number:          739,
		Order:           4,
		Title:           "Daily Temperatures",
		Slug:            "daily-temperatures",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Monotonic Stack"},
		RelatedChapters: []int{5},
		Description: `You are given a list temperatures, where temperatures[i] is the temperature on day i. For every day, work out how many days you have to wait after it until a strictly warmer day. If no later day is warmer, the answer for that day is 0.

Return a list answer of the same length, where answer[i] is the number of days from day i to the first later day with a higher temperature.`,
		Constraints: []string{
			"1 <= temperatures.length <= 10^5",
			"30 <= temperatures[i] <= 100",
		},
		Examples: []Example{
			{ID: "1", Input: "temperatures = [73,74,75,71,69,72,76,73]", Output: "[1,1,4,2,1,1,0,0]", Explanation: "Day 2 (75) waits until day 6 (76), 4 days later. Day 3 (71) waits until day 5 (72), 2 days later. Days 6 and 7 never see anything warmer."},
			{ID: "2", Input: "temperatures = [30,40,50,60]", Output: "[1,1,1,0]", Explanation: "Each day is followed directly by a warmer one, except the last."},
			{ID: "3", Input: "temperatures = [30,60,90]", Output: "[1,1,0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"temperatures": []int{73, 74, 75, 71, 69, 72, 76, 73}}, Expected: []int{1, 1, 4, 2, 1, 1, 0, 0}, Note: "example 1: one warm day resolves several waiting days at once"},
			{Input: map[string]interface{}{"temperatures": []int{30, 40, 50, 60}}, Expected: []int{1, 1, 1, 0}, Note: "strictly increasing: every day is resolved by the next"},
			{Input: map[string]interface{}{"temperatures": []int{30, 60, 90}}, Expected: []int{1, 1, 0}, Note: "example 3"},
			{Input: map[string]interface{}{"temperatures": []int{50}}, Expected: []int{0}, Hidden: true, Note: "a single day never gets a warmer day"},
			{Input: map[string]interface{}{"temperatures": []int{90, 80, 70, 60}}, Expected: []int{0, 0, 0, 0}, Hidden: true, Note: "strictly decreasing: everything stays on the stack until the end"},
			{Input: map[string]interface{}{"temperatures": []int{70, 70, 70}}, Expected: []int{0, 0, 0}, Hidden: true, Note: "equal temperatures are not warmer: pop only on strictly greater"},
			{Input: map[string]interface{}{"temperatures": []int{70, 60, 70, 80}}, Expected: []int{3, 1, 1, 0}, Hidden: true, Note: "an equal day must not resolve day 0; only 80 does"},
			{Input: map[string]interface{}{"temperatures": []int{30, 100, 30, 100}}, Expected: []int{1, 0, 1, 0}, Hidden: true, Note: "maximum and minimum temperatures alternating"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def dailyTemperatures(temperatures):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When a warm day arrives, which of the earlier days does it answer, and do those days have anything in common in how their temperatures are ordered?"},
			{Level: 2, Type: "insight", Content: "Keep the days that haven't found a warmer day yet on a stack. Their temperatures can only decrease from bottom to top (a warmer day would already have resolved the colder ones below it), so this is a monotonic decreasing stack. A new day pops every waiting day that is colder than it, and for each popped day, the new day is its answer. See Chapter 5, Monotonic Stacks.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Decide what the stack holds: indices of days still waiting for a warmer day (indices, because the answer is a distance). Start answer as all zeros.\n2. Pop what the new item resolves: while the top day is strictly colder than today, pop it.\n3. Record the answer at pop time: answer[popped] = today - popped.\n4. Push the new item: today now waits for its own warmer day.\n5. Handle what is left on the stack: those days never got warmer, and their answer is already 0."},
			{Level: 4, Type: "skeleton", Content: `def dailyTemperatures(temperatures):
    answer = [0] * len(temperatures)
    stack = []  # indices, temperatures decreasing bottom to top
    for today, temp in enumerate(temperatures):
        while stack and ____:
            colder_day = stack.pop()
            answer[colder_day] = ____
        ____
    return answer`},
		},
		Solution: Solution{
			Code: `def dailyTemperatures(temperatures):
    answer = [0] * len(temperatures)
    stack = []  # indices of days still waiting; temperatures decrease bottom to top
    for today, temp in enumerate(temperatures):
        while stack and temperatures[stack[-1]] < temp:
            colder_day = stack.pop()
            answer[colder_day] = today - colder_day
        stack.append(today)
    return answer`,
			Explanation:     "The first monotonic stack in the ladder. Like Valid Parentheses, items wait on the stack until a later item resolves them; the new idea is that one arriving day can resolve several waiting days at once, and the waiting days are always in decreasing temperature order, so the ones it resolves are all at the top.",
			BruteForce:      "For each day i, scan forward from i + 1 until a warmer day appears. On a decreasing list like [90, 80, 70, ...] every scan runs to the end, so it's about n^2/2 comparisons: 5 billion for n = 100,000. The waste: a cold day is compared with the same later days again and again by every earlier day.",
			KeyInsight:      "A day that hasn't found a warmer day yet is still \"waiting\". Waiting days are always in decreasing order of temperature, so when a warm day arrives, the days it answers are exactly the ones on top of the stack, and each day is pushed and popped once.",
			WhyItWorks:      "Invariant: before processing day t, the stack holds, bottom to top, every earlier day that has no warmer day before t, and their temperatures are non-increasing. When day t pops day d, d is on the stack, so no day between d and t was warmer than d (it would have popped d). Hence t is the FIRST warmer day after d and t - d is correct. Days left at the end never saw a warmer day, so 0 is correct. Each index is pushed once and popped at most once, so the total work is O(n) even though there is a loop inside a loop.",
			CommonMistakes:  []string{"Popping on <= instead of <: an equal temperature is not warmer, so [70, 60, 70, 80] would wrongly give day 0 the answer 2. Fix: pop only while the top is strictly colder.", "Pushing temperatures instead of indices. You then can't compute how many days apart they are. Fix: push indices and look up temperatures[stack[-1]].", "Thinking the nested while makes it O(n^2). Fix: count pushes and pops; each index is pushed once and popped at most once, so at most 2n stack operations happen."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "The stack holds indices of days still waiting for a warmer day; answer starts at 0 for everyone, which is already right for days that never get resolved.", CodeSnippet: "answer = [0] * len(temperatures)\nstack = []  # indices of days still waiting; temperatures decrease bottom to top", LineStart: 2, LineEnd: 3},
				{Title: "Pop what the new item resolves", Explanation: "While the top day is strictly colder than today, today is its first warmer day. In example 1, day 5 (72) pops day 4 (69) and then day 3 (71), and stops at day 2 (75).", CodeSnippet: "while stack and temperatures[stack[-1]] < temp:\n    colder_day = stack.pop()", LineStart: 5, LineEnd: 6},
				{Title: "Record the answer at pop time", Explanation: "The wait is the distance in days: answer[4] = 5 - 4 = 1 and answer[3] = 5 - 3 = 2.", CodeSnippet: "answer[colder_day] = today - colder_day", LineStart: 7, LineEnd: 7},
				{Title: "Push the new item", Explanation: "Today waits for its own warmer day. Everything left below it is at least as warm, so the decreasing order is kept.", CodeSnippet: "stack.append(today)", LineStart: 8, LineEnd: 8},
				{Title: "Handle what is left on the stack", Explanation: "Days still on the stack (days 6 and 7 in example 1) never found a warmer day; their answer stays 0 from the initial list, so there is nothing more to do.", CodeSnippet: "return answer", LineStart: 9, LineEnd: 9},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: scan forward from each day",
					Code: `def dailyTemperatures(temperatures):
    n = len(temperatures)
    answer = []
    for i in range(n):
        wait = 0
        for j in range(i + 1, n):
            if temperatures[j] > temperatures[i]:
                wait = j - i
                break
        answer.append(wait)
    return answer`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1) extra",
					Explanation:     "Directly follows the definition. Fine for small inputs, but a long run of falling temperatures makes every day scan to the end.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can the temperatures on the stack never increase from bottom to top?", Answer: "When a day is pushed, every colder day on top was just popped, so the new top is no warmer than the day below it. A warmer day sitting above a colder one would mean the warmer day arrived without popping it, which the while loop prevents."},
			{Prompt: "What if you wanted the number of days since the most recent PREVIOUS warmer day instead? What would you change?", Answer: "Either scan from right to left with the same code, or scan left to right and, after popping every day that is not warmer than today, read the stack's top: it is the nearest earlier day that is warmer, so the answer is today - stack[-1] (or 0 if the stack is empty). The stack is still decreasing; the answer is read at push time instead of pop time."},
		},
	},
	{
		ID:              "car-fleet",
		Number:          853,
		Order:           5,
		Title:           "Car Fleet",
		Slug:            "car-fleet",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Sorting", "Monotonic Stack"},
		RelatedChapters: []int{5, 2},
		Description: `n cars drive along a single-lane road toward a destination target miles from the start. Car i starts at position[i] (miles from the start) and drives at speed[i] miles per hour.

Cars can't overtake. When a faster car catches up with a slower car (or group of cars) ahead, it slows down and drives bumper to bumper with it from then on; treat the two as being at the same position. A car fleet is a group of cars driving together like this; a lone car is also a fleet. If a car catches up with a fleet exactly at the destination, it still counts as part of that fleet.

Return how many fleets arrive at the destination.`,
		Constraints: []string{
			"n == position.length == speed.length",
			"1 <= n <= 10^5",
			"0 < target <= 10^6",
			"0 <= position[i] < target",
			"All values in position are distinct",
			"0 < speed[i] <= 10^6",
		},
		Examples: []Example{
			{ID: "1", Input: "target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]", Output: "3", Explanation: "Alone, the cars would arrive after (12-10)/2 = 1, (12-8)/4 = 1, 12/1 = 12, (12-5)/1 = 7 and (12-3)/3 = 3 hours. The car at 8 catches the car at 10 exactly at the target: fleet [10, 8]. The car at 3 (3 hours) catches the car at 5 (7 hours) and is held to 7 hours: fleet [5, 3]. The car at 0 takes 12 hours and never catches anyone: fleet [0]. That's 3 fleets."},
			{ID: "2", Input: "target = 10, position = [3], speed = [3]", Output: "1", Explanation: "A single car is one fleet."},
			{ID: "3", Input: "target = 100, position = [0,2,4], speed = [4,2,1]", Output: "1", Explanation: "The car at 4 needs 96 hours; the others would be faster, so both catch up with it and all three arrive together."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"target": 12, "position": []int{10, 8, 0, 5, 3}, "speed": []int{2, 4, 1, 1, 3}}, Expected: 3, Note: "example 1: input not sorted by position; fleets [10,8], [5,3], [0]"},
			{Input: map[string]interface{}{"target": 10, "position": []int{3}, "speed": []int{3}}, Expected: 1, Note: "a single car"},
			{Input: map[string]interface{}{"target": 100, "position": []int{0, 2, 4}, "speed": []int{4, 2, 1}}, Expected: 1, Note: "the front car is slowest, so everyone joins it"},
			{Input: map[string]interface{}{"target": 10, "position": []int{6, 8}, "speed": []int{2, 1}}, Expected: 1, Hidden: true, Note: "both arrive after exactly 2 hours: catching up at the target still merges (use <=, not <)"},
			{Input: map[string]interface{}{"target": 12, "position": []int{7, 1}, "speed": []int{2, 4}}, Expected: 2, Hidden: true, Note: "times 2.5 and 2.75 differ by less than 1: integer division would wrongly merge them"},
			{Input: map[string]interface{}{"target": 10, "position": []int{0, 4, 2}, "speed": []int{1, 1, 1}}, Expected: 3, Hidden: true, Note: "equal speeds never catch up: every car is its own fleet"},
			{Input: map[string]interface{}{"target": 20, "position": []int{6, 2, 17}, "speed": []int{3, 9, 2}}, Expected: 2, Hidden: true, Note: "the back car catches the middle car, not the front one"},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def carFleet(target, position, speed):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "A car can only be slowed down by cars in front of it. If you look at the cars starting from the one closest to the target, what single number about the fleet just ahead decides whether the next car joins it?"},
			{Level: 2, Type: "insight", Content: "Compute each car's arrival time if it drove alone: (target - position) / speed. Going from the front car backwards, a car joins the fleet directly ahead if its own time is <= that fleet's arrival time (it would get there no later, so it catches up by the target). Otherwise it starts a new, slower fleet. The fleet arrival times form a monotonic stack that strictly increases from bottom to top. See Chapter 5.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Decide what the stack holds: the arrival time of each fleet, front fleet at the bottom. Pair positions with speeds and sort by position, closest to the target first.\n2. Compute each car's solo arrival time with real (not integer) division.\n3. Merge with the fleet on top: if the time is <= the top's time, the car joins that fleet and nothing is pushed.\n4. Push the new item: otherwise the car is a new fleet; push its time.\n5. Handle what is left on the stack: the answer is the number of entries."},
			{Level: 4, Type: "skeleton", Content: `def carFleet(target, position, speed):
    cars = sorted(zip(position, speed), reverse=True)
    stack = []
    for pos, spd in cars:
        arrival = ____
        if stack and ____:
            continue
        ____
    return len(stack)`},
		},
		Solution: Solution{
			Code: `def carFleet(target, position, speed):
    cars = sorted(zip(position, speed), reverse=True)  # closest to target first
    stack = []  # arrival time of each fleet, front fleet at the bottom
    for pos, spd in cars:
        arrival = (target - pos) / spd
        if stack and arrival <= stack[-1]:
            continue  # catches the fleet ahead by the target: joins it
        stack.append(arrival)  # slower than every fleet ahead: a new fleet
    return len(stack)`,
			Explanation:     "Builds on Daily Temperatures: again a stack of values kept in monotonic order, here fleet arrival times that strictly increase from the front fleet to the back. Sorting by position (a Chapter 2 tool) lets us meet cars in the order they could block each other. A car that would arrive no later than the fleet directly ahead merges into it; a slower car starts a new fleet on top.",
			BruteForce:      "For each car, look at every car in front of it and compute the latest solo arrival time among them; the car leads its own fleet only if it is slower than all of them. That is O(n^2) comparisons. The waste: every car rescans the same cars ahead, although only one number about them matters, the arrival time of the slowest fleet ahead, which is the stack's top.",
			KeyInsight:      "A car's fate depends only on the fleet directly ahead of it, and that fleet's arrival time is the maximum solo time of all cars ahead. Processing cars front to back keeps that maximum on top of the stack.",
			WhyItWorks:      "A fleet's arrival time is the solo time of its leader (the front car), because the cars behind it are held to its speed. Process cars from closest to farthest. Invariant: the stack holds the arrival times of the fleets formed so far, strictly increasing from bottom (front) to top (back), so the top is the time of the fleet directly ahead of the next car and the slowest so far. If the next car's solo time is <= the top, it would reach the target no later, so it catches that fleet at or before the target and joins it. If it's larger, it can't catch the fleet ahead (it is slower to the target), and since the top is the largest time ahead, it can't catch any fleet; it becomes a new fleet with a larger time, keeping the order. The stack size is the fleet count. Sorting costs O(n log n); the scan is O(n).",
			CommonMistakes:  []string{"Forgetting to sort: position isn't given in order, so example 1 would compare unrelated cars. Fix: sort (position, speed) pairs by position, descending.", "Using < instead of <= so a car that catches up exactly at the target becomes a separate fleet. Fix: equal times merge, as the statement says.", "Computing times with // (integer division). 2.5 and 2.75 hours both become 2 and merge wrongly. Fix: use / (or compare (target - p1) * s2 with (target - p2) * s1 in integers)."},
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "Pair each position with its speed and sort closest-to-target first, so each car is processed after every car that could block it. The stack holds fleet arrival times. Example 1 sorts to (10,2), (8,4), (5,1), (3,3), (0,1).", CodeSnippet: "cars = sorted(zip(position, speed), reverse=True)  # closest to target first\nstack = []  # arrival time of each fleet, front fleet at the bottom", LineStart: 2, LineEnd: 3},
				{Title: "Compute the arrival time", Explanation: "Solo arrival time = remaining distance / speed, with real division. The five cars give 1.0, 1.0, 7.0, 3.0 and 12.0 hours.", CodeSnippet: "arrival = (target - pos) / spd", LineStart: 5, LineEnd: 5},
				{Title: "Merge with the fleet on top", Explanation: "If the car would arrive no later than the fleet directly ahead, it catches it and joins; nothing is pushed. Car 8 (1.0) merges with car 10 (1.0) right at the target; car 3 (3.0) merges with car 5 (7.0).", CodeSnippet: "if stack and arrival <= stack[-1]:\n    continue  # catches the fleet ahead by the target: joins it", LineStart: 6, LineEnd: 7},
				{Title: "Push the new item", Explanation: "A car slower than the fleet ahead starts a new fleet. Car 5 pushes 7.0 and car 0 pushes 12.0, so the stack is [1.0, 7.0, 12.0], strictly increasing.", CodeSnippet: "stack.append(arrival)  # slower than every fleet ahead: a new fleet", LineStart: 8, LineEnd: 8},
				{Title: "Handle what is left on the stack", Explanation: "Each entry is one fleet: [1.0, 7.0, 12.0] gives 3, the fleets [10,8], [5,3] and [0].", CodeSnippet: "return len(stack)", LineStart: 9, LineEnd: 9},
			},
			Alternatives: []AltSolution{
				{
					Name: "Monotonic stack with pops, back to front",
					Code: `def carFleet(target, position, speed):
    cars = sorted(zip(position, speed))  # farthest from target first
    stack = []  # fleet arrival times, decreasing bottom to top
    for pos, spd in cars:
        arrival = (target - pos) / spd
        while stack and stack[-1] <= arrival:
            stack.pop()  # that faster fleet behind gets stuck behind this car
        stack.append(arrival)
    return len(stack)`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The classic monotonic-stack shape: process cars from the back, and a new car (which is ahead of everything on the stack) pops every fleet behind it that would arrive no later, because those fleets get stuck behind it. Each car is pushed and popped at most once.",
				},
				{
					Name: "Brute force: compare with every car ahead",
					Code: `def carFleet(target, position, speed):
    n = len(position)
    fleets = 0
    for i in range(n):
        leads = True
        for j in range(n):
            # car j is ahead and arrives no earlier: car i gets stuck behind it
            if position[j] > position[i] and (target - position[j]) * speed[i] >= (target - position[i]) * speed[j]:
                leads = False
                break
        if leads:
            fleets += 1
    return fleets`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "A car leads its own fleet exactly when it is strictly slower to the target than every car ahead of it. Comparing times by cross-multiplying keeps everything in exact integers. No sorting, but every car scans all others.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it enough to compare a car only with the top of the stack, not with every fleet ahead?", Answer: "The stack's times strictly increase toward the top, so the top is the slowest fleet ahead and also the one directly in front. If the car can't catch the top (its time is larger), it can't catch any faster fleet further ahead either; if it can, it reaches the top fleet first because nothing slower sits between them."},
			{Prompt: "In example 1, the car at 3 is faster than the car at 5. Why doesn't its arrival time of 3 hours count?", Answer: "It catches the car at 5 before the target and is held to that car's speed, so it actually arrives at 7 hours, with the car at 5. That's why it merges and the fleet's time stays 7.0: a fleet's time is its leader's time."},
		},
	},
	{
		ID:              "largest-rectangle-in-histogram",
		Number:          84,
		Order:           6,
		Title:           "Largest Rectangle in Histogram",
		Slug:            "largest-rectangle-in-histogram",
		Difficulty:      "Hard",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Monotonic Stack"},
		RelatedChapters: []int{5, 3},
		Description: `A histogram is drawn as a row of bars, each 1 unit wide, standing side by side on a common baseline; bar i has height heights[i].

Find the largest area of a rectangle that fits entirely inside the histogram. The rectangle's sides must be axis-aligned and it must sit on the baseline, so a rectangle covering bars i through j can be at most as tall as the shortest of those bars. Return that area.`,
		Constraints: []string{
			"1 <= heights.length <= 10^5",
			"0 <= heights[i] <= 10^4",
		},
		Examples: []Example{
			{ID: "1", Input: "heights = [2,1,5,6,2,3]", Output: "10", Explanation: "Bars 2 and 3 (heights 5 and 6) hold a rectangle of height 5 and width 2, area 10. No other choice does better: for example, bars 2 to 5 give height 2 and width 4, area 8."},
			{ID: "2", Input: "heights = [2,4]", Output: "4", Explanation: "Either the single bar of height 4 (area 4) or both bars at height 2 (area 4)."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"heights": []int{2, 1, 5, 6, 2, 3}}, Expected: 10, Note: "example 1: the best rectangle is in the middle"},
			{Input: map[string]interface{}{"heights": []int{2, 4}}, Expected: 4, Note: "example 2"},
			{Input: map[string]interface{}{"heights": []int{5}}, Expected: 5, Note: "a single bar"},
			{Input: map[string]interface{}{"heights": []int{0}}, Expected: 0, Hidden: true, Note: "a single bar of height 0"},
			{Input: map[string]interface{}{"heights": []int{1, 2, 3, 4, 5}}, Expected: 9, Hidden: true, Note: "increasing: nothing is popped until the end, so leftover bars must be handled"},
			{Input: map[string]interface{}{"heights": []int{5, 4, 3, 2, 1}}, Expected: 9, Hidden: true, Note: "decreasing: popped bars extend all the way to the left edge"},
			{Input: map[string]interface{}{"heights": []int{3, 3, 3, 3}}, Expected: 12, Hidden: true, Note: "all equal: the whole histogram is the answer"},
			{Input: map[string]interface{}{"heights": []int{2, 1, 2}}, Expected: 3, Hidden: true, Note: "the short middle bar spans the full width"},
			{Input: map[string]interface{}{"heights": []int{6, 2, 5, 4, 5, 1, 6}}, Expected: 12, Hidden: true, Note: "the answer (height 4, width 3) is bounded by lower bars on both sides"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def largestRectangleArea(heights):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose the rectangle's height is exactly heights[i]. How far left and right can it stretch, and what stops it on each side?"},
			{Level: 2, Type: "insight", Content: "The tallest rectangle that uses bar i at full height stretches until the first strictly lower bar on each side. Keep bar indices on a stack whose heights never decrease from bottom to top (a monotonic increasing stack). When a lower bar arrives, it is the right boundary for every taller bar it pops, and the bar left below each popped bar is its left boundary. See Chapter 5, Monotonic Stacks.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Decide what the stack holds: indices of bars whose rectangle can still grow to the right. Append a bar of height 0 so every bar gets popped at the end.\n2. Pop what the new item resolves: while the top bar is taller than the current bar, pop it.\n3. Record the answer at pop time: the popped bar's left boundary is the new top (or -1 if empty), its right boundary is the current index; width = right - left - 1; update the best area.\n4. Push the new item: push the current index.\n5. Handle what is left on the stack: the final 0-height bar already popped everything."},
			{Level: 4, Type: "skeleton", Content: `def largestRectangleArea(heights):
    bars = heights + [0]
    stack = []
    best = 0
    for right, h in enumerate(bars):
        while stack and bars[stack[-1]] > h:
            tall = stack.pop()
            left = ____
            width = ____
            best = max(best, ____)
        stack.append(right)
    return best`},
		},
		Solution: Solution{
			Code: `def largestRectangleArea(heights):
    bars = heights + [0]  # a final 0-height bar resolves every bar still waiting
    stack = []  # indices; heights never decrease from bottom to top
    best = 0
    for right, h in enumerate(bars):
        while stack and bars[stack[-1]] > h:
            tall = stack.pop()
            left = stack[-1] if stack else -1  # nearest lower bar on the left
            width = right - left - 1
            best = max(best, bars[tall] * width)
        stack.append(right)
    return best`,
			Explanation:     "Builds on Daily Temperatures, which found each day's next warmer day with a decreasing stack. Here the stack is increasing and each pop finds two things at once: the next lower bar on the right (the bar that caused the pop) and the previous lower bar on the left (the bar now on top). Those two boundaries give the widest rectangle at the popped bar's full height.",
			BruteForce:      "Try every pair of edges (i, j), keeping a running minimum of heights[i..j] as j moves right, and take min * (j - i + 1). That's O(n^2): 5 billion steps for 100,000 bars. The waste: for each bar, we keep extending rectangles we could have closed off the moment a lower bar appeared, and we rediscover the same boundaries from every starting point.",
			KeyInsight:      "The best rectangle has the height of some bar i, and for that height it spans from just after the previous lower bar to just before the next lower bar. A monotonic increasing stack produces both of those boundaries for every bar in one pass.",
			WhyItWorks:      "Any optimal rectangle is as tall as its shortest bar i, and it can be widened until bars lower than heights[i] block it on both sides, so it's enough to compute, for each i, heights[i] * (next lower index - previous lower index - 1). Invariant: the stack's heights never decrease from bottom to top. When bar \"right\" pops bar \"tall\", bars[right] < bars[tall] and nothing between them was lower (it would have popped tall), so right is tall's right boundary. The entry below tall is the nearest earlier bar not taller than it, and everything between those two was popped by something taller-or-equal, so it is tall's left boundary. With equal heights, an earlier equal bar may sit just below and make the width too short, but the leftmost bar of that equal run is popped last, with the true full width, so the maximum is still found. The sentinel 0 pops every real bar. Each index is pushed and popped once: O(n).",
			CommonMistakes:  []string{"Forgetting bars still on the stack at the end. [1, 2, 3, 4, 5] never pops until the end, so the answer 9 would be missed. Fix: append a 0-height sentinel bar (or loop over the leftovers with right = n).", "Using width = right - tall instead of right - left - 1. The rectangle can extend left past tall over the taller bars that were popped earlier; [5, 4, 3, 2, 1] needs width 3 for height 3. Fix: the left boundary is the new top after popping, or -1 if the stack is empty.", "Appending the sentinel with heights.append(0) and forgetting to remove it, which changes the caller's list. Fix: build a new list, bars = heights + [0]."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "The stack holds indices of bars whose rectangle could still grow to the right, with heights never decreasing upward. The extra 0-height bar at the end is lower than everything, so it will pop every real bar.", CodeSnippet: "bars = heights + [0]  # a final 0-height bar resolves every bar still waiting\nstack = []  # indices; heights never decrease from bottom to top\nbest = 0", LineStart: 2, LineEnd: 4},
				{Title: "Pop what the new item resolves", Explanation: "A lower bar ends the rectangles of all taller bars on top. In example 1, bar 4 (height 2) pops bar 3 (height 6) and then bar 2 (height 5).", CodeSnippet: "while stack and bars[stack[-1]] > h:\n    tall = stack.pop()", LineStart: 6, LineEnd: 7},
				{Title: "Record the answer at pop time", Explanation: "For bar 2 (height 5): after popping it the top is bar 1 (height 1), so left = 1, right = 4, width = 4 - 1 - 1 = 2, area 10. For bar 3 just before: left = 2, width 1, area 6.", CodeSnippet: "left = stack[-1] if stack else -1  # nearest lower bar on the left\nwidth = right - left - 1\nbest = max(best, bars[tall] * width)", LineStart: 8, LineEnd: 10},
				{Title: "Push the new item", Explanation: "The current bar's rectangle might still grow right, so it waits. Everything below it is not taller, so the order holds.", CodeSnippet: "stack.append(right)", LineStart: 11, LineEnd: 11},
				{Title: "Handle what is left on the stack", Explanation: "Nothing is left: the sentinel bar at index n popped every real bar, measuring their rectangles all the way to the right edge. So best is final.", CodeSnippet: "return best", LineStart: 12, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: every left edge with a running minimum",
					Code: `def largestRectangleArea(heights):
    best = 0
    n = len(heights)
    for i in range(n):
        lowest = heights[i]
        for j in range(i, n):
            lowest = min(lowest, heights[j])
            best = max(best, lowest * (j - i + 1))
    return best`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Checks every range [i, j] once, using the running minimum so each range is O(1). Correct, but too slow for 10^5 bars.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "When a bar is popped, why is the bar now on top of the stack its left boundary?", Answer: "The bar below it on the stack is the nearest earlier bar that is not taller: any bar between them was taller and was popped earlier (by a bar lower than it). So every bar strictly between the new top and the popped bar is at least as tall as the popped bar, and the rectangle can cover them."},
			{Prompt: "What if the histogram became a 2D grid of 0s and 1s and you wanted the largest rectangle of 1s? How could you reuse this solution?", Answer: "Process the grid row by row, keeping for each column the height of consecutive 1s ending at the current row (reset to 0 on a 0). Each row's heights form a histogram, so run this function on it and take the best over all rows. That's O(rows * cols) (LeetCode 85, Maximal Rectangle)."},
		},
	},
	{
		ID:              "generate-parentheses",
		Number:          22,
		Order:           7,
		Title:           "Generate Parentheses",
		Slug:            "generate-parentheses",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"String", "Backtracking", "Stack"},
		RelatedChapters: []int{5, 12, 8},
		Description: `Given an integer n, return every string of n '(' characters and n ')' characters that is a balanced parenthesis string, each exactly once. A string is balanced when, reading left to right, no prefix has more ')' than '(' and the totals are equal (the single-bracket-type case of Valid Parentheses).

The strings may be returned in any order.`,
		Constraints: []string{
			"1 <= n <= 8",
		},
		Examples: []Example{
			{ID: "1", Input: "n = 3", Output: `["((()))","(()())","(())()","()(())","()()()"]`, Explanation: "These are the five ways to arrange 3 pairs. For example \"())(()\" is not allowed: after \"())\" a ')' has nothing to close."},
			{ID: "2", Input: "n = 1", Output: `["()"]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 3}, Expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, Note: "example 1"},
			{Input: map[string]interface{}{"n": 1}, Expected: []string{"()"}, Note: "smallest n"},
			{Input: map[string]interface{}{"n": 2}, Expected: []string{"(())", "()()"}, Hidden: true, Note: "n = 2: nested and side by side"},
			{Input: map[string]interface{}{"n": 4}, Expected: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}, Hidden: true, Note: "n = 4: 14 strings, no duplicates and none missing"},
			{Input: map[string]interface{}{"n": 5}, Expected: []string{
				"((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())",
				"((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))",
				"(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))",
				"(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()",
				"()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()",
				"()()()(())", "()()()()()",
			}, Hidden: true, Note: "n = 5: 42 strings"},
			{Input: map[string]interface{}{"n": 6}, Expected: []string{
				"(((((())))))", "((((()()))))", "((((())())))", "((((()))()))", "((((())))())", "((((()))))()",
				"(((()(()))))", "(((()()())))", "(((()())()))", "(((()()))())", "(((()())))()", "(((())(())))",
				"(((())()()))", "(((())())())", "(((())()))()", "(((()))(()))", "(((()))()())", "(((()))())()",
				"(((())))(())", "(((())))()()", "((()((()))))", "((()(()())))", "((()(())()))", "((()(()))())",
				"((()(())))()", "((()()(())))", "((()()()()))", "((()()())())", "((()()()))()", "((()())(()))",
				"((()())()())", "((()())())()", "((()()))(())", "((()()))()()", "((())((())))", "((())(()()))",
				"((())(())())", "((())(()))()", "((())()(()))", "((())()()())", "((())()())()", "((())())(())",
				"((())())()()", "((()))((()))", "((()))(()())", "((()))(())()", "((()))()(())", "((()))()()()",
				"(()(((()))))", "(()((()())))", "(()((())()))", "(()((()))())", "(()((())))()", "(()(()(())))",
				"(()(()()()))", "(()(()())())", "(()(()()))()", "(()(())(()))", "(()(())()())", "(()(())())()",
				"(()(()))(())", "(()(()))()()", "(()()((())))", "(()()(()()))", "(()()(())())", "(()()(()))()",
				"(()()()(()))", "(()()()()())", "(()()()())()", "(()()())(())", "(()()())()()", "(()())((()))",
				"(()())(()())", "(()())(())()", "(()())()(())", "(()())()()()", "(())(((())))", "(())((()()))",
				"(())((())())", "(())((()))()", "(())(()(()))", "(())(()()())", "(())(()())()", "(())(())(())",
				"(())(())()()", "(())()((()))", "(())()(()())", "(())()(())()", "(())()()(())", "(())()()()()",
				"()((((()))))", "()(((()())))", "()(((())()))", "()(((()))())", "()(((())))()", "()((()(())))",
				"()((()()()))", "()((()())())", "()((()()))()", "()((())(()))", "()((())()())", "()((())())()",
				"()((()))(())", "()((()))()()", "()(()((())))", "()(()(()()))", "()(()(())())", "()(()(()))()",
				"()(()()(()))", "()(()()()())", "()(()()())()", "()(()())(())", "()(()())()()", "()(())((()))",
				"()(())(()())", "()(())(())()", "()(())()(())", "()(())()()()", "()()(((())))", "()()((()()))",
				"()()((())())", "()()((()))()", "()()(()(()))", "()()(()()())", "()()(()())()", "()()(())(())",
				"()()(())()()", "()()()((()))", "()()()(()())", "()()()(())()", "()()()()(())", "()()()()()()",
			}, Hidden: true, Note: "n = 6: 132 strings"},
		},
		TimeComplexity:  "O(4^n / sqrt(n))",
		SpaceComplexity: "O(n)",
		StarterCode:     "def generateParenthesis(n):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In Valid Parentheses with only '(' and ')', every stack entry is the same character. So what is the only thing about the stack you need to know while building a string one character at a time?"},
			{Level: 2, Type: "insight", Content: "The stack of open parens is fully described by its height: opened - closed. Build strings left to right: you may push another '(' while you've used fewer than n, and you may pop with ')' only while the height is above 0 (closed < opened). Exploring both choices at every position is backtracking (Chapter 12); the stack rule (Chapter 5) is what prunes every invalid string before it is finished.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Decide what the stack holds: just two counters, opened and closed; the stack height is opened - closed.\n2. Push the new item: if opened < n, try adding '(' and recurse.\n3. Pop what the new item resolves: if closed < opened, try adding ')' and recurse.\n4. Handle what is left on the stack: when the string has length 2n, opened = closed = n, so the stack is empty and the string is balanced; record it."},
			{Level: 4, Type: "skeleton", Content: `def generateParenthesis(n):
    results = []

    def build(prefix, opened, closed):
        if len(prefix) == 2 * n:
            results.append(prefix)
            return
        if ____:
            build(prefix + '(', opened + 1, closed)
        if ____:
            build(prefix + ')', opened, closed + 1)

    build('', 0, 0)
    return results`},
		},
		Solution: Solution{
			Code: `def generateParenthesis(n):
    results = []

    def build(prefix, opened, closed):
        # opened - closed is the height of the open-paren stack right now
        if len(prefix) == 2 * n:
            results.append(prefix)
            return
        if opened < n:  # push another '('
            build(prefix + '(', opened + 1, closed)
        if closed < opened:  # pop: only legal while the stack is non-empty
            build(prefix + ')', opened, closed + 1)

    build('', 0, 0)
    return results`,
			Explanation:     "This problem sits last in the stack ladder because its search technique is backtracking (Chapter 12): it tries both characters at each position and undoes the choice on return. The stack part is the validity rule from Valid Parentheses, problem 1. With only one bracket type the stack collapses to a counter, opened - closed, and \"pop only from a non-empty stack\" becomes \"add ')' only while closed < opened\". Enforcing that rule while building means only valid strings are ever completed.",
			BruteForce:      "Generate all 2^(2n) strings of '(' and ')' and keep the balanced ones, checking each with the Valid Parentheses counter. For n = 8 that's 65,536 strings of length 16 to check, only 1,430 of which are valid. The waste: strings like \")...\" are built in full even though their first character already makes them invalid.",
			KeyInsight:      "Check the stack rule after every character instead of at the end. A prefix is extendable to a valid string exactly when closed <= opened <= n, so pruning on those two conditions explores only prefixes of valid answers.",
			WhyItWorks:      "Invariant: every prefix passed to build satisfies closed <= opened <= n, i.e. the open-paren stack never underflows and never needs more than n pushes. Any such prefix can be completed to a valid string (add the remaining '(' then all ')'), so no branch is a dead end. When the length reaches 2n, opened + closed = 2n with both <= n forces opened = closed = n: the stack is empty and the string is balanced. Conversely every balanced string satisfies the invariant at each prefix, so the recursion reaches it, and it reaches it once because each string corresponds to one sequence of choices. The number of results is the n-th Catalan number, about 4^n / (n^1.5), and each costs O(n) to build.",
			CommonMistakes:  []string{"Allowing ')' whenever closed < n. That builds \"())(()\": the stack goes negative. Fix: allow ')' only while closed < opened.", "Using one shared list for the prefix and appending without removing after the call, so later strings contain leftovers. Fix: pass prefix + '(' (a new string) as here, or append before and pop after each recursive call.", "Adding the base-case check after the branches, or forgetting to return, which can record the same string twice or keep recursing. Fix: check length first and return."},
			TimeComplexity:  "O(4^n / sqrt(n))",
			SpaceComplexity: "O(n) recursion depth, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "With a single bracket type, the stack of open parens is just its height, opened - closed, so two counters replace it.", CodeSnippet: "def build(prefix, opened, closed):\n    # opened - closed is the height of the open-paren stack right now", LineStart: 4, LineEnd: 5},
				{Title: "Push the new item", Explanation: "Adding '(' pushes onto the stack. It's allowed while fewer than n openers are used.", CodeSnippet: "if opened < n:  # push another '('\n    build(prefix + '(', opened + 1, closed)", LineStart: 9, LineEnd: 10},
				{Title: "Pop what the new item resolves", Explanation: "Adding ')' closes the newest open paren, which is only legal while the stack is non-empty: closed < opened. From \"((\" with n = 2 we can't add '(' but can add ')' twice.", CodeSnippet: "if closed < opened:  # pop: only legal while the stack is non-empty\n    build(prefix + ')', opened, closed + 1)", LineStart: 11, LineEnd: 12},
				{Title: "Handle what is left on the stack", Explanation: "At length 2n both counters equal n, so the stack is empty and the string is valid: record it and stop this branch.", CodeSnippet: "if len(prefix) == 2 * n:\n    results.append(prefix)\n    return", LineStart: 6, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: all strings, filtered by the balance counter",
					Code: `from itertools import product

def generateParenthesis(n):
    results = []
    for chars in product('()', repeat=2 * n):
        balance = 0
        for ch in chars:
            balance += 1 if ch == '(' else -1
            if balance < 0:
                break
        else:
            if balance == 0:
                results.append(''.join(chars))
    return results`,
					TimeComplexity:  "O(2^(2n) * n)",
					SpaceComplexity: "O(n) plus the output",
					Explanation:     "Builds every string and checks it with the Valid Parentheses counter (the stack height). Simple and correct, but most of the 4^n strings are invalid and get thrown away.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why do we never need to check, at the end, that the generated string is balanced?", Answer: "The two guards keep closed <= opened <= n at every step, so the stack never underflows. At length 2n, opened + closed = 2n and both are at most n, so both equal n and the stack is empty. Those are exactly the conditions for a balanced string."},
			{Prompt: "What if you needed strings with n pairs of '()' and m pairs of '[]', properly nested? Would two counters still be enough?", Answer: "No. With two bracket types, the order of the open brackets matters: after \"([\" the next closer must be ']'. You'd keep a real stack of open brackets (push on open, pop on a matching close) and backtrack by undoing the push or pop after each recursive call."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, StackProblems...)
}
