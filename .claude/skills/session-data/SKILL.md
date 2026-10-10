---
name: session-data
description: What the fields in a promptgate session trace mean - event types, ordering, identity, status, tokens. Use before computing anything from session data in web/ or python/, or when a number on screen looks wrong.
---

# Reading a session trace

A session is one agent run that went through the gateway. The raw trace is
`fixtures/sessions/<id>.ndjson`, one event per line, written once the
harness disconnects. The console's payloads in `web/public/api/` are
generated from those files: `sessions.json` is the list of summaries,
`sessions/<id>.json` is `{ summary, events }` for one session. A session
that is still running has no trace file yet; its payload is what the
gateway has received so far.

## Events

One model call produces three events, in this order:

| Event | Carries |
|---|---|
| `llm.request.started` | `provider`, `model`, `is_stream` |
| `llm.request.completed` | `latency_ms`, `retries_attempted` |
| `llm.usage` | `input_tokens`, `output_tokens`, `cache_read_input_tokens`, `cache_creation_input_tokens`, and `latency_ms` again |

`llm.usage.latency_ms` repeats the value from the `llm.request.completed`
event before it. It is the same call, so do not add the two together.

One tool call produces two events:

| Event | Carries |
|---|---|
| `tool.call` | `tool`, `input`, `tool_call_id` |
| `tool.result` | `tool`, `output`, `duration_ms`, `retries_attempted`, `tool_call_id` |

Pair a call with its result by `tool_call_id`, not by position.

A session ends with `task.completed` or `task.failed` (with a `message`), or
it has neither when the trace was cut off.

## Ordering and identity

- `harness_seq` is strictly increasing within a session and is the order the
  harness emitted the events. Within one session it identifies an event.
- `ts` is wall-clock time, and it is not guaranteed to increase. A
  `tool.result` is stamped when the tool finished, which can be after the
  next event was emitted. Order by `harness_seq`, not by `ts`.

## Summary fields

- `status` - `running` while the harness is still connected, so events can
  still arrive. After it disconnects: `completed` if there is a
  `task.completed`, `failed` if there is a `task.failed`, `unknown` if there
  is neither. `unknown` means the trace has no ending; it does not mean the
  session succeeded, and it does not mean it is still running.
- `started_at` - `ts` of the first event.
- `duration_ms` - from the first event's `ts` to the last event's `ts`.
- `input_tokens`, `output_tokens`, `cache_read_input_tokens` - sums over the
  session's `llm.usage` events.
- `event_count` - number of events.

## Tokens

`input_tokens` counts every input token, including the ones served from the
prompt cache. `cache_read_input_tokens` is the part of `input_tokens` that
came from the cache, so it is never larger than `input_tokens`. Fresh input
is `input_tokens - cache_read_input_tokens`.

## Fixtures worth knowing

| Session | Why it is there |
|---|---|
| `sess-quick-lookup` | short and clean |
| `sess-cache-heavy` | most of its input is served from the cache |
| `sess-stuck-kubectl`, `sess-stuck-grep` | end in `task.failed` |
| `sess-slow-report` | long, ends in `task.completed` |
| `sess-truncated` | no terminal event |
| `sess-rollout-check` | still `running`, so no trace file yet; its last `tool.call` has no result yet |
