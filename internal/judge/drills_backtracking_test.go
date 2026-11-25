package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestBacktrackingDrillsRejectMistakes checks that each backtracking
// foundations drill fails the mistake its tests and notes promise to catch, so
// a drill can't be passed without doing the thing it drills. The unpruned
// subset-sum searches are caught by the 40-number time-limited case.
func TestBacktrackingDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-binary-strings", "forgets path.pop()", `def binaryStrings(n):
    result = []
    path = []

    def backtrack():
        if len(path) == n:
            result.append("".join(path))
            return
        for digit in "01":
            path.append(digit)
            backtrack()
            pass

    backtrack()
    return result
`},
		{"drill-binary-strings", "records the string but does not return", `def binaryStrings(n):
    result = []
    path = []

    def backtrack():
        if len(path) == n:
            result.append("".join(path))
            pass
        for digit in "01":
            path.append(digit)
            backtrack()
            path.pop()

    backtrack()
    return result
`},
		{"drill-binary-strings", "tries 1 before 0", `def binaryStrings(n):
    result = []
    path = []

    def backtrack():
        if len(path) == n:
            result.append("".join(path))
            return
        for digit in "10":
            path.append(digit)
            backtrack()
            path.pop()

    backtrack()
    return result
`},
		{"drill-binary-strings", "stops one digit too late", `def binaryStrings(n):
    result = []
    path = []

    def backtrack():
        if len(path) > n:
            result.append("".join(path))
            return
        for digit in "01":
            path.append(digit)
            backtrack()
            path.pop()

    backtrack()
    return result
`},
		{"drill-letter-case-permutation", "gives digits two options, so results repeat", `def letterCasePermutation(s):
    result = []
    path = []

    def backtrack(i):
        if i == len(s):
            result.append("".join(path))
            return
        options = [s[i]]
        if True:
            options = [s[i].lower(), s[i].upper()]
        for ch in options:
            path.append(ch)
            backtrack(i + 1)
            path.pop()

    backtrack(0)
    return result
`},
		{"drill-letter-case-permutation", "forgets path.pop()", `def letterCasePermutation(s):
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
            pass

    backtrack(0)
    return result
`},
		{"drill-letter-case-permutation", "lists the upper case form and the letter as written, repeating upper case input", `def letterCasePermutation(s):
    result = []
    path = []

    def backtrack(i):
        if i == len(s):
            result.append("".join(path))
            return
        options = [s[i]]
        if s[i].isalpha():
            options = [s[i].upper(), s[i]]
        for ch in options:
            path.append(ch)
            backtrack(i + 1)
            path.pop()

    backtrack(0)
    return result
`},
		{"drill-combinations", "appends path instead of a copy", `def combine(n, k):
    result = []
    path = []

    def backtrack(start):
        if len(path) == k:
            result.append(path)
            return
        for number in range(start, n + 1):
            path.append(number)
            backtrack(number + 1)
            path.pop()

    backtrack(1)
    return result
`},
		{"drill-combinations", "restarts the loop at 1, so numbers repeat and reorder", `def combine(n, k):
    result = []
    path = []

    def backtrack(start):
        if len(path) == k:
            result.append(path[:])
            return
        for number in range(1, n + 1):
            path.append(number)
            backtrack(number + 1)
            path.pop()

    backtrack(1)
    return result
`},
		{"drill-combinations", "recurses with start + 1 instead of number + 1", `def combine(n, k):
    result = []
    path = []

    def backtrack(start):
        if len(path) == k:
            result.append(path[:])
            return
        for number in range(start, n + 1):
            path.append(number)
            backtrack(start + 1)
            path.pop()

    backtrack(1)
    return result
`},
		{"drill-combinations", "recurses with number, so a number repeats", `def combine(n, k):
    result = []
    path = []

    def backtrack(start):
        if len(path) == k:
            result.append(path[:])
            return
        for number in range(start, n + 1):
            path.append(number)
            backtrack(number)
            path.pop()

    backtrack(1)
    return result
`},
		{"drill-combinations", "forgets path.pop()", `def combine(n, k):
    result = []
    path = []

    def backtrack(start):
        if len(path) == k:
            result.append(path[:])
            return
        for number in range(start, n + 1):
            path.append(number)
            backtrack(number + 1)
            pass

    backtrack(1)
    return result
`},
		{"drill-arrangements", "never clears used[i]", `def arrangements(letters, k):
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
            pass

    backtrack()
    return result
`},
		{"drill-arrangements", "never checks used, so letters repeat", `def arrangements(letters, k):
    result = []
    path = []
    used = [False] * len(letters)

    def backtrack():
        if len(path) == k:
            result.append("".join(path))
            return
        for i in range(len(letters)):
            if False:
                continue
            used[i] = True
            path.append(letters[i])
            backtrack()
            path.pop()
            used[i] = False

    backtrack()
    return result
`},
		{"drill-arrangements", "uses a start index and so only builds increasing selections", `def arrangements(letters, k):
    result = []
    path = []

    def backtrack(start):
        if len(path) == k:
            result.append("".join(path))
            return
        for i in range(start, len(letters)):
            path.append(letters[i])
            backtrack(i + 1)
            path.pop()

    backtrack(0)
    return result
`},
		{"drill-arrangements", "clears the flag but forgets to pop the letter", `def arrangements(letters, k):
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
            pass
            used[i] = False

    backtrack()
    return result
`},
		{"drill-subsets-with-sum", "never prunes a branch whose sum is already over the target", `def subsetsWithSum(nums, target):
    result = []
    path = []

    def backtrack(start, remaining):
        if remaining == 0:
            result.append(path[:])
            return
        for i in range(start, len(nums)):
            pass
            path.append(nums[i])
            backtrack(i + 1, remaining - nums[i])
            path.pop()

    backtrack(0, target)
    return result
`},
		{"drill-subsets-with-sum", "checks the sum only after building every subset", `def subsetsWithSum(nums, target):
    result = []
    path = []

    def backtrack(start):
        if sum(path) == target:
            result.append(path[:])
        for i in range(start, len(nums)):
            path.append(nums[i])
            backtrack(i + 1)
            path.pop()

    backtrack(0)
    return result
`},
		{"drill-subsets-with-sum", "recurses with i, so a number repeats", `def subsetsWithSum(nums, target):
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
            backtrack(i, remaining - nums[i])
            path.pop()

    backtrack(0, target)
    return result
`},
		{"drill-subsets-with-sum", "forgets path.pop()", `def subsetsWithSum(nums, target):
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
            pass

    backtrack(0, target)
    return result
`},
		{"drill-subsets-with-sum", "prunes with >= so a number that fits exactly is lost", `def subsetsWithSum(nums, target):
    result = []
    path = []

    def backtrack(start, remaining):
        if remaining == 0:
            result.append(path[:])
            return
        for i in range(start, len(nums)):
            if nums[i] >= remaining:
                break
            path.append(nums[i])
            backtrack(i + 1, remaining - nums[i])
            path.pop()

    backtrack(0, target)
    return result
`},
		{"drill-grid-paths", "never unmarks the cell", `def countPaths(grid):
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
        pass
        return total

    return backtrack(0, 0)
`},
		{"drill-grid-paths", "never marks the cell, so routes loop forever", `def countPaths(grid):
    rows, cols = len(grid), len(grid[0])
    visited = [[False] * cols for _ in range(rows)]

    def backtrack(row, col):
        if row < 0 or row >= rows or col < 0 or col >= cols:
            return 0
        if grid[row][col] == 1 or visited[row][col]:
            return 0
        if row == rows - 1 and col == cols - 1:
            return 1
        pass
        total = 0
        for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
            total += backtrack(row + dr, col + dc)
        pass
        return total

    return backtrack(0, 0)
`},
		{"drill-grid-paths", "moves only down and right", `def countPaths(grid):
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
        for dr, dc in ((1, 0), (0, 1)):
            total += backtrack(row + dr, col + dc)
        visited[row][col] = False
        return total

    return backtrack(0, 0)
`},
		{"drill-grid-paths", "does not return at the end cell, so routes continue through it", `def countPaths(grid):
    rows, cols = len(grid), len(grid[0])
    visited = [[False] * cols for _ in range(rows)]

    def backtrack(row, col):
        if row < 0 or row >= rows or col < 0 or col >= cols:
            return 0
        if grid[row][col] == 1 or visited[row][col]:
            return 0
        if row == rows - 1 and col == cols - 1:
            pass
        visited[row][col] = True
        total = 0
        for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
            total += backtrack(row + dr, col + dc)
        visited[row][col] = False
        return total

    return backtrack(0, 0)
`},
		{"drill-grid-paths", "unmarks only when a route was found", `def countPaths(grid):
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
        if total:
            visited[row][col] = False
        return total

    return backtrack(0, 0)
`},
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
