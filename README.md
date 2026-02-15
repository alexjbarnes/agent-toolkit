# agent-toolkit

Templates, configs, and AI tooling for bootstrapping new projects.

Everything here is opinionated and copy-paste ready. Grab what you need, drop it into your repo, and customise the `<!-- TODO: customise -->` markers.

## Structure

```
agent-toolkit/
├── languages/go/       # Go project template (linter, justfile, gitignore, conventions)
├── docker/             # Dockerfile templates and compose patterns
├── opencode/           # Slash commands and plugins for opencode
└── claude-code/        # Slash commands and plugins for Claude Code
```

## Usage

### Starting a new Go project

```bash
# Copy the Go scaffolding into your new repo
cp languages/go/.golangci.yml  your-repo/
cp languages/go/justfile       your-repo/
cp languages/go/.gitignore     your-repo/

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
