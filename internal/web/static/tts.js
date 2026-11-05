// Server-side narration for lesson text. The Go server exposes /api/tts
// (Piper running next to the app), so the voice is the same wherever the
// container runs and nothing depends on the browser's or OS's voices.
//
// Loaded after learn.js and app.js; renderLesson (learn.js) calls
// initLessonAudio() after each chapter render. When /api/tts/status says the
// server has no voice installed, no buttons appear at all.

const ttsState = {
  available: null, // null until /api/tts/status has been checked once
  session: null,   // the active playback: {btn, cancelled, audio, finish}
};

async function ttsAvailable() {
  if (ttsState.available === null) {
    try {
      const res = await fetch('/api/tts/status');
      ttsState.available = res.ok && (await res.json()).available === true;
    } catch (e) {
      ttsState.available = false;
    }
  }
  return ttsState.available;
}

// Every lesson card with a heading gets a Listen button. The examples and
// visualization sections stay silent: they are mostly code.
const TTS_CARDS = '.overview-card, .tutorial-card, .lesson-card, .memory-card';

async function initLessonAudio() {
  stopLessonAudio();
  if (!(await ttsAvailable())) return;
  document.querySelectorAll(`#standard-chapter-content :is(${TTS_CARDS})`).forEach(card => {
    if (card.closest('#lesson-practice-section')) return; // just a problem list
    const heading = card.querySelector('h2');
    if (!heading || heading.querySelector('.tts-btn')) return;
    const btn = document.createElement('button');
    btn.className = 'btn-small ghost tts-btn';
    btn.textContent = 'Listen';
    btn.title = 'Read this section aloud';
    btn.onclick = () => {
      if (ttsState.session && ttsState.session.btn === btn) stopLessonAudio();
      else speakSection(card, btn);
    };
    heading.appendChild(btn);
  });
}

function stopLessonAudio() {
  const s = ttsState.session;
  if (!s) return;
  ttsState.session = null;
  s.cancelled = true;
  if (s.audio) s.audio.pause();
  if (s.finish) s.finish();
  s.btn.textContent = 'Listen';
  s.btn.classList.remove('active');
}

// speechTextFrom keeps inline code (identifiers read fine as words) but drops
// code blocks, controls, collapsed answers and anything currently hidden.
function speechTextFrom(card) {
  const clone = card.cloneNode(true);
  clone.querySelectorAll('pre, button, input, textarea, select, details, .hidden, .tts-btn, .ladder, .check-choices, .pattern-choices')
    .forEach(n => n.remove());
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

async function fetchSpeech(text) {
  const res = await fetch('/api/tts', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text }),
  });
  if (!res.ok) throw new Error((await res.text()).trim() || `TTS failed (${res.status})`);
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
    session.finish = done; // lets stopLessonAudio unblock the playback loop
    audio.onended = done;
    audio.onerror = done;
    audio.play().catch(done);
  });
}

async function speakSection(card, btn) {
  stopLessonAudio();
  const text = speechTextFrom(card);
  if (!text) return;
  const session = { btn, cancelled: false, audio: null, finish: null };
  ttsState.session = session;
  btn.textContent = 'Loading...';
  btn.classList.add('active');
  const chunks = speechChunks(text);
  let failed = false;
  try {
    let next = fetchSpeech(chunks[0]);
    for (let i = 0; i < chunks.length && !session.cancelled; i++) {
      const blob = await next;
      if (session.cancelled) return;
      if (i + 1 < chunks.length) next = fetchSpeech(chunks[i + 1]);
      btn.textContent = 'Stop';
      await playSpeech(blob, session);
    }
  } catch (e) {
    console.error('TTS playback failed:', e);
    failed = !session.cancelled;
  } finally {
    if (ttsState.session === session) {
      ttsState.session = null;
      session.cancelled = true;
      btn.classList.remove('active');
      btn.textContent = failed ? 'Audio failed' : 'Listen';
      if (failed) setTimeout(() => {
        if (btn.textContent === 'Audio failed') btn.textContent = 'Listen';
      }, 3000);
    }
  }
}
