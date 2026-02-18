# Contributing

Contributions are welcome. This document covers the basics for getting a PR merged.

## Getting Started

<!-- TODO: customise -- replace with your repo URL and binary name -->

```bash
git clone https://github.com/yourorg/your-project.git
cd your-project
cp .env.example .env
just build
```

Requires Go 1.24+ and [just](https://github.com/casey/just).

## Before Submitting a PR

1. Run `just check` and make sure it passes (lint, test, build)
2. Add tests for new functionality
3. Keep commits focused on a single change

## Project Structure

<!-- TODO: customise -- describe your internal/ packages -->

```
cmd/your-app/           Entry point
internal/
  config/               Environment variable parsing
  api/                  HTTP handlers and routing
  store/                Database/storage layer
  models/               Domain types
  errors/               Sentinel errors
  logging/              Structured logging (slog)
```
