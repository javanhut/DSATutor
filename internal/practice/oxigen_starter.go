package practice

import (
	"strings"
)

// OxigenStarter returns the starter code for solving p in Oxigen. A problem
// can set StarterCodeOxigen by hand; otherwise it is derived from the Python
// StarterCode, keeping every function, class and parameter name so the judge
// calls Oxigen code exactly as it calls Python code.
func OxigenStarter(p *Problem) string {
	if p.StarterCodeOxigen != "" {
		return p.StarterCodeOxigen
	}
	funcs, classes := parsePythonStarter(p.StarterCode)

	var b strings.Builder
	for _, note := range providedStructs(p.Harness) {
		b.WriteString("// " + note + "\n")
	}
	if b.Len() > 0 {
		b.WriteString("\n")
	}

	for i, c := range classes {
		if i > 0 {
			b.WriteString("\n")
		}
		writeOxigenClass(&b, c, p.Harness.Mode == "design")
	}
	for i, f := range funcs {
		if i > 0 || len(classes) > 0 {
			b.WriteString("\n")
		}
		b.WriteString("fun " + f.name + "(" + strings.Join(f.params, ", ") + ") {\n")
		b.WriteString("    // Write your solution here\n")
		b.WriteString("}\n")
	}
	return b.String()
}

// providedStructs describes the node structs the judge adds to the
// learner's file for this problem.
func providedStructs(h Harness) []string {
	kinds := map[string]bool{h.ReturnType: true}
	for _, t := range h.ParamTypes {
		kinds[strings.SplitN(t, ":", 2)[0]] = true
	}
	var notes []string
	if kinds["ListNode"] || kinds["ListNode[]"] || kinds["CycleList"] {
		notes = append(notes, "ListNode(val, next) is provided; ListNode(5) leaves next as None.")
	}
	if kinds["TreeNode"] || kinds["TreeNodeRef"] || kinds["TreeNodeVal"] || (h.Mode == "roundtrip" && h.ClassName != "") {
		notes = append(notes, "TreeNode(val, left, right) is provided; TreeNode(5) leaves both children as None.")
	}
	if kinds["GraphNode"] {
		notes = append(notes, "Node(val, neighbors) is provided; neighbors is an array of Node.")
	}
	if kinds["RandomList"] {
		notes = append(notes, "Node(val, next, random) is provided.")
	}
	if kinds["CycleList"] || kinds["TreeNodeRef"] {
		notes = append(notes, "== compares structs field by field rather than by identity.")
	}
	return notes
}

type pyFunc struct {
	name   string
	params []string
}

type pyClass struct {
	name    string
	methods []pyFunc
}

// parsePythonStarter finds the top-level functions and the classes with
// their methods. Starter code is small and regular, so a line scan is enough.
func parsePythonStarter(src string) ([]pyFunc, []pyClass) {
	var funcs []pyFunc
	var classes []pyClass
	inClass := false
	for _, line := range strings.Split(src, "\n") {
		trimmed := strings.TrimSpace(line)
		indented := len(line) > 0 && (line[0] == ' ' || line[0] == '\t')
		switch {
		case !indented && strings.HasPrefix(trimmed, "class "):
			name := strings.TrimPrefix(trimmed, "class ")
			name = strings.TrimSpace(strings.SplitN(strings.SplitN(name, "(", 2)[0], ":", 2)[0])
			classes = append(classes, pyClass{name: name})
			inClass = true
		case !indented && strings.HasPrefix(trimmed, "def "):
			inClass = false
			if f, ok := parsePythonDef(trimmed); ok {
				funcs = append(funcs, f)
			}
		case indented && inClass && strings.HasPrefix(trimmed, "def "):
			if f, ok := parsePythonDef(trimmed); ok {
				c := &classes[len(classes)-1]
				c.methods = append(c.methods, f)
			}
		case !indented && trimmed != "" && !strings.HasPrefix(trimmed, "#"):
			inClass = false
		}
	}
	return funcs, classes
}

func parsePythonDef(line string) (pyFunc, bool) {
	rest := strings.TrimPrefix(line, "def ")
	open := strings.Index(rest, "(")
	closeIdx := strings.LastIndex(rest, ")")
	if open < 0 || closeIdx < open {
		return pyFunc{}, false
	}
	f := pyFunc{name: strings.TrimSpace(rest[:open])}
	depth, start := 0, 0
	args := rest[open+1 : closeIdx]
	var parts []string
	for i, r := range args {
		switch r {
		case '[', '(', '{':
			depth++
		case ']', ')', '}':
			depth--
		case ',':
			if depth == 0 {
				parts = append(parts, args[start:i])
				start = i + 1
			}
		}
	}
	parts = append(parts, args[start:])
	for _, raw := range parts {
		p := strings.TrimSpace(raw)
		if i := strings.IndexAny(p, ":="); i >= 0 {
			p = strings.TrimSpace(p[:i])
		}
		p = strings.TrimLeft(p, "*")
		if p != "" && p != "self" {
			f.params = append(f.params, p)
		}
	}
	return f, true
}

// writeOxigenClass turns a Python class into a struct plus an includes
// block. __init__ becomes init(), which the judge calls right after it
// declares the zero-value instance; a design class always gets one.
func writeOxigenClass(b *strings.Builder, c pyClass, design bool) {
	b.WriteString("struct " + c.name + " {\n")
	b.WriteString("    // Add the fields you need, e.g. items <array>\n")
	b.WriteString("}\n\n")
	b.WriteString(c.name + " includes {\n")

	methods := c.methods
	hasInit := false
	for _, m := range methods {
		if m.name == "__init__" {
			hasInit = true
		}
	}
	if design && !hasInit {
		methods = append([]pyFunc{{name: "__init__"}}, methods...)
	}
	for i, m := range methods {
		if i > 0 {
			b.WriteString("\n")
		}
		name := m.name
		if name == "__init__" {
			name = "init"
		}
		b.WriteString("    fun " + name + "(" + strings.Join(m.params, ", ") + ") {\n")
		b.WriteString("    }\n")
	}
	b.WriteString("}\n")
}
