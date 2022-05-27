package weights

import (
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
)

func (i *Info) ID() string {
	return i.id
}

func (i *Info) Name() string {
	return i.name
}

func (i *Info) SetName(name string) {
	i.name = name
}

func (i *Info) Weights() []*weight.Info {
	return i.weights
}

func (i *Info) SetWeights(weights []*weight.Info) {
	i.weights = weights
}

func (i *Info) Offsets() []*offset.Info {
	return i.offsets
}

func (i *Info) SetOffsets(offsets []*offset.Info) {
	i.offsets = offsets
}
