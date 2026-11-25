package practice

import "math/rand"

// SlidingWindowDrills are the sliding window foundations: slide a fixed window
// by adding what enters and removing what leaves (sums, then the best average,
// then a dictionary of counts), then the variable window in both of its shapes
// (shrink while valid for the shortest, shrink until valid for the longest) and
// counting the windows that end at each right edge, before the sliding window
// ladder applies them. The large cases carry a 5 second limit so that
// recomputing a window from scratch fails.
var SlidingWindowDrills = []*Problem{
	{
		ID:              "drill-window-sums",
		Number:          0,
		Kind:            KindDrill,
		Order:           1,
		Title:           "Fixed Window Sums",
		Slug:            "drill-window-sums",
		Difficulty:      "Easy",
		Category:        "sliding-window",
		Tags:            []string{"Sliding Window", "Array", "Prefix Sum", "Foundations"},
		RelatedChapters: []int{4},
		Description:     "Given a list of integers `nums` and a window size `k`, return the sum of every window of `k` neighbouring elements, from left to right.\n\nA **window** is a run of `k` elements next to each other: the first window is `nums[0..k-1]`, the next is `nums[1..k]`, and so on until the window ends at the last element. A list of `n` elements has `n - k + 1` windows, so the answer has that many numbers. Values can be negative or zero.\n\nThe slow way adds up `k` numbers for every window. Do it faster: **slide** the window one step by adding the element that enters on the right and subtracting the element that leaves on the left. Each window then costs a single add and a single subtract, however big `k` is.",
		Constraints:     []string{"1 <= k <= len(nums) <= 150000", "-10^4 <= nums[i] <= 10^4"},
		Examples: []Example{
			{Input: "nums = [1,3,-1,-3,5,3,6,7], k = 3", Output: "[3,-1,1,5,14,16]", Explanation: "The first window [1,3,-1] sums to 3. Sliding right adds -3 and drops 1, so 3 + (-3) - 1 = -1. Sliding again adds 5 and drops 3: -1 + 5 - 3 = 1. And so on to the last window [3,6,7] = 16."},
			{Input: "nums = [4,2,7], k = 3", Output: "[13]", Explanation: "k equals the length of the list, so there is just one window: 4 + 2 + 7."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 3, -1, -3, 5, 3, 6, 7}, "k": 3}, Expected: []int{3, -1, 1, 5, 14, 16}, Note: "example 1: six windows of three, including negative sums"},
			{Input: map[string]interface{}{"nums": []int{4, 2, 7}, "k": 3}, Expected: []int{13}, Note: "example 2: k equals the length, so there is exactly one window"},
			{Input: map[string]interface{}{"nums": []int{5, -2, 9, 1}, "k": 1}, Expected: []int{5, -2, 9, 1}, Hidden: true, Note: "k = 1: every element is its own window"},
			{Input: map[string]interface{}{"nums": []int{6}, "k": 1}, Expected: []int{6}, Hidden: true, Note: "a single element"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 2, 2, 2}, "k": 2}, Expected: []int{4, 4, 4, 4}, Hidden: true, Note: "all equal values: every window has the same sum"},
			{Input: map[string]interface{}{"nums": []int{-4, -1, -7, -3}, "k": 2}, Expected: []int{-5, -8, -10}, Hidden: true, Note: "all negative values"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 5, 0, 0, 0}, "k": 3}, Expected: []int{5, 5, 5, 0}, Hidden: true, Note: "zeros: the outgoing and incoming values can be 0"},
			{Input: map[string]interface{}{"nums": []int{10, -10, 10, -10, 10, -10, 10}, "k": 4}, Expected: []int{0, 0, 0, 0}, Hidden: true, Note: "sums that go up and down, ending on the last window"},
			{Input: map[string]interface{}{"nums": []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, "k": 4}, Expected: []int{9, 11, 19, 17, 22, 22, 16, 19}, Hidden: true, Note: "a longer list: the last window must be included (n - k + 1 windows in all)"},
			{Input: map[string]interface{}{"nums": swBigNums(120000), "k": 60000}, Expected: swWindowSumsExpected(swBigNums(120000), 60000), Hidden: true, Timeout: 5000, Note: "120000 numbers, k = 60000: recomputing each window's sum does about 3.6 billion additions and runs out of time; add-incoming/subtract-outgoing is O(n)"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1) extra (the answer list aside)",
		StarterCode:     "def windowSums(nums, k):\n    # Return the sum of every window of k neighbouring elements\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When the window moves one step right, how many of its k elements are the same as before, and which two differ?"},
			{Level: 2, Type: "insight", Content: "Only one element enters on the right and one leaves on the left; the other k - 1 stay. So new sum = old sum + entering - leaving, which costs O(1) instead of O(k). The first window has to be summed directly because there is no earlier window to slide from. See Chapter 4.", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Build the first window: sum the first k elements and record it.\n2. Slide: for each right from k to the end, the entering element is nums[right] and the leaving one is nums[right - k]. Add the first, subtract the second.\n3. Record: append the new sum to the answer.\n4. Return the answer."},
			{Level: 4, Type: "skeleton", Content: "def windowSums(nums, k):\n    window = sum(nums[:k])\n    sums = [window]\n    for right in range(k, len(nums)):\n        window += ____ - ____\n        sums.append(window)\n    return sums"},
		},
		Solution: Solution{
			Code:            "def windowSums(nums, k):\n    window = sum(nums[:k])\n    sums = [window]\n    for right in range(k, len(nums)):\n        window += nums[right] - nums[right - k]\n        sums.append(window)\n    return sums",
			Explanation:     "This is the plain fixed-size window: sum the first k elements once, then slide by adding the element that enters and subtracting the one that leaves. Every later sliding-window problem keeps some summary of the window and updates it this way instead of recomputing it; this drill is that update on its simplest summary, a sum.",
			BruteForce:      "For every starting index add up the next k elements: O(n * k). With n = 120000 and k = 60000 that is billions of additions. The waste is that neighbouring windows share k - 1 elements, and the sum of those is added again from scratch every time.",
			KeyInsight:      "Two neighbouring windows differ by exactly one element at each end. So the next sum is the old sum plus what came in minus what went out, in constant time.",
			WhyItWorks:      "Invariant: at the top of each loop iteration, window equals the sum of nums[right - k .. right - 1], the last k elements before right. It holds before the loop, because window is the sum of nums[0..k-1]. In the loop, nums[right] enters and nums[right - k] leaves, so after the update window is the sum of nums[right - k + 1 .. right], which is the next window. Appending it records one sum per window, and right visits n - k indexes after the first window, so the answer has n - k + 1 sums in all.",
			CommonMistakes:  []string{"Recomputing `sum(nums[i:i + k])` for each window. It is correct, but it costs O(n * k) and times out on large inputs. Update the sum instead.", "Subtracting the wrong outgoing element, such as `nums[right - k + 1]` (which is still inside the window) or `nums[right - 1]`. The element that leaves when `right` enters is `nums[right - k]`.", "Forgetting to subtract the outgoing element, so the 'window' sum keeps growing and covers everything seen so far.", "Starting the answer list empty and so leaving out the first window, or stopping the loop one index early and leaving out the last. The answer needs n - k + 1 sums.", "Starting the loop at `right = 0` after summing the first k elements, which adds the first elements a second time. The first slide enters index k."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1) extra (the answer list aside)",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the first window", Explanation: "There is nothing to slide from yet, so add the first k elements directly and record that sum.", CodeSnippet: "window = sum(nums[:k])\nsums = [window]", LineStart: 2, LineEnd: 3},
				{Title: "Slide the window one step", Explanation: "The window now ends at right. The element nums[right] enters and nums[right - k] leaves, so adjust the sum by the difference. On [1,3,-1,-3] with k = 3, right = 3 adds -3 and drops 1: 3 + (-3) - 1 = -1.", CodeSnippet: "window += nums[right] - nums[right - k]", LineStart: 5, LineEnd: 5},
				{Title: "Record the window", Explanation: "The window sum is now correct for the window that ends at right, so save it.", CodeSnippet: "sums.append(window)", LineStart: 6, LineEnd: 6},
				{Title: "Return every window", Explanation: "The loop visits each remaining end position once, so sums holds one value per window.", CodeSnippet: "return sums", LineStart: 7, LineEnd: 7},
			},
			Alternatives: []AltSolution{
				{Name: "Prefix sums", Code: "def windowSums(nums, k):\n    prefix = [0]\n    for value in nums:\n        prefix.append(prefix[-1] + value)\n    return [prefix[i + k] - prefix[i] for i in range(len(nums) - k + 1)]", TimeComplexity: "O(n)", SpaceComplexity: "O(n)", Explanation: "Prefix sums give any range sum in O(1): prefix[i + k] - prefix[i] is the sum of nums[i..i + k - 1]. It needs O(n) extra space; the sliding version keeps one number."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the loop start at right = k and not right = 0?", Answer: "Indexes 0 to k - 1 are the first window, which was summed directly. The first slide is the one that brings in nums[k] and drops nums[0]; starting earlier would add the first elements twice and index nums[right - k] with a negative number."},
			{Prompt: "What if the question asked for the maximum in every window, not the sum? Does the add-and-subtract trick still work?", Answer: "No. A sum can be undone by subtracting the leaving value, but a maximum cannot: if the leaving element was the maximum you do not know the next biggest one without looking at the others. That needs a different summary of the window, such as a monotonic deque (Sliding Window Maximum later in this chapter)."},
		},
	},
	{
		ID:              "drill-max-average-subarray",
		Number:          643,
		Kind:            KindDrill,
		Order:           2,
		Title:           "Maximum Average Subarray I",
		Slug:            "drill-max-average-subarray",
		Difficulty:      "Easy",
		Category:        "sliding-window",
		Tags:            []string{"Sliding Window", "Array", "Foundations"},
		RelatedChapters: []int{4},
		Description:     "Given a list of integers `nums` and a window size `k`, find the window of exactly `k` neighbouring elements with the highest **average**, and return that average as a decimal number.\n\nThe average of a window is its sum divided by `k`, so it can have a fractional part (the average of [1,2,3,4] taken three at a time can be 2.5): use `/` and not `//`. Answers within `1e-5` of the exact value are accepted. Values can be negative, so the answer can be negative too.\n\nAll windows have the same length `k`, so the window with the biggest average is the window with the biggest sum. Slide a window of sum as in the previous drill, remember the best sum, and divide once at the end.",
		Constraints:     []string{"1 <= k <= len(nums) <= 150000", "-10^4 <= nums[i] <= 10^4"},
		Examples: []Example{
			{Input: "nums = [1,12,-5,-6,50,3], k = 4", Output: "12.75", Explanation: "The windows of four sum to 2, 51 and 42. The best is [12,-5,-6,50] = 51, and 51 / 4 = 12.75."},
			{Input: "nums = [5], k = 1", Output: "5.0", Explanation: "One window with one element: its average is 5."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 12, -5, -6, 50, 3}, "k": 4}, Expected: 12.75, Note: "example 1: the best window is [12,-5,-6,50]"},
			{Input: map[string]interface{}{"nums": []int{5}, "k": 1}, Expected: 5.0, Note: "example 2: one element is its own average"},
			{Input: map[string]interface{}{"nums": []int{-3, -8, -2, -9}, "k": 2}, Expected: -5.0, Hidden: true, Note: "all negative: the answer is negative, so a best that starts at 0 is wrong"},
			{Input: map[string]interface{}{"nums": []int{7, 7, 7, 7}, "k": 3}, Expected: 7.0, Hidden: true, Note: "all equal values"},
			{Input: map[string]interface{}{"nums": []int{0, 4, 0, 3, 2}, "k": 1}, Expected: 4.0, Hidden: true, Note: "k = 1: the answer is the largest element"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5}, "k": 5}, Expected: 3.0, Hidden: true, Note: "k equals the length: the one window is the whole list"},
			{Input: map[string]interface{}{"nums": []int{9, 1, 1, 1, 1}, "k": 2}, Expected: 5.0, Hidden: true, Note: "the best window is the first one: do not skip it"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 1, 9}, "k": 2}, Expected: 5.0, Hidden: true, Note: "the best window is the last one: do not stop early"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}, "k": 3}, Expected: 3.0, Hidden: true, Note: "the average is not a whole number (2.5 is not 2), so divide with /"},
			{Input: map[string]interface{}{"nums": swBigNums(120000), "k": 60000}, Expected: swMaxAverageExpected(swBigNums(120000), 60000), Hidden: true, Timeout: 5000, Note: "120000 numbers, k = 60000: recomputing each window's sum runs out of time"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def findMaxAverage(nums, k):\n    # Return the largest average over all windows of k elements\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Every window has the same length k. What does that tell you about comparing their averages, and so what do you actually need to track?"},
			{Level: 2, Type: "insight", Content: "With equal lengths, the largest average belongs to the largest sum, so track the best window sum and divide by k once at the end. Compute each sum by sliding (add the element that enters, subtract the one that leaves), exactly as in Fixed Window Sums. The best must start from the first window's sum, not from 0, because every sum can be negative. See Chapter 4.", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Build the first window: sum the first k elements and make it the best so far.\n2. Slide: for each right from k on, add nums[right] and subtract nums[right - k].\n3. Compare: keep the larger of best and the new sum.\n4. Answer: return best divided by k."},
			{Level: 4, Type: "skeleton", Content: "def findMaxAverage(nums, k):\n    window = sum(nums[:k])\n    best = ____\n    for right in range(k, len(nums)):\n        window += ____ - ____\n        best = max(best, window)\n    return best / k"},
		},
		Solution: Solution{
			Code:            "def findMaxAverage(nums, k):\n    window = sum(nums[:k])\n    best = window\n    for right in range(k, len(nums)):\n        window += nums[right] - nums[right - k]\n        best = max(best, window)\n    return best / k",
			Explanation:     "This builds on Fixed Window Sums: the same slide-by-add-and-subtract, with one extra variable that remembers the best window seen. Comparing sums instead of averages is safe because every window has k elements, and dividing once at the end avoids doing a division per window.",
			BruteForce:      "Add up each window of k elements from scratch and keep the biggest average: O(n * k). The same shared elements are summed again for every neighbouring window.",
			KeyInsight:      "Windows of the same length are ordered the same way by sum and by average, so the problem is Fixed Window Sums plus a running maximum.",
			WhyItWorks:      "Invariant: after processing index right, window is the sum of nums[right - k + 1 .. right] and best is the largest sum among all windows ending at or before right. The first window is both, by construction. Each slide produces the next window's sum exactly as in the previous drill, and taking the max with best keeps the invariant. When the loop ends, best is the largest window sum, and because every window has k elements, dividing by k gives the largest average.",
			CommonMistakes:  []string{"Starting `best` at 0 (or at a small constant) instead of the first window's sum. When every window sum is negative, the answer would wrongly be 0.", "Using `//` and returning a whole number, so an average of 2.5 comes out as 2. Use `/`.", "Recomputing the sum of every window instead of sliding it: O(n * k), which is too slow on large inputs.", "Comparing the window against `best` only before sliding, so the last window is never checked, or skipping the first window. Compare the first window as well as every slid one."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the first window", Explanation: "Sum the first k elements and treat that sum as the best so far, because it is the only window seen.", CodeSnippet: "window = sum(nums[:k])\nbest = window", LineStart: 2, LineEnd: 3},
				{Title: "Slide the window", Explanation: "Add the element entering at right and subtract the one leaving, nums[right - k], so window is always the sum of the last k elements.", CodeSnippet: "window += nums[right] - nums[right - k]", LineStart: 5, LineEnd: 5},
				{Title: "Keep the best sum", Explanation: "Windows have equal length, so the largest sum is the largest average. On [1,12,-5,-6,50,3] with k = 4 the sums go 2, 51, 42 and best ends at 51.", CodeSnippet: "best = max(best, window)", LineStart: 6, LineEnd: 6},
				{Title: "Divide once", Explanation: "Turn the best sum into an average with real division.", CodeSnippet: "return best / k", LineStart: 7, LineEnd: 7},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to compare window sums instead of averages?", Answer: "Every window has exactly k elements, so its average is its sum divided by the same k. Dividing by a positive constant keeps the order, so the biggest sum is the biggest average. (With windows of different lengths this would be false.)"},
			{Prompt: "What if the window length could vary, so you wanted the best average over any subarray of length at least k?", Answer: "Comparing sums would no longer work, because a longer window has a bigger sum without a better average. That problem needs a different idea, such as binary searching on the answer, which is beyond this drill."},
		},
	},
	{
		ID:              "drill-distinct-in-windows",
		Number:          0,
		Kind:            KindDrill,
		Order:           3,
		Title:           "Distinct Values in Every Window",
		Slug:            "drill-distinct-in-windows",
		Difficulty:      "Easy",
		Category:        "sliding-window",
		Tags:            []string{"Sliding Window", "Hash Map", "Array", "Foundations"},
		RelatedChapters: []int{4},
		Description:     "Given a list of integers `nums` and a window size `k`, return a list with the number of **distinct** values in each window of `k` neighbouring elements, from left to right. A value that appears twice in a window counts once.\n\nYou cannot just add and subtract a number this time: when an element leaves the window you do not know whether another copy of it is still inside. So keep a dictionary `counts` mapping each value to **how many copies are in the window right now**:\n\n- when a value enters, add 1 to its count;\n- when a value leaves, subtract 1 from its count, and if the count reaches 0, **delete the key**, so that the number of keys, `len(counts)`, is always the number of distinct values in the window.\n\nThe answer has `n - k + 1` numbers.",
		Constraints:     []string{"1 <= k <= len(nums) <= 100000", "-10^9 <= nums[i] <= 10^9"},
		Examples: []Example{
			{Input: "nums = [1,2,1,3,4,2,3], k = 4", Output: "[3,4,4,3]", Explanation: "[1,2,1,3] has the values 1, 2 and 3: three distinct. [2,1,3,4] has four. [1,3,4,2] has four. [3,4,2,3] has 3, 4 and 2: three. The repeated 1 and 3 count once."},
			{Input: "nums = [5,5,5], k = 2", Output: "[1,1]", Explanation: "Both windows are [5,5]: one distinct value, with a count of 2 in the dictionary."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 1, 3, 4, 2, 3}, "k": 4}, Expected: []int{3, 4, 4, 3}, Note: "example 1: windows of four with repeats inside"},
			{Input: map[string]interface{}{"nums": []int{5, 5, 5}, "k": 2}, Expected: []int{1, 1}, Note: "example 2: all equal, so every window has one distinct value"},
			{Input: map[string]interface{}{"nums": []int{7}, "k": 1}, Expected: []int{1}, Hidden: true, Note: "a single element"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}, "k": 1}, Expected: []int{1, 1, 1, 1}, Hidden: true, Note: "k = 1: every window has exactly one distinct value"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}, "k": 4}, Expected: []int{4}, Hidden: true, Note: "k equals the length and all values are distinct"},
			{Input: map[string]interface{}{"nums": []int{2, 1, 2, 1, 2, 1}, "k": 3}, Expected: []int{2, 2, 2, 2}, Hidden: true, Note: "a value that leaves while another copy is still inside: the count must not drop"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 2, 2, 3, 3}, "k": 2}, Expected: []int{1, 2, 1, 2, 1}, Hidden: true, Note: "a value that leaves and does not return: its key must be deleted so the count drops"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 4, 0, 0}, "k": 2}, Expected: []int{1, 2, 2, 1}, Hidden: true, Note: "zeros are real values"},
			{Input: map[string]interface{}{"nums": []int{-1, -1, 2, -1, 3, 3}, "k": 3}, Expected: []int{2, 2, 3, 2}, Hidden: true, Note: "negative values"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 1, 2, 3, 1, 2, 3}, "k": 5}, Expected: []int{3, 3, 3, 3, 3}, Hidden: true, Note: "a repeating pattern: each window drops one value and gains one"},
			{Input: map[string]interface{}{"nums": swBigDistinctNums(100000), "k": 50000}, Expected: swDistinctExpected(swBigDistinctNums(100000), 50000), Hidden: true, Timeout: 5000, Note: "100000 numbers, k = 50000: rebuilding a set for every window does billions of steps and runs out of time; a count dictionary is O(n)"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(k)",
		StarterCode:     "def countDistinctWindows(nums, k):\n    # Return the number of distinct values in each window of k elements\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "When the leftmost value leaves the window, how can you tell whether it is gone for good or another copy of it is still inside?"},
			{Level: 2, Type: "insight", Content: "Keep a count of each value in the window. A leaving value drops by one, and only when its count reaches 0 is it really gone, so delete that key at 0 so that len(counts) is the number of distinct values. Entering adds one. This is the 'window as a frequency dictionary' form of the pattern, see Chapter 4.", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Grow: add the new value nums[right] to counts.\n2. Shrink: once the window is longer than k (right >= k), subtract one from nums[right - k], and delete its key if the count is 0.\n3. Record: once the first full window exists (right >= k - 1), append len(counts).\n4. Return the answers."},
			{Level: 4, Type: "skeleton", Content: "def countDistinctWindows(nums, k):\n    counts = {}\n    result = []\n    for right, value in enumerate(nums):\n        counts[value] = counts.get(value, 0) + 1\n        if right >= k:\n            leaving = nums[right - k]\n            counts[leaving] -= 1\n            if ____:\n                del counts[leaving]\n        if right >= ____:\n            result.append(len(counts))\n    return result"},
		},
		Solution: Solution{
			Code:            "def countDistinctWindows(nums, k):\n    counts = {}\n    result = []\n    for right, value in enumerate(nums):\n        counts[value] = counts.get(value, 0) + 1\n        if right >= k:\n            leaving = nums[right - k]\n            counts[leaving] -= 1\n            if counts[leaving] == 0:\n                del counts[leaving]\n        if right >= k - 1:\n            result.append(len(counts))\n    return result",
			Explanation:     "Fixed Window Sums kept one number for the window; here the summary is a dictionary of counts, because a leaving value may still have copies inside. The add-in, remove-out rhythm is the same, and the only new rule is to delete a key when its count reaches 0. Longest Substring Without Repeating Characters and Permutation in String in the ladder are built on this dictionary.",
			BruteForce:      "Build `len(set(nums[i:i + k]))` for every window: O(n * k), because each set is built from scratch. Neighbouring windows share k - 1 values and rebuild their set again each time.",
			KeyInsight:      "Store how many copies of each value are inside the window. A value is in the window exactly when its count is above 0, and deleting zero counts makes the dictionary's size the answer.",
			WhyItWorks:      "Invariant: after the update at index right, counts[v] is the number of copies of v in the last k elements ending at right (or in all elements so far, while fewer than k have been seen), and no key has a count of 0. Entering nums[right] adds one copy. Once the window would hold k + 1 elements, nums[right - k] leaves, which removes exactly one copy, and a key is deleted the moment its count would be 0. Therefore the keys are exactly the values present, and len(counts) is the number of distinct values. From right = k - 1 the window is full, so each record is for one window.",
			CommonMistakes:  []string{"Leaving keys at a count of 0 in the dictionary. Then `len(counts)` also counts values that have already left the window. Delete the key when its count reaches 0.", "Using a set instead of counts, and removing a value from it when it leaves. If another copy is still inside the window, the value is wrongly dropped.", "Rebuilding `set(nums[i:i + k])` for every window. It is correct but O(n * k) and times out on large inputs.", "Removing the wrong element, such as `nums[right - k + 1]` (still inside) instead of `nums[right - k]`.", "Appending the answer before the first full window exists (`right >= k - 1` is the first full window), which adds extra entries for the short windows at the start."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(k)",
			Walkthrough: []WalkthroughStep{
				{Title: "Grow the window", Explanation: "The value at right enters: add one to its count (the first time it is seen, start from 0).", CodeSnippet: "counts[value] = counts.get(value, 0) + 1", LineStart: 5, LineEnd: 5},
				{Title: "Shrink to size k", Explanation: "Once right >= k the window holds k + 1 values, so the value at right - k must leave: take one copy away.", CodeSnippet: "leaving = nums[right - k]\ncounts[leaving] -= 1", LineStart: 7, LineEnd: 8},
				{Title: "Delete a key at zero", Explanation: "If that was the last copy, remove the key; otherwise len(counts) would still count a value that is no longer in the window. On [1,2,1,3] sliding to [2,1,3,4], one 1 leaves but another stays, so 1 keeps a count of 1.", CodeSnippet: "if counts[leaving] == 0:\n    del counts[leaving]", LineStart: 9, LineEnd: 10},
				{Title: "Record a full window", Explanation: "From right = k - 1 the window has k elements, so the number of keys is the number of distinct values.", CodeSnippet: "if right >= k - 1:\n    result.append(len(counts))", LineStart: 11, LineEnd: 12},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why delete the key when its count reaches 0, instead of leaving a 0 in the dictionary?", Answer: "len(counts) counts keys, and a key with a 0 would be counted even though no copy of the value is in the window. Deleting at 0 keeps 'the keys are exactly the values in the window', which is what makes len(counts) correct. (Alternatively, you could count the keys with a positive count, but that costs a scan.)"},
			{Prompt: "What if you only needed to know whether any value repeats within every window? Which variable would you keep?", Answer: "Track the number of values whose count is above 1, or compare len(counts) with k: a window of k elements has all-distinct values exactly when len(counts) == k. Either one updates in O(1) as values enter and leave."},
		},
	},
	{
		ID:              "drill-min-subarray-sum",
		Number:          209,
		Kind:            KindDrill,
		Order:           4,
		Title:           "Minimum Size Subarray Sum",
		Slug:            "drill-min-subarray-sum",
		Difficulty:      "Easy",
		Category:        "sliding-window",
		Tags:            []string{"Sliding Window", "Array", "Two Pointers", "Foundations"},
		RelatedChapters: []int{4},
		Description:     "You are given a positive integer `target` and a list `nums` of **positive** integers. Return the length of the shortest run of neighbouring elements whose sum is **at least** `target`. If no run reaches the target, return 0.\n\nThis is a **variable-size window**: its length is not given, so you grow and shrink it. The recipe, with the window `nums[left..right]` and its running `total`:\n\n1. **Grow**: move `right` one step and add `nums[right]` to `total`.\n2. **Shrink while valid**: as long as `total >= target`, the window is long enough, so record its length `right - left + 1`, then remove `nums[left]` from `total` and move `left` forward to see if a shorter window still reaches the target.\n\nIt has to be a loop, a `while`, because one large new element can allow several elements to leave at once. All numbers are positive, so removing an element always makes the total smaller, and that is what makes the shrink safe.",
		Constraints:     []string{"1 <= target <= 10^9", "1 <= len(nums) <= 100000", "1 <= nums[i] <= 10^4"},
		Examples: []Example{
			{Input: "target = 7, nums = [2,3,1,2,4,3]", Output: "2", Explanation: "The run [4,3] sums to 7 and has two elements. No single element reaches 7, so 2 is the shortest."},
			{Input: "target = 15, nums = [1,2,3,4,5]", Output: "5", Explanation: "The whole list sums to exactly 15, and every shorter run sums to less."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"target": 7, "nums": []int{2, 3, 1, 2, 4, 3}}, Expected: 2, Note: "example 1: the shortest window with sum >= 7 is [4,3]"},
			{Input: map[string]interface{}{"target": 15, "nums": []int{1, 2, 3, 4, 5}}, Expected: 5, Note: "example 2: only the whole list reaches 15"},
			{Input: map[string]interface{}{"target": 100, "nums": []int{1, 2, 3}}, Expected: 0, Hidden: true, Note: "the total never reaches the target: return 0, not a huge length"},
			{Input: map[string]interface{}{"target": 4, "nums": []int{4}}, Expected: 1, Hidden: true, Note: "one element equal to the target"},
			{Input: map[string]interface{}{"target": 5, "nums": []int{1, 1, 1, 1, 1, 1, 1, 1}}, Expected: 5, Hidden: true, Note: "all small values: the window must grow to exactly 5"},
			{Input: map[string]interface{}{"target": 100, "nums": []int{1, 1, 1, 100}}, Expected: 1, Hidden: true, Note: "a big value at the end: the window must shrink several times at once (a single `if` shrink gives 4, not 1)"},
			{Input: map[string]interface{}{"target": 8, "nums": []int{5, 1, 1, 1, 1, 8, 1}}, Expected: 1, Hidden: true, Note: "the single element 8 alone is the answer, found after a long window"},
			{Input: map[string]interface{}{"target": 11, "nums": []int{1, 2, 3, 4, 5}}, Expected: 3, Hidden: true, Note: "the answer is at the end of the list"},
			{Input: map[string]interface{}{"target": 3, "nums": []int{1, 2, 1, 1}}, Expected: 2, Hidden: true, Note: "a window can reach the target exactly: sum >= target, not sum > target"},
			{Input: map[string]interface{}{"target": 6, "nums": []int{10, 2, 3}}, Expected: 1, Hidden: true, Note: "the first element alone already exceeds the target"},
			{Input: map[string]interface{}{"target": 50000, "nums": swOnes(100000)}, Expected: 50000, Hidden: true, Timeout: 5000, Note: "100000 ones and target 50000: restarting from every index does billions of steps and runs out of time; the two-pointer window is O(n)"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def minSubArrayLen(target, nums):\n    # Return the length of the shortest run with sum >= target, or 0\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "The window sum needs to be at least target, and every number is positive. When the window is already big enough, what could you try to see if a shorter one works?"},
			{Level: 2, Type: "insight", Content: "Because every number is positive, dropping the leftmost element always lowers the sum, and adding one on the right always raises it. So grow right until the sum is enough, then shrink from the left for as long as it stays enough, noting the length each time. Both pointers only move forward, so it is O(n). This is a variable window that shrinks while valid; see Chapter 4.", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Grow: add nums[right] to the total.\n2. Shrink while valid: while total >= target, record right - left + 1, subtract nums[left], and move left forward.\n3. Track the minimum: keep the smallest length recorded.\n4. Answer: return 0 if nothing was recorded, else the minimum."},
			{Level: 4, Type: "skeleton", Content: "def minSubArrayLen(target, nums):\n    left = 0\n    total = 0\n    best = len(nums) + 1\n    for right, value in enumerate(nums):\n        total += value\n        while ____:\n            best = min(best, ____)\n            total -= nums[left]\n            left += 1\n    return 0 if best == len(nums) + 1 else best"},
		},
		Solution: Solution{
			Code:            "def minSubArrayLen(target, nums):\n    left = 0\n    total = 0\n    best = len(nums) + 1\n    for right, value in enumerate(nums):\n        total += value\n        while total >= target:\n            best = min(best, right - left + 1)\n            total -= nums[left]\n            left += 1\n    return 0 if best == len(nums) + 1 else best",
			Explanation:     "The first two drills had a window of a given length; here the length is the answer, so the left edge moves too. Grow right, then shrink left for as long as the window is still valid, recording the length at every valid moment. This is the 'shrink while valid' half of the variable-window template; Longest Substring Without Repeating Characters in the ladder is its 'shrink while invalid' twin.",
			BruteForce:      "For every start index, extend to the right until the sum reaches the target: O(n^2) in the worst case. It rescans the same elements for each start, although moving the start one step right only removes one element from the sum.",
			KeyInsight:      "With positive numbers, the sum only goes up as the window grows and only goes down as it shrinks, so each pointer can move only forward: right to make the window valid, left to see how small it can get while staying valid.",
			WhyItWorks:      "Invariant: when the while loop ends, nums[left..right] has sum below target (or is empty), and every shorter window ending at right has also been ruled out, since those start further right and have an even smaller sum. Every window nums[l..r] with sum >= target is considered: when right reaches r, left has not passed l, because left only moves past an index when the window starting there already reached the target, and a window starting there ending at r reaches it too, so some valid window at least that short, ending at r, has been recorded. The best length is therefore the minimum over all end positions. left and right each move at most n times, so the time is O(n).",
			CommonMistakes:  []string{"Shrinking with `if` instead of `while`. After one big element enters, several elements may need to leave, and a single `if` removes only one, so a shorter window is missed.", "Using `right - left` for the length, which is one too small: a window from left to right inclusive has `right - left + 1` elements.", "Recording the length after moving left (after the subtraction), when the window may already be too small. Record inside the loop, before removing nums[left].", "Returning the starting value of `best` (a huge number or infinity) when no window reaches the target. The statement asks for 0.", "Restarting the sum from scratch at every start index. That is O(n^2) and times out on long inputs."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Grow the window", Explanation: "Move right one step and add the new element to the running total.", CodeSnippet: "total += value", LineStart: 6, LineEnd: 6},
				{Title: "Shrink while valid", Explanation: "As long as the window reaches the target it is a candidate, so loop (not just once): record its length, then drop the leftmost element and see if the shorter window still works. On target 7 and [2,3,1,2,4,3], once 4 arrives the window [3,1,2,4] sums to 10, and it records 5, then 4, then 3 as 2, 3 and 1 leave, and stops at [2,4] = 6.", CodeSnippet: "while total >= target:\n    best = min(best, right - left + 1)\n    total -= nums[left]\n    left += 1", LineStart: 7, LineEnd: 10},
				{Title: "Handle no answer", Explanation: "If best never changed from its starting value, no window reached the target and the answer is 0.", CodeSnippet: "return 0 if best == len(nums) + 1 else best", LineStart: 11, LineEnd: 11},
			},
			Alternatives: []AltSolution{
				{Name: "Prefix sums and binary search", Code: "import bisect\n\n\ndef minSubArrayLen(target, nums):\n    prefix = [0]\n    for value in nums:\n        prefix.append(prefix[-1] + value)\n    best = len(nums) + 1\n    for start in range(len(nums)):\n        end = bisect.bisect_left(prefix, prefix[start] + target)\n        if end <= len(nums):\n            best = min(best, end - start)\n    return 0 if best == len(nums) + 1 else best", TimeComplexity: "O(n log n)", SpaceComplexity: "O(n)", Explanation: "Positive numbers make prefix sums increasing, so the first end whose prefix reaches prefix[start] + target can be found by binary search. It is a good second tool, but the window does the same job in O(n) time and O(1) space."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does shrinking from the left stay safe? Could a shorter valid window that starts at an element we already dropped be lost?", Answer: "Dropping nums[left] only happens when the window nums[left..right] reached the target, and its length has been recorded. A window starting at left and ending at a later right is longer than that recorded one, so it can never be the shortest. Positive values are what make this true: they guarantee that a longer window starting at the same left has no smaller sum."},
			{Prompt: "What if nums could contain negative numbers? Does the window still work?", Answer: "No. A negative element makes the sum go down when the window grows, so 'sum >= target' is no longer monotonic: shrinking could make a valid window invalid and growing could make it valid again, so the pointers cannot only move forward. That case needs prefix sums with a monotonic deque (Shortest Subarray with Sum at Least K)."},
		},
	},
	{
		ID:              "drill-longest-sum-at-most-k",
		Number:          0,
		Kind:            KindDrill,
		Order:           5,
		Title:           "Longest Subarray With Sum at Most K",
		Slug:            "drill-longest-sum-at-most-k",
		Difficulty:      "Easy",
		Category:        "sliding-window",
		Tags:            []string{"Sliding Window", "Array", "Two Pointers", "Foundations"},
		RelatedChapters: []int{4},
		Description:     "You are given a list `nums` of **non-negative** integers and a non-negative integer `k`. Return the length of the longest run of neighbouring elements whose sum is **at most** `k`. If even a single element is above `k`, no run is valid, and the answer is 0.\n\nThis is the other half of the variable window. In the previous drill you shrank while the window was *valid* (sum big enough) to find the shortest. Here the window wants to be as long as possible, so you let it grow and shrink only when it breaks the rule:\n\n1. **Grow**: move `right` one step and add `nums[right]` to `total`.\n2. **Shrink until valid**: while `total > k`, remove `nums[left]` from `total` and move `left` forward. If the new element alone is above `k`, the window empties completely.\n3. **Record**: now that the window is valid again, its length `right - left + 1` is a candidate for the answer (an empty window has length 0).\n\nSince all numbers are non-negative, taking an element out never raises the sum, so shrinking always helps. Use a `while` so that the window really is valid when you record it: a big element may need several others to leave.",
		Constraints:     []string{"0 <= len(nums) <= 100000", "0 <= nums[i] <= 10^4", "0 <= k <= 10^9"},
		Examples: []Example{
			{Input: "nums = [3,1,2,1,1,4], k = 5", Output: "4", Explanation: "[1,2,1,1] sums to 5 and has four elements. Adding the 3 in front (sum 8) or the 4 behind (sum 9) breaks the limit, and no run of five elements fits."},
			{Input: "nums = [5,7], k = 4", Output: "0", Explanation: "Both elements are above 4, so each time the window empties out and its length is 0."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{3, 1, 2, 1, 1, 4}, "k": 5}, Expected: 4, Note: "example 1: the longest run with sum <= 5 is [1,2,1,1]"},
			{Input: map[string]interface{}{"nums": []int{5, 7}, "k": 4}, Expected: 0, Note: "example 2: every element is already above k, so no run is valid and the answer is 0"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 2}, "k": 10}, Expected: 3, Hidden: true, Note: "the whole list fits"},
			{Input: map[string]interface{}{"nums": []int{0, 0, 0, 0}, "k": 0}, Expected: 4, Hidden: true, Note: "zeros with k = 0: every element can stay in the window"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 9, 2, 2}, "k": 4}, Expected: 3, Hidden: true, Note: "a big element forces several elements out at once"},
			{Input: map[string]interface{}{"nums": []int{9, 1, 1}, "k": 2}, Expected: 2, Hidden: true, Note: "the big element is first: the window is empty after it, then grows again"},
			{Input: map[string]interface{}{"nums": []int{4}, "k": 4}, Expected: 1, Hidden: true, Note: "one element equal to k"},
			{Input: map[string]interface{}{"nums": []int{8, 1, 9}, "k": 7}, Expected: 1, Hidden: true, Note: "a lone element above k: the window must become empty, so the answer is 1 from [1], not 2"},
			{Input: map[string]interface{}{"nums": []int{4}, "k": 3}, Expected: 0, Hidden: true, Note: "one element above k: the empty window, length 0"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5}, "k": 6}, Expected: 3, Hidden: true, Note: "the best run is not at the start: [1,2,3] vs [2,3] vs [3]"},
			{Input: map[string]interface{}{"nums": []int{5, 1, 1, 1, 1, 1}, "k": 5}, Expected: 5, Hidden: true, Note: "the answer ends at the end of the list: do not stop early"},
			{Input: map[string]interface{}{"nums": []int{6, 0, 0, 0, 6}, "k": 5}, Expected: 3, Hidden: true, Note: "zeros on the ends of a run belong in the window: the window can be longer than the number of positive values"},
			{Input: map[string]interface{}{"nums": swOnes(100000), "k": 50000}, Expected: 50000, Hidden: true, Timeout: 5000, Note: "100000 ones and k = 50000: starting a fresh scan from every index does billions of steps and runs out of time; the window is O(n)"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def longestSubarrayAtMost(nums, k):\n    # Return the length of the longest run of neighbouring elements with sum <= k\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "You want the window as long as possible while its sum stays at most k. When a new element pushes the sum over k, which end of the window should you give up, and for how long?"},
			{Level: 2, Type: "insight", Content: "Give up elements from the left for as long as the sum is above k: with non-negative numbers, each removal can only lower it. When that loop stops, the window is valid again, and its length is a candidate for the answer. Each pointer moves forward only, so it is O(n). This is the variable window that shrinks until valid; see Chapter 4.", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Grow: add nums[right] to the total.\n2. Shrink until valid: while total > k, subtract nums[left] and move left forward.\n3. Record: the window is valid now, so best = max(best, right - left + 1).\n4. Return best."},
			{Level: 4, Type: "skeleton", Content: "def longestSubarrayAtMost(nums, k):\n    left = 0\n    total = 0\n    best = 0\n    for right, value in enumerate(nums):\n        total += value\n        while ____:\n            total -= nums[left]\n            left += 1\n        best = max(best, ____)\n    return best"},
		},
		Solution: Solution{
			Code:            "def longestSubarrayAtMost(nums, k):\n    left = 0\n    total = 0\n    best = 0\n    for right, value in enumerate(nums):\n        total += value\n        while total > k:\n            total -= nums[left]\n            left += 1\n        best = max(best, right - left + 1)\n    return best",
			Explanation:     "This builds on Minimum Size Subarray Sum, with the roles reversed: there the window shrank while valid to find the shortest, here it shrinks until valid to find the longest, and the answer is recorded after the shrink. This 'grow, shrink until valid, record' shape is exactly Longest Substring Without Repeating Characters, Longest Substring with At Most K Distinct Characters and Longest Repeating Character Replacement in the ladder, with a different rule in place of 'sum <= k'.",
			BruteForce:      "For every start index, extend to the right while the sum stays at most k: O(n^2) in the worst case. Moving the start one step right changes the sum by only one element, but the scan from the new start rebuilds the sum from nothing.",
			KeyInsight:      "With non-negative numbers, a valid window stays valid when it shrinks, and an invalid one can only be fixed by shrinking it from the left. So left never needs to move back, and one pass with two pointers sees every candidate.",
			WhyItWorks:      "Invariant: after the shrink loop, nums[left..right] is the longest valid window that ends at right. A valid window ending at right can start at any index from left onwards, because shrinking only happens while the window is invalid, so every start before left was shown to give a sum above k, and since the numbers are not negative, a longer window ending at right has a sum at least as big. So no valid window ending at right is longer. best is the maximum over all right, which covers every possible end. left and right only move forward, so the time is O(n); if nums[right] > k alone, left passes right and the window is empty, with length 0.",
			CommonMistakes:  []string{"Guarding the shrink with `left < right` (or similar) so the window can never become empty. If a single element is above k, no run is valid, and the window must shrink to nothing so that it records length 0.", "Using `right - left` for the length instead of `right - left + 1`. A window from left to right inclusive has `right - left + 1` elements.", "Recording the length before the shrink loop. The window may be invalid at that moment, so it records a run whose sum is over k.", "Starting a fresh sum from every start index. It is correct but O(n^2) and times out on long inputs."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Grow the window", Explanation: "Move right one step and add the new element to the total, even if that breaks the limit; the next step repairs it.", CodeSnippet: "total += value", LineStart: 6, LineEnd: 6},
				{Title: "Shrink until valid", Explanation: "While the total is over k, give up the leftmost element. It is a loop because a big new element can force many out; if the element alone is over k the window becomes empty. On [1,1,1,9,2,2] with k = 4, the 9 pushes the total to 12 and four removals empty the window.", CodeSnippet: "while total > k:\n    total -= nums[left]\n    left += 1", LineStart: 7, LineEnd: 9},
				{Title: "Record the valid window", Explanation: "The window now has sum at most k, so its length is a real candidate.", CodeSnippet: "best = max(best, right - left + 1)", LineStart: 10, LineEnd: 10},
				{Title: "Return the longest", Explanation: "Every end position has been tried, so best is the longest valid run (0 if none).", CodeSnippet: "return best", LineStart: 11, LineEnd: 11},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the length recorded after the while loop here, but inside it in the Minimum Size Subarray Sum drill?", Answer: "Because the rule is different. When you want the shortest window that is valid, every moment the window is valid is a candidate, so you record while shrinking. When you want the longest valid window, the window is only a candidate once the shrink has made it valid, so you record after the loop. The shape follows the question: 'shortest, shrink while valid' versus 'longest, shrink until valid'."},
			{Prompt: "What breaks if nums can contain negative numbers?", Answer: "Shrinking from the left no longer has to lower the sum, and a longer window could have a smaller sum than a shorter one. So the validity of the window is not monotonic, left cannot be moved on only by the sum, and the window method gives wrong answers. Prefix sums with a different structure are needed."},
		},
	},
	{
		ID:              "drill-subarray-product-less-than-k",
		Number:          713,
		Kind:            KindDrill,
		Order:           6,
		Title:           "Subarray Product Less Than K",
		Slug:            "drill-subarray-product-less-than-k",
		Difficulty:      "Easy",
		Category:        "sliding-window",
		Tags:            []string{"Sliding Window", "Array", "Two Pointers", "Foundations"},
		RelatedChapters: []int{4},
		Description:     "You are given a list `nums` of **positive** integers and an integer `k`. Return how many runs of neighbouring elements have a product **strictly less than** `k`. Runs at different positions count separately, even if they contain the same numbers, and a single element is a run of length 1. Return 0 when `k` is 0 or 1, since the product of positive integers is at least 1.\n\nYou will count, so this drill adds one idea to the variable window: **how many windows end at `right`?** Keep the window `nums[left..right]` valid (product below `k`), as in the previous drill. Every run that ends at `right` and starts anywhere from `left` to `right` is also valid, because its product is no bigger (all numbers are at least 1). That gives exactly `right - left + 1` new runs for each `right`, which you add to the count.\n\n1. **Grow**: multiply `product` by `nums[right]`.\n2. **Shrink until valid**: while `product >= k`, divide `product` by `nums[left]` and move `left` forward.\n3. **Count**: add `right - left + 1`.\n\nGuard `k <= 1` first: with `k = 1` the product is never below it, so the shrink loop would push `left` past the end of the list.",
		Constraints:     []string{"1 <= len(nums) <= 100000", "1 <= nums[i] <= 1000", "0 <= k <= 10^6"},
		Examples: []Example{
			{Input: "nums = [10,5,2,6], k = 100", Output: "8", Explanation: "The valid runs are [10], [5], [2], [6], [10,5], [5,2], [2,6] and [5,2,6] (product 60). [10,5,2] is 100, which is not below 100."},
			{Input: "nums = [1,2,3], k = 0", Output: "0", Explanation: "No product of positive numbers is below 0."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{10, 5, 2, 6}, "k": 100}, Expected: 8, Note: "example 1: eight runs have a product below 100"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}, "k": 0}, Expected: 0, Note: "example 2: k = 0, and no product of positive numbers is below 0"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1}, "k": 1}, Expected: 0, Hidden: true, Note: "k = 1: products are at least 1, so none are below 1 (a window must not run past the end)"},
			{Input: map[string]interface{}{"nums": []int{5}, "k": 6}, Expected: 1, Hidden: true, Note: "one element below k"},
			{Input: map[string]interface{}{"nums": []int{5}, "k": 5}, Expected: 0, Hidden: true, Note: "one element equal to k: the product must be strictly less"},
			{Input: map[string]interface{}{"nums": []int{1, 1, 1, 1}, "k": 2}, Expected: 10, Hidden: true, Note: "all ones: every run counts, 4 + 3 + 2 + 1 = 10"},
			{Input: map[string]interface{}{"nums": []int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2}, "k": 19}, Expected: 14, Hidden: true, Note: "large elements empty the window: one run per small element, and a pair when it fits"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4}, "k": 25}, Expected: 10, Hidden: true, Note: "the product 24 is below 25, so the whole list and every run inside it count"},
			{Input: map[string]interface{}{"nums": []int{100, 1, 1, 100}, "k": 50}, Expected: 3, Hidden: true, Note: "a big element forces several others out and a big element alone leaves an empty window"},
			{Input: map[string]interface{}{"nums": []int{2, 2, 2, 2, 2, 2}, "k": 9}, Expected: 15, Hidden: true, Note: "a window that slides: three 2s have a product of 8, four have 16"},
			{Input: map[string]interface{}{"nums": swSparseTwos(100000), "k": 1024}, Expected: swProductExpected(swSparseTwos(100000), 1024), Hidden: true, Timeout: 5000, Note: "100000 numbers, mostly ones with a few twos: about a billion runs qualify, so counting them one at a time runs out of time; the window counts them in O(n)"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def numSubarrayProductLessThanK(nums, k):\n    # Return how many runs of neighbouring elements have a product below k\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose nums[left..right] has a product below k. Which runs that end at right are guaranteed to be below k too, and how many are there?"},
			{Level: 2, Type: "insight", Content: "Every number is at least 1, so cutting elements off the left of a valid window never raises its product. That means every start from left to right gives a valid run ending at right, which is right - left + 1 runs, counted all at once instead of one at a time. Keep the window valid with the same shrink loop as before, and add that count for each right. Return 0 first when k <= 1. See Chapter 4.", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Guard: if k <= 1, return 0.\n2. Grow: multiply product by nums[right].\n3. Shrink until valid: while product >= k, divide by nums[left] and move left forward.\n4. Count: add right - left + 1, the number of valid runs ending at right.\n5. Return the count."},
			{Level: 4, Type: "skeleton", Content: "def numSubarrayProductLessThanK(nums, k):\n    if k <= 1:\n        return 0\n    left = 0\n    product = 1\n    count = 0\n    for right, value in enumerate(nums):\n        product *= value\n        while ____:\n            product //= nums[left]\n            left += 1\n        count += ____\n    return count"},
		},
		Solution: Solution{
			Code:            "def numSubarrayProductLessThanK(nums, k):\n    if k <= 1:\n        return 0\n    left = 0\n    product = 1\n    count = 0\n    for right, value in enumerate(nums):\n        product *= value\n        while product >= k:\n            product //= nums[left]\n            left += 1\n        count += right - left + 1\n    return count",
			Explanation:     "This extends Longest Subarray With Sum at Most K: the same grow, shrink-until-valid loop, but with a product (undone by division, as a sum is undone by subtraction), and at each step it adds the number of valid windows that end at right instead of tracking the longest. The count `right - left + 1` is the idea to keep: Subarrays with K Different Integers and many 'count the subarrays' problems rely on it.",
			BruteForce:      "Start a run at every index and extend it while the product stays below k, counting each run: O(n * L), where L is the length of the longest valid run, up to O(n^2). Each run is rebuilt and counted one at a time, though all the runs ending at the same place are valid together.",
			KeyInsight:      "If a window is valid, every shorter run that ends at the same right edge is valid too, so each right edge adds right - left + 1 runs in one step.",
			WhyItWorks:      "Invariant: after the shrink loop, nums[left..right] is the longest run ending at right whose product is below k (empty if nums[right] >= k), and product is its product. Every run nums[i..right] with left <= i <= right has a product that is at most that of nums[left..right], because all values are at least 1, so it is valid; and every run starting before left has a product of at least k, since left only moved past it when the product was too big. So exactly right - left + 1 valid runs end at right. Summing over every right counts each valid run once, by its right end. left and right only move forward, so the time is O(n).",
			CommonMistakes:  []string{"Skipping the `k <= 1` guard. With k = 0 or 1 the product is never below k, so the shrink loop keeps moving `left` past the end of the list and crashes (or runs forever).", "Adding 1 for each right, or counting only the longest window, instead of `right - left + 1`. Every shorter run that ends at right counts too.", "Using `right - left` as the count, which is one too small: that leaves out the single-element run nums[right].", "Shrinking with `if` instead of `while`. A big new element can need several others to leave, and a single `if` removes only one, so the window's product stays at or above k.", "Using `<=` in the shrink condition and so counting products equal to k as valid, when the statement says strictly less than k."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Guard the small k", Explanation: "Products of positive integers are at least 1, so no run is below k when k is 0 or 1. Returning here also keeps the shrink loop from running off the end.", CodeSnippet: "if k <= 1:\n    return 0", LineStart: 2, LineEnd: 3},
				{Title: "Grow the window", Explanation: "Multiply the new element into the product.", CodeSnippet: "product *= value", LineStart: 8, LineEnd: 8},
				{Title: "Shrink until valid", Explanation: "While the product has reached k, divide out the leftmost element and move left forward. Dividing is exact because the product contains that element as a factor. On [10,5,2,6] with k = 100, adding 6 makes 10 * 5 * 2 * 6 = 600, so 10 and 5 leave.", CodeSnippet: "while product >= k:\n    product //= nums[left]\n    left += 1", LineStart: 9, LineEnd: 11},
				{Title: "Count the runs ending at right", Explanation: "The window is valid, so every start from left to right gives a valid run ending at right. With the window [2,6] there are 2: [6] and [2,6].", CodeSnippet: "count += right - left + 1", LineStart: 12, LineEnd: 12},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it right to add right - left + 1 for each right, rather than checking each shorter run?", Answer: "All values are at least 1, so a run that is a part of a valid window, cut from the left, has a product no bigger than the window's. It is therefore valid too, and the number of such runs is the number of possible starts, right - left + 1. This lets one step count many windows. Runs that start before left were already shown to be too big."},
			{Prompt: "What if the numbers could be 0 or negative? What fails?", Answer: "The product is no longer monotonic: with a 0 the product drops to 0 and everything becomes valid again, and with negatives a longer window can have a smaller (more negative) product. Also, you cannot divide out a 0 to undo it. The 'valid windows are closed under cutting' argument breaks, so this method needs the positivity."},
		},
	},
}

// swBigNums is a fixed pseudo-random list (values -10000..10000) for the
// time-limit cases of the window-sum drills.
func swBigNums(n int) []int {
	rng := rand.New(rand.NewSource(11))
	nums := make([]int, n)
	for i := range nums {
		nums[i] = rng.Intn(20001) - 10000
	}
	return nums
}

func swWindowSumsExpected(nums []int, k int) []int {
	window := 0
	for _, v := range nums[:k] {
		window += v
	}
	sums := []int{window}
	for right := k; right < len(nums); right++ {
		window += nums[right] - nums[right-k]
		sums = append(sums, window)
	}
	return sums
}

func swMaxAverageExpected(nums []int, k int) float64 {
	best := swWindowSumsExpected(nums, k)[0]
	for _, s := range swWindowSumsExpected(nums, k) {
		if s > best {
			best = s
		}
	}
	return float64(best) / float64(k)
}

// swBigDistinctNums is a fixed pseudo-random list of mostly different values.
func swBigDistinctNums(n int) []int {
	rng := rand.New(rand.NewSource(13))
	nums := make([]int, n)
	for i := range nums {
		nums[i] = rng.Intn(1000000001)
	}
	return nums
}

func swDistinctExpected(nums []int, k int) []int {
	counts := map[int]int{}
	var result []int
	for right, v := range nums {
		counts[v]++
		if right >= k {
			leaving := nums[right-k]
			counts[leaving]--
			if counts[leaving] == 0 {
				delete(counts, leaving)
			}
		}
		if right >= k-1 {
			result = append(result, len(counts))
		}
	}
	return result
}

func swOnes(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = 1
	}
	return nums
}

// swSparseTwos is mostly ones with a 2 every 997th place.
func swSparseTwos(n int) []int {
	nums := swOnes(n)
	for i := 0; i < n; i += 997 {
		nums[i] = 2
	}
	return nums
}

func swProductExpected(nums []int, k int) int {
	left, product, count := 0, 1, 0
	for right, v := range nums {
		product *= v
		for product >= k {
			product /= nums[left]
			left++
		}
		count += right - left + 1
	}
	return count
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, SlidingWindowDrills...)
}
