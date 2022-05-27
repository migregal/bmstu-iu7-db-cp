package hotstorage

import (
	"neural_storage/cache/core/ports/config"
	"neural_storage/cache/core/services/interactors/cache"
)

type Cache struct {
	i cache.CacheInteractor
}

func New(params config.CacheConfig) *Cache {
	return &Cache{i: cache.New(params.ConnParams())}
}
