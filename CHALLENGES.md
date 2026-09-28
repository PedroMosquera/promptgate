# Challenges

promptgate has two independent tracks. Pick the one that matches the session
you were invited to, or ask if you are unsure.

## Go track

The gateway itself: `cmd/`, `internal/`. Requires Go 1.24 or newer, nothing
else. Start with `make test`.

Tasks are in `TASKS.md`.

## Frontend track

`web/`, a React console that reads the same agent session traces the gateway
would produce in production. Requires Node 18 or newer.

```
make web-install
make web-dev      # localhost:5173
```

Tasks are in `web/TASKS.md`. The main piece of work is in `web/BRIEF.md`.

## Shared data

Both tracks read from `fixtures/sessions/`, a set of raw agent session traces
in NDJSON. Neither track requires the other to be running.
