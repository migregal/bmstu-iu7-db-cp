package offset

type Info struct {
	id       string
	neuronID string
	weightId string
	offset   float64
}

func NewInfo(weightId string, id string, offset float64) *Info {
	return &Info{weightId: weightId, id: id, offset: offset}
}
