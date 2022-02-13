package helpers

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

var MemcachedClient = memcache.New("localhost:11211")

var FifteenMinutesCacheExpiration int32 = int32(15 * time.Minute.Seconds())
var TwoHoursCacheExpiration int32 = int32(2 * time.Hour.Seconds())

func GetFromCache(key string) ([]byte, error) {
	cacheItem, err := MemcachedClient.Get(key)
	if err != nil {
		return nil, err
	}
	return cacheItem.Value, nil
}

func SaveToCache(key string, buffer []byte, expiration int32) {
	MemcachedClient.Set(&memcache.Item{
		Key:        key,
		Value:      buffer,
		Expiration: expiration,
	})

	fmt.Println("[%a] Saved to cache", key)
}
