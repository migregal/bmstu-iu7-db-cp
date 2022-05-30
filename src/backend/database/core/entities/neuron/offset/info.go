package offset

type Offset struct {
	ID           int     `gorm:"-"`
	Weights      int     `gorm:"-"`
	Neuron       int     `gorm:"-"`
	InternalID   string  `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	InnerWeights string  `gorm:"type:uuid;column:weights_info_id;"`
	InnerNeuron  string  `gorm:"type:uuid;column:neuron_id;"`
	Offset       float64 `gorm:"column:value;"`
}

func (Offset) TableName() string {
	return "neuron_offsets"
}
