package userinfostat

import (
	"time"
)

type Info struct {
	ID        string    `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;"`
}

func (Info) TableName() string {
	return "users_info"
}
