package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestOneDDPDrillsRejectMistakes checks that each 1-D dynamic programming
// foundations drill fails the mistake its tests and notes promise to catch, so
// a drill can't be passed without doing the thing it drills. Plain recursion
// is caught by the large time-limited cases, and a recursive solution that is
// too deep for the stack by the n = 50000 tiling case.
func TestOneDDPDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-fibonacci", "plain recursion with no memo", `def fib(n):
    if n < 2:
        return n
    return fib(n - 1) + fib(n - 2)
`},
		{"drill-fibonacci", "uses 1 as the base case for n = 0", `def fib(n):
    if n < 1:
        return 1
    table = [0] * (n + 1)
    table[0] = 1
    table[1] = 1
    for i in range(2, n + 1):
        table[i] = table[i - 1] + table[i - 2]
    return table[n]
`},
		{"drill-fibonacci", "writes table[1] with no check for n = 0", `def fib(n):
    table = [0] * (n + 1)
    table[1] = 1
    for i in range(2, n + 1):
        table[i] = table[i - 1] + table[i - 2]
    return table[n]
`},
		{"drill-fibonacci", "table one cell too short", `def fib(n):
    if n < 2:
        return n
    table = [0] * n
    table[1] = 1
    for i in range(2, n):
        table[i] = table[i - 1] + table[i - 2]
    return table[n]
`},
		{"drill-climb-one-two-three", "adds only the two previous entries", `def waysToClimb(n):
    ways = [0] * (n + 1)
    ways[0] = 1
    for i in range(1, n + 1):
        for step in (1, 2):
            if i - step >= 0:
                ways[i] += ways[i - step]
    return ways[n]
`},
		{"drill-climb-one-two-three", "sets ways[0] = 0", `def waysToClimb(n):
    ways = [0] * (n + 1)
    ways[0] = 0
    for i in range(1, n + 1):
        for step in (1, 2, 3):
            if i - step >= 0:
                ways[i] += ways[i - step]
    return ways[n]
`},
		{"drill-climb-one-two-three", "table one cell too short", `def waysToClimb(n):
    ways = [0] * n
    ways[0] = 1
    for i in range(1, n):
        for step in (1, 2, 3):
            if i - step >= 0:
                ways[i] += ways[i - step]
    return ways[n]
`},
		{"drill-climb-one-two-three", "no guard on i - step, so negative indexes wrap around", `def waysToClimb(n):
    ways = [0] * (n + 1)
    ways[0] = 1
    for i in range(1, n + 1):
        for step in (1, 2, 3):
            ways[i] += ways[i - step]
    return ways[n]
`},
		{"drill-climb-one-two-three", "plain recursion with no memo", `def waysToClimb(n):
    if n < 0:
        return 0
    if n == 0:
        return 1
    return waysToClimb(n - 1) + waysToClimb(n - 2) + waysToClimb(n - 3)
`},
		{"drill-min-cost-climbing-stairs", "returns the entry for the last step, not the top", `def minCostClimbingStairs(cost):
    n = len(cost)
    best = [0] * (n + 1)
    for i in range(2, n + 1):
        best[i] = min(best[i - 1] + cost[i - 1], best[i - 2] + cost[i - 2])
    return best[n - 1]
`},
		{"drill-min-cost-climbing-stairs", "greedily steps to the cheaper of the next two", `def minCostClimbingStairs(cost):
    n = len(cost)
    i = 0 if cost[0] <= cost[1] else 1
    total = 0
    while i < n:
        total += cost[i]
        if i + 2 < n and cost[i + 1] > cost[i + 2]:
            i += 2
        elif i + 1 >= n:
            break
        else:
            i += 1
    return total
`},
		{"drill-min-cost-climbing-stairs", "can only start on step 0", `def minCostClimbingStairs(cost):
    n = len(cost)
    best = [0] * (n + 1)
    best[1] = cost[0]
    for i in range(2, n + 1):
        best[i] = min(best[i - 1] + cost[i - 1], best[i - 2] + cost[i - 2])
    return best[n]
`},
		{"drill-min-cost-climbing-stairs", "plain recursion with no memo", `def minCostClimbingStairs(cost):
    n = len(cost)

    def solve(i):
        if i <= 1:
            return 0
        return min(solve(i - 1) + cost[i - 1], solve(i - 2) + cost[i - 2])

    return solve(n)
`},
		{"drill-count-tilings", "coefficient 1 on the n - 2 term", `def countTilings(n):
    MOD = 10 ** 9 + 7
    ways = [1] * (n + 1)
    for i in range(2, n + 1):
        ways[i] = (ways[i - 1] + ways[i - 2]) % MOD
    return ways[n]
`},
		{"drill-count-tilings", "memoised recursion is too deep for big n", `def countTilings(n):
    MOD = 10 ** 9 + 7
    memo = {0: 1, 1: 1}

    def solve(i):
        if i not in memo:
            memo[i] = (solve(i - 1) + 2 * solve(i - 2)) % MOD
        return memo[i]

    return solve(n)
`},
		{"drill-count-tilings", "plain recursion with no memo", `def countTilings(n):
    MOD = 10 ** 9 + 7
    if n < 2:
        return 1
    return (countTilings(n - 1) + 2 * countTilings(n - 2)) % MOD
`},
		{"drill-count-tilings", "never takes the remainder", `def countTilings(n):
    ways = [1] * (n + 1)
    for i in range(2, n + 1):
        ways[i] = ways[i - 1] + 2 * ways[i - 2]
    return ways[n]
`},
		{"drill-count-tilings", "sets ways[0] = 0", `def countTilings(n):
    MOD = 10 ** 9 + 7
    ways = [1] * (n + 1)
    ways[0] = 0
    for i in range(2, n + 1):
        ways[i] = (ways[i - 1] + 2 * ways[i - 2]) % MOD
    return ways[n]
`},
		{"drill-max-subarray", "clamps at 0, allowing an empty subarray", `def maxSubArray(nums):
    ending_here = 0
    best = 0
    for value in nums:
        ending_here = max(0, ending_here + value)
        best = max(best, ending_here)
    return best
`},
		{"drill-max-subarray", "returns the last table entry instead of the largest", `def maxSubArray(nums):
    ending = [0] * len(nums)
    ending[0] = nums[0]
    for i in range(1, len(nums)):
        ending[i] = max(nums[i], ending[i - 1] + nums[i])
    return ending[-1]
`},
		{"drill-max-subarray", "always extends, never starts fresh", `def maxSubArray(nums):
    ending = [0] * len(nums)
    ending[0] = nums[0]
    for i in range(1, len(nums)):
        ending[i] = ending[i - 1] + nums[i]
    return max(ending)
`},
		{"drill-max-subarray", "tries every subarray", `def maxSubArray(nums):
    best = nums[0]
    for start in range(len(nums)):
        total = 0
        for end in range(start, len(nums)):
            total += nums[end]
            best = max(best, total)
    return best
`},
		{"drill-coin-change-ways", "amounts outside, coins inside, which counts orderings", `def change(amount, coins):
    ways = [0] * (amount + 1)
    ways[0] = 1
    for total in range(1, amount + 1):
        for coin in coins:
            if total >= coin:
                ways[total] += ways[total - coin]
    return ways[amount]
`},
		{"drill-coin-change-ways", "sets ways[0] = 0", `def change(amount, coins):
    ways = [0] * (amount + 1)
    for coin in coins:
        for total in range(coin, amount + 1):
            ways[total] += ways[total - coin]
    return ways[amount]
`},
		{"drill-coin-change-ways", "returns -1 for an impossible amount", `def change(amount, coins):
    ways = [0] * (amount + 1)
    ways[0] = 1
    for coin in coins:
        for total in range(coin, amount + 1):
            ways[total] += ways[total - coin]
    return ways[amount] if ways[amount] else -1
`},
		{"drill-coin-change-ways", "sweeps downwards, using each coin at most once", `def change(amount, coins):
    ways = [0] * (amount + 1)
    ways[0] = 1
    for coin in coins:
        for total in range(amount, coin - 1, -1):
            ways[total] += ways[total - coin]
    return ways[amount]
`},
		{"drill-coin-change-ways", "plain recursion with no memo", `def change(amount, coins):
    def count(left, index):
        if left == 0:
            return 1
        if left < 0 or index == len(coins):
            return 0
        return count(left - coins[index], index) + count(left, index + 1)

    return count(amount, 0)
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
