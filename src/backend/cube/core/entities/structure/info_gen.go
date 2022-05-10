package structure

import (
	"neural_storage/cube/core/entities/neuron"
	"neural_storage/cube/core/entities/neuron/link"
	"neural_storage/cube/core/entities/structure/layer"
	"neural_storage/cube/core/entities/structure/weights"
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

func (i *Info) Neurons() []*neuron.Info {
	return i.neurons
}

func (i *Info) SetNeurons(neurons []*neuron.Info) {
	i.neurons = neurons
}

func (i *Info) Layers() []*layer.Info {
	return i.layers
}

func (i *Info) SetLayers(layers []*layer.Info) {
	i.layers = layers
}

func (i *Info) Links() []*link.Info {
	return i.links
}

func (i *Info) SetLinks(links []*link.Info) {
	i.links = links
}

func (i *Info) Weights() []*weights.Info {
	return i.weights
}

func (i *Info) SetWeights(weights []*weights.Info) {
	i.weights = weights
}
