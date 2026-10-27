package chapter

// chapter18 teaches interval problems: the overlap test, choosing whether to
// sort by start or by end, and sweeping a timeline to count overlaps.
//
// Subgoal labels used here and in the intervals problems: "Sort the
// intervals", "Test for overlap", "Merge or keep", "Track the running end",
// "Return the answer".
// Storyboards: none.
func chapter18() Chapter {
	return Chapter{
		Number:        18,
		Slug:          "intervals",
		Title:         "Intervals",
		Category:      "intervals",
		Prerequisites: []int{17},
		Summary:       "Interval problems give you ranges [start, end] and ask whether they clash, how to merge them, how many to keep, or how many overlap at once. Almost all of them are solved by one sort and one pass: learn the overlap test, when to sort by start and when by end, and how to sweep a timeline.",
		Motivation: `A calendar app holds 10,000 busy slots for a team, as [start, end] pairs in no particular order. It needs to show the free time, which means merging overlapping busy slots into blocks.

The direct way compares every pair of slots to see whether they overlap, merges any pair that does, and repeats until nothing changes. One round alone is 10,000 * 9,999 / 2, about 50 million comparisons, and merges can trigger more rounds.

Sort the slots by start time first, and the problem collapses: each slot can only overlap the block you are currently building, because every earlier block ended before this block began. One sort (about 10,000 * 14 = 140,000 comparisons) and one pass do the whole job.

The same move, sort then sweep left to right, answers 'can one person attend all meetings?', 'how many rooms do we need?' and 'how few meetings must we cancel?'. What changes is the sort key and what you keep track of during the sweep.`,
		Objectives: []string{
			"Write the overlap test for two intervals, and adjust it for closed intervals (touching counts) versus half-open ones (touching doesn't)",
			"Explain why sorting by start puts overlaps next to each other, and why selection problems sort by end instead",
			"Trace a merge or a sweep on a small input, keeping the running end",
			"Implement Merge Intervals, Insert Interval and Meeting Rooms II, and recognise Non-overlapping Intervals as activity selection from chapter 17",
		},
		WhenToUse: []string{
			"The input is a list of [start, end] pairs: meetings, bookings, ranges, segments on a line",
			"The question asks whether ranges clash, to merge or insert ranges, or to cover points with as few ranges as possible",
			"'Minimum number of rooms/resources' or 'maximum number happening at once' means counting overlaps with a sweep",
			"'Remove the fewest so the rest don't overlap' or 'keep the most' is activity selection: sort by end",
		},
		Concepts: []Concept{
			{
				Name:        "The overlap test",
				Description: "Two intervals a and b overlap when each starts before the other ends: a.start < b.end and b.start < a.end. Whether touching endpoints count as overlap depends on the problem, so decide that first and pick < or <= to match.",
				WhyItWorks: `Two intervals fail to overlap in exactly two ways: a lies entirely to the left of b (a.end <= b.start) or entirely to the right (b.end <= a.start). Overlap is 'neither of those', and negating both gives a.start < b.end and b.start < a.end. This covers every case at once, including one interval inside the other, which is easy to miss if you reason case by case.
Half-open intervals [start, end), like meetings, don't overlap when they only touch: [1,5) and [5,10) share no moment, hence the strict <. Closed intervals [start, end], like Merge Intervals or balloon ranges, do share the touching point, so the test becomes a.start <= b.end and b.start <= a.end.
The overlapping part, when there is one, is [max(a.start, b.start), min(a.end, b.end)]; the intervals overlap exactly when that range is non-empty. After sorting by start, b.start >= a.start already, so the test shrinks to the single comparison b.start < a.end.`,
				Intuition: "Two trains on the same track section: they collide unless one has fully left before the other enters. Mapping: entering is start, leaving is end, and 'one has fully left before the other enters' is a.end <= b.start or b.end <= a.start; overlap is everything else. Whether leaving at 5 and entering at 5 is a collision is the touching rule, set by the problem.",
				CoreIdeas: []string{
					"Overlap: a.start < b.end and b.start < a.end (use <= for closed intervals where touching counts).",
					"No overlap: one ends before the other starts.",
					"Intersection: [max of starts, min of ends], non-empty exactly when the intervals overlap.",
					"Sorted by start, only b.start < a.end needs checking.",
					"Read the statement for the touching rule: meetings ([1,5],[5,10] fine) versus Merge Intervals ([1,4],[4,5] merge).",
				},
				CommonMistakes: []string{
					"Checking only 'b starts inside a' (a.start <= b.start < a.end). It misses a starting inside b; use the two-sided test, or sort so the order is known.",
					"Mixing up < and <= for touching intervals. Decide from the statement: half-open meetings use <, closed ranges use <=. Test with a touching pair.",
					"Joining the two conditions with 'or' instead of 'and'. With 'or', almost every pair looks like it overlaps.",
				},
				Examples: []CodeExample{
					{
						ID:       "overlap_test",
						Title:    "Overlap and intersection",
						Language: "python",
						Snippet: `def overlaps(a, b):            # half-open: touching is NOT overlap
    return a[0] < b[1] and b[0] < a[1]

def overlaps_closed(a, b):     # closed: touching IS overlap
    return a[0] <= b[1] and b[0] <= a[1]

def intersection(a, b):        # closed intervals
    lo, hi = max(a[0], b[0]), min(a[1], b[1])
    return [lo, hi] if lo <= hi else None

# overlaps([1, 5], [5, 10]) -> False; overlaps_closed([1, 4], [4, 5]) -> True`,
						Notes: "The two-sided test handles containment ([1,10] and [3,4]) without a special case.",
					},
				},
			},
			{
				Name:        "Sort by start or sort by end?",
				Description: "Nearly every interval solution starts with a sort, and the key decides what the sweep can assume. Sort by start when you are building or checking blocks (merge, insert, clash detection, counting rooms). Sort by end when you are choosing which intervals to keep (most non-overlapping, fewest removals, fewest points to stab every interval).",
				WhyItWorks: `Sort by start: when you reach interval b, every interval before it started no later. So b can only reach to the right of what you've built, and everything you've already closed ended before the current block began. That makes 'compare with the last block only' correct, which is what Merge Intervals and Meeting Rooms rely on.
Sort by end: this is chapter 17's activity selection. When choosing intervals to keep, the one that ends first leaves the most room for the rest, and the exchange argument proves it: in any best selection, swap the earliest-ending interval in for the selection's first one, and nothing clashes. Sorting by end makes 'the next compatible interval' always the earliest-ending one.
Using the wrong key breaks things in specific ways. Selection sorted by start and keeping the first on a clash fails on [[1,100],[1,2],[2,3],[3,4]]: it keeps [1,100] and must remove 3 intervals instead of 1. Merging sorted by end fails on [[2,3],[1,10]]: [1,10] comes second but reaches to the left of the block [2,3], so 'compare with the last block' isn't enough.`,
				Intuition: "Filing papers versus booking a room. To merge overlapping papers into folders you go through them by start date, because each new one can only stick out to the right of the current folder. To fit the most bookings into one room you go by end time, because the booking that frees the room soonest leaves the most time for others. Mapping: folders are merged blocks (sort by start), bookings are selected intervals (sort by end).",
				CoreIdeas: []string{
					"Sort by start: merge, insert, detect clashes, count rooms. The sweep compares with the latest block only.",
					"Sort by end: keep the most non-overlapping intervals or remove the fewest. This is activity selection from chapter 17.",
					"Track the running end during the sweep: the end of the current block (by start) or of the last kept interval (by end).",
					"If the input is already sorted (Insert Interval), skip the sort and do the pass in O(n).",
				},
				CommonMistakes: []string{
					"Sorting by start for a selection problem and keeping the first interval on a clash. Sort by end, or on a clash keep the interval with the smaller end.",
					"Updating the running end with the new interval's end instead of the max, in a merge. A contained interval ([2,3] inside [1,4]) would shrink the block.",
					"Starting the running end at 0 when coordinates can be negative. Use float('-inf') or seed it with the first interval.",
				},
				Examples: []CodeExample{
					{
						ID:       "merge_by_start",
						Title:    "Merge (sort by start) versus select (sort by end)",
						Language: "python",
						Snippet: `def merge(intervals):                       # sort by START
    merged = []
    for start, end in sorted(intervals, key=lambda i: i[0]):
        if merged and start <= merged[-1][1]:
            merged[-1][1] = max(merged[-1][1], end)
        else:
            merged.append([start, end])
    return merged

def most_non_overlapping(intervals):        # sort by END
    kept, current_end = 0, float('-inf')
    for start, end in sorted(intervals, key=lambda i: i[1]):
        if start >= current_end:
            kept += 1
            current_end = end
    return kept`,
						Notes: "Same loop shape, different key and different meaning of the running end.",
					},
				},
			},
			{
				Name:        "Sweep line: counting overlaps",
				Description: "To find how many intervals are active at once (rooms needed, peak load), sweep across time and keep a count: +1 when an interval starts, -1 when one ends. The answer is the largest count. A min-heap of end times (chapter 11) does the same thing while also telling you which room frees up first.",
				WhyItWorks: `The number of active intervals only changes at a start or an end, so you only need to look at those moments. Sort all starts and all ends; walking them in time order, a start means one more active interval and an end means one fewer. The peak is the maximum overlap, and that maximum is exactly the number of rooms needed: k meetings running at the same moment need k rooms, and the greedy that reuses any free room never needs more.
For half-open meetings, an end at time t must be processed before a start at time t, so a room freed at 5 can be reused by a meeting starting at 5. In the two-array version that is the strict comparison start < ends[p] to open a new room; in the heap version it is reusing a room when heap[0] <= start.
You don't need to know which meeting ended, only that one did, which is why starts and ends can be sorted separately.`,
				Intuition: "A turnstile counter at a building entrance. Each person entering adds 1, each person leaving subtracts 1, and the most people inside at once is the highest reading. Mapping: people are meetings, entering is start, leaving is end, the highest reading is the number of rooms, and if someone leaves at the moment another enters the count shouldn't spike, so process the exit first.",
				CoreIdeas: []string{
					"Events: +1 at each start, -1 at each end; the answer is the peak running count.",
					"Ties at the same time: process ends before starts for half-open intervals.",
					"Two sorted arrays (starts and ends) and a pointer give the same result without building event pairs.",
					"A min-heap of end times keeps the earliest-free room on top: reuse it if it is free by the new start.",
				},
				CommonMistakes: []string{
					"Processing a start before an end at the same time, which counts touching meetings [1,5],[5,10] as needing 2 rooms.",
					"Returning the final count (usually 0 or the heap size after popping every finished room) instead of the maximum seen during the sweep.",
					"Pairing each start with its own end in the two-array version. The arrays are sorted independently on purpose; only the counts matter.",
				},
				Examples: []CodeExample{
					{
						ID:       "sweep_line",
						Title:    "Peak overlap with sorted events",
						Language: "python",
						Snippet: `def max_overlap(intervals):          # half-open intervals
    events = []
    for start, end in intervals:
        events.append((start, 1))
        events.append((end, -1))
    events.sort()                    # at equal times, -1 sorts before +1
    active = best = 0
    for _, change in events:
        active += change
        best = max(best, active)
    return best

# max_overlap([[0, 30], [5, 10], [15, 20]]) -> 2`,
						Notes: "Sorting (time, change) puts ends (-1) before starts (+1) at the same time, which is exactly the half-open rule.",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Interval list intersections",
				Problem: "Two lists of closed intervals are each sorted by start and have no overlaps within a list. Return every intersection of an interval from the first list with one from the second, in order. Example: first = [[0,2],[5,10],[13,23],[24,25]], second = [[1,5],[8,12],[15,24],[25,26]].",
				Steps: []WorkedStep{
					{
						Subgoal:     "Sort the intervals",
						Explanation: "Both lists are already sorted by start and disjoint, so no sort is needed. Use one pointer into each list, like merging two sorted lists (chapter 3's two pointers).",
						Code:        "result = []\ni = j = 0\nwhile i < len(first) and j < len(second):\n    a_start, a_end = first[i]\n    b_start, b_end = second[j]",
						State:       "i=0 -> [0,2], j=0 -> [1,5]",
					},
					{
						Subgoal:     "Test for overlap",
						Explanation: "The candidate intersection is [max of starts, min of ends]. Closed intervals overlap when it is non-empty, i.e. lo <= hi; touching endpoints like [5,10] and [1,5] do count, giving [5,5].",
						Code:        "lo = max(a_start, b_start)\nhi = min(a_end, b_end)",
						State:       "[0,2] vs [1,5]: lo=1, hi=2 -> overlap",
					},
					{
						Subgoal:     "Merge or keep",
						Explanation: "Record the intersection when it exists. Unlike Merge Intervals we keep the overlapping part, not the union.",
						Code:        "if lo <= hi:\n    result.append([lo, hi])",
						State:       "result = [[1,2]]",
					},
					{
						Subgoal:     "Track the running end",
						Explanation: "Advance whichever interval ends first. It can't intersect anything else in the other list: the other list's later intervals start after the current one, which already reaches past this end. This is a small exchange-style argument like chapter 17's: the earlier-ending interval is the one that is finished.",
						Code:        "if a_end < b_end:\n    i += 1\nelse:\n    j += 1",
						State:       "2 < 5 -> i=1 ([5,10]); [5,10] vs [1,5]: [5,5]; 10 > 5 -> j=1 ([8,12]); [5,10] vs [8,12]: [8,10]; i=2 ...",
					},
					{
						Subgoal:     "Return the answer",
						Explanation: "The loop stops when either list runs out, because nothing can intersect an empty list. Each step advances one pointer, so it is O(m + n).",
						Code:        "return result",
						State:       "result = [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]",
					},
				},
				Code: `def intervalIntersection(first, second):
    result = []
    i = j = 0
    while i < len(first) and j < len(second):
        a_start, a_end = first[i]
        b_start, b_end = second[j]
        lo = max(a_start, b_start)
        hi = min(a_end, b_end)
        if lo <= hi:                 # closed intervals: touching counts
            result.append([lo, hi])
        if a_end < b_end:            # the one that ends first is finished
            i += 1
        else:
            j += 1
    return result`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Minimum arrows to burst balloons",
				Prompt: "Each balloon covers the closed range [start, end] on a line. An arrow shot at x bursts every balloon with start <= x <= end. Return the fewest arrows that burst all balloons. Fill in the blanks: which sort key, what overlap test, and where the running end moves.",
				Code: `def findMinArrowShots(points):
    arrows = 0
    arrow_at = float('-inf')          # Track the running end: last arrow's position
    for start, end in sorted(points, key=____):   # Sort the intervals
        if ____:                      # Test for overlap: last arrow misses this balloon
            arrows += 1
            arrow_at = ____           # shoot as far right as this balloon allows
    return arrows`,
				Blanks:      []string{"lambda p: p[1]", "start > arrow_at", "end"},
				Explanation: "This is activity selection again, so sort by end. Shoot each new arrow at the end of the earliest-ending balloon it must hit: that position hits that balloon and as many later ones as possible (exchange argument: any other arrow for this balloon is at or left of its end, so moving it right to the end can only hit more). Balloons are closed intervals, so a balloon starting exactly at arrow_at is still hit; a new arrow is needed only when start > arrow_at. [[10,16],[2,8],[1,6],[7,12]] needs 2 arrows (at 6 and 12), and [[1,2],[2,3],[3,4],[4,5]] needs 2 (at 2 and 4).",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Meetings are half-open: [1,5] and [5,10] don't clash. Which test says meetings a and b overlap?",
				Choices: []string{
					"a.start < b.end and b.start < a.end",
					"a.start <= b.start and b.end <= a.end",
					"a.start < b.end or b.start < a.end",
					"a.end == b.start",
				},
				Answer:      0,
				Explanation: "Two intervals overlap unless one ends before the other starts; negating 'a.end <= b.start or b.end <= a.start' gives the first choice. The second only detects b inside a. With 'or' almost every pair passes. a.end == b.start is the touching case, which is not an overlap for meetings.",
			},
			{
				Prompt: "Which approach fits? 'Given meeting times, find the minimum number of rooms so every meeting has a room.'",
				Choices: []string{
					"Sort by end and keep each meeting that starts after the last kept one",
					"Sort by start and keep a min-heap of room end times, reusing the earliest-free room",
					"Sliding window over meetings sorted by length",
					"Merge overlapping meetings and return the number of merged blocks",
				},
				Answer:      1,
				Explanation: "Counting rooms is counting the peak overlap: sort by start and reuse the room that frees up first (heap, chapter 11), or sweep sorted starts and ends. Sort-by-end selection answers a different question (most meetings in one room). Merging gives blocks of busy time, and a block with 3 overlapping meetings still needs 3 rooms.",
			},
			{
				Prompt:      "Trace this merge on [[1,4],[2,3],[5,6],[6,8]]:\n\nmerged = []\nfor start, end in sorted(intervals, key=lambda i: i[0]):\n    if merged and start <= merged[-1][1]:\n        merged[-1][1] = max(merged[-1][1], end)\n    else:\n        merged.append([start, end])\n\nWhat is merged after each interval?",
				Explanation: "[1,4] -> [[1,4]]. [2,3]: 2 <= 4, end stays max(4,3) = 4 -> [[1,4]]. [5,6]: 5 > 4, new block -> [[1,4],[5,6]]. [6,8]: 6 <= 6 (closed intervals touch), end becomes 8 -> [[1,4],[5,8]]. Without the max, [2,3] would have shrunk the first block to [1,3].",
			},
			{
				Prompt:      "Prerequisite (chapter 17): Non-overlapping Intervals asks for the fewest removals. Why does keeping the interval that ends first never make the answer worse?",
				Explanation: "Exchange argument: take any largest set of non-overlapping intervals and its earliest-ending member o. The overall earliest-ending interval g ends no later than o, and every other member starts at or after o ends, so g clashes with none of them. Swapping o for g keeps the set valid and the same size, so a best answer exists that starts with g; repeat on the intervals that start after g ends.",
			},
			{
				Prompt:      "In your own words: when do you sort intervals by start, and when by end? Give one problem for each.",
				Explanation: "Sort by start when building or checking blocks, because each new interval can then only extend to the right of the current block: Merge Intervals, Meeting Rooms, Meeting Rooms II. Sort by end when choosing which intervals to keep, because the earliest-ending one leaves the most room (activity selection): Non-overlapping Intervals, minimum arrows to burst balloons.",
			},
			{
				Prompt: "A Meeting Rooms II solution reuses a room only when heap[0] < start (strict). What does it return for [[1,5],[5,10]]?",
				Choices: []string{
					"2, because the room freed at 5 isn't reused by the meeting starting at 5",
					"1, because the meetings don't overlap",
					"0, because the heap is emptied",
				},
				Answer:      0,
				Explanation: "With <, 5 < 5 is false, so the second meeting opens a new room and the answer is 2. The correct answer is 1: meetings are half-open, so a room freed at 5 is available at 5. Use heap[0] <= start.",
			},
		},
	}
}
