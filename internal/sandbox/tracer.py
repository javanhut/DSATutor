#!/usr/bin/env python3
"""
DSATutor Python Tracer

Uses sys.settrace to capture line-by-line execution of Python code,
detecting and serializing data structures for visualization.
"""

import sys
import json
import copy
import io
from contextlib import redirect_stdout, redirect_stderr
from typing import Any, Dict, List, Optional, Set, Tuple

# Maximum steps to prevent infinite loops
MAX_STEPS = 1000

# Maximum recursion depth for serialization
MAX_SERIALIZE_DEPTH = 10

# Maximum array/list size to serialize
MAX_COLLECTION_SIZE = 100


class ListNode:
    """Helper class for linked list visualization."""
    def __init__(self, val: Any = 0, next: Optional['ListNode'] = None):
        self.val = val
        self.next = next

    def __repr__(self):
        return f"ListNode({self.val})"


class TreeNode:
    """Helper class for binary tree visualization."""
    def __init__(self, val: Any = 0, left: Optional['TreeNode'] = None, right: Optional['TreeNode'] = None):
        self.val = val
        self.left = left
        self.right = right

    def __repr__(self):
        return f"TreeNode({self.val})"


class GraphNode:
    """Helper class for graph visualization."""
    def __init__(self, val: Any, neighbors: Optional[List['GraphNode']] = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []

    def __repr__(self):
        return f"GraphNode({self.val})"


class DSATracer:
    """Traces Python code execution and captures state at each line."""

    def __init__(self, user_code: str, user_filename: str = "<user_code>"):
        self.user_code = user_code
        self.user_filename = user_filename
        self.steps: List[Dict] = []
        self.step_count = 0
        self.stopped = False
        self.error: Optional[str] = None
        self.code_lines = user_code.split('\n')

    def trace_function(self, frame, event: str, arg) -> Optional[callable]:
        """Trace function called by sys.settrace on each event."""
        if self.stopped:
            return None

        # Only trace user code, not imports or builtins
        if frame.f_code.co_filename != self.user_filename:
            return self.trace_function

        # Only capture on 'line' events (not call, return, exception)
        if event == 'line':
            self.capture_step(frame)

        if event == 'exception':
            exc_type, exc_value, exc_tb = arg
            self.error = f"{exc_type.__name__}: {exc_value}"
            self.stopped = True
            return None

        return self.trace_function

    def capture_step(self, frame) -> None:
        """Capture the current execution state."""
        if self.step_count >= MAX_STEPS:
            self.stopped = True
            self.error = f"Execution stopped: exceeded {MAX_STEPS} steps"
            return

        self.step_count += 1

        line_num = frame.f_lineno
        func_name = frame.f_code.co_name

        # Get local variables (excluding internal variables)
        locals_dict = {
            k: v for k, v in frame.f_locals.items()
            if not k.startswith('_') and k not in ('self', 'cls')
        }

        # Serialize variables
        serialized_locals = self.serialize_variables(locals_dict)

        # Detect data structures
        structures = self.detect_structures(locals_dict)

        # Get current line text
        line_text = ""
        if 0 < line_num <= len(self.code_lines):
            line_text = self.code_lines[line_num - 1].strip()

        step = {
            "lineNum": line_num,
            "function": func_name if func_name != "<module>" else "main",
            "state": f"Line {line_num}: {line_text[:50]}..." if len(line_text) > 50 else f"Line {line_num}: {line_text}",
            "locals": serialized_locals,
            "structures": structures
        }

        self.steps.append(step)

    def serialize_variables(self, variables: Dict[str, Any]) -> Dict[str, Any]:
        """Serialize variables to JSON-safe format."""
        result = {}
        for name, value in variables.items():
            try:
                result[name] = self.serialize_value(value, depth=0)
            except Exception:
                result[name] = f"<unserializable: {type(value).__name__}>"
        return result

    def serialize_value(self, value: Any, depth: int = 0) -> Any:
        """Recursively serialize a value to JSON-safe format."""
        if depth > MAX_SERIALIZE_DEPTH:
            return "<max depth exceeded>"

        # Primitives
        if value is None:
            return None
        if isinstance(value, bool):
            return value
        if isinstance(value, (int, float)):
            if isinstance(value, float) and (value != value):  # NaN check
                return "NaN"
            return value
        if isinstance(value, str):
            return value[:200] + "..." if len(value) > 200 else value

        # Collections
        if isinstance(value, (list, tuple)):
            if len(value) > MAX_COLLECTION_SIZE:
                serialized = [self.serialize_value(v, depth + 1) for v in value[:MAX_COLLECTION_SIZE]]
                return serialized + [f"... ({len(value) - MAX_COLLECTION_SIZE} more)"]
            return [self.serialize_value(v, depth + 1) for v in value]

        if isinstance(value, set):
            lst = sorted(list(value)[:MAX_COLLECTION_SIZE], key=lambda x: str(x))
            return {"__type__": "set", "values": [self.serialize_value(v, depth + 1) for v in lst]}

        if isinstance(value, dict):
            if len(value) > MAX_COLLECTION_SIZE:
                items = list(value.items())[:MAX_COLLECTION_SIZE]
                result = {str(k): self.serialize_value(v, depth + 1) for k, v in items}
                result["__truncated__"] = len(value) - MAX_COLLECTION_SIZE
                return result
            return {str(k): self.serialize_value(v, depth + 1) for k, v in value.items()}

        # Custom node classes
        if self.is_list_node(value):
            return {"__type__": "ListNode", "val": self.serialize_value(getattr(value, 'val', None), depth + 1)}
        if self.is_tree_node(value):
            return {"__type__": "TreeNode", "val": self.serialize_value(getattr(value, 'val', None), depth + 1)}
        if self.is_graph_node(value):
            return {"__type__": "GraphNode", "val": self.serialize_value(getattr(value, 'val', None), depth + 1)}

        # Fallback: use repr
        try:
            return repr(value)[:100]
        except Exception:
            return f"<{type(value).__name__}>"

    def is_list_node(self, obj: Any) -> bool:
        """Check if object looks like a linked list node."""
        return hasattr(obj, 'val') and hasattr(obj, 'next')

    def is_tree_node(self, obj: Any) -> bool:
        """Check if object looks like a binary tree node."""
        return hasattr(obj, 'val') and hasattr(obj, 'left') and hasattr(obj, 'right')

    def is_graph_node(self, obj: Any) -> bool:
        """Check if object looks like a graph node."""
        return hasattr(obj, 'val') and hasattr(obj, 'neighbors')

    def detect_structures(self, variables: Dict[str, Any]) -> List[Dict]:
        """Detect and serialize data structures in variables."""
        structures = []

        for name, value in variables.items():
            struct = self.classify_and_serialize_structure(name, value, variables)
            if struct:
                structures.append(struct)

        return structures

    def classify_and_serialize_structure(self, name: str, value: Any, all_locals: Dict) -> Optional[Dict]:
        """Classify a value and serialize it as a visualizable structure."""

        # Matrix (2D array) detection - check before 1D array
        if isinstance(value, list) and len(value) > 0 and all(isinstance(row, list) for row in value[:20]):
            # Check if it's a proper 2D matrix with primitive elements
            is_matrix = True
            for row in value[:20]:
                if not all(isinstance(v, (int, float, str, bool, type(None))) for v in row[:20]):
                    is_matrix = False
                    break
            if is_matrix:
                return self.serialize_matrix(name, value, all_locals)

        # Array detection (1D)
        if isinstance(value, list) and all(isinstance(v, (int, float, str, bool, type(None))) for v in value[:20]):
            highlights = self.detect_array_highlights(name, value, all_locals)
            return {
                "name": name,
                "type": "array",
                "data": value[:MAX_COLLECTION_SIZE],
                "highlights": highlights
            }

        # Dictionary/Hash table detection
        if isinstance(value, dict):
            # Check if it looks like an adjacency list (graph)
            if value and all(isinstance(v, (list, set)) for v in value.values()):
                return self.serialize_graph_adjacency(name, value)
            # Regular hash table
            return {
                "name": name,
                "type": "hash_table",
                "data": {str(k): self.serialize_value(v, 0) for k, v in list(value.items())[:MAX_COLLECTION_SIZE]}
            }

        # Linked list detection
        if self.is_list_node(value):
            return self.serialize_linked_list(name, value)

        # Binary tree detection
        if self.is_tree_node(value):
            return self.serialize_binary_tree(name, value)

        # Graph node detection
        if self.is_graph_node(value):
            return self.serialize_graph_from_node(name, value)

        # Set detection
        if isinstance(value, set):
            return {
                "name": name,
                "type": "set",
                "data": sorted([self.serialize_value(v, 0) for v in list(value)[:MAX_COLLECTION_SIZE]], key=str)
            }

        return None

    def detect_array_highlights(self, array_name: str, array: List, all_locals: Dict) -> Dict[str, int]:
        """Detect pointer variables that might reference array indices."""
        highlights = {}

        # Common pointer variable names
        pointer_vars = [
            'i', 'j', 'k', 'low', 'high', 'mid', 'left', 'right',
            'start', 'end', 'current', 'prev', 'next_idx',
            'pivot', 'min_idx', 'max_idx', 'ptr', 'slow', 'fast',
            'l', 'r', 'head', 'tail', 'front', 'back'
        ]

        for var_name in pointer_vars:
            if var_name in all_locals:
                val = all_locals[var_name]
                if isinstance(val, int) and 0 <= val < len(array):
                    highlights[var_name] = val

        return highlights

    def serialize_matrix(self, name: str, matrix: List[List], all_locals: Dict) -> Dict:
        """Serialize a 2D matrix for visualization."""
        # Limit matrix size for display
        max_rows = min(len(matrix), 20)
        max_cols = 20

        data = []
        for row in matrix[:max_rows]:
            row_data = [self.serialize_value(v, 0) for v in row[:max_cols]]
            data.append(row_data)

        # Detect row/column pointer highlights
        highlights = self.detect_matrix_highlights(name, matrix, all_locals)

        return {
            "name": name,
            "type": "matrix",
            "data": data,
            "rows": len(matrix),
            "cols": len(matrix[0]) if matrix else 0,
            "highlights": highlights
        }

    def detect_matrix_highlights(self, matrix_name: str, matrix: List[List], all_locals: Dict) -> Dict:
        """Detect row/column pointer variables for matrix highlighting."""
        highlights = {"row": {}, "col": {}, "cell": []}

        # Common row index variable names
        row_vars = ['r', 'row', 'i']
        # Common column index variable names
        col_vars = ['c', 'col', 'j']

        rows = len(matrix)
        cols = len(matrix[0]) if matrix else 0

        for var_name in row_vars:
            if var_name in all_locals:
                val = all_locals[var_name]
                if isinstance(val, int) and 0 <= val < rows:
                    highlights["row"][var_name] = val

        for var_name in col_vars:
            if var_name in all_locals:
                val = all_locals[var_name]
                if isinstance(val, int) and 0 <= val < cols:
                    highlights["col"][var_name] = val

        # If we have both row and column pointers, mark the cell
        if highlights["row"] and highlights["col"]:
            for r_name, r_val in highlights["row"].items():
                for c_name, c_val in highlights["col"].items():
                    highlights["cell"].append({"row": r_val, "col": c_val, "label": f"{r_name},{c_name}"})

        return highlights

    def serialize_linked_list(self, name: str, head) -> Dict:
        """Traverse and serialize a linked list."""
        nodes = []
        visited: Set[int] = set()
        current = head

        while current is not None and id(current) not in visited and len(nodes) < MAX_COLLECTION_SIZE:
            visited.add(id(current))
            val = getattr(current, 'val', None)
            nodes.append({"val": self.serialize_value(val, 0)})
            current = getattr(current, 'next', None)

        has_cycle = current is not None and id(current) in visited

        return {
            "name": name,
            "type": "linked_list",
            "data": {"nodes": nodes, "hasCycle": has_cycle}
        }

    def serialize_binary_tree(self, name: str, root) -> Dict:
        """Serialize a binary tree for visualization."""
        def serialize_node(node, depth: int = 0, position: int = 0) -> Optional[Dict]:
            if node is None or depth > 10:
                return None

            val = getattr(node, 'val', None)
            left = getattr(node, 'left', None)
            right = getattr(node, 'right', None)

            return {
                "val": self.serialize_value(val, 0),
                "left": serialize_node(left, depth + 1, position * 2),
                "right": serialize_node(right, depth + 1, position * 2 + 1),
                "depth": depth,
                "position": position
            }

        return {
            "name": name,
            "type": "binary_tree",
            "data": serialize_node(root)
        }

    def serialize_graph_adjacency(self, name: str, adj: Dict) -> Dict:
        """Serialize an adjacency list representation of a graph."""
        nodes = list(adj.keys())[:MAX_COLLECTION_SIZE]
        edges = []

        for node in nodes:
            neighbors = adj.get(node, [])
            for neighbor in list(neighbors)[:50]:
                edges.append({"from": str(node), "to": str(neighbor)})

        return {
            "name": name,
            "type": "graph",
            "data": {"nodes": [str(n) for n in nodes], "edges": edges}
        }

    def serialize_graph_from_node(self, name: str, start_node) -> Dict:
        """Serialize a graph starting from a GraphNode."""
        nodes = []
        edges = []
        visited: Set[int] = set()
        queue = [start_node]

        while queue and len(nodes) < MAX_COLLECTION_SIZE:
            node = queue.pop(0)
            if id(node) in visited:
                continue
            visited.add(id(node))

            val = str(getattr(node, 'val', len(nodes)))
            nodes.append(val)

            neighbors = getattr(node, 'neighbors', [])
            for neighbor in neighbors:
                neighbor_val = str(getattr(neighbor, 'val', '?'))
                edges.append({"from": val, "to": neighbor_val})
                if id(neighbor) not in visited:
                    queue.append(neighbor)

        return {
            "name": name,
            "type": "graph",
            "data": {"nodes": nodes, "edges": edges}
        }

    def run(self) -> Dict:
        """Execute the user code with tracing and return results."""
        stdout_capture = io.StringIO()
        stderr_capture = io.StringIO()

        # Build execution environment with safe builtins and helper classes
        safe_builtins = self.get_safe_builtins()
        exec_globals = {
            '__builtins__': safe_builtins,
            '__name__': '__main__',
            '__file__': self.user_filename,
            # Helper classes
            'ListNode': ListNode,
            'TreeNode': TreeNode,
            'GraphNode': GraphNode,
            # Pre-imported modules available at top level
            'math': safe_builtins['math'],
            'random': safe_builtins['random'],
            'itertools': safe_builtins['itertools'],
            'functools': safe_builtins['functools'],
            'heapq': safe_builtins['heapq'],
            'collections': safe_builtins['collections'],
            # Common collections available directly
            'deque': safe_builtins['deque'],
            'defaultdict': safe_builtins['defaultdict'],
            'Counter': safe_builtins['Counter'],
            'OrderedDict': safe_builtins['OrderedDict'],
        }

        try:
            # Compile the code
            code_obj = compile(self.user_code, self.user_filename, 'exec')

            # Set the trace function
            sys.settrace(self.trace_function)

            # Execute with captured output
            with redirect_stdout(stdout_capture), redirect_stderr(stderr_capture):
                exec(code_obj, exec_globals)

        except SyntaxError as e:
            self.error = f"SyntaxError: {e.msg} (line {e.lineno})"
        except Exception as e:
            if not self.error:
                self.error = f"{type(e).__name__}: {e}"
        finally:
            sys.settrace(None)

        output = stdout_capture.getvalue()
        if stderr_capture.getvalue():
            output += "\n[stderr]\n" + stderr_capture.getvalue()

        # Truncate output if too long
        if len(output) > 10000:
            output = output[:10000] + "\n... (output truncated)"

        return {
            "success": self.error is None,
            "steps": self.steps,
            "output": output,
            "error": self.error
        }

    def get_safe_builtins(self) -> Dict:
        """Return a restricted set of safe builtins."""
        import builtins
        import collections
        import heapq as heapq_module
        import math
        import random
        import itertools
        import functools

        # Pre-imported safe modules
        safe_modules = {
            'math': math,
            'random': random,
            'itertools': itertools,
            'functools': functools,
            'collections': collections,
            'heapq': heapq_module,
        }

        def restricted_import(name, globals=None, locals=None, fromlist=(), level=0):
            """Restricted import that only allows safe modules."""
            if name in safe_modules:
                return safe_modules[name]
            raise ImportError(f"Import of '{name}' is not allowed. Available modules: {', '.join(safe_modules.keys())}")

        safe = {
            # Constants
            'True': True,
            'False': False,
            'None': None,

            # Type constructors
            'bool': bool,
            'int': int,
            'float': float,
            'str': str,
            'list': list,
            'dict': dict,
            'set': set,
            'tuple': tuple,
            'frozenset': frozenset,
            'bytes': bytes,
            'bytearray': bytearray,
            'complex': complex,

            # Built-in functions
            'abs': abs,
            'all': all,
            'any': any,
            'bin': bin,
            'callable': callable,
            'chr': chr,
            'divmod': divmod,
            'enumerate': enumerate,
            'filter': filter,
            'format': format,
            'hash': hash,
            'hex': hex,
            'id': id,
            'isinstance': isinstance,
            'issubclass': issubclass,
            'iter': iter,
            'len': len,
            'map': map,
            'max': max,
            'min': min,
            'next': next,
            'oct': oct,
            'ord': ord,
            'pow': pow,
            'print': print,
            'range': range,
            'repr': repr,
            'reversed': reversed,
            'round': round,
            'slice': slice,
            'sorted': sorted,
            'sum': sum,
            'type': type,
            'zip': zip,

            # Exceptions (for catching)
            'Exception': Exception,
            'ValueError': ValueError,
            'TypeError': TypeError,
            'IndexError': IndexError,
            'KeyError': KeyError,
            'AttributeError': AttributeError,
            'StopIteration': StopIteration,
            'RuntimeError': RuntimeError,
            'ZeroDivisionError': ZeroDivisionError,

            # Useful modules (as objects, not import)
            'math': math,
            'random': random,
            'itertools': itertools,
            'functools': functools,

            # Collections
            'deque': collections.deque,
            'defaultdict': collections.defaultdict,
            'Counter': collections.Counter,
            'OrderedDict': collections.OrderedDict,

            # Heap operations
            'heapq': heapq_module,
            'heappush': heapq_module.heappush,
            'heappop': heapq_module.heappop,
            'heapify': heapq_module.heapify,

            # Collections module (for from collections import ...)
            'collections': collections,

            # Restricted import function (allows only safe modules)
            '__import__': restricted_import,
        }

        return safe


def main():
    """Main entry point for the tracer."""
    if len(sys.argv) < 2:
        result = {
            "success": False,
            "steps": [],
            "output": "",
            "error": "No code file provided"
        }
        print(json.dumps(result))
        return

    code_file = sys.argv[1]

    try:
        with open(code_file, 'r') as f:
            user_code = f.read()
    except Exception as e:
        result = {
            "success": False,
            "steps": [],
            "output": "",
            "error": f"Failed to read code file: {e}"
        }
        print(json.dumps(result))
        return

    tracer = DSATracer(user_code, "<user_code>")
    result = tracer.run()

    print(json.dumps(result))


if __name__ == '__main__':
    main()
