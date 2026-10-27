package practice

// randomListPrelude is the Node class LeetCode provides for Copy List with
// Random Pointer.
const randomListPrelude = `class Node:
    def __init__(self, x, next=None, random=None):
        self.val = int(x)
        self.next = next
        self.random = random
`

// randomListValidator compares the copy with the expected list and checks it
// is a real deep copy: no node of the answer may be a node of the input, and
// the input list must be left as it was.
const randomListValidator = `def check(raw, args, result, expected):
    if result != expected:
        return (False, "the copy's values or pointers differ from the original list")
    nodes = []
    node = args['head']
    while node is not None and len(nodes) <= 10000:
        nodes.append(node)
        node = node.next
    index = {}
    for i, n in enumerate(nodes):
        index[id(n)] = i
    now = []
    for n in nodes:
        r = n.random
        now.append([n.val, None if r is None else index.get(id(r), -1)])
    if now != raw['head']:
        return (False, "the original list was changed; restore its next and random pointers before returning")
    node = result_obj
    steps = 0
    while node is not None and steps <= 10000:
        if id(node) in index:
            return (False, "the copy reuses a node of the original list; every node must be new")
        if node.random is not None and id(node.random) in index:
            return (False, "a random pointer in the copy points into the original list")
        node = node.next
        steps += 1
    return True
`

// LinkedListProblems contains all linked-list category problems, in ladder
// order: each one adds a single new idea to the ones before it.
var LinkedListProblems = []*Problem{
	{
		ID:              "reverse-linked-list",
		Number:          206,
		Order:           1,
		Title:           "Reverse Linked List",
		Slug:            "reverse-linked-list",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Recursion"},
		RelatedChapters: []int{7, 8},
		Description: `You are handed the first node (head) of a singly linked list. Each node holds a value and a pointer, next, to the following node; the last node's next is None.

Turn the list around in place: every next pointer must point to the node that used to come before it, so the old last node becomes the new head. Return that new head.

If the list is empty (head is None), return None, which the judge shows as [].`,
		Constraints: []string{
			"The list has between 0 and 5000 nodes",
			"-5000 <= Node.val <= 5000",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4,5]", Output: "[5,4,3,2,1]", Explanation: "The links 1->2->3->4->5 all flip direction, so reading from the new head 5 gives 5->4->3->2->1."},
			{Input: "head = [1,2]", Output: "[2,1]"},
			{Input: "head = []", Output: "[]", Explanation: "An empty list reversed is still empty."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}}, Expected: []int{5, 4, 3, 2, 1}, Note: "typical odd-length list"},
			{Input: map[string]interface{}{"head": []int{1, 2}}, Expected: []int{2, 1}, Note: "two nodes: the head must end up pointing to None"},
			{Input: map[string]interface{}{"head": []int{}}, Expected: []int{}, Note: "empty list returns None"},
			{Input: map[string]interface{}{"head": []int{7}}, Expected: []int{7}, Hidden: true, Note: "single node"},
			{Input: map[string]interface{}{"head": []int{1, 1, 2}}, Expected: []int{2, 1, 1}, Hidden: true, Note: "duplicate values: rewire nodes, don't compare values"},
			{Input: map[string]interface{}{"head": []int{-5000, 0, 5000}}, Expected: []int{5000, 0, -5000}, Hidden: true, Note: "negative and extreme values"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}}, Expected: []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, Hidden: true, Note: "longer list: returning head instead of prev gives only [1]"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode: `# ListNode(val, next) is provided
def reverseList(head):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"head": "ListNode"},
			ReturnType: "ListNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Stand on one node and imagine flipping its next pointer so it points backwards. Which part of the list can you no longer reach after that flip, and what would you have to remember beforehand to keep going?"},
			{Level: 2, Type: "insight", Content: "Walk the list once with three pointers: prev (the head of the part you have already reversed), curr (the node you are flipping now) and nxt (the rest of the list, saved before you break the link). Each step moves one node from the front of the untouched part to the front of the reversed part. This is the in-place reversal pattern from Chapter 7.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: `1. Set up the pointers: prev starts as None (nothing reversed yet), curr starts at head.
2. Save the next node: before changing anything, remember what comes after curr.
3. Rewire the pointer: make curr point back at prev.
4. Advance the pointers: prev moves to curr, curr moves to the node you saved.
5. Return the new head: when curr runs off the end, prev is the old tail, which is the new head.`},
			{Level: 4, Type: "skeleton", Content: `def reverseList(head):
    prev = None
    curr = head
    while curr:
        nxt = ____          # save the next node
        curr.next = ____    # rewire the pointer
        prev = ____         # advance the pointers
        curr = ____
    return ____`},
		},
		Solution: Solution{
			Code: `def reverseList(head):
    prev = None
    curr = head
    while curr:
        nxt = curr.next
        curr.next = prev
        prev = curr
        curr = nxt
    return prev`,
			Explanation: "This is the first rung of the ladder and the most reused move in the chapter: walk the list once and flip each next pointer backwards. Three pointers are enough. prev marks the reversed part, curr the node being flipped, and nxt the untouched rest, saved before the link is broken. Reorder List (Order 5) reuses this loop as one of its three steps.",
			BruteForce:  "Copy every value into a Python list, reverse that list, and build brand-new nodes from it. That is O(n) time but O(n) extra memory, and it creates new nodes instead of rewiring the ones you were given. The wasted work is the second copy of the whole list: the original nodes already exist and only their pointers need to change.",
			KeyInsight:  "You don't need a second list, only the next pointers turned around. The single danger is that flipping curr.next destroys your only route to the rest of the list, so save it in nxt first. After that, each node needs a constant amount of work.",
			WhyItWorks:  "Invariant: at the top of every loop iteration, prev is the head of the already-visited nodes in reversed order, curr is the head of the not-yet-visited nodes in their original order, and the two pieces are disconnected. At the start prev = None (an empty reversed part) and curr = head (everything unvisited), so it holds. One iteration takes the first unvisited node, points it at prev, and makes it the new front of the reversed part, which keeps the invariant. The loop ends when curr is None, meaning nothing is left unvisited, so prev heads the entire list reversed. Each node is visited exactly once, so the time is O(n), and only three pointers are stored, so the extra space is O(1).",
			CommonMistakes: []string{
				"Writing curr.next = prev before saving curr.next. The rest of the list becomes unreachable and the loop stops after one node. Always save nxt = curr.next first.",
				"Returning head instead of prev. After the loop, head is the old first node, which is now the tail pointing to None, so you would return [1]. The new head is prev.",
				"Advancing with curr = curr.next after the rewire. curr.next now points backwards to prev, so you would step back instead of forward. Advance with the saved nxt.",
			},
			TimeComplexity:  "O(n): each node is visited once",
			SpaceComplexity: "O(1): three pointers, whatever the list length",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the pointers", Explanation: "Nothing is reversed yet, so prev is None. That None also becomes the next pointer of the old head, which is correct because the old head becomes the tail.", CodeSnippet: "prev = None\ncurr = head", LineStart: 2, LineEnd: 3},
				{Title: "Save the next node", Explanation: "Before touching curr.next, keep a reference to the rest of the list in nxt; it is the only way forward once the link is flipped.", CodeSnippet: "nxt = curr.next", LineStart: 5, LineEnd: 5},
				{Title: "Rewire the pointer", Explanation: "Point curr backwards at prev. On [1,2,3] the first iteration makes 1 point to None, the second makes 2 point to 1.", CodeSnippet: "curr.next = prev", LineStart: 6, LineEnd: 6},
				{Title: "Advance the pointers", Explanation: "curr is now the front of the reversed part, so it becomes prev, and the saved nxt becomes the new curr.", CodeSnippet: "prev = curr\ncurr = nxt", LineStart: 7, LineEnd: 8},
				{Title: "Return the new head", Explanation: "When curr is None every node has been flipped and prev is the old tail, which is the new head. For an empty list the loop never runs and prev is still None.", CodeSnippet: "return prev", LineStart: 9, LineEnd: 9},
			},
			Alternatives: []AltSolution{
				{
					Name: "Recursive reversal",
					Code: `def reverseList(head):
    if head is None or head.next is None:
        return head
    new_head = reverseList(head.next)
    head.next.next = head
    head.next = None
    return new_head`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n) for the call stack",
					Explanation:     "Trust the recursive call (Chapter 8): reverseList(head.next) returns the reversed rest, whose tail is head.next. Hook head onto that tail with head.next.next = head and cut head's old link. It reads neatly but keeps one stack frame per node, so a 5000-node list goes deeper than Python's default recursion limit of about 1000.",
				},
				{
					Name: "Brute force: copy the values and rebuild",
					Code: `def reverseList(head):
    values = []
    while head:
        values.append(head.val)
        head = head.next
    dummy = ListNode(0)
    tail = dummy
    for value in reversed(values):
        tail.next = ListNode(value)
        tail = tail.next
    return dummy.next`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Collect the values, then build new nodes in reverse order. It is correct and easy to trust, but it uses O(n) extra memory and doesn't reuse the given nodes, which interviewers usually ask for.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the function return prev and not curr or head?", Answer: "The loop stops when curr is None, so curr has nothing to return. head still points at the original first node, which is now the tail (its next is None). prev was the last node moved into the reversed part, the old tail, so it is the head of the reversed list."},
			{Prompt: "The recursive version is also O(n) time. Why can it crash on a 5000-node list in plain Python when the loop does not?", Answer: "Each recursive call adds a stack frame that stays alive until the deeper calls return, so the stack holds 5000 frames at once. Python's default recursion limit is about 1000, so it raises RecursionError. The loop reuses the same three variables and needs O(1) memory however long the list is."},
			{Prompt: "What if you had to reverse only the nodes from position left to position right (LeetCode 92)? What would you add to this loop?", Answer: "Walk to the node just before position left (a dummy head helps when left = 1), run the same save/rewire/advance loop for right - left + 1 nodes, then reconnect: the node before the section must point to the new front (the old node at right), and the old node at left, now the end of the section, must point to the node after right."},
		},
	},
	{
		ID:              "merge-two-sorted-lists",
		Number:          21,
		Order:           2,
		Title:           "Merge Two Sorted Lists",
		Slug:            "merge-two-sorted-lists",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Recursion"},
		RelatedChapters: []int{7, 3},
		Description: `You receive the heads of two singly linked lists, list1 and list2. Each list is already sorted in non-decreasing order.

Combine them into one sorted list by relinking the existing nodes (don't create new value nodes) and return the head of the combined list. Either input may be empty (None); if both are empty, return None, shown as [].`,
		Constraints: []string{
			"Each list has between 0 and 50 nodes",
			"-100 <= Node.val <= 100",
			"Both list1 and list2 are sorted in non-decreasing order",
		},
		Examples: []Example{
			{Input: "list1 = [1,2,4], list2 = [1,3,4]", Output: "[1,1,2,3,4,4]", Explanation: "Repeatedly take the smaller front node: 1 (list1), 1 (list2), 2, 3, 4, then the leftover 4."},
			{Input: "list1 = [], list2 = []", Output: "[]"},
			{Input: "list1 = [], list2 = [0]", Output: "[0]", Explanation: "One list is empty, so the answer is the other list."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"list1": []int{1, 2, 4}, "list2": []int{1, 3, 4}}, Expected: []int{1, 1, 2, 3, 4, 4}, Note: "interleaved lists with equal values"},
			{Input: map[string]interface{}{"list1": []int{}, "list2": []int{}}, Expected: []int{}, Note: "both lists empty"},
			{Input: map[string]interface{}{"list1": []int{}, "list2": []int{0}}, Expected: []int{0}, Note: "first list empty"},
			{Input: map[string]interface{}{"list1": []int{5}, "list2": []int{1, 2, 3}}, Expected: []int{1, 2, 3, 5}, Hidden: true, Note: "list2 runs out first, so the leftover 5 must be attached after the loop"},
			{Input: map[string]interface{}{"list1": []int{1, 1, 1}, "list2": []int{1, 1}}, Expected: []int{1, 1, 1, 1, 1}, Hidden: true, Note: "all values equal"},
			{Input: map[string]interface{}{"list1": []int{-100, 0, 100}, "list2": []int{-50, 50}}, Expected: []int{-100, -50, 0, 50, 100}, Hidden: true, Note: "negative values and constraint extremes"},
			{Input: map[string]interface{}{"list1": []int{2}, "list2": []int{1}}, Expected: []int{1, 2}, Hidden: true, Note: "the head of the answer comes from list2"},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(1)",
		StarterCode: `# ListNode(val, next) is provided
def mergeTwoLists(list1, list2):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"list1": "ListNode", "list2": "ListNode"},
			ReturnType: "ListNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Both lists are sorted, so where must the smallest remaining value always be? And when you attach the very first node, you have no 'previous node' yet. How could you avoid writing a special case for that?"},
			{Level: 2, Type: "insight", Content: "Only the two front nodes are ever candidates, like the merge step of merge sort. Build the answer behind a dummy head: a throwaway node whose next will be the real head, plus a tail pointer to the last node attached. With a dummy, attaching the first node is the same as attaching any other. See the dummy head concept in Chapter 7.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: `1. Add a dummy head: create one placeholder node and point tail at it.
2. Pick the smaller front: while both lists have nodes, attach the smaller front node to tail and step that list forward.
3. Advance the tail: move tail onto the node you just attached.
4. Attach the leftovers: when one list runs out, the other is already sorted, so link all of it at once.
5. Return the real head: dummy.next.`},
			{Level: 4, Type: "skeleton", Content: `def mergeTwoLists(list1, list2):
    dummy = ListNode(0)
    tail = dummy
    while list1 and list2:
        if ____:
            tail.next = list1
            list1 = list1.next
        else:
            tail.next = list2
            list2 = list2.next
        tail = ____
    tail.next = ____
    return ____`},
		},
		Solution: Solution{
			Code: `def mergeTwoLists(list1, list2):
    dummy = ListNode(0)
    tail = dummy
    while list1 and list2:
        if list1.val <= list2.val:
            tail.next = list1
            list1 = list1.next
        else:
            tail.next = list2
            list2 = list2.next
        tail = tail.next
    tail.next = list1 if list1 else list2
    return dummy.next`,
			Explanation: "Builds on Reverse Linked List (Order 1), where you rewired existing nodes, and adds the dummy head: a placeholder node in front of the answer so that the first attachment needs no special case. A tail pointer always sits on the last attached node, and at each step the smaller of the two front nodes is spliced on.",
			BruteForce:  "Walk both lists to collect every value, sort the n + m values, and build a new list. That costs O((n + m) log(n + m)) time and O(n + m) extra memory. The sort is wasted work: both inputs are already sorted, so their order is known and only needs interleaving.",
			KeyInsight:  "Because both lists are sorted, the smallest remaining value is always at the front of one of the two lists. Comparing the two fronts gives the next node in O(1), so a single pass produces the merge. The dummy head removes the 'is the answer still empty?' branch.",
			WhyItWorks:  "Invariant: the nodes from dummy.next to tail are the smallest values seen so far, in sorted order, and every node still in list1 or list2 is at least tail.val. Taking the smaller front keeps this true, because that node is no bigger than anything left in either list (both are sorted). When one list is empty, the other holds only values at least tail.val, already in order, so linking it whole finishes the sorted merge. Each node is attached exactly once, giving O(n + m) time with O(1) extra space, since only pointers change.",
			CommonMistakes: []string{
				"Forgetting tail = tail.next. Every new node then overwrites dummy.next and the answer holds just the last node attached. Move tail after every attachment.",
				"Not attaching the leftovers after the loop. The loop stops as soon as either list is empty, so [5] + [1,2,3] would lose the 5. Link the non-empty remainder with one assignment.",
				"Special-casing the first node (if head is None: head = ...) instead of using a dummy. The extra branches are easy to get wrong when the first node comes from list2. A dummy head makes every attachment the same.",
			},
			TimeComplexity:  "O(n + m): each node is attached once",
			SpaceComplexity: "O(1): the nodes are relinked, only the dummy is new",
			Walkthrough: []WalkthroughStep{
				{Title: "Add a dummy head", Explanation: "The dummy is a fixed node in front of the answer, so tail always has a node to attach to, even before the first real node exists.", CodeSnippet: "dummy = ListNode(0)\ntail = dummy", LineStart: 2, LineEnd: 3},
				{Title: "Pick the smaller front", Explanation: "Only the two fronts can be the next smallest value. Using <= takes list1 first on ties, which keeps equal values in a stable order.", CodeSnippet: "if list1.val <= list2.val:\n    tail.next = list1\n    list1 = list1.next", LineStart: 4, LineEnd: 10},
				{Title: "Advance the tail", Explanation: "tail moves onto the node just attached, so the next attachment goes after it.", CodeSnippet: "tail = tail.next", LineStart: 11, LineEnd: 11},
				{Title: "Attach the leftovers", Explanation: "One list is now empty. The other is sorted and every value in it is at least tail.val, so link it whole.", CodeSnippet: "tail.next = list1 if list1 else list2", LineStart: 12, LineEnd: 12},
				{Title: "Return the real head", Explanation: "The dummy itself isn't part of the answer; its next is.", CodeSnippet: "return dummy.next", LineStart: 13, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Recursive merge",
					Code: `def mergeTwoLists(list1, list2):
    if list1 is None:
        return list2
    if list2 is None:
        return list1
    if list1.val <= list2.val:
        list1.next = mergeTwoLists(list1.next, list2)
        return list1
    list2.next = mergeTwoLists(list1, list2.next)
    return list2`,
					TimeComplexity:  "O(n + m)",
					SpaceComplexity: "O(n + m) for the call stack",
					Explanation:     "The smaller front node is the head of the answer, and its next is the merge of everything else, a smaller instance of the same problem (Chapter 8). Short and elegant, but it uses one stack frame per node.",
				},
				{
					Name: "Brute force: collect, sort and rebuild",
					Code: `def mergeTwoLists(list1, list2):
    values = []
    for node in (list1, list2):
        while node:
            values.append(node.val)
            node = node.next
    dummy = ListNode(0)
    tail = dummy
    for value in sorted(values):
        tail.next = ListNode(value)
        tail = tail.next
    return dummy.next`,
					TimeComplexity:  "O((n + m) log(n + m))",
					SpaceComplexity: "O(n + m)",
					Explanation:     "Ignores that the inputs are sorted and sorts again from scratch, building new nodes. Correct, but slower and it uses extra memory.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is a single assignment enough to attach the leftovers after the loop, instead of a second loop?", Answer: "The leftover list is already a chain of sorted nodes, and every value in it is at least as large as tail.val. Linking its head to tail.next brings the whole chain along, because the nodes are still connected to each other. A loop would only copy links that already exist."},
			{Prompt: "What would go wrong if you used < instead of <= when comparing the fronts?", Answer: "The merged list would still be sorted, since on a tie either node is a valid next choice. Only stability changes: equal values from list2 would come before those from list1. For plain integers the output is identical, which is why the tests can't tell the difference."},
			{Prompt: "What if you had k sorted lists instead of two? How could you reuse this function?", Answer: "Merge them in pairs: merge lists 0 and 1, 2 and 3, and so on, then repeat on the results until one list remains. Each round touches every node once and halves the number of lists, so it costs O(N log k). That is the main solution of Merge K Sorted Lists (Order 9)."},
		},
	},
	{
		ID:              "linked-list-cycle",
		Number:          141,
		Order:           3,
		Title:           "Linked List Cycle",
		Slug:            "linked-list-cycle",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Two Pointers", "Hash Table"},
		RelatedChapters: []int{7, 3},
		Description: `You are given the head of a singly linked list. The list may contain a cycle: some node's next pointer leads back to a node that appeared earlier, so following next pointers from head would go round forever.

Return True if the list contains a cycle and False otherwise.

How the tests describe a cycle: the input pos is the 0-based index of the node that the last node points back to, or -1 for no cycle. pos is only used to build the list; it is not passed to your function.`,
		Constraints: []string{
			"The list has between 0 and 10^4 nodes",
			"-10^5 <= Node.val <= 10^5",
			"pos is -1 or a valid index in the list",
			"Follow-up: use O(1) extra memory",
		},
		Examples: []Example{
			{Input: "head = [3,2,0,-4], pos = 1", Output: "true", Explanation: "The last node (-4) points back to the node at index 1 (value 2), so 2 -> 0 -> -4 -> 2 repeats forever."},
			{Input: "head = [1,2], pos = 0", Output: "true"},
			{Input: "head = [1], pos = -1", Output: "false", Explanation: "The only node's next is None, so the walk ends."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{3, 2, 0, -4}, "pos": 1}, Expected: true, Note: "cycle that starts in the middle"},
			{Input: map[string]interface{}{"head": []int{1, 2}, "pos": 0}, Expected: true, Note: "two-node cycle back to the head"},
			{Input: map[string]interface{}{"head": []int{1}, "pos": -1}, Expected: false, Note: "single node, no cycle"},
			{Input: map[string]interface{}{"head": []int{}, "pos": -1}, Expected: false, Hidden: true, Note: "empty list"},
			{Input: map[string]interface{}{"head": []int{1}, "pos": 0}, Expected: true, Hidden: true, Note: "a node that points to itself"},
			{Input: map[string]interface{}{"head": []int{1, 2}, "pos": -1}, Expected: false, Hidden: true, Note: "even length without a cycle: fast.next.next must not crash"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}, "pos": -1}, Expected: false, Hidden: true, Note: "odd length without a cycle"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5, 6}, "pos": 5}, Expected: true, Hidden: true, Note: "the tail points to itself"},
			{Input: map[string]interface{}{"head": []int{7, 7, 7, 7}, "pos": 2}, Expected: true, Hidden: true, Note: "equal values: compare nodes, not values"},
			{Input: map[string]interface{}{"head": []int{7, 7, 7, 7}, "pos": -1}, Expected: false, Hidden: true, Note: "equal values without a cycle: comparing values would say true"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode: `# ListNode(val, next) is provided
def hasCycle(head):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"head": "CycleList:pos"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you just walk the list, a cycle means the walk never ends, so you can't wait for None. Picture two runners on a track, one twice as fast as the other. What happens if the track is a loop, and what happens if it is a straight road?"},
			{Level: 2, Type: "insight", Content: "Floyd's fast and slow pointers: slow moves one node per step and fast moves two. With no cycle, fast reaches None. With a cycle, both end up going round it, and fast gains exactly one node on slow each step, so it must land on slow. Chapter 7 explains why it can't jump over.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: `1. Start both pointers at head.
2. Check fast can move: continue only while fast and fast.next exist; otherwise the list ends, so there is no cycle.
3. Move slow one, fast two.
4. Check for a meeting: if slow and fast are the same node (identity, not value), there is a cycle.
5. If the loop ends because fast hit the end, return False.`},
			{Level: 4, Type: "skeleton", Content: `def hasCycle(head):
    slow = head
    fast = head
    while ____:
        slow = ____
        fast = ____
        if ____:
            return True
    return False`},
		},
		Solution: Solution{
			Code: `def hasCycle(head):
    slow = head
    fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True
    return False`,
			Explanation: "Introduces the fast and slow pointers of Chapter 7. Like the two pointers of Chapter 3, two positions move through the data, but here they move in the same direction at different speeds. Builds on the traversal of Reverse Linked List (Order 1); Remove Nth Node (Order 4) and Reorder List (Order 5) reuse the same two-speed idea.",
			BruteForce:  "Remember every node you visit in a set and stop when you see one twice. That is O(n) time but O(n) extra memory for the set. The memory is the waste: you store every node just to notice one repeat.",
			KeyInsight:  "Two pointers moving at speeds 1 and 2 must meet inside a cycle, because the gap between them shrinks by exactly one node per step. Without a cycle the fast pointer simply reaches None. Two pointers replace the whole set.",
			WhyItWorks:  "No cycle: fast moves ahead of slow and reaches the end (fast or fast.next is None) after about n/2 steps, so the function returns False. With a cycle of length L: after at most n steps slow is inside the cycle, and fast is already there. Measure the gap as how many nodes fast is behind slow going round the cycle, a number from 0 to L - 1. Each step fast moves 2 and slow moves 1, so the gap shrinks by exactly 1. A number that drops by 1 each step from at most L - 1 must hit 0, so they land on the same node within L more steps. Fast can't jump over slow, because the gap never drops by 2. Total time O(n), extra space O(1).",
			CommonMistakes: []string{
				"Looping on while fast.next: when fast itself is None (a list of even length) this raises AttributeError. Check fast and fast.next, in that order, so short-circuiting protects you.",
				"Comparing values (slow.val == fast.val) instead of nodes. Lists with repeated values like [7,7,7,7] would report a cycle that isn't there. Compare identity with slow is fast.",
				"Checking slow is fast before moving them. Both start at head, so the very first check would return True for every non-empty list. Move first, then compare.",
			},
			TimeComplexity:  "O(n): at most the tail length plus one lap of the cycle",
			SpaceComplexity: "O(1): two pointers",
			Walkthrough: []WalkthroughStep{
				{Title: "Start both pointers at head", Explanation: "Both runners start together. Because they start equal, the meeting check must come after the first move.", CodeSnippet: "slow = head\nfast = head", LineStart: 2, LineEnd: 3},
				{Title: "Check fast can move", Explanation: "fast needs two steps, so both fast and fast.next must exist. If either is None, the list has an end and therefore no cycle.", CodeSnippet: "while fast and fast.next:", LineStart: 4, LineEnd: 4},
				{Title: "Move slow one, fast two", Explanation: "On [3,2,0,-4] with pos 1: after step 1 slow=2, fast=0; after step 2 slow=0, fast=2; after step 3 slow=-4, fast=-4.", CodeSnippet: "slow = slow.next\nfast = fast.next.next", LineStart: 5, LineEnd: 6},
				{Title: "Check for a meeting", Explanation: "Meeting means fast caught up from behind, which is only possible by going round a loop. Use is to compare nodes, not values.", CodeSnippet: "if slow is fast:\n    return True", LineStart: 7, LineEnd: 8},
				{Title: "No meeting: the list ends", Explanation: "Leaving the loop means fast found the end of the list.", CodeSnippet: "return False", LineStart: 9, LineEnd: 9},
			},
			Alternatives: []AltSolution{
				{
					Name: "Hash set of visited nodes",
					Code: `def hasCycle(head):
    seen = set()
    node = head
    while node:
        if node in seen:
            return True
        seen.add(node)
        node = node.next
    return False`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The Chapter 2 approach: store each node (by identity) and stop at the first repeat. Easy to see it is correct, but it needs memory for every node.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can't the fast pointer skip over the slow pointer inside the cycle?", Answer: "Each step the gap between them (how far fast is behind slow along the cycle) changes by fast's speed minus slow's speed, 2 - 1 = 1. It goes L-1, L-2, ..., 1, 0 and can't jump from 1 to -1. When it reaches 0 they are on the same node."},
			{Prompt: "What if fast moved three nodes per step instead of two? Would it still always detect a cycle?", Answer: "Not with a check only after each full step. The gap would shrink by 2 per step, so if it starts odd it goes 3, 1, then wraps to L - 1 and can skip past slow; on a cycle of even length it might never land exactly on slow. Speeds 1 and 2 are the standard choice because the gap closes by exactly one."},
			{Prompt: "How would you find the node where the cycle begins (LeetCode 142)?", Answer: "After slow and fast meet, move one pointer back to head and advance both one step at a time; they meet at the start of the cycle. If the tail before the cycle has length a, the meeting point is a steps (mod the cycle length) away from the cycle's start, which is the same distance as from head."},
		},
	},
	{
		ID:              "remove-nth-node-from-end",
		Number:          19,
		Order:           4,
		Title:           "Remove Nth Node From End of List",
		Slug:            "remove-nth-node-from-end-of-list",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Two Pointers"},
		RelatedChapters: []int{7, 3},
		Description: `You are given the head of a singly linked list and an integer n. Delete the node that is n-th from the end (n = 1 means the last node, n = length means the first node) and return the head of the resulting list.

If the list had only one node and it is removed, return None, shown as [].

Follow-up: can you do it in a single pass over the list?`,
		Constraints: []string{
			"The list has sz nodes, with 1 <= sz <= 30",
			"0 <= Node.val <= 100",
			"1 <= n <= sz",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4,5], n = 2", Output: "[1,2,3,5]", Explanation: "The 2nd node from the end has value 4. Its predecessor 3 is linked straight to 5."},
			{Input: "head = [1], n = 1", Output: "[]"},
			{Input: "head = [1,2], n = 1", Output: "[1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}, "n": 2}, Expected: []int{1, 2, 3, 5}, Note: "remove from the middle"},
			{Input: map[string]interface{}{"head": []int{1}, "n": 1}, Expected: []int{}, Note: "remove the only node"},
			{Input: map[string]interface{}{"head": []int{1, 2}, "n": 1}, Expected: []int{1}, Note: "remove the last node"},
			{Input: map[string]interface{}{"head": []int{1, 2}, "n": 2}, Expected: []int{2}, Hidden: true, Note: "remove the head: needs a dummy or a special case"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3}, "n": 3}, Expected: []int{2, 3}, Hidden: true, Note: "remove the head of a longer list (n equals length)"},
			{Input: map[string]interface{}{"head": []int{10, 20, 30, 40}, "n": 1}, Expected: []int{10, 20, 30}, Hidden: true, Note: "remove the tail of an even-length list"},
			{Input: map[string]interface{}{"head": []int{1, 1, 1}, "n": 2}, Expected: []int{1, 1}, Hidden: true, Note: "duplicate values: count positions, not values"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode: `# ListNode(val, next) is provided
def removeNthFromEnd(head, n):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"head": "ListNode"},
			ReturnType: "ListNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "A singly linked list can't be walked backwards. If one pointer started n nodes ahead of another and both moved at the same speed, where would the trailing pointer be when the leading one fell off the end?"},
			{Level: 2, Type: "insight", Content: "Keep a fixed gap between two pointers. Start both at a dummy head (Merge Two Sorted Lists, Order 2), move lead n + 1 nodes ahead, then move both one step at a time. When lead becomes None, trail sits just before the node to delete, which is exactly where you need to be to unlink it. This is the gap variant of fast and slow pointers from Chapter 7.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: `1. Add a dummy head in front of head, and start lead and trail on it.
2. Open the gap: move lead n + 1 nodes forward.
3. Move together: advance both until lead is None.
4. Unlink the node: trail.next is the n-th node from the end, so skip it.
5. Return the real head: dummy.next (the old head may have been removed).`},
			{Level: 4, Type: "skeleton", Content: `def removeNthFromEnd(head, n):
    dummy = ListNode(0, head)
    lead = dummy
    trail = dummy
    for _ in range(____):
        lead = lead.next
    while lead:
        lead = lead.next
        trail = trail.next
    trail.next = ____
    return ____`},
		},
		Solution: Solution{
			Code: `def removeNthFromEnd(head, n):
    dummy = ListNode(0, head)
    lead = dummy
    trail = dummy
    for _ in range(n + 1):
        lead = lead.next
    while lead:
        lead = lead.next
        trail = trail.next
    trail.next = trail.next.next
    return dummy.next`,
			Explanation: "Combines the dummy head from Merge Two Sorted Lists (Order 2) with the two-pointer walk of Linked List Cycle (Order 3). The new idea is a fixed gap instead of different speeds: lead starts n + 1 nodes ahead of trail, so when lead falls off the end, trail is the node just before the one to delete.",
			BruteForce:  "Walk the list once to count its length L, then walk again L - n - 1 steps to the predecessor and unlink. That is still O(L) time, but two passes. The first pass exists only to learn L, information the gap gives you for free.",
			KeyInsight:  "The n-th node from the end is n steps before None. If lead is n + 1 steps ahead of trail and they move together, then when lead reaches None, trail is n + 1 steps before None, that is, the predecessor of the target. The dummy gives the head a predecessor too.",
			WhyItWorks:  "Count positions with the dummy at 0 and the real nodes at 1..L, and let None be position L + 1. After the first loop lead is at position n + 1 and trail at 0, a gap of n + 1. Moving together keeps the gap. lead stops at L + 1 (None), so trail stops at L + 1 - (n + 1) = L - n, the node just before position L - n + 1, which is the n-th from the end. Because 1 <= n <= L, trail is between the dummy and node L - 1, so trail.next always exists. When n = L, trail stays on the dummy and the head is removed through dummy.next. One pass, O(L) time, O(1) space.",
			CommonMistakes: []string{
				"Opening a gap of n instead of n + 1. trail then stops on the node to delete instead of its predecessor, and you can't unlink a node in a singly linked list without its predecessor. Count the dummy: n + 1 moves.",
				"Starting the pointers at head with no dummy. Removing the head ([1,2], n = 2) then needs trail before the head, which doesn't exist, and the code crashes or returns the wrong list. Start at a dummy and return dummy.next.",
				"Returning head instead of dummy.next. When the head itself is removed, head still points at the deleted node.",
			},
			TimeComplexity:  "O(L): one pass over the L nodes",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Add a dummy head", Explanation: "The dummy gives the real head a predecessor, so removing the head is no different from removing any other node.", CodeSnippet: "dummy = ListNode(0, head)\nlead = dummy\ntrail = dummy", LineStart: 2, LineEnd: 4},
				{Title: "Open the gap", Explanation: "Move lead n + 1 times. On [1,2,3,4,5] with n = 2, lead goes dummy -> 1 -> 2 -> 3.", CodeSnippet: "for _ in range(n + 1):\n    lead = lead.next", LineStart: 5, LineEnd: 6},
				{Title: "Move together", Explanation: "Both advance until lead is None: lead 3->4->5->None while trail dummy->1->2->3. trail stops on 3, the predecessor of 4.", CodeSnippet: "while lead:\n    lead = lead.next\n    trail = trail.next", LineStart: 7, LineEnd: 9},
				{Title: "Unlink the node", Explanation: "Skip the target by pointing trail past it: 3 now points to 5.", CodeSnippet: "trail.next = trail.next.next", LineStart: 10, LineEnd: 10},
				{Title: "Return the real head", Explanation: "dummy.next is the head of the answer, even if the old head was the node removed.", CodeSnippet: "return dummy.next", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Two passes: count, then walk",
					Code: `def removeNthFromEnd(head, n):
    length = 0
    node = head
    while node:
        length += 1
        node = node.next
    if n == length:
        return head.next
    prev = head
    for _ in range(length - n - 1):
        prev = prev.next
    prev.next = prev.next.next
    return head`,
					TimeComplexity:  "O(L)",
					SpaceComplexity: "O(1)",
					Explanation:     "Count the nodes, then walk to the node before index length - n. Same complexity, but two passes and a special case for removing the head.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the gap be n + 1 and not n?", Answer: "To delete a node in a singly linked list you must change its predecessor's next pointer, so trail has to stop one node before the target. A gap of n would park trail on the target itself. The extra step comes from starting at the dummy, which also makes the head removable."},
			{Prompt: "What if the list could be shorter than n? How would you guard against it?", Answer: "While opening the gap, check that lead is not None before each move. If lead runs out during that first loop, the list has fewer than n nodes, so there is nothing to remove: return the list unchanged (or raise an error, depending on the specification)."},
			{Prompt: "How would the same gap idea find the middle node instead?", Answer: "The middle has no fixed distance from the end, so a fixed gap doesn't help. Use different speeds instead: move fast two steps and slow one, and when fast reaches the end slow is in the middle. That is the first step of Reorder List (Order 5)."},
		},
	},
	{
		ID:              "reorder-list",
		Number:          143,
		Order:           5,
		Title:           "Reorder List",
		Slug:            "reorder-list",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Two Pointers", "Stack", "Recursion"},
		RelatedChapters: []int{7, 3},
		Description: `You are given the head of a singly linked list whose nodes, in order, are L0, L1, L2, ..., Ln-1, Ln.

Rearrange the nodes in place into the order L0, Ln, L1, Ln-1, L2, Ln-2, ...: first node, last node, second node, second-to-last node, and so on.

Change only the next pointers; don't change any node's value. The function returns nothing; the judge reads the list starting from the original head.`,
		Constraints: []string{
			"The list has between 1 and 5 * 10^4 nodes",
			"1 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4]", Output: "[1,4,2,3]", Explanation: "Take from the front and the back alternately: 1 (front), 4 (back), 2 (front), 3 (back)."},
			{Input: "head = [1,2,3,4,5]", Output: "[1,5,2,4,3]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4}}, Expected: []int{1, 4, 2, 3}, Note: "even length"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}}, Expected: []int{1, 5, 2, 4, 3}, Note: "odd length: the middle node ends last"},
			{Input: map[string]interface{}{"head": []int{1}}, Expected: []int{1}, Hidden: true, Note: "single node stays as it is"},
			{Input: map[string]interface{}{"head": []int{1, 2}}, Expected: []int{1, 2}, Hidden: true, Note: "two nodes are already in order"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3}}, Expected: []int{1, 3, 2}, Hidden: true, Note: "three nodes: catches a missing cut (cycle) after the middle"},
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5, 6}}, Expected: []int{1, 6, 2, 5, 3, 4}, Hidden: true, Note: "longer even list"},
			{Input: map[string]interface{}{"head": []int{5, 5, 7, 7, 9}}, Expected: []int{5, 9, 5, 7, 7}, Hidden: true, Note: "duplicate values: move nodes, don't swap values"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode: `# ListNode(val, next) is provided
def reorderList(head):
    # Modify the list in place; return nothing
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"head": "ListNode"},
			ReturnType: "ListNode",
			InPlace:    "head",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The answer alternates between the front half read forwards and the back half read backwards. Which three things that you have already done in this ladder would turn the list into exactly those two pieces and then zip them together?"},
			{Level: 2, Type: "insight", Content: "Three known steps in a row: find the middle with fast and slow pointers (Linked List Cycle, Order 3), reverse the second half (Reverse Linked List, Order 1), then merge the two halves by alternating nodes (like Merge Two Sorted Lists, Order 2, but taking turns instead of comparing). Cut the list at the middle so the halves are separate. See Chapter 7.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: `1. Find the middle: slow moves one, fast moves two; when fast can't move, slow is the last node of the first half.
2. Cut the list: remember slow.next as the start of the second half, then set slow.next = None so the first half ends.
3. Reverse the second half with the save / rewire / advance loop.
4. Weave the halves: save BOTH next pointers first, then point first -> second and second -> (the saved first.next), then advance both pointers to their saved nodes.`},
			{Level: 4, Type: "skeleton", Content: `def reorderList(head):
    if head is None or head.next is None:
        return
    slow, fast = head, head.next
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    second = slow.next
    slow.next = ____                      # cut the list
    prev = None
    while second:
        nxt = second.next
        second.next = prev
        prev = second
        second = nxt
    first, second = head, prev
    while second:
        first_next, second_next = ____, ____   # save both first
        first.next = second
        second.next = ____
        first, second = first_next, second_next`},
		},
		Solution: Solution{
			Code: `def reorderList(head):
    if head is None or head.next is None:
        return
    slow, fast = head, head.next
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    second = slow.next
    slow.next = None
    prev = None
    while second:
        nxt = second.next
        second.next = prev
        prev = second
        second = nxt
    first, second = head, prev
    while second:
        first_next, second_next = first.next, second.next
        first.next = second
        second.next = first_next
        first, second = first_next, second_next`,
			Explanation: "No new technique, just the first three rungs of the ladder chained together: find the middle with fast and slow pointers (Order 3), reverse the second half (Order 1), and weave the halves (the splicing of Order 2). Learning to see a harder problem as a sequence of solved subproblems is the skill this rung adds.",
			BruteForce:  "Copy every node into a Python list, then relink nodes[0], nodes[n-1], nodes[1], nodes[n-2], ... with two indices moving inwards. It is O(n) time and easy to get right, but it uses O(n) extra memory to get the random access a linked list doesn't have.",
			KeyInsight:  "Reading the back half backwards is the same as reading a reversed back half forwards. Once the second half is reversed in place, both halves can be walked forwards with two pointers, so no array is needed.",
			WhyItWorks:  "Finding the middle with fast starting at head.next makes slow stop at node ceil(n/2), so the first half has the same number of nodes as the second or one more. Setting slow.next = None makes the halves separate lists, which is what stops the result from looping. Reversing the second half gives Ln, Ln-1, ... in forward order. The weave then takes one node from each list in turn. Because the first half is never shorter, the loop can stop when the second half runs out: on odd lengths the leftover middle node is already at the end of the first half with next = None. Each node is touched a constant number of times in each phase, so the time is O(n), and only a few pointers are used, so the space is O(1).",
			CommonMistakes: []string{
				"Forgetting the cut slow.next = None. The end of the first half still points into the second half, the weave creates a cycle, and the judge reports <cycle>. Cut before reversing.",
				"Weaving with first.next = second and then second.next = first.next. By the second line first.next already IS second, so second points to itself and the list becomes a cycle. Save first.next (and second.next) in variables before changing any pointer.",
				"Swapping values instead of nodes. The task requires moving nodes, and it doesn't help anyway: you still can't reach the back of the list without a reversal or an array.",
			},
			TimeComplexity:  "O(n): three linear passes",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Find the middle", Explanation: "slow moves one and fast two. On [1,2,3,4,5] slow ends on 3, so the first half is 1,2,3 and the second half 4,5.", CodeSnippet: "slow, fast = head, head.next\nwhile fast and fast.next:\n    slow = slow.next\n    fast = fast.next.next", LineStart: 4, LineEnd: 7},
				{Title: "Cut the list", Explanation: "Keep the start of the second half, then end the first half with None so the two halves can't loop into each other.", CodeSnippet: "second = slow.next\nslow.next = None", LineStart: 8, LineEnd: 9},
				{Title: "Reverse the second half", Explanation: "The exact loop from Reverse Linked List: 4 -> 5 becomes 5 -> 4, and prev is its head.", CodeSnippet: "nxt = second.next\nsecond.next = prev\nprev = second\nsecond = nxt", LineStart: 10, LineEnd: 15},
				{Title: "Weave the halves", Explanation: "Save both next pointers, then link first -> second -> old first.next. On the example: 1 -> 5 -> 2, then 2 -> 4 -> 3, giving 1,5,2,4,3.", CodeSnippet: "first_next, second_next = first.next, second.next\nfirst.next = second\nsecond.next = first_next", LineStart: 16, LineEnd: 21},
			},
			Alternatives: []AltSolution{
				{
					Name: "Array of nodes and two indices",
					Code: `def reorderList(head):
    nodes = []
    node = head
    while node:
        nodes.append(node)
        node = node.next
    i, j = 0, len(nodes) - 1
    while i < j:
        nodes[i].next = nodes[j]
        i += 1
        if i == j:
            break
        nodes[j].next = nodes[i]
        j -= 1
    if nodes:
        nodes[i].next = None`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Store the nodes so you can index from both ends, then link front, back, front, back while i and j move inwards (converging two pointers, Chapter 3). The node where they meet becomes the tail. Simple, but O(n) extra memory.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the weave loop stop when second runs out, and not when first runs out?", Answer: "Starting fast at head.next makes the first half equal in size or one node longer than the second. So the second half always runs out first or at the same time. Any leftover first-half node is the middle node, which correctly stays last and already has next = None because of the cut."},
			{Prompt: "Why must both next pointers be saved before relinking in the weave?", Answer: "Assigning first.next = second destroys the only reference to the rest of the first half, and assigning second.next = ... destroys the rest of the second half. Saving them first means the relinking can't lose nodes or point a node at itself (the cycle bug in the hint first.next = second; second.next = first.next)."},
			{Prompt: "What if you were allowed O(n) memory? Which version would you write in an interview and why?", Answer: "The array version is shorter and harder to get wrong, so it is a reasonable first answer. Interviewers usually follow up with 'now in O(1) space', and the middle / reverse / weave version answers that by reusing three standard linked-list moves."},
		},
	},
	{
		ID:              "add-two-numbers",
		Number:          2,
		Order:           6,
		Title:           "Add Two Numbers",
		Slug:            "add-two-numbers",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Math", "Recursion"},
		RelatedChapters: []int{7, 19},
		Description: `Two non-negative integers are stored as linked lists, one decimal digit per node, with the digits in reverse order: the head holds the ones digit, the next node the tens digit, and so on. For example 342 is stored as 2 -> 4 -> 3.

Return a new linked list, in the same reverse-digit format, holding the sum of the two numbers. Neither number has leading zeros, except the number 0 itself, which is the single node [0].`,
		Constraints: []string{
			"Each list has between 1 and 100 nodes",
			"0 <= Node.val <= 9",
			"The stored numbers have no leading zeros (except 0 itself)",
		},
		Examples: []Example{
			{Input: "l1 = [2,4,3], l2 = [5,6,4]", Output: "[7,0,8]", Explanation: "342 + 465 = 807, stored in reverse as 7 -> 0 -> 8. Column by column: 2+5 = 7, 4+6 = 10 (write 0, carry 1), 3+4+1 = 8."},
			{Input: "l1 = [0], l2 = [0]", Output: "[0]"},
			{Input: "l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]", Output: "[8,9,9,9,0,0,0,1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"l1": []int{2, 4, 3}, "l2": []int{5, 6, 4}}, Expected: []int{7, 0, 8}, Note: "one carry in the middle"},
			{Input: map[string]interface{}{"l1": []int{0}, "l2": []int{0}}, Expected: []int{0}, Note: "zero plus zero"},
			{Input: map[string]interface{}{"l1": []int{9, 9, 9, 9, 9, 9, 9}, "l2": []int{9, 9, 9, 9}}, Expected: []int{8, 9, 9, 9, 0, 0, 0, 1}, Note: "different lengths and a carry that ripples through"},
			{Input: map[string]interface{}{"l1": []int{5}, "l2": []int{5}}, Expected: []int{0, 1}, Hidden: true, Note: "final carry creates a new node"},
			{Input: map[string]interface{}{"l1": []int{1, 8}, "l2": []int{0}}, Expected: []int{1, 8}, Hidden: true, Note: "second list shorter, no carry"},
			{Input: map[string]interface{}{"l1": []int{2, 4, 9}, "l2": []int{5, 6, 4, 9}}, Expected: []int{7, 0, 4, 0, 1}, Hidden: true, Note: "first list shorter, carry into the longer list and past its end"},
			{Input: map[string]interface{}{"l1": []int{0, 0, 1}, "l2": []int{0, 0, 9}}, Expected: []int{0, 0, 0, 1}, Hidden: true, Note: "zeros in low digits: 100 + 900 = 1000"},
		},
		TimeComplexity:  "O(max(m, n))",
		SpaceComplexity: "O(1)",
		StarterCode: `# ListNode(val, next) is provided
def addTwoNumbers(l1, l2):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"l1": "ListNode", "l2": "ListNode"},
			ReturnType: "ListNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When you add 342 + 465 on paper, which column do you start with, and what do you pass from one column to the next? Where is that column in these lists?"},
			{Level: 2, Type: "insight", Content: "The lists already start at the ones digit, exactly where column addition starts. Walk both lists together, add the two digits plus a carry, write total % 10 into a new node and carry total // 10 to the next column. Build the answer behind a dummy head (Merge Two Sorted Lists, Order 2, and Chapter 7). Keep going while either list has digits OR a carry remains.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: `1. Add a dummy head and a tail pointer; carry starts at 0.
2. Loop while l1, l2 or carry is non-zero.
3. Add one column: start total at carry, add l1's digit and step l1 if it exists, same for l2.
4. Split the total: the new digit is total % 10, the new carry is total // 10.
5. Append the digit: attach a new node to tail and advance the tail.
6. Return the real head: dummy.next.`},
			{Level: 4, Type: "skeleton", Content: `def addTwoNumbers(l1, l2):
    dummy = ListNode(0)
    tail = dummy
    carry = 0
    while ____:
        total = carry
        if l1:
            total += l1.val
            l1 = l1.next
        if l2:
            total += l2.val
            l2 = l2.next
        carry, digit = ____
        tail.next = ____
        tail = tail.next
    return dummy.next`},
		},
		Solution: Solution{
			Code: `def addTwoNumbers(l1, l2):
    dummy = ListNode(0)
    tail = dummy
    carry = 0
    while l1 or l2 or carry:
        total = carry
        if l1:
            total += l1.val
            l1 = l1.next
        if l2:
            total += l2.val
            l2 = l2.next
        carry, digit = divmod(total, 10)
        tail.next = ListNode(digit)
        tail = tail.next
    return dummy.next`,
			Explanation: "Builds on Merge Two Sorted Lists (Order 2): walk two lists side by side and build the answer behind a dummy head with a tail pointer. The new ideas are building brand-new nodes instead of relinking old ones, handling lists of different lengths, and carrying state (the carry) from one step to the next, which is grade-school addition.",
			BruteForce:  "Turn each list into a Python integer, add them, and turn the sum back into a list. In Python that works because integers have no size limit, but in most languages a 100-digit number overflows a 64-bit integer. It also does three passes where one would do.",
			KeyInsight:  "Storing digits in reverse puts the ones digit first, which is exactly the order column addition needs. So one left-to-right walk with a carry produces the answer digit by digit, for numbers of any length.",
			WhyItWorks:  "Invariant: after processing k columns, the nodes after dummy hold the lowest k digits of the sum, and carry (0 or 1) is the amount that column k + 1 must add. It holds after column 0 because the digit is total % 10 and the overflow total // 10 is at most 1 (9 + 9 + 1 = 19). A missing digit in the shorter list is treated as 0, which is what a leading zero means. The loop condition includes carry, so a final carry becomes one extra digit (5 + 5 = 10 gives [0,1]). Each column is O(1) work and there are max(m, n) + 1 columns at most, so the time is O(max(m, n)); the output list is required, and the extra space besides it is O(1).",
			CommonMistakes: []string{
				"Looping only while l1 and l2. Digits of the longer number are dropped ([1,8] + [0] would return [1]). Loop while either list has digits, treating a missing digit as 0.",
				"Forgetting the final carry. 5 + 5 then returns [0] instead of [0,1]. Include carry in the loop condition.",
				"Reading l1.val after l1 has become None. Guard each list with if l1 / if l2 before using its digit.",
			},
			TimeComplexity:  "O(max(m, n))",
			SpaceComplexity: "O(1) besides the output list",
			Walkthrough: []WalkthroughStep{
				{Title: "Add a dummy head", Explanation: "As in Merge Two Sorted Lists, the dummy means appending the first digit needs no special case. The carry starts at 0.", CodeSnippet: "dummy = ListNode(0)\ntail = dummy\ncarry = 0", LineStart: 2, LineEnd: 4},
				{Title: "Add one column", Explanation: "Keep going while any digit or carry remains. Missing digits count as 0, so lists of different lengths need no special case.", CodeSnippet: "while l1 or l2 or carry:\n    total = carry\n    if l1:\n        total += l1.val", LineStart: 5, LineEnd: 12},
				{Title: "Split the total", Explanation: "divmod(14, 10) is (1, 4): carry 1 to the next column and write 4 here.", CodeSnippet: "carry, digit = divmod(total, 10)", LineStart: 13, LineEnd: 13},
				{Title: "Append the digit", Explanation: "Create a node for this digit and move the tail onto it.", CodeSnippet: "tail.next = ListNode(digit)\ntail = tail.next", LineStart: 14, LineEnd: 15},
				{Title: "Return the real head", Explanation: "The digits come out lowest first, which is already the required reverse format.", CodeSnippet: "return dummy.next", LineStart: 16, LineEnd: 16},
			},
			Alternatives: []AltSolution{
				{
					Name: "Convert to integers and back",
					Code: `def addTwoNumbers(l1, l2):
    def to_int(node):
        value, place = 0, 1
        while node:
            value += node.val * place
            place *= 10
            node = node.next
        return value

    total = to_int(l1) + to_int(l2)
    dummy = ListNode(0)
    tail = dummy
    while True:
        total, digit = divmod(total, 10)
        tail.next = ListNode(digit)
        tail = tail.next
        if total == 0:
            return dummy.next`,
					TimeComplexity:  "O(max(m, n))",
					SpaceComplexity: "O(1) besides the output",
					Explanation:     "Relies on Python's unlimited-size integers. It shows what the list represents, but in Java or C++ a 100-digit number overflows, which is why the digit-by-digit carry is the expected answer.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can the carry never be more than 1?", Answer: "The largest column total is 9 + 9 + a carry of at most 1, which is 19. 19 // 10 = 1. By induction the carry stays 0 or 1 in every column, starting from 0."},
			{Prompt: "What if the digits were stored most significant first (LeetCode 445, 3 -> 4 -> 2 for 342)? What would you change?", Answer: "Column addition must start from the ones digit, which is now at the end. Reverse both lists first (Order 1), add as here, and reverse the result; or push the digits onto two stacks and pop them, building the answer by inserting each new node at the front."},
			{Prompt: "Why does the loop condition need all three of l1, l2 and carry?", Answer: "l1 and l2 cover numbers of different lengths: while either still has digits there is a column to add. carry covers the case where both lists are finished but the last column overflowed, like 5 + 5, which needs one more node holding 1."},
		},
	},
	{
		ID:              "copy-list-with-random-pointer",
		Number:          138,
		Order:           7,
		Title:           "Copy List with Random Pointer",
		Slug:            "copy-list-with-random-pointer",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Hash Table", "Linked List"},
		RelatedChapters: []int{7, 2},
		Description: `Each node of this linked list has a value, a next pointer, and an extra pointer called random that can point to any node of the same list (including itself) or be None.

Build a deep copy: exactly as many brand-new nodes, with the same values, where every copied next and random pointer points to the corresponding NEW node. No pointer in the copy may point into the original list, and the original list must be left unchanged. Return the head of the copy (None for an empty list).

How the tests show a list: each node is written as [val, random_index], where random_index is the 0-based position of the node that random points to, or null.`,
		Constraints: []string{
			"The list has between 0 and 1000 nodes",
			"-10^4 <= Node.val <= 10^4",
			"Node.random is None or points to a node in the list",
		},
		Examples: []Example{
			{Input: "head = [[7,null],[13,0],[11,4],[10,2],[1,0]]", Output: "[[7,null],[13,0],[11,4],[10,2],[1,0]]", Explanation: "The copy looks identical, e.g. the copied 13's random points to the copied 7 (index 0), but every node in it is new."},
			{Input: "head = [[1,1],[2,1]]", Output: "[[1,1],[2,1]]"},
			{Input: "head = [[3,null],[3,0],[3,null]]", Output: "[[3,null],[3,0],[3,null]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}}, Expected: [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}, Note: "random pointers forwards, backwards and None"},
			{Input: map[string]interface{}{"head": [][]interface{}{{1, 1}, {2, 1}}}, Expected: [][]interface{}{{1, 1}, {2, 1}}, Note: "a node whose random points to itself"},
			{Input: map[string]interface{}{"head": [][]interface{}{{3, nil}, {3, 0}, {3, nil}}}, Expected: [][]interface{}{{3, nil}, {3, 0}, {3, nil}}, Note: "equal values: map nodes, not values"},
			{Input: map[string]interface{}{"head": [][]interface{}{}}, Expected: [][]interface{}{}, Hidden: true, Note: "empty list returns None"},
			{Input: map[string]interface{}{"head": [][]interface{}{{5, 0}}}, Expected: [][]interface{}{{5, 0}}, Hidden: true, Note: "single node pointing to itself"},
			{Input: map[string]interface{}{"head": [][]interface{}{{1, nil}, {2, nil}, {3, nil}}}, Expected: [][]interface{}{{1, nil}, {2, nil}, {3, nil}}, Hidden: true, Note: "no random pointers at all"},
			{Input: map[string]interface{}{"head": [][]interface{}{{-10000, 2}, {0, 2}, {10000, 0}}}, Expected: [][]interface{}{{-10000, 2}, {0, 2}, {10000, 0}}, Hidden: true, Note: "extreme values and two nodes sharing a random target"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `# The Node class is provided:
# class Node:
#     def __init__(self, x, next=None, random=None):
#         self.val = int(x)
#         self.next = next
#         self.random = random
def copyRandomList(head):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"head": "RandomList"},
			ReturnType: "RandomList",
			Prelude:    randomListPrelude,
			Validator:  randomListValidator,
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Copying next pointers is easy because you create the nodes in that order. But a random pointer can point to a node you haven't copied yet. What would you need to look up, given an original node, to set a copy's pointer?"},
			{Level: 2, Type: "insight", Content: "Keep a hash map from each original node to its copy (Chapter 2). First pass: create every copy and record old -> new. Second pass: every pointer is now a lookup, copy.next = map[old.next] and copy.random = map[old.random]. Putting None -> None in the map removes the None checks. See Chapter 7 for the two-pass pattern.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: `1. Map None to None, so a missing pointer copies as missing.
2. Create the copies: walk the list once, creating a new node for each original and storing old -> new.
3. Wire the copies: walk again; for each original, set its copy's next and random to the copies of the original's next and random.
4. Return the copy of head.`},
			{Level: 4, Type: "skeleton", Content: `def copyRandomList(head):
    old_to_new = {None: None}
    node = head
    while node:
        old_to_new[node] = ____
        node = node.next
    node = head
    while node:
        copy = old_to_new[node]
        copy.next = ____
        copy.random = ____
        node = node.next
    return ____`},
		},
		Solution: Solution{
			Code: `def copyRandomList(head):
    old_to_new = {None: None}
    node = head
    while node:
        old_to_new[node] = Node(node.val)
        node = node.next
    node = head
    while node:
        copy = old_to_new[node]
        copy.next = old_to_new[node.next]
        copy.random = old_to_new[node.random]
        node = node.next
    return old_to_new[head]`,
			Explanation: "Adds a hash map (Chapter 2) to the linked-list walks of the earlier rungs. Pointers that can jump anywhere can't be copied in one forward pass, so the work splits into two passes: create all copies, then wire them by looking up each target's copy.",
			BruteForce:  "Copy the next chain first, then for each node find its random target's position by walking the original list from the head, and walk the copy the same number of steps. That is O(n) per random pointer, so O(n^2) overall. The waste is re-walking the list to answer 'where is the copy of this node?', a question a hash map answers in O(1).",
			KeyInsight:  "Every pointer in the copy is 'the copy of whatever the original pointed to'. If you can look up the copy of any original node in O(1), every pointer becomes a single lookup. So create all copies first, then wire them.",
			WhyItWorks:  "After the first pass, old_to_new holds exactly one new node per original node, plus None -> None. In the second pass, for an original node x with copy c, setting c.next = old_to_new[x.next] and c.random = old_to_new[x.random] makes c point to the copies of x's targets (or None), which is the definition of a deep copy. Because the map is keyed by node identity and not by value, nodes with equal values stay distinct. The original's pointers are only read, never written, so it is unchanged. Two passes of O(1) work per node give O(n) time and the map gives O(n) space.",
			CommonMistakes: []string{
				"Setting copy.random = node.random. That points the copy into the original list, which is a shallow copy; the judge rejects it. Always look up the copy: old_to_new[node.random].",
				"Keying the map by value (old_to_new[node.val]). Nodes with equal values ([[3,null],[3,0],[3,null]]) then share one copy. Key by the node object itself.",
				"Trying to set random in the first pass. The random target may not have been copied yet, so the lookup fails. Create all copies before wiring any pointer.",
			},
			TimeComplexity:  "O(n): two passes",
			SpaceComplexity: "O(n) for the map",
			Walkthrough: []WalkthroughStep{
				{Title: "Map None to None", Explanation: "With this entry, old_to_new[x.next] and old_to_new[x.random] work even when the pointer is None, so no if-checks are needed.", CodeSnippet: "old_to_new = {None: None}", LineStart: 2, LineEnd: 2},
				{Title: "Create the copies", Explanation: "One new node per original, values only. Pointers can't be set yet because their targets may not exist.", CodeSnippet: "old_to_new[node] = Node(node.val)\nnode = node.next", LineStart: 3, LineEnd: 6},
				{Title: "Wire the copies", Explanation: "Every copy now exists, so each pointer is a lookup. For 13 (random -> 7): copy13.random = old_to_new[node7] = copy7.", CodeSnippet: "copy = old_to_new[node]\ncopy.next = old_to_new[node.next]\ncopy.random = old_to_new[node.random]", LineStart: 7, LineEnd: 12},
				{Title: "Return the copy of head", Explanation: "For an empty list head is None and the map returns None.", CodeSnippet: "return old_to_new[head]", LineStart: 13, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Interleave copies, O(1) extra space",
					Code: `def copyRandomList(head):
    if head is None:
        return None
    node = head
    while node:
        node.next = Node(node.val, node.next)
        node = node.next.next
    node = head
    while node:
        if node.random is not None:
            node.next.random = node.random.next
        node = node.next.next
    copy_head = head.next
    node = head
    while node:
        copy = node.next
        node.next = copy.next
        copy.next = copy.next.next if copy.next else None
        node = node.next
    return copy_head`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1) besides the copy",
					Explanation:     "Put each copy right after its original (A -> A' -> B -> B'), so the copy of any node x is x.next. Then copy.random = original.random.next, and finally unzip the two lists, restoring the original. It saves the map but changes the original temporarily, so it has to be undone carefully.",
				},
				{
					Name: "Brute force: find positions by walking",
					Code: `def copyRandomList(head):
    copies = []
    node = head
    while node:
        copies.append(Node(node.val))
        node = node.next
    for i in range(len(copies) - 1):
        copies[i].next = copies[i + 1]
    node = head
    i = 0
    while node:
        if node.random is not None:
            pos, probe = 0, head
            while probe is not node.random:
                probe = probe.next
                pos += 1
            copies[i].random = copies[pos]
        node = node.next
        i += 1
    return copies[0] if copies else None`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n)",
					Explanation:     "For every random pointer, walk from the head to find the target's position, then use the copy at that position. Correct, but each lookup costs O(n), which is the work the hash map removes.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the copies be created in a separate pass before any random pointer is set?", Answer: "A random pointer may point forward to a node that hasn't been reached yet. If its copy doesn't exist, there is nothing to point to. Creating every copy first guarantees that every lookup in the second pass succeeds."},
			{Prompt: "Why is the map keyed by the node object and not by its value?", Answer: "Values can repeat, as in [[3,null],[3,0],[3,null]]. Keyed by value, the three nodes would share one copy and the copy would have one node instead of three. Node objects are unique, so keying by the object (its identity) maps each original to its own copy."},
			{Prompt: "How does the same idea copy a graph (Clone Graph, LeetCode 133)?", Answer: "A graph node can have many neighbours and there can be cycles, but the principle is identical: a map from original node to copy. Walk the graph with BFS or DFS; when you meet a node, create its copy if it isn't in the map yet, then append the copies of its neighbours. The map also serves as the visited set, which stops cycles."},
		},
	},
	{
		ID:              "lru-cache",
		Number:          146,
		Order:           8,
		Title:           "LRU Cache",
		Slug:            "lru-cache",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Hash Table", "Linked List", "Design", "Doubly-Linked List"},
		RelatedChapters: []int{7, 2},
		Description: `Design a key-value cache with a fixed capacity that evicts the Least Recently Used key when it is full. Implement the class LRUCache:

- LRUCache(capacity) creates an empty cache that can hold capacity keys (capacity is positive).
- get(key) returns the value stored for key, or -1 if key is not in the cache. A successful get counts as a use of that key.
- put(key, value) stores value for key, replacing any old value; this also counts as a use. If a new key makes the cache hold more than capacity keys, remove the key whose most recent use is the oldest.

Both get and put must run in O(1) average time.

How the tests call it: operations lists the method names (the first is the constructor) and values the arguments for each; the expected list has null for calls that return nothing.`,
		Constraints: []string{
			"1 <= capacity <= 3000",
			"0 <= key <= 10^4",
			"0 <= value <= 10^5",
			"At most 2 * 10^5 calls to get and put",
		},
		Examples: []Example{
			{Input: `["LRUCache","put","put","get","put","get","put","get","get","get"]
[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]`, Output: "[null,null,null,1,null,-1,null,-1,3,4]", Explanation: "After put(1), put(2) and get(1), key 2 is the least recently used, so put(3) evicts 2 (get(2) = -1). Then put(4) evicts 1, the oldest of {1, 3}."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}, "values": [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}}, Expected: []interface{}{nil, nil, nil, 1, nil, -1, nil, -1, 3, 4}, Note: "the example: evictions in recency order"},
			{Input: map[string]interface{}{"operations": []string{"LRUCache", "put", "get", "put", "get", "get"}, "values": [][]int{{1}, {2, 1}, {2}, {3, 2}, {2}, {3}}}, Expected: []interface{}{nil, nil, 1, nil, -1, 2}, Note: "capacity 1: every new key evicts the old one"},
			{Input: map[string]interface{}{"operations": []string{"LRUCache", "put", "put", "put", "get", "get"}, "values": [][]int{{2}, {1, 1}, {2, 2}, {1, 10}, {1}, {2}}}, Expected: []interface{}{nil, nil, nil, nil, 10, 2}, Note: "updating an existing key must not evict anything"},
			{Input: map[string]interface{}{"operations": []string{"LRUCache", "put", "put", "put", "put", "get", "get", "get"}, "values": [][]int{{2}, {1, 1}, {2, 2}, {1, 5}, {3, 3}, {1}, {2}, {3}}}, Expected: []interface{}{nil, nil, nil, nil, nil, 5, -1, 3}, Hidden: true, Note: "put on an existing key refreshes its recency"},
			{Input: map[string]interface{}{"operations": []string{"LRUCache", "put", "put", "get", "put", "get", "get", "get"}, "values": [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {1}, {2}, {3}}}, Expected: []interface{}{nil, nil, nil, 1, nil, 1, -1, 3}, Hidden: true, Note: "get refreshes recency, so the other key is evicted"},
			{Input: map[string]interface{}{"operations": []string{"LRUCache", "get", "put", "get", "get"}, "values": [][]int{{2}, {5}, {5, 50}, {5}, {6}}}, Expected: []interface{}{nil, -1, nil, 50, -1}, Hidden: true, Note: "get on an empty cache and on a missing key"},
			{Input: map[string]interface{}{"operations": []string{"LRUCache", "put", "put", "put", "get", "put", "put", "get", "get", "get", "get"}, "values": [][]int{{3}, {1, 100}, {2, 200}, {3, 300}, {1}, {4, 400}, {5, 500}, {1}, {2}, {3}, {4}}}, Expected: []interface{}{nil, nil, nil, nil, 100, nil, nil, 100, -1, -1, 400}, Hidden: true, Note: "capacity 3 with two evictions in a row"},
		},
		TimeComplexity:  "O(1) per operation",
		SpaceComplexity: "O(capacity)",
		StarterCode: `class LRUCache:
    def __init__(self, capacity):
        pass

    def get(self, key):
        pass

    def put(self, key, value):
        pass`,
		Harness: Harness{
			Mode: "design",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You need two things in O(1): find a key's value, and know which key was used least recently, even after some old key is used again. Which structure gives instant lookup, and which keeps an order you can change in the middle in O(1)?"},
			{Level: 2, Type: "insight", Content: "Combine a hash map (Chapter 2) with a doubly linked list ordered by recency: most recent next to a dummy head, least recent next to a dummy tail. The map stores key -> node, so you can jump straight to a node and unlink it in O(1) using its prev and next pointers. See the LRU storyboard in Chapter 7.", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: `1. Set up the sentinels: dummy head and tail nodes linked to each other, and an empty map.
2. Write two helpers: unlink(node) joins its neighbours; push_front(node) inserts it right after head.
3. get: if the key is missing return -1; otherwise move its node to the front (unlink + push_front) and return its value.
4. put: if the key exists, unlink its old node; create a node, store it in the map and push it to the front.
5. Evict the LRU: if the map is now over capacity, the node before tail is the least recent; unlink it and delete its key from the map (that's why nodes store their key).`},
			{Level: 4, Type: "skeleton", Content: `class DLLNode:
    def __init__(self, key=0, val=0):
        self.key, self.val = key, val
        self.prev = self.next = None

class LRUCache:
    def __init__(self, capacity):
        self.capacity = capacity
        self.nodes = {}
        self.head, self.tail = DLLNode(), DLLNode()
        self.head.next, self.tail.prev = self.tail, self.head

    def _unlink(self, node):
        node.prev.next = ____
        node.next.prev = ____

    def _push_front(self, node):
        node.prev, node.next = self.head, self.head.next
        self.head.next.prev = node
        self.head.next = node

    def get(self, key):
        if key not in self.nodes:
            return -1
        node = self.nodes[key]
        ____   # move it to the front
        return node.val

    def put(self, key, value):
        if key in self.nodes:
            self._unlink(self.nodes[key])
        node = DLLNode(key, value)
        self.nodes[key] = node
        self._push_front(node)
        if len(self.nodes) > self.capacity:
            lru = ____
            self._unlink(lru)
            del self.nodes[____]`},
		},
		Solution: Solution{
			Code: `class DLLNode:
    def __init__(self, key=0, val=0):
        self.key = key
        self.val = val
        self.prev = None
        self.next = None


class LRUCache:
    def __init__(self, capacity):
        self.capacity = capacity
        self.nodes = {}
        self.head = DLLNode()
        self.tail = DLLNode()
        self.head.next = self.tail
        self.tail.prev = self.head

    def _unlink(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev

    def _push_front(self, node):
        node.prev = self.head
        node.next = self.head.next
        self.head.next.prev = node
        self.head.next = node

    def get(self, key):
        if key not in self.nodes:
            return -1
        node = self.nodes[key]
        self._unlink(node)
        self._push_front(node)
        return node.val

    def put(self, key, value):
        if key in self.nodes:
            self._unlink(self.nodes[key])
        node = DLLNode(key, value)
        self.nodes[key] = node
        self._push_front(node)
        if len(self.nodes) > self.capacity:
            lru = self.tail.prev
            self._unlink(lru)
            del self.nodes[lru.key]`,
			Explanation: "A design problem that combines the hash map of Chapter 2 with the pointer rewiring practised in every earlier rung, now on a doubly linked list. The list keeps keys in recency order and the map finds any key's node instantly. The dummy head and tail from Merge Two Sorted Lists (Order 2) come back as sentinels at both ends.",
			BruteForce:  "Keep a dict for values and a Python list of keys in usage order. Each get or put removes the key from the order list (list.remove is O(capacity)) and appends it; eviction pops index 0 (also O(capacity)). With capacity 3000 and 2 * 10^5 calls that is up to 6 * 10^8 element moves. The waste is searching and shifting the array to move one key.",
			KeyInsight:  "Moving a key to the front is O(1) in a doubly linked list if you already hold its node: unlink it using its own prev and next, then insert it after the head. The hash map gives you that node in O(1). Together they make every operation O(1).",
			WhyItWorks:  "Invariant: the list from head to tail holds exactly the cached keys, ordered from most to least recently used, and nodes[key] is the list node for key. get and put both move the key used to the front, which keeps the order right. A new key is pushed to the front; if that pushes the size over capacity, the node just before tail is by the invariant the least recently used, so removing it (from the list and, through node.key, from the map) is the correct eviction. The sentinels guarantee that every real node has non-None prev and next, so unlink and push_front never need a special case. Each operation does a constant number of dictionary operations and pointer changes: O(1) average time, O(capacity) space.",
			CommonMistakes: []string{
				"Not refreshing recency on get. Then get(1) followed by put(3) evicts key 1 instead of key 2. Every successful get must move the node to the front.",
				"Evicting when updating an existing key. put on a key that is already cached doesn't add a key, so the size check must run after the old node is removed (or be skipped for updates); otherwise put(1, 10) on a full cache wrongly evicts another key.",
				"Storing only the value in each node. When the tail node is evicted you need its key to delete it from the map. Store key and value in the node.",
				"Using a singly linked list. Unlinking a node then needs its predecessor, which takes an O(n) walk to find. The prev pointer is what makes removal O(1).",
			},
			TimeComplexity:  "O(1) average per get and put",
			SpaceComplexity: "O(capacity)",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the sentinels", Explanation: "head and tail are dummy nodes that are never removed, so every real node always has two neighbours. Recent keys live next to head, old ones next to tail.", CodeSnippet: "self.nodes = {}\nself.head = DLLNode()\nself.tail = DLLNode()\nself.head.next = self.tail", LineStart: 10, LineEnd: 16},
				{Title: "Unlink and push to the front", Explanation: "unlink joins a node's neighbours to each other; push_front inserts the node right after head. Both are four pointer changes or fewer: O(1).", CodeSnippet: "node.prev.next = node.next\nnode.next.prev = node.prev", LineStart: 18, LineEnd: 26},
				{Title: "Move the key to the front", Explanation: "A hit in get is a use, so its node moves to the most-recent end before the value is returned.", CodeSnippet: "node = self.nodes[key]\nself._unlink(node)\nself._push_front(node)", LineStart: 28, LineEnd: 34},
				{Title: "Insert or update", Explanation: "An existing key's old node is removed first, so an update never counts as a new key. The new node goes in the map and at the front.", CodeSnippet: "if key in self.nodes:\n    self._unlink(self.nodes[key])\nnode = DLLNode(key, value)", LineStart: 36, LineEnd: 41},
				{Title: "Evict the LRU", Explanation: "Over capacity means one key too many. tail.prev is the least recently used; remove it from the list and use its stored key to remove it from the map.", CodeSnippet: "lru = self.tail.prev\nself._unlink(lru)\ndel self.nodes[lru.key]", LineStart: 42, LineEnd: 45},
			},
			Alternatives: []AltSolution{
				{
					Name: "OrderedDict",
					Code: `from collections import OrderedDict


class LRUCache:
    def __init__(self, capacity):
        self.capacity = capacity
        self.data = OrderedDict()

    def get(self, key):
        if key not in self.data:
            return -1
        self.data.move_to_end(key)
        return self.data[key]

    def put(self, key, value):
        if key in self.data:
            self.data.move_to_end(key)
        self.data[key] = value
        if len(self.data) > self.capacity:
            self.data.popitem(last=False)`,
					TimeComplexity:  "O(1) average per operation",
					SpaceComplexity: "O(capacity)",
					Explanation:     "Python's OrderedDict is internally a hash map plus a doubly linked list, exactly the structure above. move_to_end marks a use and popitem(last=False) evicts the oldest. Fine in practice, but interviews usually want you to build the list yourself.",
				},
				{
					Name: "Brute force: dict plus a list of keys",
					Code: `class LRUCache:
    def __init__(self, capacity):
        self.capacity = capacity
        self.values = {}
        self.order = []

    def get(self, key):
        if key not in self.values:
            return -1
        self.order.remove(key)
        self.order.append(key)
        return self.values[key]

    def put(self, key, value):
        if key in self.values:
            self.order.remove(key)
        self.values[key] = value
        self.order.append(key)
        if len(self.order) > self.capacity:
            oldest = self.order.pop(0)
            del self.values[oldest]`,
					TimeComplexity:  "O(capacity) per operation",
					SpaceComplexity: "O(capacity)",
					Explanation:     "Same behaviour, but order.remove and pop(0) shift the array, so each call costs O(capacity). Easy to check correctness against; too slow for the O(1) requirement.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the cache need a doubly linked list rather than a singly linked one?", Answer: "get moves an arbitrary node to the front, which means unlinking it from the middle. In a singly linked list you need the node's predecessor to unlink it, and finding the predecessor takes an O(n) walk. With a prev pointer the node knows its predecessor, so unlinking is O(1)."},
			{Prompt: "Why does each node store its key and not just its value?", Answer: "Eviction starts from the list side: tail.prev is the node to drop. To also remove it from the map you need its key, and the node is the only place to read it in O(1)."},
			{Prompt: "What if the cache had to evict the least FREQUENTLY used key instead (LFU, LeetCode 460)?", Answer: "Recency order is no longer enough. A common O(1) design keeps a map key -> (value, count) and a map count -> doubly linked list of keys with that count (in recency order, to break ties), plus the current minimum count. A use moves the key from list count to list count + 1; eviction takes the least recent key of the minimum-count list."},
		},
	},
	{
		ID:              "merge-k-sorted-lists",
		Number:          23,
		Order:           9,
		Title:           "Merge k Sorted Lists",
		Slug:            "merge-k-sorted-lists",
		Difficulty:      "Hard",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Divide and Conquer", "Heap (Priority Queue)", "Merge Sort"},
		RelatedChapters: []int{7, 8, 11},
		Description: `You are given a Python list lists of k linked lists. Each linked list is sorted in non-decreasing order, and any of them may be empty (None).

Merge all of them into a single sorted linked list and return its head. If there are no nodes at all (k = 0, or every list is empty), return None, shown as [].`,
		Constraints: []string{
			"0 <= k <= 10^4",
			"Each list has between 0 and 500 nodes",
			"-10^4 <= Node.val <= 10^4",
			"Each list is sorted in non-decreasing order",
			"The total number of nodes N is at most 10^4",
		},
		Examples: []Example{
			{Input: "lists = [[1,4,5],[1,3,4],[2,6]]", Output: "[1,1,2,3,4,4,5,6]", Explanation: "Round 1 merges [1,4,5] with [1,3,4] into [1,1,3,4,4,5] and leaves [2,6]; round 2 merges those two."},
			{Input: "lists = []", Output: "[]"},
			{Input: "lists = [[]]", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"lists": [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}}, Expected: []int{1, 1, 2, 3, 4, 4, 5, 6}, Note: "three lists, odd count"},
			{Input: map[string]interface{}{"lists": [][]int{}}, Expected: []int{}, Note: "no lists at all"},
			{Input: map[string]interface{}{"lists": [][]int{{}}}, Expected: []int{}, Note: "one empty list"},
			{Input: map[string]interface{}{"lists": [][]int{{}, {1}}}, Expected: []int{1}, Hidden: true, Note: "empty list next to a non-empty one"},
			{Input: map[string]interface{}{"lists": [][]int{{2}, {1}, {0}}}, Expected: []int{0, 1, 2}, Hidden: true, Note: "single-node lists in reverse order"},
			{Input: map[string]interface{}{"lists": [][]int{{-1, 5, 11}, {}, {6, 10}}}, Expected: []int{-1, 5, 6, 10, 11}, Hidden: true, Note: "negative values and an empty list in the middle"},
			{Input: map[string]interface{}{"lists": [][]int{{1, 1}, {1, 1}}}, Expected: []int{1, 1, 1, 1}, Hidden: true, Note: "ties everywhere: a heap of (val, node) would crash comparing nodes"},
			{Input: map[string]interface{}{"lists": [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}, {0}}}, Expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, Hidden: true, Note: "even number of lists, fully interleaved"},
		},
		TimeComplexity:  "O(N log k)",
		SpaceComplexity: "O(k)",
		StarterCode: `# ListNode(val, next) is provided
def mergeKLists(lists):
    # Write your solution here
    pass`,
		Harness: Harness{
			ParamTypes: map[string]string{"lists": "ListNode[]"},
			ReturnType: "ListNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You already know how to merge two sorted lists in linear time. If you merged the lists one after another into a growing result, how many times would the earliest nodes be walked over? Is there an order of merging that walks each node fewer times?"},
			{Level: 2, Type: "insight", Content: "Merge in pairs, like the levels of merge sort (Chapter 8): merge lists 0+1, 2+3, ... to halve the number of lists, and repeat. Each round walks every node once, and there are only log k rounds, so the total is O(N log k). Your mergeTwoLists from Order 2 is the only tool you need. (A min-heap of the k front nodes, Chapter 11, gives the same bound.)", ChapterRef: 7},
			{Level: 3, Type: "plan", Content: `1. Handle no lists: return None if lists is empty.
2. Merge two lists: reuse the dummy-head merge from Order 2 as a helper.
3. Merge in rounds: while more than one list remains, merge neighbours in pairs (a lone last list is paired with None) to form the next round's lists.
4. Return the single list that remains.`},
			{Level: 4, Type: "skeleton", Content: `def mergeKLists(lists):
    def merge_two(a, b):
        # the dummy-head merge from Merge Two Sorted Lists
        ...

    if not lists:
        return None
    while len(lists) > 1:
        merged = []
        for i in range(0, len(lists), 2):
            second = ____
            merged.append(____)
        lists = ____
    return lists[0]`},
		},
		Solution: Solution{
			Code: `def mergeKLists(lists):
    def merge_two(a, b):
        dummy = ListNode(0)
        tail = dummy
        while a and b:
            if a.val <= b.val:
                tail.next = a
                a = a.next
            else:
                tail.next = b
                b = b.next
            tail = tail.next
        tail.next = a if a else b
        return dummy.next

    if not lists:
        return None
    while len(lists) > 1:
        merged = []
        for i in range(0, len(lists), 2):
            second = lists[i + 1] if i + 1 < len(lists) else None
            merged.append(merge_two(lists[i], second))
        lists = merged
    return lists[0]`,
			Explanation: "The capstone of the ladder: Merge Two Sorted Lists (Order 2) used as a building block inside divide and conquer (Chapter 8). Merging in pairwise rounds is the bottom-up form of merge sort's combine step, so every node is walked only log k times instead of up to k times.",
			BruteForce:  "Two common slow versions. (1) Collect all N values, sort them and build a new list: O(N log N) time and O(N) extra space, ignoring that each list is already sorted. (2) Merge the lists one by one into a growing result: the result is walked again for every list, so the first list's nodes are walked k times, O(N k) in total. The waste in (2) is re-walking the same early nodes over and over.",
			KeyInsight:  "The order of merging matters. Pairing lists up halves their number each round, and a round costs O(N) because every node is in exactly one merge. With log k rounds, each node is walked O(log k) times instead of O(k).",
			WhyItWorks:  "Correctness: merge_two returns a sorted list containing exactly the nodes of its two inputs (Order 2), so after each round the lists in merged are sorted and together hold every node exactly once. The loop ends with one list, which therefore is all N nodes in sorted order. An odd list out is merged with None, which returns it unchanged. Cost: in each round every node is in exactly one merge_two call, and a merge is linear in its two inputs, so a round costs O(N). The number of lists goes k -> ceil(k/2) -> ... -> 1, which is ceil(log2 k) rounds, so the total is O(N log k). Extra space is the merged list of at most k/2 heads, O(k).",
			CommonMistakes: []string{
				"Merging one list at a time into an accumulator. It is correct, but O(N k): with 10^4 single-node lists it does about 5 * 10^7 steps. Merge in pairs instead.",
				"Using a heap of (val, node) tuples. When two values tie, Python compares the ListNode objects and raises TypeError. Add a tie-breaker: push (val, index, node).",
				"Not handling empty input: lists = [] has no lists[0], and lists = [[]] holds a single None. Return None early for an empty Python list; a None inside lists is handled by merge_two.",
			},
			TimeComplexity:  "O(N log k), N = total nodes",
			SpaceComplexity: "O(k) for the list of heads",
			Walkthrough: []WalkthroughStep{
				{Title: "Merge two lists", Explanation: "Exactly the dummy-head merge from Order 2, used here as a helper.", CodeSnippet: "dummy = ListNode(0)\ntail = dummy\nwhile a and b:", LineStart: 2, LineEnd: 14},
				{Title: "Handle no lists", Explanation: "With zero lists there is no lists[0] to return.", CodeSnippet: "if not lists:\n    return None", LineStart: 16, LineEnd: 17},
				{Title: "Merge in rounds", Explanation: "Pair neighbours: on [[1,4,5],[1,3,4],[2,6]] round 1 gives [[1,1,3,4,4,5],[2,6]] (the third list is merged with None), round 2 gives one list.", CodeSnippet: "for i in range(0, len(lists), 2):\n    second = lists[i + 1] if i + 1 < len(lists) else None\n    merged.append(merge_two(lists[i], second))\nlists = merged", LineStart: 18, LineEnd: 23},
				{Title: "Return the merged list", Explanation: "One list is left and it contains every node in sorted order. For [[]] it is None.", CodeSnippet: "return lists[0]", LineStart: 24, LineEnd: 24},
			},
			Alternatives: []AltSolution{
				{
					Name: "Min-heap of the front nodes",
					Code: `import heapq


def mergeKLists(lists):
    heap = []
    for i, node in enumerate(lists):
        if node:
            heapq.heappush(heap, (node.val, i, node))
    dummy = ListNode(0)
    tail = dummy
    while heap:
        _, i, node = heapq.heappop(heap)
        tail.next = node
        tail = node
        if node.next:
            heapq.heappush(heap, (node.next.val, i, node.next))
    return dummy.next`,
					TimeComplexity:  "O(N log k)",
					SpaceComplexity: "O(k) for the heap",
					Explanation:     "Generalises 'compare the two fronts' to k fronts: a min-heap (Chapter 11) holds one front node per list and pops the smallest in O(log k). The index i breaks ties so Python never compares two ListNode objects.",
				},
				{
					Name: "Brute force: collect, sort and rebuild",
					Code: `def mergeKLists(lists):
    values = []
    for node in lists:
        while node:
            values.append(node.val)
            node = node.next
    dummy = ListNode(0)
    tail = dummy
    for value in sorted(values):
        tail.next = ListNode(value)
        tail = tail.next
    return dummy.next`,
					TimeComplexity:  "O(N log N)",
					SpaceComplexity: "O(N)",
					Explanation:     "Throws away the fact that each list is sorted. Simple and a good correctness reference, but slower than O(N log k) when k is much smaller than N.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is merging in pairs O(N log k) while merging one list at a time is O(N k)?", Answer: "One at a time, the growing result is walked again in every merge, so the first list's nodes are walked about k times. In pairs, each round walks every node exactly once and halves the number of lists, so after log k rounds you are done and each node has been walked log k times."},
			{Prompt: "How is this the same as merge sort from Chapter 8?", Answer: "Merge sort splits an array down to single elements and then merges sorted runs in pairs, level by level, doing O(n) work per level over log n levels. Here the k lists are the sorted runs you'd have after the splitting, so only the merge levels remain: log k levels of O(N) work."},
			{Prompt: "When would you prefer the heap version?", Answer: "When the lists arrive as streams and you can't hold them all at once, or when you want the merged output one element at a time (for example the first 100 results). The heap needs only one node per list in memory and produces the next smallest value in O(log k)."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, LinkedListProblems...)
}
