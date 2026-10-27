# Custom problems

Drop extra problems here as `*.json` files, one problem per file, using the
JSON form of `practice.Problem` (see `internal/practice/problem.go`). They are
embedded into the binary at build time and loaded after the built-in problems;
a JSON problem whose `id` matches a built-in one is ignored.

This README also keeps the directory present, which `embed.go` needs to build.
