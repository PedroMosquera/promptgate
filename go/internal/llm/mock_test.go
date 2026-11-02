package llm

import (
	"context"
	"strings"
	"testing"
)

func TestMockIsDeterministic(t *testing.T) {
	m := NewMock()
	req := Request{Model: "gpt-x", Prompt: "hello there"}

	first, err := m.Complete(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	second, err := m.Complete(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if first.Completion != second.Completion {
		t.Errorf("completion not stable: %q vs %q", first.Completion, second.Completion)
	}
}

func TestMockDependsOnModel(t *testing.T) {
	m := NewMock()
	a, _ := m.Complete(context.Background(), Request{Model: "model-a", Prompt: "same"})
	b, _ := m.Complete(context.Background(), Request{Model: "model-b", Prompt: "same"})

	if a.Completion == b.Completion {
		t.Errorf("expected different completions per model, both were %q", a.Completion)
	}
	if !strings.Contains(b.Completion, "model-b") {
		t.Errorf("completion should mention its model, got %q", b.Completion)
	}
}

func TestMockHonorsCanceledContext(t *testing.T) {
	m := NewMock()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := m.Complete(ctx, Request{Model: "m", Prompt: "p"}); err == nil {
		t.Fatal("expected error from canceled context, got nil")
	}
}
