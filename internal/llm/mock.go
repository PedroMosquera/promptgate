package llm

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Mock is a deterministic, no-cost backend used for local development and tests.
type Mock struct {
	Latency time.Duration
}

func NewMock() *Mock { return &Mock{} }

func (m *Mock) Complete(ctx context.Context, req Request) (Response, error) {
	if m.Latency > 0 {
		select {
		case <-ctx.Done():
			return Response{}, ctx.Err()
		case <-time.After(m.Latency):
		}
	}
	if err := ctx.Err(); err != nil {
		return Response{}, err
	}
	return Response{
		Model:      req.Model,
		Completion: canned(req.Model, req.Prompt),
	}, nil
}

func canned(model, prompt string) string {
	words := len(strings.Fields(prompt))
	return fmt.Sprintf("[%s] %d word(s): %s", model, words, echo(prompt))
}

func echo(prompt string) string {
	const max = 60
	prompt = strings.TrimSpace(prompt)
	if len(prompt) <= max {
		return prompt
	}
	return prompt[:max] + "..."
}
