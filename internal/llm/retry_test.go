package llm

import (
	"context"
	"errors"
	"testing"
	"time"
)

type flaky struct {
	failures int
	calls    int
	err      error
}

func (f *flaky) Complete(_ context.Context, _ Request) (Response, error) {
	f.calls++
	if f.calls <= f.failures {
		return Response{}, f.err
	}
	return Response{Completion: "ok"}, nil
}

func TestRetry(t *testing.T) {
	tests := []struct {
		name        string
		failures    int
		failErr     error
		maxAttempts int
		wantCalls   int
		wantOK      bool
	}{
		{"succeeds first try", 0, ErrTransient, 3, 1, true},
		{"recovers after transient failures", 2, ErrTransient, 3, 3, true},
		{"exhausts attempts", 5, ErrTransient, 3, 3, false},
		{"does not retry permanent error", 5, errors.New("bad request"), 3, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &flaky{failures: tt.failures, err: tt.failErr}
			c := WithRetry(f, RetryOptions{MaxAttempts: tt.maxAttempts})

			_, err := c.Complete(context.Background(), Request{})
			if (err == nil) != tt.wantOK {
				t.Errorf("ok = %v, want %v (err=%v)", err == nil, tt.wantOK, err)
			}
			if f.calls != tt.wantCalls {
				t.Errorf("calls = %d, want %d", f.calls, tt.wantCalls)
			}
		})
	}
}

func TestRetryStopsWhenContextCanceled(t *testing.T) {
	f := &flaky{failures: 10, err: ErrTransient}
	c := WithRetry(f, RetryOptions{MaxAttempts: 5, BaseDelay: time.Hour})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := c.Complete(ctx, Request{})
	if !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
	if f.calls != 1 {
		t.Errorf("calls = %d, want 1 (should not keep retrying after cancel)", f.calls)
	}
}
