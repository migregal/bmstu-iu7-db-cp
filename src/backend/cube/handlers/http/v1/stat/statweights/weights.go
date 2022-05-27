package statweights

import "neural_storage/cube/core/ports/interactors"

type Handler struct {
	resolver interactors.NeuralNetworkInteractor
}

func New(resolver interactors.NeuralNetworkInteractor) Handler {
	return Handler{resolver: resolver}
}
