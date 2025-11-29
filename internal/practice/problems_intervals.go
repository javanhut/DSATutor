package practice

// IntervalsProblems contains all intervals category problems
var IntervalsProblems = []*Problem{
	{
		ID:              "merge-intervals",
		Number:          90,
		Title:           "Merge Intervals",
		Difficulty:      "Medium",
		Category:        "intervals",
		Tags:            []string{"Array", "Sorting"},
		RelatedChapters: []int{8},
		Description: `Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.`,
		Constraints: []string{
			"1 <= intervals.length <= 10^4",
			"intervals[i].length == 2",
			"0 <= starti <= endi <= 10^4",
		},
		Examples: []Example{
			{Input: "intervals = [[1,3],[2,6],[8,10],[15,18]]", Output: "[[1,6],[8,10],[15,18]]"},
			{Input: "intervals = [[1,4],[4,5]]", Output: "[[1,5]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, Expected: [][]int{{1, 6}, {8, 10}, {15, 18}}},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 4}, {4, 5}}}, Expected: [][]int{{1, 5}}},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def merge(intervals):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Sort by start time. Merge consecutive overlapping intervals by extending end time.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "Sort intervals. For each interval: if overlaps with last in result, extend end. Else add as new interval."},
			{Level: 3, Type: "code", Content: "Sort by start. result = [intervals[0]]. for interval: if interval[0] <= result[-1][1]: merge. else append."},
		},
		Solution: Solution{
			Code: `def merge(intervals):
    intervals.sort(key=lambda x: x[0])
    result = [intervals[0]]

    for start, end in intervals[1:]:
        if start <= result[-1][1]:
            result[-1][1] = max(result[-1][1], end)
        else:
            result.append([start, end])

    return result`,
			Explanation:     "Sort by start. If current overlaps with last merged, extend end. Otherwise, add as new interval.",
			TimeComplexity:  "O(n log n) for sorting",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort by start", Explanation: "Process intervals in order", CodeSnippet: "intervals.sort(key=lambda x: x[0])", LineStart: 2, LineEnd: 2},
				{Title: "Check overlap", Explanation: "Overlaps if start <= previous end", CodeSnippet: "if start <= result[-1][1]:", LineStart: 6, LineEnd: 6},
				{Title: "Merge or add", Explanation: "Extend end or add new interval", CodeSnippet: "result[-1][1] = max(result[-1][1], end)", LineStart: 7, LineEnd: 7},
			},
		},
	},
	{
		ID:              "non-overlapping-intervals",
		Number:          91,
		Title:           "Non-overlapping Intervals",
		Difficulty:      "Medium",
		Category:        "intervals",
		Tags:            []string{"Array", "Dynamic Programming", "Greedy", "Sorting"},
		RelatedChapters: []int{8},
		Description: `Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.`,
		Constraints: []string{
			"1 <= intervals.length <= 10^5",
			"intervals[i].length == 2",
			"-5 * 10^4 <= starti < endi <= 5 * 10^4",
		},
		Examples: []Example{
			{Input: "intervals = [[1,2],[2,3],[3,4],[1,3]]", Output: "1", Explanation: "Remove [1,3]."},
			{Input: "intervals = [[1,2],[1,2],[1,2]]", Output: "2"},
			{Input: "intervals = [[1,2],[2,3]]", Output: "0"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}}, Expected: 1},
			{Input: map[string]interface{}{"intervals": [][]int{{1, 2}, {1, 2}, {1, 2}}}, Expected: 2},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def eraseOverlapIntervals(intervals):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Greedy: sort by end time. Keep intervals that end earliest. Count removals when overlap detected.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "Sort by end. Track prev_end. If start < prev_end: overlap, remove one (increment count). Else update prev_end."},
			{Level: 3, Type: "code", Content: "Sort by end. prev_end = intervals[0][1], count = 0. for interval: if start < prev_end: count++. else: prev_end = end."},
		},
		Solution: Solution{
			Code: `def eraseOverlapIntervals(intervals):
    intervals.sort(key=lambda x: x[1])  # Sort by end time
    count = 0
    prev_end = float('-inf')

    for start, end in intervals:
        if start >= prev_end:
            prev_end = end
        else:
            count += 1  # Remove this interval

    return count`,
			Explanation:     "Greedy: sort by end time. Keep intervals ending earliest. Count overlaps (which need removal).",
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort by end", Explanation: "Greedy choice: keep early-ending intervals", CodeSnippet: "intervals.sort(key=lambda x: x[1])", LineStart: 2, LineEnd: 2},
				{Title: "No overlap", Explanation: "Can keep this interval", CodeSnippet: "if start >= prev_end:\n    prev_end = end", LineStart: 7, LineEnd: 8},
				{Title: "Overlap", Explanation: "Must remove this interval", CodeSnippet: "count += 1", LineStart: 10, LineEnd: 10},
			},
		},
	},
	{
		ID:              "insert-interval",
		Number:          109,
		Title:           "Insert Interval",
		Difficulty:      "Medium",
		Category:        "intervals",
		Tags:            []string{"Array"},
		RelatedChapters: []int{8},
		Description: `You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.`,
		Constraints: []string{
			"0 <= intervals.length <= 10^4",
			"intervals[i].length == 2",
			"0 <= starti <= endi <= 10^5",
			"intervals is sorted by starti in ascending order",
			"newInterval.length == 2",
			"0 <= start <= end <= 10^5",
		},
		Examples: []Example{
			{Input: "intervals = [[1,3],[6,9]], newInterval = [2,5]", Output: "[[1,5],[6,9]]"},
			{Input: "intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]", Output: "[[1,2],[3,10],[12,16]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{1, 3}, {6, 9}}, "newInterval": []int{2, 5}}, Expected: [][]int{{1, 5}, {6, 9}}},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def insert(intervals, newInterval):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Three phases: add intervals before newInterval, merge overlapping intervals, add intervals after.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "Add all intervals ending before newInterval starts. Merge overlapping. Add remaining intervals."},
			{Level: 3, Type: "code", Content: "while end < new_start: add. while start <= new_end: merge. add merged. add remaining."},
		},
		Solution: Solution{
			Code: `def insert(intervals, newInterval):
    result = []
    i = 0
    n = len(intervals)

    # Add all intervals before newInterval
    while i < n and intervals[i][1] < newInterval[0]:
        result.append(intervals[i])
        i += 1

    # Merge overlapping intervals
    while i < n and intervals[i][0] <= newInterval[1]:
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i += 1

    result.append(newInterval)

    # Add remaining intervals
    while i < n:
        result.append(intervals[i])
        i += 1

    return result`,
			Explanation:     "Three phases: add non-overlapping before, merge overlapping, add non-overlapping after.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Before new interval", Explanation: "Intervals ending before new starts", CodeSnippet: "while i < n and intervals[i][1] < newInterval[0]:", LineStart: 7, LineEnd: 9},
				{Title: "Merge overlapping", Explanation: "Expand newInterval to cover all overlapping", CodeSnippet: "newInterval[0] = min(newInterval[0], intervals[i][0])\nnewInterval[1] = max(newInterval[1], intervals[i][1])", LineStart: 13, LineEnd: 14},
				{Title: "Add remaining", Explanation: "Intervals starting after new ends", CodeSnippet: "while i < n:", LineStart: 20, LineEnd: 22},
			},
		},
	},
	{
		ID:              "meeting-rooms-ii",
		Number:          110,
		Title:           "Meeting Rooms II",
		Difficulty:      "Medium",
		Category:        "intervals",
		Tags:            []string{"Array", "Two Pointers", "Greedy", "Sorting", "Heap"},
		RelatedChapters: []int{7, 8},
		Description: `Given an array of meeting time intervals intervals where intervals[i] = [starti, endi], return the minimum number of conference rooms required.`,
		Constraints: []string{
			"1 <= intervals.length <= 10^4",
			"0 <= starti < endi <= 10^6",
		},
		Examples: []Example{
			{Input: "intervals = [[0,30],[5,10],[15,20]]", Output: "2"},
			{Input: "intervals = [[7,10],[2,4]]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{0, 30}, {5, 10}, {15, 20}}}, Expected: 2},
			{Input: map[string]interface{}{"intervals": [][]int{{7, 10}, {2, 4}}}, Expected: 1},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def minMeetingRooms(intervals):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Sort starts and ends separately. Count concurrent meetings using two pointers.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "When a meeting starts, need a room. When one ends, free a room. Track max concurrent."},
			{Level: 3, Type: "code", Content: "Sort starts, ends. Two pointers. If start < end: rooms++, move start. else: rooms--, move end. Track max."},
		},
		Solution: Solution{
			Code: `def minMeetingRooms(intervals):
    starts = sorted(i[0] for i in intervals)
    ends = sorted(i[1] for i in intervals)

    rooms = 0
    max_rooms = 0
    s = e = 0

    while s < len(intervals):
        if starts[s] < ends[e]:
            rooms += 1
            s += 1
        else:
            rooms -= 1
            e += 1
        max_rooms = max(max_rooms, rooms)

    return max_rooms`,
			Explanation:     "Sort starts and ends separately. Track concurrent meetings using two pointers.",
			TimeComplexity:  "O(n log n) for sorting",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort separately", Explanation: "We only care about when meetings start/end, not which meeting", CodeSnippet: "starts = sorted(i[0] for i in intervals)\nends = sorted(i[1] for i in intervals)", LineStart: 2, LineEnd: 3},
				{Title: "Track concurrent", Explanation: "Start before end means new room needed", CodeSnippet: "if starts[s] < ends[e]:\n    rooms += 1", LineStart: 10, LineEnd: 12},
			},
		},
	},
	{
		ID:              "meeting-rooms",
		Number:          138,
		Title:           "Meeting Rooms",
		Difficulty:      "Easy",
		Category:        "intervals",
		Tags:            []string{"Array", "Sorting"},
		RelatedChapters: []int{2, 8},
		Description: `Given an array of meeting time intervals where intervals[i] = [starti, endi], determine if a person could attend all meetings.`,
		Constraints: []string{
			"0 <= intervals.length <= 10^4",
			"intervals[i].length == 2",
			"0 <= starti < endi <= 10^6",
		},
		Examples: []Example{
			{Input: "intervals = [[0,30],[5,10],[15,20]]", Output: "false"},
			{Input: "intervals = [[7,10],[2,4]]", Output: "true"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"intervals": [][]int{{0, 30}, {5, 10}, {15, 20}}}, Expected: false},
			{Input: map[string]interface{}{"intervals": [][]int{{7, 10}, {2, 4}}}, Expected: true},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def canAttendMeetings(intervals):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Sort by start time, then check for overlaps.", ChapterRef: 8},
			{Level: 2, Type: "algorithm", Content: "After sorting, if any meeting starts before previous ends, there's overlap."},
			{Level: 3, Type: "code", Content: "intervals.sort(). for i in range(1, n): if intervals[i][0] < intervals[i-1][1]: return False."},
		},
		Solution: Solution{
			Code: `def canAttendMeetings(intervals):
    intervals.sort()

    for i in range(1, len(intervals)):
        if intervals[i][0] < intervals[i - 1][1]:
            return False

    return True`,
			Explanation:     "Sort by start time. Check if any meeting starts before previous ends.",
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort first", Explanation: "Order by start time", CodeSnippet: "intervals.sort()", LineStart: 2, LineEnd: 2},
				{Title: "Check overlap", Explanation: "Start before previous end means conflict", CodeSnippet: "if intervals[i][0] < intervals[i - 1][1]:", LineStart: 5, LineEnd: 5},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, IntervalsProblems...)
}
