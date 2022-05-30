package weight

type Info struct {
	id     int
	linkID int
	weight float64
}

func NewInfo(id int, linkID int, weight float64) *Info {
	return &Info{id: id, linkID: linkID, weight: weight}
}
