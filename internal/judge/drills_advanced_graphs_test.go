package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestAdvancedGraphsDrillsRejectMistakes checks that each advanced graphs
// foundations drill fails the mistake its tests and notes promise to catch, and
// that the reference solution of each one still passes. The mistakes are the
// ones that look right on a small example: a path built in the wrong direction
// or with the wrong parent, a Prim heap keyed like Dijkstra, a Kruskal with no
// cycle check or an unstable tie order, Bellman-Ford that ignores an extra pass
// or relaxes from an unreachable node, Floyd-Warshall with k in the wrong loop,
// and a greedy walk where Hierholzer's stack is needed.
func TestAdvancedGraphsDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-dijkstra-path", "forgets to reverse the path", `import heapq


def dijkstraPath(n, edges, start, target):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
    dist = [float("inf")] * n
    parent = [-1] * n
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
                parent[nxt] = node
                heapq.heappush(heap, (nd, nxt))
    if dist[target] == float("inf"):
        return []
    path = [target]
    while path[-1] != start:
        path.append(parent[path[-1]])

    return path`},
		{"drill-dijkstra-path", "sets the parent only when a node is first discovered", `import heapq


def dijkstraPath(n, edges, start, target):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
    dist = [float("inf")] * n
    parent = [-1] * n
    dist[start] = 0
    heap = [(0, start)]
    while heap:
        d, node = heapq.heappop(heap)
        if d > dist[node]:
            continue
        for nxt, w in adj[node]:
            nd = d + w
            if nd < dist[nxt]:
                if dist[nxt] == float("inf"):
                    parent[nxt] = node
                dist[nxt] = nd
                heapq.heappush(heap, (nd, nxt))
    if dist[target] == float("inf"):
        return []
    path = [target]
    while path[-1] != start:
        path.append(parent[path[-1]])
    path.reverse()
    return path`},
		{"drill-dijkstra-path", "overwrites the parent for every edge examined", `import heapq


def dijkstraPath(n, edges, start, target):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
    dist = [float("inf")] * n
    parent = [-1] * n
    dist[start] = 0
    heap = [(0, start)]
    while heap:
        d, node = heapq.heappop(heap)
        if d > dist[node]:
            continue
        for nxt, w in adj[node]:
            nd = d + w
            parent[nxt] = node
            if nd < dist[nxt]:
                dist[nxt] = nd
                heapq.heappush(heap, (nd, nxt))
    if dist[target] == float("inf"):
        return []
    path = [target]
    while path[-1] != start:
        path.append(parent[path[-1]])
    path.reverse()
    return path`},
		{"drill-dijkstra-path", "returns the path with the fewest edges (BFS)", `from collections import deque


def dijkstraPath(n, edges, start, target):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append(v)
    parent = {start: -1}
    queue = deque([start])
    while queue:
        node = queue.popleft()
        for nxt in adj[node]:
            if nxt not in parent:
                parent[nxt] = node
                queue.append(nxt)
    if target not in parent:
        return []
    path = [target]
    while path[-1] != start:
        path.append(parent[path[-1]])
    path.reverse()
    return path`},
		{"drill-dijkstra-path", "returns [target] when the target is unreachable", `import heapq


def dijkstraPath(n, edges, start, target):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
    dist = [float("inf")] * n
    parent = [-1] * n
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
                parent[nxt] = node
                heapq.heappush(heap, (nd, nxt))
    path = [target]
    while path[-1] != start and parent[path[-1]] != -1:
        path.append(parent[path[-1]])
    path.reverse()
    return path`},
		{"drill-prim-mst", "keys the heap by distance from the start (Dijkstra style)", `import heapq


def primMST(n, edges):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
        adj[v].append((u, w))
    seen = [False] * n
    heap = [(0, 0, 0)]
    total = 0
    while heap:
        d, w, node = heapq.heappop(heap)
        if seen[node]:
            continue
        seen[node] = True
        total += w
        for nxt, nw in adj[node]:
            if not seen[nxt]:
                heapq.heappush(heap, (d + nw, nw, nxt))
    return total`},
		{"drill-prim-mst", "counts the weight when pushing an edge", `import heapq


def primMST(n, edges):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
        adj[v].append((u, w))
    seen = [False] * n
    heap = [(0, 0)]
    total = 0
    while heap:
        w, node = heapq.heappop(heap)
        if seen[node]:
            continue
        seen[node] = True
        for nxt, nw in adj[node]:
            if not seen[nxt]:
                total += nw
                heapq.heappush(heap, (nw, nxt))
    return total`},
		{"drill-prim-mst", "adds the weight before checking whether the node is already in the tree", `import heapq


def primMST(n, edges):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
        adj[v].append((u, w))
    seen = [False] * n
    heap = [(0, 0)]
    total = 0
    while heap:
        w, node = heapq.heappop(heap)
        total += w
        if seen[node]:
            continue
        seen[node] = True
        for nxt, nw in adj[node]:
            if not seen[nxt]:
                heapq.heappush(heap, (nw, nxt))
    return total`},
		{"drill-prim-mst", "sums each node's cheapest edge", `def primMST(n, edges):
    best = [None] * n
    for u, v, w in edges:
        for x in (u, v):
            if best[x] is None or w < best[x]:
                best[x] = w
    return sum(b for b in best if b is not None)`},
		{"drill-prim-mst", "rescans every edge for each tree edge (O(n * m))", `def primMST(n, edges):
    inside = [False] * n
    inside[0] = True
    total = 0
    for _ in range(n - 1):
        best = None
        pick = -1
        for u, v, w in edges:
            if inside[u] != inside[v] and (best is None or w < best):
                best = w
                pick = v if inside[u] else u
        inside[pick] = True
        total += best
    return total`},
		{"drill-kruskal-mst", "has no cycle check", `def kruskal(n, edges):
    return [[u, v] for u, v, w in sorted(edges, key=lambda e: e[2])]`},
		{"drill-kruskal-mst", "breaks ties by node numbers instead of input order", `def kruskal(n, edges):
    parent = list(range(n))

    def find(x):
        while parent[x] != x:
            parent[x] = parent[parent[x]]
            x = parent[x]
        return x

    chosen = []
    for u, v, w in sorted(edges, key=lambda e: (e[2], e[0], e[1])):
        ru, rv = find(u), find(v)
        if ru != rv:
            parent[ru] = rv
            chosen.append([u, v])
    return chosen`},
		{"drill-kruskal-mst", "compares parent entries instead of roots", `def kruskal(n, edges):
    parent = list(range(n))
    chosen = []
    for u, v, w in sorted(edges, key=lambda e: e[2]):
        if parent[u] != parent[v]:
            parent[v] = parent[u]
            chosen.append([u, v])
    return chosen`},
		{"drill-kruskal-mst", "returns the accepted edges in input order", `def kruskal(n, edges):
    parent = list(range(n))

    def find(x):
        while parent[x] != x:
            parent[x] = parent[parent[x]]
            x = parent[x]
        return x

    order = sorted(range(len(edges)), key=lambda i: edges[i][2])
    taken = []
    for i in order:
        u, v, w = edges[i]
        ru, rv = find(u), find(v)
        if ru != rv:
            parent[ru] = rv
            taken.append(i)
    return [[edges[i][0], edges[i][1]] for i in sorted(taken)]`},
		{"drill-kruskal-mst", "returns the endpoints as [smaller, larger]", `def kruskal(n, edges):
    parent = list(range(n))

    def find(x):
        while parent[x] != x:
            parent[x] = parent[parent[x]]
            x = parent[x]
        return x

    chosen = []
    for u, v, w in sorted(edges, key=lambda e: e[2]):
        ru, rv = find(u), find(v)
        if ru != rv:
            parent[ru] = rv
            chosen.append([min(u, v), max(u, v)])
    return chosen`},
		{"drill-kruskal-mst", "assumes the graph is connected and stops at n - 1 edges only by looping a spanning tree (fails on a forest)", `def kruskal(n, edges):
    parent = list(range(n))

    def find(x):
        while parent[x] != x:
            parent[x] = parent[parent[x]]
            x = parent[x]
        return x

    chosen = []
    for u, v, w in sorted(edges, key=lambda e: e[2]):
        ru, rv = find(u), find(v)
        if ru != rv:
            parent[ru] = rv
            chosen.append([u, v])
    if len(chosen) < n - 1:
        return []
    return chosen`},
		{"drill-bellman-ford", "never checks for a negative cycle", `def bellmanFord(n, edges, start):
    INF = float("inf")
    dist = [INF] * n
    dist[start] = 0
    for _ in range(n - 1):
        for u, v, w in edges:
            if dist[u] != INF and dist[u] + w < dist[v]:
                dist[v] = dist[u] + w
    return ["inf" if d == INF else d for d in dist]`},
		{"drill-bellman-ford", "uses a big number for infinity and relaxes from unreachable nodes", `def bellmanFord(n, edges, start):
    INF = 10 ** 9
    dist = [INF] * n
    dist[start] = 0
    for _ in range(n - 1):
        for u, v, w in edges:
            if dist[u] + w < dist[v]:
                dist[v] = dist[u] + w
    for u, v, w in edges:
        if dist[u] + w < dist[v]:
            return []
    return ["inf" if d >= INF else d for d in dist]`},
		{"drill-bellman-ford", "stops after n - 2 rounds", `def bellmanFord(n, edges, start):
    INF = float("inf")
    dist = [INF] * n
    dist[start] = 0
    for _ in range(max(n - 2, 1)):
        for u, v, w in edges:
            if dist[u] != INF and dist[u] + w < dist[v]:
                dist[v] = dist[u] + w
    for u, v, w in edges:
        if dist[u] != INF and dist[u] + w < dist[v]:
            return []
    return ["inf" if d == INF else d for d in dist]`},
		{"drill-bellman-ford", "reports a cycle of weight 0 as negative (<= in the check)", `def bellmanFord(n, edges, start):
    INF = float("inf")
    dist = [INF] * n
    dist[start] = 0
    for _ in range(n - 1):
        for u, v, w in edges:
            if dist[u] != INF and dist[u] + w < dist[v]:
                dist[v] = dist[u] + w
    for u, v, w in edges:
        if dist[u] != INF and dist[u] + w <= dist[v]:
            return []
    return ["inf" if d == INF else d for d in dist]`},
		{"drill-bellman-ford", "uses Dijkstra's settled-once rule", `import heapq


def bellmanFord(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
    INF = float("inf")
    dist = [INF] * n
    dist[start] = 0
    done = [False] * n
    heap = [(0, start)]
    while heap:
        d, node = heapq.heappop(heap)
        if done[node]:
            continue
        done[node] = True
        for nxt, w in adj[node]:
            if not done[nxt] and d + w < dist[nxt]:
                dist[nxt] = d + w
                heapq.heappush(heap, (dist[nxt], nxt))
    return ["inf" if d == INF else d for d in dist]`},
		{"drill-bellman-ford", "treats a negative cycle that start cannot reach as a failure", `def bellmanFord(n, edges, start):
    INF = float("inf")
    dist = [0] * n
    dist[start] = 0
    for _ in range(n - 1):
        for u, v, w in edges:
            if dist[u] + w < dist[v]:
                dist[v] = dist[u] + w
    for u, v, w in edges:
        if dist[u] + w < dist[v]:
            return []
    return dist`},
		{"drill-floyd-warshall", "puts the k loop innermost", `def floydWarshall(n, edges):
    INF = float("inf")
    dist = [[INF] * n for _ in range(n)]
    for i in range(n):
        dist[i][i] = 0
    for u, v, w in edges:
        if w < dist[u][v]:
            dist[u][v] = w
    for i in range(n):
        for j in range(n):
            for k in range(n):
                if dist[i][k] + dist[k][j] < dist[i][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]
    return [[-1 if d == INF else d for d in row] for row in dist]`},
		{"drill-floyd-warshall", "puts the k loop in the middle", `def floydWarshall(n, edges):
    INF = float("inf")
    dist = [[INF] * n for _ in range(n)]
    for i in range(n):
        dist[i][i] = 0
    for u, v, w in edges:
        if w < dist[u][v]:
            dist[u][v] = w
    for i in range(n):
        for k in range(n):
            for j in range(n):
                if dist[i][k] + dist[k][j] < dist[i][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]
    return [[-1 if d == INF else d for d in row] for row in dist]`},
		{"drill-floyd-warshall", "keeps the last parallel edge instead of the cheapest", `def floydWarshall(n, edges):
    INF = float("inf")
    dist = [[INF] * n for _ in range(n)]
    for i in range(n):
        dist[i][i] = 0
    for u, v, w in edges:
        dist[u][v] = w
    for k in range(n):
        for i in range(n):
            for j in range(n):
                if dist[i][k] + dist[k][j] < dist[i][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]
    return [[-1 if d == INF else d for d in row] for row in dist]`},
		{"drill-floyd-warshall", "lets a self-loop overwrite the diagonal", `def floydWarshall(n, edges):
    INF = float("inf")
    dist = [[INF] * n for _ in range(n)]
    for i in range(n):
        dist[i][i] = 0
    for u, v, w in edges:
        if u == v or w < dist[u][v]:
            dist[u][v] = w
    for k in range(n):
        for i in range(n):
            for j in range(n):
                if dist[i][k] + dist[k][j] < dist[i][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]
    return [[-1 if d == INF else d for d in row] for row in dist]`},
		{"drill-floyd-warshall", "uses -1 as the unreachable value inside the algorithm", `def floydWarshall(n, edges):
    dist = [[-1] * n for _ in range(n)]
    for i in range(n):
        dist[i][i] = 0
    for u, v, w in edges:
        if dist[u][v] == -1 or w < dist[u][v]:
            dist[u][v] = w
    for k in range(n):
        for i in range(n):
            for j in range(n):
                if min(dist[i][k], dist[k][j]) < dist[i][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]
    return dist`},
		{"drill-eulerian-path", "forgets to reverse the post-order", `def eulerianPath(edges):
    adj = {}
    balance = {}
    for u, v in edges:
        adj.setdefault(u, []).append(v)
        balance[u] = balance.get(u, 0) + 1
        balance[v] = balance.get(v, 0) - 1
    start = min(adj)
    for node, b in balance.items():
        if b == 1:
            start = node
    stack = [start]
    path = []
    while stack:
        node = stack[-1]
        if adj.get(node):
            stack.append(adj[node].pop())
        else:
            path.append(stack.pop())
    return path`},
		{"drill-eulerian-path", "always starts at the smallest node", `def eulerianPath(edges):
    adj = {}
    balance = {}
    for u, v in edges:
        adj.setdefault(u, []).append(v)
        balance[u] = balance.get(u, 0) + 1
        balance[v] = balance.get(v, 0) - 1
    start = min(adj)
    stack = [start]
    path = []
    while stack:
        node = stack[-1]
        if adj.get(node):
            stack.append(adj[node].pop())
        else:
            path.append(stack.pop())
    path.reverse()
    return path`},
		{"drill-eulerian-path", "walks greedily and never backs up (takes the last edge)", `def eulerianPath(edges):
    adj = {}
    balance = {}
    for u, v in edges:
        adj.setdefault(u, []).append(v)
        balance[u] = balance.get(u, 0) + 1
        balance[v] = balance.get(v, 0) - 1
    start = min(adj)
    for node, b in balance.items():
        if b == 1:
            start = node
    path = [start]
    while adj.get(path[-1]):
        path.append(adj[path[-1]].pop())
    return path`},
		{"drill-eulerian-path", "walks greedily and never backs up (takes the first edge)", `def eulerianPath(edges):
    adj = {}
    balance = {}
    for u, v in edges:
        adj.setdefault(u, []).append(v)
        balance[u] = balance.get(u, 0) + 1
        balance[v] = balance.get(v, 0) - 1
    start = min(adj)
    for node, b in balance.items():
        if b == 1:
            start = node
    path = [start]
    while adj.get(path[-1]):
        path.append(adj[path[-1]].pop(0))
    return path`},
		{"drill-eulerian-path", "marks nodes visited instead of using edges up", `def eulerianPath(edges):
    adj = {}
    balance = {}
    for u, v in edges:
        adj.setdefault(u, []).append(v)
        balance[u] = balance.get(u, 0) + 1
        balance[v] = balance.get(v, 0) - 1
    start = min(adj)
    for node, b in balance.items():
        if b == 1:
            start = node
    seen = set()
    path = []

    def dfs(node):
        seen.add(node)
        for nxt in adj.get(node, []):
            if nxt not in seen:
                dfs(nxt)
        path.append(node)

    dfs(start)
    path.reverse()
    return path`},
		{"drill-eulerian-path", "ignores parallel edges (uses each neighbour once)", `def eulerianPath(edges):
    adj = {}
    balance = {}
    for u, v in edges:
        adj.setdefault(u, []).append(v)
        balance[u] = balance.get(u, 0) + 1
        balance[v] = balance.get(v, 0) - 1
    start = min(adj)
    for node, b in balance.items():
        if b == 1:
            start = node
    for node in adj:
        adj[node] = sorted(set(adj[node]), reverse=True)
    stack = [start]
    path = []
    while stack:
        node = stack[-1]
        if adj.get(node):
            stack.append(adj[node].pop())
        else:
            path.append(stack.pop())
    path.reverse()
    return path`},
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
