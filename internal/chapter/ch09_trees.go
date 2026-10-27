package chapter

import "time"

// chapter09 teaches binary trees and binary search trees: the DFS template
// ("what do I need from my children, what do I return to my parent"), BFS by
// levels, and the BST ordering. Storyboards (IDs registered in app.js):
// binary-tree.
func chapter09() Chapter {
	return Chapter{
		Number:        9,
		Slug:          "trees",
		Title:         "Binary Trees and BSTs",
		Category:      "trees",
		Prerequisites: []int{7, 8},
		Summary: "A binary tree is a linked structure where each node has up to two children, so almost every tree problem is solved by " +
			"one recursive question: what do I need from my children, and what do I return to my parent? This chapter teaches that " +
			"DFS template, breadth-first search one level at a time, and the binary search tree ordering that lets you discard half " +
			"the tree at every step.",
		Motivation: "Suppose a phone book app holds 1,000,000 names and must both look names up and insert new ones all day. " +
			"A sorted array finds a name in about 20 steps with binary search (Chapter 6), but inserting \"Aaron\" shifts all " +
			"1,000,000 entries one slot to the right. A linked list (Chapter 7) inserts by changing two pointers, but finding the " +
			"spot means walking up to 1,000,000 nodes. A balanced binary search tree gets both: every comparison sends you left or " +
			"right and throws away half of what is left, so a lookup or an insert touches about log2(1,000,000) = 20 nodes. Trees " +
			"also model data that is hierarchical by nature (folders, HTML documents, org charts, expression parse trees), and the " +
			"same few templates solve nearly all of the interview problems about them.",
		Objectives: []string{
			"Trace preorder, inorder and postorder DFS and level-order BFS on a small tree by hand",
			"Apply the DFS template: state what the function returns, handle the empty tree, ask the children, combine and return to the parent",
			"Explain why snapshotting len(queue) separates BFS into levels",
			"Explain why a valid BST needs bounds passed down from all ancestors, not just a parent-child comparison",
			"Implement tree solutions in O(n) time and O(h) extra space, and say when h is log n and when it is n",
		},
		WhenToUse: []string{
			"The input is a TreeNode root, or the data is hierarchical (parent/child, folders, nested structure)",
			"The answer for a node depends on answers for its subtrees (height, size, sums, balanced, same shape): use DFS",
			"The problem says level, depth, row, layer, \"closest to the root\" or \"view from the side\": use BFS by levels",
			"The tree is a binary search tree, or you need sorted order, the kth smallest, or a range: use the BST ordering and inorder traversal",
			"You must rebuild or encode a tree: use a traversal order plus a rule that pins down the structure (null markers or a second traversal)",
		},
		Concepts: []Concept{
			{
				Name:        "Binary trees and the DFS template",
				Description: "A binary tree is either empty (None) or a node with a value, a left subtree and a right subtree. Depth-first search solves a tree problem by solving it for both subtrees and combining the answers.",
				WhyItWorks: "The definition of a tree is recursive: a tree is empty, or a root plus two smaller trees. So if you know the answer for " +
					"both subtrees, you only need a constant amount of work to get the answer for the whole tree, and the recursion " +
					"stops because every call is on a strictly smaller tree until it reaches None. This is exactly Chapter 8's recipe: " +
					"state what the function returns, handle the base case (the empty tree), trust the recursive call on each child, " +
					"and combine the results. Each node is visited once, so the time is O(n); the call stack holds one frame per level " +
					"of the current path, so the extra space is O(h), the height, which is about log2 n for a balanced tree and n for a chain.",
				Intuition: "Think of a company org chart where the CEO wants the headcount. She does not count everyone herself: she asks " +
					"each of her direct reports \"how many people are in your part of the company?\", adds 1 for herself, and reports " +
					"upward. Each manager is a node, the question is the recursive call, the reply is the return value, and an empty " +
					"team (None) answers 0.",
				CoreIdeas: []string{
					"Ask two questions for every tree problem: what do I need from my children, and what do I return to my parent?",
					"The empty tree is the base case: return the neutral answer (0 for heights and sums, True for \"all nodes satisfy\", None for a missing node)",
					"Preorder (node, left, right) does work on the way down, postorder (left, right, node) works on the way up after the children answer, inorder (left, node, right) visits a BST in sorted order",
					"When the answer you report upward differs from the answer you want overall (diameter, max path sum), keep a separate global best and update it at every node",
					"Time O(n) because each node is visited once; space O(h) for the call stack",
				},
				CommonMistakes: []string{
					"Checking node.left.val before checking that node.left exists: handle None first, at the top of the function, and let the recursive call deal with missing children",
					"Returning the global answer upward instead of what the parent needs (for diameter the parent needs a height, not a diameter): write down what the function returns before coding",
					"Forgetting the +1 for the current node when combining child heights, which makes every height come out 0",
					"Assuming the tree is balanced when estimating space: a chain of 10,000 nodes makes 10,000 nested calls, so mention O(h) and that h can be n",
				},
				Examples: []CodeExample{
					{
						ID:       "tree_dfs_template",
						Title:    "The DFS template: count the nodes",
						Language: "python",
						Snippet: `def count_nodes(node):
    # State what the function returns: the number of nodes in this subtree
    if node is None:                       # Handle the empty tree
        return 0
    left_count = count_nodes(node.left)    # Ask the children
    right_count = count_nodes(node.right)
    return 1 + left_count + right_count    # Combine and return to the parent`,
						Notes: "Every DFS in this chapter has these same three parts; only the base value and the combine line change.",
					},
					{
						ID:       "tree_traversals",
						Title:    "Preorder, inorder and postorder",
						Language: "python",
						Snippet: `def preorder(node, out):
    if node:
        out.append(node.val)   # node first
        preorder(node.left, out)
        preorder(node.right, out)

def inorder(node, out):
    if node:
        inorder(node.left, out)
        out.append(node.val)   # node in the middle
        inorder(node.right, out)

def postorder(node, out):
    if node:
        postorder(node.left, out)
        postorder(node.right, out)
        out.append(node.val)   # node last

# Tree [1,2,3,4,5,null,6]:
# preorder  1 2 4 5 3 6
# inorder   4 2 5 1 3 6
# postorder 4 5 2 6 3 1`,
						Notes: "The three orders differ only in where the append sits relative to the two recursive calls.",
					},
				},
			},
			{
				Name:        "Breadth-first search by levels",
				Description: "BFS visits the tree one level at a time, from the root downward and left to right within a level, using a FIFO queue.",
				WhyItWorks: "A queue hands nodes out in the order they went in. The root goes in first; when you take a node out you put its " +
					"children in at the back, so every node of depth d is enqueued before any node of depth d + 1. That gives the " +
					"invariant: at the start of each pass of the outer loop, the queue holds exactly the nodes of one level, in left-to-right " +
					"order. Snapshotting level_size = len(queue) before the inner loop lets you take exactly that level out, even though " +
					"you are adding the next level to the back at the same time. Every node is enqueued and dequeued once, so BFS is O(n) " +
					"time and O(w) space, where w is the widest level (up to about n/2 for a full tree).",
				Intuition: "Picture a family photo taken generation by generation. The photographer calls everyone in the current generation " +
					"(the queue), and as each person steps up they bring their children to the back of the line. Counting the line before " +
					"the shoot starts (level_size) tells the photographer when one generation's photo is finished and the next begins.",
				CoreIdeas: []string{
					"Use collections.deque: popleft() is O(1), while list.pop(0) shifts every element and is O(n)",
					"Snapshot level_size = len(queue) before processing a level, then pop exactly that many nodes",
					"Only enqueue children that exist, or check for None when you pop",
					"Use BFS when the problem talks about levels, depth, rows or the first node found closest to the root; use DFS when a node's answer is built from its subtrees",
				},
				CommonMistakes: []string{
					"Looping while queue: for each node popped, without the level_size snapshot, which mixes nodes from two levels into one list: snapshot the length first",
					"Using a stack (list.pop()) instead of a queue, which turns the traversal into DFS and breaks the level order: use deque.popleft()",
					"Forgetting the empty-tree check, so deque([None]) pops None and crashes on None.left: return [] when root is None",
				},
				Examples: []CodeExample{
					{
						ID:       "tree_bfs_levels",
						Title:    "Level-by-level BFS",
						Language: "python",
						Snippet: `from collections import deque

def levels(root):
    if root is None:
        return []
    out = []
    queue = deque([root])                  # Seed the queue
    while queue:
        level = []
        for _ in range(len(queue)):        # Process one level
            node = queue.popleft()
            level.append(node.val)
            if node.left:                  # Enqueue the children
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        out.append(level)                  # Record the level
    return out`,
						Notes: "range(len(queue)) is evaluated once, before the loop starts, so it is the snapshot of the current level.",
					},
				},
			},
			{
				Name:        "Binary search trees",
				Description: "A binary search tree (BST) is a binary tree where, for every node, all values in its left subtree are smaller and all values in its right subtree are larger.",
				WhyItWorks: "The ordering rule holds for every node and for its whole subtrees, not just its children. So comparing a target " +
					"with a node tells you which subtree can contain it, and the other subtree can be discarded entirely: this is binary " +
					"search (Chapter 6) on a linked structure, costing O(h) comparisons. The same rule means an inorder traversal (left, " +
					"node, right) lists the values in increasing order: everything in the left subtree is smaller and comes first, then " +
					"the node, then everything larger. To validate a BST you must check each node against bounds inherited from all of its " +
					"ancestors, because a node deep in the right subtree of 5 must be greater than 5 even if it is smaller than its own parent.",
				Intuition: "A BST is a game of \"higher or lower\" frozen into a structure. At each node the tree answers \"lower\" (go left) or " +
					"\"higher\" (go right), so you never look at the half that cannot hold your number. The bounds are the range of numbers " +
					"still possible after all the answers you have heard so far, not just the last one.",
				CoreIdeas: []string{
					"Use the ordering to pick a side: if the target is smaller than node.val go left, if larger go right",
					"Inorder traversal of a BST gives the values in sorted order, so the kth smallest is the kth node visited",
					"Validity needs bounds: pass (low, high) down, tightening high when you go left and low when you go right",
					"Search, insert and delete cost O(h): about log2 n when balanced, but n when the tree degenerates into a chain (for example, inserting 1, 2, 3, ... in order)",
				},
				CommonMistakes: []string{
					"Validating with only node.left.val < node.val < node.right.val: [5,4,6,null,null,3,7] passes that check but 3 sits in the right subtree of 5; pass bounds down instead",
					"Allowing duplicates with <= when the problem says strictly less and greater: read the definition and use strict comparisons",
					"Using integer sentinels like -2**31 for the initial bounds, which wrongly rejects a node whose value is exactly -2**31: start with float('-inf') and float('inf') or None",
				},
				Examples: []CodeExample{
					{
						ID:       "bst_search_insert",
						Title:    "Search and insert in a BST",
						Language: "python",
						Snippet: `def search(node, target):
    while node and node.val != target:
        # Use the ordering to pick a side
        node = node.left if target < node.val else node.right
    return node

def insert(node, val):
    if node is None:
        return TreeNode(val)
    if val < node.val:
        node.left = insert(node.left, val)
    else:
        node.right = insert(node.right, val)
    return node`,
						Notes: "Both walk one root-to-leaf path, so they cost O(h).",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title: "Diameter of a binary tree (DFS template with a global answer)",
				Problem: "Given the root of a binary tree, return the length of its diameter: the number of edges on the longest path " +
					"between any two nodes. The path does not have to pass through the root. Example: [1,2,3,4,5] has diameter 3 " +
					"(the path 4 -> 2 -> 1 -> 3).",
				Steps: []WorkedStep{
					{
						Subgoal: "State what the function returns",
						Explanation: "The longest path through a node goes down into its left subtree as far as possible and down into its " +
							"right subtree as far as possible. So what a node needs from each child is that child's height (nodes on its " +
							"longest downward path). The helper returns a height, not a diameter, because a height is what a parent can " +
							"extend; a diameter that bends at some lower node cannot be extended upward.",
						Code:  "def height(node):  # nodes on the longest downward path from node",
						State: "height(None) = 0, height(leaf) = 1",
					},
					{
						Subgoal:     "Handle the empty tree",
						Explanation: "An empty subtree has height 0. This is the base case that stops the recursion at every missing child, so no code ever touches None.left.",
						Code: `if node is None:
    return 0`,
						State: "height(4.left) = 0, height(4.right) = 0",
					},
					{
						Subgoal:     "Ask the children",
						Explanation: "Trust the recursive call (Chapter 8): assume height(node.left) and height(node.right) are correct for the smaller trees, and do not trace into them.",
						Code: `left_h = height(node.left)
right_h = height(node.right)`,
						State: "at node 2: left_h = 1 (node 4), right_h = 1 (node 5)",
					},
					{
						Subgoal: "Update the global answer",
						Explanation: "The longest path that bends at this node has left_h + right_h edges (left_h edges down the left side, " +
							"right_h down the right). Record it in best. This happens at every node, so the path that bends at any node is " +
							"considered.",
						Code:  "best = max(best, left_h + right_h)",
						State: "node 4: best 0; node 5: best 0; node 2: best = 1 + 1 = 2; node 3: best 2; node 1: left_h 2, right_h 1, best = 3",
					},
					{
						Subgoal:     "Combine and return to the parent",
						Explanation: "The parent can extend only one branch, so return the longer side plus 1 for this node.",
						Code:        "return 1 + max(left_h, right_h)",
						State:       "node 2 returns 2, node 3 returns 1, node 1 returns 3; answer best = 3",
					},
				},
				Code: `def diameterOfBinaryTree(root):
    best = 0

    def height(node):
        nonlocal best
        if node is None:
            return 0
        left_h = height(node.left)
        right_h = height(node.right)
        best = max(best, left_h + right_h)
        return 1 + max(left_h, right_h)

    height(root)
    return best`,
			},
			{
				Title: "Right side view (BFS by levels)",
				Problem: "Given the root of a binary tree, imagine standing on its right side and return the values you can see from top " +
					"to bottom: the last node of each level. Example: [1,2,3,null,5,null,4] gives [1,3,4].",
				Steps: []WorkedStep{
					{
						Subgoal:     "Seed the queue",
						Explanation: "Return [] for an empty tree, otherwise put the root in a deque. The queue now holds exactly level 0.",
						Code: `if root is None:
    return []
queue = deque([root])`,
						State: "queue = [1], view = []",
					},
					{
						Subgoal:     "Process one level",
						Explanation: "Snapshot level_size = len(queue) and pop exactly that many nodes. Because of the snapshot, children added during this pass wait for the next pass.",
						Code: `level_size = len(queue)
for i in range(level_size):
    node = queue.popleft()`,
						State: "pass 1: level_size = 1 (node 1); pass 2: level_size = 2 (nodes 2, 3); pass 3: level_size = 2 (nodes 5, 4)",
					},
					{
						Subgoal:     "Record the level",
						Explanation: "The node seen from the right is the last one popped in this level, at index level_size - 1.",
						Code: `if i == level_size - 1:
    view.append(node.val)`,
						State: "pass 1 records 1, pass 2 records 3, pass 3 records 4",
					},
					{
						Subgoal:     "Enqueue the children",
						Explanation: "Append the left child before the right child so each level stays in left-to-right order; skip missing children.",
						Code: `if node.left:
    queue.append(node.left)
if node.right:
    queue.append(node.right)`,
						State: "after pass 1 queue = [2, 3]; after pass 2 queue = [5, 4]; after pass 3 queue = [], view = [1, 3, 4]",
					},
				},
				Code: `from collections import deque

def rightSideView(root):
    if root is None:
        return []
    view = []
    queue = deque([root])
    while queue:
        level_size = len(queue)
        for i in range(level_size):
            node = queue.popleft()
            if i == level_size - 1:
                view.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
    return view`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title: "Balanced binary tree",
				Prompt: "A tree is height-balanced when, at every node, the heights of the left and right subtrees differ by at most 1. " +
					"The helper returns the height of a subtree, or -1 if that subtree is already unbalanced. Fill in the base case, the " +
					"imbalance test and the combine line.",
				Code: `def isBalanced(root):
    def height(node):
        if node is None:
            return ____
        left_h = height(node.left)
        right_h = height(node.right)
        if left_h == -1 or right_h == -1 or ____:
            return -1
        return ____

    return height(root) != -1`,
				Blanks: []string{"0", "abs(left_h - right_h) > 1", "1 + max(left_h, right_h)"},
				Explanation: "Handle the empty tree: height 0. Ask the children, then combine: if either child reported -1, or the heights " +
					"differ by more than 1, pass -1 up; otherwise return to the parent the height 1 + max(left_h, right_h). Folding the " +
					"\"unbalanced\" signal into the return value keeps it one O(n) pass instead of recomputing heights at every node (O(n^2)).",
			},
			{
				Title:  "Search in a binary search tree",
				Prompt: "Return the node whose value equals val, or None if there is none. Use the ordering to pick a side.",
				Code: `def searchBST(root, val):
    node = root
    while node and node.val != val:
        if val < node.val:
            node = ____
        else:
            node = ____
    return node`,
				Blanks:      []string{"node.left", "node.right"},
				Explanation: "Everything smaller than node.val lives in the left subtree and everything larger in the right, so one comparison discards a whole subtree. The loop walks one root-to-leaf path: O(h) time and O(1) space.",
			},
		},
		Animations: []Storyboard{
			{
				ID:    "binary-tree",
				Title: "Inorder Traversal of a Binary Tree",
				Goal:  "Watch the recursion visit the tree [1,2,3,4,5,6,7] left subtree first, then the node, then the right subtree, producing 4 2 5 1 6 3 7.",
				MemoryTips: []string{
					"Handle the empty tree first: if node is None, return",
					"Ask the children: the left call finishes the whole left subtree before the node is visited",
					"Where the visit line sits picks the order: before both calls = preorder, between = inorder, after = postorder",
					"On a BST, inorder gives the values in sorted order",
					"Each node is visited once (O(n) time); the stack holds one frame per level of the current path (O(h) space)",
				},
				Steps: []StoryboardStep{
					{Cue: "left", Narration: "Start at the root 1, but do not visit it yet: first recurse into the left subtree, and keep going left (1, then 2, then 4) until the next left child is None. Each pending node waits on the call stack.", VisualHint: "Traverse left subtree", Duration: Duration(2 * time.Second), CodeRef: "inorder: inorder(node.left)"},
					{Cue: "visit", Narration: "The left call returned, so the whole left subtree is done. Now visit the current node: 4 is recorded first, then the stack unwinds to 2 and records it.", VisualHint: "Process node", Duration: Duration(2 * time.Second), CodeRef: "inorder: print(node.val)"},
					{Cue: "right", Narration: "After visiting a node, recurse into its right subtree with the same rule: from 2 go to 5, record it, and return. Then 1 is visited and the same pattern runs on the right half: 6, 3, 7.", VisualHint: "Traverse right subtree", Duration: Duration(2 * time.Second), CodeRef: "inorder: inorder(node.right)"},
					{Cue: "complete", Narration: "Every call has returned and the stack is empty. The output is 4 2 5 1 6 3 7: left subtree, root, right subtree at every level. On a binary search tree this order is sorted.", VisualHint: "Show result order", Duration: Duration(2 * time.Second), CodeRef: "inorder: return result"},
				},
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Which pattern fits best? \"Given the root of a binary tree, return the average of the node values on each level, top to bottom.\"",
				Choices: []string{
					"Sliding window (Chapter 4) over the preorder list of values",
					"BFS by levels: snapshot len(queue), sum that many nodes, divide by the count",
					"Binary search (Chapter 6) on the inorder list of values",
					"Two pointers (Chapter 3) from both ends of the level-order array",
				},
				Answer: 1,
				Explanation: "The word \"level\" is the cue for BFS by levels. A traversal list throws away depth information, so " +
					"windows, binary search and two pointers over it cannot tell where one level ends; the len(queue) snapshot does exactly that.",
			},
			{
				Prompt: "(Review of Chapter 8) A recursive DFS runs on a tree of 10,000 nodes where every node has only a left child (a chain). How many frames are on the call stack at the deepest point?",
				Choices: []string{
					"About 14, because log2(10,000) is about 13.3",
					"About 10,000, one per node on the current root-to-leaf path",
					"2, because each call makes at most two recursive calls",
					"20,000, because every node has two child calls",
				},
				Answer: 1,
				Explanation: "The stack holds one frame per call that has started but not returned, which is the current root-to-node path. " +
					"In a chain that path is all n nodes, so space is O(h) = O(n) here; log n only holds for balanced trees. Python's default " +
					"recursion limit is about 1,000, so very deep trees need an iterative version with an explicit stack.",
			},
			{
				Prompt: "A validator checks only node.left.val < node.val < node.right.val at every node. Which tree does it wrongly accept?",
				Choices: []string{
					"[2,1,3]",
					"[5,1,4,null,null,3,6]",
					"[5,4,6,null,null,3,7]",
					"[1]",
				},
				Answer: 2,
				Explanation: "In [5,4,6,null,null,3,7], node 6 has children 3 and 7, and 3 < 6 < 7 passes the local check. But 3 is in the " +
					"right subtree of 5, so it must be greater than 5. Only bounds passed down from every ancestor (low = 5 when you go right " +
					"from 5) catch it. [5,1,4,...] fails the local check already (4 < 5 on the right), and the other two are valid.",
			},
			{
				Prompt: "Trace it: for the tree [1,2,3,4,5,null,6], write the preorder, inorder and postorder traversals.",
				Explanation: "The tree is 1 with children 2 and 3; 2 has children 4 and 5; 3 has only a right child 6. Preorder (node, left, " +
					"right): 1 2 4 5 3 6. Inorder (left, node, right): 4 2 5 1 3 6. Postorder (left, right, node): 4 5 2 6 3 1. The root is " +
					"first in preorder and last in postorder, which is why preorder plus inorder can rebuild a tree.",
			},
			{
				Prompt: "Using the DFS template, explain what maxDepth(node) needs from its children and what it returns to its parent, including the base case.",
				Explanation: "It returns the number of nodes on the longest downward path from node. Base case: an empty tree has depth 0. " +
					"It needs the depths of the left and right subtrees (trust the recursive calls), and returns 1 + max(left, right) " +
					"to its parent: 1 for the node itself plus the deeper side.",
			},
			{
				Prompt: "In level-order BFS, what goes wrong if you drop the snapshot level_size = len(queue) and simply pop nodes one at a time while the queue is non-empty?",
				Choices: []string{
					"Nothing: a queue keeps levels apart automatically",
					"You cannot tell where one level ends, because children of the current level are appended to the same queue while you pop",
					"The traversal becomes depth-first",
					"It runs in O(n^2) time",
				},
				Answer: 1,
				Explanation: "The queue does keep nodes in level order, but it holds the tail of the current level and the start of the next " +
					"one at the same time. Only the count taken before the level starts tells you how many pops belong to it. The order is " +
					"still breadth-first and the time is still O(n).",
			},
		},
	}
}
