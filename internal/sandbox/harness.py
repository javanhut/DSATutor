"""
DSATutor test harness.

Runs a learner's solution against one test case. The Go side writes a JSON
spec describing how to call the code (plain function, design class, or an
encode/decode round trip), how to turn JSON inputs into ListNode/TreeNode/graph
objects, and how to turn the return value back into JSON. This module is
trusted code: it runs outside the user's file, so the tracer never records
its lines.
"""

import copy
import math

MAX_NODES = 10000


# ---------------------------------------------------------------------------
# JSON -> Python structures
# ---------------------------------------------------------------------------

def build_list(values, env):
    ListNode = env['ListNode']
    dummy = ListNode(0)
    tail = dummy
    for v in values or []:
        tail.next = ListNode(v)
        tail = tail.next
    return dummy.next


def build_cycle_list(values, pos, env):
    head = build_list(values, env)
    if head is None or pos is None or pos < 0:
        return head
    nodes = []
    node = head
    while node:
        nodes.append(node)
        node = node.next
    nodes[-1].next = nodes[pos]
    return head


def build_tree(values, env):
    """Build a binary tree from LeetCode level-order form: [1,null,2,3]."""
    TreeNode = env['TreeNode']
    if not values or values[0] is None:
        return None
    root = TreeNode(values[0])
    queue = [root]
    i = 1
    head = 0
    while head < len(queue) and i < len(values):
        node = queue[head]
        head += 1
        if i < len(values) and values[i] is not None:
            node.left = TreeNode(values[i])
            queue.append(node.left)
        i += 1
        if i < len(values) and values[i] is not None:
            node.right = TreeNode(values[i])
            queue.append(node.right)
        i += 1
    return root


def find_tree_node(root, val):
    stack = [root]
    while stack:
        node = stack.pop()
        if node is None:
            continue
        if node.val == val:
            return node
        stack.append(node.left)
        stack.append(node.right)
    return None


def build_graph(adj, env):
    """Build an undirected graph from a 1-indexed adjacency list."""
    Node = env['Node']
    if not adj:
        return None
    nodes = [Node(i + 1) for i in range(len(adj))]
    for i, nbrs in enumerate(adj):
        nodes[i].neighbors = [nodes[j - 1] for j in nbrs]
    return nodes[0]


def build_random_list(pairs, env):
    """Build a list from [[val, random_index], ...]."""
    Node = env['Node']
    if not pairs:
        return None
    nodes = [Node(p[0]) for p in pairs]
    for i, p in enumerate(pairs):
        nodes[i].next = nodes[i + 1] if i + 1 < len(nodes) else None
        nodes[i].random = nodes[p[1]] if p[1] is not None else None
    return nodes[0]


# ---------------------------------------------------------------------------
# Python structures -> JSON
# ---------------------------------------------------------------------------

def is_list_node(obj):
    return type(obj).__name__ == 'ListNode' and hasattr(obj, 'next') and hasattr(obj, 'val')


def is_tree_node(obj):
    return hasattr(obj, 'left') and hasattr(obj, 'right') and hasattr(obj, 'val')


def list_to_json(head):
    out = []
    seen = set()
    while head is not None:
        if id(head) in seen or len(out) >= MAX_NODES:
            out.append("<cycle>")
            break
        seen.add(id(head))
        out.append(to_json(head.val))
        head = head.next
    return out


def tree_to_json(root):
    if root is None:
        return []
    out = []
    queue = [root]
    head = 0
    while head < len(queue) and len(out) < MAX_NODES:
        node = queue[head]
        head += 1
        if node is None:
            out.append(None)
            continue
        out.append(to_json(node.val))
        queue.append(node.left)
        queue.append(node.right)
    while out and out[-1] is None:
        out.pop()
    return out


def graph_to_json(start):
    if start is None:
        return []
    index = {}
    order = []
    queue = [start]
    while queue:
        node = queue.pop(0)
        if id(node) in index:
            continue
        index[id(node)] = node
        order.append(node)
        for n in node.neighbors:
            if id(n) not in index:
                queue.append(n)
    order.sort(key=lambda n: n.val)
    return [sorted(n.val for n in node.neighbors) for node in order]


def random_list_to_json(head):
    nodes = []
    pos = {}
    node = head
    while node is not None and len(nodes) < MAX_NODES:
        pos[id(node)] = len(nodes)
        nodes.append(node)
        node = node.next
    out = []
    for n in nodes:
        r = getattr(n, 'random', None)
        out.append([to_json(n.val), pos.get(id(r)) if r is not None else None])
    return out


def to_json(value, return_type=""):
    """Convert any returned value into plain JSON data."""
    if return_type == 'GraphNode':
        return graph_to_json(value)
    if return_type == 'RandomList':
        return random_list_to_json(value)
    if return_type == 'TreeNodeVal':
        return None if value is None else to_json(value.val)
    if return_type in ('ListNode', 'TreeNode') and value is None:
        return []  # an empty list or tree
    if value is None or isinstance(value, (bool, int, str)):
        return value
    if isinstance(value, float):
        if math.isnan(value):
            return "NaN"
        if math.isinf(value):
            return "Infinity" if value > 0 else "-Infinity"
        return value
    if isinstance(value, (list, tuple)):
        return [to_json(v) for v in value]
    if isinstance(value, (set, frozenset)):
        items = [to_json(v) for v in value]
        try:
            return sorted(items)
        except TypeError:
            return items
    if isinstance(value, dict):
        return {str(k): to_json(v) for k, v in value.items()}
    if is_list_node(value):
        return list_to_json(value)
    if is_tree_node(value):
        return tree_to_json(value)
    return repr(value)


# ---------------------------------------------------------------------------
# Running a test case
# ---------------------------------------------------------------------------

def convert_args(spec, inputs, env):
    """Turn the JSON inputs into the Python objects the function expects."""
    types = spec.get('paramTypes') or {}
    converted = {}
    for name, value in inputs.items():
        kind = types.get(name, '')
        if kind == 'ListNode':
            converted[name] = build_list(value, env)
        elif kind == 'ListNode[]':
            converted[name] = [build_list(v, env) for v in value]
        elif kind == 'TreeNode':
            converted[name] = build_tree(value, env)
        elif kind == 'GraphNode':
            converted[name] = build_graph(value, env)
        elif kind == 'RandomList':
            converted[name] = build_random_list(value, env)
        elif kind.startswith('CycleList:'):
            converted[name] = build_cycle_list(value, inputs.get(kind.split(':', 1)[1]), env)
        else:
            converted[name] = copy.deepcopy(value)
    # Second pass: values that refer to nodes inside an already-built tree.
    for name, kind in types.items():
        if kind.startswith('TreeNodeRef:'):
            root = converted.get(kind.split(':', 1)[1])
            converted[name] = find_tree_node(root, inputs.get(name))
    return converted


def run_function(spec, args, env):
    fn = env.get(spec['function'])
    if fn is None:
        raise NameError(f"function '{spec['function']}' is not defined - keep the name from the starter code")
    params = spec.get('params') or []
    result = fn(*[args[p] for p in params])
    if spec.get('inPlace'):
        return args[spec['inPlace']]
    return result


def run_design(args, env):
    ops = args.get('operations') or []
    values = args.get('values')
    if values is None:
        values = args.get('args') or []
    if not ops:
        raise ValueError("design test case has no operations")
    cls = env.get(ops[0])
    if cls is None:
        raise NameError(f"class '{ops[0]}' is not defined - keep the name from the starter code")
    obj = cls(*(values[0] if values else []))
    out: list = [None]
    for op, vals in zip(ops[1:], values[1:]):
        method = getattr(obj, op, None)
        if method is None:
            raise AttributeError(f"'{ops[0]}' has no method '{op}'")
        out.append(to_json(method(*vals)))
    return out


def run_roundtrip(spec, args, env):
    enc_name, dec_name = [s.strip() for s in spec['function'].split(',')]
    owner = env
    if spec.get('className'):
        cls = env.get(spec['className'])
        if cls is None:
            raise NameError(f"class '{spec['className']}' is not defined - keep the name from the starter code")
        obj = cls()
        enc, dec = getattr(obj, enc_name), getattr(obj, dec_name)
    else:
        enc, dec = owner.get(enc_name), owner.get(dec_name)
        if enc is None or dec is None:
            raise NameError(f"define both '{enc_name}' and '{dec_name}'")
    params = spec.get('params') or list(args.keys())
    encoded = enc(args[params[0]])
    if not isinstance(encoded, str):
        raise TypeError(f"{enc_name} must return a string, got {type(encoded).__name__}")
    return dec(encoded)


def run_case(spec, env, safe_builtins):
    """Execute one test case. Returns (result_json, validator_verdict)."""
    inputs = spec.get('inputs') or {}
    raw = copy.deepcopy(inputs)
    args = convert_args(spec, inputs, env)
    mode = spec.get('mode') or 'function'

    if mode == 'design':
        result = run_design(args, env)
        result_json = result
    elif mode == 'roundtrip':
        result = run_roundtrip(spec, args, env)
        result_json = to_json(result, spec.get('returnType', ''))
    else:
        result = run_function(spec, args, env)
        result_json = to_json(result, spec.get('returnType', ''))

    verdict = None
    if spec.get('validator'):
        verdict = run_validator(spec, raw, args, result, result_json, safe_builtins)
    return result_json, verdict


def run_validator(spec, raw, args, result, result_json, safe_builtins):
    scope = {
        '__builtins__': safe_builtins,
        'list_to_json': list_to_json,
        'tree_to_json': tree_to_json,
        'to_json': to_json,
        'result_obj': result,
    }
    exec(compile(spec['validator'], '<validator>', 'exec'), scope)
    check = scope.get('check')
    if check is None:
        return {"passed": False, "message": "validator has no check() function"}
    try:
        out = check(raw, args, result_json, spec.get('expected'))
    except Exception as e:  # a crash in the checker means the answer has the wrong shape
        return {"passed": False, "message": f"answer has the wrong shape: {type(e).__name__}: {e}"}
    if isinstance(out, tuple):
        return {"passed": bool(out[0]), "message": str(out[1]) if len(out) > 1 else ""}
    return {"passed": bool(out), "message": ""}
