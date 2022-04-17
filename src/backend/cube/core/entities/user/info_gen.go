package user

import "time"

func (i *Info) Id() *string {
	return i.id
}

func (i *Info) SetId(id *string) {
	i.id = id
}

func (i *Info) Username() *string {
	return i.username
}

func (i *Info) SetUsername(username *string) {
	i.username = username
}

func (i *Info) Email() *string {
	return i.email
}

func (i *Info) SetEmail(email *string) {
	i.email = email
}

func (i *Info) Pwd() *string {
	return i.pwd
}

func (i *Info) SetPwd(pwdHash *string) {
	i.pwd = pwdHash
}

func (i *Info) BlockedUntil() *time.Time {
	return i.blockedUntil
}

func (i *Info) SetBlockedUntil(blockedUntil *time.Time) {
	i.blockedUntil = blockedUntil
}
