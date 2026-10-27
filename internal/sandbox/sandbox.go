// Package sandbox provides secure Python code execution with tracing
// for DSA visualization.
package sandbox

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

//go:embed tracer.py harness.py dsa_helpers.py
var sandboxFiles embed.FS

// ExecuteRequest represents a request to execute Python code.
type ExecuteRequest struct {
	Code    string `json:"code"`
	Timeout int    `json:"timeout"` // milliseconds, max 10000

	// NoTrace skips line-by-line recording. Judging a submission doesn't need
	// the visualization, and tracing slows long-running solutions a lot.
	NoTrace bool `json:"noTrace,omitempty"`

	// Harness, when set, tells the sandbox how to call the user's code for one
	// test case and report its return value (see harness.py).
	Harness *HarnessSpec `json:"harness,omitempty"`

	// Language is "python" (the default) or "oxigen" (see oxigen.go).
	Language string `json:"language,omitempty"`
}

// HarnessSpec describes how to run one test case against user code.
type HarnessSpec struct {
	Mode       string                 `json:"mode"`                 // function, design, roundtrip
	Function   string                 `json:"function,omitempty"`   // function name, or "enc,dec" for roundtrip
	ClassName  string                 `json:"className,omitempty"`  // roundtrip owner class (e.g. Codec)
	Params     []string               `json:"params,omitempty"`     // argument order
	Inputs     map[string]interface{} `json:"inputs"`               // raw JSON inputs keyed by parameter
	ParamTypes map[string]string      `json:"paramTypes,omitempty"` // ListNode, TreeNode, GraphNode, ...
	ReturnType string                 `json:"returnType,omitempty"` // GraphNode, RandomList, TreeNodeVal
	InPlace    string                 `json:"inPlace,omitempty"`    // parameter whose final value is the answer
	Prelude    string                 `json:"prelude,omitempty"`    // trusted definitions run before user code
	Validator  string                 `json:"validator,omitempty"`  // trusted check(raw, args, result, expected)
	Expected   interface{}            `json:"expected,omitempty"`
}

// Verdict is a validator's decision on a returned answer.
type Verdict struct {
	Passed  bool   `json:"passed"`
	Message string `json:"message"`
}

// ExecuteResponse represents the result of code execution.
type ExecuteResponse struct {
	Success   bool        `json:"success"`
	Steps     []Step      `json:"steps"`
	Output    string      `json:"output"`
	Error     string      `json:"error,omitempty"`
	Truncated bool        `json:"truncated,omitempty"` // more steps ran than were recorded
	HasResult bool        `json:"hasResult,omitempty"`
	Result    interface{} `json:"result,omitempty"` // harness return value as JSON
	Verdict   *Verdict    `json:"verdict,omitempty"`
}

// Step represents a single execution step captured by the tracer.
type Step struct {
	LineNum    int                    `json:"lineNum"`
	Function   string                 `json:"function"`
	State      string                 `json:"state"`
	Locals     map[string]interface{} `json:"locals"`
	Structures []Structure            `json:"structures"`
}

// Structure represents a detected data structure for visualization.
type Structure struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Data       interface{} `json:"data"`
	Highlights interface{} `json:"highlights,omitempty"`
}

// DefaultTimeout is the default execution timeout in milliseconds.
const DefaultTimeout = 10000

// MaxTimeout is the maximum allowed timeout in milliseconds.
const MaxTimeout = 10000

// MaxCodeSize is the maximum allowed code size in bytes.
const MaxCodeSize = 10240 // 10KB

// Execute runs Python code with tracing and returns the execution trace.
func Execute(req ExecuteRequest) (*ExecuteResponse, error) {
	// Validate request
	if len(req.Code) == 0 {
		return &ExecuteResponse{
			Success: false,
			Steps:   []Step{},
			Output:  "",
			Error:   "No code provided",
		}, nil
	}

	if len(req.Code) > MaxCodeSize {
		return &ExecuteResponse{
			Success: false,
			Steps:   []Step{},
			Output:  "",
			Error:   fmt.Sprintf("Code exceeds maximum size of %d bytes", MaxCodeSize),
		}, nil
	}

	// Set timeout
	timeout := req.Timeout
	if timeout <= 0 || timeout > MaxTimeout {
		timeout = DefaultTimeout
	}

	switch req.Language {
	case "", "python":
	case "oxigen":
		return executeOxigen(req, timeout)
	default:
		return failed(fmt.Sprintf("unknown language %q", req.Language)), nil
	}

	// Validate code security
	if err := ValidateCode(req.Code); err != nil {
		return &ExecuteResponse{
			Success: false,
			Steps:   []Step{},
			Output:  "",
			Error:   err.Error(),
		}, nil
	}

	// Create temp directory for execution
	tmpDir, err := os.MkdirTemp("", "dsatutor-sandbox-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	// Write tracer script
	tracerPath := filepath.Join(tmpDir, "tracer.py")
	tracerContent, err := sandboxFiles.ReadFile("tracer.py")
	if err != nil {
		return nil, fmt.Errorf("failed to read tracer script: %w", err)
	}
	if err := os.WriteFile(tracerPath, tracerContent, 0600); err != nil {
		return nil, fmt.Errorf("failed to write tracer script: %w", err)
	}

	// Write test harness
	harnessPath := filepath.Join(tmpDir, "harness.py")
	harnessContent, err := sandboxFiles.ReadFile("harness.py")
	if err != nil {
		return nil, fmt.Errorf("failed to read harness script: %w", err)
	}
	if err := os.WriteFile(harnessPath, harnessContent, 0600); err != nil {
		return nil, fmt.Errorf("failed to write harness script: %w", err)
	}

	// Write DSA helpers
	helpersPath := filepath.Join(tmpDir, "dsa_helpers.py")
	helpersContent, err := sandboxFiles.ReadFile("dsa_helpers.py")
	if err != nil {
		return nil, fmt.Errorf("failed to read helpers script: %w", err)
	}
	if err := os.WriteFile(helpersPath, helpersContent, 0600); err != nil {
		return nil, fmt.Errorf("failed to write helpers script: %w", err)
	}

	// Write user code
	codePath := filepath.Join(tmpDir, "user_code.py")
	if err := os.WriteFile(codePath, []byte(req.Code), 0600); err != nil {
		return nil, fmt.Errorf("failed to write user code: %w", err)
	}

	// Execute with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	defer cancel()

	args := []string{tracerPath, codePath}
	if req.Harness != nil {
		specPath := filepath.Join(tmpDir, "harness.json")
		spec, err := json.Marshal(req.Harness)
		if err != nil {
			return nil, fmt.Errorf("failed to encode harness spec: %w", err)
		}
		if err := os.WriteFile(specPath, spec, 0600); err != nil {
			return nil, fmt.Errorf("failed to write harness spec: %w", err)
		}
		args = append(args, specPath)
	}
	if req.NoTrace {
		args = append(args, "--no-trace")
	}

	cmd := exec.CommandContext(ctx, "python3", args...)
	cmd.Dir = tmpDir

	output, err := cmd.Output()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return &ExecuteResponse{
				Success: false,
				Steps:   []Step{},
				Output:  "",
				Error:   fmt.Sprintf("Execution timed out after %d ms", timeout),
			}, nil
		}

		// Try to get stderr for more info
		if exitErr, ok := err.(*exec.ExitError); ok {
			return &ExecuteResponse{
				Success: false,
				Steps:   []Step{},
				Output:  "",
				Error:   fmt.Sprintf("Execution failed: %s", string(exitErr.Stderr)),
			}, nil
		}

		return &ExecuteResponse{
			Success: false,
			Steps:   []Step{},
			Output:  "",
			Error:   fmt.Sprintf("Execution failed: %v", err),
		}, nil
	}

	// Parse JSON output from tracer
	var response ExecuteResponse
	if err := json.Unmarshal(output, &response); err != nil {
		return &ExecuteResponse{
			Success: false,
			Steps:   []Step{},
			Output:  string(output),
			Error:   fmt.Sprintf("Failed to parse tracer output: %v", err),
		}, nil
	}

	return &response, nil
}
