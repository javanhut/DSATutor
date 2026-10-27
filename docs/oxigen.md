# Oxigen

Practice problems and the sandbox accept [OxigenLang](https://www.oxigenlang.org)
as well as Python. Pick the language in the problem's code panel or the
sandbox header; the choice is remembered in localStorage under
`dsatutor_language`.

`imlazy setup` installs the `oxigen` binary with rvn, and the Docker image
copies it from `javanhut/oxigen` (see [docker.md](docker.md)). If the server
can't find it, Oxigen runs fail with a message saying so and Python keeps
working.

## What works, and what doesn't

| | Python | Oxigen |
|---|---|---|
| Run a visible test case | yes | yes |
| Submit against every case, hidden ones included | yes | yes |
| Line-by-line visualization of a run | yes | no: Oxigen has no tracer, so runs show the answer and printed output |
| Starter code | written per problem | generated from the Python starter |
| Hints, solution walkthrough, chapter code | Python | Python |
| Sandbox | traced, structures drawn | output only |

## Writing a solution

The starter keeps every name from the Python starter, because the judge calls
the code by those names.

- **Functions** become `fun name(params) { ... }`. Return with `give`.
- **Node types** the problem uses are added to your file by the judge, after
  your code: `ListNode(val, next)`, `TreeNode(val, left, right)`, and
  `Node(val, neighbors)` or `Node(val, next, random)` for the clone problems.
  Missing fields start as `None`, so `ListNode(5)` works. Define one yourself
  and yours is used instead.
- **Design problems** (Min Stack, LRU Cache, ...) become a struct with an
  `includes` block. `__init__` becomes `init`: the judge declares a
  zero-value instance and calls `init(...)` with the constructor arguments,
  then each operation in turn. Oxigen functions return their last expression,
  so an operation whose expected answer is `null` is treated as returning
  nothing.
- **In-place problems** (Rotate Image, Reorder List, ...) are judged on the
  final value of the named parameter, as in Python.
- `==` compares structs field by field, not by identity.

## How a test case is judged

`internal/sandbox/oxigen.go` writes three files to a temporary directory:

- `solution.oxi`: your code followed by the node structs you didn't define.
- `input.json`: the test case's inputs.
- `harness.oxi`: `harness_lib.oxi` (conversions between JSON and nodes,
  written in Oxigen) and a generated `main` block that imports your module,
  builds the arguments, makes the call, and prints the answer as JSON on a
  marked line.

It then runs `oxigen harness.oxi`. Your module can't see the harness's
imports, because Oxigen modules keep their own globals. The judge compares
the answer with the same `Compare` function it uses for Python. Problems with
a Python `Validator` have it run by `validate.py` on the JSON answer, with the
raw input as `args`. The two clone problems instead have the harness check that
the copy doesn't reuse the original's nodes. Oxigen has no `id()`, so it marks
each original node's `val` and looks for the marks in the copy.

## Security

`ValidateOxigenCode` refuses:

- `introduce`/`intro` of anything except `array`, `datetime`, `encoding`,
  `hash`, `math`, `random`, `regex`, `result`, `strings` and `time`. The
  others reach files, processes, the network, environment or C.
- Local imports (`introduce .file`).
- Any name containing `__`: Oxigen's native functions (`__exec`,
  `__read_file`, ...) are globals in every module.

Runs share the Python sandbox's limits: 10 KB of code and a 10-second timeout.

## Tests

`go test ./internal/judge/ -run Oxigen` judges Oxigen solutions to 16
problems that cover every harness feature, checks that wrong answers and
node-reusing clones fail, and runs `oxigen check` on every generated
starter. The tests skip when `oxigen` isn't installed.
