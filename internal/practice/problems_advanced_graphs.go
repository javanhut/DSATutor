package practice

// AdvancedGraphsProblems is the advanced-graphs ladder (Chapter 14): weighted
// shortest paths (Dijkstra, Bellman-Ford), minimum spanning trees, Eulerian
// paths and topological sort on derived graphs.
var AdvancedGraphsProblems = []*Problem{
	{
		ID:              "network-delay-time",
		Number:          743,
		Order:           1,
		Title:           "Network Delay Time",
		Slug:            "network-delay-time",
		Difficulty:      "Medium",
		Category:        "advanced-graphs",
		Tags:            []string{"Graph", "Shortest Path", "Dijkstra", "Heap"},
		RelatedChapters: []int{14, 11, 13},
		Description: `A network has **n** nodes labelled 1 to n. Each entry [u, v, w] of **times** is a **directed** link: a signal sent from u reaches v after w time units (w >= 0). Links are one-way: [u, v, w] says nothing about sending from v to u.

At time 0 a signal is sent from node **k**. It travels along every link, so each node hears it at the earliest time any route can deliver it.

Return the time at which the **last** node hears the signal, that is, the largest of the shortest travel times from k. If some node can never hear it, return -1.`,
		Constraints: []string{
			"1 <= k <= n <= 100",
			"1 <= times.length <= 6000",
			"times[i] = [u, v, w] with 1 <= u, v <= n, u != v and 0 <= w <= 100",
			"No pair (u, v) appears twice",
		},
		Examples: []Example{
			{Input: "times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2", Output: "2", Explanation: "From node 2: node 1 and node 3 at time 1, node 4 at time 2 (via 3). The last node hears at time 2."},
			{Input: "times = [[1,2,1]], n = 2, k = 1", Output: "1"},
			{Input: "times = [[1,2,1]], n = 2, k = 2", Output: "-1", Explanation: "The only link goes from 1 to 2, so a signal from 2 never reaches 1."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"times": [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, "n": 4, "k": 2}, Expected: 2, Hidden: false, Note: "Example 1"},
			{Input: map[string]interface{}{"times": [][]int{{1, 2, 1}}, "n": 2, "k": 1}, Expected: 1, Hidden: false, Note: "Example 2"},
			{Input: map[string]interface{}{"times": [][]int{{1, 2, 1}}, "n": 2, "k": 2}, Expected: -1, Hidden: false, Note: "Example 3: node 1 never hears"},
			{Input: map[string]interface{}{"times": [][]int{{1, 2, 10}, {1, 3, 1}, {3, 2, 1}}, "n": 3, "k": 1}, Expected: 2, Hidden: true, Note: "two short hops beat one long edge (fewest edges is not fastest)"},
			{Input: map[string]interface{}{"times": [][]int{{1, 2, 5}, {1, 3, 1}}, "n": 3, "k": 1}, Expected: 5, Hidden: true, Note: "max(dist) on a dict gives the largest label 3, not the time 5"},
			{Input: map[string]interface{}{"times": [][]int{{1, 2, 0}, {2, 3, 0}}, "n": 3, "k": 1}, Expected: 0, Hidden: true, Note: "zero-delay edges"},
			{Input: map[string]interface{}{"times": [][]int{{2, 1, 1}}, "n": 2, "k": 1}, Expected: -1, Hidden: true, Note: "edges are one-way"},
			{Input: map[string]interface{}{"times": [][]int{{1, 2, 4}, {1, 3, 1}, {3, 2, 1}, {2, 4, 1}, {3, 4, 5}}, "n": 4, "k": 1}, Expected: 3, Hidden: true, Note: "node 2 is pushed twice; the stale entry must be skipped"},
		},
		TimeComplexity:  "O(E log E)",
		SpaceComplexity: "O(V + E)",
		StarterCode:     "def networkDelayTime(times, n, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "BFS finds the route with the fewest links. Here a route with three links of time 1 beats one link of time 10. Of all the nodes whose arrival time you have estimated but not settled, which one can you be certain about right now, and why?"},
			{Level: 2, Type: "insight", Content: "Dijkstra's algorithm: keep a min-heap of (time, node). The node with the smallest time in the heap is final, because every other route to it would have to pass through a node with a time at least as large and then add a non-negative delay. Pop it, record its time, and push its neighbours with time + delay. The answer is the largest recorded time, if all n nodes were reached.", ChapterRef: 14},
			{Level: 3, Type: "plan", Content: "1. Build the graph: adjacency lists of (neighbour, delay) for nodes 1..n.\n2. Seed the heap with (0, k).\n3. Pop the cheapest entry. Skip stale entries: if the node is already finalised, continue.\n4. Finalise it: record its time.\n5. Relax the edges: push (time + delay, neighbour) for each neighbour not yet finalised.\n6. Read off the answer: if fewer than n nodes were finalised, return -1; otherwise return the largest recorded time (the maximum of the values, not of the keys)."},
			{Level: 4, Type: "skeleton", Content: `import heapq

def networkDelayTime(times, n, k):
    graph = [[] for _ in range(n + 1)]
    for source, target, delay in times:
        graph[source].append((target, delay))

    dist = {}
    heap = [(0, k)]
    while heap:
        time, node = heapq.heappop(heap)
        if ____:
            continue
        dist[node] = ____
        for nxt, delay in graph[node]:
            if nxt not in dist:
                heapq.heappush(heap, ____)
    if len(dist) < n:
        return -1
    return ____`},
		},
		Solution: Solution{
			Code: `import heapq

def networkDelayTime(times, n, k):
    graph = [[] for _ in range(n + 1)]  # nodes are labelled 1..n
    for source, target, delay in times:
        graph[source].append((target, delay))

    dist = {}  # node -> final shortest arrival time
    heap = [(0, k)]  # (tentative time, node)
    while heap:
        time, node = heapq.heappop(heap)
        if node in dist:
            continue  # stale entry: node was already finalised
        dist[node] = time  # smallest time in the heap, so it is final
        for nxt, delay in graph[node]:
            if nxt not in dist:
                heapq.heappush(heap, (time + delay, nxt))  # relax the edge
    if len(dist) < n:
        return -1  # some node never received the signal
    return max(dist.values())`,
			Explanation: "This is the chapter's Dijkstra template applied directly, and it is the first rung of this ladder. It builds on BFS from Chapter 13 (BFS finalises nodes in order of edge count; Dijkstra finalises them in order of total time) and on the min-heap from Chapter 11. Every node's shortest arrival time is settled the moment it leaves the heap; the answer is the largest of those times.",
			BruteForce:  "Enumerate every simple path from k with DFS and keep the minimum time to each node. The number of simple paths can grow exponentially with n. A better but still slow approach is Bellman-Ford (the alternative below): relax every edge, n - 1 times, which costs O(V * E), around 600,000 relaxations here, repeating edges whose endpoints did not change.",
			KeyInsight:  "With non-negative weights, the smallest tentative time in the heap can never be improved: any other route to that node leaves the settled region through some node whose time is already at least as large, and every further edge adds zero or more. So nodes can be finalised one at a time, smallest first, and each edge only needs to be relaxed once.",
			WhyItWorks:  "Invariant: every node in dist holds its true shortest time, and each time popped from the heap is at least as large as all times popped before it. Suppose node x is popped with time t but some route reaches x faster. That route starts at k (already finalised) and ends at x (not finalised), so it has a first node y that was not finalised when x was popped. The node before y on the route was finalised and pushed y with the route's time to y, which is at most the route's total (all later delays are >= 0), which is < t. Then y's entry would have been popped before x's, a contradiction. So t is final. A node may be pushed several times; the later, larger copies are stale and skipped. Each edge pushes at most once, so there are at most E heap entries: O(E log E). Finally, every node is reachable exactly when it appears in dist, so the answer is max(dist.values()) or -1.",
			CommonMistakes: []string{
				"Writing max(dist) on a dictionary. That is the largest key, a node label (3 in the hidden test), not the largest time (5). Use max(dist.values()).",
				"Using BFS or the fewest-links route. With weights, two links of 1 beat one link of 10; BFS returns 10.",
				"Recording a node's distance when it is pushed instead of when it is popped. A pushed time is only tentative; a cheaper route may still arrive. Finalise on pop and skip stale pops.",
				"Adding edges in both directions. The links are directed; [2, 1, 1] does not let a signal from 1 reach 2.",
			},
			TimeComplexity:  "O(E log E): each edge pushes at most one heap entry, and each push or pop costs O(log E). Since E <= V^2, this is also O(E log V).",
			SpaceComplexity: "O(V + E) for the adjacency lists and the heap.",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the graph", Explanation: "Directed adjacency lists of (neighbour, delay). Index 0 is unused because labels start at 1.", CodeSnippet: "graph = [[] for _ in range(n + 1)]\nfor source, target, delay in times:\n    graph[source].append((target, delay))", LineStart: 4, LineEnd: 6},
				{Title: "Seed the heap", Explanation: "The source hears the signal at time 0. dist starts empty: a node enters it only when its time is final.", CodeSnippet: "dist = {}\nheap = [(0, k)]", LineStart: 8, LineEnd: 9},
				{Title: "Pop the cheapest", Explanation: "The heap hands back the smallest tentative time. If that node is already in dist, this is a stale entry from a slower route, so skip it. Otherwise its time is final.", CodeSnippet: "time, node = heapq.heappop(heap)\nif node in dist:\n    continue\ndist[node] = time", LineStart: 11, LineEnd: 14},
				{Title: "Relax the edges", Explanation: "Offer each neighbour the route through this node. The heap keeps whichever offer is smallest at the top.", CodeSnippet: "for nxt, delay in graph[node]:\n    if nxt not in dist:\n        heapq.heappush(heap, (time + delay, nxt))", LineStart: 15, LineEnd: 17},
				{Title: "Read off the answer", Explanation: "A node missing from dist was never reached. Otherwise the last node to hear is the one with the largest final time.", CodeSnippet: "if len(dist) < n:\n    return -1\nreturn max(dist.values())", LineStart: 18, LineEnd: 20},
			},
			Alternatives: []AltSolution{
				{
					Name: "Bellman-Ford",
					Code: `def networkDelayTime(times, n, k):
    INF = float("inf")
    dist = [INF] * (n + 1)
    dist[k] = 0
    for _ in range(n - 1):  # a shortest path uses at most n - 1 edges
        changed = False
        for source, target, delay in times:
            if dist[source] + delay < dist[target]:
                dist[target] = dist[source] + delay
                changed = True
        if not changed:
            break
    slowest = max(dist[1:])
    return -1 if slowest == INF else slowest`,
					TimeComplexity:  "O(V * E)",
					SpaceComplexity: "O(V)",
					Explanation:     "Relax every edge in rounds. After round i, every node whose shortest route uses at most i edges has its correct time, and a shortest route never needs more than n - 1 edges, so n - 1 rounds suffice (stop early if a round changes nothing). Slower than Dijkstra, but it also works with negative weights, and Cheapest Flights later in this ladder uses exactly this round structure.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is a node's time final when it is popped, and which assumption does that rely on?", Answer: "It is the smallest tentative time in the heap. Any other route to it must leave the finalised set through some unfinalised node whose tentative time is already at least this large, and the remaining edges can only add time because weights are non-negative. With a negative edge that argument breaks: a later detour could subtract time and beat the popped value."},
			{Prompt: "A node can be pushed onto the heap several times. Why is that harmless?", Answer: "Only the first pop of a node carries its smallest time; that pop finalises it. Every later copy has a time at least as large and is skipped by the 'if node in dist: continue' check. The extra entries cost some heap space, at most one per edge, which is why the bound is O(E log E)."},
		},
	},
	{
		ID:              "swim-in-rising-water",
		Number:          778,
		Order:           2,
		Title:           "Swim in Rising Water",
		Slug:            "swim-in-rising-water",
		Difficulty:      "Hard",
		Category:        "advanced-graphs",
		Tags:            []string{"Graph", "Dijkstra", "Heap", "Binary Search", "Matrix"},
		RelatedChapters: []int{14, 11, 13},
		Description: `You are given an n x n grid where grid[r][c] is the elevation of cell (r, c). The elevations are all different and are exactly the numbers 0 to n^2 - 1 in some order.

Rain falls, and at time t the water level everywhere is t. You may swim from a cell to a neighbouring cell (up, down, left or right) at time t only if **both** cells have elevation at most t. Swimming any distance takes no time, but you can only wait for the water to rise.

You start at the top-left cell (0, 0). Return the smallest time t at which you can reach the bottom-right cell (n - 1, n - 1). Note that both the start and the end cell must be under water, so the answer is at least grid[0][0] and at least grid[n-1][n-1].`,
		Constraints: []string{
			"n == grid.length == grid[i].length",
			"1 <= n <= 50",
			"0 <= grid[i][j] < n^2, and all values are distinct",
		},
		Examples: []Example{
			{Input: "grid = [[0,2],[1,3]]", Output: "3", Explanation: "At time 1 you can move from 0 to 1, but the target cell has elevation 3, so you must wait until t = 3 before you can enter it."},
			{Input: "grid = [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]", Output: "16", Explanation: "The best route follows the top row and right column down to 16; every route to the corner must cross a cell of height at least 16."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"grid": [][]int{{0, 2}, {1, 3}}}, Expected: 3, Hidden: false, Note: "Example 1: the corner itself is 3"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}}, Expected: 16, Hidden: false, Note: "Example 2: the winding path tops out at 16"},
			{Input: map[string]interface{}{"grid": [][]int{{0}}}, Expected: 0, Hidden: true, Note: "a 1 x 1 grid: already there"},
			{Input: map[string]interface{}{"grid": [][]int{{3, 2}, {0, 1}}}, Expected: 3, Hidden: true, Note: "the start cell is the highest"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 8, 7}, {1, 6, 5}, {2, 3, 4}}}, Expected: 4, Hidden: true, Note: "detour around high cells"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 1, 2}, {5, 4, 3}, {6, 7, 8}}}, Expected: 8, Hidden: true, Note: "the end cell's own height decides"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 3, 8}, {1, 7, 4}, {6, 2, 5}}}, Expected: 6, Hidden: true, Note: "the best route climbs to 6 and avoids the 7 and 8"},
		},
		TimeComplexity:  "O(n^2 log n)",
		SpaceComplexity: "O(n^2)",
		StarterCode:     "def swimInWater(grid):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "For a fixed route from start to end, at what time can you swim it? So what does 'the cost of a route' mean here, and how is it different from Network Delay Time, where costs added up?"},
			{Level: 2, Type: "insight", Content: "A route can be swum at time t exactly when its highest cell is at most t, so a route's cost is the maximum elevation on it, and you want the route that minimises that maximum. Dijkstra still works when you combine costs with max instead of +: max never makes a route cheaper, which is the same property non-negative weights gave you. Pop the cell with the lowest 'highest point so far'; when the corner pops, that is the answer.", ChapterRef: 14},
			{Level: 3, Type: "plan", Content: "1. Build the graph: each cell's neighbours are its four adjacent cells.\n2. Seed the heap with (grid[0][0], 0, 0).\n3. Pop the cheapest entry; skip stale entries for cells already finalised.\n4. If the popped cell is the corner, return its level.\n5. Relax the edges: push each unfinalised neighbour with max(level, its elevation)."},
			{Level: 4, Type: "skeleton", Content: `import heapq

def swimInWater(grid):
    n = len(grid)
    heap = [(____, 0, 0)]
    done = set()
    while heap:
        level, r, c = heapq.heappop(heap)
        if (r, c) in done:
            continue
        done.add((r, c))
        if ____:
            return level
        for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
            if 0 <= nr < n and 0 <= nc < n and (nr, nc) not in done:
                heapq.heappush(heap, ____)
    return -1`},
		},
		Solution: Solution{
			Code: `import heapq

def swimInWater(grid):
    n = len(grid)
    heap = [(grid[0][0], 0, 0)]  # (highest cell on the best path so far, row, col)
    done = set()
    while heap:
        level, r, c = heapq.heappop(heap)
        if (r, c) in done:
            continue  # stale entry
        done.add((r, c))  # lowest possible level for this cell is now final
        if (r, c) == (n - 1, n - 1):
            return level
        for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
            if 0 <= nr < n and 0 <= nc < n and (nr, nc) not in done:
                # a path's cost is its highest cell, so combine with max, not +
                heapq.heappush(heap, (max(level, grid[nr][nc]), nr, nc))
    return -1`,
			Explanation: "This is Network Delay Time with one change: a route's cost is the highest elevation on it, not the sum of its edges. Dijkstra's loop stays the same, combining with max instead of +, and the cell grid is the graph. The first time the bottom-right cell leaves the heap, its level is the lowest water level at which any route is open.",
			BruteForce:  "Try every time t = 0, 1, 2, ... and flood-fill from (0, 0) through cells with elevation <= t until the corner is reachable: up to n^2 floods of O(n^2) each, O(n^4). Binary searching on t (the alternative below) cuts that to O(n^2 log n), but it still re-floods the grid from scratch for each guess.",
			KeyInsight:  "Dijkstra's correctness argument needs only one thing: extending a route can never make it cheaper. For sums that means non-negative weights; for max it is always true, since max(level, h) >= level. So the cell with the lowest 'worst elevation so far' in the heap is final, exactly like the closest node in Network Delay Time.",
			WhyItWorks:  "Define cost(route) = highest elevation on it. Invariant: when a cell is popped for the first time, level is the smallest possible cost of any route from (0, 0) to it. Suppose a cheaper route existed. Follow it from the start: it leaves the set of finalised cells at some first unfinalised cell y. The cell before y was finalised and pushed y with max(its level, grid[y]), which is at most the route's cost (the max along a route only grows), which is less than level. So y would have been popped before the current cell, a contradiction. The corner is therefore returned with its optimal cost. Each cell is finalised once and pushes at most four entries, so the heap holds O(n^2) entries and each operation costs O(log n^2) = O(log n): O(n^2 log n) total.",
			CommonMistakes: []string{
				"Adding elevations along the route (Network Delay Time's +). The time to swim a route is its highest cell, not the total climb.",
				"Starting the heap at level 0 instead of grid[0][0]. You cannot even stand in the start cell until the water covers it.",
				"Stopping as soon as the corner is pushed, not when it is popped. A pushed level is tentative; a lower route may still reach it.",
				"Using BFS by number of moves. The shortest route is often not the lowest one; the test that detours around 7 and 8 shows this.",
			},
			TimeComplexity:  "O(n^2 log n): n^2 cells, each popped once and pushing up to four entries.",
			SpaceComplexity: "O(n^2) for the heap and the finalised set.",
			Walkthrough: []WalkthroughStep{
				{Title: "Seed the heap", Explanation: "The start cell alone costs its own elevation: you must wait for the water to cover it.", CodeSnippet: "heap = [(grid[0][0], 0, 0)]\ndone = set()", LineStart: 5, LineEnd: 6},
				{Title: "Pop the cheapest", Explanation: "Take the cell with the lowest worst-elevation-so-far. Skip it if it was already finalised (a stale, higher entry). Otherwise its level is final.", CodeSnippet: "level, r, c = heapq.heappop(heap)\nif (r, c) in done:\n    continue\ndone.add((r, c))", LineStart: 8, LineEnd: 11},
				{Title: "Read off the answer", Explanation: "The corner's level becomes final when it is popped; return it right away.", CodeSnippet: "if (r, c) == (n - 1, n - 1):\n    return level", LineStart: 12, LineEnd: 13},
				{Title: "Relax the edges", Explanation: "Stepping into a neighbour costs nothing extra unless it is higher than anything so far, so the new route's cost is the max of the two.", CodeSnippet: "if 0 <= nr < n and 0 <= nc < n and (nr, nc) not in done:\nheapq.heappush(heap, (max(level, grid[nr][nc]), nr, nc))", LineStart: 15, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{
					Name: "Binary search on the time + flood fill",
					Code: `def swimInWater(grid):
    n = len(grid)

    def can_reach(t):  # can we swim to the corner once the water is at level t?
        if grid[0][0] > t:
            return False
        seen = {(0, 0)}
        stack = [(0, 0)]
        while stack:
            r, c = stack.pop()
            if (r, c) == (n - 1, n - 1):
                return True
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < n and 0 <= nc < n and (nr, nc) not in seen and grid[nr][nc] <= t:
                    seen.add((nr, nc))
                    stack.append((nr, nc))
        return False

    low, high = 0, n * n - 1
    while low < high:
        mid = (low + high) // 2
        if can_reach(mid):
            high = mid
        else:
            low = mid + 1
    return low`,
					TimeComplexity:  "O(n^2 log n)",
					SpaceComplexity: "O(n^2)",
					Explanation:     "If you can reach the corner at time t, you can also reach it at any later time, so 'reachable at t' is monotone and binary search (Chapter 6) finds the smallest t. Each guess is a flood fill (Chapter 13) over cells with elevation <= t. Same complexity, built from two older tools instead of Dijkstra.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is Dijkstra still correct when costs combine with max instead of +?", Answer: "The proof only uses the fact that extending a route never lowers its cost. With + that needs non-negative weights; with max it always holds, because max(level, h) >= level. So the lowest-cost cell in the heap cannot be improved by any route that goes through a not-yet-finalised cell."},
			{Prompt: "What changes if elevations could repeat?", Answer: "Nothing in the Dijkstra solution: it never relied on distinct values. The binary search alternative would just search over the range of elevations instead of 0..n^2 - 1. The distinctness in the statement only guarantees that the answer is a value in 0..n^2 - 1."},
		},
	},
	{
		ID:              "min-cost-to-connect-all-points",
		Number:          1584,
		Order:           3,
		Title:           "Min Cost to Connect All Points",
		Slug:            "min-cost-to-connect-all-points",
		Difficulty:      "Medium",
		Category:        "advanced-graphs",
		Tags:            []string{"Graph", "Minimum Spanning Tree", "Prim", "Kruskal", "Union-Find", "Heap"},
		RelatedChapters: []int{14, 11, 13},
		Description: `You are given **points** on a plane, where points[i] = [x, y]. Joining two points [x1, y1] and [x2, y2] costs their **Manhattan distance**, |x1 - x2| + |y1 - y2|. Any pair of points may be joined, and joins are undirected.

Return the minimum total cost to join the points so that every point is connected to every other, directly or through other points. With a single point the cost is 0.`,
		Constraints: []string{
			"1 <= points.length <= 1000",
			"-10^6 <= x, y <= 10^6",
			"All points are distinct",
		},
		Examples: []Example{
			{Input: "points = [[0,0],[2,2],[3,10],[5,2],[7,0]]", Output: "20", Explanation: "Join [0,0]-[2,2] (4), [2,2]-[5,2] (3), [5,2]-[7,0] (4) and [2,2]-[3,10] (9). Total 20; no cheaper set of joins connects all five."},
			{Input: "points = [[3,12],[-2,5],[-4,1]]", Output: "18", Explanation: "[-4,1]-[-2,5] costs 6 and [-2,5]-[3,12] costs 12."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"points": [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}}, Expected: 20, Hidden: false, Note: "Example 1"},
			{Input: map[string]interface{}{"points": [][]int{{3, 12}, {-2, 5}, {-4, 1}}}, Expected: 18, Hidden: false, Note: "Example 2: negative coordinates"},
			{Input: map[string]interface{}{"points": [][]int{{0, 0}}}, Expected: 0, Hidden: true, Note: "one point needs no edges"},
			{Input: map[string]interface{}{"points": [][]int{{0, 0}, {1, 1}}}, Expected: 2, Hidden: true, Note: "two points: one edge"},
			{Input: map[string]interface{}{"points": [][]int{{-1000000, -1000000}, {1000000, 1000000}}}, Expected: 4000000, Hidden: true, Note: "extreme coordinates"},
			{Input: map[string]interface{}{"points": [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}}, Expected: 3, Hidden: true, Note: "points on a line"},
			{Input: map[string]interface{}{"points": [][]int{{2, -3}, {-17, -8}, {13, 8}, {-17, -15}}}, Expected: 53, Hidden: true, Note: "spread-out points"},
		},
		TimeComplexity:  "O(n^2 log n)",
		SpaceComplexity: "O(n^2)",
		StarterCode:     "def minCostConnectPoints(points):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The cheapest way to connect everything never contains a loop (why pay for an edge whose endpoints are already connected?). Suppose some points are already joined into one group. Which single join is always safe to add next?"},
			{Level: 2, Type: "insight", Content: "You want a minimum spanning tree (MST): n - 1 edges that connect all points with the smallest total weight. Prim's algorithm grows one tree from point 0: repeatedly add the cheapest edge from a point in the tree to a point outside it. That edge is always part of some MST (the cut property). With a min-heap the loop looks just like Dijkstra, except the heap key is the edge's own cost, not a running total.", ChapterRef: 14},
			{Level: 3, Type: "plan", Content: "1. Seed the heap with (0, point 0): attaching the first point is free.\n2. Pop the cheapest entry. Skip stale entries: if that point is already in the tree, continue.\n3. Add the point to the tree and its edge cost to the total.\n4. Push the crossing edges: for every point not in the tree, push (distance to the new point, that point).\n5. Read off the answer when all n points are in the tree."},
			{Level: 4, Type: "skeleton", Content: `import heapq

def minCostConnectPoints(points):
    n = len(points)
    in_tree = set()
    heap = [(0, 0)]
    total = 0
    while len(in_tree) < n:
        cost, point = heapq.heappop(heap)
        if ____:
            continue
        in_tree.add(point)
        total += ____
        x, y = points[point]
        for other in range(n):
            if other not in in_tree:
                dist = ____
                heapq.heappush(heap, ____)
    return total`},
		},
		Solution: Solution{
			Code: `import heapq

def minCostConnectPoints(points):
    n = len(points)
    in_tree = set()
    heap = [(0, 0)]  # (cost of the edge that would attach this point, point)
    total = 0
    while len(in_tree) < n:
        cost, point = heapq.heappop(heap)
        if point in in_tree:
            continue  # stale: already attached by a cheaper edge
        in_tree.add(point)  # cheapest edge leaving the tree is always safe
        total += cost
        x, y = points[point]
        for other in range(n):
            if other not in in_tree:
                dist = abs(x - points[other][0]) + abs(y - points[other][1])
                heapq.heappush(heap, (dist, other))
    return total`,
			Explanation: "A different question from the first two problems: not the shortest route from one source, but the cheapest set of edges that connects everything, a minimum spanning tree. Prim's algorithm grows the tree one point at a time, always adding the cheapest edge that leaves it. The code is the Dijkstra loop from Network Delay Time with one change: the heap key is the cost of the single edge that would attach a point, not the total distance from the start.",
			BruteForce:  "Try every subset of n - 1 edges out of the n(n-1)/2 possible ones and keep the cheapest that connects all points; the number of subsets is astronomically large even for 10 points. A natural greedy that is wrong: join each point to its nearest neighbour. That can leave several disconnected clusters.",
			KeyInsight:  "Cut property: split the points into any two non-empty groups; the cheapest edge crossing between them belongs to some minimum spanning tree. Prim uses the split 'in the tree' versus 'not yet in the tree', so the cheapest edge leaving the tree is always safe to add.",
			WhyItWorks:  "A spanning tree connects all n points with exactly n - 1 edges and no cycle; an MST is one with the smallest total cost. Invariant: the edges chosen so far are part of some MST T. Let e be the cheapest edge from the current tree S to the outside. If e is in T, the invariant holds. If not, adding e to T creates a cycle, which must cross between S and the outside at least once more, through some edge f. Since e is the cheapest crossing edge, cost(e) <= cost(f), so T - f + e is a spanning tree that costs no more: also an MST, and it contains our edges plus e. After n - 1 additions, the chosen edges form an MST. Lazy deletion handles outdated heap entries: when a point is popped a second time it is already in the tree and is skipped. Every point pushes an entry for every other point, so the heap sees O(n^2) entries: O(n^2 log n) time.",
			CommonMistakes: []string{
				"Pushing the running total (cost + dist) as in Dijkstra. Prim's key is the cost of the one edge that attaches the point; the total belongs in a separate variable.",
				"Adding the cost when an entry is pushed rather than when it is popped. Only the popped, non-stale entry is the edge actually chosen.",
				"Forgetting the stale check. A point can be in the heap many times; adding it twice double-counts its cost and loops past n points.",
				"Building and sorting all n^2 / 2 edges for Kruskal without noticing the memory: 500,000 tuples for n = 1000 is fine, but it is the reason the O(n^2) array version of Prim is preferred for dense graphs.",
			},
			TimeComplexity:  "O(n^2 log n): about n^2 heap pushes for this complete graph. The array-based Prim alternative is O(n^2).",
			SpaceComplexity: "O(n^2) for the heap in the worst case.",
			Walkthrough: []WalkthroughStep{
				{Title: "Seed the heap", Explanation: "Point 0 can join the tree for free; it will be the first point popped.", CodeSnippet: "in_tree = set()\nheap = [(0, 0)]\ntotal = 0", LineStart: 5, LineEnd: 7},
				{Title: "Pop the cheapest", Explanation: "The heap gives the cheapest edge that attaches some outside point to the tree. If that point joined already (through a cheaper edge), this entry is stale.", CodeSnippet: "cost, point = heapq.heappop(heap)\nif point in in_tree:\n    continue", LineStart: 9, LineEnd: 11},
				{Title: "Add it to the tree", Explanation: "By the cut property this edge belongs to a minimum spanning tree, so pay for it.", CodeSnippet: "in_tree.add(point)\ntotal += cost", LineStart: 12, LineEnd: 13},
				{Title: "Push the crossing edges", Explanation: "Every outside point now has a new candidate edge to the tree: the one to this point. Push it with its own cost.", CodeSnippet: "for other in range(n):\n    if other not in in_tree:\n        dist = abs(x - points[other][0]) + abs(y - points[other][1])\n        heapq.heappush(heap, (dist, other))", LineStart: 15, LineEnd: 18},
				{Title: "Read off the answer", Explanation: "The loop stops once all n points are in the tree; total is the cost of its n - 1 edges.", CodeSnippet: "return total", LineStart: 19, LineEnd: 19},
			},
			Alternatives: []AltSolution{
				{
					Name: "Kruskal with union-find",
					Code: `def minCostConnectPoints(points):
    n = len(points)
    edges = sorted(
        (abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]), i, j)
        for i in range(n) for j in range(i + 1, n)
    )
    parent = list(range(n))

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    total, used = 0, 0
    for cost, i, j in edges:
        root_i, root_j = find(i), find(j)
        if root_i != root_j:  # joins two different trees: no cycle
            parent[root_i] = root_j
            total += cost
            used += 1
            if used == n - 1:
                break
    return total`,
					TimeComplexity:  "O(n^2 log n)",
					SpaceComplexity: "O(n^2)",
					Explanation:     "Sort all edges by cost and take each one that joins two different components, using the union-find from Chapter 13 to tell (the same test as Redundant Connection, with the opposite verdict). Stop after n - 1 edges. Kruskal is the natural choice when the graph comes as an edge list; here every pair is an edge, so the sort dominates.",
				},
				{
					Name: "Prim with an array (no heap)",
					Code: `def minCostConnectPoints(points):
    n = len(points)
    in_tree = [False] * n
    cheapest = [float("inf")] * n  # cheapest edge from each point to the tree
    cheapest[0] = 0
    total = 0
    for _ in range(n):
        point = -1
        for i in range(n):
            if not in_tree[i] and (point == -1 or cheapest[i] < cheapest[point]):
                point = i
        in_tree[point] = True
        total += cheapest[point]
        x, y = points[point]
        for other in range(n):
            if not in_tree[other]:
                cost = abs(x - points[other][0]) + abs(y - points[other][1])
                cheapest[other] = min(cheapest[other], cost)
    return total`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(n)",
					Explanation:     "Keep cheapest[i], the cheapest edge from point i to the tree. Each round, scan for the outside point with the smallest value, add it, and update every other point's value with its distance to the new point. On a complete graph this is the fastest option, because there is no log factor and no heap of n^2 entries.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Prim and Dijkstra look almost identical in code. What exactly differs, and why does it change the answer?", Answer: "The heap key. Dijkstra pushes time + delay, the length of the whole route from the source, because it minimises distance from one source. Prim pushes just the edge cost, because it minimises the total of the chosen edges, and a point should join through its cheapest single link to the tree regardless of how far that is from point 0."},
			{Prompt: "Why can the minimum way to connect all points never contain a cycle?", Answer: "If the chosen edges contain a cycle, removing any edge on the cycle keeps everything connected (the rest of the cycle still links its two ends) and lowers the cost, or keeps it equal for a zero-cost edge. So a cheapest connected set can always be taken cycle-free, which with n points means exactly n - 1 edges: a spanning tree."},
		},
	},
	{
		ID:              "cheapest-flights-within-k-stops",
		Number:          787,
		Order:           4,
		Title:           "Cheapest Flights Within K Stops",
		Slug:            "cheapest-flights-within-k-stops",
		Difficulty:      "Medium",
		Category:        "advanced-graphs",
		Tags:            []string{"Graph", "Shortest Path", "Bellman-Ford", "Dynamic Programming"},
		RelatedChapters: []int{14, 13},
		Description: `There are **n** cities labelled 0 to n - 1. Each entry [from, to, price] of **flights** is a **one-way** flight costing price.

Find the cheapest way to travel from city **src** to city **dst** using **at most k stops**. A stop is an intermediate city where you change planes, so a trip with at most k stops uses at most k + 1 flights. Return the cheapest total price, or -1 if no such trip exists.`,
		Constraints: []string{
			"1 <= n <= 100",
			"0 <= flights.length <= n * (n - 1) / 2",
			"flights[i] = [from, to, price] with from != to and 1 <= price <= 10^4",
			"No two flights share the same (from, to)",
			"0 <= src, dst, k < n and src != dst",
		},
		Examples: []Example{
			{Input: "n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1", Output: "700", Explanation: "0 -> 1 -> 3 costs 700 with one stop. The cheaper 0 -> 1 -> 2 -> 3 (400) needs two stops, which is too many."},
			{Input: "n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1", Output: "200", Explanation: "0 -> 1 -> 2 with one stop."},
			{Input: "n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0", Output: "500", Explanation: "No stops allowed, so only the direct flight."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 4, "flights": [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, "src": 0, "dst": 3, "k": 1}, Expected: 700, Hidden: false, Note: "Example 1: the cheaper route needs 2 stops"},
			{Input: map[string]interface{}{"n": 3, "flights": [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, "src": 0, "dst": 2, "k": 1}, Expected: 200, Hidden: false, Note: "Example 2: one stop allowed"},
			{Input: map[string]interface{}{"n": 3, "flights": [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, "src": 0, "dst": 2, "k": 0}, Expected: 500, Hidden: false, Note: "Example 3: no stops, so the direct flight"},
			{Input: map[string]interface{}{"n": 4, "flights": [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 10}}, "src": 0, "dst": 3, "k": 1}, Expected: 10, Hidden: true, Note: "updating in place within a round would chain three flights and return 3"},
			{Input: map[string]interface{}{"n": 3, "flights": [][]int{{0, 1, 100}}, "src": 0, "dst": 2, "k": 1}, Expected: -1, Hidden: true, Note: "destination unreachable"},
			{Input: map[string]interface{}{"n": 4, "flights": [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, "src": 0, "dst": 3, "k": 1}, Expected: -1, Hidden: true, Note: "reachable only with too many stops"},
			{Input: map[string]interface{}{"n": 5, "flights": [][]int{{0, 1, 5}, {1, 2, 5}, {0, 3, 2}, {3, 1, 2}, {1, 4, 1}, {2, 4, 1}}, "src": 0, "dst": 2, "k": 2}, Expected: 9, Hidden: true, Note: "two stops allowed: the cheaper detour through 3"},
			{Input: map[string]interface{}{"n": 5, "flights": [][]int{{0, 1, 5}, {1, 2, 5}, {0, 3, 2}, {3, 1, 2}, {1, 4, 1}, {2, 4, 1}}, "src": 0, "dst": 2, "k": 1}, Expected: 10, Hidden: true, Note: "same flights, one stop: plain Dijkstra would pick the detour"},
		},
		TimeComplexity:  "O(k * E)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def findCheapestPrice(n, flights, src, dst, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Plain Dijkstra settles each city once with its cheapest price. In Example 1, the cheapest way to city 2 uses one stop already; could the best trip to dst need a pricier route to city 2 that uses fewer flights? What would you have to remember about each city besides its price?"},
			{Level: 2, Type: "insight", Content: "Count flights instead of settling cities. Let cost[c] be the cheapest price to reach c using at most r flights. One round of relaxing every flight turns 'at most r flights' into 'at most r + 1 flights'. That is Bellman-Ford stopped after k + 1 rounds. The catch: within a round, read prices from the previous round only (work on a copy), or one round could chain several flights.", ChapterRef: 14},
			{Level: 3, Type: "plan", Content: "1. Seed the costs: cost[src] = 0, every other city infinity (0 flights used).\n2. Repeat k + 1 times: copy cost to next_cost; for every flight [a, b, p], relax next_cost[b] with cost[a] + p (read from cost, write to next_cost); then cost = next_cost.\n3. Read off the answer: cost[dst], or -1 if it is still infinity."},
			{Level: 4, Type: "skeleton", Content: `def findCheapestPrice(n, flights, src, dst, k):
    INF = float("inf")
    cost = [INF] * n
    cost[src] = 0
    for _ in range(____):
        next_cost = ____
        for start, end, price in flights:
            if ____ < next_cost[end]:
                next_cost[end] = ____
        cost = next_cost
    return cost[dst] if cost[dst] < INF else -1`},
		},
		Solution: Solution{
			Code: `def findCheapestPrice(n, flights, src, dst, k):
    INF = float("inf")
    cost = [INF] * n  # cheapest price to each city using at most r flights
    cost[src] = 0
    for _ in range(k + 1):  # k stops means at most k + 1 flights
        next_cost = cost[:]  # read last round's prices, write into a copy
        for start, end, price in flights:
            if cost[start] + price < next_cost[end]:
                next_cost[end] = cost[start] + price
        cost = next_cost
    return cost[dst] if cost[dst] < INF else -1`,
			Explanation: "Network Delay Time settled each node once with its cheapest time. The stop limit breaks that: the cheapest way into a city may use too many flights to continue. Bellman-Ford, introduced as the alternative in Network Delay Time, handles it naturally, because round r computes the cheapest prices using at most r flights. Running exactly k + 1 rounds gives the answer, provided each round reads only the previous round's prices.",
			BruteForce:  "Enumerate every route from src with DFS, cutting off after k + 1 flights, and keep the cheapest that ends at dst. The number of routes can grow like n^(k+1). A plain Dijkstra that ignores stops is fast but wrong: it settles a city with its cheapest price even when that price needed too many flights (the last hidden test).",
			KeyInsight:  "Index the shortest-path problem by the number of flights. 'Cheapest with at most r + 1 flights' follows from 'cheapest with at most r flights' by trying one more flight on every edge, which is exactly one Bellman-Ford round.",
			WhyItWorks:  "Invariant: after round r (r = 0 before the loop), cost[c] is the cheapest price from src to c using at most r flights. Round r + 1 computes next_cost[b] = min(cost[b], min over flights a -> b of cost[a] + price). A best route with at most r + 1 flights either uses at most r flights (already in cost[b]) or ends with some flight a -> b after a best route to a with at most r flights; both are considered, so the invariant holds. Reading from cost and writing to next_cost is essential: if a round read values written in the same round, a city updated a moment ago could be extended again, letting one round add several flights. After k + 1 rounds, cost[dst] is the cheapest price with at most k + 1 flights, which is at most k stops. Each round relaxes every flight once: O(k * E).",
			CommonMistakes: []string{
				"Updating cost in place during a round. One round can then chain several flights (0 -> 1 -> 2 -> 3 in a single round), and the hidden test returns 3 instead of 10. Read from the previous round, write to a copy.",
				"Running k rounds instead of k + 1. k stops means k + 1 flights.",
				"Using plain Dijkstra that settles each city once by price. A cheap route into a city may have used up the stops, and a pricier route with fewer flights is what reaches dst.",
				"Returning infinity instead of -1 when dst is unreachable within the limit.",
			},
			TimeComplexity:  "O(k * E): k + 1 rounds, each relaxing all E flights.",
			SpaceComplexity: "O(n) for two price arrays.",
			Walkthrough: []WalkthroughStep{
				{Title: "Seed the costs", Explanation: "With zero flights you can only be at src, for free.", CodeSnippet: "cost = [INF] * n\ncost[src] = 0", LineStart: 3, LineEnd: 4},
				{Title: "Freeze last round", Explanation: "k stops allow k + 1 flights, so run k + 1 rounds. Copy the array so this round reads only prices that used at most r flights.", CodeSnippet: "for _ in range(k + 1):\n    next_cost = cost[:]", LineStart: 5, LineEnd: 6},
				{Title: "Relax the edges", Explanation: "Try adding one more flight to every best route of the previous round; keep the cheaper price in the new array.", CodeSnippet: "for start, end, price in flights:\n    if cost[start] + price < next_cost[end]:\n        next_cost[end] = cost[start] + price\ncost = next_cost", LineStart: 7, LineEnd: 10},
				{Title: "Read off the answer", Explanation: "After k + 1 rounds, cost[dst] is the cheapest trip with at most k stops, if one exists.", CodeSnippet: "return cost[dst] if cost[dst] < INF else -1", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{
					Name: "Dijkstra over (city, flights used) states",
					Code: `import heapq

def findCheapestPrice(n, flights, src, dst, k):
    graph = [[] for _ in range(n)]
    for start, end, price in flights:
        graph[start].append((end, price))
    fewest = [float("inf")] * n  # fewest flights used when each city was popped
    heap = [(0, src, 0)]  # (cost so far, city, flights used)
    while heap:
        cost, city, used = heapq.heappop(heap)
        if city == dst:
            return cost
        if used > k or used >= fewest[city]:
            continue  # out of flights, or an earlier, cheaper visit used no more flights
        fewest[city] = used
        for nxt, price in graph[city]:
            heapq.heappush(heap, (cost + price, nxt, used + 1))
    return -1`,
					TimeComplexity:  "O(E * k * log(E * k))",
					SpaceComplexity: "O(E * k)",
					Explanation:     "Make the number of flights part of the state, so a city reached cheaply but with many flights does not block a pricier visit with fewer. Pop states by cost; the first time dst is popped, its cost is the answer. A state is skipped if it has no flights left, or if the same city was already popped (more cheaply) with no more flights used, since that earlier visit dominates it.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does each round need a copy of the cost array?", Answer: "The invariant is that round r + 1 uses routes with at most one more flight than round r. If you write into the array you are reading, a city improved earlier in the same round can be extended again later in that round, so one round may add several flights and the stop limit is silently exceeded."},
			{Prompt: "Plain Bellman-Ford runs n - 1 rounds. Why does it stop at k + 1 here, and what would n - 1 rounds compute?", Answer: "Round r gives the cheapest price using at most r flights. The question limits trips to k + 1 flights, so we stop there. Running n - 1 rounds would give the cheapest price with no stop limit, because a cheapest route without negative cycles never needs more than n - 1 flights."},
		},
	},
	{
		ID:              "reconstruct-itinerary",
		Number:          332,
		Order:           5,
		Title:           "Reconstruct Itinerary",
		Slug:            "reconstruct-itinerary",
		Difficulty:      "Hard",
		Category:        "advanced-graphs",
		Tags:            []string{"Graph", "Eulerian Path", "Hierholzer", "DFS"},
		RelatedChapters: []int{14, 13, 12},
		Description: `You are given a list of airline **tickets**, where tickets[i] = [from, to] is a one-way (**directed**) flight. All the tickets belong to one traveller, whose trip starts at "JFK".

Rebuild the itinerary: a list of airports, starting with "JFK", in which every ticket is used **exactly once** (a ticket that appears twice must be used twice). It is guaranteed that at least one such itinerary exists.

If several itineraries are possible, return the one that is smallest in lexical order when the airports are read as one sequence of strings: for example ["JFK","LGA"] comes before ["JFK","LGB"].`,
		Constraints: []string{
			"1 <= tickets.length <= 300",
			"tickets[i].length == 2; each airport code is three uppercase English letters",
			"from != to for every ticket",
			"The tickets form at least one valid itinerary starting at JFK",
		},
		Examples: []Example{
			{Input: `tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]`, Output: `["JFK","MUC","LHR","SFO","SJC"]`, Explanation: "Each airport has only one ticket out, so there is just one way to use them all."},
			{Input: `tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]`, Output: `["JFK","ATL","JFK","SFO","ATL","SFO"]`, Explanation: "[\"JFK\",\"SFO\",\"ATL\",\"JFK\",\"ATL\",\"SFO\"] also uses every ticket, but it is larger in lexical order (ATL < SFO at the second position)."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"tickets": [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}}, Expected: []string{"JFK", "MUC", "LHR", "SFO", "SJC"}, Hidden: false, Note: "Example 1: only one route"},
			{Input: map[string]interface{}{"tickets": [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}}, Expected: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}, Hidden: false, Note: "Example 2: choose the smaller airport when there is a choice"},
			{Input: map[string]interface{}{"tickets": [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}}, Expected: []string{"JFK", "NRT", "JFK", "KUL"}, Hidden: true, Note: "greedy smallest-first gets stuck at KUL; it must be the last stop"},
			{Input: map[string]interface{}{"tickets": [][]string{{"JFK", "SFO"}}}, Expected: []string{"JFK", "SFO"}, Hidden: true, Note: "a single ticket"},
			{Input: map[string]interface{}{"tickets": [][]string{{"JFK", "AAA"}, {"AAA", "JFK"}, {"JFK", "AAA"}}}, Expected: []string{"JFK", "AAA", "JFK", "AAA"}, Hidden: true, Note: "the same ticket twice: use each copy once"},
			{Input: map[string]interface{}{"tickets": [][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}}}, Expected: []string{"JFK", "ANU", "EZE", "AXA", "TIA", "ANU", "JFK", "TIA", "ANU", "TIA", "JFK"}, Hidden: true, Note: "ten tickets with several loops"},
		},
		TimeComplexity:  "O(E log E)",
		SpaceComplexity: "O(E)",
		StarterCode:     "def findItinerary(tickets):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Always flying to the alphabetically smallest destination is tempting. In the tickets JFK->KUL, JFK->NRT, NRT->JFK, where does that get you stuck? Since a valid itinerary is guaranteed, where must that dead-end airport appear in the final answer?"},
			{Level: 2, Type: "insight", Content: "Using every edge exactly once is an Eulerian path. Hierholzer's algorithm: walk greedily (smallest destination first), removing each ticket as you use it. When you reach an airport with no tickets left, it must be the end of whatever remains of the trip, so add it to the route at that moment: that is post-order. Reverse the route at the end. Dead ends you hit early simply end up last.", ChapterRef: 14},
			{Level: 3, Type: "plan", Content: "1. Build the graph: for each airport, list its destinations sorted so the smallest can be taken first (sorting in reverse and popping from the end is O(1) per ticket).\n2. Visit JFK: while the current airport still has tickets, take the smallest one (removing it) and visit that airport.\n3. When an airport has no tickets left, append it to the route (post-order).\n4. Read off the answer: reverse the route."},
			{Level: 4, Type: "skeleton", Content: `from collections import defaultdict

def findItinerary(tickets):
    destinations = defaultdict(list)
    for start, end in sorted(tickets, reverse=True):
        destinations[start].append(end)
    route = []

    def visit(airport):
        while ____:
            visit(____)
        ____

    visit("JFK")
    return ____`},
		},
		Solution: Solution{
			Code: `from collections import defaultdict

def findItinerary(tickets):
    destinations = defaultdict(list)
    for start, end in sorted(tickets, reverse=True):
        destinations[start].append(end)  # reverse order, so pop() gives the smallest
    route = []

    def visit(airport):
        while destinations[airport]:
            visit(destinations[airport].pop())  # use each ticket exactly once
        route.append(airport)  # post-order: added after all its onward tickets are used

    visit("JFK")
    return route[::-1]`,
			Explanation: "The earlier problems visited every node; this one must use every edge exactly once, which is an Eulerian path. Hierholzer's algorithm is a DFS that consumes edges instead of marking nodes, and records each airport in post-order, when it has no tickets left. Taking the smallest destination first and reversing the post-order gives the lexically smallest itinerary. It reuses DFS from Chapter 13; the backtracking alternative is the Chapter 12 approach.",
			BruteForce:  "Backtracking (the alternative below): from the current airport, try unused tickets in sorted order, recurse, and undo when you get stuck before using all tickets. The first complete route found is the smallest. It is correct, but a dead end discovered late forces undoing a long prefix, and in the worst case the search is exponential.",
			KeyInsight:  "When the walk reaches an airport with no unused tickets, that airport has to be the last stop of the remaining trip, because there is no way to leave it again. Recording airports at that moment (post-order) builds the itinerary from the end backwards, so a dead end hit early lands at the end of the route without any undoing.",
			WhyItWorks:  "Since a valid itinerary exists, every airport except the start and the end has as many tickets in as out. A walk that keeps taking unused tickets can therefore only get stuck at the end airport the first time (any other airport you enter still has an unused ticket out). Post-order adds an airport only when all of its outgoing tickets have been used, and only after every airport visited from it later has been added. So in the reversed list, each airport is followed by the trip that departs from it, every ticket appears once, and the walk is a valid Eulerian path. Side trips (loops that come back to an airport) are explored from inside the while loop and are spliced into the route before the airport they return to, which is exactly where they belong. Because destinations are taken smallest first, the first-explored (smallest) branch ends up earliest in the reversed route, unless it is a dead end, which must go last anyway. Sorting costs O(E log E); the walk uses each ticket once.",
			CommonMistakes: []string{
				"Appending each airport when you arrive (pre-order). On JFK->KUL, JFK->NRT, NRT->JFK that gives JFK, KUL, NRT, JFK, which is not even a valid trip. Append after the airport's tickets are used up, then reverse.",
				"Marking airports as visited. Airports may be visited many times; it is the tickets (edges) that are used once. Remove a ticket when you use it.",
				"Popping from the front of a sorted list (pop(0)), which is O(E) per ticket. Sort in reverse and pop from the end, or use a heap per airport.",
				"Dropping duplicate tickets by storing destinations in a set. Two identical tickets must both be used.",
			},
			TimeComplexity:  "O(E log E) for sorting; the walk itself is O(E).",
			SpaceComplexity: "O(E) for the adjacency lists, the route and the recursion stack.",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the graph", Explanation: "Sorting tickets in reverse means each airport's list ends with its smallest destination, so pop() takes the smallest in O(1) and removes the ticket at the same time.", CodeSnippet: "for start, end in sorted(tickets, reverse=True):\n    destinations[start].append(end)", LineStart: 5, LineEnd: 6},
				{Title: "Use every ticket", Explanation: "While this airport still has an unused ticket, take the smallest one and continue the walk from there. Popping the ticket is what guarantees it is used once.", CodeSnippet: "while destinations[airport]:\n    visit(destinations[airport].pop())", LineStart: 10, LineEnd: 11},
				{Title: "Record in post-order", Explanation: "Once no tickets leave this airport, everything after it in the trip is already in route, so add it now. The route is built back to front.", CodeSnippet: "route.append(airport)", LineStart: 12, LineEnd: 12},
				{Title: "Read off the answer", Explanation: "Start at JFK and reverse the post-order to get the itinerary from start to end.", CodeSnippet: "visit(\"JFK\")\nreturn route[::-1]", LineStart: 14, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{
					Name: "Backtracking in sorted order",
					Code: `def findItinerary(tickets):
    tickets = sorted(tickets)
    used = [False] * len(tickets)
    route = ["JFK"]

    def extend():
        if len(route) == len(tickets) + 1:
            return True  # every ticket used
        for i, (start, end) in enumerate(tickets):
            if not used[i] and start == route[-1]:
                used[i] = True
                route.append(end)
                if extend():
                    return True
                route.pop()  # dead end: undo and try the next ticket
                used[i] = False
        return False

    extend()
    return route`,
					TimeComplexity:  "Exponential in the worst case",
					SpaceComplexity: "O(E)",
					Explanation:     "The Chapter 12 template: choose an unused ticket from the current airport (in sorted order), recurse, and un-choose on failure. The first route that uses every ticket is the lexically smallest. Easy to trust, and fine for small inputs, but it can redo a lot of work after late dead ends.",
				},
				{
					Name: "Hierholzer with an explicit stack",
					Code: `from collections import defaultdict

def findItinerary(tickets):
    destinations = defaultdict(list)
    for start, end in sorted(tickets, reverse=True):
        destinations[start].append(end)
    route = []
    stack = ["JFK"]
    while stack:
        airport = stack[-1]
        if destinations[airport]:
            stack.append(destinations[airport].pop())
        else:
            route.append(stack.pop())
    return route[::-1]`,
					TimeComplexity:  "O(E log E)",
					SpaceComplexity: "O(E)",
					Explanation:     "The same algorithm without recursion: the stack holds the current walk. If the top airport has tickets left, push the smallest destination; otherwise pop it into the route. Useful when the itinerary is long enough to worry about recursion depth.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does recording an airport only after its tickets are used up (post-order) fix the dead-end problem?", Answer: "An airport with no unused tickets left can only be the end of the remaining trip. Post-order writes airports exactly at that moment, so the route is built from the end backwards. A dead end found early (KUL) is written first and, after reversing, ends up last, which is where it must be. The loop then continues from the previous airport with the remaining tickets."},
			{Prompt: "Why do we remove tickets rather than mark airports as visited?", Answer: "The constraint is on edges: each ticket once. Airports are often visited several times (JFK twice in Example 2). Removing a ticket when it is used enforces exactly the right constraint, and it also guarantees termination, since every step consumes one of the finitely many tickets."},
		},
	},
	{
		ID:              "alien-dictionary",
		Number:          269,
		Order:           6,
		Title:           "Alien Dictionary",
		Slug:            "alien-dictionary",
		Difficulty:      "Hard",
		Category:        "advanced-graphs",
		Tags:            []string{"Graph", "Topological Sort", "String", "BFS"},
		RelatedChapters: []int{14, 13},
		Description: `An alien language uses lowercase English letters, but in an unknown order. You are given **words**, a list that is claimed to be sorted in lexicographic order of that alien alphabet.

Lexicographic order works as in a dictionary: compare two words letter by letter; at the first position where they differ, the word with the earlier letter comes first. If one word is a **prefix** of the other (they never differ), the shorter word comes first.

Return a string containing **every distinct letter that appears in words**, each exactly once, in an order consistent with the list. If several orders are consistent (some letters are never compared), return any of them. If no order can make the list sorted, return "". In particular, a word that comes before its own prefix, such as ["abc", "ab"], makes the list invalid, and the answer is "".`,
		Constraints: []string{
			"1 <= words.length <= 100",
			"1 <= words[i].length <= 100",
			"words[i] consists of lowercase English letters",
		},
		Examples: []Example{
			{Input: `words = ["wrt","wrf","er","ett","rftt"]`, Output: `"wertf"`, Explanation: "wrt < wrf gives t before f; wrf < er gives w before e; er < ett gives r before t; ett < rftt gives e before r. Chaining them: w, e, r, t, f."},
			{Input: `words = ["z","x"]`, Output: `"zx"`},
			{Input: `words = ["z","x","z"]`, Output: `""`, Explanation: "The list says z before x and also x before z, which is impossible."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"words": []string{"wrt", "wrf", "er", "ett", "rftt"}}, Expected: "wertf", Hidden: false, Note: "Example 1: the order is fully determined"},
			{Input: map[string]interface{}{"words": []string{"z", "x"}}, Expected: "zx", Hidden: false, Note: "Example 2: z before x"},
			{Input: map[string]interface{}{"words": []string{"z", "x", "z"}}, Expected: "", Hidden: false, Note: "Example 3: z < x and x < z is a contradiction"},
			{Input: map[string]interface{}{"words": []string{"abc", "ab"}}, Expected: "", Hidden: true, Note: "a longer word before its own prefix is invalid"},
			{Input: map[string]interface{}{"words": []string{"z", "z"}}, Expected: "z", Hidden: true, Note: "a single letter"},
			{Input: map[string]interface{}{"words": []string{"ab", "adc"}}, Expected: "abcd", Hidden: true, Note: "only b < d is known; a and c can go anywhere"},
			{Input: map[string]interface{}{"words": []string{"ab", "abc"}}, Expected: "abc", Hidden: true, Note: "a prefix before the longer word is fine and gives no rule"},
			{Input: map[string]interface{}{"words": []string{"ac", "ab", "zc", "zb"}}, Expected: "aczb", Hidden: true, Note: "two independent rules: c < b and a < z"},
		},
		TimeComplexity:  "O(C)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def alienOrder(words):\n    # Write your solution here\n    pass",
		Harness: Harness{
			Validator: `def check(raw, args, result, expected):
    words = raw["words"]
    if expected == "":
        return result == "", "no letter order can explain this list, so the answer is \"\""
    if not isinstance(result, str):
        return False, "return the order as a string"
    letters = sorted({ch for word in words for ch in word})
    if sorted(result) != letters:
        return False, "the order must contain every letter of the words exactly once"
    rank = {ch: i for i, ch in enumerate(result)}
    for first, second in zip(words, words[1:]):
        if [rank[ch] for ch in first] > [rank[ch] for ch in second]:
            return False, "under order %r, %r would sort after %r" % (result, first, second)
    return True`,
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Compare \"wrt\" and \"wrf\", which are next to each other in the list. What, exactly, does their order tell you about the alphabet? Does comparing letters after the first difference tell you anything more?"},
			{Level: 2, Type: "insight", Content: "Each adjacent pair of words gives at most one rule: at the first position where they differ, the first word's letter comes before the second's. Those rules are edges of a directed graph on letters, and the answer is a topological order of it (Kahn's algorithm from Course Schedule II). A cycle means \"\", and so does a longer word placed before its own prefix.", ChapterRef: 13},
			{Level: 3, Type: "plan", Content: "1. Build the graph: create a node for every letter that appears. For each adjacent pair, find the first differing position and add the edge first-letter -> second-letter (once). If there is no difference and the first word is longer, return \"\".\n2. Count in-degrees as you add edges.\n3. Seed the frontier with every letter of in-degree 0.\n4. Expand one node: append it to the order and decrement its successors.\n5. Detect a cycle: if the order is shorter than the number of letters, return \"\"; otherwise join it."},
			{Level: 4, Type: "skeleton", Content: `from collections import deque

def alienOrder(words):
    letters = {ch for word in words for ch in word}
    after = {ch: set() for ch in letters}
    indegree = {ch: 0 for ch in letters}
    for first, second in zip(words, words[1:]):
        for a, b in zip(first, second):
            if a != b:
                ____  # add edge a -> b once, update indegree
                break
        else:
            if ____:
                return ""
    ready = deque(ch for ch in letters if indegree[ch] == 0)
    order = []
    while ready:
        ch = ready.popleft()
        order.append(ch)
        for nxt in after[ch]:
            ____
    return ____`},
		},
		Solution: Solution{
			Code: `from collections import deque

def alienOrder(words):
    letters = {ch for word in words for ch in word}
    after = {ch: set() for ch in letters}  # edge a -> b means a comes before b
    indegree = {ch: 0 for ch in letters}
    for first, second in zip(words, words[1:]):
        for a, b in zip(first, second):
            if a != b:
                if b not in after[a]:
                    after[a].add(b)
                    indegree[b] += 1
                break  # only the first difference tells us anything
        else:
            if len(first) > len(second):
                return ""  # "abc" before "ab": no order can explain it
    ready = deque(sorted(ch for ch in letters if indegree[ch] == 0))
    order = []
    while ready:
        ch = ready.popleft()
        order.append(ch)
        for nxt in sorted(after[ch]):
            indegree[nxt] -= 1
            if indegree[nxt] == 0:
                ready.append(nxt)
    return "".join(order) if len(order) == len(letters) else ""`,
			Explanation: "The final rung puts the topological sort from Course Schedule II to work on a graph you have to discover. Adjacent words reveal letter-order rules at their first difference, the rules form a directed graph, and Kahn's algorithm turns it into an alphabet or reports a contradiction. The judge accepts any order that sorts the list correctly, since letters that are never compared can go anywhere.",
			BruteForce:  "Try every ordering of the distinct letters and return the first one under which the list is sorted (the last alternative below). With up to 26 letters that is 26! orderings, far too many, but on small inputs it is a perfect cross-check because it uses the definition directly.",
			KeyInsight:  "Only the first differing letter of each adjacent pair carries information, and only adjacent pairs are needed: if each neighbouring pair is in order, the whole list is (order is transitive). That reduces the problem to a topological sort of at most 26 letters.",
			WhyItWorks:  "The list is sorted under an alphabet exactly when every adjacent pair is in order. For a pair that differs, being in order means first[i] < second[i] at the first differing position i, which is exactly the edge we add; later positions do not matter. For a pair with no difference, being in order means the first word is not longer, which no alphabet can fix, so a longer first word forces \"\". So the valid alphabets are exactly the orders of the letters that respect every edge, the topological orders. Kahn's algorithm produces one if and only if the graph has no cycle (as shown in Course Schedule); a short output means a cycle, and \"\" is correct. Comparing adjacent pairs touches each character at most once, O(C) for C total characters, and the graph has at most 26 nodes and 26 * 25 edges, a constant.",
			CommonMistakes: []string{
				"Comparing past the first difference. In wrt vs wrf only t < f is known; nothing follows about the letters after it. Break after the first differing position.",
				"Missing the prefix rule. [\"abc\", \"ab\"] has no differing position, yet no alphabet can put abc before ab, so the answer must be \"\".",
				"Leaving out letters that never appear in any rule. Every distinct letter in the words must be in the output; create a node for each one up front.",
				"Incrementing the in-degree twice for a repeated rule. If the same edge a -> b is discovered by two pairs and counted twice, b's in-degree never reaches 0. Store successors in a set and count new edges only.",
			},
			TimeComplexity:  "O(C), where C is the total number of characters in words; the topological sort works on at most 26 letters.",
			SpaceComplexity: "O(1): at most 26 nodes and 26 * 25 edges, independent of the input size.",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the graph", Explanation: "Every letter that appears is a node, even if no rule mentions it.", CodeSnippet: "letters = {ch for word in words for ch in word}\nafter = {ch: set() for ch in letters}\nindegree = {ch: 0 for ch in letters}", LineStart: 4, LineEnd: 6},
				{Title: "Read one rule per pair", Explanation: "The first differing position gives the only rule this pair supports. Use a set so a rule found twice is counted once in the in-degree.", CodeSnippet: "for a, b in zip(first, second):\n    if a != b:\n        if b not in after[a]:\n            after[a].add(b)\n            indegree[b] += 1\n        break", LineStart: 8, LineEnd: 13},
				{Title: "Reject a prefix out of order", Explanation: "The for-else runs only when the loop did not break, meaning one word is a prefix of the other. If the first is longer, no alphabet can sort the pair.", CodeSnippet: "else:\n    if len(first) > len(second):\n        return \"\"", LineStart: 14, LineEnd: 16},
				{Title: "Seed the frontier", Explanation: "Letters that no rule places after anything can come first. Sorting them just makes the output deterministic.", CodeSnippet: "ready = deque(sorted(ch for ch in letters if indegree[ch] == 0))\norder = []", LineStart: 17, LineEnd: 18},
				{Title: "Expand one node", Explanation: "Kahn's algorithm, as in Course Schedule II: place the letter, then release the letters that were waiting only for it.", CodeSnippet: "ch = ready.popleft()\norder.append(ch)\nfor nxt in sorted(after[ch]):\n    indegree[nxt] -= 1\n    if indegree[nxt] == 0:\n        ready.append(nxt)", LineStart: 20, LineEnd: 25},
				{Title: "Detect a cycle", Explanation: "Letters on a cycle never reach in-degree 0, so a short order means the rules contradict each other.", CodeSnippet: "return \"\".join(order) if len(order) == len(letters) else \"\"", LineStart: 26, LineEnd: 26},
			},
			Alternatives: []AltSolution{
				{
					Name: "Three-colour DFS",
					Code: `def alienOrder(words):
    before = {ch: set() for word in words for ch in word}  # ch -> letters that must come before it
    for first, second in zip(words, words[1:]):
        for a, b in zip(first, second):
            if a != b:
                before[b].add(a)
                break
        else:
            if len(first) > len(second):
                return ""
    WHITE, GRAY, BLACK = 0, 1, 2
    color = {ch: WHITE for ch in before}
    order = []

    def visit(ch):
        color[ch] = GRAY
        for earlier in before[ch]:
            if color[earlier] == GRAY:
                return False  # cycle
            if color[earlier] == WHITE and not visit(earlier):
                return False
        color[ch] = BLACK
        order.append(ch)  # everything that must precede ch is already placed
        return True

    for ch in before:
        if color[ch] == WHITE and not visit(ch):
            return ""
    return "".join(order)`,
					TimeComplexity:  "O(C)",
					SpaceComplexity: "O(1)",
					Explanation:     "Store, for each letter, the letters that must come before it, and DFS through them. Appending a letter after all its predecessors (post-order) yields a valid order directly; an edge into a GRAY letter is a cycle.",
				},
				{
					Name: "Try every ordering",
					Code: `from itertools import permutations

def alienOrder(words):
    letters = sorted({ch for word in words for ch in word})
    for order in permutations(letters):
        rank = {ch: i for i, ch in enumerate(order)}
        keys = [[rank[ch] for ch in word] for word in words]
        if all(keys[i] <= keys[i + 1] for i in range(len(keys) - 1)):
            return "".join(order)
    return ""`,
					TimeComplexity:  "O(L! * C) for L distinct letters",
					SpaceComplexity: "O(C)",
					Explanation:     "The definition as code: for each permutation of the letters, map each word to its list of ranks and check that the lists are in non-decreasing order (Python compares lists lexicographically, including the prefix rule). Only usable for a handful of letters, but it shares no logic with the graph solution, so it is a strong check.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it enough to compare only adjacent words, not every pair?", Answer: "Lexicographic order under a fixed alphabet is transitive: if w1 <= w2 and w2 <= w3 then w1 <= w3. So if every adjacent pair is in order, the whole list is sorted. Rules from non-adjacent pairs are implied by the adjacent ones and add nothing."},
			{Prompt: "Why must [\"abc\", \"ab\"] return \"\" even though it gives no letter rule at all?", Answer: "The two words never differ, so the prefix rule decides their order: the shorter word, ab, must come first in every alphabet. The list puts abc first, so no letter order can make it sorted. The graph never sees this because no edge is produced, which is why the check happens while reading the pairs."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, AdvancedGraphsProblems...)
}
