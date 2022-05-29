package userinfo

import (
	"database/sql"
	"neural_storage/cube/core/entities/user"
	"neural_storage/database/core/entities/user_info"
)

func toDBEntity(info user.Info) user_info.UserInfo {
	data := user_info.UserInfo{}

	if info.ID() != "" {
		data.ID = info.ID()
	}
	if info.Username() != "" {
		data.Username = sql.NullString{String: info.Username(), Valid: true}
	}
	if info.Fullname() != "" {
		data.FullName = sql.NullString{String: info.Fullname(), Valid: true}
	}
	if info.Email() != "" {
		data.Email = sql.NullString{String: info.Email(), Valid: true}
	}
	if info.Pwd() != "" {
		data.Password = sql.NullString{String: info.Pwd(), Valid: true}
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
		data.SetId(info.GetID())
	}
	if info.GetUsername() != "" {
		data.SetUsername(info.GetUsername())
	}
	if info.GetFullName() != "" {
		data.SetFullname(info.GetFullName())
	}
	if info.GetEmail() != "" {
		data.SetEmail(info.GetEmail())
	}
	if info.GetPasswordHash() != "" {
		data.SetPwd(info.GetPasswordHash())
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
