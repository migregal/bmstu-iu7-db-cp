package user_info

import (
	"database/sql"
	"time"
)

type UserInfo struct {
	ID       string         `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	Username sql.NullString `gorm:"column:username;"`
	Email    sql.NullString `gorm:"column:email;" `
	FullName sql.NullString `gorm:"column:fullname;"`
	Password sql.NullString `gorm:"column:password_hash;"`
	Flags    uint64         `gorm:"column:flags;"`
	Until    time.Time      `gorm:"column:blocked;"`
}

func (UserInfo) TableName() string {
	return "users_info"
}
