package practice

// Category groups related problems.
type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
	Order       int    `json:"order"`
	ChapterRefs []int  `json:"chapterRefs"`
	Icon        string `json:"icon,omitempty"`
}

// Blind75Categories lists the practice categories in curriculum order. Each
// category's first ChapterRef is the chapter that teaches its pattern.
var Blind75Categories = []Category{
	{
		ID:          "arrays-hashing",
		Name:        "Arrays & Hashing",
		Slug:        "arrays-hashing",
		Description: "Look things up in O(1) with hash maps and sets",
		Order:       1,
		ChapterRefs: []int{2},
	},
	{
		ID:          "two-pointers",
		Name:        "Two Pointers",
		Slug:        "two-pointers",
		Description: "Walk a sorted array or string from both ends",
		Order:       2,
		ChapterRefs: []int{3},
	},
	{
		ID:          "sliding-window",
		Name:        "Sliding Window",
		Slug:        "sliding-window",
		Description: "Grow and shrink a contiguous range in one pass",
		Order:       3,
		ChapterRefs: []int{4},
	},
	{
		ID:          "stack",
		Name:        "Stack",
		Slug:        "stack",
		Description: "Last-in-first-out, including monotonic stacks",
		Order:       4,
		ChapterRefs: []int{5},
	},
	{
		ID:          "binary-search",
		Name:        "Binary Search",
		Slug:        "binary-search",
		Description: "Halve the search space, including searching on the answer",
		Order:       5,
		ChapterRefs: []int{6},
	},
	{
		ID:          "linked-list",
		Name:        "Linked List",
		Slug:        "linked-list",
		Description: "Pointer manipulation, dummy heads and fast/slow pointers",
		Order:       6,
		ChapterRefs: []int{7},
	},
	{
		ID:          "trees",
		Name:        "Trees",
		Slug:        "trees",
		Description: "Recursive DFS, level-order BFS and BST properties",
		Order:       7,
		ChapterRefs: []int{9, 8},
	},
	{
		ID:          "tries",
		Name:        "Tries",
		Slug:        "tries",
		Description: "Prefix trees for word and prefix queries",
		Order:       8,
		ChapterRefs: []int{10},
	},
	{
		ID:          "heap-priority-queue",
		Name:        "Heap / Priority Queue",
		Slug:        "heap-priority-queue",
		Description: "Repeatedly take the smallest or largest item",
		Order:       9,
		ChapterRefs: []int{11},
	},
	{
		ID:          "backtracking",
		Name:        "Backtracking",
		Slug:        "backtracking",
		Description: "Build candidates step by step and undo dead ends",
		Order:       10,
		ChapterRefs: []int{12, 8},
	},
	{
		ID:          "graphs",
		Name:        "Graphs",
		Slug:        "graphs",
		Description: "Grid and graph traversal, topological sort, union-find",
		Order:       11,
		ChapterRefs: []int{13},
	},
	{
		ID:          "advanced-graphs",
		Name:        "Advanced Graphs",
		Slug:        "advanced-graphs",
		Description: "Weighted shortest paths and minimum spanning trees",
		Order:       12,
		ChapterRefs: []int{14},
	},
	{
		ID:          "1d-dp",
		Name:        "1-D Dynamic Programming",
		Slug:        "1d-dp",
		Description: "Reuse answers to smaller prefixes",
		Order:       13,
		ChapterRefs: []int{15},
	},
	{
		ID:          "2d-dp",
		Name:        "2-D Dynamic Programming",
		Slug:        "2d-dp",
		Description: "Tables indexed by two positions or a position and a budget",
		Order:       14,
		ChapterRefs: []int{16},
	},
	{
		ID:          "greedy",
		Name:        "Greedy",
		Slug:        "greedy",
		Description: "Commit to the locally best choice when it is provably safe",
		Order:       15,
		ChapterRefs: []int{17},
	},
	{
		ID:          "intervals",
		Name:        "Intervals",
		Slug:        "intervals",
		Description: "Sort by start or end, then sweep",
		Order:       16,
		ChapterRefs: []int{18},
	},
	{
		ID:          "math-geometry",
		Name:        "Math & Geometry",
		Slug:        "math-geometry",
		Description: "Matrix traversal and number manipulation",
		Order:       17,
		ChapterRefs: []int{19},
	},
	{
		ID:          "bit-manipulation",
		Name:        "Bit Manipulation",
		Slug:        "bit-manipulation",
		Description: "Work with the binary form of numbers",
		Order:       18,
		ChapterRefs: []int{20},
	},
}

// GetCategoryByID returns a category by its ID.
func GetCategoryByID(id string) *Category {
	for _, cat := range Blind75Categories {
		if cat.ID == id {
			return &cat
		}
	}
	return nil
}

// GetCategoryBySlug returns a category by its slug.
func GetCategoryBySlug(slug string) *Category {
	for _, cat := range Blind75Categories {
		if cat.Slug == slug {
			return &cat
		}
	}
	return nil
}

// GetCategoriesForChapter returns categories related to a chapter number.
func GetCategoriesForChapter(chapterNum int) []Category {
	result := make([]Category, 0)
	for _, cat := range Blind75Categories {
		for _, ref := range cat.ChapterRefs {
			if ref == chapterNum {
				result = append(result, cat)
				break
			}
		}
	}
	return result
}
