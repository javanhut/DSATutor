package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestLinkedListDrillsRejectMistakes checks that each linked-list foundations
// drill fails the mistake its tests and notes promise to catch, so a drill
// can't be passed without doing the thing it drills.
func TestLinkedListDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-build-linked-list", "returns the list instead of nodes", "def fromArray(values):\n    return values"},
		{"drill-build-linked-list", "treats 0 as a missing value", `class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def fromArray(values):
    dummy = ListNode()
    tail = dummy
    for v in values:
        if v:
            tail.next = ListNode(v)
            tail = tail.next
    return dummy.next`},
		{"drill-list-to-array", "stops before the last node", `def toArray(head):
    out = []
    node = head
    while node.next is not None:
        out.append(node.val)
        node = node.next
    return out`},
		{"drill-list-index-of", "returns the last match", `def indexOf(head, target):
    node = head
    index = 0
    found = -1
    while node is not None:
        if node.val == target:
            found = index
        node = node.next
        index += 1
    return found`},
		{"drill-middle-of-list", "returns the first middle on even length", `def middleNode(head):
    slow = fast = head
    while fast is not None and fast.next is not None and fast.next.next is not None:
        slow = slow.next
        fast = fast.next.next
    return slow`},
		{"drill-insert-at-index", "has no dummy, so index 0 fails", `def insertAt(head, index, val):
    prev = head
    for _ in range(index - 1):
        prev = prev.next
    prev.next = ListNode(val, prev.next)
    return head`},
		{"drill-remove-elements", "advances prev after a removal", `def removeElements(head, val):
    dummy = ListNode(0, head)
    prev = dummy
    node = head
    while node is not None:
        if node.val == val:
            prev.next = node.next
        prev = node
        node = node.next
    return dummy.next`},
		{"drill-design-linked-list", "addAtIndex at the length does not append", `class Node:
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
        if index < 0 or index >= self.size and index != 0:
            return
        prev = self._before(index)
        prev.next = Node(val, prev.next)
        self.size += 1

    def deleteAtIndex(self, index):
        if index < 0 or index >= self.size:
            return
        prev = self._before(index)
        prev.next = prev.next.next
        self.size -= 1`},
		{"drill-doubly-linked-list", "never updates the prev links of the following node", `class Node:
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
        node.next = new

    def _remove(self, node):
        node.prev.next = node.next
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
        return out`},
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
