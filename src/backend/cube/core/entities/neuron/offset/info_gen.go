package offset

func (i *Info) ID() string {
	return i.id
}

func (i *Info) NeuronID() string {
	return i.neuronID
}

func (i *Info) WeightID() string {
	return i.weightId
}

func (i *Info) Offset() float64 {
	return i.offset
}

func (i *Info) SetOffset(offset float64) {
	i.offset = offset
}
