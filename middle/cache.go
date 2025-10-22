package middle

import (
	"sync"
	"time"
)

type ExpirableCache struct {
	mutex    sync.RWMutex
	cacheMap map[string]cacheValue
}

type cacheValue struct {
	value     any
	createdOn time.Time
	ttl       time.Duration
}

func (c *ExpirableCache) Init() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.cacheMap == nil {
		c.cacheMap = make(map[string]cacheValue)
	}
}

func (c *ExpirableCache) Set(key string, value any, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.cacheMap == nil {
		c.cacheMap = make(map[string]cacheValue)
	}

	c.cacheMap[key] = cacheValue{
		value:     value,
		createdOn: time.Now(),
		ttl:       ttl,
	}
	print("cache set", key, value)
}

func (c *ExpirableCache) Get(key string) (any, bool) {
	c.mutex.RLock()
	v, ok := c.cacheMap[key]
	c.mutex.RUnlock()

	if !ok {
		println("cache miss", key)
		return "", false
	}

	if v.ttl > 0 && time.Since(v.createdOn) > v.ttl {
		println("cache miss", key)
		c.mutex.Lock()
		delete(c.cacheMap, key)
		c.mutex.Unlock()
		return "", false
	}

	println("cache hit", key)
	return v.value, true
}

func (c *ExpirableCache) Delete(key string) {
	delete(c.cacheMap, key)
}

func (c *ExpirableCache) GetString(key string) (string, bool) {
	value, ok := c.Get(key)
	if !ok {
		return "", false
	}
	valueStr, ok := value.(string)
	if !ok {
		return "", false
	}
	return valueStr, true
}
