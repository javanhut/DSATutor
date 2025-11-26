const state = {
  chapters: [],
  current: null,
  timers: [],
  visualizer: null,
  visualizerHandler: null,
  storyboardIndex: 0,
  visualizerIndex: 0,
  editMode: false,
};

// Complexity functions for Big-O curves
const complexityFunctions = {
  'O(1)':       (n) => 1,
  'O(log n)':   (n) => Math.log2(Math.max(n, 1)),
  'O(n)':       (n) => n,
  'O(n log n)': (n) => n * Math.log2(Math.max(n, 1)),
  'O(n^2)':     (n) => n * n,
};

// Curve styling configuration
const curveConfig = {
  'O(1)':       { color: '#22c55e', label: 'O(1)' },
  'O(log n)':   { color: '#38bdf8', label: 'O(log n)' },
  'O(n)':       { color: '#f59e0b', label: 'O(n)' },
  'O(n log n)': { color: '#a855f7', label: 'O(n log n)' },
  'O(n^2)':     { color: '#ef4444', label: 'O(n^2)' },
};

// Step to N-value mapping for storyboard progression
const stepToNMap = { 0: 10, 1: 25, 2: 50, 3: 75, 4: 100 };

// RuntimeShapesVisualizer: SVG-based Big-O curve visualization
class RuntimeShapesVisualizer {
  constructor(container, config) {
    this.container = container;
    this.config = config;
    this.maxN = 100;
    this.currentN = 100;
    this.curveVisibility = {};
    Object.keys(curveConfig).forEach(k => this.curveVisibility[k] = true);
    this.svg = null;
    this.paths = {};
    this.animationId = null;

    // SVG dimensions
    this.width = 600;
    this.height = 400;
    this.padding = { top: 30, right: 30, bottom: 50, left: 60 };
    this.chartWidth = this.width - this.padding.left - this.padding.right;
    this.chartHeight = this.height - this.padding.top - this.padding.bottom;
  }

  mount() {
    this.container.innerHTML = '';

    // Create controls container
    const controls = document.createElement('div');
    controls.className = 'viz-controls';
    this.container.appendChild(controls);

    // Create toggle buttons
    this.createToggles(controls);

    // Create slider
    this.createSlider(controls);

    // Create SVG canvas
    this.createSVG();

    // Initial render
    this.render();
  }

  unmount() {
    if (this.animationId) {
      cancelAnimationFrame(this.animationId);
      this.animationId = null;
    }
    this.container.innerHTML = '';
  }

  createToggles(container) {
    const togglesDiv = document.createElement('div');
    togglesDiv.className = 'curve-toggles';

    Object.entries(curveConfig).forEach(([key, cfg]) => {
      const btn = document.createElement('button');
      btn.className = 'curve-toggle active';
      btn.dataset.curve = key;
      btn.innerHTML = `<span class="toggle-dot" style="background:${cfg.color}"></span>${cfg.label}`;
      btn.onclick = () => this.toggleCurve(key, btn);
      togglesDiv.appendChild(btn);
    });

    container.appendChild(togglesDiv);
  }

  createSlider(container) {
    const sliderWrap = document.createElement('div');
    sliderWrap.className = 'slider-wrap';

    const label = document.createElement('span');
    label.className = 'n-label';
    label.textContent = `n = ${this.currentN}`;
    this.nLabel = label;

    const slider = document.createElement('input');
    slider.type = 'range';
    slider.min = 1;
    slider.max = this.maxN;
    slider.value = this.currentN;
    slider.className = 'n-slider';
    slider.oninput = (e) => {
      this.currentN = parseInt(e.target.value, 10);
      this.nLabel.textContent = `n = ${this.currentN}`;
      this.render();
    };
    this.slider = slider;

    sliderWrap.appendChild(label);
    sliderWrap.appendChild(slider);
    container.appendChild(sliderWrap);
  }

  createSVG() {
    const svgNS = 'http://www.w3.org/2000/svg';

    this.svg = document.createElementNS(svgNS, 'svg');
    this.svg.setAttribute('viewBox', `0 0 ${this.width} ${this.height}`);
    this.svg.setAttribute('class', 'viz-canvas');

    // Create axes group
    const axesGroup = document.createElementNS(svgNS, 'g');
    axesGroup.setAttribute('class', 'axes');
    this.drawAxes(axesGroup);
    this.svg.appendChild(axesGroup);

    // Create curves group
    const curvesGroup = document.createElementNS(svgNS, 'g');
    curvesGroup.setAttribute('class', 'curves');

    // Create path for each curve
    Object.entries(curveConfig).forEach(([key, cfg]) => {
      const path = document.createElementNS(svgNS, 'path');
      const safeClass = key.replace(/[()^]/g, '').replace(/\s/g, '-');
      path.setAttribute('class', `curve curve-${safeClass}`);
      path.setAttribute('stroke', cfg.color);
      path.setAttribute('fill', 'none');
      path.setAttribute('stroke-width', '2.5');
      path.setAttribute('stroke-linecap', 'round');
      path.setAttribute('stroke-linejoin', 'round');
      this.paths[key] = path;
      curvesGroup.appendChild(path);
    });

    this.svg.appendChild(curvesGroup);
    this.container.appendChild(this.svg);
  }

  drawAxes(group) {
    const svgNS = 'http://www.w3.org/2000/svg';
    const { padding, chartWidth, chartHeight } = this;

    // X-axis
    const xAxis = document.createElementNS(svgNS, 'line');
    xAxis.setAttribute('x1', padding.left);
    xAxis.setAttribute('y1', padding.top + chartHeight);
    xAxis.setAttribute('x2', padding.left + chartWidth);
    xAxis.setAttribute('y2', padding.top + chartHeight);
    xAxis.setAttribute('stroke', 'var(--border)');
    group.appendChild(xAxis);

    // Y-axis
    const yAxis = document.createElementNS(svgNS, 'line');
    yAxis.setAttribute('x1', padding.left);
    yAxis.setAttribute('y1', padding.top);
    yAxis.setAttribute('x2', padding.left);
    yAxis.setAttribute('y2', padding.top + chartHeight);
    yAxis.setAttribute('stroke', 'var(--border)');
    group.appendChild(yAxis);

    // X-axis label
    const xLabel = document.createElementNS(svgNS, 'text');
    xLabel.setAttribute('x', padding.left + chartWidth / 2);
    xLabel.setAttribute('y', padding.top + chartHeight + 40);
    xLabel.setAttribute('text-anchor', 'middle');
    xLabel.setAttribute('fill', 'var(--muted)');
    xLabel.setAttribute('font-size', '14');
    xLabel.textContent = 'Input size (n)';
    group.appendChild(xLabel);

    // Y-axis label
    const yLabel = document.createElementNS(svgNS, 'text');
    yLabel.setAttribute('x', -padding.top - chartHeight / 2);
    yLabel.setAttribute('y', 20);
    yLabel.setAttribute('text-anchor', 'middle');
    yLabel.setAttribute('fill', 'var(--muted)');
    yLabel.setAttribute('font-size', '14');
    yLabel.setAttribute('transform', 'rotate(-90)');
    yLabel.textContent = 'Operations';
    group.appendChild(yLabel);

    // X-axis tick marks
    [0, 25, 50, 75, 100].forEach(n => {
      const x = this.scaleX(n);
      const tick = document.createElementNS(svgNS, 'text');
      tick.setAttribute('x', x);
      tick.setAttribute('y', padding.top + chartHeight + 18);
      tick.setAttribute('text-anchor', 'middle');
      tick.setAttribute('fill', 'var(--muted)');
      tick.setAttribute('font-size', '11');
      tick.textContent = n;
      group.appendChild(tick);
    });
  }

  scaleX(n) {
    return this.padding.left + (n / this.maxN) * this.chartWidth;
  }

  scaleY(cost) {
    const maxCost = this.maxN * this.maxN;
    const logCost = Math.log(cost + 1);
    const logMax = Math.log(maxCost + 1);
    return this.padding.top + this.chartHeight - (logCost / logMax) * this.chartHeight;
  }

  generatePath(fn, maxN) {
    const points = [];
    for (let n = 1; n <= maxN; n++) {
      const x = this.scaleX(n);
      const y = this.scaleY(fn(n));
      points.push(`${n === 1 ? 'M' : 'L'} ${x.toFixed(2)} ${y.toFixed(2)}`);
    }
    return points.join(' ');
  }

  render() {
    Object.entries(complexityFunctions).forEach(([key, fn]) => {
      const path = this.paths[key];
      if (path) {
        const d = this.generatePath(fn, this.currentN);
        path.setAttribute('d', d);
        path.style.opacity = this.curveVisibility[key] ? 1 : 0;
      }
    });
  }

  toggleCurve(key, btn) {
    this.curveVisibility[key] = !this.curveVisibility[key];
    btn.classList.toggle('active', this.curveVisibility[key]);
    this.render();
  }

  animateToN(targetN, duration = 500) {
    if (this.animationId) {
      cancelAnimationFrame(this.animationId);
    }

    const startN = this.currentN;
    const startTime = performance.now();

    const frame = (time) => {
      const elapsed = time - startTime;
      const progress = Math.min(elapsed / duration, 1);
      const eased = 1 - Math.pow(1 - progress, 3); // ease-out cubic

      this.currentN = Math.round(startN + (targetN - startN) * eased);
      this.nLabel.textContent = `n = ${this.currentN}`;
      this.slider.value = this.currentN;
      this.render();

      if (progress < 1) {
        this.animationId = requestAnimationFrame(frame);
      } else {
        this.animationId = null;
      }
    };

    this.animationId = requestAnimationFrame(frame);
  }

  onStep(payload) {
    const stepIdx = payload.step;
    const targetN = stepToNMap[stepIdx] || this.maxN;
    this.animateToN(targetN);
  }

  onReset() {
    this.animateToN(1);
  }

  onCurveSelect(payload) {
    const { curve, visible } = payload;
    if (curve && this.curveVisibility.hasOwnProperty(curve)) {
      this.curveVisibility[curve] = visible;
      this.render();
    }
  }
}

// Visualizer registry maps visualizer IDs to handler classes
const visualizerRegistry = {
  'timeline-big-o': RuntimeShapesVisualizer,
};

// Animation registry maps storyboard IDs to animation classes
const animationRegistry = {};

// Base class for step-by-step algorithm animations
class AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    this.canvas = canvasEl;
    this.codeEl = codeEl;
    this.stateEl = stateEl;
    this.config = config;
    this.steps = [];
    this.currentStep = 0;
    this.isPlaying = false;
    this.playTimer = null;
    this.speed = 5;
    this.code = '';
    this.codeLines = [];
    this.varsEl = document.getElementById('vars-display');
    this.inputEl = document.getElementById('input-display');
  }

  setCode(code) {
    this.code = code;
    this.codeLines = code.split('\n');
    this.renderCode();
  }

  renderCode() {
    if (!this.codeEl) return;
    this.codeEl.innerHTML = this.codeLines.map((line, idx) =>
      `<div class="code-line" data-line="${idx + 1}"><span class="line-num">${idx + 1}</span><span class="line-content">${escapeHtml(line)}</span></div>`
    ).join('');
  }

  highlightLine(lineNum) {
    if (!this.codeEl) return;
    this.codeEl.querySelectorAll('.code-line').forEach(el => {
      el.classList.remove('active', 'executed');
    });
    if (lineNum > 0 && lineNum <= this.codeLines.length) {
      const lineEl = this.codeEl.querySelector(`[data-line="${lineNum}"]`);
      if (lineEl) {
        lineEl.classList.add('active');
        lineEl.scrollIntoView({ behavior: 'smooth', block: 'center' });
      }
    }
  }

  highlightLines(lineNums) {
    if (!this.codeEl) return;
    this.codeEl.querySelectorAll('.code-line').forEach(el => {
      el.classList.remove('active');
    });
    lineNums.forEach(num => {
      if (num > 0 && num <= this.codeLines.length) {
        const lineEl = this.codeEl.querySelector(`[data-line="${num}"]`);
        if (lineEl) lineEl.classList.add('active');
      }
    });
  }

  buildSteps() {
    // Override in subclass to generate animation steps
    this.steps = [];
  }

  render() {
    // Override in subclass to render current state
  }

  updateState(text) {
    if (this.stateEl) {
      this.stateEl.textContent = text;
    }
  }

  // Update the variables display panel - override in subclass
  updateVariables(vars) {
    if (!this.varsEl) return;
    if (!vars || Object.keys(vars).length === 0) {
      this.varsEl.innerHTML = '<div class="vars-empty">No variables tracked</div>';
      return;
    }
    let html = '<div class="vars-list">';
    Object.entries(vars).forEach(([name, value]) => {
      const displayValue = typeof value === 'object' ? JSON.stringify(value) : value;
      const valueClass = typeof value === 'number' ? 'var-number' :
                         typeof value === 'boolean' ? 'var-boolean' :
                         Array.isArray(value) ? 'var-array' : 'var-value';
      html += `<div class="var-item">
        <span class="var-name">${escapeHtml(name)}</span>
        <span class="${valueClass}">${escapeHtml(String(displayValue))}</span>
      </div>`;
    });
    html += '</div>';
    this.varsEl.innerHTML = html;
  }

  // Update the input data display - override in subclass
  updateInputDisplay(inputData) {
    if (!this.inputEl) return;
    if (!inputData) {
      this.inputEl.innerHTML = '';
      return;
    }
    let html = '<div class="input-section">';
    html += '<div class="input-header">Input Data</div>';
    if (Array.isArray(inputData)) {
      html += '<div class="input-array">';
      inputData.forEach((val, idx) => {
        html += `<span class="input-item" data-idx="${idx}">${val}</span>`;
      });
      html += '</div>';
    } else if (typeof inputData === 'object') {
      html += '<div class="input-object">';
      Object.entries(inputData).forEach(([key, val]) => {
        html += `<div class="input-pair"><span class="input-key">${key}:</span> <span class="input-val">${JSON.stringify(val)}</span></div>`;
      });
      html += '</div>';
    } else {
      html += `<div class="input-value">${inputData}</div>`;
    }
    html += '</div>';
    this.inputEl.innerHTML = html;
  }

  // Get current variables for display - override in subclass
  getVariables() {
    return {};
  }

  // Get input data for display - override in subclass
  getInputData() {
    return null;
  }

  goToStep(stepIdx) {
    if (stepIdx < 0) stepIdx = 0;
    if (stepIdx >= this.steps.length) stepIdx = this.steps.length - 1;
    this.currentStep = stepIdx;

    const step = this.steps[stepIdx];
    if (step) {
      if (step.lineNum) this.highlightLine(step.lineNum);
      if (step.lineNums) this.highlightLines(step.lineNums);
      if (step.state) this.updateState(step.state);
      if (step.apply) step.apply();
    }
    this.render();
    this.updateStepCounter();
    // Update variables and input display
    this.updateVariables(this.getVariables());
    this.updateInputDisplay(this.getInputData());
  }

  stepForward() {
    if (this.currentStep < this.steps.length - 1) {
      this.goToStep(this.currentStep + 1);
    } else {
      this.pause();
    }
  }

  stepBack() {
    if (this.currentStep > 0) {
      this.goToStep(this.currentStep - 1);
    }
  }

  play() {
    if (this.isPlaying) return;
    this.isPlaying = true;
    this.scheduleNext();
  }

  pause() {
    this.isPlaying = false;
    if (this.playTimer) {
      clearTimeout(this.playTimer);
      this.playTimer = null;
    }
  }

  togglePlay() {
    if (this.isPlaying) {
      this.pause();
    } else {
      this.play();
    }
    return this.isPlaying;
  }

  scheduleNext() {
    if (!this.isPlaying) return;
    const delay = 1100 - (this.speed * 100);
    this.playTimer = setTimeout(() => {
      this.stepForward();
      if (this.isPlaying && this.currentStep < this.steps.length - 1) {
        this.scheduleNext();
      } else {
        this.pause();
      }
    }, delay);
  }

  reset() {
    this.pause();
    this.buildSteps();
    this.currentStep = 0;
    this.goToStep(0);
  }

  setSpeed(speed) {
    this.speed = speed;
  }

  updateStepCounter() {
    const counter = el('anim-step-counter');
    if (counter) {
      counter.textContent = `Step ${this.currentStep + 1}/${this.steps.length}`;
    }
  }

  mount() {
    this.buildSteps();
    this.render();
    this.updateStepCounter();
    if (this.steps.length > 0) {
      this.goToStep(0);
    }
  }

  unmount() {
    this.pause();
    if (this.canvas) this.canvas.innerHTML = '';
    if (this.codeEl) this.codeEl.innerHTML = '';
    if (this.stateEl) this.stateEl.textContent = '';
  }
}

// Binary Search Visualizer
class BinarySearchAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.array = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25];
    this.target = 15;
    this.low = 0;
    this.high = this.array.length - 1;
    this.mid = -1;
    this.found = false;
    this.done = false;
    this.guess = null;

    this.setCode(`def binary_search(nums, target):
    low, high = 0, len(nums) - 1
    while low <= high:
        mid = (low + high) // 2
        guess = nums[mid]
        if guess == target:
            return mid
        if guess < target:
            low = mid + 1
        else:
            high = mid - 1
    return None`);
  }

  buildSteps() {
    this.steps = [];
    let low = 0;
    let high = this.array.length - 1;

    // Initial state
    this.steps.push({
      lineNum: 2,
      state: `Initialize: low=0, high=${high}, target=${this.target}`,
      low, high, mid: -1, found: false, done: false,
      apply: () => { this.low = low; this.high = high; this.mid = -1; this.found = false; this.done = false; }
    });

    while (low <= high) {
      const mid = Math.floor((low + high) / 2);
      const guess = this.array[mid];

      // Check condition
      this.steps.push({
        lineNum: 3,
        state: `Check: low(${low}) <= high(${high})? Yes, continue`,
        low, high, mid: -1, found: false, done: false,
        apply: () => { this.low = low; this.high = high; this.mid = -1; }
      });

      // Calculate mid
      this.steps.push({
        lineNum: 4,
        state: `Calculate mid = (${low} + ${high}) / 2 = ${mid}`,
        low, high, mid, found: false, done: false,
        apply: () => { this.mid = mid; }
      });

      // Get guess
      this.steps.push({
        lineNum: 5,
        state: `Get guess = nums[${mid}] = ${guess}`,
        low, high, mid, found: false, done: false,
        apply: () => { this.mid = mid; }
      });

      if (guess === this.target) {
        // Found
        this.steps.push({
          lineNums: [6, 7],
          state: `Found! ${guess} == ${this.target}, return index ${mid}`,
          low, high, mid, found: true, done: true,
          apply: () => { this.found = true; this.done = true; }
        });
        break;
      } else if (guess < this.target) {
        // Too low
        this.steps.push({
          lineNums: [8, 9],
          state: `${guess} < ${this.target}, search right half: low = ${mid + 1}`,
          low, high, mid, found: false, done: false,
          apply: () => {}
        });
        low = mid + 1;
      } else {
        // Too high
        this.steps.push({
          lineNums: [10, 11],
          state: `${guess} > ${this.target}, search left half: high = ${mid - 1}`,
          low, high, mid, found: false, done: false,
          apply: () => {}
        });
        high = mid - 1;
      }
    }

    if (!this.steps[this.steps.length - 1]?.found) {
      this.steps.push({
        lineNum: 12,
        state: `Not found! low(${low}) > high(${high}), return None`,
        low, high, mid: -1, found: false, done: true,
        apply: () => { this.done = true; }
      });
    }
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { low = this.low, high = this.high, mid = this.mid, found = this.found, done = this.done } = step;

    let html = '<div class="array-viz">';
    this.array.forEach((val, idx) => {
      let classes = ['array-cell'];
      if (idx < low || idx > high) classes.push('discarded');
      if (idx === mid) classes.push(found ? 'found' : 'current');
      if (idx === low && !done) classes.push('low-ptr');
      if (idx === high && !done) classes.push('high-ptr');

      // Build label for this cell
      let label = '';
      if (!done) {
        if (idx === mid && mid >= 0) label = 'mid';
        else if (idx === low && idx === high) label = 'L/H';
        else if (idx === low) label = 'low';
        else if (idx === high) label = 'high';
      }

      html += `<div class="${classes.join(' ')}">
        <div class="cell-value">${val}</div>
        <div class="cell-index">${idx}</div>
        ${label ? `<div class="cell-label ${idx === mid ? 'mid' : ''}">${label}</div>` : ''}
      </div>`;
    });
    html += '</div>';

    // Target display
    html += `<div class="target-display">Target: <strong>${this.target}</strong></div>`;

    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'low': step.low !== undefined ? step.low : this.low,
      'high': step.high !== undefined ? step.high : this.high,
      'mid': step.mid !== undefined && step.mid >= 0 ? step.mid : '-',
      'guess': step.mid >= 0 ? this.array[step.mid] : '-',
      'target': this.target,
      'found': step.found ? 'Yes' : 'No'
    };
  }

  getInputData() {
    return this.array;
  }
}

// Selection Sort Visualizer
class SelectionSortAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.originalArray = [64, 25, 12, 22, 11, 34, 90];
    this.array = [...this.originalArray];
    this.sortedBoundary = 0;
    this.scanIndex = -1;
    this.minIndex = -1;
    this.comparing = -1;

    this.setCode(`def selection_sort(arr):
    n = len(arr)
    for i in range(n):
        min_idx = i
        for j in range(i + 1, n):
            if arr[j] < arr[min_idx]:
                min_idx = j
        arr[i], arr[min_idx] = arr[min_idx], arr[i]
    return arr`);
  }

  buildSteps() {
    this.steps = [];
    const arr = [...this.originalArray];
    const n = arr.length;

    this.steps.push({
      lineNum: 2,
      state: `Array length n = ${n}`,
      array: [...arr], sortedBoundary: 0, scanIndex: -1, minIndex: -1,
      apply: () => { this.array = [...arr]; this.sortedBoundary = 0; }
    });

    for (let i = 0; i < n; i++) {
      let minIdx = i;

      this.steps.push({
        lineNum: 3,
        state: `Outer loop: i = ${i}, finding minimum in unsorted portion`,
        array: [...arr], sortedBoundary: i, scanIndex: -1, minIndex: i,
        apply: () => { this.sortedBoundary = i; this.minIndex = i; }
      });

      this.steps.push({
        lineNum: 4,
        state: `Set min_idx = ${i} (value: ${arr[i]})`,
        array: [...arr], sortedBoundary: i, scanIndex: -1, minIndex: i,
        apply: () => { this.minIndex = i; }
      });

      for (let j = i + 1; j < n; j++) {
        this.steps.push({
          lineNum: 5,
          state: `Scanning: j = ${j}`,
          array: [...arr], sortedBoundary: i, scanIndex: j, minIndex: minIdx, comparing: j,
          apply: () => { this.scanIndex = j; this.comparing = j; }
        });

        this.steps.push({
          lineNum: 6,
          state: `Compare arr[${j}]=${arr[j]} < arr[${minIdx}]=${arr[minIdx]}? ${arr[j] < arr[minIdx] ? 'Yes' : 'No'}`,
          array: [...arr], sortedBoundary: i, scanIndex: j, minIndex: minIdx, comparing: j,
          apply: () => {}
        });

        if (arr[j] < arr[minIdx]) {
          minIdx = j;
          this.steps.push({
            lineNum: 7,
            state: `New minimum found! min_idx = ${j} (value: ${arr[j]})`,
            array: [...arr], sortedBoundary: i, scanIndex: j, minIndex: minIdx,
            apply: () => { this.minIndex = minIdx; }
          });
        }
      }

      // Swap
      if (minIdx !== i) {
        this.steps.push({
          lineNum: 8,
          state: `Swap arr[${i}]=${arr[i]} with arr[${minIdx}]=${arr[minIdx]}`,
          array: [...arr], sortedBoundary: i, scanIndex: -1, minIndex: minIdx, swapping: [i, minIdx],
          apply: () => {}
        });

        [arr[i], arr[minIdx]] = [arr[minIdx], arr[i]];

        this.steps.push({
          lineNum: 8,
          state: `After swap: ${arr.join(', ')}`,
          array: [...arr], sortedBoundary: i + 1, scanIndex: -1, minIndex: -1,
          apply: () => { this.array = [...arr]; this.sortedBoundary = i + 1; }
        });
      } else {
        this.steps.push({
          lineNum: 8,
          state: `No swap needed, ${arr[i]} already in correct position`,
          array: [...arr], sortedBoundary: i + 1, scanIndex: -1, minIndex: -1,
          apply: () => { this.sortedBoundary = i + 1; }
        });
      }
    }

    this.steps.push({
      lineNum: 9,
      state: `Sorted! Final array: ${arr.join(', ')}`,
      array: [...arr], sortedBoundary: n, scanIndex: -1, minIndex: -1, done: true,
      apply: () => { this.array = [...arr]; this.sortedBoundary = n; }
    });
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const arr = step.array || this.array;
    const { sortedBoundary = 0, scanIndex = -1, minIndex = -1, swapping = [], done = false } = step;

    let html = '<div class="array-viz sort-viz">';
    arr.forEach((val, idx) => {
      let classes = ['array-cell'];
      if (idx < sortedBoundary) classes.push('sorted');
      if (idx === minIndex) classes.push('minimum');
      if (idx === scanIndex) classes.push('scanning');
      if (swapping.includes(idx)) classes.push('swapping');
      if (done) classes.push('sorted');

      const height = Math.max(20, val * 2);
      html += `<div class="${classes.join(' ')}" style="height: ${height}px">
        <div class="cell-value">${val}</div>
      </div>`;
    });
    html += '</div>';

    // Legend
    html += `<div class="sort-legend">
      <span class="legend-item"><span class="legend-color sorted"></span>Sorted</span>
      <span class="legend-item"><span class="legend-color minimum"></span>Current Min</span>
      <span class="legend-item"><span class="legend-color scanning"></span>Scanning</span>
    </div>`;

    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    const arr = step.array || this.array;
    return {
      'i (outer)': step.sortedBoundary !== undefined ? step.sortedBoundary : '-',
      'j (scan)': step.scanIndex >= 0 ? step.scanIndex : '-',
      'min_idx': step.minIndex >= 0 ? step.minIndex : '-',
      'arr[min_idx]': step.minIndex >= 0 ? arr[step.minIndex] : '-',
      'sorted count': step.sortedBoundary || 0,
      'n': arr.length
    };
  }

  getInputData() {
    const step = this.steps[this.currentStep] || {};
    return step.array || this.originalArray;
  }
}

// BFS Graph Visualizer
class BFSAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.graph = {
      'A': ['B', 'C'],
      'B': ['A', 'D', 'E'],
      'C': ['A', 'F'],
      'D': ['B'],
      'E': ['B', 'F'],
      'F': ['C', 'E']
    };
    this.nodePositions = {
      'A': { x: 150, y: 50 },
      'B': { x: 80, y: 120 },
      'C': { x: 220, y: 120 },
      'D': { x: 40, y: 200 },
      'E': { x: 120, y: 200 },
      'F': { x: 200, y: 200 }
    };
    this.startNode = 'A';
    this.visited = new Set();
    this.queue = [];
    this.currentNode = null;
    this.visitOrder = [];

    this.setCode(`from collections import deque

def bfs(start, graph):
    visited = {start}
    queue = deque([start])
    order = []
    while queue:
        node = queue.popleft()
        order.append(node)
        for nbr in graph.get(node, []):
            if nbr not in visited:
                visited.add(nbr)
                queue.append(nbr)
    return order`);
  }

  buildSteps() {
    this.steps = [];
    const visited = new Set();
    const queue = [];
    const order = [];

    // Initialize
    visited.add(this.startNode);
    queue.push(this.startNode);

    this.steps.push({
      lineNums: [4, 5],
      state: `Initialize: visited={${this.startNode}}, queue=[${this.startNode}]`,
      visited: new Set(visited), queue: [...queue], currentNode: null, order: [...order],
      apply: () => { this.visited = new Set(visited); this.queue = [...queue]; this.currentNode = null; this.visitOrder = []; }
    });

    while (queue.length > 0) {
      const node = queue.shift();
      order.push(node);

      this.steps.push({
        lineNum: 7,
        state: `Check: queue not empty? Yes, length=${queue.length + 1}`,
        visited: new Set(visited), queue: [node, ...queue], currentNode: null, order: [...order].slice(0, -1),
        apply: () => {}
      });

      this.steps.push({
        lineNum: 8,
        state: `Dequeue: node = '${node}'`,
        visited: new Set(visited), queue: [...queue], currentNode: node, order: [...order].slice(0, -1),
        apply: () => { this.queue = [...queue]; this.currentNode = node; }
      });

      this.steps.push({
        lineNum: 9,
        state: `Add '${node}' to visit order: [${order.join(', ')}]`,
        visited: new Set(visited), queue: [...queue], currentNode: node, order: [...order],
        apply: () => { this.visitOrder = [...order]; }
      });

      const neighbors = this.graph[node] || [];
      for (const nbr of neighbors) {
        this.steps.push({
          lineNum: 10,
          state: `Check neighbor '${nbr}' of '${node}'`,
          visited: new Set(visited), queue: [...queue], currentNode: node, order: [...order], checkingNeighbor: nbr,
          apply: () => {}
        });

        if (!visited.has(nbr)) {
          visited.add(nbr);
          queue.push(nbr);

          this.steps.push({
            lineNums: [11, 12, 13],
            state: `'${nbr}' not visited. Add to visited and queue. Queue: [${queue.join(', ')}]`,
            visited: new Set(visited), queue: [...queue], currentNode: node, order: [...order], newlyAdded: nbr,
            apply: () => { this.visited = new Set(visited); this.queue = [...queue]; }
          });
        } else {
          this.steps.push({
            lineNum: 11,
            state: `'${nbr}' already visited, skip`,
            visited: new Set(visited), queue: [...queue], currentNode: node, order: [...order],
            apply: () => {}
          });
        }
      }
    }

    this.steps.push({
      lineNum: 14,
      state: `BFS complete! Visit order: [${order.join(', ')}]`,
      visited: new Set(visited), queue: [], currentNode: null, order: [...order], done: true,
      apply: () => { this.queue = []; this.currentNode = null; }
    });
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { visited = this.visited, queue = this.queue, currentNode = this.currentNode, order = this.visitOrder, checkingNeighbor, newlyAdded, done } = step;

    const svgNS = 'http://www.w3.org/2000/svg';
    let html = '<svg class="graph-viz" viewBox="0 0 300 250">';

    // Draw edges
    const drawnEdges = new Set();
    Object.entries(this.graph).forEach(([node, neighbors]) => {
      neighbors.forEach(nbr => {
        const edgeKey = [node, nbr].sort().join('-');
        if (!drawnEdges.has(edgeKey)) {
          drawnEdges.add(edgeKey);
          const p1 = this.nodePositions[node];
          const p2 = this.nodePositions[nbr];
          html += `<line x1="${p1.x}" y1="${p1.y}" x2="${p2.x}" y2="${p2.y}" class="edge"/>`;
        }
      });
    });

    // Draw nodes
    Object.entries(this.nodePositions).forEach(([node, pos]) => {
      let nodeClass = 'graph-node';
      if (visited.has(node)) nodeClass += ' visited';
      if (node === currentNode) nodeClass += ' current';
      if (queue.includes(node)) nodeClass += ' in-queue';
      if (node === checkingNeighbor) nodeClass += ' checking';
      if (node === newlyAdded) nodeClass += ' newly-added';

      html += `<circle cx="${pos.x}" cy="${pos.y}" r="20" class="${nodeClass}"/>`;
      html += `<text x="${pos.x}" y="${pos.y + 5}" class="node-label">${node}</text>`;
    });

    html += '</svg>';

    // Queue display
    html += '<div class="queue-display">';
    html += '<span class="queue-label">Queue:</span>';
    html += '<div class="queue-items">';
    if (queue.length === 0) {
      html += '<span class="queue-empty">empty</span>';
    } else {
      queue.forEach((item, idx) => {
        html += `<span class="queue-item ${idx === 0 ? 'front' : ''}">${item}</span>`;
      });
    }
    html += '</div></div>';

    // Visit order
    html += '<div class="visit-order">';
    html += `<span class="order-label">Visit Order:</span> [${order.join(' -> ')}]`;
    html += '</div>';

    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    const visited = step.visited || this.visited;
    const queue = step.queue || this.queue;
    const order = step.order || this.visitOrder;
    return {
      'current node': step.currentNode || '-',
      'visited': Array.from(visited).join(', ') || '-',
      'queue': queue.length > 0 ? queue.join(', ') : 'empty',
      'visit order': order.length > 0 ? order.join(' -> ') : '-'
    };
  }

  getInputData() {
    return {
      'Start Node': this.startNode,
      'Graph': this.graph
    };
  }
}

// Call Stack Visualizer
class CallStackAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.inputN = 4;
    this.stack = [];
    this.result = null;
    this.phase = 'push'; // 'push' or 'pop'

    this.setCode(`def factorial(n):
    if n <= 1:        # base case
        return 1
    return n * factorial(n - 1)  # recursive

# Call: factorial(4)`);
  }

  buildSteps() {
    this.steps = [];
    const n = this.inputN;

    // Build push phase
    for (let i = n; i >= 1; i--) {
      this.steps.push({
        lineNum: 1,
        state: `Call factorial(${i})`,
        stack: this.buildStackState(n, i, 'push'),
        phase: 'push',
        currentCall: i,
        apply: () => { this.stack = this.buildStackState(n, i, 'push'); }
      });

      if (i === 1) {
        this.steps.push({
          lineNums: [2, 3],
          state: `Base case: n=${i} <= 1, return 1`,
          stack: this.buildStackState(n, i, 'base'),
          phase: 'base',
          currentCall: i,
          returnValue: 1,
          apply: () => {}
        });
      } else {
        this.steps.push({
          lineNum: 2,
          state: `Check: ${i} <= 1? No`,
          stack: this.buildStackState(n, i, 'push'),
          phase: 'push',
          currentCall: i,
          apply: () => {}
        });

        this.steps.push({
          lineNum: 4,
          state: `Need factorial(${i - 1}) to compute ${i} * factorial(${i - 1})`,
          stack: this.buildStackState(n, i, 'push'),
          phase: 'push',
          currentCall: i,
          waiting: true,
          apply: () => {}
        });
      }
    }

    // Build pop phase (unwinding)
    let result = 1;
    for (let i = 1; i <= n; i++) {
      const prevResult = result;
      result = i * result;

      this.steps.push({
        lineNum: i === 1 ? 3 : 4,
        state: i === 1
          ? `Return 1 from factorial(1)`
          : `Return ${i} * ${prevResult} = ${result} from factorial(${i})`,
        stack: this.buildStackState(n, i, 'pop'),
        phase: 'pop',
        currentCall: i,
        returnValue: result,
        apply: () => { this.result = result; }
      });
    }

    this.steps.push({
      lineNum: 6,
      state: `Final result: factorial(${n}) = ${result}`,
      stack: [],
      phase: 'done',
      result: result,
      apply: () => { this.stack = []; this.result = result; }
    });
  }

  buildStackState(n, current, phase) {
    const stack = [];
    if (phase === 'push' || phase === 'base') {
      for (let i = n; i >= current; i--) {
        stack.push({ n: i, status: i === current ? 'active' : 'waiting' });
      }
    } else if (phase === 'pop') {
      for (let i = n; i > current; i--) {
        stack.push({ n: i, status: 'waiting' });
      }
    }
    return stack;
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { stack = [], phase, currentCall, returnValue, waiting, result } = step;

    let html = '<div class="stack-viz">';

    // Stack frames
    html += '<div class="stack-container">';
    html += '<div class="stack-label">Call Stack</div>';
    html += '<div class="stack-frames">';

    if (stack.length === 0) {
      html += '<div class="stack-empty">Stack empty</div>';
    } else {
      stack.forEach((frame, idx) => {
        let frameClass = 'stack-frame';
        if (frame.status === 'active') frameClass += ' active';
        if (frame.status === 'returning') frameClass += ' returning';

        html += `<div class="${frameClass}">
          <span class="frame-name">factorial(${frame.n})</span>
          ${frame.status === 'active' && returnValue !== undefined ? `<span class="frame-return">= ${returnValue}</span>` : ''}
        </div>`;
      });
    }

    html += '</div></div>';

    // Phase indicator
    html += '<div class="phase-indicator">';
    if (phase === 'push') {
      html += `<div class="phase push">Pushing frames (recursing down)</div>`;
    } else if (phase === 'base') {
      html += `<div class="phase base">Base case reached!</div>`;
    } else if (phase === 'pop') {
      html += `<div class="phase pop">Popping frames (returning up)</div>`;
    } else if (phase === 'done') {
      html += `<div class="phase done">Complete! Result = ${result}</div>`;
    }
    html += '</div>';

    // Visual representation of computation
    if (phase === 'done') {
      html += '<div class="computation-trace">';
      html += `<div>4! = 4 x 3 x 2 x 1 = ${result}</div>`;
      html += '</div>';
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    const stack = step.stack || this.stack;
    return {
      'n (input)': this.inputN,
      'current call': step.currentCall || '-',
      'phase': step.phase || '-',
      'stack depth': stack.length,
      'return value': step.returnValue !== undefined ? step.returnValue : '-',
      'final result': step.result !== undefined ? step.result : '-'
    };
  }

  getInputData() {
    return {
      'Function': 'factorial(n)',
      'n': this.inputN,
      'Expected': `${this.inputN}! = ${this.factorial(this.inputN)}`
    };
  }

  factorial(n) {
    if (n <= 1) return 1;
    return n * this.factorial(n - 1);
  }
}

// Register animations
animationRegistry['binary-search'] = BinarySearchAnimator;
animationRegistry['selection-sort'] = SelectionSortAnimator;
animationRegistry['bfs'] = BFSAnimator;
animationRegistry['call-stack'] = CallStackAnimator;

// Current animation instance
let currentAnimator = null;

// Setup animation controls
function setupAnimationControls() {
  const stepBtn = el('anim-step');
  const stepBackBtn = el('anim-step-back');
  const playBtn = el('anim-play');
  const resetBtn = el('anim-reset');
  const speedSlider = el('anim-speed');

  if (stepBtn) {
    stepBtn.onclick = () => {
      if (currentAnimator) currentAnimator.stepForward();
    };
  }

  if (stepBackBtn) {
    stepBackBtn.onclick = () => {
      if (currentAnimator) currentAnimator.stepBack();
    };
  }

  if (playBtn) {
    playBtn.onclick = () => {
      if (currentAnimator) {
        const playing = currentAnimator.togglePlay();
        playBtn.textContent = playing ? 'Pause' : 'Play';
      }
    };
  }

  if (resetBtn) {
    resetBtn.onclick = () => {
      if (currentAnimator) {
        currentAnimator.reset();
        if (playBtn) playBtn.textContent = 'Play';
      }
    };
  }

  if (speedSlider) {
    speedSlider.oninput = (e) => {
      if (currentAnimator) {
        currentAnimator.setSpeed(parseInt(e.target.value, 10));
      }
    };
  }
}

// Mount animation for a storyboard
function mountAnimation(storyboardId) {
  // Cleanup previous
  if (currentAnimator) {
    currentAnimator.unmount();
    currentAnimator = null;
  }

  const AnimatorClass = animationRegistry[storyboardId];
  if (!AnimatorClass) {
    return false;
  }

  const canvas = document.getElementById('visualizer-canvas');
  const codeEl = document.getElementById('code-lines');
  const stateEl = document.getElementById('visualizer-state');

  currentAnimator = new AnimatorClass(canvas, codeEl, stateEl, {});
  currentAnimator.mount();

  const playBtn = document.getElementById('anim-play');
  if (playBtn) playBtn.textContent = 'Play';

  return true;
}

const el = (id) => document.getElementById(id);

// Helper to escape HTML
function escapeHtml(str) {
  return str.replace(/[&<>"']/g, (c) => ({
    "&": "&amp;",
    "<": "&lt;",
    ">": "&gt;",
    '"': "&quot;",
    "'": "&#39;",
  }[c]));
}

async function loadChapters() {
  const res = await fetch("/api/chapters");
  const data = await res.json();
  state.chapters = data;
  renderChapters();
  if (state.current) {
    const updated = state.chapters.find((c) => c.number === state.current.number);
    if (updated) selectChapter(updated);
  }
}

function renderChapters() {
  const container = el("chapters");
  container.innerHTML = "";
  state.chapters.forEach((ch) => {
    const item = document.createElement("div");
    item.className = "chapter-item" + (state.current && state.current.number === ch.number ? " active" : "");
    item.innerHTML = `
      <span class="chapter-num">${pad(ch.number)}</span>
      <div class="chapter-info">
        <div class="chapter-name">${ch.title}</div>
        <div class="chapter-slug">${ch.slug}</div>
      </div>
    `;
    item.onclick = () => selectChapter(ch);
    container.appendChild(item);
  });
}

function selectChapter(ch) {
  state.current = ch;
  state.storyboardIndex = 0;
  state.visualizerIndex = 0;
  el("detail-empty").classList.add("hidden");
  el("detail").classList.remove("hidden");
  el("detail-slug").textContent = ch.slug;
  el("detail-title").textContent = `${pad(ch.number)}. ${ch.title}`;
  el("detail-summary").textContent = ch.summary;

  const obj = el("detail-objectives");
  obj.innerHTML = ch.objectives.map((o) => `<li>${o}</li>`).join("");

  const vis = el("detail-visualizers");
  if (ch.visualizers && ch.visualizers.length) {
    vis.innerHTML = ch.visualizers.map((v) => `<li><strong>${v.title}</strong> — ${v.goal} (hooks: ${v.hooks?.join(", ") || "none"})</li>`).join("");
  } else {
    vis.innerHTML = `<li>No visualizers defined yet.</li>`;
  }

  populateSelectors();
  renderStoryboard(getStoryboard());
  renderExamples(ch);
  mountVisualizer(getVisualizer());
  renderEditor();
}

function renderStoryboard(sb) {
  const stepsEl = el("storyboard-steps");
  stepsEl.innerHTML = "";
  if (!sb || !sb.steps) {
    stepsEl.innerHTML = `<p class="placeholder">No storyboard defined.</p>`;
    return;
  }

  sb.steps.forEach((step, idx) => {
    const node = document.createElement("div");
    node.className = "step-item";
    node.dataset.index = idx;
    node.innerHTML = `
      <span class="step-num">${idx + 1}</span>
      <div class="step-content">
        <div class="step-cue">${step.cue}</div>
        <div class="step-narration">${step.narration}</div>
        <div class="step-meta">
          <span>${step.visualHint || ""}</span>
          ${step.codeRef ? `<span class="step-code-ref">${step.codeRef}</span>` : ""}
          <span class="step-duration">${step.duration}</span>
        </div>
      </div>
    `;
    stepsEl.appendChild(node);
  });
}

function renderExamples(ch) {
  const target = el("detail-examples");
  target.innerHTML = "";
  const examples = (ch.concepts || []).flatMap((c) => c.examples || []);
  if (ch.examples) examples.push(...ch.examples);
  if (!examples.length) {
    target.innerHTML = `<p class="placeholder">No code examples yet.</p>`;
    return;
  }
  examples.forEach((ex) => {
    const block = document.createElement("div");
    block.className = "example";
    block.dataset.id = ex.id || "";
    block.innerHTML = `
      <div class="meta">${ex.language || "code"}</div>
      <strong>${ex.title}</strong>
      <pre>${escapeHtml(ex.snippet || "")}</pre>
      ${ex.notes ? `<p>${ex.notes}</p>` : ""}
    `;
    target.appendChild(block);
  });
}

function mountVisualizer(v) {
  // Unmount previous handler if exists
  if (state.visualizerHandler) {
    state.visualizerHandler.unmount();
    state.visualizerHandler = null;
  }

  state.visualizer = v || null;
  const meta = el("visualizer-meta");
  const log = el("visualizer-log");
  const panel = el("visualizer-panel");
  log.innerHTML = "";

  if (!v) {
    meta.textContent = "No visualizer attached to this chapter.";
    panel.innerHTML = "";
    return;
  }

  meta.textContent = `${v.title} - hooks: ${v.hooks?.join(", ") || "none"}`;
  logMessage("Visualizer mounted.");

  // Check if we have a registered handler for this visualizer
  const HandlerClass = visualizerRegistry[v.id];
  if (HandlerClass) {
    state.visualizerHandler = new HandlerClass(panel, v);
    state.visualizerHandler.mount();
  } else {
    // Fallback to legacy stage rendering
    drawVisualizerStage(v);
  }
}

function logMessage(msg) {
  const log = el("visualizer-log");
  const entry = document.createElement("div");
  const now = new Date().toLocaleTimeString();
  entry.textContent = `[${now}] ${msg}`;
  log.prepend(entry);
}

function playStoryboard() {
  stopStoryboard();
  const ch = state.current;
  if (!ch || !ch.animations || !ch.animations.length) return;
  const sb = getStoryboard();
  const steps = sb.steps || [];
  let offset = 0;
  steps.forEach((step, idx) => {
    const delay = parseDuration(step.duration);
    const start = setTimeout(() => setActiveStep(idx), offset);
    state.timers.push(start);
    offset += delay;
  });

  // Fire onStep hooks if present.
  let hookOffset = 0;
  steps.forEach((step, idx) => {
    const delay = parseDuration(step.duration);
    const timer = setTimeout(() => emitVisualizerHook("onStep", { step: idx, cue: step.cue }), hookOffset);
    state.timers.push(timer);
    offset += delay;
  });
}

function stopStoryboard() {
  state.timers.forEach((t) => clearTimeout(t));
  state.timers = [];
  document.querySelectorAll(".step-item").forEach((el) => el.classList.remove("active"));
  highlightCodeRef(null);
  emitVisualizerHook("onReset", {});
  drawVisualizerStage(state.visualizer, { activeCue: null, step: null });
}

function setActiveStep(idx) {
  document.querySelectorAll(".step-item").forEach((el) => {
    if (Number(el.dataset.index) === idx) {
      el.classList.add("active");
    } else {
      el.classList.remove("active");
    }
  });
  const sb = getStoryboard();
  const step = sb?.steps?.[idx];
  if (step && step.codeRef) {
    highlightCodeRef(step.codeRef);
    drawVisualizerStage(state.visualizer, { activeCue: step.cue, step: idx, codeRef: step.codeRef });
  } else {
    highlightCodeRef(null);
    drawVisualizerStage(state.visualizer, { activeCue: step?.cue, step: idx });
  }
}

function emitVisualizerHook(name, payload) {
  const v = state.visualizer;
  if (!v || !v.hooks || !v.hooks.includes(name)) return;
  logMessage(`Hook ${name} fired: ${JSON.stringify(payload)}`);

  // Dispatch to handler if available
  const handler = state.visualizerHandler;
  if (handler && typeof handler[name] === 'function') {
    handler[name](payload);
  } else {
    // Fallback for visualizers without handlers
    drawVisualizerStage(v, { activeCue: payload.cue, step: payload.step, hook: name });
  }
}

function highlightCodeRef(ref) {
  const blocks = document.querySelectorAll(".example");
  blocks.forEach((b) => b.classList.remove("active"));
  if (!ref) return;
  const prefix = ref.split(":")[0];
  blocks.forEach((b) => {
    if (b.dataset.id === prefix) {
      b.classList.add("active");
    }
  });
}

function drawVisualizerStage(v, stateInfo = {}) {
  const panel = document.querySelector("#visualizer-panel");
  if (!panel) return;
  let stage = panel.querySelector(".stage");
  if (!stage) {
    stage = document.createElement("div");
    stage.className = "stage";
    panel.appendChild(stage);
  }
  const cues = (getStoryboard()?.steps || []).map((s) => s.cue);
  stage.innerHTML = `
    <div class="stage-title">${v?.title || "Visualizer"}</div>
    <div class="progress">
      ${cues
        .map((cue, idx) => {
          const active = stateInfo.step === idx;
          return `<div class="dot ${active ? "active" : ""}" title="${cue}"></div>`;
        })
        .join("")}
    </div>
    <div class="stage-text">${stateInfo.activeCue || "Waiting for events..."}</div>
    ${stateInfo.codeRef ? `<div class="stage-text">code: ${stateInfo.codeRef}</div>` : ""}
  `;
}

function populateSelectors() {
  const sbSelect = el("storyboard-select");
  const visSelect = el("visualizer-select");
  const ch = state.current;
  sbSelect.innerHTML = "";
  (ch.animations || []).forEach((sb, idx) => {
    const opt = document.createElement("option");
    opt.value = idx;
    opt.textContent = sb.title || `Storyboard ${idx + 1}`;
    sbSelect.appendChild(opt);
  });
  sbSelect.value = state.storyboardIndex;
  sbSelect.onchange = () => {
    state.storyboardIndex = Number(sbSelect.value);
    renderStoryboard(getStoryboard());
    renderEditor();
    // Mount animation if available
    const sb = getStoryboard();
    if (sb && sb.id) {
      mountAnimation(sb.id);
    }
  };

  // Mount initial animation
  const initialSb = getStoryboard();
  if (initialSb && initialSb.id) {
    mountAnimation(initialSb.id);
  }

  visSelect.innerHTML = "";
  (ch.visualizers || []).forEach((v, idx) => {
    const opt = document.createElement("option");
    opt.value = idx;
    opt.textContent = v.title || `Visualizer ${idx + 1}`;
    visSelect.appendChild(opt);
  });
  visSelect.value = state.visualizerIndex;
  visSelect.onchange = () => {
    state.visualizerIndex = Number(visSelect.value);
    mountVisualizer(getVisualizer());
  };
}

function getStoryboard() {
  const ch = state.current;
  if (!ch || !ch.animations || !ch.animations.length) return null;
  return ch.animations[state.storyboardIndex] || ch.animations[0];
}

function getVisualizer() {
  const ch = state.current;
  if (!ch || !ch.visualizers || !ch.visualizers.length) return null;
  return ch.visualizers[state.visualizerIndex] || ch.visualizers[0];
}

function toggleEdit() {
  state.editMode = !state.editMode;
  el("edit-panel").classList.toggle("hidden", !state.editMode);
  renderEditor();
}

function renderEditor() {
  if (!state.editMode || !state.current) return;
  el("summary-input").value = state.current.summary || "";
  el("objectives-input").value = (state.current.objectives || []).join("\n");
  const sb = getStoryboard();
  const container = el("edit-storyboard-steps");
  container.innerHTML = "";
  if (!sb) {
    container.innerHTML = "<p class='placeholder'>No storyboard to edit.</p>";
    return;
  }
  sb.steps.forEach((step, idx) => {
    const node = document.createElement("div");
    node.className = "edit-step";
    node.innerHTML = `
      <div class="row">
        <label>Cue<input type="text" data-field="cue" data-idx="${idx}" value="${step.cue || ""}"></label>
        <label>Duration<input type="text" data-field="duration" data-idx="${idx}" value="${step.duration || ""}"></label>
      </div>
      <label>Narration<input type="text" data-field="narration" data-idx="${idx}" value="${step.narration || ""}"></label>
      <label>Visual hint<input type="text" data-field="visualHint" data-idx="${idx}" value="${step.visualHint || ""}"></label>
    `;
    container.appendChild(node);
  });
}

function saveEdit() {
  const ch = state.current;
  if (!ch) return;
  ch.summary = el("summary-input").value;
  ch.objectives = el("objectives-input").value.split("\n").map((s) => s.trim()).filter(Boolean);
  const sb = getStoryboard();
  if (sb) {
    const inputs = el("edit-storyboard-steps").querySelectorAll("input");
    inputs.forEach((input) => {
      const idx = Number(input.dataset.idx);
      const field = input.dataset.field;
      if (!sb.steps[idx]) return;
      sb.steps[idx][field] = input.value;
    });
  }
  renderChapters();
  selectChapter(ch);
  state.editMode = false;
  el("edit-panel").classList.add("hidden");
}

function cancelEdit() {
  state.editMode = false;
  el("edit-panel").classList.add("hidden");
}

function downloadJSON() {
  const data = JSON.stringify(state.chapters, null, 2);
  const blob = new Blob([data], { type: "application/json" });
  const url = URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = "chapters.json";
  a.click();
  URL.revokeObjectURL(url);
}

function pad(n) {
  return String(n).padStart(2, "0");
}

function parseDuration(text) {
  if (!text) return 1000;
  const units = { ns: 1e-6, us: 1e-3, "µs": 1e-3, ms: 1, s: 1000, m: 60000, h: 3600000 };
  let total = 0;
  const regex = /([\d.]+)(ns|us|µs|ms|s|m|h)/g;
  let match;
  while ((match = regex.exec(text)) !== null) {
    const value = parseFloat(match[1]);
    const unit = units[match[2]] || 0;
    total += value * unit;
  }
  return total || 1000;
}

el("play-storyboard").onclick = playStoryboard;
el("stop-storyboard").onclick = stopStoryboard;
el("reload").onclick = loadChapters;
el("toggle-edit").onclick = toggleEdit;
el("save-edit").onclick = saveEdit;
el("cancel-edit").onclick = cancelEdit;
el("download-json").onclick = downloadJSON;

// Initialize animation controls
setupAnimationControls();

loadChapters().catch((err) => {
  console.error(err);
  alert("Failed to load chapters");
});
