package component

import (
	"context"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
	"shellrean.id/golang_auth/domain"
)

func GetCacheConnection() domain.CacheRepository {
	ctx := context.Background()
	config := bigcache.DefaultConfig(10 * time.Minute)

	cache, err := bigcache.New(ctx, config)
	if err != nil {
		log.Fatalf("Error when connecting to cache: %v", err)
	}

	return cache
}
