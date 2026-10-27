// Page shell: theme toggle and the home page's shortcuts.
// The shortcuts click the existing sidebar controls so app.js stays the
// only place that knows how modes switch.

(function () {
  'use strict';

  const root = document.documentElement;
  const darkQuery = window.matchMedia('(prefers-color-scheme: dark)');

  function currentTheme() {
    return root.dataset.theme || (darkQuery.matches ? 'dark' : 'light');
  }

  function syncThemeToggle() {
    const theme = currentTheme();
    // Pin the resolved theme so the icon rules in styles.css can match it.
    root.dataset.theme = theme;
    const label = document.getElementById('theme-toggle-label');
    if (label) label.textContent = theme === 'dark' ? 'Light theme' : 'Dark theme';
  }

  function toggleTheme() {
    const next = currentTheme() === 'dark' ? 'light' : 'dark';
    root.dataset.theme = next;
    try { localStorage.setItem('dsa-theme', next); } catch (e) { /* private mode */ }
    syncThemeToggle();
  }

  function clickById(id) {
    const target = document.getElementById(id);
    if (target) target.click();
  }

  function startFirstChapter() {
    const first = document.querySelector('#chapters .chapter-item');
    if (first) first.click();
  }

  const themeBtn = document.getElementById('theme-toggle');
  if (themeBtn) themeBtn.addEventListener('click', toggleTheme);

  // Follow the OS while the reader hasn't picked a theme themselves.
  darkQuery.addEventListener('change', () => {
    let stored = null;
    try { stored = localStorage.getItem('dsa-theme'); } catch (e) { /* private mode */ }
    if (!stored) {
      delete root.dataset.theme;
      syncThemeToggle();
    }
  });

  const handlers = {
    'home-start': startFirstChapter,
    'home-practice': () => clickById('practice-toggle'),
    'home-sandbox': () => clickById('sandbox-toggle'),
  };
  Object.entries(handlers).forEach(([id, fn]) => {
    const btn = document.getElementById(id);
    if (btn) btn.addEventListener('click', fn);
  });

  syncThemeToggle();
})();
