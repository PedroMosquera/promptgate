import type { SessionStatus } from "../types";

export default function StatusBadge({ status }: { status: SessionStatus }) {
  const known = status === "running" || status === "completed" || status === "failed";
  return (
    <span className={`status-badge ${known ? status : "unknown"}`}>
      {status}
    </span>
  );
}
