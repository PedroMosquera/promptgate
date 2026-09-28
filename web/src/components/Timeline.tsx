import { useState } from "react";
import TimelineRow from "./TimelineRow";
import type { SessionEvent } from "../types";

export default function Timeline({ events }: { events: SessionEvent[] }) {
  const [filter, setFilter] = useState<string>("all");
  const types = Array.from(new Set(events.map((e) => e.type))).sort();
  const shown = filter === "all" ? events : events.filter((e) => e.type === filter);

  return (
    <section className="timeline">
      <div className="timeline-controls">
        <label htmlFor="type-filter">Event type</label>
        <select
          id="type-filter"
          value={filter}
          onChange={(e) => setFilter(e.target.value)}
        >
          <option value="all">all ({events.length})</option>
          {types.map((t) => (
            <option key={t} value={t}>
              {t}
            </option>
          ))}
        </select>
      </div>
      <ol className="timeline-rows">
        {shown.map((event, i) => (
          <TimelineRow key={i} event={event} />
        ))}
      </ol>
    </section>
  );
}
