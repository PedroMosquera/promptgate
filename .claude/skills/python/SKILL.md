---
name: python
description: Get oriented in the promptgate session scoring service - how to run, test, and navigate it. Use when first opening python/.
---

# Working on the promptgate scoring service

A small HTTP service over agent session traces, in `python/`.

## First steps

1. `python3 -m pytest` to confirm a baseline. One test fails; that is
   expected, see `TASKS.md`.
2. `python3 server.py`, then in another terminal `curl localhost:8081/api/sessions`.
3. Read `python/README.md` for the endpoints and the pricing model.

## Where things live

- `sessions.py` - loads fixtures, builds per-session summaries.
- `cost.py` - the cost estimate.
- `server.py` - the HTTP layer, routes to JSON.
- `tests/` - the pytest suite.

## A method that works

Reproduce first, ideally with a test that fails for the same reason. Form one
hypothesis, change one thing, and let the test tell you whether you were
right.

## Next

`TASKS.md` lists the work. Do the items in order.
