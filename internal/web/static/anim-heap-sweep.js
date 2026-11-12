// Sweep line with a heap: Car Pooling, animated.
// A sweep line walks the pick-ups from left to right while a min-heap of
// (drop_off, passengers) holds the trips currently in the car. Each step
// shows the timeline, the heap as a tree and as a list, and the load meter.
// Loaded after app.js, which defines AlgorithmAnimator and animationRegistry.

class HeapSweepAnimator extends AlgorithmAnimator {
  constructor(canvasEl, codeEl, stateEl, config) {
    super(canvasEl, codeEl, stateEl, config);
    // [passengers, pick_up, drop_off], deliberately unsorted.
    this.input = [[3, 3, 7], [2, 5, 8], [2, 1, 5], [1, 7, 9], [1, 4, 6]];
    this.capacity = 6;
    this.maxKm = 10;
    // Sorted by pick-up and named A, B, C... in that order.
    this.trips = this.input
      .map((t) => t.slice())
      .sort((a, b) => a[1] - b[1])
      .map(([passengers, pickUp, dropOff], i) => ({
        name: String.fromCharCode(65 + i), passengers, pickUp, dropOff,
      }));
    this.view = null;

    this.setCode(`import heapq

def carPooling(trips, capacity):
    trips = sorted(trips, key=lambda trip: trip[1])
    riding = []
    load = 0
    for passengers, pick_up, drop_off in trips:
        while riding and riding[0][0] <= pick_up:
            load -= heapq.heappop(riding)[1]
        heapq.heappush(riding, (drop_off, passengers))
        load += passengers
        if load > capacity:
            return False
    return True`);
  }

  // heapq on entries {key: [drop_off, passengers], trip}, compared as tuples.
  static less(a, b) {
    return a.key[0] !== b.key[0] ? a.key[0] < b.key[0] : a.key[1] < b.key[1];
  }

  static push(heap, item) {
    heap.push(item);
    let i = heap.length - 1;
    while (i > 0) {
      const parent = (i - 1) >> 1;
      if (!HeapSweepAnimator.less(heap[i], heap[parent])) break;
      [heap[i], heap[parent]] = [heap[parent], heap[i]];
      i = parent;
    }
  }

  static pop(heap) {
    const top = heap[0];
    const last = heap.pop();
    if (heap.length > 0) {
      heap[0] = last;
      let i = 0;
      for (;;) {
        const l = 2 * i + 1;
        const r = l + 1;
        let small = i;
        if (l < heap.length && HeapSweepAnimator.less(heap[l], heap[small])) small = l;
        if (r < heap.length && HeapSweepAnimator.less(heap[r], heap[small])) small = r;
        if (small === i) break;
        [heap[i], heap[small]] = [heap[small], heap[i]];
        i = small;
      }
    }
    return top;
  }

  static fmt(entry) {
    return `(${entry.key[0]},${entry.key[1]})`;
  }

  buildSteps() {
    this.steps = [];
    const heap = [];
    const status = {};          // trip name -> waiting | riding | done
    this.trips.forEach((t) => { status[t.name] = 'waiting'; });
    let load = 0;
    let km = null;
    let currentTrip = null;
    const fmtHeap = () => `[${heap.map(HeapSweepAnimator.fmt).join(', ')}]`;

    const add = (lineNums, state, extra = {}) => {
      const snap = {
        lineNums,
        state,
        heap: heap.map((e) => ({ key: e.key.slice(), trip: e.trip })),
        status: { ...status },
        load,
        km,
        current: currentTrip,
        phase: 'idle',
        ...extra,
      };
      snap.apply = () => { this.view = snap; };
      this.steps.push(snap);
    };

    add([4, 5, 6],
      `Sort the trips by pick-up: ${this.trips.map((t) => `${t.name}=[${t.passengers},${t.pickUp},${t.dropOff}]`).join(' ')}. riding = [], load = 0.`,
      { phase: 'sort' });

    for (const trip of this.trips) {
      km = trip.pickUp;
      currentTrip = trip.name;
      add([7], `Sweep to km ${km}: trip ${trip.name} wants to board ${trip.passengers} passenger${trip.passengers > 1 ? 's' : ''} (km ${trip.pickUp} to ${trip.dropOff}).`,
        { phase: 'sweep' });

      while (heap.length > 0 && heap[0].key[0] <= km) {
        const top = heap[0];
        add([8], `Top ${HeapSweepAnimator.fmt(top)} is trip ${top.trip}: drop-off ${top.key[0]} <= ${km}, so it has left the car.`,
          { phase: 'expired', focus: top.trip });
        HeapSweepAnimator.pop(heap);
        status[top.trip] = 'done';
        const before = load;
        load -= top.key[1];
        add([9], `Pop ${HeapSweepAnimator.fmt(top)}: free ${top.key[1]} seat${top.key[1] > 1 ? 's' : ''}, load ${before} -> ${load}. riding = ${fmtHeap()}`,
          { phase: 'popped', popped: top });
      }
      if (heap.length === 0) {
        add([8], `Heap is empty: nobody is in the car, nothing to release.`, { phase: 'stop' });
      } else {
        add([8], `Top ${HeapSweepAnimator.fmt(heap[0])} is trip ${heap[0].trip}: drop-off ${heap[0].key[0]} > ${km}, still riding. Stop releasing${heap.length > 1 ? '; everyone below it gets off even later' : ''}.`,
          { phase: 'stop', focus: heap[0].trip });
      }

      HeapSweepAnimator.push(heap, { key: [trip.dropOff, trip.passengers], trip: trip.name });
      status[trip.name] = 'riding';
      const before = load;
      load += trip.passengers;
      add([10, 11], `Push (${trip.dropOff},${trip.passengers}) for trip ${trip.name}, load ${before} -> ${load}. riding = ${fmtHeap()}`,
        { phase: 'pushed', focus: trip.name });

      if (load > this.capacity) {
        add([12, 13], `load ${load} > capacity ${this.capacity}: the car overflows at km ${km}. Return False.`,
          { phase: 'fail' });
        return;
      }
      add([12], `load ${load} <= capacity ${this.capacity}: everyone fits at km ${km}.`, { phase: 'check' });
    }

    currentTrip = null;
    add([14], `Every pick-up fit, and between pick-ups the load only falls. Return True.`, { phase: 'done' });
  }

  // Layout of the timeline SVG, in viewBox units. The viewBox width follows
  // the canvas width (set when the skeleton is built) so text stays ~1:1.
  geom() {
    const width = this.svgWidth || 400;
    const left = 34;
    const right = width - 16;
    const top = 18;
    const lane = 30;
    const unit = (right - left) / this.maxKm;
    const axisY = top + lane * this.trips.length + 8;
    return { width, left, right, top, lane, unit, axisY, height: axisY + 30, x: (km) => left + km * unit };
  }

  buildSkeleton() {
    const avail = (this.canvas.clientWidth || 460) - 60;
    this.svgWidth = Math.max(320, Math.min(720, avail));
    const g = this.geom();
    let svg = `<svg class="hs-timeline" viewBox="0 0 ${g.width} ${g.height}" role="img" aria-label="Trips on a timeline with a sweep line">`;
    for (let km = 0; km <= this.maxKm; km++) {
      const x = g.x(km);
      svg += `<line class="hs-grid" x1="${x}" y1="${g.top - 6}" x2="${x}" y2="${g.axisY}"/>`;
      svg += `<text class="hs-tick" x="${x}" y="${g.axisY + 16}">${km}</text>`;
    }
    svg += `<line class="hs-axis" x1="${g.left}" y1="${g.axisY}" x2="${g.right}" y2="${g.axisY}"/>`;
    this.trips.forEach((t, i) => {
      const y = g.top + i * g.lane;
      const x1 = g.x(t.pickUp);
      const w = g.x(t.dropOff) - x1;
      svg += `<g class="hs-trip" data-trip="${t.name}">`;
      svg += `<text class="hs-trip-name" x="${g.left - 14}" y="${y + 15}">${t.name}</text>`;
      svg += `<rect class="hs-bar" x="${x1}" y="${y + 1}" width="${w}" height="20" rx="5"/>`;
      svg += `<text class="hs-bar-label" x="${x1 + w / 2}" y="${y + 15}">${t.passengers}p, ${t.pickUp}-${t.dropOff}</text>`;
      svg += '</g>';
    });
    svg += `<g class="hs-sweep">`;
    svg += `<line class="hs-sweep-line" x1="0" y1="${g.top - 10}" x2="0" y2="${g.axisY + 2}"/>`;
    svg += `<text class="hs-sweep-label" x="0" y="${g.top - 12}"></text>`;
    svg += '</g>';
    svg += '</svg>';

    this.canvas.innerHTML = `
      <div class="hs-viz">
        <div class="hs-panel hs-panel-timeline">
          <div class="hs-panel-title">Timeline (km)</div>
          ${svg}
          <div class="hs-load">
            <span class="hs-load-label">load</span>
            <div class="hs-load-track">
              <div class="hs-load-fill"></div>
              <div class="hs-load-cap" style="left:${(this.capacity / (this.capacity + 2)) * 100}%"></div>
            </div>
            <span class="hs-load-value"></span>
          </div>
        </div>
        <div class="hs-panel hs-panel-heap">
          <div class="hs-panel-title">riding: min-heap by drop-off</div>
          <div class="hs-heap-body">
            <div class="hs-heap"></div>
            <div class="hs-heap-side">
              <div class="hs-heap-array"></div>
              <div class="hs-event"></div>
            </div>
          </div>
        </div>
      </div>`;
  }

  renderHeap(v) {
    const heap = v.heap;
    if (heap.length === 0) {
      return '<div class="hs-heap-empty">empty: nobody in the car</div>';
    }
    const width = 260;
    const pos = heap.map((_, i) => {
      const level = Math.floor(Math.log2(i + 1));
      const first = 2 ** level - 1;
      const slots = 2 ** level;
      return { x: (width / (slots + 1)) * (i - first + 1), y: 26 + level * 58 };
    });
    const levels = Math.floor(Math.log2(heap.length)) + 1;
    let svg = `<svg class="hs-tree" viewBox="0 0 ${width} ${levels * 58 + 4}" role="img" aria-label="Heap of trips in the car">`;
    for (let i = 1; i < heap.length; i++) {
      const p = pos[(i - 1) >> 1];
      svg += `<line class="hs-edge" x1="${p.x}" y1="${p.y}" x2="${pos[i].x}" y2="${pos[i].y}"/>`;
    }
    heap.forEach((e, i) => {
      const { x, y } = pos[i];
      let cls = 'hs-node';
      if (i === 0) cls += ' top';
      if (e.trip === v.focus) {
        if (v.phase === 'expired') cls += ' expired';
        else if (v.phase === 'pushed') cls += ' pushed';
        else if (v.phase === 'stop') cls += ' active';
      }
      svg += `<g class="${cls}">`;
      svg += `<rect class="hs-node-base" x="${x - 28}" y="${y - 16}" width="56" height="32" rx="8"/>`;
      svg += `<rect class="hs-node-box" x="${x - 28}" y="${y - 16}" width="56" height="32" rx="8"/>`;
      svg += `<text class="hs-node-key" x="${x}" y="${y + 5}">(${e.key[0]},${e.key[1]})</text>`;
      svg += `<text class="hs-node-trip" x="${x}" y="${y - 20}">${i === 0 ? 'top · ' : ''}${e.trip}</text>`;
      svg += '</g>';
    });
    svg += '</svg>';
    return svg;
  }

  render() {
    if (!this.canvas) return;
    const v = this.view || this.steps[this.currentStep];
    if (!v) return;
    if (!this.canvas.querySelector('.hs-viz')) this.buildSkeleton();
    const root = this.canvas.querySelector('.hs-viz');
    const g = this.geom();

    root.dataset.phase = v.phase;
    root.querySelectorAll('.hs-trip').forEach((node) => {
      const name = node.dataset.trip;
      node.setAttribute('class', `hs-trip ${v.status[name]}${name === v.current ? ' current' : ''}${name === v.focus ? ' focus' : ''}`);
    });

    const sweep = root.querySelector('.hs-sweep');
    const shown = v.km !== null;
    sweep.classList.toggle('hs-hidden', !shown);
    if (shown) {
      sweep.style.transform = `translateX(${g.x(v.km)}px)`;
      sweep.querySelector('.hs-sweep-label').textContent = `km ${v.km}`;
    }

    const fill = root.querySelector('.hs-load-fill');
    fill.style.width = `${(Math.min(v.load, this.capacity + 2) / (this.capacity + 2)) * 100}%`;
    fill.classList.toggle('over', v.load > this.capacity);
    fill.classList.toggle('full', v.load === this.capacity);
    root.querySelector('.hs-load-value').textContent = `${v.load} / ${this.capacity}`;

    root.querySelector('.hs-heap').innerHTML = this.renderHeap(v);
    const arr = v.heap.map(HeapSweepAnimator.fmt).join(', ');
    root.querySelector('.hs-heap-array').textContent = `riding = [${arr}]`;

    const event = root.querySelector('.hs-event');
    const messages = {
      sort: 'Trips sorted by pick-up. The sweep starts at the left.',
      sweep: `Sweep line at km ${v.km}: trip ${v.current} is waiting to board.`,
      expired: `Top trip ${v.focus} has reached its drop-off: release it.`,
      popped: v.popped ? `Popped ${HeapSweepAnimator.fmt(v.popped)}: trip ${v.popped.trip} got off.` : '',
      stop: v.heap.length ? `Top trip ${v.focus} is still riding: stop releasing.` : 'Nobody to release.',
      pushed: `Trip ${v.focus} boarded: pushed onto the heap.`,
      check: `load ${v.load} fits in ${this.capacity} seats.`,
      fail: `load ${v.load} is over the capacity of ${this.capacity}.`,
      done: 'All pick-ups fit. Return True.',
    };
    event.textContent = messages[v.phase] || '';
    event.className = `hs-event ${v.phase}`;
  }

  getVariables() {
    const v = this.view || this.steps[this.currentStep] || {};
    const top = v.heap && v.heap.length ? HeapSweepAnimator.fmt(v.heap[0]) : 'none';
    return {
      pick_up: v.km === null || v.km === undefined ? '-' : v.km,
      load: v.load || 0,
      capacity: this.capacity,
      'riding[0]': top,
      riding: `[${(v.heap || []).map(HeapSweepAnimator.fmt).join(', ')}]`,
    };
  }

  getInputData() {
    return {
      trips: this.input,
      capacity: this.capacity,
    };
  }
}

animationRegistry['heap-sweep'] = HeapSweepAnimator;
