package weights

import (
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure/weights"
	httpweights "neural_storage/cube/handlers/http/v1/entities/structure/weights"
)

func weightToBL(info httpweights.Info) weights.Info {
	linkWeights := []*weight.Info{}
	for _, lw := range info.Weights {
		linkWeights = append(linkWeights, weight.NewInfo(lw.ID, lw.LinkID, lw.Weight))
	}

	offsets := []*offset.Info{}
	for _, o := range info.Offsets {
		offsets = append(offsets, offset.NewInfo(o.ID, o.NeuronID, o.Offset))
	}

	return *weights.NewInfo(info.ID, info.Name, linkWeights, offsets)
}
