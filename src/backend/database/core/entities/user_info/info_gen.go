package user_info

import "time"

func (i UserInfo) GetID() string {
	return i.ID
}

func (i UserInfo) GetUsername() string {
	return i.Username.String
}

func (i UserInfo) GetEmail() string {
	return i.Email.String
}

func (i UserInfo) GetPasswordHash() string {
	return i.Password.String
}

func (i UserInfo) GetFullName() string {
	return i.FullName.String
}

func (i UserInfo) GetFlags() uint64 {
	return i.Flags
}

func (i UserInfo) GetBlockedUntil() time.Time {
	return i.Until
}
