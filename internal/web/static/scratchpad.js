// The Teacher's Scratchpad: a floating window where the AI tutor writes and
// runs code and draws step-by-step data-structure diagrams, and where the
// learner can edit the code and run it again.
//
// Content is a list of frames the learner steps through. A traced run gives
// one frame per recorded line (code line, variables, data structures, from
// the sandbox tracer); a diagram gives frames the tutor described. Both use
// the visualizer's renderStructure (app.js), plus a bar chart here.
//
// The tutor drives it through window.dsaScratchpad (tutor.js runAction);
// the learner's edits go back to the tutor with every message (state()).

(function () {
  'use strict';

  const EDITOR_ID = 'scratchpad-editor-container';
  const sp = {
    kind: null,      // 'code' or 'diagram'
    title: '',
    explanation: '',
    frames: [],
    idx: 0,
    output: '',
    error: '',
    playing: null,
    lineHandle: null,
    ran: false,      // the code shown has been run (output belongs to it)
    items: [],       // everything shown so far, so a second diagram in one turn doesn't erase the first
    cur: -1,
    question: null,
  };
  const MAX_ITEMS = 20;

  const $ = (id) => document.getElementById(id);

  function build() {
    if ($('scratchpad')) return;
    const box = document.createElement('section');
    box.id = 'scratchpad';
    box.className = 'scratchpad hidden';
    box.setAttribute('aria-label', "Teacher's scratchpad");
    box.innerHTML = `
      <header class="sp-head" id="sp-head">
        <div class="sp-title"><strong>Teacher's scratchpad</strong><span id="sp-subtitle"></span></div>
        <div class="sp-head-actions">
          <span class="sp-items hidden" id="sp-items">
            <button class="btn-small ghost" id="sp-item-prev" title="Previous item">‹</button>
            <span id="sp-item-count"></span>
            <button class="btn-small ghost" id="sp-item-next" title="Next item">›</button>
          </span>
          <button class="btn-small ghost" id="sp-hand" title="Handwritten style for diagrams" aria-label="Handwritten style" aria-pressed="false">✎</button>
          <button class="btn-small ghost" id="sp-showcode" title="Show the code editor: write and run your own code">Code</button>
          <button class="btn-small ghost" id="sp-min" title="Minimize">–</button>
          <button class="btn-small ghost" id="sp-close" title="Close" aria-label="Close scratchpad">✕</button>
        </div>
      </header>
      <div class="sp-body" id="sp-body">
        <p class="sp-expl hidden" id="sp-expl"></p>
        <div class="sp-code" id="sp-code">
          <div id="${EDITOR_ID}" class="cm-editor-container sp-editor"></div>
          <div class="sp-code-actions">
            <button class="btn-small" id="sp-run" title="Run this code and trace it">Run</button>
            <span class="sp-hint">Edit the code and run it yourself. The tutor sees what's here.</span>
          </div>
        </div>
        <div class="sp-player" id="sp-player">
          <button class="btn-small ghost" id="sp-prev" title="Previous step">‹ Back</button>
          <button class="btn-small ghost" id="sp-play" title="Play">Play</button>
          <button class="btn-small ghost" id="sp-next" title="Next step">Next ›</button>
          <input type="range" id="sp-range" min="0" max="0" value="0" aria-label="Step">
          <span class="sp-count" id="sp-count"></span>
        </div>
        <p class="sp-caption" id="sp-caption"></p>
        <div class="sp-viz" id="sp-viz"></div>
        <div class="sp-vars hidden" id="sp-vars"></div>
        <pre class="sp-output hidden" id="sp-output"></pre>
        <section class="sp-checkpoint hidden" id="sp-checkpoint" aria-live="polite">
          <span class="sp-checkpoint-label">Your turn</span>
          <p id="sp-question"></p>
          <div class="sp-checkpoint-choices" id="sp-question-choices"></div>
          <p class="sp-checkpoint-status" id="sp-question-status"></p>
        </section>
      </div>
      <div class="sp-resize" id="sp-resize" aria-hidden="true"></div>`;
    document.body.appendChild(box);

    $('sp-close').onclick = close;
    $('sp-min').onclick = () => box.classList.toggle('minimized');
    $('sp-prev').onclick = () => go(sp.idx - 1);
    $('sp-next').onclick = () => go(sp.idx + 1);
    $('sp-play').onclick = togglePlay;
    $('sp-range').oninput = (e) => go(+e.target.value);
    $('sp-run').onclick = () => run();
    $('sp-hand').onclick = () => setHandwritten(!box.classList.contains('sketchy'));
    setHandwritten(loadPref());
    $('sp-item-prev').onclick = () => loadItem(sp.cur - 1);
    $('sp-item-next').onclick = () => loadItem(sp.cur + 1);
    $('sp-showcode').onclick = () => {
      const code = $('sp-code');
      code.classList.toggle('hidden');
      if (!code.classList.contains('hidden')) setTimeout(() => editor()?.refresh(), 30);
    };
    box.addEventListener('keydown', (e) => {
      if (e.target.closest('.CodeMirror')) return;
      if (e.key === 'ArrowRight') go(sp.idx + 1);
      if (e.key === 'ArrowLeft') go(sp.idx - 1);
    });
    makeDraggable(box, $('sp-head'));
    makeResizable(box, $('sp-resize'));
    if (window.createPythonEditor) {
      window.createPythonEditor(EDITOR_ID, { content: '', language: 'python' });
    }
  }

  function editor() {
    return window.dsaEditors && window.dsaEditors[EDITOR_ID];
  }

  function open() {
    build();
    const box = $('scratchpad');
    box.classList.remove('hidden', 'minimized');
    box.style.zIndex = '70';
    setTimeout(() => editor()?.refresh(), 30);
  }

  function close() {
    stopPlay();
    $('scratchpad')?.classList.add('hidden');
  }

  // ---------------------------------------------------------------------
  // Content

  function showCode({ code, explanation, steps, output, error }) {
    open();
    stopPlay();
    clearQuestion();
    pushItem();
    sp.explanation = explanation || '';
    sp.kind = 'code';
    sp.title = 'Code';
    setExplanation(explanation);
    $('sp-code').classList.remove('hidden');
    const ed = editor();
    if (ed) ed.setValue(code || '');
    if (steps) {
      setRun(steps, output, error);
    } else {
      // Replayed from history: run it again to get the trace.
      sp.frames = [];
      sp.output = '';
      sp.error = '';
      render();
      run();
    }
  }

  // A drawing the teacher talks through (manual) steps along with its
  // narration (step()) instead of playing on a timer.
  function showFrames({ title, frames, manual }) {
    open();
    stopPlay();
    clearQuestion();
    pushItem();
    sp.explanation = '';
    sp.kind = 'diagram';
    sp.title = title || 'Diagram';
    setExplanation('');
    $('sp-code').classList.add('hidden');
    sp.frames = (frames || []).map(f => ({ caption: f.caption || '', structures: f.structures || [] }));
    sp.output = '';
    sp.error = '';
    sp.ran = false;
    sp.idx = 0;
    render();
    if (!manual) autoPlay();
  }

  // Each showCode/showFrames becomes an item; the current one's live state
  // (edited code, run results, position) is saved before switching.
  function saveItem() {
    if (sp.cur < 0 || !sp.items[sp.cur]) return;
    Object.assign(sp.items[sp.cur], {
      kind: sp.kind, title: sp.title, explanation: sp.explanation, frames: sp.frames, idx: sp.idx,
      output: sp.output, error: sp.error, ran: sp.ran,
      code: sp.kind === 'code' && editor() ? editor().getValue() : sp.items[sp.cur].code,
    });
  }

  function pushItem() {
    saveItem();
    sp.items.push({});
    if (sp.items.length > MAX_ITEMS) sp.items.shift();
    sp.cur = sp.items.length - 1;
  }

  function loadItem(i) {
    if (i < 0 || i >= sp.items.length || i === sp.cur) return;
    stopPlay();
    saveItem();
    sp.cur = i;
    const it = sp.items[i];
    Object.assign(sp, { kind: it.kind, title: it.title, explanation: it.explanation, frames: it.frames || [], idx: it.idx || 0,
      output: it.output || '', error: it.error || '', ran: !!it.ran });
    setExplanation(sp.explanation);
    $('sp-code').classList.toggle('hidden', sp.kind !== 'code');
    sp.lastDrawn = -1;
    if (sp.kind === 'code' && editor()) { editor().setValue(it.code || ''); setTimeout(() => editor().refresh(), 30); }
    render();
  }

  // A whiteboard sketch: elements revealed step by step (element.step), so
  // frame k shows everything up to step k and draws step k's strokes live.
  function showSketch({ title, elements, captions, manual }) {
    open();
    stopPlay();
    clearQuestion();
    pushItem();
    sp.explanation = '';
    sp.kind = 'sketch';
    sp.title = title || 'Whiteboard';
    setExplanation('');
    $('sp-code').classList.add('hidden');
    const els = autoSpace((elements || []).map((e, i) => ({ ...e, step: Math.max(0, Math.floor(+e.step || 0)), seed: i + 1 })));
    const steps = Math.max(1, ...els.map(e => e.step + 1));
    sp.frames = Array.from({ length: steps }, (_, k) => ({ caption: (captions || [])[k] || '', sketch: { elements: els, upto: k } }));
    sp.output = '';
    sp.error = '';
    sp.ran = false;
    sp.idx = 0;
    sp.lastDrawn = -1;
    render();
    if (!manual) autoPlay();
  }

  // Move the current drawing to step i (the teacher's narration reached it).
  function step(i) {
    stopPlay();
    go(i);
  }

  function setExplanation(text) {
    const el = $('sp-expl');
    el.textContent = text || '';
    el.classList.toggle('hidden', !text);
  }

  function setRun(steps, output, error) {
    sp.frames = (steps || []).map(st => ({
      line: st.lineNum,
      caption: st.state || '',
      locals: st.locals || {},
      structures: st.structures || [],
    }));
    sp.output = output || '';
    sp.error = error || '';
    sp.ran = true;
    sp.idx = 0;
    render();
  }

  async function run() {
    const code = editor() ? editor().getValue() : '';
    if (!code.trim()) return;
    stopPlay();
    // Running always switches to the code view, even over a diagram.
    if (sp.cur < 0 || sp.kind !== 'code') pushItem();
    sp.kind = 'code';
    sp.title = 'Code';
    $('sp-code').classList.remove('hidden');
    $('sp-run').disabled = true;
    $('sp-run').textContent = 'Running…';
    try {
      const res = await fetch('/api/sandbox/execute', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code, timeout: 10000, language: 'python' }),
      });
      const data = await res.json();
      setRun(data.steps || [], data.output || '', data.error || '');
    } catch (e) {
      setRun([], '', e.message);
    } finally {
      $('sp-run').disabled = false;
      $('sp-run').textContent = 'Run';
    }
  }

  // ---------------------------------------------------------------------
  // Playback

  function go(i) {
    if (!sp.frames.length) return;
    sp.idx = Math.max(0, Math.min(sp.frames.length - 1, i));
    render();
    if (sp.idx === sp.frames.length - 1 && !sp.playing) syncQuestion();
  }

  function togglePlay() {
    if (sp.playing) { stopPlay(); return; }
    if (sp.idx >= sp.frames.length - 1) sp.idx = 0;
    $('sp-play').textContent = 'Pause';
    setTeachingState(true);
    sp.playing = setInterval(() => {
      if (sp.idx >= sp.frames.length - 1) { stopPlay(true); return; }
      go(sp.idx + 1);
    }, sp.kind === 'diagram' ? 1600 : 450);
  }

  function autoPlay() {
    if (sp.frames.length < 2 || sp.frames.length > 8 || window.matchMedia('(prefers-reduced-motion: reduce)').matches) {
      syncQuestion();
      return;
    }
    setTeachingState(true);
    sp.playing = setTimeout(() => {
      sp.playing = null;
      togglePlay();
    }, 650);
  }

  function stopPlay(finished = false) {
    if (sp.playing) clearInterval(sp.playing);
    sp.playing = null;
    if ($('sp-play')) $('sp-play').textContent = 'Play';
    setTeachingState(false);
    if (finished || sp.idx >= sp.frames.length - 1) syncQuestion();
  }

  function setTeachingState(active) {
    const box = $('scratchpad');
    if (!box) return;
    box.classList.toggle('teaching', active);
    box.setAttribute('aria-busy', active ? 'true' : 'false');
  }

  function clearQuestion() {
    sp.question = null;
    const checkpoint = $('sp-checkpoint');
    if (checkpoint) checkpoint.classList.add('hidden');
  }

  function showQuestion(question) {
    if (!question || !question.question || !(question.choices || []).length) return false;
    open();
    sp.question = { question: question.question, choices: question.choices, answered: false };
    $('sp-question').textContent = question.question;
    $('sp-question-choices').innerHTML = question.choices.map(choice =>
      `<button type="button" class="sp-answer" data-sp-answer="${escapeHtml(choice)}">${escapeHtml(choice)}</button>`).join('');
    $('sp-question-choices').onclick = (event) => {
      const answer = event.target.closest('[data-sp-answer]');
      if (!answer || answer.disabled || sp.question?.answered) return;
      sp.question.answered = true;
      $('sp-question-choices').querySelectorAll('button').forEach(button => {
        button.disabled = true;
        button.classList.toggle('chosen', button === answer);
      });
      $('sp-question-status').textContent = 'Listening to your reasoning…';
      window.dispatchEvent(new CustomEvent('dsatutor:scratchpad-answer', { detail: { answer: answer.dataset.spAnswer } }));
    };
    $('sp-checkpoint').classList.remove('hidden');
    syncQuestion();
    return true;
  }

  function syncQuestion() {
    const checkpoint = $('sp-checkpoint');
    if (!checkpoint || !sp.question) return;
    const waiting = !!sp.playing || (sp.frames.length > 1 && sp.idx < sp.frames.length - 1);
    checkpoint.classList.toggle('waiting', waiting);
    $('sp-question-status').textContent = waiting ? 'Watch the example unfold, then make your prediction.' : (sp.question.answered ? 'Listening to your reasoning…' : 'Choose an answer to continue the lesson.');
    $('sp-question-choices').querySelectorAll('button').forEach(button => { button.disabled = waiting || sp.question.answered; });
    if (!waiting) checkpoint.scrollIntoView({ block: 'nearest', behavior: 'smooth' });
  }

  function render() {
    const n = sp.frames.length;
    const f = sp.frames[sp.idx];
    $('sp-subtitle').textContent = sp.title ? ` · ${sp.title}` : '';
    const many = sp.items.length > 1;
    $('sp-items').classList.toggle('hidden', !many);
    if (many) {
      $('sp-item-count').textContent = `${sp.cur + 1} of ${sp.items.length}`;
      $('sp-item-prev').disabled = sp.cur <= 0;
      $('sp-item-next').disabled = sp.cur >= sp.items.length - 1;
    }
    $('sp-player').classList.toggle('hidden', n < 2);
    $('sp-range').max = Math.max(0, n - 1);
    $('sp-range').value = sp.idx;
    $('sp-count').textContent = n ? `${sp.kind === 'diagram' ? 'Frame' : 'Step'} ${sp.idx + 1} of ${n}` : '';
    $('sp-caption').textContent = f ? f.caption : (sp.kind === 'code' && !sp.ran ? 'Running…' : '');

    // A whiteboard frame draws itself; only newly revealed strokes animate.
    const viz = $('sp-viz');
    viz.classList.toggle('sketch-canvas', !!(f && f.sketch));
    if (f && f.sketch) {
      const animate = sp.lastDrawn === sp.idx - 1;
      sp.lastDrawn = sp.idx;
      viz.innerHTML = '';
      viz.appendChild(drawSketch(f.sketch, animate));
      $('sp-vars').classList.add('hidden');
      $('sp-output').classList.add('hidden');
      return;
    }
    sp.lastDrawn = -1;
    const structs = f ? f.structures : [];
    viz.innerHTML = structs.length
      ? structs.map(renderOne).join('')
      : (sp.kind === 'code' && sp.ran && !sp.error ? '<p class="sp-empty">No data structures at this step. Watch the variables below.</p>' : '');

    // Variables (traced runs only).
    const vars = $('sp-vars');
    const locals = f && f.locals ? Object.entries(f.locals) : [];
    vars.classList.toggle('hidden', !locals.length);
    vars.innerHTML = locals.length ? `<table>${locals.map(([k, v]) =>
      `<tr><th>${escapeHtml(k)}</th><td><code>${escapeHtml(short(v))}</code></td></tr>`).join('')}</table>` : '';

    // Output, shown with the run's last step.
    const out = $('sp-output');
    const showOut = sp.kind === 'code' && sp.ran && (sp.error || sp.output) && (sp.idx === n - 1 || n === 0);
    out.classList.toggle('hidden', !showOut);
    out.classList.toggle('error', !!sp.error);
    out.textContent = showOut ? (sp.error ? `${sp.output}${sp.output ? '\n' : ''}${sp.error}` : sp.output) : '';

    // The current line in the code.
    const ed = editor();
    if (ed) {
      if (sp.lineHandle) { ed.removeLineClass(sp.lineHandle, 'background', 'sp-line'); sp.lineHandle = null; }
      if (f && f.line && sp.kind === 'code') {
        sp.lineHandle = ed.addLineClass(f.line - 1, 'background', 'sp-line');
        ed.scrollIntoView({ line: f.line - 1, ch: 0 }, 40);
      }
    }
  }

  function short(v) {
    const s = typeof v === 'string' ? JSON.stringify(v) : JSON.stringify(v);
    return s && s.length > 120 ? s.slice(0, 117) + '…' : (s ?? String(v));
  }

  function renderOne(struct) {
    // A single number or string (a running sum, a counter) under any type.
    const d = struct.data;
    if (struct.type === 'value' || d === null || typeof d !== 'object') {
      return `<div class="sp-scalar"><span class="struct-label">${escapeHtml(struct.name || '')}</span><code>${escapeHtml(d === undefined ? '' : String(typeof d === 'object' ? JSON.stringify(d) : d))}</code></div>`;
    }
    if (struct.type === 'bars') return renderBars(struct);
    if (typeof renderStructure === 'function') return renderStructure(struct);
    return `<pre>${escapeHtml(JSON.stringify(struct, null, 2))}</pre>`;
  }

  // A bar chart for a list of numbers, with the same pointer and window
  // highlights as the array view.
  function renderBars(struct) {
    const data = (Array.isArray(struct.data) ? struct.data : []).map(Number).filter(v => Number.isFinite(v));
    const hl = struct.highlights || {};
    const pointers = hl.pointers || {};
    const win = hl.window || null;
    const maxAbs = Math.max(1, ...data.map(Math.abs));
    const byIdx = {};
    Object.entries(pointers).forEach(([name, i]) => { if (typeof i === 'number') (byIdx[i] = byIdx[i] || []).push(name); });
    const bars = data.map((v, i) => {
      const inWin = win && i >= win.start && i <= win.end;
      const h = Math.max(2, Math.round((Math.abs(v) / maxAbs) * 100));
      return `<div class="sp-bar${inWin ? ' in-window' : ''}${byIdx[i] ? ' pointed' : ''}">
        <span class="sp-bar-val">${escapeHtml(String(v))}</span>
        <div class="sp-bar-fill${v < 0 ? ' neg' : ''}" style="height:${h}%"></div>
        <span class="sp-bar-idx">${i}</span>
        <span class="sp-bar-ptr">${(byIdx[i] || []).map(escapeHtml).join(', ')}</span>
      </div>`;
    }).join('');
    return `<div class="sp-bars-wrap"><div class="struct-label">${escapeHtml(struct.name || '')}</div><div class="sp-bars">${bars}</div></div>`;
  }

  // ---------------------------------------------------------------------
  // Window management

  function makeDraggable(box, handle) {
    let start = null;
    handle.addEventListener('pointerdown', (e) => {
      if (e.target.closest('button')) return;
      box.dataset.placed = 'user'; // the learner chose a spot: never move it
      const r = box.getBoundingClientRect();
      start = { x: e.clientX, y: e.clientY, left: r.left, top: r.top };
      handle.setPointerCapture(e.pointerId);
    });
    handle.addEventListener('pointermove', (e) => {
      if (!start) return;
      const left = Math.max(0, Math.min(window.innerWidth - 120, start.left + e.clientX - start.x));
      const top = Math.max(0, Math.min(window.innerHeight - 40, start.top + e.clientY - start.y));
      Object.assign(box.style, { left: `${left}px`, top: `${top}px`, right: 'auto', bottom: 'auto' });
    });
    handle.addEventListener('pointerup', () => { start = null; });
  }

  function makeResizable(box, grip) {
    let start = null;
    grip.addEventListener('pointerdown', (e) => {
      const r = box.getBoundingClientRect();
      start = { x: e.clientX, y: e.clientY, w: r.width, h: r.height };
      grip.setPointerCapture(e.pointerId);
    });
    grip.addEventListener('pointermove', (e) => {
      if (!start) return;
      box.style.width = `${Math.max(340, start.w + e.clientX - start.x)}px`;
      box.style.height = `${Math.max(260, start.h + e.clientY - start.y)}px`;
      editor()?.refresh();
    });
    grip.addEventListener('pointerup', () => { start = null; });
  }

  // ---------------------------------------------------------------------
  // The whiteboard: rough.js draws the shapes (seeded per element, so a
  // shape looks the same every time it's redrawn), text is in the Caveat
  // handwriting font. Model coordinates are 0-100 across, 0-60+ down.

  const SCALE = 10;
  const FONT = { s: 34, m: 46, l: 62 };
  const NS = 'http://www.w3.org/2000/svg';

  function palette() {
    const css = getComputedStyle(document.documentElement);
    const v = (n, d) => css.getPropertyValue(n).trim() || d;
    return { ink: v('--text', '#222'), red: v('--error', '#b42318'), blue: v('--accent', '#1f4d8a'), green: v('--accent-2', '#2f7a4b'), orange: v('--warning', '#a46810') };
  }

  function num(v, d) { const n = +v; return Number.isFinite(n) ? n : d; }

  function drawSketch({ elements, upto }, animate) {
    const pal = palette();
    const shown = elements.filter(e => e.step <= upto);
    let bottom = 60;
    shown.forEach(e => {
      const y = num(e.y, 0);
      const ptrRoom = e.type === 'cells' && e.pointers && Object.keys(e.pointers).length ? 16 : 0;
      bottom = Math.max(bottom, y + num(e.h, 0) + num(e.r, 0) + (e.type === 'cells' ? 14 + ptrRoom : 6), num(e.y2, 0) + 6);
    });
    const svg = document.createElementNS(NS, 'svg');
    // The CSS gives sketches a viewport whose size is independent of the
    // drawing's aspect ratio.  Keep a little board around the model's
    // 0-100 coordinate system so rough strokes, labels and the resize grip
    // never cut into the drawing.
    const padX = 4 * SCALE;
    const padY = 3 * SCALE;
    svg.setAttribute('viewBox', `${-padX} ${-padY} ${100 * SCALE + padX * 2} ${bottom * SCALE + padY * 2}`);
    svg.setAttribute('preserveAspectRatio', 'xMidYMid meet');
    svg.setAttribute('class', 'sp-sketch');
    svg.setAttribute('role', 'img');
    const rc = window.rough ? window.rough.svg(svg) : null;
    const reduce = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
    const fresh = [];
    shown.forEach(e => {
      const g = drawElement(svg, rc, e, pal[e.color] || pal.ink, pal);
      svg.appendChild(g);
      if (animate && e.step === upto && !reduce) fresh.push(g);
    });
    // The whole step draws in about 2.5 seconds however busy it is.
    const strokes = fresh.reduce((n, g) => n + g.querySelectorAll('path, text').length, 0);
    const pace = Math.min(0.25, 2.5 / Math.max(1, strokes));
    let delay = 0;
    fresh.forEach(g => { delay = animateIn(g, delay, pace); });
    return svg;
  }

  function drawElement(svg, rc, e, color, pal) {
    const g = document.createElementNS(NS, 'g');
    const opts = { stroke: color, strokeWidth: 2.4, roughness: 1.1, bowing: 1.3, seed: e.seed };
    const S = (v, d) => num(v, d) * SCALE;
    const x = S(e.x, 0), y = S(e.y, 0);
    const shape = (node) => { if (node) g.appendChild(node); };
    switch (e.type) {
      case 'text':
        g.appendChild(textNode(e.text, x, y, FONT[e.size] || FONT.m, color, 'start'));
        break;
      case 'box': {
        const w = S(e.w, 20), h = S(e.h, 10);
        shape(rect(rc, x, y, w, h, opts));
        if (e.text) g.appendChild(textNode(e.text, x + w / 2, y + h / 2 + (FONT[e.size] || FONT.m) * 0.35, FONT[e.size] || FONT.m, color, 'middle'));
        break;
      }
      case 'circle': {
        const r = S(e.r, 4);
        shape(ellipse(rc, x, y, r * 2.3, r * 2, opts));
        break;
      }
      case 'node': {
        const r = S(e.r, 4);
        shape(ellipse(rc, x, y, r * 2, r * 2, opts));
        if (e.text !== undefined) g.appendChild(textNode(e.text, x, y + (FONT[e.size] || FONT.m) * 0.35, FONT[e.size] || FONT.m, color, 'middle'));
        break;
      }
      case 'line':
      case 'arrow': {
        const x2 = S(e.x2, num(e.x, 0) + 10), y2 = S(e.y2, num(e.y, 0));
        shape(line(rc, x, y, x2, y2, opts));
        if (e.type === 'arrow') {
          const ang = Math.atan2(y2 - y, x2 - x), len = 22;
          [0.45, -0.45].forEach(d => shape(line(rc, x2, y2, x2 - len * Math.cos(ang + d), y2 - len * Math.sin(ang + d), opts)));
        }
        if (e.text) g.appendChild(textNode(e.text, (x + x2) / 2, (y + y2) / 2 - 12, FONT[e.size] || FONT.s, color, 'middle'));
        break;
      }
      case 'underline': {
        const w = S(e.w, 20);
        shape(line(rc, x, y, x + w, y, { ...opts, strokeWidth: 3 }));
        break;
      }
      case 'cells': {
        const values = Array.isArray(e.values) ? e.values : [];
        const cw = S(e.w, 8), ch = S(e.h, 7);
        const marks = new Set((e.marks || []).map(Number));
        values.forEach((v, i) => {
          const cx = x + i * cw;
          shape(rect(rc, cx, y, cw, ch, { ...opts, seed: e.seed * 31 + i }));
          g.appendChild(textNode(v, cx + cw / 2, y + ch / 2 + FONT.m * 0.35, FONT.m, color, 'middle'));
          g.appendChild(textNode(i, cx + cw / 2, y + ch + FONT.s * 0.9, FONT.s * 0.8, pal.ink, 'middle', 0.55));
          if (marks.has(i)) shape(ellipse(rc, cx + cw / 2, y + ch / 2, cw * 1.05, ch * 1.25, { ...opts, stroke: pal.red, seed: e.seed * 17 + i }));
        });
        if (e.text) g.appendChild(textNode(e.text, x - 12, y + ch / 2 + FONT.m * 0.35, FONT.m, color, 'end'));
        // Pointers by index: an arrow up to each named cell, labels stacked
        // when several share a cell.
        const ptrs = e.pointers && typeof e.pointers === 'object' ? Object.entries(e.pointers) : [];
        const tones = [pal.blue, pal.orange, pal.green, pal.red];
        const perCell = {};
        ptrs.forEach(([name, idx], k) => {
          const i = Math.floor(+idx);
          if (!Number.isFinite(i) || i < 0 || i >= values.length) return;
          const n = perCell[i] = (perCell[i] || 0) + 1;
          const tone = tones[k % tones.length];
          const px = x + i * cw + cw / 2 + (n - 1) * 14;
          const top = y + ch + FONT.s * 1.3, bottomY = top + 70;
          shape(line(rc, px, bottomY, px, top, { ...opts, stroke: tone, seed: e.seed * 13 + k }));
          [0.5, -0.5].forEach(d => shape(line(rc, px, top, px - 18 * Math.cos(-Math.PI / 2 + d), top - 18 * Math.sin(-Math.PI / 2 + d), { ...opts, stroke: tone, seed: e.seed * 7 + k })));
          g.appendChild(textNode(name, px, bottomY + FONT.s + (n - 1) * FONT.s, FONT.s, tone, 'middle'));
        });
        break;
      }
    }
    return g;
  }

  // Models place things too close together. Blocks (text, cells, boxes,
  // nodes) that overlap an earlier block in the same columns are pushed
  // down, and everything below moves at least as far, so the layout keeps
  // its order. Arrows, circles and underlines follow the block above them.
  function autoSpace(els) {
    const unit = (size) => (FONT[size] || FONT.m) / SCALE;
    const extent = (e) => {
      const x = num(e.x, 0), y = num(e.y, 0);
      switch (e.type) {
        case 'text': {
          const u = e.type === 'text' ? unit(e.size) : unit('m');
          const lines = String(e.text ?? '').split('\n');
          const w = Math.max(...lines.map(l => l.length)) * u * 0.42;
          return { top: y - u * 0.85, bottom: y + (lines.length - 1) * u * 1.1 + u * 0.3, left: x, right: x + w };
        }
        case 'cells': {
          const n = Array.isArray(e.values) ? e.values.length : 0;
          const ptr = e.pointers && Object.keys(e.pointers).length ? 18 : 0;
          return { top: y, bottom: y + num(e.h, 7) + 4.5 + ptr, left: x - (e.text ? String(e.text).length * 2.2 + 1.5 : 0), right: x + n * num(e.w, 8) };
        }
        case 'box': return { top: y, bottom: y + num(e.h, 10), left: x, right: x + num(e.w, 20) };
        case 'node': { const r = num(e.r, 4); return { top: y - r, bottom: y + r, left: x - r, right: x + r }; }
      }
      return null;
    };
    const blocks = els.map((e, i) => ({ e, i, ext: extent(e) })).filter(b => b.ext)
      .sort((a, b) => a.ext.top - b.ext.top || a.i - b.i);
    const placed = [];
    const shiftAt = []; // [original top, shift] for the followers
    let carry = 0;
    for (const b of blocks) {
      // Redrawing the same cells in a later step (pointers moving) reuses
      // the earlier row's position instead of being pushed below it.
      const twin = placed.find(p => p.e.type === 'cells' && b.e.type === 'cells' && p.e.step !== b.e.step &&
        Math.abs(num(p.e.x, 0) - num(b.e.x, 0)) < 0.5 && Math.abs(num(p.e.y, 0) - num(b.e.y, 0)) < 0.5);
      if (twin) { b.shift = twin.shift; placed.push({ ...b, top: twin.top, bottom: twin.bottom, ghost: true }); continue; }
      let shift = carry;
      for (let guard = 0; guard < 50; guard++) {
        const hit = placed.find(p => !p.ghost && p.ext.left < b.ext.right && b.ext.left < p.ext.right &&
          b.ext.top + shift < p.bottom - 0.01 && b.ext.bottom + shift > p.top + 0.01);
        if (!hit) break;
        shift = hit.bottom + 1.5 - b.ext.top;
      }
      carry = Math.max(carry, shift);
      b.shift = shift;
      placed.push({ ...b, top: b.ext.top + shift, bottom: b.ext.bottom + shift });
      shiftAt.push([b.ext.top, shift]);
    }
    const shiftFor = (y) => {
      let s = 0;
      for (const [top, sh] of shiftAt) if (top <= y + 0.5) s = Math.max(s, sh);
      return s;
    };
    const byIndex = new Map(blocks.map(b => [b.i, b.shift]));
    return els.map((e, i) => {
      const sh = byIndex.has(i) ? byIndex.get(i) : shiftFor(Math.min(num(e.y, 0), num(e.y2, num(e.y, 0))));
      if (!sh) return e;
      const out = { ...e, y: num(e.y, 0) + sh };
      if (e.y2 !== undefined) out.y2 = num(e.y2, 0) + sh;
      return out;
    });
  }

  function textNode(text, x, y, size, color, anchor, opacity) {
    const t = document.createElementNS(NS, 'text');
    t.setAttribute('x', x);
    t.setAttribute('y', y);
    t.setAttribute('font-size', size);
    t.setAttribute('text-anchor', anchor);
    t.setAttribute('fill', color);
    t.setAttribute('class', 'sp-hand-text');
    if (opacity) t.setAttribute('opacity', opacity);
    String(text ?? '').split('\n').forEach((lineText, i) => {
      const span = document.createElementNS(NS, 'tspan');
      span.setAttribute('x', x);
      if (i) span.setAttribute('dy', size * 1.1);
      span.textContent = lineText;
      t.appendChild(span);
    });
    return t;
  }

  // rough.js shapes, or a simple wobbly fallback when it isn't loaded.
  function rect(rc, x, y, w, h, o) {
    if (rc) return rc.rectangle(x, y, w, h, o);
    return wobblyPath([[x, y], [x + w, y], [x + w, y + h], [x, y + h], [x, y]], o);
  }
  function ellipse(rc, cx, cy, w, h, o) {
    if (rc) return rc.ellipse(cx, cy, w, h, o);
    const pts = [];
    for (let i = 0; i <= 24; i++) { const a = (i / 24) * Math.PI * 2; pts.push([cx + (w / 2) * Math.cos(a), cy + (h / 2) * Math.sin(a)]); }
    return wobblyPath(pts, o);
  }
  function line(rc, x1, y1, x2, y2, o) {
    if (rc) return rc.line(x1, y1, x2, y2, o);
    return wobblyPath([[x1, y1], [(x1 + x2) / 2, (y1 + y2) / 2], [x2, y2]], o);
  }
  function wobblyPath(pts, o) {
    let seed = o.seed || 1;
    const rnd = () => { seed = (seed * 9301 + 49297) % 233280; return seed / 233280 - 0.5; };
    const p = document.createElementNS(NS, 'path');
    p.setAttribute('d', pts.map(([px, py], i) => `${i ? 'L' : 'M'}${(px + rnd() * 4).toFixed(1)} ${(py + rnd() * 4).toFixed(1)}`).join(' '));
    p.setAttribute('fill', 'none');
    p.setAttribute('stroke', o.stroke);
    p.setAttribute('stroke-width', o.strokeWidth);
    p.setAttribute('stroke-linecap', 'round');
    p.setAttribute('stroke-linejoin', 'round');
    return p;
  }

  // Draw a group's strokes in order like a pen, then fade its text in.
  function animateIn(g, delay, pace) {
    g.querySelectorAll('path').forEach(p => {
      if (p.getAttribute('fill') && p.getAttribute('fill') !== 'none') return;
      let len = 0;
      try { len = p.getTotalLength(); } catch (err) { return; }
      if (!len) return;
      const dur = Math.min(0.6, Math.max(pace * 1.5, 0.12));
      p.style.strokeDasharray = `${len}`;
      p.style.strokeDashoffset = `${len}`;
      p.style.transition = `stroke-dashoffset ${dur}s ease-out ${delay}s`;
      requestAnimationFrame(() => requestAnimationFrame(() => { p.style.strokeDashoffset = '0'; }));
      delay += pace;
    });
    g.querySelectorAll('text').forEach(t => {
      t.style.opacity = '0';
      t.style.transition = `opacity 0.3s ease ${delay}s`;
      requestAnimationFrame(() => requestAnimationFrame(() => { t.style.opacity = t.getAttribute('opacity') || '1'; }));
      delay += pace;
    });
    return delay;
  }

  // ---------------------------------------------------------------------
  // Handwritten style for every diagram (a pen-wobble filter and the
  // handwriting font), remembered per browser.

  const HAND_KEY = 'dsatutor_scratchpad_handwritten';
  function loadPref() {
    try { return localStorage.getItem(HAND_KEY) === '1'; } catch (e) { return false; }
  }
  function setHandwritten(on) {
    const box = $('scratchpad');
    if (!box) return;
    if (on && !$('sp-wobble-defs')) {
      const defs = document.createElementNS(NS, 'svg');
      defs.id = 'sp-wobble-defs';
      defs.setAttribute('width', '0');
      defs.setAttribute('height', '0');
      defs.style.position = 'absolute';
      defs.innerHTML = '<filter id="sp-wobble"><feTurbulence type="fractalNoise" baseFrequency="0.035" numOctaves="2" seed="7"/><feDisplacementMap in="SourceGraphic" scale="2"/></filter>';
      document.body.appendChild(defs);
    }
    box.classList.toggle('sketchy', on);
    $('sp-hand').classList.toggle('active', on);
    $('sp-hand').setAttribute('aria-pressed', on ? 'true' : 'false');
    try { localStorage.setItem(HAND_KEY, on ? '1' : ''); } catch (e) { /* private mode */ }
  }

  // Move the scratchpad out of the way of a spot the tutor is pointing at
  // (rect in viewport coordinates): to the top or the bottom of the window,
  // whichever leaves the spot visible. A scratchpad the learner dragged
  // stays where they put it.
  function avoid(rect) {
    const box = $('scratchpad');
    if (!box || box.classList.contains('hidden') || box.dataset.placed === 'user' || !rect) return;
    const r = box.getBoundingClientRect();
    const overlaps = r.left < rect.right && rect.left < r.right && r.top < rect.bottom && rect.top < r.bottom;
    if (!overlaps) return;
    const h = r.height, margin = 16;
    if (rect.top > h + margin * 2) {
      Object.assign(box.style, { top: `${margin}px`, bottom: 'auto' }); // spot is low: go to the top
    } else if (window.innerHeight - rect.bottom > h + margin * 2) {
      Object.assign(box.style, { top: 'auto', bottom: `${margin}px` }); // spot is high: go to the bottom
    } else {
      // No vertical room: shrink to fit under or over the spot.
      const below = window.innerHeight - rect.bottom - margin * 2;
      const above = rect.top - margin * 2;
      if (below >= above) Object.assign(box.style, { top: 'auto', bottom: `${margin}px`, height: `${Math.max(220, below)}px` });
      else Object.assign(box.style, { top: `${margin}px`, bottom: 'auto', height: `${Math.max(220, above)}px` });
    }
  }

  // What the tutor sees of the scratchpad.
  function state() {
    if (!$('scratchpad') || !sp.kind) return null;
    const frame = sp.frames[sp.idx] || {};
    const out = {
      kind: sp.kind,
      title: sp.title,
      step: sp.frames.length ? sp.idx + 1 : 0,
      steps: sp.frames.length,
      caption: frame.caption || '',
    };
    if (sp.kind === 'code') {
      out.code = editor() ? editor().getValue() : '';
      out.output = sp.ran ? `${sp.output}${sp.error ? `\n${sp.error}` : ''}`.trim() : '';
    }
    return out;
  }

  window.dsaScratchpad = { showCode, showFrames, showSketch, showQuestion, open, close, state, setHandwritten, avoid, step };
})();
