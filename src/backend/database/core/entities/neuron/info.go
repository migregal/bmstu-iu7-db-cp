package neuron

type Neuron struct {
	ID       string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	NeuronID string `gorm:"type:uuid;column:neuron_id;"`
	LayerID  string `gorm:"type:uuid;column:layer_id"`
}
