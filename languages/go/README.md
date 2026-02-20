# Go Project Setup

Opinionated guide for structuring and tooling a production Go project.

## README Badges

Copy these into your project README, replacing the placeholders.

```markdown
<!-- TODO: customise -- replace yourorg/your-project and gist ID -->
[![CI](https://github.com/yourorg/your-project/actions/workflows/ci.yml/badge.svg)](https://github.com/yourorg/your-project/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/youruser/your-gist-id/raw/coverage.json)](https://github.com/yourorg/your-project/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
```

## Package Structure

```
<!-- TODO: customise the module name -->
your-project/
├── cmd/
│   └── your-app/          # main.go lives here, one dir per binary
│       └── main.go
├── internal/              # private application code, not importable by other modules
│   ├── config/            # environment/flag parsing (caarlos0/env, godotenv)
│   ├── api/               # HTTP handlers, middleware, routing
│   ├── store/             # database/storage layer
│   ├── models/            # domain types and business logic
│   ├── errors/            # custom error types
│   └── logging/           # structured logging setup (slog)
├── pkg/                   # public library code (only if you intend others to import it)
├── .github/workflows/     # CI/CD pipelines
├── scripts/pre-commit     # TruffleHog secret scanning hook
├── .golangci.yml
├── .goreleaser.yml
├── justfile
├── renovate.json
├── go.mod
├── .env.example
├── .dockerignore
├── .gitignore
├── CHANGELOG.md
├── CONTRIBUTING.md
└── SECURITY.md
```

Key rules:
- `cmd/` contains only the `main` package. Keep it thin. Wire up dependencies and call into `internal/`.
- `internal/` is where all application logic lives. Go enforces that nothing outside this module can import it.
- `pkg/` is optional. Only use it if you are building a library meant for external consumption. Most projects do not need it.

## Module Initialisation

```bash
# Initialise the module
# <!-- TODO: customise the module path -->
go mod init github.com/yourorg/your-project

# Create the entry point
mkdir -p cmd/your-app
touch cmd/your-app/main.go

# Create internal packages as needed
mkdir -p internal/{config,api,store,models,errors,logging}
```

## Recommended Tooling

### golangci-lint

The `.golangci.yml` in this directory is ready to copy. Install the linter:

```bash
# Install (see https://golangci-lint.run/welcome/install/)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### air (hot reload for development)

```bash
go install github.com/air-verse/air@latest

# Initialise config
air init

# Run (watches for file changes and rebuilds)
air
```

The default `.air.toml` works well. Adjust `cmd` if your binary lives somewhere other than `cmd/server`.

### gofumpt (stricter formatting)

```bash
go install mvdan.cc/gofumpt@latest
```

gofumpt is a superset of gofmt. It enforces additional formatting rules that the community has largely converged on. The provided `.golangci.yml` enables it as a formatter.

### testify (assertions)

```bash
go get github.com/stretchr/testify
```

Standard library testing is fine for simple cases. testify reduces boilerplate when you have many assertions or need mocking.

## Recommended Packages

### Config/Environment

| Package | Install | Purpose |
|---------|---------|---------|
| [caarlos0/env](https://github.com/caarlos0/env) | `go get github.com/caarlos0/env/v11` | Parse environment variables into structs using tags. No boilerplate. |
| [joho/godotenv](https://github.com/joho/godotenv) | `go get github.com/joho/godotenv` | Load `.env` files in development. Call `godotenv.Load()` before `env.Parse()`. |

### HTTP

| Package | Install | Purpose |
|---------|---------|---------|
| [gin-gonic/gin](https://github.com/gin-gonic/gin) | `go get github.com/gin-gonic/gin` | HTTP framework. Fast, middleware-friendly, good error handling. |

### Auth/OAuth

| Package | Install | Purpose |
|---------|---------|---------|
| [ory/fosite](https://github.com/ory/fosite) | `go get github.com/ory/fosite/v2` | OAuth 2.0 server framework. Supports authorization code, PKCE, client credentials, DCR (RFC 7591), token introspection, and revocation. Implement the storage interfaces against your backend. |

### Database

| Package | Install | Purpose |
|---------|---------|---------|
| [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) | `go get modernc.org/sqlite` | Pure Go SQLite driver. No CGO required, cross-compiles cleanly. |

### Frontend

| Package | Install | Purpose |
|---------|---------|---------|
| [a-h/templ](https://github.com/a-h/templ) | `go install github.com/a-h/templ/cmd/templ@latest` | Type-safe HTML templates written in Go. Compile `.templ` files to Go code. |
| [htmx](https://htmx.org) | CDN or `npm install htmx.org` | Add interactivity without a JS framework. Pairs well with templ for server-rendered UIs. |
| [air-verse/air](https://github.com/air-verse/air) | `go install github.com/air-verse/air@latest` | Hot reload on file changes. Configure `.air.toml` to run `templ generate` before rebuild. |

### Testing

| Package | Install | Purpose |
|---------|---------|---------|
| [stretchr/testify](https://github.com/stretchr/testify) | `go get github.com/stretchr/testify` | Assertions (`require`, `assert`) and test suites. |
| [uber-go/mock](https://github.com/uber-go/mock) | `go install go.uber.org/mock/mockgen@latest` | Generate interface mocks. Run `mockgen` to generate, import in tests. |

## Conventions

### Error handling
- Return errors, do not panic. Panics are reserved for truly unrecoverable programmer mistakes.
- Wrap errors with context using `fmt.Errorf("doing thing: %w", err)`.
- Define custom error types in `internal/errors/` when callers need to distinguish error kinds.

### Configuration
- Use `caarlos0/env` for parsing environment variables into structs.
- Use `joho/godotenv` for loading `.env` files in development.
- Provide a `.env.example` with all required variables documented.

### Logging
- Use `log/slog` (standard library, Go 1.21+). No third-party logging libraries needed.
- Use structured fields: `slog.Info("request handled", "method", r.Method, "path", r.URL.Path)`.

### Testing
- Tests live next to the code they test (`foo_test.go` alongside `foo.go`).
- Use `_test` package suffix for black-box tests where possible.
- Run with `-race` flag always: `go test -race ./...`.

### Generated files
Generated code (templ templates, mockgen mocks, protobuf) should be excluded from both linting and coverage. The provided configs exclude these patterns:
- `_templ.go` (templ HTML templates)
- `mock_*.go` (uber-go/mock generated mocks)
- `.pb.go` (protobuf generated code)

Add your own patterns to `.golangci.yml` (exclusions), `test.yml` (coverage grep), and the `test-coverage` justfile recipe.

### Naming
- Packages are lowercase, single-word where possible. No underscores.
- Interfaces describe behaviour: `Reader`, `Syncer`, not `IReader`.
- Keep exported surface area small. If in doubt, keep it unexported.
