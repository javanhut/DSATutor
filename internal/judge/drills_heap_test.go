package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestHeapDrillsRejectMistakes checks that each heap foundations drill fails
// the mistake its tests and notes promise to catch, so a drill can't be
// passed without doing the thing it drills.
func TestHeapDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-build-min-heap", "uses child // 2 as the parent index", `class MinHeap:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)
        child = len(self.items) - 1
        while child > 0:
            parent = child // 2
            if self.items[parent] <= self.items[child]:
                break
            self.items[parent], self.items[child] = self.items[child], self.items[parent]
            child = parent

    def peek(self):
        return self.items[0] if self.items else -1

    def size(self):
        return len(self.items)

    def toList(self):
        return list(self.items)
`},
		{"drill-build-min-heap", "swaps only once instead of climbing", `class MinHeap:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)
        child = len(self.items) - 1
        if child > 0:
            parent = (child - 1) // 2
            if self.items[parent] <= self.items[child]:
                break
            self.items[parent], self.items[child] = self.items[child], self.items[parent]

    def peek(self):
        return self.items[0] if self.items else -1

    def size(self):
        return len(self.items)

    def toList(self):
        return list(self.items)
`},
		{"drill-build-min-heap", "keeps the array sorted instead of sifting up", `import bisect

class MinHeap:
    def __init__(self):
        self.items = []

    def push(self, val):
        bisect.insort(self.items, val)

    def peek(self):
        return self.items[0] if self.items else -1

    def size(self):
        return len(self.items)

    def toList(self):
        return list(self.items)
`},
		{"drill-build-min-heap", "peek returns the last item", `class MinHeap:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)
        child = len(self.items) - 1
        while child > 0:
            parent = (child - 1) // 2
            if self.items[parent] <= self.items[child]:
                break
            self.items[parent], self.items[child] = self.items[child], self.items[parent]
            child = parent

    def peek(self):
        return self.items[-1] if self.items else -1

    def size(self):
        return len(self.items)

    def toList(self):
        return list(self.items)
`},
		{"drill-heap-pop", "sifts down into the larger child", `class MinHeap:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)
        child = len(self.items) - 1
        while child > 0:
            parent = (child - 1) // 2
            if self.items[parent] <= self.items[child]:
                break
            self.items[parent], self.items[child] = self.items[child], self.items[parent]
            child = parent

    def peek(self):
        return self.items[0] if self.items else -1

    def size(self):
        return len(self.items)

    def toList(self):
        return list(self.items)

    def pop(self):
        if not self.items:
            return -1
        top = self.items[0]
        last = self.items.pop()
        if self.items:
            self.items[0] = last
            self._sift_down(0)
        return top

    def _sift_down(self, index):
        n = len(self.items)
        while True:
            child = 2 * index + 1
            if child >= n:
                break
            if child + 1 < n and self.items[child + 1] > self.items[child]:
                child += 1
            if self.items[index] <= self.items[child]:
                break
            self.items[index], self.items[child] = self.items[child], self.items[index]
            index = child
`},
		{"drill-heap-pop", "breaks a tie to the right", `class MinHeap:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)
        child = len(self.items) - 1
        while child > 0:
            parent = (child - 1) // 2
            if self.items[parent] <= self.items[child]:
                break
            self.items[parent], self.items[child] = self.items[child], self.items[parent]
            child = parent

    def peek(self):
        return self.items[0] if self.items else -1

    def size(self):
        return len(self.items)

    def toList(self):
        return list(self.items)

    def pop(self):
        if not self.items:
            return -1
        top = self.items[0]
        last = self.items.pop()
        if self.items:
            self.items[0] = last
            self._sift_down(0)
        return top

    def _sift_down(self, index):
        n = len(self.items)
        while True:
            child = 2 * index + 1
            if child >= n:
                break
            if child + 1 < n and self.items[child + 1] <= self.items[child]:
                child += 1
            if self.items[index] <= self.items[child]:
                break
            self.items[index], self.items[child] = self.items[child], self.items[index]
            index = child
`},
		{"drill-heap-pop", "always follows the left child", `class MinHeap:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)
        child = len(self.items) - 1
        while child > 0:
            parent = (child - 1) // 2
            if self.items[parent] <= self.items[child]:
                break
            self.items[parent], self.items[child] = self.items[child], self.items[parent]
            child = parent

    def peek(self):
        return self.items[0] if self.items else -1

    def size(self):
        return len(self.items)

    def toList(self):
        return list(self.items)

    def pop(self):
        if not self.items:
            return -1
        top = self.items[0]
        last = self.items.pop()
        if self.items:
            self.items[0] = last
            self._sift_down(0)
        return top

    def _sift_down(self, index):
        n = len(self.items)
        while True:
            child = 2 * index + 1
            if child >= n:
                break
            if self.items[index] <= self.items[child]:
                break
            self.items[index], self.items[child] = self.items[child], self.items[index]
            index = child
`},
		{"drill-heap-pop", "reads the right child without checking it exists", `class MinHeap:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)
        child = len(self.items) - 1
        while child > 0:
            parent = (child - 1) // 2
            if self.items[parent] <= self.items[child]:
                break
            self.items[parent], self.items[child] = self.items[child], self.items[parent]
            child = parent

    def peek(self):
        return self.items[0] if self.items else -1

    def size(self):
        return len(self.items)

    def toList(self):
        return list(self.items)

    def pop(self):
        if not self.items:
            return -1
        top = self.items[0]
        last = self.items.pop()
        if self.items:
            self.items[0] = last
            self._sift_down(0)
        return top

    def _sift_down(self, index):
        n = len(self.items)
        while True:
            child = 2 * index + 1
            if child >= n:
                break
            if self.items[child + 1] < self.items[child]:
                child += 1
            if self.items[index] <= self.items[child]:
                break
            self.items[index], self.items[child] = self.items[child], self.items[index]
            index = child
`},
		{"drill-heap-pop", "removes the root with pop(0) and never sifts", `class MinHeap:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)
        child = len(self.items) - 1
        while child > 0:
            parent = (child - 1) // 2
            if self.items[parent] <= self.items[child]:
                break
            self.items[parent], self.items[child] = self.items[child], self.items[parent]
            child = parent

    def peek(self):
        return self.items[0] if self.items else -1

    def size(self):
        return len(self.items)

    def toList(self):
        return list(self.items)

    def pop(self):
        if not self.items:
            return -1
        return self.items.pop(0)
`},
		{"drill-heap-pop", "returns the moved last item instead of the root", `class MinHeap:
    def __init__(self):
        self.items = []

    def push(self, val):
        self.items.append(val)
        child = len(self.items) - 1
        while child > 0:
            parent = (child - 1) // 2
            if self.items[parent] <= self.items[child]:
                break
            self.items[parent], self.items[child] = self.items[child], self.items[parent]
            child = parent

    def peek(self):
        return self.items[0] if self.items else -1

    def size(self):
        return len(self.items)

    def toList(self):
        return list(self.items)

    def pop(self):
        if not self.items:
            return -1
        top = self.items[0]
        last = self.items.pop()
        if self.items:
            self.items[0] = last
            self._sift_down(0)
        return last

    def _sift_down(self, index):
        n = len(self.items)
        while True:
            child = 2 * index + 1
            if child >= n:
                break
            if child + 1 < n and self.items[child + 1] < self.items[child]:
                child += 1
            if self.items[index] <= self.items[child]:
                break
            self.items[index], self.items[child] = self.items[child], self.items[index]
            index = child
`},
		{"drill-heapify", "starts one parent too early (n // 2 - 2)", `def heapify(nums):
    n = len(nums)
    for start in range(n // 2 - 2, -1, -1):
        index = start
        while True:
            child = 2 * index + 1
            if child >= n:
                break
            if child + 1 < n and nums[child + 1] < nums[child]:
                child += 1
            if nums[index] <= nums[child]:
                break
            nums[index], nums[child] = nums[child], nums[index]
            index = child
    return nums`},
		{"drill-heapify", "sifts from the root downwards in index order", `def heapify(nums):
    n = len(nums)
    for start in range(0, n // 2):
        index = start
        while True:
            child = 2 * index + 1
            if child >= n:
                break
            if child + 1 < n and nums[child + 1] < nums[child]:
                child += 1
            if nums[index] <= nums[child]:
                break
            nums[index], nums[child] = nums[child], nums[index]
            index = child
    return nums`},
		{"drill-heapify", "calls heapq.heapify, which breaks ties differently", `import heapq

def heapify(nums):
    heapq.heapify(nums)
    return nums`},
		{"drill-heapify", "builds by pushing one item at a time", `def heapify(nums):
    out = []
    for val in nums:
        out.append(val)
        child = len(out) - 1
        while child > 0:
            parent = (child - 1) // 2
            if out[parent] <= out[child]:
                break
            out[parent], out[child] = out[child], out[parent]
            child = parent
    return out`},
		{"drill-heapify", "just sorts the array", `def heapify(nums):
    nums.sort()
    return nums`},
		{"drill-heapify", "breaks a tie to the right", `def heapify(nums):
    n = len(nums)
    for start in range(n // 2 - 1, -1, -1):
        index = start
        while True:
            child = 2 * index + 1
            if child >= n:
                break
            if child + 1 < n and nums[child + 1] <= nums[child]:
                child += 1
            if nums[index] <= nums[child]:
                break
            nums[index], nums[child] = nums[child], nums[index]
            index = child
    return nums`},
		{"drill-is-min-heap", "checks that the array is sorted", `def isMinHeap(arr):
    for i in range(len(arr) - 1):
        if arr[i] > arr[i + 1]:
            return False
    return True`},
		{"drill-is-min-heap", "checks only the left child", `def isMinHeap(arr):
    for parent in range(len(arr)):
        left = 2 * parent + 1
        if left < len(arr) and arr[parent] > arr[left]:
            return False
    return True`},
		{"drill-is-min-heap", "stops one index early", `def isMinHeap(arr):
    for child in range(1, len(arr) - 1):
        if arr[(child - 1) // 2] > arr[child]:
            return False
    return True`},
		{"drill-is-min-heap", "rejects equal parent and child", `def isMinHeap(arr):
    for child in range(1, len(arr)):
        if arr[(child - 1) // 2] >= arr[child]:
            return False
    return True`},
		{"drill-heap-sort", "uses an O(n^2) selection sort", `def heapSort(nums):
    n = len(nums)
    for i in range(n):
        smallest = i
        for j in range(i + 1, n):
            if nums[j] < nums[smallest]:
                smallest = j
        nums[i], nums[smallest] = nums[smallest], nums[i]
    return nums`},
		{"drill-heap-sort", "sifts over the whole array, disturbing the sorted suffix", `def heapSort(nums):
    n = len(nums)

    def sift_down(index, size):
        while True:
            child = 2 * index + 1
            if child >= size:
                return
            if child + 1 < size and nums[child + 1] > nums[child]:
                child += 1
            if nums[index] >= nums[child]:
                return
            nums[index], nums[child] = nums[child], nums[index]
            index = child

    for start in range(n // 2 - 1, -1, -1):
        sift_down(start, n)
    for end in range(n - 1, 0, -1):
        nums[0], nums[end] = nums[end], nums[0]
        sift_down(0, n)
    return nums`},
		{"drill-heap-sort", "returns descending order", `import heapq

def heapSort(nums):
    heap = [-x for x in nums]
    heapq.heapify(heap)
    return [-heapq.heappop(heap) for _ in range(len(nums))]`},
		{"drill-heap-sort", "stops the extraction one item early", `def heapSort(nums):
    n = len(nums)

    def sift_down(index, size):
        while True:
            child = 2 * index + 1
            if child >= size:
                return
            if child + 1 < size and nums[child + 1] > nums[child]:
                child += 1
            if nums[index] >= nums[child]:
                return
            nums[index], nums[child] = nums[child], nums[index]
            index = child

    for start in range(n // 2 - 1, -1, -1):
        sift_down(start, n)
    for end in range(n - 1, 1, -1):
        nums[0], nums[end] = nums[end], nums[0]
        sift_down(0, end)
    return nums`},
		{"drill-k-largest-heapq", "forgets the negation and returns the smallest", `import heapq

def kLargest(nums, k):
    heap = list(nums)
    heapq.heapify(heap)
    return [heapq.heappop(heap) for _ in range(k)]`},
		{"drill-k-largest-heapq", "leaves the popped values negated", `import heapq

def kLargest(nums, k):
    heap = [-x for x in nums]
    heapq.heapify(heap)
    return [heapq.heappop(heap) for _ in range(k)]`},
		{"drill-k-largest-heapq", "slices the heap array instead of popping", `import heapq

def kLargest(nums, k):
    heap = [-x for x in nums]
    heapq.heapify(heap)
    return [-x for x in heap[:k]]`},
		{"drill-k-largest-heapq", "returns ascending order", `import heapq

def kLargest(nums, k):
    heap = [-x for x in nums]
    heapq.heapify(heap)
    return sorted(-heapq.heappop(heap) for _ in range(k))`},
		{"drill-merge-sorted-arrays", "reads arr[0] of an empty array", `import heapq

def mergeSortedArrays(arrays):
    heap = [(arr[0], i, 0) for i, arr in enumerate(arrays)]
    heapq.heapify(heap)
    merged = []
    while heap:
        val, i, j = heapq.heappop(heap)
        merged.append(val)
        if j + 1 < len(arrays[i]):
            heapq.heappush(heap, (arrays[i][j + 1], i, j + 1))
    return merged`},
		{"drill-merge-sorted-arrays", "never pushes the next value of the array", `import heapq

def mergeSortedArrays(arrays):
    heap = [(arr[0], i, 0) for i, arr in enumerate(arrays) if arr]
    heapq.heapify(heap)
    merged = []
    while heap:
        val, i, j = heapq.heappop(heap)
        merged.append(val)
    return merged`},
		{"drill-merge-sorted-arrays", "orders by array index before value", `import heapq

def mergeSortedArrays(arrays):
    heap = [(i, arr[0], 0) for i, arr in enumerate(arrays) if arr]
    heapq.heapify(heap)
    merged = []
    while heap:
        i, val, j = heapq.heappop(heap)
        merged.append(val)
        if j + 1 < len(arrays[i]):
            heapq.heappush(heap, (i, arrays[i][j + 1], j + 1))
    return merged`},
		{"drill-merge-sorted-arrays", "stops as soon as any array runs out", `import heapq

def mergeSortedArrays(arrays):
    heap = [(arr[0], i, 0) for i, arr in enumerate(arrays) if arr]
    heapq.heapify(heap)
    merged = []
    while heap:
        val, i, j = heapq.heappop(heap)
        merged.append(val)
        if j + 1 >= len(arrays[i]):
            break
        heapq.heappush(heap, (arrays[i][j + 1], i, j + 1))
    return merged`},
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
