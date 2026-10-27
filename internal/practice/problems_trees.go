package practice

// TreesProblems contains all trees category problems, ordered as a ladder
// (see docs/pedagogy.md). Chapter 9 teaches the DFS template, BFS by levels
// and the BST ordering that these problems practise.
var TreesProblems = []*Problem{
	{
		ID:              "maximum-depth-of-binary-tree",
		Number:          104,
		Order:           1,
		Title:           "Maximum Depth of Binary Tree",
		Slug:            "maximum-depth-of-binary-tree",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "BFS", "Binary Tree", "Recursion"},
		RelatedChapters: []int{9, 8},
		Description: `You are given the root of a binary tree. Return its **maximum depth**: the number of nodes on the longest path that starts at the root and goes down to a leaf (a node with no children).

An empty tree (root is None) has depth 0, and a tree with only a root has depth 1.`,
		Constraints: []string{
			"The tree has between 0 and 10^4 nodes",
			"-100 <= Node.val <= 100",
		},
		Examples: []Example{
			{Input: "root = [3,9,20,null,null,15,7]", Output: "3", Explanation: "The longest root-to-leaf paths are 3 -> 20 -> 15 and 3 -> 20 -> 7, each with 3 nodes. The path 3 -> 9 has only 2."},
			{Input: "root = [1,null,2]", Output: "2", Explanation: "The root has only a right child, so the path 1 -> 2 has 2 nodes."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 9, 20, nil, nil, 15, 7}}, Expected: 3, Note: "example 1: deeper on the right"},
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2}}, Expected: 2, Note: "example 2: only a right child"},
			{Input: map[string]interface{}{"root": []interface{}{}}, Expected: 0, Hidden: true, Note: "empty tree has depth 0"},
			{Input: map[string]interface{}{"root": []interface{}{0}}, Expected: 1, Hidden: true, Note: "single node counts as depth 1 (forgetting the +1 gives 0)"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, nil, 3, nil, 4}}, Expected: 4, Hidden: true, Note: "left-leaning chain of 4 nodes"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, nil, nil, 4, nil, 5}}, Expected: 4, Hidden: true, Note: "deepest leaf is under the right child's left side (checking only one side fails)"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef maxDepth(root):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If a friend told you the depth of the left subtree and the depth of the right subtree, how would you get the depth of the whole tree?"},
			{Level: 2, Type: "insight", Content: "Depth is defined recursively: an empty tree has depth 0, and any other tree has depth 1 + max(depth of left, depth of right). This is the DFS template from Chapter 9: what do I need from my children (their depths), what do I return to my parent (my own depth).", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Handle the empty tree: if root is None, the depth is 0.\n2. Ask the children: compute the depth of the left subtree and of the right subtree with recursive calls, trusting them to be correct.\n3. Combine and return to the parent: return 1 (this node) plus the larger of the two depths."},
			{Level: 4, Type: "skeleton", Content: "def maxDepth(root):\n    if root is None:\n        return ____\n    left_depth = ____\n    right_depth = maxDepth(root.right)\n    return ____"},
		},
		Solution: Solution{
			Code: `def maxDepth(root):
    if root is None:
        return 0
    left_depth = maxDepth(root.left)
    right_depth = maxDepth(root.right)
    return 1 + max(left_depth, right_depth)`,
			Explanation: "This is the first rung of the trees ladder and the plainest use of the DFS template from Chapter 9, which itself is Chapter 8's recursion recipe applied to a tree: the depth of a tree is 1 for the root plus the depth of its deeper subtree, and an empty tree has depth 0. Each node asks its two children for their depths and reports one number to its parent.",
			BruteForce:  "You could list every root-to-leaf path explicitly (for example, by copying the path so far at each step) and take the longest one. That copies up to h values for each of up to n/2 leaves, O(n * h) time and extra memory, and the work of building each path is wasted because only its length matters.",
			KeyInsight:  "The longest path from a node goes through whichever child has the longer path, so a node only needs one number from each child (its depth), not the paths themselves. That turns the problem into a single postorder pass.",
			WhyItWorks:  "Proof by induction on the tree's size. Base case: the empty tree has no nodes on any path, so 0 is correct. Step: for a non-empty tree, every root-to-leaf path is the root followed by a root-to-leaf path of the left subtree or of the right subtree. By the induction hypothesis the recursive calls return the correct depth of each subtree (they are smaller trees), so the longest path has 1 + max(left, right) nodes. Each node is visited once, giving O(n) time; the call stack holds the current path, O(h).",
			CommonMistakes: []string{
				"Returning max(left_depth, right_depth) without the + 1: every answer becomes 0. Add 1 for the current node.",
				"Returning 1 for a None node, which counts a phantom level under every leaf: the empty tree must return 0.",
				"Only following the left child (or stopping at the first leaf): the depth is the maximum over all paths, so ask both children.",
			},
			TimeComplexity:  "O(n): each node is visited exactly once",
			SpaceComplexity: "O(h) for the call stack: about log n for a balanced tree, n for a chain",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle the empty tree", Explanation: "A missing subtree contributes no nodes to any path, so it returns 0. This also stops the recursion below every leaf.", CodeSnippet: "if root is None:\n    return 0", LineStart: 2, LineEnd: 3},
				{Title: "Ask the children", Explanation: "Trust the recursive call: maxDepth on each smaller subtree returns its correct depth.", CodeSnippet: "left_depth = maxDepth(root.left)\nright_depth = maxDepth(root.right)", LineStart: 4, LineEnd: 5},
				{Title: "Combine and return to the parent", Explanation: "The longest path from this node is this node plus the deeper child's path.", CodeSnippet: "return 1 + max(left_depth, right_depth)", LineStart: 6, LineEnd: 6},
			},
			Alternatives: []AltSolution{
				{
					Name: "BFS counting levels",
					Code: `from collections import deque

def maxDepth(root):
    if root is None:
        return 0
    depth = 0
    queue = deque([root])
    while queue:
        depth += 1
        for _ in range(len(queue)):
            node = queue.popleft()
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
    return depth`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(w), the widest level",
					Explanation:     "Process the tree one level at a time and count the levels. It avoids deep recursion on a long chain, at the cost of holding a whole level in the queue.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the base case return 0 and not 1?", Answer: "The base case is the empty tree, which has no nodes, so no path through it adds anything. Returning 1 would count an imaginary node below every leaf and make every depth one too large."},
			{Prompt: "What would you change to compute the minimum depth (the shortest root-to-leaf path)?", Answer: "You cannot just swap max for min, because a node with one missing child would report 1 + 0 using the empty side, which is not a leaf. When exactly one child is None, return 1 + the depth of the other child; only when both exist take 1 + min of the two. BFS is also natural here: the first leaf you pop is at the minimum depth."},
		},
	},
	{
		ID:              "invert-binary-tree",
		Number:          226,
		Order:           2,
		Title:           "Invert Binary Tree",
		Slug:            "invert-binary-tree",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "BFS", "Binary Tree"},
		RelatedChapters: []int{9, 8},
		Description: `You are given the root of a binary tree. **Invert** it, meaning mirror it left to right: at every node, the left and right subtrees trade places. Return the root of the inverted tree.

You may change the tree in place and return the same root. An empty tree stays empty.`,
		Constraints: []string{
			"The tree has between 0 and 100 nodes",
			"-100 <= Node.val <= 100",
		},
		Examples: []Example{
			{Input: "root = [4,2,7,1,3,6,9]", Output: "[4,7,2,9,6,3,1]", Explanation: "The subtrees rooted at 2 and 7 swap sides, and inside each of them the children swap too: 1 and 3 become 3 and 1, 6 and 9 become 9 and 6."},
			{Input: "root = [2,1,3]", Output: "[2,3,1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{4, 2, 7, 1, 3, 6, 9}}, Expected: []interface{}{4, 7, 2, 9, 6, 3, 1}, Note: "example 1: full tree of three levels"},
			{Input: map[string]interface{}{"root": []interface{}{2, 1, 3}}, Expected: []interface{}{2, 3, 1}, Note: "example 2: one level of swapping"},
			{Input: map[string]interface{}{"root": []interface{}{}}, Expected: []interface{}{}, Hidden: true, Note: "empty tree"},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: []interface{}{1}, Hidden: true, Note: "single node is its own mirror"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2}}, Expected: []interface{}{1, nil, 2}, Hidden: true, Note: "a lone left child must move to the right (catches overwriting root.left before reading it)"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, nil, nil, 5}}, Expected: []interface{}{1, 3, 2, 5, nil, nil, 4}, Hidden: true, Note: "uneven tree: deeper levels must be inverted too, not just the root's children"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef invertTree(root):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
			ReturnType: "TreeNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If both subtrees of the root were already mirrored, what is the only thing left to do at the root?"},
			{Level: 2, Type: "insight", Content: "Mirroring a tree means mirroring each subtree and then swapping them. Use the DFS template from Chapter 9: ask each child to invert itself, then attach the results on the opposite sides and return this node to the parent.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Handle the empty tree: an empty tree inverts to an empty tree, so return None.\n2. Ask the children: invert the left subtree and the right subtree, and keep both results in variables before changing any pointer.\n3. Combine and return to the parent: put the inverted right subtree on the left and the inverted left subtree on the right, then return this node."},
			{Level: 4, Type: "skeleton", Content: "def invertTree(root):\n    if root is None:\n        return None\n    left_inverted = invertTree(root.left)\n    right_inverted = ____\n    root.left = ____\n    root.right = ____\n    return root"},
		},
		Solution: Solution{
			Code: `def invertTree(root):
    if root is None:
        return None
    left_inverted = invertTree(root.left)
    right_inverted = invertTree(root.right)
    root.left = right_inverted
    root.right = left_inverted
    return root`,
			Explanation: "Builds on Maximum Depth: the same DFS template, but now the value returned to the parent is a node (the root of the inverted subtree) instead of a number. Each node inverts its two subtrees through recursive calls and then swaps them.",
			BruteForce:  "You could read the tree into a level-order list, reverse each level's slice, and build a brand-new tree from it. That works in O(n) time but needs O(n) extra memory and fiddly bookkeeping for the None gaps, all to do something that only requires swapping two pointers per node.",
			KeyInsight:  "The mirror of a tree is the root with its two mirrored subtrees exchanged. So each node needs exactly one thing from each child (its inverted subtree) and does one swap.",
			WhyItWorks:  "Induction on size. The empty tree is its own mirror. For a node, the recursive calls (on smaller trees) return correctly mirrored subtrees; placing the mirrored right subtree on the left and the mirrored left subtree on the right is exactly the definition of mirroring at this node. Storing both results before assigning avoids losing a subtree. Each node is swapped once: O(n) time, O(h) stack.",
			CommonMistakes: []string{
				"Writing root.left = invertTree(root.right) and then root.right = invertTree(root.left): the second line reads the already-overwritten root.left. Save both results first, or use a tuple assignment.",
				"Swapping only the root's two children and not recursing: the deeper levels stay unmirrored. Every node must swap.",
				"Forgetting to return root at the end, so the caller gets None and the tree disappears.",
			},
			TimeComplexity:  "O(n): one swap per node",
			SpaceComplexity: "O(h) for the call stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle the empty tree", Explanation: "There is nothing to mirror below a missing child; returning None keeps that side empty.", CodeSnippet: "if root is None:\n    return None", LineStart: 2, LineEnd: 3},
				{Title: "Ask the children", Explanation: "Invert both subtrees first and keep the results in variables, so no pointer is changed before both are read.", CodeSnippet: "left_inverted = invertTree(root.left)\nright_inverted = invertTree(root.right)", LineStart: 4, LineEnd: 5},
				{Title: "Combine and return to the parent", Explanation: "Attach the mirrored subtrees on opposite sides and hand this node back up.", CodeSnippet: "root.left = right_inverted\nroot.right = left_inverted\nreturn root", LineStart: 6, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{
					Name: "Iterative BFS swap",
					Code: `from collections import deque

def invertTree(root):
    queue = deque([root] if root else [])
    while queue:
        node = queue.popleft()
        node.left, node.right = node.right, node.left
        if node.left:
            queue.append(node.left)
        if node.right:
            queue.append(node.right)
    return root`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(w), the widest level",
					Explanation:     "Visit every node in any order and swap its children. The order does not matter because each swap only touches that node's own two pointers.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Does it matter whether you swap before recursing (preorder) or after (postorder)?", Answer: "No, as long as each node's two pointers are swapped exactly once. Swapping first and then recursing on root.left and root.right still mirrors every subtree, because each subtree is mirrored independently of where it is attached. The BFS alternative shows the order is irrelevant."},
			{Prompt: "What happens if you invert the tree twice?", Answer: "You get the original tree back: mirroring is its own inverse, because every node's children are swapped twice. This is a good quick test for your implementation."},
		},
	},
	{
		ID:              "same-tree",
		Number:          100,
		Order:           3,
		Title:           "Same Tree",
		Slug:            "same-tree",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "BFS", "Binary Tree"},
		RelatedChapters: []int{9, 8},
		Description: `You are given the roots of two binary trees, p and q. Return True if they are the **same tree**, and False otherwise.

Two trees are the same when they have exactly the same shape and every pair of nodes in matching positions holds the same value. Two empty trees are the same.`,
		Constraints: []string{
			"Each tree has between 0 and 100 nodes",
			"-10^4 <= Node.val <= 10^4",
		},
		Examples: []Example{
			{Input: "p = [1,2,3], q = [1,2,3]", Output: "true", Explanation: "Both roots are 1, both left children are 2, both right children are 3, and all the children below are empty in both."},
			{Input: "p = [1,2], q = [1,null,2]", Output: "false", Explanation: "The values match but in p the 2 is a left child and in q it is a right child, so the shapes differ."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"p": []interface{}{1, 2, 3}, "q": []interface{}{1, 2, 3}}, Expected: true, Note: "example 1: identical trees"},
			{Input: map[string]interface{}{"p": []interface{}{1, 2}, "q": []interface{}{1, nil, 2}}, Expected: false, Note: "example 2: same values, different shape"},
			{Input: map[string]interface{}{"p": []interface{}{1, 2, 1}, "q": []interface{}{1, 1, 2}}, Expected: false, Hidden: true, Note: "same values in mirrored positions"},
			{Input: map[string]interface{}{"p": []interface{}{}, "q": []interface{}{}}, Expected: true, Hidden: true, Note: "two empty trees are the same"},
			{Input: map[string]interface{}{"p": []interface{}{1}, "q": []interface{}{}}, Expected: false, Hidden: true, Note: "one empty, one not (catches reading q.val on None)"},
			{Input: map[string]interface{}{"p": []interface{}{1, nil, 2, 3}, "q": []interface{}{1, nil, 2, nil, 3}}, Expected: false, Hidden: true, Note: "difference is two levels down (comparing only the roots' children fails)"},
			{Input: map[string]interface{}{"p": []interface{}{0, -5, 7}, "q": []interface{}{0, -5, 7}}, Expected: true, Hidden: true, Note: "zero and negative values"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef isSameTree(p, q):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"p": "TreeNode", "q": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If the two roots hold the same value, what else has to be true for the whole trees to be the same?"},
			{Level: 2, Type: "insight", Content: "Two trees are the same exactly when their roots match, their left subtrees are the same and their right subtrees are the same. Walk both trees together with the DFS template from Chapter 9; the only new part is a base case that handles two possibly-empty trees.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Handle the empty tree: if both are None they match; if only one is None they don't.\n2. Check this node: if the two values differ, return False.\n3. Ask the children: compare the left subtrees with each other and the right subtrees with each other.\n4. Combine and return to the parent: the trees are the same only if both comparisons are True."},
			{Level: 4, Type: "skeleton", Content: "def isSameTree(p, q):\n    if p is None and q is None:\n        return ____\n    if ____:\n        return False\n    if p.val != q.val:\n        return False\n    return ____"},
		},
		Solution: Solution{
			Code: `def isSameTree(p, q):
    if p is None and q is None:
        return True
    if p is None or q is None:
        return False
    if p.val != q.val:
        return False
    return isSameTree(p.left, q.left) and isSameTree(p.right, q.right)`,
			Explanation: "Builds on Maximum Depth and Invert Binary Tree by running the DFS template on two trees in lockstep. The new idea is the paired base case: two empty trees match, one empty and one non-empty don't. Each pair of nodes checks its own values and asks its children whether their subtrees match.",
			BruteForce:  "Serialize both trees to lists (with markers for missing children) and compare the lists. That is O(n) too, but it always builds two full lists even when the roots already differ, and forgetting the markers makes [1,2] and [1,null,2] look equal.",
			KeyInsight:  "Sameness is defined recursively, and the recursion can stop at the first mismatch. Short-circuit and returns False as soon as any pair differs, without looking at the rest.",
			WhyItWorks:  "Induction on the combined size. Two empty trees are the same; an empty and a non-empty tree are not. Otherwise the trees are the same exactly when the root values are equal and both pairs of subtrees are the same, and the recursive calls decide those smaller pairs correctly. Every pair of nodes in matching positions is compared at most once, so the time is O(min(n, m)) at worst O(n), with O(h) stack.",
			CommonMistakes: []string{
				"Checking p.val != q.val before the None checks, which crashes with AttributeError when exactly one is None. Handle the empty cases first.",
				"Comparing only a traversal's values without marking empty children: [2,1] and [1,null,2] both have inorder 1, 2 but are different trees. Compare structure, not just values.",
				"Returning True as soon as the roots match, without asking the children.",
			},
			TimeComplexity:  "O(n): each pair of matching nodes is compared once",
			SpaceComplexity: "O(h) for the call stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle the empty tree", Explanation: "Two empty trees match. If only one is empty, the shapes differ. Doing this first guarantees both are real nodes below.", CodeSnippet: "if p is None and q is None:\n    return True\nif p is None or q is None:\n    return False", LineStart: 2, LineEnd: 5},
				{Title: "Check this node", Explanation: "Different values at the same position mean the trees differ; stop early.", CodeSnippet: "if p.val != q.val:\n    return False", LineStart: 6, LineEnd: 7},
				{Title: "Combine and return to the parent", Explanation: "Ask the children pairwise and report True only if both subtrees match. and short-circuits, so a left mismatch skips the right side.", CodeSnippet: "return isSameTree(p.left, q.left) and isSameTree(p.right, q.right)", LineStart: 8, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{
					Name: "Compare preorder encodings with null markers",
					Code: `def isSameTree(p, q):
    def encode(node, out):
        if node is None:
            out.append("#")
            return
        out.append(str(node.val))
        encode(node.left, out)
        encode(node.right, out)
    first, second = [], []
    encode(p, first)
    encode(q, second)
    return first == second`,
					TimeComplexity:  "O(n + m)",
					SpaceComplexity: "O(n + m)",
					Explanation:     "A preorder listing that marks every missing child with # determines the tree uniquely, so equal encodings mean equal trees. It always walks both whole trees, unlike the recursive version, which stops at the first difference.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the None checks come before comparing p.val and q.val?", Answer: "If exactly one of them is None, reading .val on it raises an error. Checking both-None and then either-None first guarantees that both are real nodes by the time values are compared, and it also gives the right answer for those shape differences."},
			{Prompt: "How would you check whether one tree is the mirror image of another (Symmetric Tree)?", Answer: "Use the same lockstep recursion but pair the children crosswise: compare a.left with b.right and a.right with b.left. A tree is symmetric when its left and right subtrees are mirrors of each other."},
		},
	},
	{
		ID:              "subtree-of-another-tree",
		Number:          572,
		Order:           4,
		Title:           "Subtree of Another Tree",
		Slug:            "subtree-of-another-tree",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "String Matching", "Binary Tree", "Hash Function"},
		RelatedChapters: []int{9, 8},
		Description: `You are given the roots of two binary trees, root and subRoot. Return True if root contains a **subtree** that is the same tree as subRoot (same shape and same values), and False otherwise.

A subtree of a tree is one of its nodes together with **all** of that node's descendants. The whole tree counts as a subtree of itself. Matching only the top part of some node's descendants is not enough.`,
		Constraints: []string{
			"root has between 1 and 2000 nodes",
			"subRoot has between 1 and 1000 nodes",
			"-10^4 <= Node.val <= 10^4",
		},
		Examples: []Example{
			{Input: "root = [3,4,5,1,2], subRoot = [4,1,2]", Output: "true", Explanation: "The node 4 in root has children 1 and 2 and nothing below them, exactly like subRoot."},
			{Input: "root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]", Output: "false", Explanation: "Node 4 matches at the top, but its child 2 has an extra child 0, so the subtree at 4 is not the same tree as subRoot."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 4, 5, 1, 2}, "subRoot": []interface{}{4, 1, 2}}, Expected: true, Note: "example 1: match in the left subtree"},
			{Input: map[string]interface{}{"root": []interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0}, "subRoot": []interface{}{4, 1, 2}}, Expected: false, Note: "example 2: top matches but an extra descendant spoils it"},
			{Input: map[string]interface{}{"root": []interface{}{1, 1}, "subRoot": []interface{}{1}}, Expected: true, Hidden: true, Note: "root value matches but shape doesn't; the match is one level down (stopping at the first equal value fails)"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3}, "subRoot": []interface{}{1, 2}}, Expected: false, Hidden: true, Note: "subRoot is only the top part of root, not a whole subtree"},
			{Input: map[string]interface{}{"root": []interface{}{12}, "subRoot": []interface{}{2}}, Expected: false, Hidden: true, Note: "string encodings without separators would wrongly find \"2\" inside \"12\""},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, nil, nil, 4, 5, nil, nil, 6}, "subRoot": []interface{}{3, 4, 5, nil, nil, 6}}, Expected: true, Hidden: true, Note: "match deeper in the right subtree"},
			{Input: map[string]interface{}{"root": []interface{}{2}, "subRoot": []interface{}{2}}, Expected: true, Hidden: true, Note: "whole tree equals subRoot"},
		},
		TimeComplexity:  "O(n * m)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef isSubtree(root, subRoot):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode", "subRoot": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You already know how to test whether two trees are identical. At which nodes of root could a copy of subRoot start?"},
			{Level: 2, Type: "insight", Content: "A copy of subRoot can start at any node of root. So run Same Tree (problem 3) with each node of root as the candidate, using an outer DFS (Chapter 9) that asks: is it here, or in my left subtree, or in my right subtree?", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Reuse Same Tree: write a helper same(a, b) that compares two trees exactly.\n2. Handle the empty tree: an empty root cannot contain a non-empty subRoot, so return False.\n3. Check this node: if same(root, subRoot), return True.\n4. Ask the children: otherwise the copy must be inside the left subtree or the right subtree; return True if either recursive search finds it."},
			{Level: 4, Type: "skeleton", Content: "def isSubtree(root, subRoot):\n    def same(a, b):\n        if a is None and b is None:\n            return True\n        if a is None or b is None or a.val != b.val:\n            return False\n        return ____\n\n    if root is None:\n        return ____\n    if same(root, subRoot):\n        return True\n    return ____"},
		},
		Solution: Solution{
			Code: `def isSubtree(root, subRoot):
    def same(a, b):
        if a is None and b is None:
            return True
        if a is None or b is None or a.val != b.val:
            return False
        return same(a.left, b.left) and same(a.right, b.right)

    if root is None:
        return False
    if same(root, subRoot):
        return True
    return isSubtree(root.left, subRoot) or isSubtree(root.right, subRoot)`,
			Explanation: "Builds directly on Same Tree: it is two DFS passes nested inside each other. The outer DFS visits each node of root as a possible starting point; the inner DFS (Same Tree) checks whether the whole subtree there matches subRoot. The new idea is combining the children's answers with or (found anywhere) instead of and.",
			BruteForce:  "The approach shown is already the straightforward one: try every starting node and compare, O(n * m) in the worst case (for example, root is all 1s and subRoot is all 1s with one different leaf). It wastes work by re-comparing the same nodes of root against subRoot from many starting points.",
			KeyInsight:  "A subtree is fully determined by its top node, so there are only n candidates, and each candidate is a Same Tree question you already know how to answer. For a faster O(n + m) version, encode both trees with separators and null markers and do a substring search (see the alternative).",
			WhyItWorks:  "Every subtree of root is rooted at exactly one node, and the outer recursion visits every node: either the match is rooted here (checked by same), or it lies entirely inside the left subtree or entirely inside the right subtree (checked by the recursive calls). same is correct by the argument from Same Tree, which requires the entire shape below to match, so a match that is only the top part of a larger subtree is rejected. Time O(n * m), stack O(h).",
			CommonMistakes: []string{
				"Returning True as soon as a node's value equals subRoot.val: the whole subtree below must match. Call the full Same Tree check.",
				"Checking only the root, or only one side: the copy can start at any node, so combine both children's answers with or.",
				"Serializing to strings without separators or null markers, so \"12\" contains \"2\" or two different shapes encode alike. Prefix each value with a comma and mark every empty child.",
			},
			TimeComplexity:  "O(n * m) for n nodes in root and m in subRoot",
			SpaceComplexity: "O(h) for the call stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Reuse Same Tree", Explanation: "The inner helper is exactly the Same Tree solution: both empty match, one empty or different values don't, otherwise compare children pairwise.", CodeSnippet: "def same(a, b):\n    if a is None and b is None:\n        return True\n    return same(a.left, b.left) and same(a.right, b.right)", LineStart: 2, LineEnd: 7},
				{Title: "Handle the empty tree", Explanation: "subRoot has at least one node, so an empty part of root cannot contain it.", CodeSnippet: "if root is None:\n    return False", LineStart: 9, LineEnd: 10},
				{Title: "Check this node", Explanation: "Test whether the subtree rooted at this node is identical to subRoot.", CodeSnippet: "if same(root, subRoot):\n    return True", LineStart: 11, LineEnd: 12},
				{Title: "Ask the children", Explanation: "If not here, the match must lie wholly inside one child's subtree; or returns True if either side finds it.", CodeSnippet: "return isSubtree(root.left, subRoot) or isSubtree(root.right, subRoot)", LineStart: 13, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Serialize both trees and search for a substring",
					Code: `def isSubtree(root, subRoot):
    def encode(node, out):
        if node is None:
            out.append(",#")
            return
        out.append("," + str(node.val))
        encode(node.left, out)
        encode(node.right, out)
    big, small = [], []
    encode(root, big)
    encode(subRoot, small)
    return "".join(small) in "".join(big)`,
					TimeComplexity:  "O(n + m) with a linear-time string search, O(n * m) worst case with Python's in",
					SpaceComplexity: "O(n + m)",
					Explanation:     "A preorder encoding with null markers is unique, and a subtree's encoding appears as a contiguous block inside the whole tree's encoding. The leading comma on every value stops 2 from matching inside 12.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the serialization alternative need both null markers and a separator before each value?", Answer: "Without null markers, different shapes such as [1,2] and [1,null,2] produce the same preorder 1,2, so a false match is possible. Without a leading separator, the encoding of value 2 (\"2,#,#\") appears inside the encoding of 12 (\"12,#,#\"). With \",2,#,#\" against \",12,#,#\" the false match disappears."},
			{Prompt: "Why is the worst case O(n * m) and what input causes it?", Answer: "Each of the n nodes of root can start a Same Tree comparison that runs up to m steps before failing. A tree of all 1s with subRoot being all 1s except a different value at its deepest leaf makes almost every comparison go nearly to the end before finding the mismatch."},
		},
	},
	{
		ID:              "lowest-common-ancestor-bst",
		Number:          235,
		Order:           5,
		Title:           "Lowest Common Ancestor of a Binary Search Tree",
		Slug:            "lowest-common-ancestor-of-a-binary-search-tree",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Search Tree", "Binary Tree"},
		RelatedChapters: []int{9, 6},
		Description: `You are given the root of a **binary search tree** (BST) and two of its nodes, p and q. In a BST, every value in a node's left subtree is smaller than the node's value and every value in its right subtree is larger. All values are distinct.

Return the **lowest common ancestor** (LCA) of p and q: the deepest node that has both p and q in its subtree. A node counts as being in its own subtree, so if p is an ancestor of q, the answer is p.

In the tests, p and q are given by their values; the judge passes you the actual nodes with those values, and checks the value of the node you return.`,
		Constraints: []string{
			"The tree has between 2 and 10^5 nodes",
			"-10^9 <= Node.val <= 10^9",
			"All Node.val are unique",
			"p != q, and both p and q exist in the BST",
		},
		Examples: []Example{
			{Input: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8", Output: "6", Explanation: "2 is in the left subtree of 6 and 8 is in the right subtree, so 6 is the deepest node above both."},
			{Input: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4", Output: "2", Explanation: "4 is in the subtree of 2, and a node is its own ancestor, so the answer is 2."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, "p": 2, "q": 8}, Expected: 6, Note: "example 1: split at the root"},
			{Input: map[string]interface{}{"root": []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, "p": 2, "q": 4}, Expected: 2, Note: "example 2: p is an ancestor of q"},
			{Input: map[string]interface{}{"root": []interface{}{2, 1}, "p": 2, "q": 1}, Expected: 2, Hidden: true, Note: "two-node tree, answer is the root itself"},
			{Input: map[string]interface{}{"root": []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, "p": 3, "q": 5}, Expected: 4, Hidden: true, Note: "split two levels down on the left"},
			{Input: map[string]interface{}{"root": []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, "p": 9, "q": 7}, Expected: 8, Hidden: true, Note: "split on the right with p > q"},
			{Input: map[string]interface{}{"root": []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, "p": 0, "q": 5}, Expected: 2, Hidden: true, Note: "leaves on different sides of 2"},
			{Input: map[string]interface{}{"root": []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, "p": 9, "q": 0}, Expected: 6, Hidden: true, Note: "p larger than q at the root (code that assumes p < q fails)"},
		},
		TimeComplexity:  "O(h)",
		SpaceComplexity: "O(1)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef lowestCommonAncestor(root, p, q):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode", "p": "TreeNodeRef:root", "q": "TreeNodeRef:root"},
			ReturnType: "TreeNodeVal",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Standing at some node, if both p and q are smaller than it, where must they both be? And what does it mean if one is smaller and the other is larger?"},
			{Level: 2, Type: "insight", Content: "Use the BST ordering to pick a side (Chapter 9). While both values are on the same side of the current node, the LCA is on that side too. The first node where they split (or where the node equals p or q) is the LCA. No recursion over both children is needed: it is one walk down, like binary search.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Start at the root.\n2. Use the ordering to pick a side: if p and q are both smaller than the current value, move left; if both are larger, move right.\n3. Stop at the split point: otherwise one is on each side (or one of them is this node), so return the current node."},
			{Level: 4, Type: "skeleton", Content: "def lowestCommonAncestor(root, p, q):\n    node = root\n    while node:\n        if ____:\n            node = node.left\n        elif ____:\n            node = node.right\n        else:\n            return ____\n    return None"},
		},
		Solution: Solution{
			Code: `def lowestCommonAncestor(root, p, q):
    node = root
    while node:
        if p.val < node.val and q.val < node.val:
            node = node.left
        elif p.val > node.val and q.val > node.val:
            node = node.right
        else:
            return node
    return None`,
			Explanation: "The first BST problem in the ladder. Problems 1 to 4 had to visit every node; here the BST ordering from Chapter 9 lets us use the ordering to pick a side at each node, just like binary search in Chapter 6, so we walk a single path from the root and stop where p and q go different ways.",
			BruteForce:  "Ignore the ordering and use the general-tree method: recursively search both subtrees for p and q and return the node where they are found on different sides. That visits up to all n nodes, O(n), even though a single comparison at each node already tells you which side both targets are on.",
			KeyInsight:  "In a BST, comparing p and q with a node tells you which subtree each is in without searching. While both are on the same side, every common ancestor deeper down is on that side; the first node where they separate is the lowest one.",
			WhyItWorks:  "Invariant: the current node is always a common ancestor of p and q (true at the root). If both values are smaller, both nodes are in the left subtree, so the left child is also a common ancestor and is deeper; similarly on the right. When neither move applies, either p and q are on different sides or one of them is the current node. Any deeper node lies in only one subtree and so cannot contain both, which makes the current node the lowest common ancestor. Each step goes one level down: O(h) time, O(1) space.",
			CommonMistakes: []string{
				"Assuming p.val < q.val and testing only p against the left and q against the right: when p is larger the walk goes the wrong way. Test both values against the node in each condition, or swap so p is smaller.",
				"Searching both subtrees recursively as in a general tree: correct but O(n); it ignores the ordering that makes this O(h).",
				"Forgetting that a node is its own ancestor and continuing past p when q is below it: when node equals p or q, the else branch must return it.",
			},
			TimeComplexity:  "O(h): one step per level",
			SpaceComplexity: "O(1): only one pointer",
			Walkthrough: []WalkthroughStep{
				{Title: "Start at the root", Explanation: "The root is an ancestor of every node, so it is a valid common ancestor to start from.", CodeSnippet: "node = root\nwhile node:", LineStart: 2, LineEnd: 3},
				{Title: "Use the ordering to pick a side", Explanation: "If both targets are smaller, they are both in the left subtree; if both are larger, both are in the right subtree. Move there, keeping the invariant.", CodeSnippet: "if p.val < node.val and q.val < node.val:\n    node = node.left\nelif p.val > node.val and q.val > node.val:\n    node = node.right", LineStart: 4, LineEnd: 7},
				{Title: "Stop at the split point", Explanation: "Otherwise p and q are on different sides, or one of them is this node. No deeper node can hold both, so this is the LCA.", CodeSnippet: "else:\n    return node", LineStart: 8, LineEnd: 9},
			},
			Alternatives: []AltSolution{
				{
					Name: "General binary tree LCA (ignores the ordering)",
					Code: `def lowestCommonAncestor(root, p, q):
    if root is None or root is p or root is q:
        return root
    left = lowestCommonAncestor(root.left, p, q)
    right = lowestCommonAncestor(root.right, p, q)
    if left and right:
        return root
    return left if left else right`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(h)",
					Explanation:     "Each call returns p, q, their LCA, or None for its subtree. If both children report something, p and q are on different sides and this node is the LCA. Works on any binary tree, but visits every node.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the first node where p and q split guaranteed to be the lowest common ancestor, not just some common ancestor?", Answer: "At the split node, p and q are in different subtrees (or one is the node itself). Any node below it lies entirely in one of its subtrees, so it cannot have both p and q underneath. Therefore no deeper common ancestor exists."},
			{Prompt: "What changes if the tree is a general binary tree, not a BST?", Answer: "You can no longer tell from values which side p and q are on, so you must search both subtrees: the recursive alternative returns the node where the two searches come back non-empty on both sides. That costs O(n) instead of O(h)."},
		},
	},
	{
		ID:              "binary-tree-level-order-traversal",
		Number:          102,
		Order:           6,
		Title:           "Binary Tree Level Order Traversal",
		Slug:            "binary-tree-level-order-traversal",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "BFS", "Binary Tree"},
		RelatedChapters: []int{9},
		Description: `You are given the root of a binary tree. Return the values of its nodes **level by level**: a list of lists, where the first inner list holds the root's value, the second holds the values at depth 1 from left to right, and so on.

Return an empty list for an empty tree.`,
		Constraints: []string{
			"The tree has between 0 and 2000 nodes",
			"-1000 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "root = [3,9,20,null,null,15,7]", Output: "[[3],[9,20],[15,7]]", Explanation: "Depth 0 holds 3; depth 1 holds 9 and 20 from left to right; depth 2 holds 15 and 7, the children of 20."},
			{Input: "root = [1]", Output: "[[1]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 9, 20, nil, nil, 15, 7}}, Expected: [][]int{{3}, {9, 20}, {15, 7}}, Note: "example 1"},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: [][]int{{1}}, Note: "example 2: single node"},
			{Input: map[string]interface{}{"root": []interface{}{}}, Expected: [][]int{}, Hidden: true, Note: "empty tree gives []"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, nil, nil, 5}}, Expected: [][]int{{1}, {2, 3}, {4, 5}}, Hidden: true, Note: "a level whose nodes have different parents"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, nil, 3}}, Expected: [][]int{{1}, {2}, {3}}, Hidden: true, Note: "chain: one node per level"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, 5, 6, 7}}, Expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}}, Hidden: true, Note: "full tree: without a level-size snapshot, levels run together"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef levelOrder(root):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Which data structure hands items back in the same order they were put in, and how would you know where one level stops and the next begins?"},
			{Level: 2, Type: "insight", Content: "Use BFS with a queue (Chapter 9). At the start of each round the queue holds exactly one level, so snapshot its length and pop that many nodes; their children, appended at the back, form the next level.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Seed the queue: return [] for an empty tree, otherwise start the queue with the root.\n2. Process one level: take the queue's current length as the level size and pop exactly that many nodes, collecting their values.\n3. Enqueue the children: for each popped node, append its left then right child if they exist.\n4. Record the level: after the inner loop, append the collected values to the answer, and repeat while the queue is non-empty."},
			{Level: 4, Type: "skeleton", Content: "from collections import deque\n\ndef levelOrder(root):\n    if root is None:\n        return []\n    levels = []\n    queue = deque([root])\n    while queue:\n        level = []\n        for _ in range(____):\n            node = ____\n            level.append(node.val)\n            if node.left:\n                queue.append(node.left)\n            if node.right:\n                queue.append(node.right)\n        ____\n    return levels"},
		},
		Solution: Solution{
			Code: `from collections import deque

def levelOrder(root):
    if root is None:
        return []
    levels = []
    queue = deque([root])
    while queue:
        level = []
        for _ in range(len(queue)):
            node = queue.popleft()
            level.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        levels.append(level)
    return levels`,
			Explanation: "The first BFS problem of the ladder. Problems 1 to 5 went deep along one path at a time; this one goes wide, using the level-by-level BFS from Chapter 9. The key detail is the level-size snapshot, which splits one continuous queue into separate levels.",
			BruteForce:  "Compute the tree's height, then for each depth d run a fresh DFS from the root that collects the nodes at exactly depth d. Each pass walks down to depth d again, so a chain of n nodes costs 1 + 2 + ... + n = O(n^2) steps, re-visiting the top of the tree once per level.",
			KeyInsight:  "A FIFO queue naturally outputs nodes in order of depth, because a node's children are added after every node already waiting. Counting the queue at the start of a round tells you exactly how many of the waiting nodes belong to the current level.",
			WhyItWorks:  "Invariant: at the top of the while loop, the queue holds exactly the nodes of one depth d, in left-to-right order. It holds for d = 0 (just the root). The inner loop pops exactly those len(queue) nodes and appends their children, left before right, so afterwards the queue holds exactly depth d + 1 in left-to-right order. Each node is enqueued and dequeued once: O(n) time; the queue holds at most one level plus part of the next, O(n) in the worst case.",
			CommonMistakes: []string{
				"Not snapshotting the level size: popping while the queue is non-empty mixes the next level's nodes into the current list. Use for _ in range(len(queue)), which reads the length once.",
				"Using list.pop(0) instead of deque.popleft(): each pop shifts the whole list, O(n) per pop and O(n^2) overall.",
				"Putting None into the queue (or seeding it with an empty root) and then reading node.val: skip missing children and return [] for an empty tree.",
			},
			TimeComplexity:  "O(n): each node is enqueued and dequeued once",
			SpaceComplexity: "O(n): the queue can hold a whole level, up to about n/2 nodes",
			Walkthrough: []WalkthroughStep{
				{Title: "Seed the queue", Explanation: "An empty tree has no levels. Otherwise the queue starts with exactly level 0: the root.", CodeSnippet: "if root is None:\n    return []\nqueue = deque([root])", LineStart: 4, LineEnd: 7},
				{Title: "Process one level", Explanation: "range(len(queue)) is evaluated once, so it fixes how many nodes belong to this level even though children are being added behind them.", CodeSnippet: "for _ in range(len(queue)):\n    node = queue.popleft()\n    level.append(node.val)", LineStart: 8, LineEnd: 12},
				{Title: "Enqueue the children", Explanation: "Children go to the back, left first, so the next level comes out in left-to-right order.", CodeSnippet: "if node.left:\n    queue.append(node.left)\nif node.right:\n    queue.append(node.right)", LineStart: 13, LineEnd: 16},
				{Title: "Record the level", Explanation: "When the inner loop ends, this level is complete; store it and start the next round.", CodeSnippet: "levels.append(level)\nreturn levels", LineStart: 17, LineEnd: 18},
			},
			Alternatives: []AltSolution{
				{
					Name: "DFS with a depth index",
					Code: `def levelOrder(root):
    levels = []
    def visit(node, depth):
        if node is None:
            return
        if depth == len(levels):
            levels.append([])
        levels[depth].append(node.val)
        visit(node.left, depth + 1)
        visit(node.right, depth + 1)
    visit(root, 0)
    return levels`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(h) stack plus the output",
					Explanation:     "Preorder DFS passes the depth down and appends each value to the list for its depth. Left is visited before right, so each level still comes out left to right. Useful to see that levels are about depth, not about the queue.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does range(len(queue)) give the right count even though the loop body appends to the queue?", Answer: "range(len(queue)) is computed once, before the loop body runs, so it captures the number of nodes in the current level. Appending children afterwards changes the queue but not the range that was already created, so exactly the current level is popped."},
			{Prompt: "How would you return the levels from bottom to top, or in zigzag order?", Answer: "Bottom to top: build the same list and reverse it at the end (or insert each level at the front). Zigzag: keep a flag that flips every level and reverse the level list on odd levels before recording it. The BFS itself does not change; only how each level is recorded."},
		},
	},
	{
		ID:              "validate-binary-search-tree",
		Number:          98,
		Order:           7,
		Title:           "Validate Binary Search Tree",
		Slug:            "validate-binary-search-tree",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Search Tree", "Binary Tree"},
		RelatedChapters: []int{9, 8},
		Description: `You are given the root of a binary tree. Return True if it is a valid **binary search tree** (BST), and False otherwise.

A valid BST satisfies, for **every** node:
- every value in the node's left subtree is strictly less than the node's value,
- every value in the node's right subtree is strictly greater than the node's value,
- and both subtrees are valid BSTs themselves.

Duplicates are therefore not allowed.`,
		Constraints: []string{
			"The tree has between 1 and 10^4 nodes",
			"-2^31 <= Node.val <= 2^31 - 1",
		},
		Examples: []Example{
			{Input: "root = [2,1,3]", Output: "true", Explanation: "1 < 2 on the left and 3 > 2 on the right, and the leaves are trivially valid."},
			{Input: "root = [5,1,4,null,null,3,6]", Output: "false", Explanation: "The root's right child is 4, which is not greater than 5."},
			{Input: "root = [5,4,6,null,null,3,7]", Output: "false", Explanation: "Every parent-child pair looks fine (4 < 5 < 6 and 3 < 6 < 7), but 3 sits in the right subtree of 5 and is not greater than 5."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{2, 1, 3}}, Expected: true, Note: "example 1: small valid BST"},
			{Input: map[string]interface{}{"root": []interface{}{5, 1, 4, nil, nil, 3, 6}}, Expected: false, Note: "example 2: a direct child breaks the rule"},
			{Input: map[string]interface{}{"root": []interface{}{5, 4, 6, nil, nil, 3, 7}}, Expected: false, Hidden: true, Note: "trap: every parent-child pair is fine but 3 is in 5's right subtree (parent-only checks pass it)"},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: true, Hidden: true, Note: "single node"},
			{Input: map[string]interface{}{"root": []interface{}{2, 2, 2}}, Expected: false, Hidden: true, Note: "duplicates are not allowed (catches <= instead of <)"},
			{Input: map[string]interface{}{"root": []interface{}{-2147483648, nil, 2147483647}}, Expected: true, Hidden: true, Note: "values at the 32-bit limits (catches integer sentinels for the bounds)"},
			{Input: map[string]interface{}{"root": []interface{}{3, 1, 5, 0, 2, 4, 6}}, Expected: true, Hidden: true, Note: "full valid BST of three levels"},
			{Input: map[string]interface{}{"root": []interface{}{10, 5, 15, nil, nil, 6, 20}}, Expected: false, Hidden: true, Note: "6 in the right subtree of 10"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef isValidBST(root):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Look at [5,4,6,null,null,3,7]. Node 3 is smaller than its parent 6, which is fine locally. Which other node's rule does 3 break, and how could the check at node 3 know about it?"},
			{Level: 2, Type: "insight", Content: "Each node must lie in an open interval set by all of its ancestors: going left from a node caps the upper bound at that node's value, going right raises the lower bound. Pass the bounds down in the DFS (Chapter 9) instead of only comparing with the parent.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Start with no bounds: low = negative infinity, high = positive infinity at the root.\n2. Handle the empty tree: an empty subtree is valid.\n3. Check this node against its bounds: if its value is not strictly between low and high, return False.\n4. Pass bounds down: ask the left child with (low, node.val) and the right child with (node.val, high); the tree is valid only if both are."},
			{Level: 4, Type: "skeleton", Content: "def isValidBST(root):\n    def valid(node, low, high):\n        if node is None:\n            return True\n        if not (____):\n            return False\n        return valid(node.left, ____) and valid(node.right, ____)\n\n    return valid(root, float(\"-inf\"), float(\"inf\"))"},
		},
		Solution: Solution{
			Code: `def isValidBST(root):
    def valid(node, low, high):
        if node is None:
            return True
        if not (low < node.val < high):
            return False
        return valid(node.left, low, node.val) and valid(node.right, node.val, high)

    return valid(root, float("-inf"), float("inf"))`,
			Explanation: "Builds on the BST ordering from LCA of a BST and on the DFS template from problems 1 to 4, with one new twist: information flows down as well as up. Each call receives the range of values its subtree is allowed to hold, so the rule from every ancestor is enforced, not just the parent's.",
			BruteForce:  "For every node, scan its entire left subtree to check all values are smaller and its entire right subtree to check all are larger. Correct, but each node is re-scanned once per ancestor, O(n * h), which is O(n^2) on a chain. The common shortcut of comparing a node only with its children is fast but wrong: [5,4,6,null,null,3,7] passes it.",
			KeyInsight:  "Everything an ancestor demands of a node can be summarized as one open interval (low, high). Going left from node v tightens high to v; going right tightens low to v. Checking each node once against its interval enforces all ancestor constraints at once.",
			WhyItWorks:  "Invariant: valid(node, low, high) is called exactly with low = the largest ancestor value that node lies to the right of, and high = the smallest ancestor value it lies to the left of. The BST definition says node's value must be greater than every ancestor it is right of and smaller than every ancestor it is left of, which is the same as low < val < high. So the check passes at every node exactly when the tree is a BST. On [5,4,6,null,null,3,7], node 3 is called with (5, 6) and 3 < 5 fails, while the parent-only check compared 3 only with 6. Each node is checked once: O(n) time, O(h) stack.",
			CommonMistakes: []string{
				"Comparing a node only with its direct children (node.left.val < node.val < node.right.val): misses deeper violations such as the 3 in [5,4,6,null,null,3,7]. Pass bounds down from all ancestors.",
				"Using <= or >= so duplicates pass: [2,2,2] must be False. Use strict comparisons.",
				"Starting the bounds at -2**31 and 2**31 - 1 with strict comparisons, which rejects a node whose value is exactly a limit: start with float('-inf') and float('inf').",
			},
			TimeComplexity:  "O(n): each node is checked once",
			SpaceComplexity: "O(h) for the call stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle the empty tree", Explanation: "An empty subtree satisfies every bound.", CodeSnippet: "if node is None:\n    return True", LineStart: 3, LineEnd: 4},
				{Title: "Check this node against its bounds", Explanation: "low and high summarize every ancestor's rule; the value must be strictly inside.", CodeSnippet: "if not (low < node.val < high):\n    return False", LineStart: 5, LineEnd: 6},
				{Title: "Pass bounds down", Explanation: "Going left, this node's value becomes the new upper bound; going right, it becomes the new lower bound. The other bound is inherited unchanged.", CodeSnippet: "return valid(node.left, low, node.val) and valid(node.right, node.val, high)", LineStart: 7, LineEnd: 7},
				{Title: "Start with no bounds", Explanation: "The root may hold any value, so begin with infinite bounds, which also handles values at the 32-bit limits.", CodeSnippet: "return valid(root, float(\"-inf\"), float(\"inf\"))", LineStart: 9, LineEnd: 9},
			},
			Alternatives: []AltSolution{
				{
					Name: "Inorder traversal must be strictly increasing",
					Code: `def isValidBST(root):
    values = []
    def inorder(node):
        if node is None:
            return
        inorder(node.left)
        values.append(node.val)
        inorder(node.right)
    inorder(root)
    return all(values[i] < values[i + 1] for i in range(len(values) - 1))`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "A tree is a BST exactly when its inorder traversal is strictly increasing. Collect it and check neighbours. It uses O(n) extra memory; keeping only the previous value brings that down to O(h).",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the parent-only check accept [5,4,6,null,null,3,7], and what bounds does node 3 receive in the solution?", Answer: "The parent-only check compares 3 only with its parent 6 and sibling 7, and 3 < 6 < 7 holds. But 3 is in the right subtree of 5, so it must be greater than 5. In the solution, going right from 5 sets low = 5, then going left from 6 sets high = 6, so node 3 is checked against (5, 6) and fails."},
			{Prompt: "Why is \"inorder traversal is strictly increasing\" equivalent to being a valid BST?", Answer: "Inorder lists the left subtree, then the node, then the right subtree. If the tree is a BST, everything on the left is smaller and everything on the right larger, at every node, so the list is increasing. Conversely, if the list is strictly increasing, every value in a node's left subtree appears before it and so is smaller, and every value in its right subtree appears after it and so is larger, which is the BST definition."},
		},
	},
	{
		ID:              "kth-smallest-element-in-bst",
		Number:          230,
		Order:           8,
		Title:           "Kth Smallest Element in a BST",
		Slug:            "kth-smallest-element-in-a-bst",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Search Tree", "Binary Tree"},
		RelatedChapters: []int{9, 5},
		Description: `You are given the root of a binary search tree (left subtree values are smaller than the node, right subtree values are larger) and an integer k. Return the **kth smallest** value among all the node values, counting from 1 (k = 1 means the minimum).

k is always between 1 and the number of nodes.`,
		Constraints: []string{
			"The tree has n nodes, with 1 <= k <= n <= 10^4",
			"0 <= Node.val <= 10^4",
		},
		Examples: []Example{
			{Input: "root = [3,1,4,null,2], k = 1", Output: "1", Explanation: "The values in sorted order are 1, 2, 3, 4; the first is 1."},
			{Input: "root = [5,3,6,2,4,null,null,1], k = 3", Output: "3", Explanation: "Sorted order is 1, 2, 3, 4, 5, 6; the third is 3."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 1, 4, nil, 2}, "k": 1}, Expected: 1, Note: "example 1: the minimum"},
			{Input: map[string]interface{}{"root": []interface{}{5, 3, 6, 2, 4, nil, nil, 1}, "k": 3}, Expected: 3, Note: "example 2"},
			{Input: map[string]interface{}{"root": []interface{}{1}, "k": 1}, Expected: 1, Hidden: true, Note: "single node"},
			{Input: map[string]interface{}{"root": []interface{}{5, 3, 6, 2, 4, nil, nil, 1}, "k": 6}, Expected: 6, Hidden: true, Note: "k = n gives the maximum"},
			{Input: map[string]interface{}{"root": []interface{}{3, 1, 4, nil, 2}, "k": 2}, Expected: 2, Hidden: true, Note: "answer is in a right subtree of a left child (preorder or BFS order gives the wrong node)"},
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2, nil, 3, nil, 4}, "k": 3}, Expected: 3, Hidden: true, Note: "right-leaning chain"},
		},
		TimeComplexity:  "O(h + k)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef kthSmallest(root, k):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Which traversal order visits the values of a BST from smallest to largest, and do you need to finish it?"},
			{Level: 2, Type: "insight", Content: "Inorder traversal (left, node, right) of a BST visits values in sorted order (Chapter 9). Count nodes as you visit them and stop at the kth. Doing the inorder with an explicit stack (Chapter 5) makes stopping early easy.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Go left as far as possible: push every node on the way down the left edge; the top of the stack is the smallest unvisited value.\n2. Visit the next smallest: pop a node and count it; if it is the kth, return its value.\n3. Move to the right subtree: everything there is larger than the popped node but smaller than what is left on the stack, so continue from its right child."},
			{Level: 4, Type: "skeleton", Content: "def kthSmallest(root, k):\n    stack = []\n    node = root\n    while stack or node:\n        while node:\n            stack.append(node)\n            node = ____\n        node = ____\n        k -= 1\n        if k == 0:\n            return node.val\n        node = ____\n    return -1"},
		},
		Solution: Solution{
			Code: `def kthSmallest(root, k):
    stack = []
    node = root
    while stack or node:
        while node:
            stack.append(node)
            node = node.left
        node = stack.pop()
        k -= 1
        if k == 0:
            return node.val
        node = node.right
    return -1`,
			Explanation: "Builds on Validate BST, whose alternative showed that a BST's inorder traversal is sorted. Here we use that fact directly and add one twist: an iterative inorder with an explicit stack, so we can stop after k visits instead of traversing all n nodes.",
			BruteForce:  "Collect every value (any traversal), sort them, and return the (k-1)th. That is O(n log n) time and O(n) memory, and it ignores that the BST is already sorted: the sort redoes work the tree's structure has done, and it reads all n values even when k = 1.",
			KeyInsight:  "Inorder traversal of a BST yields values in increasing order, so the kth node visited in inorder is the kth smallest. With an explicit stack the traversal can stop the moment the count reaches k.",
			WhyItWorks:  "Invariant: the stack holds, from bottom to top, the ancestors whose left subtrees are being explored, and the top of the stack (after going left as far as possible) is the smallest value not yet visited. Popping it and then moving to its right subtree visits values in exactly inorder, which for a BST is ascending order. So when k reaches 0 the popped node is the kth smallest. The walk reaches the leftmost node in O(h) and then makes O(k) further pops and pushes, for O(h + k) time and O(h) stack.",
			CommonMistakes: []string{
				"Using preorder or BFS order and counting k nodes: those orders are not sorted. Only inorder visits a BST in ascending order.",
				"Off-by-one in counting (returning when k == 1 before decrementing, or treating k as 0-based): decrement once per visit and return when it reaches 0.",
				"Collecting the full inorder list and indexing it: correct, but it always visits all n nodes even when k is small.",
			},
			TimeComplexity:  "O(h + k)",
			SpaceComplexity: "O(h) for the stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Go left as far as possible", Explanation: "Push the whole left edge. The last node pushed has no left child, so it is the smallest value not yet visited.", CodeSnippet: "while node:\n    stack.append(node)\n    node = node.left", LineStart: 5, LineEnd: 7},
				{Title: "Visit the next smallest", Explanation: "Pop the top: it is the next value in ascending order. Count it and return if it is the kth.", CodeSnippet: "node = stack.pop()\nk -= 1\nif k == 0:\n    return node.val", LineStart: 8, LineEnd: 11},
				{Title: "Move to the right subtree", Explanation: "Values in the right subtree come right after this node in sorted order, before the ancestors still on the stack.", CodeSnippet: "node = node.right", LineStart: 12, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Collect all values and sort",
					Code: `def kthSmallest(root, k):
    values = []
    stack = [root]
    while stack:
        node = stack.pop()
        if node:
            values.append(node.val)
            stack.append(node.left)
            stack.append(node.right)
    values.sort()
    return values[k - 1]`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Ignores the BST ordering: gather every value in any order, sort, and index. Simple and a good correctness check, but slower and it always reads the whole tree.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the node popped from the stack always the smallest value not yet visited?", Answer: "Before popping, the code walks left until there is no left child, pushing every node on the way. The top node has no unvisited left subtree, so nothing smaller than it remains unvisited in its subtree; the nodes below it on the stack are ancestors it lies to the left of, so they are larger. Its right subtree is larger too. So it is the minimum of everything not yet visited."},
			{Prompt: "If the tree is modified often (inserts and deletes) and kthSmallest is called many times, how could you make each query faster?", Answer: "Store in every node the size of its subtree and keep it updated on insert and delete. Then at a node with left subtree size L: if k <= L go left, if k == L + 1 the node is the answer, otherwise go right with k - L - 1. Each query becomes O(h), which is O(log n) for a balanced tree."},
		},
	},
	{
		ID:              "construct-binary-tree-from-preorder-inorder",
		Number:          105,
		Order:           9,
		Title:           "Construct Binary Tree from Preorder and Inorder Traversal",
		Slug:            "construct-binary-tree-from-preorder-and-inorder-traversal",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Array", "Hash Table", "Divide and Conquer", "Tree", "Binary Tree"},
		RelatedChapters: []int{9, 8, 2},
		Description: `You are given two integer lists, preorder and inorder, which are the preorder traversal (node, then left subtree, then right subtree) and the inorder traversal (left subtree, then node, then right subtree) of the same binary tree. All values are distinct.

Rebuild that tree and return its root. Create nodes with the provided TreeNode(val, left, right) class. (The judge prints your tree in level order with null for missing children.)`,
		Constraints: []string{
			"1 <= preorder.length <= 3000",
			"inorder.length == preorder.length",
			"-3000 <= preorder[i], inorder[i] <= 3000",
			"All values are unique, and both lists are traversals of the same tree",
		},
		Examples: []Example{
			{Input: "preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]", Output: "[3,9,20,null,null,15,7]", Explanation: "preorder[0] = 3 is the root. In inorder, 9 is left of 3 (the left subtree) and 15, 20, 7 are right of it (the right subtree). Repeat on each side: the right subtree's preorder is 20, 15, 7, so 20 is its root with 15 on the left and 7 on the right."},
			{Input: "preorder = [-1], inorder = [-1]", Output: "[-1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"preorder": []int{3, 9, 20, 15, 7}, "inorder": []int{9, 3, 15, 20, 7}}, Expected: []interface{}{3, 9, 20, nil, nil, 15, 7}, Note: "example 1"},
			{Input: map[string]interface{}{"preorder": []int{-1}, "inorder": []int{-1}}, Expected: []interface{}{-1}, Note: "example 2: single node"},
			{Input: map[string]interface{}{"preorder": []int{1, 2, 3}, "inorder": []int{3, 2, 1}}, Expected: []interface{}{1, 2, nil, 3}, Hidden: true, Note: "left chain: every right subtree is empty"},
			{Input: map[string]interface{}{"preorder": []int{1, 2, 3}, "inorder": []int{1, 2, 3}}, Expected: []interface{}{1, nil, 2, nil, 3}, Hidden: true, Note: "right chain: every left subtree is empty"},
			{Input: map[string]interface{}{"preorder": []int{1, 2, 4, 5, 3, 6, 7}, "inorder": []int{4, 2, 5, 1, 6, 3, 7}}, Expected: []interface{}{1, 2, 3, 4, 5, 6, 7}, Hidden: true, Note: "full tree of three levels"},
			{Input: map[string]interface{}{"preorder": []int{3, 1, 2, 4}, "inorder": []int{1, 2, 3, 4}}, Expected: []interface{}{3, 1, 4, nil, 2}, Hidden: true, Note: "left subtree has a right child (building right before left consumes preorder in the wrong order)"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def buildTree(preorder, inorder):\n    # Build and return the root (the tree-node class is provided, see the description)\n    # Write your solution here\n    pass",
		Harness: Harness{
			ReturnType: "TreeNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Which position in preorder always holds the root? Once you know the root's value, what does its position in inorder tell you about the two subtrees?"},
			{Level: 2, Type: "insight", Content: "preorder[0] is the root. In inorder, everything left of the root is the left subtree and everything right of it is the right subtree, which also tells you how many nodes each side has. That splits the problem into two smaller copies of itself: divide and conquer (Chapter 8) producing a tree (Chapter 9). A dict from value to inorder index (Chapter 2) makes each split O(1).", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Index the inorder positions in a dict so you can find any root's position in O(1).\n2. Handle the empty tree: a range of inorder with no elements builds None.\n3. Take the next root from preorder: preorder lists roots in exactly the order a left-first recursion creates them, so keep a pointer and advance it.\n4. Split inorder and ask the children: build the left subtree from the inorder range left of the root, then the right subtree from the range to its right.\n5. Combine and return to the parent: attach both and return the new node."},
			{Level: 4, Type: "skeleton", Content: "def buildTree(preorder, inorder):\n    index_of = {val: i for i, val in enumerate(inorder)}\n    next_root = 0\n\n    def build(lo, hi):\n        nonlocal next_root\n        if lo > hi:\n            return None\n        root_val = preorder[next_root]\n        next_root += 1\n        root = TreeNode(root_val)\n        mid = ____\n        root.left = build(____)\n        root.right = build(____)\n        return root\n\n    return build(0, len(inorder) - 1)"},
		},
		Solution: Solution{
			Code: `def buildTree(preorder, inorder):
    index_of = {val: i for i, val in enumerate(inorder)}
    next_root = 0

    def build(lo, hi):
        nonlocal next_root
        if lo > hi:
            return None
        root_val = preorder[next_root]
        next_root += 1
        root = TreeNode(root_val)
        mid = index_of[root_val]
        root.left = build(lo, mid - 1)
        root.right = build(mid + 1, hi)
        return root

    return build(0, len(inorder) - 1)`,
			Explanation: "Builds on Kth Smallest and Validate BST, which read a tree through its inorder traversal; now we go the other way and rebuild a tree from two traversals. The DFS template now returns a newly created subtree to the parent, and the split is the divide step of divide and conquer from Chapter 8.",
			BruteForce:  "Recurse with list slices: the root is preorder[0], find it with inorder.index (O(n)), and recurse on preorder[1:mid+1], inorder[:mid] and preorder[mid+1:], inorder[mid+1:]. Each level copies slices and searches linearly, O(n^2) for a chain-shaped tree.",
			KeyInsight:  "Two facts remove the waste: a dict gives every root's inorder position in O(1), and preorder lists roots in exactly the order a left-subtree-first recursion needs them, so a single advancing pointer replaces all the preorder slicing. Only the inorder range (lo, hi) needs to be passed.",
			WhyItWorks:  "Preorder of a subtree is its root followed by the preorder of its left subtree, then the preorder of its right subtree. The recursion creates the root, then fully builds the left subtree, then the right subtree, so it consumes preorder values in exactly that order and next_root always points at the root of the subtree being built. Inorder places the root between its left and right subtrees, so the range lo..mid-1 holds exactly the left subtree's values and mid+1..hi the right's; an empty range means an empty subtree. By induction each call returns the correct subtree. Each value is processed once: O(n) time, O(n) for the dict and up to O(n) stack on a chain.",
			CommonMistakes: []string{
				"Building the right subtree before the left: the shared preorder pointer then hands left-subtree roots to the right side. Always build left first, matching preorder's order.",
				"Using inorder.index(root_val) on every call: correct but O(n) per call and O(n^2) overall. Precompute a dict of positions.",
				"Off-by-one in the ranges (for example build(lo, mid) for the left side), which puts the root inside its own subtree and recurses forever: the left range ends at mid - 1 and the right starts at mid + 1.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n) for the dict and recursion",
			Walkthrough: []WalkthroughStep{
				{Title: "Index the inorder positions", Explanation: "Precompute where each value sits in inorder so each split is an O(1) lookup; next_root walks through preorder.", CodeSnippet: "index_of = {val: i for i, val in enumerate(inorder)}\nnext_root = 0", LineStart: 2, LineEnd: 3},
				{Title: "Handle the empty tree", Explanation: "An empty inorder range means this subtree has no nodes.", CodeSnippet: "if lo > hi:\n    return None", LineStart: 7, LineEnd: 8},
				{Title: "Take the next root from preorder", Explanation: "The next unused preorder value is the root of the subtree being built now.", CodeSnippet: "root_val = preorder[next_root]\nnext_root += 1\nroot = TreeNode(root_val)", LineStart: 9, LineEnd: 11},
				{Title: "Split inorder and ask the children", Explanation: "Values left of mid form the left subtree, values right of mid the right subtree. Build left first so preorder is consumed in order.", CodeSnippet: "mid = index_of[root_val]\nroot.left = build(lo, mid - 1)\nroot.right = build(mid + 1, hi)", LineStart: 12, LineEnd: 14},
				{Title: "Combine and return to the parent", Explanation: "Both subtrees are attached; hand the finished subtree to the caller.", CodeSnippet: "return root", LineStart: 15, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{
					Name: "Slicing recursion",
					Code: `def buildTree(preorder, inorder):
    if not preorder:
        return None
    root = TreeNode(preorder[0])
    mid = inorder.index(preorder[0])
    root.left = buildTree(preorder[1:mid + 1], inorder[:mid])
    root.right = buildTree(preorder[mid + 1:], inorder[mid + 1:])
    return root`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n^2) worst case for the slices",
					Explanation:     "The direct translation of the idea: the left subtree has mid nodes, so its preorder is the mid values after the root. Easier to read, but each call copies lists and searches inorder linearly.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the left subtree be built before the right subtree in the pointer version?", Answer: "preorder lists the root, then the entire left subtree, then the entire right subtree. The shared pointer hands out roots in that order, so the code must request them in that order too. Building the right side first would give it the left subtree's root values."},
			{Prompt: "Could you rebuild the tree from preorder and postorder alone? What about if values could repeat?", Answer: "Not in general: preorder and postorder do not say whether a lone child is a left or a right child (for example [1,2] could be either shape), so inorder is what pins down the split. If values repeat, a value no longer identifies a unique inorder position, so the dict lookup is ambiguous and several trees can share the same traversals."},
		},
	},
	{
		ID:              "binary-tree-max-path-sum",
		Number:          124,
		Order:           10,
		Title:           "Binary Tree Maximum Path Sum",
		Slug:            "binary-tree-maximum-path-sum",
		Difficulty:      "Hard",
		Category:        "trees",
		Tags:            []string{"Dynamic Programming", "Tree", "DFS", "Binary Tree"},
		RelatedChapters: []int{9, 8},
		Description: `You are given the root of a non-empty binary tree whose node values may be negative. A **path** is a sequence of nodes where each consecutive pair is joined by an edge (parent to child or child to parent), and no node appears more than once. A path has at least one node and does not have to pass through the root.

The **path sum** is the sum of the values on the path. Return the largest path sum of any path in the tree.`,
		Constraints: []string{
			"The tree has between 1 and 3 * 10^4 nodes",
			"-1000 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "root = [1,2,3]", Output: "6", Explanation: "The path 2 -> 1 -> 3 uses every node: 2 + 1 + 3 = 6."},
			{Input: "root = [-10,9,20,null,null,15,7]", Output: "42", Explanation: "The best path is 15 -> 20 -> 7 with sum 42. Extending it up to -10 would only lower the sum, and a path through -10 cannot use both 15 and 7."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3}}, Expected: 6, Note: "example 1: path bends at the root"},
			{Input: map[string]interface{}{"root": []interface{}{-10, 9, 20, nil, nil, 15, 7}}, Expected: 42, Note: "example 2: best path avoids the root"},
			{Input: map[string]interface{}{"root": []interface{}{-3}}, Expected: -3, Hidden: true, Note: "single negative node: the path must be non-empty"},
			{Input: map[string]interface{}{"root": []interface{}{-2, -1}}, Expected: -1, Hidden: true, Note: "all negative: starting best at 0 wrongly returns 0"},
			{Input: map[string]interface{}{"root": []interface{}{-1, -2, -3}}, Expected: -1, Hidden: true, Note: "all negative: best is the single largest node"},
			{Input: map[string]interface{}{"root": []interface{}{2, -1}}, Expected: 2, Hidden: true, Note: "a negative child should be dropped, not added"},
			{Input: map[string]interface{}{"root": []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}}, Expected: 48, Hidden: true, Note: "longer path 7 -> 11 -> 4 -> 5 -> 8 -> 13"},
			{Input: map[string]interface{}{"root": []interface{}{10, 2, 10, 20, 1, nil, -25, nil, nil, nil, nil, 3, 4}}, Expected: 42, Hidden: true, Note: "returning both branches upward double-counts; the path 20 -> 2 -> 10 -> 10 is the answer"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef maxPathSum(root):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Every path has one highest node where it bends (or ends). If the path bends at node X, which pieces does it use from X's left and right subtrees? And can a parent of X extend that same path?"},
			{Level: 2, Type: "insight", Content: "Use the DFS template with a global answer (Chapter 9's diameter example). Each node returns to its parent the best downward path starting at it, using at most one child, because a parent can only extend one branch. Separately, update the global answer with the best path that bends here: node + best left gain + best right gain. Drop a child's gain when it is negative.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Handle the empty tree: an empty subtree offers a gain of 0.\n2. Ask the children: get the best downward gain from the left and from the right, and replace a negative gain with 0 (just don't use that branch).\n3. Update the global answer: the best path that bends at this node is node.val + left gain + right gain; compare it with the best so far, which starts at negative infinity.\n4. Combine and return to the parent: return node.val + the larger of the two gains, because a path going up can only continue down one side."},
			{Level: 4, Type: "skeleton", Content: "def maxPathSum(root):\n    best = float(\"-inf\")\n\n    def gain(node):\n        nonlocal best\n        if node is None:\n            return 0\n        left_gain = ____\n        right_gain = ____\n        best = max(best, ____)\n        return ____\n\n    gain(root)\n    return best"},
		},
		Solution: Solution{
			Code: `def maxPathSum(root):
    best = float("-inf")

    def gain(node):
        nonlocal best
        if node is None:
            return 0
        left_gain = max(gain(node.left), 0)
        right_gain = max(gain(node.right), 0)
        best = max(best, node.val + left_gain + right_gain)
        return node.val + max(left_gain, right_gain)

    gain(root)
    return best`,
			Explanation: "Builds on Maximum Depth (return the best single branch to the parent) and on the diameter worked example in Chapter 9 (update a global answer with the path that bends at each node). The new twists are negative values, which you drop by clamping a branch's gain at 0, and a global best that must start at negative infinity because the answer can be negative.",
			BruteForce:  "Treat the tree as an undirected graph and, from every node, walk every simple path, tracking the largest sum seen. There are O(n^2) paths (one per pair of endpoints) and each walk re-explores the same edges, O(n^2) time in total, though every path's best value depends only on a few numbers per node.",
			KeyInsight:  "Every path has a unique highest node where it bends. The best path bending at node X is X plus the best downward path into the left subtree (if positive) plus the best downward path into the right subtree (if positive). Those downward bests are exactly what the children return, so one postorder pass covers every possible bend point.",
			WhyItWorks:  "gain(node) returns the largest sum of a path that starts at node and goes straight down (possibly just node itself): it is node.val plus the better child gain, where a negative child gain is replaced by 0 meaning that branch is not taken. Any path in the tree has exactly one highest node X, and consists of X plus a downward path into at most each of X's subtrees, so its sum is at most X.val + max(left gain, 0) + max(right gain, 0), which is the value recorded at X. Taking the maximum over all X therefore finds the best path. A parent cannot reuse both branches, which is why the return value uses only max(left_gain, right_gain). Starting best at negative infinity lets an all-negative tree return its largest single node. O(n) time, O(h) stack.",
			CommonMistakes: []string{
				"Starting the global best at 0: on an all-negative tree like [-2,-1] the answer is -1, but the code returns 0. Start at float('-inf') (or at root.val).",
				"Returning node.val + left_gain + right_gain to the parent: a path that bends at this node cannot also continue upward. Return only one branch: node.val + max(left_gain, right_gain).",
				"Not clamping negative child gains at 0, so a node is forced to drag along a branch that lowers the sum ([2,-1] would give 1 instead of 2).",
			},
			TimeComplexity:  "O(n): one visit per node",
			SpaceComplexity: "O(h) for the call stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle the empty tree", Explanation: "A missing child offers nothing, so its gain is 0, which is the same as not taking that branch.", CodeSnippet: "if node is None:\n    return 0", LineStart: 6, LineEnd: 7},
				{Title: "Ask the children", Explanation: "Get the best downward path from each child; clamp at 0 so a negative branch is simply left out.", CodeSnippet: "left_gain = max(gain(node.left), 0)\nright_gain = max(gain(node.right), 0)", LineStart: 8, LineEnd: 9},
				{Title: "Update the global answer", Explanation: "The best path whose highest point is this node uses both branches. Record it; best started at negative infinity so all-negative trees work.", CodeSnippet: "best = max(best, node.val + left_gain + right_gain)", LineStart: 10, LineEnd: 10},
				{Title: "Combine and return to the parent", Explanation: "A path that continues up through the parent can only come down one side, so return the node plus its better branch.", CodeSnippet: "return node.val + max(left_gain, right_gain)", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: walk every simple path",
					Code: `def maxPathSum(root):
    neighbors = {}
    stack = [(root, None)]
    while stack:
        node, parent = stack.pop()
        neighbors[node] = [] if parent is None else [parent]
        if parent is not None:
            neighbors[parent].append(node)
        for child in (node.left, node.right):
            if child:
                stack.append((child, node))
    best = float("-inf")
    for start in neighbors:
        walk = [(start, None, start.val)]
        while walk:
            node, came_from, total = walk.pop()
            best = max(best, total)
            for nxt in neighbors[node]:
                if nxt is not came_from:
                    walk.append((nxt, node, total + nxt.val))
    return best`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n)",
					Explanation:     "Turn the tree into an undirected graph, then from every start node explore every simple path (in a tree, never going back where you came from is enough) and keep the best sum. It shares no logic with the gain recursion, which makes it a good check.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does gain return only one branch while the global update uses both?", Answer: "The global update considers the path whose highest point is this node, which may come up one side and go down the other. But the value returned is used by the parent to extend a path upward through this node, and such a path enters this node from above and can leave through only one child, otherwise it would visit this node twice."},
			{Prompt: "What does the code return for [-2,-1], and which line makes that right?", Answer: "It returns -1. At node -1 both child gains are 0, so best becomes -1; at node -2 the left gain is max(-1, 0) = 0, so the bend there gives -2 and best stays -1. Starting best at float('-inf') instead of 0 is what allows a negative answer."},
		},
	},
	{
		ID:              "serialize-and-deserialize-binary-tree",
		Number:          297,
		Order:           11,
		Title:           "Serialize and Deserialize Binary Tree",
		Slug:            "serialize-and-deserialize-binary-tree",
		Difficulty:      "Hard",
		Category:        "trees",
		Tags:            []string{"String", "Tree", "DFS", "BFS", "Design", "Binary Tree"},
		RelatedChapters: []int{9, 8},
		Description: `Design a class Codec with two methods:
- serialize(root) turns a binary tree into a string,
- deserialize(data) turns such a string back into a tree.

Any format is allowed, as long as deserialize(serialize(root)) rebuilds a tree with exactly the same shape and values as the original. The judge calls serialize on the input tree, checks that it returned a string, passes that string to deserialize, and compares the tree you rebuild with the original. Values may be negative and may repeat, and the tree may be empty.`,
		Constraints: []string{
			"The tree has between 0 and 10^4 nodes",
			"-1000 <= Node.val <= 1000",
			"serialize must return a str",
		},
		Examples: []Example{
			{Input: "root = [1,2,3,null,null,4,5]", Output: "[1,2,3,null,null,4,5]", Explanation: "One possible encoding is the preorder with # for empty children: \"1,2,#,#,3,4,#,#,5,#,#\". Reading it back in the same order rebuilds the same tree."},
			{Input: "root = []", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, nil, nil, 4, 5}}, Expected: []interface{}{1, 2, 3, nil, nil, 4, 5}, Note: "example 1"},
			{Input: map[string]interface{}{"root": []interface{}{}}, Expected: []interface{}{}, Note: "example 2: empty tree"},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: []interface{}{1}, Hidden: true, Note: "single node"},
			{Input: map[string]interface{}{"root": []interface{}{-1000, 1000}}, Expected: []interface{}{-1000, 1000}, Hidden: true, Note: "negative and multi-digit values (catches parsing one character at a time)"},
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2, nil, 3}}, Expected: []interface{}{1, nil, 2, nil, 3}, Hidden: true, Note: "right chain: shape depends on the null markers"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, nil, 3, nil, 4}}, Expected: []interface{}{1, 2, nil, 3, nil, 4}, Hidden: true, Note: "left chain"},
			{Input: map[string]interface{}{"root": []interface{}{5, 5, 5}}, Expected: []interface{}{5, 5, 5}, Hidden: true, Note: "repeated values"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `# TreeNode(val, left, right) is provided
class Codec:
    def serialize(self, root):
        # Return a string that encodes the tree
        pass

    def deserialize(self, data):
        # Rebuild the tree from the string and return its root
        pass`,
		Harness: Harness{
			Mode:       "roundtrip",
			Function:   "serialize,deserialize",
			ClassName:  "Codec",
			ParamTypes: map[string]string{"root": "TreeNode"},
			ReturnType: "TreeNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In Construct from Preorder and Inorder you needed two traversals to pin down the shape. What could you write into a single traversal so that the shape is unambiguous on its own?"},
			{Level: 2, Type: "insight", Content: "Write a preorder traversal that also records every empty child with a marker such as #. With the markers, a preorder listing determines the tree uniquely, and the reader can rebuild it with the same recursion (Chapter 9's DFS template), consuming tokens in the order the writer produced them.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Mark empty trees: when serializing, write # for every None child.\n2. Write node then children: write the value, then recursively the left subtree, then the right subtree (preorder).\n3. Join into a string with a separator such as a comma, so multi-digit and negative values stay intact.\n4. Read in the same order: split the string into tokens and read them one at a time; a # means return None.\n5. Rebuild node then children: create the node, then build its left subtree from the following tokens, then its right subtree, and return it."},
			{Level: 4, Type: "skeleton", Content: "class Codec:\n    def serialize(self, root):\n        tokens = []\n\n        def write(node):\n            if node is None:\n                ____\n                return\n            tokens.append(str(node.val))\n            ____\n            ____\n\n        write(root)\n        return \",\".join(tokens)\n\n    def deserialize(self, data):\n        tokens = iter(data.split(\",\"))\n\n        def read():\n            token = next(tokens)\n            if token == \"#\":\n                return ____\n            node = TreeNode(int(token))\n            node.left = ____\n            node.right = ____\n            return node\n\n        return read()"},
		},
		Solution: Solution{
			Code: `class Codec:
    def serialize(self, root):
        tokens = []

        def write(node):
            if node is None:
                tokens.append("#")
                return
            tokens.append(str(node.val))
            write(node.left)
            write(node.right)

        write(root)
        return ",".join(tokens)

    def deserialize(self, data):
        tokens = iter(data.split(","))

        def read():
            token = next(tokens)
            if token == "#":
                return None
            node = TreeNode(int(token))
            node.left = read()
            node.right = read()
            return node

        return read()`,
			Explanation: "The capstone of the ladder. Construct from Preorder and Inorder showed that a plain preorder is ambiguous and needed inorder to split it; here we make preorder unambiguous by writing a marker for every empty child. Serialize is a preorder DFS; deserialize is the same DFS run on a stream of tokens, the way buildTree consumed preorder with a pointer.",
			BruteForce:  "Store the tree as a full level-order array where position i has children at 2i+1 and 2i+2, padding every missing node with null. That is simple, but a chain of n nodes needs about 2^n slots, so both time and space blow up exponentially for unbalanced trees.",
			KeyInsight:  "A preorder traversal that includes a marker for every empty child determines the tree uniquely: when reading, a value always starts a new subtree, and a marker always closes one. The reader can follow the same recursion as the writer, so each token is read exactly once.",
			WhyItWorks:  "Claim: read() consumes exactly the tokens that write() produced for one subtree and rebuilds that subtree. For an empty tree write() produced a single #, and read() consumes it and returns None. For a node, write() produced its value, then the tokens of its left subtree, then those of its right subtree; read() consumes the value, then (by induction) exactly the left subtree's tokens, then exactly the right subtree's tokens, and attaches them in the same places. So deserialize(serialize(t)) equals t. Each node and each empty child produces one token, 2n + 1 tokens in total: O(n) time and space.",
			CommonMistakes: []string{
				"Leaving out the null markers: [1,2] and [1,null,2] both serialize to \"1,2\" and cannot be told apart. Write a marker for every empty child.",
				"Joining without a separator or parsing character by character: -1000 and 12 span several characters. Join with commas and split on them.",
				"Building the right subtree before the left in deserialize, or using a different order than serialize: the reader must consume tokens in exactly the order they were written.",
			},
			TimeComplexity:  "O(n) for each method",
			SpaceComplexity: "O(n) for the string and tokens, plus O(h) recursion",
			Walkthrough: []WalkthroughStep{
				{Title: "Mark empty trees", Explanation: "Every None child becomes a # token, so the shape is recorded along with the values.", CodeSnippet: "if node is None:\n    tokens.append(\"#\")\n    return", LineStart: 6, LineEnd: 8},
				{Title: "Write node then children", Explanation: "Preorder: this node's value, then the whole left subtree, then the whole right subtree.", CodeSnippet: "tokens.append(str(node.val))\nwrite(node.left)\nwrite(node.right)", LineStart: 9, LineEnd: 11},
				{Title: "Join into a string", Explanation: "Commas separate tokens so multi-digit and negative values survive the round trip.", CodeSnippet: "write(root)\nreturn \",\".join(tokens)", LineStart: 13, LineEnd: 14},
				{Title: "Read in the same order", Explanation: "An iterator hands out tokens one at a time, exactly in the order serialize wrote them; # closes an empty subtree.", CodeSnippet: "tokens = iter(data.split(\",\"))\ntoken = next(tokens)\nif token == \"#\":\n    return None", LineStart: 17, LineEnd: 22},
				{Title: "Rebuild node then children", Explanation: "Create the node, then let the recursion rebuild the left subtree from the next tokens and the right subtree after it.", CodeSnippet: "node = TreeNode(int(token))\nnode.left = read()\nnode.right = read()\nreturn node", LineStart: 23, LineEnd: 26},
			},
			Alternatives: []AltSolution{
				{
					Name: "Level-order (BFS) encoding",
					Code: `from collections import deque

class Codec:
    def serialize(self, root):
        tokens = []
        queue = deque([root])
        while queue:
            node = queue.popleft()
            if node is None:
                tokens.append("#")
                continue
            tokens.append(str(node.val))
            queue.append(node.left)
            queue.append(node.right)
        return " ".join(tokens)

    def deserialize(self, data):
        tokens = data.split(" ")
        if tokens[0] == "#":
            return None
        root = TreeNode(int(tokens[0]))
        queue = deque([root])
        i = 1
        while queue:
            node = queue.popleft()
            if tokens[i] != "#":
                node.left = TreeNode(int(tokens[i]))
                queue.append(node.left)
            if tokens[i + 1] != "#":
                node.right = TreeNode(int(tokens[i + 1]))
                queue.append(node.right)
            i += 2
        return root`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n)",
					Explanation:     "The format LeetCode itself uses: BFS writes each node's two children (or #) in level order, and the reader rebuilds level by level, giving each dequeued node the next two tokens. No recursion, so very deep trees are no problem.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is preorder with null markers enough to rebuild the tree, when plain preorder needed inorder as well?", Answer: "Plain preorder does not say where the left subtree ends and the right begins, so [1,2] could be 2 on the left or on the right. With markers, every subtree's encoding ends exactly when its last empty child's # is read, so the reader always knows when a subtree is complete and the split is unambiguous."},
			{Prompt: "How many tokens does serialize produce for a tree of n nodes, and why?", Answer: "2n + 1: one per node, plus one # for each empty child. A binary tree with n nodes has n + 1 empty child slots (each node has 2 child slots, 2n in total, and n - 1 of them are filled by the non-root nodes)."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, TreesProblems...)
}
