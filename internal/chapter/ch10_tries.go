package chapter

import "time"

// chapter10 teaches tries (prefix trees). Storyboards: trie-ops.
func chapter10() Chapter {
	return Chapter{
		Number:        10,
		Slug:          "tries",
		Title:         "Tries",
		Category:      "tries",
		Prerequisites: []int{9, 2},
		Summary:       "A trie stores words as paths of letters from a root, so words with a common start share nodes. Any prefix question (does a word start with this? which words do?) then costs time proportional to the prefix, not to the number of words.",
		Motivation: `An autocomplete box holds 100,000 words. The user types "pre" and you must show the stored words that start with it. With a hash set, the only way to answer is to test every word: 100,000 startswith checks per keystroke. The set knows nothing about how words begin; "prefix" and "preface" sit in unrelated buckets.

Now store the words as a tree of letters. From the root, an edge labelled p leads to every word starting with p; from there an edge r leads to every word starting with pr, and so on. Answering "pre" means following three edges, and the words you want are exactly the ones in the subtree you land on. The cost is 3 steps to find the subtree, whether you store a hundred words or a million. That tree is a trie.`,
		Objectives: []string{
			"Explain why a trie answers prefix queries faster than a hash set, using the cost of walking the prefix",
			"Implement insert, search and startsWith on a trie of dict-based nodes with an end-of-word flag",
			"Trace a trie walk on a small set of words and say which node each query ends on",
			"Combine a trie with depth-first search to handle wildcards and to search a grid for many words at once",
		},
		WhenToUse: []string{
			"The problem asks about prefixes: 'starts with', autocomplete, 'all words beginning with'",
			"Many words are searched against the same data (a grid, a text, a stream), so shared prefixes can be explored once",
			"Queries contain wildcards that match one letter, and you still want to skip whole groups of words",
			"You need words in sorted (lexicographic) order or the longest common prefix of a set of strings",
		},
		Concepts: []Concept{
			{
				Name:        "Trie nodes, paths and the end-of-word flag",
				Description: "A trie is a rooted tree where each edge carries one letter. Each node keeps a dict from letter to child node, and a flag saying whether a whole stored word ends at that node. A word is stored as the path from the root that spells it.",
				WhyItWorks:  "Every string s corresponds to at most one path from the root (follow s[0], then s[1], ...). Insert creates the missing nodes on that path, so after inserting a word, every prefix of it has a path. That gives the invariant the operations rely on: the path for s exists iff some stored word starts with s, and the node at its end has is_end set iff s itself was stored. Each operation follows one edge per letter, and each edge lookup is a dict lookup (O(1) on average, chapter 2), so a query of length L costs O(L), independent of how many words are stored.",
				Intuition:   "Think of a library's card catalogue drawers labelled A, then AB, then ABA... Opening drawer 'pre' gives you every card starting with 'pre' without looking at other drawers. The drawers are trie nodes, the labels are paths, and a card in a drawer (not just sub-drawers) is the end-of-word flag.",
				CoreIdeas: []string{
					"A node is children (letter -> node) plus is_end. The root represents the empty string.",
					"Words that share a prefix share the nodes of that prefix, so storage is the number of distinct prefixes, not total letters.",
					"search(word) needs the path AND is_end; startsWith(prefix) needs only the path.",
					"Every operation is a walk: one step per letter, O(L) total.",
				},
				CommonMistakes: []string{
					"Forgetting the end flag and treating 'path exists' as 'word exists'. After inserting only \"apple\", search(\"app\") must be false; check is_end.",
					"Declaring children = {} at class level instead of in __init__. All nodes then share one dict. Always create self.children inside __init__.",
					"Assuming a trie always uses less memory than a set. Each node carries a dict; for words with few shared prefixes it can use more. Choose a trie for prefix queries, not to save memory.",
				},
				Examples: []CodeExample{
					{ID: "trie-node", Title: "Trie with insert, search and startsWith", Language: "python", Snippet: "class TrieNode:\n    def __init__(self):\n        self.children = {}\n        self.is_end = False\n\nclass Trie:\n    def __init__(self):\n        self.root = TrieNode()\n\n    def insert(self, word):\n        node = self.root\n        for ch in word:                       # walk or create the path\n            node = node.children.setdefault(ch, TrieNode())\n        node.is_end = True                    # mark the word end\n\n    def _walk(self, text):\n        node = self.root\n        for ch in text:                       # walk the path\n            node = node.children.get(ch)\n            if node is None:\n                return None\n        return node\n\n    def search(self, word):\n        node = self._walk(word)\n        return node is not None and node.is_end   # check the end flag\n\n    def startsWith(self, prefix):\n        return self._walk(prefix) is not None", Notes: "setdefault creates the child only if it's missing. The three methods share one walk; only search looks at is_end."},
				},
			},
			{
				Name:        "Why a trie beats a hash set for prefixes",
				Description: "A hash set answers 'is this exact word stored?' in O(L). It cannot answer 'does any stored word start with p?' without checking every word, because hashing scatters words that start alike. A trie keeps words that start alike together, so prefix questions become one walk.",
				WhyItWorks:  "A hash function is designed so that similar keys land in unrelated buckets; that's what makes lookups fast, and it's also what destroys any notion of 'starts with'. In a set of N words, startsWith(p) has to compare p against each word: O(N * |p|). In a trie, all words starting with p live in the subtree below p's node, so finding that subtree costs |p| steps and everything outside it is never touched. The same property lets a search prune: in a grid or with wildcards, the moment a partial path has no trie edge, no word can start that way, and the search stops.",
				Intuition:   "A hash set is a coat check: hand over your ticket (the exact word) and you get your coat instantly, but you can't ask for 'all coats with tickets starting 12'. A trie is a filing cabinet sorted by prefix: you can open the drawer for '12' and everything inside matches. The ticket is the word; the drawer is the prefix's node.",
				CoreIdeas: []string{
					"Exact lookups: set and trie are both O(L). Prefix lookups: set O(N * L), trie O(L).",
					"Pruning: a missing trie edge proves that no stored word continues this way, which is what makes Word Search II fast.",
					"All words below a node share that node's prefix; a DFS of the subtree lists them (in sorted order if children are visited alphabetically).",
					"Wildcards only fan out where the query has a '.'; concrete letters still follow one edge.",
				},
				CommonMistakes: []string{
					"Reaching for a trie when only exact lookups are needed. A set is simpler, smaller and just as fast for that.",
					"Searching many words in a grid one at a time. Put them in one trie and walk it alongside the grid, so shared prefixes are explored once.",
				},
				Examples: []CodeExample{
					{ID: "trie-vs-set", Title: "Prefix check: set versus trie", Language: "python", Snippet: "words = {\"car\", \"cart\", \"dog\"}\n\n# Hash set: has to look at every word\ndef starts_with_set(prefix):\n    return any(w.startswith(prefix) for w in words)   # O(N * L)\n\n# Trie: one step per letter of the prefix\ndef starts_with_trie(root, prefix):\n    node = root\n    for ch in prefix:\n        if ch not in node.children:\n            return False                             # nothing starts this way\n        node = node.children[ch]\n    return True                                      # O(L)", Notes: "Both are correct; only the trie's cost is independent of the number of words."},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Autocomplete: list every stored word that starts with a prefix",
				Problem: "Given a list of words and a prefix, return all the words that start with the prefix, in alphabetical order. Example: words = [\"car\", \"card\", \"care\", \"cat\", \"dog\"], prefix = \"car\" -> [\"car\", \"card\", \"care\"].",
				Steps: []WorkedStep{
					{Subgoal: "Walk or create the path", Explanation: "Insert each word: from the root, follow the edge for each letter, creating a node when it's missing. Words that start alike reuse the same nodes, so 'car', 'card' and 'care' share c-a-r.", Code: "root = TrieNode()\nfor word in words:\n    node = root\n    for ch in word:\n        node = node.children.setdefault(ch, TrieNode())", State: "after 'car', 'card', 'care', 'cat', 'dog': root -> c -> a -> {r -> {d, e}, t}; root -> d -> o -> g"},
					{Subgoal: "Mark the word end", Explanation: "Set is_end on the node where each word finishes. 'car' is a word and also a prefix of 'card', so its node is marked and still has children; the mark is what tells them apart.", Code: "    node.is_end = True", State: "is_end: c-a-r yes, c-a-r-d yes, c-a-r-e yes, c-a-t yes, d-o-g yes; c-a no"},
					{Subgoal: "Walk the path", Explanation: "Follow the prefix from the root without creating anything. If a letter is missing, no word has this prefix and the answer is empty. Otherwise we stop on the node for the prefix, and every matching word is in its subtree.", Code: "node = root\nfor ch in prefix:\n    if ch not in node.children:\n        return []\n    node = node.children[ch]", State: "prefix 'car': root -c-> a -r-> node(car); prefix 'cow' stops at 'o': return []"},
					{Subgoal: "Collect the subtree", Explanation: "Depth-first search from the prefix node, carrying the letters so far. Record a word at every node with is_end; visit children in sorted order so the output is alphabetical. Only the subtree is visited, never 'cat' or 'dog'.", Code: "result = []\ndef collect(node, text):\n    if node.is_end:\n        result.append(text)\n    for ch in sorted(node.children):\n        collect(node.children[ch], text + ch)\ncollect(node, prefix)", State: "collect(car): record 'car'; child d -> record 'card'; child e -> record 'care'; result = ['car', 'card', 'care']"},
				},
				Code: "class TrieNode:\n    def __init__(self):\n        self.children = {}\n        self.is_end = False\n\ndef autocomplete(words, prefix):\n    root = TrieNode()\n    for word in words:                        # walk or create the path\n        node = root\n        for ch in word:\n            node = node.children.setdefault(ch, TrieNode())\n        node.is_end = True                    # mark the word end\n\n    node = root\n    for ch in prefix:                         # walk the path\n        if ch not in node.children:\n            return []\n        node = node.children[ch]\n\n    result = []\n    def collect(node, text):                  # collect the subtree\n        if node.is_end:\n            result.append(text)\n        for ch in sorted(node.children):\n            collect(node.children[ch], text + ch)\n    collect(node, prefix)\n    return result",
			},
		},
		Completions: []CompletionProblem{
			{
				Title:       "Count the words that start with a prefix",
				Prompt:      "Each node keeps a counter `passing`: how many inserted words pass through it. Fill in the blanks so that count_prefix(prefix) returns how many inserted words start with prefix, in O(len(prefix)).",
				Code:        "class TrieNode:\n    def __init__(self):\n        self.children = {}\n        self.passing = 0\n\nclass PrefixCounter:\n    def __init__(self):\n        self.root = TrieNode()\n\n    def insert(self, word):\n        node = self.root\n        for ch in word:                      # walk or create the path\n            if ch not in node.children:\n                node.children[ch] = ____\n            node = node.children[ch]\n            node.passing += ____\n\n    def count_prefix(self, prefix):\n        node = self.root\n        for ch in prefix:                    # walk the path\n            if ch not in node.children:\n                return ____\n            node = node.children[ch]\n        return ____",
				Blanks:      []string{"TrieNode()", "1", "0", "node.passing"},
				Explanation: "Insert walks or creates the path and adds 1 to passing on every node it enters, so a node's passing counts the words whose path goes through it: exactly the words with that node's prefix. Counting then walks the path; a missing letter means no word has the prefix (0), otherwise the answer is the counter on the last node. No subtree has to be visited.",
			},
		},
		Animations: []Storyboard{
			{
				ID:    "trie-ops",
				Title: "Trie Insert & Search",
				Goal:  "Step through character-by-character traversal.",
				MemoryTips: []string{
					"Walk or create the path: one node per letter, reuse what's there",
					"Mark the word end on the last node only",
					"Walk the path: a missing letter means 'no word starts like this'",
					"Check the end flag: search needs it, startsWith doesn't",
				},
				Steps: []StoryboardStep{
					{Cue: "insert", Narration: "To insert a word, start at the root and take one step per letter. If the edge for a letter is missing, create a new node; if it's already there (another word shares this prefix), just follow it. Inserting 'card' after 'car' only creates one new node.", VisualHint: "Grow children per letter", Duration: Duration(3 * time.Second), CodeRef: "insert"},
					{Cue: "mark-end", Narration: "On the node for the last letter, set is_end. This is how the trie knows 'car' is a stored word and not only the start of 'card'. Nodes in the middle of a word stay unmarked.", VisualHint: "End flag", Duration: Duration(2 * time.Second), CodeRef: "end"},
					{Cue: "search", Narration: "search and startsWith walk the same path without creating anything. If any letter is missing, the answer is false. If the walk finishes, startsWith is true; search is true only if the final node has is_end set.", VisualHint: "Highlight traversed edges", Duration: Duration(3 * time.Second), CodeRef: "search/startsWith"},
				},
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt:      "You insert only \"apple\" into an empty trie. What do search(\"app\") and startsWith(\"app\") return?",
				Choices:     []string{"search true, startsWith true", "search false, startsWith true", "search false, startsWith false", "search true, startsWith false"},
				Answer:      1,
				Explanation: "The path a-p-p exists because it is part of \"apple\", so startsWith is true. But is_end is only set on the final 'e' node, so no word ends at the second 'p': search is false. Confusing 'path exists' with 'word exists' is the most common trie bug.",
			},
			{
				Prompt:      "Which pattern fits best? You get 50,000 dictionary words and must answer 10,000 queries of the form 'how many words start with this prefix?'.",
				Choices:     []string{"A hash set of the words, checking each word with startswith per query", "A trie with a passing-count on each node", "Sort the words and use two pointers from both ends", "A sliding window over the concatenated words"},
				Answer:      1,
				Explanation: "Prefix queries are the trie's strength: walk the prefix and read the count, O(L) per query. The hash set needs O(N * L) per query (50,000 comparisons each). Sorting plus binary search for the range of words with that prefix also works in O(L log N); two pointers and sliding windows (chapters 3 and 4) are for contiguous ranges in one sequence, not for prefix membership.",
			},
			{
				Prompt:      "Trace: words \"to\", \"tea\", \"ten\", \"i\", \"in\" are inserted into an empty trie. How many nodes does the trie have, counting the root?",
				Choices:     []string{"7", "8", "11", "10"},
				Answer:      1,
				Explanation: "Root; t; t-o; t-e; t-e-a; t-e-n; i; i-n. That's 1 + 7 = 8 nodes. The shared prefixes 't', 'te' and 'i' are stored once, which is why it's fewer than the 11 letters in total.",
			},
			{
				Prompt:      "Prerequisite check (chapter 2): a trie node stores its children in a Python dict. What is the average cost of 'ch in node.children', and why?",
				Explanation: "O(1) on average. A dict hashes the key to find its bucket directly, without scanning other keys. That's what makes each trie step constant time, so a walk over a word of length L is O(L). (With a fixed 26-letter alphabet you can also use a list of 26 slots indexed by ord(ch) - ord('a').)",
			},
			{
				Prompt:      "Explain in your own words why Word Search II with a trie is faster than running Word Search once per word.",
				Explanation: "Running Word Search per word re-walks the same board paths for every word that shares a prefix, and each run can only prune against one word. With all words in a trie, a single DFS from each cell walks the board and the trie together: a path is extended only while it is a prefix of some word, all words sharing that prefix are found in the same walk, and the search stops as soon as no trie edge exists for the next letter.",
			},
			{
				Prompt:      "Prerequisite check (chapter 9): a trie is a tree. Which traversal lists all stored words in alphabetical order?",
				Choices:     []string{"Breadth-first (level order)", "Preorder DFS visiting children in sorted letter order, recording a word when is_end is set", "Postorder DFS in any child order", "Inorder traversal"},
				Answer:      1,
				Explanation: "A word is recorded when its node is visited, and a shorter word (\"car\") must come before its extensions (\"card\"), so record before visiting children: preorder. Visiting children in sorted order puts \"card\" before \"care\". Level order would list all short words first; inorder only applies to binary trees.",
			},
		},
	}
}
