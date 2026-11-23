// Teach Me: ask to be taught any DSA topic, not tied to a chapter.
//
// The lesson follows what learning research says works (docs/teachme.md):
// activate prior knowledge, direct instruction in small chunks with a
// retrieval check after each, a worked example revealed one step at a time,
// faded practice, explain-it-back with feedback, then practice problems.
// Each lesson's flashcards join one review deck scheduled with growing gaps
// (spaced repetition); due cards from every topic come back mixed together.

(function () {
  'use strict';

  const DECK_KEY = 'dsatutor_teachme_deck';
  const HISTORY_KEY = 'dsatutor_teachme_history';
  const LEVEL_KEY = 'dsatutor_teachme_level';
  const MODEL_KEY = 'dsatutor_tutor_model'; // shared with the tutor panel
  const VOICE_KEY = 'dsatutor_teachme_voice'; // read each new part aloud
  const DAY = 24 * 60 * 60 * 1000;
  // Days until the next review for each box; a card climbs a box per
  // successful recall and drops to the first on a miss.
  const INTERVALS = [1, 3, 7, 16, 35, 80];
  const SUGGESTIONS = [
    'Monotonic stack', 'Sliding window', 'Union-Find', 'Dijkstra',
    'Top-K with a heap', 'Prefix sums', 'Topological sort', 'Binary search on the answer',
    'Backtracking', 'Memoization vs tabulation', 'Tries', 'Big-O of recursion',
  ];

  const root = document.getElementById('teachme-mode');
  const toggle = document.getElementById('teachme-toggle');
  if (!root || !toggle) return;

  const view = {
    active: false,
    status: null,      // /api/tutor/status
    lesson: null,      // the lesson on screen
    progress: null,    // what the learner has done in it
    loading: null,     // { topic, started, controller }
    review: null,      // { queue, index, shown, done }
    tts: false,        // a narration engine works (tts.js)
    speakSel: null,    // after the next render, read this element aloud
    drafts: {},        // unsent text in reply boxes, kept across renders
    walk: null,        // { pending, error } while the walkthrough is built
  };

  // ---- storage -----------------------------------------------------------

  function load(key, fallback) {
    try { const v = localStorage.getItem(key); return v ? JSON.parse(v) : fallback; } catch (e) { return fallback; }
  }
  function save(key, value) {
    try { localStorage.setItem(key, JSON.stringify(value)); } catch (e) { /* private mode or full */ }
  }

  const deck = () => load(DECK_KEY, []);
  const dueCards = () => deck().filter(c => c.due <= Date.now());

  function addCards(lesson) {
    const cards = deck();
    const have = new Set(cards.map(c => c.front));
    const firstReview = Date.now() + DAY; // spaced: the first review is tomorrow
    let added = 0;
    for (const c of lesson.flashcards || []) {
      if (have.has(c.front)) continue;
      cards.push({ id: `${Date.now().toString(36)}-${added}`, topic: lesson.title, front: c.front, back: c.back, box: 0, due: firstReview, reviews: 0, lapses: 0 });
      added++;
    }
    save(DECK_KEY, cards);
    return added;
  }

  function rateCard(id, rating) {
    const cards = deck();
    const c = cards.find(x => x.id === id);
    if (!c) return;
    c.reviews++;
    if (rating === 'again') {
      c.lapses++;
      c.box = 0;
      c.due = Date.now() + 10 * 60 * 1000; // back in this session or soon after
    } else {
      c.box = Math.min(INTERVALS.length - 1, Math.max(0, c.box + (rating === 'hard' ? 0 : rating === 'easy' ? 2 : 1)));
      const days = INTERVALS[c.box] * (rating === 'hard' ? 0.6 : 1);
      c.due = Date.now() + days * DAY;
    }
    save(DECK_KEY, cards);
  }

  function history() { return load(HISTORY_KEY, []); }
  function remember(entry) {
    const list = history().filter(h => h.id !== entry.id);
    list.unshift(entry);
    save(HISTORY_KEY, list.slice(0, 15));
  }
  function persistProgress() {
    if (!view.lesson) return;
    const list = history();
    const h = list.find(x => x.id === view.lesson.id);
    if (h) { h.progress = view.progress; save(HISTORY_KEY, list); }
  }

  // ---- helpers -------------------------------------------------------------

  function esc(s) {
    return String(s ?? '').replace(/[&<>"']/g, c => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' }[c]));
  }
  const md = s => (typeof renderMarkdown === 'function' ? renderMarkdown(s) : `<p>${esc(s)}</p>`);
  const norm = s => String(s || '').replace(/\s+/g, '').replace(/["']/g, "'").replace(/;$/, '').toLowerCase();
  // A sentence of model text with its `code` set as code.
  const inline = s => esc(s).replace(/`([^`]+)`/g, '<code>$1</code>');
  const plural = (n, w) => `${n} ${w}${n === 1 ? '' : 's'}`;

  function model() {
    const saved = load(MODEL_KEY, '');
    const models = view.status?.models || [];
    if (saved && models.some(m => m.name === saved)) return saved;
    return view.status?.default || '';
  }

  function whenDue(ms) {
    const d = ms - Date.now();
    if (d <= 0) return 'now';
    if (d < DAY) return `in ${Math.max(1, Math.round(d / 3600000))}h`;
    return `in ${Math.round(d / DAY)}d`;
  }

  // ---- voice ---------------------------------------------------------------
  //
  // Narration reuses tts.js: the same engines and voices (phonon, Piper, ...)
  // as the chapters, the same pickers, and Listen/Skip on every stage. With
  // "Read aloud" on, each new stage, piece of feedback and reply is read as
  // it appears.

  const voiceOn = () => view.tts && load(VOICE_KEY, true);
  const pickers = () => ['tts-voice', 'tts-speaker'].map(id => document.getElementById(id)).filter(Boolean);

  // The pickers belong to the chapter header and visit Teach Me's bar;
  // they go home before the view is redrawn, or innerHTML would delete them.
  function pickersHome() {
    if (pickers().some(p => root.contains(p)) && typeof placePickers === 'function') placePickers('toggle-edit');
  }

  function paint(html) {
    if (typeof stopLessonAudio === 'function' && ttsState.session && root.contains(ttsState.session.btn)) stopLessonAudio();
    pickersHome();
    root.innerHTML = html;
    if (!view.tts) return;
    if (root.querySelector('#tm-voice-anchor') && typeof placePickers === 'function') placePickers('tm-voice-anchor');
    root.querySelectorAll('.tm-stage').forEach(stage => {
      const host = stage.querySelector('.tm-listen');
      if (host && !host.querySelector('.tts-btn')) addListenButtons(host, () => stage);
    });
    root.querySelector('#tm-voice')?.addEventListener('change', e => {
      save(VOICE_KEY, e.target.checked);
      if (!e.target.checked && typeof stopLessonAudio === 'function') stopLessonAudio();
    });
  }

  // speakSoon reads el (a selector) aloud after the coming render, when
  // Read aloud is on.
  function speakSoon(sel) { view.speakSel = sel; }
  function speakPending() {
    const sel = view.speakSel;
    view.speakSel = null;
    if (!sel || !voiceOn() || typeof speakSection !== 'function') return;
    const target = typeof sel === 'function' ? sel() : root.querySelector(sel);
    const btn = target?.closest('.tm-stage')?.querySelector('.tm-listen .tts-btn');
    if (target && btn) speakSection(target, btn);
  }

  const voiceBar = () => view.tts ? `<span class="tm-voicebar">
      <span id="tm-voice-anchor" hidden></span>
      <label class="tm-voice-toggle"><input type="checkbox" id="tm-voice" ${load(VOICE_KEY, true) ? 'checked' : ''}> Read aloud</label>
    </span>` : '';

  // ---- showing and hiding the view ------------------------------------------

  function show() {
    if (typeof leaveModes === 'function') leaveModes();
    document.getElementById('detail')?.classList.add('hidden');
    document.getElementById('detail-empty')?.classList.add('hidden');
    document.querySelectorAll('.chapter-item').forEach(i => i.classList.remove('active'));
    root.classList.remove('hidden');
    toggle.classList.add('active');
    view.active = true;
    if (!view.status) fetchStatus();
    if (typeof ttsStatus === 'function') ttsStatus().then(st => {
      const had = view.tts;
      view.tts = !!st?.available;
      if (view.tts !== had && view.active) render();
    });
    render();
    root.scrollIntoView({ block: 'start' });
  }

  // hide puts the view away. restore brings back the chapter or home page;
  // when another mode is taking over, that mode shows its own page.
  function hide(restore) {
    if (!view.active) return;
    view.active = false;
    if (typeof stopLessonAudio === 'function') stopLessonAudio();
    pickersHome();
    root.classList.add('hidden');
    toggle.classList.remove('active');
    if (restore) {
      const current = typeof state !== 'undefined' && state.current;
      document.getElementById(current ? 'detail' : 'detail-empty')?.classList.remove('hidden');
    }
  }

  toggle.addEventListener('click', () => (view.active ? hide(true) : show()));
  // Any other way of changing the page takes over from Teach Me.
  document.addEventListener('click', e => {
    if (view.active && e.target.closest('#chapters .chapter-item, #practice-toggle, #sandbox-toggle, .wordmark')) hide(false);
  }, true);

  async function fetchStatus() {
    try {
      const r = await fetch('/api/tutor/status');
      view.status = await r.json();
    } catch (e) {
      view.status = { available: false, models: [], reason: 'Could not reach the DSA Tutor server.' };
    }
    if (view.active) render();
  }

  async function openPractice(id) {
    hide(false);
    const btn = document.getElementById('practice-toggle');
    if (btn && !btn.classList.contains('active') && typeof btn.onclick === 'function') await btn.onclick();
    if (typeof openProblem === 'function') openProblem(id);
  }

  // ---- generating a lesson ------------------------------------------------

  async function teach(topic) {
    topic = topic.trim();
    if (!topic || view.loading) return;
    const level = document.getElementById('tm-level')?.value || load(LEVEL_KEY, 'intermediate');
    save(LEVEL_KEY, level);
    const controller = new AbortController();
    view.loading = { topic, started: Date.now(), controller };
    view.review = null;
    render();
    const tick = setInterval(() => {
      const t = root.querySelector('.tm-elapsed');
      if (t && view.loading) t.textContent = `${Math.round((Date.now() - view.loading.started) / 1000)}s`;
    }, 1000);
    try {
      const r = await fetch('/api/teachme/lesson', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ model: model(), topic, level }),
        signal: controller.signal,
      });
      const data = await r.json().catch(() => ({}));
      if (!r.ok || !data.lesson) throw new Error(data.error || `The server answered ${r.status}.`);
      const lesson = { ...data.lesson, id: Date.now().toString(36), model: data.model, level, created: Date.now() };
      openLesson(lesson);
      speakSoon('.tm-stage');
      remember({ id: lesson.id, title: lesson.title, topic, created: lesson.created, lesson, progress: view.progress });
      view.error = null;
    } catch (e) {
      if (e.name !== 'AbortError') view.error = e.message;
    } finally {
      clearInterval(tick);
      view.loading = null;
      render();
      root.scrollIntoView({ block: 'start' });
    }
  }

  // ---- rendering -------------------------------------------------------------

  function render() {
    badge();
    if (!view.active) return;
    if (view.loading) return renderLoading();
    if (view.review) return renderReview();
    if (view.lesson) return renderLesson();
    renderHome();
  }

  function renderHome() {
    const st = view.status;
    const due = dueCards().length;
    const total = deck().length;
    const level = load(LEVEL_KEY, 'intermediate');
    const models = st?.models || [];
    const unavailable = st && !st.available;
    const recent = history();
    paint(`
      <div class="tm-page">
        <header class="tm-head">
          <p class="eyebrow">Teach Me</p>
          <h1>What do you want to learn?</h1>
          <p class="tm-lede">Name any data structure, algorithm or pattern. You get a short lesson: the idea explained, a check after each part, a worked example, a line of code to fill in, and problems to practise. The flashcards come back for review over the next few weeks.</p>
        </header>
        ${unavailable ? `<p class="tm-notice">Teach Me needs Ollama with a chat model. ${esc(st.reason || '')}</p>` : ''}
        ${view.error ? `<p class="tm-notice tm-error">${esc(view.error)}</p>` : ''}
        <form class="tm-ask" id="tm-form">
          <label class="tm-label" for="tm-topic">Teach me…</label>
          <div class="tm-ask-row">
            <input id="tm-topic" type="text" autocomplete="off" maxlength="300" placeholder="e.g. how a monotonic stack finds the next greater element" ${unavailable ? 'disabled' : ''}>
            <button type="submit" class="primary" ${unavailable ? 'disabled' : ''}>Teach me</button>
          </div>
          <div class="tm-options">
            <label>Level
              <select id="tm-level">
                ${['beginner', 'intermediate', 'advanced'].map(l => `<option value="${l}" ${l === level ? 'selected' : ''}>${l[0].toUpperCase() + l.slice(1)}</option>`).join('')}
              </select>
            </label>
            ${models.length > 1 ? `<label>Model
              <select id="tm-model">${models.map(m => `<option ${m.name === model() ? 'selected' : ''}>${esc(m.name)}</option>`).join('')}</select>
            </label>` : ''}
          </div>
          <div class="tm-chips">${SUGGESTIONS.map(s => `<button type="button" class="tm-chip" data-topic="${esc(s)}" ${unavailable ? 'disabled' : ''}>${esc(s)}</button>`).join('')}</div>
        </form>

        <section class="tm-panel tm-review-panel">
          <div>
            <h2>Spaced review</h2>
            <p>${total ? `${plural(due, 'card')} due of ${total}. ${due ? 'Recall them before you learn something new: it makes both stick.' : nextDueText()}` : 'Finish a lesson and its flashcards land here, scheduled for tomorrow, then at growing gaps.'}</p>
          </div>
          <button id="tm-review" class="${due ? 'primary' : ''}" ${due ? '' : 'disabled'}>Review ${due ? due : ''} now</button>
        </section>

        ${recent.length ? `<section class="tm-recent">
          <h2>Your lessons</h2>
          <ul>${recent.map(h => `<li><button class="tm-link" data-open="${esc(h.id)}">${esc(h.title)}</button><span>${new Date(h.created).toLocaleDateString()}</span></li>`).join('')}</ul>
        </section>` : ''}
      </div>`);

    const input = root.querySelector('#tm-topic');
    root.querySelector('#tm-form').addEventListener('submit', e => { e.preventDefault(); teach(input.value); });
    root.querySelectorAll('.tm-chip').forEach(b => b.addEventListener('click', () => teach(b.dataset.topic)));
    root.querySelector('#tm-model')?.addEventListener('change', e => save(MODEL_KEY, e.target.value));
    root.querySelector('#tm-review')?.addEventListener('click', startReview);
    root.querySelectorAll('[data-open]').forEach(b => b.addEventListener('click', () => {
      const h = history().find(x => x.id === b.dataset.open);
      if (!h) return;
      openLesson(h.lesson, h.progress);
      render();
    }));
    if (!unavailable) input.focus();
  }

  function nextDueText() {
    const next = deck().reduce((m, c) => Math.min(m, c.due), Infinity);
    return next === Infinity ? '' : `Next card due ${whenDue(next)}.`;
  }

  function renderLoading() {
    paint(`
      <div class="tm-page">
        <header class="tm-head">
          <p class="eyebrow">Teach Me</p>
          <h1>${esc(view.loading.topic)}</h1>
        </header>
        <div class="tm-panel tm-loading">
          <span class="tm-spinner" aria-hidden="true"></span>
          <div>
            <p><strong>Writing your lesson…</strong> <span class="tm-elapsed">0s</span></p>
            <p class="tm-muted">While you wait: what do you already know about this? Say it to yourself. Recalling first, even if you're wrong, makes the lesson stick better.</p>
          </div>
          <button id="tm-cancel">Cancel</button>
        </div>
      </div>`);
    root.querySelector('#tm-cancel').addEventListener('click', () => view.loading?.controller.abort());
  }

  // The lesson unlocks one stage at a time: you answer before you move on.
  // Stages: warm-up, the parts, the worked example, for an algorithm a
  // line-by-line walkthrough of its code, faded practice, explain it back,
  // and wrap-up.
  const isAlgorithm = l => l.kind === 'algorithm' && !!l.code;

  function stages(l) {
    const list = ['warmup'];
    l.chunks.forEach((_, i) => list.push(`chunk-${i}`));
    if (l.worked_example?.steps?.length) list.push('worked');
    if (isAlgorithm(l)) list.push('walk');
    if (l.faded?.code) list.push('faded');
    if (l.explain_prompt) list.push('explain');
    list.push('wrap');
    return list;
  }

  function newProgress() {
    return { stage: 0, checks: {}, steps: 0, faded: null, fadedTries: 0, replies: {}, walkStep: 0, cardsAdded: 0 };
  }

  function openLesson(lesson, progress) {
    view.lesson = lesson;
    view.progress = { ...newProgress(), ...(progress || {}) };
    view.progress.replies ||= {};
    view.walk = null;
  }

  // Ollama answers one request at a time, so the walkthrough (about 30 s)
  // is built while the learner is on the last part or the worked example,
  // not at the start where it would hold up the warm-up reply.
  function maybeFetchWalk() {
    const l = view.lesson;
    if (!isAlgorithm(l) || l.walkthrough?.length || view.walk) return;
    const list = stages(l);
    if (view.progress.stage < list.indexOf('walk') - 2) return;
    if (Object.values(view.progress.replies).some(r => r.pending)) return;
    fetchWalk();
  }

  function persistLesson() {
    if (!view.lesson) return;
    const list = history();
    const h = list.find(x => x.id === view.lesson.id);
    if (h) { h.lesson = view.lesson; h.progress = view.progress; save(HISTORY_KEY, list); }
  }

  // The walkthrough is built in the background, so it's usually ready by
  // the time the learner gets to it.
  async function fetchWalk() {
    const l = view.lesson;
    view.walk = { pending: true };
    try {
      const r = await fetch('/api/teachme/walk', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ model: model(), topic: l.title, code: l.code }),
      });
      const data = await r.json().catch(() => ({}));
      if (!r.ok || !data.steps?.length) throw new Error(data.error || `The server answered ${r.status}.`);
      l.walkthrough = data.steps;
      if (view.lesson === l) view.walk = null;
      persistLesson();
    } catch (e) {
      if (view.lesson === l) view.walk = { error: e.message };
    }
    if (view.lesson === l && view.active && !view.loading && !view.review && currentStage() === 'walk') {
      if (l.walkthrough?.length) speakSoon('.tm-walk-step');
      renderLesson();
    }
  }

  const currentStage = () => {
    const list = stages(view.lesson);
    return list[Math.min(view.progress.stage, list.length - 1)];
  };

  function renderLesson() {
    const l = view.lesson;
    const p = view.progress;
    const list = stages(l);
    const upTo = Math.min(p.stage, list.length - 1);
    const pct = Math.round((upTo / (list.length - 1)) * 100);
    const parts = list.slice(0, upTo + 1).map((s, i) => renderStage(s, i === upTo));
    paint(`
      <div class="tm-page tm-lesson">
        <div class="tm-bar">
          <button class="tm-link" id="tm-back">← Teach me something else</button>
          ${voiceBar()}
          <span class="tm-muted">${esc(l.level || '')}${l.model ? ` · ${esc(l.model)}` : ''}</span>
        </div>
        <header class="tm-head">
          <p class="eyebrow">Teach Me</p>
          <h1>${esc(l.title)}</h1>
          <div class="tm-progress" role="progressbar" aria-valuenow="${pct}" aria-valuemin="0" aria-valuemax="100"><span style="width:${pct}%"></span></div>
        </header>
        ${parts.join('')}
      </div>`);
    bindLesson();
    speakPending();
    maybeFetchWalk();
  }

  function next() {
    view.progress.stage++;
    persistProgress();
    speakSoon(() => { const all = root.querySelectorAll('.tm-stage'); return all[all.length - 1]; });
    renderLesson();
    const last = root.querySelectorAll('.tm-stage');
    last[last.length - 1]?.scrollIntoView({ behavior: 'smooth', block: 'start' });
  }

  const continueBtn = (label = 'Continue') => `<div class="tm-actions"><button class="primary tm-next">${label}</button></div>`;
  const kicker = text => `<div class="tm-kicker-row"><p class="tm-kicker">${text}</p><span class="tm-listen"></span></div>`;

  // ---- replies to what the learner types -----------------------------------
  //
  // Anything typed is answered: each claim confirmed or corrected, a reply,
  // and, until the answer is solid, a follow-up question in new words. A
  // gated box holds the lesson there until the answer is solid; "Explain it
  // differently" asks for a new angle and a new question rather than
  // letting the learner skip past.

  function replyState(key) {
    return (view.progress.replies[key] ||= { turns: [] });
  }
  const isSolid = key => !!view.progress.replies[key]?.solid;
  // A gated line offers a way out after this many real attempts ("Explain
  // it differently" doesn't count). Taking it shows the answer and puts the
  // line in the review deck, so it still comes back.
  const ESCAPE_TRIES = 3;
  const tries = key => (view.progress.replies[key]?.turns || []).filter(t => !t.lost).length;
  const escaped = key => !!view.progress.replies[key]?.escaped;

  const CLAIM_MARK = { correct: '✓', partly: '~', wrong: '✗' };
  const VERDICT = { solid: 'You’ve got it', partial: 'Partly there', missing: 'Not yet' };

  // box: { key, prompt, placeholder, gate, warmup }
  function renderReplyBox(box, current) {
    const st = view.progress.replies[box.key] || { turns: [] };
    const last = st.turns[st.turns.length - 1];
    const ask = last?.feedback?.follow_up || '';
    const open = current && !st.solid && !st.escaped;
    const turns = st.turns.map((t, i) => `<div class="tm-exchange${i === st.turns.length - 1 ? ' tm-latest' : ''}">
        <div class="tm-you"><span>You</span>${esc(t.lost ? 'I don’t get it yet. Explain it differently.' : t.answer)}</div>
        ${t.feedback ? `<div class="tm-grade tm-${esc(t.feedback.verdict)}">
          ${t.feedback.claims?.length ? `<ul class="tm-claims">${t.feedback.claims.map(c => `<li class="tm-claim tm-claim-${esc(c.verdict)}"><span class="tm-mark" aria-label="${esc(c.verdict)}">${CLAIM_MARK[c.verdict] || '~'}</span> <em>${inline(c.claim)}</em>${c.note ? `: ${inline(c.note)}` : ''}</li>`).join('')}</ul>` : ''}
          <p><strong>${box.warmup && t.feedback.verdict !== 'solid' ? 'Good start' : VERDICT[t.feedback.verdict] || ''}.</strong> ${inline(t.feedback.reply)}</p>
          ${t.feedback.missing?.length && t.feedback.verdict !== 'solid' ? `<p class="tm-muted">Still to cover: ${t.feedback.missing.map(m => inline(m.replace(/[.\s]+$/, ''))).join('; ')}.</p>` : ''}
          ${t.feedback.follow_up && (box.gate || i === st.turns.length - 1) ? `<p class="tm-followup"><span>${box.gate ? 'Try this' : 'Think about'}</span>${inline(t.feedback.follow_up)}</p>` : ''}
        </div>` : ''}
      </div>`).join('');
    const draft = view.drafts[box.key] || '';
    return `<div class="tm-reply" data-reply="${esc(box.key)}">
      ${turns}
      ${st.error ? `<p class="tm-notice tm-error">${esc(st.error)}</p>` : ''}
      ${open ? `<textarea class="tm-input" data-draft="${esc(box.key)}" rows="3" placeholder="${esc(ask ? 'Answer the new question…' : box.placeholder || 'In your own words…')}" ${st.pending ? 'disabled' : ''}>${esc(draft)}</textarea>
        <div class="tm-actions">
          ${box.gate ? `<button data-lost="${esc(box.key)}" ${st.pending ? 'disabled' : ''}>${st.turns.length ? 'Explain it differently' : 'I don\u2019t get it yet'}</button>` : ''}
          <button class="${box.gate ? 'primary' : ''}" data-send="${esc(box.key)}" ${st.pending ? 'disabled' : ''}>${st.pending ? 'Reading your answer…' : box.sendLabel || 'Send'}</button>
        </div>` : ''}
    </div>`;
  }

  // What each reply box is about: the question, what was just taught and
  // what a good answer covers.
  function replyContext(key) {
    const l = view.lesson;
    const lessonText = [l.big_idea, ...l.chunks.map(c => `${c.heading}: ${c.explanation}`)].join('\n');
    if (key === 'warmup') return { kind: 'warmup', prompt: l.warmup, reference: l.big_idea, context: '' };
    if (key === 'explain') return { kind: 'explain', prompt: l.explain_prompt, reference: l.explain_reference, context: lessonText };
    if (key.startsWith('say-')) {
      const c = l.chunks[Number(key.slice(5))];
      return { kind: 'chunk', prompt: 'Say this part back in your own words, or ask about it.', reference: c.explanation, context: `${c.heading}: ${c.explanation}` };
    }
    if (key.startsWith('w')) {
      const s = l.walkthrough[Number(key.slice(1))];
      const numbered = l.code.split('\n').map((x, i) => `${i + 1}: ${x}`).join('\n');
      return {
        kind: 'line', prompt: s.question, reference: `${s.answer} ${s.why}`,
        context: `The code:\n${numbered}\n\nThe line(s) being taught, ${s.from}-${s.to}:\n${s.code}\nWhat it does: ${s.what}\nWhy: ${s.why}`,
      };
    }
    return { kind: 'chunk', prompt: '', reference: '', context: lessonText };
  }

  async function sendReply(key, lost) {
    const l = view.lesson;
    const st = replyState(key);
    const answer = lost ? 'I don’t understand this yet. Please explain it differently, then ask me again.' : (view.drafts[key] || '').trim();
    if (!answer || st.pending) return root.querySelector(`[data-draft="${key}"]`)?.focus();
    const ctx = replyContext(key);
    const last = st.turns[st.turns.length - 1];
    st.pending = true;
    st.error = null;
    renderLesson();
    try {
      const r = await fetch('/api/teachme/respond', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          model: model(), topic: l.title, kind: ctx.kind, context: ctx.context,
          prompt: last?.feedback?.follow_up || ctx.prompt, reference: ctx.reference,
          answer, attempt: st.turns.length + 1,
        }),
      });
      const data = await r.json().catch(() => ({}));
      if (!r.ok || !data.feedback) throw new Error(data.error || `The server answered ${r.status}.`);
      st.turns.push({ answer: lost ? '' : answer, lost: !!lost, feedback: data.feedback });
      if (data.feedback.verdict === 'solid' && !lost) st.solid = true;
      delete view.drafts[key];
      speakSoon(`[data-reply="${key}"] .tm-latest .tm-grade`);
    } catch (e) {
      st.error = `Couldn't get a reply: ${e.message}`;
    }
    st.pending = false;
    persistProgress();
    if (view.lesson === l) renderLesson();
  }

  // ---- stages ----------------------------------------------------------------

  function renderStage(stage, current) {
    const l = view.lesson;
    const p = view.progress;
    if (stage === 'warmup') {
      const st = p.replies.warmup;
      return `<section class="tm-stage">
        ${kicker('Warm up · 30 seconds')}
        <p class="tm-prompt">${esc(l.warmup)}</p>
        ${current || st?.turns?.length ? renderReplyBox({ key: 'warmup', placeholder: 'Your best guess. Wrong is fine: trying first helps.', sendLabel: 'Check my guess', warmup: true }, current) : ''}
        ${!current && p.warmup && !st?.turns?.length ? `<p class="tm-muted">You guessed: ${esc(p.warmup)}</p>` : ''}
        ${current ? continueBtn('Start the lesson') : ''}
        ${!current ? `<div class="tm-bigidea"><span>The big idea</span>${md(l.big_idea)}</div>` : ''}
      </section>`;
    }
    if (stage.startsWith('chunk-')) {
      const i = Number(stage.slice(6));
      const c = l.chunks[i];
      const right = checkRight(c.check, `c${i}`);
      return `<section class="tm-stage">
        ${kicker(`Part ${i + 1} of ${l.chunks.length}`)}
        <h2>${esc(c.heading)}</h2>
        <div class="tm-prose">${md(c.explanation)}</div>
        ${renderCheck(c.check, `c${i}`)}
        ${right && (current || p.replies[`say-${i}`]?.turns?.length) ? `<div class="tm-say">
            <div class="tm-label">Say it back or ask about it (optional)</div>
            ${renderReplyBox({ key: `say-${i}`, placeholder: 'Put this part in your own words, or ask what confused you.' }, current)}
          </div>` : ''}
        ${current && right ? continueBtn() : ''}
      </section>`;
    }
    if (stage === 'worked') {
      const w = l.worked_example;
      const shown = Math.min(p.steps || 0, w.steps.length);
      const done = shown >= w.steps.length;
      return `<section class="tm-stage">
        ${kicker('Worked example')}
        <div class="tm-prose">${md(w.problem)}</div>
        <ol class="tm-steps">${w.steps.slice(0, shown).map(s => `<li><div>${md(s.do)}</div><p class="tm-why">${esc(s.why)}</p></li>`).join('')}</ol>
        ${current && !done ? `<div class="tm-actions"><span class="tm-muted">${shown === 0 ? 'Try the first step yourself, then compare.' : 'What comes next? Decide before you reveal it.'}</span><button class="primary" id="tm-step">Reveal step ${shown + 1} of ${w.steps.length}</button></div>` : ''}
        ${done && l.code && !isAlgorithm(l) ? `<details class="tm-code" ${current ? 'open' : ''}><summary>Reference code${l.complexity ? ` · ${esc(l.complexity)}` : ''}</summary><pre class="md-code"><code>${esc(l.code)}</code></pre></details>` : ''}
        ${done && (!l.code || isAlgorithm(l)) && l.complexity ? `<p class="tm-muted">Cost: ${esc(l.complexity)}</p>` : ''}
        ${current && done ? continueBtn(isAlgorithm(l) ? 'Go through the code line by line' : 'Now you try') : ''}
      </section>`;
    }
    if (stage === 'walk') return renderWalk(current);
    if (stage === 'faded') {
      const f = l.faded;
      const solved = p.faded === 'right' || p.faded === 'revealed';
      const [before, after] = f.code.split('____');
      const blank = solved ? `<span class="tm-filled">${esc(f.answer)}</span>`
        : `<input class="tm-blank" id="tm-blank" spellcheck="false" autocomplete="off" size="${Math.max(6, f.answer.length + 2)}" aria-label="Fill in the blank">`;
      return `<section class="tm-stage">
        ${kicker('Your turn: fill in the blank')}
        <p class="tm-prompt">${esc(f.prompt)}</p>
        <pre class="md-code tm-faded"><code>${esc(before)}${blank}${esc(after)}</code></pre>
        ${!solved ? `<div class="tm-actions">
            ${p.fadedTries ? `<span class="tm-feedback tm-wrong">Not quite. Hint: ${esc(f.hint)}</span>` : ''}
            ${p.fadedTries >= 2 ? '<button id="tm-reveal">Show me</button>' : ''}
            <button class="primary" id="tm-check-blank">Check</button>
          </div>`
          : `<p class="tm-feedback ${p.faded === 'right' ? 'tm-right' : 'tm-wrong'}">${p.faded === 'right' ? 'Right.' : 'Here it is. Cover it and try it again tomorrow.'}</p>`}
        ${current && solved ? continueBtn() : ''}
      </section>`;
    }
    if (stage === 'explain') {
      const st = p.replies.explain;
      const n = tries('explain');
      const out = escaped('explain');
      return `<section class="tm-stage">
        ${kicker('Explain it back')}
        <p class="tm-prompt">${esc(l.explain_prompt)}</p>
        ${renderReplyBox({ key: 'explain', gate: true, placeholder: 'In your own words, as if to a friend. No peeking.', sendLabel: 'Get feedback' }, current)}
        ${current && !st?.solid && !out && n ? '<p class="tm-muted">Keep going: the lesson finishes once your explanation covers what happens and why.</p>' : ''}
        ${out ? `<div class="tm-grade tm-missing tm-escaped">
            <p><strong>A model answer.</strong></p>
            <div class="tm-prose">${md(l.explain_reference)}</div>
            <p class="tm-muted">This question is in your review deck now: it comes back tomorrow.</p>
          </div>`
          : st?.solid || n >= 2 ? `<details class="tm-code"><summary>Model answer</summary><div class="tm-prose">${md(l.explain_reference)}</div></details>` : ''}
        ${current && !st?.solid && !out && n >= ESCAPE_TRIES ? `<div class="tm-actions tm-escape">
            <span class="tm-muted">Still stuck after ${n} tries? You can see the answer and review it later.</span>
            <button id="tm-explain-escape">Show me the answer and finish</button>
          </div>` : ''}
        ${current && (st?.solid || out) ? continueBtn('Finish') : ''}
      </section>`;
    }
    return renderWrap();
  }

  // Line by line: the whole code with the current lines lit, what they do,
  // why they're written that way, and a question that must be answered
  // solidly before the next line.
  function renderWalk(current) {
    const l = view.lesson;
    const p = view.progress;
    const steps = l.walkthrough || [];
    if (!steps.length && !current) return `<section class="tm-stage">${kicker('Line by line')}<p class="tm-muted">Walkthrough skipped.</p></section>`;
    if (!steps.length) {
      const failed = view.walk?.error;
      return `<section class="tm-stage">
        ${kicker('Line by line')}
        ${failed ? `<p class="tm-notice tm-error">Couldn't build the walkthrough: ${esc(failed)}</p>
            <div class="tm-actions"><button id="tm-walk-skip">Skip the walkthrough</button><button class="primary" id="tm-walk-retry">Try again</button></div>`
          : `<div class="tm-loading"><span class="tm-spinner" aria-hidden="true"></span><p>Preparing a line-by-line walkthrough of the code…</p></div>`}
      </section>`;
    }
    const at = Math.min(p.walkStep || 0, steps.length - 1);
    const finished = (p.walkStep || 0) >= steps.length;
    const s = steps[at];
    const lines = l.code.split('\n');
    const lineState = n => {
      if (finished) return 'done';
      if (n >= s.from && n <= s.to) return 'now';
      return n < s.from ? 'done' : 'later';
    };
    const code = lines.map((x, i) => `<span class="tm-line tm-line-${lineState(i + 1)}"><span class="tm-ln">${i + 1}</span>${esc(x) || ' '}</span>`).join('');
    const range = st => st.from === st.to ? `Line ${st.from}` : `Lines ${st.from}–${st.to}`;
    const doneSteps = steps.slice(0, finished ? steps.length : at);
    return `<section class="tm-stage tm-walk">
      ${kicker(`Line by line · ${finished ? 'done' : `step ${at + 1} of ${steps.length}`}`)}
      <pre class="md-code tm-walk-code"><code>${code}</code></pre>
      ${doneSteps.length ? `<details class="tm-walk-done"><summary>${finished ? 'Every line, explained' : `Lines you've got (${doneSteps.length})`}</summary>
          <ol>${doneSteps.map(st => `<li><p><strong>${range(st)}</strong> <code class="tm-line-code">${esc(st.code)}</code></p><p>${md(st.what)}</p><p class="tm-why">${esc(st.why)}</p></li>`).join('')}</ol>
        </details>` : ''}
      ${!finished ? `<div class="tm-walk-step">
          <p class="tm-line-head"><strong>${range(s)}</strong> <code class="tm-line-code">${esc(s.code)}</code></p>
          <div class="tm-prose"><p><span class="tm-tag">What</span> ${inline(s.what)}</p><p><span class="tm-tag">Why</span> ${inline(s.why)}</p></div>
          <p class="tm-q"><span>Check</span>${inline(s.question)}</p>
          ${renderReplyBox({ key: `w${at}`, gate: true, placeholder: 'Answer in your own words: what and why.', sendLabel: 'Check my answer' }, current)}
          ${escaped(`w${at}`) ? `<div class="tm-grade tm-missing tm-escaped">
              <p><strong>The answer.</strong> ${inline(s.answer)}</p>
              <p class="tm-why">${inline(s.why)}</p>
              <p class="tm-muted">This line is in your review deck now: it comes back tomorrow.</p>
            </div>` : ''}
          ${current && !isSolid(`w${at}`) && !escaped(`w${at}`) && tries(`w${at}`) >= ESCAPE_TRIES ? `<div class="tm-actions tm-escape">
              <span class="tm-muted">Still stuck after ${tries(`w${at}`)} tries? You can see the answer and review this line later.</span>
              <button id="tm-walk-escape">Show me the answer and move on</button>
            </div>` : ''}
          ${current && (isSolid(`w${at}`) || escaped(`w${at}`)) ? `<div class="tm-actions"><button class="primary" id="tm-walk-next">${at + 1 < steps.length ? 'Next line' : 'That’s every line'}</button></div>` : ''}
        </div>` : ''}
      ${current && finished ? continueBtn(l.faded?.code ? 'Now you try' : 'Continue') : ''}
    </section>`;
  }


  // A check is tried until it's right: a wrong pick shows the misconception
  // behind it and is struck out, and the answer is never given away.
  const picksOf = key => {
    const v = view.progress.checks[key];
    return Array.isArray(v) ? v : v === undefined ? [] : [v];
  };
  const checkRight = (c, key) => !c?.choices?.length || picksOf(key).includes(c.answer);

  function renderCheck(c, key) {
    if (!c || !c.choices?.length) return '';
    const picks = picksOf(key);
    const right = picks.includes(c.answer);
    const lastWrong = !right && picks.length ? picks[picks.length - 1] : null;
    return `<div class="tm-check" data-key="${key}">
      <p class="tm-q"><span>Check</span>${inline(c.question)}</p>
      <div class="tm-choices">${c.choices.map((ch, i) => {
        let cls = '';
        if (right && i === c.answer) cls = 'is-right';
        else if (picks.includes(i)) cls = 'is-wrong';
        return `<button class="tm-choice ${cls}" data-i="${i}" ${right || picks.includes(i) ? 'disabled' : ''}>${esc(ch)}</button>`;
      }).join('')}</div>
      <div class="tm-said">
      ${right ? `<p class="tm-feedback tm-right"><strong>${picks.length === 1 ? 'Right.' : 'Now you have it.'}</strong> ${inline(c.why)}</p>`
        : lastWrong !== null ? `<p class="tm-feedback tm-wrong"><strong>Not quite.</strong> ${inline(c.feedback?.[lastWrong] || '')} Try again.</p>` : ''}
      </div>
    </div>`;
  }

  function renderWrap() {
    const l = view.lesson;
    const p = view.progress;
    const keys = l.chunks.map((c, i) => c.check?.choices?.length ? `c${i}` : null).filter(Boolean);
    const right = keys.filter(k => picksOf(k)[0] === l.chunks[Number(k.slice(1))].check.answer).length;
    if (!p.cardsAdded && !p.cardsDone) {
      p.cardsAdded = addCards(l);
      p.cardsDone = true;
      persistProgress();
    }
    return `<section class="tm-stage tm-wrap">
      ${kicker('Done')}
      <h2>Lock it in</h2>
      <ul class="tm-score">
        ${escaped('explain') ? '<li><strong>Explain-it-back</strong> to revisit in review</li>' : ''}
        ${keys.length ? `<li><strong>${right}/${keys.length}</strong> checks right first time</li>` : ''}
        ${l.walkthrough?.length && isAlgorithm(l) ? (() => {
          const out = l.walkthrough.filter((_, i) => escaped(`w${i}`)).length;
          return `<li><strong>${l.walkthrough.length - out}/${l.walkthrough.length}</strong> steps of code explained back</li>${out ? `<li><strong>${out}</strong> ${out === 1 ? 'line' : 'lines'} to revisit in review</li>` : ''}`;
        })() : ''}
        ${l.faded?.code ? `<li><strong>${p.faded === 'right' ? 'Solved' : 'Shown'}</strong> fill-in-the-blank</li>` : ''}
        <li><strong>${l.flashcards?.length || 0}</strong> flashcards scheduled, first review tomorrow</li>
      </ul>
      ${l.pitfalls?.length ? `<h3>Mistakes people make</h3><ul class="tm-pitfalls">${l.pitfalls.map(x => `<li>${md(x)}</li>`).join('')}</ul>` : ''}
      ${l.problems?.length ? `<h3>Practise it now</h3>
        <p class="tm-muted">Solve one while it's fresh, then another in a day or two. Struggling a bit is the point.</p>
        <ul class="tm-problems">${l.problems.map(pr => `<li><button class="tm-problem" data-problem="${esc(pr.id)}"><span>${esc(pr.title)}</span><span class="tm-diff tm-${esc(String(pr.difficulty).toLowerCase())}">${esc(pr.difficulty)}</span></button></li>`).join('')}</ul>`
        : '<p class="tm-muted">No practice problem in the bank matches this topic yet; try the Sandbox to write it from memory.</p>'}
      <div class="tm-actions"><button id="tm-again">Teach me something else</button>${dueCards().length ? `<button class="primary" id="tm-review-now">Review ${dueCards().length} due cards</button>` : ''}</div>
    </section>`;
  }

  function bindLesson() {
    const l = view.lesson;
    const p = view.progress;
    root.querySelector('#tm-back')?.addEventListener('click', home);
    root.querySelector('#tm-again')?.addEventListener('click', home);
    root.querySelector('#tm-review-now')?.addEventListener('click', startReview);
    root.querySelectorAll('.tm-next').forEach(b => b.addEventListener('click', () => {
      if (view.drafts.warmup) p.warmup = view.drafts.warmup.trim();
      next();
    }));
    root.querySelectorAll('.tm-check').forEach(box => {
      box.querySelectorAll('.tm-choice').forEach(b => b.addEventListener('click', () => {
        p.checks[box.dataset.key] = [...picksOf(box.dataset.key), Number(b.dataset.i)];
        persistProgress();
        speakSoon(`.tm-check[data-key="${box.dataset.key}"] .tm-said`);
        renderLesson();
      }));
    });
    root.querySelectorAll('[data-draft]').forEach(t => t.addEventListener('input', () => { view.drafts[t.dataset.draft] = t.value; }));
    root.querySelectorAll('[data-draft]').forEach(t => t.addEventListener('keydown', e => {
      if (e.key === 'Enter' && (e.metaKey || e.ctrlKey)) sendReply(t.dataset.draft);
    }));
    root.querySelectorAll('[data-send]').forEach(b => b.addEventListener('click', () => sendReply(b.dataset.send)));
    root.querySelectorAll('[data-lost]').forEach(b => b.addEventListener('click', () => sendReply(b.dataset.lost, true)));
    root.querySelector('#tm-step')?.addEventListener('click', () => {
      p.steps = (p.steps || 0) + 1;
      persistProgress();
      speakSoon('.tm-steps li:last-child');
      renderLesson();
    });
    root.querySelector('#tm-walk-next')?.addEventListener('click', () => {
      p.walkStep = (p.walkStep || 0) + 1;
      persistProgress();
      speakSoon(p.walkStep < (l.walkthrough?.length || 0) ? '.tm-walk-step' : null);
      renderLesson();
      root.querySelector('.tm-walk-step, .tm-walk .tm-next')?.scrollIntoView({ behavior: 'smooth', block: 'center' });
    });
    root.querySelector('#tm-walk-escape')?.addEventListener('click', () => {
      const at = p.walkStep || 0;
      const s = l.walkthrough[at];
      const range = s.from === s.to ? `line ${s.from}` : `lines ${s.from}–${s.to}`;
      replyState(`w${at}`).escaped = true;
      addCards({ title: l.title, flashcards: [{ front: `${l.title}, ${range}: \`${s.code.trim()}\`\n\n${s.question}`, back: `${s.answer}\n\n${s.why}` }] });
      persistProgress();
      badge();
      speakSoon('.tm-escaped');
      renderLesson();
    });
    root.querySelector('#tm-explain-escape')?.addEventListener('click', () => {
      replyState('explain').escaped = true;
      addCards({ title: l.title, flashcards: [{ front: `${l.title}: ${l.explain_prompt}`, back: l.explain_reference }] });
      persistProgress();
      badge();
      speakSoon('section.tm-stage:last-of-type .tm-escaped');
      renderLesson();
    });
    root.querySelector('#tm-walk-retry')?.addEventListener('click', () => { fetchWalk(); renderLesson(); });
    root.querySelector('#tm-walk-skip')?.addEventListener('click', () => { p.walkStep = 9999; next(); });
    const blank = root.querySelector('#tm-blank');
    const checkBlank = () => {
      if (!blank.value.trim()) return blank.focus();
      if (norm(blank.value) === norm(l.faded.answer)) p.faded = 'right';
      else p.fadedTries = (p.fadedTries || 0) + 1;
      persistProgress();
      renderLesson();
      root.querySelector('#tm-blank')?.focus();
    };
    root.querySelector('#tm-check-blank')?.addEventListener('click', checkBlank);
    blank?.addEventListener('keydown', e => { if (e.key === 'Enter') checkBlank(); });
    root.querySelector('#tm-reveal')?.addEventListener('click', () => { p.faded = 'revealed'; persistProgress(); renderLesson(); });
    root.querySelectorAll('[data-problem]').forEach(b => b.addEventListener('click', () => openPractice(b.dataset.problem)));
  }

  function home() {
    view.drafts = {};
    view.walk = null;
    view.lesson = null;
    view.progress = null;
    view.review = null;
    view.error = null;
    render();
  }

  // ---- spaced review ----------------------------------------------------------

  function startReview() {
    // Shuffled, so cards from different topics are interleaved.
    const queue = dueCards().map(c => c.id).sort(() => Math.random() - 0.5);
    if (!queue.length) return;
    view.review = { queue, index: 0, shown: false, done: 0, again: 0 };
    render();
  }

  function renderReview() {
    const r = view.review;
    const card = deck().find(c => c.id === r.queue[r.index]);
    if (!card) {
      paint(`
        <div class="tm-page">
          <header class="tm-head"><p class="eyebrow">Spaced review</p><h1>All caught up</h1></header>
          <section class="tm-panel"><p>You recalled ${plural(r.done, 'card')}${r.again ? `, ${r.again} needed another go` : ''}. ${nextDueText()}</p></section>
          <div class="tm-actions"><button class="primary" id="tm-done">Back to Teach Me</button></div>
        </div>`);
      root.querySelector('#tm-done').addEventListener('click', home);
      return;
    }
    paint(`
      <div class="tm-page">
        <div class="tm-bar"><button class="tm-link" id="tm-back">← Stop reviewing</button><span class="tm-muted">${r.index + 1} of ${r.queue.length}</span></div>
        <header class="tm-head"><p class="eyebrow">Spaced review · ${esc(card.topic)}</p></header>
        <section class="tm-card">
          <div class="tm-prose tm-front">${md(card.front)}</div>
          ${r.shown ? `<div class="tm-prose tm-back">${md(card.back)}</div>`
            : '<textarea class="tm-input" id="tm-recall" rows="2" placeholder="Answer from memory first (typing it is optional)"></textarea>'}
        </section>
        <div class="tm-actions tm-rate">
          ${r.shown ? `<span class="tm-muted">How well did you recall it?</span>
            <button data-rate="again">Again</button><button data-rate="hard">Hard</button><button class="primary" data-rate="good">Good</button><button data-rate="easy">Easy</button>`
            : '<button class="primary" id="tm-show">Show answer</button>'}
        </div>
      </div>`);
    root.querySelector('#tm-back').addEventListener('click', home);
    root.querySelector('#tm-show')?.addEventListener('click', () => { r.shown = true; renderReview(); });
    root.querySelector('#tm-recall')?.focus();
    root.querySelectorAll('[data-rate]').forEach(b => b.addEventListener('click', () => {
      const rating = b.dataset.rate;
      rateCard(card.id, rating);
      badge();
      if (rating === 'again') { r.queue.push(card.id); r.again++; } else r.done++;
      r.index++;
      r.shown = false;
      renderReview();
    }));
  }

  document.addEventListener('keydown', e => {
    if (!view.active || !view.review || e.target.matches('input, textarea')) return;
    const keys = { 1: 'again', 2: 'hard', 3: 'good', 4: 'easy' };
    if (e.key === ' ' && !view.review.shown) { e.preventDefault(); root.querySelector('#tm-show')?.click(); }
    else if (view.review.shown && keys[e.key]) root.querySelector(`[data-rate="${keys[e.key]}"]`)?.click();
  });

  // Let the sidebar show the number of due cards.
  function badge() {
    const n = dueCards().length;
    toggle.querySelector('.tm-badge')?.remove();
    if (n) toggle.insertAdjacentHTML('beforeend', `<span class="tm-badge" title="${plural(n, 'card')} due for review">${n}</span>`);
  }
  badge();
  setInterval(badge, 60 * 1000);
})();
