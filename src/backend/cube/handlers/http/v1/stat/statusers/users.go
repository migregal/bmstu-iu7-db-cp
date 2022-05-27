package statusers

import "neural_storage/cube/core/ports/interactors"

type Handler struct {
	resolver interactors.UserInfoInteractor
}

func New(resolver interactors.UserInfoInteractor) Handler {
	return Handler{resolver: resolver}
}
