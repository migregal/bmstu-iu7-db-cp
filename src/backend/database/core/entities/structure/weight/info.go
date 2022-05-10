package weight

type Weight struct {
	ID        string  `gorm:"type:uuid;column:id;"`
	LinkID    string  `gorm:"type:uuid;column:link_id;"`
	WeightsID string  `gorm:"type:uuid;column:weights_id;"`
	Value     float64 `gorm:"column:value;"`
}
