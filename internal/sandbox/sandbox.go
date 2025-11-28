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

//go:embed tracer.py dsa_helpers.py
var sandboxFiles embed.FS

// ExecuteRequest represents a request to execute Python code.
type ExecuteRequest struct {
	Code    string `json:"code"`
	Timeout int    `json:"timeout"` // milliseconds, max 10000
}

// ExecuteResponse represents the result of code execution.
type ExecuteResponse struct {
	Success bool        `json:"success"`
	Steps   []Step      `json:"steps"`
	Output  string      `json:"output"`
	Error   string      `json:"error,omitempty"`
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

	// Validate code security
	if err := ValidateCode(req.Code); err != nil {
		return &ExecuteResponse{
			Success: false,
			Steps:   []Step{},
			Output:  "",
			Error:   err.Error(),
		}, nil
	}

	// Set timeout
	timeout := req.Timeout
	if timeout <= 0 || timeout > MaxTimeout {
		timeout = DefaultTimeout
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

	cmd := exec.CommandContext(ctx, "python3", tracerPath, codePath)
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
