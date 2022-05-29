package layer

func (l *Layer) GetID() string {
	return l.InnerID
}

func (l *Layer) GetStructID() string {
	return l.StructureID
}

func (l *Layer) GetActivationFunc() string {
	return l.ActivationFunc
}

func (l *Layer) GetLimitFunc() string {
	return l.LimitFunc
}
