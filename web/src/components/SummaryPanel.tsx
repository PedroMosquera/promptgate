import { eventCost } from "../lib/cost";
import { formatDuration } from "../lib/format";
import StatusBadge from "./StatusBadge";
import type { SessionEvent, SessionSummary } from "../types";

export default function SummaryPanel({
  summary,
  events,
}: {
  summary: SessionSummary;
  events: SessionEvent[];
}) {
  const totalCost = events
    .filter((e) => e.type === "llm.usage")
    .reduce((sum, e) => sum + eventCost(e), 0);

  return (
    <div className="panel summary-panel">
      <div className="summary-stat">
        <span className="label">Status</span>
        <span className="value">
          <StatusBadge status={summary.status} />
        </span>
      </div>
      <div className="summary-stat">
        <span className="label">Duration</span>
        <span className="value">{formatDuration(summary.duration_ms)}</span>
      </div>
      <div className="summary-stat">
        <span className="label">Events</span>
        <span className="value">{summary.event_count}</span>
      </div>
      <div className="summary-stat">
        <span className="label">Tokens</span>
        <span className="value">
          {summary.input_tokens + summary.output_tokens}
        </span>
      </div>
      <div className="summary-stat">
        <span className="label">Cost</span>
        <span className="value">${totalCost.toFixed(4)}</span>
      </div>
    </div>
  );
}
