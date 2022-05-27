package offset

type Offset struct {
	ID      string  `gorm:"primaryKey;type:uuid;column:id;"`
	Weights string  `gorm:"type:uuid;column:weights_id;"`
	Neuron  string  `gorm:"type:uuid;column:neuron_id;"`
	Offset  float64 `gorm:"column:value;"`
}

func (Offset) TableName() string {
	return "neuron_offsets"
}
