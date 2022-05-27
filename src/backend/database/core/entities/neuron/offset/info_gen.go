package offset

func (i *Offset) GetID() string {
	return i.ID
}

func (i *Offset) GetWeightsID() string {
	return i.Weights
}

func (i *Offset) GetNeuronID() string {
	return i.Neuron
}

func (i *Offset) GetValue() float64 {
	return i.Offset
}
