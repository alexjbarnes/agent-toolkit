# ---- Build stage ----
# Chainguard Go image: minimal, secure, regularly patched.
FROM cgr.dev/chainguard/go AS builder

WORKDIR /src

# Copy dependency manifests first to cache the download layer.
# Source code changes will not invalidate this cache.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Build a statically linked binary.
# CGO_ENABLED=0 ensures no libc dependency, so the binary runs in distroless/scratch.
# <!-- TODO: customise the binary name and cmd path -->
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -ldflags "-X main.Version=${VERSION}" -o /app ./cmd/your-app

# ---- Runtime stage ----
# Chainguard static image: distroless, no shell, no package manager.
# ~2MB base. Non-root by default.
FROM cgr.dev/chainguard/static

COPY --from=builder /app /usr/local/bin/app

ENTRYPOINT ["app"]
