package modelinfo

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/neuron"
	"neural_storage/cube/core/entities/neuron/link"
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/entities/structure/layer"
	"neural_storage/cube/core/entities/structure/weights"
	dbmodel "neural_storage/database/core/entities/model"
	dbneuron "neural_storage/database/core/entities/neuron"
	dblink "neural_storage/database/core/entities/neuron/link"
	dboffset "neural_storage/database/core/entities/neuron/offset"
	dbstructure "neural_storage/database/core/entities/structure"
	dblayer "neural_storage/database/core/entities/structure/layer"
	dbweight "neural_storage/database/core/entities/structure/weight"
	dbweights "neural_storage/database/core/entities/structure/weights"
)

type accumulatedWeightInfo struct {
	weightsInfo *dbweights.Weights
	weights     []dbweight.Weight
	offsets     []dboffset.Offset
}

type accumulatedModelInfo struct {
	model     dbmodel.Model
	structure *dbstructure.Structure
	layers    []dblayer.Layer
	neurons   []dbneuron.Neuron
	links     []dblink.Link
	weights   []accumulatedWeightInfo
}

func toDBEntity(info model.Info) accumulatedModelInfo {
	data := accumulatedModelInfo{}

	data.model = dbmodel.Model{ID: info.ID(), OwnerID: info.OwnerID(), Name: info.Name()}

	if info.Structure() == nil {
		return data
	}

	data.structure = &dbstructure.Structure{ID: info.Structure().ID(), ModelID: info.ID(), Name: info.Structure().Name()}

	if len(info.Structure().Layers()) > 0 {
		var layers []dblayer.Layer
		for _, v := range info.Structure().Layers() {
			layers = append(layers, dblayer.Layer{
				ID:             v.ID(),
				LimitFunc:      v.LimitFunc(),
				ActivationFunc: v.ActivationFunc(),
			})
		}
		data.layers = layers
	}

	if len(info.Structure().Neurons()) > 0 {
		var neurons []dbneuron.Neuron
		for _, v := range info.Structure().Neurons() {
			neurons = append(neurons, dbneuron.Neuron{
				ID:      v.ID(),
				LayerID: v.ID(),
			})
		}
		data.neurons = neurons
	}

	if len(info.Structure().Links()) > 0 {
		var links []dblink.Link
		for _, v := range info.Structure().Links() {
			links = append(links, dblink.Link{
				ID:        v.ID(),
				From:      v.From(),
				To:        v.To(),
			})
		}
		data.links = links
	}

	if len(info.Structure().Weights()) > 0 {
		var weights []accumulatedWeightInfo
		for _, w := range info.Structure().Weights() {
			temp := accumulatedWeightInfo{}
			temp.weightsInfo = &dbweights.Weights{
				ID:          w.ID(),
				Name:        w.Name(),
				StructureID: info.Structure().ID(),
			}
			for _, v := range w.Weights() {
				temp.weights = append(temp.weights, dbweight.Weight{
					ID:        v.ID(),
					LinkID:    v.LinkID(),
					WeightsID: w.ID(),
					Value:     v.Weight(),
				})
			}

			for _, o := range w.Offsets() {
				temp.offsets = append(temp.offsets, dboffset.Offset{
					ID:      o.ID(),
					Neuron:  o.NeuronID(),
					Weights: w.ID(),
					Offset:  o.Offset(),
				})
			}

			weights = append(weights, temp)
		}
		data.weights = weights
	}

	return data
}

func fromDBEntity(info accumulatedModelInfo) model.Info {
	var links []*link.Info
	for i := range info.links {
		links = append(
			links,
			link.NewInfo(
				info.links[i].GetID(),
				info.links[i].GetFrom(),
				info.links[i].GetTo(),
			),
		)
	}

	var neurons []*neuron.Info
	for i := range info.neurons {
		neurons = append(
			neurons,
			neuron.NewInfo(info.neurons[i].GetID(), info.neurons[i].GetLayerID()))
	}

	var layers []*layer.Info
	for _, v := range info.layers {
		layers = append(
			layers,
			layer.NewInfo(v.GetID(), v.GetLimitFunc(), v.GetActivationFunc()))
	}

	var wholeWeightsInfo []*weights.Info
	for _, w := range info.weights {
		var offsets []*offset.Info
		for _, v := range w.offsets {
			offsets = append(offsets, offset.NewInfo(v.GetID(), v.GetNeuronID(), v.GetValue()))
		}
		var linkWeights []*weight.Info
		for _, v := range w.weights {
			linkWeights = append(linkWeights, weight.NewInfo(v.GetID(), v.GetLinkID(), v.GetValue()))
		}
		var info *weights.Info
		if w.weightsInfo != nil {
			info = weights.NewInfo(w.weightsInfo.GetID(), w.weightsInfo.GetName(), linkWeights, offsets)
		} else {
			info = weights.NewInfo("", "", linkWeights, offsets)
		}
		wholeWeightsInfo = append(wholeWeightsInfo, info)
	}

	structureInfo := structure.NewInfo(
		info.structure.GetID(),
		info.structure.GetName(),
		neurons,
		layers,
		links,
		wholeWeightsInfo)

	return *model.NewInfo(info.model.GetID(), info.model.GetOwnerID(), info.model.GetName(), structureInfo)
}
