package practice

// LinkedListProblems contains all linked-list category problems
var LinkedListProblems = []*Problem{
	{
		ID:         "reverse-linked-list",
		Number:     23,
		Title:      "Reverse Linked List",
		Difficulty: "Easy",
		Category:   "linked-list",
		Tags:       []string{"Linked List", "Recursion"},
		RelatedChapters: []int{7},
		Description: `Given the head of a singly linked list, reverse the list, and return the reversed list.`,
		Constraints: []string{
			"The number of nodes in the list is the range [0, 5000]",
			"-5000 <= Node.val <= 5000",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4,5]", Output: "[5,4,3,2,1]"},
			{Input: "head = [1,2]", Output: "[2,1]"},
			{Input: "head = []", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}}, Expected: []int{5, 4, 3, 2, 1}},
			{Input: map[string]interface{}{"head": []int{1, 2}}, Expected: []int{2, 1}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def reverseList(head):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use three pointers: prev, curr, next."},
			{Level: 2, Type: "algorithm", Content: "At each step, reverse the link and move all pointers forward."},
		},
		Solution: Solution{
			Code: `def reverseList(head):
    prev = None
    curr = head
    while curr:
        next_temp = curr.next
        curr.next = prev
        prev = curr
        curr = next_temp
    return prev`,
			Explanation: "Iteratively reverse each link. prev tracks new head.",
		},
	},
	{
		ID:         "linked-list-cycle",
		Number:     25,
		Title:      "Linked List Cycle",
		Difficulty: "Easy",
		Category:   "linked-list",
		Tags:       []string{"Hash Table", "Linked List", "Two Pointers"},
		RelatedChapters: []int{7},
		Description: `Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.

Return true if there is a cycle in the linked list. Otherwise, return false.`,
		Constraints: []string{
			"The number of the nodes in the list is in the range [0, 10^4]",
			"-10^5 <= Node.val <= 10^5",
		},
		Examples: []Example{
			{Input: "head = [3,2,0,-4], pos = 1", Output: "true", Explanation: "Tail connects to node index 1."},
			{Input: "head = [1,2], pos = 0", Output: "true"},
			{Input: "head = [1], pos = -1", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{3, 2, 0, -4}, "pos": 1}, Expected: true},
			{Input: map[string]interface{}{"head": []int{1}, "pos": -1}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def hasCycle(head):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use Floyd's cycle detection (slow and fast pointers)."},
			{Level: 2, Type: "algorithm", Content: "If fast catches up to slow, there's a cycle."},
		},
		Solution: Solution{
			Code: `def hasCycle(head):
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            return True
    return False`,
			Explanation: "Floyd's algorithm: slow moves 1 step, fast moves 2 steps. They meet if cycle exists.",
		},
	},
	{
		ID:              "merge-two-sorted-lists",
		Number:          52,
		Title:           "Merge Two Sorted Lists",
		Difficulty:      "Easy",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Recursion"},
		RelatedChapters: []int{2, 7},
		Description: `You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.`,
		Constraints: []string{
			"The number of nodes in both lists is in the range [0, 50]",
			"-100 <= Node.val <= 100",
			"Both list1 and list2 are sorted in non-decreasing order",
		},
		Examples: []Example{
			{Input: "list1 = [1,2,4], list2 = [1,3,4]", Output: "[1,1,2,3,4,4]"},
			{Input: "list1 = [], list2 = []", Output: "[]"},
			{Input: "list1 = [], list2 = [0]", Output: "[0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"list1": []int{1, 2, 4}, "list2": []int{1, 3, 4}}, Expected: []int{1, 1, 2, 3, 4, 4}},
			{Input: map[string]interface{}{"list1": []int{}, "list2": []int{}}, Expected: []int{}},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def mergeTwoLists(list1, list2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a dummy head to simplify edge cases. Compare nodes and link the smaller one.", ChapterRef: 2},
			{Level: 2, Type: "algorithm", Content: "Maintain a tail pointer. At each step, link tail.next to the smaller of the two current nodes, advance that list."},
			{Level: 3, Type: "code", Content: "dummy = ListNode(0); tail = dummy. while l1 and l2: compare and link. Append remaining nodes at end."},
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

    # Append remaining nodes
    tail.next = list1 if list1 else list2

    return dummy.next`,
			Explanation:     "Use dummy head and tail pointer. Compare and link smaller node, advance that list. Append remaining at end.",
			TimeComplexity:  "O(n + m)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Dummy head", Explanation: "Simplifies handling of head node", CodeSnippet: "dummy = ListNode(0)\ntail = dummy", LineStart: 2, LineEnd: 3},
				{Title: "Compare and link", Explanation: "Link smaller node and advance that list", CodeSnippet: "if list1.val <= list2.val:\n    tail.next = list1", LineStart: 6, LineEnd: 8},
				{Title: "Append remaining", Explanation: "One list may have remaining nodes", CodeSnippet: "tail.next = list1 if list1 else list2", LineStart: 15, LineEnd: 15},
			},
		},
	},
	{
		ID:              "reorder-list",
		Number:          53,
		Title:           "Reorder List",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Two Pointers", "Stack", "Recursion"},
		RelatedChapters: []int{2, 7},
		Description: `You are given the head of a singly linked-list. The list can be represented as:

L0 -> L1 -> ... -> Ln - 1 -> Ln

Reorder the list to be on the following form:

L0 -> Ln -> L1 -> Ln - 1 -> L2 -> Ln - 2 -> ...

You may not modify the values in the list's nodes. Only nodes themselves may be changed.`,
		Constraints: []string{
			"The number of nodes in the list is in the range [1, 5 * 10^4]",
			"1 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4]", Output: "[1,4,2,3]"},
			{Input: "head = [1,2,3,4,5]", Output: "[1,5,2,4,3]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4}}, Expected: []int{1, 4, 2, 3}},
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}}, Expected: []int{1, 5, 2, 4, 3}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def reorderList(head):\n    # Write your solution here (modify in place)\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Three steps: find middle, reverse second half, merge two halves alternating.", ChapterRef: 2},
			{Level: 2, Type: "algorithm", Content: "Use slow/fast pointers to find middle. Reverse from middle to end. Interleave first and reversed second halves."},
			{Level: 3, Type: "code", Content: "Find mid with slow/fast. Reverse(slow.next). Merge by alternating: first.next = second, then second.next = first.next."},
		},
		Solution: Solution{
			Code: `def reorderList(head):
    if not head or not head.next:
        return

    # Find middle
    slow, fast = head, head.next
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next

    # Reverse second half
    second = slow.next
    slow.next = None
    prev = None
    while second:
        tmp = second.next
        second.next = prev
        prev = second
        second = tmp

    # Merge two halves
    first, second = head, prev
    while second:
        tmp1, tmp2 = first.next, second.next
        first.next = second
        second.next = tmp1
        first, second = tmp1, tmp2`,
			Explanation:     "Find middle (slow/fast), reverse second half, merge alternating nodes from both halves.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Find middle", Explanation: "Slow pointer ends at middle after fast reaches end", CodeSnippet: "while fast and fast.next:\n    slow = slow.next", LineStart: 6, LineEnd: 9},
				{Title: "Reverse second half", Explanation: "Standard linked list reversal", CodeSnippet: "prev = None\nwhile second:", LineStart: 13, LineEnd: 18},
				{Title: "Merge halves", Explanation: "Interleave nodes from first and second halves", CodeSnippet: "first.next = second\nsecond.next = tmp1", LineStart: 22, LineEnd: 26},
			},
		},
	},
	{
		ID:              "remove-nth-node-from-end",
		Number:          54,
		Title:           "Remove Nth Node From End of List",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Two Pointers"},
		RelatedChapters: []int{2, 7},
		Description: `Given the head of a linked list, remove the nth node from the end of the list and return its head.`,
		Constraints: []string{
			"The number of nodes in the list is sz",
			"1 <= sz <= 30",
			"0 <= Node.val <= 100",
			"1 <= n <= sz",
		},
		Examples: []Example{
			{Input: "head = [1,2,3,4,5], n = 2", Output: "[1,2,3,5]"},
			{Input: "head = [1], n = 1", Output: "[]"},
			{Input: "head = [1,2], n = 1", Output: "[1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": []int{1, 2, 3, 4, 5}, "n": 2}, Expected: []int{1, 2, 3, 5}},
			{Input: map[string]interface{}{"head": []int{1}, "n": 1}, Expected: []int{}},
			{Input: map[string]interface{}{"head": []int{1, 2}, "n": 1}, Expected: []int{1}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def removeNthFromEnd(head, n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use two pointers with n nodes gap. When fast reaches end, slow is at the node before the one to remove."},
			{Level: 2, Type: "algorithm", Content: "Use dummy node to handle edge case of removing head. Advance fast n+1 steps, then move both until fast is null."},
			{Level: 3, Type: "code", Content: "dummy.next = head. Advance fast n+1 times. while fast: move both. slow.next = slow.next.next."},
		},
		Solution: Solution{
			Code: `def removeNthFromEnd(head, n):
    dummy = ListNode(0, head)
    slow = fast = dummy

    # Advance fast n+1 steps ahead
    for _ in range(n + 1):
        fast = fast.next

    # Move both until fast reaches end
    while fast:
        slow = slow.next
        fast = fast.next

    # Remove the node
    slow.next = slow.next.next

    return dummy.next`,
			Explanation:     "Two pointers n+1 apart. When fast reaches null, slow is just before the node to remove. Use dummy to handle head removal.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Use dummy node", Explanation: "Handles edge case of removing head", CodeSnippet: "dummy = ListNode(0, head)", LineStart: 2, LineEnd: 2},
				{Title: "Create gap", Explanation: "Fast is n+1 steps ahead of slow", CodeSnippet: "for _ in range(n + 1):\n    fast = fast.next", LineStart: 6, LineEnd: 7},
				{Title: "Remove node", Explanation: "Slow points to node before target", CodeSnippet: "slow.next = slow.next.next", LineStart: 15, LineEnd: 15},
			},
		},
	},
	{
		ID:              "copy-list-with-random-pointer",
		Number:          55,
		Title:           "Copy List with Random Pointer",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Hash Table", "Linked List"},
		RelatedChapters: []int{2, 5, 7},
		Description: `A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state.

Return the head of the copied linked list.`,
		Constraints: []string{
			"0 <= n <= 1000",
			"-10^4 <= Node.val <= 10^4",
			"Node.random is null or is pointing to some node in the linked list",
		},
		Examples: []Example{
			{Input: "head = [[7,null],[13,0],[11,4],[10,2],[1,0]]", Output: "[[7,null],[13,0],[11,4],[10,2],[1,0]]"},
			{Input: "head = [[1,1],[2,1]]", Output: "[[1,1],[2,1]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"head": [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}}, Expected: [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def copyRandomList(head):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a hash map to map original nodes to their copies. Two passes: create all copies, then set pointers.", ChapterRef: 5},
			{Level: 2, Type: "algorithm", Content: "First pass: create copy nodes and store in map. Second pass: set next and random pointers using the map."},
			{Level: 3, Type: "code", Content: "old_to_new = {None: None}. Create copies. For each old node: copy.next = map[old.next], copy.random = map[old.random]."},
		},
		Solution: Solution{
			Code: `def copyRandomList(head):
    if not head:
        return None

    # Map old nodes to new nodes
    old_to_new = {None: None}

    # First pass: create all copy nodes
    curr = head
    while curr:
        old_to_new[curr] = Node(curr.val)
        curr = curr.next

    # Second pass: set next and random pointers
    curr = head
    while curr:
        copy = old_to_new[curr]
        copy.next = old_to_new[curr.next]
        copy.random = old_to_new[curr.random]
        curr = curr.next

    return old_to_new[head]`,
			Explanation:     "Hash map maps original to copy. First create all copies, then set next/random pointers using the map.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n) for the hash map",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize map", Explanation: "Map None to None for handling null pointers", CodeSnippet: "old_to_new = {None: None}", LineStart: 6, LineEnd: 6},
				{Title: "Create copies", Explanation: "First pass creates all new nodes", CodeSnippet: "old_to_new[curr] = Node(curr.val)", LineStart: 10, LineEnd: 11},
				{Title: "Set pointers", Explanation: "Second pass links next and random using map", CodeSnippet: "copy.next = old_to_new[curr.next]\ncopy.random = old_to_new[curr.random]", LineStart: 17, LineEnd: 19},
			},
		},
	},
	{
		ID:              "add-two-numbers",
		Number:          56,
		Title:           "Add Two Numbers",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Math", "Recursion"},
		RelatedChapters: []int{2, 7},
		Description: `You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.`,
		Constraints: []string{
			"The number of nodes in each linked list is in the range [1, 100]",
			"0 <= Node.val <= 9",
			"It is guaranteed that the list represents a number that does not have leading zeros",
		},
		Examples: []Example{
			{Input: "l1 = [2,4,3], l2 = [5,6,4]", Output: "[7,0,8]", Explanation: "342 + 465 = 807"},
			{Input: "l1 = [0], l2 = [0]", Output: "[0]"},
			{Input: "l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]", Output: "[8,9,9,9,0,0,0,1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"l1": []int{2, 4, 3}, "l2": []int{5, 6, 4}}, Expected: []int{7, 0, 8}},
			{Input: map[string]interface{}{"l1": []int{0}, "l2": []int{0}}, Expected: []int{0}},
			{Input: map[string]interface{}{"l1": []int{9, 9, 9, 9, 9, 9, 9}, "l2": []int{9, 9, 9, 9}}, Expected: []int{8, 9, 9, 9, 0, 0, 0, 1}},
		},
		TimeComplexity:  "O(max(m, n))",
		SpaceComplexity: "O(max(m, n))",
		StarterCode:     "def addTwoNumbers(l1, l2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Simulate addition digit by digit, tracking carry. Handle different lengths and final carry."},
			{Level: 2, Type: "algorithm", Content: "At each step: sum = l1.val + l2.val + carry. New digit = sum % 10. New carry = sum // 10."},
			{Level: 3, Type: "code", Content: "Use dummy head. while l1 or l2 or carry: sum and create new node. Don't forget final carry!"},
		},
		Solution: Solution{
			Code: `def addTwoNumbers(l1, l2):
    dummy = ListNode(0)
    curr = dummy
    carry = 0

    while l1 or l2 or carry:
        val1 = l1.val if l1 else 0
        val2 = l2.val if l2 else 0

        total = val1 + val2 + carry
        carry = total // 10
        curr.next = ListNode(total % 10)
        curr = curr.next

        l1 = l1.next if l1 else None
        l2 = l2.next if l2 else None

    return dummy.next`,
			Explanation:     "Add digit by digit with carry. Handle different lengths by treating missing digits as 0. Don't forget final carry.",
			TimeComplexity:  "O(max(m, n))",
			SpaceComplexity: "O(max(m, n))",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle different lengths", Explanation: "Use 0 if list is exhausted", CodeSnippet: "val1 = l1.val if l1 else 0", LineStart: 7, LineEnd: 8},
				{Title: "Calculate digit and carry", Explanation: "digit = total % 10, carry = total // 10", CodeSnippet: "carry = total // 10\ncurr.next = ListNode(total % 10)", LineStart: 11, LineEnd: 12},
				{Title: "Continue while carry", Explanation: "Loop continues if there's a carry to process", CodeSnippet: "while l1 or l2 or carry:", LineStart: 6, LineEnd: 6},
			},
		},
	},
	{
		ID:              "merge-k-sorted-lists",
		Number:          57,
		Title:           "Merge k Sorted Lists",
		Difficulty:      "Hard",
		Category:        "linked-list",
		Tags:            []string{"Linked List", "Divide and Conquer", "Heap", "Merge Sort"},
		RelatedChapters: []int{2, 7, 11},
		Description: `You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.`,
		Constraints: []string{
			"k == lists.length",
			"0 <= k <= 10^4",
			"0 <= lists[i].length <= 500",
			"-10^4 <= lists[i][j] <= 10^4",
			"lists[i] is sorted in ascending order",
			"The sum of lists[i].length will not exceed 10^4",
		},
		Examples: []Example{
			{Input: "lists = [[1,4,5],[1,3,4],[2,6]]", Output: "[1,1,2,3,4,4,5,6]"},
			{Input: "lists = []", Output: "[]"},
			{Input: "lists = [[]]", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"lists": [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}}, Expected: []int{1, 1, 2, 3, 4, 4, 5, 6}},
			{Input: map[string]interface{}{"lists": [][]int{}}, Expected: []int{}},
		},
		TimeComplexity:  "O(N log k)",
		SpaceComplexity: "O(k)",
		StarterCode:     "def mergeKLists(lists):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a min-heap to efficiently get the smallest element among k list heads. Or use divide and conquer with pairwise merge.", ChapterRef: 11},
			{Level: 2, Type: "algorithm", Content: "Heap approach: push all heads to heap. Pop smallest, add to result, push its next. Repeat until heap empty."},
			{Level: 3, Type: "code", Content: "heapq with (node.val, index, node) tuples. Index breaks ties. Pop min, push next if exists."},
		},
		Solution: Solution{
			Code: `import heapq

def mergeKLists(lists):
    dummy = ListNode(0)
    curr = dummy
    heap = []

    # Initialize heap with heads of all lists
    for i, lst in enumerate(lists):
        if lst:
            heapq.heappush(heap, (lst.val, i, lst))

    while heap:
        val, i, node = heapq.heappop(heap)
        curr.next = node
        curr = curr.next

        if node.next:
            heapq.heappush(heap, (node.next.val, i, node.next))

    return dummy.next`,
			Explanation:     "Min-heap keeps track of smallest among k list heads. Pop smallest, add to result, push its next node.",
			TimeComplexity:  "O(N log k) - N total nodes, log k for each heap operation",
			SpaceComplexity: "O(k) for the heap",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize heap", Explanation: "Add all non-null list heads", CodeSnippet: "for i, lst in enumerate(lists):\n    if lst:\n        heapq.heappush(heap, (lst.val, i, lst))", LineStart: 9, LineEnd: 11},
				{Title: "Pop and link", Explanation: "Get smallest, link to result", CodeSnippet: "val, i, node = heapq.heappop(heap)\ncurr.next = node", LineStart: 14, LineEnd: 16},
				{Title: "Push next", Explanation: "If popped node has next, push it", CodeSnippet: "if node.next:\n    heapq.heappush(heap, (node.next.val, i, node.next))", LineStart: 18, LineEnd: 19},
			},
		},
	},
	{
		ID:              "lru-cache",
		Number:          58,
		Title:           "LRU Cache",
		Difficulty:      "Medium",
		Category:        "linked-list",
		Tags:            []string{"Hash Table", "Linked List", "Design", "Doubly-Linked List"},
		RelatedChapters: []int{2, 5, 7, 12},
		Description: `Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:
- LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
- int get(int key) Return the value of the key if the key exists, otherwise return -1.
- void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.

The functions get and put must each run in O(1) average time complexity.`,
		Constraints: []string{
			"1 <= capacity <= 3000",
			"0 <= key <= 10^4",
			"0 <= value <= 10^5",
			"At most 2 * 10^5 calls will be made to get and put",
		},
		Examples: []Example{
			{Input: `["LRUCache","put","put","get","put","get","put","get","get","get"]
[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]`, Output: "[null,null,null,1,null,-1,null,-1,3,4]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}, "values": [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}}, Expected: []interface{}{nil, nil, nil, 1, nil, -1, nil, -1, 3, 4}},
		},
		TimeComplexity:  "O(1)",
		SpaceComplexity: "O(capacity)",
		StarterCode: `class LRUCache:
    def __init__(self, capacity: int):
        pass

    def get(self, key: int) -> int:
        pass

    def put(self, key: int, value: int) -> None:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use hash map + doubly linked list. Map gives O(1) lookup. DLL gives O(1) remove/add at ends.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "DLL maintains usage order (MRU at head, LRU at tail). On access, move node to head. On capacity overflow, remove tail."},
			{Level: 3, Type: "code", Content: "Use dummy head and tail nodes. Map stores key -> node. get: move to head. put: add/update and move to head, evict if needed."},
		},
		Solution: Solution{
			Code: `class Node:
    def __init__(self, key=0, val=0):
        self.key = key
        self.val = val
        self.prev = None
        self.next = None

class LRUCache:
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.cache = {}  # key -> Node
        # Dummy head and tail
        self.head = Node()
        self.tail = Node()
        self.head.next = self.tail
        self.tail.prev = self.head

    def _remove(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev

    def _add_to_head(self, node):
        node.next = self.head.next
        node.prev = self.head
        self.head.next.prev = node
        self.head.next = node

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        node = self.cache[key]
        self._remove(node)
        self._add_to_head(node)
        return node.val

    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            node = self.cache[key]
            node.val = value
            self._remove(node)
            self._add_to_head(node)
        else:
            node = Node(key, value)
            self.cache[key] = node
            self._add_to_head(node)
            if len(self.cache) > self.capacity:
                # Remove LRU (tail.prev)
                lru = self.tail.prev
                self._remove(lru)
                del self.cache[lru.key]`,
			Explanation:     "Hash map for O(1) lookup, doubly linked list for O(1) add/remove. MRU at head, LRU at tail. Move to head on access.",
			TimeComplexity:  "O(1) for both get and put",
			SpaceComplexity: "O(capacity)",
			Walkthrough: []WalkthroughStep{
				{Title: "Data structures", Explanation: "HashMap + Doubly linked list with dummy nodes", CodeSnippet: "self.cache = {}\nself.head = Node()\nself.tail = Node()", LineStart: 11, LineEnd: 14},
				{Title: "Move to head on access", Explanation: "Remove from current position, add to head", CodeSnippet: "self._remove(node)\nself._add_to_head(node)", LineStart: 31, LineEnd: 32},
				{Title: "Evict LRU", Explanation: "Remove tail.prev when over capacity", CodeSnippet: "lru = self.tail.prev\nself._remove(lru)\ndel self.cache[lru.key]", LineStart: 46, LineEnd: 49},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, LinkedListProblems...)
}
