package adminweights

import (
	"neural_storage/cube/core/entities/structure/weights"
	httpweight "neural_storage/cube/handlers/http/v1/entities/neuron/link/weight"
	httpoffset "neural_storage/cube/handlers/http/v1/entities/neuron/offset"
	httpweights "neural_storage/cube/handlers/http/v1/entities/structure/weights"
)

func weightFromBL(info weights.Info) httpweights.Info {
	linkWeights := []httpweight.Info{}
	for _, lw := range info.Weights() {
		linkWeights = append(linkWeights, httpweight.Info{ID: lw.ID(), LinkID: lw.LinkID(), Weight: lw.Weight()})
	}

	offsets := []httpoffset.Info{}
	for _, o := range info.Offsets() {
		offsets = append(offsets, httpoffset.Info{ID: o.ID(), NeuronID: o.NeuronID(), Offset: o.Offset()})
	}

	return httpweights.Info{ID: info.ID(), Name: info.Name(), Weights: linkWeights, Offsets: offsets}
}
