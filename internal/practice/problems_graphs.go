package practice

// GraphsProblems is the graphs ladder (Chapter 13): grid and graph traversal
// with DFS and BFS, multi-source BFS, topological sort and union-find.
// Weighted shortest paths and spanning trees live in problems_advanced_graphs.go.
var GraphsProblems = []*Problem{
	{
		ID:              "number-of-islands",
		Number:          200,
		Order:           1,
		Title:           "Number of Islands",
		Slug:            "number-of-islands",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "DFS", "BFS", "Matrix", "Connected Components"},
		RelatedChapters: []int{13, 8},
		Description: `You are given a map as a grid of strings: "1" is land and "0" is water.

Two land cells belong to the same **island** when you can walk from one to the other by stepping up, down, left or right (never diagonally) and only ever standing on land. Everything outside the grid counts as water.

Return how many islands the map contains.`,
		Constraints: []string{
			"m == grid.length, n == grid[i].length",
			"1 <= m, n <= 300",
			`grid[i][j] is "0" or "1"`,
		},
		Examples: []Example{
			{Input: `grid = [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]`, Output: "1", Explanation: "All the land is connected through the top row and the left block, so it is a single island even though it has a bay of water in it."},
			{Input: `grid = [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]`, Output: "3", Explanation: "The 2 x 2 block, the single middle cell and the pair at the bottom right never touch side by side. The middle cell touches the others only at corners, which does not count."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"grid": [][]string{{"1", "1", "1", "1", "0"}, {"1", "1", "0", "1", "0"}, {"1", "1", "0", "0", "0"}, {"0", "0", "0", "0", "0"}}}, Expected: 1, Hidden: false, Note: "Example 1: one big island with a bay"},
			{Input: map[string]interface{}{"grid": [][]string{{"1", "1", "0", "0", "0"}, {"1", "1", "0", "0", "0"}, {"0", "0", "1", "0", "0"}, {"0", "0", "0", "1", "1"}}}, Expected: 3, Hidden: false, Note: "Example 2: three separate islands"},
			{Input: map[string]interface{}{"grid": [][]string{{"0"}}}, Expected: 0, Hidden: true, Note: "single water cell: no islands"},
			{Input: map[string]interface{}{"grid": [][]string{{"1"}}}, Expected: 1, Hidden: true, Note: "single land cell"},
			{Input: map[string]interface{}{"grid": [][]string{{"1", "0", "1"}, {"0", "1", "0"}, {"1", "0", "1"}}}, Expected: 5, Hidden: true, Note: "diagonal neighbours do not connect: five islands"},
			{Input: map[string]interface{}{"grid": [][]string{{"1", "0", "1"}, {"1", "0", "1"}, {"1", "1", "1"}}}, Expected: 1, Hidden: true, Note: "U shape: the two arms meet only at the bottom (checking just the up/left neighbour counts 2)"},
			{Input: map[string]interface{}{"grid": [][]string{{"1", "1", "1", "1"}, {"1", "1", "1", "1"}, {"1", "1", "1", "1"}}}, Expected: 1, Hidden: true, Note: "all land is one island"},
			{Input: map[string]interface{}{"grid": [][]string{{"1", "1", "1", "1", "1"}, {"1", "0", "0", "0", "1"}, {"1", "0", "1", "0", "1"}, {"1", "0", "0", "0", "1"}, {"1", "1", "1", "1", "1"}}}, Expected: 2, Hidden: true, Note: "a ring of land with a separate one-cell island inside it"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def numIslands(grid):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose you are scanning the map row by row and you step onto a land cell. How could you tell whether it belongs to an island you already counted, and what could you do right now so that the rest of that island is never counted again?"},
			{Level: 2, Type: "insight", Content: "Treat each land cell as a graph node joined to its land neighbours above, below, left and right. An island is a connected component. Each time the scan finds land that no earlier flood reached, that is a new island: count it, then flood the whole island with DFS or BFS and mark every cell visited (for example by turning it into \"0\").", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Build the graph: a cell's neighbours are the four cells next to it that are inside the grid.\n2. Scan for an unvisited start: loop over every cell; skip water and cells already visited.\n3. Flood the component: from that start, visit every land cell you can reach and mark each one visited.\n4. Read off the answer: the number of floods you started is the number of islands."},
			{Level: 4, Type: "skeleton", Content: `def numIslands(grid):
    rows, cols = len(grid), len(grid[0])

    def sink(r, c):
        if ____:  # outside the grid, or not unvisited land
            return
        grid[r][c] = ____
        # visit the four neighbours
        ____

    islands = 0
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == "1":
                ____
    return islands`},
		},
		Solution: Solution{
			Code: `def numIslands(grid):
    rows, cols = len(grid), len(grid[0])

    def sink(r, c):
        if r < 0 or r >= rows or c < 0 or c >= cols or grid[r][c] != "1":
            return  # off the map, water, or already visited
        grid[r][c] = "0"  # mark visited by sinking the land
        sink(r + 1, c)
        sink(r - 1, c)
        sink(r, c + 1)
        sink(r, c - 1)

    islands = 0
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == "1":  # land no earlier flood reached
                islands += 1
                sink(r, c)
    return islands`,
			Explanation: "This is the plain form of the chapter's component-counting plan, and every later grid problem in this ladder builds on it. Scan the grid; each land cell that no earlier flood has reached starts a new island. Count it and immediately flood the whole island with a depth-first search that sinks every land cell it touches (turns it to \"0\"), so no other cell of that island can start a second count.",
			BruteForce:  "A tempting shortcut is to count a land cell as a new island when its upper and left neighbours are water. That is fast but wrong: in a U shape the two arms each look like a fresh start, so it counts 2 where there is 1. The safe brute force is to keep a list of islands found so far and, for every land cell, search all of them to see which one it touches, merging islands when a cell touches two. That costs O((m * n)^2) in the worst case, because every cell can be compared against every earlier cell.",
			KeyInsight:  "An island is a connected component of a graph whose nodes are land cells and whose edges join side-by-side land. A single DFS from any cell of a component reaches exactly that component, so one flood per component, and a visited mark that stops the next flood, counts components in one pass.",
			WhyItWorks:  "Invariant: after the scan passes a cell, every island containing an already-scanned land cell has been counted exactly once and fully sunk. When the scan meets a \"1\", no earlier flood reached it (otherwise it would be \"0\"), so its island is new; counting it is correct. The flood then sinks every cell reachable from it by side steps over land, which is exactly that island (DFS reaches all of a component and nothing outside it, because it only steps onto land). So no cell of this island can trigger another count later. Each cell is sunk once and looks at four neighbours, so the total work is O(m * n). The recursion can go m * n calls deep on a snake-shaped island; the judge raises Python's recursion limit, but the iterative BFS alternative avoids the issue entirely.",
			CommonMistakes: []string{
				"Marking a cell visited after recursing into its neighbours: the neighbours then recurse back into it forever. Sink the cell first, then visit its neighbours.",
				"Treating diagonal cells as connected. The problem only joins up, down, left and right; the checkerboard test expects 5 islands, not 1.",
				"Counting a land cell as a new island just because its up and left neighbours are water. A U-shaped island breaks this; you must flood to know which island a cell belongs to.",
				"Checking grid[r][c] == 1 (an int) instead of \"1\" (a string). The grid holds strings, so the comparison is never true and the answer is always 0.",
			},
			TimeComplexity:  "O(m * n): each cell is sunk at most once and checks its four neighbours.",
			SpaceComplexity: "O(m * n) in the worst case for the recursion stack (one long snake-shaped island); no extra visited set, since the grid itself records visits.",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the graph", Explanation: "The graph is implicit. A cell's neighbours are (r+1, c), (r-1, c), (r, c+1) and (r, c-1). The first line of sink rejects anything outside the grid or not unvisited land, so the four recursive calls can be made blindly.", CodeSnippet: "if r < 0 or r >= rows or c < 0 or c >= cols or grid[r][c] != \"1\":\n    return", LineStart: 5, LineEnd: 6},
				{Title: "Flood the component", Explanation: "Mark the cell before exploring from it, so a neighbour that steps back here sees \"0\" and stops. Then visit the four neighbours; together the calls reach every land cell of this island.", CodeSnippet: "grid[r][c] = \"0\"\nsink(r + 1, c)\nsink(r - 1, c)\nsink(r, c + 1)\nsink(r, c - 1)", LineStart: 7, LineEnd: 11},
				{Title: "Scan for an unvisited start", Explanation: "Visit every cell in row order. A \"1\" here was not reached by any earlier flood, so it is the first cell we have seen of a new island.", CodeSnippet: "for r in range(rows):\n    for c in range(cols):\n        if grid[r][c] == \"1\":", LineStart: 14, LineEnd: 16},
				{Title: "Read off the answer", Explanation: "Count the island and sink it at once. When the scan ends, the counter holds the number of floods started, which is the number of islands.", CodeSnippet: "islands += 1\nsink(r, c)\nreturn islands", LineStart: 17, LineEnd: 19},
			},
			Alternatives: []AltSolution{
				{
					Name: "Iterative BFS with a visited set",
					Code: `from collections import deque

def numIslands(grid):
    rows, cols = len(grid), len(grid[0])
    seen = set()
    islands = 0
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] != "1" or (r, c) in seen:
                continue
            islands += 1
            seen.add((r, c))
            queue = deque([(r, c)])
            while queue:
                cr, cc = queue.popleft()
                for nr, nc in ((cr + 1, cc), (cr - 1, cc), (cr, cc + 1), (cr, cc - 1)):
                    if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == "1" and (nr, nc) not in seen:
                        seen.add((nr, nc))
                        queue.append((nr, nc))
    return islands`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "Same plan, but the flood uses a queue instead of recursion, so there is no recursion depth to worry about on huge islands, and a visited set leaves the input grid unchanged. Notice that a cell is marked seen when it is added to the queue, not when it is taken out; otherwise one cell can be queued several times by different neighbours.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must sink mark grid[r][c] before it calls itself on the neighbours?", Answer: "The neighbour's own call will try to step back to this cell. If the cell is still \"1\" at that moment, the two calls bounce back and forth forever (or until the recursion limit). Marking first makes the back-step see \"0\" and return at once, which is what guarantees each cell is processed once."},
			{Prompt: "What changes if diagonal cells also counted as connected?", Answer: "Only the neighbour list: add the four diagonal moves (r+1, c+1), (r+1, c-1), (r-1, c+1), (r-1, c-1). The scan, the flood and the counting stay identical. That is the value of the plan: the graph changes, the algorithm doesn't. The checkerboard test would then have 1 island instead of 5."},
			{Prompt: "The solution sinks land in the input grid. When is that a problem, and how would you avoid it?", Answer: "If the caller still needs the map afterwards, destroying it is a side effect they did not ask for. Use a separate visited set (as in the BFS alternative), or copy the grid first. It costs O(m * n) extra memory but leaves the input untouched."},
		},
	},
	{
		ID:              "clone-graph",
		Number:          133,
		Order:           2,
		Title:           "Clone Graph",
		Slug:            "clone-graph",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "DFS", "BFS", "Hash Table"},
		RelatedChapters: []int{13, 2},
		Description: `You are handed one node of a connected **undirected** graph. Each node has an integer **val** and a list **neighbors** of the nodes it is joined to. Because the graph is undirected, if A is in B's neighbours then B is in A's neighbours.

Return a **deep copy** of the whole graph: a brand-new node for every original node, with the same values and the same connections between the new nodes. No node you return, and no node reachable from it, may be one of the original nodes.

In the tests the graph is written as a 1-indexed adjacency list: entry i lists the values of the neighbours of the node whose value is i + 1, and you are given the node with value 1. An empty list means the graph is empty, the argument is None, and you should return None.`,
		Constraints: []string{
			"The number of nodes is between 0 and 100",
			"1 <= Node.val <= 100, and every value is unique",
			"No self-loops and no repeated edges",
			"The graph is connected: every node can be reached from the given node",
		},
		Examples: []Example{
			{Input: "adjList = [[2,4],[1,3],[2,4],[1,3]]", Output: "[[2,4],[1,3],[2,4],[1,3]]", Explanation: "Four nodes in a cycle 1-2-3-4-1. The copy has four new nodes wired the same way, so its adjacency list is identical, but every node object is new."},
			{Input: "adjList = [[]]", Output: "[[]]", Explanation: "One node with no neighbours: return a new node with value 1."},
			{Input: "adjList = []", Output: "[]", Explanation: "Empty graph: the argument is None, so return None."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"node": [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}}, Expected: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}, Hidden: false, Note: "Example 1: a 4-cycle 1-2-3-4-1"},
			{Input: map[string]interface{}{"node": [][]int{{}}}, Expected: [][]int{{}}, Hidden: false, Note: "Example 2: one node with no neighbours"},
			{Input: map[string]interface{}{"node": [][]int{}}, Expected: [][]int{}, Hidden: true, Note: "empty graph: return None"},
			{Input: map[string]interface{}{"node": [][]int{{2}, {1}}}, Expected: [][]int{{2}, {1}}, Hidden: true, Note: "two nodes pointing at each other: the cycle 1->2->1 must not loop forever"},
			{Input: map[string]interface{}{"node": [][]int{{2, 3}, {1, 3}, {1, 2}}}, Expected: [][]int{{2, 3}, {1, 3}, {1, 2}}, Hidden: true, Note: "triangle: node 3 is reached twice and must be cloned once"},
			{Input: map[string]interface{}{"node": [][]int{{2}, {1, 3}, {2, 4}, {3, 5}, {4}}}, Expected: [][]int{{2}, {1, 3}, {2, 4}, {3, 5}, {4}}, Hidden: true, Note: "path of five nodes"},
			{Input: map[string]interface{}{"node": [][]int{{2, 3, 4, 5, 6}, {1}, {1}, {1}, {1}, {1}}}, Expected: [][]int{{2, 3, 4, 5, 6}, {1}, {1}, {1}, {1}, {1}}, Hidden: true, Note: "star: one hub with five leaves"},
		},
		TimeComplexity:  "O(V + E)",
		SpaceComplexity: "O(V)",
		StarterCode:     "# Node is provided:\n# class Node:\n#     def __init__(self, val=0, neighbors=None):\n#         self.val = val\n#         self.neighbors = neighbors if neighbors is not None else []\n\ndef cloneGraph(node):\n    # Write your solution here\n    pass",
		Harness: Harness{
			ParamTypes: map[string]string{"node": "GraphNode"},
			ReturnType: "GraphNode",
			Prelude: `class Node:
    def __init__(self, val=0, neighbors=None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []`,
			Validator: `def check(raw, args, result, expected):
    if result != expected:
        return False, "the copy's adjacency list differs from the original's"

    def walk(start):
        found = {}
        stack = [start] if start is not None else []
        while stack:
            node = stack.pop()
            if id(node) not in found:
                found[id(node)] = node
                stack.extend(node.neighbors)
        return found

    originals = walk(args["node"])
    shared = sorted(node.val for key, node in walk(result_obj).items() if key in originals)
    if shared:
        return False, "nodes %s are the original objects, not new copies" % shared
    return True`,
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Node 1 and node 2 point at each other. If you copy node 1, then copy its neighbour 2, then copy 2's neighbour 1... when should you stop, and how would you find the copy of node 1 you already made?"},
			{Level: 2, Type: "insight", Content: "Traverse the graph (DFS or BFS) with a dictionary from each original node to its copy. The dictionary is both the visited set and the way to find an existing copy. Put a node's copy in the dictionary before you visit its neighbours, so a cycle that leads back finds the copy instead of making another.", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Handle the empty graph: None in, None out.\n2. Mark when you discover: when you first meet an original node, create its copy and store original -> copy right away.\n3. Expand one node: for each neighbour of the original, get its copy (making it if this is the first visit) and append that copy to the new node's neighbours.\n4. Read off the answer: return the copy of the node you were given."},
			{Level: 4, Type: "skeleton", Content: `def cloneGraph(node):
    if node is None:
        return None
    copies = {}

    def clone(original):
        if original in copies:
            return ____
        copy = Node(original.val)
        ____  # register the copy before going deeper
        for neighbor in original.neighbors:
            ____
        return copy

    return clone(node)`},
		},
		Solution: Solution{
			Code: `def cloneGraph(node):
    if node is None:
        return None
    copies = {}  # original node -> its clone

    def clone(original):
        if original in copies:
            return copies[original]  # already copied: reuse, don't recurse
        copy = Node(original.val)
        copies[original] = copy  # register BEFORE visiting neighbours
        for neighbor in original.neighbors:
            copy.neighbors.append(clone(neighbor))
        return copy

    return clone(node)`,
			Explanation: "Number of Islands flooded a grid with a visited mark; here the graph is explicit and the visited mark becomes a dictionary from each original node to its copy. A DFS visits every node once. On the first visit it creates the copy and records it; on any later visit (through a cycle or a second path) it returns the recorded copy. Each copy's neighbour list is built from the copies of the original's neighbours.",
			BruteForce:  "Copy recursively with no memory: copy a node, then copy each neighbour. On any cycle (even the edge 1-2, which is a cycle 1->2->1 in an undirected graph) this recurses forever. Patching it with a plain visited set stops the loop but loses the link: when you meet node 1 again you know you have seen it, but not which new object is its copy, so you cannot wire the edge.",
			KeyInsight:  "Use the dictionary original -> copy as the visited set. It answers both questions a traversal needs here: have I been here before, and if so, which copy did I make?",
			WhyItWorks:  "Invariant: once clone(original) has started, copies[original] holds the one and only copy of that node. It is stored before any neighbour is visited, so if a cycle leads back to original while its neighbours are still being processed, the lookup returns the existing copy instead of creating a duplicate or recursing again. Every node is reachable from the start node, so DFS calls clone on all of them; each creates exactly one copy. For every edge original -> neighbour, the loop appends copies[neighbour] to copies[original].neighbors, so the copies have exactly the same edges. Each node is expanded once and each adjacency entry is looked at once, so the work is O(V + E).",
			CommonMistakes: []string{
				"Storing the copy in the dictionary after the neighbour loop. A cycle reaches the node again before it is registered, so the recursion never ends. Register it right after creating it.",
				"Returning the input node, or appending original neighbours to the copy. The adjacency list looks right, but the result shares nodes with the input; the checker walks both graphs and rejects any shared node.",
				"Forgetting the empty graph: node is None, and node.val raises an AttributeError.",
				"Using node.val as the dictionary key works here because values are unique, but it breaks as soon as two nodes share a value. Keying by the node object itself is always safe.",
			},
			TimeComplexity:  "O(V + E): every node is copied once and every neighbour entry is followed once.",
			SpaceComplexity: "O(V) for the dictionary and the recursion stack (not counting the copy itself).",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle the empty graph", Explanation: "With no nodes there is nothing to copy; the tests represent this as an empty adjacency list and expect None back.", CodeSnippet: "if node is None:\n    return None\ncopies = {}", LineStart: 2, LineEnd: 4},
				{Title: "Mark when you discover", Explanation: "If this node was already copied, hand back that copy. Otherwise create the copy and record it immediately, before looking at any neighbour, so cycles find it.", CodeSnippet: "if original in copies:\n    return copies[original]\ncopy = Node(original.val)\ncopies[original] = copy", LineStart: 7, LineEnd: 10},
				{Title: "Expand one node", Explanation: "For each original neighbour, clone(neighbor) returns its copy (made now or earlier). Appending it wires the same edge between the copies.", CodeSnippet: "for neighbor in original.neighbors:\n    copy.neighbors.append(clone(neighbor))", LineStart: 11, LineEnd: 12},
				{Title: "Read off the answer", Explanation: "The copy of the given node is the entry point into the new graph; everything else is reachable from it.", CodeSnippet: "return clone(node)", LineStart: 15, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{
					Name: "Iterative BFS",
					Code: `from collections import deque

def cloneGraph(node):
    if node is None:
        return None
    copies = {node: Node(node.val)}
    queue = deque([node])
    while queue:
        original = queue.popleft()
        for neighbor in original.neighbors:
            if neighbor not in copies:
                copies[neighbor] = Node(neighbor.val)
                queue.append(neighbor)
            copies[original].neighbors.append(copies[neighbor])
    return copies[node]`,
					TimeComplexity:  "O(V + E)",
					SpaceComplexity: "O(V)",
					Explanation:     "The same dictionary drives a queue-based traversal. A node's copy is created and stored the moment the node is discovered (pushed), and every time an original edge is seen, the matching copy edge is appended. No recursion, so no depth limit.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the dictionary have to map to the copy rather than just record that a node was visited?", Answer: "When a cycle brings you back to a node, you must connect the current copy to that node's existing copy. A visited set only says 'seen it'; the dictionary also says 'here is the object to link to'. Without it you either create a duplicate copy or cannot wire the edge."},
			{Prompt: "In an undirected graph the edge 1-2 appears in both nodes' lists. Does the solution add the copied edge twice?", Answer: "It adds it once in each direction, which is exactly right: copies[1].neighbors gets copies[2] while expanding 1, and copies[2].neighbors gets copies[1] while expanding 2. Each original adjacency entry produces one copied entry, so the copy mirrors the original lists one for one."},
		},
	},
	{
		ID:              "walls-and-gates",
		Number:          286,
		Order:           3,
		Title:           "Walls and Gates",
		Slug:            "walls-and-gates",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "BFS", "Multi-source BFS", "Matrix"},
		RelatedChapters: []int{13},
		Description: `A building is drawn as a grid **rooms** of integers:

- -1 is a wall,
- 0 is a gate,
- 2147483647 (2^31 - 1, call it EMPTY) is an empty room.

From any cell you may step up, down, left or right into a cell that is not a wall. Fill every empty room, in place, with the number of steps from that room to its **nearest** gate. A room from which no gate can be reached keeps the value 2147483647. Walls and gates keep their values.

The function returns nothing; the judge reads the modified grid.`,
		Constraints: []string{
			"m == rooms.length, n == rooms[i].length",
			"1 <= m, n <= 250",
			"rooms[i][j] is -1, 0 or 2147483647",
		},
		Examples: []Example{
			{Input: "rooms = [[2147483647,-1,0,2147483647],[2147483647,2147483647,2147483647,-1],[2147483647,-1,2147483647,-1],[0,-1,2147483647,2147483647]]", Output: "[[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]", Explanation: "The top-left room is 3 steps from the gate in the bottom-left corner (down, down, down). The bottom-right room needs 4 steps from the top gate, going around the walls."},
			{Input: "rooms = [[-1]]", Output: "[[-1]]", Explanation: "A lone wall stays a wall."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"rooms": [][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}}}, Expected: [][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}}, Hidden: false, Note: "Example 1: two gates, walls in between"},
			{Input: map[string]interface{}{"rooms": [][]int{{-1}}}, Expected: [][]int{{-1}}, Hidden: false, Note: "Example 2: a single wall stays a wall"},
			{Input: map[string]interface{}{"rooms": [][]int{{2147483647}}}, Expected: [][]int{{2147483647}}, Hidden: true, Note: "no gate at all: the room stays 2147483647"},
			{Input: map[string]interface{}{"rooms": [][]int{{0, 2147483647, 2147483647, 2147483647, 0}}}, Expected: [][]int{{0, 1, 2, 1, 0}}, Hidden: true, Note: "two gates: each room takes the nearer one"},
			{Input: map[string]interface{}{"rooms": [][]int{{0, -1, 2147483647}}}, Expected: [][]int{{0, -1, 2147483647}}, Hidden: true, Note: "room walled off from every gate keeps 2147483647"},
			{Input: map[string]interface{}{"rooms": [][]int{{2147483647, 2147483647}, {2147483647, 0}}}, Expected: [][]int{{2, 1}, {1, 0}}, Hidden: true, Note: "one gate in a corner"},
			{Input: map[string]interface{}{"rooms": [][]int{{0, -1, 2147483647}, {2147483647, -1, 2147483647}, {2147483647, 2147483647, 2147483647}}}, Expected: [][]int{{0, -1, 6}, {1, -1, 5}, {2, 3, 4}}, Hidden: true, Note: "path must go around a wall"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def wallsAndGates(rooms):\n    # Modify rooms in place; return nothing\n    pass",
		Harness:         Harness{InPlace: "rooms"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Running a search from every room to find its nearest gate repeats a lot of work. What if the search ran the other way, starting from the gates? And what if all the gates started at the same moment?"},
			{Level: 2, Type: "insight", Content: "Multi-source BFS: put every gate in the queue at distance 0, then run one ordinary BFS. BFS reaches cells in order of distance from the set of starts, so the first time it reaches a room is via its nearest gate. Writing the distance into the room also marks it visited (it is no longer EMPTY).", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Build the graph: neighbours are the four adjacent cells inside the grid.\n2. Seed the frontier: queue every gate (value 0).\n3. Expand one node: pop a cell; for each neighbour that is still EMPTY, set it to this cell's value + 1 and queue it.\n4. Mark when you discover: the write in step 3 is the visited mark, so each room is set once.\n5. Read off the answer: the grid itself is the answer; rooms never reached stay EMPTY."},
			{Level: 4, Type: "skeleton", Content: `from collections import deque

def wallsAndGates(rooms):
    EMPTY = 2147483647
    rows, cols = len(rooms), len(rooms[0])
    queue = deque(____)  # every gate
    while queue:
        r, c = queue.popleft()
        for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
            if 0 <= nr < rows and 0 <= nc < cols and ____:
                rooms[nr][nc] = ____
                queue.append((nr, nc))`},
		},
		Solution: Solution{
			Code: `from collections import deque

def wallsAndGates(rooms):
    EMPTY = 2147483647
    rows, cols = len(rooms), len(rooms[0])
    # Seed the frontier with EVERY gate at once (distance 0).
    queue = deque((r, c) for r in range(rows) for c in range(cols) if rooms[r][c] == 0)
    while queue:
        r, c = queue.popleft()
        for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
            if 0 <= nr < rows and 0 <= nc < cols and rooms[nr][nc] == EMPTY:
                rooms[nr][nc] = rooms[r][c] + 1  # first visit = nearest gate
                queue.append((nr, nc))`,
			Explanation: "Number of Islands flooded from one start at a time; here we flood from all gates at once. That is multi-source BFS: seed the queue with every gate, then let a single BFS spread outward. Each room is written the first time the wave reaches it, which is from its nearest gate.",
			BruteForce:  "Run a separate BFS from every gate and keep the minimum distance in each room (this is the alternative below), or a BFS from every room until it hits a gate. With g gates the first version costs O(g * m * n), and in the worst case g is proportional to m * n, giving O((m * n)^2): each wave re-walks rooms that a closer gate already settled.",
			KeyInsight:  "The distance to the nearest gate is the distance from a virtual super-source joined to every gate. Starting BFS with all gates in the queue at distance 0 is exactly BFS from that super-source, so one pass settles every room.",
			WhyItWorks:  "BFS invariant: the queue always holds cells in non-decreasing order of distance, and the distances in it differ by at most 1. All gates enter at distance 0; their neighbours get 1 and go to the back; those neighbours' neighbours get 2, and so on. So the first time any room is reached, it is reached along a shortest path from some gate, and no other gate can be closer (a closer gate's wave would have arrived in an earlier layer). The room is then no longer EMPTY, so it is never overwritten with a larger value. Every cell is enqueued at most once and looks at four neighbours: O(m * n).",
			CommonMistakes: []string{
				"Running BFS from one gate at a time and overwriting rooms: a later, farther gate replaces a smaller distance. Either take the minimum (slow) or start all gates together (fast).",
				"Using DFS instead of BFS. DFS reaches a room along some path, not the shortest one, so the first value written can be too large.",
				"Updating cells that are walls or gates. Only rooms that are still EMPTY should be written; checking == EMPTY handles walls, gates and already-settled rooms in one comparison.",
			},
			TimeComplexity:  "O(m * n): each room is written and queued once.",
			SpaceComplexity: "O(m * n) for the queue in the worst case.",
			Walkthrough: []WalkthroughStep{
				{Title: "Seed the frontier", Explanation: "Every gate starts the BFS at distance 0. Because they are all in the queue before anything else, the waves from all gates grow together, one step at a time.", CodeSnippet: "queue = deque((r, c) for r in range(rows) for c in range(cols) if rooms[r][c] == 0)", LineStart: 7, LineEnd: 7},
				{Title: "Expand one node", Explanation: "Take the oldest cell in the queue and look at its four neighbours.", CodeSnippet: "r, c = queue.popleft()\nfor nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):", LineStart: 9, LineEnd: 10},
				{Title: "Mark when you discover", Explanation: "Only rooms that are still EMPTY are new. Writing distance + 1 both records the answer and marks the room visited, so no later (equal or longer) path changes it.", CodeSnippet: "if 0 <= nr < rows and 0 <= nc < cols and rooms[nr][nc] == EMPTY:\n    rooms[nr][nc] = rooms[r][c] + 1\n    queue.append((nr, nc))", LineStart: 11, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "One BFS per gate, keep the minimum",
					Code: `from collections import deque

def wallsAndGates(rooms):
    rows, cols = len(rooms), len(rooms[0])
    gates = [(r, c) for r in range(rows) for c in range(cols) if rooms[r][c] == 0]
    for gate in gates:
        dist = {gate: 0}
        queue = deque([gate])
        while queue:
            r, c = queue.popleft()
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < rows and 0 <= nc < cols and rooms[nr][nc] > 0 and (nr, nc) not in dist:
                    dist[(nr, nc)] = dist[(r, c)] + 1
                    rooms[nr][nc] = min(rooms[nr][nc], dist[(nr, nc)])
                    queue.append((nr, nc))`,
					TimeComplexity:  "O(g * m * n) for g gates",
					SpaceComplexity: "O(m * n)",
					Explanation:     "Each gate runs its own BFS with its own distance table and lowers a room's value whenever it finds a shorter route. Correct, but rooms near several gates are walked once per gate. It shows exactly the repeated work that seeding all gates together removes.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the first value written into a room guaranteed to be its distance to the nearest gate?", Answer: "BFS from several starts processes cells layer by layer: all cells at distance 0 (the gates), then all at distance 1, then 2. A room at true distance d from its nearest gate cannot be reached in a layer before d, and it is reached in layer d via that gate. Since the room is written in the first layer that reaches it, the value is d."},
			{Prompt: "What changes if each gate had a different opening delay (a gate opens t minutes late, so its distances start at t)?", Answer: "The layers are no longer aligned: a gate that starts at 5 is behind one that starts at 0, so plain BFS order is wrong. Use Dijkstra's algorithm (Chapter 14) with a heap seeded by (t, gate) pairs; it processes cells in order of total distance even when starts differ."},
		},
	},
	{
		ID:              "rotting-oranges",
		Number:          994,
		Order:           4,
		Title:           "Rotting Oranges",
		Slug:            "rotting-oranges",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "BFS", "Multi-source BFS", "Matrix"},
		RelatedChapters: []int{13},
		Description: `A grid holds oranges:

- 0 is an empty cell,
- 1 is a fresh orange,
- 2 is a rotten orange.

Every minute, each fresh orange that is directly above, below, left or right of a rotten orange becomes rotten. All of this happens at once: an orange that turns rotten during minute t starts spreading only in minute t + 1.

Return the number of minutes until no fresh orange is left. If some fresh orange can never rot, return -1. If there is no fresh orange at the start, the answer is 0.`,
		Constraints: []string{
			"m == grid.length, n == grid[i].length",
			"1 <= m, n <= 10",
			"grid[i][j] is 0, 1 or 2",
		},
		Examples: []Example{
			{Input: "grid = [[2,1,1],[1,1,0],[0,1,1]]", Output: "4", Explanation: "Minute 1 rots (0,1) and (1,0); minute 2 rots (0,2) and (1,1); minute 3 rots (2,1); minute 4 rots (2,2). Nothing fresh remains."},
			{Input: "grid = [[2,1,1],[0,1,1],[1,0,1]]", Output: "-1", Explanation: "The orange at (2,0) has empty cells above and to its right, so rot never reaches it."},
			{Input: "grid = [[0,2]]", Output: "0", Explanation: "There are no fresh oranges, so no time is needed."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"grid": [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}}, Expected: 4, Hidden: false, Note: "Example 1: rot spreads for 4 minutes"},
			{Input: map[string]interface{}{"grid": [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}}, Expected: -1, Hidden: false, Note: "Example 2: the bottom-left orange is cut off"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 2}}}, Expected: 0, Hidden: false, Note: "Example 3: no fresh orange, so 0 minutes"},
			{Input: map[string]interface{}{"grid": [][]int{{0}}}, Expected: 0, Hidden: true, Note: "no oranges at all"},
			{Input: map[string]interface{}{"grid": [][]int{{1}}}, Expected: -1, Hidden: true, Note: "a fresh orange with no rotten source"},
			{Input: map[string]interface{}{"grid": [][]int{{2, 1, 1, 1, 2}}}, Expected: 2, Hidden: true, Note: "two sources rot from both ends at the same time"},
			{Input: map[string]interface{}{"grid": [][]int{{2, 1}}}, Expected: 1, Hidden: true, Note: "one minute; counting the final empty level gives 2"},
			{Input: map[string]interface{}{"grid": [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}}}, Expected: 2, Hidden: true, Note: "sources in opposite corners"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def orangesRotting(grid):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The oranges that rot in minute 1 are exactly those one step from an initially rotten orange; minute 2, two steps. Which traversal visits cells in order of their step count, and how would you know where one minute ends and the next begins?"},
			{Level: 2, Type: "insight", Content: "This is Walls and Gates with a clock: multi-source BFS from every rotten orange. Process the queue one level at a time (take len(queue) cells, which are exactly the oranges that rotted in the previous minute); each level is one minute. Keep a count of fresh oranges so you can tell at the end whether any were never reached.", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Seed the frontier: queue every rotten orange and count the fresh ones.\n2. Expand one level: while there are fresh oranges and the queue is not empty, pop exactly the cells that are in the queue right now; rot their fresh neighbours, decrement the fresh count and queue them.\n3. After each level, add one minute.\n4. Read off the answer: if the fresh count is 0, return the minutes; otherwise return -1."},
			{Level: 4, Type: "skeleton", Content: `from collections import deque

def orangesRotting(grid):
    rows, cols = len(grid), len(grid[0])
    queue = deque()
    fresh = 0
    # queue every rotten orange, count the fresh ones
    ____

    minutes = 0
    while queue and ____:
        for _ in range(____):
            r, c = queue.popleft()
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == 1:
                    ____
        minutes += 1
    return ____`},
		},
		Solution: Solution{
			Code: `from collections import deque

def orangesRotting(grid):
    rows, cols = len(grid), len(grid[0])
    queue = deque()
    fresh = 0
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == 2:
                queue.append((r, c))  # every rotten orange is a source
            elif grid[r][c] == 1:
                fresh += 1

    minutes = 0
    while queue and fresh > 0:
        for _ in range(len(queue)):  # one BFS level = one minute
            r, c = queue.popleft()
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == 1:
                    grid[nr][nc] = 2
                    fresh -= 1
                    queue.append((nr, nc))
        minutes += 1
    return minutes if fresh == 0 else -1`,
			Explanation: "This builds on Walls and Gates (multi-source BFS) and adds one idea: process the queue level by level so you can count levels, which are minutes. Every rotten orange seeds the queue; each level rots the fresh neighbours of the previous level. A running count of fresh oranges tells you when to stop and whether any were unreachable.",
			BruteForce:  "Simulate the clock directly: each minute, scan the whole grid, find every fresh orange next to a rotten one, rot them all, and repeat until nothing changes. That is the alternative below. Each scan costs O(m * n) and there can be O(m * n) minutes (a long snake of oranges), so it is O((m * n)^2). Each scan re-checks oranges far from the rot that cannot possibly change this minute.",
			KeyInsight:  "The minute an orange rots equals its BFS distance from the nearest initially rotten orange. A multi-source BFS computes all those distances in one pass, and the answer is the largest one, which is the number of levels processed.",
			WhyItWorks:  "Invariant: at the start of each pass of the while loop, the queue holds exactly the oranges that became rotten in the previous minute (initially, the ones rotten at minute 0). The inner loop pops exactly len(queue) of them, so newly rotted oranges are left for the next pass, matching the rule that new rot spreads only from the next minute. Each fresh orange is rotted at the earliest level that reaches it, which is its distance from the nearest rotten source. The loop stops when no fresh oranges remain (so it doesn't count an extra, empty minute) or when the queue is empty (rot can't spread). If fresh oranges remain, they are in regions no rot can reach, so -1 is correct.",
			CommonMistakes: []string{
				"Starting BFS from only the first rotten orange found. All rotten oranges spread at the same time; two sources at either end of a row finish in half the time.",
				"Counting a minute for the final level that rots nothing, giving one too many. Stop when fresh reaches 0 (the loop condition 'queue and fresh > 0'), or subtract one at the end.",
				"Popping from the queue without fixing the level size first. Oranges rotted in this minute would spread in the same minute; take len(queue) before the inner loop.",
				"Forgetting the 'no fresh oranges at all' case: the answer is 0 even if there are no rotten oranges either.",
			},
			TimeComplexity:  "O(m * n): each orange is queued at most once.",
			SpaceComplexity: "O(m * n) for the queue.",
			Walkthrough: []WalkthroughStep{
				{Title: "Seed the frontier", Explanation: "Every rotten orange is a source at minute 0. Counting fresh oranges now lets us finish early and detect unreachable ones later.", CodeSnippet: "if grid[r][c] == 2:\n    queue.append((r, c))\nelif grid[r][c] == 1:\n    fresh += 1", LineStart: 9, LineEnd: 12},
				{Title: "Expand one level", Explanation: "len(queue) is fixed when range() is called, so the loop processes exactly the oranges that rotted last minute. Oranges rotted now wait for the next pass.", CodeSnippet: "while queue and fresh > 0:\n    for _ in range(len(queue)):\n        r, c = queue.popleft()", LineStart: 15, LineEnd: 17},
				{Title: "Mark when you discover", Explanation: "Setting the neighbour to 2 marks it visited, so no other rotten orange queues it again, and keeps the fresh count accurate.", CodeSnippet: "if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == 1:\n    grid[nr][nc] = 2\n    fresh -= 1\n    queue.append((nr, nc))", LineStart: 19, LineEnd: 22},
				{Title: "Read off the answer", Explanation: "One level is one minute. If any fresh orange is left, it was never reached, so return -1.", CodeSnippet: "minutes += 1\nreturn minutes if fresh == 0 else -1", LineStart: 23, LineEnd: 24},
			},
			Alternatives: []AltSolution{
				{
					Name: "Minute-by-minute simulation",
					Code: `def orangesRotting(grid):
    grid = [row[:] for row in grid]
    rows, cols = len(grid), len(grid[0])
    minutes = 0
    while True:
        to_rot = []
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] != 1:
                    continue
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == 2:
                        to_rot.append((r, c))
                        break
        if not to_rot:
            break
        for r, c in to_rot:
            grid[r][c] = 2
        minutes += 1
    if any(cell == 1 for row in grid for cell in row):
        return -1
    return minutes`,
					TimeComplexity:  "O((m * n)^2)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "Each minute, collect every fresh orange that touches a rotten one, then rot them all together (collecting first is what keeps new rot from spreading in the same minute). Stop when a minute changes nothing. Easy to trust, but every minute rescans the whole grid.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must the size of the level be read once, before popping, instead of looping 'while queue'?", Answer: "Oranges rotted during this minute are appended to the same queue. A plain 'while queue' loop would pop them in the same pass and let them spread immediately, merging several minutes into one and returning too small an answer. Reading len(queue) first fences off exactly the previous minute's oranges."},
			{Prompt: "How is this problem the same as Walls and Gates, and what is the one difference?", Answer: "Both are multi-source BFS on a grid: all sources (gates or rotten oranges) start in the queue at distance 0, and the first visit gives the shortest distance. Walls and Gates writes each distance into the grid; Rotting Oranges only needs the largest distance (the number of levels) plus a check that every fresh orange was reached."},
		},
	},
	{
		ID:              "pacific-atlantic-water-flow",
		Number:          417,
		Order:           5,
		Title:           "Pacific Atlantic Water Flow",
		Slug:            "pacific-atlantic-water-flow",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "DFS", "BFS", "Matrix", "Multi-source BFS"},
		RelatedChapters: []int{13},
		Description: `An island is drawn as a grid **heights**, where heights[r][c] is the height of cell (r, c). The **Pacific** ocean touches the whole top row and the whole left column; the **Atlantic** ocean touches the whole bottom row and the whole right column.

Rain water in a cell can flow to a neighbouring cell (up, down, left or right) whose height is **less than or equal to** its own. A cell on the edge drains into every ocean that its edge touches (so a corner cell, or every cell of a single-row grid, can touch both).

Return the coordinates [r, c] of every cell from which water can reach **both** oceans. You may return them in any order.`,
		Constraints: []string{
			"m == heights.length, n == heights[r].length",
			"1 <= m, n <= 200",
			"0 <= heights[r][c] <= 10^5",
		},
		Examples: []Example{
			{Input: "heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]", Output: "[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]", Explanation: "From (2,2), height 5, water can run up-left 5 -> 3 -> 2 to the top row (Pacific) and right 5 -> 3 -> 1 to the right column (Atlantic). From (0,4), the top-right corner, it is already on both oceans."},
			{Input: "heights = [[1]]", Output: "[[0,0]]", Explanation: "The single cell touches every edge, so both oceans."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"heights": [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}}, Expected: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}, Hidden: false, Note: "Example 1"},
			{Input: map[string]interface{}{"heights": [][]int{{1}}}, Expected: [][]int{{0, 0}}, Hidden: false, Note: "Example 2: a single cell touches both oceans"},
			{Input: map[string]interface{}{"heights": [][]int{{1, 1}, {1, 1}}}, Expected: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, Hidden: true, Note: "flat grid: water flows everywhere, so every cell counts"},
			{Input: map[string]interface{}{"heights": [][]int{{3, 1, 2}}}, Expected: [][]int{{0, 0}, {0, 1}, {0, 2}}, Hidden: true, Note: "single row: every cell borders both oceans"},
			{Input: map[string]interface{}{"heights": [][]int{{3, 3, 3}, {3, 1, 3}, {3, 3, 3}}}, Expected: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}, Hidden: true, Note: "a pit in the middle can't flow out"},
			{Input: map[string]interface{}{"heights": [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}}, Expected: [][]int{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}, Hidden: true, Note: "spiral of heights"},
			{Input: map[string]interface{}{"heights": [][]int{{1, 1, 1}, {1, 5, 1}, {1, 1, 1}}}, Expected: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}, Hidden: true, Note: "a peak in the middle flows to both sides"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def pacificAtlantic(heights):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Following the water downhill from each cell repeats the same paths over and over. The oceans are fixed and few. Could you instead start at an ocean and ask which cells can drain into it? Which way would you have to step?"},
			{Level: 2, Type: "insight", Content: "Reverse the flow. Water goes from a cell to a neighbour that is lower or equal, so walking backwards from the ocean you may step to a neighbour that is higher or equal. Run one multi-source search from all Pacific edge cells and one from all Atlantic edge cells; the answer is the cells reached by both.", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Build the graph: in the reversed graph, (r, c) leads to a neighbour whose height is >= heights[r][c].\n2. Seed the frontier: all cells in the top row and left column for the Pacific (the bottom row and right column for the Atlantic).\n3. Expand one node: from each reached cell, step to unreached neighbours that are at least as high; mark them as you add them.\n4. Read off the answer: return the cells in both reached sets."},
			{Level: 4, Type: "skeleton", Content: `def pacificAtlantic(heights):
    rows, cols = len(heights), len(heights[0])

    def climb(starts):
        reached = set(starts)
        stack = list(starts)
        while stack:
            r, c = stack.pop()
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < rows and 0 <= nc < cols and (nr, nc) not in reached and ____:
                    ____
        return reached

    pacific = climb(____)
    atlantic = climb(____)
    return ____`},
		},
		Solution: Solution{
			Code: `def pacificAtlantic(heights):
    rows, cols = len(heights), len(heights[0])

    def climb(starts):
        reached = set(starts)
        stack = list(starts)
        while stack:
            r, c = stack.pop()
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if (0 <= nr < rows and 0 <= nc < cols and (nr, nc) not in reached
                        and heights[nr][nc] >= heights[r][c]):  # water could flow nr,nc -> r,c
                    reached.add((nr, nc))
                    stack.append((nr, nc))
        return reached

    pacific = climb([(r, 0) for r in range(rows)] + [(0, c) for c in range(cols)])
    atlantic = climb([(r, cols - 1) for r in range(rows)] + [(rows - 1, c) for c in range(cols)])
    return [[r, c] for r, c in pacific & atlantic]`,
			Explanation: "Like Walls and Gates, start from the targets instead of from every cell: one multi-source search per ocean. The twist is that edges are directed (water only runs downhill), so the reverse search climbs: from a reached cell it steps to neighbours that are at least as high. A cell can drain to an ocean exactly when that ocean's climb reaches it, so the answer is the intersection of the two reached sets.",
			BruteForce:  "For every cell, search downhill (to neighbours of lower or equal height) and record whether the search touches a Pacific edge and an Atlantic edge. That is the alternative below. Each search can visit the whole grid, so the total is O((m * n)^2); cells in the middle of a slope are re-walked once for every cell above them.",
			KeyInsight:  "\"Can water get from X to the ocean?\" equals \"can I walk from the ocean to X against the flow?\". Reversing the edges turns m * n separate questions into two searches, one per ocean, each visiting every cell at most once.",
			WhyItWorks:  "Water can move from a to b when b is adjacent and heights[b] <= heights[a]. A drainage path X = c0, c1, ..., ck with ck on the Pacific edge exists exactly when the reverse path ck, ..., c0 exists in the reversed graph, where you may step from b to a when heights[a] >= heights[b]. The Pacific climb starts at every Pacific edge cell and follows reversed edges, so it reaches X exactly when such a path exists. The same holds for the Atlantic. Therefore X is in both sets exactly when water from X can reach both oceans. Each climb adds each cell to its set at most once and looks at four neighbours, so both together take O(m * n).",
			CommonMistakes: []string{
				"Climbing with > instead of >=. Water flows onto equal heights, so equal neighbours must be included; the flat-grid test expects every cell.",
				"Searching downhill from the ocean cells (heights[nr][nc] <= heights[r][c]). That finds where water from the edge would go, not where water comes from.",
				"Seeding only one corner, or only one edge, of each ocean. The Pacific touches the entire top row and the entire left column.",
				"Recomputing a full search for every cell, which is correct but O((m * n)^2) and too slow for a 200 x 200 grid.",
			},
			TimeComplexity:  "O(m * n): two searches, each visiting every cell at most once.",
			SpaceComplexity: "O(m * n) for the two reached sets and the stack.",
			Walkthrough: []WalkthroughStep{
				{Title: "Seed the frontier", Explanation: "Every edge cell of an ocean drains into it, so all of them start the search and are already reached.", CodeSnippet: "reached = set(starts)\nstack = list(starts)", LineStart: 5, LineEnd: 6},
				{Title: "Expand one node", Explanation: "Pop a reached cell and look at its four neighbours. A neighbour that is at least as high could pour water into this cell, so it can also reach the ocean. Mark it and keep climbing.", CodeSnippet: "r, c = stack.pop()\nand heights[nr][nc] >= heights[r][c]):\nreached.add((nr, nc))\nstack.append((nr, nc))", LineStart: 8, LineEnd: 13},
				{Title: "Seed each ocean", Explanation: "The Pacific search starts from the left column and top row; the Atlantic search starts from the right column and bottom row.", CodeSnippet: "pacific = climb([(r, 0) for r in range(rows)] + [(0, c) for c in range(cols)])\natlantic = climb([(r, cols - 1) for r in range(rows)] + [(rows - 1, c) for c in range(cols)])", LineStart: 16, LineEnd: 17},
				{Title: "Read off the answer", Explanation: "A cell reachable from both climbs drains into both oceans. Set intersection gives exactly those.", CodeSnippet: "return [[r, c] for r, c in pacific & atlantic]", LineStart: 18, LineEnd: 18},
			},
			Alternatives: []AltSolution{
				{
					Name: "Downhill search from every cell",
					Code: `def pacificAtlantic(heights):
    rows, cols = len(heights), len(heights[0])
    result = []
    for sr in range(rows):
        for sc in range(cols):
            seen = {(sr, sc)}
            stack = [(sr, sc)]
            pacific = atlantic = False
            while stack:
                r, c = stack.pop()
                if r == 0 or c == 0:
                    pacific = True
                if r == rows - 1 or c == cols - 1:
                    atlantic = True
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < rows and 0 <= nc < cols and (nr, nc) not in seen and heights[nr][nc] <= heights[r][c]:
                        seen.add((nr, nc))
                        stack.append((nr, nc))
            if pacific and atlantic:
                result.append([sr, sc])
    return result`,
					TimeComplexity:  "O((m * n)^2)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "Follows the water directly: from each cell, explore every cell it can flow to and note whether that touches each ocean. It is the definition turned into code, which makes it a good check, but it repeats each slope's search for every cell above it.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the reverse search use >= rather than <=?", Answer: "In the reverse search we move from a cell that already drains to the ocean to a neighbour that could pour water into it. Water moves from higher-or-equal to lower-or-equal, so the neighbour must be at least as high: heights[neighbour] >= heights[current]."},
			{Prompt: "Why is it enough to run only two searches, not one per edge cell?", Answer: "All the Pacific edge cells play the same role: reaching any of them means reaching the Pacific. Putting them all in the starting stack is a multi-source search from a single virtual 'Pacific' node joined to each of them, so one search answers the question for the whole ocean. The same holds for the Atlantic."},
		},
	},
	{
		ID:              "surrounded-regions",
		Number:          130,
		Order:           6,
		Title:           "Surrounded Regions",
		Slug:            "surrounded-regions",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "DFS", "BFS", "Matrix"},
		RelatedChapters: []int{13},
		Description: `A board is a grid of "X" and "O" strings. A **region** is a group of "O" cells joined by up, down, left or right steps (not diagonals).

A region is **captured** when none of its cells lies on the border of the board: then it is completely enclosed by "X". Change every cell of every captured region to "X", in place. Regions that touch the border, even at a single cell, stay as they are.

The function returns nothing; the judge reads the modified board.`,
		Constraints: []string{
			"m == board.length, n == board[i].length",
			"1 <= m, n <= 200",
			`board[i][j] is "X" or "O"`,
		},
		Examples: []Example{
			{Input: `board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]`, Output: `[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]`, Explanation: "The three O cells in the middle form one region with no border cell, so they are captured. The O at the bottom (row 3, column 1) is on the border, so it stays."},
			{Input: `board = [["X"]]`, Output: `[["X"]]`, Explanation: "Nothing to capture."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"board": [][]string{{"X", "X", "X", "X"}, {"X", "O", "O", "X"}, {"X", "X", "O", "X"}, {"X", "O", "X", "X"}}}, Expected: [][]string{{"X", "X", "X", "X"}, {"X", "X", "X", "X"}, {"X", "X", "X", "X"}, {"X", "O", "X", "X"}}, Hidden: false, Note: "Example 1: the inner region is captured, the border O stays"},
			{Input: map[string]interface{}{"board": [][]string{{"X"}}}, Expected: [][]string{{"X"}}, Hidden: false, Note: "Example 2: a single X"},
			{Input: map[string]interface{}{"board": [][]string{{"O"}}}, Expected: [][]string{{"O"}}, Hidden: true, Note: "a single O is on the border, so it stays"},
			{Input: map[string]interface{}{"board": [][]string{{"O", "O", "O"}, {"O", "O", "O"}, {"O", "O", "O"}}}, Expected: [][]string{{"O", "O", "O"}, {"O", "O", "O"}, {"O", "O", "O"}}, Hidden: true, Note: "everything is connected to the border"},
			{Input: map[string]interface{}{"board": [][]string{{"X", "X", "X", "X", "X"}, {"X", "O", "O", "O", "X"}, {"X", "O", "X", "O", "X"}, {"X", "O", "O", "O", "X"}, {"X", "X", "X", "X", "X"}}}, Expected: [][]string{{"X", "X", "X", "X", "X"}, {"X", "X", "X", "X", "X"}, {"X", "X", "X", "X", "X"}, {"X", "X", "X", "X", "X"}, {"X", "X", "X", "X", "X"}}, Hidden: true, Note: "a ring of O fully enclosed is captured"},
			{Input: map[string]interface{}{"board": [][]string{{"X", "X", "X", "X"}, {"X", "O", "X", "X"}, {"X", "X", "O", "X"}, {"X", "X", "X", "O"}}}, Expected: [][]string{{"X", "X", "X", "X"}, {"X", "X", "X", "X"}, {"X", "X", "X", "X"}, {"X", "X", "X", "O"}}, Hidden: true, Note: "diagonal contact with the border does not save a cell"},
			{Input: map[string]interface{}{"board": [][]string{{"X", "X", "X", "X", "X"}, {"X", "O", "O", "X", "X"}, {"X", "X", "O", "X", "X"}, {"X", "X", "O", "O", "O"}, {"X", "X", "X", "X", "X"}}}, Expected: [][]string{{"X", "X", "X", "X", "X"}, {"X", "O", "O", "X", "X"}, {"X", "X", "O", "X", "X"}, {"X", "X", "O", "O", "O"}, {"X", "X", "X", "X", "X"}}, Hidden: true, Note: "inner cells joined to the border by a winding path stay"},
		},
		TimeComplexity:  "O(m * n)",
		SpaceComplexity: "O(m * n)",
		StarterCode:     "def solve(board):\n    # Modify board in place; return nothing\n    pass",
		Harness:         Harness{InPlace: "board"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Deciding whether a region is enclosed means exploring all of it. Is it easier to find the regions that are captured, or the ones that are not? Where must a region that escapes capture have at least one cell?"},
			{Level: 2, Type: "insight", Content: "Flip the question: the survivors are exactly the O cells connected to an O on the border. Flood from every border O and mark what you reach as safe (for example with a temporary \"S\"). Every O left unmarked is captured. This is the reverse-search trick from Pacific Atlantic.", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Seed the frontier: every O on the border.\n2. Flood the component: from each one, mark every connected O as safe (\"S\").\n3. Read off the answer: sweep the board; any remaining \"O\" is captured and becomes \"X\", and every \"S\" goes back to \"O\"."},
			{Level: 4, Type: "skeleton", Content: `def solve(board):
    rows, cols = len(board), len(board[0])

    def mark_safe(start_r, start_c):
        stack = [(start_r, start_c)]
        while stack:
            r, c = stack.pop()
            if 0 <= r < rows and 0 <= c < cols and board[r][c] == "O":
                ____
                stack.extend(((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)))

    # flood from every border cell
    ____

    for r in range(rows):
        for c in range(cols):
            ____`},
		},
		Solution: Solution{
			Code: `def solve(board):
    rows, cols = len(board), len(board[0])

    def mark_safe(start_r, start_c):
        stack = [(start_r, start_c)]
        while stack:
            r, c = stack.pop()
            if 0 <= r < rows and 0 <= c < cols and board[r][c] == "O":
                board[r][c] = "S"  # safe: connected to the border
                stack.extend(((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)))

    for r in range(rows):
        mark_safe(r, 0)
        mark_safe(r, cols - 1)
    for c in range(cols):
        mark_safe(0, c)
        mark_safe(rows - 1, c)

    for r in range(rows):
        for c in range(cols):
            if board[r][c] == "O":
                board[r][c] = "X"  # never reached from the border: captured
            elif board[r][c] == "S":
                board[r][c] = "O"  # restore the safe cells`,
			Explanation: "This combines the island flood from Number of Islands with the reverse search from Pacific Atlantic. Instead of testing each region for enclosure, flood from every border O and mark everything reached as safe. Whatever O is left was not connected to the border, so it is captured. A final sweep flips captured cells to X and restores the safe ones.",
			BruteForce:  "Find each O region with a flood, record whether any of its cells is on the border, and flip it if not (the alternative below). That is actually linear too, but it needs a list per region and a visited set, and the obvious first attempt, starting a fresh search from every O cell to see if it can reach the border, costs O((m * n)^2).",
			KeyInsight:  "A region survives exactly when it contains a border cell, which is the same as being reachable from a border O. So you only ever need to explore the survivors, starting from the border, and everything else is captured by elimination.",
			WhyItWorks:  "Every border O is a starting point, and the flood from each one reaches its entire region (DFS reaches the whole connected component and nothing else, since it only steps onto O). So after the border loops, a cell is \"S\" exactly when its region touches the border. A region with no border cell is never started from and never entered from outside (entering would require an O neighbour in another region, which would make them the same region), so all its cells are still \"O\". The final sweep therefore turns exactly the captured cells into \"X\" and restores the survivors. The temporary \"S\" doubles as the visited mark, so each cell is flooded once: O(m * n). The flood is an explicit stack so a 200 x 200 board of O cannot hit a recursion limit.",
			CommonMistakes: []string{
				"Flipping O to X while still flooding, before you know which regions are safe. Mark first, decide in a separate sweep.",
				"Forgetting to restore \"S\" back to \"O\" in the final sweep, leaving S characters in the output.",
				"Treating diagonal contact with the border as safe. Only side-by-side steps connect cells, so an O that touches a border O only at a corner is still captured.",
				"Flooding from only the first and last row and forgetting the first and last column (or the other way round).",
			},
			TimeComplexity:  "O(m * n): every cell is marked at most once and swept once.",
			SpaceComplexity: "O(m * n) for the stack in the worst case (a board full of O).",
			Walkthrough: []WalkthroughStep{
				{Title: "Flood the component", Explanation: "An explicit stack does the DFS. A cell is marked \"S\" as soon as it is popped as an O, which also stops it from being processed again; out-of-range or non-O cells are simply skipped.", CodeSnippet: "if 0 <= r < rows and 0 <= c < cols and board[r][c] == \"O\":\n    board[r][c] = \"S\"\n    stack.extend(((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)))", LineStart: 8, LineEnd: 10},
				{Title: "Seed the frontier", Explanation: "Start a flood from every cell of the left and right columns and the top and bottom rows. mark_safe ignores cells that are X or already safe.", CodeSnippet: "mark_safe(r, 0)\nmark_safe(r, cols - 1)\nfor c in range(cols):\n    mark_safe(0, c)\n    mark_safe(rows - 1, c)", LineStart: 13, LineEnd: 17},
				{Title: "Read off the answer", Explanation: "Any O still on the board was never reached from the border, so it is captured. Safe cells go back to O.", CodeSnippet: "if board[r][c] == \"O\":\n    board[r][c] = \"X\"\nelif board[r][c] == \"S\":\n    board[r][c] = \"O\"", LineStart: 21, LineEnd: 24},
			},
			Alternatives: []AltSolution{
				{
					Name: "Explore each region and test it",
					Code: `def solve(board):
    rows, cols = len(board), len(board[0])
    seen = set()
    for sr in range(rows):
        for sc in range(cols):
            if board[sr][sc] != "O" or (sr, sc) in seen:
                continue
            region = [(sr, sc)]
            seen.add((sr, sc))
            touches_border = False
            i = 0
            while i < len(region):
                r, c = region[i]
                i += 1
                if r in (0, rows - 1) or c in (0, cols - 1):
                    touches_border = True
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < rows and 0 <= nc < cols and board[nr][nc] == "O" and (nr, nc) not in seen:
                        seen.add((nr, nc))
                        region.append((nr, nc))
            if not touches_border:
                for r, c in region:
                    board[r][c] = "X"`,
					TimeComplexity:  "O(m * n)",
					SpaceComplexity: "O(m * n)",
					Explanation:     "Flood every O region once (BFS over a growing list), note whether any of its cells is on the border, and flip the whole region if none is. Also linear, but it has to explore captured regions as well as survivors and keeps every region's cells in a list.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it easier to search from the border than to test each region for enclosure?", Answer: "Survival has a simple certificate: one border cell in the region. Searching from all border O cells finds every survivor directly, and capture follows by elimination. Testing each region needs a full flood of that region plus bookkeeping to remember its cells until you know the verdict."},
			{Prompt: "What changes if diagonal neighbours also connected cells?", Answer: "Only the neighbour list in the flood: push the four diagonal cells as well. The seeding from the border and the final sweep stay the same. Then an inner O that touches a border O at a corner would be safe."},
		},
	},
	{
		ID:              "course-schedule",
		Number:          207,
		Order:           7,
		Title:           "Course Schedule",
		Slug:            "course-schedule",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Topological Sort", "BFS", "DFS", "Cycle Detection"},
		RelatedChapters: []int{13},
		Description: `There are **numCourses** courses, labelled 0 to numCourses - 1. Each pair [a, b] in **prerequisites** means you must finish course b before you can take course a. Prerequisites are one-way (the graph is **directed**).

Return true if there is some order in which you can take every course, and false otherwise. With no prerequisites at all, the answer is true.`,
		Constraints: []string{
			"1 <= numCourses <= 2000",
			"0 <= prerequisites.length <= 5000",
			"prerequisites[i].length == 2 and 0 <= a, b < numCourses",
			"All pairs are distinct",
		},
		Examples: []Example{
			{Input: "numCourses = 2, prerequisites = [[1,0]]", Output: "true", Explanation: "Take course 0, then course 1."},
			{Input: "numCourses = 2, prerequisites = [[1,0],[0,1]]", Output: "false", Explanation: "Course 1 needs course 0 and course 0 needs course 1. Neither can be taken first."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"numCourses": 2, "prerequisites": [][]int{{1, 0}}}, Expected: true, Hidden: false, Note: "Example 1: take 0, then 1"},
			{Input: map[string]interface{}{"numCourses": 2, "prerequisites": [][]int{{1, 0}, {0, 1}}}, Expected: false, Hidden: false, Note: "Example 2: two courses need each other"},
			{Input: map[string]interface{}{"numCourses": 1, "prerequisites": [][]int{}}, Expected: true, Hidden: true, Note: "one course, no prerequisites"},
			{Input: map[string]interface{}{"numCourses": 3, "prerequisites": [][]int{{0, 1}, {1, 2}, {2, 0}}}, Expected: false, Hidden: true, Note: "a 3-cycle"},
			{Input: map[string]interface{}{"numCourses": 4, "prerequisites": [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, Expected: true, Hidden: true, Note: "diamond: course 3 is reached twice, which is not a cycle"},
			{Input: map[string]interface{}{"numCourses": 5, "prerequisites": [][]int{{1, 0}, {2, 1}, {4, 3}, {3, 4}}}, Expected: false, Hidden: true, Note: "the cycle is in a separate component"},
			{Input: map[string]interface{}{"numCourses": 6, "prerequisites": [][]int{}}, Expected: true, Hidden: true, Note: "no prerequisites at all"},
		},
		TimeComplexity:  "O(V + E)",
		SpaceComplexity: "O(V + E)",
		StarterCode:     "def canFinish(numCourses, prerequisites):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Which courses can you take on day one? Once you have taken them, which courses become available? When would this process get stuck before every course is taken?"},
			{Level: 2, Type: "insight", Content: "Draw an arrow b -> a for each pair [a, b]. You can finish everything exactly when this directed graph has no cycle. Kahn's algorithm finds out: repeatedly take a course whose in-degree (number of unmet prerequisites) is 0 and decrement the in-degree of the courses it unlocks. If a cycle exists, its courses never reach in-degree 0, so fewer than numCourses get taken. (A three-colour DFS that looks for an edge back into the current path works too.)", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Build the graph: for each [a, b], add the edge b -> a.\n2. Count in-degrees: indegree[a] += 1 for each pair.\n3. Seed the frontier: queue every course with in-degree 0.\n4. Expand one node: pop a course, count it as taken, decrement the in-degree of each course it unlocks, and queue any that reach 0.\n5. Detect a cycle: the answer is true exactly when the number taken equals numCourses."},
			{Level: 4, Type: "skeleton", Content: `from collections import deque

def canFinish(numCourses, prerequisites):
    unlocks = [[] for _ in range(numCourses)]
    indegree = [0] * numCourses
    for course, prereq in prerequisites:
        ____
        ____

    ready = deque(c for c in range(numCourses) if ____)
    taken = 0
    while ready:
        current = ready.popleft()
        taken += 1
        for nxt in unlocks[current]:
            ____
            if indegree[nxt] == 0:
                ready.append(nxt)
    return ____`},
		},
		Solution: Solution{
			Code: `from collections import deque

def canFinish(numCourses, prerequisites):
    unlocks = [[] for _ in range(numCourses)]  # edge prereq -> course
    indegree = [0] * numCourses  # prerequisites still unmet, per course
    for course, prereq in prerequisites:
        unlocks[prereq].append(course)
        indegree[course] += 1

    ready = deque(c for c in range(numCourses) if indegree[c] == 0)
    taken = 0
    while ready:
        current = ready.popleft()
        taken += 1
        for nxt in unlocks[current]:
            indegree[nxt] -= 1
            if indegree[nxt] == 0:  # its last prerequisite just got done
                ready.append(nxt)
    return taken == numCourses  # anything left over is stuck on a cycle`,
			Explanation: "The earlier problems in this ladder were undirected; here the edges have a direction (prerequisite -> course), and the question is whether the directed graph has a cycle. Kahn's algorithm answers it with the same queue loop as Walls and Gates, plus a new rule for when a node may join the queue: simulate the semester by keeping a queue of courses with no unmet prerequisites, taking them one by one, and unlocking the courses that depended on them. A cycle shows up as courses that never become available.",
			BruteForce:  "Try orders: generate permutations of the courses and check each one against every prerequisite. With n courses that is n! orders, hopeless beyond about 10. A less silly brute force runs a fresh DFS from every course to see if it can come back to itself, O(V * (V + E)), repeating the same exploration from every start.",
			KeyInsight:  "A course with in-degree 0 can always be taken next, and taking it can only lower other in-degrees. So greedily taking available courses never blocks you; the only thing that can stop the process early is a cycle, whose courses all keep at least one unmet prerequisite from inside the cycle.",
			WhyItWorks:  "If the graph has no cycle, there is always some course with in-degree 0 among those not yet taken (follow prerequisites backwards from any remaining course; without a cycle you must reach one with none left). Taking it and decrementing its successors keeps that true, so Kahn's loop takes all numCourses courses. If there is a cycle, no course on it can ever reach in-degree 0: each has a prerequisite on the cycle, and that one can only be taken after its own prerequisite on the cycle, and so on around the loop. So the count stops short of numCourses. Each course is queued once and each edge is decremented once: O(V + E).",
			CommonMistakes: []string{
				"Reversing the edge: [a, b] means b must come first, so the edge is b -> a and indegree[a] grows. The other direction answers a different question and fails asymmetric inputs.",
				"In a DFS version, using a single visited set as the cycle test. In the diamond 0->1->3, 0->2->3, course 3 is seen twice without any cycle. You need three states: unvisited, on the current path, finished.",
				"Checking only the component of course 0. The cycle may be in a separate group of courses; Kahn's count over all courses, or a DFS started from every unvisited course, catches it.",
				"Forgetting that no prerequisites at all means every course can be taken: return true.",
			},
			TimeComplexity:  "O(V + E) where V = numCourses and E = len(prerequisites).",
			SpaceComplexity: "O(V + E) for the adjacency lists and the queue.",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the graph", Explanation: "Each pair [course, prereq] becomes an edge prereq -> course: taking prereq unlocks course.", CodeSnippet: "for course, prereq in prerequisites:\n    unlocks[prereq].append(course)", LineStart: 6, LineEnd: 7},
				{Title: "Count in-degrees", Explanation: "indegree[course] is the number of prerequisites of course not yet taken.", CodeSnippet: "indegree[course] += 1", LineStart: 8, LineEnd: 8},
				{Title: "Seed the frontier", Explanation: "Courses with no prerequisites can be taken immediately.", CodeSnippet: "ready = deque(c for c in range(numCourses) if indegree[c] == 0)", LineStart: 10, LineEnd: 10},
				{Title: "Expand one node", Explanation: "Taking a course satisfies one prerequisite of everything it unlocks. A course whose count hits 0 has had its last prerequisite taken, so it joins the queue.", CodeSnippet: "current = ready.popleft()\ntaken += 1\nfor nxt in unlocks[current]:\n    indegree[nxt] -= 1\n    if indegree[nxt] == 0:\n        ready.append(nxt)", LineStart: 13, LineEnd: 18},
				{Title: "Detect a cycle", Explanation: "If some courses were never taken, they are stuck waiting on each other: a cycle.", CodeSnippet: "return taken == numCourses", LineStart: 19, LineEnd: 19},
			},
			Alternatives: []AltSolution{
				{
					Name: "Three-colour DFS",
					Code: `def canFinish(numCourses, prerequisites):
    needs = [[] for _ in range(numCourses)]  # edge course -> prereq
    for course, prereq in prerequisites:
        needs[course].append(prereq)
    WHITE, GRAY, BLACK = 0, 1, 2  # unvisited, on the current path, finished
    color = [WHITE] * numCourses

    def has_cycle(course):
        color[course] = GRAY
        for prereq in needs[course]:
            if color[prereq] == GRAY:
                return True  # back edge: we reached our own path again
            if color[prereq] == WHITE and has_cycle(prereq):
                return True
        color[course] = BLACK
        return False

    for course in range(numCourses):
        if color[course] == WHITE and has_cycle(course):
            return False
    return True`,
					TimeComplexity:  "O(V + E)",
					SpaceComplexity: "O(V + E)",
					Explanation:     "Colour each course WHITE (not visited), GRAY (on the current DFS path) or BLACK (fully explored, known to be cycle-free). Following prerequisites from a GRAY course into another GRAY course means we have walked back into our own path: a cycle. Reaching a BLACK course is fine; that is how the diamond case avoids a false alarm.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "In the three-colour DFS, why is reaching a BLACK node not a cycle but reaching a GRAY node is?", Answer: "GRAY nodes are exactly the ones on the current recursion path, so an edge to one closes a loop back to an ancestor: a cycle. A BLACK node was fully explored earlier and every path out of it was shown cycle-free; reaching it again just means two routes share a descendant, like course 3 in the diamond."},
			{Prompt: "Why does Kahn's algorithm never take a course on a cycle?", Answer: "A course on a cycle has a prerequisite that is also on the cycle. That prerequisite can only reach in-degree 0 after its own prerequisite on the cycle is taken, and so on around the loop back to the first course. No course on the cycle can go first, so none of their in-degrees ever reach 0."},
		},
	},
	{
		ID:              "course-schedule-ii",
		Number:          210,
		Order:           8,
		Title:           "Course Schedule II",
		Slug:            "course-schedule-ii",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Topological Sort", "BFS", "DFS"},
		RelatedChapters: []int{13},
		Description: `There are **numCourses** courses, labelled 0 to numCourses - 1. Each pair [a, b] in **prerequisites** means course b must be taken before course a (a **directed** dependency).

Return an order in which all the courses can be taken: a list containing every course exactly once, where each course appears after all of its prerequisites. If several orders work, return any of them. If no order exists (the prerequisites contain a cycle), return an empty list.`,
		Constraints: []string{
			"1 <= numCourses <= 2000",
			"0 <= prerequisites.length <= numCourses * (numCourses - 1)",
			"prerequisites[i].length == 2, 0 <= a, b < numCourses and a != b",
			"All pairs are distinct",
		},
		Examples: []Example{
			{Input: "numCourses = 2, prerequisites = [[1,0]]", Output: "[0,1]", Explanation: "Course 0 has no prerequisites; course 1 needs 0."},
			{Input: "numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]", Output: "[0,1,2,3]", Explanation: "Course 0 first, then 1 and 2 in either order, then 3. [0,2,1,3] is also accepted."},
			{Input: "numCourses = 1, prerequisites = []", Output: "[0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"numCourses": 2, "prerequisites": [][]int{{1, 0}}}, Expected: []int{0, 1}, Hidden: false, Note: "Example 1: the only order is 0, 1"},
			{Input: map[string]interface{}{"numCourses": 4, "prerequisites": [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, Expected: []int{0, 1, 2, 3}, Hidden: false, Note: "Example 2: 1 and 2 may come in either order"},
			{Input: map[string]interface{}{"numCourses": 1, "prerequisites": [][]int{}}, Expected: []int{0}, Hidden: false, Note: "Example 3: a single course"},
			{Input: map[string]interface{}{"numCourses": 2, "prerequisites": [][]int{{0, 1}, {1, 0}}}, Expected: []int{}, Hidden: true, Note: "a cycle: return []"},
			{Input: map[string]interface{}{"numCourses": 3, "prerequisites": [][]int{}}, Expected: []int{0, 1, 2}, Hidden: true, Note: "no prerequisites: any order of all three"},
			{Input: map[string]interface{}{"numCourses": 6, "prerequisites": [][]int{{5, 4}, {4, 3}, {3, 2}, {2, 1}, {1, 0}}}, Expected: []int{0, 1, 2, 3, 4, 5}, Hidden: true, Note: "a chain allows exactly one order"},
			{Input: map[string]interface{}{"numCourses": 5, "prerequisites": [][]int{{1, 0}, {2, 1}, {0, 2}, {4, 3}}}, Expected: []int{}, Hidden: true, Note: "cycle plus a free part: return [], not a partial order"},
		},
		TimeComplexity:  "O(V + E)",
		SpaceComplexity: "O(V + E)",
		StarterCode:     "def findOrder(numCourses, prerequisites):\n    # Write your solution here\n    pass",
		Harness: Harness{
			Validator: `def check(raw, args, result, expected):
    n = raw["numCourses"]
    if expected == []:
        return result == [], "the prerequisites contain a cycle, so the answer is []"
    if not isinstance(result, list) or sorted(result) != list(range(n)):
        return False, "the order must list every course 0..%d exactly once" % (n - 1)
    position = {course: i for i, course in enumerate(result)}
    for course, prereq in raw["prerequisites"]:
        if position[prereq] > position[course]:
            return False, "course %d is taken before its prerequisite %d" % (course, prereq)
    return True`,
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In Course Schedule you counted how many courses Kahn's algorithm could take. What if you wrote down each course at the moment you took it?"},
			{Level: 2, Type: "insight", Content: "The order in which Kahn's algorithm removes courses is a topological order: a course is only removed after all its prerequisites were removed. Record the removals; if fewer than numCourses were removed, there is a cycle and the answer is [].", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Build the graph and Count in-degrees exactly as in Course Schedule.\n2. Seed the frontier: queue every course with in-degree 0.\n3. Expand one node: pop a course, append it to the order, decrement its successors and queue those that reach 0.\n4. Detect a cycle: return the order if it has numCourses entries, otherwise []."},
			{Level: 4, Type: "skeleton", Content: `from collections import deque

def findOrder(numCourses, prerequisites):
    unlocks = [[] for _ in range(numCourses)]
    indegree = [0] * numCourses
    for course, prereq in prerequisites:
        unlocks[prereq].append(course)
        indegree[course] += 1

    ready = deque(c for c in range(numCourses) if indegree[c] == 0)
    order = []
    while ready:
        current = ready.popleft()
        ____
        for nxt in unlocks[current]:
            ____
            if ____:
                ready.append(nxt)
    return ____`},
		},
		Solution: Solution{
			Code: `from collections import deque

def findOrder(numCourses, prerequisites):
    unlocks = [[] for _ in range(numCourses)]  # edge prereq -> course
    indegree = [0] * numCourses
    for course, prereq in prerequisites:
        unlocks[prereq].append(course)
        indegree[course] += 1

    ready = deque(c for c in range(numCourses) if indegree[c] == 0)
    order = []
    while ready:
        current = ready.popleft()
        order.append(current)  # all its prerequisites are already in order
        for nxt in unlocks[current]:
            indegree[nxt] -= 1
            if indegree[nxt] == 0:
                ready.append(nxt)
    return order if len(order) == numCourses else []`,
			Explanation: "Course Schedule asked whether an order exists; this asks for the order itself. The Kahn loop is unchanged; the only addition is recording each course when it is removed. Because many orders are usually valid, the judge does not compare with one fixed answer: it checks that your list contains every course once and that every prerequisite appears before its course.",
			BruteForce:  "Generate permutations of the courses and return the first one that satisfies every prerequisite, or [] if none does. That is O(n! * E). Another slow approach repeatedly scans all remaining courses for one whose prerequisites are all taken: O(V * (V + E)), since each scan re-checks courses that are still blocked.",
			KeyInsight:  "The removal order of Kahn's algorithm is itself a valid order, because a course is removed only when its in-degree is 0, which means every prerequisite was removed (and recorded) earlier.",
			WhyItWorks:  "Invariant: every course in order appears after all of its prerequisites. A course enters the queue only when its in-degree hits 0, which happens only after each of its prerequisites was popped and appended, so the invariant holds at each append. As in Course Schedule, the loop removes every course exactly when the graph has no cycle; otherwise the courses on (or behind) a cycle are never appended, the list is shorter than numCourses, and the correct answer is []. Each course and each edge is handled once: O(V + E).",
			CommonMistakes: []string{
				"Returning the partial order when there is a cycle. The problem asks for [] whenever not every course can be taken.",
				"In the DFS version, appending a course before visiting its prerequisites (pre-order). Append after them (post-order), and with prerequisite edges course -> prereq that post-order is already the answer; with edges prereq -> course you must reverse it.",
				"Building the edge the wrong way round, which outputs a reversed (invalid) order. For [a, b], b must come first: edge b -> a.",
			},
			TimeComplexity:  "O(V + E)",
			SpaceComplexity: "O(V + E)",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the graph", Explanation: "Edges go from prerequisite to course, and indegree counts each course's unmet prerequisites.", CodeSnippet: "unlocks[prereq].append(course)\nindegree[course] += 1", LineStart: 7, LineEnd: 8},
				{Title: "Seed the frontier", Explanation: "Courses with no prerequisites can start the order.", CodeSnippet: "ready = deque(c for c in range(numCourses) if indegree[c] == 0)\norder = []", LineStart: 10, LineEnd: 11},
				{Title: "Expand one node", Explanation: "Record the course as it is taken; all of its prerequisites are already in order. Then release the courses that were waiting only for it.", CodeSnippet: "order.append(current)\nfor nxt in unlocks[current]:\n    indegree[nxt] -= 1\n    if indegree[nxt] == 0:\n        ready.append(nxt)", LineStart: 14, LineEnd: 18},
				{Title: "Detect a cycle", Explanation: "A full-length order proves there is no cycle; a short one means some courses could never be taken.", CodeSnippet: "return order if len(order) == numCourses else []", LineStart: 19, LineEnd: 19},
			},
			Alternatives: []AltSolution{
				{
					Name: "Three-colour DFS, post-order",
					Code: `def findOrder(numCourses, prerequisites):
    needs = [[] for _ in range(numCourses)]  # edge course -> prereq
    for course, prereq in prerequisites:
        needs[course].append(prereq)
    WHITE, GRAY, BLACK = 0, 1, 2
    color = [WHITE] * numCourses
    order = []

    def visit(course):
        color[course] = GRAY
        for prereq in needs[course]:
            if color[prereq] == GRAY:
                return False  # cycle
            if color[prereq] == WHITE and not visit(prereq):
                return False
        color[course] = BLACK
        order.append(course)  # post-order: every prereq was appended first
        return True

    for course in range(numCourses):
        if color[course] == WHITE and not visit(course):
            return []
    return order`,
					TimeComplexity:  "O(V + E)",
					SpaceComplexity: "O(V + E)",
					Explanation:     "DFS along course -> prerequisite edges and append a course only after all its prerequisites have been appended (post-order). That list is already a valid order. A GRAY-to-GRAY edge means a cycle, so return [].",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can the judge not just compare your answer with one expected list?", Answer: "Whenever two courses don't depend on each other, either can come first, so there are often many valid orders (any permutation when there are no prerequisites). The validator checks the defining property instead: every course appears once, and for each [a, b], b comes before a."},
			{Prompt: "Kahn's algorithm uses a FIFO queue. Would a stack (LIFO) also give a valid order?", Answer: "Yes. Any course with in-degree 0 may be taken next; the queue discipline only decides which of the available courses goes first. A stack, a heap (giving the smallest-numbered available course) or any other choice still respects every prerequisite."},
		},
	},
	{
		ID:              "number-of-connected-components",
		Number:          323,
		Order:           9,
		Title:           "Number of Connected Components in an Undirected Graph",
		Slug:            "number-of-connected-components",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Union-Find", "DFS", "Connected Components"},
		RelatedChapters: []int{13},
		Description: `You are given **n** nodes labelled 0 to n - 1 and a list **edges**, where each [a, b] is an **undirected** edge joining a and b.

A **connected component** is a largest group of nodes in which every node can reach every other by following edges. A node with no edges is a component on its own.

Return the number of connected components.`,
		Constraints: []string{
			"1 <= n <= 2000",
			"0 <= edges.length <= 5000",
			"edges[i].length == 2, 0 <= a, b < n and a != b",
			"No edge appears twice",
		},
		Examples: []Example{
			{Input: "n = 5, edges = [[0,1],[1,2],[3,4]]", Output: "2", Explanation: "Nodes 0, 1 and 2 are connected through 1; nodes 3 and 4 form a second group."},
			{Input: "n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]", Output: "1", Explanation: "One chain connects all five nodes."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 5, "edges": [][]int{{0, 1}, {1, 2}, {3, 4}}}, Expected: 2, Hidden: false, Note: "Example 1: {0,1,2} and {3,4}"},
			{Input: map[string]interface{}{"n": 5, "edges": [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}}, Expected: 1, Hidden: false, Note: "Example 2: one chain"},
			{Input: map[string]interface{}{"n": 1, "edges": [][]int{}}, Expected: 1, Hidden: true, Note: "a single node"},
			{Input: map[string]interface{}{"n": 4, "edges": [][]int{}}, Expected: 4, Hidden: true, Note: "no edges: every node is its own component"},
			{Input: map[string]interface{}{"n": 3, "edges": [][]int{{0, 1}, {1, 2}, {2, 0}}}, Expected: 1, Hidden: true, Note: "cycle: the third edge joins nodes already connected"},
			{Input: map[string]interface{}{"n": 6, "edges": [][]int{{0, 1}, {2, 3}, {4, 5}, {1, 3}, {3, 5}}}, Expected: 1, Hidden: true, Note: "pairs merged into one by later edges"},
			{Input: map[string]interface{}{"n": 7, "edges": [][]int{{0, 1}, {1, 2}, {0, 2}, {4, 5}}}, Expected: 4, Hidden: true, Note: "isolated nodes 3 and 6 count too"},
		},
		TimeComplexity:  "O(n + E * α(n))",
		SpaceComplexity: "O(n)",
		StarterCode:     "def countComponents(n, edges):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Start with n separate nodes, so n groups. Each edge might glue two groups into one. When does adding an edge lower the number of groups, and when does it not?"},
			{Level: 2, Type: "insight", Content: "An edge reduces the count by one exactly when its endpoints are in different groups. Union-find (disjoint set union) answers 'same group?' quickly: each group is a tree stored in a parent array, find(x) walks up to the root, and two nodes are in the same group when their roots match. Union joins two roots. (Counting DFS floods, as in Number of Islands, also works.)", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Start with singletons: parent[x] = x and size[x] = 1 for every node; components = n.\n2. Find the roots of both endpoints of each edge, compressing the path as you go.\n3. Union or skip: if the roots are equal, the edge adds nothing; otherwise hang the smaller tree under the larger and subtract one component.\n4. Read off the answer: the remaining count."},
			{Level: 4, Type: "skeleton", Content: `def countComponents(n, edges):
    parent = list(range(n))
    size = [1] * n

    def find(x):
        if parent[x] != x:
            parent[x] = ____
        return parent[x]

    components = n
    for a, b in edges:
        root_a, root_b = find(a), find(b)
        if ____:
            continue
        if size[root_a] < size[root_b]:
            root_a, root_b = root_b, root_a
        ____  # attach the smaller root under the larger
        size[root_a] += size[root_b]
        ____
    return components`},
		},
		Solution: Solution{
			Code: `def countComponents(n, edges):
    parent = list(range(n))  # every node starts as the root of its own set
    size = [1] * n

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])  # path compression: point straight at the root
        return parent[x]

    components = n
    for a, b in edges:
        root_a, root_b = find(a), find(b)
        if root_a == root_b:
            continue  # already in one set: this edge merges nothing
        if size[root_a] < size[root_b]:
            root_a, root_b = root_b, root_a
        parent[root_b] = root_a  # union by size: hang the smaller tree under the bigger
        size[root_a] += size[root_b]
        components -= 1
    return components`,
			Explanation: "Number of Islands counted components by flooding; this problem introduces a second tool, union-find, that counts them edge by edge without building an adjacency list. Begin with n one-node groups. For each edge, find the root of each endpoint; if they differ, merge the groups and lower the count. Path compression and union by size keep every find almost constant time.",
			BruteForce:  "For every pair of nodes, run a search to test whether they are connected, then group them: O(n^2 * (n + E)). Even counting with union-find but without its two optimisations can degrade: always hanging the first root under the second can build a chain of length n, making each find O(n) and the whole run O(n * E).",
			KeyInsight:  "Components only ever merge as edges are added, never split. A structure that supports 'which group is x in?' and 'merge these two groups' is all you need, and union-find does both in nearly O(1) amortised time by storing each group as a shallow tree identified by its root.",
			WhyItWorks:  "Invariant: two nodes have the same root exactly when the edges processed so far connect them, and components equals the number of distinct roots. It holds at the start (n roots, no edges). When an edge joins different roots, setting parent[root_b] = root_a merges exactly those two groups into one, so the count drops by one; when the roots are equal the edge connects nodes already connected, so nothing changes. Path compression only re-points nodes to their own root, so it never changes which group a node is in. Union by size keeps every tree's height at most log2(n), because a node's depth grows only when its tree is merged under one at least as large, which at least doubles its tree's size. With both tricks each operation costs O(α(n)) amortised, where α is the inverse Ackermann function, below 5 for any input that fits in memory.",
			CommonMistakes: []string{
				"Decrementing the count for every edge. An edge between nodes already connected (the third edge of a triangle) merges nothing; compare the roots first.",
				"Linking the nodes instead of their roots: parent[b] = a. If b is not a root, its old root and the rest of its tree are left behind, and later finds give wrong answers. Always link root to root.",
				"Comparing parent[a] == parent[b] instead of find(a) == find(b). Parents are only the next step up the tree, not the group's identity.",
				"Forgetting isolated nodes. Start the count at n, not at the number of nodes that appear in edges.",
			},
			TimeComplexity:  "O(n + E * α(n)), effectively linear: α(n) is the inverse Ackermann function and is at most 4 in practice.",
			SpaceComplexity: "O(n) for the parent and size arrays.",
			Walkthrough: []WalkthroughStep{
				{Title: "Start with singletons", Explanation: "Each node is the root of its own one-node tree, so there are n components before any edge is read.", CodeSnippet: "parent = list(range(n))\nsize = [1] * n", LineStart: 2, LineEnd: 3},
				{Title: "Find the roots", Explanation: "Walk up parent links to the root, the group's identity. On the way back, point each node directly at the root (path compression), so the next find from any of them takes one step.", CodeSnippet: "if parent[x] != x:\n    parent[x] = find(parent[x])\nreturn parent[x]", LineStart: 6, LineEnd: 8},
				{Title: "Union or skip", Explanation: "Same root means already connected: skip. Otherwise hang the smaller tree under the larger root (union by size) so trees stay shallow, and one fewer component remains.", CodeSnippet: "if root_a == root_b:\n    continue\nif size[root_a] < size[root_b]:\n    root_a, root_b = root_b, root_a\nparent[root_b] = root_a\nsize[root_a] += size[root_b]\ncomponents -= 1", LineStart: 13, LineEnd: 19},
				{Title: "Read off the answer", Explanation: "Every successful union removed one component, so what remains is the number of groups.", CodeSnippet: "return components", LineStart: 20, LineEnd: 20},
			},
			Alternatives: []AltSolution{
				{
					Name: "Count DFS floods",
					Code: `def countComponents(n, edges):
    graph = [[] for _ in range(n)]
    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)
    seen = [False] * n
    components = 0
    for start in range(n):
        if seen[start]:
            continue
        components += 1
        seen[start] = True
        stack = [start]
        while stack:
            node = stack.pop()
            for nxt in graph[node]:
                if not seen[nxt]:
                    seen[nxt] = True
                    stack.append(nxt)
    return components`,
					TimeComplexity:  "O(n + E)",
					SpaceComplexity: "O(n + E)",
					Explanation:     "Exactly the Number of Islands plan on an explicit graph: build adjacency lists, scan for an unvisited start, flood its component with an iterative DFS, and count the floods. Union-find wins when edges arrive one at a time (you can answer after each edge without re-flooding) and when you need to detect the edge that creates a cycle, as in the next two problems.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must union attach root to root, not node to node?", Answer: "A group is identified by its root. If you set parent[b] = a while b is not a root, b leaves its old tree but the rest of that tree still points at the old root, so the group is split instead of merged. Linking find(b) under find(a) moves the whole tree at once."},
			{Prompt: "What does union by size buy you, given that path compression already flattens trees?", Answer: "Path compression fixes trees after a find walks them, but without balancing the first walk can still be long: linking always the same way can build a chain of n nodes. Union by size guarantees height at most log2(n) from the start, and together the two give near-constant amortised time per operation."},
		},
	},
	{
		ID:              "redundant-connection",
		Number:          684,
		Order:           10,
		Title:           "Redundant Connection",
		Slug:            "redundant-connection",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Union-Find", "Cycle Detection"},
		RelatedChapters: []int{13},
		Description: `A **tree** is an undirected graph that is connected and has no cycles. Start with a tree on n nodes labelled 1 to n, and add one extra **undirected** edge between two different nodes that were not already joined. The result is given as **edges**, a list of n pairs [a, b].

Return an edge you can remove so that the remaining n - 1 edges form a tree again. If several edges would work, return the one that appears **last** in edges.`,
		Constraints: []string{
			"n == edges.length",
			"3 <= n <= 1000",
			"edges[i].length == 2, 1 <= a < b <= n",
			"No repeated edges; the graph is connected",
		},
		Examples: []Example{
			{Input: "edges = [[1,2],[1,3],[2,3]]", Output: "[2,3]", Explanation: "The three edges form a triangle. Removing any one leaves a tree; [2,3] is the last of them in the list."},
			{Input: "edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]", Output: "[1,4]", Explanation: "The cycle is 1-2-3-4-1. Removing [1,5] would disconnect node 5, so only the cycle edges qualify, and [1,4] is the last of those."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"edges": [][]int{{1, 2}, {1, 3}, {2, 3}}}, Expected: []int{2, 3}, Hidden: false, Note: "Example 1: triangle"},
			{Input: map[string]interface{}{"edges": [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}}, Expected: []int{1, 4}, Hidden: false, Note: "Example 2: the cycle closes at [1,4]; [1,5] comes later but is not on the cycle"},
			{Input: map[string]interface{}{"edges": [][]int{{1, 3}, {3, 4}, {1, 5}, {3, 5}, {2, 3}}}, Expected: []int{3, 5}, Hidden: true, Note: "cycle 1-3-5 closed by [3,5]"},
			{Input: map[string]interface{}{"edges": [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}}, Expected: []int{4, 1}, Hidden: true, Note: "the very last edge closes the cycle"},
			{Input: map[string]interface{}{"edges": [][]int{{2, 3}, {1, 2}, {3, 4}, {4, 5}, {5, 3}}}, Expected: []int{5, 3}, Hidden: true, Note: "cycle 3-4-5 that does not use the first edge"},
			{Input: map[string]interface{}{"edges": [][]int{{3, 4}, {1, 2}, {2, 4}, {3, 5}, {2, 5}}}, Expected: []int{2, 5}, Hidden: true, Note: "long cycle 2-4-3-5"},
		},
		TimeComplexity:  "O(n * α(n))",
		SpaceComplexity: "O(n)",
		StarterCode:     "def findRedundantConnection(edges):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Add the edges one at a time to an empty graph. As long as you are building a tree, each edge joins two separate pieces. What is special about the moment an edge joins two nodes that are already connected?"},
			{Level: 2, Type: "insight", Content: "Union-find: process edges in order; the first edge whose endpoints already have the same root closes the only cycle. Every earlier edge was needed to connect its endpoints, and this edge is the last cycle edge in the input, which is exactly the one the problem asks for.", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Start with singletons: parent[x] = x for nodes 1..n.\n2. Find the roots of both endpoints of the next edge.\n3. Union or detect a cycle: equal roots mean the endpoints were already connected, so return this edge; otherwise merge the two trees."},
			{Level: 4, Type: "skeleton", Content: `def findRedundantConnection(edges):
    n = len(edges)
    parent = list(range(n + 1))

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    for a, b in edges:
        root_a, root_b = find(a), find(b)
        if ____:
            return ____
        ____`},
		},
		Solution: Solution{
			Code: `def findRedundantConnection(edges):
    n = len(edges)
    parent = list(range(n + 1))  # nodes are labelled 1..n
    size = [1] * (n + 1)

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    for a, b in edges:
        root_a, root_b = find(a), find(b)
        if root_a == root_b:
            return [a, b]  # a and b were already connected: this edge closes the cycle
        if size[root_a] < size[root_b]:
            root_a, root_b = root_b, root_a
        parent[root_b] = root_a
        size[root_a] += size[root_b]
    return []`,
			Explanation: "This reuses the union-find from Number of Connected Components; the skipped case there (both endpoints already in one group) is now the answer. Adding the edges in order, every edge merges two groups until one edge finds its endpoints already connected. That edge closes the cycle, and it is the latest cycle edge in the input.",
			BruteForce:  "Try removing each edge, starting from the last, and check with a DFS whether the remaining n - 1 edges are connected (n - 1 connected edges on n nodes form a tree). That is the alternative below: O(n) checks of O(n) each, O(n^2) in total, rebuilding the graph every time.",
			KeyInsight:  "Adding edges in input order, the first one whose endpoints are already connected must be on the cycle, and every other cycle edge came before it. So it is the last cycle edge, which is exactly the tie-break the problem asks for.",
			WhyItWorks:  "The graph is a tree plus one edge, so it has exactly one cycle, and the removable edges are exactly the cycle edges (removing any of them leaves n - 1 edges that still connect everything; removing a non-cycle edge disconnects the graph). Union-find invariant: two nodes share a root exactly when the edges processed so far connect them. Before the last cycle edge is processed, the cycle edges seen so far form a path, not a loop, so no earlier edge finds equal roots (no edge outside the cycle can close a loop either). When the last cycle edge [a, b] arrives, the rest of the cycle already connects a and b, so their roots are equal and we return it. Each edge costs two near-constant finds: O(n * α(n)).",
			CommonMistakes: []string{
				"Returning the first edge of the cycle instead of the last one. The edge that triggers the equal-roots test is the last cycle edge, which is what the problem wants; don't try to find the cycle and pick an edge from it yourself.",
				"Sizing parent as n instead of n + 1. Nodes are labelled 1..n, so index n must exist.",
				"Merging before testing: if you union first, the roots are always equal afterwards and you return the first edge.",
			},
			TimeComplexity:  "O(n * α(n)) with path compression and union by size.",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Start with singletons", Explanation: "Labels run from 1 to n, so the arrays have n + 1 slots and slot 0 is unused.", CodeSnippet: "parent = list(range(n + 1))\nsize = [1] * (n + 1)", LineStart: 3, LineEnd: 4},
				{Title: "Find the roots", Explanation: "The same find with path compression as in Number of Connected Components.", CodeSnippet: "if parent[x] != x:\n    parent[x] = find(parent[x])\nreturn parent[x]", LineStart: 7, LineEnd: 9},
				{Title: "Union or detect a cycle", Explanation: "Equal roots: a and b are already connected by earlier edges, so this edge closes the cycle and is the answer. Otherwise merge the smaller tree into the larger and continue.", CodeSnippet: "if root_a == root_b:\n    return [a, b]\nif size[root_a] < size[root_b]:\n    root_a, root_b = root_b, root_a\nparent[root_b] = root_a\nsize[root_a] += size[root_b]", LineStart: 13, LineEnd: 18},
			},
			Alternatives: []AltSolution{
				{
					Name: "Remove each edge and test for a tree",
					Code: `def findRedundantConnection(edges):
    n = len(edges)

    def is_tree_without(skip):
        graph = [[] for _ in range(n + 1)]
        for i, (a, b) in enumerate(edges):
            if i != skip:
                graph[a].append(b)
                graph[b].append(a)
        seen = {1}
        stack = [1]
        while stack:
            node = stack.pop()
            for nxt in graph[node]:
                if nxt not in seen:
                    seen.add(nxt)
                    stack.append(nxt)
        return len(seen) == n  # n - 1 edges and connected means a tree

    for i in range(n - 1, -1, -1):  # try the last edge first
        if is_tree_without(i):
            return edges[i]
    return []`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n)",
					Explanation:     "Straight from the definition: for each edge from last to first, drop it, flood from node 1 and check that all n nodes are reached. The first edge that passes is the answer. It is independent of union-find, which makes it a good check, but it rebuilds the graph for each candidate.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the edge returned by union-find always the one that appears last among the cycle edges?", Answer: "Union-find only reports an edge when its endpoints are already connected. Along the cycle, that cannot happen until every other cycle edge has been added (until then the cycle edges form a path, which leaves the two ends of the missing edge unconnected). So the reported edge is the cycle edge with the largest index."},
			{Prompt: "What changes if the graph were directed (LeetCode's Redundant Connection II)?", Answer: "Union-find alone no longer suffices: a directed tree also requires every node except the root to have exactly one parent, so the extra edge may create a node with two parents instead of (or as well as) a cycle. You first look for a node with in-degree 2, then use union-find to decide which of its two incoming edges to drop."},
		},
	},
	{
		ID:              "graph-valid-tree",
		Number:          261,
		Order:           11,
		Title:           "Graph Valid Tree",
		Slug:            "graph-valid-tree",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Union-Find", "DFS", "Cycle Detection"},
		RelatedChapters: []int{13},
		Description: `You are given **n** nodes labelled 0 to n - 1 and a list **edges** of **undirected** edges [a, b].

Return true if these edges form a valid **tree**: every node can reach every other node (the graph is connected) and there is no cycle. Otherwise return false. A single node with no edges is a tree.`,
		Constraints: []string{
			"1 <= n <= 2000",
			"0 <= edges.length <= 5000",
			"edges[i].length == 2, 0 <= a, b < n and a != b",
			"No self-loops and no repeated edges",
		},
		Examples: []Example{
			{Input: "n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]", Output: "true", Explanation: "Four edges connect all five nodes and there is no loop."},
			{Input: "n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]", Output: "false", Explanation: "1-2-3-1 is a cycle."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 5, "edges": [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}}, Expected: true, Hidden: false, Note: "Example 1: a tree"},
			{Input: map[string]interface{}{"n": 5, "edges": [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}}, Expected: false, Hidden: false, Note: "Example 2: cycle 1-2-3"},
			{Input: map[string]interface{}{"n": 1, "edges": [][]int{}}, Expected: true, Hidden: true, Note: "one node, no edges: a tree"},
			{Input: map[string]interface{}{"n": 2, "edges": [][]int{}}, Expected: false, Hidden: true, Note: "two nodes, no edge: not connected"},
			{Input: map[string]interface{}{"n": 4, "edges": [][]int{{0, 1}, {1, 2}, {2, 0}}}, Expected: false, Hidden: true, Note: "right edge count (n - 1) but a cycle and a lonely node 3"},
			{Input: map[string]interface{}{"n": 4, "edges": [][]int{{0, 1}, {2, 3}}}, Expected: false, Hidden: true, Note: "too few edges: two pieces"},
			{Input: map[string]interface{}{"n": 3, "edges": [][]int{{1, 0}, {2, 0}}}, Expected: true, Hidden: true, Note: "a small star"},
		},
		TimeComplexity:  "O(n * α(n))",
		SpaceComplexity: "O(n)",
		StarterCode:     "def validTree(n, edges):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "How many edges does a tree on n nodes have? If the count is right, is that enough on its own, or can you think of n - 1 edges that still fail?"},
			{Level: 2, Type: "insight", Content: "A graph on n nodes is a tree exactly when it has n - 1 edges and no cycle (or, equivalently, n - 1 edges and is connected). Check the count first, then use union-find from Redundant Connection: if any edge joins two nodes that already share a root, there is a cycle.", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Check the edge count: anything other than n - 1 edges cannot be a tree.\n2. Start with singletons: parent[x] = x.\n3. Find the roots of both endpoints of each edge.\n4. Union or detect a cycle: equal roots mean a cycle, so return false; otherwise merge.\n5. Read off the answer: n - 1 edges with no cycle means connected, so return true."},
			{Level: 4, Type: "skeleton", Content: `def validTree(n, edges):
    if ____:
        return False
    parent = list(range(n))

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    for a, b in edges:
        root_a, root_b = find(a), find(b)
        if ____:
            return False
        ____
    return ____`},
		},
		Solution: Solution{
			Code: `def validTree(n, edges):
    if len(edges) != n - 1:
        return False  # a tree on n nodes has exactly n - 1 edges
    parent = list(range(n))

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    for a, b in edges:
        root_a, root_b = find(a), find(b)
        if root_a == root_b:
            return False  # a and b already connected: this edge makes a cycle
        parent[root_b] = root_a
    return True  # n - 1 edges and no cycle means connected`,
			Explanation: "This builds on Redundant Connection (an edge between nodes with the same root closes a cycle) and Number of Connected Components (each successful union removes one component). A tree needs both properties at once; checking that there are exactly n - 1 edges first means that 'no cycle' alone proves the rest.",
			BruteForce:  "Check each property separately with the tools you have: run a cycle-detecting DFS that remembers each node's parent, then a second traversal to see whether all n nodes are reachable. That is O(n + E) and fine, just more code. A truly naive check, testing for each edge whether removing it disconnects the graph, costs O(E * (n + E)).",
			KeyInsight:  "For a graph with n nodes, any two of 'n - 1 edges', 'connected' and 'acyclic' imply the third. So after an O(1) edge count, a single union-find pass that finds no cycle is enough.",
			WhyItWorks:  "Start with n singleton components. Each edge that joins two different roots lowers the component count by exactly one. If the count of edges is n - 1 and none of them closes a cycle, all n - 1 are successful unions, leaving n - (n - 1) = 1 component: the graph is connected, and acyclic by construction, so it is a tree. If some edge finds equal roots, its endpoints were already connected, so it creates a cycle and the graph is not a tree. If the edge count is not n - 1, it cannot be a tree: fewer edges leave at least two components, and more edges than n - 1 force at least one union to fail, which is a cycle.",
			CommonMistakes: []string{
				"Checking only the edge count. Three nodes in a triangle plus one lonely node have n - 1 = 3 edges but are neither connected nor acyclic.",
				"Checking only for cycles. Two separate trees have no cycle but are not one tree; the n - 1 check (or a final connectivity check) catches this.",
				"In a DFS version on an undirected graph, treating the edge back to the node you came from as a cycle. Every undirected edge appears in both directions; skip the parent, or use union-find, which has no such problem.",
				"Forgetting that n = 1 with no edges is a valid tree (0 = n - 1 edges).",
			},
			TimeComplexity:  "O(n * α(n)): at most n - 1 edges are processed, each with two near-constant finds.",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Check the edge count", Explanation: "Any tree on n nodes has exactly n - 1 edges, so any other count fails immediately. This also bounds the loop below.", CodeSnippet: "if len(edges) != n - 1:\n    return False", LineStart: 2, LineEnd: 3},
				{Title: "Find the roots", Explanation: "Union-find with path compression, as in the previous two problems.", CodeSnippet: "if parent[x] != x:\n    parent[x] = find(parent[x])\nreturn parent[x]", LineStart: 7, LineEnd: 9},
				{Title: "Union or detect a cycle", Explanation: "If both endpoints already share a root, this edge closes a loop. Otherwise merge their trees.", CodeSnippet: "root_a, root_b = find(a), find(b)\nif root_a == root_b:\n    return False\nparent[root_b] = root_a", LineStart: 12, LineEnd: 15},
				{Title: "Read off the answer", Explanation: "All n - 1 edges merged two components, leaving exactly one: connected and acyclic.", CodeSnippet: "return True", LineStart: 16, LineEnd: 16},
			},
			Alternatives: []AltSolution{
				{
					Name: "Edge count plus one DFS",
					Code: `def validTree(n, edges):
    if len(edges) != n - 1:
        return False
    graph = [[] for _ in range(n)]
    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)
    seen = {0}
    stack = [0]
    while stack:
        node = stack.pop()
        for nxt in graph[node]:
            if nxt not in seen:
                seen.add(nxt)
                stack.append(nxt)
    return len(seen) == n  # every node reachable from node 0`,
					TimeComplexity:  "O(n + E)",
					SpaceComplexity: "O(n + E)",
					Explanation:     "Uses the other pairing of the three tree properties: n - 1 edges and connected. One flood from node 0 checks connectivity by counting the nodes it reaches; no cycle test is needed because n - 1 edges plus connectivity already rule cycles out.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does 'n - 1 edges and no cycle' imply that the graph is connected?", Answer: "Starting from n separate nodes, an edge that does not close a cycle always joins two different components, reducing their number by one. After n - 1 such edges, n - (n - 1) = 1 component remains, so every node is connected to every other."},
			{Prompt: "This union-find omits union by size. Why is that acceptable here, and when would you add it back?", Answer: "Path compression alone already gives O(log n) amortised per operation, and n is at most 2000, so the difference is invisible. On large inputs or when the same structure answers many queries, add union by size (as in Number of Connected Components) for the near-constant α(n) bound."},
		},
	},
	{
		ID:              "word-ladder",
		Number:          127,
		Order:           12,
		Title:           "Word Ladder",
		Slug:            "word-ladder",
		Difficulty:      "Hard",
		Category:        "graphs",
		Tags:            []string{"Graph", "BFS", "Hash Table", "String", "Shortest Path"},
		RelatedChapters: []int{13, 2},
		Description: `A **ladder** from beginWord to endWord is a sequence of words that starts with beginWord and ends with endWord, where each consecutive pair differs in exactly one letter (same length, one position changed) and every word after beginWord is in **wordList**. beginWord itself does not need to be in wordList.

Return the number of words in the **shortest** ladder, counting both ends. If no ladder exists (including when endWord is not in wordList), return 0.`,
		Constraints: []string{
			"1 <= beginWord.length <= 10",
			"endWord.length == beginWord.length",
			"1 <= wordList.length <= 5000",
			"wordList[i].length == beginWord.length",
			"All words are lowercase English letters; beginWord != endWord; words in wordList are distinct",
		},
		Examples: []Example{
			{Input: `beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]`, Output: "5", Explanation: "hit -> hot -> dot -> dog -> cog is five words. No shorter ladder exists."},
			{Input: `beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]`, Output: "0", Explanation: "cog is not in the list, so no ladder can end there."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"beginWord": "hit", "endWord": "cog", "wordList": []string{"hot", "dot", "dog", "lot", "log", "cog"}}, Expected: 5, Hidden: false, Note: "Example 1: hit -> hot -> dot -> dog -> cog"},
			{Input: map[string]interface{}{"beginWord": "hit", "endWord": "cog", "wordList": []string{"hot", "dot", "dog", "lot", "log"}}, Expected: 0, Hidden: false, Note: "Example 2: endWord is not in the list"},
			{Input: map[string]interface{}{"beginWord": "a", "endWord": "c", "wordList": []string{"a", "b", "c"}}, Expected: 2, Hidden: true, Note: "one change away"},
			{Input: map[string]interface{}{"beginWord": "hot", "endWord": "dog", "wordList": []string{"hot", "dog"}}, Expected: 0, Hidden: true, Note: "two letters differ and nothing in between"},
			{Input: map[string]interface{}{"beginWord": "hot", "endWord": "dog", "wordList": []string{"hot", "dog", "dot"}}, Expected: 3, Hidden: true, Note: "one stepping stone"},
			{Input: map[string]interface{}{"beginWord": "red", "endWord": "tax", "wordList": []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}}, Expected: 4, Hidden: true, Note: "several shortest ladders of the same length"},
			{Input: map[string]interface{}{"beginWord": "lost", "endWord": "miss", "wordList": []string{"most", "mist", "miss", "lost", "fist", "fish"}}, Expected: 4, Hidden: true, Note: "four-letter words"},
		},
		TimeComplexity:  "O(N * L^2)",
		SpaceComplexity: "O(N * L^2)",
		StarterCode:     "def ladderLength(beginWord, endWord, wordList):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If every word were a node, which words would be joined by an edge? And which traversal gives the fewest steps from one node to another when every edge counts the same?"},
			{Level: 2, Type: "insight", Content: "Words are nodes, one-letter changes are edges, and the shortest ladder is a BFS shortest path. The hard part is finding neighbours fast: comparing a word with all N others costs O(N * L). Instead, group words by wildcard pattern (\"hot\" belongs to \"*ot\", \"h*t\" and \"ho*\"); two words are neighbours exactly when they share a pattern.", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Build the graph: if endWord is not in the list, return 0. Map each pattern (one letter replaced by *) to the words that match it, including beginWord.\n2. Seed the frontier: queue (beginWord, 1) and mark it seen.\n3. Expand one node: pop a word; if it is endWord, return its length. For each of its patterns, queue every unseen word in that bucket with length + 1, marking it seen as you queue it.\n4. Read off the answer: if the queue empties, return 0."},
			{Level: 4, Type: "skeleton", Content: `from collections import defaultdict, deque

def ladderLength(beginWord, endWord, wordList):
    words = set(wordList)
    if endWord not in words:
        return 0
    by_pattern = defaultdict(list)
    for word in words | {beginWord}:
        for i in range(len(word)):
            ____

    queue = deque([(beginWord, 1)])
    seen = {beginWord}
    while queue:
        word, length = queue.popleft()
        if ____:
            return length
        for i in range(len(word)):
            for nxt in ____:
                if nxt not in seen:
                    ____
    return 0`},
		},
		Solution: Solution{
			Code: `from collections import defaultdict, deque

def ladderLength(beginWord, endWord, wordList):
    words = set(wordList)
    if endWord not in words:
        return 0
    by_pattern = defaultdict(list)  # "h*t" -> every word matching it
    for word in words | {beginWord}:
        for i in range(len(word)):
            by_pattern[word[:i] + "*" + word[i + 1:]].append(word)

    queue = deque([(beginWord, 1)])  # (word, words in the sequence so far)
    seen = {beginWord}
    while queue:
        word, length = queue.popleft()
        if word == endWord:
            return length
        for i in range(len(word)):
            for nxt in by_pattern[word[:i] + "*" + word[i + 1:]]:
                if nxt not in seen:
                    seen.add(nxt)
                    queue.append((nxt, length + 1))
    return 0`,
			Explanation: "Every earlier BFS in this ladder had its graph handed to it (a grid or an edge list). Here you build the graph yourself: words are nodes and one-letter changes are edges. BFS from beginWord then finds the shortest ladder. A dictionary from wildcard patterns to words turns 'find all neighbours' from a scan of the whole list into a few dictionary lookups.",
			BruteForce:  "BFS where each step compares the current word with every unvisited word to see whether they differ in exactly one letter (the alternative below). That costs O(N * L) per word, O(N^2 * L) in total: 5000 words of length 10 means about 250 million character comparisons. Trying all 26 letters at each position (26 * L candidates, each a set lookup of cost L) is the other common approach and is fine too.",
			KeyInsight:  "Two words are one letter apart exactly when they match the same wildcard pattern, such as \"h*t\". Bucketing words by their L patterns lets each word find its neighbours by looking up L buckets instead of scanning the list.",
			WhyItWorks:  "BFS on an unweighted graph visits nodes in order of their distance from the start, so the first time endWord is popped, length is the fewest words in any ladder. Words are marked seen when queued, so each word is queued once, at its shortest distance. Neighbours are complete and correct: if two words differ only at position i, they produce the same pattern for i and are in the same bucket; if they share a bucket for position i, they agree everywhere except position i, and since the words are distinct they differ there. Building the buckets creates N * L patterns, each of length L, so O(N * L^2). The BFS builds each queued word's L patterns once (again O(N * L^2)), and each word enters the queue only once.",
			CommonMistakes: []string{
				"Returning the number of changes instead of the number of words. hit -> hot -> dot -> dog -> cog has 4 changes and 5 words; start the count at 1.",
				"Deleting words from the word set while looping over that same set, which raises 'set changed size during iteration'. Keep a separate seen set, or collect the neighbours into a list before removing them (as the alternative does).",
				"Marking words seen only when popped. The same word can then be queued from several neighbours on the same level, multiplying the work.",
				"Using DFS. It finds a ladder, not necessarily the shortest one.",
			},
			TimeComplexity:  "O(N * L^2) for N words of length L: N * L patterns, each built in O(L).",
			SpaceComplexity: "O(N * L^2) for the pattern buckets (N * L keys of length L).",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the graph", Explanation: "No ladder can end at a word outside the list. Then file every word, including beginWord, under each of its L wildcard patterns; words in one bucket are mutual neighbours.", CodeSnippet: "if endWord not in words:\n    return 0\nby_pattern = defaultdict(list)\nfor word in words | {beginWord}:\n    for i in range(len(word)):\n        by_pattern[word[:i] + \"*\" + word[i + 1:]].append(word)", LineStart: 5, LineEnd: 10},
				{Title: "Seed the frontier", Explanation: "The ladder so far is the single word beginWord, so its length is 1.", CodeSnippet: "queue = deque([(beginWord, 1)])\nseen = {beginWord}", LineStart: 12, LineEnd: 13},
				{Title: "Expand one node", Explanation: "BFS pops words in order of ladder length, so reaching endWord here means this length is the shortest. Otherwise look up the word's buckets to find its neighbours.", CodeSnippet: "word, length = queue.popleft()\nif word == endWord:\n    return length\nfor i in range(len(word)):\n    for nxt in by_pattern[word[:i] + \"*\" + word[i + 1:]]:", LineStart: 15, LineEnd: 19},
				{Title: "Mark when you discover", Explanation: "Mark a word as soon as it is queued, so it is queued only once, at its smallest length.", CodeSnippet: "if nxt not in seen:\n    seen.add(nxt)\n    queue.append((nxt, length + 1))", LineStart: 20, LineEnd: 22},
				{Title: "Read off the answer", Explanation: "If the queue empties without reaching endWord, no ladder exists.", CodeSnippet: "return 0", LineStart: 23, LineEnd: 23},
			},
			Alternatives: []AltSolution{
				{
					Name: "BFS comparing against every unvisited word",
					Code: `from collections import deque

def ladderLength(beginWord, endWord, wordList):
    if endWord not in wordList:
        return 0

    def one_apart(a, b):
        return sum(x != y for x, y in zip(a, b)) == 1

    unvisited = set(wordList)
    unvisited.discard(beginWord)
    queue = deque([(beginWord, 1)])
    while queue:
        word, length = queue.popleft()
        if word == endWord:
            return length
        for other in [w for w in unvisited if one_apart(word, w)]:
            unvisited.remove(other)
            queue.append((other, length + 1))
    return 0`,
					TimeComplexity:  "O(N^2 * L)",
					SpaceComplexity: "O(N)",
					Explanation:     "Same BFS, but neighbours are found by testing each remaining word letter by letter. Removing a word from the unvisited set when it is queued plays the role of the seen set. Simple and clearly correct, but each step scans the whole list.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the first time BFS pops endWord give the shortest ladder?", Answer: "BFS uses a FIFO queue, so it finishes every word at ladder length k before any word at length k + 1. Every edge counts as one step, so the first time endWord comes out of the queue, no ladder with fewer words can exist; it would have been popped earlier."},
			{Prompt: "What would change if each one-letter change had a different cost (say, changing a vowel costs 2)?", Answer: "The graph becomes weighted, and BFS's guarantee (fewest edges = cheapest) no longer holds. Use Dijkstra's algorithm from Chapter 14: a min-heap ordered by total cost instead of a FIFO queue, finalising a word when it is popped."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, GraphsProblems...)
}
