// Shapes match the NDJSON event schema in fixtures/sessions/ and the JSON
// payloads served from web/public/api/.

export type SessionStatus = "completed" | "failed" | "unknown";

export interface SessionSummary {
  id: string;
  status: SessionStatus;
  started_at: string;
  duration_ms: number;
  input_tokens: number;
  output_tokens: number;
  cache_read_input_tokens: number;
  event_count: number;
}

export type EventType =
  | "llm.request.started"
  | "llm.request.completed"
  | "llm.usage"
  | "tool.call"
  | "tool.result"
  | "task.completed"
  | "task.failed";

export interface SessionEvent {
  type: EventType;
  ts: string;
  harness_seq: number;
  provider?: string;
  model?: string;
  is_stream?: boolean;
  latency_ms?: number;
  retries_attempted?: number;
  input_tokens?: number;
  output_tokens?: number;
  cache_read_input_tokens?: number;
  cache_creation_input_tokens?: number;
  tool?: string;
  input?: string;
  output?: string;
  duration_ms?: number;
  tool_call_id?: string;
  status?: string;
  message?: string;
}

export interface SessionDetailPayload {
  summary: SessionSummary;
  events: SessionEvent[];
}
