package chapter

import "time"

// chapter04 teaches the sliding-window pattern: fixed windows that slide by
// adding one element and dropping one, and variable windows that grow on the
// right and shrink on the left while the window is invalid (or, for
// "shortest" problems, while it is still valid).
// Storyboards (IDs registered in app.js): sliding-window.
func chapter04() Chapter {
	return Chapter{
		Number:        4,
		Slug:          "sliding-window",
		Title:         "Sliding Window",
		Category:      "sliding-window",
		Prerequisites: []int{3},
		Summary:       "Answer questions about contiguous subarrays and substrings in one pass: keep a window [left, right], update a small summary of it as one element enters and one leaves, and move each edge only forward.",
		Motivation: `A fitness tracker logs your steps every minute for a week: n = 10,080 numbers. You want the busiest hour, the 60 consecutive minutes with the most steps. The obvious way sums each of the 10,021 windows from scratch: 10,021 x 60 = about 600,000 additions. But the window starting at minute 1 and the one starting at minute 2 share 59 minutes. Only one minute enters and one leaves, so each new sum is the old sum plus one number minus one number: 10,080 steps in total.

The gain is much bigger when the window size isn't fixed. "Find the longest stretch of a 50,000-character string with no repeated character" has n(n+1)/2 = 1.25 billion substrings to check, and checking each one for repeats costs more on top. A sliding window solves it with each index entering the window once and leaving once: about 100,000 steps. This chapter shows when that shortcut is allowed and how to write it without off-by-one bugs.`,
		Objectives: []string{
			"Recognise from a problem statement whether it needs a fixed-size or a variable-size window, or no window at all.",
			"Implement a fixed window that slides by adding the entering element and removing the leaving one.",
			"Implement a variable window with the subgoals Grow the window, Shrink until valid (or Shrink while valid), Record the answer.",
			"Explain why a nested while loop inside the for loop is still O(n), using the fact that left only moves forward.",
			"Choose the window state (a sum, a set, a count map, a formed/required counter, a monotonic deque) that makes the validity check O(1).",
		},
		WhenToUse: []string{
			"The question is about a contiguous subarray or substring (\"consecutive days\", \"substring\", \"subarray\"), not an arbitrary subset.",
			"It asks for the longest, shortest, maximum, minimum or number of windows that satisfy a condition.",
			"A window size k is given (\"every k consecutive elements\", \"a substring of length len(s1)\"): fixed window.",
			"The condition is monotonic: making a valid window smaller keeps it valid (for longest problems), or making it bigger keeps it valid (for shortest problems). Sums of non-negative numbers, counts of distinct characters and \"at most k\" rules behave this way.",
			"Warning sign: with negative numbers a sum condition is not monotonic, so a window can't decide when to shrink. Use prefix sums with a hash map (Chapter 2) instead.",
		},
		Concepts: []Concept{
			{
				Name:        "Fixed-size window",
				Description: "When every candidate has the same length k, slide one window across the input. Each step one element enters on the right and one leaves on the left, and the window's summary (sum, counts) is updated from the previous one instead of being recomputed.",
				WhyItWorks:  "Two neighbouring windows nums[i-k..i-1] and nums[i-k+1..i] share k - 1 elements. Any summary that can 'add one element' and 'remove one element' in O(1) (a sum, a count per letter) therefore moves from one window to the next in O(1). The first window costs k steps and each of the remaining n - k windows costs O(1), so the total is O(n) instead of O(n*k). Summaries that can't remove an element cheaply, like a maximum, need a smarter structure (the monotonic deque in Sliding Window Maximum).",
				Intuition:   "A picture frame sliding along a bookshelf. To know how many pages are in the frame after a move, you don't recount every book: one book entered on the right (add its pages) and one left on the left (subtract its pages). The frame is the window, the books are the elements, the running page total is the window's summary.",
				CoreIdeas: []string{
					"Build the first window of k elements, then for right = k .. n-1 add nums[right] and remove nums[right - k].",
					"The element leaving is always exactly k positions behind the one entering: index right - k.",
					"Record the answer after each slide, including for the very first window.",
					"The window's state can be a sum, a 26-slot count list (Permutation in String) or a deque of candidates (Sliding Window Maximum).",
				},
				CommonMistakes: []string{
					"Removing nums[right - k + 1] instead of nums[right - k], which leaves a window of k - 1 elements. Check with k = 1: the element leaving must be the previous one.",
					"Forgetting to compare the first window before the first slide, or the last window after the final slide. Record after every step, and test inputs where the answer is the first or last window.",
					"Using a fixed window when the length of the answer is what you are looking for. If the question asks for the longest or shortest, you need a variable window.",
				},
				Examples: []CodeExample{
					{
						ID:       "fixed-window-max-sum",
						Title:    "Largest sum of k consecutive elements",
						Language: "python",
						Snippet: `def max_sum_subarray(nums, k):
    window_sum = sum(nums[:k])        # the first window
    best = window_sum
    for right in range(k, len(nums)):
        window_sum += nums[right] - nums[right - k]   # slide: add new, drop old
        best = max(best, window_sum)
    return best

print(max_sum_subarray([2, 1, 5, 1, 3, 2], 3))  # 9, from [5, 1, 3]`,
						Notes: "Windows: [2,1,5]=8, [1,5,1]=7, [5,1,3]=9, [1,3,2]=6. Four windows, each found with one addition and one subtraction.",
					},
				},
			},
			{
				Name:        "Variable-size window: grow right, shrink left",
				Description: "When the answer's length is unknown, keep a window [left, right]. Move right forward one step at a time (Grow the window). After each step, move left forward until the window satisfies the rule again (Shrink until valid), then Record the answer. For 'shortest' problems the inner loop runs while the window is valid instead (Shrink while valid) and records inside the loop.",
				WhyItWorks:  "The trick depends on monotonic validity. For a 'longest' problem, if s[left..right] is invalid (say it contains a repeated letter), then every window that contains it is invalid too, so no answer can start at left and end at right or later. That means left never has to move backwards. For each right, the loop finds the smallest left that makes the window valid, so right - left + 1 is the longest valid window ending at right, and the best of those is the answer. Both pointers only move forward and each moves at most n times, so the nested loops do at most 2n steps in total: O(n), not O(n^2). For 'shortest' problems the argument is mirrored: if a window is valid, every larger window is valid, so for each right you shrink while it stays valid and record each window on the way.",
				Intuition:   "A caterpillar crawling along a branch. Its head (right) always moves forward to explore. When its body covers something it can't hold (a repeated letter, too many changes), it pulls its tail (left) forward until it is comfortable again. Neither end ever moves backwards, so the caterpillar crosses the branch once. The head is the right pointer, the tail is the left pointer, and 'comfortable' is the validity rule.",
				CoreIdeas: []string{
					"Template for longest: for right in range(n): add s[right]; while window invalid: remove s[left], left += 1; record right - left + 1.",
					"Template for shortest: for right in range(n): add s[right]; while window valid: record right - left + 1, remove s[left], left += 1.",
					"left only moves forward, so the total work of the inner while loop over the whole run is at most n steps (amortised O(1) per step).",
					"The window is defined by its state. Choose state that makes the validity check O(1): a set for 'no repeats', a count map for 'at most k distinct', length - max_freq for 'at most k changes'.",
					"Monotonicity is the requirement. With negative numbers, 'sum <= target' can become valid again by adding an element, so the shrink decision is wrong. Use prefix sums instead.",
				},
				CommonMistakes: []string{
					"Recording the answer before the shrink loop in a longest problem. The window is only valid after shrinking, so record after the while loop.",
					"Writing if instead of while for the shrink step. One removal may not be enough (\"abcb\" adding the second b must drop a and b). Use while unless you can prove one step always suffices.",
					"Resetting the window on a violation (left = right) instead of shrinking it one step at a time. That skips valid windows: for \"dvdf\" the answer \"vdf\" starts inside the old window.",
					"Applying the pattern to sums with negative numbers. Check that validity is monotonic before using a window.",
				},
				Examples: []CodeExample{
					{
						ID:       "variable-window-two-distinct",
						Title:    "Longest substring with at most two distinct characters",
						Language: "python",
						Snippet: `def longest_two_distinct(s):
    count = {}                     # Set up the window state
    left = 0
    best = 0
    for right in range(len(s)):    # Grow the window
        count[s[right]] = count.get(s[right], 0) + 1
        while len(count) > 2:      # Shrink until valid
            count[s[left]] -= 1
            if count[s[left]] == 0:
                del count[s[left]]
            left += 1
        best = max(best, right - left + 1)   # Record the answer
    return best

print(longest_two_distinct("eceba"))  # 3, from "ece"`,
						Notes: "The four subgoals are the same in every variable-window problem; only the window state and the validity rule change.",
					},
				},
			},
			{
				Name:        "Window state that answers 'is it valid?' in O(1)",
				Description: "A window is only fast if you can update its summary when one element enters or leaves and then check the rule without rescanning the window. Pick the summary to fit the rule.",
				WhyItWorks:  "The sliding window saves time only if each step costs O(1) (or O(alphabet size)). If checking validity meant rescanning the window, each step would cost O(window length) and you'd be back at O(n^2). So for each rule, find a summary that changes by a small known amount when one element enters or leaves: a running sum for sum rules, a set for 'no repeats', a count map plus a max count for 'at most k changes', and a formed counter that tracks how many requirements are met for 'contains all of t'. formed changes only when a count crosses its required value, so it is correct after every single add or remove.",
				Intuition:   "A bouncer at a club door with a clicker counter. He doesn't count everyone inside each time someone enters or leaves; he clicks up or down. Checking 'is the club full?' is one look at the clicker. The window is the club, people entering and leaving are elements, and the clicker is the window's state.",
				CoreIdeas: []string{
					"Every add has a matching remove: if growing does count[c] += 1, shrinking does count[c] -= 1 (and deletes the key at 0 if the rule uses len(count)).",
					"formed/required: required = number of distinct characters that need enough copies; formed = how many currently have enough. Increase formed only when a count reaches its need exactly, decrease it only when a count drops below its need.",
					"Some states can't remove cheaply. A window maximum needs a monotonic deque: drop smaller elements from the back (they can never be the maximum again) and expire the front by index. Chapter 5 develops this 'monotonic' idea for stacks.",
					"For a maximum-length answer, a slightly stale summary can be fine (Longest Repeating Character Replacement), but only with a proof; when in doubt, keep the summary exact.",
				},
				CommonMistakes: []string{
					"Updating the state on grow but forgetting the matching update on shrink, so the state drifts away from the real window. Write the add and remove lines as a pair.",
					"Counting formed += 1 every time a needed character arrives, instead of only when its count reaches the need. Extra copies would make the window look complete too early.",
					"Leaving zero-count keys in a dictionary when the rule uses len(count) (number of distinct characters). Delete keys that drop to 0.",
				},
				Examples: []CodeExample{
					{
						ID:       "window-state-formed",
						Title:    "Shortest window containing all of t (formed/required)",
						Language: "python",
						Snippet: `def min_window(s, t):
    need = {}
    for ch in t:
        need[ch] = need.get(ch, 0) + 1
    required, formed = len(need), 0
    window = {}
    left, best = 0, ""
    for right, ch in enumerate(s):                 # Grow the window
        window[ch] = window.get(ch, 0) + 1
        if ch in need and window[ch] == need[ch]:
            formed += 1
        while formed == required:                  # Shrink while valid
            if best == "" or right - left + 1 < len(best):
                best = s[left:right + 1]           # Record the answer
            out = s[left]
            window[out] -= 1
            if out in need and window[out] < need[out]:
                formed -= 1
            left += 1
    return best`,
						Notes: "This is the practice problem Minimum Window Substring in compact form. The full version stores indices instead of slicing.",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Fixed window: largest sum of k consecutive elements",
				Problem: "Given nums = [2, 1, 5, 1, 3, 2] and k = 3, return the largest sum of any 3 consecutive elements.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Set up the window state",
						Explanation: "The window has a fixed size, so its state is just its sum. Build the first window [0, k-1] directly and treat it as the best so far.",
						Code:        "window_sum = sum(nums[:k])\nbest = window_sum",
						State:       "window = [2, 1, 5], window_sum = 8, best = 8",
					},
					{
						Subgoal:     "Slide the window",
						Explanation: "Move right one step. The entering element is nums[right]; the leaving one is k positions behind it, nums[right - k]. Add one, subtract the other: O(1) instead of re-adding k numbers.",
						Code:        "for right in range(k, len(nums)):\n    window_sum += nums[right] - nums[right - k]",
						State:       "right = 3: add nums[3]=1, drop nums[0]=2 -> window [1, 5, 1], sum 7",
					},
					{
						Subgoal:     "Record the answer",
						Explanation: "After every slide the window is a complete candidate, so compare its sum with the best.",
						Code:        "    best = max(best, window_sum)",
						State:       "right = 3: best = max(8, 7) = 8; right = 4: add 3, drop 1 -> sum 9, best = 9; right = 5: add 2, drop 5 -> sum 6, best stays 9",
					},
					{
						Subgoal:     "Return the result",
						Explanation: "Every window of length k was considered exactly once: the first one in the setup and each later one after a slide. The answer is the largest sum seen.",
						Code:        "return best",
						State:       "returns 9 (window [5, 1, 3])",
					},
				},
				Code: `def max_sum_subarray(nums, k):
    window_sum = sum(nums[:k])
    best = window_sum
    for right in range(k, len(nums)):
        window_sum += nums[right] - nums[right - k]
        best = max(best, window_sum)
    return best`,
			},
			{
				Title:   "Variable window: longest substring with at most two distinct characters",
				Problem: "Given s = \"eceba\", return the length of the longest substring that contains at most 2 distinct characters. (The answer is \"ece\", length 3.)",
				Steps: []WorkedStep{
					{
						Subgoal:     "Set up the window state",
						Explanation: "The rule is about how many different characters the window holds, so keep a count per character. len(count) is then the number of distinct characters, as long as we delete keys that reach 0.",
						Code:        "count = {}\nleft = 0\nbest = 0",
						State:       "count = {}, left = 0, best = 0",
					},
					{
						Subgoal:     "Grow the window",
						Explanation: "Move right forward one character and add it to the counts. The window may now break the rule.",
						Code:        "for right in range(len(s)):\n    count[s[right]] = count.get(s[right], 0) + 1",
						State:       "right = 0..2: window \"ece\", count = {e: 2, c: 1}; right = 3 adds b: window \"eceb\", count = {e: 2, c: 1, b: 1} (3 distinct, invalid)",
					},
					{
						Subgoal:     "Shrink until valid",
						Explanation: "While there are more than 2 distinct characters, remove s[left] and move left forward. One removal may not be enough: removing the first e still leaves e in the window, so we must also remove c. Delete a key when its count reaches 0 so len(count) stays accurate.",
						Code:        "    while len(count) > 2:\n        count[s[left]] -= 1\n        if count[s[left]] == 0:\n            del count[s[left]]\n        left += 1",
						State:       "right = 3: drop e (left 0 -> 1, count e: 1), still 3 distinct; drop c (left 1 -> 2, c deleted). Window \"eb\", count = {e: 1, b: 1}",
					},
					{
						Subgoal:     "Record the answer",
						Explanation: "After the shrink loop the window is valid and it is the longest valid window that ends at right. Compare its length with the best so far.",
						Code:        "    best = max(best, right - left + 1)",
						State:       "right = 2: best = 3 (\"ece\"); right = 3: length 2, best stays 3; right = 4 adds a, drops e (left = 3), window \"ba\", best stays 3",
					},
					{
						Subgoal:     "Check the cost",
						Explanation: "The inner loop looks expensive, but left only moves forward and never passes right. Each character is added once and removed at most once, so the whole function does O(n) work.",
						Code:        "return best",
						State:       "returns 3; 5 additions and 3 removals in total for n = 5",
					},
				},
				Code: `def longest_two_distinct(s):
    count = {}
    left = 0
    best = 0
    for right in range(len(s)):
        count[s[right]] = count.get(s[right], 0) + 1
        while len(count) > 2:
            count[s[left]] -= 1
            if count[s[left]] == 0:
                del count[s[left]]
            left += 1
        best = max(best, right - left + 1)
    return best`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Shortest subarray with sum at least target",
				Prompt: "Given positive integers nums and a target, return the length of the shortest contiguous subarray whose sum is >= target, or 0 if there is none. Example: target = 7, nums = [2,3,1,2,4,3] returns 2 ([4,3]). This is a 'shortest' problem, so the inner loop runs while the window is valid. Fill in the blanks.",
				Code: `def minSubArrayLen(target, nums):
    left = 0
    window_sum = 0
    best = len(nums) + 1
    for right in range(len(nums)):
        ____                                # Grow the window
        while ____:                         # Shrink while valid
            best = ____                     # Record the answer
            window_sum -= nums[left]
            ____
    return 0 if best == len(nums) + 1 else best`,
				Blanks:      []string{"window_sum += nums[right]", "window_sum >= target", "min(best, right - left + 1)", "left += 1"},
				Explanation: "Growing adds nums[right] to the sum. Because all numbers are positive, a window with sum >= target stays valid if it grows, so for the shortest answer we shrink while it is still valid, recording the length before each removal. The removal is subtracting nums[left] and moving left forward. Trace: at right = 4 the window [3,1,2,4] (sum 10) records 4, shrinks to [1,2,4] (7) and records 3, then [2,4] (6) is invalid. At right = 5, [2,4,3] (9) records 3, [4,3] (7) records 2, and [3] stops the loop.",
			},
			{
				Title:  "Count windows of size k with a high average",
				Prompt: "Given arr, k and threshold, return how many windows of exactly k consecutive elements have an average >= threshold (equivalently, a sum >= threshold * k). Example: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4 returns 3. Fill in the blanks for this fixed-size window.",
				Code: `def numOfSubarrays(arr, k, threshold):
    count = 0
    window_sum = 0
    for right in range(len(arr)):
        window_sum += arr[right]            # Slide the window: add new
        if right >= k:
            window_sum -= ____              # Slide the window: drop old
        if ____ and window_sum >= threshold * k:
            count += 1                      # Record the answer
    return count`,
				Blanks:      []string{"arr[right - k]", "right >= k - 1"},
				Explanation: "Once the window would hold k + 1 elements (right >= k), the element leaving is arr[right - k]. A window only counts once it is complete, which first happens at right = k - 1. The qualifying windows are [2,5,5] = 12, [5,5,5] = 15 and [5,5,8] = 18, all >= 12.",
			},
		},
		Animations: []Storyboard{
			{
				ID:    "sliding-window",
				Title: "Sliding a Fixed Window",
				Goal:  "Find the largest sum of k = 3 consecutive elements in [2, 1, 5, 1, 3, 2] by sliding one window instead of re-adding every window.",
				MemoryTips: []string{
					"Set up the window state: build the first window once (sum of the first k elements).",
					"Slide the window: add the element entering on the right, subtract the one k positions behind it.",
					"Record the answer after every slide, and remember the first window counts too.",
					"For a variable window, swap 'slide' for: Grow the window, then Shrink until valid (longest) or Shrink while valid (shortest).",
					"Both edges only move forward, so the whole pass is O(n).",
				},
				Steps: []StoryboardStep{
					{Cue: "init", Narration: "Set up the window state: the first window is [2, 1, 5]. Add it up once: window_sum = 8, and the best so far is 8.", VisualHint: "First k elements highlighted with their sum", Duration: Duration(2 * time.Second), CodeRef: "sliding: sum(nums[:k])"},
					{Cue: "slide", Narration: "Slide the window one step: 1 enters on the right and 2 leaves on the left, so the sum becomes 8 + 1 - 2 = 7. The two middle elements are never re-added. Next, 3 enters and 1 leaves: 7 + 3 - 1 = 9. Then 2 enters and 5 leaves: 9 + 2 - 5 = 6.", VisualHint: "Window moves right; entering element in green, leaving in red", Duration: Duration(3 * time.Second), CodeRef: "sliding: += nums[i] - nums[i-k]"},
					{Cue: "track", Narration: "Record the answer after every slide: best goes 8, 8, 9, 9. The window [5, 1, 3] holds the record.", VisualHint: "Update max", Duration: Duration(2 * time.Second), CodeRef: "sliding: max_sum = max(...)"},
					{Cue: "complete", Narration: "Return 9. Four windows were checked with one addition and one subtraction each, instead of three additions each. With n = 100,000 and k = 1,000 that is 100,000 steps instead of 100 million.", VisualHint: "Final answer", Duration: Duration(2 * time.Second), CodeRef: "sliding: return max_sum"},
				},
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Which approach fits best? \"Given an array of positive integers and a number S, return the length of the longest contiguous subarray whose sum is at most S.\"",
				Choices: []string{
					"Two pointers from both ends of the sorted array (Chapter 3)",
					"A variable-size sliding window: grow right, shrink left while the sum is above S",
					"A fixed-size window of length S",
					"Sort the array and take the smallest elements until the sum exceeds S",
				},
				Answer:      1,
				Explanation: "\"Contiguous subarray\" plus \"longest\" plus a monotonic rule (with positive numbers, shrinking a window can only lower its sum) points to a variable window. Opposite-end two pointers need sorted input and look at pairs, not ranges. The window length is the unknown, so it can't be fixed. Sorting destroys contiguity, so it answers a different question.",
			},
			{
				Prompt: "Which approach fits best? \"Given an array that may contain negative numbers, count the contiguous subarrays whose sum equals k.\"",
				Choices: []string{
					"A variable-size sliding window that shrinks while the sum is above k",
					"Prefix sums with a hash map counting how often each prefix sum has appeared (Chapter 2)",
					"A fixed-size window of length k",
					"Binary search on the subarray length",
				},
				Answer:      1,
				Explanation: "With negative numbers, the sum is not monotonic: adding an element can lower it and removing one can raise it, so a window can't tell when to shrink. A subarray i..j sums to k exactly when prefix[j+1] - prefix[i] = k, so counting earlier prefix sums in a hash map (Chapter 2) solves it in O(n). This is the classic case where the window pattern looks right but isn't.",
			},
			{
				Prompt: "Trace this code on s = \"abba\". What are left and best when it returns?\n\nin_window = set(); left = 0; best = 0\nfor right in range(len(s)):\n    while s[right] in in_window:\n        in_window.remove(s[left])\n        left += 1\n    in_window.add(s[right])\n    best = max(best, right - left + 1)",
				Choices: []string{
					"left = 2, best = 2",
					"left = 1, best = 3",
					"left = 3, best = 2",
					"left = 0, best = 4",
				},
				Answer:      0,
				Explanation: "right = 0, 1: window \"ab\", best = 2. right = 2 (b): b is inside, so remove a (left = 1), then b (left = 2); add b, window \"b\". right = 3 (a): a is not in {b}, so add it; window \"ba\", left = 2, best stays 2. The 'jump to last index' version without max() would wrongly set left = 1 here and report 3.",
			},
			{
				Prompt:      "From memory, write the four subgoals of a variable-size window for a 'longest valid window' problem, and say how the 'shortest valid window' version differs.",
				Explanation: "Longest: (1) Set up the window state (counts/set/sum, left = 0, best). (2) Grow the window: for each right, add s[right] to the state. (3) Shrink until valid: while the window breaks the rule, remove s[left] and move left forward. (4) Record the answer: best = max(best, right - left + 1), after the loop. Shortest: the inner loop runs while the window is valid, and the answer is recorded inside the loop before each removal, because only valid windows are candidates and we want the smallest.",
			},
			{
				Prompt:      "Why is a sliding window with a while loop inside a for loop O(n) and not O(n^2)?",
				Choices:     []string{"Because the while loop runs at most once per step", "Because left only moves forward, so across the whole run it moves at most n times in total", "Because Python's set operations are O(1)", "It isn't; it is O(n^2) in the worst case"},
				Answer:      1,
				Explanation: "The inner loop can run many times in one step (in \"abcdec\" the second c removes a, b and c in one step), but every iteration moves left forward by one and left never passes right. So the total number of inner iterations over the whole run is at most n. Add the n outer steps and the total is O(n). O(1) set operations are necessary but not the reason.",
			},
			{
				Prompt:      "Prerequisite (Chapter 3): in Two Sum II on a sorted array, the pointers start at both ends. When the pair sum is too small, why is it safe to move the left pointer right, and how does that compare with the sliding window's 'left only moves forward'?",
				Explanation: "Because the array is sorted, nums[left] paired with the largest remaining element is already too small, so nums[left] can't be in any answer; discarding it loses nothing. The sliding window uses the same kind of argument: a window that is invalid at left stays invalid for every larger right, so left can be discarded forever. In both patterns each pointer moves in one direction only, which is what gives O(n). The difference is the shape: Chapter 3's pointers close in on a pair from both ends, a window's pointers move in the same direction and the elements between them all count.",
			},
			{
				Prompt: "Which of these needs a fixed-size window rather than a variable one?",
				Choices: []string{
					"Longest substring with at most k distinct characters",
					"Shortest subarray with sum at least target",
					"Does s2 contain a rearrangement of s1 as a substring?",
					"Longest substring without repeating characters",
				},
				Answer:      2,
				Explanation: "A rearrangement of s1 has exactly len(s1) characters, so every candidate window has that length and the window just slides (Permutation in String). The other three ask for the longest or shortest window, so the length is the unknown and the window must grow and shrink.",
			},
		},
	}
}
