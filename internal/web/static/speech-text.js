// Code read aloud as words. Voices read "s[right]" and "len(count) > 2"
// symbol by symbol, or skip the symbols, and the tutor's replies and the
// lessons are full of such snippets. spokenText rewrites them the way a
// person would say them ("s at right", "the length of count greater than
// 2"), always the same way, so a name sounds the same every time.
//
// Every rule is local (it looks only at the snippet it rewrites), so a
// growing reply converts to text whose start doesn't change as more of the
// reply arrives: the tutor's speech queue depends on that.

(function (root) {
  // Longest first, so "<=" wins over "<".
  const CODE_OPS = [
    ['**', ' to the power of '], ['//', ' integer divided by '],
    ['<=', ' less than or equal to '], ['>=', ' greater than or equal to '],
    ['==', ' equals '], ['!=', ' not equal to '],
    ['+=', ' plus equals '], ['-=', ' minus equals '], ['*=', ' times equals '], ['/=', ' divided equals '],
    ['&&', ' and '], ['||', ' or '], ['->', ' to '], ['=>', ' to '],
    ['++', ' plus plus '], ['--', ' minus minus '],
    ['<', ' less than '], ['>', ' greater than '], ['=', ' equals '],
    ['+', ' plus '], ['-', ' minus '], ['*', ' times '], ['/', ' divided by '], ['%', ' mod '], ['!', ' not '],
  ];
  const escape = (s) => s.replace(/[.*+?^${}()|[\]\\/-]/g, '\\$&');
  const opRe = new RegExp(CODE_OPS.map(([op]) => escape(op)).join('|'), 'g');
  const opWord = Object.fromEntries(CODE_OPS);

  // Rewrites brackets and calls from the innermost out.
  function brackets(s) {
    for (let i = 0; i < 6; i++) {
      const before = s;
      s = s
        .replace(/\[\s*::\s*-1\s*\]/g, ' reversed ')
        .replace(/\[\s*\]/g, ' an empty list ')
        .replace(/(\w)\[\s*:\s*([^[\]:]+)\]/g, '$1 up to $2')
        .replace(/(\w)\[([^[\]:]+):\s*\]/g, '$1 from $2 on')
        .replace(/(\w)\[([^[\]:]+):([^[\]:]+)\]/g, '$1 from $2 to $3')
        .replace(/\]\s*\[/g, ', ')
        .replace(/(\w)\[([^[\]]+)\]/g, '$1 at $2')
        .replace(/\[([^[\]]*)\]/g, ' $1 ')
        .replace(/\blen\(([^()]+)\)/g, ' the length of $1 ')
        .replace(/(\w)\(\s*\)/g, '$1')
        .replace(/(\w)\(([^()]+)\)/g, '$1 of $2')
        .replace(/\(([^()]*)\)/g, ' $1 ');
      if (s === before) break;
    }
    return s;
  }

  // A code snippet as words.
  function spokenCode(code) {
    let s = String(code)
      .replace(/\^2\b|²/g, ' squared ')
      .replace(/\^3\b|³/g, ' cubed ')
      .replace(/\^/g, ' to the power of ')
      .replace(/([A-Za-z_)\]])\.(?=[A-Za-z_])/g, '$1 dot ');
    s = brackets(s);
    s = s.replace(opRe, (op) => opWord[op])
      .replace(/([a-z0-9])([A-Z])/g, '$1 $2') // maxLen -> max Len
      .replace(/_+/g, ' ')                    // max_len -> max len
      .replace(/[{}[\]();:,#$@\\|&~"'`]+/g, (m) => (m.includes(',') ? ', ' : ' '));
    return s.replace(/\s+([,.!?])/g, '$1').replace(/\s+/g, ' ').trim();
  }

  // Bare code in prose: an identifier followed by [index] or (args), like
  // s[right] or len(count). "item(s)" stays as written.
  const codeToken = /\b[A-Za-z_]\w*(?:\.[A-Za-z_]\w*)*(?:\[[^[\]\n]*\]|\([^()\n]*\))+/g;
  const snakeToken = /\b[A-Za-z]\w*_\w+\b/g;
  const camelToken = /\b[a-z]+[A-Z]\w*\b/g;

  function spokenText(text) {
    return String(text)
      .replace(/\[([^\]\n]+)\]\([^)\s]+\)/g, '$1') // Markdown links: just the words
      .replace(/`([^`\n]+)`/g, (_, code) => spokenCode(code))
      .replace(codeToken, (m) => (/\((e?s)\)$/.test(m) && !m.includes('[') ? m : spokenCode(m)))
      .replace(snakeToken, (m) => spokenCode(m))
      .replace(camelToken, (m) => spokenCode(m))
      .replace(/(\w)\^(\w+)/g, (_, a, b) => `${a} ${spokenCode(`^${b}`)}`)
      .replace(/²/g, ' squared').replace(/³/g, ' cubed')
      .replace(/[ \t]*(?:→|->|=>)[ \t]*/g, ' to ')
      .replace(/[ \t]*≤[ \t]*/g, ' less than or equal to ')
      .replace(/[ \t]*≥[ \t]*/g, ' greater than or equal to ')
      .replace(/[ \t]*≠[ \t]*/g, ' not equal to ')
      // Comparisons and arithmetic between operands in prose ("a < b",
      // "right - left + 1"). A line-start ">" (a Markdown quote) and a
      // dash between words ("left - and right") are left alone.
      .replace(/(?<=[\w)\]])[ \t]*(<=|>=|==|!=|<|>|=)[ \t]*(?=[\w(-])/g, (_, op) => opWord[op])
      .replace(/(?<=[\w)\]])[ \t]*\+[ \t]*(?=[\w(])/g, ' plus ')
      .replace(/(?<=[\w)\]]) \* (?=[\w(])/g, ' times ')
      .replace(/(?<=\b(?:\d+|[a-z]))[ \t]+-[ \t]+(?=(?:\d+|[a-z])\b)/g, ' minus ')
      .replace(/[ \t]{2,}/g, ' ');
  }

  root.spokenText = spokenText;
  root.spokenCode = spokenCode;
  if (typeof module !== 'undefined') module.exports = { spokenText, spokenCode };
})(typeof window !== 'undefined' ? window : globalThis);
