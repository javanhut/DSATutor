package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestIntervalsDrillsRejectMistakes checks that each intervals foundations
// drill fails the mistake its tests and notes promise to catch, so a drill
// can't be passed without doing the thing it drills.
func TestIntervalsDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-intervals-overlap", "checks only one ordering (b starts inside a)", `def overlaps(a, b):
    return a[0] <= b[0] <= a[1]`},
		{"drill-intervals-overlap", "uses < so touching intervals do not overlap", `def overlaps(a, b):
    return max(a[0], b[0]) < min(a[1], b[1])`},
		{"drill-intervals-overlap", "checks only that a ends after b starts", `def overlaps(a, b):
    return a[1] >= b[0]`},
		{"drill-interval-list-intersection", "advances both pointers every step", `def intervalIntersection(firstList, secondList):
    result = []
    i = j = 0
    while i < len(firstList) and j < len(secondList):
        start = max(firstList[i][0], secondList[j][0])
        end = min(firstList[i][1], secondList[j][1])
        if start <= end:
            result.append([start, end])
        i += 1
        j += 1
    return result`},
		{"drill-interval-list-intersection", "uses < and drops one-point intersections", `def intervalIntersection(firstList, secondList):
    result = []
    i = j = 0
    while i < len(firstList) and j < len(secondList):
        start = max(firstList[i][0], secondList[j][0])
        end = min(firstList[i][1], secondList[j][1])
        if start < end:
            result.append([start, end])
        if firstList[i][1] < secondList[j][1]:
            i += 1
        else:
            j += 1
    return result`},
		{"drill-interval-list-intersection", "advances the interval that starts first", `def intervalIntersection(firstList, secondList):
    result = []
    i = j = 0
    while i < len(firstList) and j < len(secondList):
        start = max(firstList[i][0], secondList[j][0])
        end = min(firstList[i][1], secondList[j][1])
        if start <= end:
            result.append([start, end])
        if firstList[i][0] < secondList[j][0]:
            i += 1
        else:
            j += 1
    return result`},
		{"drill-interval-list-intersection", "appends without checking start <= end", `def intervalIntersection(firstList, secondList):
    result = []
    i = j = 0
    while i < len(firstList) and j < len(secondList):
        start = max(firstList[i][0], secondList[j][0])
        end = min(firstList[i][1], secondList[j][1])
        result.append([start, end])
        if firstList[i][1] < secondList[j][1]:
            i += 1
        else:
            j += 1
    return result`},
		{"drill-total-covered-length", "does not sort", `def coveredLength(intervals):
    if not intervals:
        return 0
    total = 0
    cur_start, cur_end = intervals[0]
    for start, end in intervals[1:]:
        if start <= cur_end:
            cur_end = max(cur_end, end)
        else:
            total += cur_end - cur_start
            cur_start, cur_end = start, end
    return total + cur_end - cur_start`},
		{"drill-total-covered-length", "overwrites the end with the new end instead of max", `def coveredLength(intervals):
    if not intervals:
        return 0
    intervals = sorted(intervals)
    total = 0
    cur_start, cur_end = intervals[0]
    for start, end in intervals[1:]:
        if start <= cur_end:
            cur_end = end
        else:
            total += cur_end - cur_start
            cur_start, cur_end = start, end
    return total + cur_end - cur_start`},
		{"drill-total-covered-length", "adds up every length (overlaps counted twice)", `def coveredLength(intervals):
    return sum(end - start for start, end in intervals)`},
		{"drill-total-covered-length", "compares with the previous interval, not the merged block", `def coveredLength(intervals):
    if not intervals:
        return 0
    intervals = sorted(intervals)
    total = 0
    cur_start, cur_end = intervals[0]
    prev_end = cur_end
    for start, end in intervals[1:]:
        if start <= prev_end:
            cur_end = max(cur_end, end)
        else:
            total += cur_end - cur_start
            cur_start, cur_end = start, end
        prev_end = end
    return total + cur_end - cur_start`},
		{"drill-total-covered-length", "forgets to add the last block", `def coveredLength(intervals):
    if not intervals:
        return 0
    intervals = sorted(intervals)
    total = 0
    cur_start, cur_end = intervals[0]
    for start, end in intervals[1:]:
        if start <= cur_end:
            cur_end = max(cur_end, end)
        else:
            total += cur_end - cur_start
            cur_start, cur_end = start, end
    return total`},
		{"drill-remove-covered-intervals", "sorts by start with ends ascending", `def removeCoveredIntervals(intervals):
    intervals = sorted(intervals)
    remaining = 0
    furthest_end = 0
    for start, end in intervals:
        if end > furthest_end:
            remaining += 1
            furthest_end = end
    return remaining`},
		{"drill-remove-covered-intervals", "uses >= so a shared end looks uncovered", `def removeCoveredIntervals(intervals):
    intervals = sorted(intervals, key=lambda iv: (iv[0], -iv[1]))
    remaining = 0
    furthest_end = 0
    for start, end in intervals:
        if end >= furthest_end:
            remaining += 1
            furthest_end = end
    return remaining`},
		{"drill-remove-covered-intervals", "compares with the previous interval's end, not the largest", `def removeCoveredIntervals(intervals):
    intervals = sorted(intervals, key=lambda iv: (iv[0], -iv[1]))
    remaining = 0
    prev_end = 0
    for start, end in intervals:
        if end > prev_end:
            remaining += 1
        prev_end = end
    return remaining`},
		{"drill-remove-covered-intervals", "does not sort", `def removeCoveredIntervals(intervals):
    remaining = 0
    furthest_end = 0
    for start, end in intervals:
        if end > furthest_end:
            remaining += 1
            furthest_end = end
    return remaining`},
		{"drill-max-overlap-closed", "processes ends before starts at the same time", `def maxOverlap(intervals):
    events = []
    for start, end in intervals:
        events.append((start, 1))
        events.append((end, 0))
    events.sort()
    current = best = 0
    for time, kind in events:
        if kind == 1:
            current += 1
            best = max(best, current)
        else:
            current -= 1
    return best`},
		{"drill-max-overlap-closed", "never subtracts at an end", `def maxOverlap(intervals):
    return len(intervals)`},
		{"drill-max-overlap-closed", "counts overlapping pairs instead of the peak", `def maxOverlap(intervals):
    pairs = 0
    for i in range(len(intervals)):
        for j in range(i + 1, len(intervals)):
            if max(intervals[i][0], intervals[j][0]) <= min(intervals[i][1], intervals[j][1]):
                pairs += 1
    return pairs`},
		{"drill-min-arrows-balloons", "sorts by start and shoots at the first end", `def findMinArrowShots(points):
    arrows = 0
    arrow_x = None
    for start, end in sorted(points):
        if arrow_x is None or start > arrow_x:
            arrows += 1
            arrow_x = end
    return arrows`},
		{"drill-min-arrows-balloons", "uses >= so a balloon touching the arrow needs another", `def findMinArrowShots(points):
    arrows = 0
    arrow_x = None
    for start, end in sorted(points, key=lambda p: p[1]):
        if arrow_x is None or start >= arrow_x:
            arrows += 1
            arrow_x = end
    return arrows`},
		{"drill-min-arrows-balloons", "counts merged groups of overlapping balloons", `def findMinArrowShots(points):
    if not points:
        return 0
    points = sorted(points)
    groups = 1
    reach = points[0][1]
    for start, end in points[1:]:
        if start > reach:
            groups += 1
            reach = end
        else:
            reach = max(reach, end)
    return groups`},
	}
	for _, c := range cases {
		c := c
		t.Run(c.id+"/"+c.why, func(t *testing.T) {
			t.Parallel()
			var p *practice.Problem
			for _, q := range practice.EmbeddedProblems {
				if q.ID == c.id {
					p = q
				}
			}
			if p == nil {
				t.Fatalf("no problem %s", c.id)
			}
			for i := range p.TestCases {
				res, err := Run(p, c.code, &p.TestCases[i], Options{})
				if err != nil {
					t.Fatalf("case %d: %v", i+1, err)
				}
				if !res.Passed {
					return
				}
			}
			t.Errorf("code that %s passed every case", c.why)
		})
	}
}
