---
name: debugging
description: How to reproduce, observe, and diagnose problems in promptgate - running targeted tests, the race detector, and exercising the running server with curl. Use when investigating a failing test or unexpected runtime behavior.
---

# Debugging promptgate

## Run one test, verbosely

```
go test -run TestName -v ./internal/gateway/
```

## Race detector

Everything on the request path is touched by many goroutines at once. To check
for data races:

```
go test -race ./...
```

A clean run prints `ok`. A race prints a `DATA RACE` report with the two stacks
that conflicted and the file and line of each access.

## Exercise the running server

```
make run
curl -i -H 'X-API-Key: demo' -d '{"model":"m","prompt":"p"}' localhost:8080/v1/complete
```

`-i` shows the status line, which is where the gateway signals rate limiting
(429), bad input (400), and upstream failures (502).

## A method that works

Reproduce first, ideally with a test that fails for the same reason. Form one
hypothesis, change one thing, and let the test tell you whether you were right.
