// CodeMirror 5 Python Editor Setup
// Uses CodeMirror 5 for stability (no ES module conflicts)

(function() {
  'use strict';

  // Python built-in completions for IntelliSense
  const pythonBuiltins = [
    // Built-in functions
    'print', 'len', 'range', 'enumerate', 'zip', 'map', 'filter', 'sorted',
    'reversed', 'sum', 'min', 'max', 'abs', 'int', 'float', 'str', 'bool',
    'list', 'dict', 'set', 'tuple', 'isinstance', 'type', 'input', 'open',
    'ord', 'chr', 'any', 'all', 'divmod', 'pow', 'round', 'hash', 'id',
    'hex', 'bin', 'oct', 'iter', 'next', 'slice', 'super', 'object',
    'getattr', 'setattr', 'hasattr', 'delattr', 'callable', 'repr', 'format',
    // Keywords
    'def', 'class', 'if', 'elif', 'else', 'for', 'while', 'break', 'continue',
    'return', 'yield', 'try', 'except', 'finally', 'raise', 'import', 'from',
    'as', 'with', 'pass', 'lambda', 'and', 'or', 'not', 'in', 'is', 'None',
    'True', 'False', 'global', 'nonlocal', 'assert', 'del', 'async', 'await',
    // Common methods
    'append', 'extend', 'insert', 'remove', 'pop', 'clear', 'index', 'count',
    'sort', 'reverse', 'copy', 'get', 'keys', 'values', 'items', 'update',
    'setdefault', 'add', 'discard', 'union', 'intersection', 'difference',
    'split', 'join', 'strip', 'replace', 'find', 'startswith', 'endswith',
    'lower', 'upper', 'isdigit', 'isalpha', 'isalnum', 'format',
    // Common modules
    'collections', 'heapq', 'bisect', 'itertools', 'functools', 'math',
    'defaultdict', 'Counter', 'deque', 'heappush', 'heappop', 'heapify',
    'bisect_left', 'bisect_right',
    // DSA helper classes
    'ListNode', 'TreeNode', 'GraphNode', 'MinHeap', 'MaxHeap',
    // Common variables
    'self', 'result', 'ans', 'res', 'left', 'right', 'mid', 'low', 'high',
    'start', 'end', 'head', 'tail', 'node', 'curr', 'prev', 'next', 'root',
    'stack', 'queue', 'visited', 'seen', 'memo', 'dp', 'graph', 'adj',
  ];

  // OxigenLang keywords, built-ins and the modules the sandbox allows
  const oxigenKeywords = [
    'fun', 'give', 'option', 'unless', 'when', 'then', 'each', 'in', 'repeat',
    'skip', 'stop', 'struct', 'includes', 'enum', 'hide', 'self', 'pattern',
    'choose', 'guard', 'fail', 'introduce', 'intro', 'from', 'as', 'and', 'or',
    'not', 'main', 'within', 'diverge', 'converge',
  ];
  const oxigenBuiltins = [
    'println', 'print', 'len', 'push', 'first', 'last', 'rest', 'has', 'range',
    'int', 'float', 'str', 'ord', 'chr', 'chars', 'set', 'tuple', 'keys',
    'values', 'insert', 'remove', 'type', 'error', 'is_error', 'is_value',
  ];
  const oxigenWords = oxigenKeywords.concat(oxigenBuiltins, [
    'True', 'False', 'None', 'generic', 'array', 'map',
    'math', 'strings', 'array', 'random', 'hash', 'regex', 'time', 'datetime',
    'ListNode', 'TreeNode', 'Node',
    'result', 'left', 'right', 'mid', 'low', 'high', 'head', 'node', 'root',
    'stack', 'queue', 'visited', 'seen', 'memo', 'dp', 'graph',
  ]);

  // Completion from a fixed word list plus the identifiers already in the
  // editor. The word being typed is left out, so finishing a word and
  // pressing Enter starts a new line instead of "accepting" it.
  function wordHint(words) {
    return function(editor) {
      const cur = editor.getCursor();
      const token = editor.getTokenAt(cur);
      const start = token.start;
      const end = cur.ch;
      const word = token.string.slice(0, end - start);
      const empty = { list: [], from: cur, to: cur };
      if (!/^[A-Za-z_]\w*$/.test(word)) return empty;

      const own = editor.getValue().match(/[A-Za-z_]\w+/g) || [];
      const seen = {};
      const list = words.concat(own).filter(function(item) {
        if (seen[item] || item === word) return false;
        seen[item] = true;
        return item.toLowerCase().startsWith(word.toLowerCase());
      });

      return {
        list: list,
        from: CodeMirror.Pos(cur.line, start),
        to: CodeMirror.Pos(cur.line, end)
      };
    };
  }

  const pythonHint = wordHint(pythonBuiltins);
  const oxigenHint = wordHint(oxigenWords);

  // Editor settings per language: CodeMirror mode and completion
  const languages = {
    python: { mode: 'python', hint: pythonHint },
    oxigen: { mode: 'oxigen', hint: oxigenHint },
  };

  if (typeof CodeMirror !== 'undefined') {
    CodeMirror.registerHelper('hint', 'python', pythonHint);
    CodeMirror.registerHelper('hint', 'oxigen', oxigenHint);

    // OxigenLang highlighting (needs the simple-mode addon)
    if (CodeMirror.defineSimpleMode) {
      const words = function(list) { return new RegExp('(?:' + list.join('|') + ')\\b'); };
      CodeMirror.defineSimpleMode('oxigen', {
        start: [
          { regex: /\/\/.*/, token: 'comment' },
          { regex: /\/\*/, token: 'comment', next: 'comment' },
          { regex: /#\[[^\]]*\]/, token: 'meta' },
          { regex: /"""/, token: 'string', next: 'tripleString' },
          { regex: /"(?:[^\\"]|\\.)*"?/, token: 'string' },
          { regex: /'(?:[^\\']|\\.)'/, token: 'string' },
          { regex: /(fun)(\s+)([A-Za-z_]\w*)/, token: ['keyword', null, 'def'] },
          { regex: /(struct|enum|pattern)(\s+)([A-Za-z_]\w*)/, token: ['keyword', null, 'def'] },
          { regex: words(oxigenKeywords), token: 'keyword' },
          { regex: /(?:True|False|None)\b/, token: 'atom' },
          { regex: words(oxigenBuiltins), token: 'builtin' },
          { regex: /<(?:int|str|float|bool|char|byte|uint|array|map|set|tuple|generic|None|[A-Z]\w*)>/, token: 'type' },
          { regex: /0x[0-9a-fA-F_]+|\d[\d_]*(?:\.\d+)?(?:[eE][+-]?\d+)?/, token: 'number' },
          { regex: /[-+\/*=<>!%&|^]+|:=/, token: 'operator' },
          { regex: /[{\[(]/, indent: true },
          { regex: /[}\])]/, dedent: true },
          { regex: /[A-Za-z_]\w*/, token: 'variable' },
        ],
        comment: [
          { regex: /.*?\*\//, token: 'comment', next: 'start' },
          { regex: /.*/, token: 'comment' },
        ],
        tripleString: [
          { regex: /.*?"""/, token: 'string', next: 'start' },
          { regex: /.*/, token: 'string' },
        ],
        meta: {
          lineComment: '//',
          blockCommentStart: '/*',
          blockCommentEnd: '*/',
          electricInput: /^\s*[}\])]$/,
        },
      });
    }
  }

  // Store editor instances globally
  window.dsaEditors = {};

  // CodeMirror measures character widths once and caches them. The editor is
  // often created before IBM Plex Mono has loaded, and measuring the fallback
  // font leaves the cursor drifting off the text, so re-measure once the web
  // fonts arrive.
  function refreshAll() {
    Object.keys(window.dsaEditors).forEach(function(id) {
      window.dsaEditors[id].refresh();
    });
  }
  if (document.fonts) {
    document.fonts.addEventListener('loadingdone', refreshAll);
    document.fonts.ready.then(refreshAll);
  }

  // Width of the line padding in styles.css (.cm-s-paper pre.CodeMirror-line)
  const LINE_PADDING = 12;

  // Create a CodeMirror editor
  function createEditor(containerId, options) {
    options = options || {};
    const container = document.getElementById(containerId);
    if (!container) {
      console.warn('Editor container ' + containerId + ' not found');
      return null;
    }

    // Check if CodeMirror is loaded
    if (typeof CodeMirror === 'undefined') {
      console.warn('CodeMirror not loaded yet');
      return null;
    }

    // Clear any existing content
    container.innerHTML = '';

    // Create textarea for CodeMirror
    const textarea = document.createElement('textarea');
    container.appendChild(textarea);

    // Get initial content from options
    const initialContent = options.content || '';

    const language = languages[options.language] || languages.python;

    // Create CodeMirror instance
    const editor = CodeMirror.fromTextArea(textarea, {
      mode: language.mode,
      theme: 'paper',
      lineNumbers: true,
      indentUnit: 4,
      tabSize: 4,
      indentWithTabs: false,
      smartIndent: true,
      electricChars: true,
      matchBrackets: true,
      autoCloseBrackets: true,
      styleActiveLine: true,
      lineWrapping: true,
      extraKeys: {
        // Spaces to the next tab stop, or indent the selected lines
        'Tab': function(cm) {
          if (cm.somethingSelected()) {
            cm.indentSelection('add');
            return;
          }
          const unit = cm.getOption('indentUnit');
          cm.replaceSelection(' '.repeat(unit - cm.getCursor().ch % unit), 'end');
        },
        // In leading whitespace, delete back to the previous tab stop
        'Backspace': function(cm) {
          const cur = cm.getCursor();
          const before = cm.getLine(cur.line).slice(0, cur.ch);
          if (cm.somethingSelected() || cur.ch === 0 || /\S/.test(before)) {
            return CodeMirror.Pass;
          }
          const unit = cm.getOption('indentUnit');
          const to = Math.floor((cur.ch - 1) / unit) * unit;
          cm.replaceRange('', CodeMirror.Pos(cur.line, to), cur, '+delete');
        },
        'Shift-Tab': function(cm) {
          cm.indentSelection('subtract');
        },
        'Ctrl-Space': 'autocomplete',
        'Ctrl-/': function(cm) {
          cm.toggleComment();
        },
        'Cmd-/': function(cm) {
          cm.toggleComment();
        }
      },
      hintOptions: {
        hint: language.hint,
        completeSingle: false
      }
    });

    // Set initial content
    editor.setValue(initialContent);

    // Auto-show hints while typing, but not inside comments or strings
    editor.on('inputRead', function(cm, change) {
      if (change.text[0] && /\w/.test(change.text[0])) {
        const cur = cm.getCursor();
        const token = cm.getTokenAt(cur);
        if (token.string.length >= 2 && !/comment|string/.test(token.type || '')) {
          cm.showHint({ completeSingle: false });
        }
      }
    });

    // Wrapped lines continue under their own indentation, not at column 0
    editor.on('renderLine', function(cm, line, elt) {
      const off = CodeMirror.countColumn(line.text, null, cm.getOption('tabSize')) * cm.defaultCharWidth();
      elt.style.textIndent = '-' + off + 'px';
      elt.style.paddingLeft = (LINE_PADDING + off) + 'px';
    });

    // CodeMirror only re-measures on window resizes; also do it when the
    // editor's own box changes, e.g. it was created hidden or a panel moved.
    if (window.ResizeObserver) {
      let size = '';
      new ResizeObserver(function() {
        const wrapper = editor.getWrapperElement();
        const now = wrapper.offsetWidth + 'x' + wrapper.offsetHeight;
        if (now !== size && wrapper.offsetWidth > 0) {
          size = now;
          editor.refresh();
        }
      }).observe(editor.getWrapperElement());
    }

    // Handle change callback
    if (options.onChange) {
      editor.on('change', function() {
        options.onChange(editor.getValue());
      });
    }

    // Store reference
    window.dsaEditors[containerId] = editor;

    // Refresh after a short delay to ensure proper rendering
    setTimeout(function() {
      editor.refresh();
    }, 100);

    return editor;
  }

  // Get editor content
  function getEditorContent(containerId) {
    const editor = window.dsaEditors[containerId];
    if (editor && editor.getValue) {
      return editor.getValue();
    }
    return '';
  }

  // Set editor content
  function setEditorContent(containerId, content) {
    const editor = window.dsaEditors[containerId];
    if (editor && editor.setValue) {
      editor.setValue(content);
    }
  }

  // Switch an editor between Python and Oxigen highlighting and completion
  function setEditorLanguage(containerId, name) {
    const editor = window.dsaEditors[containerId];
    const language = languages[name] || languages.python;
    if (editor) {
      editor.setOption('mode', language.mode);
      editor.setOption('hintOptions', { hint: language.hint, completeSingle: false });
    }
  }

  // Expose functions globally
  window.createPythonEditor = createEditor;
  window.setEditorLanguage = setEditorLanguage;
  window.getEditorContent = getEditorContent;
  window.setEditorContent = setEditorContent;

  // Signal that editor module is ready
  window.editorModuleReady = true;

  console.log('CodeMirror editor module loaded');
})();
