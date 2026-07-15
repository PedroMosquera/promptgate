---
name: onboarding
description: Get oriented in the promptgate codebase - how to build, run, and test it, and where the important code lives. Use when first opening this repository.
---

# Onboarding to promptgate

promptgate is a small Go HTTP gateway in front of a mocked language model.

## First steps

1. Confirm a clean baseline:
   - `make test` runs the unit suite.
   - `make test-race` runs it under the race detector.
2. Run it and send a request:
   - `make run`
   - `curl -s -H 'X-API-Key: demo' -d '{"model":"gpt-x","prompt":"hi"}' localhost:8080/v1/complete`
3. Read `README.md` for endpoints and configuration, and `AGENTS.md` for the
   shape of the code.

## The request path

handler -> rate limit -> cache -> model backend. Each stage lives in its own
package under `internal/`. Start in `internal/gateway/handler.go` and follow the
calls outward.

## Next

`TASKS.md` lists the work. Do the items in order.
