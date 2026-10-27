package judge

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"dsatutor/internal/practice"
	"dsatutor/internal/sandbox"
)

// oxigenSolutions are Oxigen answers to problems that between them use every
// harness feature: node conversions, in-place answers, design classes, both
// round-trip forms, validators and the clone checks.
var oxigenSolutions = map[string]string{
	"two-sum": `
fun twoSum(nums, target) {
    seen := {}
    each i, x in nums {
        need := target - x
        give [seen[need], i] when has(seen, need)
        seen[x] = i
    }
    give []
}`,
	"reverse-linked-list": `
fun reverseList(head) {
    prev := None
    node := head
    repeat when node != None {
        next := node.next
        node.next = prev
        prev := node
        node := next
    }
    give prev
}`,
	"merge-k-sorted-lists": `
introduce array

fun mergeKLists(lists) {
    vals := []
    each head in lists {
        node := head
        repeat when node != None {
            vals := push(vals, node.val)
            node := node.next
        }
    }
    dummy := ListNode(0)
    tail := dummy
    each v in array.sort(vals) {
        tail.next = ListNode(v)
        tail := tail.next
    }
    give dummy.next
}`,
	"linked-list-cycle": `
fun hasCycle(head) {
    slow := head
    fast := head
    repeat when fast != None and fast.next != None {
        slow := slow.next
        fast := fast.next.next
        give True when slow == fast
    }
    give False
}`,
	"reorder-list": `
fun reorderList(head) {
    nodes := []
    node := head
    repeat when node != None {
        nodes := push(nodes, node)
        node := node.next
    }
    i := 0
    j := len(nodes) - 1
    repeat when i < j {
        nodes[i].next = nodes[j]
        i := i + 1
        stop when i == j
        nodes[j].next = nodes[i]
        j := j - 1
    }
    nodes[i].next = None when len(nodes) > 0
}`,
	"maximum-depth-of-binary-tree": `
fun maxDepth(root) {
    give 0 when root == None
    left := maxDepth(root.left)
    right := maxDepth(root.right)
    give 1 + option { left > right, left, right }
}`,
	"invert-binary-tree": `
fun invertTree(root) {
    give None when root == None
    left := root.left
    root.left = invertTree(root.right)
    root.right = invertTree(left)
    give root
}`,
	"lowest-common-ancestor-bst": `
fun lowestCommonAncestor(root, p, q) {
    node := root
    repeat when node != None {
        option {
            p.val < node.val and q.val < node.val -> { node := node.left }
            p.val > node.val and q.val > node.val -> { node := node.right }
            { give node }
        }
    }
    give None
}`,
	"serialize-and-deserialize-binary-tree": `
introduce strings

struct Codec {
}

Codec includes {
    fun serialize(root) {
        out := []
        stack := [root]
        repeat when len(stack) > 0 {
            node := last(stack)
            stack := stack[0:len(stack) - 1]
            option {
                node == None -> { out := push(out, "#") }
                {
                    out := push(out, str(node.val))
                    stack := push(stack, node.right)
                    stack := push(stack, node.left)
                }
            }
        }
        give strings.join(out, ",")
    }

    fun deserialize(data) {
        tokens := strings.split(data, ",")
        pos := [0]
        give self.build(tokens, pos)
    }

    fun build(tokens, pos) {
        tok := tokens[pos[0]]
        pos[0] = pos[0] + 1
        give None when tok == "#"
        node := TreeNode(int(tok))
        node.left = self.build(tokens, pos)
        node.right = self.build(tokens, pos)
        give node
    }
}`,
	"encode-and-decode-strings": `
introduce strings

fun encode(strs) {
    out := ""
    each s in strs {
        out := out + str(len(s)) + "#" + s
    }
    give out
}

fun decode(s) {
    out := []
    i := 0
    repeat when i < len(s) {
        j := i
        repeat when s[j] != "#" {
            j := j + 1
        }
        n := int(s[i:j])
        out := push(out, s[j + 1:j + 1 + n])
        i := j + 1 + n
    }
    give out
}`,
	"min-stack": `
struct MinStack {
    items <array>
    mins <array>
}

MinStack includes {
    fun init() {
        self.items = []
        self.mins = []
    }

    fun push(val) {
        self.items = push(self.items, val)
        option {
            len(self.mins) == 0 or val <= last(self.mins) -> { self.mins = push(self.mins, val) }
        }
    }

    fun pop() {
        top := last(self.items)
        self.items = self.items[0:len(self.items) - 1]
        option {
            top == last(self.mins) -> { self.mins = self.mins[0:len(self.mins) - 1] }
        }
    }

    fun top() { give last(self.items) }

    fun getMin() { give last(self.mins) }
}`,
	"time-based-key-value-store": `
struct TimeMap {
    store <map>
}

TimeMap includes {
    fun init() { self.store = {} }

    fun set(key, value, timestamp) {
        option {
            not has(self.store, key) -> { self.store[key] = [] }
        }
        self.store[key] = push(self.store[key], [timestamp, value])
    }

    fun get(key, timestamp) {
        give "" when not has(self.store, key)
        entries := self.store[key]
        lo := 0
        hi := len(entries) - 1
        best := ""
        repeat when lo <= hi {
            mid := (lo + hi) / 2
            option {
                entries[mid][0] <= timestamp -> {
                    best := entries[mid][1]
                    lo := mid + 1
                }
                { hi := mid - 1 }
            }
        }
        give best
    }
}`,
	"clone-graph": `
fun cloneGraph(node) {
    give None when node == None
    copies := {}
    order := [node]
    copies[node.val] = Node(node.val, [])
    i := 0
    repeat when i < len(order) {
        cur := order[i]
        i := i + 1
        each n in cur.neighbors {
            option {
                not has(copies, n.val) -> {
                    copies[n.val] = Node(n.val, [])
                    order := push(order, n)
                }
            }
            copies[cur.val].neighbors = push(copies[cur.val].neighbors, copies[n.val])
        }
    }
    give copies[node.val]
}`,
	"copy-list-with-random-pointer": `
fun copyRandomList(head) {
    nodes := []
    node := head
    repeat when node != None {
        nodes := push(nodes, node)
        node := node.next
    }
    copies := []
    each n in nodes {
        copies := push(copies, Node(n.val))
    }
    each i, n in nodes {
        copies[i].next = copies[i + 1] when i + 1 < len(copies)
        option {
            n.random != None -> {
                each j, m in nodes {
                    option {
                        m == n.random -> {
                            copies[i].random = copies[j]
                            stop
                        }
                    }
                }
            }
        }
    }
    give option { len(copies) > 0, copies[0], None }
}`,
	"course-schedule-ii": `
fun findOrder(numCourses, prerequisites) {
    indegree := []
    graph := []
    each i in range(numCourses) {
        indegree := push(indegree, 0)
        graph := push(graph, [])
    }
    each pair in prerequisites {
        graph[pair[1]] = push(graph[pair[1]], pair[0])
        indegree[pair[0]] = indegree[pair[0]] + 1
    }
    queue := []
    each i in range(numCourses) {
        option {
            indegree[i] == 0 -> { queue := push(queue, i) }
        }
    }
    order := []
    head := 0
    repeat when head < len(queue) {
        c := queue[head]
        head := head + 1
        order := push(order, c)
        each nxt in graph[c] {
            indegree[nxt] = indegree[nxt] - 1
            option {
                indegree[nxt] == 0 -> { queue := push(queue, nxt) }
            }
        }
    }
    give option { len(order) == numCourses, order, [] }
}`,
	"rotate-image": `
fun rotate(matrix) {
    n := len(matrix)
    each i in range(n) {
        each j in range(i + 1, n) {
            tmp := matrix[i][j]
            matrix[i][j] = matrix[j][i]
            matrix[j][i] = tmp
        }
    }
    each row in matrix {
        lo := 0
        hi := n - 1
        repeat when lo < hi {
            tmp := row[lo]
            row[lo] = row[hi]
            row[hi] = tmp
            lo := lo + 1
            hi := hi - 1
        }
    }
}`,
}

func requireOxigen(t *testing.T) {
	t.Helper()
	if !sandbox.OxigenAvailable() {
		t.Skip("oxigen is not installed")
	}
}

func problemByID(t *testing.T, id string) *practice.Problem {
	t.Helper()
	for _, p := range practice.EmbeddedProblems {
		if p.ID == id {
			return p
		}
	}
	t.Fatalf("no problem %s", id)
	return nil
}

// TestOxigenSolutions judges the Oxigen answers above on every test case,
// hidden ones included, the way Submit does.
func TestOxigenSolutions(t *testing.T) {
	requireOxigen(t)
	for id, code := range oxigenSolutions {
		id, code := id, code
		t.Run(id, func(t *testing.T) {
			t.Parallel()
			p := problemByID(t, id)
			for i := range p.TestCases {
				tc := &p.TestCases[i]
				res, err := Run(p, code, tc, Options{Language: "oxigen"})
				if err != nil {
					t.Fatalf("case %d: %v", i+1, err)
				}
				if !res.Passed {
					t.Errorf("case %d (%s): expected %v, got %v; error: %s; message: %s; printed: %s",
						i+1, tc.Note, jsonish(tc.Expected), jsonish(res.Actual), res.Error, res.Message, res.Output)
				}
			}
		})
	}
}

// TestOxigenRejects checks that wrong Oxigen answers fail, including a
// clone that hands back the original nodes.
func TestOxigenRejects(t *testing.T) {
	requireOxigen(t)
	cases := []struct{ id, code, why string }{
		{"two-sum", "fun twoSum(nums, target) {\n    give [0, 0]\n}", "wrong answer"},
		{"clone-graph", "fun cloneGraph(node) {\n    give node\n}", "returns the original graph"},
		{"copy-list-with-random-pointer", "fun copyRandomList(head) {\n    give head\n}", "returns the original list"},
		{"course-schedule-ii", "fun findOrder(numCourses, prerequisites) {\n    give range(numCourses - 1, -1, -1)\n}", "order breaks prerequisites"},
		{"reverse-linked-list", "fun reverseList(head) {\n    give missing_name\n}", "runtime error"},
	}
	for _, c := range cases {
		p := problemByID(t, c.id)
		failedOne := false
		for i := range p.TestCases {
			res, err := Run(p, c.code, &p.TestCases[i], Options{Language: "oxigen"})
			if err != nil {
				t.Fatalf("%s: %v", c.id, err)
			}
			if !res.Passed {
				failedOne = true
				break
			}
		}
		if !failedOne {
			t.Errorf("%s (%s) passed every case", c.id, c.why)
		}
	}
}

// TestOxigenStartersParse runs `oxigen check` on the generated starter of
// every problem, with the node structs the judge would add.
func TestOxigenStartersParse(t *testing.T) {
	requireOxigen(t)
	dir := t.TempDir()
	for _, p := range practice.EmbeddedProblems {
		starter := practice.OxigenStarter(p)
		if !strings.Contains(starter, "fun ") {
			t.Errorf("%s: starter has no function:\n%s", p.ID, starter)
			continue
		}
		path := filepath.Join(dir, p.ID+".oxi")
		structs := "\nstruct ListNode {\n    val <generic>\n    next <ListNode> || <None>\n}\n"
		if err := os.WriteFile(path, []byte(starter+structs), 0600); err != nil {
			t.Fatal(err)
		}
		out, err := exec.Command(sandbox.OxigenBinary, "check", path).CombinedOutput()
		if err != nil {
			t.Errorf("%s: starter doesn't parse:\n%s\n%s", p.ID, starter, out)
		}
	}
}

// TestOxigenNameErrors checks that renaming what the judge calls produces a
// message about the learner's code, not about the generated harness.
func TestOxigenNameErrors(t *testing.T) {
	requireOxigen(t)
	cases := []struct{ id, code, want string }{
		{"reverse-linked-list", "fun reverse(head) {\n    give head\n}", "no function reverseList"},
		{"min-stack", "struct Stack {\n}\nStack includes {\n    fun init() {\n    }\n}", "no struct MinStack"},
		{"min-stack", "struct MinStack {\n}\nMinStack includes {\n    fun init() {\n    }\n}", "MinStack has no method push"},
	}
	for _, c := range cases {
		p := problemByID(t, c.id)
		res, err := Run(p, c.code, &p.TestCases[0], Options{Language: "oxigen"})
		if err != nil {
			t.Fatalf("%s: %v", c.id, err)
		}
		if !strings.Contains(res.Error, c.want) || strings.Contains(res.Error, "harness") {
			t.Errorf("%s: error %q, want it to mention %q", c.id, res.Error, c.want)
		}
	}
}

func TestValidateOxigenCode(t *testing.T) {
	ok := []string{
		"introduce math\nfun f() { give math.sqrt(4) }",
		"introduce {split, join} from strings",
		"// introduce os would be refused, but this is a comment\nfun f() {}",
	}
	bad := []string{
		"introduce os",
		"intro io",
		"introduce {read} from json",
		"introduce .helpers",
		"main { introduce net }",
		"fun f() { give __read_file(\"/etc/passwd\") }",
	}
	for _, code := range ok {
		if err := sandbox.ValidateOxigenCode(code); err != nil {
			t.Errorf("rejected safe code %q: %v", code, err)
		}
	}
	for _, code := range bad {
		if err := sandbox.ValidateOxigenCode(code); err == nil {
			t.Errorf("accepted unsafe code %q", code)
		}
	}
}
