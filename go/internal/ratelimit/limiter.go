// Package ratelimit is a per-key token bucket: each API key gets its own bucket.
package ratelimit

import (
	"math"
	"sync"
	"time"
)

type bucket struct {
	tokens float64
	last   time.Time
}

type Limiter struct {
	mu      sync.Mutex
	buckets map[string]*bucket
	rate    float64 // tokens per second
	burst   float64 // bucket capacity

	now func() time.Time // injected in tests
}

func New(rate, burst float64) *Limiter {
	return &Limiter{
		buckets: make(map[string]*bucket),
		rate:    rate,
		burst:   burst,
		now:     time.Now,
	}
}

// Allow reports whether key may send one request now, consuming a token if so.
func (l *Limiter) Allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.now()
	b, ok := l.buckets[key]
	if !ok {
		l.buckets[key] = &bucket{tokens: l.burst - 1, last: now}
		return true
	}

	elapsed := now.Sub(b.last).Seconds()
	b.tokens = math.Min(l.burst, b.tokens+elapsed*l.rate)
	b.last = now

	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}
