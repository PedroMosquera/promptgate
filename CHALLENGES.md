# Challenges

promptgate has three independent tracks. Pick the one that matches the
session you were invited to, or ask if you are unsure.

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

## Python track

`python/`, a small HTTP service over the same agent session traces. Requires
Python 3.9 or newer.

```
make py-install
make py-run       # localhost:8081
```

Tasks are in `python/TASKS.md`.

## Shared data

All three tracks read from `fixtures/sessions/`, a set of raw agent session
traces in NDJSON. No track requires another to be running.
