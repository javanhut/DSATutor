// Merge Intervals animator for chapter 18 (storyboard "intervals").
// Loaded after app.js, which defines AlgorithmAnimator, animationRegistry
// and escapeHtml. Styles live in anim-intervals.css.

class IntervalsAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    // Covers a contained interval ([4,5]), touching ends ([9,12] then
    // [12,13]) and gaps, so every branch of the loop shows up.
    this.intervals = [[8, 10], [2, 6], [1, 3], [15, 17], [9, 12], [4, 5], [12, 13]];
    this.maxCoord = 18;

    this.setCode(`def merge(intervals):
    merged = []
    for start, end in sorted(intervals, key=lambda i: i[0]):
        if merged and start <= merged[-1][1]:
            merged[-1][1] = max(merged[-1][1], end)
        else:
            merged.append([start, end])
    return merged`);
  }

  buildSteps() {
    this.steps = [];
    const input = this.intervals.map((iv, idx) => ({ start: iv[0], end: iv[1], id: idx }));
    const sorted = [...input].sort((a, b) => a.start - b.start);
    const merged = [];
    const snap = () => merged.map(m => ({ start: m.start, end: m.end }));
    const fmt = (list) => '[' + list.map(m => `[${m.start},${m.end}]`).join(',') + ']';

    this.steps.push({
      lineNum: 1,
      phase: 'input',
      order: input,
      state: `Input: ${input.length} intervals in no particular order: ${fmt(input)}`,
      cur: -1, merged: snap(), last: -1,
    });
    this.steps.push({
      lineNum: 2,
      phase: 'input',
      order: input,
      state: 'merged = [] : the output row starts empty',
      cur: -1, merged: snap(), last: -1,
    });
    this.steps.push({
      lineNum: 3,
      phase: 'sorted',
      order: sorted,
      prevOrder: input,
      state: `Sort by start: ${fmt(sorted)}. Overlapping intervals are now neighbours.`,
      cur: -1, merged: snap(), last: -1,
    });

    sorted.forEach((iv, idx) => {
      const lastIdx = merged.length - 1;
      const last = merged[lastIdx];
      const overlaps = !!last && iv.start <= last.end;
      let compareText;
      if (!last) {
        compareText = `[${iv.start},${iv.end}]: merged is empty, nothing to compare with`;
      } else {
        compareText = `[${iv.start},${iv.end}]: start ${iv.start} <= last end ${last.end}? ${overlaps ? 'yes, overlap' : 'no, gap'}` +
          (iv.start === last.end ? ' (closed intervals: touching counts)' : '');
      }
      this.steps.push({
        lineNum: 4,
        phase: 'sweep',
        order: sorted,
        state: `Compare with the last merged. ${compareText}`,
        cur: idx, done: idx, merged: snap(), last: lastIdx,
        verdict: last ? (overlaps ? 'overlap' : 'gap') : 'gap',
        start: iv.start, end: iv.end,
      });

      if (overlaps) {
        const oldEnd = last.end;
        last.end = Math.max(last.end, iv.end);
        const grew = last.end > oldEnd;
        this.steps.push({
          lineNum: 5,
          phase: 'sweep',
          order: sorted,
          state: grew
            ? `Extend: max(${oldEnd}, ${iv.end}) = ${last.end}, block becomes [${last.start},${last.end}]`
            : `Extend: max(${oldEnd}, ${iv.end}) = ${oldEnd}. [${iv.start},${iv.end}] lies inside [${last.start},${last.end}], so the block keeps its end`,
          cur: idx, done: idx + 1, merged: snap(), last: lastIdx,
          stretch: grew ? { idx: lastIdx, from: oldEnd } : null,
          verdict: 'overlap',
          start: iv.start, end: iv.end,
        });
      } else {
        merged.push({ start: iv.start, end: iv.end });
        this.steps.push({
          lineNum: 7,
          phase: 'sweep',
          order: sorted,
          state: `Append [${iv.start},${iv.end}] as a new block. merged = ${fmt(merged)}`,
          cur: idx, done: idx + 1, merged: snap(), last: merged.length - 1,
          appended: merged.length - 1,
          verdict: 'gap',
          start: iv.start, end: iv.end,
        });
      }
    });

    this.steps.push({
      lineNum: 8,
      phase: 'done',
      order: sorted,
      state: `Done: ${sorted.length} intervals merged into ${merged.length} disjoint blocks ${fmt(merged)}`,
      cur: -1, done: sorted.length, merged: snap(), last: -1, complete: true,
    });
  }

  pct(x) {
    return (x / this.maxCoord) * 100;
  }

  bar(start, end, cls, extraStyle = '') {
    const left = this.pct(start);
    const width = Math.max(this.pct(end) - left, 0.8);
    return `<div class="iv-bar ${cls}" style="left:${left}%;width:${width}%;${extraStyle}">` +
      `<span class="iv-bar-label">${start},${end}</span></div>`;
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep];
    if (!step) return;
    const order = step.order;
    const prevPos = new Map();
    if (step.prevOrder) step.prevOrder.forEach((iv, i) => prevPos.set(iv.id, i));

    let html = '<div class="iv-viz">';

    // Number line ticks, shared by both sections.
    html += '<div class="iv-row iv-axis"><span class="iv-label"></span><div class="iv-track iv-ticks">';
    for (let t = 0; t <= this.maxCoord; t += 2) {
      html += `<span class="iv-tick" style="left:${this.pct(t)}%">${t}</span>`;
    }
    html += '</div></div>';

    const heading = step.phase === 'input' ? 'Input (unsorted)' : 'Sorted by start';
    html += `<div class="iv-section-title">${heading}</div>`;
    html += '<div class="iv-rows">';
    if (step.cur >= 0) {
      // The sweep line sits at the current interval's start.
      html += `<div class="iv-sweep" style="--f:${step.start / this.maxCoord}"><span>start=${step.start}</span></div>`;
    }
    order.forEach((iv, i) => {
      let cls = 'iv-pending';
      if (step.phase === 'sweep' || step.phase === 'done') {
        if (i === step.cur) cls = `iv-current iv-${step.verdict}`;
        else if (i < step.done) cls = 'iv-processed';
      }
      let style = '';
      if (prevPos.has(iv.id) && prevPos.get(iv.id) !== i) {
        style = `--dy:${(prevPos.get(iv.id) - i) * 100}%;`;
        cls += ' iv-moving';
      }
      html += `<div class="iv-row" style="${style}"><span class="iv-label">[${iv.start},${iv.end}]</span>` +
        `<div class="iv-track">${this.bar(iv.start, iv.end, cls)}</div></div>`;
    });
    html += '</div>';

    html += '<div class="iv-section-title">merged</div>';
    html += '<div class="iv-row iv-output"><span class="iv-label">' + step.merged.length + ' block' +
      (step.merged.length === 1 ? '' : 's') + '</span><div class="iv-track">';
    if (step.merged.length === 0) {
      html += '<span class="iv-empty">empty</span>';
    }
    step.merged.forEach((m, i) => {
      let cls = 'iv-block';
      let style = '';
      if (i === step.last) cls += ' iv-last';
      if (step.complete) cls += ' iv-final';
      if (step.stretch && step.stretch.idx === i) {
        cls += ' iv-stretch';
        style = `--from:${Math.max(this.pct(step.stretch.from) - this.pct(m.start), 0.8)}%;`;
      }
      if (step.appended === i) cls += ' iv-appear';
      html += this.bar(m.start, m.end, cls, style);
    });
    html += '</div></div>';

    if (step.last >= 0 && !step.complete) {
      const m = step.merged[step.last];
      html += `<div class="iv-legend"><span class="iv-swatch iv-swatch-last"></span>last merged = [${m.start},${m.end}]` +
        `<span class="iv-swatch iv-swatch-current"></span>current interval</div>`;
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep];
    if (!step) return {};
    const vars = {};
    if (step.cur >= 0) {
      vars.start = step.start;
      vars.end = step.end;
    }
    if (step.merged.length > 0) {
      const m = step.merged[step.merged.length - 1];
      vars['merged[-1]'] = `[${m.start}, ${m.end}]`;
    }
    vars.merged = '[' + step.merged.map(m => `[${m.start},${m.end}]`).join(', ') + ']';
    if (step.verdict && step.lineNum === 4) {
      vars.overlap = step.verdict === 'overlap';
    }
    return vars;
  }

  getInputData() {
    return { intervals: this.intervals };
  }
}

animationRegistry['intervals'] = IntervalsAnimator;
