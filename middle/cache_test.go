package middle

import (
	"sync"
	"testing"
	"time"
)

func TestExpirableCache(t *testing.T) {
	cache := ExpirableCache{}
	cache.Init()

	// Set and Get
	cache.Set("aa", "bb", time.Minute)
	v, ok := cache.Get("aa")
	if !ok {
		t.Fatalf("expected key 'aa' to exist")
	}
	if v != "bb" {
		t.Fatalf("expected value 'bb', got '%s'", v)
	}

	// Expiration test
	cache.Set("short", "lived", 10*time.Millisecond)
	time.Sleep(15 * time.Millisecond)
	if _, ok := cache.Get("short"); ok {
		t.Fatalf("expected 'short' key to have expired")
	}

	// Nonexistent key
	if _, ok := cache.Get("does_not_exist"); ok {
		t.Fatalf("expected nonexistent key to return ok=false")
	}

	// Zero TTL
	cache.Set("forever", "alive", 0)
	time.Sleep(20 * time.Millisecond)
	v, ok = cache.Get("forever")
	if !ok || v != "alive" {
		t.Fatalf("expected key 'forever' to persist without expiration")
	}

	// Thread safety
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key" + time.Now().String()
			cache.Set(key, "val", time.Second)
			cache.Get(key)
		}(i)
	}
	wg.Wait()
}
