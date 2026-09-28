# From the on-call rotation

We keep escalating sessions to engineering that turn out to be fine. Support
sees a session sitting open for a long time in the console, assumes it is
stuck, and pages someone. Twice last week the session they paged about was
`sess-slow-report`, which just took nine minutes doing legitimate work and
finished on its own. Nobody could tell, from what the console shows today,
that it was ever going to finish.

Meanwhile we have had at least one real stuck session (ask engineering about
the `api-server` OOM incident if you want the gory details) where the
session ran for a long time doing nothing useful, and nobody noticed until it
had already burned a chunk of the budget.

Right now the console shows duration and a timeline. It does not help anyone
tell "this is slow but working" apart from "this is stuck." We would like the
console to help with that.

We do not have a fixed idea of what the fix looks like. Whatever you build,
come prepared to walk through it: what you built, what you decided not to
build, and what you think would break first if this went in front of
support tomorrow.
