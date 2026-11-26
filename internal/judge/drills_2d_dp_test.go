package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestTwoDDPDrillsRejectMistakes checks that each 2-D dynamic programming
// foundations drill fails the mistake its tests and notes promise to catch, so
// a drill can't be passed without doing the thing it drills. Plain recursion
// and the check-every-start version of the repeated subarray are caught by the
// large time-limited cases.
func TestTwoDDPDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-grid-paths-obstacles", "fills the first row with 1s ignoring obstacles", `def uniquePathsWithObstacles(obstacleGrid):
    rows, cols = len(obstacleGrid), len(obstacleGrid[0])
    paths = [[0] * cols for _ in range(rows)]
    for col in range(cols):
        paths[0][col] = 1
    for row in range(rows):
        paths[row][0] = 1
    for row in range(1, rows):
        for col in range(1, cols):
            if obstacleGrid[row][col] == 0:
                paths[row][col] = paths[row - 1][col] + paths[row][col - 1]
    return paths[rows - 1][cols - 1] if obstacleGrid[rows - 1][cols - 1] == 0 else 0
`},
		{"drill-grid-paths-obstacles", "stops the first row at an obstacle but fills the first column with 1s", `def uniquePathsWithObstacles(obstacleGrid):
    rows, cols = len(obstacleGrid), len(obstacleGrid[0])
    paths = [[0] * cols for _ in range(rows)]
    for col in range(cols):
        if obstacleGrid[0][col] == 1:
            break
        paths[0][col] = 1
    for row in range(rows):
        paths[row][0] = 1
    for row in range(1, rows):
        for col in range(1, cols):
            if obstacleGrid[row][col] == 0:
                paths[row][col] = paths[row - 1][col] + paths[row][col - 1]
    return paths[rows - 1][cols - 1]
`},
		{"drill-grid-paths-obstacles", "ignores an obstacle on the start cell", `def uniquePathsWithObstacles(obstacleGrid):
    rows, cols = len(obstacleGrid), len(obstacleGrid[0])
    paths = [[0] * cols for _ in range(rows)]
    paths[0][0] = 1
    for row in range(rows):
        for col in range(cols):
            if row == 0 and col == 0:
                continue
            if obstacleGrid[row][col] == 1:
                continue
            above = paths[row - 1][col] if row > 0 else 0
            left = paths[row][col - 1] if col > 0 else 0
            paths[row][col] = above + left
    return paths[rows - 1][cols - 1]
`},
		{"drill-grid-paths-obstacles", "table one row short", `def uniquePathsWithObstacles(obstacleGrid):
    rows, cols = len(obstacleGrid), len(obstacleGrid[0])
    paths = [[0] * cols for _ in range(rows - 1)]
    paths[0][0] = 1
    for row in range(rows):
        for col in range(cols):
            if obstacleGrid[row][col] == 0 and (row or col):
                above = paths[row - 1][col] if row > 0 else 0
                left = paths[row][col - 1] if col > 0 else 0
                paths[row][col] = above + left
    return paths[rows - 1][cols - 1]
`},
		{"drill-grid-paths-obstacles", "plain recursion with no memo", `def uniquePathsWithObstacles(obstacleGrid):
    def routes(row, col):
        if row < 0 or col < 0 or obstacleGrid[row][col] == 1:
            return 0
        if row == 0 and col == 0:
            return 1
        return routes(row - 1, col) + routes(row, col - 1)

    return routes(len(obstacleGrid) - 1, len(obstacleGrid[0]) - 1)
`},
		{"drill-min-path-sum", "goes greedily to the cheaper neighbour", `def minPathSum(grid):
    row = col = 0
    total = grid[0][0]
    while row < len(grid) - 1 or col < len(grid[0]) - 1:
        if row == len(grid) - 1:
            col += 1
        elif col == len(grid[0]) - 1:
            row += 1
        elif grid[row + 1][col] < grid[row][col + 1]:
            row += 1
        else:
            col += 1
        total += grid[row][col]
    return total
`},
		{"drill-min-path-sum", "treats the border as cost 0 so the first row and column take min with 0", `def minPathSum(grid):
    rows, cols = len(grid), len(grid[0])
    cost = [[0] * cols for _ in range(rows)]
    for row in range(rows):
        for col in range(cols):
            above = cost[row - 1][col] if row > 0 else 0
            left = cost[row][col - 1] if col > 0 else 0
            cost[row][col] = grid[row][col] + min(above, left)
    return cost[rows - 1][cols - 1]
`},
		{"drill-min-path-sum", "no guard on the first row and column, so index -1 wraps around", `def minPathSum(grid):
    rows, cols = len(grid), len(grid[0])
    cost = [[0] * cols for _ in range(rows)]
    cost[0][0] = grid[0][0]
    for row in range(rows):
        for col in range(cols):
            if row == 0 and col == 0:
                continue
            cost[row][col] = grid[row][col] + min(cost[row - 1][col], cost[row][col - 1])
    return cost[rows - 1][cols - 1]
`},
		{"drill-min-path-sum", "adds the neighbours instead of taking the min", `def minPathSum(grid):
    rows, cols = len(grid), len(grid[0])
    cost = [[0] * cols for _ in range(rows)]
    for row in range(rows):
        for col in range(cols):
            above = cost[row - 1][col] if row > 0 else 0
            left = cost[row][col - 1] if col > 0 else 0
            cost[row][col] = grid[row][col] + max(above, left) if row and col else grid[row][col] + above + left
    return cost[rows - 1][cols - 1]
`},
		{"drill-min-path-sum", "plain recursion with no memo", `def minPathSum(grid):
    def cost(row, col):
        if row == 0 and col == 0:
            return grid[0][0]
        best = float("inf")
        if row > 0:
            best = min(best, cost(row - 1, col))
        if col > 0:
            best = min(best, cost(row, col - 1))
        return grid[row][col] + best

    return cost(len(grid) - 1, len(grid[0]) - 1)
`},
		{"drill-triangle", "goes greedily to the smaller child from the top", `def minimumTotal(triangle):
    col = 0
    total = triangle[0][0]
    for row in range(1, len(triangle)):
        if triangle[row][col + 1] < triangle[row][col]:
            col += 1
        total += triangle[row][col]
    return total
`},
		{"drill-triangle", "one-row version sweeps the columns from right to left", `def minimumTotal(triangle):
    best = triangle[-1][:]
    for row in range(len(triangle) - 2, -1, -1):
        for col in range(row, -1, -1):
            best[col] = triangle[row][col] + min(best[col], best[col + 1])
    return best[0]
`},
		{"drill-triangle", "one-row version reads best[col - 1] instead of best[col + 1]", `def minimumTotal(triangle):
    best = triangle[-1][:]
    for row in range(len(triangle) - 2, -1, -1):
        for col in range(row + 1):
            left = best[col - 1] if col > 0 else best[col]
            best[col] = triangle[row][col] + min(best[col], left)
    return best[0]
`},
		{"drill-triangle", "starts from the last row itself instead of the one above it, so the loop reads past the row", `def minimumTotal(triangle):
    depth = len(triangle)
    best = [row[:] for row in triangle]
    for row in range(depth - 1, -1, -1):
        for col in range(row + 1):
            best[row][col] = triangle[row][col] + min(best[row + 1][col], best[row + 1][col + 1])
    return best[0][0]
`},
		{"drill-triangle", "plain recursion with no memo", `def minimumTotal(triangle):
    def down(row, col):
        if row == len(triangle):
            return 0
        return triangle[row][col] + min(down(row + 1, col), down(row + 1, col + 1))

    return down(0, 0)
`},
		{"drill-repeated-subarray", "carries the best value forward on a mismatch (the subsequence table)", `def findLength(nums1, nums2):
    rows, cols = len(nums1), len(nums2)
    run = [[0] * (cols + 1) for _ in range(rows + 1)]
    for i in range(1, rows + 1):
        for j in range(1, cols + 1):
            if nums1[i - 1] == nums2[j - 1]:
                run[i][j] = run[i - 1][j - 1] + 1
            else:
                run[i][j] = max(run[i - 1][j], run[i][j - 1])
    return run[rows][cols]
`},
		{"drill-repeated-subarray", "returns the bottom-right cell instead of the largest cell", `def findLength(nums1, nums2):
    rows, cols = len(nums1), len(nums2)
    run = [[0] * (cols + 1) for _ in range(rows + 1)]
    for i in range(1, rows + 1):
        for j in range(1, cols + 1):
            if nums1[i - 1] == nums2[j - 1]:
                run[i][j] = run[i - 1][j - 1] + 1
    return run[rows][cols]
`},
		{"drill-repeated-subarray", "no padding row and column, so index -1 wraps around", `def findLength(nums1, nums2):
    rows, cols = len(nums1), len(nums2)
    run = [[0] * cols for _ in range(rows)]
    longest = 0
    for i in range(rows):
        for j in range(cols):
            if nums1[i] == nums2[j]:
                run[i][j] = run[i - 1][j - 1] + 1
                longest = max(longest, run[i][j])
    return longest
`},
		{"drill-repeated-subarray", "one-row version sweeps left to right", `def findLength(nums1, nums2):
    run = [0] * (len(nums2) + 1)
    longest = 0
    for i in range(1, len(nums1) + 1):
        for j in range(1, len(nums2) + 1):
            if nums1[i - 1] == nums2[j - 1]:
                run[j] = run[j - 1] + 1
                longest = max(longest, run[j])
            else:
                run[j] = 0
    return longest
`},
		{"drill-repeated-subarray", "one-row version forgets to reset on a mismatch", `def findLength(nums1, nums2):
    run = [0] * (len(nums2) + 1)
    longest = 0
    for i in range(1, len(nums1) + 1):
        for j in range(len(nums2), 0, -1):
            if nums1[i - 1] == nums2[j - 1]:
                run[j] = run[j - 1] + 1
                longest = max(longest, run[j])
    return longest
`},
		{"drill-repeated-subarray", "extends a match from every pair of start positions", `def findLength(nums1, nums2):
    longest = 0
    for i in range(len(nums1)):
        for j in range(len(nums2)):
            length = 0
            while i + length < len(nums1) and j + length < len(nums2) and nums1[i + length] == nums2[j + length]:
                length += 1
            longest = max(longest, length)
    return longest
`},
		{"drill-knapsack-01", "one-row version sweeps the capacity upwards (unbounded knapsack)", `def maxValue(weights, values, capacity):
    best = [0] * (capacity + 1)
    for weight, value in zip(weights, values):
        for room in range(weight, capacity + 1):
            best[room] = max(best[room], best[room - weight] + value)
    return best[capacity]
`},
		{"drill-knapsack-01", "takes the item without checking that it fits", `def maxValue(weights, values, capacity):
    items = len(weights)
    best = [[0] * (capacity + 1) for _ in range(items + 1)]
    for i in range(1, items + 1):
        for room in range(capacity + 1):
            best[i][room] = max(best[i - 1][room], best[i - 1][room - weights[i - 1]] + values[i - 1])
    return best[items][capacity]
`},
		{"drill-knapsack-01", "table has capacity columns instead of capacity + 1", `def maxValue(weights, values, capacity):
    items = len(weights)
    best = [[0] * capacity for _ in range(items + 1)]
    for i in range(1, items + 1):
        for room in range(capacity):
            best[i][room] = best[i - 1][room]
            if weights[i - 1] <= room:
                best[i][room] = max(best[i][room], best[i - 1][room - weights[i - 1]] + values[i - 1])
    return best[items][capacity]
`},
		{"drill-knapsack-01", "picks items by value per weight", `def maxValue(weights, values, capacity):
    order = sorted(range(len(weights)), key=lambda i: values[i] / weights[i], reverse=True)
    total = 0
    for i in order:
        if weights[i] <= capacity:
            capacity -= weights[i]
            total += values[i]
    return total
`},
		{"drill-knapsack-01", "reads the same row instead of the previous row", `def maxValue(weights, values, capacity):
    items = len(weights)
    best = [[0] * (capacity + 1) for _ in range(items + 1)]
    for i in range(1, items + 1):
        for room in range(capacity + 1):
            best[i][room] = best[i - 1][room]
            if weights[i - 1] <= room:
                best[i][room] = max(best[i][room], best[i][room - weights[i - 1]] + values[i - 1])
    return best[items][capacity]
`},
		{"drill-knapsack-01", "plain take-or-skip recursion with no memo", `def maxValue(weights, values, capacity):
    def best(i, room):
        if i == len(weights):
            return 0
        result = best(i + 1, room)
        if weights[i] <= room:
            result = max(result, values[i] + best(i + 1, room - weights[i]))
        return result

    return best(0, capacity)
`},
		{"drill-count-squares", "takes the max of the three neighbours", `def countSquares(matrix):
    rows, cols = len(matrix), len(matrix[0])
    side = [[0] * (cols + 1) for _ in range(rows + 1)]
    total = 0
    for i in range(1, rows + 1):
        for j in range(1, cols + 1):
            if matrix[i - 1][j - 1] == 1:
                side[i][j] = 1 + max(side[i - 1][j], side[i][j - 1], side[i - 1][j - 1])
                total += side[i][j]
    return total
`},
		{"drill-count-squares", "leaves out the diagonal neighbour", `def countSquares(matrix):
    rows, cols = len(matrix), len(matrix[0])
    side = [[0] * (cols + 1) for _ in range(rows + 1)]
    total = 0
    for i in range(1, rows + 1):
        for j in range(1, cols + 1):
            if matrix[i - 1][j - 1] == 1:
                side[i][j] = 1 + min(side[i - 1][j], side[i][j - 1])
                total += side[i][j]
    return total
`},
		{"drill-count-squares", "returns the largest side instead of the number of squares", `def countSquares(matrix):
    rows, cols = len(matrix), len(matrix[0])
    side = [[0] * (cols + 1) for _ in range(rows + 1)]
    best = 0
    for i in range(1, rows + 1):
        for j in range(1, cols + 1):
            if matrix[i - 1][j - 1] == 1:
                side[i][j] = 1 + min(side[i - 1][j], side[i][j - 1], side[i - 1][j - 1])
                best = max(best, side[i][j])
    return best
`},
		{"drill-count-squares", "counts each corner once instead of adding its side", `def countSquares(matrix):
    rows, cols = len(matrix), len(matrix[0])
    side = [[0] * (cols + 1) for _ in range(rows + 1)]
    total = 0
    for i in range(1, rows + 1):
        for j in range(1, cols + 1):
            if matrix[i - 1][j - 1] == 1:
                side[i][j] = 1 + min(side[i - 1][j], side[i][j - 1], side[i - 1][j - 1])
                total += 1
    return total
`},
		{"drill-count-squares", "plain recursion with no memo", `def countSquares(matrix):
    def side(i, j):
        if i < 0 or j < 0 or matrix[i][j] == 0:
            return 0
        return 1 + min(side(i - 1, j), side(i, j - 1), side(i - 1, j - 1))

    return sum(side(i, j) for i in range(len(matrix)) for j in range(len(matrix[0])))
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
