package structure

type Structure struct {
	ID      string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	ModelID string `gorm:"type:uuid;column:model_id;"`
	Name    string `gorm:"column:title"`
}
