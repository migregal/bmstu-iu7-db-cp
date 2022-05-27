package modelstat

import "time"

type Info struct {
	ID        string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Info) TableName() string {
	return "models"
}
