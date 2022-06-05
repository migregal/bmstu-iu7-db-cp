package neuron

func (i *Neuron) GetID() string {
	return i.InnerID
}

func (i *Neuron) GetLayerID() string {
	return i.InnerLayerID
}
