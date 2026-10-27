// Learning features layered on the chapter and practice views: the lesson
// sequence on chapter pages (worked examples, completion problems, retrieval
// questions, practice ladder), spaced review of solved problems, and
// interleaved mixed practice. docs/pedagogy.md explains the research behind
// each piece.
//
// Loaded before app.js; shares its globals (el, escapeHtml, state,
// practiceState, openProblem, ...), which exist by the time anything here runs.

// ---------------------------------------------------------------------------
// Small helpers
// ---------------------------------------------------------------------------

// renderMarkdown handles the subset of Markdown used in content: fenced code,
// inline code, bold, paragraphs and "- " lists. Input is escaped first.
function renderMarkdown(text) {
  if (!text) return '';
  const blocks = [];
  let src = String(text).replace(/```[a-z]*\n?([\s\S]*?)```/g, (_, code) => {
    blocks.push(`<pre class="md-code"><code>${escapeHtml(code.replace(/\n$/, ''))}</code></pre>`);
    return `\u0000${blocks.length - 1}\u0000`;
  });
  src = escapeHtml(src)
    .replace(/`([^`]+)`/g, '<code>$1</code>')
    .replace(/\*\*([^*]+)\*\*/g, '<strong>$1</strong>');
  const html = src.split(/\n{2,}/).map(para => {
    const lines = para.split('\n');
    if (lines.every(l => /^\s*([-*]|\d+\.)\s+/.test(l) || !l.trim())) {
      const ordered = /^\s*\d+\./.test(lines[0]);
      const items = lines.filter(l => l.trim()).map(l => `<li>${l.replace(/^\s*([-*]|\d+\.)\s+/, '')}</li>`).join('');
      return ordered ? `<ol>${items}</ol>` : `<ul>${items}</ul>`;
    }
    if (/^\u0000\d+\u0000$/.test(para.trim())) return para.trim();
    return `<p>${lines.join('<br>')}</p>`;
  }).join('');
  return html.replace(/\u0000(\d+)\u0000/g, (_, i) => blocks[Number(i)]);
}

function loadJSON(key, fallback) {
  try {
    const raw = localStorage.getItem(key);
    return raw ? JSON.parse(raw) : fallback;
  } catch (e) {
    return fallback;
  }
}

function saveJSON(key, value) {
  try {
    localStorage.setItem(key, JSON.stringify(value));
  } catch (e) {
    console.error('Failed to save', key, e);
  }
}

function normalizeAnswer(s) {
  return String(s || '').replace(/\s+/g, '').replace(/["']/g, "'").toLowerCase();
}

const DAY_MS = 24 * 60 * 60 * 1000;

// ---------------------------------------------------------------------------
// Navigation between chapters and practice
// ---------------------------------------------------------------------------

async function ensureProblemsLoaded() {
  if (practiceState.problems.length === 0) {
    await loadProblems();
    await loadCategories();
  }
}

// enterPractice switches the main view to practice mode without toggling it off.
async function enterPractice() {
  const toggle = el('practice-toggle');
  if (toggle && !toggle.classList.contains('active')) {
    await toggle.onclick();
  }
}

async function goToProblem(problemId, opts = {}) {
  await enterPractice();
  await openProblem(problemId, opts);
}

async function goToCategory(categoryId) {
  await enterPractice();
  showPracticeList();
  const section = document.querySelector(`.category-section[data-category="${categoryId}"]`);
  if (section) section.scrollIntoView({ behavior: 'smooth', block: 'start' });
}

function goToChapter(num) {
  const ch = state.chapters.find(c => c.number === num);
  if (!ch) return;
  const toggle = el('practice-toggle');
  if (toggle && toggle.classList.contains('active')) {
    toggle.classList.remove('active');
    el('practice-list').classList.add('hidden');
    el('practice-mode').classList.add('hidden');
    el('chapters').classList.remove('hidden');
  }
  el('detail').classList.remove('hidden');
  el('detail-empty').classList.add('hidden');
  selectChapter(ch);
  const content = document.querySelector('.chapter-content');
  if (content) content.scrollTop = 0;
}

function chapterForCategory(categoryId) {
  return (state.chapters || []).find(c => c.category === categoryId);
}

// ---------------------------------------------------------------------------
// Chapter progress (retrieval questions answered, completions done)
// ---------------------------------------------------------------------------

const CHAPTER_PROGRESS_KEY = 'dsatutor_chapter_progress';

function chapterProgress(num) {
  const all = loadJSON(CHAPTER_PROGRESS_KEY, {});
  return all[num] || { checks: {}, completions: {} };
}

function updateChapterProgress(num, fn) {
  const all = loadJSON(CHAPTER_PROGRESS_KEY, {});
  const cur = all[num] || { checks: {}, completions: {} };
  fn(cur);
  all[num] = cur;
  saveJSON(CHAPTER_PROGRESS_KEY, all);
  markChapterDone();
}

function isChapterDone(ch) {
  const qs = ch.checkQuestions || [];
  if (!qs.length) return false;
  const prog = chapterProgress(ch.number);
  return qs.every((_, i) => prog.checks[i] === true);
}

function markChapterDone() {
  document.querySelectorAll('#chapters .chapter-item').forEach((item, i) => {
    const ch = state.chapters[i];
    if (ch) item.classList.toggle('done', isChapterDone(ch));
  });
}

// ---------------------------------------------------------------------------
// Chapter lesson sections
// ---------------------------------------------------------------------------

function renderLesson(ch) {
  renderLessonPath(ch);
  renderWhenToUse(ch);
  renderWorkedExamples(ch);
  renderCompletions(ch);
  renderCheckQuestions(ch);
  renderPracticeLadder(ch);
  const viz = document.querySelector('.visualization-section');
  const memory = document.querySelector('.memory-steps-section');
  const hasAnimations = (ch.animations || []).length > 0;
  if (viz) viz.classList.toggle('hidden', !hasAnimations);
  if (memory) memory.classList.toggle('hidden', !hasAnimations);
  markChapterDone();
}

function renderLessonPath(ch) {
  const motivation = el('lesson-motivation');
  if (motivation) {
    motivation.innerHTML = ch.motivation
      ? `<h4>Why you need this</h4>${renderMarkdown(ch.motivation)}`
      : '';
  }
  const path = el('lesson-path');
  if (!path) return;
  const prereqs = (ch.prerequisites || [])
    .map(n => state.chapters.find(c => c.number === n))
    .filter(Boolean);
  let html = '';
  if (prereqs.length) {
    html += `<div class="lesson-prereqs"><span>Before this:</span> ${prereqs.map(p =>
      `<a href="#" class="chapter-link${isChapterDone(p) ? ' done' : ''}" data-chapter="${p.number}">${pad(p.number)}. ${escapeHtml(p.title)}</a>`).join(' ')}</div>`;
  }
  html += `<ol class="lesson-steps">
    <li>Read the ideas and why they work</li>
    <li>Study the worked example step by step</li>
    <li>Fill in the completion problem</li>
    <li>Answer the check questions from memory</li>
    <li>Climb the practice ladder${ch.category ? '' : ' in later chapters'}</li>
  </ol>`;
  path.innerHTML = html;
  path.querySelectorAll('.chapter-link').forEach(a => {
    a.onclick = (e) => {
      e.preventDefault();
      goToChapter(Number(a.dataset.chapter));
    };
  });
}

function renderWhenToUse(ch) {
  const target = el('lesson-when');
  const section = el('lesson-when-section');
  if (!target || !section) return;
  const cues = ch.whenToUse || [];
  section.classList.toggle('hidden', cues.length === 0);
  target.innerHTML = cues.map(c => `<li>${renderInline(c)}</li>`).join('');
}

function renderInline(text) {
  return escapeHtml(text).replace(/`([^`]+)`/g, '<code>$1</code>');
}

// Worked examples are revealed one subgoal at a time; before each reveal the
// learner is asked to predict the next step.
function renderWorkedExamples(ch) {
  const target = el('lesson-worked');
  const section = el('lesson-worked-section');
  if (!target || !section) return;
  const examples = ch.workedExamples || [];
  section.classList.toggle('hidden', examples.length === 0);
  target.innerHTML = '';
  examples.forEach((ex, exIdx) => {
    const block = document.createElement('div');
    block.className = 'worked-example';
    block.innerHTML = `
      <h3>${escapeHtml(ex.title)}</h3>
      <div class="worked-problem">${renderMarkdown(ex.problem)}</div>
      <div class="worked-steps"></div>
      <div class="worked-controls">
        <span class="worked-predict">Before revealing: what should happen next, and why?</span>
        <button class="btn-small primary worked-next">Reveal step 1</button>
        <button class="btn-small ghost worked-all">Show all</button>
      </div>
      <details class="worked-code hidden"><summary>Full solution</summary><pre><code>${escapeHtml(ex.code || '')}</code></pre></details>
    `;
    const stepsEl = block.querySelector('.worked-steps');
    const nextBtn = block.querySelector('.worked-next');
    const allBtn = block.querySelector('.worked-all');
    const predict = block.querySelector('.worked-predict');
    const codeEl = block.querySelector('.worked-code');
    let shown = 0;
    const reveal = () => {
      const step = ex.steps[shown];
      if (!step) return;
      const node = document.createElement('div');
      node.className = 'worked-step';
      node.innerHTML = `
        <div class="worked-step-head"><span class="subgoal">${escapeHtml(step.subgoal)}</span><span class="worked-step-num">Step ${shown + 1} of ${ex.steps.length}</span></div>
        <div class="worked-step-text">${renderMarkdown(step.explanation)}</div>
        ${step.code ? `<pre class="worked-step-code"><code>${escapeHtml(step.code)}</code></pre>` : ''}
        ${step.state ? `<div class="worked-step-state"><span>Trace</span><pre>${escapeHtml(step.state)}</pre></div>` : ''}
      `;
      stepsEl.appendChild(node);
      shown++;
      if (shown >= ex.steps.length) {
        nextBtn.classList.add('hidden');
        allBtn.classList.add('hidden');
        predict.classList.add('hidden');
        codeEl.classList.remove('hidden');
      } else {
        nextBtn.textContent = `Reveal step ${shown + 1}`;
      }
    };
    nextBtn.onclick = reveal;
    allBtn.onclick = () => { while (shown < ex.steps.length) reveal(); };
    block.dataset.index = exIdx;
    target.appendChild(block);
  });
}

// Completion problems: the worked solution with key lines blanked out.
function renderCompletions(ch) {
  const target = el('lesson-completion');
  const section = el('lesson-completion-section');
  if (!target || !section) return;
  const items = ch.completions || [];
  section.classList.toggle('hidden', items.length === 0);
  target.innerHTML = '';
  const prog = chapterProgress(ch.number);
  items.forEach((item, idx) => {
    const block = document.createElement('div');
    block.className = 'completion';
    const parts = (item.code || '').split('____');
    let codeHtml = '';
    parts.forEach((part, i) => {
      codeHtml += escapeHtml(part);
      if (i < parts.length - 1) {
        codeHtml += `<input class="blank" data-i="${i}" spellcheck="false" autocomplete="off" size="${Math.max(6, (item.blanks[i] || '').length + 2)}">`;
      }
    });
    block.innerHTML = `
      <h3>${escapeHtml(item.title)}${prog.completions[idx] ? ' <span class="done-badge">done</span>' : ''}</h3>
      <div class="completion-prompt">${renderMarkdown(item.prompt)}</div>
      <pre class="completion-code"><code>${codeHtml}</code></pre>
      <div class="completion-actions">
        <button class="btn-small primary completion-check">Check</button>
        <button class="btn-small ghost completion-reveal">Show answers</button>
        <span class="completion-result"></span>
      </div>
      <div class="completion-explanation hidden">${renderMarkdown(item.explanation)}</div>
    `;
    const inputs = [...block.querySelectorAll('.blank')];
    const result = block.querySelector('.completion-result');
    const explanation = block.querySelector('.completion-explanation');
    block.querySelector('.completion-check').onclick = () => {
      let right = 0;
      inputs.forEach((input, i) => {
        const ok = normalizeAnswer(input.value) === normalizeAnswer(item.blanks[i]);
        input.classList.toggle('right', ok);
        input.classList.toggle('wrong', !ok);
        if (ok) right++;
      });
      result.textContent = `${right} of ${inputs.length} correct`;
      if (right === inputs.length) {
        explanation.classList.remove('hidden');
        updateChapterProgress(ch.number, p => { p.completions[idx] = true; });
      }
    };
    block.querySelector('.completion-reveal').onclick = () => {
      inputs.forEach((input, i) => {
        if (normalizeAnswer(input.value) !== normalizeAnswer(item.blanks[i])) {
          input.value = item.blanks[i];
          input.classList.remove('right');
          input.classList.add('revealed');
        }
      });
      explanation.classList.remove('hidden');
      result.textContent = 'Compare each answer with what you wrote, then try the next one without peeking.';
    };
    target.appendChild(block);
  });
}

// Retrieval practice: multiple choice gives feedback on click; free recall asks
// for an attempt before showing the model answer.
function renderCheckQuestions(ch) {
  const target = el('lesson-check');
  const section = el('lesson-check-section');
  if (!target || !section) return;
  const qs = ch.checkQuestions || [];
  section.classList.toggle('hidden', qs.length === 0);
  target.innerHTML = '';
  const prog = chapterProgress(ch.number);
  qs.forEach((q, idx) => {
    const block = document.createElement('div');
    block.className = 'check-question' + (prog.checks[idx] === true ? ' answered-right' : '');
    if (q.choices && q.choices.length) {
      block.innerHTML = `
        <div class="check-prompt"><span class="check-num">${idx + 1}</span>${renderMarkdown(q.prompt)}</div>
        <div class="check-choices">${q.choices.map((c, i) => `<button class="check-choice" data-i="${i}">${renderInline(c)}</button>`).join('')}</div>
        <div class="check-explanation hidden"></div>
      `;
      const explanation = block.querySelector('.check-explanation');
      block.querySelectorAll('.check-choice').forEach(btn => {
        btn.onclick = () => {
          const i = Number(btn.dataset.i);
          const correct = i === q.answer;
          block.querySelectorAll('.check-choice').forEach(b => {
            b.classList.toggle('right', Number(b.dataset.i) === q.answer && correct);
          });
          btn.classList.add(correct ? 'right' : 'wrong');
          explanation.innerHTML = `<strong>${correct ? 'Correct.' : 'Not quite.'}</strong> ${correct ? renderMarkdown(q.explanation) : 'Think about why that choice fails, then try again.'}`;
          explanation.classList.remove('hidden');
          if (correct) {
            block.classList.add('answered-right');
            updateChapterProgress(ch.number, p => { p.checks[idx] = true; });
          }
        };
      });
    } else {
      block.innerHTML = `
        <div class="check-prompt"><span class="check-num">${idx + 1}</span>${renderMarkdown(q.prompt)}</div>
        <textarea class="check-recall" rows="3" placeholder="Answer from memory first, in your own words..."></textarea>
        <div class="check-actions">
          <button class="btn-small check-show">Compare with model answer</button>
        </div>
        <div class="check-explanation hidden">
          ${renderMarkdown(q.explanation)}
          <div class="check-self">Did your answer contain the key idea?
            <button class="btn-small check-self-yes">Yes</button>
            <button class="btn-small ghost check-self-no">Not yet</button>
          </div>
        </div>
      `;
      const textarea = block.querySelector('.check-recall');
      const explanation = block.querySelector('.check-explanation');
      block.querySelector('.check-show').onclick = () => {
        if (!textarea.value.trim() && !confirm('Retrieving an answer, even a wrong one, is what makes it stick. Show the model answer anyway?')) {
          return;
        }
        explanation.classList.remove('hidden');
      };
      block.querySelector('.check-self-yes').onclick = () => {
        block.classList.add('answered-right');
        updateChapterProgress(ch.number, p => { p.checks[idx] = true; });
      };
      block.querySelector('.check-self-no').onclick = () => {
        block.classList.remove('answered-right');
        updateChapterProgress(ch.number, p => { p.checks[idx] = false; });
        textarea.value = '';
        explanation.classList.add('hidden');
        textarea.placeholder = 'Reread the concept above, then try again without looking.';
      };
    }
    target.appendChild(block);
  });
}

async function renderPracticeLadder(ch) {
  const target = el('lesson-practice');
  const section = el('lesson-practice-section');
  if (!target || !section) return;
  if (!ch.category) {
    section.classList.add('hidden');
    return;
  }
  section.classList.remove('hidden');
  target.innerHTML = '<p class="placeholder">Loading problems...</p>';
  await ensureProblemsLoaded();
  if (state.current !== ch) return; // the learner moved on while loading
  const problems = practiceState.problems.filter(p => p.category === ch.category);
  const progress = practiceState.progress;
  target.innerHTML = `
    <p class="ladder-intro">Solve these in order: each adds one idea to the last. Try each one for real before opening a hint.</p>
    <ol class="ladder">${problems.map(p => {
      const prog = progress[p.id] || {};
      const status = prog.solved ? 'solved' : (prog.attempts > 0 ? 'attempted' : '');
      return `<li class="ladder-item ${status}" data-id="${p.id}">
        <span class="ladder-title">${escapeHtml(p.title)}</span>
        <span class="difficulty-badge difficulty-${p.difficulty.toLowerCase()}">${p.difficulty}</span>
        ${prog.solved ? '<span class="solved-badge">Solved</span>' : ''}
      </li>`;
    }).join('')}</ol>`;
  target.querySelectorAll('.ladder-item').forEach(item => {
    item.onclick = () => goToProblem(item.dataset.id);
  });
}

// ---------------------------------------------------------------------------
// Spaced review
// ---------------------------------------------------------------------------

// scheduleReview sets the next review date from the learner's rating, a light
// version of SM-2: "again" resets, the others grow the interval.
// base is the interval before this solve, so re-rating doesn't compound.
function scheduleReview(problemId, rating, base) {
  const prog = practiceState.progress[problemId] || {};
  const prev = base !== undefined ? base : (prog.review?.interval || 0);
  let interval;
  switch (rating) {
    case 'again': interval = 1; break;
    case 'hard': interval = Math.max(1, Math.round(prev * 1.2) || 1); break;
    case 'easy': interval = prev ? Math.round(prev * 3.5) : 7; break;
    default: interval = prev ? Math.round(prev * 2.5) : 3; // good
  }
  interval = Math.min(interval, 180);
  updateProblemProgress(problemId, {
    review: { interval, due: Date.now() + interval * DAY_MS, reps: (prog.review?.reps || 0) + (base !== undefined ? 0 : 1), lastRating: rating },
  });
  return interval;
}

function dueReviews() {
  const now = Date.now();
  return practiceState.problems.filter(p => {
    const r = practiceState.progress[p.id]?.review;
    return r && r.due <= now;
  });
}

function renderReviewBar() {
  const bar = el('practice-learning-bar');
  if (!bar) return;
  const due = dueReviews();
  const unlocked = unlockedCategories();
  bar.innerHTML = `
    <div class="learning-bar-item">
      <button id="start-review" class="btn-small ${due.length ? 'primary' : 'ghost'}" ${due.length ? '' : 'disabled'}>Review${due.length ? ` (${due.length} due)` : ''}</button>
      <span class="learning-bar-hint">${due.length ? 'Solve these again from scratch. Spacing out practice makes it last.' : 'Solved problems come back here on a schedule.'}</span>
    </div>
    <div class="learning-bar-item">
      <button id="start-mixed" class="btn-small ${unlocked.length >= 2 ? 'primary' : 'ghost'}" ${unlocked.length ? '' : 'disabled'}>Mixed Practice</button>
      <span class="learning-bar-hint">${unlocked.length >= 2
        ? `Random problems from ${unlocked.length} topics you've started, so you practise recognising the pattern.`
        : 'Unlocks once you solve a problem; best after two or more topics.'}</span>
    </div>
  `;
  const reviewBtn = el('start-review');
  if (reviewBtn) reviewBtn.onclick = () => {
    const next = due.sort((a, b) => practiceState.progress[a.id].review.due - practiceState.progress[b.id].review.due)[0];
    if (next) openProblem(next.id, { review: true });
  };
  const mixedBtn = el('start-mixed');
  if (mixedBtn) mixedBtn.onclick = () => startMixed();
}

// ---------------------------------------------------------------------------
// Interleaved mixed practice
// ---------------------------------------------------------------------------

function unlockedCategories() {
  const ids = new Set();
  practiceState.problems.forEach(p => {
    if (practiceState.progress[p.id]?.solved) ids.add(p.category);
  });
  return practiceState.categories.filter(c => ids.has(c.id));
}

function pickMixedProblem(excludeId) {
  const unlocked = new Set(unlockedCategories().map(c => c.id));
  const due = dueReviews().filter(p => p.id !== excludeId);
  if (due.length) return due[Math.floor(Math.random() * due.length)];
  // Otherwise, the next unsolved rung of each unlocked ladder.
  const candidates = [];
  unlocked.forEach(cat => {
    const next = practiceState.problems.find(p => p.category === cat && !practiceState.progress[p.id]?.solved && p.id !== excludeId);
    if (next) candidates.push(next);
  });
  if (!candidates.length) {
    // Everything unlocked is solved: revisit a random solved one.
    const solved = practiceState.problems.filter(p => unlocked.has(p.category) && p.id !== excludeId);
    return solved[Math.floor(Math.random() * solved.length)];
  }
  return candidates[Math.floor(Math.random() * candidates.length)];
}

async function startMixed(excludeId) {
  await ensureProblemsLoaded();
  const next = pickMixedProblem(excludeId);
  if (!next) {
    alert('Solve at least one problem in a topic to add it to Mixed Practice.');
    return;
  }
  openProblem(next.id, { mixed: true });
}

// Before a mixed problem is shown, the learner names the pattern; this trains
// recognising which technique fits, which blocked practice never asks for.
function renderPatternPrompt(problem) {
  const target = el('pattern-prompt');
  if (!target) return;
  if (!practiceState.mode?.mixed) {
    target.classList.add('hidden');
    target.innerHTML = '';
    document.querySelector('#problem-description .chapter-links')?.classList.remove('hidden');
    return;
  }
  const correct = practiceState.categories.find(c => c.id === problem.category);
  const pool = unlockedCategories().filter(c => c.id !== problem.category);
  const others = practiceState.categories.filter(c => c.id !== problem.category && !pool.includes(c));
  while (pool.length < 3 && others.length) {
    pool.push(others.splice(Math.floor(Math.random() * others.length), 1)[0]);
  }
  const choices = [correct, ...pool.sort(() => Math.random() - 0.5).slice(0, 4)].filter(Boolean).sort(() => Math.random() - 0.5);
  document.querySelector('#problem-description .chapter-links')?.classList.add('hidden');
  target.classList.remove('hidden');
  target.innerHTML = `
    <div class="pattern-question">Mixed practice: read the problem, then decide which pattern you'd use before writing code.</div>
    <div class="pattern-choices">${choices.map(c => `<button class="check-choice" data-id="${c.id}">${escapeHtml(c.name)}</button>`).join('')}</div>
    <div class="pattern-feedback hidden"></div>
  `;
  target.querySelectorAll('.check-choice').forEach(btn => {
    btn.onclick = () => {
      const right = btn.dataset.id === problem.category;
      btn.classList.add(right ? 'right' : 'wrong');
      target.querySelectorAll('.check-choice').forEach(b => {
        if (b.dataset.id === problem.category) b.classList.add('right');
        b.disabled = true;
      });
      const stats = loadJSON('dsatutor_pattern_stats', { right: 0, total: 0 });
      stats.total++;
      if (right) stats.right++;
      saveJSON('dsatutor_pattern_stats', stats);
      const fb = target.querySelector('.pattern-feedback');
      const ch = chapterForCategory(problem.category);
      fb.innerHTML = `${right ? 'Right' : `It's <strong>${escapeHtml(correct?.name || problem.category)}</strong>`}: ${escapeHtml(correct?.description || '')}.
        ${ch ? `<a href="#" class="chapter-link" data-chapter="${ch.number}">Chapter ${ch.number}</a>` : ''}
        <span class="pattern-stats">Pattern accuracy so far: ${stats.right}/${stats.total}</span>`;
      fb.classList.remove('hidden');
      fb.querySelector('.chapter-link')?.addEventListener('click', e => { e.preventDefault(); goToChapter(ch.number); });
      document.querySelector('#problem-description .chapter-links')?.classList.remove('hidden');
    };
  });
}

// ---------------------------------------------------------------------------
// After an accepted submission: self-explanation and a review rating
// ---------------------------------------------------------------------------

function renderAfterSolve(problem) {
  const target = el('after-solve');
  if (!target) return;
  const prior = practiceState.progress[problem.id]?.review;
  const base = prior && !practiceState.mode?.reviewScheduled ? prior.interval : (practiceState.mode?.reviewBase || 0);
  const checks = problem.selfCheck || [];
  target.classList.remove('hidden');
  target.innerHTML = `
    <div class="after-solve-head">Accepted. Before moving on:</div>
    ${checks.length ? `<div class="after-solve-section">
      <h4>Explain it to yourself</h4>
      ${checks.map((q, i) => `
        <div class="self-check">
          <div class="self-check-q">${renderMarkdown(q.prompt)}</div>
          <textarea rows="2" placeholder="Your explanation..."></textarea>
          <button class="btn-small ghost self-check-show" data-i="${i}">Compare</button>
          <div class="self-check-a hidden">${renderMarkdown(q.answer)}</div>
        </div>`).join('')}
    </div>` : ''}
    <div class="after-solve-section">
      <h4>How hard was that? (sets when it comes back for review)</h4>
      <div class="rating-buttons">
        <button class="btn-small" data-rating="again">Needed the solution</button>
        <button class="btn-small" data-rating="hard">Hard</button>
        <button class="btn-small primary" data-rating="good">Good</button>
        <button class="btn-small" data-rating="easy">Easy</button>
      </div>
      <div class="rating-result"></div>
    </div>
    ${practiceState.mode?.mixed ? '<button id="mixed-next" class="btn-small primary">Next mixed problem</button>' : ''}
  `;
  target.querySelectorAll('.self-check-show').forEach(btn => {
    btn.onclick = () => btn.nextElementSibling.classList.remove('hidden');
  });
  target.querySelectorAll('.rating-buttons button').forEach(btn => {
    btn.onclick = () => {
      const days = scheduleReview(problem.id, btn.dataset.rating, base);
      target.querySelector('.rating-result').textContent = `It'll come back for review in ${days} day${days === 1 ? '' : 's'}.`;
      target.querySelectorAll('.rating-buttons button').forEach(b => b.classList.toggle('primary', b === btn));
    };
  });
  const next = el('mixed-next');
  if (next) next.onclick = () => startMixed(problem.id);
  // Default schedule in case the learner skips the rating, once per visit.
  if (!practiceState.mode) practiceState.mode = {};
  if (!practiceState.mode.reviewScheduled) {
    const viewed = practiceState.progress[problem.id]?.solutionViewed;
    practiceState.mode.reviewBase = base;
    practiceState.mode.reviewScheduled = true;
    scheduleReview(problem.id, viewed ? 'again' : 'good', base);
    updateProblemProgress(problem.id, { review: { ...practiceState.progress[problem.id].review, reps: (prior?.reps || 0) + 1 } });
  }
}
