---
name: frontend
description: Get oriented in the promptgate web console - how to run, test, and navigate it. Use when first opening web/.
---

# Working on the promptgate console

A React console over agent session traces, in `web/`.

## First steps

1. `npm install`, then `npm test` to confirm a clean baseline.
2. `npm run dev`, open `localhost:5173`, click through both pages.
3. Read `web/README.md` for the pages, the data it reads, and the layout.

## Where things live

- `src/pages` - the two routed pages.
- `src/components` - the timeline, its rows, the summary panel, the status
  badge.
- `src/lib` - cost calculation and formatting, both pure functions.
- `src/api` - the fetch wrapper and the two calls it exposes.
- `public/api` - the session payloads themselves, generated from
  `fixtures/sessions/`.

## A method that works

Reproduce first. Several of the bugs in `TASKS.md` only show up after a
specific sequence of clicks, not from reading the code alone. Get the
symptom on screen before you go looking for the cause.

## Next

`TASKS.md` has the bug pool. `BRIEF.md` is the larger piece of work.
