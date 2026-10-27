package chapter

// chapter03 teaches the two-pointer technique for the two-pointers ladder.
// Storyboards (IDs registered in app.js): none.
func chapter03() Chapter {
	return Chapter{
		Number:        3,
		Slug:          "two-pointers",
		Title:         "Two Pointers",
		Category:      "two-pointers",
		Prerequisites: []int{2},
		Summary: "Two indices moving through an array can replace a loop inside a loop. Learn why, in sorted input or under the right " +
			"argument, each comparison proves one element useless so a pointer can safely move past it, and use it for pair sums, " +
			"palindromes, in-place rewriting and two-wall area problems, all in O(n) time and O(1) extra memory.",
		Motivation: `Find two numbers that add up to a target in a sorted list of 100,000 numbers, using no extra memory. Chapter 2's hash map is out (it needs O(n) memory), and the brute force checks about 5 billion pairs.

Watch what one comparison can tell you. Take [1, 3, 4, 6, 8, 11] and target 10, and add the two ends: 1 + 11 = 12, too big. 1 is the smallest value, so 11 is too big even with the smallest possible partner: 11 is in no valid pair, and you can throw it away. Next, 1 + 8 = 9, too small, and 8 is now the largest value left, so 1 is too small even with the largest partner: throw 1 away. 3 + 8 = 11, too big: drop 8. 3 + 6 = 9, too small: drop 3. 4 + 6 = 10: found.

Five comparisons instead of the 15 pairs. Each comparison discarded one value for good, so for n = 100,000 you need at most 99,999 comparisons, not 5 billion, and only two index variables. That is the two-pointer idea: keep two positions, and use what you know about the data to prove that one of them can move.`,
		Objectives: []string{
			"Explain, with a discard argument, why moving one pointer in a sorted pair search can never skip the answer.",
			"Implement opposite-end pointers for pair sums and palindromes, and same-direction (read/write) pointers for rewriting an array in place.",
			"Turn a three-number search into a pair search by sorting, fixing one number and skipping duplicates.",
			"Decide between a hash map (Chapter 2) and two pointers based on whether the input is sorted, whether indices are needed and how much memory is allowed.",
			"Trace the pointers on a small input and state the invariant the loop keeps.",
		},
		WhenToUse: []string{
			"The input is sorted, or you're allowed to sort it, and you need a pair or triplet with a given sum or closest to a target.",
			"You compare an array or string from both ends: palindromes, reversing, mirrored checks.",
			"The problem requires O(1) extra memory, which rules out a set or hash map.",
			"You must change an array in place: remove duplicates or a value, move zeros, compact the kept items to the front.",
			"You choose a left and a right boundary and the answer is limited by the smaller of the two (water between walls).",
		},
		Concepts: []Concept{
			{
				Name:        "Opposite-end pointers",
				Description: "Start one pointer at each end and move them toward each other; each step uses a comparison to discard one end.",
				WhyItWorks: `In a sorted array, nums[left] + nums[right] is the largest sum nums[left] can make with anything still in the window (its partner is the biggest value left) and the smallest sum nums[right] can make (its partner is the smallest value left). So if the sum is too small, nums[left] is too small with every remaining partner and can be discarded; if it is too big, nums[right] is too big with every remaining partner and can be discarded.

The invariant: if an answer pair exists, both of its indices are still inside [left, right]. Each move only discards an index proven to be in no answer, so the invariant survives every step. Each step shrinks the window by one, so the loop ends after at most n - 1 steps: O(n) time and O(1) extra memory. For a palindrome the argument is simpler: position i from the left must match position i from the right, so once a mirror pair matches, both ends are settled and both pointers move.`,
				Intuition: "You have $100 to spend on exactly two items from a shelf sorted by price. Hold the cheapest item in your left hand and the most expensive in your right. Over budget? The expensive item can't work even with the cheapest partner, so put it back and pick up the next most expensive. Under budget? The cheap item can't reach $100 even with the priciest partner, so put it back and pick the next cheapest. Each put-back is permanent, and your hands meet after at most one pass along the shelf.",
				CoreIdeas: []string{
					"Place the pointers at index 0 and len(nums) - 1.",
					"Compare and decide: the sum (or the two characters) tells you what to do.",
					"Move the pointer that can't help: too small -> left += 1, too big -> right -= 1.",
					"Record the answer when you find it, or keep the best seen so far.",
					"Stop when they meet (left >= right): every element has been settled.",
				},
				CommonMistakes: []string{
					"Using opposite-end pointers on unsorted input for a sum problem. The discard argument depends on sorted order; sort first, or use a hash map (Chapter 2) if you need the original indices.",
					"Moving the wrong pointer, or both, after a miss. Only the pointer proven useless may move; moving the other can skip the answer.",
					"Writing while left <= right for a pair search, which lets an element pair with itself when left == right. Use left < right.",
				},
				Examples: []CodeExample{
					{
						ID:       "pair-sum-sorted",
						Title:    "Pair with a given sum in a sorted list",
						Language: "python",
						Snippet: `def has_pair_with_sum(nums, target):
    left, right = 0, len(nums) - 1
    while left < right:
        total = nums[left] + nums[right]
        if total == target:
            return True
        if total < target:
            left += 1    # nums[left] is too small even with the largest partner
        else:
            right -= 1   # nums[right] is too big even with the smallest partner
    return False`,
						Notes: "On [1, 3, 4, 6, 8, 11] with target 10 the pointers visit (1,11) (1,8) (3,8) (3,6) (4,6): five comparisons.",
					},
				},
			},
			{
				Name:        "Same-direction (read/write) pointers",
				Description: "A read pointer scans every element; a write pointer marks where the next element you keep should go. Used to filter or compact an array in place.",
				WhyItWorks: `The invariant: nums[0:write] holds exactly the elements kept so far, in their original order, and write <= read. The read pointer looks at each element once. If the element should be kept, it is copied to nums[write] and write advances; otherwise only read advances. Because write never passes read, copying never overwrites an element that hasn't been read yet.

When read reaches the end, the first write positions are the answer, and everything after them is leftover that the caller ignores. Each element is read once and written at most once: O(n) time, O(1) extra memory. The keep rule depends on the problem: "differs from the last kept value" removes duplicates from a sorted array, "is not zero" moves zeros to the back, "is not val" removes a value.`,
				Intuition: "A librarian shelving books from a cart. One hand picks up each book in turn (read). The other marks the next free spot on the shelf (write). A book that belongs is placed at the mark and the mark moves along; a book that doesn't is set aside and the mark stays put. The shelf fills from the left in cart order, and the reading hand is never behind the shelving hand.",
				CoreIdeas: []string{
					"Place the pointers: write starts where the kept part ends (often 0 or 1); read scans from left to right.",
					"Compare and decide: test nums[read] against the keep rule, often comparing with nums[write - 1].",
					"Record the answer: copy (or swap) the kept element to nums[write], then write += 1.",
					"Return write as the new length; nums[:write] is the result.",
				},
				CommonMistakes: []string{
					"Deleting elements from the list while looping over it (nums.pop(i) in a loop). Each delete is O(n), so the whole thing is O(n^2), and indices shift under you. Overwrite with a write pointer instead.",
					"Comparing nums[read] with nums[read - 1] after earlier writes have changed it. When removing duplicates, compare with the last kept value nums[write - 1].",
					"Returning the list instead of the count, or forgetting that the elements after write are leftovers.",
				},
				Examples: []CodeExample{
					{
						ID:       "remove-value",
						Title:    "Remove every copy of a value in place",
						Language: "python",
						Snippet: `def remove_value(nums, val):
    write = 0
    for read in range(len(nums)):
        if nums[read] != val:        # keep it
            nums[write] = nums[read]
            write += 1
    return write                      # nums[:write] is the result`,
						Notes: "remove_value([3, 2, 2, 3], 3) returns 2 and leaves nums starting with [2, 2].",
					},
				},
			},
			{
				Name:        "Sort, fix one, then pair-search",
				Description: "For triplets (or k numbers), sort once, fix the first number with a loop, and find the rest with opposite-end pointers. Sorting also puts duplicates side by side so they are easy to skip.",
				WhyItWorks: `Fixing nums[i] turns "three numbers summing to 0" into "two numbers to the right of i summing to -nums[i]", which is exactly the sorted pair search above: O(n) per i, O(n^2) in total, instead of O(n^3) for all triples. Sorting costs O(n log n), which is smaller than O(n^2), so it's free in Big O.

Sorting brings a second benefit: equal values become neighbours. If nums[i] equals nums[i - 1], fixing it again would find exactly the same triplets, so you skip it. After recording a match, you advance left past copies of the value just used, because with the first two values fixed the third is determined and would repeat. Comparing with the previous element (i - 1), not the next, keeps the first copy so it can still pair with the other copies.`,
				Intuition: "Seating three guests whose ages must add up to a set number. Line everyone up by age. Pick the first guest, and then you only need two people from further down the line whose ages add up to what's left, which is the two-hands shelf search again. If the next person you'd pick first is the same age as the last, they'd produce the same seatings, so you skip straight to someone different.",
				CoreIdeas: []string{
					"Sort the input: O(n log n), dominated by the O(n^2) search that follows.",
					"Fix one value with an outer loop, then place the pointers just right of it and at the end.",
					"Skip duplicates: skip a fixed value equal to the previous one; after a match, skip equal values on the left.",
					"Total cost O(n^2) time and O(1) extra memory besides the sort and the output.",
				},
				CommonMistakes: []string{
					"Producing duplicate triplets because no duplicates are skipped. Skip nums[i] == nums[i - 1] and advance left past equal values after a match.",
					"Skipping with nums[i] == nums[i + 1], which drops the first copy and loses triplets that need two equal values, like [-1, -1, 2].",
					"Sorting when you must return original indices. Sorting destroys positions; use a hash map, or sort (value, index) pairs.",
				},
				Examples: []CodeExample{
					{
						ID:       "fix-one-pair-search",
						Title:    "Count triplets with sum below a target",
						Language: "python",
						Snippet: `def count_triplets_below(nums, target):
    nums = sorted(nums)
    count = 0
    for i in range(len(nums) - 2):
        left, right = i + 1, len(nums) - 1
        while left < right:
            if nums[i] + nums[left] + nums[right] < target:
                count += right - left   # every partner from left+1..right works
                left += 1
            else:
                right -= 1
    return count`,
						Notes: "When the sum with nums[right] is below target, so is the sum with any smaller partner, so right - left triplets are counted at once. [-2, 0, 1, 3] with target 2 gives 2.",
					},
				},
			},
			{
				Name:        "Discard the side that limits the answer",
				Description: "When the answer depends on the smaller of two boundaries (like water held between two walls), move the pointer at the smaller boundary: it is the one whose best possible answer is already known.",
				WhyItWorks: `Not every two-pointer problem uses sorted order. In the water-between-walls problems the input isn't sorted, but the formula itself gives the discard argument. If a container's height is min(height[left], height[right]) and the left wall is the shorter one, then any narrower container that keeps the left wall is no taller (the short wall still caps it) and is strictly narrower, so it holds less. The short wall has nothing better to offer, and moving past it can't lose the optimum.

The same reasoning extends to running maxima: if you always process the side with the lower wall, the other side is guaranteed to contain a wall at least as tall, so the lower side's own running maximum decides the answer there. Whenever you design a two-pointer move, state this kind of proof in one sentence: "this pointer can move because every answer that still uses it is no better than one we've already seen".`,
				Intuition: "Two people carrying a long plank level over rough ground: the plank can only be as high as the shorter person can lift it. If you want to try a higher carry, swapping out the taller person can't help, because the shorter one still limits the height. Only replacing the shorter person can raise it.",
				CoreIdeas: []string{
					"Record the answer for the current pair before moving anything.",
					"Move the pointer that can't help: the one at the smaller (limiting) boundary.",
					"Justify the move: every untried pair that keeps the limiting boundary is no better than the current one.",
					"On a tie, either pointer can move, because both are limiting.",
				},
				CommonMistakes: []string{
					"Moving the taller (non-limiting) side. That can discard a boundary that belongs to the best pair.",
					"Moving a pointer without being able to say why it is safe. If you can't state the discard argument, the greedy move may be wrong; fall back to the brute force and look for the argument.",
				},
				Examples: []CodeExample{
					{
						ID:       "limiting-side",
						Title:    "Largest rectangle between two walls",
						Language: "python",
						Snippet: `def best_pair_area(height):
    left, right = 0, len(height) - 1
    best = 0
    while left < right:
        best = max(best, (right - left) * min(height[left], height[right]))
        if height[left] < height[right]:
            left += 1      # the left wall caps every narrower container using it
        else:
            right -= 1
    return best`,
						Notes: "Each step discards the wall that limits the current container. On [2, 1, 3] it measures 2*2=4, then 1*1=1, so returns 4.",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Opposite ends: pair sum closest to a target",
				Problem: "Given a sorted list nums (at least two integers) and an integer target, return the sum of two different elements that is closest to target. Example: nums = [-4, -1, 2, 6, 9], target = 4 gives 5 (-4 + 9 or -1 + 6; 5 is 1 away from 4 and no pair sums to 3 or 4).",
				Steps: []WorkedStep{
					{
						Subgoal:     "Brute force first",
						Explanation: "Try all n(n-1)/2 pairs and keep the sum with the smallest distance to target: O(n^2). The waste, as in the motivation: once a sum is too big, the right element is too big with every remaining partner, yet the brute force still tries them.",
						Code: `best = nums[0] + nums[1]
for i in range(len(nums)):
    for j in range(i + 1, len(nums)):
        if abs(nums[i] + nums[j] - target) < abs(best - target):
            best = nums[i] + nums[j]`,
						State: "n = 5 -> 10 pairs; sums -5 -2 2 5 1 5 8 8 11 15 -> closest to 4 is 5",
					},
					{
						Subgoal:     "Place the pointers",
						Explanation: "Put left on the smallest value and right on the largest. Their sum is a reasonable first guess for best.",
						Code: `left, right = 0, len(nums) - 1
best = nums[left] + nums[right]`,
						State: "left=0 (-4), right=4 (9), best = 5",
					},
					{
						Subgoal:     "Record the answer",
						Explanation: "At each step compute the current sum and keep it if it is closer to target than best. If it hits target exactly, nothing can be closer, so stop.",
						Code: `total = nums[left] + nums[right]
if abs(total - target) < abs(best - target):
    best = total
if total == target:
    return total`,
						State: "total = -4 + 9 = 5, distance 1 -> best = 5",
					},
					{
						Subgoal:     "Move the pointer that can't help",
						Explanation: "If total > target, every other partner of nums[right] in the window is at least nums[left], so their sums are at least total: further above target, never closer. Discard right. If total < target, symmetrically discard left.",
						Code: `if total < target:
    left += 1
else:
    right -= 1`,
						State: "5 > 4 -> right=3 (6): -4+6=2, dist 2 | 2 < 4 -> left=1 (-1): -1+6=5, dist 1 (tie, keep 5) | 5 > 4 -> right=2 (2): -1+2=1, dist 3 | 1 < 4 -> left=2",
					},
					{
						Subgoal:     "Stop when they meet",
						Explanation: "When left == right there is no pair left. Every pair was either measured or discarded by the argument above, so best is the closest sum. 4 comparisons instead of 10; O(n) time, O(1) memory.",
						Code:        `return best`,
						State:       "left = right = 2 -> return 5",
					},
				},
				Code: `def closest_pair_sum(nums, target):
    left, right = 0, len(nums) - 1
    best = nums[left] + nums[right]
    while left < right:
        total = nums[left] + nums[right]
        if abs(total - target) < abs(best - target):
            best = total
        if total == target:
            return total
        if total < target:
            left += 1
        else:
            right -= 1
    return best`,
			},
			{
				Title:   "Same direction: remove duplicates from a sorted list in place",
				Problem: "Given a sorted list nums, rearrange it in place so that each distinct value appears once at the front, in order, and return k, the number of distinct values. What is left after the first k positions doesn't matter. Use O(1) extra memory. Example: [1, 1, 2, 3, 3] becomes [1, 2, 3, _, _] and returns 3.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Brute force first",
						Explanation: "Building a new list of unique values uses O(n) memory, which is not allowed. Deleting each duplicate with nums.pop(i) is in place but each pop shifts the tail: O(n) per pop, O(n^2) overall. The waste: values are moved many times when each one only needs to be moved once, to its final slot.",
						State:       "pop-based on [1,1,1,1,2]: three pops, each shifting the rest left",
					},
					{
						Subgoal:     "Place the pointers",
						Explanation: "The first element is always kept, so the kept part is nums[0:1] and write = 1. The read pointer starts at 1 and scans to the end. An empty list has no distinct values, so return 0 for it.",
						Code: `if not nums:
    return 0
write = 1
for read in range(1, len(nums)):`,
						State: "nums = [1, 1, 2, 3, 3], write = 1",
					},
					{
						Subgoal:     "Compare and decide",
						Explanation: "Because the list is sorted, a value is new exactly when it differs from the last kept value, nums[write - 1]. Compare with that, not with nums[read - 1], which an earlier write may have overwritten.",
						Code:        `    if nums[read] != nums[write - 1]:`,
						State:       "read=1: 1 == 1 skip | read=2: 2 != 1 keep | read=3: 3 != 2 keep | read=4: 3 == 3 skip",
					},
					{
						Subgoal:     "Record the answer",
						Explanation: "Copy the new value into the next free slot of the kept part and advance write. write never passes read, so nothing unread is overwritten.",
						Code: `        nums[write] = nums[read]
        write += 1`,
						State: "read=2 -> nums = [1, 2, 2, 3, 3], write = 2 | read=3 -> nums = [1, 2, 3, 3, 3], write = 3",
					},
					{
						Subgoal:     "Stop when the read pointer reaches the end",
						Explanation: "Every element was read once and written at most once: O(n) time, O(1) extra memory. nums[:write] holds the distinct values in order.",
						Code:        `return write`,
						State:       "return 3; nums[:3] = [1, 2, 3]",
					},
				},
				Code: `def remove_duplicates(nums):
    if not nums:
        return 0
    write = 1
    for read in range(1, len(nums)):
        if nums[read] != nums[write - 1]:
            nums[write] = nums[read]
            write += 1
    return write`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Squares of a sorted list",
				Prompt: "nums is sorted and may contain negatives. Return a new sorted list of the squares of its values in O(n). Example: [-4, -1, 0, 3, 10] gives [0, 1, 9, 16, 100]. The largest square is always at one of the two ends, so place the pointers at both ends and fill the result from the back.",
				Code: `def sorted_squares(nums):
    result = [0] * len(nums)
    left, right = 0, len(nums) - 1
    for pos in range(len(nums) - 1, -1, -1):
        if abs(nums[left]) > abs(nums[right]):
            result[pos] = ____
            ____
        else:
            result[pos] = nums[right] * nums[right]
            ____
    return result`,
				Blanks:      []string{"nums[left] * nums[left]", "left += 1", "right -= 1"},
				Explanation: "Compare and decide: the end with the larger absolute value has the largest remaining square, so it goes into the last empty slot. Move the pointer that was used, because its value is now placed. For [-4, -1, 0, 3, 10]: 10 vs -4 -> 100; -4 vs 3 -> 16; -1 vs 3 -> 9; -1 vs 0 -> 1; then 0. Each step places one value, so O(n).",
			},
			{
				Title:  "Move the zeros to the end",
				Prompt: "Move every 0 in nums to the end, in place, keeping the other values in their original order. Return nums. Example: [0, 1, 0, 3, 12] becomes [1, 3, 12, 0, 0]. Use a read and a write pointer.",
				Code: `def move_zeroes(nums):
    write = 0
    for read in range(len(nums)):
        if ____:
            nums[write], nums[read] = nums[read], nums[write]
            ____
    return nums`,
				Blanks:      []string{"nums[read] != 0", "write += 1"},
				Explanation: "Invariant: nums[0:write] holds the nonzero values seen so far, in order, and nums[write:read] holds only zeros. A nonzero value at read is swapped into slot write (swapping a zero back to read), and write advances. When read finishes, all zeros are behind write. O(n) time, O(1) memory.",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt:      "In the sorted list [1, 2, 4, 7, 11] with target 9, left = 0 and right = 4, so the sum is 12. What should happen next?",
				Choices:     []string{"left += 1, because the sum needs to change", "right -= 1, because 11 is too big even with the smallest value, so it is in no valid pair", "Move both pointers inward", "Stop: no pair can sum to 9"},
				Answer:      1,
				Explanation: "12 > 9, and nums[left] = 1 is the smallest partner 11 could ever have, so every sum with 11 is at least 12. Discard 11. Moving left would only make the sum bigger. The next steps are 1 + 7 = 8 (move left), then 2 + 7 = 9: found.",
			},
			{
				Prompt:      "Trace remove_duplicates (read/write pointers) on [0, 0, 1, 1, 1, 2]. What does it return, and what are the first k elements?",
				Choices:     []string{"3, with [0, 1, 2]", "6, with [0, 0, 1, 1, 1, 2]", "2, with [0, 1]", "3, with [0, 1, 1]"},
				Answer:      0,
				Explanation: "write starts at 1. read=1: 0 equals the last kept 0, skip. read=2: 1 is new -> nums[1] = 1, write = 2. read=3, 4: 1 again, skip. read=5: 2 is new -> nums[2] = 2, write = 3. Return 3; nums[:3] = [0, 1, 2].",
			},
			{
				Prompt:      "Which approach fits: \"Find two numbers in an unsorted list that add up to target and return their original indices\"?",
				Choices:     []string{"Opposite-end pointers on the list as given", "A hash map of value -> index with the seen-so-far pattern (Chapter 2)", "Sort the list, then opposite-end pointers, returning the positions in the sorted list", "A read/write pointer pass"},
				Answer:      1,
				Explanation: "The input is unsorted, so the pointer discard argument doesn't hold on it, and sorting would scramble the indices you must return (you'd have to sort (value, index) pairs). The hash map is O(n) time and keeps the original indices. Two pointers win when the input is already sorted or extra memory is forbidden.",
			},
			{
				Prompt:      "Chapter 2 detected duplicates with a set, using O(n) extra memory. If you may modify the input, how can you detect duplicates with O(1) extra memory?",
				Choices:     []string{"Sort the list in place, then compare each element with its neighbour in one same-direction scan", "Use a dict instead of a set", "Use a list instead of a set to store seen values", "Run opposite-end pointers on the unsorted list"},
				Answer:      0,
				Explanation: "After sorting, equal values are adjacent, so one scan comparing nums[i] with nums[i - 1] finds them: O(n log n) time with an in-place sort (nums.sort()) and O(1) extra memory. A dict also uses O(n) memory, and a list of seen values makes each lookup O(n). This is the classic trade: a hash set buys O(n) time with O(n) memory.",
			},
			{
				Prompt:      "State the invariant that makes the opposite-end pair search correct, and explain why each pointer move keeps it true.",
				Explanation: "Invariant: if a pair with the target sum exists, both of its indices lie inside [left, right]. If nums[left] + nums[right] < target, then nums[left] plus any partner in the window is at most that sum (nums[right] is the largest remaining), so nums[left] is in no valid pair and left += 1 keeps the invariant. If the sum is > target, nums[right] plus any partner is at least that sum (nums[left] is the smallest remaining), so right -= 1 is safe. Each step shrinks the window, so the loop ends in at most n - 1 steps, having found the pair if it exists.",
			},
			{
				Prompt:      "When should you choose a hash map and when two pointers for a pair-sum problem?",
				Explanation: "Hash map (Chapter 2): the input is unsorted, you need original indices, and O(n) extra memory is allowed; it runs in O(n). Two pointers: the input is sorted (or sorting is cheap and indices don't matter) or extra memory must be O(1); it runs in O(n) on sorted input, O(n log n) if you sort first. Also prefer two pointers when you need the closest sum rather than an exact one, because a hash lookup can only find exact values.",
			},
		},
	}
}
