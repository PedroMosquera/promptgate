// Package cache is a small in-memory response cache shared across requests.
package cache

type Cache struct {
	entries map[string]string
}

func New() *Cache {
	return &Cache{entries: make(map[string]string)}
}

func (c *Cache) Get(model, prompt string) (string, bool) {
	v, ok := c.entries[key(model, prompt)]
	return v, ok
}

func (c *Cache) Set(model, prompt, completion string) {
	c.entries[key(model, prompt)] = completion
}

func key(model, prompt string) string {
	return prompt
}
