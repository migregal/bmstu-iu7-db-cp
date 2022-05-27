package weight

func (w *Weight) GetID() string {
	return w.ID
}

func (w *Weight) GetWeightsID() string {
	return w.WeightsID
}

func (w *Weight) GetLinkID() string {
	return w.LinkID
}

func (w *Weight) GetValue() float64 {
	return w.Value
}
