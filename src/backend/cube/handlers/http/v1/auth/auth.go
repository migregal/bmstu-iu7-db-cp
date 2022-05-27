package auth

import "neural_storage/cube/core/ports/interactors"

type Handler struct {
	resolver interactors.UserInfoInteractor
}

func NewHandler(resolver interactors.UserInfoInteractor) Handler {
	return Handler{resolver: resolver}
}

var UserIdIdentityKey = "user_id"
var UserFlagsIdentityKey = "flags"

type User struct {
	ID       string
	Email    string
	Username string
	Flags    uint64
}
