# Tasks (Go track)

promptgate works, mostly. Spend some time with it and work through the items
below in order. You are not expected to finish all of them. Getting cleanly
through the first two is a good outcome, and we care more about how you work than
how far you get.

Think out loud as you go. Use whatever tools you normally reach for, including AI
assistants if that is part of your workflow. Either way, be ready to explain the
code you end up with as if you had written it yourself.

## 1. A failing test

`go test ./...` has one failing test. Start there. Understand what it is telling
you, find the cause, and make it pass. A clear explanation of the root cause is
worth as much as the fix.

## 2. Concurrency under load

promptgate handles many requests at the same time. Convince yourself that the
state shared on the request path is safe when it does. If it is not, fix it, and
show how you proved the fix holds.

## 3. Per-request timeout

Add a configurable timeout for the call to the model backend. If the backend
takes longer than the configured duration, cancel it and return the client a
`504 Gateway Timeout`. Add a test that covers it.

## Stretch, only if there is time

Add `GET /v1/stats` returning the number of cache hits and misses since startup.
Think about what has to hold for those counters to be correct while many requests
are in flight.
