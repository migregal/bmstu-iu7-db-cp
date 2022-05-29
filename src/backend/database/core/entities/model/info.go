package model

type Model struct {
	ID      string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	OwnerID string `gorm:"column:owner_id;"`
	Name    string `gorm:"column:title;"`
}

func (Model) TableName() string {
	return "models"
}
