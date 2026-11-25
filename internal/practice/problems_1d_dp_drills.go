package practice

// OneDDPDrills are the 1-D dynamic programming foundations: name the state
// ("what does table[i] mean?"), write the recurrence from the last step, fill
// the table from the base cases up, and shrink it to a few variables. The
// ladder runs Fibonacci (memo, table, two variables), a three-term recurrence,
// a min over two choices, a recurrence with a coefficient on a big n where
// recursion runs out of stack, best-ending-here (Kadane), and the loop order
// that separates combinations from orderings. The cases with large inputs
// carry a 5 second limit so that plain recursion (exponential) or recomputing
// every subarray (quadratic) fails.
var OneDDPDrills = []*Problem{
	{
		ID:              "drill-fibonacci",
		Number:          509,
		Kind:            KindDrill,
		Order:           1,
		Title:           "Fibonacci Number",
		Slug:            "drill-fibonacci",
		Difficulty:      "Easy",
		Category:        "1d-dp",
		Tags:            []string{"Dynamic Programming", "Memoization", "Math", "Foundations"},
		RelatedChapters: []int{15, 8},
		Description:     "The Fibonacci numbers start `fib(0) = 0`, `fib(1) = 1`, and every later one is the sum of the two before it: `fib(n) = fib(n - 1) + fib(n - 2)`. The sequence runs 0, 1, 1, 2, 3, 5, 8, 13, ... Given `n`, return `fib(n)`.\n\nThe recurrence is already a recursive function, but written that way it recomputes the same values again and again: `fib(5)` calls `fib(3)` twice and `fib(2)` three times, and the number of calls roughly doubles with every step of `n`. Make each value cost one step:\n\n1. **Memoise**: keep the recursion but remember every answer in a dictionary.\n2. **Table**: the **state** is `table[i] = fib(i)`; fill it from the **base cases** `table[0] = 0` and `table[1] = 1` upwards.\n3. **Two variables**: `table[i]` only reads the two entries before it, so keep just those.\n\nAny of the three is accepted.",
		Constraints:     []string{"0 <= n <= 90"},
		Examples: []Example{
			{Input: "n = 2", Output: "1", Explanation: "fib(2) = fib(1) + fib(0) = 1 + 0 = 1."},
			{Input: "n = 6", Output: "8", Explanation: "The table runs 0, 1, 1, 2, 3, 5, 8, so fib(6) = 8."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 2}, Expected: 1, Note: "example 1: the first value built from the recurrence"},
			{Input: map[string]interface{}{"n": 6}, Expected: 8, Note: "example 2: a short run of the table"},
			{Input: map[string]interface{}{"n": 0}, Expected: 0, Hidden: true, Note: "n = 0: a table of size n + 1 has no index 1, so n = 0 must be handled before writing table[1]"},
			{Input: map[string]interface{}{"n": 1}, Expected: 1, Hidden: true, Note: "n = 1: the second base case"},
			{Input: map[string]interface{}{"n": 10}, Expected: 55, Hidden: true, Note: "n = 10: base cases of 1 and 1 give 89 instead of 55"},
			{Input: map[string]interface{}{"n": 20}, Expected: 6765, Hidden: true, Note: "n = 20"},
			{Input: map[string]interface{}{"n": 45}, Expected: 1134903170, Hidden: true, Timeout: 5000, Note: "n = 45: plain recursion makes billions of calls and runs out of time"},
			{Input: map[string]interface{}{"n": 90}, Expected: odFib(90), Hidden: true, Timeout: 5000, Note: "n = 90, the largest input: only a version that computes each value once finishes, and the answer is still a large whole number"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def fib(n):\n    # Return the n-th Fibonacci number\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "fib(6) needs fib(5) and fib(4), and fib(5) needs fib(4) as well. How many times is fib(4) worked out, and what could you do the first time so the other times are free?"},
			{Level: 2, Type: "insight", Content: "The same subproblem fib(i) is asked for many times, but its answer never changes. Compute each fib(i) once and store it: top-down that is a memo dictionary, bottom-up it is a table where table[i] = fib(i), filled from the base cases table[0] = 0 and table[1] = 1. That is dynamic programming; see Chapter 15 (and Chapter 8 for the recursion it starts from).", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. State: table[i] holds fib(i).\n2. Base cases: handle n below 2 directly, and set table[1] = 1.\n3. Recurrence: for i from 2 to n, table[i] = table[i - 1] + table[i - 2].\n4. Answer: return table[n]."},
			{Level: 4, Type: "skeleton", Content: "def fib(n):\n    if n < 2:\n        return n\n    table = [0] * (n + 1)\n    table[1] = 1\n    for i in range(2, n + 1):\n        table[i] = ____ + ____\n    return table[n]"},
		},
		Solution: Solution{
			Code:            "def fib(n):\n    if n < 2:\n        return n\n    table = [0] * (n + 1)\n    table[1] = 1\n    for i in range(2, n + 1):\n        table[i] = table[i - 1] + table[i - 2]\n    return table[n]",
			Explanation:     "This is the smallest complete dynamic program: a state (table[i] = fib(i)), base cases (0 and 1), a recurrence (the sum of the two before) and an order to fill it in (increasing i, so both inputs are ready). Every later drill and the whole 1-D DP ladder, starting with Climbing Stairs, are this shape with a different state and recurrence. The two alternatives show the same idea top-down (memo) and with the table squeezed to two variables.",
			BruteForce:      "Write the recurrence as a function: fib(n) = fib(n - 1) + fib(n - 2). It is correct but makes about 1.6^n calls, so n = 45 is billions of calls and n = 90 is out of reach. The waste is that fib(i) for small i is recomputed in every branch of the call tree.",
			KeyInsight:      "There are only n + 1 different subproblems, fib(0) to fib(n), and each one's answer never changes, so compute it once and look it up afterwards.",
			WhyItWorks:      "State: table[i] = fib(i). Invariant: when the loop reaches i, table[0..i - 1] already hold the correct values. It holds at i = 2 because table[0] = 0 and table[1] = 1 are the definition. The recurrence says fib(i) is exactly fib(i - 1) + fib(i - 2), both of which are filled in, so table[i] is correct and the invariant carries on. The recurrence is complete because the definition has no other case for i >= 2. The loop does n - 1 additions in all.",
			CommonMistakes:  []string{"Writing the recursion with no memo. It is correct on small inputs and then takes exponential time; store each answer, or build the table bottom-up.", "Setting the base cases wrongly, for example fib(0) = 1 or returning 1 for every n <= 1. The recurrence faithfully copies a wrong start, so every later answer is wrong; check n = 0, 1 and 10.", "Writing `table[1] = 1` without handling n = 0 first: the table `[0] * (n + 1)` has one cell when n = 0, so index 1 does not exist. Return n directly when n < 2, or make the table n + 2 long.", "Making the table one cell too short (`[0] * n`), so `table[n]` is out of range: a table that answers for 0..n needs n + 1 cells."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Handle the smallest inputs", Explanation: "fib(0) = 0 and fib(1) = 1, which is n itself, so answer these straight away. This also keeps the table below from being too small for index 1.", CodeSnippet: "if n < 2:\n    return n", LineStart: 2, LineEnd: 3},
				{Title: "Define the state and base cases", Explanation: "table[i] will hold fib(i) for every i from 0 to n, so it needs n + 1 cells. table[0] = 0 is already there; set table[1] = 1.", CodeSnippet: "table = [0] * (n + 1)\ntable[1] = 1", LineStart: 4, LineEnd: 5},
				{Title: "Apply the recurrence", Explanation: "Each cell is the sum of the two before it, which are already filled because i increases. For n = 6 the table becomes 0, 1, 1, 2, 3, 5, 8.", CodeSnippet: "for i in range(2, n + 1):\n    table[i] = table[i - 1] + table[i - 2]", LineStart: 6, LineEnd: 7},
				{Title: "Read the answer", Explanation: "The last cell is the answer to the question that was asked.", CodeSnippet: "return table[n]", LineStart: 8, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{Name: "Memoised recursion", Code: "def fib(n):\n    memo = {0: 0, 1: 1}\n\n    def solve(i):\n        if i not in memo:\n            memo[i] = solve(i - 1) + solve(i - 2)\n        return memo[i]\n\n    return solve(n)", TimeComplexity: "O(n)", SpaceComplexity: "O(n)", Explanation: "The recurrence as a recursive function, with a dictionary so that each fib(i) is computed once. It is the top-down way of filling the same table, and uses recursion depth n."},
				{Name: "Two variables", Code: "def fib(n):\n    previous, current = 0, 1\n    for _ in range(n):\n        previous, current = current, previous + current\n    return previous", TimeComplexity: "O(n)", SpaceComplexity: "O(1)", Explanation: "table[i] reads only table[i - 1] and table[i - 2], so keep those two as variables and slide them forward. Same time, constant space."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is memoised recursion fast when plain recursion is not, if they run the same recurrence?", Answer: "The recurrence only has n + 1 distinct subproblems, fib(0) to fib(n), but the plain call tree asks for them an exponential number of times in total. With a memo, the first call for each i does the work and every later call is one dictionary lookup, so the total work is one addition per subproblem."},
			{Prompt: "How would you know you could shrink the table to two variables, and what would stop you?", Answer: "Look at what table[i] reads: only table[i - 1] and table[i - 2], so older cells are never used again and can be dropped. If the recurrence read table[i - 10], or an arbitrary earlier cell, you would need to keep that many cells (or the whole table)."},
		},
	},
	{
		ID:              "drill-climb-one-two-three",
		Number:          0,
		Kind:            KindDrill,
		Order:           2,
		Title:           "Climb with One, Two or Three Steps",
		Slug:            "drill-climb-one-two-three",
		Difficulty:      "Easy",
		Category:        "1d-dp",
		Tags:            []string{"Dynamic Programming", "Counting", "Foundations"},
		RelatedChapters: []int{15, 8},
		Description:     "A staircase has `n` steps. You start on the ground (step 0) and each move climbs **1, 2 or 3 steps**. Return the number of different move sequences that take you from the ground to step `n` exactly. Two sequences are different if they differ in at least one move, so 1+2 and 2+1 count as two ways. The answer fits in a 64-bit integer.\n\nWrite the recurrence by asking about the **last move**. Let `ways[i]` be the number of ways to reach step `i`. Every way to reach step `i` ends with a move of 1, 2 or 3, which came from step `i - 1`, `i - 2` or `i - 3`, and different last moves give different sequences. So `ways[i] = ways[i - 1] + ways[i - 2] + ways[i - 3]`, counting only the steps that exist.\n\nThe base case is `ways[0] = 1`: there is exactly one way to be on the ground, the empty sequence. This is Climbing Stairs (the ladder's first problem) with one more allowed move.",
		Constraints:     []string{"1 <= n <= 60"},
		Examples: []Example{
			{Input: "n = 3", Output: "4", Explanation: "1+1+1, 1+2, 2+1 and 3 itself: four ways. By the recurrence, ways[3] = ways[2] + ways[1] + ways[0] = 2 + 1 + 1."},
			{Input: "n = 5", Output: "13", Explanation: "ways[4] = 7 and ways[3] = 4 and ways[2] = 2, so ways[5] = 7 + 4 + 2 = 13."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 3}, Expected: 4, Note: "example 1: the 3-step move appears for the first time"},
			{Input: map[string]interface{}{"n": 5}, Expected: 13, Note: "example 2: three terms are added"},
			{Input: map[string]interface{}{"n": 1}, Expected: 1, Hidden: true, Note: "n = 1, the smallest input: a table must not index ways[2] or ways[-1]"},
			{Input: map[string]interface{}{"n": 2}, Expected: 2, Hidden: true, Note: "n = 2: only the moves 1+1 and 2 exist (ways[-1] counts for nothing)"},
			{Input: map[string]interface{}{"n": 4}, Expected: 7, Hidden: true, Note: "n = 4: Fibonacci-style two-term sums give 5"},
			{Input: map[string]interface{}{"n": 10}, Expected: 274, Hidden: true, Note: "n = 10: a wrong base case drifts further from 274 with each step"},
			{Input: map[string]interface{}{"n": 25}, Expected: odClimb123(25), Hidden: true, Note: "n = 25"},
			{Input: map[string]interface{}{"n": 60}, Expected: odClimb123(60), Hidden: true, Timeout: 5000, Note: "n = 60: plain recursion branches three ways each step and cannot finish; the answer is above 10^15"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def waysToClimb(n):\n    # Return the number of move sequences of 1, 2 or 3 steps that reach step n\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Imagine you are standing on step i. What could your very last move have been, and which step were you on before it?"},
			{Level: 2, Type: "insight", Content: "Group the ways to reach step i by their last move. A last move of 1 came from i - 1, of 2 from i - 2, of 3 from i - 3, and no sequence ends two ways, so ways[i] is the sum of those three entries (skip any step below 0). The ground has exactly one way (do nothing), so ways[0] = 1 and the other base cases come out of the recurrence. See Chapter 15.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. State: ways[i] is the number of ways to reach step i, for i from 0 to n.\n2. Base case: ways[0] = 1.\n3. Recurrence: for each i from 1 to n, add ways[i - step] for each step in 1, 2, 3 whose i - step is not below 0.\n4. Answer: return ways[n]."},
			{Level: 4, Type: "skeleton", Content: "def waysToClimb(n):\n    ways = [0] * (n + 1)\n    ways[0] = ____\n    for i in range(1, n + 1):\n        for step in (1, 2, 3):\n            if ____ >= 0:\n                ways[i] += ways[i - step]\n    return ways[n]"},
		},
		Solution: Solution{
			Code:            "def waysToClimb(n):\n    ways = [0] * (n + 1)\n    ways[0] = 1\n    for i in range(1, n + 1):\n        for step in (1, 2, 3):\n            if i - step >= 0:\n                ways[i] += ways[i - step]\n    return ways[n]",
			Explanation:     "This builds on Fibonacci: the same table filled from the base cases up, but the recurrence now has three terms and you write it by asking what the last move was. Putting the single base case ways[0] = 1 in the table and guarding with i - step >= 0 means n = 1 and n = 2 need no special cases. Climbing Stairs, the first problem of the ladder, is this drill with the 3-step move removed.",
			BruteForce:      "Recurse: ways(i) = ways(i - 1) + ways(i - 2) + ways(i - 3) with ways(0) = 1 and ways(negative) = 0. That branches three ways per step and makes exponentially many calls (about 1.8^n), so n = 60 never finishes. The waste is that ways(i) for small i is recomputed in every branch.",
			KeyInsight:      "Classify every way to reach step i by its last move. The three classes are separate and complete, and each is counted by an earlier table entry, so the whole count is the sum of three cells.",
			WhyItWorks:      "State: ways[i] is the number of move sequences from the ground to step i. Every such sequence for i >= 1 ends with exactly one move of 1, 2 or 3 (and the move is only possible if the step it started on is at least 0). Removing that last move leaves a sequence that reaches step i - 1, i - 2 or i - 3, and any such sequence plus the move gives a different sequence for i. So the three classes are disjoint, they cover everything, and ways[i] is their sum. The base case ways[0] = 1 counts the empty sequence; ways of negative steps count nothing because the guard skips them.",
			CommonMistakes:  []string{"Adding only two previous entries (the Fibonacci recurrence from Climbing Stairs). The 3-step move is a third class of last move, and n = 3 already gives 3 instead of 4.", "Setting `ways[0] = 0`. Every value then stays 0: the ground has one way to be reached (the empty sequence), which is what seeds every other count.", "Making the table `[0] * n` (one cell short), so `ways[n]` fails; a table that answers for 0..n has n + 1 cells.", "Reading `ways[i - step]` without the `i - step >= 0` guard, which in Python wraps around to the end of the table and silently adds wrong values for n = 1 and n = 2.", "Writing the recursion without a memo: it passes the small cases and then times out at n = 60."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "ways[i] is the number of ways to stand on step i, for 0 to n, so the table has n + 1 cells that start at 0.", CodeSnippet: "ways = [0] * (n + 1)", LineStart: 2, LineEnd: 2},
				{Title: "Set the base case", Explanation: "There is one way to be on the ground: make no move. Every other count grows out of this 1.", CodeSnippet: "ways[0] = 1", LineStart: 3, LineEnd: 3},
				{Title: "Split by the last move", Explanation: "For step i, the last move was 1, 2 or 3, from step i - step. Add the count for each possible starting step, and skip starting steps below the ground.", CodeSnippet: "for step in (1, 2, 3):\n    if i - step >= 0:\n        ways[i] += ways[i - step]", LineStart: 5, LineEnd: 7},
				{Title: "Read the answer", Explanation: "For n = 5 the table is 1, 1, 2, 4, 7, 13, so the answer is its last cell.", CodeSnippet: "return ways[n]", LineStart: 8, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{Name: "Three variables", Code: "def waysToClimb(n):\n    two_back, one_back, current = 0, 0, 1\n    for _ in range(n):\n        two_back, one_back, current = one_back, current, two_back + one_back + current\n    return current", TimeComplexity: "O(n)", SpaceComplexity: "O(1)", Explanation: "ways[i] only reads the three entries before it, so keep those three as variables and slide them. Starting from (0, 0, 1), which stands for ways[-2], ways[-1] and ways[0], makes the first step behave like the table version."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it right to add ways[i - 1], ways[i - 2] and ways[i - 3] and not count some sequences twice?", Answer: "Each sequence has exactly one last move. Grouping by that last move splits the sequences into three groups with no overlap (a sequence cannot end in both a 1 and a 2), and every sequence is in one of them. Within a group, dropping the last move is a one-to-one match with the sequences that reach the earlier step, so each group has exactly ways[i - step] members."},
			{Prompt: "What changes if a move of 2 steps were not allowed, only 1 or 3?", Answer: "Only the set of steps in the recurrence: ways[i] = ways[i - 1] + ways[i - 3]. The state, the base case ways[0] = 1 and the fill order are identical. Writing the allowed steps as a tuple that the loop walks over is what makes this a one-line change."},
		},
	},
	{
		ID:              "drill-min-cost-climbing-stairs",
		Number:          746,
		Kind:            KindDrill,
		Order:           3,
		Title:           "Min Cost Climbing Stairs",
		Slug:            "drill-min-cost-climbing-stairs",
		Difficulty:      "Easy",
		Category:        "1d-dp",
		Tags:            []string{"Dynamic Programming", "Array", "Foundations"},
		RelatedChapters: []int{15, 8},
		Description:     "A staircase has steps numbered `0` to `len(cost) - 1`, and `cost[i]` is the price you pay **when you step off step `i`** (to climb 1 or 2 steps up from it). You may start on step 0 or on step 1 without paying anything yet. The **top** is the position `len(cost)`, one past the last step; you reach it when a move takes you there or beyond.\n\nReturn the minimum total price to reach the top.\n\nThe new idea compared with counting ways is that you **minimise** over the two choices for the last move. Let `best[i]` be the cheapest price to be standing on position `i`. You got to `i` from `i - 1` (paying `cost[i - 1]`) or from `i - 2` (paying `cost[i - 2]`), so `best[i] = min(best[i - 1] + cost[i - 1], best[i - 2] + cost[i - 2])`. Because you may start on step 0 or 1 for free, `best[0] = best[1] = 0`. The answer is `best[len(cost)]`.",
		Constraints:     []string{"2 <= len(cost) <= 1000", "0 <= cost[i] <= 999"},
		Examples: []Example{
			{Input: "cost = [10,15,20]", Output: "15", Explanation: "Start on step 1, pay 15, and climb 2 steps to the top. Total 15. Starting on step 0 would cost 10 + 20 = 30."},
			{Input: "cost = [1,100,1,1,1,100,1,1,100,1]", Output: "6", Explanation: "Start on step 0 and use the cheap steps: pay 1 at steps 0, 2, 4, 6, 7 and 9 and jump over the 100s, total 6."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"cost": []int{10, 15, 20}}, Expected: 15, Note: "example 1: the top can be reached from the last step or the one before it, and the one before is cheaper"},
			{Input: map[string]interface{}{"cost": []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, Expected: 6, Note: "example 2: jump over every expensive step"},
			{Input: map[string]interface{}{"cost": []int{5, 8}}, Expected: 5, Hidden: true, Note: "two steps: you can start on either and step off once"},
			{Input: map[string]interface{}{"cost": []int{9, 1}}, Expected: 1, Hidden: true, Note: "two steps: starting on step 1 beats starting on step 0"},
			{Input: map[string]interface{}{"cost": []int{50, 1, 1, 50}}, Expected: 2, Hidden: true, Note: "start on step 1: starting only on step 0 forces the 50"},
			{Input: map[string]interface{}{"cost": []int{0, 0, 0, 0}}, Expected: 0, Hidden: true, Note: "all free"},
			{Input: map[string]interface{}{"cost": []int{1, 2, 100, 3, 4, 100}}, Expected: 9, Hidden: true, Note: "the last step is expensive: finish by jumping from the step before it. Reading best[last step] instead of best[top] pays the 100"},
			{Input: map[string]interface{}{"cost": []int{2, 5, 1, 1, 9, 1, 1, 7}}, Expected: 6, Hidden: true, Note: "picking the cheaper next step greedily does not always win"},
			{Input: map[string]interface{}{"cost": odStairCosts(60)}, Expected: odMinCost(odStairCosts(60)), Hidden: true, Timeout: 5000, Note: "60 steps: plain recursion with two branches per step cannot finish"},
			{Input: map[string]interface{}{"cost": odStairCosts(1000)}, Expected: odMinCost(odStairCosts(1000)), Hidden: true, Note: "the largest input"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def minCostClimbingStairs(cost):\n    # Return the minimum price to reach the top (position len(cost))\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "To stand on position i, you came from one of two positions. What does each of them cost to leave, and which would you choose?"},
			{Level: 2, Type: "insight", Content: "Let best[i] be the cheapest price to be standing on position i (the top is position len(cost)). The last move came from i - 1 or i - 2, so best[i] = min(best[i - 1] + cost[i - 1], best[i - 2] + cost[i - 2]). Starting on step 0 or step 1 is free, so best[0] = best[1] = 0. The answer is at the top, not at the last step. See Chapter 15.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. State: best[i] is the cheapest price to stand on position i, for i from 0 to n (n is len(cost)).\n2. Base cases: best[0] = best[1] = 0, since you may start on either step.\n3. Recurrence: for i from 2 to n, take the smaller of best[i - 1] + cost[i - 1] and best[i - 2] + cost[i - 2].\n4. Answer: return best[n]."},
			{Level: 4, Type: "skeleton", Content: "def minCostClimbingStairs(cost):\n    n = len(cost)\n    best = [0] * (n + 1)\n    for i in range(2, n + 1):\n        best[i] = min(best[i - 1] + ____, best[i - 2] + ____)\n    return best[____]"},
		},
		Solution: Solution{
			Code:            "def minCostClimbingStairs(cost):\n    n = len(cost)\n    best = [0] * (n + 1)\n    for i in range(2, n + 1):\n        best[i] = min(best[i - 1] + cost[i - 1], best[i - 2] + cost[i - 2])\n    return best[n]",
			Explanation:     "This builds on Climb with One, Two or Three Steps: the same last-move reasoning and left-to-right table, but each choice now has a price and the recurrence takes the minimum of the choices instead of the sum. The base cases are the free starts, and the answer sits at position n, one past the last step. House Robber, the second problem of the ladder, uses the same min/max-over-two-choices shape.",
			BruteForce:      "Try both moves from every step recursively and keep the cheaper total: about 1.6^n calls, so 60 steps never finishes. The waste is that the cheapest way to finish from step i is worked out again every time some path reaches i.",
			KeyInsight:      "The cheapest way to reach a position only depends on the cheapest ways to reach the two positions before it, so a table of cheapest prices solves every position once, left to right.",
			WhyItWorks:      "State: best[i] is the minimum price to stand on position i, where leaving step j costs cost[j]. Positions 0 and 1 are free starts, so best[0] = best[1] = 0 (reaching position 1 from step 0 would cost cost[0], which is never less than 0). For i >= 2 the last move came from i - 1 or i - 2 and paid that step's price. The best way to reach the earlier position plus its price is the cheapest way through that choice, so the minimum of the two choices is the cheapest way to reach i. Positions are filled in increasing order, so both inputs are ready, and best[n] is the cheapest way to the top.",
			CommonMistakes:  []string{"Returning the table entry for the last step (index n - 1) instead of the top (index n). The top can be reached by stepping off the last step or the one before it, so the answer must look at both; best[n] already takes that min.", "Going greedy: stepping to whichever of the next two steps is cheaper. A cheap step can be followed by an expensive one, as in [2,5,1,1,9,1,1,7]; only comparing complete prices works.", "Starting only from step 0, so a free start on step 1 is never used."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "best[i] is the cheapest price to be standing on position i. There are n steps and the top is position n, so the table has n + 1 cells. The two starting positions are free, so they stay at 0.", CodeSnippet: "n = len(cost)\nbest = [0] * (n + 1)", LineStart: 2, LineEnd: 3},
				{Title: "Take the cheaper last move", Explanation: "You reached position i from i - 1 or from i - 2. Each route costs the best price to get there plus the price of leaving that step. On [10,15,20], best[2] = min(0 + 15, 0 + 10) = 10, and best[3] = min(10 + 20, 0 + 15) = 15.", CodeSnippet: "for i in range(2, n + 1):\n    best[i] = min(best[i - 1] + cost[i - 1], best[i - 2] + cost[i - 2])", LineStart: 4, LineEnd: 5},
				{Title: "Read the top", Explanation: "The top is position n, one past the last step, so its entry already accounts for both ways of getting there.", CodeSnippet: "return best[n]", LineStart: 6, LineEnd: 6},
			},
			Alternatives: []AltSolution{
				{Name: "Two variables", Code: "def minCostClimbingStairs(cost):\n    two_back, one_back = 0, 0\n    for i in range(2, len(cost) + 1):\n        two_back, one_back = one_back, min(one_back + cost[i - 1], two_back + cost[i - 2])\n    return one_back", TimeComplexity: "O(n)", SpaceComplexity: "O(1)", Explanation: "best[i] reads only the two entries before it, so keep best[i - 2] and best[i - 1] and slide them forward."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the answer best[n] and not best[n - 1]?", Answer: "The top is one position past the last step. You can arrive there by stepping off the last step (1 move) or the one before it (2 moves), and best[n] is the minimum over exactly those two. best[n - 1] only describes standing on the last step, which you still have to pay to leave."},
			{Prompt: "Counting ways used a sum and this uses a min. What stays the same in the table?", Answer: "Everything except how the choices are combined: the state (a value per position), the base cases, the fill order and the last-move argument. Sum counts the sequences through each choice; min keeps the cheapest through each choice."},
		},
	},
	{
		ID:              "drill-count-tilings",
		Number:          0,
		Kind:            KindDrill,
		Order:           4,
		Title:           "Count Tilings of a 2 by n Strip",
		Slug:            "drill-count-tilings",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Dynamic Programming", "Counting", "Math", "Foundations"},
		RelatedChapters: []int{15, 8},
		Description:     "A strip is 2 cells tall and `n` cells wide. You can fill it with two kinds of tile: a **domino** that covers 2 cells (placed standing up, covering one whole column, or lying down, covering 2 cells of one row) and a **2 by 2 square**. Every cell must be covered exactly once. Return the number of different ways to tile the strip, **modulo 1_000_000_007** (a prime used to keep numbers small).\n\nLook at the last column. Either a standing domino fills it, leaving a strip of width `n - 1`; or the last two columns are filled on their own, by two lying dominoes stacked on top of each other, or by one square, and that is **2 ways** leaving width `n - 2`. So with `ways[i]` the number of tilings of a strip of width `i`, `ways[i] = ways[i - 1] + 2 * ways[i - 2]`, and `ways[0] = ways[1] = 1` (the empty strip has one tiling, and so does a strip one column wide).\n\nThe point of this drill is the size of `n`: up to 100,000. A recursive version recurses `n` levels deep, which is more than the stack allows here. Build the answer **bottom-up**, from the smallest width to `n`, taking the remainder at every step.",
		Constraints:     []string{"1 <= n <= 100000"},
		Examples: []Example{
			{Input: "n = 2", Output: "3", Explanation: "Two standing dominoes; two lying dominoes stacked; or one square. By the recurrence ways[2] = ways[1] + 2 * ways[0] = 1 + 2."},
			{Input: "n = 4", Output: "11", Explanation: "ways[3] = 3 + 2 * 1 = 5, and ways[4] = 5 + 2 * 3 = 11."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 2}, Expected: 3, Note: "example 1: the last two columns can be filled in two extra ways"},
			{Input: map[string]interface{}{"n": 4}, Expected: 11, Note: "example 2: the coefficient 2 on the n - 2 term"},
			{Input: map[string]interface{}{"n": 1}, Expected: 1, Hidden: true, Note: "n = 1, the smallest input: one standing domino; a table must not index ways[2] or ways[-1]"},
			{Input: map[string]interface{}{"n": 3}, Expected: 5, Hidden: true, Note: "n = 3: a coefficient of 1 on the n - 2 term gives 3"},
			{Input: map[string]interface{}{"n": 10}, Expected: 683, Hidden: true, Note: "n = 10"},
			{Input: map[string]interface{}{"n": 40}, Expected: odTilings(40), Hidden: true, Timeout: 5000, Note: "n = 40: plain recursion makes billions of calls; the answer has already been through the modulus"},
			{Input: map[string]interface{}{"n": 20000}, Expected: odTilings(20000), Hidden: true, Timeout: 5000, Note: "n = 20000: the answer must be reduced modulo 1_000_000_007, and a recursion this deep (about 20000 levels) already hits the stack limit"},
			{Input: map[string]interface{}{"n": 50000}, Expected: odTilings(50000), Hidden: true, Timeout: 5000, Note: "n = 50000: more than 20000 recursion levels, which the stack here does not allow; a memoised recursion fails and the table version passes"},
			{Input: map[string]interface{}{"n": 100000}, Expected: odTilings(100000), Hidden: true, Timeout: 5000, Note: "n = 100000, the largest input"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def countTilings(n):\n    # Return the number of tilings of a 2 by n strip, modulo 1_000_000_007\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Look only at the rightmost column of the strip. What can cover it, and how much of the strip is still left to tile in each case?"},
			{Level: 2, Type: "insight", Content: "A standing domino fills the last column and leaves width n - 1. Otherwise the last two columns are filled together, in 2 different ways (two lying dominoes, or a square), leaving width n - 2. So ways[i] = ways[i - 1] + 2 * ways[i - 2], with ways[0] = ways[1] = 1. With n up to 100000 you must fill the table from the bottom up, because recursion that deep overflows the stack. See Chapter 15.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. State: ways[i] is the number of tilings of a strip of width i, for i from 0 to n.\n2. Base cases: ways[0] = ways[1] = 1.\n3. Recurrence: for i from 2 to n, ways[i] = (ways[i - 1] + 2 * ways[i - 2]) modulo 1_000_000_007.\n4. Answer: return ways[n]."},
			{Level: 4, Type: "skeleton", Content: "def countTilings(n):\n    MOD = 10 ** 9 + 7\n    ways = [1] * (n + 1)\n    for i in range(2, n + 1):\n        ways[i] = (ways[i - 1] + ____ * ways[i - 2]) % ____\n    return ways[n]"},
		},
		Solution: Solution{
			Code:            "def countTilings(n):\n    MOD = 10 ** 9 + 7\n    ways = [1] * (n + 1)\n    for i in range(2, n + 1):\n        ways[i] = (ways[i - 1] + 2 * ways[i - 2]) % MOD\n    return ways[n]",
			Explanation:     "This builds on Climb with One, Two or Three Steps, with two new ingredients: a coefficient in the recurrence (the last two columns can be filled 2 ways, so ways[i - 2] counts twice) and a big n, where the bottom-up table is required because recursion that deep overflows the stack. The same loop-from-the-base-cases shape solves Decode Ways in the ladder at any length.",
			BruteForce:      "Recurse: tilings(i) = tilings(i - 1) + 2 * tilings(i - 2). That makes exponentially many calls, so even n = 40 is out of reach, and a memoised version avoids the repeated work but still recurses n levels deep, which fails for n in the tens of thousands.",
			KeyInsight:      "Split by what covers the last column(s). The cases do not overlap and cover everything, and each leaves a smaller strip of the same kind, so a table indexed by width needs only two earlier cells.",
			WhyItWorks:      "State: ways[i] is the number of tilings of a 2 by i strip. Consider the cells of the last column. If a standing domino covers both, the rest is any tiling of width i - 1. Otherwise the top cell is covered by a lying domino or a square, which extends into column i - 1; the lying case forces a second lying domino under it (nothing else fits), and the square covers both rows, so exactly 2 ways fill the last two columns, leaving any tiling of width i - 2. These cases are disjoint and complete, so ways[i] = ways[i - 1] + 2 * ways[i - 2]. Taking the remainder after each addition is safe because addition and multiplication by 2 agree with the remainder, and it keeps every number small.",
			CommonMistakes:  []string{"Using a coefficient of 1 on ways[i - 2], forgetting that two different fillings of the last two columns exist; n = 3 gives 3 instead of 5.", "Making the recursion deeper than the stack allows. A recursive solution, even with a memo, recurses n levels, and the cases with n from 20000 up are past the limit, so build the table in a loop from 2 up to n.", "Taking the remainder only once at the end instead of at every step. The numbers grow to tens of thousands of digits, and the answer for the big cases would be slow; take the remainder inside the loop.", "Setting `ways[0] = 0`, which makes the 2-wide strip count wrong (the empty strip has one tiling, the empty one)."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "ways[i] is the number of tilings of a strip of width i. The empty strip and the one-wide strip each have exactly one tiling, so the whole table starts as 1s.", CodeSnippet: "MOD = 10 ** 9 + 7\nways = [1] * (n + 1)", LineStart: 2, LineEnd: 3},
				{Title: "Split on the last columns", Explanation: "A standing domino leaves ways[i - 1]. Filling the last two columns together can be done 2 ways (two lying dominoes, or a square), each leaving ways[i - 2]. For n = 4 this gives 1, 1, 3, 5, 11.", CodeSnippet: "for i in range(2, n + 1):\n    ways[i] = (ways[i - 1] + 2 * ways[i - 2]) % MOD", LineStart: 4, LineEnd: 5},
				{Title: "Keep the numbers small", Explanation: "The answer is asked modulo 1_000_000_007, and reducing at every step keeps each number below that prime.", CodeSnippet: "% MOD", LineStart: 5, LineEnd: 5},
				{Title: "Read the answer", Explanation: "The table was filled upwards in a loop, so no recursion was needed however large n is.", CodeSnippet: "return ways[n]", LineStart: 6, LineEnd: 6},
			},
			Alternatives: []AltSolution{
				{Name: "Two variables", Code: "def countTilings(n):\n    MOD = 10 ** 9 + 7\n    two_back, one_back = 1, 1\n    for _ in range(n - 1):\n        two_back, one_back = one_back, (one_back + 2 * two_back) % MOD\n    return one_back", TimeComplexity: "O(n)", SpaceComplexity: "O(1)", Explanation: "ways[i] reads only ways[i - 1] and ways[i - 2], so the table shrinks to two variables that slide forward."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is the table version safe for n = 100000 when a recursive version is not?", Answer: "A loop runs in constant stack space, but each recursive call waits on the one below it, so n calls are on the stack at the deepest point (here the sandbox limits recursion to 20000 levels). Bottom-up also never needs a memo dictionary, since each cell is computed once in order."},
			{Prompt: "Why is it fine to take the remainder at every step, when the question only asks for it at the end?", Answer: "Remainders respect addition and multiplication: (a + b) % m equals ((a % m) + (b % m)) % m, and the same for multiplying by 2. So reducing early gives the same final remainder, while keeping every intermediate number small. Never use remainders with subtraction or division this way without care."},
		},
	},
	{
		ID:              "drill-max-subarray",
		Number:          0,
		Kind:            KindDrill,
		Order:           5,
		Title:           "Maximum Subarray",
		Slug:            "drill-max-subarray",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Dynamic Programming", "Array", "Foundations"},
		RelatedChapters: []int{15, 8},
		Description:     "Given a non-empty list of integers `nums`, return the largest sum of any **subarray**: a run of one or more elements next to each other. The numbers can be negative, so the best subarray might be a single element, and if every number is negative the answer is the largest (least negative) one.\n\nThis is a 1-D DP whose state is not the answer itself. Let `ending[i]` be the best sum of a subarray that **ends exactly at index `i`**. That subarray is either the single element `nums[i]`, or `nums[i]` added to the best subarray ending at `i - 1`, so `ending[i] = max(nums[i], ending[i - 1] + nums[i])` with `ending[0] = nums[0]`. The answer is the **largest value anywhere in the table**, since the best subarray ends somewhere.",
		Constraints:     []string{"1 <= len(nums) <= 100000", "-10^4 <= nums[i] <= 10^4"},
		Examples: []Example{
			{Input: "nums = [-2,1,-3,4,-1,2,1,-5,4]", Output: "6", Explanation: "The table of best sums ending at each index is -2, 1, -2, 4, 3, 5, 6, 1, 5. The largest is 6, from the subarray [4,-1,2,1]."},
			{Input: "nums = [-3,-1,-2]", Output: "-1", Explanation: "Every number is negative, so the best subarray is the single element -1. An answer of 0 would mean an empty subarray, which is not allowed."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, Expected: 6, Note: "example 1: the best subarray is in the middle"},
			{Input: map[string]interface{}{"nums": []int{-3, -1, -2}}, Expected: -1, Note: "example 2: all negative, so the answer is the largest single element"},
			{Input: map[string]interface{}{"nums": []int{7}}, Expected: 7, Hidden: true, Note: "a single element"},
			{Input: map[string]interface{}{"nums": []int{-5}}, Expected: -5, Hidden: true, Note: "a single negative element: no empty subarray"},
			{Input: map[string]interface{}{"nums": []int{5, 4, -1, 7, 8}}, Expected: 23, Hidden: true, Note: "the whole list is the best subarray"},
			{Input: map[string]interface{}{"nums": []int{1, -1, 1, -1, 5}}, Expected: 5, Hidden: true, Note: "the best run ends at the last element"},
			{Input: map[string]interface{}{"nums": []int{8, -20, 3, 2}}, Expected: 8, Hidden: true, Note: "the best run is at the very start; the table's maximum is at index 0, so returning its last cell is wrong"},
			{Input: map[string]interface{}{"nums": []int{4, -1, -1, 4}}, Expected: 6, Hidden: true, Note: "a dip that is worth carrying through: 4 + (-1) + (-1) + 4 beats either 4 alone"},
			{Input: map[string]interface{}{"nums": []int{-1, 0, -2}}, Expected: 0, Hidden: true, Note: "a zero is a real element, so 0 is the answer here"},
			{Input: map[string]interface{}{"nums": odNoise(100000)}, Expected: odMaxSub(odNoise(100000)), Hidden: true, Timeout: 5000, Note: "100000 numbers: trying every subarray is billions of sums and runs out of time; one pass is O(n)"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(n)",
		StarterCode:     "def maxSubArray(nums):\n    # Return the largest sum of a non-empty subarray\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose you only look at subarrays that end exactly at index i. When you move to index i + 1, what two choices do you have for the best subarray ending there?"},
			{Level: 2, Type: "insight", Content: "Let ending[i] be the best sum of a subarray that ends at i. It either starts fresh at i (just nums[i]) or extends the best one ending at i - 1, so ending[i] = max(nums[i], ending[i - 1] + nums[i]). Carrying a negative sum forward only hurts, so the max drops it. The overall answer is the largest ending[i], because the best subarray ends somewhere. See Chapter 15.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. State: ending[i] is the best sum of a subarray ending at index i.\n2. Base case: ending[0] = nums[0].\n3. Recurrence: for i from 1, ending[i] = max(nums[i], ending[i - 1] + nums[i]).\n4. Answer: return the largest entry of ending (not the last one)."},
			{Level: 4, Type: "skeleton", Content: "def maxSubArray(nums):\n    ending = [0] * len(nums)\n    ending[0] = nums[0]\n    for i in range(1, len(nums)):\n        ending[i] = max(____, ending[i - 1] + ____)\n    return max(____)"},
		},
		Solution: Solution{
			Code:            "def maxSubArray(nums):\n    ending = [0] * len(nums)\n    ending[0] = nums[0]\n    for i in range(1, len(nums)):\n        ending[i] = max(nums[i], ending[i - 1] + nums[i])\n    return max(ending)",
			Explanation:     "This builds on Min Cost Climbing Stairs: a table with a max over two choices (start over or extend) instead of a min. What is new is that the state is not the thing asked for: ending[i] is the best subarray that ends at i, and the answer is the best entry in the whole table. This best-ending-here idea (Kadane's algorithm) is the pattern behind Maximum Product Subarray and Longest Increasing Subsequence in the ladder.",
			BruteForce:      "Sum every subarray, with a running sum for each start: O(n^2). For 100000 numbers that is about 5 billion additions. The waste is that the best subarray ending at i is rebuilt from scratch for every start, when it can be read off the best one ending at i - 1.",
			KeyInsight:      "Define the state as the best subarray ending exactly at i. Then the choices for i are only two (start here, or extend the best ending at i - 1), and the global answer is the maximum over all i.",
			WhyItWorks:      "State: ending[i] is the largest sum among subarrays whose last element is nums[i]. Such a subarray is either just nums[i], or nums[i] preceded by a subarray that ends at i - 1; for the second kind the sum is largest when the earlier part is the best one, ending[i - 1]. So ending[i] = max(nums[i], ending[i - 1] + nums[i]), and ending[0] = nums[0] because only one subarray ends at index 0. Every subarray ends at some index, so the best subarray overall is the largest ending[i].",
			CommonMistakes:  []string{"Starting the best answer at 0, or clamping each entry with `max(0, ...)`. That allows an empty subarray, so an all-negative list wrongly returns 0; the answer must be at least one element.", "Returning the last table entry. ending[n - 1] is the best subarray that ends at the last index, not the best anywhere; [8,-20,3,2] must return 8.", "Always extending (`ending[i] = ending[i - 1] + nums[i]`), so a negative running sum is carried forward for no reason. The option of starting fresh at nums[i] is the other half of the recurrence.", "Trying every subarray with two loops. It is correct but O(n^2), and it times out on 100000 numbers."},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state", Explanation: "ending[i] will be the best sum of a subarray that ends at index i. The table has one cell per index.", CodeSnippet: "ending = [0] * len(nums)", LineStart: 2, LineEnd: 2},
				{Title: "Set the base case", Explanation: "Only one subarray ends at index 0, namely [nums[0]], so that is its best.", CodeSnippet: "ending[0] = nums[0]", LineStart: 3, LineEnd: 3},
				{Title: "Start over or extend", Explanation: "At index i the best subarray ending there is nums[i] alone, or nums[i] added to the best one ending at i - 1. On [-2,1,-3,4] the cells are -2, 1 (start over, since -2 + 1 is smaller), -2, and 4 (start over, since -2 + 4 = 2 is smaller).", CodeSnippet: "for i in range(1, len(nums)):\n    ending[i] = max(nums[i], ending[i - 1] + nums[i])", LineStart: 4, LineEnd: 5},
				{Title: "Take the best end", Explanation: "The best subarray overall ends at some index, so the answer is the largest cell in the table.", CodeSnippet: "return max(ending)", LineStart: 6, LineEnd: 6},
			},
			Alternatives: []AltSolution{
				{Name: "Kadane with two variables", Code: "def maxSubArray(nums):\n    ending_here = best = nums[0]\n    for value in nums[1:]:\n        ending_here = max(value, ending_here + value)\n        best = max(best, ending_here)\n    return best", TimeComplexity: "O(n)", SpaceComplexity: "O(1)", Explanation: "ending[i] only reads ending[i - 1], and the answer is the maximum so far, so keep the previous cell and a running maximum instead of the table."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why define the state as the best subarray ending at i, and not simply the best subarray in the first i elements?", Answer: "The best in the first i + 1 elements cannot be built from the best in the first i: whether nums[i + 1] helps depends on the sum of the run just before it, which that value does not record. 'Ends exactly at i' is the extra information that makes the next cell a function of the previous one, and it is why the answer is the max over the table."},
			{Prompt: "What would change if the subarray were allowed to be empty?", Answer: "Then the answer is never below 0, so you could clamp each cell with max(0, ...) or start the best at 0, and an all-negative list returns 0. Here the statement requires at least one element, so those clamps are bugs."},
		},
	},
	{
		ID:              "drill-coin-change-ways",
		Number:          0,
		Kind:            KindDrill,
		Order:           6,
		Title:           "Coin Change II",
		Slug:            "drill-coin-change-ways",
		Difficulty:      "Medium",
		Category:        "1d-dp",
		Tags:            []string{"Dynamic Programming", "Array", "Counting", "Foundations"},
		RelatedChapters: []int{15, 8},
		Description:     "You have coins of the values in `coins` (distinct positive integers) and an unlimited number of each. Return the number of **combinations** of coins that add up to exactly `amount`. A combination ignores order: 2+3 and 3+2 are the same combination, so they count once. Return 0 if no combination works, and 1 if `amount` is 0 (the empty combination). The answer fits in a 64-bit integer.\n\nLet `ways[a]` be the number of combinations that make `a`, with `ways[0] = 1`. Process the coins **one at a time**, in an outer loop, and for each coin sweep the amounts upwards: `ways[a] += ways[a - coin]`. After finishing a coin, `ways[a]` counts the combinations that use only the coins seen so far; that is what stops a different order of the same coins being counted again. **The order of the two loops is the whole lesson**: with the amounts in the outer loop and the coins inside, the same table counts ordered sequences (permutations) instead.",
		Constraints:     []string{"1 <= len(coins) <= 6", "1 <= coins[i] <= 100, all different", "0 <= amount <= 10000"},
		Examples: []Example{
			{Input: "amount = 5, coins = [1,2,5]", Output: "4", Explanation: "5, 2+2+1, 2+1+1+1 and 1+1+1+1+1. Counting 2+2+1 and 1+2+2 separately would give 9, which counts orders."},
			{Input: "amount = 3, coins = [2]", Output: "0", Explanation: "Two cannot add up to 3, so no combination works."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"amount": 5, "coins": []int{1, 2, 5}}, Expected: 4, Note: "example 1: combinations, not orderings (9 if order counted)"},
			{Input: map[string]interface{}{"amount": 3, "coins": []int{2}}, Expected: 0, Note: "example 2: impossible amounts count 0 ways, not -1"},
			{Input: map[string]interface{}{"amount": 0, "coins": []int{7}}, Expected: 1, Hidden: true, Note: "amount 0: the empty combination is the one way"},
			{Input: map[string]interface{}{"amount": 10, "coins": []int{10}}, Expected: 1, Hidden: true, Note: "one coin that fits exactly"},
			{Input: map[string]interface{}{"amount": 4, "coins": []int{1, 2, 3}}, Expected: 4, Hidden: true, Note: "4 combinations (7 orderings): loop order matters"},
			{Input: map[string]interface{}{"amount": 8, "coins": []int{3, 5}}, Expected: 1, Hidden: true, Note: "only 3 + 5 works"},
			{Input: map[string]interface{}{"amount": 7, "coins": []int{2, 4}}, Expected: 0, Hidden: true, Note: "an odd amount from even coins"},
			{Input: map[string]interface{}{"amount": 100, "coins": []int{1, 5, 10, 25, 50}}, Expected: 292, Hidden: true, Note: "the classic: 292 ways to make 100 from US coins"},
			{Input: map[string]interface{}{"amount": 10000, "coins": []int{1, 2, 5, 10, 20, 50}}, Expected: odCoinWays(10000, []int{1, 2, 5, 10, 20, 50}), Hidden: true, Timeout: 5000, Note: "amount 10000: there are trillions of combinations, so counting them one by one (plain recursion) cannot finish; the table does about 60000 additions"},
		},
		TimeComplexity:  "O(amount * len(coins))",
		SpaceComplexity: "O(amount)",
		StarterCode:     "def change(amount, coins):\n    # Return the number of coin combinations that make amount\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you let the table remember every order, 1+2 and 2+1 are separate entries. How could you make the table add a coin type once, so that a combination is built in only one order?"},
			{Level: 2, Type: "insight", Content: "Let ways[a] be the number of combinations that make a, with ways[0] = 1. Put the coins in the outer loop: for each coin, sweep a upwards and add ways[a - coin] to ways[a]. After processing a coin, the table counts combinations that use only the coins processed so far, so each combination is built in the order of the coin list, once. Swapping the two loops counts orderings instead. See Chapter 15.", ChapterRef: 15},
			{Level: 3, Type: "plan", Content: "1. State: ways[a] is the number of combinations of the coins so far that make a, for a from 0 to amount.\n2. Base case: ways[0] = 1 and every other entry 0.\n3. Recurrence: for each coin in turn (outer loop), for a from coin to amount (inner loop), ways[a] += ways[a - coin].\n4. Answer: return ways[amount]."},
			{Level: 4, Type: "skeleton", Content: "def change(amount, coins):\n    ways = [0] * (amount + 1)\n    ways[0] = 1\n    for coin in ____:\n        for total in range(coin, amount + 1):\n            ways[total] += ways[____]\n    return ways[amount]"},
		},
		Solution: Solution{
			Code:            "def change(amount, coins):\n    ways = [0] * (amount + 1)\n    ways[0] = 1\n    for coin in coins:\n        for total in range(coin, amount + 1):\n            ways[total] += ways[total - coin]\n    return ways[amount]",
			Explanation:     "This builds on Climb with One, Two or Three Steps, where the counts were of ordered sequences; here the same 'add the table entry a move earlier' recurrence counts unordered combinations because the coins are the outer loop. The state is still one number per amount, with ways[0] = 1 as the base case, so this is the table the ladder's Coin Change reuses with a min in place of the sum. If the loops were swapped (amounts outside, coins inside), the table would count ordered sequences, as in Combination Sum IV (LeetCode 377).",
			BruteForce:      "Recurse over (amount left, index of the coin you may still use): either use the coin again or move on to the next one. It is correct but explores every combination one at a time, so with trillions of answers it never finishes. The waste is that many different prefixes leave the same (amount left, coin index) state.",
			KeyInsight:      "Add the coin types one at a time. Once a coin is finished, ways[a] counts combinations from the coins so far, so a later coin can only extend them: a combination can only be built in the order of the coin list, once.",
			WhyItWorks:      "Invariant: after the outer loop has processed the first j coins, ways[a] is the number of combinations of those j coins that make a. It holds for j = 0 (only the empty combination makes 0, so ways[0] = 1 and the rest 0). Processing a new coin c: a combination of the first j + 1 coins that makes a uses c some number of times, and removing one copy of c leaves a combination of the first j + 1 coins that makes a - c (or none, if there was no c, which is the old ways[a]). The in-place sweep upwards makes ways[a - c] already include uses of c, so adding it counts every number of copies once. Because the coins are handled in a fixed order, a combination such as 2+1 is built as 'one coin 1, then one coin 2' and no other way, so orderings are not counted.",
			CommonMistakes:  []string{"Putting the amounts in the outer loop and the coins in the inner loop. That counts ordered sequences: 1+2 and 2+1 are counted twice, so amount 5 with [1,2,5] returns 9 and not 4. Coins outside counts combinations.", "Setting `ways[0] = 0`. Every entry then stays 0 because everything is built on top of the one empty combination.", "Using `-1` for impossible amounts, as in Coin Change. Here the table counts, so an impossible amount has 0 ways and the sum just works.", "Plain recursion without a memo. It is correct on the small cases and cannot finish on an amount of 10000.", "Sweeping the amounts downwards in the inner loop. That lets each coin be used at most once, which turns this into the 0/1 version of the problem."},
			TimeComplexity:  "O(amount * len(coins))",
			SpaceComplexity: "O(amount)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the state and base case", Explanation: "ways[a] counts combinations that make a. The empty combination makes 0, so ways[0] = 1 and everything else starts at 0.", CodeSnippet: "ways = [0] * (amount + 1)\nways[0] = 1", LineStart: 2, LineEnd: 3},
				{Title: "Add one coin type at a time", Explanation: "The outer loop is over coins, so a combination is only ever built in the order of the coin list. This is what makes 2+1 and 1+2 count once.", CodeSnippet: "for coin in coins:", LineStart: 4, LineEnd: 4},
				{Title: "Extend by one more copy", Explanation: "For each amount from coin up, every combination that makes total - coin gets one more copy of this coin. Going upwards lets the same coin be used again. With [1,2,5]: after coin 1 every ways[a] is 1; after coin 2 ways[4] = 3; after coin 5 ways[5] = 4.", CodeSnippet: "for total in range(coin, amount + 1):\n    ways[total] += ways[total - coin]", LineStart: 5, LineEnd: 6},
				{Title: "Read the answer", Explanation: "After every coin has been added, ways[amount] is the number of combinations. Impossible amounts stay at 0.", CodeSnippet: "return ways[amount]", LineStart: 7, LineEnd: 7},
			},
			Alternatives: []AltSolution{
				{Name: "Table of coins by amount", Code: "def change(amount, coins):\n    # ways[j][a]: combinations of the first j coins that make a\n    ways = [[0] * (amount + 1) for _ in range(len(coins) + 1)]\n    for j in range(len(coins) + 1):\n        ways[j][0] = 1\n    for j in range(1, len(coins) + 1):\n        coin = coins[j - 1]\n        for total in range(1, amount + 1):\n            ways[j][total] = ways[j - 1][total]\n            if total >= coin:\n                ways[j][total] += ways[j][total - coin]\n    return ways[len(coins)][amount]", TimeComplexity: "O(amount * len(coins))", SpaceComplexity: "O(amount * len(coins))", Explanation: "The same recurrence with the coin index kept as a second dimension: either skip coin j (ways[j - 1][a]) or use one more copy of it (ways[j][a - coin]). The one-row table above is this one with the coin dimension folded away."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does swapping the two loops change the answer from 4 to 9 for amount 5 with [1,2,5]?", Answer: "With the amounts outside, ways[a] adds ways[a - coin] for every coin as the last move, so it counts sequences by their last coin and 1+2 and 2+1 end differently and are both counted. With the coins outside, each coin's pass only extends combinations made from earlier coins, so a multiset of coins can be built in just one order."},
			{Prompt: "How does this table change to find the fewest coins, as in Coin Change, and what plays the role of the 0 for impossible amounts?", Answer: "Replace the sum with a min: best[a] = min(best[a], best[a - coin] + 1), with best[0] = 0, and use infinity for unreachable amounts (not 0 or -1 inside the loop, which would look like a real answer); return -1 at the end if the result is still infinity. Counting uses 0 for impossible because adding 0 changes nothing."},
		},
	},
}

func odFib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func odClimb123(n int) int {
	ways := make([]int, n+1)
	ways[0] = 1
	for i := 1; i <= n; i++ {
		for s := 1; s <= 3; s++ {
			if i-s >= 0 {
				ways[i] += ways[i-s]
			}
		}
	}
	return ways[n]
}

// odStairCosts is a deterministic mix of cheap and expensive steps.
func odStairCosts(n int) []int {
	cost := make([]int, n)
	seed := 12345
	for i := range cost {
		seed = (seed*1103515245 + 12345) % 2147483648
		cost[i] = (seed >> 8) % 100
	}
	return cost
}

func odMinCost(cost []int) int {
	n := len(cost)
	best := make([]int, n+1)
	for i := 2; i <= n; i++ {
		a, b := best[i-1]+cost[i-1], best[i-2]+cost[i-2]
		if b < a {
			a = b
		}
		best[i] = a
	}
	return best[n]
}

func odTilings(n int) int {
	const mod = 1000000007
	ways := make([]int, n+1)
	ways[0] = 1
	if n >= 1 {
		ways[1] = 1
	}
	for i := 2; i <= n; i++ {
		ways[i] = (ways[i-1] + 2*ways[i-2]) % mod
	}
	return ways[n]
}

// odNoise is a deterministic list of positive and negative numbers.
func odNoise(n int) []int {
	nums := make([]int, n)
	seed := 987
	for i := range nums {
		seed = (seed*1103515245 + 12345) % 2147483648
		nums[i] = (seed>>8)%201 - 101
	}
	return nums
}

func odMaxSub(nums []int) int {
	ending, best := nums[0], nums[0]
	for _, v := range nums[1:] {
		if ending+v > v {
			ending += v
		} else {
			ending = v
		}
		if ending > best {
			best = ending
		}
	}
	return best
}

func odCoinWays(amount int, coins []int) int {
	ways := make([]int, amount+1)
	ways[0] = 1
	for _, c := range coins {
		for t := c; t <= amount; t++ {
			ways[t] += ways[t-c]
		}
	}
	return ways[amount]
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, OneDDPDrills...)
}
