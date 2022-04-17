package weight

type Info struct {
	id     string
	weight float64
}

func NewInfo(id string, weight float64) *Info {
	return &Info{id: id, weight: weight}
}
