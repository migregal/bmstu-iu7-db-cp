package weights

import (
	"time"
)

type Weights struct {
	ID          int `gorm:"-"`
	InnerID     string `gorm:"type:uuid;column:id;default:generated();"`
	StructureID string `gorm:"type:uuid;column:structure_id;"`
	Name        string `gorm:"column:name"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Weights) TableName() string {
	return "weights_info"
}
