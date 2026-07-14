package cache

import "testing"

func TestGetMissThenHit(t *testing.T) {
	c := New()

	if _, ok := c.Get("m", "p"); ok {
		t.Fatal("expected miss on empty cache")
	}

	c.Set("m", "p", "answer")
	got, ok := c.Get("m", "p")
	if !ok || got != "answer" {
		t.Errorf("Get = (%q, %v), want (%q, true)", got, ok, "answer")
	}
}

func TestSetOverwrites(t *testing.T) {
	c := New()
	c.Set("m", "p", "first")
	c.Set("m", "p", "second")

	if got, _ := c.Get("m", "p"); got != "second" {
		t.Errorf("Get = %q, want %q", got, "second")
	}
}
