package link

func (i *Link) GetStructureID() string {
	return i.Structure
}

func (i *Link) GetID() string {
	return i.ID
}

func (i *Link) GetFrom() string {
	return i.From
}

func (i *Link) GetTo() string {
	return i.To
}
