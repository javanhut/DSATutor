// Mock interview mode: a timed, Google-style coding interview on one
// practice problem, with the AI tutor as the interviewer.
//
// While an interview runs, the problem page looks like a shared document:
// Run and Submit, hints, the solution, the visualizer, related chapters and
// the LeetCode number are hidden, and the code starts from the bare
// signature. The candidate talks to the interviewer in the tutor panel (by
// text or microphone). Ending the interview runs the final code against
// every test and asks the model for a scorecard on Google's four coding
// criteria (internal/tutor/interview.go). Past interviews are kept in
// localStorage.
//
// Loaded after app.js, learn.js and tutor.js; it only appears when the tutor
// is available.

(function () {
  'use strict';

  const HISTORY_KEY = 'dsatutor_interviews';
  const MAX_HISTORY = 25;
  const LEVELS = ['L3', 'L4', 'L5'];
  const CRITERIA = [
    ['algorithms', 'Algorithms', 'approach, data structures, complexity'],
    ['coding', 'Coding', 'correct, clean code that handles edge cases'],
    ['communication', 'Communication', 'clarifying, explaining, thinking out loud'],
    ['problem_solving', 'Problem solving', 'brute force to optimal, testing, few hints'],
  ];
  const SCORE_WORDS = ['', 'Poor', 'Mixed', 'Good', 'Excellent'];

  const iv = {
    active: null, // { problemId, title, level, duration, startedAt, key }
    timer: null,
    ending: false,
  };

  const $ = (id) => document.getElementById(id);

  function load(key, fallback) {
    try { const v = localStorage.getItem(key); return v ? JSON.parse(v) : fallback; } catch (e) { return fallback; }
  }
  function save(key, value) {
    try { localStorage.setItem(key, JSON.stringify(value)); } catch (e) { /* full or private */ }
  }

  function elapsedSeconds() {
    return iv.active ? Math.floor((Date.now() - iv.active.startedAt) / 1000) : 0;
  }

  function fmtClock(sec) {
    const s = Math.max(0, sec);
    return `${Math.floor(s / 60)}:${String(s % 60).padStart(2, '0')}`;
  }

  function levelOf(problem) {
    const tag = (problem.tags || []).find(t => /^Google L[345]$/.test(t));
    if (tag) return tag.slice(-2);
    return { Easy: 'L3', Medium: 'L4', Hard: 'L5' }[problem.difficulty] || 'L4';
  }

  // ---------------------------------------------------------------------
  // The tutor mode: while an interview runs on the open problem, the tutor
  // panel is the interviewer.

  window.tutorModes = window.tutorModes || [];
  window.tutorModes.unshift({
    context() {
      const a = iv.active;
      const open = typeof practiceState !== 'undefined' && practiceState.currentProblem;
      if (!a || !open || practiceState.currentProblem.id !== a.problemId) return null;
      if ($('practice-mode')?.classList.contains('hidden')) return null;
      return {
        key: a.key,
        label: 'Mock interview',
        kind: 'interview',
        noClear: true,
        contextLine: `Mock interview · ${a.level} · the interviewer sees your document`,
        placeholder: 'Talk to your interviewer: clarify, explain your approach, think out loud…',
        emptyHtml: '<p>Your interviewer is joining…</p>',
        quick: [
          ['Clarifying question', 'I have a clarifying question: ', true],
          ['My approach', "Here's my approach: ", true],
          ['Done coding', "I'm done coding. Let me walk through an example to test it.", false],
          ['Ask for a hint', "I'm stuck. Could I get a small hint?", false],
        ],
        ctx: {
          view: 'problem',
          problemId: a.problemId,
          language: practiceState.language,
          code: getEditorCode('practice-code'),
          interview: { elapsedSeconds: elapsedSeconds(), durationSeconds: a.duration, level: a.level },
        },
      };
    },
  });

  // ---------------------------------------------------------------------
  // Modal dialogs

  function modal(html, { wide = false, onClose } = {}) {
    closeModal();
    const back = document.createElement('div');
    back.className = 'dsa-modal-backdrop';
    back.id = 'dsa-modal';
    back.innerHTML = `<div class="dsa-modal${wide ? ' wide' : ''}" role="dialog" aria-modal="true">${html}</div>`;
    document.body.appendChild(back);
    back.addEventListener('click', (e) => {
      if (e.target === back || e.target.closest('[data-close]')) { closeModal(); if (onClose) onClose(); }
    });
    return back.firstElementChild;
  }

  function closeModal() {
    $('dsa-modal')?.remove();
  }

  document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape' && $('dsa-modal') && !$('dsa-modal').dataset.locked) closeModal();
  });

  function lengthPicker(selected = 45) {
    return `<label class="iv-field"><span>Length</span><select id="iv-length">
      ${[30, 45, 60].map(m => `<option value="${m}"${m === selected ? ' selected' : ''}>${m} minutes</option>`).join('')}
    </select></label>`;
  }

  function levelPicker(selected) {
    return `<label class="iv-field"><span>Level</span><select id="iv-level">
      ${LEVELS.map(l => `<option value="${l}"${l === selected ? ' selected' : ''}>${l}${l === 'L3' ? ' (new grad)' : l === 'L4' ? ' (mid-level)' : ' (senior)'}</option>`).join('')}
    </select></label>`;
  }

  const RULES = `<ul class="iv-rules">
    <li>The AI tutor is your interviewer. Talk to it like a person: ask clarifying questions, explain your approach and its complexity before coding, and think out loud. The microphone works too.</li>
    <li>Like a Google Doc: <strong>no Run, no hints, no solution</strong>. Test by walking through examples by hand.</li>
    <li>When you finish (or time runs out), your code runs against every test, hidden ones included, and you get a scorecard on algorithms, coding, communication and problem solving.</li>
  </ul>`;

  function historyHtml() {
    const past = load(HISTORY_KEY, []) || [];
    if (!past.length) return '';
    return `<div class="iv-history"><h3>Past interviews</h3><ul>${past.slice(0, 6).map((h, i) => `
      <li><button type="button" class="iv-history-item" data-i="${i}">
        <span class="iv-verdict-chip ${verdictClass(h.scorecard?.verdict)}">${escapeHtml(h.scorecard?.verdict || 'No scorecard')}</span>
        <span>${escapeHtml(h.title)} · ${h.level}</span>
        <span class="iv-date">${new Date(h.startedAt).toLocaleDateString()}</span>
      </button></li>`).join('')}</ul></div>`;
  }

  function bindHistory(root) {
    root.querySelectorAll('.iv-history-item').forEach(btn => {
      btn.onclick = () => {
        const h = (load(HISTORY_KEY, []) || [])[+btn.dataset.i];
        if (h) showScorecard(h);
      };
    });
  }

  // From the problem list: pick a Google problem at a level.
  function openLauncher() {
    const box = modal(`
      <h2>Mock interview</h2>
      <p class="iv-lede">A timed coding interview on a problem Google candidates reported, with the tutor as your interviewer.</p>
      <div class="iv-fields">${levelPicker('L4')}${lengthPicker(45)}
        <label class="iv-check"><input type="checkbox" id="iv-unsolved" checked> Only problems I haven't solved</label>
      </div>
      ${RULES}
      <div class="iv-actions"><button class="btn-small ghost" data-close>Cancel</button><button class="btn-small primary" id="iv-go">Start interview</button></div>
      ${historyHtml()}`);
    bindHistory(box);
    $('iv-go').onclick = async () => {
      const level = $('iv-level').value;
      const duration = +$('iv-length').value * 60;
      const unsolved = $('iv-unsolved').checked;
      const pool = pickPool(level, unsolved);
      if (!pool.length) {
        alert(`No ${unsolved ? 'unsolved ' : ''}${level} problems left. Untick "Only problems I haven't solved" to repeat one.`);
        return;
      }
      const choice = pool[Math.floor(Math.random() * pool.length)];
      closeModal();
      await start(choice.id, level, duration);
    };
  }

  function pickPool(level, unsolved) {
    const progress = practiceState.progress || {};
    const ok = p => !unsolved || !progress[p.id]?.solved;
    let pool = practiceState.problems.filter(p => (p.tags || []).includes(`Google ${level}`) && ok(p));
    if (!pool.length) {
      const diff = { L3: 'Easy', L4: 'Medium', L5: 'Hard' }[level];
      pool = practiceState.problems.filter(p => p.difficulty === diff && ok(p));
    }
    return pool;
  }

  // From an open problem: interview on this one.
  function openForProblem() {
    const p = practiceState.currentProblem;
    if (!p) return;
    const box = modal(`
      <h2>Interview on “${escapeHtml(p.title)}”</h2>
      <p class="iv-lede">Your code resets to the bare function signature for the interview; your draft is kept and comes back afterwards.</p>
      <div class="iv-fields">${levelPicker(levelOf(p))}${lengthPicker(45)}</div>
      ${RULES}
      <div class="iv-actions"><button class="btn-small ghost" data-close>Cancel</button><button class="btn-small primary" id="iv-go">Start interview</button></div>
      ${historyHtml()}`);
    bindHistory(box);
    $('iv-go').onclick = async () => {
      const level = $('iv-level').value;
      const duration = +$('iv-length').value * 60;
      closeModal();
      await start(p.id, level, duration);
    };
  }

  // ---------------------------------------------------------------------
  // Running an interview

  async function start(problemId, level, duration) {
    if (!practiceState.currentProblem || practiceState.currentProblem.id !== problemId || $('practice-mode').classList.contains('hidden')) {
      if (typeof goToProblem === 'function') await goToProblem(problemId);
      else await openProblem(problemId);
    }
    const p = practiceState.currentProblem;
    if (!p || p.id !== problemId) return;

    // Keep the learner's draft; the interview starts from the signature.
    if (typeof saveDraft === 'function') saveDraft();
    const draftKeyNow = typeof draftKey === 'function' ? draftKey(p, practiceState.language) : null;
    const draft = draftKeyNow ? practiceState.drafts[draftKeyNow] : undefined;
    setEditorCode('practice-code', starterFor(p, practiceState.language));

    const startedAt = Date.now();
    iv.active = { problemId, title: p.title, level, duration, startedAt, key: `interview:${problemId}:${startedAt}`, draftKey: draftKeyNow, draft };
    iv.ending = false;
    document.body.classList.add('interview-active');
    document.querySelector('#practice-mode .tab-btn[data-tab="description"]')?.click();
    if (typeof stopLessonAudio === 'function') stopLessonAudio();
    mountBar();
    iv.timer = setInterval(tick, 1000);
    tick();

    await window.dsaTutor.open();
    window.dsaTutor.sync();
    window.dsaTutor.send('(The candidate has joined the interview. Greet them in one sentence and present the problem.)', { hidden: true });
  }

  function mountBar() {
    $('interview-bar')?.remove();
    const bar = document.createElement('div');
    bar.id = 'interview-bar';
    bar.className = 'interview-bar';
    bar.innerHTML = `
      <span class="iv-live" aria-hidden="true"></span>
      <span class="iv-bar-label">Interview · ${iv.active.level}</span>
      <span class="iv-clock" id="iv-clock" title="Time left">--:--</span>
      <button class="btn-small ghost" id="iv-talk" title="Open the conversation with your interviewer">Interviewer</button>
      <button class="btn-small primary" id="iv-end" title="Finish and get your scorecard">End interview</button>`;
    document.querySelector('#practice-mode .practice-header .header-controls')?.prepend(bar);
    $('iv-end').onclick = () => end(false);
    $('iv-talk').onclick = () => window.dsaTutor.open();
  }

  function tick() {
    if (!iv.active) return;
    const left = iv.active.duration - elapsedSeconds();
    const clock = $('iv-clock');
    if (clock) {
      clock.textContent = fmtClock(left);
      clock.classList.toggle('warn', left <= 10 * 60 && left > 5 * 60);
      clock.classList.toggle('urgent', left <= 5 * 60);
    }
    if (left <= 0 && !iv.ending) end(true);
  }

  function teardown() {
    clearInterval(iv.timer);
    iv.timer = null;
    document.body.classList.remove('interview-active');
    $('interview-bar')?.remove();
    const a = iv.active;
    iv.active = null;
    return a;
  }

  function restoreDraft(a) {
    // Put the learner's own draft back after the interview code is graded.
    const p = practiceState.currentProblem;
    if (!p || p.id !== a.problemId) return;
    if (a.draftKey) {
      if (a.draft === undefined) delete practiceState.drafts[a.draftKey];
      else practiceState.drafts[a.draftKey] = a.draft;
    }
    setEditorCode('practice-code', a.draft ?? starterFor(p, practiceState.language));
  }

  async function end(timeUp) {
    if (!iv.active || iv.ending) return;
    if (!timeUp && !confirm('End the interview and get your scorecard?')) return;
    iv.ending = true;
    window.dsaTutor.stop();
    const elapsed = Math.min(elapsedSeconds(), iv.active.duration);
    const code = getEditorCode('practice-code');
    const language = practiceState.language;
    const a = teardown();
    const transcript = window.dsaTutor.conversation(a.key)
      .filter(m => !m.hidden)
      .map(m => ({ role: m.role, content: m.content }));
    window.dsaTutor.sync();

    const box = modal(`
      <h2>${timeUp ? "Time's up" : 'Interview finished'}</h2>
      <p class="iv-lede">Running your code against every test and writing your feedback…</p>
      <div class="iv-spinner" aria-hidden="true"></div>`);
    box.parentElement.dataset.locked = '1';

    let result;
    try {
      const res = await fetch('/api/tutor/interview/score', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          model: window.dsaTutor.model(),
          problemId: a.problemId,
          language,
          code,
          messages: transcript,
          elapsedSeconds: elapsed,
          durationSeconds: a.duration,
          level: a.level,
        }),
      });
      if (!res.ok) throw new Error((await res.text()).trim() || `Scoring failed (${res.status})`);
      result = await res.json();
    } catch (e) {
      result = { error: e.message };
    }

    const record = {
      problemId: a.problemId,
      title: a.title,
      level: a.level,
      startedAt: a.startedAt,
      elapsed,
      duration: a.duration,
      code,
      tests: result.tests || null,
      scorecard: result.scorecard || null,
      error: result.error || '',
      model: result.model || '',
      transcript: transcript.slice(-60),
    };
    const past = load(HISTORY_KEY, []) || [];
    past.unshift(record);
    save(HISTORY_KEY, past.slice(0, MAX_HISTORY));

    if (result.tests && result.tests.status === 'Accepted' && typeof saveAcceptedSolution === 'function') {
      saveAcceptedSolution(a.problemId, code, language, 'interview');
    }
    if (result.tests) {
      const prev = practiceState.progress[a.problemId] || {};
      updateProblemProgress(a.problemId, {
        attempts: (prev.attempts || 0) + 1,
        solved: result.tests.status === 'Accepted' || !!prev.solved,
        lastAttempt: new Date().toISOString(),
      });
      renderProblemsList();
      updateProgressCount();
    }
    restoreDraft(a);
    if (typeof renderMySolutions === 'function') renderMySolutions();
    showScorecard(record, code);
  }

  function abandon() {
    const a = teardown();
    if (a) restoreDraft(a);
    window.dsaTutor.stop();
    window.dsaTutor.sync();
  }

  // Leaving the problem mid-interview asks first.
  document.addEventListener('click', (e) => {
    if (!iv.active || iv.ending) return;
    const leaving = e.target.closest('#practice-back, #practice-toggle, #sandbox-toggle, .chapter-item, .chapter-link, #home-start, #home-practice, #home-sandbox');
    if (!leaving) return;
    if (confirm('Abandon the interview? It will not be scored.')) {
      abandon();
    } else {
      e.preventDefault();
      e.stopImmediatePropagation();
    }
  }, true);

  window.addEventListener('beforeunload', (e) => {
    if (iv.active) { e.preventDefault(); e.returnValue = ''; }
  });

  // ---------------------------------------------------------------------
  // Scorecard

  function verdictClass(v) {
    if (!v) return 'none';
    if (/Strong Hire/.test(v)) return 'strong-hire';
    if (/Lean Hire/.test(v)) return 'lean-hire';
    if (/^Hire/.test(v)) return 'hire';
    if (/Lean No/.test(v)) return 'lean-no';
    return 'no-hire';
  }

  function showScorecard(rec) {
    const card = rec.scorecard;
    const t = rec.tests;
    const meta = [
      `${escapeHtml(rec.title)} · ${rec.level}`,
      `${fmtClock(rec.elapsed)} of ${Math.round(rec.duration / 60)} min`,
      t ? `${t.passed}/${t.total} tests · ${escapeHtml(t.status)}` : 'tests not run',
    ];
    if (card) meta.push(`${card.hints_given} hint${card.hints_given === 1 ? '' : 's'}`);

    const ratings = card ? CRITERIA.map(([key, name, what]) => {
      const r = card.ratings?.[key] || { score: 1, evidence: '' };
      return `<div class="iv-rating">
        <div class="iv-rating-head">
          <span class="iv-rating-name">${name}</span>
          <span class="iv-meter" aria-label="${r.score} of 4">${[1, 2, 3, 4].map(n => `<i class="${n <= r.score ? 'on' : ''}"></i>`).join('')}</span>
          <span class="iv-score-word">${SCORE_WORDS[r.score] || ''}</span>
        </div>
        <div class="iv-rating-what">${what}</div>
        <p>${escapeHtml(r.evidence || '')}</p>
      </div>`;
    }).join('') : '';

    const list = (items) => (items || []).map(x => `<li>${escapeHtml(x)}</li>`).join('');
    const failures = t && t.failures && t.failures.length
      ? `<div class="iv-section"><h3>Tests your code failed</h3><ul>${list(t.failures)}</ul></div>` : '';
    const transcript = (rec.transcript || []).map(m =>
      `<div class="iv-line ${m.role}"><strong>${m.role === 'assistant' ? 'Interviewer' : 'You'}:</strong> ${escapeHtml(m.content)}</div>`).join('');

    const box = modal(`
      <div class="iv-card-head">
        <span class="iv-verdict ${verdictClass(card?.verdict)}">${escapeHtml(card?.verdict || 'No scorecard')}</span>
        <div class="iv-meta">${meta.join(' · ')}</div>
      </div>
      ${card ? `<p class="iv-summary">${escapeHtml(card.summary)}</p>` : `<p class="iv-summary error">The scorecard couldn't be written: ${escapeHtml(rec.error || 'unknown error')}. Your test results are below.</p>`}
      ${ratings ? `<div class="iv-ratings">${ratings}</div>` : ''}
      ${card ? `<div class="iv-section"><h3>Complexity</h3><p>${escapeHtml(card.complexity)}${card.reached_optimal ? '' : ' <em>(not the optimal complexity)</em>'}</p></div>` : ''}
      ${card ? `<div class="iv-columns">
        <div class="iv-section"><h3>Strengths</h3><ul>${list(card.strengths)}</ul></div>
        <div class="iv-section"><h3>Do differently next time</h3><ul>${list(card.improvements)}</ul></div>
      </div>` : ''}
      ${failures}
      <details class="iv-section"><summary>Your final code</summary><pre class="tutor-code"><code>${escapeHtml(rec.code || '(none)')}</code></pre></details>
      <details class="iv-section"><summary>Transcript (${(rec.transcript || []).length} messages)</summary><div class="iv-transcript">${transcript || '<p>(empty)</p>'}</div></details>
      <p class="iv-note">Scored by ${escapeHtml(rec.model || 'the tutor model')}. It's a practice signal, not a real hiring decision.</p>
      <div class="iv-actions">
        <button class="btn-small ghost" id="iv-again">Interview again on this problem</button>
        <button class="btn-small primary" data-close>Done</button>
      </div>`, { wide: true });
    box.querySelector('#iv-again').onclick = async () => {
      closeModal();
      if (typeof goToProblem === 'function') await goToProblem(rec.problemId);
      openForProblem();
    };
  }

  // ---------------------------------------------------------------------
  // Entry points: shown only when the tutor works.

  async function init() {
    const st = await window.dsaTutor?.status();
    if (!st || !st.available) return;
    const listControls = document.querySelector('#practice-list .header-controls .control-group');
    if (listControls && !$('mock-interview')) {
      const b = document.createElement('button');
      b.id = 'mock-interview';
      b.className = 'btn-small primary';
      b.textContent = 'Mock interview';
      b.title = 'A timed Google-style interview with the AI tutor as interviewer';
      b.onclick = openLauncher;
      listControls.prepend(b);
    }
    const ask = $('practice-ask-tutor');
    if (ask && !$('practice-interview')) {
      const b = document.createElement('button');
      b.id = 'practice-interview';
      b.className = 'btn-small ghost';
      b.textContent = 'Interview mode';
      b.title = 'Practise this problem as a timed interview';
      b.onclick = openForProblem;
      ask.insertAdjacentElement('beforebegin', b);
    }
  }

  window.dsaInterview = { start, isActive: () => !!iv.active };
  init();
})();
