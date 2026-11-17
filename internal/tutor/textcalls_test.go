package tutor

import (
	"strings"
	"testing"
)

// From a real reply: a self-closing highlight and a malformed ask_learner
// (note the stray quote), both written as text.
const textCallReply = `<highlight target="code" start_line="6" end_line="8" note="Shrink from the left edge &amp; record." pointer="hand"/>

Here's the mental template.

<ask_learner"> <parameter name="question">For at most 2 distinct characters, what is the while condition?</parameter> <parameter name="choices">["while s[right] in seen", "while len(count) > 2"]</parameter> </ask_learner>`

func TestTextCalls(t *testing.T) {
	text, calls := TextCalls(textCallReply)
	if strings.Contains(text, "<") || strings.TrimSpace(text) != "Here's the mental template." {
		t.Fatalf("text = %q", text)
	}
	if len(calls) != 2 {
		t.Fatalf("got %d calls, want 2", len(calls))
	}
	h := calls[0].Function
	if h.Name != "highlight" || intArg(h.Arguments, "start_line") != 6 || intArg(h.Arguments, "end_line") != 8 ||
		str(h.Arguments, "target") != "code" || str(h.Arguments, "note") != "Shrink from the left edge & record." {
		t.Errorf("highlight = %+v", h)
	}
	a := calls[1].Function
	if a.Name != "ask_learner" || !strings.HasPrefix(str(a.Arguments, "question"), "For at most") || len(listArg(a.Arguments, "choices")) != 2 {
		t.Errorf("ask_learner = %+v", a)
	}
}

func TestTextCallsOtherFormats(t *testing.T) {
	for _, in := range []string{
		"<tool_call>\n<function=open_tab>\n<parameter=tab>\ntests\n</parameter>\n</function>\n</tool_call>",
		`<function_calls><invoke name="open_tab"><parameter name="tab">tests</parameter></invoke></function_calls>`,
	} {
		text, calls := TextCalls("Look. " + in + " Done.")
		if strings.Contains(text, "<") || len(calls) != 1 || calls[0].Function.Name != "open_tab" || str(calls[0].Function.Arguments, "tab") != "tests" {
			t.Errorf("%q: text %q, calls %+v", in, text, calls)
		}
	}
	// Prose and code that merely contain < are left alone.
	for _, in := range []string{"if a < b and b <= c", "List<Integer> xs", "use the highlight tool"} {
		if text, calls := TextCalls(in); text != in || calls != nil {
			t.Errorf("%q changed to %q (%d calls)", in, text, len(calls))
		}
	}
}

func TestMarkupFilterStreaming(t *testing.T) {
	// Fed a few characters at a time, no tag (or part of one) gets through.
	for _, size := range []int{1, 3, 7, 50} {
		var f MarkupFilter
		var out strings.Builder
		for i := 0; i < len(textCallReply); i += size {
			chunk := textCallReply[i:min(i+size, len(textCallReply))]
			got := f.Write(chunk)
			if strings.Contains(got, "<") || strings.Contains(got, "parameter") {
				t.Fatalf("size %d: leaked %q", size, got)
			}
			out.WriteString(got)
		}
		out.WriteString(f.Flush())
		if strings.TrimSpace(out.String()) != "Here's the mental template." {
			t.Errorf("size %d: streamed %q", size, out.String())
		}
	}
	var f MarkupFilter
	if got := f.Write("so a < b, and "); got != "so a < b, and " {
		t.Errorf("comparison held back: %q", got)
	}
}
