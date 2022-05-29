package offset

type Offset struct {
	InternalID string  `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	ID          string  `gorm:"-"`
	Weights     string  `gorm:"type:uuid;column:weights_info_id;"`
	Neuron      string  `gorm:"type:uuid;column:neuron_id;"`
	Offset      float64 `gorm:"column:value;"`
}

func (Offset) TableName() string {
	return "neuron_offsets"
}
