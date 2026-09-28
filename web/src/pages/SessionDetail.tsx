import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { getSession } from "../api/sessions";
import SummaryPanel from "../components/SummaryPanel";
import Timeline from "../components/Timeline";
import type { SessionDetailPayload } from "../types";

export default function SessionDetail() {
  const { id } = useParams<{ id: string }>();
  const [data, setData] = useState<SessionDetailPayload | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;
    setData(null);
    getSession(id).then(setData).catch((e) => setError(String(e)));
  }, [id]);

  return (
    <div>
      <Link to="/" className="back-link">
        &larr; back to sessions
      </Link>
      {error ? <p className="error">{error}</p> : null}
      {!error && !data ? <p className="loading">Loading session...</p> : null}
      {data ? (
        <>
          <SummaryPanel summary={data.summary} events={data.events} />
          <Timeline events={data.events} />
        </>
      ) : null}
    </div>
  );
}
