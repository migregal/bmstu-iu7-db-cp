package layer

type Info struct {
	id             int
	limitFunc      string
	activationFunc string
}

func NewInfo(id int, limitFunc string, activationFunc string) *Info {
	return &Info{id: id, limitFunc: limitFunc, activationFunc: activationFunc}
}
