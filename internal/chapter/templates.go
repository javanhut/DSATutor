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
				{
					ID:    "linked-list",
					Title: "Linked List Traversal",
					Goal:  "Show searching through a linked list node by node.",
					Steps: []StoryboardStep{
						{Cue: "start", Narration: "Start at the head of the list.", VisualHint: "Pointer at head", Duration: Duration(2 * time.Second), CodeRef: "linked_list: current = head"},
						{Cue: "compare", Narration: "Compare current node's value with target.", VisualHint: "Highlight comparison", Duration: Duration(2 * time.Second), CodeRef: "linked_list: if current.val == target"},
						{Cue: "advance", Narration: "Move to the next node.", VisualHint: "Follow pointer", Duration: Duration(2 * time.Second), CodeRef: "linked_list: current = current.next"},
						{Cue: "found", Narration: "Continue until found or end reached.", VisualHint: "Loop indicator", Duration: Duration(2 * time.Second), CodeRef: "linked_list: while current"},
					},
				},
				{
					ID:    "two-pointers",
					Title: "Floyd's Cycle Detection",
					Goal:  "Detect cycles using slow and fast pointers.",
					Steps: []StoryboardStep{
						{Cue: "init", Narration: "Initialize slow and fast pointers at head.", VisualHint: "Two pointers at start", Duration: Duration(2 * time.Second), CodeRef: "cycle: slow = fast = head"},
						{Cue: "move", Narration: "Slow moves 1 step, fast moves 2 steps.", VisualHint: "Different speeds", Duration: Duration(3 * time.Second), CodeRef: "cycle: slow.next, fast.next.next"},
						{Cue: "check", Narration: "Check if pointers meet.", VisualHint: "Collision detection", Duration: Duration(2 * time.Second), CodeRef: "cycle: if slow == fast"},
						{Cue: "cycle", Narration: "If they meet, a cycle exists!", VisualHint: "Cycle found", Duration: Duration(2 * time.Second), CodeRef: "cycle: return True"},
					},
				},
				{
					ID:    "merge-sort",
					Title: "Merge Sort Divide and Conquer",
					Goal:  "Show recursive splitting and merging of arrays.",
					Steps: []StoryboardStep{
						{Cue: "divide", Narration: "Recursively divide the array in half.", VisualHint: "Split animation", Duration: Duration(3 * time.Second), CodeRef: "merge_sort: mid = len(arr) // 2"},
						{Cue: "base", Narration: "Stop when arrays have 1 element.", VisualHint: "Base case", Duration: Duration(2 * time.Second), CodeRef: "merge_sort: if len(arr) <= 1"},
						{Cue: "merge", Narration: "Merge sorted subarrays back together.", VisualHint: "Merge animation", Duration: Duration(3 * time.Second), CodeRef: "merge: compare and combine"},
						{Cue: "complete", Narration: "Final sorted array emerges.", VisualHint: "Sorted result", Duration: Duration(2 * time.Second), CodeRef: "merge_sort: return result"},
					},
				},
				{
					ID:    "sliding-window",
					Title: "Sliding Window Technique",
					Goal:  "Find maximum sum of k consecutive elements.",
					Steps: []StoryboardStep{
						{Cue: "init", Narration: "Calculate initial window sum.", VisualHint: "First k elements", Duration: Duration(2 * time.Second), CodeRef: "sliding: sum(nums[:k])"},
						{Cue: "slide", Narration: "Slide window: add new, remove old.", VisualHint: "Window moves right", Duration: Duration(3 * time.Second), CodeRef: "sliding: += nums[i] - nums[i-k]"},
						{Cue: "track", Narration: "Track maximum sum seen.", VisualHint: "Update max", Duration: Duration(2 * time.Second), CodeRef: "sliding: max_sum = max(...)"},
						{Cue: "complete", Narration: "Return the maximum sum found.", VisualHint: "Final answer", Duration: Duration(2 * time.Second), CodeRef: "sliding: return max_sum"},
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
				{
					ID:    "rotated-array",
					Title: "Search in Rotated Sorted Array",
					Goal:  "Binary search on a rotated array.",
					Steps: []StoryboardStep{
						{Cue: "init", Narration: "Initialize low and high pointers.", VisualHint: "Array with rotation point", Duration: Duration(2 * time.Second), CodeRef: "rotated: low, high = 0, len-1"},
						{Cue: "find-sorted", Narration: "Identify which half is sorted.", VisualHint: "Highlight sorted half", Duration: Duration(3 * time.Second), CodeRef: "rotated: nums[low] <= nums[mid]"},
						{Cue: "narrow", Narration: "Narrow search to correct half.", VisualHint: "Discard half", Duration: Duration(3 * time.Second), CodeRef: "rotated: adjust bounds"},
						{Cue: "found", Narration: "Target found or not in array.", VisualHint: "Result", Duration: Duration(2 * time.Second), CodeRef: "rotated: return mid or -1"},
					},
				},
				{
					ID:    "backtracking",
					Title: "N-Queens Backtracking",
					Goal:  "Solve N-Queens using backtracking.",
					Steps: []StoryboardStep{
						{Cue: "place", Narration: "Try placing a queen in a row.", VisualHint: "Queen placement", Duration: Duration(2 * time.Second), CodeRef: "nqueens: place queen"},
						{Cue: "check", Narration: "Check if placement is safe.", VisualHint: "Conflict detection", Duration: Duration(2 * time.Second), CodeRef: "nqueens: is_safe()"},
						{Cue: "recurse", Narration: "Recurse to next row if safe.", VisualHint: "Continue building", Duration: Duration(2 * time.Second), CodeRef: "nqueens: backtrack(row+1)"},
						{Cue: "backtrack", Narration: "Backtrack if no valid placement.", VisualHint: "Remove queen", Duration: Duration(2 * time.Second), CodeRef: "nqueens: board[row] = -1"},
					},
				},
				{
					ID:    "permutations",
					Title: "Generate Permutations",
					Goal:  "Generate all permutations using backtracking.",
					Steps: []StoryboardStep{
						{Cue: "choose", Narration: "Choose an element from remaining.", VisualHint: "Select item", Duration: Duration(2 * time.Second), CodeRef: "permute: for num in remaining"},
						{Cue: "add", Narration: "Add to current path.", VisualHint: "Build path", Duration: Duration(2 * time.Second), CodeRef: "permute: path.append(num)"},
						{Cue: "recurse", Narration: "Recurse with remaining elements.", VisualHint: "Continue", Duration: Duration(2 * time.Second), CodeRef: "permute: backtrack(path, remaining)"},
						{Cue: "pop", Narration: "Backtrack by removing last element.", VisualHint: "Pop from path", Duration: Duration(2 * time.Second), CodeRef: "permute: path.pop()"},
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
					ID:    "hash-table",
					Title: "Hash Table Operations",
					Goal:  "Show hashing, collision handling with chaining, and lookup.",
					Steps: []StoryboardStep{
						{Cue: "hash", Narration: "Hash the key to find a bucket.", VisualHint: "Key -> hash -> bucket index", Duration: Duration(3 * time.Second), CodeRef: "hash_table: hash(key) -> bucket"},
						{Cue: "collision", Narration: "Handle collisions with chaining.", VisualHint: "Linked list in bucket", Duration: Duration(3 * time.Second), CodeRef: "hash_table: collision handling"},
						{Cue: "insert", Narration: "Insert key-value pair into bucket.", VisualHint: "Append to chain", Duration: Duration(2 * time.Second), CodeRef: "hash_table: bucket.append"},
						{Cue: "lookup", Narration: "Lookup hashes key and searches chain.", VisualHint: "O(1) average", Duration: Duration(2 * time.Second), CodeRef: "hash_table: lookup"},
					},
				},
				{
					ID:    "string-match",
					Title: "String Pattern Matching",
					Goal:  "Find pattern occurrences in text.",
					Steps: []StoryboardStep{
						{Cue: "init", Narration: "Start comparing pattern at position 0.", VisualHint: "Align pattern", Duration: Duration(2 * time.Second), CodeRef: "string: for i in range(n-m+1)"},
						{Cue: "compare", Narration: "Compare characters one by one.", VisualHint: "Character match/mismatch", Duration: Duration(3 * time.Second), CodeRef: "string: while text[i+j] == pattern[j]"},
						{Cue: "match", Narration: "Full match found!", VisualHint: "Pattern matched", Duration: Duration(2 * time.Second), CodeRef: "string: if j == m: return i"},
						{Cue: "shift", Narration: "Shift pattern and try next position.", VisualHint: "Move pattern right", Duration: Duration(2 * time.Second), CodeRef: "string: continue loop"},
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
				{
					ID:    "dfs",
					Title: "DFS Depth-First Traversal",
					Goal:  "Show stack-based deep exploration of a graph.",
					Steps: []StoryboardStep{
						{Cue: "push-start", Narration: "Push the start node onto the stack.", VisualHint: "Stack with start node", Duration: Duration(2 * time.Second), CodeRef: "dfs: stack.append(start)"},
						{Cue: "pop", Narration: "Pop a node and visit if not seen.", VisualHint: "LIFO order", Duration: Duration(3 * time.Second), CodeRef: "dfs: node = stack.pop()"},
						{Cue: "explore", Narration: "Push unvisited neighbors onto stack.", VisualHint: "Go deep first", Duration: Duration(3 * time.Second), CodeRef: "dfs: push neighbors"},
						{Cue: "backtrack", Narration: "Backtrack when no more unvisited neighbors.", VisualHint: "Stack shrinks", Duration: Duration(2 * time.Second), CodeRef: "dfs: implicit backtrack"},
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
					Title: "Dijkstra's Algorithm",
					Goal:  "Show cost updates, priority queue, and edge relaxation.",
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
					ID:    "greedy",
					Title: "Activity Selection (Greedy)",
					Goal:  "Show picking earliest finishing activities first.",
					Steps: []StoryboardStep{
						{Cue: "sort", Narration: "Sort activities by finish time.", VisualHint: "Timeline sorted", Duration: Duration(3 * time.Second), CodeRef: "activity_selection: sort by end"},
						{Cue: "consider", Narration: "Consider each activity in sorted order.", VisualHint: "Highlight current", Duration: Duration(2 * time.Second), CodeRef: "activity_selection: for activity"},
						{Cue: "check", Narration: "Check if start >= current_end.", VisualHint: "Compatibility check", Duration: Duration(2 * time.Second), CodeRef: "activity_selection: if compatible"},
						{Cue: "select", Narration: "Select compatible activity, update end.", VisualHint: "Add to solution", Duration: Duration(3 * time.Second), CodeRef: "activity_selection: select"},
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
					ID:    "knapsack",
					Title: "0/1 Knapsack DP",
					Goal:  "Show building the 2D value table for knapsack.",
					Steps: []StoryboardStep{
						{Cue: "init", Narration: "Initialize DP table with zeros.", VisualHint: "Base case: 0 items = 0 value", Duration: Duration(2 * time.Second), CodeRef: "knapsack: init dp"},
						{Cue: "consider", Narration: "Consider each item one by one.", VisualHint: "Highlight current item", Duration: Duration(2 * time.Second), CodeRef: "knapsack: for each item"},
						{Cue: "decide", Narration: "For each capacity, decide include vs exclude.", VisualHint: "Compare options", Duration: Duration(3 * time.Second), CodeRef: "knapsack: max(exclude, include)"},
						{Cue: "fill", Narration: "Fill cell with maximum value.", VisualHint: "Update table", Duration: Duration(2 * time.Second), CodeRef: "knapsack: dp[i][w] = ..."},
						{Cue: "traceback", Narration: "Traceback to find selected items.", VisualHint: "Recover solution", Duration: Duration(3 * time.Second), CodeRef: "knapsack: traceback"},
					},
				},
				{
					ID:    "memoization",
					Title: "Memoization with Fibonacci",
					Goal:  "Show how memoization avoids recomputation.",
					Steps: []StoryboardStep{
						{Cue: "call", Narration: "Make recursive call for fib(n).", VisualHint: "Call stack grows", Duration: Duration(2 * time.Second), CodeRef: "fib: call fib(n)"},
						{Cue: "check-memo", Narration: "Check if result is already in memo.", VisualHint: "Memo table lookup", Duration: Duration(2 * time.Second), CodeRef: "fib: if n in memo"},
						{Cue: "compute", Narration: "Compute and store result if not memoized.", VisualHint: "Store in memo", Duration: Duration(3 * time.Second), CodeRef: "fib: memo[n] = ..."},
						{Cue: "return", Narration: "Return memoized result.", VisualHint: "Stack unwinds", Duration: Duration(2 * time.Second), CodeRef: "fib: return memo[n]"},
					},
				},
				{
					ID:    "iterative-dp",
					Title: "Iterative DP (Bottom-Up)",
					Goal:  "Show bottom-up tabulation for Fibonacci.",
					Steps: []StoryboardStep{
						{Cue: "init", Narration: "Initialize dp array with base cases.", VisualHint: "dp[0]=0, dp[1]=1", Duration: Duration(2 * time.Second), CodeRef: "iterative: dp[0], dp[1]"},
						{Cue: "fill", Narration: "Fill table from small to large.", VisualHint: "Build up solutions", Duration: Duration(3 * time.Second), CodeRef: "iterative: dp[i] = dp[i-1] + dp[i-2]"},
						{Cue: "dependencies", Narration: "Each cell depends on previous cells.", VisualHint: "Show dependencies", Duration: Duration(2 * time.Second), CodeRef: "iterative: use previous values"},
						{Cue: "result", Narration: "Final answer in dp[n].", VisualHint: "Return result", Duration: Duration(2 * time.Second), CodeRef: "iterative: return dp[n]"},
					},
				},
				{
					ID:    "topsort",
					Title: "Topological Sort (Kahn's)",
					Goal:  "Show in-degree based topological ordering.",
					Steps: []StoryboardStep{
						{Cue: "calc-indegree", Narration: "Calculate in-degree for each node.", VisualHint: "Count incoming edges", Duration: Duration(3 * time.Second), CodeRef: "topsort: in_degree[v] += 1"},
						{Cue: "init-queue", Narration: "Queue all nodes with in-degree 0.", VisualHint: "Start nodes queued", Duration: Duration(2 * time.Second), CodeRef: "topsort: queue zero-degree"},
						{Cue: "process", Narration: "Dequeue, add to result, decrement neighbors.", VisualHint: "Process and update", Duration: Duration(4 * time.Second), CodeRef: "topsort: process node"},
						{Cue: "complete", Narration: "Continue until queue is empty.", VisualHint: "Ordering complete", Duration: Duration(2 * time.Second), CodeRef: "topsort: return result"},
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
				{
					ID:    "binary-tree",
					Title: "Binary Tree Traversals",
					Goal:  "Show inorder, preorder, and postorder traversals.",
					Steps: []StoryboardStep{
						{Cue: "left", Narration: "Go left recursively first.", VisualHint: "Traverse left subtree", Duration: Duration(2 * time.Second), CodeRef: "inorder: inorder(node.left)"},
						{Cue: "visit", Narration: "Visit the current node.", VisualHint: "Process node", Duration: Duration(2 * time.Second), CodeRef: "inorder: print(node.val)"},
						{Cue: "right", Narration: "Then go right.", VisualHint: "Traverse right subtree", Duration: Duration(2 * time.Second), CodeRef: "inorder: inorder(node.right)"},
						{Cue: "complete", Narration: "Traversal complete.", VisualHint: "Show result order", Duration: Duration(2 * time.Second), CodeRef: "inorder: return result"},
					},
				},
				{
					ID:    "heap",
					Title: "Min-Heap Operations",
					Goal:  "Show heap insertion and bubble-up.",
					Steps: []StoryboardStep{
						{Cue: "insert", Narration: "Insert element at the end.", VisualHint: "Add to array", Duration: Duration(2 * time.Second), CodeRef: "heap: heap.append(val)"},
						{Cue: "compare", Narration: "Compare with parent.", VisualHint: "Parent-child comparison", Duration: Duration(2 * time.Second), CodeRef: "heap: if heap[idx] < heap[parent]"},
						{Cue: "swap", Narration: "Swap if child is smaller.", VisualHint: "Bubble up", Duration: Duration(2 * time.Second), CodeRef: "heap: swap elements"},
						{Cue: "done", Narration: "Heap property restored.", VisualHint: "Min at root", Duration: Duration(2 * time.Second), CodeRef: "heap: property satisfied"},
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
		{
			Number:  12,
			Slug:    "blind-75-visuals",
			Title:   "Blind 75 Visual Walkthroughs",
			Summary: "High-value interview problems with compact, purpose-built animations.",
			Objectives: []string{
				"Trace how supporting data structures evolve step by step",
				"Connect problem statements to concrete algorithm flows",
				"Build intuition for backtracking, heaps, and binary search partitions",
			},
			Animations: []Storyboard{
				{
					ID:    "min-stack",
					Title: "Min Stack Operations",
					Goal:  "See how value/min pairs keep O(1) min queries.",
					Steps: []StoryboardStep{
						{Cue: "push-min", Narration: "Push values while tracking current min.", VisualHint: "Stack nodes carry value + min", Duration: Duration(3 * time.Second), CodeRef: "push"},
						{Cue: "query", Narration: "Get top and min in O(1).", VisualHint: "Read last tuple", Duration: Duration(2 * time.Second), CodeRef: "top/getMin"},
						{Cue: "pop", Narration: "Pop restores previous min.", VisualHint: "Remove top tuple", Duration: Duration(2 * time.Second), CodeRef: "pop"},
					},
				},
				{
					ID:    "lru-cache",
					Title: "LRU Cache",
					Goal:  "Visualize MRU/LRU order alongside hash map lookups.",
					Steps: []StoryboardStep{
						{Cue: "put", Narration: "Insert or update a key and move it to MRU.", VisualHint: "Order list bumps key to front", Duration: Duration(3 * time.Second), CodeRef: "put"},
						{Cue: "evict", Narration: "Evict LRU when capacity is exceeded.", VisualHint: "Drop tail node", Duration: Duration(3 * time.Second), CodeRef: "evict"},
						{Cue: "get", Narration: "Reads also refresh recency.", VisualHint: "Move hit to front", Duration: Duration(2 * time.Second), CodeRef: "get"},
					},
				},
				{
					ID:    "trie-ops",
					Title: "Trie Insert & Search",
					Goal:  "Step through character-by-character traversal.",
					Steps: []StoryboardStep{
						{Cue: "insert", Narration: "Create nodes along the word path.", VisualHint: "Grow children per letter", Duration: Duration(3 * time.Second), CodeRef: "insert"},
						{Cue: "mark-end", Narration: "Mark terminal nodes to store full words.", VisualHint: "End flag", Duration: Duration(2 * time.Second), CodeRef: "end"},
						{Cue: "search", Narration: "Walk the path for search and prefix checks.", VisualHint: "Highlight traversed edges", Duration: Duration(3 * time.Second), CodeRef: "search/startsWith"},
					},
				},
				{
					ID:    "car-fleet",
					Title: "Car Fleet",
					Goal:  "Show how arrival times form fleets with a monotonic stack.",
					Steps: []StoryboardStep{
						{Cue: "sort", Narration: "Sort cars from closest to farthest from target.", VisualHint: "Positions descending", Duration: Duration(2 * time.Second), CodeRef: "sort"},
						{Cue: "time", Narration: "Compute each car's arrival time.", VisualHint: "Time to target", Duration: Duration(2 * time.Second), CodeRef: "time"},
						{Cue: "stack", Narration: "Merge into fleets when arrival time is slower.", VisualHint: "Monotonic stack of times", Duration: Duration(3 * time.Second), CodeRef: "stack"},
					},
				},
				{
					ID:    "median-two-arrays",
					Title: "Median of Two Sorted Arrays",
					Goal:  "Follow binary search over partitions to find the median.",
					Steps: []StoryboardStep{
						{Cue: "partition", Narration: "Place partition in the shorter array.", VisualHint: "Partition bars", Duration: Duration(3 * time.Second), CodeRef: "partition"},
						{Cue: "check", Narration: "Validate partition correctness with max-left/min-right.", VisualHint: "Compare boundaries", Duration: Duration(3 * time.Second), CodeRef: "bounds"},
						{Cue: "adjust", Narration: "Move left/right until balanced, then compute median.", VisualHint: "Shift partition", Duration: Duration(3 * time.Second), CodeRef: "adjust"},
					},
				},
				{
					ID:    "kth-largest",
					Title: "Kth Largest via Min-Heap",
					Goal:  "Keep only k largest elements with a min-heap.",
					Steps: []StoryboardStep{
						{Cue: "push", Narration: "Push each element onto the heap.", VisualHint: "Heap grows", Duration: Duration(2 * time.Second), CodeRef: "heappush"},
						{Cue: "trim", Narration: "Pop when heap size exceeds k.", VisualHint: "Remove smallest", Duration: Duration(2 * time.Second), CodeRef: "heappop"},
						{Cue: "answer", Narration: "Heap root holds the kth largest.", VisualHint: "Root = answer", Duration: Duration(2 * time.Second), CodeRef: "heap[0]"},
					},
				},
				{
					ID:    "n-queens",
					Title: "N-Queens Backtracking",
					Goal:  "Place queens row by row and backtrack on conflicts.",
					Steps: []StoryboardStep{
						{Cue: "place", Narration: "Try a column for the current row.", VisualHint: "Drop queen", Duration: Duration(2 * time.Second), CodeRef: "place"},
						{Cue: "attack-check", Narration: "Skip columns under attack.", VisualHint: "Mark diagonals/columns", Duration: Duration(3 * time.Second), CodeRef: "is_safe"},
						{Cue: "solution", Narration: "Record a full board when all rows filled.", VisualHint: "Board snapshot", Duration: Duration(2 * time.Second), CodeRef: "append solution"},
					},
				},
			},
		},
	}
}
