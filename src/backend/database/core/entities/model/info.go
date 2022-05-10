package model

type Model struct {
	ID   string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	Name string `gorm:"column:name;"`
}
