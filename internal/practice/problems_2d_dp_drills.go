package practice

// TwoDDPDrills are the 2-D dynamic programming foundations: name what
// table[i][j] means, write the recurrence from the last step, fill a table
// with two indexes from its base row and column, and shrink it to one row where
// the fill order allows. The ladder runs grid paths with obstacles (counting
// with blocked cells and the first row and column), minimum path sum (a min
// instead of a sum), a triangle filled from the bottom (and a one-row version
// whose sweep direction matters), the longest common block (reset to 0, answer
// anywhere in the table), the 0/1 knapsack (the second index is a capacity, and
// the one-row version must sweep it downwards) and counting all-ones squares
// (three neighbours). The cases with large inputs carry a 5 second limit so that
// plain recursion (exponential) or checking every start (cubic) fails.
var TwoDDPDrills = []*Problem{
	{
		ID:              "drill-grid-paths-obstacles",
		Number:          63,
		Kind:            KindDrill,
		Order:           1,
		Title:           "Unique Paths with Obstacles",
		Slug:            "drill-grid-paths-obstacles",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Dynamic Programming", "Matrix", "Foundations"},
		RelatedChapters: []int{16, 15},
		Description:     "You are given a grid `obstacleGrid` with `rows` rows and `cols` columns, where `0` is a free cell and `1` is an obstacle. A robot starts at the top-left cell and wants to reach the bottom-right cell. Each move goes one cell **down** or one cell **right**, and it can never enter an obstacle. Return the number of different routes. If the start or the finish is an obstacle, the answer is 0. The answer fits in a 64-bit integer.\n\nThis is the grid DP in its plainest form (the ladder's Unique Paths with one extra rule). Define the **state** first: `paths[i][j]` is the number of routes from the start to cell `(i, j)`. A route into `(i, j)` came from the cell above or from the cell to the left, and those two groups share no route, so the **recurrence** is `paths[i][j] = paths[i - 1][j] + paths[i][j - 1]`, except that an obstacle cell has `paths = 0`.\n\nThe **base case** is the start: `paths[0][0] = 1` if it is free. The first row and the first column are special because they have only one neighbour: a cell in the first row can only be entered from its left, so one obstacle there makes every cell to its right unreachable. Your table gets that for free if cells outside the grid count as 0 and obstacle cells are left at 0.",
		Constraints:     []string{"1 <= rows, cols <= 20", "obstacleGrid[i][j] is 0 or 1"},
		Examples: []Example{
			{Input: "obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]", Output: "2", Explanation: "The obstacle in the middle forces the route around it: right, right, down, down or down, down, right, right."},
			{Input: "obstacleGrid = [[0,1],[0,0]]", Output: "1", Explanation: "The top-right cell is blocked, so the only route is down and then right."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, Expected: 2, Note: "example 1: one obstacle in the middle leaves two routes"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{0, 1}, {0, 0}}}, Expected: 1, Note: "example 2: the only route goes down first"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{0}}}, Expected: 1, Hidden: true, Note: "a single free cell: one path of length zero"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{1}}}, Expected: 0, Hidden: true, Note: "the start cell is an obstacle, so there is no path at all"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{1, 0}}}, Expected: 0, Hidden: true, Note: "the start is blocked in a one-row grid: the answer is 0 even though every other cell is free"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{0, 0}, {0, 1}}}, Expected: 0, Hidden: true, Note: "the finish is blocked: 0"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{0, 1, 0, 0}, {0, 0, 0, 0}}}, Expected: 1, Hidden: true, Note: "an obstacle in the first row blocks every cell to its right in that row, but the second row still reaches them from below"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{0, 0}, {1, 0}, {0, 0}}}, Expected: 1, Hidden: true, Note: "an obstacle in the first column blocks every cell below it in that column"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{0}, {1}, {0}}}, Expected: 0, Hidden: true, Note: "a one-column grid with an obstacle in the middle: no path"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{0, 0, 0, 0, 0}}}, Expected: 1, Hidden: true, Note: "a single free row: exactly one path"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 1}}}, Expected: 0, Hidden: true, Note: "blocked finish in a larger grid"},
			{Input: map[string]interface{}{"obstacleGrid": [][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 1}, {1, 0, 0, 0}}}, Expected: 3, Hidden: true, Note: "several obstacles in the interior"},
			{Input: map[string]interface{}{"obstacleGrid": tddEmptyGrid(18, 18)}, Expected: tddPathsObstacles(tddEmptyGrid(18, 18)), Hidden: true, Timeout: 5000, Note: "18 by 18 with no obstacles: plain recursion makes billions of calls and runs out of time, and the answer is over two billion"},
		},
		TimeComplexity:  "O(rows * cols)",
		SpaceComplexity: "O(rows * cols)",
		StarterCode:     "def uniquePathsWithObstacles(obstacleGrid):\n    # Return the number of down/right routes from the top-left to the bottom-right cell\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Think about the last move into a cell. Which two cells could it have come from, and what happens to the count if one of them is an obstacle or off the grid?"},
			{Level: 2, Type: "insight", Content: "Every route into cell (i, j) ends with a move down from (i - 1, j) or a move right from (i, j - 1), so paths[i][j] is the sum of those two counts, and an obstacle cell has no routes through it. Fill the table from paths[0][0] = 1 row by row so both inputs are ready. This is 2-D dynamic programming; see Chapter 16.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. State: paths[i][j] is the number of routes from the start to cell (i, j).\n2. Base case: paths[0][0] = 1 when the start is free; every other cell begins at 0.\n3. Recurrence: skip obstacle cells (they stay 0); otherwise add the count from above (if there is a row above) and from the left (if there is a column to the left).\n4. Answer: return the bottom-right cell."},
			{Level: 4, Type: "skeleton", Content: "def uniquePathsWithObstacles(obstacleGrid):\n    rows, cols = len(obstacleGrid), len(obstacleGrid[0])\n    paths = [[0] * cols for _ in range(rows)]\n    for row in range(rows):\n        for col in range(cols):\n            if obstacleGrid[row][col] == 1:\n                continue\n            if row == 0 and col == 0:\n                paths[row][col] = ____\n                continue\n            from_above = paths[row - 1][col] if row > 0 else 0\n            from_left = ____ if col > 0 else 0\n            paths[row][col] = from_above + from_left\n    return paths[rows - 1][cols - 1]"},
		},
		Solution: Solution{
			Code:           "def uniquePathsWithObstacles(obstacleGrid):\n    rows, cols = len(obstacleGrid), len(obstacleGrid[0])\n    paths = [[0] * cols for _ in range(rows)]\n    for row in range(rows):\n        for col in range(cols):\n            if obstacleGrid[row][col] == 1:\n                continue\n            if row == 0 and col == 0:\n                paths[row][col] = 1\n                continue\n            from_above = paths[row - 1][col] if row > 0 else 0\n            from_left = paths[row][col - 1] if col > 0 else 0\n            paths[row][col] = from_above + from_left\n    return paths[rows - 1][cols - 1]",
			Explanation:    "This is the first table with two indexes: paths[i][j] counts routes to one cell, filled row by row from the base case paths[0][0] = 1. It builds on 1-D DP counting (Climbing Stairs: sum the cells you could have come from) with two neighbours instead of two earlier steps, and the ladder's Unique Paths is this drill with no obstacles. Leaving obstacle cells at 0 and treating off-grid neighbours as 0 means the first row and column need no separate loops, and an obstacle blocks everything behind it automatically.",
			BruteForce:     "Recurse from the finish: routes(i, j) = routes(i - 1, j) + routes(i, j - 1), stopping at obstacles and the border. That explores every route one at a time, and the number of routes grows like a binomial coefficient (about 2 * 10^9 on an empty 18 by 18 grid), so it is far too slow. The waste is that routes(i, j) is recomputed for every route that passes through that cell.",
			KeyInsight:     "There are only rows * cols different cells, and the number of routes to a cell never changes, so compute each cell once from its two earlier neighbours.",
			WhyItWorks:     "State: paths[i][j] is the number of down/right routes from (0, 0) to (i, j) that avoid obstacles. For a free cell other than the start, every such route has a last move, from (i - 1, j) or from (i, j - 1). Removing it leaves a route to that neighbour, and adding the move to any route to a free neighbour gives a different route to (i, j), so the two groups are disjoint and complete and paths[i][j] is their sum. An obstacle cell has no routes, and a missing neighbour (above the first row or left of the first column) contributes 0. Cells are filled row by row, so both neighbours are final when they are read. The base case paths[0][0] = 1 is the empty route, and it is 0 if the start is an obstacle, which then keeps every cell at 0.",
			CommonMistakes: []string{"Filling the first row (or column) with 1s without checking for obstacles. An obstacle at (0, 2) must make (0, 3), (0, 4), ... unreachable, because the only way into them is from the left; stop at the first obstacle.", "Setting `paths[0][0] = 1` without checking whether the start is an obstacle. A blocked start means 0 routes, and the answer must stay 0.", "Making the table one row or column short, so the bottom-right cell does not exist (an IndexError) or the answer is read from the wrong cell. The table has the same shape as the grid.", "Writing the recursion with no memo. It is right on small grids and takes exponential time on an 18 by 18 grid."},
			TimeComplexity: "O(rows * cols)", SpaceComplexity: "O(rows * cols)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "paths[i][j] will hold the number of routes to cell (i, j), so the table has the same shape as the grid and starts at 0. A cell that stays 0 is unreachable or an obstacle.", CodeSnippet: "paths = [[0] * cols for _ in range(rows)]", LineStart: 3, LineEnd: 3},
				{Title: "Skip obstacles", Explanation: "No route can pass through an obstacle, so leave its count at 0. This is also what blocks the rest of a row or column: the cells after it read a 0 from their only neighbour.", CodeSnippet: "if obstacleGrid[row][col] == 1:\n    continue", LineStart: 6, LineEnd: 7},
				{Title: "Seed the start", Explanation: "The start has exactly one route, the empty one. Because the obstacle check came first, a blocked start keeps 0.", CodeSnippet: "if row == 0 and col == 0:\n    paths[row][col] = 1\n    continue", LineStart: 8, LineEnd: 10},
				{Title: "Add the two ways in", Explanation: "A cell is entered from above or from the left; a neighbour outside the grid counts as 0. On the first example the table reads 1 1 1 / 1 0 1 / 1 1 2.", CodeSnippet: "from_above = paths[row - 1][col] if row > 0 else 0\nfrom_left = paths[row][col - 1] if col > 0 else 0\npaths[row][col] = from_above + from_left", LineStart: 11, LineEnd: 13},
				{Title: "Read the answer", Explanation: "The bottom-right cell is the number of routes to the finish.", CodeSnippet: "return paths[rows - 1][cols - 1]", LineStart: 14, LineEnd: 14},
			},
			Alternatives: []AltSolution{
				{Name: "One row", Code: "def uniquePathsWithObstacles(obstacleGrid):\n    cols = len(obstacleGrid[0])\n    paths = [0] * cols\n    paths[0] = 1\n    for row in obstacleGrid:\n        for col in range(cols):\n            if row[col] == 1:\n                paths[col] = 0\n            elif col > 0:\n                paths[col] += paths[col - 1]\n    return paths[cols - 1]", TimeComplexity: "O(rows * cols)", SpaceComplexity: "O(cols)", Explanation: "Row i only reads row i - 1 (the same column, still holding the old value) and its own left neighbour (already updated this row), so one array is enough: paths[col] += paths[col - 1] means old value (from above) plus new left. An obstacle resets the cell to 0. Seeding paths[0] = 1 before the first row makes the start cell behave like every other cell."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does one obstacle in the first row make every cell to its right in that row unreachable, when the table has no special code for that?", Answer: "A cell in the first row has no row above it, so its from_above is 0 and its count is just its left neighbour's count. The obstacle's own count is 0, so the next cell copies 0, and so on along the row. The rule is simply the recurrence applied with a missing neighbour counted as 0."},
			{Prompt: "What changes if the robot may also move diagonally down-right?", Answer: "Add a third term: paths[i][j] = from_above + from_left + from_diagonal, where the diagonal is paths[i - 1][j - 1] when both indexes exist. The state, the base case and the row-by-row order are unchanged, because the new neighbour is also in an earlier row."},
		},
	},
	{
		ID:              "drill-min-path-sum",
		Number:          64,
		Kind:            KindDrill,
		Order:           2,
		Title:           "Minimum Path Sum",
		Slug:            "drill-min-path-sum",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Dynamic Programming", "Matrix", "Foundations"},
		RelatedChapters: []int{16, 15},
		Description:     "You are given a `rows` by `cols` grid `grid` of non-negative numbers. Start at the top-left cell and walk to the bottom-right cell, moving only **down** or **right**. The cost of a route is the sum of the numbers on every cell you step on, including the first and the last. Return the smallest possible cost.\n\nThe **state** is `cost[i][j]`, the smallest cost of any route from the start to cell `(i, j)` (including that cell). The last move into `(i, j)` came from above or from the left, so the **recurrence** is `cost[i][j] = grid[i][j] + min(cost[i - 1][j], cost[i][j - 1])`. This is the counting drill with the sum of counts replaced by a min of costs.\n\nThe **base cases** are the borders. `cost[0][0] = grid[0][0]`. A cell in the first row has no cell above, so it can only come from its left: `cost[0][j] = cost[0][j - 1] + grid[0][j]`; the first column is the same with the cell above. Do not let a missing neighbour count as cost 0, which would look like the cheapest choice.",
		Constraints:     []string{"1 <= rows, cols <= 100", "0 <= grid[i][j] <= 200"},
		Examples: []Example{
			{Input: "grid = [[1,3,1],[1,5,1],[4,2,1]]", Output: "7", Explanation: "The route 1, 3, 1, 1, 1 (right, right, down, down) costs 7. Going to the smaller neighbour at each step (down first) would cost 9."},
			{Input: "grid = [[1,2,3],[4,5,6]]", Output: "12", Explanation: "Right, right, down costs 1 + 2 + 3 + 6 = 12."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"grid": [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}}, Expected: 7, Note: "example 1: the cheapest route is not the greedy one"},
			{Input: map[string]interface{}{"grid": [][]int{{1, 2, 3}, {4, 5, 6}}}, Expected: 12, Note: "example 2: a 2 by 3 grid"},
			{Input: map[string]interface{}{"grid": [][]int{{5}}}, Expected: 5, Hidden: true, Note: "a single cell: its own cost"},
			{Input: map[string]interface{}{"grid": [][]int{{1, 2, 3, 4}}}, Expected: 10, Hidden: true, Note: "a single row: the only route is left to right, so the answer is the sum"},
			{Input: map[string]interface{}{"grid": [][]int{{2}, {3}, {4}}}, Expected: 9, Hidden: true, Note: "a single column: the only route is top to bottom"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 0}, {0, 0}}}, Expected: 0, Hidden: true, Note: "all zeros: the answer is 0"},
			{Input: map[string]interface{}{"grid": [][]int{{3, 1, 9}, {8, 7, 1}, {2, 6, 1}}}, Expected: 13, Hidden: true, Note: "the first row is cheap but the column below is not: first-row and first-column cells have one neighbour only"},
			{Input: map[string]interface{}{"grid": [][]int{{9, 1, 1, 1}, {9, 9, 9, 1}, {1, 1, 1, 1}}}, Expected: 14, Hidden: true, Note: "a detour: the cheapest route goes right across the top and then down"},
			{Input: map[string]interface{}{"grid": [][]int{{1, 9}, {1, 9}, {1, 1}}}, Expected: 4, Hidden: true, Note: "the cheapest route goes down the first column and along the bottom row"},
			{Input: map[string]interface{}{"grid": [][]int{{7, 1}, {2, 9}}}, Expected: 17, Hidden: true, Note: "a 2 by 2 grid where the left and right route differ"},
			{Input: map[string]interface{}{"grid": tddCostGrid(18, 18, 7)}, Expected: tddMinPath(tddCostGrid(18, 18, 7)), Hidden: true, Timeout: 5000, Note: "18 by 18: plain recursion tries billions of routes and runs out of time"},
			{Input: map[string]interface{}{"grid": tddCostGrid(80, 90, 11)}, Expected: tddMinPath(tddCostGrid(80, 90, 11)), Hidden: true, Timeout: 5000, Note: "80 by 90: the table version does 7,200 cells of work"},
		},
		TimeComplexity:  "O(rows * cols)",
		SpaceComplexity: "O(rows * cols)",
		StarterCode:     "def minPathSum(grid):\n    # Return the smallest sum of cells on a down/right route from top-left to bottom-right\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "For the last cell, the route came from the cell above or the cell to its left. Which of the two would you want to have arrived from, and what do you need to know about each?"},
			{Level: 2, Type: "insight", Content: "The cheapest route to a cell extends the cheaper of the cheapest routes to its two neighbours, because the cost of the rest of the route does not depend on how you got there. So cost[i][j] = grid[i][j] + min(cost[i - 1][j], cost[i][j - 1]), filled row by row. A cell on the first row or column has only one neighbour. This is 2-D dynamic programming; see Chapter 16.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. State: cost[i][j] is the cheapest cost of a route from (0, 0) to (i, j).\n2. Base case: cost[0][0] = grid[0][0].\n3. Recurrence: on the first row the cheapest way in is from the left; on the first column from above; elsewhere take the smaller of the two, then add the cell's own number.\n4. Answer: return the bottom-right cell."},
			{Level: 4, Type: "skeleton", Content: "def minPathSum(grid):\n    rows, cols = len(grid), len(grid[0])\n    cost = [[0] * cols for _ in range(rows)]\n    for row in range(rows):\n        for col in range(cols):\n            if row == 0 and col == 0:\n                best_before = 0\n            elif row == 0:\n                best_before = cost[row][col - 1]\n            elif col == 0:\n                best_before = ____\n            else:\n                best_before = min(____, cost[row][col - 1])\n            cost[row][col] = grid[row][col] + best_before\n    return cost[rows - 1][cols - 1]"},
		},
		Solution: Solution{
			Code:           "def minPathSum(grid):\n    rows, cols = len(grid), len(grid[0])\n    cost = [[0] * cols for _ in range(rows)]\n    for row in range(rows):\n        for col in range(cols):\n            if row == 0 and col == 0:\n                best_before = 0\n            elif row == 0:\n                best_before = cost[row][col - 1]\n            elif col == 0:\n                best_before = cost[row - 1][col]\n            else:\n                best_before = min(cost[row - 1][col], cost[row][col - 1])\n            cost[row][col] = grid[row][col] + best_before\n    return cost[rows - 1][cols - 1]",
			Explanation:    "This builds on the obstacles drill by replacing the sum of two counts with a min of two costs: the same table, the same fill order, the same special first row and column. The three-way split into start, first row or column, and interior is spelled out so each border case is visible; it is the same shape the ladder's Unique Paths uses, and the min-of-neighbours idea returns in Triangle next and in Edit Distance later.",
			BruteForce:     "Try every route: recurse cost(i, j) = grid[i][j] + min(cost(i - 1, j), cost(i, j - 1)) with no memory. The number of calls is the number of routes, which is exponential in rows + cols (billions on an 18 by 18 grid). Going to the cheaper neighbour at each step is fast but wrong: a cheap cell can lead into expensive ones. The waste in the recursion is recomputing cost(i, j) for every route that reaches that cell.",
			KeyInsight:     "The cheapest route to a cell depends only on the cheapest routes to its two neighbours, not on the route itself, so one number per cell is enough.",
			WhyItWorks:     "State: cost[i][j] is the minimum, over all down/right routes from (0, 0) to (i, j), of the sum of the cells on the route. Any such route's last step comes from (i - 1, j) or (i, j - 1). The part of the route before that step is itself a route to that neighbour, and replacing it with a cheaper route to the same neighbour gives a cheaper route to (i, j) (the last cell's cost is added either way). So the best route to (i, j) is the cell's number plus the smaller of the two neighbours' bests. On the first row or column only one neighbour exists, so the min is just that neighbour. Row-by-row order has both neighbours final before they are read.",
			CommonMistakes: []string{"Treating a missing neighbour as 0 in the interior formula (for example padding the table with zeros and taking `min` over the border). On the first row that min picks the 0 from outside the grid instead of the real left neighbour, so the first row and column are undercounted. Handle them separately, or pad with infinity.", "Going greedily to the cheaper of the two next cells. That is a decision made with no knowledge of what follows; example 1 gives 9 instead of 7. Compute the best way into each cell instead.", "Reading `cost[row - 1][col]` when `row` is 0: Python's index -1 reads the last row and silently uses a cell that has not been computed yet.", "Writing the recursion with no memo. It is correct on small grids and takes exponential time on an 18 by 18 grid."},
			TimeComplexity: "O(rows * cols)", SpaceComplexity: "O(rows * cols)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "cost[i][j] holds the cheapest total from the start up to and including cell (i, j). The table has the grid's shape.", CodeSnippet: "cost = [[0] * cols for _ in range(rows)]", LineStart: 3, LineEnd: 3},
				{Title: "Find the cheapest way in", Explanation: "The start has nothing before it. The first row can only be entered from the left and the first column only from above. Everywhere else take the smaller of the two.", CodeSnippet: "if row == 0 and col == 0:\n    best_before = 0\nelif row == 0:\n    best_before = cost[row][col - 1]\nelif col == 0:\n    best_before = cost[row - 1][col]\nelse:\n    best_before = min(cost[row - 1][col], cost[row][col - 1])", LineStart: 6, LineEnd: 13},
				{Title: "Add the cell itself", Explanation: "Stepping on the cell costs its number on top of the cheapest way in. On example 1 the table becomes 1 4 5 / 2 7 6 / 6 8 7.", CodeSnippet: "cost[row][col] = grid[row][col] + best_before", LineStart: 14, LineEnd: 14},
				{Title: "Read the answer", Explanation: "The bottom-right cell is the cheapest cost of the whole route.", CodeSnippet: "return cost[rows - 1][cols - 1]", LineStart: 15, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{Name: "One row", Code: "def minPathSum(grid):\n    cols = len(grid[0])\n    cost = [float(\"inf\")] * cols\n    cost[0] = 0\n    for row in grid:\n        for col in range(cols):\n            if col > 0:\n                cost[col] = min(cost[col], cost[col - 1])\n            cost[col] += row[col]\n    return cost[cols - 1]", TimeComplexity: "O(rows * cols)", SpaceComplexity: "O(cols)", Explanation: "cost[col] before the update is the cell above (the previous row's value) and cost[col - 1] is the cell to the left (already updated for this row), so one array holds the previous row and the current row together. Infinity stands for the missing cell above row 0, so it never wins the min."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the cheapest route to a cell built from the cheapest routes to its neighbours, and not from some other route to them?", Answer: "What happens after the neighbour does not depend on how you reached it: the remaining cost is the same cell and step either way. So if a route to the neighbour were not the cheapest, swapping in the cheapest one would make the whole route cheaper, which means the best whole route never uses a non-best prefix."},
			{Prompt: "What would you change to also return the route itself, not only its cost?", Answer: "Keep the table, then walk back from the bottom-right cell: at each cell step to whichever neighbour has the smaller cost value (the one the min chose), until you reach the start, and reverse the list. Alternatively store a parent pointer when filling each cell."},
		},
	},
	{
		ID:              "drill-triangle",
		Number:          120,
		Kind:            KindDrill,
		Order:           3,
		Title:           "Triangle",
		Slug:            "drill-triangle",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Dynamic Programming", "Matrix", "Foundations"},
		RelatedChapters: []int{16, 15},
		Description:     "You are given `triangle`, a list of rows where row `r` has `r + 1` numbers (the first row has one number). Start at the top and move down one row at a time. From the number at position `c` in row `r` you may step to position `c` or position `c + 1` in row `r + 1`. The cost of a path is the sum of the numbers it visits. Return the smallest cost of any path from the top row to the bottom row.\n\nHere a bottom-up table is cleaner than a top-down one. The **state** is `best[r][c]`, the smallest cost of a path that starts at `(r, c)` and goes down to the last row. The **recurrence** is `best[r][c] = triangle[r][c] + min(best[r + 1][c], best[r + 1][c + 1])`. The **base row** is the last row: `best` there is the numbers themselves, since the path ends at once. Fill the rows from the bottom upwards; the answer is `best[0][0]`.\n\nBecause row `r` only reads row `r + 1`, a table with one row is enough once you have the full one working: see the alternative solution and mind the order you sweep the columns.",
		Constraints:     []string{"1 <= number of rows <= 200", "-10000 <= triangle[r][c] <= 10000"},
		Examples: []Example{
			{Input: "triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]", Output: "11", Explanation: "The path 2 + 3 + 5 + 1 = 11 is the cheapest."},
			{Input: "triangle = [[-10]]", Output: "-10", Explanation: "A single row: the path is the one number."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"triangle": [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}}, Expected: 11, Note: "example 1: the cheapest path is 2, 3, 5, 1"},
			{Input: map[string]interface{}{"triangle": [][]int{{-10}}}, Expected: -10, Note: "example 2: a single row"},
			{Input: map[string]interface{}{"triangle": [][]int{{1}, {2, 3}}}, Expected: 3, Hidden: true, Note: "two rows: the smaller of the two children"},
			{Input: map[string]interface{}{"triangle": [][]int{{1}, {2, 3}, {10, 10, 1}}}, Expected: 5, Hidden: true, Note: "greedy trap: the smaller child 2 leads only to 10s, while 3 leads to a 1"},
			{Input: map[string]interface{}{"triangle": [][]int{{-1}, {-2, -3}, {-4, -5, -6}}}, Expected: -10, Hidden: true, Note: "all negative numbers: the answer is the most negative route, not the first one found"},
			{Input: map[string]interface{}{"triangle": [][]int{{5}, {5, 5}, {5, 5, 5}, {5, 5, 5, 5}}}, Expected: 20, Hidden: true, Note: "all equal: every path has the same sum"},
			{Input: map[string]interface{}{"triangle": [][]int{{0}, {0, 0}, {0, 0, 0}}}, Expected: 0, Hidden: true, Note: "all zeros"},
			{Input: map[string]interface{}{"triangle": [][]int{{1}, {1, 100}, {100, 1, 100}, {100, 100, 1, 100}, {1, 100, 100, 100, 1}}}, Expected: 104, Hidden: true, Note: "the cheap path zigzags along the left edge: edge cells have only one parent or child"},
			{Input: map[string]interface{}{"triangle": [][]int{{3}, {7, 4}, {2, 4, 6}, {8, 5, 9, 3}}}, Expected: 16, Hidden: true, Note: "a second small triangle"},
			{Input: map[string]interface{}{"triangle": tddTriangle(30)}, Expected: tddMinTriangle(tddTriangle(30)), Hidden: true, Timeout: 5000, Note: "30 rows: there are 2^29 paths, so plain recursion runs out of time"},
			{Input: map[string]interface{}{"triangle": tddTriangle(200)}, Expected: tddMinTriangle(tddTriangle(200)), Hidden: true, Timeout: 5000, Note: "200 rows, the largest input: only about 20,000 cells of work for the table"},
		},
		TimeComplexity:  "O(n^2)",
		SpaceComplexity: "O(n^2)",
		StarterCode:     "def minimumTotal(triangle):\n    # Return the smallest sum of a top-to-bottom path\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "A path from the top has to pick a child at every row. If someone told you the cheapest cost from each of the two children down to the bottom, how would you pick the best choice at the top number?"},
			{Level: 2, Type: "insight", Content: "Work from the bottom. For a number at (r, c), the cheapest path from it down is its own value plus the smaller of the cheapest paths from its two children (r + 1, c) and (r + 1, c + 1). The last row is the base: the path ends there. Filling upwards means both children are final when read, and the top cell holds the answer. This is 2-D dynamic programming; see Chapter 16.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. State: best[r][c] is the smallest cost of a path from (r, c) down to the bottom row.\n2. Base case: copy the last row; those paths end immediately.\n3. Recurrence: for each earlier row, from the second to last up to the first, best[r][c] = triangle[r][c] + min of the two cells below.\n4. Answer: return best[0][0]."},
			{Level: 4, Type: "skeleton", Content: "def minimumTotal(triangle):\n    depth = len(triangle)\n    best = [row[:] for row in triangle]\n    for row in range(depth - 2, -1, -1):\n        for col in range(row + 1):\n            best[row][col] = triangle[row][col] + min(____, best[row + 1][col + 1])\n    return ____"},
		},
		Solution: Solution{
			Code:           "def minimumTotal(triangle):\n    depth = len(triangle)\n    best = [row[:] for row in triangle]\n    for row in range(depth - 2, -1, -1):\n        for col in range(row + 1):\n            best[row][col] = triangle[row][col] + min(best[row + 1][col], best[row + 1][col + 1])\n    return best[0][0]",
			Explanation:    "This is the min-path-sum recurrence turned upside down: instead of the cheapest way into a cell from its two parents, the cheapest way out of a cell through its two children. Going from the bottom up makes the table's base row the input's last row and puts the answer in one cell at the top, with no border cases at all (every cell above the last row has two children). The one-row alternative shows when a table can shrink: row r reads only row r + 1.",
			BruteForce:     "Try every path: recurse down from (0, 0), taking the smaller of the two children. A triangle with n rows has 2^(n - 1) paths, so 30 rows is already half a billion. Choosing the smaller child at each step (greedy) is fast but wrong, since a small number can lead to large ones. The waste in the recursion is that two paths that meet in the same cell redo all the work below it.",
			KeyInsight:     "The cheapest path from a cell down depends only on the cheapest paths from its two children, and the cells in a row share children, so compute each cell once from the row below.",
			WhyItWorks:     "State: best[r][c] is the minimum sum over all paths from (r, c) to the last row. In the last row the only path is the cell itself, which is the base. For an earlier row, every path from (r, c) takes its own number and then continues from (r + 1, c) or (r + 1, c + 1); the best continuation from each child is best of that child, so the best total is the cell's number plus the smaller of the two. Rows are filled from the last to the first, so both children are final when read. Every cell is computed once, and the answer is the best path from the only cell in the top row.",
			CommonMistakes: []string{"Going greedily from the top to the smaller child at each row. A small number can lead only to large ones; on [[1],[2,3],[10,10,1]] the greedy path costs 13 and the best costs 5. Decide each cell from the finished values below it.", "In the one-row version, sweeping the columns from right to left. Cell `c` needs the old `best[c + 1]` (the child's value), but a right-to-left sweep has already overwritten it with this row's value. Sweep left to right, so `best[c + 1]` is still the row below when `best[c]` is computed.", "Reading `best[row + 1][col + 1]` on the last row, or looping over `range(depth)` and stepping off the table. The base row has no children: start the loop at the second to last row.", "Writing the recursion with no memo. It passes the small cases and takes 2^(rows - 1) steps on 30 rows."},
			TimeComplexity: "O(n^2)", SpaceComplexity: "O(n^2)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state and base row", Explanation: "best starts as a copy of the triangle. In the last row it is already right: a path that starts there is only that number.", CodeSnippet: "best = [row[:] for row in triangle]", LineStart: 3, LineEnd: 3},
				{Title: "Go upwards", Explanation: "Process the rows from the second to last to the first, so the row below is final. Row `row` has `row + 1` cells.", CodeSnippet: "for row in range(depth - 2, -1, -1):\n    for col in range(row + 1):", LineStart: 4, LineEnd: 5},
				{Title: "Pick the cheaper child", Explanation: "A cell's cheapest path down is its own number plus the smaller of the two children below it. On example 1 the rows become 11 / 9 10 / 7 6 10 / 4 1 8 3.", CodeSnippet: "best[row][col] = triangle[row][col] + min(best[row + 1][col], best[row + 1][col + 1])", LineStart: 6, LineEnd: 6},
				{Title: "Read the answer", Explanation: "The single cell in the top row now holds the cheapest path from the top.", CodeSnippet: "return best[0][0]", LineStart: 7, LineEnd: 7},
			},
			Alternatives: []AltSolution{
				{Name: "One row, left to right", Code: "def minimumTotal(triangle):\n    best = triangle[-1][:]\n    for row in range(len(triangle) - 2, -1, -1):\n        for col in range(row + 1):\n            best[col] = triangle[row][col] + min(best[col], best[col + 1])\n    return best[0]", TimeComplexity: "O(n^2)", SpaceComplexity: "O(n)", Explanation: "Row r reads only row r + 1, so one array holds the row below and is updated in place into the current row. Cell col needs best[col] and best[col + 1] from the row below; sweeping left to right overwrites best[col] after both are read, and best[col + 1] has not been touched yet. Sweeping right to left would read an already updated best[col + 1]."},
				{Name: "Top-down with a memo", Code: "def minimumTotal(triangle):\n    memo = {}\n\n    def down(row, col):\n        if row == len(triangle):\n            return 0\n        if (row, col) not in memo:\n            memo[(row, col)] = triangle[row][col] + min(down(row + 1, col), down(row + 1, col + 1))\n        return memo[(row, col)]\n\n    return down(0, 0)", TimeComplexity: "O(n^2)", SpaceComplexity: "O(n^2)", Explanation: "The same recurrence as a recursive function that stores each answer. It visits each cell once instead of once per path; depth is the number of rows."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to fill the triangle from the bottom, and what would go wrong with a top-down table that stores the cheapest path into each cell?", Answer: "A top-down table works too (best path into (r, c) is its number plus the smaller of its two parents), but then the answer is the minimum over the whole last row, and the edge cells have only one parent. Going bottom-up has no edge cases and one answer cell. Both are correct because each cell's best depends only on its two neighbours in the adjacent row."},
			{Prompt: "In the one-row version, why does the sweep direction matter?", Answer: "Cell col reads best[col] and best[col + 1], both from the row below, and writes best[col]. Going left to right, best[col + 1] is not written until the next step, so it is still the old value when needed. Going right to left, best[col + 1] was already overwritten with the current row's value, which gives wrong answers. Whenever you reuse a row, ask which entries the update reads and make sure they are still old when read."},
		},
	},
	{
		ID:              "drill-repeated-subarray",
		Number:          718,
		Kind:            KindDrill,
		Order:           4,
		Title:           "Maximum Length of Repeated Subarray",
		Slug:            "drill-repeated-subarray",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Dynamic Programming", "Array", "Foundations"},
		RelatedChapters: []int{16, 15},
		Description:     "Given two integer arrays `nums1` and `nums2`, return the length of the longest block of **consecutive** elements that appears in both arrays (a common subarray, also called a longest common substring). If they share no element, return 0.\n\nThis is the table of Longest Common Subsequence (a ladder problem) with one rule changed, and the change is the point of the drill. The **state** is `run[i][j]`, the length of the longest common block that **ends exactly** at `nums1[i - 1]` and `nums2[j - 1]`. If those two elements are equal, the block extends the one ending at the previous pair: `run[i][j] = run[i - 1][j - 1] + 1`. If they differ, no block can end at this pair, so `run[i][j] = 0`: the run is **reset**, where the subsequence table would carry the best value forward.\n\nThe **base row and column** are `run[0][j] = run[i][0] = 0` (an empty prefix has no block), so the table is `(len(nums1) + 1)` by `(len(nums2) + 1)`. Because a block can end anywhere, the answer is the **largest cell** in the whole table, not the bottom-right cell.",
		Constraints:     []string{"1 <= len(nums1), len(nums2) <= 1000", "0 <= nums1[i], nums2[j] <= 100"},
		Examples: []Example{
			{Input: "nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]", Output: "3", Explanation: "The block [3, 2, 1] appears in both. run is 1, 2, 3 along the matching diagonal."},
			{Input: "nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]", Output: "5", Explanation: "Every element matches, so the run grows to 5."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums1": []int{1, 2, 3, 2, 1}, "nums2": []int{3, 2, 1, 4, 7}}, Expected: 3, Note: "example 1: the shared block 3, 2, 1"},
			{Input: map[string]interface{}{"nums1": []int{0, 0, 0, 0, 0}, "nums2": []int{0, 0, 0, 0, 0}}, Expected: 5, Note: "example 2: every element matches, so the answer is the full length"},
			{Input: map[string]interface{}{"nums1": []int{1}, "nums2": []int{1}}, Expected: 1, Hidden: true, Note: "one matching element"},
			{Input: map[string]interface{}{"nums1": []int{1}, "nums2": []int{2}}, Expected: 0, Hidden: true, Note: "nothing in common: 0, not the length of any subsequence"},
			{Input: map[string]interface{}{"nums1": []int{1, 2, 3, 4}, "nums2": []int{1, 3, 4, 5}}, Expected: 2, Hidden: true, Note: "a common subsequence of length 3 (1, 3, 4) but the longest contiguous block is 2 (3, 4): treating it like a subsequence gives 3"},
			{Input: map[string]interface{}{"nums1": []int{1, 2, 1, 2, 3}, "nums2": []int{2, 3, 1, 2, 1}}, Expected: 3, Hidden: true, Note: "a break in the middle: the run 1, 2 stops when the next elements differ, and the table must reset to 0"},
			{Input: map[string]interface{}{"nums1": []int{5, 5, 5}, "nums2": []int{5}}, Expected: 1, Hidden: true, Note: "different lengths: the run is limited by the shorter array"},
			{Input: map[string]interface{}{"nums1": []int{4, 1, 2, 3}, "nums2": []int{1, 2, 3, 4}}, Expected: 3, Hidden: true, Note: "the block sits at the start of one array and the middle of the other, so the answer is not in the last cell of the table"},
			{Input: map[string]interface{}{"nums1": []int{7, 8, 9, 1, 2}, "nums2": []int{1, 2, 9, 8, 7}}, Expected: 2, Hidden: true, Note: "matches at the end of both arrays, and a decoy at the start"},
			{Input: map[string]interface{}{"nums1": []int{1, 0, 0, 0, 1}, "nums2": []int{0, 0, 1, 0, 0}}, Expected: 3, Hidden: true, Note: "a run of 2 appears twice; a run ending later must not add to an earlier one"},
			{Input: map[string]interface{}{"nums1": []int{3, 1, 2, 3, 1, 2, 3, 4}, "nums2": []int{1, 2, 3, 4, 1, 2, 3}}, Expected: 4, Hidden: true, Note: "repeated blocks: the longest is 4"},
			{Input: map[string]interface{}{"nums1": tddRepeat(1000, 0), "nums2": tddRepeat(1000, 0)}, Expected: 1000, Hidden: true, Timeout: 5000, Note: "1000 equal elements on both sides: extending a match from every pair of positions does about 3 * 10^8 steps and runs out of time"},
			{Input: map[string]interface{}{"nums1": tddSmall(1000, 3, 21), "nums2": tddSmall(1000, 3, 22)}, Expected: tddRepeatedLen(tddSmall(1000, 3, 21), tddSmall(1000, 3, 22)), Hidden: true, Timeout: 5000, Note: "1000 elements drawn from {0, 1, 2}: many partial matches"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def findLength(nums1, nums2):\n    # Return the length of the longest block of consecutive elements shared by both arrays\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose two blocks end at nums1[i - 1] and nums2[j - 1]. What do you need to know about the elements just before them, and what if the two last elements differ?"},
			{Level: 2, Type: "insight", Content: "A common block that ends at this pair of positions exists only when the two elements are equal, and then it is one longer than the common block ending at the previous pair. If they differ the block length is 0. Unlike a subsequence, a block cannot skip an element, so nothing is carried over from other cells, and the best block can end anywhere: keep a running maximum. This is 2-D dynamic programming; see Chapter 16.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. State: run[i][j] is the length of the longest common block ending at nums1[i - 1] and nums2[j - 1]. Size the table (m + 1) by (n + 1) so row 0 and column 0 are the empty-prefix base of 0.\n2. Recurrence: equal elements give run[i - 1][j - 1] + 1; different elements give 0.\n3. Track the largest value written.\n4. Answer: return the largest value, not the bottom-right cell."},
			{Level: 4, Type: "skeleton", Content: "def findLength(nums1, nums2):\n    rows, cols = len(nums1), len(nums2)\n    run = [[0] * (cols + 1) for _ in range(rows + 1)]\n    longest = 0\n    for i in range(1, rows + 1):\n        for j in range(1, cols + 1):\n            if nums1[i - 1] == nums2[j - 1]:\n                run[i][j] = ____ + 1\n                longest = max(longest, run[i][j])\n    return ____"},
		},
		Solution: Solution{
			Code:           "def findLength(nums1, nums2):\n    rows, cols = len(nums1), len(nums2)\n    run = [[0] * (cols + 1) for _ in range(rows + 1)]\n    longest = 0\n    for i in range(1, rows + 1):\n        for j in range(1, cols + 1):\n            if nums1[i - 1] == nums2[j - 1]:\n                run[i][j] = run[i - 1][j - 1] + 1\n                longest = max(longest, run[i][j])\n    return longest",
			Explanation:    "This is the Longest Common Subsequence table with two changes that the contiguity requirement forces: a mismatch resets the cell to 0 (the table is initialised to 0 and mismatches simply leave it), and the answer is the maximum over all cells because a block can end anywhere. The padding row and column of zeros make the first row and column need no special case, which is the first of those tricks you will reuse in Edit Distance and the other string problems on the ladder.",
			BruteForce:     "For every pair of start positions (i, j), extend while the elements match and keep the longest. That is O(m * n * min(m, n)); on two arrays of 1000 equal elements it does about 3 * 10^8 steps. The waste is that the block starting at (i, j) is mostly the block starting at (i + 1, j + 1) with one more element at the front, and that overlap is recomputed from scratch.",
			KeyInsight:     "A common block ending at a pair of positions is exactly one element longer than the common block ending at the previous pair, so one number per pair, with a reset on mismatch, holds all the blocks.",
			WhyItWorks:     "State: run[i][j] is the length of the longest common block of nums1 and nums2 that ends at nums1[i - 1] and nums2[j - 1], and 0 if there is none. If the two elements differ, no block can end at this pair, so 0. If they are equal, a block ending here is one element plus a block that ends at nums1[i - 2] and nums2[j - 2] (possibly empty), and the longest such is run[i - 1][j - 1] + 1. Every common block ends at some pair, so the longest block of all is the largest cell. Row 0 and column 0 are 0 because an empty prefix ends no block.",
			CommonMistakes: []string{"Carrying the best value forward on a mismatch, `run[i][j] = max(run[i - 1][j], run[i][j - 1])`. That is the subsequence table, which allows gaps: [1,2,3,4] and [1,3,4,5] would give 3 instead of 2. A block must be contiguous, so a mismatch resets the run to 0.", "Returning the bottom-right cell `run[rows][cols]`. A block can end anywhere, and that cell is 0 whenever the last elements differ; return the largest value seen.", "Using a table with no padding row and column, so `run[i - 1][j - 1]` at `i = 0` or `j = 0` reads index -1 and wraps to the last row or column. Make the table (m + 1) by (n + 1) and start the loops at 1.", "Checking every start pair and extending (the brute force). It is correct and does about 3 * 10^8 steps on two arrays of 1000 equal elements."},
			TimeComplexity: "O(m * n)", SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state with padding", Explanation: "run[i][j] is the block ending at the i-th and j-th elements, counting from 1. The extra row and column of zeros stand for empty prefixes.", CodeSnippet: "run = [[0] * (cols + 1) for _ in range(rows + 1)]", LineStart: 3, LineEnd: 3},
				{Title: "Extend on a match", Explanation: "If the two elements are equal, the block is one longer than the block ending at the previous pair, which sits diagonally up and left. A mismatch leaves the 0 already there, which is the reset.", CodeSnippet: "if nums1[i - 1] == nums2[j - 1]:\n    run[i][j] = run[i - 1][j - 1] + 1", LineStart: 7, LineEnd: 8},
				{Title: "Remember the best block", Explanation: "A block can end at any pair, so keep the largest value ever written. On example 1 the run 1, 2, 3 builds along one diagonal.", CodeSnippet: "longest = max(longest, run[i][j])", LineStart: 9, LineEnd: 9},
				{Title: "Return the maximum", Explanation: "The answer is the longest block seen, not the corner cell.", CodeSnippet: "return longest", LineStart: 10, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{Name: "One row, right to left", Code: "def findLength(nums1, nums2):\n    run = [0] * (len(nums2) + 1)\n    longest = 0\n    for i in range(1, len(nums1) + 1):\n        for j in range(len(nums2), 0, -1):\n            if nums1[i - 1] == nums2[j - 1]:\n                run[j] = run[j - 1] + 1\n                longest = max(longest, run[j])\n            else:\n                run[j] = 0\n    return longest", TimeComplexity: "O(m * n)", SpaceComplexity: "O(n)", Explanation: "run[i][j] reads run[i - 1][j - 1], which is the previous row's cell one to the left. With a single array, sweeping j from right to left means run[j - 1] has not been overwritten yet, so it still holds the previous row. The else branch must write 0 because the cell otherwise keeps the previous row's value. Sweeping left to right would read this row's value and grow runs that are not diagonal."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does a mismatch set the cell to 0 here, when the Longest Common Subsequence table copies the best of its neighbours?", Answer: "A subsequence may skip elements, so the best length up to here is still available after a mismatch and is carried along. A block must be contiguous, so a mismatch ends every block at that pair, and the longest block that ends here is empty. The carried-forward value would let a block jump over a gap."},
			{Prompt: "Why can the single-row version only run right to left, and what goes wrong in the other direction?", Answer: "Cell j reads the previous row's cell j - 1. Right to left, cell j - 1 is overwritten only after cell j has read it. Left to right, cell j - 1 already holds the current row's value, so a run is extended using a pair from the same row, which counts matches that are not on a common diagonal."},
		},
	},
	{
		ID:              "drill-knapsack-01",
		Number:          0,
		Kind:            KindDrill,
		Order:           5,
		Title:           "0/1 Knapsack",
		Slug:            "drill-knapsack-01",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Dynamic Programming", "Knapsack", "Foundations"},
		RelatedChapters: []int{16, 15},
		Description:     "You have items numbered from 0, where item `i` weighs `weights[i]` and is worth `values[i]`, and a bag that holds a total weight of at most `capacity`. Each item can be put in the bag **at most once**. Return the largest total value that fits.\n\nThe **state** is `best[i][room]`, the largest value using only the first `i` items with `room` units of capacity. For item `i`, either skip it, `best[i - 1][room]`, or, if `weights[i - 1] <= room`, take it and add its value to the best with the remaining room, `best[i - 1][room - weight] + value`. The recurrence is the larger of the two.\n\nThe **base row** is `best[0][room] = 0` for every room (no items, no value), and the **base column** is `best[i][0] = 0` (no room), so the table has `items + 1` rows and `capacity + 1` columns. Once that works, notice that row `i` only reads row `i - 1`: you can keep a single row of `capacity + 1` cells, **but only if the capacity loop goes downwards**. See the alternative solution and think about what an upward loop does.",
		Constraints:     []string{"0 <= number of items <= 100", "1 <= weights[i] <= 100", "0 <= values[i] <= 1000", "0 <= capacity <= 1000"},
		Examples: []Example{
			{Input: "weights = [1,3,4,5], values = [1,4,5,7], capacity = 7", Output: "9", Explanation: "Take the items of weight 3 and 4: weight 7, value 4 + 5 = 9. Taking weights 1 and 5 gives only 8."},
			{Input: "weights = [10,20,30], values = [60,100,120], capacity = 50", Output: "220", Explanation: "Take the items of weight 20 and 30 for 100 + 120. The first two together give only 160."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"weights": []int{1, 3, 4, 5}, "values": []int{1, 4, 5, 7}, "capacity": 7}, Expected: 9, Note: "example 1: items of weight 3 and 4 give 4 + 5 = 9"},
			{Input: map[string]interface{}{"weights": []int{10, 20, 30}, "values": []int{60, 100, 120}, "capacity": 50}, Expected: 220, Note: "example 2: the weights 20 and 30 beat the best value per weight"},
			{Input: map[string]interface{}{"weights": []int{5}, "values": []int{10}, "capacity": 0}, Expected: 0, Hidden: true, Note: "capacity 0: nothing fits, so 0 (a table with no column 0 breaks here)"},
			{Input: map[string]interface{}{"weights": []int{}, "values": []int{}, "capacity": 5}, Expected: 0, Hidden: true, Note: "no items: 0"},
			{Input: map[string]interface{}{"weights": []int{3}, "values": []int{5}, "capacity": 9}, Expected: 5, Hidden: true, Note: "one item, large capacity: each item may be used once, so 5 and not 15"},
			{Input: map[string]interface{}{"weights": []int{6}, "values": []int{9}, "capacity": 5}, Expected: 0, Hidden: true, Note: "the only item is too heavy: a table that adds the value without checking the weight reads a negative index"},
			{Input: map[string]interface{}{"weights": []int{4, 4}, "values": []int{7, 8}, "capacity": 4}, Expected: 8, Hidden: true, Note: "two items that both fit alone but not together"},
			{Input: map[string]interface{}{"weights": []int{2, 3, 4, 5}, "values": []int{3, 4, 5, 6}, "capacity": 5}, Expected: 7, Hidden: true, Note: "weights 2 and 3 together give 7, better than any single item"},
			{Input: map[string]interface{}{"weights": []int{1, 2, 3}, "values": []int{6, 10, 12}, "capacity": 5}, Expected: 22, Hidden: true, Note: "ratio trap: the best value per weight picks weight 1 and 2 for 16, but weights 2 and 3 give 22"},
			{Input: map[string]interface{}{"weights": []int{5, 4, 6, 3}, "values": []int{10, 40, 30, 50}, "capacity": 10}, Expected: 90, Hidden: true, Note: "the best pair is the items of weight 4 and 3 (value 90), not a pair that fills the bag exactly"},
			{Input: map[string]interface{}{"weights": []int{1, 1, 1, 1}, "values": []int{1, 1, 1, 1}, "capacity": 4}, Expected: 4, Hidden: true, Note: "all items fit exactly"},
			{Input: map[string]interface{}{"weights": []int{3, 3, 3}, "values": []int{4, 4, 4}, "capacity": 7}, Expected: 8, Hidden: true, Note: "three equal items, room for two"},
			{Input: map[string]interface{}{"weights": tddRepeat(28, 1), "values": tddRepeat(28, 1), "capacity": 28}, Expected: 28, Hidden: true, Timeout: 5000, Note: "28 items of weight 1: trying take or skip on every item makes 2^28 calls and runs out of time"},
			{Input: map[string]interface{}{"weights": tddKnapWeights(100), "values": tddKnapValues(100), "capacity": 1000}, Expected: tddKnapsack(tddKnapWeights(100), tddKnapValues(100), 1000), Hidden: true, Timeout: 5000, Note: "100 items and capacity 1000: a table of 100,000 cells"},
		},
		TimeComplexity:  "O(n * capacity)",
		SpaceComplexity: "O(n * capacity)",
		StarterCode:     "def maxValue(weights, values, capacity):\n    # Return the largest total value of items (each used at most once) with total weight <= capacity\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Look at the last item only. There are two possibilities for it: it is in the bag or it is not. What would you need to know about the rest of the items in each case?"},
			{Level: 2, Type: "insight", Content: "If the last item is left out, the best is the best for the earlier items with the same room. If it is taken, you get its value plus the best for the earlier items with the room it leaves. Take the larger. The state is (how many items are considered, how much room is left), which is a 2-D table. This is 2-D dynamic programming; see Chapter 16.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. State: best[i][room] is the largest value from the first i items with at most room capacity. The table has items + 1 rows and capacity + 1 columns.\n2. Base case: row 0 is all zeros (no items); that also covers room 0.\n3. Recurrence: start from best[i - 1][room]; if the item fits (weight <= room), also try best[i - 1][room - weight] + value and keep the larger.\n4. Answer: return best[items][capacity]."},
			{Level: 4, Type: "skeleton", Content: "def maxValue(weights, values, capacity):\n    items = len(weights)\n    best = [[0] * (capacity + 1) for _ in range(items + 1)]\n    for i in range(1, items + 1):\n        weight, value = weights[i - 1], values[i - 1]\n        for room in range(capacity + 1):\n            best[i][room] = best[i - 1][room]\n            if weight <= room:\n                best[i][room] = max(best[i][room], ____ + value)\n    return ____"},
		},
		Solution: Solution{
			Code:           "def maxValue(weights, values, capacity):\n    items = len(weights)\n    best = [[0] * (capacity + 1) for _ in range(items + 1)]\n    for i in range(1, items + 1):\n        weight, value = weights[i - 1], values[i - 1]\n        for room in range(capacity + 1):\n            best[i][room] = best[i - 1][room]\n            if weight <= room:\n                best[i][room] = max(best[i][room], best[i - 1][room - weight] + value)\n    return best[items][capacity]",
			Explanation:    "This is the first table whose second index is a quantity (the remaining capacity) rather than a position in a second sequence. It builds on Coin Change II from the 1-D drills: there each coin could be reused (the loop ran upwards on one array), here each item is used once, which is why the table needs the row for 'the first i - 1 items' and the single-row version must go downwards. The ladder's Target Sum, Partition Equal Subset Sum and Coin Change II are all this table with a different value to maximise or count.",
			BruteForce:     "For each item choose take or skip, and recurse: best(i, room) = max(best(i - 1, room), value + best(i - 1, room - weight)). There are 2^n choices, so 28 items is 2.7 * 10^8 calls. The waste is that the same (i, room) pair is reached through many different sets of choices, and its answer is always the same.",
			KeyInsight:     "What matters about the earlier choices is only how many items are still undecided and how much room is left, a pair of small numbers, not which items were taken.",
			WhyItWorks:     "State: best[i][room] is the maximum total value of any subset of the first i items whose total weight is at most room. Row 0 is 0 because the only subset is empty. For i >= 1, a subset either leaves out item i - 1, which is a subset of the first i - 1 items with the same room, or contains it, which is item i - 1 plus a subset of the first i - 1 items with room reduced by its weight (only possible when weight <= room). The best subset of the two kinds is the larger of the two best values, and each kind's best is a table entry in row i - 1. Reading only the previous row is what enforces that each item is used at most once.",
			CommonMistakes: []string{"Sweeping the capacity upwards in a single-row version. Then `best[room - weight]` has already been updated for the same item, so the item can be taken again and again: with one item of weight 3 and value 5, capacity 9 would give 15. Go from `capacity` down to `weight`, or keep the full table.", "Taking the item without checking `weight <= room`. The index `room - weight` is then negative and Python reads from the end of the row, adding a value for room that does not exist.", "Making the table `capacity` columns instead of `capacity + 1`, so `best[items][capacity]` is out of range and room 0 is missing.", "Picking items by best value per weight. It is fast and wrong for a 0/1 bag: example 2 gives 160 instead of 220.", "Writing take-or-skip recursion with no memo. It is correct on small inputs and takes 2^n steps on 28 items."},
			TimeComplexity: "O(n * capacity)", SpaceComplexity: "O(n * capacity)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "best[i][room] is the best value from the first i items with room units of capacity. Row 0 (no items) and column 0 (no room) are 0, so the table is (items + 1) by (capacity + 1).", CodeSnippet: "best = [[0] * (capacity + 1) for _ in range(items + 1)]", LineStart: 3, LineEnd: 3},
				{Title: "Skip the item", Explanation: "The default for every room is to leave item i - 1 out, which is the best for the earlier items with the same room.", CodeSnippet: "best[i][room] = best[i - 1][room]", LineStart: 7, LineEnd: 7},
				{Title: "Take the item if it fits", Explanation: "If the weight fits in this room, compare skipping with taking: the item's value plus the best for the earlier items with the room left over. On example 1 with capacity 7 the last row ends in 9.", CodeSnippet: "if weight <= room:\n    best[i][room] = max(best[i][room], best[i - 1][room - weight] + value)", LineStart: 8, LineEnd: 9},
				{Title: "Read the answer", Explanation: "All items considered, all the capacity available.", CodeSnippet: "return best[items][capacity]", LineStart: 10, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{Name: "One row, capacity downwards", Code: "def maxValue(weights, values, capacity):\n    best = [0] * (capacity + 1)\n    for weight, value in zip(weights, values):\n        for room in range(capacity, weight - 1, -1):\n            best[room] = max(best[room], best[room - weight] + value)\n    return best[capacity]", TimeComplexity: "O(n * capacity)", SpaceComplexity: "O(capacity)", Explanation: "Row i reads only row i - 1 at room and at room - weight, a smaller index. Sweeping room from capacity down to weight means best[room - weight] has not been updated for this item yet, so it still holds row i - 1. Sweeping upwards would read the updated value, which lets one item be taken repeatedly (the unbounded knapsack, like Coin Change II)."},
				{Name: "Take or skip with a memo", Code: "def maxValue(weights, values, capacity):\n    memo = {}\n\n    def best(i, room):\n        if i == len(weights):\n            return 0\n        if (i, room) not in memo:\n            result = best(i + 1, room)\n            if weights[i] <= room:\n                result = max(result, values[i] + best(i + 1, room - weights[i]))\n            memo[(i, room)] = result\n        return memo[(i, room)]\n\n    return best(0, capacity)", TimeComplexity: "O(n * capacity)", SpaceComplexity: "O(n * capacity)", Explanation: "The recurrence as a recursive function over (next item, room left), with a dictionary so that each pair is solved once. The recursion depth is the number of items."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the single-row version have to sweep the capacity downwards, and what problem does an upward sweep solve instead?", Answer: "At item i the new value of best[room] needs the OLD value at room - weight (from before this item). Going downwards, the smaller index is untouched when it is read. Going upwards, it has already been updated for this item, so the item can be counted again; that computes the unbounded knapsack, where every item can be used any number of times."},
			{Prompt: "Why does picking items by value per weight fail for this problem, and when would it be right?", Answer: "Items cannot be split, so a high-ratio item can use up room that two lower-ratio items would fill exactly (weights 10, 20, 30 with capacity 50: the ratios choose 10 and 20 for 160, but 20 and 30 give 220). If items could be split into fractions, the ratio rule is optimal: that is the fractional knapsack, a greedy problem."},
		},
	},
	{
		ID:              "drill-count-squares",
		Number:          1277,
		Kind:            KindDrill,
		Order:           6,
		Title:           "Count Square Submatrices with All Ones",
		Slug:            "drill-count-squares",
		Difficulty:      "Medium",
		Category:        "2d-dp",
		Tags:            []string{"Dynamic Programming", "Matrix", "Foundations"},
		RelatedChapters: []int{16, 15},
		Description:     "Given a `rows` by `cols` matrix of `0`s and `1`s, return the number of square submatrices (blocks of consecutive rows and columns, with equal height and width) that contain only `1`s. Squares of different sizes and positions are counted separately, and a single `1` is a 1 by 1 square.\n\nThe **state** is `side[i][j]`, the side length of the largest all-ones square whose **bottom-right corner** is cell `(i, j)`. The key fact: that number is also the **count of squares** that end at that corner, because a square of side `k` ending there contains one of every smaller side, so the corner contributes exactly `side[i][j]` squares. The answer is the sum of the table.\n\nThe **recurrence** for a cell holding 1 is `side[i][j] = 1 + min(side[i - 1][j], side[i][j - 1], side[i - 1][j - 1])`: a square of side `k` at the corner needs squares of side `k - 1` ending at the cell above, the cell on the left **and** the cell diagonally up-left. A cell holding 0 has `side = 0`. The **base**: pad the table with a row and a column of zeros so the first row and column need no special case.",
		Constraints:     []string{"1 <= rows, cols <= 300", "matrix[i][j] is 0 or 1"},
		Examples: []Example{
			{Input: "matrix = [[0,1,1,1],[1,1,1,1],[0,1,1,1]]", Output: "15", Explanation: "There are 10 squares of side 1, 4 of side 2 and 1 of side 3, which is 15."},
			{Input: "matrix = [[1,0,1],[1,1,0],[1,1,0]]", Output: "7", Explanation: "There are 6 squares of side 1 and 1 of side 2 (the lower-left block), which is 7."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"matrix": [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}}, Expected: 15, Note: "example 1: ten 1 by 1, four 2 by 2 and one 3 by 3 square"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}}, Expected: 7, Note: "example 2: six 1 by 1 and one 2 by 2 square"},
			{Input: map[string]interface{}{"matrix": [][]int{{0}}}, Expected: 0, Hidden: true, Note: "a single 0: no squares"},
			{Input: map[string]interface{}{"matrix": [][]int{{1}}}, Expected: 1, Hidden: true, Note: "a single 1: one square"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 1, 1, 1}}}, Expected: 4, Hidden: true, Note: "a single row: only 1 by 1 squares, one per 1"},
			{Input: map[string]interface{}{"matrix": [][]int{{1}, {0}, {1}, {1}}}, Expected: 3, Hidden: true, Note: "a single column"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 1}, {1, 1}}}, Expected: 5, Hidden: true, Note: "a full 2 by 2: four 1 by 1 and one 2 by 2"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 1}, {1, 0}}}, Expected: 3, Hidden: true, Note: "a hole in the corner: a recurrence using only the left and the upper neighbours counts a 2 by 2 square"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}, Expected: 8, Hidden: true, Note: "a hole in the middle: a diagonal neighbour of 0 stops the square"},
			{Input: map[string]interface{}{"matrix": [][]int{{0, 0}, {0, 0}}}, Expected: 0, Hidden: true, Note: "all zeros: 0"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}}, Expected: 14, Hidden: true, Note: "a full 3 by 3: 9 + 4 + 1 = 14"},
			{Input: map[string]interface{}{"matrix": [][]int{{1, 1, 0, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 0}}}, Expected: 17, Hidden: true, Note: "mixed squares of size 2"},
			{Input: map[string]interface{}{"matrix": tddOnes(14, 14)}, Expected: tddCountSquares(tddOnes(14, 14)), Hidden: true, Timeout: 5000, Note: "14 by 14 of ones: plain recursion on the side length branches three ways with no memory and runs out of time"},
			{Input: map[string]interface{}{"matrix": tddRandomBits(300, 300, 3)}, Expected: tddCountSquares(tddRandomBits(300, 300, 3)), Hidden: true, Timeout: 5000, Note: "300 by 300 mostly ones: the table version does 90,000 cells of work"},
		},
		TimeComplexity:  "O(rows * cols)",
		SpaceComplexity: "O(rows * cols)",
		StarterCode:     "def countSquares(matrix):\n    # Return the number of all-ones square submatrices\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Fix the bottom-right corner of a square. If a 3 by 3 all-ones square ends there, what must be true of the cells just above, to the left and diagonally up-left of the corner?"},
			{Level: 2, Type: "insight", Content: "A square of side k with its corner at (i, j) is a 1 at the corner plus three overlapping squares of side k - 1 ending at the cells above, to the left and diagonally above-left. So the largest side at (i, j) is 1 plus the minimum of those three largest sides, and the number of squares ending at that corner equals that side. Add them all up. This is 2-D dynamic programming; see Chapter 16.", ChapterRef: 16},
			{Level: 3, Type: "plan", Content: "1. State: side[i][j] is the side of the largest all-ones square with its bottom-right corner at (i, j). Pad the table with a row and column of zeros.\n2. Recurrence: if the cell is 1, side = 1 + min(above, left, diagonal); if it is 0, side = 0.\n3. Each corner contributes side[i][j] squares, so add it to a running total.\n4. Answer: return the total."},
			{Level: 4, Type: "skeleton", Content: "def countSquares(matrix):\n    rows, cols = len(matrix), len(matrix[0])\n    side = [[0] * (cols + 1) for _ in range(rows + 1)]\n    total = 0\n    for i in range(1, rows + 1):\n        for j in range(1, cols + 1):\n            if matrix[i - 1][j - 1] == 1:\n                side[i][j] = 1 + min(side[i - 1][j], side[i][j - 1], ____)\n                total += ____\n    return total"},
		},
		Solution: Solution{
			Code:           "def countSquares(matrix):\n    rows, cols = len(matrix), len(matrix[0])\n    side = [[0] * (cols + 1) for _ in range(rows + 1)]\n    total = 0\n    for i in range(1, rows + 1):\n        for j in range(1, cols + 1):\n            if matrix[i - 1][j - 1] == 1:\n                side[i][j] = 1 + min(side[i - 1][j], side[i][j - 1], side[i - 1][j - 1])\n                total += side[i][j]\n    return total",
			Explanation:    "This rounds out the grid drills with a recurrence that reads three neighbours (above, left and diagonal) and takes a min, where the earlier drills read two and added or took the best. The padded row and column of zeros, the table indexed from 1, and the choice of 'ends at this corner' as the state are the same moves as in the repeated-subarray drill, and the sum over every cell (instead of the corner cell) comes from the same fact: the answer can end anywhere. It is the idea behind Maximal Square and a good model for the ladder's Edit Distance, which also reads above, left and diagonal.",
			BruteForce:     "For every cell as a top-left corner and every side length, check that the whole square is ones. That is up to O(rows * cols * min(rows, cols)^3) work with naive checks, or O(rows * cols * min(rows, cols)) with prefix sums. Recursing on the side length with no memo, side(i, j) = 1 + min of three recursive calls, makes up to 3^(i + j) calls. The waste is that squares ending at neighbouring corners overlap almost entirely, and the same corner's side is asked for many times.",
			KeyInsight:     "A square with its corner at (i, j) is built from three smaller squares ending at the three neighbours, so the largest side at a corner is determined by its neighbours' largest sides, and that side is also how many squares end there.",
			WhyItWorks:     "State: side[i][j] is the largest k such that the k by k block with bottom-right corner (i, j) is all ones (0 if the cell is 0). If the cell is 1, a k by k block is all ones exactly when the corner is 1 and the three (k - 1) by (k - 1) blocks ending at the cell above, the cell to the left and the diagonal cell are all ones, because together those three blocks and the corner cover the whole k by k block. So the largest k is 1 plus the minimum of the three neighbours' sides. Squares ending at the corner have sides 1 up to side[i][j] (each smaller block inside a large one is also all ones), so the corner adds side[i][j] squares, and each square has exactly one bottom-right corner so none is counted twice.",
			CommonMistakes: []string{"Using `max` instead of `min` of the three neighbours. The square is only as large as its smallest supporting square: a 0 anywhere around breaks it, and on a full 2 by 2 with a hole the max version grows a square that has a 0 inside.", "Leaving out the diagonal neighbour (using `min(above, left)`). The above and left squares can both be full while the cell diagonally up-left holds a 0; then there is no bigger square. A grid like [[1,1],[1,0]] already shows it.", "Returning the largest side, or the corner cell. The problem counts all squares, so add each corner's side to a running total.", "Recursing on side(i, j) with no memo. It takes 3^(i + j) calls, which a 14 by 14 matrix of ones already breaks."},
			TimeComplexity: "O(rows * cols)", SpaceComplexity: "O(rows * cols)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state with padding", Explanation: "side[i][j] is the largest all-ones square whose bottom-right corner is cell (i, j), counting cells from 1. The extra row and column of zeros mean cells outside the matrix support no square.", CodeSnippet: "side = [[0] * (cols + 1) for _ in range(rows + 1)]", LineStart: 3, LineEnd: 3},
				{Title: "Check the cell", Explanation: "A 0 cell ends no square, so its entry stays 0. Only a 1 can be a corner.", CodeSnippet: "if matrix[i - 1][j - 1] == 1:", LineStart: 7, LineEnd: 7},
				{Title: "Take the smallest support", Explanation: "The square at this corner is one larger than the smallest of the squares ending at the cell above, the cell to the left and the diagonal cell.", CodeSnippet: "side[i][j] = 1 + min(side[i - 1][j], side[i][j - 1], side[i - 1][j - 1])", LineStart: 8, LineEnd: 8},
				{Title: "Count the squares at this corner", Explanation: "A corner with side 3 ends squares of sides 1, 2 and 3, so add the side itself. Example 1 has side values summing to 15.", CodeSnippet: "total += side[i][j]", LineStart: 9, LineEnd: 9},
				{Title: "Return the total", Explanation: "Every square is counted once, at its bottom-right corner.", CodeSnippet: "return total", LineStart: 10, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{Name: "Top-down with a memo", Code: "from functools import lru_cache\n\n\ndef countSquares(matrix):\n    @lru_cache(maxsize=None)\n    def side(i, j):\n        if i < 0 or j < 0 or matrix[i][j] == 0:\n            return 0\n        return 1 + min(side(i - 1, j), side(i, j - 1), side(i - 1, j - 1))\n\n    return sum(side(i, j) for i in range(len(matrix)) for j in range(len(matrix[0])))", TimeComplexity: "O(rows * cols)", SpaceComplexity: "O(rows * cols)", Explanation: "The same recurrence as a recursive function; lru_cache stores each corner's side so it is computed once. It is a good check of the state definition, but the recursion can reach depth rows + cols."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does adding side[i][j] for every corner count each square exactly once?", Answer: "A square is identified by its bottom-right corner and its side, so it has exactly one corner and one side. If the largest square at a corner has side k, then all smaller sides 1 to k at that corner are also all ones (they sit inside it), and no bigger side is. So the corner holds exactly k squares and across all corners nothing repeats."},
			{Prompt: "What changes if you only want the side of the largest square (Maximal Square)?", Answer: "Nothing in the table: the same side[i][j] with the same recurrence. Instead of summing the table, return the maximum entry (squared if the area is wanted). The answer-can-end-anywhere habit is the same as in the repeated-subarray drill."},
		},
	},
}

func tddEmptyGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	return grid
}

func tddPathsObstacles(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	paths := tddEmptyGrid(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				continue
			}
			if i == 0 && j == 0 {
				paths[i][j] = 1
				continue
			}
			if i > 0 {
				paths[i][j] += paths[i-1][j]
			}
			if j > 0 {
				paths[i][j] += paths[i][j-1]
			}
		}
	}
	return paths[rows-1][cols-1]
}

// tddSeq is a deterministic pseudo-random sequence.
func tddSeq(seed int) func() int {
	return func() int {
		seed = (seed*1103515245 + 12345) % 2147483648
		return seed >> 8
	}
}

// tddCostGrid is a grid of costs from 0 to 9.
func tddCostGrid(rows, cols, seed int) [][]int {
	next := tddSeq(seed)
	grid := tddEmptyGrid(rows, cols)
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = next() % 10
		}
	}
	return grid
}

func tddMinPath(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	cost := tddEmptyGrid(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			before := 0
			switch {
			case i == 0 && j == 0:
			case i == 0:
				before = cost[i][j-1]
			case j == 0:
				before = cost[i-1][j]
			default:
				before = cost[i-1][j]
				if cost[i][j-1] < before {
					before = cost[i][j-1]
				}
			}
			cost[i][j] = grid[i][j] + before
		}
	}
	return cost[rows-1][cols-1]
}

// tddTriangle has numbers from -50 to 49 in row r with r + 1 entries.
func tddTriangle(rows int) [][]int {
	next := tddSeq(31)
	tri := make([][]int, rows)
	for r := range tri {
		tri[r] = make([]int, r+1)
		for c := range tri[r] {
			tri[r][c] = next()%100 - 50
		}
	}
	return tri
}

func tddMinTriangle(tri [][]int) int {
	best := append([]int(nil), tri[len(tri)-1]...)
	for r := len(tri) - 2; r >= 0; r-- {
		for c := 0; c <= r; c++ {
			m := best[c]
			if best[c+1] < m {
				m = best[c+1]
			}
			best[c] = tri[r][c] + m
		}
	}
	return best[0]
}

func tddRepeat(n, v int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = v
	}
	return out
}

// tddSmall is n numbers from 0 to k - 1.
func tddSmall(n, k, seed int) []int {
	next := tddSeq(seed)
	out := make([]int, n)
	for i := range out {
		out[i] = next() % k
	}
	return out
}

func tddRepeatedLen(a, b []int) int {
	run := make([][]int, len(a)+1)
	for i := range run {
		run[i] = make([]int, len(b)+1)
	}
	longest := 0
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				run[i][j] = run[i-1][j-1] + 1
				if run[i][j] > longest {
					longest = run[i][j]
				}
			}
		}
	}
	return longest
}

func tddKnapWeights(n int) []int {
	next := tddSeq(77)
	out := make([]int, n)
	for i := range out {
		out[i] = next()%60 + 1
	}
	return out
}

func tddKnapValues(n int) []int {
	next := tddSeq(78)
	out := make([]int, n)
	for i := range out {
		out[i] = next() % 500
	}
	return out
}

func tddKnapsack(weights, values []int, capacity int) int {
	best := make([]int, capacity+1)
	for i, w := range weights {
		for room := capacity; room >= w; room-- {
			if v := best[room-w] + values[i]; v > best[room] {
				best[room] = v
			}
		}
	}
	return best[capacity]
}

func tddOnes(rows, cols int) [][]int {
	grid := tddEmptyGrid(rows, cols)
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = 1
		}
	}
	return grid
}

// tddRandomBits is a grid where about one cell in ten is 0.
func tddRandomBits(rows, cols, seed int) [][]int {
	next := tddSeq(seed)
	grid := tddEmptyGrid(rows, cols)
	for i := range grid {
		for j := range grid[i] {
			if next()%10 != 0 {
				grid[i][j] = 1
			}
		}
	}
	return grid
}

func tddCountSquares(m [][]int) int {
	rows, cols := len(m), len(m[0])
	side := tddEmptyGrid(rows+1, cols+1)
	total := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if m[i-1][j-1] == 1 {
				s := side[i-1][j]
				if side[i][j-1] < s {
					s = side[i][j-1]
				}
				if side[i-1][j-1] < s {
					s = side[i-1][j-1]
				}
				side[i][j] = s + 1
				total += s + 1
			}
		}
	}
	return total
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, TwoDDPDrills...)
}
