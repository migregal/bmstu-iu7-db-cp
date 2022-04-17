package neuron

func (i *Info) Id() string {
	return i.id
}

func (i *Info) SetId(id string) {
	i.id = id
}

func (i *Info) LayerID() string {
	return i.layerID
}

func (i *Info) SetLayerID(layerID string) {
	i.layerID = layerID
}
