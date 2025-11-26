package storage

import (
	"encoding/json"
	"os"

	"dsatutor/internal/chapter"
)

// SaveJSON writes chapters to a file so they can be edited without recompiling.
func SaveJSON(path string, chapters []chapter.Chapter) error {
	data, err := json.MarshalIndent(chapters, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// LoadJSON loads chapters from a JSON file, allowing runtime customization.
func LoadJSON(path string) ([]chapter.Chapter, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var chapters []chapter.Chapter
	if err := json.Unmarshal(raw, &chapters); err != nil {
		return nil, err
	}
	return chapters, nil
}
