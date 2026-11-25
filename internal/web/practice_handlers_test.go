package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"dsatutor/internal/practice"
)

func practiceTestServer(t *testing.T) *http.ServeMux {
	t.Helper()
	loader := practice.NewProblemLoader(practice.ProblemsFS, practice.ProblemsBasePath)
	if err := loader.Load(); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	NewPracticeServer(loader).RegisterRoutes(mux)
	return mux
}

func getJSON(t *testing.T, mux *http.ServeMux, path string, out interface{}) {
	t.Helper()
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("GET %s = %d", path, rec.Code)
	}
	if err := json.Unmarshal(rec.Body.Bytes(), out); err != nil {
		t.Fatal(err)
	}
}

// The UI keeps drills out of mixed practice, interviews and the solved count
// by their kind, so both the list and the single-problem response carry it,
// and the ladder total doesn't count them.
func TestDrillsCarryTheirKind(t *testing.T) {
	mux := practiceTestServer(t)

	var list struct {
		Problems []struct{ ID, Kind, Category string }
		Stats    map[string]int
	}
	getJSON(t, mux, "/api/practice/problems", &list)
	drills, firstTree := 0, ""
	for _, p := range list.Problems {
		if p.Kind == practice.KindDrill {
			drills++
		}
		if p.Category == "trees" && firstTree == "" {
			firstTree = p.ID
		}
	}
	if drills == 0 || list.Stats["drills"] != drills {
		t.Errorf("list has %d drills, stats say %d", drills, list.Stats["drills"])
	}
	if list.Stats["total"] != len(list.Problems)-drills {
		t.Errorf("stats total %d should count only ladder problems (%d)", list.Stats["total"], len(list.Problems)-drills)
	}
	if firstTree != "drill-build-binary-tree" {
		t.Errorf("trees should start with its drills, starts with %s", firstTree)
	}

	var one struct{ Kind string }
	getJSON(t, mux, "/api/practice/problems/drill-build-binary-tree", &one)
	if one.Kind != practice.KindDrill {
		t.Errorf("single-problem response has kind %q", one.Kind)
	}
}
