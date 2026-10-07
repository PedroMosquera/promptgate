# Keep the console live

**From:** support lead, on-call rotation
**Priority:** high, we feel this on every incident

## What happens today

During an incident, whoever is on support keeps the console open on a second
screen, usually for hours. The console loads its data once. To see a session
that started after they opened it, or new events in the session they are
watching, they reload the page, find the session again, and reopen the event
they were reading.

Most of our incidents involve the gateway itself, so this is exactly when it
is slow or flaky.

## What we want

The console keeps itself up to date while it is open:

- New sessions show up in the list within about 10 seconds, without a reload.
- The open session picks up new events within about 10 seconds, without a
  reload.

## Done when

- Both of the above work in `npm run dev`.
- There are tests for the refresh.

## Trying it locally

There is no gateway in development. The payloads are static files under
`public/api/`, and the dev server reads them from disk on every request. To
make "the gateway" record something, edit one of those files while
`npm run dev` is running: append an event to a session's `events`, or add a
session to `sessions.json` along with a matching `sessions/<id>.json`.
Revert those edits when you are done.
