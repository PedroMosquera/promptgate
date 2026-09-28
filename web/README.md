# promptgate console

A React console over agent session traces: the kind of record a gateway like
`promptgate` would emit for each request it proxies to a model. It lists
sessions and lets you drill into a single session's event timeline, with cost
and duration computed from the trace data.

Session data ships as static JSON under `public/api/`, generated from the raw
traces in `../fixtures/sessions/`. There is no backend process to run.

## Running it

```
npm install
npm run dev        # localhost:5173
```

```
npm test            # unit tests
npm run build        # production build
```

## Pages

- `/` - the sessions list in a persistent sidebar, plus an analytics
  overview panel showing cost per session.
- `/sessions/:id` - a summary panel plus the full event timeline for one
  session, with a filter by event type, shown alongside the sidebar.

## Data

- `GET /api/sessions.json` - array of session summaries.
- `GET /api/sessions/<id>.json` - one session's summary plus its full event
  list.

Each event in a session's list carries, among other fields, the tool called,
its input and output, and how long the call took (`latency_ms`) and how many
times it was retried (`retries_attempted`). Neither of the last two is shown
in the timeline today.

## Layout

```
src/api        fetch wrapper and typed session calls
src/lib        cost calculation and formatting helpers
src/components shared UI: status badge, timeline, timeline row, summary panel
src/pages      the two routed pages
public/api     generated session payloads, source of truth is fixtures/sessions
```

## Working on this

`TASKS.md` has a small set of bugs to work through. `BRIEF.md` describes the
larger piece of work. `HANDOFF.md` is where you leave notes on what you built,
what you left out, and what you would do next.
