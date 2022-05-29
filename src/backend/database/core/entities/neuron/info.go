package neuron

type Neuron struct {
	ID      string `gorm:"-"`
	InnerID string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	LayerID string `gorm:"type:uuid;column:layer_id"`
}
