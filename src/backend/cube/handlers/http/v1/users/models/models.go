package models

import (
	"neural_storage/cube/core/ports/cache"
	"neural_storage/cube/core/ports/interactors"
)

type Handler struct {
	resolver interactors.NeuralNetworkInteractor
	cache    cache.CacheInteractor
}

func New(resolver interactors.NeuralNetworkInteractor, cache cache.CacheInteractor) Handler {
	return Handler{resolver: resolver, cache: cache}
}
