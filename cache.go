package cache

// Status return status of cache
type Status struct {
	Gets        int64
	Hits        int64
	MaxItemSize int
	CurrentSize int
}

// Cache this is an interface which defines some common functions
type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Delete(key string)
	Status() *Status
}
