package practice

// TriesProblems contains all tries category problems, in ladder order. The
// subgoal labels (walk or create the path, mark the word end, walk the path,
// check the end flag) are the ones chapter 10 teaches.
var TriesProblems = []*Problem{
	{
		ID:              "implement-trie",
		Number:          208,
		Order:           1,
		Title:           "Implement Trie (Prefix Tree)",
		Slug:            "implement-trie-prefix-tree",
		Difficulty:      "Medium",
		Category:        "tries",
		Tags:            []string{"Hash Table", "String", "Design", "Trie"},
		RelatedChapters: []int{10, 9, 2},
		Description: `A **trie** (say "try"), or prefix tree, stores a set of words as a tree of letters: each edge is one letter, and words that start the same way share the same path from the root. It answers "is this word stored?" and "does any stored word start with this prefix?" in time proportional to the length of the query.

Implement the ` + "`Trie`" + ` class:
- ` + "`Trie()`" + ` creates an empty trie.
- ` + "`insert(word)`" + ` adds ` + "`word`" + `.
- ` + "`search(word)`" + ` returns ` + "`true`" + ` if ` + "`word`" + ` was inserted before, and ` + "`false`" + ` otherwise.
- ` + "`startsWith(prefix)`" + ` returns ` + "`true`" + ` if some inserted word begins with ` + "`prefix`" + ` (a word counts as a prefix of itself), and ` + "`false`" + ` otherwise.`,
		Constraints: []string{
			"1 <= word.length, prefix.length <= 2000",
			"word and prefix consist only of lowercase English letters",
			"At most 3 * 10^4 calls in total to insert, search and startsWith",
		},
		Examples: []Example{
			{Input: "[\"Trie\",\"insert\",\"search\",\"search\",\"startsWith\",\"insert\",\"search\"]\n[[],[\"apple\"],[\"apple\"],[\"app\"],[\"app\"],[\"app\"],[\"app\"]]", Output: "[null,null,true,false,true,null,true]", Explanation: "After inserting \"apple\", the path a-p-p exists, so startsWith(\"app\") is true, but search(\"app\") is false: no word ends at that node. After insert(\"app\") the node is marked as a word end and search(\"app\") becomes true."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}, "values": [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}}, Expected: []interface{}{nil, nil, true, false, true, nil, true}, Hidden: false, Note: "the example: 'app' is a prefix before it is a word"},
			{Input: map[string]interface{}{"operations": []string{"Trie", "search", "startsWith"}, "values": [][]string{{}, {"a"}, {"a"}}}, Expected: []interface{}{nil, false, false}, Hidden: false, Note: "an empty trie contains nothing"},
			{Input: map[string]interface{}{"operations": []string{"Trie", "insert", "insert", "insert", "search", "search", "startsWith", "startsWith", "search"}, "values": [][]string{{}, {"cat"}, {"car"}, {"dog"}, {"car"}, {"ca"}, {"ca"}, {"do"}, {"cart"}}}, Expected: []interface{}{nil, nil, nil, nil, true, false, true, true, false}, Hidden: true, Note: "branches that share the prefix 'ca'"},
			{Input: map[string]interface{}{"operations": []string{"Trie", "insert", "insert", "search", "search", "startsWith"}, "values": [][]string{{}, {"app"}, {"apple"}, {"app"}, {"appl"}, {"apples"}}}, Expected: []interface{}{nil, nil, nil, true, false, false}, Hidden: true, Note: "a short word inserted before a longer one keeps its end flag"},
			{Input: map[string]interface{}{"operations": []string{"Trie", "insert", "insert", "search", "startsWith", "startsWith"}, "values": [][]string{{}, {"a"}, {"a"}, {"a"}, {"a"}, {"b"}}}, Expected: []interface{}{nil, nil, nil, true, true, false}, Hidden: true, Note: "inserting the same word twice; one-letter words"},
			{Input: map[string]interface{}{"operations": []string{"Trie", "insert", "startsWith", "startsWith", "search"}, "values": [][]string{{}, {"hello"}, {"hello"}, {"helloo"}, {"hell"}}}, Expected: []interface{}{nil, nil, true, false, false}, Hidden: true, Note: "a whole word is its own prefix; a longer prefix is not"},
		},
		TimeComplexity:  "O(L) per operation, L = length of the word or prefix",
		SpaceComplexity: "O(total letters inserted)",
		StarterCode:     "class Trie:\n    def __init__(self):\n        pass\n\n    def insert(self, word: str) -> None:\n        pass\n\n    def search(self, word: str) -> bool:\n        pass\n\n    def startsWith(self, prefix: str) -> bool:\n        pass",
		Harness:         Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "After inserting \"apple\", the letters a-p-p are stored on the way to \"apple\". How will search(\"app\") tell the difference between 'app was inserted' and 'app is only part of a longer word'?"},
			{Level: 2, Type: "insight", Content: "Make each node a dict of children (letter -> node) plus an is_end flag. insert walks from the root letter by letter, creating missing children, and sets is_end on the last node. search and startsWith both walk the path; if a letter is missing the answer is false. search additionally needs is_end on the final node; startsWith doesn't.", ChapterRef: 10},
			{Level: 3, Type: "plan", Content: "1. Define TrieNode with children = {} and is_end = False; the Trie holds a root node.\n2. insert: walk or create the path, one child per letter; then mark the word end.\n3. A helper walk(text): walk the path from the root; return None if a letter is missing, else the last node.\n4. search: walk the path, then check the end flag.\n5. startsWith: walk the path; reaching the end is enough."},
			{Level: 4, Type: "skeleton", Content: "class TrieNode:\n    def __init__(self):\n        self.children = {}\n        self.is_end = False\n\nclass Trie:\n    def __init__(self):\n        self.root = TrieNode()\n\n    def insert(self, word: str) -> None:\n        node = self.root\n        for ch in word:\n            if ch not in node.children:\n                ____\n            node = node.children[ch]\n        ____\n\n    def _walk(self, text):\n        node = self.root\n        for ch in text:\n            if ch not in node.children:\n                return ____\n            node = node.children[ch]\n        return node\n\n    def search(self, word: str) -> bool:\n        node = self._walk(word)\n        return ____\n\n    def startsWith(self, prefix: str) -> bool:\n        return ____"},
		},
		Solution: Solution{
			Code: `class TrieNode:
    def __init__(self):
        self.children = {}   # letter -> TrieNode
        self.is_end = False  # True if a whole word ends at this node


class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        node = self.root
        for ch in word:                       # walk or create the path
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True                    # mark the word end

    def _walk(self, text):
        node = self.root
        for ch in text:                       # walk the path
            if ch not in node.children:
                return None                   # fell off: nothing starts with text
            node = node.children[ch]
        return node

    def search(self, word: str) -> bool:
        node = self._walk(word)
        return node is not None and node.is_end   # check the end flag

    def startsWith(self, prefix: str) -> bool:
        return self._walk(prefix) is not None     # any node on the path will do`,
			Explanation: "The first rung builds the data structure itself, exactly as chapter 10 presents it. Each node maps letters to child nodes (the hash maps from chapter 2) and has a flag for 'a word ends here'. All three operations are a walk from the root, one letter at a time: insert creates missing nodes on the way, search and startsWith only follow existing edges, and they differ only in whether the end flag is checked.",
			BruteForce:  "Store the words in a Python set. search is then an O(L) hash lookup, but startsWith has to test every stored word: O(N * L) for N words. With 3 * 10^4 operations on 10^4 stored words that is up to 3 * 10^8 startswith checks. The waste: words that share a prefix are compared separately, letter by letter, over and over.",
			KeyInsight:  "Store shared prefixes once. Then 'does any word start with p?' is 'does the path spelled by p exist?', which takes len(p) steps no matter how many words are stored.",
			WhyItWorks:  "Invariant: a path from the root spelling s exists iff some inserted word starts with s, and the node at the end of that path has is_end == True iff s itself was inserted. insert keeps both true: it creates exactly the missing nodes along word's path (so every prefix of word now has a path) and sets is_end only on word's last node. So startsWith(p) is 'the path for p exists' and search(w) is 'the path for w exists and ends at a node with is_end'. Each operation touches one node per letter, so it costs O(L), independent of how many words are stored.",
			CommonMistakes: []string{
				"Returning True from search whenever the path exists. Then search(\"app\") is true after inserting only \"apple\". search must also check is_end.",
				"Using a class attribute like 'children = {}' in the class body instead of self.children in __init__. Every node then shares one dict and the trie becomes a mess. Create the dict per node in __init__.",
				"Setting is_end on every node during insert. Then every prefix of every word counts as a word.",
			},
			TimeComplexity:  "O(L) per insert, search or startsWith",
			SpaceComplexity: "O(total letters inserted) in the worst case (no shared prefixes)",
			Walkthrough: []WalkthroughStep{
				{Title: "Define the node", Explanation: "A node is a dict of children plus a flag that says whether a whole word ends here.", CodeSnippet: "self.children = {}\nself.is_end = False", LineStart: 3, LineEnd: 4},
				{Title: "Walk or create the path", Explanation: "For each letter, follow the child edge, creating it if it doesn't exist yet.", CodeSnippet: "for ch in word:\n    if ch not in node.children:\n        node.children[ch] = TrieNode()\n    node = node.children[ch]", LineStart: 13, LineEnd: 16},
				{Title: "Mark the word end", Explanation: "The last node of the path is where this word ends; prefixes of it are not marked.", CodeSnippet: "node.is_end = True", LineStart: 17, LineEnd: 17},
				{Title: "Walk the path", Explanation: "Shared by search and startsWith: follow existing edges only, and stop with None at the first missing letter.", CodeSnippet: "for ch in text:\n    if ch not in node.children:\n        return None", LineStart: 21, LineEnd: 23},
				{Title: "Check the end flag", Explanation: "search needs a stored word ending exactly here; startsWith is satisfied by the path alone.", CodeSnippet: "return node is not None and node.is_end", LineStart: 29, LineEnd: 32},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: a set of words", Code: `class Trie:
    def __init__(self):
        self.words = set()

    def insert(self, word):
        self.words.add(word)

    def search(self, word):
        return word in self.words

    def startsWith(self, prefix):
        return any(w.startswith(prefix) for w in self.words)   # scans every word`, TimeComplexity: "O(L) insert and search, O(N * L) startsWith", SpaceComplexity: "O(total letters inserted)", Explanation: "A hash set answers search just as fast, but has no structure for prefixes, so startsWith must check every word. This is why a trie beats a hash set as soon as prefix queries matter."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why can't search and startsWith be the same function?", Answer: "Both need the path to exist, but search asks whether a word ends exactly there. After inserting only \"apple\", the path for \"app\" exists (so startsWith is true) but the node for the second p has is_end False (so search is false). The end flag is the only difference."},
			{Prompt: "What would you change if the alphabet were only lowercase a-z and you needed it faster?", Answer: "Replace the dict with a fixed list of 26 child slots indexed by ord(ch) - ord('a'). Lookups become plain array indexing with no hashing. It uses more memory per node (26 slots even when most are empty), which is why the dict version is the default in Python."},
			{Prompt: "How would you support delete(word)?", Answer: "Walk to the word's last node and clear is_end. Then, optionally, remove nodes on the way back up that have no children and are not word ends, so the trie doesn't keep dead branches. A recursive delete that returns 'this child is now empty' makes that cleanup easy."},
		},
	},
	{
		ID:              "design-add-and-search-words",
		Number:          211,
		Order:           2,
		Title:           "Design Add and Search Words Data Structure",
		Slug:            "design-add-and-search-words-data-structure",
		Difficulty:      "Medium",
		Category:        "tries",
		Tags:            []string{"String", "Depth-First Search", "Design", "Trie"},
		RelatedChapters: []int{10, 8, 12},
		Description: `Design a word dictionary that supports adding words and searching with a wildcard.

Implement the ` + "`WordDictionary`" + ` class:
- ` + "`WordDictionary()`" + ` creates an empty dictionary.
- ` + "`addWord(word)`" + ` adds ` + "`word`" + `.
- ` + "`search(word)`" + ` returns ` + "`true`" + ` if some added word matches ` + "`word`" + `, and ` + "`false`" + ` otherwise. ` + "`word`" + ` may contain the character ` + "`'.'`" + `, which matches **exactly one** letter of any kind. A match must have the same length: ` + "`\"a.\"`" + ` matches ` + "`\"ab\"`" + ` but not ` + "`\"a\"`" + ` or ` + "`\"abc\"`" + `.`,
		Constraints: []string{
			"1 <= word.length <= 25",
			"word in addWord consists of lowercase English letters",
			"word in search consists of '.' or lowercase English letters",
			"There are at most 2 dots in each search query",
			"At most 10^4 calls in total to addWord and search",
		},
		Examples: []Example{
			{Input: "[\"WordDictionary\",\"addWord\",\"addWord\",\"addWord\",\"search\",\"search\",\"search\",\"search\"]\n[[],[\"bad\"],[\"dad\"],[\"mad\"],[\"pad\"],[\"bad\"],[\".ad\"],[\"b..\"]]", Output: "[null,null,null,null,false,true,true,true]", Explanation: "\"pad\" was never added. \".ad\" matches bad, dad and mad. \"b..\" matches bad."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"}, "values": [][]string{{}, {"bad"}, {"dad"}, {"mad"}, {"pad"}, {"bad"}, {".ad"}, {"b.."}}}, Expected: []interface{}{nil, nil, nil, nil, false, true, true, true}, Hidden: false, Note: "the example"},
			{Input: map[string]interface{}{"operations": []string{"WordDictionary", "search", "addWord", "search"}, "values": [][]string{{}, {"."}, {"a"}, {"."}}}, Expected: []interface{}{nil, false, nil, true}, Hidden: false, Note: "search before any word is added"},
			{Input: map[string]interface{}{"operations": []string{"WordDictionary", "addWord", "addWord", "search", "search", "search", "search"}, "values": [][]string{{}, {"a"}, {"ab"}, {"."}, {".."}, {"..."}, {"a."}}}, Expected: []interface{}{nil, nil, nil, true, true, false, true}, Hidden: true, Note: "dots must match exactly one letter: length matters"},
			{Input: map[string]interface{}{"operations": []string{"WordDictionary", "addWord", "addWord", "search", "search", "search"}, "values": [][]string{{}, {"at"}, {"and"}, {"an."}, {"a.d."}, {"..d"}}}, Expected: []interface{}{nil, nil, nil, true, false, true}, Hidden: true, Note: "a dot must try every child, not just the first"},
			{Input: map[string]interface{}{"operations": []string{"WordDictionary", "addWord", "search", "search"}, "values": [][]string{{}, {"abc"}, {"ab"}, {"abc."}}}, Expected: []interface{}{nil, nil, false, false}, Hidden: true, Note: "a prefix of a word is not a match"},
			{Input: map[string]interface{}{"operations": []string{"WordDictionary", "addWord", "addWord", "search", "search", "search"}, "values": [][]string{{}, {"xyz"}, {"xbz"}, {"x.z"}, {"..a"}, {"x.."}}}, Expected: []interface{}{nil, nil, nil, true, false, true}, Hidden: true, Note: "wildcard in the middle reaches both branches"},
		},
		TimeComplexity:  "O(L) addWord; O(L) search without dots, O(26^d * L) with d dots",
		SpaceComplexity: "O(total letters added)",
		StarterCode:     "class WordDictionary:\n    def __init__(self):\n        pass\n\n    def addWord(self, word: str) -> None:\n        pass\n\n    def search(self, word: str) -> bool:\n        pass",
		Harness:         Harness{Mode: "design"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "In Implement Trie, search followed exactly one child per letter. When the letter is '.', which children could the matching word continue through, and how would you try them all?"},
			{Level: 2, Type: "insight", Content: "Use the same trie. addWord is insert. search becomes a recursive walk match(node, i): a normal letter follows its one child; a '.' tries every child and succeeds if any branch matches the rest. This is depth-first search over the trie, and it is still pruned: a missing letter ends that branch immediately. At the end of the word, check is_end, so lengths must agree.", ChapterRef: 10},
			{Level: 3, Type: "plan", Content: "1. Walk or create the path and mark the word end (addWord, exactly as in Implement Trie).\n2. search defines match(node, i): 'can word[i:] be matched starting at node?'.\n3. Check the end flag: if i == len(word), return node.is_end.\n4. Branch on a wildcard: if word[i] == '.', return True if match(child, i + 1) is True for any child.\n5. Walk the path: otherwise follow the child for word[i], or return False if it doesn't exist."},
			{Level: 4, Type: "skeleton", Content: "def search(self, word: str) -> bool:\n    def match(node, i):\n        if i == len(word):\n            return ____\n        ch = word[i]\n        if ch == '.':\n            for child in ____:\n                if ____:\n                    return True\n            return False\n        if ch not in node.children:\n            return False\n        return ____\n\n    return match(self.root, 0)"},
		},
		Solution: Solution{
			Code: `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False


class WordDictionary:
    def __init__(self):
        self.root = TrieNode()

    def addWord(self, word: str) -> None:
        node = self.root
        for ch in word:                          # walk or create the path
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True                       # mark the word end

    def search(self, word: str) -> bool:
        def match(node, i):
            if i == len(word):
                return node.is_end               # check the end flag
            ch = word[i]
            if ch == '.':                        # branch on a wildcard
                for child in node.children.values():
                    if match(child, i + 1):
                        return True
                return False
            if ch not in node.children:          # walk the path
                return False
            return match(node.children[ch], i + 1)

        return match(self.root, 0)`,
			Explanation: "Builds on Implement Trie (problem 1): addWord is unchanged, and search keeps the same walk for normal letters. The new idea is the wildcard. A '.' can't pick one child, so search becomes a recursive depth-first search (chapter 8) that tries every child at a dot and stops as soon as one branch matches. The trie still prunes: any concrete letter that has no edge ends that branch.",
			BruteForce:  "Keep a list of all words and, for each query, compare it with every word of the same length, letter by letter, treating '.' as a match. With 10^4 words of length 25 that's up to 2.5 * 10^5 character comparisons per query. The waste: if the query starts with 'q' and only 3 words start with 'q', the other words are still compared.",
			KeyInsight:  "Concrete letters narrow the search to one branch, exactly as in a normal trie lookup; only dots fan out. With at most 2 dots, a query explores at most 26 * 26 branches, each walked in O(L), regardless of how many words were added.",
			WhyItWorks:  "match(node, i) returns True iff some added word has the letters on the path to node as its first i letters and matches word[i:] from there. Base case: when i == len(word), the word matched so far is complete, and it's an added word iff node.is_end (this enforces equal length). Step: a letter matches only the child with that letter, so following it (or failing if absent) is exact; a dot matches any letter, so the answer is True iff some child's subtree matches word[i+1:], which the loop checks. Returning on the first success is safe because one match is enough.",
			CommonMistakes: []string{
				"Treating a '.' at the end as 'any word with this prefix'. The match must end on a node with is_end; \"a.\" must not match \"abc\". Check is_end when i == len(word), as for any letter.",
				"Returning the result of the first child at a dot ('return match(child, i + 1)' inside the loop). That gives up after one child. Return True on success, and only return False after every child failed.",
				"Writing search iteratively with a single node pointer. A dot needs several branches at once; use recursion (or a list of current nodes) instead.",
			},
			TimeComplexity:  "O(L) addWord; search is O(L) without dots and O(26^d * L) with d dots (d <= 2 here)",
			SpaceComplexity: "O(total letters added) for the trie; O(L) recursion per search",
			Walkthrough: []WalkthroughStep{
				{Title: "Walk or create the path", Explanation: "addWord is the Implement Trie insert, unchanged.", CodeSnippet: "for ch in word:\n    if ch not in node.children:\n        node.children[ch] = TrieNode()", LineStart: 13, LineEnd: 17},
				{Title: "Check the end flag", Explanation: "Every character of the query is matched: it's a hit only if an added word ends here, which also enforces equal length.", CodeSnippet: "if i == len(word):\n    return node.is_end", LineStart: 21, LineEnd: 22},
				{Title: "Branch on a wildcard", Explanation: "A dot can be any letter: try each child, succeed on the first that matches the rest, fail only if all fail.", CodeSnippet: "if ch == '.':\n    for child in node.children.values():\n        if match(child, i + 1):\n            return True\n    return False", LineStart: 24, LineEnd: 28},
				{Title: "Walk the path", Explanation: "A normal letter has at most one matching child; a missing edge ends this branch.", CodeSnippet: "if ch not in node.children:\n    return False\nreturn match(node.children[ch], i + 1)", LineStart: 29, LineEnd: 31},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: compare with every word", Code: `class WordDictionary:
    def __init__(self):
        self.words = []

    def addWord(self, word):
        self.words.append(word)

    def search(self, word):
        for w in self.words:                     # compare against every stored word
            if len(w) == len(word) and all(p == '.' or p == c for p, c in zip(word, w)):
                return True
        return False`, TimeComplexity: "O(N * L) per search for N words", SpaceComplexity: "O(total letters added)", Explanation: "Stores words in a list and checks each one of equal length against the pattern. No pruning by prefix, so every query scans every word."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why is search fast for queries without dots, no matter how many words are stored?", Answer: "Without dots the recursion follows exactly one child per letter, like a normal trie lookup: at most L steps. The number of stored words only affects how wide the trie is, and a concrete letter never looks at sibling branches."},
			{Prompt: "What would change if '*' could match any sequence of letters (including none)?", Answer: "At a '*' you'd try two things: skip it (match(node, i + 1)) or let it consume one letter and stay on the '*' (match(child, i) for every child). That's the wildcard-matching recursion; adding memoization on (node, i) keeps it from blowing up."},
		},
	},
	{
		ID:              "word-search-ii",
		Number:          212,
		Order:           3,
		Title:           "Word Search II",
		Slug:            "word-search-ii",
		Difficulty:      "Hard",
		Category:        "tries",
		Tags:            []string{"Array", "String", "Backtracking", "Trie", "Matrix"},
		RelatedChapters: []int{10, 12, 8},
		Description: `You are given an ` + "`m x n`" + ` grid of lowercase letters ` + "`board`" + ` and a list of distinct ` + "`words`" + `. Return every word from the list that can be traced on the board.

To trace a word, start on a cell holding its first letter, then move step by step to a cell directly above, below, left or right (not diagonally) holding the next letter. A cell may be used **at most once within the same word** (different words may reuse cells). Return each found word once, in any order.`,
		Constraints: []string{
			"m == board.length, n == board[i].length",
			"1 <= m, n <= 12",
			"board[i][j] is a lowercase English letter",
			"1 <= words.length <= 3 * 10^4",
			"1 <= words[i].length <= 10",
			"words[i] consists of lowercase English letters, and all words are distinct",
		},
		Examples: []Example{
			{Input: `board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]`, Output: `["eat","oath"]`, Explanation: "\"oath\": o(0,0) a(0,1) t(1,1) h(2,1). \"eat\": e(1,3) a(1,2) t(1,1). There is no 'p' for \"pea\", and \"rain\" has no adjacent path."},
			{Input: `board = [["a","b"],["c","d"]], words = ["abcb"]`, Output: "[]", Explanation: "The only b is used once already."},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"board": [][]string{{"o", "a", "a", "n"}, {"e", "t", "a", "e"}, {"i", "h", "k", "r"}, {"i", "f", "l", "v"}}, "words": []string{"oath", "pea", "eat", "rain"}}, Expected: []string{"oath", "eat"}, Hidden: false, Note: "the example: two of four words"},
			{Input: map[string]interface{}{"board": [][]string{{"a", "b"}, {"c", "d"}}, "words": []string{"abcb"}}, Expected: []interface{}{}, Hidden: false, Note: "would need to reuse a cell"},
			{Input: map[string]interface{}{"board": [][]string{{"a"}}, "words": []string{"a", "b"}}, Expected: []string{"a"}, Hidden: true, Note: "single cell board"},
			{Input: map[string]interface{}{"board": [][]string{{"a", "a"}}, "words": []string{"aa", "aaa"}}, Expected: []string{"aa"}, Hidden: true, Note: "'aa' is found from both cells but reported once; 'aaa' needs a reused cell"},
			{Input: map[string]interface{}{"board": [][]string{{"o", "a", "a", "n"}, {"e", "t", "a", "e"}, {"i", "h", "k", "r"}, {"i", "f", "l", "v"}}, "words": []string{"oa", "oaa", "oat", "oath", "eta", "hike"}}, Expected: []string{"oa", "oat", "oath", "oaa", "eta"}, Hidden: true, Note: "words that share prefixes are all found in one walk"},
			{Input: map[string]interface{}{"board": [][]string{{"a", "b", "c"}, {"a", "e", "d"}, {"a", "f", "g"}}, "words": []string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"}}, Expected: []string{"abcdefg", "befa", "eaabcdgfa", "gfedcbaaa"}, Hidden: true, Note: "long snaking paths: cells must be restored after each branch"},
			{Input: map[string]interface{}{"board": [][]string{{"x", "y"}, {"z", "w"}}, "words": []string{"q", "xyzw", "xy"}}, Expected: []string{"xy"}, Hidden: true, Note: "letters not on the board are pruned immediately"},
		},
		TimeComplexity:  "O(m * n * 3^L), L = longest word length",
		SpaceComplexity: "O(W) for W total letters in words",
		StarterCode:     "def findWords(board, words):\n    # Write your solution here\n    pass",
		Harness:         Harness{Compare: "unordered"},
		Hints: []Hint{
			{Level: 1, Type: "nudge", Content: "Running Word Search (backtracking, chapter 12) once per word repeats work: \"oath\", \"oat\" and \"oak\" all walk o-a from the same cells. How could a single walk from each cell check all the words at once?"},
			{Level: 2, Type: "insight", Content: "Put all words in a trie, then backtrack over the board while walking the trie in parallel. From a cell, you may step to a neighbour only if the trie node has a child for that neighbour's letter: the trie prunes every path that isn't a prefix of some word. Store the whole word at its end node so you can record it when you arrive, and clear it so it's recorded once.", ChapterRef: 10},
			{Level: 3, Type: "plan", Content: "1. Build the trie of words; the node where a word ends stores the word.\n2. For every cell, start a DFS with the root as the parent node.\n3. Follow the trie edge: look up the child for this cell's letter; if there's none, prune (return).\n4. Record a found word: if the child stores a word, append it and clear it.\n5. Choose: mark the cell '#'. Explore: recurse into the four neighbours with the child node. Un-choose: restore the letter.\n6. Optional pruning: if the child now has no children and no word, delete it from its parent."},
			{Level: 4, Type: "skeleton", Content: "def dfs(r, c, parent):\n    ch = board[r][c]\n    node = parent.children.get(ch)\n    if ____:\n        return\n    if node.word is not None:\n        found.append(node.word)\n        ____\n    board[r][c] = '#'\n    for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):\n        if 0 <= nr < rows and 0 <= nc < cols and board[nr][nc] != '#':\n            dfs(nr, nc, ____)\n    board[r][c] = ____"},
		},
		Solution: Solution{
			Code: `class TrieNode:
    def __init__(self):
        self.children = {}
        self.word = None     # the whole word that ends here, if any


def findWords(board, words):
    root = TrieNode()
    for word in words:                         # build the trie of words
        node = root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.word = word

    rows, cols = len(board), len(board[0])
    found = []

    def dfs(r, c, parent):
        ch = board[r][c]
        node = parent.children.get(ch)
        if node is None:                       # prune: no word continues this way
            return
        if node.word is not None:              # record a found word, once
            found.append(node.word)
            node.word = None
        board[r][c] = '#'                      # choose: mark the cell as used
        for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
            if 0 <= nr < rows and 0 <= nc < cols and board[nr][nc] != '#':
                dfs(nr, nc, node)              # explore with the child node
        board[r][c] = ch                       # un-choose: restore the cell
        if not node.children and node.word is None:
            del parent.children[ch]            # prune: this branch is used up

    for r in range(rows):
        for c in range(cols):
            dfs(r, c, root)
    return found`,
			Explanation: "Builds on Implement Trie (problem 1) for the word store and on Word Search (chapter 12, backtracking ladder problem 8) for the grid DFS with choose / explore / un-choose. The new idea is walking both at once: each DFS step on the board is also one step down the trie, so a single search from each cell checks every word simultaneously and stops the moment the path is no longer a prefix of any word.",
			BruteForce:  "Run Word Search separately for each word: for each of W words, try every start cell and backtrack. Each run is O(m * n * 3^L), so the total is O(W * m * n * 3^L); with 3 * 10^4 words on a 12 x 12 board that's 4.3 million start-cell searches before any branching. The waste: words sharing a prefix re-walk the same board paths, and each run explores only one word's letters even when the board path spells the start of many words.",
			KeyInsight:  "The trie merges all words' prefixes into one tree, so one DFS that follows the trie explores each board path once for all words. A path is abandoned as soon as the trie has no edge for its next letter, which prunes far more than any single word would.",
			WhyItWorks:  "Invariant: when dfs(r, c, parent) is called, the cells marked '#' form a path whose letters spell the string on the trie path to parent, and (r, c) is an unmarked neighbour of its end. Following parent.children[board[r][c]] keeps this pairing for one more letter; if that child doesn't exist, no word has this path as a prefix, so nothing below can be found and returning is safe. Any word that can be traced is a path of distinct cells whose prefixes are all trie paths, so the DFS from its first cell reaches its end node and records it. Clearing node.word after recording ensures each word appears once even if it can be traced several ways. Deleting exhausted leaves only removes branches with no words left to find, so later searches skip them.",
			CommonMistakes: []string{
				"Not clearing node.word after recording it. A word reachable by two paths (like \"aa\" on [[\"a\",\"a\"]]) is then reported twice.",
				"Forgetting to restore the cell after exploring. Cells stay marked '#' and later searches from other cells can't use them, so words go missing.",
				"Checking the trie with the parent's letter instead of the current cell's letter, or recursing with parent instead of node. The trie must advance exactly one level per board step.",
			},
			TimeComplexity:  "O(m * n * 3^L): from each cell, at most 3 new directions per step and at most L steps, where L <= 10; the trie usually prunes far earlier",
			SpaceComplexity: "O(W) for the trie, W = total letters in words; O(L) recursion",
			Walkthrough: []WalkthroughStep{
				{Title: "Build the trie of words", Explanation: "Insert every word; the end node stores the whole word so it can be recorded without rebuilding it from the path.", CodeSnippet: "for word in words:\n    node = root", LineStart: 9, LineEnd: 15},
				{Title: "Follow the trie edge", Explanation: "Advance the trie by this cell's letter. No child means no word continues this way: prune.", CodeSnippet: "node = parent.children.get(ch)\nif node is None:\n    return", LineStart: 22, LineEnd: 24},
				{Title: "Record a found word", Explanation: "Arriving at a node that stores a word means the path spells it. Clear it so it's recorded once.", CodeSnippet: "if node.word is not None:\n    found.append(node.word)\n    node.word = None", LineStart: 25, LineEnd: 27},
				{Title: "Choose, explore, un-choose", Explanation: "Mark the cell, visit unmarked neighbours with the child node, then restore the letter.", CodeSnippet: "board[r][c] = '#'\nfor nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):\nboard[r][c] = ch", LineStart: 28, LineEnd: 32},
				{Title: "Prune used-up branches", Explanation: "A node with no children and no word left can never yield anything again; delete it so later starts skip it immediately.", CodeSnippet: "if not node.children and node.word is None:\n    del parent.children[ch]", LineStart: 33, LineEnd: 34},
			},
			Alternatives: []AltSolution{
				{Name: "Brute force: Word Search once per word", Code: `def findWords(board, words):
    rows, cols = len(board), len(board[0])

    def exists(word):                          # Word Search, once per word
        def dfs(r, c, i, used):
            if board[r][c] != word[i] or (r, c) in used:
                return False
            if i == len(word) - 1:
                return True
            used.add((r, c))
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < rows and 0 <= nc < cols and dfs(nr, nc, i + 1, used):
                    used.discard((r, c))
                    return True
            used.discard((r, c))
            return False

        return any(dfs(r, c, 0, set()) for r in range(rows) for c in range(cols))

    return [w for w in words if exists(w)]`, TimeComplexity: "O(W * m * n * 3^L) for W words", SpaceComplexity: "O(L)", Explanation: "Runs the chapter 12 Word Search for each word independently. Correct and simple, but repeats the same board walks for every word that shares a prefix."},
			},
		},
		SelfCheck: []SelfCheck{
			{Prompt: "Why does a trie help here, when a hash set of words would also tell you whether a path spells a word?", Answer: "A set can only answer 'is this exact string a word?'. To prune, the DFS needs 'is this string the start of any word?', which a set can't answer without checking every word. The trie answers it in O(1) per step (does this child exist?), so dead paths stop immediately."},
			{Prompt: "What does deleting exhausted leaves buy, and is it needed for correctness?", Answer: "It isn't needed for correctness: without it, the DFS would still find every word once (node.word is cleared). It speeds things up: after a word is found, its private branch of the trie is removed, so later searches stop earlier instead of re-walking paths that can't produce anything new."},
		},
	},
}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, TriesProblems...)
}
