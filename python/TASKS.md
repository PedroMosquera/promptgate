# Tasks (Python track)

The service works, mostly. Work through the items below in order. You are not
expected to finish all of them. Getting cleanly through the first two is a
good outcome, and we care more about how you work than how far you get.

Think out loud as you go. Use whatever tools you normally reach for, including
AI assistants if that is part of your workflow. Either way, be ready to
explain the code you end up with as if you had written it yourself.

## 1. A failing test

`python3 -m pytest` has one failing test. Start there. Understand what it is
telling you, find the cause, and make it pass. A clear explanation of the
root cause is worth as much as the fix.

## 2. Every session says it succeeded

`GET /api/sessions` reports every session as `"status": "completed"`.
Spot-check the raw traces in `fixtures/sessions/` against what the endpoint
returns for each one. At least one of them did not actually finish.

Fix the status so it reflects what really happened, and think about whether
"completed" and "failed" are the only two states a session can honestly be
in.

## 3. A per-session signals endpoint

Add `GET /api/sessions/<id>/signals`. A "repeat" is a run of consecutive
tool calls that share the same tool and the same input; a run of 4 such
calls has a repeat count of 4, not 3. For each run of 2 or more, report the
tool, the input, the repeat count, and the total duration of that run
(including its first call). The response shape (one object per run, a flat
list, whatever) is your call, state your reasoning for it. A test is part
of the deliverable, and should cover a session with no repeats as well as
one with them.

## Stretch, only if there is time

Add `GET /api/stats` returning fleet-wide cost and the p50 and p95 session
duration. Think about what a percentile means once some sessions never
properly finished.
