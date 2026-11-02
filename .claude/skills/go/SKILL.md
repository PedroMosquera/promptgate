---
name: go
description: Get oriented in the promptgate Go gateway in go/ - how to build, run, and test it, and where the important code lives. Use when first opening go/.
---

# Working on the promptgate gateway

A small Go HTTP gateway in front of a mocked language model, in `go/`. Run
every command below from `go/`.

## First steps

1. Confirm a baseline:
   - `go test ./...` runs the unit suite.
   - `go test -race ./...` runs it under the race detector.
2. Run it and send a request:
   - `go run ./cmd/promptgate`
   - `curl -s -H 'X-API-Key: demo' -d '{"model":"gpt-x","prompt":"hi"}' localhost:8080/v1/complete`
3. Read `go/README.md` for endpoints and configuration, and `go/AGENTS.md`
   for the shape of the code.

## The request path

handler -> rate limit -> cache -> model backend. Each stage lives in its own
package under `go/internal/`. Start in `internal/gateway/handler.go` and
follow the calls outward.

## Next

`go/TASKS.md` lists the work. Do the items in order.
