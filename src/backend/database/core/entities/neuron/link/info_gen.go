package link


func (i *Link) GetID() string {
	return i.InnerID
}

func (i *Link) GetFrom() string {
	return i.InnerFrom
}

func (i *Link) GetTo() string {
	return i.InnerTo
}
