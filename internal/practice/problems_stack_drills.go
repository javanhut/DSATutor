package practice

// StackDrills are the stack foundations: build a stack by hand, use push-then-pop
// to reverse and to cancel, build a queue out of two stacks, then the monotonic
// stack template in both directions and a stack that saves state, before the
// stack ladder applies them. They follow docs/pedagogy.md like any problem,
// minus the brute force (a drill is the plain algorithm itself).
var StackDrills = []*Problem{
	{
		ID:              "drill-build-stack",
		Number:          0,
		Kind:            KindDrill,
		Order:           1,
		Title:           "Build a Stack",
		Slug:            "build-a-stack",
		Difficulty:      "Easy",
		Category:        "stack",
		Tags:            []string{"Stack", "Design", "Foundations"},
		RelatedChapters: []int{5},
		Description: `Build your own stack class, ` + "`" + `MyStack` + "`" + `. A stack is last in, first out (LIFO): the item you pushed most recently is the first one to come back out, like a pile of plates where you only touch the top plate. It supports these methods:

- ` + "`" + `push(val)` + "`" + `: put ` + "`" + `val` + "`" + ` on top of the stack.
- ` + "`" + `pop()` + "`" + `: remove the top item and return it, or return ` + "`" + `-1` + "`" + ` if the stack is empty.
- ` + "`" + `peek()` + "`" + `: return the top item without removing it, or return ` + "`" + `-1` + "`" + ` if the stack is empty.
- ` + "`" + `isEmpty()` + "`" + `: return ` + "`" + `True` + "`" + ` if the stack holds no items, otherwise ` + "`" + `False` + "`" + `.
- ` + "`" + `size()` + "`" + `: return how many items the stack holds.

Every pushed value is 0 or more, so ` + "`" + `-1` + "`" + ` can only mean "nothing there". You may keep the items in a Python list (` + "`" + `append` + "`" + ` and ` + "`" + `pop` + "`" + ` work at its end), or chain your own nodes so that each one points at the item below it. Either way, what matters is the interface: nothing but the top item is ever read or removed.

How the tests call it: ` + "`" + `operations` + "`" + ` lists the method names (the first is the constructor) and ` + "`" + `values` + "`" + ` the arguments for each; the expected list has null for calls that return nothing.`,
		Constraints: []string{
			"0 <= val <= 1000",
			"At most 1000 calls to push, pop, peek, isEmpty and size",
		},
		Examples: []Example{
			{Input: "[\"MyStack\",\"push\",\"push\",\"peek\",\"pop\",\"size\",\"isEmpty\"]\n[[],[1],[2],[],[],[],[]]", Output: "[null,null,null,2,2,1,false]", Explanation: "After push(1), push(2) the stack is 1 then 2 on top. peek() reads 2 and leaves it, pop() removes 2, so one item (the 1) is left: size() is 1 and isEmpty() is false."},
			{Input: "[\"MyStack\",\"pop\",\"peek\",\"isEmpty\",\"size\"]\n[[],[],[],[],[]]", Output: "[null,-1,-1,true,0]", Explanation: "A new stack is empty: pop and peek return -1, isEmpty is true and the size is 0."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"MyStack", "push", "push", "peek", "pop", "size", "isEmpty"}, "values": [][]int{{}, {1}, {2}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, 2, 2, 1, false}, Note: "example 1: peek reads the top, pop removes it"},
			{Input: map[string]interface{}{"operations": []string{"MyStack", "pop", "peek", "isEmpty", "size"}, "values": [][]int{{}, {}, {}, {}, {}}}, Expected: []interface{}{nil, -1, -1, true, 0}, Note: "example 2: a new stack is empty, and pop and peek on it return -1"},
			{Input: map[string]interface{}{"operations": []string{"MyStack", "push", "push", "push", "pop", "pop", "pop", "pop"}, "values": [][]int{{}, {1}, {2}, {3}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, 3, 2, 1, -1}, Hidden: true, Note: "last in, first out: 3, 2, 1, and then -1 once the stack is empty"},
			{Input: map[string]interface{}{"operations": []string{"MyStack", "push", "push", "pop", "push", "pop", "pop", "pop"}, "values": [][]int{{}, {5}, {6}, {}, {7}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, 6, nil, 7, 5, -1}, Hidden: true, Note: "pushes and pops interleaved: the top is whatever was pushed last"},
			{Input: map[string]interface{}{"operations": []string{"MyStack", "push", "peek", "peek", "size", "pop", "isEmpty"}, "values": [][]int{{}, {4}, {}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, 4, 4, 1, 4, true}, Hidden: true, Note: "peek must not remove: reading twice gives the same answer and the size stays 1"},
			{Input: map[string]interface{}{"operations": []string{"MyStack", "push", "peek", "isEmpty", "pop", "isEmpty"}, "values": [][]int{{}, {0}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, 0, false, 0, true}, Hidden: true, Note: "0 is a real value (testing `if top:` mistakes it for empty)"},
			{Input: map[string]interface{}{"operations": []string{"MyStack", "pop", "size", "push", "size", "pop", "size"}, "values": [][]int{{}, {}, {}, {9}, {}, {}, {}}}, Expected: []interface{}{nil, -1, 0, nil, 1, 9, 0}, Hidden: true, Note: "popping an empty stack must not make the size negative"},
			{Input: map[string]interface{}{"operations": []string{"MyStack", "push", "push", "push", "pop", "pop", "size", "push", "size", "peek"}, "values": [][]int{{}, {1}, {2}, {3}, {}, {}, {}, {8}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, 3, 2, 1, nil, 2, 8}, Hidden: true, Note: "size follows pushes and pops; peek sees the newest push"},
		},
		TimeComplexity:  "O(1) per operation",
		SpaceComplexity: "O(n)",
		StarterCode: `class MyStack:
    def __init__(self):
        # Set up an empty stack (define your own node class if you chain nodes)
        pass

    def push(self, val):
        pass

    def pop(self):
        pass

    def peek(self):
        pass

    def isEmpty(self):
        pass

    def size(self):
        pass`,
		Harness: Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "A stack only ever touches one end. If you remember just the item on top, and which item sits below it, what do push and pop each have to change?"},
			{Level: 2, Type: "insight", Content: "Push and pop both work on the same end, the top. With a Python list that end is the right-hand end: append pushes and pop() removes. With nodes, keep a reference to the top node; a push makes a new node that points at the old top, and a pop moves the reference to the node below. Handle the empty case in pop and peek before touching the top. See Chapter 5.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Store the items: a top reference (None when empty) and a count, or a list.\n2. push: put the new item on top and add 1 to the count.\n3. pop: return -1 if empty; otherwise remember the top value, move the top down one, subtract 1, return the value.\n4. peek: return -1 if empty; otherwise the top value, with nothing removed.\n5. isEmpty and size read the stored state."},
			{Level: 4, Type: "skeleton", Content: "class Node:\n    def __init__(self, val, below=None):\n        self.val = val\n        self.below = below\n\n\nclass MyStack:\n    def __init__(self):\n        self.top = None\n        self.count = 0\n\n    def push(self, val):\n        self.top = Node(val, ____)\n        self.count += 1\n\n    def pop(self):\n        if self.top is None:\n            return -1\n        val = self.top.val\n        self.top = ____\n        self.count -= 1\n        return val\n\n    def peek(self):\n        if self.top is None:\n            return -1\n        return ____\n\n    def isEmpty(self):\n        return ____\n\n    def size(self):\n        return self.count"},
		},
		Solution: Solution{
			Code: `class Node:
    def __init__(self, val, below=None):
        self.val = val
        self.below = below


class MyStack:
    def __init__(self):
        self.top = None
        self.count = 0

    def push(self, val):
        self.top = Node(val, self.top)
        self.count += 1

    def pop(self):
        if self.top is None:
            return -1
        val = self.top.val
        self.top = self.top.below
        self.count -= 1
        return val

    def peek(self):
        if self.top is None:
            return -1
        return self.top.val

    def isEmpty(self):
        return self.top is None

    def size(self):
        return self.count`,
			Explanation: "A stack is a pile you can only add to or take from at the top. Here each node remembers the node below it, so push and pop only ever change the top reference. This is the interface that Valid Parentheses and every later stack problem rely on; Python's list.append and list.pop give you the same thing in one line.",
			KeyInsight:  "Only the top is ever visible, so a stack needs just two things: a reference to the top item and, for each item, a way to reach the one below it. Push and pop are then the same move in opposite directions, which is why both take constant time.",
			WhyItWorks:  "Invariant: following below links from top visits the items in the reverse of the order they were pushed, and count equals the number of nodes on that chain. An empty stack has top None and count 0. push puts a new node in front of the chain, so the newest item is first; pop removes the first node of the chain, so it returns the newest item still present, which is exactly last in, first out. peek reads the first node without changing the chain, and isEmpty and size read top and count. Each method does a fixed number of steps, so O(1) time, and the chain uses O(n) space for n items.",
			CommonMistakes: []string{
				"Letting pop or peek crash on an empty stack (an IndexError on a list, an AttributeError on None). Check for empty first and return -1.",
				"Making peek remove the item, so a second peek gives a different answer. peek only reads the top.",
				"Taking items from the wrong end, so the oldest item comes out first. That is a queue (first in, first out), not a stack.",
				"Subtracting 1 from the size on every pop, even when the stack was empty and nothing was removed, so the size goes negative.",
				"Testing `if top:` or `if not value:` to detect an empty stack, which treats a stored 0 as missing. Compare against None, or check the length.",
			},
			TimeComplexity:  "O(1) per operation: push and pop only change the top",
			SpaceComplexity: "O(n) for the items",
			Alternatives: []AltSolution{
				{Name: "Python list", Code: `class MyStack:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)

    def pop(self):
        return self.items.pop() if self.items else -1

    def peek(self):
        return self.items[-1] if self.items else -1

    def isEmpty(self):
        return not self.items

    def size(self):
        return len(self.items)`, TimeComplexity: "O(1) per operation (amortised)", SpaceComplexity: "O(n)", Explanation: "A Python list already behaves as a stack at its right-hand end: append is push and pop() is pop. This is what you will write in every later stack problem; the node version above shows what the list is doing for you."},
			},
			Walkthrough: []WalkthroughStep{
				{Title: "Define what is stored", Explanation: "Each node holds a value and a link to the node below it. The stack keeps the top node and a count, which starts at None and 0.", CodeSnippet: "self.top = None\nself.count = 0", LineStart: 9, LineEnd: 10},
				{Title: "Push onto the top", Explanation: "The new node points at the current top and becomes the top. After push(1), push(2) the chain is 2, then 1.", CodeSnippet: "self.top = Node(val, self.top)\nself.count += 1", LineStart: 13, LineEnd: 14},
				{Title: "Pop from the top", Explanation: "Guard the empty case first. Otherwise save the top value, move top down to the node below, and return the saved value. The stack 2, 1 returns 2 and is left as 1.", CodeSnippet: "if self.top is None:\n    return -1\nval = self.top.val\nself.top = self.top.below", LineStart: 17, LineEnd: 20},
				{Title: "Read without removing", Explanation: "peek returns the top value and changes nothing, so calling it twice gives the same answer.", CodeSnippet: "return self.top.val", LineStart: 27, LineEnd: 27},
				{Title: "Report the state", Explanation: "isEmpty is true exactly when there is no top node, and size is the count kept up to date by push and pop.", CodeSnippet: "return self.top is None", LineStart: 30, LineEnd: 30},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why do push and pop both cost O(1), however many items are stored?", Answer: "Both only touch the top: push puts one node in front of the chain, and pop unlinks the first node. Neither walks the stack, so the number of items does not matter. A list gives the same cost because it adds and removes at its right-hand end, where nothing has to shift."},
			{Prompt: "What would change if pop took from the bottom instead of the top?", Answer: "It would be a queue: first in, first out. The oldest item would come out first. Popping from the bottom of a list would also cost O(n), because every other item shifts down, which is why a real queue uses a deque instead."},
		},
	},
	{
		ID:              "drill-reverse-each-word",
		Number:          557,
		Kind:            KindDrill,
		Order:           2,
		Title:           "Reverse Each Word with a Stack",
		Slug:            "reverse-words-in-a-string-iii",
		Difficulty:      "Easy",
		Category:        "stack",
		Tags:            []string{"String", "Stack", "Foundations"},
		RelatedChapters: []int{5},
		Description: `You are given a string ` + "`" + `s` + "`" + ` made of words separated by single spaces (no leading or trailing spaces). Return a string with the letters of **each word reversed**, while the words stay in the same order and the single spaces stay where they were.

Use a stack to do the reversing: push the letters of a word one by one, and when the word ends, pop them all. Items come out of a stack in the opposite order they went in, so the popped letters spell the word backwards.`,
		Constraints: []string{
			"1 <= s.length <= 10000",
			"s contains printable ASCII characters",
			"Words are separated by exactly one space, with no leading or trailing space",
		},
		Examples: []Example{
			{Input: "s = \"Mr Ding\"", Output: "\"rM gniD\"", Explanation: "Push M, r, then at the space pop them: r, M. Then push D, i, n, g and pop at the end: g, n, i, D."},
			{Input: "s = \"a\"", Output: "\"a\"", Explanation: "One letter, so reversed it is the same."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "Mr Ding"}, Expected: "rM gniD", Note: "example 1: two words, each reversed in place"},
			{Input: map[string]interface{}{"s": "Let's take LeetCode contest"}, Expected: "s'teL ekat edoCteeL tsetnoc", Note: "example 2: punctuation and mixed case move with the letters"},
			{Input: map[string]interface{}{"s": "a"}, Expected: "a", Hidden: true, Note: "a single letter"},
			{Input: map[string]interface{}{"s": "abc"}, Expected: "cba", Hidden: true, Note: "a single word: the last word must be flushed at the end of the string"},
			{Input: map[string]interface{}{"s": "hello world"}, Expected: "olleh dlrow", Hidden: true, Note: "reversing the whole string would give \"dlrow olleh\"; the word order must stay"},
			{Input: map[string]interface{}{"s": "racecar level"}, Expected: "racecar level", Hidden: true, Note: "palindromic words stay unchanged"},
			{Input: map[string]interface{}{"s": "a bc d"}, Expected: "a cb d", Hidden: true, Note: "short words next to each other"},
			{Input: map[string]interface{}{"s": "one two three four"}, Expected: "eno owt eerht ruof", Hidden: true, Note: "four words; every one reversed, none skipped"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `def reverseWords(s):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "What happens to the order of items when you put them on a stack and take them off again? When should you start taking letters off?"},
			{Level: 2, Type: "insight", Content: "A stack hands items back in reverse order, so pushing the letters of one word and then popping them all spells that word backwards. A space marks the end of a word: that is the moment to pop everything. The last word has no space after it, so it needs a flush at the end of the string as well. See Chapter 5.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Make an empty output list and an empty stack.\n2. Push every letter that is not a space.\n3. On a space, pop every letter into the output, then add the space.\n4. Flush the last word, which has no space after it.\n5. Join the output into one string."},
			{Level: 4, Type: "skeleton", Content: "def reverseWords(s):\n    out = []\n    stack = []\n    for ch in s + \" \":  # the extra space flushes the last word\n        if ch != \" \":\n            stack.append(____)\n        else:\n            while ____:\n                out.append(stack.pop())\n            out.append(\" \")\n    out.pop()  # drop the space added after the last word\n    return \"\".join(out)"},
		},
		Solution: Solution{
			Code: `def reverseWords(s):
    out = []
    stack = []
    for ch in s + " ":  # the extra space flushes the last word
        if ch != " ":
            stack.append(ch)
        else:
            while stack:
                out.append(stack.pop())
            out.append(" ")
    out.pop()  # drop the space added after the last word
    return "".join(out)`,
			Explanation: "Pushing a word's letters and then popping them all reverses the word, because a stack returns items in the opposite order. This is the plainest use of the stack: push as things arrive, pop when something tells you the group is complete. Nothing is matched or compared yet; Remove Adjacent Duplicates, next, adds a comparison on the top.",
			KeyInsight:  "Push-then-pop reverses. The only decisions are when to pop (at a space) and not forgetting the last word, which has no space after it; appending a sentinel space to the input makes the last word end like every other.",
			WhyItWorks:  "Invariant: at the start of each iteration, stack holds exactly the letters of the current word read so far, in reading order, and out holds the finished words (reversed) with a space after each. A letter extends the current word. A space ends the word: popping the stack yields its letters last to first, which is the word reversed, and the space then separates it from the next one. The extra space appended to s ends the final word the same way. One push and one pop per letter, so O(n) time and O(n) space.",
			CommonMistakes: []string{
				"Forgetting to flush the last word, because the string ends without a space. The last word is then missing or unreversed. Add a sentinel space, or flush after the loop.",
				"Pushing the whole string and popping once at the end, which reverses the entire string and also the order of the words. Pop at each space.",
				"Pushing the spaces too, so they get popped in the wrong place. A space is the signal to pop; it is never stored.",
			},
			TimeComplexity:  "O(n): each character is pushed once and popped at most once",
			SpaceComplexity: "O(n) for the output and the stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "The stack holds the letters of the word being read, in reading order. out collects the finished words.", CodeSnippet: "out = []\nstack = []", LineStart: 2, LineEnd: 3},
				{Title: "Push the new item", Explanation: "A letter is not finished with yet, so it waits on the stack. For \"Mr\", M and then r are pushed.", CodeSnippet: "stack.append(ch)", LineStart: 6, LineEnd: 6},
				{Title: "Pop what the new item resolves", Explanation: "A space ends the word, so pop everything: r, then M. The popped letters are the word reversed.", CodeSnippet: "while stack:\n    out.append(stack.pop())", LineStart: 8, LineEnd: 9},
				{Title: "Handle what is left on the stack", Explanation: "The appended space ends the final word the same way. The space it adds to out is then removed, and the pieces are joined.", CodeSnippet: "out.pop()  # drop the space added after the last word", LineStart: 11, LineEnd: 11},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the code append a space to `s` before the loop?", Answer: "The last word has no space after it, so nothing would tell the loop to pop it. The extra space ends the last word like every other one, which removes a special case after the loop. The matching `out.pop()` removes the extra space again."},
			{Prompt: "What would the stack-based version return for the whole string if you only popped once, at the very end?", Answer: "It would return the whole string reversed, for example \"dlrow olleh\" for \"hello world\". The words would come out in reverse order too, because the stack reverses everything pushed since the last pop."},
		},
	},
	{
		ID:              "drill-remove-adjacent-duplicates",
		Number:          1047,
		Kind:            KindDrill,
		Order:           3,
		Title:           "Remove All Adjacent Duplicates",
		Slug:            "remove-all-adjacent-duplicates-in-string",
		Difficulty:      "Easy",
		Category:        "stack",
		Tags:            []string{"String", "Stack", "Foundations"},
		RelatedChapters: []int{5},
		Description: `You are given a string ` + "`" + `s` + "`" + ` of lowercase letters. One **duplicate removal** picks two equal letters that sit next to each other and deletes both; the pieces on either side then join up. Keep removing until no two equal letters are next to each other, and return the final string (which may be empty).

A removal can create a new equal pair where the pieces join: in ` + "`" + `"abba"` + "`" + ` removing ` + "`" + `bb` + "`" + ` leaves ` + "`" + `aa` + "`" + `, which can then be removed too. The result is the same whichever pair you remove first.`,
		Constraints: []string{
			"1 <= s.length <= 100000",
			"s consists of lowercase English letters",
		},
		Examples: []Example{
			{Input: "s = \"abbaca\"", Output: "\"ca\"", Explanation: "Remove \"bb\" to get \"aaca\". Remove \"aa\" to get \"ca\". No equal neighbours remain."},
			{Input: "s = \"azxxzy\"", Output: "\"ay\"", Explanation: "Remove \"xx\" to get \"azzy\". Remove \"zz\" to get \"ay\"."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "abbaca"}, Expected: "ca", Note: "example 1: removing one pair exposes another"},
			{Input: map[string]interface{}{"s": "azxxzy"}, Expected: "ay", Note: "example 2: a nested pair cascades outwards"},
			{Input: map[string]interface{}{"s": "a"}, Expected: "a", Hidden: true, Note: "a single letter stays"},
			{Input: map[string]interface{}{"s": "aa"}, Expected: "", Hidden: true, Note: "a string that removes down to the empty string"},
			{Input: map[string]interface{}{"s": "abba"}, Expected: "", Hidden: true, Note: "cascade: a single replace() pass leaves \"aa\""},
			{Input: map[string]interface{}{"s": "aaa"}, Expected: "a", Hidden: true, Note: "three equal letters: one pair goes and one letter stays (not all three)"},
			{Input: map[string]interface{}{"s": "abc"}, Expected: "abc", Hidden: true, Note: "no equal neighbours: nothing is removed"},
			{Input: map[string]interface{}{"s": "abccba"}, Expected: "", Hidden: true, Note: "a long palindrome removes completely, from the middle out"},
			{Input: map[string]interface{}{"s": "mississippi"}, Expected: "m", Hidden: true, Note: "several removals separated by letters that stay"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `def removeDuplicates(s):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Read the string from left to right. When a new letter arrives, which single earlier letter decides whether it cancels something?"},
			{Level: 2, Type: "insight", Content: "A new letter can only cancel the letter right before it among the survivors, and after a removal the survivor before that becomes the new neighbour. 'Look at the most recent survivor, maybe remove it' is a stack: if the top equals the new letter, pop; otherwise push. See Chapter 5.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Make an empty stack for the letters that survive so far.\n2. For each letter: if the stack is not empty and its top equals the letter, pop (both vanish).\n3. Otherwise push the letter.\n4. Join what is left on the stack, bottom to top."},
			{Level: 4, Type: "skeleton", Content: "def removeDuplicates(s):\n    stack = []\n    for ch in s:\n        if stack and stack[-1] == ____:\n            stack.____()\n        else:\n            stack.append(____)\n    return \"\".join(stack)"},
		},
		Solution: Solution{
			Code: `def removeDuplicates(s):
    stack = []
    for ch in s:
        if stack and stack[-1] == ch:
            stack.pop()
        else:
            stack.append(ch)
    return "".join(stack)`,
			Explanation: "The stack holds the letters that have survived so far, so its top is the left neighbour of the next letter. A new letter either cancels the top or waits on it. This is the matching idea from Valid Parentheses applied to a string, and it builds on the push-then-pop of the previous drill; Minimum String Length After Removing Substrings uses the same loop with a different pair rule.",
			KeyInsight:  "After a removal, the letter that becomes the new neighbour is exactly the next item down the stack. So one pass is enough: you never need to rescan the string, because the stack is already ordered by position.",
			WhyItWorks:  "Invariant: after reading a prefix of s, the stack is that prefix with all removals applied, and it has no two equal letters side by side. Reading a new letter c: if it equals the top, c and the top are an adjacent pair in the reduced prefix, so removing them gives the reduced prefix plus c, which is the stack with its top popped; if not, c is not part of any pair with the letter before it, so it is pushed and the stack still has no equal neighbours. The final stack is the reduced string. The order of removals does not matter. Each letter is pushed at most once and popped at most once, so O(n) time and O(n) space.",
			CommonMistakes: []string{
				"Doing a single pass of `s.replace(...)` for each letter pair, which misses pairs created by an earlier removal: \"abba\" needs two rounds. The stack handles the cascade in one pass.",
				"Removing a whole run of equal letters instead of one pair, so \"aaa\" becomes \"\" instead of \"a\". Only two letters cancel at a time.",
				"Comparing with `stack[0]` or the previous character of `s` instead of the top of the stack. After a removal, the neighbour is the item below, not the character that was just read.",
			},
			TimeComplexity:  "O(n): each letter is pushed and popped at most once",
			SpaceComplexity: "O(n) for the stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "The stack holds the letters that survived so far; its top is the left neighbour of the next letter.", CodeSnippet: "stack = []", LineStart: 2, LineEnd: 2},
				{Title: "Pop what the new item resolves", Explanation: "If the new letter equals the top, they are an adjacent pair, so the top is popped and the new letter is dropped. In \"abba\", the second b pops the first.", CodeSnippet: "if stack and stack[-1] == ch:\n    stack.pop()", LineStart: 4, LineEnd: 5},
				{Title: "Push the new item", Explanation: "Otherwise the letter waits on the stack. In \"abba\", a and b are pushed before the second b cancels.", CodeSnippet: "else:\n    stack.append(ch)", LineStart: 6, LineEnd: 7},
				{Title: "Handle what is left on the stack", Explanation: "The survivors, bottom to top, are the answer. \"abba\" ends with the second a popping the first, so the answer is the empty string.", CodeSnippet: "return \"\".join(stack)", LineStart: 8, LineEnd: 8},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is one left-to-right pass enough, when a removal can create a new pair?", Answer: "A new pair can only form between the letters on either side of the removal, and those are the new top of the stack and the next letter to be read. The stack always exposes the right neighbour, so the pair is found the moment the next letter arrives. Nothing before the top ever needs to be rescanned."},
			{Prompt: "What changes if equal letters had to be removed in groups of three?", Answer: "Compare the new letter with the top two items (or keep a count next to each stack entry), and pop both when all three match. The structure is the same: the stack still holds the survivors in order, and you still look only at the top."},
		},
	},
	{
		ID:              "drill-queue-using-stacks",
		Number:          232,
		Kind:            KindDrill,
		Order:           4,
		Title:           "Implement Queue using Stacks",
		Slug:            "implement-queue-using-stacks",
		Difficulty:      "Easy",
		Category:        "stack",
		Tags:            []string{"Stack", "Queue", "Design", "Foundations"},
		RelatedChapters: []int{5},
		Description: `Build a first-in, first-out (FIFO) queue, ` + "`" + `MyQueue` + "`" + `, using only **two stacks**. In your code, treat each stack as a Python list that you only touch with ` + "`" + `append` + "`" + ` (push), ` + "`" + `pop()` + "`" + ` from the end, ` + "`" + `[-1]` + "`" + ` for the top, and ` + "`" + `len` + "`" + ` or truthiness for emptiness; no ` + "`" + `pop(0)` + "`" + ` and no ` + "`" + `deque` + "`" + `. The class supports:

- ` + "`" + `push(x)` + "`" + `: add ` + "`" + `x` + "`" + ` to the back of the queue.
- ` + "`" + `pop()` + "`" + `: remove the element at the front of the queue and return it.
- ` + "`" + `peek()` + "`" + `: return the element at the front without removing it.
- ` + "`" + `empty()` + "`" + `: return ` + "`" + `True` + "`" + ` if the queue holds no elements, otherwise ` + "`" + `False` + "`" + `.

` + "`" + `pop` + "`" + ` and ` + "`" + `peek` + "`" + ` are only called when the queue is not empty. Moving everything from one stack onto another reverses its order, which is the whole trick. Aim for each operation to cost O(1) on average (amortised).

How the tests call it: ` + "`" + `operations` + "`" + ` lists the method names (the first is the constructor) and ` + "`" + `values` + "`" + ` the arguments for each; the expected list has null for calls that return nothing.`,
		Constraints: []string{
			"1 <= x <= 9",
			"At most 100 calls to push, pop, peek and empty",
			"pop and peek are only called on a non-empty queue",
		},
		Examples: []Example{
			{Input: "[\"MyQueue\",\"push\",\"push\",\"peek\",\"pop\",\"empty\"]\n[[],[1],[2],[],[],[]]", Output: "[null,null,null,1,1,false]", Explanation: "The queue is 1 then 2. The front is 1, so peek() and pop() both return 1. The 2 is left, so empty() is false."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"MyQueue", "push", "push", "peek", "pop", "empty"}, "values": [][]int{{}, {1}, {2}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, 1, 1, false}, Note: "example: the front is the first element pushed"},
			{Input: map[string]interface{}{"operations": []string{"MyQueue", "push", "pop", "empty"}, "values": [][]int{{}, {1}, {}, {}}}, Expected: []interface{}{nil, nil, 1, true}, Note: "a single element goes in and out; the queue ends empty"},
			{Input: map[string]interface{}{"operations": []string{"MyQueue", "push", "push", "pop", "push", "pop", "pop", "empty"}, "values": [][]int{{}, {1}, {2}, {}, {3}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, 1, nil, 2, 3, true}, Hidden: true, Note: "pushing between pops: 3 must wait behind 2 (moving the new push onto a non-empty outbox breaks the order)"},
			{Input: map[string]interface{}{"operations": []string{"MyQueue", "push", "push", "peek", "push", "pop", "pop", "pop"}, "values": [][]int{{}, {1}, {2}, {}, {3}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, 1, nil, 1, 2, 3}, Hidden: true, Note: "peek moves elements across; later pushes must still queue up behind them"},
			{Input: map[string]interface{}{"operations": []string{"MyQueue", "push", "peek", "empty"}, "values": [][]int{{}, {4}, {}, {}}}, Expected: []interface{}{nil, nil, 4, false}, Hidden: true, Note: "after peek moves the element to the outbox, the queue is still not empty (empty must check both stacks)"},
			{Input: map[string]interface{}{"operations": []string{"MyQueue", "push", "push", "push", "pop", "pop", "pop", "empty"}, "values": [][]int{{}, {5}, {6}, {7}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, 5, 6, 7, true}, Hidden: true, Note: "FIFO order 5, 6, 7 (a single stack would return 7, 6, 5)"},
			{Input: map[string]interface{}{"operations": []string{"MyQueue", "push", "push", "push", "pop", "push", "push", "pop", "pop", "pop", "pop", "empty"}, "values": [][]int{{}, {1}, {2}, {3}, {}, {4}, {5}, {}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, 1, nil, nil, 2, 3, 4, 5, true}, Hidden: true, Note: "a longer mix that drains the outbox and refills it twice"},
			{Input: map[string]interface{}{"operations": []string{"MyQueue", "push", "pop", "push", "peek", "pop", "empty"}, "values": [][]int{{}, {8}, {}, {9}, {}, {}, {}}}, Expected: []interface{}{nil, nil, 8, nil, 9, 9, true}, Hidden: true, Note: "the queue is reused after becoming empty"},
		},
		TimeComplexity:  "O(1) amortised per operation",
		SpaceComplexity: "O(n)",
		StarterCode: `class MyQueue:
    def __init__(self):
        # Two stacks: use plain lists with append / pop() / [-1]
        pass

    def push(self, x):
        pass

    def pop(self):
        pass

    def peek(self):
        pass

    def empty(self):
        pass`,
		Harness: Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "A stack gives back the newest item first, but a queue must give back the oldest. What happens to the order of a pile of items if you move them one by one onto a second pile?"},
			{Level: 2, Type: "insight", Content: "Moving every item from one stack to another reverses their order, so the oldest item ends on top. Keep an inbox for pushes and an outbox for pops and peeks. Move items across only when the outbox is empty, and then move all of them: items already in the outbox are older than anything in the inbox, so they must be served first. Each item is moved at most once, which is why the cost is O(1) amortised. See Chapter 5.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Keep two lists: inbox (new pushes) and outbox (items in front-first order, front on top).\n2. push: append to inbox.\n3. A helper refill: only if outbox is empty, pop everything from inbox onto outbox.\n4. pop: refill, then pop from outbox.\n5. peek: refill, then read outbox[-1].\n6. empty: true only if both stacks are empty."},
			{Level: 4, Type: "skeleton", Content: "class MyQueue:\n    def __init__(self):\n        self.inbox = []\n        self.outbox = []\n\n    def push(self, x):\n        self.inbox.append(x)\n\n    def _refill(self):\n        if ____:\n            while self.inbox:\n                self.outbox.append(self.inbox.____())\n\n    def pop(self):\n        self._refill()\n        return self.outbox.pop()\n\n    def peek(self):\n        self._refill()\n        return ____\n\n    def empty(self):\n        return not self.inbox and ____"},
		},
		Solution: Solution{
			Code: `class MyQueue:
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
        return not self.inbox and not self.outbox`,
			Explanation: "One stack takes the pushes; the other holds the items in the order they must leave. Moving a whole stack onto another reverses it, which turns newest-first into oldest-first. This drill shows what 'reverse by popping' (Reverse Each Word) is good for, and the amortised argument, each item moved at most once, comes back in later chapters.",
			KeyInsight:  "Only refill the outbox when it is empty, and then move everything. The outbox then always holds the oldest items in front-first order, and every item makes exactly one trip from inbox to outbox, so the total transfer work is at most the number of pushes.",
			WhyItWorks:  "Invariant: reading the outbox from its top down and then the inbox from its bottom up lists all elements in queue order, front first. push adds at the inbox top, which is the back of the queue, so the invariant holds. When the outbox is empty, popping the inbox item by item onto the outbox reverses it, so the oldest element lands on top: the invariant still holds, and the front is now outbox[-1]. pop and peek read that front. If the outbox is not empty its top is already the front, so no transfer is allowed (it would put newer elements under or over older ones). Each element is pushed to the inbox once, moved once and popped once, so n operations cost O(n) in total: O(1) amortised.",
			CommonMistakes: []string{
				"Moving the inbox onto the outbox on every pop or peek, even when the outbox still has items. Newer items end up on top of older ones and come out first. Only refill when the outbox is empty.",
				"Moving only some of the inbox, or only when asked to peek. When the outbox is refilled, move everything; otherwise the front is wrong.",
				"Checking only one stack in `empty`. After a peek the element lives in the outbox while the inbox is empty, so both must be checked.",
				"Returning the top of the inbox from pop, which makes the class a stack (last in, first out) instead of a queue.",
			},
			TimeComplexity:  "O(1) amortised: each element is moved from inbox to outbox at most once",
			SpaceComplexity: "O(n) for the elements",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stacks hold", Explanation: "The inbox takes new pushes (newest on top). The outbox holds items ready to leave, with the front of the queue on top.", CodeSnippet: "self.inbox = []\nself.outbox = []", LineStart: 3, LineEnd: 4},
				{Title: "Push the new item", Explanation: "A push goes straight onto the inbox. It is the back of the queue, so it must wait behind everything already there.", CodeSnippet: "self.inbox.append(x)", LineStart: 7, LineEnd: 7},
				{Title: "Transfer only when the outbox is empty", Explanation: "Popping the inbox onto the outbox reverses it, so the oldest item is on top. With 1, 2, 3 pushed, 3 is moved first and 1 last, leaving 1 on top.", CodeSnippet: "if not self.outbox:\n    while self.inbox:\n        self.outbox.append(self.inbox.pop())", LineStart: 10, LineEnd: 12},
				{Title: "Read the front from the outbox top", Explanation: "After a refill, the front of the queue is the outbox top: pop removes it and peek reads it.", CodeSnippet: "return self.outbox.pop()", LineStart: 16, LineEnd: 16},
				{Title: "Check both stacks for empty", Explanation: "Elements may be in either stack, so the queue is empty only when both are.", CodeSnippet: "return not self.inbox and not self.outbox", LineStart: 23, LineEnd: 23},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the transfer wait until the outbox is empty?", Answer: "Items in the outbox are older than anything in the inbox. Moving inbox items onto a non-empty outbox would put newer items above older ones, so a pop would return a newer item first and break first in, first out. Waiting until the outbox is empty guarantees everything moved is older than nothing left behind."},
			{Prompt: "A pop can cost O(n) when it triggers a transfer, so why is the cost O(1) amortised?", Answer: "Each element is moved from the inbox to the outbox at most once in its lifetime, and each move is paid for by the push that put it there. Over any run of n operations the total moving work is at most n, so the average cost per operation is constant even though one pop can be expensive."},
		},
	},
	{
		ID:              "drill-next-greater-element",
		Number:          0,
		Kind:            KindDrill,
		Order:           5,
		Title:           "Next Greater Element to the Right",
		Slug:            "next-greater-element-to-the-right",
		Difficulty:      "Easy",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Monotonic Stack", "Foundations"},
		RelatedChapters: []int{5},
		Description: `Given an array of integers ` + "`" + `nums` + "`" + `, return an array ` + "`" + `answer` + "`" + ` of the same length where ` + "`" + `answer[i]` + "`" + ` is the **first value to the right of position i that is strictly greater than ` + "`" + `nums[i]` + "`" + `**, or ` + "`" + `-1` + "`" + ` if there is none.

"First" means closest to ` + "`" + `i` + "`" + `: scan right from ` + "`" + `i + 1` + "`" + ` and stop at the first larger value. A value equal to ` + "`" + `nums[i]` + "`" + ` is not greater. The last element has nothing to its right, so its answer is always ` + "`" + `-1` + "`" + `.

A pair of loops checks every ` + "`" + `i` + "`" + ` against every later position, which is slow for long arrays. Try to do it in a single left-to-right pass with a stack. This is the monotonic-stack template that Daily Temperatures and the harder problems in the Stack ladder are built from.`,
		Constraints: []string{
			"1 <= nums.length <= 10000",
			"-1000000 <= nums[i] <= 1000000",
		},
		Examples: []Example{
			{Input: "nums = [2,1,2,4,3]", Output: "[4,2,4,-1,-1]", Explanation: "For the first 2, the values 1 and 2 are not greater, so its answer is 4. For the 1, the next value 2 is greater. For the second 2 it is 4. Nothing to the right of 4 is larger, nor of 3."},
			{Input: "nums = [3,2,1]", Output: "[-1,-1,-1]", Explanation: "Each value is followed only by smaller values."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 1, 2, 4, 3}}, Expected: []int{4, 2, 4, -1, -1}, Note: "example 1: some elements wait for several steps"},
			{Input: map[string]interface{}{"nums": []int{3, 2, 1}}, Expected: []int{-1, -1, -1}, Note: "example 2: decreasing, so nobody finds a greater value"},
			{Input: map[string]interface{}{"nums": []int{5}}, Expected: []int{-1}, Hidden: true, Note: "a single element has nothing to its right"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}}, Expected: []int{2, 3, 4, -1}, Hidden: true, Note: "increasing: every element is resolved by its immediate neighbour"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 2}}, Expected: []int{-1, -1, -1}, Hidden: true, Note: "equal values are not greater (a <= comparison gives the wrong answer)"},
			{Input: map[string]interface{}{"nums": []int{3, 2, 1, 4}}, Expected: []int{4, 4, 4, -1}, Hidden: true, Note: "one big element resolves several waiting elements at once (`if` instead of `while` pops only one)"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 2, 4}}, Expected: []int{3, 4, 4, -1}, Hidden: true, Note: "answers are values, not indices, and the nearest greater is used"},
			{Input: map[string]interface{}{"nums": []int{-1, -2, 0, -3}}, Expected: []int{0, 0, -1, -1}, Hidden: true, Note: "negative numbers and a zero"},
			{Input: map[string]interface{}{"nums": []int{5, 1, 2, 3, 6, 0, 7}}, Expected: []int{6, 2, 3, 6, 7, 7, -1}, Hidden: true, Note: "a mix of waiting elements resolved at different times"},
			{Input: map[string]interface{}{"nums": stackDrillDescendingThen(3000, 5000)}, Expected: stackDrillAllThenLast(3001, 5000), Hidden: true, Note: "a long strictly decreasing run, all resolved by the last element"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `def nextGreater(nums):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Going left to right, at the moment you meet a new number, which of the earlier numbers just found out their answer? Which earlier numbers are still waiting?"},
			{Level: 2, Type: "insight", Content: "Keep the positions that are still waiting for a greater value. A new number resolves every waiting position whose value is strictly smaller than it, and those are always the most recent ones, so they sit on top of the stack. The waiting values never increase from bottom to top: this is a monotonic stack. Pop while the top is smaller, write the answer for each popped position, then push the new position. See Chapter 5.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Make the answer array filled with -1 (the default for anything never resolved).\n2. Make an empty stack of indices that are still waiting.\n3. For each number: while the stack is not empty and the number at the top index is smaller, pop it and set its answer to the current number.\n4. Push the current index.\n5. Return the answer."},
			{Level: 4, Type: "skeleton", Content: "def nextGreater(nums):\n    answer = [-1] * len(nums)\n    stack = []  # indices still waiting\n    for i, num in enumerate(nums):\n        while stack and nums[stack[-1]] ____ num:\n            answer[____] = num\n        stack.append(____)\n    return answer"},
		},
		Solution: Solution{
			Code: `def nextGreater(nums):
    answer = [-1] * len(nums)
    stack = []  # indices still waiting; their values never increase bottom to top
    for i, num in enumerate(nums):
        while stack and nums[stack[-1]] < num:
            answer[stack.pop()] = num
        stack.append(i)
    return answer`,
			Explanation: "The stack holds the positions that have not yet seen a larger value, with values that never increase from bottom to top. A new number pops every smaller waiting value and is their answer, then waits itself. This is the monotonic-stack template: Daily Temperatures records the distance instead of the value, and the previous-smaller drill that follows mirrors it.",
			KeyInsight:  "Instead of searching to the right for each element, let each element wait on a stack and be resolved the moment a larger value arrives. The stack stays ordered, so the elements a new number can resolve are always at the top, and anything it cannot resolve cannot be resolved by it at all.",
			WhyItWorks:  "Invariant: before reading index i, the stack holds exactly the earlier indices with no greater value yet seen, and their values never increase from bottom to top. If nums[i] is greater than the top, that top has found its first greater value (nothing between them was greater, or it would have been popped earlier), so its answer is nums[i]. Popping continues while the new number beats the top; once the top is not smaller, everything beneath it is at least as large, so nothing deeper can be resolved either. Pushing i keeps the order valid. Indices left at the end never found a greater value and keep -1. Each index is pushed and popped at most once: O(n) time, O(n) space.",
			CommonMistakes: []string{
				"Popping while the top is `<=` the new number, which counts an equal value as greater. [2, 2, 2] must give -1 for all three.",
				"Using `if` instead of `while`, so a large number resolves only one waiting element and leaves the rest with -1. One new value can resolve many.",
				"Starting the answer list at 0 (or leaving it empty) instead of -1, so elements that never find a greater value report the wrong thing.",
				"Storing values on the stack instead of indices, which cannot tell you which answer slot to fill when there are duplicates. Store the index and look the value up.",
			},
			TimeComplexity:  "O(n): each index is pushed once and popped at most once",
			SpaceComplexity: "O(n) for the stack and the answer",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "answer starts at -1 for everyone, which is right for elements that never get resolved. The stack holds the indices still waiting for a greater value.", CodeSnippet: "answer = [-1] * len(nums)\nstack = []  # indices still waiting; their values never increase bottom to top", LineStart: 2, LineEnd: 3},
				{Title: "Pop what the new item resolves", Explanation: "While the top waiting value is strictly smaller than the new number, the new number is its first greater value. In [3,2,1,4], the 4 pops 1, then 2, then 3.", CodeSnippet: "while stack and nums[stack[-1]] < num:", LineStart: 5, LineEnd: 5},
				{Title: "Record the answer at pop time", Explanation: "Fill the answer slot of the popped index with the new number. For [2,1,2,4,3], the second 2 pops the 1 and writes answer[1] = 2.", CodeSnippet: "answer[stack.pop()] = num", LineStart: 6, LineEnd: 6},
				{Title: "Push the new item", Explanation: "The new number now waits for its own greater value. Everything left below it is at least as large, so the order is kept.", CodeSnippet: "stack.append(i)", LineStart: 7, LineEnd: 7},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to stop popping at the first stack entry that is not smaller than the new number?", Answer: "The values on the stack never increase from bottom to top, so every entry below the top is at least as large as the top. If the top is not smaller than the new number, nothing underneath is either, so the new number cannot resolve any of them and the loop can stop."},
			{Prompt: "What changes if you want the next greater element to the left instead?", Answer: "Scan the array from right to left with the same stack, or equivalently ask for the previous greater element (the mirrored template in the next drill). Another option is to run the same code on the reversed array and reverse the answers, since the right-hand neighbour in the reversed array is the left-hand neighbour in the original."},
		},
	},
	{
		ID:              "drill-previous-smaller-element",
		Number:          0,
		Kind:            KindDrill,
		Order:           6,
		Title:           "Previous Smaller Element",
		Slug:            "previous-smaller-element",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Array", "Stack", "Monotonic Stack", "Foundations"},
		RelatedChapters: []int{5},
		Description: `Given an array of integers ` + "`" + `nums` + "`" + `, return an array ` + "`" + `answer` + "`" + ` of the same length where ` + "`" + `answer[i]` + "`" + ` is the **index** of the nearest element to the **left** of position ` + "`" + `i` + "`" + ` whose value is **strictly smaller** than ` + "`" + `nums[i]` + "`" + `, or ` + "`" + `-1` + "`" + ` if there is none.

"Nearest" means the largest index ` + "`" + `j < i` + "`" + ` with ` + "`" + `nums[j] < nums[i]` + "`" + `. A value equal to ` + "`" + `nums[i]` + "`" + ` is not smaller. The first element has nothing to its left, so its answer is always ` + "`" + `-1` + "`" + `.

This mirrors the previous drill: there each element waited for a greater value to its right; here each element looks back for a smaller one to its left. It is the building block for Largest Rectangle in Histogram, which needs the nearest smaller bar on each side of every bar.`,
		Constraints: []string{
			"1 <= nums.length <= 10000",
			"-1000000 <= nums[i] <= 1000000",
		},
		Examples: []Example{
			{Input: "nums = [2,1,5,6,2,3]", Output: "[-1,-1,1,2,1,4]", Explanation: "Index 2 (value 5): the nearest smaller to its left is the 1 at index 1. Index 4 (value 2): 6 and 5 are not smaller, the 1 at index 1 is. Index 5 (value 3): the 2 at index 4."},
			{Input: "nums = [3,2,1]", Output: "[-1,-1,-1]", Explanation: "Every value is preceded only by larger values."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{2, 1, 5, 6, 2, 3}}, Expected: []int{-1, -1, 1, 2, 1, 4}, Note: "example 1: some elements look back past several larger ones"},
			{Input: map[string]interface{}{"nums": []int{3, 2, 1}}, Expected: []int{-1, -1, -1}, Note: "example 2: decreasing, so no element has a smaller one to its left"},
			{Input: map[string]interface{}{"nums": []int{4}}, Expected: []int{-1}, Hidden: true, Note: "a single element"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}}, Expected: []int{-1, 0, 1}, Hidden: true, Note: "increasing: each answer is the index just before"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 2}}, Expected: []int{-1, -1, -1}, Hidden: true, Note: "equal values are not smaller (a `>` pop condition gives [-1,0,1])"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 4, 5, 2}}, Expected: []int{-1, 0, 1, 2, 0}, Hidden: true, Note: "the new value must pop every larger entry, not just one (`if` instead of `while`)"},
			{Input: map[string]interface{}{"nums": []int{3, 7, 8, 4}}, Expected: []int{-1, 0, 1, 0}, Hidden: true, Note: "the answer is an index, not the value"},
			{Input: map[string]interface{}{"nums": []int{5, -3, -3, 0, -8, 2}}, Expected: []int{-1, -1, -1, 2, -1, 4}, Hidden: true, Note: "negatives and repeated values"},
			{Input: map[string]interface{}{"nums": []int{4, 5, 2, 10, 8, 3, 3, 6}}, Expected: []int{-1, 0, -1, 2, 2, 2, 2, 6}, Hidden: true, Note: "a longer mix: entries are popped and re-exposed at different times"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `def prevSmaller(nums):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Going left to right, which of the earlier numbers could still be the answer for some later number, and which can never be the answer again once a smaller number has appeared after them?"},
			{Level: 2, Type: "insight", Content: "An earlier number that is larger than or equal to the current one can never be the nearest smaller for the current number or for anything after it, because the current number is closer and not larger. So pop those. What remains has values that strictly increase from bottom to top, and the top is the answer. This is the monotonic stack from the previous drill, with the answer read off the stack after popping instead of at pop time. See Chapter 5.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Make the answer array filled with -1 and an empty stack of indices.\n2. For each index i: while the stack is not empty and the value at the top is greater than or equal to nums[i], pop.\n3. If the stack is not empty, the top index is the answer for i.\n4. Push i.\n5. Return the answer."},
			{Level: 4, Type: "skeleton", Content: "def prevSmaller(nums):\n    answer = [-1] * len(nums)\n    stack = []  # indices; values strictly increase bottom to top\n    for i, num in enumerate(nums):\n        while stack and nums[stack[-1]] ____ num:\n            stack.pop()\n        if stack:\n            answer[i] = ____\n        stack.append(____)\n    return answer"},
		},
		Solution: Solution{
			Code: `def prevSmaller(nums):
    answer = [-1] * len(nums)
    stack = []  # indices; values strictly increase from bottom to top
    for i, num in enumerate(nums):
        while stack and nums[stack[-1]] >= num:
            stack.pop()
        if stack:
            answer[i] = stack[-1]
        stack.append(i)
    return answer`,
			Explanation: "The stack keeps the candidates for 'nearest smaller on my left': indices whose values strictly increase from bottom to top. A new number discards every candidate that is not smaller, and the one left on top is its answer. It mirrors Next Greater Element: there the answer is written when elements are popped, here it is read from the stack after the popping. Both are needed for Largest Rectangle in Histogram.",
			KeyInsight:  "An earlier value that is larger than or equal to the current one is useless as an answer from now on: the current value is closer and not larger, so it beats it for every later element. Discarding those leaves a stack in which the nearest smaller value is simply the top.",
			WhyItWorks:  "Invariant: before reading index i, the stack holds the indices j < i that have no later index k with i > k > j and nums[k] <= nums[j] (the candidates), in increasing order, so their values strictly increase from bottom to top. Popping while the top is >= nums[i] removes the candidates that i is closer than and not larger than. The top is then the largest index with a strictly smaller value: any index between it and i was popped because something at least as small followed it, and that something is either on the stack or was itself popped in the same way, down to an entry that is smaller than nums[i] only if it is on the stack. Pushing i keeps the order strict. Each index is pushed and popped at most once: O(n) time, O(n) space.",
			CommonMistakes: []string{
				"Popping only while the top is `>` the new value, which keeps an equal value on the stack and returns it as the 'smaller' one. The condition for popping is `>=`.",
				"Using `if` instead of `while`, so a small number removes just one larger candidate and then reads a larger value as its answer.",
				"Returning the value of the smaller element instead of its index. The answer is a position.",
				"Reading the answer before popping, so it is the previous element whatever its size.",
			},
			TimeComplexity:  "O(n): each index is pushed once and popped at most once",
			SpaceComplexity: "O(n) for the stack and the answer",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "The stack holds candidate indices whose values strictly increase bottom to top. answer starts at -1 for elements with no smaller value on their left.", CodeSnippet: "answer = [-1] * len(nums)\nstack = []  # indices; values strictly increase from bottom to top", LineStart: 2, LineEnd: 3},
				{Title: "Pop what the new item resolves", Explanation: "Candidates that are not smaller than the new value can never be an answer again, so they are popped. In [1,3,4,5,2], the 2 pops 5, 4 and 3, leaving the 1.", CodeSnippet: "while stack and nums[stack[-1]] >= num:\n    stack.pop()", LineStart: 5, LineEnd: 6},
				{Title: "Record the answer", Explanation: "What is left on top is the nearest strictly smaller value on the left. An empty stack means there is none, so the -1 stays.", CodeSnippet: "if stack:\n    answer[i] = stack[-1]", LineStart: 7, LineEnd: 8},
				{Title: "Push the new item", Explanation: "The new index becomes the newest candidate for everything to its right.", CodeSnippet: "stack.append(i)", LineStart: 9, LineEnd: 9},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can a popped index never be the answer for a later element?", Answer: "A popped value is greater than or equal to the value that popped it, and that value sits closer to every later element. Whenever the popped value would qualify as 'smaller than a later number', the closer, not-larger value qualifies too, and is the better answer. So nothing is lost by discarding it."},
			{Prompt: "How would you get the nearest smaller on the right as well?", Answer: "Run the same code mirrored: either scan from the right, or notice that the pops from the left-to-right pass already tell you. When an index is popped by a smaller value at position i, that i is its nearest smaller on the right (the first smaller value to arrive). Largest Rectangle in Histogram uses both sides."},
		},
	},
	{
		ID:              "drill-evaluate-expression",
		Number:          224,
		Kind:            KindDrill,
		Order:           7,
		Title:           "Evaluate an Expression with Parentheses",
		Slug:            "basic-calculator",
		Difficulty:      "Medium",
		Category:        "stack",
		Tags:            []string{"Math", "String", "Stack", "Foundations"},
		RelatedChapters: []int{5},
		Description: `You are given a string ` + "`" + `s` + "`" + ` that is a valid arithmetic expression made of non-negative integers (any number of digits), the operators ` + "`" + `+` + "`" + ` and ` + "`" + `-` + "`" + `, parentheses ` + "`" + `(` + "`" + ` and ` + "`" + `)` + "`" + `, and spaces. Return the value of the expression as an integer. There is no ` + "`" + `*` + "`" + ` or ` + "`" + `/` + "`" + `, and you may not use ` + "`" + `eval` + "`" + `.

A ` + "`" + `-` + "`" + ` can also be unary (apply to what follows with nothing before it): ` + "`" + `"-2"` + "`" + `, ` + "`" + `"(-2)"` + "`" + ` and ` + "`" + `"3-(-2)"` + "`" + ` are all valid. Parentheses can nest, and a ` + "`" + `-` + "`" + ` in front of a parenthesised group flips the sign of everything inside it: ` + "`" + `1-(2+3)` + "`" + ` is ` + "`" + `1 - 5` + "`" + `.

Parentheses are what make this a stack problem. When a ` + "`" + `(` + "`" + ` opens, you must set aside what you have computed and the sign in front of the group, work out the group on its own, and combine the two when the ` + "`" + `)` + "`" + ` closes it. That is what a recursive solution does implicitly; here you keep the saved state on an explicit stack.`,
		Constraints: []string{
			"1 <= s.length <= 30000",
			"s contains digits, '+', '-', '(', ')' and spaces",
			"s is a valid expression; every intermediate and final value fits in a 32-bit signed integer",
		},
		Examples: []Example{
			{Input: "s = \"1 + 1\"", Output: "2", Explanation: "Spaces are ignored."},
			{Input: "s = \"1-(2+3)\"", Output: "-4", Explanation: "The group evaluates to 5, and the minus in front flips it: 1 - 5 = -4."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "1 + 1"}, Expected: 2, Note: "example 1: spaces are ignored"},
			{Input: map[string]interface{}{"s": "1-(2+3)"}, Expected: -4, Note: "example 2: a minus in front of a group flips the whole group"},
			{Input: map[string]interface{}{"s": "123"}, Expected: 123, Hidden: true, Note: "a single multi-digit number with no operator (the last number must be added at the end)"},
			{Input: map[string]interface{}{"s": "10+20"}, Expected: 30, Hidden: true, Note: "multi-digit numbers (digits are not separate operands)"},
			{Input: map[string]interface{}{"s": " 2-1 + 2 "}, Expected: 3, Hidden: true, Note: "spaces around numbers and operators; left to right: (2-1)+2"},
			{Input: map[string]interface{}{"s": "(1+(4+5+2)-3)+(6+8)"}, Expected: 23, Hidden: true, Note: "nested and side-by-side groups: 23"},
			{Input: map[string]interface{}{"s": "2-(5-6)"}, Expected: 3, Hidden: true, Note: "a negative group result: 2 - (-1) = 3"},
			{Input: map[string]interface{}{"s": "-(2+3)"}, Expected: -5, Hidden: true, Note: "a leading unary minus on a group"},
			{Input: map[string]interface{}{"s": "2-(-3)"}, Expected: 5, Hidden: true, Note: "a unary minus inside a group"},
			{Input: map[string]interface{}{"s": "10-(2-(3-(4-5)))"}, Expected: 12, Hidden: true, Note: "three levels of nesting, each flipping the sign inside it"},
			{Input: map[string]interface{}{"s": "(((7)))"}, Expected: 7, Hidden: true, Note: "redundant parentheses around one number"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `def calculate(s):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When you reach an opening parenthesis, what do you have to remember so that, after the closing parenthesis, you can carry on as if the group were one number?"},
			{Level: 2, Type: "insight", Content: "Evaluate left to right, keeping a running total, the sign of the next number and the number being read. At a '(' push the total and the sign in front of the group onto a stack and start a fresh total for the inside. At the matching ')' finish the group's total, pop the saved total and sign, and combine them: saved total + saved sign * group total. This is the explicit-stack version of recursion; the stack holds what a function call would have held. See Chapter 5.", ChapterRef: 5},
			{Level: 3, Type: "plan", Content: "1. Keep total = 0, sign = 1, number = 0, and an empty stack.\n2. Digit: extend the number (number * 10 + digit).\n3. '+' or '-': add sign * number to total, reset the number, set the sign for the next number.\n4. '(': push (total, sign), then reset total and sign for the inside of the group.\n5. ')': add the last number inside, pop the saved pair, and set total = saved total + saved sign * total.\n6. After the loop, add the last number and return the total."},
			{Level: 4, Type: "skeleton", Content: "def calculate(s):\n    total = 0\n    sign = 1\n    number = 0\n    stack = []  # (total, sign) saved when a \"(\" opens\n    for ch in s:\n        if ch.isdigit():\n            number = number * 10 + int(ch)\n        elif ch == \"+\" or ch == \"-\":\n            total += sign * number\n            number = 0\n            sign = 1 if ch == \"+\" else -1\n        elif ch == \"(\":\n            stack.append((____, ____))\n            total = 0\n            sign = 1\n        elif ch == \")\":\n            total += sign * number\n            number = 0\n            saved_total, saved_sign = stack.pop()\n            total = ____ + saved_sign * ____\n    return total + sign * number"},
		},
		Solution: Solution{
			Code: `def calculate(s):
    total = 0
    sign = 1
    number = 0
    stack = []  # (total, sign) saved when a "(" opens
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
    return total + sign * number`,
			Explanation: "Between parentheses the expression is a plain running sum, so the only difficulty is a group. On '(' the code saves the total so far and the sign in front of the group on a stack and starts again from zero; on ')' it pops them and folds the finished group in as a single signed number. It builds on Evaluate Reverse Polish Notation (a stack of pending values) by saving state instead of operands.",
			KeyInsight:  "A parenthesised group is a sub-problem with its own running total. Save the outer state when the group opens, restore it when it closes, and the same left-to-right loop handles every depth of nesting; the stack is just the call stack of a recursive solution written out by hand.",
			WhyItWorks:  "Invariant: at any point, total is the value of the part of the current group read so far (not counting the pending number), sign is the sign that applies to the number being read, and the stack holds one (total, sign) pair for each group still open, outermost at the bottom. A digit extends the number; an operator folds sign * number into the total and sets the sign for the next number. A '(' saves the outer state and starts a fresh group, so everything inside is evaluated as if it were a whole expression. A ')' finishes the inner total, restores the saved pair, and adds the group's value with the sign that was in front of it, which is correct because a minus before a group negates the sum of everything inside. Each character is handled once and each pair is pushed and popped once: O(n) time and O(n) space for the nesting depth.",
			CommonMistakes: []string{
				"Ignoring the sign in front of the parenthesis: pushing only the total, so `1-(2+3)` becomes 6 instead of -4. The sign in front of the group must be saved and applied when the group closes.",
				"Treating every digit as its own number instead of building it up with `number * 10 + digit`, so `10+20` gives 3 rather than 30.",
				"Forgetting the last number, which has no operator after it: `123` returns 0. Add `sign * number` after the loop (and before closing each group).",
				"Not adding the pending number to the total before restoring at a ')', so the last number inside a group is lost.",
			},
			TimeComplexity:  "O(n): one pass, each pair pushed and popped once",
			SpaceComplexity: "O(n) for the stack (at most the nesting depth)",
			Walkthrough: []WalkthroughStep{
				{Title: "Decide what the stack holds", Explanation: "total is the sum so far in the current group, sign is the sign for the number being read, and number is the digits read so far. The stack will hold the outer state for each open group.", CodeSnippet: "total = 0\nsign = 1\nnumber = 0\nstack = []  # (total, sign) saved when a \"(\" opens", LineStart: 2, LineEnd: 5},
				{Title: "Read the number and fold it in", Explanation: "Digits build the number. An operator adds the finished number with its sign to the total, and the new operator sets the next sign. For \"10-4\": 1, 10, then '-' adds +10 and sets sign -1.", CodeSnippet: "number = number * 10 + int(ch)\nelif ch == \"+\" or ch == \"-\":\n    total += sign * number\n    number = 0\n    sign = 1 if ch == \"+\" else -1", LineStart: 8, LineEnd: 12},
				{Title: "Save state on the stack", Explanation: "A '(' starts a group. The outer total and the sign in front of the group are pushed, and the group starts again from zero with a plus sign. For \"1-(2+3)\": push (1, -1).", CodeSnippet: "stack.append((total, sign))\ntotal = 0\nsign = 1", LineStart: 14, LineEnd: 16},
				{Title: "Pop and combine", Explanation: "A ')' finishes the group, then restores the saved pair. The group is added with the sign saved in front of it: 1 + (-1) * 5 = -4.", CodeSnippet: "total += sign * number\nnumber = 0\nsaved_total, saved_sign = stack.pop()\ntotal = saved_total + saved_sign * total", LineStart: 18, LineEnd: 21},
				{Title: "Handle what is left", Explanation: "The last number has no operator after it, so it is folded in after the loop.", CodeSnippet: "return total + sign * number", LineStart: 22, LineEnd: 22},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the sign in front of '(' pushed on the stack along with the total?", Answer: "The sign applies to the entire group, and the group's value is only known when the matching ')' arrives. By then the loop has moved on and the sign is gone, so it has to be saved. Pushing it with the outer total lets one step combine them: saved total + saved sign * group total."},
			{Prompt: "How would you handle this problem recursively, and what does the stack replace?", Answer: "Write a function that evaluates from the current position until it meets ')' or the end, and call it on '('; the call returns the group's value to the caller. The saved total and sign are local variables of the caller, which the language keeps on the call stack. The explicit stack stores the same data, which avoids a recursion-depth limit on deeply nested input."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, StackDrills...)
}

// stackDrillDescendingThen returns n, n-1, ..., 1 followed by last.
func stackDrillDescendingThen(n, last int) []int {
	out := make([]int, 0, n+1)
	for v := n; v >= 1; v-- {
		out = append(out, v)
	}
	return append(out, last)
}

// stackDrillAllThenLast returns length-1 copies of v followed by a single -1.
func stackDrillAllThenLast(length, v int) []int {
	out := make([]int, length)
	for i := range out {
		out[i] = v
	}
	out[length-1] = -1
	return out
}
