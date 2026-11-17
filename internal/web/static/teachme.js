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
    render();
    root.scrollIntoView({ block: 'start' });
  }

  // hide puts the view away. restore brings back the chapter or home page;
  // when another mode is taking over, that mode shows its own page.
  function hide(restore) {
    if (!view.active) return;
    view.active = false;
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
      view.lesson = lesson;
      view.progress = { stage: 0, checks: {}, steps: 0, faded: null, fadedTries: 0, explain: null, cardsAdded: 0 };
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
    root.innerHTML = `
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
      </div>`;

    const input = root.querySelector('#tm-topic');
    root.querySelector('#tm-form').addEventListener('submit', e => { e.preventDefault(); teach(input.value); });
    root.querySelectorAll('.tm-chip').forEach(b => b.addEventListener('click', () => teach(b.dataset.topic)));
    root.querySelector('#tm-model')?.addEventListener('change', e => save(MODEL_KEY, e.target.value));
    root.querySelector('#tm-review')?.addEventListener('click', startReview);
    root.querySelectorAll('[data-open]').forEach(b => b.addEventListener('click', () => {
      const h = history().find(x => x.id === b.dataset.open);
      if (!h) return;
      view.lesson = h.lesson;
      view.progress = h.progress || { stage: 0, checks: {}, steps: 0 };
      render();
    }));
    if (!unavailable) input.focus();
  }

  function nextDueText() {
    const next = deck().reduce((m, c) => Math.min(m, c.due), Infinity);
    return next === Infinity ? '' : `Next card due ${whenDue(next)}.`;
  }

  function renderLoading() {
    root.innerHTML = `
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
      </div>`;
    root.querySelector('#tm-cancel').addEventListener('click', () => view.loading?.controller.abort());
  }

  // The lesson unlocks one stage at a time: you answer before you move on.
  // Stages: 0 warm-up, 1..n chunks, then worked example, faded practice,
  // explain it back, and wrap-up.
  function stages(l) {
    const list = ['warmup'];
    l.chunks.forEach((_, i) => list.push(`chunk-${i}`));
    if (l.worked_example?.steps?.length) list.push('worked');
    if (l.faded?.code) list.push('faded');
    if (l.explain_prompt) list.push('explain');
    list.push('wrap');
    return list;
  }

  function renderLesson() {
    const l = view.lesson;
    const p = view.progress;
    const list = stages(l);
    const upTo = Math.min(p.stage, list.length - 1);
    const pct = Math.round((upTo / (list.length - 1)) * 100);
    const parts = list.slice(0, upTo + 1).map((s, i) => renderStage(s, i === upTo));
    root.innerHTML = `
      <div class="tm-page tm-lesson">
        <div class="tm-bar">
          <button class="tm-link" id="tm-back">← Teach me something else</button>
          <span class="tm-muted">${esc(l.level || '')}${l.model ? ` · ${esc(l.model)}` : ''}</span>
        </div>
        <header class="tm-head">
          <p class="eyebrow">Teach Me</p>
          <h1>${esc(l.title)}</h1>
          <div class="tm-progress" role="progressbar" aria-valuenow="${pct}" aria-valuemin="0" aria-valuemax="100"><span style="width:${pct}%"></span></div>
        </header>
        ${parts.join('')}
      </div>`;
    bindLesson();
  }

  function next() {
    view.progress.stage++;
    persistProgress();
    renderLesson();
    const last = root.querySelectorAll('.tm-stage');
    last[last.length - 1]?.scrollIntoView({ behavior: 'smooth', block: 'start' });
  }

  const continueBtn = (label = 'Continue') => `<div class="tm-actions"><button class="primary tm-next">${label}</button></div>`;

  function renderStage(stage, current) {
    const l = view.lesson;
    const p = view.progress;
    if (stage === 'warmup') {
      return `<section class="tm-stage">
        <p class="tm-kicker">Warm up · 30 seconds</p>
        <p class="tm-prompt">${esc(l.warmup)}</p>
        ${current ? `<textarea class="tm-input" id="tm-warmup" rows="2" placeholder="Your best guess (optional, nobody grades this)">${esc(p.warmup || '')}</textarea>${continueBtn('Start the lesson')}`
          : p.warmup ? `<p class="tm-muted">You guessed: ${esc(p.warmup)}</p>` : ''}
        ${!current || p.stage > 0 ? `<div class="tm-bigidea"><span>The big idea</span>${md(l.big_idea)}</div>` : ''}
      </section>`;
    }
    if (stage.startsWith('chunk-')) {
      const i = Number(stage.slice(6));
      const c = l.chunks[i];
      return `<section class="tm-stage">
        <p class="tm-kicker">Part ${i + 1} of ${l.chunks.length}</p>
        <h2>${esc(c.heading)}</h2>
        <div class="tm-prose">${md(c.explanation)}</div>
        ${renderCheck(c.check, `c${i}`, current)}
      </section>`;
    }
    if (stage === 'worked') {
      const w = l.worked_example;
      const shown = Math.min(p.steps || 0, w.steps.length);
      const done = shown >= w.steps.length;
      return `<section class="tm-stage">
        <p class="tm-kicker">Worked example</p>
        <div class="tm-prose">${md(w.problem)}</div>
        <ol class="tm-steps">${w.steps.slice(0, shown).map(s => `<li><div>${md(s.do)}</div><p class="tm-why">${esc(s.why)}</p></li>`).join('')}</ol>
        ${current && !done ? `<div class="tm-actions"><span class="tm-muted">${shown === 0 ? 'Try the first step yourself, then compare.' : 'What comes next? Decide before you reveal it.'}</span><button class="primary" id="tm-step">Reveal step ${shown + 1} of ${w.steps.length}</button></div>` : ''}
        ${done && l.code ? `<details class="tm-code" ${current ? 'open' : ''}><summary>Reference code${l.complexity ? ` · ${esc(l.complexity)}` : ''}</summary><pre class="md-code"><code>${esc(l.code)}</code></pre></details>` : ''}
        ${done && !l.code && l.complexity ? `<p class="tm-muted">Cost: ${esc(l.complexity)}</p>` : ''}
        ${current && done ? continueBtn('Now you try') : ''}
      </section>`;
    }
    if (stage === 'faded') {
      const f = l.faded;
      const solved = p.faded === 'right' || p.faded === 'revealed';
      const [before, after] = f.code.split('____');
      const blank = solved ? `<span class="tm-filled">${esc(f.answer)}</span>`
        : `<input class="tm-blank" id="tm-blank" spellcheck="false" autocomplete="off" size="${Math.max(6, f.answer.length + 2)}" aria-label="Fill in the blank">`;
      return `<section class="tm-stage">
        <p class="tm-kicker">Your turn: fill in the blank</p>
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
      const ex = p.explain;
      return `<section class="tm-stage">
        <p class="tm-kicker">Explain it back</p>
        <p class="tm-prompt">${esc(l.explain_prompt)}</p>
        <textarea class="tm-input" id="tm-explain" rows="4" placeholder="In your own words, as if to a friend. No peeking." ${ex?.feedback ? 'readonly' : ''}>${esc(ex?.answer || '')}</textarea>
        ${ex?.error ? `<p class="tm-notice tm-error">${esc(ex.error)}</p>` : ''}
        ${ex?.feedback ? `<div class="tm-grade tm-${esc(ex.feedback.verdict)}">
            <p><strong>${{ solid: 'Solid', partial: 'Partly there', missing: 'Not yet' }[ex.feedback.verdict] || ''}.</strong> ${esc(ex.feedback.feedback)}</p>
            ${ex.feedback.missing?.length ? `<ul>${ex.feedback.missing.map(m => `<li>${esc(m)}</li>`).join('')}</ul>` : ''}
          </div>` : ''}
        ${ex?.feedback || ex?.skipped ? `<details class="tm-code"><summary>Model answer</summary><div class="tm-prose">${md(l.explain_reference)}</div></details>` : ''}
        ${current ? (ex?.feedback || ex?.skipped ? continueBtn('Finish')
          : `<div class="tm-actions"><button id="tm-skip-explain">Skip</button><button class="primary" id="tm-grade" ${ex?.pending ? 'disabled' : ''}>${ex?.pending ? 'Reading your answer…' : 'Get feedback'}</button></div>`) : ''}
      </section>`;
    }
    return renderWrap();
  }

  function renderCheck(c, key, current) {
    if (!c || !c.choices?.length) return current ? continueBtn() : '';
    const picked = view.progress.checks[key];
    const answered = picked !== undefined;
    return `<div class="tm-check" data-key="${key}">
      <p class="tm-q"><span>Check</span>${esc(c.question)}</p>
      <div class="tm-choices">${c.choices.map((ch, i) => {
        let cls = '';
        if (answered && i === c.answer) cls = 'is-right';
        else if (answered && i === picked) cls = 'is-wrong';
        return `<button class="tm-choice ${cls}" data-i="${i}" ${answered ? 'disabled' : ''}>${esc(ch)}</button>`;
      }).join('')}</div>
      ${answered ? (picked === c.answer
        ? `<p class="tm-feedback tm-right"><strong>Right.</strong> ${esc(c.why)}</p>`
        : `<p class="tm-feedback tm-wrong"><strong>Not quite.</strong> ${esc(c.feedback?.[picked] || '')}</p>
           <p class="tm-feedback"><strong class="tm-right">${esc(c.choices[c.answer])}</strong> is right: ${esc(c.why)}</p>`) : ''}
      ${current && answered ? continueBtn() : ''}
    </div>`;
  }

  function renderWrap() {
    const l = view.lesson;
    const p = view.progress;
    const keys = l.chunks.map((c, i) => c.check?.choices?.length ? `c${i}` : null).filter(Boolean);
    const right = keys.filter(k => p.checks[k] === l.chunks[Number(k.slice(1))].check.answer).length;
    if (!p.cardsAdded && !p.cardsDone) {
      p.cardsAdded = addCards(l);
      p.cardsDone = true;
      persistProgress();
    }
    return `<section class="tm-stage tm-wrap">
      <p class="tm-kicker">Done</p>
      <h2>Lock it in</h2>
      <ul class="tm-score">
        ${keys.length ? `<li><strong>${right}/${keys.length}</strong> checks right</li>` : ''}
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
      const w = root.querySelector('#tm-warmup');
      if (w) p.warmup = w.value.trim();
      next();
    }));
    root.querySelectorAll('.tm-check').forEach(box => {
      box.querySelectorAll('.tm-choice').forEach(b => b.addEventListener('click', () => {
        p.checks[box.dataset.key] = Number(b.dataset.i);
        persistProgress();
        renderLesson();
      }));
    });
    root.querySelector('#tm-step')?.addEventListener('click', () => {
      p.steps = (p.steps || 0) + 1;
      persistProgress();
      renderLesson();
    });
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
    root.querySelector('#tm-skip-explain')?.addEventListener('click', () => { p.explain = { skipped: true }; persistProgress(); renderLesson(); });
    root.querySelector('#tm-grade')?.addEventListener('click', () => grade(root.querySelector('#tm-explain').value));
    root.querySelectorAll('[data-problem]').forEach(b => b.addEventListener('click', () => openPractice(b.dataset.problem)));
  }

  async function grade(answer) {
    const l = view.lesson;
    const p = view.progress;
    if (!answer.trim()) return root.querySelector('#tm-explain')?.focus();
    p.explain = { answer, pending: true };
    renderLesson();
    try {
      const r = await fetch('/api/teachme/feedback', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ model: model(), topic: l.title, prompt: l.explain_prompt, reference: l.explain_reference, answer }),
      });
      const data = await r.json().catch(() => ({}));
      if (!r.ok || !data.feedback) throw new Error(data.error || `The server answered ${r.status}.`);
      p.explain = { answer, feedback: data.feedback };
    } catch (e) {
      p.explain = { answer, error: `Couldn't get feedback: ${e.message}` };
    }
    persistProgress();
    if (view.lesson === l) renderLesson();
  }

  function home() {
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
      root.innerHTML = `
        <div class="tm-page">
          <header class="tm-head"><p class="eyebrow">Spaced review</p><h1>All caught up</h1></header>
          <section class="tm-panel"><p>You recalled ${plural(r.done, 'card')}${r.again ? `, ${r.again} needed another go` : ''}. ${nextDueText()}</p></section>
          <div class="tm-actions"><button class="primary" id="tm-done">Back to Teach Me</button></div>
        </div>`;
      root.querySelector('#tm-done').addEventListener('click', home);
      return;
    }
    root.innerHTML = `
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
      </div>`;
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
