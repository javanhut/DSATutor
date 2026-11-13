package tutor

import (
	"regexp"
	"strings"
	"unicode"
)

// Models narrate their tool use ("Let me write it onto the lesson page.",
// "I'm waiting for your answer.") even when told not to. In Teacher Mode the
// learner sees every tool's result, so those sentences are noise, and read
// aloud they are worse. StripStageDirections drops them, along with any
// "[Your last reply also showed: ...]" record a model copies from its
// history (from the marker to the end of its line: the record can contain
// brackets of its own, like nums[2]).

var (
	shownRecord = regexp.MustCompile(`(?i)[\[(]?(?:your last reply also showed|shown with this reply):[^\n]*`)
	extraBlank  = regexp.MustCompile(`\n{3,}`)
)

// sentences splits text into sentences, each keeping its trailing space. A
// sentence ends at a newline, or at . ! ? followed by space and a capital
// letter, digit or Markdown mark (so "vs. Lists" and "e.g. a" stay whole).
func sentences(text string) []string {
	var out []string
	start := 0
	rs := []rune(text)
	for i := 0; i < len(rs); i++ {
		end := -1
		switch {
		case rs[i] == '\n':
			j := i
			for j < len(rs) && (rs[j] == '\n' || rs[j] == ' ') {
				j++
			}
			end = j
		case rs[i] == '.' || rs[i] == '!' || rs[i] == '?':
			j := i + 1
			for j < len(rs) && strings.ContainsRune(".!?\"'”’)", rs[j]) {
				j++
			}
			k := j
			for k < len(rs) && rs[k] == ' ' {
				k++
			}
			if rs[i] == '.' && abbreviation(rs[start:i]) {
				break
			}
			if k == len(rs) || (k > j && (unicode.IsUpper(rs[k]) || unicode.IsDigit(rs[k]) || strings.ContainsRune("*_[`(\"“", rs[k]))) {
				end = k
			}
		}
		if end > 0 {
			out = append(out, string(rs[start:end]))
			start = end
			i = end - 1
		}
	}
	if start < len(rs) {
		out = append(out, string(rs[start:]))
	}
	return out
}

var narrationStarts = []string{
	"let me ", "now let me ", "next, let me ", "first, let me ", "i'll ", "i will ", "i've ", "i have ",
	"i'm going to ", "i am going to ", "i'm waiting", "i am waiting", "waiting for your", "here's the sketch",
	"here is the sketch", "here's a sketch", "take a look at the scratchpad", "take a look at the board",
}

var narrationTopics = []string{
	"lesson page", "scratchpad", "whiteboard", "board", "sketch", "draw", "section", "your answer",
	"question", "waiting", "write", "wrote", "written", "show you",
}

// StripStageDirections removes sentences that narrate the lesson's
// mechanics rather than teach.
func StripStageDirections(text string) string {
	text = shownRecord.ReplaceAllString(text, "")
	var b strings.Builder
	for _, sent := range sentences(text) {
		if isStageDirection(sent) {
			continue
		}
		b.WriteString(sent)
	}
	out := extraBlank.ReplaceAllString(b.String(), "\n\n")
	return strings.TrimSpace(out)
}

// abbreviation reports whether the text before a period ends in a word
// that is not a sentence end: "vs", "e.g", "etc", or a single letter.
func abbreviation(before []rune) bool {
	i := len(before)
	for i > 0 && (unicode.IsLetter(before[i-1]) || before[i-1] == '.') {
		i--
	}
	w := strings.ToLower(string(before[i:]))
	switch w {
	case "vs", "e.g", "i.e", "etc", "eg", "ie", "approx", "dr", "mr", "ms", "mrs", "no":
		return true
	}
	return len([]rune(w)) == 1
}

func isStageDirection(sentence string) bool {
	s := strings.ToLower(strings.TrimSpace(sentence))
	s = strings.TrimLeft(s, "*_ ")
	if s == "" {
		return false
	}
	starts := false
	for _, p := range narrationStarts {
		if strings.HasPrefix(s, p) {
			starts = true
			break
		}
	}
	if !starts {
		return false
	}
	for _, t := range narrationTopics {
		if strings.Contains(s, t) {
			return true
		}
	}
	return false
}
