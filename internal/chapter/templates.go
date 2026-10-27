package chapter

// DefaultChapters returns the curriculum in learning order. Each chapter lives
// in its own chNN_*.go file; docs/pedagogy.md describes what a chapter must
// contain and why.
func DefaultChapters() []Chapter {
	return []Chapter{
		chapter01(),
		chapter02(),
		chapter03(),
		chapter04(),
		chapter05(),
		chapter06(),
		chapter07(),
		chapter08(),
		chapter09(),
		chapter10(),
		chapter11(),
		chapter12(),
		chapter13(),
		chapter14(),
		chapter15(),
		chapter16(),
		chapter17(),
		chapter18(),
		chapter19(),
		chapter20(),
		chapter21(),
	}
}
