package task_sobes

import (
	"context"
	"sync"
	"time"
)

type Loader func(ctx context.Context, key string) (string, error)

type cacheEntry struct {
	value     string
	expiresAt time.Time
}

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
}

type call struct {
	done  chan struct{}
	value string
	err   error
}
type cache struct {
	loader  Loader
	ttl     time.Duration
	loading map[string]*call
	record  map[string]cacheEntry
	mu      sync.Mutex
}

func NewCache(loader Loader, ttl time.Duration) Cache {
	return &cache{
		loader:  loader,
		ttl:     ttl,
		record:  make(map[string]cacheEntry),
		loading: make(map[string]*call),
	}
}

func (c *cache) Get(ctx context.Context, key string) (string, error) {
	now := time.Now()
	c.mu.Lock()
	entry, ok := c.record[key]
	if ok && now.Before(entry.expiresAt) {
		value := entry.value
		c.mu.Unlock()
		return value, nil
	}

	load, ok := c.loading[key]
	if ok {
		c.mu.Unlock()
		select {
		case <-load.done:
			return load.value, load.err
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}

	newCall := &call{
		done: make(chan struct{}),
	}

	c.loading[key] = newCall

	c.mu.Unlock()

	value, err := c.loader(ctx, key)

	c.mu.Lock()

	newCall.value = value
	newCall.err = err

	if err == nil {
		c.record[key] = cacheEntry{
			value:     value,
			expiresAt: time.Now().Add(c.ttl),
		}
	}
	delete(c.loading, key)
	close(newCall.done)

	c.mu.Unlock()

	return value, err
}
