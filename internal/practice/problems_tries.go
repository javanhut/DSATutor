package practice

// TriesProblems contains all tries category problems
var TriesProblems = []*Problem{
	{
		ID:              "implement-trie",
		Number:          68,
		Title:           "Implement Trie (Prefix Tree)",
		Difficulty:      "Medium",
		Category:        "tries",
		Tags:            []string{"Hash Table", "String", "Design", "Trie"},
		RelatedChapters: []int{5, 11, 12},
		Description: `A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings.

Implement the Trie class:
- Trie() Initializes the trie object.
- void insert(String word) Inserts the string word into the trie.
- boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
- boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.`,
		Constraints: []string{
			"1 <= word.length, prefix.length <= 2000",
			"word and prefix consist only of lowercase English letters",
			"At most 3 * 10^4 calls in total will be made to insert, search, and startsWith",
		},
		Examples: []Example{
			{Input: `["Trie","insert","search","search","startsWith","insert","search"]
[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]`, Output: "[null,null,true,false,true,null,true]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}, "values": [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}}, Expected: []interface{}{nil, nil, true, false, true, nil, true}},
		},
		TimeComplexity:  "O(m) per operation",
		SpaceComplexity: "O(n * m)",
		StarterCode: `class Trie:
    def __init__(self):
        pass

    def insert(self, word: str) -> None:
        pass

    def search(self, word: str) -> bool:
        pass

    def startsWith(self, prefix: str) -> bool:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Each node has children (dict/array) and an end-of-word flag. Traverse character by character.", ChapterRef: 11},
			{Level: 2, Type: "algorithm", Content: "Insert: create nodes as needed. Search: traverse and check end flag. StartsWith: traverse only, no end check."},
			{Level: 3, Type: "code", Content: "class TrieNode: children = {}, is_end = False. Insert: for char, create if not exists. Search: traverse, check is_end."},
		},
		Solution: Solution{
			Code: `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        node = self.root
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        node.is_end = True

    def search(self, word: str) -> bool:
        node = self._find_node(word)
        return node is not None and node.is_end

    def startsWith(self, prefix: str) -> bool:
        return self._find_node(prefix) is not None

    def _find_node(self, prefix: str):
        node = self.root
        for char in prefix:
            if char not in node.children:
                return None
            node = node.children[char]
        return node`,
			Explanation:     "Trie nodes have children dict and end flag. Insert creates path. Search/startsWith traverse, search also checks end flag.",
			TimeComplexity:  "O(m) where m is word/prefix length",
			SpaceComplexity: "O(n * m) for all inserted characters",
			Walkthrough: []WalkthroughStep{
				{Title: "TrieNode structure", Explanation: "Children dict and end-of-word flag", CodeSnippet: "self.children = {}\nself.is_end = False", LineStart: 3, LineEnd: 4},
				{Title: "Insert word", Explanation: "Create nodes along path, mark end", CodeSnippet: "if char not in node.children:\n    node.children[char] = TrieNode()\nnode.is_end = True", LineStart: 13, LineEnd: 16},
				{Title: "Search vs startsWith", Explanation: "Search checks is_end, startsWith doesn't", CodeSnippet: "return node is not None and node.is_end", LineStart: 20, LineEnd: 20},
			},
		},
	},
	{
		ID:              "design-add-and-search-words",
		Number:          69,
		Title:           "Design Add and Search Words Data Structure",
		Difficulty:      "Medium",
		Category:        "tries",
		Tags:            []string{"String", "DFS", "Design", "Trie"},
		RelatedChapters: []int{5, 11, 12},
		Description: `Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:
- WordDictionary() Initializes the object.
- void addWord(word) Adds word to the data structure, it can be matched later.
- bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.`,
		Constraints: []string{
			"1 <= word.length <= 25",
			"word in addWord consists of lowercase English letters",
			"word in search consist of '.' or lowercase English letters",
			"There will be at most 2 dots in word for search queries",
			"At most 10^4 calls will be made to addWord and search",
		},
		Examples: []Example{
			{Input: `["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]`, Output: "[null,null,null,null,false,true,true,true]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"operations": []string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"}, "values": [][]string{{}, {"bad"}, {"dad"}, {"mad"}, {"pad"}, {"bad"}, {".ad"}, {"b.."}}}, Expected: []interface{}{nil, nil, nil, nil, false, true, true, true}},
		},
		TimeComplexity:  "O(m) add, O(26^d * m) search",
		SpaceComplexity: "O(n * m)",
		StarterCode: `class WordDictionary:
    def __init__(self):
        pass

    def addWord(self, word: str) -> None:
        pass

    def search(self, word: str) -> bool:
        pass`,
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Use a Trie. For '.', try all children recursively instead of specific character.", ChapterRef: 11},
			{Level: 2, Type: "algorithm", Content: "addWord is standard Trie insert. search uses DFS: for '.', branch to all children; for letter, follow that child."},
			{Level: 3, Type: "code", Content: "def dfs(node, i): if i == len(word): return node.is_end. if word[i] == '.': try all children. else follow char."},
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
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        node.is_end = True

    def search(self, word: str) -> bool:
        def dfs(node, i):
            if i == len(word):
                return node.is_end

            char = word[i]
            if char == '.':
                # Try all children
                for child in node.children.values():
                    if dfs(child, i + 1):
                        return True
                return False
            else:
                if char not in node.children:
                    return False
                return dfs(node.children[char], i + 1)

        return dfs(self.root, 0)`,
			Explanation:     "Trie with wildcard search. For '.', try all children (backtracking). For letters, follow specific path.",
			TimeComplexity:  "O(m) for add, O(26^d * m) for search with d dots",
			SpaceComplexity: "O(n * m) for Trie storage",
			Walkthrough: []WalkthroughStep{
				{Title: "Standard Trie add", Explanation: "Same as regular Trie insert", CodeSnippet: "for char in word:\n    if char not in node.children:", LineStart: 11, LineEnd: 14},
				{Title: "Handle wildcard", Explanation: "Try all children for '.'", CodeSnippet: "if char == '.':\n    for child in node.children.values():", LineStart: 24, LineEnd: 28},
				{Title: "Regular search", Explanation: "Follow specific character path", CodeSnippet: "return dfs(node.children[char], i + 1)", LineStart: 33, LineEnd: 33},
			},
		},
	},
	{
		ID:              "word-search-ii",
		Number:          70,
		Title:           "Word Search II",
		Difficulty:      "Hard",
		Category:        "tries",
		Tags:            []string{"Array", "String", "Backtracking", "Trie", "Matrix"},
		RelatedChapters: []int{4, 5, 11, 12},
		Description: `Given an m x n board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.`,
		Constraints: []string{
			"m == board.length",
			"n == board[i].length",
			"1 <= m, n <= 12",
			"board[i][j] is a lowercase English letter",
			"1 <= words.length <= 3 * 10^4",
			"1 <= words[i].length <= 10",
			"words[i] consists of lowercase English letters",
			"All the strings of words are unique",
		},
		Examples: []Example{
			{Input: `board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]`, Output: `["eat","oath"]`},
			{Input: `board = [["a","b"],["c","d"]], words = ["abcb"]`, Output: "[]"},
		},
		TestCases: []TestCase{
			{Input: map[string]interface{}{"board": [][]string{{"o", "a", "a", "n"}, {"e", "t", "a", "e"}, {"i", "h", "k", "r"}, {"i", "f", "l", "v"}}, "words": []string{"oath", "pea", "eat", "rain"}}, Expected: []string{"eat", "oath"}},
		},
		TimeComplexity:  "O(m * n * 4^L)",
		SpaceComplexity: "O(N)",
		StarterCode:     "def findWords(board, words):\n    # Write your solution here\n    pass",
		Hints: []Hint{
			{Level: 1, Type: "approach", Content: "Build Trie from words. DFS from each cell, pruning paths not in Trie. More efficient than searching each word separately.", ChapterRef: 11},
			{Level: 2, Type: "algorithm", Content: "Insert all words into Trie. DFS from each cell following Trie. Mark visited cells. When reach end-of-word, add to result."},
			{Level: 3, Type: "code", Content: "Build Trie. For each cell, dfs with Trie node. If char not in children, return. If is_end, add word. Mark visited, explore 4 dirs, unmark."},
		},
		Solution: Solution{
			Code: `def findWords(board, words):
    class TrieNode:
        def __init__(self):
            self.children = {}
            self.word = None

    # Build Trie
    root = TrieNode()
    for word in words:
        node = root
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        node.word = word

    result = []
    m, n = len(board), len(board[0])

    def dfs(r, c, node):
        char = board[r][c]
        if char not in node.children:
            return

        next_node = node.children[char]
        if next_node.word:
            result.append(next_node.word)
            next_node.word = None  # Avoid duplicates

        # Mark visited
        board[r][c] = '#'

        # Explore 4 directions
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nr, nc = r + dr, c + dc
            if 0 <= nr < m and 0 <= nc < n and board[nr][nc] != '#':
                dfs(nr, nc, next_node)

        # Restore
        board[r][c] = char

    for r in range(m):
        for c in range(n):
            dfs(r, c, root)

    return result`,
			Explanation:     "Build Trie from words. DFS from each cell following Trie paths. Much faster than checking each word separately.",
			TimeComplexity:  "O(m * n * 4^L) where L is max word length",
			SpaceComplexity: "O(N) for Trie where N is total characters in words",
			Walkthrough: []WalkthroughStep{
				{Title: "Store word at end", Explanation: "Store full word instead of just flag for easy retrieval", CodeSnippet: "node.word = word", LineStart: 15, LineEnd: 15},
				{Title: "Prune with Trie", Explanation: "Stop DFS if path not in Trie", CodeSnippet: "if char not in node.children:\n    return", LineStart: 22, LineEnd: 23},
				{Title: "Mark and restore", Explanation: "Prevent revisiting same cell in path", CodeSnippet: "board[r][c] = '#'\n# explore...\nboard[r][c] = char", LineStart: 31, LineEnd: 40},
			},
		},
	},

}

func init() {
	EmbeddedProblems = append(EmbeddedProblems, TriesProblems...)
}
