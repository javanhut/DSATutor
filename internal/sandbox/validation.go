package sandbox

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// ValidationResult represents the result of validating user output against expected.
type ValidationResult struct {
	Passed   bool   `json:"passed"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
	Message  string `json:"message,omitempty"`
}

// ValidateOutput compares user output against expected output.
// It handles various output formats: arrays, numbers, strings, booleans.
func ValidateOutput(actual, expected string) ValidationResult {
	actual = strings.TrimSpace(actual)
	expected = strings.TrimSpace(expected)

	// Direct string match
	if actual == expected {
		return ValidationResult{
			Passed:   true,
			Expected: expected,
			Actual:   actual,
		}
	}

	// Try parsing as JSON for structural comparison
	var actualVal, expectedVal interface{}
	actualErr := json.Unmarshal([]byte(actual), &actualVal)
	expectedErr := json.Unmarshal([]byte(expected), &expectedVal)

	if actualErr == nil && expectedErr == nil {
		if deepEqual(actualVal, expectedVal) {
			return ValidationResult{
				Passed:   true,
				Expected: expected,
				Actual:   actual,
			}
		}

		// Check for order-independent array comparison (for problems where order doesn't matter)
		if isUnorderedArrayMatch(actualVal, expectedVal) {
			return ValidationResult{
				Passed:   true,
				Expected: expected,
				Actual:   actual,
				Message:  "Arrays match (order independent)",
			}
		}
	}

	// Normalize Python representations
	normalizedActual := normalizePythonOutput(actual)
	normalizedExpected := normalizePythonOutput(expected)

	if normalizedActual == normalizedExpected {
		return ValidationResult{
			Passed:   true,
			Expected: expected,
			Actual:   actual,
		}
	}

	return ValidationResult{
		Passed:   false,
		Expected: expected,
		Actual:   actual,
		Message:  "Output does not match expected",
	}
}

// normalizePythonOutput converts Python-style output to normalized form.
func normalizePythonOutput(s string) string {
	// Replace Python True/False with JSON true/false
	s = strings.ReplaceAll(s, "True", "true")
	s = strings.ReplaceAll(s, "False", "false")
	s = strings.ReplaceAll(s, "None", "null")

	// Replace single quotes with double quotes for arrays/dicts
	s = strings.ReplaceAll(s, "'", "\"")

	// Remove whitespace around brackets and commas
	s = regexp.MustCompile(`\s*([[\]{}:,])\s*`).ReplaceAllString(s, "$1")

	// Handle tuples - convert (1, 2) to [1, 2]
	s = regexp.MustCompile(`\(([^()]*)\)`).ReplaceAllString(s, "[$1]")

	return s
}

// deepEqual performs deep equality check for JSON-parsed values.
func deepEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// isUnorderedArrayMatch checks if two values are arrays with same elements (order independent).
func isUnorderedArrayMatch(a, b interface{}) bool {
	aArr, aOk := a.([]interface{})
	bArr, bOk := b.([]interface{})

	if !aOk || !bOk || len(aArr) != len(bArr) {
		return false
	}

	// For simple arrays of primitives, check if they contain same elements
	aCounts := make(map[string]int)
	bCounts := make(map[string]int)

	for _, v := range aArr {
		key := fmt.Sprintf("%v", v)
		aCounts[key]++
	}

	for _, v := range bArr {
		key := fmt.Sprintf("%v", v)
		bCounts[key]++
	}

	return reflect.DeepEqual(aCounts, bCounts)
}

// ParseTestInput parses a test input string into function arguments.
// Handles formats like: "[1,2,3], 5" or "nums = [1,2,3], target = 5"
func ParseTestInput(input string) ([]string, error) {
	input = strings.TrimSpace(input)

	// Remove parameter names if present (e.g., "nums = [1,2,3]" -> "[1,2,3]")
	paramPattern := regexp.MustCompile(`\w+\s*=\s*`)
	input = paramPattern.ReplaceAllString(input, "")

	// Split by comma, but respect brackets and quotes
	args := splitArguments(input)

	return args, nil
}

// splitArguments splits a comma-separated argument list respecting brackets and quotes.
func splitArguments(input string) []string {
	var args []string
	var current strings.Builder
	depth := 0
	inString := false
	stringChar := rune(0)

	for i, ch := range input {
		switch {
		case (ch == '"' || ch == '\'') && !inString:
			inString = true
			stringChar = ch
			current.WriteRune(ch)
		case ch == stringChar && inString:
			// Check if escaped
			if i > 0 && input[i-1] == '\\' {
				current.WriteRune(ch)
			} else {
				inString = false
				stringChar = 0
				current.WriteRune(ch)
			}
		case inString:
			current.WriteRune(ch)
		case ch == '[' || ch == '(' || ch == '{':
			depth++
			current.WriteRune(ch)
		case ch == ']' || ch == ')' || ch == '}':
			depth--
			current.WriteRune(ch)
		case ch == ',' && depth == 0:
			args = append(args, strings.TrimSpace(current.String()))
			current.Reset()
		default:
			current.WriteRune(ch)
		}
	}

	// Add last argument
	if current.Len() > 0 {
		args = append(args, strings.TrimSpace(current.String()))
	}

	return args
}

// FormatTestArgs formats parsed arguments for Python function call.
func FormatTestArgs(args []string) string {
	return strings.Join(args, ", ")
}

// ExtractReturnValue extracts the return value from execution output.
// Handles cases where output might have multiple lines.
func ExtractReturnValue(output string) string {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) == 0 {
		return ""
	}
	// Return the last non-empty line as the result
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			return line
		}
	}
	return ""
}

// CompareNumeric compares two numeric values with tolerance for floating point.
func CompareNumeric(actual, expected string, tolerance float64) bool {
	actualNum, err1 := strconv.ParseFloat(actual, 64)
	expectedNum, err2 := strconv.ParseFloat(expected, 64)

	if err1 != nil || err2 != nil {
		return actual == expected
	}

	diff := actualNum - expectedNum
	if diff < 0 {
		diff = -diff
	}

	return diff <= tolerance
}
