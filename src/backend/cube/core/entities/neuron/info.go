package neuron

type Info struct {
	id      string
	layerID string
}

func NewInfo(id string, layerID string) *Info {
	return &Info{id: id, layerID: layerID}
}
