package practice

// BinarySearchDrills are the binary search foundations: one boundary template
// (a half-open range [lo, hi), loop while lo < hi, hi = mid or lo = mid + 1)
// practised on a plain array, then on counting duplicates, on the first and
// last position, on a bare yes/no array, on an answer range (integer square
// root), on a "compare with the neighbour" rule (peak element) and on a
// feasibility check (ship packages), before the binary search ladder applies
// them. They follow docs/pedagogy.md like any problem, minus the brute force
// (a drill is the plain algorithm itself). The cases check answers, so a
// linear scan is not rejected by the judge; the statements ask for O(log n)
// and the self-check questions explain why.
var BinarySearchDrills = []*Problem{
	{
		ID:              "drill-search-insert-position",
		Number:          35,
		Kind:            KindDrill,
		Order:           1,
		Title:           "Search Insert Position",
		Slug:            "search-insert-position",
		Difficulty:      "Easy",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search", "Foundations"},
		RelatedChapters: []int{6},
		Description: `You are given an array ` + "`" + `nums` + "`" + ` of distinct integers sorted in ascending order, and a number ` + "`" + `target` + "`" + `. Return the index where ` + "`" + `target` + "`" + ` is found. If ` + "`" + `target` + "`" + ` is not in the array, return the index where it would have to be inserted to keep the array sorted (this can be ` + "`" + `len(nums)` + "`" + `, one past the last item).

Do it in O(log n) time. This drill teaches the **boundary template** of binary search that every later drill reuses: the answer is *the first index ` + "`" + `i` + "`" + ` with ` + "`" + `nums[i] >= target` + "`" + `*, and you find it by searching a half-open range ` + "`" + `[lo, hi)` + "`" + `, which starts as ` + "`" + `lo = 0` + "`" + `, ` + "`" + `hi = len(nums)` + "`" + `.`,
		Constraints: []string{
			"1 <= nums.length <= 100000",
			"-10000 <= nums[i] <= 10000",
			"nums contains distinct values sorted in ascending order",
			"-10000 <= target <= 10000",
		},
		Examples: []Example{
			{Input: "nums = [1,3,5,6], target = 5", Output: "2", Explanation: "5 is at index 2."},
			{Input: "nums = [1,3,5,6], target = 2", Output: "1", Explanation: "2 is not in the array. Inserting it at index 1 gives [1,2,3,5,6], which is still sorted."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 3, 5, 6}, "target": 5}, Expected: 2, Note: "example 1: the target is present, so its own index comes back"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 5, 6}, "target": 2}, Expected: 1, Note: "example 2: between 1 and 3, so it would go at index 1"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 5, 6}, "target": 7}, Expected: 4, Hidden: true, Note: "bigger than everything: the insert position is len(nums), one past the last index"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 5, 6}, "target": 0}, Expected: 0, Hidden: true, Note: "smaller than everything: insert at index 0"},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 1}, Expected: 0, Hidden: true, Note: "a single element equal to the target"},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 2}, Expected: 1, Hidden: true, Note: "a single element, target after it"},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 0}, Expected: 0, Hidden: true, Note: "a single element, target before it"},
			{Input: map[string]interface{}{"nums": []int{-10, -4, 0, 3, 8}, "target": -4}, Expected: 1, Hidden: true, Note: "negative numbers, target present"},
			{Input: map[string]interface{}{"nums": []int{-10, -4, 0, 3, 8}, "target": 1}, Expected: 3, Hidden: true, Note: "negative numbers, target missing between 0 and 3"},
			{Input: map[string]interface{}{"nums": []int{2, 4, 6, 8, 10, 12}, "target": 12}, Expected: 5, Hidden: true, Note: "target is the last element: the answer is len(nums) - 1, not len(nums)"},
			{Input: map[string]interface{}{"nums": []int{2, 4, 6, 8, 10, 12}, "target": 2}, Expected: 0, Hidden: true, Note: "target is the first element"},
			{Input: map[string]interface{}{"nums": bsDrillEvens(200000), "target": 399999}, Expected: 200000, Hidden: true, Note: "200000 items, odd target just below the last element: about 18 halvings are needed"},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode: `def searchInsert(nums, target):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Every number before the answer is smaller than target, and every number from the answer on is at least target. Looking only at nums[mid], what does `nums[mid] < target` tell you about where the answer is?"},
			{Level: 2, Type: "insight", Content: "The array reads as 'too small, too small, ..., big enough, big enough'. The answer is the first big-enough index, so one comparison at mid discards half the range: if nums[mid] is too small the answer is after mid, otherwise mid could still be the answer. Search the half-open range [lo, hi) with while lo < hi, and use lo = mid + 1 or hi = mid. This is the boundary template from Chapter 6.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Set the range: lo = 0 and hi = len(nums). hi is a legal answer (insert at the end), so it is included as a candidate but never read.\n2. Loop while lo < hi, and take mid = (lo + hi) // 2.\n3. If nums[mid] < target, then mid and everything left of it is too small: lo = mid + 1.\n4. Otherwise nums[mid] is big enough and could be the answer: hi = mid.\n5. When lo == hi, the range has one candidate left. Return lo."},
			{Level: 4, Type: "skeleton", Content: "def searchInsert(nums, target):\n    lo, hi = 0, len(nums)\n    while lo < ____:\n        mid = (lo + hi) // 2\n        if nums[mid] < target:\n            lo = ____\n        else:\n            hi = ____\n    return ____"},
		},
		Solution: Solution{
			Code: `def searchInsert(nums, target):
    lo, hi = 0, len(nums)
    while lo < hi:
        mid = (lo + hi) // 2
        if nums[mid] < target:
            lo = mid + 1
        else:
            hi = mid
    return lo`,
			Explanation: "The array is a run of too-small values followed by a run of big-enough ones, and the answer is the boundary between them. Each probe at mid says which side of the boundary mid is on, so half of the range is thrown away. This is the first rung of the binary search ladder in its plainest form: Binary Search (704) uses the same halving but stops at an exact match, while this boundary template always ends with lo on the first index that is big enough, and every later drill and Koko Eating Bananas reuse it.",
			KeyInsight:  "You do not search for the target; you search for the boundary where 'nums[i] < target' flips from true to false. Because that flip happens once in a sorted array, one comparison per step is enough to throw half the range away.",
			WhyItWorks:  "Invariant: every index before lo holds a value smaller than target, and every index from hi on is either len(nums) or holds a value at least target, so the answer is always in [lo, hi]. If nums[mid] < target, mid and everything left of it are too small, so lo = mid + 1 keeps the first half of the invariant. Otherwise mid is big enough, so the answer is at most mid and hi = mid keeps the second half. Because mid < hi when lo < hi, both moves shrink the range, so the loop ends. At lo == hi the two sides touch, which makes lo the first big-enough index (or len(nums)). The range halves each time, so about log2(n) steps.",
			CommonMistakes: []string{
				"Mixing templates: the half-open range starts with hi = len(nums), but the loop is written `while lo <= hi`. When lo reaches len(nums), `nums[mid]` is read one past the end and crashes. Pick one template: half-open means `while lo < hi` and `hi = mid`.",
				"Writing `hi = mid - 1` in the else branch. Here mid might be the answer itself, and this throws it away: for [1,3,5,6] and target 5 it returns 1 instead of 2. Closed-range templates use mid - 1; the half-open one keeps mid.",
				"Writing `lo = mid` instead of `lo = mid + 1`. When hi = lo + 1, mid equals lo, nothing changes, and the loop never ends. The probe at mid has already said 'too small', so skip past it.",
				"Returning -1 when the target is missing, as the exact-value version does. This problem always has an answer, because the insert position is where the boundary lies.",
			},
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Alternatives: []AltSolution{
				{Name: "Closed range [lo, hi]", Code: `def searchInsert(nums, target):
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return lo`, TimeComplexity: "O(log n)", SpaceComplexity: "O(1)", Explanation: "The other common template: the range [lo, hi] is closed, so hi starts at len(nums) - 1, the loop runs while lo <= hi, and both moves skip mid. It ends with lo one past the last too-small item, the same answer. It is correct as long as every line belongs to this template; mixing it with the half-open version above is the classic bug."},
				{Name: "bisect_left", Code: `import bisect


def searchInsert(nums, target):
    return bisect.bisect_left(nums, target)`, TimeComplexity: "O(log n)", SpaceComplexity: "O(1)", Explanation: "Python's bisect module has this exact function. Worth using in real code, but writing the template by hand is the point of this drill."},
			},
			Walkthrough: []WalkthroughStep{
				{Title: "Set the search range", Explanation: "The answer is an index from 0 to len(nums), where len(nums) means 'insert at the end'. So the half-open range [lo, hi) starts as the whole array and hi = len(nums) is the last candidate.", CodeSnippet: "lo, hi = 0, len(nums)", LineStart: 2, LineEnd: 2},
				{Title: "Pick the middle", Explanation: "While more than one candidate is left, probe the middle. Integer division rounds down, so mid is always less than hi and the array read is in bounds.", CodeSnippet: "while lo < hi:\nmid = (lo + hi) // 2", LineStart: 3, LineEnd: 4},
				{Title: "Discard half", Explanation: "If nums[mid] is too small, mid and everything before it cannot be the answer. In [1,3,5,6] with target 5, mid = 2 gives 5, which is not too small, so hi = 2 and the second half is gone.", CodeSnippet: "if nums[mid] < target:\nlo = mid + 1\nelse:\nhi = mid", LineStart: 5, LineEnd: 8},
				{Title: "Read the answer", Explanation: "When lo == hi one candidate remains, and it is the first index whose value is at least target.", CodeSnippet: "return lo", LineStart: 9, LineEnd: 9},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the else branch `hi = mid` and not `hi = mid - 1`?", Answer: "If nums[mid] is big enough, mid is a valid answer and might be the first such index, so it must stay in the range. `hi = mid - 1` would drop it. The range [lo, hi) is half-open, which also means hi itself is never read as an element, so hi = mid keeps the invariant that everything from hi on is big enough."},
			{Prompt: "What changes if the array may contain duplicates of the target?", Answer: "Nothing: the loop already returns the first index with nums[i] >= target, so with duplicates it lands on the first copy. That is the idea the next drill uses to count how many copies there are. A version that returns as soon as it sees nums[mid] == target would return whichever copy it hit first."},
		},
	},
	{
		ID:              "drill-count-occurrences",
		Number:          0,
		Kind:            KindDrill,
		Order:           2,
		Title:           "Count Occurrences in a Sorted Array",
		Slug:            "count-occurrences-in-a-sorted-array",
		Difficulty:      "Easy",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search", "Foundations"},
		RelatedChapters: []int{6},
		Description: `You are given an array ` + "`" + `nums` + "`" + ` of integers sorted in ascending order. It may contain duplicates. Return how many times ` + "`" + `target` + "`" + ` appears in it, or 0 if it does not appear. Do it in O(log n) time, so you may not scan the array to count.

All copies of a value sit next to each other in a sorted array, so the count is the width of one block: the index just after the last copy minus the index of the first copy. You already have the first one from the last drill, the **lower bound**: the first index with ` + "`" + `nums[i] >= target` + "`" + `. The other end is the **upper bound**: the first index with ` + "`" + `nums[i] > target` + "`" + `. The upper bound is the same search with a single ` + "`" + `<` + "`" + ` changed to ` + "`" + `<=` + "`" + `.`,
		Constraints: []string{
			"0 <= nums.length <= 100000",
			"-1000000000 <= nums[i], target <= 1000000000",
			"nums is sorted in ascending order and may contain duplicates",
		},
		Examples: []Example{
			{Input: "nums = [1,2,2,2,3,5], target = 2", Output: "3", Explanation: "The lower bound of 2 is index 1 and the upper bound is index 4, so 4 - 1 = 3 copies."},
			{Input: "nums = [1,2,2,2,3,5], target = 4", Output: "0", Explanation: "Both bounds are index 5, the place where 4 would go, so the block has width 0."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 2, 2, 3, 5}, "target": 2}, Expected: 3, Note: "example 1: three copies in the middle"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 2, 2, 3, 5}, "target": 4}, Expected: 0, Note: "example 2: missing, but between two values"},
			{Input: map[string]interface{}{"nums": []int{}, "target": 5}, Expected: 0, Hidden: true, Note: "empty array: nothing to count"},
			{Input: map[string]interface{}{"nums": []int{4}, "target": 4}, Expected: 1, Hidden: true, Note: "a single matching element"},
			{Input: map[string]interface{}{"nums": []int{4}, "target": 5}, Expected: 0, Hidden: true, Note: "a single element that does not match"},
			{Input: map[string]interface{}{"nums": []int{7, 7, 7, 7, 7}, "target": 7}, Expected: 5, Hidden: true, Note: "every element equals the target: from index 0 to the end"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 2, 3, 3, 3}, "target": 1}, Expected: 2, Hidden: true, Note: "the copies are at the very start"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 3, 3}, "target": 3}, Expected: 3, Hidden: true, Note: "the copies are at the very end: first_greater must be able to return len(nums)"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 5, 6}, "target": 0}, Expected: 0, Hidden: true, Note: "smaller than everything"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 5, 6}, "target": 9}, Expected: 0, Hidden: true, Note: "bigger than everything"},
			{Input: map[string]interface{}{"nums": []int{-3, -3, -1, 0, 0, 0, 0, 2}, "target": 0}, Expected: 4, Hidden: true, Note: "negatives and four copies of zero"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 3, 5}, "target": 3}, Expected: 2, Hidden: true, Note: "two copies: an answer of 1 means the code counted 'found' instead of how many"},
			{Input: map[string]interface{}{"nums": bsDrillRuns(60000, 80000, 60000), "target": 7}, Expected: 80000, Hidden: true, Note: "200000 items with 80000 copies in the middle: both bounds are needed"},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode: `def countOccurrences(nums, target):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In a sorted array, equal values form one block. If you could find where the block starts and where it ends, how would you get its length without visiting each element?"},
			{Level: 2, Type: "insight", Content: "Run the boundary template twice. The lower bound is the first index where nums[i] >= target, found with `nums[mid] < target` meaning 'go right'. The upper bound is the first index where nums[i] > target, found with `nums[mid] <= target` meaning 'go right'. The count is upper - lower, and it is 0 automatically when the target is missing because both bounds land on the same index. See the boundary template in Chapter 6.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Write the lower bound: lo, hi = 0, len(nums); while lo < hi: if nums[mid] < target, lo = mid + 1, else hi = mid.\n2. Write the upper bound the same way, but the 'go right' test is nums[mid] <= target, so it skips past all copies.\n3. Return the upper bound minus the lower bound.\n4. Check an empty array: both loops are skipped and return 0 - 0."},
			{Level: 4, Type: "skeleton", Content: "def countOccurrences(nums, target):\n    def first_at_least():\n        lo, hi = 0, len(nums)\n        while lo < hi:\n            mid = (lo + hi) // 2\n            if nums[mid] < target:\n                lo = mid + 1\n            else:\n                hi = mid\n        return lo\n\n    def first_greater():\n        lo, hi = 0, len(nums)\n        while lo < hi:\n            mid = (lo + hi) // 2\n            if nums[mid] ____ target:\n                lo = mid + 1\n            else:\n                hi = mid\n        return lo\n\n    return ____ - ____"},
		},
		Solution: Solution{
			Code: `def countOccurrences(nums, target):
    def first_at_least():
        lo, hi = 0, len(nums)
        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] < target:
                lo = mid + 1
            else:
                hi = mid
        return lo

    def first_greater():
        lo, hi = 0, len(nums)
        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] <= target:
                lo = mid + 1
            else:
                hi = mid
        return lo

    return first_greater() - first_at_least()`,
			Explanation: "With duplicates there is no single 'the' index of the target, so instead of finding one copy the code finds the two edges of the block of copies and subtracts. It builds directly on Search Insert Position: the first helper is that drill unchanged, and the second differs by one comparison. This pair of bounds is also what the next drill (first and last position) and Time Based Key-Value Store use.",
			KeyInsight:  "A comparison decides which boundary you find. Moving right while the value is `< target` stops at the first copy; moving right while the value is `<= target` stops just after the last copy. The same template gives both edges.",
			WhyItWorks:  "Each helper keeps the invariant of the previous drill: everything before lo answers 'go right' and everything from hi on answers 'stop', for its own test. The test is monotone because the array is sorted: `nums[i] < target` is true for a prefix, and `nums[i] <= target` is true for a longer prefix. So first_at_least returns the length of the prefix of values below target, and first_greater returns the length of the prefix of values at most target. Their difference is the number of values equal to target. If the target is missing the two prefixes are the same, so the difference is 0. Each search halves its range, so O(log n) in total.",
			CommonMistakes: []string{
				"Using `<` in both helpers, so both searches return the same index and the count is always 0. The upper bound needs `<=` so that it moves past every copy.",
				"Binary searching for any copy and counting 1 (or returning mid) when nums[mid] == target. That finds one copy, not the block's two edges, so [1,3,3,5] gives 1 instead of 2. Duplicates make 'found' not enough: you need the first and the one after the last.",
				"Starting the upper-bound search with `hi = len(nums) - 1`. If the last copy is the final element, the answer len(nums) can no longer be reached and the count comes out one short. In the half-open template, hi starts at len(nums).",
			},
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Alternatives: []AltSolution{
				{Name: "bisect", Code: `import bisect


def countOccurrences(nums, target):
    return bisect.bisect_right(nums, target) - bisect.bisect_left(nums, target)`, TimeComplexity: "O(log n)", SpaceComplexity: "O(1)", Explanation: "bisect_left is the lower bound and bisect_right is the upper bound, so the standard library already has both edges. Writing them by hand is this drill's point."},
			},
			Walkthrough: []WalkthroughStep{
				{Title: "Find the first edge", Explanation: "The lower bound moves right while the value is smaller than target, so it stops on the first copy (or where the target would go). On [1,2,2,2,3,5] with target 2 it stops at index 1.", CodeSnippet: "if nums[mid] < target:\nlo = mid + 1", LineStart: 6, LineEnd: 7},
				{Title: "Find the edge after the last copy", Explanation: "The upper bound moves right while the value is at most target. That also steps over every copy, so it stops just after the last one. On the same array it stops at index 4.", CodeSnippet: "if nums[mid] <= target:\nlo = mid + 1", LineStart: 16, LineEnd: 17},
				{Title: "Keep the same template in both", Explanation: "Both helpers start with hi = len(nums), loop while lo < hi, and keep the middle on the else branch with hi = mid. Only the comparison differs.", CodeSnippet: "while lo < hi:\nmid = (lo + hi) // 2", LineStart: 4, LineEnd: 5},
				{Title: "Subtract the edges", Explanation: "The width of the block is the upper edge minus the lower edge: 4 - 1 = 3. For a missing target both are the same index and the result is 0.", CodeSnippet: "return first_greater() - first_at_least()", LineStart: 22, LineEnd: 22},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does `<=` find the position after the last copy?", Answer: "The test `nums[mid] <= target` is true for every copy as well as everything smaller, so the search keeps moving right until it reaches a value larger than target. The first index holding such a value is directly after the last copy. With `<` the copies make the test false, so the search stops at the first one."},
			{Prompt: "How would you find the upper bound if the values were not integers (for example, floats)?", Answer: "The same code works, because it only uses comparisons. The shortcut of using the lower bound of `target + 1` is the only version that depends on integers, so writing `<=` is the more general habit."},
		},
	},
	{
		ID:              "drill-first-and-last-position",
		Number:          34,
		Kind:            KindDrill,
		Order:           3,
		Title:           "Find First and Last Position of Element in Sorted Array",
		Slug:            "find-first-and-last-position-of-element-in-sorted-array",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search", "Foundations"},
		RelatedChapters: []int{6},
		Description: `You are given an array ` + "`" + `nums` + "`" + ` of integers sorted in ascending order, which may contain duplicates, and a number ` + "`" + `target` + "`" + `. Return a list ` + "`" + `[first, last]` + "`" + ` holding the index of the first and the index of the last occurrence of ` + "`" + `target` + "`" + `. If ` + "`" + `target` + "`" + ` does not appear, return ` + "`" + `[-1, -1]` + "`" + `. Do it in O(log n) time.

The last drill counted copies with two bounds. Here you must also decide whether the target exists at all: the lower bound only gives *where the target would go*, which is not the same as *where it is*. After a search lands on an index, check that it is inside the array and that ` + "`" + `nums[index] == target` + "`" + ` before trusting it.`,
		Constraints: []string{
			"0 <= nums.length <= 100000",
			"-1000000000 <= nums[i], target <= 1000000000",
			"nums is sorted in ascending order and may contain duplicates",
		},
		Examples: []Example{
			{Input: "nums = [5,7,7,8,8,10], target = 8", Output: "[3,4]", Explanation: "8 first appears at index 3 and last appears at index 4."},
			{Input: "nums = [5,7,7,8,8,10], target = 6", Output: "[-1,-1]", Explanation: "6 is not in the array. The lower bound is index 1, which holds 7, not 6, so the target is missing."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{5, 7, 7, 8, 8, 10}, "target": 8}, Expected: []int{3, 4}, Note: "example 1: the copies of 8 are at indices 3 and 4"},
			{Input: map[string]interface{}{"nums": []int{5, 7, 7, 8, 8, 10}, "target": 6}, Expected: []int{-1, -1}, Note: "example 2: 6 would sit between 5 and 7 but is not there"},
			{Input: map[string]interface{}{"nums": []int{}, "target": 0}, Expected: []int{-1, -1}, Hidden: true, Note: "empty array"},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 1}, Expected: []int{0, 0}, Hidden: true, Note: "a single element that matches: first and last are both 0"},
			{Input: map[string]interface{}{"nums": []int{1}, "target": 2}, Expected: []int{-1, -1}, Hidden: true, Note: "a single element, target missing above it"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 2, 2}, "target": 2}, Expected: []int{0, 3}, Hidden: true, Note: "every element matches: [0, 3]"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 3, 3}, "target": 3}, Expected: []int{2, 4}, Hidden: true, Note: "the copies run to the end of the array"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 2, 3}, "target": 1}, Expected: []int{0, 2}, Hidden: true, Note: "the copies start at index 0"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 5, 7}, "target": 9}, Expected: []int{-1, -1}, Hidden: true, Note: "target above everything: the lower bound is len(nums), so nums[first] must not be read"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 5, 7}, "target": 0}, Expected: []int{-1, -1}, Hidden: true, Note: "target below everything: the lower bound is 0 but nums[0] is not the target"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 3, 5}, "target": 4}, Expected: []int{-1, -1}, Hidden: true, Note: "missing between two values: the lower bound points at a larger value"},
			{Input: map[string]interface{}{"nums": []int{-5, -5, -2, 0, 3}, "target": -5}, Expected: []int{0, 1}, Hidden: true, Note: "negative duplicates at the start"},
			{Input: map[string]interface{}{"nums": bsDrillRuns(60000, 80000, 60000), "target": 7}, Expected: []int{60000, 139999}, Hidden: true, Note: "200000 items with the target in the middle: [60000, 139999]"},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode: `def searchRange(nums, target):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you search for the lower bound of target, you get an index. What two checks tell you whether that index really holds the target, and not just a place it could be inserted?"},
			{Level: 2, Type: "insight", Content: "The lower bound of target is the first index with nums[i] >= target. If that index is len(nums), or holds a value other than target, the target is missing and the answer is [-1, -1]. Otherwise it is the first position. The last position is one before the lower bound of target + 1, because the next larger integer starts right after the block. Chapter 6 calls this the boundary template.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Write one helper, first_at_least(value), with the half-open template and `nums[mid] < value` meaning go right.\n2. first = first_at_least(target).\n3. Check for absence: first == len(nums) (test this before reading nums[first]) or nums[first] != target means return [-1, -1].\n4. last = first_at_least(target + 1) - 1, the index just before the next larger value begins.\n5. Return [first, last]."},
			{Level: 4, Type: "skeleton", Content: "def searchRange(nums, target):\n    def first_at_least(value):\n        lo, hi = 0, len(nums)\n        while lo < hi:\n            mid = (lo + hi) // 2\n            if nums[mid] < value:\n                lo = mid + 1\n            else:\n                hi = mid\n        return lo\n\n    first = first_at_least(target)\n    if first == ____ or nums[first] != ____:\n        return [-1, -1]\n    return [first, first_at_least(____) - 1]"},
		},
		Solution: Solution{
			Code: `def searchRange(nums, target):
    def first_at_least(value):
        lo, hi = 0, len(nums)
        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] < value:
                lo = mid + 1
            else:
                hi = mid
        return lo

    first = first_at_least(target)
    if first == len(nums) or nums[first] != target:
        return [-1, -1]
    return [first, first_at_least(target + 1) - 1]`,
			Explanation: "This combines the two earlier drills: the lower bound from Search Insert Position gives the first copy, and the idea behind Count Occurrences (the block ends where the next larger value begins) gives the last. The new step is existence: a bound is only a position, so you check it before answering. Search in Rotated Sorted Array and Time Based Key-Value Store later lean on the same 'search, then verify what you landed on' habit.",
			KeyInsight:  "Searching for the target is really two boundary searches. The first copy is the lower bound of target, and the last copy is one before the lower bound of target + 1. A boundary only says where a value belongs, so verify that the value is really there.",
			WhyItWorks:  "Invariant of the helper: everything before lo is smaller than value and everything from hi on is at least value, so it returns the first index holding a value >= value (or len(nums)). If target is in the array, its first copy is the first index with a value >= target, so first is right and nums[first] == target holds. If the target is missing, that index either is len(nums) or holds a larger value, and the check returns [-1, -1]. The values greater than target start right after the last copy, and for integers that is the first value >= target + 1, so subtracting one gives the last copy. Two binary searches, O(log n).",
			CommonMistakes: []string{
				"Reading `nums[first]` before checking `first == len(nums)`. When the target is larger than everything, first is len(nums) and the read raises IndexError. Check the bound first, as in `first == len(nums) or nums[first] != target`.",
				"Skipping the existence check and returning [first, first_at_least(target + 1) - 1] for a missing target. For [1,3,5,7] and target 4 that gives [2, 1], a nonsense pair, instead of [-1, -1].",
				"Binary searching until nums[mid] == target and returning that mid for both ends (or expanding left and right one step at a time). With duplicates, mid can be any copy, and walking outwards is O(n) when most of the array equals the target.",
				"Forgetting the `- 1` in the last position, which returns the first index after the block. The lower bound of target + 1 is already one past the last copy.",
			},
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Alternatives: []AltSolution{
				{Name: "bisect", Code: `import bisect


def searchRange(nums, target):
    first = bisect.bisect_left(nums, target)
    if first == len(nums) or nums[first] != target:
        return [-1, -1]
    return [first, bisect.bisect_right(nums, target) - 1]`, TimeComplexity: "O(log n)", SpaceComplexity: "O(1)", Explanation: "The standard library has both bounds. The existence check is still needed, and it is the part this drill is about."},
			},
			Walkthrough: []WalkthroughStep{
				{Title: "Write the boundary search once", Explanation: "One helper finds the first index whose value is at least `value`, using the half-open template from the earlier drills. Calling it twice with different values gives both edges.", CodeSnippet: "def first_at_least(value):\nlo, hi = 0, len(nums)\nwhile lo < hi:", LineStart: 2, LineEnd: 4},
				{Title: "Find the first copy", Explanation: "The lower bound of target is the only place the first copy can be. On [5,7,7,8,8,10] with target 8 it is index 3.", CodeSnippet: "first = first_at_least(target)", LineStart: 12, LineEnd: 12},
				{Title: "Check the target is really there", Explanation: "A bound is only a position. If it is past the end or holds a different value (target 6 gives index 1, which holds 7), the target is missing. Test the index before reading nums[first].", CodeSnippet: "if first == len(nums) or nums[first] != target:\nreturn [-1, -1]", LineStart: 13, LineEnd: 14},
				{Title: "Find the last copy", Explanation: "The next larger integer starts right after the block, so one before the lower bound of target + 1 is the last copy: 5 - 1 = 4.", CodeSnippet: "return [first, first_at_least(target + 1) - 1]", LineStart: 15, LineEnd: 15},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the check `first == len(nums)` come before `nums[first] != target`?", Answer: "Python evaluates `or` from left to right and stops at the first true part. When first equals len(nums) the first part is true, so the read `nums[first]` never runs. In the other order, a target bigger than everything would index past the end and raise an error."},
			{Prompt: "What if the array were sorted in descending order?", Answer: "Flip the comparison: go right while nums[mid] is greater than the value, so the 'too early' prefix is the large values. The structure and the half-open template stay identical. Only the meaning of 'before' changes, which is a good habit: write down the yes/no test first, then the loop."},
		},
	},
	{
		ID:              "drill-first-true",
		Number:          0,
		Kind:            KindDrill,
		Order:           4,
		Title:           "First True in a Sorted Yes/No Array",
		Slug:            "first-true-in-a-sorted-yes-no-array",
		Difficulty:      "Easy",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search", "Foundations"},
		RelatedChapters: []int{6},
		Description: `You are given a list ` + "`" + `flags` + "`" + ` of booleans that is **monotone**: some number of ` + "`" + `False` + "`" + ` values come first, followed by ` + "`" + `True` + "`" + ` values only (either part may be empty). Return the index of the first ` + "`" + `True` + "`" + `. If there is no ` + "`" + `True` + "`" + `, return ` + "`" + `len(flags)` + "`" + `.

This is the bare form of the boundary template: the "comparison" is just reading ` + "`" + `flags[mid]` + "`" + `. Every later binary search is this same loop with a different question hidden inside ` + "`" + `flags[mid]` + "`" + `: *is ` + "`" + `nums[mid] >= target` + "`" + `?*, *is this speed fast enough?*, *is ` + "`" + `mid * mid > x` + "`" + `?* Once you can answer a yes/no question that flips only once, you can binary search for the flip.`,
		Constraints: []string{
			"0 <= flags.length <= 100000",
			"flags is False for every index below some boundary and True from there on",
		},
		Examples: []Example{
			{Input: "flags = [False, False, False, True, True, True]", Output: "3", Explanation: "The first True is at index 3."},
			{Input: "flags = [False, False, False]", Output: "3", Explanation: "No item is True, so the answer is len(flags) = 3."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"flags": []bool{false, false, false, true, true, true}}, Expected: 3, Note: "example 1: the first True is at index 3"},
			{Input: map[string]interface{}{"flags": []bool{false, false, false}}, Expected: 3, Note: "example 2: no True anywhere, so the answer is len(flags)"},
			{Input: map[string]interface{}{"flags": []bool{true, true, true, true}}, Expected: 0, Hidden: true, Note: "all True: the first one is at index 0"},
			{Input: map[string]interface{}{"flags": []bool{true}}, Expected: 0, Hidden: true, Note: "a single True"},
			{Input: map[string]interface{}{"flags": []bool{false}}, Expected: 1, Hidden: true, Note: "a single False: nothing is True"},
			{Input: map[string]interface{}{"flags": []int{}}, Expected: 0, Hidden: true, Note: "empty array: the answer is 0 = len(flags)"},
			{Input: map[string]interface{}{"flags": []bool{false, true}}, Expected: 1, Hidden: true, Note: "the smallest array that has a boundary in the middle"},
			{Input: map[string]interface{}{"flags": []bool{false, false, false, false, true}}, Expected: 4, Hidden: true, Note: "only the last item is True"},
			{Input: map[string]interface{}{"flags": []bool{false, true, true, true, true, true, true}}, Expected: 1, Hidden: true, Note: "only the first item is False"},
			{Input: map[string]interface{}{"flags": []bool{false, false, false, false, false, false, false, true, true, true, true, true, true, true, true}}, Expected: 7, Hidden: true, Note: "an odd boundary position with an odd length"},
			{Input: map[string]interface{}{"flags": []bool{false, false, false, false, false, false, false, false, false, false}}, Expected: 10, Hidden: true, Note: "ten Falses: lo must end at len(flags)"},
			{Input: map[string]interface{}{"flags": bsDrillFlags(200000, 123457)}, Expected: 123457, Hidden: true, Note: "200000 flags with the first True at index 123457"},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode: `def firstTrue(flags):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Because the list looks like 'False, False, ..., True, True', one look at flags[mid] splits the range in two. If flags[mid] is True, which side holds the first True, and could mid itself be it?"},
			{Level: 2, Type: "insight", Content: "If flags[mid] is True, the first True is at mid or to its left, so keep mid: hi = mid. If it is False, then mid and everything before it are False, so lo = mid + 1. The range is half-open, [lo, hi), with hi = len(flags) as the 'no True' answer. This is the boundary template from Chapter 6, with the yes/no question made explicit.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Set lo = 0 and hi = len(flags). hi is a legal answer when everything is False.\n2. While lo < hi, take mid = (lo + hi) // 2.\n3. If flags[mid] is True, the answer is at most mid: hi = mid.\n4. Otherwise flags[mid] is False, so the answer is after mid: lo = mid + 1.\n5. Return lo."},
			{Level: 4, Type: "skeleton", Content: "def firstTrue(flags):\n    lo, hi = 0, len(flags)\n    while lo < hi:\n        mid = (lo + hi) // 2\n        if flags[mid]:\n            hi = ____\n        else:\n            lo = ____\n    return ____"},
		},
		Solution: Solution{
			Code: `def firstTrue(flags):
    lo, hi = 0, len(flags)
    while lo < hi:
        mid = (lo + hi) // 2
        if flags[mid]:
            hi = mid
        else:
            lo = mid + 1
    return lo`,
			Explanation: "This is the boundary template with nothing else attached: a monotone yes/no sequence and a search for the point where it flips. It generalises Search Insert Position, where the yes/no question was `nums[mid] >= target`, and it is the shape behind Koko Eating Bananas (875), Find Minimum in Rotated Sorted Array (153) and the next two drills.",
			KeyInsight:  "A binary search needs only one thing: a question whose answers are all 'no' up to some point and all 'yes' after it. Find that question, and the code is always the same eight lines.",
			WhyItWorks:  "Invariant: every index below lo is False, and hi is either len(flags) or an index holding True, so the first True (or len(flags)) is in [lo, hi]. A True at mid means the first True is at most mid, so hi = mid keeps the invariant. A False at mid means mid and everything before it are False (monotone), so lo = mid + 1 keeps it. The range shrinks every step, and when lo == hi it names the boundary. That takes about log2(n) steps, because each step discards half of the range.",
			CommonMistakes: []string{
				"Returning mid as soon as flags[mid] is True. That index is True but not necessarily the first True: [False, True, True] could return 2. The first True is at mid or to the left, so keep searching with hi = mid.",
				"Writing `hi = mid - 1` when flags[mid] is True. Mid might be the first True, and this discards it, so the code returns the position after it.",
				"Setting `hi = len(flags) - 1`, which makes the all-False case unreachable: the loop ends at the last index and returns len(flags) - 1, but the correct answer is len(flags). The upper end of a half-open range is a candidate that is never read.",
				"Using `lo = mid` when flags[mid] is False. When hi = lo + 1, mid equals lo and the range never shrinks, so the loop runs forever. A False at mid has been ruled out, so move past it with mid + 1.",
			},
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Alternatives: []AltSolution{
				{Name: "bisect with a key", Code: `import bisect


def firstTrue(flags):
    return bisect.bisect_left(flags, True)`, TimeComplexity: "O(log n)", SpaceComplexity: "O(1)", Explanation: "Python compares False < True, so a monotone list of booleans is already sorted and bisect_left finds the first True. It is a neat trick; the drill's point is to write the loop yourself."},
			},
			Walkthrough: []WalkthroughStep{
				{Title: "Set the search range", Explanation: "The answer is an index in 0..len(flags), where len(flags) means no True exists. So hi starts one past the last index.", CodeSnippet: "lo, hi = 0, len(flags)", LineStart: 2, LineEnd: 2},
				{Title: "Pick the middle", Explanation: "Probe the middle of what is left. Because lo < hi, mid is always a real index.", CodeSnippet: "while lo < hi:\nmid = (lo + hi) // 2", LineStart: 3, LineEnd: 4},
				{Title: "Keep the half that holds the boundary", Explanation: "True at mid means the first True is at mid or earlier, so mid stays in the range (hi = mid). False at mid means it is later, so mid is dropped (lo = mid + 1).", CodeSnippet: "if flags[mid]:\nhi = mid\nelse:\nlo = mid + 1", LineStart: 5, LineEnd: 8},
				{Title: "Read the answer", Explanation: "When lo == hi one candidate is left, and it is the boundary. On [False, False, False, True, True, True] the range goes 0..6, then 0..3, then 2..3, then 3..3.", CodeSnippet: "return lo", LineStart: 9, LineEnd: 9},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "What does the code return for an all-False list, and why is that the right answer?", Answer: "It returns len(flags). Every probe says False, so lo keeps moving right until it reaches hi, which started at len(flags). That is the 'no True' answer the statement asks for, and the reason hi must start at len(flags), not len(flags) - 1."},
			{Prompt: "How would you find the *last* True in a list of True values followed by False values?", Answer: "Flip the yes/no question: the first False is one past the last True, so search for the first False the same way and subtract 1. The loop never changes. Only the question you ask changes, which is the transferable skill."},
		},
	},
	{
		ID:              "drill-integer-square-root",
		Number:          69,
		Kind:            KindDrill,
		Order:           5,
		Title:           "Sqrt(x)",
		Slug:            "sqrtx",
		Difficulty:      "Easy",
		Category:        "binary-search",
		Tags:            []string{"Math", "Binary Search", "Foundations"},
		RelatedChapters: []int{6},
		Description: `Given a non-negative integer ` + "`" + `x` + "`" + `, return the integer square root of ` + "`" + `x` + "`" + `: the largest integer ` + "`" + `r` + "`" + ` with ` + "`" + `r * r <= x` + "`" + `, which is the square root rounded down. Do not use ` + "`" + `math.sqrt` + "`" + `, ` + "`" + `math.isqrt` + "`" + `, ` + "`" + `** 0.5` + "`" + ` or any built-in power function.

Here the thing you search is not an array. It is the set of possible *answers*, the integers 0, 1, 2, ..., and the yes/no question is ` + "`" + `mid * mid > x` + "`" + `. That question is ` + "`" + `False` + "`" + ` for small numbers and ` + "`" + `True` + "`" + ` from some point on, which is exactly the monotone array of the last drill, but one that you never build: you compute ` + "`" + `flags[mid]` + "`" + ` on demand.`,
		Constraints: []string{
			"0 <= x <= 2147483647",
		},
		Examples: []Example{
			{Input: "x = 4", Output: "2", Explanation: "2 * 2 = 4."},
			{Input: "x = 8", Output: "2", Explanation: "The square root of 8 is about 2.83. Rounded down it is 2, since 2 * 2 = 4 <= 8 but 3 * 3 = 9 > 8."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"x": 4}, Expected: 2, Note: "example 1: a perfect square"},
			{Input: map[string]interface{}{"x": 8}, Expected: 2, Note: "example 2: between 2 * 2 = 4 and 3 * 3 = 9, so the answer rounds down to 2"},
			{Input: map[string]interface{}{"x": 0}, Expected: 0, Hidden: true, Note: "zero: the answer is 0"},
			{Input: map[string]interface{}{"x": 1}, Expected: 1, Hidden: true, Note: "one: the smallest positive input, answer 1 (a search that stops below x fails here)"},
			{Input: map[string]interface{}{"x": 2}, Expected: 1, Hidden: true, Note: "just above 1: still 1"},
			{Input: map[string]interface{}{"x": 3}, Expected: 1, Hidden: true, Note: "the last number before 4: answer 1"},
			{Input: map[string]interface{}{"x": 15}, Expected: 3, Hidden: true, Note: "one below a perfect square: 3 * 3 = 9 <= 15 < 16"},
			{Input: map[string]interface{}{"x": 16}, Expected: 4, Hidden: true, Note: "a perfect square above the tiny ones"},
			{Input: map[string]interface{}{"x": 17}, Expected: 4, Hidden: true, Note: "one above a perfect square"},
			{Input: map[string]interface{}{"x": 99980001}, Expected: 9999, Hidden: true, Note: "9999 * 9999 exactly"},
			{Input: map[string]interface{}{"x": 2147395599}, Expected: 46339, Hidden: true, Note: "one below 46340 * 46340: the answer is 46339"},
			{Input: map[string]interface{}{"x": 2147395600}, Expected: 46340, Hidden: true, Note: "exactly 46340 * 46340: large squares must be hit exactly"},
			{Input: map[string]interface{}{"x": 2147483647}, Expected: 46340, Hidden: true, Note: "the maximum input, 2^31 - 1: the answer is 46340"},
		},
		TimeComplexity:  "O(log x)",
		SpaceComplexity: "O(1)",
		StarterCode: `def mySqrt(x):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Consider the numbers 0, 1, 2, 3, ... and ask of each: is mid * mid larger than x? What does that sequence of answers look like?"},
			{Level: 2, Type: "insight", Content: "It looks like 'no, no, ..., no, yes, yes, ...': once a number's square is above x, every bigger number's square is too. The integer square root is the last 'no', which is one before the first 'yes'. Find the first yes with the boundary template from Chapter 6, searching the range [0, x + 1) (x + 1 is a safe 'yes'), and return lo - 1.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Set the answer range: lo = 0, hi = x + 1. The number x + 1 always says yes, because (x + 1) * (x + 1) > x.\n2. While lo < hi, take mid = (lo + hi) // 2.\n3. If mid * mid > x, mid is a yes, so the first yes is at most mid: hi = mid.\n4. Otherwise mid is a no, so move past it: lo = mid + 1.\n5. lo is now the first number whose square is above x. The answer is one before it: return lo - 1."},
			{Level: 4, Type: "skeleton", Content: "def mySqrt(x):\n    lo, hi = 0, x + 1\n    while lo < hi:\n        mid = (lo + hi) // 2\n        if mid * mid ____ x:\n            hi = mid\n        else:\n            lo = ____\n    return lo ____ 1"},
		},
		Solution: Solution{
			Code: `def mySqrt(x):
    lo, hi = 0, x + 1
    while lo < hi:
        mid = (lo + hi) // 2
        if mid * mid > x:
            hi = mid
        else:
            lo = mid + 1
    return lo - 1`,
			Explanation: "Instead of searching an array, this searches the answer itself: the integers from 0 up, with the question 'is the square above x?' standing in for flags[mid]. It reuses the first-True drill exactly, then subtracts one, because the integer square root is the last no. This is the idea behind Koko Eating Bananas (875), where the thing searched is an eating speed.",
			KeyInsight:  "When the answer is a number and you can test any candidate, search the candidates. If 'candidate is too big' is False and then True, with a single flip, the first True is a boundary you can find in O(log range) tests instead of trying them all.",
			WhyItWorks:  "Invariant: every number below lo has a square at most x, and hi is either x + 1 or a number whose square is above x, so the first number with a square above x is in [lo, hi]. Squares grow with the number, so 'mid * mid > x' flips only once. A yes at mid means the boundary is at most mid (hi = mid); a no means mid and everything below it fit (lo = mid + 1). When lo == hi it is the first number whose square is above x, and the number before it is the last one that fits, which is the largest r with r * r <= x. x + 1 is always a yes because (x + 1)^2 = x^2 + 2x + 1 > x, so the first yes exists. The range is x + 1 wide, so O(log x) steps.",
			CommonMistakes: []string{
				"Returning `lo` instead of `lo - 1`. The loop finds the first number whose square is too big, which is one more than the answer, so the code returns 3 for x = 8.",
				"Starting with `hi = x`. For x = 1 the range [0, 1) holds only 0, so the code ends with lo = 1 and returns 0 instead of 1. The first yes may be x + 1, so the range must reach it.",
				"Searching for an exact square, `if mid * mid == x: return mid`, with nothing else for the non-square case. A non-square like 8 never matches, so the code needs the 'last number that fits' rule instead of equality.",
				"Mixing templates: `while lo <= hi` with `hi = mid`. When lo == hi the range never shrinks, and the loop does not end. A closed range needs `hi = mid - 1`, a half-open one needs `while lo < hi`.",
			},
			TimeComplexity:  "O(log x)",
			SpaceComplexity: "O(1)",
			Alternatives: []AltSolution{
				{Name: "math.isqrt", Code: `import math


def mySqrt(x):
    return math.isqrt(x)`, TimeComplexity: "O(log x)", SpaceComplexity: "O(1)", Explanation: "Python 3.8+ ships an exact integer square root. The statement forbids it so you practise the search, but in real code this is the right call; floating-point `int(x ** 0.5)` is not, since it can be off by one for large inputs."},
			},
			Walkthrough: []WalkthroughStep{
				{Title: "Define the answer range", Explanation: "The candidates are the integers 0 to x, plus x + 1 as a guaranteed 'too big'. The half-open range [0, x + 1) holds the first number whose square is above x.", CodeSnippet: "lo, hi = 0, x + 1", LineStart: 2, LineEnd: 2},
				{Title: "Pick the middle", Explanation: "Probe a candidate in the middle of the remaining range. For x = 8 the first probe is mid = 4.", CodeSnippet: "while lo < hi:\nmid = (lo + hi) // 2", LineStart: 3, LineEnd: 4},
				{Title: "Discard half", Explanation: "If mid * mid is above x (4 * 4 = 16 > 8), then mid and everything above are too big, so hi = mid. Otherwise mid fits and everything below it fits too, so lo = mid + 1.", CodeSnippet: "if mid * mid > x:\nhi = mid\nelse:\nlo = mid + 1", LineStart: 5, LineEnd: 8},
				{Title: "Step back one", Explanation: "lo is the first number that is too big (3 for x = 8, because 3 * 3 = 9 > 8). The largest one that still fits is the number before it.", CodeSnippet: "return lo - 1", LineStart: 9, LineEnd: 9},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the answer `lo - 1`, and when would you not subtract one?", Answer: "The search finds the first 'yes' to a 'too big' question, so the last 'no' (the largest number that fits) is the one before it. If you phrase the question the other way, for example 'is mid * mid <= x?' with True first, you would be looking for the last True, which needs a different loop shape. Choose the question that flips from False to True and the template does not change."},
			{Prompt: "Why is `int(x ** 0.5)` risky for large x, even though it is O(1)?", Answer: "Floats have 53 bits of precision, so for big integers the square root can be a hair below or above the true value and `int` rounds the wrong way, giving an answer that is off by one. An integer binary search uses exact arithmetic, so it is never off."},
		},
	},
	{
		ID:              "drill-find-peak-element",
		Number:          162,
		Kind:            KindDrill,
		Order:           6,
		Title:           "Find Peak Element",
		Slug:            "find-peak-element",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search", "Foundations"},
		RelatedChapters: []int{6},
		Description: `A **peak** is an element strictly greater than its neighbours. You are given an array ` + "`" + `nums` + "`" + ` in which any two neighbouring values are different (` + "`" + `nums[i] != nums[i + 1]` + "`" + `). Treat the two ends as if they have ` + "`" + `-infinity` + "`" + ` beyond them, so ` + "`" + `nums[-1] = nums[len(nums)] = -infinity` + "`" + ` for the comparison only: the first element is a peak if it is larger than ` + "`" + `nums[1]` + "`" + `, and the last is a peak if it is larger than ` + "`" + `nums[len(nums) - 2]` + "`" + `.

Return the **index of any peak**. There may be several, and any one is accepted. The array is not sorted, but you must still run in O(log n) time. The key is that you do not need to find the *highest* peak: comparing ` + "`" + `nums[mid]` + "`" + ` with its right neighbour is enough to tell you which half is *guaranteed* to hold one.`,
		Constraints: []string{
			"1 <= nums.length <= 1000",
			"-2147483648 <= nums[i] <= 2147483647",
			"nums[i] != nums[i + 1] for every valid i",
		},
		Examples: []Example{
			{Input: "nums = [1,2,3,1]", Output: "2", Explanation: "nums[2] = 3 is larger than both neighbours (2 and 1)."},
			{Input: "nums = [1,2,1,3,5,6,4]", Output: "5", Explanation: "Index 1 (value 2) and index 5 (value 6) are both peaks. Returning either is correct."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 1}}, Expected: 2, Note: "example 1: a single peak at index 2"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 1, 3, 5, 6, 4}}, Expected: 5, Note: "example 2: two peaks (index 1 and index 5): either is accepted"},
			{Input: map[string]interface{}{"nums": []int{5}}, Expected: 0, Hidden: true, Note: "a single element is a peak"},
			{Input: map[string]interface{}{"nums": []int{1, 2}}, Expected: 1, Hidden: true, Note: "increasing pair: the peak is the last element, whose right neighbour counts as -infinity"},
			{Input: map[string]interface{}{"nums": []int{2, 1}}, Expected: 0, Hidden: true, Note: "decreasing pair: the peak is the first element"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5}}, Expected: 4, Hidden: true, Note: "strictly increasing: the only peak is the last index, so a search that stops in the middle is wrong"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 2, 1, 0}}, Expected: 1, Hidden: true, Note: "the only peak is strictly inside the array, and neither end is a peak"},
			{Input: map[string]interface{}{"nums": []int{5, 4, 3, 2, 1}}, Expected: 0, Hidden: true, Note: "strictly decreasing: the only peak is index 0"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 2, 4, 2, 5, 1, 0}}, Expected: 1, Hidden: true, Note: "several peaks (3, 4, 5): whichever the search lands on is accepted"},
			{Input: map[string]interface{}{"nums": []int{-9, -5, -7, -8}}, Expected: 1, Hidden: true, Note: "negative numbers with the peak at index 1"},
			{Input: map[string]interface{}{"nums": []int{3, 1, 2, 1, 2, 1, 3}}, Expected: 6, Hidden: true, Note: "peaks at both ends and in the middle"},
			{Input: map[string]interface{}{"nums": []int{1, 5, 9, 8, 7, 6, 5, 4, 3, 2}}, Expected: 2, Hidden: true, Note: "the maximum is near the left, the right side is a long fall"},
			{Input: map[string]interface{}{"nums": bsDrillPeaks(1000, 612)}, Expected: 612, Hidden: true, Note: "1000 items rising to index 612 and falling: the only peak"},
		},
		TimeComplexity:  "O(log n)",
		SpaceComplexity: "O(1)",
		StarterCode: `def findPeakElement(nums):
    # Write your solution here
    pass`,
		Harness: Harness{Validator: `def check(raw, args, result, expected):
    nums = raw["nums"]
    if isinstance(result, bool) or not isinstance(result, int):
        return (False, "return the index of a peak as an int")
    if result < 0 or result >= len(nums):
        return (False, "index " + str(result) + " is outside the array")
    here = nums[result]
    if result > 0 and nums[result - 1] > here:
        return (False, "nums[" + str(result) + "] = " + str(here) + " is not a peak: its left neighbour " + str(nums[result - 1]) + " is larger")
    if result + 1 < len(nums) and nums[result + 1] > here:
        return (False, "nums[" + str(result) + "] = " + str(here) + " is not a peak: its right neighbour " + str(nums[result + 1]) + " is larger")
    return True`},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Stand at index mid and look at the right neighbour. If nums[mid] < nums[mid + 1], you can walk uphill to the right. Can that walk fall off a peak-less cliff, given that the end of the array counts as a drop?"},
			{Level: 2, Type: "insight", Content: "If nums[mid] < nums[mid + 1], the values rise to the right of mid, and they must come down by the end of the array (the ends count as -infinity), so a peak exists to the right: lo = mid + 1. Otherwise nums[mid] > nums[mid + 1], so mid itself or something to its left is a peak: hi = mid. This is the boundary template from Chapter 6, where the 'yes' question is 'is nums[mid] > nums[mid + 1]?'.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Set lo = 0 and hi = len(nums) - 1. The last index is always a candidate (its right side is -infinity), and it is the one index never compared with a right neighbour.\n2. While lo < hi, take mid = (lo + hi) // 2. Since mid < hi, mid + 1 is a valid index.\n3. If nums[mid] < nums[mid + 1], go uphill: lo = mid + 1.\n4. Otherwise a peak is at mid or to its left: hi = mid.\n5. When lo == hi, return lo."},
			{Level: 4, Type: "skeleton", Content: "def findPeakElement(nums):\n    lo, hi = 0, len(nums) - 1\n    while lo < hi:\n        mid = (lo + hi) // 2\n        if nums[mid] ____ nums[mid + 1]:\n            lo = ____\n        else:\n            hi = ____\n    return lo"},
		},
		Solution: Solution{
			Code: `def findPeakElement(nums):
    lo, hi = 0, len(nums) - 1
    while lo < hi:
        mid = (lo + hi) // 2
        if nums[mid] < nums[mid + 1]:
            lo = mid + 1
        else:
            hi = mid
    return lo`,
			Explanation: "The array is not sorted, yet the same boundary template works, because the question 'is nums[mid] > nums[mid + 1]?' has a guarantee behind it: going uphill always ends at a peak. Think of it as First True over the hidden array `nums[i] > nums[i + 1]`, where the last index counts as True because its right side is -infinity. The step beyond First True is that the yes/no question compares neighbours, which prepares you for Find Minimum in Rotated Sorted Array (153), where the question is a comparison with an endpoint.",
			KeyInsight:  "You do not need the array to be sorted. You need a rule that, from any probe, names a half that must contain an answer. Rising to the right means a peak is on the right, because the array has to come back down by the end.",
			WhyItWorks:  "Invariant: a peak always exists in [lo, hi]. It holds at the start since the whole array contains one (the maximum is a peak). If nums[mid] < nums[mid + 1], then walking right from mid + 1 the values either keep rising until the last element, which is a peak because its right side is -infinity, or they fall at some point, and the turning point is a peak. Either way there is a peak in [mid + 1, hi], so lo = mid + 1 keeps the invariant. Otherwise nums[mid] > nums[mid + 1] (they are never equal), so by the same argument walking left from mid there is a peak in [lo, mid]: either mid itself, or the values rise leftwards until a peak or the start of the array. So hi = mid keeps it. The range shrinks every step and ends on one index, which is therefore a peak. About log2(n) steps.",
			CommonMistakes: []string{
				"Comparing with the left neighbour or in the wrong direction (`nums[mid] > nums[mid + 1]` sending lo to mid + 1). That walks toward a valley instead of a peak, and on [1,2,3,4,5] it lands on index 0. Walk toward the larger neighbour.",
				"Checking only the two ends (`return 0 if nums[0] > nums[1] else len(nums) - 1`). A peak can sit anywhere in the middle, as in [1,3,2,1,0], where neither end is one. Only the larger-neighbour walk is guaranteed.",
				"Reading `nums[mid + 1]` when mid can be the last index, which crashes with IndexError. Starting with hi = len(nums) - 1 and looping while lo < hi keeps mid < hi, so mid + 1 is always valid.",
				"Writing `hi = mid - 1` in the else branch. Mid itself may be the only peak, and the code throws it away; on [1,3,2] the search ends at index 0, which is not a peak. Keep mid with hi = mid.",
			},
			TimeComplexity:  "O(log n)",
			SpaceComplexity: "O(1)",
			Alternatives: []AltSolution{
				{Name: "Linear scan", Code: `def findPeakElement(nums):
    for i in range(len(nums) - 1):
        if nums[i] > nums[i + 1]:
            return i
    return len(nums) - 1`, TimeComplexity: "O(n)", SpaceComplexity: "O(1)", Explanation: "Return the first index that is larger than its right neighbour, or the last index. It is correct and simple, but it reads every element when the peak is far to the right. The binary search needs O(log n), which matters for large arrays."},
			},
			Walkthrough: []WalkthroughStep{
				{Title: "Set the search range", Explanation: "Every peak candidate is an index from 0 to len(nums) - 1. The last index is the answer if nothing before it turns downhill, and it is never compared with a right neighbour, so it can be the hi of a half-open range.", CodeSnippet: "lo, hi = 0, len(nums) - 1", LineStart: 2, LineEnd: 2},
				{Title: "Pick the middle", Explanation: "Because lo < hi, mid is less than hi, so nums[mid + 1] is always a valid read.", CodeSnippet: "while lo < hi:\nmid = (lo + hi) // 2", LineStart: 3, LineEnd: 4},
				{Title: "Walk uphill", Explanation: "If the next value is larger, a peak lies to the right of mid. In [1,2,1,3,5,6,4] with mid = 3 (value 3 < 5), lo = 4.", CodeSnippet: "if nums[mid] < nums[mid + 1]:\nlo = mid + 1", LineStart: 5, LineEnd: 6},
				{Title: "Keep mid when it is already downhill", Explanation: "If the next value is smaller, mid is either a peak or has one to its left, so mid stays in the range.", CodeSnippet: "else:\nhi = mid", LineStart: 7, LineEnd: 8},
				{Title: "Read the answer", Explanation: "When one index is left, it is a peak: the walk uphill cannot get stuck anywhere else.", CodeSnippet: "return lo", LineStart: 9, LineEnd: 9},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it enough to compare with only the right neighbour?", Answer: "Because the question you ask must have a one-way answer. If nums[mid] < nums[mid + 1], a peak is certain on the right. If not, the values are not equal, so nums[mid] > nums[mid + 1] and a peak is certain at mid or to its left. The left neighbour is not needed to choose a half; it would only be needed to decide that mid itself is a peak, which the loop finds on its own when the range shrinks to it."},
			{Prompt: "What breaks if two neighbouring values may be equal?", Answer: "A tie nums[mid] == nums[mid + 1] gives no information about which side climbs: a flat stretch could continue in either direction, so you cannot rule out either half, and the guarantee that a peak exists on a chosen side fails. The O(log n) argument depends on neighbours always being different."},
		},
	},
	{
		ID:              "drill-capacity-to-ship-packages",
		Number:          1011,
		Kind:            KindDrill,
		Order:           7,
		Title:           "Capacity To Ship Packages Within D Days",
		Slug:            "capacity-to-ship-packages-within-d-days",
		Difficulty:      "Medium",
		Category:        "binary-search",
		Tags:            []string{"Array", "Binary Search", "Foundations"},
		RelatedChapters: []int{6},
		Description: `A conveyor belt carries packages whose weights are in the array ` + "`" + `weights` + "`" + `, in a fixed order. A ship loads them **in that order** (you cannot reorder or skip), and each day it carries some consecutive packages whose total weight does not exceed the ship's **capacity**. Given ` + "`" + `days` + "`" + `, return the smallest capacity that gets every package shipped within ` + "`" + `days` + "`" + ` days.

You cannot compute the answer directly, but you can *test* any capacity: load packages greedily and count the days used. That test is ` + "`" + `False` + "`" + ` for small capacities (too many days) and ` + "`" + `True` + "`" + ` from some point on (once the capacity is big enough, every bigger one works too), which is the monotone yes/no sequence from the earlier drills. The answer lies between the heaviest single package (a smaller ship could never lift it) and the total weight (one day is always enough). So binary search the capacity, using a helper that answers 'does capacity ` + "`" + `c` + "`" + ` fit in ` + "`" + `days` + "`" + ` days?'.

This is the pattern behind Koko Eating Bananas (875) in the ladder: search the answer, check feasibility.`,
		Constraints: []string{
			"1 <= days <= weights.length <= 50000",
			"1 <= weights[i] <= 500",
		},
		Examples: []Example{
			{Input: "weights = [1,2,3,4,5,6,7,8,9,10], days = 5", Output: "15", Explanation: "Capacity 15 works: [1,2,3,4,5], [6,7], [8], [9], [10] uses 5 days. Capacity 14 needs 6 days, because [1..4], [5,6], [7], [8], [9], [10] is already 6 groups."},
			{Input: "weights = [3,2,2,4,1,4], days = 3", Output: "6", Explanation: "Capacity 6: [3,2], [2,4], [1,4] is three days. Capacity 5 would need four."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"weights": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "days": 5}, Expected: 15, Note: "example 1: the best split is [1..5], [6,7], [8], [9], [10]"},
			{Input: map[string]interface{}{"weights": []int{3, 2, 2, 4, 1, 4}, "days": 3}, Expected: 6, Note: "example 2: capacity 6 ships it in three days"},
			{Input: map[string]interface{}{"weights": []int{1, 2, 3, 1, 1}, "days": 4}, Expected: 3, Hidden: true, Note: "example 3: four days for five packages"},
			{Input: map[string]interface{}{"weights": []int{7}, "days": 1}, Expected: 7, Hidden: true, Note: "a single package: capacity must be at least its weight"},
			{Input: map[string]interface{}{"weights": []int{10}, "days": 5}, Expected: 10, Hidden: true, Note: "more days than packages: capacity is still the heaviest package"},
			{Input: map[string]interface{}{"weights": []int{5, 5, 5, 5}, "days": 4}, Expected: 5, Hidden: true, Note: "one package per day: the answer is the heaviest, 5"},
			{Input: map[string]interface{}{"weights": []int{5, 5, 5, 5}, "days": 1}, Expected: 20, Hidden: true, Note: "everything in one day: the answer is the sum, 20"},
			{Input: map[string]interface{}{"weights": []int{1, 10}, "days": 2}, Expected: 10, Hidden: true, Note: "one heavy package: capacity below 10 can never move it, even if the day count looks fine"},
			{Input: map[string]interface{}{"weights": []int{10, 1, 1, 1, 1, 10}, "days": 2}, Expected: 12, Hidden: true, Note: "two heavy packages at the ends"},
			{Input: map[string]interface{}{"weights": []int{2, 2, 2, 2, 2, 2}, "days": 3}, Expected: 4, Hidden: true, Note: "equal weights: three days, two packages each"},
			{Input: map[string]interface{}{"weights": []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, "days": 3}, Expected: 17, Hidden: true, Note: "descending order: the order of packages cannot be changed"},
			{Input: map[string]interface{}{"weights": []int{1, 1, 1, 1, 1, 1, 1, 100}, "days": 2}, Expected: 100, Hidden: true, Note: "the last package is the heaviest and ships alone"},
			{Input: map[string]interface{}{"weights": bsDrillWeights(50000), "days": 1}, Expected: 12525000, Hidden: true, Note: "50000 packages, all in one day: the capacity range is about 12.5 million, so a scan over capacities is far too slow"},
			{Input: map[string]interface{}{"weights": bsDrillWeights(50000), "days": 7}, Expected: 1789390, Hidden: true, Note: "50000 packages over 7 days: the same range, a different answer"},
		},
		TimeComplexity:  "O(n log S) where S is the total weight",
		SpaceComplexity: "O(1)",
		StarterCode: `def shipWithinDays(weights, days):
    # Write your solution here
    pass`,
		Harness: Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose someone told you a capacity. How would you check, in one pass over weights, whether the packages fit in days? And if capacity c works, does a bigger capacity work too?"},
			{Level: 2, Type: "insight", Content: "A greedy pass answers 'does capacity c fit?': keep loading until the next package would overflow, then start a new day. Fitting is monotone: if c works, every bigger capacity works, so the capacities read 'no, no, ..., yes, yes'. The answer is the first yes, which is the boundary template from Chapter 6 applied to the answer range [max(weights), sum(weights)], the same idea as First True and Sqrt.", ChapterRef: 6},
			{Level: 3, Type: "plan", Content: "1. Write fits(capacity): used = 1, load = 0; for each weight, if load + weight > capacity start a new day (used += 1, load = 0); then load += weight. Return used <= days.\n2. Set the range: lo = max(weights) (a smaller ship cannot lift the heaviest package) and hi = sum(weights) (one day always fits).\n3. While lo < hi, mid = (lo + hi) // 2.\n4. If fits(mid), the answer is at most mid: hi = mid. Otherwise lo = mid + 1.\n5. Return lo."},
			{Level: 4, Type: "skeleton", Content: "def shipWithinDays(weights, days):\n    def fits(capacity):\n        used, load = 1, 0\n        for weight in weights:\n            if load + weight > capacity:\n                used += 1\n                load = 0\n            load += weight\n        return used <= days\n\n    lo, hi = ____, ____\n    while lo < hi:\n        mid = (lo + hi) // 2\n        if fits(mid):\n            hi = ____\n        else:\n            lo = ____\n    return lo"},
		},
		Solution: Solution{
			Code: `def shipWithinDays(weights, days):
    def fits(capacity):
        used, load = 1, 0
        for weight in weights:
            if load + weight > capacity:
                used += 1
                load = 0
            load += weight
        return used <= days

    lo, hi = max(weights), sum(weights)
    while lo < hi:
        mid = (lo + hi) // 2
        if fits(mid):
            hi = mid
        else:
            lo = mid + 1
    return lo`,
			Explanation: "The yes/no question is no longer a single comparison. It is a helper that scans the whole array, so each probe costs O(n), but the pattern of the search is unchanged from First True and Sqrt: search the answer range, ask a monotone question, keep mid on a yes. It is the bridge to Koko Eating Bananas (875), where the helper adds up hours instead of counting days, and to the many 'minimise the maximum' problems that follow it.",
			KeyInsight:  "If you can check a candidate answer quickly and the check flips once as the candidate grows, you do not need a formula for the answer: binary search the candidates. A bigger capacity never needs more days, so the smallest working capacity is a boundary.",
			WhyItWorks:  "Monotonicity: a ship that fits in days days with capacity c can use the same loads with any bigger capacity, so fits(c) implies fits(c + 1), and the sequence of answers is 'no, ..., no, yes, ..., yes'. The greedy check is right because loading as many packages as possible each day never uses more days than stopping early (an exchange argument: any schedule can be shifted so that each day carries the longest possible prefix of what is left). The range is [max(weights), sum(weights)]: below the heaviest weight a package cannot be carried at all, and the total weight always fits in one day, so fits(hi) is true and a first yes exists. The invariant is that everything below lo is a no and hi is a yes, so the first yes is in [lo, hi]; fits(mid) true gives hi = mid, false gives lo = mid + 1. There are log2(sum) probes, each O(n).",
			CommonMistakes: []string{
				"Starting the range at `lo = 1` (or 0). A capacity below the heaviest package cannot move it, but the greedy loop can still give a small count: [1,10] with 2 days returns a capacity below 10. Start at max(weights).",
				"Scanning capacities one at a time (`for c in range(max(weights), sum(weights) + 1): if fits(c): return c`). It is correct, but with 50000 packages the range holds millions of values, each costing a pass over the array, and it runs out of time. Binary search the range.",
				"Using `>=` in the overflow test (`load + weight >= capacity`), which starts a new day when a package fits exactly. Packages that add up to exactly the capacity are allowed, so the test is `>`.",
				"Forgetting to count the first day: starting with `used = 0` makes every schedule look one day shorter, so the answer is too small. A ship is always in use on day one.",
				"Moving `hi = mid - 1` after a success, which throws away a capacity that works and might be the smallest. Keep mid when fits(mid) is true.",
			},
			TimeComplexity:  "O(n log S)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Write the yes/no question", Explanation: "fits(capacity) loads packages in order and counts days: a new day starts when the next package would overflow. It is the 'flags[mid]' of this problem, computed on demand.", CodeSnippet: "def fits(capacity):\nused, load = 1, 0", LineStart: 2, LineEnd: 3},
				{Title: "Count the days", Explanation: "A package that does not fit starts a new day with an empty ship. With [3,2,2,4,1,4] and capacity 6 the days are [3,2], [2,4], [1,4].", CodeSnippet: "if load + weight > capacity:\nused += 1\nload = 0\nload += weight", LineStart: 5, LineEnd: 8},
				{Title: "Bound the answer range", Explanation: "No ship smaller than the heaviest package can work, and a ship as big as the total weight always takes one day. The answer is in [max, sum].", CodeSnippet: "lo, hi = max(weights), sum(weights)", LineStart: 11, LineEnd: 11},
				{Title: "Keep the half that holds the boundary", Explanation: "If capacity mid is enough, the answer is mid or smaller (hi = mid); otherwise it is bigger (lo = mid + 1).", CodeSnippet: "if fits(mid):\nhi = mid\nelse:\nlo = mid + 1", LineStart: 14, LineEnd: 17},
				{Title: "Read the answer", Explanation: "When lo == hi the range has closed on the smallest capacity that fits.", CodeSnippet: "return lo", LineStart: 18, LineEnd: 18},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why do the packages have to stay in order, and how would the answer change if you could reorder them?", Answer: "The check loads packages consecutively, so it assumes order. If reordering were allowed this would become bin packing, a much harder problem with no known fast exact algorithm. The fixed order is what makes the greedy check correct and the problem monotone."},
			{Prompt: "Why is the binary search over capacities and not over the array indices?", Answer: "Nothing about the index tells you the answer: the array is not sorted. What is monotone is the property 'this capacity is enough', which only depends on the capacity. So the search space is the range of possible answers, and the array is only used inside the check. Recognising what to search is the main skill the ladder asks of you next."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, BinarySearchDrills...)
}

// bsDrillEvens returns 0, 2, 4, ... with n items.
func bsDrillEvens(n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = 2 * i
	}
	return out
}

// bsDrillRuns returns n items made of runs of the value 7 in the middle:
// the first low items are 3, the next mid items are 7, the rest are 9.
func bsDrillRuns(low, mid, high int) []int {
	out := make([]int, 0, low+mid+high)
	for i := 0; i < low; i++ {
		out = append(out, 3)
	}
	for i := 0; i < mid; i++ {
		out = append(out, 7)
	}
	for i := 0; i < high; i++ {
		out = append(out, 9)
	}
	return out
}

// bsDrillFlags returns n flags that are false before index firstTrue and true
// from there on.
func bsDrillFlags(n, firstTrue int) []bool {
	out := make([]bool, n)
	for i := firstTrue; i < n; i++ {
		out[i] = true
	}
	return out
}

// bsDrillWeights returns n deterministic weights between 1 and 500.
func bsDrillWeights(n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = (i*37)%500 + 1
	}
	return out
}

// bsDrillPeaks returns an array that rises to a peak at index peak and falls
// after it, with n items and all neighbours different.
func bsDrillPeaks(n, peak int) []int {
	out := make([]int, n)
	for i := range out {
		if i <= peak {
			out[i] = i
		} else {
			out[i] = 2*peak - i
		}
	}
	return out
}
