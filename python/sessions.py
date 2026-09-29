"""Loads agent session traces from fixtures/sessions/ and builds summaries."""
import json
import os
from datetime import datetime, timezone

FIXTURES_DIR = os.path.join(os.path.dirname(__file__), "..", "fixtures", "sessions")

TERMINAL_TYPES = {"task.completed", "task.failed"}


def _parse_ts(ts):
    return datetime.fromisoformat(ts.replace("Z", "+00:00"))


def load_events(session_id):
    """Returns the ordered list of event dicts for one session."""
    path = os.path.join(FIXTURES_DIR, f"{session_id}.ndjson")
    events = []
    with open(path) as f:
        for line in f:
            line = line.strip()
            if line:
                events.append(json.loads(line))
    return events


def list_session_ids():
    return sorted(
        name[:-len(".ndjson")]
        for name in os.listdir(FIXTURES_DIR)
        if name.endswith(".ndjson")
    )


def _status(events):
    """Infers a session's status from its events.

    Looks for a task.failed event; if there isn't one, calls it completed.
    """
    for e in events:
        if e["type"] == "task.failed":
            return "failed"
    return "completed"


def summarize(session_id):
    events = load_events(session_id)
    input_tokens = sum(e.get("input_tokens", 0) for e in events if e["type"] == "llm.usage")
    output_tokens = sum(e.get("output_tokens", 0) for e in events if e["type"] == "llm.usage")
    cache_read_input_tokens = sum(
        e.get("cache_read_input_tokens", 0) for e in events if e["type"] == "llm.usage"
    )
    started_at = events[0]["ts"]
    duration_ms = int((_parse_ts(events[-1]["ts"]) - _parse_ts(started_at)).total_seconds() * 1000)
    return {
        "id": session_id,
        "status": _status(events),
        "started_at": started_at,
        "duration_ms": duration_ms,
        "input_tokens": input_tokens,
        "output_tokens": output_tokens,
        "cache_read_input_tokens": cache_read_input_tokens,
        "event_count": len(events),
    }


def detail(session_id):
    events = load_events(session_id)
    return {"summary": summarize(session_id), "events": events}
