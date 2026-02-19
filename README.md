# agent-toolkit

Templates, configs, and AI tooling for bootstrapping new projects.

Everything here is opinionated and copy-paste ready. Grab what you need, drop it into your repo, and customise the `<!-- TODO: customise -->` markers.

## Structure

```
agent-toolkit/
├── languages/go/               # Go project template
│   ├── .github/workflows/      #   CI/CD pipelines (lint, test, build, dev-image, secrets, renovate)
│   ├── scripts/pre-commit      #   TruffleHog pre-commit hook
│   ├── .golangci.yml           #   Linter config (v2)
│   ├── .goreleaser.yml         #   Cross-platform release builds
│   ├── justfile                #   Build/test/lint recipes
│   ├── renovate.json           #   Dependency update config
│   ├── .gitignore              #   Go-specific ignores
│   ├── .dockerignore           #   Docker build context ignores
│   ├── .env.example            #   Environment variable template
│   ├── CONTRIBUTING.md         #   PR guidelines
│   ├── SECURITY.md             #   Vulnerability reporting
│   ├── CHANGELOG.md            #   Release notes template
│   └── README.md               #   Setup guide, packages, conventions
├── docker/                     # Docker templates
│   ├── Dockerfile.go           #   Multi-stage Go build (Chainguard)
│   ├── docker-compose.yml      #   App + Postgres + Redis
│   └── README.md               #   Best practices
├── opencode/                   # Slash commands and plugins for opencode
└── claude-code/                # Slash commands and plugins for Claude Code
```

## Usage

### Starting a new Go project

```bash
# Copy the full Go template into your new repo
cp -r languages/go/.github     your-repo/
cp -r languages/go/scripts     your-repo/
cp languages/go/.golangci.yml  your-repo/
cp languages/go/.goreleaser.yml your-repo/
cp languages/go/justfile       your-repo/
cp languages/go/renovate.json  your-repo/
cp languages/go/.gitignore     your-repo/
cp languages/go/.dockerignore  your-repo/
cp languages/go/.env.example   your-repo/
cp languages/go/CONTRIBUTING.md your-repo/
cp languages/go/SECURITY.md    your-repo/
cp languages/go/CHANGELOG.md   your-repo/

# Copy the Docker files if you need containerisation
cp docker/Dockerfile.go        your-repo/Dockerfile
cp docker/docker-compose.yml   your-repo/docker-compose.yml
```

Then search for `<!-- TODO: customise -->` comments and fill in your project-specific values.

### Adding AI tooling

Check the `opencode/` and `claude-code/` directories for available slash commands and plugins. Each has its own README with installation instructions.

## Conventions

- Config files include comments explaining the reasoning behind each choice
- Templates target production use, not toy examples
- Docs are concise and actionable
