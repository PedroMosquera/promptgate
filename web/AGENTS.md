# Notes for coding agents: web console

React 18 and TypeScript on Vite, `react-router-dom` 6, Vitest with Testing
Library. There is no backend in development: session payloads are static
files under `public/api/`, served by the Vite dev server.

## Commands

Run these from `web/`.

- `npm ci` installs exactly what `package-lock.json` pins.
- `npm run dev` serves the console on `localhost:5173`.
- `npm test` runs the Vitest suite once.
- `npm run build` produces the production bundle in `dist/`.

## Definition of done

A change is done when `npm test` is green and the behavior has been checked
in the browser. The suite runs with type-checking on, so a green run also
means the change compiles; there is no need to run a full `npm run build`
locally, CI builds every push.

## Layout

```
src/api         fetchJson wrapper and the two session calls
src/lib         cost and formatting, pure functions
src/components  layout, timeline, timeline row, summary panel, status badge
src/pages       routed pages: overview (/), session detail (/sessions/:id),
                and the sessions list, which ConsoleLayout keeps mounted
                as a sidebar on every route
src/types.ts    payload shapes, mirroring fixtures/sessions
public/api      session payloads, generated from ../fixtures/sessions
```

`fetchJson` adds an artificial delay that grows with payload size, so local
loading behaves like the deployed gateway. Large sessions take over a second
to arrive. Keep that in mind for anything involving loading, ordering, or
timing, and in tests.

## Rules

The files in `../.claude/rules/` whose name starts with `web-` apply to
everything under `web/`. Read them before editing:

- `web-data-fetching.md` - how pages load data, and what not to add.
- `web-cost.md` - which cost figure is authoritative.
- `web-ui.md` - styling, accessibility, and component conventions.

## Skills

- `frontend` - getting oriented in the console.
- `session-data` - what each field in a session trace means.
- `frontend-testing` - writing component tests here, including timers.
