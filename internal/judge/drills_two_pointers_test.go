package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestTwoPointersDrillsRejectMistakes checks that each two pointers foundations
// drill fails the mistake its tests and notes promise to catch, so a drill
// can't be passed without doing the thing it drills.
func TestTwoPointersDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-reverse-in-place", "swaps every pair twice (loops over every index)", `def reverseString(s):
    n = len(s)
    for i in range(n):
        s[i], s[n - 1 - i] = s[n - 1 - i], s[i]`},
		{"drill-reverse-in-place", "rebinds s to a reversed copy", `def reverseString(s):
    s = s[::-1]`},
		{"drill-reverse-in-place", "stops one step early (left < right - 1)", `def reverseString(s):
    left, right = 0, len(s) - 1
    while left < right - 1:
        s[left], s[right] = s[right], s[left]
        left += 1
        right -= 1`},
		{"drill-remove-duplicates-sorted", "counts distinct values without moving them forward", `def removeDuplicates(nums):
    return len(set(nums))`},
		{"drill-remove-duplicates-sorted", "off by one: returns write - 1", `def removeDuplicates(nums):
    if not nums:
        return 0
    write = 1
    for read in range(1, len(nums)):
        if nums[read] != nums[write - 1]:
            nums[write] = nums[read]
            write += 1
    return write - 1`},
		{"drill-remove-duplicates-sorted", "reads nums[0] on an empty array", `def removeDuplicates(nums):
    last = nums[0]
    write = 1
    for read in range(1, len(nums)):
        if nums[read] != last:
            nums[write] = nums[read]
            last = nums[read]
            write += 1
    return write`},
		{"drill-move-zeroes", "treats negatives as zeros (> 0)", `def moveZeroes(nums):
    write = 0
    for read in range(len(nums)):
        if nums[read] > 0:
            nums[write], nums[read] = nums[read], nums[write]
            write += 1`},
		{"drill-move-zeroes", "swaps from both ends (not stable)", `def moveZeroes(nums):
    left, right = 0, len(nums) - 1
    while left < right:
        if nums[left] != 0:
            left += 1
        elif nums[right] == 0:
            right -= 1
        else:
            nums[left], nums[right] = nums[right], nums[left]`},
		{"drill-move-zeroes", "copies nonzero items forward but never fills zeros", `def moveZeroes(nums):
    write = 0
    for read in range(len(nums)):
        if nums[read] != 0:
            nums[write] = nums[read]
            write += 1`},
		{"drill-move-zeroes", "returns a new list instead of changing nums", `def moveZeroes(nums):
    nums = [x for x in nums if x != 0] + [0] * nums.count(0)
    return nums`},
		{"drill-move-zeroes", "removes zeros one at a time and appends", `def moveZeroes(nums):
    zeros = nums.count(0)
    for _ in range(zeros):
        nums.remove(0)
        nums.append(0)`},
		{"drill-merge-two-sorted-arrays", "drops the leftover tail", `def mergeSorted(a, b):
    merged = []
    i = j = 0
    while i < len(a) and j < len(b):
        if a[i] <= b[j]:
            merged.append(a[i])
            i += 1
        else:
            merged.append(b[j])
            j += 1
    return merged`},
		{"drill-merge-two-sorted-arrays", "drops an item when the two fronts tie", `def mergeSorted(a, b):
    merged = []
    i = j = 0
    while i < len(a) and j < len(b):
        if a[i] < b[j]:
            merged.append(a[i])
            i += 1
        elif a[i] > b[j]:
            merged.append(b[j])
            j += 1
        else:
            merged.append(a[i])
            i += 1
            j += 1
    merged.extend(a[i:])
    merged.extend(b[j:])
    return merged`},
		{"drill-merge-two-sorted-arrays", "indexes past the end of the shorter array", `def mergeSorted(a, b):
    merged = []
    i = j = 0
    while i < len(a) or j < len(b):
        if a[i] <= b[j]:
            merged.append(a[i])
            i += 1
        else:
            merged.append(b[j])
            j += 1
    return merged`},
		{"drill-squares-sorted-array", "compares raw values instead of absolute values", `def sortedSquares(nums):
    result = [0] * len(nums)
    left, right = 0, len(nums) - 1
    for pos in range(len(nums) - 1, -1, -1):
        if nums[left] > nums[right]:
            result[pos] = nums[left] * nums[left]
            left += 1
        else:
            result[pos] = nums[right] * nums[right]
            right -= 1
    return result`},
		{"drill-squares-sorted-array", "fills the result from the front (descending)", `def sortedSquares(nums):
    result = []
    left, right = 0, len(nums) - 1
    while left <= right:
        if abs(nums[left]) > abs(nums[right]):
            result.append(nums[left] * nums[left])
            left += 1
        else:
            result.append(nums[right] * nums[right])
            right -= 1
    return result`},
		{"drill-squares-sorted-array", "stops when the pointers meet (left < right)", `def sortedSquares(nums):
    result = [0] * len(nums)
    left, right = 0, len(nums) - 1
    pos = len(nums) - 1
    while left < right:
        if abs(nums[left]) > abs(nums[right]):
            result[pos] = nums[left] * nums[left]
            left += 1
        else:
            result[pos] = nums[right] * nums[right]
            right -= 1
        pos -= 1
    return result`},
		{"drill-squares-sorted-array", "squares without reordering", `def sortedSquares(nums):
    return [x * x for x in nums]`},
		{"drill-is-subsequence", "checks only that every letter appears (ignores order)", `def isSubsequence(s, t):
    return all(ch in t for ch in s)`},
		{"drill-is-subsequence", "moves t only on a mismatch (reuses a letter of t)", `def isSubsequence(s, t):
    i = j = 0
    while i < len(s) and j < len(t):
        if s[i] == t[j]:
            i += 1
        else:
            j += 1
    return i == len(s)`},
		{"drill-is-subsequence", "requires a substring", `def isSubsequence(s, t):
    return s in t`},
		{"drill-is-subsequence", "reads s[0] on an empty s", `def isSubsequence(s, t):
    i = 0
    for ch in t:
        if ch == s[i]:
            i += 1
            if i == len(s):
                return True
    return False`},
		{"drill-intersect-sorted-arrays", "advances only one pointer after a match", `def intersect(nums1, nums2):
    common = []
    i = j = 0
    while i < len(nums1) and j < len(nums2):
        if nums1[i] < nums2[j]:
            i += 1
        elif nums1[i] > nums2[j]:
            j += 1
        else:
            common.append(nums1[i])
            i += 1
    return common`},
		{"drill-intersect-sorted-arrays", "advances both pointers on a mismatch", `def intersect(nums1, nums2):
    common = []
    i = j = 0
    while i < len(nums1) and j < len(nums2):
        if nums1[i] == nums2[j]:
            common.append(nums1[i])
        i += 1
        j += 1
    return common`},
		{"drill-intersect-sorted-arrays", "uses a set intersection (loses repeats)", `def intersect(nums1, nums2):
    return sorted(set(nums1) & set(nums2))`},
		{"drill-count-pairs-sorted", "moves both pointers by one on a hit", `def countPairs(nums, target):
    count = 0
    left, right = 0, len(nums) - 1
    while left < right:
        total = nums[left] + nums[right]
        if total < target:
            left += 1
        elif total > target:
            right -= 1
        else:
            count += 1
            left += 1
            right -= 1
    return count`},
		{"drill-count-pairs-sorted", "multiplies run lengths even when both pointers share one run", `def countPairs(nums, target):
    count = 0
    left, right = 0, len(nums) - 1
    while left < right:
        total = nums[left] + nums[right]
        if total < target:
            left += 1
        elif total > target:
            right -= 1
        else:
            low, high = nums[left], nums[right]
            low_count = high_count = 0
            while left < len(nums) and nums[left] == low:
                left += 1
                low_count += 1
            while right >= 0 and nums[right] == high:
                right -= 1
                high_count += 1
            count += low_count * high_count
    return count`},
		{"drill-count-pairs-sorted", "counts distinct value pairs, not positions", `def countPairs(nums, target):
    seen = set(nums)
    return len([x for x in seen if x < target - x and target - x in seen]) + (1 if target % 2 == 0 and nums.count(target // 2) >= 2 else 0)`},
		{"drill-count-pairs-sorted", "checks every pair of positions", `def countPairs(nums, target):
    count = 0
    for i in range(len(nums)):
        for j in range(i + 1, len(nums)):
            if nums[i] + nums[j] == target:
                count += 1
    return count`},
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
