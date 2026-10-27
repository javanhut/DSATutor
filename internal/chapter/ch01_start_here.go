package chapter

import "time"

// chapter01 is the learner's entry point: how to study with the app, the
// problem-solving recipe every later chapter reuses, and Big O.
// Storyboards (IDs registered in app.js): runtime-shapes, selection-sort.
func chapter01() Chapter {
	return Chapter{
		Number:        1,
		Slug:          "start-here",
		Title:         "Start Here: How to Learn Algorithms and Big O",
		Category:      "",
		Prerequisites: []int{},
		Summary: "How this course works and the one tool every later chapter needs. You'll learn the study loop the app is built around " +
			"(worked example, completion problem, recall questions, practice ladder, review), the five-step problem-solving recipe used in " +
			"every chapter, and how to count operations to predict how fast code runs and how much memory it needs (Big O).",
		Motivation: `You have the arrival times of 100,000 messages and want the two that arrived closest together. The first idea most people have is to compare every pair. That is n(n-1)/2 pairs, about 5 billion comparisons for n = 100,000. Python does roughly 10 to 50 million simple steps per second, so this takes several minutes.

A second idea: sort the times first. In sorted order, the closest partner of any time is one of its neighbours, so you only compare neighbours. Sorting takes about n log2 n, roughly 1.7 million steps, and the neighbour pass 100,000 more. That finishes in well under a second.

Same answer, same computer, a difference of minutes against a blink. The difference isn't clever code; it is the shape of the work: one approach grows like n squared, the other like n log n. This chapter teaches you to see that shape before you run anything (Big O), gives you a recipe for finding the faster idea (understand, brute force, find the wasted work, optimise, test), and explains how to use this app so that what you learn sticks.`,
		Objectives: []string{
			"Describe the study loop in this app (chapter, worked example, completion problem, check questions, practice ladder, Review queue, Mixed Practice) and explain why each step is there.",
			"Apply the five-step recipe to a new problem: understand with examples, brute force first, find the wasted work, optimise, test edge cases.",
			"Count the operations of a short program and state its running time in Big O, dropping constants and lower-order terms.",
			"Trace selection sort on a small list and explain why it makes n(n-1)/2 comparisons, which is O(n^2) time and O(1) extra space.",
			"Use a problem's constraints (for example n up to 10^5) to decide which complexities are fast enough.",
		},
		WhenToUse: []string{
			"Read the constraints before coding: n up to about 20 allows O(2^n); up to about 10^3 or 10^4 allows O(n^2); 10^5 to 10^6 needs O(n log n) or O(n).",
			"Your first idea has a loop inside a loop over the same data: that is O(n^2), so ask what work is repeated.",
			"A problem says \"in O(n) time\" or \"with O(1) extra memory\": it is telling you which brute force is not allowed.",
			"Two solutions both pass small tests: compare their Big O to know which survives the hidden large tests.",
		},
		Concepts: []Concept{
			{
				Name:        "How to learn with this app",
				Description: "The study loop: read the chapter, study the worked example, fill in the completion problem, answer the check questions from memory, then climb the practice ladder, and come back through Review and Mixed Practice.",
				WhyItWorks: `Each step matches a finding that holds up across many learning studies.

Worked examples first: beginners learn a new method faster by studying a fully solved example than by struggling alone, because all their attention goes to the method instead of to searching blindly (Sweller; Atkinson et al.). Then the support fades: the completion problem leaves only the key lines blank, the practice ladder gives you a full problem, and within each problem the hints go from a question to a plan to a code skeleton (Renkl and Atkinson). Labelled subgoals ("Find the wasted work", "Look up before you store") give you a plan you can reuse on new problems (Catrambone; Margulieux and Guzdial).

Answering check questions from memory strengthens memory far more than re-reading (retrieval practice, Roediger and Karpicke). Coming back to a solved problem days later in the Review queue, at growing gaps of 1, 3, 7, 16 and 35 days, beats doing it many times in one sitting (spacing, Cepeda et al.). Mixed Practice hides the category, so you have to recognise which pattern fits, which is the skill a real interview or job needs (interleaving, Rohrer and Taylor). Trying a problem for real before looking at hints, and explaining each line to yourself, makes the solution stick when you do see it (productive struggle, Kapur; self-explanation, Chi et al.).`,
				Intuition: "Think of learning to cook. First you watch a chef cook a dish while explaining each step (worked example). Then you cook it with half the prep already done (completion problem). Then you write the recipe down from memory (check questions). Then you cook a series of dishes that each add one new technique (practice ladder). Next week you cook last week's dish again without the book (Review queue). Finally someone hands you a mystery box and you must decide what to make (Mixed Practice).",
				CoreIdeas: []string{
					"Read the chapter, then study the worked example step by step, saying why each step is there before you read the explanation.",
					"Fill in the completion problem: the code is there with the key lines blank, so you practise exactly the part that matters.",
					"Answer the check questions from memory before revealing the answers; a wrong guess followed by the answer still helps you learn.",
					"Climb the practice ladder in order: problem 1 uses the pattern in its plainest form and each later one adds one idea.",
					"Try each problem for real before opening a hint. Hints fade: 1 is a question, 2 the key observation, 3 a step-by-step plan, 4 a code skeleton. Take only the next one you need.",
					"The model solution unlocks after you submit once or use all hints. Read it even when you passed, and answer its self-check questions in your own words.",
					"Come back to the Review queue when it has problems due; rate honestly how hard each one was so the next gap fits.",
					"Use Mixed Practice to learn to name the pattern before you code: the category is hidden, just as in an interview.",
				},
				CommonMistakes: []string{
					"Opening the solution after two minutes of being stuck. Being stuck for a while is where learning happens; take one hint, think again, and only then take the next.",
					"Re-reading a chapter instead of recalling it. Close the page and answer the check questions; re-reading feels productive but builds far weaker memory.",
					"Skipping the Review queue because the problem \"was easy last week\". Spaced repetition only works if you actually do the reviews when they are due.",
					"Copying code from the worked example without explaining it. Say out loud (or write as a comment) why each line is there; that self-explanation is what transfers to new problems.",
					"Practising one category for weeks and never mixing. You'll recognise patterns only when you're told the category; use Mixed Practice once two or more categories are unlocked.",
				},
				Examples: []CodeExample{
					{
						ID:       "self-explanation",
						Title:    "Self-explanation: write why next to each line",
						Language: "python",
						Snippet: `def has_negative(nums):
    for x in nums:        # look at each value once: at most n steps
        if x < 0:         # one negative is enough to answer
            return True   # stop early: best case is 1 step
    return False          # reached only after checking all n values`,
						Notes: "Before running code, explain each line to yourself: what it does and why it is needed. This habit is one of the best-supported ways to learn from examples.",
					},
				},
			},
			{
				Name:        "The problem-solving recipe",
				Description: "Five steps used in every chapter: understand with examples, brute force first, find the wasted work, optimise, test edge cases.",
				WhyItWorks: `Most algorithm problems are not solved by a flash of insight but by improving a slow solution. Understanding with examples catches misread requirements before you write code. A brute force gives you a correct baseline: it proves you understand the problem, and its answers let you check the fast version. Finding the wasted work is where the speed-up comes from: every pattern in this course (hashing, two pointers, sliding window, binary search, dynamic programming) removes a specific kind of repeated work, so asking "what does the brute force recompute or re-scan?" points to the right pattern. Optimising then means replacing that waste with a data structure or an argument. Edge cases (empty input, one element, duplicates, negatives, extremes) are where correct-looking code breaks, so you test them on purpose.

The five labels are subgoals. Using the same labels in every worked example, hint and walkthrough gives you a plan to fall back on when you face a problem you've never seen.`,
				Intuition: "A doctor doesn't jump straight to surgery. They listen to the symptoms (understand with examples), start with the standard safe treatment (brute force), notice what isn't working (find the wasted work), choose a targeted treatment (optimise), and check for side effects (test edge cases). Skipping the first step is how wrong problems get solved; skipping the last is how bugs ship.",
				CoreIdeas: []string{
					"Understand with examples: work two or three small inputs by hand, including one you invent yourself, and write the expected outputs.",
					"Brute force first: write the most direct correct solution and state its Big O. It doesn't need to be fast.",
					"Find the wasted work: look for work the brute force repeats, such as re-scanning the same items or recomputing the same value.",
					"Optimise: remove that waste with a structure (set, dict, sorted order, pointers, a table) and state the new Big O.",
					"Test edge cases: empty input, one element, all equal, duplicates, negatives, and the largest allowed values.",
				},
				CommonMistakes: []string{
					"Coding before working an example by hand. Write at least one input and its expected output first; many wrong answers come from misreading the task.",
					"Skipping the brute force because it's \"too slow\". Write it (or at least describe it): the fast solution is usually a fix to its waste, and it doubles as a checker.",
					"Testing only the examples in the statement. Add the edge cases yourself; hidden tests always include them.",
				},
				Examples: []CodeExample{
					{
						ID:       "smallest-gap-brute",
						Title:    "Brute force first: smallest gap between two numbers",
						Language: "python",
						Snippet: `def smallest_gap(nums):
    best = None
    for i in range(len(nums)):
        for j in range(i + 1, len(nums)):   # every pair once
            gap = abs(nums[i] - nums[j])
            if best is None or gap < best:
                best = gap
    return best`,
						Notes: "Correct and easy to trust, with n(n-1)/2 comparisons: O(n^2). The worked example below finds the wasted work and fixes it.",
					},
				},
			},
			{
				Name:        "Big O: counting how work grows",
				Description: "Describe running time by how the number of basic operations grows with the input size n, ignoring constant factors and smaller terms.",
				WhyItWorks: `Exact running time depends on the machine, the language and the input, so we count basic operations (a comparison, an addition, a lookup) as a function of n instead. Then we keep only the part that dominates for large n. For f(n) = 3n^2 + 5n + 2 at n = 1,000, the 3n^2 term is 3,000,000 while 5n + 2 is 5,002, under 0.2% of the total, and the gap widens as n grows. Constant factors like the 3 change with hardware, but the shape (squared) does not: doubling n always makes an n^2 algorithm about 4 times slower and an n algorithm about 2 times slower. So we write O(n^2).

Rules of thumb: steps one after another add (O(n) + O(n) = O(n)); a loop inside a loop multiplies (n times n is O(n^2)); halving the remaining work each step gives O(log n), because you can halve n only about log2 n times (20 times for a million). Unless a problem says otherwise, Big O describes the worst case.`,
				Intuition: "A taxi fare is a fixed starting fee plus a price per kilometre. For a long trip the starting fee hardly matters; what matters is how the price grows per kilometre. Big O is that per-kilometre shape: the fixed fee is a dropped constant, and a fare that grew with the square of the distance would lose to any per-kilometre fare on a long enough trip, no matter how cheap its start.",
				CoreIdeas: []string{
					"Count basic operations as a function of the input size n; name what n is (length of the list, number of nodes, and so on).",
					"Drop constant factors: O(3n) is O(n) and O(n/2) is O(n).",
					"Keep only the dominant term: O(n^2 + n) is O(n^2), and O(n + log n) is O(n).",
					"Sequential steps add; nested loops multiply (when the inner loop also grows with n).",
					"Repeated halving is O(log n); sorting with a good algorithm is O(n log n).",
					"Common classes from fast to slow: O(1), O(log n), O(n), O(n log n), O(n^2), O(2^n).",
				},
				CommonMistakes: []string{
					"Calling every nested loop O(n^2). If the inner loop runs over a fixed 26 letters, the total is 26n, which is O(n). Check what each loop actually runs over.",
					"Treating built-in operations as free. In Python, x in some_list, list.insert(0, x), slicing and sorted() all cost O(n) or O(n log n). Count them.",
					"Writing constants inside Big O, such as O(2n) or O(26). Drop them: O(n), O(1).",
					"Forgetting that two different inputs need two variables: looping over a list of size n inside a loop over a list of size m is O(n * m), not O(n^2).",
				},
				Examples: []CodeExample{
					{
						ID:       "count-operations",
						Title:    "Counting operations",
						Language: "python",
						Snippet: `def summary(nums):
    total = 0
    for x in nums:              # n iterations, O(1) each: O(n)
        total += x
    pairs = 0
    for i in range(len(nums)):  # n iterations ...
        for j in range(len(nums)):   # ... times n: O(n^2)
            if nums[i] < nums[j]:
                pairs += 1
    return total, pairs         # O(n) + O(n^2) = O(n^2)`,
						Notes: "Add the costs of the parts that run one after another, then keep only the dominant term.",
					},
				},
			},
			{
				Name:        "Space complexity",
				Description: "How the extra memory an algorithm uses grows with n, measured the same way as time.",
				WhyItWorks: `Memory is a resource just like time, and many problems limit it ("use O(1) extra space"). We count the extra memory the algorithm creates as n grows, usually not counting the input itself or the required output. A handful of variables is O(1); a new list, set or dict that can hold every element is O(n); a table of all pairs is O(n^2). Recursion also uses memory: each pending call keeps a stack frame, so recursion n levels deep costs O(n) space even if no list is created.

Time and space often trade against each other. Storing what you've already seen in a set costs O(n) memory but can turn an O(n^2) search into O(n) time; this trade is the heart of the next chapter.`,
				Intuition: "Time is how many trips you make to the shelf; space is how big a desk you need. Sorting papers in place on your lap needs no desk (O(1) space). Laying every paper out on the desk so you can spot one at a glance needs a desk as big as the pile (O(n) space), but finding things becomes instant.",
				CoreIdeas: []string{
					"Count extra memory created by the algorithm, as a function of n.",
					"A fixed number of variables is O(1) space, however large the numbers in them are.",
					"A copy, set, dict or list that grows with the input is O(n) space.",
					"Recursion depth counts: n nested calls use O(n) stack space.",
					"Trading memory for time is common and often worth it, unless the problem forbids it.",
				},
				CommonMistakes: []string{
					"Calling an algorithm O(1) space while it builds a sorted copy with sorted(nums) or a slice nums[1:]. Each creates an O(n) list.",
					"Counting the required output as extra space and then rejecting a correct O(1)-extra-space solution. State what you count.",
					"Ignoring recursion depth: a recursive function that recurses once per element uses O(n) stack memory.",
				},
				Examples: []CodeExample{
					{
						ID:       "reverse-two-ways",
						Title:    "O(n) versus O(1) extra space",
						Language: "python",
						Snippet: `def reversed_copy(nums):         # O(n) extra: builds a new list
    return [nums[i] for i in range(len(nums) - 1, -1, -1)]

def reverse_in_place(nums):      # O(1) extra: only two indices
    left, right = 0, len(nums) - 1
    while left < right:
        nums[left], nums[right] = nums[right], nums[left]
        left += 1
        right -= 1
    return nums`,
						Notes: "Both are O(n) time. The second uses only a constant number of variables, which matters when the problem forbids extra memory.",
					},
				},
			},
			{
				Name:        "Selection sort: a first O(n^2) algorithm",
				Description: "Repeatedly find the smallest remaining value and swap it to the front of the unsorted part.",
				WhyItWorks: `Invariant: after pass i, the first i + 1 positions hold the i + 1 smallest values in sorted order, and they never move again. Each pass scans the unsorted part to find its minimum and swaps it into the next position, so the sorted prefix grows by one and the invariant holds. After n - 1 passes only the largest value is left, and it is already in the last position.

Counting the work: pass 0 compares n - 1 pairs, pass 1 compares n - 2, and so on down to 1. The total is (n - 1) + (n - 2) + ... + 1 = n(n - 1)/2 comparisons, which is n^2/2 - n/2, so O(n^2). This count is the same for every input, even an already sorted one, so the best case is also O(n^2). It makes at most n - 1 swaps and sorts in place using O(1) extra space.`,
				Intuition: "Picking a team by height: look along the whole line for the shortest person and bring them to the front. Then look along everyone except the first person for the next shortest and bring them to second place. Each search covers one fewer person, but you still walk most of the line every time, which is why the total work grows with the square of the line's length.",
				CoreIdeas: []string{
					"The list is split into a sorted prefix and an unsorted rest; each pass moves the rest's minimum to the end of the prefix.",
					"Comparisons: (n-1) + (n-2) + ... + 1 = n(n-1)/2, so O(n^2) time in every case.",
					"Swaps: at most n - 1, which is why selection sort is sometimes used when writes are expensive.",
					"In place: only a few index variables, so O(1) extra space.",
				},
				CommonMistakes: []string{
					"Starting the inner scan at index 0 instead of i + 1, which re-scans the sorted prefix and can move settled values. Scan only the unsorted part.",
					"Swapping every time a smaller value is found instead of once per pass. Track the index of the minimum, then do a single swap after the scan.",
					"Thinking a sorted input makes it O(n). The inner loop still scans the whole unsorted part on every pass.",
				},
				Examples: []CodeExample{
					{
						ID:       "selection-sort",
						Title:    "Selection sort",
						Language: "python",
						Snippet: `def selection_sort(nums):
    n = len(nums)
    for i in range(n - 1):
        smallest = i
        for j in range(i + 1, n):      # scan the unsorted part
            if nums[j] < nums[smallest]:
                smallest = j
        nums[i], nums[smallest] = nums[smallest], nums[i]
    return nums`,
						Notes: "Two nested loops over shrinking ranges: n(n-1)/2 comparisons, O(n^2) time, O(1) extra space.",
					},
				},
			},
		},
		Animations: []Storyboard{
			{
				ID:    "runtime-shapes",
				Title: "Runtime Shapes",
				Goal:  "See how O(1), O(log n), O(n), O(n log n) and O(n^2) grow as n increases, and why the shape matters more than constants.",
				MemoryTips: []string{
					"Find the input size: say what n is before anything else.",
					"Count the work per loop: one pass over n is O(n); a loop over n inside a loop over n is O(n^2); halving each step is O(log n).",
					"Add it up: steps in sequence add, nested steps multiply.",
					"Keep the dominant term: drop constants and smaller terms (3n^2 + 5n becomes O(n^2)).",
					"Measure the space: count the extra lists, sets and recursion depth the same way.",
					"Check the constraints: n = 10^5 needs O(n log n) or better; O(n^2) would be 10^10 steps.",
				},
				Steps: []StoryboardStep{
					{Cue: "intro-bigo", Narration: "Big O describes how the number of steps grows as the input grows. It ignores constant factors, because those depend on the machine; the growth shape does not.", VisualHint: "Asymptotic intro panel", Duration: Duration(3 * time.Second), CodeRef: "concepts: asymptotic notation"},
					{Cue: "plot-axes", Narration: "The horizontal axis is the input size n; the vertical axis is the number of operations. Each curve is one complexity class.", VisualHint: "Axes: n vs operations", Duration: Duration(2 * time.Second), CodeRef: "runtime_shapes: setup plot"},
					{Cue: "draw-constant", Narration: "O(1) is a flat line: reading nums[i] or checking a set takes the same time for 10 items or 10 million.", VisualHint: "Flat green line", Duration: Duration(3 * time.Second), CodeRef: "arr[index] // O(1)"},
					{Cue: "draw-log", Narration: "O(log n) barely rises: halving a sorted list a million long reaches one item in about 20 steps, a billion in about 30.", VisualHint: "Gentle blue curve", Duration: Duration(3 * time.Second), CodeRef: "while (lo <= hi) mid = (lo+hi)/2"},
					{Cue: "draw-linear", Narration: "O(n) is a straight line: one pass over the input. Double n and the work doubles.", VisualHint: "Orange diagonal line", Duration: Duration(3 * time.Second), CodeRef: "for item in items: process(item)"},
					{Cue: "draw-nlogn", Narration: "O(n log n) sits just above linear. Good sorting algorithms live here: for n = 100,000 it is about 1.7 million steps.", VisualHint: "Purple curve above linear", Duration: Duration(3 * time.Second), CodeRef: "merge_sort: split, sort, merge"},
					{Cue: "draw-quadratic", Narration: "O(n^2) shoots up: a loop inside a loop, like selection sort or checking every pair. For n = 100,000 that is 10 billion steps.", VisualHint: "Red steep curve", Duration: Duration(3 * time.Second), CodeRef: "for i: for j: compare(i,j)"},
					{Cue: "scaling-demo", Narration: "Drag n and watch the table. At n = 100, n^2 is 10,000; at n = 1,000 it is a million. Ten times the input, a hundred times the work.", VisualHint: "Operations table updates", Duration: Duration(4 * time.Second), CodeRef: "scaling: watch numbers grow"},
					{Cue: "time-space", Narration: "Space grows by the same rules. A set of seen values costs O(n) memory but can turn an O(n^2) search into O(n) time: the trade you'll make in Chapter 2.", VisualHint: "Time-space reference table", Duration: Duration(3 * time.Second), CodeRef: "tradeoffs: memoization, caching"},
					{Cue: "compare", Narration: "For large inputs the algorithm matters more than the hardware: an O(n log n) program on a phone beats an O(n^2) program on a supercomputer once n is big enough.", VisualHint: "Large n comparison", Duration: Duration(3 * time.Second), CodeRef: "insight: choose wisely"},
				},
			},
			{
				ID:    "selection-sort",
				Title: "Selection Sort Mechanics",
				Goal:  "Watch the sorted prefix grow by one per pass, and count the comparisons that make selection sort O(n^2).",
				MemoryTips: []string{
					"Split the list: a sorted prefix on the left, the unsorted rest on the right (the prefix starts empty).",
					"Scan the unsorted part for the index of its smallest value.",
					"Swap that value into the first unsorted position, once per pass.",
					"Grow the prefix by one and repeat until one value is left.",
					"Count the work: (n-1) + (n-2) + ... + 1 = n(n-1)/2 comparisons, O(n^2) time, O(1) extra space.",
				},
				Steps: []StoryboardStep{
					{Cue: "scan", Narration: "Scan the unsorted part from left to right, comparing each value with the smallest found so far. On the first pass that is all n values: n - 1 comparisons.", VisualHint: "Sweeping pointer", Duration: Duration(3 * time.Second), CodeRef: "selection_sort: inner scan"},
					{Cue: "select", Narration: "Remember the index of the smallest value, not the value itself, so you know where to swap from.", VisualHint: "Highlight min", Duration: Duration(2 * time.Second), CodeRef: "selection_sort: track min index"},
					{Cue: "swap", Narration: "Swap the smallest value into the first unsorted position. It is now in its final place and will never move again.", VisualHint: "Animate swap", Duration: Duration(3 * time.Second), CodeRef: "selection_sort: swap"},
					{Cue: "shrink", Narration: "The sorted prefix grows by one, so the next scan is one shorter. Adding up every scan gives n(n-1)/2 comparisons: O(n^2), even if the list started sorted.", VisualHint: "Move boundary", Duration: Duration(2 * time.Second), CodeRef: "selection_sort: outer loop"},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "The recipe in action: smallest gap between two numbers",
				Problem: "Given a list nums of at least two integers, return the smallest absolute difference between two values at different positions. For example, nums = [8, 1, 5, 12] gives 3 (between 5 and 8).",
				Steps: []WorkedStep{
					{
						Subgoal:     "Understand with examples",
						Explanation: "Work small cases by hand before coding. List every pair of [8, 1, 5, 12] and its difference. Then invent edge cases: a repeated value means the answer is 0, and negative numbers must work too.",
						State:       "|8-1|=7 |8-5|=3 |8-12|=4 |1-5|=4 |1-12|=11 |5-12|=7 -> answer 3; [4, 9, 4] -> 0; [-3, 7, -1] -> 2",
					},
					{
						Subgoal:     "Brute force first",
						Explanation: "Compare every pair once (j starts after i). It is clearly correct, and its answers can check the faster version later. Count it: n(n-1)/2 comparisons, so O(n^2) time and O(1) extra space.",
						Code: `best = None
for i in range(len(nums)):
    for j in range(i + 1, len(nums)):
        gap = abs(nums[i] - nums[j])
        if best is None or gap < best:
            best = gap`,
						State: "n = 4 -> 6 comparisons; n = 100,000 -> about 5,000,000,000",
					},
					{
						Subgoal:     "Find the wasted work",
						Explanation: "Picture the values in sorted order: 1, 5, 8, 12. Comparing 1 with 8 can never win, because 5 sits between them and is closer to both. In sorted order the closest partner of any value is one of its neighbours, so every non-neighbour comparison is wasted.",
						State:       "sorted: 1 5 8 12; useful pairs: (1,5) (5,8) (8,12); wasted: (1,8) (1,12) (5,12)",
					},
					{
						Subgoal:     "Optimise",
						Explanation: "Sort once, then compare only neighbours. Sorting is O(n log n) and the neighbour pass is O(n); the sum is O(n log n). sorted() builds a new list, so extra space is O(n).",
						Code: `ordered = sorted(nums)
best = ordered[1] - ordered[0]
for i in range(2, len(ordered)):
    best = min(best, ordered[i] - ordered[i - 1])`,
						State: "ordered = [1, 5, 8, 12]; gaps 4, 3, 4; best = 3",
					},
					{
						Subgoal:     "Test edge cases",
						Explanation: "Run the cases you invented in step 1 and compare with the brute force. With two elements the loop doesn't run and best is their gap; a repeat gives a neighbour gap of 0; negatives sort correctly. Because ordered is sorted, each gap is already non-negative, so abs is not needed.",
						State:       "[3, 10] -> 7; [4, 9, 4] -> ordered [4, 4, 9] -> 0; [-3, 7, -1] -> ordered [-3, -1, 7] -> 2",
					},
				},
				Code: `def smallest_gap(nums):
    ordered = sorted(nums)
    best = ordered[1] - ordered[0]
    for i in range(2, len(ordered)):
        best = min(best, ordered[i] - ordered[i - 1])
    return best`,
			},
			{
				Title:   "Counting the operations in selection sort",
				Problem: "Find the time and space complexity of selection sort (code below) and trace it on nums = [29, 10, 14, 37, 13].",
				Steps: []WorkedStep{
					{
						Subgoal:     "Find the input size",
						Explanation: "The work depends on how many values there are, so n = len(nums). The values themselves don't change the number of comparisons.",
						Code:        `n = len(nums)`,
						State:       "nums = [29, 10, 14, 37, 13], n = 5",
					},
					{
						Subgoal:     "Count the work per loop",
						Explanation: "The outer loop runs n - 1 times (i = 0 .. n - 2). On pass i the inner loop compares positions i + 1 .. n - 1, which is n - 1 - i comparisons. Each comparison and the one swap per pass are O(1).",
						Code: `for i in range(n - 1):
    smallest = i
    for j in range(i + 1, n):
        if nums[j] < nums[smallest]:
            smallest = j
    nums[i], nums[smallest] = nums[smallest], nums[i]`,
						State: "pass 0: 4 comparisons, min 10 -> [10, 29, 14, 37, 13]; pass 1: 3, min 13 -> [10, 13, 14, 37, 29]; pass 2: 2, min 14 stays; pass 3: 1, min 29 -> [10, 13, 14, 29, 37]",
					},
					{
						Subgoal:     "Add it up",
						Explanation: "The passes run one after another, so add their costs: (n - 1) + (n - 2) + ... + 1 = n(n - 1)/2. For n = 5 that is 4 + 3 + 2 + 1 = 10 comparisons, plus at most n - 1 = 4 swaps.",
						State:       "total comparisons = 10 = 5 * 4 / 2",
					},
					{
						Subgoal:     "Keep the dominant term",
						Explanation: "n(n - 1)/2 = n^2/2 - n/2. Drop the constant 1/2 and the smaller term n/2: the time is O(n^2). The count doesn't depend on the order of the input, so a sorted list costs the same: best, average and worst case are all O(n^2).",
						State:       "n = 1,000 -> 499,500 comparisons; n = 100,000 -> about 5 billion",
					},
					{
						Subgoal:     "Measure the space",
						Explanation: "Only n, i, j and smallest are created, and the swaps happen inside nums. The extra memory doesn't grow with n, so space is O(1).",
						State:       "extra variables: n, i, j, smallest -> O(1)",
					},
				},
				Code: `def selection_sort(nums):
    n = len(nums)
    for i in range(n - 1):
        smallest = i
        for j in range(i + 1, n):
            if nums[j] < nums[smallest]:
                smallest = j
        nums[i], nums[smallest] = nums[smallest], nums[i]
    return nums`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Selection sort: fill in the scan and the swap",
				Prompt: "Complete selection sort. The inner loop must scan only the unsorted part, keep the index of the smallest value, and the pass ends with a single swap. Check yourself on [4, 3, 1, 2] -> [1, 2, 3, 4].",
				Code: `def selection_sort(nums):
    n = len(nums)
    for i in range(n - 1):
        smallest = i
        for j in range(____, n):
            if nums[j] < nums[____]:
                smallest = j
        nums[i], nums[smallest] = ____
    return nums`,
				Blanks:      []string{"i + 1", "smallest", "nums[smallest], nums[i]"},
				Explanation: "The scan starts at i + 1 because positions 0..i-1 are already sorted and position i is the current candidate. Comparing against nums[smallest] (not nums[i]) keeps the minimum of everything scanned so far. The swap puts that minimum at position i; Python's tuple assignment swaps without a temporary variable.",
			},
			{
				Title:  "Largest gap between sorted neighbours",
				Prompt: "Use the same recipe as the smallest-gap example. Given at least two numbers, return the largest difference between two values that are next to each other once the list is sorted. Example: [3, 9, 1, 6] sorts to [1, 3, 6, 9], whose neighbour gaps are 2, 3 and 3, so the answer is 3. Target: O(n log n).",
				Code: `def largest_sorted_gap(nums):
    ordered = ____
    best = 0
    for i in range(1, len(ordered)):
        best = max(best, ____)
    return best`,
				Blanks:      []string{"sorted(nums)", "ordered[i] - ordered[i - 1]"},
				Explanation: "Optimise: sorting makes neighbours the only pairs that count, exactly as in the worked example. Each gap is ordered[i] - ordered[i - 1], already non-negative because the list is sorted. Sorting is O(n log n) and the pass O(n), so the total is O(n log n).",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt:      "What is the time complexity of this code?\n\ntotal = 0\nfor i in range(n):\n    for j in range(i, n):\n        total += 1",
				Choices:     []string{"O(n)", "O(n log n)", "O(n^2)", "O(n^2 / 2), which is different from O(n^2)"},
				Answer:      2,
				Explanation: "The inner loop runs n, then n - 1, ..., then 1 times: n(n + 1)/2 steps in total. Dropping the constant 1/2 and the smaller term gives O(n^2). O(n^2 / 2) is the same class as O(n^2), because Big O ignores constant factors.",
			},
			{
				Prompt:      "An algorithm does 3n + 100 log n + 7 operations. What is its Big O?",
				Choices:     []string{"O(3n)", "O(n)", "O(log n)", "O(n log n)"},
				Answer:      1,
				Explanation: "Keep the dominant term (3n grows faster than 100 log n for large n) and drop its constant: O(n). O(3n) is not wrong in meaning but constants aren't written inside Big O. O(n log n) would be a product, not a sum.",
			},
			{
				Prompt:      "Trace selection sort on [4, 3, 1, 2]. What is the list after the first two passes of the outer loop?",
				Choices:     []string{"[1, 2, 4, 3]", "[1, 3, 4, 2]", "[1, 2, 3, 4]", "[3, 4, 1, 2]"},
				Answer:      0,
				Explanation: "Pass 1: the smallest value is 1 at index 2; swap with index 0 -> [1, 3, 4, 2]. Pass 2: the smallest of [3, 4, 2] is 2 at index 3; swap with index 1 -> [1, 2, 4, 3]. A third pass would finish the sort.",
			},
			{
				Prompt:      "A problem says n can be as large as 100,000. Which approach is fast enough?",
				Choices:     []string{"Check every pair: O(n^2)", "Sort, then make one pass: O(n log n)", "Try every subset: O(2^n)", "They are all fine on a modern computer"},
				Answer:      1,
				Explanation: "O(n log n) is about 1.7 million steps for n = 100,000. O(n^2) is 10 billion steps, which takes minutes in Python, and O(2^n) would never finish. Reading the constraints this way tells you which brute force must be improved.",
			},
			{
				Prompt:      "You've read a chapter and opened the first practice problem. After a few minutes you are stuck. What is the best next step?",
				Choices:     []string{"Open the model solution to see how it's done", "Reveal hint 1 (a question), think about it, and try again before taking the next hint", "Skip to the next problem in the ladder", "Re-read the whole chapter from the top"},
				Answer:      1,
				Explanation: "Struggling a bit first makes the idea stick when it arrives. Hints fade from a question to the key observation, a plan and a skeleton, so take only the next one you need. The ladder is ordered so each problem prepares the next, and re-reading is much weaker than trying to recall.",
			},
			{
				Prompt:      "Name the five steps of the problem-solving recipe and say what each one gives you.",
				Explanation: "1. Understand with examples: confirms you've read the task right and gives test cases. 2. Brute force first: a correct baseline and a checker, with its Big O. 3. Find the wasted work: the repeated scanning or recomputing the brute force does, which points to the pattern to use. 4. Optimise: remove that waste with a better structure or argument, and state the new Big O. 5. Test edge cases: empty, single, duplicates, negatives and extremes, where code usually breaks.",
			},
			{
				Prompt:      "Why does the app bring solved problems back in the Review queue days later, and why does Mixed Practice hide each problem's category?",
				Explanation: "Recalling something after a gap (1, 3, 7, 16, 35 days) strengthens memory much more than repeating it in one sitting, which only feels effective: that is spaced practice. Mixed Practice hides the category because in a real interview or job nobody tells you \"this is a two-pointer problem\"; mixing categories forces you to recognise which pattern fits, a skill that blocked practice by category never trains.",
			},
			{
				Prompt:      "How much extra space does selection sort use?",
				Choices:     []string{"O(1)", "O(n)", "O(n^2)", "O(log n)"},
				Answer:      0,
				Explanation: "It sorts inside the given list and only creates a few index variables (n, i, j, smallest), whose number doesn't grow with n. Building a sorted copy would be O(n), but selection sort doesn't.",
			},
		},
	}
}
