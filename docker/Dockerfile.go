# ---- Build stage ----
# Chainguard Go image: minimal, secure, regularly patched.
FROM cgr.dev/chainguard/go AS builder

WORKDIR /src

# Copy dependency manifests first to cache the download layer.
# Source code changes will not invalidated this cache.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Build a statically linked binary.
# CGO_ENABLED=0 ensures no libc dependency.
# <!-- TODO: customise the binary name and cmd path -->
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -ldflags "-X main.Version=${VERSION}" -o /app ./cmd/your-app

# ---- Runtime stage ----
# Chainguard wolfi-base: minimal, secure, regularly patched.
# Use wolfi-base over chainguard/static when the app needs to write
# files at runtime (e.g. SQLite). Switch to chainguard/static if
# the binary is fully stateless (no filesystem writes).
FROM cgr.dev/chainguard/wolfi-base

RUN mkdir -p /data && chown nonroot:nonroot /data
USER nonroot

COPY --from=builder /app /usr/local/bin/app

ENTRYPOINT ["app"]
