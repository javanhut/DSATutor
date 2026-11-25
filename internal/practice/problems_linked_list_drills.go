package practice

// LinkedListDrills are the linked-list foundations: build a list by hand, then
// run each core walk and pointer edit on its own, before the linked-list
// ladder applies them. They follow docs/pedagogy.md like any problem, minus
// the brute force (a drill is the plain algorithm itself).
var LinkedListDrills = []*Problem{
	{
		ID:              "drill-build-linked-list",
		Number:          0,
		Kind:            KindDrill,
		Order:           1,
		Title:           "Build a Linked List by Hand",
		Slug:            "drill-build-linked-list",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Foundations"},
		RelatedChapters: []int{7},
		Description: `Every linked list problem hands you a ready-made ` + "`" + `head` + "`" + `. This drill makes you build one yourself.

1. Finish the ` + "`" + `ListNode` + "`" + ` class: a node stores its value in ` + "`" + `val` + "`" + ` and a link to the node after it in ` + "`" + `next` + "`" + ` (None when there is no node after it).
2. Write ` + "`" + `fromArray(values)` + "`" + `, which creates one node per entry of ` + "`" + `values` + "`" + `, in order, links them together, and returns the **head** (the first node).

Return None when ` + "`" + `values` + "`" + ` is empty. The judge walks your list from the head and reads the values back, so return real ` + "`" + `ListNode` + "`" + ` objects, not a list.`,
		Constraints: []string{
			"0 <= len(values) <= 1000",
			"-1000 <= values[i] <= 1000",
		},
		Examples: []Example{
			{Input: "values = [1,2,3,4,5]", Output: "[1,2,3,4,5]", Explanation: "Five nodes, each pointing at the next, the last one at None. The head is the node holding 1."},
			{Input: "values = []", Output: "[]", Explanation: "No values, no nodes: the head is None."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"values": []int{1, 2, 3, 4, 5}}, Expected: []int{1, 2, 3, 4, 5}, Note: "example 1: a list of five nodes"},
			{Input: map[string]interface{}{"values": []int{7}}, Expected: []int{7}, Note: "example 2: a single node; its next must stay None"},
			{Input: map[string]interface{}{"values": []int{}}, Expected: []int{}, Hidden: true, Note: "empty list builds the empty list (None)"},
			{Input: map[string]interface{}{"values": []int{0}}, Expected: []int{0}, Hidden: true, Note: "a value of 0 is a real node (`if v:` or `if not v:` treats it as missing)"},
			{Input: map[string]interface{}{"values": []int{3, 0, 0, 2}}, Expected: []int{3, 0, 0, 2}, Hidden: true, Note: "zeros in the middle and repeated values must all become nodes"},
			{Input: map[string]interface{}{"values": []int{-1, -5, 4, -2}}, Expected: []int{-1, -5, 4, -2}, Hidden: true, Note: "negative values are stored as given"},
			{Input: map[string]interface{}{"values": []int{1, 2}}, Expected: []int{1, 2}, Hidden: true, Note: "two nodes: the first must point at the second, and the second at None"},
			{Input: map[string]interface{}{"values": []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}}, Expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, Hidden: true, Note: "a longer list keeps the order of the input"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `class ListNode:
    def __init__(self, val=0, next=None):
        # Store the value and the link to the next node
        pass


def fromArray(values):
    # Return the head ListNode, or None for an empty list
    pass`,
		StarterCodeOxigen: "// ListNode(val, next) is provided; ListNode(5) leaves next as None.\n\nfun fromArray(values) {\n    // Write your solution here\n}\n",
		Harness: Harness{
			ReturnType: "ListNode",
			Validator: `def check(raw, args, result, expected):
    if isinstance(result_obj, (list, tuple, dict, str, int, float)):
        return (False, "return the head ListNode, not a " + type(result_obj).__name__)
    if result_obj is not None:
        result_obj.val, result_obj.next
    if result != expected:
        return (False, "your list reads back from the head as " + str(result))
    return True`,
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You can only attach a node to a node you already hold. While you build the list, which single node do you always need a reference to, so the next new node has somewhere to hang?"},
			{Level: 2, Type: "insight", Content: "Keep a tail pointer to the last node built so far. For each value, create a node, hang it on tail.next, and move tail onto it. Starting from a throwaway dummy node means the first value needs no special case, and dummy.next is the head at the end. This is the dummy-head idea from Chapter 7. Test with the loop itself, not truthiness: 0 is a real value.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: "1. Define the node: store val and next in __init__.\n2. Make a dummy node and let tail point at it.\n3. For each value: create a node, set tail.next to it, move tail forward.\n4. Return dummy.next, which is None if the loop never ran."},
			{Level: 4, Type: "skeleton", Content: "class ListNode:\n    def __init__(self, val=0, next=None):\n        self.val = ____\n        self.next = ____\n\n\ndef fromArray(values):\n    dummy = ListNode()\n    tail = dummy\n    for v in values:\n        tail.next = ____\n        tail = ____\n    return dummy.next"},
		},
		Solution: Solution{
			Code: `class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def fromArray(values):
    dummy = ListNode()
    tail = dummy
    for v in values:
        tail.next = ListNode(v)
        tail = tail.next
    return dummy.next`,
			Explanation: "A linked list is nodes holding links to other nodes, so the class stores a value and one link. Building one is the dummy-head pattern from Chapter 7: a tail pointer follows the end of the list, and each new node is hung on it.",
			KeyInsight:  "The head of the list is only known once the first node exists, which makes the first value a special case. A dummy node placed before the head removes it: every node, including the first, is attached with tail.next = new node, and the real head is always dummy.next.",
			WhyItWorks:  "Invariant: before each iteration, the chain from dummy.next to tail holds exactly the values read so far, in order, and tail.next is None. It holds at the start (an empty chain, tail is the dummy). Each iteration hangs one new node on tail and moves tail onto it, so the chain gains exactly the next value at its end. After the last value the chain is the whole input, and dummy.next is its head, or None when no node was added. One pass, one node per value: O(n) time and O(n) space for the nodes.",
			CommonMistakes: []string{
				"Testing `if v:` instead of looping over every entry, which drops each node whose value is 0. Every entry becomes a node, whatever its value.",
				"Returning `values`, or a Python list of nodes, instead of the head node. The function hands back one ListNode, and the rest of the list hangs off it.",
				"Forgetting to move tail forward, so every new node overwrites tail.next and the list ends up with only the last value.",
				"Returning `dummy` instead of `dummy.next`, which puts a stray 0 node at the front of the list.",
			},
			TimeComplexity:  "O(n): one node is created per value",
			SpaceComplexity: "O(n) for the nodes themselves",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the node", Explanation: "A node is two fields: its value and the link to the next node. The default next=None makes ListNode(5) a lone last node.", CodeSnippet: "self.val = val\nself.next = next", LineStart: 3, LineEnd: 4},
				{Title: "Start from a dummy node", Explanation: "dummy is a placeholder before the head. tail starts on it, so the first real node is attached the same way as every other.", CodeSnippet: "dummy = ListNode()\ntail = dummy", LineStart: 8, LineEnd: 9},
				{Title: "Hang each value on the tail", Explanation: "On [3,0,2]: node 3 hangs on dummy, tail moves to it; node 0 hangs on 3 (a real node, even though 0 is falsy); node 2 hangs on 0.", CodeSnippet: "for v in values:\n    tail.next = ListNode(v)\n    tail = tail.next", LineStart: 10, LineEnd: 12},
				{Title: "Return the real head", Explanation: "The head is the node after the dummy. With no values the loop never ran, so dummy.next is still None: the empty list.", CodeSnippet: "return dummy.next", LineStart: 13, LineEnd: 13},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why start with a dummy node instead of making the first node separately?", Answer: "Without it, the first value has to be handled differently: there is no tail to hang it on yet, so you need a branch to set the head. A dummy gives every value, including the first, a tail to attach to, and the head is simply dummy.next. The same trick removes the special case at the head in insert and delete."},
			{Prompt: "What does the list look like if you forget `tail = tail.next`?", Answer: "tail never leaves the dummy, so each new node replaces dummy.next. Only the last value survives, and the result is a one-node list holding the final entry."},
		},
	},
	{
		ID:              "drill-list-to-array",
		Number:          0,
		Kind:            KindDrill,
		Order:           2,
		Title:           "Linked List to Array",
		Slug:            "drill-list-to-array",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Traversal", "Foundations"},
		RelatedChapters: []int{7},
		Description: `You are handed the head of a singly linked list. Return its values, in order from head to tail, as an ordinary Python list.

Walking a list is the move every other linked list operation is built on: start at the head, do something with the node you are on, step to its ` + "`" + `next` + "`" + `, and stop when you step off the end (None). An empty list (head is None) gives ` + "`" + `[]` + "`" + `.`,
		Constraints: []string{
			"The list has between 0 and 1000 nodes",
			"-1000 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4,5]", Output: "[1,2,3,4,5]", Explanation: "Visit the nodes from the head: 1, then 2, and so on until next is None."},
			{Input: "head = []", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}}, Expected: []int{1, 2, 3, 4, 5}, Note: "example 1: walk all five nodes"},
			{Input: map[string]interface{}{"head": []int{8}}, Expected: []int{8}, Note: "example 2: a single node"},
			{Input: map[string]interface{}{"head": []int{}}, Expected: []int{}, Hidden: true, Note: "empty list gives [] (the loop must not touch None.val)"},
			{Input: map[string]interface{}{"head": []int{0, 0, 3}}, Expected: []int{0, 0, 3}, Hidden: true, Note: "zeros are values, not the end of the list (`while node.val` stops early)"},
			{Input: map[string]interface{}{"head": []int{4, -2, 4, -2}}, Expected: []int{4, -2, 4, -2}, Hidden: true, Note: "repeated and negative values keep their order"},
			{Input: map[string]interface{}{"head": []int{1, 2}}, Expected: []int{1, 2}, Hidden: true, Note: "the last node must be included (`while node.next` stops one short)"},
			{Input: map[string]interface{}{"head": []int{5, 4, 3, 2, 1, 0}}, Expected: []int{5, 4, 3, 2, 1, 0}, Hidden: true, Note: "a longer list comes out in head-to-tail order"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `# ListNode(val, next) is provided
def toArray(head):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"head": "ListNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You can only reach a node through the one before it, so you cannot jump to the end. What single variable do you need to keep moving along the list, and what value tells you that you have walked off the end?"},
			{Level: 2, Type: "insight", Content: "Keep a pointer called node, starting at head. While node is not None, record node.val and move node = node.next. The loop condition is on the node itself, not node.next, because the last node must still be visited. See the traversal in Chapter 7.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: "1. Make the answer list.\n2. Set node to head.\n3. While node is not None: append node.val, then move node to node.next.\n4. Return the answer list."},
			{Level: 4, Type: "skeleton", Content: "def toArray(head):\n    out = []\n    node = head\n    while ____:\n        out.append(____)\n        node = ____\n    return out"},
		},
		Solution: Solution{
			Code: `def toArray(head):
    out = []
    node = head
    while node is not None:
        out.append(node.val)
        node = node.next
    return out`,
			Explanation: "The basic walk. A pointer starts on the head and steps along next links, recording each value, until it falls off the end of the list. Every later list algorithm is this loop with something extra done at each node.",
			KeyInsight:  "Loop while the node itself exists, not while it has a next. The loop body runs once per real node, including the last one, and the empty list simply never enters the loop.",
			WhyItWorks:  "Invariant: at the top of each iteration, node is the first node whose value is not yet in out, and out holds the values of all nodes before it, in order. It holds at the start (node is the head, out is empty). Each iteration appends node's value and moves to the next node, which keeps the invariant. The loop ends when node is None, meaning every node has been recorded. One step per node: O(n) time, and the output list is O(n) space.",
			CommonMistakes: []string{
				"Looping `while node.next`, which stops on the last node without recording it, and crashes with an AttributeError on an empty list because None has no next.",
				"Looping `while node.val` (or `while node`-style truthiness on the value), which stops at the first 0. Test the node against None.",
				"Advancing with `head = head.next` and also reading head, then losing the start of the list. Use a separate pointer so head stays where it was.",
				"Forgetting to step: without `node = node.next` the loop never ends.",
			},
			TimeComplexity:  "O(n): each node is visited once",
			SpaceComplexity: "O(n) for the output list",
			Walkthrough: []WalkthroughStep{
				{Title: "Make the answer and the pointer", Explanation: "out collects the values; node is the walker and starts on the head.", CodeSnippet: "out = []\nnode = head", LineStart: 2, LineEnd: 3},
				{Title: "Stop only when you fall off", Explanation: "Testing node itself means the last node (whose next is None) is still visited, and an empty list skips the loop.", CodeSnippet: "while node is not None:", LineStart: 4, LineEnd: 4},
				{Title: "Record, then step", Explanation: "On [4,5]: node is 4, record 4, step to 5; record 5, step to None; the loop ends with out = [4,5].", CodeSnippet: "out.append(node.val)\nnode = node.next", LineStart: 5, LineEnd: 6},
				{Title: "Return the values", Explanation: "Every node was visited exactly once, so out is the whole list.", CodeSnippet: "return out", LineStart: 7, LineEnd: 7},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the loop condition `node is not None` and not `node.next is not None`?", Answer: "The condition decides whether there is a node to process. `node.next is not None` asks whether there is a node after this one, so it exits while standing on the last node and never records it, and it crashes on an empty list. `node is not None` processes the last node and then exits when stepping past it."},
			{Prompt: "How would you change the loop to stop one node early, so the walker ends on the last node?", Answer: "Loop while node.next is not None. The walker then ends standing on the last node instead of past it, which is what you want when you need the tail itself (to append there) rather than the values. Check the list is non-empty first."},
		},
	},
	{
		ID:              "drill-list-index-of",
		Number:          0,
		Kind:            KindDrill,
		Order:           3,
		Title:           "Index of a Value in a Linked List",
		Slug:            "drill-list-index-of",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Traversal", "Foundations"},
		RelatedChapters: []int{7},
		Description: `You are handed the head of a singly linked list and a ` + "`" + `target` + "`" + ` value. Return the **0-based position** of the first node whose value equals ` + "`" + `target` + "`" + `, counting from the head, or ` + "`" + `-1` + "`" + ` if no node holds it.

A list has no index to jump to, so you have to walk it and count as you go. If the value appears more than once, the first (closest to the head) match wins. An empty list returns -1.`,
		Constraints: []string{
			"The list has between 0 and 1000 nodes",
			"-1000 <= Node.val, target <= 1000",
		},
		Examples: []Example{
			{Input: "head = [10,20,30], target = 20", Output: "1", Explanation: "Node 10 is at index 0 and node 20 at index 1."},
			{Input: "head = [7,3,7,3], target = 7", Output: "0", Explanation: "7 appears at index 0 and index 2; the first one counts."},
			{Input: "head = [10,20,30], target = 99", Output: "-1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{10, 20, 30}, "target": 20}, Expected: 1, Note: "example 1: found in the middle"},
			{Input: map[string]interface{}{"head": []int{10, 20, 30}, "target": 99}, Expected: -1, Note: "example 2: absent, so -1"},
			{Input: map[string]interface{}{"head": []int{}, "target": 5}, Expected: -1, Hidden: true, Note: "empty list: nothing to find"},
			{Input: map[string]interface{}{"head": []int{4, 5, 6}, "target": 4}, Expected: 0, Hidden: true, Note: "found at the head is index 0"},
			{Input: map[string]interface{}{"head": []int{4, 5, 6}, "target": 6}, Expected: 2, Hidden: true, Note: "found at the last node (the loop must reach it)"},
			{Input: map[string]interface{}{"head": []int{7, 3, 7, 3}, "target": 7}, Expected: 0, Hidden: true, Note: "duplicates: return the first match, not the last"},
			{Input: map[string]interface{}{"head": []int{2, 9, 2, 9, 2}, "target": 2}, Expected: 0, Hidden: true, Note: "first of three matches"},
			{Input: map[string]interface{}{"head": []int{0, 1, 0}, "target": 0}, Expected: 0, Hidden: true, Note: "a target of 0 is a value like any other"},
			{Input: map[string]interface{}{"head": []int{-3, -2, -1}, "target": -2}, Expected: 1, Hidden: true, Note: "negative values compare normally"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode: `# ListNode(val, next) is provided
def indexOf(head, target):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"head": "ListNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "An array knows each element's position, but a node does not. What do you have to carry along while you walk, besides the node you are standing on?"},
			{Level: 2, Type: "insight", Content: "Walk the list as in the previous drill, with a counter that starts at 0 and goes up by one each time you step to the next node. The counter is the index of the node you are on. Return it the moment a node's value equals the target; falling off the end means -1.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: "1. Start node at head and index at 0.\n2. While node is not None: if node.val equals target, return index.\n3. Otherwise step to node.next and add 1 to index.\n4. If the loop ends, the value is not in the list: return -1."},
			{Level: 4, Type: "skeleton", Content: "def indexOf(head, target):\n    node = head\n    index = 0\n    while node is not None:\n        if ____ == target:\n            return ____\n        node = ____\n        index += 1\n    return ____"},
		},
		Solution: Solution{
			Code: `def indexOf(head, target):
    node = head
    index = 0
    while node is not None:
        if node.val == target:
            return index
        node = node.next
        index += 1
    return -1`,
			Explanation: "The walk with a counter. The counter turns the position along the chain into an index, and returning on the first match makes the answer the earliest occurrence.",
			KeyInsight:  "Return inside the loop, the moment a match is seen. Because the walk goes from the head, the first match found is the match closest to the head, so duplicates need no extra handling, and the -1 after the loop is only reached when nothing matched.",
			WhyItWorks:  "Invariant: at the top of each iteration, node is the node at position index, and no node before it holds target. It holds at the start (the head is at index 0). If node matches, index is its position and everything before it did not match, so it is the first match. Otherwise stepping to the next node and adding 1 keeps both facts. If the loop ends, every node was checked and none matched. At most one step per node: O(n) time, O(1) extra space.",
			CommonMistakes: []string{
				"Updating the answer instead of returning on a match, so the last matching index wins. Return immediately at the first match.",
				"Forgetting to increase the counter on a step (or increasing it only on mismatches in the wrong place), so the returned index is off.",
				"Returning the node or its value instead of the index. The question asks for a position.",
				"Using `while node.next`, which never examines the last node, so a target in the last position is reported as -1.",
			},
			TimeComplexity:  "O(n): at most one step per node",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Start at the head with a counter", Explanation: "node is the walker, and index is its position: the head is position 0.", CodeSnippet: "node = head\nindex = 0", LineStart: 2, LineEnd: 3},
				{Title: "Check each node in order", Explanation: "On [7,3,7] with target 7, the very first node matches and 0 is returned, even though another 7 follows.", CodeSnippet: "while node is not None:\n    if node.val == target:\n        return index", LineStart: 4, LineEnd: 6},
				{Title: "Step and count together", Explanation: "Moving to the next node always goes with adding 1 to the index, so the two never disagree.", CodeSnippet: "node = node.next\nindex += 1", LineStart: 7, LineEnd: 8},
				{Title: "Fell off the end", Explanation: "Every node was checked and none matched, so the value is not in the list.", CodeSnippet: "return -1", LineStart: 9, LineEnd: 9},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can the function return from inside the loop?", Answer: "Nothing after the first match can change the answer, because the first match is by definition the one closest to the head. Returning at once is correct and also stops walking early, so a hit near the front costs only a few steps."},
			{Prompt: "What is the worst case for this function, and which inputs cause it?", Answer: "The target sits in the last node or is absent, so every node is visited: n steps. There is no faster way on a plain linked list, because without an index the only route to a node is through all the nodes before it."},
		},
	},
	{
		ID:              "drill-middle-of-list",
		Number:          876,
		Kind:            KindDrill,
		Order:           4,
		Title:           "Middle of the Linked List",
		Slug:            "middle-of-the-linked-list",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Two Pointers", "Fast and Slow", "Foundations"},
		RelatedChapters: []int{7},
		Description: `You are handed the head of a singly linked list. Return the **middle node**. When the list has an even number of nodes there are two middle nodes; return the **second** one.

You return the node itself, so the answer is the rest of the list from that node to the end (the judge shows it as a list of values). An empty list returns None.

Do it without first counting the nodes if you can: one pass, using two pointers that move at different speeds.`,
		Constraints: []string{
			"The list has between 0 and 100 nodes",
			"1 <= Node.val <= 100 (the tests also use 0 and negatives)",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4,5]", Output: "[3,4,5]", Explanation: "Five nodes: the middle is the third, and the list from there is 3,4,5."},
			{Input: "head = [1,2,3,4,5,6]", Output: "[4,5,6]", Explanation: "Six nodes: the middles are 3 and 4, and the second one, 4, is returned."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}}, Expected: []int{3, 4, 5}, Note: "example 1: odd length, the exact middle"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5, 6}}, Expected: []int{4, 5, 6}, Note: "example 2: even length, the second of the two middles"},
			{Input: map[string]interface{}{"head": []int{}}, Expected: []int{}, Hidden: true, Note: "empty list returns None, which reads as []"},
			{Input: map[string]interface{}{"head": []int{7}}, Expected: []int{7}, Hidden: true, Note: "a single node is its own middle"},
			{Input: map[string]interface{}{"head": []int{1, 2}}, Expected: []int{2}, Hidden: true, Note: "two nodes: the second one (the first middle is wrong)"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3}}, Expected: []int{2, 3}, Hidden: true, Note: "three nodes: the node in the center"},
			{Input: map[string]interface{}{"head": []int{9, 8, 7, 6}}, Expected: []int{7, 6}, Hidden: true, Note: "four nodes: the third node"},
			{Input: map[string]interface{}{"head": []int{0, -1, 0, -1, 0, -1, 0, -1}}, Expected: []int{0, -1, 0, -1}, Hidden: true, Note: "repeated values: the middle is a node, not a value (even length again)"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode: `# ListNode(val, next) is provided
def middleNode(head):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"head": "ListNode"},
			ReturnType: "ListNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If one walker covers the whole list in n steps and a second walker moves at half that speed, where is the second walker when the first one finishes?"},
			{Level: 2, Type: "insight", Content: "Use a fast and a slow pointer. Both start at the head; slow moves one node per round and fast moves two. When fast runs off the end, slow has covered half the distance, so it stands on the middle. This is the fast/slow pointer pattern from Chapter 7: no counting pass and no length needed.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: "1. Set slow and fast to head.\n2. Loop while fast exists and fast.next exists (so that fast can safely move two steps).\n3. Each round: slow moves one node, fast moves two.\n4. When the loop ends, return slow."},
			{Level: 4, Type: "skeleton", Content: "def middleNode(head):\n    slow = head\n    fast = head\n    while ____ and ____:\n        slow = ____\n        fast = ____\n    return slow"},
		},
		Solution: Solution{
			Code: `def middleNode(head):
    slow = head
    fast = head
    while fast is not None and fast.next is not None:
        slow = slow.next
        fast = fast.next.next
    return slow`,
			Explanation: "Fast and slow pointers. Because fast moves twice as fast as slow, slow is halfway along the list when fast reaches the end, so a single pass finds the middle without knowing the length.",
			KeyInsight:  "Two pointers at different speeds turn a distance into a position: when the faster one has gone the whole length, the slower one has gone half of it. The loop condition fast and fast.next is what makes an even-length list stop on the second middle.",
			WhyItWorks:  "After k rounds slow is at position k and fast at position 2k (if it has not run off the end). The loop continues while fast has a node to stand on and one more after it, so it runs until 2k reaches n-1 (odd n, fast on the last node) or n (even n, fast just past the end). For odd n that happens at k = (n-1)/2, the exact middle. For even n it happens at k = n/2, which is the second of the two middles. An empty list never enters the loop and returns None. Each round advances fast two nodes, so about n/2 rounds: O(n) time, O(1) space.",
			CommonMistakes: []string{
				"Looping only while fast.next exists, which crashes on an empty list, or only while fast exists, which crashes when fast.next is None and fast.next.next is evaluated. Check both fast and fast.next.",
				"Using `while fast.next and fast.next.next`, which stops one step earlier and returns the first middle on an even-length list. The problem asks for the second.",
				"Counting the nodes first and walking n // 2 steps. It works, but needs two passes; the point of the drill is the single pass with two speeds.",
				"Returning slow.val instead of slow. The answer is the node, and the rest of the list hangs off it.",
			},
			TimeComplexity:  "O(n): fast takes about n/2 rounds",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Start both pointers at the head", Explanation: "They begin together; only their speeds differ.", CodeSnippet: "slow = head\nfast = head", LineStart: 2, LineEnd: 3},
				{Title: "Run while fast can take two steps", Explanation: "Fast needs a node and a node after it, otherwise moving two steps would step through None.", CodeSnippet: "while fast is not None and fast.next is not None:", LineStart: 4, LineEnd: 4},
				{Title: "Slow takes one, fast takes two", Explanation: "On [1,2,3,4]: round 1 puts slow on 2 and fast on 3; round 2 puts slow on 3 and fast past the end, and the loop stops.", CodeSnippet: "slow = slow.next\nfast = fast.next.next", LineStart: 5, LineEnd: 6},
				{Title: "Slow is on the middle", Explanation: "Fast has covered the list, so slow has covered half: it is on the middle, the second one when the length is even.", CodeSnippet: "return slow", LineStart: 7, LineEnd: 7},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the condition need both `fast is not None` and `fast.next is not None`?", Answer: "Fast moves two nodes with fast.next.next. Before that move, fast.next must exist (otherwise None.next fails), and to ask for fast.next at all, fast must exist (an empty list, or fast already past the end on even lengths). Python's `and` checks them left to right, so the second check is only reached when the first passed."},
			{Prompt: "What one change would make an even-length list return the first middle instead?", Answer: "Stop one step earlier: loop while fast.next and fast.next.next. Then on [1,2,3,4] slow stops on 2 rather than 3. Which middle you need decides, for example, where to cut a list in half when splitting it for merge sort."},
		},
	},
	{
		ID:              "drill-insert-at-index",
		Number:          0,
		Kind:            KindDrill,
		Order:           5,
		Title:           "Insert at an Index",
		Slug:            "drill-insert-at-index",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Dummy Node", "Pointer Rewiring", "Foundations"},
		RelatedChapters: []int{7},
		Description: `You are handed the head of a singly linked list, an ` + "`" + `index` + "`" + ` and a value ` + "`" + `val` + "`" + `. Insert a new node holding ` + "`" + `val` + "`" + ` so that it ends up **at position ` + "`" + `index` + "`" + `** (0-based) in the list, and return the head of the resulting list.

` + "`" + `index = 0` + "`" + ` puts the new node at the front, so the head changes. ` + "`" + `index` + "`" + ` equal to the current length puts it at the end. The index is always between 0 and the length, so you never need to handle an index past the end. The list is changed in place by relinking nodes; you only create the one new node.`,
		Constraints: []string{
			"The list has between 0 and 1000 nodes",
			"0 <= index <= the number of nodes in the list",
			"-1000 <= Node.val, val <= 1000",
		},
		Examples: []Example{
			{Input: "head = [1,2,4], index = 2, val = 3", Output: "[1,2,3,4]", Explanation: "The new node must sit at position 2, between 2 and 4."},
			{Input: "head = [1,2,3], index = 0, val = 0", Output: "[0,1,2,3]", Explanation: "A new head: the old head is now at position 1."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 4}, "index": 2, "val": 3}, Expected: []int{1, 2, 3, 4}, Note: "example 1: insert in the middle"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3}, "index": 0, "val": 0}, Expected: []int{0, 1, 2, 3}, Note: "example 2: index 0 makes a new head (needs a dummy node)"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3}, "index": 3, "val": 4}, Expected: []int{1, 2, 3, 4}, Hidden: true, Note: "index equal to the length appends at the end"},
			{Input: map[string]interface{}{"head": []int{}, "index": 0, "val": 9}, Expected: []int{9}, Hidden: true, Note: "empty list: the new node becomes the whole list"},
			{Input: map[string]interface{}{"head": []int{5}, "index": 1, "val": 6}, Expected: []int{5, 6}, Hidden: true, Note: "append to a one-node list"},
			{Input: map[string]interface{}{"head": []int{5}, "index": 0, "val": 6}, Expected: []int{6, 5}, Hidden: true, Note: "prepend to a one-node list"},
			{Input: map[string]interface{}{"head": []int{1, 3}, "index": 1, "val": 2}, Expected: []int{1, 2, 3}, Hidden: true, Note: "between two nodes"},
			{Input: map[string]interface{}{"head": []int{-1, -2}, "index": 1, "val": -5}, Expected: []int{-1, -5, -2}, Hidden: true, Note: "negative values are just values"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode: `# ListNode(val, next) is provided
def insertAt(head, index, val):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"head": "ListNode"},
			ReturnType: "ListNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "To put a node at position index you have to change the next pointer of the node at index - 1. What do you do when index is 0 and there is no node before the position?"},
			{Level: 2, Type: "insight", Content: "Put a dummy node in front of the head, so there is always a node before position index, even for index 0. Walk a pointer prev forward index times from the dummy; now prev.next is the node that should end up after the new one. Then the insert is two pointer changes: the new node's next takes prev.next, and prev.next becomes the new node. The head is dummy.next. See the dummy node in Chapter 7.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: "1. Create dummy whose next is head, and set prev to dummy.\n2. Move prev forward index times, so prev is the node just before the insertion point.\n3. Create the new node with next set to prev.next.\n4. Set prev.next to the new node.\n5. Return dummy.next."},
			{Level: 4, Type: "skeleton", Content: "def insertAt(head, index, val):\n    dummy = ListNode(0, head)\n    prev = dummy\n    for _ in range(____):\n        prev = prev.next\n    prev.next = ListNode(val, ____)\n    return ____"},
		},
		Solution: Solution{
			Code: `def insertAt(head, index, val):
    dummy = ListNode(0, head)
    prev = dummy
    for _ in range(index):
        prev = prev.next
    prev.next = ListNode(val, prev.next)
    return dummy.next`,
			Explanation: "Insertion is pointer rewiring at one place. A dummy node in front of the head guarantees a node exists before every position, so inserting at the front is no different from inserting anywhere else.",
			KeyInsight:  "To insert you must hold the node before the gap. The dummy node makes one exist for every index from 0 to the length. Order matters: the new node must take the old prev.next first, and only then may prev.next be redirected, or the rest of the list is lost.",
			WhyItWorks:  "Starting at the dummy and moving prev index times puts prev at position index - 1 of the real list (the dummy is position -1), that is, the node directly before the gap. The new node is created already pointing at prev.next, the node currently at position index (or None at the end), and then prev.next is pointed at the new node. Everything before prev is unchanged, the new node is at position index, and everything that was at position index or later follows it. For index 0, prev is the dummy, so dummy.next, the head, becomes the new node. The walk takes index steps: O(n) time, O(1) space besides the new node.",
			CommonMistakes: []string{
				"Walking to the node at position index - 1 from the head with no dummy. For index 0 there is no such node, so the code crashes or inserts after the head instead of before it.",
				"Doing prev.next = new first and then new.next = prev.next, which makes the new node point at itself and loses the rest of the list. Set the new node's next first.",
				"Walking one step too far or too few, so the value lands one position off. prev must stop on the node before the gap.",
				"Returning head instead of dummy.next. When index is 0 the head has changed.",
			},
			TimeComplexity:  "O(n): the walk takes at most n steps",
			SpaceComplexity: "O(1) besides the new node",
			Walkthrough: []WalkthroughStep{
				{Title: "Put a dummy in front of the head", Explanation: "dummy.next is the head, and prev starts on the dummy: position -1, the node before position 0.", CodeSnippet: "dummy = ListNode(0, head)\nprev = dummy", LineStart: 2, LineEnd: 3},
				{Title: "Walk to the node before the gap", Explanation: "index steps from the dummy land prev on position index - 1. For index 0 it does not move at all.", CodeSnippet: "for _ in range(index):\n    prev = prev.next", LineStart: 4, LineEnd: 5},
				{Title: "Splice in the new node", Explanation: "On [1,2,4] with index 2: prev is node 2. The new node 3 takes next = 4 first, then 2 is pointed at 3, giving 1,2,3,4.", CodeSnippet: "prev.next = ListNode(val, prev.next)", LineStart: 6, LineEnd: 6},
				{Title: "Return the head", Explanation: "The head is dummy.next: still the old head, or the new node when index was 0.", CodeSnippet: "return dummy.next", LineStart: 7, LineEnd: 7},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the new node's next be set before prev.next is changed?", Answer: "prev.next is the only link to the rest of the list from this point. If you overwrite it first, nothing refers to the old tail any more, and the nodes after the gap are lost. Reading prev.next into the new node first keeps them attached, and then redirecting prev.next completes the splice."},
			{Prompt: "How does the dummy node help when the list is empty?", Answer: "With an empty list, head is None, so there is no node to start from. The dummy is a real node whose next is None; with index 0, prev stays on it, the new node is attached to dummy.next, and dummy.next is returned as the new head. No empty-list branch is needed."},
		},
	},
	{
		ID:              "drill-remove-elements",
		Number:          203,
		Kind:            KindDrill,
		Order:           6,
		Title:           "Remove Linked List Elements",
		Slug:            "remove-linked-list-elements",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Dummy Node", "Pointer Rewiring", "Foundations"},
		RelatedChapters: []int{7},
		Description: `You are handed the head of a singly linked list and a value ` + "`" + `val` + "`" + `. Remove **every** node whose value equals ` + "`" + `val` + "`" + ` and return the head of the resulting list.

The matching nodes may be at the head, at the tail, or several in a row, and they may be all of the list. Remove nodes by relinking the ones around them; do not build a new list. An empty result is None, which the judge shows as ` + "`" + `[]` + "`" + `.`,
		Constraints: []string{
			"The list has between 0 and 10000 nodes",
			"0 <= Node.val, val <= 50 (the tests also use negatives)",
		},
		Examples: []Example{
			{Input: "head = [1,2,6,3,4,5,6], val = 6", Output: "[1,2,3,4,5]", Explanation: "Both nodes holding 6 are unlinked; the others keep their order."},
			{Input: "head = [7,7,7,7], val = 7", Output: "[]", Explanation: "Every node matches, so the head moves right off the end of the list."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 6, 3, 4, 5, 6}, "val": 6}, Expected: []int{1, 2, 3, 4, 5}, Note: "example 1: a match in the middle and one at the tail"},
			{Input: map[string]interface{}{"head": []int{}, "val": 1}, Expected: []int{}, Note: "example 2: empty list stays empty"},
			{Input: map[string]interface{}{"head": []int{7, 7, 7, 7}, "val": 7}, Expected: []int{}, Note: "every node matches, so nothing is left (the head changes repeatedly)"},
			{Input: map[string]interface{}{"head": []int{1, 2, 2, 1}, "val": 2}, Expected: []int{1, 1}, Hidden: true, Note: "two matches in a row: the second must not be skipped"},
			{Input: map[string]interface{}{"head": []int{6, 6, 1}, "val": 6}, Expected: []int{1}, Hidden: true, Note: "matches at the head: the answer's head is a later node"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3}, "val": 4}, Expected: []int{1, 2, 3}, Hidden: true, Note: "no match leaves the list as it is"},
			{Input: map[string]interface{}{"head": []int{0, 1, 0}, "val": 0}, Expected: []int{1}, Hidden: true, Note: "removing 0: matches at both ends"},
			{Input: map[string]interface{}{"head": []int{-1, -1, 2, -1}, "val": -1}, Expected: []int{2}, Hidden: true, Note: "a run of matches, then a survivor, then another match"},
			{Input: map[string]interface{}{"head": []int{3, 5, 5, 5, 3}, "val": 5}, Expected: []int{3, 3}, Hidden: true, Note: "a long run in the middle"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode: `# ListNode(val, next) is provided
def removeElements(head, val):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"head": "ListNode"},
			ReturnType: "ListNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "To delete a node you must change the next pointer of the node before it. What does this mean for a match at the head, and what should happen to your 'previous node' after you delete something?"},
			{Level: 2, Type: "insight", Content: "Use a dummy node in front of the head, so the head is deleted by the same rule as any other node. Keep prev, the last node that is staying in the list, and node, the one being examined. A match is skipped with prev.next = node.next, and prev does NOT move, because the node after the deleted one might match too. Only a node that stays moves prev onto itself. See Chapter 7.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: "1. Make dummy pointing at head, with prev on the dummy and node on head.\n2. While node exists: if node.val equals val, set prev.next to node.next (delete it).\n3. Otherwise node stays, so move prev onto it.\n4. Either way, step node to node.next.\n5. Return dummy.next."},
			{Level: 4, Type: "skeleton", Content: "def removeElements(head, val):\n    dummy = ListNode(0, head)\n    prev = dummy\n    node = head\n    while node is not None:\n        if node.val == val:\n            prev.next = ____\n        else:\n            prev = ____\n        node = ____\n    return ____"},
		},
		Solution: Solution{
			Code: `def removeElements(head, val):
    dummy = ListNode(0, head)
    prev = dummy
    node = head
    while node is not None:
        if node.val == val:
            prev.next = node.next
        else:
            prev = node
        node = node.next
    return dummy.next`,
			Explanation: "A single pass with a dummy head and a trailing pointer. prev always points at the last node that is staying in the list, so a deletion is one pointer change on it, and the head is handled by the same code as everything else.",
			KeyInsight:  "prev must only advance onto nodes that are kept. After a deletion prev stays put, so a run of matching nodes is removed one by one by repeatedly redirecting the same prev.next, and the dummy node means the head is just another node that has a prev.",
			WhyItWorks:  "Invariant: the chain from the dummy to prev contains exactly the already-examined nodes that do not match val, in order, and prev.next is node. A matching node is bypassed with prev.next = node.next, which keeps the invariant: the kept chain is unchanged, and the next examined node is again prev.next. A non-matching node is appended to the kept chain by moving prev onto it. When node reaches None every node was examined, so the chain holds exactly the survivors, and dummy.next is the head (None if there are none). One step per node: O(n) time, O(1) space.",
			CommonMistakes: []string{
				"Moving prev forward after a deletion too. Then prev points at the node that was just removed, and the next deletion changes a pointer on a node that is no longer in the list, so a run of consecutive matches is only partly removed.",
				"Handling the head separately with a `while head and head.val == val` loop and forgetting the empty and all-match cases. A dummy node removes that special case.",
				"Comparing the value of node.next in a loop that stands on prev, and then skipping the check of the node after a deletion. Examine the node you are on and step exactly once per round.",
				"Returning head instead of dummy.next, which is wrong whenever the original head was removed.",
			},
			TimeComplexity:  "O(n): one step per node",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the dummy and the two pointers", Explanation: "prev is the last kept node (the dummy at first); node is the one under examination.", CodeSnippet: "dummy = ListNode(0, head)\nprev = dummy\nnode = head", LineStart: 2, LineEnd: 4},
				{Title: "Delete a match by bypassing it", Explanation: "On [6,6,1] with val 6: prev (the dummy) skips the first 6, and stays put, so the second 6 is skipped by the same prev.next change.", CodeSnippet: "if node.val == val:\n    prev.next = node.next", LineStart: 6, LineEnd: 7},
				{Title: "Keep a non-match by moving prev onto it", Explanation: "Only a node that stays in the list becomes the new prev.", CodeSnippet: "else:\n    prev = node", LineStart: 8, LineEnd: 9},
				{Title: "Step and return", Explanation: "node always steps to the next node; the head is whatever now follows the dummy.", CodeSnippet: "node = node.next\nreturn dummy.next", LineStart: 10, LineEnd: 11},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does prev stay where it is after a deletion?", Answer: "prev must be the last node still in the list, because that is the node whose next you rewrite. After deleting the node after prev, the new node after prev has not been checked yet; if it also matches, it must be bypassed by changing the same prev.next. Moving prev onto a deleted node would make later deletions edit a node that is no longer connected."},
			{Prompt: "What goes wrong without a dummy node when the first two nodes match?", Answer: "The head itself has no previous node, so you need a separate loop that advances head while it matches, then a second loop for the rest, and both must cope with the list becoming empty. The dummy gives the head a prev, so one loop covers every case, and the answer is dummy.next."},
		},
	},
	{
		ID:              "drill-design-linked-list",
		Number:          707,
		Kind:            KindDrill,
		Order:           7,
		Title:           "Design Linked List",
		Slug:            "design-linked-list",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Design", "Dummy Node", "Foundations"},
		RelatedChapters: []int{7},
		Description: `Build your own singly linked list class, ` + "`" + `MyLinkedList` + "`" + `, from nodes you define yourself. It supports these methods; indices are 0-based:

- ` + "`" + `get(index)` + "`" + `: return the value of the node at ` + "`" + `index` + "`" + `, or -1 if the index is not valid.
- ` + "`" + `addAtHead(val)` + "`" + `: insert a new node at the front.
- ` + "`" + `addAtTail(val)` + "`" + `: insert a new node at the end.
- ` + "`" + `addAtIndex(index, val)` + "`" + `: insert a new node so that it sits at ` + "`" + `index` + "`" + `. If ` + "`" + `index` + "`" + ` equals the length, it is appended; if ` + "`" + `index` + "`" + ` is greater than the length, nothing happens.
- ` + "`" + `deleteAtIndex(index)` + "`" + `: remove the node at ` + "`" + `index` + "`" + ` if the index is valid; otherwise do nothing.

An index is valid for get and delete when it is at least 0 and less than the number of nodes. This drill puts the earlier ones together: a walk, an insert and a delete, all behind one class.

How the tests call it: operations lists the method names (the first is the constructor) and values the arguments for each; the expected list has null for calls that return nothing.`,
		Constraints: []string{
			"0 <= index, val <= 1000",
			"At most 2000 calls to get, addAtHead, addAtTail, addAtIndex and deleteAtIndex",
		},
		Examples: []Example{
			{Input: "[\"MyLinkedList\",\"addAtHead\",\"addAtTail\",\"addAtIndex\",\"get\",\"deleteAtIndex\",\"get\"]\n[[],[1],[3],[1,2],[1],[1],[1]]", Output: "[null,null,null,null,2,null,3]", Explanation: "The list goes 1, then 1,3, then 1,2,3. get(1) reads 2. After deleting index 1 the list is 1,3, so get(1) reads 3."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"}, "values": [][]int{{}, {1}, {3}, {1, 2}, {1}, {1}, {1}}}, Expected: []interface{}{nil, nil, nil, nil, 2, nil, 3}, Note: "example: build 1,2,3, read the middle, delete it, read again"},
			{Input: map[string]interface{}{"operations": []string{"MyLinkedList", "get", "addAtHead", "get", "get"}, "values": [][]int{{}, {0}, {5}, {0}, {1}}}, Expected: []interface{}{nil, -1, nil, 5, -1}, Note: "get on an empty list and past the end returns -1"},
			{Input: map[string]interface{}{"operations": []string{"MyLinkedList", "addAtIndex", "addAtIndex", "get", "addAtIndex", "get", "get"}, "values": [][]int{{}, {1, 7}, {0, 7}, {0}, {1, 8}, {1}, {2}}}, Expected: []interface{}{nil, nil, nil, 7, nil, 8, -1}, Hidden: true, Note: "addAtIndex past the length does nothing; index equal to the length appends"},
			{Input: map[string]interface{}{"operations": []string{"MyLinkedList", "addAtTail", "addAtTail", "addAtHead", "addAtHead", "get", "get", "get", "get", "get"}, "values": [][]int{{}, {2}, {3}, {1}, {0}, {0}, {1}, {2}, {3}, {4}}}, Expected: []interface{}{nil, nil, nil, nil, nil, 0, 1, 2, 3, -1}, Hidden: true, Note: "head and tail adds on a list that starts empty"},
			{Input: map[string]interface{}{"operations": []string{"MyLinkedList", "addAtHead", "addAtHead", "addAtHead", "deleteAtIndex", "get", "get", "get"}, "values": [][]int{{}, {3}, {2}, {1}, {0}, {0}, {1}, {2}}}, Expected: []interface{}{nil, nil, nil, nil, nil, 2, 3, -1}, Hidden: true, Note: "deleting the head leaves the next node first"},
			{Input: map[string]interface{}{"operations": []string{"MyLinkedList", "addAtTail", "addAtTail", "addAtTail", "deleteAtIndex", "get", "addAtTail", "get", "get"}, "values": [][]int{{}, {1}, {2}, {3}, {2}, {1}, {4}, {2}, {3}}}, Expected: []interface{}{nil, nil, nil, nil, nil, 2, nil, 4, -1}, Hidden: true, Note: "deleting the tail, then appending, must reconnect correctly"},
			{Input: map[string]interface{}{"operations": []string{"MyLinkedList", "addAtHead", "deleteAtIndex", "deleteAtIndex", "get", "addAtTail", "get"}, "values": [][]int{{}, {4}, {1}, {0}, {0}, {6}, {0}}}, Expected: []interface{}{nil, nil, nil, nil, -1, nil, 6}, Hidden: true, Note: "an invalid delete does nothing; deleting the last node leaves an empty list that still works"},
			{Input: map[string]interface{}{"operations": []string{"MyLinkedList", "addAtIndex", "addAtIndex", "addAtIndex", "addAtIndex", "get", "get", "get", "get", "deleteAtIndex", "get", "get"}, "values": [][]int{{}, {0, 1}, {1, 3}, {1, 2}, {3, 4}, {0}, {1}, {2}, {3}, {2}, {2}, {3}}}, Expected: []interface{}{nil, nil, nil, nil, nil, 1, 2, 3, 4, nil, 4, -1}, Hidden: true, Note: "inserts in the middle and at the end, then a middle delete"},
			{Input: map[string]interface{}{"operations": []string{"MyLinkedList", "addAtHead", "deleteAtIndex", "get", "addAtIndex", "get"}, "values": [][]int{{}, {0}, {5}, {0}, {2, 9}, {1}}}, Expected: []interface{}{nil, nil, nil, 0, nil, -1}, Hidden: true, Note: "a value of 0 is a real node; invalid delete and invalid add change nothing"},
		},
		TimeComplexity:  "O(n) per operation, O(1) for addAtHead",
		SpaceComplexity: "O(n)",
		StarterCode: `class MyLinkedList:
    def __init__(self):
        # Set up an empty list (define your own node class)
        pass

    def get(self, index):
        pass

    def addAtHead(self, val):
        pass

    def addAtTail(self, val):
        pass

    def addAtIndex(self, index, val):
        pass

    def deleteAtIndex(self, index):
        pass`,
		Harness: Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You wrote insert-at-index and removal as separate functions earlier. If one helper finds the node just before a position, how many of these five methods can be built on top of it?"},
			{Level: 2, Type: "insight", Content: "Keep a dummy node and a size count. A helper _before(index) walks from the dummy and returns the node before position index. Then get reads _before(index).next.val, addAtIndex splices after it, deleteAtIndex relinks around it, and addAtHead and addAtTail are just addAtIndex(0, ...) and addAtIndex(size, ...). The size is what tells valid indices from invalid ones. See Chapter 7.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: "1. Define a node class, and in __init__ create the dummy and set size to 0.\n2. Write _before(index): start at the dummy and step index times.\n3. get: return -1 unless 0 <= index < size; otherwise the answer is _before(index).next.val.\n4. addAtIndex: do nothing unless 0 <= index <= size; otherwise splice a new node after _before(index) and add 1 to size.\n5. deleteAtIndex: do nothing unless 0 <= index < size; otherwise bypass the node and subtract 1 from size.\n6. addAtHead and addAtTail call addAtIndex with 0 and size."},
			{Level: 4, Type: "skeleton", Content: "class Node:\n    def __init__(self, val=0, next=None):\n        self.val = val\n        self.next = next\n\n\nclass MyLinkedList:\n    def __init__(self):\n        self.dummy = Node()\n        self.size = 0\n\n    def _before(self, index):\n        prev = self.dummy\n        for _ in range(index):\n            prev = prev.next\n        return prev\n\n    def get(self, index):\n        if index < 0 or index >= ____:\n            return -1\n        return self._before(index).next.val\n\n    def addAtHead(self, val):\n        self.addAtIndex(0, val)\n\n    def addAtTail(self, val):\n        self.addAtIndex(____, val)\n\n    def addAtIndex(self, index, val):\n        if index < 0 or index > ____:\n            return\n        prev = self._before(index)\n        prev.next = Node(val, ____)\n        self.size += 1\n\n    def deleteAtIndex(self, index):\n        if index < 0 or index >= self.size:\n            return\n        prev = self._before(index)\n        prev.next = ____\n        self.size -= 1"},
		},
		Solution: Solution{
			Code: `class Node:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class MyLinkedList:
    def __init__(self):
        self.dummy = Node()
        self.size = 0

    def _before(self, index):
        prev = self.dummy
        for _ in range(index):
            prev = prev.next
        return prev

    def get(self, index):
        if index < 0 or index >= self.size:
            return -1
        return self._before(index).next.val

    def addAtHead(self, val):
        self.addAtIndex(0, val)

    def addAtTail(self, val):
        self.addAtIndex(self.size, val)

    def addAtIndex(self, index, val):
        if index < 0 or index > self.size:
            return
        prev = self._before(index)
        prev.next = Node(val, prev.next)
        self.size += 1

    def deleteAtIndex(self, index):
        if index < 0 or index >= self.size:
            return
        prev = self._before(index)
        prev.next = prev.next.next
        self.size -= 1`,
			Explanation: "Everything here is a variation of the same move: find the node before a position, then read or rewire its next. The dummy head removes the edge case at the front, and the size count separates valid indices from invalid ones without walking the list.",
			KeyInsight:  "One helper that returns the node before a position serves every method: reading, inserting and deleting all start there. The two bounds differ by one on purpose: an insert may use index == size (append) but a read or delete may not, because no node exists at that position.",
			WhyItWorks:  "Invariant: the chain after the dummy holds exactly size nodes, in list order. _before(index) returns the node at position index - 1 (the dummy for index 0), valid for 0 <= index <= size. get and delete first require index < size, so _before(index).next is a real node. addAtIndex allows index up to size, where _before(index) is the last node (or the dummy) and the new node is appended after it. Each add or delete changes the chain by one node and size by one, preserving the invariant, and invalid indices return before touching anything. The walk is the only non-constant part: O(n) per operation, except addAtHead, which walks zero steps.",
			CommonMistakes: []string{
				"Rejecting index == size in addAtIndex (using >= instead of >), so appending by index silently does nothing. Index equal to the length is valid for insertion.",
				"Allowing index == size in get or deleteAtIndex, so _before(size).next is None and the code crashes on None.val. Reads and deletes need a node to exist.",
				"Forgetting to update size after an insert or a delete, so every later bounds check is wrong.",
				"Changing the head with a special case and forgetting the empty list: the first add on an empty list, and the delete of the only node, are where the front-of-list code usually breaks. The dummy node removes these cases.",
			},
			TimeComplexity:  "O(n) per operation, O(1) for addAtHead",
			SpaceComplexity: "O(n) for the nodes",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the dummy and the size", Explanation: "The dummy never holds data; size is the number of real nodes, so bounds checks need no walking.", CodeSnippet: "self.dummy = Node()\nself.size = 0", LineStart: 9, LineEnd: 10},
				{Title: "Find the node before a position", Explanation: "Stepping index times from the dummy lands on position index - 1, and on the dummy itself for index 0.", CodeSnippet: "prev = self.dummy\nfor _ in range(index):\n    prev = prev.next\nreturn prev", LineStart: 13, LineEnd: 16},
				{Title: "Read with a strict bound", Explanation: "Valid positions are 0 to size - 1, so anything else is -1. The value is in the node after the one before.", CodeSnippet: "if index < 0 or index >= self.size:\n    return -1\nreturn self._before(index).next.val", LineStart: 19, LineEnd: 21},
				{Title: "Head and tail are special indices", Explanation: "Adding at the head is index 0 and adding at the tail is index size, so both reuse addAtIndex.", CodeSnippet: "self.addAtIndex(0, val)\nself.addAtIndex(self.size, val)", LineStart: 24, LineEnd: 27},
				{Title: "Insert with an inclusive bound", Explanation: "index == size is allowed here, because appending is valid; anything larger is ignored. The new node takes prev's old next first.", CodeSnippet: "if index < 0 or index > self.size:\n    return\nprev = self._before(index)\nprev.next = Node(val, prev.next)\nself.size += 1", LineStart: 30, LineEnd: 34},
				{Title: "Delete by bypassing the node", Explanation: "The same strict bound as get. The node before simply points past the deleted one.", CodeSnippet: "prev.next = prev.next.next\nself.size -= 1", LineStart: 40, LineEnd: 41},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is index == size valid for addAtIndex but not for get?", Answer: "To insert at position size you need only the last node (or the dummy) to hang the new node on; the position just past the end is a legal place for a new node. To read or delete, a node must already exist at that position, and none does at size."},
			{Prompt: "Which of the five methods takes O(1), and why?", Answer: "addAtHead. It splices after the dummy, so _before(0) takes zero steps. addAtTail still costs O(n) here because it walks to the end; keeping a tail pointer would make it O(1), at the price of updating that pointer when the last node is deleted."},
		},
	},
	{
		ID:              "drill-doubly-linked-list",
		Number:          0,
		Kind:            KindDrill,
		Order:           8,
		Title:           "Doubly Linked List Deque",
		Slug:            "drill-doubly-linked-list",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Doubly Linked List", "Design", "Sentinel Nodes", "Foundations"},
		RelatedChapters: []int{7},
		Description: `A **doubly linked list** gives every node two links: ` + "`" + `next` + "`" + ` to the node after it and ` + "`" + `prev` + "`" + ` to the node before it. That lets you remove a node, and reach either end, without walking the list. The LRU Cache later in this ladder is built on one.

Build a double-ended queue, ` + "`" + `Deque` + "`" + `, on a doubly linked list with two **sentinel** nodes: a dummy head before the first real node and a dummy tail after the last, so no real node ever has a missing neighbour. Methods:

- ` + "`" + `pushFront(val)` + "`" + ` and ` + "`" + `pushBack(val)` + "`" + `: add a node at the front or at the back.
- ` + "`" + `popFront()` + "`" + ` and ` + "`" + `popBack()` + "`" + `: remove the node at that end and return its value, or -1 if the deque is empty.
- ` + "`" + `toList()` + "`" + `: the values from front to back, following the ` + "`" + `next` + "`" + ` links.
- ` + "`" + `toListReversed()` + "`" + `: the values from back to front, following the ` + "`" + `prev` + "`" + ` links from the tail.

All five operations must be O(1), so use doubly linked nodes; the tests can only check what comes back, but ` + "`" + `toListReversed` + "`" + ` reads the list through ` + "`" + `prev` + "`" + `, so a deque that lets those links go stale is caught.

How the tests call it: operations lists the method names (the first is the constructor) and values the arguments for each; the expected list has null for calls that return nothing.`,
		Constraints: []string{
			"-1000 <= val <= 1000",
			"At most 1000 calls",
			"popFront and popBack return -1 only on an empty deque",
		},
		Examples: []Example{
			{Input: "[\"Deque\",\"pushBack\",\"pushBack\",\"pushFront\",\"toList\",\"toListReversed\"]\n[[],[2],[3],[1],[],[]]", Output: "[null,null,null,null,[1,2,3],[3,2,1]]", Explanation: "pushBack(2), pushBack(3) and pushFront(1) build 1,2,3. Reading from the front gives [1,2,3]; reading back to front through prev gives [3,2,1]."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"Deque", "pushBack", "pushBack", "pushFront", "toList", "toListReversed"}, "values": [][]int{{}, {2}, {3}, {1}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, []int{1, 2, 3}, []int{3, 2, 1}}, Note: "push at both ends: front to back is 1,2,3, back to front is 3,2,1"},
			{Input: map[string]interface{}{"operations": []string{"Deque", "popFront", "popBack", "toList"}, "values": [][]int{{}, {}, {}, {}}}, Expected: []interface{}{nil, -1, -1, []int{}}, Note: "popping an empty deque returns -1 at both ends and changes nothing"},
			{Input: map[string]interface{}{"operations": []string{"Deque", "pushBack", "pushBack", "pushBack", "popFront", "popBack", "toList", "toListReversed"}, "values": [][]int{{}, {1}, {2}, {3}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, 1, 3, []int{2}, []int{2}}, Hidden: true, Note: "pops from both ends, then both reads must show only the middle node"},
			{Input: map[string]interface{}{"operations": []string{"Deque", "pushFront", "pushFront", "pushFront", "pushFront", "popBack", "popBack", "toListReversed", "toList"}, "values": [][]int{{}, {1}, {2}, {3}, {4}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, nil, 1, 2, []int{3, 4}, []int{4, 3}}, Hidden: true, Note: "popBack after pushing at the front must find the right node through prev, then the reversed read must agree"},
			{Input: map[string]interface{}{"operations": []string{"Deque", "pushBack", "popBack", "popBack", "pushFront", "popFront", "popFront"}, "values": [][]int{{}, {5}, {}, {}, {6}, {}, {}}}, Expected: []interface{}{nil, nil, 5, -1, nil, 6, -1}, Hidden: true, Note: "a deque of one node empties and then works again from both ends"},
			{Input: map[string]interface{}{"operations": []string{"Deque", "pushBack", "pushBack", "pushBack", "pushBack", "popFront", "popFront", "toListReversed", "pushFront", "toListReversed", "popBack", "toList"}, "values": [][]int{{}, {1}, {2}, {3}, {4}, {}, {}, {}, {9}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, nil, 1, 2, []int{4, 3}, nil, []int{4, 3, 9}, 4, []int{9, 3}}, Hidden: true, Note: "prev links must survive pops from the front and a later push"},
			{Input: map[string]interface{}{"operations": []string{"Deque", "pushFront", "pushBack", "pushFront", "pushBack", "toList", "toListReversed", "popBack", "popBack", "popFront", "popFront", "popFront"}, "values": [][]int{{}, {2}, {3}, {1}, {4}, {}, {}, {}, {}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, nil, nil, []int{1, 2, 3, 4}, []int{4, 3, 2, 1}, 4, 3, 1, 2, -1}, Hidden: true, Note: "alternating pushes, then draining the deque and popping past empty"},
			{Input: map[string]interface{}{"operations": []string{"Deque", "pushBack", "pushFront", "popBack", "pushBack", "toListReversed", "popFront", "toList"}, "values": [][]int{{}, {0}, {0}, {}, {-7}, {}, {}, {}}}, Expected: []interface{}{nil, nil, nil, 0, nil, []int{-7, 0}, 0, []int{-7}}, Hidden: true, Note: "zeros and negatives are values; -1 is only returned for an empty deque"},
		},
		TimeComplexity:  "O(1) per operation, O(n) for toList and toListReversed",
		SpaceComplexity: "O(n)",
		StarterCode: `class Deque:
    def __init__(self):
        # Define your own doubly linked node (val, prev, next)
        # and set up the head and tail sentinels
        pass

    def pushFront(self, val):
        pass

    def pushBack(self, val):
        pass

    def popFront(self):
        # Return the value, or -1 if empty
        pass

    def popBack(self):
        # Return the value, or -1 if empty
        pass

    def toList(self):
        # Values from front to back
        pass

    def toListReversed(self):
        # Values from back to front, following prev
        pass`,
		Harness: Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Inserting and removing in a doubly linked list always touches the same four pointers, whichever end it happens at. With a dummy node at each end, can you write that once and use it for both the front and the back?"},
			{Level: 2, Type: "insight", Content: "Write two private helpers: one that inserts a new node right after a given node, and one that removes a given node by joining its two neighbours. pushFront inserts after the head sentinel, pushBack inserts after tail.prev (the last real node), popFront removes head.next and popBack removes tail.prev. Because the sentinels are always there, no real node has a missing neighbour, so there are no special cases for the first or last element. The deque is empty exactly when head.next is the tail. See the doubly linked list and the LRU storyboard in Chapter 7.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: "1. Define a node with val, prev and next. In __init__ create head and tail sentinels and link them to each other.\n2. _insert_after(node, val): the new node's prev is node and its next is node.next; then fix node.next.prev and node.next, in that order.\n3. _remove(node): node.prev.next = node.next and node.next.prev = node.prev; return the value.\n4. The four push and pop methods call the helpers; a pop checks for empty first.\n5. toList walks next from head.next to the tail; toListReversed walks prev from tail.prev to the head."},
			{Level: 4, Type: "skeleton", Content: "class Node:\n    def __init__(self, val=0):\n        self.val = val\n        self.prev = None\n        self.next = None\n\n\nclass Deque:\n    def __init__(self):\n        self.head = Node()\n        self.tail = Node()\n        self.head.next = self.tail\n        self.tail.prev = ____\n\n    def _insert_after(self, node, val):\n        new = Node(val)\n        new.prev = node\n        new.next = node.next\n        node.next.prev = ____\n        node.next = ____\n\n    def _remove(self, node):\n        node.prev.next = ____\n        node.next.prev = ____\n        return node.val\n\n    def pushFront(self, val):\n        self._insert_after(self.head, val)\n\n    def pushBack(self, val):\n        self._insert_after(____, val)\n\n    def popFront(self):\n        if self.head.next is self.tail:\n            return -1\n        return self._remove(self.head.next)\n\n    def popBack(self):\n        if ____:\n            return -1\n        return self._remove(self.tail.prev)\n\n    def toList(self):\n        out = []\n        node = self.head.next\n        while node is not self.tail:\n            out.append(node.val)\n            node = node.next\n        return out\n\n    def toListReversed(self):\n        out = []\n        node = self.tail.prev\n        while node is not self.head:\n            out.append(node.val)\n            node = ____\n        return out"},
		},
		Solution: Solution{
			Code: `class Node:
    def __init__(self, val=0):
        self.val = val
        self.prev = None
        self.next = None


class Deque:
    def __init__(self):
        self.head = Node()
        self.tail = Node()
        self.head.next = self.tail
        self.tail.prev = self.head

    def _insert_after(self, node, val):
        new = Node(val)
        new.prev = node
        new.next = node.next
        node.next.prev = new
        node.next = new

    def _remove(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev
        return node.val

    def pushFront(self, val):
        self._insert_after(self.head, val)

    def pushBack(self, val):
        self._insert_after(self.tail.prev, val)

    def popFront(self):
        if self.head.next is self.tail:
            return -1
        return self._remove(self.head.next)

    def popBack(self):
        if self.tail.prev is self.head:
            return -1
        return self._remove(self.tail.prev)

    def toList(self):
        out = []
        node = self.head.next
        while node is not self.tail:
            out.append(node.val)
            node = node.next
        return out

    def toListReversed(self):
        out = []
        node = self.tail.prev
        while node is not self.head:
            out.append(node.val)
            node = node.prev
        return out`,
			Explanation: "A doubly linked list with sentinels. Each node knows both neighbours, so removing any node is a local change and both ends are reachable in O(1). The two sentinels guarantee every real node has a prev and a next, so one insert helper and one remove helper serve both ends.",
			KeyInsight:  "Every insert or remove is the same four pointers, and a doubly linked list is only correct when both directions agree. The sentinels remove the empty-list and end-of-list special cases, so the whole class is two helpers plus thin wrappers. The test that reads the deque backwards is what exposes a prev link that was not updated.",
			WhyItWorks:  "Invariant: following next from the head sentinel visits the real nodes front to back and ends at the tail sentinel, and following prev from the tail visits the same nodes back to front; for every node n, n.next.prev is n and n.prev.next is n. The empty deque satisfies it (head and tail point at each other). _insert_after sets the new node's two links from its neighbours and then updates both neighbours' links to the new node, so the two-way property holds again. _remove makes its two neighbours point at each other, which does the same. A pop checks head.next is tail (equivalently tail.prev is head) before removing. Each operation changes a constant number of pointers: O(1). toList and toListReversed visit every node once: O(n).",
			CommonMistakes: []string{
				"Updating next links but not prev links, as in a singly linked list. The forward reads look right, but popBack and toListReversed follow stale prev pointers and return wrong values or -1.",
				"Updating the pointers in the wrong order in the insert: if node.next is changed before node.next.prev is read, the old neighbour is lost. Set the new node's links first, then fix the neighbours.",
				"Not checking for empty before a pop. On an empty deque, head.next is the tail sentinel, and removing it corrupts the sentinels instead of returning -1.",
				"Using a Python list or collections.deque inside. The behaviour matches, but the point is to practise the node links the LRU Cache and other designs depend on.",
			},
			TimeComplexity:  "O(1) per push and pop; O(n) for toList and toListReversed",
			SpaceComplexity: "O(n) for the nodes",
			Walkthrough: []WalkthroughStep{
				{Title: "Link the two sentinels", Explanation: "An empty deque is two dummy nodes pointing at each other in both directions.", CodeSnippet: "self.head.next = self.tail\nself.tail.prev = self.head", LineStart: 12, LineEnd: 13},
				{Title: "Insert after any node", Explanation: "The new node first copies its neighbours from node, and only then do the neighbours point back at it. Four pointers, in an order that never loses the old neighbour.", CodeSnippet: "new.prev = node\nnew.next = node.next\nnode.next.prev = new\nnode.next = new", LineStart: 17, LineEnd: 20},
				{Title: "Remove a node by joining its neighbours", Explanation: "Each side is told to skip the node. The node needs no cleanup of its own.", CodeSnippet: "node.prev.next = node.next\nnode.next.prev = node.prev", LineStart: 23, LineEnd: 24},
				{Title: "Push at both ends with the same helper", Explanation: "Front means after the head sentinel; back means after the last real node, which is tail.prev.", CodeSnippet: "self._insert_after(self.head, val)\nself._insert_after(self.tail.prev, val)", LineStart: 28, LineEnd: 31},
				{Title: "Pop only when something is there", Explanation: "head.next being the tail means no real node exists, so the answer is -1; otherwise remove the first (or last) real node.", CodeSnippet: "if self.head.next is self.tail:\n    return -1\nreturn self._remove(self.head.next)", LineStart: 34, LineEnd: 36},
				{Title: "Read backwards through prev", Explanation: "Starting from tail.prev and following prev until the head sentinel lists the values back to front. A stale prev shows up here.", CodeSnippet: "node = self.tail.prev\nwhile node is not self.head:\n    out.append(node.val)\n    node = node.prev", LineStart: 53, LineEnd: 56},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does popBack need the prev pointers, when popFront does not?", Answer: "The last real node is reached in O(1) as tail.prev, but removing it means making its predecessor point at the tail, and that predecessor is node.prev. In a singly linked list you would have to walk from the head to find it, O(n). The prev links are what make removal at the back O(1)."},
			{Prompt: "What would break if the deque had no sentinels?", Answer: "The first insert into an empty deque, and the removal of the only node, would have to set or clear head and tail themselves, and every helper would need checks for a missing prev or next. With the two sentinels, a real node always has both neighbours, so the same four pointer changes work in every case."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, LinkedListDrills...)
}
