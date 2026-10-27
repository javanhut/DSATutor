package chapter

import "time"

// chapter05 teaches stacks and monotonic stacks; see docs/pedagogy.md.
// Storyboards (IDs registered in app.js): min-stack, car-fleet.
//
// Subgoal labels reused across the worked examples, the completions and the
// stack problems' plan hints and walkthroughs:
//
//	Decide what the stack holds / Pop what the new item resolves /
//	Record the answer at pop time / Push the new item /
//	Handle what is left on the stack
func chapter05() Chapter {
	return Chapter{
		Number:        5,
		Slug:          "stack",
		Title:         "Stacks and Monotonic Stacks",
		Category:      "stack",
		Prerequisites: []int{2},
		Summary:       "A stack hands back the most recent unfinished item first. That fits nesting (brackets, expressions), undo, and, kept in sorted order as a monotonic stack, every \"next greater / next smaller element\" question in O(n).",
		Motivation: `Weather data for 100,000 days: for each day, how many days until it gets warmer? The obvious way scans forward from every day until a warmer one shows up. If the temperatures mostly fall, every scan runs to the end of the list: about 100,000 x 100,000 / 2 = 5 billion comparisons, which takes minutes in Python.

Look at what the scans repeat. While the weather keeps getting colder, nothing is answered: day 1 is waiting, day 2 is waiting, day 3 is waiting. Then a warm day arrives and answers a whole group of waiting days at once, starting with the most recent one. "Most recent first" is exactly how a stack works. Keep the waiting days on a stack, and when a warm day arrives, pop every colder day off the top and record its answer. Each day is pushed once and popped at most once, so the 5 billion comparisons become about 200,000 stack operations.

This chapter starts with the plain stack (matching brackets, evaluating expressions, a stack that remembers its minimum) and then turns it into a monotonic stack, the tool behind Daily Temperatures, Car Fleet and Largest Rectangle in Histogram.`,
		Objectives: []string{
			"Trace a stack through push and pop operations and write down its contents after each step.",
			"Explain why a closing bracket must be matched against the top of the stack, and why counting brackets isn't enough.",
			"Define a monotonic stack, and explain why each pop is the moment that element's next greater (or smaller) element is found.",
			"Explain why a monotonic stack runs in O(n) even though it has a while loop inside a for loop.",
			"Implement next greater element, Daily Temperatures and a min stack using the five subgoals: decide what the stack holds, pop what the new item resolves, record the answer at pop time, push the new item, handle what is left.",
			"Tell stack problems apart from two-pointer (Chapter 3) and sliding-window (Chapter 4) problems.",
		},
		WhenToUse: []string{
			"Things must be closed or undone in the reverse order they were opened: brackets, HTML tags, nested calls, undo, backspace.",
			"The statement asks for the next / previous greater or smaller element, or \"how many days until\" something larger appears.",
			"An operator or event acts on the most recent unfinished items, as in postfix (Reverse Polish) expressions.",
			"You need the widest range around each element bounded by the nearest lower (or higher) element on each side, such as rectangles in a histogram.",
			"A container must answer \"what is the minimum / maximum right now?\" in O(1) while items are pushed and popped at one end.",
		},
		Concepts: []Concept{
			{
				Name:        "The stack: last in, first out",
				Description: "A stack is a list where you only add (push) and remove (pop) at one end, the top. In Python a list is a stack: append pushes, pop() pops, and stack[-1] peeks at the top, all O(1).",
				WhyItWorks:  `A stack fits any problem where the item that must be finished next is always the most recently started unfinished one. Nested brackets are the classic case: in "{[()]}", the '(' opened last must close first, otherwise two pairs would cross, as in "([)]". So the invariant is: the stack holds exactly the unfinished items, oldest at the bottom and newest on top. When a new item arrives it either finishes the top item (pop and check) or starts something new (push). Both steps keep the invariant, and each costs O(1), so one left-to-right scan does the whole job in O(n).`,
				Intuition:   `A stack of plates: you put plates on top and take them off the top. The plate you can reach is the one you put down last. Mapping: a plate is an unfinished item (an open bracket, a waiting number); putting it down is push; the only plate you can take is the newest unfinished item, which is exactly the one a closing bracket must match.`,
				CoreIdeas: []string{
					"Push when an item can't be finished yet; pop when the arriving item finishes the top one.",
					"Peek with stack[-1] only after checking the stack isn't empty.",
					"At the end, whatever is left on the stack was never finished; decide what that means (invalid brackets, answer 0, a final count).",
					"Decide what each entry holds: the item itself, its index (when you need distances), or a pair with extra state (Min Stack stores (value, minimum so far)).",
					"A stack entry can carry a snapshot of a summary of everything below it, because nothing below the top changes while the entry is on the stack.",
				},
				CommonMistakes: []string{
					"Calling stack.pop() or stack[-1] on an empty stack. Fix: check \"if stack\" first; an unexpected closer like ']' with an empty stack means invalid input.",
					"Forgetting the leftovers after the loop, e.g. returning True for \"((\". Fix: make \"handle what is left on the stack\" an explicit last step.",
					"Popping operands in the wrong order: in postfix \"3 4 -\" the first pop is the right operand (4). Fix: the later token sits higher on the stack, so it comes off first.",
				},
				Examples: []CodeExample{
					{
						ID:       "stack-balanced",
						Title:    "Balanced brackets with a stack",
						Language: "python",
						Snippet: `def is_balanced(s):
    closer_to_opener = {')': '(', ']': '[', '}': '{'}
    stack = []
    for ch in s:
        if ch in closer_to_opener:
            if not stack or stack[-1] != closer_to_opener[ch]:
                return False
            stack.pop()
        else:
            stack.append(ch)
    return not stack`,
						Notes: "\"{[]}\": push '{', push '[', ']' pops '[', '}' pops '{', stack empty: True. \"([)]\": ')' finds '[' on top: False.",
					},
				},
			},
			{
				Name:        "Monotonic stacks: next greater and next smaller",
				Description: "A monotonic stack is a stack whose values stay in sorted order from bottom to top. A decreasing stack (each entry smaller than the one below it) finds each element's next greater element; an increasing stack finds each element's next smaller element.",
				WhyItWorks:  `Definition: before pushing a new element x onto a decreasing stack, pop every entry that is smaller than x, because leaving it would break the order. Each pop is the moment that entry's next greater element is found, and it is x. Why x is the FIRST greater element to the right of the popped entry e: e is still on the stack when x arrives, so no element between e and x was greater than e, since that element would already have popped e. Why it is O(n): every index is pushed exactly once and popped at most once, so all the while loops together do at most n pops, about 2n stack operations in total, even though the code has a loop inside a loop. Entries left on the stack at the end never met a greater element. An increasing stack is the mirror image: pop while the top is larger than x, and x is the popped entry's next smaller element. After the pops, the entry left on top is the new element's previous greater (or smaller) element, so one pass can give both neighbours. The same idea with a deque, where old entries also leave from the bottom when they fall out of a window, is the monotonic deque of Sliding Window Maximum in Chapter 4.`,
				Intuition:   `People waiting in a queue at a viewpoint, each blocking the view of shorter people behind them. Mapping: each person is an element; the stack is the line of people who still can't see anyone taller ahead of them, tallest at the back (bottom) and shortest at the front (top). When a tall person walks up, every shorter person at the front now sees someone taller, the newcomer, so they step out of line (pop) and record who it was. The newcomer stops at the first person at least as tall, and joins the line.`,
				CoreIdeas: []string{
					"Decreasing stack (values shrink from bottom to top) answers \"next greater\"; increasing stack answers \"next smaller\".",
					"Push indices, not values, whenever the answer is a distance or a width; look values up with nums[stack[-1]].",
					"Record the answer for an element at the moment it is popped; the element that caused the pop is its answer.",
					"Strict or non-strict comparison decides how ties behave: Daily Temperatures pops on < because an equal day is not warmer.",
					"Each element is pushed once and popped at most once, so the whole scan is O(n).",
					"A monotonic deque (Chapter 4, Sliding Window Maximum) is the same structure plus removal of expired entries from the bottom.",
				},
				CommonMistakes: []string{
					"Using the wrong direction: a decreasing stack can't find next smaller elements. Fix: ask \"which arriving element resolves a waiting one?\"; if a larger one does, keep the stack decreasing.",
					"Getting ties wrong, e.g. popping on <= in Daily Temperatures so 70 resolves an earlier 70. Fix: decide from the statement whether equal counts, and test with equal values.",
					"Believing the nested while makes it O(n^2) and reaching for something else. Fix: count total pushes and pops, not loop nesting.",
				},
				Examples: []CodeExample{
					{
						ID:       "next-greater-element",
						Title:    "Next greater element with a decreasing stack",
						Language: "python",
						Snippet: `def next_greater(nums):
    answer = [-1] * len(nums)
    stack = []  # indices; values decrease from bottom to top
    for i, x in enumerate(nums):
        while stack and nums[stack[-1]] < x:
            answer[stack.pop()] = x
        stack.append(i)
    return answer`,
						Notes: "next_greater([2, 7, 3, 5, 4, 6, 8]) == [7, 8, 5, 6, 6, 8, -1]. When 6 arrives it pops 4 and then 5; it stops at 7.",
					},
				},
			},
			{
				Name:        "Store the answer with each entry",
				Description: "When a stack must answer a summary question (minimum, maximum, sum) in O(1), store the summary of everything at or below each entry inside the entry itself, as a pair.",
				WhyItWorks:  `A stack only changes at its top. So for any entry, the set of elements at or below it stays the same from the moment it is pushed until the moment it is popped. A summary of that set, such as its minimum, can therefore be computed once at push time, as min(new value, summary stored in the entry below), and it never goes stale. Popping needs no repair work: the entry below already holds the correct summary for the smaller stack. That gives O(1) getMin with O(n) extra space, where recomputing min over the whole list would cost O(n) per query.`,
				Intuition:   `A pile of paperwork where each sheet has a sticky note saying "the most urgent deadline in this pile from here down". Mapping: the sheet is the pushed value, the sticky note is the stored minimum, and the pile is the stack. To know the most urgent deadline you read the top note. When you remove the top sheet, the note on the next sheet is already correct, because nothing under it changed.`,
				CoreIdeas: []string{
					"Push (value, summary) where summary = combine(value, summary of the entry below).",
					"Pop does nothing special: the entry below already carries the right summary.",
					"Any summary that can be updated from (new value, old summary) works: min, max, sum, count.",
					"An alternative stores summaries on a second stack only when they change, but then ties must be pushed too (use <=, not <).",
				},
				CommonMistakes: []string{
					"Keeping one global minimum variable, which can't be restored after the minimum is popped. Fix: one snapshot per entry.",
					"With a separate stack of minimums, pushing only strictly smaller values: push 2, push 2, pop leaves the minimum stack empty while a 2 remains. Fix: push when val <= current minimum.",
				},
				Examples: []CodeExample{
					{
						ID:       "min-stack-pairs",
						Title:    "Min stack with (value, minimum) pairs",
						Language: "python",
						Snippet: `class MinStack:
    def __init__(self):
        self.stack = []  # (value, smallest at or below)

    def push(self, val):
        smallest = min(val, self.stack[-1][1]) if self.stack else val
        self.stack.append((val, smallest))

    def pop(self):
        self.stack.pop()

    def getMin(self):
        return self.stack[-1][1]`,
						Notes: "push 3, push 5, push 2 stores (3,3), (5,3), (2,2). getMin is 2; after pop it is 3 again.",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Remove adjacent duplicates (a plain stack)",
				Problem: "Given a lowercase string s, repeatedly delete any two adjacent equal letters until no such pair remains, and return the result. For s = \"abbaca\", deleting \"bb\" gives \"aaca\", deleting \"aa\" gives \"ca\". Do it in one pass.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Decide what the stack holds",
						Explanation: "The stack holds the letters kept so far, in order. Only the last kept letter can be touched by the next one, so the last kept letter must be the easiest to reach: the top of a stack.",
						Code:        "stack = []  # letters kept so far",
						State:       "s = \"abbaca\", stack = []",
					},
					{
						Subgoal:     "Pop what the new item resolves",
						Explanation: "If the new letter equals the top, the two are adjacent and equal, so both disappear: pop the top and don't push. Removing them may make the letter below touch the next one, which the stack handles automatically.",
						Code:        "if stack and stack[-1] == ch:\n    stack.pop()",
						State:       "after 'a','b': [a, b]; 'b' == top: pop -> [a]; next 'a' == top: pop -> []",
					},
					{
						Subgoal:     "Push the new item",
						Explanation: "Otherwise the new letter is kept and becomes the letter the next one must be compared with.",
						Code:        "else:\n    stack.append(ch)",
						State:       "'c': push -> [c]; 'a': push -> [c, a]",
					},
					{
						Subgoal:     "Handle what is left on the stack",
						Explanation: "The letters still on the stack survived every deletion, in their original order (bottom to top), so they are the answer.",
						Code:        "return ''.join(stack)",
						State:       "stack = [c, a] -> \"ca\"",
					},
				},
				Code: `def remove_duplicates(s):
    stack = []  # letters kept so far
    for ch in s:
        if stack and stack[-1] == ch:
            stack.pop()  # ch cancels the letter it touches
        else:
            stack.append(ch)
    return ''.join(stack)`,
			},
			{
				Title:   "Next greater element (a monotonic stack)",
				Problem: "For each element of nums, find the first element to its right that is strictly greater, or -1 if there is none. For nums = [2, 7, 3, 5, 4, 6, 8] the answer is [7, 8, 5, 6, 6, 8, -1].",
				Steps: []WorkedStep{
					{
						Subgoal:     "Decide what the stack holds",
						Explanation: "The stack holds indices of elements still waiting for a greater element. Their values decrease from bottom to top: a larger value would already have popped the smaller ones. Start every answer at -1, which is right for anything never resolved.",
						Code:        "answer = [-1] * len(nums)\nstack = []  # indices; values decrease bottom to top",
						State:       "i=0 (2): stack empty, push -> [0] (values [2])",
					},
					{
						Subgoal:     "Pop what the new item resolves",
						Explanation: "A new value x resolves every waiting element smaller than it, and those are all on top because the stack is decreasing. Stop at the first value that is not smaller.",
						Code:        "while stack and nums[stack[-1]] < x:",
						State:       "i=1 (7): 2 < 7, pop index 0. i=2 (3): 7 < 3? no. i=3 (5): 3 < 5, pop index 2; 7 < 5? no",
					},
					{
						Subgoal:     "Record the answer at pop time",
						Explanation: "The popped element is still on the stack when x arrives, so nothing between them was greater than it: x is its FIRST greater element.",
						Code:        "    answer[stack.pop()] = x",
						State:       "answer[0] = 7, answer[2] = 5; i=5 (6) pops index 4 (4) then index 3 (5): answer[4] = 6, answer[3] = 6",
					},
					{
						Subgoal:     "Push the new item",
						Explanation: "x now waits for its own greater element. Everything below it is at least as large, so the stack stays decreasing.",
						Code:        "stack.append(i)",
						State:       "after i=5: stack [1, 5] (values [7, 6]); i=6 (8) pops 5 and 1, then stack [6]",
					},
					{
						Subgoal:     "Handle what is left on the stack",
						Explanation: "Index 6 (value 8) never met a greater value; its answer stays -1 from the initial list. Every index was pushed once and popped at most once: 7 pushes and 6 pops in total, O(n).",
						Code:        "return answer",
						State:       "answer = [7, 8, 5, 6, 6, 8, -1]",
					},
				},
				Code: `def next_greater(nums):
    answer = [-1] * len(nums)
    stack = []  # indices; values decrease from bottom to top
    for i, x in enumerate(nums):
        while stack and nums[stack[-1]] < x:
            answer[stack.pop()] = x
        stack.append(i)
    return answer`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Apply backspaces",
				Prompt: "In the string s, '#' means backspace: it deletes the previous character still on screen, if any. Return the final text. For \"ab#c\" the answer is \"ac\"; for \"a##b\" it is \"b\". Fill in the blanks using the same subgoals as Remove adjacent duplicates.",
				Code: `def apply_backspaces(s):
    stack = []  # characters on screen
    for ch in s:
        if ch == '#':
            if ____:
                ____
        else:
            ____
    return ''.join(stack)`,
				Blanks:      []string{"stack", "stack.pop()", "stack.append(ch)"},
				Explanation: "A backspace resolves the most recent character on screen, the top, but only if there is one: check \"stack\" before popping (\"a##b\" has a second '#' with nothing to delete). Any other character is pushed. What's left on the stack is the text on screen.",
			},
			{
				Title:  "Next smaller element",
				Prompt: "For each element, find the first element to its right that is strictly smaller, or -1. For [4, 8, 5, 2, 6] the answer is [2, 5, 2, -1, -1]. This mirrors the next-greater worked example with an increasing stack. Fill in the blanks.",
				Code: `def next_smaller(nums):
    answer = [-1] * len(nums)
    stack = []  # indices; values increase from bottom to top
    for i, x in enumerate(nums):
        while stack and ____:
            answer[____] = x
        stack.append(i)
    return answer`,
				Blanks:      []string{"nums[stack[-1]] > x", "stack.pop()"},
				Explanation: "Now a smaller arriving value resolves the waiting ones, so pop while the top's value is greater than x, and record x as the popped index's answer at pop time. The stack stays increasing because everything larger than x leaves before x is pushed.",
			},
		},
		Animations: []Storyboard{
			{
				ID:    "min-stack",
				Title: "Min Stack Operations",
				Goal:  "See how storing (value, minimum so far) pairs makes getMin O(1) and lets pop restore the previous minimum for free.",
				MemoryTips: []string{
					"Decide what the stack holds: pairs (value, smallest at or below this entry).",
					"Push the new item: its minimum is min(val, minimum stored in the top pair), or val on an empty stack.",
					"Pop just removes the top pair; the pair below already knows the previous minimum.",
					"top reads the first part of the top pair, getMin the second part: both O(1).",
				},
				Steps: []StoryboardStep{
					{Cue: "push-min", Narration: "Push 3: the stack is empty, so store (3, 3). Push 5: 5 is bigger than the minimum below (3), so store (5, 3). getMin reads the top pair's minimum: 3. Push 2 stores (2, 2) and push 1 stores (1, 1): each pair records the smallest value at or below it.", VisualHint: "Stack nodes carry value + min", Duration: Duration(3 * time.Second), CodeRef: "push"},
					{Cue: "query", Narration: "getMin now reads (1, 1) on top and returns 1 without scanning anything. top would return the value part, 1. Both are a single read of the top pair.", VisualHint: "Read last tuple", Duration: Duration(2 * time.Second), CodeRef: "top/getMin"},
					{Cue: "pop", Narration: "Pop removes (1, 1). The new top is (2, 2), so getMin returns 2: the previous minimum came back without any recomputation. Pop again removes (2, 2); top is now (5, 3), so top returns 5 and getMin would return 3.", VisualHint: "Remove top tuple", Duration: Duration(2 * time.Second), CodeRef: "pop"},
				},
			},
			{
				ID:    "car-fleet",
				Title: "Car Fleet",
				Goal:  "Watch fleet arrival times build a stack that strictly increases from the front fleet to the back for target 12, positions [10, 8, 0, 5, 3], speeds [2, 4, 1, 1, 3].",
				MemoryTips: []string{
					"Decide what the stack holds: one arrival time per fleet. Sort cars closest to the target first.",
					"Compute each car's solo time with real division: (target - position) / speed.",
					"Merge with the fleet on top if the time is <= the top's time (catching up exactly at the target still merges).",
					"Push the new item otherwise: a slower car starts a new fleet.",
					"Handle what is left on the stack: its size is the number of fleets.",
				},
				Steps: []StoryboardStep{
					{Cue: "sort", Narration: "Sort the cars from closest to the target to farthest: positions 10, 8, 5, 3, 0. A car can only be blocked by cars ahead of it, so each car is processed after every car that could block it.", VisualHint: "Positions descending", Duration: Duration(2 * time.Second), CodeRef: "sort"},
					{Cue: "time", Narration: "Solo arrival times: car 10 needs (12-10)/2 = 1.0 hour, car 8 needs 4/4 = 1.0, car 5 needs 7/1 = 7.0, car 3 needs 9/3 = 3.0, car 0 needs 12/1 = 12.0.", VisualHint: "Time to target", Duration: Duration(2 * time.Second), CodeRef: "time"},
					{Cue: "stack", Narration: "Car 10 pushes 1.0. Car 8 (1.0) is not slower than 1.0, so it catches car 10 exactly at the target and joins: fleet [10, 8]. Car 5 (7.0) is slower, so it pushes 7.0. Car 3 (3.0) would arrive before 7.0, so it catches car 5 and joins: fleet [5, 3]. Car 0 (12.0) is slower than everything ahead and pushes 12.0: fleet [0]. The stack [1.0, 7.0, 12.0] holds 3 fleets.", VisualHint: "Monotonic stack of times", Duration: Duration(3 * time.Second), CodeRef: "stack"},
				},
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Which approach fits best? \"For each element of an unsorted array, find the nearest element to its right that is larger than it.\"",
				Choices: []string{
					"Two pointers moving in from both ends (Chapter 3)",
					"A sliding window that grows and shrinks (Chapter 4)",
					"A monotonic decreasing stack",
					"A hash map from value to index (Chapter 2)",
				},
				Answer:      2,
				Explanation: "\"Nearest larger to the right\" is the next-greater-element question, the signature of a monotonic stack. Two pointers need sorted data or a pair condition; a sliding window needs a contiguous range with a validity rule; a hash map finds exact values, not the nearest larger one.",
			},
			{
				Prompt: "Trace this code on nums = [2, 1, 2, 4, 3]. What does it return?\n\ndef next_greater(nums):\n    answer = [-1] * len(nums)\n    stack = []\n    for i, x in enumerate(nums):\n        while stack and nums[stack[-1]] < x:\n            answer[stack.pop()] = x\n        stack.append(i)\n    return answer",
				Choices: []string{
					"[4, 2, 4, -1, -1]",
					"[2, 2, 4, -1, -1]",
					"[4, 4, 4, -1, -1]",
					"[4, 2, 4, 3, -1]",
				},
				Answer:      0,
				Explanation: "i=1 (1) is pushed on top of 2. i=2 (2) pops index 1 (answer 2) but stops at index 0, because 2 < 2 is false; an equal value is not greater. i=3 (4) pops index 2 and index 0 (both get 4). Indices 3 and 4 are never popped, so they stay -1. [2, 2, ...] would come from popping on <=; [4, 4, 4, ...] confuses \"next greater\" with \"largest to the right\".",
			},
			{
				Prompt: "Valid Parentheses uses closer_to_opener = {')': '(', ']': '[', '}': '{'} and tests \"if ch in closer_to_opener\". From Chapter 2, what does that test check, and how long does it take?",
				Choices: []string{
					"Whether ch is one of the dictionary's keys, in O(1) average time",
					"Whether ch is one of the dictionary's values, in O(1) average time",
					"Whether ch is one of the keys, in O(n) time because it scans the dictionary",
					"Whether ch is a key or a value, in O(log n) time",
				},
				Answer:      0,
				Explanation: "\"in\" on a dict checks keys only, using the hash table, so it is O(1) on average. Here the keys are the closers, so the test means \"is ch a closing bracket?\", and closer_to_opener[ch] then gives the opener it must match.",
			},
			{
				Prompt: "A learner implements Min Stack with a values stack and a mins stack. push(val) appends to mins only when val < mins[-1] (or mins is empty); pop() pops mins when the popped value equals mins[-1]. What happens on push(2), push(2), pop(), getMin()?",
				Choices: []string{
					"getMin returns 2",
					"getMin crashes with an IndexError because mins is empty",
					"getMin returns the wrong value, 0",
				},
				Answer:      1,
				Explanation: "The second 2 is not strictly smaller, so mins stays [2]. pop removes a 2, which equals mins[-1], so mins becomes [] while values is still [2]. getMin then reads mins[-1] on an empty list. Pushing when val <= mins[-1] (or storing a minimum with every entry) fixes it.",
			},
			{
				Prompt:      "Daily Temperatures has a while loop inside a for loop. Explain in your own words why it still runs in O(n).",
				Explanation: "Count stack operations instead of loop nesting. Each index is pushed exactly once (at the end of its own iteration) and can be popped at most once, since a popped index never returns. So across the whole run the while loop's body executes at most n times in total, and the for loop does n pushes: at most 2n operations, O(n). A single iteration can pop many entries, but those entries can't be popped again later.",
			},
			{
				Prompt:      "Sliding Window Maximum (Chapter 4) and Daily Temperatures both keep values in decreasing order. Why does Sliding Window Maximum need a deque while Daily Temperatures only needs a stack?",
				Explanation: "In both, a new value removes smaller values from the newest end, since they can never be the answer again; that's the monotonic-stack part. But in Sliding Window Maximum entries also expire when the window's left edge passes them, and the oldest entry (the current maximum) sits at the other end, so you need removal at both ends: a deque. In Daily Temperatures an entry only ever leaves when a warmer day resolves it from the top, so one end is enough.",
			},
			{
				Prompt:      "In Car Fleet, why must the cars be sorted by position (closest to the target first) before using the stack, and what does each stack entry mean?",
				Explanation: "A car can only be slowed down by cars ahead of it, so we must know every fleet ahead before deciding a car's fate; sorting closest-first guarantees that. Each stack entry is the arrival time of one fleet (its leader's solo time). The times strictly increase toward the top, so the top is the fleet directly ahead and the slowest so far: a car with time <= top joins it, a slower car pushes a new fleet, and the stack size is the answer.",
			},
		},
	}
}
