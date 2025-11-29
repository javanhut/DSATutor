package practice

// GraphsProblems contains all graphs category problems
var GraphsProblems = []*Problem{
	{
		ID:         "number-of-islands",
		Number:     32,
		Title:      "Number of Islands",
		Difficulty: "Medium",
		Category:   "graphs",
		Tags:       []string{"Graph", "BFS", "DFS", "Matrix"},
		RelatedChapters: []int{11},
		Description: `Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.`,
		Constraints: []string{
			"m == grid.length",
			"n == grid[i].length",
			"1 <= m, n <= 300",
			"grid[i][j] is '0' or '1'",
		},
		Examples: []Example{
			{Input: `grid = [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]`, Output: "1"},
			{Input: `grid = [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]`, Output: "3"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"grid": [][]string{{"1", "1", "0"}, {"0", "1", "0"}, {"0", "0", "1"}}}, Expected: 2},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def numIslands(grid):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use DFS/BFS to explore and mark each island."},
			{Level: 2, Type: "algorithm", Content: "For each '1' found, increment count and flood-fill to mark entire island."},
		},
		Solution: Solution{
			Code: `def numIslands(grid):
    if not grid:
        return 0

    rows, cols = len(grid), len(grid[0])
    count = 0

    def dfs(r, c):
        if r < 0 or r >= rows or c < 0 or c >= cols or grid[r][c] != '1':
            return
        grid[r][c] = '0'  # Mark as visited
        dfs(r+1, c)
        dfs(r-1, c)
        dfs(r, c+1)
        dfs(r, c-1)

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '1':
                count += 1
                dfs(r, c)

    return count`,
			Explanation: "DFS from each unvisited land cell, marking visited cells. Count DFS starts.",
		},
	},
	{
		ID:         "course-schedule",
		Number:     35,
		Title:      "Course Schedule",
		Difficulty: "Medium",
		Category:   "graphs",
		Tags:       []string{"DFS", "BFS", "Graph", "Topological Sort"},
		RelatedChapters: []int{11},
		Description: `There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

Return true if you can finish all courses. Otherwise, return false.`,
		Constraints: []string{
			"1 <= numCourses <= 2000",
			"0 <= prerequisites.length <= 5000",
			"prerequisites[i].length == 2",
			"0 <= ai, bi < numCourses",
			"All the pairs prerequisites[i] are unique",
		},
		Examples: []Example{
			{Input: "numCourses = 2, prerequisites = [[1,0]]", Output: "true", Explanation: "Take course 0, then course 1."},
			{Input: "numCourses = 2, prerequisites = [[1,0],[0,1]]", Output: "false", Explanation: "Circular dependency."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"numCourses": 2, "prerequisites": [][]int{{1, 0}}}, Expected: true},
			{Input: map[string]interface{}{"numCourses": 2, "prerequisites": [][]int{{1, 0}, {0, 1}}}, Expected: false},
		},
		TimeComplexity:  "O(V + E)",
		SpaceComplexity: "O(V + E)",
		StarterCode:     "def canFinish(numCourses, prerequisites):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "This is cycle detection in a directed graph."},
			{Level: 2, Type: "algorithm", Content: "Use DFS with three states: unvisited, visiting, visited. Cycle if we hit 'visiting'."},
		},
		Solution: Solution{
			Code: `def canFinish(numCourses, prerequisites):
    graph = [[] for _ in range(numCourses)]
    for course, prereq in prerequisites:
        graph[course].append(prereq)

    # 0 = unvisited, 1 = visiting, 2 = visited
    state = [0] * numCourses

    def has_cycle(course):
        if state[course] == 1:
            return True
        if state[course] == 2:
            return False

        state[course] = 1
        for prereq in graph[course]:
            if has_cycle(prereq):
                return True
        state[course] = 2
        return False

    for course in range(numCourses):
        if has_cycle(course):
            return False
    return True`,
			Explanation: "DFS cycle detection. Three states track if node is currently being processed.",
		},
	},
	{
		ID:              "clone-graph",
		Number:          79,
		Title:           "Clone Graph",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Hash Table", "DFS", "BFS", "Graph"},
		RelatedChapters: []int{6, 10},
		Description: `Given a reference of a node in a connected undirected graph, return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.`,
		Constraints: []string{
			"The number of nodes in the graph is in the range [0, 100]",
			"1 <= Node.val <= 100",
			"Node.val is unique for each node",
			"There are no repeated edges and no self-loops in the graph",
			"The Graph is connected and all nodes can be visited starting from the given node",
		},
		Examples: []Example{
			{Input: "adjList = [[2,4],[1,3],[2,4],[1,3]]", Output: "[[2,4],[1,3],[2,4],[1,3]]"},
			{Input: "adjList = [[]]", Output: "[[]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"adjList": [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}}, Expected: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}},
		},
		TimeComplexity:  "O(V + E)",
		SpaceComplexity: "O(V)",
		StarterCode:     "def cloneGraph(node):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use hash map to track original->clone mapping. DFS/BFS to traverse. Create clones and connect neighbors.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "DFS: if node already cloned, return clone. Create new node, add to map, recursively clone neighbors."},
			{Level: 3, Type: "code", Content: "def dfs(node): if node in old_to_new: return it. Create clone, add to map. For each neighbor, append dfs(neighbor) to clone.neighbors."},
		},
		Solution: Solution{
			Code: `def cloneGraph(node):
    if not node:
        return None

    old_to_new = {}

    def dfs(node):
        if node in old_to_new:
            return old_to_new[node]

        clone = Node(node.val)
        old_to_new[node] = clone

        for neighbor in node.neighbors:
            clone.neighbors.append(dfs(neighbor))

        return clone

    return dfs(node)`,
			Explanation:     "DFS with memoization. Map original nodes to clones. Recursively clone neighbors.",
			TimeComplexity:  "O(V + E)",
			SpaceComplexity: "O(V) for hash map",
			Walkthrough: []WalkthroughStep{
				{Title: "Check if already cloned", Explanation: "Return existing clone to avoid infinite loop", CodeSnippet: "if node in old_to_new:\n    return old_to_new[node]", LineStart: 8, LineEnd: 9},
				{Title: "Create clone first", Explanation: "Add to map before recursing to handle cycles", CodeSnippet: "clone = Node(node.val)\nold_to_new[node] = clone", LineStart: 11, LineEnd: 12},
				{Title: "Clone neighbors", Explanation: "Recursively clone and connect", CodeSnippet: "clone.neighbors.append(dfs(neighbor))", LineStart: 15, LineEnd: 15},
			},
		},
	},
	{
		ID:              "pacific-atlantic-water-flow",
		Number:          80,
		Title:           "Pacific Atlantic Water Flow",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Array", "DFS", "BFS", "Matrix"},
		RelatedChapters: []int{6, 10},
		Description: `There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.`,
		Constraints: []string{
			"m == heights.length",
			"n == heights[r].length",
			"1 <= m, n <= 200",
			"0 <= heights[r][c] <= 10^5",
		},
		Examples: []Example{
			{Input: "heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]", Output: "[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"heights": [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}}, Expected: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def pacificAtlantic(heights):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Reverse the flow: DFS from ocean borders going uphill. Find cells reachable from both oceans.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "DFS from Pacific borders (top, left) and Atlantic borders (bottom, right). Return intersection of reachable cells."},
			{Level: 3, Type: "code", Content: "pacific, atlantic = set(), set(). DFS from borders with condition heights[nr][nc] >= heights[r][c]. Return intersection."},
		},
		Solution: Solution{
			Code: `def pacificAtlantic(heights):
    if not heights:
        return []

    m, n = len(heights), len(heights[0])
    pacific = set()
    atlantic = set()

    def dfs(r, c, reachable, prev_height):
        if (r, c) in reachable or r < 0 or r >= m or c < 0 or c >= n:
            return
        if heights[r][c] < prev_height:
            return

        reachable.add((r, c))
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            dfs(r + dr, c + dc, reachable, heights[r][c])

    # DFS from Pacific (top and left borders)
    for c in range(n):
        dfs(0, c, pacific, 0)
    for r in range(m):
        dfs(r, 0, pacific, 0)

    # DFS from Atlantic (bottom and right borders)
    for c in range(n):
        dfs(m - 1, c, atlantic, 0)
    for r in range(m):
        dfs(r, n - 1, atlantic, 0)

    return list(pacific & atlantic)`,
			Explanation:     "Reverse flow: DFS uphill from ocean borders. Cells in both sets can reach both oceans.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Reverse flow", Explanation: "Go uphill: heights[nr][nc] >= heights[r][c]", CodeSnippet: "if heights[r][c] < prev_height:\n    return", LineStart: 11, LineEnd: 12},
				{Title: "DFS from Pacific", Explanation: "Top row and left column", CodeSnippet: "for c in range(n):\n    dfs(0, c, pacific, 0)", LineStart: 20, LineEnd: 23},
				{Title: "Intersection", Explanation: "Cells reachable from both oceans", CodeSnippet: "return list(pacific & atlantic)", LineStart: 31, LineEnd: 31},
			},
		},
	},
	{
		ID:              "rotting-oranges",
		Number:          81,
		Title:           "Rotting Oranges",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Array", "BFS", "Matrix"},
		RelatedChapters: []int{6, 10},
		Description: `You are given an m x n grid where each cell can have one of three values:
- 0 representing an empty cell,
- 1 representing a fresh orange, or
- 2 representing a rotten orange.

Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.`,
		Constraints: []string{
			"m == grid.length",
			"n == grid[i].length",
			"1 <= m, n <= 10",
			"grid[i][j] is 0, 1, or 2",
		},
		Examples: []Example{
			{Input: "grid = [[2,1,1],[1,1,0],[0,1,1]]", Output: "4"},
			{Input: "grid = [[2,1,1],[0,1,1],[1,0,1]]", Output: "-1", Explanation: "Bottom left orange can never rot."},
			{Input: "grid = [[0,2]]", Output: "0"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"grid": [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}}, Expected: 4},
			{Input: map[string]interface{}{"grid": [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}}, Expected: -1},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def orangesRotting(grid):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Multi-source BFS from all rotten oranges simultaneously. Count fresh oranges. Track minutes.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Add all rotten to queue, count fresh. BFS level by level (each level = 1 minute). Decrement fresh count."},
			{Level: 3, Type: "code", Content: "queue = all rotten. while queue: process level, rot adjacent fresh. if fresh > 0: return -1. return minutes."},
		},
		Solution: Solution{
			Code: `from collections import deque

def orangesRotting(grid):
    m, n = len(grid), len(grid[0])
    queue = deque()
    fresh = 0

    # Find all rotten and count fresh
    for r in range(m):
        for c in range(n):
            if grid[r][c] == 2:
                queue.append((r, c))
            elif grid[r][c] == 1:
                fresh += 1

    if fresh == 0:
        return 0

    minutes = 0
    while queue:
        minutes += 1
        for _ in range(len(queue)):
            r, c = queue.popleft()
            for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                nr, nc = r + dr, c + dc
                if 0 <= nr < m and 0 <= nc < n and grid[nr][nc] == 1:
                    grid[nr][nc] = 2
                    fresh -= 1
                    queue.append((nr, nc))

    return minutes - 1 if fresh == 0 else -1`,
			Explanation:     "Multi-source BFS from all rotten oranges. Each BFS level is one minute. Check if all fresh oranges rotted.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n) for queue",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize", Explanation: "Queue all rotten, count fresh", CodeSnippet: "if grid[r][c] == 2:\n    queue.append((r, c))\nelif grid[r][c] == 1:\n    fresh += 1", LineStart: 11, LineEnd: 14},
				{Title: "BFS by level", Explanation: "Each level is one minute", CodeSnippet: "for _ in range(len(queue)):", LineStart: 22, LineEnd: 22},
				{Title: "Rot adjacent", Explanation: "Mark as rotten and add to queue", CodeSnippet: "grid[nr][nc] = 2\nfresh -= 1\nqueue.append((nr, nc))", LineStart: 27, LineEnd: 29},
			},
		},
	},
	{
		ID:              "course-schedule-ii",
		Number:          82,
		Title:           "Course Schedule II",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"DFS", "BFS", "Graph", "Topological Sort"},
		RelatedChapters: []int{6, 10},
		Description: `There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.`,
		Constraints: []string{
			"1 <= numCourses <= 2000",
			"0 <= prerequisites.length <= numCourses * (numCourses - 1)",
			"prerequisites[i].length == 2",
			"0 <= ai, bi < numCourses",
			"ai != bi",
			"All the pairs [ai, bi] are distinct",
		},
		Examples: []Example{
			{Input: "numCourses = 2, prerequisites = [[1,0]]", Output: "[0,1]"},
			{Input: "numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]", Output: "[0,2,1,3]"},
			{Input: "numCourses = 1, prerequisites = []", Output: "[0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"numCourses": 2, "prerequisites": [][]int{{1, 0}}}, Expected: []int{0, 1}},
			{Input: map[string]interface{}{"numCourses": 4, "prerequisites": [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, Expected: []int{0, 1, 2, 3}},
		},
		TimeComplexity:  "O(V + E)",
		SpaceComplexity: "O(V + E)",
		StarterCode:     "def findOrder(numCourses, prerequisites):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Topological sort using Kahn's algorithm (BFS) or DFS. Return empty array if cycle detected.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Kahn's: Track in-degrees. Add courses with in-degree 0 to queue. Process, reduce in-degrees, add new zeros."},
			{Level: 3, Type: "code", Content: "Build graph and in-degree array. Queue courses with in-degree 0. While queue: pop, add to result, decrement neighbors' in-degrees."},
		},
		Solution: Solution{
			Code: `from collections import deque, defaultdict

def findOrder(numCourses, prerequisites):
    graph = defaultdict(list)
    in_degree = [0] * numCourses

    # Build graph
    for course, prereq in prerequisites:
        graph[prereq].append(course)
        in_degree[course] += 1

    # Start with courses having no prerequisites
    queue = deque([i for i in range(numCourses) if in_degree[i] == 0])
    result = []

    while queue:
        course = queue.popleft()
        result.append(course)

        for next_course in graph[course]:
            in_degree[next_course] -= 1
            if in_degree[next_course] == 0:
                queue.append(next_course)

    return result if len(result) == numCourses else []`,
			Explanation:     "Kahn's algorithm for topological sort. Track in-degrees. Process nodes with in-degree 0. Return empty if cycle.",
			TimeComplexity:  "O(V + E)",
			SpaceComplexity: "O(V + E)",
			Walkthrough: []WalkthroughStep{
				{Title: "Build graph", Explanation: "Adjacency list and in-degree count", CodeSnippet: "graph[prereq].append(course)\nin_degree[course] += 1", LineStart: 9, LineEnd: 10},
				{Title: "Start with zeros", Explanation: "Courses with no prerequisites", CodeSnippet: "queue = deque([i for i in range(numCourses) if in_degree[i] == 0])", LineStart: 13, LineEnd: 13},
				{Title: "Process and update", Explanation: "Add to result, decrement in-degrees of neighbors", CodeSnippet: "in_degree[next_course] -= 1\nif in_degree[next_course] == 0:", LineStart: 21, LineEnd: 23},
			},
		},
	},
	{
		ID:              "surrounded-regions",
		Number:          111,
		Title:           "Surrounded Regions",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Array", "DFS", "BFS", "Union Find", "Matrix"},
		RelatedChapters: []int{6, 10},
		Description: `Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.`,
		Constraints: []string{
			"m == board.length",
			"n == board[i].length",
			"1 <= m, n <= 200",
			"board[i][j] is 'X' or 'O'",
		},
		Examples: []Example{
			{Input: `board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]`, Output: `[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"board": [][]string{{"X", "X", "X", "X"}, {"X", "O", "O", "X"}, {"X", "X", "O", "X"}, {"X", "O", "X", "X"}}}, Expected: [][]string{{"X", "X", "X", "X"}, {"X", "X", "X", "X"}, {"X", "X", "X", "X"}, {"X", "O", "X", "X"}}},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def solve(board):\n    # Write your solution here (modify in place)\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "O's on border can't be captured. DFS from border O's to mark safe cells. Flip remaining O's.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Mark border-connected O's as 'S' (safe). Then flip all 'O' to 'X', all 'S' to 'O'."},
			{Level: 3, Type: "code", Content: "DFS from border O's, mark as 'S'. Loop through board: O->X, S->O."},
		},
		Solution: Solution{
			Code: `def solve(board):
    if not board:
        return

    m, n = len(board), len(board[0])

    def dfs(r, c):
        if r < 0 or r >= m or c < 0 or c >= n or board[r][c] != 'O':
            return
        board[r][c] = 'S'  # Mark as safe
        dfs(r + 1, c)
        dfs(r - 1, c)
        dfs(r, c + 1)
        dfs(r, c - 1)

    # Mark border-connected O's as safe
    for r in range(m):
        dfs(r, 0)
        dfs(r, n - 1)
    for c in range(n):
        dfs(0, c)
        dfs(m - 1, c)

    # Flip: O -> X (surrounded), S -> O (safe)
    for r in range(m):
        for c in range(n):
            if board[r][c] == 'O':
                board[r][c] = 'X'
            elif board[r][c] == 'S':
                board[r][c] = 'O'`,
			Explanation:     "Mark border-connected O's as safe. Then flip remaining O's to X, restore safe cells to O.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n) for recursion stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Mark safe", Explanation: "DFS from border, mark connected O's", CodeSnippet: "board[r][c] = 'S'", LineStart: 10, LineEnd: 10},
				{Title: "Process borders", Explanation: "DFS from all border cells", CodeSnippet: "dfs(r, 0)\ndfs(r, n - 1)", LineStart: 18, LineEnd: 19},
				{Title: "Final flip", Explanation: "O->X (captured), S->O (safe)", CodeSnippet: "if board[r][c] == 'O':\n    board[r][c] = 'X'", LineStart: 27, LineEnd: 30},
			},
		},
	},
	{
		ID:              "word-ladder",
		Number:          112,
		Title:           "Word Ladder",
		Difficulty:      "Hard",
		Category:        "graphs",
		Tags:            []string{"Hash Table", "String", "BFS"},
		RelatedChapters: []int{5, 6},
		Description: `A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

- Every adjacent pair of words differs by a single letter.
- Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
- sk == endWord

Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.`,
		Constraints: []string{
			"1 <= beginWord.length <= 10",
			"endWord.length == beginWord.length",
			"1 <= wordList.length <= 5000",
			"wordList[i].length == beginWord.length",
			"beginWord, endWord, and wordList[i] consist of lowercase English letters",
			"beginWord != endWord",
			"All the words in wordList are unique",
		},
		Examples: []Example{
			{Input: `beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]`, Output: "5", Explanation: "hit -> hot -> dot -> dog -> cog"},
			{Input: `beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]`, Output: "0"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"beginWord": "hit", "endWord": "cog", "wordList": []string{"hot", "dot", "dog", "lot", "log", "cog"}}, Expected: 5},
		},
		TimeComplexity:  "O(M^2 * N)",
		SpaceComplexity: "O(M^2 * N)",
		StarterCode:     "def ladderLength(beginWord, endWord, wordList):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "BFS finds shortest path. Build graph where edges connect words differing by one letter.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Use pattern matching: 'hot' -> '*ot', 'h*t', 'ho*'. Words with same pattern are neighbors."},
			{Level: 3, Type: "code", Content: "Build pattern -> words map. BFS from beginWord. For each word, try all patterns to find neighbors."},
		},
		Solution: Solution{
			Code: `from collections import defaultdict, deque

def ladderLength(beginWord, endWord, wordList):
    if endWord not in wordList:
        return 0

    word_set = set(wordList)
    queue = deque([(beginWord, 1)])
    visited = {beginWord}

    while queue:
        word, length = queue.popleft()

        if word == endWord:
            return length

        for i in range(len(word)):
            for c in 'abcdefghijklmnopqrstuvwxyz':
                next_word = word[:i] + c + word[i+1:]
                if next_word in word_set and next_word not in visited:
                    visited.add(next_word)
                    queue.append((next_word, length + 1))

    return 0`,
			Explanation:     "BFS from beginWord. Try changing each position to each letter. Track visited to avoid cycles.",
			TimeComplexity:  "O(M^2 * N) where M is word length, N is word count",
			SpaceComplexity: "O(M * N)",
			Walkthrough: []WalkthroughStep{
				{Title: "BFS for shortest path", Explanation: "Queue stores (word, path_length)", CodeSnippet: "queue = deque([(beginWord, 1)])", LineStart: 8, LineEnd: 8},
				{Title: "Generate neighbors", Explanation: "Try all single-letter changes", CodeSnippet: "next_word = word[:i] + c + word[i+1:]", LineStart: 19, LineEnd: 19},
				{Title: "Check valid neighbor", Explanation: "Must be in dictionary and not visited", CodeSnippet: "if next_word in word_set and next_word not in visited:", LineStart: 20, LineEnd: 22},
			},
		},
	},
	{
		ID:              "reconstruct-itinerary",
		Number:          122,
		Title:           "Reconstruct Itinerary",
		Difficulty:      "Hard",
		Category:        "graphs",
		Tags:            []string{"Graph", "DFS", "Eulerian Path"},
		RelatedChapters: []int{6, 12},
		Description: `You are given a list of airline tickets where tickets[i] = [fromi, toi] represent the departure and the arrival airports of one flight. Reconstruct the itinerary in order and return it.

All of the tickets belong to a man who departs from "JFK", thus, the itinerary must begin with "JFK". If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string.

You may assume all tickets form at least one valid itinerary. You must use all the tickets once and only once.`,
		Constraints: []string{
			"1 <= tickets.length <= 300",
			"tickets[i].length == 2",
			"fromi.length == 3",
			"toi.length == 3",
			"fromi and toi consist of uppercase English letters",
			"fromi != toi",
		},
		Examples: []Example{
			{Input: `tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]`, Output: `["JFK","MUC","LHR","SFO","SJC"]`},
			{Input: `tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]`, Output: `["JFK","ATL","JFK","SFO","ATL","SFO"]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"tickets": [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}}, Expected: []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
			{Input: map[string]interface{}{"tickets": [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}}, Expected: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
		},
		TimeComplexity:  "O(E log E)",
		SpaceComplexity: "O(E)",
		StarterCode:     "def findItinerary(tickets):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "This is Eulerian path problem. Use Hierholzer's algorithm with DFS.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Build adjacency list sorted in reverse. DFS from JFK, pop destinations, add to result in reverse."},
			{Level: 3, Type: "code", Content: "graph = defaultdict(list). Sort destinations reversed. DFS: while graph[node]: dfs(graph[node].pop()). result.append(node). Return reversed."},
		},
		Solution: Solution{
			Code: `def findItinerary(tickets):
    from collections import defaultdict

    graph = defaultdict(list)
    for src, dst in sorted(tickets, reverse=True):
        graph[src].append(dst)

    result = []

    def dfs(node):
        while graph[node]:
            dfs(graph[node].pop())
        result.append(node)

    dfs("JFK")
    return result[::-1]`,
			Explanation:     "Hierholzer's algorithm for Eulerian path. Sort in reverse for lexical order, build path backwards.",
			TimeComplexity:  "O(E log E)",
			SpaceComplexity: "O(E)",
			Walkthrough: []WalkthroughStep{
				{Title: "Sort reverse", Explanation: "Reverse sort so pop gives smallest", CodeSnippet: "for src, dst in sorted(tickets, reverse=True):", LineStart: 5, LineEnd: 5},
				{Title: "DFS with pop", Explanation: "Visit and remove edges", CodeSnippet: "while graph[node]:\n    dfs(graph[node].pop())", LineStart: 11, LineEnd: 12},
				{Title: "Build backwards", Explanation: "Append after visiting all neighbors", CodeSnippet: "result.append(node)", LineStart: 13, LineEnd: 13},
			},
		},
	},
	{
		ID:              "min-cost-to-connect-all-points",
		Number:          123,
		Title:           "Min Cost to Connect All Points",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "MST", "Prim's", "Union Find"},
		RelatedChapters: []int{6, 7, 12},
		Description: `You are given an array points representing integer coordinates of some points on a 2D-plane, where points[i] = [xi, yi].

The cost of connecting two points [xi, yi] and [xj, yj] is the manhattan distance between them: |xi - xj| + |yi - yj|.

Return the minimum cost to make all points connected. All points are connected if there is exactly one simple path between any two points.`,
		Constraints: []string{
			"1 <= points.length <= 1000",
			"-10^6 <= xi, yi <= 10^6",
			"All pairs (xi, yi) are distinct",
		},
		Examples: []Example{
			{Input: "points = [[0,0],[2,2],[3,10],[5,2],[7,0]]", Output: "20"},
			{Input: "points = [[3,12],[-2,5],[-4,1]]", Output: "18"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"points": [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}}, Expected: 20},
			{Input: map[string]interface{}{"points": [][]int{{3, 12}, {-2, 5}, {-4, 1}}}, Expected: 18},
		},
		TimeComplexity:  "O(n^2 log n)",
		SpaceComplexity: "O(n^2)",
		StarterCode:     "def minCostConnectPoints(points):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "This is Minimum Spanning Tree (MST). Use Prim's or Kruskal's algorithm.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Prim's: Start from any node, greedily add closest unvisited node. Use min-heap."},
			{Level: 3, Type: "code", Content: "heap = [(0, 0)]. visited = set(). while len(visited) < n: cost, node = heappop. if not visited: add to visited, add neighbors."},
		},
		Solution: Solution{
			Code: `def minCostConnectPoints(points):
    import heapq

    n = len(points)
    visited = set()
    heap = [(0, 0)]  # (cost, point_index)
    total_cost = 0

    while len(visited) < n:
        cost, i = heapq.heappop(heap)
        if i in visited:
            continue
        visited.add(i)
        total_cost += cost

        for j in range(n):
            if j not in visited:
                dist = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])
                heapq.heappush(heap, (dist, j))

    return total_cost`,
			Explanation:     "Prim's algorithm: greedily add closest unconnected point until all connected.",
			TimeComplexity:  "O(n^2 log n)",
			SpaceComplexity: "O(n^2)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize heap", Explanation: "Start from point 0 with cost 0", CodeSnippet: "heap = [(0, 0)]", LineStart: 6, LineEnd: 6},
				{Title: "Greedy selection", Explanation: "Pop minimum cost unvisited point", CodeSnippet: "cost, i = heapq.heappop(heap)\nif i in visited:\n    continue", LineStart: 10, LineEnd: 12},
				{Title: "Add neighbors", Explanation: "Push distances to all unvisited points", CodeSnippet: "dist = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])", LineStart: 18, LineEnd: 18},
			},
		},
	},
	{
		ID:              "network-delay-time",
		Number:          124,
		Title:           "Network Delay Time",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Dijkstra", "Shortest Path"},
		RelatedChapters: []int{7, 12},
		Description: `You are given a network of n nodes, labeled from 1 to n. You are also given times, a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node, vi is the target node, and wi is the time it takes for a signal to travel from source to target.

We will send a signal from a given node k. Return the minimum time it takes for all the n nodes to receive the signal. If it is impossible for all the n nodes to receive the signal, return -1.`,
		Constraints: []string{
			"1 <= k <= n <= 100",
			"1 <= times.length <= 6000",
			"times[i].length == 3",
			"1 <= ui, vi <= n",
			"ui != vi",
			"0 <= wi <= 100",
			"All the pairs (ui, vi) are unique",
		},
		Examples: []Example{
			{Input: "times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2", Output: "2"},
			{Input: "times = [[1,2,1]], n = 2, k = 1", Output: "1"},
			{Input: "times = [[1,2,1]], n = 2, k = 2", Output: "-1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"times": [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, "n": 4, "k": 2}, Expected: 2},
			{Input: map[string]interface{}{"times": [][]int{{1, 2, 1}}, "n": 2, "k": 1}, Expected: 1},
			{Input: map[string]interface{}{"times": [][]int{{1, 2, 1}}, "n": 2, "k": 2}, Expected: -1},
		},
		TimeComplexity:  "O(E log V)",
		SpaceComplexity: "O(V + E)",
		StarterCode:     "def networkDelayTime(times, n, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use Dijkstra's algorithm to find shortest path from k to all nodes.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Build graph, use min-heap starting from k. Track minimum time to reach each node."},
			{Level: 3, Type: "code", Content: "heap = [(0, k)]. dist = {}. while heap: if node not in dist: dist[node] = time; add neighbors. return max(dist) if len == n else -1."},
		},
		Solution: Solution{
			Code: `def networkDelayTime(times, n, k):
    import heapq
    from collections import defaultdict

    graph = defaultdict(list)
    for u, v, w in times:
        graph[u].append((v, w))

    heap = [(0, k)]
    dist = {}

    while heap:
        time, node = heapq.heappop(heap)
        if node in dist:
            continue
        dist[node] = time

        for neighbor, weight in graph[node]:
            if neighbor not in dist:
                heapq.heappush(heap, (time + weight, neighbor))

    return max(dist.values()) if len(dist) == n else -1`,
			Explanation:     "Dijkstra finds shortest path to all nodes. Answer is max of all shortest paths.",
			TimeComplexity:  "O(E log V)",
			SpaceComplexity: "O(V + E)",
			Walkthrough: []WalkthroughStep{
				{Title: "Build graph", Explanation: "Adjacency list with weights", CodeSnippet: "for u, v, w in times:\n    graph[u].append((v, w))", LineStart: 6, LineEnd: 7},
				{Title: "Dijkstra", Explanation: "Pop minimum, record if first visit", CodeSnippet: "if node in dist:\n    continue\ndist[node] = time", LineStart: 14, LineEnd: 16},
				{Title: "Check reachability", Explanation: "All n nodes must be reached", CodeSnippet: "return max(dist.values()) if len(dist) == n else -1", LineStart: 22, LineEnd: 22},
			},
		},
	},
	{
		ID:              "swim-in-rising-water",
		Number:          125,
		Title:           "Swim in Rising Water",
		Difficulty:      "Hard",
		Category:        "graphs",
		Tags:            []string{"Graph", "Binary Search", "BFS", "Union Find"},
		RelatedChapters: []int{1, 6, 7, 12},
		Description: `You are given an n x n integer matrix grid where each value grid[i][j] represents the elevation at that point (i, j).

The rain starts to fall. At time t, the depth of the water everywhere is t. You can swim from a square to another 4-directionally adjacent square if and only if the elevation of both squares individually are at most t. You can swim infinite distances in zero time. Of course, you must stay within the boundaries of the grid during your swim.

Return the least time until you can reach the bottom right square (n - 1, n - 1) if you start at the top left square (0, 0).`,
		Constraints: []string{
			"n == grid.length",
			"n == grid[i].length",
			"1 <= n <= 50",
			"0 <= grid[i][j] < n^2",
			"Each value grid[i][j] is unique",
		},
		Examples: []Example{
			{Input: "grid = [[0,2],[1,3]]", Output: "3"},
			{Input: "grid = [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]", Output: "16"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"grid": [][]int{{0, 2}, {1, 3}}}, Expected: 3},
			{Input: map[string]interface{}{"grid": [][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}}, Expected: 16},
		},
		TimeComplexity:  "O(n^2 log n)",
		SpaceComplexity: "O(n^2)",
		StarterCode:     "def swimInWater(grid):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use modified Dijkstra or binary search + BFS. Find minimum time to reach destination.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Dijkstra variant: cost to reach cell is max(current_max, cell_value). Use min-heap."},
			{Level: 3, Type: "code", Content: "heap = [(grid[0][0], 0, 0)]. while heap: t, r, c = pop. if (r,c) == (n-1,n-1): return t. For neighbors: push (max(t, grid[nr][nc]), nr, nc)."},
		},
		Solution: Solution{
			Code: `def swimInWater(grid):
    import heapq

    n = len(grid)
    visited = set()
    heap = [(grid[0][0], 0, 0)]

    while heap:
        t, r, c = heapq.heappop(heap)
        if (r, c) in visited:
            continue
        visited.add((r, c))

        if r == n - 1 and c == n - 1:
            return t

        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nr, nc = r + dr, c + dc
            if 0 <= nr < n and 0 <= nc < n and (nr, nc) not in visited:
                heapq.heappush(heap, (max(t, grid[nr][nc]), nr, nc))

    return -1`,
			Explanation:     "Modified Dijkstra where edge weight is max elevation seen. First path to end is optimal.",
			TimeComplexity:  "O(n^2 log n)",
			SpaceComplexity: "O(n^2)",
			Walkthrough: []WalkthroughStep{
				{Title: "Track max elevation", Explanation: "Time needed is max elevation on path", CodeSnippet: "(max(t, grid[nr][nc]), nr, nc)", LineStart: 20, LineEnd: 20},
				{Title: "First arrival wins", Explanation: "Min-heap ensures optimal path found first", CodeSnippet: "if r == n - 1 and c == n - 1:\n    return t", LineStart: 14, LineEnd: 15},
			},
		},
	},
	{
		ID:              "cheapest-flights-within-k-stops",
		Number:          126,
		Title:           "Cheapest Flights Within K Stops",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "BFS", "Dijkstra", "Dynamic Programming"},
		RelatedChapters: []int{7, 9, 12},
		Description: `There are n cities connected by some number of flights. You are given an array flights where flights[i] = [fromi, toi, pricei] indicates that there is a flight from city fromi to city toi with cost pricei.

You are also given three integers src, dst, and k, return the cheapest price from src to dst with at most k stops. If there is no such route, return -1.`,
		Constraints: []string{
			"1 <= n <= 100",
			"0 <= flights.length <= (n * (n - 1) / 2)",
			"flights[i].length == 3",
			"0 <= fromi, toi < n",
			"fromi != toi",
			"1 <= pricei <= 10^4",
			"There will not be any multiple flights between two cities",
			"0 <= src, dst, k < n",
			"src != dst",
		},
		Examples: []Example{
			{Input: "n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1", Output: "700"},
			{Input: "n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1", Output: "200"},
			{Input: "n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0", Output: "500"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 4, "flights": [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, "src": 0, "dst": 3, "k": 1}, Expected: 700},
			{Input: map[string]interface{}{"n": 3, "flights": [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, "src": 0, "dst": 2, "k": 1}, Expected: 200},
		},
		TimeComplexity:  "O(k * E)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def findCheapestPrice(n, flights, src, dst, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use Bellman-Ford variant limited to k+1 iterations, or BFS with level tracking.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Bellman-Ford: relax all edges k+1 times. Or modified Dijkstra tracking stops."},
			{Level: 3, Type: "code", Content: "prices = [inf]*n. prices[src]=0. for k+1 times: temp=prices.copy(). for u,v,p: temp[v]=min(temp[v], prices[u]+p). prices=temp."},
		},
		Solution: Solution{
			Code: `def findCheapestPrice(n, flights, src, dst, k):
    prices = [float('inf')] * n
    prices[src] = 0

    for _ in range(k + 1):
        temp = prices.copy()
        for u, v, p in flights:
            if prices[u] != float('inf'):
                temp[v] = min(temp[v], prices[u] + p)
        prices = temp

    return prices[dst] if prices[dst] != float('inf') else -1`,
			Explanation:     "Bellman-Ford with k+1 iterations. Each iteration allows one more stop.",
			TimeComplexity:  "O(k * E)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Limited iterations", Explanation: "k+1 iterations for at most k stops", CodeSnippet: "for _ in range(k + 1):", LineStart: 5, LineEnd: 5},
				{Title: "Copy before update", Explanation: "Prevent using updates from same iteration", CodeSnippet: "temp = prices.copy()", LineStart: 6, LineEnd: 6},
				{Title: "Relax edges", Explanation: "Update if path through u is cheaper", CodeSnippet: "temp[v] = min(temp[v], prices[u] + p)", LineStart: 9, LineEnd: 9},
			},
		},
	},
	{
		ID:              "redundant-connection",
		Number:          132,
		Title:           "Redundant Connection",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Union Find", "DFS"},
		RelatedChapters: []int{6, 12},
		Description: `In this problem, a tree is an undirected graph that is connected and has no cycles.

You are given a graph that started as a tree with n nodes labeled from 1 to n, with one additional edge added. The added edge has two different vertices chosen from 1 to n, and was not an edge that already existed.

Return an edge that can be removed so that the resulting graph is a tree of n nodes. If there are multiple answers, return the answer that occurs last in the input.`,
		Constraints: []string{
			"n == edges.length",
			"3 <= n <= 1000",
			"edges[i].length == 2",
			"1 <= ai < bi <= edges.length",
			"ai != bi",
			"There are no repeated edges",
			"The given graph is connected",
		},
		Examples: []Example{
			{Input: "edges = [[1,2],[1,3],[2,3]]", Output: "[2,3]"},
			{Input: "edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]", Output: "[1,4]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"edges": [][]int{{1, 2}, {1, 3}, {2, 3}}}, Expected: []int{2, 3}},
			{Input: map[string]interface{}{"edges": [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}}, Expected: []int{1, 4}},
		},
		TimeComplexity:  "O(n * alpha(n))",
		SpaceComplexity: "O(n)",
		StarterCode:     "def findRedundantConnection(edges):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use Union-Find. The edge that creates a cycle is the redundant one.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Process edges in order. If two nodes already connected, that edge is redundant."},
			{Level: 3, Type: "code", Content: "parent = list(range(n+1)). find with path compression. union: if find(u)==find(v): return [u,v]."},
		},
		Solution: Solution{
			Code: `def findRedundantConnection(edges):
    parent = list(range(len(edges) + 1))
    rank = [0] * (len(edges) + 1)

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    def union(x, y):
        px, py = find(x), find(y)
        if px == py:
            return False
        if rank[px] < rank[py]:
            px, py = py, px
        parent[py] = px
        if rank[px] == rank[py]:
            rank[px] += 1
        return True

    for u, v in edges:
        if not union(u, v):
            return [u, v]

    return []`,
			Explanation:     "Union-Find detects cycle. When union fails (same component), that's the redundant edge.",
			TimeComplexity:  "O(n * alpha(n))",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Path compression", Explanation: "Flatten tree during find", CodeSnippet: "parent[x] = find(parent[x])", LineStart: 7, LineEnd: 7},
				{Title: "Cycle detection", Explanation: "Same root means already connected", CodeSnippet: "if px == py:\n    return False", LineStart: 12, LineEnd: 13},
				{Title: "Return last", Explanation: "Process in order, return first failure", CodeSnippet: "if not union(u, v):\n    return [u, v]", LineStart: 22, LineEnd: 23},
			},
		},
	},
	{
		ID:              "graph-valid-tree",
		Number:          133,
		Title:           "Graph Valid Tree",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Union Find", "BFS", "DFS"},
		RelatedChapters: []int{6, 12},
		Description: `You have a graph of n nodes labeled from 0 to n - 1. You are given an integer n and a list of edges where edges[i] = [ai, bi] indicates that there is an undirected edge between nodes ai and bi in the graph.

Return true if the edges of the given graph make up a valid tree, and false otherwise.`,
		Constraints: []string{
			"1 <= n <= 2000",
			"0 <= edges.length <= 5000",
			"edges[i].length == 2",
			"0 <= ai, bi < n",
			"ai != bi",
			"There are no self-loops or repeated edges",
		},
		Examples: []Example{
			{Input: "n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]", Output: "true"},
			{Input: "n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]", Output: "false"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 5, "edges": [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}}, Expected: true},
			{Input: map[string]interface{}{"n": 5, "edges": [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}}, Expected: false},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def validTree(n, edges):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "A valid tree has exactly n-1 edges, is connected, and has no cycles.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Check edges == n-1. Then use Union-Find or DFS to verify connectivity and no cycles."},
			{Level: 3, Type: "code", Content: "if len(edges) != n-1: return False. Use Union-Find: all unions must succeed (no cycle)."},
		},
		Solution: Solution{
			Code: `def validTree(n, edges):
    if len(edges) != n - 1:
        return False

    parent = list(range(n))

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    def union(x, y):
        px, py = find(x), find(y)
        if px == py:
            return False
        parent[px] = py
        return True

    return all(union(u, v) for u, v in edges)`,
			Explanation:     "Tree needs exactly n-1 edges and no cycles. Union-Find checks both.",
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Edge count", Explanation: "Tree has exactly n-1 edges", CodeSnippet: "if len(edges) != n - 1:\n    return False", LineStart: 2, LineEnd: 3},
				{Title: "No cycles", Explanation: "All unions must succeed", CodeSnippet: "return all(union(u, v) for u, v in edges)", LineStart: 19, LineEnd: 19},
			},
		},
	},
	{
		ID:              "number-of-connected-components",
		Number:          134,
		Title:           "Number of Connected Components in an Undirected Graph",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Union Find", "BFS", "DFS"},
		RelatedChapters: []int{6, 12},
		Description: `You have a graph of n nodes. You are given an integer n and an array edges where edges[i] = [ai, bi] indicates that there is an edge between ai and bi in the graph.

Return the number of connected components in the graph.`,
		Constraints: []string{
			"1 <= n <= 2000",
			"1 <= edges.length <= 5000",
			"edges[i].length == 2",
			"0 <= ai <= bi < n",
			"ai != bi",
			"There are no repeated edges",
		},
		Examples: []Example{
			{Input: "n = 5, edges = [[0,1],[1,2],[3,4]]", Output: "2"},
			{Input: "n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 5, "edges": [][]int{{0, 1}, {1, 2}, {3, 4}}}, Expected: 2},
			{Input: map[string]interface{}{"n": 5, "edges": [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}}, Expected: 1},
		},
		TimeComplexity:  "O(n + E)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def countComponents(n, edges):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use Union-Find or DFS/BFS from each unvisited node.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Union-Find: start with n components. Each successful union decreases count by 1."},
			{Level: 3, Type: "code", Content: "count = n. for u,v in edges: if find(u) != find(v): union(u,v); count -= 1. return count."},
		},
		Solution: Solution{
			Code: `def countComponents(n, edges):
    parent = list(range(n))

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    def union(x, y):
        px, py = find(x), find(y)
        if px != py:
            parent[px] = py
            return True
        return False

    count = n
    for u, v in edges:
        if union(u, v):
            count -= 1

    return count`,
			Explanation:     "Start with n components. Each successful union merges two components.",
			TimeComplexity:  "O(n + E)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Start with n", Explanation: "Initially each node is its own component", CodeSnippet: "count = n", LineStart: 16, LineEnd: 16},
				{Title: "Merge components", Explanation: "Successful union reduces count", CodeSnippet: "if union(u, v):\n    count -= 1", LineStart: 18, LineEnd: 19},
			},
		},
	},
	{
		ID:              "alien-dictionary",
		Number:          135,
		Title:           "Alien Dictionary",
		Difficulty:      "Hard",
		Category:        "graphs",
		Tags:            []string{"Graph", "Topological Sort", "BFS"},
		RelatedChapters: []int{6, 12},
		Description: `There is a new alien language that uses the English alphabet. However, the order of the letters is unknown to you.

You are given a list of strings words from the alien language's dictionary. The strings in words are sorted lexicographically by the rules of this new language.

Return a string of the unique letters in the new alien language sorted in lexicographically increasing order by the new language's rules. If there is no solution, return "". If there are multiple solutions, return any of them.`,
		Constraints: []string{
			"1 <= words.length <= 100",
			"1 <= words[i].length <= 100",
			"words[i] consists of only lowercase English letters",
		},
		Examples: []Example{
			{Input: `words = ["wrt","wrf","er","ett","rftt"]`, Output: `"wertf"`},
			{Input: `words = ["z","x"]`, Output: `"zx"`},
			{Input: `words = ["z","x","z"]`, Output: `""`, Explanation: "Invalid order."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"words": []string{"wrt", "wrf", "er", "ett", "rftt"}}, Expected: "wertf"},
			{Input: map[string]interface{}{"words": []string{"z", "x"}}, Expected: "zx"},
			{Input: map[string]interface{}{"words": []string{"z", "x", "z"}}, Expected: ""},
		},
		TimeComplexity:  "O(C)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def alienOrder(words):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Build a graph from adjacent word pairs, then topological sort.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Compare adjacent words to find ordering. First different char gives edge. Use Kahn's algorithm for topo sort."},
			{Level: 3, Type: "code", Content: "Build graph and indegree. BFS from zero-indegree nodes. If result length != unique chars, return empty."},
		},
		Solution: Solution{
			Code: `def alienOrder(words):
    from collections import defaultdict, deque

    # Build adjacency list and indegree
    adj = defaultdict(set)
    indegree = {c: 0 for word in words for c in word}

    for i in range(len(words) - 1):
        w1, w2 = words[i], words[i + 1]
        min_len = min(len(w1), len(w2))

        # Check for invalid case: prefix is longer
        if len(w1) > len(w2) and w1[:min_len] == w2[:min_len]:
            return ""

        for j in range(min_len):
            if w1[j] != w2[j]:
                if w2[j] not in adj[w1[j]]:
                    adj[w1[j]].add(w2[j])
                    indegree[w2[j]] += 1
                break

    # Topological sort
    queue = deque([c for c in indegree if indegree[c] == 0])
    result = []

    while queue:
        c = queue.popleft()
        result.append(c)
        for neighbor in adj[c]:
            indegree[neighbor] -= 1
            if indegree[neighbor] == 0:
                queue.append(neighbor)

    return "".join(result) if len(result) == len(indegree) else ""`,
			Explanation:     "Build graph from word orderings. Topological sort gives valid alphabet order.",
			TimeComplexity:  "O(C)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Invalid prefix", Explanation: "Longer word can't come before its prefix", CodeSnippet: "if len(w1) > len(w2) and w1[:min_len] == w2[:min_len]:\n    return \"\"", LineStart: 13, LineEnd: 14},
				{Title: "First difference", Explanation: "Gives ordering between chars", CodeSnippet: "if w1[j] != w2[j]:", LineStart: 17, LineEnd: 17},
				{Title: "Cycle check", Explanation: "Must include all characters", CodeSnippet: "len(result) == len(indegree)", LineStart: 35, LineEnd: 35},
			},
		},
	},
	{
		ID:              "walls-and-gates",
		Number:          139,
		Title:           "Walls and Gates",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Array", "BFS", "Matrix"},
		RelatedChapters: []int{6, 12},
		Description: `You are given an m x n grid rooms initialized with these three possible values.

-1 A wall or an obstacle.
0 A gate.
INF Infinity means an empty room. We use the value 2^31 - 1 = 2147483647 to represent INF.

Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, leave it as INF.`,
		Constraints: []string{
			"m == rooms.length",
			"n == rooms[i].length",
			"1 <= m, n <= 250",
			"rooms[i][j] is -1, 0, or 2^31 - 1",
		},
		Examples: []Example{
			{Input: "rooms = [[2147483647,-1,0,2147483647],[2147483647,2147483647,2147483647,-1],[2147483647,-1,2147483647,-1],[0,-1,2147483647,2147483647]]", Output: "[[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]"},
			{Input: "rooms = [[-1]]", Output: "[[-1]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"rooms": [][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}}}, Expected: [][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}}},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def wallsAndGates(rooms):\n    # Write your solution here (modify in place)\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Multi-source BFS starting from all gates simultaneously.", ChapterRef: 6},
			{Level: 2, Type: "algorithm", Content: "Add all gates to queue. BFS outward, updating distances level by level."},
			{Level: 3, Type: "code", Content: "queue = all gates. while queue: for neighbors: if INF: rooms[nr][nc] = rooms[r][c] + 1; add to queue."},
		},
		Solution: Solution{
			Code: `def wallsAndGates(rooms):
    from collections import deque

    if not rooms:
        return

    m, n = len(rooms), len(rooms[0])
    INF = 2147483647
    queue = deque()

    # Add all gates to queue
    for i in range(m):
        for j in range(n):
            if rooms[i][j] == 0:
                queue.append((i, j))

    # BFS from all gates
    while queue:
        r, c = queue.popleft()
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nr, nc = r + dr, c + dc
            if 0 <= nr < m and 0 <= nc < n and rooms[nr][nc] == INF:
                rooms[nr][nc] = rooms[r][c] + 1
                queue.append((nr, nc))`,
			Explanation:     "Multi-source BFS from all gates. First visit to each cell gives shortest distance.",
			TimeComplexity:  "O(m * n)",
			SpaceComplexity: "O(m * n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Collect gates", Explanation: "All gates are starting points", CodeSnippet: "if rooms[i][j] == 0:\n    queue.append((i, j))", LineStart: 14, LineEnd: 15},
				{Title: "Update distance", Explanation: "INF cells get distance from neighbor + 1", CodeSnippet: "rooms[nr][nc] = rooms[r][c] + 1", LineStart: 23, LineEnd: 23},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, GraphsProblems...)
}
