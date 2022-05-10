package user

import "time"

type Info struct {
	id           *string
	username     *string
	fullname     *string
	email        *string
	pwd          *string
	blockedUntil *time.Time
}

func NewInfo(id *string, username *string, fullname *string, email *string, pwd *string, blockedUntil *time.Time) *Info {
	return &Info{id: id, username: username, email: email, pwd: pwd, blockedUntil: blockedUntil}
}
