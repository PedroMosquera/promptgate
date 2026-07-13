package ratelimit

import (
	"testing"
	"time"
)

func TestAllowConsumesBurst(t *testing.T) {
	l := New(0, 3) // no refill, capacity 3
	l.now = func() time.Time { return time.Unix(0, 0) }

	for i := 0; i < 3; i++ {
		if !l.Allow("k") {
			t.Fatalf("request %d should be allowed within burst", i+1)
		}
	}
	if l.Allow("k") {
		t.Error("4th request should be denied once burst is spent")
	}
}

func TestRefillOverTime(t *testing.T) {
	now := time.Unix(0, 0)
	l := New(1, 1) // 1 token/sec, capacity 1
	l.now = func() time.Time { return now }

	if !l.Allow("k") {
		t.Fatal("first request should be allowed")
	}
	if l.Allow("k") {
		t.Fatal("second immediate request should be denied")
	}

	now = now.Add(time.Second)
	if !l.Allow("k") {
		t.Error("request should be allowed after a token refills")
	}
}

func TestKeysAreIsolated(t *testing.T) {
	l := New(0, 1)
	l.now = func() time.Time { return time.Unix(0, 0) }

	if !l.Allow("a") {
		t.Fatal("key a should be allowed")
	}
	if !l.Allow("b") {
		t.Error("key b should have its own bucket")
	}
}
