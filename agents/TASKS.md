# Tasks (agent trace)

Read `README.md` first for how the session started and what the harness
records. Then work through the questions below, out loud. Your interviewer
will ask follow-ups as you go. There is no code to write, and a clear chain
from evidence to conclusion is worth more than getting through every
question.

Use any tool you like to read the trace, AI assistants included. If you use
one, check what it tells you against the trace itself.

## 1. What happened?

Walk through the session in your own words. What did the agent find out, and
when did it have enough to answer the alert?

## 2. What went wrong in the harness?

The session did not end the way it should have. Say what you think went
wrong, and point at the events that support it. Be explicit about what the
trace proves and what it only suggests.

## 3. Where would you look, and what would you change?

The harness code is not in this repo. Describe where you would expect the
fault to live, what you would check first, and what you would change so a
session like this one cannot burn its whole time budget again.

## Stretch

Every model call this agent made would go through a gateway like the one in
`go/`. What could the gateway notice about a session like this one, and
should it act on it or leave that to the harness?
