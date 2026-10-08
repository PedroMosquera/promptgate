---
paths:
  - "web/src/**"
---

# Data fetching in the console

- All network access goes through `src/api/sessions.ts`, which uses
  `fetchJson` from `src/api/client.ts`. Components never call `fetch`
  directly.
- Pages own their data. A page fetches in a `useEffect` keyed on the route
  params it depends on, and passes data down as props.
- Do not add a data-fetching or state-management library (React Query, SWR,
  Redux, Zustand, and the like). Any new runtime dependency in
  `package.json` needs sign-off from the team; solve it with React and the
  browser.
- Every load has a visible loading state and a visible error state. Do not
  swallow errors: `.catch(() => {})`, or a catch that quietly substitutes an
  empty list, hides a failure from the person on call.
- Whenever a page loads or reloads its data, clear what it was showing first
  (`setData(null)`) so the loading state shows, and if the load fails, show
  the error in place of the data. Never leave old data on screen next to a
  loading or error state: the person on call must not read stale data as if
  it were current.
