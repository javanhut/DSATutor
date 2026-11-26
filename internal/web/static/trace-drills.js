// Trace by hand: drills that build intuition for an algorithm with no code.
//
// The learner gets a freshly generated structure (a tree, a graph or a linked
// list) and clicks nodes: in the order an algorithm visits them, or where its
// pointers end up. A wrong answer says which step went wrong and why, and
// "Show me" replays the algorithm one step at a time with its call stack,
// queue or pointers and the code line it is on. Results live in
// practiceState.progress under "trace:<cat>:<id>", so they are exported,
// imported and reset with the rest of practice.
//
// The core below runs any "kit". A kit lists exercises and builds questions:
//   { prompt, aside?, mode: 'sequence' | 'choice', ids, label(id), svg(view),
//     expected, fixedLength, slots?, validate?(answer) -> {ok, k},
//     wrong(k, answer), right(), code, steps(), outLabel? }
// A choice question has { sequence, choices: [{id, label}], answerId,
//   explain(picked) } instead.
(function () {
  const STREAK_TO_MASTER = 3;
  const KITS = {};
  const S = { cat: 'trees', exId: null, size: 'small' };
  const $ = (id) => document.getElementById(id);
  const esc = (s) => (typeof escapeHtml === 'function' ? escapeHtml(String(s)) : String(s));
  const cap = (s) => s[0].toUpperCase() + s.slice(1);

  // -------------------------------------------------------------------------
  // Shared helpers

  function randInt(a, b) { return a + Math.floor(Math.random() * (b - a + 1)); }
  function pick(list) { return list[randInt(0, list.length - 1)]; }
  function shuffle(a) {
    for (let i = a.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [a[i], a[j]] = [a[j], a[i]];
    }
    return a;
  }
  function distinctValues(n, max = 99) {
    return shuffle(Array.from({ length: max }, (_, i) => i + 1)).slice(0, n);
  }
  function firstMismatch(answer, expected) {
    let k = 0;
    while (k < answer.length && k < expected.length && answer[k] === expected[k]) k++;
    return k;
  }

  // nodeSvg draws one round node with the core's state classes.
  function nodeSvg(id, x, y, label, view, extra = '') {
    const cls = ['tnode'];
    const idx = view.picked.get(id);
    if (idx) cls.push('picked');
    if (view.current === id) cls.push('current');
    if (view.held.has(id) && view.current !== id) cls.push('held');
    if (view.wrong === id) cls.push('wrong');
    if (view.target === id) cls.push('target');
    const clickable = view.clickable && (!idx || view.repeats);
    if (clickable) cls.push('clickable');
    return `<g class="${cls.join(' ')}" data-id="${esc(id)}" ${clickable ? `role="button" tabindex="0" aria-label="Node ${esc(label)}"` : ''}>
      <circle cx="${x}" cy="${y}" r="17" />
      <text x="${x}" y="${y}" dy="0.35em">${esc(label)}</text>
      ${idx ? `<text class="badge" x="${x + 20}" y="${y - 15}">${idx}</text>` : ''}
      ${extra}
    </g>`;
  }

  function svgWrap(width, height, body, label) {
    return `<svg class="trace-tree" viewBox="0 0 ${width} ${height}" style="width:100%;max-width:${Math.round(width * 1.3)}px" role="img" aria-label="${esc(label)}">
      <defs><marker id="trace-arrow" viewBox="0 0 10 10" refX="9" refY="5" markerUnits="userSpaceOnUse" markerWidth="14" markerHeight="14" orient="auto-start-reverse"><path d="M0,0 L10,5 L0,10 z" class="arrowhead" /></marker></defs>
      ${body}</svg>`;
  }

  // ===========================================================================
  // Trees
  // ===========================================================================

  const ORDER_NAMES = { pre: 'preorder', in: 'inorder', post: 'postorder', level: 'level order' };
  const TREE_SIZES = { small: { n: 6, levels: 3 }, medium: { n: 9, levels: 4 }, large: { n: 13, levels: 5 } };

  // A random shape of n nodes with at most `levels` levels: the left subtree
  // takes any size that still lets both sides fit.
  function shape(n, levels, depth = 0, parent = null, side = null) {
    if (n === 0) return null;
    const room = 2 ** (levels - 1) - 1;
    const k = randInt(Math.max(0, n - 1 - room), Math.min(n - 1, room));
    const node = { val: 0, depth, parent, side, left: null, right: null };
    node.left = shape(k, levels - 1, depth + 1, node, 'left');
    node.right = shape(n - 1 - k, levels - 1, depth + 1, node, 'right');
    return node;
  }

  function nodesInorder(root, out = []) {
    if (root) { nodesInorder(root.left, out); out.push(root); nodesInorder(root.right, out); }
    return out;
  }

  // makeTree labels a random shape. A BST gets sorted values in inorder
  // position; any other tree gets shuffled values, so inorder gives nothing away.
  function makeTree(size, bst) {
    const { n, levels } = TREE_SIZES[size];
    const root = shape(n, levels);
    const nodes = nodesInorder(root);
    const vals = distinctValues(n);
    if (bst) vals.sort((a, b) => a - b);
    nodes.forEach((node, i) => { node.val = vals[i]; node.x = i; node.id = i; });
    return { root, nodes, levels: Math.max(...nodes.map(nd => nd.depth)) + 1 };
  }

  function traverse(root, order) {
    const out = [];
    if (order === 'level') {
      const queue = root ? [root] : [];
      while (queue.length) {
        const node = queue.shift();
        out.push(node);
        if (node.left) queue.push(node.left);
        if (node.right) queue.push(node.right);
      }
      return out;
    }
    (function visit(node) {
      if (!node) return;
      if (order === 'pre') out.push(node);
      visit(node.left);
      if (order === 'in') out.push(node);
      visit(node.right);
      if (order === 'post') out.push(node);
    })(root);
    return out;
  }

  // bstPath is the nodes a search for t compares against. found tells whether
  // it stopped on t; otherwise t would be inserted under the last node.
  function bstPath(root, t) {
    const path = [];
    let node = root;
    while (node) {
      path.push(node);
      if (t === node.val) return { path, found: true };
      node = t < node.val ? node.left : node.right;
    }
    const last = path[path.length - 1];
    return { path, found: false, side: t < last.val ? 'left' : 'right' };
  }

  const vals = (list) => list.map(n => n.val).join(', ');

  function treeReason(order, node, index, seq) {
    const v = node.val;
    const p = node.parent;
    if (order === 'pre') {
      if (!p) return `Preorder records a node the moment it arrives, before either child, so the root ${v} comes first.`;
      if (node.side === 'left') return `${v} is ${p.val}'s left child. Preorder records ${p.val}, then goes straight down its left side.`;
      return `Everything on ${p.val}'s left side is recorded${p.left ? '' : ' (it has none)'}, so preorder moves to its right child ${v} and records it on arrival.`;
    }
    if (order === 'in') {
      if (index === 0) return `Inorder starts at the leftmost node: from the root, keep going left until there is no left child. That is ${v}.`;
      if (node.left) return `Everything in ${v}'s left subtree is recorded, so ${v} comes next, before anything in its right subtree.`;
      return `${v} has no left child, so inorder records it as soon as it gets there (it is the leftmost node of ${p && node.side === 'right' ? `${p.val}'s right subtree` : 'what is left'}).`;
    }
    if (order === 'post') {
      if (!node.left && !node.right) return `${v} is a leaf: both of its subtrees are empty, so it is recorded the moment it is reached${index === 0 ? ', and it is the first leaf on the leftmost path' : ''}.`;
      return `Both of ${v}'s subtrees are fully recorded, so ${v} is recorded on the way back up${p ? '' : '. The root always comes last'}.`;
    }
    if (!p) return 'Level order starts at the root.';
    const prev = seq[index - 1];
    if (prev && prev.depth < node.depth) return `Level ${node.depth - 1} is finished, so level ${node.depth} starts, from the left. ${v} was queued when its parent ${p.val} was recorded.`;
    return `${v} is the next node on level ${node.depth}, going left to right. It was queued when its parent ${p.val} was recorded, behind everything queued before it.`;
  }

  const TREE_CODE = {
    pre: ['def visit(node):', '    if node is None:', '        return', '    out.append(node.val)', '    visit(node.left)', '    visit(node.right)'],
    in: ['def visit(node):', '    if node is None:', '        return', '    visit(node.left)', '    out.append(node.val)', '    visit(node.right)'],
    post: ['def visit(node):', '    if node is None:', '        return', '    visit(node.left)', '    visit(node.right)', '    out.append(node.val)'],
    level: ['queue = deque([root])', 'while queue:', '    node = queue.popleft()', '    out.append(node.val)', '    if node.left: queue.append(node.left)', '    if node.right: queue.append(node.right)'],
    bst: ['node = root', 'while node:', '    if t == node.val:', '        return node          # search: found', '    parent = node', '    node = node.left if t < node.val else node.right', '# fell off the tree: insert t under parent'],
  };

  function treeDfsSteps(tree, order) {
    const steps = [];
    const stack = [];
    const out = [];
    const lines = TREE_CODE[order];
    const recLine = lines.indexOf('    out.append(node.val)');
    const leftLine = lines.indexOf('    visit(node.left)');
    const rightLine = lines.indexOf('    visit(node.right)');
    const seq = traverse(tree.root, order);
    const snap = (node, line, text) => steps.push({ current: node.id, line, text, held: stack.map(n => n.val), heldIds: stack.map(n => n.id), heldLabel: 'Call stack (bottom to top)', out: out.map(n => n.id) });
    (function visit(node) {
      stack.push(node);
      snap(node, 0, node.parent ? `Go ${node.side} from ${node.parent.val}: call visit(${node.val}).` : `Start: call visit(${node.val}) on the root.`);
      const rec = () => {
        out.push(node);
        snap(node, recLine, `Record ${node.val}. ${treeReason(order, node, out.length - 1, seq)}`);
      };
      if (order === 'pre') rec();
      if (node.left) visit(node.left);
      else snap(node, leftLine, `${node.val} has no left child: visit(None) returns at once.`);
      if (order === 'in') rec();
      if (node.right) visit(node.right);
      else snap(node, rightLine, `${node.val} has no right child: visit(None) returns at once.`);
      if (order === 'post') rec();
      stack.pop();
      snap(node, -1, stack.length
        ? `visit(${node.val}) is finished. Return to ${stack[stack.length - 1].val}, which carries on from where it called.`
        : `visit(${node.val}) is finished, and so is the whole tree.`);
    })(tree.root);
    return steps;
  }

  function treeBfsSteps(tree) {
    const steps = [];
    const out = [];
    const queue = [tree.root];
    const seq = traverse(tree.root, 'level');
    const snap = (node, line, text) => steps.push({ current: node?.id, line, text, held: queue.map(n => n.val), heldIds: queue.map(n => n.id), heldLabel: 'Queue (front to back)', out: out.map(n => n.id) });
    snap(tree.root, 0, `Start: put the root ${tree.root.val} in the queue.`);
    while (queue.length) {
      const node = queue.shift();
      out.push(node);
      snap(node, 3, `Take ${node.val} from the front of the queue and record it. ${treeReason('level', node, out.length - 1, seq)}`);
      if (node.left) queue.push(node.left);
      snap(node, 4, node.left ? `Add ${node.val}'s left child ${node.left.val} to the back of the queue.` : `${node.val} has no left child.`);
      if (node.right) queue.push(node.right);
      snap(node, 5, node.right ? `Add ${node.val}'s right child ${node.right.val} to the back of the queue.` : `${node.val} has no right child.`);
    }
    snap(null, 1, 'The queue is empty, so every node has been recorded.');
    return steps;
  }

  function treeBstSteps(tree, t) {
    const { path, found, side } = bstPath(tree.root, t);
    const steps = [];
    const out = [];
    path.forEach((node, i) => {
      out.push(node.id);
      if (node.val === t) {
        steps.push({ current: node.id, line: 3, text: `${t} == ${node.val}: found it.`, out: out.slice() });
        return;
      }
      const dir = t < node.val ? 'left' : 'right';
      const next = node[dir];
      steps.push({
        current: node.id, line: 5, out: out.slice(),
        text: `${i === 0 ? 'Start at the root. ' : ''}${t} ${t < node.val ? '<' : '>'} ${node.val}, so go ${dir}${next ? ` to ${next.val}` : `, but ${node.val} has no ${dir} child`}. Everything on the ${dir === 'left' ? 'right' : 'left'} of ${node.val} can be ignored.`,
      });
    });
    if (!found) {
      const last = path[path.length - 1];
      steps.push({ current: last.id, line: 6, out: out.slice(), ghost: true, text: `The search fell off the tree, so ${t} is not there. Inserting it makes it ${last.val}'s new ${side} child.` });
    }
    return steps;
  }

  function treeSvg(tree, view, ghostInfo) {
    const dx = 46;
    const dy = 66;
    const pad = 31;
    const ghost = ghostInfo && (view.step?.ghost || view.done);
    const xs = tree.nodes.map(n => n.x);
    const ghostX = ghost ? ghostInfo.parent.x + (ghostInfo.side === 'left' ? -0.5 : 0.5) : null;
    const minX = Math.min(...xs, ghostX ?? Infinity);
    const maxX = Math.max(...xs, ghostX ?? -Infinity);
    const px = (x) => (x - minX) * dx + pad;
    const py = (d) => d * dy + pad;
    let edges = '';
    let nodes = '';
    tree.nodes.forEach(n => {
      [n.left, n.right].forEach(c => {
        if (c) edges += `<line x1="${px(n.x)}" y1="${py(n.depth)}" x2="${px(c.x)}" y2="${py(c.depth)}" />`;
      });
      nodes += nodeSvg(n.id, px(n.x), py(n.depth), n.val, view);
    });
    if (ghost) {
      const last = ghostInfo.parent;
      const gx = px(ghostX);
      const gy = py(last.depth + 1);
      edges += `<line class="ghost" x1="${px(last.x)}" y1="${py(last.depth)}" x2="${gx}" y2="${gy}" />`;
      nodes += `<g class="tnode ghost"><circle cx="${gx}" cy="${gy}" r="17" /><text x="${gx}" y="${gy}" dy="0.35em">${ghostInfo.val}</text></g>`;
    }
    const width = (maxX - minX) * dx + 2 * pad;
    const height = (tree.levels + (ghost ? 1 : 0) - 1) * dy + 2 * pad;
    return svgWrap(width, height, `<g class="edges">${edges}</g>${nodes}`, 'Binary tree');
  }

  function treeQuestion(exId, size) {
    if (exId === 'name') return treeNameQuestion(size);
    if (exId === 'bst') return treeBstQuestion(size);
    const order = { preorder: 'pre', inorder: 'in', postorder: 'post', levelorder: 'level' }[exId];
    const tree = makeTree(size, false);
    const seq = traverse(tree.root, order);
    const byId = (id) => tree.nodes[id];
    return {
      prompt: `Click the nodes in <strong>${ORDER_NAMES[order]}</strong>, one at a time.`,
      mode: 'sequence',
      ids: tree.nodes.map(n => n.id),
      label: (id) => byId(id).val,
      svg: (view) => treeSvg(tree, view),
      expected: seq.map(n => n.id),
      fixedLength: true,
      code: TREE_CODE[order],
      steps: () => (order === 'level' ? treeBfsSteps(tree) : treeDfsSteps(tree, order)),
      right: () => `${cap(ORDER_NAMES[order])}: ${vals(seq)}.`,
      wrong: (k, answer) => {
        const right = k === 0 ? 'The first node is wrong.' : `The first ${k} ${k === 1 ? 'is' : 'are'} right.`;
        const given = answer.map(byId).map(n => n.val).join();
        const other = ['pre', 'in', 'post', 'level'].find(o => o !== order && traverse(tree.root, o).map(n => n.val).join() === given);
        return `<p class="trace-bad">${right} Next comes ${seq[k].val}, not ${byId(answer[k]).val}.${other ? ` (Your order is ${ORDER_NAMES[other]}, not ${ORDER_NAMES[order]}.)` : ''}</p>
          <p class="trace-why">${treeReason(order, seq[k], k, seq)}</p>`;
      },
    };
  }

  function treeBstQuestion(size) {
    const tree = makeTree(size, true);
    const present = tree.nodes.map(n => n.val);
    const insert = Math.random() < 0.5;
    const target = insert
      ? pick(Array.from({ length: 99 }, (_, i) => i + 1).filter(v => !present.includes(v)))
      : pick(tree.nodes.filter(n => n.depth >= 1)).val; // below the root, so the path has some length
    const { path, found, side } = bstPath(tree.root, target);
    const last = path[path.length - 1];
    const byId = (id) => tree.nodes[id];
    const targetNode = tree.nodes.find(n => n.val === target);
    return {
      prompt: insert
        ? `<strong>Insert ${target}</strong> into this BST. Click every node ${target} is compared against, in order. The last one becomes its parent.`
        : `<strong>Search for ${target}</strong> in this BST. Click every node the search compares ${target} against, in order, ending on ${target}.`,
      mode: 'sequence',
      ids: tree.nodes.map(n => n.id),
      label: (id) => byId(id).val,
      svg: (view) => treeSvg(tree, { ...view, target: view.checked && targetNode ? targetNode.id : undefined }, found ? null : { parent: last, side, val: target }),
      expected: path.map(n => n.id),
      fixedLength: false,
      code: TREE_CODE.bst,
      steps: () => treeBstSteps(tree, target),
      right: () => (found
        ? `The search compares ${target} against ${vals(path)} and finds it.`
        : `The search compares ${target} against ${vals(path)}, then falls off the tree: ${target} becomes ${last.val}'s ${side} child.`),
      wrong: (k, answer) => {
        const right = k === 0 ? 'The first node is wrong.' : `The first ${k} ${k === 1 ? 'is' : 'are'} right.`;
        if (k === path.length) {
          return `<p class="trace-bad">${right} But the search stops at ${last.val}: ${found ? `that is ${target} itself.` : `${target} ${target < last.val ? '<' : '>'} ${last.val} and ${last.val} has no ${side} child, so there is nowhere further to go.`}</p>`;
        }
        const next = path[k];
        const why = k === 0
          ? `Every search starts at the root, ${next.val}.`
          : `${target} ${target < path[k - 1].val ? '<' : '>'} ${path[k - 1].val}, so the search goes ${target < path[k - 1].val ? 'left' : 'right'} to ${next.val}. It never visits the other side.`;
        return `<p class="trace-bad">${right} ${k < answer.length ? `Next comes ${next.val}, not ${byId(answer[k]).val}.` : `Keep going: next comes ${next.val}.`} ${why}</p>`;
      },
    };
  }

  function treeNameQuestion(size) {
    let tree;
    let all;
    // Regenerate until the four orders all differ, so exactly one fits.
    for (let tries = 0; tries < 50; tries++) {
      tree = makeTree(size, false);
      all = { pre: traverse(tree.root, 'pre'), in: traverse(tree.root, 'in'), post: traverse(tree.root, 'post'), level: traverse(tree.root, 'level') };
      if (new Set(Object.values(all).map(vals)).size === 4) break;
    }
    const answerId = pick(Object.keys(all));
    const r = tree.root;
    const clue = {
      pre: `the root ${r.val} comes first and the next value is the root's ${r.left ? `left child ${r.left.val}` : `right child ${r.right?.val}`}: each node, then straight down its left side`,
      in: `the root ${r.val} sits between its left and right subtrees, and the sequence starts at the leftmost node ${all.in[0].val}`,
      post: `the root ${r.val} comes last, and the sequence starts with a leaf, ${all.post[0].val}`,
      level: `the root ${r.val} comes first, then its children side by side (${vals(all.level.slice(1, 3))}), a level at a time`,
    }[answerId];
    return {
      prompt: 'Which traversal of this tree produces this sequence?',
      mode: 'choice',
      svg: (view) => treeSvg(tree, view),
      sequence: all[answerId].map(n => n.val),
      choices: ['pre', 'in', 'post', 'level'].map(o => ({ id: o, label: ORDER_NAMES[o] })),
      answerId,
      explain: (picked) => `<p class="${picked === answerId ? 'trace-ok' : 'trace-bad'}"><strong>${picked === answerId ? 'Right' : 'Not quite'}:</strong> it is ${ORDER_NAMES[answerId]}. The giveaway: ${clue}.</p>
        <table class="trace-orders">${['pre', 'in', 'post', 'level'].map(o => `<tr class="${o === answerId ? 'is-answer' : ''}"><th>${ORDER_NAMES[o]}</th><td>${vals(all[o])}</td></tr>`).join('')}</table>`,
    };
  }

  KITS.trees = {
    title: 'Trees',
    exercises: [
      { id: 'preorder', name: 'Preorder' },
      { id: 'inorder', name: 'Inorder' },
      { id: 'postorder', name: 'Postorder' },
      { id: 'levelorder', name: 'Level order' },
      { id: 'bst', name: 'BST search & insert' },
      { id: 'name', name: 'Name the order' },
    ],
    make: treeQuestion,
  };

  // ===========================================================================
  // Graphs
  // ===========================================================================

  const GRAPH_SIZES = { small: 6, medium: 8, large: 10 };

  // makeGraph places n nodes on a jittered grid (filled row by row, so the
  // grid's own edges connect them), joins them with a random spanning tree
  // of neighbouring cells plus some extra short edges, and shuffles the
  // labels so a label says nothing about position. Crossing diagonals are
  // skipped to keep the drawing readable. directed orients every edge along
  // a random ranking, which can't make a cycle.
  function makeGraph(size, { directed = false, weighted = false } = {}) {
    const n = GRAPH_SIZES[size];
    const cols = n <= 6 ? 3 : n <= 8 ? 4 : 4;
    const labels = shuffle(Array.from({ length: n }, (_, i) => i));
    const cells = Array.from({ length: n }, (_, i) => ({ r: Math.floor(i / cols), c: i % cols }));
    const pos = cells.map(({ r, c }) => ({ x: 50 + c * 120 + randInt(-16, 16), y: 45 + r * 105 + randInt(-14, 14) }));
    const cand = [];
    for (let i = 0; i < n; i++) {
      for (let j = i + 1; j < n; j++) {
        const dr = Math.abs(cells[i].r - cells[j].r);
        const dc = Math.abs(cells[i].c - cells[j].c);
        if (dr <= 1 && dc <= 1) cand.push([i, j, dr + dc === 2]);
      }
    }
    shuffle(cand);
    const parent = Array.from({ length: n }, (_, i) => i);
    const find = (x) => (parent[x] === x ? x : (parent[x] = find(parent[x])));
    const chosen = [];
    const has = (i, j) => chosen.some(([a, b]) => (a === i && b === j) || (a === j && b === i));
    const crosses = (i, j) => {
      // A diagonal crosses the other diagonal of the same 2x2 square.
      const a = cells[i];
      const b = cells[j];
      const top = Math.min(a.r, b.r);
      const left = Math.min(a.c, b.c);
      const idx = (r, c) => cells.findIndex(x => x.r === r && x.c === c);
      const other = (a.r === top) === (a.c === left)
        ? [idx(top, left + 1), idx(top + 1, left)]
        : [idx(top, left), idx(top + 1, left + 1)];
      return other[0] >= 0 && other[1] >= 0 && has(other[0], other[1]);
    };
    // Spanning tree first, preferring straight edges so diagonals stay rare.
    [...cand.filter(e => !e[2]), ...cand.filter(e => e[2])].forEach(([i, j, diag]) => {
      if (find(i) !== find(j) && !(diag && crosses(i, j))) { parent[find(i)] = find(j); chosen.push([i, j]); }
    });
    cand.forEach(([i, j, diag]) => {
      if (!has(i, j) && Math.random() < (diag ? 0.18 : 0.35) && !(diag && crosses(i, j))) chosen.push([i, j]);
    });
    const rank = shuffle(Array.from({ length: n }, (_, i) => i));
    const nodes = Array.from({ length: n }, (_, i) => ({ id: labels[i], ...pos[i] }));
    const byId = new Map(nodes.map(nd => [nd.id, nd]));
    const edges = chosen.map(([i, j]) => {
      let [u, v] = [labels[i], labels[j]];
      if (directed && rank[i] > rank[j]) [u, v] = [v, u];
      return { u, v, w: weighted ? randInt(1, 9) : 1 };
    });
    const adj = new Map(nodes.map(nd => [nd.id, []]));
    edges.forEach(({ u, v, w }) => {
      adj.get(u).push({ v, w });
      if (!directed) adj.get(v).push({ v: u, w });
    });
    adj.forEach(list => list.sort((a, b) => a.v - b.v));
    return { n, nodes, byId, edges, adj, directed, weighted };
  }

  // graphSvg draws g. A step's tree (a Set of edge ids) is drawn bold. With
  // edgeIds, each weight badge is clickable and carries its edge's id, and
  // the nodes are not clickable.
  function graphSvg(g, view, { edgeIds = false } = {}) {
    const r = 17;
    let edges = '';
    let weights = '';
    const tree = view.step?.tree || (edgeIds ? new Set(view.picked.keys()) : null);
    g.edges.forEach(({ u, v, w, id }) => {
      const a = g.byId.get(u);
      const b = g.byId.get(v);
      const dx = b.x - a.x;
      const dy = b.y - a.y;
      const len = Math.hypot(dx, dy);
      const ux = dx / len;
      const uy = dy / len;
      const x1 = a.x + ux * r;
      const y1 = a.y + uy * r;
      const x2 = b.x - ux * (r + (g.directed ? 3 : 0));
      const y2 = b.y - uy * (r + (g.directed ? 3 : 0));
      const hot = view.step?.edge && ((view.step.edge[0] === u && view.step.edge[1] === v) || (!g.directed && view.step.edge[0] === v && view.step.edge[1] === u));
      const cls = [hot ? 'hot' : '', tree?.has(id) ? 'tree' : ''].filter(Boolean).join(' ');
      edges += `<line class="${cls}" x1="${x1}" y1="${y1}" x2="${x2}" y2="${y2}" ${g.directed ? 'marker-end="url(#trace-arrow)"' : ''} />`;
      if (g.weighted) {
        const mx = (a.x + b.x) / 2 - uy * (edgeIds ? 0 : 9);
        const my = (a.y + b.y) / 2 + ux * (edgeIds ? 0 : 9);
        weights += edgeIds ? edgeBadge(id, mx, my, w, `${u}–${v}`, view) : `<g class="weight"><circle cx="${mx}" cy="${my}" r="9" /><text x="${mx}" y="${my}" dy="0.35em">${w}</text></g>`;
      }
    });
    const dist = view.step?.dist;
    const nodes = g.nodes.map(nd => {
      const extra = dist && dist[nd.id] !== undefined
        ? `<text class="dist" x="${nd.x}" y="${nd.y + 31}">${dist[nd.id] === Infinity ? '∞' : dist[nd.id]}</text>` : '';
      return nodeSvg(nd.id, nd.x, nd.y, nd.id, edgeIds ? { ...view, clickable: false, picked: new Map(), wrong: undefined } : view, extra);
    }).join('');
    const width = Math.max(...g.nodes.map(nd => nd.x)) + 50;
    const height = Math.max(...g.nodes.map(nd => nd.y)) + 50;
    return svgWrap(width, height, `<g class="edges">${edges}</g>${weights}${nodes}`, g.directed ? 'Directed graph' : 'Graph');
  }

  // edgeBadge is a clickable weight on an edge, styled like a small node.
  function edgeBadge(id, x, y, w, name, view) {
    const cls = ['tnode', 'wedge'];
    const idx = view.picked.get(id);
    if (idx) cls.push('picked');
    if (view.wrong === id) cls.push('wrong');
    const clickable = view.clickable && !idx;
    if (clickable) cls.push('clickable');
    return `<g class="${cls.join(' ')}" data-id="${esc(id)}" ${clickable ? `role="button" tabindex="0" aria-label="Edge ${esc(name)}, weight ${w}"` : ''}>
      <circle cx="${x}" cy="${y}" r="12" style="fill: var(--panel); stroke: none" />
      <circle cx="${x}" cy="${y}" r="12" />
      <text x="${x}" y="${y}" dy="0.35em">${w}</text>
      ${idx ? `<text class="badge" x="${x + 13}" y="${y - 10}">${idx}</text>` : ''}
    </g>`;
  }

  function adjacencyHtml(g) {
    const rows = [...g.adj.keys()].sort((a, b) => a - b).map(u => {
      const list = g.adj.get(u).map(({ v, w }) => (g.weighted ? `${v} (${w})` : v)).join(', ');
      return `<tr><th>${u}</th><td>${list || '<em>none</em>'}</td></tr>`;
    }).join('');
    return `<details class="trace-adj" open><summary>Adjacency list${g.weighted ? ' (neighbour (weight))' : ''}${g.directed ? ', edges point one way' : ''}</summary><table>${rows}</table></details>`;
  }

  const GRAPH_CODE = {
    bfs: ['visited = {start}', 'queue = deque([start])', 'while queue:', '    u = queue.popleft()', '    out.append(u)', '    for v in sorted(adj[u]):', '        if v not in visited:', '            visited.add(v)', '            queue.append(v)'],
    dfs: ['def dfs(u):', '    visited.add(u)', '    out.append(u)', '    for v in sorted(adj[u]):', '        if v not in visited:', '            dfs(v)'],
    topo: ['indegree = count incoming edges of each node', 'ready = [u for u in nodes if indegree[u] == 0]', 'while ready:', '    u = smallest node in ready   # any ready node works', '    out.append(u)', '    for v in adj[u]:', '        indegree[v] -= 1', '        if indegree[v] == 0: ready.append(v)'],
    dijkstra: ['dist = {start: 0}, all others infinity', 'heap = [(0, start)]', 'while heap:', '    d, u = heappop(heap)', '    if u is settled: continue', '    settle u                     # its distance is final', '    for v, w in adj[u]:', '        if d + w < dist[v]:', '            dist[v] = d + w', '            heappush(heap, (dist[v], v))'],
  };

  function bfsRun(g, start) {
    const steps = [];
    const out = [];
    const visited = new Set([start]);
    const queue = [start];
    const parent = new Map();
    const snap = (current, line, text, edge) => steps.push({ current, line, text, edge, held: queue.slice(), heldIds: queue.slice(), heldLabel: 'Queue (front to back)', out: out.slice() });
    snap(start, 1, `Start at ${start}: mark it visited and put it in the queue.`);
    while (queue.length) {
      const u = queue.shift();
      out.push(u);
      snap(u, 4, `Take ${u} from the front of the queue and record it.${parent.has(u) ? ` It was queued when ${parent.get(u)} was recorded.` : ''}`);
      g.adj.get(u).forEach(({ v }) => {
        if (visited.has(v)) {
          snap(u, 6, `${v} is already visited (it is queued or recorded), so skip it.`, [u, v]);
        } else {
          visited.add(v);
          parent.set(v, u);
          queue.push(v);
          snap(u, 8, `${v} is new: mark it visited now and add it to the back of the queue.`, [u, v]);
        }
      });
    }
    snap(null, 2, 'The queue is empty: every node reachable from the start has been recorded.');
    return { steps, order: out, parent };
  }

  function dfsRun(g, start) {
    const steps = [];
    const out = [];
    const visited = new Set();
    const stack = [];
    const parent = new Map();
    const snap = (current, line, text, edge) => steps.push({ current, line, text, edge, held: stack.slice(), heldIds: stack.slice(), heldLabel: 'Call stack (bottom to top)', out: out.slice() });
    (function dfs(u) {
      stack.push(u);
      visited.add(u);
      out.push(u);
      snap(u, 2, parent.has(u) ? `Call dfs(${u}) from ${parent.get(u)}: mark ${u} visited and record it.` : `Start: dfs(${u}) marks ${u} visited and records it.`);
      g.adj.get(u).forEach(({ v }) => {
        if (visited.has(v)) {
          snap(u, 4, `${u}'s neighbour ${v} is already visited, so skip it.`, [u, v]);
        } else {
          parent.set(v, u);
          snap(u, 5, `${v} is ${u}'s smallest unvisited neighbour: go deeper, before looking at ${u}'s other neighbours.`, [u, v]);
          dfs(v);
        }
      });
      stack.pop();
      snap(u, -1, stack.length
        ? `Every neighbour of ${u} is visited, so dfs(${u}) returns and ${stack[stack.length - 1]} carries on with its next neighbour.`
        : `Every neighbour of ${u} is visited and the stack is empty: done.`);
    })(start);
    return { steps, order: out, parent };
  }

  function topoRun(g) {
    const indeg = new Map(g.nodes.map(nd => [nd.id, 0]));
    g.edges.forEach(({ v }) => indeg.set(v, indeg.get(v) + 1));
    const ready = [...indeg].filter(([, d]) => d === 0).map(([u]) => u).sort((a, b) => a - b);
    const out = [];
    const steps = [];
    const snap = (current, line, text, edge) => steps.push({ current, line, text, edge, held: ready.slice(), heldIds: ready.slice(), heldLabel: 'Ready (in-degree 0)', out: out.slice(), indeg: new Map(indeg) });
    snap(null, 1, `Count each node's incoming edges. Nodes with none can go first: ${ready.join(', ')}.`);
    while (ready.length) {
      ready.sort((a, b) => a - b);
      const u = ready.shift();
      out.push(u);
      snap(u, 4, `Take ${u} from the ready set and record it. Nothing it depends on is left.`);
      g.adj.get(u).forEach(({ v }) => {
        indeg.set(v, indeg.get(v) - 1);
        if (indeg.get(v) === 0) ready.push(v);
        snap(u, indeg.get(v) === 0 ? 7 : 6, indeg.get(v) === 0
          ? `Remove the edge ${u} → ${v}. ${v} has no incoming edges left, so it is ready.`
          : `Remove the edge ${u} → ${v}. ${v} still waits for ${indeg.get(v)} more.`, [u, v]);
      });
    }
    snap(null, 2, 'Nothing is ready and every node is recorded: that is a topological order.');
    return { steps, order: out };
  }

  function dijkstraRun(g, start) {
    const dist = new Map(g.nodes.map(nd => [nd.id, Infinity]));
    const via = new Map();
    dist.set(start, 0);
    const heap = [[0, start]];
    const settled = [];
    const done = new Set();
    const steps = [];
    const heapSort = () => heap.sort((a, b) => a[0] - b[0] || a[1] - b[1]);
    const snap = (current, line, text, edge) => steps.push({
      current, line, text, edge, out: settled.slice(),
      held: heap.map(([d, v]) => `${v}:${d}`), heldIds: heap.map(([, v]) => v), heldLabel: 'Heap, smallest first (node:distance)',
      dist: Object.fromEntries(dist),
    });
    snap(start, 1, `Start: ${start} is at distance 0 and everything else at infinity.`);
    while (heap.length) {
      heapSort();
      const [d, u] = heap.shift();
      if (done.has(u)) {
        snap(u, 4, `Pop ${u} at ${d}: it is already settled with a shorter distance, so this entry is stale. Skip it.`);
        continue;
      }
      done.add(u);
      settled.push(u);
      snap(u, 5, `Pop ${u} at distance ${d}, the smallest in the heap. No other route can be shorter, because every other node is at least ${d} away and weights are never negative: settle ${u}.${via.has(u) ? ` (Its best route arrives from ${via.get(u)}.)` : ''}`);
      g.adj.get(u).forEach(({ v, w }) => {
        if (done.has(v)) return;
        if (d + w < dist.get(v)) {
          const old = dist.get(v);
          dist.set(v, d + w);
          via.set(v, u);
          heap.push([d + w, v]);
          heapSort();
          snap(u, 8, `${u} → ${v} costs ${d} + ${w} = ${d + w}, better than ${old === Infinity ? '∞' : old}: update ${v} and push it.`, [u, v]);
        } else {
          snap(u, 7, `${u} → ${v} costs ${d} + ${w} = ${d + w}, no better than ${v}'s ${dist.get(v)}: leave it.`, [u, v]);
        }
      });
    }
    snap(null, 2, 'The heap is empty: every reachable node is settled with its shortest distance.');
    return { steps, order: settled, dist, via };
  }

  function graphQuestion(exId, size) {
    if (exId === 'topo') return graphTopoQuestion(size);
    const g = makeGraph(size, { weighted: exId === 'dijkstra' });
    const start = pick(g.nodes).id;
    const run = exId === 'bfs' ? bfsRun(g, start) : exId === 'dfs' ? dfsRun(g, start) : dijkstraRun(g, start);
    const seq = run.order;
    const what = { bfs: 'breadth-first search (BFS)', dfs: 'depth-first search (DFS)', dijkstra: "Dijkstra's algorithm" }[exId];
    const reasonFor = (k) => {
      const v = seq[k];
      if (k === 0) return `Every search starts at the start node, ${start}.`;
      if (exId === 'bfs') {
        const p = run.parent.get(v);
        return `${v} is at the front of the queue. It was queued when ${p} was recorded, and everything queued before it has been taken out already: BFS finishes each distance before the next.`;
      }
      if (exId === 'dfs') {
        const p = run.parent.get(v);
        return p === seq[k - 1]
          ? `${v} is the smallest unvisited neighbour of ${p}, the node DFS just reached, so DFS goes there next.`
          : `${seq[k - 1]} has no unvisited neighbours left, so DFS backs up the call stack to ${p}, the closest node that still has one: ${v} is ${p}'s smallest unvisited neighbour.`;
      }
      return `When ${seq[k - 1]} is settled, ${v} has the smallest tentative distance (${run.dist.get(v)}, via ${run.via.get(v)}) of everything not settled yet.`;
    };
    return {
      prompt: exId === 'dijkstra'
        ? `Run <strong>${what}</strong> from <strong>${start}</strong>. Click the nodes in the order they are <em>settled</em> (popped with their final distance). Ties go to the smaller label.`
        : `Run <strong>${what}</strong> from <strong>${start}</strong>, taking neighbours in increasing order. Click the nodes in the order they are visited.`,
      aside: adjacencyHtml(g),
      mode: 'sequence',
      ids: g.nodes.map(nd => nd.id),
      label: (id) => id,
      svg: (view) => graphSvg(g, { ...view, target: view.target ?? (view.step ? undefined : start) }),
      expected: seq,
      fixedLength: true,
      code: GRAPH_CODE[exId],
      steps: () => run.steps,
      right: () => (exId === 'dijkstra'
        ? `Settled in order ${seq.join(', ')}, at distances ${seq.map(v => run.dist.get(v)).join(', ')}.`
        : `${exId.toUpperCase()} from ${start}: ${seq.join(', ')}.`),
      wrong: (k, answer) => {
        const right = k === 0 ? 'The first node is wrong.' : `The first ${k} ${k === 1 ? 'is' : 'are'} right.`;
        const other = exId === 'bfs' ? dfsRun(g, start).order : exId === 'dfs' ? bfsRun(g, start).order : null;
        const named = other && other.join() === answer.join() ? ` (That is the ${exId === 'bfs' ? 'DFS' : 'BFS'} order.)` : '';
        return `<p class="trace-bad">${right} Next comes ${seq[k]}, not ${answer[k]}.${named}</p><p class="trace-why">${reasonFor(k)}</p>`;
      },
    };
  }

  function graphTopoQuestion(size) {
    const g = makeGraph(size, { directed: true });
    const run = topoRun(g);
    const preds = new Map(g.nodes.map(nd => [nd.id, []]));
    g.edges.forEach(({ u, v }) => preds.get(v).push(u));
    // The first position where a node is placed before something it needs.
    const firstBad = (answer) => {
      const placed = new Set();
      for (let k = 0; k < answer.length; k++) {
        if (preds.get(answer[k]).some(p => !placed.has(p))) return k;
        placed.add(answer[k]);
      }
      return answer.length;
    };
    return {
      prompt: 'An arrow <strong>u → v</strong> means u must come before v. Click the nodes in a <strong>topological order</strong>. Any valid order counts.',
      aside: adjacencyHtml(g),
      mode: 'sequence',
      ids: g.nodes.map(nd => nd.id),
      label: (id) => id,
      svg: (view) => graphSvg(g, view),
      expected: run.order,
      fixedLength: true,
      validate: (answer) => {
        const k = firstBad(answer);
        return { ok: k === answer.length && answer.length === g.n, k };
      },
      code: GRAPH_CODE.topo,
      steps: () => run.steps,
      right: () => 'Every arrow points forward in your order. (Kahn\'s algorithm, taking the smallest ready node each time, gives ' + run.order.join(', ') + '.)',
      wrong: (k, answer) => {
        const x = answer[k];
        const placed = new Set(answer.slice(0, k));
        const missing = preds.get(x).filter(p => !placed.has(p));
        const ready = g.nodes.map(nd => nd.id).filter(v => !placed.has(v) && preds.get(v).every(p => placed.has(p))).sort((a, b) => a - b);
        return `<p class="trace-bad">${k === 0 ? 'The first node is wrong.' : `The first ${k} ${k === 1 ? 'is' : 'are'} fine.`} ${x} can't come yet: it needs ${missing.join(' and ')} first (${missing.map(p => `${p} → ${x}`).join(', ')}).</p>
          <p class="trace-why">At that point the nodes with nothing left to wait for are ${ready.join(', ')}. Any of them can go next.</p>`;
      },
    };
  }

  KITS.graphs = {
    title: 'Graphs',
    exercises: [
      { id: 'bfs', name: 'BFS order' },
      { id: 'dfs', name: 'DFS order' },
      { id: 'topo', name: 'Topological order' },
      { id: 'dijkstra', name: 'Dijkstra' },
    ],
    make: graphQuestion,
  };

  // ===========================================================================
  // Linked lists
  // ===========================================================================

  const LIST_SIZES = { small: [4, 6], medium: [6, 8], large: [8, 10] };
  const NONE = 'none';
  const DUMMY = 'dummy';

  // listSvg draws nodes in a row with their next arrows. next maps each id to
  // the id it points at (default: the following node, then None). Pointers
  // ({name: id}) are drawn as tags above their nodes.
  function listSvg(list, view, { next, dummy = false, cycleTo = null, showNone = true, clickNone = false } = {}) {
    const step = view.step;
    const ids = [...(dummy ? [DUMMY] : []), ...list.map((_, i) => i), ...(showNone ? [NONE] : [])];
    const gap = 74;
    const x0 = 40;
    const y = 70;
    const xOf = new Map(ids.map((id, i) => [id, x0 + i * gap]));
    const label = (id) => (id === NONE ? 'None' : id === DUMMY ? 'dummy' : list[id]);
    const nextOf = step?.next || next || ((id) => {
      if (id === DUMMY) return 0;
      if (id === list.length - 1) return cycleTo ?? NONE;
      return id + 1;
    });
    let arrows = '';
    ids.forEach(id => {
      if (id === NONE) return;
      const t = typeof nextOf === 'function' ? nextOf(id) : nextOf[id];
      if (t === undefined || !xOf.has(t)) return;
      const a = xOf.get(id);
      const b = xOf.get(t);
      if (b === a + gap) {
        arrows += `<line x1="${a + 18}" y1="${y}" x2="${b - 21}" y2="${y}" marker-end="url(#trace-arrow)" />`;
      } else {
        // Backward links arc below (a cycle) or above (a reversed pointer).
        const below = cycleTo !== null && t === cycleTo && id === list.length - 1;
        const dir = below ? 1 : -1;
        const h = 26 + Math.abs(b - a) * 0.12;
        const s = b < a ? -1 : 1;
        arrows += `<path d="M ${a} ${y + dir * 17} C ${a} ${y + dir * h * 1.6}, ${b} ${y + dir * h * 1.6}, ${b + s * -2} ${y + dir * 20}" fill="none" marker-end="url(#trace-arrow)" />`;
      }
    });
    const tags = {};
    Object.entries(step?.pointers || {}).forEach(([name, id]) => { (tags[id] = tags[id] || []).push(name); });
    let nodes = '';
    ids.forEach(id => {
      const x = xOf.get(id);
      const isNone = id === NONE;
      const allowed = isNone ? clickNone : id === DUMMY ? dummy : true;
      const v = { ...view, clickable: view.clickable && allowed };
      const tagText = (tags[id] || []).join(', ');
      const extra = tagText ? `<text class="ptr" x="${x}" y="${y - 30}">${esc(tagText)}</text>` : '';
      nodes += nodeSvg(id, x, y, label(id), v, extra).replace('<circle', isNone || id === DUMMY ? '<circle class="sentinel"' : '<circle');
    });
    const width = x0 * 2 + (ids.length - 1) * gap;
    return svgWrap(width, 150, `<g class="edges list-edges">${arrows}</g>${nodes}`, 'Linked list');
  }

  const LIST_CODE = {
    reverse: ['prev, curr = None, head', 'while curr:', '    nxt = curr.next', '    curr.next = prev', '    prev = curr', '    curr = nxt'],
    middle: ['slow = fast = head', 'while fast and fast.next:', '    slow = slow.next', '    fast = fast.next.next', 'return slow'],
    cycle: ['slow = fast = head', 'while fast and fast.next:', '    slow = slow.next', '    fast = fast.next.next', '    if slow is fast:', '        return True'],
    nth: ['dummy = ListNode(0, head)', 'lead = trail = dummy', 'for _ in range(n + 1):', '    lead = lead.next', 'while lead:', '    lead = lead.next', '    trail = trail.next', 'trail.next = trail.next.next'],
  };

  function listQuestion(exId, size) {
    const n = randInt(...LIST_SIZES[size]);
    const list = distinctValues(n);
    const lab = (id) => (id === NONE ? 'None' : id === DUMMY ? 'dummy' : list[id]);

    if (exId === 'reverse') {
      const k = randInt(1, n);
      const steps = [];
      const next = Object.fromEntries(list.map((_, i) => [i, i + 1 < n ? i + 1 : NONE]));
      let prev = NONE;
      let curr = 0;
      const snap = (line, text, extra = {}) => steps.push({ line, text, next: { ...next }, pointers: { prev, curr, ...extra }, current: curr === NONE ? undefined : curr, out: [] });
      snap(0, `Start: prev is None and curr is the head, ${lab(0)}.`);
      for (let it = 1; it <= k; it++) {
        const nxt = next[curr];
        snap(2, `Iteration ${it}: save curr.next (${lab(nxt)}) in nxt before it is overwritten.`, { nxt });
        next[curr] = prev;
        snap(3, `Flip ${lab(curr)}'s arrow to point back at prev (${lab(prev)}).`, { nxt });
        prev = curr;
        curr = nxt;
        snap(5, `Move both pointers one step: prev = ${lab(prev)}, curr = ${lab(curr)}.`);
      }
      snap(1, `That was iteration ${k}. prev is ${lab(prev)} and curr is ${lab(curr)}.`);
      return {
        prompt: `This loop reverses the list in place. Stop it after <strong>${k} iteration${k === 1 ? '' : 's'}</strong>. Click the node <strong>prev</strong> points to, then the node <strong>curr</strong> points to (None counts).`,
        mode: 'sequence',
        slots: ['prev', 'curr'],
        ids: [...list.map((_, i) => i), NONE],
        label: lab,
        svg: (view) => listSvg(list, view, { clickNone: true }),
        expected: [k - 1, k < n ? k : NONE],
        fixedLength: true,
        code: LIST_CODE.reverse,
        steps: () => steps,
        right: () => `After ${k} iteration${k === 1 ? '' : 's'} the first ${k} node${k === 1 ? ' points' : 's point'} backwards, prev is the last of them (${lab(k - 1)}) and curr is the first node not yet flipped (${lab(k < n ? k : NONE)}).`,
        wrong: () => `<p class="trace-bad">prev should be ${lab(k - 1)} and curr ${lab(k < n ? k : NONE)}.</p>
          <p class="trace-why">Each iteration flips one arrow and moves both pointers one node to the right. prev starts at None and curr at the head, so after ${k} iteration${k === 1 ? '' : 's'} prev is node ${k} of the list (${lab(k - 1)}) and curr is node ${k + 1}${k === n ? ', past the end, which is None' : ` (${lab(k)})`}.</p>`,
      };
    }

    if (exId === 'middle') {
      const steps = [];
      let slow = 0;
      let fast = 0;
      let iters = 0;
      const nx = (id) => (id === NONE || id + 1 >= n ? NONE : id + 1);
      steps.push({ line: 0, text: `Both start at the head, ${lab(0)}.`, pointers: { slow, fast }, current: slow, out: [] });
      while (fast !== NONE && nx(fast) !== NONE) {
        slow = nx(slow);
        fast = nx(nx(fast));
        iters++;
        steps.push({ line: 3, text: `slow moves one node to ${lab(slow)}; fast moves two to ${lab(fast)}.`, pointers: { slow, fast }, current: slow, out: [] });
      }
      steps.push({ line: 4, text: fast === NONE ? 'fast ran off the end, so the loop stops.' : `fast is on the last node and can't take two steps, so the loop stops. slow is on ${lab(slow)}.`, pointers: { slow, fast }, current: slow, out: [] });
      return {
        prompt: 'Fast and slow pointers start at the head; the loop below runs until fast can\'t take two steps. Click the node <strong>slow</strong> ends on.',
        mode: 'sequence',
        slots: ['slow'],
        ids: list.map((_, i) => i),
        label: lab,
        svg: (view) => listSvg(list, view),
        expected: [slow],
        fixedLength: true,
        code: LIST_CODE.middle,
        steps: () => steps,
        right: () => `slow ends on ${lab(slow)}, node ${slow + 1} of ${n}${n % 2 === 0 ? ' (the second of the two middles)' : ' (the middle)'}.`,
        wrong: () => `<p class="trace-bad">slow ends on ${lab(slow)}.</p>
          <p class="trace-why">fast moves twice as far as slow. The loop runs ${iters} time${iters === 1 ? '' : 's'} for ${n} nodes (until fast is on the last node or past it), so slow moves ${iters} node${iters === 1 ? '' : 's'} from the head to node ${slow + 1}.${n % 2 === 0 ? ' With an even length that is the second middle.' : ''}</p>`,
      };
    }

    if (exId === 'cycle') {
      const pos = randInt(0, n - 2);
      const nx = (id) => (id === n - 1 ? pos : id + 1);
      const steps = [];
      let slow = 0;
      let fast = 0;
      const trail = [];
      steps.push({ line: 0, text: `Both start at the head, ${lab(0)}.`, pointers: { slow, fast }, current: slow, out: [] });
      do {
        slow = nx(slow);
        fast = nx(nx(fast));
        trail.push(`${lab(slow)}/${lab(fast)}`);
        steps.push({ line: slow === fast ? 4 : 3, text: slow === fast ? `slow moves to ${lab(slow)} and fast lands on it too: they meet.` : `slow moves to ${lab(slow)}, fast two nodes to ${lab(fast)}.`, pointers: { slow, fast }, current: slow, out: [] });
      } while (slow !== fast);
      return {
        prompt: `The last node links back to ${lab(pos)}, making a cycle. Floyd's slow and fast pointers both start at the head. Click the node where they <strong>first meet</strong>.`,
        mode: 'sequence',
        slots: ['meet'],
        ids: list.map((_, i) => i),
        label: lab,
        svg: (view) => listSvg(list, view, { cycleTo: pos, showNone: false }),
        expected: [slow],
        fixedLength: true,
        code: LIST_CODE.cycle,
        steps: () => steps,
        right: () => `They meet at ${lab(slow)} after ${trail.length} move${trail.length === 1 ? '' : 's'}.`,
        wrong: () => `<p class="trace-bad">They first meet at ${lab(slow)}.</p>
          <p class="trace-why">After each move (slow/fast): ${trail.join(' → ')}. Once both are in the cycle, fast gains one node per move, so it catches slow within one lap.</p>`,
      };
    }

    // Remove the m-th node from the end with a gap of m + 1.
    const m = randInt(1, n);
    const order = [DUMMY, ...list.map((_, i) => i), NONE];
    const steps = [];
    let lead = 0; // indexes into order
    let trail = 0;
    const at = (i) => order[i];
    steps.push({ line: 1, text: 'lead and trail both start at the dummy node in front of the head.', pointers: { lead: at(lead), trail: at(trail) }, current: at(trail), out: [] });
    for (let i = 0; i <= m; i++) {
      lead++;
      steps.push({ line: 3, text: `Open the gap: lead moves to ${lab(at(lead))} (${i + 1} of ${m + 1} moves).`, pointers: { lead: at(lead), trail: at(trail) }, current: at(trail), out: [] });
    }
    while (at(lead) !== NONE) {
      lead++;
      trail++;
      steps.push({ line: 6, text: `Move together: lead to ${lab(at(lead))}, trail to ${lab(at(trail))}.`, pointers: { lead: at(lead), trail: at(trail) }, current: at(trail), out: [] });
    }
    steps.push({ line: 7, text: `lead is None, so trail (${lab(at(trail))}) is just before the node to remove, ${lab(at(trail + 1))}. Its next pointer skips over it.`, pointers: { lead: at(lead), trail: at(trail) }, current: at(trail), out: [] });
    const ordinal = (x) => `${x}${['th', 'st', 'nd', 'rd'][(x % 100 > 10 && x % 100 < 14) || x % 10 > 3 ? 0 : x % 10]}`;
    return {
      prompt: `Remove the <strong>${ordinal(m)} node from the end</strong> (${lab(n - m)}) with two pointers and a dummy head. Click the node <strong>trail</strong> stops on: the one whose next pointer changes. The dummy counts.`,
      mode: 'sequence',
      slots: ['trail'],
      ids: [DUMMY, ...list.map((_, i) => i)],
      label: lab,
      svg: (view) => listSvg(list, view, { dummy: true }),
      expected: [at(trail)],
      fixedLength: true,
      code: LIST_CODE.nth,
      steps: () => steps,
      right: () => `trail stops on ${lab(at(trail))}, and ${lab(at(trail))}.next skips ${lab(n - m)}.${at(trail) === DUMMY ? ' This is why the dummy exists: removing the head needs a node in front of it.' : ''}`,
      wrong: () => `<p class="trace-bad">trail stops on ${lab(at(trail))}.</p>
        <p class="trace-why">lead starts ${m + 1} nodes ahead of trail and they move together, so when lead falls off the end (None), trail is ${m + 1} nodes before None: the node just before the ${ordinal(m)} from the end.${at(trail) === DUMMY ? ' Here that is the dummy itself, because the head is the one being removed.' : ''}</p>`,
    };
  }

  KITS['linked-list'] = {
    title: 'Linked lists',
    exercises: [
      { id: 'reverse', name: 'Reverse: where are prev & curr?' },
      { id: 'middle', name: 'Fast & slow: middle' },
      { id: 'cycle', name: 'Cycle: where do they meet?' },
      { id: 'nth', name: 'Remove nth from end' },
    ],
    make: listQuestion,
  };

  // ===========================================================================
  // Heaps
  // ===========================================================================

  const HEAP_SIZES = { small: 6, medium: 9, large: 13 };
  const parentOf = (i) => (i - 1) >> 1;
  const ALL_VALUES = Array.from({ length: 99 }, (_, i) => i + 1);

  function heapPush(a, x) {
    a.push(x);
    let i = a.length - 1;
    while (i > 0 && a[parentOf(i)] > a[i]) {
      [a[parentOf(i)], a[i]] = [a[i], a[parentOf(i)]];
      i = parentOf(i);
    }
    return i;
  }

  // siftDown swaps with the smaller child (the left one on a tie) until the
  // value is no bigger than both children. log gets one entry per swap.
  function siftDown(a, i, log = []) {
    for (;;) {
      let c = 2 * i + 1;
      if (c >= a.length) return i;
      if (c + 1 < a.length && a[c + 1] < a[c]) c++;
      if (a[i] <= a[c]) return i;
      log.push(a[c]);
      [a[i], a[c]] = [a[c], a[i]];
      i = c;
    }
  }

  function makeHeap(values) {
    const a = [];
    values.forEach(v => heapPush(a, v));
    return a;
  }

  function heapify(input, notes = []) {
    const a = [...input];
    for (let i = (a.length >> 1) - 1; i >= 0; i--) {
      const v = a[i];
      const swaps = [];
      const end = siftDown(a, i, swaps);
      notes.push(swaps.length
        ? `index ${i} (${v}) swaps with ${swaps.join(', then ')} and ends at index ${end}`
        : `index ${i} (${v}) is no bigger than its children and stays`);
    }
    return a;
  }

  const isMinHeap = (a) => a.every((v, i) => i === 0 || a[parentOf(i)] <= v);

  // heapSvg draws a heap twice: as the tree it means and as the array it is
  // stored in. Node ids are array indexes; step.vals overrides the values
  // during a replay, and mark rings one index before the replay starts.
  function heapSvg(vals, view, { mark = null, clickableIds = null } = {}) {
    const a = view.step?.vals || vals;
    const n = a.length;
    const levels = Math.floor(Math.log2(n)) + 1;
    const treeW = 2 ** (levels - 1) * 50;
    const cell = 36;
    const width = Math.max(treeW, n * cell) + 40;
    const off = (width - treeW) / 2;
    const dy = 62;
    const pos = (i) => {
      const d = Math.floor(Math.log2(i + 1));
      return { x: off + (i - (2 ** d - 1) + 0.5) * treeW / 2 ** d, y: 30 + d * dy };
    };
    const hot = view.step?.hot || [];
    let edges = '';
    for (let i = 1; i < n; i++) {
      const p = pos(parentOf(i));
      const c = pos(i);
      const isHot = hot.includes(i) && hot.includes(parentOf(i));
      edges += `<line class="${isHot ? 'hot' : ''}" x1="${p.x}" y1="${p.y}" x2="${c.x}" y2="${c.y}" />`;
    }
    let nodes = '';
    let cells = '';
    const ay = 30 + levels * dy - 10;
    const ax = (width - n * cell) / 2;
    for (let i = 0; i < n; i++) {
      const { x, y } = pos(i);
      const allowed = !clickableIds || clickableIds.includes(i);
      const v = { ...view, clickable: view.clickable && allowed, target: !view.step && mark === i ? i : view.target };
      nodes += nodeSvg(i, x, y, a[i], v);
      const cls = ['heap-cell'];
      if (view.current === i) cls.push('current');
      else if (view.held.has(i)) cls.push('held');
      if (view.picked.get(i)) cls.push('picked');
      cells += `<g class="${cls.join(' ')}"><rect x="${ax + i * cell + 2}" y="${ay}" width="${cell - 4}" height="28" rx="4" />
        <text x="${ax + i * cell + cell / 2}" y="${ay + 14}" dy="0.35em">${a[i]}</text>
        <text class="idx" x="${ax + i * cell + cell / 2}" y="${ay + 42}">${i}</text></g>`;
    }
    return svgWrap(width, ay + 52, `<g class="edges">${edges}</g>${nodes}<g class="heap-array">${cells}</g>`, 'Min-heap as a tree and as an array');
  }

  const HEAP_CODE = {
    push: ['heap.append(x)', 'i = len(heap) - 1', 'while i > 0:', '    p = (i - 1) // 2', '    if heap[p] <= heap[i]: break', '    heap[p], heap[i] = heap[i], heap[p]', '    i = p'],
    pop: ['top = heap[0]', 'heap[0] = heap.pop()', 'i = 0', 'while True:', '    c = 2 * i + 1', '    if c >= len(heap): break', '    if c + 1 < len(heap) and heap[c + 1] < heap[c]: c += 1', '    if heap[i] <= heap[c]: break', '    heap[i], heap[c] = heap[c], heap[i]', '    i = c'],
    spot: ['for i in range(1, len(heap)):', '    p = (i - 1) // 2', '    if heap[p] > heap[i]:', '        return i  # smaller than its parent'],
  };

  function heapQuestion(exId, size) {
    const n = HEAP_SIZES[size];
    const idx = (id) => `index ${id}`;

    if (exId === 'push') {
      const base = makeHeap(distinctValues(n - 1));
      const x = pick(ALL_VALUES.filter(v => !base.includes(v)));
      const a = [...base, x];
      const steps = [];
      const cmp = [];
      const snap = (line, text, extra = {}) => steps.push({ line, text, vals: [...a], ...extra });
      let i = n - 1;
      snap(0, `Append ${x} in the first free slot, index ${i}, so the tree stays complete.`, { current: i });
      while (i > 0) {
        const p = parentOf(i);
        if (a[p] <= a[i]) {
          cmp.push(`${x} > ${a[p]}, stop`);
          snap(4, `Its parent at index ${p} holds ${a[p]}, which is smaller, so ${x} stops at index ${i}.`, { current: i, heldIds: [p], hot: [i, p] });
          break;
        }
        cmp.push(`${x} < ${a[p]}, swap`);
        snap(5, `${x} is smaller than its parent ${a[p]} at index ${p}: swap them.`, { current: i, heldIds: [p], hot: [i, p] });
        [a[p], a[i]] = [a[i], a[p]];
        i = p;
        snap(6, `${x} is now at index ${i}.`, { current: i });
      }
      if (i === 0) { cmp.push('it reached the root'); snap(2, `${x} reached the root: it is the new minimum.`, { current: 0 }); }
      return {
        prompt: `<strong>Push ${x}.</strong> It starts in the next free slot (index ${n - 1}, ringed) and sifts up. Click the node where ${x} comes to rest.`,
        mode: 'sequence',
        slots: ['rests at'],
        ids: a.map((_, j) => j),
        label: idx,
        svg: (view) => heapSvg([...base, x], view, { mark: n - 1 }),
        expected: [i],
        fixedLength: true,
        code: HEAP_CODE.push,
        steps: () => steps,
        right: () => `${cmp.join('; ')}. ${x} settles at index ${i}.`,
        wrong: () => `<p class="trace-bad">${x} comes to rest at index ${i}.</p>
          <p class="trace-why">Sift up only ever compares with the parent at (i - 1) // 2, never a sibling, and swaps while the new value is smaller: ${cmp.join('; ')}.</p>`,
      };
    }

    if (exId === 'pop') {
      const heap = makeHeap(distinctValues(n));
      const a = [...heap];
      const top = a[0];
      const last = a[n - 1];
      const steps = [];
      const cmp = [];
      const snap = (line, text, extra = {}) => steps.push({ line, text, vals: [...a], ...extra });
      snap(0, `The minimum, ${top}, is at the root. That is what pop returns.`, { current: 0 });
      a[0] = a.pop();
      snap(1, `Move the last value, ${last}, to the root, so the tree stays complete.`, { current: 0 });
      let i = 0;
      for (;;) {
        let c = 2 * i + 1;
        if (c >= a.length) { cmp.push('no children left'); snap(5, `Index ${i} has no children, so ${last} stays here.`, { current: i }); break; }
        const two = c + 1 < a.length;
        if (two && a[c + 1] < a[c]) c++;
        const kids = two ? `the smaller child is ${a[c]} (of ${a[2 * i + 1]} and ${a[2 * i + 2]})` : `the only child is ${a[c]}`;
        if (a[i] <= a[c]) {
          cmp.push(`${kids}, ${last} < ${a[c]}, stop`);
          snap(7, `At index ${i}, ${kids}. ${last} is smaller, so it stops.`, { current: i, heldIds: [c], hot: [i, c] });
          break;
        }
        cmp.push(`${kids}, swap`);
        snap(8, `At index ${i}, ${kids}. ${last} is bigger: swap with it.`, { current: i, heldIds: [c], hot: [i, c] });
        [a[i], a[c]] = [a[c], a[i]];
        i = c;
        snap(9, `${last} is now at index ${i}.`, { current: i });
      }
      const ids = a.map((_, j) => j);
      return {
        prompt: `<strong>Pop the minimum</strong> (${top}). The last value, ${last} (ringed), moves to the root and sifts down. Click the node where ${last} comes to rest.`,
        mode: 'sequence',
        slots: ['rests at'],
        ids,
        label: idx,
        svg: (view) => heapSvg(heap, view, { mark: n - 1, clickableIds: ids }),
        expected: [i],
        fixedLength: true,
        code: HEAP_CODE.pop,
        steps: () => steps,
        right: () => `${cap(cmp.join('; '))}. ${last} settles at index ${i}.`,
        wrong: () => `<p class="trace-bad">${last} comes to rest at index ${i}.</p>
          <p class="trace-why">Sift down swaps with the <em>smaller</em> child, so the child that moves up is no bigger than its sibling and the heap rule holds on both sides. Step by step: ${cmp.join('; ')}.</p>`,
      };
    }

    if (exId === 'heapify') {
      let input;
      let right;
      let notes;
      let wrongs;
      for (let tries = 0; tries < 100; tries++) {
        input = distinctValues(n);
        if (isMinHeap(input)) continue;
        notes = [];
        right = heapify(input, notes);
        const topDown = [...input];
        for (let i = 0; i < n >> 1; i++) siftDown(topDown, i);
        const seen = new Set([right.join()]);
        wrongs = [
          { a: makeHeap(input), why: 'pushing the values one at a time (sift up), which is O(n log n) and usually gives a different array' },
          { a: [...input].sort((x, y) => x - y), why: 'sorting, which is a valid heap but not what heapify does' },
          { a: topDown, why: 'sifting down from the root forwards; values deep in the tree never get a chance to rise' },
        ].filter(w => !seen.has(w.a.join()) && seen.add(w.a.join()));
        if (wrongs.length >= 2) break;
      }
      const choices = shuffle([{ id: 'right', a: right }, ...wrongs.map((w, i) => ({ id: `w${i}`, a: w.a, why: w.why }))]);
      return {
        prompt: 'Bottom-up <strong>heapify</strong> sifts down every parent, from the last parent back to the root. Which array does it turn this one into?',
        mode: 'choice',
        svg: (view) => heapSvg(input, view),
        sequence: input,
        choices: choices.map(c => ({ id: c.id, label: `[${c.a.join(', ')}]` })),
        answerId: 'right',
        explain: (picked) => {
          const w = choices.find(c => c.id === picked);
          return `<p class="${picked === 'right' ? 'trace-ok' : 'trace-bad'}"><strong>${picked === 'right' ? 'Right' : 'Not quite'}:</strong> heapify gives [${right.join(', ')}].${w?.why ? ` The one you picked comes from ${w.why}.` : ''}</p>
            <p class="trace-why">The last parent is at index ${(n >> 1) - 1} (n // 2 - 1); leaves are already heaps. Then:</p>
            <ol class="trace-notes">${notes.map(t => `<li>${esc(t)}</li>`).join('')}</ol>`;
        },
      };
    }

    // spot: one value is smaller than its parent.
    let a;
    let bad;
    for (let tries = 0; tries < 100 && !a; tries++) {
      const heap = makeHeap(distinctValues(n));
      bad = randInt(1, n - 1);
      const room = ALL_VALUES.filter(v => v < heap[parentOf(bad)] && !heap.includes(v));
      if (room.length) { a = [...heap]; a[bad] = pick(room); }
    }
    const p = parentOf(bad);
    const steps = [];
    for (let i = 1; i <= bad; i++) {
      const q = parentOf(i);
      steps.push({ line: i === bad ? 3 : 2, text: i === bad ? `Index ${i}: parent ${a[q]} > ${a[i]}. ${a[i]} is smaller than its parent, so the heap is broken here.` : `Index ${i}: parent ${a[q]} <= ${a[i]}, fine.`, vals: a, current: i, heldIds: [q], hot: [i, q] });
    }
    return {
      prompt: 'This array should be a min-heap, but <strong>one value is smaller than its parent</strong>. Click it.',
      mode: 'sequence',
      slots: ['broken'],
      ids: a.map((_, j) => j),
      label: (id) => String(a[id]),
      svg: (view) => heapSvg(a, view),
      expected: [bad],
      fixedLength: true,
      code: HEAP_CODE.spot,
      steps: () => steps,
      right: () => `${a[bad]} sits under ${a[p]}. In a min-heap every parent is no bigger than its children; in the array, heap[(i - 1) // 2] <= heap[i].`,
      wrong: (k, answer) => `<p class="trace-bad">It is ${a[bad]} at index ${bad}.</p>
        <p class="trace-why">${answer[0] === p ? `${a[p]} is the parent of the broken pair; the value that is out of place relative to its parent is its child ${a[bad]}.` : `Check each node against its parent only: ${a[bad]} is under ${a[p]}, which is bigger. Siblings and cousins can be in any order.`}</p>`,
    };
  }

  KITS['heap-priority-queue'] = {
    title: 'Heaps',
    exercises: [
      { id: 'push', name: 'Push: where does it settle?' },
      { id: 'pop', name: 'Pop: where does the last value settle?' },
      { id: 'heapify', name: 'Heapify: which array?' },
      { id: 'spot', name: 'Spot the broken node' },
    ],
    make: heapQuestion,
  };

  // ===========================================================================
  // Tries
  // ===========================================================================

  const TRIE_BANK = {
    a: ['a', 'an', 'and', 'ant', 'any', 'app', 'apple', 'apply', 'ape'],
    b: ['ban', 'band', 'bat', 'bath', 'be', 'bee', 'beef', 'bet', 'bed'],
    c: ['can', 'cane', 'cap', 'cape', 'car', 'card', 'care', 'cart', 'cat'],
    d: ['do', 'dog', 'dot', 'dove', 'dote', 'dig', 'dim', 'dip'],
    t: ['tea', 'team', 'tear', 'ten', 'tent', 'to', 'toe', 'top', 'tops'],
  };
  const TRIE_SIZES = { small: [1, 4], medium: [2, 3], large: [3, 3] }; // groups, words per group
  const LETTERS = 'abcdefghijklmnopqrstuvwxyz';

  function makeTrie(words) {
    const root = { id: '^', ch: '', children: new Map(), end: false, depth: 0 };
    const nodes = [root];
    words.forEach(w => {
      let node = root;
      for (const ch of w) {
        if (!node.children.has(ch)) {
          const c = { id: node.id + ch, ch, children: new Map(), end: false, depth: node.depth + 1, parent: node };
          node.children.set(ch, c);
          nodes.push(c);
        }
        node = node.children.get(ch);
      }
      node.end = true;
    });
    // Leaves left to right in alphabetical order; a parent sits over its children.
    let leaf = 0;
    (function place(node) {
      node.kids = [...node.children.keys()].sort().map(k => node.children.get(k));
      if (!node.kids.length) { node.x = leaf++; return; }
      node.kids.forEach(place);
      node.x = (node.kids[0].x + node.kids[node.kids.length - 1].x) / 2;
    })(root);
    return { root, nodes, byId: new Map(nodes.map(n => [n.id, n])), words, depth: Math.max(...nodes.map(n => n.depth)) };
  }

  function trieWords(size) {
    const [groups, per] = TRIE_SIZES[size];
    return shuffle(Object.keys(TRIE_BANK)).slice(0, groups).flatMap(g => shuffle([...TRIE_BANK[g]]).slice(0, per));
  }

  const word = (id) => id.slice(1);
  const quote = (s) => `'${s}'`;

  function trieSvg(trie, view, { clickableIds = null } = {}) {
    const dx = 44;
    const dy = 56;
    const pad = 30;
    const px = (x) => x * dx + pad;
    const py = (d) => d * dy + pad;
    const hot = new Set(view.step?.out || []);
    let edges = '';
    let nodes = '';
    trie.nodes.forEach(n => {
      n.kids.forEach(c => {
        edges += `<line class="${hot.has(n.id) && hot.has(c.id) ? 'hot' : ''}" x1="${px(n.x)}" y1="${py(n.depth)}" x2="${px(c.x)}" y2="${py(c.depth)}" />`;
      });
      const allowed = !clickableIds || clickableIds.includes(n.id);
      const v = { ...view, clickable: view.clickable && allowed };
      const ring = n.end ? `<circle class="end-ring" cx="${px(n.x)}" cy="${py(n.depth)}" r="21" />` : '';
      let g = nodeSvg(n.id, px(n.x), py(n.depth), n.depth ? n.ch : 'root', v, ring);
      if (!n.depth) g = g.replace('<circle', '<circle class="sentinel"');
      nodes += g;
    });
    const width = Math.max(...trie.nodes.map(n => n.x)) * dx + 2 * pad;
    return svgWrap(width, trie.depth * dy + 2 * pad, `<g class="edges">${edges}</g>${nodes}`, 'Trie');
  }

  const TRIE_CODE = {
    search: ['node = root', 'for ch in word:', '    if ch not in node.children:', '        return False', '    node = node.children[ch]', 'return node.is_word'],
    insert: ['node = root', 'for ch in word:', '    if ch not in node.children:', '        node.children[ch] = TrieNode()', '    node = node.children[ch]', 'node.is_word = True'],
    collect: ['node = walk(prefix)  # None if missing', 'def dfs(node, path):', '    if node.is_word:', '        out.append(path)', '    for ch in sorted(node.children):', '        dfs(node.children[ch], path + ch)'],
  };

  // trieQuery picks a word to look up: a stored word, a prefix that is not a
  // word, or a word that falls off the trie.
  function trieQuery(trie, kind) {
    const inner = trie.nodes.filter(n => n.depth && !n.end);
    if (kind === 'prefix' && inner.length) return word(pick(inner).id);
    if (kind === 'word') return pick(trie.words);
    for (;;) {
      const w = pick(trie.words);
      const j = randInt(1, w.length);
      const node = trie.byId.get(`^${w.slice(0, j)}`);
      const free = [...LETTERS].filter(ch => !node.children.has(ch));
      const q = w.slice(0, j) + pick(free) + w.slice(j + 1);
      if (q.length <= 6) return q;
    }
  }

  function trieSearch(trie, q) {
    const path = [trie.root];
    let node = trie.root;
    for (const ch of q) {
      node = node.children.get(ch);
      if (!node) return { path, result: 'off', stuck: ch };
      path.push(node);
    }
    return { path, result: node.end ? 'yes' : 'flag' };
  }

  function trieSearchSteps(q, { path, result, stuck }) {
    const steps = [];
    const out = [];
    path.forEach((node, i) => {
      out.push(node.id);
      steps.push({ line: i ? 4 : 0, text: i ? `'${node.ch}' has a child: move to it. Spelled so far: ${quote(word(node.id))}.` : 'Start at the root.', current: node.id, out: [...out] });
    });
    const last = path[path.length - 1];
    steps.push(result === 'off'
      ? { line: 3, text: `No child for '${stuck}' under ${last.depth ? quote(word(last.id)) : 'the root'}: return False.`, current: last.id, out: [...out] }
      : { line: 5, text: `Every letter matched. ${quote(q)} ${result === 'yes' ? 'is marked as a word end: return True' : 'is not marked as a word end: return False'}.`, current: last.id, out: [...out] });
    return steps;
  }

  function trieQuestion(exId, size) {
    const trie = makeTrie(trieWords(size));
    const label = (id) => (id === '^' ? 'root' : word(id));
    const aside = `<p class="trace-aside">Words: ${trie.words.slice().sort().map(w => `<code>${esc(w)}</code>`).join(' ')}. A ring marks a word end.</p>`;

    if (exId === 'path' || exId === 'result') {
      const q = trieQuery(trie, pick(['word', 'prefix', 'missing']));
      const res = trieSearch(trie, q);
      const steps = trieSearchSteps(q, res);
      const last = res.path[res.path.length - 1];
      const why = {
        yes: `every letter of ${quote(q)} has a node and the last one is marked as a word end`,
        flag: `the path ${quote(q)} exists, but only as the start of a longer word: its last node is not marked as a word end`,
        off: `after ${last.depth ? quote(word(last.id)) : 'the root'} there is no child for '${res.stuck}', so the search falls off the trie`,
      }[res.result];
      if (exId === 'result') {
        const names = { yes: 'True', off: 'False: the path runs out', flag: 'False: no word ends there' };
        return {
          prompt: `What does <strong>search(${quote(q)})</strong> return?`,
          aside,
          mode: 'choice',
          svg: (view) => trieSvg(trie, { ...view, picked: new Map(view.checked ? res.path.map((n, i) => [n.id, i + 1]) : []) }),
          sequence: [q],
          choices: ['yes', 'off', 'flag'].map(id => ({ id, label: names[id] })),
          answerId: res.result,
          explain: (picked) => `<p class="${picked === res.result ? 'trace-ok' : 'trace-bad'}"><strong>${picked === res.result ? 'Right' : 'Not quite'}:</strong> ${names[res.result]}. ${cap(why)}.</p>
            ${res.result !== 'off' ? `<p class="trace-why">The end flag is what separates a stored word from a prefix of one: startsWith(${quote(q)}) is True either way.</p>` : ''}`,
        };
      }
      return {
        prompt: `<strong>search(${quote(q)})</strong>: click the nodes it visits, starting with the root.`,
        aside,
        mode: 'sequence',
        ids: trie.nodes.map(n => n.id),
        label,
        svg: (view) => trieSvg(trie, view),
        expected: res.path.map(n => n.id),
        fixedLength: false,
        code: TRIE_CODE.search,
        outLabel: 'Visited',
        steps: () => steps,
        right: () => `It visits ${res.path.map(n => label(n.id)).join(' → ')} and returns ${res.result === 'yes' ? 'True' : 'False'}: ${why}.`,
        wrong: (k, answer) => {
          const ok = k === 0 ? 'The first node is wrong.' : `The first ${k} ${k === 1 ? 'is' : 'are'} right.`;
          if (k === res.path.length) return `<p class="trace-bad">${ok} The search stops there: ${why}.</p>`;
          const next = res.path[k];
          const reason = k === 0 ? 'Every search starts at the root.' : `The next letter is '${next.ch}', so it follows that child only.`;
          return `<p class="trace-bad">${ok} ${k < answer.length ? `Next comes ${quote(label(next.id))}, not ${quote(label(answer[k]))}.` : `Keep going: next comes ${quote(label(next.id))}.`} ${reason}</p>`;
        },
      };
    }

    if (exId === 'insert') {
      let q;
      const r = Math.random();
      if (r < 0.2) q = trieQuery(trie, 'prefix');
      else if (r < 0.3) q = trieQuery(trie, 'word');
      else q = trieQuery(trie, 'missing');
      let node = trie.root;
      let shared = 0;
      for (const ch of q) {
        if (!node.children.has(ch)) break;
        node = node.children.get(ch);
        shared++;
      }
      const made = q.length - shared;
      const already = made === 0 && node.end;
      const opts = [...new Set([made, q.length, made + 1, Math.max(0, made - 1), 0])].slice(0, 4).sort((x, y) => x - y);
      const lead = shared ? `${quote(q.slice(0, shared))} already has a path, so insert walks it` : 'The root has no child for its first letter';
      const tail = made
        ? `then creates one node per remaining letter (${[...q.slice(shared)].map(c => `'${c}'`).join(', ')}): ${made}, and marks the last as a word end.`
        : already ? `and ${quote(q)} is already a word: nothing changes.` : 'and only sets the word-end flag on the last node: no new nodes.';
      return {
        prompt: `How many <strong>new nodes</strong> does <strong>insert(${quote(q)})</strong> create?`,
        aside,
        mode: 'choice',
        svg: (view) => trieSvg(trie, view),
        sequence: [q],
        choices: opts.map(o => ({ id: String(o), label: String(o) })),
        answerId: String(made),
        explain: (picked) => `<p class="${picked === String(made) ? 'trace-ok' : 'trace-bad'}"><strong>${picked === String(made) ? 'Right' : 'Not quite'}:</strong> ${made}. ${lead}${shared ? ',' : ','} ${tail}</p>
          <p class="trace-why">Shared prefixes are stored once: that is the whole point of a trie.</p>`,
      };
    }

    // collect: every word under a prefix, in DFS order.
    const under = (n) => (n.end ? 1 : 0) + n.kids.reduce((s, c) => s + under(c), 0);
    const starts = trie.nodes.filter(n => n.depth && under(n) >= 2);
    const start = starts.length ? pick(starts) : trie.root;
    const prefix = word(start.id);
    const steps = [];
    const out = [];
    let node = trie.root;
    steps.push({ line: 0, text: prefix ? `Walk down ${quote(prefix)} to the node where the prefix ends.` : 'The prefix is empty, so start at the root.', current: node.id, out: [] });
    for (const ch of prefix) {
      node = node.children.get(ch);
      steps.push({ line: 0, text: `'${ch}': move down to ${quote(word(node.id))}.`, current: node.id, out: [...out] });
    }
    (function dfs(n) {
      if (n.end) {
        out.push(n.id);
        steps.push({ line: 3, text: `${quote(word(n.id))} is a word end: record it, before looking at its children.`, current: n.id, out: [...out] });
      } else {
        steps.push({ line: 1, text: `Visit ${quote(word(n.id) || 'root')}: not a word end.`, current: n.id, out: [...out] });
      }
      n.kids.forEach(c => dfs(c));
    })(start);
    const ends = trie.nodes.filter(n => n.end).map(n => n.id);
    return {
      prompt: `<strong>wordsWithPrefix(${quote(prefix)})</strong> walks to the prefix, then does a DFS visiting children in alphabetical order. Click the words in the order it records them.`,
      aside,
      mode: 'sequence',
      ids: ends,
      label,
      svg: (view) => trieSvg(trie, view, { clickableIds: ends }),
      expected: out,
      fixedLength: false,
      code: TRIE_CODE.collect,
      steps: () => steps,
      right: () => `${out.map(label).join(', ')}: alphabetical, because the DFS takes children in sorted order and records a word before its longer extensions.`,
      wrong: (k, answer) => {
        const ok = k === 0 ? 'The first word is wrong.' : `The first ${k} ${k === 1 ? 'is' : 'are'} right.`;
        if (k === out.length) return `<p class="trace-bad">${ok} That is every word starting with ${quote(prefix)}.</p>`;
        const given = answer[k];
        let why = `It records a word as soon as it reaches its end node, then goes into the children in alphabetical order, so ${quote(label(out[k]))} comes next.`;
        if (given && !label(given).startsWith(prefix)) why = `${quote(label(given))} does not start with ${quote(prefix)}, so the DFS never reaches it.`;
        return `<p class="trace-bad">${ok} ${given ? `Next is ${quote(label(out[k]))}, not ${quote(label(given))}.` : `Keep going: next is ${quote(label(out[k]))}.`}</p><p class="trace-why">${why}</p>`;
      },
    };
  }

  KITS.tries = {
    title: 'Tries',
    exercises: [
      { id: 'path', name: 'Search: which nodes?' },
      { id: 'result', name: 'Search: what does it return?' },
      { id: 'insert', name: 'Insert: how many new nodes?' },
      { id: 'collect', name: 'Autocomplete order' },
    ],
    make: trieQuestion,
  };

  // ===========================================================================
  // Stacks
  // ===========================================================================

  const STACK_SIZES = { small: 5, medium: 7, large: 9 };

  // rowSvg draws values in a row with their index underneath. mark rings one
  // position before a replay; step.heldIds (what is on the stack) are dashed.
  function rowSvg(labels, view, { mark = null, clickableIds = null, showIndex = true } = {}) {
    const gap = 50;
    const x0 = 34;
    const y = 36;
    let nodes = '';
    labels.forEach((l, i) => {
      const allowed = !clickableIds || clickableIds.includes(i);
      const v = { ...view, clickable: view.clickable && allowed, target: !view.step && mark === i ? i : view.target };
      const extra = showIndex ? `<text class="idx" x="${x0 + i * gap}" y="${y + 36}">${i}</text>` : '';
      nodes += nodeSvg(i, x0 + i * gap, y, l, v, extra);
    });
    return svgWrap(x0 * 2 + (labels.length - 1) * gap, y + 50, nodes, 'Sequence of values');
  }

  const STACK_CODE = {
    lifo: ['stack = []', 'stack.append(x)        # push: onto the top', 'x = stack.pop()        # pop: off the top'],
    greater: ['ans = [-1] * len(nums)', 'stack = []  # indexes still waiting', 'for i, x in enumerate(nums):', '    while stack and nums[stack[-1]] < x:', '        ans[stack.pop()] = x', '    stack.append(i)'],
    match: ['stack = []', 'for i, ch in enumerate(s):', '    if ch in "([{":', '        stack.append(i)', '    else:', '        j = stack.pop()  # s[j] is closed by s[i]'],
  };

  // greaterRun is the monotonic stack for "next greater element", up to and
  // including index upTo. It returns the replay and the order of pops.
  function greaterRun(nums, upTo = nums.length - 1) {
    const steps = [];
    const stack = [];
    const pops = [];
    const show = () => ({ heldLabel: 'Stack (bottom → top)', held: stack.map(j => nums[j]), heldIds: [...stack], out: [...pops] });
    for (let i = 0; i <= upTo; i++) {
      const x = nums[i];
      while (stack.length && nums[stack[stack.length - 1]] < x) {
        const j = stack.pop();
        pops.push(j);
        steps.push({ line: 4, text: `${x} > ${nums[j]} on top: pop ${nums[j]}. Its next greater element is ${x}.`, current: i, ...show() });
      }
      stack.push(i);
      steps.push({ line: 5, text: stack.length > 1 ? `${x} is not bigger than the top, ${nums[stack[stack.length - 2]]}: push it. The stack still decreases from bottom to top.` : `The stack is empty${pops.length && steps.length ? ' now' : ''}: push ${x}.`, current: i, ...show() });
    }
    return { steps, stack, pops };
  }

  function stackQuestion(exId, size) {
    const n = STACK_SIZES[size];

    if (exId === 'lifo') {
      const vals = distinctValues(n);
      const ops = [];
      const stack = [];
      const pops = [];
      const steps = [];
      let pushed = 0;
      while (pushed < n || stack.length) {
        if (pushed < n && (!stack.length || Math.random() < 0.55)) {
          stack.push(pushed);
          ops.push(`push ${vals[pushed]}`);
          steps.push({ line: 1, text: `push ${vals[pushed]}: it goes on top.`, current: pushed, heldLabel: 'Stack (bottom → top)', held: stack.map(j => vals[j]), heldIds: [...stack], out: [...pops] });
          pushed++;
        } else {
          const j = stack.pop();
          pops.push(j);
          ops.push('pop');
          steps.push({ line: 2, text: `pop returns the top, ${vals[j]}: the most recent push still on the stack.`, current: j, heldLabel: 'Stack (bottom → top)', held: stack.map(k => vals[k]), heldIds: [...stack], out: [...pops] });
        }
      }
      return {
        prompt: 'Run these operations on an empty stack. Click the values in the order <strong>pop</strong> returns them.',
        aside: `<p class="trace-ops">${ops.map(o => `<code>${o}</code>`).join(' ')}</p>`,
        mode: 'sequence',
        ids: vals.map((_, i) => i),
        label: (id) => String(vals[id]),
        svg: (view) => rowSvg(vals, view, { showIndex: false }),
        expected: pops,
        fixedLength: true,
        code: STACK_CODE.lifo,
        outLabel: 'Popped',
        steps: () => steps,
        right: () => `Pops: ${pops.map(j => vals[j]).join(', ')}. Each pop takes the newest value still there.`,
        wrong: (k) => {
          const popIndex = ops.map((o, i) => [o, i]).filter(([o]) => o === 'pop')[k][1];
          const before = steps[popIndex].held.concat(vals[pops[k]]);
          return `<p class="trace-bad">${k === 0 ? 'The first pop is wrong.' : `The first ${k} ${k === 1 ? 'is' : 'are'} right.`} Pop number ${k + 1} returns ${vals[pops[k]]}.</p>
            <p class="trace-why">Just before it, the stack (bottom → top) is ${before.join(', ')}. Pop takes the top: last in, first out.</p>`;
        },
      };
    }

    if (exId === 'state' || exId === 'pops') {
      let nums;
      let run;
      let k;
      for (let tries = 0; tries < 100; tries++) {
        nums = distinctValues(n, 30);
        k = exId === 'state' ? randInt(2, n - 1) : n - 1;
        run = greaterRun(nums, k);
        if (exId === 'pops' ? run.pops.length >= 3 : run.pops.length >= 1 && run.stack.length >= 2) break;
      }
      const v = (ids) => ids.map(j => nums[j]).join(', ');
      if (exId === 'state') {
        return {
          prompt: `Next greater element with a monotonic stack. Stop after the loop has handled <strong>index ${k}</strong> (${nums[k]}, ringed). Click what is on the stack, <strong>bottom to top</strong>.`,
          mode: 'sequence',
          ids: nums.map((_, i) => i),
          label: (id) => String(nums[id]),
          svg: (view) => rowSvg(nums, view, { mark: k }),
          expected: run.stack,
          fixedLength: false,
          code: STACK_CODE.greater,
          outLabel: 'Popped',
          steps: () => run.steps,
          right: () => `The stack holds ${v(run.stack)}: the values so far with nothing bigger after them, decreasing from bottom to top.`,
          wrong: (m, answer) => `<p class="trace-bad">The stack is ${v(run.stack)} (bottom → top).</p>
            <p class="trace-why">A value stays on the stack until a bigger one arrives and pops it. So far ${run.pops.length ? `${v(run.pops)} ${run.pops.length === 1 ? 'was' : 'were'} popped` : 'nothing was popped'}; what is left decreases from bottom to top${answer.some(j => j > k) ? `, and nothing after index ${k} has been read yet` : ''}.</p>`,
        };
      }
      const ans = Object.fromEntries(run.steps.filter(s => s.line === 4).map((s, i) => [run.pops[i], nums[s.current]]));
      return {
        prompt: 'Next greater element with a monotonic stack. Click the values in the order they are <strong>popped</strong> (the moment each finds its next greater element).',
        mode: 'sequence',
        ids: nums.map((_, i) => i),
        label: (id) => String(nums[id]),
        svg: (view) => rowSvg(nums, view),
        expected: run.pops,
        fixedLength: false,
        code: STACK_CODE.greater,
        outLabel: 'Popped',
        steps: () => run.steps,
        right: () => `Pops: ${run.pops.map(j => `${nums[j]} (by ${ans[j]})`).join(', ')}. ${run.stack.length ? `${v(run.stack)} never get popped: nothing bigger comes after them, so their answer stays -1.` : ''}`,
        wrong: (m) => {
          const ok = m === 0 ? 'The first pop is wrong.' : `The first ${m} ${m === 1 ? 'is' : 'are'} right.`;
          if (m === run.pops.length) return `<p class="trace-bad">${ok} Nothing else is popped: ${v(run.stack)} stay on the stack with no greater element after them.</p>`;
          const j = run.pops[m];
          return `<p class="trace-bad">${ok} Next ${nums[j]} is popped, when ${ans[j]} arrives.</p>
            <p class="trace-why">Each new value pops everything smaller from the top of the stack, nearest first, before it is pushed itself.</p>`;
        },
      };
    }

    // match: which bracket closes the ringed one?
    const pairs = n;
    let s;
    let match;
    let t;
    for (let tries = 0; tries < 100; tries++) {
      s = [];
      const open = [];
      let left = pairs;
      while (s.length < 2 * pairs) {
        if (left && (!open.length || Math.random() < 0.5)) {
          const ch = pick(['(', '[', '{']);
          open.push(ch);
          s.push(ch);
          left--;
        } else s.push({ '(': ')', '[': ']', '{': '}' }[open.pop()]);
      }
      match = {};
      const st = [];
      s.forEach((ch, i) => { if ('([{'.includes(ch)) st.push(i); else match[st.pop()] = i; });
      const good = Object.keys(match).map(Number).filter(i => match[i] - i >= 3);
      if (good.length) { t = pick(good); break; }
    }
    const steps = [];
    const st = [];
    for (let i = 0; i <= match[t]; i++) {
      const ch = s[i];
      if ('([{'.includes(ch)) {
        st.push(i);
        steps.push({ line: 3, text: `'${ch}' opens: push it.`, current: i, heldLabel: 'Stack (bottom → top)', held: st.map(j => s[j]), heldIds: [...st] });
      } else {
        const j = st.pop();
        steps.push({ line: 5, text: `'${ch}' closes the opener on top, '${s[j]}' at index ${j}${j === t ? ': the ringed one.' : '.'}`, current: i, heldLabel: 'Stack (bottom → top)', held: st.map(x => s[x]), heldIds: [...st, j] });
      }
    }
    const closers = s.map((ch, i) => i).filter(i => !'([{'.includes(s[i]));
    return {
      prompt: `Valid Parentheses pushes each opener and pops one for each closer. Click the bracket that closes <strong>the ringed '${s[t]}'</strong> at index ${t}.`,
      mode: 'sequence',
      slots: ['closed by'],
      ids: s.map((_, i) => i),
      label: (id) => `${s[id]} at ${id}`,
      svg: (view) => rowSvg(s, view, { mark: t, clickableIds: closers }),
      expected: [match[t]],
      fixedLength: true,
      code: STACK_CODE.match,
      steps: () => steps,
      right: () => `The '${s[match[t]]}' at index ${match[t]}: everything opened after index ${t} is closed first, so ${t} is back on top when it arrives.`,
      wrong: (k, answer) => `<p class="trace-bad">It is the '${s[match[t]]}' at index ${match[t]}.</p>
        <p class="trace-why">${answer[0] < t ? `Index ${answer[0]} comes before the opener, so it cannot close it.` : answer[0] < match[t] ? `When index ${answer[0]} arrives, something opened after index ${t} is on top of the stack, so that is what it closes.` : `By index ${answer[0]} the opener was already closed at index ${match[t]}.`} A closer always matches the most recent opener still open.</p>`,
    };
  }

  KITS.stack = {
    title: 'Stacks',
    exercises: [
      { id: 'lifo', name: 'Push & pop order' },
      { id: 'match', name: 'Matching brackets' },
      { id: 'state', name: 'Monotonic stack: contents' },
      { id: 'pops', name: 'Monotonic stack: pop order' },
    ],
    make: stackQuestion,
  };

  // ===========================================================================
  // Arrays with pointers: binary search, sliding window, two pointers
  // ===========================================================================

  // arraySvg draws cells with their index underneath. Ids are indexes; with
  // end set, an extra "end" cell at index n stands for len(nums). A step can
  // override the values (step.vals), name pointers (step.pointers, drawn as
  // tags above their cells) and shade a window (step.window = [lo, hi],
  // inclusive). mark rings one cell before a replay starts.
  function arraySvg(cells, view, { mark = null, clickableIds = null, end = false, window: win = null } = {}) {
    const step = view.step;
    const a = step?.vals || cells;
    const w = 44;
    const x0 = 24;
    const y = 54;
    const ids = [...a.map((_, i) => i), ...(end ? [a.length] : [])];
    const xOf = (i) => x0 + i * w;
    const range = step ? step.window : win;
    let body = '';
    if (range && range[0] <= range[1]) {
      body += `<rect class="trace-window" x="${xOf(range[0]) - 2}" y="${y - 4}" width="${(range[1] - range[0] + 1) * w}" height="44" rx="7" />`;
    }
    const tags = {};
    Object.entries(step?.pointers || {}).forEach(([name, id]) => { (tags[id] = tags[id] || []).push(name); });
    ids.forEach(i => {
      const isEnd = i === a.length;
      const cls = ['tnode', 'acell'];
      const idx = view.picked.get(i);
      if (idx) cls.push('picked');
      if (view.current === i) cls.push('current');
      if (view.held.has(i) && view.current !== i) cls.push('held');
      if (view.wrong === i) cls.push('wrong');
      if ((!step && mark === i) || view.target === i) cls.push('target');
      if (isEnd) cls.push('end');
      const allowed = !clickableIds || clickableIds.includes(i);
      const clickable = view.clickable && allowed && !idx;
      if (clickable) cls.push('clickable');
      const x = xOf(i);
      const tag = tags[i] ? `<text class="ptr" x="${x + w / 2 - 2}" y="${y - 14}">${esc(tags[i].join(', '))}</text>` : '';
      body += `<g class="${cls.join(' ')}" data-id="${i}" ${clickable ? `role="button" tabindex="0" aria-label="Index ${i}"` : ''}>
        <rect x="${x + 2}" y="${y}" width="${w - 8}" height="36" rx="5" />
        <text x="${x + w / 2 - 2}" y="${y + 18}" dy="0.35em">${isEnd ? 'end' : esc(a[i])}</text>
        <text class="idx" x="${x + w / 2 - 2}" y="${y + 52}">${i}</text>
        ${idx ? `<text class="badge" x="${x + w - 10}" y="${y - 2}">${idx}</text>` : ''}
        ${tag}</g>`;
    });
    // Tags for pointers that are off the array (like hi = -1) go at the edge.
    Object.entries(tags).forEach(([id, names]) => {
      if (+id >= 0 && +id < ids.length) return;
      const x = +id < 0 ? x0 - 6 : xOf(ids.length) + 10;
      body += `<text class="ptr off" x="${x}" y="${y + 22}">${esc(names.join(', '))}</text>`;
    });
    return svgWrap(x0 * 2 + ids.length * w + 10, y + 64, body, 'Array');
  }

  const cellLabel = (a) => (id) => (id >= 0 && id < a.length ? `[${id}] ${a[id]}` : id < 0 ? `${id} (before the start)` : `${id} (end)`);
  const ARRAY_SIZES = { small: 7, medium: 10, large: 14 };
  const sortedDistinct = (n) => distinctValues(n).sort((x, y) => x - y);
  const said = (k) => (k === 0 ? 'The first one is wrong.' : `The first ${k} ${k === 1 ? 'is' : 'are'} right.`);

  // ---------------------------------------------------------------------------
  // Binary search

  const BS_CODE = {
    search: ['lo, hi = 0, len(nums) - 1', 'while lo <= hi:', '    mid = (lo + hi) // 2', '    if nums[mid] == target: return mid', '    if nums[mid] < target: lo = mid + 1', '    else: hi = mid - 1', 'return -1'],
    lower: ['lo, hi = 0, len(nums)', 'while lo < hi:', '    mid = (lo + hi) // 2', '    if nums[mid] < target: lo = mid + 1', '    else: hi = mid', 'return lo'],
    first: ['lo, hi = 0, len(ok)', 'while lo < hi:', '    mid = (lo + hi) // 2', '    if ok[mid]: hi = mid       # mid might be the first', '    else: lo = mid + 1', 'return lo'],
  };

  function bsSearchRun(a, t) {
    const steps = [];
    const mids = [];
    let lo = 0;
    let hi = a.length - 1;
    const states = [];
    steps.push({ line: 0, text: `The answer, if any, is somewhere in lo..hi: the whole array.`, pointers: { lo, hi }, window: [lo, hi] });
    while (lo <= hi) {
      const mid = (lo + hi) >> 1;
      mids.push(mid);
      steps.push({ line: 2, text: `mid = (${lo} + ${hi}) // 2 = ${mid}, which holds ${a[mid]}.`, pointers: { lo, hi, mid }, window: [lo, hi], current: mid, out: [...mids] });
      if (a[mid] === t) {
        steps.push({ line: 3, text: `${a[mid]} is the target: return ${mid}.`, pointers: { lo, hi, mid }, window: [mid, mid], current: mid, out: [...mids] });
        return { steps, mids, found: mid, states };
      }
      if (a[mid] < t) {
        lo = mid + 1;
        steps.push({ line: 4, text: `${a[mid]} < ${t}, so the target can only be right of mid: lo = ${lo}.`, pointers: { lo, hi }, window: [lo, hi], out: [...mids] });
      } else {
        hi = mid - 1;
        steps.push({ line: 5, text: `${a[mid]} > ${t}, so the target can only be left of mid: hi = ${hi}.`, pointers: { lo, hi }, window: [lo, hi], out: [...mids] });
      }
      states.push({ lo, hi, stepIndex: steps.length });
    }
    steps.push({ line: 6, text: `lo (${lo}) passed hi (${hi}): the range is empty, so ${t} is not here. Return -1.`, pointers: { lo, hi }, out: [...mids] });
    return { steps, mids, found: -1, states };
  }

  // halfOpenRun is the lo < hi template: lower bound when ok(i) is
  // nums[i] >= target, or "first True" for a boolean array.
  function halfOpenRun(n, ok, say) {
    const steps = [];
    const mids = [];
    let lo = 0;
    let hi = n;
    steps.push({ line: 0, text: `The answer is in lo..hi, and hi = ${n} (the end) is allowed: it means "none".`, pointers: { lo, hi }, window: [lo, hi - 1] });
    while (lo < hi) {
      const mid = (lo + hi) >> 1;
      mids.push(mid);
      steps.push({ line: 2, text: `mid = (${lo} + ${hi}) // 2 = ${mid}.`, pointers: { lo, hi, mid }, window: [lo, hi - 1], current: mid, out: [...mids] });
      if (ok(mid)) {
        hi = mid;
        steps.push({ line: say.yesLine, text: `${say.yes(mid)}, so the answer is mid or left of it: hi = ${hi}.`, pointers: { lo, hi }, window: [lo, hi - 1], out: [...mids] });
      } else {
        lo = mid + 1;
        steps.push({ line: say.noLine, text: `${say.no(mid)}, so the answer is right of mid: lo = ${lo}.`, pointers: { lo, hi }, window: [lo, hi - 1], out: [...mids] });
      }
    }
    steps.push({ line: 5, text: `lo == hi == ${lo}: the range is down to one spot. Return ${lo}.`, pointers: { lo, hi }, current: lo, out: [...mids] });
    return { steps, mids, answer: lo };
  }

  function binarySearchQuestion(exId, size) {
    const n = ARRAY_SIZES[size];

    if (exId === 'probes' || exId === 'state') {
      let a;
      let t;
      let run;
      for (let tries = 0; tries < 100; tries++) {
        a = sortedDistinct(n);
        const absent = exId === 'state' || Math.random() < 0.4;
        t = absent ? pick(ALL_VALUES.filter(v => !a.includes(v) && v > a[0] - 5 && v < a[n - 1] + 5)) : pick(a);
        run = bsSearchRun(a, t);
        if (exId === 'probes' ? run.mids.length >= 2 : run.states.length >= 2) break;
      }
      const label = cellLabel(a);
      if (exId === 'probes') {
        return {
          prompt: `Binary search for <strong>${t}</strong>. Click each index mid lands on, in order.`,
          mode: 'sequence',
          ids: a.map((_, i) => i),
          label,
          svg: (view) => arraySvg(a, view),
          expected: run.mids,
          fixedLength: false,
          code: BS_CODE.search,
          outLabel: 'mid so far',
          steps: () => run.steps,
          right: () => `mid visits ${run.mids.map(m => a[m]).join(', ')}${run.found >= 0 ? ` and finds ${t} at index ${run.found}` : `, then lo passes hi: ${t} is not in the array`}. ${run.mids.length} probes for ${n} values: each one halves the range.`,
          wrong: (k, answer) => {
            if (k === run.mids.length) return `<p class="trace-bad">${said(k)} Then the search stops: ${run.found >= 0 ? `index ${run.found} holds ${t}` : 'lo has passed hi, so the range is empty'}.</p>`;
            const st = run.steps.find(s => s.line === 2 && s.current === run.mids[k]);
            const { lo, hi } = st.pointers;
            const why = k === 0 ? `The range starts as the whole array, 0..${n - 1}.` : `After comparing with ${a[run.mids[k - 1]]}, the range is ${lo}..${hi}.`;
            return `<p class="trace-bad">${said(k)} ${k < answer.length ? `Next mid is ${run.mids[k]}, not ${answer[k]}.` : `Keep going: next mid is ${run.mids[k]}.`}</p>
              <p class="trace-why">${why} mid = (${lo} + ${hi}) // 2 = ${run.mids[k]}, rounding down.</p>`;
          },
        };
      }
      const k = randInt(1, run.states.length - 1);
      const { lo, hi, stepIndex } = run.states[k - 1];
      const steps = run.steps.slice(0, stepIndex);
      return {
        prompt: `Binary search for <strong>${t}</strong>. Stop after <strong>${k} iteration${k === 1 ? '' : 's'}</strong> of the loop. Click where <strong>lo</strong> is, then <strong>hi</strong>.`,
        mode: 'sequence',
        slots: ['lo', 'hi'],
        ids: a.map((_, i) => i),
        label,
        svg: (view) => arraySvg(a, view),
        expected: [lo, hi],
        fixedLength: true,
        code: BS_CODE.search,
        steps: () => steps,
        right: () => `lo = ${lo} and hi = ${hi}: ${t} can only be in ${a.slice(lo, hi + 1).join(', ')}.`,
        wrong: () => {
          const trail = run.mids.slice(0, k).map((m, i) => {
            const s = run.states[i];
            return `mid ${m} (${a[m]}) ${a[m] < t ? `< ${t}, lo = ${s.lo}` : `> ${t}, hi = ${s.hi}`}`;
          });
          return `<p class="trace-bad">lo is ${lo} and hi is ${hi}.</p>
            <p class="trace-why">Each iteration moves one bound past mid: ${trail.join('; ')}. The bound moves to mid ± 1, not to mid, because mid has just been ruled out.</p>`;
        },
      };
    }

    if (exId === 'lower') {
      let a;
      let t;
      for (let tries = 0; tries < 100; tries++) {
        const pool = shuffle(Array.from({ length: 30 }, (_, i) => i + 1)).slice(0, Math.ceil(n / 2));
        a = Array.from({ length: n }, () => pick(pool)).sort((x, y) => x - y);
        const r = Math.random();
        t = r < 0.6 ? pick(a.filter((v, i) => i > 0 && a[i - 1] === v).concat(a)) : r < 0.9 ? randInt(a[0], a[n - 1]) : pick([a[n - 1] + randInt(1, 3), Math.max(0, a[0] - randInt(1, 3))]);
        if (new Set(a).size < n) break;
      }
      const run = halfOpenRun(n, (i) => a[i] >= t, {
        yesLine: 4, noLine: 3,
        yes: (m) => `${a[m]} >= ${t}`,
        no: (m) => `${a[m]} < ${t}`,
      });
      const label = cellLabel(a);
      const ans = run.answer;
      return {
        prompt: `<strong>lowerBound(nums, ${t})</strong> returns the first index whose value is <strong>>= ${t}</strong> (or the end if there is none). Click the index it returns.`,
        mode: 'sequence',
        slots: ['returns'],
        ids: [...a.map((_, i) => i), n],
        label,
        svg: (view) => arraySvg(a, view, { end: true }),
        expected: [ans],
        fixedLength: true,
        code: BS_CODE.lower,
        outLabel: 'mid so far',
        steps: () => run.steps,
        right: () => (ans === n ? `Every value is < ${t}, so it returns ${n}, the end: that is where ${t} would be inserted.` : `Index ${ans} (${a[ans]}). ${a[ans] === t ? `It is the first ${t}, even with copies after it: hi = mid keeps a match and keeps looking left.` : `${t} is not in the array; ${ans} is where it would be inserted.`}`),
        wrong: (k, answer) => `<p class="trace-bad">It returns ${ans === n ? `${n} (the end)` : `${ans} (${a[ans]})`}.</p>
          <p class="trace-why">${answer[0] < ans ? `Index ${answer[0]} holds ${a[answer[0]]}, which is < ${t}.` : answer[0] < n && a[answer[0]] >= t ? `Index ${answer[0]} is >= ${t}, but it is not the first: index ${ans} comes earlier.` : 'Look for the first value that is not smaller than the target.'} When nums[mid] >= target the loop sets hi = mid rather than returning, so it keeps searching left for an earlier match.</p>`,
      };
    }

    // first True in a monotone boolean array
    const firstT = randInt(1, n);
    const ok = Array.from({ length: n }, (_, i) => i >= firstT);
    const cells = ok.map(b => (b ? 'T' : 'F'));
    const run = halfOpenRun(n, (i) => ok[i], {
      yesLine: 3, noLine: 4,
      yes: (m) => `ok[${m}] is True`,
      no: (m) => `ok[${m}] is False`,
    });
    return {
      prompt: `ok is False up to some index and True from there on (like "first bad version"). Binary search for the first True: click each index mid lands on, in order.`,
      mode: 'sequence',
      ids: cells.map((_, i) => i),
      label: (id) => `[${id}] ${cells[id]}`,
      svg: (view) => arraySvg(cells, view, { end: true, clickableIds: cells.map((_, i) => i) }),
      expected: run.mids,
      fixedLength: false,
      code: BS_CODE.first,
      outLabel: 'mid so far',
      steps: () => run.steps,
      right: () => `mid checks ${run.mids.join(', ')} and the search returns ${run.answer}${run.answer === n ? ' (no True at all)' : ''}. Only ${run.mids.length} checks: this is how you search an answer space where you can only ask "is this big enough?".`,
      wrong: (k, answer) => {
        if (k === run.mids.length) return `<p class="trace-bad">${said(k)} Then lo == hi, so the loop ends and returns ${run.answer}.</p>`;
        const st = run.steps.find(s => s.line === 2 && s.current === run.mids[k]);
        const { lo, hi } = st.pointers;
        return `<p class="trace-bad">${said(k)} ${k < answer.length ? `Next mid is ${run.mids[k]}, not ${answer[k]}.` : `Keep going: next mid is ${run.mids[k]}.`}</p>
          <p class="trace-why">The range is lo = ${lo}, hi = ${hi} (hi is exclusive). mid = (${lo} + ${hi}) // 2 = ${run.mids[k]}. A True sets hi = mid (mid could be the first True); a False sets lo = mid + 1.</p>`;
      },
    };
  }

  KITS['binary-search'] = {
    title: 'Binary search',
    exercises: [
      { id: 'probes', name: 'Which indexes does mid hit?' },
      { id: 'state', name: 'Where are lo & hi?' },
      { id: 'lower', name: 'Lower bound' },
      { id: 'first', name: 'First True' },
    ],
    make: binarySearchQuestion,
  };

  // ---------------------------------------------------------------------------
  // Sliding window

  const SW_CODE = {
    fixed: ['s = sum(nums[:k])', 'best = s', 'for r in range(k, len(nums)):', '    s += nums[r] - nums[r - k]   # add new, drop old', '    best = max(best, s)'],
    minlen: ['left = total = 0', 'for right, x in enumerate(nums):', '    total += x', '    while total >= target:', '        best = min(best, right - left + 1)', '        total -= nums[left]', '        left += 1'],
    unique: ['last = {}   # char -> index it was last seen', 'left = 0', 'for right, ch in enumerate(s):', '    if ch in last and last[ch] >= left:', '        left = last[ch] + 1', '    last[ch] = right', '    best = max(best, right - left + 1)'],
  };

  function uniqueRun(s, upTo = s.length - 1) {
    const steps = [];
    const last = {};
    let left = 0;
    let best = 0;
    let bestRange = null;
    for (let right = 0; right <= upTo; right++) {
      const ch = s[right];
      if (ch in last && last[ch] >= left) {
        const old = left;
        left = last[ch] + 1;
        steps.push({ line: 4, text: `'${ch}' is already in the window at index ${last[ch]}: jump left from ${old} to ${left}, just past it.`, pointers: { left, right }, window: [left, right], current: right, heldIds: [last[ch]] });
      }
      last[ch] = right;
      const len = right - left + 1;
      if (len > best) { best = len; bestRange = [left, right]; }
      steps.push({ line: 6, text: `Window ${left}..${right} = '${s.slice(left, right + 1).join('')}', length ${len}.${len === best && bestRange[1] === right ? ' A new longest.' : ` The longest so far is ${best}.`}`, pointers: { left, right }, window: [left, right], current: right });
    }
    return { steps, left, best, bestRange };
  }

  function slidingWindowQuestion(exId, size) {
    const n = ARRAY_SIZES[size];

    if (exId === 'fixed') {
      const k = { small: 3, medium: 3, large: 4 }[size];
      let a;
      let sums;
      for (let tries = 0; tries < 200; tries++) {
        a = Array.from({ length: n }, () => randInt(1, 9));
        sums = [];
        for (let i = 0; i + k <= n; i++) sums.push(a.slice(i, i + k).reduce((x, y) => x + y, 0));
        const max = Math.max(...sums);
        if (sums.filter(v => v === max).length === 1) break;
      }
      const best = sums.indexOf(Math.max(...sums));
      const steps = [{ line: 0, text: `The first window, 0..${k - 1}, sums to ${sums[0]}.`, pointers: { r: k - 1 }, window: [0, k - 1] }];
      for (let r = k; r < n; r++) {
        const i = r - k + 1;
        steps.push({ line: 3, text: `Slide: add nums[${r}] = ${a[r]}, drop nums[${r - k}] = ${a[r - k]}. Sum ${sums[i - 1]} → ${sums[i]}.${sums[i] === Math.max(...sums.slice(0, i + 1)) && sums[i] > Math.max(...sums.slice(0, i)) ? ' A new best.' : ''}`, pointers: { r }, window: [i, r], current: r, heldIds: [r - k] });
      }
      return {
        prompt: `Slide a window of <strong>${k}</strong> across the array. Click the first and last cell of the window with the <strong>largest sum</strong>.`,
        mode: 'sequence',
        slots: ['first', 'last'],
        ids: a.map((_, i) => i),
        label: cellLabel(a),
        svg: (view) => arraySvg(a, view),
        expected: [best, best + k - 1],
        fixedLength: true,
        code: SW_CODE.fixed,
        steps: () => steps,
        right: () => `${best}..${best + k - 1} sums to ${sums[best]}. Each slide is one add and one subtract, so all ${sums.length} windows cost O(n), not O(n·k).`,
        wrong: (m, answer) => `<p class="trace-bad">The best window is ${best}..${best + k - 1}, with sum ${sums[best]}.</p>
          <p class="trace-why">${answer[1] - answer[0] + 1 !== k ? `A window here always has exactly ${k} cells: last = first + ${k - 1}. ` : ''}The window sums, left to right: ${sums.join(', ')}.</p>`,
      };
    }

    if (exId === 'minlen') {
      let a;
      let target;
      let r;
      let steps;
      let left;
      for (let tries = 0; tries < 200; tries++) {
        a = Array.from({ length: n }, () => randInt(1, 9));
        target = randInt(12, 20);
        r = randInt(3, n - 1);
        steps = [];
        left = 0;
        let total = 0;
        let shrank = false;
        for (let right = 0; right <= r; right++) {
          total += a[right];
          steps.push({ line: 2, text: `right = ${right}: add ${a[right]}, total = ${total}.${total >= target ? ` That is >= ${target}.` : ''}`, pointers: { left, right }, window: [left, right], current: right });
          while (total >= target) {
            total -= a[left];
            left++;
            if (right === r) shrank = true;
            steps.push({ line: 6, text: `Record length ${right - left + 2}, then drop nums[${left - 1}] = ${a[left - 1]}: total = ${total}, left = ${left}.${total >= target ? ' Still >= target, keep shrinking.' : ''}`, pointers: { left, right }, window: [left, right], current: right });
          }
        }
        if (shrank) break;
      }
      return {
        prompt: `Shortest subarray with sum <strong>>= ${target}</strong>: grow right, then shrink left while the total is still big enough. When right reaches <strong>index ${r}</strong> (ringed) and the while loop ends, where is <strong>left</strong>?`,
        mode: 'sequence',
        slots: ['left'],
        ids: a.map((_, i) => i),
        label: cellLabel(a),
        svg: (view) => arraySvg(a, view, { mark: r }),
        expected: [left],
        fixedLength: true,
        code: SW_CODE.minlen,
        steps: () => steps,
        right: () => `left = ${left}: the window ${left}..${r} sums to ${a.slice(left, r + 1).reduce((x, y) => x + y, 0)}, below ${target} again. Removing one more would only shrink it further, and putting nums[${left - 1}] back would make it >= ${target}.`,
        wrong: () => `<p class="trace-bad">left ends at ${left}.</p>
          <p class="trace-why">The loop is a <em>while</em>, not an if: it keeps dropping nums[left] until the total falls below ${target}. ${a.slice(left - 1, r + 1).reduce((x, y) => x + y, 0)} (from ${left - 1}) is still >= ${target}, and ${a.slice(left, r + 1).reduce((x, y) => x + y, 0)} (from ${left}) is not. left never moves backwards, which is why the whole run is O(n).</p>`,
      };
    }

    // unique / longest: longest substring without repeats.
    const alpha = size === 'small' ? 'abcd' : 'abcde';
    let s;
    let run;
    let r;
    for (let tries = 0; tries < 200; tries++) {
      s = Array.from({ length: n }, () => pick([...alpha]));
      r = randInt(3, n - 1);
      run = uniqueRun(s, exId === 'uniqueLeft' ? r : n - 1);
      if (exId === 'uniqueLeft' ? run.left > 0 : run.best >= 3) break;
    }
    if (exId === 'uniqueLeft') {
      return {
        prompt: `Longest substring without repeats. When <strong>right reaches index ${r}</strong> (ringed), where is <strong>left</strong> after the update?`,
        mode: 'sequence',
        slots: ['left'],
        ids: s.map((_, i) => i),
        label: cellLabel(s),
        svg: (view) => arraySvg(s, view, { mark: r }),
        expected: [run.left],
        fixedLength: true,
        code: SW_CODE.unique,
        steps: () => run.steps,
        right: () => `left = ${run.left}: '${s.slice(run.left, r + 1).join('')}' has no repeats, and starting any earlier would include a letter twice.`,
        wrong: () => `<p class="trace-bad">left is at ${run.left}.</p>
          <p class="trace-why">left only moves when the new letter is already inside the window, and then it jumps just past that letter's last index. It never moves backwards, even when the letter's earlier copy is left of the window. The window ${run.left}..${r} is '${s.slice(run.left, r + 1).join('')}'.</p>`,
      };
    }
    const [b0, b1] = run.bestRange;
    return {
      prompt: 'Longest substring without repeats. Click the first and last letter of the window the algorithm records as longest (the first one it finds, if there is a tie).',
      mode: 'sequence',
      slots: ['first', 'last'],
      ids: s.map((_, i) => i),
      label: cellLabel(s),
      svg: (view) => arraySvg(s, view),
      expected: [b0, b1],
      fixedLength: true,
      code: SW_CODE.unique,
      steps: () => run.steps,
      right: () => `'${s.slice(b0, b1 + 1).join('')}' at ${b0}..${b1}, length ${run.best}.`,
      wrong: (m, answer) => {
        const pickd = s.slice(answer[0], answer[1] + 1);
        const dup = pickd.find((c, i) => pickd.indexOf(c) !== i);
        return `<p class="trace-bad">It records ${b0}..${b1}, '${s.slice(b0, b1 + 1).join('')}' (length ${run.best}).</p>
          <p class="trace-why">${answer[1] < answer[0] ? 'The first letter comes before the last.' : dup ? `Your window has '${dup}' twice.` : `Yours has length ${answer[1] - answer[0] + 1}${answer[1] - answer[0] + 1 === run.best ? ', the same, but the algorithm only replaces the best on a strictly longer window, so the earlier one stays' : ''}.`}</p>`;
      },
    };
  }

  KITS['sliding-window'] = {
    title: 'Sliding window',
    exercises: [
      { id: 'fixed', name: 'Fixed window: best sum' },
      { id: 'minlen', name: 'Shrink: where is left?' },
      { id: 'uniqueLeft', name: 'No repeats: where is left?' },
      { id: 'longest', name: 'No repeats: longest window' },
    ],
    make: slidingWindowQuestion,
  };

  // ---------------------------------------------------------------------------
  // Two pointers

  const TP_CODE = {
    pair: ['l, r = 0, len(nums) - 1', 'while l < r:', '    s = nums[l] + nums[r]', '    if s == target: return [l, r]', '    if s < target: l += 1', '    else: r -= 1'],
    dedupe: ['w = 1   # next slot to write', 'for r in range(1, len(nums)):', '    if nums[r] != nums[w - 1]:', '        nums[w] = nums[r]', '        w += 1', 'return w'],
    water: ['l, r = 0, len(h) - 1', 'while l < r:', '    best = max(best, min(h[l], h[r]) * (r - l))', '    if h[l] < h[r]: l += 1', '    else: r -= 1'],
    palindrome: ['l, r = 0, len(s) - 1', 'while l < r:', '    if s[l] != s[r]: return False', '    l, r = l + 1, r - 1', 'return True'],
  };

  function twoPointersQuestion(exId, size) {
    const n = ARRAY_SIZES[size];

    if (exId === 'pair' || exId === 'water') {
      let a;
      let target;
      let states;
      let steps;
      for (let tries = 0; tries < 200; tries++) {
        a = exId === 'pair' ? sortedDistinct(n) : Array.from({ length: n }, () => randInt(1, 9));
        let l = 0;
        let r = n - 1;
        if (exId === 'pair') {
          const i = randInt(0, n - 2);
          target = a[i] + a[randInt(i + 1, n - 1)];
        }
        states = [];
        steps = [{ line: 0, text: 'l starts at the left end, r at the right end.', pointers: { l, r }, window: [l, r] }];
        while (l < r) {
          if (exId === 'pair') {
            const sum = a[l] + a[r];
            if (sum === target) { steps.push({ line: 3, text: `${a[l]} + ${a[r]} = ${target}: found.`, pointers: { l, r }, window: [l, r] }); break; }
            if (sum < target) { l++; steps.push({ line: 4, text: `${a[l - 1]} + ${a[r]} = ${sum} < ${target}: the sum must grow, and only moving l right can do that. l = ${l}.`, pointers: { l, r }, window: [l, r], heldIds: [l - 1] }); }
            else { r--; steps.push({ line: 5, text: `${a[l]} + ${a[r + 1]} = ${sum} > ${target}: the sum must shrink, so r moves left. r = ${r}.`, pointers: { l, r }, window: [l, r], heldIds: [r + 1] }); }
          } else {
            const area = Math.min(a[l], a[r]) * (r - l);
            if (a[l] < a[r]) { l++; steps.push({ line: 3, text: `Area min(${a[l - 1]}, ${a[r]}) × ${r - l + 1} = ${area}. The left line is shorter, and no wider container is left for it: move l. l = ${l}.`, pointers: { l, r }, window: [l, r], heldIds: [l - 1] }); }
            else { r--; steps.push({ line: 4, text: `Area min(${a[l]}, ${a[r + 1]}) × ${r - l + 1} = ${area}. The right line is ${a[r + 1] < a[l] ? 'shorter' : 'as short (a tie moves r)'}: move r. r = ${r}.`, pointers: { l, r }, window: [l, r], heldIds: [r + 1] }); }
          }
          states.push({ l, r, at: steps.length });
        }
        if (states.length >= 3) break;
      }
      const k = randInt(1, states.length - 1);
      const st = states[k - 1];
      const shown = steps.slice(0, st.at);
      const moves = shown.slice(1).map(s => s.text.split(':')[0]);
      return {
        prompt: exId === 'pair'
          ? `Sorted array, target sum <strong>${target}</strong>. After <strong>${k} move${k === 1 ? '' : 's'}</strong>, where are <strong>l</strong> and <strong>r</strong>? Click l, then r.`
          : `Container with most water (the numbers are line heights). After <strong>${k} move${k === 1 ? '' : 's'}</strong>, where are <strong>l</strong> and <strong>r</strong>? Click l, then r. On a tie, r moves.`,
        mode: 'sequence',
        slots: ['l', 'r'],
        ids: a.map((_, i) => i),
        label: cellLabel(a),
        svg: (view) => arraySvg(a, view),
        expected: [st.l, st.r],
        fixedLength: true,
        code: exId === 'pair' ? TP_CODE.pair : TP_CODE.water,
        steps: () => shown,
        right: () => `l = ${st.l}, r = ${st.r}. ${exId === 'pair' ? 'Each move rules out one value for good: everything left of l is too small for any partner, everything right of r too big.' : 'The shorter line moves because keeping it can never beat the area just measured: any narrower container is at most as tall.'}`,
        wrong: () => `<p class="trace-bad">l = ${st.l} and r = ${st.r}.</p>
          <p class="trace-why">Exactly one pointer moves per step. The moves: ${moves.map(m => esc(m)).join('; ')}.</p>`,
      };
    }

    if (exId === 'dedupe') {
      let a;
      let R;
      for (let tries = 0; tries < 100; tries++) {
        const pool = sortedDistinct(Math.ceil(n / 2));
        a = Array.from({ length: n }, () => pick(pool)).sort((x, y) => x - y);
        R = randInt(2, n - 2);
        if (new Set(a.slice(0, R + 1)).size < R + 1 && new Set(a.slice(0, R + 1)).size > 1) break;
      }
      const vals = [...a];
      const steps = [{ line: 0, text: 'nums[0] is always kept, so w starts at 1.', pointers: { w: 1 }, vals: [...vals], window: [0, 0] }];
      let w = 1;
      for (let r = 1; r <= R; r++) {
        if (vals[r] !== vals[w - 1]) {
          vals[w] = vals[r];
          w++;
          steps.push({ line: 4, text: `nums[${r}] = ${vals[r]} differs from the last kept value: copy it to index ${w - 1}, then w = ${w}.`, pointers: { r, w }, vals: [...vals], window: [0, w - 1], current: r });
        } else {
          steps.push({ line: 2, text: `nums[${r}] = ${vals[r]} repeats the last kept value (${vals[w - 1]}): skip it. w stays ${w}.`, pointers: { r, w }, vals: [...vals], window: [0, w - 1], current: r });
        }
      }
      const kept = [...new Set(a.slice(0, R + 1))];
      return {
        prompt: `Remove duplicates in place: r reads every value, w is the next slot to write a new value to. After the loop handles <strong>r = ${R}</strong> (ringed), where is <strong>w</strong>?`,
        mode: 'sequence',
        slots: ['w'],
        ids: a.map((_, i) => i),
        label: (id) => `index ${id}`,
        svg: (view) => arraySvg(a, view, { mark: R }),
        expected: [w],
        fixedLength: true,
        code: TP_CODE.dedupe,
        steps: () => steps,
        right: () => `w = ${w}: indexes 0..${w - 1} now hold the ${kept.length} distinct values read so far (${kept.join(', ')}).`,
        wrong: () => `<p class="trace-bad">w is ${w}.</p>
          <p class="trace-why">w only moves when r finds a new value, so it always equals the number of distinct values seen. Up to index ${R} those are ${kept.join(', ')}: ${kept.length} of them.</p>`,
      };
    }

    // palindrome: where does the check fail?
    const alpha = 'abcdefgh';
    const half = Array.from({ length: n >> 1 }, () => pick([...alpha]));
    const s = [...half, ...(n % 2 ? [pick([...alpha])] : []), ...[...half].reverse()];
    const p = randInt(0, (n >> 1) - 1);
    s[p] = pick([...alpha].filter(c => c !== s[n - 1 - p]));
    const steps = [{ line: 0, text: 'l at the start, r at the end.', pointers: { l: 0, r: n - 1 }, window: [0, n - 1] }];
    for (let l = 0; l <= p; l++) {
      const r = n - 1 - l;
      steps.push(l === p
        ? { line: 2, text: `'${s[l]}' != '${s[r]}': return False.`, pointers: { l, r }, window: [l, r], current: l, heldIds: [r] }
        : { line: 3, text: `'${s[l]}' == '${s[r]}': move both inwards.`, pointers: { l, r }, window: [l, r], current: l, heldIds: [r] });
    }
    return {
      prompt: 'Palindrome check with pointers from both ends. This string is not a palindrome. Click where <strong>l</strong> and <strong>r</strong> are when it returns False.',
      mode: 'sequence',
      slots: ['l', 'r'],
      ids: s.map((_, i) => i),
      label: cellLabel(s),
      svg: (view) => arraySvg(s, view),
      expected: [p, n - 1 - p],
      fixedLength: true,
      code: TP_CODE.palindrome,
      steps: () => steps,
      right: () => `l = ${p}, r = ${n - 1 - p}: '${s[p]}' vs '${s[n - 1 - p]}'. ${p ? `The ${p} pair${p === 1 ? '' : 's'} outside matched.` : 'The very first pair differs.'}`,
      wrong: (m, answer) => `<p class="trace-bad">l = ${p} and r = ${n - 1 - p}.</p>
        <p class="trace-why">${answer.length === 2 && answer[0] + answer[1] !== n - 1 ? `l and r always mirror each other: l + r = ${n - 1}. ` : ''}The pointers move inwards together and stop at the first pair that differs, counting from the outside.</p>`,
    };
  }

  KITS['two-pointers'] = {
    title: 'Two pointers',
    exercises: [
      { id: 'pair', name: 'Pair sum: where are l & r?' },
      { id: 'water', name: 'Container: where are l & r?' },
      { id: 'dedupe', name: 'Read/write: where is w?' },
      { id: 'palindrome', name: 'Palindrome: where does it fail?' },
    ],
    make: twoPointersQuestion,
  };

  // ===========================================================================
  // Backtracking
  // ===========================================================================

  // setAnswer builds validate() for questions whose answer is a set of ids.
  const sameSet = (answer, expected) => answer.length === expected.length && expected.every(id => answer.includes(id));
  const fmtPath = (p) => `[${p.join(', ')}]`;

  const BT_CODE = {
    subsets: ['def backtrack(start, path):', '    out.append(path[:])   # every node is a subset', '    for i in range(start, len(nums)):', '        path.append(nums[i])', '        backtrack(i + 1, path)', '        path.pop()         # undo the choice'],
    combine: ['def backtrack(start, path):', '    if len(path) == k:', '        out.append(path[:])', '        return', '    for x in range(start, n + 1):', '        path.append(x)', '        backtrack(x + 1, path)', '        path.pop()'],
    permute: ['def backtrack(path):', '    if len(path) == len(items):', '        out.append(path[:])', '        return', '    for i, x in enumerate(items):', '        if used[i]: continue', '        used[i] = True; path.append(x)', '        backtrack(path)', '        used[i] = False; path.pop()'],
    queens: ['def place(row):', '    if row == n: return True', '    for col in range(n):', '        if safe(row, col):', '            queens[row] = col', '            if place(row + 1): return True', '            queens[row] = None    # dead end: undo', '    return False'],
  };

  // btRun runs a generator and records a replay up to the m-th recorded result.
  function btRun(kind, items, k, m) {
    const steps = [];
    const out = [];
    const path = [];
    const used = items.map(() => false);
    const lab = (ids) => fmtPath(ids.map(i => items[i]));
    const snap = (line, text) => steps.push({ line, text, heldIds: [...path], out: [...path] });
    const record = (line) => {
      out.push([...path]);
      snap(line, `Record #${out.length}: ${lab(path)}.`);
      return out.length >= m;
    };
    const go = (start) => {
      if (kind === 'subsets' && record(1)) return true;
      if (kind === 'combine' && path.length === k) return record(2);
      if (kind === 'permute' && path.length === items.length) return record(2);
      for (let i = kind === 'permute' ? 0 : start; i < items.length; i++) {
        if (kind === 'permute' && used[i]) continue;
        used[i] = true;
        path.push(i);
        snap(kind === 'permute' ? 6 : kind === 'combine' ? 5 : 3, `Choose ${items[i]}: path is ${lab(path)}.`);
        if (go(i + 1)) return true;
        path.pop();
        used[i] = false;
        snap(kind === 'permute' ? 8 : kind === 'combine' ? 7 : 5, `Back from ${items[i]}: undo it, path is ${lab(path)}.`);
      }
      return false;
    };
    go(0);
    return { steps, out };
  }

  function queensRun(n, c0) {
    const queens = [c0];
    const steps = [{ line: 0, text: `Row 0 is fixed: its queen is in column ${c0}.`, queens: [0 * n + c0] }];
    const ids = () => queens.map((c, r) => r * n + c);
    const safe = (r, c) => queens.every((qc, qr) => qc !== c && Math.abs(qc - c) !== r - qr);
    const go = (r) => {
      if (r === n) return true;
      const attacked = [];
      for (let c = 0; c < n; c++) {
        if (!safe(r, c)) { attacked.push(c); continue; }
        queens[r] = c;
        steps.push({ line: 4, text: `Row ${r}: ${attacked.length ? `column${attacked.length > 1 ? 's' : ''} ${attacked.join(', ')} ${attacked.length > 1 ? 'are' : 'is'} attacked; ` : ''}place a queen in column ${c}.`, queens: ids(), current: r * n + c });
        attacked.length = 0;
        if (go(r + 1)) return true;
        queens.length = r;
        steps.push({ line: 6, text: `Nothing works below row ${r}: take the queen off column ${c} and try further right.`, queens: ids(), heldIds: [r * n + c] });
      }
      steps.push({ line: 7, text: `Row ${r}: ${attacked.length === n ? 'every column is attacked' : 'no column left to try'}. Go back up a row.`, queens: ids() });
      return false;
    };
    const ok = go(1);
    return { ok, steps, queens: [...queens] };
  }

  function gridSvg(n, view, { fixed = [] } = {}) {
    const size = 46;
    const pad = 20;
    const queens = new Set(view.step ? view.step.queens : fixed);
    let body = '';
    for (let r = 0; r < n; r++) {
      for (let c = 0; c < n; c++) {
        const id = r * n + c;
        const cls = ['tnode', 'acell', 'gcell', (r + c) % 2 ? 'dark' : 'light'];
        const idx = view.picked.get(id);
        if (idx) cls.push('picked');
        if (view.current === id) cls.push('current');
        if (view.held.has(id)) cls.push('held');
        if (view.wrong === id) cls.push('wrong');
        const clickable = view.clickable && r > 0 && !idx;
        if (clickable) cls.push('clickable');
        const x = pad + c * size;
        const y = pad + r * size;
        body += `<g class="${cls.join(' ')}" data-id="${id}" ${clickable ? `role="button" tabindex="0" aria-label="Row ${r}, column ${c}"` : ''}>
          <rect x="${x}" y="${y}" width="${size}" height="${size}" rx="2" />
          <text x="${x + size / 2}" y="${y + size / 2}" dy="0.35em">${queens.has(id) ? '♛' : ''}</text>
          ${idx ? `<text class="badge" x="${x + size - 12}" y="${y + 13}">${idx}</text>` : ''}</g>`;
      }
    }
    return svgWrap(n * size + 2 * pad, n * size + 2 * pad, body, `${n} by ${n} board`);
  }

  function backtrackingQuestion(exId, size) {
    if (exId === 'queens') {
      const n = { small: 4, medium: 5, large: 6 }[size];
      let run;
      let c0;
      for (const c of shuffle(Array.from({ length: n }, (_, i) => i))) {
        run = queensRun(n, c);
        c0 = c;
        if (run.ok) break;
      }
      const expected = run.queens.slice(1).map((c, i) => (i + 1) * n + c);
      return {
        prompt: `<strong>${n}-queens</strong>: no two queens may share a row, column or diagonal. Row 0's queen is fixed in column ${c0}. Each row tries columns left to right and backtracks on a dead end. Click the first solution it finds, one queen per row, top to bottom.`,
        mode: 'sequence',
        slots: Array.from({ length: n - 1 }, (_, i) => `row ${i + 1}`),
        ids: Array.from({ length: n * n }, (_, i) => i).filter(i => i >= n),
        label: (id) => `col ${id % n}`,
        svg: (view) => gridSvg(n, view, { fixed: [c0] }),
        expected,
        fixedLength: true,
        code: BT_CODE.queens,
        steps: () => run.steps,
        right: () => `Columns ${run.queens.join(', ')} by row. ${run.steps.filter(s => s.line === 6).length} dead end${run.steps.filter(s => s.line === 6).length === 1 ? '' : 's'} were undone on the way.`,
        wrong: (k, answer) => {
          const r = k + 1;
          const c = answer[k] % n;
          const prev = run.queens.slice(0, r);
          const clash = prev.findIndex((qc, qr) => qc === c || Math.abs(qc - c) === r - qr);
          return `<p class="trace-bad">${k ? `Rows 1 to ${k} are right. ` : ''}Row ${r} gets column ${run.queens[r]}, not ${c}.</p>
            <p class="trace-why">${clash >= 0 ? `Column ${c} is attacked by the queen in row ${clash} (column ${prev[clash]}).` : `Column ${c} is safe from the rows above, but ${c < run.queens[r] ? 'it leads to a dead end further down, so the search takes it back and tries further right' : `the search tries columns left to right and column ${run.queens[r]} comes first and works`}.`} Use "Show me" to watch the dead ends being undone.</p>`;
        },
      };
    }

    const kind = exId;
    let items;
    let k = 0;
    if (kind === 'combine') {
      const n = { small: 4, medium: 5, large: 6 }[size];
      items = Array.from({ length: n }, (_, i) => i + 1);
      k = { small: 2, medium: 3, large: 3 }[size];
    } else if (kind === 'permute') {
      const n = { small: 3, medium: 4, large: 4 }[size];
      items = shuffle([...'ABCDEFG']).slice(0, n).sort();
    } else {
      const n = { small: 3, medium: 4, large: 5 }[size];
      items = distinctValues(n, 9).sort((x, y) => x - y);
    }
    const full = btRun(kind, items, k, Infinity).out;
    const m = randInt(kind === 'subsets' ? 3 : 2, Math.min(full.length, kind === 'subsets' ? 12 : 9));
    const run = btRun(kind, items, k, m);
    const target = run.out[m - 1];
    const ord = (x) => `${x}${['th', 'st', 'nd', 'rd'][(x % 100 > 10 && x % 100 < 14) || x % 10 > 3 ? 0 : x % 10]}`;
    const what = { subsets: 'subset', combine: 'combination', permute: 'permutation' }[kind];
    const call = { subsets: `subsets(${fmtPath(items)})`, combine: `combine(n=${items.length}, k=${k})`, permute: `permute(${fmtPath(items)})` }[kind];
    const listed = run.out.map((p, i) => `${i + 1}. ${fmtPath(p.map(j => items[j]))}`).join('  ');
    return {
      prompt: `<strong>${call}</strong> records results in the order the code below reaches them. Click the <strong>${ord(m)} ${what}</strong> it records, element by element.${kind === 'subsets' ? ' (The 1st is the empty subset.)' : ''}`,
      mode: 'sequence',
      ids: items.map((_, i) => i),
      label: (id) => String(items[id]),
      svg: (view) => arraySvg(items, view),
      expected: target,
      fixedLength: kind !== 'subsets',
      code: BT_CODE[kind],
      outLabel: 'path',
      steps: () => run.steps,
      right: () => `${fmtPath(target.map(j => items[j]))}. The order comes straight from the loop: smallest choice first, go as deep as possible, undo, try the next.`,
      wrong: () => `<p class="trace-bad">The ${ord(m)} is ${fmtPath(target.map(j => items[j]))}.</p>
        <p class="trace-why">In order: ${esc(listed)}. ${kind === 'permute' ? 'used[] lets every unused item be chosen at each depth.' : 'The start index means each item is only followed by later items, so no result is produced twice.'}</p>`,
    };
  }

  KITS.backtracking = {
    title: 'Backtracking',
    exercises: [
      { id: 'subsets', name: 'Subsets: the m-th recorded' },
      { id: 'combine', name: 'Combinations: the m-th' },
      { id: 'permute', name: 'Permutations: the m-th' },
      { id: 'queens', name: 'N-queens: first solution' },
    ],
    make: backtrackingQuestion,
  };

  // ===========================================================================
  // 1-D dynamic programming
  // ===========================================================================

  const DP_CODE = {
    rob: ['dp[0] = nums[0]', 'dp[1] = max(nums[0], nums[1])', 'for i in range(2, n):', '    dp[i] = max(dp[i - 1],            # skip house i', '                dp[i - 2] + nums[i])  # rob house i'],
    kadane: ['cur = best = nums[0]', 'for x in nums[1:]:', '    cur = max(x, cur + x)   # best sum ending here', '    best = max(best, cur)'],
    coins: ['ways = [1] + [0] * amount', 'for c in coins:                  # coins outside:', '    for a in range(c, amount + 1):  # each set counted once', '        ways[a] += ways[a - c]'],
    lis: ['for i in range(n):', '    dp[i] = 1', '    for j in range(i):', '        if nums[j] < nums[i]:', '            dp[i] = max(dp[i], dp[j] + 1)'],
  };

  function dpQuestion(exId, size) {
    const n = { small: 6, medium: 8, large: 10 }[size];
    const dpHeld = (dp, upTo) => dp.slice(0, upTo + 1).map((v, i) => `${i}: ${v}`);

    if (exId === 'rob') {
      let a;
      let best;
      let pickSet;
      for (let tries = 0; tries < 300; tries++) {
        a = Array.from({ length: n }, () => randInt(1, 20));
        best = -1;
        let count = 0;
        for (let mask = 0; mask < 1 << n; mask++) {
          if (mask & (mask << 1)) continue;
          const s = a.reduce((acc, v, i) => acc + (mask >> i & 1 ? v : 0), 0);
          if (s > best) { best = s; count = 1; pickSet = mask; } else if (s === best) count++;
        }
        if (count === 1) break;
      }
      const expected = a.map((_, i) => i).filter(i => pickSet >> i & 1);
      const dp = [];
      const steps = [];
      a.forEach((v, i) => {
        if (i === 0) { dp.push(v); steps.push({ line: 0, text: `One house: rob it. dp[0] = ${v}.`, current: 0, heldLabel: 'dp', held: dpHeld(dp, 0) }); return; }
        if (i === 1) { dp.push(Math.max(a[0], v)); steps.push({ line: 1, text: `Two houses: take the richer. dp[1] = max(${a[0]}, ${v}) = ${dp[1]}.`, current: 1, heldIds: [0], heldLabel: 'dp', held: dpHeld(dp, 1) }); return; }
        const skip = dp[i - 1];
        const rob = dp[i - 2] + v;
        dp.push(Math.max(skip, rob));
        steps.push({ line: 3, text: `dp[${i}] = max(skip: dp[${i - 1}] = ${skip}, rob: dp[${i - 2}] + ${v} = ${rob}) = ${dp[i]}${rob > skip ? ', robbing' : skip > rob ? ', skipping' : ''}.`, current: i, heldIds: [i - 1, i - 2], heldLabel: 'dp', held: dpHeld(dp, i) });
      });
      steps.push({ line: 4, text: `Best total: dp[${n - 1}] = ${best}. Walking back through the choices gives houses ${expected.join(', ')}.`, out: expected, heldLabel: 'dp', held: dpHeld(dp, n - 1) });
      return {
        prompt: 'House robber: you cannot rob two houses next to each other. Click the houses in the <strong>best</strong> plan (any order), then Check.',
        mode: 'sequence',
        ids: a.map((_, i) => i),
        label: cellLabel(a),
        svg: (view) => arraySvg(a, view),
        expected,
        fixedLength: false,
        validate: (ans) => ({ ok: sameSet(ans, expected), k: 0 }),
        code: DP_CODE.rob,
        outLabel: 'Robbed',
        steps: () => steps,
        right: () => `Houses ${expected.join(', ')} for ${best}. dp[i] is the best total from the first i + 1 houses; each house either is skipped (dp[i - 1]) or robbed on top of dp[i - 2].`,
        wrong: (k, ans) => {
          const sorted = [...ans].sort((x, y) => x - y);
          const adj = sorted.find((x, i) => i && sorted[i - 1] === x - 1);
          const total = ans.reduce((s, i) => s + a[i], 0);
          return `<p class="trace-bad">The best plan is houses ${expected.join(', ')}, total ${best}.</p>
            <p class="trace-why">${adj !== undefined ? `Houses ${adj - 1} and ${adj} are next to each other, which is not allowed.` : `Yours totals ${total}.`} Greedy picks (the biggest house first, or every other house) can miss it; the table compares both choices at every house. dp: ${dp.join(', ')}.</p>`;
        },
      };
    }

    if (exId === 'kadane') {
      let a;
      let best;
      let range;
      for (let tries = 0; tries < 300; tries++) {
        a = Array.from({ length: n }, () => randInt(-9, 9));
        best = -Infinity;
        let count = 0;
        for (let i = 0; i < n; i++) for (let j = i; j < n; j++) {
          const s = a.slice(i, j + 1).reduce((x, y) => x + y, 0);
          if (s > best) { best = s; count = 1; range = [i, j]; } else if (s === best) count++;
        }
        if (count === 1 && best > 0 && range[1] > range[0] && a.some(v => v < 0)) break;
      }
      const steps = [];
      let cur = a[0];
      let bst = a[0];
      let start = 0;
      steps.push({ line: 0, text: `Start: cur = best = ${a[0]}.`, current: 0, window: [0, 0], heldLabel: 'cur, best', held: [cur, bst] });
      for (let i = 1; i < n; i++) {
        const x = a[i];
        const fresh = x > cur + x;
        const prevCur = cur;
        cur = Math.max(x, cur + x);
        if (fresh) start = i;
        bst = Math.max(bst, cur);
        steps.push({ line: 2, text: fresh ? `${x} on its own beats ${prevCur} + ${x}: start a new run here. cur = ${cur}.` : `Extend the run: cur = ${prevCur} + ${x} = ${cur}.${cur === bst ? ` best = ${bst}.` : ''}`, current: i, window: [start, i], heldLabel: 'cur, best', held: [cur, bst] });
      }
      return {
        prompt: 'Maximum subarray (Kadane): cur is the best sum of a run ending at the current index. Click the first and last cell of the subarray with the <strong>largest sum</strong>.',
        mode: 'sequence',
        slots: ['first', 'last'],
        ids: a.map((_, i) => i),
        label: cellLabel(a),
        svg: (view) => arraySvg(a, view),
        expected: range,
        fixedLength: true,
        code: DP_CODE.kadane,
        steps: () => steps,
        right: () => `${range[0]}..${range[1]} sums to ${best}. The only decision at each index: extend the run, or drop it because it is dragging the sum below starting fresh.`,
        wrong: (k, ans) => `<p class="trace-bad">The best is ${range[0]}..${range[1]}, sum ${best}.</p>
          <p class="trace-why">${ans[1] >= ans[0] ? `Yours sums to ${a.slice(ans[0], ans[1] + 1).reduce((x, y) => x + y, 0)}. ` : ''}A run is worth keeping while its sum stays positive; once it goes negative, starting fresh is better.</p>`,
      };
    }

    if (exId === 'coins') {
      let coins;
      let amount;
      let combos;
      let perms;
      for (let tries = 0; tries < 300; tries++) {
        coins = shuffle([1, 2, 3, 5]).slice(0, size === 'small' ? 2 : 3).sort((x, y) => x - y);
        amount = randInt(size === 'small' ? 4 : 6, size === 'large' ? 11 : 8);
        const w = [1, ...Array(amount).fill(0)];
        coins.forEach(c => { for (let x = c; x <= amount; x++) w[x] += w[x - c]; });
        combos = w[amount];
        const p = [1, ...Array(amount).fill(0)];
        for (let x = 1; x <= amount; x++) coins.forEach(c => { if (c <= x) p[x] += p[x - c]; });
        perms = p[amount];
        if (combos !== perms && combos >= 2) break;
      }
      const table = [];
      const w = [1, ...Array(amount).fill(0)];
      coins.forEach(c => { for (let x = c; x <= amount; x++) w[x] += w[x - c]; table.push({ c, row: [...w] }); });
      const extras = shuffle([combos + 1, combos - 1, perms + 1, combos + 2].filter(v => v > 0 && v !== combos && v !== perms));
      const opts = [...new Set([combos, perms, ...extras])].slice(0, 4).sort((x, y) => x - y);
      return {
        prompt: `How many ways can you make <strong>${amount}</strong> from coins <strong>${coins.join(', ')}</strong> (unlimited of each, order does not matter: 1 + 2 and 2 + 1 are the same way)?`,
        mode: 'choice',
        svg: (view) => arraySvg(coins, view),
        sequence: [`amount ${amount}`, ...coins.map(c => `coin ${c}`)],
        choices: opts.map(v => ({ id: String(v), label: String(v) })),
        answerId: String(combos),
        explain: (picked) => `<p class="${picked === String(combos) ? 'trace-ok' : 'trace-bad'}"><strong>${picked === String(combos) ? 'Right' : 'Not quite'}:</strong> ${combos}.${picked === String(perms) ? ` ${perms} counts orderings too: that is what you get with the amount loop outside and the coin loop inside.` : ''}</p>
          <p class="trace-why">With the coin loop outside, each way is built in coin order, so it is counted once. ways[a] after each coin:</p>
          <table class="trace-orders">${table.map(t => `<tr><th>with ${t.c}</th><td>${t.row.join(' ')}</td></tr>`).join('')}</table>`,
      };
    }

    // lis: any longest strictly increasing subsequence
    let a;
    let dp;
    let L;
    for (let tries = 0; tries < 300; tries++) {
      a = distinctValues(n, 30);
      dp = a.map(() => 1);
      for (let i = 0; i < n; i++) for (let j = 0; j < i; j++) if (a[j] < a[i]) dp[i] = Math.max(dp[i], dp[j] + 1);
      L = Math.max(...dp);
      if (L >= 3 && L < n - 1) break;
    }
    const steps = a.map((v, i) => {
      const from = a.map((_, j) => j).filter(j => j < i && a[j] < v);
      const bestJ = from.length ? from.reduce((b, j) => (dp[j] > dp[b] ? j : b), from[0]) : -1;
      return { line: 4, text: bestJ < 0 ? `dp[${i}] = 1: nothing before ${v} is smaller.` : `dp[${i}] = dp[${bestJ}] + 1 = ${dp[i]}: the longest run ending in a smaller value is the one ending at ${a[bestJ]}.`, current: i, heldIds: from, heldLabel: 'dp', held: dpHeld(dp, i) };
    });
    let i = dp.indexOf(L);
    const one = [i];
    while (dp[i] > 1) { const j = a.findIndex((v, jj) => jj < i && v < a[i] && dp[jj] === dp[i] - 1); one.unshift(j); i = j; }
    return {
      prompt: 'Click a <strong>longest strictly increasing subsequence</strong>, left to right. Any one of the longest counts.',
      mode: 'sequence',
      ids: a.map((_, j) => j),
      label: cellLabel(a),
      svg: (view) => arraySvg(a, view),
      expected: one,
      fixedLength: false,
      validate: (ans) => {
        const inc = ans.every((id, j) => j === 0 || (id > ans[j - 1] && a[id] > a[ans[j - 1]]));
        let k = 1;
        while (k < ans.length && ans[k] > ans[k - 1] && a[ans[k]] > a[ans[k - 1]]) k++;
        return { ok: inc && ans.length === L, k: inc ? ans.length : k };
      },
      code: DP_CODE.lis,
      steps: () => steps,
      right: () => `Length ${L}. dp[i] is the longest increasing subsequence ending at i: ${dp.join(', ')}.`,
      wrong: (k, ans) => {
        const inc = k === 0 || k >= ans.length;
        return `<p class="trace-bad">${inc ? `Yours is increasing but has length ${ans.length}; the longest has ${L}, for example ${one.map(j => a[j]).join(', ')}.` : `${a[ans[k]]} at index ${ans[k]} does not continue an increasing sequence: it must be to the right of and bigger than ${a[ans[k - 1]]}.`}</p>
          <p class="trace-why">dp[i] = 1 + the best dp[j] for j < i with a smaller value. dp: ${dp.join(', ')}.</p>`;
      },
    };
  }

  KITS['1d-dp'] = {
    title: '1-D dynamic programming',
    exercises: [
      { id: 'rob', name: 'House robber: which houses?' },
      { id: 'kadane', name: 'Maximum subarray' },
      { id: 'coins', name: 'Coin change: how many ways?' },
      { id: 'lis', name: 'Longest increasing subsequence' },
    ],
    make: dpQuestion,
  };

  // ===========================================================================
  // Intervals
  // ===========================================================================

  // intervalSvg draws one bar per interval over a number line. step.lines
  // draws vertical markers (a sweep position or arrows).
  function intervalSvg(iv, view, { mark = null, lines = [] } = {}) {
    const unit = 24;
    const x0 = 30;
    const rowH = 30;
    const top = 16;
    const maxX = Math.max(...iv.map(([, e]) => e)) + 1;
    const px = (v) => x0 + v * unit;
    const h = top + iv.length * rowH + 30;
    let body = '';
    for (let v = 0; v <= maxX; v++) {
      body += `<line class="grid" x1="${px(v)}" y1="${top - 6}" x2="${px(v)}" y2="${h - 24}" />`;
      if (v % 2 === 0) body += `<text class="idx" x="${px(v)}" y="${h - 8}">${v}</text>`;
    }
    (view.step?.lines || lines).forEach(x => { body += `<line class="sweep" x1="${px(x)}" y1="${top - 10}" x2="${px(x)}" y2="${h - 20}" />`; });
    iv.forEach(([s, e], i) => {
      const y = top + i * rowH;
      const cls = ['tnode', 'acell', 'ibar'];
      const idx = view.picked.get(i);
      if (idx) cls.push('picked');
      if (view.current === i) cls.push('current');
      if (view.held.has(i) && view.current !== i) cls.push('held');
      if (view.wrong === i) cls.push('wrong');
      if ((!view.step && mark === i) || view.target === i) cls.push('target');
      const clickable = view.clickable && !idx;
      if (clickable) cls.push('clickable');
      body += `<g class="${cls.join(' ')}" data-id="${i}" ${clickable ? `role="button" tabindex="0" aria-label="Interval ${s} to ${e}"` : ''}>
        <rect x="${px(s) - 4}" y="${y}" width="${(e - s) * unit + 8}" height="20" rx="10" />
        <text x="${px(e) + 10}" y="${y + 10}" dy="0.35em" class="ilabel">[${s}, ${e}]</text>
        ${idx ? `<text class="badge" x="${px(s) - 14}" y="${y + 10}" dy="0.35em">${idx}</text>` : ''}</g>`;
    });
    return svgWrap(px(maxX) + 60, h, body, 'Intervals on a number line');
  }

  const IV_CODE = {
    merge: ['intervals.sort()                 # by start', 'merged = [intervals[0]]', 'for s, e in intervals[1:]:', '    if s <= merged[-1][1]:       # touches or overlaps', '        merged[-1][1] = max(merged[-1][1], e)', '    else:', '        merged.append([s, e])'],
    sweep: ['events = [(s, +1) for s, e in iv]', 'events += [(e, -1) for s, e in iv]', 'events.sort(key=lambda t: (t[0], -t[1]))  # starts 1st', 'open_now = best = 0', 'for x, d in events:', '    open_now += d', '    best = max(best, open_now)'],
    arrows: ['balloons.sort(key=lambda b: b[1])   # by end', 'arrow = None', 'for s, e in balloons:', '    if arrow is not None and s <= arrow:', '        continue                     # already burst', '    arrow = e                        # shoot at its end', '    arrows += 1'],
    covered: ['iv.sort(key=lambda t: (t[0], -t[1]))  # end desc on ties', 'far = -1', 'for s, e in iv:', '    if e <= far:', '        covered.append([s, e])', '    else:', '        far = e'],
  };

  function makeIntervals(count, maxLen = 6) {
    const out = [];
    const seen = new Set();
    while (out.length < count) {
      const s = randInt(0, 16);
      const e = s + randInt(1, maxLen);
      if (e > 20 || seen.has(`${s},${e}`)) continue;
      seen.add(`${s},${e}`);
      out.push([s, e]);
    }
    return out;
  }

  function intervalsQuestion(exId, size) {
    const count = { small: 5, medium: 7, large: 9 }[size];
    const fmt = ([s, e]) => `[${s}, ${e}]`;

    if (exId === 'group') {
      let iv;
      let groups;
      let t;
      for (let tries = 0; tries < 300; tries++) {
        iv = makeIntervals(count, 5);
        const order = iv.map((_, i) => i).sort((x, y) => iv[x][0] - iv[y][0] || iv[x][1] - iv[y][1]);
        groups = [];
        let end = -Infinity;
        order.forEach(i => {
          if (iv[i][0] <= end) { groups[groups.length - 1].push(i); end = Math.max(end, iv[i][1]); } else { groups.push([i]); end = iv[i][1]; }
        });
        const big = groups.filter(g => g.length >= 2);
        if (groups.length >= 2 && big.length) { t = pick(pick(big)); break; }
      }
      const g = groups.find(gr => gr.includes(t));
      const order = groups.flat();
      const steps = [];
      let gi = 0;
      let end = -Infinity;
      order.forEach((i, pos) => {
        const [s, e] = iv[i];
        if (s <= end) {
          const was = end;
          end = Math.max(end, e);
          steps.push({ line: 4, text: `${fmt(iv[i])} starts at ${s} <= the block's end, ${was}: merge it.${end > was ? ` The end grows to ${end}.` : ` The end stays ${end} (max, not overwrite).`}`, current: i, heldIds: groups[gi].filter(x => order.indexOf(x) < pos), out: groups[gi].filter(x => order.indexOf(x) <= pos), lines: [end] });
        } else {
          if (pos) gi++;
          end = e;
          steps.push({ line: pos ? 6 : 1, text: pos ? `${fmt(iv[i])} starts at ${s}, after the block's end: start a new block.` : `Sorted by start, ${fmt(iv[i])} opens the first block.`, current: i, out: [i], lines: [end] });
        }
      });
      return {
        prompt: `Sort by start and merge (closed intervals: [1, 3] and [3, 5] touch, so they merge). Click <strong>every interval in the same merged block</strong> as the ringed ${fmt(iv[t])}, including it.`,
        mode: 'sequence',
        ids: iv.map((_, i) => i),
        label: (id) => fmt(iv[id]),
        svg: (view) => intervalSvg(iv, view, { mark: t }),
        expected: g,
        fixedLength: false,
        validate: (ans) => ({ ok: sameSet(ans, g), k: 0 }),
        code: IV_CODE.merge,
        outLabel: 'Current block',
        steps: () => steps,
        right: () => `${g.map(i => fmt(iv[i])).join(', ')} merge into [${Math.min(...g.map(i => iv[i][0]))}, ${Math.max(...g.map(i => iv[i][1]))}].`,
        wrong: (k, ans) => {
          const extra = ans.find(i => !g.includes(i));
          const missing = g.find(i => !ans.includes(i));
          return `<p class="trace-bad">The block is ${g.map(i => fmt(iv[i])).join(', ')}.</p>
            <p class="trace-why">${missing !== undefined ? `${fmt(iv[missing])} belongs to it: chained overlaps count even if it never touches ${fmt(iv[t])} directly. ` : ''}${extra !== undefined ? `${fmt(iv[extra])} starts after the block's end, so it begins a new block. ` : ''}After sorting, an interval joins the block when its start is <= the block's furthest end so far.</p>`;
        },
      };
    }

    if (exId === 'sweep') {
      let iv;
      let best;
      for (let tries = 0; tries < 300; tries++) {
        iv = makeIntervals(count, 7);
        best = 0;
        for (let x = 0; x <= 21; x++) best = Math.max(best, iv.filter(([s, e]) => s <= x && x <= e).length);
        if (best >= 3) break;
      }
      const events = [];
      iv.forEach(([s, e], i) => { events.push([s, 1, i]); events.push([e, -1, i]); });
      events.sort((p, q) => p[0] - q[0] || q[1] - p[1]);
      const steps = [];
      const open = [];
      let top = 0;
      events.forEach(([x, d, i]) => {
        if (d > 0) open.push(i); else open.splice(open.indexOf(i), 1);
        if (open.length > top) top = open.length;
        steps.push({ line: 5, text: `x = ${x}: ${fmt(iv[i])} ${d > 0 ? 'opens' : 'closes'}. ${open.length} open${open.length === top && d > 0 ? ` (the most so far: ${top})` : ''}.`, current: i, heldIds: [...open], lines: [x], heldLabel: 'open', held: open.map(j => fmt(iv[j])) });
      });
      return {
        prompt: 'Sweep the line left to right (closed intervals: touching counts as overlapping, so starts are processed before ends at the same x). Click a set of intervals that are <strong>all open at once</strong> when the count is highest.',
        mode: 'sequence',
        ids: iv.map((_, i) => i),
        label: (id) => fmt(iv[id]),
        svg: (view) => intervalSvg(iv, view),
        expected: (() => { for (let x = 0; x <= 21; x++) { const at = iv.map((_, i) => i).filter(i => iv[i][0] <= x && x <= iv[i][1]); if (at.length === best) return at; } return []; })(),
        fixedLength: false,
        validate: (ans) => {
          const common = Math.max(...ans.map(i => iv[i][0])) <= Math.min(...ans.map(i => iv[i][1]));
          return { ok: ans.length === best && common, k: 0 };
        },
        code: IV_CODE.sweep,
        steps: () => steps,
        right: () => `${best} at once. The +1/-1 sweep finds the count without comparing every pair; it is the same idea as Meeting Rooms II.`,
        wrong: (k, ans) => {
          const common = ans.length && Math.max(...ans.map(i => iv[i][0])) <= Math.min(...ans.map(i => iv[i][1]));
          return `<p class="trace-bad">${common ? `Those ${ans.length} do overlap at one point, but the most open at once is ${best}.` : `Those are not all open at the same moment: the latest start (${Math.max(...ans.map(i => iv[i][0]))}) is after the earliest end (${Math.min(...ans.map(i => iv[i][1]))}).`}</p>
            <p class="trace-why">A set is open at once when its latest start is <= its earliest end. Use "Show me" to watch the count rise and fall.</p>`;
        },
      };
    }

    if (exId === 'arrows') {
      let iv;
      for (let tries = 0; tries < 300; tries++) {
        iv = makeIntervals(count, 5);
        if (new Set(iv.map(([, e]) => e)).size === iv.length) break;
      }
      const order = iv.map((_, i) => i).sort((x, y) => iv[x][1] - iv[y][1]);
      const shots = [];
      const lines = [];
      const steps = [];
      let arrow = null;
      order.forEach(i => {
        const [s, e] = iv[i];
        if (arrow !== null && s <= arrow) {
          steps.push({ line: 4, text: `${fmt(iv[i])} starts at ${s} <= ${arrow}: the arrow at ${arrow} already bursts it.`, current: i, out: [...shots], lines: [...lines] });
          return;
        }
        arrow = e;
        shots.push(i);
        lines.push(e);
        steps.push({ line: 5, text: `${fmt(iv[i])} is not burst yet and ends soonest: shoot at its end, x = ${e}.`, current: i, out: [...shots], lines: [...lines] });
      });
      return {
        prompt: 'Minimum arrows to burst balloons: sort by end, and shoot at the end of each balloon that is not burst yet. Click the balloons whose ends get an arrow, in order.',
        mode: 'sequence',
        ids: iv.map((_, i) => i),
        label: (id) => fmt(iv[id]),
        svg: (view) => intervalSvg(iv, view),
        expected: shots,
        fixedLength: false,
        code: IV_CODE.arrows,
        outLabel: 'Arrows at',
        steps: () => steps,
        right: () => `${shots.length} arrows, at x = ${lines.join(', ')}. Shooting at the earliest end is as far right as the arrow can go while still hitting that balloon, so it catches the most others.`,
        wrong: (k, ans) => {
          if (k === shots.length) return `<p class="trace-bad">${said(k)} After that, every balloon is burst.</p>`;
          const next = shots[k];
          const prev = k ? iv[shots[k - 1]][1] : null;
          return `<p class="trace-bad">${said(k)} ${k < ans.length ? `Next is ${fmt(iv[next])}, not ${fmt(iv[ans[k]])}.` : `Keep going: next is ${fmt(iv[next])}.`}</p>
            <p class="trace-why">${prev !== null ? `Everything starting at or before ${prev} was burst by the last arrow. ` : ''}Of the balloons left, ${fmt(iv[next])} ends first, at ${iv[next][1]}.${k < ans.length && prev !== null && iv[ans[k]][0] <= prev ? ` ${fmt(iv[ans[k]])} is already burst.` : ''}</p>`;
        },
      };
    }

    // covered: intervals inside another
    let iv;
    let covered;
    for (let tries = 0; tries < 300; tries++) {
      iv = makeIntervals(count, 7);
      covered = iv.map((_, i) => i).filter(i => iv.some(([s, e], j) => j !== i && s <= iv[i][0] && iv[i][1] <= e));
      if (covered.length >= 1 && covered.length <= count - 2) break;
    }
    const order = iv.map((_, i) => i).sort((x, y) => iv[x][0] - iv[y][0] || iv[y][1] - iv[x][1]);
    const steps = [];
    let far = -1;
    const found = [];
    order.forEach(i => {
      const [, e] = iv[i];
      if (e <= far) { found.push(i); steps.push({ line: 4, text: `${fmt(iv[i])} ends at ${e} <= ${far}, the furthest end so far, and started no earlier: it is covered.`, current: i, out: [...found], lines: [far] }); }
      else { far = e; steps.push({ line: 6, text: `${fmt(iv[i])} reaches further (${e}): not covered. far = ${e}.`, current: i, out: [...found], lines: [far] }); }
    });
    return {
      prompt: 'Click <strong>every interval that lies inside another one</strong> ([2, 4] is inside [1, 6], and inside [2, 6] too). Then Check.',
      mode: 'sequence',
      ids: iv.map((_, i) => i),
      label: (id) => fmt(iv[id]),
      svg: (view) => intervalSvg(iv, view),
      expected: covered,
      fixedLength: false,
      validate: (ans) => ({ ok: sameSet(ans, covered), k: 0 }),
      code: IV_CODE.covered,
      outLabel: 'Covered',
      steps: () => steps,
      right: () => `${covered.map(i => fmt(iv[i])).join(', ')}. Sorting by start, and by end descending on ties, puts every container before what it contains, so one running "furthest end" is enough.`,
      wrong: (k, ans) => {
        const extra = ans.find(i => !covered.includes(i));
        const missing = covered.find(i => !ans.includes(i));
        const host = (i) => iv.findIndex(([s, e], j) => j !== i && s <= iv[i][0] && iv[i][1] <= e);
        return `<p class="trace-bad">The covered ones are ${covered.map(i => fmt(iv[i])).join(', ')}.</p>
          <p class="trace-why">${missing !== undefined ? `${fmt(iv[missing])} is inside ${fmt(iv[host(missing)])}. ` : ''}${extra !== undefined ? `${fmt(iv[extra])} sticks out of every other interval. ` : ''}Covered means start >= the other's start and end <= its end; equal ends or starts still count.</p>`;
      },
    };
  }

  KITS.intervals = {
    title: 'Intervals',
    exercises: [
      { id: 'group', name: 'Merge: which block?' },
      { id: 'sweep', name: 'Sweep: most at once' },
      { id: 'arrows', name: 'Arrows: where to shoot' },
      { id: 'covered', name: 'Covered intervals' },
    ],
    make: intervalsQuestion,
  };

  // ===========================================================================
  // Arrays & hashing
  // ===========================================================================

  const AH_CODE = {
    twosum: ['seen = {}   # value -> index', 'for j, x in enumerate(nums):', '    if target - x in seen:', '        return [seen[target - x], j]', '    seen[x] = j'],
    prefix: ['prefix = [0]', 'for x in nums:', '    prefix.append(prefix[-1] + x)', '# sum(nums[l..r]) == prefix[r + 1] - prefix[l]'],
    starts: ['s = set(nums)', 'for x in s:', '    if x - 1 in s:', '        continue          # not the start of a run', '    length = 1', '    while x + length in s:', '        length += 1'],
    count: ['count = {0: 1}   # prefix sum -> how often seen', 'total = ans = 0', 'for x in nums:', '    total += x', '    ans += count.get(total - k, 0)', '    count[total] = count.get(total, 0) + 1'],
  };

  const fmtMap = (m) => (Object.keys(m).length ? Object.entries(m).map(([k, v]) => `${k}: ${v}`) : []);

  function arraysQuestion(exId, size) {
    const n = { small: 6, medium: 8, large: 10 }[size];

    if (exId === 'twosum') {
      const a = distinctValues(n, 30);
      const i0 = randInt(0, n - 2);
      const target = a[i0] + a[randInt(i0 + 1, n - 1)];
      const seen = {};
      const steps = [];
      let ans;
      for (let j = 0; j < n; j++) {
        const x = a[j];
        if ((target - x) in seen) {
          ans = [seen[target - x], j];
          steps.push({ line: 3, text: `${target} - ${x} = ${target - x} is in seen (index ${seen[target - x]}): return [${ans.join(', ')}].`, current: j, heldIds: [seen[target - x]], heldLabel: 'seen', held: fmtMap(seen) });
          break;
        }
        seen[x] = j;
        steps.push({ line: 4, text: `${target} - ${x} = ${target - x} is not in seen yet. Remember ${x} at index ${j}.`, current: j, heldLabel: 'seen', held: fmtMap(seen) });
      }
      return {
        prompt: `Two Sum, one pass with a hash map, target <strong>${target}</strong>. Click the pair it returns: the earlier index, then the later.`,
        mode: 'sequence',
        slots: ['i', 'j'],
        ids: a.map((_, i) => i),
        label: cellLabel(a),
        svg: (view) => arraySvg(a, view),
        expected: ans,
        fixedLength: true,
        code: AH_CODE.twosum,
        steps: () => steps,
        right: () => `[${ans.join(', ')}]: ${a[ans[0]]} + ${a[ans[1]]} = ${target}. It returns as soon as the second number of a pair arrives, because the first is already in the map.`,
        wrong: (k, answer) => {
          const ok = a[answer[0]] + a[answer[1]] === target;
          return `<p class="trace-bad">It returns [${ans.join(', ')}].</p>
            <p class="trace-why">${ok ? `${a[answer[0]]} + ${a[answer[1]]} is also ${target}, but the loop stops at the first index j whose complement was already seen, which is ${ans[1]}.` : `${a[answer[0]]} + ${a[answer[1]]} is not ${target}.`} Each number looks up target - x among the numbers <em>before</em> it, so one pass is enough.</p>`;
        },
      };
    }

    if (exId === 'prefix') {
      const a = Array.from({ length: n }, () => randInt(-5, 9));
      const p = [0];
      a.forEach(x => p.push(p[p.length - 1] + x));
      const l = randInt(0, n - 2);
      const r = randInt(l + 1, n - 1);
      const steps = [{ line: 0, text: 'prefix[0] = 0: the sum of no numbers.', current: 0, vals: p.map((v, i) => (i === 0 ? v : '')) }];
      for (let i = 1; i <= n; i++) {
        steps.push({ line: 2, text: `prefix[${i}] = prefix[${i - 1}] + nums[${i - 1}] = ${p[i - 1]} + ${a[i - 1]} = ${p[i]}: the sum of the first ${i}.`, current: i, vals: p.map((v, j) => (j <= i ? v : '')) });
      }
      steps.push({ line: 3, text: `sum(nums[${l}..${r}]) = prefix[${r + 1}] - prefix[${l}] = ${p[r + 1]} - ${p[l]} = ${p[r + 1] - p[l]}.`, vals: p, heldIds: [r + 1, l], current: r + 1 });
      return {
        prompt: `The cells are the <strong>prefix sums</strong> of nums: prefix[i] is the sum of the first i numbers. To get sum(nums[${l}..${r}]) (inclusive), click the prefix you take, then the one you subtract.`,
        aside: `<p class="trace-ops">nums = <code>[${a.join(', ')}]</code></p>`,
        mode: 'sequence',
        slots: ['take', 'minus'],
        ids: p.map((_, i) => i),
        label: (id) => `prefix[${id}]`,
        svg: (view) => arraySvg(p, view),
        expected: [r + 1, l],
        fixedLength: true,
        code: AH_CODE.prefix,
        steps: () => steps,
        right: () => `prefix[${r + 1}] - prefix[${l}] = ${p[r + 1]} - ${p[l]} = ${p[r + 1] - p[l]}. prefix[${r + 1}] covers nums[0..${r}]; subtracting prefix[${l}] removes nums[0..${l - 1}].`,
        wrong: (k, answer) => `<p class="trace-bad">It is prefix[${r + 1}] - prefix[${l}].</p>
          <p class="trace-why">${answer[0] === r ? `prefix[${r}] stops before nums[${r}]: prefix[i] covers indexes 0..i - 1, so to include nums[${r}] take prefix[${r + 1}]. ` : ''}${answer[1] === l - 1 ? `prefix[${l - 1}] would leave nums[${l - 1}] in. ` : answer[1] === l + 1 ? `prefix[${l + 1}] would also remove nums[${l}]. ` : ''}The prefix array has one more cell than nums, and that shift is where the off-by-ones come from.</p>`,
      };
    }

    if (exId === 'starts') {
      let vals;
      let starts;
      for (let tries = 0; tries < 200; tries++) {
        const set = new Set();
        while (set.size < n) {
          const s = randInt(1, 40);
          const len = randInt(1, 4);
          for (let v = s; v < s + len && set.size < n; v++) set.add(v);
        }
        vals = shuffle([...set]);
        starts = vals.map((v, i) => i).filter(i => !set.has(vals[i] - 1));
        if (starts.length >= 2 && starts.length <= n - 2) break;
      }
      const set = new Set(vals);
      const steps = vals.map((x, i) => {
        if (set.has(x - 1)) return { line: 3, text: `${x - 1} is in the set, so ${x} is inside a run: skip it. Its run is counted from its start.`, current: i, out: starts.filter(j => j <= i) };
        let len = 1;
        while (set.has(x + len)) len++;
        return { line: 6, text: `${x - 1} is not in the set, so ${x} starts a run: ${Array.from({ length: len }, (_, t) => x + t).join(', ')} (length ${len}).`, current: i, out: starts.filter(j => j <= i) };
      });
      return {
        prompt: 'Longest consecutive sequence in O(n): only start counting from a number that begins a run. Click <strong>every number the while loop starts from</strong>, then Check.',
        mode: 'sequence',
        ids: vals.map((_, i) => i),
        label: (id) => String(vals[id]),
        svg: (view) => arraySvg(vals, view),
        expected: starts,
        fixedLength: false,
        validate: (ans) => ({ ok: sameSet(ans, starts), k: 0 }),
        code: AH_CODE.starts,
        outLabel: 'Run starts',
        steps: () => steps,
        right: () => `${starts.map(i => vals[i]).sort((x, y) => x - y).join(', ')}. Each number is walked over by exactly one while loop, from its run's start, so the total work is O(n).`,
        wrong: (k, ans) => {
          const extra = ans.find(i => !starts.includes(i));
          const missing = starts.find(i => !ans.includes(i));
          return `<p class="trace-bad">The starts are ${starts.map(i => vals[i]).sort((x, y) => x - y).join(', ')}.</p>
            <p class="trace-why">${extra !== undefined ? `${vals[extra] - 1} is in the set, so ${vals[extra]} is in the middle of a run. ` : ''}${missing !== undefined ? `${vals[missing] - 1} is not in the set, so ${vals[missing]} starts a run. ` : ''}A start is a number whose predecessor is missing; position in the array does not matter.</p>`;
        },
      };
    }

    // count: subarrays summing to k with prefix counts
    let a;
    let k;
    let ans;
    let unseeded;
    for (let tries = 0; tries < 300; tries++) {
      a = Array.from({ length: n }, () => randInt(-3, 5));
      k = randInt(2, 8);
      ans = 0;
      for (let i = 0; i < n; i++) for (let j = i; j < n; j++) if (a.slice(i, j + 1).reduce((x, y) => x + y, 0) === k) ans++;
      let t = 0;
      unseeded = ans;
      a.forEach(x => { t += x; if (t === k) unseeded--; });
      if (ans >= 2 && unseeded !== ans) break;
    }
    const count = { 0: 1 };
    let total = 0;
    let run = 0;
    const lines = [];
    a.forEach((x, i) => {
      total += x;
      const add = count[total - k] || 0;
      run += add;
      lines.push(`i = ${i}: total ${total}, look up ${total} - ${k} = ${total - k}: +${add} (ans ${run})`);
      count[total] = (count[total] || 0) + 1;
    });
    const opts = [...new Set([ans, unseeded, ans + 1, Math.max(0, ans - 1)])].slice(0, 4).sort((x, y) => x - y);
    return {
      prompt: `How many subarrays (contiguous, non-empty) of this array sum to <strong>${k}</strong>?`,
      mode: 'choice',
      svg: (view) => arraySvg(a, view),
      sequence: [`k = ${k}`],
      choices: opts.map(v => ({ id: String(v), label: String(v) })),
      answerId: String(ans),
      explain: (picked) => `<p class="${picked === String(ans) ? 'trace-ok' : 'trace-bad'}"><strong>${picked === String(ans) ? 'Right' : 'Not quite'}:</strong> ${ans}.${picked === String(unseeded) ? ` ${unseeded} misses the subarrays that start at index 0: that is what happens without the {0: 1} seed in the map.` : ''}</p>
        <p class="trace-why">A subarray ending at i sums to k when some earlier prefix equals total - k, so the map counts prefixes seen so far:</p>
        <ol class="trace-notes">${lines.map(t => `<li>${esc(t)}</li>`).join('')}</ol>`,
    };
  }

  KITS['arrays-hashing'] = {
    title: 'Arrays & hashing',
    exercises: [
      { id: 'twosum', name: 'Two Sum: which pair?' },
      { id: 'prefix', name: 'Prefix sums: range sum' },
      { id: 'starts', name: 'Consecutive runs: starts' },
      { id: 'count', name: 'Subarrays summing to k' },
    ],
    make: arraysQuestion,
  };

  // ===========================================================================
  // 2-D dynamic programming
  // ===========================================================================

  // tableSvg draws an R x C grid of numbers. Ids are r * C + c. A step can
  // show a dp table instead (step.table, with '' for cells not filled yet).
  function tableSvg(grid, view, { blocked = new Set(), clickableIds = null, dim = new Set(), mark = null, heads = false } = {}) {
    const R = grid.length;
    const C = grid[0].length;
    const size = 46;
    const pad = heads ? 34 : 18;
    const table = view.step?.table;
    let body = '';
    for (let r = 0; r < R; r++) {
      for (let c = 0; c < C; c++) {
        const id = r * C + c;
        const cls = ['tnode', 'acell', 'tcell'];
        if (blocked.has(id)) cls.push('blocked');
        if ((view.step?.dim || dim).has(id)) cls.push('dim');
        if (!view.step && mark === id) cls.push('target');
        const idx = view.picked.get(id);
        if (idx) cls.push('picked');
        if (view.current === id) cls.push('current');
        if (view.held.has(id) && view.current !== id) cls.push('held');
        if (view.wrong === id) cls.push('wrong');
        const clickable = view.clickable && !blocked.has(id) && !idx && (!clickableIds || clickableIds.includes(id));
        if (clickable) cls.push('clickable');
        const x = pad + c * size;
        const y = pad + r * size;
        const text = blocked.has(id) ? '■' : table ? table[r][c] : grid[r][c];
        body += `<g class="${cls.join(' ')}" data-id="${id}" ${clickable ? `role="button" tabindex="0" aria-label="Row ${r}, column ${c}"` : ''}>
          <rect x="${x + 2}" y="${y + 2}" width="${size - 4}" height="${size - 4}" rx="5" />
          <text x="${x + size / 2}" y="${y + size / 2}" dy="0.35em">${esc(text)}</text>
          ${table && !blocked.has(id) ? `<text class="idx" x="${x + 10}" y="${y + 13}">${esc(grid[r][c])}</text>` : ''}
          ${idx ? `<text class="badge" x="${x + size - 11}" y="${y + 13}">${idx}</text>` : ''}</g>`;
      }
    }
    if (heads) {
      for (let c = 0; c < C; c++) body += `<text class="idx head" x="${pad + c * size + size / 2}" y="${pad - 10}">${c}</text>`;
      for (let r = 0; r < R; r++) body += `<text class="idx head" x="${pad - 14}" y="${pad + r * size + size / 2}" dy="0.35em">${r}</text>`;
    }
    return svgWrap(C * size + 2 * pad, R * size + 2 * pad, body, 'Grid');
  }

  const DP2_CODE = {
    minpath: ['dp[0][0] = grid[0][0]', 'for r in range(R):', '    for c in range(C):', '        up = dp[r - 1][c] if r else inf', '        left = dp[r][c - 1] if c else inf', '        dp[r][c] = grid[r][c] + min(up, left)'],
    lcs: ['for i in range(1, len(a) + 1):', '    for j in range(1, len(b) + 1):', '        if a[i - 1] == b[j - 1]:', '            dp[i][j] = dp[i - 1][j - 1] + 1', '        else:', '            dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])'],
  };

  const tableHtml = (rows, rowHead = (r) => r, colHead = null) => `<table class="trace-orders trace-dp">${colHead ? `<tr><th></th>${colHead.map(h => `<th>${esc(h)}</th>`).join('')}</tr>` : ''}${rows.map((row, r) => `<tr><th>${esc(rowHead(r))}</th>${row.map(v => `<td>${esc(v)}</td>`).join('')}</tr>`).join('')}</table>`;

  function dp2Question(exId, size) {
    const [R, C] = { small: [3, 3], medium: [3, 4], large: [4, 5] }[size];

    if (exId === 'minpath') {
      let grid;
      let dp;
      for (let tries = 0; tries < 100; tries++) {
        grid = Array.from({ length: R }, () => Array.from({ length: C }, () => randInt(1, 9)));
        dp = grid.map(row => row.map(() => 0));
        for (let r = 0; r < R; r++) for (let c = 0; c < C; c++) dp[r][c] = grid[r][c] + (r || c ? Math.min(r ? dp[r - 1][c] : Infinity, c ? dp[r][c - 1] : Infinity) : 0);
        // Not just the obvious all-right-then-down path.
        const edge = grid[0].reduce((x, y) => x + y, 0) + grid.slice(1).reduce((x, row) => x + row[C - 1], 0);
        if (edge !== dp[R - 1][C - 1]) break;
      }
      const best = dp[R - 1][C - 1];
      const steps = [];
      const shown = grid.map(row => row.map(() => ''));
      for (let r = 0; r < R; r++) for (let c = 0; c < C; c++) {
        shown[r][c] = dp[r][c];
        const from = !r && !c ? 'the start' : !r ? `left (${dp[r][c - 1]})` : !c ? `above (${dp[r - 1][c]})` : `min(above ${dp[r - 1][c]}, left ${dp[r][c - 1]})`;
        steps.push({ line: 5, text: `dp[${r}][${c}] = ${grid[r][c]} + ${from === 'the start' ? '0' : from} = ${dp[r][c]}.`, current: r * C + c, heldIds: [r ? (r - 1) * C + c : -1, c ? r * C + c - 1 : -1], table: shown.map(row => [...row]) });
      }
      const path = [];
      let r = R - 1;
      let c = C - 1;
      while (r || c) { path.unshift(r * C + c); if (!r) c--; else if (!c) r--; else if (dp[r - 1][c] <= dp[r][c - 1]) r--; else c--; }
      steps.push({ line: 5, text: `Best: ${best}. Walk back from the corner, always to the smaller neighbour, to recover a path.`, table: shown.map(row => [...row]), out: [0, ...path] });
      const id = (rr, cc) => rr * C + cc;
      return {
        prompt: `Minimum path sum, moving only right or down. Click a cheapest path cell by cell, from the top-left to the bottom-right.`,
        mode: 'sequence',
        ids: grid.flatMap((row, rr) => row.map((_, cc) => id(rr, cc))),
        label: (x) => String(grid[Math.floor(x / C)][x % C]),
        svg: (view) => tableSvg(grid, view),
        expected: [0, ...path],
        fixedLength: true,
        validate: (ans) => {
          let k = 0;
          if (ans[0] !== 0) return { ok: false, k: 0 };
          for (k = 1; k < ans.length; k++) {
            const [pr, pc] = [Math.floor(ans[k - 1] / C), ans[k - 1] % C];
            const [cr, cc] = [Math.floor(ans[k] / C), ans[k] % C];
            if (!((cr === pr && cc === pc + 1) || (cc === pc && cr === pr + 1))) break;
          }
          const sum = ans.reduce((s, x) => s + grid[Math.floor(x / C)][x % C], 0);
          return { ok: k === ans.length && sum === best, k };
        },
        code: DP2_CODE.minpath,
        steps: () => steps,
        right: () => `Sum ${best}. dp[r][c] is the cheapest way to reach (r, c): its own cost plus the cheaper of the cell above and the cell to the left.`,
        wrong: (k, ans) => {
          if (k < ans.length) return `<p class="trace-bad">${k === 0 ? 'The path starts at the top-left cell.' : `Step ${k + 1} is not one move right or down from the cell before.`}</p>`;
          const sum = ans.reduce((s, x) => s + grid[Math.floor(x / C)][x % C], 0);
          return `<p class="trace-bad">That path costs ${sum}; the cheapest costs ${best}.</p>
            <p class="trace-why">Picking the cheaper next cell each time is greedy and can walk into an expensive region. The table compares every way in. dp:</p>${tableHtml(dp)}`;
        },
      };
    }

    if (exId === 'paths') {
      let blocked;
      let ways;
      let naive;
      let free;
      for (let tries = 0; tries < 300; tries++) {
        blocked = new Set();
        const count = randInt(1, Math.max(1, Math.floor(R * C / 5)));
        while (blocked.size < count) { const x = randInt(1, R * C - 2); blocked.add(x); }
        const solve = (edgeBug, ignore) => {
          const w = Array.from({ length: R }, () => Array(C).fill(0));
          for (let r = 0; r < R; r++) for (let c = 0; c < C; c++) {
            if (!ignore && blocked.has(r * C + c)) { w[r][c] = 0; continue; }
            if (!r || !c) w[r][c] = edgeBug ? 1 : (!r && !c ? 1 : !r ? w[r][c - 1] : w[r - 1][c]);
            else w[r][c] = w[r - 1][c] + w[r][c - 1];
          }
          return w;
        };
        ways = solve(false, false);
        naive = solve(true, false)[R - 1][C - 1];
        free = solve(false, true)[R - 1][C - 1];
        if (ways[R - 1][C - 1] >= 1 && naive !== ways[R - 1][C - 1]) break;
      }
      const ans = ways[R - 1][C - 1];
      const grid = Array.from({ length: R }, () => Array(C).fill(''));
      const opts = [...new Set([ans, naive, free, ans + 1])].filter(v => v >= 0).slice(0, 4).sort((x, y) => x - y);
      return {
        prompt: 'How many paths go from the top-left to the bottom-right, moving only right or down and never through a ■?',
        mode: 'choice',
        svg: (view) => tableSvg(grid, view, { blocked }),
        sequence: [`${R} × ${C} grid`],
        choices: opts.map(v => ({ id: String(v), label: String(v) })),
        answerId: String(ans),
        explain: (picked) => `<p class="${picked === String(ans) ? 'trace-ok' : 'trace-bad'}"><strong>${picked === String(ans) ? 'Right' : 'Not quite'}:</strong> ${ans}.${picked === String(naive) ? ` ${naive} comes from setting the whole first row and column to 1: a block there cuts off every cell after it.` : picked === String(free) ? ` ${free} ignores the blocks.` : ''}</p>
          <p class="trace-why">ways[r][c] = ways above + ways to the left, and 0 on a block:</p>${tableHtml(ways.map((row, r) => row.map((v, c) => (blocked.has(r * C + c) ? '■' : v))))}`,
      };
    }

    if (exId === 'lcs') {
      const alpha = 'ABCD';
      const la = { small: 5, medium: 6, large: 8 }[size];
      let a;
      let b;
      let dp;
      for (let tries = 0; tries < 200; tries++) {
        a = Array.from({ length: la }, () => pick([...alpha]));
        b = Array.from({ length: la - 1 + randInt(0, 2) }, () => pick([...alpha]));
        dp = Array.from({ length: a.length + 1 }, () => Array(b.length + 1).fill(0));
        for (let i = 1; i <= a.length; i++) for (let j = 1; j <= b.length; j++) dp[i][j] = a[i - 1] === b[j - 1] ? dp[i - 1][j - 1] + 1 : Math.max(dp[i - 1][j], dp[i][j - 1]);
        if (dp[a.length][b.length] >= 3 && dp[a.length][b.length] < a.length - 1) break;
      }
      const L = dp[a.length][b.length];
      const isSub = (s) => { let j = 0; for (const ch of b) if (j < s.length && ch === s[j]) j++; return j === s.length; };
      const one = [];
      let i = a.length;
      let j = b.length;
      while (i && j) { if (a[i - 1] === b[j - 1]) { one.unshift(i - 1); i--; j--; } else if (dp[i - 1][j] >= dp[i][j - 1]) i--; else j--; }
      const steps = a.map((ch, r) => ({ line: 5, text: `Row ${r + 1} ('${ch}'): ${dp[r + 1].slice(1).join(' ')}. Matching '${ch}' with the same letter in b adds 1 to the diagonal; otherwise take the better of above and left.`, current: r, heldIds: a.map((_, x) => x).filter(x => x < r) }));
      return {
        prompt: `Click a <strong>longest common subsequence</strong> of a and b, as letters of a, left to right. Any one of the longest counts.`,
        aside: `<p class="trace-ops">a = <code>${a.join('')}</code> (the cells) &nbsp; b = <code>${b.join('')}</code></p>`,
        mode: 'sequence',
        ids: a.map((_, x) => x),
        label: cellLabel(a),
        svg: (view) => arraySvg(a, view),
        expected: one,
        fixedLength: false,
        validate: (ans) => {
          const inc = ans.every((x, t) => !t || x > ans[t - 1]);
          const sub = isSub(ans.map(x => a[x]));
          return { ok: inc && sub && ans.length === L, k: inc && sub ? ans.length : 0 };
        },
        code: DP2_CODE.lcs,
        steps: () => steps,
        right: () => `Length ${L}, for example ${one.map(x => a[x]).join('')}. dp[i][j] is the LCS of the first i letters of a and the first j of b.`,
        wrong: (k, ans) => {
          const inc = ans.every((x, t) => !t || x > ans[t - 1]);
          const word = ans.map(x => a[x]).join('');
          return `<p class="trace-bad">${!inc ? 'Pick the letters left to right.' : !isSub(ans.map(x => a[x])) ? `${word} is not a subsequence of b (${b.join('')}).` : `${word} is common to both, but the longest has ${L} letters, for example ${one.map(x => a[x]).join('')}.`}</p>
            <p class="trace-why">dp table (rows: a, columns: b):</p>${tableHtml(dp.slice(1).map(row => row.slice(1)), (r) => a[r], b)}`;
        },
      };
    }

    // knapsack (0/1), with the unbounded and greedy answers as traps
    let items;
    let cap;
    let best;
    let unb;
    let greedy;
    for (let tries = 0; tries < 300; tries++) {
      const count = { small: 3, medium: 4, large: 5 }[size];
      items = Array.from({ length: count }, () => ({ w: randInt(1, 6), v: randInt(2, 15) }));
      cap = randInt(5, 10);
      const zero = Array(cap + 1).fill(0);
      items.forEach(({ w, v }) => { for (let c = cap; c >= w; c--) zero[c] = Math.max(zero[c], zero[c - w] + v); });
      best = zero[cap];
      const u = Array(cap + 1).fill(0);
      items.forEach(({ w, v }) => { for (let c = w; c <= cap; c++) u[c] = Math.max(u[c], u[c - w] + v); });
      unb = u[cap];
      let room = cap;
      greedy = 0;
      [...items].sort((x, y) => y.v / y.w - x.v / x.w).forEach(({ w, v }) => { if (w <= room) { room -= w; greedy += v; } });
      if (unb !== best && greedy !== best) break;
    }
    const rows = [];
    const t = Array(cap + 1).fill(0);
    items.forEach(({ w, v }) => { for (let c = cap; c >= w; c--) t[c] = Math.max(t[c], t[c - w] + v); rows.push([...t]); });
    const opts = [...new Set([best, unb, greedy, best + 1])].slice(0, 4).sort((x, y) => x - y);
    return {
      prompt: `0/1 knapsack: each item at most once, capacity <strong>${cap}</strong>. What is the most value you can carry?`,
      mode: 'choice',
      svg: (view) => arraySvg(items.map(it => `${it.w}|${it.v}`), view),
      sequence: items.map(it => `w ${it.w}, v ${it.v}`),
      choices: opts.map(v => ({ id: String(v), label: String(v) })),
      answerId: String(best),
      explain: (picked) => `<p class="${picked === String(best) ? 'trace-ok' : 'trace-bad'}"><strong>${picked === String(best) ? 'Right' : 'Not quite'}:</strong> ${best}.${picked === String(unb) ? ` ${unb} reuses an item: that is what the one-row table gives if capacity runs upwards.` : picked === String(greedy) ? ` ${greedy} is greedy by value per weight, which can leave capacity unused.` : ''}</p>
        <p class="trace-why">best[c] after each item (capacity looping downwards, so each item is used once):</p>${tableHtml(rows, (r) => `+ w${items[r].w} v${items[r].v}`, Array.from({ length: cap + 1 }, (_, c) => c))}`,
    };
  }

  KITS['2d-dp'] = {
    title: '2-D dynamic programming',
    exercises: [
      { id: 'minpath', name: 'Minimum path sum' },
      { id: 'paths', name: 'Paths around blocks' },
      { id: 'lcs', name: 'Longest common subsequence' },
      { id: 'knapsack', name: '0/1 knapsack' },
    ],
    make: dp2Question,
  };

  // ===========================================================================
  // Greedy
  // ===========================================================================

  const GR_CODE = {
    jump: ['jumps = end = far = 0', 'for i in range(len(nums) - 1):', '    far = max(far, i + nums[i])', '    if i == end:            # used up this jump', '        jumps += 1', '        end = far'],
    gas: ['tank = start = 0', 'for i in range(len(gas)):', '    tank += gas[i] - cost[i]', '    if tank < 0:      # stuck before i + 1', '        start = i + 1', '        tank = 0', 'return start       # total checked first'],
    partition: ['last = {ch: i for i, ch in enumerate(s)}', 'start = end = 0', 'for i, ch in enumerate(s):', '    end = max(end, last[ch])', '    if i == end:      # nothing here appears later', '        parts.append(i - start + 1)', '        start = i + 1'],
  };

  function greedyQuestion(exId, size) {
    const n = { small: 7, medium: 9, large: 12 }[size];

    if (exId === 'jump') {
      let a;
      let minJ;
      for (let tries = 0; tries < 300; tries++) {
        a = Array.from({ length: n }, () => randInt(1, 3));
        const d = Array(n).fill(Infinity);
        d[0] = 0;
        for (let i = 0; i < n; i++) for (let s = 1; s <= a[i] && i + s < n; s++) d[i + s] = Math.min(d[i + s], d[i] + 1);
        minJ = d[n - 1];
        // More than one way to land should exist, so it is not forced.
        if (minJ >= 3 && a.filter(v => v > 1).length >= 3) break;
      }
      const steps = [];
      let jumps = 0;
      let end = 0;
      let far = 0;
      const lands = [];
      for (let i = 0; i < n - 1; i++) {
        far = Math.max(far, i + a[i]);
        if (i === end) {
          jumps++;
          end = far;
          steps.push({ line: 5, text: `Index ${i} is the edge of what ${jumps - 1} jump${jumps - 1 === 1 ? '' : 's'} can reach. Jump ${jumps} can reach up to index ${Math.min(end, n - 1)}.`, current: i, window: [i, Math.min(end, n - 1)] });
        } else {
          steps.push({ line: 2, text: `From index ${i} you could reach ${i + a[i]}; the farthest so far is ${far}.`, current: i, window: [i, Math.min(far, n - 1)] });
        }
      }
      // One landing path with the minimum: from each window jump to the cell that reaches farthest.
      let i = 0;
      while (i < n - 1) {
        if (i + a[i] >= n - 1) { lands.push(n - 1); break; }
        let bestNext = i + 1;
        for (let s = 1; s <= a[i]; s++) if (i + s + a[i + s] > bestNext + a[bestNext]) bestNext = i + s;
        lands.push(bestNext);
        i = bestNext;
      }
      return {
        prompt: `Jump game: from index i you may jump 1 to nums[i] steps. Starting at index 0, click where you land on each jump of a route with the <strong>fewest jumps</strong>, ending on the last index.`,
        mode: 'sequence',
        ids: a.map((_, x) => x).slice(1),
        label: cellLabel(a),
        svg: (view) => arraySvg(a, view, { mark: 0, clickableIds: a.map((_, x) => x).slice(1) }),
        expected: lands,
        fixedLength: false,
        validate: (ans) => {
          let prev = 0;
          let k = 0;
          for (; k < ans.length; k++) { if (ans[k] <= prev || ans[k] > prev + a[prev]) break; prev = ans[k]; }
          return { ok: k === ans.length && prev === n - 1 && ans.length === minJ, k };
        },
        code: GR_CODE.jump,
        steps: () => steps,
        right: () => `${minJ} jumps. The greedy never has to pick a landing spot: it tracks how far j jumps can reach, and when it runs out of a jump's range it takes another.`,
        wrong: (k, ans) => {
          if (k < ans.length) return `<p class="trace-bad">Index ${ans[k]} cannot be reached in one jump from index ${k ? ans[k - 1] : 0} (it allows up to ${a[k ? ans[k - 1] : 0]} steps forward).</p>`;
          if (ans[ans.length - 1] !== n - 1) return `<p class="trace-bad">Keep going until you land on the last index.</p>`;
          return `<p class="trace-bad">That takes ${ans.length} jumps; ${minJ} is enough, for example ${lands.join(' → ')}.</p>
            <p class="trace-why">Jumping as far as possible each time is not the rule; jump to the spot that lets the <em>next</em> jump reach farthest.</p>`;
        },
      };
    }

    if (exId === 'gas') {
      let gas;
      let cost;
      let start;
      for (let tries = 0; tries < 500; tries++) {
        gas = Array.from({ length: n }, () => randInt(1, 8));
        cost = Array.from({ length: n }, () => randInt(1, 8));
        const total = gas.reduce((x, y) => x + y, 0) - cost.reduce((x, y) => x + y, 0);
        if (total < 0 || total > 3) continue;
        const ok = [];
        for (let s = 0; s < n; s++) { let t = 0; let fine = true; for (let q = 0; q < n; q++) { const i = (s + q) % n; t += gas[i] - cost[i]; if (t < 0) { fine = false; break; } } if (fine) ok.push(s); }
        if (ok.length === 1 && ok[0] >= 2) { start = ok[0]; break; }
      }
      const diff = gas.map((g, i) => g - cost[i]);
      const steps = [];
      let tank = 0;
      let st = 0;
      diff.forEach((d, i) => {
        tank += d;
        if (tank < 0) {
          steps.push({ line: 4, text: `Station ${i}: tank ${tank - d} + ${d} = ${tank} < 0. No start from ${st} to ${i} works, so try ${i + 1} with an empty tank.`, current: i, window: [st, i] });
          st = i + 1;
          tank = 0;
        } else {
          steps.push({ line: 2, text: `Station ${i}: gas ${gas[i]} - cost ${cost[i]} = ${d >= 0 ? '+' : ''}${d}, tank ${tank}.`, current: i, window: [st, i] });
        }
      });
      steps.push({ line: 6, text: `The total is not negative, so the last start standing, ${st}, makes it all the way round.`, current: st });
      return {
        prompt: 'Gas station: each cell is gas[i] - cost[i], the change in your tank driving from station i to the next. The total is not negative, so exactly one start works. Click it.',
        aside: `<p class="trace-ops">gas = <code>[${gas.join(', ')}]</code> cost = <code>[${cost.join(', ')}]</code></p>`,
        mode: 'sequence',
        slots: ['start'],
        ids: diff.map((_, i) => i),
        label: (id) => `station ${id}`,
        svg: (view) => arraySvg(diff.map(d => (d > 0 ? `+${d}` : String(d))), view),
        expected: [start],
        fixedLength: true,
        code: GR_CODE.gas,
        steps: () => steps,
        right: () => `Station ${start}. One pass: whenever the tank goes negative, every start so far is ruled out at once, because each of them arrived at that point with no more gas than the current start had.`,
        wrong: (k, ans) => {
          let t = 0;
          let fail = -1;
          for (let q = 0; q < n; q++) { const i = (ans[0] + q) % n; t += diff[i]; if (t < 0) { fail = i; break; } }
          return `<p class="trace-bad">It is station ${start}.</p>
            <p class="trace-why">Starting at ${ans[0]}, the tank goes negative after station ${fail}. ${diff[ans[0]] === Math.max(...diff) ? 'The biggest single gain is not always the right start; what matters is never dipping below zero on the way round.' : 'Use "Show me" to see each failed start ruled out.'}</p>`;
        },
      };
    }

    if (exId === 'partition') {
      let s;
      let ends;
      for (let tries = 0; tries < 300; tries++) {
        s = Array.from({ length: n + 2 }, () => pick([...'abcdefg']));
        const last = {};
        s.forEach((ch, i) => { last[ch] = i; });
        ends = [];
        let end = 0;
        s.forEach((ch, i) => { end = Math.max(end, last[ch]); if (i === end) ends.push(i); });
        if (ends.length >= 3) break;
      }
      const last = {};
      s.forEach((ch, i) => { last[ch] = i; });
      const steps = [];
      let start = 0;
      let end = 0;
      s.forEach((ch, i) => {
        const was = end;
        end = Math.max(end, last[ch]);
        steps.push(i === end
          ? { line: 5, text: `i = ${i} reaches the part's end: nothing in ${start}..${i} appears later. Cut here.`, current: i, window: [start, i], out: ends.filter(e => e <= i) }
          : { line: 3, text: `'${ch}' last appears at ${last[ch]}${last[ch] > was ? `, so the part must stretch to ${end}` : ''}.`, current: i, window: [start, end], heldIds: [last[ch]], out: ends.filter(e => e < i) });
        if (i === end) start = i + 1;
      });
      return {
        prompt: 'Partition labels: cut the string into as many parts as possible so each letter appears in only one part. Click the <strong>last index of each part</strong>, in order.',
        mode: 'sequence',
        ids: s.map((_, i) => i),
        label: cellLabel(s),
        svg: (view) => arraySvg(s, view),
        expected: ends,
        fixedLength: false,
        code: GR_CODE.partition,
        outLabel: 'Cuts after',
        steps: () => steps,
        right: () => `Parts: ${ends.map((e, i) => s.slice(i ? ends[i - 1] + 1 : 0, e + 1).join('')).join(' | ')}. Cutting as soon as possible is safe because a later cut could only merge parts.`,
        wrong: (k, ans) => {
          if (k === ends.length) return `<p class="trace-bad">${said(k)} That was the end of the string.</p>`;
          const st = k ? ends[k - 1] + 1 : 0;
          const part = s.slice(st, ends[k] + 1);
          const tooEarly = k < ans.length && ans[k] < ends[k];
          const culprit = tooEarly ? s.slice(st, ans[k] + 1).find(ch => last[ch] > ans[k]) : null;
          return `<p class="trace-bad">${said(k)} The next part ends at ${ends[k]} ('${part.join('')}').</p>
            <p class="trace-why">${culprit ? `Cutting at ${ans[k]} would split '${culprit}', which appears again at ${last[culprit]}.` : k < ans.length ? `Every letter in ${st}..${ends[k]} is finished by ${ends[k]}, so the part can end there; waiting longer would give fewer parts.` : `Keep going: the next part ends at ${ends[k]}.`}</p>`;
        },
      };
    }

    // coins: does largest-first greedy give the fewest coins?
    const sets = [[1, 5, 10, 25], [1, 2, 5, 10], [1, 3, 4], [1, 5, 6, 9], [1, 7, 10], [1, 4, 6]];
    let coins;
    let amount;
    let opt;
    let gr;
    let grPicks;
    for (let tries = 0; tries < 300; tries++) {
      coins = pick(sets);
      amount = randInt(6, size === 'small' ? 15 : 30);
      const d = [0, ...Array(amount).fill(Infinity)];
      for (let x = 1; x <= amount; x++) coins.forEach(c => { if (c <= x) d[x] = Math.min(d[x], d[x - c] + 1); });
      opt = d[amount];
      let left = amount;
      grPicks = [];
      [...coins].reverse().forEach(c => { while (left >= c) { left -= c; grPicks.push(c); } });
      gr = grPicks.length;
      if (Math.random() < 0.5 ? gr !== opt : gr === opt) break;
    }
    const opts = [...new Set([opt, gr, opt + 1, Math.max(1, opt - 1)])].slice(0, 4).sort((x, y) => x - y);
    const canonical = coins.join() === '1,5,10,25' || coins.join() === '1,2,5,10';
    return {
      prompt: `Coins <strong>${coins.join(', ')}</strong>, unlimited of each. What is the <strong>fewest</strong> coins that make <strong>${amount}</strong>?`,
      mode: 'choice',
      svg: (view) => arraySvg(coins, view),
      sequence: [`amount ${amount}`],
      choices: opts.map(v => ({ id: String(v), label: String(v) })),
      answerId: String(opt),
      explain: (picked) => `<p class="${picked === String(opt) ? 'trace-ok' : 'trace-bad'}"><strong>${picked === String(opt) ? 'Right' : 'Not quite'}:</strong> ${opt}.</p>
        <p class="trace-why">Largest coin first takes ${grPicks.join(' + ')}: ${gr} coin${gr === 1 ? '' : 's'}. ${gr === opt ? (canonical ? 'For this coin system greedy is always optimal (it is "canonical"), but that is a property you have to prove, not assume.' : 'Here greedy happens to match, but for this coin system it is not always optimal, so only DP is safe.') : `That is not optimal: a bigger coin early blocks a better combination. Greedy needs a proof; when there is none, use the DP (Coin Change, 322).`}</p>`,
    };
  }

  KITS.greedy = {
    title: 'Greedy',
    exercises: [
      { id: 'jump', name: 'Jump game: fewest jumps' },
      { id: 'gas', name: 'Gas station: where to start' },
      { id: 'partition', name: 'Partition labels' },
      { id: 'coins', name: 'When greedy fails: coins' },
    ],
    make: greedyQuestion,
  };

  // ===========================================================================
  // Bit manipulation
  // ===========================================================================

  // bitsSvg draws rows of W bits, most significant on the left, with each
  // column's place value underneath. Only the row marked answer is
  // clickable; its cells' ids are bit positions.
  function bitsSvg(W, rows, view) {
    const w = 40;
    const x0 = 70;
    const rowH = 46;
    const top = 14;
    let body = '';
    rows.forEach((row, r) => {
      const y = top + r * rowH;
      body += `<text class="ilabel bitname" x="${x0 - 12}" y="${y + 18}" dy="0.35em">${esc(row.name)}</text>`;
      for (let col = 0; col < W; col++) {
        const pos = W - 1 - col;
        const x = x0 + col * w;
        const cls = ['tnode', 'acell'];
        let attrs = '';
        let text = row.bits ? String(row.bits[pos]) : '';
        if (row.answer) {
          const shown = view.step?.bits;
          if (shown) text = String(shown[pos]);
          const idx = view.picked.get(pos);
          if (idx || (view.step && shown && shown[pos] === 1)) cls.push('picked');
          if (view.wrong === pos) cls.push('wrong');
          const clickable = view.clickable && !idx;
          if (clickable) cls.push('clickable');
          if (!view.step && idx) text = '1';
          attrs = `data-id="${pos}" ${clickable ? `role="button" tabindex="0" aria-label="Bit ${pos}"` : ''}`;
        } else {
          if (row.bits && row.bits[pos] === 1) cls.push('one');
          if (view.step?.col === pos) cls.push('current');
        }
        body += `<g class="${cls.join(' ')}" ${attrs}><rect x="${x + 2}" y="${y}" width="${w - 6}" height="36" rx="5" /><text x="${x + w / 2 - 1}" y="${y + 18}" dy="0.35em">${esc(text)}</text></g>`;
      }
    });
    const by = top + rows.length * rowH + 4;
    for (let col = 0; col < W; col++) body += `<text class="idx" x="${x0 + col * w + w / 2 - 1}" y="${by}">${2 ** (W - 1 - col)}</text>`;
    return svgWrap(x0 + W * w + 20, by + 12, body, 'Bits');
  }

  const toBits = (n, W) => Array.from({ length: W }, (_, i) => (n >> i) & 1);
  const fromBits = (b) => b.reduce((s, v, i) => s + (v << i), 0);
  const binStr = (n, W) => n.toString(2).padStart(W, '0');

  function bitQuestion(exId, size) {
    const W = { small: 4, medium: 6, large: 8 }[size];
    const ones = (n) => toBits(n, W).map((v, i) => (v ? i : -1)).filter(i => i >= 0);
    const answerRow = { name: '?', answer: true };
    const base = (expectedN, prompt, rows, code, steps, right, why) => {
      const expected = ones(expectedN);
      return {
        prompt: `${prompt} Click every bit that is 1 in the answer (none if it is 0), then Check.`,
        mode: 'sequence',
        ids: Array.from({ length: W }, (_, i) => i),
        label: (id) => `2^${id}`,
        svg: (view) => bitsSvg(W, rows, view),
        expected,
        fixedLength: false,
        validate: (ans) => ({ ok: sameSet(ans, expected), k: 0 }),
        code,
        steps: () => steps,
        right: () => right,
        wrong: (k, ans) => `<p class="trace-bad">The answer is ${binStr(expectedN, W)} (${expectedN}); you made ${binStr(ans.reduce((s, i) => s | (1 << i), 0), W)}.</p><p class="trace-why">${why}</p>`,
      };
    };

    if (exId === 'read') {
      const n = randInt(1, 2 ** W - 1);
      const steps = [];
      let x = n;
      const got = Array(W).fill(0);
      for (let i = 0; x; i++) {
        got[i] = x % 2;
        steps.push({ line: 2, text: `${x} % 2 = ${x % 2}, so bit ${i} (worth ${2 ** i}) is ${x % 2}. Then ${x} // 2 = ${x >> 1}.`, bits: [...got], col: i });
        x >>= 1;
      }
      const row = { name: '?', answer: true };
      const q = base(n, `Write <strong>${n}</strong> in binary.`, [row], ['bits = []', 'while n:', '    bits.append(n % 2)   # lowest bit first', '    n //= 2'], steps,
        `${n} = ${ones(n).map(i => 2 ** i).reverse().join(' + ')} = ${binStr(n, W)}.`,
        `Each bit is a place value (1, 2, 4, 8, ...). Repeatedly take n % 2 for the lowest bit and halve n: ${steps.map(s => s.text.split(',')[0]).join('; ')}.`);
      return q;
    }

    if (exId === 'ops') {
      const op = pick(['&', '|', '^']);
      const a = randInt(1, 2 ** W - 1);
      const b = randInt(1, 2 ** W - 1);
      const r = op === '&' ? a & b : op === '|' ? a | b : a ^ b;
      const rule = { '&': 'AND is 1 only where both bits are 1', '|': 'OR is 1 where either bit is 1', '^': 'XOR is 1 where the bits differ' }[op];
      const steps = [];
      for (let i = W - 1; i >= 0; i--) {
        steps.push({ line: 0, text: `Bit ${i}: ${(a >> i) & 1} ${op} ${(b >> i) & 1} = ${(r >> i) & 1}.`, col: i, bits: toBits(r, W).map((v, j) => (j >= i ? v : '')) });
      }
      return base(r, `Compute <strong>a ${op} b</strong>, bit by bit.`, [{ name: `a = ${a}`, bits: toBits(a, W) }, { name: `b = ${b}`, bits: toBits(b, W) }, { name: `a ${op} b`, answer: true }],
        [`r = a ${op} b   # ${rule}`], steps, `${binStr(a, W)} ${op} ${binStr(b, W)} = ${binStr(r, W)} (${r}). ${cap(rule)}.`, `${cap(rule)}; each column is independent.`);
    }

    if (exId === 'lowbit') {
      let n;
      do { n = randInt(2, 2 ** W - 1); } while ((n & (n - 1)) === 0 || (n & 1));
      const kind = pick(['clear', 'keep']);
      const r = kind === 'clear' ? n & (n - 1) : n & -n;
      const low = Math.log2(n & -n);
      const steps = [
        { line: 0, text: `n - 1 = ${n - 1}: subtracting 1 turns the lowest 1 (bit ${low}) into 0 and every 0 below it into 1.`, col: low, bits: toBits(n - 1, W) },
        { line: 1, text: kind === 'clear' ? `n & (n - 1) keeps the bits above bit ${low}, where n and n - 1 agree, and loses bit ${low}: ${binStr(r, W)}.` : `-n is ~n + 1, which agrees with n only at bit ${low}; n & -n = ${binStr(r, W)}.`, col: low, bits: toBits(r, W) },
      ];
      return base(r, `n = <strong>${n}</strong>. Compute <strong>${kind === 'clear' ? 'n & (n - 1)' : 'n & -n'}</strong>.`,
        [{ name: `n = ${n}`, bits: toBits(n, W) }, { name: `n - 1`, bits: toBits(n - 1, W) }, { name: kind === 'clear' ? 'n & (n-1)' : 'n & -n', answer: true }],
        ['m = n - 1        # lowest 1 -> 0, the 0s below -> 1', kind === 'clear' ? 'r = n & m        # clears the lowest set bit' : 'r = n & -n       # keeps only the lowest set bit'], steps,
        kind === 'clear' ? `${binStr(r, W)}: the lowest set bit (bit ${low}) is gone. Loop this to count bits, one per set bit.` : `${binStr(r, W)}: only the lowest set bit (bit ${low}, worth ${2 ** low}) survives.`,
        kind === 'clear' ? `n & (n - 1) clears exactly the lowest 1. Everything above it is unchanged.` : 'n & -n isolates the lowest 1: in two\'s complement -n flips every bit above it.');
    }

    // shift
    const left = Math.random() < 0.5;
    let n;
    let k;
    do { n = randInt(1, 2 ** W - 1); k = randInt(1, Math.max(1, W >> 1)); } while (left && (n << k) >= 2 ** W);
    const r = left ? n << k : n >> k;
    const steps = [{ line: 0, text: left ? `Every bit moves ${k} place${k === 1 ? '' : 's'} left (towards bigger values), with 0s filling in on the right: ${n} × 2^${k} = ${r}.` : `Every bit moves ${k} place${k === 1 ? '' : 's'} right; the lowest ${k} bit${k === 1 ? ' falls' : 's fall'} off: ${n} // 2^${k} = ${r}.`, bits: toBits(r, W) }];
    return base(r, `n = <strong>${n}</strong>. Compute <strong>n ${left ? '<<' : '>>'} ${k}</strong>.`,
      [{ name: `n = ${n}`, bits: toBits(n, W) }, { name: `n ${left ? '<<' : '>>'} ${k}`, answer: true }],
      [left ? `r = n << ${k}   # times 2 ** ${k}` : `r = n >> ${k}   # floor-divide by 2 ** ${k}`], steps,
      `${binStr(r, W)} (${r}). ${left ? `Shifting left by ${k} multiplies by ${2 ** k}.` : `Shifting right by ${k} floor-divides by ${2 ** k}.`}`,
      left ? `Each 1 moves ${k} column${k === 1 ? '' : 's'} to the left.` : `Each 1 moves ${k} column${k === 1 ? '' : 's'} to the right, and the ones that pass bit 0 are dropped.`);
  }

  KITS['bit-manipulation'] = {
    title: 'Bit manipulation',
    exercises: [
      { id: 'read', name: 'Number to binary' },
      { id: 'ops', name: 'AND, OR, XOR' },
      { id: 'lowbit', name: 'The lowest set bit' },
      { id: 'shift', name: 'Shifts' },
    ],
    make: bitQuestion,
  };

  // ===========================================================================
  // Math & geometry
  // ===========================================================================

  // pointsSvg draws labelled points on a small integer grid, with an
  // optional arrow from the first point to the second.
  function pointsSvg(pts, view, { arrow = null } = {}) {
    const unit = 34;
    const pad = 30;
    const maxC = 8;
    const px = (x) => pad + x * unit;
    const py = (y) => pad + (maxC - y) * unit;
    let body = '';
    for (let v = 0; v <= maxC; v++) {
      body += `<line class="grid" x1="${px(v)}" y1="${py(0)}" x2="${px(v)}" y2="${py(maxC)}" /><line class="grid" x1="${px(0)}" y1="${py(v)}" x2="${px(maxC)}" y2="${py(v)}" />`;
      if (v % 2 === 0) body += `<text class="idx" x="${px(v)}" y="${py(0) + 16}">${v}</text><text class="idx" x="${px(0) - 14}" y="${py(v) + 4}">${v}</text>`;
    }
    if (arrow) {
      const [a, b] = arrow;
      const dx = b.x - a.x;
      const dy = b.y - a.y;
      const t = 3;
      body += `<clipPath id="trace-grid-clip"><rect x="${px(0)}" y="${py(maxC)}" width="${maxC * unit}" height="${maxC * unit}" /></clipPath>`;
      body += `<line class="sweep" clip-path="url(#trace-grid-clip)" x1="${px(a.x - dx * t * 4)}" y1="${py(a.y - dy * t * 4)}" x2="${px(a.x + dx * (t * 4 + 1))}" y2="${py(a.y + dy * (t * 4 + 1))}" />`;
      body += `<g class="edges"><line x1="${px(a.x)}" y1="${py(a.y)}" x2="${px(b.x)}" y2="${py(b.y)}" marker-end="url(#trace-arrow)" /></g>`;
    }
    pts.forEach(p => { body += nodeSvg(p.name, px(p.x), py(p.y), p.name, { ...view, clickable: false }, `<text class="idx" x="${px(p.x) + 24}" y="${py(p.y) - 14}">(${p.x}, ${p.y})</text>`); });
    return `<svg class="trace-tree" viewBox="0 0 ${px(maxC) + pad + 40} ${py(0) + pad}" style="width:100%;max-width:460px" role="img" aria-label="Points">
      <defs><marker id="trace-arrow" viewBox="0 0 10 10" refX="9" refY="5" markerWidth="7" markerHeight="7" orient="auto-start-reverse"><path d="M0,0 L10,5 L0,10 z" class="arrowhead" /></marker></defs>${body}</svg>`;
  }

  const MATH_CODE = {
    sieve: ['is_prime = [True] * (n + 1)', 'for p in range(2, isqrt(n) + 1):', '    if is_prime[p]:', '        for m in range(p * p, n + 1, p):', '            is_prime[m] = False'],
    spiral: ['top, bottom, left, right = 0, R - 1, 0, C - 1', 'while top <= bottom and left <= right:', '    walk right along top;    top += 1', '    walk down along right;   right -= 1', '    walk left along bottom;  bottom -= 1   (if rows remain)', '    walk up along left;      left += 1     (if columns remain)'],
    rotate: ['# 90 degrees clockwise: (r, c) -> (c, n - 1 - r)', 'for r in range(n):', '    for c in range(n):', '        out[c][n - 1 - r] = m[r][c]'],
  };

  function mathQuestion(exId, size) {
    if (exId === 'gcd') {
      let a;
      let b;
      let lines;
      for (let tries = 0; tries < 300; tries++) {
        const g = randInt(2, 9);
        a = g * randInt(3, size === 'small' ? 12 : 30);
        b = g * randInt(2, size === 'small' ? 10 : 25);
        if (a === b) continue;
        lines = [];
        let x = Math.max(a, b);
        let y = Math.min(a, b);
        while (y) { lines.push(`gcd(${x}, ${y}): ${x} % ${y} = ${x % y}`); [x, y] = [y, x % y]; }
        lines.push(`gcd(${x}, 0) = ${x}`);
        if (lines.length >= 3) break;
      }
      const gcd = (x, y) => (y ? gcd(y, x % y) : x);
      const ans = gcd(a, b);
      const firstRem = Math.max(a, b) % Math.min(a, b);
      const opts = [...new Set([ans, ans * 2, firstRem, ans > 2 ? ans - 1 : ans + 3].filter(v => v > 0))].slice(0, 4).sort((x, y) => x - y);
      return {
        prompt: `Use Euclid's algorithm: what is <strong>gcd(${a}, ${b})</strong>?`,
        mode: 'choice',
        svg: (view) => arraySvg([a, b], view),
        sequence: [`a = ${a}`, `b = ${b}`],
        choices: opts.map(v => ({ id: String(v), label: String(v) })),
        answerId: String(ans),
        explain: (picked) => `<p class="${picked === String(ans) ? 'trace-ok' : 'trace-bad'}"><strong>${picked === String(ans) ? 'Right' : 'Not quite'}:</strong> ${ans}.</p>
          <p class="trace-why">gcd(a, b) = gcd(b, a % b), because anything that divides a and b also divides a % b. The second number shrinks every step, so it ends at 0:</p>
          <ol class="trace-notes">${lines.map(t => `<li>${esc(t)}</li>`).join('')}</ol>`,
      };
    }

    if (exId === 'sieve') {
      const N = { small: 30, medium: 50, large: 80 }[size];
      const cols = 10;
      const R = Math.ceil(N / cols);
      const grid = Array.from({ length: R }, (_, r) => Array.from({ length: cols }, (_, c) => r * cols + c + 1));
      const P = pick([2, 3, 5, 7].filter(p => p * p <= N));
      const crossedBefore = new Set();
      [2, 3, 5, 7].filter(p => p < P).forEach(p => { for (let m = p * p; m <= N; m += p) crossedBefore.add(m); });
      const expected = [];
      for (let m = P * P; m <= N; m += P) if (!crossedBefore.has(m)) expected.push(m - 1);
      const idOf = (v) => v - 1;
      const dim = new Set([0, ...[...crossedBefore].map(idOf), ...grid.flat().filter(v => v > N).map(idOf)]);
      const steps = [];
      const crossed = new Set(dim);
      for (let m = P * P; m <= N; m += P) {
        const already = crossedBefore.has(m);
        crossed.add(idOf(m));
        steps.push({ line: 4, text: already ? `${m} = ${P} × ${m / P} was already crossed out by a smaller prime.` : `${m} = ${P} × ${m / P}: cross it out.`, current: idOf(m), dim: new Set(crossed) });
      }
      steps.unshift({ line: 3, text: `Start at ${P} × ${P} = ${P * P}: every smaller multiple of ${P} has a smaller prime factor and is already crossed out.`, current: idOf(P), dim });
      return {
        prompt: `Sieve of Eratosthenes up to ${N}. ${P === 2 ? 'This is the first pass' : `The passes for ${[2, 3, 5, 7].filter(x => x < P).join(' and ')} are done (crossed out)`}. Now p = <strong>${P}</strong>: click every number this pass <strong>newly</strong> crosses out, then Check.`,
        mode: 'sequence',
        ids: grid.flat().filter(v => v <= N).map(idOf).filter(id => !dim.has(id)),
        label: (id) => String(id + 1),
        svg: (view) => tableSvg(grid.map(row => row.map(v => (v <= N ? v : ''))), view, { dim }),
        expected,
        fixedLength: false,
        validate: (ans) => ({ ok: sameSet(ans, expected), k: 0 }),
        code: MATH_CODE.sieve,
        outLabel: 'Crossed',
        steps: () => steps,
        right: () => `${expected.map(i => i + 1).join(', ')}. Starting at p² is what makes the sieve O(n log log n): smaller multiples are already gone.`,
        wrong: (k, ans) => {
          const extra = ans.find(i => !expected.includes(i));
          const missing = expected.find(i => !ans.includes(i));
          return `<p class="trace-bad">This pass newly crosses out ${expected.map(i => i + 1).join(', ')}.</p>
            <p class="trace-why">${extra !== undefined ? `${extra + 1} ${(extra + 1) % P ? `is not a multiple of ${P}` : (extra + 1) < P * P ? `is below ${P * P}; it has a smaller prime factor` : 'was already crossed out'}. ` : ''}${missing !== undefined ? `${missing + 1} = ${P} × ${(missing + 1) / P} is still uncrossed. ` : ''}The pass crosses p², p² + p, p² + 2p, ...</p>`;
        },
      };
    }

    if (exId === 'spiral') {
      const [R, C] = { small: [3, 3], medium: [3, 4], large: [4, 4] }[size];
      const vals = distinctValues(R * C, 99);
      const grid = Array.from({ length: R }, (_, r) => vals.slice(r * C, r * C + C));
      const order = [];
      const steps = [];
      let top = 0;
      let bottom = R - 1;
      let left = 0;
      let right = C - 1;
      const push = (r, c, line, text) => { order.push(r * C + c); steps.push({ line, text, current: r * C + c, out: [...order] }); };
      while (top <= bottom && left <= right) {
        for (let c = left; c <= right; c++) push(top, c, 2, `Right along row ${top}.`);
        top++;
        for (let r = top; r <= bottom; r++) push(r, right, 3, `Down column ${right}.`);
        right--;
        if (top <= bottom) { for (let c = right; c >= left; c--) push(bottom, c, 4, `Left along row ${bottom}.`); bottom--; }
        if (left <= right) { for (let r = bottom; r >= top; r--) push(r, left, 5, `Up column ${left}.`); left++; }
      }
      return {
        prompt: 'Click every cell in <strong>spiral order</strong>: clockwise from the top-left, layer by layer.',
        mode: 'sequence',
        ids: order.slice().sort((x, y) => x - y),
        label: (id) => String(grid[Math.floor(id / C)][id % C]),
        svg: (view) => tableSvg(grid, view),
        expected: order,
        fixedLength: true,
        code: MATH_CODE.spiral,
        outLabel: 'Visited',
        steps: () => steps,
        right: () => `${order.map(id => grid[Math.floor(id / C)][id % C]).join(', ')}. Four boundaries shrink by one after each side is walked.`,
        wrong: (k, ans) => `<p class="trace-bad">${said(k)} Next is ${grid[Math.floor(order[k] / C)][order[k] % C]}, not ${grid[Math.floor(ans[k] / C)][ans[k] % C]}.</p>
          <p class="trace-why">${steps[k].text} Each side stops one cell short of where the previous side started, because that boundary has already moved in.</p>`,
      };
    }

    if (exId === 'rotate') {
      const n = { small: 3, medium: 4, large: 5 }[size];
      const vals = distinctValues(n * n, 99);
      const grid = Array.from({ length: n }, (_, r) => vals.slice(r * n, r * n + n));
      let r0;
      let c0;
      do { r0 = randInt(0, n - 1); c0 = randInt(0, n - 1); } while (n % 2 && r0 === (n >> 1) && c0 === (n >> 1));
      const tr = c0;
      const tc = n - 1 - r0;
      const rotated = Array.from({ length: n }, (_, r) => Array.from({ length: n }, (_, c) => grid[n - 1 - c][r]));
      const steps = [{ line: 0, text: `Row ${r0} becomes column ${n - 1 - r0} (the top row becomes the right column), and column ${c0} becomes row ${c0}.`, current: r0 * n + c0, table: grid.map(row => row.map(String)) },
        { line: 3, text: `So (${r0}, ${c0}) moves to (${tr}, ${tc}). This is the whole matrix rotated.`, current: tr * n + tc, table: rotated.map(row => row.map(String)) }];
      return {
        prompt: `Rotate the matrix <strong>90° clockwise</strong>. Click the cell where <strong>${grid[r0][c0]}</strong> (row ${r0}, column ${c0}) ends up.`,
        mode: 'sequence',
        slots: [`${grid[r0][c0]} goes to`],
        ids: vals.map((_, i) => i),
        label: (id) => `(${Math.floor(id / n)}, ${id % n})`,
        svg: (view) => tableSvg(grid, view, { mark: r0 * n + c0 }),
        expected: [tr * n + tc],
        fixedLength: true,
        code: MATH_CODE.rotate,
        steps: () => steps,
        right: () => `(${tr}, ${tc}): (r, c) goes to (c, n - 1 - r). In place, it is a transpose followed by reversing each row.`,
        wrong: (k, ans) => {
          const r = Math.floor(ans[0] / n);
          const c = ans[0] % n;
          const ccw = r === n - 1 - c0 && c === r0;
          const tp = r === c0 && c === r0;
          return `<p class="trace-bad">It goes to (${tr}, ${tc}).</p><p class="trace-why">${ccw ? 'That is where a counter-clockwise rotation sends it. ' : tp ? 'That is the transpose only; a clockwise rotation also reverses each row. ' : ''}Clockwise, the top row becomes the right column: (r, c) → (c, n - 1 - r) = (${c0}, ${n - 1} - ${r0}).</p>`;
        },
      };
    }

    // orient: which side of p -> q is r on? (cross product sign)
    let p;
    let q;
    let r;
    let cross;
    for (let tries = 0; tries < 300; tries++) {
      p = { name: 'p', x: randInt(0, 4), y: randInt(0, 4) };
      q = { name: 'q', x: randInt(3, 8), y: randInt(2, 8) };
      if (p.x === q.x && p.y === q.y) continue;
      const g = (x, y) => (y ? g(y, x % y) : Math.abs(x));
      const dx = q.x - p.x;
      const dy = q.y - p.y;
      if (Math.random() < 0.3) {
        const s = g(Math.abs(dx), Math.abs(dy));
        const t = pick([-1, 2]);
        r = { name: 'r', x: p.x + (dx / s) * t, y: p.y + (dy / s) * t };
      } else r = { name: 'r', x: randInt(0, 8), y: randInt(0, 8) };
      if (r.x < 0 || r.x > 8 || r.y < 0 || r.y > 8 || (r.x === p.x && r.y === p.y) || (r.x === q.x && r.y === q.y)) continue;
      cross = dx * (r.y - p.y) - dy * (r.x - p.x);
      break;
    }
    const ans = cross > 0 ? 'left' : cross < 0 ? 'right' : 'on';
    const names = { left: 'Left (counter-clockwise)', right: 'Right (clockwise)', on: 'On the line (collinear)' };
    return {
      prompt: 'Walking from <strong>p</strong> to <strong>q</strong>, is <strong>r</strong> on your left, on your right, or on the line?',
      mode: 'choice',
      svg: (view) => pointsSvg([p, q, r], view, { arrow: [p, q] }),
      sequence: [`p (${p.x}, ${p.y})`, `q (${q.x}, ${q.y})`, `r (${r.x}, ${r.y})`],
      choices: ['left', 'right', 'on'].map(id => ({ id, label: names[id] })),
      answerId: ans,
      explain: (picked) => `<p class="${picked === ans ? 'trace-ok' : 'trace-bad'}"><strong>${picked === ans ? 'Right' : 'Not quite'}:</strong> ${names[ans]}.</p>
        <p class="trace-why">cross = (q.x - p.x)(r.y - p.y) - (q.y - p.y)(r.x - p.x) = (${q.x - p.x})(${r.y - p.y}) - (${q.y - p.y})(${r.x - p.x}) = ${cross}. Positive is left, negative is right, zero is collinear. Integers only, so no floating-point slope errors and no division by zero on vertical lines.</p>`,
    };
  }

  KITS['math-geometry'] = {
    title: 'Math & geometry',
    exercises: [
      { id: 'gcd', name: 'Euclid: gcd' },
      { id: 'sieve', name: 'Sieve: one pass' },
      { id: 'spiral', name: 'Spiral order' },
      { id: 'rotate', name: 'Rotate 90°' },
      { id: 'orient', name: 'Left, right or on?' },
    ],
    make: mathQuestion,
  };

  // ===========================================================================
  // Advanced graphs
  // ===========================================================================

  const AG_CODE = {
    prim: ['in_tree = {start}', 'heap = [(w, start, v) for v, w in adj[start]]', 'while len(in_tree) < n:', '    w, u, v = heappop(heap)', '    if v in in_tree: continue   # both ends inside', '    in_tree.add(v); total += w', '    for x, wx in adj[v]:', '        if x not in in_tree:', '            heappush(heap, (wx, v, x))'],
    kruskal: ['parent = {u: u for u in nodes}', 'for w, u, v in sorted(edges):', '    ru, rv = find(u), find(v)', '    if ru == rv: continue       # would close a cycle', '    parent[ru] = rv             # union the components', '    tree.append((u, v)); total += w', '    if len(tree) == n - 1: break'],
    floyd: ['for k in range(n):        # pass k: via k', '    for i in range(n):', '        for j in range(n):', '            if dist[i][k] + dist[k][j] < dist[i][j]:', '                dist[i][j] = dist[i][k] + dist[k][j]'],
    hierholzer: ['def visit(u):', '    while out[u]:              # smallest first', '        v = out[u].pop(0)', '        visit(v)', '    route.append(u)', 'visit(start)', 'return route[::-1]'],
    flights: ['cost = [inf] * n; cost[src] = 0', 'for _ in range(k + 1):        # k stops = k + 1 flights', '    nxt = cost[:]              # freeze last round', '    for u, v, w in flights:', '        if cost[u] + w < nxt[v]:', '            nxt[v] = cost[u] + w', '    cost = nxt', 'return cost[dst] if cost[dst] < inf else -1'],
  };

  // distinctWeights gives g's edges the weights 1..m in random order, so the
  // minimum spanning tree is unique, and names each edge e0, e1, ...
  function distinctWeights(g) {
    const ws = shuffle(Array.from({ length: g.edges.length }, (_, i) => i + 1));
    g.edges.forEach((e, i) => { e.w = ws[i]; e.id = `e${i}`; });
    g.adj = new Map(g.nodes.map(nd => [nd.id, []]));
    g.edges.forEach(({ u, v, w }) => {
      g.adj.get(u).push({ v, w });
      if (!g.directed) g.adj.get(v).push({ v: u, w });
    });
    g.adj.forEach(list => list.sort((a, b) => a.v - b.v));
    return g;
  }

  const edgeOf = (g, u, v) => g.edges.find(e => (e.u === u && e.v === v) || (!g.directed && e.u === v && e.v === u));
  const edgeName = (e) => `${e.u}–${e.v} (${e.w})`;

  function primRun(g, start) {
    const inTree = new Set([start]);
    const order = [start];
    const tree = new Set();
    const via = new Map();
    const heap = [];
    const steps = [];
    let total = 0;
    const push = (u) => {
      g.adj.get(u).forEach(({ v, w }) => { if (!inTree.has(v)) heap.push({ w, u, v }); });
      heap.sort((a, b) => a.w - b.w);
    };
    const snap = (current, line, text, edge) => steps.push({
      current, line, text, edge, out: order.slice(), tree: new Set(tree),
      held: heap.map(e => `${e.u}–${e.v}:${e.w}`), heldIds: [], heldLabel: 'Heap, cheapest first (edge:weight)',
    });
    push(start);
    snap(start, 1, `Start the tree at ${start} and push every edge leaving it.`);
    while (order.length < g.n) {
      const e = heap.shift();
      if (inTree.has(e.v)) {
        snap(e.u, 4, `Pop ${e.u}–${e.v} (${e.w}). ${e.v} joined the tree after this edge was pushed, so both ends are inside and it would close a cycle. Skip it.`, [e.u, e.v]);
        continue;
      }
      inTree.add(e.v);
      order.push(e.v);
      tree.add(edgeOf(g, e.u, e.v).id);
      via.set(e.v, e);
      total += e.w;
      push(e.v);
      snap(e.v, 5, `Pop ${e.u}–${e.v} (${e.w}), the cheapest edge leaving the tree: ${e.v} joins, total ${total}. Then push ${e.v}'s edges to nodes still outside.`, [e.u, e.v]);
    }
    snap(null, 2, `All ${g.n} nodes are in the tree: ${g.n - 1} edges, total weight ${total}.`);
    return { steps, order, via, total };
  }

  function kruskalRun(g) {
    const parent = new Map(g.nodes.map(nd => [nd.id, nd.id]));
    const find = (x) => { while (parent.get(x) !== x) x = parent.get(x); return x; };
    const sorted = g.edges.slice().sort((a, b) => a.w - b.w);
    const tree = new Set();
    const order = [];
    const rejected = [];
    const steps = [];
    let total = 0;
    const groups = () => {
      const m = new Map();
      g.nodes.map(nd => nd.id).sort((a, b) => a - b).forEach(u => {
        const r = find(u);
        if (!m.has(r)) m.set(r, []);
        m.get(r).push(u);
      });
      return [...m.values()].map(c => `{${c.join(', ')}}`);
    };
    const snap = (line, text, edge) => steps.push({
      line, text, edge, out: order.slice(), tree: new Set(tree),
      held: groups(), heldIds: [], heldLabel: 'Components',
    });
    snap(0, `Every node starts as its own component. Edges in weight order: ${sorted.map(e => e.w).join(', ')}.`);
    for (const e of sorted) {
      const ru = find(e.u);
      const rv = find(e.v);
      if (ru === rv) {
        rejected.push(e.id);
        snap(3, `${edgeName(e)}: ${e.u} and ${e.v} are already in the same component, so this edge would close a cycle. Skip it.`, [e.u, e.v]);
        continue;
      }
      parent.set(ru, rv);
      tree.add(e.id);
      order.push(e.id);
      total += e.w;
      snap(5, `${edgeName(e)} joins two different components: add it and merge them. Total ${total}.`, [e.u, e.v]);
      if (order.length === g.n - 1) break;
    }
    snap(6, `${g.n - 1} edges for ${g.n} nodes: the tree is complete, total weight ${total}. Later edges can only close cycles.`);
    return { steps, order, rejected, total };
  }

  // flightsRun is Bellman-Ford limited to k + 1 rounds. With copy it reads
  // only last round's prices (the right way); without, one round can chain
  // several flights. rows holds the prices after each round.
  function flightsRun(g, flights, src, rounds, copy = true) {
    const ids = g.nodes.map(nd => nd.id);
    let cost = new Map(ids.map(u => [u, u === src ? 0 : Infinity]));
    const rows = [new Map(cost)];
    const steps = [];
    const fmt = (c) => (c === Infinity ? '∞' : c);
    const snap = (line, text, shown, frozen, edge) => steps.push({
      line, text, edge, dist: Object.fromEntries(shown),
      held: ids.filter(u => frozen.get(u) < Infinity).sort((a, b) => a - b).map(u => `${u}:${frozen.get(u)}`), heldIds: [], heldLabel: 'Last round (frozen)',
    });
    snap(0, `With no flights you can only be at ${src}, for free.`, cost, cost);
    for (let r = 1; r <= rounds; r++) {
      const nxt = copy ? new Map(cost) : cost;
      snap(2, `Round ${r}: copy the prices. Reading only the copy means every route found this round uses at most ${r} flight${r === 1 ? '' : 's'}.`, nxt, cost);
      flights.forEach(({ u, v, w }) => {
        if (cost.get(u) === Infinity) return;
        const c = cost.get(u) + w;
        if (c < nxt.get(v)) {
          const old = nxt.get(v);
          nxt.set(v, c);
          snap(5, `${u} → ${v}: ${cost.get(u)} + ${w} = ${c}, better than ${fmt(old)}: nxt[${v}] = ${c}.`, nxt, cost, [u, v]);
        } else {
          snap(4, `${u} → ${v}: ${cost.get(u)} + ${w} = ${c}, no better than ${fmt(nxt.get(v))}.`, nxt, cost, [u, v]);
        }
      });
      cost = nxt;
      rows.push(new Map(cost));
      snap(6, `End of round ${r}: each price is the cheapest using at most ${r} flight${r === 1 ? '' : 's'}.`, cost, cost);
    }
    return { cost, rows, steps };
  }

  // circleGraph places n nodes round a circle (slightly jittered) and adds
  // the given directed edges, each named e0, e1, ...
  function circleGraph(n, edges, weighted) {
    const turn = Math.random() * Math.PI * 2;
    const nodes = Array.from({ length: n }, (_, i) => {
      const a = turn + (i * 2 * Math.PI) / n;
      return { id: i, x: Math.round(170 + 125 * Math.cos(a) + randInt(-8, 8)), y: Math.round(150 + 115 * Math.sin(a) + randInt(-8, 8)) };
    });
    const minX = Math.min(...nodes.map(nd => nd.x));
    const minY = Math.min(...nodes.map(nd => nd.y));
    nodes.forEach(nd => { nd.x += 40 - minX; nd.y += 40 - minY; });
    const es = edges.map(([u, v, w = 1], i) => ({ u, v, w, id: `e${i}` }));
    const adj = new Map(nodes.map(nd => [nd.id, []]));
    es.forEach(({ u, v, w }) => adj.get(u).push({ v, w }));
    adj.forEach(list => list.sort((a, b) => a.v - b.v));
    return { n, nodes, byId: new Map(nodes.map(nd => [nd.id, nd])), edges: es, adj, directed: true, weighted };
  }

  // floydPass runs passes 0..K-1 and then reports what pass K does.
  function floydPass(n, edges, K) {
    const d = Array.from({ length: n }, (_, i) => Array.from({ length: n }, (_, j) => (i === j ? 0 : Infinity)));
    edges.forEach(([u, v, w]) => { d[u][v] = Math.min(d[u][v], w); });
    for (let k = 0; k < K; k++) {
      for (let i = 0; i < n; i++) for (let j = 0; j < n; j++) if (d[i][k] + d[k][j] < d[i][j]) d[i][j] = d[i][k] + d[k][j];
    }
    const before = d.map(row => row.slice());
    const changed = [];
    const steps = [];
    const fmt = (c) => (c === Infinity ? '∞' : String(c));
    const show = () => d.map(row => row.map(fmt));
    const snap = (line, text, current, heldIds) => steps.push({ line, text, current, heldIds, held: [], table: show(), out: changed.slice() });
    snap(0, `Before pass ${K}, dist[i][j] is the cheapest path using ${K === 0 ? 'no node' : K === 1 ? 'only node 0' : `only nodes 0..${K - 1}`} in the middle. Pass ${K} lets paths go through ${K} as well.`, undefined, []);
    for (let i = 0; i < n; i++) {
      for (let j = 0; j < n; j++) {
        if (i === K || j === K || i === j) continue;
        if (d[i][K] === Infinity || d[K][j] === Infinity) continue;
        const via = d[i][K] + d[K][j];
        const cells = [i * n + K, K * n + j];
        if (via < d[i][j]) {
          const old = d[i][j];
          d[i][j] = via;
          changed.push(i * n + j);
          snap(4, `${i} → ${K} → ${j}: ${before[i][K]} + ${before[K][j]} = ${via}, better than ${fmt(old)}. dist[${i}][${j}] = ${via}.`, i * n + j, cells);
        } else {
          snap(3, `${i} → ${K} → ${j}: ${d[i][K]} + ${d[K][j]} = ${via}, no better than ${fmt(d[i][j])}.`, i * n + j, cells);
        }
      }
    }
    snap(0, changed.length
      ? `Pass ${K} done: ${changed.length} cell${changed.length === 1 ? '' : 's'} got smaller. Row ${K} and column ${K} never change in their own pass, and neither does the diagonal. Cells that skip the check had ∞ in row or column ${K}.`
      : `Pass ${K} changed nothing.`, undefined, []);
    return { before, changed, steps };
  }

  function hierholzerRun(g, start) {
    const out = new Map([...g.adj].map(([u, list]) => [u, list.map(x => x.v)]));
    const route = [];
    const stack = [];
    const used = new Set();
    const steps = [];
    const snap = (current, line, text, edge, shown = route) => steps.push({
      current, line, text, edge, out: shown.slice(), tree: new Set(used),
      held: stack.slice(), heldIds: stack.slice(), heldLabel: 'Call stack (bottom to top)',
    });
    (function visit(u) {
      stack.push(u);
      snap(u, 0, stack.length === 1 ? `visit(${start}): start the walk at ${start}.` : `visit(${u}).`);
      while (out.get(u).length) {
        const v = out.get(u).shift();
        used.add(g.edges.find(e => e.u === u && e.v === v).id);
        snap(u, 2, `${u} → ${v} is ${u}'s smallest unused edge: use it and walk on to ${v}.${out.get(u).length ? ` (${u} still has ${out.get(u).join(', ')} for later.)` : ''}`, [u, v]);
        visit(v);
      }
      route.push(u);
      stack.pop();
      snap(u, 4, `${u} has no unused edges left, so it is finished: append it to route. ${stack.length ? `Back in visit(${stack[stack.length - 1]}).` : ''}`);
    })(start);
    const path = route.slice().reverse();
    snap(null, 6, `Every edge is used. route lists the path backwards, so reverse it: ${path.join(' → ')}.`, undefined, path);
    return { steps, path, route };
  }

  // A random directed walk of m edges, never repeating an edge or using both
  // directions between two nodes, so it is an Eulerian path of its own graph.
  function eulerWalk(n, m) {
    for (let tries = 0; tries < 500; tries++) {
      const seen = new Set();
      const walk = [randInt(0, n - 1)];
      while (walk.length <= m) {
        const u = walk[walk.length - 1];
        const opts = Array.from({ length: n }, (_, v) => v).filter(v => v !== u && !seen.has(`${u},${v}`) && !seen.has(`${v},${u}`));
        if (!opts.length) break;
        const v = pick(opts);
        seen.add(`${u},${v}`);
        walk.push(v);
      }
      if (walk.length === m + 1) return walk;
    }
    return null;
  }

  // greedyStuck follows smallest unused edges from u until there are none.
  function greedyStuck(g, usedPairs, u) {
    const used = new Set(usedPairs);
    for (;;) {
      const next = g.adj.get(u).map(x => x.v).find(v => !used.has(`${u},${v}`));
      if (next === undefined) return { at: u, left: g.edges.length - used.size };
      used.add(`${u},${next}`);
      u = next;
    }
  }

  function floydQuestion(size) {
    const n = { small: 4, medium: 5, large: 5 }[size];
    let edges;
    let K;
    let run;
    const want = size === 'large' ? 3 : 2;
    for (let tries = 0; tries < 400; tries++) {
      const pairs = [];
      for (let u = 0; u < n; u++) for (let v = u + 1; v < n; v++) pairs.push(Math.random() < 0.5 ? [u, v] : [v, u]);
      edges = shuffle(pairs).slice(0, Math.round(n * (size === 'small' ? 1.3 : 1.5))).map(([u, v]) => [u, v, randInt(1, 9)]);
      K = randInt(0, n - 1);
      run = floydPass(n, edges, K);
      if (run.changed.length >= want || (tries > 300 && run.changed.length)) break;
    }
    const g = circleGraph(n, edges, true);
    const fmt = (c) => (c === Infinity ? '∞' : String(c));
    const grid = run.before.map(row => row.map(fmt));
    const expected = run.changed;
    const name = (id) => `${Math.floor(id / n)}→${id % n}`;
    return {
      prompt: `The table is <strong>dist</strong> ${K === 0 ? 'before any pass' : K === 1 ? 'after pass 0' : `after passes 0..${K - 1}`} of Floyd-Warshall (row = from, column = to). Click every cell that <strong>pass ${K}</strong> makes smaller, by routing through node ${K}.`,
      mode: 'sequence',
      ids: Array.from({ length: n * n }, (_, i) => i),
      label: name,
      svg: (view) => `<div class="trace-stack">${graphSvg(g, { ...view, picked: new Map(), clickable: false, current: undefined, held: new Set(), wrong: undefined, step: undefined, target: K })}${tableSvg(grid, view, { heads: true, mark: null })}</div>`,
      expected,
      code: AG_CODE.floyd,
      outLabel: 'Changed',
      validate: (ans) => {
        const k = ans.findIndex(id => !expected.includes(id));
        return { ok: k < 0 && ans.length === expected.length, k: k < 0 ? ans.length : k };
      },
      steps: () => run.steps,
      right: () => `Pass ${K} improves ${expected.map(name).join(', ')}.`,
      wrong: (k, ans) => {
        const b = run.before;
        if (k < ans.length) {
          const id = ans[k];
          const [i, j] = [Math.floor(id / n), id % n];
          const why = i === j ? 'The diagonal stays 0: no path back to yourself is cheaper than staying put.'
            : i === K || j === K ? `Row ${K} and column ${K} can't change in pass ${K}: going through ${K} on the way to or from ${K} adds nothing.`
              : b[i][K] === Infinity || b[K][j] === Infinity ? `There is no path ${b[i][K] === Infinity ? `${i} → ${K}` : `${K} → ${j}`} yet (∞), so going through ${K} doesn't help.`
                : `${i} → ${K} → ${j} costs ${b[i][K]} + ${b[K][j]} = ${b[i][K] + b[K][j]}, which doesn't beat ${fmt(b[i][j])}.`;
          return `<p class="trace-bad">${i}→${j} doesn't change in pass ${K}.</p><p class="trace-why">${why}</p>`;
        }
        const missed = expected.filter(id => !ans.includes(id));
        const ex = missed[0];
        const [i, j] = [Math.floor(ex / n), ex % n];
        return `<p class="trace-bad">Those all change, but you missed ${missed.map(name).join(', ')}.</p>
          <p class="trace-why">For example ${i} → ${K} → ${j} costs ${b[i][K]} + ${b[K][j]} = ${b[i][K] + b[K][j]}, less than ${fmt(b[i][j])}. Check every cell whose row has a number in column ${K} and whose column has a number in row ${K}.</p>`;
      },
    };
  }

  function hierholzerQuestion(size) {
    const [n, m] = { small: [5, 6], medium: [5, 8], large: [6, 10] }[size];
    let g;
    let run;
    let start;
    const wantStuck = Math.random() < 0.85;
    for (let tries = 0; tries < 2000; tries++) {
      const walk = eulerWalk(n, m);
      if (!walk) continue;
      g = circleGraph(n, walk.slice(1).map((v, i) => [walk[i], v]), false);
      start = walk[0];
      run = hierholzerRun(g, start);
      const stuck = greedyStuck(g, [], start).left > 0;
      if (stuck === wantStuck) break;
    }
    const path = run.path;
    return {
      prompt: `Find a path from <strong>${start}</strong> that uses every arrow exactly once, the way Reconstruct Itinerary wants it: whenever there is a choice, go to the <strong>smallest</strong> node that still lets you use every arrow. Click the nodes in order (a node can be clicked again).`,
      mode: 'sequence',
      repeats: true,
      ids: g.nodes.map(nd => nd.id),
      label: (id) => id,
      svg: (view) => graphSvg(g, { ...view, target: view.target ?? (view.step ? undefined : start) }),
      expected: path,
      fixedLength: true,
      code: AG_CODE.hierholzer,
      outLabel: 'route',
      steps: () => run.steps,
      right: () => `${path.join(' → ')}. Hierholzer's algorithm gets there without any look-ahead: walk smallest-first, and when a node runs out of edges, add it to the route and back up.`,
      wrong: (k, ans) => {
        if (k === 0) return `<p class="trace-bad">The path starts at ${start}.</p><p class="trace-why">${g.edges.filter(e => e.u === start).length > g.edges.filter(e => e.v === start).length ? `${start} has one more arrow out than in, so a path that uses every arrow must start there.` : `Every node has as many arrows in as out, so the path is a loop and could start anywhere; this one fixes the start at ${start}.`}</p>`;
        const prev = path[k - 1];
        const x = ans[k];
        const want = path[k];
        const usedPairs = path.slice(1, k).map((v, i) => `${path[i]},${v}`);
        const left = g.adj.get(prev).map(e => e.v).filter(v => !usedPairs.includes(`${prev},${v}`));
        let why;
        if (!left.includes(x)) {
          why = `There is no unused arrow ${prev} → ${x}. From ${prev} the unused arrows go to ${left.join(', ')}.`;
        } else if (x < want) {
          const s = greedyStuck(g, [...usedPairs, `${prev},${x}`], x);
          why = `Going ${prev} → ${x} now gets stuck at ${s.at} with ${s.left} arrow${s.left === 1 ? '' : 's'} unused. Hierholzer's algorithm does try ${x} first, but that walk ends at ${s.at}, so it lands at the end of the route; the rest of ${prev}'s arrows, starting with ${prev} → ${want}, come before it.`;
        } else {
          why = `${prev} → ${want} is smaller and still lets you use every arrow.`;
        }
        return `<p class="trace-bad">${said(k)} After ${prev} comes ${want}, not ${x}.</p><p class="trace-why">${why}</p>`;
      },
    };
  }

  function agQuestion(exId, size) {
    if (exId === 'dijkstra') return graphQuestion('dijkstra', size);
    if (exId === 'floyd') return floydQuestion(size);
    if (exId === 'hierholzer') return hierholzerQuestion(size);

    if (exId === 'prim') {
      const g = distinctWeights(makeGraph(size, { weighted: true }));
      const start = pick(g.nodes).id;
      const run = primRun(g, start);
      const seq = run.order;
      return {
        prompt: `Run <strong>Prim's algorithm</strong> from <strong>${start}</strong>. Click the nodes in the order they join the tree. Each step adds the cheapest edge with exactly one end in the tree.`,
        aside: adjacencyHtml(g),
        mode: 'sequence',
        ids: g.nodes.map(nd => nd.id),
        label: (id) => id,
        svg: (view) => graphSvg(g, { ...view, target: view.target ?? (view.step ? undefined : start) }),
        expected: seq,
        fixedLength: true,
        code: AG_CODE.prim,
        outLabel: 'In the tree',
        steps: () => run.steps,
        right: () => `Joined in order ${seq.join(', ')}, through ${seq.slice(1).map(v => edgeName(edgeOf(g, run.via.get(v).u, v))).join(', ')}. Total weight ${run.total}.`,
        wrong: (k, ans) => {
          if (k === 0) return `<p class="trace-bad">${said(0)} Prim starts with just ${start} in the tree.</p>`;
          const inTree = new Set(seq.slice(0, k));
          const best = run.via.get(seq[k]);
          const x = ans[k];
          const links = g.adj.get(x).filter(({ v }) => inTree.has(v)).sort((a, b) => a.w - b.w);
          const why = links.length
            ? `${x} could join through ${x}–${links[0].v} (${links[0].w}), but ${best.u}–${best.v} (${best.w}) is the cheapest edge leaving the tree.`
            : `${x} has no edge to the tree yet (${[...inTree].join(', ')}). Prim only grows the tree it has.`;
          const dij = dijkstraRun(g, start).order;
          const note = dij.join() === ans.join() && dij.join() !== seq.join()
            ? ' Your order is Dijkstra\'s: it compares the whole distance from the start, while Prim compares only the one edge that joins the tree.' : '';
          return `<p class="trace-bad">${said(k)} Next ${seq[k]} joins, not ${x}.</p><p class="trace-why">${why}${note}</p>`;
        },
      };
    }

    if (exId === 'kruskal') {
      let g;
      let run;
      const want = size === 'small' ? 1 : 2;
      for (let tries = 0; tries < 200; tries++) {
        g = distinctWeights(makeGraph(size, { weighted: true }));
        run = kruskalRun(g);
        if (run.rejected.length >= want) break;
      }
      const byId = new Map(g.edges.map(e => [e.id, e]));
      const seq = run.order;
      return {
        prompt: `Run <strong>Kruskal's algorithm</strong>. Click the edges it adds, in order, by clicking their weights. It takes edges cheapest first and skips any whose ends are already connected.`,
        mode: 'sequence',
        ids: g.edges.map(e => e.id),
        label: (id) => edgeName(byId.get(id)),
        svg: (view) => graphSvg(g, view, { edgeIds: true }),
        expected: seq,
        fixedLength: true,
        code: AG_CODE.kruskal,
        outLabel: 'In the tree',
        steps: () => run.steps,
        right: () => `The tree is ${seq.map(id => edgeName(byId.get(id))).join(', ')}, total weight ${run.total}. Skipped on the way: ${run.rejected.map(id => edgeName(byId.get(id))).join(', ')}.`,
        wrong: (k, ans) => {
          const x = byId.get(ans[k]);
          const next = byId.get(seq[k]);
          const parent = new Map(g.nodes.map(nd => [nd.id, nd.id]));
          const find = (a) => { while (parent.get(a) !== a) a = parent.get(a); return a; };
          seq.slice(0, k).forEach(id => { const e = byId.get(id); parent.set(find(e.u), find(e.v)); });
          const why = find(x.u) === find(x.v)
            ? `${x.u} and ${x.v} are already connected by the edges chosen so far, so ${edgeName(x)} would close a cycle.`
            : `${edgeName(x)} does join two components, but ${edgeName(next)} is cheaper and also joins two. Kruskal always takes the cheapest edge that does not close a cycle.`;
          return `<p class="trace-bad">${said(k)} Next comes ${edgeName(next)}, not ${edgeName(x)}.</p><p class="trace-why">${why}</p>`;
        },
      };
    }

    // flights: Bellman-Ford limited to k + 1 rounds (Cheapest Flights Within K Stops).
    // Mostly a trip the limit makes dearer; sometimes no trip at all, or one
    // the limit doesn't affect. The last complete candidate is kept if none fits.
    const roll = Math.random();
    const wantNone = roll < 0.15;
    const wantDiff = roll < 0.8;
    let c = null;
    for (let tries = 0; tries < 300; tries++) {
      const g0 = makeGraph(size, { directed: true, weighted: true });
      const fl = g0.edges.slice().sort((a, b) => a.u - b.u || a.v - b.v);
      const s0 = pick(g0.nodes).id;
      const k0 = randInt(0, size === 'small' ? 1 : 2);
      const all = flightsRun(g0, fl, s0, g0.n - 1);
      const reach = g0.nodes.map(nd => nd.id).filter(u => u !== s0 && all.cost.get(u) < Infinity);
      if (!reach.length) continue;
      const d0 = pick(reach);
      c = { g: g0, flights: fl, src: s0, dst: d0, k: k0, free: all, right: flightsRun(g0, fl, s0, k0 + 1), inPlace: flightsRun(g0, fl, s0, k0 + 1, false) };
      const r = c.right.cost.get(d0);
      if ((r === Infinity) === wantNone && (r !== all.cost.get(d0)) === wantDiff) break;
    }
    const { g, flights, src, dst, k, free, right, inPlace } = c;
    const val = (c) => (c === Infinity ? -1 : c);
    const ans = val(right.cost.get(dst));
    const unlimited = val(free.cost.get(dst));
    const chained = val(inPlace.cost.get(dst));
    const needed = free.rows.findIndex(row => row.get(dst) === free.cost.get(dst));
    const opts = [...new Set([ans, unlimited, chained, -1, ans === -1 ? unlimited + 2 : ans + randInt(1, 3)])].slice(0, 4).sort((a, b) => a - b);
    const ids = g.nodes.map(nd => nd.id).sort((a, b) => a - b);
    const fmt = (c) => (c === Infinity ? '∞' : c);
    const table = tableHtml(right.rows.map(row => ids.map(u => fmt(row.get(u)))), (r) => `${r} flight${r === 1 ? '' : 's'}`, ids);
    const stops = `${k} stop${k === 1 ? '' : 's'}`;
    return {
      prompt: `Arrows are one-way flights and the numbers are prices. What is the cheapest trip from <strong>${src}</strong> to <strong>${dst}</strong> with at most <strong>${stops}</strong> (${k + 1} flight${k ? 's' : ''})? Answer -1 if there is none.`,
      aside: `<p class="trace-notes">The flights are relaxed in this order: ${flights.map(e => `${e.u}→${e.v} (${e.w})`).join(', ')}.</p>`,
      mode: 'choice',
      svg: (view) => graphSvg(g, { ...view, current: view.step ? view.current : src, target: dst }),
      sequence: [`${src} → ${dst}`, `≤ ${stops}`],
      choices: opts.map(v => ({ id: String(v), label: v === -1 ? '-1 (no route)' : String(v) })),
      answerId: String(ans),
      code: AG_CODE.flights,
      steps: () => right.steps,
      explain: (picked) => {
        const notes = [];
        if (unlimited !== ans) notes.push(`Without the limit the cheapest is ${unlimited}, but that route needs ${needed} flights.`);
        if (picked !== String(ans) && picked === String(chained) && chained !== ans) notes.push(`${chained} is what you get by updating the prices in place instead of copying them: one round then chains several flights, so the stop limit leaks.`);
        return `<p class="${picked === String(ans) ? 'trace-ok' : 'trace-bad'}"><strong>${picked === String(ans) ? 'Right' : 'Not quite'}:</strong> ${ans}.</p>
          <p class="trace-why">Bellman-Ford runs one round per flight allowed. Prices after each round:</p>${table}
          ${notes.length ? `<p class="trace-why">${notes.join(' ')}</p>` : ''}`;
      },
    };
  }

  KITS['advanced-graphs'] = {
    title: 'Advanced graphs',
    exercises: [
      { id: 'prim', name: "Prim's MST" },
      { id: 'kruskal', name: "Kruskal's MST" },
      { id: 'dijkstra', name: 'Dijkstra' },
      { id: 'flights', name: 'K stops: Bellman-Ford' },
      { id: 'floyd', name: 'Floyd-Warshall: one pass' },
      { id: 'hierholzer', name: 'Hierholzer: Euler path' },
    ],
    make: agQuestion,
  };

  // ===========================================================================
  // Core: progress, checking, rendering
  // ===========================================================================

  function key(cat, id) { return `trace:${cat}:${id}`; }

  function stats(cat, id) {
    return (typeof practiceState !== 'undefined' && practiceState.progress[key(cat, id)]) || {};
  }

  function record(correct) {
    const k = key(S.cat, S.exId);
    const prev = practiceState.progress[k] || {};
    const streak = correct ? (prev.streak || 0) + 1 : 0;
    const data = {
      traceAttempts: (prev.traceAttempts || 0) + 1,
      traceCorrect: (prev.traceCorrect || 0) + (correct ? 1 : 0),
      streak,
      best: Math.max(prev.best || 0, streak),
      mastered: !!prev.mastered || streak >= STREAK_TO_MASTER,
      lastAttempt: Date.now(),
    };
    if (typeof updateProblemProgress === 'function') updateProblemProgress(k, data);
  }

  function summary(cat) {
    const list = KITS[cat]?.exercises || [];
    const mastered = list.filter(e => stats(cat, e.id).mastered).length;
    const names = list.map(e => e.name.split(':')[0].toLowerCase()).join(', ');
    return mastered
      ? `${mastered} of ${list.length} mastered (${STREAK_TO_MASTER} right in a row). No code, a new one every time.`
      : `Click nodes to show what the algorithm does: ${names}. No code, a new one every time.`;
  }

  function streakHtml() {
    const st = stats(S.cat, S.exId);
    return `<p class="trace-streak">${st.streak ? `${st.streak} right in a row` : 'Correct on a later try (streaks count first tries).'}${st.mastered ? ' · mastered' : st.streak ? ` · ${STREAK_TO_MASTER - st.streak} more to master it` : ''}</p>`;
  }

  function newQuestion() {
    stopTimer();
    S.q = KITS[S.cat].make(S.exId, S.size);
    S.answer = [];
    S.checked = false;
    S.correct = false;
    S.firstTry = true;
    S.feedback = '';
    S.replay = null;
    S.picked = null;
    render();
  }

  function check() {
    const q = S.q;
    const res = q.validate ? q.validate(S.answer) : (() => {
      const k = firstMismatch(S.answer, q.expected);
      return { ok: k === q.expected.length && S.answer.length === q.expected.length, k };
    })();
    if (S.firstTry) record(res.ok);
    S.firstTry = false;
    S.checked = true;
    S.correct = res.ok;
    S.mismatch = res.k;
    S.feedback = res.ok
      ? `<p class="trace-ok"><strong>Right.</strong> ${q.right()}</p>${streakHtml()}`
      : q.wrong(res.k, S.answer);
    render();
  }

  function choose(id) {
    if (S.checked) return;
    S.picked = id;
    const ok = id === S.q.answerId;
    record(ok);
    S.checked = true;
    S.correct = ok;
    S.feedback = S.q.explain(id) + (ok ? streakHtml() : '');
    render();
  }

  function view() {
    const step = S.replay?.steps[S.replay.i];
    const picked = new Map();
    // A node clicked more than once (repeats) shows all its positions.
    (step ? step.out || [] : (S.q.mode === 'choice' ? [] : S.answer)).forEach((id, i) => picked.set(id, picked.has(id) ? `${picked.get(id)},${i + 1}` : i + 1));
    const wrongIdx = S.checked && !S.correct && !step && S.q.mode !== 'choice' ? (S.mismatch ?? -1) : -1;
    return {
      step,
      picked,
      current: step?.current,
      held: new Set(step?.heldIds || []),
      wrong: wrongIdx >= 0 ? S.answer[wrongIdx] : undefined,
      clickable: !step && !S.checked && S.q.mode !== 'choice',
      repeats: !!S.q.repeats,
      checked: S.checked,
      done: S.checked && S.correct,
    };
  }

  function codeHtml(line) {
    const lines = S.q.code;
    if (!lines) return '';
    return `<pre class="trace-code">${lines.map((l, i) => `<span class="${i === line ? 'hl' : ''}">${esc(l)}</span>`).join('\n')}</pre>`;
  }

  function chip(id, i) {
    const q = S.q;
    const bad = S.checked && !S.correct && i >= (S.mismatch ?? Infinity);
    const name = q.slots ? `<i>${esc(q.slots[i])}</i> ` : '';
    return `<b class="${bad ? 'wrong' : ''}">${name}${esc(q.label(id))}</b>`;
  }

  function sideHtml() {
    const q = S.q;
    if (S.replay) {
      const { steps, i } = S.replay;
      const step = steps[i];
      return `<div class="trace-replay">
        <div class="trace-replay-head">Step ${i + 1} of ${steps.length}</div>
        ${codeHtml(step.line)}
        <p class="trace-step-text">${esc(step.text)}</p>
        ${step.heldLabel ? `<div class="trace-held"><span>${esc(step.heldLabel)}</span>${step.held.length ? step.held.map(v => `<b>${esc(v)}</b>`).join('') : '<em>empty</em>'}</div>` : ''}
        ${step.pointers ? `<div class="trace-held"><span>Pointers</span>${Object.entries(step.pointers).map(([k, v]) => `<b><i>${esc(k)}</i> ${esc(q.label(v))}</b>`).join('')}</div>` : ''}
        ${step.out && !step.pointers ? `<div class="trace-held"><span>${q.outLabel || (q.code === GRAPH_CODE.dijkstra ? 'Settled' : 'Recorded')}</span>${step.out.length ? step.out.map(id => `<b class="out">${esc(q.label(id))}</b>`).join('') : '<em>nothing yet</em>'}</div>` : ''}
        <div class="trace-buttons">
          <button class="btn-small" data-act="back" ${i === 0 ? 'disabled' : ''}>Back</button>
          <button class="btn-small primary" data-act="next" ${i === steps.length - 1 ? 'disabled' : ''}>Next</button>
          <button class="btn-small" data-act="play">${S.replay.timer ? 'Pause' : 'Play'}</button>
          <button class="btn-small ghost" data-act="end-replay">Done</button>
        </div>
      </div>`;
    }
    if (q.mode === 'choice') {
      return `<p class="trace-prompt">${q.prompt}</p>
        ${q.aside || ''}
        <div class="trace-seq big">${q.sequence.map(v => `<b>${esc(v)}</b>`).join('')}</div>
        <div class="trace-buttons trace-choices">${q.choices.map(c => `<button class="btn-small ${S.picked === c.id ? (S.correct ? 'primary' : 'danger') : ''}" data-choice="${esc(c.id)}" ${S.checked ? 'disabled' : ''}>${esc(c.label)}</button>`).join('')}</div>
        <div class="trace-feedback">${S.feedback}</div>
        <div class="trace-buttons">${S.checked && q.steps ? '<button class="btn-small" data-act="show">Show me step by step</button>' : ''}<button class="btn-small ${S.checked ? 'primary' : 'ghost'}" data-act="new">New one</button></div>`;
    }
    const full = q.fixedLength ? S.answer.length === q.expected.length : S.answer.length > 0;
    const placeholder = q.slots ? `Click: ${q.slots.join(', then ')}` : 'Your order appears here';
    return `<p class="trace-prompt">${q.prompt}</p>
      ${q.aside || ''}
      <div class="trace-seq">${S.answer.length ? S.answer.map(chip).join('') : `<em>${esc(placeholder)}</em>`}</div>
      <div class="trace-buttons">
        ${S.checked ? '' : `<button class="btn-small" data-act="undo" ${S.answer.length ? '' : 'disabled'}>Undo</button>
        <button class="btn-small primary" data-act="check" ${full ? '' : 'disabled'}>Check</button>`}
        ${S.checked && !S.correct && !q.slots ? '<button class="btn-small primary" data-act="retry">Try again from there</button>' : ''}
        <button class="btn-small" data-act="show">Show me step by step</button>
        <button class="btn-small ${S.checked && S.correct ? 'primary' : 'ghost'}" data-act="new">New one</button>
      </div>
      <div class="trace-feedback">${S.feedback}</div>
      ${S.checked ? '' : `<details class="trace-code-peek"><summary>The code being traced</summary>${codeHtml(-1)}</details>`}`;
  }

  function tabsHtml() {
    return KITS[S.cat].exercises.map(e => {
      const st = stats(S.cat, e.id);
      return `<button class="trace-tab ${e.id === S.exId ? 'active' : ''}" data-ex="${esc(e.id)}">${esc(e.name)}${st.mastered ? ' <span class="trace-mastered">✓</span>' : ''}</button>`;
    }).join('');
  }

  function render() {
    const box = $('trace-drills');
    if (!box) return;
    box.querySelector('.trace-tabs').innerHTML = tabsHtml();
    box.querySelector('.trace-stage').innerHTML = S.q.svg(view());
    box.querySelector('.trace-side').innerHTML = sideHtml();
    box.querySelectorAll('.trace-size button').forEach(b => b.classList.toggle('active', b.dataset.size === S.size));
  }

  function pickNode(raw) {
    const id = S.q.ids.find(x => String(x) === raw);
    if (id === undefined || (!S.q.repeats && S.answer.includes(id)) || S.checked || S.replay) return;
    if (S.q.fixedLength && S.answer.length >= S.q.expected.length) return;
    S.answer.push(id);
    render();
  }

  function startReplay() {
    stopTimer();
    // Seeing the answer before checking counts as a miss for the streak.
    if (!S.checked && S.firstTry) { record(false); S.firstTry = false; }
    S.replay = { steps: S.q.steps(), i: 0, timer: null };
    render();
  }

  function stopTimer() {
    if (S.replay?.timer) { clearInterval(S.replay.timer); S.replay.timer = null; }
  }

  function onAction(act) {
    if (act === 'undo') { S.answer.pop(); render(); }
    else if (act === 'check') check();
    else if (act === 'new') newQuestion();
    else if (act === 'retry') {
      S.answer = S.answer.slice(0, S.mismatch);
      S.checked = false;
      S.feedback = '';
      render();
    } else if (act === 'show') startReplay();
    else if (act === 'back' && S.replay) { stopTimer(); S.replay.i = Math.max(0, S.replay.i - 1); render(); }
    else if (act === 'next' && S.replay) { stopTimer(); S.replay.i = Math.min(S.replay.steps.length - 1, S.replay.i + 1); render(); }
    else if (act === 'play' && S.replay) {
      if (S.replay.timer) { stopTimer(); render(); return; }
      if (S.replay.i >= S.replay.steps.length - 1) S.replay.i = 0;
      S.replay.timer = setInterval(() => {
        if (!S.replay || S.replay.i >= S.replay.steps.length - 1) { stopTimer(); render(); return; }
        S.replay.i++;
        render();
      }, 1100);
      render();
    } else if (act === 'end-replay') {
      stopTimer();
      S.replay = null;
      if (!S.checked) S.answer = [];
      render();
    }
  }

  // -------------------------------------------------------------------------
  // Open and close

  function open(cat = 'trees') {
    const kit = KITS[cat];
    if (!kit) return;
    close();
    if (S.cat !== cat || !kit.exercises.some(e => e.id === S.exId)) S.exId = kit.exercises[0].id;
    S.cat = cat;
    const back = document.createElement('div');
    back.id = 'trace-drills';
    back.className = 'trace-backdrop';
    back.innerHTML = `<div class="trace-panel" role="dialog" aria-modal="true" aria-label="Trace by hand">
      <header class="trace-header">
        <div>
          <span class="trace-kicker">Foundations · ${esc(kit.title)}</span>
          <h2>Trace by hand</h2>
        </div>
        <div class="trace-size" role="group" aria-label="Size">
          ${['small', 'medium', 'large'].map(s => `<button class="btn-small" data-size="${s}">${cap(s)}</button>`).join('')}
        </div>
        <button class="btn-small ghost trace-close" data-act="close" title="Close (Esc)">Close</button>
      </header>
      <nav class="trace-tabs"></nav>
      <div class="trace-body">
        <div class="trace-stage"></div>
        <div class="trace-side"></div>
      </div>
    </div>`;
    document.body.appendChild(back);

    back.addEventListener('click', (e) => {
      if (e.target === back) return close();
      const tab = e.target.closest('[data-ex]');
      if (tab) { S.exId = tab.dataset.ex; return newQuestion(); }
      const size = e.target.closest('[data-size]');
      if (size) { S.size = size.dataset.size; return newQuestion(); }
      const choice = e.target.closest('[data-choice]');
      if (choice) return choose(choice.dataset.choice);
      const act = e.target.closest('[data-act]');
      if (act) return act.dataset.act === 'close' ? close() : onAction(act.dataset.act);
      const node = e.target.closest('.tnode.clickable');
      if (node) pickNode(node.dataset.id);
    });
    back.addEventListener('keydown', (e) => {
      const node = e.target.closest?.('.tnode.clickable');
      if (node && (e.key === 'Enter' || e.key === ' ')) {
        e.preventDefault();
        pickNode(node.dataset.id);
      }
    });
    newQuestion();
  }

  function close() {
    const box = $('trace-drills');
    if (!box) return;
    stopTimer();
    box.remove();
    // Refresh the mastery summary on whatever is underneath.
    if (typeof renderProblemsList === 'function' && !$('practice-list')?.classList.contains('hidden')) renderProblemsList();
    if (typeof state !== 'undefined' && state.current?.category && typeof renderPracticeLadder === 'function' && !$('detail')?.classList.contains('hidden')) {
      renderPracticeLadder(state.current);
    }
  }

  document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape' && $('trace-drills')) close();
  });

  window.traceDrills = {
    has: (cat) => !!KITS[cat],
    summary,
    open,
    // Exposed for tests and the browser console.
    _internals: {
      KITS, state: S, makeTree, traverse, bstPath, treeDfsSteps, treeBfsSteps, treeBstSteps,
      makeGraph, bfsRun, dfsRun, topoRun, dijkstraRun, primRun, kruskalRun, flightsRun, distinctWeights, floydPass, hierholzerRun,
    },
  };
})();
