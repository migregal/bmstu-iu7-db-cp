package offset

func (i *Offset) GetID() string {
	return i.InternalID
}

func (i *Offset) GetWeightsID() string {
	return i.InnerWeights
}

func (i *Offset) GetNeuronID() string {
	return i.InnerNeuron
}

func (i *Offset) GetValue() float64 {
	return i.Offset
}
