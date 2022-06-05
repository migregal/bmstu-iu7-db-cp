package offset

type Info struct {
	id       int
	neuronID int
	offset   float64
}

func NewInfo(id int, neuronID int, offset float64) *Info {
	return &Info{id: id, neuronID: neuronID, offset: offset}
}
