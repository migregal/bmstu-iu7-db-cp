package cachecfg

import "neural_storage/cache/core/services/interactors/cache"

type CacheConfig struct {
	CacheAdapter string
	CacheParams  cache.Params
}

func (cfg *CacheConfig) IsMocked() bool {
	return false
}

func (cfg *CacheConfig) Adapter() string {
	return cfg.CacheAdapter
}

func (cfg *CacheConfig) ConnParams() cache.Params {
	return cfg.CacheParams
}
