package chapter

import "time"

// chapter11 teaches heaps and priority queues. Storyboards: heap, kth-largest, heap-sweep.
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
			"Sweep a sorted list of events with a heap of the active items, and explain why lazy deletion of expired items is safe",
		},
		WhenToUse: []string{
			"The problem repeatedly asks for the smallest, largest, earliest or most urgent item while items keep being added or removed",
			"'k largest', 'k smallest', 'k closest', 'k most frequent', or 'the k-th largest' (keep a heap of size k)",
			"Merging several sorted lists or streams into one sorted output (one entry per list in the heap)",
			"A running median or any 'middle of a stream' question (two heaps)",
			"A simulation where the next event is the one with the smallest time or the biggest count",
			"Items that start and end on a timeline (meetings, buildings, trips) and, at each event, you need the max or min among the ones active right now (sweep line with a heap)",
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
			{
				Name:        "Sweep line with a heap of active items",
				Description: "Many problems give you items that live on a timeline: meetings with a start and an end, buildings with a left and a right wall, trips with a pick-up and a drop-off. A sweep line sorts the start (and end) points and walks over them from left to right, like a vertical line moving across the timeline. At each event it updates the set of items that are active at that position, and a heap keyed by what you need next (the tallest building, the earliest-ending meeting, the lowest free room number) answers the question for that position from its top.",
				WhyItWorks:  "Between two consecutive events no item starts or ends, so whatever you are measuring (the tallest active height, the number of riders, which rooms are busy) is constant there; checking only at the event points loses nothing. At each event you first push the items that have started, then pop expired items from the top until the top is still active. After that the top is the best of the active items: every active item is in the heap, and any expired item left behind is buried below an active one, where it cannot affect the answer. This is lazy deletion, and it is needed because a heap can only remove its top cheaply; finding an arbitrary item inside it is O(n). Each item is pushed once and popped at most once, so after the O(n log n) sort the whole sweep is O(n log n).",
				Intuition:   "Picture a receptionist walking along a day's schedule with a stack of name cards, ordered so the person she cares about most is always on top. When the clock reaches someone's arrival she adds their card (push). She never digs through the stack to remove people who have left; only when she glances at the top card does she check whether that person is still in the building, and if not she throws the card away and looks again (lazy deletion). The clock times are the sorted events, the stack is the heap, and the top card after cleaning is the answer at that moment.",
				CoreIdeas: []string{
					"Sort the events by time (or x). Process them in order; the answer can only change at an event.",
					"Key the heap by what you must read at each event: (-height, end) for the tallest active item, (end, ...) for the next one to finish, (room_number) for the lowest free room.",
					"Release what has ended before reading the top: while heap[0] has expired, pop it. Expired items below the top can stay until they surface.",
					"If the heap is keyed by end time, the top is always the next to expire, so the release step pops exactly the expired items (no lazy deletion needed). Lazy deletion matters when the heap is keyed by something else, like height.",
				},
				CommonMistakes: []string{
					"Reading heap[0] before popping expired tops. The top may be a building that ended long ago; always run the 'while the top has expired, pop' loop first, then read.",
					"Getting ties at equal times wrong. Decide from the problem whether an item ending at t and one starting at t overlap: with half-open ranges [start, end), test end <= t so the old item is released before the new one counts.",
					"Trying to remove an ended item from the middle of the heap with list.remove or a search. That is O(n) per removal and breaks the heap order unless you re-heapify. Leave it in and drop it when it reaches the top.",
				},
				Examples: []CodeExample{
					{ID: "sweep-heap", Title: "The tallest active interval at each query time", Language: "python", Snippet: "import heapq\n\ndef tallest_active(intervals, times):\n    # intervals: [start, end, value], active while start <= t < end\n    # times: query times, sorted ascending\n    intervals = sorted(intervals)                 # sort the events\n    active = []                                   # max-heap of (-value, end)\n    answers = []\n    i = 0\n    for t in times:\n        while i < len(intervals) and intervals[i][0] <= t:\n            start, end, value = intervals[i]\n            heapq.heappush(active, (-value, end)) # add what has started\n            i += 1\n        while active and active[0][1] <= t:       # release what has ended (lazily)\n            heapq.heappop(active)\n        answers.append(-active[0][0] if active else 0)   # take the best active item\n    return answers\n\nprint(tallest_active([[1, 6, 5], [2, 4, 8], [3, 9, 6]], [2, 3, 4, 7, 9]))\n# [8, 8, 6, 6, 0]: at t = 7 the 5 (ended at 6) is still buried in the heap, harmlessly", Notes: "Times 2 and 3: the 8 is on top. Time 4: the 8 has ended (end 4 <= 4) and is popped, so the 6 shows. Time 9: the 6 and the buried 5 both surface and are popped in turn, leaving nothing active."},
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
			{
				Title:   "Car Pooling with a drop-off heap",
				Problem: "A car with `capacity` empty seats drives east and never turns back. trips[i] = [passengers, pick_up, drop_off] means that many people board at kilometre pick_up and leave at kilometre drop_off (they are out of the car by then, so their seats are free for anyone boarding at drop_off). Return True if every trip fits without ever exceeding capacity. Example: trips = [[3,3,7],[2,1,5],[2,5,8]], capacity = 5 -> True (the car never holds more than 5); with capacity = 4 -> False (5 people are aboard from kilometre 3 to 5).",
				Steps: []WorkedStep{
					{Subgoal: "Sort the events", Explanation: "The load can only change where someone boards or leaves, so sweep over the pick-ups in order. Sort the trips by pick-up; the drop-offs will be handled by the heap. riding is a min-heap of (drop_off, passengers): the trips currently in the car, with the next one to leave on top.", Code: "trips = sorted(trips, key=lambda trip: trip[1])\nriding = []\nload = 0", State: "trips = [[2,1,5],[3,3,7],[2,5,8]], riding = [], load = 0"},
					{Subgoal: "Release what has ended", Explanation: "Before anyone boards at pick_up, everyone whose drop_off <= pick_up has left. The heap is keyed by drop_off, so they are exactly the items at the top: pop until the top's drop_off is later than pick_up, subtracting their passengers. Using <= rather than < frees seats at the same kilometre where someone else boards.", Code: "for passengers, pick_up, drop_off in trips:\n    while riding and riding[0][0] <= pick_up:\n        load -= heapq.heappop(riding)[1]", State: "pick-up 1: heap empty, nothing to release. pick-up 3: top (5,2), 5 > 3, nothing released. pick-up 5: top (5,2), 5 <= 5 -> pop, load 5 -> 3, riding = [(7,3)]"},
					{Subgoal: "Add the new item", Explanation: "The new trip is now active: push it with its drop-off as the key so the release step can find it later, and add its passengers to the running load.", Code: "    heapq.heappush(riding, (drop_off, passengers))\n    load += passengers", State: "pick-up 1: riding = [(5,2)], load 2. pick-up 3: riding = [(5,2),(7,3)], load 5. pick-up 5: riding = [(7,3),(8,2)], load 5"},
					{Subgoal: "Check the answer", Explanation: "The load only grows at a pick-up, so a pick-up is the only place it can exceed capacity. Check right after boarding; one overload is enough to answer False.", Code: "    if load > capacity:\n        return False", State: "capacity 5: load 2, 5, 5 never exceeds 5, keep going. capacity 4: at pick-up 3 load 5 > 4 -> return False"},
					{Subgoal: "Return the result", Explanation: "If every pick-up passed the check, the car was never over capacity at any kilometre, because between pick-ups the load can only drop. Each trip is pushed and popped at most once: O(n log n) including the sort.", Code: "return True", State: "capacity 5: all three pick-ups pass -> return True"},
				},
				Code: "import heapq\n\ndef carPooling(trips, capacity):\n    trips = sorted(trips, key=lambda trip: trip[1])     # sort the events\n    riding = []                                         # min-heap of (drop_off, passengers)\n    load = 0\n    for passengers, pick_up, drop_off in trips:\n        while riding and riding[0][0] <= pick_up:       # release what has ended\n            load -= heapq.heappop(riding)[1]\n        heapq.heappush(riding, (drop_off, passengers))  # add the new item\n        load += passengers\n        if load > capacity:                             # check the answer\n            return False\n    return True",
			},
			{
				Title:   "Reorganize String with a cooldown slot",
				Problem: "Rearrange the letters of a string s so that no two equal letters are next to each other, and return any such arrangement, or \"\" if none exists. Example: s = \"aaabbc\" -> \"ababac\" (\"abacab\" is also fine); s = \"aaab\" -> \"\", because three a's need two separators and there is only one b. The greedy rule: always write the letter with the most copies left, except the letter you just wrote.",
				Steps: []WorkedStep{
					{Subgoal: "Count the letters", Explanation: "Which letter goes next depends only on how many copies of each letter are left, so count them first. A letter with many copies is the hardest to place, which is why the greedy rule picks it as early as possible.", Code: "counts = Counter(s)", State: "s = \"aaabbc\": counts = {a: 3, b: 2, c: 1}"},
					{Subgoal: "Build the heap", Explanation: "Put one (-count, letter) entry per distinct letter into a heap. heapq is a min-heap, so negating the count puts the letter with the most copies left on top; equal counts fall back to alphabetical order, which makes the output predictable. heapify does this in O(n).", Code: "heap = [(-count, ch) for ch, count in counts.items()]\nheapq.heapify(heap)\nresult = []\nheld = None", State: "heap = [(-3,'a'), (-2,'b'), (-1,'c')], held = None, result = \"\""},
					{Subgoal: "Read the top", Explanation: "Pop the letter with the most copies left and write one copy of it. It cannot equal the last letter written, because that letter is not in the heap at all: it is sitting in the cooldown slot held.", Code: "while heap:\n    count, ch = heapq.heappop(heap)\n    result.append(ch)", State: "round 1: pop (-3,'a') -> result \"a\", heap = [(-2,'b'), (-1,'c')]. round 2: pop (-2,'b') -> result \"ab\", heap = [(-1,'c')]"},
					{Subgoal: "Release the cooled item", Explanation: "Now that a different letter has been written, the letter in the cooldown slot may appear again, so push it back and let it compete for the next round. Releasing after the pop, not before, is what keeps it from being picked twice in a row.", Code: "    if held:\n        heapq.heappush(heap, held)", State: "round 1: held is None, nothing to release. round 2: push (-2,'a') -> heap = [(-2,'a'), (-1,'c')]"},
					{Subgoal: "Park it while it cools", Explanation: "The letter just written has one copy fewer: count + 1, since counts are negative. If copies remain, park it in held for one round instead of pushing it back; if none remain (count + 1 == 0) it is finished and held becomes None.", Code: "    held = (count + 1, ch) if count + 1 < 0 else None", State: "round 1: held = (-2,'a'). round 2: held = (-1,'b'). round 3: pop (-2,'a'), result \"aba\", release b -> heap = [(-1,'b'), (-1,'c')], held = (-1,'a')"},
					{Subgoal: "Repeat until the heap is empty", Explanation: "Each round writes one letter and moves at most one entry into and out of the heap, so a round costs O(log k) for k distinct letters. The loop stops when nothing is left to choose from.", Code: "while heap:  # rounds 4-6", State: "round 4: pop (-1,'b') -> \"abab\", release a -> heap = [(-1,'a'), (-1,'c')], held = None. round 5: pop (-1,'a') -> \"ababa\", heap = [(-1,'c')]. round 6: pop (-1,'c') -> \"ababac\", heap = [], held = None"},
					{Subgoal: "Check the answer", Explanation: "If the heap ran dry while a letter still had copies in the cooldown slot, there is no other letter left to separate them, so no valid arrangement exists. Picking the most frequent letter first never causes this when an answer exists, so this check is exact.", Code: "if held:\n    return \"\"", State: "\"aaabbc\": held = None, keep going. \"aaab\": pop a -> \"a\", held (-2,'a'); pop b -> \"ab\", release a, held None; pop a -> \"aba\", heap = [], held = (-1,'a') -> return \"\""},
					{Subgoal: "Return the result", Explanation: "Every letter was written and no round wrote the letter that was cooling down, so no two neighbours are equal. Total cost O(n log k), and k is at most 26 for lowercase letters.", Code: "return \"\".join(result)", State: "return \"ababac\""},
				},
				Code: "import heapq\nfrom collections import Counter\n\ndef reorganizeString(s):\n    counts = Counter(s)                                    # count the letters\n    heap = [(-count, ch) for ch, count in counts.items()]  # build the heap\n    heapq.heapify(heap)\n    result = []\n    held = None                       # the letter just written, cooling down\n    while heap:\n        count, ch = heapq.heappop(heap)                    # read the top\n        result.append(ch)\n        if held:                                           # release the cooled item\n            heapq.heappush(heap, held)\n        held = (count + 1, ch) if count + 1 < 0 else None  # park it while it cools\n    if held:                                               # check the answer\n        return \"\"\n    return \"\".join(result)                                # return the result",
			},
		},
		Completions: []CompletionProblem{
			{
				Title:       "The k smallest values, with a max-heap",
				Prompt:      "Return the k smallest numbers in ascending order. Keep a heap of size k, but this time you must evict the largest kept value, so use a max-heap by negation. Fill in the blanks.",
				Code:        "import heapq\n\ndef k_smallest(nums, k):\n    heap = []                                # max-heap of negated values\n    for num in nums:\n        heapq.heappush(heap, ____)           # push the new item\n        if ____:                             # trim to size k\n            heapq.heappop(heap)              # evicts the largest kept value\n    return sorted(____ for x in heap)        # read the heap",
				Blanks:      []string{"-num", "len(heap) > k", "-x"},
				Explanation: "Blank by blank. heapq only builds min-heaps, so pushing -num stores every value negated: the smallest stored negative is the largest original number, which puts exactly the value to evict on top. len(heap) > k is the trim test: push first, then pop once the heap holds k + 1 (popping before pushing could discard a value that deserved to stay). That pop removes the most negative entry, i.e. the largest kept number. What remains are the k smallest, stored as negatives, so the read-out undoes the trick with -x before sorting. For [5, 1, 8, 3] with k = 2: push -5, -1; -8 arrives and the trim pops it right back out (8 is the largest so far); -3 arrives and the trim evicts 5; sorted(-x) over {-3, -1} returns [1, 3].",
			},
			{
				Title:       "The busiest moment of the day",
				Prompt:      "Meetings are [start, end] and are in progress while start <= t < end (a meeting ending at 5 is over for one starting at 5). Return (count, time): the largest number of meetings in progress at once and the earliest time that count is reached, or (0, None) for no meetings. Sweep the meetings by start with a min-heap of the end times of the ones still running. Fill in the blanks.",
				Code:        "import heapq\n\ndef busiest_moment(meetings):\n    # meetings: [start, end], in progress while start <= t < end\n    meetings = sorted(meetings)                         # sort the events\n    running = []                                        # min-heap of end times\n    best_count, best_time = 0, None\n    for start, end in meetings:\n        while ____:                                     # release what has ended\n            heapq.heappop(running)\n        heapq.heappush(running, ____)                   # add the new item\n        if ____:                                        # record the answer\n            best_count, best_time = len(running), start\n    return best_count, best_time",
				Blanks:      []string{"running and running[0] <= start", "end", "len(running) > best_count"},
				Explanation: "Blank by blank. Before a meeting starts, release every meeting that has already ended: the heap holds end times, so the earliest end is on top, and we pop while it is <= start (<= because a meeting ending exactly at start no longer overlaps). The new meeting is now running, so push its end, the key the release step will need later. The heap now holds exactly the meetings in progress at start, so len(running) is the overlap there; the overlap only rises at a start, so checking at starts is enough. Using > rather than >= keeps the first (earliest) time the record count is reached. For [[6,9],[1,5],[3,8],[2,4]], sorted by start: t = 1 -> running [5], best (1, 1); t = 2 -> [4, 5], best (2, 2); t = 3 -> [4, 5, 8], best (3, 3); t = 6 pops 4 and 5, pushes 9 -> [8, 9], 2 is not > 3. Returns (3, 3).",
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
			{
				ID:    "heap-sweep",
				Title: "Sweep Line with a Drop-off Heap",
				Goal:  "Sweep the pick-ups of Car Pooling from left to right, keeping the trips in the car in a min-heap keyed by drop-off.",
				MemoryTips: []string{
					"Sort the events: walk the pick-ups from left to right",
					"Release what has ended: pop while the top's drop-off <= pick-up",
					"Add the new item: push (drop_off, passengers) and add to the load",
					"Check the answer right after boarding: that's the only place the load rises",
				},
				Steps: []StoryboardStep{
					{Cue: "sort", Narration: "Sort the trips by pick-up kilometre. The load in the car can only change where someone boards or leaves, so the sweep line only has to stop at the pick-ups; the heap will tell it about the drop-offs.", VisualHint: "Trips laid out on a timeline", Duration: Duration(2 * time.Second), CodeRef: "heap-sweep: trips = sorted(trips, key=lambda trip: trip[1])"},
					{Cue: "sweep", Narration: "Move the sweep line to the next pick-up. Everything to its left has already happened; the heap holds exactly the trips that boarded earlier and may still be in the car.", VisualHint: "Sweep line jumps to the pick-up", Duration: Duration(2 * time.Second), CodeRef: "heap-sweep: for passengers, pick_up, drop_off in trips"},
					{Cue: "release", Narration: "Look at the top of the heap: the trip in the car with the earliest drop-off. If it gets off at or before this kilometre, pop it and free its seats, then look again. Stop at the first top that is still riding; every trip below it gets off even later.", VisualHint: "Expired top pops off the tree", Duration: Duration(3 * time.Second), CodeRef: "heap-sweep: while riding and riding[0][0] <= pick_up"},
					{Cue: "push", Narration: "The new trip boards: push (drop_off, passengers) so the release step can find it when the sweep passes its drop-off, and add its passengers to the load. The new entry sifts up if it gets off sooner than its parent.", VisualHint: "New node joins the heap", Duration: Duration(2 * time.Second), CodeRef: "heap-sweep: heapq.heappush(riding, (drop_off, passengers))"},
					{Cue: "check", Narration: "Compare the load with the capacity right after boarding. Between pick-ups the load can only fall, so if it fits at every pick-up it fits everywhere.", VisualHint: "Load meter against the capacity mark", Duration: Duration(2 * time.Second), CodeRef: "heap-sweep: if load > capacity"},
					{Cue: "done", Narration: "Every pick-up passed the check, so the car never overflowed. Each trip was pushed once and popped at most once: O(n log n) with the sort.", VisualHint: "All trips done", Duration: Duration(2 * time.Second), CodeRef: "heap-sweep: return True"},
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
			{
				Prompt:      "In tallest_active (the sweep-line concept example, a max-heap of (-value, end)), the heap at t = 7 holds (-6, 9) on top and (-5, 6) below it, although the interval with value 5 ended at t = 6. Why is it fine to leave it there?",
				Choices:     []string{"It isn't fine: every ended interval must be removed as soon as the sweep passes its end, or the answer is wrong", "Only the top is ever read, and the release loop checks the top before each read, so a buried ended interval is popped the moment it would surface", "Heaps automatically drop items whose end has passed", "It is fine only because 5 < 6; if the ended interval had a larger value, the answer would be wrong"},
				Answer:      1,
				Explanation: "Lazy deletion relies on the fact that the answer comes only from heap[0], and the release loop ('while the top has ended, pop') runs before every read. A buried ended item never influences a read; when it reaches the top it is popped. Option 1 would need an O(n) search per removal. Heaps know nothing about time (option 3). Option 4 misses that a larger ended interval would be on top and popped by the release loop first, so it can't be read either.",
			},
			{
				Prompt:      "Trace: tallest_active from the sweep-line concept (push what has started, pop ended tops while end <= t, read -top or 0). What does tallest_active([[0, 4, 3], [1, 3, 7], [2, 6, 5]], [1, 3, 4, 6]) return?",
				Choices:     []string{"[7, 5, 5, 0]", "[7, 7, 5, 0]", "[7, 5, 3, 0]", "[7, 5, 5, 3]"},
				Answer:      0,
				Explanation: "t = 1: push (-3, end 4) and (-7, end 3); top is 7 -> 7. t = 3: push (-5, end 6); the top 7 has end 3 <= 3, pop it; the top 5 is active -> 5. t = 4: the top 5 (end 6) is active -> 5; the 3 ended at 4 but is buried, so it stays. t = 6: the top 5 has end 6 <= 6, pop; the 3 surfaces with end 4 <= 6, pop; the heap is empty -> 0. [7, 7, 5, 0] reads the top before popping (or treats the end as inclusive), and [7, 5, 5, 3] forgets that buried ended items must be popped when they surface.",
			},
		},
	}
}
