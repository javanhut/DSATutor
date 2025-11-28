package sandbox

import (
	"fmt"
	"regexp"
	"strings"
)

// BlockedImports contains Python modules that are not allowed.
var BlockedImports = []string{
	// System access
	"os",
	"sys",
	"subprocess",
	"shutil",
	"pathlib",
	"glob",

	// File operations
	"io",
	"tempfile",

	// Network
	"socket",
	"urllib",
	"requests",
	"http",
	"ftplib",
	"smtplib",
	"telnetlib",
	"ssl",

	// Code execution
	"code",
	"codeop",
	"compileall",
	"py_compile",

	// Serialization (can execute code)
	"pickle",
	"marshal",
	"shelve",

	// Low-level
	"ctypes",
	"cffi",
	"_thread",
	"threading",
	"multiprocessing",

	// Introspection
	"inspect",
	"dis",
	"gc",

	// Import manipulation
	"importlib",
	"pkgutil",
	"modulefinder",

	// Dangerous
	"pty",
	"tty",
	"termios",
	"resource",
	"signal",
	"fcntl",
	"mmap",

	// Async (can be exploited)
	"asyncio",
	"concurrent",
}

// BlockedBuiltins contains Python builtins that are not allowed.
// Note: __import__ is allowed because tracer.py provides a restricted version
// that only permits importing from safe modules.
var BlockedBuiltins = []string{
	"open",
	"exec",
	"eval",
	"compile",
	"input",
	"breakpoint",
	"globals",
	"locals",
	"vars",
	"dir",
	"getattr",
	"setattr",
	"delattr",
	"hasattr",
}

// BlockedPatterns contains regex patterns for dangerous code.
var BlockedPatterns = []*regexp.Regexp{
	// Attribute access to dunder methods
	regexp.MustCompile(`\.__class__`),
	regexp.MustCompile(`\.__bases__`),
	regexp.MustCompile(`\.__subclasses__`),
	regexp.MustCompile(`\.__mro__`),
	regexp.MustCompile(`\.__globals__`),
	regexp.MustCompile(`\.__code__`),
	regexp.MustCompile(`\.__builtins__`),
	regexp.MustCompile(`\.__import__`),
	regexp.MustCompile(`\.__dict__`),
	regexp.MustCompile(`\.__module__`),

	// File operations via alternative methods
	regexp.MustCompile(`Path\s*\(`),
	regexp.MustCompile(`file\s*=`),
}

// ValidateCode checks if Python code is safe to execute.
func ValidateCode(code string) error {
	// Normalize line endings
	code = strings.ReplaceAll(code, "\r\n", "\n")
	code = strings.ReplaceAll(code, "\r", "\n")

	lines := strings.Split(code, "\n")

	for lineNum, line := range lines {
		// Skip comments and empty lines
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Check for blocked imports
		if err := checkImports(line, lineNum+1); err != nil {
			return err
		}

		// Check for blocked builtins
		if err := checkBuiltins(line, lineNum+1); err != nil {
			return err
		}

		// Check for blocked patterns
		if err := checkPatterns(line, lineNum+1); err != nil {
			return err
		}
	}

	return nil
}

// checkImports checks for blocked import statements.
func checkImports(line string, lineNum int) error {
	trimmed := strings.TrimSpace(line)

	// Match import statements
	// import os
	// import os, sys
	// from os import path
	// from os.path import join
	// import os as o

	isImport := strings.HasPrefix(trimmed, "import ") || strings.HasPrefix(trimmed, "from ")
	if !isImport {
		return nil
	}

	for _, blocked := range BlockedImports {
		// Check various import patterns
		patterns := []string{
			fmt.Sprintf("import %s", blocked),
			fmt.Sprintf("import %s,", blocked),
			fmt.Sprintf("import %s ", blocked),
			fmt.Sprintf("from %s", blocked),
			fmt.Sprintf(", %s", blocked),
			fmt.Sprintf(",%s", blocked),
		}

		for _, pattern := range patterns {
			if strings.Contains(trimmed, pattern) {
				return fmt.Errorf("line %d: import '%s' is not allowed for security reasons", lineNum, blocked)
			}
		}
	}

	return nil
}

// checkBuiltins checks for blocked builtin function calls.
func checkBuiltins(line string, lineNum int) error {
	for _, blocked := range BlockedBuiltins {
		// Match function calls: builtin(
		pattern := fmt.Sprintf(`\b%s\s*\(`, regexp.QuoteMeta(blocked))
		re := regexp.MustCompile(pattern)

		if re.MatchString(line) {
			return fmt.Errorf("line %d: builtin '%s()' is not allowed for security reasons", lineNum, blocked)
		}
	}

	return nil
}

// checkPatterns checks for blocked code patterns.
func checkPatterns(line string, lineNum int) error {
	for _, pattern := range BlockedPatterns {
		if pattern.MatchString(line) {
			return fmt.Errorf("line %d: code pattern not allowed for security reasons", lineNum)
		}
	}

	return nil
}
