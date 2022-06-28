package cache

import (
	"container/list"
	"sync"
	"sync/atomic"
	//"fmt"
)

// An AtomicInt is an int64 to be accessed atomically.
type AtomicInt int64

// MemCache is an LRU cache. It is safe for concurrent access.
type MemCache struct {
	mutex       sync.RWMutex
	maxItemSize int
	cacheList   *list.List
	cache       map[interface{}]*list.Element
	hits, gets  AtomicInt
}

type Entry struct {
	key   interface{}
	value interface{}
}

//NewMemCache If maxItemSize is zero, the cache has no limit.
//if maxItemSize is not zero, when cache's size beyond maxItemSize,start to swap
func NewMemCache(maxItemSize int) *MemCache {
	return &MemCache{
		maxItemSize: maxItemSize,
		cacheList:   list.New(),
		cache:       make(map[interface{}]*list.Element),
	}
}

//Status return the status of cache
func (c *MemCache) Status() *Status {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return &Status{
		MaxItemSize: c.maxItemSize,
		CurrentSize: c.cacheList.Len(),
		Gets:        c.gets.Get(),
		Hits:        c.hits.Get(),
	}
}

//Get value with key
func (c *MemCache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	c.gets.Add(1)
	if ele, hit := c.cache[key]; hit {
		c.hits.Add(1)
		c.cacheList.MoveToFront(ele)
		return ele.Value.(*Entry).value, true
	}
	return nil, false
}

//Set a value with key
func (c *MemCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.cacheList.MoveToFront(ele)
		ele.Value.(*Entry).value = value
		return
	}

	ele := c.cacheList.PushFront(&Entry{key: key, value: value})
	c.cache[key] = ele
	if c.maxItemSize > 0 && c.cacheList.Len() > c.maxItemSize {
		c.removeOldest()
	}
}

//Delete delete the key
func (c *MemCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if ele, ok := c.cache[key]; ok {
		c.cacheList.Remove(ele)
		key := ele.Value.(*Entry).key
		delete(c.cache, key)
		return
	}
}

//RemoveOldest remove the oldest key
func (c *MemCache) removeOldest() {
	ele := c.cacheList.Back()
	if ele != nil {
		c.cacheList.Remove(ele)
		key := ele.Value.(*Entry).key
		delete(c.cache, key)
	}
}

// Add atomically adds n to i.
func (i *AtomicInt) Add(n int64) {
	atomic.AddInt64((*int64)(i), n)
}

// Get atomically gets the value of i.
func (i *AtomicInt) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}
