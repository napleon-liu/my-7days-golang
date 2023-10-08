package geecache

/*
*	cache implements lru, add mutex mu
 */

import (
	"gee-cache/lru"
	"sync"
)

type cache struct {
	m          sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

// add is one implement to lru cache's Add method
func (c *cache) add(key string, value ByteView) {
	c.m.Lock()
	defer c.m.Unlock()

	// Lazy initialization
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}

	return
}
