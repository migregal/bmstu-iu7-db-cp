package weight

type Weight struct {
	ID             int     `gorm:"-"`
	LinkID         int     `gorm:"-"`
	WeightsID      int  `gorm:"-"`
	InnerID        string  `gorm:"type:uuid;column:id;default:generated();"`
	InnerLinkID    string  `gorm:"type:uuid;column:link_id;"`
	InnerWeightsID string  `gorm:"type:uuid;column:weights_info_id;"`
	Value          float64 `gorm:"column:value;"`
}

func (Weight) TableName() string {
	return "link_weights"
}
