package practice

// HeapProblems contains all heap-priority-queue category problems, in ladder
// order. The subgoal labels (build the heap, push the new item, trim to size
// k, read the top) are the ones chapter 11 teaches.
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
		ID:              "k-closest-points-to-origin",
		Number:          973,
		Order:           3,
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
		ID:              "task-scheduler",
		Number:          621,
		Order:           4,
		Title:           "Task Scheduler",
		Slug:            "task-scheduler",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Hash Table", "Greedy", "Heap (Priority Queue)", "Counting"},
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
		Order:           5,
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
			Explanation: "Builds on K Closest Points (problem 3), where heap entries were tuples ordered by a key. Here the key is a negated timestamp (max-heap: newest first), and each entry also remembers which user's list it came from and where, so after popping it we can push the next tweet from the same list. That is the k-way merge from chapter 7's Merge K Sorted Lists, stopped after 10 items.",
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
		ID:              "find-median-from-data-stream",
		Number:          295,
		Order:           6,
		Title:           "Find Median from Data Stream",
		Slug:            "find-median-from-data-stream",
		Difficulty:      "Hard",
		Category:        "heap-priority-queue",
		Tags:            []string{"Two Pointers", "Design", "Sorting", "Heap (Priority Queue)", "Data Stream"},
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
			Explanation: "The ladder's last rung combines everything so far: a max-heap by negation (Last Stone Weight), a heap whose top is a boundary value (Kth Largest), and two heaps working together. The new idea is keeping two heaps balanced: the max-heap's top and the min-heap's top are the two middle values, so the median is always read from the tops in O(1).",
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
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, HeapProblems...)
}
