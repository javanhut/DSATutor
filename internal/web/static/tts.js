// Server-side narration for lesson text. The Go server exposes /api/tts
// backed by selectable engines: Ollama speech models (lfm3-audio, phonon)
// when an Ollama server has them installed, and Piper, which ships in the
// Docker image, as the always-there fallback. The voice picker in the
// chapter header remembers the learner's choice; the server falls back to
// the other engines automatically if the chosen one fails. Engines with
// several named voices (phonon's Zoey, Alfie, ...) get a second picker.
//
// Playback is paragraph-aware: each section reads block by block with the
// current block highlighted, hovering any paragraph shows a play button that
// starts reading from there, and a Skip control jumps to the next block.
//
// Practice problems get the same controls: one Listen button in the problem
// panel's header reads the open tab (Description, Hints or Solution).
//
// Loaded after learn.js and app.js; renderLesson (learn.js) calls
// initLessonAudio() after each chapter render and openProblem (app.js) calls
// initPracticeAudio() after each problem render. When /api/tts/status says
// no engine works, none of this UI appears.

const TTS_ENGINE_KEY = 'dsatutor_tts_engine';
const TTS_SPEAKER_KEY = 'dsatutor_tts_speakers'; // engine id -> chosen voice name

// Cards that get narration controls, and what counts as one readable block.
const TTS_CARDS = '.overview-card, .tutorial-card, .lesson-card, .memory-card';
// A practice problem is read one tab at a time.
const TTS_PRACTICE_TABS = '#problem-description, #problem-hints, #problem-solution';
// The entries after .worked-step-head are the parts of a practice problem
// that are not plain paragraphs: its examples, hint labels and the titles of
// the solution walkthrough steps.
const TTS_BLOCKS = 'p, li, h3, h4, .worked-step-head, .example-header, .example-io > div, .example-explanation, .hint-type, .walkthrough-header .subgoal';
// Constraints are skipped like code: they are all symbols (1 <= n <= 10^4).
const TTS_EXCLUDED = 'pre, details, .hidden, .ladder, .check-choices, .pattern-choices, .problem-constraints, .chapter-links';

const ttsState = {
  status: null,  // /api/tts/status payload, null until fetched once
  engine: '',    // selected engine id; '' lets the server pick its default
  speaker: '',   // named voice of that engine; '' is the engine's default
  session: null, // the active playback; see speakSection
};

async function ttsStatus() {
  if (ttsState.status === null) {
    try {
      const res = await fetch('/api/tts/status');
      ttsState.status = res.ok ? await res.json() : { available: false, engines: [] };
    } catch (e) {
      ttsState.status = { available: false, engines: [] };
    }
    const saved = loadJSON(TTS_ENGINE_KEY, '');
    const engines = (ttsState.status.engines || []).filter(e => e.available);
    ttsState.engine = engines.some(e => e.id === saved) ? saved : '';
    initVoicePicker(engines);
    initSpeakerPicker();
  }
  return ttsState.status;
}

// The picker lives in the header of the chapter or practice problem on
// screen (see placePickers); it only shows when there is an actual choice to
// make.
function initVoicePicker(engines) {
  const picker = el('tts-voice');
  if (!picker) return;
  if (engines.length < 2) {
    picker.classList.add('hidden');
    return;
  }
  picker.innerHTML = engines.map(e =>
    `<option value="${escapeHtml(e.id)}"${e.id === ttsState.engine ? ' selected' : ''}>${escapeHtml(e.label)}</option>`
  ).join('');
  picker.classList.remove('hidden');
  picker.onchange = () => {
    stopLessonAudio();
    ttsState.engine = picker.value;
    saveJSON(TTS_ENGINE_KEY, picker.value);
    initSpeakerPicker();
  };
}

// The speaker picker sits beside the voice picker and lists the named voices
// of the engine in use. It only shows for an engine that has some, and the
// choice is remembered per engine.
function initSpeakerPicker() {
  const picker = el('tts-speaker');
  if (!picker) return;
  const status = ttsState.status || {};
  const id = ttsState.engine || status.default;
  const engine = (status.engines || []).find(e => e.id === id);
  const voices = (engine && engine.voices) || [];
  const saved = loadJSON(TTS_SPEAKER_KEY, {}) || {};
  ttsState.speaker = voices.includes(saved[id]) ? saved[id] : '';
  if (!voices.length) {
    picker.classList.add('hidden');
    return;
  }
  picker.innerHTML = ['', ...voices].map(v =>
    `<option value="${escapeHtml(v)}"${v === ttsState.speaker ? ' selected' : ''}>${escapeHtml(v || 'Default voice')}</option>`
  ).join('');
  picker.classList.remove('hidden');
  picker.onchange = () => {
    stopLessonAudio();
    ttsState.speaker = picker.value;
    saveJSON(TTS_SPEAKER_KEY, { ...(loadJSON(TTS_SPEAKER_KEY, {}) || {}), [id]: picker.value });
  };
}

// One pair of pickers serves lessons and practice problems: it moves into
// the header of whichever is on screen, ahead of that header's own buttons
// (anchorId names one of them).
function placePickers(anchorId) {
  const anchor = el(anchorId);
  const pickers = [el('tts-voice'), el('tts-speaker')].filter(Boolean);
  if (anchor && pickers.length) anchor.parentElement.prepend(...pickers);
}

// Adds the Listen and Skip buttons to host; section() returns the element
// to read when Listen is pressed.
function addListenButtons(host, section) {
  const btn = document.createElement('button');
  btn.className = 'btn-small ghost tts-btn';
  btn.textContent = 'Listen';
  btn.title = 'Read this section aloud';
  btn.onclick = () => {
    const card = section();
    if (ttsState.session && ttsState.session.btn === btn) stopLessonAudio();
    else if (card) speakSection(card, btn);
  };
  host.appendChild(btn);
  const skip = document.createElement('button');
  skip.className = 'btn-small ghost tts-btn tts-skip hidden';
  skip.textContent = 'Skip ›';
  skip.title = 'Skip to the next paragraph';
  skip.onclick = skipLessonAudio;
  host.appendChild(skip);
}

async function initLessonAudio() {
  stopLessonAudio();
  if (!(await ttsStatus()).available) return;
  placePickers('toggle-edit');
  initParagraphPlay();
  document.querySelectorAll(`#standard-chapter-content :is(${TTS_CARDS})`).forEach(card => {
    if (card.closest('#lesson-practice-section')) return; // just a problem list
    const heading = card.querySelector('h2');
    if (!heading || heading.querySelector('.tts-btn')) return;
    addListenButtons(heading, () => card);
  });
}

// The problem panel is reused for every problem, so its buttons are added
// once; each call still stops whatever the previous problem was reading.
async function initPracticeAudio() {
  stopLessonAudio();
  if (!(await ttsStatus()).available) return;
  placePickers('practice-run');
  initParagraphPlay();
  const panel = document.querySelector('#practice-mode .problem-panel');
  const header = panel && panel.querySelector('.panel-header');
  if (!header || el('practice-tts')) return;
  const host = document.createElement('div');
  host.id = 'practice-tts';
  host.className = 'panel-tts';
  header.appendChild(host);
  addListenButtons(host, () => panel.querySelector(`:is(${TTS_PRACTICE_TABS}).active`));
  // Switching tabs, revealing a hint and unlocking the solution all replace
  // the text being read, so they end the narration. The Visualize tab has
  // nothing to read.
  panel.addEventListener('click', e => {
    const tab = e.target.closest('.tab-btn');
    if (tab) host.classList.toggle('hidden', tab.dataset.tab === 'visualization' || tab.dataset.tab === 'mine');
    if (tab || e.target.closest('#reveal-hint, #unlock-solution')) stopLessonAudio();
  });
  el('practice-back')?.addEventListener('click', stopLessonAudio);
}

function stopLessonAudio() {
  const s = ttsState.session;
  if (!s) return;
  ttsState.session = null;
  s.cancelled = true;
  if (s.audio) s.audio.pause();
  if (s.finish) s.finish();
  endSessionUI(s);
}

function endSessionUI(s) {
  s.btn.textContent = 'Listen';
  s.btn.classList.remove('active');
  if (s.skipBtn) s.skipBtn.classList.add('hidden');
  if (s.highlighted) s.highlighted.classList.remove('tts-reading');
  s.highlighted = null;
}

// Skip to the next readable block of the playing section.
function skipLessonAudio() {
  const s = ttsState.session;
  if (!s) return;
  s.skipTo = s.currentBlock + 1;
  if (s.audio) s.audio.pause();
  if (s.finish) s.finish(); // unblocks the playback loop, which handles skipTo
}

// speechBlocks lists the readable blocks of a card in document order:
// paragraphs, list items and small headings, skipping code, controls,
// collapsed answers and anything currently hidden. Inline code stays:
// identifiers read fine as words.
function speechBlocks(card) {
  const blocks = [];
  card.querySelectorAll(TTS_BLOCKS).forEach(node => {
    if (node.closest(TTS_EXCLUDED)) return;
    // A block nested in another block (a p inside an li) is read by its parent.
    const parentBlock = node.parentElement && node.parentElement.closest(TTS_BLOCKS);
    if (parentBlock && card.contains(parentBlock)) return;
    const text = blockText(node);
    if (text) blocks.push({ el: node, text });
  });
  if (blocks.length) return blocks;
  // A card with no block-level markup still gets whole-card narration.
  const clone = card.cloneNode(true);
  clone.querySelectorAll(`${TTS_EXCLUDED}, button, input, textarea, select, .tts-btn`).forEach(n => n.remove());
  const text = clone.textContent.replace(/\s+/g, ' ').trim();
  return text ? [{ el: card, text }] : [];
}

function blockText(node) {
  const clone = node.cloneNode(true);
  clone.querySelectorAll('pre, button, input, textarea, select, .hidden').forEach(n => n.remove());
  return clone.textContent.replace(/\s+/g, ' ').trim();
}

// Sentences are grouped into small requests: small enough that the first
// audio arrives in a couple of seconds, large enough that playback outruns
// synthesis once the next chunk is prefetched.
function speechChunks(text, max = 600) {
  const sentences = text.match(/[^.!?]*[.!?]+[\s"')\]]*|[^.!?]+$/g) || [text];
  const chunks = [];
  let cur = '';
  for (let s of sentences) {
    while (s.length > max) { // a sentence longer than a whole chunk: hard split
      if (cur) { chunks.push(cur.trim()); cur = ''; }
      chunks.push(s.slice(0, max).trim());
      s = s.slice(max);
    }
    if ((cur + s).length > max) {
      chunks.push(cur.trim());
      cur = '';
    }
    cur += s;
  }
  if (cur.trim()) chunks.push(cur.trim());
  return chunks.filter(Boolean);
}

async function fetchSpeech(text, retried) {
  const body = { text };
  if (ttsState.engine) body.engine = ttsState.engine;
  if (ttsState.speaker) body.voice = ttsState.speaker;
  const res = await fetch('/api/tts', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  });
  if (!res.ok) {
    const msg = (await res.text()).trim();
    // A remembered engine that no longer exists (model removed): drop the
    // choice, refresh the picker, retry once with the server default.
    if (res.status === 400 && msg.includes('unknown narration engine') && !retried) {
      ttsState.engine = '';
      ttsState.status = null;
      try { localStorage.removeItem(TTS_ENGINE_KEY); } catch (e) { /* ignore */ }
      await ttsStatus();
      return fetchSpeech(text, true);
    }
    throw new Error(msg || `TTS failed (${res.status})`);
  }
  return res.blob();
}

function playSpeech(blob, session) {
  return new Promise(resolve => {
    const url = URL.createObjectURL(blob);
    const audio = new Audio(url);
    let finished = false;
    const done = () => {
      if (finished) return;
      finished = true;
      URL.revokeObjectURL(url);
      resolve();
    };
    session.audio = audio;
    session.finish = done; // lets stop and skip unblock the playback loop
    audio.onended = done;
    audio.onerror = done;
    audio.play().catch(err => {
      // The browser refused to play without a click (autoplay policy):
      // callers can offer a play button instead of staying silent.
      // Brave/Safari do not consistently use NotAllowedError here, so every
      // rejected play promise needs the same visible recovery control.
      session.blocked = true;
      session.playError = err;
      done();
    });
  });
}

function highlightBlock(session, blockEl) {
  if (session.highlighted === blockEl) return;
  if (session.highlighted) session.highlighted.classList.remove('tts-reading');
  session.highlighted = null;
  if (blockEl && blockEl !== session.card) {
    blockEl.classList.add('tts-reading');
    blockEl.scrollIntoView({ block: 'nearest', behavior: 'smooth' });
    session.highlighted = blockEl;
  }
}

async function speakSection(card, btn, startBlock = 0) {
  stopLessonAudio();
  const blocks = speechBlocks(card);
  if (!blocks.length) return;
  // Chunk every block, remembering which block each chunk belongs to, so the
  // highlight and Skip work on paragraphs while requests stay small.
  const chunks = [];
  blocks.forEach((b, bi) => speechChunks(b.text).forEach(text => chunks.push({ text, block: bi })));
  let pos = chunks.findIndex(c => c.block >= startBlock);
  if (pos < 0) return;

  const skipBtn = btn.parentElement.querySelector('.tts-skip');
  const session = {
    btn, skipBtn, card,
    cancelled: false, audio: null, finish: null,
    currentBlock: chunks[pos].block, skipTo: -1, highlighted: null,
  };
  ttsState.session = session;
  btn.textContent = 'Loading...';
  btn.classList.add('active');
  let failed = false;
  let prefetch = { idx: pos, blob: fetchSpeech(chunks[pos].text) };
  try {
    while (pos < chunks.length && !session.cancelled) {
      let blobPromise;
      if (prefetch && prefetch.idx === pos) {
        blobPromise = prefetch.blob;
      } else {
        if (prefetch) prefetch.blob.catch(() => {}); // discarded by a skip
        blobPromise = fetchSpeech(chunks[pos].text);
      }
      const blob = await blobPromise;
      if (session.cancelled) return;
      prefetch = pos + 1 < chunks.length
        ? { idx: pos + 1, blob: fetchSpeech(chunks[pos + 1].text) }
        : null;
      btn.textContent = 'Stop';
      if (skipBtn) skipBtn.classList.remove('hidden');
      session.currentBlock = chunks[pos].block;
      highlightBlock(session, blocks[chunks[pos].block].el);
      await playSpeech(blob, session);
      if (session.cancelled) return;
      if (session.skipTo >= 0) {
        const target = chunks.findIndex(c => c.block >= session.skipTo);
        session.skipTo = -1;
        if (target < 0) break;
        pos = target;
      } else {
        pos++;
      }
    }
  } catch (e) {
    console.error('TTS playback failed:', e);
    failed = !session.cancelled;
  } finally {
    if (prefetch) prefetch.blob.catch(() => {});
    if (ttsState.session === session) {
      ttsState.session = null;
      session.cancelled = true;
      endSessionUI(session);
      if (failed) {
        session.btn.textContent = 'Audio failed';
        setTimeout(() => {
          if (session.btn.textContent === 'Audio failed') session.btn.textContent = 'Listen';
        }, 3000);
      }
    }
  }
}

// blockSection finds the section a block is read with and that section's
// Listen button: a lesson card, or the open tab of a practice problem.
function blockSection(block) {
  const card = block.closest(`#standard-chapter-content :is(${TTS_CARDS})`);
  if (card) {
    if (card.closest('#lesson-practice-section')) return null; // just a problem list
    return { card, btn: card.querySelector('h2 .tts-btn') };
  }
  const tab = block.closest(TTS_PRACTICE_TABS);
  return tab && { card: tab, btn: document.querySelector('#practice-tts .tts-btn') };
}

// A single floating play button follows the hovered paragraph, so any block
// can start playback from there without cluttering the text with controls.
function initParagraphPlay() {
  if (document.getElementById('tts-para-play')) return;
  const fab = document.createElement('button');
  fab.id = 'tts-para-play';
  fab.className = 'tts-para-play hidden';
  fab.title = 'Listen from this paragraph';
  fab.textContent = '▶';
  document.body.appendChild(fab);
  let target = null; // the hovered block, its section and that section's Listen button
  const hide = () => { fab.classList.add('hidden'); target = null; };

  document.addEventListener('mouseover', e => {
    if (!e.target.closest) return;
    if (e.target.closest('#tts-para-play')) return; // hovering the button itself
    const block = e.target.closest(TTS_BLOCKS);
    const section = block && !block.closest(TTS_EXCLUDED) && blockSection(block);
    if (!section || !section.btn) {
      hide();
      return;
    }
    target = { ...section, block };
    const r = block.getBoundingClientRect();
    fab.style.top = `${r.top - 1}px`;
    fab.style.left = `${Math.max(2, r.left - 26)}px`;
    fab.classList.remove('hidden');
  });
  // Fixed positioning goes stale as soon as anything scrolls.
  document.addEventListener('scroll', hide, true);

  fab.onclick = () => {
    if (!target) return;
    const { card, btn, block } = target;
    hide();
    const blocks = speechBlocks(card);
    const idx = blocks.findIndex(b => b.el === block || b.el.contains(block) || block.contains(b.el));
    speakSection(card, btn, Math.max(0, idx));
  };
}
