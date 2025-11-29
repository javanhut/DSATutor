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

  // Python hint function
  function pythonHint(editor) {
    const cur = editor.getCursor();
    const token = editor.getTokenAt(cur);
    const start = token.start;
    const end = cur.ch;
    const word = token.string.slice(0, end - start);

    const list = pythonBuiltins.filter(function(item) {
      return item.toLowerCase().startsWith(word.toLowerCase());
    }).map(function(item) {
      return {
        text: item,
        displayText: item
      };
    });

    return {
      list: list,
      from: CodeMirror.Pos(cur.line, start),
      to: CodeMirror.Pos(cur.line, end)
    };
  }

  // Register Python hint
  if (typeof CodeMirror !== 'undefined') {
    CodeMirror.registerHelper('hint', 'python', pythonHint);
  }

  // Store editor instances globally
  window.dsaEditors = {};

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

    // Create CodeMirror instance
    const editor = CodeMirror.fromTextArea(textarea, {
      mode: 'python',
      theme: 'material-darker',
      lineNumbers: true,
      indentUnit: 4,
      tabSize: 4,
      indentWithTabs: false,
      smartIndent: true,
      electricChars: true,
      matchBrackets: true,
      autoCloseBrackets: true,
      lineWrapping: true,
      extraKeys: {
        'Tab': function(cm) {
          if (cm.somethingSelected()) {
            cm.indentSelection('add');
          } else {
            cm.replaceSelection('    ', 'end');
          }
        },
        'Shift-Tab': function(cm) {
          cm.indentSelection('subtract');
        },
        'Ctrl-Space': 'autocomplete',
        'Ctrl-/': function(cm) {
          cm.toggleComment();
        }
      },
      hintOptions: {
        hint: pythonHint,
        completeSingle: false
      }
    });

    // Set initial content
    editor.setValue(initialContent);

    // Auto-show hints while typing
    editor.on('inputRead', function(cm, change) {
      if (change.text[0] && /\w/.test(change.text[0])) {
        const cur = cm.getCursor();
        const token = cm.getTokenAt(cur);
        if (token.string.length >= 2) {
          cm.showHint({ completeSingle: false });
        }
      }
    });

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

  // Expose functions globally
  window.createPythonEditor = createEditor;
  window.getEditorContent = getEditorContent;
  window.setEditorContent = setEditorContent;

  // Signal that editor module is ready
  window.editorModuleReady = true;

  console.log('CodeMirror editor module loaded');
})();
