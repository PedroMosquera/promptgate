# promptgate

A small HTTP gateway that sits in front of a large language model. A client POSTs
a prompt, promptgate rate-limits per API key, checks an in-memory cache, and on a
miss calls the model backend (retrying transient failures) before caching and
returning the result.

The backend is mocked, so the service runs with no API keys and costs nothing.
Swapping in a real provider means implementing a single interface in
`internal/llm`.

## Running it

```
make run                 # listens on :8080
PORT=9000 make run       # or pick a port
```

Send a request:

```
curl -s -H 'X-API-Key: demo' \
  -d '{"model":"gpt-x","prompt":"hello"}' \
  localhost:8080/v1/complete
```

```
{"model":"gpt-x","completion":"[gpt-x] 1 word(s): hello","cached":false}
```

A second identical request comes back with `"cached":true`.

## Endpoints

- `POST /v1/complete` - body `{"model": "...", "prompt": "..."}`, requires an
  `X-API-Key` header.
- `GET /healthz` - liveness check.

## Configuration

| Env var        | Default | Meaning                       |
|----------------|---------|-------------------------------|
| `PORT`         | 8080    | Port to listen on             |
| `RATE_PER_SEC` | 5       | Token refill rate per API key |
| `BURST`        | 10      | Bucket size per API key       |

## Tests

```
make test         # unit tests
make test-race    # the same suite under the race detector
```

## Layout

```
cmd/promptgate      entrypoint, config, graceful shutdown
internal/gateway    HTTP handlers and routing
internal/llm        model-backend interface, mock, retry wrapper
internal/cache      in-memory response cache
internal/ratelimit  per-key token bucket
```

## Working on this

`TASKS.md` describes what we would like you to look at. `AGENTS.md` has notes for
coding assistants. Use an AI agent or don't, whichever matches how you work, and
be ready to walk through your reasoning and your changes.
