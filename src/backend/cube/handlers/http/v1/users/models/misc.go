package models

import (
	"neural_storage/cube/core/entities/neuron"
	"neural_storage/cube/core/entities/neuron/link"
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/entities/structure/layer"
	"neural_storage/cube/core/entities/structure/weights"
	httpstructure "neural_storage/cube/handlers/http/v1/entities/structure"
	httpweights "neural_storage/cube/handlers/http/v1/entities/structure/weights"
)

func structToBL(info httpstructure.Info) *structure.Info {
	neurons := []*neuron.Info{}
	for _, v := range info.Neurons {
		neurons = append(neurons, neuron.NewInfo(v.ID, v.LayerID))
	}

	layers := []*layer.Info{}
	for _, v := range info.Layers {
		layers = append(layers, layer.NewInfo(v.ID, v.LimitFunc, v.ActivationFunc))
	}

	links := []*link.Info{}
	for _, v := range info.Links {
		links = append(links, link.NewInfo(v.ID, v.From, v.To))
	}

	weights := []*weights.Info{}
	for _, w := range info.Weights {
		weights = append(weights, weightToBL(w))
	}
	return structure.NewInfo(info.ID, neurons, layers, links, weights)
}

func weightToBL(info httpweights.Info) *weights.Info {
	linkWeights := []*weight.Info{}
	for _, lw := range info.Weights {
		linkWeights = append(linkWeights, weight.NewInfo(lw.ID, lw.LinkID, lw.Weight))
	}

	offsets := []*offset.Info{}
	for _, o := range info.Offsets {
		offsets = append(offsets, offset.NewInfo(o.ID, o.NeuronID, o.Offset))
	}

	return weights.NewInfo(info.ID, info.Name, linkWeights, offsets)
}
