import { useState } from "react";
import { formatTimestamp } from "../lib/format";
import type { SessionEvent } from "../types";

export default function TimelineRow({ event }: { event: SessionEvent }) {
  const [open, setOpen] = useState(false);
  const body = event.output ?? event.input ?? event.message;

  return (
    <li className="timeline-row">
      <button className="row-head" onClick={() => setOpen((v) => !v)}>
        <span className="row-seq">{event.harness_seq}</span>
        <span className="row-type">{event.type}</span>
        <span className="row-ts">{formatTimestamp(event.ts)}</span>
        {body ? <span className="row-caret">{open ? "-" : "+"}</span> : null}
      </button>
      {open && body ? <pre className="row-body">{body}</pre> : null}
    </li>
  );
}
