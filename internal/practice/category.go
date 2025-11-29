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

// Blind75Categories returns the standard Blind 75 category organization.
var Blind75Categories = []Category{
	{
		ID:          "arrays-hashing",
		Name:        "Arrays & Hashing",
		Slug:        "arrays-hashing",
		Description: "Array manipulation and hash table problems",
		Order:       1,
		ChapterRefs: []int{2, 5},
	},
	{
		ID:          "two-pointers",
		Name:        "Two Pointers",
		Slug:        "two-pointers",
		Description: "Problems using the two-pointer technique",
		Order:       2,
		ChapterRefs: []int{2},
	},
	{
		ID:          "sliding-window",
		Name:        "Sliding Window",
		Slug:        "sliding-window",
		Description: "Sliding window pattern problems",
		Order:       3,
		ChapterRefs: []int{2},
	},
	{
		ID:          "stack",
		Name:        "Stack",
		Slug:        "stack",
		Description: "Stack-based problems",
		Order:       4,
		ChapterRefs: []int{3},
	},
	{
		ID:          "binary-search",
		Name:        "Binary Search",
		Slug:        "binary-search",
		Description: "Binary search and variants",
		Order:       5,
		ChapterRefs: []int{1, 4},
	},
	{
		ID:          "linked-list",
		Name:        "Linked List",
		Slug:        "linked-list",
		Description: "Linked list manipulation problems",
		Order:       6,
		ChapterRefs: []int{2},
	},
	{
		ID:          "trees",
		Name:        "Trees",
		Slug:        "trees",
		Description: "Binary tree and BST problems",
		Order:       7,
		ChapterRefs: []int{11},
	},
	{
		ID:          "tries",
		Name:        "Tries",
		Slug:        "tries",
		Description: "Trie (prefix tree) problems",
		Order:       8,
		ChapterRefs: []int{11},
	},
	{
		ID:          "heap-priority-queue",
		Name:        "Heap / Priority Queue",
		Slug:        "heap-priority-queue",
		Description: "Heap and priority queue problems",
		Order:       9,
		ChapterRefs: []int{7},
	},
	{
		ID:          "backtracking",
		Name:        "Backtracking",
		Slug:        "backtracking",
		Description: "Backtracking and recursive enumeration",
		Order:       10,
		ChapterRefs: []int{3, 4},
	},
	{
		ID:          "graphs",
		Name:        "Graphs",
		Slug:        "graphs",
		Description: "Graph traversal and algorithms",
		Order:       11,
		ChapterRefs: []int{6, 7},
	},
	{
		ID:          "1d-dp",
		Name:        "1-D Dynamic Programming",
		Slug:        "1d-dp",
		Description: "One-dimensional DP problems",
		Order:       12,
		ChapterRefs: []int{9},
	},
	{
		ID:          "2d-dp",
		Name:        "2-D Dynamic Programming",
		Slug:        "2d-dp",
		Description: "Two-dimensional DP problems",
		Order:       13,
		ChapterRefs: []int{9},
	},
	{
		ID:          "greedy",
		Name:        "Greedy",
		Slug:        "greedy",
		Description: "Greedy algorithm problems",
		Order:       14,
		ChapterRefs: []int{8},
	},
	{
		ID:          "intervals",
		Name:        "Intervals",
		Slug:        "intervals",
		Description: "Interval scheduling and merging",
		Order:       15,
		ChapterRefs: []int{8},
	},
	{
		ID:          "math-geometry",
		Name:        "Math & Geometry",
		Slug:        "math-geometry",
		Description: "Mathematical and geometric problems",
		Order:       16,
		ChapterRefs: []int{1},
	},
	{
		ID:          "bit-manipulation",
		Name:        "Bit Manipulation",
		Slug:        "bit-manipulation",
		Description: "Bitwise operation problems",
		Order:       17,
		ChapterRefs: []int{},
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
