package chapter

import "time"

// chapter06 teaches binary search: the exact-value template, the boundary
// ("first true") template, binary search on the answer, and searching data
// that is only partly sorted. See docs/pedagogy.md.
// Storyboards (IDs registered in app.js): binary-search, rotated-array, median-two-arrays.
func chapter06() Chapter {
	return Chapter{
		Number:        6,
		Slug:          "binary-search",
		Title:         "Binary Search",
		Category:      "binary-search",
		Prerequisites: []int{1},
		Summary:       "Halve the search space with every comparison. Learn two loop templates (exact value with lo <= hi, first true with lo < hi), when to use each, how to avoid infinite loops, and how to binary-search an answer range with a monotonic yes/no test.",
		Motivation: `A sorted array holds 1,000,000 numbers and you need to know whether 734,519 is among them. A loop that checks every element does up to 1,000,000 comparisons, and if you have 1,000,000 such questions that's 10^12 comparisons: hours of work.

Now use the order. Look at the middle element, position 500,000. Say it holds 512,000. Since the array is sorted, everything to its left is smaller still, so half a million candidates are gone after one comparison. Look at the middle of what's left, and again, and again: 1,000,000 -> 500,000 -> 250,000 -> ... -> 1 takes about 20 steps, because 2^20 is just over a million. That's binary search: O(log n) instead of O(n).

The same halving works even when there's no array to search. Koko must eat 1,000,000,000 bananas in 2 hours: what's the slowest speed that works? Trying speeds 1, 2, 3, ... could take 500 million tries. But speeds have an order too: if speed k is fast enough, every faster speed is as well. So you can binary-search the speeds themselves, in about 30 checks.`,
		Objectives: []string{
			"Implement the exact-value template (while lo <= hi, lo = mid + 1, hi = mid - 1) and explain why each bound excludes mid.",
			"Implement the boundary template (while lo < hi, hi = mid, lo = mid + 1) to find the first index where a yes/no test becomes true, and explain why it cannot loop forever.",
			"Choose between the two templates from the wording of a problem: 'find this value' versus 'find the first/smallest/last position that ...'.",
			"Turn a 'minimum value that satisfies a condition' problem into binary search on the answer by writing a monotonic yes/no predicate.",
			"Trace lo, hi and mid on a small input, including rotated arrays where only one half is sorted.",
		},
		WhenToUse: []string{
			"The input is sorted (or sorted by rows, or sorted and then rotated) and the problem asks you to find something in it.",
			"The problem demands O(log n) time, or the input is so large that O(n) per query is too slow.",
			"You're asked for the minimum (or maximum) value such that some condition holds, and if a value works then every larger (or smaller) value also works: 'smallest speed', 'least capacity', 'minimum days'.",
			"You need the first or last position where something becomes true: first element >= x, last timestamp <= t, first bad version.",
		},
		Concepts: []Concept{
			{
				Name:        "Template 1: find an exact value (closed range, lo <= hi)",
				Description: "Keep a range [lo, hi] of indices where the target could still be, both ends included. Compare the middle with the target; return it on a match, otherwise move lo to mid + 1 or hi to mid - 1. Stop when the range is empty (lo > hi).",
				WhyItWorks:  "Invariant: if the target is in the array, its index lies in [lo, hi]. It's true at the start, when the range is the whole array. If nums[mid] < target, then by sortedness every index up to mid also holds a value < target, so lo = mid + 1 keeps the invariant; the other case is symmetric. mid has been compared already, so excluding it is safe, and it also guarantees the range shrinks by at least one every iteration. When the range is empty, the invariant says the target can't be anywhere, so returning -1 is correct. Each step at least halves the range, so there are at most about log2(n) + 1 steps.",
				Intuition:   "It's the number-guessing game: 'I'm thinking of a number from 1 to 100.' You guess 50 and hear 'higher'. You don't guess 51 next; you guess 75, the middle of what's left. Your guess is mid, 'higher' is nums[mid] < target, and the remaining range 51..100 is the new [lo, hi]. Since 50 was already wrong, you never include it again.",
				CoreIdeas: []string{
					"Use it when you're looking for one specific value that might not be there, and you can stop the moment you see it.",
					"lo and hi are both live candidates, so the loop runs while lo <= hi; a one-element range (lo == hi) still needs checking.",
					"Both updates skip mid (mid + 1, mid - 1) because mid was already checked; that's what makes the loop finish.",
					"mid = (lo + hi) // 2 always lies in [lo, hi]. Python integers don't overflow; in Java or C++ write lo + (hi - lo) / 2.",
				},
				CommonMistakes: []string{
					"Looping while lo < hi with a closed range: the last remaining candidate is never checked, so [5] with target 5 returns -1. With hi = len(nums) - 1, loop while lo <= hi.",
					"Writing lo = mid or hi = mid in this template: with lo = 3, hi = 4, mid = 3, and lo = mid changes nothing, so the loop never ends. Skip mid, since it was compared already.",
					"Starting hi at len(nums) and then reading nums[mid] when mid can equal len(nums): IndexError for a target larger than everything. The closed range starts at hi = len(nums) - 1.",
				},
				Examples: []CodeExample{
					{
						ID:       "exact-template",
						Title:    "Exact-value binary search",
						Language: "python",
						Snippet: `def search(nums, target):
    lo, hi = 0, len(nums) - 1          # Define the search space
    while lo <= hi:
        mid = (lo + hi) // 2           # Pick the middle
        if nums[mid] == target:        # Write the yes/no test
            return mid
        if nums[mid] < target:         # Discard half
            lo = mid + 1
        else:
            hi = mid - 1
    return -1                          # Read the answer: not found`,
						Notes: "nums = [1,3,5,7,9,11], target = 7: mid = 2 (5, too small), mid = 4 (9, too big), mid = 3 (7, found).",
					},
				},
			},
			{
				Name:        "Template 2: find the first true (half-open range, lo < hi, hi = mid)",
				Description: "Many questions aren't 'where is x?' but 'where does something start?': the first element >= x, the first bad version, the first speed that's fast enough. Write the question as a yes/no test that is false for a prefix and true for the rest (no, no, no, yes, yes). Then search for the first yes: if mid says yes, the answer is mid or earlier (hi = mid); if no, it's later (lo = mid + 1). Stop when lo == hi.",
				WhyItWorks:  "Invariant: the first index where the test is true lies in [lo, hi]. Set hi to a position that's guaranteed true, or one past the end to stand for 'nothing is true'. If test(mid) is true, the first true can't be after mid, but mid itself might be it, so hi = mid keeps it (hi = mid - 1 could lose the answer). If test(mid) is false, the test is false for everything up to mid because it's monotonic, so lo = mid + 1. Why no infinite loop: mid = (lo + hi) // 2 rounds down, so while lo < hi, mid < hi. hi = mid therefore always shrinks the range, and lo = mid + 1 does too. The loop ends with lo == hi, the single remaining candidate. The danger is the mirror image: lo = mid. With lo = 3, hi = 4, mid = 3, and lo = mid changes nothing, so that version loops forever unless mid rounds up, (lo + hi + 1) // 2.",
				Intuition:   "Finding the first bad photo in a camera roll of 1,000 pictures, where some setting got switched on at one point and every photo after it is bad. You open photo 500: bad. The first bad one is 500 or earlier, and 500 itself might be it, so you keep it: 'hi = 500'. Open 250: good, so the change happened after 250: 'lo = 251'. The test 'is this photo bad?' is the predicate, and bad photos form one block at the end, which is what makes halving valid.",
				CoreIdeas: []string{
					"Use it for 'first/smallest/earliest position or value such that ...' and, by stepping back one, for 'last position such that ...' (the last <= t is one before the first > t).",
					"The predicate must be monotonic: once true, true for all later positions. Check that before you code.",
					"Loop while lo < hi, set hi = mid on true and lo = mid + 1 on false, and return lo. Every boundary problem in this chapter uses this exact skeleton.",
					"If 'nothing is true' is a possible outcome, let hi start one past the last candidate (for example len(nums)) and treat that value as 'not found'.",
					"Rule of thumb for the two templates: exact value that may be missing, use template 1; a boundary, first or last, use template 2. Don't mix their loop conditions and updates.",
				},
				CommonMistakes: []string{
					"Pairing hi = mid with while lo <= hi: once lo == hi and the test is true, hi = mid changes nothing and the loop spins forever. Template 2 always uses while lo < hi.",
					"Writing lo = mid for the 'false' branch: because mid rounds down, lo = mid can leave the range unchanged (lo = 3, hi = 4). Use lo = mid + 1, or round mid up if you really need lo = mid.",
					"Setting hi to the last index when the answer may be 'none of them': the search can then never report 'none'. Use one past the end, as in 'last timestamp <= t' when every timestamp qualifies.",
				},
				Examples: []CodeExample{
					{
						ID:       "first-true-template",
						Title:    "First index with nums[i] >= target (lower bound)",
						Language: "python",
						Snippet: `def lower_bound(nums, target):
    lo, hi = 0, len(nums)              # Define the search space (len = "none")
    while lo < hi:
        mid = (lo + hi) // 2           # Pick the middle (rounds down, so mid < hi)
        if nums[mid] >= target:        # Write the yes/no test
            hi = mid                   # Discard half: mid may be the answer, keep it
        else:
            lo = mid + 1               # mid is too small, drop it
    return lo                          # Read the answer: first true`,
						Notes: "nums = [1,3,3,3,5,8], target = 3 gives 1; target = 9 gives 6, meaning no element is >= 9.",
					},
				},
			},
			{
				Name:        "Binary search on the answer",
				Description: "Sometimes the thing you binary-search isn't in an array at all: it's the answer itself. If a problem asks for the minimum value v (speed, capacity, days, size) such that feasible(v) is true, and feasibility is monotonic (if v works, v + 1 works), then the values from the lowest possible to the highest possible form a sorted no...no yes...yes line. Binary-search that line for the first yes with template 2, calling a feasibility check at each mid.",
				WhyItWorks:  "Monotonicity is everything. Take Koko: hours(k) = sum of ceil(pile / k). Raising k can only lower or keep each term, so hours(k) never increases, and 'hours(k) <= h' is false for small k and true from some k on. That's the shape template 2 needs. Pick lo as a value that might work and hi as one that surely works (k = max(piles) eats one pile per hour), and the first-true search returns the smallest feasible k. Cost: the check is O(n) and the range has size R, so the total is O(n log R): for R = 10^9 that's about 30 checks instead of up to 10^9.",
				Intuition:   "Setting an oven timer for a cake you've never baked, with a tester who only says 'done' or 'not done'. If 40 minutes is done, 41 is too. Instead of trying 1, 2, 3, ... minutes, try 60 (done), then 30 (not done), then 45 (done), and so on. Minutes play the role of indices, and 'done?' is the predicate. You're searching the space of possible answers, not a list of data.",
				CoreIdeas: []string{
					"Recognise it: 'minimum/maximum X such that ...', a feasibility check that's easy for a fixed X, and a monotonic yes/no answer.",
					"Write the feasibility check first, as its own function. It's usually a greedy O(n) pass (count hours, count days, count pieces).",
					"Choose bounds that bracket the answer: lo = smallest value that could possibly work, hi = a value that certainly works. For Koko, [1, max(piles)]; for shipping, [max(weights), sum(weights)].",
					"Then it's template 2, unchanged: hi = mid when feasible, lo = mid + 1 when not.",
				},
				CommonMistakes: []string{
					"Choosing a lower bound that can never work or causes errors, like speed 0 (division by zero) or a ship capacity below the heaviest package (that package can never be loaded). Start lo at the smallest value that's legal.",
					"Choosing an upper bound that might not be feasible, like max(piles) - 1: when the answer is exactly max(piles), the search returns a wrong value. hi must be a value you know works.",
					"Not checking monotonicity: if 'works' can flip back to 'doesn't work' for a larger value, binary search is not valid and may skip the answer.",
				},
				Examples: []CodeExample{
					{
						ID:       "search-on-answer",
						Title:    "Smallest feasible value with a predicate",
						Language: "python",
						Snippet: `def smallest_feasible(lo, hi, feasible):
    # Requires: feasible is monotonic and feasible(hi) is True.
    while lo < hi:
        mid = (lo + hi) // 2
        if feasible(mid):
            hi = mid
        else:
            lo = mid + 1
    return lo

piles, h = [3, 6, 7, 11], 8
hours = lambda k: sum((p + k - 1) // k for p in piles)
print(smallest_feasible(1, max(piles), lambda k: hours(k) <= h))  # 4`,
						Notes: "The same function solves Koko, shipping capacity and many more: only feasible and the bounds change.",
					},
				},
			},
			{
				Name:        "Partly sorted data: find the sorted half",
				Description: "A sorted array that was rotated, like [4,5,6,7,0,1,2], is not sorted, but it's two sorted runs. Split any range at mid and at least one half has no drop in it, so it is sorted. For that half, two comparisons with its ends tell you whether the target can be inside. Keep that half if so, otherwise keep the other one.",
				WhyItWorks:  "A rotated sorted array has at most one drop (a place where a value is followed by a smaller one). The two halves [lo, mid] and [mid, hi] can't both contain it, so one is in ascending order, and nums[lo] <= nums[mid] detects when it's the left one. In a sorted half, the target is inside exactly when it lies between that half's end values, so that test decides which half can't hold the target. The invariant 'if present, the target is in [lo, hi]' holds as in template 1, and the range halves each time. The same idea, 'compare with an end to know which run you're in', gives the minimum too: elements before the minimum are all > nums[-1], the rest are <= nums[-1], a perfect first-true predicate.",
				Intuition:   "A clock face starting at 4: 4,5,6,7,0,1,2. Cut it anywhere and at least one piece reads in order without wrapping past 7 back to 0. You can only trust the in-order piece to answer 'is 0 between your ends?', so you ask it, and that tells you which side to throw away.",
				CoreIdeas: []string{
					"First find the sorted half (nums[lo] <= nums[mid] means the left half is sorted), then decide: is the target within its ends?",
					"For the minimum of a rotated array, compare with the last element: the predicate nums[i] <= nums[-1] is monotonic, so use template 2.",
					"The median of two sorted arrays is the same spirit taken further: binary-search a cut position i in the shorter array, check the four border values, and move the cut toward the side that's too big.",
				},
				CommonMistakes: []string{
					"Using a strict nums[lo] < nums[mid]: when lo == mid the left half is a single element, which is sorted, but the strict test says otherwise and sends the search the wrong way ([3,1], target 1).",
					"Comparing only target with nums[mid], as in plain binary search: in [4,5,6,7,0,1,2], 0 < 7 but 0 is to the right. Identify the sorted half first.",
				},
				Examples: []CodeExample{
					{
						ID:       "rotated-search",
						Title:    "Search a rotated sorted array",
						Language: "python",
						Snippet: `def search_rotated(nums, target):
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] == target:
            return mid
        if nums[lo] <= nums[mid]:                 # Find the sorted half: left
            if nums[lo] <= target < nums[mid]:    # target inside it?
                hi = mid - 1
            else:
                lo = mid + 1
        else:                                     # right half is sorted
            if nums[mid] < target <= nums[hi]:
                lo = mid + 1
            else:
                hi = mid - 1
    return -1`,
						Notes: "[4,5,6,7,0,1,2], target 0: mid 3 (left [4..7] sorted, 0 not in it, lo = 4), mid 5 (left [0,1] sorted, 0 in it, hi = 4), mid 4: found.",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "First position of a value in a sorted array with duplicates",
				Problem: "Given a sorted array nums that may contain duplicates and a value target, return the first index i with nums[i] >= target, or len(nums) if every element is smaller. For nums = [1,3,3,3,5,8] and target = 3 the answer is 1. (This is Python's bisect_left, and the heart of LeetCode 34 and 35.)",
				Steps: []WorkedStep{
					{
						Subgoal:     "Write the yes/no test",
						Explanation: "We want where the 3s start, not just any 3, so stopping at the first match (template 1) could return index 2 or 3. Ask a question whose answers form one switch: 'is nums[i] >= target?'. Since nums is sorted, it's false for the small values and true from the first 3 on.",
						Code:        "test(i) = nums[i] >= target",
						State:       "nums = [1,3,3,3,5,8], target 3: test gives [no, yes, yes, yes, yes, yes]",
					},
					{
						Subgoal:     "Define the search space",
						Explanation: "The answer is an index from 0 to len(nums), where len(nums) means 'no element is >= target'. Including that extra slot lets the search report 'none' without a special case.",
						Code:        "lo, hi = 0, len(nums)",
						State:       "lo = 0, hi = 6",
					},
					{
						Subgoal:     "Pick the middle",
						Explanation: "Loop while more than one candidate remains (lo < hi). mid rounds down, so mid is always strictly less than hi, which is what lets hi = mid make progress.",
						Code:        "while lo < hi:\n    mid = (lo + hi) // 2",
						State:       "lo = 0, hi = 6: mid = 3, nums[3] = 3",
					},
					{
						Subgoal:     "Discard half",
						Explanation: "If the test is true at mid, the first true is mid or to its left; mid could be the answer, so keep it with hi = mid. If false, mid and everything left of it are too small: lo = mid + 1.",
						Code:        "    if nums[mid] >= target:\n        hi = mid\n    else:\n        lo = mid + 1",
						State:       "mid 3: 3 >= 3 yes, hi = 3. mid 1: 3 >= 3 yes, hi = 1. mid 0: 1 >= 3 no, lo = 1. Now lo == hi == 1.",
					},
					{
						Subgoal:     "Read the answer",
						Explanation: "When lo == hi a single candidate is left, and by the invariant it's the first true. For target = 9 the test is never true, every step does lo = mid + 1, and lo ends at 6 = len(nums), meaning 'none'.",
						Code:        "return lo",
						State:       "returns 1 (the first 3). target 9 returns 6.",
					},
				},
				Code: `def lower_bound(nums, target):
    lo, hi = 0, len(nums)
    while lo < hi:
        mid = (lo + hi) // 2
        if nums[mid] >= target:
            hi = mid
        else:
            lo = mid + 1
    return lo`,
			},
			{
				Title:   "Binary search on the answer: shipping capacity",
				Problem: "Packages with weights [3,2,2,4,1,4] must be shipped in order within days = 3. Each day the ship is loaded with packages in order until the next one would exceed its capacity. What is the smallest capacity that ships everything in 3 days? (LeetCode 1011.) The answer is 6: day 1 takes 3+2, day 2 takes 2+4, day 3 takes 1+4.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Write the yes/no test",
						Explanation: "For a fixed capacity, counting the days needed is easy: fill each day greedily. A bigger ship never needs more days, so 'days_needed(c) <= days' is monotonic: false, false, ..., true, true. That makes it searchable.",
						Code: `def days_needed(capacity):
    used, load = 1, 0
    for w in weights:
        if load + w > capacity:
            used += 1
            load = 0
        load += w
    return used`,
						State: "days_needed(5) = 4 ([3,2] [2] [4,1] [4]), days_needed(6) = 3 ([3,2] [2,4] [1,4])",
					},
					{
						Subgoal:     "Define the search space",
						Explanation: "The capacity can't be below the heaviest package (4), or that package never fits. At capacity sum(weights) = 16 everything goes in one day, so that surely works. The answer is in [4, 16].",
						Code:        "lo, hi = max(weights), sum(weights)",
						State:       "lo = 4, hi = 16",
					},
					{
						Subgoal:     "Pick the middle",
						Explanation: "Same template 2 loop as the lower bound: while lo < hi, mid = (lo + hi) // 2. The only difference is that the test calls days_needed instead of reading an array.",
						Code:        "while lo < hi:\n    mid = (lo + hi) // 2",
						State:       "lo = 4, hi = 16: mid = 10, days_needed(10) = 2",
					},
					{
						Subgoal:     "Discard half",
						Explanation: "Feasible means this capacity or a smaller one is the answer (hi = mid); infeasible means we need more (lo = mid + 1).",
						Code:        "    if days_needed(mid) <= days:\n        hi = mid\n    else:\n        lo = mid + 1",
						State:       "mid 10: 2 days, yes, hi = 10. mid 7: 3 days, yes, hi = 7. mid 5: 4 days, no, lo = 6. mid 6: 3 days, yes, hi = 6.",
					},
					{
						Subgoal:     "Read the answer",
						Explanation: "lo == hi == 6 is the smallest capacity whose test is true. Four checks instead of trying all 13 capacities from 4 to 16; for weights summing to 10^7 it would be about 24 checks.",
						Code:        "return lo",
						State:       "returns 6",
					},
				},
				Code: `def shipWithinDays(weights, days):
    def days_needed(capacity):
        used, load = 1, 0
        for w in weights:
            if load + w > capacity:
                used += 1
                load = 0
            load += w
        return used

    lo, hi = max(weights), sum(weights)
    while lo < hi:
        mid = (lo + hi) // 2
        if days_needed(mid) <= days:
            hi = mid
        else:
            lo = mid + 1
    return lo`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Integer square root (LeetCode 69)",
				Prompt: "Return the integer part of the square root of a non-negative integer x, e.g. 8 -> 2 and 16 -> 4, without math.sqrt. Use the same subgoals: the largest k with k * k <= x is one before the first k with k * k > x. Fill in the blanks of the template 2 search over k in [0, x + 1].",
				Code: `def mySqrt(x):
    lo, hi = 0, x + 1          # Define the search space
    while ____:
        mid = (lo + hi) // 2   # Pick the middle
        if ____:               # Write the yes/no test: too big?
            ____
        else:
            ____
    return ____                # Read the answer`,
				Blanks:      []string{"lo < hi", "mid * mid > x", "hi = mid", "lo = mid + 1", "lo - 1"},
				Explanation: "'mid * mid > x' is false for small k and true from some k on, so it's a first-true search: while lo < hi, hi = mid when true, lo = mid + 1 when false. hi starts at x + 1 because (x + 1)^2 > x always holds, so a true is guaranteed in range. The loop finds the first k that's too big; the answer is the one before it, lo - 1. For x = 8: (lo, hi, mid) goes (0,9,4) too big, (0,4,2) fine so lo = 3, (3,4,3) too big so hi = 3, giving lo - 1 = 2.",
			},
		},
		Animations: []Storyboard{
			{
				ID:    "binary-search",
				Title: "Binary Search Walkthrough",
				Goal:  "Find 15 in the sorted array [1,3,5,...,25] (13 elements) with the exact-value template, in three looks instead of eight.",
				MemoryTips: []string{
					"Define the search space: lo = 0, hi = len - 1, both ends included.",
					"Pick the middle: mid = (lo + hi) // 2 while lo <= hi.",
					"Write the yes/no test: equal? Return mid.",
					"Discard half: too small means lo = mid + 1, too big means hi = mid - 1 (mid is already checked, so skip it).",
					"Read the answer: if lo passes hi, the value isn't there.",
				},
				Steps: []StoryboardStep{
					{Cue: "sorted-lineup", Narration: "13 sorted numbers, 1 to 25, and we want 15. The whole array is still possible: lo = 0, hi = 12. Sorted order is what will let one comparison rule out many elements.", VisualHint: "Array laid out left to right", Duration: Duration(2 * time.Second), CodeRef: "binary_search: init low/high"},
					{Cue: "pick-middle", Narration: "mid = (0 + 12) // 2 = 6, which holds 13. 13 is smaller than 15, and so is everything to its left, because the array is sorted.", VisualHint: "Highlight mid index", Duration: Duration(3 * time.Second), CodeRef: "binary_search: mid calc"},
					{Cue: "discard-half", Narration: "Discard indices 0 to 6 in one move: lo = mid + 1 = 7. mid itself is excluded since we've already checked it. Six candidates are left, 15 to 25.", VisualHint: "Fade out irrelevant section", Duration: Duration(3 * time.Second), CodeRef: "binary_search: adjust bounds"},
					{Cue: "repeat", Narration: "Repeat: mid = (7 + 12) // 2 = 9 holds 19, too big, so hi = 8. Then mid = (7 + 8) // 2 = 7 holds 15: found at index 7 after 3 looks. If lo had passed hi, 15 would not be in the array.", VisualHint: "Loop indicator", Duration: Duration(3 * time.Second), CodeRef: "binary_search: while loop"},
				},
			},
			{
				ID:    "rotated-array",
				Title: "Search in Rotated Sorted Array",
				Goal:  "Find 0 in [4,5,6,7,0,1,2] by first finding the sorted half, then asking whether the target lies between its ends.",
				MemoryTips: []string{
					"Define the search space: lo = 0, hi = len - 1, loop while lo <= hi.",
					"Pick the middle and check it: equal? Return mid.",
					"Find the sorted half: nums[lo] <= nums[mid] means the left half is sorted, otherwise the right half is.",
					"Discard half: if the target is between the sorted half's ends, keep that half; otherwise keep the other.",
					"Use inclusive ends (<=) for the sorted half; only mid is excluded.",
				},
				Steps: []StoryboardStep{
					{Cue: "init", Narration: "The array was sorted, then rotated, so it is two ascending runs: [4,5,6,7] and [0,1,2]. Start with lo = 0, hi = 6, target 0.", VisualHint: "Array with rotation point", Duration: Duration(2 * time.Second), CodeRef: "rotated: low, high = 0, len-1"},
					{Cue: "find-sorted", Narration: "mid = 3 holds 7, not the target. nums[lo] = 4 <= 7, so the left half [4,5,6,7] has no drop: it is sorted. Only a sorted half can answer 'is the target between your ends?' reliably.", VisualHint: "Highlight sorted half", Duration: Duration(3 * time.Second), CodeRef: "rotated: nums[low] <= nums[mid]"},
					{Cue: "narrow", Narration: "Is 4 <= 0 < 7? No, so 0 can't be in the left half: lo = 4. Next mid = 5 holds 1; nums[4] = 0 <= 1, so [0,1] is the sorted half, and 0 <= 0 < 1 holds, so hi = 4.", VisualHint: "Discard half", Duration: Duration(3 * time.Second), CodeRef: "rotated: adjust bounds"},
					{Cue: "found", Narration: "mid = 4 holds 0: found at index 4 after three looks. Had the range emptied (lo > hi), the answer would be -1.", VisualHint: "Result", Duration: Duration(2 * time.Second), CodeRef: "rotated: return mid or -1"},
				},
			},
			{
				ID:    "median-two-arrays",
				Title: "Median of Two Sorted Arrays",
				Goal:  "Find the median of A = [1,3,8] and B = [7,9,10,11] by binary-searching how many elements A puts in the left half.",
				MemoryTips: []string{
					"Search the shorter array; the left half holds half = (m + n + 1) // 2 elements.",
					"Place the partition: i from A, j = half - i from B.",
					"Check the partition: A's left border <= B's right border and B's left border <= A's right border (missing borders are -inf / +inf).",
					"Move the partition: A's left too big means fewer from A (hi = i - 1); B's left too big means more from A (lo = i + 1).",
					"Compute the median: odd total takes max of the left borders; even total averages it with min of the right borders.",
				},
				Steps: []StoryboardStep{
					{Cue: "partition", Narration: "7 numbers in total, so the left half holds half = 4 of them. A is shorter, so we search i, the count A gives to the left, over 0..3. First try i = 1, which forces j = 4 - 1 = 3 from B.", VisualHint: "Partition bars", Duration: Duration(3 * time.Second), CodeRef: "partition"},
					{Cue: "check", Narration: "Borders: A's left is 1, A's right is 3, B's left is 10, B's right is 11. 1 <= 11 is fine, but 10 <= 3 fails: B's left side holds a number bigger than something on A's right side.", VisualHint: "Compare boundaries", Duration: Duration(3 * time.Second), CodeRef: "bounds"},
					{Cue: "adjust", Narration: "So A must give more: lo = 2. i = 2, j = 2: 9 <= 8 fails again, lo = 3. i = 3, j = 1: 8 <= 9 and 7 <= +inf both hold. The total is odd, so the median is max(8, 7) = 8. Check: merged it's [1,3,7,8,9,10,11].", VisualHint: "Shift partition", Duration: Duration(3 * time.Second), CodeRef: "adjust"},
				},
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Which pattern fits best? 'Given an array of package weights and a number of days, find the smallest ship capacity that ships all packages, in order, within that many days.'",
				Choices: []string{
					"Sliding window (chapter 4): the packages shipped on one day form a contiguous block",
					"Monotonic stack (chapter 5): track the heaviest package seen so far",
					"Binary search on the answer: 'can capacity c do it in time?' is monotonic in c",
					"Two pointers (chapter 3): one pointer at each end of the array",
				},
				Answer:      2,
				Explanation: "The cue is 'smallest value such that a condition holds', with a condition that only gets easier as the value grows. Each day's load is contiguous, but no single window answers the question, and nothing is being matched across ends or popped as a 'next greater' element. Write days_needed(c), search c over [max(weights), sum(weights)] with template 2.",
			},
			{
				Prompt: "Chapter 1 review: a sorted array has 1,000,000 elements. About how many comparisons does binary search need in the worst case?",
				Choices: []string{
					"About 1,000,000, because it may still have to visit every element",
					"About 500,000, because it discards half the array",
					"About 20, because each step halves the range and 2^20 is roughly 1,000,000",
					"About 1,000, the square root of n",
				},
				Answer:      2,
				Explanation: "Each comparison halves the remaining range: 1,000,000 -> 500,000 -> ... -> 1 takes log2(1,000,000), about 20, steps. That's O(log n). Halving once (500,000) is only the first step, and visiting everything is the O(n) linear scan.",
			},
			{
				Prompt:      "Trace the exact-value template on nums = [1,3,5,7,9,11] with target = 7 (lo = 0, hi = 5, mid = (lo + hi) // 2). Which indices does mid take, in order?",
				Explanation: "mid = 2 (value 5 < 7, so lo = 3); mid = (3 + 5) // 2 = 4 (value 9 > 7, so hi = 3); mid = (3 + 3) // 2 = 3 (value 7, found). The answer is 2, 4, 3, and the function returns 3.",
			},
			{
				Prompt: "This loop is meant to find the LAST index where ok(i) is true (ok is true then false):\n\nwhile lo < hi:\n    mid = (lo + hi) // 2\n    if ok(mid):\n        lo = mid\n    else:\n        hi = mid - 1\n\nWhat happens when lo = 3, hi = 4 and ok(3) is true?",
				Choices: []string{
					"It returns 3 correctly",
					"It loops forever: mid = 3, lo = mid leaves the range [3, 4] unchanged",
					"It skips index 4 and returns a wrong answer",
					"It raises an IndexError",
				},
				Answer:      1,
				Explanation: "mid rounds down, so with lo = 3, hi = 4, mid = 3 and lo = mid changes nothing; the loop repeats forever. hi = mid is safe in template 2 because mid < hi always, but lo = mid is not. Fixes: round mid up with (lo + hi + 1) // 2, or rewrite it as a first-true search for 'not ok' and subtract 1.",
			},
			{
				Prompt:      "In your own words: when do you use while lo <= hi with mid + 1 / mid - 1, and when do you use while lo < hi with hi = mid? What's the danger of mixing them?",
				Explanation: "Use lo <= hi when you're looking for one exact value that may be missing: you check mid, return on a match, and exclude mid otherwise; the loop ends when the range is empty. Use lo < hi with hi = mid when you're looking for a boundary, the first index where a monotonic yes/no test turns true (or the smallest feasible answer): mid might be the answer, so it stays in range, and the loop ends when one candidate is left. Mixing them breaks things: lo <= hi with hi = mid loops forever once lo == hi, and lo < hi with mid - 1 can skip the last candidate.",
			},
			{
				Prompt: "Why is 'can Koko finish in h hours at speed k?' a valid predicate for binary search?",
				Choices: []string{
					"Because the piles array is sorted",
					"Because it is monotonic: if speed k works, every faster speed also works",
					"Because the answer is always max(piles)",
					"Because ceil(p / k) is always an integer",
				},
				Answer:      1,
				Explanation: "Binary search needs a test that switches from no to yes only once. Eating faster never takes more hours, so once a speed is fast enough all faster speeds are too. The piles are not sorted at all; the search runs over the speeds, not the array.",
			},
			{
				Prompt:      "Which pattern: 'Given a sorted array and a target, find two numbers that add up to target.' Binary search for target - nums[i] for each i works. Is it the best choice?",
				Explanation: "It works in O(n log n), but two pointers from chapter 3 do it in O(n): start at both ends and move the left pointer up when the sum is too small, the right one down when it's too big. Binary search shines when you need a single lookup or boundary in sorted data, or to search an answer range; for pairs in a sorted array, two pointers use the order better.",
			},
		},
	}
}
