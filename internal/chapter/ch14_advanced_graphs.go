package chapter

import "time"

// chapter14 teaches weighted graphs: Dijkstra's algorithm (and why a popped
// node is final), Bellman-Ford with bounded rounds, minimum spanning trees
// (Prim and Kruskal) and Eulerian paths with Hierholzer's algorithm.
// Storyboards (IDs registered in app.js): dijkstra.
func chapter14() Chapter {
	return Chapter{
		Number:        14,
		Slug:          "advanced-graphs",
		Title:         "Shortest Paths and Spanning Trees",
		Category:      "advanced-graphs",
		Prerequisites: []int{13, 11},
		Summary:       "When edges have different costs, fewest edges is no longer cheapest. Learn Dijkstra's algorithm (BFS with a min-heap) and why it needs non-negative weights, Bellman-Ford for limited hops and negative edges, Prim and Kruskal for the cheapest network that connects everything, and Hierholzer's algorithm for using every edge exactly once.",
		Motivation: `A courier leaves the depot A and must know the fastest time to every other stop. The roads are one-way with travel times: A->B 4, A->C 2, B->C 1, B->D 5, C->D 8, C->E 10, D->E 2.

BFS from Chapter 13 counts edges, so it says D is two roads away and happily reports A->C->D, which takes 2 + 8 = 10 minutes. But A->B->D takes 4 + 5 = 9. Fewest roads is not fastest.

The brute force is to try every route. On a map of 20 stops where every stop has a road to every other, the number of simple routes between two stops is more than 10^16; even at a billion routes per second that takes over six months. And most of that work is wasted: once you know the fastest way to reach C, every route that reaches C more slowly is useless, no matter what comes after.

Dijkstra's algorithm uses exactly that observation. It settles stops one at a time in order of their travel time, using the min-heap from Chapter 11 to find the next one, and never revisits a settled stop. For this map it settles A (0), C (2), B (4), D (9), E (11) after about a dozen heap operations. The rest of the chapter adapts the same idea: limit the number of hops (Bellman-Ford), connect everything as cheaply as possible (minimum spanning trees), or use every road exactly once (Eulerian paths).`,
		Objectives: []string{
			"Trace Dijkstra's algorithm with a min-heap, writing the heap and the finalised distances after every pop, including stale entries that are skipped.",
			"Explain why a node's distance is final when it is popped, and give a small graph with a negative edge where that argument fails.",
			"Implement Bellman-Ford with a fixed number of rounds, and explain why each round must read the previous round's distances when hops are limited.",
			"Define a minimum spanning tree, state the cut property, and implement Prim's algorithm with a heap and Kruskal's algorithm with union-find.",
			"Recognise an 'use every edge exactly once' problem and implement Hierholzer's algorithm, explaining why recording nodes in post-order works.",
		},
		WhenToUse: []string{
			"Edges carry different costs (times, prices, distances) and you need the cheapest or fastest route from a source: Dijkstra, when costs are non-negative.",
			"The cost of a route is its worst edge or highest cell rather than a sum (\"minimum effort\", \"lowest water level\"): Dijkstra with max instead of +.",
			"A limit on the number of edges (\"at most k stops\") or negative edge weights: Bellman-Ford with the right number of rounds.",
			"\"Connect all points/cities with minimum total cost\": a minimum spanning tree with Prim or Kruskal.",
			"\"Use every ticket/edge exactly once\" or \"draw the figure without lifting the pen\": an Eulerian path with Hierholzer's algorithm.",
			"Ordering constraints hidden in data (letters in a sorted alien dictionary): build the graph, then topological sort from Chapter 13.",
		},
		Concepts: []Concept{
			{
				Name:        "Dijkstra's algorithm",
				Description: "Dijkstra finds the cheapest route from one source to every node when all edge weights are non-negative. It is BFS with the FIFO queue replaced by a min-heap of (distance, node): always expand the node with the smallest tentative distance, finalise it, and relax its outgoing edges.",
				WhyItWorks:  "Claim: when a node x is popped for the first time with distance d, d is its true shortest distance. Suppose a shorter route to x existed. It starts at the source, which is finalised, and ends at x, which is not, so it has a first node y that was not finalised when x was popped. The node just before y on that route was finalised, and when it was, it pushed y with the route's distance up to y. All weights are non-negative, so that distance is at most the whole route's length, which is less than d. Then y would have come out of the heap before x: a contradiction. So the popped distance is final. The argument needs every edge to add zero or more; one negative edge lets a later detour undercut an already-finalised node. Pushing a new entry on every improvement and skipping entries for finalised nodes (lazy deletion) keeps the code simple: at most one push per edge, so O(E log E) time.",
				Intuition:   "Water flooding a network of pipes from one tap, where pipe length is the time water takes to pass through it. Water reaches junctions in order of their distance from the tap; the moment water first arrives at a junction, nothing can get there sooner, because any other way in comes through pipes the water has not finished filling. The heap is the list of junctions about to get wet, sorted by when.",
				CoreIdeas: []string{
					"Subgoals: Build the graph, Seed the heap, Pop the cheapest, Skip stale entries, Relax the edges, Read off the answer.",
					"Finalise on pop, not on push: a pushed distance is only a tentative offer.",
					"Lazy deletion: a node may sit in the heap several times; when a stale copy is popped, skip it.",
					"Relaxing an edge u -> v with weight w means offering v the distance dist[u] + w.",
					"The combine step can change: Swim in Rising Water uses max(level, height) instead of +. Dijkstra stays correct as long as extending a route never makes it cheaper.",
					"Complexity O(E log E) = O(E log V) with a binary heap.",
				},
				CommonMistakes: []string{
					"Finalising a node when it is pushed. A cheaper route may still arrive; finalise when popped.",
					"Forgetting the stale check, which re-expands nodes and can produce wrong answers in variants that count things on pop.",
					"Using Dijkstra with negative weights. Use Bellman-Ford instead.",
					"Taking max(dist) of a dictionary, which is the largest key, not the largest distance. Use max(dist.values()).",
				},
				Examples: []CodeExample{
					{
						ID:       "dijkstra",
						Title:    "Dijkstra with a heap and lazy deletion",
						Language: "python",
						Snippet: `import heapq

def dijkstra(graph, start):  # graph[u] = [(v, weight), ...]
    dist = {}  # finalised distances only
    heap = [(0, start)]
    while heap:
        d, u = heapq.heappop(heap)
        if u in dist:
            continue  # stale entry
        dist[u] = d  # smallest in the heap, so final
        for v, w in graph[u]:
            if v not in dist:
                heapq.heappush(heap, (d + w, v))  # relax u -> v
    return dist`,
						Notes: "Nodes missing from dist are unreachable.",
					},
				},
			},
			{
				Name:        "Bellman-Ford and limited hops",
				Description: "Bellman-Ford relaxes every edge in rounds. After round i, each node's distance is at most the cheapest route that uses i edges or fewer. It handles negative weights, detects negative cycles, and, stopped after k + 1 rounds, answers 'cheapest with at most k stops'.",
				WhyItWorks:  "Invariant: after round i, dist[v] is no more than the cheapest route to v with at most i edges. Round i + 1 tries extending every route by one edge, so any best route with i + 1 edges (a best route with i edges to its second-to-last node, plus the last edge) is considered. Without negative cycles a cheapest route repeats no node, so it has at most n - 1 edges, and n - 1 rounds suffice; if round n still improves something, a negative cycle exists. When the number of edges is limited, the round must read only the previous round's values: updating in place lets an edge relaxed early in a round be extended again later in the same round, so one round could add several edges and break the 'at most i edges' meaning.",
				Intuition:   "Rumours spreading through a town, one day at a time. On day i, everyone knows the cheapest rumour that took at most i retellings. If you are only allowed k + 1 retellings, you stop after day k + 1 and must make sure that news heard today is not passed on again until tomorrow, which is what the copy of the array does.",
				CoreIdeas: []string{
					"Each round: for every edge (u, v, w), if dist[u] + w < dist[v], lower dist[v].",
					"n - 1 rounds for unlimited hops; k + 1 rounds for 'at most k stops'.",
					"With a hop limit, read from the previous round's array and write into a copy.",
					"An improvement in round n proves a negative cycle reachable from the source.",
					"Cost O(rounds * E): slower than Dijkstra, but it allows negative weights and hop limits.",
				},
				CommonMistakes: []string{
					"Updating in place when hops are limited, so a single round chains several edges.",
					"Running k rounds for 'k stops'; k stops means k + 1 edges.",
					"Adding to infinity with integers like 10**9 that can overflow in other languages; in Python, float('inf') + w is still infinity, which keeps comparisons safe.",
				},
				Examples: []CodeExample{
					{
						ID:       "bellman_ford_k",
						Title:    "Cheapest price with at most k + 1 edges",
						Language: "python",
						Snippet: `def cheapest_within(n, edges, src, dst, k):
    INF = float("inf")
    cost = [INF] * n
    cost[src] = 0
    for _ in range(k + 1):
        new_cost = cost[:]  # read last round, write this round
        for u, v, w in edges:
            if cost[u] + w < new_cost[v]:
                new_cost[v] = cost[u] + w
        cost = new_cost
    return cost[dst] if cost[dst] < INF else -1`,
						Notes: "Without the copy, round one could already reach dst through a long chain.",
					},
				},
			},
			{
				Name:        "Minimum spanning trees: Prim and Kruskal",
				Description: "A spanning tree of a connected, undirected, weighted graph is a set of n - 1 edges that connects all n nodes with no cycle. A minimum spanning tree (MST) is one with the smallest total weight: the cheapest way to wire every house to the network. Prim grows one tree from a start node by repeatedly adding the cheapest edge leaving it; Kruskal sorts all edges and adds each one that joins two different components.",
				WhyItWorks:  "Both rest on the cut property: split the nodes into two non-empty groups; the cheapest edge crossing between them belongs to some MST. Proof by exchange: take any MST T without that edge e. Adding e to T creates a cycle, which must cross the split a second time through some edge f. Since e is the cheapest crossing edge, swapping f for e gives a spanning tree that costs no more, so it is also minimum and contains e. Prim's split is 'in the tree' versus 'not yet'; Kruskal's is 'the component containing one end' versus the rest, and union-find tells whether the ends are already in one component (then the edge would close a cycle, so it is skipped). Each greedy step adds a safe edge, so after n - 1 steps the result is an MST.",
				Intuition:   "Prim is a city growing its road network outward: each year it builds the cheapest road from the city to any village not yet connected. Kruskal is a planner with a list of all possible roads sorted by price, approving each road unless the two places it joins are already connected.",
				CoreIdeas: []string{
					"An MST has exactly n - 1 edges; a cheapest connecting set never needs a cycle.",
					"Prim with a heap looks like Dijkstra, but the heap key is the single edge's cost, not the total distance from the start.",
					"Kruskal = sort edges + union-find from Chapter 13; skip an edge whose ends share a root.",
					"Dense graphs (every pair is an edge): array-based Prim in O(V^2) beats both heap Prim and Kruskal.",
					"An MST is not a shortest-path tree: it minimises the total of all edges, not the distance from one node.",
				},
				CommonMistakes: []string{
					"Pushing the running total in Prim (Dijkstra's key) instead of the edge cost.",
					"Counting an edge's cost when it is pushed instead of when it is popped and accepted.",
					"In Kruskal, adding an edge without checking whether both ends are already connected, which builds cycles and overpays.",
				},
				Examples: []CodeExample{
					{
						ID:       "prim",
						Title:    "Prim's algorithm with a heap",
						Language: "python",
						Snippet: `import heapq

def mst_cost(n, graph):  # graph[u] = [(v, cost), ...], undirected
    in_tree = set()
    heap = [(0, 0)]  # (edge cost to attach node, node)
    total = 0
    while heap and len(in_tree) < n:
        cost, u = heapq.heappop(heap)
        if u in in_tree:
            continue  # stale
        in_tree.add(u)
        total += cost
        for v, c in graph[u]:
            if v not in in_tree:
                heapq.heappush(heap, (c, v))  # edge cost, not cost + c
    return total if len(in_tree) == n else -1  # -1: graph not connected`,
						Notes: "Compare with Dijkstra above: only the pushed key differs.",
					},
				},
			},
			{
				Name:        "Eulerian paths and Hierholzer's algorithm",
				Description: "An Eulerian path uses every edge of a graph exactly once. In a directed graph it exists when every node has in-degree equal to out-degree, except possibly a start node with one extra out-edge and an end node with one extra in-edge, and all edges are connected. Hierholzer's algorithm finds one in O(E): walk along unused edges, and record a node in post-order, when it has no unused edges left; the reversed record is the path.",
				WhyItWorks:  "Because the degrees balance, a walk that keeps taking unused edges can only get stuck at the path's end node (any other node you enter still has an unused edge out). The node where you get stuck must therefore be the last stop of everything that remains, so it is written to the record first. Then the walk backs up to the previous node and, if that node still has unused edges, explores them: they form a loop that returns to it, and since the loop is recorded before the node itself, it ends up spliced in exactly where it belongs once the record is reversed. Every edge is consumed once, giving O(E) plus any sorting.",
				Intuition:   "Planning a trip where you must use every ticket. If you charge ahead and land somewhere with no tickets out, that city must be the last stop of the trip, so you write it at the bottom of the page and then fill in the page upwards as you back out and use the tickets you skipped.",
				CoreIdeas: []string{
					"Consume edges (remove a ticket when used); do not mark nodes, which may be visited many times.",
					"Record a node after its outgoing edges are exhausted (post-order), then reverse.",
					"To get the lexicographically smallest path, try destinations in sorted order; sorting in reverse and popping from the end gives O(1) removals.",
					"Backtracking (Chapter 12) also works but can be exponential; Hierholzer never undoes a step.",
				},
				CommonMistakes: []string{
					"Recording nodes in pre-order (when you arrive), which can put a dead end in the middle of the route.",
					"Using a visited set for nodes, which forbids the repeated visits an Eulerian path needs.",
					"Storing destinations in a set, which loses duplicate edges.",
				},
				Examples: []CodeExample{
					{
						ID:       "hierholzer",
						Title:    "Hierholzer's algorithm",
						Language: "python",
						Snippet: `from collections import defaultdict

def euler_path(edges, start):
    out = defaultdict(list)
    for u, v in sorted(edges, reverse=True):
        out[u].append(v)  # pop() returns the smallest
    path = []

    def walk(u):
        while out[u]:
            walk(out[u].pop())  # use the edge once
        path.append(u)  # post-order

    walk(start)
    return path[::-1]`,
						Notes: "The iterative version keeps the walk on an explicit stack: push while the top has edges, otherwise pop into path.",
					},
				},
			},
		},
		Animations: []Storyboard{
			{
				ID:    "dijkstra",
				Title: "Dijkstra's algorithm",
				Goal:  "Watch tentative distances fall as edges are relaxed, and see why the cheapest node in the heap can be finalised.",
				MemoryTips: []string{
					"Build the graph: adjacency lists of (neighbour, weight).",
					"Seed the heap with (0, start); nothing is final yet.",
					"Pop the cheapest; skip stale entries for nodes already final.",
					"Finalise it: with non-negative weights no later route can be cheaper.",
					"Relax the edges: push (distance + weight, neighbour) for every neighbour not yet final.",
					"Read off the answer from the finalised distances.",
				},
				Steps: []StoryboardStep{
					{Cue: "init", Narration: "Seed the heap. Every node starts at infinity except the source A at 0, so the heap holds just (0, A). Nothing is final yet.", VisualHint: "Distance table: A 0, all others infinity", Duration: Duration(3 * time.Second), CodeRef: "dijkstra: init dist/parent"},
					{Cue: "pick-node", Narration: "Pop the cheapest node in the heap. After A, the heap holds C at 2 and B at 4, so C comes next even though B was pushed first. If a popped node is already final, the entry is stale and is skipped.", VisualHint: "Highlight the smallest entry leaving the heap", Duration: Duration(3 * time.Second), CodeRef: "dijkstra: pop heap"},
					{Cue: "relax", Narration: "Relax the edges: offer each neighbour the distance through this node. From C: D gets 2 + 8 = 10 and E gets 2 + 10 = 12. Later, from B at 4, D improves to 4 + 5 = 9; the old offer of 10 stays in the heap as a stale entry.", VisualHint: "Edge flashes and the neighbour's distance drops when the offer is cheaper", Duration: Duration(4 * time.Second), CodeRef: "dijkstra: relax edge"},
					{Cue: "finalize", Narration: "Finalise and repeat. A popped node's distance can never improve: any other route would leave the finished nodes through something at least as expensive and then add non-negative weights. Final distances: A 0, C 2, B 4, D 9, E 11 (via D, 9 + 2).", VisualHint: "Node moves into the finished set; table entry turns solid", Duration: Duration(3 * time.Second), CodeRef: "dijkstra: visited set"},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Fastest delivery times (Dijkstra)",
				Problem: "One-way roads with travel times: A->B 4, A->C 2, B->C 1, B->D 5, C->D 8, C->E 10, D->E 2. Return the fastest time from A to every stop.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Build the graph",
						Explanation: "Store each stop's outgoing roads as (neighbour, time) pairs. Roads are one-way, so each appears once.",
						Code:        `graph = {"A": [("B", 4), ("C", 2)], "B": [("C", 1), ("D", 5)], "C": [("D", 8), ("E", 10)], "D": [("E", 2)], "E": []}`,
						State:       "5 stops, 7 directed roads",
					},
					{
						Subgoal:     "Seed the heap",
						Explanation: "The source costs 0. dist will only ever hold final distances, so it starts empty.",
						Code: `dist = {}
heap = [(0, start)]`,
						State: "heap = [(0, A)]   dist = {}",
					},
					{
						Subgoal:     "Pop the cheapest",
						Explanation: "Take the smallest tentative time. If that stop is already final, this entry is stale (a slower offer made earlier), so skip it. Otherwise its time is final.",
						Code: `time, node = heapq.heappop(heap)
if node in dist:
    continue
dist[node] = time`,
						State: "pops in order: A 0, C 2, B 4, D 9, D 10 (stale, skipped), E 11, E 12 (stale, skipped)",
					},
					{
						Subgoal:     "Relax the edges",
						Explanation: "Offer each neighbour that is not final the time through this stop. The heap keeps the best offer on top.",
						Code: `for nxt, weight in graph[node]:
    if nxt not in dist:
        heapq.heappush(heap, (time + weight, nxt))`,
						State: "from A: push (4,B), (2,C). from C: push (10,D), (12,E). from B: C is final, push (9,D). from D: push (11,E)",
					},
					{
						Subgoal:     "Read off the answer",
						Explanation: "When the heap is empty every reachable stop is final. A stop not in dist would be unreachable.",
						Code:        `return dist`,
						State:       "dist = {A: 0, C: 2, B: 4, D: 9, E: 11}. BFS would have said D via C (10) and E via C (12).",
					},
				},
				Code: `import heapq

def shortest_times(graph, start):
    dist = {}
    heap = [(0, start)]
    while heap:
        time, node = heapq.heappop(heap)
        if node in dist:
            continue
        dist[node] = time
        for nxt, weight in graph[node]:
            if nxt not in dist:
                heapq.heappush(heap, (time + weight, nxt))
    return dist`,
			},
			{
				Title:   "Cheapest cable network (Prim's MST)",
				Problem: "Four houses 0..3 can be joined by cables (undirected) with costs 0-1: 1, 0-2: 4, 1-2: 2, 1-3: 6, 2-3: 3. Return the cheapest total cost that connects all houses, or -1 if impossible.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Build the graph",
						Explanation: "Cables work both ways, so each one goes into both houses' lists.",
						Code: `graph = [[] for _ in range(n)]
for a, b, cost in edges:
    graph[a].append((b, cost))
    graph[b].append((a, cost))`,
						State: "graph[1] = [(0,1), (2,2), (3,6)], and so on",
					},
					{
						Subgoal:     "Seed the heap",
						Explanation: "Start the tree at house 0; attaching the first house is free.",
						Code: `in_tree = set()
heap = [(0, 0)]
total = 0`,
						State: "heap = [(0, 0)]   in_tree = {}   total = 0",
					},
					{
						Subgoal:     "Pop the cheapest",
						Explanation: "The heap gives the cheapest cable that would attach some house to the tree. If that house is already in, the entry is stale. Otherwise the cut property says this cable is safe: add the house and pay for the cable.",
						Code: `cost, node = heapq.heappop(heap)
if node in in_tree:
    continue
in_tree.add(node)
total += cost`,
						State: "add 0 (cost 0, total 0); add 1 (cost 1, total 1); add 2 (cost 2, total 3); add 3 (cost 3, total 6)",
					},
					{
						Subgoal:     "Push the crossing edges",
						Explanation: "The new house offers its cables to houses outside the tree. The key is the cable's own cost, not a running total; that is the one difference from Dijkstra.",
						Code: `for nxt, edge_cost in graph[node]:
    if nxt not in in_tree:
        heapq.heappush(heap, (edge_cost, nxt))`,
						State: "after 0: heap [(1,1), (4,2)]; after 1: [(2,2), (4,2), (6,3)]; after 2: [(3,3), (4,2), (6,3)]",
					},
					{
						Subgoal:     "Read off the answer",
						Explanation: "If every house joined, total is the MST cost. If the heap ran dry first, some house is unreachable.",
						Code:        `return total if len(in_tree) == n else -1`,
						State:       "total = 6 using cables 0-1, 1-2, 2-3; the leftover entries (4,2) and (6,3) were never needed",
					},
				},
				Code: `import heapq

def mst_cost(n, edges):
    graph = [[] for _ in range(n)]
    for a, b, cost in edges:
        graph[a].append((b, cost))
        graph[b].append((a, cost))
    in_tree = set()
    heap = [(0, 0)]
    total = 0
    while heap and len(in_tree) < n:
        cost, node = heapq.heappop(heap)
        if node in in_tree:
            continue
        in_tree.add(node)
        total += cost
        for nxt, edge_cost in graph[node]:
            if nxt not in in_tree:
                heapq.heappush(heap, (edge_cost, nxt))
    return total if len(in_tree) == n else -1`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Bellman-Ford with negative edges",
				Prompt: "Return the shortest distances from start in a directed graph whose edges (a, b, w) may have negative weights, or None if a negative cycle is reachable. Use rounds of edge relaxation. Fill in the blanks.",
				Code: `def bellman_ford(n, edges, start):
    INF = float("inf")
    dist = [INF] * n
    dist[start] = 0
    for _ in range(____):
        for a, b, w in edges:
            if ____:
                dist[b] = dist[a] + w
    for a, b, w in edges:
        if dist[a] + w < dist[b]:
            return ____  # still improving: negative cycle
    return dist`,
				Blanks:      []string{"n - 1", "dist[a] + w < dist[b]", "None"},
				Explanation: "A shortest route without a negative cycle repeats no node, so it has at most n - 1 edges and n - 1 rounds are enough. Here updating in place is fine because there is no hop limit; it only makes convergence faster. If a further pass still finds an improvement, some cycle keeps lowering the cost forever. On edges (0,1,2), (0,2,3), (2,1,-2) the answer is [0, 1, 3]; Dijkstra would have finalised node 1 at 2 too early.",
			},
			{
				Title:  "Kruskal's MST with union-find",
				Prompt: "Return the total cost of a minimum spanning tree of a connected undirected graph with edges (a, b, cost). Sort the edges and keep each one that joins two different components. Fill in the blanks.",
				Code: `def kruskal(n, edges):
    parent = list(range(n))

    def find(x):
        if parent[x] != x:
            parent[x] = ____
        return parent[x]

    total = 0
    for cost, a, b in sorted(____):
        root_a, root_b = find(a), find(b)
        if ____:
            parent[root_a] = root_b
            total += cost
    return total`,
				Blanks:      []string{"find(parent[x])", "(cost, a, b) for a, b, cost in edges", "root_a != root_b"},
				Explanation: "Putting cost first in each tuple makes sorted() order edges from cheapest to most expensive. An edge between two different roots joins two components, and by the cut property the cheapest such edge is safe; an edge inside one component would close a cycle, so it is skipped. On the cable example from Prim's worked example this also returns 6.",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Which pattern fits? A grid where each cell has an entry fee from 1 to 9; find the cheapest way from the top-left to the bottom-right moving in four directions.",
				Choices: []string{
					"BFS (Chapter 13), because it finds shortest paths in grids",
					"Dijkstra (this chapter), because moves have different, non-negative costs",
					"Backtracking (Chapter 12) over all paths, keeping the cheapest",
					"A minimum spanning tree of the grid",
				},
				Answer:      1,
				Explanation: "BFS minimises the number of moves, which equals cost only when every move costs the same. With fees from 1 to 9, a longer path through cheap cells can win, so you need Dijkstra. Backtracking is exponential. An MST connects all cells cheaply but says nothing about the cheapest route between two specific cells.",
			},
			{
				Prompt:      "Why does Dijkstra fail with a negative edge? Use the graph 0->1 (2), 0->2 (3), 2->1 (-2), source 0.",
				Explanation: "Dijkstra pops node 1 at distance 2 (smaller than node 2's 3) and finalises it. Then it pops node 2 at 3, and the edge 2->1 offers 3 + (-2) = 1, a better route to a node already declared final. The proof that a popped node is final assumes every remaining edge adds zero or more; a negative edge breaks that. Bellman-Ford handles it: it returns distances [0, 1, 3].",
			},
			{
				Prompt: "Trace it. Run heap-based Dijkstra from A on A->B 1, A->C 4, B->C 2, C->D 1. In which order are nodes finalised, and with which distances?",
				Choices: []string{
					"A 0, B 1, C 4, D 5",
					"A 0, B 1, C 3, D 4",
					"A 0, C 4, B 1, D 5",
				},
				Answer:      1,
				Explanation: "Pop A (0): push B 1, C 4. Pop B (1): push C 1 + 2 = 3. Pop C (3), which beats the older offer C 4: push D 3 + 1 = 4. The heap now holds (4, C) and (4, D); the stale (4, C) comes out first and is skipped, then D is finalised at 4. Finalising C at 4 as soon as it was pushed would be the 'finalise on push' mistake.",
			},
			{
				Prompt: "From Chapter 13: why does BFS give the shortest path in an unweighted graph?",
				Choices: []string{
					"Because it explores neighbours in sorted order",
					"Because its FIFO queue processes nodes in non-decreasing order of edge count, so a node is first reached through a fewest-edge path",
					"Because it marks nodes visited, which removes all longer paths from the graph",
				},
				Answer:      1,
				Explanation: "The queue holds nodes of distance d followed by nodes of distance d + 1, never anything else, so nodes are discovered ring by ring. That is exactly Dijkstra's invariant in the special case where every edge weighs 1, which is why a plain queue can replace the heap there.",
			},
			{
				Prompt: "From Chapter 11: a heap holds tuples (distance, node). What does heapq.heappop return, and what does it cost?",
				Choices: []string{
					"The most recently pushed tuple, in O(1)",
					"The tuple with the smallest distance (ties broken by node), in O(log n)",
					"The tuple with the largest distance, in O(log n)",
				},
				Answer:      1,
				Explanation: "Python's heapq is a min-heap and compares tuples element by element, so the smallest distance comes out first, and on a tie the node values are compared. Each push and pop costs O(log n), which is where Dijkstra's log factor comes from.",
			},
			{
				Prompt:      "Is a minimum spanning tree the same as the tree of shortest paths from node 0? Explain with an example.",
				Explanation: "No. Take a triangle with edges 0-1 cost 2, 0-2 cost 2, 1-2 cost 1. The shortest paths from 0 use edges 0-1 and 0-2 (total 4), because each of 1 and 2 is 2 away directly. The MST uses 1-2 and one of the 2-edges (total 3), which is cheaper overall but makes the path from 0 to the far node cost 3. MSTs minimise the sum of all chosen edges; shortest-path trees minimise each node's distance from the source.",
			},
			{
				Prompt:      "In 'cheapest flights within k stops', why must each Bellman-Ford round read from a copy of the previous round's prices?",
				Explanation: "Round i is supposed to allow one more flight than round i - 1. If you write improvements into the same array you are reading, a city improved earlier in the round can be extended again later in that round, so a single round may chain several flights and the stop limit is silently exceeded. Reading old values and writing new ones keeps exactly one extra flight per round.",
			},
		},
	}
}
