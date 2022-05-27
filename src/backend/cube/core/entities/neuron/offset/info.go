package offset

type Info struct {
	id       string
	neuronID string
	offset   float64
}

func NewInfo(id string, neuronID string, offset float64) *Info {
	return &Info{id: id, neuronID: neuronID, offset: offset}
}
