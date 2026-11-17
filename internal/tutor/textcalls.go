package tutor

import (
	"encoding/json"
	"html"
	"regexp"
	"strings"
)

// Models sometimes write a tool call into their reply as markup instead of
// calling the tool: <highlight target="code" start_line="6"/>, or
// <ask_learner><parameter name="question">…</parameter></ask_learner>.
// Ollama only parses the format the model's own template defines, so these
// arrive as ordinary text, which the learner would see, keep in the chat
// and hear read aloud. TextCalls finds them so they can be run and cut out,
// and MarkupFilter keeps them out of the streamed text.

// markupTools are the tool names recognised as tags. Keep in step with the
// fn(...) definitions in tools.go and scratchpad.go.
var markupTools = []string{
	"highlight", "write_code", "open_tab", "select_test", "ask_learner", "write_lesson",
	"scratchpad_sketch", "scratchpad_visualize", "scratchpad_run_code",
}

var (
	// An opening tag of a call: a tool's own name, or a generic
	// <invoke name="tool"> / <function=tool>.
	callOpen = regexp.MustCompile(`<\s*(` + strings.Join(markupTools, "|") + `|invoke|function)\b([^<>]*)>`)
	// Tags that carry no call of their own: wrappers, and stray closers.
	callWrapper = regexp.MustCompile(`</?\s*(?:tool_calls?|function_calls?)\s*>|</\s*(?:` + strings.Join(markupTools, "|") + `|invoke|function|parameter)\s*>`)
	callAttr    = regexp.MustCompile(`([\w-]+)\s*=\s*(?:"([^"]*)"|'([^']*)'|([^\s"'/>]+))`)
	callParam   = regexp.MustCompile(`<parameter(?:\s+name\s*=\s*["']?|\s*=\s*["']?)([\w-]+)["']?\s*>([\s\S]*?)</parameter\s*>`)
	paramClose  = regexp.MustCompile(`</parameter\s*>`)
	partialTag  = regexp.MustCompile(`<\s*/?([A-Za-z_]*)([^<>]*)$`)
	extraBlanks = regexp.MustCompile(`\n{3,}`)
)

// markup is one tag (or tag with its body) found in text.
type markup struct {
	start, end int
	complete   bool      // closed (or self-closing); otherwise end is a best guess
	call       *ToolCall // nil for wrappers and stray closers
}

// nextMarkup finds the first call or wrapper tag in s.
func nextMarkup(s string) (markup, bool) {
	open := callOpen.FindStringSubmatchIndex(s)
	wrap := callWrapper.FindStringIndex(s)
	if wrap != nil && (open == nil || wrap[0] < open[0]) {
		return markup{start: wrap[0], end: wrap[1], complete: true}, true
	}
	if open == nil {
		return markup{}, false
	}
	tag := s[open[2]:open[3]]
	attrs := s[open[4]:open[5]]
	m := markup{start: open[0], end: open[1]}
	if strings.HasSuffix(strings.TrimSpace(attrs), "/") {
		m.complete = true
	} else {
		rest := s[open[1]:]
		if c := regexp.MustCompile(`</\s*` + tag + `\s*>`).FindStringIndex(rest); c != nil {
			m.end, m.complete = open[1]+c[1], true
		} else if cs := paramClose.FindAllStringIndex(rest, -1); cs != nil {
			m.end = open[1] + cs[len(cs)-1][1]
		}
	}
	m.call = parseCall(tag, attrs, s[open[1]:m.end])
	return m, true
}

// parseCall reads a call's name and arguments from its tag and body.
func parseCall(tag, attrs, body string) *ToolCall {
	args := map[string]any{}
	name := tag
	if tag == "invoke" || tag == "function" {
		name = ""
		if t := strings.TrimSpace(attrs); strings.HasPrefix(t, "=") {
			name = strings.Trim(strings.Fields(t[1:]+" ")[0], `"'`)
		}
	}
	for _, a := range callAttr.FindAllStringSubmatch(attrs, -1) {
		v := a[2] + a[3] + a[4]
		if (tag == "invoke" || tag == "function") && a[1] == "name" {
			name = v
			continue
		}
		args[a[1]] = argValue(html.UnescapeString(v))
	}
	for _, p := range callParam.FindAllStringSubmatch(body, -1) {
		args[p[1]] = argValue(p[2])
	}
	if !knownTool(name) {
		return nil
	}
	var c ToolCall
	c.Function.Name = name
	c.Function.Arguments = args
	return &c
}

// argValue keeps text as text (intArg reads numbers from strings) but
// decodes JSON lists and objects, which some tools take.
func argValue(v string) any {
	v = strings.TrimSpace(v)
	if strings.HasPrefix(v, "[") || strings.HasPrefix(v, "{") {
		var out any
		if json.Unmarshal([]byte(v), &out) == nil {
			return out
		}
	}
	return v
}

func knownTool(name string) bool {
	for _, t := range markupTools {
		if t == name {
			return true
		}
	}
	return false
}

// TextCalls removes tool-call markup from text and returns the cleaned text
// and the calls it held (calls to unknown tools are just removed).
func TextCalls(text string) (string, []ToolCall) {
	var out strings.Builder
	var calls []ToolCall
	for {
		m, ok := nextMarkup(text)
		if !ok {
			out.WriteString(text)
			break
		}
		out.WriteString(text[:m.start])
		if m.call != nil {
			calls = append(calls, *m.call)
		}
		text = text[m.end:]
	}
	if len(calls) == 0 && out.Len() == 0 {
		return "", nil
	}
	return extraBlanks.ReplaceAllString(out.String(), "\n\n"), calls
}

// MarkupFilter strips tool-call markup from streamed text. It holds back
// text from the start of a possible tag until the tag is complete (then
// drops it) or the stream ends.
type MarkupFilter struct {
	buf string
}

// maxHeldTag bounds how long an unterminated "<name" is held before it is
// treated as ordinary text.
const maxHeldTag = 200

// Write takes the next piece of the stream and returns what can be shown.
func (f *MarkupFilter) Write(s string) string {
	f.buf += s
	var out strings.Builder
	for {
		m, ok := nextMarkup(f.buf)
		if !ok {
			cut := len(f.buf)
			if at := partialStart(f.buf); at >= 0 {
				cut = at
			}
			out.WriteString(f.buf[:cut])
			f.buf = f.buf[cut:]
			return out.String()
		}
		out.WriteString(f.buf[:m.start])
		if !m.complete {
			f.buf = f.buf[m.start:]
			return out.String()
		}
		f.buf = f.buf[m.end:]
	}
}

// Flush returns the rest of the stream, cleaned.
func (f *MarkupFilter) Flush() string {
	s, _ := TextCalls(f.buf)
	f.buf = ""
	return s
}

// partialStart is where a tag that may still become markup starts at the
// end of s ("<highl", "<highlight target=\"co"), or -1. Comparisons like
// "a < b" are not held.
func partialStart(s string) int {
	loc := partialTag.FindStringSubmatchIndex(s)
	if loc == nil || len(s)-loc[0] > maxHeldTag {
		return -1
	}
	name, after := s[loc[2]:loc[3]], s[loc[4]:loc[5]]
	if (after == "" && tagPrefix(name)) || knownTool(name) || name == "invoke" || name == "function" {
		return loc[0]
	}
	return -1
}

// tagPrefix reports whether a tag name still being written could become one
// of the recognised tags.
func tagPrefix(name string) bool {
	for _, t := range append(markupTools, "invoke", "function", "parameter", "tool_call", "tool_calls", "function_call", "function_calls") {
		if strings.HasPrefix(t, name) {
			return true
		}
	}
	return false
}
