package practice

// TwoPointersDrills are the two-pointer foundations: the three shapes on their
// own before the ladder applies them. Opposite ends (reverse in place),
// read/write pointers moving the same way (remove duplicates, move zeroes),
// and one pointer per sequence (merge, is subsequence, intersect), then
// opposite ends that choose between the ends (squares of a sorted array) and
// the run-skipping on a hit that 3Sum needs (count pairs). They follow
// docs/pedagogy.md like any problem; the in-place drills are judged by the
// final contents of the list, and Remove Duplicates by a validator that
// checks the returned k and the kept front of the list.
var TwoPointersDrills = []*Problem{
	{
		ID:              "drill-reverse-in-place",
		Number:          344,
		Kind:            KindDrill,
		Order:           1,
		Title:           "Reverse a List in Place",
		Slug:            "reverse-string",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Two Pointers", "Array", "In Place", "Foundations"},
		RelatedChapters: []int{3},
		Description:     "You are given a list `s` of single-character strings. **Reverse it in place**: change the list itself so that its items appear in the opposite order, and return nothing. The judge reads `s` after your function finishes.\n\nThis is the plainest two-pointer shape, **opposite ends**. Put one pointer `left` on the first item and one pointer `right` on the last. Swap the two items, move `left` one step right and `right` one step left, and repeat. Stop when the pointers meet or cross: every item has then been swapped into its mirror position. Use the pointers; do not build a reversed copy and assign it to `s` (that would rebind the name and leave the caller's list unchanged).",
		Constraints:     []string{"0 <= s.length <= 100000", "Every item of s is a single character"},
		Examples: []Example{
			{Input: "s = [\"h\",\"e\",\"l\",\"l\",\"o\"]", Output: "[\"o\",\"l\",\"l\",\"e\",\"h\"]", Explanation: "Swap h and o, then e and l. The middle l is alone, so left and right meet and the loop stops."},
			{Input: "s = [\"H\",\"a\",\"n\",\"n\",\"a\",\"h\"]", Output: "[\"h\",\"a\",\"n\",\"n\",\"a\",\"H\"]", Explanation: "Three swaps: H with h, a with a, n with n. The pointers cross after the third swap."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": []string{"h", "e", "l", "l", "o"}}, Expected: []string{"o", "l", "l", "e", "h"}, Note: "example 1: odd length; the middle letter stays where it is"},
			{Input: map[string]interface{}{"s": []string{"H", "a", "n", "n", "a", "h"}}, Expected: []string{"h", "a", "n", "n", "a", "H"}, Note: "example 2: even length, and the answer is the same letters in reverse order, not the same word"},
			{Input: map[string]interface{}{"s": []string{}}, Expected: []string{}, Hidden: true, Note: "empty list: nothing to swap"},
			{Input: map[string]interface{}{"s": []string{"a"}}, Expected: []string{"a"}, Hidden: true, Note: "one item: left and right start on the same index and must not swap or crash"},
			{Input: map[string]interface{}{"s": []string{"a", "b"}}, Expected: []string{"b", "a"}, Hidden: true, Note: "two items: exactly one swap (a loop that stops one step early does nothing)"},
			{Input: map[string]interface{}{"s": []string{"a", "b", "c", "d", "e", "f", "g"}}, Expected: []string{"g", "f", "e", "d", "c", "b", "a"}, Hidden: true, Note: "odd length 7: three swaps, the middle stays"},
			{Input: map[string]interface{}{"s": []string{"a", "b", "c", "d", "e", "f"}}, Expected: []string{"f", "e", "d", "c", "b", "a"}, Hidden: true, Note: "even length 6: three swaps and the pointers cross without meeting"},
			{Input: map[string]interface{}{"s": []string{"a", "a", "b", "b"}}, Expected: []string{"b", "b", "a", "a"}, Hidden: true, Note: "repeated letters still end up mirrored: [b,b,a,a]"},
			{Input: map[string]interface{}{"s": []string{"1", "2", "3", "4"}}, Expected: []string{"4", "3", "2", "1"}, Hidden: true, Note: "digits as characters, reversed like any other items"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def reverseString(s):\n    # Modify s in place; return nothing\n    pass",
		Harness:         Harness{InPlace: "s"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Which two items end up swapping places in the reversed list? Where are they when you start, and where should you look next once they are swapped?"},
			{Level: 2, Type: "insight", Content: "Item 0 and the last item trade places, then item 1 and the second to last, and so on. Keep one pointer at each end and move them toward each other: this is two pointers from opposite ends (Chapter 3). Stop when left is no longer less than right, because after that every pair has been swapped (or the single middle item needs no swap).", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Set up the pointers: left at the first index, right at the last index.\n2. Repeat while left < right.\n3. Swap s[left] and s[right].\n4. Move both pointers one step inward.\n5. Return nothing: the list was changed in place."},
			{Level: 4, Type: "skeleton", Content: "def reverseString(s):\n    left, right = ____, ____\n    while ____:\n        s[left], s[right] = s[right], s[left]\n        left += 1\n        right -= 1"},
		},
		Solution: Solution{
			Code:            "def reverseString(s):\n    left, right = 0, len(s) - 1\n    while left < right:\n        s[left], s[right] = s[right], s[left]\n        left += 1\n        right -= 1",
			Explanation:     "Two pointers start at the two ends of the list and walk toward each other, swapping what they point at. It is the simplest use of the opposite-ends shape and uses no extra list. Later drills keep the same loop and change only what happens at each step: Squares of a Sorted Array chooses between the ends instead of swapping them, and Two Sum II and Container With Most Water move the pointers according to a comparison.",
			BruteForce:      "Build a new list by walking s from the back and appending each item, then copy it over s: O(n) time but O(n) extra space, and the copy is work the swap never needs.",
			KeyInsight:      "Each item has exactly one destination: the mirror position. Item i belongs at n-1-i, and that position's item belongs at i. So one swap puts two items in their final places, and walking in from both ends visits every pair once.",
			WhyItWorks:      "Invariant: before each iteration, everything outside the window s[left..right] has already been swapped into its final mirrored place, and the window is the reversed middle still to do. Initially the window is the whole list. A swap fixes the two window ends, and the pointers step inward, so the window shrinks by two while the invariant still holds. When left >= right the window holds 0 or 1 items; an empty window means nothing is left and a single item is its own mirror. Every pair is swapped once, so the loop runs n // 2 times: O(n) time and O(1) space.",
			CommonMistakes:  []string{"Looping over every index (`for i in range(len(s))`) and swapping s[i] with s[n-1-i], which swaps every pair twice and restores the original order. Stop when the pointers meet: only n // 2 swaps are needed.", "Assigning a reversed copy (`s = s[::-1]`). That rebinds the local name; the caller's list is unchanged. Change the items of the list itself (swap them, or `s[:] = ...`).", "An off-by-one loop condition such as `while left < right - 1`, which leaves the innermost pair unswapped (a two-item list never changes).", "Moving only one pointer, or moving them the wrong way, so the loop never ends or swaps the same pair repeatedly."},
			TimeComplexity:  "O(n): each item is touched once",
			SpaceComplexity: "O(1): only two index variables",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the pointers", Explanation: "left starts on the first index and right on the last. For [h,e,l,l,o] that is 0 and 4.", CodeSnippet: "left, right = 0, len(s) - 1", LineStart: 2, LineEnd: 2},
				{Title: "Loop while the pointers have not met", Explanation: "As long as left < right there are two different items to exchange. Equal pointers mean one middle item, and crossed pointers mean every pair is done.", CodeSnippet: "while left < right:", LineStart: 3, LineEnd: 3},
				{Title: "Swap the ends", Explanation: "Exchange the two items. After the first swap [h,e,l,l,o] is [o,e,l,l,h].", CodeSnippet: "s[left], s[right] = s[right], s[left]", LineStart: 4, LineEnd: 4},
				{Title: "Move inward", Explanation: "Step left right and right left so the next pair is the next-to-outer items. Now left is 1 and right is 3.", CodeSnippet: "left += 1\nright -= 1", LineStart: 5, LineEnd: 6},
			},
			Alternatives: []AltSolution{
				{Name: "Slice assignment", Code: "def reverseString(s):\n    s[:] = s[::-1]", TimeComplexity: "O(n)", SpaceComplexity: "O(n)", Explanation: "The slice s[::-1] builds a reversed copy, and s[:] = ... writes it back into the same list. It is correct and short, but it uses O(n) extra space; the two-pointer version needs none."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the loop condition use left < right and not left <= right?", Answer: "When the pointers are equal they point at the same item, so swapping it with itself does nothing, and the middle item of an odd-length list is already in its final place. Using left <= right would be harmless here but would run one pointless iteration; using left > right would never start."},
			{Prompt: "What if the list were a linked list instead of an array?", Answer: "You could not move right backward in O(1), because a singly linked node has no previous link. Reversing a linked list rewrites next pointers instead, which is the Linked Lists chapter."},
		},
	},
	{
		ID:              "drill-remove-duplicates-sorted",
		Number:          26,
		Kind:            KindDrill,
		Order:           2,
		Title:           "Remove Duplicates from a Sorted Array",
		Slug:            "remove-duplicates-from-sorted-array",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Two Pointers", "Array", "Read/Write Pointers", "Foundations"},
		RelatedChapters: []int{3},
		Description:     "You are given a list `nums` sorted in non-decreasing order. Rewrite the front of `nums` so that it holds each distinct value once, in the same order, and return `k`, the number of distinct values. What is stored after index `k - 1` does not matter.\n\nThis is the second two-pointer shape, **read and write pointers moving the same direction**. A `read` pointer scans every item. A `write` pointer marks where the next kept item goes: everything before `write` is the finished answer. When `read` finds a value that differs from the last kept one, copy it to `write` and advance `write`. Do it with O(1) extra space; the judge checks both the returned `k` and the first `k` items of `nums`.",
		Constraints:     []string{"0 <= nums.length <= 100000", "-100 <= nums[i] <= 100", "nums is sorted in non-decreasing order"},
		Examples: []Example{
			{Input: "nums = [1,1,2]", Output: "2, with nums starting [1,2,...]", Explanation: "read skips the second 1 because it equals the last kept value, then copies 2 to index 1. k is 2."},
			{Input: "nums = [0,0,1,1,1,2,2,3,3,4]", Output: "5, with nums starting [0,1,2,3,4,...]", Explanation: "Each run of equal values is written once, so five values are kept."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 1, 2}}, Expected: 2, Note: "example 1: [1,2] are kept, so the answer is 2"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, Expected: 5, Note: "example 2: five distinct values, each run of duplicates collapses to one"},
			{Input: map[string]interface{}{"nums": []int{}}, Expected: 0, Hidden: true, Note: "empty array: return 0 (nums[0] must not be read)"},
			{Input: map[string]interface{}{"nums": []int{7}}, Expected: 1, Hidden: true, Note: "one item: nothing to compare, k is 1"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}}, Expected: 3, Hidden: true, Note: "all distinct: every item is kept in place"},
			{Input: map[string]interface{}{"nums": []int{5, 5, 5, 5}}, Expected: 1, Hidden: true, Note: "all equal: only the first is kept, k is 1"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 2, 2, 2}}, Expected: 2, Hidden: true, Note: "a long run of duplicates at the end"},
			{Input: map[string]interface{}{"nums": []int{-3, -3, 0, 0, 2}}, Expected: 3, Hidden: true, Note: "negative values and zero are ordinary values"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 2, 3, 3, 3, 4, 4, 5}}, Expected: 5, Hidden: true, Note: "runs of different lengths: the kept prefix must be exactly [1,2,3,4,5]"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def removeDuplicates(nums):\n    # Rewrite the front of nums; return k, the number of distinct values\n    pass",
		Harness:         Harness{Validator: "def check(raw, args, result, expected):\n    kept = []\n    for x in raw[\"nums\"]:\n        if not kept or kept[-1] != x:\n            kept.append(x)\n    if not isinstance(result, int):\n        return False, \"return k, the number of distinct values (an integer), after rewriting the front of nums\"\n    if result != len(kept):\n        return False, \"returned \" + str(result) + \" but the array has \" + str(len(kept)) + \" distinct values\"\n    front = args[\"nums\"][:result]\n    if front != kept:\n        return False, \"the first \" + str(result) + \" items of nums are \" + str(front) + \" but should be \" + str(kept)\n    return True, \"\""},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "After you have decided to keep some values, what do you compare the next value against to know whether it is new? And where in the array would you put it?"},
			{Level: 2, Type: "insight", Content: "Because the array is sorted, a value is a duplicate exactly when it equals the last value you kept. Keep two pointers moving the same direction: read looks at every item, and write is the next free slot for a kept item (Chapter 3, read/write pointers). Everything before write is already the answer.", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Handle the empty array: return 0.\n2. The first item is always kept, so set write to 1.\n3. Move read across the rest of the array.\n4. If nums[read] differs from the last kept item nums[write - 1], copy it to nums[write] and add 1 to write.\n5. Return write."},
			{Level: 4, Type: "skeleton", Content: "def removeDuplicates(nums):\n    if not nums:\n        return 0\n    write = 1\n    for read in range(1, len(nums)):\n        if nums[read] != ____:\n            nums[____] = nums[read]\n            write += 1\n    return ____"},
		},
		Solution: Solution{
			Code:            "def removeDuplicates(nums):\n    if not nums:\n        return 0\n    write = 1\n    for read in range(1, len(nums)):\n        if nums[read] != nums[write - 1]:\n            nums[write] = nums[read]\n            write += 1\n    return write",
			Explanation:     "A read pointer scans the sorted array and a write pointer builds the answer in the same array. The array being sorted turns 'have I seen this value?' into 'does it equal the last kept value?', which needs no set. This read/write pair is the engine behind every in-place filter; Move Zeroes, next, uses it with a different keep rule.",
			BruteForce:      "Put the values in a set and sort them (or scan with a growing list of seen values): O(n) to O(n log n) time with O(n) extra space, or delete duplicates from the middle of the list one by one, which shifts the rest each time and costs O(n^2). The waste is not using the sorted order: duplicates are always next to each other.",
			KeyInsight:      "In a sorted array equal values sit side by side, so a duplicate is just a value equal to the previous kept one. One comparison per item is enough, and the kept items can be packed into the front as you go because write never gets ahead of read.",
			WhyItWorks:      "Invariant: before examining nums[read], nums[:write] holds the distinct values of nums[:read] in order, and nums[write - 1] is the largest of them. If nums[read] equals nums[write - 1] it is a duplicate (equal values are adjacent, so it cannot also be an earlier different value) and is skipped. Otherwise it is larger (the array is sorted) and new, so it is copied to nums[write] and write grows by one. Since write <= read, a copy never overwrites an item that has not been read. When read reaches the end, nums[:write] holds every distinct value once.",
			CommonMistakes:  []string{"Counting the distinct values but never copying them forward (for example returning `len(set(nums))`). The judge also checks the first k items of nums.", "An off-by-one in the starting positions: setting write to 0 while also starting read at 0, or returning `write - 1`, so the count is wrong by one. The first item is always kept, so write starts at 1.", "Reading `nums[0]` or `nums[write - 1]` on an empty array. Return 0 first.", "Deleting items from the list while scanning it (`del nums[i]` or `nums.remove`). Every deletion shifts the rest of the list, which is O(n^2) overall and also skips items."},
			TimeComplexity:  "O(n): read visits each item once",
			SpaceComplexity: "O(1): only two index variables",
			Walkthrough: []WalkthroughStep{
				{Title: "Keep the first item", Explanation: "Any non-empty array keeps its first item, so write, the next free slot, starts at 1. An empty array returns 0.", CodeSnippet: "if not nums:\n    return 0\nwrite = 1", LineStart: 2, LineEnd: 4},
				{Title: "Scan with the read pointer", Explanation: "read looks at every later item. In [1,1,2] read starts at index 1.", CodeSnippet: "for read in range(1, len(nums)):", LineStart: 5, LineEnd: 5},
				{Title: "Compare with the last kept item", Explanation: "A sorted array makes the last kept item nums[write - 1] the only value to compare with. Equal means a duplicate, so nothing happens.", CodeSnippet: "if nums[read] != nums[write - 1]:", LineStart: 6, LineEnd: 6},
				{Title: "Write the new item", Explanation: "Copy the new value into the next free slot and advance write. For [1,1,2], 2 goes to index 1 and write becomes 2.", CodeSnippet: "nums[write] = nums[read]\nwrite += 1", LineStart: 7, LineEnd: 8},
				{Title: "Return the count", Explanation: "write is both the next free slot and the number of kept items.", CodeSnippet: "return write", LineStart: 9, LineEnd: 9},
			},
			Alternatives: []AltSolution{
				{Name: "Compare with the previous item", Code: "def removeDuplicates(nums):\n    write = 0\n    for read in range(len(nums)):\n        if write == 0 or nums[read] != nums[write - 1]:\n            nums[write] = nums[read]\n            write += 1\n    return write", TimeComplexity: "O(n)", SpaceComplexity: "O(1)", Explanation: "The same idea without the empty-array special case: the first item is kept because write is 0."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to overwrite nums[write] while read is still scanning?", Answer: "write never exceeds read, so the slot being overwritten is at or behind the item just read. Everything behind read has already been examined, so no unread item is destroyed."},
			{Prompt: "What changes if the array is not sorted?", Answer: "Equal values are no longer adjacent, so comparing with the last kept item is not enough: [1,2,1] would keep both 1s. You would need a set of seen values (O(n) space) or to sort first."},
		},
	},
	{
		ID:              "drill-move-zeroes",
		Number:          283,
		Kind:            KindDrill,
		Order:           3,
		Title:           "Move Zeroes",
		Slug:            "move-zeroes",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Two Pointers", "Array", "Read/Write Pointers", "Foundations"},
		RelatedChapters: []int{3},
		Description:     "You are given a list `nums` of integers. Move every `0` to the end of the list while keeping the relative order of the nonzero items. Change `nums` in place and return nothing; the judge reads `nums` after your function finishes.\n\nThis uses the same read/write pointers as removing duplicates, with a different keep rule. `read` scans every item. `write` is the index where the next nonzero item belongs, so everything before `write` is the nonzero items found so far, in order. Whenever `read` finds a nonzero item, put it at `write` and advance `write`. Swapping the two (instead of only copying) leaves the zeros behind automatically, so you never have to fill the end with zeros in a second pass.",
		Constraints:     []string{"0 <= nums.length <= 100000", "-1000000 <= nums[i] <= 1000000"},
		Examples: []Example{
			{Input: "nums = [0,1,0,3,12]", Output: "[1,3,12,0,0]", Explanation: "The nonzero items 1, 3, 12 keep their order and move to the front; the zeros are pushed behind them."},
			{Input: "nums = [0]", Output: "[0]", Explanation: "A list of one zero has nothing to move."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{0, 1, 0, 3, 12}}, Expected: []int{1, 3, 12, 0, 0}, Note: "example 1: nonzero items keep their order, zeros go to the end"},
			{Input: map[string]interface{}{"nums": []int{0}}, Expected: []int{0}, Note: "example 2: a single zero"},
			{Input: map[string]interface{}{"nums": []int{}}, Expected: []int{}, Hidden: true, Note: "empty list"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}}, Expected: []int{1, 2, 3}, Hidden: true, Note: "no zeros: nothing moves"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0}}, Expected: []int{0, 0, 0}, Hidden: true, Note: "all zeros: nothing to move"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 1}}, Expected: []int{1, 0, 0}, Hidden: true, Note: "one nonzero item that must travel past two zeros"},
			{Input: map[string]interface{}{"nums": []int{-1, 0, -2, 0, 3}}, Expected: []int{-1, -2, 3, 0, 0}, Hidden: true, Note: "negative numbers are nonzero and must be kept (an `> 0` test drops them)"},
			{Input: map[string]interface{}{"nums": []int{4, 0, 0, 5}}, Expected: []int{4, 5, 0, 0}, Hidden: true, Note: "zeros in the middle: the order 4, 5 must be kept (swapping from both ends gives 5, 4)"},
			{Input: map[string]interface{}{"nums": []int{2, 1, 0, 0, 3, 1}}, Expected: []int{2, 1, 3, 1, 0, 0}, Hidden: true, Note: "duplicates among the nonzero items keep their order too"},
			{Input: map[string]interface{}{"nums": tpDrillZeroEveryOther(100000)}, Expected: tpDrillZeroEveryOtherAnswer(100000), Hidden: true, Timeout: 5000, Note: "100000 items with every other one a zero: needs O(n); removing zeros from the front and appending runs out of time"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def moveZeroes(nums):\n    # Modify nums in place; return nothing\n    pass",
		Harness:         Harness{InPlace: "nums"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Imagine collecting the nonzero items one at a time at the front. What do you need to remember to know where the next one goes?"},
			{Level: 2, Type: "insight", Content: "Use read/write pointers (Chapter 3): read scans every item, and write is the position where the next nonzero item belongs. Everything before write is the nonzero items found so far, in their original order. The zeros are whatever lies between write and read.", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Start write at 0.\n2. Move read across the whole list.\n3. When nums[read] is not 0, swap nums[write] and nums[read].\n4. Then advance write by one.\n5. Return nothing: the list was changed in place."},
			{Level: 4, Type: "skeleton", Content: "def moveZeroes(nums):\n    write = 0\n    for read in range(len(nums)):\n        if nums[read] ____ 0:\n            nums[write], nums[read] = ____, ____\n            write += 1"},
		},
		Solution: Solution{
			Code:            "def moveZeroes(nums):\n    write = 0\n    for read in range(len(nums)):\n        if nums[read] != 0:\n            nums[write], nums[read] = nums[read], nums[write]\n            write += 1",
			Explanation:     "Same pointers as Remove Duplicates, new keep rule: keep the nonzero items. Because read never falls behind write, swapping a nonzero item forward drops a zero into the slot it left, so the zeros end up behind the kept items without a second pass. It builds directly on the read/write invariant from the previous drill.",
			BruteForce:      "Make a new list of the nonzero items and then append the zeros: O(n) time but O(n) extra space. A common in-place version removes each zero and appends a new one, which shifts the rest of the list every time: O(n^2).",
			KeyInsight:      "The zeros need not be handled at all. If every nonzero item is moved to the next free slot, in the order it is found, the zeros are exactly what is left over behind them, and a swap keeps the list a permutation of what it was.",
			WhyItWorks:      "Invariant: before examining nums[read], nums[:write] is the nonzero items of nums[:read] in their original order, and nums[write:read] is all zeros. If nums[read] is zero, read moves on and the zero block grows by one. If it is nonzero, swapping it with nums[write] (the first zero of the block, or itself when the block is empty) puts it at the end of the nonzero prefix and moves a zero to the back of the block; write grows by one and the invariant still holds. When read reaches the end, nums[:write] is all the nonzero items in order and nums[write:] is all zeros. The order is stable because items are only placed in the order read finds them.",
			CommonMistakes:  []string{"Treating 'not zero' as 'positive' (`if nums[read] > 0`), which leaves negative numbers where they were. Compare with `!= 0`.", "Swapping from both ends (a left pointer for zeros and a right pointer for nonzero items), which finishes in one pass but reverses the order of the nonzero items. The order must stay stable, so the pointers move in the same direction.", "Copying the nonzero items forward without writing zeros into the rest of the list, which leaves stale values behind the new end.", "Building a new list and returning it (or assigning `nums = [...]`). The judge reads the original list, so the change must be made to nums itself.", "Removing zeros from the list as you meet them, which shifts the remaining items and is O(n^2)."},
			TimeComplexity:  "O(n): each item is read once",
			SpaceComplexity: "O(1): swaps in place",
			Walkthrough: []WalkthroughStep{
				{Title: "Start the write pointer", Explanation: "write is where the next nonzero item belongs. Nothing has been kept yet, so it is 0.", CodeSnippet: "write = 0", LineStart: 2, LineEnd: 2},
				{Title: "Scan with the read pointer", Explanation: "read visits every index once, left to right.", CodeSnippet: "for read in range(len(nums)):", LineStart: 3, LineEnd: 3},
				{Title: "Keep only the nonzero items", Explanation: "A zero is skipped and costs nothing. A nonzero item is moved to write. In [0,1,0,3,12], read at index 1 finds the 1.", CodeSnippet: "if nums[read] != 0:", LineStart: 4, LineEnd: 4},
				{Title: "Swap and advance", Explanation: "Swapping puts the nonzero item at write and the zero that was there at read. Now [1,0,0,3,12] and write is 1.", CodeSnippet: "nums[write], nums[read] = nums[read], nums[write]\nwrite += 1", LineStart: 5, LineEnd: 6},
			},
			Alternatives: []AltSolution{
				{Name: "Copy then fill", Code: "def moveZeroes(nums):\n    write = 0\n    for read in range(len(nums)):\n        if nums[read] != 0:\n            nums[write] = nums[read]\n            write += 1\n    for i in range(write, len(nums)):\n        nums[i] = 0", TimeComplexity: "O(n)", SpaceComplexity: "O(1)", Explanation: "Two passes: pack the nonzero items to the front exactly like Remove Duplicates, then overwrite everything after write with zeros. It avoids swapping at the cost of a second loop."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must both pointers move in the same direction here, while Reverse a List uses opposite ends?", Answer: "Reversing wants the first item to swap with the last, so the pointers start at the ends. Moving zeroes must keep the nonzero items in their original order, which only a scan from left to right does: each kept item is placed after the one before it."},
			{Prompt: "What if you had to move a different value, say all 5s, to the end?", Answer: "Only the test changes: compare with 5 instead of 0 (keep when nums[read] != 5). The invariant is the same: the prefix up to write holds the kept items in order."},
		},
	},
	{
		ID:              "drill-merge-two-sorted-arrays",
		Number:          0,
		Kind:            KindDrill,
		Order:           4,
		Title:           "Merge Two Sorted Arrays",
		Slug:            "merge-two-sorted-arrays",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Two Pointers", "Array", "Two Arrays", "Merge", "Foundations"},
		RelatedChapters: []int{3},
		Description:     "You are given two lists `a` and `b`, each sorted in non-decreasing order. Return a **new** list containing every item of both, in non-decreasing order. Items that are equal in the two lists must all appear in the result (it keeps duplicates).\n\nThis is the third two-pointer shape, **one pointer per array**. `i` points at the smallest item of `a` not yet used, and `j` does the same for `b`. At every step the next item of the answer is the smaller of `a[i]` and `b[j]`: take it, and advance only the pointer you took from. When one array runs out, the rest of the other is already sorted and goes straight to the end. (Python's `sorted(a + b)` also works, but this drill is about doing it with the two pointers, which is also how merge sort combines its halves. LeetCode's Merge Sorted Array, number 88, merges into the first array in place; here you return a new list.)",
		Constraints:     []string{"0 <= a.length, b.length <= 100000", "-1000000000 <= a[i], b[j] <= 1000000000", "a and b are sorted in non-decreasing order"},
		Examples: []Example{
			{Input: "a = [1,2,4], b = [1,3,4]", Output: "[1,1,2,3,4,4]", Explanation: "Compare 1 and 1: take a's 1. Then 2 and 1: take b's 1. Then 2 and 3: take 2, then 4 and 3: take 3, then 4 and 4: take a's 4, and a is empty, so b's final 4 is copied over."},
			{Input: "a = [1], b = []", Output: "[1]", Explanation: "b is empty, so a's items are the whole answer."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"a": []int{1, 2, 4}, "b": []int{1, 3, 4}}, Expected: []int{1, 1, 2, 3, 4, 4}, Note: "example 1: interleaved values with a duplicate across the arrays"},
			{Input: map[string]interface{}{"a": []int{1}, "b": []int{}}, Expected: []int{1}, Note: "example 2: the second array is empty"},
			{Input: map[string]interface{}{"a": []int{}, "b": []int{}}, Expected: []int{}, Hidden: true, Note: "both empty"},
			{Input: map[string]interface{}{"a": []int{}, "b": []int{2, 5}}, Expected: []int{2, 5}, Hidden: true, Note: "first array empty: the whole second array is the tail"},
			{Input: map[string]interface{}{"a": []int{1, 2, 3}, "b": []int{4, 5, 6}}, Expected: []int{1, 2, 3, 4, 5, 6}, Hidden: true, Note: "every item of a comes first: b is entirely a leftover tail"},
			{Input: map[string]interface{}{"a": []int{4, 5, 6}, "b": []int{1, 2, 3}}, Expected: []int{1, 2, 3, 4, 5, 6}, Hidden: true, Note: "every item of b comes first: a is entirely a leftover tail"},
			{Input: map[string]interface{}{"a": []int{1, 1, 1}, "b": []int{1, 1}}, Expected: []int{1, 1, 1, 1, 1}, Hidden: true, Note: "all equal: no value may be dropped when the two sides tie"},
			{Input: map[string]interface{}{"a": []int{-5, -1, 0, 7}, "b": []int{-3, -1, 2}}, Expected: []int{-5, -3, -1, -1, 0, 2, 7}, Hidden: true, Note: "negative numbers and duplicates across arrays"},
			{Input: map[string]interface{}{"a": []int{1, 3, 5, 7, 9}, "b": []int{2}}, Expected: []int{1, 2, 3, 5, 7, 9}, Hidden: true, Note: "the single item of b lands in the middle, and a has a long tail"},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(n + m) for the result",
		StarterCode:     "def mergeSorted(a, b):\n    # Return a new sorted list with every item of a and b\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "At each step, which single item can be the next one of the merged list? How many items do you have to look at to find it?"},
			{Level: 2, Type: "insight", Content: "Both lists are sorted, so the smallest unused item overall is either the smallest unused item of a or the smallest unused item of b. Keep one pointer in each list (Chapter 3, two arrays) and advance only the pointer whose item you took. When one list is used up, the other's remaining items are already in order.", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Make an empty result and set i and j to 0.\n2. While both pointers are inside their lists, compare a[i] and b[j].\n3. Append the smaller (a[i] on a tie) and advance that pointer.\n4. When the loop ends, at most one list has items left: append the rest of a from i, then the rest of b from j.\n5. Return the result."},
			{Level: 4, Type: "skeleton", Content: "def mergeSorted(a, b):\n    merged = []\n    i = j = 0\n    while i < len(a) and ____:\n        if a[i] <= b[j]:\n            merged.append(a[i])\n            i += 1\n        else:\n            merged.append(b[j])\n            j += 1\n    merged.extend(____)\n    merged.extend(____)\n    return merged"},
		},
		Solution: Solution{
			Code:            "def mergeSorted(a, b):\n    merged = []\n    i = j = 0\n    while i < len(a) and j < len(b):\n        if a[i] <= b[j]:\n            merged.append(a[i])\n            i += 1\n        else:\n            merged.append(b[j])\n            j += 1\n    merged.extend(a[i:])\n    merged.extend(b[j:])\n    return merged",
			Explanation:     "One pointer walks each sorted array, and the smaller of the two front items is always the next answer. It builds on the read pointer from the previous drills, now with two readers and one output. It is also the merge step of merge sort, and Intersection of Two Sorted Arrays, next, uses the same walk but keeps only the matches.",
			BruteForce:      "Concatenate and sort: O(n log n) time for n + m items, which ignores that both inputs are already sorted. Inserting each item of b into a at the right place by scanning is O(n * m).",
			KeyInsight:      "Sorted order means the smallest unused item of each list is its front. So the overall smallest unused item is one of just two candidates, and one comparison decides it.",
			WhyItWorks:      "Invariant: merged holds the i + j smallest items of a and b in order, every unused item (a[i:], b[j:]) is at least as large as merged[-1], and both a[i:] and b[j:] are still sorted. Taking the smaller of a[i] and b[j] gives the smallest unused item overall, because everything else in a[i:] is at least a[i] and everything in b[j:] is at least b[j]. The invariant therefore still holds. When one list is exhausted the other's remaining items are sorted and all at least merged[-1], so they can be appended as they are. Each loop iteration advances i or j, so there are at most n + m iterations.",
			CommonMistakes:  []string{"Stopping when either list runs out and forgetting the leftover tail of the other, so the largest items are lost. Always append a[i:] and b[j:] after the loop.", "Looping `while i < len(a) or j < len(b)` and indexing both lists inside, which reads past the end of the shorter one. Guard each index, or use `and` and finish with the tails.", "Using `<` for one branch and `>` for the other with nothing for equal values, which drops an item whenever the two fronts tie. Every item must be taken: send ties to one side.", "Removing items from the front of a list (`a.pop(0)`) instead of moving an index. Each pop shifts the rest of the list, making the merge O(n^2)."},
			TimeComplexity:  "O(n + m): each item is appended once",
			SpaceComplexity: "O(n + m) for the result; O(1) extra besides it",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up one pointer per array", Explanation: "i is the first unused item of a, j of b, and merged starts empty.", CodeSnippet: "merged = []\ni = j = 0", LineStart: 2, LineEnd: 3},
				{Title: "Compare the two fronts", Explanation: "While both arrays have unused items, the smaller front is the next answer. A tie goes to a.", CodeSnippet: "while i < len(a) and j < len(b):\n    if a[i] <= b[j]:", LineStart: 4, LineEnd: 5},
				{Title: "Take the smaller and advance that pointer", Explanation: "Append the chosen item and move only its pointer. For [1,2,4] and [1,3,4] the first two steps take a's 1 and b's 1.", CodeSnippet: "merged.append(a[i])\ni += 1", LineStart: 6, LineEnd: 7},
				{Title: "Append the leftover tail", Explanation: "One array is used up; the other's rest is sorted and no smaller item can come, so copy it as it is.", CodeSnippet: "merged.extend(a[i:])\nmerged.extend(b[j:])", LineStart: 11, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{Name: "heapq.merge", Code: "import heapq\n\ndef mergeSorted(a, b):\n    return list(heapq.merge(a, b))", TimeComplexity: "O(n + m)", SpaceComplexity: "O(n + m)", Explanation: "The standard library merges already-sorted inputs with the same idea, using a heap to pick the next front item (it generalises to any number of lists). Worth knowing for the k-way merge in the heap chapter."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it enough to compare only a[i] and b[j] and not other items?", Answer: "Each list is sorted, so a[i] is the smallest unused item of a and b[j] the smallest unused item of b. Any other unused item is at least as large as one of those, so it cannot be the next smallest overall."},
			{Prompt: "How would you merge k sorted lists?", Answer: "Keep one pointer per list, but then finding the smallest front takes k comparisons per step. A min-heap holding each list's front gets the smallest in O(log k): the k-way merge from the heap chapter. Merging pairs repeatedly (as merge sort does) also works."},
		},
	},
	{
		ID:              "drill-squares-sorted-array",
		Number:          977,
		Kind:            KindDrill,
		Order:           5,
		Title:           "Squares of a Sorted Array",
		Slug:            "squares-of-a-sorted-array",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Two Pointers", "Array", "Sorting", "Foundations"},
		RelatedChapters: []int{3},
		Description:     "You are given a list `nums` sorted in non-decreasing order, possibly with negative numbers. Return a new list of the **squares** of every item, also sorted in non-decreasing order.\n\nSquaring does not keep the order: `[-4, -1, 0, 3, 10]` squares to `[16, 1, 0, 9, 100]`. But the largest square is always at one of the two **ends**: the item with the biggest absolute value is either the most negative or the most positive. So use opposite-end pointers again, but this time fill the answer **from the back**: compare `abs(nums[left])` with `abs(nums[right])`, put the larger square in the next free slot from the end, and move that pointer inward. Sorting the squares works too, but the point of the drill is the O(n) two-pointer fill.",
		Constraints:     []string{"0 <= nums.length <= 100000", "-10000 <= nums[i] <= 10000", "nums is sorted in non-decreasing order"},
		Examples: []Example{
			{Input: "nums = [-4,-1,0,3,10]", Output: "[0,1,9,16,100]", Explanation: "The ends are -4 and 10: 100 is larger, so it goes last. Next -4 and 3: 16 beats 9, then -1 and 3: 9 beats 1, then 1, then 0."},
			{Input: "nums = [-7,-3,2,3,11]", Output: "[4,9,9,49,121]", Explanation: "11 gives 121 (last), then -7 gives 49, then -3 and 3 tie at 9, then 2 gives 4."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{-4, -1, 0, 3, 10}}, Expected: []int{0, 1, 9, 16, 100}, Note: "example 1: negative and positive items; the biggest square comes from the left end"},
			{Input: map[string]interface{}{"nums": []int{-7, -3, 2, 3, 11}}, Expected: []int{4, 9, 9, 49, 121}, Note: "example 2: 49 from the left end is overtaken by 121 from the right end"},
			{Input: map[string]interface{}{"nums": []int{}}, Expected: []int{}, Hidden: true, Note: "empty list"},
			{Input: map[string]interface{}{"nums": []int{5}}, Expected: []int{25}, Hidden: true, Note: "one item"},
			{Input: map[string]interface{}{"nums": []int{-5}}, Expected: []int{25}, Hidden: true, Note: "one negative item: left and right start on the same index and it must be placed"},
			{Input: map[string]interface{}{"nums": []int{-5, -3, -1}}, Expected: []int{1, 9, 25}, Hidden: true, Note: "all negative: the squares come out in the opposite order of the input"},
			{Input: map[string]interface{}{"nums": []int{0, 2, 4, 9}}, Expected: []int{0, 4, 16, 81}, Hidden: true, Note: "all non-negative: the squares are already sorted"},
			{Input: map[string]interface{}{"nums": []int{-3, 3}}, Expected: []int{9, 9}, Hidden: true, Note: "a tie between the two ends: both squares are 9"},
			{Input: map[string]interface{}{"nums": []int{-6, -6, -2, 1, 1, 6}}, Expected: []int{1, 1, 4, 36, 36, 36}, Hidden: true, Note: "duplicates and ties between the ends in the middle of the run"},
			{Input: map[string]interface{}{"nums": []int{-10, -9, 1, 2, 3}}, Expected: []int{1, 4, 9, 81, 100}, Hidden: true, Note: "the left end wins several times in a row before the right end is used"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n) for the result",
		StarterCode:     "def sortedSquares(nums):\n    # Return the squares of nums as a sorted list\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Which item of the sorted list has the largest square? And once that is placed, which one has the largest square among the rest?"},
			{Level: 2, Type: "insight", Content: "The biggest absolute value is always at the left end or the right end, because the list is sorted. Keep one pointer on each end (opposite ends, Chapter 3), compare absolute values, and place the larger square at the back of the answer; then move that pointer inward and repeat for the next slot from the back.", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Make a result list of the same length as nums.\n2. Set left to the first index and right to the last.\n3. Fill result from the last position down to the first.\n4. Compare abs(nums[left]) and abs(nums[right]); the larger one is squared into the current position.\n5. Move the pointer that was used inward, then return result."},
			{Level: 4, Type: "skeleton", Content: "def sortedSquares(nums):\n    result = [0] * len(nums)\n    left, right = 0, len(nums) - 1\n    for pos in range(len(nums) - 1, -1, -1):\n        if abs(nums[left]) > ____:\n            result[pos] = ____\n            left += 1\n        else:\n            result[pos] = nums[right] * nums[right]\n            right -= 1\n    return result"},
		},
		Solution: Solution{
			Code:            "def sortedSquares(nums):\n    result = [0] * len(nums)\n    left, right = 0, len(nums) - 1\n    for pos in range(len(nums) - 1, -1, -1):\n        if abs(nums[left]) > abs(nums[right]):\n            result[pos] = nums[left] * nums[left]\n            left += 1\n        else:\n            result[pos] = nums[right] * nums[right]\n            right -= 1\n    return result",
			Explanation:     "Reverse a List used both ends to swap; here the two ends compete: the larger absolute value is the largest remaining square, so it fills the answer from the back. It is the opposite-ends pattern with a decision at each step, the same shape as Two Sum II and Container With Most Water in the ladder, which also move one of two end pointers by comparing what they see.",
			BruteForce:      "Square every item and sort: O(n log n) time. It ignores that the input is sorted; the squares of the negative items are in descending order and those of the non-negative items ascending, so the list is a merge of two ordered runs.",
			KeyInsight:      "After squaring, order depends only on absolute value, and the sorted input guarantees the largest absolute value is at one of the two ends. So the largest square is always one comparison away, and filling from the back never has to look inside the list.",
			WhyItWorks:      "Invariant: result[pos + 1:] holds the squares of the items already consumed (those outside nums[left..right]), in sorted order, and every one of them is at least as large as the square of anything inside the window. The window nums[left..right] is sorted, so the largest absolute value inside it is at left or at right; squaring it and placing it at result[pos] keeps the invariant, and the other end becomes the candidate for the next slot. Each iteration consumes one item, so after n iterations the whole list is filled and sorted.",
			CommonMistakes:  []string{"Comparing the raw values (`nums[left] > nums[right]`) instead of their absolute values, so a large negative item is never chosen. Compare `abs(...)` (or the squares).", "Filling the result from the front, so the largest squares come first and the answer is in descending order. The larger square goes into the last free slot.", "Looping `while left < right` and leaving the final item unplaced when the pointers meet. When left equals right there is still one item whose square belongs in the last free slot, so drive the loop by the number of slots, or use `<=`.", "Squaring the items but returning them in the input order, which is only sorted when every item is non-negative."},
			TimeComplexity:  "O(n): each item is squared and placed once",
			SpaceComplexity: "O(n): the result list",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the answer and the two ends", Explanation: "The result is filled from the back, so one pointer sits on each end of the sorted input.", CodeSnippet: "result = [0] * len(nums)\nleft, right = 0, len(nums) - 1", LineStart: 2, LineEnd: 3},
				{Title: "Fill the slots from the back", Explanation: "pos counts down from the last index to 0, so each iteration decides the largest square still unplaced.", CodeSnippet: "for pos in range(len(nums) - 1, -1, -1):", LineStart: 4, LineEnd: 4},
				{Title: "Compare the ends by absolute value", Explanation: "The end farther from zero has the larger square. For [-4,-1,0,3,10], 4 against 10: the right end wins.", CodeSnippet: "if abs(nums[left]) > abs(nums[right]):", LineStart: 5, LineEnd: 5},
				{Title: "Place it and move that pointer inward", Explanation: "Write the square, then step the pointer used toward the middle. The other pointer stays, because its item has not been placed.", CodeSnippet: "result[pos] = nums[left] * nums[left]\nleft += 1", LineStart: 6, LineEnd: 7},
			},
			Alternatives: []AltSolution{
				{Name: "Square and sort", Code: "def sortedSquares(nums):\n    return sorted(x * x for x in nums)", TimeComplexity: "O(n log n)", SpaceComplexity: "O(n)", Explanation: "The simplest answer and a fine first attempt. It re-sorts data that already has structure, so it costs a log n factor the two-pointer fill avoids."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can the largest square never be in the middle of the list?", Answer: "Squares grow with absolute value. In a sorted list the absolute value is largest at one of the two ends: moving inward from either end goes toward zero, or toward the other end's smaller side. So the maximum is always at left or right."},
			{Prompt: "What changes if nums is not sorted?", Answer: "Then the largest absolute value could be anywhere, and there is no two-pointer shortcut: square everything and sort (O(n log n))."},
		},
	},
	{
		ID:              "drill-is-subsequence",
		Number:          392,
		Kind:            KindDrill,
		Order:           6,
		Title:           "Is Subsequence",
		Slug:            "is-subsequence",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Two Pointers", "String", "Two Arrays", "Foundations"},
		RelatedChapters: []int{3},
		Description:     "A string `s` is a **subsequence** of a string `t` if you can get `s` by deleting some (possibly none) of the characters of `t` without changing the order of the rest. For example `\"ace\"` is a subsequence of `\"abcde\"`, but `\"aec\"` is not, because the `e` comes after the `c` in `t`. Return `true` if `s` is a subsequence of `t`, otherwise `false`. The empty string is a subsequence of every string.\n\nThis is the **two sequences** shape again, with a twist: the pointers have different jobs. `j` walks along `t` one step at a time. `i` marks the next character of `s` that still has to be found, and it moves only when `t[j]` matches it. `s` is a subsequence exactly when `i` gets all the way to the end of `s` before `t` runs out.",
		Constraints:     []string{"0 <= s.length <= 100", "0 <= t.length <= 10000", "s and t consist of lowercase English letters"},
		Examples: []Example{
			{Input: "s = \"abc\", t = \"ahbgdc\"", Output: "true", Explanation: "j finds a at 0, b at 2 and c at 5 in that order; the letters in between are skipped."},
			{Input: "s = \"axc\", t = \"ahbgdc\"", Output: "false", Explanation: "a matches, but there is no x in t, so i is stuck at 1 when t runs out."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "abc", "t": "ahbgdc"}, Expected: true, Note: "example 1: a, b and c appear in t in that order, with other letters between"},
			{Input: map[string]interface{}{"s": "axc", "t": "ahbgdc"}, Expected: false, Note: "example 2: there is no x in t"},
			{Input: map[string]interface{}{"s": "", "t": "ahbgdc"}, Expected: true, Hidden: true, Note: "the empty string is a subsequence of anything"},
			{Input: map[string]interface{}{"s": "", "t": ""}, Expected: true, Hidden: true, Note: "empty s and empty t"},
			{Input: map[string]interface{}{"s": "a", "t": ""}, Expected: false, Hidden: true, Note: "a non-empty s is never a subsequence of an empty t"},
			{Input: map[string]interface{}{"s": "aa", "t": "a"}, Expected: false, Hidden: true, Note: "t has the letter, but only once: a match uses up a letter of t"},
			{Input: map[string]interface{}{"s": "abc", "t": "cba"}, Expected: false, Hidden: true, Note: "same letters, wrong order: this is not a subsequence"},
			{Input: map[string]interface{}{"s": "ace", "t": "abcde"}, Expected: true, Hidden: true, Note: "skipping letters of t is allowed"},
			{Input: map[string]interface{}{"s": "abc", "t": "abc"}, Expected: true, Hidden: true, Note: "s equal to t"},
			{Input: map[string]interface{}{"s": "ab", "t": "baab"}, Expected: true, Hidden: true, Note: "the first b cannot be used, but a later a and b can"},
			{Input: map[string]interface{}{"s": "abcd", "t": "abc"}, Expected: false, Hidden: true, Note: "s is longer than t"},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def isSubsequence(s, t):\n    # Return True if s is a subsequence of t\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you read t from left to right, which character of s are you waiting for at any moment? What should happen when you meet it, and when you meet any other character?"},
			{Level: 2, Type: "insight", Content: "Greedy matching works: take every match as early as possible. Keep one pointer j that scans t and one pointer i at the next unmatched character of s (Chapter 3, two sequences). j moves every step; i moves only on a match. s is a subsequence when i reaches len(s).", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Start i (position in s) and j (position in t) at 0.\n2. Loop while both are inside their strings.\n3. If s[i] equals t[j], advance i.\n4. Always advance j.\n5. After the loop, return whether i reached the end of s."},
			{Level: 4, Type: "skeleton", Content: "def isSubsequence(s, t):\n    i = j = 0\n    while i < len(s) and j < len(t):\n        if s[i] == ____:\n            i += 1\n        ____\n    return i == ____"},
		},
		Solution: Solution{
			Code:            "def isSubsequence(s, t):\n    i = j = 0\n    while i < len(s) and j < len(t):\n        if s[i] == t[j]:\n            i += 1\n        j += 1\n    return i == len(s)",
			Explanation:     "Two pointers walk two different sequences at different speeds: j always moves, i moves only on a match. Matching as early as possible is safe, so a single left-to-right pass decides the answer. It is the two-sequence shape of Merge Two Sorted Arrays with a different move rule, and the same idea reappears in string problems such as Merge Strings and the sliding window chapter.",
			BruteForce:      "Try every way of choosing len(s) positions of t in order (or recursion with two choices at each character): exponential in the worst case. Checking only that every character of s occurs in t is quick but answers a different question, since it ignores order.",
			KeyInsight:      "Greedy matching is safe: if s[i] can be matched by t[j], matching it now leaves at least as much of t for the rest of s as matching it later would. So one pass with no backtracking decides the answer.",
			WhyItWorks:      "Invariant: i is the length of the longest prefix of s that is a subsequence of t[:j]. When t[j] equals s[i], that prefix grows to i + 1 (taking the earliest possible t[j] leaves every later character of t available), otherwise it is unchanged. j grows by one each step, so the loop ends after at most len(t) steps. s is a subsequence of t exactly when its longest matched prefix is all of s, that is, when i == len(s).",
			CommonMistakes:  []string{"Checking only that each character of s appears somewhere in t (a set or a count), which ignores order: `abc` would be a subsequence of `cba`. The characters must be found left to right.", "Advancing j only on a mismatch, so a character of t can match several characters of s: `aa` would be a subsequence of `a`. j moves on every step.", "Reading `s[0]` before checking that s is not empty, which crashes on the empty string. Let the loop condition guard the indexing.", "Requiring the characters to be next to each other (a substring test such as `s in t`). Letters of t may be skipped in between."},
			TimeComplexity:  "O(n + m): j moves at most len(t) times",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Place a pointer in each string", Explanation: "i is the character of s still to be found, j the current character of t. Both start at the front.", CodeSnippet: "i = j = 0", LineStart: 2, LineEnd: 2},
				{Title: "Scan t while s is unfinished", Explanation: "Stop as soon as s is fully matched or t has no characters left.", CodeSnippet: "while i < len(s) and j < len(t):", LineStart: 3, LineEnd: 3},
				{Title: "Advance s only on a match", Explanation: "A match uses t[j] for s[i]. For \"abc\" in \"ahbgdc\", j = 0 matches the a and i becomes 1.", CodeSnippet: "if s[i] == t[j]:\n    i += 1", LineStart: 4, LineEnd: 5},
				{Title: "Always advance t", Explanation: "Whether or not it matched, t[j] has been used or skipped; the next step looks at the next character.", CodeSnippet: "j += 1", LineStart: 6, LineEnd: 6},
				{Title: "Decide from i", Explanation: "If i reached len(s), every character of s was found in order.", CodeSnippet: "return i == len(s)", LineStart: 7, LineEnd: 7},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to match each character of s at the earliest position in t?", Answer: "Any later match of the same character would leave a shorter suffix of t for the rest of s. Whatever characters can be matched after a later match can also be matched after the earlier one, so the greedy choice never loses a solution."},
			{Prompt: "What changes if you must answer this for thousands of different s against the same long t?", Answer: "The linear scan for each s is too slow. Preprocess t once into, for each letter, the sorted list of its positions, and find each next match with binary search (Chapter 6): O(len(s) log len(t)) per query."},
		},
	},
	{
		ID:              "drill-intersect-sorted-arrays",
		Number:          350,
		Kind:            KindDrill,
		Order:           7,
		Title:           "Intersection of Two Sorted Arrays",
		Slug:            "intersection-of-two-arrays-ii",
		Difficulty:      "Easy",
		Category:        "two-pointers",
		Tags:            []string{"Two Pointers", "Array", "Two Arrays", "Foundations"},
		RelatedChapters: []int{3},
		Description:     "You are given two lists `nums1` and `nums2`, **each sorted in non-decreasing order**. Return a sorted list of their intersection: each value that occurs in both lists, repeated as many times as it occurs in **both** (that is, the smaller of its two counts). For example `[1,1,2,2]` and `[2,2]` give `[2,2]`, and `[1,1,1]` with `[1,1]` gives `[1,1]`. (LeetCode's Intersection of Two Arrays II allows unsorted input; here both are sorted, which is what lets two pointers replace the hash map.)\n\nWalk one pointer along each list. If the two items differ, the smaller one cannot match anything further in the other list, so advance **that** pointer only. If they are equal, you have found a common value: record it and advance **both** pointers, so that each copy is matched with exactly one copy on the other side.",
		Constraints:     []string{"0 <= nums1.length, nums2.length <= 1000", "-1000 <= nums1[i], nums2[j] <= 1000", "Both lists are sorted in non-decreasing order"},
		Examples: []Example{
			{Input: "nums1 = [1,1,2,2], nums2 = [2,2]", Output: "[2,2]", Explanation: "1 < 2 so the first pointer skips both 1s; then 2 = 2 twice, so two 2s are matched."},
			{Input: "nums1 = [4,5,9], nums2 = [4,4,8,9,9]", Output: "[4,9]", Explanation: "4 matches 4 (once, nums1 has only one 4); 5 is skipped, 8 is skipped, 9 matches 9."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums1": []int{1, 1, 2, 2}, "nums2": []int{2, 2}}, Expected: []int{2, 2}, Note: "example 1: 2 appears twice in both, so it is returned twice"},
			{Input: map[string]interface{}{"nums1": []int{4, 5, 9}, "nums2": []int{4, 4, 8, 9, 9}}, Expected: []int{4, 9}, Note: "example 2: 4 and 9 are in both; each appears once in nums1, so once in the answer"},
			{Input: map[string]interface{}{"nums1": []int{}, "nums2": []int{1, 2}}, Expected: []int{}, Hidden: true, Note: "first array empty"},
			{Input: map[string]interface{}{"nums1": []int{1, 2}, "nums2": []int{}}, Expected: []int{}, Hidden: true, Note: "second array empty"},
			{Input: map[string]interface{}{"nums1": []int{1, 2, 3}, "nums2": []int{4, 5, 6}}, Expected: []int{}, Hidden: true, Note: "no common value: the answer is empty"},
			{Input: map[string]interface{}{"nums1": []int{1, 2, 3}, "nums2": []int{2, 3, 4}}, Expected: []int{2, 3}, Hidden: true, Note: "the pointers must not both skip on a mismatch: 2 and 3 are common"},
			{Input: map[string]interface{}{"nums1": []int{1, 1, 1}, "nums2": []int{1, 1}}, Expected: []int{1, 1}, Hidden: true, Note: "a value that repeats: the answer has the smaller count, two 1s"},
			{Input: map[string]interface{}{"nums1": []int{-3, -1, 0, 2, 2}, "nums2": []int{-3, -3, 2, 2, 2}}, Expected: []int{-3, 2, 2}, Hidden: true, Note: "negative numbers and different counts on each side"},
			{Input: map[string]interface{}{"nums1": []int{5, 5, 5}, "nums2": []int{5, 5, 5}}, Expected: []int{5, 5, 5}, Hidden: true, Note: "identical arrays"},
			{Input: map[string]interface{}{"nums1": []int{1, 2, 2, 3}, "nums2": []int{2, 3, 3}}, Expected: []int{2, 3}, Hidden: true, Note: "a run in one array but not the other: one 2 and one 3 are common"},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(1) besides the answer",
		StarterCode:     "def intersect(nums1, nums2):\n    # Return the sorted intersection, with repeated values\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If the two items you are looking at are different, which of them can you safely discard? What do you do when they are equal?"},
			{Level: 2, Type: "insight", Content: "In sorted lists, if nums1[i] < nums2[j], then nums1[i] is smaller than everything left in nums2, so it can never match: advance i. The other way round advances j. On equality, record the value and advance both, so each copy matches one copy (Chapter 3, two arrays).", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Start i and j at 0 and an empty answer.\n2. Loop while both pointers are inside their lists.\n3. If nums1[i] is smaller, advance i. If nums2[j] is smaller, advance j.\n4. If they are equal, append the value and advance both.\n5. Return the answer."},
			{Level: 4, Type: "skeleton", Content: "def intersect(nums1, nums2):\n    common = []\n    i = j = 0\n    while i < len(nums1) and j < len(nums2):\n        if nums1[i] < nums2[j]:\n            i += 1\n        elif ____:\n            j += 1\n        else:\n            common.append(nums1[i])\n            ____\n            ____\n    return common"},
		},
		Solution: Solution{
			Code:            "def intersect(nums1, nums2):\n    common = []\n    i = j = 0\n    while i < len(nums1) and j < len(nums2):\n        if nums1[i] < nums2[j]:\n            i += 1\n        elif nums1[i] > nums2[j]:\n            j += 1\n        else:\n            common.append(nums1[i])\n            i += 1\n            j += 1\n    return common",
			Explanation:     "The same two-array walk as Merge Two Sorted Arrays, but instead of taking the smaller item every time it discards the smaller item unless it matches. Sorted order tells each pointer which side has fallen behind. It builds on the merge drill's loop and is the pointer-movement logic that 3Sum (15) and Two Sum II (167) reuse.",
			BruteForce:      "Count the values of one list in a hash map and look up each item of the other (O(n + m) time but O(n) space), or compare every pair (O(n * m)). Neither uses the sorted order, which lets two pointers do it in O(1) extra space.",
			KeyInsight:      "In sorted lists the smaller front item is smaller than everything remaining on the other side, so a mismatch always tells you which pointer to move. A match consumes one copy from each list, which gives the smaller of the two counts automatically.",
			WhyItWorks:      "Invariant: common holds the intersection (with multiplicities) of nums1[:i] and nums2[:j], and every item of nums1[:i] and nums2[:j] that is not in common cannot match anything in the remaining parts. If nums1[i] < nums2[j], all of nums2[j:] is at least nums2[j] > nums1[i], so nums1[i] has no partner in nums2[j:] (and the earlier part was already handled): advancing i loses nothing. The symmetric case advances j. If they are equal, pairing these two copies is always right, since either copy could only be paired with an equal item, and it is consumed on both sides. A value with count p in one list and q in the other is matched min(p, q) times before one list runs out of it.",
			CommonMistakes:  []string{"Advancing only one pointer after a match, so the same item is matched against several copies on the other side and the answer repeats it too often (`[1,1,1]` against `[1]` should give one 1). A match uses one copy from each list: advance both.", "Advancing both pointers on a mismatch, which can skip an item that matches the next one: `[1,2,3]` and `[2,3,4]` would give an empty answer. Advance only the pointer with the smaller item.", "Using a set intersection, which drops repeated values: `[2,2]` and `[2,2]` would give `[2]`.", "Looping until both lists are exhausted and indexing past the end of one of them. Stop as soon as either pointer runs out: nothing more can match."},
			TimeComplexity:  "O(n + m)",
			SpaceComplexity: "O(1) besides the answer",
			Walkthrough: []WalkthroughStep{
				{Title: "Place a pointer in each list", Explanation: "i and j are the first unused items of the two sorted lists.", CodeSnippet: "common = []\ni = j = 0", LineStart: 2, LineEnd: 3},
				{Title: "Loop while both have items", Explanation: "Once either list is used up there is nothing left to match against.", CodeSnippet: "while i < len(nums1) and j < len(nums2):", LineStart: 4, LineEnd: 4},
				{Title: "Discard the smaller item", Explanation: "If nums1[i] is smaller, it is smaller than everything left in nums2, so it can never match. In [4,5,9] against [4,4,8,9,9], 5 < 8 is discarded.", CodeSnippet: "if nums1[i] < nums2[j]:\n    i += 1\nelif nums1[i] > nums2[j]:\n    j += 1", LineStart: 5, LineEnd: 8},
				{Title: "Record a match and advance both", Explanation: "The equal items are one copy each. Take them as a pair and move on, so the next copies are compared with each other.", CodeSnippet: "common.append(nums1[i])\ni += 1\nj += 1", LineStart: 10, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{Name: "Count with a Counter", Code: "from collections import Counter\n\ndef intersect(nums1, nums2):\n    counts = Counter(nums1)\n    common = []\n    for x in nums2:\n        if counts[x] > 0:\n            common.append(x)\n            counts[x] -= 1\n    return common", TimeComplexity: "O(n + m)", SpaceComplexity: "O(n)", Explanation: "A hash map of counts works for unsorted input as well. Because nums2 is sorted here, the answer comes out sorted too. It spends O(n) space where two pointers need none."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to discard the smaller item when the two fronts differ?", Answer: "Everything left in the other list is at least as large as its front, which is larger than the smaller item. So the smaller item has no equal partner left, and moving past it loses nothing."},
			{Prompt: "What changes if the arrays are not sorted?", Answer: "The pointers can no longer tell which side is behind. Sort both first (O(n log n)) and then use this loop, or use a hash map of counts."},
		},
	},
	{
		ID:              "drill-count-pairs-sorted",
		Number:          0,
		Kind:            KindDrill,
		Order:           8,
		Title:           "Count Pairs with a Target Sum",
		Slug:            "count-pairs-with-target-sum",
		Difficulty:      "Medium",
		Category:        "two-pointers",
		Tags:            []string{"Two Pointers", "Array", "Duplicates", "Counting", "Foundations"},
		RelatedChapters: []int{3},
		Description:     "You are given a list `nums` sorted in non-decreasing order, which **may contain duplicates**, and an integer `target`. Return how many pairs of **positions** `(i, j)` with `i < j` have `nums[i] + nums[j] == target`. Two pairs count separately if they use different positions, even when the values are the same: in `[1,1,3]` with target 4 there are two pairs, `(0,2)` and `(1,2)`.\n\nStart as in Two Sum II: `left` at the first item and `right` at the last. If the sum is too small, move `left` right; if too large, move `right` left. The new part is the case where the sum equals the target. Do not move just one step: count every pair at once. If `nums[left] != nums[right]`, then `left` begins a run of `low_count` equal items and `right` ends a run of `high_count` equal items, and each item of one run pairs with each of the other: add `low_count * high_count`, then jump both pointers past their runs. If `nums[left] == nums[right]`, every item between them is the same value, and the number of pairs among `width` items is `width * (width - 1) / 2`; add that and stop. This duplicate skipping is the same step that 3Sum needs.",
		Constraints:     []string{"0 <= nums.length <= 100000", "-1000000000 <= nums[i], target <= 1000000000", "nums is sorted in non-decreasing order"},
		Examples: []Example{
			{Input: "nums = [1,1,2,3,3], target = 4", Output: "4", Explanation: "The 1 run (two items) pairs with the 3 run (two items): 2 * 2 = 4 pairs. The single 2 would need another 2."},
			{Input: "nums = [2,2,2,2], target = 4", Output: "6", Explanation: "Every pair of positions sums to 4. There are 4 choose 2 = 6 of them, which is 4 * 3 / 2."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}, "target": 5}, Expected: 2, Note: "the pairs (1,4) and (2,3): no duplicates, the plain two-pointer walk"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1}, "target": 2}, Expected: 3, Note: "all equal: three pairs of positions (0,1), (0,2), (1,2)"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 2, 2, 3, 3}, "target": 4}, Expected: 5, Hidden: true, Note: "two copies of 1 and of 3 give 4 pairs, and (2,2) gives 1 more: 5"},
			{Input: map[string]interface{}{"nums": []int{}, "target": 5}, Expected: 0, Hidden: true, Note: "empty list"},
			{Input: map[string]interface{}{"nums": []int{3}, "target": 6}, Expected: 0, Hidden: true, Note: "one item cannot make a pair, even though 3 + 3 = 6"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 2, 2}, "target": 4}, Expected: 6, Hidden: true, Note: "four equal items: choose 2 of 4 = 6 (4 * 4 would overcount)"},
			{Input: map[string]interface{}{"nums": []int{1, 2}, "target": 10}, Expected: 0, Hidden: true, Note: "no pair reaches the target"},
			{Input: map[string]interface{}{"nums": []int{-2, -1, 0, 1, 2}, "target": 0}, Expected: 2, Hidden: true, Note: "negative numbers: (-2,2) and (-1,1)"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 2, 3, 3}, "target": 4}, Expected: 4, Hidden: true, Note: "the 1-run and the 3-run multiply: 2 * 2 = 4; the single 2 has no partner"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 3, 3}, "target": 4}, Expected: 3, Hidden: true, Note: "one 1 with three 3s: 1 * 3 = 3"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0, 0, 5}, "target": 5}, Expected: 4, Hidden: true, Note: "four zeros each pair with the 5"},
			{Input: map[string]interface{}{"nums": tpDrillAllSame(50000, 7), "target": 14}, Expected: 1249975000, Hidden: true, Timeout: 5000, Note: "50000 equal items: the answer is 50000 * 49999 / 2; needs O(n), and checking every pair runs out of time"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def countPairs(nums, target):\n    # Return the number of index pairs i < j with nums[i] + nums[j] == target\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When the sum hits the target, there may be several equal items next to each of the two pointers. How many pairs do they make together, and where should the pointers go afterwards?"},
			{Level: 2, Type: "insight", Content: "Move the pointers as in Two Sum II (Chapter 3): too small, advance left; too large, retreat right. On a hit, count the whole run of equal values on each side at once: different values give low_count * high_count pairs, and when left and right point at equal values, all the items between them are equal and make width * (width - 1) / 2 pairs. Then skip past the runs.", ChapterRef: 3},
			{Level: 3, Type: "plan", Content: "1. Set left to the start, right to the end, count to 0.\n2. While left < right, compute the sum.\n3. Sum below the target: left += 1. Sum above the target: right -= 1.\n4. Equal and nums[left] == nums[right]: add width * (width - 1) // 2 and stop.\n5. Equal and the values differ: measure the run at each pointer, add low_count * high_count, and move both pointers past their runs.\n6. Return count."},
			{Level: 4, Type: "skeleton", Content: "def countPairs(nums, target):\n    count = 0\n    left, right = 0, len(nums) - 1\n    while left < right:\n        total = nums[left] + nums[right]\n        if total < target:\n            left += 1\n        elif total > target:\n            right -= 1\n        elif nums[left] == nums[right]:\n            width = right - left + 1\n            count += ____\n            break\n        else:\n            low, high = nums[left], nums[right]\n            low_count = high_count = 0\n            while nums[left] == low:\n                left += 1\n                low_count += 1\n            while nums[right] == high:\n                right -= 1\n                high_count += 1\n            count += ____\n    return count"},
		},
		Solution: Solution{
			Code:            "def countPairs(nums, target):\n    count = 0\n    left, right = 0, len(nums) - 1\n    while left < right:\n        total = nums[left] + nums[right]\n        if total < target:\n            left += 1\n        elif total > target:\n            right -= 1\n        elif nums[left] == nums[right]:\n            width = right - left + 1\n            count += width * (width - 1) // 2\n            break\n        else:\n            low, high = nums[left], nums[right]\n            low_count = high_count = 0\n            while nums[left] == low:\n                left += 1\n                low_count += 1\n            while nums[right] == high:\n                right -= 1\n                high_count += 1\n            count += low_count * high_count\n    return count",
			Explanation:     "Two Sum II (167) finds one pair with opposite-end pointers; this drill counts every pair, so on a match it must handle runs of equal values instead of stepping one position. That run-skipping is exactly the duplicate handling of 3Sum (15), which fixes one item and then runs this walk on the rest. It builds on Squares of a Sorted Array for the opposite-ends loop and on Remove Duplicates for the idea of passing over a run.",
			BruteForce:      "Check every pair of positions: O(n^2) time, which on 50,000 items is over a billion steps. A hash map of counts runs in O(n) but needs O(n) space and does not teach the sorted-array walk that 3Sum builds on.",
			KeyInsight:      "Because the array is sorted, equal values form runs, and every item in one run behaves identically against the other end. So a hit does not need to be counted pair by pair: the run lengths give the whole count in one step, and the pointers can jump past both runs.",
			WhyItWorks:      "Invariant: count is the number of valid pairs among the positions already discarded (those before left or after right), and no valid pair joins a discarded position with one still inside the window nums[left..right]. If the sum is too small, nums[left] with anything at or before right is too small (they are all at most nums[right]), so left has no partner and is discarded; symmetrically for too large. On a hit with different values, the run of low (all items from left while equal) and the run of high (down from right while equal) pair completely with each other: each of the low_count * high_count pairs sums to the target, and no item of either run can pair with anything else, because its only partner value is the other run's value. Both runs are discarded. If the values are equal, all items in the window are the same value v and 2v = target, so every pair of positions in the window counts: width * (width - 1) / 2. Each step discards at least one item, so the loop is O(n).",
			CommonMistakes:  []string{"Moving both pointers by one on a hit (as Two Sum II does), which counts one pair per hit and misses the other copies: `[1,1,3,3]` with target 4 would give 1 instead of 4.", "Multiplying the two run lengths when both pointers sit in the same run of equal values. The runs overlap there, so low_count * high_count overcounts: four 2s give 6 pairs, not 4 * 4. Count width * (width - 1) / 2 instead.", "Counting distinct value pairs instead of pairs of positions, so `[1,1,1]` with target 2 gives 1 rather than 3.", "Checking every pair of positions, which is correct but O(n^2): on 50000 equal items it runs out of time.", "Walking a run with `while nums[left + 1] == nums[left]` and no bound, which reads past the end. The run loop here is safe because the other pointer's value differs, so it stops before reaching it."},
			TimeComplexity:  "O(n): every step discards at least one position",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Place the two pointers at the ends", Explanation: "left starts at the smallest item and right at the largest, as in Two Sum II.", CodeSnippet: "count = 0\nleft, right = 0, len(nums) - 1", LineStart: 2, LineEnd: 3},
				{Title: "Move the pointer that is wrong", Explanation: "A sum below the target needs a larger left item; above it needs a smaller right item. Neither discarded item could have worked.", CodeSnippet: "if total < target:\n    left += 1\nelif total > target:\n    right -= 1", LineStart: 6, LineEnd: 9},
				{Title: "Count a run of one repeated value", Explanation: "Equal values at both pointers means the whole window is one value, so every pair of its items counts, and there is nothing left to find.", CodeSnippet: "width = right - left + 1\ncount += width * (width - 1) // 2\nbreak", LineStart: 11, LineEnd: 13},
				{Title: "Count two runs at once and skip them", Explanation: "The run starting at left and the run ending at right pair completely. For [1,1,2,3,3] with target 4: 2 and 2 items give 4 pairs. Both runs are then stepped over.", CodeSnippet: "low, high = nums[left], nums[right]\nlow_count = high_count = 0\nwhile nums[left] == low:\n    left += 1\n    low_count += 1\nwhile nums[right] == high:\n    right -= 1\n    high_count += 1\ncount += low_count * high_count", LineStart: 15, LineEnd: 23},
			},
			Alternatives: []AltSolution{
				{Name: "Hash map of counts", Code: "from collections import Counter\n\ndef countPairs(nums, target):\n    seen = Counter()\n    count = 0\n    for x in nums:\n        count += seen[target - x]\n        seen[x] += 1\n    return count", TimeComplexity: "O(n)", SpaceComplexity: "O(n)", Explanation: "For each item, add how many earlier items complete it, then record the item. It works without the array being sorted (Chapter 2), at the cost of O(n) space."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to skip the whole run at left after a hit?", Answer: "Every item in that run has the same value, so it needs the same partner value, and the partner run at right has already been counted against all of them. No item of the run can pair with anything else, because the target and its own value fix the partner's value."},
			{Prompt: "How does this help with 3Sum?", Answer: "3Sum fixes the first item and then needs the pairs among the remaining sorted items that sum to minus that item, listing each distinct value pair once. It uses the same walk, but on a hit it records the pair and skips equal values to avoid repeated triples. Counting runs here is the same skipping step."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, TwoPointersDrills...)
}

// tpDrillZeroEveryOther returns n items where the even indexes are 0 and the
// odd indexes hold i.
func tpDrillZeroEveryOther(n int) []int {
	out := make([]int, n)
	for i := range out {
		if i%2 == 1 {
			out[i] = i
		}
	}
	return out
}

// tpDrillZeroEveryOtherAnswer is tpDrillZeroEveryOther(n) with the zeros moved
// to the end.
func tpDrillZeroEveryOtherAnswer(n int) []int {
	out := make([]int, 0, n)
	for i := 1; i < n; i += 2 {
		out = append(out, i)
	}
	for len(out) < n {
		out = append(out, 0)
	}
	return out
}

// tpDrillAllSame returns n copies of v.
func tpDrillAllSame(n, v int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = v
	}
	return out
}
