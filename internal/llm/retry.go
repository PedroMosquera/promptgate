package llm

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type RetryOptions struct {
	MaxAttempts int
	BaseDelay   time.Duration
}

type retrying struct {
	next Client
	opts RetryOptions
}

// WithRetry retries transient failures with exponential backoff; other errors
// are returned on the first try.
func WithRetry(next Client, opts RetryOptions) Client {
	if opts.MaxAttempts < 1 {
		opts.MaxAttempts = 1
	}
	return &retrying{next: next, opts: opts}
}

func (r *retrying) Complete(ctx context.Context, req Request) (Response, error) {
	var lastErr error
	for attempt := 0; attempt < r.opts.MaxAttempts; attempt++ {
		resp, err := r.next.Complete(ctx, req)
		if err == nil {
			return resp, nil
		}
		if !errors.Is(err, ErrTransient) {
			return Response{}, err
		}
		lastErr = err
		if attempt == r.opts.MaxAttempts-1 {
			break
		}
		select {
		case <-ctx.Done():
			return Response{}, ctx.Err()
		case <-time.After(backoff(attempt, r.opts.BaseDelay)):
		}
	}
	return Response{}, fmt.Errorf("llm: giving up after %d attempts: %w", r.opts.MaxAttempts, lastErr)
}

func backoff(attempt int, base time.Duration) time.Duration {
	if base <= 0 {
		return 0
	}
	return base * (1 << attempt)
}
