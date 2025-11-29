package practice

// HeapProblems contains all heap-priority-queue category problems
var HeapProblems = []*Problem{
	{
		ID:              "kth-largest-element-in-stream",
		Number:          71,
		Title:           "Kth Largest Element in a Stream",
		Difficulty:      "Easy",
		Category:        "heap-priority-queue",
		Tags:            []string{"Tree", "Design", "Binary Search Tree", "Heap", "Binary Tree", "Data Stream"},
		RelatedChapters: []int{7, 11, 12},
		Description: `Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:
- KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
- int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.`,
		Constraints: []string{
			"1 <= k <= 10^4",
			"0 <= nums.length <= 10^4",
			"-10^4 <= nums[i] <= 10^4",
			"-10^4 <= val <= 10^4",
			"At most 10^4 calls will be made to add",
			"It is guaranteed that there will be at least k elements when add is called",
		},
		Examples: []Example{
			{Input: `["KthLargest","add","add","add","add","add"]
[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]`, Output: "[null,4,5,5,8,8]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"KthLargest", "add", "add", "add", "add", "add"}, "values": []interface{}{[]int{3, 4, 5, 8, 2}, 3, 5, 10, 9, 4}}, Expected: []interface{}{nil, 4, 5, 5, 8, 8}},
		},
		TimeComplexity:  "O(log k) per add",
		SpaceComplexity: "O(k)",
		StarterCode: `class KthLargest:
    def __init__(self, k: int, nums: List[int]):
        pass

    def add(self, val: int) -> int:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a min-heap of size k. The smallest element in heap is the kth largest overall.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Keep only k largest elements in min-heap. On add: push value, if size > k, pop smallest. Top is kth largest."},
			{Level: 3, Type: "code", Content: "heapq.heappush(heap, val). if len(heap) > k: heapq.heappop(heap). return heap[0]."},
		},
		Solution: Solution{
			Code: `import heapq

class KthLargest:
    def __init__(self, k: int, nums):
        self.k = k
        self.heap = []
        for num in nums:
            self.add(num)

    def add(self, val: int) -> int:
        heapq.heappush(self.heap, val)
        if len(self.heap) > self.k:
            heapq.heappop(self.heap)
        return self.heap[0]`,
			Explanation:     "Min-heap of size k keeps k largest elements. Smallest of these (heap top) is the kth largest.",
			TimeComplexity:  "O(log k) per add",
			SpaceComplexity: "O(k)",
			Walkthrough: []WalkthroughStep{
				{Title: "Initialize with adds", Explanation: "Use add method to build initial heap", CodeSnippet: "for num in nums:\n    self.add(num)", LineStart: 7, LineEnd: 8},
				{Title: "Maintain size k", Explanation: "Pop smallest when over k elements", CodeSnippet: "if len(self.heap) > self.k:\n    heapq.heappop(self.heap)", LineStart: 12, LineEnd: 13},
				{Title: "Return kth largest", Explanation: "Min of k largest = kth largest", CodeSnippet: "return self.heap[0]", LineStart: 14, LineEnd: 14},
			},
		},
	},
	{
		ID:              "last-stone-weight",
		Number:          72,
		Title:           "Last Stone Weight",
		Difficulty:      "Easy",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Heap"},
		RelatedChapters: []int{7, 11},
		Description: `You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:
- If x == y, both stones are destroyed, and
- If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.

At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left, return 0.`,
		Constraints: []string{
			"1 <= stones.length <= 30",
			"1 <= stones[i] <= 1000",
		},
		Examples: []Example{
			{Input: "stones = [2,7,4,1,8,1]", Output: "1"},
			{Input: "stones = [1]", Output: "1"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"stones": []int{2, 7, 4, 1, 8, 1}}, Expected: 1},
			{Input: map[string]interface{}{"stones": []int{1}}, Expected: 1},
		},
		TimeComplexity:  "O(n log n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def lastStoneWeight(stones):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use max-heap to efficiently get two heaviest stones. Python has min-heap, so negate values.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Pop two largest, if different, push difference back. Continue until 0 or 1 stone remains."},
			{Level: 3, Type: "code", Content: "heap = [-s for s in stones]. heapify. while len > 1: pop two, push -diff if nonzero. return -heap[0] or 0."},
		},
		Solution: Solution{
			Code: `import heapq

def lastStoneWeight(stones):
    # Use negative values for max-heap behavior
    heap = [-s for s in stones]
    heapq.heapify(heap)

    while len(heap) > 1:
        y = -heapq.heappop(heap)  # Heaviest
        x = -heapq.heappop(heap)  # Second heaviest
        if x != y:
            heapq.heappush(heap, -(y - x))

    return -heap[0] if heap else 0`,
			Explanation:     "Max-heap (via negation) for heaviest stones. Pop two, push difference if nonzero. Repeat until done.",
			TimeComplexity:  "O(n log n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Max-heap via negation", Explanation: "Python only has min-heap, negate for max", CodeSnippet: "heap = [-s for s in stones]", LineStart: 5, LineEnd: 5},
				{Title: "Smash stones", Explanation: "Pop two heaviest, push difference if nonzero", CodeSnippet: "if x != y:\n    heapq.heappush(heap, -(y - x))", LineStart: 11, LineEnd: 12},
				{Title: "Return result", Explanation: "Remaining stone weight or 0", CodeSnippet: "return -heap[0] if heap else 0", LineStart: 14, LineEnd: 14},
			},
		},
	},
	{
		ID:              "k-closest-points-to-origin",
		Number:          73,
		Title:           "K Closest Points to Origin",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Math", "Divide and Conquer", "Geometry", "Sorting", "Heap", "Quickselect"},
		RelatedChapters: []int{7, 11, 12},
		Description: `Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., sqrt((x1 - x2)^2 + (y1 - y2)^2)).

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).`,
		Constraints: []string{
			"1 <= k <= points.length <= 10^4",
			"-10^4 <= xi, yi <= 10^4",
		},
		Examples: []Example{
			{Input: "points = [[1,3],[-2,2]], k = 1", Output: "[[-2,2]]"},
			{Input: "points = [[3,3],[5,-1],[-2,4]], k = 2", Output: "[[3,3],[-2,4]]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"points": [][]int{{1, 3}, {-2, 2}}, "k": 1}, Expected: [][]int{{-2, 2}}},
			{Input: map[string]interface{}{"points": [][]int{{3, 3}, {5, -1}, {-2, 4}}, "k": 2}, Expected: [][]int{{3, 3}, {-2, 4}}},
		},
		TimeComplexity:  "O(n log k)",
		SpaceComplexity: "O(k)",
		StarterCode:     "def kClosest(points, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use max-heap of size k. Compare by squared distance (no need for sqrt). Keep k smallest distances.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "For each point, push (-dist, point) to heap. If size > k, pop largest. Result is k closest points."},
			{Level: 3, Type: "code", Content: "dist = x*x + y*y. heappush with -dist for max-heap. if len > k: heappop. Return points from heap."},
		},
		Solution: Solution{
			Code: `import heapq

def kClosest(points, k):
    heap = []

    for x, y in points:
        dist = x * x + y * y
        heapq.heappush(heap, (-dist, [x, y]))
        if len(heap) > k:
            heapq.heappop(heap)

    return [point for _, point in heap]`,
			Explanation:     "Max-heap of size k (negated distance). Keeps k points with smallest distances. No sqrt needed for comparison.",
			TimeComplexity:  "O(n log k)",
			SpaceComplexity: "O(k)",
			Walkthrough: []WalkthroughStep{
				{Title: "Squared distance", Explanation: "Skip sqrt, relative order unchanged", CodeSnippet: "dist = x * x + y * y", LineStart: 7, LineEnd: 7},
				{Title: "Max-heap for k smallest", Explanation: "Negate distance, pop largest when over k", CodeSnippet: "heapq.heappush(heap, (-dist, [x, y]))\nif len(heap) > k:", LineStart: 8, LineEnd: 10},
				{Title: "Extract points", Explanation: "Return points from heap", CodeSnippet: "return [point for _, point in heap]", LineStart: 12, LineEnd: 12},
			},
		},
	},
	{
		ID:              "task-scheduler",
		Number:          74,
		Title:           "Task Scheduler",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Array", "Hash Table", "Greedy", "Sorting", "Heap", "Counting"},
		RelatedChapters: []int{7, 8, 11, 12},
		Description: `You are given an array of CPU tasks, each represented by letters A to Z, and a cooling interval n. Each cycle or interval allows the completion of one task. Tasks can be completed in any order, but there's a constraint: identical tasks must be separated by at least n intervals due to cooling time.

Return the minimum number of intervals required to complete all tasks.`,
		Constraints: []string{
			"1 <= tasks.length <= 10^4",
			"tasks[i] is an uppercase English letter",
			"0 <= n <= 100",
		},
		Examples: []Example{
			{Input: `tasks = ["A","A","A","B","B","B"], n = 2`, Output: "8", Explanation: "A -> B -> idle -> A -> B -> idle -> A -> B"},
			{Input: `tasks = ["A","A","A","B","B","B"], n = 0`, Output: "6"},
			{Input: `tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2`, Output: "16"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"tasks": []string{"A", "A", "A", "B", "B", "B"}, "n": 2}, Expected: 8},
			{Input: map[string]interface{}{"tasks": []string{"A", "A", "A", "B", "B", "B"}, "n": 0}, Expected: 6},
		},
		TimeComplexity:  "O(n * m)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def leastInterval(tasks, n):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Max-heap for task counts. Each cycle: pop up to n+1 tasks, decrement, track time for cooldown.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Greedy: always pick task with most remaining. After cooldown, re-add tasks with remaining count."},
			{Level: 3, Type: "code", Content: "heap with -count. Each round: pop n+1 tasks, add to cooldown queue. After n+1 cycles, re-add from queue."},
		},
		Solution: Solution{
			Code: `from collections import Counter, deque
import heapq

def leastInterval(tasks, n):
    count = Counter(tasks)
    heap = [-cnt for cnt in count.values()]
    heapq.heapify(heap)

    time = 0
    queue = deque()  # (time_available, count)

    while heap or queue:
        time += 1

        if heap:
            cnt = heapq.heappop(heap) + 1  # Decrement (negative)
            if cnt != 0:
                queue.append((time + n, cnt))

        if queue and queue[0][0] == time:
            heapq.heappush(heap, queue.popleft()[1])

    return time`,
			Explanation:     "Max-heap tracks remaining counts. Queue tracks cooling tasks. Each tick: run a task, re-add after cooldown.",
			TimeComplexity:  "O(total_time * log 26) = O(n)",
			SpaceComplexity: "O(26) = O(1) for task counts",
			Walkthrough: []WalkthroughStep{
				{Title: "Count frequencies", Explanation: "Max-heap of task counts", CodeSnippet: "heap = [-cnt for cnt in count.values()]", LineStart: 6, LineEnd: 6},
				{Title: "Process each tick", Explanation: "Decrement count, add to cooldown queue", CodeSnippet: "if cnt != 0:\n    queue.append((time + n, cnt))", LineStart: 17, LineEnd: 18},
				{Title: "Re-add after cooldown", Explanation: "Task available again", CodeSnippet: "if queue and queue[0][0] == time:\n    heapq.heappush(heap, queue.popleft()[1])", LineStart: 20, LineEnd: 21},
			},
		},
	},
	{
		ID:              "find-median-from-data-stream",
		Number:          75,
		Title:           "Find Median from Data Stream",
		Difficulty:      "Hard",
		Category:        "heap-priority-queue",
		Tags:            []string{"Two Pointers", "Design", "Sorting", "Heap", "Data Stream"},
		RelatedChapters: []int{7, 11, 12},
		Description: `The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

Implement the MedianFinder class:
- MedianFinder() initializes the MedianFinder object.
- void addNum(int num) adds the integer num from the data stream to the data structure.
- double findMedian() returns the median of all elements so far.`,
		Constraints: []string{
			"-10^5 <= num <= 10^5",
			"There will be at least one element in the data structure before calling findMedian",
			"At most 5 * 10^4 calls will be made to addNum and findMedian",
		},
		Examples: []Example{
			{Input: `["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]`, Output: "[null,null,null,1.5,null,2.0]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"}, "values": [][]int{{}, {1}, {2}, {}, {3}, {}}}, Expected: []interface{}{nil, nil, nil, 1.5, nil, 2.0}},
		},
		TimeComplexity:  "O(log n) add, O(1) median",
		SpaceComplexity: "O(n)",
		StarterCode: `class MedianFinder:
    def __init__(self):
        pass

    def addNum(self, num: int) -> None:
        pass

    def findMedian(self) -> float:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Two heaps: max-heap for smaller half, min-heap for larger half. Median is from top of heaps.", ChapterRef: 12},
			{Level: 2, Type: "algorithm", Content: "Keep heaps balanced (size differs by at most 1). Add to small, move max to large, rebalance if needed."},
			{Level: 3, Type: "code", Content: "small = max-heap, large = min-heap. Add to small, push -heappop(small) to large. If large bigger, push -heappop(large) to small."},
		},
		Solution: Solution{
			Code: `import heapq

class MedianFinder:
    def __init__(self):
        self.small = []  # Max-heap (negated)
        self.large = []  # Min-heap

    def addNum(self, num: int) -> None:
        # Add to max-heap (small)
        heapq.heappush(self.small, -num)

        # Move largest from small to large
        heapq.heappush(self.large, -heapq.heappop(self.small))

        # Balance: small should have >= elements than large
        if len(self.large) > len(self.small):
            heapq.heappush(self.small, -heapq.heappop(self.large))

    def findMedian(self) -> float:
        if len(self.small) > len(self.large):
            return -self.small[0]
        return (-self.small[0] + self.large[0]) / 2`,
			Explanation:     "Two heaps partition data. Small (max-heap) has smaller half, large (min-heap) has larger half. Median from tops.",
			TimeComplexity:  "O(log n) for addNum, O(1) for findMedian",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Two heaps", Explanation: "Max-heap for small half, min-heap for large half", CodeSnippet: "self.small = []  # Max-heap (negated)\nself.large = []  # Min-heap", LineStart: 5, LineEnd: 6},
				{Title: "Add and rebalance", Explanation: "Always add to small, move to large, rebalance", CodeSnippet: "heapq.heappush(self.small, -num)\nheapq.heappush(self.large, -heapq.heappop(self.small))", LineStart: 10, LineEnd: 13},
				{Title: "Get median", Explanation: "Middle from small's max or average of both tops", CodeSnippet: "return (-self.small[0] + self.large[0]) / 2", LineStart: 22, LineEnd: 22},
			},
		},
	},
	{
		ID:              "design-twitter",
		Number:          136,
		Title:           "Design Twitter",
		Difficulty:      "Medium",
		Category:        "heap-priority-queue",
		Tags:            []string{"Design", "Heap", "Hash Table"},
		RelatedChapters: []int{7, 11, 12},
		Description: `Design a simplified version of Twitter where users can post tweets, follow/unfollow another user, and is able to see the 10 most recent tweets in the user's news feed.

Implement the Twitter class:
- Twitter() Initializes your twitter object.
- void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId.
- List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed.
- void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID followeeId.
- void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID followeeId.`,
		Constraints: []string{
			"1 <= userId, followerId, followeeId <= 500",
			"0 <= tweetId <= 10^4",
			"All the tweets have unique IDs",
			"At most 3 * 10^4 calls will be made to postTweet, getNewsFeed, follow, and unfollow",
		},
		Examples: []Example{
			{Input: `["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]`, Output: `[null, null, [5], null, null, [6, 5], null, [5]]`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"Twitter", "postTweet", "getNewsFeed"}, "args": [][]int{{}, {1, 5}, {1}}}, Expected: []interface{}{nil, nil, []int{5}}},
		},
		TimeComplexity:  "O(n log n) for getNewsFeed",
		SpaceComplexity: "O(users * tweets + follows)",
		StarterCode:     "class Twitter:\n    def __init__(self):\n        pass\n\n    def postTweet(self, userId, tweetId):\n        pass\n\n    def getNewsFeed(self, userId):\n        pass\n\n    def follow(self, followerId, followeeId):\n        pass\n\n    def unfollow(self, followerId, followeeId):\n        pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use hash maps for tweets and follows. Heap for merging feeds.", ChapterRef: 7},
			{Level: 2, Type: "algorithm", Content: "Store tweets with timestamps. GetNewsFeed: merge k sorted lists using min-heap."},
			{Level: 3, Type: "code", Content: "tweets = defaultdict(list). follows = defaultdict(set). Use heap with (time, tweetId, userId, index)."},
		},
		Solution: Solution{
			Code: `class Twitter:
    def __init__(self):
        from collections import defaultdict
        import heapq
        self.time = 0
        self.tweets = defaultdict(list)  # userId -> [(time, tweetId)]
        self.follows = defaultdict(set)  # userId -> set of followeeIds

    def postTweet(self, userId, tweetId):
        self.tweets[userId].append((self.time, tweetId))
        self.time += 1

    def getNewsFeed(self, userId):
        import heapq
        heap = []
        self.follows[userId].add(userId)  # Include own tweets

        for followeeId in self.follows[userId]:
            tweets = self.tweets[followeeId]
            if tweets:
                idx = len(tweets) - 1
                time, tweetId = tweets[idx]
                heapq.heappush(heap, (-time, tweetId, followeeId, idx))

        result = []
        while heap and len(result) < 10:
            time, tweetId, followeeId, idx = heapq.heappop(heap)
            result.append(tweetId)
            if idx > 0:
                idx -= 1
                time, tweetId = self.tweets[followeeId][idx]
                heapq.heappush(heap, (-time, tweetId, followeeId, idx))

        return result

    def follow(self, followerId, followeeId):
        self.follows[followerId].add(followeeId)

    def unfollow(self, followerId, followeeId):
        self.follows[followerId].discard(followeeId)`,
			Explanation:     "Merge k sorted lists (each user's tweets) using heap to get 10 most recent.",
			TimeComplexity:  "O(n log n) for getNewsFeed",
			SpaceComplexity: "O(users * tweets + follows)",
			Walkthrough: []WalkthroughStep{
				{Title: "Timestamp tweets", Explanation: "Track posting order globally", CodeSnippet: "self.tweets[userId].append((self.time, tweetId))\nself.time += 1", LineStart: 10, LineEnd: 11},
				{Title: "Merge with heap", Explanation: "K-way merge of sorted lists", CodeSnippet: "heapq.heappush(heap, (-time, tweetId, followeeId, idx))", LineStart: 23, LineEnd: 23},
				{Title: "Get next tweet", Explanation: "Move to previous tweet in same user's list", CodeSnippet: "if idx > 0:\n    idx -= 1", LineStart: 29, LineEnd: 30},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, HeapProblems...)
}
