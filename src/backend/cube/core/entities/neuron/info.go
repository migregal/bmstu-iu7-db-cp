package neuron

type Info struct {
	id      int
	layerID int
}

func NewInfo(id int, layerID int) *Info {
	return &Info{id: id, layerID: layerID}
}
