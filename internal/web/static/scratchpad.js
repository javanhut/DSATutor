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

  function showFrames({ title, frames }) {
    open();
    stopPlay();
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
    if (sp.kind === 'code' && editor()) { editor().setValue(it.code || ''); setTimeout(() => editor().refresh(), 30); }
    render();
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
  }

  function togglePlay() {
    if (sp.playing) { stopPlay(); return; }
    if (sp.idx >= sp.frames.length - 1) sp.idx = 0;
    $('sp-play').textContent = 'Pause';
    sp.playing = setInterval(() => {
      if (sp.idx >= sp.frames.length - 1) { stopPlay(); return; }
      go(sp.idx + 1);
    }, sp.kind === 'diagram' ? 1600 : 450);
  }

  function stopPlay() {
    if (sp.playing) clearInterval(sp.playing);
    sp.playing = null;
    if ($('sp-play')) $('sp-play').textContent = 'Play';
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
    $('sp-count').textContent = n ? `${sp.kind === 'code' ? 'Step' : 'Frame'} ${sp.idx + 1} of ${n}` : '';
    $('sp-caption').textContent = f ? f.caption : (sp.kind === 'code' && !sp.ran ? 'Running…' : '');

    // The structures of this frame.
    const viz = $('sp-viz');
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

  // What the tutor sees of the scratchpad.
  function state() {
    if (!$('scratchpad') || sp.kind !== 'code') return null;
    const code = editor() ? editor().getValue() : '';
    if (!code.trim()) return null;
    return { code, output: sp.ran ? `${sp.output}${sp.error ? `\n${sp.error}` : ''}`.trim() : '' };
  }

  window.dsaScratchpad = { showCode, showFrames, open, close, state };
})();
