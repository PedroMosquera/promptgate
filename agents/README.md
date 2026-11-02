# Agent session trace

`stuck-loop.ndjson` is the full event stream of one on-call agent session,
recorded by the harness that runs our agents. There is no code to run here:
the work is reading the trace and reasoning about the agent and the harness
that produced it.

## How the session started

The harness opened the session in response to this alert:

```
api-server-7d9f8b-xk2p4 in production has restarted 47 times in the last hour
```

The agent can run shell commands through one tool, `exec_cli`. The session
has ended.

## What the harness records

One JSON event per line, ordered by `harness_seq`. The `session-data` skill in
`../.claude/skills/` describes the common events and fields in detail. In
short:

| Event | Emitted when |
|---|---|
| `llm.request.started` | the harness sends a request to the model |
| `llm.request.completed` | the model's response has arrived |
| `llm.usage` | token counts for that model call |
| `tool.call` | the model asked for a tool, with its input |
| `tool.progress` | a long-running tool is still working |
| `tool.result` | the tool finished, with its output and duration |
| `loop.forced_stop` | the harness's loop guard stopped the session |
| `task.completed` | the agent finished and reported an answer |
| `task.failed` | the session ended without an answer, with a `message` |

Two mechanisms can end a session that is not making progress:

- **The loop guard**, inside the harness. It counts model calls against a
  per-session budget and stops the session when the budget is spent.
- **The stuck-session reaper**, a watchdog outside the harness. It kills any
  session that has been running for more than 15 minutes of wall-clock time.

## Reading the trace

Any tool works: a short script, an editor, an AI assistant, or `jq` if you
have it:

```
jq -c '{seq: .harness_seq, ts, type}' stuck-loop.ndjson
```

`TASKS.md` has the questions.
