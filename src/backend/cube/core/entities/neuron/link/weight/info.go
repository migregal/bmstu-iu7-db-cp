package weight

type Info struct {
	id     string
	linkID string
	weight float64
}

func NewInfo(id string, linkID string, weight float64) *Info {
	return &Info{id: id, linkID: linkID, weight: weight}
}
