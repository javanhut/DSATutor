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

  // Modes (interview.js, teacher.js) can take over the conversation: the
  // first one whose context() returns something decides the key, persona,
  // quick prompts and placeholder.
  window.tutorModes = window.tutorModes || [];

  function currentContext() {
    for (const mode of window.tutorModes) {
      const c = mode.context && mode.context();
      if (c) return c;
    }
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
    $('tutor-messages').addEventListener('click', (e) => {
      const answer = e.target.closest('[data-answer]');
      if (answer) {
        if (tutor.busy) return;
        answer.closest('.tutor-ask')?.querySelectorAll('.tutor-answer').forEach(b => b.classList.toggle('chosen', b === answer));
        send(answer.dataset.answer);
        return;
      }
      const chip = e.target.closest('[data-replay]');
      if (!chip || !chip.dataset.replay) return;
      const [mi, ai] = chip.dataset.replay.split(':').map(Number);
      const a = tutor.messages[mi]?.actions?.[ai];
      if (a) { clearSpotlights(); runAction(a); }
    });
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
    tutor.view = c;
    $('tutor-context').textContent = c.contextLine || (c.kind === 'problem' ? `${c.label} · sees your code and runs`
      : c.kind === 'sandbox' ? 'Sandbox · sees your code'
      : c.label);
    $('tutor-clear').classList.toggle('hidden', !!c.noClear);
    if (c.key !== tutor.key) {
      if (tutor.busy) stopReply();
      clearSpotlights();
      tutor.key = c.key;
      tutor.messages = loadConversation(c.key);
      renderMessages();
    }
    $('tutor-text').placeholder = c.placeholder || {
      problem: 'Ask about this problem, your code, complexity...',
      chapter: 'Ask about this chapter, or say "quiz me"...',
      sandbox: 'Ask about your sandbox code...',
      home: 'Ask anything about data structures and algorithms...',
    }[c.kind];
    // A quick prompt is [label, text] and sends at once; [label, text, true]
    // only fills the input, for prompts the learner finishes themselves.
    const quick = c.quick || QUICK_PROMPTS[c.kind] || [];
    $('tutor-quick').innerHTML = quick.map(([label], i) =>
      `<button type="button" class="tutor-chip" data-i="${i}">${escapeHtml(label)}</button>`
    ).join('');
    $('tutor-quick').onclick = (e) => {
      const chip = e.target.closest('.tutor-chip');
      if (!chip) return;
      const [, text, prefill] = quick[+chip.dataset.i];
      if (prefill) {
        const input = $('tutor-text');
        input.value = text;
        input.focus();
        input.setSelectionRange(text.length, text.length);
      } else {
        send(text);
      }
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
    if (!tutor.messages.some(m => !m.hidden)) {
      if (tutor.view && tutor.view.emptyHtml) {
        box.innerHTML = `<div class="tutor-empty">${tutor.view.emptyHtml}</div>`;
        return;
      }
      box.innerHTML = `<div class="tutor-empty">
        <p>Ask anything: why your code fails a test, how to start, what the complexity is, or to explain an idea again.</p>
        <p class="tutor-empty-note">The tutor sees what's on screen${tutor.key.startsWith('problem:') ? ', including your code and your last run' : ''}. It coaches with hints instead of handing you the answer, unless you ask for it.</p>
      </div>`;
      return;
    }
    box.innerHTML = tutor.messages.map((m, i) => (m.hidden ? '' : messageHtml(m, i))).join('');
    box.scrollTop = box.scrollHeight;
  }

  function messageHtml(m, idx) {
    const body = m.role === 'assistant' ? renderMarkdown(m.content) : `<p>${escapeHtml(m.content).replace(/\n/g, '<br>')}</p>`;
    const acts = m.actions || [];
    const chips = acts.map((a, ai) => (a.name === 'ask_learner' ? '' : chipHtml(a, `${idx}:${ai}`))).join('');
    const asks = acts.filter(a => a.name === 'ask_learner').map(askHtml).join('');
    const text = m.content === '(done)' && acts.length ? '' : `<div class="tutor-msg-text">${body}</div>`;
    return `<div class="tutor-msg ${m.role}">${chips ? `<div class="tutor-actions">${chips}</div>` : ''}${text}${asks}</div>`;
  }

  // A chip records one thing the tutor did. Highlights and scratchpad work
  // can be shown again; edits can't be replayed from history.
  const REPLAYABLE = ['highlight', 'scratchpad_code', 'scratchpad_visualize', 'scratchpad_sketch'];
  function chipHtml(a, ref) {
    const replay = REPLAYABLE.includes(a.name) ? ` data-replay="${ref}" title="Show it again"` : '';
    return `<button type="button" class="tutor-action-chip ${a.name}"${replay}>${escapeHtml(actionLabel(a))}</button>`;
  }

  // A question from ask_learner: clicking an answer sends it.
  function askHtml(a) {
    return `<div class="tutor-ask"><p>${escapeHtml(a.question)}</p><div class="tutor-ask-choices">${
      (a.choices || []).map(c => `<button type="button" class="tutor-answer" data-answer="${escapeHtml(c)}">${escapeHtml(c)}</button>`).join('')
    }</div></div>`;
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
    if (!opts.hidden) $('tutor-text').value = '';
    // Hidden messages cue the model (an interviewer starting the interview)
    // without showing up in the chat.
    tutor.messages.push(opts.hidden ? { role: 'user', content: text, hidden: true } : { role: 'user', content: text });
    renderMessages();
    stopSpeaking();
    if (typeof stopLessonAudio === 'function') stopLessonAudio();

    const box = $('tutor-messages');
    const bubble = document.createElement('div');
    bubble.className = 'tutor-msg assistant pending';
    bubble.innerHTML = '<div class="tutor-actions hidden"></div><div class="tutor-msg-text"><p class="tutor-typing">Thinking…</p></div><div class="tutor-asks"></div>';
    box.appendChild(bubble);
    const textEl = bubble.querySelector('.tutor-msg-text');
    const actionsEl = bubble.querySelector('.tutor-actions');
    const actions = []; // what the tutor did on the page during this reply
    clearSpotlights();
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
          messages: tutor.messages.map(m => ({ role: m.role, content: m.content })),
          context: { ...c.ctx, voice: speak, targets: targetList(), testCount: testCount(), scratchpad: window.dsaScratchpad?.state() || undefined },
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
            if (reply.trim()) textEl.innerHTML = renderMarkdown(reply);
            if (speech) speech.push(reply);
          } else if (ev.type === 'thinking') {
            thinking += ev.text;
            if (!reply.trim()) textEl.innerHTML = `<p class="tutor-typing">Thinking… <span class="tutor-think-count">${thinking.length} chars of reasoning</span></p>`;
          } else if (ev.type === 'action') {
            const done = runAction(ev);
            if (done && done.name === 'ask_learner') {
              actions.push(done);
              bubble.querySelector('.tutor-asks').insertAdjacentHTML('beforeend', askHtml(done));
            } else if (done) {
              actions.push(done);
              actionsEl.classList.remove('hidden');
              actionsEl.insertAdjacentHTML('beforeend', chipHtml(done, ''));
              if (done.undo) actionsEl.lastElementChild.insertAdjacentElement('afterend', undoButton(done.undo));
            }
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
    const saved = actions.map(({ undo, ...a }) => a); // undo handles don't survive a reload
    if (reply.trim() || actions.length) {
      tutor.messages.push({ role: 'assistant', content: reply.trim() || '(done)', actions: saved });
      if (speech) speech.end(reply);
    } else {
      bubble.remove();
      if (speech) stopSpeaking();
      if (!failed) tutor.messages.pop(); // stopped before any answer: drop the question too
    }
    if (failed) setStatus(failed, true);
    saveConversation();
    if (!bubble.isConnected || !(reply.trim() || actions.length)) renderMessages();
    else {
      if (reply.trim()) textEl.innerHTML = renderMarkdown(reply);
      else textEl.remove();
      if (!actionsEl.children.length) actionsEl.classList.add('hidden');
      // Replay chips need their message index now that it is saved.
      const idx = tutor.messages.length - 1;
      const chipActs = actions.map((x, ai) => [x, ai]).filter(([x]) => x.name !== 'ask_learner');
      actionsEl.querySelectorAll('.tutor-action-chip').forEach((chip, ci) => {
        if (chip.dataset.replay !== undefined && chipActs[ci]) chip.dataset.replay = `${idx}:${chipActs[ci][1]}`;
      });
      box.scrollTop = box.scrollHeight;
    }
  }

  // ---------------------------------------------------------------------
  // Acting on the page. The tutor points at regions by id; the ids and
  // their labels are sent with each message (targetList) and the server only
  // accepts those, so pageTargets is the whole vocabulary.

  const TAB_FOR_PANE = { 'problem-description': 'description', 'problem-hints': 'hints', 'visualization-tab': 'visualization' };
  const textOf = (node, n = 80) => (node ? node.innerText.replace(/\s+/g, ' ').trim().slice(0, n) : '');
  const hasText = (node) => !!node && node.innerText.trim() !== '';
  // A heading's own text, without the Listen/Skip buttons inside it.
  const headingText = (h) => {
    const c = h.cloneNode(true);
    c.querySelectorAll('button, select').forEach(n => n.remove());
    return textOf(c, 60);
  };

  function pageTargets() {
    const kind = currentContext().kind;
    const t = [];
    const add = (id, label, get, extra = {}) => { if (get()) t.push({ id, label, get, ...extra }); };
    if (kind === 'interview') {
      // An interviewer points at the candidate's code and nothing else: the
      // tests, hints and buttons aren't part of the interview.
      add('code', "The candidate's code document", () => $('practice-editor-container'), { lines: true, editor: 'practice-editor-container', textarea: 'practice-code' });
    } else if (kind === 'problem') {
      add('code', "The learner's code editor", () => $('practice-editor-container'), { lines: true, writable: true, editor: 'practice-editor-container', textarea: 'practice-code' });
      add('statement', 'Problem statement', () => $('problem-statement'));
      document.querySelectorAll('#problem-examples .example').forEach((ex, i) => {
        add(`example-${i + 1}`, `Example ${i + 1}: ${textOf(ex.querySelector('.example-io'))}`, () => document.querySelectorAll('#problem-examples .example')[i]);
      });
      if (hasText($('problem-constraints'))) add('constraints', 'Constraints', () => $('problem-constraints'));
      add('tests', 'Test cases panel: case buttons, input and expected output', () => document.querySelector('#practice-mode .results-panel'));
      if (hasText($('current-test-output'))) add('run-output', "Result of the learner's last run", () => $('current-test-output'));
      if (hasText($('test-results'))) add('submit-results', "Result of the learner's last submission", () => $('test-results'));
      if (hasText($('hints-list'))) add('hints', 'Revealed hints (Hints tab)', () => $('hints-list'));
      add('run-button', 'Run button (runs the selected test case)', () => $('practice-run'), { button: true });
      add('submit-button', 'Submit button (runs every test)', () => $('practice-submit'), { button: true });
    } else if (kind === 'sandbox') {
      add('code', "The learner's sandbox code editor", () => $('sandbox-editor-container'), { lines: true, writable: true, editor: 'sandbox-editor-container', textarea: 'sandbox-code' });
      if (hasText($('sandbox-output'))) add('output', 'Output of their last run', () => $('sandbox-output'));
    } else if (kind === 'chapter' || kind === 'teacher') {
      const cards = [...document.querySelectorAll('#standard-chapter-content :is(.overview-card, .tutorial-card, .lesson-card, .memory-card)')]
        .filter(card => card.offsetParent && card.querySelector('h2'));
      cards.forEach((card, i) => {
        const title = headingText(card.querySelector('h2'));
        add(`section-${i + 1}`, `Lesson section "${title}"`, () => card);
        card.querySelectorAll('.worked-step').forEach((step, j) => {
          add(`section-${i + 1}-step-${j + 1}`, `Worked example step ${j + 1} in "${title}": ${textOf(step.querySelector('.subgoal'), 60)}`, () => step);
        });
        card.querySelectorAll('.check-question').forEach((q, j) => {
          add(`section-${i + 1}-question-${j + 1}`, `Check question ${j + 1}: ${textOf(q.querySelector('.check-prompt'), 70)}`, () => q);
        });
        [...card.querySelectorAll('pre')].filter(p => !p.closest('.worked-step, .check-question')).slice(0, 6).forEach((pre, j) => {
          add(`section-${i + 1}-code-${j + 1}`, `Code block ${j + 1} in "${title}": ${textOf(pre, 50)}`, () => pre);
        });
      });
    }
    return t.slice(0, 60);
  }

  function targetList() {
    return pageTargets().map(({ id, label, lines, writable }) => ({ id, label, lines, writable }));
  }

  function testCount() {
    return document.querySelectorAll('#test-case-selector .test-case-btn').length;
  }

  function cmFor(containerId) {
    return (window.dsaEditors && window.dsaEditors[containerId]) || null;
  }

  // Everything highlighted during the current reply, so the next one starts
  // from a clean page.
  const spot = { nodes: [], lines: [], widgets: [], timers: [], keys: new Set() };

  function clearSpotlights() {
    spot.nodes.forEach(n => n.classList.remove('tutor-spotlight', 'tutor-spotlight-pulse'));
    spot.lines.forEach(({ cm, handle, cls }) => cm.removeLineClass(handle, 'background', cls));
    spot.widgets.forEach(w => w.clear ? w.clear() : w.remove());
    spot.timers.forEach(clearTimeout);
    document.querySelectorAll('.tutor-callout').forEach(n => n.remove());
    Object.assign(spot, { nodes: [], lines: [], widgets: [], timers: [], keys: new Set() });
  }

  // A pane inside a hidden problem tab has to be shown before it can be
  // pointed at.
  function reveal(node) {
    const pane = node.closest('.tab-content');
    if (pane && !pane.classList.contains('active') && TAB_FOR_PANE[pane.id]) {
      document.querySelector(`#practice-mode .tab-btn[data-tab="${TAB_FOR_PANE[pane.id]}"]`)?.click();
    }
  }

  function calloutNode(note) {
    const box = document.createElement('div');
    box.className = 'tutor-callout';
    box.innerHTML = `<span class="tutor-callout-label">Tutor</span><span class="tutor-callout-text"></span><button type="button" class="tutor-callout-close" aria-label="Dismiss">✕</button>`;
    box.querySelector('.tutor-callout-text').textContent = note;
    return box;
  }

  function highlightLines(cm, start, end, cls, note) {
    const last = cm.lastLine();
    const from = Math.max(0, Math.min(start - 1, last));
    const to = Math.max(from, Math.min(end - 1, last));
    for (let l = from; l <= to; l++) {
      const handle = cm.addLineClass(l, 'background', cls);
      spot.lines.push({ cm, handle, cls });
    }
    if (note) {
      const node = calloutNode(note);
      const widget = cm.addLineWidget(to, node, { coverGutter: false, noHScroll: true });
      node.querySelector('.tutor-callout-close').onclick = () => widget.clear();
      spot.widgets.push(widget);
    }
    cm.scrollIntoView({ from: { line: from, ch: 0 }, to: { line: to + (note ? 1 : 0), ch: 0 } }, 60);
    cm.getWrapperElement().scrollIntoView({ block: 'nearest', behavior: 'smooth' });
  }

  function actionLabel(a) {
    switch (a.name) {
      case 'highlight': {
        const where = a.start_line ? `lines ${a.start_line}${a.end_line && a.end_line !== a.start_line ? `–${a.end_line}` : ''} of your code` : (a.label || a.target);
        return `Pointed at ${where}`;
      }
      case 'write_code': return `Edited your code${a.summary ? `: ${a.summary}` : ''}`;
      case 'open_tab': return `Opened the ${a.tab} tab`;
      case 'scratchpad_code': return `Ran code in the scratchpad${a.explanation ? `: ${a.explanation}` : ''}`;
      case 'scratchpad_visualize': return `Drew “${a.title || 'a diagram'}” in the scratchpad`;
      case 'scratchpad_sketch': return `Sketched “${a.title || 'a drawing'}” on the whiteboard`;
      case 'select_test': return `Selected test case ${a.case}`;
    }
    return a.name;
  }

  // runAction carries out one action from the server. It returns the record
  // shown as a chip (with an undo handle for edits), or null when the page
  // has changed so much that the target is gone.
  function runAction(a) {
    const targets = pageTargets();
    const target = targets.find(t => t.id === a.target);
    switch (a.name) {
      case 'highlight': {
        if (!target) return null;
        // Models sometimes point at the same thing twice in one reply.
        const key = `${a.target}:${a.start_line || ''}:${a.end_line || ''}`;
        if (spot.keys.has(key)) return null;
        spot.keys.add(key);
        const node = target.get();
        reveal(node);
        const record = { name: 'highlight', target: a.target, label: target.label.split(':')[0], note: a.note, start_line: a.start_line, end_line: a.end_line };
        const cm = target.editor && cmFor(target.editor);
        if (cm && a.start_line) {
          highlightLines(cm, a.start_line, a.end_line || a.start_line, 'tutor-line-hl', a.note);
          return record;
        }
        node.classList.add('tutor-spotlight');
        spot.nodes.push(node);
        if (target.button) {
          node.classList.add('tutor-spotlight-pulse');
        } else if (a.note) {
          const call = calloutNode(a.note);
          call.querySelector('.tutor-callout-close').onclick = () => call.remove();
          node.insertAdjacentElement('afterend', call);
        }
        node.scrollIntoView({ block: 'center', behavior: 'smooth' });
        return record;
      }
      case 'write_code': {
        if (!target || !target.writable) return null;
        const cm = cmFor(target.editor);
        const ta = $(target.textarea);
        const before = cm ? cm.getValue() : (ta?.value || '');
        const after = a.code;
        if (cm) {
          // One replaceRange is one undo step.
          cm.replaceRange(after, { line: 0, ch: 0 }, { line: cm.lastLine(), ch: cm.getLine(cm.lastLine()).length }, '+tutor');
        }
        if (ta) ta.value = after;
        const gen = cm ? cm.changeGeneration(true) : 0;
        // Flash the lines that changed.
        const a1 = before.split('\n'), b1 = after.split('\n');
        let p = 0;
        while (p < a1.length && p < b1.length && a1[p] === b1[p]) p++;
        let q = 0;
        while (q < a1.length - p && q < b1.length - p && a1[a1.length - 1 - q] === b1[b1.length - 1 - q]) q++;
        if (cm && b1.length - q > p) highlightLines(cm, p + 1, b1.length - q, 'tutor-line-new', '');
        return {
          name: 'write_code', target: a.target, summary: a.summary,
          undo: () => {
            if (cm) {
              if (!cm.isClean(gen)) return "You've changed the code since; use Ctrl+Z / Cmd+Z in the editor.";
              cm.undo();
            } else if (ta) {
              ta.value = before;
            }
            return '';
          },
        };
      }
      case 'open_tab': {
        const btn = document.querySelector(`#practice-mode .tab-btn[data-tab="${a.tab}"]`);
        if (!btn || a.tab === 'solution') return null;
        btn.click();
        return { name: 'open_tab', tab: a.tab };
      }
      case 'scratchpad_code': {
        if (!window.dsaScratchpad) return null;
        window.dsaScratchpad.showCode(a);
        // The trace isn't kept in history; replaying runs the code again.
        return { name: 'scratchpad_code', code: a.code, explanation: a.explanation };
      }
      case 'scratchpad_visualize': {
        if (!window.dsaScratchpad) return null;
        window.dsaScratchpad.showFrames(a);
        return { name: 'scratchpad_visualize', title: a.title, frames: a.frames };
      }
      case 'scratchpad_sketch': {
        if (!window.dsaScratchpad) return null;
        window.dsaScratchpad.showSketch(a);
        return { name: 'scratchpad_sketch', title: a.title, elements: a.elements, captions: a.captions };
      }
      case 'ask_learner':
        return { name: 'ask_learner', question: a.question, choices: a.choices };
      case 'select_test': {
        const btn = $('test-case-selector')?.querySelector(`[data-index="${a.case - 1}"]`);
        if (!btn) return null;
        btn.click();
        btn.scrollIntoView({ block: 'nearest', behavior: 'smooth' });
        return { name: 'select_test', case: a.case };
      }
    }
    return null;
  }

  function undoButton(undo) {
    const btn = document.createElement('button');
    btn.type = 'button';
    btn.className = 'tutor-undo';
    btn.textContent = 'Undo';
    btn.onclick = () => {
      const problem = undo();
      if (problem) { setStatus(problem, true); return; }
      btn.textContent = 'Undone';
      btn.disabled = true;
    };
    return btn;
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

  // The API the modes use.
  window.dsaTutor = {
    open: openPanel,
    close: closePanel,
    send: (text, opts) => send(text, opts),
    sync: () => { if (tutor.open) syncContext(); },
    isBusy: () => tutor.busy,
    stop: stopReply,
    // The saved messages of a conversation, whether or not it is on screen.
    conversation: (key) => (key === tutor.key ? tutor.messages : loadConversation(key)).slice(),
    model: () => tutor.model,
    status: loadStatus,
  };
  window.openTutor = openPanel;
  init();
})();
