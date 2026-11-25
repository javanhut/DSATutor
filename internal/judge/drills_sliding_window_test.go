package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestSlidingWindowDrillsRejectMistakes checks that each sliding window
// foundations drill fails the mistake its tests and notes promise to catch,
// so a drill can't be passed without doing the thing it drills. The
// recompute-from-scratch cases are caught by the large time-limited case.
func TestSlidingWindowDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-window-sums", "recomputes each window's sum from scratch", `def windowSums(nums, k):
    return [sum(nums[i:i + k]) for i in range(len(nums) - k + 1)]
`},
		{"drill-window-sums", "subtracts nums[right - k + 1], which is still in the window", `def windowSums(nums, k):
    window = sum(nums[:k])
    sums = [window]
    for right in range(k, len(nums)):
        window += nums[right] - nums[right - k + 1]
        sums.append(window)
    return sums
`},
		{"drill-window-sums", "forgets to subtract the outgoing element", `def windowSums(nums, k):
    window = sum(nums[:k])
    sums = [window]
    for right in range(k, len(nums)):
        window += nums[right]
        sums.append(window)
    return sums
`},
		{"drill-window-sums", "leaves out the last window", `def windowSums(nums, k):
    window = sum(nums[:k])
    sums = [window]
    for right in range(k, len(nums) - 1):
        window += nums[right] - nums[right - k]
        sums.append(window)
    return sums
`},
		{"drill-window-sums", "leaves out the first window", `def windowSums(nums, k):
    window = sum(nums[:k])
    sums = []
    for right in range(k, len(nums)):
        window += nums[right] - nums[right - k]
        sums.append(window)
    return sums
`},
		{"drill-window-sums", "starts sliding at index 0 and adds the first elements twice", `def windowSums(nums, k):
    window = sum(nums[:k])
    sums = [window]
    for right in range(0, len(nums)):
        window += nums[right] - nums[right - k]
        sums.append(window)
    return sums[:len(nums) - k + 1]
`},
		{"drill-max-average-subarray", "starts best at 0, wrong when every sum is negative", `def findMaxAverage(nums, k):
    window = sum(nums[:k])
    best = 0
    for right in range(k, len(nums)):
        window += nums[right] - nums[right - k]
        best = max(best, window)
    return best / k
`},
		{"drill-max-average-subarray", "divides with //", `def findMaxAverage(nums, k):
    window = sum(nums[:k])
    best = window
    for right in range(k, len(nums)):
        window += nums[right] - nums[right - k]
        best = max(best, window)
    return best // k
`},
		{"drill-max-average-subarray", "recomputes each window's sum from scratch", `def findMaxAverage(nums, k):
    return max(sum(nums[i:i + k]) for i in range(len(nums) - k + 1)) / k
`},
		{"drill-max-average-subarray", "never compares the last window", `def findMaxAverage(nums, k):
    window = sum(nums[:k])
    best = window
    for right in range(k, len(nums)):
        best = max(best, window)
        window += nums[right] - nums[right - k]
    return best / k
`},
		{"drill-max-average-subarray", "never compares the first window", `def findMaxAverage(nums, k):
    window = sum(nums[:k])
    best = -10 ** 9
    for right in range(k, len(nums)):
        window += nums[right] - nums[right - k]
        best = max(best, window)
    return best / k
`},
		{"drill-distinct-in-windows", "leaves zero-count keys in the dictionary", `def countDistinctWindows(nums, k):
    counts = {}
    result = []
    for right, value in enumerate(nums):
        counts[value] = counts.get(value, 0) + 1
        if right >= k:
            leaving = nums[right - k]
            counts[leaving] -= 1
            pass
        if right >= k - 1:
            result.append(len(counts))
    return result
`},
		{"drill-distinct-in-windows", "removes the wrong element (still inside the window)", `def countDistinctWindows(nums, k):
    counts = {}
    result = []
    for right, value in enumerate(nums):
        counts[value] = counts.get(value, 0) + 1
        if right >= k:
            leaving = nums[right - k + 1]
            counts[leaving] -= 1
            if counts[leaving] == 0:
                del counts[leaving]
        if right >= k - 1:
            result.append(len(counts))
    return result
`},
		{"drill-distinct-in-windows", "uses a set, dropping a value while another copy is inside", `def countDistinctWindows(nums, k):
    seen = set()
    result = []
    for right, value in enumerate(nums):
        seen.add(value)
        if right >= k:
            seen.discard(nums[right - k])
        if right >= k - 1:
            result.append(len(seen))
    return result
`},
		{"drill-distinct-in-windows", "rebuilds a set for every window", `def countDistinctWindows(nums, k):
    return [len(set(nums[i:i + k])) for i in range(len(nums) - k + 1)]
`},
		{"drill-distinct-in-windows", "records before the first full window exists", `def countDistinctWindows(nums, k):
    counts = {}
    result = []
    for right, value in enumerate(nums):
        counts[value] = counts.get(value, 0) + 1
        if right >= k:
            leaving = nums[right - k]
            counts[leaving] -= 1
            if counts[leaving] == 0:
                del counts[leaving]
        if right >= 0:
            result.append(len(counts))
    return result
`},
		{"drill-min-subarray-sum", "shrinks with if instead of while", `def minSubArrayLen(target, nums):
    left = 0
    total = 0
    best = len(nums) + 1
    for right, value in enumerate(nums):
        total += value
        if total >= target:
            best = min(best, right - left + 1)
            total -= nums[left]
            left += 1
    return 0 if best == len(nums) + 1 else best
`},
		{"drill-min-subarray-sum", "uses right - left as the length", `def minSubArrayLen(target, nums):
    left = 0
    total = 0
    best = len(nums) + 1
    for right, value in enumerate(nums):
        total += value
        while total >= target:
            best = min(best, right - left)
            total -= nums[left]
            left += 1
    return 0 if best == len(nums) + 1 else best
`},
		{"drill-min-subarray-sum", "records the length after removing nums[left]", `def minSubArrayLen(target, nums):
    left = 0
    total = 0
    best = len(nums) + 1
    for right, value in enumerate(nums):
        total += value
        while total >= target:
            total -= nums[left]
            left += 1
            best = min(best, right - left + 1)
    return 0 if best == len(nums) + 1 else best
`},
		{"drill-min-subarray-sum", "returns the starting best when nothing reaches the target", `def minSubArrayLen(target, nums):
    left = 0
    total = 0
    best = len(nums) + 1
    for right, value in enumerate(nums):
        total += value
        while total >= target:
            best = min(best, right - left + 1)
            total -= nums[left]
            left += 1
    return best
`},
		{"drill-min-subarray-sum", "restarts the sum from every start index", `def minSubArrayLen(target, nums):
    best = 0
    for start in range(len(nums)):
        total = 0
        for end in range(start, len(nums)):
            total += nums[end]
            if total >= target:
                if best == 0 or end - start + 1 < best:
                    best = end - start + 1
                break
    return best
`},
		{"drill-min-subarray-sum", "requires sum > target instead of >=", `def minSubArrayLen(target, nums):
    left = 0
    total = 0
    best = len(nums) + 1
    for right, value in enumerate(nums):
        total += value
        while total > target:
            best = min(best, right - left + 1)
            total -= nums[left]
            left += 1
    return 0 if best == len(nums) + 1 else best
`},
		{"drill-longest-sum-at-most-k", "never lets the window become empty", `def longestSubarrayAtMost(nums, k):
    left = 0
    total = 0
    best = 0
    for right, value in enumerate(nums):
        total += value
        while left < right and total > k:
            total -= nums[left]
            left += 1
        best = max(best, right - left + 1)
    return best
`},
		{"drill-longest-sum-at-most-k", "uses right - left as the length", `def longestSubarrayAtMost(nums, k):
    left = 0
    total = 0
    best = 0
    for right, value in enumerate(nums):
        total += value
        while total > k:
            total -= nums[left]
            left += 1
        best = max(best, right - left)
    return best
`},
		{"drill-longest-sum-at-most-k", "records the length before shrinking", `def longestSubarrayAtMost(nums, k):
    left = 0
    total = 0
    best = 0
    for right, value in enumerate(nums):
        total += value
        best = max(best, right - left + 1)
        while total > k:
            total -= nums[left]
            left += 1
    return best
`},
		{"drill-longest-sum-at-most-k", "restarts the sum from every start index", `def longestSubarrayAtMost(nums, k):
    best = 0
    for start in range(len(nums)):
        total = 0
        for end in range(start, len(nums)):
            total += nums[end]
            if total > k:
                break
            best = max(best, end - start + 1)
    return best
`},
		{"drill-longest-sum-at-most-k", "shrinks only while sum >= k, dropping valid windows", `def longestSubarrayAtMost(nums, k):
    left = 0
    total = 0
    best = 0
    for right, value in enumerate(nums):
        total += value
        while total >= k:
            total -= nums[left]
            left += 1
        best = max(best, right - left + 1)
    return best
`},
		{"drill-subarray-product-less-than-k", "skips the k <= 1 guard", `def numSubarrayProductLessThanK(nums, k):
    pass
    left = 0
    product = 1
    count = 0
    for right, value in enumerate(nums):
        product *= value
        while product >= k:
            product //= nums[left]
            left += 1
        count += right - left + 1
    return count
`},
		{"drill-subarray-product-less-than-k", "adds 1 per right instead of right - left + 1", `def numSubarrayProductLessThanK(nums, k):
    if k <= 1:
        return 0
    left = 0
    product = 1
    count = 0
    for right, value in enumerate(nums):
        product *= value
        while product >= k:
            product //= nums[left]
            left += 1
        count += 1
    return count
`},
		{"drill-subarray-product-less-than-k", "adds right - left, missing the single-element run", `def numSubarrayProductLessThanK(nums, k):
    if k <= 1:
        return 0
    left = 0
    product = 1
    count = 0
    for right, value in enumerate(nums):
        product *= value
        while product >= k:
            product //= nums[left]
            left += 1
        count += right - left
    return count
`},
		{"drill-subarray-product-less-than-k", "shrinks with if instead of while", `def numSubarrayProductLessThanK(nums, k):
    if k <= 1:
        return 0
    left = 0
    product = 1
    count = 0
    for right, value in enumerate(nums):
        product *= value
        if product >= k:
            product //= nums[left]
            left += 1
        count += right - left + 1
    return count
`},
		{"drill-subarray-product-less-than-k", "accepts a product equal to k", `def numSubarrayProductLessThanK(nums, k):
    if k <= 1:
        return 0
    left = 0
    product = 1
    count = 0
    for right, value in enumerate(nums):
        product *= value
        while product > k:
            product //= nums[left]
            left += 1
        count += right - left + 1
    return count
`},
		{"drill-subarray-product-less-than-k", "counts every run one at a time", `def numSubarrayProductLessThanK(nums, k):
    count = 0
    for start in range(len(nums)):
        product = 1
        for end in range(start, len(nums)):
            product *= nums[end]
            if product >= k:
                break
            count += 1
    return count
`},
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
