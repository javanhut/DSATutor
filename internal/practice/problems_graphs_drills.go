package practice

// GraphsDrills are the graphs foundations: build the two standard representations, run each core traversal
// and the three classic algorithms on their own, before the graphs ladder applies them. They follow
// docs/pedagogy.md like any problem, minus the brute force (a drill is the plain algorithm itself).
var GraphsDrills = []*Problem{
	{
		ID:              "drill-build-adjacency-list",
		Kind:            KindDrill,
		Order:           1,
		Title:           "Build an Adjacency List",
		Slug:            "drill-build-adjacency-list",
		Difficulty:      "Easy",
		Category:        "graphs",
		Tags:            []string{"Graph", "Adjacency List", "Foundations"},
		RelatedChapters: []int{13},
		Description: `Almost every graph problem starts by turning a list of edges into something you can walk. The most common shape is the **adjacency list**: for each node, the list of nodes it has an edge to.

Write ` + "`" + `buildAdjacencyList(n, edges, directed)` + "`" + `. There are ` + "`" + `n` + "`" + ` nodes labelled ` + "`" + `0` + "`" + ` to ` + "`" + `n-1` + "`" + `, and each entry of ` + "`" + `edges` + "`" + ` is a pair ` + "`" + `[u, v]` + "`" + `. Return a list of ` + "`" + `n` + "`" + ` lists, where entry ` + "`" + `i` + "`" + ` holds the neighbours of node ` + "`" + `i` + "`" + `.

- If ` + "`" + `directed` + "`" + ` is true, the edge ` + "`" + `[u, v]` + "`" + ` goes from ` + "`" + `u` + "`" + ` to ` + "`" + `v` + "`" + ` only, so ` + "`" + `v` + "`" + ` is added to ` + "`" + `u` + "`" + `'s list and nothing else.
- If ` + "`" + `directed` + "`" + ` is false, the edge works both ways: ` + "`" + `v` + "`" + ` is added to ` + "`" + `u` + "`" + `'s list and then ` + "`" + `u` + "`" + ` is added to ` + "`" + `v` + "`" + `'s list.

Neighbours appear in the order their edges appear in ` + "`" + `edges` + "`" + `. Do not sort them.`,
		Constraints: []string{
			"1 <= n <= 1000",
			"0 <= len(edges) <= 5000",
			"0 <= u, v < n",
			"No self-loops and no repeated edges",
		},
		Examples: []Example{
			{Input: "n = 4, edges = [[0,1],[0,2],[1,2],[2,3]], directed = false", Output: "[[1,2],[0,2],[0,1,3],[2]]", Explanation: "Each edge is stored at both ends. Node 2 sees 0, then 1, then 3, because [0,2], [1,2], [2,3] appear in that order."},
			{Input: "n = 4, edges = [[0,1],[0,2],[1,2],[2,3]], directed = true", Output: "[[1,2],[2],[3],[]]", Explanation: "Directed: each edge is stored only at its source. Node 3 has no outgoing edges, so its list is empty."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 2}, []interface{}{2, 3}}, "directed": false}, Expected: []interface{}{[]interface{}{1, 2}, []interface{}{0, 2}, []interface{}{0, 1, 3}, []interface{}{2}}, Note: "example 1: an undirected edge is stored at both ends"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 2}, []interface{}{2, 3}}, "directed": true}, Expected: []interface{}{[]interface{}{1, 2}, []interface{}{2}, []interface{}{3}, []interface{}{}}, Note: "example 2: the same edges, directed: each edge is stored once, at its source"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{3, 4}}, "directed": false}, Expected: []interface{}{[]interface{}{1}, []interface{}{0}, []interface{}{}, []interface{}{4}, []interface{}{3}}, Hidden: true, Note: "node 2 has no edges: its row is an empty list"},
			{Input: map[string]interface{}{"n": 1, "edges": []interface{}{}, "directed": false}, Expected: []interface{}{[]interface{}{}}, Hidden: true, Note: "one node, no edges: [[]]"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{2, 0}, []interface{}{1, 0}, []interface{}{0, 3}}, "directed": false}, Expected: []interface{}{[]interface{}{2, 1, 3}, []interface{}{0}, []interface{}{0}, []interface{}{0}}, Hidden: true, Note: "neighbours come in the order the edges appear, not sorted (0 gets 2, then 1, then 3)"},
			{Input: map[string]interface{}{"n": 2, "edges": []interface{}{[]interface{}{1, 0}}, "directed": true}, Expected: []interface{}{[]interface{}{}, []interface{}{0}}, Hidden: true, Note: "directed edge 1 -> 0: only row 1 has an entry (adding the reverse edge always fails)"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{0, 3}}, "directed": true}, Expected: []interface{}{[]interface{}{1, 2, 3}, []interface{}{}, []interface{}{}, []interface{}{}}, Hidden: true, Note: "a directed star: the three leaf rows must be separate empty lists (`[[]] * n` shares one list)"},
			{Input: map[string]interface{}{"n": 3, "edges": []interface{}{}, "directed": false}, Expected: []interface{}{[]interface{}{}, []interface{}{}, []interface{}{}}, Hidden: true, Note: "three nodes, no edges: [[], [], []]"},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(n + m)",
		StarterCode: `def buildAdjacencyList(n, edges, directed):
    # Write your solution here
    pass`,
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: `You need one list of neighbours per node, and then one step per edge. What should every node start with before any edge is read?`},
			{Level: 2, Type: "insight", Content: `Make ` + "`" + `n` + "`" + ` separate empty lists first, so a node with no edges still has a row. Then loop over the edges once: an edge ` + "`" + `[u, v]` + "`" + ` always puts ` + "`" + `v` + "`" + ` in ` + "`" + `u` + "`" + `'s list, and only an undirected edge also puts ` + "`" + `u` + "`" + ` in ` + "`" + `v` + "`" + `'s list. This is the adjacency list from Chapter 13.`, ChapterRef: 13},
			{Level: 3, Type: "plan", Content: `1. Create ` + "`" + `n` + "`" + ` empty lists, one per node (not one list repeated ` + "`" + `n` + "`" + ` times).
2. Loop over each edge ` + "`" + `[u, v]` + "`" + `.
3. Append ` + "`" + `v` + "`" + ` to the list of ` + "`" + `u` + "`" + `.
4. If the graph is undirected, also append ` + "`" + `u` + "`" + ` to the list of ` + "`" + `v` + "`" + `.
5. Return the lists.`},
			{Level: 4, Type: "skeleton", Content: `def buildAdjacencyList(n, edges, directed):
    adj = ____
    for u, v in edges:
        adj[u].append(____)
        if ____:
            adj[v].append(____)
    return adj`},
		},
		Solution: Solution{
			Code: `def buildAdjacencyList(n, edges, directed):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        if not directed:
            adj[v].append(u)
    return adj`,
			Explanation: "An adjacency list stores, for each node, only the neighbours it actually has. Build `n` empty lists, then read each edge once and file it under the node it leaves from, and, for an undirected edge, also under the node it arrives at.",
			KeyInsight:  "An undirected edge is two directed edges, one each way. So the only difference between the two modes is whether the reverse entry is added, and forgetting it is the classic bug.",
			WhyItWorks:  "Invariant: after the loop has handled the first k edges, `adj[x]` holds exactly the neighbours of x that those k edges give it, in the order the edges came. Each edge adds one entry (directed) or two (undirected) and touches nothing else, so the invariant holds after every edge. The lists start separate and empty, so rows never share data. One pass over m edges plus n list creations: O(n + m) time, and the lists hold at most 2m entries.",
			CommonMistakes: []string{
				"Writing `adj = [[]] * n`. That makes n references to one list, so every append shows up in every row. Use a list comprehension that creates a fresh list per node.",
				"Forgetting `adj[v].append(u)` for an undirected graph. Then a traversal can walk from u to v but never back, and most answers go wrong quietly.",
				"Always adding the reverse edge, even for a directed graph. A directed edge 1 -> 0 must not appear in node 0's list.",
				"Sorting each list. The drill asks for edge order, and later problems rely on knowing exactly what order you built.",
			},
			TimeComplexity:  "O(n + m): one pass over the edges, plus one empty list per node",
			SpaceComplexity: "O(n + m) for the lists",
			Walkthrough: []WalkthroughStep{
				{Title: "Make one empty list per node", Explanation: "A comprehension runs the `[]` expression n times, so each node gets its own list, and a node with no edges ends up with a visible empty row.", CodeSnippet: `adj = [[] for _ in range(n)]`, LineStart: 2, LineEnd: 2},
				{Title: "Store the edge at its source", Explanation: "For [2,0] this puts 0 into node 2's list. Doing it in edge order is what fixes the neighbour order.", CodeSnippet: `for u, v in edges:
    adj[u].append(v)`, LineStart: 3, LineEnd: 4},
				{Title: "Store the reverse edge when undirected", Explanation: "An undirected edge also lets you travel from v back to u, so u joins v's list. A directed edge skips this.", CodeSnippet: `if not directed:
    adj[v].append(u)`, LineStart: 5, LineEnd: 6},
				{Title: "Return the lists", Explanation: "Entry i is the neighbour list of node i.", CodeSnippet: `return adj`, LineStart: 7, LineEnd: 7},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is an adjacency list usually better than an n x n matrix for a graph with few edges?", Answer: "The list stores only the edges that exist, O(n + m) in total, while the matrix always stores n * n cells even when almost all are 0. Walking a node's neighbours is also faster: the list gives exactly them, while a matrix row makes you scan all n cells. The matrix wins only when you need an O(1) 'is there an edge from u to v' test or the graph is dense."},
			{Prompt: "In an undirected graph with m edges, how many entries do the lists hold in total, and why?", Answer: "2m. Each undirected edge is stored once at each end. For a directed graph it is m, one entry per edge."},
		},
	},
	{
		ID:              "drill-build-adjacency-matrix",
		Kind:            KindDrill,
		Order:           2,
		Title:           "Build an Adjacency Matrix",
		Slug:            "drill-build-adjacency-matrix",
		Difficulty:      "Easy",
		Category:        "graphs",
		Tags:            []string{"Graph", "Adjacency Matrix", "Matrix", "Foundations"},
		RelatedChapters: []int{13},
		Description: `The other standard way to store a graph is an **adjacency matrix**: an n x n grid where cell ` + "`" + `[u][v]` + "`" + ` is ` + "`" + `1` + "`" + ` if there is an edge from ` + "`" + `u` + "`" + ` to ` + "`" + `v` + "`" + `, and ` + "`" + `0` + "`" + ` if not.

Write ` + "`" + `buildAdjacencyMatrix(n, edges, directed)` + "`" + `. There are ` + "`" + `n` + "`" + ` nodes labelled ` + "`" + `0` + "`" + ` to ` + "`" + `n-1` + "`" + `, and each entry of ` + "`" + `edges` + "`" + ` is a pair ` + "`" + `[u, v]` + "`" + `. Return the n x n matrix of 0s and 1s.

- If ` + "`" + `directed` + "`" + ` is true, the edge ` + "`" + `[u, v]` + "`" + ` sets only cell ` + "`" + `[u][v]` + "`" + `.
- If ` + "`" + `directed` + "`" + ` is false, the edge works both ways, so it sets both ` + "`" + `[u][v]` + "`" + ` and ` + "`" + `[v][u]` + "`" + `. An undirected graph's matrix is therefore symmetric.`,
		Constraints: []string{
			"1 <= n <= 200",
			"0 <= len(edges) <= 2000",
			"0 <= u, v < n",
			"No self-loops and no repeated edges",
		},
		Examples: []Example{
			{Input: "n = 4, edges = [[0,1],[0,2],[1,2],[2,3]], directed = false", Output: "[[0,1,1,0],[1,0,1,0],[1,1,0,1],[0,0,1,0]]", Explanation: "Every edge sets two mirrored cells, so row 2 reads 1,1,0,1: node 2 touches 0, 1 and 3."},
			{Input: "n = 3, edges = [[0,1]], directed = true", Output: "[[0,1,0],[0,0,0],[0,0,0]]", Explanation: "Directed: only cell [0][1] is set. Cell [1][0] stays 0 because there is no edge from 1 to 0."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 2}, []interface{}{2, 3}}, "directed": false}, Expected: []interface{}{[]interface{}{0, 1, 1, 0}, []interface{}{1, 0, 1, 0}, []interface{}{1, 1, 0, 1}, []interface{}{0, 0, 1, 0}}, Note: "example 1: an undirected edge sets both mirrored cells, so the matrix is symmetric"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 2}, []interface{}{2, 3}}, "directed": true}, Expected: []interface{}{[]interface{}{0, 1, 1, 0}, []interface{}{0, 0, 1, 0}, []interface{}{0, 0, 0, 1}, []interface{}{0, 0, 0, 0}}, Note: "example 2: the same edges, directed: only cell [u][v] is set"},
			{Input: map[string]interface{}{"n": 3, "edges": []interface{}{[]interface{}{0, 1}}, "directed": true}, Expected: []interface{}{[]interface{}{0, 1, 0}, []interface{}{0, 0, 0}, []interface{}{0, 0, 0}}, Hidden: true, Note: "directed edge 0 -> 1 leaves cell [1][0] at 0 (adding the reverse always fails)"},
			{Input: map[string]interface{}{"n": 1, "edges": []interface{}{}, "directed": false}, Expected: []interface{}{[]interface{}{0}}, Hidden: true, Note: "one node, no edges: [[0]]"},
			{Input: map[string]interface{}{"n": 3, "edges": []interface{}{}, "directed": false}, Expected: []interface{}{[]interface{}{0, 0, 0}, []interface{}{0, 0, 0}, []interface{}{0, 0, 0}}, Hidden: true, Note: "no edges: all zeros"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 4}, []interface{}{4, 2}, []interface{}{1, 3}}, "directed": false}, Expected: []interface{}{[]interface{}{0, 0, 0, 0, 1}, []interface{}{0, 0, 0, 1, 0}, []interface{}{0, 0, 0, 0, 1}, []interface{}{0, 1, 0, 0, 0}, []interface{}{1, 0, 1, 0, 0}}, Hidden: true, Note: "labels far apart, with node 4 on two edges"},
			{Input: map[string]interface{}{"n": 3, "edges": []interface{}{[]interface{}{2, 0}, []interface{}{1, 2}}, "directed": true}, Expected: []interface{}{[]interface{}{0, 0, 0}, []interface{}{0, 0, 1}, []interface{}{1, 0, 0}}, Hidden: true, Note: "directed edges listed from higher to lower labels"},
			{Input: map[string]interface{}{"n": 3, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{0, 2}}, "directed": false}, Expected: []interface{}{[]interface{}{0, 1, 1}, []interface{}{1, 0, 1}, []interface{}{1, 1, 0}}, Hidden: true, Note: "a triangle fills every off-diagonal cell and leaves the diagonal at 0"},
		},
		TimeComplexity:  "O(n^2 + m)",
		SpaceComplexity: "O(n^2)",
		StarterCode: `def buildAdjacencyMatrix(n, edges, directed):
    # Write your solution here
    pass`,
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: `A matrix has a cell for every possible pair of nodes. What should every cell hold before any edge is read, and which cells does one edge change?`},
			{Level: 2, Type: "insight", Content: `Start with an n x n grid of zeros, building each row separately so they do not share memory. An edge ` + "`" + `[u, v]` + "`" + ` sets cell ` + "`" + `[u][v]` + "`" + ` to 1; an undirected edge also sets ` + "`" + `[v][u]` + "`" + `. Compare with the adjacency list from the last drill and Chapter 13.`, ChapterRef: 13},
			{Level: 3, Type: "plan", Content: `1. Create an n x n grid of 0s, one fresh row per node.
2. Loop over each edge ` + "`" + `[u, v]` + "`" + `.
3. Set ` + "`" + `matrix[u][v] = 1` + "`" + `.
4. If undirected, also set ` + "`" + `matrix[v][u] = 1` + "`" + `.
5. Return the matrix.`},
			{Level: 4, Type: "skeleton", Content: `def buildAdjacencyMatrix(n, edges, directed):
    matrix = [____ for _ in range(n)]
    for u, v in edges:
        matrix[u][v] = ____
        if ____:
            matrix[v][u] = 1
    return matrix`},
		},
		Solution: Solution{
			Code: `def buildAdjacencyMatrix(n, edges, directed):
    matrix = [[0] * n for _ in range(n)]
    for u, v in edges:
        matrix[u][v] = 1
        if not directed:
            matrix[v][u] = 1
    return matrix`,
			Explanation: "Make an n x n grid of zeros, then for each edge write a 1 into the cell for that pair, and into the mirrored cell when the edge works both ways. Asking whether two nodes are joined is then a single cell lookup.",
			KeyInsight:  "Cell [u][v] means 'there is an edge from u to v', so an undirected edge is simply two such facts, one in each direction. That is why undirected matrices are symmetric and directed ones need not be.",
			WhyItWorks:  "Invariant: after the loop has handled the first k edges, cell [a][b] is 1 exactly when one of those edges is [a, b] (or [b, a] when undirected), and 0 otherwise. Each edge writes only the one or two cells it describes, so the invariant holds after each step. Each row is built by its own `[0] * n`, so writing one row never changes another. Building the grid costs O(n^2) and each edge O(1), so O(n^2 + m) time and O(n^2) space.",
			CommonMistakes: []string{
				"Building the grid as `[[0] * n] * n`. That repeats one row n times, so setting a cell changes the same column in every row. Build each row with a comprehension.",
				"Forgetting the mirrored cell for an undirected edge, which leaves the matrix lopsided: node 0 can reach 1, but 1 cannot reach 0.",
				"Always setting the mirrored cell, even when directed. A directed edge 0 -> 1 must leave [1][0] at 0.",
				"Swapping the cell's meaning halfway, treating [u][v] as 'v to u' in one place and 'u to v' in another. Pick one and stick with it.",
			},
			TimeComplexity:  "O(n^2 + m): fill the grid, then one write per edge",
			SpaceComplexity: "O(n^2) for the grid, whatever the number of edges",
			Walkthrough: []WalkthroughStep{
				{Title: "Make a grid of zeros", Explanation: "Each of the n rows is its own list of n zeros. Starting at 0 means 'no edge'.", CodeSnippet: `matrix = [[0] * n for _ in range(n)]`, LineStart: 2, LineEnd: 2},
				{Title: "Mark the edge", Explanation: "For [0,1] this sets cell [0][1]. Rows are indexed by where the edge starts and columns by where it ends.", CodeSnippet: `for u, v in edges:
    matrix[u][v] = 1`, LineStart: 3, LineEnd: 4},
				{Title: "Mark the mirrored cell when undirected", Explanation: "An undirected edge is also an edge from v to u, so cell [v][u] is set too. Directed edges skip this.", CodeSnippet: `if not directed:
    matrix[v][u] = 1`, LineStart: 5, LineEnd: 6},
				{Title: "Return the grid", Explanation: "Row u now lists, cell by cell, which nodes u has an edge to.", CodeSnippet: `return matrix`, LineStart: 7, LineEnd: 7},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "When would you choose an adjacency matrix over an adjacency list?", Answer: "When the graph is dense (m close to n^2), or when you need to test 'is there an edge between u and v?' in O(1), as in Floyd-Warshall or a small fixed-size grid of relationships. For sparse graphs the matrix wastes O(n^2) memory on zeros and makes 'list a node's neighbours' cost O(n)."},
			{Prompt: "How can you tell from the matrix alone whether a graph was undirected?", Answer: "Check that it is symmetric: matrix[u][v] == matrix[v][u] for every pair. An undirected graph always produces a symmetric matrix. A directed graph can still be symmetric by chance, if every edge happens to have a partner pointing back."},
		},
	},
	{
		ID:              "drill-bfs-order",
		Kind:            KindDrill,
		Order:           3,
		Title:           "Breadth-First Search Order",
		Slug:            "drill-bfs-order",
		Difficulty:      "Easy",
		Category:        "graphs",
		Tags:            []string{"Graph", "BFS", "Queue", "Foundations"},
		RelatedChapters: []int{13},
		Description: `Given an undirected graph and a start node, walk it **breadth-first** and return the nodes in the order they are visited. Breadth-first means: the start first, then everything one edge away, then everything two edges away, and so on.

There are ` + "`" + `n` + "`" + ` nodes labelled ` + "`" + `0` + "`" + ` to ` + "`" + `n-1` + "`" + `, and each entry of ` + "`" + `edges` + "`" + ` is a pair ` + "`" + `[u, v]` + "`" + ` meaning an undirected edge. When a node has several unvisited neighbours, visit them in **increasing label order**, so there is exactly one right answer.

Only nodes reachable from ` + "`" + `start` + "`" + ` are visited, so the answer may be shorter than ` + "`" + `n` + "`" + `. Each node appears at most once.`,
		Constraints: []string{
			"1 <= n <= 1000",
			"0 <= len(edges) <= 5000",
			"0 <= u, v, start < n",
			"No self-loops and no repeated edges",
		},
		Examples: []Example{
			{Input: "n = 6, edges = [[0,1],[0,2],[1,3],[2,3],[3,4],[4,5]], start = 0", Output: "[0,1,2,3,4,5]", Explanation: "First 0, then its neighbours 1 and 2, then their new neighbour 3, then 4, then 5. Node 3 is reachable from both 1 and 2 but is visited only once."},
			{Input: "n = 6, edges = [[0,1],[1,2],[3,4]], start = 0", Output: "[0,1,2]", Explanation: "Nodes 3, 4 and 5 cannot be reached from 0, so they are not visited."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 3}, []interface{}{2, 3}, []interface{}{3, 4}, []interface{}{4, 5}}, "start": 0}, Expected: []interface{}{0, 1, 2, 3, 4, 5}, Note: "example 1: node 3 is reachable two ways but is visited once"},
			{Input: map[string]interface{}{"n": 7, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 3}, []interface{}{1, 4}, []interface{}{2, 5}, []interface{}{2, 6}}, "start": 0}, Expected: []interface{}{0, 1, 2, 3, 4, 5, 6}, Note: "example 2: a full tree, visited level by level (a stack would go 0,1,3,4,2,5,6)"},
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{3, 4}}, "start": 0}, Expected: []interface{}{0, 1, 2}, Hidden: true, Note: "nodes 3, 4 and 5 are not reachable from 0, so they are not in the answer"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 3}, []interface{}{3, 4}}, "start": 2}, Expected: []interface{}{2, 1, 3, 0, 4}, Hidden: true, Note: "start in the middle of a path: both sides are explored, nearer nodes first"},
			{Input: map[string]interface{}{"n": 1, "edges": []interface{}{}, "start": 0}, Expected: []interface{}{0}, Hidden: true, Note: "a single node"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 3}, []interface{}{0, 1}, []interface{}{0, 2}, []interface{}{2, 4}}, "start": 0}, Expected: []interface{}{0, 1, 2, 3, 4}, Hidden: true, Note: "edges are listed out of order: neighbours must be visited in increasing label order"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 3}, []interface{}{3, 0}}, "start": 0}, Expected: []interface{}{0, 1, 3, 2}, Hidden: true, Note: "a cycle: the walk must stop when it comes back around"},
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 2}, []interface{}{1, 3}, []interface{}{2, 3}, []interface{}{3, 4}, []interface{}{2, 5}}, "start": 1}, Expected: []interface{}{1, 0, 2, 3, 5, 4}, Hidden: true, Note: "a dense little graph started away from 0: each node is visited once"},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(n + m)",
		StarterCode: `def bfsOrder(n, edges, start):
    # Write your solution here
    pass`,
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: `Breadth-first means finishing everything one step away before anything two steps away. Which structure hands items back in exactly the order you put them in?`},
			{Level: 2, Type: "insight", Content: `A queue. Take a node off the front, record it, and put its unvisited neighbours on the back. The key detail: mark a node visited at the moment you put it on the queue, not when you take it off, or a node reachable two ways gets queued twice. This is BFS from Chapter 13.`, ChapterRef: 13},
			{Level: 3, Type: "plan", Content: `1. Build an adjacency list from the edges (both directions) and sort each list.
2. Make a visited list; mark ` + "`" + `start` + "`" + ` visited and put it in a queue.
3. While the queue is not empty: pop the front node and record it.
4. For each neighbour not yet visited: mark it visited now, and append it to the queue.
5. Return the recorded order.`},
			{Level: 4, Type: "skeleton", Content: `from collections import deque


def bfsOrder(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for neighbours in adj:
        neighbours.sort()
    order = []
    visited = [False] * n
    visited[start] = True
    queue = ____
    while queue:
        node = ____
        order.append(node)
        for nxt in adj[node]:
            if not visited[nxt]:
                ____
                queue.append(nxt)
    return order`},
		},
		Solution: Solution{
			Code: `from collections import deque


def bfsOrder(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for neighbours in adj:
        neighbours.sort()
    order = []
    visited = [False] * n
    visited[start] = True
    queue = deque([start])
    while queue:
        node = queue.popleft()
        order.append(node)
        for nxt in adj[node]:
            if not visited[nxt]:
                visited[nxt] = True
                queue.append(nxt)
    return order`,
			Explanation: "Build the adjacency lists, then run the queue loop of BFS. The queue holds nodes that are discovered but not yet processed, front to back in discovery order, so nodes closer to the start are always processed before farther ones.",
			KeyInsight:  "Mark a node visited when it joins the queue, not when it leaves. Then every node enters the queue at most once, so the queue stays small and the order recorded is exactly the discovery order.",
			WhyItWorks:  "Invariant: the queue holds, in order, nodes at distance d followed by nodes at distance d+1 from the start (never farther), and `visited` marks exactly the nodes that have ever been queued. A node is queued the first time any processed node reaches it, which is along a shortest path, so it comes out of the queue after every node nearer to the start. Marking at enqueue means no node is queued twice. Each node is queued and popped once and each adjacency entry is looked at once: O(n + m) time; the queue and visited list use O(n).",
			CommonMistakes: []string{
				"Marking visited only when a node is popped. A node reachable from two queued nodes then gets queued twice, and in a large graph the queue can grow to size m instead of n.",
				"Using a list's `pop()` (a stack) instead of `popleft()` of a deque. That is depth-first-like behaviour, and the order on a tree is no longer level by level.",
				"Using `list.pop(0)` as the queue. It is correct but each pop shifts the whole list, O(n); use `collections.deque`.",
				"Forgetting to sort the neighbour lists. The answer then follows the order edges were given, not the increasing-label order the drill asks for.",
			},
			TimeComplexity:  "O(n + m): each node is queued once and each edge examined twice",
			SpaceComplexity: "O(n + m) for the lists, the queue and the visited marks",
			Walkthrough: []WalkthroughStep{
				{Title: "Build sorted neighbour lists", Explanation: "Each edge goes in both lists, and sorting makes 'increasing label order' automatic when we later loop over a list.", CodeSnippet: `adj = [[] for _ in range(n)]
for u, v in edges:
    adj[u].append(v)
    adj[v].append(u)
for neighbours in adj:
    neighbours.sort()`, LineStart: 5, LineEnd: 10},
				{Title: "Start the queue with the start node", Explanation: "Mark it visited right away, because it is already in the queue.", CodeSnippet: `order = []
visited = [False] * n
visited[start] = True
queue = deque([start])`, LineStart: 11, LineEnd: 14},
				{Title: "Take the front node and record it", Explanation: "On the first example 0 is taken, and then 1, 2, 3...: the front of the queue is always the earliest-discovered node waiting.", CodeSnippet: `while queue:
    node = queue.popleft()
    order.append(node)`, LineStart: 15, LineEnd: 17},
				{Title: "Queue each new neighbour, marking it now", Explanation: "Marking before appending is what stops node 3 being added once via 1 and again via 2.", CodeSnippet: `for nxt in adj[node]:
    if not visited[nxt]:
        visited[nxt] = True
        queue.append(nxt)`, LineStart: 18, LineEnd: 21},
				{Title: "Return the visit order", Explanation: "Unreachable nodes were never queued, so they are simply not in the list.", CodeSnippet: `return order`, LineStart: 22, LineEnd: 22},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does BFS visit nodes in order of distance from the start?", Answer: "The queue is first-in-first-out, and every node is queued by a node one step nearer. So all nodes at distance d are queued before any node at distance d+1 can be queued (those are queued only while processing distance-d nodes), and they come out in the same order. That is also why BFS finds fewest-edge paths, which the next drills use."},
			{Prompt: "What goes wrong if you mark a node visited only when you pop it?", Answer: "Two nodes in the queue can both have the same unvisited neighbour, so it is queued twice. The answer can then contain duplicates unless you also check on pop, and the queue can grow far larger than n. Marking when queued gives each node exactly one place in the queue."},
		},
	},
	{
		ID:              "drill-dfs-order",
		Kind:            KindDrill,
		Order:           4,
		Title:           "Depth-First Search Order",
		Slug:            "drill-dfs-order",
		Difficulty:      "Easy",
		Category:        "graphs",
		Tags:            []string{"Graph", "DFS", "Recursion", "Foundations"},
		RelatedChapters: []int{13},
		Description: `Given an undirected graph and a start node, walk it **depth-first** and return the nodes in the order they are first visited. Depth-first means: from a node, go to its first unvisited neighbour and keep going as deep as you can, and only when you are stuck back up and try the next neighbour.

There are ` + "`" + `n` + "`" + ` nodes labelled ` + "`" + `0` + "`" + ` to ` + "`" + `n-1` + "`" + `, and each entry of ` + "`" + `edges` + "`" + ` is a pair ` + "`" + `[u, v]` + "`" + ` meaning an undirected edge. Try the neighbours of a node in **increasing label order**, so there is exactly one right answer. The order a node is first reached is a **preorder**: it is recorded before any of the nodes found through it.

Only nodes reachable from ` + "`" + `start` + "`" + ` are visited. Each node appears at most once.`,
		Constraints: []string{
			"1 <= n <= 1000",
			"0 <= len(edges) <= 5000",
			"0 <= u, v, start < n",
			"No self-loops and no repeated edges",
		},
		Examples: []Example{
			{Input: "n = 6, edges = [[0,1],[0,2],[1,3],[2,3],[3,4],[4,5]], start = 0", Output: "[0,1,3,2,4,5]", Explanation: "From 0 go to 1, then from 1 to 3, then (3's neighbours in order) 2, then 4 and 5. Compare breadth-first, which gives [0,1,2,3,4,5]."},
			{Input: "n = 4, edges = [[0,1],[1,2],[2,3],[3,0]], start = 0", Output: "[0,1,2,3]", Explanation: "A cycle: after 3, the edge back to 0 leads to an already-visited node and is ignored."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 3}, []interface{}{2, 3}, []interface{}{3, 4}, []interface{}{4, 5}}, "start": 0}, Expected: []interface{}{0, 1, 3, 2, 4, 5}, Note: "example 1: go as deep as possible first: 0,1,3,2,4,5 (BFS order would be 0,1,2,3,4,5)"},
			{Input: map[string]interface{}{"n": 7, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 3}, []interface{}{1, 4}, []interface{}{2, 5}, []interface{}{2, 6}}, "start": 0}, Expected: []interface{}{0, 1, 3, 4, 2, 5, 6}, Note: "example 2: a full tree: the whole left side comes before the right (preorder)"},
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{3, 4}}, "start": 0}, Expected: []interface{}{0, 1, 2}, Hidden: true, Note: "nodes 3, 4 and 5 are not reachable from 0, so they are not in the answer"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 3}, []interface{}{3, 4}}, "start": 2}, Expected: []interface{}{2, 1, 0, 3, 4}, Hidden: true, Note: "start in the middle of a path: it goes down to the lowest label before 3 and 4"},
			{Input: map[string]interface{}{"n": 1, "edges": []interface{}{}, "start": 0}, Expected: []interface{}{0}, Hidden: true, Note: "a single node"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 3}, []interface{}{0, 1}, []interface{}{0, 2}, []interface{}{2, 4}}, "start": 0}, Expected: []interface{}{0, 1, 2, 4, 3}, Hidden: true, Note: "edges are listed out of order: neighbours must be tried in increasing label order"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 3}, []interface{}{3, 0}}, "start": 0}, Expected: []interface{}{0, 1, 2, 3}, Hidden: true, Note: "a cycle: without a visited set the recursion never ends"},
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 2}, []interface{}{1, 3}, []interface{}{2, 3}, []interface{}{3, 4}, []interface{}{2, 5}}, "start": 1}, Expected: []interface{}{1, 0, 2, 3, 4, 5}, Hidden: true, Note: "a dense little graph started away from 0: each node is visited once"},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(n + m)",
		StarterCode: `def dfsOrder(n, edges, start):
    # Write your solution here
    pass`,
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: `Depth-first means following one path as far as it goes before trying alternatives. What do you need to remember so you never walk around a cycle forever?`},
			{Level: 2, Type: "insight", Content: `Write a helper ` + "`" + `visit(node)` + "`" + `: mark the node visited, record it, then call ` + "`" + `visit` + "`" + ` on each neighbour that is not yet visited. The recursion's call stack remembers where to back up to. The visited marks are what stop a cycle from looping. This is the DFS template from Chapter 13, the same walk as the tree traversals with a visited check added.`, ChapterRef: 13},
			{Level: 3, Type: "plan", Content: `1. Build adjacency lists in both directions and sort each list.
2. Make ` + "`" + `order` + "`" + ` and a visited list.
3. Write ` + "`" + `visit(node)` + "`" + `: mark it visited and record it.
4. In ` + "`" + `visit` + "`" + `, loop over the node's neighbours in order and recurse into each one that is not visited.
5. Call ` + "`" + `visit(start)` + "`" + ` and return ` + "`" + `order` + "`" + `.`},
			{Level: 4, Type: "skeleton", Content: `def dfsOrder(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for neighbours in adj:
        neighbours.sort()
    order = []
    visited = [False] * n

    def visit(node):
        ____
        order.append(node)
        for nxt in adj[node]:
            if ____:
                ____

    visit(start)
    return order`},
		},
		Solution: Solution{
			Code: `def dfsOrder(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for neighbours in adj:
        neighbours.sort()
    order = []
    visited = [False] * n

    def visit(node):
        visited[node] = True
        order.append(node)
        for nxt in adj[node]:
            if not visited[nxt]:
                visit(nxt)

    visit(start)
    return order`,
			Explanation: "A recursive helper marks a node, records it, and then recurses into each neighbour that has not been visited yet. The call stack is the 'path back' that depth-first search needs, and the visited marks keep it from walking around cycles.",
			KeyInsight:  "The tree traversals never needed a visited set because a tree has no cycles. A graph can lead you back to a node you already visited, so every node is marked when first entered and never entered again.",
			WhyItWorks:  "By induction on the number of unvisited nodes reachable. `visit(x)` records x, then handles each neighbour in order: a neighbour already marked is skipped, and an unmarked one is fully explored (everything reachable from it that is still unvisited) before the next neighbour is looked at. So when `visit(x)` returns, every node reachable from x without passing through an already-visited node has been visited, and each exactly once because it is marked before anything else happens in its call. Each node's call runs once and scans its list once: O(n + m) time, with up to n frames on the stack in the worst case.",
			CommonMistakes: []string{
				"Not keeping a visited set. On any graph with a cycle the recursion goes round forever until Python raises a recursion error.",
				"Marking visited after the recursive calls return instead of on entry. A neighbour deeper in the walk can then reach this node again before it is marked.",
				"Writing the walk as a queue loop. That is breadth-first: on a graph with branches the order differs (0,1,2,3 instead of 0,1,3,2 in the first example).",
				"Forgetting to sort the neighbour lists, so the order follows the edge list instead of increasing labels.",
			},
			TimeComplexity:  "O(n + m): each node is entered once and each adjacency entry scanned once",
			SpaceComplexity: "O(n + m) for the lists, plus up to n frames of recursion",
			Walkthrough: []WalkthroughStep{
				{Title: "Build sorted neighbour lists", Explanation: "Each edge goes in both lists. Sorting means the recursion later tries neighbours smallest-label first.", CodeSnippet: `adj = [[] for _ in range(n)]
for u, v in edges:
    adj[u].append(v)
    adj[v].append(u)
for neighbours in adj:
    neighbours.sort()`, LineStart: 2, LineEnd: 7},
				{Title: "Set up the answer and the visited marks", Explanation: "One shared list and one visited list, used by every call.", CodeSnippet: `order = []
visited = [False] * n`, LineStart: 8, LineEnd: 9},
				{Title: "Mark and record on entry", Explanation: "Marking first is what stops a cycle from re-entering this node. Recording now makes the order a preorder.", CodeSnippet: `visited[node] = True
order.append(node)`, LineStart: 12, LineEnd: 13},
				{Title: "Go deeper into each unvisited neighbour", Explanation: "On the first example, `visit(0)` goes into 1 and, inside it, into 3 before coming back to try 2. A neighbour that is already visited is skipped.", CodeSnippet: `for nxt in adj[node]:
    if not visited[nxt]:
        visit(nxt)`, LineStart: 14, LineEnd: 16},
				{Title: "Start the walk", Explanation: "One call from the start explores everything reachable.", CodeSnippet: `visit(start)
return order`, LineStart: 18, LineEnd: 19},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can a tree traversal skip the visited set but a graph DFS cannot?", Answer: "In a tree every node has exactly one path from the root, so you can never reach a node twice, and children only point downward. A graph can have cycles or two routes to the same node, so without a record of visited nodes DFS would re-enter nodes and, on a cycle, never stop."},
			{Prompt: "Both BFS and DFS visit each reachable node once. What decides the different order?", Answer: "Which node is taken next. BFS takes the oldest discovered node (a queue), so it explores in rings of equal distance. DFS takes the newest one (the recursion's call stack), so it follows one path as deep as it goes before backing up."},
		},
	},
	{
		ID:              "drill-dfs-iterative",
		Kind:            KindDrill,
		Order:           5,
		Title:           "Depth-First Search with a Stack",
		Slug:            "drill-dfs-iterative",
		Difficulty:      "Easy",
		Category:        "graphs",
		Tags:            []string{"Graph", "DFS", "Stack", "Foundations"},
		RelatedChapters: []int{13},
		Description: `Do the previous drill again, this time **without recursion**. Given an undirected graph and a start node, return the nodes in depth-first order, with neighbours tried in increasing label order, using an explicit stack in place of the call stack.

The answer must be **exactly the same** as the recursive depth-first order: from each node go to its smallest unvisited neighbour first and as deep as possible, and only back up when stuck.

There are ` + "`" + `n` + "`" + ` nodes labelled ` + "`" + `0` + "`" + ` to ` + "`" + `n-1` + "`" + ` and each entry of ` + "`" + `edges` + "`" + ` is a pair ` + "`" + `[u, v]` + "`" + ` meaning an undirected edge. Only nodes reachable from ` + "`" + `start` + "`" + ` are visited, each once.

Why bother? Python stops a recursion at about 20,000 calls here, and a long path graph (a linked list shaped like a graph) is deeper than that. A stack you manage yourself has no such limit.`,
		Constraints: []string{
			"1 <= n <= 30000",
			"0 <= len(edges) <= 30000",
			"0 <= u, v, start < n",
			"No self-loops and no repeated edges",
		},
		Examples: []Example{
			{Input: "n = 6, edges = [[0,1],[0,2],[1,3],[2,3],[3,4],[4,5]], start = 0", Output: "[0,1,3,2,4,5]", Explanation: "Same as the recursive walk: 0, then 1, then 3, then back at 3 the next smallest unvisited neighbour is 2, then 4 and 5."},
			{Input: "n = 5, edges = [[0,3],[0,1],[0,2],[2,4]], start = 0", Output: "[0,1,2,4,3]", Explanation: "0's neighbours in order are 1, 2, 3. Go to 1 (a dead end), then 2 and on to 4, and finally 3."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 3}, []interface{}{2, 3}, []interface{}{3, 4}, []interface{}{4, 5}}, "start": 0}, Expected: []interface{}{0, 1, 3, 2, 4, 5}, Note: "example 1: go as deep as possible first: 0,1,3,2,4,5 (BFS order would be 0,1,2,3,4,5)"},
			{Input: map[string]interface{}{"n": 7, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 3}, []interface{}{1, 4}, []interface{}{2, 5}, []interface{}{2, 6}}, "start": 0}, Expected: []interface{}{0, 1, 3, 4, 2, 5, 6}, Note: "example 2: a full tree: the whole left side comes before the right (preorder)"},
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{3, 4}}, "start": 0}, Expected: []interface{}{0, 1, 2}, Hidden: true, Note: "nodes 3, 4 and 5 are not reachable from 0, so they are not in the answer"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 3}, []interface{}{3, 4}}, "start": 2}, Expected: []interface{}{2, 1, 0, 3, 4}, Hidden: true, Note: "start in the middle of a path: it goes down to the lowest label before 3 and 4"},
			{Input: map[string]interface{}{"n": 1, "edges": []interface{}{}, "start": 0}, Expected: []interface{}{0}, Hidden: true, Note: "a single node"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 3}, []interface{}{0, 1}, []interface{}{0, 2}, []interface{}{2, 4}}, "start": 0}, Expected: []interface{}{0, 1, 2, 4, 3}, Hidden: true, Note: "edges are listed out of order: neighbours must be tried in increasing label order"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 3}, []interface{}{3, 0}}, "start": 0}, Expected: []interface{}{0, 1, 2, 3}, Hidden: true, Note: "a cycle: without a visited set the recursion never ends"},
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 2}, []interface{}{1, 3}, []interface{}{2, 3}, []interface{}{3, 4}, []interface{}{2, 5}}, "start": 1}, Expected: []interface{}{1, 0, 2, 3, 4, 5}, Hidden: true, Note: "a dense little graph started away from 0: each node is visited once"},
			{Input: map[string]interface{}{"n": 25000, "edges": graphDrillPathEdges(25000), "start": 0}, Expected: graphDrillCountUp(25000), Hidden: true, Note: "a 25,000-node path graph, deeper than Python's recursion limit: only an explicit stack gets through"},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(n + m)",
		StarterCode: `def dfsOrderIterative(n, edges, start):
    # Write your solution here
    pass`,
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: `A recursive walk remembers where to come back to using the call stack. If you keep your own stack, what goes on it, and in which order, so the smallest neighbour comes off first?`},
			{Level: 2, Type: "insight", Content: `Push the start. Then loop: pop a node; if it was already visited, skip it; otherwise mark and record it and push its neighbours. A stack returns the last thing pushed, so push the neighbours from largest to smallest and the smallest is popped first. The same node can be pushed more than once, so the visited check belongs when you pop. See the DFS template in Chapter 13.`, ChapterRef: 13},
			{Level: 3, Type: "plan", Content: `1. Build sorted adjacency lists.
2. Start a stack holding ` + "`" + `start` + "`" + `.
3. Pop a node. If it is visited, continue to the next pop.
4. Mark it visited and record it.
5. Push its neighbours in decreasing order, skipping visited ones.
6. When the stack is empty, return the order.`},
			{Level: 4, Type: "skeleton", Content: `def dfsOrderIterative(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for neighbours in adj:
        neighbours.sort()
    order = []
    visited = [False] * n
    stack = [start]
    while stack:
        node = ____
        if visited[node]:
            ____
        visited[node] = True
        order.append(node)
        for nxt in ____(adj[node]):
            if not visited[nxt]:
                stack.append(nxt)
    return order`},
		},
		Solution: Solution{
			Code: `def dfsOrderIterative(n, edges, start):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    for neighbours in adj:
        neighbours.sort()
    order = []
    visited = [False] * n
    stack = [start]
    while stack:
        node = stack.pop()
        if visited[node]:
            continue
        visited[node] = True
        order.append(node)
        for nxt in reversed(adj[node]):
            if not visited[nxt]:
                stack.append(nxt)
    return order`,
			Explanation: "Replace the call stack with a list used as a stack. A node is only counted as visited when it is popped, and its neighbours are pushed from largest label to smallest so that the smallest is popped next, exactly the order the recursion would have chosen.",
			KeyInsight:  "Pushing neighbours in reverse order makes the stack pop them smallest first, and checking `visited` at pop time (not at push time) is what reproduces recursive order, since a node's correct place is where the walk actually reaches it, not where it was first noticed.",
			WhyItWorks:  "Compare with the recursion. `visit(x)` goes into x's neighbours one at a time, smallest first, and each is explored fully before the next. On the stack, x's unvisited neighbours sit on top in order smallest-on-top, and any node pushed while exploring the top one lands above the rest, so the top neighbour's whole subtree is finished before the next neighbour is popped. A neighbour that got visited meanwhile is skipped when popped. So pops that are not skipped happen in the same order as the first visits of the recursion. Each node is expanded once, and pushes are bounded by the number of adjacency entries: O(n + m) time and space.",
			CommonMistakes: []string{
				"Pushing neighbours in increasing order. The largest is then popped first and the whole answer comes out in a different order from the recursive walk.",
				"Marking nodes visited when they are pushed. That gives a valid depth-first-like traversal but not the same order: a node can be pushed early by a far-away ancestor and visited too soon.",
				"Skipping the visited check at pop time. A node pushed twice is then recorded twice.",
				"Using the recursive version because it passes small tests. A path graph of 25,000 nodes is deeper than Python's recursion limit and fails with a recursion error.",
			},
			TimeComplexity:  "O(n + m): each node is expanded once and each adjacency entry is pushed at most once",
			SpaceComplexity: "O(n + m) for the lists, the visited marks and the stack",
			Walkthrough: []WalkthroughStep{
				{Title: "Build sorted neighbour lists", Explanation: "Each edge goes in both lists, and sorting fixes the order of neighbours.", CodeSnippet: `adj = [[] for _ in range(n)]
for u, v in edges:
    adj[u].append(v)
    adj[v].append(u)
for neighbours in adj:
    neighbours.sort()`, LineStart: 2, LineEnd: 7},
				{Title: "Start with the start node on the stack", Explanation: "The stack holds nodes still to be explored, newest on top.", CodeSnippet: `order = []
visited = [False] * n
stack = [start]`, LineStart: 8, LineEnd: 10},
				{Title: "Pop, and skip nodes already visited", Explanation: "The same node can be on the stack several times; only the first pop counts.", CodeSnippet: `node = stack.pop()
if visited[node]:
    continue`, LineStart: 12, LineEnd: 14},
				{Title: "Mark and record", Explanation: "This is the moment the node is really reached, so the order is a preorder just as in the recursion.", CodeSnippet: `visited[node] = True
order.append(node)`, LineStart: 15, LineEnd: 16},
				{Title: "Push neighbours largest first", Explanation: "Reversing means the smallest label ends up on top and is popped next. On [0,3],[0,1],[0,2] this pushes 3, 2, 1, and 1 comes off first.", CodeSnippet: `for nxt in reversed(adj[node]):
    if not visited[nxt]:
        stack.append(nxt)`, LineStart: 17, LineEnd: 19},
				{Title: "Return the order", Explanation: "When the stack is empty, every reachable node has been recorded once.", CodeSnippet: `return order`, LineStart: 20, LineEnd: 20},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must neighbours be pushed in decreasing order to match recursion?", Answer: "Because a stack returns the most recently pushed item first. The recursion tries the smallest neighbour first, so the smallest must be pushed last and sit on top. Pushing in increasing order pops the largest neighbour first and reverses the exploration order."},
			{Prompt: "Why does the visited check go at pop time and not at push time?", Answer: "A node can be pushed by several nodes before it is ever popped. If it were marked at push time, the later pushers would skip it, and it would be visited from the earliest pusher's position rather than where recursion reaches it, which changes the order (and the result is not true DFS order). Checking at pop time lets a node stay on the stack in several places, and only the one popped first, the deepest-first reach, counts."},
		},
	},
	{
		ID:              "drill-shortest-path-unweighted",
		Kind:            KindDrill,
		Order:           6,
		Title:           "Shortest Path in an Unweighted Graph",
		Slug:            "drill-shortest-path-unweighted",
		Difficulty:      "Easy",
		Category:        "graphs",
		Tags:            []string{"Graph", "BFS", "Shortest Path", "Foundations"},
		RelatedChapters: []int{13},
		Description: `Given an undirected graph, a ` + "`" + `start` + "`" + ` node and a ` + "`" + `target` + "`" + ` node, return the **fewest number of edges** on any path from ` + "`" + `start` + "`" + ` to ` + "`" + `target` + "`" + `. Return ` + "`" + `-1` + "`" + ` if there is no path, and ` + "`" + `0` + "`" + ` if ` + "`" + `start` + "`" + ` and ` + "`" + `target` + "`" + ` are the same node.

There are ` + "`" + `n` + "`" + ` nodes labelled ` + "`" + `0` + "`" + ` to ` + "`" + `n-1` + "`" + `, and each entry of ` + "`" + `edges` + "`" + ` is a pair ` + "`" + `[u, v]` + "`" + ` meaning an undirected edge. Every edge counts as 1; there are no weights.`,
		Constraints: []string{
			"1 <= n <= 1000",
			"0 <= len(edges) <= 5000",
			"0 <= u, v, start, target < n",
			"No self-loops and no repeated edges",
		},
		Examples: []Example{
			{Input: "n = 6, edges = [[0,1],[0,2],[1,3],[2,3],[3,4],[4,5]], start = 0, target = 5", Output: "4", Explanation: "The path 0 - 1 - 3 - 4 - 5 uses four edges, and no path is shorter."},
			{Input: "n = 5, edges = [[0,1],[2,3]], start = 0, target = 3", Output: "-1", Explanation: "0 and 3 are in different connected pieces, so no path exists."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 3}, []interface{}{2, 3}, []interface{}{3, 4}, []interface{}{4, 5}}, "start": 0, "target": 5}, Expected: 4, Note: "example 1: 0 -> 1 -> 3 -> 4 -> 5 is four edges"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{2, 3}}, "start": 0, "target": 3}, Expected: -1, Note: "example 2: 3 is in a different component: -1"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}}, "start": 2, "target": 2}, Expected: 0, Hidden: true, Note: "start equals target: 0 edges, even before looking at any edge"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 3}, []interface{}{0, 3}}, "start": 0, "target": 3}, Expected: 1, Hidden: true, Note: "a direct edge beats the long way round (a depth-first search finds the long way first)"},
			{Input: map[string]interface{}{"n": 7, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 3}, []interface{}{3, 4}, []interface{}{4, 5}, []interface{}{5, 6}, []interface{}{0, 4}}, "start": 0, "target": 6}, Expected: 3, Hidden: true, Note: "a shortcut in the middle of a long path: 0 -> 4 -> 5 -> 6"},
			{Input: map[string]interface{}{"n": 1, "edges": []interface{}{}, "start": 0, "target": 0}, Expected: 0, Hidden: true, Note: "one node, start equals target"},
			{Input: map[string]interface{}{"n": 2, "edges": []interface{}{[]interface{}{1, 0}}, "start": 0, "target": 1}, Expected: 1, Hidden: true, Note: "a target next to the start, the edge written in the other direction"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}}, "start": 0, "target": 3}, Expected: -1, Hidden: true, Note: "the target has no edges at all: -1"},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(n + m)",
		StarterCode: `def shortestPath(n, edges, start, target):
    # Write your solution here
    pass`,
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: `Which walk reaches every node one edge away, then every node two edges away, in that order? What does that mean the first time you reach the target?`},
			{Level: 2, Type: "insight", Content: `Run BFS from ` + "`" + `start` + "`" + ` and store each node's distance when you first discover it: a neighbour's distance is the current node's distance plus 1. BFS reaches nodes in order of distance, so the first time you reach the target is by a shortest path. A distance list that starts at -1 doubles as the visited marks. See BFS and shortest paths in Chapter 13.`, ChapterRef: 13},
			{Level: 3, Type: "plan", Content: `1. Build adjacency lists (both directions).
2. Make ` + "`" + `dist` + "`" + ` filled with -1, set ` + "`" + `dist[start] = 0` + "`" + `, and queue the start.
3. Pop a node; if it is the target return its distance.
4. For each neighbour with distance -1, set its distance to this node's distance plus 1 and queue it.
5. If the queue empties without reaching the target, return -1.`},
			{Level: 4, Type: "skeleton", Content: `from collections import deque


def shortestPath(n, edges, start, target):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    dist = [-1] * n
    dist[start] = ____
    queue = deque([start])
    while queue:
        node = queue.popleft()
        if node == target:
            return ____
        for nxt in adj[node]:
            if dist[nxt] == -1:
                dist[nxt] = ____
                queue.append(nxt)
    return -1`},
		},
		Solution: Solution{
			Code: `from collections import deque


def shortestPath(n, edges, start, target):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
        adj[v].append(u)
    dist = [-1] * n
    dist[start] = 0
    queue = deque([start])
    while queue:
        node = queue.popleft()
        if node == target:
            return dist[node]
        for nxt in adj[node]:
            if dist[nxt] == -1:
                dist[nxt] = dist[node] + 1
                queue.append(nxt)
    return -1`,
			Explanation: "BFS finds nodes in order of how many edges away they are, so recording a distance as each node is discovered gives every node its fewest-edge distance. The first time the target comes off the queue, its distance is the answer; if the queue runs dry first, the target is unreachable.",
			KeyInsight:  "In a graph where every edge counts the same, BFS is a shortest-path algorithm: nodes leave the queue in order of distance, and a node's first discovery is along a shortest path, so its distance never needs updating.",
			WhyItWorks:  "Invariant: when a node at distance d is popped, every node at distance less than d has already been discovered with its correct distance, and the queue holds only nodes of distance d and d+1. A newly discovered neighbour of a distance-d node is at most d+1, and it cannot be less than d+1 or it would already have been discovered from an earlier node. So the stored distance is exact. If the target is reachable it is discovered after at most n-1 rounds; if the queue empties, the whole reachable piece has been seen without the target, so no path exists. Time O(n + m).",
			CommonMistakes: []string{
				"Using a depth-first search and returning the length of the first path found. DFS finds some path, often a long one; on a cycle with a direct edge it can report the long way round.",
				"Counting nodes on the path instead of edges, which is off by one. A path of 3 edges visits 4 nodes.",
				"Forgetting the case start == target. Here it works naturally (distance 0), but a version that only returns when it discovers the target from a neighbour returns -1.",
				"Returning 0 when the target is unreachable, which looks identical to start == target. Use -1 for 'no path'.",
			},
			TimeComplexity:  "O(n + m): each node is queued once and each edge examined twice",
			SpaceComplexity: "O(n + m) for the lists, the distance list and the queue",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the adjacency lists", Explanation: "Each undirected edge is stored in both directions so you can step either way along it.", CodeSnippet: `adj = [[] for _ in range(n)]
for u, v in edges:
    adj[u].append(v)
    adj[v].append(u)`, LineStart: 5, LineEnd: 8},
				{Title: "Start at distance 0", Explanation: "The distance list is also the visited record: -1 means not yet discovered. The start is zero edges from itself.", CodeSnippet: `dist = [-1] * n
dist[start] = 0
queue = deque([start])`, LineStart: 9, LineEnd: 11},
				{Title: "Stop when the target is reached", Explanation: "Nodes leave the queue nearest first, so the first time the target is popped its distance is the fewest edges.", CodeSnippet: `node = queue.popleft()
if node == target:
    return dist[node]`, LineStart: 13, LineEnd: 15},
				{Title: "Discover neighbours one edge further", Explanation: "A node found from `node` is one edge beyond it. In the first example 3 is discovered from 1 at distance 2 and 2 does not change it.", CodeSnippet: `for nxt in adj[node]:
    if dist[nxt] == -1:
        dist[nxt] = dist[node] + 1
        queue.append(nxt)`, LineStart: 16, LineEnd: 19},
				{Title: "No path", Explanation: "The queue is empty, so everything reachable has been tried and the target was not among it.", CodeSnippet: `return -1`, LineStart: 20, LineEnd: 20},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the first time BFS reaches the target guaranteed to be a shortest path?", Answer: "BFS processes nodes in order of distance: everything one edge away, then two, and so on. A node is discovered from the earliest processed node that touches it, which lies on a path of minimum length. So the distance stored when a node is first discovered is already the smallest possible."},
			{Prompt: "What changes if the edges have different weights?", Answer: "BFS is no longer enough, because the path with fewest edges is not always the cheapest. You need Dijkstra's algorithm (Chapter 14), which replaces the plain queue with a min-heap ordered by total distance."},
		},
	},
	{
		ID:              "drill-flood-fill",
		Number:          733,
		Kind:            KindDrill,
		Order:           7,
		Title:           "Flood Fill",
		Slug:            "flood-fill",
		Difficulty:      "Easy",
		Category:        "graphs",
		Tags:            []string{"Graph", "Grid", "DFS", "BFS", "Matrix", "Foundations"},
		RelatedChapters: []int{13},
		Description: `An image is a grid of integers, where each number is a pixel's colour. You are given the image and a starting pixel at row ` + "`" + `sr` + "`" + `, column ` + "`" + `sc` + "`" + `, and a new ` + "`" + `color` + "`" + `.

Repaint the starting pixel and every pixel **connected** to it that has the same original colour with the new colour. Two pixels are connected if they are next to each other up, down, left or right (not diagonally) and share that original colour; connection can chain through any number of pixels.

Return the modified image. This is the paint-bucket tool in an image editor, and it is graph traversal where every pixel is a node and its four neighbours are its edges.`,
		Constraints: []string{
			"1 <= rows, columns <= 50",
			"0 <= image[i][j], color < 65536",
			"0 <= sr < rows and 0 <= sc < columns",
		},
		Examples: []Example{
			{Input: "image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2", Output: "[[2,2,2],[2,2,0],[2,0,1]]", Explanation: "The 1s connected to the centre pixel become 2. The 1 in the bottom-right corner touches the region only diagonally, so it stays 1."},
			{Input: "image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0", Output: "[[0,0,0],[0,0,0]]", Explanation: "The new colour is the colour already there, so nothing changes."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"image": []interface{}{[]interface{}{1, 1, 1}, []interface{}{1, 1, 0}, []interface{}{1, 0, 1}}, "sr": 1, "sc": 1, "color": 2}, Expected: []interface{}{[]interface{}{2, 2, 2}, []interface{}{2, 2, 0}, []interface{}{2, 0, 1}}, Note: "example 1: the middle region of 1s is repainted; the 1 in the corner touches it only diagonally, so it stays"},
			{Input: map[string]interface{}{"image": []interface{}{[]interface{}{0, 0, 0}, []interface{}{0, 0, 0}}, "sr": 0, "sc": 0, "color": 0}, Expected: []interface{}{[]interface{}{0, 0, 0}, []interface{}{0, 0, 0}}, Note: "example 2: the new colour equals the old one: return at once (the fill would never stop)"},
			{Input: map[string]interface{}{"image": []interface{}{[]interface{}{5}}, "sr": 0, "sc": 0, "color": 7}, Expected: []interface{}{[]interface{}{7}}, Hidden: true, Note: "a single pixel"},
			{Input: map[string]interface{}{"image": []interface{}{[]interface{}{1, 0}, []interface{}{0, 1}}, "sr": 0, "sc": 0, "color": 3}, Expected: []interface{}{[]interface{}{3, 0}, []interface{}{0, 1}}, Hidden: true, Note: "the other 1 touches the start only diagonally, so it is not repainted"},
			{Input: map[string]interface{}{"image": []interface{}{[]interface{}{1, 1, 1}, []interface{}{1, 1, 1}}, "sr": 1, "sc": 2, "color": 9}, Expected: []interface{}{[]interface{}{9, 9, 9}, []interface{}{9, 9, 9}}, Hidden: true, Note: "the whole image is one region"},
			{Input: map[string]interface{}{"image": []interface{}{[]interface{}{1, 1, 0, 1}, []interface{}{1, 0, 0, 1}, []interface{}{0, 0, 1, 1}}, "sr": 0, "sc": 3, "color": 4}, Expected: []interface{}{[]interface{}{1, 1, 0, 4}, []interface{}{1, 0, 0, 4}, []interface{}{0, 0, 4, 4}}, Hidden: true, Note: "two separate regions of 1s: only the one containing the start changes"},
			{Input: map[string]interface{}{"image": []interface{}{[]interface{}{2, 2, 2}, []interface{}{2, 1, 2}, []interface{}{2, 2, 2}}, "sr": 1, "sc": 1, "color": 2}, Expected: []interface{}{[]interface{}{2, 2, 2}, []interface{}{2, 2, 2}, []interface{}{2, 2, 2}}, Hidden: true, Note: "the new colour equals the surrounding colour: the single 1 becomes 2"},
			{Input: map[string]interface{}{"image": []interface{}{[]interface{}{1, 1}, []interface{}{1, 0}}, "sr": 0, "sc": 0, "color": 0}, Expected: []interface{}{[]interface{}{0, 0}, []interface{}{0, 0}}, Hidden: true, Note: "the new colour 0 is already present elsewhere in the image: the region of 1s still gets repainted"},
		},
		TimeComplexity:  "O(rows * cols)",
		SpaceComplexity: "O(rows * cols)",
		StarterCode: `def floodFill(image, sr, sc, color):
    # Write your solution here
    pass`,
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: `Painting a pixel does not stop the paint: it spreads to neighbours of the same old colour. What do you need so that spreading ends, even when the new colour equals the old one?`},
			{Level: 2, Type: "insight", Content: `This is DFS on a grid. From the start pixel, repaint it and then try its four neighbours, stopping at anything outside the image or not of the original colour. The repaint itself serves as the visited mark, because a repainted pixel no longer has the old colour. But if the new colour equals the old one, repainting changes nothing and the search never ends, so return early. See grids as graphs in Chapter 13.`, ChapterRef: 13},
			{Level: 3, Type: "plan", Content: `1. Remember the starting pixel's original colour.
2. If it already equals the new colour, return the image unchanged.
3. Write ` + "`" + `fill(r, c)` + "`" + `: return if outside the grid or the pixel is not the original colour.
4. Otherwise paint it, then call ` + "`" + `fill` + "`" + ` on the four neighbours.
5. Call ` + "`" + `fill(sr, sc)` + "`" + ` and return the image.`},
			{Level: 4, Type: "skeleton", Content: `def floodFill(image, sr, sc, color):
    old = image[sr][sc]
    if ____:
        return image
    rows, cols = len(image), len(image[0])

    def fill(r, c):
        if r < 0 or r >= rows or c < 0 or c >= cols or ____:
            return
        image[r][c] = ____
        fill(r + 1, c)
        fill(r - 1, c)
        fill(r, c + 1)
        fill(r, c - 1)

    fill(sr, sc)
    return image`},
		},
		Solution: Solution{
			Code: `def floodFill(image, sr, sc, color):
    old = image[sr][sc]
    if old == color:
        return image
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
    return image`,
			Explanation: "Treat each pixel as a node whose edges go to its four neighbours. A DFS from the start paints each pixel as it enters it, and refuses to enter pixels that are off the grid or not the original colour, so it spreads exactly over the connected region of that colour.",
			KeyInsight:  "Painting a pixel is also marking it visited, because it no longer has the old colour. That works only if the new colour differs from the old one, so the case `old == color` must return before the search starts.",
			WhyItWorks:  "Invariant: a pixel is painted exactly when `fill` is called on it while it holds the old colour, and `fill` then calls itself on all four neighbours. By induction on the distance from the start, every pixel of the old colour connected to the start is reached through a chain of old-coloured neighbours, so it is painted; a pixel not connected is never reached, because the recursion only moves into old-coloured neighbours. A pixel is entered while still old-coloured at most once, since entering repaints it with a colour different from old. Each pixel does O(1) work: O(rows * cols) time, and the recursion is at most that deep.",
			CommonMistakes: []string{
				"Skipping the `old == color` check. Repainting a pixel with the colour it already has leaves it 'unvisited', so the four calls bounce between neighbours forever and Python raises a recursion error.",
				"Also spreading to diagonal neighbours. Connection here is up, down, left and right only.",
				"Comparing neighbours against the new colour instead of the original. Remember the original colour before painting anything, because the start pixel's colour changes.",
				"Forgetting the bounds checks before reading `image[r][c]`, which raises an IndexError at the edge (or silently wraps around on a negative index).",
			},
			TimeComplexity:  "O(rows * cols): each pixel is painted at most once and does four checks",
			SpaceComplexity: "O(rows * cols) for the recursion stack in the worst case",
			Walkthrough: []WalkthroughStep{
				{Title: "Remember the original colour and guard", Explanation: "Everything depends on the colour you are replacing. If it already equals the new colour, return now: the fill would never mark anything as done.", CodeSnippet: `old = image[sr][sc]
if old == color:
    return image`, LineStart: 2, LineEnd: 4},
				{Title: "Know the grid size", Explanation: "The bounds checks below need the number of rows and columns.", CodeSnippet: `rows, cols = len(image), len(image[0])`, LineStart: 5, LineEnd: 5},
				{Title: "Stop at the edge or at the wrong colour", Explanation: "A pixel outside the grid, or one that is not the original colour (a different colour, or already repainted), is not part of the region.", CodeSnippet: `if r < 0 or r >= rows or c < 0 or c >= cols or image[r][c] != old:
    return`, LineStart: 8, LineEnd: 9},
				{Title: "Paint, then spread to the four neighbours", Explanation: "Painting turns the pixel into the new colour, so it can never match `old` again. On example 1 this spreads from the centre over all the 1s that touch it by an edge.", CodeSnippet: `image[r][c] = color
fill(r + 1, c)
fill(r - 1, c)
fill(r, c + 1)
fill(r, c - 1)`, LineStart: 10, LineEnd: 14},
				{Title: "Start the fill", Explanation: "One call from the start pixel does the whole region.", CodeSnippet: `fill(sr, sc)
return image`, LineStart: 16, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{
					Name: "Iterative with a queue (BFS)",
					Code: `from collections import deque


def floodFill(image, sr, sc, color):
    old = image[sr][sc]
    if old == color:
        return image
    rows, cols = len(image), len(image[0])
    image[sr][sc] = color
    queue = deque([(sr, sc)])
    while queue:
        r, c = queue.popleft()
        for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
            if 0 <= nr < rows and 0 <= nc < cols and image[nr][nc] == old:
                image[nr][nc] = color
                queue.append((nr, nc))
    return image`,
					TimeComplexity:  "O(rows * cols)",
					SpaceComplexity: "O(rows * cols)",
					Explanation:     "The same spreading with a queue and no recursion. Paint a pixel when it is queued so that it is never queued twice. Useful when a grid could be so large that recursion would go too deep.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the case where the new colour equals the old one need its own check?", Answer: "The algorithm uses 'the pixel still has the old colour' as its visited mark. If the new colour is the old colour, painting does not change that, so pixels never look visited and the four recursive calls keep re-entering each other's pixels without end. Returning early is correct anyway, since nothing needs to change."},
			{Prompt: "How is this the same as the graph traversals earlier in the set?", Answer: "Each pixel is a node, its edges go to its four neighbours, and 'visited' is stored in the grid itself (the colour change) instead of a separate list. The recursion is the DFS template from the DFS drill, and the queue variant is the BFS template."},
		},
	},
	{
		ID:              "drill-topological-sort",
		Kind:            KindDrill,
		Order:           8,
		Title:           "Topological Sort",
		Slug:            "drill-topological-sort",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Topological Sort", "BFS", "Directed Graph", "Foundations"},
		RelatedChapters: []int{13},
		Description: `You have ` + "`" + `n` + "`" + ` tasks labelled ` + "`" + `0` + "`" + ` to ` + "`" + `n-1` + "`" + `. Each entry ` + "`" + `[u, v]` + "`" + ` of ` + "`" + `edges` + "`" + ` is a **directed** edge meaning task ` + "`" + `u` + "`" + ` must be done before task ` + "`" + `v` + "`" + `.

Return an order in which all ` + "`" + `n` + "`" + ` tasks can be done so that every edge goes **forward** (` + "`" + `u` + "`" + ` appears before ` + "`" + `v` + "`" + ` for every edge). If the requirements contain a cycle, so that no such order exists, return an empty list ` + "`" + `[]` + "`" + `.

Several orders can be valid, and any valid one is accepted.`,
		Constraints: []string{
			"1 <= n <= 1000",
			"0 <= len(edges) <= 5000",
			"0 <= u, v < n",
			"No self-loops and no repeated edges",
		},
		Examples: []Example{
			{Input: "n = 6, edges = [[5,2],[5,0],[4,0],[4,1],[2,3],[3,1]]", Output: "[4,5,0,2,3,1]", Explanation: "5 before 2 and 0, 4 before 0 and 1, 2 before 3, 3 before 1. Other orders such as [5,4,2,0,3,1] are also valid."},
			{Input: "n = 3, edges = [[0,1],[1,2],[2,0]]", Output: "[]", Explanation: "0 must precede 1, 1 must precede 2, and 2 must precede 0: no task can go first, so there is no valid order."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 6, "edges": []interface{}{[]interface{}{5, 2}, []interface{}{5, 0}, []interface{}{4, 0}, []interface{}{4, 1}, []interface{}{2, 3}, []interface{}{3, 1}}}, Expected: []interface{}{4, 5, 2, 0, 3, 1}, Note: "example 1: several valid orders exist: any order with every edge going forward is accepted"},
			{Input: map[string]interface{}{"n": 3, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 0}}}, Expected: []interface{}{}, Note: "example 2: a cycle: no valid order exists, so the answer is []"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{0, 2}, []interface{}{1, 3}, []interface{}{2, 3}}}, Expected: []interface{}{0, 1, 2, 3}, Hidden: true, Note: "a diamond: 0 first, 3 last, 1 and 2 in either order"},
			{Input: map[string]interface{}{"n": 3, "edges": []interface{}{}}, Expected: []interface{}{0, 1, 2}, Hidden: true, Note: "no edges: every order is valid, but all three nodes must appear"},
			{Input: map[string]interface{}{"n": 5, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 1}, []interface{}{3, 4}}}, Expected: []interface{}{}, Hidden: true, Note: "a cycle (1 <-> 2) in one part of the graph while the rest could be ordered: still []"},
			{Input: map[string]interface{}{"n": 1, "edges": []interface{}{}}, Expected: []interface{}{0}, Hidden: true, Note: "a single node"},
			{Input: map[string]interface{}{"n": 2, "edges": []interface{}{[]interface{}{1, 0}}}, Expected: []interface{}{1, 0}, Hidden: true, Note: "one edge written from the higher label: [1, 0] is the only valid order"},
			{Input: map[string]interface{}{"n": 4, "edges": []interface{}{[]interface{}{3, 2}, []interface{}{2, 1}, []interface{}{1, 0}}}, Expected: []interface{}{3, 2, 1, 0}, Hidden: true, Note: "a chain against the label order: [3, 2, 1, 0] is the only valid order"},
			{Input: map[string]interface{}{"n": 2, "edges": []interface{}{[]interface{}{0, 1}, []interface{}{1, 0}}}, Expected: []interface{}{}, Hidden: true, Note: "two nodes pointing at each other: a cycle"},
		},
		TimeComplexity:  "O(n + m)",
		SpaceComplexity: "O(n + m)",
		StarterCode: `def topologicalSort(n, edges):
    # Write your solution here
    pass`,
		Harness: Harness{
			Validator: `def check(raw, args, result, expected):
    n = raw['n']
    if expected == []:
        if result == []:
            return True
        return (False, "the requirements contain a cycle, so no order exists and the answer is []")
    if not isinstance(result, list) or sorted(result) != list(range(n)):
        return (False, "the answer must list every node from 0 to " + str(n - 1) + " exactly once, got " + str(result))
    pos = {node: i for i, node in enumerate(result)}
    for u, v in raw['edges']:
        if pos[u] > pos[v]:
            return (False, "edge " + str([u, v]) + " goes backward: " + str(u) + " must come before " + str(v) + " in " + str(result))
    return True`,
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: `Which tasks can safely go first, and what changes for the other tasks once one of them is done?`},
			{Level: 2, Type: "insight", Content: `Count each node's in-degree, the number of unfinished prerequisites. Any node with in-degree 0 can go now. Do it, then lower the in-degree of everything it points to, and queue those that reach 0. This is Kahn's algorithm from Chapter 13. If you finish and have placed fewer than n nodes, the rest sit on a cycle and never reach 0.`, ChapterRef: 13},
			{Level: 3, Type: "plan", Content: `1. Build adjacency lists and count each node's in-degree (the number of edges pointing to it).
2. Queue every node with in-degree 0.
3. Pop a node and append it to the order.
4. For each node it points to, lower that node's in-degree; if it becomes 0, queue it.
5. If the order holds all n nodes return it, otherwise return [].`},
			{Level: 4, Type: "skeleton", Content: `from collections import deque


def topologicalSort(n, edges):
    adj = [[] for _ in range(n)]
    indegree = [0] * n
    for u, v in edges:
        adj[u].append(v)
        indegree[v] += 1
    queue = deque(i for i in range(n) if ____)
    order = []
    while queue:
        node = queue.popleft()
        order.append(node)
        for nxt in adj[node]:
            ____
            if indegree[nxt] == 0:
                ____
    return order if ____ else []`},
		},
		Solution: Solution{
			Code: `from collections import deque


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
    return order if len(order) == n else []`,
			Explanation: "Kahn's algorithm repeatedly places a task that has no unfinished prerequisites, then 'completes' it by removing its edges from the counts of the tasks that depend on it. Tasks that become free join the queue. If the queue empties before every task is placed, the remaining tasks depend on each other in a cycle.",
			KeyInsight:  "A node with no incoming edges can always go first, and removing it can only free other nodes, never block them. In a cycle, every node has an incoming edge from another cycle node, so none of them ever reaches in-degree 0, and the count of placed nodes falls short of n.",
			WhyItWorks:  "Invariant: `indegree[x]` is the number of x's prerequisites not yet placed in `order`, and a node is queued exactly when that reaches 0. A node is appended only after all of its prerequisites were appended earlier, so every edge goes forward. If the order ends with n nodes, it is a valid order. If it ends short, the unplaced nodes each still have an unplaced prerequisite; following those backwards must eventually repeat a node, so there is a cycle, and no valid order exists. Each node is queued once and each edge is processed once: O(n + m).",
			CommonMistakes: []string{
				"Returning whatever order you collected without comparing its length to n. On a graph with a cycle that returns a partial order, instead of [].",
				"Counting in-degree from the wrong end of the edge. For [u, v] it is `v` that gains a prerequisite, so increase `indegree[v]`, not `indegree[u]`.",
				"Treating the edges as undirected. Direction is the entire point: [1, 0] means 1 goes before 0.",
				"Forgetting that nodes with no edges at all still belong in the answer. They start with in-degree 0 and are queued at the start.",
			},
			TimeComplexity:  "O(n + m): each node is queued once and each edge's counter is lowered once",
			SpaceComplexity: "O(n + m) for the lists, the counts and the queue",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the graph and count prerequisites", Explanation: "Each edge [u, v] gives u a neighbour v and gives v one more prerequisite.", CodeSnippet: `adj = [[] for _ in range(n)]
indegree = [0] * n
for u, v in edges:
    adj[u].append(v)
    indegree[v] += 1`, LineStart: 5, LineEnd: 9},
				{Title: "Queue every node that is ready", Explanation: "A node with in-degree 0 has no prerequisites, so it can go first. Isolated nodes are in this group.", CodeSnippet: `queue = deque(i for i in range(n) if indegree[i] == 0)
order = []`, LineStart: 10, LineEnd: 11},
				{Title: "Place the next ready node", Explanation: "Popping it and appending it fixes its position in the answer.", CodeSnippet: `node = queue.popleft()
order.append(node)`, LineStart: 13, LineEnd: 14},
				{Title: "Free the nodes that were waiting for it", Explanation: "Placing `node` satisfies one prerequisite of each neighbour. A neighbour whose count reaches 0 has nothing left to wait for and is queued.", CodeSnippet: `for nxt in adj[node]:
    indegree[nxt] -= 1
    if indegree[nxt] == 0:
        queue.append(nxt)`, LineStart: 15, LineEnd: 18},
				{Title: "Detect a cycle", Explanation: "If every node was placed there is a valid order. If some were not, they are stuck on a cycle, and the answer is [].", CodeSnippet: `return order if len(order) == n else []`, LineStart: 19, LineEnd: 19},
			},
			Alternatives: []AltSolution{
				{
					Name: "DFS with three colours (reverse postorder)",
					Code: `def topologicalSort(n, edges):
    adj = [[] for _ in range(n)]
    for u, v in edges:
        adj[u].append(v)
    state = [0] * n  # 0 = new, 1 = on the current path, 2 = finished
    order = []

    def visit(node):
        state[node] = 1
        for nxt in adj[node]:
            if state[nxt] == 1:
                return False
            if state[nxt] == 0 and not visit(nxt):
                return False
        state[node] = 2
        order.append(node)
        return True

    for i in range(n):
        if state[i] == 0 and not visit(i):
            return []
    return order[::-1]`,
					TimeComplexity:  "O(n + m)",
					SpaceComplexity: "O(n + m)",
					Explanation:     "Depth-first search: a node is added to the list after everything it points to is finished, so the reversed list has each node before everything it points to. Meeting a node that is still on the current path means a cycle. Uses recursion, so very deep graphs need the iterative Kahn version.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does a cycle make it impossible to place every node?", Answer: "Every node on a cycle has a prerequisite that is also on the cycle, and that one is waiting for the next, all the way round. So none of them can ever have all prerequisites done: none reaches in-degree 0, none is queued, and the order stops short of n nodes."},
			{Prompt: "Why can the answer be different from the expected one and still be correct?", Answer: "A topological order is only required to put every edge forward. Whenever two tasks have no path between them, either may go first, so many orders are valid. The judge here checks the edge rule rather than comparing with one answer."},
		},
	},
	{
		ID:              "drill-union-find",
		Kind:            KindDrill,
		Order:           9,
		Title:           "Union-Find (Disjoint Sets)",
		Slug:            "drill-union-find",
		Difficulty:      "Medium",
		Category:        "graphs",
		Tags:            []string{"Graph", "Union Find", "Design", "Foundations"},
		RelatedChapters: []int{13},
		Description: `Build a **union-find** structure (also called disjoint sets) that tracks which elements belong together. At the start there are ` + "`" + `n` + "`" + ` elements labelled ` + "`" + `0` + "`" + ` to ` + "`" + `n-1` + "`" + `, each alone in its own set.

Implement ` + "`" + `class UnionFind` + "`" + ` with:

- ` + "`" + `__init__(self, n)` + "`" + `: start with ` + "`" + `n` + "`" + ` separate sets.
- ` + "`" + `find(self, x)` + "`" + `: return the **representative** (root) of the set containing ` + "`" + `x` + "`" + `. Two elements are in the same set exactly when ` + "`" + `find` + "`" + ` gives them the same answer.
- ` + "`" + `union(self, a, b)` + "`" + `: merge the sets containing ` + "`" + `a` + "`" + ` and ` + "`" + `b` + "`" + `. Return ` + "`" + `True` + "`" + ` if they were in different sets and have been merged, ` + "`" + `False` + "`" + ` if they were already in the same set.
- ` + "`" + `connected(self, a, b)` + "`" + `: return ` + "`" + `True` + "`" + ` if ` + "`" + `a` + "`" + ` and ` + "`" + `b` + "`" + ` are in the same set.
- ` + "`" + `count(self)` + "`" + `: return the current number of separate sets.

Which element is chosen as a set's representative is up to you, so the tests call only ` + "`" + `union` + "`" + `, ` + "`" + `connected` + "`" + ` and ` + "`" + `count` + "`" + `. Make it fast: with path compression and union by size, every operation is close to O(1).`,
		Constraints: []string{
			"1 <= n <= 30000",
			"0 <= a, b, x < n",
			"Up to 60000 calls to union, connected and count in total",
		},
		Examples: []Example{
			{Input: "[\"UnionFind\",\"union\",\"union\",\"connected\",\"connected\",\"count\",\"union\",\"count\"]\n[[5],[0,1],[1,2],[0,2],[0,3],[],[0,2],[]]", Output: "[null,true,true,true,false,3,false,3]", Explanation: "Five separate sets. union(0,1) and union(1,2) merge them, leaving {0,1,2}, {3}, {4}: 3 sets. 0 and 2 are connected, 0 and 3 are not. union(0,2) finds them already together, so it returns false and the count stays 3."},
			{Input: "[\"UnionFind\",\"count\",\"union\",\"union\",\"connected\",\"union\",\"connected\",\"count\"]\n[[4],[],[0,1],[2,3],[1,2],[1,2],[0,3],[]]", Output: "[null,4,true,true,false,true,true,1]", Explanation: "{0,1} and {2,3} are separate, so 1 and 2 are not connected, until union(1,2) joins them. Then 0 and 3 are connected and there is 1 set."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"UnionFind", "union", "union", "connected", "connected", "count", "union", "count"}, "values": [][]interface{}{{5}, {0, 1}, {1, 2}, {0, 2}, {0, 3}, {}, {0, 2}, {}}}, Expected: []interface{}{nil, true, true, true, false, 3, false, 3}, Note: "example 1: joining sets, a connected query, and a union that changes nothing"},
			{Input: map[string]interface{}{"operations": []string{"UnionFind", "count", "union", "union", "connected", "union", "connected", "count"}, "values": [][]interface{}{{4}, {}, {0, 1}, {2, 3}, {1, 2}, {1, 2}, {0, 3}, {}}}, Expected: []interface{}{nil, 4, true, true, false, true, true, 1}, Note: "example 2: two separate pairs are not connected until a union links them"},
			{Input: map[string]interface{}{"operations": []string{"UnionFind", "count", "connected", "union"}, "values": [][]interface{}{{1}, {}, {0, 0}, {0, 0}}}, Expected: []interface{}{nil, 1, true, false}, Hidden: true, Note: "one element: it is connected to itself, and union(0, 0) merges nothing"},
			{Input: map[string]interface{}{"operations": []string{"UnionFind", "union", "union", "union", "count"}, "values": [][]interface{}{{4}, {3, 2}, {2, 3}, {3, 2}, {}}}, Expected: []interface{}{nil, true, false, false, 3}, Hidden: true, Note: "the same pair joined again, in either order, is False and the count stays put"},
			{Input: map[string]interface{}{"operations": []string{"UnionFind", "union", "union", "union", "union", "union", "connected", "count"}, "values": [][]interface{}{{6}, {0, 1}, {2, 3}, {4, 5}, {0, 2}, {0, 4}, {1, 5}, {}}}, Expected: []interface{}{nil, true, true, true, true, true, true, 1}, Hidden: true, Note: "merging bigger sets into bigger sets: everything ends up in one set"},
			{Input: map[string]interface{}{"operations": []string{"UnionFind", "connected", "connected", "count"}, "values": [][]interface{}{{3}, {0, 1}, {1, 2}, {}}}, Expected: []interface{}{nil, false, false, 3}, Hidden: true, Note: "nothing joined yet: no two different elements are connected"},
			{Input: map[string]interface{}{"operations": []string{"UnionFind", "union", "union", "union", "union", "union", "connected", "count"}, "values": [][]interface{}{{6}, {0, 1}, {2, 3}, {1, 3}, {0, 2}, {4, 5}, {2, 4}, {}}}, Expected: []interface{}{nil, true, true, true, false, true, false, 2}, Hidden: true, Note: "a union between two members of the same set (0 and 2 are already linked through 1 and 3) is False"},
			{Input: map[string]interface{}{"operations": graphDrillUnionOps(25000), "values": graphDrillUnionValues(25000)}, Expected: graphDrillUnionExpected(25000), Hidden: true, Note: "25,000 unions in a row, then one query across the whole chain: a deep chain of parents must not be walked recursively"},
		},
		TimeComplexity:  "O(1) per operation",
		SpaceComplexity: "O(n)",
		StarterCode: `class UnionFind:
    def __init__(self, n):
        pass

    def find(self, x):
        pass

    def union(self, a, b):
        pass

    def connected(self, a, b):
        pass

    def count(self):
        pass`,
		Harness: Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: `If every element knows one other element in its own set, following those links can reach one special element for the whole set. What would that special element look like, and how do you merge two sets?`},
			{Level: 2, Type: "insight", Content: `Keep a ` + "`" + `parent` + "`" + ` list where every element starts as its own parent. ` + "`" + `find` + "`" + ` follows parents until it reaches an element that is its own parent: the root. ` + "`" + `union` + "`" + ` finds both roots and, if they differ, makes one root the parent of the other. Two refinements make it fast: attach the smaller set under the larger (union by size), and point every element you pass during ` + "`" + `find` + "`" + ` straight at the root (path compression). See union-find in Chapter 13.`, ChapterRef: 13},
			{Level: 3, Type: "plan", Content: `1. Start ` + "`" + `parent[i] = i` + "`" + `, ` + "`" + `size[i] = 1` + "`" + `, and ` + "`" + `sets = n` + "`" + `.
2. ` + "`" + `find(x)` + "`" + `: walk to the root, then walk again from ` + "`" + `x` + "`" + ` pointing each element's parent at the root; return the root.
3. ` + "`" + `union(a, b)` + "`" + `: find both roots; if they are the same return False.
4. Otherwise attach the root of the smaller set under the larger, add the sizes, decrease ` + "`" + `sets` + "`" + ` by one, and return True.
5. ` + "`" + `connected` + "`" + ` compares two roots; ` + "`" + `count` + "`" + ` returns ` + "`" + `sets` + "`" + `.`},
			{Level: 4, Type: "skeleton", Content: `class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.size = [1] * n
        self.sets = n

    def find(self, x):
        root = x
        while self.parent[root] != root:
            root = ____
        while self.parent[x] != root:
            self.parent[x], x = ____, self.parent[x]
        return root

    def union(self, a, b):
        ra, rb = self.find(a), self.find(b)
        if ____:
            return False
        if self.size[ra] < self.size[rb]:
            ra, rb = rb, ra
        self.parent[rb] = ra
        ____
        self.sets -= 1
        return True

    def connected(self, a, b):
        return ____

    def count(self):
        return self.sets`},
		},
		Solution: Solution{
			Code: `class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.size = [1] * n
        self.sets = n

    def find(self, x):
        root = x
        while self.parent[root] != root:
            root = self.parent[root]
        while self.parent[x] != root:
            self.parent[x], x = root, self.parent[x]
        return root

    def union(self, a, b):
        ra, rb = self.find(a), self.find(b)
        if ra == rb:
            return False
        if self.size[ra] < self.size[rb]:
            ra, rb = rb, ra
        self.parent[rb] = ra
        self.size[ra] += self.size[rb]
        self.sets -= 1
        return True

    def connected(self, a, b):
        return self.find(a) == self.find(b)

    def count(self):
        return self.sets`,
			Explanation: "Each set is a tree of parent links whose root is the set's representative. `find` climbs to the root, `union` hangs one root under the other, and two elements are connected exactly when they have the same root. Path compression and union by size keep the trees almost flat.",
			KeyInsight:  "Whether two elements are in the same set is just whether `find` gives them the same root, so merging two sets costs one link between two roots. Keeping trees short (union by size) and flattening them as you search (path compression) make each operation effectively constant.",
			WhyItWorks:  "Invariant: following parent links from any element ends at a root, and the elements that reach a given root are exactly one set. Initially each element is its own root. `union` changes parent links only for one root, hanging it under another root, which merges exactly those two sets and nothing else. Path compression only moves an element's link to a higher ancestor, which is the same root, so it never changes the sets. `union` returns True and lowers the count exactly when the roots differ, which is when two different sets merge. Attaching the smaller tree under the larger means an element's depth grows only when its set at least doubles, so depth is at most log n, and with compression the cost per operation is nearly constant (the inverse Ackermann function).",
			CommonMistakes: []string{
				"Comparing the elements instead of their roots in `union` or `connected`. `connected(0, 3)` must compare `find(0)` and `find(3)`; two elements in the same set usually have different parents.",
				"Returning True from `union` even when the elements were already together, or lowering the count then. Both must only happen when two different roots are merged.",
				"Linking the elements, `parent[a] = b`, instead of their roots. That cuts a away from its own set and silently splits it.",
				"Writing `find` recursively without thinking about depth. Joining a long chain by naive unions produces a path 25,000 deep, which overflows Python's recursion limit; an iterative find with union by size avoids both problems.",
			},
			TimeComplexity:  "O(1) per operation: close to constant with path compression and union by size (inverse Ackermann)",
			SpaceComplexity: "O(n) for the parent, size and count",
			Walkthrough: []WalkthroughStep{
				{Title: "Start with every element alone", Explanation: "Each element is its own parent (so its own root) with size 1, and there are n sets.", CodeSnippet: `self.parent = list(range(n))
self.size = [1] * n
self.sets = n`, LineStart: 3, LineEnd: 5},
				{Title: "Climb to the root", Explanation: "A root is an element that is its own parent. Following parent links from x always ends there.", CodeSnippet: `root = x
while self.parent[root] != root:
    root = self.parent[root]`, LineStart: 8, LineEnd: 10},
				{Title: "Compress the path", Explanation: "A second walk points every element on the path directly at the root, so the next find from any of them takes one step.", CodeSnippet: `while self.parent[x] != root:
    self.parent[x], x = root, self.parent[x]
return root`, LineStart: 11, LineEnd: 13},
				{Title: "Stop if both are already together", Explanation: "Equal roots mean one set: nothing to merge, so return False and leave the count alone.", CodeSnippet: `ra, rb = self.find(a), self.find(b)
if ra == rb:
    return False`, LineStart: 16, LineEnd: 18},
				{Title: "Hang the smaller tree under the larger", Explanation: "Swapping makes `ra` the bigger root. Attaching `rb` under it keeps trees short, and the sizes and set count are updated.", CodeSnippet: `if self.size[ra] < self.size[rb]:
    ra, rb = rb, ra
self.parent[rb] = ra
self.size[ra] += self.size[rb]
self.sets -= 1
return True`, LineStart: 19, LineEnd: 24},
				{Title: "Answer connected and count", Explanation: "Connected means same root. The count is a number we maintain, so there is no need to scan anything.", CodeSnippet: `return self.find(a) == self.find(b)`, LineStart: 27, LineEnd: 27},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does union by size keep the trees shallow?", Answer: "An element's depth increases by one only when its root is hung under a bigger root, which means its set at least doubled in size. A set can double at most log2(n) times before it holds all n elements, so no element is ever deeper than log2(n). Without it, unions like (0,1), (1,2), ... can build a chain n deep."},
			{Prompt: "What does path compression do, and does it change which elements are in the same set?", Answer: "During `find`, it re-points every element on the way up at the root, so later finds from those elements take one step. It only replaces an element's parent with one of its own ancestors, which leads to the same root, so the sets are unchanged; only the shape of the trees gets flatter."},
		},
	},
}

// graphDrillPathEdges is the edge list of the path graph 0-1-2-...-(n-1).
func graphDrillPathEdges(n int) []interface{} {
	out := make([]interface{}, 0, n-1)
	for i := 0; i+1 < n; i++ {
		out = append(out, []interface{}{i, i + 1})
	}
	return out
}

// graphDrillCountUp is [0, 1, ..., n-1], the DFS order of the path graph from node 0.
func graphDrillCountUp(n int) []interface{} {
	out := make([]interface{}, n)
	for i := range out {
		out[i] = i
	}
	return out
}

// graphDrillUnionOps is the operation list for n-1 unions in a row
// (union(0,1), union(1,2), ...), then one connected query and one count.
func graphDrillUnionOps(n int) []string {
	ops := make([]string, 0, n+2)
	ops = append(ops, "UnionFind")
	for i := 0; i+1 < n; i++ {
		ops = append(ops, "union")
	}
	return append(ops, "connected", "count")
}

// graphDrillUnionValues holds the arguments that go with graphDrillUnionOps.
func graphDrillUnionValues(n int) [][]interface{} {
	vals := make([][]interface{}, 0, n+2)
	vals = append(vals, []interface{}{n})
	for i := 0; i+1 < n; i++ {
		vals = append(vals, []interface{}{i, i + 1})
	}
	return append(vals, []interface{}{0, n - 1}, []interface{}{})
}

// graphDrillUnionExpected is the answer list for graphDrillUnionOps: every
// union merges two sets, the ends are connected, and one set is left.
func graphDrillUnionExpected(n int) []interface{} {
	out := make([]interface{}, 0, n+2)
	out = append(out, nil)
	for i := 0; i+1 < n; i++ {
		out = append(out, true)
	}
	return append(out, true, 1)
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, GraphsDrills...)
}
