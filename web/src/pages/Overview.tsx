import { useEffect, useState } from "react";
import { listSessions } from "../api/sessions";
import { sessionCostFlat } from "../lib/cost";
import type { SessionSummary } from "../types";

// Deliberately unstyled compared to the rest of the console: this is the
// first thing we shipped to answer "can we see cost at a glance", and it
// shows. Whoever picks this up next should make it match.
export default function Overview() {
  const [sessions, setSessions] = useState<SessionSummary[] | null>(null);

  useEffect(() => {
    listSessions().then(setSessions).catch(() => setSessions([]));
  }, []);

  if (!sessions) return <p>Loading...</p>;

  const withCost = sessions.map((s) => ({ s, cost: sessionCostFlat(s) }));
  const maxCost = Math.max(...withCost.map((x) => x.cost), 0.0001);

  return (
    <div
      style={{
        background: "#f4f5f7",
        borderRadius: "8px",
        padding: "24px",
        boxShadow: "0 2px 6px rgba(0, 0, 0, 0.15)",
        fontFamily: "system-ui, sans-serif",
        color: "#1a1a1a",
      }}
    >
      <h2 style={{ fontSize: "1.3rem", marginTop: 0 }}>Analytics Overview</h2>
      <p style={{ color: "#555" }}>
        Select a session on the left to dive in, or check out the cost
        breakdown below.
      </p>
      <div
        style={{
          background: "#ffffff",
          border: "1px solid #ddd",
          borderRadius: "6px",
          padding: "20px",
          marginTop: "16px",
        }}
      >
        {withCost.map(({ s, cost }) => (
          <div key={s.id} style={{ marginBottom: "12px" }}>
            <div
              style={{
                display: "flex",
                justifyContent: "space-between",
                fontSize: "0.85rem",
                marginBottom: "4px",
              }}
            >
              <span>{s.id}</span>
              <span>${cost.toFixed(4)}</span>
            </div>
            <div
              style={{
                background: "#e5e7eb",
                borderRadius: "4px",
                height: "10px",
                overflow: "hidden",
              }}
            >
              <div
                style={{
                  width: `${(cost / maxCost) * 100}%`,
                  height: "100%",
                  background: "#3b82f6",
                  borderRadius: "4px",
                }}
              />
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
