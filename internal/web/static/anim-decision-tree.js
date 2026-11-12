// Decision Tree Animator (chapter 12, storyboard "decision-tree").
// Draws the recursion tree of Combinations (n = 4, k = 3) as the search runs:
// chosen nodes appear on the way down, the current path is highlighted, a
// node greys out when its choice is undone, pruned branches are crossed out,
// and complete answers are collected in a results list.
// Loaded after app.js, which defines AlgorithmAnimator and animationRegistry.
class DecisionTreeAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.n = 4;
    this.k = 3;

    this.setCode(`def combine(n, k):
    result = []
    path = []

    def backtrack(start):
        if len(path) == k:              # record a complete answer
            result.append(path[:])
            return
        for num in range(start, n + 1):
            need = k - len(path)
            if n - num + 1 < need:      # prune: not enough numbers left
                break
            path.append(num)            # choose
            backtrack(num + 1)          # explore
            path.pop()                  # un-choose

    backtrack(1)
    return result`);
  }

  buildSteps() {
    const { n, k } = this;
    this.steps = [];
    this.nodes = [];

    // Mutable search state; every step stores a snapshot of it.
    const base = {};        // node id -> 'open' | 'done' | 'found' | 'pruned'
    const pathIds = [0];    // node ids from the root to the deepest chosen node
    const path = [];
    const result = [];
    let current = 0;
    let pruned = 0;

    const addNode = (parent, num, kind) => {
      const node = {
        id: this.nodes.length,
        parent,
        num,
        kind,
        depth: parent === null ? 0 : this.nodes[parent].depth + 1,
        path: parent === null ? [] : [...this.nodes[parent].path, num],
        children: [],
      };
      this.nodes.push(node);
      if (parent !== null) this.nodes[parent].children.push(node.id);
      return node.id;
    };

    const snap = (fields) => {
      this.steps.push({
        base: { ...base },
        pathIds: [...pathIds],
        current,
        path: [...path],
        result: result.map(r => [...r]),
        pruned,
        ...fields,
      });
    };

    addNode(null, null, 'node');
    base[0] = 'open';
    snap({ lineNums: [17], newId: 0, start: 1,
      state: 'Call backtrack(1) with path = []. This call is the root of the decision tree: nothing chosen yet.' });

    const backtrack = (start, nodeId) => {
      if (path.length === k) {
        result.push([...path]);
        base[nodeId] = 'found';
        snap({ lineNums: [6, 7, 8], start,
          state: `Record a complete answer: path [${path.join(', ')}] holds k = ${k} numbers, so save a copy and return.` });
        return;
      }
      for (let num = start; num <= n; num++) {
        const need = k - path.length;
        const left = n - num + 1;
        if (left < need) {
          const id = addNode(nodeId, num, 'pruned');
          base[id] = 'pruned';
          pruned++;
          snap({ lineNums: [10, 11, 12], newId: id, start, num, need, left,
            state: `Prune: path [${path.join(', ')}] still needs ${need} numbers, but only ${left} (${num}..${n}) ${left === 1 ? 'is' : 'are'} left. Cross out ${num} and break: larger numbers leave even fewer.` });
          break;
        }
        const id = addNode(nodeId, num, 'node');
        base[id] = 'open';
        path.push(num);
        pathIds.push(id);
        snap({ lineNums: [9, 10, 11, 13], newId: id, start, num, need, left,
          state: `Choose ${num}: need ${need}, ${left} left, so it can work. Append it: path = [${path.join(', ')}].` });

        current = id;
        snap({ lineNums: [14], start: num + 1, num, need: k - path.length, left: n - num,
          state: `Explore: call backtrack(${num + 1}) one level deeper. Only numbers above ${num} may follow.` });

        backtrack(num + 1, id);

        path.pop();
        pathIds.pop();
        current = nodeId;
        if (base[id] !== 'found') base[id] = 'done';
        snap({ lineNums: [15], start, num, need, left,
          state: `Un-choose ${num}: pop it, so path = [${path.join(', ')}] again. ${num < n ? `Back in backtrack(${start}), try ${num + 1}.` : `backtrack(${start}) has no numbers left to try, so it returns.`}` });
      }
    };

    backtrack(1, 0);

    base[0] = 'done';
    snap({ lineNums: [18], start: 1,
      state: `Done: ${result.length} combinations found and ${pruned} branches pruned without exploring them.` });

    this.layout();
  }

  // Leaves (in DFS order) get consecutive columns; a parent sits centred
  // over its first and last child.
  layout() {
    let col = 0;
    const place = (id) => {
      const node = this.nodes[id];
      if (node.children.length === 0) {
        node.col = col++;
      } else {
        node.children.forEach(place);
        const first = this.nodes[node.children[0]];
        const last = this.nodes[node.children[node.children.length - 1]];
        node.col = (first.col + last.col) / 2;
      }
    };
    place(0);
    this.cols = col;
    this.depth = Math.max(...this.nodes.map(nd => nd.depth));
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep];
    if (!step) return;

    const colW = 62, rowH = 72, padX = 34, padTop = 26, padBottom = 40, r = 16;
    const width = padX * 2 + (this.cols - 1) * colW;
    const height = padTop + padBottom + this.depth * rowH;
    const pos = (nd) => ({ x: padX + nd.col * colW, y: padTop + nd.depth * rowH });
    const onPath = new Set(step.pathIds);

    const visible = this.nodes.filter(nd => step.base[nd.id]);
    let edges = '';
    let nodes = '';
    visible.forEach(nd => {
      const status = step.base[nd.id];
      const cls = ['dt-node', `dt-${status}`];
      if (onPath.has(nd.id)) cls.push('dt-on-path');
      if (nd.id === step.current) cls.push('dt-current');
      if (nd.id === step.newId) cls.push('dt-new');
      const p = pos(nd);

      if (nd.parent !== null) {
        const q = pos(this.nodes[nd.parent]);
        const ecls = ['dt-edge', `dt-${status}`];
        if (onPath.has(nd.id)) ecls.push('dt-on-path');
        if (nd.id === step.newId) ecls.push('dt-new');
        edges += `<line class="${ecls.join(' ')}" x1="${q.x}" y1="${q.y + r}" x2="${p.x}" y2="${p.y - r}"/>`;
      }

      const label = nd.parent === null ? '[ ]' : String(nd.num);
      let cross = '';
      let extra = '';
      if (status === 'pruned') {
        const d = r * 0.62;
        cross = `<path class="dt-cross" d="M${p.x - d} ${p.y - d}L${p.x + d} ${p.y + d}M${p.x + d} ${p.y - d}L${p.x - d} ${p.y + d}"/>`;
        extra += `<text class="dt-sub dt-sub-pruned" x="${p.x}" y="${p.y + r + 14}">cut</text>`;
      } else if (status === 'found') {
        extra += `<text class="dt-sub dt-sub-found" x="${p.x}" y="${p.y + r + 14}">[${nd.path.join(',')}]</text>`;
      }
      nodes += `<g class="${cls.join(' ')}"><title>path [${nd.path.join(', ')}]${status === 'pruned' ? ' (pruned)' : ''}</title>` +
        `<circle cx="${p.x}" cy="${p.y}" r="${r}"/>${cross}` +
        `<text class="dt-label" x="${p.x}" y="${p.y}">${label}</text>${extra}</g>`;
    });

    let html = '<div class="dt-viz">';
    html += `<svg class="dt-tree" viewBox="0 0 ${width} ${height}" role="img" aria-label="Decision tree for combinations of ${this.k} from 1..${this.n}">${edges}${nodes}</svg>`;

    html += '<div class="dt-legend">' +
      '<span><i class="dt-key dt-key-path"></i>current path</span>' +
      '<span><i class="dt-key dt-key-done"></i>undone</span>' +
      '<span><i class="dt-key dt-key-found"></i>answer</span>' +
      '<span><i class="dt-key dt-key-pruned"></i>pruned</span></div>';

    html += '<div class="dt-path"><span class="dt-caption">path</span>';
    if (step.path.length === 0) html += '<span class="dt-empty">empty</span>';
    step.path.forEach(v => { html += `<span class="dt-chip">${v}</span>`; });
    html += '</div>';

    html += `<div class="dt-results"><div class="dt-caption">result (${step.result.length})</div><div class="dt-result-list">`;
    if (step.result.length === 0) html += '<span class="dt-empty">none yet</span>';
    step.result.forEach((res, i) => {
      const fresh = i === step.result.length - 1 && step.lineNums.includes(7);
      html += `<span class="dt-result${fresh ? ' dt-new' : ''}">[${res.join(', ')}]</span>`;
    });
    html += '</div></div></div>';

    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    const vars = { start: step.start };
    if (step.num !== undefined) {
      vars.num = step.num;
      vars.need = step.need;
      vars.left = step.left;
    }
    vars.path = `[${(step.path || []).join(', ')}]`;
    vars.results = (step.result || []).length;
    vars.pruned = step.pruned || 0;
    return vars;
  }

  getInputData() {
    let count = 1;
    for (let i = 0; i < this.k; i++) count = count * (this.n - i) / (i + 1);
    return { n: this.n, k: this.k, 'Answers': `C(${this.n}, ${this.k}) = ${count}` };
  }
}

animationRegistry['decision-tree'] = DecisionTreeAnimator;
