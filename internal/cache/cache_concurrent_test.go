//go:build race

package cache

import (
	"fmt"
	"sync"
	"testing"
)

// TestConcurrentAccess drives the cache from many goroutines. Race-tagged: it
// exists to be run under `go test -race`.
func TestConcurrentAccess(t *testing.T) {
	c := New()
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("p%d", i%8)
			c.Set("m", key, "v")
			c.Get("m", key)
		}(i)
	}
	wg.Wait()
}
