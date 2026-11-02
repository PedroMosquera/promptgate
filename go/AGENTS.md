# Notes for coding agents: Go gateway

A small Go service, an LLM gateway. Standard library only, no web framework.

## Build, run, test

Run these from `go/`. The root `Makefile` has the same commands as `go-run`,
`go-test`, `go-test-race` and `go-vet`.

- `go run ./cmd/promptgate` starts the server on `:8080`.
- `go test ./...` runs the unit tests.
- `go test -race ./...` runs them under the race detector. Reach for it
  whenever you touch state that is shared across requests.
- `go build ./...` and `go vet ./...` should stay clean.

## Shape of the code

The request path is: handler -> rate limit -> cache -> model backend.

- `internal/gateway` owns HTTP. `handler.go` has the request handling,
  `server.go` has the server plumbing and the JSON helpers.
- `internal/llm` is the seam to the model. `Client` is the interface the gateway
  depends on. `Mock` is the deterministic backend used everywhere. `WithRetry`
  wraps a client with backoff on `ErrTransient`.
- `internal/cache` is the shared response cache.
- `internal/ratelimit` is a per-key token bucket with an injectable clock.

## Conventions

- Standard library first. Prefer a small interface over a dependency.
- Errors are wrapped with `%w` and returned, never swallowed.
- Handlers map a failure to a status code and a small JSON body. Internal detail
  goes to the logs, not the response.
- `context.Context` flows from the request into the model call. Work that blocks
  or runs in the background should honor it.

## Skills

- `go` - getting oriented in the gateway.
- `go-debugging` - targeted tests, the race detector, and curl against the
  running server.
