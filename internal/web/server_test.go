package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// The "/" handler must serve the index only at exactly "/", not act as a
// catch-all: unknown paths (e.g. an API route from a newer client against an
// older server) need an honest 404, not the index page with a 200.
func TestUnknownPathsReturn404(t *testing.T) {
	s := NewServer(nil)
	for _, path := range []string{"/no/such/page", "/api/nope", "/index.html"} {
		rec := httptest.NewRecorder()
		s.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
		if rec.Code != http.StatusNotFound {
			t.Errorf("GET %s = %d, want 404", path, rec.Code)
		}
	}
}

func TestIndexServedAtRoot(t *testing.T) {
	s := NewServer(nil)
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("GET / = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "<!doctype html>") {
		t.Error("GET / did not return the index page")
	}
}
