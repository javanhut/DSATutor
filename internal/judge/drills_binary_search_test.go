package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestBinarySearchDrillsRejectMistakes checks that each binary search
// foundations drill fails the mistake its tests and notes promise to catch, so
// a drill can't be passed without doing the thing it drills. A linear scan is
// not rejected for the array drills (their inputs are too small to time out),
// so only the answer-range scan in the shipping drill is in this list.
func TestBinarySearchDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-search-insert-position", "reads past the end with while lo <= hi on a half-open range", `def searchInsert(nums, target):
    lo, hi = 0, len(nums)
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] < target:
            lo = mid + 1
        else:
            hi = mid
    return lo`},
		{"drill-search-insert-position", "uses hi = mid - 1 on a half-open range", `def searchInsert(nums, target):
    lo, hi = 0, len(nums)
    while lo < hi:
        mid = (lo + hi) // 2
        if nums[mid] < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return lo`},
		{"drill-search-insert-position", "uses lo = mid and never ends", `def searchInsert(nums, target):
    lo, hi = 0, len(nums)
    while lo < hi:
        mid = (lo + hi) // 2
        if nums[mid] < target:
            lo = mid
        else:
            hi = mid
    return lo`},
		{"drill-search-insert-position", "returns -1 when the target is missing", `def searchInsert(nums, target):
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] == target:
            return mid
        if nums[mid] < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return -1`},
		{"drill-count-occurrences", "uses < in both bounds, so the count is always 0", `def countOccurrences(nums, target):
    def bound():
        lo, hi = 0, len(nums)
        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] < target:
                lo = mid + 1
            else:
                hi = mid
        return lo

    return bound() - bound()`},
		{"drill-count-occurrences", "finds one copy and reports 1", `def countOccurrences(nums, target):
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] == target:
            return 1
        if nums[mid] < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return 0`},
		{"drill-count-occurrences", "starts the upper bound with hi = len(nums) - 1", `def countOccurrences(nums, target):
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
        lo, hi = 0, len(nums) - 1
        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] <= target:
                lo = mid + 1
            else:
                hi = mid
        return lo

    return first_greater() - first_at_least()`},
		{"drill-first-and-last-position", "reads nums[first] before checking the bound", `def searchRange(nums, target):
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
    if nums[first] != target or first == len(nums):
        return [-1, -1]
    return [first, first_at_least(target + 1) - 1]`},
		{"drill-first-and-last-position", "skips the existence check", `def searchRange(nums, target):
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
    return [first, first_at_least(target + 1) - 1]`},
		{"drill-first-and-last-position", "returns the same index found for both ends", `def searchRange(nums, target):
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] == target:
            return [mid, mid]
        if nums[mid] < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return [-1, -1]`},
		{"drill-first-and-last-position", "forgets the - 1 on the last position", `def searchRange(nums, target):
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
    return [first, first_at_least(target + 1)]`},
		{"drill-first-true", "returns mid as soon as it sees a True", `def firstTrue(flags):
    lo, hi = 0, len(flags)
    while lo < hi:
        mid = (lo + hi) // 2
        if flags[mid]:
            return mid
        lo = mid + 1
    return lo`},
		{"drill-first-true", "uses hi = mid - 1 after a True", `def firstTrue(flags):
    lo, hi = 0, len(flags)
    while lo < hi:
        mid = (lo + hi) // 2
        if flags[mid]:
            hi = mid - 1
        else:
            lo = mid + 1
    return lo`},
		{"drill-first-true", "starts with hi = len(flags) - 1, so all-False answers len - 1", `def firstTrue(flags):
    lo, hi = 0, len(flags) - 1
    while lo < hi:
        mid = (lo + hi) // 2
        if flags[mid]:
            hi = mid
        else:
            lo = mid + 1
    return lo`},
		{"drill-first-true", "uses lo = mid after a False and never ends", `def firstTrue(flags):
    lo, hi = 0, len(flags)
    while lo < hi:
        mid = (lo + hi) // 2
        if flags[mid]:
            hi = mid
        else:
            lo = mid
    return lo`},
		{"drill-integer-square-root", "returns lo, one too big", `def mySqrt(x):
    lo, hi = 0, x + 1
    while lo < hi:
        mid = (lo + hi) // 2
        if mid * mid > x:
            hi = mid
        else:
            lo = mid + 1
    return lo`},
		{"drill-integer-square-root", "starts with hi = x", `def mySqrt(x):
    lo, hi = 0, x
    while lo < hi:
        mid = (lo + hi) // 2
        if mid * mid > x:
            hi = mid
        else:
            lo = mid + 1
    return lo - 1`},
		{"drill-integer-square-root", "only matches exact squares", `def mySqrt(x):
    lo, hi = 0, x
    while lo <= hi:
        mid = (lo + hi) // 2
        if mid * mid == x:
            return mid
        if mid * mid < x:
            lo = mid + 1
        else:
            hi = mid - 1
    return lo`},
		{"drill-integer-square-root", "mixes while lo <= hi with hi = mid and never ends", `def mySqrt(x):
    lo, hi = 0, x + 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if mid * mid > x:
            hi = mid
        else:
            lo = mid + 1
    return lo - 1`},
		{"drill-find-peak-element", "walks toward a valley (wrong direction)", `def findPeakElement(nums):
    lo, hi = 0, len(nums) - 1
    while lo < hi:
        mid = (lo + hi) // 2
        if nums[mid] > nums[mid + 1]:
            lo = mid + 1
        else:
            hi = mid
    return lo`},
		{"drill-find-peak-element", "only checks the two ends", `def findPeakElement(nums):
    if len(nums) == 1 or nums[0] > nums[1]:
        return 0
    return len(nums) - 1`},
		{"drill-find-peak-element", "reads nums[mid + 1] when mid can be the last index", `def findPeakElement(nums):
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if nums[mid] < nums[mid + 1]:
            lo = mid + 1
        else:
            hi = mid - 1
    return lo`},
		{"drill-find-peak-element", "uses hi = mid - 1 in the else branch", `def findPeakElement(nums):
    lo, hi = 0, len(nums) - 1
    while lo < hi:
        mid = (lo + hi) // 2
        if nums[mid] < nums[mid + 1]:
            lo = mid + 1
        else:
            hi = mid - 1
    return lo`},
		{"drill-capacity-to-ship-packages", "starts the range at 1 instead of max(weights)", `def shipWithinDays(weights, days):
    def fits(capacity):
        used, load = 1, 0
        for weight in weights:
            if load + weight > capacity:
                used += 1
                load = 0
            load += weight
        return used <= days

    lo, hi = 1, sum(weights)
    while lo < hi:
        mid = (lo + hi) // 2
        if fits(mid):
            hi = mid
        else:
            lo = mid + 1
    return lo`},
		{"drill-capacity-to-ship-packages", "uses >= in the overflow test", `def shipWithinDays(weights, days):
    def fits(capacity):
        used, load = 1, 0
        for weight in weights:
            if load + weight >= capacity:
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
    return lo`},
		{"drill-capacity-to-ship-packages", "starts counting days at 0", `def shipWithinDays(weights, days):
    def fits(capacity):
        used, load = 0, 0
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
    return lo`},
		{"drill-capacity-to-ship-packages", "uses hi = mid - 1 after a success", `def shipWithinDays(weights, days):
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
            hi = mid - 1
        else:
            lo = mid + 1
    return lo`},
		{"drill-capacity-to-ship-packages", "scans every capacity one at a time (too slow)", `def shipWithinDays(weights, days):
    def fits(capacity):
        used, load = 1, 0
        for weight in weights:
            if load + weight > capacity:
                used += 1
                load = 0
            load += weight
        return used <= days

    capacity = max(weights)
    while not fits(capacity):
        capacity += 1
    return capacity`},
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
