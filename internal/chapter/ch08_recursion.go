package chapter

import "time"

// chapter08 teaches recursion and divide and conquer: the foundation for
// trees (9), backtracking (12), graph DFS (13) and dynamic programming (15).
// Storyboards (IDs registered in app.js): call-stack, merge-sort, quicksort.
func chapter08() Chapter {
	return Chapter{
		Number:        8,
		Slug:          "recursion",
		Title:         "Recursion and Divide and Conquer",
		Category:      "",
		Prerequisites: []int{1},
		Summary: "Solve a problem by solving a smaller copy of the same problem. You learn to write a base case, " +
			"trust the recursive call as a contract, read the call stack, price a recursion tree, and use divide " +
			"and conquer (merge sort, quicksort) to turn O(n^2) work into O(n log n). Trees, backtracking, graph DFS " +
			"and dynamic programming are all built on this chapter.",
		Motivation: "You need to sort 1,000,000 records. The obvious method, selection sort, scans the unsorted part for " +
			"the minimum again and again: n(n-1)/2 = 499,999,500,000 comparisons, about 5 * 10^11. At 100 million " +
			"comparisons per second that is more than an hour and a half. Merge sort does the same job with about " +
			"n * log2(n) = 1,000,000 * 20 = 2 * 10^7 comparisons, a fraction of a second, which is 25,000 times less work. " +
			"The trick is not a clever loop. It is one sentence: \"to sort a list, sort its left half, sort its right " +
			"half, and merge the two sorted halves\". That sentence uses the word *sort* to define sorting, and it works " +
			"because each half is a smaller copy of the same problem, and a list of one item is already sorted. " +
			"This chapter teaches you to write and trust such definitions, to see what they cost, and to know when a loop is better.",
		Objectives: []string{
			"Write a recursive function by stating what it returns, handling the base case, trusting the recursive call and combining the results",
			"Trace a recursive call on a small input by drawing the call stack, including the order in which frames are pushed and popped",
			"Draw the recursion tree of a function and estimate its cost as (number of calls) x (work per call), e.g. explain why naive Fibonacci is exponential",
			"Explain why merge sort is O(n log n) in every case and why quicksort is O(n log n) on average but O(n^2) on sorted input with a first-element pivot",
			"Convert a recursive function to an iterative one (accumulator or explicit stack) and say when Python's recursion limit makes that necessary",
		},
		WhenToUse: []string{
			"The input is itself recursive: a tree, a nested list, a linked list (\"the rest of the list\" is again a list)",
			"The problem can be stated in terms of a smaller instance of itself: \"the answer for n uses the answer for n-1 or n/2\"",
			"You can split the input into halves, solve each half independently and combine the two answers (sorting, counting inversions, closest pair)",
			"You must explore every choice (all subsets, all orderings): each choice leads to a smaller problem of the same kind (backtracking, chapter 12)",
		},
		Concepts: []Concept{
			{
				Name:        "Base case and recursive case",
				Description: "Every recursive function has two parts: a base case small enough to answer directly, and a recursive case that reduces the input and calls the function again.",
				WhyItWorks: "A recursive function is correct for the same reason a proof by induction is. If the base case returns the right answer for the smallest inputs, " +
					"and the recursive case returns the right answer for size n whenever the call on the smaller input is right, then by induction it is right for every size. " +
					"Termination needs one more fact: every recursive call must move strictly closer to a base case (n-1, n//2, a shorter slice), so the chain of calls " +
					"cannot go on forever. If some input can skip past the base case (n jumps from 1 to -1), the induction has a hole and the function never stops.",
				Intuition: "Think of a line of people where you want to know your position. You ask the person in front of you \"what is your position?\" and add 1. " +
					"They do the same. The first person in line (the base case) knows the answer without asking: position 1. " +
					"Each person = one call, asking forward = the recursive call, adding 1 = combining, the first person = the base case.",
				CoreIdeas: []string{
					"The base case is the smallest input whose answer you know without recursion: an empty list (sum 0), n <= 1 (factorial 1), a None node (depth 0)",
					"The recursive case must shrink the input on every call, so it is guaranteed to reach a base case",
					"Some functions need two base cases (fib(0) = 0 and fib(1) = 1) because the recursive case looks two steps back",
					"Check the base case first, before touching the input, so the empty or None case never crashes",
				},
				CommonMistakes: []string{
					"A base case that can be skipped: count_down(n - 2) with `if n == 0` never stops for odd n. Use `if n <= 0` so every path lands on it.",
					"Forgetting to return the recursive result: writing `sum_list(nums[1:])` instead of `return nums[0] + sum_list(nums[1:])` makes the function return None. Every branch must return a value.",
					"Recursing on an input that does not shrink (calling f(n) from f(n)), which recurses until Python raises RecursionError. Check that each call's argument is strictly smaller.",
				},
				Examples: []CodeExample{
					{
						ID:       "factorial",
						Title:    "Factorial: one base case, one recursive call",
						Language: "python",
						Snippet: `def factorial(n):
    if n <= 1:          # base case
        return 1
    return n * factorial(n - 1)   # recursive case

# factorial(5) = 5 * 4 * 3 * 2 * 1 = 120`,
						Notes: "n - 1 shrinks the input on every call, so n <= 1 is always reached for n >= 0.",
					},
					{
						ID:       "recursive_sum",
						Title:    "Sum a list recursively",
						Language: "python",
						Snippet: `def sum_list(nums):
    if len(nums) == 0:  # base case: empty list sums to 0
        return 0
    return nums[0] + sum_list(nums[1:])

# sum_list([1, 2, 3, 4]) = 1 + 2 + 3 + 4 = 10`,
						Notes: "Each call removes one element. (nums[1:] copies the list, so this is O(n^2) time; passing an index instead makes it O(n).)",
					},
					{
						ID:       "count_digits",
						Title:    "Count digits in a number",
						Language: "python",
						Snippet: `def count_digits(n):
    n = abs(n)
    if n < 10:          # base case: a single digit
        return 1
    return 1 + count_digits(n // 10)

# count_digits(12345) = 5`,
						Notes: "n // 10 drops the last digit, so the input shrinks by one digit per call.",
					},
				},
			},
			{
				Name:        "Trust the recursive call (the leap of faith)",
				Description: "When writing the recursive case, assume the call on the smaller input already returns the correct answer, exactly as its contract says, and only decide how to use it.",
				WhyItWorks: "Tracing every call in your head does not scale: a call on a tree of 1,000 nodes makes 1,000 more calls. Instead, write the function's contract in one sentence " +
					"(\"height(node) returns the number of nodes on the longest path from node down to a leaf\"). When you write the recursive case, the call on a smaller input " +
					"is covered by the induction hypothesis: if your base case is right and your combining step is right assuming correct sub-answers, the whole function is right. " +
					"So the only questions you need to answer are: what is the smallest input, and how do I build my answer from the answers for the smaller pieces? " +
					"This is the same reasoning as a proof by induction, and it is what makes tree problems in chapter 9 short.",
				Intuition: "A manager asked for a report on a 300-person department does not interview everyone. She asks each of her 3 team leads for their team's report, " +
					"trusts those reports, and writes a one-page summary. The manager = the current call, the team leads = recursive calls on smaller inputs, " +
					"trusting their reports = the leap of faith, the summary = combining the results, and an individual contributor with no reports = the base case.",
				CoreIdeas: []string{
					"Write the contract first: what does the function return, for which input? Keep it to one sentence",
					"Inside the recursive case, read the recursive call as \"the correct answer for the smaller input\" and do not trace into it",
					"Your only jobs are the base case and the combine step; induction does the rest",
					"Trace a tiny input (size 0, 1, 2) afterwards to check the base case and the combine step, not to understand the whole recursion",
				},
				CommonMistakes: []string{
					"Trying to trace every level before writing anything, and getting lost. Write the contract, trust it, and only trace sizes 0 to 2 to check.",
					"Letting the code drift from the contract: the contract says \"returns the height\" but the base case returns -1 for None in one place and 0 in another. Each branch must honour the same contract.",
					"Doing the sub-problem's work again in the caller (looping over the rest of the list after calling sum_list on it). Trust the call; it already did that work.",
				},
				Examples: []CodeExample{
					{
						ID:       "tree_height",
						Title:    "Height of a binary tree",
						Language: "python",
						Snippet: `def height(node):
    # Contract: number of nodes on the longest path from node down to a leaf
    if node is None:            # base case: empty tree has height 0
        return 0
    left_h = height(node.left)    # trust: correct height of the left subtree
    right_h = height(node.right)  # trust: correct height of the right subtree
    return 1 + max(left_h, right_h)  # combine`,
						Notes: "You never trace the subtrees: you trust their heights and add 1 for this node.",
					},
					{
						ID:       "linked_list_length",
						Title:    "Length of a linked list",
						Language: "python",
						Snippet: `def length(node):
    if node is None:            # base case: empty list
        return 0
    return 1 + length(node.next)   # this node plus the (trusted) length of the rest`,
						Notes: "\"The rest of the list\" is again a list, so the same function answers it.",
					},
					{
						ID:       "reverse_string",
						Title:    "Reverse a string",
						Language: "python",
						Snippet: `def reverse(s):
    if len(s) <= 1:             # base case
        return s
    return reverse(s[1:]) + s[0]   # trust: reverse(s[1:]) is the rest reversed

# reverse("hello") = reverse("ello") + "h" = "olle" + "h" = "olleh"`,
						Notes: "Trust that reverse(\"ello\") is \"olle\"; then the combine step is obvious.",
					},
				},
			},
			{
				Name:        "The call stack",
				Description: "Each call gets its own frame (its own copies of its parameters and local variables) pushed on the call stack; the frame is popped when the call returns.",
				WhyItWorks: "When factorial(4) calls factorial(3), factorial(4) is not finished: it still has to multiply by 4. So Python saves its frame (n = 4, \"resume at the multiplication\") " +
					"on a stack and starts a new frame for factorial(3). Frames pile up until a base case returns without calling anything; then they are popped in reverse order, " +
					"each one finishing its own pending work with the value it got back. Because frames are separate, n = 4 and n = 3 never overwrite each other. " +
					"The space cost of recursion is the maximum number of frames alive at once, the recursion depth, not the total number of calls. " +
					"Python caps the depth at about 1,000 frames (sys.getrecursionlimit()) and raises RecursionError beyond that, so recursing once per element of a 100,000-element list fails.",
				Intuition: "A stack of plates in a cafeteria: each new call puts a plate on top with its own notes written on it. You can only work on the top plate. " +
					"When a call returns, its plate is removed and you continue with the notes on the plate underneath. " +
					"Plate = stack frame, notes = parameters and local variables, top plate = the running call, removing a plate = returning a value.",
				CoreIdeas: []string{
					"Calls are pushed on the way down (until a base case) and popped on the way back up, last in first out",
					"Code before the recursive call runs on the way down; code after it runs on the way back up, in reverse order",
					"Extra space is O(depth): O(n) for one call per element, O(log n) when the input halves each time",
					"Python's default limit is about 1,000 frames; deeper recursion needs an iterative rewrite (or sys.setrecursionlimit, with care)",
				},
				CommonMistakes: []string{
					"Thinking recursion uses no memory because there is no explicit list. Every pending call holds a frame, so a depth-n recursion uses O(n) space; count it in your space complexity.",
					"Recursing once per element on large inputs (a 10^5-node linked list) and hitting RecursionError. Use a loop when the depth can grow with n, and recursion when the depth is O(log n) or small.",
					"Expecting code after the recursive call to run in the same order as the calls. It runs in reverse (unwinding) order: print before the call to see 5, 3, 1 and after it to see 1, 3, 5.",
				},
				Examples: []CodeExample{
					{
						ID:       "call_stack_trace",
						Title:    "Watching frames push and pop",
						Language: "python",
						Snippet: `def f(n):
    if n <= 0:
        return 0
    print(n)            # runs on the way down
    result = n + f(n - 2)
    print("back", n)    # runs on the way back up
    return result

f(5)   # prints 5, 3, 1, back 1, back 3, back 5 and returns 9`,
						Notes: "Three frames (n = 5, 3, 1) are alive at the deepest point, plus the base-case call f(-1).",
					},
				},
			},
			{
				Name:        "Recursion trees and their cost",
				Description: "Draw each call as a node and the calls it makes as its children. The running time is the sum of the work done at every node of this tree.",
				WhyItWorks: "Every call does some work of its own (comparisons, a merge, an addition) plus its recursive calls. Adding the work of every node in the recursion tree counts " +
					"everything exactly once, so time = (number of calls) x (work per call) when the work per call is the same, or the sum over levels when it is not. " +
					"Naive fib(n) calls fib(n-1) and fib(n-2), so the tree branches twice per level and has about 1.6^n nodes: fib(30) makes 2,692,537 calls and fib(40) makes 331,160,281, " +
					"mostly recomputing the same values. Fast power(b, n) makes one call on n/2, so its tree is a path of about log2(n) nodes. " +
					"The tree also gives the space: the stack holds one root-to-leaf path at a time, so space = tree height.",
				Intuition: "A phone tree for a school closure: the principal calls 2 teachers, each calls 2 parents, each parent calls 2 more. Counting all phone calls = counting nodes " +
					"in the recursion tree; the number of rounds = the tree height. If every call branches twice and the rounds go n deep, you get 2^n calls; if each call makes only one smaller call of half the size, you get log n.",
				CoreIdeas: []string{
					"One recursive call per invocation on n-1: a chain of n calls, O(n) if each does O(1) work",
					"One call on n/2: a chain of log n calls (binary search, fast power), O(log n)",
					"Two calls on n-1 or n-2: a branching tree with exponential calls (naive Fibonacci); identical subtrees mean memoization (chapter 15) can collapse it to O(n)",
					"Two calls on n/2 with O(n) work each level (merge sort): log n levels x O(n) per level = O(n log n)",
				},
				CommonMistakes: []string{
					"Counting only the depth: fib(n) has depth n but about 1.6^n calls. Time is the number of nodes in the tree, space is its height.",
					"Ignoring hidden work per call: sum_list(nums[1:]) copies the slice each time, so n calls x O(n) copy = O(n^2). Pass an index to keep each call O(1).",
					"Calling the same recursive function twice with the same argument (power(b, n//2) * power(b, n//2)) and turning an O(log n) chain into an O(n) tree. Store the result in a variable and reuse it.",
				},
				Examples: []CodeExample{
					{
						ID:       "fibonacci",
						Title:    "Naive Fibonacci: an exponential recursion tree",
						Language: "python",
						Snippet: `def fib(n):
    if n <= 1:          # base cases: fib(0) = 0, fib(1) = 1
        return n
    return fib(n - 1) + fib(n - 2)

# fib(6) = 8, but fib(30) already makes 2,692,537 calls`,
						Notes: "fib(2) is recomputed many times in the tree. Chapter 15 caches each answer so every fib(k) is computed once.",
					},
					{
						ID:       "power_recursive",
						Title:    "Fast power: a chain of log n calls",
						Language: "python",
						Snippet: `def power(base, exp):
    if exp == 0:        # base case
        return 1
    half = power(base, exp // 2)   # one call on half the exponent
    if exp % 2 == 0:
        return half * half
    return half * half * base

# power(2, 10): exponents 10, 5, 2, 1, 0 -> 5 calls, returns 1024`,
						Notes: "Storing `half` matters: calling power twice would double the calls at every level.",
					},
					{
						ID:       "binary_search_recursive",
						Title:    "Binary search, recursively",
						Language: "python",
						Snippet: `def binary_search(arr, target, lo, hi):
    if lo > hi:                 # base case: empty range, not found
        return -1
    mid = (lo + hi) // 2
    if arr[mid] == target:
        return mid              # base case: found
    if arr[mid] < target:
        return binary_search(arr, target, mid + 1, hi)
    return binary_search(arr, target, lo, mid - 1)`,
						Notes: "One call on half the range: O(log n) time and O(log n) stack depth (the loop version from chapter 6 uses O(1) space).",
					},
				},
			},
			{
				Name:        "Divide and conquer: merge sort",
				Description: "Divide the input into independent parts, conquer each part recursively, and combine the part answers into the whole answer. Merge sort splits in half, sorts each half and merges.",
				WhyItWorks: "Correctness: by the leap of faith both halves come back sorted, and merging two sorted lists produces a sorted list, because the smallest remaining item overall " +
					"is always at the front of one of the two lists, so taking the smaller front each time appends items in order. " +
					"Cost: halving 1,000,000 items reaches size 1 after about log2(1,000,000) = 20 levels. On each level the merges together touch every item once, O(n) per level. " +
					"So the total is O(n) x log n levels = O(n log n), in the best, average and worst case alike, because the split is always exactly in half. " +
					"The merge writes into a new list, so merge sort needs O(n) extra memory, plus O(log n) stack depth.",
				Intuition: "Sorting a pile of 1,000 exam papers with a friend: split the pile, each of you sorts your half (and splits again with more friends, down to single papers, " +
					"which are already sorted), then you merge two sorted piles by repeatedly taking the top paper with the smaller name. " +
					"Splitting = Divide, friends sorting their halves = Conquer, taking the smaller top paper = Combine.",
				CoreIdeas: []string{
					"Divide: mid = len(arr) // 2, split into arr[:mid] and arr[mid:]",
					"Conquer: sort both halves recursively; a list of 0 or 1 items is the base case",
					"Combine: merge with two pointers, always taking the smaller front item (`<=` keeps equal items in their original order, so the sort is stable)",
					"O(n log n) time in every case, O(n) extra space",
				},
				CommonMistakes: []string{
					"Forgetting to append the leftovers after the merge loop: when one list runs out, the rest of the other list is still unsorted output. Add result.extend(left[i:]) and result.extend(right[j:]).",
					"Using `mid = len(arr) // 2` with a base case of `len(arr) == 0` only: a 1-element list splits into [] and [x] forever. The base case must be len(arr) <= 1.",
					"Thinking merge sort is O(n^2) because there are n - 1 merges each costing up to O(n). Most merges are tiny; count per level instead: each level costs O(n) and there are log n levels.",
				},
				Examples: []CodeExample{
					{
						ID:       "merge_sort",
						Title:    "Merge sort",
						Language: "python",
						Snippet: `def merge_sort(arr):
    if len(arr) <= 1:              # base case
        return arr
    mid = len(arr) // 2            # Divide
    left = merge_sort(arr[:mid])   # Conquer
    right = merge_sort(arr[mid:])  # Conquer
    return merge(left, right)      # Combine

def merge(left, right):
    result = []
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result`,
						Notes: "The merge is the two-pointer technique from chapter 3 applied to two sorted lists.",
					},
				},
			},
			{
				Name:        "Divide and conquer: quicksort and pivot choice",
				Description: "Quicksort does the work before recursing: it partitions the array around a pivot (smaller items left, larger right), then sorts the two sides recursively. No combine step is needed.",
				WhyItWorks: "After partitioning, the pivot is in its final sorted position: everything left of it is <= pivot and everything right is >= pivot, so sorting each side independently " +
					"sorts the whole array. Cost depends on the split. If the pivot lands near the middle, the sizes halve, there are about log n levels of O(n) partition work: O(n log n). " +
					"If the pivot is always the smallest item (first-element pivot on already sorted input), one side is empty and the other has n - 1 items: n levels with n, n-1, ..., 1 " +
					"comparisons, about n^2/2. Sorting [0..999] that way takes 499,500 comparisons and 999 nested calls, right at Python's recursion limit. " +
					"A random pivot makes a bad split on every level extremely unlikely, giving O(n log n) expected time and O(log n) expected stack depth, and it sorts in place with no extra array.",
				Intuition: "Sorting students by height: pick one student as the pivot, send everyone shorter to the left wall and everyone taller to the right wall. The pivot is now standing in their final spot. " +
					"Repeat inside each group. Picking the pivot = choose-pivot, walking to a wall = partition, repeating inside each group = recurse. If you always pick the shortest student, one wall stays empty and you make no progress.",
				CoreIdeas: []string{
					"Divide = partition around the pivot; this is where all the comparisons happen",
					"Conquer = quicksort the left side and the right side; the pivot itself is already placed",
					"Combine = nothing: the array is sorted in place once both sides are",
					"Average O(n log n), worst O(n^2); a random pivot (random.randrange) avoids the sorted-input worst case",
				},
				CommonMistakes: []string{
					"Always choosing the first (or last) element as the pivot. On sorted or reverse-sorted input every split is 0 / n-1 and the sort becomes O(n^2) with O(n) recursion depth. Pick a random pivot.",
					"Including the pivot in a recursive call (quicksort(nums, lo, p) instead of p - 1). When the pivot is the largest item the range never shrinks and the recursion never ends.",
					"Assuming quicksort is O(n log n) in the worst case like merge sort. It is only the average; say \"expected O(n log n)\" and know the input that breaks it.",
				},
				Examples: []CodeExample{
					{
						ID:       "quicksort",
						Title:    "Quicksort, readable version",
						Language: "python",
						Snippet: `def quicksort(nums):
    if len(nums) <= 1:
        return nums
    pivot = nums[len(nums) // 2]
    less = [x for x in nums if x < pivot]
    equal = [x for x in nums if x == pivot]
    greater = [x for x in nums if x > pivot]
    return quicksort(less) + equal + quicksort(greater)`,
						Notes: "Clear but uses O(n) extra lists per level. The in-place version partitions with two indices.",
					},
					{
						ID:       "pivot_choice",
						Title:    "Randomized pivot selection",
						Language: "python",
						Snippet: `import random

def choose_pivot(nums, lo, hi):
    k = random.randrange(lo, hi + 1)
    nums[k], nums[hi] = nums[hi], nums[k]   # move it to the end, then partition as usual
    return nums[hi]`,
						Notes: "No fixed input can force bad splits every time, so the expected time is O(n log n).",
					},
				},
			},
			{
				Name:        "Recursion versus iteration",
				Description: "Anything written recursively can be written with a loop, using an accumulator for a single chain of calls or an explicit stack for branching calls.",
				WhyItWorks: "A recursive call only does two things a loop cannot do by itself: it remembers pending work (the frame) and returns to it later. " +
					"If no work is pending after the call (a tail call, like factorial_tail(n - 1, n * acc)), nothing needs remembering, so the call is just \"update the variables and loop\". " +
					"If work is pending at many branches (tree traversal), push that work onto your own list used as a stack; the list plays the role of the call stack. " +
					"Python does not optimise tail calls, so tail recursion still uses one frame per call. Prefer recursion when the depth is small (balanced trees, O(log n) divide and conquer) " +
					"and the code is clearer; prefer a loop when the depth grows with n (linked lists, 10^5-element arrays).",
				Intuition: "Recursion is like leaving sticky notes for yourself on a stack of trays (\"come back and multiply by 4\"); iteration is carrying a running total in your hand. " +
					"If the note would say nothing (a tail call), you don't need the tray at all. Sticky notes = frames, running total = accumulator, your own tray stack = an explicit stack.",
				CoreIdeas: []string{
					"Tail recursion: carry the partial answer in an accumulator parameter, then turn the call into a loop",
					"Branching recursion: replace the call stack with an explicit list, push children, pop to process",
					"Python has no tail-call optimisation and a ~1,000-frame limit, so linear-depth recursion over large inputs should be a loop",
					"Choose recursion for recursive structures with small depth, where it makes the code shorter and easier to prove correct",
				},
				CommonMistakes: []string{
					"Believing a tail-recursive function is safe from RecursionError in Python. It is not; rewrite it as a while loop for deep inputs.",
					"Pushing children onto an explicit stack in the wrong order and getting a different traversal. A stack is last in first out, so push the right child before the left to visit left first.",
				},
				Examples: []CodeExample{
					{
						ID:       "factorial_tail",
						Title:    "Tail-recursive factorial with an accumulator",
						Language: "python",
						Snippet: `def factorial_tail(n, acc=1):
    if n <= 1:
        return acc
    return factorial_tail(n - 1, n * acc)

# factorial_tail(5) -> (4, 5) -> (3, 20) -> (2, 60) -> (1, 120) -> 120`,
						Notes: "Nothing is left to do after the call returns, so this is directly a loop.",
					},
					{
						ID:       "factorial_iterative",
						Title:    "Factorial: recursive vs iterative",
						Language: "python",
						Snippet: `def fact_rec(n):
    if n <= 1:
        return 1
    return n * fact_rec(n - 1)

def fact_iter(n):
    result = 1
    for i in range(2, n + 1):
        result *= i
    return result`,
						Notes: "Same answer; the loop uses O(1) space instead of O(n) frames.",
					},
					{
						ID:       "tree_iterative",
						Title:    "Inorder traversal: recursive vs explicit stack",
						Language: "python",
						Snippet: `def inorder_rec(node, result):
    if node:
        inorder_rec(node.left, result)
        result.append(node.val)
        inorder_rec(node.right, result)

def inorder_iter(root):
    result, stack = [], []
    current = root
    while current or stack:
        while current:              # go left, remembering the path
            stack.append(current)
            current = current.left
        current = stack.pop()       # the deepest unvisited node
        result.append(current.val)
        current = current.right
    return result`,
						Notes: "The list `stack` holds exactly the frames the recursive version would have pending.",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Sum a list recursively (and watch the call stack)",
				Problem: "Write sum_list(nums, i) that returns the sum of nums[i:], without a loop. Trace it on nums = [3, 1, 4] starting at i = 0.",
				Steps: []WorkedStep{
					{
						Subgoal:     "State what the function returns",
						Explanation: "Write the contract before any code: sum_list(nums, i) returns nums[i] + nums[i+1] + ... + nums[-1], the sum of the suffix starting at i. Passing an index instead of slicing keeps each call O(1).",
						Code:        "def sum_list(nums, i=0):\n    # returns the sum of nums[i:]",
						State:       "contract: sum_list([3,1,4], 0) = 8, sum_list([3,1,4], 1) = 5, sum_list([3,1,4], 3) = 0",
					},
					{
						Subgoal:     "Handle the base case",
						Explanation: "The smallest suffix is the empty one, i == len(nums), and the sum of nothing is 0. Every call increases i by 1, so this base case is always reached.",
						Code:        "    if i == len(nums):\n        return 0",
						State:       "sum_list([3,1,4], 3) -> i == 3 == len -> return 0",
					},
					{
						Subgoal:     "Trust the recursive call",
						Explanation: "By the contract, sum_list(nums, i + 1) is the sum of everything after position i. Do not trace into it; just name it.",
						Code:        "    rest = sum_list(nums, i + 1)",
						State:       "inside sum_list([3,1,4], 0): rest = sum_list([3,1,4], 1), which the contract says is 1 + 4 = 5",
					},
					{
						Subgoal:     "Combine the results",
						Explanation: "The suffix from i is nums[i] followed by the suffix from i + 1, so its sum is nums[i] + rest.",
						Code:        "    return nums[i] + rest",
						State:       "stack on the way down: [i=0] -> [i=0, i=1] -> [i=0, i=1, i=2] -> [i=0, i=1, i=2, i=3]; popping: i=3 returns 0, i=2 returns 4 + 0 = 4, i=1 returns 1 + 4 = 5, i=0 returns 3 + 5 = 8",
					},
				},
				Code: `def sum_list(nums, i=0):
    # returns the sum of nums[i:]
    if i == len(nums):
        return 0
    rest = sum_list(nums, i + 1)
    return nums[i] + rest`,
			},
			{
				Title:   "Merge sort on [5, 2, 4, 1]",
				Problem: "Sort a list of numbers in O(n log n) time with divide and conquer. Trace merge_sort([5, 2, 4, 1]).",
				Steps: []WorkedStep{
					{
						Subgoal:     "State what the function returns",
						Explanation: "merge_sort(arr) returns a new list with the items of arr in ascending order. Writing this down is what lets us trust the two recursive calls later.",
						Code:        "def merge_sort(arr):\n    # returns a sorted copy of arr",
						State:       "merge_sort([5,2,4,1]) should be [1,2,4,5]",
					},
					{
						Subgoal:     "Handle the base case",
						Explanation: "A list with 0 or 1 items is already sorted. It must be len(arr) <= 1, not == 0: splitting a 1-item list gives [] and the same 1-item list, which would recurse forever.",
						Code:        "    if len(arr) <= 1:\n        return arr",
						State:       "merge_sort([5]) -> [5]; merge_sort([2]) -> [2]",
					},
					{
						Subgoal:     "Divide",
						Explanation: "Split at the middle so both halves are about n/2. Halving is what gives only log2(n) levels.",
						Code:        "    mid = len(arr) // 2",
						State:       "[5,2,4,1]: mid = 2 -> left part [5,2], right part [4,1]",
					},
					{
						Subgoal:     "Conquer",
						Explanation: "Trust the recursive call: by the contract each call returns its half sorted. Each half is smaller, so the recursion ends.",
						Code:        "    left = merge_sort(arr[:mid])\n    right = merge_sort(arr[mid:])",
						State:       "merge_sort([5,2]) splits into [5] and [2] and returns [2,5]; merge_sort([4,1]) returns [1,4]",
					},
					{
						Subgoal:     "Combine",
						Explanation: "Merge two sorted lists with two pointers: the smallest remaining item is at the front of one of them, so take the smaller front each time, then append whatever is left. Each item is copied once, so the merge is O(len(left) + len(right)).",
						Code:        "    return merge(left, right)",
						State:       "merge([2,5], [1,4]): take 1 (from right), 2, 4, then right is empty, append leftover 5 -> [1,2,4,5]",
					},
					{
						Subgoal:     "Count the cost",
						Explanation: "Level 0 merges 4 items, level 1 does two merges of 2 items (4 total). Every level touches all n items once and there are log2(n) levels: O(n log n). For n = 1,000,000 that is about 20 levels x 1,000,000 = 2 * 10^7 steps.",
						State:       "levels: [5,2,4,1] -> [5,2] [4,1] -> [5] [2] [4] [1]; merge work: 2 + 2 at level 1, 4 at level 0 -> 8 = 4 items x 2 levels",
					},
				},
				Code: `def merge_sort(arr):
    if len(arr) <= 1:
        return arr
    mid = len(arr) // 2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])
    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Fast power",
				Prompt: "power(base, exp) returns base ** exp for exp >= 0 using O(log exp) multiplications. Use the same subgoals: handle the base case, trust the recursive call on exp // 2, combine the results (remember an odd exponent needs one extra factor of base).",
				Code: `def power(base, exp):
    # State what the function returns: base ** exp
    if exp == 0:                     # Handle the base case
        return ____
    half = power(base, ____)         # Trust the recursive call
    if exp % 2 == 0:                 # Combine the results
        return ____
    return ____`,
				Blanks:      []string{"1", "exp // 2", "half * half", "half * half * base"},
				Explanation: "Anything to the power 0 is 1. The call on exp // 2 is trusted to return base ** (exp // 2). If exp is even, exp = 2 * (exp // 2), so the answer is half * half. If exp is odd, exp = 2 * (exp // 2) + 1, so it needs one more factor: half * half * base. power(2, 10) visits exponents 10, 5, 2, 1, 0 and returns 1024; storing half avoids making the same call twice.",
			},
			{
				Title:  "Quicksort in place",
				Prompt: "Sort nums[lo..hi] in place. Partition around the last element (Divide), then quicksort each side of the pivot (Conquer). Nothing is left to combine.",
				Code: `def quicksort(nums, lo, hi):
    if ____:                         # Handle the base case: 0 or 1 items
        return
    pivot = nums[hi]                 # Divide: partition around the pivot
    i = lo                           # nums[lo:i] holds items < pivot
    for j in range(lo, hi):
        if ____:
            nums[i], nums[j] = nums[j], nums[i]
            i += 1
    nums[i], nums[hi] = nums[hi], nums[i]   # pivot lands in its final spot
    quicksort(nums, lo, ____)        # Conquer the left side
    quicksort(nums, ____, hi)        # Conquer the right side`,
				Blanks:      []string{"lo >= hi", "nums[j] < pivot", "i - 1", "i + 1"},
				Explanation: "A range with lo >= hi has at most one item and is sorted. The loop keeps the invariant that nums[lo:i] are all smaller than the pivot, so each smaller item is swapped to position i. After the swap the pivot sits at index i in its final place, so the recursive calls must exclude it: lo..i-1 and i+1..hi. Including i would let a range fail to shrink. On [3,6,1,5,2,4] the first partition gives [3,1,2,4,6,5] with the pivot 4 at index 3.",
			},
		},
		Animations: []Storyboard{
			{
				ID:    "call-stack",
				Title: "Recursive Call Stack",
				Goal:  "Watch factorial(4) push a frame per call, hit the base case, and pop frames while multiplying on the way back up.",
				MemoryTips: []string{
					"State what the function returns: factorial(n) = n * (n-1) * ... * 1",
					"Handle the base case first: n <= 1 returns 1 without calling anything",
					"Trust the recursive call: factorial(n - 1) is the right answer for n - 1",
					"Combine the results on the way back up: multiply by n, then pop the frame",
					"Space = deepest stack (n frames here), not the total number of calls",
				},
				Steps: []StoryboardStep{
					{Cue: "push", Narration: "factorial(4) cannot answer yet: it needs factorial(3). Its frame (n = 4, waiting to multiply by 4) stays on the stack and a new frame starts for n = 3, then n = 2, then n = 1. Each frame has its own n.", VisualHint: "Stack grows downward: factorial(4), factorial(3), factorial(2), factorial(1)", Duration: Duration(3 * time.Second), CodeRef: "recursion: call self"},
					{Cue: "base-case", Narration: "factorial(1) hits the base case n <= 1 and returns 1 without making another call. This is the deepest point: 4 frames alive at once, which is the O(n) space the recursion uses.", VisualHint: "Highlight terminal frame", Duration: Duration(3 * time.Second), CodeRef: "recursion: base check"},
					{Cue: "pop", Narration: "Frames pop in reverse order, each finishing its pending multiplication: factorial(2) returns 2 * 1 = 2, factorial(3) returns 3 * 2 = 6, factorial(4) returns 4 * 6 = 24.", VisualHint: "Stack shrinks upward with return values shown", Duration: Duration(3 * time.Second), CodeRef: "recursion: return unwind"},
				},
			},
			{
				ID:    "merge-sort",
				Title: "Merge Sort Divide and Conquer",
				Goal:  "See merge sort split [38, 27, 43, 3, 9, 82, 10] down to single items and merge sorted halves back into [3, 9, 10, 27, 38, 43, 82].",
				MemoryTips: []string{
					"Divide: split at mid = len(arr) // 2",
					"Handle the base case: 0 or 1 items is already sorted",
					"Conquer: trust merge_sort to return each half sorted",
					"Combine: merge by taking the smaller front item, then append the leftovers",
					"Cost: log n levels x O(n) merge work per level = O(n log n), always",
				},
				Steps: []StoryboardStep{
					{Cue: "divide", Narration: "Split the 7 items at mid = 3 into [38, 27, 43] and [3, 9, 82, 10], and keep halving each part. Because sizes halve, there are only about log2(7) ~ 3 levels.", VisualHint: "Split animation", Duration: Duration(3 * time.Second), CodeRef: "merge_sort: mid = len(arr) // 2"},
					{Cue: "base", Narration: "A part with one item, like [38] or [27], is already sorted, so it is returned as is. This base case is what stops the splitting.", VisualHint: "Base case", Duration: Duration(2 * time.Second), CodeRef: "merge_sort: if len(arr) <= 1"},
					{Cue: "merge", Narration: "Merge two sorted parts by comparing their front items and taking the smaller: [27, 38, 43] and [3, 9, 10, 82] give 3, 9, 10, 27, ... Each level of merges touches every item once: O(n) per level.", VisualHint: "Merge animation", Duration: Duration(3 * time.Second), CodeRef: "merge: compare and combine"},
					{Cue: "complete", Narration: "The top-level merge returns [3, 9, 10, 27, 38, 43, 82]. Total work: O(n) per level times log n levels = O(n log n), with O(n) extra memory for the merged lists.", VisualHint: "Sorted result", Duration: Duration(2 * time.Second), CodeRef: "merge_sort: return result"},
				},
			},
			{
				ID:    "quicksort",
				Title: "Quicksort Partitioning",
				Goal:  "Watch quicksort place a pivot in its final position, then sort each side of it, on [38, 27, 43, 3, 9, 82, 10].",
				MemoryTips: []string{
					"Handle the base case: a range with lo >= hi is sorted",
					"Divide: choose a pivot and partition (smaller items left, larger right); the pivot is now in its final spot",
					"Conquer: quicksort the left side and the right side, excluding the pivot",
					"Combine: nothing to do, the array is sorted in place",
					"Pick a random pivot: a first or last pivot on sorted input gives O(n^2)",
				},
				Steps: []StoryboardStep{
					{Cue: "choose-pivot", Narration: "Choose the last item, 10, as the pivot. A good pivot lands near the middle; a pivot that is the smallest or largest item makes a useless split.", VisualHint: "Highlight pivot", Duration: Duration(2 * time.Second), CodeRef: "quicksort: pivot"},
					{Cue: "partition", Narration: "Sweep j across the range, swapping every item <= 10 to the front. Only 3 and 9 qualify, then the pivot swaps into index 2: [3, 9, 10, 38, 27, 82, 43]. 10 will never move again.", VisualHint: "Two-pointer sweep", Duration: Duration(4 * time.Second), CodeRef: "quicksort: partition lists"},
					{Cue: "recurse", Narration: "Sort [3, 9] and [38, 27, 82, 43] the same way, leaving the pivot out. Balanced splits give about log n levels of O(n) work, O(n log n) on average; with a random pivot no input forces the O(n^2) worst case.", VisualHint: "Call tree view", Duration: Duration(3 * time.Second), CodeRef: "quicksort: recursive calls"},
				},
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "What happens when you call count_down(5)?\n\ndef count_down(n):\n    if n == 0:\n        return\n    print(n)\n    count_down(n - 2)",
				Choices: []string{
					"It prints 5, 3, 1 and returns",
					"It prints 5, 3, 1, -1, -3, ... until Python raises RecursionError",
					"It prints 5, 3, 1, 0 and returns",
					"Python detects the missing base case and returns None immediately",
				},
				Answer:      1,
				Explanation: "n goes 5, 3, 1, -1: it jumps over 0, so the base case is never reached and the stack grows until RecursionError (about 1,000 frames). Expecting it to stop after 1 is the classic mistake: a base case must catch every path, so write `if n <= 0`.",
			},
			{
				Prompt: "Merge sort makes n - 1 merges and a merge can cost O(n). Why is merge sort O(n log n) and not O(n^2)?",
				Choices: []string{
					"Because the merges on each level of the recursion tree together touch each item once (O(n) per level), and there are only about log2(n) levels",
					"Because each merge costs only O(log n)",
					"It is O(n^2) in the worst case, like quicksort; O(n log n) is only the average",
					"Because merge sort skips the merge when the halves are already sorted",
				},
				Answer:      0,
				Explanation: "Most merges are tiny. Summing per level: level k has 2^k merges of n/2^k items each, so n per level, and halving gives log2(n) levels. Merge sort always splits exactly in half, so unlike quicksort its worst case is also O(n log n). A single merge is O(n), not O(log n), and the standard algorithm never skips merges.",
			},
			{
				Prompt: "Which pattern fits best? Given an array of 100,000 numbers, count the pairs (i, j) with i < j and nums[i] > nums[j] (the number of inversions).",
				Choices: []string{
					"A hash map of values seen so far (chapter 2)",
					"Two pointers from both ends of the array (chapter 3)",
					"Divide and conquer: count inversions inside each half recursively and count the cross pairs while merging (merge sort)",
					"A sliding window over the array (chapter 4)",
				},
				Answer:      2,
				Explanation: "A hash map answers \"have I seen this exact value?\", not \"how many larger values came before?\". Two pointers need sorted input and a monotone rule, and a sliding window needs a contiguous range, but inversions are pairs anywhere. Divide and conquer works: inversions = inversions in the left half + in the right half + pairs across the halves, and while merging two sorted halves, each time an item from the right is taken first it is smaller than every remaining left item, so add len(left) - i. That is O(n log n) instead of ~5 * 10^9 pair checks.",
			},
			{
				Prompt: "Chapter 1 review: about how many times can you halve 1,000,000 before reaching 1, and what does that number tell you about merge sort on 1,000,000 items?",
				Choices: []string{
					"About 1,000; merge sort makes about 1,000 recursive calls",
					"About 20 (log2 of 1,000,000); merge sort has about 20 levels of splitting, so about 20 x 1,000,000 = 2 * 10^7 steps",
					"About 500,000; merge sort halves the list once per element",
					"About 6 (log10 of 1,000,000); merge sort has 6 levels",
				},
				Answer:      1,
				Explanation: "2^20 = 1,048,576, so 20 halvings bring 1,000,000 down to 1. Big O uses log base 2 for halving (the base only changes a constant factor). This is why binary search needs about 20 steps and why merge sort's recursion tree has about 20 levels of O(n) work each.",
			},
			{
				Prompt:      "Trace it: what does mystery(6) return?\n\ndef mystery(n):\n    if n == 0:\n        return \"\"\n    return mystery(n // 2) + str(n % 2)",
				Explanation: "\"110\". mystery(6) = mystery(3) + \"0\"; mystery(3) = mystery(1) + \"1\"; mystery(1) = mystery(0) + \"1\" = \"1\". Unwinding: mystery(3) = \"11\", mystery(6) = \"110\". The function returns the binary representation of n: the call on n // 2 builds the higher bits, and the combine step appends the lowest bit.",
			},
			{
				Prompt:      "What is the leap of faith in recursion, and why is it safe? Use height(node) = 1 + max(height(node.left), height(node.right)) as your example.",
				Explanation: "You state a contract (height(node) returns the number of nodes on the longest downward path from node) and, when writing the recursive case, assume the calls on the smaller subtrees already satisfy it, instead of tracing them. It is safe by induction: the base case (None has height 0) is correct, and if the children's heights are correct then 1 + the larger one is correct for the node; every call is on a strictly smaller subtree, so the recursion reaches the base case.",
			},
			{
				Prompt:      "You quicksort an already sorted list of 1,000 items, always choosing the first element as the pivot. What goes wrong, and how do you fix it?",
				Explanation: "The pivot is always the smallest item, so every partition puts 0 items on the left and all the rest on the right. There are n levels instead of log n, doing n-1, n-2, ..., 1 comparisons: 499,500 for 1,000 items, O(n^2), and the recursion goes 999 calls deep, right at Python's ~1,000-frame limit. Choose the pivot at random (or median-of-three) so balanced splits are expected on any input, giving expected O(n log n) time and O(log n) depth.",
			},
			{
				Prompt: "Naive fib(n) = fib(n - 1) + fib(n - 2) has recursion depth n. What are its time and extra space?",
				Choices: []string{
					"O(n) time and O(n) space, because the depth is n",
					"Exponential time (about 1.6^n calls; fib(30) makes 2,692,537) and O(n) space",
					"O(n log n) time, like merge sort, because it makes two calls",
					"Exponential time and exponential space",
				},
				Answer:      1,
				Explanation: "Time counts every node of the recursion tree, and the tree branches twice per level, so the number of calls grows like 1.6^n. Space counts only the frames alive at once, one root-to-leaf path, so O(n). Merge sort's two calls are on n/2 with O(log n) depth; fib's are on n-1 and n-2, which barely shrink. Caching repeated answers (chapter 15) brings fib down to O(n) time.",
			},
		},
	}
}
