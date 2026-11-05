# promptgate

promptgate is a small LLM gateway plus the tools around it. A client sends a
prompt, the gateway rate-limits it, checks a cache, and calls a model
backend. Every agent session that goes through it leaves a trace, and the
other parts of this repo read those traces.

The model backend is mocked, so nothing here needs an API key, a network
connection, or money.

## Where to start

Open `CHALLENGES.md`. It tells you which folder is yours for the session you
were invited to, and you can ignore the rest of the repo.

| Folder | What it is | Stack | Start with |
|---|---|---|---|
| `go/` | the gateway itself | Go, standard library only | `go/README.md` |
| `python/` | a scoring service over session traces | Python, standard library plus pytest | `python/README.md` |
| `web/` | the on-call console over session traces | React, TypeScript, Vite | `web/README.md` |
| `agents/` | a recorded agent session that went wrong | NDJSON trace, no code | `agents/README.md` |
| `fixtures/` | session traces shared by `python/` and `web/` | NDJSON | shared, read only |

Each track folder has its own `README.md` (what it is, how to run it) and
`TASKS.md` (what to work on). Read the ones in your folder, not the others.
The tracks do not depend on each other, so none of them needs another one
running.

## Files at the root

- `CHALLENGES.md` - which track is yours, and what each one needs installed.
- `HANDOFF.md` - a short template to fill in at the end, whatever track you
  worked on.
- `AGENTS.md` and `CLAUDE.md` - notes for coding assistants. `AGENTS.md`
  points each assistant at the notes for its track.
- `.claude/skills/` and `.claude/rules/` - short guides and conventions that
  assistants load. They are plain Markdown, so you can read them too.
- `Makefile` - shortcuts for every track, prefixed `go-`, `py-` and `web-`.

## Tools and AI assistants

Use any tool you like during the challenge, AI agents included. We know a
paid AI tool can be an extra cost for you, so feel free not to use one: it
does not count against you. If you do use one, it helps us see how you work
with AI agents in your day-to-day development. Either way, be ready to walk
through your reasoning and every change you made.
