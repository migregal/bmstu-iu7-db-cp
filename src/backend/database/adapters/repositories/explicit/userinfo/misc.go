package userinfo

import (
	"database/sql"
	"neural_storage/cube/core/entities/user"
	"neural_storage/database/core/entities/user_info"
)

func toDBEntity(info user.Info) user_info.UserInfo {
	data := user_info.UserInfo{}

	if info.ID() != nil {
		data.ID = *info.ID()
	}
	if info.Username() != nil {
		data.Username = sql.NullString{String: *info.Username(), Valid: true}
	}
	if info.Fullname() != nil {
		data.FullName = sql.NullString{String: *info.Fullname(), Valid: true}
	}
	if info.Email() != nil {
		data.Email = sql.NullString{String: *info.Email(), Valid: true}
	}
	if info.Pwd() != nil {
		data.Password = sql.NullString{String: *info.Pwd(), Valid: true}
	}
	if info.Flags() != 0 {
		data.Flags = info.Flags()
	}
	if !info.BlockedUntil().IsZero() {
		data.Until = info.BlockedUntil()
	}

	data.Flags = info.Flags()

	return data
}

func fromDBEntity(info user_info.UserInfo) user.Info {
	data := user.Info{}

	if info.GetID() != "" {
		temp := info.GetID()
		data.SetId(&temp)
	}
	if info.GetUsername() != "" {
		temp := info.GetUsername()
		data.SetUsername(&temp)
	}
	if info.GetFullName() != "" {
		temp := info.GetFullName()
		data.SetFullname(&temp)
	}
	if info.GetEmail() != "" {
		temp := info.GetEmail()
		data.SetEmail(&temp)
	}
	if info.GetPasswordHash() != "" {
		temp := info.GetPasswordHash()
		data.SetPwd(&temp)
	}
	data.SetFlags(info.GetFlags())
	data.SetBlockedUntil(info.GetBlockedUntil())

	return data
}

func fromDBEntities(info []user_info.UserInfo) []user.Info {
	data := make([]user.Info, 0, len(info))

	for i := range info {
		data = append(data, fromDBEntity(info[i]))
	}

	return data
}
