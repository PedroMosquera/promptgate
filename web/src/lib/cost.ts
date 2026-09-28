export const RATE_INPUT_PER_1K = 0.003;
export const RATE_OUTPUT_PER_1K = 0.015;
export const CACHE_READ_DISCOUNT = 0.1;

/** Rough per-session estimate for the list view. */
export function sessionCostFlat(s: {
  input_tokens: number;
  output_tokens: number;
}): number {
  return (
    (s.input_tokens / 1000) * RATE_INPUT_PER_1K +
    (s.output_tokens / 1000) * RATE_OUTPUT_PER_1K
  );
}

/** Per-event cost, billing cached reads at the reduced rate. */
export function eventCost(e: {
  input_tokens?: number;
  output_tokens?: number;
  cache_read_input_tokens?: number;
}): number {
  const cached = e.cache_read_input_tokens ?? 0;
  const fresh = Math.max(0, (e.input_tokens ?? 0) - cached);
  return (
    (fresh / 1000) * RATE_INPUT_PER_1K +
    (cached / 1000) * RATE_INPUT_PER_1K * CACHE_READ_DISCOUNT +
    ((e.output_tokens ?? 0) / 1000) * RATE_OUTPUT_PER_1K
  );
}
