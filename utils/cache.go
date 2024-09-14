package utils

import "sync"

var cache = struct {
	sync.RWMutex
	m map[string]interface{}
}{
	m: make(map[string]interface{}),
}

// InitCache initializes the in-memory cache
func InitCache() {
	cache = struct {
		sync.RWMutex
		m map[string]interface{}
	}{
		m: make(map[string]interface{}),
	}
}

// GetCache retrieves a value from the cache
func GetCache(key string) (interface{}, bool) {
	cache.RLock()
	defer cache.RUnlock()
	val, ok := cache.m[key]
	return val, ok
}

// SetCache sets a value in the cache
func SetCache(key string, value interface{}) {
	cache.Lock()
	defer cache.Unlock()
	cache.m[key] = value
}
