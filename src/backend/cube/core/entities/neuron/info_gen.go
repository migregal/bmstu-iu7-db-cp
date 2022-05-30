package neuron

func (i *Info) ID() int {
	return i.id
}

func (i *Info) SetId(id int) {
	i.id = id
}

func (i *Info) LayerID() int {
	return i.layerID
}

func (i *Info) SetLayerID(layerID int) {
	i.layerID = layerID
}
