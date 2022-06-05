package weights

import (
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
)

type Info struct {
	id      string
	name    string
	weights []*weight.Info
	offsets []*offset.Info
}

func NewInfo(id, name string, weights []*weight.Info, offsets []*offset.Info) *Info {
	return &Info{id: id, name: name, weights: weights, offsets: offsets}
}
