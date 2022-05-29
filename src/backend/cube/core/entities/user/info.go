package user

import "time"

type Info struct {
	id           string
	username     string
	fullname     string
	email        string
	pwd          string
	blockedUntil time.Time
	flags        uint64
}

func NewInfo(id string, username string, fullname string, email string, pwd string, flags uint64, blockedUntil time.Time) *Info {
	return &Info{id: id, username: username, fullname: fullname, email: email, pwd: pwd, flags: flags, blockedUntil: blockedUntil}
}
