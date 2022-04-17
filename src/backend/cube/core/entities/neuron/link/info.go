package link

type Info struct {
	id   string
	from string
	to   string
}

func NewInfo(id string, from string, to string) *Info {
	return &Info{id: id, from: from, to: to}
}
