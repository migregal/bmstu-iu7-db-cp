package usercreds

import (
	"database/sql"
	"neural_storage/cube/core/entities/user"
	"neural_storage/database/core/entities/user_creds"
)

func toDBEntity(info user.Info) user_creds.UserCreds {
	data := user_creds.UserCreds{}

	if info.Id() != nil {
		data.ID = *info.Id()
	}
	if info.Username() != nil {
		data.Uname = sql.NullString{String: *info.Username(), Valid: true}
	}
	if info.Email() != nil {
		data.EmailAddress = sql.NullString{String: *info.Email(), Valid: true}
	}
	if info.Pwd() != nil {
		data.PasswordHash = sql.NullString{String: *info.Pwd(), Valid: true}
	}
	return data
}

func fromDBEntity(info user_creds.UserCreds) user.Info {
	data := user.Info{}

	if info.Id() != "" {
		temp := info.Id()
		data.SetId(&temp)
	}
	if info.Username() != "" {
		temp := info.Username()
		data.SetUsername(&temp)
	}
	if info.Pwd() != "" {
		temp := info.Pwd()
		data.SetPwd(&temp)
	}
	if info.Email() != "" {
		temp := info.Email()
		data.SetFullname(&temp)
	}

	return data
}

func fromDBEntities(info []user_creds.UserCreds) []user.Info {
	data := make([]user.Info, 0, len(info))

	for i := range info {
		data = append(data, fromDBEntity(info[i]))
	}

	return data
}
