package chapter

import "time"

// chapter15 teaches one-dimensional dynamic programming through a six-step
// recipe whose subgoal labels are reused by every 1d-dp practice problem:
// Define the state, Write the recurrence, Set the base cases, Choose the
// evaluation order, Locate the answer, Optimise space.
// Storyboards (IDs registered in app.js): memoization, iterative-dp.
func chapter15() Chapter {
	return Chapter{
		Number:        15,
		Slug:          "1d-dp",
		Title:         "Dynamic Programming in One Dimension",
		Category:      "1d-dp",
		Prerequisites: []int{8},
		Summary:       "Dynamic programming (DP) is recursion that remembers. When a recursive solution asks the same smaller question many times, you answer each question once and store it. This chapter gives you a six-step recipe (define the state, write the recurrence, set the base cases, choose the evaluation order, locate the answer, optimise space) and walks Climbing Stairs from an exponential recursion to memoization, a bottom-up table and finally two variables.",
		Motivation:    "Climbing Stairs: a staircase has n steps and each move climbs 1 or 2 of them; how many different move sequences reach the top? The last move came from step n-1 or step n-2, so ways(n) = ways(n-1) + ways(n-2), and Chapter 8 makes that a three-line recursive function. Run it for n = 45 and it makes 3,672,623,805 calls, several minutes of Python, to produce 1,836,311,903. Yet there are only 46 different questions it can ask, ways(0) to ways(45). The recursion is slow because it asks the same question over and over: inside ways(45), ways(2) alone is recomputed 701,408,733 times. If we write each answer down the first time and look it up afterwards, the 3.6 billion calls become 89 calls, or a loop of 44 additions. That is the whole idea of dynamic programming. The hard part is not the remembering; it is finding the right question to remember answers to, and this chapter gives you a fixed recipe for that.",
		Objectives: []string{
			"Explain overlapping subproblems and optimal substructure, using call counts (15 calls for n = 5, 3.6 billion for n = 45) and the Coin Change counterexample to greedy",
			"Apply the six-step DP recipe: define the state in words, write the recurrence from the last choice, set base cases, choose an evaluation order, locate the answer, optimise space",
			"Convert a plain recursion into a memoized (top-down) version and into a tabulated (bottom-up) version, and trace the table on a small input",
			"Reduce a table to rolling variables when the recurrence only reads a fixed number of earlier entries, and explain when that is not possible",
			"Recognise when a problem needs DP rather than greedy, backtracking or a sliding window",
		},
		WhenToUse: []string{
			"The question asks for a count of ways, a minimum or maximum, or a yes/no 'is it possible', not for a list of every solution",
			"A natural recursive solution exists, but it calls itself on the same arguments many times (draw the call tree and look for repeats)",
			"Each answer can be built from answers to smaller versions of the same question: a shorter prefix, a smaller amount, an earlier position",
			"There is a sequence of choices (take or skip, which coin, how many digits, which word) and the best result after a choice does not depend on how you got there",
			"Greedy looks tempting but you can construct a small counterexample (coins [1,3,4] for amount 6)",
			"Input sizes around 10^2 to 10^4 where O(n^2) or O(n * amount) is acceptable but 2^n is not",
		},
		Concepts: []Concept{
			{
				Name:        "The DP recipe: six steps, every time",
				Description: "Every DP solution in this chapter and its problem ladder is built with the same six steps. Name them as you work; the names are the plan.",
				WhyItWorks:  "A DP is a proof by induction turned into code. Step 1 (Define the state) fixes what each subproblem means, in words, e.g. 'ways[i] = number of move sequences ending on step i'. Step 2 (Write the recurrence) is the inductive step: look at the last choice made (last move, last coin, last house) and express the answer through smaller states. It is correct when the cases you list are complete (every solution makes some last choice) and, for counting, disjoint. Step 3 (Set the base cases) is the induction base: the smallest states that the recurrence cannot reach. Step 4 (Choose the evaluation order) makes sure every state you read is already final. Step 5 (Locate the answer) says which cell (or the max over which cells) answers the original question. Step 6 (Optimise space) drops what is never read again. Skipping step 1 is the usual cause of wrong DPs: if you cannot say what dp[i] means, you cannot check the recurrence.",
				Intuition:   "It is like filling in a times table in primary school. First you decide what each box means (row times column), then the rule for filling a box from boxes already filled, then the boxes you fill by hand (the 1-row), then the order (row by row), then which box answers your question. Once the boxes mean something precise, the rule almost writes itself.",
				CoreIdeas: []string{
					"1. Define the state: say in one sentence what dp[i] means, including the units (ways, coins, money, True/False).",
					"2. Write the recurrence: ask 'what was the last choice?' and combine the smaller states for each possible last choice (+ for counting, max/min for optimising, or for yes/no).",
					"3. Set the base cases: the states the recurrence cannot compute, usually the empty input (dp[0]).",
					"4. Choose the evaluation order: every state must be computed after the states it reads; for prefix states that means left to right.",
					"5. Locate the answer: dp[n] for prefix states; max(dp) when states mean 'ending exactly at i'.",
					"6. Optimise space: if the recurrence reads only the last k entries, keep k variables instead of a table.",
				},
				CommonMistakes: []string{
					"Writing code before defining the state in words. Then off-by-one errors appear everywhere (is dp[i] about the first i items or item i?). Write the definition as a comment first.",
					"Choosing base cases by trial and error until the examples pass. Derive them from the state's meaning: ways[0] = 1 because the empty move sequence is one way to stand on the ground.",
					"Optimising space before the table version works. Get the full table right (it is easy to print and trace), then remove what is never read.",
				},
				Examples: []CodeExample{
					{
						ID:       "dp-recipe-template",
						Title:    "The recipe as a code template",
						Language: "python",
						Snippet: `def solve(items):
    n = len(items)
    # 1. Define the state: dp[i] = answer for the first i items
    dp = [0] * (n + 1)
    # 3. Set the base cases
    dp[0] = BASE
    # 4. Choose the evaluation order: smaller prefixes first
    for i in range(1, n + 1):
        # 2. Write the recurrence: combine states for each possible last choice
        dp[i] = combine(dp[i - 1], dp[i - 2], items[i - 1])
    # 5. Locate the answer
    return dp[n]
    # 6. Optimise space: if only dp[i-1], dp[i-2] are read, keep two variables`,
						Notes: "BASE and combine are placeholders; every ladder problem fills them in differently.",
					},
				},
			},
			{
				Name:        "Overlapping subproblems and memoization (top-down)",
				Description: "A recursion has overlapping subproblems when it calls itself with the same arguments more than once. Memoization stores each result the first time and returns the stored value on every repeat.",
				WhyItWorks:  "Count the calls. The plain recursion climb(n) = climb(n-1) + climb(n-2) makes 15 calls for n = 5, 177 for n = 10, 21,891 for n = 20 and 3,672,623,805 for n = 45, growing by a factor of about 1.618 per step. But there are only n + 1 different arguments. For n = 5 the 15 calls are: climb(5) once, climb(4) once, climb(3) twice, climb(2) three times, climb(1) five times, climb(0) three times. Those repeats are the overlap. A function whose result depends only on its arguments (no hidden state) gives the same answer every time, so caching is safe. With a memo, each argument's work is done once; every other call is an O(1) lookup, so the total is (number of states) x (work per state) = (n + 1) x O(1). For n = 45 the memoized version makes 89 calls. The same formula, states times work per state, gives the running time of every DP.",
				Intuition:   "A student asked 'what is 7 x 8?' every five minutes can work it out each time, or write '56' on a sticky note the first time. The sticky notes are the memo, the question '7 x 8' is the state, and the reason it works is that the answer never changes between askings.",
				CoreIdeas: []string{
					"Overlap test: draw the first few levels of the call tree; if any argument appears twice, memoization will help.",
					"Memoize by checking the memo before computing and storing the result before returning; the base cases can go straight into the memo.",
					"Running time = number of distinct states x work per state (climbing stairs: n states x O(1); coin change: amount states x k coins).",
					"Top-down recursion visits only the states actually needed, but uses a call stack as deep as the longest chain of calls.",
				},
				CommonMistakes: []string{
					"Memoizing a function whose result depends on something not in the key (a global list being modified, a 'path so far' parameter). Every input that changes the answer must be part of the state.",
					"Storing the result but forgetting to check the memo first (or checking it but never storing). Then nothing is reused; count calls on a small input to make sure each state is computed once.",
					"Recursing thousands of levels deep, e.g. memoized coin change with amount 10,000 and a 1-coin. Python's default limit is 1,000 frames; the bottom-up table has no such limit.",
				},
				Examples: []CodeExample{
					{
						ID:       "climb-memo",
						Title:    "Plain recursion, then memoized",
						Language: "python",
						Snippet: `def climb_plain(n):                 # 3,672,623,805 calls for n = 45
    if n <= 1:
        return 1
    return climb_plain(n - 1) + climb_plain(n - 2)

def climb_memo(n, memo=None):        # 89 calls for n = 45
    if memo is None:
        memo = {0: 1, 1: 1}
    if n not in memo:
        memo[n] = climb_memo(n - 1, memo) + climb_memo(n - 2, memo)
    return memo[n]`,
						Notes: "The memoized version is the plain recursion plus three lines. Both compute the same recurrence.",
					},
				},
			},
			{
				Name:        "Optimal substructure, and when greedy fails",
				Description: "A problem has optimal substructure when an optimal solution is built from optimal solutions to its subproblems. It is what makes a max/min recurrence correct. Greedy algorithms make one locally best choice; DP tries every possible last choice.",
				WhyItWorks:  "Coin Change: coins [1,3,4], amount 6, fewest coins. Greedy takes the biggest coin that fits: 4, then 1, then 1, which is 3 coins. But 3 + 3 uses only 2. Greedy commits to 4 and never revisits that choice. DP instead asks 'what was the last coin?' and tries all of them. Optimal substructure is the justification: if an optimal pile for 6 ends with coin c, the rest must be an optimal pile for 6 - c (otherwise swap in a smaller pile and beat the optimum). So fewest[a] = 1 + min(fewest[a - c]) over coins c <= a. The table for amounts 0..6 is [0, 1, 2, 1, 1, 2, 2]: fewest[6] = min(fewest[5] + 1, fewest[3] + 1, fewest[2] + 1) = min(3, 2, 3) = 2. Greedy is only correct when you can prove its one choice is always safe (an exchange argument, Chapter 17); DP needs no such proof, because it checks every choice.",
				Intuition:   "Planning a road trip through checkpoints: if the fastest route from home to the coast passes through Lyon, then the part from Lyon to the coast must be the fastest Lyon-to-coast route, or you could swap it in and arrive sooner. That is optimal substructure. Greedy is taking the motorway exit that looks fastest right now; DP is comparing every possible last checkpoint before the coast.",
				CoreIdeas: []string{
					"Optimal substructure: the best answer for a state is built from the best answers of smaller states, so storing only the best per state loses nothing.",
					"Overlapping subproblems make DP fast; optimal substructure makes it correct. You need both.",
					"Greedy picks one choice per step; DP takes a max/min over all choices per step. Test a greedy idea on small cases before trusting it: [1,3,4] and 6 breaks largest-coin-first.",
					"If the best answer for a state depends on how you reached it, the state is missing information; add it (e.g. 'was house 0 taken?', or the max AND min product).",
				},
				CommonMistakes: []string{
					"Trusting greedy because it works for familiar coins (1, 5, 10, 25). Always try to break it with a small example; if you can, use DP.",
					"Keeping only the best value per state when the future also needs another value. In Maximum Product Subarray, a negative number turns the smallest product into the largest, so the state must hold both.",
					"Counting with max or optimising with +: count recurrences add disjoint cases, optimisation recurrences take max/min over all cases.",
				},
				Examples: []CodeExample{
					{
						ID:       "coin-greedy-vs-dp",
						Title:    "Greedy versus DP on coins [1,3,4], amount 6",
						Language: "python",
						Snippet: `def greedy(coins, amount):           # returns 3 (4 + 1 + 1): wrong
    count = 0
    for coin in sorted(coins, reverse=True):
        count += amount // coin
        amount %= coin
    return count if amount == 0 else -1

def fewest_coins(coins, amount):     # returns 2 (3 + 3)
    INF = float("inf")
    fewest = [0] + [INF] * amount
    for a in range(1, amount + 1):
        for coin in coins:
            if coin <= a:
                fewest[a] = min(fewest[a], fewest[a - coin] + 1)
    return fewest[amount] if fewest[amount] != INF else -1`,
						Notes: "The DP table for amounts 0..6 is [0, 1, 2, 1, 1, 2, 2].",
					},
				},
			},
			{
				Name:        "Tabulation (bottom-up) and rolling variables",
				Description: "Tabulation fills a table from the base cases upward in a loop, so every state is computed before it is needed. When the recurrence reads only the last few entries, the table can shrink to a few variables.",
				WhyItWorks:  "Memoization lets the recursion discover the order; tabulation chooses it in advance. For ways[i] = ways[i-1] + ways[i-2], every state reads smaller indices only, so the loop i = 2, 3, ..., n is a valid order: by the time it reaches i, both inputs are final (this is the induction made explicit). The table for n = 5 is [1, 1, 2, 3, 5, 8]: four additions, no recursion, no stack. Now look at which cells are read: computing ways[i] reads only ways[i-1] and ways[i-2]; after that, ways[i-2] is never read again. So two variables are enough: (two_back, one_back) = (one_back, one_back + two_back), O(1) space. This works whenever the recurrence reads a fixed window of earlier entries. It does not work for Coin Change (fewest[a] reads back as far as the largest coin) or Longest Increasing Subsequence (ends_at[i] reads every j < i), where the table stays.",
				Intuition:   "Climbing an actual staircase while counting: you only need to remember the counts written on the last two steps, not the whole staircase. The full table is a notebook with one page per step; rolling variables are two sticky notes you keep rewriting as you climb.",
				CoreIdeas: []string{
					"Bottom-up = base cases first, then a loop in an order where every dependency comes earlier. No recursion, no stack depth limit.",
					"Top-down and bottom-up compute the same states with the same recurrence; they differ only in who decides the order.",
					"Rolling variables: if dp[i] reads only dp[i-1], ..., dp[i-k], keep k variables and update them with a simultaneous assignment.",
					"Keep the full table while debugging, and whenever you need to reconstruct the choices (which houses, which coins), not just the value.",
				},
				CommonMistakes: []string{
					"Updating rolling variables one at a time: one_back = one_back + two_back, then two_back = one_back, which uses the new value. Use a tuple assignment, or a temporary variable.",
					"Looping in an order that reads unfinished cells (e.g. filling a palindrome table row by row with i increasing, when is_pal[i][j] needs row i+1). Check that every cell read is already final.",
					"Rolling the table when the recurrence reads further back than you kept, or when you still need the table to reconstruct the answer.",
				},
				Examples: []CodeExample{
					{
						ID:       "climb-table-rolling",
						Title:    "Bottom-up table, then two rolling variables",
						Language: "python",
						Snippet: `def climb_table(n):
    ways = [0] * (n + 1)
    ways[0] = ways[1] = 1
    for i in range(2, n + 1):
        ways[i] = ways[i - 1] + ways[i - 2]
    return ways[n]                       # n = 5: table [1, 1, 2, 3, 5, 8]

def climb_rolling(n):
    two_back, one_back = 1, 1            # ways[0], ways[1]
    for _ in range(2, n + 1):
        two_back, one_back = one_back, one_back + two_back
    return one_back`,
						Notes: "Same recurrence, same order; the rolling version just forgets cells that will never be read again.",
					},
				},
			},
		},
		Animations: []Storyboard{
			{
				ID:    "memoization",
				Title: "Memoization with Fibonacci",
				Goal:  "See which calls do real work and which become lookups when fib(6) is memoized: 11 calls instead of 25.",
				MemoryTips: []string{
					"1. Define the state: fib(n) depends only on n, so n is the memo key.",
					"2. Write the recurrence: fib(n) = fib(n-1) + fib(n-2).",
					"3. Set the base cases: fib(0) = 0, fib(1) = 1 stop the recursion.",
					"4. Evaluation order comes free: the recursion goes deep first and stores answers on the way back up.",
					"Check the memo before computing; store before returning. Each state is computed once.",
				},
				Steps: []StoryboardStep{
					{Cue: "call", Narration: "fib(6) is called. Without a memo this single call would grow into 25 calls; watch how many are actually computed. The call stack grows downwards: fib(6) needs fib(5), which needs fib(4), and so on to the base cases.", VisualHint: "Call stack grows; memo table is empty", Duration: Duration(2 * time.Second), CodeRef: "fib: call fib(n)"},
					{Cue: "check-memo", Narration: "Every call first looks in the memo. The first time we reach fib(4) the memo has no entry, so we must compute it. The second time fib(4) is needed (by fib(6)), the entry is there and the call returns immediately: that lookup replaces a whole subtree of 9 calls.", VisualHint: "Memo lookup: miss on first visit, hit on repeat", Duration: Duration(2 * time.Second), CodeRef: "fib: if n in memo"},
					{Cue: "compute", Narration: "On a miss we apply the recurrence: fib(n) = fib(n-1) + fib(n-2), and store the result in the memo before returning. fib(2) = 1, fib(3) = 2, fib(4) = 3, fib(5) = 5 are each stored exactly once.", VisualHint: "Store the new value in the memo table", Duration: Duration(3 * time.Second), CodeRef: "fib: memo[n] = ..."},
					{Cue: "return", Narration: "The stack unwinds, and each answer flows back to its caller. fib(6) = 8 after 11 calls. Plain recursion makes 25 calls for fib(6) and over 3 billion for fib(45); the memoized version makes 2n - 1, because each state is computed once and looked up at most once more.", VisualHint: "Stack unwinds; memo holds fib(2)..fib(6)", Duration: Duration(2 * time.Second), CodeRef: "fib: return memo[n]"},
				},
			},
			{
				ID:    "iterative-dp",
				Title: "Iterative DP (Bottom-Up)",
				Goal:  "Fill the Fibonacci table from the base cases up to dp[8], seeing that each cell reads only the two cells before it.",
				MemoryTips: []string{
					"1. Define the state: dp[i] = the i-th Fibonacci number (for stairs: ways to reach step i).",
					"2. Write the recurrence: dp[i] = dp[i-1] + dp[i-2].",
					"3. Set the base cases: dp[0] = 0 and dp[1] = 1, filled before the loop.",
					"4. Choose the evaluation order: i from 2 upward, so both inputs are already final.",
					"5. Locate the answer: dp[n]. 6. Optimise space: only two cells are read, so two variables suffice.",
				},
				Steps: []StoryboardStep{
					{Cue: "init", Narration: "Create a table with one cell per state, dp[0] to dp[8], and fill the base cases by hand: dp[0] = 0, dp[1] = 1. The recurrence reads two cells back, so these are exactly the cells it cannot compute itself.", VisualHint: "dp[0]=0, dp[1]=1 highlighted as base cases", Duration: Duration(2 * time.Second), CodeRef: "iterative: dp[0], dp[1]"},
					{Cue: "fill", Narration: "Fill the remaining cells left to right with dp[i] = dp[i-1] + dp[i-2]: 1, 2, 3, 5, 8, 13, 21. No recursion, no call stack; one addition per cell.", VisualHint: "Cells fill one at a time from left to right", Duration: Duration(3 * time.Second), CodeRef: "iterative: dp[i] = dp[i-1] + dp[i-2]"},
					{Cue: "dependencies", Narration: "Watch the arrows: each new cell reads only the two cells just before it, which are already final. That is why left to right is a valid evaluation order. It also shows that older cells are never read again, which is what lets us replace the table by two rolling variables.", VisualHint: "Arrows from dp[i-1] and dp[i-2] into dp[i]", Duration: Duration(2 * time.Second), CodeRef: "iterative: use previous values"},
					{Cue: "result", Narration: "The answer is the last cell, dp[8] = 21, computed with 7 additions. The memoized recursion computes the same values in the same dependency order; tabulation just decides that order up front.", VisualHint: "dp[8] highlighted as the answer", Duration: Duration(2 * time.Second), CodeRef: "iterative: return dp[n]"},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Climbing Stairs with the six-step recipe",
				Problem: "A staircase has n steps; each move climbs 1 or 2 steps. Count the move sequences that start on the ground (step 0) and end exactly on step n. Trace n = 5.",
				Steps: []WorkedStep{
					{Subgoal: "Define the state", Explanation: "Pick a smaller version of the same question and say it in words: ways[i] = the number of move sequences that end exactly on step i. The original question is ways[n], and the states are i = 0..n.", Code: "ways = [0] * (n + 1)   # ways[i] = sequences ending on step i", State: "n = 5: cells ways[0..5], all unknown"},
					{Subgoal: "Write the recurrence", Explanation: "Look at the last move. A sequence ending on step i either ended with a 1-step (from i-1) or a 2-step (from i-2). Every sequence has exactly one last move, so the two groups are disjoint and cover everything: add them.", Code: "ways[i] = ways[i - 1] + ways[i - 2]", State: "ways[5] = ways[4] + ways[3]"},
					{Subgoal: "Set the base cases", Explanation: "The recurrence reads two cells back, so ways[0] and ways[1] must be given. ways[0] = 1: the empty sequence is one way to be on the ground. ways[1] = 1: a single 1-step. Check: ways[2] = 1 + 1 = 2 (1+1 or 2), which is right.", Code: "ways[0] = 1\nways[1] = 1", State: "ways = [1, 1, ?, ?, ?, ?]"},
					{Subgoal: "Choose the evaluation order", Explanation: "ways[i] reads only smaller indices, so fill i = 2, 3, ..., n. When the loop reaches i, both inputs are already final.", Code: "for i in range(2, n + 1):\n    ways[i] = ways[i - 1] + ways[i - 2]", State: "i=2: 2 | i=3: 3 | i=4: 5 | i=5: 8 -> ways = [1, 1, 2, 3, 5, 8]"},
					{Subgoal: "Locate the answer", Explanation: "The question asked about step n, which is the last cell.", Code: "return ways[n]", State: "ways[5] = 8"},
					{Subgoal: "Optimise space", Explanation: "Computing ways[i] reads only the two previous cells, and nothing reads ways[i-2] afterwards. Keep two variables and roll them forward with a simultaneous assignment: O(1) space instead of O(n).", Code: "two_back, one_back = 1, 1\nfor _ in range(2, n + 1):\n    two_back, one_back = one_back, one_back + two_back\nreturn one_back", State: "(1,1) -> (1,2) -> (2,3) -> (3,5) -> (5,8): answer 8"},
				},
				Code: `def climbStairs(n):
    # Define the state: ways[i] = move sequences ending on step i
    ways = [0] * (n + 1)
    # Set the base cases
    ways[0] = 1
    ways[1] = 1
    # Choose the evaluation order: low steps first
    for i in range(2, n + 1):
        # Write the recurrence: last move was a 1-step or a 2-step
        ways[i] = ways[i - 1] + ways[i - 2]
    # Locate the answer
    return ways[n]`,
			},
			{
				Title:   "One recurrence, four implementations",
				Problem: "Implement ways(n) = ways(n-1) + ways(n-2), ways(0) = ways(1) = 1 four ways, and compare how much work each does for n = 5 and n = 45.",
				Steps: []WorkedStep{
					{Subgoal: "Write the recurrence", Explanation: "Translate the recurrence directly into a recursive function (Chapter 8): base case first, then the two recursive calls. It is correct, but the call tree branches twice per level and repeats the same arguments.", Code: "def climb(n):\n    if n <= 1:\n        return 1\n    return climb(n - 1) + climb(n - 2)", State: "n = 5: 15 calls (climb(2) 3 times, climb(1) 5 times) | n = 45: 3,672,623,805 calls"},
					{Subgoal: "Memoize the states", Explanation: "The result depends only on n, so cache it by n. Check the memo before computing, store before returning. Each state is now computed once (overlapping subproblems removed); this is top-down DP.", Code: "memo = {0: 1, 1: 1}\ndef climb(n):\n    if n not in memo:\n        memo[n] = climb(n - 1) + climb(n - 2)\n    return memo[n]", State: "n = 5: 9 calls, memo {0:1, 1:1, 2:2, 3:3, 4:5, 5:8} | n = 45: 89 calls"},
					{Subgoal: "Choose the evaluation order", Explanation: "Instead of letting recursion discover the order, fill a table from the base cases upward. Every cell reads only earlier cells, so a simple loop works: bottom-up tabulation, with no call stack.", Code: "ways = [0] * (n + 1)\nways[0] = ways[1] = 1\nfor i in range(2, n + 1):\n    ways[i] = ways[i - 1] + ways[i - 2]\nreturn ways[n]", State: "n = 5: 4 additions, table [1, 1, 2, 3, 5, 8] | n = 45: 44 additions"},
					{Subgoal: "Optimise space", Explanation: "Only the last two cells are ever read, so replace the table by two rolling variables. Time stays O(n), space drops to O(1).", Code: "two_back, one_back = 1, 1\nfor _ in range(2, n + 1):\n    two_back, one_back = one_back, one_back + two_back\nreturn one_back", State: "n = 5: (1,1) -> (1,2) -> (2,3) -> (3,5) -> (5,8)"},
				},
				Code: `def climb_rolling(n):
    two_back, one_back = 1, 1      # ways[0], ways[1]
    for _ in range(2, n + 1):
        two_back, one_back = one_back, one_back + two_back
    return one_back`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Min Cost Climbing Stairs",
				Prompt: "cost[i] is the price of stepping off step i; you may start on step 0 or step 1 for free, each move climbs 1 or 2 steps, and the top is step n = len(cost). Fill in the base cases, the recurrence and the answer. Example: cost = [10, 15, 20] gives 15 (start on step 1, pay 15, jump 2 to the top).",
				Code: `def minCostClimbingStairs(cost):
    n = len(cost)
    # Define the state: pay[i] = cheapest total to stand on step i (step n is the top)
    pay = [0] * (n + 1)
    # Set the base cases: you may start on step 0 or step 1 for free
    pay[0] = ____
    pay[1] = ____
    # Choose the evaluation order: low steps first
    for i in range(2, n + 1):
        # Write the recurrence: the last move came from step i-1 or step i-2
        pay[i] = ____
    # Locate the answer
    return ____`,
				Blanks:      []string{"0", "0", "min(pay[i - 1] + cost[i - 1], pay[i - 2] + cost[i - 2])", "pay[n]"},
				Explanation: "Same state shape and evaluation order as Climbing Stairs, but optimising instead of counting. Starting is free, so pay[0] = pay[1] = 0. To stand on step i, the last move left step i-1 (paying cost[i-1]) or step i-2 (paying cost[i-2]); keep the cheaper. For [10, 15, 20]: pay = [0, 0, 10, 15], so the answer is pay[3] = 15.",
			},
			{
				Title:  "N-th Tribonacci number with rolling variables",
				Prompt: "T(0) = 0, T(1) = 1, T(2) = 1 and T(i) = T(i-1) + T(i-2) + T(i-3). The recurrence reads three earlier values, so three variables are enough. Fill in the rolling update and the answer. Example: T(4) = 4, T(25) = 1389537.",
				Code: `def tribonacci(n):
    if n == 0:
        return 0
    if n <= 2:
        return 1
    a, b, c = 0, 1, 1          # T(i-3), T(i-2), T(i-1) for i = 3
    for _ in range(3, n + 1):
        a, b, c = ____
    return ____`,
				Blanks:      []string{"b, c, a + b + c", "c"},
				Explanation: "Optimise space: each step shifts the window by one, so the old T(i-2) becomes the new T(i-3), the old T(i-1) becomes T(i-2), and the new T(i-1) is the sum of all three old values. The tuple assignment computes a + b + c from the old values before any variable changes. After the last pass, c holds T(n).",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt:      "Plain recursion for Climbing Stairs (climb(n) = climb(n-1) + climb(n-2)) makes 15 calls for n = 5. Roughly how many calls does it make for n = 45, and how many different arguments does it see?",
				Choices:     []string{"About 45 calls, 45 arguments", "About 2,000 calls (45 squared), 46 arguments", "About 3.7 billion calls, but only 46 different arguments", "It never finishes, because the recursion has no base case"},
				Answer:      2,
				Explanation: "3,672,623,805 calls, growing by about 1.618 per extra step, but only arguments 0..45 ever occur. The gap between billions of calls and 46 distinct questions is the overlapping-subproblems property, and it is exactly what memoization (89 calls) or a table (44 additions) removes.",
			},
			{
				Prompt:      "Coins [1, 3, 4], amount 6. What does greedy (always take the largest coin that fits) return, and what is the true minimum?",
				Choices:     []string{"Greedy 2, minimum 2: greedy is optimal for any coins", "Greedy 3 (4+1+1), minimum 2 (3+3)", "Greedy 2 (3+3), minimum 2", "Greedy -1, because 6 is not a multiple of 4"},
				Answer:      1,
				Explanation: "Greedy commits to 4 and is left with 2, which needs 1 + 1: three coins. DP tries every last coin: fewest[6] = min(fewest[5] + 1, fewest[3] + 1, fewest[2] + 1) = min(3, 2, 3) = 2. Greedy needs a proof that its choice is always safe; DP needs none, because it checks every choice.",
			},
			{
				Prompt:      "Which pattern fits best? 'Given a string and a dictionary, return the NUMBER of ways to split the string into dictionary words.'",
				Choices:     []string{"Backtracking (Chapter 12): generate every split and count them", "Sliding window (Chapter 4): grow and shrink a window over the string", "1-D DP over prefixes (Chapter 15): ways[i] = sum of ways[j] for each word s[j:i]", "Two pointers (Chapter 3) from both ends"},
				Answer:      2,
				Explanation: "Only a count is asked, and the number of splits can be exponential (for 'aaaa...' with words 'a' and 'aa' it is a Fibonacci number), so listing them with backtracking is too slow. The count for a prefix depends only on counts for shorter prefixes: overlapping subproblems, so DP. Backtracking is right when every split must be returned. A sliding window needs a contiguous range with a condition that grows or shrinks monotonically, which does not describe cutting a string into words.",
			},
			{
				Prompt:      "Which pattern fits each? (a) Longest substring without repeating characters. (b) Longest strictly increasing subsequence.",
				Choices:     []string{"Both sliding window", "(a) sliding window, (b) DP", "(a) DP, (b) sliding window", "Both DP"},
				Answer:      1,
				Explanation: "(a) asks for a contiguous substring whose validity only breaks by adding a repeat and is restored by removing from the left: a sliding window (Chapter 4), O(n). (b) asks for a subsequence, which can skip elements; whether an element extends the answer depends on which earlier element was last, so the state is 'longest ending at i' and the recurrence looks back at every earlier j: DP, O(n^2) (or O(n log n) with binary search).",
			},
			{
				Prompt:      "Prerequisite (Chapter 8, recursion): climb(n) = climb(n-1) + climb(n-2). Which base case makes the recursion correct and terminating for every n >= 1?",
				Choices:     []string{"if n == 1: return 1", "if n <= 1: return 1", "if n == 2: return 2", "if n <= 0: return 0"},
				Answer:      1,
				Explanation: "climb(2) calls climb(1) and climb(0), so both must be base cases: n <= 1 returns 1 covers them. 'n == 1' alone lets climb(0) recurse to climb(-1), climb(-2), ... forever. 'n == 2' alone misses climb(1) -> climb(0) -> negative. 'n <= 0: return 0' terminates but gives climb(2) = climb(1) + climb(0) = 0 + 0 instead of 2. Every recursion needs base cases that every call path reaches, which becomes step 3 of the DP recipe.",
			},
			{
				Prompt:      "Trace this code. What does it return?\n\ndp = [0] * 6\ndp[0] = 1\nfor i in range(1, 6):\n    dp[i] = dp[i - 1] + (dp[i - 3] if i >= 3 else 0)\nreturn dp[5]",
				Choices:     []string{"3", "4", "5", "8"},
				Answer:      1,
				Explanation: "This counts move sequences with steps of 1 or 3. dp fills as [1, 1, 1, 2, 3, 4]: dp[3] = dp[2] + dp[0] = 2, dp[4] = dp[3] + dp[1] = 3, dp[5] = dp[4] + dp[2] = 4 (1+1+1+1+1, 1+1+3, 1+3+1, 3+1+1).",
			},
			{
				Prompt:      "Name the six steps of the DP recipe in order, and say in one sentence what each one decides.",
				Explanation: "1. Define the state: what dp[i] means, in words. 2. Write the recurrence: how dp[i] follows from smaller states, by looking at the last choice. 3. Set the base cases: the smallest states, which the recurrence cannot compute. 4. Choose the evaluation order: an order in which every state is computed after the states it reads. 5. Locate the answer: which cell (dp[n]) or which combination (max over all dp[i]) answers the question. 6. Optimise space: keep only the entries that will still be read.",
			},
			{
				Prompt:      "When is the answer max(dp) rather than dp[n]? Give an example from this chapter's problems.",
				Explanation: "When the state is defined as 'ending exactly at i' instead of 'using the first i elements'. Then dp[n-1] only covers answers that end at the last element, while the best answer can end anywhere. Longest Increasing Subsequence (ends_at[i]) and Maximum Product Subarray (high[i]) both return the maximum over all i; for [1,3,6,7,9,4,10,5,6] the LIS ending at the last element has length 5 but the answer is 6.",
			},
			{
				Prompt:      "A recurrence is dp[i] = max(dp[i-1], dp[i-3] + nums[i-1]). How many variables do you need for a rolling (O(1) space) version, and why?",
				Choices:     []string{"One: only the result is needed", "Two, as in Climbing Stairs", "Three: dp[i-1], dp[i-2] and dp[i-3] must all be kept", "None: rolling only works for sums, not for max"},
				Answer:      2,
				Explanation: "dp[i] reads dp[i-3], so the window must reach three cells back: dp[i-1], dp[i-2] and dp[i-3]. dp[i-2] is not read now, but it becomes dp[i-3] for the next i. Rolling works for any recurrence that reads a fixed number of earlier cells, whether it combines them with +, max, min or or.",
			},
		},
	}
}
