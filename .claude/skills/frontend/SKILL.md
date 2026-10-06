---
name: frontend
description: Get oriented in the promptgate web console - how to run, test, and navigate it, and which notes and rules apply. Use when first opening web/.
---

# Working on the promptgate console

A React console over agent session traces, in `web/`.

## First steps

1. `npm ci`, then `npm test` to confirm a clean baseline.
2. `npm run dev`, open `localhost:5173`, click through both pages.
3. Read `web/README.md` for the pages and the data, and `web/AGENTS.md` for
   commands, layout, and the definition of done.
4. Read the `web-*` rules in `.claude/rules/` before changing code.

## Where things live

- `src/pages` - the routed pages, plus the sessions list that stays mounted
  as a sidebar.
- `src/components` - the layout, the timeline, its rows, the summary panel,
  the status badge.
- `src/lib` - cost calculation and formatting, both pure functions.
- `src/api` - the fetch wrapper and the two calls it exposes.
- `public/api` - the session payloads themselves, generated from
  `fixtures/sessions/`.

For what the fields in a payload mean, load the `session-data` skill. For
tests, load `frontend-testing`.

## A method that works

Reproduce first. Several of the bugs in `TASKS.md` only show up after a
specific sequence of clicks, not from reading the code alone. Get the
symptom on screen before you go looking for the cause, and check the fix the
same way you found the bug.

## Next

`TASKS.md` has the bug pool and points at the larger piece of work.
