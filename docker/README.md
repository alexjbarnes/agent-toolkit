# Docker

Templates for containerising Go applications.

## Files

- `Dockerfile.go` -- Multi-stage build for a Go binary. Produces a minimal image.
- `docker-compose.yml` -- Typical setup with app, Postgres, and Redis.

## Dockerfile Best Practices

1. **Use multi-stage builds.** Build in a full Go image, copy only the binary into a minimal runtime image. This keeps final images small (often under 20MB).

2. **Use Chainguard base images.** `cgr.dev/chainguard/go` for the build stage and `cgr.dev/chainguard/static` for the runtime. These are distroless images with minimal attack surface and no shell, package manager, or unnecessary libraries.

3. **Copy go.mod and go.sum first.** This lets Docker cache the dependency download layer separately from your source code. Source changes do not re-download dependencies.

4. **Set `CGO_ENABLED=0`.** This produces a statically linked binary that runs in a scratch/distroless container without libc.

5. **Do not run as root.** Chainguard images run as a non-root user by default. If using other base images, add a `USER` directive.

6. **Use a `.dockerignore`.** Exclude `bin/`, `.env`, `.git/`, coverage files, and anything else that should not end up in the build context.

## When to Use Compose

Use `docker-compose.yml` for:
- **Local development** when your app depends on external services (databases, caches, queues)
- **CI integration tests** that need real service dependencies
- **Self-hosted deployments** where a simple orchestrator is sufficient

Do not use compose for production orchestration at scale. Use Kubernetes, Nomad, or a managed platform for that.

## Quick Start

```bash
# Build the image
docker build -f Dockerfile.go -t your-app:latest .

# Run with compose
docker compose up -d

# View logs
docker compose logs -f app

# Tear down
docker compose down -v
```
