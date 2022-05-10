package weight

func (i *Info) ID() string {
	return i.id
}

func (i *Info) LinkID() string {
	return i.linkID
}

func (i *Info) Weight() float64 {
	return i.weight
}

func (i *Info) SetWeight(weight float64) {
	i.weight = weight
}
