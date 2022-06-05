package structure

import (
	"neural_storage/cube/core/entities/neuron"
	"neural_storage/cube/core/entities/neuron/link"
	"neural_storage/cube/core/entities/structure/layer"
	"neural_storage/cube/core/entities/structure/weights"
)

type Info struct {
	id      string
	name    string
	neurons []*neuron.Info
	layers  []*layer.Info
	links   []*link.Info
	weights []*weights.Info
}

func NewInfo(id, name string, neurons []*neuron.Info, layers []*layer.Info, links []*link.Info, weights []*weights.Info) *Info {
	return &Info{id: id, name: name, neurons: neurons, layers: layers, links: links, weights: weights}
}
