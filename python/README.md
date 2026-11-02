# promptgate session scoring service

A small HTTP service over agent session traces: the kind of record a gateway
like `promptgate` would emit for each request it proxies to a model. It reads
raw traces from `../fixtures/sessions/` and serves per-session summaries with
an estimated cost.

Standard library only, plus `pytest` for the test suite. No framework, no
database, no external services.

## Running it

Requires Python 3.9 or newer. From the repo root, the `py-` targets in the
`Makefile` set up a virtual environment in `python/.venv` and use it:

```
make py-install    # python3 -m venv python/.venv, then installs pytest
make py-test       # unit tests
make py-run        # listens on :8081
```

Or by hand, from `python/`:

```
python3 -m venv .venv
.venv/bin/pip install -r requirements.txt
.venv/bin/python -m pytest
PORT=8081 .venv/bin/python server.py
```

The virtual environment also avoids pip's "externally managed environment"
error on Homebrew's Python.

## Endpoints

- `GET /api/sessions` - array of session summaries.
- `GET /api/sessions/<id>` - one session's summary plus its full event list.

## Pricing

Cost is estimated per session from its token usage:

| Rate | Value |
|---|---|
| Input tokens | $0.003 / 1K |
| Output tokens | $0.015 / 1K |
| Cached input tokens | 10% of the input rate |

A session that read most of its input from cache should cost noticeably less
than one that didn't, for the same token counts.

## Layout

```
sessions.py   loads fixtures, builds per-session summaries
cost.py       cost estimate from a summary
server.py     HTTP layer: routes summaries and details to JSON
tests/        pytest suite
```

## Working on this

`TASKS.md` has the tasks. `AGENTS.md` has notes for coding assistants.
`../fixtures/sessions/` is the data, shared with the frontend track; nothing
here depends on that track running.
