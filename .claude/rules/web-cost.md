---
paths:
  - "web/src/**"
---

# Cost figures

- `sessionCostFlat` in `src/lib/cost.ts` is the billing-of-record figure. It
  is the number Finance reconciles against the provider invoice each month,
  and the one the sessions list shows.
- `eventCost` re-derives cost from individual `llm.usage` events. The summary
  panel uses it as a per-event breakdown; treat it as an approximation.
- If the two disagree for a session, the list's figure is the correct one.
  Bring the other view in line with it, not the other way round.
- Show cost as `$` with four decimals, as the list does.
- Do not change the rate constants without sign-off from Finance.
