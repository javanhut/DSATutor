# Python Sandbox

The Python Sandbox is an interactive feature that allows you to write custom Python algorithms and visualize their execution step-by-step with automatic data structure detection.

## Overview

The sandbox provides:
- A code editor for writing Python algorithms
- Automatic line-by-line execution tracing
- Visual detection and rendering of common data structures
- Step-by-step playback controls
- Variable state tracking at each step

## Accessing the Sandbox

1. Click the **Sandbox** button (marked with "S") in the sidebar footer
2. The UI switches to sandbox mode, replacing the chapter view
3. Write your Python code in the editor
4. Click **Run Code** to execute and visualize

## Code Editor Features

The sandbox code editor includes Python-specific features to make writing code easier:

### Automatic Indentation

- **Tab**: Inserts 4 spaces (Python standard indentation)
- **Shift+Tab**: Removes 4 spaces (dedent)
- **Enter after colon**: Automatically increases indentation after lines ending with `:` (for `def`, `if`, `for`, `while`, `class`, etc.)
- **Enter**: Maintains current indentation level on new lines
- **Smart Backspace**: In leading whitespace, deletes entire indent units (4 spaces at a time)

### Multi-line Selection

- Select multiple lines and press **Tab** to indent all selected lines
- Select multiple lines and press **Shift+Tab** to dedent all selected lines

## Available Helper Classes

The sandbox provides pre-defined helper classes for creating visualizable data structures:

### ListNode

For creating linked lists:

```python
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
```

Example usage:
```python
# Create a linked list: 1 -> 2 -> 3
head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)

# Or use chaining
head = ListNode(1, ListNode(2, ListNode(3)))
```

### TreeNode

For creating binary trees:

```python
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
```

Example usage:
```python
# Create a binary tree:
#       1
#      / \
#     2   3
root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
```

### GraphNode

For creating graphs:

```python
class GraphNode:
    def __init__(self, val, neighbors=None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
```

Example usage:
```python
# Create a graph with two connected nodes
node1 = GraphNode(1)
node2 = GraphNode(2)
node1.neighbors = [node2]
node2.neighbors = [node1]
```

## Automatic Structure Detection

The sandbox automatically detects and visualizes these data structures:

| Python Type | Visualization |
|-------------|---------------|
| `list` (primitives) | Array with index/pointer highlights |
| `dict` (scalar values) | Hash table with key-value pairs |
| `dict` (list values) | Graph (adjacency list representation) |
| `set` | Set visualization |
| Object with `val`/`next` | Linked list |
| Object with `val`/`left`/`right` | Binary tree |
| Object with `val`/`neighbors` | Graph |

## Pointer Highlighting

The sandbox automatically detects common pointer variable names and highlights the corresponding array indices:

- `low`, `high` - Shown in green/red for binary search boundaries
- `mid`, `current`, `i` - Shown in blue for current position
- `j` - Shown in orange for secondary pointer
- `left`, `right`, `start`, `end`, `pivot`, etc.

## Available Modules

The following modules are available in the sandbox:

- `math` - Mathematical functions
- `random` - Random number generation
- `itertools` - Iterator building blocks
- `functools` - Higher-order functions
- `heapq` - Heap queue operations (`heappush`, `heappop`, `heapify`)
- `collections.deque` - Double-ended queue
- `collections.defaultdict` - Dict with default factory
- `collections.Counter` - Element counting

## Example Algorithms

### Binary Search

```python
def binary_search(arr, target):
    low, high = 0, len(arr) - 1
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1
    return -1

arr = [1, 3, 5, 7, 9, 11, 13]
result = binary_search(arr, 7)
print(f"Found at index: {result}")
```

### Linked List Traversal

```python
# Create linked list
head = ListNode(1, ListNode(2, ListNode(3, ListNode(4))))

# Traverse
current = head
while current:
    print(current.val)
    current = current.next
```

### Binary Tree Operations

```python
# Create tree
root = TreeNode(4)
root.left = TreeNode(2)
root.right = TreeNode(6)
root.left.left = TreeNode(1)
root.left.right = TreeNode(3)

# Inorder traversal
def inorder(node):
    if node:
        inorder(node.left)
        print(node.val)
        inorder(node.right)

inorder(root)
```

### Graph BFS

```python
from collections import deque

# Create graph as adjacency list
graph = {
    'A': ['B', 'C'],
    'B': ['A', 'D', 'E'],
    'C': ['A', 'F'],
    'D': ['B'],
    'E': ['B', 'F'],
    'F': ['C', 'E']
}

def bfs(graph, start):
    visited = set()
    queue = deque([start])
    visited.add(start)

    while queue:
        node = queue.popleft()
        print(node)

        for neighbor in graph[node]:
            if neighbor not in visited:
                visited.add(neighbor)
                queue.append(neighbor)

bfs(graph, 'A')
```

## Limitations

### Security Restrictions

For security, the following are blocked:
- File I/O operations (`open`, file paths)
- System access (`os`, `sys`, `subprocess`)
- Network operations (`socket`, `urllib`, `requests`)
- Code execution (`exec`, `eval`, `compile`)
- Import manipulation (`__import__`, `importlib`)

### Resource Limits

| Resource | Limit |
|----------|-------|
| Execution time | 10 seconds |
| Maximum steps | 1000 |
| Code size | 10 KB |
| Output size | 10 KB |

### Blocked Patterns

The following patterns are blocked for security:
- Dunder attribute access (`__class__`, `__bases__`, `__globals__`, etc.)
- Attribute manipulation functions (`getattr`, `setattr`, `delattr`)
- Introspection builtins (`globals`, `locals`, `vars`, `dir`)

## Playback Controls

After running code:

- **Step** - Advance one step
- **Back** - Go back one step
- **Play/Pause** - Auto-play through steps
- **Reset** - Go back to step 1
- **Speed** - Adjust playback speed (1-10)

## Troubleshooting

### "Import not allowed" error

The sandbox restricts certain imports for security. Use only the allowed modules listed above.

### "Execution timed out" error

Your code took too long to execute. Check for infinite loops or reduce the problem size.

### "Exceeded step limit" error

The code executed more than 1000 steps. This usually indicates an infinite loop or very long algorithm. Reduce the input size.

### No visualization appears

Make sure your code creates recognizable data structures (lists, dicts, or node objects with standard attributes). Simple variables without data structures won't generate visualizations.

## Technical Details

### How It Works

1. Your Python code is sent to the Go backend
2. The backend validates the code for security
3. A Python subprocess runs the code with `sys.settrace` enabled
4. Each line execution captures:
   - Line number
   - Function name
   - Local variables
   - Detected data structures
5. The trace is returned as JSON
6. The frontend renders each step as an animation frame

### Files

- `internal/sandbox/tracer.py` - Python execution tracer
- `internal/sandbox/dsa_helpers.py` - Helper class definitions
- `internal/sandbox/sandbox.go` - Go execution handler
- `internal/sandbox/security.go` - Code validation
- `internal/web/static/app.js` - SandboxAnimator class
