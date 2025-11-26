package chapter

// Registry keeps chapters discoverable and lets you swap in new templates easily.
type Registry struct {
	chapters map[int]Chapter
	order    []int
}

// NewRegistry builds a registry from a slice of chapters.
func NewRegistry(chapters []Chapter) *Registry {
	chaptersByNumber := make(map[int]Chapter, len(chapters))
	order := make([]int, 0, len(chapters))
	for _, ch := range chapters {
		chaptersByNumber[ch.Number] = ch
		order = append(order, ch.Number)
	}
	return &Registry{
		chapters: chaptersByNumber,
		order:    order,
	}
}

// Chapter returns a single chapter by number, if it exists.
func (r *Registry) Chapter(number int) (Chapter, bool) {
	ch, ok := r.chapters[number]
	return ch, ok
}

// All returns chapters in the order they were registered.
func (r *Registry) All() []Chapter {
	result := make([]Chapter, 0, len(r.order))
	for _, number := range r.order {
		result = append(result, r.chapters[number])
	}
	return result
}
