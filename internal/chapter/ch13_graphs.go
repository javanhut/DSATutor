package chapter

import "time"

// chapter13 teaches graph traversal: how to store a graph, BFS for fewest
// steps (including multi-source BFS), DFS and flood fill for components,
// topological sort with Kahn's algorithm and three-colour DFS, and union-find.
// Storyboards (IDs registered in app.js): bfs, dfs, topsort.
func chapter13() Chapter {
	return Chapter{
		Number:        13,
		Slug:          "graphs",
		Title:         "Graphs: BFS, DFS, Topological Sort, Union-Find",
		Category:      "graphs",
		Prerequisites: []int{9, 8},
		Summary:       "A graph is a set of nodes joined by edges; grids, friendships, course prerequisites and word puzzles are all graphs in disguise. Learn to store one as adjacency lists, explore it with BFS (fewest steps) and DFS (reach everything), order it with topological sort, and track groups with union-find, always visiting each node once.",
		Motivation: `A social network has 10,000 people and 50,000 friendships. Is Ana connected to Ben through some chain of friends, and how short is the shortest chain?

The slow way is to try chains. Ana has about 10 friends, each of them has about 10 more, so there are roughly 10^6 chains of length 6 starting at Ana, and many of them reach the same people again and again (Ana's friend's friend is often Ana's friend too). Worse, chains can loop forever: Ana -> Cleo -> Ana -> Cleo...

The fix is one small idea that the whole chapter is built on: remember who you have already visited, and never visit anyone twice. Breadth-first search does exactly that and explores in rings, everyone 1 step away, then everyone 2 steps away, so the first time it meets Ben, the number of rings is the shortest chain. It touches each person once and each friendship twice: about 10,000 + 2 * 50,000 = 110,000 steps instead of millions.

Trees from Chapter 9 were graphs with no cycles and one path between any two nodes, which is why tree recursion never needed a visited set. General graphs have cycles and many paths, and that changes everything.`,
		Objectives: []string{
			"Build an adjacency list from an edge list, and recognise a grid as an implicit graph whose neighbours are the four adjacent cells.",
			"Trace BFS on a small graph, writing the queue and the visited set after every step, and explain why it finds fewest-step paths in an unweighted graph.",
			"Implement DFS and flood fill (recursive and with an explicit stack) to count connected components, and say when recursion depth becomes a problem.",
			"Use multi-source BFS and reverse search (starting from the targets) to turn many searches into one.",
			"Produce a topological order with Kahn's algorithm, detect a cycle with it or with three-colour DFS, and explain why a cycle makes ordering impossible.",
			"Implement union-find with path compression and union by size from scratch, and explain why an edge between two nodes with the same root closes a cycle.",
		},
		WhenToUse: []string{
			"The input is a grid and the question is about connected regions, flooding, spreading or the fewest steps between cells.",
			"The input is a list of pairs (edges, friendships, prerequisites, \"a must come before b\") and you need groups, reachability or an order.",
			"The words \"fewest moves\", \"minimum number of steps\" or \"shortest transformation\" where every move costs the same: BFS.",
			"Dependencies or ordering constraints, \"can all tasks be finished?\", \"find a valid order\": topological sort and cycle detection.",
			"Edges arrive one at a time and you must know whether two nodes are already connected, or which edge creates a cycle: union-find.",
			"Many starting points spreading at once (all gates, all rotten oranges, all ocean edges): multi-source BFS.",
		},
		Concepts: []Concept{
			{
				Name:        "Graphs and how to store them",
				Description: "A graph is a set of nodes (vertices) and edges between pairs of them. Edges are directed (one-way, like a prerequisite) or undirected (two-way, like a friendship), and may carry a weight (Chapter 14). Most problems hand you an edge list; you turn it into an adjacency list, where adj[u] lists u's neighbours. A grid is a graph you never build: each cell's neighbours are computed on the fly.",
				WhyItWorks:  "Every traversal repeatedly asks 'what are the neighbours of this node?'. An adjacency list answers that in time proportional to the number of neighbours, so visiting every node and scanning every list costs O(V + E) in total: each node once, each edge once (twice if undirected, once from each end). An adjacency matrix answers 'is u joined to v?' in O(1) but takes O(V^2) memory and O(V) to list one node's neighbours, which wastes time on sparse graphs, the common case in interviews.",
				Intuition:   "An adjacency list is a contacts app: each person (node) has their own list of contacts (neighbours). To find everyone reachable from you, you open your contacts, then theirs. Directed means you saved their number but they may not have saved yours; undirected means both of you saved each other.",
				CoreIdeas: []string{
					"Undirected edge [a, b]: append b to adj[a] AND a to adj[b]. Directed edge a -> b: only adj[a].",
					"Grid as a graph: the neighbours of (r, c) are (r+1, c), (r-1, c), (r, c+1), (r, c-1), kept only when inside the grid and passable. Nothing needs to be built.",
					"Implicit graphs: in Word Ladder the nodes are words and edges are one-letter changes; you discover neighbours with a helper (wildcard patterns) instead of storing a list.",
					"A tree (Chapter 9) is a connected graph with n - 1 edges and no cycle. General graphs can have cycles, so every traversal needs a visited set.",
					"Traversal cost is O(V + E) with adjacency lists; for an m x n grid that is O(m * n).",
				},
				CommonMistakes: []string{
					"Adding only one direction for an undirected edge, so half the graph becomes unreachable. Append both ways.",
					"Reversing a directed edge: in a prerequisite pair [a, b] (b before a), the edge is b -> a. Write the meaning in a comment before coding.",
					"Forgetting nodes that have no edges. Create adjacency entries for all n nodes (range(n)), not just the ones that appear in edges.",
				},
				Examples: []CodeExample{
					{
						ID:       "adjacency_list",
						Title:    "Edge list to adjacency list",
						Language: "python",
						Snippet: `def build_graph(n, edges, directed=False):
    adj = [[] for _ in range(n)]  # every node gets a list, even with no edges
    for a, b in edges:
        adj[a].append(b)
        if not directed:
            adj[b].append(a)  # undirected: store both directions
    return adj

# build_graph(4, [[0, 1], [1, 2]]) -> [[1], [0, 2], [1], []]`,
						Notes: "Node 3 has no edges but still gets an empty list, so loops over range(n) never crash.",
					},
				},
			},
			{
				Name:        "Breadth-first search (BFS)",
				Description: "BFS explores a graph in rings around the start: first every node 1 edge away, then every node 2 edges away, and so on. It uses a FIFO queue and a visited set, and it finds the fewest-edge path to every node when all edges cost the same. Starting with several nodes in the queue (multi-source BFS) finds each node's distance to the nearest start.",
				WhyItWorks:  "Invariant: the queue holds nodes in non-decreasing order of distance from the start, and the first and last distances in it differ by at most 1. It holds at the start (just the source at distance 0). Popping a node at distance d and appending its unvisited neighbours at distance d + 1 to the back keeps it true. So nodes leave the queue in order of distance, and the first time a node is discovered it is from a node one ring closer, which makes its recorded distance the shortest. Marking a node visited when it is enqueued means it is enqueued at most once, so BFS costs O(V + E).",
				Intuition:   "Drop a stone in a pond. The ripple reaches everything 1 metre away before anything 2 metres away. The queue is the ripple's edge: new points join at the back, so the ring you are expanding is always the nearest unfinished one. Several stones dropped at once (multi-source) make ripples that grow together, and each point is first touched by the closest stone.",
				CoreIdeas: []string{
					"Subgoals: Build the graph, Seed the frontier, Expand one node, Mark when you discover, Read off the answer.",
					"Use collections.deque and popleft(); list.pop(0) is O(n) per pop.",
					"Mark a node visited when you enqueue it, not when you dequeue it, so it is never queued twice.",
					"To count levels (minutes, moves), process exactly len(queue) nodes per round.",
					"Multi-source BFS: put every source in the queue at distance 0 before the loop. Reverse search: start from the targets when there are fewer of them, or when 'can X reach a target?' is asked for every X.",
					"BFS gives fewest edges, not least total weight. With weights, use Dijkstra (Chapter 14).",
				},
				CommonMistakes: []string{
					"Marking visited only on dequeue. Distances stay correct, but a node can be queued once per neighbour, which multiplies the work (and memory) on dense graphs and grids. Mark on enqueue.",
					"Using DFS when the question asks for the fewest steps. DFS finds some path, not the shortest one.",
					"Running one BFS per source and keeping the minimum, when all sources could start together in a single multi-source BFS.",
				},
				Examples: []CodeExample{
					{
						ID:       "bfs",
						Title:    "BFS distances from a start node",
						Language: "python",
						Snippet: `from collections import deque

def bfs_distances(adj, start):
    dist = {start: 0}  # doubles as the visited set
    queue = deque([start])
    while queue:
        node = queue.popleft()
        for nxt in adj[node]:
            if nxt not in dist:  # mark when you discover
                dist[nxt] = dist[node] + 1
                queue.append(nxt)
    return dist`,
						Notes: "A node absent from dist was never reached. Storing the parent of each node as well lets you rebuild the path itself.",
					},
				},
			},
			{
				Name:        "Depth-first search (DFS) and flood fill",
				Description: "DFS follows one path as deep as it can, then backs up and tries the next branch. It is the natural way to visit every node of a connected component (flood fill), to count components, and to detect cycles. It can be written recursively, like the tree recursion of Chapter 9 plus a visited set, or with an explicit stack.",
				WhyItWorks:  "DFS from a start node visits exactly the nodes reachable from it. Every node it visits is reachable, because it only steps along edges. And every reachable node is visited: if some reachable node were missed, take a path to it and look at the first missed node on that path; its predecessor was visited, and DFS looks at all neighbours of every visited node, so it would have been found. The visited mark makes each node expand once, so the cost is O(V + E). Counting components is therefore: scan all nodes, and every unvisited node you meet starts one new flood.",
				Intuition:   "Exploring a maze with a ball of string and a piece of chalk. You walk forward until you hit a dead end or a chalk-marked room, then follow the string back to the last junction with an unexplored corridor. The chalk is the visited set; the string is the recursion stack.",
				CoreIdeas: []string{
					"Component-counting subgoals: Scan for an unvisited start, Flood the component, Read off the answer.",
					"Mark a node before exploring its neighbours, or two neighbours will bounce back and forth forever.",
					"In a grid you can mark visits in the input itself (sink land to \"0\") when the caller doesn't need it afterwards.",
					"Recursion depth equals the longest path DFS walks, up to m * n on a snake-shaped grid region. The judge raises Python's recursion limit, but an explicit stack (or BFS) avoids the limit entirely.",
					"Order rarely matters for reachability or counting; it does matter for post-order uses such as topological sort and Hierholzer's algorithm (Chapter 14).",
				},
				CommonMistakes: []string{
					"Marking visited after the recursive calls, which recurses forever on any cycle. Mark first, then recurse.",
					"Forgetting to start a DFS from every unvisited node, which counts or checks only the first component.",
					"Using recursion on a 300 x 300 grid of land without thinking about depth: up to 90,000 nested calls. Prefer an explicit stack when regions can be huge.",
				},
				Examples: []CodeExample{
					{
						ID:       "dfs_components",
						Title:    "Count components with an iterative DFS",
						Language: "python",
						Snippet: `def count_components(adj):
    seen = set()
    count = 0
    for start in range(len(adj)):
        if start in seen:
            continue  # already flooded from an earlier start
        count += 1
        seen.add(start)
        stack = [start]
        while stack:
            node = stack.pop()
            for nxt in adj[node]:
                if nxt not in seen:
                    seen.add(nxt)
                    stack.append(nxt)
    return count`,
						Notes: "Swapping the stack for a deque and pop() for popleft() turns this into BFS; the count is the same.",
					},
				},
			},
			{
				Name:        "Topological sort and cycle detection",
				Description: "A topological order of a directed graph lists every node so that each edge u -> v has u before v: a valid order to take courses or run build steps. It exists exactly when the graph has no directed cycle (a DAG). Kahn's algorithm builds the order with in-degrees and a queue; three-colour DFS builds it in post-order and detects cycles as edges back into the current path.",
				WhyItWorks:  "Kahn: a node with in-degree 0 has no unmet prerequisites, so it can go next, and removing it only lowers other in-degrees. In a DAG some remaining node always has in-degree 0 (walk prerequisites backwards from any node; without a cycle the walk must stop at one), so the process outputs every node. On a cycle, every node waits for its predecessor on the cycle, so none ever reaches 0 and the output is short: that is the cycle test. Three-colour DFS: WHITE is unvisited, GRAY is on the current recursion path, BLACK is finished. An edge to a GRAY node goes back to an ancestor on the current path, which closes a cycle; an edge to a BLACK node does not, since everything below a BLACK node is already known to be cycle-free.",
				Intuition:   "Kahn is getting dressed: anything with nothing left to wait for (underwear, socks) can go on now; putting it on may free up something else (trousers, then shoes). If shoes needed socks and socks needed shoes, you would be stuck forever with both still waiting. Three-colour DFS is walking a trail with breadcrumbs: GRAY breadcrumbs mark the path you are on right now; stepping onto your own fresh breadcrumb means you walked in a circle.",
				CoreIdeas: []string{
					"Kahn subgoals: Build the graph, Count in-degrees, Seed the frontier (in-degree 0), Expand one node (decrement successors, queue those that hit 0), Detect a cycle (output shorter than n).",
					"Many topological orders usually exist; any node with in-degree 0 may go next.",
					"Three-colour DFS appends a node after all its successors (post-order). With edges u -> v meaning 'u before v', reverse the post-order; with edges pointing to prerequisites, post-order is already the answer.",
					"A single visited set cannot detect directed cycles: in a diamond, the bottom node is reached twice without any cycle.",
				},
				CommonMistakes: []string{
					"Treating 'already visited' as 'cycle' in a directed graph. Use three colours: only an edge into a GRAY node is a cycle.",
					"Returning a partial order when a cycle exists. If the output is shorter than the number of nodes, the answer is 'impossible'.",
					"Getting the edge direction backwards for pairs like [course, prerequisite], which produces the reversed order.",
				},
				Examples: []CodeExample{
					{
						ID:       "kahn",
						Title:    "Kahn's algorithm",
						Language: "python",
						Snippet: `from collections import deque

def topo_order(n, edges):  # edge (u, v): u must come before v
    adj = [[] for _ in range(n)]
    indegree = [0] * n
    for u, v in edges:
        adj[u].append(v)
        indegree[v] += 1
    ready = deque(v for v in range(n) if indegree[v] == 0)
    order = []
    while ready:
        u = ready.popleft()
        order.append(u)
        for v in adj[u]:
            indegree[v] -= 1
            if indegree[v] == 0:
                ready.append(v)
    return order if len(order) == n else []  # [] means a cycle`,
						Notes: "Each node is queued once and each edge decremented once: O(V + E).",
					},
				},
			},
			{
				Name:        "Union-find (disjoint set union)",
				Description: "Union-find keeps nodes in groups and answers two questions fast: which group is x in (find), and merge the groups of a and b (union). Each group is a tree stored in a parent array; the root is the group's name. It shines when edges arrive one at a time and you need to know whether two nodes are already connected, which is also how you spot the edge that closes a cycle in an undirected graph.",
				WhyItWorks:  "Invariant: two nodes have the same root exactly when the edges added so far connect them. Initially every node is its own root. union(a, b) links root to root (parent[find(b)] = find(a)), which merges exactly those two groups. Path compression re-points nodes on a find path straight to the root; that changes the shape of a tree but not its root, so the invariant survives. Union by size hangs the smaller tree under the larger, so a node's depth grows only when its group at least doubles, giving depth at most log2(n). Together they make each operation O(α(n)) amortised, where α, the inverse Ackermann function, is below 5 for any realistic n. If find(a) == find(b) before adding edge [a, b], a path already joins them, so the new edge would close a cycle.",
				Intuition:   "Clubs with presidents. Everyone knows who to ask ('my parent'), and asking up the chain eventually reaches the club president (the root). Two people are in the same club when they reach the same president. Merging clubs means one president agrees to report to the other. Path compression is everyone you asked along the way writing down the president's number so next time they ask directly.",
				CoreIdeas: []string{
					"Subgoals: Start with singletons (parent[x] = x), Find the roots, Union or detect a cycle (or skip), Read off the answer.",
					"find(x): if parent[x] != x, set parent[x] = find(parent[x]) and return it (path compression).",
					"union: compare roots; if different, attach the smaller-size root under the larger and add the sizes.",
					"Number of groups = n minus the number of successful unions.",
					"Union-find handles undirected connectivity only. For directed cycles use topological sort or three-colour DFS.",
				},
				CommonMistakes: []string{
					"Linking nodes instead of roots (parent[b] = a). If b is not a root, the rest of its tree is left behind and the groups split. Always link find(b) under find(a).",
					"Comparing parent[a] == parent[b] instead of find(a) == find(b); the parent is just the next step up, not the group's name.",
					"Counting a component merge for every edge, even when both ends already share a root.",
				},
				Examples: []CodeExample{
					{
						ID:       "union_find",
						Title:    "Union-find with path compression and union by size",
						Language: "python",
						Snippet: `parent = list(range(n))  # every node is its own root
size = [1] * n

def find(x):
    if parent[x] != x:
        parent[x] = find(parent[x])  # path compression
    return parent[x]

def union(a, b):
    root_a, root_b = find(a), find(b)
    if root_a == root_b:
        return False  # already connected: this edge would close a cycle
    if size[root_a] < size[root_b]:
        root_a, root_b = root_b, root_a
    parent[root_b] = root_a  # smaller tree under the larger
    size[root_a] += size[root_b]
    return True`,
						Notes: "union returns whether it merged anything, which is exactly what Redundant Connection and Graph Valid Tree need.",
					},
				},
			},
		},
		Animations: []Storyboard{
			{
				ID:    "bfs",
				Title: "BFS: exploring in rings",
				Goal:  "See how a FIFO queue makes BFS finish every node at distance d before any node at distance d + 1.",
				MemoryTips: []string{
					"Build the graph: adjacency lists, or the four neighbours of a grid cell.",
					"Seed the frontier: put the start (or every start) in the queue and mark it visited.",
					"Expand one node: popleft the oldest node and look at all its neighbours.",
					"Mark when you discover: mark a neighbour the moment you enqueue it, so it is never queued twice.",
					"Read off the answer: the first time you reach a node, its ring number is its fewest-step distance.",
				},
				Steps: []StoryboardStep{
					{Cue: "enqueue-start", Narration: "Seed the frontier. A goes into the queue and is marked visited at once. It is ring 0: zero steps from itself.", VisualHint: "Queue: [A]; A highlighted as visited", Duration: Duration(2 * time.Second), CodeRef: "bfs: init queue"},
					{Cue: "expand", Narration: "Expand one node at a time from the front of the queue. Popping A discovers B and C (ring 1); popping B discovers D and E (ring 2); popping C discovers F (ring 2). Each new node is marked as it joins the back of the queue, so A, which B also points to, is never queued again.", VisualHint: "Queue grows at the back and shrinks at the front; ring colours spread outward", Duration: Duration(4 * time.Second), CodeRef: "bfs: neighbor loop"},
					{Cue: "goal", Narration: "Read off the answer. Nodes leave the queue in ring order A, B, C, D, E, F, so the first time a target appears, its ring is the fewest steps to reach it. Keeping each node's parent lets you trace that shortest path back to A.", VisualHint: "Path traced back from the goal through parent links", Duration: Duration(3 * time.Second), CodeRef: "bfs: goal check/parent map"},
				},
			},
			{
				ID:    "dfs",
				Title: "DFS: deep first, then back up",
				Goal:  "See how a stack (or recursion) makes DFS follow one branch to the end before trying the next.",
				MemoryTips: []string{
					"Push the start: the stack (or the call stack) holds the path you are on.",
					"Pop, skip if already visited, otherwise mark it and record it.",
					"Push its unvisited neighbours; the last one pushed is explored first.",
					"When a node has nothing new, the stack shrinks: that is backtracking.",
					"Scan for an unvisited start to count components: each new start floods one component.",
				},
				Steps: []StoryboardStep{
					{Cue: "push-start", Narration: "Push the start node A onto the stack. A stack is last in, first out, so whatever we push last is what we explore next.", VisualHint: "Stack: [A]", Duration: Duration(2 * time.Second), CodeRef: "dfs: stack.append(start)"},
					{Cue: "pop", Narration: "Pop the top node. If it is already visited, skip it; otherwise mark it visited and record it. A is popped and visited first.", VisualHint: "Top of the stack lifts off; node turns visited", Duration: Duration(3 * time.Second), CodeRef: "dfs: node = stack.pop()"},
					{Cue: "explore", Narration: "Push the unvisited neighbours (in reverse, so the leftmost is popped first). From A we push C then B, pop B, push E then D, and pop D: we are already two levels deep before C is touched.", VisualHint: "Path A, B, D highlighted going down the left side", Duration: Duration(3 * time.Second), CodeRef: "dfs: push neighbors"},
					{Cue: "backtrack", Narration: "D has no neighbours, so nothing is pushed and the next pop returns to the most recent unexplored branch: E, then F through E. When F's only neighbour is visited, we back up again to C. Final order: A, B, D, E, F, C.", VisualHint: "Stack shrinks; exploration resumes at the last junction", Duration: Duration(2 * time.Second), CodeRef: "dfs: implicit backtrack"},
				},
			},
			{
				ID:    "topsort",
				Title: "Topological sort with Kahn's algorithm",
				Goal:  "See in-degrees fall to zero as prerequisites are taken, and how a cycle would leave nodes stuck.",
				MemoryTips: []string{
					"Build the graph with edges from prerequisite to dependent.",
					"Count in-degrees: how many prerequisites each node still waits for.",
					"Seed the frontier with every node of in-degree 0.",
					"Expand one node: output it, decrement each successor, queue any that reach 0.",
					"Detect a cycle: if fewer than n nodes come out, the rest are stuck on a cycle.",
				},
				Steps: []StoryboardStep{
					{Cue: "calc-indegree", Narration: "Count in-degrees: the number of arrows pointing into each node. A and B have 0, C has 2 (from A and B), D has 1, E has 2, F has 1.", VisualHint: "Each node labelled with its incoming-arrow count", Duration: Duration(3 * time.Second), CodeRef: "topsort: in_degree[v] += 1"},
					{Cue: "init-queue", Narration: "Seed the frontier with every node whose in-degree is 0: A and B wait for nothing, so either can go first.", VisualHint: "Queue: [A, B]", Duration: Duration(2 * time.Second), CodeRef: "topsort: queue zero-degree"},
					{Cue: "process", Narration: "Expand one node: output it and decrement its successors. Taking A drops C to 1. Taking B drops C to 0 and D to 0, so both join the queue. Taking C, then D, drops E to 0; taking E drops F to 0.", VisualHint: "Counters tick down; nodes join the queue when they reach 0", Duration: Duration(4 * time.Second), CodeRef: "topsort: process node"},
					{Cue: "complete", Narration: "The queue is empty and all six nodes came out: A, B, C, D, E, F is a valid order. Had there been a cycle, its nodes would never reach 0 and the output would be shorter than six: that is the cycle test.", VisualHint: "Final order shown left to right", Duration: Duration(2 * time.Second), CodeRef: "topsort: return result"},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Fewest steps through a maze (BFS)",
				Problem: "A maze is a grid of 0 (open) and 1 (wall). Starting at the top-left cell, moving up, down, left or right through open cells, return the fewest steps to reach the bottom-right cell, or -1 if it can't be reached. Trace it on maze = [[0,0,1],[1,0,0],[1,1,0]].",
				Steps: []WorkedStep{
					{
						Subgoal:     "Build the graph",
						Explanation: "The grid is an implicit graph: a cell's neighbours are the four cells next to it that are inside the grid and open. We compute them on the fly instead of building lists. If the start or the end is a wall, no path exists.",
						Code: `rows, cols = len(maze), len(maze[0])
if maze[0][0] == 1 or maze[rows - 1][cols - 1] == 1:
    return -1`,
						State: "rows = 3, cols = 3; start (0,0) and end (2,2) are both open",
					},
					{
						Subgoal:     "Seed the frontier",
						Explanation: "Every move costs the same, so BFS gives fewest steps. Put the start in the queue at distance 0. The dist dictionary doubles as the visited set.",
						Code: `dist = {(0, 0): 0}
queue = deque([(0, 0)])`,
						State: "queue = [(0,0)]   dist = {(0,0): 0}",
					},
					{
						Subgoal:     "Expand one node",
						Explanation: "Pop the oldest cell. If it is the target, its distance is final because BFS pops cells in order of distance. Otherwise look at its four neighbours.",
						Code: `r, c = queue.popleft()
if (r, c) == (rows - 1, cols - 1):
    return dist[(r, c)]
for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):`,
						State: "pop (0,0): (1,0) is a wall, (-1,0) and (0,-1) are outside, (0,1) is open",
					},
					{
						Subgoal:     "Mark when you discover",
						Explanation: "An open, unseen neighbour is one step further than the current cell. Record its distance immediately so no other cell queues it again.",
						Code: `if 0 <= nr < rows and 0 <= nc < cols and maze[nr][nc] == 0 and (nr, nc) not in dist:
    dist[(nr, nc)] = dist[(r, c)] + 1
    queue.append((nr, nc))`,
						State: "dist[(0,1)] = 1, queue = [(0,1)]; pop (0,1) -> dist[(1,1)] = 2; pop (1,1) -> dist[(1,2)] = 3; pop (1,2) -> dist[(2,2)] = 4",
					},
					{
						Subgoal:     "Read off the answer",
						Explanation: "The target (2,2) is popped with distance 4 and returned. If the queue had emptied first, the target would be unreachable, so the function ends with -1.",
						Code:        `return -1`,
						State:       "pop (2,2) -> return 4. For [[0,1],[1,0]] the queue empties after (0,0), so -1.",
					},
				},
				Code: `from collections import deque

def shortest_path(maze):
    rows, cols = len(maze), len(maze[0])
    if maze[0][0] == 1 or maze[rows - 1][cols - 1] == 1:
        return -1
    dist = {(0, 0): 0}
    queue = deque([(0, 0)])
    while queue:
        r, c = queue.popleft()
        if (r, c) == (rows - 1, cols - 1):
            return dist[(r, c)]
        for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
            if 0 <= nr < rows and 0 <= nc < cols and maze[nr][nc] == 0 and (nr, nc) not in dist:
                dist[(nr, nc)] = dist[(r, c)] + 1
                queue.append((nr, nc))
    return -1`,
			},
			{
				Title:   "Ordering tasks with dependencies (Kahn's algorithm)",
				Problem: "Given a list of tasks and pairs (first, then) meaning 'first must be done before then', return the tasks in a valid order, or [] if the dependencies contain a cycle. Trace it on tasks = [shower, dress, coffee, breakfast, leave] with (shower, dress), (coffee, breakfast), (dress, leave), (breakfast, leave).",
				Steps: []WorkedStep{
					{
						Subgoal:     "Build the graph",
						Explanation: "Each pair is a directed edge first -> then: finishing first unlocks then. Give every task an entry, even tasks with no edges.",
						Code: `after = {t: [] for t in tasks}
for first, then in deps:
    after[first].append(then)`,
						State: "after = {shower: [dress], coffee: [breakfast], dress: [leave], breakfast: [leave], leave: []}",
					},
					{
						Subgoal:     "Count in-degrees",
						Explanation: "A task's in-degree is how many tasks it still waits for. It can start exactly when this reaches 0.",
						Code: `indegree = {t: 0 for t in tasks}
for first, then in deps:
    indegree[then] += 1`,
						State: "indegree = {shower: 0, dress: 1, coffee: 0, breakfast: 1, leave: 2}",
					},
					{
						Subgoal:     "Seed the frontier",
						Explanation: "Tasks with in-degree 0 wait for nothing and can go first.",
						Code:        `ready = deque(t for t in tasks if indegree[t] == 0)`,
						State:       "ready = [shower, coffee]",
					},
					{
						Subgoal:     "Expand one node",
						Explanation: "Do the next ready task and record it. Each task it unlocks has one fewer thing to wait for; any that reach 0 become ready.",
						Code: `task = ready.popleft()
order.append(task)
for nxt in after[task]:
    indegree[nxt] -= 1
    if indegree[nxt] == 0:
        ready.append(nxt)`,
						State: "shower -> dress 0, ready [coffee, dress]; coffee -> breakfast 0, ready [dress, breakfast]; dress -> leave 1; breakfast -> leave 0, ready [leave]; leave",
					},
					{
						Subgoal:     "Detect a cycle",
						Explanation: "If every task came out, the order is valid. If some never reached in-degree 0, they wait on each other in a cycle, and no order exists.",
						Code:        `return order if len(order) == len(tasks) else []`,
						State:       "order = [shower, coffee, dress, breakfast, leave], all 5 tasks, so it is returned",
					},
				},
				Code: `from collections import deque

def task_order(tasks, deps):
    after = {t: [] for t in tasks}
    indegree = {t: 0 for t in tasks}
    for first, then in deps:
        after[first].append(then)
        indegree[then] += 1
    ready = deque(t for t in tasks if indegree[t] == 0)
    order = []
    while ready:
        task = ready.popleft()
        order.append(task)
        for nxt in after[task]:
            indegree[nxt] -= 1
            if indegree[nxt] == 0:
                ready.append(nxt)
    return order if len(order) == len(tasks) else []`,
			},
			{
				Title:   "Are two computers on the same network? (union-find)",
				Problem: "There are n computers and a list of cables [a, b]. For each query [a, b], answer whether a and b are connected through cables. Trace it on n = 5, cables = [[0,1],[2,3],[1,3]], queries = [[0,2],[0,4]].",
				Steps: []WorkedStep{
					{
						Subgoal:     "Start with singletons",
						Explanation: "Before any cable, each computer is alone: it is its own root, and its group has size 1.",
						Code: `parent = list(range(n))
size = [1] * n`,
						State: "parent = [0, 1, 2, 3, 4]   size = [1, 1, 1, 1, 1]",
					},
					{
						Subgoal:     "Find the roots",
						Explanation: "Follow parent links up to the root, which names the group. On the way back, point every node on the path straight at the root (path compression), so the next find is one step.",
						Code: `def find(x):
    if parent[x] != x:
        parent[x] = find(parent[x])
    return parent[x]`,
						State: "find(1) with parent = [0, 0, 0, 2, 4]: 1 -> 0, root 0. find(3): 3 -> 2 -> 0, root 0, and parent[3] becomes 0",
					},
					{
						Subgoal:     "Union or skip",
						Explanation: "Find both roots. If they match, the cable adds nothing new (it would close a cycle). Otherwise hang the smaller tree under the larger root.",
						Code: `root_a, root_b = find(a), find(b)
if root_a == root_b:
    return False
if size[root_a] < size[root_b]:
    root_a, root_b = root_b, root_a
parent[root_b] = root_a
size[root_a] += size[root_b]
return True`,
						State: "[0,1]: parent = [0,0,2,3,4]; [2,3]: parent = [0,0,2,2,4]; [1,3]: roots 0 and 2, sizes 2 and 2, so parent[2] = 0 -> parent = [0,0,0,2,4], size[0] = 4",
					},
					{
						Subgoal:     "Read off the answer",
						Explanation: "Two computers are connected exactly when they have the same root.",
						Code:        `return [find(a) == find(b) for a, b in queries]`,
						State:       "[0,2]: roots 0 and 0 -> True; [0,4]: roots 0 and 4 -> False. Answer [True, False]",
					},
				},
				Code: `def network_queries(n, cables, queries):
    parent = list(range(n))
    size = [1] * n

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    def union(a, b):
        root_a, root_b = find(a), find(b)
        if root_a == root_b:
            return False
        if size[root_a] < size[root_b]:
            root_a, root_b = root_b, root_a
        parent[root_b] = root_a
        size[root_a] += size[root_b]
        return True

    for a, b in cables:
        union(a, b)
    return [find(a) == find(b) for a, b in queries]`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Count friend groups (DFS flood)",
				Prompt: "n people are labelled 0..n-1 and each pair in friends is a mutual (undirected) friendship. Count the friend groups. Use the subgoals Build the graph, Scan for an unvisited start, Flood the component, Read off the answer. Fill in the blanks.",
				Code: `def count_groups(n, friends):
    graph = [[] for _ in range(n)]
    for a, b in friends:
        graph[a].append(b)
        ____
    seen = [False] * n
    groups = 0
    for start in range(n):
        if ____:
            continue
        groups += 1
        seen[start] = True
        stack = [start]
        while stack:
            person = stack.pop()
            for friend in graph[person]:
                if not seen[friend]:
                    ____
                    stack.append(friend)
    return groups`,
				Blanks:      []string{"graph[b].append(a)", "seen[start]", "seen[friend] = True"},
				Explanation: "Friendship is undirected, so the edge goes into both lists. A start that is already seen belongs to a group counted earlier. Marking a friend when it is pushed (mark when you discover) keeps it from being pushed again by another friend. count_groups(6, [(0,1),(1,2),(3,4)]) returns 3: {0,1,2}, {3,4} and {5}.",
			},
			{
				Title:  "Detect a directed cycle (three-colour DFS)",
				Prompt: "Return True if the directed graph with nodes 0..n-1 and edges (a, b) meaning a -> b contains a cycle. WHITE means unvisited, GRAY means on the current path, BLACK means finished. Fill in the blanks.",
				Code: `def has_cycle(n, edges):
    graph = [[] for _ in range(n)]
    for a, b in edges:
        graph[a].append(b)
    WHITE, GRAY, BLACK = 0, 1, 2
    color = [WHITE] * n

    def visit(node):
        color[node] = ____
        for nxt in graph[node]:
            if color[nxt] == ____:
                return True
            if color[nxt] == WHITE and visit(nxt):
                return True
        color[node] = ____
        return False

    return any(color[v] == WHITE and visit(v) for v in range(n))`,
				Blanks:      []string{"GRAY", "GRAY", "BLACK"},
				Explanation: "A node turns GRAY when we enter it and BLACK when all its descendants are done. Meeting a GRAY node means we walked back into the current path: a cycle. Meeting a BLACK node is harmless, which is why the diamond 0->1, 0->2, 1->3, 2->3 returns False while 0->1->2->0 returns True.",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Which pattern fits? A knight stands on square (0, 0) of an infinite chessboard. Return the fewest knight moves needed to reach square (x, y).",
				Choices: []string{
					"Backtracking (Chapter 12): try every sequence of moves and keep the shortest",
					"BFS over squares (this chapter): each square is a node, each knight move an edge of cost 1",
					"A heap of squares ordered by distance (Chapter 11), since we want the minimum",
					"Two pointers (Chapter 3) moving toward (x, y) from both sides",
				},
				Answer:      1,
				Explanation: "Fewest moves with every move costing the same is the signature of BFS: the first time BFS reaches (x, y), the ring number is the answer. Backtracking enumerates sequences and revisits squares, which is exponential. A heap is Dijkstra, needed only when moves have different costs; with equal costs a plain queue already pops in distance order.",
			},
			{
				Prompt: "Which pattern fits? Friendships are announced one at a time. After each announcement, report how many separate friend groups there are.",
				Choices: []string{
					"Run a fresh DFS over the whole graph after every announcement",
					"Union-find: each announcement is a union, and the group count drops by one only when the two roots differ",
					"Topological sort of the friendship graph",
				},
				Answer:      1,
				Explanation: "Union-find answers each announcement in near-constant time and keeps the count up to date. Re-running DFS costs O(V + E) per announcement. Topological sort is for directed dependencies, and friendships are undirected.",
			},
			{
				Prompt: "Trace it. parent = [0, 0, 1, 2] encodes the chain 3 -> 2 -> 1 -> 0. What is parent after calling find(3) with path compression (parent[x] = find(parent[x]))?",
				Choices: []string{
					"[0, 0, 1, 2]: find only reads the array",
					"[0, 0, 1, 0]: only node 3 is re-pointed",
					"[0, 0, 0, 0]: every node on the path is re-pointed to the root",
				},
				Answer:      2,
				Explanation: "find(3) calls find(2), which calls find(1), which returns 0. On the way back each call sets its own parent to the root: parent[1] = 0, parent[2] = 0, parent[3] = 0. The whole path is flattened, so later finds on any of these nodes take one step.",
			},
			{
				Prompt:      "Trace it. Run Kahn's algorithm on nodes A, B, C, D with edges A -> C, B -> C, C -> D, starting with ready = [A, B]. Give the in-degrees at the start and the order produced.",
				Explanation: "In-degrees: A 0, B 0, C 2, D 1. Pop A: C drops to 1. Pop B: C drops to 0, so C is queued. Pop C: D drops to 0, so D is queued. Pop D. Order A, B, C, D, all four nodes, so there is no cycle. Starting with ready = [B, A] would give B, A, C, D, which is equally valid.",
			},
			{
				Prompt: "Your BFS marks a node visited when it is popped instead of when it is pushed. On a grid, what goes wrong?",
				Choices: []string{
					"Distances become wrong, because a longer path may be recorded first",
					"Distances are still right, but a cell can be pushed once by each neighbour, multiplying queue size and work",
					"BFS turns into DFS",
					"Nothing at all changes",
				},
				Answer:      1,
				Explanation: "The first pop of a cell still happens at its shortest distance, so answers stay correct. But until it is popped, every neighbour that sees it as unvisited pushes another copy. On a large open grid that can blow up memory and time. Marking on push guarantees one entry per cell.",
			},
			{
				Prompt:      "From Chapter 9 (trees): recursive tree traversal never used a visited set, but graph DFS must. Why?",
				Explanation: "In a tree each node is reached by exactly one path from the root, following child links only, so recursion can never arrive at a node twice and can never loop. A general graph can have cycles and several paths to the same node, and undirected edges point back to where you came from. Without a visited set, DFS revisits nodes and on any cycle recurses forever. A tree is just the special case of a connected graph with n - 1 edges and no cycle.",
			},
			{
				Prompt: "In three-colour DFS on a directed graph, you follow an edge and land on a BLACK node. What does that mean?",
				Choices: []string{
					"A cycle: we have been here before",
					"No cycle through this edge: the BLACK node and everything below it were fully explored and found cycle-free",
					"The graph is disconnected",
				},
				Answer:      1,
				Explanation: "Only GRAY nodes are on the current path; an edge to one of them closes a loop. A BLACK node was finished earlier, which happens whenever two routes share a descendant (the bottom of a diamond). Treating it as a cycle is the classic single-visited-set bug.",
			},
		},
	}
}
