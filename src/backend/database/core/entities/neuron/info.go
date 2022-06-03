package neuron

type Neuron struct {
	ID           int    `gorm:"column:neuron_id;"`
	LayerID      int    `gorm:"-"`
	InnerID      string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	InnerLayerID string `gorm:"type:uuid;column:layer_id"`
}
