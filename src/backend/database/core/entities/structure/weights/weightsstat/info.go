package weightsstat

import "time"

type Info struct {
	ID        string `gorm:"type:uuid;column:id;default:generated();"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Info) TableName() string {
	return "weights_info"
}
