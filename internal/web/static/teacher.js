// Teacher Mode: on a chapter page, the AI tutor stops waiting for questions
// and teaches the chapter's topic one idea at a time, from a lesson it plans
// and writes itself. The chapter's prewritten page is put away: the lesson
// page starts blank and the teacher writes each section onto it as it
// teaches (write_lesson), pointing at what it wrote as reference. The
// Teacher's Scratchpad holds only the examples (sketches, traced code). Each
// turn ends with a question, often with clickable answers, and the lesson
// adapts to the answers. Normal chapter view and the regular tutor keep
// using the prewritten pages.
//
// The plan comes from the server (internal/tutor/lesson.go) when a lesson
// starts. Each chapter keeps its own plan, written sections and lesson
// conversation, so leaving and coming back resumes where it stopped. Loaded
// after tutor.js and scratchpad.js.

(function () {
  'use strict';

  const ON_KEY = 'dsatutor_teacher_mode';
  const LESSON_KEY = (n) => `dsatutor_teacher_lesson:${n}`;
  const teach = { on: false };
  const $ = (id) => document.getElementById(id);

  function chapterOnScreen() {
    const detail = $('detail');
    if (!detail || detail.classList.contains('hidden')) return null;
    return (typeof state !== 'undefined' && state.current) || null;
  }

  // ---------------------------------------------------------------------
  // The lesson for one chapter: { plan: {title, sections}, current,
  // sections: {index: {heading, body, code}} }.

  function load(n) {
    try { return JSON.parse(localStorage.getItem(LESSON_KEY(n)) || 'null') || {}; } catch (e) { return {}; }
  }
  function save(n, lesson) {
    try { localStorage.setItem(LESSON_KEY(n), JSON.stringify(lesson)); } catch (e) { /* private mode */ }
  }

  // What the server sees of the lesson: the plan and where it stands.
  function lessonContext(n) {
    const l = load(n);
    if (!l.plan || !(l.plan.sections || []).length) return undefined;
    return {
      title: l.plan.title,
      sections: l.plan.sections,
      current: l.current || 0,
      written: Object.keys(l.sections || {}).map(Number),
      // The text already on the page, so the teacher builds on it.
      pages: Object.entries(l.sections || {}).map(([i, p]) => ({ section: +i, heading: p.heading, body: p.body })),
    };
  }

  function setLesson(plan) {
    const ch = chapterOnScreen();
    if (!ch || !plan) return;
    save(ch.number, { plan: { title: plan.title, sections: plan.sections || [] }, current: 0, sections: {} });
    render();
  }

  function writeSection(a) {
    const ch = chapterOnScreen();
    if (!ch) return false;
    const l = load(ch.number);
    l.sections = l.sections || {};
    l.sections[a.section] = { heading: a.heading, body: a.body, code: a.code || '' };
    l.current = a.section;
    save(ch.number, l);
    render();
    const el = $(`tl-section-${a.section}`);
    if (el) {
      el.classList.remove('tl-fresh');
      void el.offsetWidth; // restart the "just written" animation
      el.classList.add('tl-fresh');
      el.scrollIntoView({ block: 'start', behavior: 'smooth' });
    }
    return true;
  }

  // ---------------------------------------------------------------------
  // The lesson page, shown in place of the chapter's content.

  function page() {
    let el = $('teacher-lesson');
    if (el) return el;
    const host = document.querySelector('#detail .chapter-content');
    if (!host) return null;
    el = document.createElement('article');
    el.id = 'teacher-lesson';
    el.className = 'teacher-lesson';
    el.setAttribute('aria-live', 'polite');
    host.prepend(el);
    return el;
  }

  function md(text) {
    return window.dsaTutor?.renderMarkdown ? window.dsaTutor.renderMarkdown(text || '') : `<p>${escapeHtml(text || '')}</p>`;
  }

  function render() {
    const host = document.querySelector('#detail .chapter-content');
    const ch = chapterOnScreen();
    const showing = teach.on && !!ch;
    if (host) host.classList.toggle('teaching', showing);
    if (!showing) return;
    const el = page();
    if (!el) return;
    const l = load(ch.number);
    const plan = l.plan;
    const written = l.sections || {};
    if (!plan || !(plan.sections || []).length) {
      el.innerHTML = `<header class="tl-head"><span class="tl-kicker">Teacher mode · Ch. ${ch.number}</span><h2>${escapeHtml(ch.title)}</h2></header>
        <p class="tl-empty">Your teacher is planning this lesson. It will be written here as you go.</p>`;
      return;
    }
    const current = l.current || 0;
    const outline = plan.sections.map((s, i) => {
      const cls = [written[i] ? 'done' : '', i === current ? 'current' : ''].filter(Boolean).join(' ');
      return `<li class="${cls}">${written[i] ? `<a href="#tl-section-${i}" data-tl-jump="${i}">${escapeHtml(s.heading)}</a>` : escapeHtml(s.heading)}</li>`;
    }).join('');
    const sections = plan.sections.map((s, i) => {
      const w = written[i];
      if (!w) return '';
      return `<section class="tl-section" id="tl-section-${i}" data-section="${i}">
          <span class="tl-num">Part ${i + 1} of ${plan.sections.length}</span>
          <h3>${escapeHtml(w.heading || s.heading)}</h3>
          <div class="tl-body">${md(w.body)}</div>
          ${w.code ? `<pre class="tl-code"><code>${escapeHtml(w.code)}</code></pre>` : ''}
        </section>`;
    }).join('');
    const remaining = plan.sections.length - Object.keys(written).length;
    el.innerHTML = `<header class="tl-head"><span class="tl-kicker">Teacher mode · Ch. ${ch.number} · written live</span><h2>${escapeHtml(plan.title || ch.title)}</h2></header>
      <nav class="tl-outline" aria-label="Lesson plan"><ol>${outline}</ol></nav>
      ${sections || '<p class="tl-empty">Your teacher will write the first part here in a moment.</p>'}
      ${sections && remaining > 0 ? `<p class="tl-more">${remaining} more part${remaining === 1 ? '' : 's'} to come.</p>` : ''}`;
    el.querySelectorAll('[data-tl-jump]').forEach(a => {
      a.onclick = (e) => { e.preventDefault(); $(`tl-section-${a.dataset.tlJump}`)?.scrollIntoView({ block: 'start', behavior: 'smooth' }); };
    });
  }

  // The parts of the lesson page the teacher may point at.
  function targets() {
    const out = [];
    document.querySelectorAll('#teacher-lesson .tl-section').forEach(sec => {
      const i = sec.dataset.section;
      const heading = sec.querySelector('h3')?.textContent.trim() || '';
      out.push({ id: `lesson-${i}`, label: `Lesson page, section ${i}: ${heading}`, get: () => $(`tl-section-${i}`) });
      if (sec.querySelector('.tl-code')) {
        out.push({ id: `lesson-${i}-code`, label: `Code in lesson page section ${i}: ${heading}`, get: () => $(`tl-section-${i}`)?.querySelector('.tl-code'), lines: true });
      }
    });
    return out;
  }

  // ---------------------------------------------------------------------
  // The tutor mode

  window.tutorModes = window.tutorModes || [];
  window.tutorModes.push({
    context() {
      if (!teach.on) return null;
      const ch = chapterOnScreen();
      if (!ch) return null;
      return {
        key: `teach:${ch.number}`,
        label: `Ch. ${ch.number}: ${ch.title}`,
        kind: 'teacher',
        contextLine: `Teacher mode · Ch. ${ch.number}: ${ch.title}`,
        placeholder: 'Answer, ask a question, or say "next"…',
        emptyHtml: '<p>Your teacher is getting the lesson ready…</p>',
        // Sent (hidden) whenever this chapter's lesson has nothing in it yet:
        // on turning Teacher Mode on, opening the panel, Clear, reloads and
        // chapter changes alike. The server plans a fresh lesson for it.
        kickoff: '(Start the lesson on this topic from the beginning. Write the first part on the lesson page, keep it to one idea, then ask me a question.)',
        kickoffLabel: 'Start the lesson',
        quick: [
          ['Next', "Got it. Let's move on to the next idea."],
          ['Show me', 'Show me that in the scratchpad with a small example.'],
          ["I'm confused", "I'm confused. Explain it another way, with a smaller example."],
          ['Quiz me', 'Quiz me on this before we move on.'],
          ['Let me try', 'Give me a small exercise to try in the scratchpad.'],
        ],
        ctx: { view: 'chapter', chapterNumber: ch.number, teach: true, lesson: lessonContext(ch.number), learner: window.dsaLearning?.chapterProfile(ch) },
      };
    },
  });

  // Open the lesson for the chapter on screen. A chapter with nothing yet
  // starts its lesson (the tutor sends the kickoff); a chapter taught before
  // resumes where it stopped.
  async function begin(force) {
    const ch = chapterOnScreen();
    if (!ch) return;
    render();
    await window.dsaTutor.open();
    if (force) window.dsaTutor.kickoff();
    else window.dsaTutor.sync();
  }

  function setOn(on) {
    teach.on = on;
    try { localStorage.setItem(ON_KEY, on ? '1' : ''); } catch (e) { /* private mode */ }
    const btn = $('teacher-mode');
    if (btn) {
      btn.classList.toggle('active', on);
      btn.setAttribute('aria-pressed', on ? 'true' : 'false');
      btn.textContent = on ? 'Teacher mode: on' : 'Teacher mode';
    }
    render();
    if (on) begin(true);
    else { window.dsaTutor.clear(); window.dsaTutor.sync(); }
  }

  // Moving to another chapter with Teacher Mode on starts or resumes that
  // chapter's lesson, however the chapter was opened (sidebar, links, the
  // page restoring the last chapter): watch the chapter title.
  let shownChapter = null;
  let pending = 0;
  const onChapterChange = () => {
    clearTimeout(pending);
    pending = setTimeout(() => {
      const ch = chapterOnScreen();
      const n = ch ? ch.number : null;
      if (n === shownChapter) return;
      shownChapter = n;
      render();
      if (teach.on && ch) begin(false);
    }, 300);
  };
  const title = $('detail-title');
  if (title) new MutationObserver(onChapterChange).observe(title, { childList: true, characterData: true, subtree: true });
  const detail = $('detail');
  if (detail) new MutationObserver(onChapterChange).observe(detail, { attributes: true, attributeFilter: ['class'] });

  async function init() {
    const st = await window.dsaTutor?.status();
    if (!st || !st.available) return;
    const group = $('toggle-edit')?.parentElement;
    if (!group || $('teacher-mode')) return;
    const btn = document.createElement('button');
    btn.id = 'teacher-mode';
    btn.className = 'btn-small teacher-btn';
    btn.textContent = 'Teacher mode';
    btn.title = 'Have the AI tutor teach this topic live: it writes the lesson as it goes, with examples in a scratchpad';
    btn.setAttribute('aria-pressed', 'false');
    btn.onclick = () => setOn(!teach.on);
    group.prepend(btn);
    let saved = false;
    try { saved = localStorage.getItem(ON_KEY) === '1'; } catch (e) { /* ignore */ }
    if (saved) {
      teach.on = true;
      btn.classList.add('active');
      btn.setAttribute('aria-pressed', 'true');
      btn.textContent = 'Teacher mode: on';
      render();
    }
  }

  window.dsaTeacher = {
    start: () => setOn(true),
    stop: () => setOn(false),
    isOn: () => teach.on,
    setLesson,
    writeSection,
    targets,
  };
  init();
})();
