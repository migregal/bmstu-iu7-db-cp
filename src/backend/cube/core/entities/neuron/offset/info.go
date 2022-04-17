package offset

type Info struct {
	id     string
	offset float64
}

func NewInfo(id string, offset float64) *Info {
	return &Info{id: id, offset: offset}
}
