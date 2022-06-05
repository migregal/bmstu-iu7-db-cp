package weights

func (w *Weights) GetID() string {
	return w.InnerID
}

func (w *Weights) GetStructureID() string {
	return w.StructureID
}

func (w *Weights) GetName() string {
	return w.Name
}
