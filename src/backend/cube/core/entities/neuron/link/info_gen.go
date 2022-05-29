package link

func (i *Info) ID() string {
	return i.id
}

func (i *Info) From() string {
	return i.from
}

func (i *Info) SetFrom(from string) {
	i.from = from
}

func (i *Info) To() string {
	return i.to
}

func (i *Info) SetTo(to string) {
	i.to = to
}
