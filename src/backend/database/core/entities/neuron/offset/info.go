package offset

type Offset struct {
	ID        string  `gorm:"type:uuid;column:id;"`
	NeuronID  string  `gorm:"type:uuid;column:neuron_id;"`
	WeightsID string  `gorm:"type:uuid;column:weights_id;"`
	Value     float64 `gorm:"column:value;"`
}
