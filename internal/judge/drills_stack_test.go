package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestStackDrillsRejectMistakes checks that each stack foundations drill fails
// the mistake its tests and notes promise to catch, so a drill can't be
// passed without doing the thing it drills.
func TestStackDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-build-stack", "crashes on pop and peek when empty", `class MyStack:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)

    def pop(self):
        return self.items.pop()

    def peek(self):
        return self.items[-1]

    def isEmpty(self):
        return not self.items

    def size(self):
        return len(self.items)`},
		{"drill-build-stack", "peek removes the top", `class MyStack:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)

    def pop(self):
        return self.items.pop() if self.items else -1

    def peek(self):
        return self.items.pop() if self.items else -1

    def isEmpty(self):
        return not self.items

    def size(self):
        return len(self.items)`},
		{"drill-build-stack", "takes from the bottom (a queue)", `class MyStack:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)

    def pop(self):
        return self.items.pop(0) if self.items else -1

    def peek(self):
        return self.items[0] if self.items else -1

    def isEmpty(self):
        return not self.items

    def size(self):
        return len(self.items)`},
		{"drill-build-stack", "size goes negative when popping an empty stack", `class MyStack:
    def __init__(self):
        self.items = []
        self.count = 0

    def push(self, val):
        self.items.append(val)
        self.count += 1

    def pop(self):
        self.count -= 1
        return self.items.pop() if self.items else -1

    def peek(self):
        return self.items[-1] if self.items else -1

    def isEmpty(self):
        return not self.items

    def size(self):
        return self.count`},
		{"drill-build-stack", "treats a stored 0 as empty", `class MyStack:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)

    def pop(self):
        return self.items.pop() if self.items else -1

    def peek(self):
        return self.items[-1] if self.items and self.items[-1] else -1

    def isEmpty(self):
        return not self.items or not self.items[-1]

    def size(self):
        return len(self.items)`},
		{"drill-reverse-each-word", "never flushes the last word", `def reverseWords(s):
    out = []
    stack = []
    for ch in s:
        if ch != " ":
            stack.append(ch)
        else:
            while stack:
                out.append(stack.pop())
            out.append(" ")
    return "".join(out)`},
		{"drill-reverse-each-word", "reverses the whole string", `def reverseWords(s):
    stack = list(s)
    out = []
    while stack:
        out.append(stack.pop())
    return "".join(out)`},
		{"drill-remove-adjacent-duplicates", "one replace pass, no cascade", `def removeDuplicates(s):
    for c in "abcdefghijklmnopqrstuvwxyz":
        s = s.replace(c + c, "")
    return s`},
		{"drill-remove-adjacent-duplicates", "removes the whole run of equal letters", `def removeDuplicates(s):
    out = []
    i = 0
    while i < len(s):
        j = i
        while j < len(s) and s[j] == s[i]:
            j += 1
        if j - i == 1:
            out.append(s[i])
        i = j
    return "".join(out)`},
		{"drill-queue-using-stacks", "moves the inbox on every pop, even when the outbox has items", `class MyQueue:
    def __init__(self):
        self.inbox = []
        self.outbox = []

    def push(self, x):
        self.inbox.append(x)

    def _refill(self):
        while self.inbox:
            self.outbox.append(self.inbox.pop())

    def pop(self):
        self._refill()
        return self.outbox.pop()

    def peek(self):
        self._refill()
        return self.outbox[-1]

    def empty(self):
        return not self.inbox and not self.outbox`},
		{"drill-queue-using-stacks", "empty only checks the inbox", `class MyQueue:
    def __init__(self):
        self.inbox = []
        self.outbox = []

    def push(self, x):
        self.inbox.append(x)

    def _refill(self):
        if not self.outbox:
            while self.inbox:
                self.outbox.append(self.inbox.pop())

    def pop(self):
        self._refill()
        return self.outbox.pop()

    def peek(self):
        self._refill()
        return self.outbox[-1]

    def empty(self):
        return not self.inbox`},
		{"drill-queue-using-stacks", "pops the newest element (a stack)", `class MyQueue:
    def __init__(self):
        self.items = []

    def push(self, x):
        self.items.append(x)

    def pop(self):
        return self.items.pop()

    def peek(self):
        return self.items[-1]

    def empty(self):
        return not self.items`},
		{"drill-next-greater-element", "counts an equal value as greater", `def nextGreater(nums):
    answer = [-1] * len(nums)
    stack = []
    for i, num in enumerate(nums):
        while stack and nums[stack[-1]] <= num:
            answer[stack.pop()] = num
        stack.append(i)
    return answer`},
		{"drill-next-greater-element", "pops only one waiting element", `def nextGreater(nums):
    answer = [-1] * len(nums)
    stack = []
    for i, num in enumerate(nums):
        if stack and nums[stack[-1]] < num:
            answer[stack.pop()] = num
        stack.append(i)
    return answer`},
		{"drill-next-greater-element", "starts the answer at 0 instead of -1", `def nextGreater(nums):
    answer = [0] * len(nums)
    stack = []
    for i, num in enumerate(nums):
        while stack and nums[stack[-1]] < num:
            answer[stack.pop()] = num
        stack.append(i)
    return answer`},
		{"drill-previous-smaller-element", "keeps an equal value as the answer", `def prevSmaller(nums):
    answer = [-1] * len(nums)
    stack = []
    for i, num in enumerate(nums):
        while stack and nums[stack[-1]] > num:
            stack.pop()
        if stack:
            answer[i] = stack[-1]
        stack.append(i)
    return answer`},
		{"drill-previous-smaller-element", "pops only one larger candidate", `def prevSmaller(nums):
    answer = [-1] * len(nums)
    stack = []
    for i, num in enumerate(nums):
        if stack and nums[stack[-1]] >= num:
            stack.pop()
        if stack:
            answer[i] = stack[-1]
        stack.append(i)
    return answer`},
		{"drill-previous-smaller-element", "returns the value instead of the index", `def prevSmaller(nums):
    answer = [-1] * len(nums)
    stack = []
    for i, num in enumerate(nums):
        while stack and nums[stack[-1]] >= num:
            stack.pop()
        if stack:
            answer[i] = nums[stack[-1]]
        stack.append(i)
    return answer`},
		{"drill-evaluate-expression", "drops the sign in front of a group", `def calculate(s):
    total = 0
    sign = 1
    number = 0
    stack = []
    for ch in s:
        if ch.isdigit():
            number = number * 10 + int(ch)
        elif ch == "+" or ch == "-":
            total += sign * number
            number = 0
            sign = 1 if ch == "+" else -1
        elif ch == "(":
            stack.append(total)
            total = 0
            sign = 1
        elif ch == ")":
            total += sign * number
            number = 0
            total = stack.pop() + total
    return total + sign * number`},
		{"drill-evaluate-expression", "reads each digit as its own number", `def calculate(s):
    total = 0
    sign = 1
    stack = []
    for ch in s:
        if ch.isdigit():
            total += sign * int(ch)
        elif ch == "+" or ch == "-":
            sign = 1 if ch == "+" else -1
        elif ch == "(":
            stack.append((total, sign))
            total = 0
            sign = 1
        elif ch == ")":
            saved_total, saved_sign = stack.pop()
            total = saved_total + saved_sign * total
    return total`},
		{"drill-evaluate-expression", "forgets the last number", `def calculate(s):
    total = 0
    sign = 1
    number = 0
    stack = []
    for ch in s:
        if ch.isdigit():
            number = number * 10 + int(ch)
        elif ch == "+" or ch == "-":
            total += sign * number
            number = 0
            sign = 1 if ch == "+" else -1
        elif ch == "(":
            stack.append((total, sign))
            total = 0
            sign = 1
        elif ch == ")":
            total += sign * number
            number = 0
            saved_total, saved_sign = stack.pop()
            total = saved_total + saved_sign * total
    return total`},
	}
	for _, c := range cases {
		c := c
		t.Run(c.id+"/"+c.why, func(t *testing.T) {
			t.Parallel()
			var p *practice.Problem
			for _, q := range practice.EmbeddedProblems {
				if q.ID == c.id {
					p = q
				}
			}
			if p == nil {
				t.Fatalf("no problem %s", c.id)
			}
			for i := range p.TestCases {
				res, err := Run(p, c.code, &p.TestCases[i], Options{})
				if err != nil {
					t.Fatalf("case %d: %v", i+1, err)
				}
				if !res.Passed {
					return
				}
			}
			t.Errorf("code that %s passed every case", c.why)
		})
	}
}
