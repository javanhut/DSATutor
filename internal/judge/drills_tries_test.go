package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestTriesDrillsRejectMistakes checks that each tries foundations drill
// fails the mistake its tests and notes promise to catch, so a drill can't be
// passed without doing the thing it drills.
func TestTriesDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-trie-build-nodes", "counts every letter instead of only new nodes", `class TrieNode:
    def __init__(self):
        self.children = {}

class Trie:
    def __init__(self):
        self.root = TrieNode()
        self.nodes = 0

    def insert(self, word):
        node = self.root
        for ch in word:
            node.children[ch] = TrieNode()
            self.nodes += 1
            node = node.children[ch]

    def countNodes(self):
        return self.nodes`},
		{"drill-trie-build-nodes", "counts the root as a node", `class TrieNode:
    def __init__(self):
        self.children = {}

class Trie:
    def __init__(self):
        self.root = TrieNode()
        self.nodes = 1

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
                self.nodes += 1
            node = node.children[ch]

    def countNodes(self):
        return self.nodes`},
		{"drill-trie-build-nodes", "shares one children dict between all nodes", `class TrieNode:
    children = {}

class Trie:
    def __init__(self):
        self.root = TrieNode()
        self.nodes = 0

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
                self.nodes += 1
            node = node.children[ch]

    def countNodes(self):
        return self.nodes`},
		{"drill-trie-end-flag", "contains accepts any existing path", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def contains(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return True

    def longestWordPrefix(self, text):
        node = self.root
        best = 0
        for i, ch in enumerate(text):
            if ch not in node.children:
                break
            node = node.children[ch]
            if node.is_end:
                best = i + 1
        return text[:best]`},
		{"drill-trie-end-flag", "longestWordPrefix returns the deepest path, not the deepest word", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def contains(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return node.is_end

    def longestWordPrefix(self, text):
        node = self.root
        best = 0
        for i, ch in enumerate(text):
            if ch not in node.children:
                break
            node = node.children[ch]
            best = i + 1
        return text[:best]`},
		{"drill-trie-end-flag", "longestWordPrefix returns the first word found, not the longest", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def contains(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return node.is_end

    def longestWordPrefix(self, text):
        node = self.root
        for i, ch in enumerate(text):
            if ch not in node.children:
                break
            node = node.children[ch]
            if node.is_end:
                return text[:i + 1]
        return ""`},
		{"drill-trie-counts", "a boolean end flag counts duplicate copies once", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.pass_count = 0
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
            node.pass_count += 1
        node.is_end = True

    def _walk(self, text):
        node = self.root
        for ch in text:
            if ch not in node.children:
                return None
            node = node.children[ch]
        return node

    def countWordsEqualTo(self, word):
        node = self._walk(word)
        return 1 if node and node.is_end else 0

    def countWordsStartingWith(self, prefix):
        node = self._walk(prefix)
        return node.pass_count if node else 0

    def erase(self, word):
        node = self.root
        for ch in word:
            node = node.children[ch]
            node.pass_count -= 1
        node.is_end = False`},
		{"drill-trie-counts", "prefix count looks only at words ending at the prefix node", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.end_count = 0

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.end_count += 1

    def _walk(self, text):
        node = self.root
        for ch in text:
            if ch not in node.children:
                return None
            node = node.children[ch]
        return node

    def countWordsEqualTo(self, word):
        node = self._walk(word)
        return node.end_count if node else 0

    def countWordsStartingWith(self, prefix):
        node = self._walk(prefix)
        return node.end_count if node else 0

    def erase(self, word):
        self._walk(word).end_count -= 1`},
		{"drill-trie-counts", "erase leaves the prefix counts on the path unchanged", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.pass_count = 0
        self.end_count = 0

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
            node.pass_count += 1
        node.end_count += 1

    def _walk(self, text):
        node = self.root
        for ch in text:
            if ch not in node.children:
                return None
            node = node.children[ch]
        return node

    def countWordsEqualTo(self, word):
        node = self._walk(word)
        return node.end_count if node else 0

    def countWordsStartingWith(self, prefix):
        node = self._walk(prefix)
        return node.pass_count if node else 0

    def erase(self, word):
        self._walk(word).end_count -= 1`},
		{"drill-trie-counts", "crashes on a path that does not exist", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.pass_count = 0
        self.end_count = 0

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
            node.pass_count += 1
        node.end_count += 1

    def _walk(self, text):
        node = self.root
        for ch in text:
            node = node.children[ch]
        return node

    def countWordsEqualTo(self, word):
        return self._walk(word).end_count

    def countWordsStartingWith(self, prefix):
        return self._walk(prefix).pass_count

    def erase(self, word):
        node = self.root
        for ch in word:
            node = node.children[ch]
            node.pass_count -= 1
        node.end_count -= 1`},
		{"drill-trie-autocomplete", "lists words in insertion order, not alphabetical", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def wordsWithPrefix(self, prefix):
        node = self.root
        for ch in prefix:
            if ch not in node.children:
                return []
            node = node.children[ch]
        words = []

        def collect(current, path):
            if current.is_end:
                words.append(path)
            for ch in current.children:
                collect(current.children[ch], path + ch)

        collect(node, prefix)
        return words`},
		{"drill-trie-autocomplete", "forgets the prefix itself when it is a word", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def wordsWithPrefix(self, prefix):
        node = self.root
        for ch in prefix:
            if ch not in node.children:
                return []
            node = node.children[ch]
        words = []

        def collect(current, path):
            for ch in sorted(current.children):
                child = current.children[ch]
                if child.is_end:
                    words.append(path + ch)
                collect(child, path + ch)

        collect(node, prefix)
        return words`},
		{"drill-trie-autocomplete", "explores children before the node's own word", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def wordsWithPrefix(self, prefix):
        node = self.root
        for ch in prefix:
            if ch not in node.children:
                return []
            node = node.children[ch]
        words = []

        def collect(current, path):
            for ch in sorted(current.children):
                collect(current.children[ch], path + ch)
            if current.is_end:
                words.append(path)

        collect(node, prefix)
        return words`},
		{"drill-trie-autocomplete", "crashes when the prefix is not in the trie", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def wordsWithPrefix(self, prefix):
        node = self.root
        for ch in prefix:
            node = node.children[ch]
        words = []

        def collect(current, path):
            if current.is_end:
                words.append(path)
            for ch in sorted(current.children):
                collect(current.children[ch], path + ch)

        collect(node, prefix)
        return words`},
		{"drill-trie-delete", "removes the whole path even when other words share it", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def _count(self, node):
        return sum(1 + self._count(c) for c in node.children.values())

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def contains(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return node.is_end

    def startsWith(self, prefix):
        node = self.root
        for ch in prefix:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return True

    def countNodes(self):
        return self._count(self.root)

    def delete(self, word):
        if not self.contains(word):
            return False
        del self.root.children[word[0]]
        return True`},
		{"drill-trie-delete", "only clears the end flag and never prunes", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def _count(self, node):
        return sum(1 + self._count(c) for c in node.children.values())

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def _walk(self, text):
        node = self.root
        for ch in text:
            if ch not in node.children:
                return None
            node = node.children[ch]
        return node

    def contains(self, word):
        node = self._walk(word)
        return bool(node and node.is_end)

    def startsWith(self, prefix):
        return self._walk(prefix) is not None

    def countNodes(self):
        return self._count(self.root)

    def delete(self, word):
        if not self.contains(word):
            return False
        self._walk(word).is_end = False
        return True`},
		{"drill-trie-delete", "deletes a word that is only a prefix", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def _count(self, node):
        return sum(1 + self._count(c) for c in node.children.values())

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def _walk(self, text):
        node = self.root
        for ch in text:
            if ch not in node.children:
                return None
            node = node.children[ch]
        return node

    def contains(self, word):
        node = self._walk(word)
        return bool(node and node.is_end)

    def startsWith(self, prefix):
        return self._walk(prefix) is not None

    def countNodes(self):
        return self._count(self.root)

    def delete(self, word):
        path = [self.root]
        for ch in word:
            if ch not in path[-1].children:
                return False
            path.append(path[-1].children[ch])
        path[-1].is_end = False
        for i in range(len(word), 0, -1):
            if path[i].is_end or path[i].children:
                break
            del path[i - 1].children[word[i - 1]]
        return True`},
		{"drill-trie-delete", "prunes a node that still ends another word", `class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def _count(self, node):
        return sum(1 + self._count(c) for c in node.children.values())

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def _walk(self, text):
        node = self.root
        for ch in text:
            if ch not in node.children:
                return None
            node = node.children[ch]
        return node

    def contains(self, word):
        node = self._walk(word)
        return bool(node and node.is_end)

    def startsWith(self, prefix):
        return self._walk(prefix) is not None

    def countNodes(self):
        return self._count(self.root)

    def delete(self, word):
        if not self.contains(word):
            return False
        path = [self.root]
        for ch in word:
            path.append(path[-1].children[ch])
        path[-1].is_end = False
        for i in range(len(word), 0, -1):
            if path[i].children:
                break
            del path[i - 1].children[word[i - 1]]
        return True`},
		{"drill-trie-longest-common-prefix", "ignores word ends and keeps walking single-child nodes", `def longestCommonPrefix(strs):
    root = {}
    for word in strs:
        node = root
        for ch in word:
            node = node.setdefault(ch, {})
    prefix = []
    node = root
    while len(node) == 1:
        ch = next(iter(node))
        prefix.append(ch)
        node = node[ch]
    return "".join(prefix)`},
		{"drill-trie-longest-common-prefix", "returns the first string without looking for branches", `def longestCommonPrefix(strs):
    return strs[0]`},
		{"drill-trie-longest-common-prefix", "skips empty strings when building", `def longestCommonPrefix(strs):
    words = [w for w in strs if w]
    first, last = min(words), max(words)
    i = 0
    while i < len(first) and i < len(last) and first[i] == last[i]:
        i += 1
    return first[:i]`},
	}
	for _, c := range cases {
		c := c
		t.Run(c.id+"/"+c.why, func(t *testing.T) {
			t.Parallel()
			var p *practice.Problem
			for _, q := range practice.EmbeddedProblems {
				if q.ID == c.id {
					p = q
				}
			}
			if p == nil {
				t.Fatalf("no problem %s", c.id)
			}
			for i := range p.TestCases {
				res, err := Run(p, c.code, &p.TestCases[i], Options{})
				if err != nil {
					t.Fatalf("case %d: %v", i+1, err)
				}
				if !res.Passed {
					return
				}
			}
			t.Errorf("code that %s passed every case", c.why)
		})
	}
}
