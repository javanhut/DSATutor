package chapter

import "time"

// chapter11 teaches heaps and priority queues. Storyboards: heap, kth-largest.
func chapter11() Chapter {
	return Chapter{
		Number:        11,
		Slug:          "heap-priority-queue",
		Title:         "Heaps and Priority Queues",
		Category:      "heap-priority-queue",
		Prerequisites: []int{9},
		Summary:       "A heap keeps the smallest (or largest) item on top of a nearly complete binary tree stored in an array. Reading the top is O(1); adding an item or removing the top is O(log n). Use it when you repeatedly need the current minimum or maximum of a changing collection.",
		Motivation: `A hospital emergency room has 1,000 waiting patients, and new ones arrive all the time. Every time a doctor is free, the most urgent patient must be called next. Keep the patients in a plain list and each call means scanning all 1,000 to find the most urgent: over a day with 50,000 calls, that's 50 million comparisons. Keep the list sorted instead and each call is instant, but every arrival must be inserted in place, shifting up to 1,000 entries.

A heap is the compromise that wins both ways. It doesn't keep the patients fully sorted, only "each parent is at least as urgent as its children". That weak order is enough to keep the most urgent at the top, and it can be repaired after an arrival or a call by fixing one path of about log2(1000), roughly 10 nodes. 50,000 calls then cost about half a million steps instead of 50 million.`,
		Objectives: []string{
			"Explain the heap property and why it puts the minimum at index 0 of the array",
			"Trace a push (sift up) and a pop (sift down) on a small heap and say which elements move",
			"Use Python's heapq as a min-heap, negate values to get a max-heap, and order entries with tuples",
			"Apply the 'keep a heap of size k' pattern and say whether it needs a min-heap or a max-heap",
			"Explain why heapify is O(n) while n pushes are O(n log n)",
		},
		WhenToUse: []string{
			"The problem repeatedly asks for the smallest, largest, earliest or most urgent item while items keep being added or removed",
			"'k largest', 'k smallest', 'k closest', 'k most frequent', or 'the k-th largest' (keep a heap of size k)",
			"Merging several sorted lists or streams into one sorted output (one entry per list in the heap)",
			"A running median or any 'middle of a stream' question (two heaps)",
			"A simulation where the next event is the one with the smallest time or the biggest count",
		},
		Concepts: []Concept{
			{
				Name:        "The heap property and its array layout",
				Description: "A binary heap is a complete binary tree (every level full except possibly the last, which fills left to right) where every parent is <= its children (a min-heap). It is stored in a plain list: the node at index i has children at 2i + 1 and 2i + 2 and its parent at (i - 1) // 2.",
				WhyItWorks:  "The heap property chains: a parent <= its children <= their children..., so the root is <= every element and the minimum is always heap[0]. Because the tree is complete, its height is floor(log2 n), about 20 for a million items. Push appends at the end (keeping the tree complete) and swaps the new item up while it's smaller than its parent; only that one path can be out of order, so at most log n swaps fix it. Pop moves the last item to the root and swaps it down with its smaller child until both children are larger; again only one path changes. That's why both cost O(log n) while the top is O(1).",
				Intuition:   "Think of a company org chart where every manager is senior to the people they manage. The CEO is guaranteed to be the most senior person, but nothing is promised between two people in different departments. When someone senior is hired at the bottom, they get promoted step by step past less senior managers (sift up); when the CEO leaves, someone from the bottom fills the top seat and is demoted step by step (sift down). The managers are parents, seniority is the value.",
				CoreIdeas: []string{
					"Min-heap property: heap[(i - 1) // 2] <= heap[i] for every i > 0. The minimum is heap[0].",
					"The heap is only partially ordered: heap[1] and heap[2] can be in either order, and the list is usually not sorted.",
					"push: append, then sift up along one path, O(log n). pop: move the last item to the root, then sift down, O(log n).",
					"heapify builds a heap from a list in O(n) by sifting down from the last parent to the root; most nodes are near the bottom and barely move.",
				},
				CommonMistakes: []string{
					"Assuming heap[1] is the second smallest or that the list is sorted. Only heap[0] is guaranteed; pop to get items in order.",
					"Building a heap with n separate pushes when you already have all the items. heapq.heapify(list) is O(n) instead of O(n log n).",
					"Searching or deleting an arbitrary item in a heap. That's O(n); heaps only give fast access to the top. Use a sorted structure or lazy deletion if you need it.",
				},
				Examples: []CodeExample{
					{ID: "heap-push", Title: "Push with sift up, by hand", Language: "python", Snippet: "def push(heap, value):\n    heap.append(value)                 # keep the tree complete\n    i = len(heap) - 1\n    while i > 0:\n        parent = (i - 1) // 2\n        if heap[parent] <= heap[i]:    # heap property holds: stop\n            break\n        heap[parent], heap[i] = heap[i], heap[parent]\n        i = parent                     # continue up one level\n\nheap = [1, 3, 2, 7, 4]\npush(heap, 0)\n# 0 lands at index 5, swaps with 2 (index 2), then with 1 (index 0)\n# heap == [0, 3, 1, 7, 4, 2]", Notes: "heapq.heappush does exactly this. Only the path from the new leaf to the root changes."},
				},
			},
			{
				Name:        "heapq is a min-heap: negation and tuples",
				Description: "Python's heapq module turns a plain list into a min-heap: heappush, heappop, heapify, heap[0] to peek, and heappushpop/heapreplace for a push and pop in one step. There is no max-heap version, so you store negated keys. To order records by a key, push tuples whose first element is the key.",
				WhyItWorks:  "Negation reverses every comparison: a > b exactly when -a < -b. So a min-heap of negated values keeps the largest original value on top, and negating again on the way out gives it back. Tuples compare element by element, left to right, so (key, payload) is ordered by key first; the payload only matters when keys tie. That's how K Closest Points pushes (-distance, x, y) and Design Twitter pushes (-time, tweetId, user, index): the heap orders by the first field and carries the rest along.",
				Intuition:   "A min-heap is a queue that always lets the shortest person go first. If you want the tallest to go first instead, give everyone a ticket with minus their height: the most negative ticket, the tallest person, now goes first. The ticket is the negated key; the person is the payload in the tuple.",
				CoreIdeas: []string{
					"Min-heap: push x, read heap[0]. Max-heap: push -x, read -heap[0].",
					"Order by a key: push (key, data). Negate the key for 'largest key first'.",
					"Ties compare the next tuple element, so only put comparable values after the key (numbers, strings), or add a counter as a tie-breaker.",
					"heapify(list) is O(n); heappushpop(heap, x) pushes then pops in one O(log n) step.",
				},
				CommonMistakes: []string{
					"Forgetting to negate on the way out. You push -5 and return heap[0], which is -5, not 5. Negate when pushing and when reading.",
					"Pushing (priority, object) where objects can't be compared. If two priorities tie, Python compares the objects and raises TypeError. Add a unique counter: (priority, counter, object).",
					"Negating a key but not the comparison you reason with: in a negated heap, 'bigger than the top' means the stored value is smaller. Convert back before comparing.",
				},
				Examples: []CodeExample{
					{ID: "heapq-basics", Title: "Min-heap, max-heap and keyed entries", Language: "python", Snippet: "import heapq\n\nnums = [5, 1, 8, 3]\nheapq.heapify(nums)               # O(n)\nprint(heapq.heappop(nums))        # 1 (smallest)\n\nmax_heap = [-x for x in [5, 1, 8, 3]]\nheapq.heapify(max_heap)\nprint(-heapq.heappop(max_heap))   # 8 (largest)\n\ntasks = []\nheapq.heappush(tasks, (2, \"write tests\"))\nheapq.heappush(tasks, (1, \"fix bug\"))\nprint(heapq.heappop(tasks))       # (1, 'fix bug'): ordered by the first field", Notes: "One module, three uses: min-heap, max-heap by negation, and priority queue of records by tuple key."},
				},
			},
			{
				Name:        "Keep a heap of size k",
				Description: "To track the k largest items of a collection or stream, keep a min-heap of at most k items: push each new item, and if the heap grows to k + 1, pop the top. The heap then holds the k largest, and its top is the k-th largest. For the k smallest, mirror it with a max-heap. Two heaps facing each other (a max-heap for the lower half, a min-heap for the upper half) extend the idea to the median.",
				WhyItWorks:  "Invariant: the heap holds the k largest items seen so far. When a new item arrives, the k largest of (old items + new item) are the old k plus the new one, minus the smallest of those k + 1, which is exactly push then pop-the-minimum. The item you evict is always the weakest of the kept ones, so the top of the heap must be the weakest: a min-heap when keeping the largest, a max-heap when keeping the smallest. Each item costs O(log k), so n items cost O(n log k) time and O(k) memory, better than sorting (O(n log n), O(n)) when k is much smaller than n, and it works on streams you can't store.",
				Intuition:   "A sports team with k roster spots. Every new player who tries out joins the roster; if that makes k + 1 players, the weakest is cut. The coach only needs to know who the weakest player is, so that's who stands at the front of the line. The roster is the heap, the weakest player is heap[0], and the k-th best player overall is that weakest roster member.",
				CoreIdeas: []string{
					"Keep the largest -> min-heap of size k (evict the smallest). Keep the smallest -> max-heap of size k (evict the largest).",
					"Push first, then trim: if len(heap) > k, pop. The top is the boundary value (the k-th best).",
					"Cost O(n log k) time, O(k) memory: good for small k and for streams.",
					"Two heaps (low max-heap, high min-heap, sizes within one) keep the middle values on top: a running median in O(log n) per insert.",
				},
				CommonMistakes: []string{
					"Picking the heap direction by the word in the question ('largest' -> max-heap). Pick it by what you evict: keeping the largest means evicting the smallest, so use a min-heap.",
					"Popping before pushing to keep the size at k. If the new item is worse than everything kept, you've thrown away a good item. Push then pop, or use heappushpop.",
					"Reading heap[0] before the heap has k items. Until k items have arrived, the top is not the k-th largest.",
				},
				Examples: []CodeExample{
					{ID: "top-k", Title: "The k largest values of a list", Language: "python", Snippet: "import heapq\n\ndef k_largest(nums, k):\n    heap = []\n    for num in nums:\n        heapq.heappush(heap, num)     # push the new item\n        if len(heap) > k:             # trim to size k\n            heapq.heappop(heap)       # evict the smallest kept value\n    return sorted(heap, reverse=True) # heap[0] alone is the k-th largest\n\nprint(k_largest([3, 2, 1, 5, 6, 4], 2))   # [6, 5]", Notes: "O(n log k). heapq.nlargest(k, nums) uses the same idea."},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Kth Largest Element in an Array",
				Problem: "Given an array nums and an integer k, return the k-th largest element (in sorted order, counting duplicates). Example: nums = [3,2,1,5,6,4], k = 2 -> 5, since sorted descending the array is [6,5,4,3,2,1].",
				Steps: []WorkedStep{
					{Subgoal: "Build the heap", Explanation: "Start with an empty min-heap. It will hold the k largest values seen so far, so its top is the weakest of them: the current k-th largest. A min-heap is right because the value we'll evict is always the smallest kept one.", Code: "heap = []", State: "heap = []"},
					{Subgoal: "Push the new item", Explanation: "Each number might belong to the top k, so push it. The heap may now hold k + 1 values.", Code: "for num in nums:\n    heapq.heappush(heap, num)", State: "push 3 -> [3]; push 2 -> [2,3]; push 1 -> [1,3,2]"},
					{Subgoal: "Trim to size k", Explanation: "With k + 1 values kept, the smallest of them can never be among the k largest (there are already k bigger-or-equal values), so pop it. This keeps the invariant: the heap holds the k largest values seen so far.", Code: "    if len(heap) > k:\n        heapq.heappop(heap)", State: "k=2: [1,3,2] pop 1 -> [2,3]; push 5 -> [2,3,5] pop 2 -> [3,5]; push 6 -> pop 3 -> [5,6]; push 4 -> [4,6,5] pop 4 -> [5,6]"},
					{Subgoal: "Read the top", Explanation: "After all numbers, the heap holds the 2 largest values {5, 6}; its top is the smaller of them, which is the 2nd largest overall. Total cost O(n log k) instead of O(n log n) for sorting.", Code: "return heap[0]", State: "heap = [5,6] -> return 5"},
				},
				Code: "import heapq\n\ndef findKthLargest(nums, k):\n    heap = []                          # build the heap\n    for num in nums:\n        heapq.heappush(heap, num)      # push the new item\n        if len(heap) > k:              # trim to size k\n            heapq.heappop(heap)\n    return heap[0]                     # read the top",
			},
		},
		Completions: []CompletionProblem{
			{
				Title:       "The k smallest values, with a max-heap",
				Prompt:      "Return the k smallest numbers in ascending order. Keep a heap of size k, but this time you must evict the largest kept value, so use a max-heap by negation. Fill in the blanks.",
				Code:        "import heapq\n\ndef k_smallest(nums, k):\n    heap = []                                # max-heap of negated values\n    for num in nums:\n        heapq.heappush(heap, ____)           # push the new item\n        if ____:                             # trim to size k\n            heapq.heappop(heap)              # evicts the largest kept value\n    return sorted(____ for x in heap)        # read the heap",
				Blanks:      []string{"-num", "len(heap) > k", "-x"},
				Explanation: "To keep the k smallest, the item to evict is the largest kept one, so it must be on top: a max-heap, built from heapq by pushing -num. After pushing, a heap of k + 1 items is trimmed by one pop, which removes the most negative stored value, i.e. the largest number. Reading back requires negating again.",
			},
		},
		Animations: []Storyboard{
			{
				ID:    "heap",
				Title: "Min-Heap Operations",
				Goal:  "Show heap insertion and bubble-up.",
				MemoryTips: []string{
					"Append at the end so the tree stays complete",
					"Compare with the parent at (i - 1) // 2",
					"Swap up while smaller than the parent; stop as soon as it isn't",
					"Only one root-to-leaf path changes: O(log n)",
				},
				Steps: []StoryboardStep{
					{Cue: "insert", Narration: "A push starts by appending the new value at the end of the list: the next free spot in the bottom level of the tree. The tree stays complete, but the new value may be smaller than its parent, breaking the heap property on this one path.", VisualHint: "Add to array", Duration: Duration(2 * time.Second), CodeRef: "heap: heap.append(val)"},
					{Cue: "compare", Narration: "Find the parent at index (i - 1) // 2 and compare. If the parent is smaller or equal, the heap property already holds and we're done.", VisualHint: "Parent-child comparison", Duration: Duration(2 * time.Second), CodeRef: "heap: if heap[idx] < heap[parent]"},
					{Cue: "swap", Narration: "If the new value is smaller than its parent, swap them and repeat from the parent's index. Each swap moves the value up one level, so there are at most log2(n) swaps.", VisualHint: "Bubble up", Duration: Duration(2 * time.Second), CodeRef: "heap: swap elements"},
					{Cue: "done", Narration: "The value stops when its parent is smaller or it reaches the root. Every parent is now <= its children again, so the minimum sits at index 0. A pop is the mirror image: move the last value to the root and sift it down.", VisualHint: "Min at root", Duration: Duration(2 * time.Second), CodeRef: "heap: property satisfied"},
				},
			},
			{
				ID:    "kth-largest",
				Title: "Kth Largest via Min-Heap",
				Goal:  "Keep only k largest elements with a min-heap.",
				MemoryTips: []string{
					"Push the new item",
					"Trim to size k: pop when the heap holds k + 1",
					"Read the top: the weakest of the k best is the k-th best",
					"Keeping the largest? Evict the smallest, so use a min-heap",
				},
				Steps: []StoryboardStep{
					{Cue: "push", Narration: "Push each new number onto a min-heap. The heap is meant to hold the k largest numbers seen so far, and the new one might belong there.", VisualHint: "Heap grows", Duration: Duration(2 * time.Second), CodeRef: "heappush"},
					{Cue: "trim", Narration: "If the heap now holds k + 1 numbers, pop once. The popped value is the smallest of them, and with k values at least as big it can never be the k-th largest again.", VisualHint: "Remove smallest", Duration: Duration(2 * time.Second), CodeRef: "heappop"},
					{Cue: "answer", Narration: "The heap holds exactly the k largest numbers, and its root is the smallest of those: the k-th largest overall. Each step cost O(log k).", VisualHint: "Root = answer", Duration: Duration(2 * time.Second), CodeRef: "heap[0]"},
				},
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt:      "Trace: the min-heap is stored as [1, 3, 2, 7, 4]. Which values are the children of 3?",
				Choices:     []string{"2 and 7", "7 and 4", "1 and 2", "4 only"},
				Answer:      1,
				Explanation: "3 is at index 1, so its children are at 2*1 + 1 = 3 and 2*1 + 2 = 4: the values 7 and 4. Both are >= 3, as the heap property requires.",
			},
			{
				Prompt:      "Which pattern fits best? Numbers arrive one at a time (millions of them), and after each one you must report the 100th largest number so far.",
				Choices:     []string{"Append to a list and sort it after every number", "A min-heap holding the 100 largest numbers seen so far", "Binary search on the list of numbers received", "A sliding window of the last 100 numbers"},
				Answer:      1,
				Explanation: "A size-k min-heap answers each update in O(log 100) and reads the answer at heap[0]. Re-sorting is O(n log n) per number. Binary search (chapter 6) needs the list sorted, and inserting into a sorted list costs O(n). A sliding window (chapter 4) only sees the last 100 numbers, not the largest 100.",
			},
			{
				Prompt:      "You want to keep the k *closest* points while scanning a list. Which heap do you use with heapq?",
				Choices:     []string{"Push (distance, point) into a min-heap and pop when size > k", "Push (-distance, point) into heapq and pop when size > k", "Push every point, then pop k times at the end", "A heap isn't useful here"},
				Answer:      1,
				Explanation: "Keeping the smallest distances means evicting the largest, so the largest kept distance must be on top: a max-heap, which in heapq means negating the key. Option 1 evicts the closest point. Option 3 is correct but O(n log n) time and O(n) memory, and doesn't work on a stream.",
			},
			{
				Prompt:      "Prerequisite check (chapter 9): why is a heap of 1,000,000 items only about 20 levels tall, and why does that matter?",
				Explanation: "A heap is a complete binary tree: every level except the last is full, so level d holds 2^d nodes and 20 levels hold about 2^20 = 1,048,576 nodes. Push and pop each move one item along a single root-to-leaf path, so they take at most about 20 swaps. That's what makes them O(log n).",
			},
			{
				Prompt:      "Why does heapq.heapify(list) run in O(n), when pushing the same n items one by one takes O(n log n)?",
				Explanation: "heapify sifts nodes down starting from the last parent and moving towards the root. Half the nodes are leaves and never move; a quarter can move at most 1 level; an eighth at most 2; and so on. The total, n/4 * 1 + n/8 * 2 + n/16 * 3 + ..., is less than n. Pushing one by one sifts each new item up from the bottom, which can take log n swaps for most of them.",
			},
			{
				Prompt:      "What does this print?\n\nimport heapq\nh = []\nfor x in [4, 9, 1, 7]:\n    heapq.heappush(h, -x)\n    if len(h) > 2:\n        heapq.heappop(h)\nprint(sorted(-v for v in h))",
				Choices:     []string{"[7, 9]", "[1, 4]", "[4, 9]", "[1, 7]"},
				Answer:      1,
				Explanation: "Pushing -x makes it a max-heap, and popping when size > 2 evicts the largest kept value. So the heap keeps the 2 smallest: after 4, 9 -> {4, 9}; push 1, evict 9 -> {1, 4}; push 7, evict 7 -> {1, 4}. It prints [1, 4].",
			},
		},
	}
}
