package practice

// AdvancedGraphsDrills is the advanced graphs foundation: Dijkstra on its own before the weighted-graph ladder.
var AdvancedGraphsDrills = []*Problem{
	{
		ID:              "drill-dijkstra",
		Kind:            KindDrill,
		Order:           1,
		Title:           "Dijkstra's Shortest Paths",
		Slug:            "drill-dijkstra",
		Difficulty:      "Medium",
		Category:        "advanced-graphs",
		Tags:            []string{"Graph", "Dijkstra", "Heap", "Shortest Path", "Foundations"},
		RelatedChapters: []int{14, 13},
		Description: `Given a **directed** graph with non-negative edge weights and a ` + "`" + `start` + "`" + ` node, find the cheapest total cost to reach every node from ` + "`" + `start` + "`" + `.

There are ` + "`" + `n` + "`" + ` nodes labelled ` + "`" + `0` + "`" + ` to ` + "`" + `n-1` + "`" + `, and each entry of ` + "`" + `edges` + "`" + ` is a triple ` + "`" + `[u, v, w]` + "`" + `: an edge from ` + "`" + `u` + "`" + ` to ` + "`" + `v` + "`" + ` that costs ` + "`" + `w` + "`" + `, where ` + "`" + `w >= 0` + "`" + `. Return a list of ` + "`" + `n` + "`" + ` numbers where entry ` + "`" + `i` + "`" + ` is the cost of the cheapest path from ` + "`" + `start` + "`" + ` to ` + "`" + `i` + "`" + `, or ` + "`" + `-1` + "`" + ` if node ` + "`" + `i` + "`" + ` cannot be reached. The distance from ` + "`" + `start` + "`" + ` to itself is ` + "`" + `0` + "`" + `.

This is **Dijkstra's algorithm**. The cheapest path is not always the one with the fewest edges, so the plain BFS from earlier in the set no longer works. Two edges between the same pair of nodes may both appear; the cheaper one counts.`,
		Constraints: []string{
			"1 <= n <= 1000",
			"0 <= len(edges) <= 5000",
			"0 <= u, v, start < n",
			"0 <= w <= 1000",
			"No self-loops",
		},
		Examples: []Example{
			{Input: "n = 5, edges = [[0,1,4],[0,2,1],[2,1,2],[1,3,1],[2,3,5]], start = 0", Output: "[0,3,1,4,-1]", Explanation: "Node 1 is cheaper via 0 -> 2 -> 1 (cost 3) than directly (cost 4). Node 3 costs 4 via 1. Node 4 has no incoming edge, so it is -1."},
			{Input: "n = 4, edges = [[0,3,10],[0,1,2],[1,2,2],[2,3,2]], start = 0", Output: "[0,2,4,6]", Explanation: "The direct edge to node 3 costs 10, but 0 -> 1 -> 2 -> 3 costs 6, so the three-edge route wins."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 1, 4}, []interface{}{0, 2, 1}, []interface{}{2, 1, 2}, []interface{}{1, 3, 1}, []interface{}{2, 3, 5}}, "start": 0}, Expected: []interface{}{0, 3, 1, 4, -1}, Note: "example 1: 0 -> 2 -> 1 -> 3 costs 4, cheaper than going 0 -> 1 directly then on"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 3, 10}, []interface{}{0, 1, 2}, []interface{}{1, 2, 2}, []interface{}{2, 3, 2}}, "start": 0}, Expected: []interface{}{0, 2, 4, 6}, Note: "example 2: the one-edge route to node 3 costs 10 but the three-edge route costs 6 (fewest edges is not cheapest)"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1, 0}, []interface{}{1, 2, 5}, []interface{}{0, 2, 6}, []interface{}{2, 3, 0}}, "start": 0}, Expected: []interface{}{0, 0, 5, 5}, Hidden: true, Note: "zero-weight edges: node 1 costs 0 and node 3 costs the same as node 2"},
			{Input: map[string]interface{}{"n": 3, "edges": []interface{}{[]interface{}{1, 0, 5}, []interface{}{1, 2, 1}}, "start": 0}, Expected: []interface{}{0, -1, -1}, Hidden: true, Note: "edges are directed: nothing leaves node 0, so every other node is -1"},
			{Input: map[string]interface{}{"n": 1, "edges": []interface{}{}, "start": 0}, Expected: []interface{}{0}, Hidden: true, Note: "a single node: its distance is 0"},
			{Input: map[string]interface{}{"n": 2, "edges": []interface{}{[]interface{}{0, 1, 9}, []interface{}{0, 1, 3}}, "start": 0}, Expected: []interface{}{0, 3}, Hidden: true, Note: "two edges between the same pair: the cheaper one counts"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1, 1}, []interface{}{1, 2, 1}, []interface{}{2, 3, 1}}, "start": 2}, Expected: []interface{}{-1, -1, 0, 1}, Hidden: true, Note: "start in the middle: nodes behind it are -1"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 1, 1}, []interface{}{0, 2, 5}, []interface{}{1, 2, 1}, []interface{}{2, 3, 1}, []interface{}{3, 1, 1}, []interface{}{3, 4, 2}}, "start": 0}, Expected: []interface{}{0, 1, 2, 3, 5}, Hidden: true, Note: "a cycle and a node first reached by a costly edge (2 is first seen at 5, later improved to 2)"},
		},
		TimeComplexity:  "O((n + m) log n)",
		SpaceComplexity: "O(n + m)",
		StarterCode: `def dijkstra(n, edges, start):
    # Write your solution here
    pass`,
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: `If you always settle the unsettled node with the smallest known cost next, what can you say about the cost you have for it at that moment, given that weights are never negative?`},
			{Level: 2, Type: "insight", Content: `Keep the best known distance to each node and a min-heap of (distance, node). Pop the cheapest entry; that node's distance is now final. Try each outgoing edge: if going through this node gives a cheaper way to a neighbour, record it and push the neighbour. The heap can hold stale entries, so skip a popped entry whose distance is larger than the recorded one. This is Dijkstra's algorithm from Chapter 14.`, ChapterRef: 14},
			{Level: 3, Type: "plan", Content: `1. Build adjacency lists of (neighbour, weight).
2. Set every distance to infinity except ` + "`" + `dist[start] = 0` + "`" + `, and push ` + "`" + `(0, start)` + "`" + ` on a min-heap.
3. Pop the smallest ` + "`" + `(d, node)` + "`" + `; if ` + "`" + `d` + "`" + ` is bigger than ` + "`" + `dist[node]` + "`" + `, skip it.
4. For each edge ` + "`" + `(nxt, w)` + "`" + `, compute ` + "`" + `d + w` + "`" + `; if it beats ` + "`" + `dist[nxt]` + "`" + `, store it and push ` + "`" + `(d + w, nxt)` + "`" + `.
5. When the heap is empty, return the distances, turning infinity into -1.`},
			{Level: 4, Type: "skeleton", Content: `import heapq


def dijkstra(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
    dist = [float("inf")] * n
    dist[start] = 0
    heap = [(0, start)]
    while heap:
        d, node = heapq.heappop(heap)
        if ____:
            continue
        for nxt, w in adj[node]:
            nd = ____
            if nd < dist[nxt]:
                dist[nxt] = nd
                ____
    return [-1 if d == float("inf") else d for d in dist]`},
		},
		Solution: Solution{
			Code: `import heapq


def dijkstra(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
    dist = [float("inf")] * n
    dist[start] = 0
    heap = [(0, start)]
    while heap:
        d, node = heapq.heappop(heap)
        if d > dist[node]:
            continue
        for nxt, w in adj[node]:
            nd = d + w
            if nd < dist[nxt]:
                dist[nxt] = nd
                heapq.heappush(heap, (nd, nxt))
    return [-1 if d == float("inf") else d for d in dist]`,
			Explanation: "Dijkstra is BFS with a min-heap instead of a plain queue. It repeatedly finalizes the unsettled node with the smallest known distance, then offers each of that node's edges as a possible shortcut to its neighbours. Distances only ever decrease, and the heap hands back the cheapest candidate first.",
			KeyInsight:  "Because weights are never negative, the unsettled node with the smallest tentative distance cannot be improved by any route through other unsettled nodes, since those routes already cost at least as much. So when it is popped, its distance is final.",
			WhyItWorks:  "Invariant: every node popped (and not skipped) has its final shortest distance, and each tentative `dist` is the cost of some real path. Take the cheapest unsettled node x. Any path to x passes through a first unsettled node y (the settled ones have already offered their edges); y's tentative distance is at least x's, since x is the cheapest, and the rest of the path costs at least 0, so the path costs at least dist[x]. So dist[x] is final. A stale heap entry has a larger distance than the recorded one, and the check `d > dist[node]` skips it. Each edge relaxes at most once per settled node, so there are at most m pushes, each O(log m): O((n + m) log n) overall.",
			CommonMistakes: []string{
				"Using BFS or a plain queue and counting edges. That finds the fewest edges, not the lowest cost: in the second example it returns 10 for node 3 instead of 6.",
				"Locking in a node's distance the first time it is discovered. A node found through an expensive edge can be reached more cheaply later; keep updating while `nd < dist[nxt]`.",
				"Leaving out the `if d > dist[node]: continue` check. The answer stays right, but stale heap entries are expanded again and the running time can blow up.",
				"Printing `inf` for unreachable nodes. The problem asks for -1, so convert at the end, and do not add weights to -1 along the way.",
			},
			TimeComplexity:  "O((n + m) log n): each edge pushes at most one heap entry, and each heap operation costs a logarithm",
			SpaceComplexity: "O(n + m) for the lists, the distances and the heap",
			Walkthrough: []WalkthroughStep{
				{Title: "Build weighted adjacency lists", Explanation: "Each edge keeps its cost next to the neighbour, as a (neighbour, weight) pair.", CodeSnippet: `adj = [[] for _ in range(n)]
for u, v, w in edges:
    adj[u].append((v, w))`, LineStart: 5, LineEnd: 7},
				{Title: "Start with infinite distances, 0 at the start", Explanation: "Infinity means 'no route found yet'. The start costs 0 and goes in the heap.", CodeSnippet: `dist = [float("inf")] * n
dist[start] = 0
heap = [(0, start)]`, LineStart: 8, LineEnd: 10},
				{Title: "Take the cheapest unsettled node", Explanation: "The heap returns the smallest (distance, node) pair. An entry bigger than the recorded distance is stale, left over from before a cheaper route was found, so skip it.", CodeSnippet: `d, node = heapq.heappop(heap)
if d > dist[node]:
    continue`, LineStart: 12, LineEnd: 14},
				{Title: "Offer each edge as a shortcut", Explanation: "If going through `node` is cheaper than the best known route to the neighbour, record it and push the new candidate. In example 2 node 3 first gets 10, then 6 replaces it.", CodeSnippet: `for nxt, w in adj[node]:
    nd = d + w
    if nd < dist[nxt]:
        dist[nxt] = nd
        heapq.heappush(heap, (nd, nxt))`, LineStart: 15, LineEnd: 19},
				{Title: "Report unreachable nodes as -1", Explanation: "A distance still at infinity was never reached.", CodeSnippet: `return [-1 if d == float("inf") else d for d in dist]`, LineStart: 20, LineEnd: 20},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does Dijkstra's algorithm fail when an edge has a negative weight?", Answer: "It relies on the cheapest unsettled node being final: no later route can beat it, because every extra edge only adds cost. A negative edge breaks that, since a longer route through a not-yet-settled node could end up cheaper than the one you already finalized. Bellman-Ford handles negative weights."},
			{Prompt: "What is the `if d > dist[node]: continue` line for?", Answer: "A node can be pushed on the heap several times, once for each time a cheaper route to it was found. The cheapest entry is popped first and does the real work; the later, more expensive entries are stale. This line recognizes them and skips the repeated work."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, AdvancedGraphsDrills...)
}
