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
        quick: [
          ['Next', "Got it. Let's move on to the next idea."],
          ['Show me', 'Show me that in the scratchpad with a small example.'],
          ["I'm confused", "I'm confused. Explain it another way, with a smaller example."],
          ['Quiz me', 'Quiz me on this before we move on.'],
          ['Let me try', 'Give me a small exercise to try in the scratchpad.'],
        ],
        ctx: { view: 'chapter', chapterNumber: ch.number, teach: true },
      };
    },
  });

  async function begin() {
    const ch = chapterOnScreen();
    if (!ch) return;
    await window.dsaTutor.open();
    window.dsaTutor.sync();
    // A fresh chapter starts the lesson; a chapter taught before resumes.
    const convo = window.dsaTutor.conversation(`teach:${ch.number}`);
    if (!convo.length && !window.dsaTutor.isBusy()) {
      window.dsaTutor.send('(Start the lesson on this chapter from the beginning. Keep it to one idea, then ask me a question.)', { hidden: true });
    }
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
    if (on) begin();
    else window.dsaTutor.sync();
  }

  // Moving to another chapter with Teacher Mode on starts or resumes that
  // chapter's lesson.
  document.addEventListener('click', (e) => {
    if (!teach.on || !e.target.closest('.chapter-item, .chapter-link, #home-start')) return;
    setTimeout(() => { if (teach.on && chapterOnScreen()) begin(); }, 600);
  });

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
