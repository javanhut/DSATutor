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

// DFS (Depth-First Search) Animator
class DFSAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.graph = {
      'A': ['B', 'C'],
      'B': ['D', 'E'],
      'C': ['F'],
      'D': [],
      'E': ['F'],
      'F': []
    };
    this.nodePositions = {
      'A': { x: 150, y: 30 },
      'B': { x: 75, y: 100 },
      'C': { x: 225, y: 100 },
      'D': { x: 40, y: 170 },
      'E': { x: 110, y: 170 },
      'F': { x: 225, y: 170 }
    };
    this.startNode = 'A';
    this.visited = new Set();
    this.stack = [];
    this.currentNode = null;
    this.visitOrder = [];

    this.setCode(`def dfs(start, graph):
    visited = set()
    stack = [start]
    order = []
    while stack:
        node = stack.pop()
        if node in visited:
            continue
        visited.add(node)
        order.append(node)
        # Add neighbors in reverse for left-to-right order
        for nbr in reversed(graph.get(node, [])):
            if nbr not in visited:
                stack.append(nbr)
    return order`);
  }

  buildSteps() {
    this.steps = [];
    const visited = new Set();
    const stack = [this.startNode];
    const order = [];

    this.steps.push({
      lineNum: 3,
      state: `Initialize: stack=[${this.startNode}]`,
      visited: new Set(), stack: [...stack], currentNode: null, order: [],
      apply: () => { this.visited = new Set(); this.stack = [...stack]; this.visitOrder = []; }
    });

    while (stack.length > 0) {
      const node = stack.pop();

      this.steps.push({
        lineNum: 5,
        state: `Pop '${node}' from stack`,
        visited: new Set(visited), stack: [...stack], currentNode: node, order: [...order], popping: node,
        apply: () => { this.currentNode = node; }
      });

      if (visited.has(node)) {
        this.steps.push({
          lineNums: [6, 7],
          state: `'${node}' already visited, skip`,
          visited: new Set(visited), stack: [...stack], currentNode: node, order: [...order],
          apply: () => {}
        });
        continue;
      }

      visited.add(node);
      order.push(node);

      this.steps.push({
        lineNums: [8, 9],
        state: `Visit '${node}', order: [${order.join(', ')}]`,
        visited: new Set(visited), stack: [...stack], currentNode: node, order: [...order],
        apply: () => { this.visited.add(node); this.visitOrder = [...order]; }
      });

      const neighbors = [...(this.graph[node] || [])].reverse();
      for (const nbr of neighbors) {
        if (!visited.has(nbr)) {
          stack.push(nbr);
          this.steps.push({
            lineNums: [11, 12, 13],
            state: `Push '${nbr}' to stack. Stack: [${stack.join(', ')}]`,
            visited: new Set(visited), stack: [...stack], currentNode: node, order: [...order], pushing: nbr,
            apply: () => { this.stack = [...stack]; }
          });
        }
      }
    }

    this.steps.push({
      lineNum: 14,
      state: `DFS complete! Order: [${order.join(', ')}]`,
      visited: new Set(visited), stack: [], currentNode: null, order: [...order], done: true,
      apply: () => { this.stack = []; this.currentNode = null; }
    });
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { visited = this.visited, stack = this.stack, currentNode, order = this.visitOrder, popping, pushing, done } = step;

    let html = '<svg class="graph-viz" viewBox="0 0 300 220">';

    // Draw edges
    const drawnEdges = new Set();
    Object.entries(this.graph).forEach(([node, neighbors]) => {
      neighbors.forEach(nbr => {
        const p1 = this.nodePositions[node];
        const p2 = this.nodePositions[nbr];
        html += `<line x1="${p1.x}" y1="${p1.y}" x2="${p2.x}" y2="${p2.y}" class="edge" marker-end="url(#arrowhead)"/>`;
      });
    });

    // Arrowhead marker
    html += `<defs><marker id="arrowhead" markerWidth="10" markerHeight="7" refX="9" refY="3.5" orient="auto">
      <polygon points="0 0, 10 3.5, 0 7" fill="var(--border-light)"/>
    </marker></defs>`;

    // Draw nodes
    Object.entries(this.nodePositions).forEach(([node, pos]) => {
      let nodeClass = 'graph-node';
      if (visited.has(node)) nodeClass += ' visited';
      if (node === currentNode) nodeClass += ' current';
      if (stack.includes(node)) nodeClass += ' in-stack';
      if (node === popping) nodeClass += ' popping';
      if (node === pushing) nodeClass += ' pushing';

      html += `<circle cx="${pos.x}" cy="${pos.y}" r="18" class="${nodeClass}"/>`;
      html += `<text x="${pos.x}" y="${pos.y + 5}" class="node-label">${node}</text>`;
    });

    html += '</svg>';

    // Stack display
    html += '<div class="stack-display">';
    html += '<span class="stack-label">Stack (LIFO):</span>';
    html += '<div class="stack-items">';
    if (stack.length === 0) {
      html += '<span class="stack-empty">empty</span>';
    } else {
      [...stack].reverse().forEach((item, idx) => {
        html += `<span class="stack-item ${idx === 0 ? 'top' : ''}">${item}</span>`;
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
    return {
      'current': step.currentNode || '-',
      'visited': Array.from(step.visited || this.visited).join(', ') || '-',
      'stack': (step.stack || this.stack).join(', ') || 'empty',
      'order': (step.order || this.visitOrder).join(' -> ') || '-'
    };
  }

  getInputData() {
    return { 'Start': this.startNode, 'Graph': this.graph };
  }
}

// Linked List Animator
class LinkedListAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.nodes = [
      { val: 1, id: 0 },
      { val: 3, id: 1 },
      { val: 5, id: 2 },
      { val: 7, id: 3 },
      { val: 9, id: 4 }
    ];
    this.head = 0;
    this.current = -1;
    this.target = 5;
    this.found = false;

    this.setCode(`class Node:
    def __init__(self, val, next=None):
        self.val = val
        self.next = next

def search(head, target):
    current = head
    while current is not None:
        if current.val == target:
            return current  # Found!
        current = current.next
    return None  # Not found`);
  }

  buildSteps() {
    this.steps = [];

    this.steps.push({
      lineNum: 7,
      state: `Start search for ${this.target}, current = head`,
      current: 0, found: false,
      apply: () => { this.current = 0; this.found = false; }
    });

    for (let i = 0; i < this.nodes.length; i++) {
      this.steps.push({
        lineNum: 8,
        state: `Check: current (${this.nodes[i].val}) is not None? Yes`,
        current: i, found: false,
        apply: () => { this.current = i; }
      });

      this.steps.push({
        lineNum: 9,
        state: `Compare: ${this.nodes[i].val} == ${this.target}? ${this.nodes[i].val === this.target ? 'Yes!' : 'No'}`,
        current: i, found: this.nodes[i].val === this.target, comparing: true,
        apply: () => {}
      });

      if (this.nodes[i].val === this.target) {
        this.steps.push({
          lineNum: 10,
          state: `Found ${this.target} at node ${i}!`,
          current: i, found: true, done: true,
          apply: () => { this.found = true; }
        });
        return;
      }

      this.steps.push({
        lineNum: 11,
        state: `Move to next node`,
        current: i, found: false, moving: true,
        apply: () => {}
      });
    }

    this.steps.push({
      lineNum: 12,
      state: `Reached end (None), ${this.target} not found`,
      current: -1, found: false, done: true,
      apply: () => { this.current = -1; }
    });
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { current, found, comparing, moving, done } = step;

    let html = '<div class="linked-list-viz">';

    // Head pointer
    html += '<div class="list-head-ptr">head</div>';

    // Nodes
    html += '<div class="list-nodes">';
    this.nodes.forEach((node, idx) => {
      let nodeClass = 'list-node';
      if (idx === current) nodeClass += ' current';
      if (idx === current && found) nodeClass += ' found';
      if (idx === current && comparing) nodeClass += ' comparing';
      if (idx < current) nodeClass += ' visited';

      html += `<div class="${nodeClass}">
        <div class="node-val">${node.val}</div>
        <div class="node-ptr">${idx < this.nodes.length - 1 ? '->' : 'None'}</div>
      </div>`;

      if (idx < this.nodes.length - 1) {
        html += '<div class="list-arrow">-></div>';
      }
    });
    html += '<div class="list-null">None</div>';
    html += '</div>';

    // Current pointer
    if (current >= 0 && current < this.nodes.length) {
      html += `<div class="current-ptr" style="left: calc(${current} * 90px + 45px)">current</div>`;
    }

    html += '</div>';
    html += `<div class="target-display">Target: <strong>${this.target}</strong></div>`;

    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    const curr = step.current;
    return {
      'current': curr >= 0 && curr < this.nodes.length ? `Node(${this.nodes[curr].val})` : 'None',
      'current.val': curr >= 0 && curr < this.nodes.length ? this.nodes[curr].val : '-',
      'target': this.target,
      'found': step.found ? 'Yes' : 'No'
    };
  }

  getInputData() {
    return this.nodes.map(n => n.val);
  }
}

// Two Pointers (Floyd's Cycle Detection) Animator
class TwoPointersAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    // Linked list with cycle: 1->2->3->4->5->6->3 (cycle back to 3)
    this.nodes = [
      { val: 1, next: 1 },
      { val: 2, next: 2 },
      { val: 3, next: 3 },
      { val: 4, next: 4 },
      { val: 5, next: 5 },
      { val: 6, next: 2 }  // Points back to node index 2 (val=3)
    ];
    this.slow = 0;
    this.fast = 0;
    this.hasCycle = false;

    this.setCode(`def has_cycle(head):
    if not head or not head.next:
        return False
    slow = head
    fast = head
    while fast and fast.next:
        slow = slow.next        # Move 1 step
        fast = fast.next.next   # Move 2 steps
        if slow == fast:
            return True  # Cycle detected!
    return False  # No cycle`);
  }

  buildSteps() {
    this.steps = [];
    let slow = 0;
    let fast = 0;

    this.steps.push({
      lineNums: [4, 5],
      state: 'Initialize: slow = head, fast = head',
      slow: 0, fast: 0, hasCycle: false,
      apply: () => { this.slow = 0; this.fast = 0; }
    });

    let iterations = 0;
    const maxIterations = 20;

    while (iterations < maxIterations) {
      iterations++;

      // Check while condition
      const fastNode = this.nodes[fast];
      const fastNextIdx = fastNode ? fastNode.next : -1;
      const canContinue = fast < this.nodes.length && fastNextIdx < this.nodes.length;

      this.steps.push({
        lineNum: 6,
        state: `Check: fast and fast.next exist? ${canContinue ? 'Yes' : 'No'}`,
        slow, fast, hasCycle: false,
        apply: () => {}
      });

      if (!canContinue) break;

      // Move slow
      const oldSlow = slow;
      slow = this.nodes[slow].next;
      this.steps.push({
        lineNum: 7,
        state: `slow = slow.next (${this.nodes[oldSlow].val} -> ${this.nodes[slow].val})`,
        slow, fast, hasCycle: false, slowMoving: true,
        apply: () => { this.slow = slow; }
      });

      // Move fast (2 steps)
      const oldFast = fast;
      fast = this.nodes[fast].next;
      if (fast < this.nodes.length) {
        fast = this.nodes[fast].next;
      }
      this.steps.push({
        lineNum: 8,
        state: `fast = fast.next.next (${this.nodes[oldFast].val} -> ${this.nodes[fast].val})`,
        slow, fast, hasCycle: false, fastMoving: true,
        apply: () => { this.fast = fast; }
      });

      // Check if they meet
      this.steps.push({
        lineNum: 9,
        state: `Check: slow == fast? (${this.nodes[slow].val} == ${this.nodes[fast].val}) ${slow === fast ? 'Yes!' : 'No'}`,
        slow, fast, hasCycle: slow === fast, checking: true,
        apply: () => {}
      });

      if (slow === fast) {
        this.steps.push({
          lineNum: 10,
          state: 'Cycle detected! Pointers met.',
          slow, fast, hasCycle: true, done: true,
          apply: () => { this.hasCycle = true; }
        });
        return;
      }
    }

    this.steps.push({
      lineNum: 11,
      state: 'No cycle found (fast reached end)',
      slow, fast, hasCycle: false, done: true,
      apply: () => {}
    });
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { slow, fast, hasCycle, slowMoving, fastMoving, checking, done } = step;

    let html = '<div class="two-pointers-viz">';

    // Draw nodes in a line with cycle arrow
    html += '<div class="cycle-nodes">';
    this.nodes.forEach((node, idx) => {
      let nodeClass = 'cycle-node';
      if (idx === slow && idx === fast) nodeClass += ' both-ptrs';
      else if (idx === slow) nodeClass += ' slow-ptr';
      else if (idx === fast) nodeClass += ' fast-ptr';
      if (hasCycle && idx === slow) nodeClass += ' cycle-found';

      html += `<div class="${nodeClass}">
        <div class="node-val">${node.val}</div>
        <div class="node-idx">${idx}</div>
      </div>`;
    });
    html += '</div>';

    // Cycle arrow indicator
    html += '<div class="cycle-arrow">Cycle: node 5 -> node 2</div>';

    // Pointer labels
    html += '<div class="pointer-info">';
    html += `<span class="slow-label">Slow (1x): ${slow < this.nodes.length ? this.nodes[slow].val : '-'}</span>`;
    html += `<span class="fast-label">Fast (2x): ${fast < this.nodes.length ? this.nodes[fast].val : '-'}</span>`;
    html += '</div>';

    if (hasCycle) {
      html += '<div class="cycle-result found">Cycle Detected!</div>';
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'slow': step.slow < this.nodes.length ? `Node(${this.nodes[step.slow].val})` : '-',
      'fast': step.fast < this.nodes.length ? `Node(${this.nodes[step.fast].val})` : '-',
      'slow position': step.slow,
      'fast position': step.fast,
      'cycle found': step.hasCycle ? 'Yes' : 'No'
    };
  }

  getInputData() {
    return { 'List': '1->2->3->4->5->6->3 (cycle)', 'Cycle at': 'Node 3' };
  }
}

// Quicksort Animator
class QuicksortAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.originalArray = [38, 27, 43, 3, 9, 82, 10];
    this.array = [...this.originalArray];
    this.pivot = -1;
    this.left = -1;
    this.right = -1;
    this.partitionStart = 0;
    this.partitionEnd = this.array.length - 1;

    this.setCode(`def quicksort(arr, lo, hi):
    if lo < hi:
        pivot_idx = partition(arr, lo, hi)
        quicksort(arr, lo, pivot_idx - 1)
        quicksort(arr, pivot_idx + 1, hi)

def partition(arr, lo, hi):
    pivot = arr[hi]  # Choose last as pivot
    i = lo - 1
    for j in range(lo, hi):
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i + 1], arr[hi] = arr[hi], arr[i + 1]
    return i + 1`);
  }

  buildSteps() {
    this.steps = [];
    const arr = [...this.originalArray];

    this.steps.push({
      lineNum: 1,
      state: `Initial array: [${arr.join(', ')}]`,
      array: [...arr], pivot: -1, i: -1, j: -1, lo: 0, hi: arr.length - 1,
      apply: () => { this.array = [...arr]; }
    });

    this.quicksortSteps(arr, 0, arr.length - 1, 0);
  }

  quicksortSteps(arr, lo, hi, depth) {
    if (lo >= hi) return;

    // Show partition start
    const pivot = arr[hi];
    const pivotIdx = hi;

    this.steps.push({
      lineNum: 8,
      state: `Partition [${lo}:${hi}], pivot = ${pivot}`,
      array: [...arr], pivot: pivotIdx, i: lo - 1, j: lo, lo, hi, depth,
      apply: () => {}
    });

    let i = lo - 1;
    for (let j = lo; j < hi; j++) {
      this.steps.push({
        lineNum: 10,
        state: `j=${j}: arr[${j}]=${arr[j]} <= pivot(${pivot})? ${arr[j] <= pivot ? 'Yes' : 'No'}`,
        array: [...arr], pivot: pivotIdx, i, j, lo, hi, depth, comparing: j,
        apply: () => {}
      });

      if (arr[j] <= pivot) {
        i++;
        if (i !== j) {
          [arr[i], arr[j]] = [arr[j], arr[i]];
          this.steps.push({
            lineNums: [11, 12, 13],
            state: `Swap arr[${i}] and arr[${j}]: [${arr.join(', ')}]`,
            array: [...arr], pivot: pivotIdx, i, j, lo, hi, depth, swapping: [i, j],
            apply: () => { this.array = [...arr]; }
          });
        }
      }
    }

    // Final swap: pivot into place
    [arr[i + 1], arr[hi]] = [arr[hi], arr[i + 1]];
    const newPivotIdx = i + 1;

    this.steps.push({
      lineNum: 14,
      state: `Place pivot: swap arr[${i + 1}] and arr[${hi}]: [${arr.join(', ')}]`,
      array: [...arr], pivot: newPivotIdx, i: i + 1, j: -1, lo, hi, depth, pivotPlaced: true,
      apply: () => { this.array = [...arr]; }
    });

    // Recurse left
    if (lo < newPivotIdx - 1) {
      this.quicksortSteps(arr, lo, newPivotIdx - 1, depth + 1);
    }
    // Recurse right
    if (newPivotIdx + 1 < hi) {
      this.quicksortSteps(arr, newPivotIdx + 1, hi, depth + 1);
    }
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { array, pivot, i, j, lo, hi, swapping = [], comparing, pivotPlaced } = step;
    const arr = array || this.array;

    let html = '<div class="quicksort-viz">';
    html += '<div class="qs-array">';

    arr.forEach((val, idx) => {
      let cellClass = 'qs-cell';
      if (idx === pivot) cellClass += ' pivot';
      if (idx === i) cellClass += ' i-ptr';
      if (idx === comparing) cellClass += ' comparing';
      if (swapping.includes(idx)) cellClass += ' swapping';
      if (idx < lo || idx > hi) cellClass += ' inactive';
      if (pivotPlaced && idx === pivot) cellClass += ' placed';

      const height = Math.max(25, val * 2.5);
      html += `<div class="${cellClass}" style="height: ${height}px">
        <div class="cell-value">${val}</div>
      </div>`;
    });

    html += '</div>';

    // Pointer indicators
    html += '<div class="qs-pointers">';
    if (lo !== undefined && hi !== undefined) {
      html += `<span>Range: [${lo}, ${hi}]</span>`;
    }
    if (pivot >= 0) {
      html += `<span class="pivot-info">Pivot: ${arr[pivot]}</span>`;
    }
    html += '</div>';

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    const arr = step.array || this.array;
    return {
      'pivot': step.pivot >= 0 ? arr[step.pivot] : '-',
      'i': step.i >= 0 ? step.i : '-',
      'j': step.j >= 0 ? step.j : '-',
      'range': `[${step.lo}, ${step.hi}]`,
      'array': arr.join(', ')
    };
  }

  getInputData() {
    return this.originalArray;
  }
}

// Hash Table Animator
class HashTableAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.bucketCount = 7;
    this.buckets = Array(this.bucketCount).fill(null).map(() => []);
    this.insertions = [
      { key: 'apple', val: 5 },
      { key: 'banana', val: 3 },
      { key: 'cherry', val: 8 },
      { key: 'date', val: 2 },
      { key: 'elder', val: 6 },  // Will collide
      { key: 'fig', val: 4 }
    ];
    this.currentKey = null;
    this.currentBucket = -1;

    this.setCode(`class HashTable:
    def __init__(self, size=7):
        self.buckets = [[] for _ in range(size)]

    def _hash(self, key):
        return hash(key) % len(self.buckets)

    def insert(self, key, value):
        idx = self._hash(key)
        bucket = self.buckets[idx]
        for i, (k, v) in enumerate(bucket):
            if k == key:
                bucket[i] = (key, value)  # Update
                return
        bucket.append((key, value))  # Insert`);
  }

  buildSteps() {
    this.steps = [];
    const buckets = Array(this.bucketCount).fill(null).map(() => []);

    this.steps.push({
      lineNums: [2, 3],
      state: `Initialize hash table with ${this.bucketCount} buckets`,
      buckets: buckets.map(b => [...b]), currentKey: null, currentBucket: -1,
      apply: () => { this.buckets = buckets.map(b => [...b]); }
    });

    for (const item of this.insertions) {
      const hashVal = this.simpleHash(item.key);
      const bucketIdx = hashVal % this.bucketCount;

      this.steps.push({
        lineNum: 9,
        state: `Insert '${item.key}': hash('${item.key}') = ${hashVal}`,
        buckets: buckets.map(b => [...b]), currentKey: item.key, currentBucket: -1, hashing: true,
        apply: () => { this.currentKey = item.key; }
      });

      this.steps.push({
        lineNum: 9,
        state: `Bucket index = ${hashVal} % ${this.bucketCount} = ${bucketIdx}`,
        buckets: buckets.map(b => [...b]), currentKey: item.key, currentBucket: bucketIdx,
        apply: () => { this.currentBucket = bucketIdx; }
      });

      const existingIdx = buckets[bucketIdx].findIndex(([k]) => k === item.key);
      if (existingIdx >= 0) {
        buckets[bucketIdx][existingIdx] = [item.key, item.val];
        this.steps.push({
          lineNums: [12, 13],
          state: `Key exists, update value to ${item.val}`,
          buckets: buckets.map(b => [...b]), currentKey: item.key, currentBucket: bucketIdx, updating: true,
          apply: () => { this.buckets = buckets.map(b => [...b]); }
        });
      } else {
        const hasCollision = buckets[bucketIdx].length > 0;
        buckets[bucketIdx].push([item.key, item.val]);
        this.steps.push({
          lineNum: 15,
          state: hasCollision
            ? `Collision! Chain '${item.key}' in bucket ${bucketIdx}`
            : `Insert '${item.key}' -> ${item.val} in bucket ${bucketIdx}`,
          buckets: buckets.map(b => [...b]), currentKey: item.key, currentBucket: bucketIdx,
          collision: hasCollision, inserting: true,
          apply: () => { this.buckets = buckets.map(b => [...b]); }
        });
      }
    }

    this.steps.push({
      lineNum: 15,
      state: 'All items inserted!',
      buckets: buckets.map(b => [...b]), currentKey: null, currentBucket: -1, done: true,
      apply: () => {}
    });
  }

  simpleHash(key) {
    let hash = 0;
    for (let i = 0; i < key.length; i++) {
      hash = (hash * 31 + key.charCodeAt(i)) >>> 0;
    }
    return hash;
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { buckets, currentBucket, collision, inserting, hashing } = step;
    const bkts = buckets || this.buckets;

    let html = '<div class="hashtable-viz">';

    bkts.forEach((bucket, idx) => {
      let bucketClass = 'ht-bucket';
      if (idx === currentBucket) bucketClass += ' active';
      if (idx === currentBucket && collision) bucketClass += ' collision';

      html += `<div class="${bucketClass}">
        <div class="bucket-idx">${idx}</div>
        <div class="bucket-chain">`;

      if (bucket.length === 0) {
        html += '<span class="bucket-empty">empty</span>';
      } else {
        bucket.forEach(([k, v], i) => {
          let entryClass = 'bucket-entry';
          if (idx === currentBucket && i === bucket.length - 1 && inserting) {
            entryClass += ' new-entry';
          }
          html += `<div class="${entryClass}">${k}: ${v}</div>`;
        });
      }

      html += '</div></div>';
    });

    html += '</div>';

    if (step.currentKey) {
      html += `<div class="hash-info">Key: <strong>${step.currentKey}</strong></div>`;
    }

    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'key': step.currentKey || '-',
      'bucket': step.currentBucket >= 0 ? step.currentBucket : '-',
      'collision': step.collision ? 'Yes' : 'No',
      'total items': (step.buckets || this.buckets).reduce((sum, b) => sum + b.length, 0)
    };
  }

  getInputData() {
    return this.insertions.map(i => `${i.key}: ${i.val}`);
  }
}

// Dijkstra's Algorithm Animator
class DijkstraAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    // Weighted graph
    this.graph = {
      'A': [['B', 4], ['C', 2]],
      'B': [['C', 1], ['D', 5]],
      'C': [['D', 8], ['E', 10]],
      'D': [['E', 2]],
      'E': []
    };
    this.nodePositions = {
      'A': { x: 50, y: 100 },
      'B': { x: 130, y: 40 },
      'C': { x: 130, y: 160 },
      'D': { x: 220, y: 100 },
      'E': { x: 300, y: 100 }
    };
    this.start = 'A';
    this.dist = {};
    this.prev = {};
    this.visited = new Set();
    this.current = null;

    this.setCode(`import heapq

def dijkstra(graph, start):
    dist = {v: float('inf') for v in graph}
    dist[start] = 0
    prev = {v: None for v in graph}
    pq = [(0, start)]
    visited = set()

    while pq:
        d, u = heapq.heappop(pq)
        if u in visited:
            continue
        visited.add(u)
        for v, w in graph[u]:
            if dist[u] + w < dist[v]:
                dist[v] = dist[u] + w
                prev[v] = u
                heapq.heappush(pq, (dist[v], v))
    return dist, prev`);
  }

  buildSteps() {
    this.steps = [];
    const nodes = Object.keys(this.graph);
    const dist = {};
    const prev = {};
    const visited = new Set();
    const pq = [];

    // Initialize
    nodes.forEach(v => {
      dist[v] = v === this.start ? 0 : Infinity;
      prev[v] = null;
    });
    pq.push([0, this.start]);

    this.steps.push({
      lineNums: [4, 5, 6, 7],
      state: `Initialize: dist[${this.start}]=0, others=Inf`,
      dist: {...dist}, prev: {...prev}, visited: new Set(), current: null, pq: [...pq],
      apply: () => { this.dist = {...dist}; this.prev = {...prev}; this.visited = new Set(); }
    });

    while (pq.length > 0) {
      pq.sort((a, b) => a[0] - b[0]);
      const [d, u] = pq.shift();

      this.steps.push({
        lineNum: 11,
        state: `Pop (${d}, '${u}') from priority queue`,
        dist: {...dist}, prev: {...prev}, visited: new Set(visited), current: u, pq: [...pq], popping: u,
        apply: () => { this.current = u; }
      });

      if (visited.has(u)) {
        this.steps.push({
          lineNums: [12, 13],
          state: `'${u}' already visited, skip`,
          dist: {...dist}, prev: {...prev}, visited: new Set(visited), current: u, pq: [...pq],
          apply: () => {}
        });
        continue;
      }

      visited.add(u);
      this.steps.push({
        lineNum: 14,
        state: `Mark '${u}' as visited`,
        dist: {...dist}, prev: {...prev}, visited: new Set(visited), current: u, pq: [...pq],
        apply: () => { this.visited.add(u); }
      });

      // Relax edges
      for (const [v, w] of this.graph[u]) {
        const newDist = dist[u] + w;

        this.steps.push({
          lineNum: 16,
          state: `Check edge ${u}->${v} (w=${w}): ${dist[u]}+${w}=${newDist} < ${dist[v] === Infinity ? 'Inf' : dist[v]}?`,
          dist: {...dist}, prev: {...prev}, visited: new Set(visited), current: u, pq: [...pq],
          checking: v, edgeWeight: w,
          apply: () => {}
        });

        if (newDist < dist[v]) {
          dist[v] = newDist;
          prev[v] = u;
          pq.push([newDist, v]);

          this.steps.push({
            lineNums: [17, 18, 19],
            state: `Relax! dist[${v}] = ${newDist}, prev[${v}] = ${u}`,
            dist: {...dist}, prev: {...prev}, visited: new Set(visited), current: u, pq: [...pq],
            relaxed: v,
            apply: () => { this.dist = {...dist}; this.prev = {...prev}; }
          });
        }
      }
    }

    this.steps.push({
      lineNum: 20,
      state: `Done! Shortest distances from '${this.start}'`,
      dist: {...dist}, prev: {...prev}, visited: new Set(visited), current: null, pq: [], done: true,
      apply: () => {}
    });
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { dist, visited, current, checking, relaxed, done } = step;

    let html = '<svg class="dijkstra-viz" viewBox="0 0 350 200">';

    // Draw edges with weights
    Object.entries(this.graph).forEach(([u, edges]) => {
      edges.forEach(([v, w]) => {
        const p1 = this.nodePositions[u];
        const p2 = this.nodePositions[v];
        const midX = (p1.x + p2.x) / 2;
        const midY = (p1.y + p2.y) / 2;

        let edgeClass = 'edge';
        if (checking === v && current === u) edgeClass += ' checking';
        if (relaxed === v && current === u) edgeClass += ' relaxed';

        html += `<line x1="${p1.x}" y1="${p1.y}" x2="${p2.x}" y2="${p2.y}" class="${edgeClass}"/>`;
        html += `<text x="${midX}" y="${midY - 5}" class="edge-weight">${w}</text>`;
      });
    });

    // Draw nodes
    Object.entries(this.nodePositions).forEach(([node, pos]) => {
      let nodeClass = 'dijkstra-node';
      if (visited && visited.has(node)) nodeClass += ' visited';
      if (node === current) nodeClass += ' current';
      if (node === checking) nodeClass += ' checking';
      if (node === relaxed) nodeClass += ' relaxed';

      const d = dist ? (dist[node] === Infinity ? 'Inf' : dist[node]) : 'Inf';

      html += `<circle cx="${pos.x}" cy="${pos.y}" r="20" class="${nodeClass}"/>`;
      html += `<text x="${pos.x}" y="${pos.y + 4}" class="node-label">${node}</text>`;
      html += `<text x="${pos.x}" y="${pos.y + 32}" class="dist-label">${d}</text>`;
    });

    html += '</svg>';

    // Distance table
    html += '<div class="dist-table">';
    Object.keys(this.graph).forEach(node => {
      const d = dist ? (dist[node] === Infinity ? 'Inf' : dist[node]) : 'Inf';
      let cellClass = 'dist-cell';
      if (visited && visited.has(node)) cellClass += ' final';
      html += `<div class="${cellClass}"><span>${node}</span><span>${d}</span></div>`;
    });
    html += '</div>';

    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    const dist = step.dist || this.dist;
    return {
      'current': step.current || '-',
      'dist': Object.entries(dist).map(([k, v]) => `${k}:${v === Infinity ? 'Inf' : v}`).join(', '),
      'visited': Array.from(step.visited || this.visited).join(', ') || '-'
    };
  }

  getInputData() {
    return { 'Start': this.start, 'Graph': 'Weighted directed' };
  }
}

// Memoization (Fibonacci) Animator
class MemoizationAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.n = 6;
    this.memo = {};
    this.callStack = [];
    this.currentCall = null;

    this.setCode(`def fib(n, memo={}):
    if n in memo:
        return memo[n]
    if n <= 1:
        return n
    memo[n] = fib(n-1, memo) + fib(n-2, memo)
    return memo[n]

# Call: fib(6)`);
  }

  buildSteps() {
    this.steps = [];
    this.fibSteps(this.n, {}, []);
  }

  fibSteps(n, memo, stack) {
    const newStack = [...stack, n];

    this.steps.push({
      lineNum: 1,
      state: `Call fib(${n})`,
      memo: {...memo}, callStack: [...newStack], currentCall: n, phase: 'call',
      apply: () => { this.callStack = [...newStack]; this.currentCall = n; }
    });

    // Check memo
    if (memo[n] !== undefined) {
      this.steps.push({
        lineNums: [2, 3],
        state: `memo[${n}] exists! Return ${memo[n]}`,
        memo: {...memo}, callStack: [...newStack], currentCall: n, phase: 'memo-hit', memoHit: n,
        apply: () => {}
      });
      return memo[n];
    }

    this.steps.push({
      lineNum: 2,
      state: `memo[${n}] not found, compute it`,
      memo: {...memo}, callStack: [...newStack], currentCall: n, phase: 'compute',
      apply: () => {}
    });

    // Base case
    if (n <= 1) {
      this.steps.push({
        lineNums: [4, 5],
        state: `Base case: fib(${n}) = ${n}`,
        memo: {...memo}, callStack: [...newStack], currentCall: n, phase: 'base',
        apply: () => {}
      });
      return n;
    }

    // Recursive calls
    this.steps.push({
      lineNum: 6,
      state: `Need fib(${n-1}) + fib(${n-2})`,
      memo: {...memo}, callStack: [...newStack], currentCall: n, phase: 'recurse',
      apply: () => {}
    });

    const left = this.fibSteps(n - 1, memo, newStack);
    const right = this.fibSteps(n - 2, memo, newStack);
    const result = left + right;
    memo[n] = result;

    this.steps.push({
      lineNum: 6,
      state: `memo[${n}] = fib(${n-1}) + fib(${n-2}) = ${left} + ${right} = ${result}`,
      memo: {...memo}, callStack: [...newStack], currentCall: n, phase: 'store', storing: n,
      apply: () => { this.memo = {...memo}; }
    });

    this.steps.push({
      lineNum: 7,
      state: `Return fib(${n}) = ${result}`,
      memo: {...memo}, callStack: stack, currentCall: stack[stack.length - 1] || null, phase: 'return',
      apply: () => { this.callStack = [...stack]; }
    });

    return result;
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { memo, callStack, currentCall, phase, memoHit, storing } = step;

    let html = '<div class="memo-viz">';

    // Memo table
    html += '<div class="memo-table">';
    html += '<div class="memo-header">Memo Table</div>';
    html += '<div class="memo-entries">';
    for (let i = 0; i <= this.n; i++) {
      let cellClass = 'memo-cell';
      if (memo && memo[i] !== undefined) cellClass += ' filled';
      if (i === memoHit) cellClass += ' hit';
      if (i === storing) cellClass += ' storing';
      html += `<div class="${cellClass}">
        <span class="memo-key">fib(${i})</span>
        <span class="memo-val">${memo && memo[i] !== undefined ? memo[i] : '-'}</span>
      </div>`;
    }
    html += '</div></div>';

    // Call stack
    html += '<div class="memo-stack">';
    html += '<div class="stack-header">Call Stack</div>';
    html += '<div class="stack-frames">';
    if (callStack && callStack.length > 0) {
      [...callStack].reverse().forEach((n, idx) => {
        let frameClass = 'memo-frame';
        if (idx === 0) frameClass += ' active';
        html += `<div class="${frameClass}">fib(${n})</div>`;
      });
    } else {
      html += '<div class="stack-empty">empty</div>';
    }
    html += '</div></div>';

    html += '</div>';

    // Fibonacci sequence preview
    html += '<div class="fib-sequence">';
    html += 'Sequence: ';
    for (let i = 0; i <= this.n; i++) {
      if (memo && memo[i] !== undefined) {
        html += `<span class="fib-num">${memo[i]}</span>`;
      }
    }
    html += '</div>';

    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    const memo = step.memo || this.memo;
    return {
      'n': step.currentCall !== null ? step.currentCall : '-',
      'memo entries': Object.keys(memo).length,
      'stack depth': (step.callStack || this.callStack).length,
      'phase': step.phase || '-'
    };
  }

  getInputData() {
    return { 'Computing': `fib(${this.n})`, 'Expected': this.fibResult(this.n) };
  }

  fibResult(n) {
    if (n <= 1) return n;
    let a = 0, b = 1;
    for (let i = 2; i <= n; i++) {
      [a, b] = [b, a + b];
    }
    return b;
  }
}

// Topological Sort Animator
class TopSortAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    // DAG representing course prerequisites
    this.graph = {
      'A': ['C'],
      'B': ['C', 'D'],
      'C': ['E'],
      'D': ['E'],
      'E': ['F'],
      'F': []
    };
    this.nodePositions = {
      'A': { x: 50, y: 50 },
      'B': { x: 50, y: 150 },
      'C': { x: 150, y: 100 },
      'D': { x: 150, y: 180 },
      'E': { x: 250, y: 140 },
      'F': { x: 330, y: 140 }
    };
    this.inDegree = {};
    this.queue = [];
    this.result = [];
    this.current = null;

    this.setCode(`from collections import deque

def topsort(graph):
    in_degree = {u: 0 for u in graph}
    for u in graph:
        for v in graph[u]:
            in_degree[v] += 1

    queue = deque([u for u in graph if in_degree[u] == 0])
    result = []

    while queue:
        u = queue.popleft()
        result.append(u)
        for v in graph[u]:
            in_degree[v] -= 1
            if in_degree[v] == 0:
                queue.append(v)
    return result`);
  }

  buildSteps() {
    this.steps = [];
    const nodes = Object.keys(this.graph);
    const inDegree = {};
    nodes.forEach(u => inDegree[u] = 0);

    // Calculate in-degrees
    Object.entries(this.graph).forEach(([u, neighbors]) => {
      neighbors.forEach(v => inDegree[v]++);
    });

    this.steps.push({
      lineNums: [4, 5, 6, 7],
      state: `Calculate in-degrees: ${Object.entries(inDegree).map(([k,v]) => `${k}:${v}`).join(', ')}`,
      inDegree: {...inDegree}, queue: [], result: [], current: null,
      apply: () => { this.inDegree = {...inDegree}; }
    });

    // Initialize queue with zero in-degree nodes
    const queue = nodes.filter(u => inDegree[u] === 0);

    this.steps.push({
      lineNum: 9,
      state: `Queue nodes with in-degree 0: [${queue.join(', ')}]`,
      inDegree: {...inDegree}, queue: [...queue], result: [], current: null,
      apply: () => { this.queue = [...queue]; }
    });

    const result = [];

    while (queue.length > 0) {
      const u = queue.shift();

      this.steps.push({
        lineNum: 13,
        state: `Dequeue '${u}'`,
        inDegree: {...inDegree}, queue: [...queue], result: [...result], current: u, dequeuing: u,
        apply: () => { this.current = u; }
      });

      result.push(u);

      this.steps.push({
        lineNum: 14,
        state: `Add '${u}' to result: [${result.join(', ')}]`,
        inDegree: {...inDegree}, queue: [...queue], result: [...result], current: u,
        apply: () => { this.result = [...result]; }
      });

      // Decrease in-degree of neighbors
      for (const v of this.graph[u]) {
        inDegree[v]--;

        this.steps.push({
          lineNum: 16,
          state: `Decrement in_degree[${v}] to ${inDegree[v]}`,
          inDegree: {...inDegree}, queue: [...queue], result: [...result], current: u,
          decrementing: v,
          apply: () => { this.inDegree = {...inDegree}; }
        });

        if (inDegree[v] === 0) {
          queue.push(v);
          this.steps.push({
            lineNums: [17, 18],
            state: `in_degree[${v}] == 0, enqueue '${v}'`,
            inDegree: {...inDegree}, queue: [...queue], result: [...result], current: u,
            enqueuing: v,
            apply: () => { this.queue = [...queue]; }
          });
        }
      }
    }

    this.steps.push({
      lineNum: 19,
      state: `Topological order: [${result.join(' -> ')}]`,
      inDegree: {...inDegree}, queue: [], result: [...result], current: null, done: true,
      apply: () => {}
    });
  }

  render() {
    if (!this.canvas) return;
    const step = this.steps[this.currentStep] || {};
    const { inDegree, queue, result, current, dequeuing, enqueuing, decrementing, done } = step;

    let html = '<svg class="topsort-viz" viewBox="0 0 380 220">';

    // Draw edges
    Object.entries(this.graph).forEach(([u, neighbors]) => {
      neighbors.forEach(v => {
        const p1 = this.nodePositions[u];
        const p2 = this.nodePositions[v];
        let edgeClass = 'edge';
        if (current === u && decrementing === v) edgeClass += ' processing';
        html += `<line x1="${p1.x}" y1="${p1.y}" x2="${p2.x}" y2="${p2.y}" class="${edgeClass}" marker-end="url(#arrow)"/>`;
      });
    });

    html += `<defs><marker id="arrow" markerWidth="10" markerHeight="7" refX="9" refY="3.5" orient="auto">
      <polygon points="0 0, 10 3.5, 0 7" fill="var(--border-light)"/>
    </marker></defs>`;

    // Draw nodes with in-degree
    Object.entries(this.nodePositions).forEach(([node, pos]) => {
      let nodeClass = 'topsort-node';
      if (result && result.includes(node)) nodeClass += ' processed';
      if (node === current) nodeClass += ' current';
      if (node === dequeuing) nodeClass += ' dequeuing';
      if (node === enqueuing) nodeClass += ' enqueuing';
      if (queue && queue.includes(node)) nodeClass += ' in-queue';

      const deg = inDegree ? inDegree[node] : 0;

      html += `<circle cx="${pos.x}" cy="${pos.y}" r="20" class="${nodeClass}"/>`;
      html += `<text x="${pos.x}" y="${pos.y + 4}" class="node-label">${node}</text>`;
      html += `<text x="${pos.x + 25}" y="${pos.y - 15}" class="degree-label">${deg}</text>`;
    });

    html += '</svg>';

    // Queue and result display
    html += '<div class="topsort-info">';
    html += `<div class="ts-queue"><span>Queue:</span> [${(queue || []).join(', ') || 'empty'}]</div>`;
    html += `<div class="ts-result"><span>Result:</span> [${(result || []).join(' -> ') || ''}]</div>`;
    html += '</div>';

    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'current': step.current || '-',
      'queue': (step.queue || []).join(', ') || 'empty',
      'result': (step.result || []).join(' -> ') || '-',
      'in_degrees': Object.entries(step.inDegree || {}).map(([k,v]) => `${k}:${v}`).join(', ')
    };
  }

  getInputData() {
    return { 'DAG': 'Course prerequisites', 'Nodes': Object.keys(this.graph).join(', ') };
  }
}

// String Pattern Matching Animator (KMP-style visualization)
class StringMatchAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    this.text = "ABABDABACDABABCABAB";
    this.pattern = "ABABCABAB";
    this.textIdx = 0;
    this.patternIdx = 0;
    this.matches = [];

    this.setCode(`def string_match(text, pattern):
    n, m = len(text), len(pattern)
    for i in range(n - m + 1):
        j = 0
        while j < m and text[i + j] == pattern[j]:
            j += 1
        if j == m:
            return i  # Match found at index i
    return -1  # No match`);
  }

  buildSteps() {
    this.steps = [];
    const n = this.text.length;
    const m = this.pattern.length;

    this.steps.push({
      lineNum: 1,
      state: `Initialize: text length = ${n}, pattern length = ${m}`,
      textIdx: -1,
      patternIdx: -1,
      comparing: false,
      matches: [],
      apply: () => {}
    });

    for (let i = 0; i <= n - m; i++) {
      this.steps.push({
        lineNum: 2,
        state: `Starting comparison at text index ${i}`,
        textIdx: i,
        patternIdx: 0,
        comparing: false,
        matches: [...this.matches],
        apply: () => {}
      });

      let j = 0;
      while (j < m && this.text[i + j] === this.pattern[j]) {
        this.steps.push({
          lineNum: 4,
          state: `Comparing text[${i + j}]='${this.text[i + j]}' with pattern[${j}]='${this.pattern[j]}' - Match!`,
          textIdx: i,
          patternIdx: j,
          comparePos: i + j,
          comparing: true,
          isMatch: true,
          matches: [...this.matches],
          apply: () => {}
        });
        j++;
      }

      if (j === m) {
        this.matches.push(i);
        this.steps.push({
          lineNum: 6,
          state: `Pattern found at index ${i}!`,
          textIdx: i,
          patternIdx: j - 1,
          comparing: false,
          foundMatch: true,
          matchStart: i,
          matches: [...this.matches],
          apply: () => {}
        });
      } else if (j < m && i + j < n) {
        this.steps.push({
          lineNum: 4,
          state: `Mismatch: text[${i + j}]='${this.text[i + j]}' != pattern[${j}]='${this.pattern[j]}'`,
          textIdx: i,
          patternIdx: j,
          comparePos: i + j,
          comparing: true,
          isMatch: false,
          matches: [...this.matches],
          apply: () => {}
        });
      }
    }

    this.steps.push({
      lineNum: 7,
      state: `Search complete. Found ${this.matches.length} match(es).`,
      textIdx: -1,
      patternIdx: -1,
      comparing: false,
      matches: [...this.matches],
      apply: () => {}
    });
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { textIdx, patternIdx, comparePos, comparing, isMatch, foundMatch, matchStart, matches } = step;

    let html = '<div class="string-match-viz">';

    // Text display
    html += '<div class="string-row"><span class="string-label">Text:</span><div class="string-chars">';
    for (let i = 0; i < this.text.length; i++) {
      let charClass = 'string-char';
      if (matches && matches.some(m => i >= m && i < m + this.pattern.length)) {
        charClass += ' matched';
      }
      if (comparing && i === comparePos) {
        charClass += isMatch ? ' comparing match' : ' comparing mismatch';
      }
      if (foundMatch && i >= matchStart && i < matchStart + this.pattern.length) {
        charClass += ' found';
      }
      html += `<span class="${charClass}">${this.text[i]}<span class="char-idx">${i}</span></span>`;
    }
    html += '</div></div>';

    // Pattern display
    html += '<div class="string-row"><span class="string-label">Pattern:</span><div class="string-chars">';
    const offset = textIdx >= 0 ? textIdx : 0;
    for (let i = 0; i < offset && textIdx >= 0; i++) {
      html += '<span class="string-char spacer"></span>';
    }
    for (let i = 0; i < this.pattern.length; i++) {
      let charClass = 'string-char pattern';
      if (comparing && i === patternIdx) {
        charClass += isMatch ? ' comparing match' : ' comparing mismatch';
      }
      if (foundMatch) {
        charClass += ' found';
      }
      html += `<span class="${charClass}">${this.pattern[i]}</span>`;
    }
    html += '</div></div>';

    // Match positions
    if (matches && matches.length > 0) {
      html += `<div class="match-info">Matches found at indices: [${matches.join(', ')}]</div>`;
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'text_index': step.textIdx >= 0 ? step.textIdx : '-',
      'pattern_index': step.patternIdx >= 0 ? step.patternIdx : '-',
      'matches_found': (step.matches || []).length
    };
  }

  getInputData() {
    return { 'Text': this.text, 'Pattern': this.pattern };
  }
}

// Backtracking Animator (N-Queens)
class BacktrackingAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    this.n = 4;
    this.board = Array(this.n).fill(-1);
    this.solutions = [];

    this.setCode(`def solve_n_queens(n):
    board = [-1] * n  # board[row] = col
    solutions = []

    def is_safe(row, col):
        for r in range(row):
            c = board[r]
            if c == col or abs(c - col) == row - r:
                return False
        return True

    def backtrack(row):
        if row == n:
            solutions.append(board[:])
            return
        for col in range(n):
            if is_safe(row, col):
                board[row] = col
                backtrack(row + 1)
                board[row] = -1  # backtrack

    backtrack(0)
    return solutions`);
  }

  buildSteps() {
    this.steps = [];
    this.board = Array(this.n).fill(-1);
    this.solutions = [];

    this.steps.push({
      lineNum: 1,
      state: `Initialize ${this.n}x${this.n} board for N-Queens`,
      board: [...this.board],
      row: -1,
      col: -1,
      checking: false,
      solutions: [],
      apply: () => {}
    });

    this.backtrackBuild(0);

    this.steps.push({
      lineNum: 20,
      state: `Found ${this.solutions.length} solution(s)!`,
      board: Array(this.n).fill(-1),
      row: -1,
      col: -1,
      checking: false,
      solutions: [...this.solutions],
      apply: () => {}
    });
  }

  backtrackBuild(row) {
    if (row === this.n) {
      this.solutions.push([...this.board]);
      this.steps.push({
        lineNum: 12,
        state: `Solution found! Queens at columns: [${this.board.join(', ')}]`,
        board: [...this.board],
        row: row,
        col: -1,
        checking: false,
        foundSolution: true,
        solutions: [...this.solutions],
        apply: () => {}
      });
      return;
    }

    this.steps.push({
      lineNum: 14,
      state: `Trying row ${row}`,
      board: [...this.board],
      row: row,
      col: -1,
      checking: false,
      solutions: [...this.solutions],
      apply: () => {}
    });

    for (let col = 0; col < this.n; col++) {
      this.steps.push({
        lineNum: 15,
        state: `Checking if position (${row}, ${col}) is safe`,
        board: [...this.board],
        row: row,
        col: col,
        checking: true,
        solutions: [...this.solutions],
        apply: () => {}
      });

      if (this.isSafe(row, col)) {
        this.board[row] = col;
        this.steps.push({
          lineNum: 16,
          state: `Safe! Placing queen at (${row}, ${col})`,
          board: [...this.board],
          row: row,
          col: col,
          checking: false,
          placed: true,
          solutions: [...this.solutions],
          apply: () => {}
        });

        this.backtrackBuild(row + 1);

        this.board[row] = -1;
        this.steps.push({
          lineNum: 18,
          state: `Backtracking: removing queen from (${row}, ${col})`,
          board: [...this.board],
          row: row,
          col: col,
          checking: false,
          backtracking: true,
          solutions: [...this.solutions],
          apply: () => {}
        });
      } else {
        this.steps.push({
          lineNum: 15,
          state: `Not safe at (${row}, ${col}) - conflict detected`,
          board: [...this.board],
          row: row,
          col: col,
          checking: false,
          conflict: true,
          solutions: [...this.solutions],
          apply: () => {}
        });
      }
    }
  }

  isSafe(row, col) {
    for (let r = 0; r < row; r++) {
      const c = this.board[r];
      if (c === col || Math.abs(c - col) === row - r) {
        return false;
      }
    }
    return true;
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { board, row, col, checking, placed, conflict, backtracking, foundSolution, solutions } = step;

    let html = '<div class="nqueens-viz">';
    html += '<div class="chess-board">';

    for (let r = 0; r < this.n; r++) {
      html += '<div class="chess-row">';
      for (let c = 0; c < this.n; c++) {
        let cellClass = 'chess-cell';
        cellClass += (r + c) % 2 === 0 ? ' light' : ' dark';

        if (r === row && c === col) {
          if (checking) cellClass += ' checking';
          if (conflict) cellClass += ' conflict';
          if (placed) cellClass += ' placed';
          if (backtracking) cellClass += ' backtracking';
        }
        if (foundSolution) cellClass += ' solution';

        const hasQueen = board && board[r] === c;
        html += `<div class="${cellClass}">${hasQueen ? 'Q' : ''}</div>`;
      }
      html += '</div>';
    }

    html += '</div>';

    if (solutions && solutions.length > 0) {
      html += `<div class="solutions-count">Solutions found: ${solutions.length}</div>`;
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'current_row': step.row >= 0 ? step.row : '-',
      'trying_col': step.col >= 0 ? step.col : '-',
      'board': (step.board || []).map((c, r) => c >= 0 ? `(${r},${c})` : '-').filter(x => x !== '-').join(', ') || 'empty',
      'solutions': (step.solutions || []).length
    };
  }

  getInputData() {
    return { 'Board Size': `${this.n}x${this.n}`, 'Problem': 'N-Queens' };
  }
}

// Permutations Animator
class PermutationsAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    this.nums = [1, 2, 3];
    this.permutations = [];

    this.setCode(`def permute(nums):
    result = []

    def backtrack(path, remaining):
        if not remaining:
            result.append(path[:])
            return
        for i, num in enumerate(remaining):
            path.append(num)
            new_remaining = remaining[:i] + remaining[i+1:]
            backtrack(path, new_remaining)
            path.pop()  # backtrack

    backtrack([], nums)
    return result`);
  }

  buildSteps() {
    this.steps = [];
    this.permutations = [];

    this.steps.push({
      lineNum: 1,
      state: `Generate permutations of [${this.nums.join(', ')}]`,
      path: [],
      remaining: [...this.nums],
      permutations: [],
      apply: () => {}
    });

    this.permuteBuild([], [...this.nums]);

    this.steps.push({
      lineNum: 13,
      state: `Complete! Generated ${this.permutations.length} permutations`,
      path: [],
      remaining: [],
      permutations: [...this.permutations],
      apply: () => {}
    });
  }

  permuteBuild(path, remaining) {
    if (remaining.length === 0) {
      this.permutations.push([...path]);
      this.steps.push({
        lineNum: 5,
        state: `Found permutation: [${path.join(', ')}]`,
        path: [...path],
        remaining: [],
        foundPermutation: true,
        permutations: [...this.permutations],
        apply: () => {}
      });
      return;
    }

    for (let i = 0; i < remaining.length; i++) {
      const num = remaining[i];

      this.steps.push({
        lineNum: 7,
        state: `Choose ${num} from remaining [${remaining.join(', ')}]`,
        path: [...path],
        remaining: [...remaining],
        choosing: i,
        permutations: [...this.permutations],
        apply: () => {}
      });

      path.push(num);
      const newRemaining = [...remaining.slice(0, i), ...remaining.slice(i + 1)];

      this.steps.push({
        lineNum: 8,
        state: `Path: [${path.join(', ')}], Remaining: [${newRemaining.join(', ')}]`,
        path: [...path],
        remaining: [...newRemaining],
        permutations: [...this.permutations],
        apply: () => {}
      });

      this.permuteBuild(path, newRemaining);

      path.pop();
      this.steps.push({
        lineNum: 10,
        state: `Backtrack: removed ${num}, path now [${path.join(', ')}]`,
        path: [...path],
        remaining: [...remaining],
        backtracking: true,
        permutations: [...this.permutations],
        apply: () => {}
      });
    }
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { path, remaining, choosing, foundPermutation, backtracking, permutations } = step;

    let html = '<div class="permutations-viz">';

    // Current state
    html += '<div class="perm-state">';
    html += '<div class="perm-row"><span class="perm-label">Path:</span><div class="perm-array">';
    (path || []).forEach((n, i) => {
      html += `<span class="perm-item${foundPermutation ? ' found' : ''}">${n}</span>`;
    });
    if ((path || []).length === 0) html += '<span class="perm-empty">empty</span>';
    html += '</div></div>';

    html += '<div class="perm-row"><span class="perm-label">Remaining:</span><div class="perm-array">';
    (remaining || []).forEach((n, i) => {
      let itemClass = 'perm-item';
      if (choosing === i) itemClass += ' choosing';
      html += `<span class="${itemClass}">${n}</span>`;
    });
    if ((remaining || []).length === 0) html += '<span class="perm-empty">empty</span>';
    html += '</div></div>';
    html += '</div>';

    // Generated permutations
    if (permutations && permutations.length > 0) {
      html += '<div class="perm-results">';
      html += `<div class="perm-header">Permutations (${permutations.length}):</div>`;
      html += '<div class="perm-list">';
      permutations.forEach((p, i) => {
        html += `<span class="perm-result">[${p.join(',')}]</span>`;
      });
      html += '</div></div>';
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'path': `[${(step.path || []).join(', ')}]`,
      'remaining': `[${(step.remaining || []).join(', ')}]`,
      'permutations_found': (step.permutations || []).length
    };
  }

  getInputData() {
    return { 'Input': `[${this.nums.join(', ')}]`, 'Total Permutations': `${this.factorial(this.nums.length)}` };
  }

  factorial(n) {
    return n <= 1 ? 1 : n * this.factorial(n - 1);
  }
}

// Iterative DP Animator (Fibonacci bottom-up)
class IterativeDPAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    this.n = 8;
    this.dp = [];

    this.setCode(`def fib_iterative(n):
    if n <= 1:
        return n

    dp = [0] * (n + 1)
    dp[0] = 0
    dp[1] = 1

    for i in range(2, n + 1):
        dp[i] = dp[i-1] + dp[i-2]

    return dp[n]`);
  }

  buildSteps() {
    this.steps = [];
    this.dp = new Array(this.n + 1).fill(null);

    this.steps.push({
      lineNum: 1,
      state: `Calculate Fibonacci(${this.n}) using bottom-up DP`,
      dp: [...this.dp],
      currentIdx: -1,
      apply: () => {}
    });

    this.dp[0] = 0;
    this.steps.push({
      lineNum: 6,
      state: 'Base case: dp[0] = 0',
      dp: [...this.dp],
      currentIdx: 0,
      isBase: true,
      apply: () => {}
    });

    this.dp[1] = 1;
    this.steps.push({
      lineNum: 7,
      state: 'Base case: dp[1] = 1',
      dp: [...this.dp],
      currentIdx: 1,
      isBase: true,
      apply: () => {}
    });

    for (let i = 2; i <= this.n; i++) {
      this.steps.push({
        lineNum: 9,
        state: `Computing dp[${i}] = dp[${i-1}] + dp[${i-2}] = ${this.dp[i-1]} + ${this.dp[i-2]}`,
        dp: [...this.dp],
        currentIdx: i,
        deps: [i-1, i-2],
        computing: true,
        apply: () => {}
      });

      this.dp[i] = this.dp[i-1] + this.dp[i-2];
      this.steps.push({
        lineNum: 9,
        state: `dp[${i}] = ${this.dp[i]}`,
        dp: [...this.dp],
        currentIdx: i,
        computed: true,
        apply: () => {}
      });
    }

    this.steps.push({
      lineNum: 11,
      state: `Result: Fibonacci(${this.n}) = ${this.dp[this.n]}`,
      dp: [...this.dp],
      currentIdx: this.n,
      final: true,
      apply: () => {}
    });
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { dp, currentIdx, deps, isBase, computing, computed, final } = step;

    let html = '<div class="iterative-dp-viz">';

    // DP table
    html += '<div class="dp-table">';
    html += '<div class="dp-indices">';
    for (let i = 0; i <= this.n; i++) {
      html += `<span class="dp-idx">i=${i}</span>`;
    }
    html += '</div>';

    html += '<div class="dp-values">';
    for (let i = 0; i <= this.n; i++) {
      let cellClass = 'dp-cell';
      if (i === currentIdx) {
        if (isBase) cellClass += ' base';
        if (computing) cellClass += ' computing';
        if (computed) cellClass += ' computed';
        if (final) cellClass += ' final';
      }
      if (deps && deps.includes(i)) cellClass += ' dependency';

      const val = dp && dp[i] !== null ? dp[i] : '?';
      html += `<span class="${cellClass}">${val}</span>`;
    }
    html += '</div>';
    html += '</div>';

    // Formula visualization
    if (computing && deps) {
      html += '<div class="dp-formula">';
      html += `dp[${currentIdx}] = dp[${deps[0]}] + dp[${deps[1]}]`;
      html += `<br>= ${dp[deps[0]]} + ${dp[deps[1]]}`;
      html += '</div>';
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'current_index': step.currentIdx >= 0 ? step.currentIdx : '-',
      'dp_value': step.dp && step.currentIdx >= 0 && step.dp[step.currentIdx] !== null ? step.dp[step.currentIdx] : '-',
      'n': this.n
    };
  }

  getInputData() {
    return { 'Calculate': `Fibonacci(${this.n})`, 'Method': 'Bottom-up DP' };
  }
}

// Rotated Array Search Animator
class RotatedArrayAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    this.nums = [4, 5, 6, 7, 0, 1, 2];
    this.target = 0;
    this.low = 0;
    this.high = this.nums.length - 1;

    this.setCode(`def search_rotated(nums, target):
    low, high = 0, len(nums) - 1

    while low <= high:
        mid = (low + high) // 2

        if nums[mid] == target:
            return mid

        # Left half is sorted
        if nums[low] <= nums[mid]:
            if nums[low] <= target < nums[mid]:
                high = mid - 1
            else:
                low = mid + 1
        # Right half is sorted
        else:
            if nums[mid] < target <= nums[high]:
                low = mid + 1
            else:
                high = mid - 1

    return -1`);
  }

  buildSteps() {
    this.steps = [];
    let low = 0;
    let high = this.nums.length - 1;

    this.steps.push({
      lineNum: 1,
      state: `Search for ${this.target} in rotated array [${this.nums.join(', ')}]`,
      array: [...this.nums],
      low: low,
      high: high,
      mid: -1,
      target: this.target,
      apply: () => {}
    });

    while (low <= high) {
      const mid = Math.floor((low + high) / 2);

      this.steps.push({
        lineNum: 5,
        state: `Calculate mid = (${low} + ${high}) / 2 = ${mid}`,
        array: [...this.nums],
        low: low,
        high: high,
        mid: mid,
        target: this.target,
        apply: () => {}
      });

      if (this.nums[mid] === this.target) {
        this.steps.push({
          lineNum: 7,
          state: `Found! nums[${mid}] = ${this.target}`,
          array: [...this.nums],
          low: low,
          high: high,
          mid: mid,
          target: this.target,
          found: true,
          apply: () => {}
        });
        return;
      }

      // Left half is sorted
      if (this.nums[low] <= this.nums[mid]) {
        this.steps.push({
          lineNum: 10,
          state: `Left half [${low}..${mid}] is sorted (${this.nums[low]} <= ${this.nums[mid]})`,
          array: [...this.nums],
          low: low,
          high: high,
          mid: mid,
          target: this.target,
          sortedHalf: 'left',
          apply: () => {}
        });

        if (this.nums[low] <= this.target && this.target < this.nums[mid]) {
          this.steps.push({
            lineNum: 12,
            state: `Target ${this.target} is in left half, search left`,
            array: [...this.nums],
            low: low,
            high: mid - 1,
            mid: mid,
            target: this.target,
            searchDirection: 'left',
            apply: () => {}
          });
          high = mid - 1;
        } else {
          this.steps.push({
            lineNum: 14,
            state: `Target ${this.target} is in right half, search right`,
            array: [...this.nums],
            low: mid + 1,
            high: high,
            mid: mid,
            target: this.target,
            searchDirection: 'right',
            apply: () => {}
          });
          low = mid + 1;
        }
      } else {
        this.steps.push({
          lineNum: 16,
          state: `Right half [${mid}..${high}] is sorted (${this.nums[mid]} > ${this.nums[low]})`,
          array: [...this.nums],
          low: low,
          high: high,
          mid: mid,
          target: this.target,
          sortedHalf: 'right',
          apply: () => {}
        });

        if (this.nums[mid] < this.target && this.target <= this.nums[high]) {
          this.steps.push({
            lineNum: 18,
            state: `Target ${this.target} is in right half, search right`,
            array: [...this.nums],
            low: mid + 1,
            high: high,
            mid: mid,
            target: this.target,
            searchDirection: 'right',
            apply: () => {}
          });
          low = mid + 1;
        } else {
          this.steps.push({
            lineNum: 20,
            state: `Target ${this.target} is in left half, search left`,
            array: [...this.nums],
            low: low,
            high: mid - 1,
            mid: mid,
            target: this.target,
            searchDirection: 'left',
            apply: () => {}
          });
          high = mid - 1;
        }
      }
    }

    this.steps.push({
      lineNum: 22,
      state: `Target ${this.target} not found`,
      array: [...this.nums],
      low: low,
      high: high,
      mid: -1,
      target: this.target,
      notFound: true,
      apply: () => {}
    });
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { array, low, high, mid, target, found, sortedHalf, searchDirection, notFound } = step;

    let html = '<div class="rotated-array-viz">';

    // Array visualization
    html += '<div class="array-viz rotated">';
    (array || this.nums).forEach((val, i) => {
      let cellClass = 'array-cell';
      if (i < low || i > high) cellClass += ' excluded';
      if (i === mid) cellClass += found ? ' found' : ' mid';
      if (i === low) cellClass += ' low';
      if (i === high) cellClass += ' high';
      if (val === target) cellClass += ' target';
      if (sortedHalf === 'left' && i >= low && i <= mid) cellClass += ' sorted-half';
      if (sortedHalf === 'right' && i >= mid && i <= high) cellClass += ' sorted-half';

      html += `<div class="${cellClass}">
        <span class="cell-value">${val}</span>
        <span class="cell-index">${i}</span>
      </div>`;
    });
    html += '</div>';

    // Pointer labels
    html += '<div class="pointer-labels">';
    if (low >= 0) html += `<span class="ptr-label">low=${low}</span>`;
    if (mid >= 0) html += `<span class="ptr-label">mid=${mid}</span>`;
    if (high >= 0) html += `<span class="ptr-label">high=${high}</span>`;
    html += '</div>';

    // Pivot indicator
    const pivotIdx = array ? array.findIndex((v, i) => i > 0 && v < array[i-1]) : -1;
    if (pivotIdx > 0) {
      html += `<div class="pivot-info">Rotation pivot at index ${pivotIdx}</div>`;
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'low': step.low >= 0 ? step.low : '-',
      'mid': step.mid >= 0 ? step.mid : '-',
      'high': step.high >= 0 ? step.high : '-',
      'target': this.target
    };
  }

  getInputData() {
    return { 'Array': `[${this.nums.join(', ')}]`, 'Target': this.target };
  }
}

// Merge Sort Animator
class MergeSortAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    this.original = [38, 27, 43, 3, 9, 82, 10];
    this.array = [...this.original];

    this.setCode(`def merge_sort(arr):
    if len(arr) <= 1:
        return arr

    mid = len(arr) // 2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])

    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result`);
  }

  buildSteps() {
    this.steps = [];

    this.steps.push({
      lineNum: 1,
      state: `Merge sort array [${this.original.join(', ')}]`,
      levels: [[...this.original]],
      currentLevel: 0,
      phase: 'divide',
      apply: () => {}
    });

    // Build divide phase
    this.buildDivide([[...this.original]], 0);

    // Build merge phase
    this.buildMerge();
  }

  buildDivide(levels, level) {
    const current = levels[level];
    if (!current || current.every(arr => arr.length <= 1)) {
      return levels;
    }

    const nextLevel = [];
    current.forEach((arr, idx) => {
      if (arr.length <= 1) {
        nextLevel.push(arr);
      } else {
        const mid = Math.floor(arr.length / 2);
        const left = arr.slice(0, mid);
        const right = arr.slice(mid);

        this.steps.push({
          lineNum: 5,
          state: `Split [${arr.join(', ')}] into [${left.join(', ')}] and [${right.join(', ')}]`,
          levels: levels.map(l => [...l]),
          currentLevel: level,
          splitting: idx,
          left: left,
          right: right,
          phase: 'divide',
          apply: () => {}
        });

        nextLevel.push(left, right);
      }
    });

    levels.push(nextLevel);
    this.steps.push({
      lineNum: 6,
      state: `Level ${level + 1}: ${nextLevel.map(a => `[${a.join(',')}]`).join(' ')}`,
      levels: levels.map(l => [...l]),
      currentLevel: level + 1,
      phase: 'divide',
      apply: () => {}
    });

    return this.buildDivide(levels, level + 1);
  }

  buildMerge() {
    // Simulate the merge phase
    let levels = [
      [[38, 27, 43, 3, 9, 82, 10]],
      [[38, 27, 43], [3, 9, 82, 10]],
      [[38], [27, 43], [3, 9], [82, 10]],
      [[38], [27], [43], [3], [9], [82], [10]]
    ];

    // Merge level 3 -> 2
    this.steps.push({
      lineNum: 11,
      state: 'Begin merging phase',
      levels: levels,
      currentLevel: 3,
      phase: 'merge',
      apply: () => {}
    });

    // Merge pairs
    const mergeOps = [
      { left: [27], right: [43], result: [27, 43], state: 'Merge [27] and [43]' },
      { left: [3], right: [9], result: [3, 9], state: 'Merge [3] and [9]' },
      { left: [82], right: [10], result: [10, 82], state: 'Merge [82] and [10]' },
      { left: [38], right: [27, 43], result: [27, 38, 43], state: 'Merge [38] and [27,43]' },
      { left: [3, 9], right: [10, 82], result: [3, 9, 10, 82], state: 'Merge [3,9] and [10,82]' },
      { left: [27, 38, 43], right: [3, 9, 10, 82], result: [3, 9, 10, 27, 38, 43, 82], state: 'Final merge' }
    ];

    mergeOps.forEach((op, i) => {
      this.steps.push({
        lineNum: 14,
        state: op.state,
        merging: { left: op.left, right: op.right },
        result: op.result,
        phase: 'merge',
        apply: () => {}
      });

      this.steps.push({
        lineNum: 22,
        state: `Result: [${op.result.join(', ')}]`,
        merged: op.result,
        phase: 'merge',
        apply: () => {}
      });
    });

    this.steps.push({
      lineNum: 23,
      state: `Sorted: [${[3, 9, 10, 27, 38, 43, 82].join(', ')}]`,
      sorted: [3, 9, 10, 27, 38, 43, 82],
      phase: 'complete',
      apply: () => {}
    });
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { levels, phase, merging, result, merged, sorted, splitting, left, right } = step;

    let html = '<div class="merge-sort-viz">';

    if (phase === 'divide' && levels) {
      html += '<div class="divide-phase">';
      levels.forEach((level, i) => {
        html += `<div class="merge-level level-${i}">`;
        level.forEach((arr, j) => {
          let arrClass = 'merge-array';
          if (splitting === j && i === step.currentLevel) arrClass += ' splitting';
          html += `<span class="${arrClass}">[${arr.join(', ')}]</span>`;
        });
        html += '</div>';
      });
      html += '</div>';
    }

    if (phase === 'merge' && merging) {
      html += '<div class="merge-phase">';
      html += `<div class="merge-inputs">`;
      html += `<span class="merge-array left">[${merging.left.join(', ')}]</span>`;
      html += `<span class="merge-op">+</span>`;
      html += `<span class="merge-array right">[${merging.right.join(', ')}]</span>`;
      html += '</div>';
      if (result) {
        html += `<div class="merge-result">= [${result.join(', ')}]</div>`;
      }
      html += '</div>';
    }

    if (merged) {
      html += `<div class="merged-result">[${merged.join(', ')}]</div>`;
    }

    if (sorted) {
      html += `<div class="final-sorted">Sorted: [${sorted.join(', ')}]</div>`;
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'phase': step.phase || '-',
      'current_level': step.currentLevel >= 0 ? step.currentLevel : '-'
    };
  }

  getInputData() {
    return { 'Input': `[${this.original.join(', ')}]`, 'Algorithm': 'Merge Sort O(n log n)' };
  }
}

// Sliding Window Animator
class SlidingWindowAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    this.nums = [2, 1, 5, 1, 3, 2];
    this.k = 3;
    this.maxSum = 0;

    this.setCode(`def max_sum_subarray(nums, k):
    if len(nums) < k:
        return 0

    # Initial window sum
    window_sum = sum(nums[:k])
    max_sum = window_sum

    # Slide the window
    for i in range(k, len(nums)):
        window_sum += nums[i] - nums[i - k]
        max_sum = max(max_sum, window_sum)

    return max_sum`);
  }

  buildSteps() {
    this.steps = [];

    this.steps.push({
      lineNum: 1,
      state: `Find max sum of ${this.k} consecutive elements in [${this.nums.join(', ')}]`,
      windowStart: -1,
      windowEnd: -1,
      windowSum: 0,
      maxSum: 0,
      apply: () => {}
    });

    // Initial window
    let windowSum = 0;
    for (let i = 0; i < this.k; i++) {
      windowSum += this.nums[i];
    }

    this.steps.push({
      lineNum: 5,
      state: `Initial window [0..${this.k-1}]: sum = ${windowSum}`,
      windowStart: 0,
      windowEnd: this.k - 1,
      windowSum: windowSum,
      maxSum: windowSum,
      isInitial: true,
      apply: () => {}
    });

    let maxSum = windowSum;

    // Slide the window
    for (let i = this.k; i < this.nums.length; i++) {
      const removing = this.nums[i - this.k];
      const adding = this.nums[i];

      this.steps.push({
        lineNum: 10,
        state: `Slide: remove nums[${i - this.k}]=${removing}, add nums[${i}]=${adding}`,
        windowStart: i - this.k,
        windowEnd: i - 1,
        windowSum: windowSum,
        maxSum: maxSum,
        removing: i - this.k,
        adding: i,
        apply: () => {}
      });

      windowSum = windowSum + adding - removing;
      maxSum = Math.max(maxSum, windowSum);

      this.steps.push({
        lineNum: 11,
        state: `New window [${i - this.k + 1}..${i}]: sum = ${windowSum}, max = ${maxSum}`,
        windowStart: i - this.k + 1,
        windowEnd: i,
        windowSum: windowSum,
        maxSum: maxSum,
        newMax: windowSum === maxSum,
        apply: () => {}
      });
    }

    this.steps.push({
      lineNum: 13,
      state: `Maximum sum of ${this.k} consecutive elements: ${maxSum}`,
      windowStart: -1,
      windowEnd: -1,
      windowSum: windowSum,
      maxSum: maxSum,
      final: true,
      apply: () => {}
    });
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { windowStart, windowEnd, windowSum, maxSum, removing, adding, isInitial, newMax, final } = step;

    let html = '<div class="sliding-window-viz">';

    // Array with window
    html += '<div class="array-viz window-array">';
    this.nums.forEach((val, i) => {
      let cellClass = 'array-cell';
      if (i >= windowStart && i <= windowEnd) cellClass += ' in-window';
      if (i === removing) cellClass += ' removing';
      if (i === adding) cellClass += ' adding';

      html += `<div class="${cellClass}">
        <span class="cell-value">${val}</span>
        <span class="cell-index">${i}</span>
      </div>`;
    });
    html += '</div>';

    // Window info
    html += '<div class="window-info">';
    if (windowStart >= 0) {
      html += `<div class="window-range">Window: [${windowStart}..${windowEnd}]</div>`;
    }
    html += `<div class="window-sum${newMax ? ' new-max' : ''}">Window Sum: ${windowSum}</div>`;
    html += `<div class="max-sum">Max Sum: ${maxSum}</div>`;
    html += '</div>';

    if (final) {
      html += `<div class="final-result">Answer: ${maxSum}</div>`;
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'window_start': step.windowStart >= 0 ? step.windowStart : '-',
      'window_end': step.windowEnd >= 0 ? step.windowEnd : '-',
      'window_sum': step.windowSum,
      'max_sum': step.maxSum
    };
  }

  getInputData() {
    return { 'Array': `[${this.nums.join(', ')}]`, 'Window Size (k)': this.k };
  }
}

// Binary Tree Traversal Animator
class BinaryTreeAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    // Tree structure: { val, left, right }
    this.root = {
      val: 1,
      left: {
        val: 2,
        left: { val: 4, left: null, right: null },
        right: { val: 5, left: null, right: null }
      },
      right: {
        val: 3,
        left: { val: 6, left: null, right: null },
        right: { val: 7, left: null, right: null }
      }
    };

    this.traversalType = 'inorder';
    this.result = [];

    this.setCode(`def inorder(node):
    if node is None:
        return
    inorder(node.left)   # Visit left
    print(node.val)      # Process node
    inorder(node.right)  # Visit right

# Preorder: node -> left -> right
# Postorder: left -> right -> node`);
  }

  buildSteps() {
    this.steps = [];
    this.result = [];

    this.steps.push({
      lineNum: 1,
      state: `Inorder traversal: Left -> Node -> Right`,
      current: null,
      visited: [],
      result: [],
      apply: () => {}
    });

    this.inorderBuild(this.root, []);

    this.steps.push({
      lineNum: 6,
      state: `Traversal complete: [${this.result.join(', ')}]`,
      current: null,
      visited: [...this.result],
      result: [...this.result],
      complete: true,
      apply: () => {}
    });
  }

  inorderBuild(node, visited) {
    if (!node) return;

    this.steps.push({
      lineNum: 4,
      state: `Go left from ${node.val}`,
      current: node.val,
      visited: [...visited],
      result: [...this.result],
      direction: 'left',
      apply: () => {}
    });

    this.inorderBuild(node.left, visited);

    visited.push(node.val);
    this.result.push(node.val);

    this.steps.push({
      lineNum: 5,
      state: `Visit node ${node.val}`,
      current: node.val,
      visited: [...visited],
      result: [...this.result],
      processing: true,
      apply: () => {}
    });

    this.steps.push({
      lineNum: 6,
      state: `Go right from ${node.val}`,
      current: node.val,
      visited: [...visited],
      result: [...this.result],
      direction: 'right',
      apply: () => {}
    });

    this.inorderBuild(node.right, visited);
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { current, visited, result, processing, direction, complete } = step;

    let html = '<div class="binary-tree-viz">';

    // Tree visualization using SVG
    html += '<svg class="tree-svg" viewBox="0 0 300 200">';

    // Draw edges
    html += '<g class="tree-edges">';
    // Level 1 -> 2
    html += '<line x1="150" y1="30" x2="75" y2="80" class="tree-edge"/>';
    html += '<line x1="150" y1="30" x2="225" y2="80" class="tree-edge"/>';
    // Level 2 -> 3
    html += '<line x1="75" y1="80" x2="40" y2="140" class="tree-edge"/>';
    html += '<line x1="75" y1="80" x2="110" y2="140" class="tree-edge"/>';
    html += '<line x1="225" y1="80" x2="190" y2="140" class="tree-edge"/>';
    html += '<line x1="225" y1="80" x2="260" y2="140" class="tree-edge"/>';
    html += '</g>';

    // Draw nodes
    const positions = [
      { val: 1, x: 150, y: 30 },
      { val: 2, x: 75, y: 80 },
      { val: 3, x: 225, y: 80 },
      { val: 4, x: 40, y: 140 },
      { val: 5, x: 110, y: 140 },
      { val: 6, x: 190, y: 140 },
      { val: 7, x: 260, y: 140 }
    ];

    html += '<g class="tree-nodes">';
    positions.forEach(pos => {
      let nodeClass = 'tree-node';
      if (visited && visited.includes(pos.val)) nodeClass += ' visited';
      if (pos.val === current) {
        nodeClass += ' current';
        if (processing) nodeClass += ' processing';
      }

      html += `<circle cx="${pos.x}" cy="${pos.y}" r="18" class="${nodeClass}"/>`;
      html += `<text x="${pos.x}" y="${pos.y + 5}" class="node-label">${pos.val}</text>`;
    });
    html += '</g>';

    html += '</svg>';

    // Result display
    html += '<div class="traversal-result">';
    html += `<span class="result-label">Inorder:</span>`;
    html += `<span class="result-values">[${(result || []).join(', ')}]</span>`;
    html += '</div>';

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'current_node': step.current || '-',
      'result': `[${(step.result || []).join(', ')}]`
    };
  }

  getInputData() {
    return { 'Tree': 'Complete binary tree', 'Traversal': 'Inorder (Left-Node-Right)' };
  }
}

// Heap/Priority Queue Animator
class HeapAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    this.heap = [];
    this.insertions = [4, 10, 3, 5, 1, 8];

    this.setCode(`class MinHeap:
    def __init__(self):
        self.heap = []

    def insert(self, val):
        self.heap.append(val)
        self._bubble_up(len(self.heap) - 1)

    def _bubble_up(self, idx):
        parent = (idx - 1) // 2
        if idx > 0 and self.heap[idx] < self.heap[parent]:
            self.heap[idx], self.heap[parent] = \\
                self.heap[parent], self.heap[idx]
            self._bubble_up(parent)`);
  }

  buildSteps() {
    this.steps = [];
    this.heap = [];

    this.steps.push({
      lineNum: 1,
      state: `Build min-heap by inserting [${this.insertions.join(', ')}]`,
      heap: [],
      apply: () => {}
    });

    this.insertions.forEach((val, i) => {
      this.steps.push({
        lineNum: 5,
        state: `Insert ${val}`,
        heap: [...this.heap],
        inserting: val,
        apply: () => {}
      });

      this.heap.push(val);
      let idx = this.heap.length - 1;

      this.steps.push({
        lineNum: 6,
        state: `Added ${val} at index ${idx}`,
        heap: [...this.heap],
        currentIdx: idx,
        apply: () => {}
      });

      // Bubble up
      while (idx > 0) {
        const parentIdx = Math.floor((idx - 1) / 2);

        if (this.heap[idx] < this.heap[parentIdx]) {
          this.steps.push({
            lineNum: 10,
            state: `Bubble up: ${this.heap[idx]} < ${this.heap[parentIdx]}, swap`,
            heap: [...this.heap],
            currentIdx: idx,
            parentIdx: parentIdx,
            swapping: true,
            apply: () => {}
          });

          [this.heap[idx], this.heap[parentIdx]] = [this.heap[parentIdx], this.heap[idx]];

          this.steps.push({
            lineNum: 12,
            state: `After swap: heap = [${this.heap.join(', ')}]`,
            heap: [...this.heap],
            currentIdx: parentIdx,
            apply: () => {}
          });

          idx = parentIdx;
        } else {
          this.steps.push({
            lineNum: 9,
            state: `Heap property satisfied: ${this.heap[idx]} >= ${this.heap[parentIdx]}`,
            heap: [...this.heap],
            currentIdx: idx,
            parentIdx: parentIdx,
            satisfied: true,
            apply: () => {}
          });
          break;
        }
      }

      if (idx === 0 && this.heap.length > 1) {
        this.steps.push({
          lineNum: 9,
          state: `${val} reached root position`,
          heap: [...this.heap],
          currentIdx: 0,
          isRoot: true,
          apply: () => {}
        });
      }
    });

    this.steps.push({
      lineNum: 13,
      state: `Min-heap built: [${this.heap.join(', ')}], min = ${this.heap[0]}`,
      heap: [...this.heap],
      complete: true,
      apply: () => {}
    });
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { heap, currentIdx, parentIdx, swapping, inserting, complete } = step;

    let html = '<div class="heap-viz">';

    // Tree visualization
    if (heap && heap.length > 0) {
      html += '<svg class="heap-tree" viewBox="0 0 320 180">';

      const positions = this.getHeapPositions(heap.length);

      // Draw edges
      html += '<g class="heap-edges">';
      for (let i = 1; i < heap.length; i++) {
        const parentPos = positions[Math.floor((i - 1) / 2)];
        const childPos = positions[i];
        let edgeClass = 'heap-edge';
        if ((i === currentIdx || i === parentIdx) && swapping) edgeClass += ' swapping';
        html += `<line x1="${parentPos.x}" y1="${parentPos.y}" x2="${childPos.x}" y2="${childPos.y}" class="${edgeClass}"/>`;
      }
      html += '</g>';

      // Draw nodes
      html += '<g class="heap-nodes">';
      heap.forEach((val, i) => {
        const pos = positions[i];
        let nodeClass = 'heap-node';
        if (i === currentIdx) nodeClass += ' current';
        if (i === parentIdx) nodeClass += ' parent';
        if (i === 0 && complete) nodeClass += ' min';

        html += `<circle cx="${pos.x}" cy="${pos.y}" r="18" class="${nodeClass}"/>`;
        html += `<text x="${pos.x}" y="${pos.y + 5}" class="node-label">${val}</text>`;
        html += `<text x="${pos.x}" y="${pos.y - 25}" class="idx-label">[${i}]</text>`;
      });
      html += '</g>';

      html += '</svg>';
    }

    // Array representation
    html += '<div class="heap-array">';
    html += `<span class="array-label">Array:</span>`;
    html += `<span class="array-values">[${(heap || []).join(', ')}]</span>`;
    html += '</div>';

    if (inserting !== undefined) {
      html += `<div class="inserting">Inserting: ${inserting}</div>`;
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getHeapPositions(size) {
    const positions = [];
    const levels = Math.ceil(Math.log2(size + 1));
    let idx = 0;

    for (let level = 0; level < levels && idx < size; level++) {
      const nodesInLevel = Math.pow(2, level);
      const levelWidth = 300;
      const spacing = levelWidth / (nodesInLevel + 1);

      for (let i = 0; i < nodesInLevel && idx < size; i++) {
        positions.push({
          x: 10 + spacing * (i + 1),
          y: 25 + level * 45
        });
        idx++;
      }
    }

    return positions;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'heap_size': (step.heap || []).length,
      'current_idx': step.currentIdx >= 0 ? step.currentIdx : '-',
      'min_value': step.heap && step.heap.length > 0 ? step.heap[0] : '-'
    };
  }

  getInputData() {
    return { 'Insertions': `[${this.insertions.join(', ')}]`, 'Type': 'Min-Heap' };
  }
}

// Greedy Algorithm Animator (Activity Selection / Interval Scheduling)
class GreedyAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    // Activities: [start, end, name]
    this.activities = [
      [1, 4, 'A'],
      [3, 5, 'B'],
      [0, 6, 'C'],
      [5, 7, 'D'],
      [3, 9, 'E'],
      [5, 9, 'F'],
      [6, 10, 'G'],
      [8, 11, 'H'],
      [8, 12, 'I'],
      [2, 14, 'J'],
      [12, 16, 'K']
    ];
    this.selected = [];
    this.currentEnd = 0;

    this.setCode(`def activity_selection(activities):
    # Sort by finish time (greedy choice)
    sorted_acts = sorted(activities, key=lambda x: x[1])

    selected = []
    current_end = 0

    for start, end, name in sorted_acts:
        if start >= current_end:
            # Activity is compatible
            selected.append((start, end, name))
            current_end = end

    return selected`);
  }

  buildSteps() {
    this.steps = [];
    this.selected = [];
    this.currentEnd = 0;

    // Sort activities by end time
    const sorted = [...this.activities].sort((a, b) => a[1] - b[1]);

    this.steps.push({
      lineNum: 1,
      state: `Activity Selection: maximize non-overlapping activities`,
      activities: this.activities.map(a => ({ start: a[0], end: a[1], name: a[2] })),
      sorted: null,
      selected: [],
      currentEnd: 0,
      considering: -1,
      apply: () => {}
    });

    this.steps.push({
      lineNum: 2,
      state: `Sort activities by finish time (greedy choice property)`,
      activities: this.activities.map(a => ({ start: a[0], end: a[1], name: a[2] })),
      sorted: sorted.map(a => ({ start: a[0], end: a[1], name: a[2] })),
      selected: [],
      currentEnd: 0,
      considering: -1,
      sorting: true,
      apply: () => {}
    });

    let currentEnd = 0;
    const selectedList = [];

    sorted.forEach((activity, idx) => {
      const [start, end, name] = activity;

      this.steps.push({
        lineNum: 7,
        state: `Consider activity ${name} [${start}, ${end}]`,
        activities: this.activities.map(a => ({ start: a[0], end: a[1], name: a[2] })),
        sorted: sorted.map(a => ({ start: a[0], end: a[1], name: a[2] })),
        selected: [...selectedList],
        currentEnd: currentEnd,
        considering: idx,
        checkingStart: start,
        apply: () => {}
      });

      if (start >= currentEnd) {
        this.steps.push({
          lineNum: 8,
          state: `${name} is compatible! start(${start}) >= current_end(${currentEnd})`,
          activities: this.activities.map(a => ({ start: a[0], end: a[1], name: a[2] })),
          sorted: sorted.map(a => ({ start: a[0], end: a[1], name: a[2] })),
          selected: [...selectedList],
          currentEnd: currentEnd,
          considering: idx,
          compatible: true,
          apply: () => {}
        });

        selectedList.push({ start, end, name });
        currentEnd = end;

        this.steps.push({
          lineNum: 10,
          state: `Select ${name}, update current_end to ${end}`,
          activities: this.activities.map(a => ({ start: a[0], end: a[1], name: a[2] })),
          sorted: sorted.map(a => ({ start: a[0], end: a[1], name: a[2] })),
          selected: [...selectedList],
          currentEnd: currentEnd,
          considering: idx,
          justSelected: true,
          apply: () => {}
        });
      } else {
        this.steps.push({
          lineNum: 7,
          state: `${name} overlaps! start(${start}) < current_end(${currentEnd}). Skip.`,
          activities: this.activities.map(a => ({ start: a[0], end: a[1], name: a[2] })),
          sorted: sorted.map(a => ({ start: a[0], end: a[1], name: a[2] })),
          selected: [...selectedList],
          currentEnd: currentEnd,
          considering: idx,
          overlaps: true,
          apply: () => {}
        });
      }
    });

    this.steps.push({
      lineNum: 12,
      state: `Done! Selected ${selectedList.length} activities: ${selectedList.map(a => a.name).join(', ')}`,
      activities: this.activities.map(a => ({ start: a[0], end: a[1], name: a[2] })),
      sorted: sorted.map(a => ({ start: a[0], end: a[1], name: a[2] })),
      selected: [...selectedList],
      currentEnd: currentEnd,
      considering: -1,
      complete: true,
      apply: () => {}
    });
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { sorted, selected, currentEnd, considering, compatible, overlaps, justSelected, complete, sorting } = step;

    const maxTime = 16;
    const timelineWidth = 100; // percentage

    let html = '<div class="greedy-viz">';

    // Timeline header
    html += '<div class="timeline-header">';
    for (let t = 0; t <= maxTime; t += 2) {
      html += `<span class="time-mark" style="left: ${(t / maxTime) * 100}%">${t}</span>`;
    }
    html += '</div>';

    // Current end marker
    if (currentEnd > 0) {
      html += `<div class="current-end-marker" style="left: ${(currentEnd / maxTime) * 100}%">`;
      html += `<div class="marker-line"></div>`;
      html += `<div class="marker-label">end=${currentEnd}</div>`;
      html += '</div>';
    }

    // Activities timeline
    html += '<div class="activities-timeline">';
    const displayActivities = sorted || step.activities || [];

    displayActivities.forEach((act, idx) => {
      const left = (act.start / maxTime) * 100;
      const width = ((act.end - act.start) / maxTime) * 100;

      let barClass = 'activity-bar';
      const isSelected = selected && selected.some(s => s.name === act.name);

      if (isSelected) {
        barClass += ' selected';
      }
      if (considering === idx) {
        barClass += ' considering';
        if (compatible || justSelected) barClass += ' compatible';
        if (overlaps) barClass += ' overlaps';
      }
      if (complete && isSelected) {
        barClass += ' final';
      }

      html += `<div class="activity-row">`;
      html += `<span class="activity-name">${act.name}</span>`;
      html += `<div class="activity-track">`;
      html += `<div class="${barClass}" style="left: ${left}%; width: ${width}%;">`;
      html += `<span class="activity-range">${act.start}-${act.end}</span>`;
      html += `</div>`;
      html += `</div>`;
      html += `</div>`;
    });
    html += '</div>';

    // Selected activities summary
    if (selected && selected.length > 0) {
      html += '<div class="selected-summary">';
      html += `<span class="summary-label">Selected:</span>`;
      html += `<span class="summary-items">${selected.map(a => a.name).join(' -> ')}</span>`;
      html += '</div>';
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    const considering = step.sorted && step.considering >= 0 ? step.sorted[step.considering] : null;
    return {
      'current_end': step.currentEnd || 0,
      'considering': considering ? `${considering.name} [${considering.start},${considering.end}]` : '-',
      'selected_count': (step.selected || []).length
    };
  }

  getInputData() {
    return {
      'Problem': 'Activity Selection',
      'Activities': this.activities.length,
      'Strategy': 'Earliest finish time'
    };
  }
}

// Knapsack DP Animator (0/1 Knapsack)
class KnapsackAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);

    // Items: [weight, value, name]
    this.items = [
      [2, 3, 'A'],
      [3, 4, 'B'],
      [4, 5, 'C'],
      [5, 6, 'D']
    ];
    this.capacity = 8;
    this.dp = [];

    this.setCode(`def knapsack(items, capacity):
    n = len(items)
    # dp[i][w] = max value using first i items with capacity w
    dp = [[0] * (capacity + 1) for _ in range(n + 1)]

    for i in range(1, n + 1):
        weight, value = items[i-1]
        for w in range(capacity + 1):
            if weight <= w:
                # Can include item: max of include vs exclude
                dp[i][w] = max(dp[i-1][w],
                               dp[i-1][w-weight] + value)
            else:
                # Cannot include item
                dp[i][w] = dp[i-1][w]

    return dp[n][capacity]`);
  }

  buildSteps() {
    this.steps = [];
    const n = this.items.length;
    const W = this.capacity;

    // Initialize DP table
    this.dp = Array(n + 1).fill(null).map(() => Array(W + 1).fill(0));

    this.steps.push({
      lineNum: 1,
      state: `0/1 Knapsack: ${n} items, capacity ${W}`,
      dp: this.dp.map(row => [...row]),
      currentI: -1,
      currentW: -1,
      items: this.items.map(([w, v, name]) => ({ weight: w, value: v, name })),
      apply: () => {}
    });

    this.steps.push({
      lineNum: 4,
      state: `Initialize DP table with zeros (base case: 0 items = 0 value)`,
      dp: this.dp.map(row => [...row]),
      currentI: 0,
      currentW: -1,
      items: this.items.map(([w, v, name]) => ({ weight: w, value: v, name })),
      initializing: true,
      apply: () => {}
    });

    // Fill DP table
    for (let i = 1; i <= n; i++) {
      const [weight, value, name] = this.items[i - 1];

      this.steps.push({
        lineNum: 6,
        state: `Consider item ${name}: weight=${weight}, value=${value}`,
        dp: this.dp.map(row => [...row]),
        currentI: i,
        currentW: -1,
        currentItem: { weight, value, name },
        items: this.items.map(([w, v, name]) => ({ weight: w, value: v, name })),
        apply: () => {}
      });

      for (let w = 0; w <= W; w++) {
        if (weight <= w) {
          const exclude = this.dp[i - 1][w];
          const include = this.dp[i - 1][w - weight] + value;

          this.steps.push({
            lineNum: 9,
            state: `w=${w}: Can include ${name}. Exclude=${exclude}, Include=${include}`,
            dp: this.dp.map(row => [...row]),
            currentI: i,
            currentW: w,
            currentItem: { weight, value, name },
            items: this.items.map(([w, v, name]) => ({ weight: w, value: v, name })),
            comparing: true,
            excludeValue: exclude,
            includeValue: include,
            lookupCells: [[i - 1, w], [i - 1, w - weight]],
            apply: () => {}
          });

          this.dp[i][w] = Math.max(exclude, include);
          const chose = include > exclude ? 'include' : 'exclude';

          this.steps.push({
            lineNum: 10,
            state: `dp[${i}][${w}] = max(${exclude}, ${include}) = ${this.dp[i][w]} (${chose})`,
            dp: this.dp.map(row => [...row]),
            currentI: i,
            currentW: w,
            currentItem: { weight, value, name },
            items: this.items.map(([w, v, name]) => ({ weight: w, value: v, name })),
            filled: true,
            choice: chose,
            apply: () => {}
          });
        } else {
          this.steps.push({
            lineNum: 13,
            state: `w=${w}: Cannot include ${name} (weight ${weight} > ${w})`,
            dp: this.dp.map(row => [...row]),
            currentI: i,
            currentW: w,
            currentItem: { weight, value, name },
            items: this.items.map(([w, v, name]) => ({ weight: w, value: v, name })),
            cannotInclude: true,
            apply: () => {}
          });

          this.dp[i][w] = this.dp[i - 1][w];

          this.steps.push({
            lineNum: 14,
            state: `dp[${i}][${w}] = dp[${i-1}][${w}] = ${this.dp[i][w]}`,
            dp: this.dp.map(row => [...row]),
            currentI: i,
            currentW: w,
            currentItem: { weight, value, name },
            items: this.items.map(([w, v, name]) => ({ weight: w, value: v, name })),
            filled: true,
            choice: 'exclude',
            apply: () => {}
          });
        }
      }
    }

    // Traceback to find selected items
    const selected = [];
    let w = W;
    for (let i = n; i > 0 && w > 0; i--) {
      if (this.dp[i][w] !== this.dp[i - 1][w]) {
        selected.push(this.items[i - 1][2]);
        w -= this.items[i - 1][0];
      }
    }

    this.steps.push({
      lineNum: 16,
      state: `Maximum value: ${this.dp[n][W]}. Items: ${selected.reverse().join(', ')}`,
      dp: this.dp.map(row => [...row]),
      currentI: n,
      currentW: W,
      items: this.items.map(([w, v, name]) => ({ weight: w, value: v, name })),
      complete: true,
      selectedItems: selected,
      maxValue: this.dp[n][W],
      apply: () => {}
    });
  }

  render() {
    const step = this.steps[this.currentStep] || {};
    const { dp, currentI, currentW, currentItem, comparing, filled, lookupCells, complete, selectedItems, cannotInclude, choice } = step;

    const n = this.items.length;
    const W = this.capacity;

    let html = '<div class="knapsack-viz">';

    // Items display
    html += '<div class="knapsack-items">';
    this.items.forEach(([weight, value, name], idx) => {
      let itemClass = 'knapsack-item';
      if (currentItem && currentItem.name === name) itemClass += ' current';
      if (selectedItems && selectedItems.includes(name)) itemClass += ' selected';

      html += `<div class="${itemClass}">`;
      html += `<span class="item-name">${name}</span>`;
      html += `<span class="item-stats">w:${weight} v:${value}</span>`;
      html += `</div>`;
    });
    html += '</div>';

    // DP Table
    html += '<div class="dp-grid-container">';
    html += '<table class="dp-grid">';

    // Header row (capacity)
    html += '<tr><th></th>';
    for (let w = 0; w <= W; w++) {
      html += `<th class="capacity-header">w=${w}</th>`;
    }
    html += '</tr>';

    // DP rows
    for (let i = 0; i <= n; i++) {
      html += '<tr>';
      html += `<th class="item-header">${i === 0 ? '0' : this.items[i - 1][2]}</th>`;

      for (let w = 0; w <= W; w++) {
        let cellClass = 'dp-table-cell';
        const val = dp ? dp[i][w] : 0;

        if (i === currentI && w === currentW) {
          cellClass += ' current';
          if (comparing) cellClass += ' comparing';
          if (filled) cellClass += ' filled';
          if (cannotInclude) cellClass += ' cannot-include';
          if (choice === 'include') cellClass += ' included';
        }

        if (lookupCells) {
          lookupCells.forEach(([li, lw]) => {
            if (i === li && w === lw) cellClass += ' lookup';
          });
        }

        if (complete && i === n && w === W) cellClass += ' result';

        html += `<td class="${cellClass}">${val}</td>`;
      }
      html += '</tr>';
    }

    html += '</table>';
    html += '</div>';

    // Current decision info
    if (comparing) {
      html += '<div class="decision-info">';
      html += `<span>Exclude: ${step.excludeValue}</span>`;
      html += `<span>Include: ${step.includeValue}</span>`;
      html += '</div>';
    }

    if (complete) {
      html += '<div class="knapsack-result">';
      html += `<span>Max Value: ${step.maxValue}</span>`;
      html += `<span>Items: ${selectedItems.join(', ')}</span>`;
      html += '</div>';
    }

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const step = this.steps[this.currentStep] || {};
    return {
      'item': step.currentItem ? `${step.currentItem.name} (w:${step.currentItem.weight}, v:${step.currentItem.value})` : '-',
      'i': step.currentI >= 0 ? step.currentI : '-',
      'w': step.currentW >= 0 ? step.currentW : '-',
      'dp[i][w]': step.dp && step.currentI >= 0 && step.currentW >= 0 ? step.dp[step.currentI][step.currentW] : '-'
    };
  }

  getInputData() {
    return {
      'Items': this.items.map(([w, v, n]) => `${n}(${w},${v})`).join(' '),
      'Capacity': this.capacity
    };
  }
}

// Min Stack Visualizer
class MinStackAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.operations = [
      { op: 'push', val: 3 },
      { op: 'push', val: 5 },
      { op: 'getMin' },
      { op: 'push', val: 2 },
      { op: 'push', val: 1 },
      { op: 'getMin' },
      { op: 'pop' },
      { op: 'getMin' },
      { op: 'pop' },
      { op: 'top' },
    ];
    this.stack = [];
    this.currentOp = null;

    this.setCode(`class MinStack:
    def __init__(self):
        self.stack = []

    def push(self, val):
        cur_min = val if not self.stack else min(val, self.stack[-1][1])
        self.stack.append((val, cur_min))

    def pop(self):
        self.stack.pop()

    def top(self):
        return self.stack[-1][0]

    def getMin(self):
        return self.stack[-1][1]`);
  }

  buildSteps() {
    this.steps = [];
    let stack = [];

    this.steps.push({
      lineNum: 2,
      state: 'Initialize empty min stack',
      stack: [],
      op: { op: 'init' },
      apply: () => { this.stack = []; this.currentOp = 'init'; },
    });

    this.operations.forEach((operation) => {
      let stateText = '';
      if (operation.op === 'push') {
        const curMin = stack.length === 0 ? operation.val : Math.min(operation.val, stack[stack.length - 1].min);
        stack = [...stack, { val: operation.val, min: curMin }];
        stateText = `push(${operation.val}), min now ${curMin}`;
      } else if (operation.op === 'pop') {
        const popped = stack.pop();
        stateText = popped ? `pop() removed ${popped.val}` : 'pop() on empty stack';
      } else if (operation.op === 'top') {
        const top = stack[stack.length - 1];
        stateText = top ? `top() -> ${top.val}` : 'top() on empty stack';
      } else if (operation.op === 'getMin') {
        const top = stack[stack.length - 1];
        stateText = top ? `getMin() -> ${top.min}` : 'getMin() on empty stack';
      }

      const snapshot = stack.map(item => ({ ...item }));
      this.steps.push({
        lineNums: this.linesForOp(operation.op),
        state: stateText,
        stack: snapshot,
        op: operation,
        apply: () => { this.stack = snapshot; this.currentOp = operation.op; },
      });
    });
  }

  linesForOp(op) {
    switch (op) {
      case 'push': return [5, 6, 7];
      case 'pop': return [9];
      case 'top': return [12];
      case 'getMin': return [15];
      default: return [2];
    }
  }

  render() {
    if (!this.canvas) return;
    const items = [...this.stack].map((item, idx) => ({ ...item, idx }));
    let html = '<div class="stack-viz">';
    if (items.length === 0) {
      html += '<div class="stack-empty">Empty stack</div>';
    } else {
      for (let i = items.length - 1; i >= 0; i--) {
        const item = items[i];
        html += `<div class="stack-node">
          <div class="stack-value">${item.val}</div>
          <div class="stack-min">min ${item.min}</div>
        </div>`;
      }
    }
    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    const top = this.stack[this.stack.length - 1];
    return {
      size: this.stack.length,
      top: top ? top.val : 'none',
      min: top ? top.min : 'none',
      op: this.currentOp || 'init',
    };
  }

  getInputData() {
    return this.operations.map(o => o.val !== undefined ? `${o.op}(${o.val})` : o.op);
  }
}

// LRU Cache Visualizer (capacity 2)
class LRUCacheAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.capacity = 2;
    this.operations = [
      { op: 'put', key: 1, val: 1 },
      { op: 'put', key: 2, val: 2 },
      { op: 'get', key: 1 },
      { op: 'put', key: 3, val: 3 },
      { op: 'get', key: 2 },
      { op: 'put', key: 4, val: 4 },
      { op: 'get', key: 1 },
      { op: 'get', key: 3 },
      { op: 'get', key: 4 },
    ];
    this.order = [];
    this.map = {};
    this.current = null;

    this.setCode(`class LRUCache:
    def __init__(self, capacity):
        self.cap = capacity
        self.map = {}       # key -> (val, node)
        self.order = []     # most recent at front

    def get(self, key):
        if key not in self.map:
            return -1
        val = self.map[key]
        self._bump(key)
        return val

    def put(self, key, value):
        if key in self.map:
            self.map[key] = value
            self._bump(key)
            return
        if len(self.order) == self.cap:
            lru = self.order.pop()  # remove LRU
            self.map.pop(lru, None)
        self.order.insert(0, key)
        self.map[key] = value`);
  }

  buildSteps() {
    this.steps = [];
    let map = {};
    let order = [];

    this.steps.push({
      lineNum: 2,
      state: `Capacity ${this.capacity}, cache empty`,
      apply: () => { this.map = {}; this.order = []; this.current = 'init'; },
    });

    this.operations.forEach((op) => {
      let stateText = '';
      if (op.op === 'put') {
        if (map.hasOwnProperty(op.key)) {
          map = { ...map, [op.key]: op.val };
          order = [op.key, ...order.filter(k => k !== op.key)];
          stateText = `put(${op.key}, ${op.val}) update & move to MRU`;
        } else {
          if (order.length === this.capacity) {
            const evict = order.pop();
            const { [evict]: _, ...rest } = map;
            map = rest;
            stateText = `put(${op.key}, ${op.val}) evicts ${evict}`;
          } else {
            stateText = `put(${op.key}, ${op.val})`;
          }
          order = [op.key, ...order];
          map = { ...map, [op.key]: op.val };
        }
      } else if (op.op === 'get') {
        if (map.hasOwnProperty(op.key)) {
          order = [op.key, ...order.filter(k => k !== op.key)];
          stateText = `get(${op.key}) -> ${map[op.key]} (bump to MRU)`;
        } else {
          stateText = `get(${op.key}) -> -1 (miss)`;
        }
      }

      const snapOrder = [...order];
      const snapMap = { ...map };

      this.steps.push({
        lineNums: op.op === 'put' ? [12, 13, 19] : [7, 8, 9],
        state: stateText,
        order: snapOrder,
        map: snapMap,
        op,
        apply: () => { this.order = snapOrder; this.map = snapMap; this.current = `${op.op}(${op.key}${op.val !== undefined ? ',' + op.val : ''})`; },
      });
    });
  }

  render() {
    if (!this.canvas) return;
    let html = '<div class="lru-viz">';
    html += '<div class="lru-order"><div class="lru-label">Order (MRU → LRU)</div>';
    if (this.order.length === 0) {
      html += '<div class="lru-empty">Cache empty</div>';
    } else {
      html += '<div class="lru-row">';
      this.order.forEach((key, idx) => {
        html += `<div class="lru-card ${idx === 0 ? 'mru' : ''}"><div class="lru-key">k=${key}</div><div class="lru-val">v=${this.map[key]}</div></div>`;
      });
      html += '</div>';
    }
    html += '</div>';

    html += '<div class="lru-map"><div class="lru-label">Hash Map</div>';
    if (Object.keys(this.map).length === 0) {
      html += '<div class="lru-empty">No entries</div>';
    } else {
      html += '<div class="lru-map-grid">';
      Object.entries(this.map).forEach(([k, v]) => {
        html += `<div class="lru-map-entry"><span class="map-key">key ${k}</span><span class="map-val">→ ${v}</span></div>`;
      });
      html += '</div>';
    }
    html += '</div></div>';

    this.canvas.innerHTML = html;
  }

  getVariables() {
    return {
      capacity: this.capacity,
      size: Object.keys(this.map).length,
      order: this.order.join(' → ') || 'empty',
      op: this.current || 'init',
    };
  }

  getInputData() {
    return this.operations.map(o => o.val !== undefined ? `${o.op}(${o.key},${o.val})` : `${o.op}(${o.key})`);
  }
}

// Trie (insert/search) Visualizer
class TrieAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.trie = this.createNode();
    this.operations = [
      { op: 'insert', word: 'cat' },
      { op: 'insert', word: 'car' },
      { op: 'insert', word: 'dog' },
      { op: 'search', word: 'car' },
      { op: 'startsWith', word: 'ca' },
      { op: 'search', word: 'cow' },
    ];
    this.currentWord = '';
    this.charIndex = -1;

    this.setCode(`class TrieNode:
    def __init__(self):
        self.children = {}
        self.end = False

class Trie:
    def insert(self, word):
        node = self.root
        for ch in word:
            node = node.children.setdefault(ch, TrieNode())
        node.end = True

    def search(self, word):
        node = self._walk(word)
        return bool(node and node.end)

    def startsWith(self, prefix):
        return bool(self._walk(prefix))`);
  }

  createNode() {
    return { children: {}, end: false };
  }

  cloneNode(node) {
    const cloned = { end: node.end, children: {} };
    Object.entries(node.children).forEach(([ch, child]) => {
      cloned.children[ch] = this.cloneNode(child);
    });
    return cloned;
  }

  buildSteps() {
    this.steps = [];
    let root = this.createNode();

    this.steps.push({
      lineNum: 2,
      state: 'Initialize empty trie',
      trie: this.cloneNode(root),
      op: { op: 'init' },
      apply: () => { this.trie = this.cloneNode(root); this.currentWord = ''; this.charIndex = -1; },
    });

    this.operations.forEach((op) => {
      if (op.op === 'insert') {
        let node = root;
        op.word.split('').forEach((ch, idx) => {
          if (!node.children[ch]) node.children[ch] = this.createNode();
          node = node.children[ch];
          const snapshot = this.cloneNode(root);
          this.steps.push({
            lineNums: [9, 10, 11],
            state: `insert("${op.word}") visiting '${ch}' (idx ${idx})`,
            trie: snapshot,
            op,
            charIndex: idx,
            currentWord: op.word,
            apply: () => { this.trie = snapshot; this.currentWord = op.word; this.charIndex = idx; },
          });
        });
        node.end = true;
        const snapshot = this.cloneNode(root);
        this.steps.push({
          lineNum: 12,
          state: `mark end of "${op.word}"`,
          trie: snapshot,
          op,
          charIndex: op.word.length - 1,
          currentWord: op.word,
          apply: () => { this.trie = snapshot; this.currentWord = op.word; this.charIndex = op.word.length - 1; },
        });
      } else {
        // search / startsWith
        let node = root;
        let failedAt = -1;
        op.word.split('').forEach((ch, idx) => {
          if (node && node.children[ch]) {
            node = node.children[ch];
          } else {
            node = null;
            failedAt = idx;
          }
          const snapshot = this.cloneNode(root);
          this.steps.push({
            lineNums: op.op === 'search' ? [14, 15, 16] : [19],
            state: `${op.op}("${op.word}") at '${ch}' (idx ${idx})` + (failedAt === -1 ? '' : ' - missing'),
            trie: snapshot,
            op,
            charIndex: idx,
            currentWord: op.word,
            apply: () => { this.trie = snapshot; this.currentWord = op.word; this.charIndex = idx; },
          });
        });
        const found = !!node && (op.op === 'startsWith' || node.end);
        const snapshot = this.cloneNode(root);
        this.steps.push({
          lineNums: op.op === 'search' ? [17] : [19],
          state: `${op.op}("${op.word}") -> ${found}`,
          trie: snapshot,
          op,
          charIndex: op.word.length - 1,
          currentWord: op.word,
          apply: () => { this.trie = snapshot; this.currentWord = op.word; this.charIndex = op.word.length - 1; },
        });
      }
    });
  }

  render() {
    if (!this.canvas) return;

    const buildList = (node, prefix) => {
      let html = '';
      Object.entries(node.children).forEach(([ch, child]) => {
        const nextPrefix = prefix + ch;
        html += `<div class="trie-node">
          <div class="trie-chip ${nextPrefix === this.currentPrefix() ? 'active' : ''}">
            ${escapeHtml(nextPrefix)}${child.end ? '<span class="trie-end">★</span>' : ''}
          </div>
          <div class="trie-children">${buildList(child, nextPrefix)}</div>
        </div>`;
      });
      return html;
    };

    let html = '<div class="trie-viz">';
    html += '<div class="trie-root">root</div>';
    html += `<div class="trie-children">${buildList(this.trie, '')}</div>`;
    html += '</div>';
    this.canvas.innerHTML = html;
  }

  currentPrefix() {
    if (!this.currentWord || this.charIndex < 0) return '';
    return this.currentWord.slice(0, this.charIndex + 1);
  }

  getVariables() {
    return {
      word: this.currentWord || 'n/a',
      prefix: this.currentPrefix() || 'n/a',
      idx: this.charIndex,
    };
  }

  getInputData() {
    return this.operations.map(o => `${o.op}("${o.word}")`);
  }
}

// Car Fleet Visualizer
class CarFleetAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.target = 12;
    this.positions = [10, 8, 0, 5, 3];
    this.speeds = [2, 4, 1, 1, 3];
    this.cars = this.positions.map((pos, idx) => ({ pos, speed: this.speeds[idx] }));
    this.stack = [];
    this.current = null;

    this.setCode(`def carFleet(target, position, speed):
    pairs = sorted(zip(position, speed), reverse=True)
    stack = []
    for pos, spd in pairs:
        time = (target - pos) / spd
        if not stack or time > stack[-1]:
            stack.append(time)
    return len(stack)`);
  }

  buildSteps() {
    this.steps = [];
    const pairs = this.cars.map(c => ({ ...c })).sort((a, b) => b.pos - a.pos);
    let stack = [];
    this.steps.push({
      lineNum: 2,
      state: 'Sort cars by position descending',
      stack: [],
      current: null,
      apply: () => { this.stack = []; this.current = null; },
    });

    pairs.forEach((car, idx) => {
      const time = (this.target - car.pos) / car.speed;
      let stateText = `Car at ${car.pos} (spd ${car.speed}) time ${time.toFixed(2)}`;
      if (stack.length === 0 || time > stack[stack.length - 1]) {
        stack.push(time);
        stateText += ' starts new fleet';
      } else {
        stateText += ' joins fleet ahead';
      }
      const snap = [...stack];
      this.steps.push({
        lineNums: [3, 4, 5, 6],
        state: stateText,
        stack: snap,
        current: { ...car, time },
        apply: () => { this.stack = snap; this.current = { ...car, time }; },
      });
    });
  }

  render() {
    if (!this.canvas) return;
    const roadLength = this.target;
    let html = '<div class="fleet-viz">';
    html += `<div class="fleet-road">`;
    this.cars.sort((a, b) => b.pos - a.pos).forEach((car, idx) => {
      const percent = Math.max(0, Math.min(100, (car.pos / roadLength) * 100));
      html += `<div class="fleet-car" style="left:${percent}%">
        <div class="car-label">pos ${car.pos}, v ${car.speed}</div>
      </div>`;
    });
    html += `<div class="fleet-target" style="left:100%;">Target ${this.target}</div>`;
    html += '</div>';

    html += '<div class="fleet-stack"><div class="fleet-label">Fleet arrival times (stack)</div>';
    if (this.stack.length === 0) {
      html += '<div class="fleet-empty">No fleets yet</div>';
    } else {
      html += '<div class="fleet-stack-items">';
      this.stack.slice().reverse().forEach(time => {
        html += `<div class="fleet-time">time ${time.toFixed(2)}</div>`;
      });
      html += '</div>';
    }
    html += '</div></div>';

    this.canvas.innerHTML = html;
  }

  getVariables() {
    return {
      target: this.target,
      fleets: this.stack.length,
      current: this.current ? `pos ${this.current.pos}, v ${this.current.speed}, t ${this.current.time.toFixed(2)}` : 'n/a',
    };
  }

  getInputData() {
    return { position: this.positions, speed: this.speeds, target: this.target };
  }
}

// Median of Two Sorted Arrays Visualizer
class MedianTwoArraysAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.A = [1, 3, 8];
    this.B = [7, 9, 10, 11];
    this.low = 0;
    this.high = this.A.length;
    this.partA = 0;
    this.partB = 0;
    this.median = null;
    this.stateText = '';

    this.setCode(`def findMedianSortedArrays(nums1, nums2):
    A, B = nums1, nums2
    if len(A) > len(B):
        A, B = B, A
    low, high = 0, len(A)
    while low <= high:
        i = (low + high) // 2
        j = (len(A) + len(B) + 1)//2 - i
        maxLeftA = -inf if i == 0 else A[i-1]
        minRightA = inf if i == len(A) else A[i]
        maxLeftB = -inf if j == 0 else B[j-1]
        minRightB = inf if j == len(B) else B[j]
        if maxLeftA <= minRightB and maxLeftB <= minRightA:
            if (len(A)+len(B)) % 2 == 0:
                return (max(maxLeftA, maxLeftB) + min(minRightA, minRightB))/2
            return max(maxLeftA, maxLeftB)
        elif maxLeftA > minRightB:
            high = i - 1
        else:
            low = i + 1`);
  }

  buildSteps() {
    this.steps = [];
    let low = 0, high = this.A.length;
    const total = this.A.length + this.B.length;

    while (low <= high) {
      const i = Math.floor((low + high) / 2);
      const j = Math.floor((total + 1) / 2) - i;

      const maxLeftA = i === 0 ? -Infinity : this.A[i - 1];
      const minRightA = i === this.A.length ? Infinity : this.A[i];
      const maxLeftB = j === 0 ? -Infinity : this.B[j - 1];
      const minRightB = j === this.B.length ? Infinity : this.B[j];

      let state = `i=${i}, j=${j}, maxLeftA=${maxLeftA}, minRightA=${minRightA}, maxLeftB=${maxLeftB}, minRightB=${minRightB}`;

      let found = false;
      let median = null;
      if (maxLeftA <= minRightB && maxLeftB <= minRightA) {
        found = true;
        if (total % 2 === 0) {
          median = (Math.max(maxLeftA, maxLeftB) + Math.min(minRightA, minRightB)) / 2;
        } else {
          median = Math.max(maxLeftA, maxLeftB);
        }
        state += ` -> median ${median}`;
      } else if (maxLeftA > minRightB) {
        state += ' (move left)';
        high = i - 1;
      } else {
        state += ' (move right)';
        low = i + 1;
      }

      this.steps.push({
        lineNums: [6, 7, 8, 12, 13],
        state,
        low,
        high,
        partA: i,
        partB: j,
        median,
        found,
        apply: () => {
          this.low = low;
          this.high = high;
          this.partA = i;
          this.partB = j;
          this.median = median;
          this.stateText = state;
        },
      });

      if (found) break;
    }
  }

  render() {
    if (!this.canvas) return;
    const renderRow = (arr, part) => {
      let html = '<div class="median-row">';
      arr.forEach((val, idx) => {
        const before = idx < part;
        const boundary = idx === part;
        html += `<div class="median-cell ${before ? 'left' : 'right'}">${val}</div>`;
        if (boundary) {
          html += '<div class="median-partition">|</div>';
        }
      });
      if (part === arr.length) {
        html += '<div class="median-partition">|</div>';
      }
      html += '</div>';
      return html;
    };

    let html = '<div class="median-viz">';
    html += '<div class="median-label">Array A (shorter)</div>';
    html += renderRow(this.A, this.partA);
    html += '<div class="median-label">Array B</div>';
    html += renderRow(this.B, this.partB);
    html += `<div class="median-state">${escapeHtml(this.stateText)}</div>`;
    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    return {
      low: this.low,
      high: this.high,
      partA: this.partA,
      partB: this.partB,
      median: this.median !== null ? this.median : 'pending',
    };
  }

  getInputData() {
    return { nums1: this.A, nums2: this.B };
  }
}

// Kth Largest Element Visualizer (min-heap approach)
class KthLargestAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.nums = [4, 5, 8, 2, 3, 5, 10, 9, 4];
    this.k = 3;
    this.heap = [];
    this.current = null;

    this.setCode(`import heapq

def kthLargest(nums, k):
    heap = []
    for num in nums:
        heapq.heappush(heap, num)
        if len(heap) > k:
            heapq.heappop(heap)
    return heap[0]`);
  }

  buildSteps() {
    this.steps = [];
    let heap = [];
    this.steps.push({
      lineNum: 2,
      state: `k=${this.k}, start with empty min-heap`,
      heap: [],
      apply: () => { this.heap = []; this.current = null; },
    });

    this.nums.forEach((num) => {
      heap.push(num);
      heap.sort((a, b) => a - b);
      let state = `push ${num}`;
      if (heap.length > this.k) {
        const removed = heap.shift();
        state += `, pop ${removed} to keep size ${this.k}`;
      }
      const snap = [...heap];
      this.steps.push({
        lineNums: [6, 7, 8],
        state,
        heap: snap,
        current: num,
        apply: () => { this.heap = snap; this.current = num; },
      });
    });
  }

  render() {
    if (!this.canvas) return;
    let html = '<div class="heap-viz">';
    if (this.heap.length === 0) {
      html += '<div class="heap-empty">Heap empty</div>';
    } else {
      html += '<div class="heap-levels">';
      this.heap.forEach((val, idx) => {
        html += `<div class="heap-node">${val}</div>`;
      });
      html += '</div>';
    }
    html += `<div class="heap-footer">Top of heap = current ${this.k}th largest${this.heap.length > 0 ? ` (${this.heap[0]})` : ''}</div>`;
    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    return {
      k: this.k,
      heap: this.heap,
      current: this.current !== null ? this.current : 'n/a',
    };
  }

  getInputData() {
    return { nums: this.nums, k: this.k };
  }
}

// N-Queens Visualizer (n=4)
class NQueensAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.n = 4;
    this.boards = [];
    this.currentBoard = this.emptyBoard();
    this.stepBoards = this.precomputeBoards();

    this.setCode(`def solveNQueens(n):
    res = []
    board = [['.'] * n for _ in range(n)]

    def backtrack(r):
        if r == n:
            res.append([''.join(row) for row in board])
            return
        for c in range(n):
            if is_safe(r, c):
                board[r][c] = 'Q'
                backtrack(r + 1)
                board[r][c] = '.'

    backtrack(0)
    return res`);
  }

  emptyBoard() {
    return Array.from({ length: this.n }, () => Array(this.n).fill('.'));
  }

  precomputeBoards() {
    // Simple scripted path to one solution for n=4
    return [
      { row: 0, col: 1 },
      { row: 1, col: 3 },
      { row: 2, col: 0 },
      { row: 3, col: 2 },
    ];
  }

  buildSteps() {
    this.steps = [];
    let board = this.emptyBoard();

    this.steps.push({
      lineNum: 2,
      state: 'Start backtracking on 4x4 board',
      board: board.map(row => [...row]),
      apply: () => { this.currentBoard = board.map(row => [...row]); },
    });

    this.stepBoards.forEach((placement, idx) => {
      board[placement.row][placement.col] = 'Q';
      const snap = board.map(row => [...row]);
      this.steps.push({
        lineNums: [5, 6, 7, 8],
        state: `Place queen at row ${placement.row}, col ${placement.col}`,
        board: snap,
        apply: () => { this.currentBoard = snap.map(row => [...row]); },
      });
    });

    this.steps.push({
      lineNum: 11,
      state: 'Solution found',
      board: board.map(row => [...row]),
      apply: () => { this.currentBoard = board.map(row => [...row]); },
    });
  }

  render() {
    if (!this.canvas) return;
    let html = '<div class="board-viz">';
    this.currentBoard.forEach((row, rIdx) => {
      html += '<div class="board-row">';
      row.forEach((cell, cIdx) => {
        const isQueen = cell === 'Q';
        html += `<div class="board-cell ${isQueen ? 'queen' : ''}">${isQueen ? '♛' : ''}</div>`;
      });
      html += '</div>';
    });
    html += '</div>';
    this.canvas.innerHTML = html;
  }

  getVariables() {
    return {
      n: this.n,
      queensPlaced: this.currentBoard.flat().filter(c => c === 'Q').length,
    };
  }

  getInputData() {
    return { n: this.n };
  }
}

// Sandbox Animator - Dynamic visualization of user Python code
class SandboxAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    this.executionTrace = null;
    this.currentLocals = {};
    this.currentStructures = [];
    this.userCode = '';
    this.varsEl = document.getElementById('sandbox-vars');
    this.inputEl = null;
  }

  async executeCode(code) {
    this.userCode = code;
    this.setCode(code);

    try {
      const response = await fetch('/api/sandbox/execute', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code, timeout: 10000 })
      });

      const result = await response.json();

      if (result.success) {
        this.executionTrace = result;
        this.buildStepsFromTrace();
        this.updateStepCounter();
        if (this.steps.length > 0) {
          this.goToStep(0);
        }
        return { success: true, output: result.output, stepCount: this.steps.length };
      } else {
        return { success: false, error: result.error };
      }
    } catch (err) {
      return { success: false, error: err.message };
    }
  }

  buildStepsFromTrace() {
    if (!this.executionTrace || !this.executionTrace.steps) {
      this.steps = [];
      return;
    }

    this.steps = this.executionTrace.steps.map((traceStep, idx) => ({
      lineNum: traceStep.lineNum,
      state: traceStep.state || `Line ${traceStep.lineNum}`,
      locals: traceStep.locals || {},
      structures: traceStep.structures || [],
      apply: () => {
        this.currentLocals = traceStep.locals || {};
        this.currentStructures = traceStep.structures || [];
      }
    }));
  }

  buildSteps() {
    // Override - steps are built from trace
    if (this.executionTrace) {
      this.buildStepsFromTrace();
    }
  }

  render() {
    if (!this.canvas) return;

    if (!this.currentStructures || this.currentStructures.length === 0) {
      this.canvas.innerHTML = '<div class="sandbox-empty">No data structures detected in current step</div>';
      return;
    }

    let html = '<div class="sandbox-viz">';

    this.currentStructures.forEach(struct => {
      html += `<div class="structure-section">`;
      html += `<div class="structure-header"><span class="structure-name">${escapeHtml(struct.name)}</span><span class="structure-type">${struct.type}</span></div>`;

      switch (struct.type) {
        case 'array':
          html += this.renderArray(struct.data, struct.highlights);
          break;
        case 'linked_list':
          html += this.renderLinkedList(struct.data);
          break;
        case 'binary_tree':
          html += this.renderBinaryTree(struct.data);
          break;
        case 'graph':
          html += this.renderGraph(struct.data);
          break;
        case 'hash_table':
          html += this.renderHashTable(struct.data);
          break;
        case 'set':
          html += this.renderSet(struct.data);
          break;
        default:
          html += `<div class="structure-data">${escapeHtml(JSON.stringify(struct.data))}</div>`;
      }

      html += '</div>';
    });

    html += '</div>';
    this.canvas.innerHTML = html;
  }

  renderArray(data, highlights) {
    if (!Array.isArray(data)) return '<div class="error">Invalid array data</div>';

    let html = '<div class="array-viz">';
    data.forEach((val, idx) => {
      let classes = ['array-cell'];

      if (highlights) {
        Object.entries(highlights).forEach(([name, hIdx]) => {
          if (hIdx === idx) {
            classes.push('highlight');
            classes.push(`ptr-${name}`);
          }
        });
      }

      html += `<div class="${classes.join(' ')}">`;
      html += `<div class="cell-value">${escapeHtml(String(val))}</div>`;
      html += `<div class="cell-index">${idx}</div>`;

      // Show pointer labels
      if (highlights) {
        const pointers = Object.entries(highlights)
          .filter(([_, hIdx]) => hIdx === idx)
          .map(([name, _]) => name);
        if (pointers.length > 0) {
          html += `<div class="cell-pointers">${pointers.join(', ')}</div>`;
        }
      }

      html += '</div>';
    });
    html += '</div>';
    return html;
  }

  renderLinkedList(data) {
    if (!data || !data.nodes) return '<div class="error">Invalid linked list data</div>';

    let html = '<div class="linked-list-viz">';
    data.nodes.forEach((node, idx) => {
      html += `<div class="list-node">`;
      html += `<div class="node-val">${escapeHtml(String(node.val))}</div>`;
      html += '</div>';
      if (idx < data.nodes.length - 1) {
        html += '<div class="list-arrow">-&gt;</div>';
      }
    });
    if (data.hasCycle) {
      html += '<div class="cycle-indicator">(cycle detected)</div>';
    }
    html += '</div>';
    return html;
  }

  renderBinaryTree(data) {
    if (!data) return '<div class="tree-empty">Empty tree</div>';

    // Calculate tree dimensions
    const getDepth = (node) => {
      if (!node) return 0;
      return 1 + Math.max(getDepth(node.left), getDepth(node.right));
    };

    const depth = getDepth(data);
    const width = Math.pow(2, depth) * 60;
    const height = depth * 80;

    let html = `<div class="binary-tree-viz" style="min-width: ${width}px; min-height: ${height}px;">`;
    html += '<svg class="tree-svg" width="100%" height="100%" viewBox="0 0 ' + width + ' ' + height + '">';

    // Render nodes recursively
    const renderNode = (node, x, y, spread) => {
      if (!node) return '';

      let svg = '';
      const nodeRadius = 20;
      const verticalGap = 70;

      // Draw edges first
      if (node.left) {
        const childX = x - spread / 2;
        const childY = y + verticalGap;
        svg += `<line x1="${x}" y1="${y + nodeRadius}" x2="${childX}" y2="${childY - nodeRadius}" class="tree-edge"/>`;
        svg += renderNode(node.left, childX, childY, spread / 2);
      }
      if (node.right) {
        const childX = x + spread / 2;
        const childY = y + verticalGap;
        svg += `<line x1="${x}" y1="${y + nodeRadius}" x2="${childX}" y2="${childY - nodeRadius}" class="tree-edge"/>`;
        svg += renderNode(node.right, childX, childY, spread / 2);
      }

      // Draw node
      svg += `<circle cx="${x}" cy="${y}" r="${nodeRadius}" class="tree-node"/>`;
      svg += `<text x="${x}" y="${y + 5}" class="tree-label">${escapeHtml(String(node.val))}</text>`;

      return svg;
    };

    html += renderNode(data, width / 2, 30, width / 4);
    html += '</svg></div>';
    return html;
  }

  renderGraph(data) {
    if (!data || !data.nodes) return '<div class="error">Invalid graph data</div>';

    const nodes = data.nodes;
    const edges = data.edges || [];

    // Simple circular layout
    const width = 400;
    const height = 300;
    const centerX = width / 2;
    const centerY = height / 2;
    const radius = Math.min(width, height) / 3;

    const positions = {};
    nodes.forEach((node, idx) => {
      const angle = (2 * Math.PI * idx) / nodes.length - Math.PI / 2;
      positions[node] = {
        x: centerX + radius * Math.cos(angle),
        y: centerY + radius * Math.sin(angle)
      };
    });

    let html = `<div class="graph-viz">`;
    html += `<svg class="graph-svg" viewBox="0 0 ${width} ${height}">`;

    // Draw edges
    edges.forEach(edge => {
      const from = positions[edge.from];
      const to = positions[edge.to];
      if (from && to) {
        html += `<line x1="${from.x}" y1="${from.y}" x2="${to.x}" y2="${to.y}" class="graph-edge"/>`;
      }
    });

    // Draw nodes
    nodes.forEach(node => {
      const pos = positions[node];
      html += `<circle cx="${pos.x}" cy="${pos.y}" r="20" class="graph-node"/>`;
      html += `<text x="${pos.x}" y="${pos.y + 5}" class="graph-label">${escapeHtml(String(node))}</text>`;
    });

    html += '</svg></div>';
    return html;
  }

  renderHashTable(data) {
    if (!data || typeof data !== 'object') return '<div class="error">Invalid hash table data</div>';

    let html = '<div class="hash-table-viz">';
    Object.entries(data).forEach(([key, value]) => {
      html += `<div class="hash-entry">`;
      html += `<span class="hash-key">${escapeHtml(key)}</span>`;
      html += `<span class="hash-arrow">:</span>`;
      html += `<span class="hash-value">${escapeHtml(JSON.stringify(value))}</span>`;
      html += '</div>';
    });
    html += '</div>';
    return html;
  }

  renderSet(data) {
    if (!Array.isArray(data)) return '<div class="error">Invalid set data</div>';

    let html = '<div class="set-viz">';
    html += '<span class="set-brace">{</span>';
    data.forEach((val, idx) => {
      html += `<span class="set-item">${escapeHtml(String(val))}</span>`;
      if (idx < data.length - 1) {
        html += '<span class="set-comma">, </span>';
      }
    });
    html += '<span class="set-brace">}</span>';
    html += '</div>';
    return html;
  }

  getVariables() {
    return this.currentLocals || {};
  }

  updateStepCounter() {
    const counter = document.getElementById('sandbox-step-counter');
    if (counter) {
      counter.textContent = `Step ${this.currentStep + 1}/${this.steps.length}`;
    }
  }

  reset() {
    this.pause();
    this.currentStep = 0;
    if (this.steps.length > 0) {
      this.goToStep(0);
    }
  }
}

// Sandbox state
const sandboxState = {
  isActive: false,
  animator: null,
  isExecuting: false
};

// Toggle sandbox mode
function toggleSandboxMode() {
  sandboxState.isActive = !sandboxState.isActive;

  const sandboxMode = document.getElementById('sandbox-mode');
  const detailEmpty = document.getElementById('detail-empty');
  const detail = document.getElementById('detail');
  const toggleBtn = document.getElementById('sandbox-toggle');

  if (sandboxState.isActive) {
    // Activate sandbox mode
    sandboxMode.classList.remove('hidden');
    detailEmpty.classList.add('hidden');
    detail.classList.add('hidden');

    if (toggleBtn) {
      toggleBtn.classList.add('active');
    }

    // Initialize sandbox animator if not already
    if (!sandboxState.animator) {
      initSandboxAnimator();
    }

    // Deselect any chapter
    document.querySelectorAll('.chapter-item').forEach(el => el.classList.remove('active'));
  } else {
    // Deactivate sandbox mode
    sandboxMode.classList.add('hidden');

    if (toggleBtn) {
      toggleBtn.classList.remove('active');
    }

    // Show empty state if no chapter selected
    if (!state.current) {
      detailEmpty.classList.remove('hidden');
    } else {
      detail.classList.remove('hidden');
    }

    // Cleanup animator
    if (sandboxState.animator) {
      sandboxState.animator.pause();
    }
  }
}

function initSandboxAnimator() {
  const canvas = document.getElementById('sandbox-canvas');
  const codeEl = document.getElementById('code-lines');
  const stateEl = document.getElementById('sandbox-state');

  sandboxState.animator = new SandboxAnimator(canvas, codeEl, stateEl, {});
}

async function runSandboxCode() {
  if (sandboxState.isExecuting) return;

  const codeTextarea = document.getElementById('sandbox-code');
  const outputEl = document.getElementById('sandbox-output');
  const stateEl = document.getElementById('sandbox-state');
  const runBtn = document.getElementById('sandbox-run');
  const structuresEl = document.getElementById('sandbox-structures');

  const code = codeTextarea.value.trim();
  if (!code) {
    outputEl.textContent = 'Error: No code provided';
    return;
  }

  // Update UI for execution
  sandboxState.isExecuting = true;
  runBtn.textContent = 'Running...';
  runBtn.disabled = true;
  stateEl.textContent = 'Executing code...';
  outputEl.textContent = '';

  try {
    if (!sandboxState.animator) {
      initSandboxAnimator();
    }

    const result = await sandboxState.animator.executeCode(code);

    if (result.success) {
      outputEl.textContent = result.output || '(no output)';
      stateEl.textContent = `Execution complete. ${result.stepCount} steps captured. Use controls to step through.`;

      // Update detected structures
      const structures = sandboxState.animator.currentStructures;
      if (structures && structures.length > 0) {
        const types = [...new Set(structures.map(s => s.type))];
        structuresEl.textContent = types.join(', ');
      } else {
        structuresEl.textContent = 'None detected';
      }
    } else {
      outputEl.textContent = `Error: ${result.error}`;
      stateEl.textContent = 'Execution failed. Check the error above.';
      structuresEl.textContent = 'None';
    }
  } catch (err) {
    outputEl.textContent = `Error: ${err.message}`;
    stateEl.textContent = 'Execution failed.';
  } finally {
    sandboxState.isExecuting = false;
    runBtn.textContent = 'Run Code';
    runBtn.disabled = false;
  }
}

function clearSandbox() {
  const codeTextarea = document.getElementById('sandbox-code');
  const outputEl = document.getElementById('sandbox-output');
  const stateEl = document.getElementById('sandbox-state');
  const canvas = document.getElementById('sandbox-canvas');
  const varsEl = document.getElementById('sandbox-vars');
  const structuresEl = document.getElementById('sandbox-structures');
  const stepCounter = document.getElementById('sandbox-step-counter');

  codeTextarea.value = '';
  outputEl.textContent = '';
  stateEl.textContent = 'Click "Run Code" to execute your Python code and visualize the algorithm.';
  canvas.innerHTML = '';
  varsEl.innerHTML = '<div class="vars-empty">No variables tracked</div>';
  structuresEl.textContent = 'None';
  stepCounter.textContent = 'Step 0/0';

  if (sandboxState.animator) {
    sandboxState.animator.steps = [];
    sandboxState.animator.currentStep = 0;
    sandboxState.animator.currentLocals = {};
    sandboxState.animator.currentStructures = [];
  }
}

// Handle Python-specific indentation in the code editor
function handlePythonIndentation(e) {
  const textarea = e.target;
  const INDENT = '    '; // 4 spaces for Python
  const INDENT_SIZE = 4;

  // Get cursor position and text
  const start = textarea.selectionStart;
  const end = textarea.selectionEnd;
  const value = textarea.value;

  // Find the start of the current line
  const lineStart = value.lastIndexOf('\n', start - 1) + 1;
  const lineEnd = value.indexOf('\n', start);
  const currentLine = value.substring(lineStart, lineEnd === -1 ? value.length : lineEnd);

  // Get current line's indentation
  const currentIndent = currentLine.match(/^(\s*)/)[1];

  if (e.key === 'Tab') {
    e.preventDefault();

    if (e.shiftKey) {
      // Shift+Tab: Dedent
      if (start === end) {
        // No selection - dedent current line
        if (currentIndent.length >= INDENT_SIZE) {
          const newIndent = currentIndent.substring(INDENT_SIZE);
          const newValue = value.substring(0, lineStart) + newIndent + currentLine.trimStart() + value.substring(lineEnd === -1 ? value.length : lineEnd);
          textarea.value = newValue;
          textarea.selectionStart = textarea.selectionEnd = start - INDENT_SIZE;
        }
      } else {
        // Selection - dedent all selected lines
        const selStart = value.lastIndexOf('\n', start - 1) + 1;
        const selEnd = value.indexOf('\n', end - 1);
        const selectedText = value.substring(selStart, selEnd === -1 ? value.length : selEnd);
        const lines = selectedText.split('\n');
        const dedentedLines = lines.map(line => {
          if (line.startsWith(INDENT)) {
            return line.substring(INDENT_SIZE);
          } else if (line.match(/^\s+/)) {
            return line.replace(/^\s{1,4}/, '');
          }
          return line;
        });
        const newText = dedentedLines.join('\n');
        textarea.value = value.substring(0, selStart) + newText + value.substring(selEnd === -1 ? value.length : selEnd);
        textarea.selectionStart = selStart;
        textarea.selectionEnd = selStart + newText.length;
      }
    } else {
      // Tab: Indent
      if (start === end) {
        // No selection - insert tab at cursor
        textarea.value = value.substring(0, start) + INDENT + value.substring(end);
        textarea.selectionStart = textarea.selectionEnd = start + INDENT_SIZE;
      } else {
        // Selection - indent all selected lines
        const selStart = value.lastIndexOf('\n', start - 1) + 1;
        const selEnd = value.indexOf('\n', end - 1);
        const selectedText = value.substring(selStart, selEnd === -1 ? value.length : selEnd);
        const lines = selectedText.split('\n');
        const indentedLines = lines.map(line => INDENT + line);
        const newText = indentedLines.join('\n');
        textarea.value = value.substring(0, selStart) + newText + value.substring(selEnd === -1 ? value.length : selEnd);
        textarea.selectionStart = selStart;
        textarea.selectionEnd = selStart + newText.length;
      }
    }
  } else if (e.key === 'Enter') {
    e.preventDefault();

    // Check if previous line ends with a colon (Python block start)
    const trimmedLine = currentLine.trimEnd();
    const endsWithColon = trimmedLine.endsWith(':');

    // Calculate new indentation
    let newIndent = currentIndent;
    if (endsWithColon) {
      // Increase indent after colon
      newIndent = currentIndent + INDENT;
    }

    // Check for dedent keywords at start of line
    const dedentKeywords = ['return', 'break', 'continue', 'pass', 'raise'];
    const lineContent = currentLine.trim();
    const shouldDedent = dedentKeywords.some(kw =>
      lineContent === kw || lineContent.startsWith(kw + ' ')
    );

    // Insert newline with appropriate indentation
    textarea.value = value.substring(0, start) + '\n' + newIndent + value.substring(end);
    textarea.selectionStart = textarea.selectionEnd = start + 1 + newIndent.length;
  } else if (e.key === 'Backspace' && start === end) {
    // Smart backspace: delete entire indent if cursor is in leading whitespace
    const beforeCursor = value.substring(lineStart, start);
    if (beforeCursor.length > 0 && beforeCursor.trim() === '') {
      // Cursor is in leading whitespace
      const spacesToDelete = beforeCursor.length % INDENT_SIZE || INDENT_SIZE;
      if (spacesToDelete > 0 && beforeCursor.length >= spacesToDelete) {
        e.preventDefault();
        textarea.value = value.substring(0, start - spacesToDelete) + value.substring(start);
        textarea.selectionStart = textarea.selectionEnd = start - spacesToDelete;
      }
    }
  }
}

function setupSandboxControls() {
  const toggleBtn = document.getElementById('sandbox-toggle');
  const runBtn = document.getElementById('sandbox-run');
  const clearBtn = document.getElementById('sandbox-clear');
  const stepBtn = document.getElementById('sandbox-step');
  const stepBackBtn = document.getElementById('sandbox-step-back');
  const playBtn = document.getElementById('sandbox-play');
  const resetBtn = document.getElementById('sandbox-reset');
  const speedSlider = document.getElementById('sandbox-speed');
  const codeTextarea = document.getElementById('sandbox-code');

  // Setup Python code editor indentation handling
  if (codeTextarea) {
    codeTextarea.addEventListener('keydown', handlePythonIndentation);
  }

  if (toggleBtn) {
    toggleBtn.onclick = toggleSandboxMode;
  }

  if (runBtn) {
    runBtn.onclick = runSandboxCode;
  }

  if (clearBtn) {
    clearBtn.onclick = clearSandbox;
  }

  if (stepBtn) {
    stepBtn.onclick = () => {
      if (sandboxState.animator) sandboxState.animator.stepForward();
    };
  }

  if (stepBackBtn) {
    stepBackBtn.onclick = () => {
      if (sandboxState.animator) sandboxState.animator.stepBack();
    };
  }

  if (playBtn) {
    playBtn.onclick = () => {
      if (sandboxState.animator) {
        const playing = sandboxState.animator.togglePlay();
        playBtn.textContent = playing ? 'Pause' : 'Play';
      }
    };
  }

  if (resetBtn) {
    resetBtn.onclick = () => {
      if (sandboxState.animator) {
        sandboxState.animator.reset();
        if (playBtn) playBtn.textContent = 'Play';
      }
    };
  }

  if (speedSlider) {
    speedSlider.oninput = (e) => {
      if (sandboxState.animator) {
        sandboxState.animator.setSpeed(parseInt(e.target.value, 10));
      }
    };
  }
}

// Register animations
animationRegistry['binary-search'] = BinarySearchAnimator;
animationRegistry['selection-sort'] = SelectionSortAnimator;
animationRegistry['bfs'] = BFSAnimator;
animationRegistry['call-stack'] = CallStackAnimator;
animationRegistry['dfs'] = DFSAnimator;
animationRegistry['linked-list'] = LinkedListAnimator;
animationRegistry['two-pointers'] = TwoPointersAnimator;
animationRegistry['quicksort'] = QuicksortAnimator;
animationRegistry['hash-table'] = HashTableAnimator;
animationRegistry['dijkstra'] = DijkstraAnimator;
animationRegistry['memoization'] = MemoizationAnimator;
animationRegistry['topsort'] = TopSortAnimator;
animationRegistry['string-match'] = StringMatchAnimator;
animationRegistry['backtracking'] = BacktrackingAnimator;
animationRegistry['permutations'] = PermutationsAnimator;
animationRegistry['iterative-dp'] = IterativeDPAnimator;
animationRegistry['rotated-array'] = RotatedArrayAnimator;
animationRegistry['merge-sort'] = MergeSortAnimator;
animationRegistry['sliding-window'] = SlidingWindowAnimator;
animationRegistry['binary-tree'] = BinaryTreeAnimator;
animationRegistry['heap'] = HeapAnimator;
animationRegistry['greedy'] = GreedyAnimator;
animationRegistry['knapsack'] = KnapsackAnimator;
animationRegistry['min-stack'] = MinStackAnimator;
animationRegistry['lru-cache'] = LRUCacheAnimator;
animationRegistry['trie-ops'] = TrieAnimator;
animationRegistry['car-fleet'] = CarFleetAnimator;
animationRegistry['median-two-arrays'] = MedianTwoArraysAnimator;
animationRegistry['kth-largest'] = KthLargestAnimator;
animationRegistry['n-queens'] = NQueensAnimator;

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

// Initialize sandbox controls
setupSandboxControls();

loadChapters().catch((err) => {
  console.error(err);
  alert("Failed to load chapters");
});

// Practice Mode State and Functions
const practiceState = {
  problems: [],
  categories: [],
  currentProblem: null,
  progress: {},
  hintsRevealed: 0,
  currentTestCase: 0,
  practiceAnimator: null,
};

function loadProgress() {
  try {
    const saved = localStorage.getItem('dsatutor_practice_progress');
    if (saved) {
      practiceState.progress = JSON.parse(saved);
    }
  } catch (e) {
    console.error('Failed to load progress:', e);
  }
}

function saveProgress() {
  try {
    localStorage.setItem('dsatutor_practice_progress', JSON.stringify(practiceState.progress));
  } catch (e) {
    console.error('Failed to save progress:', e);
  }
}

function updateProblemProgress(problemId, data) {
  if (!practiceState.progress[problemId]) {
    practiceState.progress[problemId] = {
      attempts: 0,
      solved: false,
      hintsUsed: 0,
      lastAttempt: null,
    };
  }
  Object.assign(practiceState.progress[problemId], data);
  saveProgress();
}

async function loadProblems() {
  try {
    const response = await fetch('/api/practice/problems');
    if (!response.ok) throw new Error('Failed to load problems');
    const data = await response.json();
    // API returns {problems: [...], total: N, stats: {...}}
    practiceState.problems = data.problems || [];
    practiceState.stats = data.stats || {};
    return practiceState.problems;
  } catch (e) {
    console.error('Failed to load problems:', e);
    return [];
  }
}

async function loadCategories() {
  try {
    const response = await fetch('/api/practice/categories');
    if (!response.ok) throw new Error('Failed to load categories');
    practiceState.categories = await response.json();
    return practiceState.categories;
  } catch (e) {
    console.error('Failed to load categories:', e);
    return [];
  }
}

async function loadProblem(id) {
  try {
    const response = await fetch(`/api/practice/problems/${id}`);
    if (!response.ok) throw new Error('Failed to load problem');
    return await response.json();
  } catch (e) {
    console.error('Failed to load problem:', e);
    return null;
  }
}

function showPracticeList() {
  el('practice-list').classList.remove('hidden');
  el('practice-mode').classList.add('hidden');
}

function showProblem() {
  el('practice-list').classList.add('hidden');
  el('practice-mode').classList.remove('hidden');
}

function renderProblemsList() {
  const container = el('problems-by-category');
  if (!container) return;

  const problems = practiceState.problems;
  const categories = practiceState.categories;
  const progress = practiceState.progress;

  // Group problems by category
  const byCategory = {};
  problems.forEach(p => {
    if (!byCategory[p.category]) {
      byCategory[p.category] = [];
    }
    byCategory[p.category].push(p);
  });

  let html = '';
  categories.forEach(cat => {
    const catProblems = byCategory[cat.id] || [];
    if (catProblems.length === 0) return;

    const solved = catProblems.filter(p => progress[p.id]?.solved).length;

    html += `
      <div class="category-section">
        <div class="category-header">
          <h3>${escapeHtml(cat.name)}</h3>
          <span class="category-progress">${solved}/${catProblems.length}</span>
        </div>
        <div class="category-problems">
    `;

    catProblems.forEach(p => {
      const prog = progress[p.id] || {};
      const statusClass = prog.solved ? 'solved' : (prog.attempts > 0 ? 'attempted' : '');
      const diffClass = `difficulty-${p.difficulty.toLowerCase()}`;

      html += `
        <div class="problem-card ${statusClass}" data-problem-id="${p.id}">
          <div class="problem-card-header">
            <span class="problem-number">${p.number}.</span>
            <span class="problem-title">${escapeHtml(p.title)}</span>
          </div>
          <div class="problem-card-meta">
            <span class="difficulty-badge ${diffClass}">${p.difficulty}</span>
            ${prog.solved ? '<span class="solved-badge">Solved</span>' : ''}
          </div>
        </div>
      `;
    });

    html += '</div></div>';
  });

  container.innerHTML = html;

  // Add click handlers
  container.querySelectorAll('.problem-card').forEach(card => {
    card.onclick = () => {
      const problemId = card.dataset.problemId;
      openProblem(problemId);
    };
  });
}

async function openProblem(problemId) {
  const problem = await loadProblem(problemId);
  if (!problem) {
    alert('Failed to load problem');
    return;
  }

  practiceState.currentProblem = problem;
  practiceState.hintsRevealed = 0;
  practiceState.currentTestCase = 0;

  renderProblemDescription();
  renderTestCases();
  renderHints();
  renderStarterCode();

  showProblem();
}

function renderProblemDescription() {
  const problem = practiceState.currentProblem;
  if (!problem) return;

  // Update header
  const titleEl = el('practice-title');
  const diffBadge = el('practice-difficulty');

  if (titleEl) {
    titleEl.textContent = `${problem.number}. ${problem.title}`;
  }

  if (diffBadge) {
    diffBadge.textContent = problem.difficulty.toUpperCase();
    diffBadge.className = `chapter-badge practice-badge difficulty-${problem.difficulty.toLowerCase()}`;
  }

  // Update problem statement
  const statementEl = el('problem-statement');
  if (statementEl) {
    statementEl.innerHTML = `<div class="problem-text">${problem.description}</div>`;
  }

  // Update examples
  const examplesEl = el('problem-examples');
  if (examplesEl && problem.examples && problem.examples.length > 0) {
    let html = '<h4>Examples:</h4>';
    problem.examples.forEach((ex, i) => {
      html += `
        <div class="example">
          <div class="example-header">Example ${i + 1}:</div>
          <div class="example-io">
            <div><strong>Input:</strong> <code>${escapeHtml(ex.input)}</code></div>
            <div><strong>Output:</strong> <code>${escapeHtml(ex.output)}</code></div>
          </div>
          ${ex.explanation ? `<div class="example-explanation"><strong>Explanation:</strong> ${escapeHtml(ex.explanation)}</div>` : ''}
        </div>
      `;
    });
    examplesEl.innerHTML = html;
  }

  // Update constraints
  const constraintsEl = el('problem-constraints');
  if (constraintsEl && problem.constraints && problem.constraints.length > 0) {
    let html = '<h4>Constraints:</h4><ul>';
    problem.constraints.forEach(c => {
      html += `<li><code>${escapeHtml(c)}</code></li>`;
    });
    html += '</ul>';
    constraintsEl.innerHTML = html;
  }

  // Update related chapters
  const relatedEl = el('related-chapters');
  if (relatedEl && problem.relatedChapters && problem.relatedChapters.length > 0) {
    let html = '';
    problem.relatedChapters.forEach(chNum => {
      const ch = state.chapters.find(c => c.number === chNum);
      if (ch) {
        html += `<a href="#" class="chapter-link" data-chapter="${chNum}">${ch.title}</a>`;
      }
    });
    relatedEl.innerHTML = html;

    // Add click handlers
    relatedEl.querySelectorAll('.chapter-link').forEach(link => {
      link.onclick = (e) => {
        e.preventDefault();
        const chNum = parseInt(link.dataset.chapter);
        const ch = state.chapters.find(c => c.number === chNum);
        if (ch) {
          el('practice-toggle').classList.remove('active');
          el('practice-list').classList.add('hidden');
          el('practice-mode').classList.add('hidden');
          el('detail').classList.remove('hidden');
          el('detail-empty').classList.add('hidden');
          selectChapter(ch);
        }
      };
    });
  }
}

// Format a value for display - handles strings, objects, arrays
function formatTestValue(value) {
  if (value === null || value === undefined) return 'null';
  if (typeof value === 'string') return value;
  if (typeof value === 'object') {
    // For objects like {nums: [1,2,3], target: 5}, format nicely
    if (Array.isArray(value)) {
      return JSON.stringify(value);
    }
    // Format as key=value pairs
    return Object.entries(value)
      .map(([k, v]) => `${k} = ${JSON.stringify(v)}`)
      .join(', ');
  }
  return String(value);
}

function renderTestCases() {
  const problem = practiceState.currentProblem;
  const selectorEl = el('test-case-selector');
  const resultsEl = el('test-results');
  if (!selectorEl || !problem) return;

  // Render test case selector buttons
  let selectorHtml = '';
  problem.testCases?.forEach((tc, i) => {
    if (!tc.hidden) {
      const isActive = i === practiceState.currentTestCase;
      selectorHtml += `<button class="test-case-btn ${isActive ? 'active' : ''}" data-index="${i}">Case ${i + 1}</button>`;
    }
  });
  selectorEl.innerHTML = selectorHtml;

  // Render current test case details
  const currentTC = problem.testCases?.[practiceState.currentTestCase];
  if (resultsEl && currentTC) {
    const inputDisplay = escapeHtml(formatTestValue(currentTC.input));
    const expectedDisplay = escapeHtml(formatTestValue(currentTC.expected));
    resultsEl.innerHTML = `
      <div class="test-case-detail">
        <div class="test-io">
          <div class="test-input">
            <strong>Input:</strong>
            <pre><code>${inputDisplay}</code></pre>
          </div>
          <div class="test-expected">
            <strong>Expected Output:</strong>
            <pre><code>${expectedDisplay}</code></pre>
          </div>
        </div>
        <div class="test-output" id="current-test-output"></div>
      </div>
    `;
  }

  // Add click handlers to test case buttons
  selectorEl.querySelectorAll('.test-case-btn').forEach(btn => {
    btn.onclick = () => {
      practiceState.currentTestCase = parseInt(btn.dataset.index);
      selectorEl.querySelectorAll('.test-case-btn').forEach(b => b.classList.remove('active'));
      btn.classList.add('active');
      renderTestCases();
    };
  });
}

function renderHints() {
  const problem = practiceState.currentProblem;
  const hintsListEl = el('hints-list');
  const revealBtn = el('reveal-hint');
  if (!hintsListEl || !problem) return;

  let html = '';
  problem.hints?.forEach((hint, i) => {
    const revealed = i < practiceState.hintsRevealed;
    if (revealed) {
      html += `
        <div class="hint revealed">
          <div class="hint-header">
            <span class="hint-type">${escapeHtml(hint.type || 'Hint')} ${i + 1}</span>
          </div>
          <div class="hint-content">${escapeHtml(hint.content)}</div>
        </div>
      `;
    }
  });

  if (practiceState.hintsRevealed === 0) {
    html = '<p class="no-hints">No hints revealed yet. Click "Reveal Next Hint" to get started.</p>';
  }

  hintsListEl.innerHTML = html;

  // Update reveal button
  if (revealBtn) {
    const remainingHints = (problem.hints?.length || 0) - practiceState.hintsRevealed;
    if (remainingHints > 0) {
      revealBtn.textContent = `Reveal Next Hint (${remainingHints} remaining)`;
      revealBtn.disabled = false;
      revealBtn.onclick = () => {
        practiceState.hintsRevealed++;
        updateProblemProgress(problem.id, { hintsUsed: practiceState.hintsRevealed });
        renderHints();
      };
    } else {
      revealBtn.textContent = 'All hints revealed';
      revealBtn.disabled = true;
    }
  }
}

function renderStarterCode() {
  const problem = practiceState.currentProblem;
  const editor = el('practice-code');
  if (!editor || !problem) return;

  editor.value = problem.starterCode || '# Write your solution here\n';
}

async function runPracticeCode() {
  const problem = practiceState.currentProblem;
  const editor = el('practice-code');
  const outputEl = el('current-test-output');
  if (!problem || !editor) return;

  const code = editor.value;
  const testCase = problem.testCases?.[practiceState.currentTestCase];
  if (!testCase) return;

  if (outputEl) {
    outputEl.innerHTML = '<div class="running">Running...</div>';
  }

  try {
    const response = await fetch('/api/practice/run', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        problemId: problem.id,
        code: code,
        testCaseIndex: practiceState.currentTestCase,
      }),
    });

    const result = await response.json();

    let html = '';
    if (result.error) {
      html = `<div class="test-error">${escapeHtml(result.error)}</div>`;
    } else {
      html = `
        <div class="test-result-detail ${result.passed ? 'passed' : 'failed'}">
          <div class="result-status">${result.passed ? 'Passed' : 'Failed'}</div>
          <div class="result-output"><strong>Your Output:</strong> <pre><code>${escapeHtml(result.output || 'None')}</code></pre></div>
        </div>
      `;

      // Update test case button status
      const btn = el('test-case-selector')?.querySelector(`[data-index="${practiceState.currentTestCase}"]`);
      if (btn) {
        btn.classList.remove('passed', 'failed');
        btn.classList.add(result.passed ? 'passed' : 'failed');
      }

      // Handle visualization steps
      if (result.steps && result.steps.length > 0) {
        renderPracticeVisualization(result.steps);
      }
    }

    if (outputEl) {
      outputEl.innerHTML = html;
    }
  } catch (e) {
    if (outputEl) {
      outputEl.innerHTML = `<div class="test-error">Error: ${escapeHtml(e.message)}</div>`;
    }
  }
}

async function submitPracticeCode() {
  const problem = practiceState.currentProblem;
  const editor = el('practice-code');
  const outputEl = el('test-results');
  if (!problem || !editor) return;

  const code = editor.value;

  if (outputEl) {
    outputEl.innerHTML = '<div class="running">Submitting all test cases...</div>';
  }

  try {
    const response = await fetch('/api/practice/submit', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        problemId: problem.id,
        code: code,
      }),
    });

    const result = await response.json();

    let html = '';
    if (result.error) {
      html = `<div class="test-error">${escapeHtml(result.error)}</div>`;
    } else {
      const allPassed = result.results?.every(r => r.passed);

      html = `
        <div class="submit-result ${allPassed ? 'all-passed' : 'some-failed'}">
          <div class="submit-status">${allPassed ? 'All Tests Passed!' : 'Some Tests Failed'}</div>
          <div class="submit-summary">Passed: ${result.passed}/${result.total}</div>
          <div class="submit-details">
      `;

      result.results?.forEach((r, i) => {
        html += `
          <div class="submit-test ${r.passed ? 'passed' : 'failed'}">
            <span class="test-num">Test ${i + 1}</span>
            <span class="test-status">${r.passed ? 'Passed' : 'Failed'}</span>
            ${!r.passed && r.output ? `<div class="test-output">Got: <code>${escapeHtml(r.output)}</code></div>` : ''}
          </div>
        `;
      });

      html += '</div></div>';

      // Update progress
      updateProblemProgress(problem.id, {
        attempts: (practiceState.progress[problem.id]?.attempts || 0) + 1,
        solved: allPassed || practiceState.progress[problem.id]?.solved,
        lastAttempt: new Date().toISOString(),
      });

      // Refresh problems list to show updated status
      if (allPassed) {
        await loadProblems();
        renderProblemsList();
        updateProgressCount();
      }
    }

    if (outputEl) {
      outputEl.innerHTML = html;
    }
  } catch (e) {
    if (outputEl) {
      outputEl.innerHTML = `<div class="test-error">Error: ${escapeHtml(e.message)}</div>`;
    }
  }
}

function updateProgressCount() {
  const countEl = el('progress-count');
  if (!countEl) return;

  const solved = Object.values(practiceState.progress).filter(p => p.solved).length;
  const total = practiceState.problems.length;
  countEl.textContent = `${solved}/${total} Solved`;
}

function renderPracticeVisualization(steps) {
  const vizContainer = el('practice-canvas');
  if (!vizContainer || !steps || steps.length === 0) return;

  practiceState.vizSteps = steps;
  practiceState.currentVizStep = 0;

  renderCurrentVizStep();
}

function renderCurrentVizStep() {
  const vizContainer = el('practice-canvas');
  const varsContainer = el('practice-vars');
  const stateContainer = el('practice-state');
  const stepCounter = el('practice-step-counter');

  if (!vizContainer || !practiceState.vizSteps) return;

  const steps = practiceState.vizSteps;
  const step = steps[practiceState.currentVizStep];

  if (!step) return;

  // Update step counter
  if (stepCounter) {
    stepCounter.textContent = `Step ${practiceState.currentVizStep + 1}/${steps.length}`;
  }

  // Render variables
  if (varsContainer && step.locals) {
    let varsHtml = '<div class="vars-list">';
    Object.entries(step.locals).forEach(([name, value]) => {
      const displayValue = typeof value === 'object' ? JSON.stringify(value) : String(value);
      varsHtml += `
        <div class="var-item">
          <span class="var-name">${escapeHtml(name)}</span>
          <span class="var-value">${escapeHtml(displayValue)}</span>
        </div>
      `;
    });
    varsHtml += '</div>';
    varsContainer.innerHTML = varsHtml;
  }

  // Update state text
  if (stateContainer) {
    stateContainer.textContent = `Line ${step.line}: ${step.function || 'main'}`;
  }

  // Render visualization based on detected structures
  if (step.structures && step.structures.length > 0) {
    let vizHtml = '';
    step.structures.forEach(struct => {
      vizHtml += renderStructure(struct);
    });
    vizContainer.innerHTML = vizHtml;
  } else {
    // Simple variable display
    vizContainer.innerHTML = '<div class="simple-viz">See variables panel for current state</div>';
  }
}

function renderStructure(struct) {
  switch (struct.type) {
    case 'array':
      return renderArrayViz(struct);
    case 'linked_list':
      return renderLinkedListViz(struct);
    case 'tree':
      return renderTreeViz(struct);
    default:
      return `<div class="struct-unknown">${struct.name}: ${JSON.stringify(struct.data)}</div>`;
  }
}

function renderArrayViz(struct) {
  const data = struct.data || [];
  const highlights = struct.highlights || {};

  let html = `<div class="array-viz"><div class="struct-name">${escapeHtml(struct.name)}</div><div class="array-cells">`;
  data.forEach((val, i) => {
    let cellClass = 'array-cell';
    if (highlights.current === i) cellClass += ' current';
    if (highlights.comparing?.includes(i)) cellClass += ' comparing';
    if (highlights.found === i) cellClass += ' found';
    html += `<div class="${cellClass}"><span class="cell-value">${val}</span><span class="cell-index">${i}</span></div>`;
  });
  html += '</div></div>';
  return html;
}

function renderLinkedListViz(struct) {
  const nodes = struct.data || [];
  let html = `<div class="linked-list-viz"><div class="struct-name">${escapeHtml(struct.name)}</div><div class="list-nodes">`;
  nodes.forEach((val, i) => {
    html += `<div class="list-node"><span class="node-value">${val}</span></div>`;
    if (i < nodes.length - 1) {
      html += '<div class="list-arrow">-></div>';
    }
  });
  html += '</div></div>';
  return html;
}

function renderTreeViz(struct) {
  // Simple tree rendering - could be enhanced
  return `<div class="tree-viz"><div class="struct-name">${escapeHtml(struct.name)}</div><pre>${JSON.stringify(struct.data, null, 2)}</pre></div>`;
}

function showSolution() {
  const problem = practiceState.currentProblem;
  if (!problem || !problem.solution) return;

  const lockedEl = el('solution-locked');
  const contentEl = el('solution-content');
  const walkthroughEl = el('solution-walkthrough');

  if (lockedEl) {
    lockedEl.classList.add('hidden');
  }

  if (contentEl) {
    let html = `
      <h4>Solution</h4>
      <div class="solution-approach"><strong>Approach:</strong> ${escapeHtml(problem.solution.approach || '')}</div>
      <div class="solution-complexity">
        <span><strong>Time:</strong> ${escapeHtml(problem.timeComplexity || 'N/A')}</span>
        <span><strong>Space:</strong> ${escapeHtml(problem.spaceComplexity || 'N/A')}</span>
      </div>
      <pre class="solution-code"><code>${escapeHtml(problem.solution.code || '')}</code></pre>
    `;
    contentEl.innerHTML = html;
    contentEl.classList.remove('hidden');
  }

  if (walkthroughEl && problem.solution.walkthrough && problem.solution.walkthrough.length > 0) {
    let html = '<h4>Walkthrough</h4>';
    problem.solution.walkthrough.forEach((step, i) => {
      html += `
        <div class="walkthrough-step">
          <div class="walkthrough-header">Step ${i + 1}: ${escapeHtml(step.title || '')}</div>
          <div class="walkthrough-explanation">${escapeHtml(step.explanation || '')}</div>
          ${step.code ? `<pre class="walkthrough-code"><code>${escapeHtml(step.code)}</code></pre>` : ''}
        </div>
      `;
    });
    walkthroughEl.innerHTML = html;
  }
}

function setupPracticeControls() {
  // Practice toggle button
  const toggleBtn = el('practice-toggle');
  if (toggleBtn) {
    toggleBtn.onclick = async () => {
      const isActive = toggleBtn.classList.toggle('active');
      const chaptersEl = el('chapters');
      const detailEl = el('detail');
      const detailEmptyEl = el('detail-empty');
      const sandboxEl = el('sandbox-mode');
      const practiceListEl = el('practice-list');
      const practiceModeEl = el('practice-mode');

      if (isActive) {
        // Entering practice mode - hide chapter content, show practice
        if (chaptersEl) chaptersEl.classList.add('hidden');
        if (detailEl) detailEl.classList.add('hidden');
        if (detailEmptyEl) detailEmptyEl.classList.add('hidden');
        if (sandboxEl) sandboxEl.classList.add('hidden');

        if (practiceState.problems.length === 0) {
          await loadProblems();
          await loadCategories();
        }
        renderProblemsList();
        updateProgressCount();
        showPracticeList();
      } else {
        // Exiting practice mode - show chapter content, hide practice
        if (practiceListEl) practiceListEl.classList.add('hidden');
        if (practiceModeEl) practiceModeEl.classList.add('hidden');

        if (chaptersEl) chaptersEl.classList.remove('hidden');
        // Show detail-empty or detail depending on whether a chapter is selected
        if (state.currentChapter) {
          if (detailEl) detailEl.classList.remove('hidden');
          if (detailEmptyEl) detailEmptyEl.classList.add('hidden');
        } else {
          if (detailEmptyEl) detailEmptyEl.classList.remove('hidden');
          if (detailEl) detailEl.classList.add('hidden');
        }
      }
    };
  }

  // Back button
  const backBtn = el('practice-back');
  if (backBtn) {
    backBtn.onclick = () => {
      showPracticeList();
    };
  }

  // Run button
  const runBtn = el('practice-run');
  if (runBtn) {
    runBtn.onclick = runPracticeCode;
  }

  // Submit button
  const submitBtn = el('practice-submit');
  if (submitBtn) {
    submitBtn.onclick = submitPracticeCode;
  }

  // Solution button
  const solutionBtn = el('unlock-solution');
  if (solutionBtn) {
    solutionBtn.onclick = showSolution;
  }

  // Panel tabs (Description/Hints/Solution and TestCases/Visualization)
  const panelTabs = document.querySelectorAll('#practice-mode .tab-btn');
  panelTabs.forEach(tab => {
    tab.onclick = () => {
      const tabContainer = tab.closest('.panel-tabs');
      const panelBody = tab.closest('.panel-header')?.nextElementSibling;
      if (!tabContainer || !panelBody) return;

      // Update active tab button
      tabContainer.querySelectorAll('.tab-btn').forEach(t => t.classList.remove('active'));
      tab.classList.add('active');

      // Show/hide tab content based on data-tab attribute
      const targetTab = tab.dataset.tab;
      panelBody.querySelectorAll('.tab-content').forEach(content => {
        // Map content IDs to tab names:
        // test-cases-tab -> testcases, visualization-tab -> visualization
        // problem-description -> description, problem-hints -> hints, problem-solution -> solution
        const contentId = content.id;
        let contentTab = contentId
          .replace('-tab', '')
          .replace('problem-', '')
          .replace(/-/g, ''); // Remove all hyphens: test-cases -> testcases

        const isMatch = contentTab === targetTab;
        content.classList.toggle('hidden', !isMatch);
        content.classList.toggle('active', isMatch);
      });
    };
  });

  // Filter controls
  const difficultyFilter = el('filter-difficulty');
  if (difficultyFilter) {
    difficultyFilter.onchange = () => filterProblems();
  }

  const categoryFilter = el('filter-category');
  if (categoryFilter) {
    categoryFilter.onchange = () => filterProblems();
  }

  const statusFilter = el('filter-status');
  if (statusFilter) {
    statusFilter.onchange = () => filterProblems();
  }

  // Search filter
  const searchFilter = el('filter-search');
  if (searchFilter) {
    searchFilter.oninput = () => filterProblems();
  }

  // Practice code editor - add Python indentation handling
  const practiceCodeEditor = el('practice-code');
  if (practiceCodeEditor) {
    practiceCodeEditor.addEventListener('keydown', handlePythonIndentation);
  }

  // Visualization controls
  setupPracticeVizControls();
}

function filterProblems() {
  const difficulty = el('filter-difficulty')?.value || 'all';
  const category = el('filter-category')?.value || 'all';
  const status = el('filter-status')?.value || 'all';
  const search = el('filter-search')?.value?.toLowerCase() || '';

  const cards = document.querySelectorAll('.problem-card');
  cards.forEach(card => {
    const problem = practiceState.problems.find(p => p.id === card.dataset.problemId);
    if (!problem) return;

    const progress = practiceState.progress[problem.id] || {};

    let show = true;

    if (difficulty !== 'all' && problem.difficulty.toLowerCase() !== difficulty.toLowerCase()) {
      show = false;
    }

    if (category !== 'all' && problem.category !== category) {
      show = false;
    }

    if (status === 'solved' && !progress.solved) {
      show = false;
    } else if (status === 'unsolved' && progress.solved) {
      show = false;
    } else if (status === 'attempted' && (!progress.attempts || progress.solved)) {
      show = false;
    }

    // Search filter
    if (search && !problem.title.toLowerCase().includes(search) &&
        !problem.id.toLowerCase().includes(search)) {
      show = false;
    }

    card.style.display = show ? '' : 'none';
  });

  // Hide empty categories
  document.querySelectorAll('.category-section').forEach(section => {
    const visibleCards = section.querySelectorAll('.problem-card:not([style*="display: none"])');
    section.style.display = visibleCards.length > 0 ? '' : 'none';
  });
}

function exportProgress() {
  const data = JSON.stringify(practiceState.progress, null, 2);
  const blob = new Blob([data], { type: 'application/json' });
  const url = URL.createObjectURL(blob);
  const a = document.createElement('a');
  a.href = url;
  a.download = 'dsatutor_progress.json';
  a.click();
  URL.revokeObjectURL(url);
}

function importProgress(file) {
  const reader = new FileReader();
  reader.onload = (e) => {
    try {
      const data = JSON.parse(e.target.result);
      practiceState.progress = data;
      saveProgress();
      renderProblemsList();
    } catch (err) {
      alert('Failed to import progress: Invalid file format');
    }
  };
  reader.readAsText(file);
}

function setupPracticeVizControls() {
  // Practice visualization controls
  const resetBtn = el('practice-reset');
  const stepBackBtn = el('practice-step-back');
  const playBtn = el('practice-play');
  const stepBtn = el('practice-step');

  if (resetBtn) {
    resetBtn.onclick = () => {
      if (practiceState.vizSteps && practiceState.vizSteps.length > 0) {
        practiceState.currentVizStep = 0;
        renderCurrentVizStep();
      }
    };
  }

  if (stepBackBtn) {
    stepBackBtn.onclick = () => {
      if (practiceState.vizSteps && practiceState.currentVizStep > 0) {
        practiceState.currentVizStep--;
        renderCurrentVizStep();
      }
    };
  }

  if (stepBtn) {
    stepBtn.onclick = () => {
      if (practiceState.vizSteps && practiceState.currentVizStep < practiceState.vizSteps.length - 1) {
        practiceState.currentVizStep++;
        renderCurrentVizStep();
      }
    };
  }

  if (playBtn) {
    playBtn.onclick = () => {
      if (practiceState.vizPlaying) {
        // Stop playback
        practiceState.vizPlaying = false;
        playBtn.textContent = 'Play';
        if (practiceState.vizTimer) {
          clearInterval(practiceState.vizTimer);
          practiceState.vizTimer = null;
        }
      } else {
        // Start playback
        practiceState.vizPlaying = true;
        playBtn.textContent = 'Pause';
        practiceState.vizTimer = setInterval(() => {
          if (practiceState.currentVizStep < practiceState.vizSteps.length - 1) {
            practiceState.currentVizStep++;
            renderCurrentVizStep();
          } else {
            // Reached the end
            practiceState.vizPlaying = false;
            playBtn.textContent = 'Play';
            clearInterval(practiceState.vizTimer);
            practiceState.vizTimer = null;
          }
        }, 500);
      }
    };
  }
}

// Initialize practice mode
loadProgress();
setupPracticeControls();
