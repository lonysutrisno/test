package script

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]string
	lock sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) {
	// c.lock.Lock()
	// defer c.lock.Unlock()

	c.data[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) Process(key string) {
	c.lock.RLock()
	value, exists := c.data[key]
	c.lock.RUnlock()

	if exists {
		fmt.Printf("Processing key: %s, value: %s\n", key, value)
		time.Sleep(1 * time.Second)
		fmt.Printf("Finished processing key: %s\n", key)
	} else {
		fmt.Printf("Key: %s does not exist\n", key)
	}
}

func Caching() {
	cache := NewCache()

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	var wg sync.WaitGroup
	keys := []string{"key1", "key2", "key3", "key4"}

	wg.Add(len(keys))

	for _, key := range keys {
		go func(k string) {
			defer wg.Done()
			cache.Process(k)
		}(key)
	}

	wg.Wait()
}
