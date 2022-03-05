package internal

import (
	"time"
)

type Cache struct {
	Value     interface{}
	ExpiresIn time.Time
}

// simplest cache on memory using hashmap
var cacheData = make(map[string]Cache)

func SetCache(key string, value interface{}, expiresIn time.Time) {
	cacheData[key] = Cache{
		Value:     value,
		ExpiresIn: expiresIn,
	}
}

func GetCache(key string) (interface{}, error) {
	data, ok := cacheData[key]
	if !ok {
		return nil, ErrCacheNotFound
	}
	if time.Until(data.ExpiresIn) < 0 {
		delete(cacheData, key)
		return nil, ErrCacheNotFound
	}
	return data.Value, nil
}
