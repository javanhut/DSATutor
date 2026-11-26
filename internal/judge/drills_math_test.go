package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestMathGeometryDrillsRejectMistakes checks that each math and geometry
// foundations drill fails the mistake its tests and notes promise to catch:
// ignored signs, off-by-one bounds (the square root, the strict "less than n"),
// shared rows, float slopes, and the slow versions that only the large cases
// (counting down for a gcd, a prime test per number, multiplying exp times)
// can tell from the real algorithm.
func TestMathGeometryDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-reverse-digits", "ignores negative numbers (loops on n > 0 only)", `def reverseDigits(n):
    result = 0
    while n > 0:
        result = result * 10 + n % 10
        n //= 10
    return result`},
		{"drill-reverse-digits", "shifts after adding the digit", `def reverseDigits(n):
    sign = -1 if n < 0 else 1
    n = abs(n)
    result = 0
    while n > 0:
        result = (result + n % 10) * 10
        n //= 10
    return sign * result`},
		{"drill-reverse-digits", "reverses the string of a negative number with its sign", `def reverseDigits(n):
    return int(str(n)[::-1])`},
		{"drill-gcd-lcm", "computes the lcm as a * b", `def gcdLcm(a, b):
    x, y = a, b
    while y:
        x, y = y, x % y
    return [x, a * b]`},
		{"drill-gcd-lcm", "stops the Euclid loop when the remainder reaches 1", `def gcdLcm(a, b):
    x, y = a, b
    while y > 1:
        x, y = y, x % y
    return [x, a // x * b]`},
		{"drill-gcd-lcm", "counts down from the smaller number", `def gcdLcm(a, b):
    g = min(a, b)
    while a % g or b % g:
        g -= 1
    return [g, a // g * b]`},
		{"drill-is-prime", "treats 0 and 1 as prime", `def isPrime(n):
    divisor = 2
    while divisor * divisor <= n:
        if n % divisor == 0:
            return False
        divisor += 1
    return True`},
		{"drill-is-prime", "stops before the square root", `def isPrime(n):
    if n < 2:
        return False
    for divisor in range(2, int(n ** 0.5)):
        if n % divisor == 0:
            return False
    return True`},
		{"drill-is-prime", "tries every divisor up to n - 1", `def isPrime(n):
    if n < 2:
        return False
    for divisor in range(2, n):
        if n % divisor == 0:
            return False
    return True`},
		{"drill-sieve-primes", "counts primes up to and including n", `def countPrimes(n):
    n += 1
    if n < 3:
        return 0
    is_prime = [True] * n
    is_prime[0] = is_prime[1] = False
    for p in range(2, int(n ** 0.5) + 1):
        if is_prime[p]:
            for multiple in range(p * p, n, p):
                is_prime[multiple] = False
    return sum(is_prime)`},
		{"drill-sieve-primes", "stops the outer loop before the square root", `def countPrimes(n):
    if n < 3:
        return 0
    is_prime = [True] * n
    is_prime[0] = is_prime[1] = False
    for p in range(2, int(n ** 0.5)):
        if is_prime[p]:
            for multiple in range(p * p, n, p):
                is_prime[multiple] = False
    return sum(is_prime)`},
		{"drill-sieve-primes", "crosses out from p so p itself is removed", `def countPrimes(n):
    if n < 3:
        return 0
    is_prime = [True] * n
    is_prime[0] = is_prime[1] = False
    for p in range(2, int(n ** 0.5) + 1):
        if is_prime[p]:
            for multiple in range(p, n, p):
                is_prime[multiple] = False
    return sum(is_prime)`},
		{"drill-sieve-primes", "does not mark 0 and 1 as not prime", `def countPrimes(n):
    if n < 3:
        return 0
    is_prime = [True] * n
    for p in range(2, int(n ** 0.5) + 1):
        if is_prime[p]:
            for multiple in range(p * p, n, p):
                is_prime[multiple] = False
    return sum(is_prime)`},
		{"drill-sieve-primes", "tests every number with trial division", `def countPrimes(n):
    count = 0
    for number in range(2, n):
        divisor = 2
        while divisor * divisor <= number:
            if number % divisor == 0:
                break
            divisor += 1
        else:
            count += 1
    return count`},
		{"drill-pow-mod", "starts the result at 1 even when mod is 1", `def powMod(base, exp, mod):
    result = 1
    base %= mod
    while exp > 0:
        if exp % 2 == 1:
            result = result * base % mod
        base = base * base % mod
        exp //= 2
    return result`},
		{"drill-pow-mod", "squares without reducing modulo mod", `def powMod(base, exp, mod):
    result = 1 % mod
    while exp > 0:
        if exp % 2 == 1:
            result = result * base % mod
        base = base * base
        exp //= 2
    return result`},
		{"drill-pow-mod", "stops the loop at exp > 1 and drops the top bit", `def powMod(base, exp, mod):
    result = 1 % mod
    base %= mod
    while exp > 1:
        if exp % 2 == 1:
            result = result * base % mod
        base = base * base % mod
        exp //= 2
    return result`},
		{"drill-pow-mod", "multiplies exp times", `def powMod(base, exp, mod):
    result = 1 % mod
    for _ in range(exp):
        result = result * base % mod
    return result`},
		{"drill-transpose-matrix", "shares one row list between all rows", `def transpose(matrix):
    rows, cols = len(matrix), len(matrix[0])
    result = [[0] * rows] * cols
    for r in range(rows):
        for c in range(cols):
            result[c][r] = matrix[r][c]
    return result`},
		{"drill-transpose-matrix", "swaps every pair twice in place", `def transpose(matrix):
    n = len(matrix)
    for r in range(n):
        for c in range(n):
            matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
    return matrix`},
		{"drill-transpose-matrix", "allocates the input's shape instead of the flipped one", `def transpose(matrix):
    rows, cols = len(matrix), len(matrix[0])
    result = [[0] * cols for _ in range(rows)]
    for r in range(rows):
        for c in range(cols):
            result[c][r] = matrix[r][c]
    return result`},
		{"drill-diagonal-order", "starts every diagonal in row 0", `def findDiagonalOrder(mat):
    rows, cols = len(mat), len(mat[0])
    result = []
    for d in range(rows + cols - 1):
        r, c = 0, d
        diagonal = []
        while r < rows and c >= 0:
            diagonal.append(mat[r][c])
            r += 1
            c -= 1
        if d % 2 == 0:
            diagonal.reverse()
        result.extend(diagonal)
    return result`},
		{"drill-diagonal-order", "uses rows + cols - 2 diagonals", `def findDiagonalOrder(mat):
    rows, cols = len(mat), len(mat[0])
    result = []
    for d in range(rows + cols - 2):
        r = 0 if d < cols else d - cols + 1
        c = d - r
        diagonal = []
        while r < rows and c >= 0:
            diagonal.append(mat[r][c])
            r += 1
            c -= 1
        if d % 2 == 0:
            diagonal.reverse()
        result.extend(diagonal)
    return result`},
		{"drill-diagonal-order", "reverses the wrong parity", `def findDiagonalOrder(mat):
    rows, cols = len(mat), len(mat[0])
    result = []
    for d in range(rows + cols - 1):
        r = 0 if d < cols else d - cols + 1
        c = d - r
        diagonal = []
        while r < rows and c >= 0:
            diagonal.append(mat[r][c])
            r += 1
            c -= 1
        if d % 2 == 1:
            diagonal.reverse()
        result.extend(diagonal)
    return result`},
		{"drill-orientation", "compares slopes with floating-point division", `def orientation(p, q, r):
    s1 = (q[1] - p[1]) / (q[0] - p[0])
    s2 = (r[1] - p[1]) / (r[0] - p[0])
    return 0 if s1 == s2 else (1 if s2 > s1 else -1)`},
		{"drill-orientation", "forgets to subtract p", `def orientation(p, q, r):
    cross = q[0] * r[1] - q[1] * r[0]
    return (cross > 0) - (cross < 0)`},
		{"drill-orientation", "flips the sign of the cross product", `def orientation(p, q, r):
    cross = (q[1] - p[1]) * (r[0] - p[0]) - (q[0] - p[0]) * (r[1] - p[1])
    return (cross > 0) - (cross < 0)`},
		{"drill-orientation", "returns the raw cross product", `def orientation(p, q, r):
    return (q[0] - p[0]) * (r[1] - p[1]) - (q[1] - p[1]) * (r[0] - p[0])`},
		{"drill-orientation", "compares slopes after a float division that rounds near-collinear points together", `def orientation(p, q, r):
    if q[0] == p[0] or r[0] == p[0]:
        cross = (q[0] - p[0]) * (r[1] - p[1]) - (q[1] - p[1]) * (r[0] - p[0])
        return (cross > 0) - (cross < 0)
    s1 = (q[1] - p[1]) / (q[0] - p[0])
    s2 = (r[1] - p[1]) / (r[0] - p[0])
    if s1 == s2:
        return 0
    return 1 if (r[0] > p[0]) == (q[0] > p[0]) and s2 > s1 else -1`},
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
