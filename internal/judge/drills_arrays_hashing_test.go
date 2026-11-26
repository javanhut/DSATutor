package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestArraysHashingDrillsRejectMistakes checks that each arrays and hashing
// foundations drill fails the mistake its tests and notes promise to catch, so
// a drill can't be passed without doing the thing it drills. Re-counting or
// re-summing for every value or query, and checking every pair, are caught by
// the large time-limited cases.
func TestArraysHashingDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-count-frequencies", "overwrites the count with 1 instead of adding", `def countFrequencies(nums):
    counts = {}
    for value in nums:
        counts[value] = 1
    return [[value, count] for value, count in counts.items()]
`},
		{"drill-count-frequencies", "adds to a key that does not exist yet", `def countFrequencies(nums):
    counts = {}
    for value in nums:
        counts[value] += 1
    return [[value, count] for value, count in counts.items()]
`},
		{"drill-count-frequencies", "returns one pair per element, not per distinct value", `def countFrequencies(nums):
    return [[value, nums.count(value)] for value in nums]
`},
		{"drill-count-frequencies", "puts the count before the value", `def countFrequencies(nums):
    counts = {}
    for value in nums:
        counts[value] = counts.get(value, 0) + 1
    return [[count, value] for value, count in counts.items()]
`},
		{"drill-count-frequencies", "recounts every value with a scan", `def countFrequencies(nums):
    return [[value, nums.count(value)] for value in set(nums)]
`},
		{"drill-first-unique-character", "returns the last unique character", `def firstUniqChar(s):
    counts = {}
    for ch in s:
        counts[ch] = counts.get(ch, 0) + 1
    for i in range(len(s) - 1, -1, -1):
        if counts[s[i]] == 1:
            return i
    return -1
`},
		{"drill-first-unique-character", "returns the character, not its index", `def firstUniqChar(s):
    counts = {}
    for ch in s:
        counts[ch] = counts.get(ch, 0) + 1
    for ch in s:
        if counts[ch] == 1:
            return ch
    return -1
`},
		{"drill-first-unique-character", "returns the first character not seen before", `def firstUniqChar(s):
    seen = set()
    for i, ch in enumerate(s):
        if ch not in seen:
            return i
        seen.add(ch)
    return -1
`},
		{"drill-first-unique-character", "returns 0 when nothing is unique", `def firstUniqChar(s):
    counts = {}
    for ch in s:
        counts[ch] = counts.get(ch, 0) + 1
    for i, ch in enumerate(s):
        if counts[ch] == 1:
            return i
    return 0
`},
		{"drill-first-unique-character", "compares every pair in a double loop", `def firstUniqChar(s):
    for i in range(len(s)):
        unique = True
        for j in range(len(s)):
            if i != j and s[i] == s[j]:
                unique = False
                break
        if unique:
            return i
    return -1
`},
		{"drill-range-sum-queries", "prefix array has no leading zero", `def rangeSums(nums, queries):
    prefix = []
    total = 0
    for value in nums:
        total += value
        prefix.append(total)
    return [prefix[right] - prefix[left - 1] for left, right in queries]
`},
		{"drill-range-sum-queries", "excludes the right end", `def rangeSums(nums, queries):
    prefix = [0]
    for value in nums:
        prefix.append(prefix[-1] + value)
    return [prefix[right] - prefix[left] for left, right in queries]
`},
		{"drill-range-sum-queries", "re-sums every range", `def rangeSums(nums, queries):
    return [sum(nums[left:right + 1]) for left, right in queries]
`},
		{"drill-subarray-sum-equals-k", "map is not seeded with the empty prefix", `def subarraySum(nums, k):
    seen = {}
    prefix = 0
    count = 0
    for value in nums:
        prefix += value
        count += seen.get(prefix - k, 0)
        seen[prefix] = seen.get(prefix, 0) + 1
    return count
`},
		{"drill-subarray-sum-equals-k", "stores only whether a prefix was seen", `def subarraySum(nums, k):
    seen = {0}
    prefix = 0
    count = 0
    for value in nums:
        prefix += value
        if prefix - k in seen:
            count += 1
        seen.add(prefix)
    return count
`},
		{"drill-subarray-sum-equals-k", "records the prefix before looking up", `def subarraySum(nums, k):
    seen = {0: 1}
    prefix = 0
    count = 0
    for value in nums:
        prefix += value
        seen[prefix] = seen.get(prefix, 0) + 1
        count += seen.get(prefix - k, 0)
    return count
`},
		{"drill-subarray-sum-equals-k", "sliding window that assumes positive values", `def subarraySum(nums, k):
    count = 0
    left = 0
    total = 0
    for right, value in enumerate(nums):
        total += value
        while total > k and left <= right:
            total -= nums[left]
            left += 1
        if total == k:
            count += 1
    return count
`},
		{"drill-subarray-sum-equals-k", "checks every start and end pair", `def subarraySum(nums, k):
    count = 0
    for start in range(len(nums)):
        total = 0
        for end in range(start, len(nums)):
            total += nums[end]
            if total == k:
                count += 1
    return count
`},
		{"drill-count-pairs-with-difference", "looks up only value - k", `def countKDifference(nums, k):
    seen = {}
    pairs = 0
    for value in nums:
        pairs += seen.get(value - k, 0)
        seen[value] = seen.get(value, 0) + 1
    return pairs
`},
		{"drill-count-pairs-with-difference", "looks up only value + k", `def countKDifference(nums, k):
    seen = {}
    pairs = 0
    for value in nums:
        pairs += seen.get(value + k, 0)
        seen[value] = seen.get(value, 0) + 1
    return pairs
`},
		{"drill-count-pairs-with-difference", "stores only whether a value was seen", `def countKDifference(nums, k):
    seen = set()
    pairs = 0
    for value in nums:
        if value - k in seen:
            pairs += 1
        if value + k in seen:
            pairs += 1
        seen.add(value)
    return pairs
`},
		{"drill-count-pairs-with-difference", "counts every pair from both ends", `def countKDifference(nums, k):
    counts = {}
    for value in nums:
        counts[value] = counts.get(value, 0) + 1
    pairs = 0
    for value in nums:
        pairs += counts.get(value - k, 0) + counts.get(value + k, 0)
    return pairs
`},
		{"drill-count-pairs-with-difference", "checks every pair", `def countKDifference(nums, k):
    pairs = 0
    for i in range(len(nums)):
        for j in range(i + 1, len(nums)):
            if abs(nums[i] - nums[j]) == k:
                pairs += 1
    return pairs
`},
		{"drill-rotate-array", "skips k % n", `def rotate(nums, k):
    size = len(nums)

    def reverse(left, right):
        while left < right:
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1

    reverse(0, size - 1)
    reverse(0, k - 1)
    reverse(k, size - 1)
`},
		{"drill-rotate-array", "rotates to the left", `def rotate(nums, k):
    size = len(nums)
    k %= size

    def reverse(left, right):
        while left < right:
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1

    reverse(0, k - 1)
    reverse(k, size - 1)
    reverse(0, size - 1)
`},
		{"drill-rotate-array", "rebinds the name instead of changing the list", `def rotate(nums, k):
    k %= len(nums)
    if k:
        nums = nums[-k:] + nums[:-k]
`},
		{"drill-rotate-array", "shifts one place at a time without reducing k", `def rotate(nums, k):
    for _ in range(k):
        nums.insert(0, nums.pop())
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
