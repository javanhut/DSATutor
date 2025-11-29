package practice

// TreesProblems contains all trees category problems
var TreesProblems = []*Problem{
	{
		ID:         "invert-binary-tree",
		Number:     26,
		Title:      "Invert Binary Tree",
		Difficulty: "Easy",
		Category:   "trees",
		Tags:       []string{"Tree", "Binary Tree", "DFS", "BFS"},
		RelatedChapters: []int{9, 10},
		Description: `Given the root of a binary tree, invert the tree, and return its root.`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [0, 100]",
			"-100 <= Node.val <= 100",
		},
		Examples: []Example{
			{Input: "root = [4,2,7,1,3,6,9]", Output: "[4,7,2,9,6,3,1]"},
			{Input: "root = [2,1,3]", Output: "[2,3,1]"},
			{Input: "root = []", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{4, 2, 7, 1, 3, 6, 9}}, Expected: []interface{}{4, 7, 2, 9, 6, 3, 1}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def invertTree(root):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Recursively swap left and right children."},
		},
		Solution: Solution{
			Code: `def invertTree(root):
    if not root:
        return None
    root.left, root.right = root.right, root.left
    invertTree(root.left)
    invertTree(root.right)
    return root`,
			Explanation: "Swap left and right children, then recursively invert subtrees.",
		},
	},
	{
		ID:         "maximum-depth-of-binary-tree",
		Number:     27,
		Title:      "Maximum Depth of Binary Tree",
		Difficulty: "Easy",
		Category:   "trees",
		Tags:       []string{"Tree", "DFS", "BFS", "Binary Tree"},
		RelatedChapters: []int{9, 10},
		Description: `Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [0, 10^4]",
			"-100 <= Node.val <= 100",
		},
		Examples: []Example{
			{Input: "root = [3,9,20,null,null,15,7]", Output: "3"},
			{Input: "root = [1,null,2]", Output: "2"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 9, 20, nil, nil, 15, 7}}, Expected: 3},
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2}}, Expected: 2},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def maxDepth(root):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Recursively find max depth of left and right subtrees."},
		},
		Solution: Solution{
			Code: `def maxDepth(root):
    if not root:
        return 0
    return 1 + max(maxDepth(root.left), maxDepth(root.right))`,
			Explanation: "Base case: empty tree has depth 0. Otherwise, 1 + max of subtree depths.",
		},
	},
	{
		ID:              "same-tree",
		Number:          59,
		Title:           "Same Tree",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "BFS", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11},
		Description: `Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.`,
		Constraints: []string{
			"The number of nodes in both trees is in the range [0, 100]",
			"-10^4 <= Node.val <= 10^4",
		},
		Examples: []Example{
			{Input: "p = [1,2,3], q = [1,2,3]", Output: "true"},
			{Input: "p = [1,2], q = [1,null,2]", Output: "false"},
			{Input: "p = [1,2,1], q = [1,1,2]", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"p": []interface{}{1, 2, 3}, "q": []interface{}{1, 2, 3}}, Expected: true},
			{Input: map[string]interface{}{"p": []interface{}{1, 2}, "q": []interface{}{1, nil, 2}}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def isSameTree(p, q):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Recursively compare nodes: both null (same), one null (different), values differ (different), else check children."},
			{Level: 2, Type: "algorithm", Content: "Base cases: both None -> True, one None -> False, different values -> False. Recurse on left and right."},
			{Level: 3, Type: "code", Content: "if not p and not q: return True. if not p or not q: return False. return p.val == q.val and recurse."},
		},
		Solution: Solution{
			Code: `def isSameTree(p, q):
    if not p and not q:
        return True
    if not p or not q:
        return False
    if p.val != q.val:
        return False
    return isSameTree(p.left, q.left) and isSameTree(p.right, q.right)`,
			Explanation:     "Recursively compare: both null = same, one null = different, different values = different, else check both children.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(h) - recursion stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Both null", Explanation: "Empty trees are the same", CodeSnippet: "if not p and not q:\n    return True", LineStart: 2, LineEnd: 3},
				{Title: "One null", Explanation: "Structure differs", CodeSnippet: "if not p or not q:\n    return False", LineStart: 4, LineEnd: 5},
				{Title: "Check children", Explanation: "Both subtrees must match", CodeSnippet: "return isSameTree(p.left, q.left) and isSameTree(p.right, q.right)", LineStart: 8, LineEnd: 8},
			},
		},
	},
	{
		ID:              "subtree-of-another-tree",
		Number:          60,
		Title:           "Subtree of Another Tree",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "String Matching", "Binary Tree", "Hash Function"},
		RelatedChapters: []int{9, 10, 11},
		Description: `Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.`,
		Constraints: []string{
			"The number of nodes in the root tree is in the range [1, 2000]",
			"The number of nodes in the subRoot tree is in the range [1, 1000]",
			"-10^4 <= root.val <= 10^4",
			"-10^4 <= subRoot.val <= 10^4",
		},
		Examples: []Example{
			{Input: "root = [3,4,5,1,2], subRoot = [4,1,2]", Output: "true"},
			{Input: "root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 4, 5, 1, 2}, "subRoot": []interface{}{4, 1, 2}}, Expected: true},
			{Input: map[string]interface{}{"root": []interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0}, "subRoot": []interface{}{4, 1, 2}}, Expected: false},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def isSubtree(root, subRoot):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "For each node in root, check if subtree starting there equals subRoot. Use isSameTree helper."},
			{Level: 2, Type: "algorithm", Content: "Base: if root null, return false (unless subRoot also null). Check if trees are same, or recurse on children."},
			{Level: 3, Type: "code", Content: "def isSubtree(root, sub): return isSameTree(root, sub) or isSubtree(root.left, sub) or isSubtree(root.right, sub)."},
		},
		Solution: Solution{
			Code: `def isSubtree(root, subRoot):
    def isSameTree(p, q):
        if not p and not q:
            return True
        if not p or not q:
            return False
        return p.val == q.val and isSameTree(p.left, q.left) and isSameTree(p.right, q.right)

    if not root:
        return False
    if isSameTree(root, subRoot):
        return True
    return isSubtree(root.left, subRoot) or isSubtree(root.right, subRoot)`,
			Explanation:     "At each node, check if tree matches subRoot. If not, recursively check left and right children.",
			TimeComplexity:  "O(m * n) - for each of m nodes, potentially compare n nodes",
			SpaceComplexity: "O(h) - recursion depth",
			Walkthrough: []WalkthroughStep{
				{Title: "isSameTree helper", Explanation: "Checks if two trees are identical", CodeSnippet: "def isSameTree(p, q):", LineStart: 2, LineEnd: 7},
				{Title: "Check current node", Explanation: "See if subtree rooted here matches", CodeSnippet: "if isSameTree(root, subRoot):\n    return True", LineStart: 11, LineEnd: 12},
				{Title: "Recurse on children", Explanation: "Check if subRoot exists in left or right subtree", CodeSnippet: "return isSubtree(root.left, subRoot) or isSubtree(root.right, subRoot)", LineStart: 13, LineEnd: 13},
			},
		},
	},
	{
		ID:              "lowest-common-ancestor-bst",
		Number:          61,
		Title:           "Lowest Common Ancestor of a BST",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Search Tree", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11},
		Description: `Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself)."`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [2, 10^5]",
			"-10^9 <= Node.val <= 10^9",
			"All Node.val are unique",
			"p != q",
			"p and q will exist in the BST",
		},
		Examples: []Example{
			{Input: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8", Output: "6"},
			{Input: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4", Output: "2"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, "p": 2, "q": 8}, Expected: 6},
			{Input: map[string]interface{}{"root": []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, "p": 2, "q": 4}, Expected: 2},
		},
		TimeComplexity:  "O(h)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def lowestCommonAncestor(root, p, q):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use BST property: if both p and q are smaller, go left. If both larger, go right. Otherwise, current is LCA."},
			{Level: 2, Type: "algorithm", Content: "LCA is where p and q split (one goes left, one goes right, or one equals current node)."},
			{Level: 3, Type: "code", Content: "while root: if p.val < root.val and q.val < root.val: go left. elif both > root.val: go right. else: return root."},
		},
		Solution: Solution{
			Code: `def lowestCommonAncestor(root, p, q):
    while root:
        if p.val < root.val and q.val < root.val:
            root = root.left
        elif p.val > root.val and q.val > root.val:
            root = root.right
        else:
            return root`,
			Explanation:     "Use BST property. If both values smaller, go left. Both larger, go right. Otherwise, we're at the LCA (split point).",
			TimeComplexity:  "O(h) - height of tree",
			SpaceComplexity: "O(1) - iterative",
			Walkthrough: []WalkthroughStep{
				{Title: "Both smaller", Explanation: "p and q both in left subtree", CodeSnippet: "if p.val < root.val and q.val < root.val:\n    root = root.left", LineStart: 3, LineEnd: 4},
				{Title: "Both larger", Explanation: "p and q both in right subtree", CodeSnippet: "elif p.val > root.val and q.val > root.val:\n    root = root.right", LineStart: 5, LineEnd: 6},
				{Title: "Split point", Explanation: "p and q are on different sides (or one equals root)", CodeSnippet: "else:\n    return root", LineStart: 7, LineEnd: 8},
			},
		},
	},
	{
		ID:              "binary-tree-level-order-traversal",
		Number:          62,
		Title:           "Binary Tree Level Order Traversal",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "BFS", "Binary Tree"},
		RelatedChapters: []int{6, 9, 10, 11},
		Description: `Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [0, 2000]",
			"-1000 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "root = [3,9,20,null,null,15,7]", Output: "[[3],[9,20],[15,7]]"},
			{Input: "root = [1]", Output: "[[1]]"},
			{Input: "root = []", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 9, 20, nil, nil, 15, 7}}, Expected: [][]int{{3}, {9, 20}, {15, 7}}},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: [][]int{{1}}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def levelOrder(root):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use BFS with a queue. Process all nodes at current level before moving to next level.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Track level size. Process exactly that many nodes, adding their children. All processed nodes form one level."},
			{Level: 3, Type: "code", Content: "while queue: level_size = len(queue). for i in range(level_size): pop, add to level list, push children."},
		},
		Solution: Solution{
			Code: `from collections import deque

def levelOrder(root):
    if not root:
        return []

    result = []
    queue = deque([root])

    while queue:
        level = []
        level_size = len(queue)

        for _ in range(level_size):
            node = queue.popleft()
            level.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)

        result.append(level)

    return result`,
			Explanation:     "BFS with level tracking. Process exactly level_size nodes per iteration, adding children for next level.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n) - queue can hold entire level",
			Walkthrough: []WalkthroughStep{
				{Title: "Track level size", Explanation: "Know how many nodes belong to current level", CodeSnippet: "level_size = len(queue)", LineStart: 12, LineEnd: 12},
				{Title: "Process level", Explanation: "Pop exactly level_size nodes", CodeSnippet: "for _ in range(level_size):\n    node = queue.popleft()", LineStart: 14, LineEnd: 15},
				{Title: "Add children", Explanation: "Children form next level", CodeSnippet: "if node.left:\n    queue.append(node.left)", LineStart: 17, LineEnd: 20},
			},
		},
	},
	{
		ID:              "validate-binary-search-tree",
		Number:          63,
		Title:           "Validate Binary Search Tree",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Search Tree", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11},
		Description: `Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [1, 10^4]",
			"-2^31 <= Node.val <= 2^31 - 1",
		},
		Examples: []Example{
			{Input: "root = [2,1,3]", Output: "true"},
			{Input: "root = [5,1,4,null,null,3,6]", Output: "false", Explanation: "Root is 5 but right child is 4."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{2, 1, 3}}, Expected: true},
			{Input: map[string]interface{}{"root": []interface{}{5, 1, 4, nil, nil, 3, 6}}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def isValidBST(root):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Track valid range [min, max] for each node. Going left updates max, going right updates min."},
			{Level: 2, Type: "algorithm", Content: "Initially range is (-inf, +inf). Left child: (min, parent). Right child: (parent, max). Node must be in range."},
			{Level: 3, Type: "code", Content: "def valid(node, min_val, max_val): if not node: return True. check range, recurse with updated bounds."},
		},
		Solution: Solution{
			Code: `def isValidBST(root):
    def valid(node, min_val, max_val):
        if not node:
            return True
        if node.val <= min_val or node.val >= max_val:
            return False
        return valid(node.left, min_val, node.val) and valid(node.right, node.val, max_val)

    return valid(root, float('-inf'), float('inf'))`,
			Explanation:     "Track valid range for each node. Going left: upper bound becomes parent. Going right: lower bound becomes parent.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(h) - recursion depth",
			Walkthrough: []WalkthroughStep{
				{Title: "Base case", Explanation: "Empty node is valid", CodeSnippet: "if not node:\n    return True", LineStart: 3, LineEnd: 4},
				{Title: "Check range", Explanation: "Node value must be in (min, max)", CodeSnippet: "if node.val <= min_val or node.val >= max_val:\n    return False", LineStart: 5, LineEnd: 6},
				{Title: "Recurse with bounds", Explanation: "Left: update max. Right: update min", CodeSnippet: "return valid(node.left, min_val, node.val) and valid(node.right, node.val, max_val)", LineStart: 7, LineEnd: 7},
			},
		},
	},
	{
		ID:              "kth-smallest-element-in-bst",
		Number:          64,
		Title:           "Kth Smallest Element in a BST",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Search Tree", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11},
		Description: `Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.`,
		Constraints: []string{
			"The number of nodes in the tree is n",
			"1 <= k <= n <= 10^4",
			"0 <= Node.val <= 10^4",
		},
		Examples: []Example{
			{Input: "root = [3,1,4,null,2], k = 1", Output: "1"},
			{Input: "root = [5,3,6,2,4,null,null,1], k = 3", Output: "3"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 1, 4, nil, 2}, "k": 1}, Expected: 1},
			{Input: map[string]interface{}{"root": []interface{}{5, 3, 6, 2, 4, nil, nil, 1}, "k": 3}, Expected: 3},
		},
		TimeComplexity:  "O(H + k)",
		SpaceComplexity: "O(H)",
		StarterCode:     "def kthSmallest(root, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Inorder traversal of BST visits nodes in sorted order. Stop at kth node."},
			{Level: 2, Type: "algorithm", Content: "Use iterative inorder with stack. Push all left children, pop, count, push right."},
			{Level: 3, Type: "code", Content: "stack = []. Go left as far as possible. Pop, decrement k. If k == 0, return. Go right."},
		},
		Solution: Solution{
			Code: `def kthSmallest(root, k):
    stack = []
    curr = root

    while stack or curr:
        # Go left as far as possible
        while curr:
            stack.append(curr)
            curr = curr.left

        curr = stack.pop()
        k -= 1
        if k == 0:
            return curr.val

        curr = curr.right

    return -1`,
			Explanation:     "Iterative inorder traversal. BST inorder gives sorted values. Stop when we've visited k nodes.",
			TimeComplexity:  "O(H + k) - go down H levels, then visit k nodes",
			SpaceComplexity: "O(H) for stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Go left", Explanation: "Push all left children to stack", CodeSnippet: "while curr:\n    stack.append(curr)\n    curr = curr.left", LineStart: 7, LineEnd: 9},
				{Title: "Visit node", Explanation: "Pop and count", CodeSnippet: "curr = stack.pop()\nk -= 1", LineStart: 11, LineEnd: 12},
				{Title: "Check k", Explanation: "Return when kth node visited", CodeSnippet: "if k == 0:\n    return curr.val", LineStart: 13, LineEnd: 14},
			},
		},
	},
	{
		ID:              "construct-binary-tree-from-preorder-inorder",
		Number:          65,
		Title:           "Construct Binary Tree from Preorder and Inorder Traversal",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Array", "Hash Table", "Divide and Conquer", "Tree", "Binary Tree"},
		RelatedChapters: []int{4, 9, 10, 11},
		Description: `Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.`,
		Constraints: []string{
			"1 <= preorder.length <= 3000",
			"inorder.length == preorder.length",
			"-3000 <= preorder[i], inorder[i] <= 3000",
			"preorder and inorder consist of unique values",
			"Each value of inorder also appears in preorder",
			"preorder is guaranteed to be the preorder traversal of the tree",
			"inorder is guaranteed to be the inorder traversal of the tree",
		},
		Examples: []Example{
			{Input: "preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]", Output: "[3,9,20,null,null,15,7]"},
			{Input: "preorder = [-1], inorder = [-1]", Output: "[-1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"preorder": []int{3, 9, 20, 15, 7}, "inorder": []int{9, 3, 15, 20, 7}}, Expected: []interface{}{3, 9, 20, nil, nil, 15, 7}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def buildTree(preorder, inorder):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "First element of preorder is root. Find root in inorder - elements to left are left subtree, to right are right subtree.", ChapterRef: 4},
			{Level: 2, Type: "algorithm", Content: "Use hash map for O(1) lookup in inorder. Recursively build left subtree then right subtree."},
			{Level: 3, Type: "code", Content: "Build index map for inorder. Pop from preorder for root. Partition inorder by root index. Recurse on left, then right."},
		},
		Solution: Solution{
			Code: `def buildTree(preorder, inorder):
    inorder_map = {val: idx for idx, val in enumerate(inorder)}
    pre_idx = [0]  # Use list for mutable reference

    def build(left, right):
        if left > right:
            return None

        root_val = preorder[pre_idx[0]]
        pre_idx[0] += 1
        root = TreeNode(root_val)

        mid = inorder_map[root_val]
        root.left = build(left, mid - 1)
        root.right = build(mid + 1, right)

        return root

    return build(0, len(inorder) - 1)`,
			Explanation:     "Preorder gives roots. Inorder partitions into left/right subtrees. Hash map speeds up finding root position.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n) for hash map and recursion",
			Walkthrough: []WalkthroughStep{
				{Title: "Build index map", Explanation: "O(1) lookup for root position in inorder", CodeSnippet: "inorder_map = {val: idx for idx, val in enumerate(inorder)}", LineStart: 2, LineEnd: 2},
				{Title: "Get root from preorder", Explanation: "First unprocessed preorder element is current root", CodeSnippet: "root_val = preorder[pre_idx[0]]\npre_idx[0] += 1", LineStart: 9, LineEnd: 10},
				{Title: "Partition and recurse", Explanation: "Build left subtree before right (preorder order)", CodeSnippet: "root.left = build(left, mid - 1)\nroot.right = build(mid + 1, right)", LineStart: 14, LineEnd: 15},
			},
		},
	},
	{
		ID:              "binary-tree-max-path-sum",
		Number:          66,
		Title:           "Binary Tree Maximum Path Sum",
		Difficulty:      "Hard",
		Category:        "trees",
		Tags:            []string{"Dynamic Programming", "Tree", "DFS", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11, 12},
		Description: `A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path.`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [1, 3 * 10^4]",
			"-1000 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "root = [1,2,3]", Output: "6", Explanation: "Optimal path is 2 -> 1 -> 3 with sum 6."},
			{Input: "root = [-10,9,20,null,null,15,7]", Output: "42", Explanation: "Optimal path is 15 -> 20 -> 7 with sum 42."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3}}, Expected: 6},
			{Input: map[string]interface{}{"root": []interface{}{-10, 9, 20, nil, nil, 15, 7}}, Expected: 42},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "def maxPathSum(root):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "At each node, consider: path through node (left + node + right) vs. path continuing upward (node + max(left, right)).", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "DFS returns max gain if path continues upward. Track global max for paths that end at current node (including both children)."},
			{Level: 3, Type: "code", Content: "def dfs(node): returns max(0, node.val + max(left, right)). Update global max with node.val + left + right."},
		},
		Solution: Solution{
			Code: `def maxPathSum(root):
    max_sum = [float('-inf')]

    def dfs(node):
        if not node:
            return 0

        # Max gain from left/right children (ignore if negative)
        left_gain = max(0, dfs(node.left))
        right_gain = max(0, dfs(node.right))

        # Path through current node (using both children)
        path_sum = node.val + left_gain + right_gain
        max_sum[0] = max(max_sum[0], path_sum)

        # Return max gain for path continuing upward
        return node.val + max(left_gain, right_gain)

    dfs(root)
    return max_sum[0]`,
			Explanation:     "DFS computes max path continuing upward. At each node, also check if path through node (both children) is best. Track global max.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(h) - recursion depth",
			Walkthrough: []WalkthroughStep{
				{Title: "Ignore negative paths", Explanation: "Only use child paths if they add positive value", CodeSnippet: "left_gain = max(0, dfs(node.left))", LineStart: 9, LineEnd: 10},
				{Title: "Check path through node", Explanation: "Path using both children ends here", CodeSnippet: "path_sum = node.val + left_gain + right_gain", LineStart: 13, LineEnd: 14},
				{Title: "Return upward gain", Explanation: "Can only use one child if path continues up", CodeSnippet: "return node.val + max(left_gain, right_gain)", LineStart: 17, LineEnd: 17},
			},
		},
	},
	{
		ID:              "serialize-and-deserialize-binary-tree",
		Number:          67,
		Title:           "Serialize and Deserialize Binary Tree",
		Difficulty:      "Hard",
		Category:        "trees",
		Tags:            []string{"String", "Tree", "DFS", "BFS", "Design", "Binary Tree"},
		RelatedChapters: []int{9, 10, 11, 12},
		Description: `Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later.

Design an algorithm to serialize and deserialize a binary tree.`,
		Constraints: []string{
			"The number of nodes in the tree is in the range [0, 10^4]",
			"-1000 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "root = [1,2,3,null,null,4,5]", Output: "[1,2,3,null,null,4,5]"},
			{Input: "root = []", Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, nil, nil, 4, 5}}, Expected: []interface{}{1, 2, 3, nil, nil, 4, 5}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `class Codec:
    def serialize(self, root):
        # Write your solution here
        pass

    def deserialize(self, data):
        # Write your solution here
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use preorder traversal with null markers. Serialize: node value or 'N' for null. Deserialize: recursively build from string.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "Preorder with nulls: '1,2,N,N,3,4,N,N,5,N,N'. Deserialize by reading values in order and recursively building."},
			{Level: 3, Type: "code", Content: "serialize: dfs, append val or 'N'. deserialize: use iterator, pop value, if 'N' return None, else create node and recurse."},
		},
		Solution: Solution{
			Code: `class Codec:
    def serialize(self, root):
        result = []

        def dfs(node):
            if not node:
                result.append('N')
                return
            result.append(str(node.val))
            dfs(node.left)
            dfs(node.right)

        dfs(root)
        return ','.join(result)

    def deserialize(self, data):
        values = iter(data.split(','))

        def dfs():
            val = next(values)
            if val == 'N':
                return None
            node = TreeNode(int(val))
            node.left = dfs()
            node.right = dfs()
            return node

        return dfs()`,
			Explanation:     "Preorder traversal with null markers. Serialize to comma-separated string. Deserialize by consuming values in preorder.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Mark nulls", Explanation: "Use 'N' to mark empty nodes", CodeSnippet: "if not node:\n    result.append('N')", LineStart: 6, LineEnd: 8},
				{Title: "Preorder serialize", Explanation: "Root, then left, then right", CodeSnippet: "result.append(str(node.val))\ndfs(node.left)\ndfs(node.right)", LineStart: 9, LineEnd: 11},
				{Title: "Iterator deserialize", Explanation: "Consume values in same order", CodeSnippet: "val = next(values)\nif val == 'N': return None", LineStart: 20, LineEnd: 22},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, TreesProblems...)
}
