# Copilot instructions for this repository

This is a minimal Go repository with a single executable entrypoint at `main.go`.
These instructions help an AI coding agent quickly become productive and make safe, repository-consistent changes.

## Big picture
- **Repo type:** single-package Go CLI/executable. All logic should live under `main` or a small `pkg/` package if the project grows.
- **Primary intent:** small utility or starter; ask the maintainer before adding non-trivial features.

## Developer workflows (commands)
- Format & vet: `gofmt -w .` and `go vet ./...`
- Build: `go build -o bin/app .` or `go build ./...`
- Run: `go run ./` or `go run main.go`
- Test: `go test ./...` (no tests exist yet; add under `*_test.go` when required)

## Project-specific conventions
- Keep changes minimal and idiomatic Go. Favor small functions and packages.
- If adding features, request the expected CLI flags, environment variables, input/output format, and whether to use modules (`go.mod`).
- Place library code in a `pkg/` package if it will be reused; keep `main.go` as the thin entrypoint.

## Integration points & dependencies
- No external services or deps detected. If adding third-party modules, run `go get` and update `go.mod` and `go.sum`.

## PR & change guidance
- For small fixes (formatting, tiny helpers) create the change directly and include a short PR description with example run commands.
- For any design or input-format decisions (CLI args, config, APIs), open an issue or ask the user before implementing.

## Examples from this codebase
- Entrypoint: `main.go` — assume CLI behavior is unspecified; ask before implementing behavior beyond printing a TODO message.

If anything here is unclear or you want the agent to follow stricter conventions (packaging, testing, CI), tell me which areas to expand.
# Copilot instructions for this repo

This repository currently contains a single `main.go` at the project root (empty at time of inspection). The guidance below is tuned to help an AI coding agent be immediately productive in this workspace.

1. Big picture
- Repo type: minimal Go project (presence of `main.go`). No other source, configs, or README were found.
- Primary goal for agents: ask the user for the intended executable behavior and required inputs before implementing non-trivial logic.

2. What to look for / confirm before coding
- If implementing features, request: expected CLI args, environment variables, input/output format, and any backend/service contracts.
- Ask whether we should follow modules (Go modules) or a single-file script approach.

3. Build / test / debug commands (use these when adding code)
- Format and vet: `gofmt -w .` and `go vet ./...`
- Build: `go build ./...` or `go build -o bin/app .`
- Test: `go test ./...`
- Run (for `main`): `go run ./` or `go run main.go`

4. Project-specific conventions (discoverable here)
- Keep code minimal and idiomatic Go: small packages, exported names only when required by other packages.
- Place executable entry in `main.go` and library code in a `pkg` or top-level package when repo grows.
- Use module-aware commands; if `go.mod` is added later, prefer `go test ./...` and `go build ./...`.

5. Integration points and dependencies
- No external dependencies detected. If adding third-party modules, update `go.mod` with `go get` and ensure `go.sum` is committed.

6. When to open a PR vs. ask the user
- For any implementation that requires design decisions (input formats, error-handling semantics, CLI flags), open an issue or ask the user first.
- For small, well-scoped tasks (formatting, small helper funcs) create a PR with a short description and example usage.

7. Examples / templates
- Minimal `main` scaffold to suggest when user asks for a starter:

```go
package main

import "fmt"

func main() {
    fmt.Println("TODO: implement")
}
```

8. Notes for reviewers and future agents
- This repo lacks tests and documentation; prioritize clarifying requirements before adding complex code.
- If the user wants a specific build/test matrix or CI, ask for target platforms and Go version.

If any section is unclear or you want the instructions adjusted (e.g., prefer workspace-level CI templates or a different code layout), tell me what to change and I'll update this file.
