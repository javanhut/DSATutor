package practice

// IntervalsProblems contains all intervals category problems, in ladder order.
//
// Subgoal labels shared with chapter 18: "Sort the intervals",
// "Test for overlap", "Merge or keep", "Track the running end",
// "Return the answer".
var IntervalsProblems = []*Problem{
	{
		ID:              "meeting-rooms",
		Number:          252,
		Order:           1,
		Title:           "Meeting Rooms",
		Slug:            "meeting-rooms",
		Difficulty:      "Easy",
		Category:        "intervals",
		Tags:            []string{"Array", "Sorting", "Intervals"},
		RelatedChapters: []int{18, 17},
		Description: `You are given a list of meeting times, where intervals[i] = [start, end] means meeting i runs from time start up to time end. A meeting occupies the half-open time range [start, end): one meeting may start at the exact moment another ends, so [1,5] and [5,10] do **not** clash.

Return true if one person could attend every meeting (no two meetings overlap), and false otherwise. An empty list returns true.`,
		Constraints: []string{
			"0 <= intervals.length <= 10^4",
			"intervals[i].length == 2",
			"0 <= start < end <= 10^6",
		},
		Examples: []Example{
			{Input: "intervals = [[0,30],[5,10],[15,20]]", Output: "false", Explanation: "Sorted by start, [5,10] begins at 5, before [0,30] ends at 30, so those two clash."},
			{Input: "intervals = [[7,10],[2,4]]", Output: "true", Explanation: "Sorted: [2,4] then [7,10]. The second starts at 7, after the first ends at 4."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{0, 30}, {5, 10}, {15, 20}}}, Expected: false, Note: "example 1: a long meeting covers others"},
			{Input: map[string]interface{}{"intervals": [][]int{{7, 10}, {2, 4}}}, Expected: true, Note: "example 2: unsorted, no clash"},
			{Input: map[string]interface{}{"intervals": [][]int{}}, Expected: true, Hidden: true, Note: "empty list"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 5}}}, Expected: true, Hidden: true, Note: "single meeting"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 5}, {5, 10}}}, Expected: true, Hidden: true, Note: "touching meetings don't clash (catches <= instead of <)"},
			{Input: map[string]interface{}{"intervals": [][]int{{5, 10}, {1, 6}}}, Expected: false, Hidden: true, Note: "clash only visible after sorting"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 10}, {11, 12}, {3, 4}}}, Expected: false, Hidden: true, Note: "a meeting nested inside another"},
			{Input: map[string]interface{}{"intervals": [][]int{{6, 8}, {1, 3}, {3, 6}}}, Expected: true, Hidden: true, Note: "a chain of touching meetings"},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def canAttendMeetings(intervals):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Checking every pair of meetings works but is slow. If the meetings were listed in the order they start, which meetings would you need to compare each one with?"},
			{Level: 2, Type: "insight", Content: "After sorting by start time, if any two meetings clash then some neighbouring pair clashes: a meeting can only collide with a later-starting one if it collides with the very next one. So sort by start and compare each meeting's start with the previous meeting's end. Chapter 18 covers the overlap test and why sorting by start puts clashes next to each other.", ChapterRef: 18},
			{Level: 3, Type: "plan", Content: "1. Sort the intervals by start time.\n2. Test for overlap: for each meeting after the first, does it start before the previous one ends?\n3. Return the answer: false at the first clash, true if none is found.\nThink about touching meetings: is start == previous end a clash here?"},
			{Level: 4, Type: "skeleton", Content: "def canAttendMeetings(intervals):\n    meetings = sorted(intervals, key=____)\n    for i in range(1, len(meetings)):\n        if ____:\n            return False\n    return True"},
		},
		Solution: Solution{
			Code: `def canAttendMeetings(intervals):
    meetings = sorted(intervals, key=lambda meeting: meeting[0])
    for i in range(1, len(meetings)):
        if meetings[i][0] < meetings[i - 1][1]:
            return False  # starts before the previous meeting ends
    return True`,
			Explanation: "The first rung of the intervals ladder: sort by start, then compare neighbours. It uses the two tools every later problem needs, the overlap test and sorting by start, with nothing else on top.",
			BruteForce:  "Compare every pair with the overlap test a.start < b.end and b.start < a.end: n(n-1)/2 checks, about 5 * 10^7 for n = 10^4. The waste: most pairs are far apart in time, and after sorting only neighbours can be the first clash.",
			KeyInsight:  "Sorted by start, a meeting that clashes with any later meeting also clashes with the next one, so n - 1 neighbour checks replace all pairs.",
			WhyItWorks: `Sort by start. Suppose meetings i < j clash, so meetings[j].start < meetings[i].end. If no neighbouring pair between them clashes, then every meeting in between ends no later than the next one starts, so meetings[i].end <= meetings[i+1].start <= ... <= meetings[j].start, which contradicts the clash. So some neighbouring pair clashes, and checking neighbours finds it.
For sorted neighbours the general overlap test a.start < b.end and b.start < a.end simplifies: b.start >= a.start already, and b.start < b.end always, so only b.start < a.end remains. It's strict because a meeting ending at 5 and one starting at 5 don't share any moment.`,
			CommonMistakes: []string{
				"Using <= instead of <. Then [1,5],[5,10] counts as a clash, but one person can leave at 5 and walk into the next meeting.",
				"Forgetting to sort. On [[5,10],[1,6]] the neighbours in input order are not the neighbours in time, and an unsorted scan can miss clashes.",
			},
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort the intervals", Explanation: "Sorting by start lines meetings up in time order, so any clash shows up between neighbours. sorted() makes a copy and leaves the input alone.", CodeSnippet: "meetings = sorted(intervals, key=lambda meeting: meeting[0])", LineStart: 2, LineEnd: 2},
				{Title: "Test for overlap", Explanation: "With sorted neighbours the overlap test is just 'starts before the previous one ends'. Example 1 sorted is [0,30],[5,10],[15,20]; 5 < 30 is a clash.", CodeSnippet: "for i in range(1, len(meetings)):\n    if meetings[i][0] < meetings[i - 1][1]:\n        return False  # starts before the previous meeting ends", LineStart: 3, LineEnd: 5},
				{Title: "Return the answer", Explanation: "No neighbouring clash means no clash at all. An empty or one-meeting list never enters the loop and returns true.", CodeSnippet: "return True", LineStart: 6, LineEnd: 6},
			},
			Alternatives: []AltSolution{
				{
					Name: "Check every pair (brute force)",
					Code: `def canAttendMeetings(intervals):
    for i in range(len(intervals)):
        for j in range(i + 1, len(intervals)):
            a, b = intervals[i], intervals[j]
            if a[0] < b[1] and b[0] < a[1]:
                return False
    return True`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Applies the full two-sided overlap test to every pair. It needs no sorting, which makes it a good way to check the sorted version.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it enough to compare each meeting only with the one just before it in sorted order?", Answer: "If meeting i clashes with some later meeting j but no neighbours in between clash, then each meeting ends by the time the next starts, so i ends by the time j starts, contradicting the clash. So the earliest clash is always between neighbours."},
			{Prompt: "What changes if meetings were closed intervals, so [1,5] and [5,10] did clash?", Answer: "Only the comparison: use meetings[i][0] <= meetings[i - 1][1]. Whether endpoints touch counts as overlap is part of the problem statement, so read it before choosing < or <=."},
		},
	},
	{
		ID:              "merge-intervals",
		Number:          56,
		Order:           2,
		Title:           "Merge Intervals",
		Slug:            "merge-intervals",
		Difficulty:      "Medium",
		Category:        "intervals",
		Tags:            []string{"Array", "Sorting", "Intervals"},
		RelatedChapters: []int{18, 17},
		Description: `You are given a list of intervals, where intervals[i] = [start, end] with start <= end. Two intervals overlap if they share at least one point; here intervals are closed, so [1,4] and [4,5] overlap at 4.

Merge every group of overlapping intervals into one interval that covers exactly the same points, and return the resulting list of non-overlapping intervals. The intervals in the answer may be in any order.`,
		Constraints: []string{
			"1 <= intervals.length <= 10^4",
			"intervals[i].length == 2",
			"0 <= start <= end <= 10^4",
		},
		Examples: []Example{
			{Input: "intervals = [[1,3],[2,6],[8,10],[15,18]]", Output: "[[1,6],[8,10],[15,18]]", Explanation: "[1,3] and [2,6] share 2..3, so they merge into [1,6]. The others don't touch anything."},
			{Input: "intervals = [[1,4],[4,5]]", Output: "[[1,5]]", Explanation: "They share the point 4, so they merge."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, Expected: [][]int{{1, 6}, {8, 10}, {15, 18}}, Note: "example 1"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 4}, {4, 5}}}, Expected: [][]int{{1, 5}}, Note: "example 2: touching intervals merge"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 4}}}, Expected: [][]int{{1, 4}}, Hidden: true, Note: "single interval"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 4}, {2, 3}}}, Expected: [][]int{{1, 4}}, Hidden: true, Note: "contained interval (catches end = new end without max)"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 10}, {2, 3}, {4, 5}}}, Expected: [][]int{{1, 10}}, Hidden: true, Note: "several intervals inside one long one"},
			{Input: map[string]interface{}{"intervals": [][]int{{8, 10}, {1, 3}, {2, 6}}}, Expected: [][]int{{1, 6}, {8, 10}}, Hidden: true, Note: "unsorted input"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 4}, {0, 4}}}, Expected: [][]int{{0, 4}}, Hidden: true, Note: "the later interval in the input starts first"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {3, 4}}}, Expected: [][]int{{1, 2}, {3, 4}}, Hidden: true, Note: "a gap of one: nothing merges"},
			{Input: map[string]interface{}{"intervals": [][]int{{2, 2}, {2, 2}}}, Expected: [][]int{{2, 2}}, Hidden: true, Note: "duplicate single-point intervals"},
			{Input: map[string]interface{}{"intervals": [][]int{{5, 6}, {1, 2}, {2, 3}, {7, 8}, {3, 4}}}, Expected: [][]int{{1, 4}, {5, 6}, {7, 8}}, Hidden: true, Note: "a chain of touching intervals"},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def merge(intervals):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In Meeting Rooms you sorted by start and compared neighbours. If you sort the same way here, which already-merged interval could the next interval possibly overlap?"},
			{Level: 2, Type: "insight", Content: "After sorting by start, the next interval can only overlap the most recently merged block: every earlier block ended before this one began. So keep a list of merged blocks and compare each interval with the last block only. If it starts at or before that block's end, extend the end (with max, since it may be contained); otherwise start a new block. See chapter 18.", ChapterRef: 18},
			{Level: 3, Type: "plan", Content: "1. Sort the intervals by start.\n2. Track the running end: the last block in the result.\n3. Test for overlap: does the next interval start at or before that block's end?\n4. Merge or keep: if yes, set the block's end to the max of the two ends; if no, append it as a new block.\n5. Return the answer: the list of blocks."},
			{Level: 4, Type: "skeleton", Content: "def merge(intervals):\n    merged = []\n    for start, end in sorted(intervals, key=lambda interval: interval[0]):\n        if merged and ____:\n            merged[-1][1] = ____\n        else:\n            merged.append([start, end])\n    return merged"},
		},
		Solution: Solution{
			Code: `def merge(intervals):
    merged = []
    for start, end in sorted(intervals, key=lambda interval: interval[0]):
        if merged and start <= merged[-1][1]:
            merged[-1][1] = max(merged[-1][1], end)  # overlaps: extend
        else:
            merged.append([start, end])  # gap: start a new block
    return merged`,
			Explanation: "Builds on Meeting Rooms: same sort by start, same 'compare with what came just before', but now instead of stopping at the first overlap we absorb it into a growing block. The new idea is the running end, merged[-1][1], which has to be a max because a short interval can sit entirely inside a long one.",
			BruteForce:  "Repeatedly look for any two overlapping intervals, replace them with their union, and start over until no pair overlaps: O(n^2) per pass and up to n passes. The waste: after sorting, an interval can only overlap the latest block, so all other comparisons are pointless.",
			KeyInsight:  "Sorted by start, each interval either overlaps the last merged block or starts after everything so far; there's never a need to look further back.",
			WhyItWorks: `Invariant: after processing the first k sorted intervals, merged holds the correct merge of exactly those intervals, its blocks are disjoint and in order, and the last block has the largest end seen so far.
Step: the next interval [s, e] has s >= every start seen so far. Every block before the last one ended before the last block began, so it ended before s too: [s, e] can't touch it. So [s, e] either overlaps the last block (s <= its end), and the union of the two is [block start, max(block end, e)], or it starts after every point covered so far and is a new block.
The end must be max(block end, e), not e: in [[1,4],[2,3]] the second interval is inside the first, and setting the end to 3 would lose the points 3..4.`,
			CommonMistakes: []string{
				"Setting merged[-1][1] = end instead of max(...). A contained interval like [2,3] inside [1,4] would shrink the block to [1,3].",
				"Using < instead of <= for the overlap test. These intervals are closed, so [1,4] and [4,5] share the point 4 and must merge into [1,5].",
				"Not sorting first, or sorting by end. With [[1,4],[0,4]] in input order, [0,4] arrives after the block [1,4] and the code would either miss the merge or produce [[1,4],[0,4]].",
			},
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort the intervals", Explanation: "Sort by start so each interval can only overlap the latest block. Unpacking start, end makes the test below readable.", CodeSnippet: "for start, end in sorted(intervals, key=lambda interval: interval[0]):", LineStart: 3, LineEnd: 3},
				{Title: "Test for overlap", Explanation: "merged[-1][1] is the running end of the current block. Closed intervals overlap when start <= that end. Example 1: [2,6] starts at 2 <= 3, so it overlaps [1,3].", CodeSnippet: "if merged and start <= merged[-1][1]:", LineStart: 4, LineEnd: 4},
				{Title: "Merge or keep", Explanation: "Overlap: extend the block to the larger end ([1,3] becomes [1,6]). Gap: start a new block ([8,10] starts after 6).", CodeSnippet: "merged[-1][1] = max(merged[-1][1], end)  # overlaps: extend\nelse:\n    merged.append([start, end])  # gap: start a new block", LineStart: 5, LineEnd: 7},
				{Title: "Return the answer", Explanation: "The blocks are disjoint and come out sorted by start as a side effect.", CodeSnippet: "return merged", LineStart: 8, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{
					Name: "Merge any overlapping pair until none is left (brute force)",
					Code: `def merge(intervals):
    blocks = [list(interval) for interval in intervals]
    changed = True
    while changed:
        changed = False
        for i in range(len(blocks)):
            for j in range(i + 1, len(blocks)):
                a, b = blocks[i], blocks[j]
                if a[0] <= b[1] and b[0] <= a[1]:
                    blocks[i] = [min(a[0], b[0]), max(a[1], b[1])]
                    blocks.pop(j)
                    changed = True
                    break
            if changed:
                break
    return blocks`,
					TimeComplexity:  "O(n^3)",
					SpaceComplexity: "O(n)",
					Explanation:     "Uses the two-sided closed-interval overlap test on every pair and merges one pair at a time. No sorting, so it is an independent check, but each merge restarts the search.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why do you only ever compare the new interval with the last merged block?", Answer: "Intervals arrive in order of start. Every earlier block ended before the last block started, and the last block started no later than the new interval, so those earlier blocks end before the new interval begins and can't overlap it."},
			{Prompt: "Would sorting by end instead of start work here?", Answer: "No. With [[2,3],[1,10]] sorted by end, [2,3] comes first and then [1,10] starts before the block began; the 'compare with the last block' test only extends the end, so you'd need to fix the start too and could need to merge backwards into several blocks. Sorting by start guarantees new intervals never reach to the left of the current block."},
		},
	},
	{
		ID:              "insert-interval",
		Number:          57,
		Order:           3,
		Title:           "Insert Interval",
		Slug:            "insert-interval",
		Difficulty:      "Medium",
		Category:        "intervals",
		Tags:            []string{"Array", "Intervals"},
		RelatedChapters: []int{18},
		Description: `You are given a list of intervals that are already **sorted by start and non-overlapping**, and one more interval newInterval = [start, end]. Intervals are closed, so [1,5] and [5,7] overlap at 5.

Insert newInterval so that the list is still sorted by start and has no overlapping intervals, merging newInterval with any intervals it overlaps. Return the new list. The input list may be empty.`,
		Constraints: []string{
			"0 <= intervals.length <= 10^4",
			"intervals[i].length == 2 and 0 <= start <= end <= 10^5",
			"intervals is sorted by start and has no overlaps",
			"newInterval.length == 2 and 0 <= start <= end <= 10^5",
		},
		Examples: []Example{
			{Input: "intervals = [[1,3],[6,9]], newInterval = [2,5]", Output: "[[1,5],[6,9]]", Explanation: "[2,5] overlaps [1,3], giving [1,5]. [6,9] starts after 5, so it is copied unchanged."},
			{Input: "intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]", Output: "[[1,2],[3,10],[12,16]]", Explanation: "[1,2] ends before 4: copy it. [3,5], [6,7] and [8,10] all overlap the growing interval, which becomes [3,10]. [12,16] starts after 10: copy it."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{1, 3}, {6, 9}}, "newInterval": []int{2, 5}}, Expected: [][]int{{1, 5}, {6, 9}}, Note: "example 1"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, "newInterval": []int{4, 8}}, Expected: [][]int{{1, 2}, {3, 10}, {12, 16}}, Note: "example 2: absorbs three intervals"},
			{Input: map[string]interface{}{"intervals": [][]int{}, "newInterval": []int{5, 7}}, Expected: [][]int{{5, 7}}, Hidden: true, Note: "empty list"},
			{Input: map[string]interface{}{"intervals": [][]int{{3, 5}}, "newInterval": []int{1, 2}}, Expected: [][]int{{1, 2}, {3, 5}}, Hidden: true, Note: "goes before everything"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}}, "newInterval": []int{5, 6}}, Expected: [][]int{{1, 2}, {5, 6}}, Hidden: true, Note: "goes after everything (catches forgetting to append it)"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {6, 7}}, "newInterval": []int{3, 4}}, Expected: [][]int{{1, 2}, {3, 4}, {6, 7}}, Hidden: true, Note: "fits in a gap without touching"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 5}}, "newInterval": []int{5, 7}}, Expected: [][]int{{1, 7}}, Hidden: true, Note: "touching at one point merges"},
			{Input: map[string]interface{}{"intervals": [][]int{{2, 3}, {4, 5}}, "newInterval": []int{1, 10}}, Expected: [][]int{{1, 10}}, Hidden: true, Note: "new interval covers everything"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 10}}, "newInterval": []int{3, 4}}, Expected: [][]int{{1, 10}}, Hidden: true, Note: "new interval inside an existing one"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def insert(intervals, newInterval):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You could append newInterval and run Merge Intervals again, but the list is already sorted. Relative to newInterval, what three kinds of existing intervals are there, and in what order do they appear?"},
			{Level: 2, Type: "insight", Content: "Because the list is sorted and disjoint, it splits into three runs: intervals that end before newInterval starts, intervals that overlap it, and intervals that start after it ends. Copy the first run, fold the middle run into newInterval with min/max, append it, then copy the rest. One pass, no sorting. Chapter 18 covers the overlap test.", ChapterRef: 18},
			{Level: 3, Type: "plan", Content: "1. Sort the intervals: already done for you, which is what makes one pass enough.\n2. Test for overlap, part 1: while the current interval ends before newInterval starts, keep it as is.\n3. Merge or keep: while the current interval starts at or before newInterval's end, absorb it (min of starts, max of ends).\n4. Append the merged newInterval.\n5. Return the answer: append whatever is left and return."},
			{Level: 4, Type: "skeleton", Content: "def insert(intervals, newInterval):\n    result = []\n    new_start, new_end = newInterval\n    i = 0\n    n = len(intervals)\n    while i < n and ____:\n        result.append(intervals[i])\n        i += 1\n    while i < n and ____:\n        new_start = ____\n        new_end = ____\n        i += 1\n    result.append([new_start, new_end])\n    while i < n:\n        result.append(intervals[i])\n        i += 1\n    return result"},
		},
		Solution: Solution{
			Code: `def insert(intervals, newInterval):
    result = []
    new_start, new_end = newInterval
    i = 0
    n = len(intervals)
    while i < n and intervals[i][1] < new_start:
        result.append(intervals[i])  # ends before the new one starts
        i += 1
    while i < n and intervals[i][0] <= new_end:
        new_start = min(new_start, intervals[i][0])  # overlaps: absorb it
        new_end = max(new_end, intervals[i][1])
        i += 1
    result.append([new_start, new_end])
    while i < n:
        result.append(intervals[i])  # starts after the new one ends
        i += 1
    return result`,
			Explanation: "Builds on Merge Intervals: the same overlap test and the same min/max absorption, but the input is already sorted, so we skip the sort and do the merge in one O(n) pass. The new idea is splitting the list into before, overlapping and after runs.",
			BruteForce:  "Append newInterval and run Merge Intervals: sort O(n log n), then merge. Correct, but it re-sorts a list that was already sorted and re-checks intervals that can't possibly change.",
			KeyInsight:  "In a sorted, disjoint list the intervals overlapping newInterval are one contiguous run. Everything before it ends before new_start and everything after it starts after new_end.",
			WhyItWorks: `Closed intervals [a, b] and [s, e] overlap iff a <= e and s <= b. In the sorted, disjoint list, starts and ends both increase.
Run 1: while intervals[i].end < new_start, the interval is entirely to the left, so it doesn't overlap. Once one ends at or after new_start, every later one does too (ends increase).
Run 2: those intervals already satisfy end >= new_start; while start <= new_end they overlap the growing interval, so absorb them. Once a start is > new_end, every later start is larger, so nothing else overlaps. Growing new_end can't pull in extra intervals either: new_end only grows past the original end when an absorbed interval sticks out to the right, and because the input is disjoint the next interval starts after that one ends.
Run 3: the rest start after new_end; copy them. Each run keeps sorted order, and the merged interval sits between runs 1 and 3, so the output is sorted and disjoint.`,
			CommonMistakes: []string{
				"Forgetting to append the merged newInterval when no interval overlaps it, or when it goes after everything. Append it unconditionally after the middle loop.",
				"Using <= in the first loop (intervals[i][1] <= new_start). Then [1,5] with new [5,7] is copied unmerged and the answer has two touching intervals; closed intervals touching at 5 must merge into [1,7].",
				"Crashing on an empty list by reading intervals[0]. Use while-loop guards (i < n) so the empty case just appends newInterval.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track the running end", Explanation: "new_start and new_end describe the interval being inserted; they will grow as overlapping intervals are absorbed.", CodeSnippet: "new_start, new_end = newInterval", LineStart: 3, LineEnd: 3},
				{Title: "Test for overlap", Explanation: "Intervals ending before new_start are entirely to the left: keep them unchanged. Example 2: [1,2] ends at 2 < 4.", CodeSnippet: "while i < n and intervals[i][1] < new_start:\n    result.append(intervals[i])  # ends before the new one starts\n    i += 1", LineStart: 6, LineEnd: 8},
				{Title: "Merge or keep", Explanation: "Intervals starting at or before new_end overlap: widen the new interval to cover them. Example 2: [3,5] makes it [3,8], [6,7] keeps it [3,8], [8,10] makes it [3,10]; [12,16] starts after 10 and stops the loop.", CodeSnippet: "while i < n and intervals[i][0] <= new_end:\n    new_start = min(new_start, intervals[i][0])  # overlaps: absorb it\n    new_end = max(new_end, intervals[i][1])\n    i += 1", LineStart: 9, LineEnd: 12},
				{Title: "Merge or keep", Explanation: "The merged interval goes in exactly once, between the left run and the right run.", CodeSnippet: "result.append([new_start, new_end])", LineStart: 13, LineEnd: 13},
				{Title: "Return the answer", Explanation: "Everything left starts after new_end, so copy it unchanged.", CodeSnippet: "while i < n:\n    result.append(intervals[i])  # starts after the new one ends\n    i += 1\nreturn result", LineStart: 14, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{
					Name: "Append and re-merge",
					Code: `def insert(intervals, newInterval):
    merged = []
    for start, end in sorted(intervals + [newInterval]):
        if merged and start <= merged[-1][1]:
            merged[-1][1] = max(merged[-1][1], end)
        else:
            merged.append([start, end])
    return merged`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Reuses Merge Intervals directly. Simple and correct, but it pays for a sort the input didn't need.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Absorbing an interval can push new_end to the right. Why can't that make the merged interval overlap something in the third run?", Answer: "new_end only grows when an absorbed interval ends past it, and that can only be the last absorbed one. The input intervals are disjoint and sorted, so the next interval starts strictly after that interval's end, which is the new new_end. So the third run still starts after the merged interval. In example 2, [8,10] raises new_end to 10, and [12,16] starts at 12 > 10."},
			{Prompt: "If you had to insert k new intervals one at a time, what would that cost, and what would you do instead?", Answer: "Each insert is O(n), so k inserts cost O(k * n). If you get them all at once, append them all and run Merge Intervals once in O((n + k) log(n + k)). If they arrive over time, keep the intervals in a balanced tree keyed by start so each insert only touches its neighbours."},
		},
	},
	{
		ID:              "non-overlapping-intervals",
		Number:          435,
		Order:           4,
		Title:           "Non-overlapping Intervals",
		Slug:            "non-overlapping-intervals",
		Difficulty:      "Medium",
		Category:        "intervals",
		Tags:            []string{"Array", "Greedy", "Sorting", "Intervals"},
		RelatedChapters: []int{18, 17, 15},
		Description: `You are given a list of intervals, where intervals[i] = [start, end] with start < end. Intervals that only touch, such as [1,2] and [2,3], do **not** overlap.

Return the minimum number of intervals you must remove so that the remaining intervals don't overlap each other.`,
		Constraints: []string{
			"1 <= intervals.length <= 10^5",
			"intervals[i].length == 2",
			"-5 * 10^4 <= start < end <= 5 * 10^4",
		},
		Examples: []Example{
			{Input: "intervals = [[1,2],[2,3],[3,4],[1,3]]", Output: "1", Explanation: "Remove [1,3]; [1,2], [2,3], [3,4] only touch. Keeping [1,3] would clash with both [1,2] and [2,3], so removing it is best."},
			{Input: "intervals = [[1,2],[1,2],[1,2]]", Output: "2", Explanation: "Only one copy can stay."},
			{Input: "intervals = [[1,2],[2,3]]", Output: "0", Explanation: "They touch but don't overlap."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}}, Expected: 1, Note: "example 1"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {1, 2}, {1, 2}}}, Expected: 2, Note: "example 2: duplicates"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {2, 3}}}, Expected: 0, Note: "example 3: touching is fine"},
			{Input: map[string]interface{}{"intervals": [][]int{{0, 5}}}, Expected: 0, Hidden: true, Note: "single interval"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 100}, {1, 2}, {2, 3}, {3, 4}}}, Expected: 1, Hidden: true, Note: "one long interval (catches sorting by start and keeping the first)"},
			{Input: map[string]interface{}{"intervals": [][]int{{-10, -5}, {-4, -2}}}, Expected: 0, Hidden: true, Note: "negative coordinates (catches starting current_end at 0)"},
			{Input: map[string]interface{}{"intervals": [][]int{{-5, -3}, {-4, -1}}}, Expected: 1, Hidden: true, Note: "negative coordinates that overlap"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 5}, {2, 3}, {3, 4}, {4, 6}}}, Expected: 1, Hidden: true, Note: "short intervals inside a long one"},
			{Input: map[string]interface{}{"intervals": [][]int{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}}, Expected: 2, Hidden: true, Note: "staircase of overlaps"},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def eraseOverlapIntervals(intervals):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Removing as few as possible is the same as keeping as many as possible. If you could keep only one interval to start your schedule, which one would leave the most room for the others?"},
			{Level: 2, Type: "insight", Content: "This is activity selection from chapter 17: sort by END time and keep each interval that starts at or after the end of the last one kept. The interval that ends first leaves the most room, and an exchange argument shows some best answer starts with it. Answer = total - kept.", ChapterRef: 17},
			{Level: 3, Type: "plan", Content: "1. Sort the intervals by end time.\n2. Track the running end: the end of the last kept interval, starting at minus infinity (coordinates can be negative).\n3. Test for overlap: does the next interval start at or after that end?\n4. Merge or keep: if yes, keep it and move the running end to its end; if no, it is removed.\n5. Return the answer: total minus kept."},
			{Level: 4, Type: "skeleton", Content: "def eraseOverlapIntervals(intervals):\n    kept = 0\n    current_end = ____\n    for start, end in sorted(intervals, key=____):\n        if ____:\n            kept += 1\n            current_end = end\n    return len(intervals) - kept"},
		},
		Solution: Solution{
			Code: `def eraseOverlapIntervals(intervals):
    kept = 0
    current_end = float('-inf')  # end of the last interval kept
    for start, end in sorted(intervals, key=lambda interval: interval[1]):
        if start >= current_end:
            kept += 1  # compatible with everything kept: keep it
            current_end = end
    return len(intervals) - kept`,
			Explanation: "The first interval problem sorted by end instead of start. It is chapter 17's activity selection in disguise: keep the most non-overlapping intervals, and the answer is how many you had to drop. It builds on Meeting Rooms' neighbour test, but now a clash means 'drop one' instead of 'stop'.",
			BruteForce:  "Try every subset of intervals, keep those that are non-overlapping, and take the largest: O(2^n). A DP like longest increasing subsequence (sort by start, best[i] = 1 + max best[j] over j that end by start i) is O(n^2), 10^10 steps for n = 10^5. The waste: it considers every earlier interval as the predecessor, while the exchange argument shows the earliest-ending one is always a safe choice.",
			KeyInsight:  "Among intervals compatible with what you have kept, the one that ends first is always safe to keep: it leaves the most time for everything after it.",
			WhyItWorks: `Exchange argument. Sort by end and let g be the interval that ends first. Take any largest set O of non-overlapping intervals, and let o be the interval in O that ends first. Swap o for g: g ends no later than o, and every other interval in O starts at or after o ends, so also at or after g ends. The swapped set is still non-overlapping and has the same size, so some optimal set contains g.
After keeping g, the remaining problem is the same question on the intervals that start at or after g.end, and the same argument applies again. By induction the greedy keeps as many as any optimal set, so len - kept is the fewest removals.
Greedy stays ahead, the same fact said differently: after keeping k intervals, the greedy's current_end is <= the end of the k-th interval of any valid set.
start >= current_end is not strict because touching intervals are allowed here.`,
			CommonMistakes: []string{
				"Sorting by start and keeping the first interval on a clash. On [[1,100],[1,2],[2,3],[3,4]] that keeps [1,100] and removes 3, but removing just [1,100] is enough. Sort by end, or when sorting by start keep the one with the smaller end.",
				"Starting current_end at 0. Coordinates can be negative, so [[-10,-5],[-4,-2]] would treat both as clashing with an imaginary interval ending at 0. Start at float('-inf'), or keep the first sorted interval explicitly.",
				"Using > instead of >= for compatibility. [1,2] and [2,3] only touch, so the second must be kept; with > example 3 returns 1 instead of 0.",
			},
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track the running end", Explanation: "current_end is when the last kept interval finishes. Minus infinity means nothing is kept yet, which works for negative coordinates too.", CodeSnippet: "kept = 0\ncurrent_end = float('-inf')  # end of the last interval kept", LineStart: 2, LineEnd: 3},
				{Title: "Sort the intervals", Explanation: "Sort by end, so the next compatible interval we meet is always the one that finishes earliest. Example 1 sorted by end: [1,2],[2,3],[1,3],[3,4].", CodeSnippet: "for start, end in sorted(intervals, key=lambda interval: interval[1]):", LineStart: 4, LineEnd: 4},
				{Title: "Test for overlap", Explanation: "An interval is compatible if it starts at or after current_end. Example 1: [1,2] kept (end 2), [2,3] kept (end 3), [1,3] starts at 1 < 3 so it is removed, [3,4] kept.", CodeSnippet: "if start >= current_end:", LineStart: 5, LineEnd: 5},
				{Title: "Merge or keep", Explanation: "Keep it and move the running end. Intervals we skip are the removed ones.", CodeSnippet: "kept += 1  # compatible with everything kept: keep it\ncurrent_end = end", LineStart: 6, LineEnd: 7},
				{Title: "Return the answer", Explanation: "Removed = total - kept. Example 1: 4 - 3 = 1.", CodeSnippet: "return len(intervals) - kept", LineStart: 8, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{
					Name: "Longest chain DP (O(n^2))",
					Code: `def eraseOverlapIntervals(intervals):
    ordered = sorted(intervals)
    n = len(ordered)
    longest = [1] * n  # most intervals kept, ending with interval i
    for i in range(n):
        for j in range(i):
            if ordered[j][1] <= ordered[i][0]:
                longest[i] = max(longest[i], longest[j] + 1)
    return n - max(longest)`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n)",
					Explanation:     "The chapter 15 approach, like longest increasing subsequence: longest[i] is the largest non-overlapping chain ending with interval i. It tries every predecessor, which is exactly the work the exchange argument proves unnecessary.",
				},
				{
					Name:            "Sort by start, drop the longer end",
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Sort by start; on a clash, count a removal and keep whichever of the two intervals ends earlier (prev_end = min(prev_end, end)). It is the same exchange argument applied pairwise.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can keeping the interval that ends first never make the answer worse? Explain it in your own words.", Answer: "Take any best set and swap its earliest-ending interval for the overall earliest-ending one. The new interval ends no later, so it can't clash with the rest of the set, which all start after the old one ended. The set stays the same size, so a best answer exists that agrees with the greedy's first choice; repeat for the rest."},
			{Prompt: "Why doesn't 'keep the shortest interval first' work?", Answer: "A short interval can sit across the boundary of two others. With [1,5],[4,7],[6,10], the shortest is [4,7], which clashes with both others, so you keep 1. Earliest end keeps [1,5] and [6,10], which is 2. Length is not what decides how much room is left; the end time is."},
		},
	},
	{
		ID:              "meeting-rooms-ii",
		Number:          253,
		Order:           5,
		Title:           "Meeting Rooms II",
		Slug:            "meeting-rooms-ii",
		Difficulty:      "Medium",
		Category:        "intervals",
		Tags:            []string{"Array", "Sorting", "Heap (Priority Queue)", "Intervals"},
		RelatedChapters: []int{18, 11, 17},
		Description: `You are given a list of meetings, where intervals[i] = [start, end] means meeting i runs during the half-open time range [start, end). A room is free again at the exact moment its meeting ends, so a meeting ending at 5 and one starting at 5 can use the same room.

Return the minimum number of conference rooms needed to hold every meeting.`,
		Constraints: []string{
			"1 <= intervals.length <= 10^4",
			"intervals[i].length == 2",
			"0 <= start < end <= 10^6",
		},
		Examples: []Example{
			{Input: "intervals = [[0,30],[5,10],[15,20]]", Output: "2", Explanation: "[0,30] needs a room for the whole time. [5,10] needs a second room; at 15 the second room is free again, so [15,20] reuses it."},
			{Input: "intervals = [[7,10],[2,4]]", Output: "1", Explanation: "The meetings don't overlap, so one room is enough."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{0, 30}, {5, 10}, {15, 20}}}, Expected: 2, Note: "example 1"},
			{Input: map[string]interface{}{"intervals": [][]int{{7, 10}, {2, 4}}}, Expected: 1, Note: "example 2: no overlap"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 5}}}, Expected: 1, Hidden: true, Note: "single meeting"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 5}, {5, 10}}}, Expected: 1, Hidden: true, Note: "touching meetings share a room (catches < instead of <=)"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 10}, {2, 9}, {3, 8}}}, Expected: 3, Hidden: true, Note: "nested meetings all overlap"},
			{Input: map[string]interface{}{"intervals": [][]int{{9, 10}, {4, 9}, {4, 17}}}, Expected: 2, Hidden: true, Note: "unsorted, with a room freed at 9"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}}}, Expected: 4, Hidden: true, Note: "identical meetings"},
			{Input: map[string]interface{}{"intervals": [][]int{{2, 7}, {3, 4}, {5, 6}}}, Expected: 2, Hidden: true, Note: "a room is reused inside a long meeting"},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 4}, {2, 5}, {4, 6}, {5, 7}}}, Expected: 2, Hidden: true, Note: "rooms freed exactly when the next meetings start"},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def minMeetingRooms(intervals):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Process meetings in order of start time. When a new meeting starts, which of the rooms already in use is the only one worth checking to see if it's free?"},
			{Level: 2, Type: "insight", Content: "Only the room whose meeting ends earliest matters: if even that room is still busy, every room is. Keep the end times of the rooms in a min-heap (chapter 11). For each meeting sorted by start: if the smallest end time is <= the start, that room is free, so pop it; then push this meeting's end. The heap's size is the number of rooms. See chapter 18 for the sweep-line view.", ChapterRef: 18},
			{Level: 3, Type: "plan", Content: "1. Sort the intervals by start time.\n2. Track the running end of every room: a min-heap of end times.\n3. Test for overlap: is the earliest end time <= this meeting's start?\n4. Merge or keep: if so, reuse that room (pop it); either way push this meeting's end.\n5. Return the answer: the heap size."},
			{Level: 4, Type: "skeleton", Content: "import heapq\n\ndef minMeetingRooms(intervals):\n    end_times = []\n    for start, end in sorted(intervals, key=lambda meeting: meeting[0]):\n        if end_times and ____:\n            heapq.heappop(end_times)\n        heapq.heappush(end_times, ____)\n    return ____"},
		},
		Solution: Solution{
			Code: `import heapq

def minMeetingRooms(intervals):
    end_times = []  # min-heap: when each room in use becomes free
    for start, end in sorted(intervals, key=lambda meeting: meeting[0]):
        if end_times and end_times[0] <= start:
            heapq.heappop(end_times)  # earliest room is free: reuse it
        heapq.heappush(end_times, end)
    return len(end_times)`,
			Explanation: "The top rung: Meeting Rooms asked whether one room is enough; now we count rooms. It keeps the sort by start from the earlier rungs, but instead of one running end it tracks one end per room, in a min-heap (chapter 11), so the earliest-free room is always at the top.",
			BruteForce:  "The answer is the largest number of meetings happening at the same moment, and that maximum is reached at some meeting's start. For each meeting, count how many meetings are running at its start: O(n^2), 10^8 checks for n = 10^4. The waste: each count rescans all meetings instead of updating a running count as time moves forward.",
			KeyInsight:  "When a meeting starts, only the room that frees up earliest is worth checking. A min-heap of end times gives you that room in O(log n).",
			WhyItWorks: `The number of rooms the algorithm opens equals the largest number of meetings that overlap at one moment, and no schedule can do better than that.
Lower bound: if k meetings are all running at time t, they need k different rooms.
Greedy matches it: the heap grows only when a meeting starts at time s and the earliest end in the heap is > s. Every room in the heap holds a meeting that started at or before s (we process by start) and ends after s, so all of them plus the new one are running at time s. So the heap size never exceeds the true maximum overlap, and it never undercounts because each meeting gets a room that is actually free.
Why <= and not <: meetings are half-open, so a room whose meeting ends at 5 is free for a meeting starting at 5. With < the code opens a second room for [1,5],[5,10].
We pop at most one room per meeting on purpose: the heap size is the number of rooms ever opened, so it should only grow, never shrink.`,
			CommonMistakes: []string{
				"Using < instead of <= when checking the earliest end. Touching meetings [1,5],[5,10] then need 2 rooms instead of 1.",
				"Popping every finished room with a while loop and returning the final heap size. The heap then holds only meetings still running at the end, not the maximum; either pop one at a time as here, or track max(len(heap)) along the way.",
				"Pushing the start time instead of the end time, or forgetting to sort by start. The heap must answer 'when does each room become free', and meetings must arrive in time order for that answer to mean anything.",
			},
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track the running end", Explanation: "end_times holds one entry per room: the time it becomes free. The smallest is at end_times[0].", CodeSnippet: "end_times = []  # min-heap: when each room in use becomes free", LineStart: 4, LineEnd: 4},
				{Title: "Sort the intervals", Explanation: "Handle meetings in the order they start, so when a meeting arrives, every meeting already placed started earlier.", CodeSnippet: "for start, end in sorted(intervals, key=lambda meeting: meeting[0]):", LineStart: 5, LineEnd: 5},
				{Title: "Test for overlap", Explanation: "If the earliest-free room is free by this start, reuse it. Example 1: [0,30] opens room A (heap [30]); [5,10]: 30 > 5, so room B (heap [10,30]); [15,20]: 10 <= 15, so pop 10 and reuse room B.", CodeSnippet: "if end_times and end_times[0] <= start:\n    heapq.heappop(end_times)  # earliest room is free: reuse it", LineStart: 6, LineEnd: 7},
				{Title: "Merge or keep", Explanation: "Either way, the room this meeting uses is now busy until its end.", CodeSnippet: "heapq.heappush(end_times, end)", LineStart: 8, LineEnd: 8},
				{Title: "Return the answer", Explanation: "Each heap entry is a room that was opened, so the heap size is the room count. Example 1: [20,30], 2 rooms.", CodeSnippet: "return len(end_times)", LineStart: 9, LineEnd: 9},
			},
			Alternatives: []AltSolution{
				{
					Name: "Sweep line with two sorted arrays",
					Code: `def minMeetingRooms(intervals):
    starts = sorted(meeting[0] for meeting in intervals)
    ends = sorted(meeting[1] for meeting in intervals)
    rooms = 0
    end_ptr = 0
    for start in starts:
        if start < ends[end_ptr]:
            rooms += 1     # nothing has ended yet: open a room
        else:
            end_ptr += 1   # a meeting ended by now: reuse its room
    return rooms`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Treat starts and ends as separate events on a timeline. Walk the starts in order; ends[end_ptr] is the earliest end not yet used. If the start is before it, all rooms are busy, so open one; otherwise a room is free, so reuse it and move to the next end. start < ends[end_ptr] is strict for the same reason as the heap's <=: a meeting ending at 5 frees its room for one starting at 5. You don't need to know which meeting ended, only that one did.",
				},
				{
					Name: "Count overlaps at each start (brute force)",
					Code: `def minMeetingRooms(intervals):
    best = 0
    for meeting in intervals:
        moment = meeting[0]
        running = sum(1 for s, e in intervals if s <= moment < e)
        best = max(best, running)
    return best`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "The maximum overlap always happens at some meeting's start, so count the meetings running at each start time. s <= moment < e is the half-open membership test.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it enough to look at only the earliest-ending room when a new meeting starts?", Answer: "If the room that frees up first is still busy at the new start, every other room is busy too, since they all free up later. If it is free, using it is as good as any other free room: all free rooms are equally available for this meeting, and the meetings still to come only care about when rooms become free."},
			{Prompt: "In the two-array sweep, ends are sorted separately from starts, so a start is no longer paired with its own end. Why is that still correct?", Answer: "The question is only how many meetings are running at each moment, which is (starts so far) minus (ends so far). Which meeting ended doesn't matter for counting, just that some room became free. So sorting starts and ends independently loses nothing."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, IntervalsProblems...)
}
