# cache
A golang memcache with LRU to swap

Install
===

	go get github.com/g4zhuj/cache

Usage
===


### init

```
// If maxItemSize is zero, the cache has no limit.
	var cache Cache
	maxSize = 1024
	cache = NewMemCache(maxSize)
```

### Get

```
value, ok := cache.Get("key")
```

### Set

```
cache.Set("key", "value")
```

### Delete

```
cache.Delete("delKey")
```


### Status

```
status := cache.Status()
fmt.Println("Gets Count: ",	status.Gets)
fmt.Println("Hits Count: ",	status.Hits)
fmt.Println("MaxItemSize: ",	status.MaxItemSize)
fmt.Println("CurrentSize: ",	status.CurrentSize)
```
	
	

