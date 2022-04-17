package offset

func (i *Info) Id() string {
	return i.id
}

func (i *Info) Offset() float64 {
	return i.offset
}

func (i *Info) SetOffset(offset float64) {
	i.offset = offset
}
