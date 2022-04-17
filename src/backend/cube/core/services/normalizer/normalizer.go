package normalizer

import (
	"neural_storage/cube/core/ports/config"
	ports "neural_storage/cube/core/ports/normalizer"
)

type Normalizer struct {
	conf config.NormalizerConfig
}

func NewNormalizer(conf config.NormalizerConfig) ports.Normalizer {
	return &Normalizer{conf}
}
