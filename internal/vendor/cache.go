package vendor

import (
	"time"

	"github.com/bluele/gcache"
)

const (
	cacheSize = 1_000_000
	cacheTTL  = 1 * time.Hour // default expiration
)

type gCache[T any] struct {
	data gcache.Cache
}

func newTCache[T any]() *gCache[T] {
	return &gCache[T]{
		data: gcache.New(cacheSize).Expiration(cacheTTL).ARC().Build(),
	}
}
func (gc *gCache[T]) update(queryString string, characters T) error {
	return gc.data.SetWithExpire(queryString, characters, 1*time.Hour)
}

func (gc *gCache[T]) read(queryString string) (T, bool) {
	if val, err := gc.data.Get(queryString); err != nil {
		return *new(T), false
	} else {
		if characters, found := val.(T); found {
			return characters, true
		}
	}
	return *new(T), false
}

func (gc *gCache[T]) delete(queryString string) {
	gc.data.Remove(queryString)
}
