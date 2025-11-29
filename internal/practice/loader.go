package practice

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"sort"
	"strings"
)

// ProblemLoader handles loading problems from embedded Go code and optional JSON files.
type ProblemLoader struct {
	problems   map[string]*Problem
	categories []Category
	embedFS    embed.FS
	basePath   string
}

// NewProblemLoader creates a new loader.
func NewProblemLoader(embedFS embed.FS, basePath string) *ProblemLoader {
	return &ProblemLoader{
		problems:   make(map[string]*Problem),
		categories: Blind75Categories,
		embedFS:    embedFS,
		basePath:   basePath,
	}
}

// Load loads all problems - first from embedded Go code, then from any JSON files.
func (pl *ProblemLoader) Load() error {
	// Step 1: Load embedded problems from Go code
	for _, p := range EmbeddedProblems {
		pl.problems[p.ID] = p
	}
	log.Printf("Loaded %d embedded problems", len(EmbeddedProblems))

	// Step 2: Try to load additional problems from JSON files (optional)
	pl.loadJSONProblems()

	log.Printf("Total problems loaded: %d", len(pl.problems))
	return nil
}

// loadJSONProblems attempts to load additional problems from JSON files.
// These are optional and can extend the embedded problems.
func (pl *ProblemLoader) loadJSONProblems() {
	// Try to find and load custom problems from the embedded filesystem
	err := fs.WalkDir(pl.embedFS, pl.basePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil // Skip errors
		}
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		// Skip index.json and categories.json
		if strings.HasSuffix(path, "index.json") || strings.HasSuffix(path, "categories.json") {
			return nil
		}

		data, err := pl.embedFS.ReadFile(path)
		if err != nil {
			log.Printf("Warning: Could not read %s: %v", path, err)
			return nil
		}

		var problem Problem
		if err := json.Unmarshal(data, &problem); err != nil {
			log.Printf("Warning: Could not parse %s: %v", path, err)
			return nil
		}

		// Only add if not already in embedded problems (JSON can override)
		if _, exists := pl.problems[problem.ID]; !exists {
			pl.problems[problem.ID] = &problem
			log.Printf("Loaded additional problem from JSON: %s", problem.ID)
		}

		return nil
	})

	if err != nil {
		log.Printf("Warning: Error walking problem directory: %v", err)
	}
}

// LoadFromDirectory loads additional problems from a directory (for external JSON files).
func (pl *ProblemLoader) LoadFromDirectory(dirPath string, fsys fs.FS) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		if strings.HasSuffix(path, "index.json") || strings.HasSuffix(path, "categories.json") {
			return nil
		}

		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return nil
		}

		var problem Problem
		if err := json.Unmarshal(data, &problem); err != nil {
			log.Printf("Warning: Could not parse %s: %v", path, err)
			return nil
		}

		pl.problems[problem.ID] = &problem
		return nil
	})
}

// GetProblem returns a problem by ID.
func (pl *ProblemLoader) GetProblem(id string) (*Problem, error) {
	problem, ok := pl.problems[id]
	if !ok {
		return nil, fmt.Errorf("problem not found: %s", id)
	}
	return problem, nil
}

// GetAllProblems returns all loaded problems.
func (pl *ProblemLoader) GetAllProblems() []*Problem {
	problems := make([]*Problem, 0, len(pl.problems))
	for _, p := range pl.problems {
		problems = append(problems, p)
	}

	// Sort by number
	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})

	return problems
}

// GetProblemsByCategory returns problems in a specific category.
func (pl *ProblemLoader) GetProblemsByCategory(categoryID string) []*Problem {
	problems := make([]*Problem, 0)
	for _, p := range pl.problems {
		if p.Category == categoryID {
			problems = append(problems, p)
		}
	}

	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})

	return problems
}

// GetProblemsByDifficulty returns problems of a specific difficulty.
func (pl *ProblemLoader) GetProblemsByDifficulty(difficulty string) []*Problem {
	problems := make([]*Problem, 0)
	for _, p := range pl.problems {
		if strings.EqualFold(p.Difficulty, difficulty) {
			problems = append(problems, p)
		}
	}

	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})

	return problems
}

// GetProblemsByTag returns problems with a specific tag.
func (pl *ProblemLoader) GetProblemsByTag(tag string) []*Problem {
	problems := make([]*Problem, 0)
	for _, p := range pl.problems {
		for _, t := range p.Tags {
			if strings.EqualFold(t, tag) {
				problems = append(problems, p)
				break
			}
		}
	}

	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})

	return problems
}

// GetProblemsForChapter returns problems related to a chapter.
func (pl *ProblemLoader) GetProblemsForChapter(chapterNum int) []*Problem {
	problems := make([]*Problem, 0)
	for _, p := range pl.problems {
		for _, ref := range p.RelatedChapters {
			if ref == chapterNum {
				problems = append(problems, p)
				break
			}
		}
	}

	sort.Slice(problems, func(i, j int) bool {
		return problems[i].Number < problems[j].Number
	})

	return problems
}

// SearchProblems searches problems by title or description.
func (pl *ProblemLoader) SearchProblems(query string) []*Problem {
	query = strings.ToLower(query)
	problems := make([]*Problem, 0)

	for _, p := range pl.problems {
		if strings.Contains(strings.ToLower(p.Title), query) ||
			strings.Contains(strings.ToLower(p.Description), query) {
			problems = append(problems, p)
		}
	}

	sort.Slice(problems, func(i, j int) bool {
		iTitle := strings.Contains(strings.ToLower(problems[i].Title), query)
		jTitle := strings.Contains(strings.ToLower(problems[j].Title), query)
		if iTitle != jTitle {
			return iTitle
		}
		return problems[i].Number < problems[j].Number
	})

	return problems
}

// GetCategories returns all categories.
func (pl *ProblemLoader) GetCategories() []Category {
	return pl.categories
}

// GetStats returns statistics about loaded problems.
func (pl *ProblemLoader) GetStats() map[string]int {
	stats := map[string]int{
		"total":  len(pl.problems),
		"easy":   0,
		"medium": 0,
		"hard":   0,
	}

	for _, p := range pl.problems {
		switch strings.ToLower(p.Difficulty) {
		case "easy":
			stats["easy"]++
		case "medium":
			stats["medium"]++
		case "hard":
			stats["hard"]++
		}
	}

	return stats
}

// FilterProblems applies multiple filters and returns matching problems.
func (pl *ProblemLoader) FilterProblems(category, difficulty, tag, search string) []*Problem {
	problems := pl.GetAllProblems()
	result := make([]*Problem, 0)

	for _, p := range problems {
		if category != "" && p.Category != category {
			continue
		}

		if difficulty != "" && !strings.EqualFold(p.Difficulty, difficulty) {
			continue
		}

		if tag != "" {
			hasTag := false
			for _, t := range p.Tags {
				if strings.EqualFold(t, tag) {
					hasTag = true
					break
				}
			}
			if !hasTag {
				continue
			}
		}

		if search != "" {
			searchLower := strings.ToLower(search)
			if !strings.Contains(strings.ToLower(p.Title), searchLower) &&
				!strings.Contains(strings.ToLower(p.Description), searchLower) {
				continue
			}
		}

		result = append(result, p)
	}

	return result
}
