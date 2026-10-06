# Notes for coding agents

promptgate is three small, independent pieces built around the same agent
session traces. If you are helping someone work on it, read this file, then
the notes for the part you are touching.

| Track | Code | Agent notes | Skills |
|---|---|---|---|
| Go gateway | `cmd/`, `internal/` | "Go gateway" below | `onboarding`, `debugging` |
| Web console | `web/` | `web/AGENTS.md` | `frontend`, `session-data`, `frontend-testing` |
| Python service | `python/` | `python/README.md` | `python` |

`CHALLENGES.md` explains the tracks. They do not depend on each other: none
of them needs another one running.

## Skills and rules

- Skills live in `.claude/skills/<name>/SKILL.md`. Each one is a short guide
  for a specific job (getting oriented, debugging, testing). They are plain
  Markdown: if your assistant does not load skills on its own, open the one
  you need and read it.
- Rules live in `.claude/rules/`. Each rule names the paths it applies to in
  its front matter. Read the ones that match the files you are editing.

## Shared data

`fixtures/sessions/*.ndjson` are raw agent session traces, one JSON event per
line. `web/` and `python/` both read them, and tests and documented numbers in
both tracks depend on their exact contents. Do not edit them to make a change
work.

## Go gateway

A small Go service, an LLM gateway. Standard library only, no web framework.

### Build, run, test

- `make run` starts the server on `:8080`.
- `make test` runs the unit tests.
- `make test-race` runs them under the race detector. Reach for it whenever you
  touch state that is shared across requests.
- `go build ./...` and `go vet ./...` should stay clean.

### Shape of the code

The request path is: handler -> rate limit -> cache -> model backend.

- `internal/gateway` owns HTTP. `handler.go` has the request handling,
  `server.go` has the server plumbing and the JSON helpers.
- `internal/llm` is the seam to the model. `Client` is the interface the gateway
  depends on. `Mock` is the deterministic backend used everywhere. `WithRetry`
  wraps a client with backoff on `ErrTransient`.
- `internal/cache` is the shared response cache.
- `internal/ratelimit` is a per-key token bucket with an injectable clock.

### Conventions

- Standard library first. Prefer a small interface over a dependency.
- Errors are wrapped with `%w` and returned, never swallowed.
- Handlers map a failure to a status code and a small JSON body. Internal detail
  goes to the logs, not the response.
- `context.Context` flows from the request into the model call. Work that blocks
  or runs in the background should honor it.
