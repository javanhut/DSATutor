package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestDrillsRejectMistakes checks that each foundations drill fails the
// mistake its tests and notes promise to catch, so a drill can't be passed
// without doing the thing it drills.
func TestDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-build-binary-tree", "returns the list instead of nodes", "def fromLevelOrder(values):\n    return values"},
		{"drill-build-binary-tree", "treats 0 as a missing node", `from collections import deque

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val, self.left, self.right = val, left, right

def fromLevelOrder(values):
    if not values or values[0] is None:
        return None
    root = TreeNode(values[0])
    queue = deque([root])
    i = 1
    while queue and i < len(values):
        node = queue.popleft()
        if values[i]:
            node.left = TreeNode(values[i])
            queue.append(node.left)
        i += 1
        if i < len(values) and values[i]:
            node.right = TreeNode(values[i])
            queue.append(node.right)
        i += 1
    return root`},
		{"drill-build-binary-tree", "uses the 2*i+1 array formula", `class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val, self.left, self.right = val, left, right

def fromLevelOrder(values):
    def make(i):
        if i >= len(values) or values[i] is None:
            return None
        return TreeNode(values[i], make(2 * i + 1), make(2 * i + 2))
    return make(0)`},
		{"drill-inorder-traversal", "records in preorder", `def inorderTraversal(root):
    out = []
    def visit(node):
        if node:
            out.append(node.val)
            visit(node.left)
            visit(node.right)
    visit(root)
    return out`},
		{"drill-inorder-iterative", "recurses", `def inorderIterative(root):
    out = []
    def visit(node):
        if node:
            visit(node.left)
            out.append(node.val)
            visit(node.right)
    visit(root)
    return out`},
		{"drill-level-order-values", "uses a stack instead of a queue", `def levelOrderValues(root):
    out, stack = [], [root] if root else []
    while stack:
        node = stack.pop()
        out.append(node.val)
        if node.left:
            stack.append(node.left)
        if node.right:
            stack.append(node.right)
    return out`},
		{"drill-postorder-traversal", "reverses preorder", `def postorderTraversal(root):
    out = []
    def visit(node):
        if node:
            out.append(node.val)
            visit(node.left)
            visit(node.right)
    visit(root)
    return out[::-1]`},
		{"drill-delete-from-bst", "drops the successor's right child", `def deleteNode(root, key):
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
        parent, successor = root, root.right
        while successor.left:
            parent, successor = successor, successor.left
        root.val = successor.val
        if parent is root:
            root.right = successor.right
        else:
            parent.left = None
    return root`},
	}
	for _, c := range cases {
		c := c
		t.Run(c.id+"/"+c.why, func(t *testing.T) {
			t.Parallel()
			var p *practice.Problem
			for _, q := range practice.EmbeddedProblems {
				if q.ID == c.id {
					p = q
				}
			}
			if p == nil {
				t.Fatalf("no problem %s", c.id)
			}
			for i := range p.TestCases {
				res, err := Run(p, c.code, &p.TestCases[i], Options{})
				if err != nil {
					t.Fatalf("case %d: %v", i+1, err)
				}
				if !res.Passed {
					return
				}
			}
			t.Errorf("code that %s passed every case", c.why)
		})
	}
}
