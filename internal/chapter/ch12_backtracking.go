package chapter

import "time"

// chapter12 teaches backtracking. Storyboards: permutations, n-queens, decision-tree.
func chapter12() Chapter {
	return Chapter{
		Number:        12,
		Slug:          "backtracking",
		Title:         "Backtracking",
		Category:      "backtracking",
		Prerequisites: []int{8},
		Summary:       "Backtracking builds answers one choice at a time, walking a tree of decisions depth first: choose, explore everything that follows, then un-choose and try the next option. Pruning cuts off branches that can't lead to a valid answer. It is the tool for 'list every...' and constraint puzzles.",
		Motivation: `Place 8 queens on a chessboard so that none attacks another. The obvious method is to try every way of putting 8 queens on 64 squares and test each board: C(64, 8) is about 4.4 billion boards. Even if you only try one queen per row, that's 8^8 = 16.7 million boards, and almost all of them are hopeless after the second queen: if the queens in rows 0 and 1 already attack each other, the 262,144 ways of filling rows 2 to 7 are all wasted.

Backtracking places queens row by row and checks each one the moment it's placed. When a queen is attacked, it doesn't fill the remaining rows; it takes that queen back and tries the next square. That simple habit (check early, undo, try the next option) finds all 92 solutions after trying about 2,000 placements instead of millions of boards.`,
		Objectives: []string{
			"Draw the decision tree for a small input and say what a node, a branch and a leaf represent",
			"Implement the template: record a complete answer, prune, choose, explore, un-choose",
			"Choose between a start index (order doesn't matter) and a used[] array (order matters) to avoid duplicates",
			"Skip repeated values at the same tree level when the input has duplicates",
			"Explain how the r - c and r + c keys detect diagonal attacks in N-Queens",
		},
		WhenToUse: []string{
			"The problem says 'return all', 'list every', 'generate all possible' (subsets, combinations, permutations, partitions)",
			"Constraint puzzles: place items so that no two conflict (N-Queens, Sudoku), or trace a path that uses each cell once (Word Search)",
			"The input is small (n up to about 10-20, or a 6 x 6 grid), a hint that exponential search is expected",
			"You need any one valid configuration and can check partial configurations early",
		},
		Concepts: []Concept{
			{
				Name:        "The decision tree: choose, explore, un-choose",
				Description: "Every backtracking problem is a depth-first walk of a tree of decisions. The root is 'nothing chosen yet'; each edge is one choice; each node is a partial answer (the path); leaves (or some nodes) are complete answers. One shared path list is extended when going down and shrunk when coming back up.",
				WhyItWorks:  "Recursion (chapter 8) gives us depth-first traversal for free: each call is a node, and the loop inside it tries each child. Using one path list for the whole search is safe because of a simple invariant: when a call returns, path is exactly as it was when the call started. Choose appends, explore recurses (and restores path itself before returning, by the same invariant), un-choose pops. So every sibling branch starts from the same state, and memory is O(depth) rather than a copy per node. Because each complete answer is one root-to-node path, it is produced exactly once.",
				Intuition:   "Exploring a maze with a ball of string. At each junction you choose a corridor (choose), walk it as far as you can (explore), and if it dead-ends you wind the string back to the junction (un-choose) and try the next corridor. The string is the path list; winding it back is pop().",
				CoreIdeas: []string{
					"Template: record a complete answer (base case) -> loop over choices -> prune invalid ones -> choose -> explore -> un-choose.",
					"Record a copy (path[:]), because path keeps changing after you record it.",
					"Un-choose must undo everything choose did (pop the path, clear used[], remove from sets), in reverse order.",
					"The output size sets the cost: 2^n subsets, n! permutations. Backtracking can't beat the output size, but pruning keeps it from doing much more.",
				},
				CommonMistakes: []string{
					"Appending path instead of path[:]. Every recorded answer is the same list object, which ends up empty. Always record a copy.",
					"Un-choosing only part of the state (popping path but leaving used[i] = True, or leaving a key in a set). Siblings then see stale state and answers go missing.",
					"Returning from inside the loop after the first recursive call. That explores only the first choice at each level. Return only from the base case (or when you need just one answer).",
				},
				Examples: []CodeExample{
					{ID: "backtracking-template", Title: "The backtracking template", Language: "python", Snippet: "def solve(input):\n    result = []\n    path = []\n\n    def backtrack(state):\n        if is_complete(state):           # record a complete answer\n            result.append(path[:])\n            return\n        for choice in choices(state):\n            if not is_valid(choice):     # prune\n                continue\n            path.append(choice)          # choose\n            backtrack(next_state(state, choice))  # explore\n            path.pop()                   # un-choose\n\n    backtrack(initial_state)\n    return result", Notes: "Every problem in this chapter fills in the four blanks: what's complete, what the choices are, what's invalid, and how the state advances."},
				},
			},
			{
				Name:        "Defining the choices: start index, used[], and duplicates",
				Description: "Most of the skill is deciding which choices are allowed at each node. If order inside an answer doesn't matter (subsets, combinations), only allow elements after the last one chosen: a start index. If order matters (permutations), allow any element not already used: a used[] array. If the input has duplicate values, sort it and skip a value equal to the previous one at the same level.",
				WhyItWorks:  "Duplicates come from reaching the same answer by different paths. With a start index, every subset can only be built in increasing index order, so it has exactly one path; recursing with i + 1 means each element is used at most once, with i it may repeat (Combination Sum). For permutations, order is the answer, so any unused element may come next, and used[] makes the check O(1). With duplicate values, two siblings with the same value would root identical subtrees; after sorting, 'i > start and nums[i] == nums[i - 1]' identifies exactly the second sibling, while deeper levels can still take the second copy (so [2, 2] is still built).",
				Intuition:   "Picking a team from a line of players. If the team is just a set of people, you walk down the line and only ever pick players further along (start index), so you never pick the same team twice in different orders. If you're choosing a batting order, anyone not yet in the order can bat next (used[]). If two players are identical twins, choosing 'the first twin' or 'the second twin' first gives the same team, so at each pick you only consider the first twin.",
				CoreIdeas: []string{
					"Order doesn't matter -> loop from start; recurse with i + 1 (use once) or i (reuse allowed).",
					"Order matters -> loop over all indices, skip used ones, mark and unmark used[i].",
					"Duplicate values -> sort, then skip i when i > start and nums[i] == nums[i - 1].",
					"Subsets record at every node; permutations and combinations of a fixed size or sum record only when complete.",
				},
				CommonMistakes: []string{
					"Recursing with start + 1 instead of i + 1 in the start-index loop. That lets earlier skipped elements back in and builds [2, 1] after [1, 2].",
					"Skipping duplicates with 'i > 0' instead of 'i > start'. That forbids a second copy even when it's chosen one level deeper, so [2, 2] is lost.",
					"Skipping duplicates without sorting first. Equal values aren't adjacent, so the comparison with nums[i - 1] misses them.",
				},
				Examples: []CodeExample{
					{ID: "start-vs-used", Title: "Start index versus used[]", Language: "python", Snippet: "# Subsets: order doesn't matter -> start index\ndef backtrack(start):\n    result.append(path[:])\n    for i in range(start, len(nums)):\n        path.append(nums[i])\n        backtrack(i + 1)\n        path.pop()\n\n# Permutations: order matters -> used[]\ndef backtrack():\n    if len(path) == len(nums):\n        result.append(path[:])\n        return\n    for i in range(len(nums)):\n        if used[i]:\n            continue\n        used[i] = True\n        path.append(nums[i])\n        backtrack()\n        path.pop()\n        used[i] = False", Notes: "Same template; the only difference is which choices the loop offers and when an answer counts as complete."},
				},
			},
			{
				Name:        "Pruning with constraints (N-Queens diagonals)",
				Description: "Pruning means rejecting a choice as soon as it breaks a rule, instead of building a full answer and testing it at the end. The check must be cheap, so keep the information it needs in sets or arrays that choose and un-choose update.",
				WhyItWorks:  "If a partial answer already breaks a rule, every extension of it breaks the rule too, so its whole subtree can be skipped without losing any answer. In N-Queens, placing one queen per row removes row conflicts by construction. A column is a set lookup. For diagonals: moving one step down-right adds 1 to both r and c, so r - c is the same on the whole diagonal; moving down-left adds 1 to r and subtracts 1 from c, so r + c is constant. Squares (0,1), (1,2), (2,3) share r - c = -1; squares (0,2), (1,1), (2,0) share r + c = 2. So 'is this diagonal attacked?' is one set lookup each, and the check is O(1).",
				Intuition:   "Seating guests who mustn't sit next to rivals. You don't seat everyone and then check the whole table; you check each guest as they sit down and send them to another chair at once. The 'who's at this table' list is the set; seating and unseating update it.",
				CoreIdeas: []string{
					"Check a choice before recursing; a failed check skips the entire subtree.",
					"With sorted inputs, a too-big value can 'break' the loop, since all later values are bigger (Combination Sum).",
					"N-Queens: three sets: cols, r - c (down-right diagonals), r + c (down-left diagonals). Add on choose, remove on un-choose.",
					"In grids, mark the current path's cells (e.g. overwrite with '#') and restore them on un-choose (Word Search).",
				},
				CommonMistakes: []string{
					"Checking only one diagonal direction. Both r - c and r + c are needed; with one, n = 6 returns far more than the 4 real solutions.",
					"Pruning with break when the remaining choices aren't sorted. Use continue unless you know every later choice also fails.",
				},
				Examples: []CodeExample{
					{ID: "n-queens-check", Title: "O(1) attack check for N-Queens", Language: "python", Snippet: "cols, diag, anti = set(), set(), set()\n\ndef safe(row, col):\n    return col not in cols and (row - col) not in diag and (row + col) not in anti\n\ndef place(row, col):        # choose\n    cols.add(col); diag.add(row - col); anti.add(row + col)\n\ndef remove(row, col):       # un-choose\n    cols.remove(col); diag.remove(row - col); anti.remove(row + col)", Notes: "Each set lookup is O(1), so testing a square doesn't depend on how many queens are placed."},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Combinations: every k-element subset of 1..n",
				Problem: "Given n and k, return all ways to choose k distinct numbers from 1..n, in any order (order inside a combination doesn't matter). Example: n = 4, k = 2 -> [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]].",
				Steps: []WorkedStep{
					{Subgoal: "Set up the state", Explanation: "Order inside a combination doesn't matter, so use a start index: after choosing a number, only larger numbers may follow. That gives each combination exactly one path. path is shared; result collects copies.", Code: "result = []\npath = []\ndef backtrack(start):", State: "n=4, k=2: backtrack(1), path=[]"},
					{Subgoal: "Record a complete answer", Explanation: "A combination is complete when it has k numbers. Record a copy and return: adding more would exceed k.", Code: "    if len(path) == k:\n        result.append(path[:])\n        return", State: "path=[1,2] has 2 numbers -> record [1,2]"},
					{Subgoal: "Choose", Explanation: "Try each number from start to n as the next element of the combination.", Code: "    for num in range(start, n + 1):\n        path.append(num)", State: "in backtrack(1): choose 1 -> path=[1]"},
					{Subgoal: "Explore", Explanation: "Recurse with num + 1 so only larger numbers can follow. This call records every combination that begins with the current path.", Code: "        backtrack(num + 1)", State: "backtrack(2) with path=[1] records [1,2], [1,3], [1,4]"},
					{Subgoal: "Un-choose", Explanation: "Pop the number so path is back to how it was, and the loop can try the next number at this level.", Code: "        path.pop()", State: "pop 1 -> path=[]; choose 2 -> records [2,3], [2,4]; choose 3 -> [3,4]; choose 4 -> nothing (no larger numbers)"},
				},
				Code: "def combine(n, k):\n    result = []\n    path = []\n\n    def backtrack(start):                 # set up the state\n        if len(path) == k:                # record a complete answer\n            result.append(path[:])\n            return\n        for num in range(start, n + 1):\n            path.append(num)              # choose\n            backtrack(num + 1)            # explore\n            path.pop()                    # un-choose\n\n    backtrack(1)\n    return result",
			},
			{
				Title:   "Restore IP Addresses: pruning by digit count and part validity",
				Problem: "Given a string s of digits, return every valid IPv4 address you can make by inserting exactly three dots into s, without reordering, adding or dropping digits. An address has four parts; each part is a number from 0 to 255 written without leading zeros (\"0\" is allowed, \"01\" is not). Example: s = \"101023\" -> [\"1.0.10.23\", \"1.0.102.3\", \"10.1.0.23\", \"10.10.2.3\", \"101.0.2.3\"].",
				Steps: []WorkedStep{
					{Subgoal: "Set up the state", Explanation: "Each level of the tree picks one part, so the path is the list of parts chosen so far and the state is start, the index of the first unused digit. left and need are computed on entry because both prunes and the base case are written in terms of them.", Code: "result = []\nparts = []\ndef backtrack(start):\n    left = len(s) - start\n    need = 4 - len(parts)", State: "s=\"101023\": backtrack(0), parts=[], left=6, need=4"},
					{Subgoal: "Record a complete answer", Explanation: "Four parts means the address is complete, but it only counts if every digit was used. Either way return: a fifth part is never allowed, so going deeper is pointless.", Code: "    if need == 0:\n        if left == 0:\n            result.append('.'.join(parts))\n        return", State: "parts=['1','0','10','23'], left=0 -> record \"1.0.10.23\"; parts=['1','0','1','0'], left=2 -> two digits unused, not recorded"},
					{Subgoal: "Prune: digits can't fit", Explanation: "Every remaining part needs 1 to 3 digits, so the leftover digits must be between need and 3 * need. If not, no way of splitting them works, so return before trying a single part and skip the whole subtree.", Code: "    if left < need or left > 3 * need:\n        return", State: "parts=['10','102'], start=5: left=1, need=2 -> 1 < 2, return at once (the root passes: 4 <= 6 <= 12)"},
					{Subgoal: "Choose", Explanation: "The choices at each node are the next part's length: 1, 2 or 3 digits starting at start. Slicing gives the candidate part; stop when the slice would run past the end of s.", Code: "    for length in range(1, 4):\n        if start + length > len(s):\n            break\n        part = s[start:start + length]", State: "in backtrack(0): length 1 -> part '1', later length 2 -> '10', length 3 -> '101'"},
					{Subgoal: "Prune: invalid part", Explanation: "A part with a leading zero or a value above 255 is illegal. Use break, not continue: making the part longer keeps the leading zero and only makes the number bigger, so every longer part fails too.", Code: "        if (part[0] == '0' and length > 1) or int(part) > 255:\n            break", State: "parts=['1'], start=1: part '0' is fine, part '01' has a leading zero -> break, so '010' is never built"},
					{Subgoal: "Explore", Explanation: "Append the part and recurse from the digit after it. The call returns once it has recorded every address that starts with the current parts, with parts restored.", Code: "        parts.append(part)\n        backtrack(start + length)", State: "parts=['1','0'] -> backtrack(2): left=4, need=2; this subtree records \"1.0.10.23\" and \"1.0.102.3\""},
					{Subgoal: "Un-choose", Explanation: "Pop the part so parts is exactly as it was at the start of this loop iteration; the next, longer part then starts from the same prefix.", Code: "        parts.pop()", State: "pop '0' -> parts=['1']; '01' pruned; pop '1' -> parts=[]; choose '10' -> records \"10.1.0.23\", \"10.10.2.3\"; choose '101' -> records \"101.0.2.3\"; result has 5 addresses"},
				},
				Code: "def restore_ip_addresses(s):\n    result = []\n    parts = []\n\n    def backtrack(start):                     # set up the state\n        left = len(s) - start                 # digits not yet used\n        need = 4 - len(parts)                 # parts still to fill\n        if need == 0:                         # record a complete answer\n            if left == 0:\n                result.append('.'.join(parts))\n            return\n        if left < need or left > 3 * need:    # prune: digits can't fit\n            return\n        for length in range(1, 4):            # choose a 1-3 digit part\n            if start + length > len(s):\n                break\n            part = s[start:start + length]\n            if (part[0] == '0' and length > 1) or int(part) > 255:\n                break                         # prune: longer parts fail too\n            parts.append(part)                # choose\n            backtrack(start + length)         # explore\n            parts.pop()                       # un-choose\n\n    backtrack(0)\n    return result",
			},
		},
		Completions: []CompletionProblem{
			{
				Title:       "Binary strings with no two adjacent 1s",
				Prompt:      "Return every string of n characters from '0' and '1' that never has two 1s next to each other. For n = 3: [\"000\", \"001\", \"010\", \"100\", \"101\"]. Fill in the blanks of the backtracking template; the prune rejects a '1' right after a '1'.",
				Code:        "def no_adjacent_ones(n):\n    result = []\n    path = []\n\n    def backtrack():\n        if ____:                                # record a complete answer\n            result.append(''.join(path))\n            return\n        for bit in '01':\n            if bit == '1' and ____:             # prune\n                continue\n            path.append(bit)                    # choose\n            backtrack()                         # explore\n            ____                                # un-choose\n\n    backtrack()\n    return result",
				Blanks:      []string{"len(path) == n", "path and path[-1] == '1'", "path.pop()"},
				Explanation: "Blank by blank. The path holds one character per level, so an answer is complete when len(path) == n: record it and return so this branch stops growing. The prune: a '1' is illegal only when the previous character is also '1', which is path[-1] == '1' — and the path and guard matters because at the first level path is empty, where a bare path[-1] would raise IndexError. continue skips only that bit; '0' at the same level is still tried. path.pop() is the un-choose: it removes this level's bit so the loop's next bit extends the same prefix the level started from. For n = 2: '0','0' records \"00\"; '0','1' records \"01\" (previous is '0', allowed); '1','0' records \"10\"; '1','1' is pruned — exactly the strings with no adjacent 1s.",
			},
		},
		Animations: []Storyboard{
			{
				ID:    "permutations",
				Title: "Generate Permutations",
				Goal:  "Generate all permutations using backtracking.",
				MemoryTips: []string{
					"Record a complete answer: only when the path holds every element",
					"Prune: skip elements already used",
					"Choose: mark used, append",
					"Explore: recurse for the next slot",
					"Un-choose: pop and unmark, in reverse order",
				},
				Steps: []StoryboardStep{
					{Cue: "choose", Narration: "At each level, any element not yet used can fill the next slot, because order matters in a permutation. Skip the elements already in the path.", VisualHint: "Select item", Duration: Duration(2 * time.Second), CodeRef: "permute: for num in remaining"},
					{Cue: "add", Narration: "Choose: mark the element as used and append it to the path. The path now describes one more slot of the arrangement.", VisualHint: "Build path", Duration: Duration(2 * time.Second), CodeRef: "permute: path.append(num)"},
					{Cue: "recurse", Narration: "Explore: recurse to fill the next slot with the remaining elements. When the path holds all n elements, record a copy; that leaf is one complete permutation.", VisualHint: "Continue", Duration: Duration(2 * time.Second), CodeRef: "permute: backtrack(path, remaining)"},
					{Cue: "pop", Narration: "Un-choose: pop the element and mark it unused. The path is exactly as it was before this choice, so the loop can try the next element in this slot. For 3 elements this produces 3 x 2 x 1 = 6 leaves.", VisualHint: "Pop from path", Duration: Duration(2 * time.Second), CodeRef: "permute: path.pop()"},
				},
			},
			{
				ID:    "n-queens",
				Title: "N-Queens Backtracking",
				Goal:  "Place queens row by row and backtrack on conflicts.",
				MemoryTips: []string{
					"One queen per row: the recursion depth is the row",
					"Prune: column, r - c and r + c must all be free",
					"Choose: add the column and both diagonal keys",
					"Un-choose: remove all three keys before trying the next column",
				},
				Steps: []StoryboardStep{
					{Cue: "place", Narration: "Work row by row, so two queens can never share a row. In the current row, try each column from left to right.", VisualHint: "Drop queen", Duration: Duration(2 * time.Second), CodeRef: "place"},
					{Cue: "attack-check", Narration: "Skip a square if its column is taken, or its down-right diagonal (same r - c) or down-left diagonal (same r + c) already has a queen. If it's safe, place the queen, record its three keys and move to the next row. If no column in a row is safe, remove the previous queen and try its next column.", VisualHint: "Mark diagonals/columns", Duration: Duration(3 * time.Second), CodeRef: "is_safe"},
					{Cue: "solution", Narration: "When every row has a queen, the board is a solution: record it as strings, then keep backtracking to find the others. For n = 4 there are exactly two.", VisualHint: "Board snapshot", Duration: Duration(2 * time.Second), CodeRef: "append solution"},
				},
			},
			{
				ID:    "decision-tree",
				Title: "Decision Tree with Pruning",
				Goal:  "Watch the recursion tree of Combinations (n = 4, k = 3) grow, backtrack and get pruned.",
				MemoryTips: []string{
					"Record a complete answer: when the path holds k numbers, save a copy and return",
					"Prune: if fewer than k - len(path) numbers are left, break; no later number can work either",
					"Choose: append num, which adds a child node below the current one",
					"Explore: recurse with num + 1, one level deeper in the tree",
					"Un-choose: pop num, grey the node out and try the next sibling",
				},
				Steps: []StoryboardStep{
					{Cue: "choose", Narration: "Each node is a call of backtrack and each edge is one chosen number. Choosing appends num to the path and draws a new child; the highlighted root-to-node line is exactly the path list.", VisualHint: "New child node, path highlighted", Duration: Duration(2 * time.Second), CodeRef: "combine: path.append(num)"},
					{Cue: "explore", Narration: "Explore: recurse with num + 1, so only larger numbers can follow. The search always goes as deep as it can before trying a sibling.", VisualHint: "Move down one level", Duration: Duration(2 * time.Second), CodeRef: "combine: backtrack(num + 1)"},
					{Cue: "record", Narration: "When the path holds k = 3 numbers, the node is a leaf: record a copy in the results list and return.", VisualHint: "Leaf turns green, answer added to results", Duration: Duration(2 * time.Second), CodeRef: "combine: result.append(path[:])"},
					{Cue: "un-choose", Narration: "Un-choose: pop the number. The node greys out and the current position moves back up to its parent, which tries its next number.", VisualHint: "Node greys out, path shrinks", Duration: Duration(2 * time.Second), CodeRef: "combine: path.pop()"},
					{Cue: "prune", Narration: "Prune: with path [1] we still need 2 numbers, but from 4 only one number (4) is left. That branch is crossed out and the loop breaks, since larger numbers leave even fewer. At the root, 3 is cut the same way: 4 leaves are found instead of exploring every branch.", VisualHint: "Branch drawn crossed out", Duration: Duration(3 * time.Second), CodeRef: "combine: if n - num + 1 < need: break"},
				},
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt:      "Trace: this code runs on nums = [1, 2, 3]. What is the 4th list appended to result?\n\ndef backtrack(start):\n    result.append(path[:])\n    for i in range(start, len(nums)):\n        path.append(nums[i])\n        backtrack(i + 1)\n        path.pop()",
				Choices:     []string{"[2]", "[1, 2, 3]", "[1, 3]", "[3]"},
				Answer:      1,
				Explanation: "The order is depth first: [] (root), [1], [1, 2], [1, 2, 3], then back up to [1, 3], [2], [2, 3], [3]. The 4th is [1, 2, 3]. The search goes as deep as possible before trying the next sibling.",
			},
			{
				Prompt:      "Which pattern fits best? 'Return every combination of k numbers from 1..9 that add up to n, each number used at most once.'",
				Choices:     []string{"Sliding window over 1..9", "Backtracking with a start index, recursing with i + 1 and pruning when the sum exceeds n", "Dynamic programming that counts the combinations", "Binary search for each number"},
				Answer:      1,
				Explanation: "'Return every combination' means the answers must be listed, so backtracking is the tool. A sliding window (chapter 4) only finds contiguous runs; DP (chapter 15) can count the ways but doesn't list them; binary search (chapter 6) finds one value, not sets of values. Use i + 1 because each number is used once.",
			},
			{
				Prompt:      "In N-Queens, a queen is at row 2, column 0. Which of these squares does it attack along a diagonal?",
				Choices:     []string{"(0, 0)", "(0, 2)", "(1, 2)", "(3, 2)"},
				Answer:      1,
				Explanation: "(2, 0) has r + c = 2 and r - c = 2. (0, 2) has r + c = 2, the same down-left diagonal. (0, 0) shares only the column; (1, 2) has r + c = 3 and r - c = -1; (3, 2) has r + c = 5 and r - c = 1.",
			},
			{
				Prompt:      "Why does the template append path[:] and not path when recording an answer?",
				Explanation: "path is one list shared by the whole search and keeps changing: later appends and pops modify it. Appending path itself would store a reference to that same list, so every recorded answer would end up as whatever path holds at the end (usually []). path[:] stores a snapshot copy that later changes can't touch.",
			},
			{
				Prompt:      "Prerequisite check (chapter 8): for Subsets on n = 10 elements, what is the maximum number of backtrack calls on the call stack at once, and why does that matter?",
				Choices:     []string{"2^10 = 1024", "11", "10! = 3,628,800", "1"},
				Answer:      1,
				Explanation: "The call stack holds one frame per level of the current path: the root call plus one per chosen element, so at most n + 1 = 11. The total number of calls is 2^n, but they don't all exist at once. That's why backtracking uses O(n) memory besides the output, and why Python's recursion limit (about 1000) is no problem here.",
			},
			{
				Prompt:      "Subsets uses a start index, Permutations uses used[]. In your own words, how do you decide which one a new problem needs?",
				Explanation: "Ask whether two answers with the same elements in a different order count as different. If not (subsets, combinations, combination sums), force one canonical order by only choosing elements after the last one: a start index. If yes (permutations, arrangements, orderings), any unused element can come next, so track which are used with used[] (or a set).",
			},
		},
	}
}
