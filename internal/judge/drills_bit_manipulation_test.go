package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestBitManipulationDrillsRejectMistakes checks that each bit manipulation
// foundations drill fails the mistake its tests and notes promise to catch.
// The mistakes are the ones Python makes easy: operator precedence (+ and -
// bind tighter than <<, ~ tighter than <<), the wrong identity (n & (n + 1),
// x & (x - 1) for x & -x), off-by-one bit positions and missing edge cases
// (zero, 1, the order of bits). Python integers do not overflow, so the
// 32-bit mask pitfall belongs to Sum of Two Integers in the ladder, and the
// negative >> loop has no valid input here (every drill takes n >= 0 except
// Power of Two and Single Number III, which have no shifting loop to hang).
func TestBitManipulationDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-to-binary", "returns an empty string for 0", `def toBinary(n):
    bits = []
    while n > 0:
        bits.append(str(n % 2))
        n //= 2
    return "".join(reversed(bits))`},
		{"drill-to-binary", "joins the digits without reversing", `def toBinary(n):
    if n == 0:
        return "0"
    bits = []
    while n > 0:
        bits.append(str(n % 2))
        n //= 2
    return "".join(bits)`},
		{"drill-to-binary", "stops at n > 1 and drops the leading 1", `def toBinary(n):
    if n == 0:
        return "0"
    bits = []
    while n > 1:
        bits.append(str(n % 2))
        n //= 2
    return "".join(reversed(bits))`},
		{"drill-to-binary", "returns bin(n) with its 0b prefix", `def toBinary(n):
    return bin(n)`},
		{"drill-from-binary", "reads the string from the right end", `def fromBinary(bits):
    value = 0
    for char in reversed(bits):
        value = value * 2 + int(char)
    return value`},
		{"drill-from-binary", "writes value << 1 + digit without parentheses", `def fromBinary(bits):
    value = 0
    for char in bits:
        value = value << 1 + int(char)
    return value`},
		{"drill-from-binary", "uses place values that are one power too high", `def fromBinary(bits):
    total = 0
    n = len(bits)
    for i, char in enumerate(bits):
        total += int(char) * 2 ** (n - i)
    return total`},
		{"drill-from-binary", "counts the ones without doubling", `def fromBinary(bits):
    value = 0
    for char in bits:
        value += int(char)
    return value`},
		{"drill-bit-ops", "returns the masked value instead of 0 or 1 for get", `def bitOps(n, i):
    mask = 1 << i
    return [n & mask, n | mask, n & ~mask, n ^ mask]`},
		{"drill-bit-ops", "writes n & ~1 << i for clear", `def bitOps(n, i):
    mask = 1 << i
    return [(n >> i) & 1, n | mask, n & ~1 << i, n ^ mask]`},
		{"drill-bit-ops", "clears with XOR", `def bitOps(n, i):
    mask = 1 << i
    return [(n >> i) & 1, n | mask, n ^ mask, n ^ mask]`},
		{"drill-bit-ops", "clears with subtraction", `def bitOps(n, i):
    mask = 1 << i
    return [(n >> i) & 1, n | mask, n - mask, n ^ mask]`},
		{"drill-bit-ops", "sets with addition", `def bitOps(n, i):
    mask = 1 << i
    return [(n >> i) & 1, n + mask, n & ~mask, n ^ mask]`},
		{"drill-bit-ops", "sets with XOR", `def bitOps(n, i):
    mask = 1 << i
    return [(n >> i) & 1, n ^ mask, n & ~mask, n ^ mask]`},
		{"drill-bit-ops", "builds the mask one position too high", `def bitOps(n, i):
    mask = 1 << (i + 1)
    return [(n >> (i + 1)) & 1, n | mask, n & ~mask, n ^ mask]`},
		{"drill-power-of-two", "has no n > 0 test", `def isPowerOfTwo(n):
    return n & (n - 1) == 0`},
		{"drill-power-of-two", "tests only whether n is even", `def isPowerOfTwo(n):
    return n > 0 and n & 1 == 0`},
		{"drill-power-of-two", "uses n & (n + 1)", `def isPowerOfTwo(n):
    return n > 0 and n & (n + 1) == 0`},
		{"drill-power-of-two", "rejects 1", `def isPowerOfTwo(n):
    return n > 1 and n & (n - 1) == 0`},
		{"drill-set-bit-positions", "reports 1-indexed positions", `def setBitPositions(n):
    positions = []
    while n:
        positions.append((n & -n).bit_length())
        n &= n - 1
    return positions`},
		{"drill-set-bit-positions", "appends the isolated value, not its position", `def setBitPositions(n):
    positions = []
    while n:
        positions.append(n & -n)
        n &= n - 1
    return positions`},
		{"drill-set-bit-positions", "lists the positions from the top down", `def setBitPositions(n):
    return [i for i in range(30, -1, -1) if (n >> i) & 1]`},
		{"drill-hamming-distance", "counts how far apart the values are", `def hammingDistance(x, y):
    diff = abs(x - y)
    count = 0
    while diff:
        diff &= diff - 1
        count += 1
    return count`},
		{"drill-hamming-distance", "counts the bits of x | y", `def hammingDistance(x, y):
    diff = x | y
    count = 0
    while diff:
        diff &= diff - 1
        count += 1
    return count`},
		{"drill-hamming-distance", "counts the bits of x & y", `def hammingDistance(x, y):
    diff = x & y
    count = 0
    while diff:
        diff &= diff - 1
        count += 1
    return count`},
		{"drill-hamming-distance", "returns x ^ y without counting", `def hammingDistance(x, y):
    return x ^ y`},
		{"drill-hamming-distance", "stops while diff > 1", `def hammingDistance(x, y):
    diff = x ^ y
    count = 0
    while diff > 1:
        diff &= diff - 1
        count += 1
    return count`},
		{"drill-two-single-numbers", "splits on the last bit", `def singleNumber(nums):
    a = b = 0
    for num in nums:
        if num & 1:
            a ^= num
        else:
            b ^= num
    return [a, b]`},
		{"drill-two-single-numbers", "splits on x & (x - 1)", `def singleNumber(nums):
    x = 0
    for num in nums:
        x ^= num
    low = x & (x - 1)
    a = b = 0
    for num in nums:
        if num & low:
            a ^= num
        else:
            b ^= num
    return [a, b]`},
		{"drill-two-single-numbers", "splits on all the bits of x", `def singleNumber(nums):
    x = 0
    for num in nums:
        x ^= num
    a = b = 0
    for num in nums:
        if num & x:
            a ^= num
        else:
            b ^= num
    return [a, b]`},
		{"drill-two-single-numbers", "returns the XOR of the two answers", `def singleNumber(nums):
    x = 0
    for num in nums:
        x ^= num
    return [x, 0]`},
		{"drill-subsets-bitmask", "writes range(1 << n - 1)", `def subsets(nums):
    n = len(nums)
    result = []
    for mask in range(1 << n - 1):
        result.append([nums[i] for i in range(n) if (mask >> i) & 1])
    return result`},
		{"drill-subsets-bitmask", "tests mask & i instead of mask & (1 << i)", `def subsets(nums):
    n = len(nums)
    result = []
    for mask in range(1 << n):
        result.append([nums[i] for i in range(n) if mask & i])
    return result`},
		{"drill-subsets-bitmask", "walks the positions from the top down", `def subsets(nums):
    n = len(nums)
    result = []
    for mask in range(1 << n):
        result.append([nums[i] for i in range(n - 1, -1, -1) if (mask >> i) & 1])
    return result`},
		{"drill-subsets-bitmask", "lists subsets by size with combinations", `def subsets(nums):
    from itertools import combinations
    result = []
    for size in range(len(nums) + 1):
        for combo in combinations(nums, size):
            result.append(list(combo))
    return result`},
		{"drill-subsets-bitmask", "reads bit i + 1 of the mask", `def subsets(nums):
    n = len(nums)
    result = []
    for mask in range(1 << n):
        result.append([nums[i] for i in range(n) if (mask >> (i + 1)) & 1])
    return result`},
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
