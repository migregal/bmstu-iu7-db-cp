package user_creds

import "database/sql"

type UserCreds struct {
	ID           string         `gorm:"primaryKey;type:uuid;column:user_id;default:generated();"`
	Uname        sql.NullString `gorm:"column:username;"`
	EmailAddress sql.NullString `gorm:"column:email;"`
	PasswordHash sql.NullString `gorm:"column:password_hash;"`
}

func (UserCreds) TableName() string {
	return "users_creds"
}
