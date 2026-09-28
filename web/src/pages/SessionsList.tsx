import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { listSessions } from "../api/sessions";
import { sessionCostFlat } from "../lib/cost";
import { formatDuration, formatTimestamp } from "../lib/format";
import StatusBadge from "../components/StatusBadge";
import type { SessionSummary } from "../types";

export default function SessionsList() {
  const [sessions, setSessions] = useState<SessionSummary[] | null>(null);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    listSessions().then(setSessions).catch((e) => setError(String(e)));
  }, []);

  if (error) return <p className="error">{error}</p>;
  if (!sessions) return <p className="loading">Loading sessions...</p>;

  return (
    <div className="panel">
      <table className="sessions-table">
        <thead>
          <tr>
            <th>Session</th>
            <th>Status</th>
            <th>Started</th>
            <th>Duration</th>
            <th>Events</th>
            <th>Cost</th>
          </tr>
        </thead>
        <tbody>
          {sessions.map((s) => (
            <tr
              key={s.id}
              className="session-row"
              onClick={() => navigate(`/sessions/${s.id}`)}
            >
              <td>
                <Link to={`/sessions/${s.id}`}>{s.id}</Link>
              </td>
              <td>
                <StatusBadge status={s.status} />
              </td>
              <td>{formatTimestamp(s.started_at)}</td>
              <td>{formatDuration(s.duration_ms)}</td>
              <td>{s.event_count}</td>
              <td>${sessionCostFlat(s).toFixed(4)}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
