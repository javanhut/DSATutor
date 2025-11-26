package chapter

import "time"

// DefaultChapters returns reusable scaffolding for chapters 1–11.
// Each chapter includes placeholders for animations, visualizers, tutorials,
// and exercises so you can drop in concrete assets later.
func DefaultChapters() []Chapter {
	return []Chapter{
		{
			Number: 1,
			Slug:   "introduction-to-algorithms",
			Title:  "Introduction to Algorithms",
			Summary: "Warm up with binary search, storyboards for Big O intuition, " +
				"and a gentle on-ramp to recursion.",
			Objectives: []string{
				"Explain binary search in your own words",
				"Classify algorithm growth using Big O notation",
				"Recognize base vs. recursive cases at a glance",
			},
			Concepts: []Concept{
				{
					Name:        "Binary Search",
					Description: "Narrow a sorted list by halving.",
					CoreIdeas:   []string{"Requires sorted data", "Logarithmic steps"},
					Examples: []CodeExample{
						{
							ID:       "binary_search",
							Title:    "Binary search in Python",
							Language: "python",
							Snippet: `def binary_search(nums, target):
    low, high = 0, len(nums) - 1
    while low <= high:
        mid = (low + high) // 2
        guess = nums[mid]
        if guess == target:
            return mid
        if guess < target:
            low = mid + 1
        else:
            high = mid - 1
    return None`,
							Notes: "Requires sorted input; returns index or None if missing.",
						},
					},
				},
				{
					Name:        "Big O Notation",
					Description: "Describe how work scales with input size.",
					CoreIdeas:   []string{"Worst-case framing", "Compare shapes, not constants"},
					Examples: []CodeExample{
						{
							ID:       "big_o_loop",
							Title:    "Loop cost intuition",
							Language: "python",
							Snippet: `# O(n): touches each element once
for item in items:
    pass`,
							Notes: "Use Big O to compare growth, not exact timings.",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "binary-search",
					Title: "Binary Search Walkthrough",
					Goal:  "Show halving a sorted array until the target is found or missed.",
					Steps: []StoryboardStep{
						{Cue: "sorted-lineup", Narration: "Start with a sorted list of numbers.", VisualHint: "Array laid out left to right", Duration: Duration(2 * time.Second), CodeRef: "binary_search: init low/high"},
						{Cue: "pick-middle", Narration: "Inspect the middle item; compare against the target.", VisualHint: "Highlight mid index", Duration: Duration(3 * time.Second), CodeRef: "binary_search: mid calc"},
						{Cue: "discard-half", Narration: "Discard the half that cannot contain the target.", VisualHint: "Fade out irrelevant section", Duration: Duration(3 * time.Second), CodeRef: "binary_search: adjust bounds"},
						{Cue: "repeat", Narration: "Repeat until found or the interval is empty.", VisualHint: "Loop indicator", Duration: Duration(3 * time.Second), CodeRef: "binary_search: while loop"},
					},
				},
				{
					ID:    "runtime-shapes",
					Title: "Runtime Shapes",
					Goal:  "Visualize O(1), O(log n), O(n), and O(n log n) growth.",
					Steps: []StoryboardStep{
						{Cue: "plot-axes", Narration: "Plot n on the x-axis and operations on the y-axis.", VisualHint: "Blank axes with labels", Duration: Duration(2 * time.Second), CodeRef: "runtime_shapes: setup plot"},
						{Cue: "draw-log", Narration: "Sketch O(log n) rising slowly.", VisualHint: "Gentle curve", Duration: Duration(3 * time.Second), CodeRef: "runtime_shapes: log curve"},
						{Cue: "draw-linear", Narration: "Sketch O(n) as a straight line.", VisualHint: "Line through origin", Duration: Duration(3 * time.Second), CodeRef: "runtime_shapes: linear"},
						{Cue: "draw-nlogn", Narration: "Sketch O(n log n) just above linear.", VisualHint: "Slightly steeper curve", Duration: Duration(3 * time.Second), CodeRef: "runtime_shapes: nlogn"},
						{Cue: "compare", Narration: "Highlight how shapes diverge as n grows.", VisualHint: "Callouts on large n region", Duration: Duration(3 * time.Second), CodeRef: "runtime_shapes: compare"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "big-o-sense",
					Title:   "Big O Speed Intuition",
					Outcome: "Map common operations to O(1), O(log n), O(n), and O(n log n).",
					Steps: []TutorialStep{
						{Prompt: "Classify checking an element in a hash map.", Guidance: "Hash maps average to O(1).", CodeFocus: "lookup", Checkpoint: "Hash access in Python is O(1)."},
						{Prompt: "Classify scanning for a value in an unsorted list.", Guidance: "You may need every item.", CodeFocus: "for loops", Checkpoint: "Linear work is O(n)."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "bs-mental-model", Prompt: "Walk through binary search on [1,3,4,6,8,9] looking for 7.", Difficulty: "Easy", Outcome: "Understands failure case behavior."},
				{ID: "big-o-sorting", Prompt: "Explain why selection sort is O(n^2).", Difficulty: "Easy", Outcome: "Connects nested loops to quadratic growth."},
			},
			Visualizers: []Visualizer{
				{ID: "timeline-big-o", Title: "Runtime Shapes", Goal: "Compare n vs. log n vs. n log n growth visually.", DataModel: "series[n] -> cost", Interactions: []string{"scrub-n", "toggle-curve"}, Hooks: []string{"onStep", "onReset", "onCurveSelect"}},
			},
		},
		{
			Number: 2,
			Slug:   "selection-sort",
			Title:  "Selection Sort, Arrays, and Linked Lists",
			Summary: "Introduce arrays vs. linked lists and animate selection sort " +
				"as a first pass at reasoning about O(n^2) work.",
			Objectives: []string{
				"Contrast indexed arrays with pointer-based linked lists",
				"Trace selection sort swaps and comparisons",
				"Relate data structure choice to insertion and lookup costs",
			},
			Concepts: []Concept{
				{
					Name:        "Arrays",
					Description: "Contiguous memory with O(1) index lookup.",
					CoreIdeas:   []string{"Fast reads", "Slow middle insertions"},
					Examples: []CodeExample{
						{
							ID:       "array_index",
							Title:    "Array indexing",
							Language: "python",
							Snippet: `vals = [2, 4, 6]
second = vals[1]  # O(1) read`,
							Notes: "Random access is constant time; middle inserts cost shifting.",
						},
					},
				},
				{
					Name:        "Linked Lists",
					Description: "Nodes linked by pointers.",
					CoreIdeas:   []string{"Fast inserts/removals", "Sequential reads"},
					Examples: []CodeExample{
						{
							ID:       "linked_list_insert",
							Title:    "Append to a singly linked list node",
							Language: "python",
							Snippet: `class Node:
    def __init__(self, val, next=None):
        self.val = val
        self.next = next

def insert_after(node, val):
    node.next = Node(val, node.next)`,
							Notes: "Insertion avoids shifting elements; traversal stays O(n).",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "selection-sort",
					Title: "Selection Sort Mechanics",
					Goal:  "Show selecting the minimum and swapping forward.",
					Steps: []StoryboardStep{
						{Cue: "scan", Narration: "Scan the unsorted portion for the minimum.", VisualHint: "Sweeping pointer", Duration: Duration(3 * time.Second), CodeRef: "selection_sort: inner scan"},
						{Cue: "select", Narration: "Mark the smallest element found.", VisualHint: "Highlight min", Duration: Duration(2 * time.Second), CodeRef: "selection_sort: track min index"},
						{Cue: "swap", Narration: "Swap it with the first unsorted position.", VisualHint: "Animate swap", Duration: Duration(3 * time.Second), CodeRef: "selection_sort: swap"},
						{Cue: "shrink", Narration: "Shrink the unsorted window and repeat.", VisualHint: "Move boundary", Duration: Duration(2 * time.Second), CodeRef: "selection_sort: outer loop"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "array-vs-list",
					Title:   "Choosing Arrays or Linked Lists",
					Outcome: "Pick the right structure based on required operations.",
					Steps: []TutorialStep{
						{Prompt: "Frequent middle insertions?", Guidance: "Prefer linked lists.", CodeFocus: "Insert operations", Checkpoint: "Linked lists avoid shifting elements."},
						{Prompt: "Need random access by index?", Guidance: "Use arrays.", CodeFocus: "Indexing", Checkpoint: "Arrays give O(1) access."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "count-swaps", Prompt: "How many swaps does selection sort make for n items?", Difficulty: "Easy", Outcome: "Understands swap count is n-1."},
			},
			Visualizers: []Visualizer{
				{ID: "array-list-comparer", Title: "Array vs. Linked List", Goal: "Play insertion and access operations against both structures.", DataModel: "operation sequence", Interactions: []string{"step", "reset"}, Hooks: []string{"onStep", "onReset", "onOperation"}},
			},
		},
		{
			Number:  3,
			Slug:    "recursion",
			Title:   "Recursion and the Call Stack",
			Summary: "Decompose problems into base and recursive cases and visualize the call stack.",
			Objectives: []string{
				"Identify base vs. recursive cases in code",
				"Trace call stacks and unwinding",
				"Recognize when recursion is a good fit",
			},
			Concepts: []Concept{
				{
					Name:        "Base Case",
					Description: "The simplest input handled directly.",
					CoreIdeas:   []string{"Stops recursion", "Often size 0 or 1"},
					Examples: []CodeExample{
						{
							ID:       "factorial",
							Title:    "Base case for factorial",
							Language: "python",
							Snippet: `def fact(n):
    if n <= 1:  # base case
        return 1
    return n * fact(n - 1)`,
							Notes: "Base case must be reachable to avoid infinite recursion.",
						},
					},
				},
				{
					Name:        "Recursive Case",
					Description: "Reduces the problem size and calls itself.",
					CoreIdeas:   []string{"Moves toward base case", "Accumulates work"},
					Examples: []CodeExample{
						{
							ID:       "recursive_sum",
							Title:    "Recursive sum",
							Language: "python",
							Snippet: `def sum_list(nums):
    if len(nums) == 0:  # base
        return 0
    return nums[0] + sum_list(nums[1:])  # recursive case`,
							Notes: "Each call shrinks the input, ensuring termination.",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "call-stack",
					Title: "Recursive Call Stack",
					Goal:  "Show stack frames building up and collapsing.",
					Steps: []StoryboardStep{
						{Cue: "push", Narration: "Each recursive call pushes a new frame.", VisualHint: "Stack grows downward", Duration: Duration(3 * time.Second), CodeRef: "recursion: call self"},
						{Cue: "base-case", Narration: "At the base case, no further calls are made.", VisualHint: "Highlight terminal frame", Duration: Duration(3 * time.Second), CodeRef: "recursion: base check"},
						{Cue: "pop", Narration: "Frames pop as results return.", VisualHint: "Stack shrinks upward", Duration: Duration(3 * time.Second), CodeRef: "recursion: return unwind"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "sum-recursive",
					Title:   "Recursive Summation",
					Outcome: "Write a recursive function to sum a list.",
					Steps: []TutorialStep{
						{Prompt: "Define the base case.", Guidance: "Empty slice returns 0.", CodeFocus: "if len(nums)==0", Checkpoint: "Base case avoids infinite calls."},
						{Prompt: "Define the recursive step.", Guidance: "First element plus sum of the rest.", CodeFocus: "nums[0] + sum(nums[1:])", Checkpoint: "Problem size shrinks each call."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "recursion-proof", Prompt: "Explain why every recursive algorithm needs a base case.", Difficulty: "Easy", Outcome: "Connects base cases to termination."},
			},
			Visualizers: []Visualizer{
				{ID: "stack-tracer", Title: "Stack Tracer", Goal: "Step through recursive calls with a visible stack.", DataModel: "function frames", Interactions: []string{"step", "rewind", "auto"}, Hooks: []string{"onStep", "onReset", "onFrameChange"}},
			},
		},
		{
			Number:  4,
			Slug:    "quicksort",
			Title:   "Divide & Conquer with Quicksort",
			Summary: "Use divide-and-conquer to sort and compare average vs. worst cases.",
			Objectives: []string{
				"Walk through quicksort partitioning",
				"Contrast average O(n log n) vs. worst O(n^2)",
				"Choose pivots that avoid degeneration",
			},
			Concepts: []Concept{
				{
					Name:        "Divide & Conquer",
					Description: "Split a problem, solve parts, combine results.",
					CoreIdeas:   []string{"Divide", "Conquer", "Combine"},
					Examples: []CodeExample{
						{
							ID:       "quicksort",
							Title:    "Quicksort skeleton",
							Language: "python",
							Snippet: `def quicksort(nums):
    if len(nums) <= 1:
        return nums
    pivot = nums[len(nums) // 2]
    less  = [x for x in nums if x < pivot]
    equal = [x for x in nums if x == pivot]
    greater = [x for x in nums if x > pivot]
    return quicksort(less) + equal + quicksort(greater)`,
							Notes: "Divide on pivot, conquer sub-arrays, combine results.",
						},
					},
				},
				{
					Name:        "Pivot Choice",
					Description: "Picking a pivot influences performance.",
					CoreIdeas:   []string{"Randomized pivoting", "Median-of-three"},
					Examples: []CodeExample{
						{
							ID:       "pivot_choice",
							Title:    "Randomized pivot selection",
							Language: "python",
							Snippet: `import random
pivot = nums[random.randrange(len(nums))]`,
							Notes: "Random pivots help avoid consistently bad partitions.",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "quicksort",
					Title: "Quicksort Partitioning",
					Goal:  "Show pivot selection and partition steps.",
					Steps: []StoryboardStep{
						{Cue: "choose-pivot", Narration: "Pick a pivot element.", VisualHint: "Highlight pivot", Duration: Duration(2 * time.Second), CodeRef: "quicksort: pivot"},
						{Cue: "partition", Narration: "Reorder so smaller items are left of the pivot.", VisualHint: "Two-pointer sweep", Duration: Duration(4 * time.Second), CodeRef: "quicksort: partition lists"},
						{Cue: "recurse", Narration: "Recurse on each sub-array.", VisualHint: "Call tree view", Duration: Duration(3 * time.Second), CodeRef: "quicksort: recursive calls"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "quicksort-annotate",
					Title:   "Implement Quicksort",
					Outcome: "Write a simple quicksort with clear base and recursive cases.",
					Steps: []TutorialStep{
						{Prompt: "Base case size?", Guidance: "Return slice if len <= 1.", CodeFocus: "len(nums) <= 1", Checkpoint: "No more splitting."},
						{Prompt: "Partition logic?", Guidance: "Split into less/equal/greater around pivot.", CodeFocus: "loops + appends", Checkpoint: "All items processed once."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "pivot-risk", Prompt: "What happens if the array is already sorted and you pick the first element as pivot?", Difficulty: "Medium", Outcome: "Understands worst-case behavior."},
			},
			Visualizers: []Visualizer{
				{ID: "pivot-simulator", Title: "Pivot Simulator", Goal: "Try different pivot strategies and watch depth.", DataModel: "array + pivot function", Interactions: []string{"randomize", "step"}, Hooks: []string{"onStep", "onReset", "onPartition"}},
			},
		},
		{
			Number:  5,
			Slug:    "hash-tables",
			Title:   "Hash Tables",
			Summary: "Cover hashing, collision handling, and performance trade-offs.",
			Objectives: []string{
				"Describe hash functions and buckets",
				"Choose collision strategies",
				"Monitor load factor and resizing",
			},
			Concepts: []Concept{
				{
					Name:        "Hash Function",
					Description: "Map keys to bucket indexes.",
					CoreIdeas:   []string{"Uniform distribution", "Deterministic"},
					Examples: []CodeExample{
						{
							ID:       "hash_lookup",
							Title:    "Go map lookup",
							Language: "python",
							Snippet: `emails = {"ada": "ada@example.com"}
_ = emails["ada"]  # hash + bucket lookup`,
							Notes: "Python dicts handle hashing internally; collisions are abstracted away.",
						},
					},
				},
				{
					Name:        "Collisions",
					Description: "Two keys can land in the same bucket.",
					CoreIdeas:   []string{"Chaining", "Open addressing"},
					Examples: []CodeExample{
						{
							ID:       "hash_chaining",
							Title:    "Chaining concept",
							Language: "python",
							Snippet: `class Entry:
    def __init__(self, key, val, next=None):
        self.key = key
        self.val = val
        self.next = next
# Bucket holds a linked list of entries sharing a hash.`,
							Notes: "Chaining keeps multiple entries per bucket via linked lists.",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "hash-insert",
					Title: "Hash Insert and Lookup",
					Goal:  "Show hashing, collision handling, and lookup.",
					Steps: []StoryboardStep{
						{Cue: "hash", Narration: "Hash the key to find a bucket.", VisualHint: "Key -> hash -> bucket index", Duration: Duration(3 * time.Second), CodeRef: "hash_table: hash(key) -> bucket"},
						{Cue: "collision", Narration: "Handle collisions gracefully.", VisualHint: "Linked list or probe sequence", Duration: Duration(3 * time.Second), CodeRef: "hash_table: collision handling"},
						{Cue: "lookup", Narration: "Repeat hashing for fast lookups.", VisualHint: "O(1) pointer jump", Duration: Duration(2 * time.Second), CodeRef: "hash_table: lookup"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "design-hash-table",
					Title:   "Design a Simple Hash Table",
					Outcome: "Outline buckets, hashing, and resizing steps.",
					Steps: []TutorialStep{
						{Prompt: "Pick bucket count.", Guidance: "Start with a small power of two.", CodeFocus: "make([]bucket, n)", Checkpoint: "Bucket array allocated."},
						{Prompt: "Define collision strategy.", Guidance: "Use chaining with slices of pairs.", CodeFocus: "append to bucket slice", Checkpoint: "Buckets store multiple entries."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "load-factor", Prompt: "Why do lookups slow down when the load factor grows?", Difficulty: "Medium", Outcome: "Connects load factor to collision frequency."},
			},
			Visualizers: []Visualizer{
				{ID: "bucket-viewer", Title: "Bucket Viewer", Goal: "Animate insertions and collisions.", DataModel: "key/value pairs", Interactions: []string{"insert", "delete", "search"}, Hooks: []string{"onStep", "onReset", "onResize"}},
			},
		},
		{
			Number:  6,
			Slug:    "breadth-first-search",
			Title:   "Breadth-First Search",
			Summary: "Model problems as graphs and use BFS to find shortest paths in unweighted graphs.",
			Objectives: []string{
				"Represent graphs with adjacency lists",
				"Use queues to explore neighbors level by level",
				"Detect shortest paths in unweighted graphs",
			},
			Concepts: []Concept{
				{
					Name:        "Graph",
					Description: "Nodes and edges to model relationships.",
					CoreIdeas:   []string{"Directed vs. undirected", "Weighted vs. unweighted"},
					Examples: []CodeExample{
						{
							ID:       "adjacency_list",
							Title:    "Adjacency list",
							Language: "python",
							Snippet: `graph = {
    "A": ["B", "C"],
    "B": ["D"],
    "C": ["D"],
}`,
							Notes: "Adjacency lists keep neighbors per node.",
						},
					},
				},
				{
					Name:        "Queue",
					Description: "FIFO structure driving BFS order.",
					CoreIdeas:   []string{"Enqueue neighbors", "Dequeue next frontier"},
					Examples: []CodeExample{
						{
							ID:       "bfs",
							Title:    "BFS over adjacency list",
							Language: "python",
							Snippet: `from collections import deque

def bfs(start, graph):
    visited = {start}
    queue = deque([start])
    order = []
    while queue:
        node = queue.popleft()
        order.append(node)
        for nbr in graph.get(node, []):
            if nbr not in visited:
                visited.add(nbr)
                queue.append(nbr)
    return order`,
							Notes: "FIFO queue ensures level-order traversal.",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "bfs",
					Title: "BFS Level Order",
					Goal:  "Show frontier expansion using a queue.",
					Steps: []StoryboardStep{
						{Cue: "enqueue-start", Narration: "Begin with the start node in the queue.", VisualHint: "Queue with start node", Duration: Duration(2 * time.Second), CodeRef: "bfs: init queue"},
						{Cue: "expand", Narration: "Dequeue and enqueue unseen neighbors.", VisualHint: "Highlight visited nodes", Duration: Duration(4 * time.Second), CodeRef: "bfs: neighbor loop"},
						{Cue: "goal", Narration: "Stop when the goal is dequeued.", VisualHint: "Path traced back", Duration: Duration(3 * time.Second), CodeRef: "bfs: goal check/parent map"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "shortest-path",
					Title:   "Build a BFS Shortest Path Finder",
					Outcome: "Trace predecessors to reconstruct a path.",
					Steps: []TutorialStep{
						{Prompt: "Track visited?", Guidance: "Use a map from node to bool.", CodeFocus: "visited map", Checkpoint: "No reprocessing nodes."},
						{Prompt: "Track parents?", Guidance: "Store predecessor per node.", CodeFocus: "parent map", Checkpoint: "Can reconstruct path."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "queue-why", Prompt: "Why does BFS need a queue instead of a stack?", Difficulty: "Easy", Outcome: "Understands level-order traversal."},
			},
			Visualizers: []Visualizer{
				{ID: "graph-builder", Title: "Graph Builder", Goal: "Draw nodes and edges, then animate BFS.", DataModel: "nodes + edges", Interactions: []string{"add-node", "add-edge", "run-bfs"}, Hooks: []string{"onStep", "onReset", "onVisit"}},
			},
		},
		{
			Number:  7,
			Slug:    "dijkstras-algorithm",
			Title:   "Dijkstra's Algorithm",
			Summary: "Find shortest weighted paths without negative edges using greedy relaxation.",
			Objectives: []string{
				"Maintain tentative distances and parents",
				"Pick the lowest-cost unprocessed node each iteration",
				"Detect when graphs break the no-negative-edges rule",
			},
			Concepts: []Concept{
				{
					Name:        "Relaxation",
					Description: "Improve path estimates when a shorter route is found.",
					CoreIdeas:   []string{"Distance updates", "Parent tracking"},
					Examples: []CodeExample{
						{
							ID:       "relaxation",
							Title:    "Edge relaxation",
							Language: "python",
							Snippet: `if dist[u] + w < dist[v]:
    dist[v] = dist[u] + w
    parent[v] = u`,
							Notes: "Only update when a shorter path is discovered.",
						},
					},
				},
				{
					Name:        "Priority Queue",
					Description: "Fetch next cheapest node efficiently.",
					CoreIdeas:   []string{"Min-heap behavior", "Avoid O(n^2) scans"},
					Examples: []CodeExample{
						{
							ID:       "min_heap",
							Title:    "Min-heap with container/heap",
							Language: "python",
							Snippet: `import heapq

heap = []
heapq.heappush(heap, (cost, node))
cost, node = heapq.heappop(heap)`,
							Notes: "A heap keeps pop-min at O(log n) instead of scanning.",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "dijkstra",
					Title: "Dijkstra in Action",
					Goal:  "Show cost updates and finalized nodes.",
					Steps: []StoryboardStep{
						{Cue: "init", Narration: "Initialize all node costs to infinity except the start.", VisualHint: "Table of costs", Duration: Duration(3 * time.Second), CodeRef: "dijkstra: init dist/parent"},
						{Cue: "pick-node", Narration: "Pick the lowest-cost unprocessed node.", VisualHint: "Highlight min-cost", Duration: Duration(3 * time.Second), CodeRef: "dijkstra: pop heap"},
						{Cue: "relax", Narration: "Relax edges: update neighbor costs if cheaper.", VisualHint: "Arrows with updated weights", Duration: Duration(4 * time.Second), CodeRef: "dijkstra: relax edge"},
						{Cue: "finalize", Narration: "Mark node as processed and repeat.", VisualHint: "Move node to done set", Duration: Duration(3 * time.Second), CodeRef: "dijkstra: visited set"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "implement-dijkstra",
					Title:   "Implement Dijkstra",
					Outcome: "Write the loop with a priority queue or simple scan.",
					Steps: []TutorialStep{
						{Prompt: "Detect negative edges.", Guidance: "Reject graphs with negatives.", CodeFocus: "edge weight check", Checkpoint: "Fails fast on invalid input."},
						{Prompt: "Update costs and parents.", Guidance: "Compare newCost < cost[neighbor].", CodeFocus: "relaxation", Checkpoint: "Parents set only on improvement."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "negative-edge", Prompt: "Why does a negative edge break Dijkstra?", Difficulty: "Medium", Outcome: "Understands greedy assumption fails."},
			},
			Visualizers: []Visualizer{
				{ID: "weighted-paths", Title: "Weighted Paths", Goal: "Animate cost table updates alongside the graph.", DataModel: "weighted graph", Interactions: []string{"add-edge", "run-dijkstra"}, Hooks: []string{"onStep", "onReset", "onRelax"}},
			},
		},
		{
			Number:  8,
			Slug:    "greedy-algorithms",
			Title:   "Greedy Algorithms",
			Summary: "Select locally optimal steps for near-optimal solutions; explore proofs and counterexamples.",
			Objectives: []string{
				"Model problems where greedy works",
				"Recognize when approximation is acceptable",
				"State the exchange argument intuition",
			},
			Concepts: []Concept{
				{
					Name:        "Greedy Choice",
					Description: "Pick the best local option and commit.",
					CoreIdeas:   []string{"No backtracking", "Works with certain properties"},
					Examples: []CodeExample{
						{
							ID:       "activity_selection",
							Title:    "Activity selection greedy",
							Language: "python",
							Snippet: `classes = sorted(classes, key=lambda c: c.end)
schedule = []
current_end = 0
for c in classes:
    if c.start >= current_end:
        schedule.append(c)
        current_end = c.end`,
							Notes: "Earliest-finish-time is a valid greedy choice here.",
						},
					},
				},
				{
					Name:        "Approximation",
					Description: "Close-to-optimal solutions for hard problems.",
					CoreIdeas:   []string{"Set cover", "Knapsack heuristics"},
					Examples: []CodeExample{
						{
							ID:       "set_cover",
							Title:    "Set cover greedy step",
							Language: "python",
							Snippet: `while uncovered:
    best = max(sets, key=lambda s: len(uncovered & s))
    solution.append(best)
    uncovered -= best`,
							Notes: "Greedy yields a logarithmic-factor approximation for set cover.",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "classroom-scheduling",
					Title: "Activity Selection",
					Goal:  "Show picking earliest finishing classes first.",
					Steps: []StoryboardStep{
						{Cue: "sort", Narration: "Sort intervals by finish time.", VisualHint: "Timeline sorted", Duration: Duration(3 * time.Second), CodeRef: "activity_selection: sort by end"},
						{Cue: "pick", Narration: "Pick the earliest finishing class compatible with current schedule.", VisualHint: "Greedy selection highlight", Duration: Duration(3 * time.Second), CodeRef: "activity_selection: choose interval"},
						{Cue: "repeat", Narration: "Repeat with the remaining intervals.", VisualHint: "Timeline shrinks", Duration: Duration(3 * time.Second), CodeRef: "activity_selection: loop"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "set-cover-approx",
					Title:   "Set Cover Approximation",
					Outcome: "Iteratively pick the set covering the most uncovered items.",
					Steps: []TutorialStep{
						{Prompt: "Pick the best set.", Guidance: "Maximize uncovered elements per cost.", CodeFocus: "greedy loop", Checkpoint: "Progress toward full cover."},
						{Prompt: "Stop condition.", Guidance: "When all elements are covered.", CodeFocus: "while len(uncovered)>0", Checkpoint: "Algorithm terminates."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "counterexample", Prompt: "Give a case where greedy fails for the 0/1 knapsack problem.", Difficulty: "Medium", Outcome: "Understands limitations of greedy."},
			},
			Visualizers: []Visualizer{
				{ID: "greedy-playground", Title: "Greedy Playground", Goal: "Try greedy heuristics on scheduling and set cover.", DataModel: "intervals/sets", Interactions: []string{"shuffle", "step"}, Hooks: []string{"onStep", "onReset", "onChoice"}},
			},
		},
		{
			Number:  9,
			Slug:    "dynamic-programming",
			Title:   "Dynamic Programming",
			Summary: "Fill tables to reuse subproblem results; practice with knapsack and longest-common-substring.",
			Objectives: []string{
				"Break problems into overlapping subproblems",
				"Fill DP grids row-by-row or column-by-column",
				"Trace decisions to reconstruct solutions",
			},
			Concepts: []Concept{
				{
					Name:        "Subproblem Overlap",
					Description: "Solutions depend on repeated smaller cases.",
					CoreIdeas:   []string{"Memoization", "Bottom-up"},
					Examples: []CodeExample{
						{
							ID:       "fib_memo",
							Title:    "Memoized Fibonacci",
							Language: "python",
							Snippet: `memo = {0: 0, 1: 1}
def fib(n):
    if n in memo:
        return memo[n]
    memo[n] = fib(n - 1) + fib(n - 2)
    return memo[n]`,
							Notes: "Memoization reuses overlapping subproblem results.",
						},
					},
				},
				{
					Name:        "State",
					Description: "Minimal data needed to describe a subproblem.",
					CoreIdeas:   []string{"Indices", "Capacity", "Suffix/Prefix"},
					Examples: []CodeExample{
						{
							ID:       "knapsack_dp",
							Title:    "Knapsack DP state",
							Language: "python",
							Snippet: `dp = [[0]*(capacity+1) for _ in range(n_items+1)]
# dp[i][w] = best value using first i items within weight w`,
							Notes: "State picks just enough data (items considered, capacity).",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "knapsack-grid",
					Title: "Knapsack Table Fill",
					Goal:  "Show building the 2D value table.",
					Steps: []StoryboardStep{
						{Cue: "init-row", Narration: "Initialize base cases with zero capacity or zero items.", VisualHint: "First row/column zeros", Duration: Duration(3 * time.Second), CodeRef: "knapsack: init dp"},
						{Cue: "fill", Narration: "Decide include vs. exclude for each cell.", VisualHint: "Highlight cell dependency", Duration: Duration(4 * time.Second), CodeRef: "knapsack: dp transition"},
						{Cue: "traceback", Narration: "Trace choices back to recover items.", VisualHint: "Arrows backward", Duration: Duration(3 * time.Second), CodeRef: "knapsack: traceback"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "lcs",
					Title:   "Longest Common Substring Grid",
					Outcome: "Populate a DP grid and read the best length.",
					Steps: []TutorialStep{
						{Prompt: "Define state.", Guidance: "dp[i][j] = longest suffix match ending at i,j.", CodeFocus: "2D slice", Checkpoint: "State minimal and sufficient."},
						{Prompt: "Transition.", Guidance: "If chars match, extend from dp[i-1][j-1], else 0.", CodeFocus: "if a[i-1]==b[j-1]", Checkpoint: "Reset on mismatch."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "order-matters", Prompt: "Does filling rows vs. columns change correctness?", Difficulty: "Easy", Outcome: "Understands independence of traversal order when dependencies are met."},
			},
			Visualizers: []Visualizer{
				{ID: "dp-grid", Title: "DP Grid Animator", Goal: "Animate cell dependencies and traceback.", DataModel: "2D table", Interactions: []string{"step", "auto", "reset"}, Hooks: []string{"onStep", "onReset", "onCellUpdate"}},
			},
		},
		{
			Number:  10,
			Slug:    "k-nearest-neighbors",
			Title:   "K-Nearest Neighbors",
			Summary: "Classify and regress using distance metrics and feature scaling.",
			Objectives: []string{
				"Compute distances between feature vectors",
				"Pick k and distance metrics thoughtfully",
				"Handle classification vs. regression outputs",
			},
			Concepts: []Concept{
				{
					Name:        "Distance Metrics",
					Description: "Measure closeness between points.",
					CoreIdeas:   []string{"Euclidean", "Manhattan", "Cosine"},
					Examples: []CodeExample{
						{
							ID:       "euclidean_distance",
							Title:    "Euclidean distance",
							Language: "python",
							Snippet: `import math
def dist(a, b):
    return math.sqrt(sum((x - y) ** 2 for x, y in zip(a, b)))`,
							Notes: "Swap in L1 (Manhattan) or cosine similarity as needed.",
						},
					},
				},
				{
					Name:        "Feature Scaling",
					Description: "Normalize inputs so distances are meaningful.",
					CoreIdeas:   []string{"Standardization", "Min-max"},
					Examples: []CodeExample{
						{
							ID:       "min_max_scale",
							Title:    "Min-max scaling",
							Language: "python",
							Snippet: `def min_max_scale(x, min_val, max_val):
    return (x - min_val) / (max_val - min_val)`,
							Notes: "Scale each feature so no single dimension dominates distance.",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "knn",
					Title: "KNN Classification",
					Goal:  "Show voting among nearest neighbors.",
					Steps: []StoryboardStep{
						{Cue: "plot", Narration: "Plot training points with labels.", VisualHint: "Scatter plot", Duration: Duration(3 * time.Second), CodeRef: "knn: training data"},
						{Cue: "query", Narration: "Drop in a query point.", VisualHint: "New point highlight", Duration: Duration(2 * time.Second), CodeRef: "knn: query vector"},
						{Cue: "neighbors", Narration: "Draw circles to capture k nearest neighbors.", VisualHint: "Expanding radius", Duration: Duration(3 * time.Second), CodeRef: "knn: distance + sort"},
						{Cue: "vote", Narration: "Neighbors vote for the label.", VisualHint: "Label tally", Duration: Duration(3 * time.Second), CodeRef: "knn: majority vote"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "feature-extract",
					Title:   "Feature Extraction",
					Outcome: "List candidate features and normalize them.",
					Steps: []TutorialStep{
						{Prompt: "List raw features.", Guidance: "Think length, color, weight.", CodeFocus: "struct fields", Checkpoint: "Features enumerated."},
						{Prompt: "Scale features.", Guidance: "Apply min-max or z-score.", CodeFocus: "normalize(values)", Checkpoint: "Comparable magnitudes."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "choose-k", Prompt: "How does choosing k too small or too large affect bias/variance?", Difficulty: "Medium", Outcome: "Understands sensitivity to k."},
			},
			Visualizers: []Visualizer{
				{ID: "knn-playground", Title: "KNN Playground", Goal: "Manipulate k and distance to see decision boundaries.", DataModel: "points + labels", Interactions: []string{"drag-point", "change-k"}, Hooks: []string{"onStep", "onReset", "onBoundaryChange"}},
			},
		},
		{
			Number:  11,
			Slug:    "where-to-go-next",
			Title:   "Where to Go Next",
			Summary: "Preview advanced topics and plug in new modules as learners progress.",
			Objectives: []string{
				"Offer next-step algorithms to explore",
				"Encourage experimentation with new visualizers",
				"Bridge toward distributed and probabilistic data structures",
			},
			Concepts: []Concept{
				{
					Name:        "Trees",
					Description: "Binary search trees and tries as next steps.",
					CoreIdeas:   []string{"Balancing", "Prefix storage"},
					Examples: []CodeExample{
						{
							ID:       "bst_insert",
							Title:    "BST insert",
							Language: "python",
							Snippet: `class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def insert(root, v):
    if root is None:
        return Node(v)
    if v < root.val:
        root.left = insert(root.left, v)
    else:
        root.right = insert(root.right, v)
    return root`,
							Notes: "Balancing strategies keep depth near log n.",
						},
					},
				},
				{
					Name:        "Distributed Thinking",
					Description: "MapReduce basics and parallelism.",
					CoreIdeas:   []string{"Map", "Reduce"},
					Examples: []CodeExample{
						{
							ID:       "mapreduce",
							Title:    "Map + reduce idea (Go map/reduce style)",
							Language: "python",
							Snippet: `mapped = []
for line in lines:
    mapped.extend(map_fn(line))
grouped = shuffle(mapped)  # group by key
results = {key: reduce_fn(key, vals) for key, vals in grouped.items()}`,
							Notes: "Map splits work; shuffle groups keys; reduce aggregates.",
						},
					},
				},
			},
			Animations: []Storyboard{
				{
					ID:    "mapreduce-mini",
					Title: "Mini MapReduce",
					Goal:  "Conceptually show map, shuffle, and reduce stages.",
					Steps: []StoryboardStep{
						{Cue: "map", Narration: "Map splits work across workers.", VisualHint: "Parallel boxes", Duration: Duration(3 * time.Second), CodeRef: "mapreduce: map_fn"},
						{Cue: "shuffle", Narration: "Shuffle groups keys together.", VisualHint: "Buckets per key", Duration: Duration(3 * time.Second), CodeRef: "mapreduce: shuffle/group"},
						{Cue: "reduce", Narration: "Reduce aggregates grouped values.", VisualHint: "Summaries per key", Duration: Duration(3 * time.Second), CodeRef: "mapreduce: reduce_fn"},
					},
				},
			},
			Tutorials: []Tutorial{
				{
					ID:      "next-steps",
					Title:   "Pick Your Next Topic",
					Outcome: "Draft a personalized next-steps track.",
					Steps: []TutorialStep{
						{Prompt: "Data structure depth?", Guidance: "Trees, tries, heaps.", CodeFocus: "topic selection", Checkpoint: "Focus area chosen."},
						{Prompt: "Systems angle?", Guidance: "Parallel algorithms, MapReduce.", CodeFocus: "topic selection", Checkpoint: "Stretch goal picked."},
					},
				},
			},
			Exercises: []Exercise{
				{ID: "design-plan", Prompt: "Design a 2-week plan to learn one advanced topic.", Difficulty: "Open", Outcome: "Concrete next steps."},
			},
			Visualizers: []Visualizer{
				{ID: "topic-map", Title: "Topic Map", Goal: "Display unlocked topics and suggested paths.", DataModel: "topic graph", Interactions: []string{"expand-topic", "add-note"}, Hooks: []string{"onStep", "onReset", "onProgressUpdate"}},
			},
		},
	}
}
