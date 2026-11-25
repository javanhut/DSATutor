package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestGraphDrillsRejectMistakes checks that each graphs foundations drill
// fails the main mistake its tests and notes promise to catch, so a drill
// can't be passed without doing the thing it drills.
func TestGraphDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-build-adjacency-list", "forgets the reverse edge when undirected", `def buildAdjacencyList(n, edges, directed):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
    return adj`},
		{"drill-build-adjacency-list", "shares one list between every row", `def buildAdjacencyList(n, edges, directed):
    adj = [[]] * n
    for u, v in edges:
        adj[u].append(v)
        if not directed:
            adj[v].append(u)
    return adj`},
		{"drill-build-adjacency-matrix", "forgets the mirrored cell when undirected", `def buildAdjacencyMatrix(n, edges, directed):
    matrix = [[0] * n for _ in range(n)]
    for u, v in edges:
        matrix[u][v] = 1
    return matrix`},
		{"drill-build-adjacency-matrix", "always mirrors, even when directed", `def buildAdjacencyMatrix(n, edges, directed):
    matrix = [[0] * n for _ in range(n)]
    for u, v in edges:
        matrix[u][v] = 1
        matrix[v][u] = 1
    return matrix`},
		{"drill-bfs-order", "uses a stack instead of a queue", `def bfsOrder(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for x in adj:
        x.sort()
    order, visited, stack = [], [False] * n, [start]
    visited[start] = True
    while stack:
        node = stack.pop()
        order.append(node)
        for nxt in adj[node]:
            if not visited[nxt]:
                visited[nxt] = True
                stack.append(nxt)
    return order`},
		{"drill-bfs-order", "does not sort the neighbours", `from collections import deque

def bfsOrder(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    order, visited, queue = [], [False] * n, deque([start])
    visited[start] = True
    while queue:
        node = queue.popleft()
        order.append(node)
        for nxt in adj[node]:
            if not visited[nxt]:
                visited[nxt] = True
                queue.append(nxt)
    return order`},
		{"drill-dfs-order", "has no visited set, so a cycle never ends", `def dfsOrder(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for x in adj:
        x.sort()
    order = []
    def visit(node, parent):
        order.append(node)
        for nxt in adj[node]:
            if nxt != parent:
                visit(nxt, node)
    visit(start, -1)
    return order`},
		{"drill-dfs-order", "walks breadth-first", `from collections import deque

def dfsOrder(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for x in adj:
        x.sort()
    order, visited, queue = [], [False] * n, deque([start])
    visited[start] = True
    while queue:
        node = queue.popleft()
        order.append(node)
        for nxt in adj[node]:
            if not visited[nxt]:
                visited[nxt] = True
                queue.append(nxt)
    return order`},
		{"drill-dfs-iterative", "recurses, so a 25,000-node path overflows", `def dfsOrderIterative(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for x in adj:
        x.sort()
    order, visited = [], [False] * n
    def visit(node):
        visited[node] = True
        order.append(node)
        for nxt in adj[node]:
            if not visited[nxt]:
                visit(nxt)
    visit(start)
    return order`},
		{"drill-dfs-iterative", "pushes neighbours in increasing order", `def dfsOrderIterative(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for x in adj:
        x.sort()
    order, visited, stack = [], [False] * n, [start]
    while stack:
        node = stack.pop()
        if visited[node]:
            continue
        visited[node] = True
        order.append(node)
        for nxt in adj[node]:
            if not visited[nxt]:
                stack.append(nxt)
    return order`},
		{"drill-shortest-path-unweighted", "returns the length of the first path a depth-first search finds", `def shortestPath(n, edges, start, target):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for x in adj:
        x.sort()
    seen = [False] * n
    def go(node, d):
        if node == target:
            return d
        seen[node] = True
        for nxt in adj[node]:
            if not seen[nxt]:
                r = go(nxt, d + 1)
                if r != -1:
                    return r
        return -1
    return go(start, 0)`},
		{"drill-flood-fill", "has no guard for a new colour equal to the old one", `def floodFill(image, sr, sc, color):
    old = image[sr][sc]
    rows, cols = len(image), len(image[0])
    def fill(r, c):
        if r < 0 or r >= rows or c < 0 or c >= cols or image[r][c] != old:
            return
        image[r][c] = color
        fill(r + 1, c)
        fill(r - 1, c)
        fill(r, c + 1)
        fill(r, c - 1)
    fill(sr, sc)
    return image`},
		{"drill-flood-fill", "spreads to diagonal neighbours", `def floodFill(image, sr, sc, color):
    old = image[sr][sc]
    if old == color:
        return image
    rows, cols = len(image), len(image[0])
    def fill(r, c):
        if r < 0 or r >= rows or c < 0 or c >= cols or image[r][c] != old:
            return
        image[r][c] = color
        for dr in (-1, 0, 1):
            for dc in (-1, 0, 1):
                fill(r + dr, c + dc)
    fill(sr, sc)
    return image`},
		{"drill-topological-sort", "returns a partial order for a graph with a cycle", `from collections import deque

def topologicalSort(n, edges):
    adj = [[] for _ in range(n)]
    indegree = [0] * n
    for u, v in edges:
        adj[u].append(v)
        indegree[v] += 1
    queue = deque(i for i in range(n) if indegree[i] == 0)
    order = []
    while queue:
        node = queue.popleft()
        order.append(node)
        for nxt in adj[node]:
            indegree[nxt] -= 1
            if indegree[nxt] == 0:
                queue.append(nxt)
    return order`},
		{"drill-topological-sort", "counts the in-degree at the wrong end of each edge", `from collections import deque

def topologicalSort(n, edges):
    adj = [[] for _ in range(n)]
    indegree = [0] * n
    for u, v in edges:
        adj[v].append(u)
        indegree[u] += 1
    queue = deque(i for i in range(n) if indegree[i] == 0)
    order = []
    while queue:
        node = queue.popleft()
        order.append(node)
        for nxt in adj[node]:
            indegree[nxt] -= 1
            if indegree[nxt] == 0:
                queue.append(nxt)
    return order if len(order) == n else []`},
		{"drill-union-find", "union returns True even when already connected", `class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.sets = n
    def find(self, x):
        while self.parent[x] != x:
            x = self.parent[x]
        return x
    def union(self, a, b):
        ra, rb = self.find(a), self.find(b)
        if ra != rb:
            self.parent[ra] = rb
            self.sets -= 1
        return True
    def connected(self, a, b):
        return self.find(a) == self.find(b)
    def count(self):
        return self.sets`},
		{"drill-union-find", "links the elements instead of their roots", `class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.sets = n
    def find(self, x):
        while self.parent[x] != x:
            x = self.parent[x]
        return x
    def union(self, a, b):
        if self.find(a) == self.find(b):
            return False
        self.parent[a] = b
        self.sets -= 1
        return True
    def connected(self, a, b):
        return self.find(a) == self.find(b)
    def count(self):
        return self.sets`},
		{"drill-union-find", "find recurses down a 25,000-deep chain", `class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.sets = n
    def find(self, x):
        if self.parent[x] != x:
            return self.find(self.parent[x])
        return x
    def union(self, a, b):
        ra, rb = self.find(a), self.find(b)
        if ra == rb:
            return False
        self.parent[ra] = rb
        self.sets -= 1
        return True
    def connected(self, a, b):
        return self.find(a) == self.find(b)
    def count(self):
        return self.sets`},
		{"drill-dijkstra", "counts edges instead of weights (BFS)", `from collections import deque

def dijkstra(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
    dist = [-1] * n
    dist[start] = 0
    queue = deque([start])
    while queue:
        node = queue.popleft()
        for nxt, w in adj[node]:
            if dist[nxt] == -1:
                dist[nxt] = dist[node] + w
                queue.append(nxt)
    return dist`},
		{"drill-dijkstra", "never improves a distance once it is set", `import heapq

def dijkstra(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v, w in edges:
        adj[u].append((v, w))
    dist = [-1] * n
    dist[start] = 0
    heap = [(0, start)]
    while heap:
        d, node = heapq.heappop(heap)
        for nxt, w in adj[node]:
            if dist[nxt] == -1:
                dist[nxt] = d + w
                heapq.heappush(heap, (d + w, nxt))
    return dist`},
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
