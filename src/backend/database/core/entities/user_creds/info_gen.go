package user_creds

func (i UserCreds) Id() string {
	return i.ID
}

func (i UserCreds) Username() string {
	return i.Uname.String
}

func (i UserCreds) Email() string {
	return i.EmailAddress.String
}

func (i UserCreds) Pwd() string {
	return i.PasswordHash.String
}
