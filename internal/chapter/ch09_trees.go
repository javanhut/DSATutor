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
			"Apply rerooting to get an answer for every node as the root in O(n): one bottom-up pass for subtree facts, one top-down pass that moves the root across each edge",
		},
		WhenToUse: []string{
			"The input is a TreeNode root, or the data is hierarchical (parent/child, folders, nested structure)",
			"The answer for a node depends on answers for its subtrees (height, size, sums, balanced, same shape): use DFS",
			"The problem says level, depth, row, layer, \"closest to the root\" or \"view from the side\": use BFS by levels",
			"The tree is a binary search tree, or you need sorted order, the kth smallest, or a range: use the BST ordering and inorder traversal",
			"You must rebuild or encode a tree: use a traversal order plus a rule that pins down the structure (null markers or a second traversal)",
			"An unrooted tree (n nodes, n - 1 undirected edges) asks for an answer \"for every node\" or \"if each node were the root\": use rerooting instead of one search per node",
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
			{
				Name:        "Rerooting: an answer for every root in two passes",
				Description: "Some problems give an unrooted tree as n and an undirected edge list and ask for a value for every node as if that node were the root (total distance to everyone, edges to flip, guesses that hold). Rerooting computes the answer for one root, then slides the root across each edge, fixing the answer in O(1) per move.",
				WhyItWorks: "Pick any root, say node 0, and compute its answer plus a few facts about each subtree (usually its size) in one " +
					"bottom-up pass. Now move the root from p to its child c. Every path that does not cross the edge p-c looks exactly the " +
					"same as before; only the relationship across that one edge changes. For a sum of distances, the size[c] nodes in c's " +
					"subtree each get one step closer and the n - size[c] other nodes each get one step farther, so answer[c] = answer[p] - " +
					"size[c] + (n - size[c]). That is O(1) work per edge instead of an O(n) recomputation, and walking the edges top-down " +
					"(parents before children) fills in all n answers in O(n) total instead of O(n^2).",
				Intuition: "Picture a group of friends spread over a street map that is a tree, choosing a meeting corner. The total walking " +
					"distance at corner p is the old answer. Moving the meeting point one street over to corner c (moving the root across " +
					"one edge) shortens the walk by one block for everyone who lives on c's side of that street (the size[c] nodes in c's " +
					"subtree) and lengthens it by one block for everyone on p's side (the other n - size[c] nodes). Nobody has to re-walk " +
					"the whole map to price the new corner: you just count heads on each side of the street.",
				CoreIdeas: []string{
					"Root it anywhere: build an undirected adjacency list, then a BFS from node 0 gives a parent array and an order where every parent comes before its children",
					"Collect subtree facts bottom-up: walk that order in reverse (children before parents, a postorder) and add each node's size or sum into its parent",
					"Get the answer for the first root from those facts, then walk the order forward and set answer[child] from answer[parent] with an O(1) formula",
					"Find the formula by asking what changes across the single edge parent-child when the root moves over it; everything else stays put",
					"Two linear passes over n nodes: O(n) time and O(n) space, with no recursion at all",
				},
				CommonMistakes: []string{
					"Running a fresh BFS or DFS from every node, which is O(n^2) and times out for n = 30,000: compute one root's answer, then derive the others with the O(1) move",
					"Writing the passes as recursion, which hits Python's recursion limit of about 1,000 on a path-shaped tree: use the BFS order and loop over it backward, then forward",
					"Adding each edge in one direction only (graph[a].append(b) without graph[b].append(a)), so a BFS from node 0 cannot leave a node through an edge that was listed the other way round and misses part of the tree: an undirected edge goes into both lists",
					"Running the forward pass in reversed order, so answer[parent] is still 0 when a child reads it: walk the BFS order forward so every parent is final before its children",
				},
				Examples: []CodeExample{
					{
						ID:       "tree_rerooting_reversals",
						Title:    "Rerooting: one-way roads to reverse for every capital",
						Language: "python",
						Snippet: `def minReversals(n, edges):
    # edges[i] = [u, v] is a one-way road u -> v; the roads form a tree.
    # answer[r] = roads to reverse so every city can be reached from r.
    graph = [[] for _ in range(n)]
    for u, v in edges:
        graph[u].append((v, 0))   # walking u -> v follows the arrow
        graph[v].append((u, 1))   # walking v -> u needs one reversal
    parent = [-1] * n
    down_cost = [0] * n           # cost of walking parent -> node
    seen = [False] * n
    seen[0] = True
    order = [0]                   # Root it anywhere: BFS order from 0
    for node in order:
        for nxt, cost in graph[node]:
            if not seen[nxt]:
                seen[nxt] = True
                parent[nxt] = node
                down_cost[nxt] = cost
                order.append(nxt)
    answer = [0] * n
    answer[0] = sum(down_cost)    # every road must point away from 0
    for node in order[1:]:        # Move the root across an edge
        # only the road parent-node is now walked the other way
        answer[node] = answer[parent[node]] + (1 if down_cost[node] == 0 else -1)
    return answer

# minReversals(4, [[2, 0], [2, 1], [1, 3]]) -> [1, 1, 0, 2]`,
						Notes: "Here the root's answer needs no subtree facts, so the bottom-up pass shrinks to a sum. When it does need them (sizes, weights), add a reversed(order) pass, as in the worked example.",
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

			{
				Title: "Total walking distance to every meeting point (rerooting)",
				Problem: "A town has n houses numbered 0 to n - 1 joined by n - 1 two-way streets that form a tree: edges[i] = [a, b] is " +
					"an undirected street of length 1, and every house can reach every other. people[i] people live in house i. For " +
					"every house h, if everyone walks to h along the tree, the total walking distance is the sum over all houses v of " +
					"people[v] * dist(h, v). Return a list answer where answer[h] is that total. Example: n = 6, edges = " +
					"[[0,1],[0,2],[2,3],[2,4],[4,5]], people = [1,2,1,3,1,2] gives [17, 23, 13, 17, 17, 23], so house 2 is the cheapest meeting point.",
				Steps: []WorkedStep{
					{
						Subgoal: "Build the undirected adjacency",
						Explanation: "A street can be walked both ways, so each edge goes into both endpoints' lists. Any node can then become " +
							"the root later, because the BFS can leave every node in every direction.",
						Code: `graph = [[] for _ in range(n)]
for a, b in edges:
    graph[a].append(b)
    graph[b].append(a)`,
						State: "graph = [[1, 2], [0], [0, 3, 4], [2], [2, 5], [4]]",
					},
					{
						Subgoal: "Root it anywhere",
						Explanation: "Choose node 0 as the first root and run a BFS that records each node's parent. Iterating over order while " +
							"appending to it is a queue without popping, and it gives an order where every parent comes before its children. " +
							"No recursion, so a 100,000-node path is fine.",
						Code: `parent = [-1] * n
seen = [False] * n
seen[0] = True
order = [0]
for node in order:
    for nxt in graph[node]:
        if not seen[nxt]:
            seen[nxt] = True
            parent[nxt] = node
            order.append(nxt)`,
						State: "order = [0, 1, 2, 3, 4, 5], parent = [-1, 0, 0, 2, 2, 4]",
					},
					{
						Subgoal: "Collect subtree facts bottom-up",
						Explanation: "Walk the order backward, so every child is finished before its parent (a postorder, as in the DFS template). " +
							"weight[v] becomes the number of people in v's subtree. cost[v] is the walking total for v's subtree to reach v: " +
							"the child's own total plus one extra step for each of the weight[child] people crossing the edge up to v.",
						Code: `weight = people[:]
cost = [0] * n
for node in reversed(order):
    p = parent[node]
    if p != -1:
        weight[p] += weight[node]
        cost[p] += cost[node] + weight[node]`,
						State: "node 5: weight[4] = 3, cost[4] = 2; node 4: weight[2] = 4, cost[2] = 5; node 3: weight[2] = 7, cost[2] = 8; " +
							"node 2: weight[0] = 8, cost[0] = 15; node 1: weight[0] = 10, cost[0] = 17. Final weight = [10, 2, 7, 3, 3, 2]",
					},
					{
						Subgoal: "Answer for the first root",
						Explanation: "The root's subtree is the whole tree, so cost[0] is already the full answer for house 0, and weight[0] is " +
							"the total number of people. These two numbers seed the top-down pass.",
						Code: `total = weight[0]
answer = [0] * n
answer[0] = cost[0]`,
						State: "total = 10, answer[0] = 17 (2*1 + 1*1 + 3*2 + 1*2 + 2*3)",
					},
					{
						Subgoal: "Move the root across an edge",
						Explanation: "Walk the order forward so answer[p] is final before its child reads it. Moving from p to child c, the " +
							"weight[c] people on c's side each walk one step less and the total - weight[c] people on p's side each walk one " +
							"step more. Nothing else changes, so the update is O(1).",
						Code: `for node in order[1:]:
    p = parent[node]
    answer[node] = answer[p] - weight[node] + (total - weight[node])`,
						State: "node 1: 17 - 2 + 8 = 23; node 2: 17 - 7 + 3 = 13; node 3: 13 - 3 + 7 = 17; node 4: 13 - 3 + 7 = 17; " +
							"node 5: 17 - 2 + 8 = 23. answer = [17, 23, 13, 17, 17, 23]",
					},
				},
				Code: `def meetingCosts(n, edges, people):
    graph = [[] for _ in range(n)]
    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)
    parent = [-1] * n
    seen = [False] * n
    seen[0] = True
    order = [0]
    for node in order:
        for nxt in graph[node]:
            if not seen[nxt]:
                seen[nxt] = True
                parent[nxt] = node
                order.append(nxt)
    weight = people[:]
    cost = [0] * n
    for node in reversed(order):
        p = parent[node]
        if p != -1:
            weight[p] += weight[node]
            cost[p] += cost[node] + weight[node]
    total = weight[0]
    answer = [0] * n
    answer[0] = cost[0]
    for node in order[1:]:
        p = parent[node]
        answer[node] = answer[p] - weight[node] + (total - weight[node])
    return answer`,
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
				Explanation: "Blank by blank. The empty tree has height 0, and that base case feeds every leaf's children. The imbalance test " +
					"is abs(left_h - right_h) > 1, straight from the definition; it sits beside the two == -1 checks so a failure anywhere " +
					"below also bubbles up. The combine line returns 1 + max(left_h, right_h): a node is one level above its taller child. " +
					"Folding the \"unbalanced\" signal into the return value keeps it one O(n) pass instead of recomputing heights at every " +
					"node (O(n^2)). For the left chain 1 -> 2 -> 3: height(3) = 1 + max(0, 0) = 1, height(2) = 1 + max(1, 0) = 2, and at " +
					"the root abs(2 - 0) > 1 sends -1 up: unbalanced.",
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
				Explanation: "Everything smaller than node.val lives in the left subtree and everything larger in the right. So when val < node.val, the target can only be on the left: node = node.left; otherwise val is larger (the loop condition already rules out equal) and node = node.right. Each comparison discards a whole subtree, and the loop walks one root-to-leaf path: O(h) time, O(1) space. Searching 5 under root 4 with right child 7 and 7's left child 5: 5 > 4 goes right, 5 < 7 goes left, found.",
			},
			{
				Title: "Roots that make enough guesses true (rerooting)",
				Prompt: "A tree has n = len(edges) + 1 nodes and undirected edges. Each guess [u, v] (no duplicates) claims \"u is the parent " +
					"of v\". For a chosen root, a guess is true when u really is v's parent. Return how many roots make at least k guesses " +
					"true. The tree is rooted at 0 with a BFS order and parent array. Fill in the count for root 0 and the two O(1) " +
					"corrections when the root moves from p to its child node.",
				Code: `def rootCount(edges, guesses, k):
    n = len(edges) + 1
    graph = [[] for _ in range(n)]
    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)
    guessed = set(map(tuple, guesses))
    parent = [-1] * n
    seen = [False] * n
    seen[0] = True
    order = [0]
    for node in order:
        for nxt in graph[node]:
            if not seen[nxt]:
                seen[nxt] = True
                parent[nxt] = node
                order.append(nxt)
    correct = [0] * n
    correct[0] = sum(____ for v in order[1:])
    for node in order[1:]:
        p = parent[node]
        lose = ____
        gain = ____
        correct[node] = correct[p] - lose + gain
    return sum(c >= k for c in correct)`,
				Blanks: []string{"(parent[v], v) in guessed", "(p, node) in guessed", "(node, p) in guessed"},
				Explanation: "Blank by blank. With root 0, every non-root v has parent[v], so the true guesses are exactly the pairs " +
					"(parent[v], v) found in guessed; True counts as 1 in sum. Moving the root from p to its child node flips only the " +
					"edge p-node: p stops being node's parent, so a guess (p, node) that was true becomes false (lose), and node becomes " +
					"p's parent, so a guess (node, p) becomes true (gain). Every other parent link is unchanged, so each move is O(1) and " +
					"the whole thing is O(n + g). For edges [[0,1],[1,2],[1,3],[3,4]] and guesses [[1,3],[0,1],[1,0],[4,3]]: correct[0] = 2; " +
					"node 1 loses (0,1) and gains (1,0): 2; node 2: 2; node 3 loses (1,3): 1; node 4 gains (4,3): 2. With k = 2 the answer is 4.",
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
			{
				Prompt: "A tree is rooted at node 0 and answer[p] is the sum of distances from p to every node. You move the root from p to its child c. Which nodes become exactly one step closer to the root?",
				Choices: []string{
					"Only c itself, because it is the only node whose distance drops to 0",
					"The size[c] nodes in c's subtree, c included",
					"The n - size[c] nodes outside c's subtree, because they are now below the new root",
					"Every node except p, because the root moved downward",
				},
				Answer: 1,
				Explanation: "Every path from c's subtree to p went through c, so starting at c instead saves exactly one step for all " +
					"size[c] of them (c drops from 1 to 0). Every node outside that subtree must now cross the edge p-c to reach c, so those " +
					"n - size[c] nodes are one step farther. That is why answer[c] = answer[p] - size[c] + (n - size[c]). Choice 3 has the " +
					"two groups reversed, and choice 1 forgets that c's descendants also gain.",
			},
			{
				Prompt: "Trace it: four houses on a line, edges [[0,1],[1,2],[2,3]], with people = [3,1,1,1]. Rooted at 0, compute weight[] " +
					"(people in each subtree) and answer[0], then use the rerooting update to get answer[1], answer[2] and answer[3].",
				Explanation: "Bottom-up from 0: weight = [6, 3, 2, 1] (total 6). answer[0] = 1*1 + 1*2 + 1*3 = 6. Moving the root forward with " +
					"answer[c] = answer[p] - weight[c] + (total - weight[c]): answer[1] = 6 - 3 + 3 = 6, answer[2] = 6 - 2 + 4 = 8, " +
					"answer[3] = 8 - 1 + 5 = 12. Check one directly: at house 2 the walks are 3*2 + 1*1 + 0 + 1*1 = 8. So answer = [6, 6, 8, 12], " +
					"and houses 0 and 1 tie as the cheapest meeting point because the 3 people at house 0 pull it left.",
			},
		},
	}
}
