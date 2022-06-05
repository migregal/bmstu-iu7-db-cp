package weight

func (w *Weight) GetID() string {
	return w.InnerID
}

func (w *Weight) GetWeightsID() string {
	return w.InnerWeightsID
}

func (w *Weight) GetLinkID() string {
	return w.InnerLinkID
}

func (w *Weight) GetValue() float64 {
	return w.Value
}
