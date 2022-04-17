package layer

type Info struct {
	id             string
	limitFunc      string
	activationFunc string
}

func NewInfo(id string, limitFunc string, activationFunc string) *Info {
	return &Info{id: id, limitFunc: limitFunc, activationFunc: activationFunc}
}
