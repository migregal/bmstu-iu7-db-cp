package user_info

import "database/sql"

type UserInfo struct {
	ID           string         `gorm:"primaryKey;type:uuid;column:user_id;default:generated();"`
	Uname        sql.NullString `gorm:"column:username;"`
	EmailAddress sql.NullString `gorm:"column:email;" `
	FName        sql.NullString `gorm:"column:fullname;"`
}

func (UserInfo) TableName() string {
	return "users_info"
}
