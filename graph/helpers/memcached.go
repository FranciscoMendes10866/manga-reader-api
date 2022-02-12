package helpers

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

var MemcachedClient = memcache.New("localhost:11211")

var CacheExpirationSeconds int32 = int32(15 * time.Minute.Seconds())

func SaveToCache(key string, buffer []byte) {
	MemcachedClient.Set(&memcache.Item{
		Key:        key,
		Value:      buffer,
		Expiration: CacheExpirationSeconds,
	})

	fmt.Println("[%a] Saved to cache", key)
}
