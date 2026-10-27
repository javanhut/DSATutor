package chapter

import "time"

// chapter16 extends the Chapter 15 recipe (Define the state, Write the
// recurrence, Set the base cases, Choose the evaluation order, Locate the
// answer, Optimise space) to states with two indices: grid paths, two-string
// alignment, knapsack (item x budget), state machines and interval DP. The
// 2d-dp practice problems reuse the same subgoal labels.
// Storyboards (IDs registered in app.js): knapsack.
func chapter16() Chapter {
	return Chapter{
		Number:        16,
		Slug:          "2d-dp",
		Title:         "Dynamic Programming on Grids and Strings",
		Category:      "2d-dp",
		Prerequisites: []int{15},
		Summary:       "When one index is not enough to describe a subproblem, give the state two: a cell of a grid, a prefix of each of two strings, an item and a remaining budget, a day and a situation, or the two ends of a range. The Chapter 15 recipe stays the same; the table becomes a grid that you fill row by row (or by range length), and you can draw it, read it and shrink it to one row afterwards. This chapter covers the three big families (grid paths, two-string alignment, knapsack) plus state machines and interval DP.",
		Motivation:    "A diff tool compares two versions of a file and shows which lines were kept. Underneath it is Longest Common Subsequence: the longest list of lines that appears, in order, in both versions. The natural recursion looks at the last line of each version: if they match, keep both; if not, drop one of them, trying both choices. For two 20-line files with nothing in common it makes 275,693,057,639 calls (2 x C(40, 20) - 1), hours of Python, and the files are tiny. Yet every call is described by just two numbers, how many lines of each file are left: (i, j) with 0 <= i, j <= 20, only 441 different questions. Chapter 15 taught you to store each answer once; with two indices the store is a table with a row for each i and a column for each j, and filling it takes 441 steps instead of 275 billion. Two 1,000-line files need about a million cells, well under a second. This chapter is about choosing those two indices, filling the table in an order where every cell's inputs are ready, and recognising the handful of shapes that 2D DP problems come in.",
		Objectives: []string{
			"Define a two-index state in words (what dp[i][j] means, including the extra row and column for empty prefixes) and draw the filled table for a small input",
			"Choose a valid evaluation order (row by row, by interval length, by value with a memoised DFS) and explain why every cell a cell reads is already final",
			"Write the recurrences for the three families: grid paths (above + left), two-string alignment (match, or skip which side) and knapsack (item index x budget)",
			"Explain why a one-row 0/1 knapsack loops the budget backwards while an unbounded one loops forwards, and predict the wrong answer when the direction is swapped",
			"Model a problem as a state machine (a day and a situation) or as an interval with a last choice (Burst Balloons)",
			"Shrink a working 2D table to one or two rows, and say when you should keep the full table instead",
		},
		WhenToUse: []string{
			"Two strings or arrays are compared or combined: 'longest common subsequence', 'fewest edits/deletions to turn one into the other', 'can s3 be formed by interleaving', 'does the pattern match the whole string'",
			"A grid where moves only go right/down (or only to strictly larger values), asking to count routes or find the cheapest one",
			"Items with a budget: choose some numbers or coins so the total is exactly (or at most) X, and count the ways or maximise the value. The state is (item index, remaining budget)",
			"A sequence of decisions where what you may do next depends on a small 'mode' (holding a stock, in cooldown, transactions left): add the mode as a second index",
			"Removing or merging pieces of a row where each step's cost depends on the current neighbours (bursting balloons, multiplying a chain of matrices): an interval state (left, right)",
			"The answer is a count, a min/max or a yes/no rather than a list of every solution, and the two dimensions multiply to at most a few million cells",
		},
		Concepts: []Concept{
			{
				Name:        "The 2D table: what a cell means and the order you fill it",
				Description: "A two-index state is a table with one cell per pair of indices. Before writing code, say in one sentence what a single cell means, then check that every cell reads only cells that are already filled.",
				WhyItWorks:  "Nothing new is needed beyond Chapter 15's induction: a cell is correct if the cells it reads are correct and the recurrence lists every possible last choice. What is new is the evaluation order. In a grid where the robot moves only right or down, cell (r, c) is reached from (r - 1, c) or (r, c - 1); both are in an earlier row or earlier in the same row, so filling row by row, left to right, always has them ready. Draw arrows from each cell to the cells it reads: if every arrow points up, left or up-left, row-major order works. If some arrow points down or right (as in Longest Increasing Path, where a cell reads whichever neighbours are larger), row order is wrong and you need another order: by value, by interval length, or a memoised DFS that discovers it. The cost of a 2D DP is (number of cells) x (work per cell): m x n cells with O(1) work each is O(m * n).",
				Intuition:   "A table is a spreadsheet where each cell has a formula that refers to other cells. A spreadsheet recalculates in an order where every referenced cell is computed first, and it reports an error if two cells refer to each other (a cycle). Row-by-row filling is that recalculation order done by hand: the formula in B3 may use B2 and A3, which are already there when you reach B3.",
				CoreIdeas: []string{
					"State in words first: 'paths[r][c] = number of routes from the start to (r, c)'. The units (routes, letters, coins, True/False) tell you whether to add, max/min or 'or'.",
					"Index 0 often means 'empty': a table of (m + 1) x (n + 1) for prefixes, with row 0 and column 0 as base cases, keeps the recurrence free of special cases.",
					"Valid order = every arrow from a cell to the cells it reads points to something already filled. Up/left/up-left arrows -> row by row. Shorter ranges -> by length. Larger values -> sort by value or memoised DFS.",
					"Draw the table for a 3 x 3 example by hand before coding; most bugs show up as one wrong edge row or column.",
				},
				CommonMistakes: []string{
					"Reading paths[r - 1][c] when r = 0. Python's paths[-1] is the LAST row, so there is no error, just a wrong number. Fill the first row and column as base cases (or guard with r > 0) and start the loops at 1.",
					"Building the table with [[0] * n] * m. That makes m references to the SAME row, so writing one cell changes the whole column. Use [[0] * n for _ in range(m)].",
					"Filling in an order that reads unfilled cells, e.g. row by row in Burst Balloons. The result is silently too small. Draw the dependency arrows and check the order before trusting the output.",
				},
				Examples: []CodeExample{
					{
						ID:       "grid-paths-with-blocks",
						Title:    "Grid paths with blocked cells (row by row)",
						Language: "python",
						Snippet: `def paths_with_blocks(grid):          # grid[r][c] == 1 means blocked
    rows, cols = len(grid), len(grid[0])
    # paths[r][c] = number of right/down routes from (0, 0) to (r, c)
    paths = [[0] * cols for _ in range(rows)]
    for r in range(rows):                 # row by row: above and left are ready
        for c in range(cols):
            if grid[r][c] == 1:
                paths[r][c] = 0           # no route ends on a wall
            elif r == 0 and c == 0:
                paths[r][c] = 1           # the start
            else:
                above = paths[r - 1][c] if r > 0 else 0
                left = paths[r][c - 1] if c > 0 else 0
                paths[r][c] = above + left
    return paths[rows - 1][cols - 1]

# [[0,0,0],[0,1,0],[0,0,0]] -> table [[1,1,1],[1,0,1],[1,1,2]] -> 2`,
						Notes: "Unique Paths II. The guards r > 0 and c > 0 replace separate base-case loops; a blocked cell is simply a cell with 0 routes.",
					},
				},
			},
			{
				Name:        "Two strings: match, or skip which side",
				Description: "When two sequences are compared, the state is a pair of prefix lengths (i, j). The recurrence looks at the last letter of each prefix: if they match, both are used (move diagonally); if not, one of them is skipped, and you try each side.",
				WhyItWorks:  "Every alignment of a[:i] with b[:j] ends in one of three ways: the last letters a[i - 1] and b[j - 1] are paired (diagonal cell (i - 1, j - 1)), a[i - 1] is left unpaired (cell above, (i - 1, j)), or b[j - 1] is left unpaired (cell to the left, (i, j - 1)). That covers every alignment, so combining those three cells (max for LCS, min-plus-cost for Edit Distance, sum for counting, 'or' for yes/no) is complete. The specific problems just decide which of the three moves are allowed and what they cost: LCS allows the diagonal only on a match; Edit Distance also allows it on a mismatch at cost 1 (replace); Distinct Subsequences never lets the second string skip. The base cases come from the empty prefixes: row 0 and column 0 say what happens when one string is used up (0 for LCS, i or j edits for Edit Distance, 1 way to spell the empty string).",
				Intuition:   "Two people read two lists of words aloud, line by line, trying to keep in step. At each moment either both say the same word (they move on together: diagonal), or one of them skips a word and the other waits (up or left). An alignment is a path through the table from the top-left corner to the bottom-right, made only of those three kinds of steps; the DP finds the best path without listing them all.",
				CoreIdeas: []string{
					"State: dp[i][j] about the prefixes a[:i] and b[:j]; the last letters are a[i - 1] and b[j - 1] (watch the minus one).",
					"Diagonal = use both last letters; up = skip a's letter; left = skip b's letter.",
					"Max for 'longest', min for 'fewest edits', + for 'how many ways', or for 'is it possible'.",
					"Row 0 and column 0 are the 'one string is empty' answers, and they differ by problem: decide them from the state's meaning, not by habit.",
				},
				CommonMistakes: []string{
					"Comparing a[i] with b[j] in a table that has an extra empty row and column. Cell (i, j) is about the first i letters, whose last letter is a[i - 1]. Write the state as a comment and the indexing follows.",
					"On a mismatch, moving only diagonally (skipping both letters). That throws away letters that might match later: for LCS of 'abc' and 'cab' it returns 0 instead of 2. A mismatch means skip ONE side, and you try both.",
					"Copying the base cases from another problem: LCS's row 0 is all 0s, but Edit Distance's row 0 is 0, 1, 2, ..., because turning '' into 'abc' takes 3 inserts.",
				},
				Examples: []CodeExample{
					{
						ID:       "shortest-common-supersequence-length",
						Title:    "Shortest common supersequence length",
						Language: "python",
						Snippet: `def scs_length(a, b):
    m, n = len(a), len(b)
    # scs[i][j] = length of the shortest string that has both a[:i] and b[:j] as subsequences
    scs = [[0] * (n + 1) for _ in range(m + 1)]
    for i in range(m + 1):
        scs[i][0] = i                     # only a's letters left
    for j in range(n + 1):
        scs[0][j] = j                     # only b's letters left
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if a[i - 1] == b[j - 1]:
                scs[i][j] = scs[i - 1][j - 1] + 1      # match: write the letter once
            else:
                scs[i][j] = 1 + min(scs[i - 1][j],     # write a's letter
                                    scs[i][j - 1])     # write b's letter
    return scs[m][n]

# scs_length("abac", "cab") -> 5 (for example "cabac")`,
						Notes: "Same three neighbours as LCS and Edit Distance; only the costs and the edges change.",
					},
				},
			},
			{
				Name:        "Knapsack: item index x budget, and which way the inner loop runs",
				Description: "When you choose a subset (or a multiset) of items under a budget, the state is (how many items have been considered, budget used or left). 0/1 knapsack allows each item once; unbounded knapsack allows unlimited copies. In the 2D table the difference is one index; in a one-row version it is the direction of the inner loop.",
				WhyItWorks:  "Let best[i][w] = the best value using only the first i items with capacity w. Item i (weight wt, value v) is either skipped, best[i - 1][w], or taken. In 0/1 knapsack, after taking it you may not take it again, so the rest comes from the PREVIOUS row: best[i - 1][w - wt] + v. In unbounded knapsack you may take it again, so the rest comes from the SAME row: best[i][w - wt] + v. Now shrink the table to one row, updated in place. When you are about to write best[w], the cells to its left (smaller w) have already been updated if the loop runs forwards, and have not if it runs backwards. So forwards reads the same row (unbounded), and backwards reads the previous row (0/1). With one item of weight 2 and value 3 and capacity 4: forwards gives best = [0, 0, 3, 3, 6], where 6 means the item was taken twice; backwards gives [0, 0, 3, 3, 3]. Coin Change II (unlimited coins) loops forwards; Target Sum and Partition Equal Subset Sum (each number once) loop backwards.",
				Intuition:   "Packing a suitcase from a list of items. 0/1 knapsack is packing souvenirs: each one is unique, so once an item is in, it is crossed off the list (you look at the list as it was BEFORE this item: the previous row). Unbounded knapsack is packing from a shop shelf: after putting a T-shirt in, the same T-shirt is still on the shelf (you look at the list as it is NOW, including this item: the same row).",
				CoreIdeas: []string{
					"State: (items considered so far, budget). Rows = items, columns = budget 0..W.",
					"0/1: take reads the previous row, best[i - 1][w - wt]. One-row version: budget loop BACKWARDS.",
					"Unbounded: take reads the same row, best[i][w - wt]. One-row version: budget loop FORWARDS.",
					"Counting versions use + instead of max, and the base case is 'budget 0 is reached in exactly 1 way' (the empty choice).",
					"Putting the items in the outer loop counts combinations (each multiset once); putting the budget outside counts ordered sequences.",
				},
				CommonMistakes: []string{
					"Copying a forwards loop into a 0/1 problem. Items get reused silently: for Target Sum with [1, 1, 1, 1, 1] and target 3 the forwards loop returns 70 instead of 5. Ask 'may this item be used again?' and pick the direction from the answer.",
					"Swapping the loops in a counting problem (budget outside, items inside). That counts 1 + 2 and 2 + 1 as different, so Coin Change II with amount 4 and coins [1, 2] returns 5 instead of 3.",
					"Shrinking to one row before the 2D version works. The table makes the previous-row versus same-row read visible; once it is right, the loop direction follows from it.",
				},
				Examples: []CodeExample{
					{
						ID:       "knapsack-loop-direction",
						Title:    "0/1 versus unbounded knapsack in one row",
						Language: "python",
						Snippet: `def best_value_01(weights, values, capacity):          # each item at most once
    best = [0] * (capacity + 1)                         # best[w] = most value within weight w
    for wt, v in zip(weights, values):
        for w in range(capacity, wt - 1, -1):           # BACKWARDS: best[w - wt] is still the previous row
            best[w] = max(best[w], best[w - wt] + v)
    return best[capacity]

def best_value_unbounded(weights, values, capacity):   # unlimited copies of each item
    best = [0] * (capacity + 1)
    for wt, v in zip(weights, values):
        for w in range(wt, capacity + 1):               # FORWARDS: best[w - wt] may already include this item
            best[w] = max(best[w], best[w - wt] + v)
    return best[capacity]

# weights [2, 3, 4, 5], values [3, 4, 5, 6], capacity 8:
# 0/1 -> 10 (items of weight 3 and 5), unbounded -> 12 (four items of weight 2)`,
						Notes: "The two functions differ only in the direction of the inner loop.",
					},
				},
			},
			{
				Name:        "State machines: a second index for the situation you are in",
				Description: "Sometimes the second index is not a position in a second input but a small set of named situations: holding a stock or not, in a cooldown or not, how many transactions are left. The table has one row per step and one column per situation, and each column lists the ways to arrive in that situation from yesterday's situations.",
				WhyItWorks:  "DP needs the state to contain everything the future depends on. In Best Time to Buy and Sell Stock with Cooldown, whether you may buy today depends on whether you sold yesterday, which the day number alone doesn't tell you. Adding the situation (HOLD, SOLD, REST) restores the property: two histories that end day i in the same situation have the same possible futures, so only the one with the most cash matters (optimal substructure). The recurrence is then a list of allowed transitions, like a little diagram: REST -> HOLD (buy), HOLD -> SOLD (sell), SOLD -> REST (cooldown), and each situation may also stay as it is where the rules allow. Each column of row i takes the max over the arrows that point into it from row i - 1.",
				Intuition:   "A traffic light has a few states (red, green, amber) and fixed rules for which state can follow which. To plan a drive you don't need the whole history of the light, only its current state. A state-machine DP keeps, for every day and every state of the 'light', the best result so far, and moves forward one day at a time along the allowed arrows.",
				CoreIdeas: []string{
					"Draw the situations as circles and the allowed moves as arrows before writing code; each arrow is one term in a max.",
					"Impossible situations (sold before ever buying) start at -infinity, so a max never picks them.",
					"The rows only read the previous row, so the table always shrinks to one variable per situation.",
				},
				CommonMistakes: []string{
					"Too few situations: merging 'just sold' with 'resting' lets you buy during the cooldown. If the rules for tomorrow differ, the situations must be separate columns.",
					"Updating the per-situation variables one after another so that later lines read today's values. Compute every situation from yesterday's values (a table row, or a simultaneous assignment).",
				},
				Examples: []CodeExample{
					{
						ID:       "stock-with-fee-state-machine",
						Title:    "Stock trading with a fee: two situations",
						Language: "python",
						Snippet: `def max_profit_with_fee(prices, fee):
    # hold = most cash while owning a share; free = most cash while owning nothing
    hold, free = -prices[0], 0
    for price in prices[1:]:
        hold, free = (max(hold, free - price),         # keep holding, or buy today
                      max(free, hold + price - fee))   # stay out, or sell today and pay the fee
    return free

# max_profit_with_fee([1, 3, 2, 8, 4, 9], 2) -> 8 (buy 1, sell 8, buy 4, sell 9)`,
						Notes: "Two situations and four arrows. The cooldown version adds a third situation (just sold).",
					},
				},
			},
			{
				Name:        "Other fill orders: intervals by length, cells by value",
				Description: "Two families don't fill row by row. Interval DP has a state (left, right) for a range and fills short ranges before long ones; the key move is to choose the LAST action inside the range. Memoised DFS handles tables where a cell can depend on neighbours in any direction, as long as the dependencies have no cycle.",
				WhyItWorks:  "Burst Balloons: if you pick the balloon k that is burst FIRST, its neighbours become adjacent and the two sides interact, so they are not independent subproblems. If you pick the balloon k burst LAST in the range, it stands as a wall while everything else is burst: the left part only ever sees left and k as its outer neighbours, the right part only k and right, and k itself finally earns vals[left] * vals[k] * vals[right]. So best[left][right] = max over k of best[left][k] + vals[left] * vals[k] * vals[right] + best[k][right]. Both subranges are shorter, so filling by increasing length is a valid order. Longest Increasing Path in a Matrix: longest[r][c] = 1 + max of longest at the strictly larger neighbours. Those neighbours can be in any direction, but values strictly increase along every dependency, so there is no cycle (it is a DAG, Chapter 13). A DFS that computes the neighbours first and stores every result visits each cell once: a valid order found automatically.",
				Intuition:   "Planning the order of speeches at a dinner where each speaker's applause depends on who is sitting next to them at the time. Deciding who speaks first leaves a messy table. Deciding who speaks last splits the table into a left half and a right half that never interact, each with the last speaker as its fixed neighbour, which is a much smaller version of the same question.",
				CoreIdeas: []string{
					"Interval state: dp[left][right] for the range between two ends; fill by gap = right - left from small to large.",
					"Choose the LAST action in the range so it splits the range into two independent parts.",
					"Padding with neutral values (1s around the balloons) removes edge cases.",
					"If dependencies go in every direction but can't form a cycle, memoise a DFS (or sort the cells into a valid order) instead of looping row by row.",
				},
				CommonMistakes: []string{
					"Choosing the first action in interval DP. The parts are not independent afterwards and the formula gives wrong coins. Ask which action happens LAST.",
					"Filling an interval table row by row. best[k][right] for k > left is in a later row and still 0; the answer comes out too small (63 instead of 167 for [3, 1, 5, 8]). Fill by gap.",
					"Memoising a DFS on a graph that has cycles (moving to equal or smaller values). The memo then stores half-finished answers. Memoised DFS is only correct when the dependencies form a DAG.",
				},
				Examples: []CodeExample{
					{
						ID:       "matrix-chain-interval-dp",
						Title:    "Matrix chain multiplication (interval DP, last split)",
						Language: "python",
						Snippet: `def min_multiplications(dims):
    # matrix k has shape dims[k] x dims[k + 1]; cost[left][right] = cheapest way to
    # multiply the matrices between boundaries left and right
    n = len(dims)
    cost = [[0] * n for _ in range(n)]          # gap 1: a single matrix, nothing to do
    for gap in range(2, n):                     # short chains first
        for left in range(n - gap):
            right = left + gap
            cost[left][right] = min(
                cost[left][k] + dims[left] * dims[k] * dims[right] + cost[k][right]
                for k in range(left + 1, right)  # k = the boundary of the LAST multiplication
            )
    return cost[0][n - 1]

# min_multiplications([10, 20, 30, 40]) -> 18000: (AB)C costs 6000 + 12000`,
						Notes: "The same shape as Burst Balloons: a range, a last choice k that splits it, and a fill order by length.",
					},
				},
			},
		},
		Animations: []Storyboard{
			{
				ID:    "knapsack",
				Title: "0/1 Knapsack: filling the item x capacity table",
				Goal:  "Watch the table best[i][w] fill row by row for items A (weight 2, value 3), B (3, 4), C (4, 5), D (5, 6) and capacity 8, and see why each cell reads only the row above.",
				MemoryTips: []string{
					"Define the state: best[i][w] = most value using only the first i items with capacity w.",
					"Set the base cases: row 0 (no items) and column 0 (no capacity) are 0.",
					"Write the recurrence: skip the item (cell above) or take it (row above, w - weight columns left, plus its value); keep the max.",
					"Choose the evaluation order: one row per item; every read is in the row above, so the order within a row doesn't matter.",
					"Locate the answer: the bottom-right cell; walk back up to find which items were taken.",
					"Optimise space: one row, capacity loop BACKWARDS so the row above is still there when you read it.",
				},
				Steps: []StoryboardStep{
					{Cue: "init", Narration: "Define the state and set the base cases. The table has a row for 'first i items' (i = 0..4) and a column for every capacity 0..8. Row 0 means no items are allowed, so every cell is 0; column 0 means no room, also 0.", VisualHint: "Empty 5 x 9 grid; row 0 and column 0 shaded as base cases", Duration: Duration(3 * time.Second), CodeRef: "knapsack: init dp"},
					{Cue: "consider", Narration: "Row 1 adds item A (weight 2, value 3). Each row reads only the finished row above it, so we can fill it left to right. Capacities 0 and 1 can't fit A, so they copy the cell above: 0.", VisualHint: "Highlight item A and row 1; arrows from row 0 down to row 1", Duration: Duration(3 * time.Second), CodeRef: "knapsack: for each item"},
					{Cue: "decide", Narration: "Write the recurrence. At row 3 (item C, weight 4, value 5), capacity 8: skip C gives the cell above, best[2][8] = 7; take C gives best[2][8 - 4] + 5 = 4 + 5 = 9. The max is 9. Taking reads the PREVIOUS row, so C can't be used twice.", VisualHint: "Two arrows into cell (3, 8): one straight down from (2, 8) = 7, one from (2, 4) = 4 plus 5", Duration: Duration(4 * time.Second), CodeRef: "knapsack: max(exclude, include)"},
					{Cue: "fill", Narration: "Filling every cell this way gives rows A: 0 0 3 3 3 3 3 3 3, B: 0 0 3 4 4 7 7 7 7, C: 0 0 3 4 5 7 8 9 9, D: 0 0 3 4 5 7 8 9 10. Locate the answer: the bottom-right cell, 10.", VisualHint: "Table fills row by row; bottom-right cell glows with 10", Duration: Duration(3 * time.Second), CodeRef: "knapsack: dp[i][w] = ..."},
					{Cue: "traceback", Narration: "Walk back from (4, 8). 10 differs from the cell above (9), so D was taken: move to (3, 8 - 5) = (3, 3). There 4 equals the cell above, so C was skipped. At (2, 3) the 4 differs from the 3 above it, so B was taken, leaving capacity 0. Items B and D: weight 8, value 10. This walk-back is why it pays to keep the full table.", VisualHint: "Path from bottom-right up through the table; taken items D and B highlighted", Duration: Duration(4 * time.Second), CodeRef: "knapsack: traceback"},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Minimum Path Sum: a grid table filled row by row",
				Problem: "A grid of non-negative costs; start at the top-left, finish at the bottom-right, and move only right or down. Return the smallest possible sum of the costs on the route (both ends included). Trace grid = [[1,3,1],[1,5,1],[4,2,1]], whose answer is 7 (1 -> 3 -> 1 -> 1 -> 1).",
				Steps: []WorkedStep{
					{Subgoal: "Define the state", Explanation: "Pick the smaller question: the cheapest route not to the finish but to any cell. cost[r][c] = the smallest sum of a right/down route from (0, 0) to (r, c), including both cells. The table has the same shape as the grid.", Code: "cost = [[0] * cols for _ in range(rows)]   # cost[r][c] = cheapest route sum to (r, c)", State: "3 x 3 table, every cell unknown"},
					{Subgoal: "Write the recurrence", Explanation: "A route into (r, c) arrives from above or from the left, and whichever it was, the part before it should be the cheapest route to that neighbour. So take the cheaper neighbour and add this cell's cost. It is min (not +) because we optimise instead of counting.", Code: "cost[r][c] = grid[r][c] + min(cost[r - 1][c], cost[r][c - 1])", State: "cell (1, 1): 5 + min(cost[0][1], cost[1][0]) = 5 + min(4, 2) = 7"},
					{Subgoal: "Set the base cases", Explanation: "The recurrence needs a cell above and a cell to the left, which the top row and the left column don't have. The top row can only be reached by moving right, so it is a running sum; the same for the left column moving down.", Code: "cost[0][0] = grid[0][0]\nfor c in range(1, cols):\n    cost[0][c] = cost[0][c - 1] + grid[0][c]\nfor r in range(1, rows):\n    cost[r][0] = cost[r - 1][0] + grid[r][0]", State: "top row: 1 4 5 | left column: 1, 2, 6"},
					{Subgoal: "Choose the evaluation order", Explanation: "Every cell reads the cell above and the cell to its left, so row by row, left to right, always has both ready. Here is the table after each row is filled.", Code: "for r in range(1, rows):\n    for c in range(1, cols):\n        cost[r][c] = grid[r][c] + min(cost[r - 1][c], cost[r][c - 1])", State: "after row 0:  1 4 5\nafter row 1:  1 4 5 / 2 7 6\nafter row 2:  1 4 5 / 2 7 6 / 6 8 7\n\n      c=0 c=1 c=2\nr=0    1   4   5\nr=1    2   7   6\nr=2    6   8   7"},
					{Subgoal: "Locate the answer", Explanation: "The finish is the bottom-right cell. To recover the route, walk back from it, each time to the neighbour (above or left) with the smaller value: 7 <- 6 <- 5 <- 4 <- 1, i.e. right, right, down, down.", Code: "return cost[rows - 1][cols - 1]", State: "cost[2][2] = 7"},
					{Subgoal: "Optimise space", Explanation: "Each row reads only the row above, so one row updated in place is enough: before best[c] is overwritten it is the cell above; best[c - 1] is already the cell to the left in this row. Do this only after the table version works, and keep the table if you need the route.", Code: "best = [0] * cols\nfor r in range(rows):\n    for c in range(cols):\n        if r == 0 and c == 0:\n            best[c] = grid[0][0]\n        elif r == 0:\n            best[c] = best[c - 1] + grid[r][c]\n        elif c == 0:\n            best[c] = best[c] + grid[r][c]\n        else:\n            best[c] = grid[r][c] + min(best[c], best[c - 1])\nreturn best[cols - 1]", State: "row 0: [1, 4, 5] -> row 1: [2, 7, 6] -> row 2: [6, 8, 7]: answer 7"},
				},
				Code: `def minPathSum(grid):
    rows, cols = len(grid), len(grid[0])
    # Define the state: cost[r][c] = cheapest right/down route sum from (0, 0) to (r, c)
    cost = [[0] * cols for _ in range(rows)]
    # Set the base cases: the top row and left column have only one route
    cost[0][0] = grid[0][0]
    for c in range(1, cols):
        cost[0][c] = cost[0][c - 1] + grid[0][c]
    for r in range(1, rows):
        cost[r][0] = cost[r - 1][0] + grid[r][0]
    # Choose the evaluation order: row by row, left to right
    for r in range(1, rows):
        for c in range(1, cols):
            # Write the recurrence: arrive from above or from the left
            cost[r][c] = grid[r][c] + min(cost[r - 1][c], cost[r][c - 1])
    # Locate the answer
    return cost[rows - 1][cols - 1]`,
			},
			{
				Title:   "Fewest deletions to make two strings equal: match, or skip which side",
				Problem: "Given word1 and word2, delete letters from either string (one letter per step) until the two strings are equal. Return the fewest deletions. Trace word1 = \"sea\", word2 = \"eat\", whose answer is 2 (delete s from \"sea\" and t from \"eat\", leaving \"ea\").",
				Steps: []WorkedStep{
					{Subgoal: "Define the state", Explanation: "The state is a pair of prefixes, one from each word. dels[i][j] = fewest deletions to make word1[:i] and word2[:j] equal. An extra row and column (index 0) stand for the empty prefix, so the table is (m + 1) x (n + 1).", Code: "dels = [[0] * (n + 1) for _ in range(m + 1)]   # dels[i][j] for word1[:i], word2[:j]", State: "word1 = \"sea\" (m = 3), word2 = \"eat\" (n = 3): a 4 x 4 table"},
					{Subgoal: "Write the recurrence", Explanation: "Look at the last letters word1[i - 1] and word2[j - 1]. If they match, keep both: no cost, and the rest is the diagonal cell. If not, at least one of them must be deleted: delete word1's (cell above) or word2's (cell to the left), pay 1, and keep the cheaper. Match, or skip which side.", Code: "if word1[i - 1] == word2[j - 1]:\n    dels[i][j] = dels[i - 1][j - 1]\nelse:\n    dels[i][j] = 1 + min(dels[i - 1][j], dels[i][j - 1])", State: "cell (2, 1): 'e' == 'e' -> dels[1][0] = 1 | cell (1, 1): 's' != 'e' -> 1 + min(dels[0][1], dels[1][0]) = 1 + 1 = 2"},
					{Subgoal: "Set the base cases", Explanation: "If one prefix is empty, every letter of the other must be deleted: dels[i][0] = i and dels[0][j] = j. These edges come straight from the state's meaning (they are not all 0, unlike LCS).", Code: "for i in range(m + 1):\n    dels[i][0] = i\nfor j in range(n + 1):\n    dels[0][j] = j", State: "row 0: 0 1 2 3 | column 0: 0 1 2 3"},
					{Subgoal: "Choose the evaluation order", Explanation: "Each cell reads up, left and up-left, so fill row by row, left to right. The full table for \"sea\" vs \"eat\":", Code: "for i in range(1, m + 1):\n    for j in range(1, n + 1):\n        ...", State: "        \"\"  e  a  t\n  \"\"     0  1  2  3\n  s      1  2  3  4\n  e      2  1  2  3\n  a      3  2  1  2"},
					{Subgoal: "Locate the answer", Explanation: "Both full words: the bottom-right cell. Walking back, (3, 3) 't' vs 'a' came from the left (delete t), (3, 2) 'a' == 'a' diagonal, (2, 1) 'e' == 'e' diagonal, (1, 0) delete s.", Code: "return dels[m][n]", State: "dels[3][3] = 2"},
				},
				Code: `def minDeletions(word1, word2):
    m, n = len(word1), len(word2)
    # Define the state: dels[i][j] = fewest deletions to make word1[:i] and word2[:j] equal
    dels = [[0] * (n + 1) for _ in range(m + 1)]
    # Set the base cases: against an empty string, delete everything
    for i in range(m + 1):
        dels[i][0] = i
    for j in range(n + 1):
        dels[0][j] = j
    # Choose the evaluation order: row by row
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            # Write the recurrence: match, or skip (delete) which side
            if word1[i - 1] == word2[j - 1]:
                dels[i][j] = dels[i - 1][j - 1]
            else:
                dels[i][j] = 1 + min(dels[i - 1][j], dels[i][j - 1])
    # Locate the answer
    return dels[m][n]`,
			},
			{
				Title:   "0/1 knapsack: item index x capacity, then one row backwards",
				Problem: "Items A (weight 2, value 3), B (3, 4), C (4, 5), D (5, 6); a bag holds total weight 8; each item can be packed at most once. Return the largest total value. The answer is 10 (B and D).",
				Steps: []WorkedStep{
					{Subgoal: "Define the state", Explanation: "Two things describe a subproblem: which items we may still use and how much room is left. best[i][w] = the most value using only the first i items with capacity w. Rows 0..4, columns 0..8.", Code: "best = [[0] * (capacity + 1) for _ in range(n + 1)]", State: "5 x 9 table"},
					{Subgoal: "Write the recurrence", Explanation: "Item i is either skipped (the cell above) or taken, if it fits: its value plus the best for the remaining capacity using only the EARLIER items (previous row), because it can't be packed twice.", Code: "best[i][w] = best[i - 1][w]\nif wt <= w:\n    best[i][w] = max(best[i][w], best[i - 1][w - wt] + v)", State: "C at w = 8: max(best[2][8], best[2][4] + 5) = max(7, 4 + 5) = 9"},
					{Subgoal: "Set the base cases", Explanation: "With no items (row 0) or no capacity (column 0) the value is 0, which the zero-filled table already holds.", Code: "# row 0 and column 0 stay 0", State: "row 0: 0 0 0 0 0 0 0 0 0"},
					{Subgoal: "Choose the evaluation order", Explanation: "Row by row: every read is in the row above. The finished table:", Code: "for i in range(1, n + 1):\n    wt, v = weights[i - 1], values[i - 1]\n    for w in range(capacity + 1):\n        ...", State: "      w: 0 1 2 3 4 5 6 7 8\n none:   0 0 0 0 0 0 0 0 0\n +A:     0 0 3 3 3 3 3 3 3\n +B:     0 0 3 4 4 7 7 7 7\n +C:     0 0 3 4 5 7 8 9 9\n +D:     0 0 3 4 5 7 8 9 10"},
					{Subgoal: "Locate the answer", Explanation: "All items, full capacity: best[4][8] = 10. To find the items, walk back up: a cell that differs from the cell above means that row's item was taken. best[4][8] = 10 differs from best[3][8] = 9, so D was taken; move to capacity 8 - 5 = 3. best[3][3] = 4 equals best[2][3] = 4, so C was skipped. best[2][3] = 4 differs from best[1][3] = 3, so B was taken; capacity 0 is left. Items B and D.", Code: "return best[n][capacity]", State: "best[4][8] = 10 = B (4) + D (6)"},
					{Subgoal: "Optimise space", Explanation: "Keep one row. 'Take' must read the previous row at a SMALLER capacity, so run w backwards: the cells to the left haven't been overwritten yet. Forwards would read cells already updated for this item and pack it twice. Check with item A alone and capacity 4: forwards gives 6 (two A's), backwards gives 3.", Code: "best = [0] * (capacity + 1)\nfor wt, v in zip(weights, values):\n    for w in range(capacity, wt - 1, -1):\n        best[w] = max(best[w], best[w - wt] + v)\nreturn best[capacity]", State: "A only, capacity 4 | forwards: [0, 0, 3, 3, 6] (wrong: A twice) | backwards: [0, 0, 3, 3, 3]"},
				},
				Code: `def knapsack(weights, values, capacity):
    n = len(weights)
    # Define the state: best[i][w] = most value from the first i items within capacity w
    best = [[0] * (capacity + 1) for _ in range(n + 1)]   # base cases: row 0 and column 0 are 0
    # Choose the evaluation order: one row per item
    for i in range(1, n + 1):
        wt, v = weights[i - 1], values[i - 1]
        for w in range(capacity + 1):
            # Write the recurrence: skip the item, or take it once (previous row)
            best[i][w] = best[i - 1][w]
            if wt <= w:
                best[i][w] = max(best[i][w], best[i - 1][w - wt] + v)
    # Locate the answer
    return best[n][capacity]`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Longest common substring (contiguous)",
				Prompt: "Return the length of the longest string that appears as a contiguous block in both a and b. Unlike a subsequence, the letters must be next to each other, so a mismatch breaks the run. Fill in the recurrence for a match and the line that records the best run. Example: a = \"abcde\", b = \"xbcdy\" gives 3 (\"bcd\").",
				Code: `def longest_common_substring(a, b):
    m, n = len(a), len(b)
    # Define the state: run[i][j] = length of the common block that ENDS at a[i - 1] and b[j - 1]
    run = [[0] * (n + 1) for _ in range(m + 1)]   # base cases: empty prefixes, run 0
    best = 0
    # Choose the evaluation order: row by row
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if a[i - 1] == b[j - 1]:
                # Write the recurrence: extend the block that ended one letter earlier in both
                run[i][j] = ____
                # Locate the answer: the best block can end anywhere
                best = ____
            # on a mismatch the block is broken: run[i][j] stays 0
    return best`,
				Blanks:      []string{"run[i - 1][j - 1] + 1", "max(best, run[i][j])"},
				Explanation: "Same table as LCS, but the state means 'ending exactly here', so a match can only extend the diagonal (the block must continue in both strings) and a mismatch resets to 0 instead of skipping one side. Because a block can end anywhere, the answer is the maximum over all cells, not the bottom-right cell, the same 'max over states' answer as Chapter 15's Longest Increasing Subsequence. For \"abcde\" and \"xbcdy\": run[2][2] = 1 (b), run[3][3] = 2 (bc), run[4][4] = 3 (bcd).",
			},
			{
				Title:  "Unbounded knapsack in one row",
				Prompt: "Each item can be packed any number of times. Fill in the capacity loop (pick the direction!) and the 'take one more' value. Example: weights [2, 3, 4, 5], values [3, 4, 5, 6], capacity 8 gives 12 (four copies of the weight-2 item).",
				Code: `def unbounded_knapsack(weights, values, capacity):
    # Define the state: best[w] = most value within weight w, using the items seen so far
    best = [0] * (capacity + 1)                 # base cases: no items, value 0
    for i in range(len(weights)):
        # Choose the evaluation order: may this item be used again?
        for w in ____:
            # Write the recurrence: skip it, or take one more copy
            best[w] = max(best[w], ____)
    # Locate the answer
    return best[capacity]`,
				Blanks:      []string{"range(weights[i], capacity + 1)", "best[w - weights[i]] + values[i]"},
				Explanation: "Items may be reused, so 'take one more' must read the CURRENT row at w - weights[i]: in the 2D table that is best[i][w - weight]. Looping w forwards does exactly that, because the smaller capacities have already been updated for this item. Starting at weights[i] skips capacities the item doesn't fit in. A backwards loop would read the previous row and give the 0/1 answer, 10, instead of 12.",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Which pattern fits best? 'Given two strings, return the fewest letters you must delete (from either string) so that they become equal.'",
				Choices: []string{
					"A sliding window over the longer string (Chapter 4)",
					"Two pointers, one in each string, advancing together when the letters are equal (Chapter 3)",
					"A 2D DP table over pairs of prefixes: match, or skip which side (this chapter)",
					"A 1D DP over the positions of one string (Chapter 15)",
				},
				Answer:      2,
				Explanation: "Deletions can happen anywhere in either string, and whether to delete a letter depends on letters that come later in both strings, so the state needs a position in each: (i, j). Two pointers commit greedily to the first equal pair and fail on 'abc' vs 'cab'; a sliding window needs a contiguous block; one index can't remember how far you are in the other string.",
			},
			{
				Prompt: "Which technique fits? 'A maze is a grid of open and wall cells. You may move up, down, left or right. Return the fewest moves from the start to the exit.'",
				Choices: []string{
					"A 2D DP table filled row by row, like Minimum Path Sum",
					"Breadth-first search from the start (Chapter 13)",
					"Backtracking over every path, keeping the shortest (Chapter 12)",
					"Interval DP over (left, right) ranges",
				},
				Answer:      1,
				Explanation: "With moves in all four directions, cells depend on each other in cycles (A is next to B and B is next to A), so there is no evaluation order for a table. BFS on an unweighted graph finds the shortest path in O(cells). Grid DP works when moves go only right/down (Unique Paths, Minimum Path Sum) or only to strictly larger values (Longest Increasing Path): then the dependencies form no cycle. Backtracking would work but is exponential.",
			},
			{
				Prompt: "One item with weight 2 and value 3, capacity 4, a one-row table best = [0, 0, 0, 0, 0] updated with 'for w in range(2, 5): best[w] = max(best[w], best[w - 2] + 3)' (forwards). What is best[4] afterwards?",
				Choices: []string{
					"3: the item packed once, which is right for 0/1 knapsack",
					"6: the item packed twice, which is right only if items may be reused",
					"0: the loop never reaches capacity 4",
				},
				Answer:      1,
				Explanation: "Forwards: w = 2 sets best[2] = 3; w = 3 sets best[3] = 3; w = 4 reads best[2], which was ALREADY updated for this item, so best[4] = 3 + 3 = 6. That is the unbounded rule (same row). For 0/1 knapsack loop backwards (w = 4, 3, 2): best[4] reads best[2] while it is still 0, giving 3.",
			},
			{
				Prompt:      "Trace this code. What does it return?\n\ndef f(x, y):\n    m, n = len(x), len(y)\n    t = [[0] * (n + 1) for _ in range(m + 1)]\n    for i in range(1, m + 1):\n        for j in range(1, n + 1):\n            if x[i - 1] == y[j - 1]:\n                t[i][j] = t[i - 1][j - 1] + 1\n            else:\n                t[i][j] = max(t[i - 1][j], t[i][j - 1])\n    return t[m][n]\n\nf(\"aab\", \"aba\")",
				Choices:     []string{"1", "2", "3", "0"},
				Answer:      1,
				Explanation: "It is LCS. Rows of t (columns j = 0..3): i=0: 0 0 0 0; i=1 ('a'): 0 1 1 1; i=2 ('aa'): 0 1 1 2; i=3 ('aab'): 0 1 2 2. The answer t[3][3] = 2, e.g. \"ab\" or \"aa\". It can't be 3 because \"aab\" is not a subsequence of \"aba\".",
			},
			{
				Prompt:      "Chapter 15 review: name the six steps of the DP recipe in order. Which of them need the most new thinking when the state gets a second index?",
				Explanation: "Define the state, Write the recurrence, Set the base cases, Choose the evaluation order, Locate the answer, Optimise space. With two indices the state definition (what does dp[i][j] mean, and is index 0 'empty'?) and the evaluation order (do all dependency arrows point to filled cells: row by row, by length, by value?) need the most care. Base cases become a whole row and column instead of one or two cells, and optimising space means keeping one or two rows instead of a few variables.",
			},
			{
				Prompt:      "Chapter 15 review: what is the difference between memoisation (top-down) and tabulation (bottom-up)? Give a case from this chapter where the top-down version can fail in Python while the table works.",
				Explanation: "Memoisation keeps the recursive function and caches each state's answer, so the recursion discovers the order and only visits states it needs. Tabulation fills a table in an order you choose, with loops and no call stack. Top-down can fail on deep recursion: Longest Increasing Path on a 200 x 200 snake-shaped grid recurses 40,000 calls deep, and Coin Change II with amount 5000 and a coin of 1 recurses about 5000 deep. The bottom-up versions (cells sorted by value, or the table) have no depth limit.",
			},
			{
				Prompt:      "In Burst Balloons, why does the recurrence choose the balloon burst LAST in a range rather than FIRST?",
				Explanation: "The last balloon k stays standing while every other balloon in the range is burst, so it acts as a wall: the balloons to its left only ever have the left wall and k as outer neighbours, and those to its right only k and the right wall. The two sides are independent subproblems, and k's own coins are vals[left] * vals[k] * vals[right]. If you choose the first balloon, its neighbours become adjacent after it bursts, the two sides interact, and the subproblems are no longer ranges.",
			},
			{
				Prompt:      "Coin Change II puts the coins in the outer loop and the amounts in the inner loop. What does the code count if you swap the loops, and what does it return for amount 4 with coins [1, 2]?",
				Explanation: "With amounts outside, ways[a] adds ways[a - c] for every coin c, so it asks 'what was the last coin?' and every coin can be last: it counts ordered sequences, where 1 + 2 + 1 and 2 + 1 + 1 differ. For amount 4 and coins [1, 2] it returns 5 (1111, 112, 121, 211, 22) instead of the 3 combinations (1111, 112, 22). Coins outside forces each combination to be built in one fixed order: all 1s, then all 2s.",
			},
			{
				Prompt: "In the Edit Distance table dist[i][j] (word1[:i] into word2[:j]), a cell takes 1 + dist[i][j - 1]. Which edit does that move stand for?",
				Choices: []string{
					"Delete word1[i - 1]",
					"Insert word2[j - 1]",
					"Replace word1[i - 1] with word2[j - 1]",
					"Keep a matching letter for free",
				},
				Answer:      1,
				Explanation: "Moving from (i, j - 1) to (i, j) produces one more letter of word2 without using a letter of word1: an insert of word2[j - 1]. The cell above, (i - 1, j), uses up a letter of word1 without producing one of word2 (delete), and the diagonal uses one of each (replace, or free on a match).",
			},
		},
	}
}
