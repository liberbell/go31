package cache

import (
	"net/http"
	"sync"
)

type response struct {
	header http.Header
	code   int
	body   []byte
}

type memCache struct {
	lock sync.RWMutex
	data map[string]response
}

var (
	cache = memCache{data: map[string]response{}}
)

func set(resource string, response *response) {
	cache.lock.Lock()

	if response == nil {
		delete(cache.data, resource)
	} else {
		cache.data[resource] = *response
	}
	cache.lock.Unlock()
}

func get(resource string) *response {
	cache.lock.Lock()
}
