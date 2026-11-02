# Notes for coding agents

promptgate is four small, independent tracks built around the same agent
session traces. Work out which track the person you are helping is on, then
read that track's notes and ignore the others. If it is not clear, ask them
or point them at `CHALLENGES.md`.

| Track | Folder | Agent notes | Skills |
|---|---|---|---|
| Go gateway | `go/` | `go/AGENTS.md` | `go`, `go-debugging` |
| Web console | `web/` | `web/AGENTS.md` | `frontend`, `frontend-testing`, `session-data` |
| Python service | `python/` | `python/AGENTS.md` | `python` |
| Agent trace | `agents/` | `agents/README.md` | `session-data` |

Each folder's `README.md` says how to run it and `TASKS.md` lists the work.
The tracks do not depend on each other: none of them needs another one
running.

## Skills and rules

- Skills live in `.claude/skills/<name>/SKILL.md`. Each one is a short guide
  for a specific job (getting oriented, debugging, testing). They are plain
  Markdown: if your assistant does not load skills on its own, open the one
  you need and read it.
- Rules live in `.claude/rules/`. Each rule names the paths it applies to in
  its front matter. Read the ones that match the files you are editing.

## Shared data

`fixtures/sessions/*.ndjson` are raw agent session traces, one JSON event per
line. `python/` reads them directly and `web/public/api/` is generated from
them. Tests and documented numbers in both tracks depend on their exact
contents. Do not edit them to make a change work.
