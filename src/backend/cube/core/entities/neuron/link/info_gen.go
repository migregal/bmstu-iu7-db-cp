package link

func (i *Info) ID() int {
	return i.id
}

func (i *Info) From() int {
	return i.from
}

func (i *Info) SetFrom(from int) {
	i.from = from
}

func (i *Info) To() int {
	return i.to
}

func (i *Info) SetTo(to int) {
	i.to = to
}
