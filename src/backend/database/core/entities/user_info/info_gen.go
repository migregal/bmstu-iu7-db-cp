package user_info

func (i UserInfo) Id() string {
	return i.ID
}

func (i UserInfo) Username() string {
	return i.Uname.String
}

func (i UserInfo) Email() string {
	return i.EmailAddress.String
}

func (i UserInfo) FullName() string {
	return i.FName.String
}
