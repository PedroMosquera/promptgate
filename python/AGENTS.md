# Notes for coding agents: Python service

A small HTTP service over the agent session traces in `../fixtures/sessions/`.
Standard library only, plus `pytest` for the tests. No framework.

## Commands

From the repo root: `make py-install`, `make py-test`, `make py-run`. From
`python/`, once `.venv` exists:

- `.venv/bin/python -m pytest` runs the suite.
- `PORT=8081 .venv/bin/python server.py` serves on `localhost:8081`.

## Layout

```
sessions.py   loads fixtures, builds per-session summaries
cost.py       cost estimate from a summary
server.py     HTTP layer: routes summaries and details to JSON
tests/        pytest suite
```

## Conventions

- Standard library only in the service. `pytest` is the one dependency, and
  only for tests.
- The pricing table in `README.md` is the reference for cost.

## Skills

- `python` - getting oriented in the service.
