package vendor

import (
	"time"

	"github.com/bluele/gcache"
)

type gCache struct {
	characters gcache.Cache
}

const (
	cacheSize = 1_000_000
	cacheTTL  = 1 * time.Hour // default expiration
)

func newGCache() *gCache {
	return &gCache{
		characters: gcache.New(cacheSize).Expiration(cacheTTL).ARC().Build(),
	}
}
func (gc *gCache) update(queryString string, characters []Character, expireIn time.Duration) error {
	return gc.characters.SetWithExpire(queryString, characters, expireIn)
}

func (gc *gCache) read(queryString string) ([]Character, bool) {
	if val, err := gc.characters.Get(queryString); err != nil {
		return nil, false
	} else {
		if characters, found := val.([]Character); found {
			return characters, true
		}
	}
	return nil, false
}

func (gc *gCache) delete(queryString string) {
	gc.characters.Remove(queryString)
}
