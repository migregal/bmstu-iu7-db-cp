//go:generate mockery --name=CacheConfig --outpkg=mock --output=../../../../config/adapters/cache/mock/ --filename=cache_config.go --structname=CacheConfig

package config

import "neural_storage/cache/core/services/interactors/cache"

type CacheConfig interface {
	IsMocked() bool
	Adapter() string
	ConnParams() cache.Params
}
