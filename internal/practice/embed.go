package practice

import (
	"embed"
)

//go:embed all:problems
var ProblemsFS embed.FS

// ProblemsBasePath is the base path within the embedded filesystem.
const ProblemsBasePath = "problems"
