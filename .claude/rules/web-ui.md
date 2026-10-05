---
paths:
  - "web/src/**"
---

# UI conventions in the console

- Styling is plain CSS in `src/index.css`, using the custom properties
  defined in `:root` (`--panel`, `--text-dim`, `--ok`, `--fail`,
  `--unknown`, ...). No inline `style` props, no CSS-in-JS, no UI kit. The
  overview panel is the one place that breaks this and is due for cleanup.
- One component per file, default export, named after the file.
- Interactive elements are real `button`s and `a`s, never a clickable `div`.
  Anything that expands or collapses carries `aria-expanded`.
- Status is never conveyed by color alone: a badge also carries its text.
- Durations go through `formatDuration`, timestamps through
  `formatTimestamp`, and long strings through `truncate`, all in
  `src/lib/format.ts`.
