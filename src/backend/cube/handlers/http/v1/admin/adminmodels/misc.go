package adminmodels

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/entities/structure/weights"
	httpmodel "neural_storage/cube/handlers/http/v1/entities/model"
	httpneuron "neural_storage/cube/handlers/http/v1/entities/neuron"
	httplink "neural_storage/cube/handlers/http/v1/entities/neuron/link"
	httpweight "neural_storage/cube/handlers/http/v1/entities/neuron/link/weight"
	httpoffset "neural_storage/cube/handlers/http/v1/entities/neuron/offset"
	httpstructure "neural_storage/cube/handlers/http/v1/entities/structure"
	httplayer "neural_storage/cube/handlers/http/v1/entities/structure/layer"
	httpweights "neural_storage/cube/handlers/http/v1/entities/structure/weights"
)

func modelFromBL(info *model.Info) httpmodel.Info {
	return httpmodel.Info{
		ID:        info.ID(),
		OwnerID:   info.OwnerID(),
		Name:      info.Name(),
		Structure: structFromBL(info.Structure()),
	}
}

func structFromBL(info *structure.Info) *httpstructure.Info {
	if info == nil {
		return nil
	}

	layers := []httplayer.Info{}
	for _, v := range info.Layers() {
		layers = append(layers, httplayer.Info{ID: v.ID(), ActivationFunc: v.ActivationFunc(), LimitFunc: v.LimitFunc()})
	}

	neurons := []httpneuron.Info{}
	for _, v := range info.Neurons() {
		neurons = append(neurons, httpneuron.Info{ID: v.ID(), LayerID: v.LayerID()})
	}

	links := []httplink.Info{}
	for _, v := range info.Links() {
		links = append(links, httplink.Info{ID: v.ID(), From: v.From(), To: v.To()})
	}

	return &httpstructure.Info{
		ID:      info.ID(),
		Name:    info.Name(),
		Layers:  layers,
		Neurons: neurons,
		Links:   links,
		Weights: weightFromBL(info.Weights()),
	}
}

func weightFromBL(info []*weights.Info) []httpweights.Info {
	weights := []httpweights.Info{}
	for _, i := range info {
		linkWeights := []httpweight.Info{}
		for _, lw := range i.Weights() {
			linkWeights = append(linkWeights, httpweight.Info{ID: lw.ID(), LinkID: lw.LinkID(), Weight: lw.Weight()})
		}

		offsets := []httpoffset.Info{}
		for _, o := range i.Offsets() {
			offsets = append(offsets, httpoffset.Info{ID: o.ID(), NeuronID: o.NeuronID(), Offset: o.Offset()})
		}

		weights = append(weights, httpweights.Info{ID: i.ID(), Name: i.Name(), Weights: linkWeights, Offsets: offsets})
	}

	return weights
}
