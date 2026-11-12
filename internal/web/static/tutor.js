// AI tutor: an optional chat panel backed by an Ollama language model.
//
// The panel opens only when the learner presses a Tutor button (sidebar, or
// "Ask tutor" in a practice problem's header). Every message carries what is
// on screen: the open problem or chapter, the code in the editor, the
// revealed hints and the last run or submission, so the tutor can point at
// the line where the logic breaks. The server builds the prompt from the
// real problem data (see internal/tutor).
//
// Voice: the microphone button records a question, the server transcribes it
// with an audio-capable Ollama model, and the question is sent. Replies can
// be read aloud sentence by sentence while they stream, through the same
// narration engines as the lessons (phonon, lfm3-audio, Piper; see tts.js).
//
// Loaded after app.js and tts.js. When /api/tutor/status reports no chat
// model, none of this UI appears.

(function () {
  'use strict';

  const MODEL_KEY = 'dsatutor_tutor_model';
  const PREFS_KEY = 'dsatutor_tutor_prefs';   // { think, speak }
  const HISTORY_KEY = 'dsatutor_tutor_chats'; // context key -> messages
  const MAX_SAVED_CHATS = 30;
  const MAX_SAVED_MESSAGES = 40;

  const tutor = {
    status: null,
    model: '',
    prefs: { think: false, speak: false },
    open: false,
    key: '',          // which conversation is showing (problem, chapter...)
    messages: [],     // [{ role, content }]
    busy: false,
    abort: null,      // AbortController of the streaming reply
    recorder: null,   // the active microphone recording
    speech: null,     // the active spoken reply
  };

  const $ = (id) => document.getElementById(id);
  const visible = (id) => { const n = $(id); return !!n && !n.classList.contains('hidden'); };

  function load(key, fallback) {
    try { const v = localStorage.getItem(key); return v ? JSON.parse(v) : fallback; } catch (e) { return fallback; }
  }
  function save(key, value) {
    try { localStorage.setItem(key, JSON.stringify(value)); } catch (e) { /* private mode or full */ }
  }

  // ---------------------------------------------------------------------
  // What the learner is looking at

  function practiceRunText() {
    const parts = [];
    const problem = practiceState.currentProblem;
    const single = $('current-test-output');
    if (single && single.innerText.trim()) {
      const tc = problem?.testCases?.[practiceState.currentTestCase];
      const input = tc ? formatTestValue(tc.input) : '';
      const expected = tc ? JSON.stringify(tc.expected) : '';
      parts.push(`Ran test case ${practiceState.currentTestCase + 1}${input ? ` (input ${input}, expected ${expected})` : ''}:\n${single.innerText.trim()}`);
    }
    const all = $('test-results');
    if (all && all.innerText.trim()) parts.push(`Submission:\n${all.innerText.trim()}`);
    return parts.join('\n\n');
  }

  function currentContext() {
    if (visible('practice-mode') && typeof practiceState !== 'undefined' && practiceState.currentProblem) {
      const p = practiceState.currentProblem;
      const prog = practiceState.progress[p.id] || {};
      return {
        key: `problem:${p.id}`,
        label: p.title,
        kind: 'problem',
        ctx: {
          view: 'problem',
          problemId: p.id,
          language: practiceState.language,
          code: getEditorCode('practice-code'),
          lastRun: practiceRunText(),
          hintsRevealed: Math.max(practiceState.hintsRevealed || 0, prog.hintsUsed || 0),
          solutionUnlocked: !!(prog.solved || prog.solutionViewed),
        },
      };
    }
    if (visible('sandbox-mode')) {
      const out = $('sandbox-output');
      return {
        key: 'sandbox',
        label: 'Sandbox',
        kind: 'sandbox',
        ctx: {
          view: 'sandbox',
          language: typeof sandboxState !== 'undefined' ? sandboxState.language : 'python',
          code: getEditorCode('sandbox-code'),
          lastRun: out ? out.innerText.trim() : '',
        },
      };
    }
    if (visible('detail') && typeof state !== 'undefined' && state.current) {
      const ch = state.current;
      return {
        key: `chapter:${ch.number}`,
        label: `Ch. ${ch.number}: ${ch.title}`,
        kind: 'chapter',
        ctx: { view: 'chapter', chapterNumber: ch.number },
      };
    }
    return { key: 'home', label: 'General', kind: 'home', ctx: { view: 'home' } };
  }

  const QUICK_PROMPTS = {
    problem: [
      ['Where does my logic break?', 'Look at my current code and my last run. Where does my logic go wrong? Point to the line and give me an input that breaks it, but let me fix it myself.'],
      ['Give me a nudge', "I'm stuck. Give me the smallest nudge toward the right approach, without code."],
      ['Check my complexity', "What are the time and space complexity of my current code? Ask me what I think first if I haven't said, then explain."],
      ['Better approach?', 'Is there a better approach than what I have? Compare it with brute force and tell me what pattern to think about.'],
      ['Explain the problem', 'Can you restate this problem in plain words and walk through the first example with me?'],
    ],
    chapter: [
      ['Explain it simpler', 'Explain the main idea of this chapter in simpler terms, with a tiny example.'],
      ['Quiz me', 'Quiz me with one question about this chapter. Wait for my answer before giving the next one.'],
      ['When do I use this?', 'How do I recognize in an interview problem that this pattern applies?'],
      ['Walk an example', 'Walk me through a small example of this technique step by step, and ask me to predict the next step.'],
    ],
    sandbox: [
      ['What does my code do?', 'Explain what my sandbox code does, step by step.'],
      ['Find the bug', 'Is there a bug in my code? Point to the line and explain why.'],
      ['Complexity?', 'What is the time and space complexity of my code? Explain how you counted.'],
    ],
    home: [
      ['What should I study?', 'I want to get better at DSA interviews. How should I go through this course, and what should I focus on first?'],
      ['Explain Big-O', 'Explain Big-O notation with simple examples.'],
    ],
  };

  // ---------------------------------------------------------------------
  // Conversation storage, one chat per problem / chapter / sandbox

  function loadConversation(key) {
    const all = load(HISTORY_KEY, {}) || {};
    return Array.isArray(all[key]?.messages) ? all[key].messages : [];
  }

  function saveConversation() {
    const all = load(HISTORY_KEY, {}) || {};
    if (tutor.messages.length) {
      all[tutor.key] = { at: Date.now(), messages: tutor.messages.slice(-MAX_SAVED_MESSAGES) };
    } else {
      delete all[tutor.key];
    }
    const keys = Object.keys(all).sort((a, b) => (all[b].at || 0) - (all[a].at || 0));
    keys.slice(MAX_SAVED_CHATS).forEach(k => delete all[k]);
    save(HISTORY_KEY, all);
  }

  // ---------------------------------------------------------------------
  // Markdown (just what tutors write: code, lists, emphasis, headings)

  function inlineMd(text) {
    return escapeHtml(text)
      .replace(/`([^`]+)`/g, '<code>$1</code>')
      .replace(/\*\*([^*]+)\*\*/g, '<strong>$1</strong>')
      .replace(/(^|[^*])\*([^*\s][^*]*)\*/g, '$1<em>$2</em>');
  }

  function renderMarkdown(src) {
    const out = [];
    const parts = src.split(/```/);
    parts.forEach((part, i) => {
      if (i % 2 === 1) { // inside a fence; the first line may name the language
        const code = part.replace(/^[\w+-]*\n/, '');
        out.push(`<pre class="tutor-code"><code>${escapeHtml(code.replace(/\n$/, ''))}</code></pre>`);
        return;
      }
      let list = null; // 'ul' or 'ol'
      let para = [];
      const flushPara = () => { if (para.length) { out.push(`<p>${inlineMd(para.join(' '))}</p>`); para = []; } };
      const closeList = () => { if (list) { out.push(`</${list}>`); list = null; } };
      part.split('\n').forEach(line => {
        const t = line.trim();
        let m;
        if (!t) { flushPara(); closeList(); return; }
        if ((m = t.match(/^#{1,6}\s+(.*)$/))) { flushPara(); closeList(); out.push(`<h4>${inlineMd(m[1])}</h4>`); return; }
        if ((m = t.match(/^[-*•]\s+(.*)$/)) || (m = t.match(/^\d+[.)]\s+(.*)$/))) {
          const kind = /^\d/.test(t) ? 'ol' : 'ul';
          flushPara();
          if (list !== kind) { closeList(); out.push(`<${kind}>`); list = kind; }
          out.push(`<li>${inlineMd(m[1])}</li>`);
          return;
        }
        closeList();
        para.push(t);
      });
      flushPara();
      closeList();
    });
    return out.join('');
  }

  // Text fit for speaking: no code blocks, no Markdown symbols.
  function speakable(text) {
    return text
      .replace(/```[\s\S]*?(```|$)/g, ' ')
      .replace(/`([^`]*)`/g, '$1')
      .replace(/[*_#>]+/g, '')
      .replace(/^\s*[-•]\s+/gm, '')
      .replace(/\s+/g, ' ')
      .trim();
  }

  // ---------------------------------------------------------------------
  // Panel

  function buildPanel() {
    const panel = document.createElement('aside');
    panel.id = 'tutor-panel';
    panel.className = 'tutor-panel';
    panel.setAttribute('aria-label', 'AI tutor');
    panel.innerHTML = `
      <header class="tutor-head">
        <div class="tutor-title">
          <h2>Tutor</h2>
          <span class="tutor-context" id="tutor-context" title="What the tutor can see"></span>
        </div>
        <div class="tutor-head-actions">
          <button class="btn-small ghost" id="tutor-clear" title="Start this conversation over">Clear</button>
          <button class="btn-small ghost tutor-close" id="tutor-close" title="Close (Esc)" aria-label="Close tutor">✕</button>
        </div>
      </header>
      <div class="tutor-settings">
        <select id="tutor-model" title="Language model"></select>
        <label class="tutor-toggle hidden" id="tutor-think-wrap" title="Let the model reason before answering: slower, better at finding bugs">
          <input type="checkbox" id="tutor-think"> Think deeper
        </label>
        <label class="tutor-toggle hidden" id="tutor-speak-wrap" title="Read replies aloud with the narration voice">
          <input type="checkbox" id="tutor-speak"> Speak replies
        </label>
      </div>
      <div class="tutor-messages" id="tutor-messages" aria-live="polite"></div>
      <div class="tutor-quick" id="tutor-quick"></div>
      <form class="tutor-input" id="tutor-form">
        <textarea id="tutor-text" rows="2" placeholder="Ask about this problem, your code, complexity..."></textarea>
        <div class="tutor-input-actions">
          <button type="button" class="btn-small ghost tutor-mic hidden" id="tutor-mic" title="Ask by voice (click to start, click again to send)" aria-label="Record a question">
            <svg class="icon" viewBox="0 0 24 24" aria-hidden="true"><rect x="9" y="2" width="6" height="12" rx="3"/><path d="M19 10v1a7 7 0 0 1-14 0v-1"/><path d="M12 18v4"/></svg>
          </button>
          <button type="submit" class="btn-small primary" id="tutor-send">Send</button>
        </div>
      </form>
      <p class="tutor-status" id="tutor-status"></p>
    `;
    document.body.appendChild(panel);

    $('tutor-close').onclick = closePanel;
    $('tutor-clear').onclick = () => {
      stopReply();
      stopSpeaking();
      tutor.messages = [];
      saveConversation();
      renderMessages();
    };
    $('tutor-form').onsubmit = (e) => { e.preventDefault(); send($('tutor-text').value); };
    $('tutor-text').addEventListener('keydown', (e) => {
      if (e.key === 'Enter' && !e.shiftKey && !e.isComposing) { e.preventDefault(); send($('tutor-text').value); }
    });
    $('tutor-mic').onclick = toggleRecording;
    $('tutor-think').onchange = (e) => { tutor.prefs.think = e.target.checked; save(PREFS_KEY, tutor.prefs); };
    $('tutor-speak').onchange = (e) => {
      tutor.prefs.speak = e.target.checked;
      save(PREFS_KEY, tutor.prefs);
      if (!e.target.checked) stopSpeaking();
    };
    $('tutor-model').onchange = (e) => {
      tutor.model = e.target.value;
      save(MODEL_KEY, tutor.model);
      syncThinkToggle();
    };
    document.addEventListener('keydown', (e) => {
      if (e.key === 'Escape' && tutor.open && !e.defaultPrevented) {
        if (tutor.recorder) cancelRecording();
        else closePanel();
      }
    });
  }

  function populateSettings() {
    const models = tutor.status.models || [];
    const saved = load(MODEL_KEY, '');
    tutor.model = models.some(m => m.name === saved) ? saved : tutor.status.default;
    $('tutor-model').innerHTML = models.map(m =>
      `<option value="${escapeHtml(m.name)}"${m.name === tutor.model ? ' selected' : ''}>${escapeHtml(m.name)}${m.cloud ? ' (cloud)' : ''}</option>`
    ).join('');
    tutor.prefs = { think: false, speak: false, ...(load(PREFS_KEY, {}) || {}) };
    $('tutor-think').checked = tutor.prefs.think;
    $('tutor-speak').checked = tutor.prefs.speak;
    syncThinkToggle();
    $('tutor-mic').classList.toggle('hidden', !tutor.status.transcribe);
    if (typeof ttsStatus === 'function') {
      ttsStatus().then(st => $('tutor-speak-wrap').classList.toggle('hidden', !st.available));
    }
  }

  function syncThinkToggle() {
    const m = (tutor.status.models || []).find(x => x.name === tutor.model);
    $('tutor-think-wrap').classList.toggle('hidden', !(m && m.thinking));
  }

  // Each problem, chapter and the sandbox keep their own conversation; the
  // panel switches to the right one whenever it opens or a message is sent.
  function syncContext() {
    const c = currentContext();
    $('tutor-context').textContent = c.kind === 'problem' ? `${c.label} · sees your code and runs`
      : c.kind === 'sandbox' ? 'Sandbox · sees your code'
      : c.label;
    if (c.key !== tutor.key) {
      if (tutor.busy) stopReply();
      tutor.key = c.key;
      tutor.messages = loadConversation(c.key);
      renderMessages();
    }
    $('tutor-text').placeholder = {
      problem: 'Ask about this problem, your code, complexity...',
      chapter: 'Ask about this chapter, or say "quiz me"...',
      sandbox: 'Ask about your sandbox code...',
      home: 'Ask anything about data structures and algorithms...',
    }[c.kind];
    $('tutor-quick').innerHTML = QUICK_PROMPTS[c.kind].map(([label], i) =>
      `<button type="button" class="tutor-chip" data-i="${i}">${escapeHtml(label)}</button>`
    ).join('');
    $('tutor-quick').onclick = (e) => {
      const chip = e.target.closest('.tutor-chip');
      if (chip) send(QUICK_PROMPTS[c.kind][+chip.dataset.i][1]);
    };
    return c;
  }

  async function openPanel(prefill) {
    const st = await loadStatus();
    if (!st.available) return;
    if (!$('tutor-panel')) { buildPanel(); populateSettings(); }
    tutor.open = true;
    document.body.classList.add('tutor-open');
    $('tutor-panel').classList.add('open');
    $('tutor-toggle')?.classList.add('active');
    syncContext();
    setStatus('');
    if (prefill) $('tutor-text').value = prefill;
    $('tutor-text').focus();
  }

  function closePanel() {
    tutor.open = false;
    document.body.classList.remove('tutor-open');
    $('tutor-panel')?.classList.remove('open');
    $('tutor-toggle')?.classList.remove('active');
    if (tutor.recorder) cancelRecording();
    stopSpeaking();
  }

  function togglePanel() {
    if (tutor.open) closePanel();
    else openPanel();
  }

  function setStatus(text, isError) {
    const s = $('tutor-status');
    if (!s) return;
    s.textContent = text;
    s.classList.toggle('error', !!isError);
  }

  function renderMessages() {
    const box = $('tutor-messages');
    if (!box) return;
    if (!tutor.messages.length) {
      box.innerHTML = `<div class="tutor-empty">
        <p>Ask anything: why your code fails a test, how to start, what the complexity is, or to explain an idea again.</p>
        <p class="tutor-empty-note">The tutor sees what's on screen${tutor.key.startsWith('problem:') ? ', including your code and your last run' : ''}. It coaches with hints instead of handing you the answer, unless you ask for it.</p>
      </div>`;
      return;
    }
    box.innerHTML = tutor.messages.map(m => messageHtml(m)).join('');
    box.scrollTop = box.scrollHeight;
  }

  function messageHtml(m) {
    const body = m.role === 'assistant' ? renderMarkdown(m.content) : `<p>${escapeHtml(m.content).replace(/\n/g, '<br>')}</p>`;
    return `<div class="tutor-msg ${m.role}">${body}</div>`;
  }

  // ---------------------------------------------------------------------
  // Chat

  function setBusy(busy) {
    tutor.busy = busy;
    const btn = $('tutor-send');
    btn.textContent = busy ? 'Stop' : 'Send';
    btn.type = busy ? 'button' : 'submit';
    btn.onclick = busy ? (e) => { e.preventDefault(); stopReply(); } : null;
  }

  function stopReply() {
    if (tutor.abort) tutor.abort.abort();
  }

  async function send(text, opts = {}) {
    text = (text || '').trim();
    if (!text || tutor.busy) return;
    const c = syncContext();
    $('tutor-text').value = '';
    tutor.messages.push({ role: 'user', content: text });
    renderMessages();
    stopSpeaking();
    if (typeof stopLessonAudio === 'function') stopLessonAudio();

    const box = $('tutor-messages');
    const bubble = document.createElement('div');
    bubble.className = 'tutor-msg assistant pending';
    bubble.innerHTML = '<p class="tutor-typing">Thinking…</p>';
    box.appendChild(bubble);
    box.scrollTop = box.scrollHeight;

    // A question asked by voice gets a spoken answer even with Speak replies
    // off, so a conversation can go back and forth hands-free.
    const speak = (tutor.prefs.speak || !!opts.voice) && !$('tutor-speak-wrap').classList.contains('hidden');
    const speech = speak ? startSpeaking() : null;
    const ctrl = new AbortController();
    tutor.abort = ctrl;
    setBusy(true);
    setStatus('');

    let reply = '';
    let thinking = '';
    let failed = '';
    try {
      const res = await fetch('/api/tutor/chat', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        signal: ctrl.signal,
        body: JSON.stringify({
          model: tutor.model,
          think: tutor.prefs.think,
          messages: tutor.messages,
          context: { ...c.ctx, voice: speak },
        }),
      });
      if (!res.ok) {
        const msg = (await res.text()).trim();
        if (res.status === 400 && msg.includes('unknown tutor model')) {
          // The remembered model was removed; fall back to the default.
          tutor.status = null;
          try { localStorage.removeItem(MODEL_KEY); } catch (e) { /* ignore */ }
        }
        throw new Error(msg || `Tutor failed (${res.status})`);
      }
      const reader = res.body.getReader();
      const decoder = new TextDecoder();
      let buf = '';
      for (;;) {
        const { value, done } = await reader.read();
        if (done) break;
        buf += decoder.decode(value, { stream: true });
        let nl;
        while ((nl = buf.indexOf('\n')) >= 0) {
          const line = buf.slice(0, nl).trim();
          buf = buf.slice(nl + 1);
          if (!line) continue;
          const ev = JSON.parse(line);
          if (ev.type === 'delta') {
            reply += ev.text;
            bubble.innerHTML = renderMarkdown(reply);
            if (speech) speech.push(reply);
          } else if (ev.type === 'thinking') {
            thinking += ev.text;
            if (!reply) bubble.innerHTML = `<p class="tutor-typing">Thinking… <span class="tutor-think-count">${thinking.length} chars of reasoning</span></p>`;
          } else if (ev.type === 'error') {
            throw new Error(ev.error);
          }
        }
        box.scrollTop = box.scrollHeight;
      }
    } catch (e) {
      if (e.name !== 'AbortError') failed = e.message || String(e);
    } finally {
      tutor.abort = null;
      setBusy(false);
    }

    bubble.classList.remove('pending');
    if (reply.trim()) {
      tutor.messages.push({ role: 'assistant', content: reply });
      bubble.innerHTML = renderMarkdown(reply);
      if (speech) speech.end(reply);
    } else {
      bubble.remove();
      if (speech) stopSpeaking();
      if (!failed) tutor.messages.pop(); // stopped before any answer: drop the question too
    }
    if (failed) setStatus(failed, true);
    saveConversation();
    renderMessages();
  }

  // ---------------------------------------------------------------------
  // Spoken replies: complete sentences are queued while the reply streams
  // and played in order, prefetching the next one, through /api/tts.

  function startSpeaking() {
    stopSpeaking();
    const s = {
      queue: [],     // text chunks still to fetch and play
      spokenUpTo: 0, // characters of the speakable text already queued
      done: false,
      cancelled: false,
      audio: null,  // set by playSpeech (tts.js), with finish() to cut it short
      finish: null,
      wake: null,
    };
    tutor.speech = s;
    s.push = (full) => {
      const text = speakable(full);
      // Queue only finished sentences; the rest may still be growing.
      const re = /[^.!?]*[.!?]+(\s|$)/g;
      re.lastIndex = s.spokenUpTo;
      let m, end = s.spokenUpTo;
      while ((m = re.exec(text)) && m[0].length) end = re.lastIndex;
      if (end - s.spokenUpTo >= 40 || (end > s.spokenUpTo && s.queue.length === 0 && !s.audio)) {
        s.queue.push(text.slice(s.spokenUpTo, end).trim());
        s.spokenUpTo = end;
        if (s.wake) s.wake();
      }
    };
    s.end = (full) => {
      const text = speakable(full);
      const rest = text.slice(s.spokenUpTo).trim();
      if (rest) s.queue.push(rest);
      s.spokenUpTo = text.length;
      s.done = true;
      if (s.wake) s.wake();
    };
    playLoop(s);
    return s;
  }

  async function playLoop(s) {
    let next = null; // prefetched { text, blob }
    const fetchChunk = (text) => ({ text, blob: fetchSpeech(text) });
    while (!s.cancelled) {
      if (!next && s.queue.length) next = fetchChunk(s.queue.shift());
      if (!next) {
        if (s.done) break;
        await new Promise(r => { s.wake = r; });
        s.wake = null;
        continue;
      }
      let blob;
      try { blob = await next.blob; } catch (e) { console.error('Tutor speech failed:', e); break; }
      next = s.queue.length ? fetchChunk(s.queue.shift()) : null;
      if (s.cancelled) break;
      await playSpeech(blob, s);
      s.audio = null;
      s.finish = null;
    }
    if (next) next.blob.catch(() => {});
    if (tutor.speech === s) tutor.speech = null;
  }

  function stopSpeaking() {
    const s = tutor.speech;
    if (!s) return;
    tutor.speech = null;
    s.cancelled = true;
    if (s.audio) s.audio.pause();
    if (s.finish) s.finish(); // unblocks playSpeech
    if (s.wake) s.wake();
  }

  // ---------------------------------------------------------------------
  // Microphone: record 16 kHz mono WAV, transcribe on the server, send.

  async function toggleRecording() {
    if (tutor.recorder) return finishRecording();
    if (!navigator.mediaDevices?.getUserMedia) {
      setStatus('The microphone needs a secure page: open the tutor at http://localhost.', true);
      return;
    }
    stopSpeaking();
    let stream;
    try {
      stream = await navigator.mediaDevices.getUserMedia({ audio: { channelCount: 1, echoCancellation: true, noiseSuppression: true } });
    } catch (e) {
      setStatus('Microphone access was denied.', true);
      return;
    }
    const ctx = new (window.AudioContext || window.webkitAudioContext)();
    const source = ctx.createMediaStreamSource(stream);
    // ScriptProcessor is deprecated but universally supported and needs no
    // separate worklet file; a question is short, so its cost is irrelevant.
    const node = ctx.createScriptProcessor(4096, 1, 1);
    const chunks = [];
    node.onaudioprocess = (e) => chunks.push(new Float32Array(e.inputBuffer.getChannelData(0)));
    source.connect(node);
    node.connect(ctx.destination);
    tutor.recorder = { stream, ctx, node, source, chunks, rate: ctx.sampleRate };
    $('tutor-mic').classList.add('recording');
    $('tutor-mic').title = 'Stop and send (Esc cancels)';
    setStatus('Listening… click the microphone again to send.');
  }

  function releaseRecorder() {
    const r = tutor.recorder;
    tutor.recorder = null;
    if (!r) return null;
    r.node.disconnect();
    r.source.disconnect();
    r.stream.getTracks().forEach(t => t.stop());
    r.ctx.close();
    $('tutor-mic').classList.remove('recording');
    $('tutor-mic').title = 'Ask by voice (click to start, click again to send)';
    return r;
  }

  function cancelRecording() {
    releaseRecorder();
    setStatus('');
  }

  async function finishRecording() {
    const r = releaseRecorder();
    if (!r) return;
    const samples = downsample(concat(r.chunks), r.rate, 16000);
    if (samples.length < 16000 * 0.3) { setStatus('That was too short to hear anything.', true); return; }
    setStatus('Transcribing…');
    $('tutor-mic').disabled = true;
    try {
      const res = await fetch(`/api/tutor/transcribe?model=${encodeURIComponent(tutor.model)}`, {
        method: 'POST',
        headers: { 'Content-Type': 'audio/wav' },
        body: encodeWav(samples, 16000),
      });
      if (!res.ok) throw new Error((await res.text()).trim() || `Transcription failed (${res.status})`);
      const { text } = await res.json();
      setStatus('');
      if (!text) { setStatus("Couldn't make out any words. Try again closer to the mic.", true); return; }
      send(text, { voice: true });
    } catch (e) {
      setStatus(e.message, true);
    } finally {
      $('tutor-mic').disabled = false;
    }
  }

  function concat(chunks) {
    const out = new Float32Array(chunks.reduce((n, c) => n + c.length, 0));
    let off = 0;
    chunks.forEach(c => { out.set(c, off); off += c.length; });
    return out;
  }

  // Averages each output sample's input window: a cheap low-pass that keeps
  // speech intelligible at 16 kHz.
  function downsample(input, from, to) {
    if (from === to) return input;
    const ratio = from / to;
    const out = new Float32Array(Math.floor(input.length / ratio));
    for (let i = 0; i < out.length; i++) {
      const start = Math.floor(i * ratio), end = Math.min(input.length, Math.floor((i + 1) * ratio));
      let sum = 0;
      for (let j = start; j < end; j++) sum += input[j];
      out[i] = sum / Math.max(1, end - start);
    }
    return out;
  }

  function encodeWav(samples, rate) {
    const buf = new ArrayBuffer(44 + samples.length * 2);
    const v = new DataView(buf);
    const str = (off, s) => { for (let i = 0; i < s.length; i++) v.setUint8(off + i, s.charCodeAt(i)); };
    str(0, 'RIFF'); v.setUint32(4, 36 + samples.length * 2, true); str(8, 'WAVE');
    str(12, 'fmt '); v.setUint32(16, 16, true); v.setUint16(20, 1, true); v.setUint16(22, 1, true);
    v.setUint32(24, rate, true); v.setUint32(28, rate * 2, true); v.setUint16(32, 2, true); v.setUint16(34, 16, true);
    str(36, 'data'); v.setUint32(40, samples.length * 2, true);
    for (let i = 0; i < samples.length; i++) {
      const s = Math.max(-1, Math.min(1, samples[i]));
      v.setInt16(44 + i * 2, s < 0 ? s * 0x8000 : s * 0x7fff, true);
    }
    return new Blob([buf], { type: 'audio/wav' });
  }

  // ---------------------------------------------------------------------
  // Entry points: the Tutor buttons appear only when a chat model exists.

  async function loadStatus() {
    if (!tutor.status) {
      try {
        const res = await fetch('/api/tutor/status');
        tutor.status = res.ok ? await res.json() : { available: false };
      } catch (e) {
        tutor.status = { available: false };
      }
    }
    return tutor.status;
  }

  async function init() {
    const st = await loadStatus();
    if (!st.available) return;
    $('tutor-toggle')?.classList.remove('hidden');
    $('practice-ask-tutor')?.classList.remove('hidden');
    $('tutor-toggle')?.addEventListener('click', togglePanel);
    $('practice-ask-tutor')?.addEventListener('click', () => openPanel());
  }

  // When the learner moves to another problem or chapter with the panel
  // open, follow them to that conversation.
  document.addEventListener('click', () => {
    if (tutor.open) setTimeout(() => { if (tutor.open && !tutor.busy) syncContext(); }, 400);
  });

  // Resetting practice also forgets the tutor's problem conversations: they
  // talk about attempts that no longer exist.
  document.addEventListener('dsatutor:practice-reset', () => {
    const all = load(HISTORY_KEY, {}) || {};
    Object.keys(all).filter(k => k.startsWith('problem:')).forEach(k => delete all[k]);
    save(HISTORY_KEY, all);
    if (tutor.key.startsWith('problem:')) {
      stopReply();
      tutor.messages = [];
      renderMessages();
    }
  });

  window.openTutor = openPanel;
  init();
})();
