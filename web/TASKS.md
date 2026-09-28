# Tasks (frontend track)

The console works, mostly. Pick one bug below as a warm-up, then spend most of
your time on `BRIEF.md`. You are not expected to fix every bug or finish the
brief. Getting cleanly through a warm-up bug and making real, well-reasoned
progress on the brief is a good outcome.

Think out loud as you go. Use whatever tools you normally reach for, including
AI assistants if that is part of your workflow. Either way, be ready to explain
the code you end up with as if you had written it yourself.

## Bug pool

Your interviewer will point you at one of these to start. The others are
there if you have time left.

### A. A row that expands itself

1. Run `npm run dev`, open a session with a mix of event types (for example
   `sess-stuck-kubectl`).
2. Expand any row in the timeline by clicking it.
3. Change the event-type filter to a type that still leaves several rows
   visible.
4. Look at the row that ends up in the same position the expanded one was in.

What you should see: a row is expanded that you never clicked, showing a
different event's output than the one you actually opened.

### B. Two numbers that disagree

1. On `/`, note the cost shown for `sess-cache-heavy`.
2. Open that session and look at the cost in the summary panel.
3. Compare the two numbers.

They should be the same session's cost. They are not.

### C. A session that shows someone else's data

1. On `/`, click into `sess-stuck-kubectl`.
2. Before the page finishes loading, click back, then immediately click into
   `sess-quick-lookup`.
3. Watch what renders.

The URL and the page heading say `sess-quick-lookup`. The timeline sometimes
does not.

## Stretch

If you finish the brief with time to spare, pick up a second bug from the
pool, or look for a second field in the fixtures that the console does not
yet surface.
