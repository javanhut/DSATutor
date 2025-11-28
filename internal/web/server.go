package web

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"time"

	"dsatutor/internal/chapter"
	"dsatutor/internal/sandbox"
)

//go:embed static/*
var staticFiles embed.FS
var staticSub, _ = fs.Sub(staticFiles, "static")

// Server hosts a minimal web UI to browse chapters and play storyboards.
type Server struct {
	chapters []chapter.Chapter
	mux      *http.ServeMux
}

// NewServer wires handlers for static content and chapter data.
func NewServer(chapters []chapter.Chapter) *Server {
	s := &Server{
		chapters: chapters,
		mux:      http.NewServeMux(),
	}
	s.routes()
	return s
}

// ServeHTTP satisfies http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// Listen starts the HTTP server.
func (s *Server) Listen(addr string) error {
	server := &http.Server{
		Addr:              addr,
		Handler:           s.mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Web UI listening on %s\n", addr)
	return server.ListenAndServe()
}

func (s *Server) routes() {
	s.mux.HandleFunc("/api/chapters", s.handleChapters)
	s.mux.HandleFunc("/api/sandbox/execute", s.handleSandboxExecute)
	s.mux.HandleFunc("/", s.handleIndex)
	s.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticSub))))
}

func (s *Server) handleChapters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(s.chapters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	data, err := staticFiles.ReadFile("static/index.html")
	if err != nil {
		http.Error(w, "index not found", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func (s *Server) handleSandboxExecute(w http.ResponseWriter, r *http.Request) {
	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req sandbox.ExecuteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(sandbox.ExecuteResponse{
			Success: false,
			Steps:   []sandbox.Step{},
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}

	// Execute code
	response, err := sandbox.Execute(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(sandbox.ExecuteResponse{
			Success: false,
			Steps:   []sandbox.Step{},
			Error:   "Execution error: " + err.Error(),
		})
		return
	}

	// Return result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
