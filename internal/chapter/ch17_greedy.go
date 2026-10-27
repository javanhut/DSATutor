package chapter

import "time"

// chapter17 teaches greedy algorithms: choosing a rule, proving it with an
// exchange or stays-ahead argument, and recognising when greedy fails and
// dynamic programming (chapter 15) is needed instead.
//
// Subgoal labels used here and in the greedy problems: "Choose the greedy
// rule", "Order the input", "Track the greedy state", "Commit or skip",
// "Check for failure", "Return the answer", "Justify the choice".
// Storyboards (IDs registered in app.js): greedy.
func chapter17() Chapter {
	return Chapter{
		Number:        17,
		Slug:          "greedy",
		Title:         "Greedy Algorithms",
		Category:      "greedy",
		Prerequisites: []int{15},
		Summary:       "A greedy algorithm makes the locally best choice at each step and never takes it back. It is fast and short, but only correct when you can prove the choice is safe. Learn to choose a greedy rule, justify it with an exchange or stays-ahead argument, and spot the counterexample that sends you back to dynamic programming.",
		Motivation: `A conference has 40 talks, each with a start and end time, all in one room. You want to schedule as many talks as possible with no two overlapping (one may start the moment another ends).

The safe way is to try every subset: 2^40 is about 1.1 trillion subsets, and each must be checked for clashes. Even at a billion subsets per second that is 18 minutes, and with 60 talks it's 36 years.

Dynamic programming (chapter 15) can do it in O(n^2) by trying every possible previous talk. But there is something faster: sort the talks by end time, and walk through them keeping every talk that starts after the last one you kept. That is one sort and one pass, O(n log n), a few hundred steps for 40 talks.

The catch: plenty of rules that sound just as sensible are wrong. 'Shortest talk first' fails on [1,5], [4,7], [6,10]: it picks [4,7], which blocks both others, so you get 1 talk instead of 2. 'Earliest start first' fails on [0,10], [1,2], [3,4]. A greedy algorithm is only as good as the argument that its choice is safe, and this chapter is about making that argument.`,
		Objectives: []string{
			"Explain what makes an algorithm greedy: one locally best choice per step, never undone",
			"Justify a greedy rule with an exchange argument (swap the optimal solution's first choice for the greedy one without making it worse)",
			"Justify a greedy rule with a 'greedy stays ahead' invariant",
			"Find a counterexample that breaks a tempting greedy rule, such as largest-coin-first for coins [1,3,4] and amount 6",
			"Decide between greedy and dynamic programming for a new problem, and implement activity selection, Jump Game and Kadane's algorithm",
		},
		WhenToUse: []string{
			"The problem asks for a maximum or minimum count (most meetings, fewest jumps, fewest removals) and you can describe a 'best next move' in one sentence",
			"Sorting the input by one key (end time, smallest value, deadline) seems to make the decisions obvious",
			"A running summary of the past (farthest reachable index, current sum, fuel in the tank) is enough to decide the next step; you never need to know which earlier choices produced it",
			"A DP solution exists but every state seems to pick the same option; that is a hint the choice can be made greedily",
			"Warning sign against greedy: you can build a small input where the locally best move blocks a better overall answer (coin change with arbitrary coins, 0/1 knapsack); use DP instead",
		},
		Concepts: []Concept{
			{
				Name:        "The greedy choice and the exchange argument",
				Description: "A greedy algorithm repeatedly picks the option that looks best right now according to a fixed rule, commits to it, and moves on. The exchange argument is the standard proof that the rule is safe: take any optimal solution and show you can swap its first choice for the greedy one without making it worse.",
				WhyItWorks: `Proof template (exchange argument):
1. Let g be the greedy's first choice, and let O be any optimal solution.
2. If O already contains g, fine. Otherwise O made a different first choice o.
3. Swap o for g in O. Show the result is still valid (no new conflicts) and no worse (same size or cost).
4. So some optimal solution starts with g. After committing to g, what is left is a smaller instance of the same problem, so repeat (induction).
For activity selection, g is the talk that ends first. If O's first talk o ends later than g, replace it: g ends no later than o, so it can't clash with the rest of O, which all start after o ends. Same number of talks, still valid. That is the whole proof.
The argument fails exactly when the swap can make things worse, and that is how you find counterexamples.`,
				Intuition: "Packing a day with as many errands as possible: always do next the errand that finishes soonest. Mapping: errands are intervals, 'finishes soonest' is the greedy rule, and the exchange argument is noticing that if your friend's best plan starts with a later-finishing errand, swapping in yours frees up at least as much of the afternoon.",
				CoreIdeas: []string{
					"Greedy = choose by a fixed rule, commit, never undo. No backtracking, no table of alternatives.",
					"Every greedy needs a proof; the rule is not obvious just because it sounds sensible.",
					"Exchange argument: an optimal solution that disagrees with the greedy can be changed to agree, without getting worse.",
					"A forced choice is the easiest case: in Hand of Straights the smallest card must start its group, so every solution agrees with the greedy.",
					"Sorting by the right key (end time, smallest value) is usually how the greedy rule is put into practice.",
				},
				CommonMistakes: []string{
					"Trusting a rule because it worked on the examples. Before coding, try to break it on 2 or 3 tiny inputs, e.g. shortest-first on [1,5],[4,7],[6,10].",
					"Sorting by the wrong key. Activity selection must sort by end time; sorting by start or by length both have counterexamples.",
					"Proving the greedy choice is 'good' instead of 'safe'. The proof must show an optimal solution exists that contains it, not just that it looks efficient.",
				},
				Examples: []CodeExample{
					{
						ID:       "activity_selection",
						Title:    "Activity selection: earliest finish first",
						Language: "python",
						Snippet: `def max_activities(activities):
    chosen = []
    current_end = float('-inf')
    for start, end in sorted(activities, key=lambda a: a[1]):
        if start >= current_end:   # compatible with everything chosen
            chosen.append([start, end])
            current_end = end
    return chosen`,
						Notes: "Earliest finish is a safe choice by the exchange argument. Touching intervals (start == current_end) are allowed.",
					},
				},
			},
			{
				Name:        "Greedy stays ahead",
				Description: "The second proof style: define a measure of progress and show that after every step the greedy is at least as far along as any other solution after the same number of steps. If it is never behind, it can't finish behind.",
				WhyItWorks: `Proof template (greedy stays ahead):
1. Pick a measure that you can compare step by step (index reached after k jumps, end time of the k-th chosen talk, best sum of a run ending at i).
2. Show by induction that after step k, the greedy's measure is at least as good as any valid solution's after step k.
3. Conclude the greedy reaches the goal no later than the optimum.
Jump Game II: let E_k be the furthest index reachable with k jumps, which is what the greedy tracks. Any path's (k+1)-th jump starts at some index <= E_k, so it lands at most at max(i + nums[i]) over i <= E_k = E_{k+1}. The greedy is never behind, so the first k with E_k >= n-1 is the minimum.
Many scan greedies are this argument with a one-number state: Jump Game's farthest, Kadane's current sum, Gas Station's tank. The invariant says the number summarises everything the past can offer the future.`,
				Intuition: "A relay race where you may pass the baton to anyone within reach. If after every leg your team's baton is at least as far down the track as any other team's, you can't cross the line later than them. Mapping: legs are greedy steps, baton position is the measure, and 'at least as far' is the invariant you prove by induction.",
				CoreIdeas: []string{
					"Choose a measure and prove, step by step, that the greedy is never behind any other solution.",
					"Scan greedies keep a tiny state (farthest, current, tank, low/high) and the invariant says what it means.",
					"Dropping a part that can't help is a greedy step too: Kadane drops a negative prefix, Gas Station drops a block of failed starts.",
					"When you can't commit to one choice, keep the best and worst cases (Valid Parenthesis String's [low, high] range).",
				},
				CommonMistakes: []string{
					"Stating the invariant vaguely ('farthest is the best so far'). Say exactly what it means: 'every index in [0, farthest] is reachable and none beyond it is'.",
					"Forgetting the boundary conditions the invariant needs: Kadane's best must start at nums[0], not 0, or the all-negative case breaks.",
					"Following one path instead of a frontier: always taking the full jump in Jump Game gets stuck on [2,3,0,0,0], while tracking the furthest reach over all reachable indices does not.",
				},
				Examples: []CodeExample{
					{
						ID:       "jump_game_frontier",
						Title:    "Jump Game: the reachable frontier",
						Language: "python",
						Snippet: `def canJump(nums):
    farthest = 0
    for i, jump in enumerate(nums):
        if i > farthest:          # i is beyond every reachable index
            return False
        farthest = max(farthest, i + jump)
    return True`,
						Notes: "Invariant: indices 0..farthest are all reachable. No route can get past farthest, so it stays ahead of every path.",
					},
				},
			},
			{
				Name:        "When greedy fails: counterexamples and dynamic programming",
				Description: "Greedy and DP both build an answer from choices. DP tries every option for each subproblem and keeps the best; greedy keeps only one. When the locally best option can block a better overall answer, greedy is wrong and you need DP from chapter 15.",
				WhyItWorks: `Coin change with coins [1, 3, 4] and amount 6. Greedy 'largest coin first' takes 4, then 1, then 1: 3 coins. The optimum is 3 + 3: 2 coins. The exchange argument breaks: swapping the optimum's first coin (3) for the greedy's (4) leaves 2, which costs two more coins, so the swap made things worse. DP fixes it by trying every coin: dp[a] = 1 + min(dp[a - c]) over coins c, giving dp[6] = 2.
0/1 knapsack, capacity 50, items (weight 10, value 60), (20, 100), (30, 120). Greedy by value per weight takes 10 and 20 for 160, then 30 doesn't fit. The optimum takes 20 and 30 for 220. Greedy works for the fractional knapsack (you may take part of an item: 60 + 100 + 80 = 240) because the exchange argument holds there; whole items break it.
With US coins [1, 5, 10, 25] largest-first happens to be optimal (63 cents: 25+25+10+1+1+1, 6 coins, same as DP), which is why the counterexample has to be searched for, not assumed.
Some problems, such as set cover, have no fast exact algorithm at all; there a greedy (pick the set covering the most uncovered items) is used knowingly as an approximation, within a factor of about ln n of optimal.`,
				Intuition: "Hiking toward the highest peak by always walking uphill. On a single hill it works; with two hills you can climb the small one and get stuck at its top. Mapping: 'always uphill' is the greedy rule, the small hill is a locally best choice that blocks the better answer, and DP is the map that shows every route.",
				CoreIdeas: []string{
					"To reject a greedy rule, one counterexample is enough; search small inputs (2 to 4 items) for it.",
					"Coins [1,3,4], amount 6: greedy 3 coins, optimal 2. 0/1 knapsack: greedy by ratio 160, optimal 220.",
					"If the exchange argument fails (the swap can make things worse), expect a counterexample and reach for DP.",
					"Greedy is the special case of DP where one option per state is always best, so you never need the table.",
					"A brute-force or DP solution on small random inputs is the fastest way to test a greedy idea.",
				},
				CommonMistakes: []string{
					"Assuming largest-coin-first works for any coin set because it works for real currency. Always test with odd denominations such as [1,3,4].",
					"Using the fractional-knapsack greedy on 0/1 knapsack. When items can't be split, greedy by value per weight is not optimal.",
					"Giving up on greedy too early: if DP ends up always choosing the same option, look for the proof that makes that option safe and drop the table.",
				},
				Examples: []CodeExample{
					{
						ID:       "coin_change_greedy_vs_dp",
						Title:    "Coin change: greedy fails, DP succeeds",
						Language: "python",
						Snippet: `def coins_greedy(coins, amount):          # NOT always optimal
    count = 0
    for coin in sorted(coins, reverse=True):
        count += amount // coin
        amount %= coin
    return count if amount == 0 else -1

def coins_dp(coins, amount):              # chapter 15
    dp = [0] + [float('inf')] * amount
    for a in range(1, amount + 1):
        for coin in coins:
            if coin <= a:
                dp[a] = min(dp[a], dp[a - coin] + 1)
    return dp[amount]

# coins_greedy([1, 3, 4], 6) == 3, coins_dp([1, 3, 4], 6) == 2`,
						Notes: "Greedy commits to 4 and can't recover; DP tries every first coin.",
					},
					{
						ID:       "set_cover",
						Title:    "Set cover: greedy as an approximation",
						Language: "python",
						Snippet: `def greedy_set_cover(universe, sets):
    uncovered = set(universe)
    chosen = []
    while uncovered:
        best = max(sets, key=lambda s: len(uncovered & s))
        if not uncovered & best:
            return None              # some item can't be covered
        chosen.append(best)
        uncovered -= best
    return chosen`,
						Notes: "Not always optimal, but provably within about ln n of the best. Greedy without an exchange proof gives an approximation, not an exact answer.",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Activity selection: schedule the most talks",
				Problem: "Given talks as [start, end] (a talk may start exactly when another ends), choose as many talks as possible with no two overlapping. Input: A[1,4], B[3,5], C[0,6], D[5,7], E[3,9], F[5,9], G[6,10], H[8,11], I[8,12], J[2,14], K[12,16].",
				Steps: []WorkedStep{
					{
						Subgoal:     "Choose the greedy rule",
						Explanation: "List candidate rules and try to break each on a tiny input. Earliest start fails on [0,10],[1,2],[3,4]. Shortest first fails on [1,5],[4,7],[6,10]. Earliest end survives: the talk that ends first leaves the most time for everything after it.",
						State:       "earliest start: picks [0,10] -> 1 talk (best 2); shortest: picks [4,7] -> 1 talk (best 2); earliest end: no counterexample",
					},
					{
						Subgoal:     "Order the input",
						Explanation: "Sort by end time, so that every time we look at a new talk it is the earliest-ending talk not yet considered.",
						Code:        "for start, end in sorted(activities, key=lambda a: a[1]):",
						State:       "sorted by end: A(1,4) B(3,5) C(0,6) D(5,7) E(3,9) F(5,9) G(6,10) H(8,11) I(8,12) J(2,14) K(12,16)",
					},
					{
						Subgoal:     "Track the greedy state",
						Explanation: "We only need one number: when the last chosen talk ends. Every chosen talk is compatible with all earlier ones, so a new talk only has to start after this time. Start at minus infinity so the first talk is always accepted.",
						Code:        "chosen = []\ncurrent_end = float('-inf')",
						State:       "chosen=[] current_end=-inf",
					},
					{
						Subgoal:     "Commit or skip",
						Explanation: "If the talk starts at or after current_end it fits: commit to it and move current_end. Otherwise it overlaps the last chosen talk: skip it forever. Note >=, because a talk may start the moment another ends.",
						Code:        "if start >= current_end:\n    chosen.append([start, end])\n    current_end = end",
						State:       "A: 1>=-inf take, end=4 | B: 3<4 skip | C: 0<4 skip | D: 5>=4 take, end=7 | E,F,G skip | H: 8>=7 take, end=11 | I,J skip | K: 12>=11 take, end=16",
					},
					{
						Subgoal:     "Return the answer",
						Explanation: "The chosen list is the schedule. Here 4 talks: A, D, H, K. One sort plus one pass is O(n log n).",
						Code:        "return chosen",
						State:       "chosen = [A(1,4), D(5,7), H(8,11), K(12,16)], 4 talks",
					},
					{
						Subgoal:     "Justify the choice",
						Explanation: "Exchange argument: let O be any best schedule and o its earliest-ending talk. A ends no later than o, and every other talk in O starts after o ends, so replacing o with A keeps O valid and the same size. So some best schedule starts with A; the rest is the same problem on talks starting at or after 4. By induction greedy is optimal.",
						State:       "e.g. O = [B(3,5), D, H, K] -> swap B for A -> [A, D, H, K], still 4 talks, still valid",
					},
				},
				Code: `def max_activities(activities):
    chosen = []
    current_end = float('-inf')  # end of the last chosen talk
    for start, end in sorted(activities, key=lambda a: a[1]):
        if start >= current_end:  # fits after everything chosen
            chosen.append([start, end])
            current_end = end
    return chosen`,
			},
			{
				Title:   "Testing a greedy rule: coin change with coins [1, 3, 4]",
				Problem: "Find the fewest coins that sum to an amount, with coin values [1, 3, 4] and unlimited copies of each. Does 'always take the largest coin that fits' work?",
				Steps: []WorkedStep{
					{
						Subgoal:     "Choose the greedy rule",
						Explanation: "The tempting rule is largest coin first: it reduces the amount fastest, and it is optimal for everyday coins like [1, 5, 10, 25].",
						Code:        "for coin in sorted(coins, reverse=True):\n    count += amount // coin\n    amount %= coin",
						State:       "amount 6: take 4 (left 2), take 1 (left 1), take 1 (left 0) -> 3 coins",
					},
					{
						Subgoal:     "Justify the choice",
						Explanation: "Try the exchange argument. An optimal answer for 6 is [3, 3]. Swap its first coin for the greedy's 4: the remaining amount is 2, which needs [1, 1]. The swapped solution has 3 coins, worse than 2. The swap step fails, so the proof fails.",
						State:       "optimal [3,3] (2 coins) -> swap first 3 for 4 -> [4,1,1] (3 coins): worse",
					},
					{
						Subgoal:     "Check for failure",
						Explanation: "A failed proof isn't yet a wrong algorithm, so confirm with a concrete counterexample. Compare greedy with the true best (from brute force) on amounts 1..8. Amount 6 is the only disagreement in that range, and one disagreement is enough to reject the rule.",
						State:       "amount: 1 2 3 4 5 6 7 8\ngreedy: 1 2 1 1 2 3 2 2\nbest:   1 2 1 1 2 2 2 2",
					},
					{
						Subgoal:     "Return the answer",
						Explanation: "Greedy is wrong here, so use the DP from chapter 15: dp[a] is the fewest coins for amount a, trying every coin as the last one. It is O(amount * coins) instead of greedy's O(coins), the price of not having a safe choice.",
						Code:        "dp = [0] + [float('inf')] * amount\nfor a in range(1, amount + 1):\n    for coin in coins:\n        if coin <= a:\n            dp[a] = min(dp[a], dp[a - coin] + 1)",
						State:       "dp = [0, 1, 2, 1, 1, 2, 2, 2] for amounts 0..7, so dp[6] = 2",
					},
				},
				Code: `def fewest_coins(coins, amount):
    # Greedy (largest coin first) gives 3 for coins [1,3,4], amount 6.
    # DP tries every last coin and gives the true answer, 2.
    dp = [0] + [float('inf')] * amount
    for a in range(1, amount + 1):
        for coin in coins:
            if coin <= a:
                dp[a] = min(dp[a], dp[a - coin] + 1)
    return dp[amount] if dp[amount] != float('inf') else -1`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Assign cookies",
				Prompt: "Each child has a greed value: child i is happy with any cookie of size >= greed[i]. Each child gets at most one cookie. Return the largest number of happy children. Fill in the blanks using the same subgoals: order the input, track the greedy state, commit or skip.",
				Code: `def findContentChildren(greed, cookies):
    greed = sorted(greed)
    cookies = ____                   # Order the input
    child = 0                        # Track the greedy state: children made happy
    for cookie in cookies:           # smallest cookie first
        if child < len(greed) and ____:
            ____                     # Commit: this cookie goes to this child
    return child`,
				Blanks:      []string{"sorted(cookies)", "cookie >= greed[child]", "child += 1"},
				Explanation: "Sort both lists, then give each cookie, smallest first, to the least greedy child still waiting if it satisfies them; otherwise the cookie is too small for every waiting child and is skipped. Exchange argument: in any best assignment, the least greedy child can be given the smallest cookie that satisfies them without making anything worse, because any other child that cookie could please is at least as greedy and can use the larger cookie instead. On greed [1,2,3] with cookies [1,1] the answer is 1; on greed [1,2] with cookies [1,2,3] it is 2.",
			},
		},
		Animations: []Storyboard{
			{
				ID:    "greedy",
				Title: "Activity Selection (Greedy)",
				Goal:  "See why picking the talk that finishes first, again and again, gives the largest schedule, and how the exchange argument justifies each pick.",
				MemoryTips: []string{
					"Choose the greedy rule: earliest finish, and test it on tiny inputs first",
					"Order the input: sort by end time",
					"Track the greedy state: current_end, the end of the last chosen talk",
					"Commit or skip: take a talk if start >= current_end, otherwise skip it forever",
					"Justify the choice: swap the best schedule's first talk for the earliest-ending one; it stays valid and the same size",
				},
				Steps: []StoryboardStep{
					{Cue: "sort", Narration: "Sort the talks by finish time. The talk that ends first leaves the most room for the rest, which is the greedy rule we will justify.", VisualHint: "Timeline bars reorder by end time", Duration: Duration(3 * time.Second), CodeRef: "activity_selection: sort by end"},
					{Cue: "consider", Narration: "Look at the next talk in finish order. current_end is when the last chosen talk ends; at the start it is minus infinity, so the first talk always fits.", VisualHint: "Highlight current bar and the current_end marker", Duration: Duration(2 * time.Second), CodeRef: "activity_selection: for activity"},
					{Cue: "check", Narration: "Does it start at or after current_end? If so it can't overlap anything chosen, because every chosen talk ended by current_end.", VisualHint: "Compare the bar's left edge with the current_end line", Duration: Duration(2 * time.Second), CodeRef: "activity_selection: if compatible"},
					{Cue: "select", Narration: "Compatible: commit to it and move current_end to its end. We never revisit this decision.", VisualHint: "Bar turns green and joins the schedule; current_end line moves", Duration: Duration(3 * time.Second), CodeRef: "activity_selection: select"},
					{Cue: "skip", Narration: "Overlapping: skip it for good. Any schedule that used it instead of the chosen talk would end later and leave less room.", VisualHint: "Bar greys out", Duration: Duration(2 * time.Second), CodeRef: "activity_selection: if compatible"},
					{Cue: "result", Narration: "After one pass the chosen talks are a largest schedule: A, D, H and K. The exchange argument guarantees no schedule has more.", VisualHint: "Chosen bars highlighted on the timeline", Duration: Duration(3 * time.Second), CodeRef: "activity_selection: return"},
				},
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "To schedule the most non-overlapping talks, which greedy rule is always optimal?",
				Choices: []string{
					"Pick the shortest talk first",
					"Pick the talk that starts earliest",
					"Pick the talk that ends earliest",
					"Pick the talk that overlaps the fewest others",
				},
				Answer:      2,
				Explanation: "Earliest end is safe by the exchange argument: it can replace the first talk of any best schedule. Shortest first fails on [1,5],[4,7],[6,10] (it picks [4,7] and gets 1 instead of 2), and earliest start fails on [0,10],[1,2],[3,4]. Fewest overlaps sounds clever but also has counterexamples.",
			},
			{
				Prompt: "Which pattern fits best? 'Given coin values (any set, e.g. [1, 3, 4]) and an amount, return the fewest coins that make the amount.'",
				Choices: []string{
					"Greedy: take the largest coin that fits, repeat",
					"Dynamic programming over amounts (chapter 15)",
					"Sliding window over the sorted coins (chapter 4)",
					"Binary search on the number of coins (chapter 6)",
				},
				Answer:      1,
				Explanation: "With arbitrary coins the largest-first greedy fails: [1,3,4] for 6 gives 4+1+1 = 3 coins while 3+3 = 2. The exchange argument breaks, so you need DP: dp[a] = 1 + min(dp[a - c]). There is no contiguous window to slide, and the number of coins isn't something you can check cheaply for binary search.",
			},
			{
				Prompt:      "Trace canJump on nums = [3, 2, 1, 0, 4]:\n\nfarthest = 0\nfor i, jump in enumerate(nums):\n    if i > farthest:\n        return False\n    farthest = max(farthest, i + jump)\nreturn True\n\nWhat is farthest after each index, and what is returned?",
				Explanation: "i=0: farthest = 3. i=1: max(3, 3) = 3. i=2: max(3, 3) = 3. i=3: max(3, 3) = 3. i=4: 4 > 3, so it returns False. Every index before the 0 at index 3 reaches exactly index 3, so the frontier never passes it.",
			},
			{
				Prompt:      "Prerequisite (chapter 15): greedy and dynamic programming both build an answer from choices. In your own words, what does DP do at each state that greedy does not, and when does that extra work become unnecessary?",
				Explanation: "DP tries every option for a state and keeps the best (e.g. dp[a] = 1 + min over every coin c of dp[a - c]), so it never commits too early, at the cost of filling a table. Greedy picks one option by a rule and never compares alternatives. The extra work is unnecessary when you can prove one option is always among the best, with an exchange or stays-ahead argument; then the table always takes that option and you can drop it.",
			},
			{
				Prompt:      "Write the exchange argument for activity selection (earliest end first) in three or four sentences.",
				Explanation: "Let g be the talk that ends earliest and O any largest schedule, with o its earliest-ending talk. g ends no later than o, and every other talk in O starts at or after o ends, so also at or after g ends. Replacing o with g keeps O valid and the same size, so some largest schedule contains g. After taking g, the remaining talks that start at or after g's end form a smaller instance of the same problem, so by induction greedy is optimal.",
			},
			{
				Prompt: "0/1 knapsack with capacity 50 and items (weight 10, value 60), (20, 100), (30, 120). What does 'take the best value-per-weight first' return, and what is optimal?",
				Choices: []string{
					"Greedy 160, optimal 220",
					"Greedy 220, optimal 220",
					"Greedy 240, optimal 240",
					"Greedy 180, optimal 220",
				},
				Answer:      0,
				Explanation: "Ratios are 6, 5 and 4, so greedy takes the 10 and the 20 (weight 30, value 160); the 30 no longer fits. Taking 20 and 30 (weight 50) gives 220. 240 is the fractional knapsack answer, where you may take 2/3 of the last item; there the greedy is correct.",
			},
			{
				Prompt:      "In Gas Station, why does running out of fuel on the way out of station i rule out every start from the current candidate up to i, not just the candidate itself?",
				Explanation: "From the candidate, you arrived at each station p in that range with a tank >= 0 (i was the first failure). Starting at p means arriving with 0 instead, which is no more fuel, over the same legs from p to i, so you'd fail at i or sooner. That is why the next candidate can jump straight to i + 1.",
			},
		},
	}
}
