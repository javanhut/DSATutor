package sandbox

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

// ComplexityClass represents a Big-O complexity class.
type ComplexityClass string

const (
	O1       ComplexityClass = "O(1)"
	OLogN    ComplexityClass = "O(log n)"
	ON       ComplexityClass = "O(n)"
	ONLogN   ComplexityClass = "O(n log n)"
	ON2      ComplexityClass = "O(n^2)"
	ON3      ComplexityClass = "O(n^3)"
	O2N      ComplexityClass = "O(2^n)"
	ONFact   ComplexityClass = "O(n!)"
	OUnknown ComplexityClass = "Unknown"
)

// ComplexityResult represents the estimated complexity of code.
type ComplexityResult struct {
	TimeComplexity  ComplexityClass `json:"timeComplexity"`
	SpaceComplexity ComplexityClass `json:"spaceComplexity"`
	Confidence      float64         `json:"confidence"` // 0.0 to 1.0
	Explanation     string          `json:"explanation"`
	Warning         string          `json:"warning,omitempty"`
}

// ComplexityAnalyzer analyzes code for time and space complexity.
type ComplexityAnalyzer struct {
	code      string
	lines     []string
	loopDepth int
}

// NewComplexityAnalyzer creates a new analyzer for the given code.
func NewComplexityAnalyzer(code string) *ComplexityAnalyzer {
	return &ComplexityAnalyzer{
		code:  code,
		lines: strings.Split(code, "\n"),
	}
}

// Analyze performs static analysis to estimate complexity.
func (ca *ComplexityAnalyzer) Analyze() ComplexityResult {
	timeComplexity := ca.analyzeTimeComplexity()
	spaceComplexity := ca.analyzeSpaceComplexity()

	return ComplexityResult{
		TimeComplexity:  timeComplexity,
		SpaceComplexity: spaceComplexity,
		Confidence:      0.7, // Static analysis has limited accuracy
		Explanation:     ca.generateExplanation(timeComplexity, spaceComplexity),
	}
}

// analyzeTimeComplexity estimates time complexity from code structure.
func (ca *ComplexityAnalyzer) analyzeTimeComplexity() ComplexityClass {
	maxLoopDepth := ca.countMaxLoopDepth()
	hasRecursion := ca.hasRecursion()
	hasBinarySearch := ca.hasBinarySearchPattern()
	hasSorting := ca.hasSortingCall()
	hasHeapOp := ca.hasHeapOperations()

	// Check for specific patterns
	if hasBinarySearch && maxLoopDepth <= 1 {
		return OLogN
	}

	if hasSorting {
		if maxLoopDepth <= 1 {
			return ONLogN
		}
		// Sorting with additional loops
		return ON2
	}

	if hasHeapOp && maxLoopDepth <= 1 {
		return ONLogN
	}

	// Recursive complexity estimation
	if hasRecursion {
		if ca.hasMemoization() {
			return ON // With memoization, often reduces to O(n)
		}
		if ca.hasExponentialRecursion() {
			return O2N
		}
		return ON // Default for simple recursion
	}

	// Loop-based complexity
	switch maxLoopDepth {
	case 0:
		return O1
	case 1:
		return ON
	case 2:
		return ON2
	case 3:
		return ON3
	default:
		return OUnknown
	}
}

// analyzeSpaceComplexity estimates space complexity from code structure.
func (ca *ComplexityAnalyzer) analyzeSpaceComplexity() ComplexityClass {
	hasRecursion := ca.hasRecursion()
	creates2DArray := ca.creates2DArray()
	createsArray := ca.createsArray()
	createsHashMap := ca.createsHashMap()

	if creates2DArray {
		return ON2
	}

	if hasRecursion && !ca.hasTailRecursion() {
		return ON // Recursion stack
	}

	if createsArray || createsHashMap {
		return ON
	}

	return O1
}

// countMaxLoopDepth counts the maximum nested loop depth.
func (ca *ComplexityAnalyzer) countMaxLoopDepth() int {
	maxDepth := 0
	currentDepth := 0

	forPattern := regexp.MustCompile(`^\s*(for|while)\s+`)

	for _, line := range ca.lines {
		if forPattern.MatchString(line) {
			currentDepth++
			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}
		}

		// Simple indent-based depth tracking (for Python)
		trimmed := strings.TrimLeft(line, " \t")
		if len(trimmed) > 0 && !strings.HasPrefix(trimmed, "#") {
			indent := len(line) - len(trimmed)
			// Rough heuristic: 4 spaces = 1 indent level
			level := indent / 4
			if level < currentDepth {
				currentDepth = level
			}
		}
	}

	return maxDepth
}

// hasRecursion checks if code contains recursive function calls.
func (ca *ComplexityAnalyzer) hasRecursion() bool {
	funcNamePattern := regexp.MustCompile(`def\s+(\w+)\s*\(`)
	matches := funcNamePattern.FindAllStringSubmatch(ca.code, -1)

	for _, match := range matches {
		if len(match) > 1 {
			funcName := match[1]
			// Check if function name appears in its own body
			callPattern := regexp.MustCompile(funcName + `\s*\(`)
			// Find where function is defined and check after
			defIdx := strings.Index(ca.code, "def "+funcName)
			if defIdx >= 0 {
				bodyStart := defIdx + len("def "+funcName)
				if callPattern.MatchString(ca.code[bodyStart:]) {
					return true
				}
			}
		}
	}

	return false
}

// hasBinarySearchPattern checks for binary search patterns.
func (ca *ComplexityAnalyzer) hasBinarySearchPattern() bool {
	patterns := []string{
		`mid\s*=`,
		`left\s*=.*right`,
		`low\s*=.*high`,
		`bisect`,
		`binary.*search`,
		`//\s*2`,
		`>>\s*1`,
	}

	for _, pattern := range patterns {
		if matched, _ := regexp.MatchString(`(?i)`+pattern, ca.code); matched {
			return true
		}
	}

	return false
}

// hasSortingCall checks for sorting function calls.
func (ca *ComplexityAnalyzer) hasSortingCall() bool {
	patterns := []string{
		`\.sort\(`,
		`sorted\(`,
		`heapq\.heapify`,
	}

	for _, pattern := range patterns {
		if matched, _ := regexp.MatchString(pattern, ca.code); matched {
			return true
		}
	}

	return false
}

// hasHeapOperations checks for heap operations.
func (ca *ComplexityAnalyzer) hasHeapOperations() bool {
	patterns := []string{
		`heapq\.`,
		`heappush`,
		`heappop`,
		`PriorityQueue`,
	}

	for _, pattern := range patterns {
		if matched, _ := regexp.MatchString(pattern, ca.code); matched {
			return true
		}
	}

	return false
}

// hasMemoization checks for memoization patterns.
func (ca *ComplexityAnalyzer) hasMemoization() bool {
	patterns := []string{
		`@lru_cache`,
		`@cache`,
		`@functools\.cache`,
		`memo\s*=`,
		`cache\s*=`,
		`dp\s*=`,
	}

	for _, pattern := range patterns {
		if matched, _ := regexp.MatchString(`(?i)`+pattern, ca.code); matched {
			return true
		}
	}

	return false
}

// hasExponentialRecursion checks for patterns suggesting exponential recursion.
func (ca *ComplexityAnalyzer) hasExponentialRecursion() bool {
	// Check for two recursive calls on same line or consecutive lines
	funcNamePattern := regexp.MustCompile(`def\s+(\w+)\s*\(`)
	matches := funcNamePattern.FindAllStringSubmatch(ca.code, -1)

	for _, match := range matches {
		if len(match) > 1 {
			funcName := match[1]
			// Check for pattern like: f(n-1) + f(n-2)
			doubleCallPattern := regexp.MustCompile(funcName + `\s*\([^)]*\)\s*[+]\s*` + funcName)
			if doubleCallPattern.MatchString(ca.code) {
				return true
			}
		}
	}

	return false
}

// hasTailRecursion checks if recursion is tail-recursive.
func (ca *ComplexityAnalyzer) hasTailRecursion() bool {
	// Simple heuristic: check if return statement directly calls the function
	funcNamePattern := regexp.MustCompile(`def\s+(\w+)\s*\(`)
	matches := funcNamePattern.FindAllStringSubmatch(ca.code, -1)

	for _, match := range matches {
		if len(match) > 1 {
			funcName := match[1]
			tailPattern := regexp.MustCompile(`return\s+` + funcName + `\s*\(`)
			if tailPattern.MatchString(ca.code) {
				return true
			}
		}
	}

	return false
}

// creates2DArray checks if code creates 2D arrays/matrices.
func (ca *ComplexityAnalyzer) creates2DArray() bool {
	patterns := []string{
		`\[\s*\[`,                   // [[...]]
		`\[\s*\]\s*\*`,              // [] * n
		`for.*for.*\[\]`,            // nested loops creating arrays
		`\[\[.*\]\s*for.*for.*\]`,   // list comprehension 2D
		`dp\s*=\s*\[\[`,             // dp = [[
		`matrix\s*=\s*\[\[`,         // matrix = [[
		`grid\s*=\s*\[\[`,           // grid = [[
	}

	for _, pattern := range patterns {
		if matched, _ := regexp.MatchString(pattern, ca.code); matched {
			return true
		}
	}

	return false
}

// createsArray checks if code creates arrays/lists.
func (ca *ComplexityAnalyzer) createsArray() bool {
	patterns := []string{
		`\[\s*\]\s*$`,       // = []
		`list\s*\(`,         // list()
		`\[.*for.*in.*\]`,   // list comprehension
		`\[\s*0\s*\]\s*\*`,  // [0] * n
	}

	for _, pattern := range patterns {
		if matched, _ := regexp.MatchString(pattern, ca.code); matched {
			return true
		}
	}

	return false
}

// createsHashMap checks if code creates hash maps/dictionaries/sets.
func (ca *ComplexityAnalyzer) createsHashMap() bool {
	patterns := []string{
		`\{\s*\}`,                    // {}
		`dict\s*\(`,                  // dict()
		`set\s*\(`,                   // set()
		`defaultdict`,               // defaultdict
		`Counter\s*\(`,              // Counter()
		`\{.*:.*for.*in.*\}`,        // dict comprehension
		`\{.*for.*in.*\}`,           // set comprehension
	}

	for _, pattern := range patterns {
		if matched, _ := regexp.MatchString(pattern, ca.code); matched {
			return true
		}
	}

	return false
}

// generateExplanation creates a human-readable explanation of the complexity.
func (ca *ComplexityAnalyzer) generateExplanation(time, space ComplexityClass) string {
	var parts []string

	parts = append(parts, fmt.Sprintf("Time: %s", time))
	parts = append(parts, fmt.Sprintf("Space: %s", space))

	if ca.hasRecursion() {
		parts = append(parts, "Uses recursion")
	}
	if ca.hasSortingCall() {
		parts = append(parts, "Uses sorting")
	}
	if ca.hasBinarySearchPattern() {
		parts = append(parts, "Uses binary search pattern")
	}
	if ca.hasMemoization() {
		parts = append(parts, "Uses memoization/DP")
	}

	return strings.Join(parts, "; ")
}

// ValidateComplexity checks if code meets expected complexity requirements.
func ValidateComplexity(code, expectedTime, expectedSpace string) ComplexityResult {
	analyzer := NewComplexityAnalyzer(code)
	result := analyzer.Analyze()

	// Compare with expected
	expectedTimeClass := parseComplexityClass(expectedTime)
	expectedSpaceClass := parseComplexityClass(expectedSpace)

	if !complexityMeets(result.TimeComplexity, expectedTimeClass) {
		result.Warning = fmt.Sprintf("Time complexity %s may not meet required %s",
			result.TimeComplexity, expectedTimeClass)
	}

	if !complexityMeets(result.SpaceComplexity, expectedSpaceClass) {
		if result.Warning != "" {
			result.Warning += "; "
		}
		result.Warning += fmt.Sprintf("Space complexity %s may not meet required %s",
			result.SpaceComplexity, expectedSpaceClass)
	}

	return result
}

// parseComplexityClass converts a string to ComplexityClass.
func parseComplexityClass(s string) ComplexityClass {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, " ", "")

	switch s {
	case "o(1)", "constant":
		return O1
	case "o(logn)", "o(log n)", "logarithmic":
		return OLogN
	case "o(n)", "linear":
		return ON
	case "o(nlogn)", "o(n log n)", "linearithmic":
		return ONLogN
	case "o(n^2)", "o(n2)", "quadratic":
		return ON2
	case "o(n^3)", "o(n3)", "cubic":
		return ON3
	case "o(2^n)", "o(2n)", "exponential":
		return O2N
	case "o(n!)", "factorial":
		return ONFact
	default:
		return OUnknown
	}
}

// complexityMeets checks if actual complexity meets or beats expected.
func complexityMeets(actual, expected ComplexityClass) bool {
	if expected == OUnknown {
		return true // Can't validate against unknown
	}

	order := map[ComplexityClass]int{
		O1:       1,
		OLogN:    2,
		ON:       3,
		ONLogN:   4,
		ON2:      5,
		ON3:      6,
		O2N:      7,
		ONFact:   8,
		OUnknown: 100,
	}

	return order[actual] <= order[expected]
}

// EstimateFromExecution estimates complexity from execution trace data.
func EstimateFromExecution(stepCounts []int, inputSizes []int) ComplexityClass {
	if len(stepCounts) < 3 || len(inputSizes) < 3 {
		return OUnknown
	}

	// Calculate ratios between consecutive runs
	// For O(n): ratio should be ~1
	// For O(n^2): ratio should be ~(n2/n1)^2
	// For O(log n): ratio should be ~log(n2)/log(n1)

	n1, n2 := float64(inputSizes[0]), float64(inputSizes[len(inputSizes)-1])
	s1, s2 := float64(stepCounts[0]), float64(stepCounts[len(stepCounts)-1])

	if s1 == 0 || n1 == 0 {
		return OUnknown
	}

	ratio := s2 / s1
	nRatio := n2 / n1

	// Check against expected ratios
	o1Ratio := 1.0
	onRatio := nRatio
	on2Ratio := nRatio * nRatio
	ologNRatio := math.Log2(n2) / math.Log2(n1)
	onlogNRatio := (n2 * math.Log2(n2)) / (n1 * math.Log2(n1))

	// Find closest match
	diffs := map[ComplexityClass]float64{
		O1:     math.Abs(ratio - o1Ratio),
		ON:     math.Abs(ratio - onRatio),
		ON2:    math.Abs(ratio - on2Ratio),
		OLogN:  math.Abs(ratio - ologNRatio),
		ONLogN: math.Abs(ratio - onlogNRatio),
	}

	minDiff := math.MaxFloat64
	result := OUnknown

	for class, diff := range diffs {
		if diff < minDiff {
			minDiff = diff
			result = class
		}
	}

	return result
}
