package userinfo

import (
	"database/sql"
	"neural_storage/cube/core/entities/user"
	"neural_storage/database/core/entities/user_info"
)

func toDBEntity(info user.Info) user_info.UserInfo {
	data := user_info.UserInfo{}

	if info.Id() != nil {
		data.ID = *info.Id()
	}
	if info.Username() != nil {
		data.Uname = sql.NullString{String: *info.Username(), Valid: true}
	}
	if info.Fullname() != nil {
		data.FName = sql.NullString{String: *info.Fullname(), Valid: true}
	}
	if info.Email() != nil {
		data.EmailAddress = sql.NullString{String: *info.Email(), Valid: true}
	}
	return data
}

func fromDBEntity(info user_info.UserInfo) user.Info {
	data := user.Info{}

	if info.Id() != "" {
		temp := info.Id()
		data.SetId(&temp)
	}
	if info.Username() != "" {
		temp := info.Username()
		data.SetUsername(&temp)
	}
	if info.FullName() != "" {
		temp := info.FullName()
		data.SetFullname(&temp)
	}
	if info.Email() != "" {
		temp := info.Email()
		data.SetFullname(&temp)
	}

	return data
}

func fromDBEntities(info []user_info.UserInfo) []user.Info {
	data := make([]user.Info, 0, len(info))

	for i := range info {
		data = append(data, fromDBEntity(info[i]))
	}

	return data
}
