package practice

// SlidingWindowProblems contains all sliding-window category problems.
//
// Ladder (Chapter 4): the buy/sell window (left edge = cheapest day so far),
// then variable windows that grow right and shrink left while invalid
// (no repeats, then a count-based validity rule), a fixed-size window with
// counts, a "shortest valid window" with formed/required bookkeeping, and
// finally a fixed window whose state is a monotonic deque.
var SlidingWindowProblems = []*Problem{
	{
		ID:              "best-time-to-buy-and-sell-stock",
		Number:          121,
		Order:           1,
		Title:           "Best Time to Buy and Sell Stock",
		Slug:            "best-time-to-buy-and-sell-stock",
		Difficulty:      "Easy",
		Category:        "sliding-window",
		Tags:            []string{"Array", "Sliding Window", "Greedy"},
		RelatedChapters: []int{4, 3},
		Description: `You are given a list prices where prices[i] is the price of one share of a stock on day i.

You may make exactly one trade: pick one day to buy a single share and a strictly later day to sell it. Return the largest profit (sell price minus buy price) you can make. If every possible trade loses money or breaks even, don't trade and return 0.

Think of it as a window [buy, sell] over the days: the left edge is the day you buy, the right edge the day you sell, and the left edge must stay to the left of the right edge.`,
		Constraints: []string{
			"1 <= prices.length <= 10^5",
			"0 <= prices[i] <= 10^4",
		},
		Examples: []Example{
			{Input: "prices = [7,1,5,3,6,4]", Output: "5", Explanation: "Buy on day 1 (price 1) and sell on day 4 (price 6): 6 - 1 = 5. Selling on day 0 at 7 is not allowed because it comes before the buy."},
			{Input: "prices = [7,6,4,3,1]", Output: "0", Explanation: "The price only falls, so every trade loses money. Don't trade."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"prices": []int{7, 1, 5, 3, 6, 4}}, Expected: 5, Note: "example: best buy is in the middle, best sell later"},
			{Input: map[string]interface{}{"prices": []int{7, 6, 4, 3, 1}}, Expected: 0, Note: "falling prices: profit can't go negative"},
			{Input: map[string]interface{}{"prices": []int{1}}, Expected: 0, Hidden: true, Note: "single day: no trade possible"},
			{Input: map[string]interface{}{"prices": []int{2, 4, 1}}, Expected: 2, Hidden: true, Note: "the lowest price comes after the highest, so max(prices) - min(prices) = 3 is wrong"},
			{Input: map[string]interface{}{"prices": []int{3, 3, 3}}, Expected: 0, Hidden: true, Note: "all prices equal: breaking even is 0"},
			{Input: map[string]interface{}{"prices": []int{5, 2, 6, 1, 4}}, Expected: 4, Hidden: true, Note: "a newer, lower buy day (1) appears but the older one (2) gives the best trade"},
			{Input: map[string]interface{}{"prices": []int{2, 1, 2, 1, 0, 1, 2}}, Expected: 2, Hidden: true, Note: "the buy day moves several times; the answer uses the last low"},
			{Input: map[string]interface{}{"prices": []int{1, 2}}, Expected: 1, Note: "two days, rising"},
			{Input: map[string]interface{}{"prices": []int{3, 1, 5, 4, 6}}, Expected: 5, Hidden: true, Note: "a small dip (5 to 4) must not move the buy day off the real low"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def maxProfit(prices):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Suppose you have decided to sell today. Which single earlier day would you want to have bought on, and what do you need to remember from the days you have already passed to know it?"},
			{Level: 2, Type: "insight", Content: "For a fixed sell day, the best buy day is the cheapest day before it. So keep a window [buy, sell]: the right edge (sell) moves forward every day, and the left edge (buy) only jumps forward when today is cheaper than the current buy day, because a cheaper day is a better buy for every future sell. This is the sliding-window pattern in its simplest form (Chapter 4).", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Set up the window state: buy = day 0, best = 0.\n2. Grow the window: move sell from day 1 to the last day.\n3. Shrink until valid: if prices[sell] is lower than prices[buy], the window would lose money, so move buy to sell.\n4. Record the answer: otherwise compare prices[sell] - prices[buy] with best.\n5. Return best."},
			{Level: 4, Type: "skeleton", Content: "def maxProfit(prices):\n    buy = 0\n    best = 0\n    for sell in range(1, len(prices)):\n        if ____:\n            buy = ____\n        else:\n            best = ____\n    return best"},
		},
		Solution: Solution{
			Code: `def maxProfit(prices):
    buy = 0    # left edge: the cheapest day seen so far
    best = 0
    for sell in range(1, len(prices)):  # right edge moves every day
        if prices[sell] < prices[buy]:
            buy = sell  # cheaper day found: restart the window here
        else:
            best = max(best, prices[sell] - prices[buy])
    return best`,
			Explanation: "Treat the trade as a window [buy, sell] over the days. The right edge walks forward one day at a time. The left edge stays on the cheapest day seen so far and jumps to the right edge whenever a new low appears. Each step records the profit of selling today. This is the plainest sliding window: two indices that only move forward, with the left one moved by a rule about what makes the window worth keeping. Later problems in this ladder keep the same shape and make the rule richer.",
			BruteForce:  "Try every pair of days i < j and keep the largest prices[j] - prices[i]. That is n(n-1)/2 pairs, about 5 billion for n = 100,000, and O(n^2) time. The waste: for each sell day j it re-scans every earlier day to find the cheapest one, although the cheapest day before j is just the cheapest day before j - 1 or day j - 1 itself.",
			KeyInsight:  "For any sell day, the only buy day worth considering is the cheapest day before it. That minimum can be kept up to date in O(1) as the sell day advances, so one pass is enough.",
			WhyItWorks:  "Invariant: at the start of each step, prices[buy] is the minimum of prices[0..sell-1]. It holds at the start (buy = 0, sell = 1) and the if-branch restores it whenever prices[sell] is a new minimum. Given the invariant, prices[sell] - prices[buy] is the best profit for selling on day sell, and best is the maximum of those values over all sell days, which is the answer. When the left edge jumps to a new low, no trade that buys before it can beat buying at the new low for any later sell day, so skipping those buy days loses nothing.",
			CommonMistakes: []string{
				"Returning max(prices) - min(prices). That ignores the order of days: for [2,4,1] it gives 3, but the 1 comes after the 4. Track the minimum only over days before the sell day.",
				"Starting best at a negative number or allowing a negative result. If prices only fall, the right answer is 0 (no trade). Start best at 0.",
				"Comparing today with yesterday instead of with the buy day (moving buy whenever the price dips). For [3,1,5,4,6] the dip to 4 would move buy off the true low of 1 and give 4 instead of 5. Move the left edge only when today is cheaper than prices[buy].",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the window state", Explanation: "The window starts with buy on day 0. best starts at 0 because not trading is always allowed.", CodeSnippet: "buy = 0    # left edge: the cheapest day seen so far\nbest = 0", LineStart: 2, LineEnd: 3},
				{Title: "Grow the window", Explanation: "The sell day (right edge) moves forward one day per step and never goes back. Starting at 1 guarantees sell is after buy.", CodeSnippet: "for sell in range(1, len(prices)):", LineStart: 4, LineEnd: 4},
				{Title: "Shrink until valid", Explanation: "If today is cheaper than the buy day, this window would lose money and today is a better buy for every later day. The left edge jumps to today. For [7,1,5,...] this happens at day 1: buy moves from 7 to 1.", CodeSnippet: "if prices[sell] < prices[buy]:\n    buy = sell", LineStart: 5, LineEnd: 6},
				{Title: "Record the answer", Explanation: "Otherwise the window is worth checking: selling today earns prices[sell] - prices[buy]. For [7,1,5,3,6,4] the recorded profits are 4, 2, 5, 3, so best ends at 5.", CodeSnippet: "best = max(best, prices[sell] - prices[buy])", LineStart: 8, LineEnd: 8},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: every pair of days",
					Code: `def maxProfit(prices):
    best = 0
    for i in range(len(prices)):
        for j in range(i + 1, len(prices)):
            best = max(best, prices[j] - prices[i])
    return best`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Checks all n(n-1)/2 buy/sell pairs. Correct, and a good way to check the fast version on small inputs, but far too slow for 10^5 days.",
				},
				{
					Name: "Running minimum",
					Code: `def maxProfit(prices):
    lowest = prices[0]
    best = 0
    for price in prices[1:]:
        best = max(best, price - lowest)
        lowest = min(lowest, price)
    return best`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "The same idea written with values instead of indices: lowest is prices[buy]. Compute today's profit before updating lowest, so you never buy and sell on the same day.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is it safe to move the left edge (buy) all the way to today when today is a new low, skipping every day in between?", Answer: "Any later sell day earns more from buying at the new low than from buying at any earlier day, because the new low is cheaper than all of them. So no future trade needs those earlier days as its buy day. Trades that sell before today were already recorded in best."},
			{Prompt: "What if you could make up to two trades instead of one? Does one running minimum still work?", Answer: "No. Two trades need more state: for example, the best single-trade profit ending on or before each day (computed left to right) and the best starting on or after each day (computed right to left), then the best split point. One running minimum only describes one trade. That is a dynamic programming problem (Chapter 15)."},
		},
	},
	{
		ID:              "longest-substring-without-repeating",
		Number:          3,
		Order:           2,
		Title:           "Longest Substring Without Repeating Characters",
		Slug:            "longest-substring-without-repeating-characters",
		Difficulty:      "Medium",
		Category:        "sliding-window",
		Tags:            []string{"Hash Table", "String", "Sliding Window"},
		RelatedChapters: []int{4, 2},
		Description: `Given a string s, return the length of the longest substring of s in which no character appears more than once.

A substring is a contiguous run of characters (so "pwke" is not a substring of "pwwkew", because its letters are not next to each other). If s is empty, return 0.`,
		Constraints: []string{
			"0 <= s.length <= 5 * 10^4",
			"s consists of English letters, digits, symbols and spaces",
		},
		Examples: []Example{
			{Input: `s = "abcabcbb"`, Output: "3", Explanation: `"abc" has no repeats. Every substring of length 4 contains some letter twice.`},
			{Input: `s = "bbbbb"`, Output: "1", Explanation: `Any two characters are "bb", a repeat, so the best is a single "b".`},
			{Input: `s = "pwwkew"`, Output: "3", Explanation: `"wke" works. "pwke" has no repeats but isn't contiguous.`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "abcabcbb"}, Expected: 3, Note: "example: repeats force the window to shrink"},
			{Input: map[string]interface{}{"s": "bbbbb"}, Expected: 1, Note: "all characters equal"},
			{Input: map[string]interface{}{"s": "pwwkew"}, Expected: 3, Note: "the answer is in the middle, not at the start"},
			{Input: map[string]interface{}{"s": ""}, Expected: 0, Hidden: true, Note: "empty string"},
			{Input: map[string]interface{}{"s": " "}, Expected: 1, Hidden: true, Note: "a single space is a character too"},
			{Input: map[string]interface{}{"s": "abba"}, Expected: 2, Hidden: true, Note: "a jump-to-last-index version must never move left backwards (the second a's last index is 0)"},
			{Input: map[string]interface{}{"s": "dvdf"}, Expected: 3, Hidden: true, Note: "on a repeat, shrink just past the old copy; don't restart from the current character"},
			{Input: map[string]interface{}{"s": "tmmzuxt"}, Expected: 5, Hidden: true, Note: "the final t repeats one far to the left, outside the current window"},
			{Input: map[string]interface{}{"s": "a b!a"}, Expected: 4, Hidden: true, Note: "spaces and symbols"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(min(n, m)) where m is the alphabet size",
		StarterCode:     "def lengthOfLongestSubstring(s):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Say the substring s[left..right] has no repeats and you add s[right+1], which is already inside. Does any substring that still starts at left have a chance now? Where is the first start that could work?"},
			{Level: 2, Type: "insight", Content: "If a window contains a repeated character, every bigger window around it does too. So when s[right] is a repeat, move left forward (removing characters) until the old copy is gone; left never has to move back. Keep the characters of the window in a set so you can test 'is this a repeat?' in O(1). This is the variable-size sliding window from Chapter 4.", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Set up the window state: an empty set of the characters in the window, left = 0, best = 0.\n2. Grow the window: for each right, look at s[right].\n3. Shrink until valid: while s[right] is already in the set, remove s[left] from the set and move left forward.\n4. Add s[right] to the set.\n5. Record the answer: best = max(best, window length right - left + 1)."},
			{Level: 4, Type: "skeleton", Content: "def lengthOfLongestSubstring(s):\n    in_window = set()\n    left = 0\n    best = 0\n    for right in range(len(s)):\n        while ____:\n            in_window.remove(____)\n            left += 1\n        in_window.add(s[right])\n        best = ____\n    return best"},
		},
		Solution: Solution{
			Code: `def lengthOfLongestSubstring(s):
    in_window = set()  # characters in s[left..right-1]
    left = 0
    best = 0
    for right in range(len(s)):
        while s[right] in in_window:
            in_window.remove(s[left])
            left += 1
        in_window.add(s[right])
        best = max(best, right - left + 1)
    return best`,
			Explanation: "This is the first variable-size window in the ladder. Like Best Time to Buy and Sell Stock, the right edge moves one step at a time and the left edge only moves forward. Here the left edge creeps forward one character at a time, removing characters until the window is valid again, which is the general 'shrink until valid' loop. A set holds the window's characters so the validity check (is s[right] already inside?) is O(1).",
			BruteForce:  "Check every start i, extend j while s[j] isn't in a set built for that start, and keep the longest run. That is O(n^2) in the worst case (about 1.25 billion steps for n = 50,000). The waste: after a repeat stops start i, start i + 1 rebuilds almost the same set from scratch, even though the window s[i+1..j-1] is already known to have no repeats.",
			KeyInsight:  "Validity is monotonic: if s[left..right] contains a repeat, so does every longer window containing it. So when a repeat appears you never need an earlier left again; just slide left forward past the old copy and keep everything you already know about the rest of the window.",
			WhyItWorks:  "Invariant: after the while loop, s[left..right] has no repeated character and left is the smallest start for which that is true. Before adding s[right], s[left..right-1] had no repeats; removing characters from the left until s[right] is gone makes s[left..right] repeat-free, and any smaller left would still contain the old copy of s[right]. So right - left + 1 is the longest valid substring ending at right, and best is the maximum over all right. Each index enters the set once and leaves at most once, so the total work is O(n) even though there is a loop inside a loop.",
			CommonMistakes: []string{
				"Resetting the window to start at right when a repeat appears (left = right, clear the set). For \"dvdf\" that gives 2 instead of 3 (\"vdf\"). Remove characters only until the old copy of s[right] is gone.",
				"In the faster 'last seen index' version, writing left = last[c] + 1 without max(left, ...). For \"abba\" the final a's last index is 0, which would move left backwards to 1 and count \"bba\". Use left = max(left, last[c] + 1).",
				"Recording the length before shrinking. The window is only valid after the while loop, so record the answer after it.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(min(n, m)) where m is the alphabet size",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the window state", Explanation: "The set holds exactly the characters of the current window. left marks the start of the window.", CodeSnippet: "in_window = set()  # characters in s[left..right-1]\nleft = 0\nbest = 0", LineStart: 2, LineEnd: 4},
				{Title: "Grow the window", Explanation: "Each step considers adding one new character on the right.", CodeSnippet: "for right in range(len(s)):", LineStart: 5, LineEnd: 5},
				{Title: "Shrink until valid", Explanation: "While the new character is already inside, drop characters from the left. In \"abcabcbb\" at right = 3 (a), the window \"abc\" drops a and becomes \"bca\" after adding.", CodeSnippet: "while s[right] in in_window:\n    in_window.remove(s[left])\n    left += 1", LineStart: 6, LineEnd: 8},
				{Title: "Record the answer", Explanation: "Now s[left..right] has no repeats; add s[right] and compare its length with the best so far.", CodeSnippet: "in_window.add(s[right])\nbest = max(best, right - left + 1)", LineStart: 9, LineEnd: 10},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: try every start",
					Code: `def lengthOfLongestSubstring(s):
    best = 0
    for i in range(len(s)):
        seen = set()
        for j in range(i, len(s)):
            if s[j] in seen:
                break
            seen.add(s[j])
            best = max(best, j - i + 1)
    return best`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(min(n, m))",
					Explanation:     "For each start, extend until the first repeat. Simple and correct, but it rebuilds the set for every start.",
				},
				{
					Name: "Jump left using last-seen indices",
					Code: `def lengthOfLongestSubstring(s):
    last = {}
    left = 0
    best = 0
    for right, ch in enumerate(s):
        if ch in last:
            left = max(left, last[ch] + 1)
        last[ch] = right
        best = max(best, right - left + 1)
    return best`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(min(n, m))",
					Explanation:     "Instead of removing characters one by one, remember where each character was last seen and jump left past it. The max() matters: an old index that is already left of the window must not pull left backwards (\"abba\").",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "The code has a while loop inside a for loop. Why is it still O(n) and not O(n^2)?", Answer: "Count the work per character instead of per loop. Each index is added to the set once (when right reaches it) and removed at most once (when left passes it). left only moves forward and never past right, so across the whole run the inner loop executes at most n times in total. That is about 2n set operations, which is O(n)."},
			{Prompt: "What changes if the question asks for the longest substring with at most 2 repeats of any character, or at most k distinct characters?", Answer: "Only the validity rule and the window state. Replace the set with a count dictionary; grow by incrementing count[s[right]]; shrink while the rule is broken (some count above the limit, or more than k keys), decrementing count[s[left]] and deleting zero entries. The grow / shrink until valid / record skeleton stays the same."},
		},
	},
	{
		ID:              "longest-repeating-character-replacement",
		Number:          424,
		Order:           3,
		Title:           "Longest Repeating Character Replacement",
		Slug:            "longest-repeating-character-replacement",
		Difficulty:      "Medium",
		Category:        "sliding-window",
		Tags:            []string{"Hash Table", "String", "Sliding Window"},
		RelatedChapters: []int{4, 2},
		Description: `You are given a string s of uppercase English letters and an integer k. You may pick any position and change its letter to any other uppercase letter, at most k times in total.

Return the length of the longest substring (contiguous run) that can be made of one repeated letter after at most k such changes.`,
		Constraints: []string{
			"1 <= s.length <= 10^5",
			"s consists of only uppercase English letters",
			"0 <= k <= s.length",
		},
		Examples: []Example{
			{Input: `s = "ABAB", k = 2`, Output: "4", Explanation: `Change both A's to B (or both B's to A): "BBBB" has length 4.`},
			{Input: `s = "AABABBA", k = 1`, Output: "4", Explanation: `Change the A at index 3 to B: "AABBBBA". The run "BBBB" has length 4. No window of length 5 has at most one letter that differs from its most common letter.`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "ABAB", "k": 2}, Expected: 4, Note: "example: k is enough to fix the whole string"},
			{Input: map[string]interface{}{"s": "AABABBA", "k": 1}, Expected: 4, Note: "example: the window must shrink"},
			{Input: map[string]interface{}{"s": "A", "k": 0}, Expected: 1, Hidden: true, Note: "single character, no changes allowed"},
			{Input: map[string]interface{}{"s": "AAAA", "k": 0}, Expected: 4, Hidden: true, Note: "all letters equal"},
			{Input: map[string]interface{}{"s": "ABCDE", "k": 1}, Expected: 2, Hidden: true, Note: "all letters distinct: one change makes a pair"},
			{Input: map[string]interface{}{"s": "ABAA", "k": 0}, Expected: 2, Hidden: true, Note: "k = 0 means the longest run that already exists"},
			{Input: map[string]interface{}{"s": "BAAAB", "k": 2}, Expected: 5, Hidden: true, Note: "the most common letter is in the middle"},
			{Input: map[string]interface{}{"s": "AAABBBCCCA", "k": 2}, Expected: 5, Hidden: true, Note: "three letters compete for the majority"},
			{Input: map[string]interface{}{"s": "AABCDAA", "k": 2}, Expected: 4, Hidden: true, Note: "the majority letter leaves the window, so the stored max_freq goes stale"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def characterReplacement(s, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Look at one window, for example \"AABAB\". Which letter would you keep, and how many changes would turning the whole window into that letter cost? Can you compute that from counts alone?"},
			{Level: 2, Type: "insight", Content: "A window can be made uniform by keeping its most frequent letter and changing everything else, which costs (window length - count of the most frequent letter). The window is valid when that cost is <= k. Keep letter counts for the window and use the variable-size sliding window from Chapter 4: grow right, shrink left while the cost is above k.", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Set up the window state: a count for each letter in the window, max_freq = 0, left = 0, best = 0.\n2. Grow the window: add s[right] to the counts and raise max_freq if that letter's count is now the largest.\n3. Shrink until valid: while (window length - max_freq) > k, decrement the count of s[left] and move left forward.\n4. Record the answer: best = max(best, window length)."},
			{Level: 4, Type: "skeleton", Content: "def characterReplacement(s, k):\n    count = {}\n    max_freq = 0\n    left = 0\n    best = 0\n    for right in range(len(s)):\n        count[s[right]] = count.get(s[right], 0) + 1\n        max_freq = ____\n        while ____ > k:\n            count[s[left]] -= 1\n            left += 1\n        best = ____\n    return best"},
		},
		Solution: Solution{
			Code: `def characterReplacement(s, k):
    count = {}      # letter -> how many times it appears in the window
    max_freq = 0    # the largest count seen in any window so far
    left = 0
    best = 0
    for right in range(len(s)):
        count[s[right]] = count.get(s[right], 0) + 1
        max_freq = max(max_freq, count[s[right]])
        while (right - left + 1) - max_freq > k:
            count[s[left]] -= 1
            left += 1
        best = max(best, right - left + 1)
    return best`,
			Explanation: "This builds on Longest Substring Without Repeating Characters: the same grow / shrink until valid / record loop, but the validity rule is now a number computed from counts instead of 'no repeats'. A window of length L whose most common letter appears f times needs L - f changes, so it is valid when L - f <= k. The twist is that max_freq is never decreased when the left edge moves, and the answer is still right.",
			BruteForce:  "For every start i, extend j and keep letter counts; the window s[i..j] is fixable when (j - i + 1) - max(counts) <= k. That is O(n^2) windows (5 billion for n = 10^5), each checked in O(1) with at most 26 counts. The waste: start i + 1 recounts letters that start i already counted.",
			KeyInsight:  "Fixability is monotonic in the window: shrinking a fixable window keeps it fixable, and growing an unfixable one keeps it unfixable. So the two-pointer window applies, with cost = length - max_freq as the validity test.",
			WhyItWorks:  "Why the stale max_freq is still correct: max_freq only goes up, even when the letter that produced it leaves the window. Before the first shrink, no letter has left, so max_freq is exact. The loop shrinks only when length - max_freq > k, and it then removes exactly one letter, after which length = max_freq + k. From then on the window never gets shorter: each step adds one letter and removes at most one. It gets longer only if max_freq rises, and max_freq rises only when count[s[right]], a true count inside the current window, beats every earlier count. At that moment the window really is fixable. So every time best increases it records a genuinely valid window. When max_freq is stale, the window may be invalid, but it just slides at a length that was already achieved, so it can't inflate best. Since we only want the maximum length, a smaller window would never help, and we never need to recompute the true maximum.",
			CommonMistakes: []string{
				"Using the count of s[right] as the majority instead of the largest count in the window. For \"AABABBA\" with k = 1, the window \"AABA\" ends in A but you need the overall maximum. Track max_freq across all letters.",
				"Recomputing max(count.values()) and shrinking the window back down whenever the majority letter leaves. It is correct but unnecessary work; understand why the stale max_freq can't produce a too-large answer (the window length never exceeds a previously valid length unless max_freq truly increases).",
				"Writing the condition as (window length - max_freq) >= k. Exactly k changes is allowed, so shrink only when the cost is strictly greater than k; \"ABAB\" with k = 2 must return 4.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the window state", Explanation: "count holds letter counts for s[left..right]; there are at most 26 keys, so the space is O(1). max_freq remembers the highest count seen.", CodeSnippet: "count = {}      # letter -> how many times it appears in the window\nmax_freq = 0    # the largest count seen in any window so far", LineStart: 2, LineEnd: 3},
				{Title: "Grow the window", Explanation: "Add s[right] and raise max_freq if this letter is now the most common one. That count is a real count inside the current window.", CodeSnippet: "count[s[right]] = count.get(s[right], 0) + 1\nmax_freq = max(max_freq, count[s[right]])", LineStart: 7, LineEnd: 8},
				{Title: "Shrink until valid", Explanation: "The cost to fix the window is its length minus max_freq. While that is above k, drop the leftmost letter. For \"AABABBA\", k = 1 at right = 4 the window \"AABAB\" costs 5 - 3 = 2, so the first A leaves.", CodeSnippet: "while (right - left + 1) - max_freq > k:\n    count[s[left]] -= 1\n    left += 1", LineStart: 9, LineEnd: 11},
				{Title: "Record the answer", Explanation: "Compare the window length with the best. Because max_freq never decreases, the shrink loop runs at most once per step and the window never gets shorter.", CodeSnippet: "best = max(best, right - left + 1)", LineStart: 12, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: every window with counts",
					Code: `def characterReplacement(s, k):
    best = 0
    for i in range(len(s)):
        count = {}
        for j in range(i, len(s)):
            count[s[j]] = count.get(s[j], 0) + 1
            if (j - i + 1) - max(count.values()) <= k:
                best = max(best, j - i + 1)
    return best`,
					TimeComplexity:  "O(n^2)",
					SpaceComplexity: "O(1)",
					Explanation:     "Checks every window with the exact majority count. Clear and correct; useful to test the fast version.",
				},
				{
					Name: "One window per target letter",
					Code: `def characterReplacement(s, k):
    best = 0
    for letter in set(s):
        left = 0
        others = 0  # letters in the window that are not 'letter'
        for right in range(len(s)):
            if s[right] != letter:
                others += 1
            while others > k:
                if s[left] != letter:
                    others -= 1
                left += 1
            best = max(best, right - left + 1)
    return best`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Fix the letter to keep, then find the longest window with at most k other letters. There are at most 26 letters, so this is still linear, and it avoids the stale max_freq subtlety entirely.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why may max_freq go stale (too large) without making the answer too big?", Answer: "Once the window has shrunk once, its length is always max_freq + k. It grows only in a step where max_freq grows, and max_freq grows only when the newly added letter's real count in the current window beats every earlier count, so at that moment the window truly needs at most k changes. While max_freq is stale, the window keeps the same length and just slides, so it never records a length that wasn't already achieved by a valid window."},
			{Prompt: "Suppose the question asked you to return the substring itself (or the number of valid windows) instead of the best length. Would the stale max_freq trick still be safe?", Answer: "No. The trick only works because we care about the maximum length, so a window that is temporarily invalid but no longer than the best is harmless. To return the substring or count valid windows you need every recorded window to be truly valid, so recompute the real maximum count (at most 26 letters) when shrinking, or use the one-window-per-letter version."},
		},
	},
	{
		ID:              "permutation-in-string",
		Number:          567,
		Order:           4,
		Title:           "Permutation in String",
		Slug:            "permutation-in-string",
		Difficulty:      "Medium",
		Category:        "sliding-window",
		Tags:            []string{"Hash Table", "Two Pointers", "String", "Sliding Window"},
		RelatedChapters: []int{4, 2},
		Description: `Given two strings s1 and s2, return true if some substring of s2 is a rearrangement (permutation) of s1, and false otherwise.

A permutation of s1 uses exactly the same letters as s1, each the same number of times, in any order. The substring must be contiguous. If s1 is longer than s2, the answer is false.`,
		Constraints: []string{
			"1 <= s1.length, s2.length <= 10^4",
			"s1 and s2 consist of lowercase English letters",
		},
		Examples: []Example{
			{Input: `s1 = "ab", s2 = "eidbaooo"`, Output: "true", Explanation: `The substring "ba" (indices 3-4) has one a and one b, the same letters as "ab".`},
			{Input: `s1 = "ab", s2 = "eidboaoo"`, Output: "false", Explanation: `b and a both appear, but never next to each other.`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s1": "ab", "s2": "eidbaooo"}, Expected: true, Note: "example: match in the middle"},
			{Input: map[string]interface{}{"s1": "ab", "s2": "eidboaoo"}, Expected: false, Note: "example: the letters are present but not contiguous"},
			{Input: map[string]interface{}{"s1": "abc", "s2": "ab"}, Expected: false, Hidden: true, Note: "s1 longer than s2"},
			{Input: map[string]interface{}{"s1": "a", "s2": "a"}, Expected: true, Hidden: true, Note: "both single characters"},
			{Input: map[string]interface{}{"s1": "ab", "s2": "bacc"}, Expected: true, Hidden: true, Note: "match is the very first window: check before sliding"},
			{Input: map[string]interface{}{"s1": "ab", "s2": "ccccab"}, Expected: true, Hidden: true, Note: "match is the very last window: check after the final slide"},
			{Input: map[string]interface{}{"s1": "aab", "s2": "abab"}, Expected: true, Hidden: true, Note: "duplicates in s1: counts matter, not just which letters"},
			{Input: map[string]interface{}{"s1": "aa", "s2": "ab"}, Expected: false, Hidden: true, Note: "need two a's, s2 has one"},
			{Input: map[string]interface{}{"s1": "aab", "s2": "xabbx"}, Expected: false, Hidden: true, Note: "the window \"abb\" has the same letters as \"aab\" but different counts"},
			{Input: map[string]interface{}{"s1": "hello", "s2": "ooolleoooleh"}, Expected: false, Hidden: true, Note: "all letters of s1 appear, scattered"},
			{Input: map[string]interface{}{"s1": "adc", "s2": "dcda"}, Expected: true, Hidden: true, Note: "the match overlaps a non-match window"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(1)",
		StarterCode:     "def checkInclusion(s1, s2):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "How long must a substring of s2 be to be a rearrangement of s1? And how could you tell two strings are rearrangements of each other without trying every order?"},
			{Level: 2, Type: "insight", Content: "Two strings are permutations of each other exactly when their letter counts are equal. Every candidate substring has length len(s1), so this is a fixed-size sliding window (Chapter 4): slide a window of len(s1) across s2, adding the entering letter's count and subtracting the leaving letter's count, and compare the window's counts with s1's counts.", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Set up the window state: if s1 is longer than s2, return False. Count the letters of s1 into need (26 slots) and start an empty window count.\n2. Slide the window: for each right, add s2[right]; once right >= len(s1), drop s2[right - len(s1)], the letter that just fell out.\n3. Record the answer: if the window counts equal need, return True.\n4. After the loop, return False."},
			{Level: 4, Type: "skeleton", Content: "def checkInclusion(s1, s2):\n    m = len(s1)\n    if m > len(s2):\n        return False\n    need = [0] * 26\n    window = [0] * 26\n    for ch in s1:\n        need[ord(ch) - ord('a')] += 1\n    for right in range(len(s2)):\n        window[ord(s2[right]) - ord('a')] += 1\n        if ____:\n            window[ord(____) - ord('a')] -= 1\n        if ____:\n            return True\n    return False"},
		},
		Solution: Solution{
			Code: `def checkInclusion(s1, s2):
    m = len(s1)
    if m > len(s2):
        return False
    need = [0] * 26     # letter counts of s1
    window = [0] * 26   # letter counts of s2[right-m+1..right]
    for ch in s1:
        need[ord(ch) - ord('a')] += 1
    for right in range(len(s2)):
        window[ord(s2[right]) - ord('a')] += 1
        if right >= m:
            window[ord(s2[right - m]) - ord('a')] -= 1
        if window == need:
            return True
    return False`,
			Explanation: "This keeps the letter counts from Longest Repeating Character Replacement but changes the window type: every candidate has exactly len(s1) letters, so the window is fixed-size. Each step adds one letter on the right and removes the one that fell off the left, and then compares two 26-slot count lists. No shrink-until-valid loop is needed because the size never changes.",
			BruteForce:  "For each of the len(s2) - len(s1) + 1 start positions, sort the substring and compare it with sorted(s1). With m = len(s1) and n = len(s2) that is O(n * m log m). The waste: neighbouring windows share m - 1 letters, yet each one is re-sorted from scratch.",
			KeyInsight:  "Permutation means equal letter counts, and the counts of the next window differ from the current one by exactly one added and one removed letter. So the window's counts can be updated in O(1) per step.",
			WhyItWorks:  "Invariant: after processing index right, window holds the letter counts of s2[max(0, right-m+1)..right]. Adding s2[right] extends the window, and once right >= m, subtracting s2[right - m] removes the letter that is now m+1 positions back, so the window has exactly m letters. Every length-m substring of s2 ends at some right >= m - 1, so each one is compared with need exactly once; two strings are permutations of each other if and only if their counts are equal, so the function returns True exactly when a matching substring exists. Comparing two lists of 26 counts is constant work, so the total is O(n).",
			CommonMistakes: []string{
				"Comparing only which letters appear (sets) instead of how many times. For s1 = \"aab\", the window \"abb\" has the same set {a, b} but is not a permutation. Compare counts.",
				"Removing the wrong letter when sliding: the letter leaving is s2[right - m], not s2[right - m + 1]. An off-by-one here makes the window m - 1 long and misses matches.",
				"Checking for a match only after the first full slide, or only before the loop ends. Compare after every step so the first window (\"bacc\") and the last window (\"ccccab\") are both checked.",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(1)",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the window state", Explanation: "A substring can't be longer than s2. need holds s1's counts; window will hold the counts of the current length-m window. 26 slots each, so O(1) space.", CodeSnippet: "if m > len(s2):\n    return False\nneed = [0] * 26     # letter counts of s1", LineStart: 3, LineEnd: 8},
				{Title: "Slide the window", Explanation: "Add the entering letter; once the window would exceed m letters, subtract the one that fell out. For s2 = \"eidbaooo\", m = 2, at right = 4 we add a and remove d (index 2), leaving {b:1, a:1}.", CodeSnippet: "window[ord(s2[right]) - ord('a')] += 1\nif right >= m:\n    window[ord(s2[right - m]) - ord('a')] -= 1", LineStart: 10, LineEnd: 12},
				{Title: "Record the answer", Explanation: "If the window has exactly s1's counts, this window is a permutation of s1, so stop early.", CodeSnippet: "if window == need:\n    return True", LineStart: 13, LineEnd: 14},
				{Title: "No window matched", Explanation: "Every length-m window was compared once; none matched.", CodeSnippet: "return False", LineStart: 15, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: sort every window",
					Code: `def checkInclusion(s1, s2):
    target = sorted(s1)
    m = len(s1)
    for start in range(len(s2) - m + 1):
        if sorted(s2[start:start + m]) == target:
            return True
    return False`,
					TimeComplexity:  "O(n * m log m)",
					SpaceComplexity: "O(m)",
					Explanation:     "Sorting puts any permutation into the same order, so sorted strings are equal exactly for permutations. Correct but recomputes each window from scratch.",
				},
				{
					Name: "Count how many letters match",
					Code: `def checkInclusion(s1, s2):
    m = len(s1)
    if m > len(s2):
        return False
    need = [0] * 26
    window = [0] * 26
    for ch in s1:
        need[ord(ch) - ord('a')] += 1
    matches = sum(1 for i in range(26) if need[i] == window[i])

    def change(i, delta):
        nonlocal matches
        if window[i] == need[i]:
            matches -= 1
        window[i] += delta
        if window[i] == need[i]:
            matches += 1

    for right in range(len(s2)):
        change(ord(s2[right]) - ord('a'), 1)
        if right >= m:
            change(ord(s2[right - m]) - ord('a'), -1)
        if matches == 26:
            return True
    return False`,
					TimeComplexity:  "O(n)",
					SpaceComplexity: "O(1)",
					Explanation:     "Keeps a running count of how many of the 26 letters have the right count, so each step is a few operations instead of a 26-slot comparison. This is the same bookkeeping idea as formed/required in Minimum Window Substring.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is this a fixed-size window rather than a grow/shrink window like Longest Substring Without Repeating Characters?", Answer: "Because the length of a valid answer is known in advance: a permutation of s1 has exactly len(s1) letters. There is nothing to optimise about the length, so the window keeps that size and just slides, adding one letter and removing one each step. Variable windows are for problems where the best length is the unknown."},
			{Prompt: "What changes if you must return the start index of every substring of s2 that is a permutation of s1 (LeetCode 438, Find All Anagrams)?", Answer: "Almost nothing: instead of returning True at the first match, append right - m + 1 (the window's start) to a list whenever window == need, and return the list at the end. The sliding, adding and removing stay identical."},
		},
	},
	{
		ID:              "minimum-window-substring",
		Number:          76,
		Order:           5,
		Title:           "Minimum Window Substring",
		Slug:            "minimum-window-substring",
		Difficulty:      "Hard",
		Category:        "sliding-window",
		Tags:            []string{"Hash Table", "String", "Sliding Window"},
		RelatedChapters: []int{4, 2},
		Description: `Given strings s and t, return the shortest substring of s that contains every character of t, counting duplicates (if t has two a's, the substring needs at least two a's). Characters are case-sensitive: "a" and "A" are different.

If no substring of s contains all of t, return the empty string "". The tests are built so that the shortest window is unique.`,
		Constraints: []string{
			"1 <= s.length, t.length <= 10^5",
			"s and t consist of uppercase and lowercase English letters",
			"The answer, when it exists, is unique",
		},
		Examples: []Example{
			{Input: `s = "ADOBECODEBANC", t = "ABC"`, Output: `"BANC"`, Explanation: `"ADOBEC" (length 6) also contains A, B and C, but "BANC" (length 4) is the shortest window that does.`},
			{Input: `s = "a", t = "a"`, Output: `"a"`},
			{Input: `s = "a", t = "aa"`, Output: `""`, Explanation: `t needs two a's but s has only one.`},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "ADOBECODEBANC", "t": "ABC"}, Expected: "BANC", Note: "example: the window must shrink past several useless letters"},
			{Input: map[string]interface{}{"s": "a", "t": "a"}, Expected: "a", Note: "whole string is the answer"},
			{Input: map[string]interface{}{"s": "a", "t": "aa"}, Expected: "", Note: "duplicates in t: one a is not enough"},
			{Input: map[string]interface{}{"s": "abc", "t": "d"}, Expected: "", Hidden: true, Note: "a character of t never appears"},
			{Input: map[string]interface{}{"s": "AaB", "t": "ab"}, Expected: "", Hidden: true, Note: "case-sensitive: B is not b"},
			{Input: map[string]interface{}{"s": "bba", "t": "ab"}, Expected: "ba", Hidden: true, Note: "an extra b must be dropped from the left"},
			{Input: map[string]interface{}{"s": "ab", "t": "b"}, Expected: "b", Hidden: true, Note: "answer is at the end"},
			{Input: map[string]interface{}{"s": "aa", "t": "aa"}, Expected: "aa", Hidden: true, Note: "t is made of one repeated letter"},
			{Input: map[string]interface{}{"s": "cabwefgewcwaefgcf", "t": "cae"}, Expected: "cwae", Hidden: true, Note: "longer string; the shortest window appears late"},
			{Input: map[string]interface{}{"s": "aaflslflsldkalskaaa", "t": "aaa"}, Expected: "aaa", Hidden: true, Note: "count-based: three a's needed, only the last run has them close together"},
		},
		TimeComplexity:  "O(m + n)",
		SpaceComplexity: "O(m + n)",
		StarterCode:     "def minWindow(s, t):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Once a window s[left..right] contains all of t, would extending it further right ever give a shorter answer? What should you try instead?"},
			{Level: 2, Type: "insight", Content: "Validity is monotonic again: if a window covers t, every bigger window does too. So grow right until the window covers t, then shrink left while it still covers t, recording the length at each step. To test 'covers t' in O(1), count how many distinct characters already have enough copies (formed) and compare with how many distinct characters t has (required). This is the 'shortest valid window' form of the variable-size sliding window (Chapter 4).", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Set up the window state: need = counts of t; required = number of distinct characters in t; window counts empty; formed = 0.\n2. Grow the window: add s[right]; if its window count just reached its need, formed += 1.\n3. Shrink while valid: while formed == required, record the window if it is the shortest so far, then remove s[left]; if that drops its count below its need, formed -= 1; move left forward.\n4. Record the answer: return the best window, or \"\" if none was found."},
			{Level: 4, Type: "skeleton", Content: "def minWindow(s, t):\n    need = {}\n    for ch in t:\n        need[ch] = need.get(ch, 0) + 1\n    required = len(need)\n    window = {}\n    formed = 0\n    left = 0\n    best_len, best_left = len(s) + 1, 0\n    for right in range(len(s)):\n        ch = s[right]\n        window[ch] = window.get(ch, 0) + 1\n        if ch in need and ____:\n            formed += 1\n        while ____:\n            if right - left + 1 < best_len:\n                best_len, best_left = right - left + 1, left\n            out = s[left]\n            window[out] -= 1\n            if out in need and ____:\n                formed -= 1\n            left += 1\n    return \"\" if best_len > len(s) else s[best_left:best_left + best_len]"},
		},
		Solution: Solution{
			Code: `def minWindow(s, t):
    need = {}
    for ch in t:
        need[ch] = need.get(ch, 0) + 1
    required = len(need)   # distinct characters that must be satisfied
    window = {}
    formed = 0             # distinct characters with enough copies now
    left = 0
    best_len, best_left = len(s) + 1, 0
    for right in range(len(s)):
        ch = s[right]
        window[ch] = window.get(ch, 0) + 1
        if ch in need and window[ch] == need[ch]:
            formed += 1
        while formed == required:
            if right - left + 1 < best_len:
                best_len, best_left = right - left + 1, left
            out = s[left]
            window[out] -= 1
            if out in need and window[out] < need[out]:
                formed -= 1
            left += 1
    if best_len > len(s):
        return ""
    return s[best_left:best_left + best_len]`,
			Explanation: "This combines two earlier rungs. From Permutation in String it takes count matching against a target (need), and from Longest Substring Without Repeating Characters it takes the grow/shrink window. The difference is the goal: we want the shortest valid window, so the inner loop runs while the window is valid (recording, then shrinking) instead of while it is invalid. The formed/required pair turns 'does the window cover t?' into one integer comparison.",
			BruteForce:  "For each start i, extend j until s[i..j] covers t, checking coverage with a count comparison, and keep the shortest. That is O(n^2) windows, each with a coverage check, which is far too slow for n = 10^5. The waste: after start i finds its first covering window ending at j, start i + 1 can't end before j, yet it scans from i + 1 again.",
			KeyInsight:  "If s[left..right] covers t, then so does any window containing it, and the shortest covering window ending at right has the largest possible left. Both pointers can only move forward, and a running counter (formed) tells in O(1) whether all required characters are covered.",
			WhyItWorks:  "formed counts the distinct characters c of t with window[c] >= need[c]. It changes only when a count crosses its need: formed += 1 exactly when window[c] rises to need[c], and formed -= 1 exactly when window[c] falls below need[c]. Extra copies (window[c] above need[c]) never change it. So formed == required holds exactly when the window covers t. For each right, the while loop records every covering window ending at right, from the current left to the largest left that still covers t, so the shortest covering window ending at right is considered. Left never needs to move back: a window starting before the current left that ends at a later right is longer than one we already recorded. Each index is added once and removed once: O(m + n).",
			CommonMistakes: []string{
				"Incrementing formed on every matching character instead of only when window[c] reaches need[c] exactly. With t = \"aa\" and s = \"a\", counting each a as 'formed' would report a cover after one a. Increment only on the == crossing.",
				"Decrementing formed whenever a character of t leaves, even if extra copies remain. For s = \"bba\", t = \"ab\", dropping the first b leaves another b, so the window \"ba\" is still valid. Decrement only when window[c] drops below need[c].",
				"Recording the window after shrinking instead of before. The window is valid at the top of the while loop; after removing s[left] it may no longer be. Record first, then remove.",
				"Treating letters case-insensitively. \"B\" does not satisfy \"b\".",
			},
			TimeComplexity:  "O(m + n)",
			SpaceComplexity: "O(m + n)",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the window state", Explanation: "need counts t. For t = \"ABC\", required = 3: three distinct characters must each have enough copies. formed starts at 0 because the window is empty.", CodeSnippet: "required = len(need)   # distinct characters that must be satisfied\nformed = 0             # distinct characters with enough copies now", LineStart: 2, LineEnd: 9},
				{Title: "Grow the window", Explanation: "Add s[right]. If its count just reached the amount t needs, one more requirement is satisfied. In \"ADOBECODEBANC\", formed reaches 3 at right = 5 (C), giving the window \"ADOBEC\".", CodeSnippet: "window[ch] = window.get(ch, 0) + 1\nif ch in need and window[ch] == need[ch]:\n    formed += 1", LineStart: 10, LineEnd: 14},
				{Title: "Shrink while valid", Explanation: "While every requirement is met, record the window if it is the shortest so far, then drop s[left]. If that leaves too few copies of a needed character, formed drops and the loop ends.", CodeSnippet: "while formed == required:\n    if right - left + 1 < best_len:\n        best_len, best_left = right - left + 1, left\nout = s[left]\nwindow[out] -= 1\nif out in need and window[out] < need[out]:\n    formed -= 1\nleft += 1", LineStart: 15, LineEnd: 22},
				{Title: "Record the answer", Explanation: "best_len starts one larger than any real window, so if it never changed, no window covered t and the answer is \"\". Otherwise slice out the best window.", CodeSnippet: "if best_len > len(s):\n    return \"\"\nreturn s[best_left:best_left + best_len]", LineStart: 23, LineEnd: 25},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: shortest cover from every start",
					Code: `from collections import Counter

def minWindow(s, t):
    need = Counter(t)
    best = ""
    for i in range(len(s)):
        for j in range(i, len(s)):
            piece = s[i:j + 1]
            if not (need - Counter(piece)):
                if best == "" or len(piece) < len(best):
                    best = piece
                break
    return best`,
					TimeComplexity:  "O(n^3)",
					SpaceComplexity: "O(n)",
					Explanation:     "need - Counter(piece) is empty exactly when piece has at least as many of each character as t. Stopping at the first cover for each start keeps it simple. Fine for tests, hopeless for n = 10^5.",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does formed count distinct characters (required = len(need)) instead of total characters (len(t))?", Answer: "Because coverage is decided per character: the window covers t when, for every distinct character c, window[c] >= need[c]. Counting distinct characters lets formed change only when a character crosses its threshold, so extra copies can't make it look covered too early. (A variant counts total needed characters still missing and decrements only while window[c] <= need[c]; it works too, but the bookkeeping rule is the same idea.)"},
			{Prompt: "Longest Substring Without Repeating Characters shrinks while the window is invalid; this problem shrinks while it is valid. How do you decide which one a problem needs?", Answer: "Ask which way validity moves and what you optimise. If you want the longest window and bigger windows break validity (repeats, too many changes), grow freely and shrink while invalid, recording after the loop. If you want the shortest window and bigger windows keep validity (covering t, sum at least a target), grow until valid and shrink while valid, recording inside the loop."},
		},
	},
	{
		ID:              "sliding-window-maximum",
		Number:          239,
		Order:           6,
		Title:           "Sliding Window Maximum",
		Slug:            "sliding-window-maximum",
		Difficulty:      "Hard",
		Category:        "sliding-window",
		Tags:            []string{"Array", "Queue", "Sliding Window", "Monotonic Queue"},
		RelatedChapters: []int{4, 5},
		Description: `You are given an integer list nums and a window size k. A window of k consecutive elements starts at the left end of nums and moves right one position at a time until it reaches the right end.

Return a list with the maximum of each window, in order. There are len(nums) - k + 1 windows.`,
		Constraints: []string{
			"1 <= nums.length <= 10^5",
			"-10^4 <= nums[i] <= 10^4",
			"1 <= k <= nums.length",
		},
		Examples: []Example{
			{Input: "nums = [1,3,-1,-3,5,3,6,7], k = 3", Output: "[3,3,5,5,6,7]", Explanation: "Windows: [1,3,-1] max 3, [3,-1,-3] max 3, [-1,-3,5] max 5, [-3,5,3] max 5, [5,3,6] max 6, [3,6,7] max 7."},
			{Input: "nums = [1], k = 1", Output: "[1]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 3, -1, -3, 5, 3, 6, 7}, "k": 3}, Expected: []int{3, 3, 5, 5, 6, 7}, Note: "example: the maximum changes and repeats"},
			{Input: map[string]interface{}{"nums": []int{1}, "k": 1}, Expected: []int{1}, Note: "single element"},
			{Input: map[string]interface{}{"nums": []int{1, -1}, "k": 1}, Expected: []int{1, -1}, Hidden: true, Note: "k = 1: each element is its own window"},
			{Input: map[string]interface{}{"nums": []int{3, 1, 2}, "k": 3}, Expected: []int{3}, Hidden: true, Note: "k = n: one window"},
			{Input: map[string]interface{}{"nums": []int{9, 8, 7, 6, 5}, "k": 2}, Expected: []int{9, 8, 7, 6}, Hidden: true, Note: "decreasing input: the old maximum must expire from the front"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5}, "k": 3}, Expected: []int{3, 4, 5}, Hidden: true, Note: "increasing input: each new element clears the deque"},
			{Input: map[string]interface{}{"nums": []int{4, 4, 4, 4}, "k": 2}, Expected: []int{4, 4, 4}, Hidden: true, Note: "duplicates: equal values must not break expiry"},
			{Input: map[string]interface{}{"nums": []int{-7, -8, 7, 5, 7, 1, 6, 0}, "k": 4}, Expected: []int{7, 7, 7, 7, 7}, Hidden: true, Note: "negatives and a repeated maximum"},
			{Input: map[string]interface{}{"nums": []int{1, 3, 1, 2, 0, 5}, "k": 3}, Expected: []int{3, 3, 2, 5}, Hidden: true, Note: "after 3 expires, a smaller element kept in the deque (2) becomes the maximum"},
		},
		TimeComplexity:  "O(n)",
		SpaceComplexity: "O(k)",
		StarterCode:     "def maxSlidingWindow(nums, k):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In the window [1, 3, -1], a 5 arrives. Can 1, 3 or -1 ever be the maximum of any later window that also contains the 5? Which elements are worth remembering at all?"},
			{Level: 2, Type: "insight", Content: "An element that has a larger element to its right inside the window can never be a maximum again: the larger one stays in the window at least as long. So keep only the 'still possible' candidates, which are in decreasing order of value from front to back. A deque of indices does this (a monotonic deque): pop smaller values off the back when a new element arrives, and pop the front when its index slides out. The front is always the window's maximum. Chapter 4 covers the fixed window; Chapter 5 explains monotonic stacks, the same 'pop what can never matter again' idea.", ChapterRef: 4},
			{Level: 3, Type: "plan", Content: "1. Set up the window state: an empty deque of indices whose values decrease from front to back.\n2. Grow the window: before appending index right, pop indices from the back whose values are <= nums[right]; they can never be a maximum again. Append right.\n3. Slide the window: if the front index is <= right - k, it has left the window; pop it from the front.\n4. Record the answer: once right >= k - 1, the front index holds the window maximum; append nums[front]."},
			{Level: 4, Type: "skeleton", Content: "from collections import deque\n\ndef maxSlidingWindow(nums, k):\n    candidates = deque()\n    result = []\n    for right in range(len(nums)):\n        while candidates and ____:\n            candidates.pop()\n        candidates.append(right)\n        if ____:\n            candidates.popleft()\n        if ____:\n            result.append(nums[candidates[0]])\n    return result"},
		},
		Solution: Solution{
			Code: `from collections import deque

def maxSlidingWindow(nums, k):
    candidates = deque()  # indices; their values decrease front to back
    result = []
    for right in range(len(nums)):
        while candidates and nums[candidates[-1]] <= nums[right]:
            candidates.pop()
        candidates.append(right)
        if candidates[0] <= right - k:
            candidates.popleft()
        if right >= k - 1:
            result.append(nums[candidates[0]])
    return result`,
			Explanation: "This is a fixed-size window like Permutation in String, but the window state is harder: a maximum can't be updated by 'add the new, subtract the old' because when the maximum leaves you don't know the next largest. The fix is a monotonic deque: a double-ended queue of indices whose values strictly decrease from front to back. New elements knock out every smaller element at the back (those can never be a maximum again), and the front falls off when it slides out of the window. The front is always the current maximum. Chapter 5 teaches the same idea for stacks (Daily Temperatures); a deque adds the ability to expire old items from the front.",
			BruteForce:  "Compute max(nums[i:i+k]) for every window: (n - k + 1) windows of k elements, O(n*k), about 5 * 10^9 operations for n = 10^5 and k = 5 * 10^4. The waste: consecutive windows share k - 1 elements, but each max is computed from scratch.",
			KeyInsight:  "If j < i and nums[j] <= nums[i], then j can never be the maximum of a window that contains i, and every future window that contains j also contains i (j leaves first). So j can be discarded forever. What remains is a decreasing sequence, and its first element is the maximum.",
			WhyItWorks:  "Invariant after each step: the deque holds, in increasing index order, exactly the indices in the current window that have no element >= them to their right inside the window, and their values strictly decrease from front to back. Popping smaller-or-equal values from the back when right arrives keeps this true (those indices now have nums[right] to their right). Popping the front when its index is right - k removes the one index that just left the window; only the front can be that old, because indices are in increasing order. The window's maximum has nothing >= it to its right (taking the rightmost copy of the maximum), so it is in the deque, and since values decrease, it is at the front. Each index is appended once and popped at most once, so the total work is O(n) and the deque never holds more than k indices.",
			CommonMistakes: []string{
				"Storing values instead of indices. With values you can't tell when the front has left the window (especially with duplicates like [4,4,4,4]). Store indices and compare nums[index].",
				"Checking expiry with < instead of <=: the front index is outside the window [right-k+1, right] when it is <= right - k. For [9,8,7,6,5], k = 2, an off-by-one keeps 9 for too long.",
				"Using a heap without removing expired items. A max-heap gives the largest value ever seen, which may be outside the window; you must pop entries whose index is too old (lazy deletion), which costs O(n log n).",
				"Appending to the result before the first full window (right < k - 1). The first answer belongs to window [0, k-1].",
			},
			TimeComplexity:  "O(n)",
			SpaceComplexity: "O(k)",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the window state", Explanation: "The deque holds indices of elements that could still be a window maximum, with decreasing values from front to back.", CodeSnippet: "candidates = deque()  # indices; their values decrease front to back", LineStart: 4, LineEnd: 5},
				{Title: "Grow the window", Explanation: "A new element makes every smaller-or-equal element behind it useless, so pop them from the back, then append. For [1,3,-1,-3,5], when 5 arrives the deque [3,-1,-3] (as values) is emptied.", CodeSnippet: "while candidates and nums[candidates[-1]] <= nums[right]:\n    candidates.pop()\ncandidates.append(right)", LineStart: 7, LineEnd: 9},
				{Title: "Slide the window", Explanation: "The window is [right - k + 1, right]. If the front index is right - k or smaller, it has slid out; pop it. At most one index expires per step.", CodeSnippet: "if candidates[0] <= right - k:\n    candidates.popleft()", LineStart: 10, LineEnd: 11},
				{Title: "Record the answer", Explanation: "Once the first full window exists, the front of the deque is its maximum.", CodeSnippet: "if right >= k - 1:\n    result.append(nums[candidates[0]])", LineStart: 12, LineEnd: 13},
			},
			Alternatives: []AltSolution{
				{
					Name: "Brute force: max of every window",
					Code: `def maxSlidingWindow(nums, k):
    result = []
    for start in range(len(nums) - k + 1):
        result.append(max(nums[start:start + k]))
    return result`,
					TimeComplexity:  "O(n*k)",
					SpaceComplexity: "O(k)",
					Explanation:     "Direct and obviously correct; too slow when both n and k are large.",
				},
				{
					Name: "Max-heap with lazy deletion",
					Code: `import heapq

def maxSlidingWindow(nums, k):
    heap = []
    result = []
    for right, value in enumerate(nums):
        heapq.heappush(heap, (-value, right))
        while heap[0][1] <= right - k:
            heapq.heappop(heap)
        if right >= k - 1:
            result.append(-heap[0][0])
    return result`,
					TimeComplexity:  "O(n log n)",
					SpaceComplexity: "O(n)",
					Explanation:     "Python's heapq is a min-heap, so values are negated. Expired entries stay in the heap until they reach the top, then get popped. Simpler to invent than the deque, slightly slower (Chapter 11 covers heaps).",
				},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does popping from the back use <= rather than <, and would < also be correct?", Answer: "With <=, an older equal value is removed because the newer copy will stay in the window at least as long, so the older one is redundant. With <, both copies stay; the front still holds a maximum and expiry still works because we store indices, so the answer is also correct. <= just keeps the deque smaller. What would break is storing values instead of indices."},
			{Prompt: "Why is the whole algorithm O(n) even though the while loop can pop many elements in one step?", Answer: "Each index is appended to the deque exactly once and can be popped (from the back or the front) at most once. So across all n steps there are at most n appends and n pops in total: amortised O(1) per step, O(n) overall. One step may pop many, but only because earlier steps appended them."},
			{Prompt: "How would you change the code to return the window minimums instead?", Answer: "Flip the comparison so the deque keeps values increasing from front to back: pop from the back while nums[candidates[-1]] >= nums[right]. The front then holds the window minimum. Expiry and recording stay the same."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, SlidingWindowProblems...)
}
