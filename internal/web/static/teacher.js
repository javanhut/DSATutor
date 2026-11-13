// Teacher Mode: on a chapter page, the AI tutor stops waiting for questions
// and teaches the chapter one idea at a time. Each turn it points at the
// section it's on, shows the idea in the Teacher's Scratchpad (a diagram or
// a traced program), and asks the learner something, often with clickable
// answers. It adapts to the answers: right moves on, wrong gets another
// angle, "I'm confused" gets a smaller example. The persona and the full
// chapter content it teaches from are in internal/tutor/prompt.go.
//
// Each chapter keeps its own lesson conversation, so leaving and coming
// back resumes where it stopped. Loaded after tutor.js and scratchpad.js.

(function () {
  'use strict';

  const ON_KEY = 'dsatutor_teacher_mode';
  const teach = { on: false };
  const $ = (id) => document.getElementById(id);

  function chapterOnScreen() {
    const detail = $('detail');
    if (!detail || detail.classList.contains('hidden')) return null;
    return (typeof state !== 'undefined' && state.current) || null;
  }

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
        // chapter changes alike.
        kickoff: '(Start the lesson on this chapter from the beginning. Keep it to one idea, then ask me a question.)',
        kickoffLabel: 'Start the lesson',
        quick: [
          ['Next', "Got it. Let's move on to the next idea."],
          ['Show me', 'Show me that in the scratchpad with a small example.'],
          ["I'm confused", "I'm confused. Explain it another way, with a smaller example."],
          ['Quiz me', 'Quiz me on this before we move on.'],
          ['Let me try', 'Give me a small exercise to try in the scratchpad.'],
        ],
        ctx: { view: 'chapter', chapterNumber: ch.number, teach: true, learner: window.dsaLearning?.chapterProfile(ch) },
      };
    },
  });

  // Some chapters open on a full-page view (Chapter 1's Runtime Shapes)
  // that hides the lesson sections. The teacher teaches from — and points
  // at — the lesson, so switch to a topic that shows it.
  function showLessonContent() {
    const standard = $('standard-chapter-content');
    const select = $('storyboard-select');
    if (!standard || !standard.classList.contains('hidden') || !select) return;
    for (const opt of select.options) {
      select.value = opt.value;
      select.dispatchEvent(new Event('change'));
      if (!standard.classList.contains('hidden')) return;
    }
  }

  // Open the lesson for the chapter on screen. A chapter with nothing yet
  // starts its lesson (the tutor sends the kickoff); a chapter taught before
  // resumes where it stopped.
  async function begin(force) {
    const ch = chapterOnScreen();
    if (!ch) return;
    showLessonContent();
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
    if (on) begin(true);
    else window.dsaTutor.sync();
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
    btn.title = 'Have the AI tutor teach this chapter interactively, with visualizations in a scratchpad';
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
    }
  }

  window.dsaTeacher = { start: () => setOn(true), stop: () => setOn(false), isOn: () => teach.on };
  init();
})();
