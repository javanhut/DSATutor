package practice

import (
	"sort"
	"strings"
)

// BacktrackingDrills are the backtracking foundations: the bare choose,
// explore, un-choose skeleton (binary strings), a choice that only branches
// sometimes (letter case), the start index that stops reordered duplicates
// (combinations), a used[] array that lets any unused item come next
// (arrangements), pruning a branch that can no longer succeed (subsets with a
// target sum), and restoring shared state after a branch (paths through a
// grid), before the backtracking ladder applies them. The target-sum drill has
// a 40-number case with a 5 second limit that only a pruned search finishes.
var BacktrackingDrills = []*Problem{
	{
		ID:              "drill-binary-strings",
		Number:          0,
		Kind:            KindDrill,
		Order:           1,
		Title:           "All Binary Strings",
		Slug:            "drill-binary-strings",
		Difficulty:      "Easy",
		Category:        "backtracking",
		Tags:            []string{"Backtracking", "Recursion", "String", "Foundations"},
		RelatedChapters: []int{12, 8},
		Description:     "Given an integer `n`, return every string of exactly `n` characters, each of them `\"0\"` or `\"1\"`, so `2^n` strings in all.\n\nReturn them in **increasing order**, as if counting in binary: `\"00\"`, `\"01\"`, `\"10\"`, `\"11\"` for `n = 2`. The order is part of the task here: it is what you get when you try `\"0\"` before `\"1\"` at every position, and it shows you the shape of the search.\n\nBuild each string one character at a time on a shared `path` list: **choose** a digit and append it, **explore** the rest of the positions with a recursive call, then **un-choose** by removing the digit so the next digit starts from the same `path`. When `path` has `n` digits, record it and return. This choose, explore, un-choose loop is the skeleton of every backtracking problem.",
		Constraints:     []string{"1 <= n <= 10"},
		Examples: []Example{
			{Input: "n = 2", Output: "[\"00\",\"01\",\"10\",\"11\"]", Explanation: "At position 0 try 0, then at position 1 try 0 and 1 (\"00\", \"01\"); undo the second digit, undo the first, try 1 at position 0 and repeat (\"10\", \"11\")."},
			{Input: "n = 1", Output: "[\"0\",\"1\"]", Explanation: "One position with two choices."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 2}, Expected: []string{"00", "01", "10", "11"}, Note: "example 1: two positions, four strings in counting order"},
			{Input: map[string]interface{}{"n": 1}, Expected: []string{"0", "1"}, Note: "example 2: a single position"},
			{Input: map[string]interface{}{"n": 3}, Expected: btBinary(3), Hidden: true, Note: "eight strings: each digit must be removed before the next is tried"},
			{Input: map[string]interface{}{"n": 4}, Expected: btBinary(4), Hidden: true, Note: "sixteen strings: the order must stay 0 before 1 at every position"},
			{Input: map[string]interface{}{"n": 6}, Expected: btBinary(6), Hidden: true, Note: "64 strings, none repeated or cut short"},
			{Input: map[string]interface{}{"n": 10}, Expected: btBinary(10), Hidden: true, Note: "the largest size: 1024 strings of length 10"},
		},
		TimeComplexity:  "O(n * 2^n)",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def binaryStrings(n):\n    # Return every binary string of length n, in increasing order\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Fill the first position with 0 and ask for every way to fill the rest. When you have them all, what do you have to do to the first position before you try 1?"},
			{Level: 2, Type: "insight", Content: "Keep one shared list `path`. At each position, for each digit: append it (choose), recurse for the next position (explore), then pop it (un-choose) so `path` is exactly as it was. When `path` has n digits it is a complete answer: record it and return. This choose, explore, un-choose loop is backtracking (Chapter 12), and the recursion is the one from Chapter 8.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: a result list and a path list.\n2. Stop: when path has n digits, join it into a string, append it to the result and return.\n3. Choose: for each digit in \"0\", then \"1\", append it to path.\n4. Explore: call backtrack() to fill the next position.\n5. Un-choose: pop the digit before trying the next one."},
			{Level: 4, Type: "skeleton", Content: "def binaryStrings(n):\n    result = []\n    path = []\n\n    def backtrack():\n        if len(path) == ____:\n            result.append(\"\".join(path))\n            ____\n        for digit in \"01\":\n            path.append(digit)   # choose\n            ____                 # explore\n            ____                 # un-choose\n\n    backtrack()\n    return result"},
		},
		Solution: Solution{
			Code: `def binaryStrings(n):
    result = []
    path = []

    def backtrack():
        if len(path) == n:
            result.append("".join(path))
            return
        for digit in "01":
            path.append(digit)
            backtrack()
            path.pop()

    backtrack()
    return result`,
			Explanation:     "This is the bare backtracking skeleton: a shared path, a base case that records a complete answer, and a loop that chooses, explores and un-chooses. Because each position tries 0 before 1, the strings come out in increasing order. Every later drill keeps this shape and changes only what the choices are, which ones are allowed, and what shared state must be restored.",
			KeyInsight:      "The recursion does the nesting you can't write with a loop: n nested for loops would be needed for n positions, but one function that calls itself once per position has the depth of the recursion equal to n. The pop after the call is what lets the one shared path serve every branch.",
			WhyItWorks:      "Invariant: when backtrack() is entered, path holds the digits chosen for positions 0..len(path)-1 and nothing else, and when it returns, path is exactly as it was on entry. The loop keeps that true, because every append is matched by a pop after the call that may have changed path and restored it. So each recursive call sees the correct prefix. The base case fires at depth n, once per way of picking a digit at each of the n positions, which is 2^n different paths in the tree, each recorded once; trying 0 before 1 at every level visits them in increasing order.",
			CommonMistakes:  []string{"Forgetting `path.pop()`. The un-choose step is what resets `path` for the next digit; without it the digits pile up and later strings are too long or built on leftovers.", "Recording the string but not returning. The loop then keeps adding digits past length n and never stops (RecursionError). The base case has to return.", "Trying digits in the wrong order, such as `\"10\"`, which gives the right strings in the wrong order. Here the order is part of the answer.", "Writing the base case as `len(path) > n`, so every string is built one digit too long; stop exactly when the length equals n."},
			TimeComplexity:  "O(n * 2^n): 2^n strings, each joined in O(n)",
			SpaceComplexity: "O(n) for the path and recursion depth, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the state", Explanation: "path is the string being built, one digit per slot, and result collects the finished strings. The inner function reads and changes both.", CodeSnippet: "result = []\npath = []", LineStart: 2, LineEnd: 3},
				{Title: "Record a complete answer", Explanation: "A path with n digits is a finished string. Join it into a string (which also copies it), save it and return so no more digits are added.", CodeSnippet: "if len(path) == n:\n    result.append(\"\".join(path))\n    return", LineStart: 6, LineEnd: 8},
				{Title: "Choose", Explanation: "Try 0 and then 1 for the next position by appending it to path.", CodeSnippet: "for digit in \"01\":\n    path.append(digit)", LineStart: 9, LineEnd: 10},
				{Title: "Explore", Explanation: "Let the recursive call fill every later position with this digit in place.", CodeSnippet: "backtrack()", LineStart: 11, LineEnd: 11},
				{Title: "Un-choose", Explanation: "Remove the digit so path is back to what it was before this choice, ready for the next digit. For n = 2 this is how \"00\" turns into \"01\" and then into \"10\".", CodeSnippet: "path.pop()", LineStart: 12, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{Name: "Counting in binary", Code: "def binaryStrings(n):\n    return [format(number, \"b\").zfill(n) for number in range(2 ** n)]", TimeComplexity: "O(n * 2^n)", SpaceComplexity: "O(n) besides the output", Explanation: "For this task the strings are just the numbers 0 to 2^n - 1 written in binary. It is not backtracking, so it will not carry over to problems whose choices depend on earlier choices, which is why the skeleton above is the one to learn."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "After the call `backtrack()` returns inside the loop, what is in `path`, and why does that matter?", Answer: "Exactly what it held before the loop appended the digit, because the deeper calls each undid their own choices. The pop then removes this call's digit. Every call can therefore assume path holds only its own prefix, no matter which branches ran before it."},
			{Prompt: "How many times is the base case reached, and how many times does the loop body run in total for n = 3?", Answer: "The base case is reached 8 times (2^3). The loop body runs once for every node of the tree below the root: 2 + 4 + 8 = 14 times. Every digit that is appended is later popped."},
		},
	},
	{
		ID:              "drill-letter-case-permutation",
		Number:          784,
		Kind:            KindDrill,
		Order:           2,
		Title:           "Letter Case Permutation",
		Slug:            "drill-letter-case-permutation",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Backtracking", "String", "Bit Manipulation", "Foundations"},
		RelatedChapters: []int{12, 8},
		Description:     "Given a string `s` of letters and digits, return every string you can make by changing each letter to lower case or upper case, independently, while the digits stay as they are. Return the strings in **any order**.\n\nA string with `L` letters has `2^L` results: digits give no choice, letters give two. So `\"a1b2\"` has four results, and `\"3z4\"` has two.\n\nUse the skeleton from the previous drill, but let the number of choices depend on the position: a digit has just one option (itself), a letter has two (lower and upper). The recursion walks the positions of `s` left to right with an index `i`; there is no start index to worry about here, because every position is used exactly once.",
		Constraints:     []string{"1 <= len(s) <= 12", "s consists of lower case letters, upper case letters and digits"},
		Examples: []Example{
			{Input: "s = \"a1b2\"", Output: "[\"a1b2\",\"a1B2\",\"A1b2\",\"A1B2\"]", Explanation: "Two letters give 2 * 2 = 4 results; the digits 1 and 2 never change."},
			{Input: "s = \"3z4\"", Output: "[\"3z4\",\"3Z4\"]", Explanation: "One letter gives two results."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"s": "a1b2"}, Expected: []string{"a1b2", "a1B2", "A1b2", "A1B2"}, Note: "example 1: two letters among digits"},
			{Input: map[string]interface{}{"s": "3z4"}, Expected: []string{"3z4", "3Z4"}, Note: "example 2: one letter"},
			{Input: map[string]interface{}{"s": "2020"}, Expected: []string{"2020"}, Hidden: true, Note: "digits only: exactly one result, no duplicates"},
			{Input: map[string]interface{}{"s": "7"}, Expected: []string{"7"}, Hidden: true, Note: "a single digit"},
			{Input: map[string]interface{}{"s": "q"}, Expected: []string{"q", "Q"}, Hidden: true, Note: "a single letter"},
			{Input: map[string]interface{}{"s": "aB"}, Expected: []string{"ab", "aB", "Ab", "AB"}, Hidden: true, Note: "an upper case letter in the input: a mixed-case string still gives each result once"},
			{Input: map[string]interface{}{"s": "C0d3"}, Expected: btCaseVariants("C0d3"), Hidden: true, Note: "letters and digits mixed, starting with an upper case letter"},
			{Input: map[string]interface{}{"s": "aBc1De2fG3h"}, Expected: btCaseVariants("aBc1De2fG3h"), Hidden: true, Note: "eight letters: 256 results, each letter flipped independently"},
		},
		TimeComplexity:  "O(n * 2^L)",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def letterCasePermutation(s):\n    # Return every string made by changing the case of letters in s\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "At the position of a digit, how many choices do you have? At the position of a letter?"},
			{Level: 2, Type: "insight", Content: "It is the same choose, explore, un-choose skeleton as before, but the list of options at position i depends on the character: just [s[i]] for a digit, [lower, upper] for a letter. The base case is when i reaches len(s). Offering two options for a digit would create duplicates. See Chapter 12.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: a result list, a path list, and backtrack(i) for position i.\n2. Stop: when i == len(s), join the path into a string, record it and return.\n3. Decide the options: the character itself, or its lower and upper case if it is a letter.\n4. Choose and explore: append the option, call backtrack(i + 1).\n5. Un-choose: pop the option before trying the next."},
			{Level: 4, Type: "skeleton", Content: "def letterCasePermutation(s):\n    result = []\n    path = []\n\n    def backtrack(i):\n        if i == ____:\n            result.append(\"\".join(path))\n            return\n        options = [s[i]]\n        if s[i].isalpha():\n            options = [s[i].lower(), ____]\n        for ch in options:\n            path.append(ch)   # choose\n            ____              # explore\n            ____              # un-choose\n\n    backtrack(0)\n    return result"},
		},
		Solution: Solution{
			Code: `def letterCasePermutation(s):
    result = []
    path = []

    def backtrack(i):
        if i == len(s):
            result.append("".join(path))
            return
        options = [s[i]]
        if s[i].isalpha():
            options = [s[i].lower(), s[i].upper()]
        for ch in options:
            path.append(ch)
            backtrack(i + 1)
            path.pop()

    backtrack(0)
    return result`,
			Explanation:     "This builds on All Binary Strings: the same skeleton, except the number of options depends on the position, and the recursion carries the position as an index. A digit contributes its single option; a letter contributes its two cases. Each leaf of the tree is one result.",
			KeyInsight:      "The tree branches only at letters. Treating the options as a list per position, one for a digit and two for a letter, keeps the loop the same for every position and creates each result once.",
			WhyItWorks:      "Invariant: when backtrack(i) is entered, path holds the chosen characters for positions 0..i-1, one per position, and on return it is as it was. A result is complete when i == len(s). Because exactly one option is taken per position and a digit has just one option, each distinct result corresponds to exactly one path in the tree, so every result appears once; all 2^L paths are explored because every letter offers both cases.",
			CommonMistakes:  []string{"Offering a lower and an upper option for every character, digits included. `\"1\".lower()` and `\"1\".upper()` are the same, so every result is returned twice per digit. Branch only when `isalpha()` is true.", "Offering only `[s[i], s[i].swapcase()]` is fine, but offering `[s[i].upper(), s[i]]` for a letter that is already upper case gives the same character twice. Use the lower and the upper form, whatever the input case.", "Forgetting `path.pop()`, so each result also carries the characters from the previous branches.", "Recording at every call rather than when `i == len(s)`: partial strings then show up in the answer."},
			TimeComplexity:  "O(n * 2^L) for L letters",
			SpaceComplexity: "O(n) for the path and recursion depth, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Record a complete answer", Explanation: "When every position has a character, the path is one finished string. Join, save and return.", CodeSnippet: "if i == len(s):\n    result.append(\"\".join(path))\n    return", LineStart: 6, LineEnd: 8},
				{Title: "Decide the options", Explanation: "A digit can only be itself. A letter can be lower or upper case, whichever its case in the input, so both forms are listed.", CodeSnippet: "options = [s[i]]\nif s[i].isalpha():\n    options = [s[i].lower(), s[i].upper()]", LineStart: 9, LineEnd: 11},
				{Title: "Choose", Explanation: "Append one option to the path for position i.", CodeSnippet: "for ch in options:\n    path.append(ch)", LineStart: 12, LineEnd: 13},
				{Title: "Explore", Explanation: "Recurse on the next position; i + 1 is the next slot to fill.", CodeSnippet: "backtrack(i + 1)", LineStart: 14, LineEnd: 14},
				{Title: "Un-choose", Explanation: "Pop the character so the other case can take its place.", CodeSnippet: "path.pop()", LineStart: 15, LineEnd: 15},
			},
			Alternatives: []AltSolution{
				{Name: "Build on the results so far", Code: "def letterCasePermutation(s):\n    results = [\"\"]\n    for ch in s:\n        if ch.isalpha():\n            results = [r + c for r in results for c in (ch.lower(), ch.upper())]\n        else:\n            results = [r + ch for r in results]\n    return results", TimeComplexity: "O(n * 2^L)", SpaceComplexity: "O(n * 2^L)", Explanation: "Instead of one shared path, extend every partial string at each position. It works for this task, but it keeps all partial results in memory and cannot prune a branch, so backtracking is the tool to learn."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does the input \"2020\" produce exactly one string, and which line makes sure of it?", Answer: "Every position is a digit, so each has a single option and the tree never branches. The `isalpha()` check keeps digits at one option; without it each digit would offer two equal options and the result would have 16 copies."},
			{Prompt: "What changes if you may also remove a letter (so \"a1\" gives \"a1\", \"A1\", \"1\")?", Answer: "A letter would have a third option, \"nothing\", which appends no character. The loop over options stays the same, but you would skip the append and the pop for that option, or append an empty string."},
		},
	},
	{
		ID:              "drill-combinations",
		Number:          77,
		Kind:            KindDrill,
		Order:           3,
		Title:           "Combinations",
		Slug:            "drill-combinations",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Backtracking", "Recursion", "Foundations"},
		RelatedChapters: []int{12, 8},
		Description:     "Given two integers `n` and `k`, return every way to choose `k` different numbers from `1, 2, ..., n`. Return the combinations in **any order**, and each combination in any order too, but each set of numbers only once.\n\nA combination is a *set*: `[1,2]` and `[2,1]` are the same combination, so only one of them belongs in the answer. There are `n! / (k! * (n - k)!)` of them: choosing 2 of 4 gives 6.\n\nThe fix that removes the reordered duplicates is a **start index**. After choosing a number, only numbers *larger* than it may come next, so every combination is built in increasing order and in exactly one way. Pass the next allowed number down as the argument of the recursive call.",
		Constraints:     []string{"1 <= k <= n <= 12"},
		Examples: []Example{
			{Input: "n = 4, k = 2", Output: "[[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]", Explanation: "After choosing 1 the next number can be 2, 3 or 4; after choosing 2 only 3 or 4 can follow, so [2,1] is never built."},
			{Input: "n = 1, k = 1", Output: "[[1]]", Explanation: "Only one number to choose."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"n": 4, "k": 2}, Expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}, Note: "example 1: six pairs, each only once"},
			{Input: map[string]interface{}{"n": 1, "k": 1}, Expected: [][]int{{1}}, Note: "example 2: the smallest input"},
			{Input: map[string]interface{}{"n": 5, "k": 5}, Expected: [][]int{{1, 2, 3, 4, 5}}, Hidden: true, Note: "k equals n: one combination, the whole range"},
			{Input: map[string]interface{}{"n": 5, "k": 1}, Expected: [][]int{{1}, {2}, {3}, {4}, {5}}, Hidden: true, Note: "k = 1: each number alone"},
			{Input: map[string]interface{}{"n": 5, "k": 3}, Expected: btCombos(5, 3), Hidden: true, Note: "ten triples: reordered copies such as [2,1,3] must not appear"},
			{Input: map[string]interface{}{"n": 6, "k": 4}, Expected: btCombos(6, 4), Hidden: true, Note: "fifteen combinations of four"},
			{Input: map[string]interface{}{"n": 9, "k": 4}, Expected: btCombos(9, 4), Hidden: true, Note: "126 combinations: a returned list must be a copy, not the shared path"},
			{Input: map[string]interface{}{"n": 12, "k": 6}, Expected: btCombos(12, 6), Hidden: true, Note: "924 combinations at the largest size"},
		},
		TimeComplexity:  "O(k * C(n, k))",
		SpaceComplexity: "O(k) besides the output",
		StarterCode:     "def combine(n, k):\n    # Return every combination of k numbers chosen from 1..n\n    pass",
		Harness:         Harness{Compare: "unordered-nested"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If you have already chosen 2, which numbers are you allowed to pick next if [2,1] and [1,2] must not both appear?"},
			{Level: 2, Type: "insight", Content: "Only allow numbers larger than the last one chosen: pass `start` down, loop from `start` to n, and recurse with `number + 1`. Every combination is then built in increasing order, which is the only way to build it. A path with k numbers is a complete answer. See Chapter 12.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: a result list, a path list, and backtrack(start).\n2. Stop: when path has k numbers, append a copy of it and return.\n3. Choose: for each number from start to n, append it.\n4. Explore: call backtrack(number + 1) so only larger numbers follow.\n5. Un-choose: pop it before trying the next number."},
			{Level: 4, Type: "skeleton", Content: "def combine(n, k):\n    result = []\n    path = []\n\n    def backtrack(start):\n        if len(path) == ____:\n            result.append(____)   # a copy, not path itself\n            return\n        for number in range(____, n + 1):\n            path.append(number)   # choose\n            backtrack(____)       # explore\n            ____                  # un-choose\n\n    backtrack(1)\n    return result"},
		},
		Solution: Solution{
			Code: `def combine(n, k):
    result = []
    path = []

    def backtrack(start):
        if len(path) == k:
            result.append(path[:])
            return
        for number in range(start, n + 1):
            path.append(number)
            backtrack(number + 1)
            path.pop()

    backtrack(1)
    return result`,
			Explanation:     "This builds on the skeleton of the first two drills and adds the start index. Choosing in increasing order means a set can only be built one way, so no reordered duplicates appear; the base case is a path of exactly k numbers. The same start index is what Subsets and Combination Sum use in the ladder.",
			BruteForce:      "Generate every ordered selection of k numbers, sort each one and keep it only if you haven't seen it. That builds k! times more lists than needed (for n = 12, k = 6 about 665,000 ordered picks for 924 sets) and needs a set to remove the copies. The waste is creating [2,1] when [1,2] already exists.",
			KeyInsight:      "A set has one increasing arrangement. If a number may only be followed by larger ones, each set is built along exactly one path, so there is nothing to filter out.",
			WhyItWorks:      "Invariant: when backtrack(start) runs, path holds increasing numbers all smaller than start, and the call records every way to extend it with k - len(path) numbers from start..n. A set {a1 < a2 < ... < ak} is reached by choosing a1, then a2 (which is allowed since a2 >= a1 + 1), and so on; it can't be reached in any other order, because the loop never goes below start. So each combination is recorded once. path[:] copies it, since the shared path changes afterwards.",
			CommonMistakes:  []string{"Looping from 1 every time (`range(1, n + 1)`) or recursing with `backtrack(start + 1)`. Both let smaller numbers return: you get [1,2] and [2,1], or repeated numbers like [1,1]. Loop from `start` and recurse with `number + 1`.", "Appending `path` instead of `path[:]`. Every saved answer is the same list object, which is empty when the recursion finishes, so the result is a list of empty lists.", "Forgetting `path.pop()`, so numbers from earlier branches stay in the path and later combinations are too long.", "Forgetting the base-case `return`: the loop keeps adding numbers to a path that already has k. The answer stays right, but every full path is extended pointlessly, and in a problem whose base case tests `<` or `==` on a sum it can record wrong answers."},
			TimeComplexity:  "O(k * C(n, k)): each combination is copied in O(k)",
			SpaceComplexity: "O(k) for the path and recursion depth, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the state", Explanation: "path holds the numbers chosen so far, in increasing order; result holds finished combinations.", CodeSnippet: "result = []\npath = []", LineStart: 2, LineEnd: 3},
				{Title: "Record a complete answer", Explanation: "With k numbers the combination is finished. Save a copy and stop; a longer path would be useless.", CodeSnippet: "if len(path) == k:\n    result.append(path[:])\n    return", LineStart: 6, LineEnd: 8},
				{Title: "Choose", Explanation: "Try every number from start up to n as the next one. Starting at start, not 1, is what keeps the numbers increasing.", CodeSnippet: "for number in range(start, n + 1):\n    path.append(number)", LineStart: 9, LineEnd: 10},
				{Title: "Explore", Explanation: "The next number must be larger than this one, so the new start is number + 1.", CodeSnippet: "backtrack(number + 1)", LineStart: 11, LineEnd: 11},
				{Title: "Un-choose", Explanation: "Pop the number so the next candidate takes its place in the same slot.", CodeSnippet: "path.pop()", LineStart: 12, LineEnd: 12},
			},
			Alternatives: []AltSolution{
				{Name: "itertools.combinations", Code: "from itertools import combinations\n\ndef combine(n, k):\n    return [list(c) for c in combinations(range(1, n + 1), k)]", TimeComplexity: "O(k * C(n, k))", SpaceComplexity: "O(k) besides the output", Explanation: "The standard library does the same enumeration. Handy to know and to check against, but the point of the drill is to be able to write the recursion yourself, because Combination Sum and Subsets II change its rules."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why do we recurse with number + 1 and not with start + 1?", Answer: "After choosing `number`, the next pick must come after it. start + 1 only moves one step past where this loop began, so if number is larger than start the numbers between start + 1 and number could be chosen again: you would build [3,2] after having built [2,3]."},
			{Prompt: "How could you stop early when too few numbers are left to reach k?", Answer: "If path has m numbers, you need k - m more, so the loop only has to run up to n - (k - m) + 1: any larger starting number leaves too few numbers after it. That prune doesn't change the answer but skips branches that can never finish."},
		},
	},
	{
		ID:              "drill-arrangements",
		Number:          0,
		Kind:            KindDrill,
		Order:           4,
		Title:           "Arrangements of k Letters",
		Slug:            "drill-arrangements",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Backtracking", "Recursion", "String", "Foundations"},
		RelatedChapters: []int{12, 8},
		Description:     "Given a string `letters` of **distinct** lower case letters and an integer `k`, return every string of exactly `k` letters that can be made by taking letters from `letters`, using each letter of `letters` at most once. Order matters: `\"ab\"` and `\"ba\"` are different results. Return them in **any order**.\n\nFor `letters = \"abc\"` and `k = 2` the results are `\"ab\"`, `\"ac\"`, `\"ba\"`, `\"bc\"`, `\"ca\"`, `\"cb\"`. With `k = len(letters)` you get every permutation.\n\nA start index no longer works: any unused letter may come next, including one that sits earlier in `letters`. Track which letters are in use with a boolean list `used`, set `used[i]` when you **choose** letter `i` and clear it again when you **un-choose**, so the flag describes exactly the letters currently on the path.",
		Constraints:     []string{"1 <= k <= len(letters) <= 7", "letters has distinct lower case letters"},
		Examples: []Example{
			{Input: "letters = \"abc\", k = 2", Output: "[\"ab\",\"ac\",\"ba\",\"bc\",\"ca\",\"cb\"]", Explanation: "3 choices for the first letter, then 2 left for the second: 3 * 2 = 6 results."},
			{Input: "letters = \"xy\", k = 1", Output: "[\"x\",\"y\"]", Explanation: "With k = 1 each letter alone is a result."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"letters": "abc", "k": 2}, Expected: []string{"ab", "ac", "ba", "bc", "ca", "cb"}, Note: "example 1: six ordered pairs"},
			{Input: map[string]interface{}{"letters": "xy", "k": 1}, Expected: []string{"x", "y"}, Note: "example 2: k = 1"},
			{Input: map[string]interface{}{"letters": "q", "k": 1}, Expected: []string{"q"}, Hidden: true, Note: "a single letter"},
			{Input: map[string]interface{}{"letters": "abc", "k": 3}, Expected: btArrangements("abc", 3), Hidden: true, Note: "k equals the length: all six permutations"},
			{Input: map[string]interface{}{"letters": "dog", "k": 2}, Expected: btArrangements("dog", 2), Hidden: true, Note: "letters not in alphabetical order: both orders of every pair"},
			{Input: map[string]interface{}{"letters": "stop", "k": 3}, Expected: btArrangements("stop", 3), Hidden: true, Note: "24 results: a letter must not repeat inside one result"},
			{Input: map[string]interface{}{"letters": "planet", "k": 4}, Expected: btArrangements("planet", 4), Hidden: true, Note: "360 results: every used flag must be cleared after the call"},
			{Input: map[string]interface{}{"letters": "zyxwvut", "k": 5}, Expected: btArrangements("zyxwvut", 5), Hidden: true, Note: "2520 results at the largest size"},
		},
		TimeComplexity:  "O(k * n! / (n - k)!)",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def arrangements(letters, k):\n    # Return every ordered selection of k different letters of the string\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Once \"b\" is on the path for \"abc\", which letters may follow it, and how can the code know that \"b\" is taken?"},
			{Level: 2, Type: "insight", Content: "Any letter that is not already on the path may come next, whatever its position, so every call loops over all of `letters` and skips the used ones. A `used` list marks the letters on the path: set it when you choose a letter and clear it when you un-choose, so it is right for the next branch. See Chapter 12.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: a result list, a path list, and a used list of False, one per letter.\n2. Stop: when path has k letters, record the joined string and return.\n3. Choose: for each index i with used[i] false, set used[i] and append letters[i].\n4. Explore: call backtrack().\n5. Un-choose: pop the letter and clear used[i]."},
			{Level: 4, Type: "skeleton", Content: "def arrangements(letters, k):\n    result = []\n    path = []\n    used = [False] * len(letters)\n\n    def backtrack():\n        if len(path) == k:\n            result.append(\"\".join(path))\n            return\n        for i in range(len(letters)):\n            if ____:\n                continue\n            used[i] = True        # choose\n            path.append(letters[i])\n            ____                  # explore\n            path.pop()            # un-choose\n            ____\n\n    backtrack()\n    return result"},
		},
		Solution: Solution{
			Code: `def arrangements(letters, k):
    result = []
    path = []
    used = [False] * len(letters)

    def backtrack():
        if len(path) == k:
            result.append("".join(path))
            return
        for i in range(len(letters)):
            if used[i]:
                continue
            used[i] = True
            path.append(letters[i])
            backtrack()
            path.pop()
            used[i] = False

    backtrack()
    return result`,
			Explanation:     "This builds on Combinations. There a start index forbade going back; here order matters, so every unused letter is allowed and a used list says which ones are taken. The shared state is now two things, the path and the used flags, and each choice has to be undone in both. It is the engine behind Permutations in the ladder.",
			BruteForce:      "Generate all n^k strings of k letters (every letter at every position, repeats allowed) and filter out the ones with a repeated letter. For 7 letters and k = 5 that is 16,807 strings to keep 2,520. The waste is building a path like \"aa\" and carrying it on to full length before throwing it away.",
			KeyInsight:      "Keep a flag per letter so a letter on the path can never be chosen again; invalid strings are then never built at all. The flag has to be cleared when the letter leaves the path, or it would block that letter in every other branch.",
			WhyItWorks:      "Invariant: when backtrack() is entered, used[i] is true exactly for the letters on the path, and when it returns, path and used are as they were on entry (each append and set is matched by a pop and a clear after the recursive call). So each call sees the correct set of free letters. Every result is a sequence of k different letters; each such sequence is one root-to-leaf path in the tree, reached by choosing its letters in order, and no path repeats a letter. Each is recorded once.",
			CommonMistakes:  []string{"Not clearing `used[i]` after the recursive call. The first branch marks every letter as used and leaves them marked, so the later branches find nothing free and most results are missing.", "Skipping the `used[i]` check, so a letter can repeat (\"aa\", \"bb\"): that is a different problem, with n^k results.", "Using a start index as in Combinations. It builds only increasing selections, so \"ba\" is never produced.", "Popping the letter but not clearing the flag (or clearing the flag and not popping): both pieces of state need undoing, always together."},
			TimeComplexity:  "O(k * n! / (n - k)!): each result is joined in O(k)",
			SpaceComplexity: "O(n) for the path, the flags and the recursion depth, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Set up the state", Explanation: "Besides the path, keep one flag per letter saying whether that letter is currently on the path.", CodeSnippet: "result = []\npath = []\nused = [False] * len(letters)", LineStart: 2, LineEnd: 4},
				{Title: "Record a complete answer", Explanation: "When the path has k letters, join them and stop.", CodeSnippet: "if len(path) == k:\n    result.append(\"\".join(path))\n    return", LineStart: 7, LineEnd: 9},
				{Title: "Skip used letters", Explanation: "Loop over every letter, from the start each time, and pass over the ones already on the path.", CodeSnippet: "for i in range(len(letters)):\n    if used[i]:\n        continue", LineStart: 10, LineEnd: 12},
				{Title: "Choose", Explanation: "Mark the letter as taken and put it on the path.", CodeSnippet: "used[i] = True\npath.append(letters[i])", LineStart: 13, LineEnd: 14},
				{Title: "Explore", Explanation: "The recursive call fills the next slot from the letters that are still free.", CodeSnippet: "backtrack()", LineStart: 15, LineEnd: 15},
				{Title: "Un-choose", Explanation: "Remove the letter and free its flag, so both the path and used are as before this choice.", CodeSnippet: "path.pop()\nused[i] = False", LineStart: 16, LineEnd: 17},
			},
			Alternatives: []AltSolution{
				{Name: "itertools.permutations", Code: "from itertools import permutations\n\ndef arrangements(letters, k):\n    return [\"\".join(p) for p in permutations(letters, k)]", TimeComplexity: "O(k * n! / (n - k)!)", SpaceComplexity: "O(k) besides the output", Explanation: "The library does the same enumeration. Write the recursion yourself first: the used-flag pattern is what Permutations, N-Queens and Word Search need."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can't a start index replace the used list here?", Answer: "A start index only allows letters after the last one chosen, which builds each set in one order. Arrangements need every order: after \"c\" the letter \"a\", which is earlier in the string, must still be allowed. So we allow any letter and rule out only the ones on the path, which is what used records."},
			{Prompt: "What if letters could contain repeated characters but results had to be unique?", Answer: "Then equal letters would build identical strings along different paths. One fix is to sort letters and, in a loop, skip a letter equal to the previous one when the previous one is not currently used. That is the rule Permutations II and Subsets II use to skip duplicates at the same depth."},
		},
	},
	{
		ID:              "drill-subsets-with-sum",
		Number:          0,
		Kind:            KindDrill,
		Order:           5,
		Title:           "Subsets That Add Up to a Target",
		Slug:            "drill-subsets-with-sum",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Backtracking", "Array", "Pruning", "Foundations"},
		RelatedChapters: []int{12, 8},
		Description:     "You are given a list `nums` of **distinct positive integers sorted in increasing order**, and a positive integer `target`. Return every subset of `nums` whose numbers add up to exactly `target`. Each subset is written in the same increasing order as `nums`, and the subsets can be returned in **any order**. Return an empty list when no subset works.\n\nSo `nums = [1,2,3,4,5]` and `target = 5` give `[[1,4],[2,3],[5]]`.\n\nThis is Subsets with one extra rule, and the rule is a chance to **prune**: all numbers are positive, so once the sum of the path is more than `target`, adding more numbers can only make it worse. Don't explore that branch. Because `nums` is sorted, if one number is too big for the remaining sum then all later numbers are too, so you can stop the loop (`break`) rather than only skipping that number. Some inputs have 40 numbers: the search must skip whole branches, because trying every one of the 2^40 subsets is far too slow.",
		Constraints:     []string{"1 <= len(nums) <= 40", "1 <= nums[i] <= 100, nums is sorted in increasing order, all distinct", "1 <= target <= 100"},
		Examples: []Example{
			{Input: "nums = [1,2,3,4,5], target = 5", Output: "[[1,4],[2,3],[5]]", Explanation: "1 + 4, 2 + 3 and 5 alone all give 5. After choosing 1, 2 and 3 the sum is 6, so that branch is dropped immediately."},
			{Input: "nums = [2,4], target = 3", Output: "[]", Explanation: "2 is too small alone, 4 is too big, and 2 + 4 = 6: no subset gives 3."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5}, "target": 5}, Expected: [][]int{{1, 4}, {2, 3}, {5}}, Note: "example 1: three subsets of different sizes"},
			{Input: map[string]interface{}{"nums": []int{2, 4}, "target": 3}, Expected: [][]int{}, Note: "example 2: no subset works, so the answer is an empty list"},
			{Input: map[string]interface{}{"nums": []int{7}, "target": 7}, Expected: [][]int{{7}}, Hidden: true, Note: "a single number equal to the target"},
			{Input: map[string]interface{}{"nums": []int{3, 5, 8}, "target": 1}, Expected: [][]int{}, Hidden: true, Note: "every number is already too big"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3, 4, 5, 6, 7, 8}, "target": 12}, Expected: btSubsetSums([]int{1, 2, 3, 4, 5, 6, 7, 8}, 12), Hidden: true, Note: "many subsets of 2 to 4 numbers: no number may repeat inside one subset"},
			{Input: map[string]interface{}{"nums": []int{2, 3, 5, 6, 9, 10}, "target": 15}, Expected: btSubsetSums([]int{2, 3, 5, 6, 9, 10}, 15), Hidden: true, Note: "numbers with gaps; the answer includes the three-number subsets"},
			{Input: map[string]interface{}{"nums": []int{1, 2, 3}, "target": 6}, Expected: [][]int{{1, 2, 3}}, Hidden: true, Note: "the whole list adds up to the target: the last number must be allowed in"},
			{Input: map[string]interface{}{"nums": btRange(1, 40), "target": 12}, Expected: btSubsetSums(btRange(1, 40), 12), Hidden: true, Timeout: 5000, Note: "40 numbers (2^40 subsets), target 12: a search that does not stop at a sum above the target runs out of time"},
		},
		TimeComplexity:  "O(2^n) in the worst case, far less with pruning",
		SpaceComplexity: "O(n) besides the output",
		StarterCode:     "def subsetsWithSum(nums, target):\n    # Return every subset of nums (sorted, distinct, positive) that adds up to target\n    pass",
		Harness:         Harness{Compare: "unordered-nested"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "With the path summing to 8 and a target of 5, is there any point in exploring what could come next? What about when only 2 is still needed and the next number is 6?"},
			{Level: 2, Type: "insight", Content: "Track the amount still needed (`remaining`). A path with remaining == 0 is an answer. Since every number is positive and nums is sorted, as soon as nums[i] > remaining, so is every later number: `break` out of the loop. That one line cuts all the branches that cannot work, and it is the difference between milliseconds and years on 40 numbers. Everything else is the Combinations skeleton. See Chapter 12.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: result, path and backtrack(start, remaining), with remaining starting at target.\n2. Stop: when remaining == 0, record a copy of path and return.\n3. Prune: in the loop, if nums[i] > remaining, break.\n4. Choose and explore: append nums[i], call backtrack(i + 1, remaining - nums[i]).\n5. Un-choose: pop nums[i]."},
			{Level: 4, Type: "skeleton", Content: "def subsetsWithSum(nums, target):\n    result = []\n    path = []\n\n    def backtrack(start, remaining):\n        if remaining == ____:\n            result.append(path[:])\n            return\n        for i in range(start, len(nums)):\n            if ____ > remaining:\n                break             # prune: later numbers are even bigger\n            path.append(nums[i])  # choose\n            backtrack(____, ____) # explore\n            ____                  # un-choose\n\n    backtrack(0, target)\n    return result"},
		},
		Solution: Solution{
			Code: `def subsetsWithSum(nums, target):
    result = []
    path = []

    def backtrack(start, remaining):
        if remaining == 0:
            result.append(path[:])
            return
        for i in range(start, len(nums)):
            if nums[i] > remaining:
                break
            path.append(nums[i])
            backtrack(i + 1, remaining - nums[i])
            path.pop()

    backtrack(0, target)
    return result`,
			Explanation:     "This builds on Combinations: the same start index, but the stop rule is a sum, and a new step prunes branches that can no longer succeed. Carrying the amount still needed makes the check one comparison, and sorting turns \"skip this number\" into \"stop the loop\". It is the pruning idea behind Combination Sum and N-Queens in the ladder.",
			BruteForce:      "Generate all 2^n subsets (Subsets) and keep those whose sum equals the target. With 40 numbers that is over a trillion subsets. The waste is extending a path that is already over the target: every subset built from it is also over the target, and they are all generated and summed anyway.",
			KeyInsight:      "Positive numbers only ever raise the sum, so a path that is too big stays too big. In a sorted list the first number that is too big for the remaining sum rules out all the numbers after it, so one break prunes the rest of the loop.",
			WhyItWorks:      "Invariant: remaining equals target minus the sum of path, and path is increasing with indexes below start. A subset adding to the target is reached by choosing its numbers in index order; the pruning never removes it, because at each of its choices nums[i] <= remaining (the rest of the subset adds up to remaining and every number is positive), so the break is not taken before i. A branch that is cut has nums[i] > remaining, and since all later numbers are at least nums[i] and positive, every subset through it exceeds the target. So no answer is lost and each is recorded once, when remaining reaches 0.",
			CommonMistakes:  []string{"Not pruning: building every subset and checking the sum at the end. It is correct on small inputs but needs 2^n steps, so the 40-number case times out. Stop as soon as nums[i] > remaining.", "Using `continue` instead of `break`. It stays correct, but it still looks at every later number, which are all too big in a sorted list. `break` is the version that ends the loop.", "Recursing with `backtrack(i, ...)` instead of `i + 1`: it lets one number be used several times (that is Combination Sum, a different problem).", "Forgetting `path.pop()`, so numbers from earlier branches stay on the path and the recorded subsets contain extras."},
			TimeComplexity:  "O(2^n) in the worst case; each path is at most a prefix whose sum is at most the target",
			SpaceComplexity: "O(n) for the path and recursion depth, plus the output",
			Walkthrough: []WalkthroughStep{
				{Title: "Record a complete answer", Explanation: "remaining is the amount still needed. Zero means the path adds up to the target exactly, so save a copy.", CodeSnippet: "if remaining == 0:\n    result.append(path[:])\n    return", LineStart: 6, LineEnd: 8},
				{Title: "Prune what cannot work", Explanation: "Try each number from start on, but nums is sorted and positive: if this number is more than the amount still needed, every later one is as well, so leave the loop.", CodeSnippet: "for i in range(start, len(nums)):\n    if nums[i] > remaining:\n        break", LineStart: 9, LineEnd: 11},
				{Title: "Choose", Explanation: "Add nums[i] to the path; looping from start keeps numbers in index order.", CodeSnippet: "path.append(nums[i])", LineStart: 12, LineEnd: 12},
				{Title: "Explore", Explanation: "Continue after i, with the amount still needed reduced by the number just chosen.", CodeSnippet: "backtrack(i + 1, remaining - nums[i])", LineStart: 13, LineEnd: 13},
				{Title: "Un-choose", Explanation: "Pop it. remaining is passed as an argument, so it needs no undo.", CodeSnippet: "path.pop()", LineStart: 14, LineEnd: 14},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is `break` correct here, and what input would make it wrong?", Answer: "nums is sorted ascending and positive, so if nums[i] > remaining then nums[i+1], nums[i+2], ... are at least as big and also exceed remaining. If nums were not sorted (or could hold zeros and negatives) a later number could still fit, so you would have to use continue for the number and could prune on the sum only if all numbers were non-negative."},
			{Prompt: "Why does `remaining` need no undo step after the call, while `path` does?", Answer: "remaining is an int passed by value into each call, so each call has its own copy and changing it for the child doesn't affect the parent. path is one list shared by every call, so a change made by one branch is seen by all unless it is reversed."},
		},
	},
	{
		ID:              "drill-grid-paths",
		Number:          0,
		Kind:            KindDrill,
		Order:           6,
		Title:           "Count Paths Through a Grid",
		Slug:            "drill-grid-paths",
		Difficulty:      "Medium",
		Category:        "backtracking",
		Tags:            []string{"Backtracking", "Matrix", "Recursion", "Foundations"},
		RelatedChapters: []int{12, 8},
		Description:     "You are given a rectangular `grid` of `0`s (open cells) and `1`s (walls). Count the different routes from the top-left cell `(0, 0)` to the bottom-right cell `(rows - 1, cols - 1)`.\n\nA route moves one step at a time **up, down, left or right** onto an open cell, and **never visits the same cell twice** in the same route. Two routes are different if their sequences of cells differ. If the start or the end is a wall, the answer is `0`. A grid with one open cell has one route (the start is also the end).\n\nTrack the cells on the current route in a `visited` grid: **mark** a cell when you step on it and **unmark** it when you leave it, so that other routes may use that cell again. Forgetting to unmark is the classic bug of Word Search and every maze problem, and this drill is where you practise restoring a shared grid.",
		Constraints:     []string{"1 <= rows, cols <= 5", "grid[r][c] is 0 or 1"},
		Examples: []Example{
			{Input: "grid = [[0,0],[0,0]]", Output: "2", Explanation: "Right then down, or down then right. Going right, down, left would visit the start's neighbour and then have nowhere new to go."},
			{Input: "grid = [[0,0,0],[1,1,0],[0,0,0]]", Output: "1", Explanation: "The middle row is walled except at the right, so the only route is across the top, down the right side and along the bottom."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"grid": [][]int{{0, 0}, {0, 0}}}, Expected: 2, Note: "example 1: an open 2 by 2 grid has two routes"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}}}, Expected: 1, Note: "example 2: walls leave one route"},
			{Input: map[string]interface{}{"grid": [][]int{{0}}}, Expected: 1, Hidden: true, Note: "one open cell: the start is the end"},
			{Input: map[string]interface{}{"grid": [][]int{{1}}}, Expected: 0, Hidden: true, Note: "one wall: no route at all"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 1}, {1, 0}}}, Expected: 0, Hidden: true, Note: "walls cut the start off from the end"},
			{Input: map[string]interface{}{"grid": [][]int{{1, 0}, {0, 0}}}, Expected: 0, Hidden: true, Note: "the start cell is a wall"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}}, Expected: 12, Hidden: true, Note: "open 3 by 3: 12 routes, including ones that turn back (up or left)"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, Expected: 2, Hidden: true, Note: "a wall in the middle: around either side"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}}, Expected: 184, Hidden: true, Note: "open 4 by 4: 184 routes, so a cell left marked after a branch loses most of them"},
			{Input: map[string]interface{}{"grid": [][]int{{0, 0, 0, 0, 0}, {0, 1, 0, 0, 0}, {0, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {0, 1, 0, 0, 0}}}, Expected: btGridPaths([][]int{{0, 0, 0, 0, 0}, {0, 1, 0, 0, 0}, {0, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {0, 1, 0, 0, 0}}), Hidden: true, Note: "5 by 5 with walls: many routes that wander"},
		},
		TimeComplexity:  "O(4^(rows * cols)) in the worst case",
		SpaceComplexity: "O(rows * cols) for the visited grid and recursion depth",
		StarterCode:     "def countPaths(grid):\n    # Count routes from top-left to bottom-right (up/down/left/right, no cell twice)\n    pass",
		Harness:         Harness{},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "If a route goes through cell X and finds the exit, and you then look for a different route that also passes through X, what must be true of X when you get there?"},
			{Level: 2, Type: "insight", Content: "Mark a cell visited when you step on it, so the same route can't go through it again, and unmark it when the call ends, so other routes can. The recursive call tries the four neighbours and returns 0 for anything off the grid, a wall or a visited cell, and 1 for the end. The un-choose step is `visited[row][col] = False`. See Chapter 12.", ChapterRef: 12},
			{Level: 3, Type: "plan", Content: "1. Set up the state: a visited grid of False.\n2. Reject: off the grid, a wall or already visited returns 0.\n3. Stop: the bottom-right cell returns 1.\n4. Choose: mark the cell as visited.\n5. Explore: add up backtrack() for the four neighbours.\n6. Un-choose: unmark the cell, then return the total."},
			{Level: 4, Type: "skeleton", Content: "def countPaths(grid):\n    rows, cols = len(grid), len(grid[0])\n    visited = [[False] * cols for _ in range(rows)]\n\n    def backtrack(row, col):\n        if row < 0 or row >= rows or col < 0 or col >= cols:\n            return 0\n        if grid[row][col] == 1 or visited[row][col]:\n            return 0\n        if row == rows - 1 and col == cols - 1:\n            return 1\n        ____                     # choose\n        total = 0\n        for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):\n            total += backtrack(____, ____)   # explore\n        ____                     # un-choose\n        return total\n\n    return backtrack(0, 0)"},
		},
		Solution: Solution{
			Code: `def countPaths(grid):
    rows, cols = len(grid), len(grid[0])
    visited = [[False] * cols for _ in range(rows)]

    def backtrack(row, col):
        if row < 0 or row >= rows or col < 0 or col >= cols:
            return 0
        if grid[row][col] == 1 or visited[row][col]:
            return 0
        if row == rows - 1 and col == cols - 1:
            return 1
        visited[row][col] = True
        total = 0
        for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
            total += backtrack(row + dr, col + dc)
        visited[row][col] = False
        return total

    return backtrack(0, 0)`,
			Explanation:     "This builds on Arrangements: there a used flag per letter marked what was on the path; here a visited flag per cell does the same on a grid, and the recursion counts instead of listing. Marking before exploring and unmarking afterwards makes the grid describe only the current route. This is exactly the mark and unmark of Word Search and N-Queens in the ladder.",
			BruteForce:      "For a grid this small, there is no cheaper way to count every route: the number of routes can itself be in the thousands. The wasteful approach is to build every sequence of moves (4 choices per step) and test it at the end; marking cells cuts off a route the moment it steps onto a cell it has already used.",
			KeyInsight:      "The visited grid must always describe the route that is being built, no more and no less. Mark on the way in, unmark on the way out, and a cell left over from a finished branch can never hide a route in another branch.",
			WhyItWorks:      "Invariant: while backtrack(row, col) runs on an open cell, visited is true exactly for the cells of the current route from (0, 0) to (row, col); on return it is back to the state at entry. Each call reaches the end only by a route of unvisited open cells, so a count of 1 is a valid route. Every valid route is found: it is a sequence of moves into unvisited open cells, and each prefix of it is explored since the loop tries all four directions. Routes differ in their cells, so each is counted once. Because every route ends at the end cell (it returns 1 before marking or moving on), no route is extended past it.",
			CommonMistakes:  []string{"Not unmarking the cell after exploring. The first route marks its cells and they stay marked, so other routes find them blocked and the count comes out far too low.", "Not marking the cell at all (or only checking walls): the route can step back and forth forever, and the recursion never ends.", "Moving only right and down. That counts a different problem (Unique Paths): the routes that turn back up or left are missing.", "Checking the end only after the wall and visited checks is fine, but forgetting to return 1 at the end cell lets the route walk on through it and be counted more than once.", "Reading `grid[row][col]` before the bounds check: a negative index silently wraps around to the other side of the grid in Python."},
			TimeComplexity:  "O(4^(rows * cols)) in the worst case; the visited marks cut it far lower",
			SpaceComplexity: "O(rows * cols) for the visited grid and recursion depth",
			Walkthrough: []WalkthroughStep{
				{Title: "Reject a bad cell", Explanation: "Off the grid, a wall, or a cell already on this route: no route goes this way, so count 0. The bounds are checked first so the grid is never indexed out of range.", CodeSnippet: "if row < 0 or row >= rows or col < 0 or col >= cols:\n    return 0\nif grid[row][col] == 1 or visited[row][col]:\n    return 0", LineStart: 6, LineEnd: 9},
				{Title: "Record a complete answer", Explanation: "The end cell finishes the route; count it as one and stop.", CodeSnippet: "if row == rows - 1 and col == cols - 1:\n    return 1", LineStart: 10, LineEnd: 11},
				{Title: "Choose", Explanation: "Mark the cell as part of the current route.", CodeSnippet: "visited[row][col] = True", LineStart: 12, LineEnd: 12},
				{Title: "Explore", Explanation: "Add up the routes that continue from each of the four neighbours.", CodeSnippet: "for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):\n    total += backtrack(row + dr, col + dc)", LineStart: 14, LineEnd: 15},
				{Title: "Un-choose", Explanation: "Unmark the cell before returning, so the next route that gets here finds it free.", CodeSnippet: "visited[row][col] = False\nreturn total", LineStart: 16, LineEnd: 17},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "In the 2 by 2 grid, why does the count stay at 2 even though the code can step into a cell from several sides?", Answer: "A route can't visit a cell twice, because visited blocks it. From (0,0) going right then down reaches the end; going right then trying left is blocked since (0,0) is marked. The only complete routes are right-down and down-right."},
			{Prompt: "What would you change to list the routes themselves as strings like \"RD\" instead of counting them?", Answer: "Keep a path list of moves, append a letter when you choose a direction and pop it after the call (a second piece of state to undo alongside visited), and record a copy of the joined path when you reach the end instead of returning 1."},
		},
	},
}

// btBinary lists every binary string of length n in increasing order.
func btBinary(n int) []string {
	out := []string{}
	for x := 0; x < 1<<n; x++ {
		var b strings.Builder
		for i := n - 1; i >= 0; i-- {
			b.WriteByte(byte('0' + (x>>i)&1))
		}
		out = append(out, b.String())
	}
	return out
}

// btCaseVariants lists every way to change the case of the letters in s.
func btCaseVariants(s string) []string {
	out := []string{""}
	for _, ch := range s {
		c := string(ch)
		lower, upper := strings.ToLower(c), strings.ToUpper(c)
		next := []string{}
		for _, prefix := range out {
			next = append(next, prefix+lower)
			if upper != lower {
				next = append(next, prefix+upper)
			}
		}
		out = next
	}
	return out
}

// btCombos lists every k-subset of 1..n in increasing order.
func btCombos(n, k int) [][]int {
	out := [][]int{}
	var path []int
	var walk func(start int)
	walk = func(start int) {
		if len(path) == k {
			out = append(out, append([]int{}, path...))
			return
		}
		for x := start; x <= n; x++ {
			path = append(path, x)
			walk(x + 1)
			path = path[:len(path)-1]
		}
	}
	walk(1)
	return out
}

// btArrangements lists every ordered selection of k distinct letters.
func btArrangements(letters string, k int) []string {
	out := []string{}
	used := make([]bool, len(letters))
	var path []byte
	var walk func()
	walk = func() {
		if len(path) == k {
			out = append(out, string(path))
			return
		}
		for i := 0; i < len(letters); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, letters[i])
			walk()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	walk()
	sort.Strings(out)
	return out
}

// btRange is the numbers lo..hi inclusive.
func btRange(lo, hi int) []int {
	out := []int{}
	for x := lo; x <= hi; x++ {
		out = append(out, x)
	}
	return out
}

// btSubsetSums lists the subsets of sorted positive nums that add up to target.
func btSubsetSums(nums []int, target int) [][]int {
	out := [][]int{}
	var path []int
	var walk func(start, remaining int)
	walk = func(start, remaining int) {
		if remaining == 0 {
			out = append(out, append([]int{}, path...))
			return
		}
		for i := start; i < len(nums) && nums[i] <= remaining; i++ {
			path = append(path, nums[i])
			walk(i+1, remaining-nums[i])
			path = path[:len(path)-1]
		}
	}
	walk(0, target)
	return out
}

// btGridPaths counts routes from the top-left to the bottom-right cell that
// never visit a cell twice.
func btGridPaths(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	var walk func(r, c int) int
	walk = func(r, c int) int {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 1 || visited[r][c] {
			return 0
		}
		if r == rows-1 && c == cols-1 {
			return 1
		}
		visited[r][c] = true
		total := walk(r+1, c) + walk(r-1, c) + walk(r, c+1) + walk(r, c-1)
		visited[r][c] = false
		return total
	}
	return walk(0, 0)
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, BacktrackingDrills...)
}
