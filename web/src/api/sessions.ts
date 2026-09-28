import { fetchJson } from "./client";
import type { SessionDetailPayload, SessionSummary } from "../types";

export async function listSessions(): Promise<SessionSummary[]> {
  return fetchJson<SessionSummary[]>("/api/sessions.json");
}

export async function getSession(id: string): Promise<SessionDetailPayload> {
  return fetchJson<SessionDetailPayload>(`/api/sessions/${id}.json`);
}
