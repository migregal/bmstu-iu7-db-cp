package link

type Info struct {
	id   int
	from int
	to   int
}

func NewInfo(id int, from int, to int) *Info {
	return &Info{id: id, from: from, to: to}
}
