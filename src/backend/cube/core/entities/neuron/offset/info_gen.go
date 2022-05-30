package offset

func (i *Info) ID() int {
	return i.id
}

func (i *Info) NeuronID() int {
	return i.neuronID
}

func (i *Info) Offset() float64 {
	return i.offset
}

func (i *Info) SetOffset(offset float64) {
	i.offset = offset
}
