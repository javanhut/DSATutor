package chapter

import "time"

// chapter07 teaches linked lists: pointer rewiring, the dummy head, fast and
// slow pointers, and the hash map + doubly linked list combination behind an
// LRU cache. Storyboards (IDs registered in app.js): linked-list,
// two-pointers, lru-cache.
func chapter07() Chapter {
	return Chapter{
		Number:        7,
		Slug:          "linked-list",
		Title:         "Linked Lists",
		Category:      "linked-list",
		Prerequisites: []int{2, 3},
		Summary: "A linked list is a chain of nodes joined by pointers. You give up O(1) indexing and gain O(1) insertion and removal " +
			"anywhere you already hold a node. You learn to rewire pointers without losing the rest of the list (save, rewire, advance), " +
			"to remove first-node special cases with a dummy head, to find middles, cycles and n-th-from-the-end nodes with fast and " +
			"slow pointers, and to pair a doubly linked list with a hash map for O(1) caches.",
		Motivation: "Suppose a text editor keeps the lines of a document in a Python list and the user pastes 100,000 lines, each at the " +
			"top of the file. lines.insert(0, x) has to shift every existing line one slot to the right, so the inserts cost " +
			"1 + 2 + ... + 100,000 = 5,000,050,000 element moves, about 5 billion. A linked list does each insert by changing two " +
			"pointers: 100,000 inserts, 200,000 pointer writes. The price is that you can no longer jump to line 50,000; you have to " +
			"walk 50,000 nodes to get there. Linked-list problems are therefore about doing everything in one or two walks and " +
			"rewiring pointers in the right order, because one wrong assignment can cut off the rest of the list or create a loop. " +
			"The same trade-off powers real systems: an LRU cache, used in browsers and databases, moves an item to the front of a " +
			"linked list in O(1) every time it is used.",
		Objectives: []string{
			"Trace pointer assignments on a small list by drawing the nodes and arrows before and after each line",
			"Implement in-place reversal with the subgoals save the next node, rewire the pointer, advance the pointers",
			"Use a dummy head so that inserting or deleting the first node needs no special case",
			"Explain why fast and slow pointers find the middle of a list and why they must meet inside a cycle",
			"Combine a hash map with a doubly linked list to support lookup, move-to-front and eviction in O(1)",
		},
		WhenToUse: []string{
			"The input is the head of a linked list, or the task says \"return the head of the modified list\"",
			"\"In place\" or \"O(1) extra memory\" on a list: rewire the existing nodes instead of copying them into an array",
			"\"middle\", \"cycle\", \"n-th node from the end\": two pointers at different speeds or with a fixed gap",
			"The output list is built node by node (merging, adding digits): a dummy head plus a tail pointer",
			"A design needs O(1) lookup AND O(1) reordering or removal in the middle of an order (LRU cache): hash map + doubly linked list",
		},
		Concepts: []Concept{
			{
				Name:        "Nodes, pointers and safe rewiring",
				Description: "A node stores a value and a next pointer. The list is just the head pointer; everything else is reached by following next. Changing a list means changing next pointers, in an order that never loses a node.",
				WhyItWorks: "A singly linked list has exactly one route to every node: from the head, through the next pointers. So the only " +
					"thing that keeps a node reachable is some variable or some next pointer referring to it. When you overwrite curr.next, " +
					"the old target is reachable only if another variable still holds it. That is why every rewiring step starts by saving " +
					"what you are about to overwrite (nxt = curr.next). In-place reversal keeps the invariant \"prev heads the reversed part, " +
					"curr heads the untouched part\", moving one node across per step, so after n steps the whole list is reversed in O(n) " +
					"time with O(1) extra memory.",
				Intuition: "Think of a conga line where each dancer holds the shoulder of the next one. To turn one dancer around you must " +
					"first note who was in front of them (save the next node), then let them grab the person behind instead (rewire the " +
					"pointer), then step to the next dancer (advance). Dancer = node, hand on shoulder = next pointer, noting who was in " +
					"front = the saved nxt variable.",
				CoreIdeas: []string{
					"Access is O(n): reaching the k-th node takes k hops, so solutions aim for one or two passes",
					"Insert or delete after a node you hold is O(1): change one or two pointers, nothing shifts",
					"Before overwriting a next pointer, save the node it points to, or that part of the list is lost",
					"Compare nodes with is (identity), not values; different nodes can hold equal values",
					"Always ask: what happens on an empty list, a one-node list, and at the last node (next is None)?",
				},
				CommonMistakes: []string{
					"Overwriting curr.next before saving it, which cuts off the rest of the list. Save nxt = curr.next first, then rewire.",
					"Accessing node.next.val without checking node.next is not None, which crashes at the last node. Put the None check in the loop condition.",
					"Losing the head: walking with the head variable itself and then returning it. Walk with a separate pointer, or return dummy.next.",
				},
				Examples: []CodeExample{
					{
						ID:       "reverse_list_iterative",
						Title:    "Reverse a list in place: save, rewire, advance",
						Language: "python",
						Snippet: `def reverse_list(head):
    prev = None
    curr = head
    while curr:
        nxt = curr.next     # save the next node
        curr.next = prev    # rewire the pointer
        prev = curr         # advance the pointers
        curr = nxt
    return prev             # return the new head

# [1,2,3]: prev=None curr=1 -> prev=1 curr=2 -> prev=2 curr=3 -> prev=3 curr=None
# result: 3 -> 2 -> 1 -> None`,
						Notes: "Three pointers, one pass, O(1) extra memory. The comments are the subgoal labels used in every linked-list problem.",
					},
					{
						ID:       "linked_list_insert",
						Title:    "Insert after a node",
						Language: "python",
						Snippet: `def insert_after(node, val):
    node.next = ListNode(val, node.next)   # the new node takes over the old link`,
						Notes: "O(1): the new node's next is set to the old successor before node.next is overwritten, so nothing is lost.",
					},
				},
			},
			{
				Name:        "The dummy head",
				Description: "A throwaway node placed in front of the real head. Every real node now has a predecessor, so deleting or inserting at the front looks exactly like doing it anywhere else. Return dummy.next at the end.",
				WhyItWorks: "Most list edits need the node BEFORE the one you change: to delete x you set prev.next = x.next. The first node " +
					"has no predecessor, so code without a dummy needs a separate branch for it, and that branch is where bugs hide (removing " +
					"the head, building a list whose first node is not known yet). A dummy node gives the head a predecessor, so one code " +
					"path covers every position. The real head may change during the algorithm, but dummy.next always points at the current " +
					"real head, so returning it is always correct.",
				Intuition: "It's like the empty hook at the front of a key ring. You never take it off, so adding or removing the first real key " +
					"works exactly like adding or removing any other. Hook = dummy node, keys = real nodes, the key after the hook = dummy.next, the real head.",
				CoreIdeas: []string{
					"dummy = ListNode(0, head) and return dummy.next",
					"When building a new list, keep a tail pointer to the last node and attach with tail.next = node; tail = tail.next",
					"When deleting, walk with prev (starting at dummy) so you always hold the predecessor",
					"The dummy's value is never read, so any value works",
				},
				CommonMistakes: []string{
					"Returning head instead of dummy.next. If the old head was removed, head still points at the deleted node. Return dummy.next.",
					"Forgetting to advance tail after attaching a node, so every attachment overwrites dummy.next. Write tail = tail.next right after tail.next = ...",
					"Advancing prev after deleting a node. After prev.next = curr.next, prev is still the correct predecessor of the next candidate; move only curr.",
				},
				Examples: []CodeExample{
					{
						ID:       "dummy_head_merge",
						Title:    "Build a list behind a dummy head",
						Language: "python",
						Snippet: `def merge(a, b):
    dummy = ListNode(0)
    tail = dummy
    while a and b:
        if a.val <= b.val:
            tail.next, a = a, a.next
        else:
            tail.next, b = b, b.next
        tail = tail.next
    tail.next = a if a else b
    return dummy.next`,
						Notes: "No 'if the result is still empty' branch anywhere: the dummy is always there to attach to.",
					},
				},
			},
			{
				Name:        "Fast and slow pointers",
				Description: "Two pointers walk the same list in the same direction, either at different speeds (slow 1 step, fast 2) or with a fixed gap. They find the middle, detect cycles and locate the n-th node from the end in one pass with O(1) memory.",
				WhyItWorks: "Middle: fast covers twice the distance of slow, so when fast reaches the end (n nodes), slow has covered n/2: the middle. " +
					"Cycle (Floyd): once both pointers are inside a cycle of length L, look at how many nodes fast is behind slow. Each step " +
					"fast moves 2 and slow 1, so that gap shrinks by exactly 1. It starts below L, so within L steps it reaches 0 and they " +
					"stand on the same node; it can't jump over 0 because it only ever drops by 1. Without a cycle fast reaches None. " +
					"Gap: if lead starts k nodes ahead of trail and both move one step at a time, the gap stays k, so when lead reaches the " +
					"end, trail is exactly k nodes from it.",
				Intuition: "Two runners on a track, one twice as fast. On a straight road the fast one reaches the finish while the slow one " +
					"is halfway (middle of the list). On a circular track the fast one laps the slow one and they must meet (a cycle). " +
					"Runner = pointer, track = list, finish line = None, circular track = cycle.",
				CoreIdeas: []string{
					"Loop condition while fast and fast.next guards both of fast's hops",
					"Starting both at head, slow ends on the middle node (the second middle for even length)",
					"Meeting means a cycle; reaching None means no cycle",
					"A fixed gap of n + 1 starting from a dummy parks the trailing pointer just before the n-th node from the end",
					"Unlike Chapter 3's pointers, these never move backwards: a singly linked list has no prev pointer",
				},
				CommonMistakes: []string{
					"Writing while fast.next only: it crashes with AttributeError when fast is already None. Check fast and fast.next.",
					"Comparing values (slow.val == fast.val) instead of nodes (slow is fast), which reports false cycles in lists with repeated values.",
					"Checking slow is fast before the first move: they start equal, so every list looks cyclic. Move, then compare.",
				},
				Examples: []CodeExample{
					{
						ID:       "find_middle",
						Title:    "Middle of the list and cycle detection",
						Language: "python",
						Snippet: `def middle(head):
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    return slow            # [1,2,3,4,5] -> 3, [1,2,3,4] -> 3

def has_cycle(head):
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True
    return False`,
						Notes: "Same loop, different question: middle reads slow when the loop ends; cycle detection checks for a meeting inside the loop.",
					},
				},
			},
			{
				Name:        "Doubly linked list + hash map",
				Description: "Each node also points to its previous node. Given a node, you can unlink it in O(1) without searching for its predecessor. A hash map from key to node lets you find that node in O(1), which is the design behind the LRU cache.",
				WhyItWorks: "Removing a node needs its predecessor. In a singly linked list finding it is an O(n) walk; with a prev pointer the " +
					"node carries its predecessor, so node.prev.next = node.next and node.next.prev = node.prev remove it in O(1). A hash " +
					"map (Chapter 2) answers \"where is the node for key k?\" in O(1). Together: look up the node, unlink it, push it to the " +
					"front, all O(1). With dummy head and tail sentinels every real node always has both neighbours, so unlink and insert " +
					"never need None checks. The list order then records recency, and the node just before the tail is the least recently used.",
				Intuition: "A stack of papers on a desk, most recently touched on top, plus an index card telling you exactly where each paper is. " +
					"Pulling a paper out and putting it on top is instant because the index card says where it is and the papers above and " +
					"below just close the gap. Papers = nodes, index card = hash map, top of the stack = front of the list, bottom = the LRU item.",
				CoreIdeas: []string{
					"Unlink: node.prev.next = node.next; node.next.prev = node.prev",
					"Insert after head: set the node's prev and next first, then update the neighbours",
					"Sentinel head and tail nodes remove every edge case at both ends",
					"Store the key inside the node, so evicting from the list side can also delete from the map",
				},
				CommonMistakes: []string{
					"Updating head.next before reading head.next to set the new node's next, which loses the old first node. Set the new node's pointers first.",
					"Removing a node from the list but not from the map (or the reverse), so the two structures disagree. Every eviction updates both.",
					"Forgetting that a read (get) also changes recency. Every access moves the node to the front.",
				},
				Examples: []CodeExample{
					{
						ID:       "dll_unlink_push",
						Title:    "O(1) unlink and push to front",
						Language: "python",
						Snippet: `def unlink(node):
    node.prev.next = node.next
    node.next.prev = node.prev

def push_front(head, node):      # head is a sentinel
    node.prev = head
    node.next = head.next
    head.next.prev = node
    head.next = node`,
						Notes: "Four pointer writes at most, whatever the list length. The sentinel guarantees head.next is never None.",
					},
				},
			},
		},
		Animations: []Storyboard{
			{
				ID:    "linked-list",
				Title: "Walking a Linked List",
				Goal:  "Search for 5 in 1 -> 3 -> 5 -> 7 -> 9 by following next pointers, and see why access costs O(n).",
				MemoryTips: []string{
					"Start at the head with a separate pointer (current = head), so you never lose the start",
					"Look at one node at a time: compare current.val, then follow current.next",
					"Stop when you find it or when current is None, the end of the list",
					"Reaching the k-th node costs k hops; there is no jumping to an index",
				},
				Steps: []StoryboardStep{
					{Cue: "start", Narration: "The list is only its head pointer. Put a separate pointer, current, on the head (value 1), so the head itself is never moved or lost.", VisualHint: "Pointer at head", Duration: Duration(2 * time.Second), CodeRef: "linked_list: current = head"},
					{Cue: "compare", Narration: "Compare the current node's value with the target 5. 1 is not 5, and nothing tells us where 5 is, so we must keep walking.", VisualHint: "Highlight comparison", Duration: Duration(2 * time.Second), CodeRef: "linked_list: if current.val == target"},
					{Cue: "advance", Narration: "Follow the next pointer: current = current.next. This is the only way to move; there is no index to jump to. 1 -> 3 -> 5 takes two hops.", VisualHint: "Follow pointer", Duration: Duration(2 * time.Second), CodeRef: "linked_list: current = current.next"},
					{Cue: "found", Narration: "Repeat until the value matches (found 5 at the third node) or current becomes None (not in the list). In the worst case every node is visited: O(n).", VisualHint: "Loop indicator", Duration: Duration(2 * time.Second), CodeRef: "linked_list: while current"},
				},
			},
			{
				ID:    "two-pointers",
				Title: "Floyd's Cycle Detection",
				Goal:  "Watch slow (1 step) and fast (2 steps) on 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> back to 3, and see why fast must land on slow.",
				MemoryTips: []string{
					"Start both pointers at head",
					"Check fast can move: while fast and fast.next",
					"Move slow one, fast two, THEN compare",
					"Check for a meeting with is: same node means a cycle; fast reaching None means no cycle",
					"Inside the cycle the gap shrinks by exactly one per step, so fast can't jump over slow",
				},
				Steps: []StoryboardStep{
					{Cue: "init", Narration: "Both pointers start at the head (1). They start equal, so we must move before checking for a meeting.", VisualHint: "Two pointers at start", Duration: Duration(2 * time.Second), CodeRef: "cycle: slow = fast = head"},
					{Cue: "move", Narration: "Each step slow moves one node and fast moves two: slow 2, 3, 4, 5 while fast goes 3, 5, 3, 5. Once both are in the loop 3-4-5-6 (length 4), fast gains exactly one node on slow per step.", VisualHint: "Different speeds", Duration: Duration(2 * time.Second), CodeRef: "cycle: slow.next, fast.next.next"},
					{Cue: "check", Narration: "After every move, compare the nodes themselves with is. A gap that shrinks by one each step must reach zero; it can't skip from 1 to -1.", VisualHint: "Collision detection", Duration: Duration(2 * time.Second), CodeRef: "cycle: if slow == fast"},
					{Cue: "cycle", Narration: "After 4 steps both stand on node 5: a cycle exists. Without a cycle, fast would have hit None after about n/2 steps and we'd return False. O(n) time, O(1) memory.", VisualHint: "Cycle found", Duration: Duration(2 * time.Second), CodeRef: "cycle: return True"},
				},
			},
			{
				ID:    "lru-cache",
				Title: "LRU Cache: Hash Map + Doubly Linked List",
				Goal:  "With capacity 2, run put(1), put(2), get(1), put(3), get(2), put(4) and watch the recency order and the evictions.",
				MemoryTips: []string{
					"The map finds the node, the list keeps the order: most recent next to head, least recent next to tail",
					"Every use (get or put) moves the key to the front: unlink, then push to the front",
					"Over capacity? Evict tail.prev and delete its key from the map too",
					"Sentinel head and tail mean no None checks when unlinking or inserting",
				},
				Steps: []StoryboardStep{
					{Cue: "put", Narration: "put(1,1), put(2,2): each new key gets a node at the front and a map entry. Order (most recent first): 2, 1.", VisualHint: "Order list bumps key to front", Duration: Duration(2 * time.Second), CodeRef: "put"},
					{Cue: "evict", Narration: "After get(1) the order is 1, 2. put(3,3) makes three keys with capacity 2, so we evict tail.prev, key 2, from the list AND the map. Order: 3, 1.", VisualHint: "Drop tail node", Duration: Duration(2 * time.Second), CodeRef: "evict"},
					{Cue: "get", Narration: "Reads count as uses: get(1) earlier moved 1 to the front, which is why 2 was evicted instead of 1. get(2) now returns -1, and put(4,4) evicts 1, the least recent of 3 and 1.", VisualHint: "Move hit to front", Duration: Duration(2 * time.Second), CodeRef: "get"},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Reverse a linked list in place",
				Problem: "Given the head of [1,2,3], reverse the list in place and return the new head (3 -> 2 -> 1). Use O(1) extra memory.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Set up the pointers",
						Explanation: "Nothing is reversed yet, so prev = None. That None will become the next of the old head, which is right: the old head ends up as the tail. curr starts on the head.",
						Code:        "prev = None\ncurr = head",
						State:       "prev=None  curr=1  list: 1->2->3->None",
					},
					{
						Subgoal:     "Save the next node",
						Explanation: "We are about to overwrite curr.next. It is the only route to 2 and 3, so keep it in nxt first.",
						Code:        "nxt = curr.next",
						State:       "curr=1  nxt=2",
					},
					{
						Subgoal:     "Rewire the pointer",
						Explanation: "Point the current node backwards, at the reversed part. The rest of the list is still safe in nxt.",
						Code:        "curr.next = prev",
						State:       "reversed: 1->None   untouched: 2->3->None (held by nxt)",
					},
					{
						Subgoal:     "Advance the pointers",
						Explanation: "curr is now the front of the reversed part, so it becomes prev; the saved node becomes curr. Repeat the save / rewire / advance steps while curr is not None.",
						Code:        "prev = curr\ncurr = nxt",
						State:       "after node 1: prev=1 curr=2 | after node 2: prev=2 (2->1) curr=3 | after node 3: prev=3 (3->2->1) curr=None",
					},
					{
						Subgoal:     "Return the new head",
						Explanation: "curr is None, so every node is flipped. prev is the old tail, now the head. For an empty list the loop never runs and prev = None is also correct.",
						Code:        "return prev",
						State:       "return 3: 3->2->1->None",
					},
				},
				Code: `def reverse_list(head):
    prev = None
    curr = head
    while curr:
        nxt = curr.next
        curr.next = prev
        prev = curr
        curr = nxt
    return prev`,
			},
			{
				Title:   "Find the middle node with fast and slow pointers",
				Problem: "Return the middle node of [1,2,3,4,5] (node 3) in one pass with O(1) memory. For an even length such as [1,2,3,4], return the second middle (node 3).",
				Steps: []WorkedStep{
					{
						Subgoal:     "Start both pointers at head",
						Explanation: "Both runners start on the first node. The list length is unknown, and we never compute it.",
						Code:        "slow = head\nfast = head",
						State:       "slow=1 fast=1",
					},
					{
						Subgoal:     "Check fast can move",
						Explanation: "fast needs two hops, so fast and fast.next must both exist. Checking fast first protects the fast.next access.",
						Code:        "while fast and fast.next:",
						State:       "fast=1, fast.next=2: continue",
					},
					{
						Subgoal:     "Move slow one, fast two",
						Explanation: "fast covers twice the distance of slow, so when fast has covered the whole list, slow has covered half of it.",
						Code:        "    slow = slow.next\n    fast = fast.next.next",
						State:       "step 1: slow=2 fast=3 | step 2: slow=3 fast=5 | fast.next is None: stop",
					},
					{
						Subgoal:     "Read the answer from slow",
						Explanation: "When fast can't move any further, slow is at the middle. On [1,2,3,4], fast goes 1 -> 3 -> None while slow goes 1 -> 2 -> 3, the second middle.",
						Code:        "return slow",
						State:       "[1,2,3,4,5] -> 3   [1,2,3,4] -> 3",
					},
				},
				Code: `def middle_node(head):
    slow = head
    fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    return slow`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Remove every node with a given value",
				Prompt: "Remove all nodes whose value equals val and return the new head (LeetCode 203). [6,1,6,2,6] with val = 6 becomes [1,2]; [7,7] with val = 7 becomes []. Fill in the blanks using the subgoals: add a dummy head, rewire the pointer (to skip a node), advance the pointers, return the real head.",
				Code: `def removeElements(head, val):
    dummy = ListNode(0, ____)
    prev = dummy
    curr = head
    while curr:
        if curr.val == val:
            prev.next = ____
        else:
            prev = ____
        curr = curr.next
    return ____`,
				Blanks:      []string{"head", "curr.next", "curr", "dummy.next"},
				Explanation: "The dummy gives the first node a predecessor, so deleting a leading 6 is the same as deleting any other. To delete curr, its predecessor skips over it: prev.next = curr.next. prev moves only when curr is kept, because after a deletion prev is still the predecessor of the next candidate. Return dummy.next, since the old head may have been removed (as in [7,7]).",
			},
			{
				Title:  "Delete the middle node",
				Prompt: "Delete the middle node (index n // 2, 0-based) of a list and return the head (LeetCode 2095). [1,3,4,7,1,2,6] becomes [1,3,4,1,2,6]; [1,2,3,4] becomes [1,2,4]; [1] becomes []. Combine a dummy head with fast and slow pointers so that slow stops just BEFORE the middle.",
				Code: `def deleteMiddle(head):
    dummy = ListNode(0, head)
    slow = ____
    fast = head
    while fast and fast.next:
        slow = slow.next
        fast = ____
    slow.next = ____
    return dummy.next`,
				Blanks:      []string{"dummy", "fast.next.next", "slow.next.next"},
				Explanation: "Starting slow one node behind (on the dummy) while fast starts on head makes slow stop on the predecessor of the middle instead of on the middle itself, which is what deletion needs. Then slow.next = slow.next.next unlinks the middle. For [1] the loop never runs, slow is still the dummy, and the only node is removed.",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Trace the reversal loop (save, rewire, advance) on [1,2,3]. What are prev and curr after exactly two iterations?",
				Choices: []string{
					"prev = node 2, with 2 -> 1 -> None; curr = node 3",
					"prev = node 1; curr = node 2",
					"prev = node 3, with 3 -> 2 -> 1; curr = None",
					"prev = node 2, with 2 -> 3 -> None; curr = node 3",
				},
				Answer:      0,
				Explanation: "Iteration 1 flips 1 (1 -> None) and leaves prev = 1, curr = 2. Iteration 2 flips 2 (2 -> 1) and leaves prev = 2, curr = 3. Choice B is the state after one iteration, C after three, and D forgets that 2's next was rewired.",
			},
			{
				Prompt: "Which pattern fits? \"Return the middle node of a singly linked list in one pass, using O(1) extra memory.\"",
				Choices: []string{
					"Converging two pointers from both ends, as for pair sum on a sorted array (Chapter 3)",
					"A hash map from index to node, then look up index n // 2 (Chapter 2)",
					"Fast and slow pointers moving in the same direction (Chapter 7)",
					"A sliding window over the nodes (Chapter 4)",
				},
				Answer:      2,
				Explanation: "Converging pointers need a right end that can move left, but a singly linked list can't be walked backwards. The hash map works but uses O(n) memory. There is no window condition to maintain. Fast (2 steps) and slow (1 step) pointers leave slow at the middle when fast reaches the end, in one pass with two variables.",
			},
			{
				Prompt: "Why do linked-list solutions often start with dummy = ListNode(0, head)?",
				Choices: []string{
					"It makes the list one node longer, which fixes off-by-one errors in the length",
					"It gives the first real node a predecessor, so deleting or inserting at the front needs no special case, and dummy.next is always the current head",
					"It is required by Python, because head may be None",
					"It speeds the algorithm up from O(n) to O(1)",
				},
				Answer:      1,
				Explanation: "Edits need the predecessor of the node being changed, and the head has none. The dummy supplies one, so a single code path handles every position, and dummy.next tracks the head even if the original head is removed. It doesn't change the complexity and has nothing to do with Python itself.",
			},
			{
				Prompt: "What does this function return for the list [1,2,3,4]?\n\ndef mystery(head):\n    slow = fast = head\n    while fast and fast.next:\n        slow = slow.next\n        fast = fast.next.next\n    return slow.val",
				Choices: []string{
					"2",
					"3",
					"4",
					"It raises AttributeError",
				},
				Answer:      1,
				Explanation: "slow goes 1 -> 2 -> 3 while fast goes 1 -> 3 -> None. When fast is None the loop stops and slow is on 3, the second of the two middle nodes. The fast and fast.next check prevents any AttributeError.",
			},
			{
				Prompt:      "Explain in your own words why the fast pointer must land on the slow pointer when the list has a cycle, and why it can't jump over it.",
				Explanation: "Once both pointers are inside the cycle (length L), measure how many nodes fast is behind slow along the cycle: a number between 0 and L - 1. Each step fast moves 2 and slow moves 1, so the gap shrinks by exactly 1. A whole number that decreases by 1 each step must pass through 0, and gap 0 means both are on the same node. It can't jump from 1 to -1 because it never decreases by 2. So they meet within L steps after slow enters the cycle, O(n) in total.",
			},
			{
				Prompt:      "Prerequisite check (Chapter 3): two pointers on a sorted array often start at both ends and move inwards. Why can't you do that on a singly linked list, and what do linked-list problems do instead?",
				Explanation: "Moving the right pointer inwards means stepping to the previous element, and a singly linked list has no prev pointer; reaching the last node also costs a full walk. Linked-list problems keep both pointers moving forwards: at different speeds (middle, cycle), with a fixed gap (n-th from the end), or they reverse half the list first so it can be walked forwards (Reorder List, palindrome checks).",
			},
			{
				Prompt:      "Prerequisite check (Chapter 2): an LRU cache must find any key and move it to the front in O(1). What does the hash map contribute, and why does a doubly linked list alone not give O(1) get?",
				Explanation: "The list keeps keys in recency order and can move or remove a node in O(1), but only once you hold that node: finding the node for key k means walking the list, O(n). The hash map from Chapter 2 stores key -> node, so the node is found in O(1) average time. Then unlinking (using the node's prev and next) and pushing it to the front are O(1) too.",
			},
			{
				Prompt: "You hold a reference to a node in the middle of a list and want to delete it. How long does it take in a singly linked list and in a doubly linked list?",
				Choices: []string{
					"O(1) in both, because you already hold the node",
					"O(n) in the singly linked list (you must find its predecessor from the head), O(1) in the doubly linked list (node.prev is the predecessor)",
					"O(n) in both, because the nodes after it must shift",
					"O(log n) in both",
				},
				Answer:      1,
				Explanation: "Deleting means changing the predecessor's next pointer. A singly linked node doesn't know its predecessor, so you walk from the head to find it, O(n). A doubly linked node has prev, so node.prev.next = node.next and node.next.prev = node.prev finish in O(1). Nothing ever shifts in a linked list; shifting is an array cost.",
			},
		},
	}
}
