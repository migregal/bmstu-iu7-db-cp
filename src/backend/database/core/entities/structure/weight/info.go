package weight

type Weight struct {
	ID        string  `gorm:"-"`
	InnerID   string  `gorm:"type:uuid;column:id;default:generated();"`
	LinkID    string  `gorm:"type:uuid;column:link_id;"`
	WeightsID string  `gorm:"type:uuid;column:weights_info_id;"`
	Value     float64 `gorm:"column:value;"`
}

func (Weight) TableName() string {
	return "link_weights"
}
