// Package llm is the gateway's seam to the model backend: callers depend only on
// Client, so the mock can be swapped for a real provider.
package llm

import (
	"context"
	"errors"
)

// ErrTransient marks a failure worth retrying; permanent failures are returned bare.
var ErrTransient = errors.New("llm: transient failure")

type Request struct {
	Model  string
	Prompt string
}

type Response struct {
	Model      string
	Completion string
}

type Client interface {
	Complete(ctx context.Context, req Request) (Response, error)
}
