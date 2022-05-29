package auth

import (
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
)

type Handler struct {
	resolver interactors.UserInfoInteractor

	lg *logger.Logger
}

func NewHandler(lg *logger.Logger, resolver interactors.UserInfoInteractor) Handler {
	return Handler{resolver: resolver, lg: lg}
}

var UserIdIdentityKey = "user_id"
var UserFlagsIdentityKey = "flags"

type User struct {
	ID       string
	Email    string
	Username string
	Flags    uint64
}
