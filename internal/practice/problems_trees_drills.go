package practice

// TreesDrills are the trees foundations: build a tree by hand, then run each
// core traversal and the two BST updates on their own, before the trees
// ladder applies them. They follow docs/pedagogy.md like any problem, minus
// the brute force (a drill is the plain algorithm itself). The by-hand trace
// drills that go with them run in the browser (static/trace-drills.js).
var TreesDrills = []*Problem{
	{
		ID:              "drill-build-binary-tree",
		Kind:            KindDrill,
		Order:           1,
		Title:           "Build a Binary Tree by Hand",
		Slug:            "drill-build-binary-tree",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "Binary Tree", "BFS", "Foundations"},
		RelatedChapters: []int{9},
		Description: `Every tree problem hands you a ready-made ` + "`root`" + `. This drill makes you build one yourself.

1. Finish the ` + "`TreeNode`" + ` class: a node stores its value in ` + "`val`" + ` and its children in ` + "`left`" + ` and ` + "`right`" + ` (None when a child is missing).
2. Write ` + "`fromLevelOrder(values)`" + `, which builds a tree from the **level-order** list format that LeetCode uses, and returns its root.

How the format works: ` + "`values[0]`" + ` is the root. Then, taking the real nodes one at a time in the order they were created, the next two entries are that node's left child and right child. ` + "`None`" + ` means "no child here", and a None slot gets no children of its own, so it uses up no further entries. Trailing Nones may be left off, so the list can end before every node has had both slots filled.

Return None when ` + "`values`" + ` is empty or ` + "`values[0]`" + ` is None. The judge reads your tree back in level order, so return real ` + "`TreeNode`" + ` objects, not a list.`,
		Constraints: []string{
			"0 <= len(values) <= 1000",
			"-1000 <= values[i] <= 1000, or values[i] is None",
			"values[0] is None only when the tree is empty",
		},
		Examples: []Example{
			{Input: "values = [1,2,3,4,5,6,7]", Output: "[1,2,3,4,5,6,7]", Explanation: "1 is the root. Its children are the next two entries, 2 and 3. Then 2 takes 4 and 5, and 3 takes 6 and 7: a full tree of three levels."},
			{Input: "values = [1,null,2,3]", Output: "[1,null,2,3]", Explanation: "1 has no left child and 2 as its right child. The None slot has no children, so the next real node, 2, takes 3 as its left child."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"values": []interface{}{1, 2, 3, 4, 5, 6, 7}}, Expected: []interface{}{1, 2, 3, 4, 5, 6, 7}, Note: "example 1: a full tree of three levels"},
			{Input: map[string]interface{}{"values": []interface{}{1, nil, 2, 3}}, Expected: []interface{}{1, nil, 2, 3}, Note: "example 2: a None slot takes no children (the 2*i+1 index formula gets this wrong)"},
			{Input: map[string]interface{}{"values": []interface{}{}}, Expected: []interface{}{}, Hidden: true, Note: "empty list builds the empty tree (None)"},
			{Input: map[string]interface{}{"values": []interface{}{0}}, Expected: []interface{}{0}, Hidden: true, Note: "a root of value 0 is a real node (`if not values[0]` treats it as missing)"},
			{Input: map[string]interface{}{"values": []interface{}{5, 3, 8, nil, 4, 0, 9}}, Expected: []interface{}{5, 3, 8, nil, 4, 0, 9}, Hidden: true, Note: "a child of value 0 must still become a node (`if values[i]:` drops it)"},
			{Input: map[string]interface{}{"values": []interface{}{1, 2, nil, 3, nil, 4}}, Expected: []interface{}{1, 2, nil, 3, nil, 4}, Hidden: true, Note: "left chain: each node's right slot is None"},
			{Input: map[string]interface{}{"values": []interface{}{-1, nil, -2, nil, -3}}, Expected: []interface{}{-1, nil, -2, nil, -3}, Hidden: true, Note: "right chain of negative values"},
			{Input: map[string]interface{}{"values": []interface{}{1, 2}}, Expected: []interface{}{1, 2}, Hidden: true, Note: "the list ends after a left child, so the right slot is missing (an unchecked index raises IndexError)"},
			{Input: map[string]interface{}{"values": []interface{}{1, 2, 3, nil, nil, nil, nil}}, Expected: []interface{}{1, 2, 3}, Hidden: true, Note: "trailing Nones add no nodes"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode: `class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        # Store the value and the two child links
        pass


def fromLevelOrder(values):
    # Return the root TreeNode, or None for an empty tree
    pass`,
		StarterCodeOxigen: "// TreeNode(val, left, right) is provided; TreeNode(5) leaves both children as None.\n\nfun fromLevelOrder(values) {\n    // Write your solution here\n}\n",
		Harness: Harness{
			ReturnType: "TreeNode",
			Validator: `def check(raw, args, result, expected):
    if isinstance(result_obj, (list, tuple, dict, str, int, float)):
        return (False, "return the root TreeNode, not a " + type(result_obj).__name__)
    if result_obj is not None:
        result_obj.val, result_obj.left, result_obj.right
    if result != expected:
        return (False, "your tree reads back in level order as " + str(result))
    return True`,
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The entries come in pairs, one pair per real node, in the order the nodes were created. What structure hands you nodes back in exactly the order you put them in?"},
			{Level: 2, Type: "insight", Content: "A queue. Every node you create joins the back of the queue; the node at the front is the next parent, and it takes the next two entries as its left and right children. This is the level-by-level BFS from Chapter 9, run while building instead of while reading. Test entries with `is not None`, because 0 is a real value.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Define the node: store val, left and right in `__init__`.\n2. Handle the empty tree: return None for an empty list or a None root.\n3. Seed the queue with the root, and point an index at entry 1.\n4. Give the next parent its children: pop the front node; if the entry at the index is a value, make it the left child and queue it; move the index on; do the same for the right child, checking the index is still inside the list.\n5. Return the root once the entries or the parents run out."},
			{Level: 4, Type: "skeleton", Content: "from collections import deque\n\n\nclass TreeNode:\n    def __init__(self, val=0, left=None, right=None):\n        self.val = val\n        self.left = ____\n        self.right = ____\n\n\ndef fromLevelOrder(values):\n    if not values or values[0] is None:\n        return None\n    root = TreeNode(values[0])\n    queue = deque([root])\n    i = 1\n    while queue and i < len(values):\n        node = ____\n        if values[i] is not None:\n            node.left = TreeNode(values[i])\n            ____\n        i += 1\n        if ____:\n            node.right = TreeNode(values[i])\n            queue.append(node.right)\n        i += 1\n    return root"},
		},
		Solution: Solution{
			Code: `from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def fromLevelOrder(values):
    if not values or values[0] is None:
        return None
    root = TreeNode(values[0])
    queue = deque([root])
    i = 1
    while queue and i < len(values):
        node = queue.popleft()
        if values[i] is not None:
            node.left = TreeNode(values[i])
            queue.append(node.left)
        i += 1
        if i < len(values) and values[i] is not None:
            node.right = TreeNode(values[i])
            queue.append(node.right)
        i += 1
    return root`,
			Explanation: "A tree is just nodes holding links to other nodes, so the class stores a value and two links. Building from level order is the Chapter 9 BFS run in reverse: a queue holds the nodes that still need children, and each one in turn takes the next two entries of the list.",
			KeyInsight:  "The list gives children to real nodes in the order the nodes were created, and a queue returns nodes in the order they went in. So the front of the queue is always the node whose children come next, and a None slot never joins the queue, which is why it uses no entries.",
			WhyItWorks:  "Invariant: when the index reaches entry i, the queue holds, in creation order, exactly the real nodes that have not yet been given their two child slots, and its front is the node that entries i and i+1 belong to. It holds at the start (only the root, entries 1 and 2). Each round pops that node, fills its two slots from entries i and i+1, and queues only the children that are real, so the next parent in creation order moves to the front. Each entry is read once and each node is queued once: O(n) time, and the queue holds at most one level, O(n) in the worst case.",
			CommonMistakes: []string{
				"Testing `if values[i]:` instead of `if values[i] is not None:`, which drops every node whose value is 0. Compare with None explicitly.",
				"Using the array formula (children of index i at 2*i+1 and 2*i+2). That only works when no None slot is followed by more nodes; LeetCode's format gives children only to real nodes. Use the queue.",
				"Reading values[i] for the right child without checking `i < len(values)`: the list can end right after a left child. Check the index before each read.",
				"Returning `values`, or a list of nodes, instead of the root node. The function hands back one TreeNode, and the rest of the tree hangs off it.",
			},
			TimeComplexity:  "O(n): each entry is read once and each node is created and queued once",
			SpaceComplexity: "O(n) for the nodes themselves; the queue holds at most one level",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the node", Explanation: "A node is three fields: its value and the two links. The defaults make TreeNode(5) a leaf.", CodeSnippet: "self.val = val\nself.left = left\nself.right = right", LineStart: 4, LineEnd: 8},
				{Title: "Handle the empty tree", Explanation: "No entries, or a None root, means there is no tree to build, so the answer is None.", CodeSnippet: "if not values or values[0] is None:\n    return None", LineStart: 12, LineEnd: 13},
				{Title: "Seed the queue with the root", Explanation: "The root is waiting for its children, and they are entries 1 and 2, so i starts at 1.", CodeSnippet: "root = TreeNode(values[0])\nqueue = deque([root])\ni = 1", LineStart: 14, LineEnd: 16},
				{Title: "Give the next parent its children", Explanation: "On [1,null,2,3]: node 1 reads None (no left child) and 2 (right child, queued). Then node 2 reads 3 as its left child, and the list ends. Only real children join the queue, so a None slot never claims entries.", CodeSnippet: "node = queue.popleft()\nif values[i] is not None:\n    node.left = TreeNode(values[i])\n    queue.append(node.left)\nif i < len(values) and values[i] is not None:", LineStart: 17, LineEnd: 26},
				{Title: "Return the root", Explanation: "Every node is reachable from the root, so the root is the whole tree.", CodeSnippet: "return root", LineStart: 27, LineEnd: 27},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does a None entry use up no further entries of the list?", Answer: "Entries are handed out to real nodes only: each node popped from the queue takes the next two. A None slot never becomes a node, so it is never queued and never asks for children. That is what lets the format skip whole missing subtrees instead of writing Nones for every empty position."},
			{Prompt: "What would you write to turn a tree back into this list?", Answer: "The same BFS in the other direction: start a queue with the root, and for each node popped, append its value, or None for a missing node, and queue both of its children, None included, without queueing children of None. At the end, drop the trailing Nones. The judge does exactly this to read your tree back."},
		},
	},
	{
		ID:              "drill-preorder-traversal",
		Number:          144,
		Kind:            KindDrill,
		Order:           2,
		Title:           "Preorder Traversal",
		Slug:            "binary-tree-preorder-traversal",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Tree", "Recursion", "Foundations"},
		RelatedChapters: []int{9, 8},
		Description: `You are given the root of a binary tree. Return its node values in **preorder**: for each node, record the node itself first, then everything in its left subtree (in preorder), then everything in its right subtree (in preorder).

An empty tree (root is None) gives an empty list.`,
		Constraints: []string{
			"The tree has between 0 and 100 nodes",
			"-100 <= Node.val <= 100",
		},
		Examples: []Example{
			{Input: "root = [1,null,2,3]", Output: "[1,2,3]", Explanation: "Record 1 first. Its left subtree is empty, so move to its right subtree: record 2, then 2's left child 3."},
			{Input: "root = [1,2,3,4,5,null,8,null,null,6,7,9]", Output: "[1,2,4,5,6,7,3,8,9]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2, 3}}, Expected: []interface{}{1, 2, 3}, Note: "example 1: an empty left subtree"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}}, Expected: []interface{}{1, 2, 4, 5, 6, 7, 3, 8, 9}, Note: "example 2: the whole left subtree comes before the right one"},
			{Input: map[string]interface{}{"root": []interface{}{}}, Expected: []interface{}{}, Hidden: true, Note: "empty tree gives []"},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: []interface{}{1}, Hidden: true, Note: "a single node"},
			{Input: map[string]interface{}{"root": []interface{}{3, 1, 2}}, Expected: []interface{}{3, 1, 2}, Hidden: true, Note: "three nodes give a different answer for each order (inorder would be 1,3,2)"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, 5, 6, 7}}, Expected: []interface{}{1, 2, 4, 5, 3, 6, 7}, Hidden: true, Note: "full tree: not level order (level order would be 1..7)"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, nil, 3, nil, 4}}, Expected: []interface{}{1, 2, 3, 4}, Hidden: true, Note: "left chain: the node comes before its descendants"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef preorderTraversal(root):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose a helper already gives you the preorder list of any smaller tree. Where does the root's own value go relative to the lists for its two subtrees?"},
			{Level: 2, Type: "insight", Content: "Preorder is the DFS template with the work done before the two recursive calls: record the node, then visit left, then visit right. The name tells you where the visit goes: pre = before the children. Collect values in one list shared by every call so nothing is copied. See Chapter 9's preorder, inorder and postorder.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Make the answer list.\n2. Write a helper visit(node): stop at an empty tree.\n3. Record the node, then visit its left subtree, then its right subtree.\n4. Call visit on the root and return the list."},
			{Level: 4, Type: "skeleton", Content: "def preorderTraversal(root):\n    out = []\n\n    def visit(node):\n        if node is None:\n            return\n        ____\n        visit(node.left)\n        ____\n\n    visit(root)\n    return out"},
		},
		Solution: Solution{
			Code: `def preorderTraversal(root):
    out = []

    def visit(node):
        if node is None:
            return
        out.append(node.val)
        visit(node.left)
        visit(node.right)

    visit(root)
    return out`,
			Explanation: "The first of the three depth-first orders. A helper walks the tree with the Chapter 9 DFS template and appends each value just before descending into its children, so every node appears before everything below it.",
			KeyInsight:  "The three depth-first orders run the exact same walk; they differ only in where the line that records the node sits. Putting it before both recursive calls gives preorder, which is why the root always comes first and each subtree's root comes before its descendants.",
			WhyItWorks:  "By induction on tree size. An empty tree records nothing, which is its preorder. For a node, the helper records the node, then by the induction hypothesis the call on the left subtree appends exactly the left subtree's preorder, and the call on the right appends the right subtree's preorder. That concatenation is the definition of preorder. Each node is visited once, O(n), and the call stack holds one frame per level of the current path, O(h).",
			CommonMistakes: []string{
				"Writing `return [root.val] + preorderTraversal(root.left) + preorderTraversal(root.right)`: correct, but it copies lists at every level, O(n * h) in the worst case. Append to one shared list.",
				"Recording the node after visiting the left subtree, which gives inorder. The record line goes before both calls.",
				"Forgetting the None check, so the helper reads `.val` of a missing child and crashes on every leaf.",
			},
			TimeComplexity:  "O(n): each node is visited once",
			SpaceComplexity: "O(h) for the call stack, besides the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Make the answer list", Explanation: "One list that every call appends to, so no list is ever copied.", CodeSnippet: "out = []", LineStart: 2, LineEnd: 2},
				{Title: "Stop at an empty tree", Explanation: "A missing child has no values to record, and stopping here is what ends the recursion below each leaf.", CodeSnippet: "if node is None:\n    return", LineStart: 5, LineEnd: 6},
				{Title: "Record, then go left, then go right", Explanation: "On [3,1,2]: record 3, then the left call records 1, then the right call records 2, giving [3,1,2].", CodeSnippet: "out.append(node.val)\nvisit(node.left)\nvisit(node.right)", LineStart: 7, LineEnd: 9},
				{Title: "Start at the root", Explanation: "One call on the root walks the whole tree; then the list is complete.", CodeSnippet: "visit(root)\nreturn out", LineStart: 11, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Iterative with an explicit stack",
					Code: `def preorderTraversal(root):
    out = []
    stack = [root] if root else []
    while stack:
        node = stack.pop()
        out.append(node.val)
        if node.right:
            stack.append(node.right)
        if node.left:
            stack.append(node.left)
    return out`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(h)",
					Explanation:     "The stack stands in for the call stack. Push the right child before the left one, so the left child is popped, and recorded, first.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "In the iterative version, why is the right child pushed before the left?", Answer: "A stack gives back the last thing pushed. Preorder must finish the left subtree before starting the right one, so the left child has to come off first, which means it has to go on last."},
			{Prompt: "Given only a preorder list, can you tell which value is the root? Can you rebuild the tree?", Answer: "The first value is always the root. But the list alone does not say where the left subtree ends and the right begins: [1,2] could be 1 with a left child 2 or with a right child 2. You need a second order (inorder) to split the rest, which is exactly how Construct Binary Tree from Preorder and Inorder Traversal works later in the ladder."},
		},
	},
	{
		ID:              "drill-inorder-traversal",
		Number:          94,
		Kind:            KindDrill,
		Order:           3,
		Title:           "Inorder Traversal",
		Slug:            "binary-tree-inorder-traversal",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Tree", "Recursion", "Foundations"},
		RelatedChapters: []int{9, 8},
		Description: `You are given the root of a binary tree. Return its node values in **inorder**: for each node, first everything in its left subtree (in inorder), then the node itself, then everything in its right subtree (in inorder).

An empty tree (root is None) gives an empty list. For a binary search tree, inorder lists the values in sorted order.`,
		Constraints: []string{
			"The tree has between 0 and 100 nodes",
			"-100 <= Node.val <= 100",
		},
		Examples: []Example{
			{Input: "root = [1,null,2,3]", Output: "[1,3,2]", Explanation: "1 has no left subtree, so record 1. Then do its right subtree in inorder: 2's left subtree (3) comes first, then 2."},
			{Input: "root = [1,2,3,4,5,null,8,null,null,6,7,9]", Output: "[4,2,6,5,7,1,3,9,8]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2, 3}}, Expected: []interface{}{1, 3, 2}, Note: "example 1: a left child is recorded before its parent"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}}, Expected: []interface{}{4, 2, 6, 5, 7, 1, 3, 9, 8}, Note: "example 2: the root sits between its two subtrees"},
			{Input: map[string]interface{}{"root": []interface{}{}}, Expected: []interface{}{}, Hidden: true, Note: "empty tree gives []"},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: []interface{}{1}, Hidden: true, Note: "a single node"},
			{Input: map[string]interface{}{"root": []interface{}{3, 1, 2}}, Expected: []interface{}{1, 3, 2}, Hidden: true, Note: "three nodes give a different answer for each order (preorder would be 3,1,2)"},
			{Input: map[string]interface{}{"root": []interface{}{0, -1, 1}}, Expected: []interface{}{-1, 0, 1}, Hidden: true, Note: "a small BST comes out sorted"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, nil, 3, nil, 4}}, Expected: []interface{}{4, 3, 2, 1}, Hidden: true, Note: "left chain: the deepest node comes first"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef inorderTraversal(root):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You wrote preorder by recording a node before visiting its children. Inorder lists the left subtree before the node. Which single line of your preorder moves, and where to?"},
			{Level: 2, Type: "insight", Content: "Inorder is the same DFS walk with the record line between the two recursive calls: visit left, record the node, visit right. In a BST everything on the left is smaller and everything on the right is larger, so inorder gives sorted order. See Chapter 9.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Make the answer list.\n2. Write a helper visit(node): stop at an empty tree.\n3. Visit the left subtree, record the node, then visit the right subtree.\n4. Call visit on the root and return the list."},
			{Level: 4, Type: "skeleton", Content: "def inorderTraversal(root):\n    out = []\n\n    def visit(node):\n        if node is None:\n            return\n        ____\n        ____\n        visit(node.right)\n\n    visit(root)\n    return out"},
		},
		Solution: Solution{
			Code: `def inorderTraversal(root):
    out = []

    def visit(node):
        if node is None:
            return
        visit(node.left)
        out.append(node.val)
        visit(node.right)

    visit(root)
    return out`,
			Explanation: "The preorder drill with one line moved. The helper finishes a node's entire left subtree, records the node, and then does its right subtree, so each node sits between the values below it on the left and on the right.",
			KeyInsight:  "Moving the record line between the two recursive calls makes each node wait until its left subtree is done. That puts every node after all of its left descendants and before all of its right ones, which for a BST is exactly ascending order.",
			WhyItWorks:  "By induction on tree size. An empty tree records nothing. For a node, the left call appends the left subtree's inorder (induction hypothesis), then the node is appended, then the right call appends the right subtree's inorder; that concatenation is the definition of inorder. In a BST every left value is smaller than the node and every right value larger, so by the same induction the list is sorted. Each node is visited once: O(n) time, O(h) stack.",
			CommonMistakes: []string{
				"Recording the node before the left call (preorder) or after the right call (postorder). The record line sits between the two calls.",
				"Building and concatenating new lists at every call, which copies values over and over. Append to one shared list.",
				"Expecting sorted output from any binary tree. Only a BST's inorder is sorted; a general tree's inorder follows its shape.",
			},
			TimeComplexity:  "O(n): each node is visited once",
			SpaceComplexity: "O(h) for the call stack, besides the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Make the answer list", Explanation: "One shared list, appended to by every call.", CodeSnippet: "out = []", LineStart: 2, LineEnd: 2},
				{Title: "Stop at an empty tree", Explanation: "A missing child contributes nothing and ends the recursion.", CodeSnippet: "if node is None:\n    return", LineStart: 5, LineEnd: 6},
				{Title: "Go left, record, then go right", Explanation: "On [3,1,2]: the left call records 1, then 3 is recorded, then the right call records 2, giving [1,3,2].", CodeSnippet: "visit(node.left)\nout.append(node.val)\nvisit(node.right)", LineStart: 7, LineEnd: 9},
				{Title: "Start at the root", Explanation: "One call on the root walks the whole tree.", CodeSnippet: "visit(root)\nreturn out", LineStart: 11, LineEnd: 12},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does inorder on a BST give sorted values?", Answer: "In a BST every value in a node's left subtree is smaller and every value in its right subtree is larger. Inorder records the whole left subtree, then the node, then the whole right subtree, so each node lands after everything smaller below it and before everything larger. Applied at every node, that makes the full list ascending."},
			{Prompt: "In the left chain 1 -> 2 -> 3 -> 4 (each a left child), why does 4 come first?", Answer: "Each node must finish its left subtree before recording itself, so 1 waits for 2, which waits for 3, which waits for 4. Node 4 has no left child, so it is recorded first, and then each waiting node is recorded as its call resumes: 4, 3, 2, 1. Those waiting calls are the stack the iterative version manages by hand."},
		},
	},
	{
		ID:              "drill-postorder-traversal",
		Number:          145,
		Kind:            KindDrill,
		Order:           4,
		Title:           "Postorder Traversal",
		Slug:            "binary-tree-postorder-traversal",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Tree", "Recursion", "Foundations"},
		RelatedChapters: []int{9, 8},
		Description: `You are given the root of a binary tree. Return its node values in **postorder**: for each node, first everything in its left subtree (in postorder), then everything in its right subtree (in postorder), and the node itself last.

An empty tree (root is None) gives an empty list.`,
		Constraints: []string{
			"The tree has between 0 and 100 nodes",
			"-100 <= Node.val <= 100",
		},
		Examples: []Example{
			{Input: "root = [1,null,2,3]", Output: "[3,2,1]", Explanation: "1 must wait for its right subtree. Inside it, 2 must wait for its left child 3. So 3 is recorded first, then 2, then 1."},
			{Input: "root = [1,2,3,4,5,null,8,null,null,6,7,9]", Output: "[4,6,7,5,2,9,8,3,1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2, 3}}, Expected: []interface{}{3, 2, 1}, Note: "example 1: every node waits for its children"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}}, Expected: []interface{}{4, 6, 7, 5, 2, 9, 8, 3, 1}, Note: "example 2: the root is last"},
			{Input: map[string]interface{}{"root": []interface{}{}}, Expected: []interface{}{}, Hidden: true, Note: "empty tree gives []"},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: []interface{}{1}, Hidden: true, Note: "a single node"},
			{Input: map[string]interface{}{"root": []interface{}{3, 1, 2}}, Expected: []interface{}{1, 2, 3}, Hidden: true, Note: "three nodes give a different answer for each order (preorder would be 3,1,2)"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, 5, 6, 7}}, Expected: []interface{}{4, 5, 2, 6, 7, 3, 1}, Hidden: true, Note: "full tree: not reversed preorder (that would be 7,6,3,5,4,2,1)"},
			{Input: map[string]interface{}{"root": []interface{}{0, -1, 1}}, Expected: []interface{}{-1, 1, 0}, Hidden: true, Note: "values of 0 and below"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef postorderTraversal(root):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In postorder a node is recorded only once everything below it has been recorded. Where must the record line go relative to the two recursive calls?"},
			{Level: 2, Type: "insight", Content: "Postorder puts the record line after both recursive calls: visit left, visit right, record the node. It is the order the DFS template from Chapter 9 uses when a node needs its children's answers before it can compute its own (depth, height, sums).", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Make the answer list.\n2. Write a helper visit(node): stop at an empty tree.\n3. Visit the left subtree, then the right subtree, then record the node.\n4. Call visit on the root and return the list."},
			{Level: 4, Type: "skeleton", Content: "def postorderTraversal(root):\n    out = []\n\n    def visit(node):\n        if node is None:\n            return\n        ____\n        ____\n        ____\n\n    visit(root)\n    return out"},
		},
		Solution: Solution{
			Code: `def postorderTraversal(root):
    out = []

    def visit(node):
        if node is None:
            return
        visit(node.left)
        visit(node.right)
        out.append(node.val)

    visit(root)
    return out`,
			Explanation: "The last of the three depth-first orders: the same walk as preorder and inorder, with the record line after both recursive calls. A node is recorded only when its whole subtree is finished, so the root comes last.",
			KeyInsight:  "Recording after both calls means a node is written down on the way back up, once both children are done. That is the moment a parent has its children's results, which is why most tree problems that return a value per subtree (Maximum Depth, Diameter, Balanced Tree) are postorder at heart.",
			WhyItWorks:  "By induction on tree size. An empty tree records nothing. For a node, the left call appends the left subtree's postorder, the right call appends the right subtree's postorder (induction hypothesis), and then the node is appended; that is the definition of postorder. Every node is appended after all of its descendants. Each node is visited once: O(n) time, O(h) stack.",
			CommonMistakes: []string{
				"Reversing the preorder list and calling it postorder. Reversed preorder is right, left, node backwards; it matches postorder only when you also swap which child goes first (the iterative alternative below does exactly that).",
				"Putting the record line between the calls (inorder) or before them (preorder). It goes after both.",
				"Using the order of the record line to mean the order of the calls. All three orders call left before right; only the record moves.",
			},
			TimeComplexity:  "O(n): each node is visited once",
			SpaceComplexity: "O(h) for the call stack, besides the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Make the answer list", Explanation: "One shared list, appended to by every call.", CodeSnippet: "out = []", LineStart: 2, LineEnd: 2},
				{Title: "Stop at an empty tree", Explanation: "A missing child contributes nothing and ends the recursion.", CodeSnippet: "if node is None:\n    return", LineStart: 5, LineEnd: 6},
				{Title: "Go left, go right, then record", Explanation: "On [3,1,2]: the left call records 1, the right call records 2, and only then is 3 recorded, giving [1,2,3].", CodeSnippet: "visit(node.left)\nvisit(node.right)\nout.append(node.val)", LineStart: 7, LineEnd: 9},
				{Title: "Start at the root", Explanation: "One call on the root walks the whole tree, and the root is the last value appended.", CodeSnippet: "visit(root)\nreturn out", LineStart: 11, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Iterative: reversed node-right-left",
					Code: `def postorderTraversal(root):
    out = []
    stack = [root] if root else []
    while stack:
        node = stack.pop()
        out.append(node.val)
        if node.left:
            stack.append(node.left)
        if node.right:
            stack.append(node.right)
    return out[::-1]`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(n) for the reversed list",
					Explanation:     "An iterative preorder that goes right before left records node, right, left. Reversed, that is left, right, node: postorder.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why do problems like Maximum Depth use postorder?", Answer: "A node's depth is 1 + the larger of its children's depths, so it can only be computed after both children have returned their answers. Doing the node's work after both recursive calls is exactly postorder."},
			{Prompt: "To delete every node of a tree in a language without garbage collection, which order is safe, and why?", Answer: "Postorder. Freeing a node before its children would lose the only links to them. Postorder frees both subtrees first, so every node is freed after the nodes it points to."},
		},
	},
	{
		ID:              "drill-level-order-values",
		Kind:            KindDrill,
		Order:           5,
		Title:           "Level-Order (BFS) Traversal",
		Slug:            "drill-level-order-values",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "BFS", "Binary Tree", "Queue", "Foundations"},
		RelatedChapters: []int{9},
		Description: `You are given the root of a binary tree. Return its node values in **level order** as one flat list: the root, then every node one level down from left to right, then every node two levels down from left to right, and so on.

An empty tree (root is None) gives an empty list. (Binary Tree Level Order Traversal, later in the ladder, asks for one list per level; this drill is the plain walk underneath it.)`,
		Constraints: []string{
			"The tree has between 0 and 2000 nodes",
			"-1000 <= Node.val <= 1000",
		},
		Examples: []Example{
			{Input: "root = [3,9,20,null,null,15,7]", Output: "[3,9,20,15,7]", Explanation: "Level 0 is 3. Level 1 is 9 then 20. Level 2 is 15 then 7, the children of 20."},
			{Input: "root = [1,2,3,4,null,null,5]", Output: "[1,2,3,4,5]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{3, 9, 20, nil, nil, 15, 7}}, Expected: []interface{}{3, 9, 20, 15, 7}, Note: "example 1: a missing subtree on the left"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, nil, nil, 5}}, Expected: []interface{}{1, 2, 3, 4, 5}, Note: "example 2: a level's nodes come from different parents"},
			{Input: map[string]interface{}{"root": []interface{}{}}, Expected: []interface{}{}, Hidden: true, Note: "empty tree gives []"},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: []interface{}{1}, Hidden: true, Note: "a single node"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, 5, 6, 7}}, Expected: []interface{}{1, 2, 3, 4, 5, 6, 7}, Hidden: true, Note: "full tree (a stack instead of a queue gives 1,3,7,6,2,5,4 or the preorder 1,2,4,5,3,6,7)"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}}, Expected: []interface{}{1, 2, 3, 4, 5, 8, 6, 7, 9}, Hidden: true, Note: "four levels: 8 (level 2) comes before 6 (level 3) though 6 is further left"},
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2, nil, 3}}, Expected: []interface{}{1, 2, 3}, Hidden: true, Note: "right chain: one node per level"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(w)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef levelOrderValues(root):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When you record a node, its children belong to the next level, but the rest of the current level still has to come first. What structure lets the children wait behind everything already waiting?"},
			{Level: 2, Type: "insight", Content: "A first-in, first-out queue. Start it with the root; repeatedly take the front node, record it, and add its children to the back. Children always join behind every node of the current level, so levels come out in order. This is Chapter 9's level-by-level BFS. In Python use collections.deque, whose popleft is O(1).", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Handle the empty tree: return [].\n2. Seed the queue with the root.\n3. Take the front node and record it.\n4. Queue its children, left then right, skipping missing ones.\n5. Repeat until the queue is empty, then return the list."},
			{Level: 4, Type: "skeleton", Content: "from collections import deque\n\ndef levelOrderValues(root):\n    if root is None:\n        return []\n    out = []\n    queue = deque([root])\n    while queue:\n        node = ____\n        out.append(node.val)\n        if node.left:\n            ____\n        if node.right:\n            ____\n    return out"},
		},
		Solution: Solution{
			Code: `from collections import deque


def levelOrderValues(root):
    if root is None:
        return []
    out = []
    queue = deque([root])
    while queue:
        node = queue.popleft()
        out.append(node.val)
        if node.left:
            queue.append(node.left)
        if node.right:
            queue.append(node.right)
    return out`,
			Explanation: "Breadth-first search on a tree. A queue holds the nodes that have been discovered but not yet recorded; the front is always the shallowest of them, so the values come out level by level, left to right.",
			KeyInsight:  "Depth-first orders use a stack (the call stack), which always resumes the most recent node and so dives down. Swapping the stack for a queue resumes the oldest node instead, which finishes each level before starting the next.",
			WhyItWorks:  "Invariant: the queue holds nodes in order of level, and within a level from left to right, with at most two levels present (the rest of level d, then the part of level d+1 found so far). It holds at the start (just the root). Popping the front node of level d and appending its children, left then right, adds level-(d+1) nodes behind the rest of level d in left-to-right order, because the parents themselves are popped left to right. So nodes are recorded in level order. Each node enters and leaves the queue once: O(n) time, and the queue never holds more than about two levels, O(w).",
			CommonMistakes: []string{
				"Using a list with pop() (a stack): it records the most recently added node, which is depth-first, not level order. Use a deque and popleft().",
				"Using list.pop(0): the order is right, but every pop shifts the whole list, O(n) each. collections.deque does it in O(1).",
				"Queueing None children and then reading `.val` of None. Check each child before appending it.",
			},
			TimeComplexity:  "O(n): each node is queued and recorded once",
			SpaceComplexity: "O(w), where w is the widest level (up to n/2 for a full tree)",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle the empty tree", Explanation: "With no root there is nothing to queue.", CodeSnippet: "if root is None:\n    return []", LineStart: 5, LineEnd: 6},
				{Title: "Seed the queue with the root", Explanation: "The root is the only node of level 0.", CodeSnippet: "queue = deque([root])", LineStart: 8, LineEnd: 8},
				{Title: "Take the front node and record it", Explanation: "popleft takes the oldest waiting node, the leftmost unrecorded node of the shallowest level.", CodeSnippet: "node = queue.popleft()\nout.append(node.val)", LineStart: 10, LineEnd: 11},
				{Title: "Queue its children", Explanation: "On [3,9,20,null,null,15,7]: popping 3 queues 9 and 20; popping 9 queues nothing; popping 20 queues 15 and 7. Children join behind the rest of their parent's level.", CodeSnippet: "if node.left:\n    queue.append(node.left)\nif node.right:\n    queue.append(node.right)", LineStart: 12, LineEnd: 15},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "What happens to the output if you replace popleft() with pop()?", Answer: "The deque becomes a stack, so the newest node is taken first. Because the right child was pushed last, you get node, right subtree, left subtree, a depth-first order (reversed-child preorder). On [1,2,3,4,5,6,7] that is 1,3,7,6,2,5,4 instead of 1..7."},
			{Prompt: "How would you change this to return one list per level?", Answer: "At the top of each round, read len(queue): that is exactly how many nodes the current level has, because the queue holds only that level at that moment. Pop that many nodes into a new inner list, queueing their children as before, then append the inner list. That is Binary Tree Level Order Traversal."},
		},
	},
	{
		ID:              "drill-inorder-iterative",
		Kind:            KindDrill,
		Order:           6,
		Title:           "Inorder Without Recursion",
		Slug:            "drill-inorder-iterative",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "DFS", "Binary Tree", "Stack", "Foundations"},
		RelatedChapters: []int{9, 8, 5},
		Description: `Return the **inorder** values of a binary tree (left subtree, then the node, then the right subtree) **without recursion**: manage the stack yourself.

Recursive code keeps its place on Python's call stack, which has a fixed depth limit. A tree that is one long chain of left children has as many levels as nodes, and one hidden test is a chain of 25,000 nodes, deeper than the limit, so a recursive solution fails it with a RecursionError. An explicit stack is an ordinary list and can grow as large as it needs to.

An empty tree (root is None) gives an empty list.`,
		Constraints: []string{
			"The tree has between 0 and 25,000 nodes",
			"-10^5 <= Node.val <= 10^5",
			"Do not use recursion",
		},
		Examples: []Example{
			{Input: "root = [1,null,2,3]", Output: "[1,3,2]", Explanation: "Push 1 and try to go left: there is nothing, so pop 1 and record it, then move to its right child 2. Push 2, go left to 3, push 3; 3 has no left child, so pop and record 3, then pop and record 2."},
			{Input: "root = [1,2,3,4,5,6,7]", Output: "[4,2,5,1,6,3,7]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2, 3}}, Expected: []interface{}{1, 3, 2}, Note: "example 1: a node with no left child is recorded straight away"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, 5, 6, 7}}, Expected: []interface{}{4, 2, 5, 1, 6, 3, 7}, Note: "example 2: a full tree"},
			{Input: map[string]interface{}{"root": []interface{}{}}, Expected: []interface{}{}, Hidden: true, Note: "empty tree gives []"},
			{Input: map[string]interface{}{"root": []interface{}{1}}, Expected: []interface{}{1}, Hidden: true, Note: "a single node"},
			{Input: map[string]interface{}{"root": []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}}, Expected: []interface{}{4, 2, 6, 5, 7, 1, 3, 9, 8}, Hidden: true, Note: "after a right move, the walk must go all the way left again (5 -> 6)"},
			{Input: map[string]interface{}{"root": []interface{}{1, nil, 2, nil, 3}}, Expected: []interface{}{1, 2, 3}, Hidden: true, Note: "right chain: the stack never holds more than one node"},
			{Input: map[string]interface{}{"root": drillLeftChain(25000)}, Expected: drillCountdown(25000), Hidden: true, Note: "a 25,000-node left chain, deeper than Python's recursion limit: only an explicit stack gets through"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\n# Don't use recursion: manage your own stack.\ndef inorderIterative(root):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In the recursive version, when a call goes left, the node it came from waits on the call stack to be recorded later. If you keep those waiting nodes in your own list, which one do you need back first?"},
			{Level: 2, Type: "insight", Content: "The most recently pushed one, so the list is a stack (Chapter 5). The recursive walk does three things, and you do them by hand: go left as far as you can, pushing each node (they are waiting for their left subtrees); pop one, which means its left subtree is finished, so record it; then move to its right child and repeat. See the explicit-stack inorder in Chapter 8.", ChapterRef: 8},
			{Level: 3, Type: "plan", Content: "1. Start with an empty stack and a pointer at the root.\n2. Loop while the pointer is a node or the stack is not empty.\n3. Go left as far as possible, pushing every node on the way.\n4. Pop a node and record it: everything to its left is done.\n5. Move the pointer to its right child (which may be None) and loop."},
			{Level: 4, Type: "skeleton", Content: "def inorderIterative(root):\n    out = []\n    stack = []\n    node = root\n    while node or stack:\n        while node:\n            ____\n            node = ____\n        node = stack.pop()\n        ____\n        node = ____\n    return out"},
		},
		Solution: Solution{
			Code: `def inorderIterative(root):
    out = []
    stack = []
    node = root
    while node or stack:
        while node:
            stack.append(node)
            node = node.left
        node = stack.pop()
        out.append(node.val)
        node = node.right
    return out`,
			Explanation: "The recursive inorder from the earlier drill, with the call stack made visible. Each node on the stack is a node whose left subtree is still being walked; popping it means that walk has finished, so it is recorded and its right subtree starts.",
			KeyInsight:  "A recursive call that is waiting is just a node that has not been recorded yet. Pushing nodes while going left and popping them when the left side runs out reproduces exactly the order in which the recursive calls would resume, without any limit on depth.",
			WhyItWorks:  "Invariant: the stack holds, bottom to top, the ancestors of the current position whose left subtrees are being walked and who have not been recorded, which is precisely the set of recursive calls that would be paused at visit(node.left). The inner loop pushes the left spine of the current subtree, as the recursion would descend. When it reaches None, the top of the stack has a finished left subtree, so recording it next is correct inorder, and moving to its right child starts that subtree just as the recursive call on node.right would. Every node is pushed and popped once: O(n) time and O(h) stack space, held in a list rather than in call frames.",
			CommonMistakes: []string{
				"Pushing only node.left once instead of looping all the way down the left spine. Every left descendant must be pushed before anything is recorded.",
				"Looping only `while stack:`: after popping the last node the stack is empty, but its right child may still be unvisited. Continue while `node or stack`.",
				"Pushing the right child onto the stack straight away (as in iterative preorder). In inorder the right subtree must start only after the node is recorded, so move the pointer there instead.",
			},
			TimeComplexity:  "O(n): each node is pushed and popped once",
			SpaceComplexity: "O(h) for the explicit stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Start at the root with an empty stack", Explanation: "node is where the walk is now; the stack holds nodes waiting to be recorded.", CodeSnippet: "stack = []\nnode = root", LineStart: 3, LineEnd: 4},
				{Title: "Go left as far as possible", Explanation: "Each node pushed is waiting for its left subtree, just as a recursive call would wait at visit(node.left). On [1,2,3,4,5,6,7] this pushes 1, 2, 4.", CodeSnippet: "while node:\n    stack.append(node)\n    node = node.left", LineStart: 6, LineEnd: 8},
				{Title: "Pop and record", Explanation: "The top node's left subtree is finished, so it is next in inorder: 4, then 2.", CodeSnippet: "node = stack.pop()\nout.append(node.val)", LineStart: 9, LineEnd: 10},
				{Title: "Move to the right subtree", Explanation: "After recording 2, the walk continues at 5, which is pushed by the next inner loop. If the right child is None, the next pop resumes an ancestor.", CodeSnippet: "node = node.right", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Morris traversal (no stack at all)",
					Code: `def inorderIterative(root):
    out = []
    node = root
    while node:
        if node.left is None:
            out.append(node.val)
            node = node.right
        else:
            pred = node.left
            while pred.right and pred.right is not node:
                pred = pred.right
            if pred.right is None:
                pred.right = node
                node = node.left
            else:
                pred.right = None
                out.append(node.val)
                node = node.right
    return out`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Instead of a stack, temporarily point each node's inorder predecessor back at it, so the walk can climb back up; the link is removed on the second visit. O(1) extra space, at the cost of changing the tree while walking it.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the outer loop test `node or stack` and not just `stack`?", Answer: "At the very start the stack is empty but node is the root, and later, right after popping the last waiting node, the stack is empty again while node has moved to that node's right child, which still has to be walked. The walk is only over when there is no current node and no waiting node."},
			{Prompt: "How large does the stack get for a full tree of n nodes, and for a left chain?", Answer: "It never holds more than one node per level of the current path, so its size is at most the height h: about log2(n) for a full tree, and n for a left chain, where every node waits for the one below it. That is the same depth the recursion would reach, but here it lives in a list instead of the call stack."},
		},
	},
	{
		ID:              "drill-insert-into-bst",
		Number:          701,
		Kind:            KindDrill,
		Order:           7,
		Title:           "Insert into a BST",
		Slug:            "insert-into-a-binary-search-tree",
		Difficulty:      "Easy",
		Category:        "trees",
		Tags:            []string{"Tree", "Binary Search Tree", "Binary Tree", "Foundations"},
		RelatedChapters: []int{9},
		Description: `You are given the root of a **binary search tree** (BST) and a value ` + "`val`" + ` that is not in the tree. In a BST, every value in a node's left subtree is smaller than the node's value and every value in its right subtree is larger.

Insert ` + "`val`" + ` as a **new leaf**: start at the root, go left when ` + "`val`" + ` is smaller than the current node and right when it is larger, and attach the new node at the first missing child you reach. Return the root of the tree (for an empty tree, the new node is the root).`,
		Constraints: []string{
			"The tree has between 0 and 10^4 nodes",
			"-10^8 <= Node.val, val <= 10^8",
			"All values are distinct, and val is not in the tree",
		},
		Examples: []Example{
			{Input: "root = [4,2,7,1,3], val = 5", Output: "[4,2,7,1,3,5]", Explanation: "5 > 4, so go right to 7. 5 < 7, so go left; 7 has no left child, so 5 becomes it."},
			{Input: "root = [40,20,60,10,30,50,70], val = 25", Output: "[40,20,60,10,30,50,70,null,null,25]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{4, 2, 7, 1, 3}, "val": 5}, Expected: []interface{}{4, 2, 7, 1, 3, 5}, Note: "example 1: right, then left"},
			{Input: map[string]interface{}{"root": []interface{}{40, 20, 60, 10, 30, 50, 70}, "val": 25}, Expected: []interface{}{40, 20, 60, 10, 30, 50, 70, nil, nil, 25}, Note: "example 2: left, then right, then left"},
			{Input: map[string]interface{}{"root": []interface{}{}, "val": 5}, Expected: []interface{}{5}, Hidden: true, Note: "empty tree: the new node is the root"},
			{Input: map[string]interface{}{"root": []interface{}{1}, "val": 2}, Expected: []interface{}{1, nil, 2}, Hidden: true, Note: "a larger value goes right"},
			{Input: map[string]interface{}{"root": []interface{}{5, 3}, "val": 4}, Expected: []interface{}{5, 3, nil, nil, 4}, Hidden: true, Note: "smaller than the root but larger than its left child"},
			{Input: map[string]interface{}{"root": []interface{}{5, 3, 8}, "val": 1}, Expected: []interface{}{5, 3, 8, 1}, Hidden: true, Note: "the smallest value goes to the far left"},
			{Input: map[string]interface{}{"root": []interface{}{8, nil, 10, nil, 12}, "val": 11}, Expected: []interface{}{8, nil, 10, nil, 12, 11}, Hidden: true, Note: "deep in a right chain, then one step left"},
		},
		TimeComplexity:  "O(h)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef insertIntoBST(root, val):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
			ReturnType: "TreeNode",
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Standing at one node, the BST rule tells you which half of the tree val belongs in. Do you ever need to look at the other half?"},
			{Level: 2, Type: "insight", Content: "No: compare once and go to one child, exactly like BST search in Chapter 9. The insert spot is where that search falls off the tree. Recursively, insert into the correct subtree and store the result back in that child link; an empty subtree returns the new node, which attaches it.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Reach an empty spot: if root is None, return a new node holding val.\n2. Pick a side: smaller goes into the left subtree, larger into the right.\n3. Insert into that subtree and store the returned subtree root back in the child link.\n4. Return this node, which stays the root of its subtree."},
			{Level: 4, Type: "skeleton", Content: "def insertIntoBST(root, val):\n    if root is None:\n        return ____\n    if val < root.val:\n        root.left = ____\n    else:\n        ____\n    return root"},
		},
		Solution: Solution{
			Code: `def insertIntoBST(root, val):
    if root is None:
        return TreeNode(val)
    if val < root.val:
        root.left = insertIntoBST(root.left, val)
    else:
        root.right = insertIntoBST(root.right, val)
    return root`,
			Explanation: "BST search with a write at the end. Each call compares val with one node and recurses into the only subtree val can belong to; the call that reaches an empty spot creates the node, and the assignment in its parent's call links it in.",
			KeyInsight:  "Returning the subtree's root and assigning it back (`root.left = insertIntoBST(root.left, val)`) means no call has to know whether it is the one attaching the node: every call returns its own node unchanged, except the one at the empty spot, which returns the new node.",
			WhyItWorks:  "At each node the BST rule says every value on the left is smaller and every value on the right larger, so val can only belong in one subtree, and putting it there keeps the rule true at this node. By induction, the recursive call returns a valid BST of that subtree plus val, and assigning it back keeps the rest of the tree unchanged. The empty tree plus val is a single node. One node per level is visited: O(h) time and O(h) call stack.",
			CommonMistakes: []string{
				"Calling `insertIntoBST(root.left, val)` without storing the result in root.left: the new node is created and then lost. Assign the returned subtree back.",
				"Returning the new node from the top-level call instead of the original root, which throws away the rest of the tree. Only an empty tree returns the new node.",
				"Comparing with the wrong child, such as going right when val is smaller. Smaller goes left, larger goes right, at every level.",
			},
			TimeComplexity:  "O(h): one comparison per level, about log n when balanced and n for a chain",
			SpaceComplexity: "O(h) for the call stack (the iterative version uses O(1))",
			Walkthrough: []WalkthroughStep{
				{Title: "Reach an empty spot", Explanation: "Falling off the tree means this is where val belongs. The new node is returned to the parent's call, which attaches it.", CodeSnippet: "if root is None:\n    return TreeNode(val)", LineStart: 2, LineEnd: 3},
				{Title: "Pick a side and insert there", Explanation: "On [4,2,7,1,3] with 5: 5 > 4 so go right; 5 < 7 so go left, which is empty, so 7.left becomes the new node.", CodeSnippet: "if val < root.val:\n    root.left = insertIntoBST(root.left, val)\nroot.right = insertIntoBST(root.right, val)", LineStart: 4, LineEnd: 7},
				{Title: "Return this node", Explanation: "This node is still the root of its subtree, so the parent's link to it is unchanged.", CodeSnippet: "return root", LineStart: 8, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{
					Name: "Iterative walk",
					Code: `def insertIntoBST(root, val):
    new_node = TreeNode(val)
    if root is None:
        return new_node
    node = root
    while True:
        if val < node.val:
            if node.left is None:
                node.left = new_node
                return root
            node = node.left
        else:
            if node.right is None:
                node.right = new_node
                return root
            node = node.right`,
					TimeComplexity:  "O(h)",
					SpaceComplexity: "O(1)",
					Explanation:     "Walk down with a pointer and attach the node at the first missing child on the path. No call stack, at the cost of checking for the missing child before each step.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "If you insert 1, 2, 3, 4, 5 in that order into an empty BST, what shape do you get, and what does that do to the running time?", Answer: "Each value is larger than everything before it, so each goes right of the last: a chain of 5 nodes with height 5. Insert and search then take O(n) instead of O(log n). Self-balancing trees (AVL, red-black) exist to stop this, and inserting values in random order avoids it on average."},
			{Prompt: "Why does `root.left = insertIntoBST(root.left, val)` reassign the link even when the left child already existed?", Answer: "Because the recursive call returns the root of the left subtree after the insert, and for an existing child that is the same node, so the assignment changes nothing. Only when the child was None does it return a new node, and then the assignment is what attaches it. Writing it one way for both cases removes a special case."},
		},
	},
	{
		ID:              "drill-delete-from-bst",
		Number:          450,
		Kind:            KindDrill,
		Order:           8,
		Title:           "Delete from a BST",
		Slug:            "delete-node-in-a-bst",
		Difficulty:      "Medium",
		Category:        "trees",
		Tags:            []string{"Tree", "Binary Search Tree", "Binary Tree", "Foundations"},
		RelatedChapters: []int{9},
		Description: `You are given the root of a **binary search tree** (BST: every value in a node's left subtree is smaller, every value in its right subtree larger) and a ` + "`key`" + `. Remove the node holding ` + "`key`" + ` if there is one, keep the tree a valid BST, and return its root.

If ` + "`key`" + ` is not in the tree, return the tree unchanged. Removing a node with no children or one child is a matter of relinking its parent; a node with two children needs a replacement value from one of its subtrees. Any valid BST that holds exactly the remaining values is accepted.`,
		Constraints: []string{
			"The tree has between 0 and 10^4 nodes",
			"-10^5 <= Node.val, key <= 10^5",
			"All values in the tree are distinct",
		},
		Examples: []Example{
			{Input: "root = [5,3,6,2,4,null,7], key = 3", Output: "[5,4,6,2,null,null,7]", Explanation: "3 has two children. Its successor (the smallest value in its right subtree) is 4: copy 4 into the node, then delete the old 4, a leaf. Using the predecessor 2 instead, [5,2,6,null,4,null,7], is also accepted."},
			{Input: "root = [5,3,6,2,4,null,7], key = 0", Output: "[5,3,6,2,4,null,7]", Explanation: "0 is not in the tree, so nothing changes."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"root": []interface{}{5, 3, 6, 2, 4, nil, 7}, "key": 3}, Expected: []interface{}{5, 4, 6, 2, nil, nil, 7}, Note: "example 1: a node with two children"},
			{Input: map[string]interface{}{"root": []interface{}{5, 3, 6, 2, 4, nil, 7}, "key": 0}, Expected: []interface{}{5, 3, 6, 2, 4, nil, 7}, Note: "example 2: key not in the tree"},
			{Input: map[string]interface{}{"root": []interface{}{}, "key": 0}, Expected: []interface{}{}, Hidden: true, Note: "empty tree"},
			{Input: map[string]interface{}{"root": []interface{}{5, 3, 6, 2, 4, nil, 7}, "key": 7}, Expected: []interface{}{5, 3, 6, 2, 4}, Hidden: true, Note: "a leaf: its parent's link becomes None"},
			{Input: map[string]interface{}{"root": []interface{}{5, 3, 6, 2}, "key": 3}, Expected: []interface{}{5, 2, 6}, Hidden: true, Note: "only a left child: it takes the node's place"},
			{Input: map[string]interface{}{"root": []interface{}{5, 3, 6, 2, 4, nil, 7}, "key": 5}, Expected: []interface{}{6, 3, 7, 2, 4}, Hidden: true, Note: "the root, which has two children"},
			{Input: map[string]interface{}{"root": []interface{}{1}, "key": 1}, Expected: []interface{}{}, Hidden: true, Note: "the only node: the tree becomes empty"},
			{Input: map[string]interface{}{"root": []interface{}{8, 3, 12, nil, nil, 10, 14, nil, 11}, "key": 8}, Expected: []interface{}{10, 3, 12, nil, nil, 11, 14}, Hidden: true, Note: "the successor 10 has a right child 11, which must be kept"},
		},
		TimeComplexity:  "O(h)",
		SpaceComplexity: "O(h)",
		StarterCode:     "# TreeNode(val, left, right) is provided\ndef deleteNode(root, key):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"root": "TreeNode"},
			ReturnType: "TreeNode",
			Validator: `def check(raw, args, result, expected):
    want = sorted(v for v in raw['root'] if v is not None and v != raw['key'])
    got = []
    stack = []
    node = result_obj
    while stack or node is not None:
        while node is not None:
            stack.append(node)
            if len(stack) > len(want) + 1:
                return (False, "the returned tree has more nodes than it should (or a cycle)")
            node = node.left
        node = stack.pop()
        got.append(node.val)
        if len(got) > len(want):
            return (False, "the returned tree has more nodes than it should: " + str(result))
        node = node.right
    if got == want:
        return True
    if sorted(got) == got:
        return (False, "the tree should hold " + str(want) + " but holds " + str(got))
    return (False, "not a valid BST: its inorder " + str(got) + " is not sorted")`,
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Removing a leaf, or a node with one child, is easy: what can take its place? For a node with two children, which value in the tree could sit in its spot without breaking the BST rule on either side?"},
			{Level: 2, Type: "insight", Content: "First find the node with BST search (Chapter 9), assigning results back as in Insert into a BST. A node with at most one child is replaced by that child (or None). A node with two children takes the value of its inorder successor, the leftmost node of its right subtree: it is larger than everything on the left and smaller than everything else on the right. Then delete that successor from the right subtree; it has no left child, so that deletion is the easy case.", ChapterRef: 9},
			{Level: 3, Type: "plan", Content: "1. Handle the empty tree: return None (the key was not found).\n2. Search: if key is smaller, delete from the left subtree and assign the result to root.left; if larger, do the same on the right.\n3. Found it, at most one child: return the other child.\n4. Found it, two children: find the leftmost node of the right subtree, copy its value here, and delete that value from the right subtree.\n5. Return this node."},
			{Level: 4, Type: "skeleton", Content: "def deleteNode(root, key):\n    if root is None:\n        return None\n    if key < root.val:\n        root.left = deleteNode(root.left, key)\n    elif key > root.val:\n        root.right = deleteNode(root.right, key)\n    else:\n        if root.left is None:\n            return ____\n        if root.right is None:\n            return ____\n        successor = root.right\n        while ____:\n            successor = successor.left\n        root.val = ____\n        root.right = ____\n    return root"},
		},
		Solution: Solution{
			Code: `def deleteNode(root, key):
    if root is None:
        return None
    if key < root.val:
        root.left = deleteNode(root.left, key)
    elif key > root.val:
        root.right = deleteNode(root.right, key)
    else:
        if root.left is None:
            return root.right
        if root.right is None:
            return root.left
        successor = root.right
        while successor.left:
            successor = successor.left
        root.val = successor.val
        root.right = deleteNode(root.right, successor.val)
    return root`,
			Explanation: "Builds on Insert into a BST: the same search that assigns each recursive result back into the child link. Once the node is found, there are three cases by number of children, and the two-children case is turned into a one-child case by swapping in the inorder successor.",
			KeyInsight:  "The inorder successor of a node with two children is the smallest value in its right subtree. It is larger than everything in the left subtree and smaller than everything else in the right subtree, so it can take the node's place, and since it has no left child, removing it from its old spot is the easy case.",
			WhyItWorks:  "The search visits one node per level and changes nothing until it finds the key; each call returns its subtree's new root and the parent stores it, so only the path is affected. A node with no children or one child can be replaced by that child, because the whole child subtree was already on the correct side of the parent. For two children, copying the successor's value keeps the BST rule at this node (every left value is smaller than it, every remaining right value larger), and deleting the successor from the right subtree is a recursive call on a node with no left child, which returns its right child. Time O(h) for the search plus the walk to the successor, O(h) stack.",
			CommonMistakes: []string{
				"Deleting the successor by setting its parent's left link to None, which loses the successor's right subtree (the [8,3,12,null,null,10,14,null,11] case). Delete it with the same function, which promotes its right child.",
				"Using `deleteNode(root.left, key)` without assigning it back to root.left, so the removal never reaches the tree.",
				"Picking the successor from the left subtree (its largest value is the predecessor, which is fine) but then walking left in it instead of right. The predecessor is the rightmost node of the left subtree.",
			},
			TimeComplexity:  "O(h): one path down to the key and one more down to its successor",
			SpaceComplexity: "O(h) for the call stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Search for the key", Explanation: "Exactly the insert recursion: go to the one side the key can be on, and store the updated subtree back in that link.", CodeSnippet: "if key < root.val:\n    root.left = deleteNode(root.left, key)\nroot.right = deleteNode(root.right, key)", LineStart: 4, LineEnd: 7},
				{Title: "At most one child: promote it", Explanation: "Returning the other child makes the parent link skip this node. For a leaf both are None, so it returns None.", CodeSnippet: "if root.left is None:\n    return root.right\nif root.right is None:\n    return root.left", LineStart: 9, LineEnd: 12},
				{Title: "Two children: find the successor", Explanation: "The leftmost node of the right subtree. On [5,3,6,2,4,null,7] deleting 3, the right subtree is just 4.", CodeSnippet: "successor = root.right\nwhile successor.left:\n    successor = successor.left", LineStart: 13, LineEnd: 15},
				{Title: "Copy it in and delete it below", Explanation: "The node now holds 4, and the old 4 is removed from the right subtree by the same function, which handles its possible right child.", CodeSnippet: "root.val = successor.val\nroot.right = deleteNode(root.right, successor.val)", LineStart: 16, LineEnd: 17},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can the successor never have a left child?", Answer: "It is found by walking left from the right child until there is no left child. If it had a left child, that child would be smaller and still in the right subtree, so it would be the successor instead. That is what makes the second deletion the easy one-child case."},
			{Prompt: "Would the predecessor (the largest value in the left subtree) work just as well? What changes?", Answer: "Yes. The predecessor is larger than every other left value and smaller than every right value, so it can take the node's place too. You find it by walking right from the left child, copy it in, and delete it from the left subtree; it has no right child. Both answers are valid BSTs, which is why this problem uses a validator instead of one exact answer."},
		},
	},
}

// drillLeftChain is a tree of n nodes in which each node is the left child
// of the one before, in level-order form: [1, 2, None, 3, None, ..., n].
func drillLeftChain(n int) []interface{} {
	// Each node's two slots are (next value, None); the last has none.
	out := make([]interface{}, 0, 2*n)
	out = append(out, 1)
	for i := 2; i <= n; i++ {
		out = append(out, i)
		if i < n {
			out = append(out, nil)
		}
	}
	return out
}

// drillCountdown is [n, n-1, ..., 1], the inorder of drillLeftChain(n).
func drillCountdown(n int) []interface{} {
	out := make([]interface{}, n)
	for i := range out {
		out[i] = n - i
	}
	return out
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, TreesDrills...)
}
