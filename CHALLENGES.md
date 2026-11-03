# Challenges

Find your role below and work in the folder it names. Your interviewer will
confirm which one at the start. If you are unsure, ask.

| Role | Folder | Tasks |
|---|---|---|
| Backend, Go | `go/` | `go/TASKS.md` |
| Backend, Python | `python/` | `python/TASKS.md` |
| Frontend, React | `web/` | `web/TASKS.md` |
| Agent engineering | `go/`, then `agents/` | `go/TASKS.md`, then `agents/TASKS.md` |

When you finish, or when time is up, fill in `HANDOFF.md` at the root.

## Backend, Go

The gateway in `go/`. Requires Go 1.26.10 or newer, nothing else. From the
repo root:

```
make go-test       # or: cd go && go test ./...
```

## Backend, Python

The scoring service in `python/`. Requires Python 3.10 or newer.

```
make py-install    # creates python/.venv and installs pytest into it
make py-test
make py-run        # localhost:8081
```

## Frontend, React

The console in `web/`. Requires Node 18 or newer.

```
make web-install
make web-dev       # localhost:5173
make web-test
```

Your interviewer will tell you which piece of work in `web/TASKS.md` is your
main one.

## Agent engineering

Start with the Go gateway, then move to the recorded agent session in
`agents/`. Your interviewer will tell you when to switch. `agents/` needs
nothing installed: it is a trace to read and reason about. Start with
`agents/README.md`.

## Shared data

`fixtures/sessions/` holds raw agent session traces, one JSON event per line.
`python/` and `web/` both read them. Do not edit them to make a change work.
