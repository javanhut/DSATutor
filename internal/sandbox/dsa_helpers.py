"""
DSA Helper Classes for DSATutor Sandbox

These classes are available by default in the sandbox environment.
Users can use them to create data structures that will be automatically
visualized during code execution.

Available classes:
    - ListNode: For linked list visualization
    - TreeNode: For binary tree visualization
    - GraphNode: For graph visualization

Example usage:
    # Linked List
    head = ListNode(1)
    head.next = ListNode(2)
    head.next.next = ListNode(3)

    # Binary Tree
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)

    # Graph
    node1 = GraphNode(1)
    node2 = GraphNode(2)
    node1.neighbors = [node2]
"""

from typing import Any, Optional, List


class ListNode:
    """
    Node class for singly linked list.

    Attributes:
        val: The value stored in the node
        next: Reference to the next node (or None)

    Example:
        # Create a linked list: 1 -> 2 -> 3
        head = ListNode(1)
        head.next = ListNode(2)
        head.next.next = ListNode(3)

        # Or create with chaining
        head = ListNode(1, ListNode(2, ListNode(3)))
    """

    def __init__(self, val: Any = 0, next: Optional['ListNode'] = None):
        self.val = val
        self.next = next

    def __repr__(self) -> str:
        return f"ListNode({self.val})"

    def __str__(self) -> str:
        values = []
        current = self
        visited = set()
        while current and id(current) not in visited:
            visited.add(id(current))
            values.append(str(current.val))
            current = current.next
        if current:
            values.append("...")  # Cycle detected
        return " -> ".join(values)

    @classmethod
    def from_list(cls, values: List[Any]) -> Optional['ListNode']:
        """Create a linked list from a Python list."""
        if not values:
            return None
        head = cls(values[0])
        current = head
        for val in values[1:]:
            current.next = cls(val)
            current = current.next
        return head

    def to_list(self) -> List[Any]:
        """Convert linked list to Python list."""
        values = []
        current = self
        visited = set()
        while current and id(current) not in visited:
            visited.add(id(current))
            values.append(current.val)
            current = current.next
        return values


class TreeNode:
    """
    Node class for binary tree.

    Attributes:
        val: The value stored in the node
        left: Reference to left child (or None)
        right: Reference to right child (or None)

    Example:
        # Create a binary tree:
        #       1
        #      / \\
        #     2   3
        #    / \\
        #   4   5

        root = TreeNode(1)
        root.left = TreeNode(2)
        root.right = TreeNode(3)
        root.left.left = TreeNode(4)
        root.left.right = TreeNode(5)
    """

    def __init__(self, val: Any = 0, left: Optional['TreeNode'] = None, right: Optional['TreeNode'] = None):
        self.val = val
        self.left = left
        self.right = right

    def __repr__(self) -> str:
        return f"TreeNode({self.val})"

    def __str__(self) -> str:
        return f"TreeNode({self.val}, left={self.left}, right={self.right})"

    @classmethod
    def from_list(cls, values: List[Any]) -> Optional['TreeNode']:
        """
        Create a binary tree from a level-order list.
        None values represent missing nodes.

        Example:
            TreeNode.from_list([1, 2, 3, None, 4])
            Creates:
                  1
                 / \\
                2   3
                 \\
                  4
        """
        if not values or values[0] is None:
            return None

        root = cls(values[0])
        queue = [root]
        i = 1

        while queue and i < len(values):
            node = queue.pop(0)

            # Left child
            if i < len(values) and values[i] is not None:
                node.left = cls(values[i])
                queue.append(node.left)
            i += 1

            # Right child
            if i < len(values) and values[i] is not None:
                node.right = cls(values[i])
                queue.append(node.right)
            i += 1

        return root

    def to_list(self) -> List[Any]:
        """Convert binary tree to level-order list."""
        if not self:
            return []

        result = []
        queue = [self]

        while queue:
            node = queue.pop(0)
            if node:
                result.append(node.val)
                queue.append(node.left)
                queue.append(node.right)
            else:
                result.append(None)

        # Remove trailing Nones
        while result and result[-1] is None:
            result.pop()

        return result


class GraphNode:
    """
    Node class for graph (adjacency list representation).

    Attributes:
        val: The value/identifier of the node
        neighbors: List of adjacent GraphNode objects

    Example:
        # Create a graph:
        # 1 -- 2
        # |    |
        # 3 -- 4

        n1, n2, n3, n4 = GraphNode(1), GraphNode(2), GraphNode(3), GraphNode(4)
        n1.neighbors = [n2, n3]
        n2.neighbors = [n1, n4]
        n3.neighbors = [n1, n4]
        n4.neighbors = [n2, n3]
    """

    def __init__(self, val: Any, neighbors: Optional[List['GraphNode']] = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []

    def __repr__(self) -> str:
        neighbor_vals = [n.val for n in self.neighbors]
        return f"GraphNode({self.val}, neighbors={neighbor_vals})"

    def __str__(self) -> str:
        return f"GraphNode({self.val})"

    @classmethod
    def from_adjacency_list(cls, adj: dict) -> dict:
        """
        Create graph nodes from an adjacency list dictionary.

        Example:
            adj = {
                1: [2, 3],
                2: [1, 4],
                3: [1, 4],
                4: [2, 3]
            }
            nodes = GraphNode.from_adjacency_list(adj)
            # nodes[1] is the GraphNode for value 1
        """
        # First pass: create all nodes
        nodes = {val: cls(val) for val in adj}

        # Second pass: connect neighbors
        for val, neighbor_vals in adj.items():
            nodes[val].neighbors = [nodes[nv] for nv in neighbor_vals if nv in nodes]

        return nodes


class MinHeap:
    """
    Min-heap implementation for visualization.

    The heap is backed by a list, making it easy to visualize
    both as a tree and as an array.

    Example:
        heap = MinHeap()
        heap.push(3)
        heap.push(1)
        heap.push(2)
        print(heap.pop())  # 1
        print(heap.pop())  # 2
    """

    def __init__(self, initial: Optional[List[Any]] = None):
        self.data = []
        if initial:
            for item in initial:
                self.push(item)

    def __len__(self) -> int:
        return len(self.data)

    def __bool__(self) -> bool:
        return len(self.data) > 0

    def __repr__(self) -> str:
        return f"MinHeap({self.data})"

    def push(self, val: Any) -> None:
        """Add a value to the heap."""
        self.data.append(val)
        self._sift_up(len(self.data) - 1)

    def pop(self) -> Any:
        """Remove and return the minimum value."""
        if not self.data:
            raise IndexError("pop from empty heap")

        min_val = self.data[0]
        last = self.data.pop()

        if self.data:
            self.data[0] = last
            self._sift_down(0)

        return min_val

    def peek(self) -> Any:
        """Return the minimum value without removing it."""
        if not self.data:
            raise IndexError("peek at empty heap")
        return self.data[0]

    def _sift_up(self, idx: int) -> None:
        """Move element up to maintain heap property."""
        parent = (idx - 1) // 2
        while idx > 0 and self.data[idx] < self.data[parent]:
            self.data[idx], self.data[parent] = self.data[parent], self.data[idx]
            idx = parent
            parent = (idx - 1) // 2

    def _sift_down(self, idx: int) -> None:
        """Move element down to maintain heap property."""
        size = len(self.data)
        while True:
            smallest = idx
            left = 2 * idx + 1
            right = 2 * idx + 2

            if left < size and self.data[left] < self.data[smallest]:
                smallest = left
            if right < size and self.data[right] < self.data[smallest]:
                smallest = right

            if smallest == idx:
                break

            self.data[idx], self.data[smallest] = self.data[smallest], self.data[idx]
            idx = smallest


class MaxHeap:
    """
    Max-heap implementation for visualization.

    Example:
        heap = MaxHeap()
        heap.push(1)
        heap.push(3)
        heap.push(2)
        print(heap.pop())  # 3
        print(heap.pop())  # 2
    """

    def __init__(self, initial: Optional[List[Any]] = None):
        self.data = []
        if initial:
            for item in initial:
                self.push(item)

    def __len__(self) -> int:
        return len(self.data)

    def __bool__(self) -> bool:
        return len(self.data) > 0

    def __repr__(self) -> str:
        return f"MaxHeap({self.data})"

    def push(self, val: Any) -> None:
        """Add a value to the heap."""
        self.data.append(val)
        self._sift_up(len(self.data) - 1)

    def pop(self) -> Any:
        """Remove and return the maximum value."""
        if not self.data:
            raise IndexError("pop from empty heap")

        max_val = self.data[0]
        last = self.data.pop()

        if self.data:
            self.data[0] = last
            self._sift_down(0)

        return max_val

    def peek(self) -> Any:
        """Return the maximum value without removing it."""
        if not self.data:
            raise IndexError("peek at empty heap")
        return self.data[0]

    def _sift_up(self, idx: int) -> None:
        """Move element up to maintain heap property."""
        parent = (idx - 1) // 2
        while idx > 0 and self.data[idx] > self.data[parent]:
            self.data[idx], self.data[parent] = self.data[parent], self.data[idx]
            idx = parent
            parent = (idx - 1) // 2

    def _sift_down(self, idx: int) -> None:
        """Move element down to maintain heap property."""
        size = len(self.data)
        while True:
            largest = idx
            left = 2 * idx + 1
            right = 2 * idx + 2

            if left < size and self.data[left] > self.data[largest]:
                largest = left
            if right < size and self.data[right] > self.data[largest]:
                largest = right

            if largest == idx:
                break

            self.data[idx], self.data[largest] = self.data[largest], self.data[idx]
            idx = largest
