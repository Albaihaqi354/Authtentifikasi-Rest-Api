package component

import (
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
	"shellrean.id/golang_auth/domain"
)

func GetCacheConnection() domain.CacheRepository {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Fatalf("error When Connect Cache %s", err.Error())
	}
	return cache
}
