# Notes for coding agents

This is a small Go service, an LLM gateway. Standard library only, no web
framework. If you are helping someone work on it, here is the lay of the land.

## Build, run, test

- `make run` starts the server on `:8080`.
- `make test` runs the unit tests.
- `make test-race` runs them under the race detector. Reach for it whenever you
  touch state that is shared across requests.
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
