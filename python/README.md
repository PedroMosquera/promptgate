# promptgate session scoring service

A small HTTP service over agent session traces: the kind of record a gateway
like `promptgate` would emit for each request it proxies to a model. It reads
raw traces from `../fixtures/sessions/` and serves per-session summaries with
an estimated cost.

Standard library only, plus `pytest` for the test suite. No framework, no
database, no external services.

## Running it

```
python3 -m venv .venv
.venv/bin/python3 -m pip install -r requirements.txt   # only pytest
.venv/bin/python3 server.py                            # listens on :8081
```

A virtual environment avoids `pip`'s "externally managed environment" error
on systems (Homebrew's Python included) that block installing into the
system interpreter directly. `make py-install` does the same thing from the
repo root.

```
PORT=9000 .venv/bin/python3 server.py       # or pick a port
```

```
.venv/bin/python3 -m pytest                 # unit tests
```

If `python3 -m pytest` already works without any of this (pytest already on
your machine), you can skip the virtual environment and use `python3`
directly throughout.

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

`TASKS.md` has the tasks. `../fixtures/sessions/` is the data, shared with
the frontend track; nothing here depends on that track running.
