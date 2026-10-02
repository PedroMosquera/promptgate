---
name: frontend-testing
description: How to write and run tests for the promptgate web console - Vitest and Testing Library setup, rendering routed pages, controlling the fetch delay, and fake timers. Use when adding or fixing a test under web/src.
---

# Testing the console

Vitest runs in jsdom with Testing Library's matchers loaded from
`src/setupTests.ts`. Globals (`describe`, `it`, `expect`, `vi`) are on.

```
npm test                          # whole suite, once
npx vitest run src/lib            # one directory
npx vitest src/components         # watch mode
```

`npm test` runs with type-checking on, so a type error in a test or in the
code it imports fails the run.

## Where tests go

Next to the code: `format.ts` and `format.test.ts`. Pure helpers in
`src/lib` get plain unit tests. Components get Testing Library tests.

## Querying

Query the way a user finds things: `getByRole`, `getByLabelText`,
`getByText`. The timeline's filter is a labelled `<select>`
(`getByLabelText("Event type")`), and each timeline row head is a `button`.
Avoid querying by class name.

`@testing-library/user-event` is not installed. Use `fireEvent` from
`@testing-library/react`; do not add a dependency for this.

## Pages need a router

Anything that uses `useParams`, `Link`, or `useNavigate` needs a router
around it:

```tsx
import { MemoryRouter, Route, Routes } from "react-router-dom";

render(
  <MemoryRouter initialEntries={["/sessions/sess-quick-lookup"]}>
    <Routes>
      <Route path="/sessions/:id" element={<SessionDetail />} />
    </Routes>
  </MemoryRouter>,
);
```

## Data and the fetch delay

`fetchJson` waits between 100 and 1500 ms after the response arrives, longer
for bigger payloads. Two ways to deal with it:

- **Mock the API module** when the test is about rendering, not loading:

  ```ts
  vi.mock("../api/sessions", () => ({
    listSessions: vi.fn(),
    getSession: vi.fn(),
  }));
  ```

  then `vi.mocked(getSession).mockResolvedValue(payload)` per test. Load a
  real payload with `import payload from "../../public/api/sessions/sess-quick-lookup.json"`.

- **Control time** when the test is about ordering or timing: mock `fetch`
  with `vi.spyOn(globalThis, "fetch")`, switch to fake timers, and advance
  them yourself.

## Fake timers

```ts
beforeEach(() => vi.useFakeTimers());
afterEach(() => vi.useRealTimers());

// inside a test
await act(async () => {
  await vi.advanceTimersByTimeAsync(2000);
});
```

- Use the `Async` variants. They flush promise callbacks between timers,
  which `fetchJson`'s delay needs.
- `findBy*` and `waitFor` poll on timers. Under fake timers, advance time
  explicitly and then use `getBy*`, rather than waiting.
- Restore real timers in `afterEach`, or the next test inherits fake ones.

## A test is evidence only if it can fail

Before trusting a new test, break the code it covers and watch it go red for
the reason you expect. Then put the code back.
