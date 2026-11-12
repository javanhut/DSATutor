package practice

// HeapProblems contains all heap-priority-queue category problems, in ladder
// order. The subgoal labels (build the heap, push the new item, trim to size
// k, read the top) are the ones chapter 11 teaches.
//
// Ladder: 1 Last Stone Weight (max-heap by negation) -> 2 Kth Largest in a
// Stream (size-k min-heap) -> 3 Minimum Cost to Connect Sticks (pop two, push
// the sum: greedy min-heap) -> 4 K Closest Points (tuple keys, size-k max-heap)
// -> 5 Top K Frequent Words (two-part key with an alphabetical tie-break) ->
// 6 Task Scheduler (heap plus a cooling queue) -> 7 Design Twitter (k-way
// merge of sorted lists) -> 8 Find K Pairs with Smallest Sums (k-way merge over
// implicit rows) -> 9 Find Median from Data Stream (two balanced heaps) ->
// 10 Meeting Rooms III (free and busy heaps, delayed starts) -> 11 The Skyline
// Problem (sweep with a max-heap and lazy deletion).
var HeapProblems = []*Problem{
	{
		ID:              "last-stone-weight",
		Number:          1046,
		Order:           1,
		Title:           "Last Stone Weight",
		Slug:            "last-stone-weight",
		Difficulty:      "Easy",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Heap (Priority Queue)"},
		RelatedChapters: []int{11, 9},
		Description: `You have a pile of stones; ` + "`stones[i]`" + ` is the weight of stone ` + "`i`" + `.

Each turn, take the **two heaviest** stones, with weights ` + "`x <= y`" + `, and smash them together:
- if ` + "`x == y`" + `, both stones are destroyed;
- otherwise the lighter stone is destroyed and the heavier one now weighs ` + "`y - x`" + ` and goes back into the pile.

Repeat until at most one stone is left. Return the weight of that last stone, or ` + "`0`" + ` if no stones are left.`,
		Constraints: []string{
			"1 <= stones.length <= 30",
			"1 <= stones[i] <= 1000",
		},
		Examples: []Example{
			{Input: "stones = [2,7,4,1,8,1]", Output: "1", Explanation: "Smash 8 and 7 -> 1, pile [2,4,1,1,1]. Smash 4 and 2 -> 2, pile [2,1,1,1]. Smash 2 and 1 -> 1, pile [1,1,1]. Smash 1 and 1 -> nothing, pile [1]. The last stone weighs 1."},
			{Input: "stones = [1]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"stones": []int{2, 7, 4, 1, 8, 1}}, Expected: 1, Hidden: false, Note: "the example: smash 8/7, 4/2, 2/1, 1/1"},
			{Input: map[string]interface{}{"stones": []int{1}}, Expected: 1, Hidden: false, Note: "one stone is already the last stone"},
			{Input: map[string]interface{}{"stones": []int{2, 2}}, Expected: 0, Hidden: false, Note: "equal stones destroy each other: return 0"},
			{Input: map[string]interface{}{"stones": []int{3, 7, 2}}, Expected: 2, Hidden: true, Note: "the leftover must go back into the heap"},
			{Input: map[string]interface{}{"stones": []int{10, 4, 2, 10}}, Expected: 2, Hidden: true, Note: "the two heaviest are equal"},
			{Input: map[string]interface{}{"stones": []int{9, 3, 2, 10}}, Expected: 0, Hidden: true, Note: "everything cancels out"},
			{Input: map[string]interface{}{"stones": []int{1, 3}}, Expected: 2, Hidden: true, Note: "unsorted pair: heaviest minus lighter"},
			{Input: map[string]interface{}{"stones": []int{1000, 1, 1, 1, 1, 1}}, Expected: 995, Hidden: true, Note: "one giant stone: catches using a min-heap without negating"},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def lastStoneWeight(stones):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Every turn you need the two largest values of a collection that keeps changing (a new stone can go back in). Sorting once isn't enough. Which structure gives you the largest item quickly, even after insertions?"},
			{Level: 2, Type: "insight", Content: "A max-heap: pop gives the largest in O(log n) and push inserts in O(log n). Python's heapq is a min-heap only, so store negated weights: the smallest negative number is the heaviest stone. Build it with heapify in O(n), then repeat: pop two, push the difference if it isn't zero.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. Build the heap: negate every weight and heapify.\n2. Read the top: while at least two stones remain, pop the heaviest and the second heaviest (un-negate them).\n3. Push the new item: if they differ, push -(heaviest - second).\n4. Return the last weight (un-negated), or 0 if the heap is empty."},
			{Level: 4, Type: "skeleton", Content: "import heapq\n\ndef lastStoneWeight(stones):\n    heap = [____ for s in stones]\n    heapq.heapify(heap)\n    while len(heap) > 1:\n        heaviest = ____\n        second = -heapq.heappop(heap)\n        if heaviest != second:\n            heapq.heappush(heap, ____)\n    return ____ if heap else 0"},
		},
		Solution: Solution{
			Code: `import heapq


def lastStoneWeight(stones):
    heap = [-s for s in stones]          # negate: heapq is a min-heap
    heapq.heapify(heap)                  # build the heap in O(n)
    while len(heap) > 1:
        heaviest = -heapq.heappop(heap)  # read the top twice
        second = -heapq.heappop(heap)
        if heaviest != second:
            heapq.heappush(heap, -(heaviest - second))  # push the leftover back
    return -heap[0] if heap else 0`,
			Explanation: "The first rung of the heap ladder: a plain function that only needs a priority queue used directly. It practises two habits from chapter 11 that every later problem uses: heapq is a min-heap, so negate values to get a max-heap, and heapify builds the heap in O(n). After that, each smash is two pops and at most one push.",
			BruteForce:  "Sort the pile every turn and take the last two. Each sort is O(n log n) and there can be up to n - 1 turns, so O(n^2 log n). The waste: after one smash only one stone changed, yet the whole list is re-sorted from scratch.",
			KeyInsight:  "Only the maximum matters each turn, and only one item is inserted. A heap gives exactly that: the top is the maximum, and pop or push repairs the order along one root-to-leaf path in O(log n), without re-sorting everything.",
			WhyItWorks:  "Invariant: the heap holds exactly the negated weights of the stones still in the pile, and it satisfies the heap property (each parent <= its children), so heap[0] is the most negative value, i.e. the heaviest stone. Each loop iteration pops the two heaviest stones, which is what the rules require, and pushes the difference back if one survives, so the invariant still holds. Every turn removes at least one stone, so there are at most n - 1 turns of O(log n) each.",
			CommonMistakes: []string{
				"Using heapq without negating. heappop then returns the lightest stones, and the game is simulated backwards; [1000,1,1,1,1,1] gives 0 instead of 995.",
				"Forgetting to un-negate when reading or returning. You'd return -1 instead of 1. Negate on the way in and on the way out.",
				"Returning heap[0] without checking the heap is empty. When the last two stones are equal there is nothing left: return 0.",
			},
			TimeComplexity:  "O(n log n): heapify is O(n), then at most n - 1 turns of O(log n)",
			SpaceComplexity: "O(n) for the heap",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the heap", Explanation: "Negate so the min-heap behaves as a max-heap, then heapify in O(n).", CodeSnippet: "heap = [-s for s in stones]\nheapq.heapify(heap)", LineStart: 5, LineEnd: 6},
				{Title: "Read the top", Explanation: "Pop twice to get the two heaviest stones, un-negating each.", CodeSnippet: "heaviest = -heapq.heappop(heap)\nsecond = -heapq.heappop(heap)", LineStart: 8, LineEnd: 9},
				{Title: "Push the new item", Explanation: "If one stone survives, it goes back into the pile with its reduced weight (negated again).", CodeSnippet: "if heaviest != second:\n    heapq.heappush(heap, -(heaviest - second))", LineStart: 10, LineEnd: 11},
				{Title: "Return the answer", Explanation: "At most one stone is left; an empty heap means 0.", CodeSnippet: "return -heap[0] if heap else 0", LineStart: 12, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: sort every turn", Code: `def lastStoneWeight(stones):
    stones = list(stones)
    while len(stones) > 1:
        stones.sort()                    # re-sort every round: O(n log n) each time
        heaviest = stones.pop()
        second = stones.pop()
        if heaviest != second:
            stones.append(heaviest - second)
    return stones[0] if stones else 0`, TimeComplexity: "O(n^2 log n)", SpaceComplexity: "O(n)", Explanation: "Re-sorts the pile each turn and takes the two largest from the end. Easy to get right, but it redoes the whole ordering after each small change."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does storing -weight in a min-heap give you the heaviest stone first?", Answer: "Negation reverses order: if a > b then -a < -b. The min-heap's top is the smallest stored value, which is the negation of the largest weight. So popping and negating again returns the heaviest stone."},
			{Prompt: "Why is heapify O(n) and not O(n log n)?", Answer: "heapify sifts nodes down starting from the bottom. Half the nodes are leaves and don't move; a quarter move at most one level; an eighth at most two, and so on. The sum n/4 * 1 + n/8 * 2 + ... is less than n, so the total work is O(n). Pushing n items one by one would be O(n log n)."},
		},
	},
	{
		ID:              "kth-largest-element-in-stream",
		Number:          703,
		Order:           2,
		Title:           "Kth Largest Element in a Stream",
		Slug:            "kth-largest-element-in-a-stream",
		Difficulty:      "Easy",
		Category:        "heap-priority-queue",
		Tags:            []string{"Design", "Heap (Priority Queue)", "Data Stream"},
		RelatedChapters: []int{11, 9},
		Description: `Design a class that keeps track of the **k-th largest** number seen so far in a stream of numbers. "k-th largest" means k-th in sorted order from the top, counting duplicates (so in [5,5,4] the 2nd largest is 5).

Implement the ` + "`KthLargest`" + ` class:
- ` + "`KthLargest(k, nums)`" + ` stores ` + "`k`" + ` and the initial numbers ` + "`nums`" + ` (which may be empty).
- ` + "`add(val)`" + ` adds ` + "`val`" + ` to the stream and returns the current k-th largest number.

It is guaranteed that whenever ` + "`add`" + ` returns, at least ` + "`k`" + ` numbers have been seen.`,
		Constraints: []string{
			"1 <= k <= 10^4",
			"0 <= nums.length <= 10^4",
			"-10^4 <= nums[i], val <= 10^4",
			"At most 10^4 calls will be made to add",
		},
		Examples: []Example{
			{Input: "[\"KthLargest\",\"add\",\"add\",\"add\",\"add\",\"add\"]\n[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]", Output: "[null,4,5,5,8,8]", Explanation: "k = 3. After adding 3 the stream is [2,3,4,5,8]: the 3rd largest is 4. After 5: [2,3,4,5,5,8] -> 5. After 10 -> 5. After 9 -> 8. After 4 -> 8."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"KthLargest", "add", "add", "add", "add", "add"}, "values": []interface{}{[]interface{}{3, []int{4, 5, 8, 2}}, []int{3}, []int{5}, []int{10}, []int{9}, []int{4}}}, Expected: []interface{}{nil, 4, 5, 5, 8, 8}, Hidden: false, Note: "the example: k = 3"},
			{Input: map[string]interface{}{"operations": []string{"KthLargest", "add", "add", "add"}, "values": []interface{}{[]interface{}{1, []interface{}{}}, []int{-3}, []int{-2}, []int{-4}}}, Expected: []interface{}{nil, -3, -2, -2}, Hidden: false, Note: "k = 1 starting empty: the answer is the running maximum"},
			{Input: map[string]interface{}{"operations": []string{"KthLargest", "add", "add", "add", "add", "add"}, "values": []interface{}{[]interface{}{2, []int{0}}, []int{-1}, []int{1}, []int{-2}, []int{-4}, []int{3}}}, Expected: []interface{}{nil, -1, 0, 0, 0, 1}, Hidden: true, Note: "starts with fewer than k numbers"},
			{Input: map[string]interface{}{"operations": []string{"KthLargest", "add", "add", "add"}, "values": []interface{}{[]interface{}{2, []int{5, 5, 5}}, []int{5}, []int{1}, []int{6}}}, Expected: []interface{}{nil, 5, 5, 5}, Hidden: true, Note: "duplicates count separately (not kth distinct)"},
			{Input: map[string]interface{}{"operations": []string{"KthLargest", "add", "add"}, "values": []interface{}{[]interface{}{4, []int{7, 1, 9, 3, 8, 2}}, []int{0}, []int{10}}}, Expected: []interface{}{nil, 3, 7}, Hidden: true, Note: "constructor gets more than k numbers and must trim"},
			{Input: map[string]interface{}{"operations": []string{"KthLargest", "add", "add", "add", "add"}, "values": []interface{}{[]interface{}{3, []int{-10, -20, -30}}, []int{-25}, []int{-5}, []int{-1}, []int{-15}}}, Expected: []interface{}{nil, -25, -20, -10, -10}, Hidden: true, Note: "all negative values"},
		},
		TimeComplexity:  "O(log k) per add",
		SpaceComplexity: "O(k)",
		StarterCode:     "class KthLargest:\n    def __init__(self, k: int, nums: List[int]):\n        pass\n\n    def add(self, val: int) -> int:\n        pass",
		Harness:         Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "To answer 'what is the 3rd largest?' do you need to remember the numbers smaller than the 3rd largest at all? What is the smallest set of numbers you need to keep?"},
			{Level: 2, Type: "insight", Content: "Keep only the k largest numbers seen so far, in a min-heap. The smallest of those k is exactly the k-th largest overall, and it sits at heap[0]. When a new number arrives, push it; if the heap now has k + 1 items, pop the smallest. This is the 'keep a heap of size k' pattern from chapter 11.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. Build the heap: in the constructor, heapify the initial nums.\n2. Trim to size k: pop while it holds more than k items.\n3. Push the new item: in add, push val.\n4. Trim to size k: if the size is now k + 1, pop once.\n5. Read the top: return heap[0]."},
			{Level: 4, Type: "skeleton", Content: "import heapq\n\nclass KthLargest:\n    def __init__(self, k: int, nums: List[int]):\n        self.k = k\n        self.heap = list(nums)\n        heapq.heapify(self.heap)\n        while ____:\n            heapq.heappop(self.heap)\n\n    def add(self, val: int) -> int:\n        heapq.heappush(self.heap, val)\n        if ____:\n            heapq.heappop(self.heap)\n        return ____"},
		},
		Solution: Solution{
			Code: `import heapq


class KthLargest:
    def __init__(self, k: int, nums: List[int]):
        self.k = k
        self.heap = list(nums)
        heapq.heapify(self.heap)          # build the heap in O(n)
        while len(self.heap) > k:         # trim to size k
            heapq.heappop(self.heap)

    def add(self, val: int) -> int:
        heapq.heappush(self.heap, val)    # push the new item
        if len(self.heap) > self.k:       # trim to size k
            heapq.heappop(self.heap)
        return self.heap[0]               # read the top: smallest of the k largest`,
			Explanation: "Builds on Last Stone Weight (problem 1), which used the heap's top directly. Here the new idea is the size-k heap: instead of keeping everything, keep only the k largest values in a min-heap. Its top is then the k-th largest, and each new value costs one push and at most one pop on a heap of size k.",
			BruteForce:  "Keep every number in a list; on each add, sort it (or insert in sorted position) and read the element k from the end. Sorting is O(n log n) per add; bisect.insort is O(n) because it shifts elements. With 10^4 adds on a growing list that is up to about 10^8 element moves. The waste: numbers below the k-th largest are stored and shifted forever, though they can never be the answer again.",
			KeyInsight:  "Once a number has k larger-or-equal numbers above it, it can never become the k-th largest again (numbers are only added). So only the top k matter, and among them we only ever ask for the smallest: a min-heap of size k.",
			WhyItWorks:  "Invariant: after each call the heap holds the k largest numbers seen so far (or all of them, if fewer than k have been seen). Adding val: the new top-k is the old top-k plus val, minus the smallest of those k + 1, which is exactly what push-then-pop-the-minimum does. The k-th largest overall is the smallest member of the top k, which the min-heap keeps at heap[0]. Each add touches a heap of size at most k + 1, so it is O(log k).",
			CommonMistakes: []string{
				"Using a max-heap of everything and popping k - 1 times per query. That's O(k log n) per add and destroys the heap. A min-heap of size k answers in O(1) after an O(log k) update.",
				"Assuming nums has at least k elements in the constructor. It may have k - 1 (even 0). Only trim while the size is greater than k; never read heap[0] in the constructor.",
				"Popping before pushing ('if len == k: pop; push'). If val is smaller than everything in the heap, you've thrown away a top-k value and kept val. Push first, then trim (or use heapq.heappushpop).",
			},
			TimeComplexity:  "O(n log n) to build (heapify, then trim), O(log k) per add",
			SpaceComplexity: "O(k)",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the heap", Explanation: "heapify the initial numbers in O(n).", CodeSnippet: "self.heap = list(nums)\nheapq.heapify(self.heap)", LineStart: 7, LineEnd: 8},
				{Title: "Trim to size k", Explanation: "Pop the smallest until only the k largest remain. Each pop removes a number that can never be the answer again.", CodeSnippet: "while len(self.heap) > k:\n    heapq.heappop(self.heap)", LineStart: 9, LineEnd: 10},
				{Title: "Push the new item", Explanation: "The new value might belong in the top k, so it goes in first.", CodeSnippet: "heapq.heappush(self.heap, val)", LineStart: 13, LineEnd: 13},
				{Title: "Read the top", Explanation: "After trimming back to k items, the smallest of the k largest is the answer.", CodeSnippet: "if len(self.heap) > self.k:\n    heapq.heappop(self.heap)\nreturn self.heap[0]", LineStart: 14, LineEnd: 16},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: sorted list", Code: `import bisect


class KthLargest:
    def __init__(self, k, nums):
        self.k = k
        self.sorted_nums = sorted(nums)

    def add(self, val):
        bisect.insort(self.sorted_nums, val)   # O(n) shift to keep the list sorted
        return self.sorted_nums[-self.k]`, TimeComplexity: "O(n) per add", SpaceComplexity: "O(n)", Explanation: "Keeps every number in sorted order with bisect.insort and reads index -k. Finding the position is O(log n), but inserting shifts up to n elements."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why a min-heap, when we want the k-th *largest*?", Answer: "We keep the k largest numbers, and the answer is the smallest of them. The question we ask most is 'which kept number is weakest?', both to answer and to decide what to evict. A min-heap puts exactly that number on top."},
			{Prompt: "What changes if the stream could also remove numbers?", Answer: "A size-k heap no longer works, because a removed number might be in the top k and an evicted number might have to come back. You'd need all values, e.g. a sorted container or two heaps with lazy deletion (mark removed values and skip them when they reach the top)."},
		},
	},
	{
		ID:              "minimum-cost-to-connect-sticks",
		Number:          1167,
		Order:           3,
		Title:           "Minimum Cost to Connect Sticks",
		Slug:            "minimum-cost-to-connect-sticks",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Greedy", "Heap (Priority Queue)", "Google", "Google L4", "Google Phone Screen"},
		RelatedChapters: []int{11, 17},
		Description: `You have some sticks with positive integer lengths; ` + "`sticks[i]`" + ` is the length of stick ` + "`i`" + `.

You may join any two sticks of lengths ` + "`x`" + ` and ` + "`y`" + ` into one stick of length ` + "`x + y`" + `. Each join **costs** ` + "`x + y`" + `, the length of the new stick. The new stick can be joined again later, like any other.

Keep joining until exactly one stick is left. Return the **minimum total cost** of all the joins. If you start with a single stick, no join is needed and the cost is ` + "`0`" + `.`,
		Constraints: []string{
			"1 <= sticks.length <= 10^4",
			"1 <= sticks[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "sticks = [2,4,3]", Output: "14", Explanation: "Join 2 and 3 for a cost of 5; the sticks are now [5,4]. Join 5 and 4 for a cost of 9; one stick of length 9 is left. Total 5 + 9 = 14. Joining 4 and 3 first would cost 7 + 9 = 16."},
			{Input: "sticks = [1,8,3,5]", Output: "30", Explanation: "Join 1 and 3 (cost 4), then 4 and 5 (cost 9), then 9 and 8 (cost 17): 4 + 9 + 17 = 30."},
			{Input: "sticks = [5]", Output: "0", Explanation: "There is already only one stick."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"sticks": []int{2, 4, 3}}, Expected: 14, Hidden: false, Note: "the example: join 2 and 3 first"},
			{Input: map[string]interface{}{"sticks": []int{1, 8, 3, 5}}, Expected: 30, Hidden: false, Note: "the joined stick keeps being the shortest"},
			{Input: map[string]interface{}{"sticks": []int{5}}, Expected: 0, Hidden: false, Note: "a single stick costs nothing"},
			{Input: map[string]interface{}{"sticks": []int{3, 3}}, Expected: 6, Hidden: true, Note: "two sticks: exactly one join"},
			{Input: map[string]interface{}{"sticks": []int{1, 1, 1, 1}}, Expected: 8, Hidden: true, Note: "all equal: pair them up (2 + 2), then join the pairs (4)"},
			{Input: map[string]interface{}{"sticks": []int{4, 4, 6, 6}}, Expected: 40, Hidden: true, Note: "sorting once and chaining left to right gives 42; the joined 8 must wait while 6 and 6 join"},
			{Input: map[string]interface{}{"sticks": []int{10, 1, 2}}, Expected: 16, Hidden: true, Note: "joining the longest sticks first gives 25"},
			{Input: map[string]interface{}{"sticks": []int{10000, 10000, 10000}}, Expected: 50000, Hidden: true, Note: "maximum lengths"},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def connectSticks(sticks):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "A stick that is joined early becomes part of every later stick it ends up in, so its length is paid again and again. Which sticks would you rather pay for many times, and which only once?"},
			{Level: 2, Type: "insight", Content: "Always join the two shortest sticks currently in the pile. A min-heap gives you them in O(log n): pop two, add their sum to the cost, push the sum back. It is Last Stone Weight with a min-heap instead of a max-heap, and the new stick is a sum instead of a difference.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. Build the heap: heapify all the stick lengths into a min-heap.\n2. Read the top: while more than one stick is left, pop the two shortest.\n3. Pay for the join: add their sum to the total.\n4. Push the new item: push the sum back into the heap.\n5. Return the total (0 if there was only one stick)."},
			{Level: 4, Type: "skeleton", Content: "import heapq\n\ndef connectSticks(sticks):\n    heap = list(sticks)\n    heapq.heapify(heap)\n    total = 0\n    while ____:\n        first = heapq.heappop(heap)\n        second = heapq.heappop(heap)\n        joined = ____\n        total += ____\n        heapq.heappush(heap, ____)\n    return total"},
		},
		Solution: Solution{
			Code: `import heapq


def connectSticks(sticks):
    heap = list(sticks)                    # build the heap (min-heap)
    heapq.heapify(heap)
    total = 0
    while len(heap) > 1:
        first = heapq.heappop(heap)        # read the top: the two shortest
        second = heapq.heappop(heap)
        joined = first + second
        total += joined                    # pay for this join
        heapq.heappush(heap, joined)       # push the new item
    return total`,
			Explanation: "Builds on Last Stone Weight (problem 1): pop the top two, combine them, push the result back, repeat until one is left. Two things change: we want the two *shortest* sticks, so Python's min-heap is used as-is with no negation, and every join adds to a running cost. The new idea is the greedy argument for why the two shortest is always the right pair.",
			BruteForce:  "Try every possible pair to join at every step and recurse on the resulting pile, keeping the cheapest total. With n sticks there are about n^2/2 choices at the first step, (n-1)^2/2 at the next and so on, so the search grows faster than n!. The waste: most of those orders join long sticks early, which makes their length be paid again in every later join; one fixed rule picks a best pair every time.",
			KeyInsight:  "Each stick's length is paid once for every join it takes part in, directly or inside a bigger stick. To keep the total small, the longest sticks should take part in the fewest joins, so the two shortest sticks are always the ones to join now. A min-heap hands them out in O(log n).",
			WhyItWorks:  "Draw the joins as a tree: original sticks are leaves, and each join is a parent whose length is the sum of its children. A stick at depth d is paid d times, so the total cost is the sum of length * depth over the leaves. In some optimal tree the two shortest sticks are siblings at the deepest level: if a longer stick sat deeper than a shorter one, swapping them would not raise the cost. So joining the two shortest first is safe, and what remains is the same problem with one fewer stick (the joined stick replaces the two). Repeating this argument at every step shows the greedy total is optimal. This is the Huffman coding argument.",
			CommonMistakes: []string{
				"Sorting once and then joining left to right, always adding the next stick to the growing stick. On [4, 4, 6, 6] that pays 8 + 14 + 20 = 42 instead of 8 + 12 + 20 = 40: the joined stick may no longer be among the two shortest. Push it back into a heap so it competes with the rest.",
				"Using negation out of habit from Last Stone Weight. That pops the two longest sticks, which is the worst choice ([10, 1, 2] costs 25 instead of 16). Here we want the smallest, which is Python's default.",
				"Forgetting to push the joined stick back, or returning the length of the last stick instead of the running total. The cost is the sum of every join, not the final length.",
			},
			TimeComplexity:  "O(n log n): n - 1 joins, each with two pops and one push",
			SpaceComplexity: "O(n) for the heap",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the heap", Explanation: "Copy the lengths (so the caller's list is untouched) and heapify in O(n). The shortest stick is now on top.", CodeSnippet: "heap = list(sticks)\nheapq.heapify(heap)", LineStart: 5, LineEnd: 6},
				{Title: "Read the top", Explanation: "Pop twice to get the two shortest sticks in the pile, which may include sticks we made earlier.", CodeSnippet: "first = heapq.heappop(heap)\nsecond = heapq.heappop(heap)", LineStart: 9, LineEnd: 10},
				{Title: "Pay for the join", Explanation: "The join costs the length of the new stick.", CodeSnippet: "joined = first + second\ntotal += joined", LineStart: 11, LineEnd: 12},
				{Title: "Push the new item", Explanation: "The joined stick goes back into the pile and competes with the others for the next join.", CodeSnippet: "heapq.heappush(heap, joined)", LineStart: 13, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{Name: "Re-sort every round", Code: `def connectSticks(sticks):
    pile = list(sticks)
    total = 0
    while len(pile) > 1:
        pile.sort()                        # find the two shortest by sorting again
        joined = pile[0] + pile[1]
        total += joined
        pile = pile[2:] + [joined]
    return total`, TimeComplexity: "O(n^2 log n)", SpaceComplexity: "O(n)", Explanation: "The same greedy, but it re-sorts the whole pile before every join just to find the two shortest. Each round costs O(n log n), and there are n - 1 rounds; the heap gets the same two sticks in O(log n)."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does a stick joined early cost more overall than one joined late?", Answer: "Its length becomes part of the new stick, and that stick's length is paid again at every later join it is part of. A stick at depth d in the join tree is paid d times. Long sticks should therefore be joined as late as possible, which is exactly what joining the two shortest first does."},
			{Prompt: "If the sticks were already sorted, could you avoid the heap?", Answer: "Yes. Keep the sorted sticks in one queue and the joined sticks in a second queue. Joined sticks come out in non-decreasing order, so the second queue stays sorted too. The two shortest are always among the fronts of the two queues, giving O(n) after the sort. This is the classic two-queue Huffman construction."},
		},
	},
	{
		ID:              "k-closest-points-to-origin",
		Number:          973,
		Order:           4,
		Title:           "K Closest Points to Origin",
		Slug:            "k-closest-points-to-origin",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Math", "Geometry", "Sorting", "Heap (Priority Queue)"},
		RelatedChapters: []int{11, 9},
		Description: `You are given a list of points on a plane, ` + "`points[i] = [x, y]`" + `, and an integer ` + "`k`" + `. Return the ` + "`k`" + ` points closest to the origin ` + "`(0, 0)`" + `.

Distance is ordinary straight-line (Euclidean) distance: ` + "`sqrt(x^2 + y^2)`" + `. You may return the points in any order. The input guarantees the answer is unique (no tie at the k-th place).`,
		Constraints: []string{
			"1 <= k <= points.length <= 10^4",
			"-10^4 <= x, y <= 10^4",
		},
		Examples: []Example{
			{Input: "points = [[1,3],[-2,2]], k = 1", Output: "[[-2,2]]", Explanation: "(1,3) is at distance sqrt(10); (-2,2) is at sqrt(8), which is smaller."},
			{Input: "points = [[3,3],[5,-1],[-2,4]], k = 2", Output: "[[3,3],[-2,4]]", Explanation: "Squared distances are 18, 26 and 20; the two smallest are 18 and 20."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"points": [][]int{{1, 3}, {-2, 2}}, "k": 1}, Expected: [][]int{{-2, 2}}, Hidden: false, Note: "the example: 8 < 10"},
			{Input: map[string]interface{}{"points": [][]int{{3, 3}, {5, -1}, {-2, 4}}, "k": 2}, Expected: [][]int{{-2, 4}, {3, 3}}, Hidden: false, Note: "keep two of three"},
			{Input: map[string]interface{}{"points": [][]int{{0, 1}}, "k": 1}, Expected: [][]int{{0, 1}}, Hidden: true, Note: "a single point"},
			{Input: map[string]interface{}{"points": [][]int{{1, 1}, {2, 2}, {3, 3}}, "k": 3}, Expected: [][]int{{3, 3}, {1, 1}, {2, 2}}, Hidden: true, Note: "k equals the number of points"},
			{Input: map[string]interface{}{"points": [][]int{{0, 0}, {5, 5}, {-1, 0}, {10, -10}}, "k": 2}, Expected: [][]int{{-1, 0}, {0, 0}}, Hidden: true, Note: "the origin itself is distance 0"},
			{Input: map[string]interface{}{"points": [][]int{{-5, -5}, {4, -3}, {-6, 0}, {1, 7}, {2, -2}}, "k": 3}, Expected: [][]int{{-6, 0}, {4, -3}, {2, -2}}, Hidden: true, Note: "negative coordinates: squaring removes the sign"},
			{Input: map[string]interface{}{"points": [][]int{{10000, 10000}, {-10000, 9999}, {1, -1}}, "k": 1}, Expected: [][]int{{1, -1}}, Hidden: true, Note: "large coordinates near the limits"},
		},
		TimeComplexity:  "O(n log k)",
		SpaceComplexity: "O(k)",
		StarterCode:     "def kClosest(points, k):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In Kth Largest you kept the k largest numbers. Here you want the k smallest distances. Which number should sit on top of the heap so you know what to throw away when a new point arrives?"},
			{Level: 2, Type: "insight", Content: "Keep a heap of size k holding the k closest points so far. To evict, you need the farthest of them on top, so this is a max-heap by distance: push (-distance, x, y). When the heap grows to k + 1, pop: that removes the farthest. Compare squared distances x*x + y*y; the square root doesn't change the order.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. For each point, compute the squared distance.\n2. Push the new item: push (-dist, x, y).\n3. Trim to size k: if the heap has k + 1 entries, pop (the farthest).\n4. Read the heap: return the points left in it, in any order."},
			{Level: 4, Type: "skeleton", Content: "import heapq\n\ndef kClosest(points, k):\n    heap = []\n    for x, y in points:\n        dist = ____\n        heapq.heappush(heap, (____, x, y))\n        if ____:\n            heapq.heappop(heap)\n    return [[x, y] for _, x, y in heap]"},
		},
		Solution: Solution{
			Code: `import heapq


def kClosest(points, k):
    heap = []                              # max-heap of the k closest so far
    for x, y in points:
        dist = x * x + y * y               # squared distance: same order, no sqrt
        heapq.heappush(heap, (-dist, x, y))  # push the new item (negated for max)
        if len(heap) > k:                  # trim to size k
            heapq.heappop(heap)            # drops the farthest of the k + 1
    return [[x, y] for _, x, y in heap]`,
			Explanation: "Uses the same 'heap of size k' pattern as Kth Largest (problem 2), flipped: to keep the k *smallest* distances we need the *largest* of them on top, so we use a max-heap via negation, as in Last Stone Weight (problem 1). The new idea is heap entries that are tuples, ordered by their first element (the negated distance), with the point carried along.",
			BruteForce:  "Sort all n points by distance and take the first k: O(n log n). It's simple and fine for 10^4 points, but it orders all n points when we only need to know which k are smallest; when k is much smaller than n, that's wasted work.",
			KeyInsight:  "To keep the k best by 'smallest', evict the worst of the kept ones, which is the largest. So the eviction heap is a max-heap of size k, and each point costs O(log k) instead of contributing to an O(n log n) sort.",
			WhyItWorks:  "Invariant: after processing a prefix of points, the heap holds the k closest points of that prefix (or all of them if there are fewer than k), with the farthest of those on top (most negative -dist). When a new point arrives, the k closest of the prefix plus the new point are the old k plus the new one minus the farthest of those k + 1: push then pop. Squared distance preserves order because sqrt is increasing for non-negative numbers. At the end the heap holds the k closest overall.",
			CommonMistakes: []string{
				"Pushing +dist into a min-heap of size k. Popping then evicts the closest point, and you return the k farthest. To keep the smallest values, evict the largest: negate.",
				"Pushing [x, y] lists and relying on tuple comparison for ties. It works in Python, but if you push (dist, point_object) and two distances tie, Python compares the objects and can crash. Put only comparable values (numbers) after the key.",
				"Using math.sqrt and floats. Not wrong, but slower, and float rounding can blur near-ties. Squared integers order the points exactly.",
			},
			TimeComplexity:  "O(n log k): n pushes and at most n pops on a heap of size k + 1",
			SpaceComplexity: "O(k)",
			Walkthrough: []WalkthroughStep{
				{Title: "Compute the key", Explanation: "Squared distance orders points the same way as real distance and stays an exact integer.", CodeSnippet: "dist = x * x + y * y", LineStart: 7, LineEnd: 7},
				{Title: "Push the new item", Explanation: "Negate so the farthest kept point is on top of Python's min-heap.", CodeSnippet: "heapq.heappush(heap, (-dist, x, y))", LineStart: 8, LineEnd: 8},
				{Title: "Trim to size k", Explanation: "With k + 1 points kept, the top is the farthest of them; it can't be in the answer.", CodeSnippet: "if len(heap) > k:\n    heapq.heappop(heap)", LineStart: 9, LineEnd: 10},
				{Title: "Read the heap", Explanation: "The k points left are the answer; any order is accepted.", CodeSnippet: "return [[x, y] for _, x, y in heap]", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{Name: "Sort by distance", Code: `def kClosest(points, k):
    by_distance = sorted(points, key=lambda p: p[0] * p[0] + p[1] * p[1])
    return by_distance[:k]`, TimeComplexity: "O(n log n)", SpaceComplexity: "O(n)", Explanation: "Sort everything by squared distance and slice. Shorter, and perfectly fine when k is close to n; the heap wins when k is small. Quickselect (average O(n)) is a third option."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Kth Largest kept the k largest with a min-heap; here we keep the k smallest with a max-heap. What's the general rule?", Answer: "The top of a size-k heap must be the element you'd evict: the worst of the kept ones. Keeping the largest, the worst is the smallest, so use a min-heap. Keeping the smallest, the worst is the largest, so use a max-heap (negate in Python)."},
			{Prompt: "When is heapq.nsmallest(k, points, key=...) or a full sort the better choice?", Answer: "When k is close to n, a sort is O(n log n) either way and simpler. When k is small and n is huge, or points arrive as a stream you can't store, the size-k heap's O(n log k) time and O(k) memory win. heapq.nsmallest uses the same heap idea internally."},
		},
	},
	{
		ID:              "top-k-frequent-words",
		Number:          692,
		Order:           5,
		Title:           "Top K Frequent Words",
		Slug:            "top-k-frequent-words",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Hash Table", "String", "Sorting", "Heap (Priority Queue)", "Counting", "Google", "Google L3", "Google Phone Screen"},
		RelatedChapters: []int{11, 2},
		Description: `You are given a list of lowercase ` + "`words`" + ` (the same word may appear many times) and an integer ` + "`k`" + `. Return the ` + "`k`" + ` words that appear most often.

The order of the answer matters: sort it from the **most frequent** word to the least frequent. When two words appear the same number of times, the one that comes first in **alphabetical (dictionary) order** goes first, and the same rule decides which word makes the cut when several words tie for the last place.`,
		Constraints: []string{
			"1 <= words.length <= 500",
			"1 <= words[i].length <= 10",
			"words[i] consists of lowercase English letters",
			"1 <= k <= the number of distinct words",
		},
		Examples: []Example{
			{Input: `words = ["red","blue","red","green","blue","red"], k = 2`, Output: `["red","blue"]`, Explanation: "red appears 3 times, blue 2 times and green once. The two most frequent are red, then blue."},
			{Input: `words = ["go","run","go","fly","run","go","sky","fly","run"], k = 3`, Output: `["go","run","fly"]`, Explanation: "go and run both appear 3 times, so they are ordered alphabetically: go before run. fly (2 times) is third; sky (once) doesn't make the cut."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"words": []string{"red", "blue", "red", "green", "blue", "red"}, "k": 2}, Expected: []string{"red", "blue"}, Hidden: false, Note: "the example: distinct counts"},
			{Input: map[string]interface{}{"words": []string{"go", "run", "go", "fly", "run", "go", "sky", "fly", "run"}, "k": 3}, Expected: []string{"go", "run", "fly"}, Hidden: false, Note: "a tie in count is broken alphabetically"},
			{Input: map[string]interface{}{"words": []string{"b", "a", "c"}, "k": 3}, Expected: []string{"a", "b", "c"}, Hidden: true, Note: "every word once: purely alphabetical"},
			{Input: map[string]interface{}{"words": []string{"solo"}, "k": 1}, Expected: []string{"solo"}, Hidden: true, Note: "a single word"},
			{Input: map[string]interface{}{"words": []string{"cat", "bat", "cat", "bat", "ant", "ant", "dog"}, "k": 2}, Expected: []string{"ant", "bat"}, Hidden: true, Note: "three words tie at the cut: the alphabetically first two win, not cat"},
			{Input: map[string]interface{}{"words": []string{"app", "apple", "app", "apple", "ap"}, "k": 2}, Expected: []string{"app", "apple"}, Hidden: true, Note: "a prefix sorts before the longer word"},
			{Input: map[string]interface{}{"words": []string{"x", "y", "y", "z", "z", "z"}, "k": 1}, Expected: []string{"z"}, Hidden: true, Note: "the alphabetically last word can still be the most frequent"},
			{Input: map[string]interface{}{"words": []string{"m", "n", "n", "m", "o", "o", "o", "p"}, "k": 4}, Expected: []string{"o", "m", "n", "p"}, Hidden: true, Note: "k equals the number of distinct words: the whole ranking"},
		},
		TimeComplexity:  "O(n + k log u) for u distinct words",
		SpaceComplexity: "O(u)",
		StarterCode:     "def topKFrequent(words, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In K Closest Points every heap entry was a tuple with one key in front. Here the ranking has two rules: count first, then alphabetical order. How could a single tuple encode both rules so that the heap's 'smallest' entry is exactly the word that should come first?"},
			{Level: 2, Type: "insight", Content: "Count the words with a Counter, then put (-count, word) tuples in a heap. Python compares tuples left to right: the negated count makes higher counts smaller, and when counts tie the word itself breaks the tie alphabetically, which is just what we want. Popping k times reads the answer in the required order.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. Count the words: build a word -> count map.\n2. Build the heap: one (-count, word) entry per distinct word, then heapify.\n3. Read the top k times: each pop is the next word in the answer, already in the right order.\n4. Return the k words."},
			{Level: 4, Type: "skeleton", Content: "import heapq\nfrom collections import Counter\n\ndef topKFrequent(words, k):\n    counts = Counter(words)\n    heap = [____ for word, count in counts.items()]\n    heapq.heapify(heap)\n    result = []\n    for _ in range(k):\n        _, word = ____\n        result.append(word)\n    return result"},
		},
		Solution: Solution{
			Code: `import heapq
from collections import Counter


def topKFrequent(words, k):
    counts = Counter(words)                       # count the words
    heap = [(-count, word) for word, count in counts.items()]
    heapq.heapify(heap)                           # build the heap
    result = []
    for _ in range(k):
        _, word = heapq.heappop(heap)             # read the top: most frequent, then a-z
        result.append(word)
    return result`,
			Explanation: "Builds on K Closest Points (problem 4), where heap entries were tuples ordered by their first element. The new twist is a two-part key: (-count, word). The negated count puts the most frequent word on top, and the word itself breaks ties alphabetically because Python compares tuples element by element. Because the tie-break is part of the answer, we heapify every distinct word and pop k times, which returns the words already in the required order.",
			BruteForce:  "For each distinct word, count its occurrences by scanning the whole list, then sort the distinct words by (count descending, word ascending) and take the first k. Counting this way is O(u * n) for u distinct words: every scan re-reads words it has already counted. A hash map counts everything in one pass, and a heap avoids ordering all u words when only k are needed.",
			KeyInsight:  "A tuple key can encode a whole ranking: (-count, word) is smaller exactly when the word ranks higher (more frequent, or equally frequent and alphabetically earlier). So a plain min-heap of these tuples hands out the answer in order, one pop per word.",
			WhyItWorks:  "After heapify, the heap property holds: every entry is <= its children, so the top is the smallest tuple. Tuple order compares -count first, so a larger count always wins; only when counts are equal does it compare the words, and the alphabetically earlier word is smaller. So the smallest tuple is exactly the highest-ranked word. Popping removes it and restores the heap property, so the next pop is the highest-ranked of the rest. After k pops we have the top k in ranking order, and every word that tied at the cut was decided by the same alphabetical rule.",
			CommonMistakes: []string{
				"Breaking count ties the wrong way, e.g. heapq.nlargest(k, counts.items(), key=lambda p: (p[1], p[0])), which puts the alphabetically *later* word first among equal counts. Test: cat, bat and ant all appear twice and k = 2; the answer is [ant, bat]. Use (-count, word) with a min-heap or sort.",
				"Keeping a size-k min-heap of (count, word) like Kth Largest. Its eviction rule is wrong for ties: with equal counts it evicts the alphabetically *first* word, which is the one you want to keep. A size-k heap needs 'reverse alphabetical' as the tie-break, which a plain tuple can't express (see the Entry class alternative).",
				"Returning the heap list itself, or the k entries in heap order. A heap is only partially ordered; pop the entries one by one (or sort) to get the required order.",
			},
			TimeComplexity:  "O(n + k log u): O(n) to count, O(u) to heapify the u distinct words, and k pops",
			SpaceComplexity: "O(u) for the counts and the heap",
			Walkthrough: []WalkthroughStep{
				{Title: "Count the words", Explanation: "One pass over the words gives each distinct word's count.", CodeSnippet: "counts = Counter(words)", LineStart: 6, LineEnd: 6},
				{Title: "Build the heap", Explanation: "One (-count, word) tuple per distinct word. The negated count makes the most frequent word smallest; equal counts fall back to comparing words alphabetically.", CodeSnippet: "heap = [(-count, word) for word, count in counts.items()]\nheapq.heapify(heap)", LineStart: 7, LineEnd: 8},
				{Title: "Read the top", Explanation: "Each pop is the highest-ranked word still in the heap, so the answer is built in the required order.", CodeSnippet: "for _ in range(k):\n    _, word = heapq.heappop(heap)\n    result.append(word)", LineStart: 10, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{Name: "Sort every distinct word", Code: `from collections import Counter


def topKFrequent(words, k):
    counts = Counter(words)
    ranked = sorted(counts, key=lambda word: (-counts[word], word))
    return ranked[:k]`, TimeComplexity: "O(n + u log u)", SpaceComplexity: "O(u)", Explanation: "Sort all distinct words by the same two-part key and slice. Simplest to write and fine for small inputs, but it orders all u words when only k are needed."},
				{Name: "Size-k heap with a custom order", Code: `import heapq
from collections import Counter


class Entry:
    def __init__(self, count, word):
        self.count, self.word = count, word

    def __lt__(self, other):              # 'less' means 'ranks lower'
        if self.count != other.count:
            return self.count < other.count
        return self.word > other.word     # later in a-z ranks lower


def topKFrequent(words, k):
    heap = []
    for word, count in Counter(words).items():
        heapq.heappush(heap, Entry(count, word))
        if len(heap) > k:                 # trim to size k: drop the lowest-ranked
            heapq.heappop(heap)
    result = []
    while heap:
        result.append(heapq.heappop(heap).word)
    return result[::-1]                   # popped lowest first, so reverse`, TimeComplexity: "O(n + u log k)", SpaceComplexity: "O(u)", Explanation: "The Kth Largest pattern: keep only the k best, with the worst on top to evict. The worst word has the fewest uses or, on a tie, comes later alphabetically. Strings can't be negated, so a small class defines that order with __lt__. Popping yields the kept words from worst to best, so the list is reversed at the end."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can't you simply use (count, word) in a size-k min-heap, as in Kth Largest?", Answer: "The top of a size-k heap must be the entry to evict: the lowest-ranked one. With (count, word), equal counts compare words alphabetically, so the alphabetically *first* word sits on top and gets evicted, but that is the one that should be kept. You need the opposite order for the word part, and strings can't be negated, so you either heapify everything with (-count, word) or define a custom comparison."},
			{Prompt: "What changes if ties should be broken by first appearance in the list instead of alphabetically?", Answer: "Replace the word in the key with the index where it first appears: (-count, first_index, word). The smaller index ranks higher, and the rest of the algorithm is unchanged. Counting the first index is one extra dictionary filled during the counting pass."},
		},
	},
	{
		ID:              "task-scheduler",
		Number:          621,
		Order:           6,
		Title:           "Task Scheduler",
		Slug:            "task-scheduler",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Hash Table", "Greedy", "Heap (Priority Queue)", "Counting", "Google", "Google L4", "Google Phone Screen"},
		RelatedChapters: []int{11, 2, 17},
		Description: `A CPU must run a list of ` + "`tasks`" + `, each named by a capital letter A-Z; the same letter may appear many times. Time is split into equal intervals, and in each interval the CPU either runs one task or sits idle.

After running a task, the CPU must wait at least ` + "`n`" + ` intervals before running **the same letter** again (other letters may run in between). Tasks can run in any order. Return the minimum number of intervals needed to finish every task, idle intervals included.`,
		Constraints: []string{
			"1 <= tasks.length <= 10^4",
			"tasks[i] is an uppercase English letter",
			"0 <= n <= 100",
		},
		Examples: []Example{
			{Input: `tasks = ["A","A","A","B","B","B"], n = 2`, Output: "8", Explanation: "A B idle A B idle A B. Each A is followed by at least 2 other intervals before the next A, and the same for B."},
			{Input: `tasks = ["A","A","A","B","B","B"], n = 0`, Output: "6", Explanation: "No cooling, so the 6 tasks run back to back."},
			{Input: `tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2`, Output: "16", Explanation: "A B C A D E A F G A idle idle A idle idle A."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"tasks": []string{"A", "A", "A", "B", "B", "B"}, "n": 2}, Expected: 8, Hidden: false, Note: "the example: A B idle A B idle A B"},
			{Input: map[string]interface{}{"tasks": []string{"A", "A", "A", "B", "B", "B"}, "n": 0}, Expected: 6, Hidden: false, Note: "no cooling: just the task count"},
			{Input: map[string]interface{}{"tasks": []string{"A", "A", "A", "A", "A", "A", "B", "C", "D", "E", "F", "G"}, "n": 2}, Expected: 16, Hidden: false, Note: "one dominant task forces idle slots"},
			{Input: map[string]interface{}{"tasks": []string{"A", "C", "A", "B", "D", "B"}, "n": 1}, Expected: 6, Hidden: true, Note: "enough variety to fill every gap"},
			{Input: map[string]interface{}{"tasks": []string{"A", "A", "A", "B", "B", "B"}, "n": 3}, Expected: 10, Hidden: true, Note: "long cooldown: two idles per frame"},
			{Input: map[string]interface{}{"tasks": []string{"A"}, "n": 5}, Expected: 1, Hidden: true, Note: "a single task needs no idle time"},
			{Input: map[string]interface{}{"tasks": []string{"A", "A"}, "n": 2}, Expected: 4, Hidden: true, Note: "same task twice: A idle idle A"},
			{Input: map[string]interface{}{"tasks": []string{"A", "A", "A", "A", "B", "B", "C"}, "n": 1}, Expected: 7, Hidden: true, Note: "other tasks fill exactly the gaps"},
			{Input: map[string]interface{}{"tasks": []string{"A", "B", "C", "D", "E", "A", "B", "C", "D", "E"}, "n": 4}, Expected: 10, Hidden: true, Note: "five tied tasks fill each frame completely"},
		},
		TimeComplexity:  "O(T) where T is the answer",
		SpaceComplexity: "O(1)",
		StarterCode:     "def leastInterval(tasks, n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Imagine you're the CPU at some interval and several letters are allowed to run. Which one would you run to reduce idle time later, and why?"},
			{Level: 2, Type: "insight", Content: "Greedy simulation: at each interval run the ready letter with the most copies left, because the most frequent letter is the one most likely to force idle gaps at the end. A max-heap of remaining counts gives you that letter. After running it, it can't run for n intervals, so park it in a queue tagged with the interval when its cooldown ends; at that interval, put it back into the heap.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. Count the tasks: build a max-heap of the counts (negated).\n2. Each interval, time += 1. Run the most frequent ready task: if the heap isn't empty, pop it and use up one copy.\n3. Park it while it cools: if copies remain, append (time + n, count) to a FIFO queue.\n4. Release a cooled task: if the front of the queue's cooldown ends at this interval, push it back onto the heap.\n5. Stop when both the heap and the queue are empty; time is the answer."},
			{Level: 4, Type: "skeleton", Content: "import heapq\nfrom collections import Counter, deque\n\ndef leastInterval(tasks, n):\n    ready = [-count for count in Counter(tasks).values()]\n    heapq.heapify(ready)\n    cooling = deque()\n    time = 0\n    while ready or cooling:\n        time += 1\n        if ready:\n            remaining = ____\n            if remaining != 0:\n                cooling.append((____, remaining))\n        if cooling and ____:\n            heapq.heappush(ready, cooling.popleft()[1])\n    return time"},
		},
		Solution: Solution{
			Code: `import heapq
from collections import Counter, deque


def leastInterval(tasks, n):
    ready = [-count for count in Counter(tasks).values()]  # max-heap of remaining counts
    heapq.heapify(ready)
    cooling = deque()                  # (last tick it must wait, -remaining)
    time = 0
    while ready or cooling:
        time += 1
        if ready:                      # run the most frequent ready task
            remaining = heapq.heappop(ready) + 1
            if remaining != 0:         # park it while it cools
                cooling.append((time + n, remaining))
        if cooling and cooling[0][0] == time:
            heapq.heappush(ready, cooling.popleft()[1])  # release a cooled task
    return time`,
			Explanation: "Builds on Last Stone Weight (problem 1): a max-heap of counts, popped and pushed back after a change. The new idea is that an item can't go straight back: it spends n intervals in a cooling queue first. Each interval we run the ready letter with the most copies left (greedy), park it, and release whatever finished cooling. The counting formula in the alternatives computes the same number directly.",
			BruteForce:  "Try every order of the tasks and, for each, count the intervals including forced idles; take the minimum. With 12 tasks that's up to 12! = 479 million orders. The waste: most orders are clearly bad (running all the rare letters first leaves the frequent one to idle alone at the end), and a simple rule picks a best order directly.",
			KeyInsight:  "The letter with the most remaining copies is the bottleneck: if it isn't run when it's allowed, its copies bunch up at the end and every gap between them becomes idle time. So running the most frequent ready letter is always safe, and a max-heap finds it in O(log 26) time.",
			WhyItWorks:  "The simulation keeps two invariants: every letter with copies left is either in the ready heap (allowed to run now) or in the cooling queue with the last interval it must still wait, and the queue is ordered by that interval because tasks are parked in time order. A letter run at interval t is parked until t + n and pushed back at the end of interval t + n, so it can next run at t + n + 1, exactly n intervals later. The greedy choice is optimal by an exchange argument: if an optimal schedule runs a less frequent letter (or idles) where a more frequent one was ready, swapping their future copies never creates a cooling violation for the frequent letter and never lengthens the schedule. The counting formula alternative proves the same total: max((maxCount - 1) * (n + 1) + numTied, len(tasks)).",
			CommonMistakes: []string{
				"Re-adding a task to the heap at time + n but checking it before running the current interval, so it runs after only n - 1 waits. Run first, then release tasks whose wait ends at this interval, so they are ready from the next one.",
				"Assuming the answer is always (maxCount - 1) * (n + 1) + numTied. When there are many different letters, no idle is needed and the answer is just len(tasks); take the max of the two.",
				"Popping from the heap when it's empty during an idle interval. The heap can be empty while tasks are cooling: that interval is idle, so only advance time.",
			},
			TimeComplexity:  "O(T) where T is the returned number of intervals (at most len(tasks) * (n + 1)); each interval does O(1) work because the heap never holds more than 26 letters",
			SpaceComplexity: "O(1): at most 26 letters in the heap and the queue",
			Walkthrough: []WalkthroughStep{
				{Title: "Count the tasks", Explanation: "Only how many copies each letter has matters, not which letter it is. Negate the counts for a max-heap.", CodeSnippet: "ready = [-count for count in Counter(tasks).values()]\nheapq.heapify(ready)", LineStart: 6, LineEnd: 7},
				{Title: "Run the most frequent ready task", Explanation: "Each interval, pop the letter with the most copies left and use one copy (adding 1 to a negative count). If the heap is empty, the interval is idle.", CodeSnippet: "if ready:\n    remaining = heapq.heappop(ready) + 1", LineStart: 12, LineEnd: 13},
				{Title: "Park it while it cools", Explanation: "If copies remain, it may run again only after interval time + n; queue it with that interval.", CodeSnippet: "if remaining != 0:\n    cooling.append((time + n, remaining))", LineStart: 14, LineEnd: 15},
				{Title: "Release a cooled task", Explanation: "At the end of interval time + n the parked letter becomes ready for the next interval. The queue is in time order, so only the front can be due.", CodeSnippet: "if cooling and cooling[0][0] == time:\n    heapq.heappush(ready, cooling.popleft()[1])", LineStart: 16, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{Name: "Counting formula", Code: `from collections import Counter


def leastInterval(tasks, n):
    counts = Counter(tasks).values()
    most = max(counts)
    tied = sum(1 for c in counts if c == most)   # tasks sharing the top count
    frame = (most - 1) * (n + 1) + tied          # rows of n + 1 slots, plus the last row
    return max(frame, len(tasks))`, TimeComplexity: "O(n)", SpaceComplexity: "O(1)", Explanation: "Picture the most frequent letter (count m) laid out as m - 1 frames of n + 1 slots plus a last row. The other letters fill frame slots; idles appear only in unfilled slots. The last row holds every letter tied for count m. That gives (m - 1) * (n + 1) + tied intervals, unless there are so many tasks that the frames overflow, in which case no idle is ever needed and the answer is len(tasks). This is O(len(tasks)) to count, with no simulation."},
				{Name: "Simulation without a heap", Code: `from collections import Counter


def leastInterval(tasks, n):
    remaining = Counter(tasks)
    next_free = {task: 1 for task in remaining}  # earliest tick each task may run
    time = 0
    while remaining:
        time += 1
        best = None
        for task, count in remaining.items():    # scan every task: no heap
            if next_free[task] <= time and (best is None or count > remaining[best]):
                best = task
        if best is not None:
            remaining[best] -= 1
            next_free[best] = time + n + 1
            if remaining[best] == 0:
                del remaining[best]
    return time`, TimeComplexity: "O(T)", SpaceComplexity: "O(1)", Explanation: "The same greedy, but each interval scans all (at most 26) letters for the ready one with most copies left, instead of using a heap and a queue. Since the alphabet is fixed, the scan is also O(1) per interval; the heap version is the one that scales to a large alphabet."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "For tasks A A A B B B and n = 2, what happens to the heap and the queue in each of the first four intervals?", Answer: "Start: ready = {A:3, B:3}. t=1: run A (2 left), park (3, A). t=2: run B (2 left), park (4, B). t=3: heap empty, idle; front of queue is due at 3, so A goes back. t=4: run A (1 left), park (6, A); B is due at 4 and goes back. So far: A B idle A, matching the optimal A B idle A B idle A B."},
			{Prompt: "Why does the formula take max(..., len(tasks))?", Answer: "The frame formula counts slots assuming idles are needed. If there are more other tasks than idle slots, they simply extend the frames, no idle is ever forced, and the CPU is busy every interval. The answer can never be less than the number of tasks, so the max picks whichever bound applies."},
		},
	},
	{
		ID:              "design-twitter",
		Number:          355,
		Order:           7,
		Title:           "Design Twitter",
		Slug:            "design-twitter",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Hash Table", "Linked List", "Design", "Heap (Priority Queue)"},
		RelatedChapters: []int{11, 2, 7},
		Description: `Design a tiny social network. Users post tweets, follow and unfollow each other, and read a news feed.

Implement the ` + "`Twitter`" + ` class:
- ` + "`Twitter()`" + ` creates an empty network.
- ` + "`postTweet(userId, tweetId)`" + ` records that ` + "`userId`" + ` posted a new tweet with the given id. Every call is later in time than the one before; tweet ids are unique but not in time order.
- ` + "`getNewsFeed(userId)`" + ` returns the ids of the **10 most recent** tweets posted by the user or by anyone they currently follow, newest first (fewer than 10 if there aren't that many, an empty list if there are none).
- ` + "`follow(followerId, followeeId)`" + ` makes ` + "`followerId`" + ` follow ` + "`followeeId`" + `.
- ` + "`unfollow(followerId, followeeId)`" + ` stops that; unfollowing someone you don't follow does nothing.

A user always sees their own tweets.`,
		Constraints: []string{
			"1 <= userId, followerId, followeeId <= 500",
			"0 <= tweetId <= 10^4",
			"All tweet ids are unique",
			"At most 3 * 10^4 calls in total to postTweet, getNewsFeed, follow and unfollow",
		},
		Examples: []Example{
			{Input: "[\"Twitter\",\"postTweet\",\"getNewsFeed\",\"follow\",\"postTweet\",\"getNewsFeed\",\"unfollow\",\"getNewsFeed\"]\n[[],[1,5],[1],[1,2],[2,6],[1],[1,2],[1]]", Output: "[null,null,[5],null,null,[6,5],null,[5]]", Explanation: "User 1 posts 5 and sees it. After following user 2, who posts 6, user 1's feed is [6,5], newest first. After unfollowing 2, only [5] remains."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"}, "values": [][]int{{}, {1, 5}, {1}, {1, 2}, {2, 6}, {1}, {1, 2}, {1}}}, Expected: []interface{}{nil, nil, []int{5}, nil, nil, []int{6, 5}, nil, []int{5}}, Hidden: false, Note: "the example: follow adds 6 to the feed, unfollow removes it"},
			{Input: map[string]interface{}{"operations": []string{"Twitter", "getNewsFeed", "follow", "getNewsFeed"}, "values": [][]int{{}, {3}, {3, 4}, {3}}}, Expected: []interface{}{nil, []interface{}{}, nil, []interface{}{}}, Hidden: false, Note: "a feed with no tweets is an empty list"},
			{Input: map[string]interface{}{"operations": []string{"Twitter", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "getNewsFeed"}, "values": [][]int{{}, {1, 100}, {1, 101}, {1, 102}, {1, 103}, {1, 104}, {1, 105}, {1, 106}, {1, 107}, {1, 108}, {1, 109}, {1, 110}, {1, 111}, {1}}}, Expected: []interface{}{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, []int{111, 110, 109, 108, 107, 106, 105, 104, 103, 102}}, Hidden: true, Note: "more than ten tweets: only the ten newest, newest first"},
			{Input: map[string]interface{}{"operations": []string{"Twitter", "follow", "follow", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "getNewsFeed", "getNewsFeed"}, "values": [][]int{{}, {1, 2}, {1, 3}, {2, 21}, {3, 31}, {1, 11}, {2, 22}, {4, 41}, {3, 32}, {1}, {2}}}, Expected: []interface{}{nil, nil, nil, nil, nil, nil, nil, nil, nil, []int{32, 22, 11, 31, 21}, []int{22, 21}}, Hidden: true, Note: "interleaved tweets from three users merge by time, not by user"},
			{Input: map[string]interface{}{"operations": []string{"Twitter", "postTweet", "postTweet", "unfollow", "follow", "follow", "getNewsFeed"}, "values": [][]int{{}, {1, 7}, {2, 8}, {1, 2}, {1, 1}, {1, 2}, {1}}}, Expected: []interface{}{nil, nil, nil, nil, nil, nil, []int{8, 7}}, Hidden: true, Note: "unfollowing someone not followed and following yourself are harmless"},
			{Input: map[string]interface{}{"operations": []string{"Twitter", "postTweet", "postTweet", "follow", "postTweet", "getNewsFeed", "unfollow", "postTweet", "getNewsFeed"}, "values": [][]int{{}, {1, 0}, {2, 9}, {2, 1}, {1, 4}, {2}, {2, 1}, {1, 6}, {2}}}, Expected: []interface{}{nil, nil, nil, nil, nil, []int{4, 9, 0}, nil, nil, []int{9}}, Hidden: true, Note: "tweet id 0 and tweet ids not in time order"},
			{Input: map[string]interface{}{"operations": []string{"Twitter", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "follow", "getNewsFeed"}, "values": [][]int{{}, {1, 1}, {2, 2}, {1, 3}, {2, 4}, {1, 5}, {2, 6}, {2, 7}, {1, 2}, {1}}}, Expected: []interface{}{nil, nil, nil, nil, nil, nil, nil, nil, nil, []int{7, 6, 5, 4, 3, 2, 1}}, Hidden: true, Note: "the cap of ten is not reached: return every visible tweet"},
		},
		TimeComplexity:  "O(f) per getNewsFeed (f = users followed), O(1) for the other operations",
		SpaceComplexity: "O(t + e) for t tweets and e follow relations",
		StarterCode:     "class Twitter:\n    def __init__(self):\n        pass\n\n    def postTweet(self, userId: int, tweetId: int) -> None:\n        pass\n\n    def getNewsFeed(self, userId: int) -> List[int]:\n        pass\n\n    def follow(self, followerId: int, followeeId: int) -> None:\n        pass\n\n    def unfollow(self, followerId: int, followeeId: int) -> None:\n        pass",
		Harness:         Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Each user's own tweets are already in time order. If you had one sorted list per followed user, how would you pick the newest tweet overall, then the next newest, without merging everything?"},
			{Level: 2, Type: "insight", Content: "This is 'merge k sorted lists' (chapter 7) with a heap. Give every tweet a global timestamp. For a feed, put each followed user's newest tweet into a max-heap keyed by time. Pop the newest overall, then push that same user's next older tweet. Ten pops give the feed; you never touch older tweets.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. Store: a clock, a dict user -> list of (time, tweetId), and a dict user -> set of followees.\n2. Seed the heap: for the user and everyone they follow, push (-time, tweetId, user, index) for their newest tweet.\n3. Pop the newest overall: pop and append the tweet id to the feed.\n4. Refill from the same user: if that user has an older tweet (index - 1), push it.\n5. Stop after 10 tweets or when the heap is empty."},
			{Level: 4, Type: "skeleton", Content: "def getNewsFeed(self, userId):\n    heap = []\n    for uid in self.following[userId] | {userId}:\n        if self.tweets[uid]:\n            idx = len(self.tweets[uid]) - 1\n            time, tweet_id = self.tweets[uid][idx]\n            heap.append((____, tweet_id, uid, idx))\n    heapq.heapify(heap)\n    feed = []\n    while heap and ____:\n        _, tweet_id, uid, idx = heapq.heappop(heap)\n        feed.append(tweet_id)\n        if ____:\n            time, older_id = self.tweets[uid][idx - 1]\n            heapq.heappush(heap, (____, older_id, uid, idx - 1))\n    return feed"},
		},
		Solution: Solution{
			Code: `import heapq
from collections import defaultdict


class Twitter:
    def __init__(self):
        self.clock = 0
        self.tweets = defaultdict(list)    # userId -> [(time, tweetId)], oldest first
        self.following = defaultdict(set)  # followerId -> set of followeeIds

    def postTweet(self, userId: int, tweetId: int) -> None:
        self.tweets[userId].append((self.clock, tweetId))
        self.clock += 1

    def getNewsFeed(self, userId: int) -> List[int]:
        heap = []
        for uid in self.following[userId] | {userId}:   # seed with each newest tweet
            if self.tweets[uid]:
                idx = len(self.tweets[uid]) - 1
                time, tweet_id = self.tweets[uid][idx]
                heap.append((-time, tweet_id, uid, idx))
        heapq.heapify(heap)
        feed = []
        while heap and len(feed) < 10:                  # pop the newest overall
            _, tweet_id, uid, idx = heapq.heappop(heap)
            feed.append(tweet_id)
            if idx > 0:                                 # refill from the same user
                time, older_id = self.tweets[uid][idx - 1]
                heapq.heappush(heap, (-time, older_id, uid, idx - 1))
        return feed

    def follow(self, followerId: int, followeeId: int) -> None:
        self.following[followerId].add(followeeId)

    def unfollow(self, followerId: int, followeeId: int) -> None:
        self.following[followerId].discard(followeeId)`,
			Explanation: "Builds on K Closest Points (problem 4), where heap entries were tuples ordered by a key. Here the key is a negated timestamp (max-heap: newest first), and each entry also remembers which user's list it came from and where, so after popping it we can push the next tweet from the same list. That is the k-way merge from chapter 7's Merge K Sorted Lists, stopped after 10 items.",
			BruteForce:  "Keep one global list of every tweet; for a feed, scan it from newest to oldest and keep tweets whose author is visible, stopping at 10. That's O(T) per feed in the worst case (T = all tweets ever), e.g. when the user follows nobody active. Another brute force gathers all tweets of all followees and sorts them. The waste: it reads tweets that can never be in the top 10.",
			KeyInsight:  "Each user's tweets are already sorted by time. The newest tweet overall must be the newest of some user, so a heap holding one 'frontier' tweet per user always has the next answer on top; after popping it, only that user's frontier moves.",
			WhyItWorks:  "Invariant: during the merge, for every visible user with unread tweets, the heap holds exactly one entry: their newest tweet not yet put in the feed. The overall newest unread tweet is the newest of some user's unread tweets, so it is the heap's top (largest time, i.e. smallest -time). After popping it, pushing that user's next older tweet restores the invariant. So the feed comes out in exact newest-first order. Time stamps come from a global counter, so they are unique and ties never need a tie-breaker. Cost: building the heap is O(f), and 10 pops and pushes are O(log f) each.",
			CommonMistakes: []string{
				"Sorting by tweetId instead of by time. Ids are unique but not in posting order (tweet 4 can be posted after tweet 9). Keep a global counter.",
				"Forgetting the user's own tweets. The feed includes them even if they don't follow themselves; add userId to the set of authors you merge.",
				"Using remove() for unfollow on a set that may not contain the followee. That raises KeyError; discard() does nothing if it's absent.",
				"Adding the user to their own following set to include their tweets. It works until they unfollow themselves; merging following | {userId} keeps the data clean.",
			},
			TimeComplexity:  "O(f) per getNewsFeed: heapify f entries, then 10 pops/pushes of O(log f); postTweet, follow and unfollow are O(1)",
			SpaceComplexity: "O(t + e) for t tweets and e follow relations; O(f) extra per feed",
			Walkthrough: []WalkthroughStep{
				{Title: "Store tweets with a global clock", Explanation: "Each user's list is in posting order, and the clock gives every tweet a unique time.", CodeSnippet: "self.tweets[userId].append((self.clock, tweetId))\nself.clock += 1", LineStart: 12, LineEnd: 13},
				{Title: "Seed the heap", Explanation: "One entry per visible user: their newest tweet, with the index to continue from.", CodeSnippet: "for uid in self.following[userId] | {userId}:\n    if self.tweets[uid]:", LineStart: 17, LineEnd: 22},
				{Title: "Pop the newest overall", Explanation: "The top has the largest time among all frontiers, so it is the next tweet in the feed.", CodeSnippet: "while heap and len(feed) < 10:\n    _, tweet_id, uid, idx = heapq.heappop(heap)\n    feed.append(tweet_id)", LineStart: 24, LineEnd: 26},
				{Title: "Refill from the same user", Explanation: "Only the popped user's frontier moved; push their next older tweet if there is one.", CodeSnippet: "if idx > 0:\n    time, older_id = self.tweets[uid][idx - 1]\n    heapq.heappush(heap, (-time, older_id, uid, idx - 1))", LineStart: 27, LineEnd: 29},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: scan the global timeline", Code: `class Twitter:
    def __init__(self):
        self.all_tweets = []        # (userId, tweetId), oldest first
        self.following = {}

    def postTweet(self, userId, tweetId):
        self.all_tweets.append((userId, tweetId))

    def getNewsFeed(self, userId):
        visible = self.following.get(userId, set()) | {userId}
        feed = []
        for author, tweet_id in reversed(self.all_tweets):   # scan every tweet ever posted
            if author in visible:
                feed.append(tweet_id)
                if len(feed) == 10:
                    break
        return feed

    def follow(self, followerId, followeeId):
        self.following.setdefault(followerId, set()).add(followeeId)

    def unfollow(self, followerId, followeeId):
        self.following.get(followerId, set()).discard(followeeId)`, TimeComplexity: "O(T) per getNewsFeed for T tweets in total", SpaceComplexity: "O(T + e)", Explanation: "Keeps every tweet in one list and scans backwards for visible authors. Simple, and fine when everyone posts a lot, but a user following quiet accounts forces a scan of the entire history."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does each heap entry carry the user and an index, not just the tweet?", Answer: "After popping a tweet we must push the next one from the same user's list. The entry has to say which list it came from (uid) and where in it (idx) so we can find tweets[uid][idx - 1]. Without that, the heap couldn't advance the right frontier."},
			{Prompt: "How is this the same algorithm as Merge K Sorted Lists?", Answer: "Each followed user's tweet list is a sorted list (by time). The feed is the merge of those lists, newest first, truncated to 10. Merge K Sorted Lists keeps one node per list in a min-heap, pops the smallest and pushes its next node; here we keep one tweet per user in a max-heap, pop the newest and push that user's next older tweet."},
		},
	},
	{
		ID:              "find-k-pairs-with-smallest-sums",
		Number:          373,
		Order:           8,
		Title:           "Find K Pairs with Smallest Sums",
		Slug:            "find-k-pairs-with-smallest-sums",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Heap (Priority Queue)", "Google", "Google L4", "Google Phone Screen"},
		RelatedChapters: []int{11, 7},
		Description: `You are given two integer arrays ` + "`nums1`" + ` and ` + "`nums2`" + `, each sorted in non-decreasing order, and an integer ` + "`k`" + `.

A **pair** ` + "`[u, v]`" + ` takes one value ` + "`u`" + ` from ` + "`nums1`" + ` and one value ` + "`v`" + ` from ` + "`nums2`" + `. Pairs are formed from positions, so if a value appears twice in an array, it can form two different pairs with the same partner.

Return the ` + "`k`" + ` pairs with the **smallest sums** ` + "`u + v`" + `. Return the pairs in **any order**. When several pairs tie for the last place, any of them may be chosen, as long as the sums you return are the ` + "`k`" + ` smallest.`,
		Constraints: []string{
			"1 <= nums1.length, nums2.length <= 10^5",
			"-10^9 <= nums1[i], nums2[i] <= 10^9",
			"nums1 and nums2 are sorted in non-decreasing order",
			"1 <= k <= 10^4 and k <= nums1.length * nums2.length",
		},
		Examples: []Example{
			{Input: "nums1 = [1,3,9], nums2 = [2,5,8], k = 3", Output: "[[1,2],[3,2],[1,5]]", Explanation: "The sums of all nine pairs are 3, 6, 9 (with u = 1), 5, 8, 11 (u = 3) and 11, 14, 17 (u = 9). The three smallest are 3, 5 and 6."},
			{Input: "nums1 = [1,1,2], nums2 = [1,2,3], k = 2", Output: "[[1,1],[1,1]]", Explanation: "Both 1s in nums1 pair with nums2[0] = 1 for a sum of 2. They come from different positions, so the pair appears twice."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums1": []int{1, 3, 9}, "nums2": []int{2, 5, 8}, "k": 3}, Expected: [][]int{{1, 2}, {3, 2}, {1, 5}}, Hidden: false, Note: "the example: the answer moves along both arrays"},
			{Input: map[string]interface{}{"nums1": []int{1, 1, 2}, "nums2": []int{1, 2, 3}, "k": 2}, Expected: [][]int{{1, 1}, {1, 1}}, Hidden: false, Note: "duplicate values form separate pairs"},
			{Input: map[string]interface{}{"nums1": []int{1, 2}, "nums2": []int{3}, "k": 2}, Expected: [][]int{{1, 3}, {2, 3}}, Hidden: true, Note: "k is every possible pair"},
			{Input: map[string]interface{}{"nums1": []int{5}, "nums2": []int{5}, "k": 1}, Expected: [][]int{{5, 5}}, Hidden: true, Note: "one value in each array"},
			{Input: map[string]interface{}{"nums1": []int{-10, -4, 0}, "nums2": []int{-3, 5, 6}, "k": 4}, Expected: [][]int{{-10, -3}, {-4, -3}, {-10, 5}, {-10, 6}}, Hidden: true, Note: "negative values"},
			{Input: map[string]interface{}{"nums1": []int{1, 2, 3}, "nums2": []int{1, 100, 200}, "k": 3}, Expected: [][]int{{1, 1}, {2, 1}, {3, 1}}, Hidden: true, Note: "walking only along nums2 with u = nums1[0] fails here"},
			{Input: map[string]interface{}{"nums1": []int{1, 100}, "nums2": []int{1, 2, 3}, "k": 3}, Expected: [][]int{{1, 1}, {1, 2}, {1, 3}}, Hidden: true, Note: "walking only along nums1 with v = nums2[0] fails here"},
			{Input: map[string]interface{}{"nums1": []int{1, 2, 3}, "nums2": []int{1, 2, 3}, "k": 6}, Expected: [][]int{{1, 1}, {1, 2}, {2, 1}, {1, 3}, {2, 2}, {3, 1}}, Hidden: true, Note: "pushing both neighbours without a visited set returns [2,2] twice"},
		},
		TimeComplexity:  "O(k log k)",
		SpaceComplexity: "O(k)",
		StarterCode:     "def kSmallestPairs(nums1, nums2, k):\n    # Write your solution here\n    pass",
		Harness: Harness{
			Validator: `def check(raw, args, result, expected):
    from collections import Counter
    nums1, nums2, k = raw["nums1"], raw["nums2"], raw["k"]
    if not isinstance(result, list) or len(result) != k:
        return False, "return exactly k pairs"
    have1, have2 = Counter(nums1), Counter(nums2)
    used = Counter()
    for pair in result:
        if not isinstance(pair, list) or len(pair) != 2:
            return False, "each pair must be a list [u, v]"
        u, v = pair
        if have1[u] == 0 or have2[v] == 0:
            return False, str(pair) + " does not take u from nums1 and v from nums2"
        used[(u, v)] += 1
        if used[(u, v)] > have1[u] * have2[v]:
            return False, str(pair) + " appears more times than it can be formed"
    got = sorted(u + v for u, v in result)
    want = sorted(u + v for u, v in expected)
    if got != want:
        return False, "the sums " + str(got) + " are not the k smallest " + str(want)
    return True, ""`,
		},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Picture all pairs as a grid: row i uses nums1[i], column j uses nums2[j]. Because both arrays are sorted, each row's sums only grow to the right. Which single pair in each row is the only one that could be the next smallest overall?"},
			{Level: 2, Type: "insight", Content: "Each row is a sorted list, so this is the k-way merge from Design Twitter: keep one 'frontier' entry per row in a min-heap, keyed by sum. Pop the smallest; it is the next answer. Then push the next pair from the same row, (i, j + 1). Seed the heap with column 0 of the first k rows only: a row beyond k can't contribute, since the k rows above it already have smaller first sums.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. Seed the heap: for the first k values of nums1, push (nums1[i] + nums2[0], i, 0).\n2. Pop the smallest sum: append [nums1[i], nums2[j]] to the answer.\n3. Refill from the same row: if j + 1 is in range, push (nums1[i] + nums2[j + 1], i, j + 1).\n4. Stop after k pairs."},
			{Level: 4, Type: "skeleton", Content: "import heapq\n\ndef kSmallestPairs(nums1, nums2, k):\n    heap = [(____, i, 0) for i, a in enumerate(nums1[:k])]\n    heapq.heapify(heap)\n    result = []\n    while len(result) < k:\n        _, i, j = heapq.heappop(heap)\n        result.append([nums1[i], nums2[j]])\n        if ____:\n            heapq.heappush(heap, (____, i, j + 1))\n    return result"},
		},
		Solution: Solution{
			Code: `import heapq


def kSmallestPairs(nums1, nums2, k):
    # seed the heap: each of the first k rows starts at column 0
    heap = [(a + nums2[0], i, 0) for i, a in enumerate(nums1[:k])]
    heapq.heapify(heap)
    result = []
    while len(result) < k:
        _, i, j = heapq.heappop(heap)          # pop the smallest sum
        result.append([nums1[i], nums2[j]])
        if j + 1 < len(nums2):                 # refill from the same row
            heapq.heappush(heap, (nums1[i] + nums2[j + 1], i, j + 1))
    return result`,
			Explanation: "Builds on Design Twitter (problem 7), which merged several sorted tweet lists by keeping each list's frontier in a heap and refilling from the list that was popped. Here the sorted lists are implicit: row i is nums1[i] paired with every nums2[j], in increasing sum order. The new twist is that there could be 10^5 rows, but only the first k can ever matter, so the heap is seeded with at most k frontiers.",
			BruteForce:  "Build all m * n pairs, sort them by sum and return the first k: O(mn log(mn)) time and O(mn) memory. With 10^5 values per array that is 10^10 pairs. The waste: we build and sort every pair, though only k of them, and only pairs near the top-left of the grid, can ever be in the answer.",
			KeyInsight:  "In row i the sums increase with j, so the smallest unused pair of that row is always its leftmost unused one. The next smallest pair overall is therefore one of the row frontiers, and a min-heap of frontiers finds it in O(log k). Each pop exposes just one new candidate: the next pair in the same row.",
			WhyItWorks:  "Invariant: for every seeded row i, the heap holds exactly one entry, the leftmost pair of that row not yet returned (unless the row is used up), and every returned pair has a sum <= every pair still unreturned. Any unreturned pair (i, j) has a sum >= its row's frontier, which is >= the heap's top, so popping the top returns a smallest remaining pair. Pushing (i, j + 1) restores the one-per-row invariant. Rows i >= k never matter: rows 0..k-1 each have a first pair whose sum is <= nums1[i] + nums2[0], so they alone already hold k pairs no larger than any pair in a later row. Each pair is pushed at most once (only from its left neighbour), so no pair can be returned twice.",
			CommonMistakes: []string{
				"Pushing both neighbours (i + 1, j) and (i, j + 1) after each pop without a visited set. Pair (1, 1) is reached from (0, 1) and from (1, 0), so it is returned twice. Either keep a visited set or, as here, seed the whole first column and only ever move right.",
				"Walking a single row or column greedily, e.g. pairing nums1[0] with every value of nums2. On nums1 = [1,2,3], nums2 = [1,100,200] the three smallest sums all use nums2[0]. Every row's frontier has to compete in the heap.",
				"Seeding every row of nums1. It is correct, but with 10^5 rows and k = 1 it does far more work than needed; only the first min(k, len(nums1)) rows can contribute.",
			},
			TimeComplexity:  "O(k log k): at most k entries are seeded, and k pops and pushes each cost O(log k)",
			SpaceComplexity: "O(k) for the heap",
			Walkthrough: []WalkthroughStep{
				{Title: "Seed the heap", Explanation: "Each of the first k rows starts at its smallest pair, column 0. Rows after the k-th can never be needed.", CodeSnippet: "heap = [(a + nums2[0], i, 0) for i, a in enumerate(nums1[:k])]\nheapq.heapify(heap)", LineStart: 6, LineEnd: 7},
				{Title: "Pop the smallest sum", Explanation: "The top is the smallest frontier, and every unused pair is at least its row's frontier, so the top is the next answer.", CodeSnippet: "_, i, j = heapq.heappop(heap)\nresult.append([nums1[i], nums2[j]])", LineStart: 10, LineEnd: 11},
				{Title: "Refill from the same row", Explanation: "Only row i's frontier moved; its next candidate is one column to the right, if there is one.", CodeSnippet: "if j + 1 < len(nums2):\n    heapq.heappush(heap, (nums1[i] + nums2[j + 1], i, j + 1))", LineStart: 12, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{Name: "Sort every pair", Code: `def kSmallestPairs(nums1, nums2, k):
    pairs = [[u, v] for u in nums1 for v in nums2]   # all m * n pairs
    pairs.sort(key=lambda pair: pair[0] + pair[1])
    return pairs[:k]`, TimeComplexity: "O(mn log(mn))", SpaceComplexity: "O(mn)", Explanation: "The brute force: build the whole grid of pairs and sort by sum. Easy to trust, and useful to check the heap version on small inputs, but it can't handle arrays of 10^5 values."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to seed only the first k rows?", Answer: "The first pair of row i, nums1[i] + nums2[0], is the smallest sum in that row, and nums1 is sorted, so it is >= the first pair of every earlier row. If i >= k, the k rows before it each already have a pair with a sum <= every sum in row i, so k answers can be found without row i (ties can be broken in favour of the earlier rows)."},
			{Prompt: "How does this relate to finding the k-th smallest element of a sorted matrix?", Answer: "It is the same problem: the grid of sums has rows and columns sorted. Seed the first column, pop k times and refill from the same row; the k-th pop is the k-th smallest. The heap holds at most one entry per row, so it stays O(min(k, rows))."},
		},
	},
	{
		ID:              "find-median-from-data-stream",
		Number:          295,
		Order:           9,
		Title:           "Find Median from Data Stream",
		Slug:            "find-median-from-data-stream",
		Difficulty:      "Hard",
		Category:        "heap-priority-queue",
		Tags:            []string{"Two Pointers", "Design", "Sorting", "Heap (Priority Queue)", "Data Stream", "Google", "Google L5", "Google Onsite"},
		RelatedChapters: []int{11, 9},
		Description: `The **median** of a list of numbers is the middle value once the list is sorted. If the list has an even length, it is the average of the two middle values: the median of [2,3,4] is 3, and of [2,3] it is 2.5.

Implement the ` + "`MedianFinder`" + ` class:
- ` + "`MedianFinder()`" + ` starts with no numbers.
- ` + "`addNum(num)`" + ` adds an integer from the stream.
- ` + "`findMedian()`" + ` returns the median of every number added so far, as a float. It is only called after at least one number has been added. Answers within 10^-5 of the true value are accepted.`,
		Constraints: []string{
			"-10^5 <= num <= 10^5",
			"findMedian is only called when at least one number has been added",
			"At most 5 * 10^4 calls in total to addNum and findMedian",
		},
		Examples: []Example{
			{Input: "[\"MedianFinder\",\"addNum\",\"addNum\",\"findMedian\",\"addNum\",\"findMedian\"]\n[[],[1],[2],[],[3],[]]", Output: "[null,null,null,1.5,null,2.0]", Explanation: "After adding 1 and 2 the median is (1 + 2) / 2 = 1.5. After adding 3, the sorted list is [1,2,3] and the median is 2."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"}, "values": [][]int{{}, {1}, {2}, {}, {3}, {}}}, Expected: []interface{}{nil, nil, nil, 1.5, nil, 2.0}, Hidden: false, Note: "the example: 1.5 then 2.0"},
			{Input: map[string]interface{}{"operations": []string{"MedianFinder", "addNum", "findMedian"}, "values": [][]int{{}, {7}, {}}}, Expected: []interface{}{nil, nil, 7.0}, Hidden: false, Note: "a single number is its own median"},
			{Input: map[string]interface{}{"operations": []string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"}, "values": [][]int{{}, {-1}, {}, {-2}, {}, {-3}, {}, {-4}, {}}}, Expected: []interface{}{nil, nil, -1.0, nil, -1.5, nil, -2.0, nil, -2.5}, Hidden: true, Note: "negative, decreasing numbers"},
			{Input: map[string]interface{}{"operations": []string{"MedianFinder", "addNum", "addNum", "addNum", "addNum", "findMedian", "addNum", "findMedian"}, "values": [][]int{{}, {5}, {15}, {1}, {3}, {}, {8}, {}}}, Expected: []interface{}{nil, nil, nil, nil, nil, 4.0, nil, 5.0}, Hidden: true, Note: "numbers arrive out of order: a new small number must move to the low half"},
			{Input: map[string]interface{}{"operations": []string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian", "addNum", "findMedian"}, "values": [][]int{{}, {4}, {4}, {4}, {}, {4}, {}}}, Expected: []interface{}{nil, nil, nil, nil, 4.0, nil, 4.0}, Hidden: true, Note: "all equal values"},
			{Input: map[string]interface{}{"operations": []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "addNum", "findMedian", "addNum", "findMedian"}, "values": [][]int{{}, {100000}, {-100000}, {}, {0}, {1}, {}, {2}, {}}}, Expected: []interface{}{nil, nil, nil, 0.0, nil, nil, 0.5, nil, 1.0}, Hidden: true, Note: "extreme values and a fractional even median"},
			{Input: map[string]interface{}{"operations": []string{"MedianFinder", "addNum", "addNum", "addNum", "addNum", "addNum", "addNum", "findMedian"}, "values": [][]int{{}, {1}, {2}, {3}, {4}, {5}, {6}, {}}}, Expected: []interface{}{nil, nil, nil, nil, nil, nil, nil, 3.5}, Hidden: true, Note: "increasing input: catches never moving numbers to the high half"},
		},
		TimeComplexity:  "O(log n) per addNum, O(1) per findMedian",
		SpaceComplexity: "O(n)",
		StarterCode:     "class MedianFinder:\n    def __init__(self):\n        pass\n\n    def addNum(self, num: int) -> None:\n        pass\n\n    def findMedian(self) -> float:\n        pass",
		Harness:         Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The median splits the sorted numbers into a lower half and an upper half. Which one or two numbers do you actually need to see, and what would you need to quickly get at in each half?"},
			{Level: 2, Type: "insight", Content: "Keep two heaps: a max-heap 'low' with the smaller half and a min-heap 'high' with the larger half. Keep every number in low <= every number in high, and the sizes equal or low one bigger. Then the median is low's top (odd count) or the average of both tops (even count). To add, push into low, move low's largest to high, and if high is now bigger, move high's smallest back.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. Set up: low = [] (max-heap, store negated), high = [] (min-heap).\n2. Route the number: push -num onto low, then pop low's top and push it (un-negated) onto high. This keeps every low <= every high.\n3. Rebalance: if high has more items than low, move high's top back to low.\n4. Read the median: if low is bigger, its top; otherwise the average of the two tops."},
			{Level: 4, Type: "skeleton", Content: "import heapq\n\nclass MedianFinder:\n    def __init__(self):\n        self.low = []\n        self.high = []\n\n    def addNum(self, num: int) -> None:\n        heapq.heappush(self.low, ____)\n        heapq.heappush(self.high, ____)\n        if ____:\n            heapq.heappush(self.low, -heapq.heappop(self.high))\n\n    def findMedian(self) -> float:\n        if len(self.low) > len(self.high):\n            return ____\n        return ____"},
		},
		Solution: Solution{
			Code: `import heapq


class MedianFinder:
    def __init__(self):
        self.low = []    # max-heap (negated): the smaller half
        self.high = []   # min-heap: the larger half

    def addNum(self, num: int) -> None:
        heapq.heappush(self.low, -num)                    # route through low...
        heapq.heappush(self.high, -heapq.heappop(self.low))  # ...so high gets low's max
        if len(self.high) > len(self.low):                # rebalance: low may hold one extra
            heapq.heappush(self.low, -heapq.heappop(self.high))

    def findMedian(self) -> float:
        if len(self.low) > len(self.high):                # odd count: low's top
            return float(-self.low[0])
        return (-self.low[0] + self.high[0]) / 2          # even count: average of the tops`,
			Explanation: "This rung combines the core tools of the ladder so far: a max-heap by negation (Last Stone Weight), a heap whose top is a boundary value (Kth Largest), and two heaps working together. The new idea is keeping two heaps balanced: the max-heap's top and the min-heap's top are the two middle values, so the median is always read from the tops in O(1).",
			BruteForce:  "Keep all numbers in a list and sort it on every findMedian: O(n log n) per query. Keeping the list sorted with bisect.insort is better but each insert still shifts up to n elements: O(n) per add, around 1.25 * 10^9 element moves for 5 * 10^4 adds in the worst case. The waste: the median only needs the middle, but we maintain the full order of both halves.",
			KeyInsight:  "We never need the full order, only the boundary between the halves. A max-heap exposes the largest of the lower half and a min-heap the smallest of the upper half, which are exactly the one or two middle values, and each insert costs O(log n).",
			WhyItWorks:  "Invariant after each addNum: (a) every value in low is <= every value in high, and (b) len(low) == len(high) or len(low) == len(high) + 1. Adding num: push it into low, then move low's maximum m to high. Everything left in low is <= m, since m was the largest; and everything left in low is <= the old contents of high, since it is either an old low value (true by (a)) or num when num < m, and then m was an old low value, so num < m <= old high. So (a) holds. If high is now bigger, moving its minimum back to low keeps (a) (that minimum is >= all of low and <= all of high) and restores (b). With (a) and (b), the sorted list is low's values followed by high's, so the middle value is low's maximum (odd count) and the two middle values are low's maximum and high's minimum (even count).",
			CommonMistakes: []string{
				"Pushing every number straight into one heap by comparing with the median and never rebalancing. The sizes drift apart and the tops stop being the middle values; increasing input [1..6] would put everything in one half.",
				"Forgetting to negate when moving between the heaps. low stores -value; push -num into low, and push -heappop(low) into high (and the reverse).",
				"Returning integer division (a + b) // 2. The median of [1,2] is 1.5, not 1. Use / 2.",
			},
			TimeComplexity:  "O(log n) per addNum (at most three heap operations), O(1) per findMedian",
			SpaceComplexity: "O(n) for the two heaps",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up two halves", Explanation: "low is a max-heap (negated values) of the smaller half; high is a min-heap of the larger half.", CodeSnippet: "self.low = []\nself.high = []", LineStart: 6, LineEnd: 7},
				{Title: "Route the number", Explanation: "Pass num through low so that the largest of low's values (num included) is what ends up in high. This keeps every low value <= every high value.", CodeSnippet: "heapq.heappush(self.low, -num)\nheapq.heappush(self.high, -heapq.heappop(self.low))", LineStart: 10, LineEnd: 11},
				{Title: "Rebalance", Explanation: "low may hold one extra value but never fewer than high; move high's smallest back if needed.", CodeSnippet: "if len(self.high) > len(self.low):\n    heapq.heappush(self.low, -heapq.heappop(self.high))", LineStart: 12, LineEnd: 13},
				{Title: "Read the median", Explanation: "Odd count: low's top is the middle value. Even count: average the two tops.", CodeSnippet: "if len(self.low) > len(self.high):\n    return float(-self.low[0])\nreturn (-self.low[0] + self.high[0]) / 2", LineStart: 16, LineEnd: 18},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: sorted list", Code: `import bisect


class MedianFinder:
    def __init__(self):
        self.values = []

    def addNum(self, num):
        bisect.insort(self.values, num)       # O(n): shifts later elements right

    def findMedian(self):
        n = len(self.values)
        mid = n // 2
        if n % 2 == 1:
            return float(self.values[mid])
        return (self.values[mid - 1] + self.values[mid]) / 2`, TimeComplexity: "O(n) per addNum, O(1) per findMedian", SpaceComplexity: "O(n)", Explanation: "Keeps all values sorted with bisect.insort and reads the middle. Correct and short; insertion shifts elements, so it's O(n) per add."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why route every number through low and then move low's top to high, instead of comparing num with the median first?", Answer: "Routing always works without special cases: after pushing num into low, low's maximum is the only value that could break 'low <= high', and moving it to high fixes that. Comparing with the current median also works, but needs care when a heap is empty and still needs the rebalance step."},
			{Prompt: "If 99% of the numbers were known to lie in 0..100, how could you make addNum O(1)?", Answer: "Keep a count array for 0..100 plus counts (or small heaps) for the rare outliers below and above. findMedian walks the 101 counters to find the middle position, which is O(1) for a fixed range. This is a classic follow-up; the two-heap answer is the general one."},
		},
	},
	{
		ID:              "meeting-rooms-iii",
		Number:          2402,
		Order:           10,
		Title:           "Meeting Rooms III",
		Slug:            "meeting-rooms-iii",
		Difficulty:      "Hard",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Hash Table", "Sorting", "Heap (Priority Queue)", "Simulation", "Google", "Google L5", "Google Phone Screen"},
		RelatedChapters: []int{11, 18},
		Description: `An office has ` + "`n`" + ` meeting rooms numbered ` + "`0`" + ` to ` + "`n - 1`" + `. You are given a list of ` + "`meetings`" + `, where ` + "`meetings[i] = [start, end]`" + ` is a meeting that wants the half-open time range ` + "`[start, end)`" + `: it occupies a room from ` + "`start`" + ` up to, but not including, ` + "`end`" + `, so a room whose meeting ends at time 5 is free for a meeting starting at 5. Every ` + "`start`" + ` is different. The list is **not** necessarily sorted.

Rooms are assigned to meetings in order of their **original start times**, using these rules:
1. A meeting takes the free room with the **lowest number**.
2. If no room is free, the meeting is **delayed** until a room frees up. It takes the room that becomes free **earliest** (the lowest-numbered one if several free up at the same time), and it keeps its original **duration** ` + "`end - start`" + `.
3. When rooms free up, waiting meetings are served in order of their original start times.

Return the number of the room that hosted the **most meetings**. If several rooms tie, return the lowest number.`,
		Constraints: []string{
			"1 <= n <= 100",
			"1 <= meetings.length <= 10^5",
			"meetings[i].length == 2",
			"0 <= start < end <= 5 * 10^5",
			"all start values are unique",
		},
		Examples: []Example{
			{Input: "n = 2, meetings = [[0,6],[1,3],[2,5],[4,9]]", Output: "0", Explanation: "[0,6) takes room 0 and [1,3) takes room 1. At time 2 both rooms are busy; room 1 frees first (at 3), so [2,5) is delayed to [3,6) in room 1. At time 4 both rooms are busy until 6; room 0 is the lower number, so [4,9) runs as [6,11) in room 0. Each room hosted 2 meetings; the tie goes to room 0."},
			{Input: "n = 3, meetings = [[1,20],[2,10],[3,5],[4,9],[6,8]]", Output: "1", Explanation: "Rooms 0, 1, 2 take the first three meetings (ending at 20, 10, 5). [4,9) waits for room 2 and runs as [5,10). [6,8) waits: rooms 1 and 2 both free at 10, so it takes room 1 as [10,12). Counts are [1,2,2], so room 1 wins the tie."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 2, "meetings": [][]int{{0, 6}, {1, 3}, {2, 5}, {4, 9}}}, Expected: 0, Hidden: false, Note: "the example: rooms tie on count, lowest number wins"},
			{Input: map[string]interface{}{"n": 3, "meetings": [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}}, Expected: 1, Hidden: false, Note: "delayed meetings; two rooms free at the same time"},
			{Input: map[string]interface{}{"n": 2, "meetings": [][]int{{9, 12}, {0, 5}, {7, 9}}}, Expected: 0, Hidden: true, Note: "unsorted input: processing in the given order picks room 1"},
			{Input: map[string]interface{}{"n": 2, "meetings": [][]int{{6, 10}, {7, 9}, {5, 11}, {9, 10}}}, Expected: 0, Hidden: true, Note: "a delayed meeting keeps its duration and does not end at its original end"},
			{Input: map[string]interface{}{"n": 3, "meetings": [][]int{{3, 6}, {4, 5}, {6, 8}}}, Expected: 0, Hidden: true, Note: "a room ending at 6 is free for a meeting starting at 6"},
			{Input: map[string]interface{}{"n": 2, "meetings": [][]int{{1, 5}, {3, 4}, {9, 14}}}, Expected: 0, Hidden: true, Note: "lowest free room number, not the room that freed up first"},
			{Input: map[string]interface{}{"n": 3, "meetings": [][]int{{0, 10}, {1, 10}, {2, 3}, {4, 5}, {6, 7}}}, Expected: 2, Hidden: true, Note: "the highest-numbered room can win"},
			{Input: map[string]interface{}{"n": 1, "meetings": [][]int{{5, 10}, {0, 3}, {2, 4}}}, Expected: 0, Hidden: true, Note: "a single room hosts everything"},
			{Input: map[string]interface{}{"n": 2, "meetings": [][]int{{3, 6}, {0, 4}, {1, 2}, {5, 8}, {7, 9}}}, Expected: 1, Hidden: true, Note: "rooms are reused as soon as they free up"},
		},
		TimeComplexity:  "O(m log m + m log n)",
		SpaceComplexity: "O(m + n)",
		StarterCode:     "def mostBooked(n, meetings):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When a meeting arrives, you need two different answers quickly: 'which free room has the lowest number?' and, if none is free, 'which busy room frees up first?'. Can one ordering answer both questions?"},
			{Level: 2, Type: "insight", Content: "Use two heaps. free is a min-heap of room numbers. busy is a min-heap of (end time, room). Process meetings sorted by start. First move every busy room whose end time is <= start into free. If free is non-empty, pop the lowest room. Otherwise pop the top of busy: it frees earliest (lowest number on a tie, thanks to the tuple), and the delayed meeting ends at that time plus its duration.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. Set up two heaps: free holds all room numbers; busy starts empty; count is a list of zeros.\n2. Sort the meetings by start.\n3. Release finished rooms: pop every (end, room) from busy with end <= start and push room into free.\n4. Take the lowest free room: if free is non-empty, pop it and push (end, room) into busy.\n5. Delay into the earliest room: otherwise pop (finish, room) from busy and push (finish + duration, room).\n6. Count the meeting for that room; at the end return the lowest room with the highest count."},
			{Level: 4, Type: "skeleton", Content: "import heapq\n\ndef mostBooked(n, meetings):\n    free = list(range(n))\n    busy = []\n    count = [0] * n\n    for start, end in sorted(meetings):\n        while busy and ____:\n            _, room = heapq.heappop(busy)\n            heapq.heappush(free, room)\n        if free:\n            room = heapq.heappop(free)\n            heapq.heappush(busy, ____)\n        else:\n            finish, room = heapq.heappop(busy)\n            heapq.heappush(busy, ____)\n        count[room] += 1\n    return ____"},
		},
		Solution: Solution{
			Code: `import heapq


def mostBooked(n, meetings):
    free = list(range(n))              # min-heap of free room numbers
    busy = []                          # min-heap of (end time, room)
    count = [0] * n
    for start, end in sorted(meetings):
        while busy and busy[0][0] <= start:    # release rooms that are done
            _, room = heapq.heappop(busy)
            heapq.heappush(free, room)
        if free:                               # take the lowest free room
            room = heapq.heappop(free)
            heapq.heappush(busy, (end, room))
        else:                                  # delay into the earliest room
            finish, room = heapq.heappop(busy)
            heapq.heappush(busy, (finish + end - start, room))
        count[room] += 1
    return count.index(max(count))`,
			Explanation: "Builds on Task Scheduler (problem 6), where a task waited in a cooling queue until a known time and then moved back to the ready heap. Here rooms do the same: they sit in a busy heap keyed by end time and move to a free heap once a meeting starts at or after that time. Like Find Median (problem 9), two heaps work together, but now each answers its own question: free gives the lowest room number, busy gives the room that frees up first.",
			BruteForce:  "Keep each room's next free time in an array. For every meeting, scan all n rooms for the lowest-numbered free one; if none is free, scan again for the earliest free time. That is O(m log m + m * n), which is fine for n <= 100 but wasteful: each scan re-reads rooms whose state hasn't changed since the last meeting.",
			KeyInsight:  "The two questions need two orders: free rooms by number, busy rooms by end time (then number). Keeping each set in its own heap makes both questions O(log n), and a room moves between the heaps only when its end time is reached.",
			WhyItWorks:  "Invariant: just before a meeting is placed, free holds exactly the rooms whose current meeting ends at or before its start, and busy holds the rest with their true end times. Meetings are handled in start order, and delayed meetings only ever start later, so a room released for one start is also free for every later start. If free is non-empty, its top is the lowest free number (rule 1). If it is empty, every room is busy past start; busy's top has the smallest end time, ties broken by room number because the tuple compares the room next (rule 2). A delayed meeting starts at that end time, so it ends at finish + (end - start). Because meetings are processed in start order, a waiting meeting always gets the next room before any later meeting (rule 3).",
			CommonMistakes: []string{
				"Using one heap of (free time, room) and always popping its top. When several rooms are already free, that picks the room that freed up earliest, not the lowest-numbered one. Release every finished room into a separate heap ordered by number first.",
				"Giving a delayed meeting its original end time. It starts when the room frees up, so it ends at finish + (end - start).",
				"Releasing rooms with end < start instead of end <= start. Meetings are half-open, so a room whose meeting ends at 6 is free for one starting at 6.",
				"Forgetting to sort the meetings. Rooms must be handed out in order of start time; the input can be in any order.",
			},
			TimeComplexity:  "O(m log m + m log n): sorting m meetings, then each room move or assignment is a heap operation on at most n rooms",
			SpaceComplexity: "O(m + n): the sorted copy of the meetings and the two heaps",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up two heaps", Explanation: "list(range(n)) is already sorted, so it is a valid min-heap of free rooms. busy will hold (end time, room).", CodeSnippet: "free = list(range(n))\nbusy = []\ncount = [0] * n", LineStart: 5, LineEnd: 7},
				{Title: "Release finished rooms", Explanation: "Every room whose meeting ends at or before this start goes back to free. The busy heap gives them in end-time order, so we stop at the first room still in use.", CodeSnippet: "while busy and busy[0][0] <= start:\n    _, room = heapq.heappop(busy)\n    heapq.heappush(free, room)", LineStart: 9, LineEnd: 11},
				{Title: "Take the lowest free room", Explanation: "The free heap's top is the lowest room number; the meeting holds it until end.", CodeSnippet: "room = heapq.heappop(free)\nheapq.heappush(busy, (end, room))", LineStart: 13, LineEnd: 14},
				{Title: "Delay into the earliest room", Explanation: "No room is free. The busy top frees first (lowest number on a tie); the meeting starts then and keeps its duration.", CodeSnippet: "finish, room = heapq.heappop(busy)\nheapq.heappush(busy, (finish + end - start, room))", LineStart: 16, LineEnd: 17},
				{Title: "Read the answer", Explanation: "Count each meeting for its room. list.index returns the first maximum, which is the lowest room number on a tie.", CodeSnippet: "count[room] += 1\nreturn count.index(max(count))", LineStart: 18, LineEnd: 19},
			},
			Alternatives: []AltSolution{
				{Name: "Scan every room", Code: `def mostBooked(n, meetings):
    free_at = [0] * n                  # when each room is next free
    count = [0] * n
    for start, end in sorted(meetings):
        room = next((r for r in range(n) if free_at[r] <= start), None)
        if room is None:               # all busy: earliest, lowest number on a tie
            room = min(range(n), key=lambda r: free_at[r])
            free_at[room] += end - start
        else:
            free_at[room] = end
        count[room] += 1
    return count.index(max(count))`, TimeComplexity: "O(m log m + m * n)", SpaceComplexity: "O(m + n)", Explanation: "The brute force: one array of free times, scanned for every meeting. With n <= 100 it is fast enough, and it is a good way to check the heap version, but the heaps make each meeting O(log n) instead of O(n)."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why must every finished room be released before choosing, rather than just checking the top of busy?", Answer: "Several rooms may have finished by this start, and the rule is to take the lowest-numbered of them, not the one that finished first. Only after all finished rooms are in the free heap does its top answer 'lowest free number'. Rooms that finished long ago are already there from earlier releases."},
			{Prompt: "What changes if a delayed meeting should take the lowest-numbered room among all rooms that become free within the next 5 time units?", Answer: "The busy heap's top alone is no longer enough. You would pop every busy room with end <= top_end + 5 into a temporary list, choose the lowest number among them, and push the others back. The two-heap structure stays; only the delay rule changes."},
		},
	},
	{
		ID:              "the-skyline-problem",
		Number:          218,
		Order:           11,
		Title:           "The Skyline Problem",
		Slug:            "the-skyline-problem",
		Difficulty:      "Hard",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Divide and Conquer", "Binary Indexed Tree", "Segment Tree", "Line Sweep", "Heap (Priority Queue)", "Ordered Set", "Google", "Google L5", "Google Onsite"},
		RelatedChapters: []int{11, 8, 18},
		Description: `A city is drawn as rectangles standing on flat ground (height 0). Each building is ` + "`buildings[i] = [left, right, height]`" + `: it covers the x-range from ` + "`left`" + ` up to ` + "`right`" + ` (the building's left wall is at ` + "`left`" + `, its right wall at ` + "`right`" + `) and is ` + "`height`" + ` tall. The list is sorted by ` + "`left`" + `.

The **skyline** is the outline you see from far away: at every x, its height is the height of the tallest building covering x, where a building covers x when ` + "`left <= x < right`" + `, or 0 if no building does.

Return the skyline as a list of **key points** ` + "`[x, y]`" + `, sorted by x. A key point marks an x where the skyline's height changes to ` + "`y`" + `, and the height stays ` + "`y`" + ` until the next key point. The last key point is where the rightmost building ends, with ` + "`y = 0`" + `. Two consecutive key points must never have the same height: if the height doesn't change, don't add a point.`,
		Constraints: []string{
			"1 <= buildings.length <= 10^4",
			"0 <= left < right <= 2^31 - 1",
			"1 <= height <= 2^31 - 1",
			"buildings is sorted by left in non-decreasing order",
		},
		Examples: []Example{
			{Input: "buildings = [[1,6,5],[2,4,8],[3,9,6],[10,14,4],[12,16,3]]", Output: "[[1,5],[2,8],[4,6],[9,0],[10,4],[14,3],[16,0]]", Explanation: "At 1 the skyline rises to 5. At 2 the 8-tall building takes over. At 4 it ends; the tallest still standing is the 6-tall one (the 5-tall one runs until 6, but it's hidden). At 9 nothing is standing, so the height drops to 0. The second group gives 4 at 10, 3 at 14 when the 4-tall building ends, and 0 at 16."},
			{Input: "buildings = [[0,3,4],[3,6,4]]", Output: "[[0,4],[6,0]]", Explanation: "The two buildings touch at x = 3 with the same height, so the height never changes there and no key point is added."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"buildings": [][]int{{1, 6, 5}, {2, 4, 8}, {3, 9, 6}, {10, 14, 4}, {12, 16, 3}}}, Expected: [][]int{{1, 5}, {2, 8}, {4, 6}, {9, 0}, {10, 4}, {14, 3}, {16, 0}}, Hidden: false, Note: "the example: overlaps, hidden buildings and a gap"},
			{Input: map[string]interface{}{"buildings": [][]int{{0, 3, 4}, {3, 6, 4}}}, Expected: [][]int{{0, 4}, {6, 0}}, Hidden: false, Note: "touching buildings of equal height: no point at 3"},
			{Input: map[string]interface{}{"buildings": [][]int{{1, 5, 3}}}, Expected: [][]int{{1, 3}, {5, 0}}, Hidden: true, Note: "a single building"},
			{Input: map[string]interface{}{"buildings": [][]int{{1, 10, 8}, {3, 5, 4}}}, Expected: [][]int{{1, 8}, {10, 0}}, Hidden: true, Note: "a shorter building fully hidden inside a taller one"},
			{Input: map[string]interface{}{"buildings": [][]int{{1, 3, 2}, {5, 7, 4}}}, Expected: [][]int{{1, 2}, {3, 0}, {5, 4}, {7, 0}}, Hidden: true, Note: "a gap drops the skyline to the ground"},
			{Input: map[string]interface{}{"buildings": [][]int{{1, 6, 4}, {2, 6, 7}}}, Expected: [][]int{{1, 4}, {2, 7}, {6, 0}}, Hidden: true, Note: "two buildings end at the same x: one point, not two"},
			{Input: map[string]interface{}{"buildings": [][]int{{1, 4, 2}, {1, 3, 5}}}, Expected: [][]int{{1, 5}, {3, 2}, {4, 0}}, Hidden: true, Note: "two buildings start at the same x: only the taller one shows"},
			{Input: map[string]interface{}{"buildings": [][]int{{1, 3, 5}, {3, 6, 2}}}, Expected: [][]int{{1, 5}, {3, 2}, {6, 0}}, Hidden: true, Note: "right walls are exclusive: at x = 3 only the 2-tall building stands"},
			{Input: map[string]interface{}{"buildings": [][]int{{0, 2147483647, 2147483647}}}, Expected: [][]int{{0, 2147483647}, {2147483647, 0}}, Hidden: true, Note: "maximum coordinates and height"},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def getSkyline(buildings):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The skyline can only change at an x where some building starts or ends. If you swept from left to right over just those x values, what one number would you need to read at each of them, and what set of buildings would you need to keep to read it?"},
			{Level: 2, Type: "insight", Content: "Sweep over the sorted wall positions and keep a max-heap of the buildings standing there, as (-height, right). At each x, push every building whose left wall is at or before x. Removing an ended building from the middle of a heap is expensive, so delete lazily: while the top's right <= x, pop it. Ended buildings below the top don't matter until they surface. The top's height is the skyline height at x; record [x, height] only if it differs from the last point.", ChapterRef: 11},
			{Level: 3, Type: "plan", Content: "1. Collect the critical x's: every left and right wall, sorted and de-duplicated.\n2. Start the heap with a ground entry (height 0, never ends).\n3. Push buildings that start: for each x, push (-height, right) for every building with left <= x not pushed yet.\n4. Drop ended buildings: while the top's right <= x, pop it.\n5. Read the top and record a change: if -top height differs from the last recorded height, append [x, height].\n6. Return the key points."},
			{Level: 4, Type: "skeleton", Content: "import heapq\n\ndef getSkyline(buildings):\n    xs = sorted({x for left, right, _ in buildings for x in (left, right)})\n    live = [(0, float(\"inf\"))]\n    result = []\n    i = 0\n    for x in xs:\n        while i < len(buildings) and buildings[i][0] <= x:\n            left, right, height = buildings[i]\n            heapq.heappush(live, ____)\n            i += 1\n        while ____:\n            heapq.heappop(live)\n        height = ____\n        if not result or ____:\n            result.append([x, height])\n    return result"},
		},
		Solution: Solution{
			Code: `import heapq


def getSkyline(buildings):
    xs = sorted({x for left, right, _ in buildings for x in (left, right)})
    live = [(0, float("inf"))]          # max-heap of (-height, right); 0 = ground
    result = []
    i = 0
    for x in xs:
        while i < len(buildings) and buildings[i][0] <= x:
            left, right, height = buildings[i]     # push buildings that start here
            heapq.heappush(live, (-height, right))
            i += 1
        while live[0][1] <= x:                     # drop ended buildings lazily
            heapq.heappop(live)
        height = -live[0][0]                       # read the top
        if not result or result[-1][1] != height:  # record only real changes
            result.append([x, height])
    return result`,
			Explanation: "Builds on Meeting Rooms III (problem 10), where rooms left a heap once their end time was reached. Here buildings leave when the sweep passes their right wall, but the heap is ordered by height, not by end, so an ended building can be buried below the top where we can't cheaply reach it. The new idea is lazy deletion (mentioned in Kth Largest's self-check): leave ended buildings in the heap and discard them only when they reach the top, because only the top is ever read.",
			BruteForce:  "For every critical x (each left or right wall), scan all n buildings to find the tallest one with left <= x < right, and record the height when it changes. That is O(n^2): there are up to 2n critical x's, and each scan re-checks every building, although only the buildings that start or end near x changed since the previous x.",
			KeyInsight:  "Between two consecutive critical x's the set of standing buildings doesn't change, so the skyline is the maximum of a set that gains a building at each left wall and loses one at each right wall. A max-heap gives that maximum, and lazy deletion lets us remove ended buildings only when they would otherwise be read.",
			WhyItWorks:  "Invariant: after processing x, every building with left <= x has been pushed, and the heap's top is a building with right > x (or the ground entry, which never ends). Buildings below the top may have ended, but each is at most as tall as the top, which is still standing, so they can't change the maximum until they surface, and at that point the while loop pops them. So -top height is the tallest building with left <= x < right, which is the skyline height at x. The height is constant between consecutive critical x's, so checking only those x's finds every change, and comparing with the last recorded height drops points that don't change it. Each building is pushed and popped at most once, giving O(n log n).",
			CommonMistakes: []string{
				"Treating the right wall as part of the building (left <= x <= right). At x = right the building must already be gone: on [[1,3,5],[3,6,2]] the height at 3 is 2, not 5. Pop while right <= x.",
				"Appending a point at every critical x, even when the height didn't change. Two touching buildings of equal height ([[0,3,4],[3,6,4]]) would produce a spurious [3,4]. Compare with the last recorded height first.",
				"Handling each wall as a separate event with its own output point. When two buildings start or end at the same x, that yields two points with the same x. Process all buildings at one x before reading the top, as the sorted set of x's does here.",
				"Trying to remove an ended building from the middle of the heap with list.remove and re-heapify. That is O(n) per removal and O(n^2) overall; lazy deletion keeps it O(log n).",
			},
			TimeComplexity:  "O(n log n): sorting up to 2n critical x's, and every building is pushed and popped at most once",
			SpaceComplexity: "O(n) for the critical x's, the heap and the result",
			Walkthrough: []WalkthroughStep{
				{Title: "Collect the critical x's", Explanation: "The skyline can only change where a wall stands. A set removes duplicates, so buildings sharing a wall produce one x.", CodeSnippet: "xs = sorted({x for left, right, _ in buildings for x in (left, right)})", LineStart: 5, LineEnd: 5},
				{Title: "Push buildings that start", Explanation: "buildings is sorted by left, so a single pointer i adds each building once, as soon as the sweep reaches its left wall. Negate the height for a max-heap and keep right to know when it ends.", CodeSnippet: "while i < len(buildings) and buildings[i][0] <= x:\n    left, right, height = buildings[i]\n    heapq.heappush(live, (-height, right))\n    i += 1", LineStart: 10, LineEnd: 13},
				{Title: "Drop ended buildings", Explanation: "Lazy deletion: only the top is ever read, so only an ended building on top has to go. The ground entry has right = infinity and is never popped.", CodeSnippet: "while live[0][1] <= x:\n    heapq.heappop(live)", LineStart: 14, LineEnd: 15},
				{Title: "Read the top and record a change", Explanation: "The top is the tallest standing building. Add a key point only when the height differs from the previous one.", CodeSnippet: "height = -live[0][0]\nif not result or result[-1][1] != height:\n    result.append([x, height])", LineStart: 16, LineEnd: 18},
			},
			Alternatives: []AltSolution{
				{Name: "Scan every building at each x", Code: `def getSkyline(buildings):
    result = []
    for x in sorted({x for left, right, _ in buildings for x in (left, right)}):
        height = max((h for left, right, h in buildings if left <= x < right), default=0)
        if not result or result[-1][1] != height:
            result.append([x, height])
    return result`, TimeComplexity: "O(n^2)", SpaceComplexity: "O(n)", Explanation: "The brute force: at each critical x, take the tallest building covering it directly. It uses the same critical x's and the same 'record only changes' rule, so it is a good check for the heap version, but each x re-scans all n buildings."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to leave ended buildings in the heap?", Answer: "We only ever read the top. An ended building below the top is no taller than the top, which is still standing, so it can't affect the maximum. If it later reaches the top, the while loop sees right <= x and pops it before the top is read. Each building is still popped at most once, so the total cost is unchanged."},
			{Prompt: "What changes if the buildings are not sorted by left?", Answer: "Sort them by left first (O(n log n), which doesn't change the overall bound); the pointer i depends on that order. Alternatively, build an event list of (x, -height, right) for left walls and handle right walls by lazy deletion, sorting the events by x."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, HeapProblems...)
}
