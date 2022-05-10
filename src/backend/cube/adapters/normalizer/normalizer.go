package normalizer

import (
	"neural_storage/cube/core/ports/config"
	ports "neural_storage/cube/core/ports/normalizer"
	"neural_storage/cube/core/services/normalizer"
)

func NewNormalizer(conf config.NormalizerConfig) ports.Normalizer {
	return normalizer.NewNormalizer(conf)
}
