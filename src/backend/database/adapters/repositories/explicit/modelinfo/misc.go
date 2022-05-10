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

func extractWeightsIDs(neurons []dbweights.Weights) []string {
	var res []string

	for i := range neurons {
		res = append(res, neurons[i].ID)
	}

	return res
}

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

	data.model = dbmodel.Model{ID: info.Id(), Name: info.Name()}

	if info.Structure() != nil {
		data.structure = &dbstructure.Structure{ID: info.Structure().ID(), Name: info.Structure().Name()}
	}

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
				NeuronID: v.Id(),
				LayerID:  v.Id(),
			})
		}
		data.neurons = neurons
	}

	if len(info.Structure().Links()) > 0 {
		var links []dblink.Link
		for _, v := range info.Structure().Links() {
			links = append(links, dblink.Link{
				StructureID: data.structure.ID,
				LinkID:      v.Id(),
				FromID:      v.From(),
				ToID:        v.To(),
			})
		}
		data.links = links
	}

	if len(info.Structure().Weights()) > 0 {
		var weights []accumulatedWeightInfo
		for _, w := range info.Structure().Weights() {
			temp := accumulatedWeightInfo{}
			temp.weightsInfo = &dbweights.Weights{ID: w.Id(), Name: w.Name()}
			for _, v := range w.Weights() {
				temp.weights = append(temp.weights, dbweight.Weight{
					LinkID:    v.ID(),
					WeightsID: v.LinkID(),
					Value:     v.Weight(),
				})
			}

			for _, o := range w.Offsets() {
				temp.offsets = append(temp.offsets, dboffset.Offset{
					NeuronID:  o.ID(),
					WeightsID: o.WeightID(),
					Value:     o.Offset(),
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
				info.links[i].LinkID,
				info.links[i].FromID,
				info.links[i].ToID,
			),
		)
	}

	var neurons []*neuron.Info
	for i := range info.neurons {
		neurons = append(
			neurons,
			neuron.NewInfo(info.neurons[i].NeuronID, info.neurons[i].LayerID))
	}

	var layers []*layer.Info
	for _, v := range info.layers {
		layers = append(
			layers,
			layer.NewInfo(v.ID, v.LimitFunc, v.ActivationFunc))
	}

	var wholeWeightsInfo []*weights.Info
	for _, w := range info.weights {
		var offsets []*offset.Info
		for _, v := range w.offsets {
			offsets = append(offsets, offset.NewInfo(v.WeightsID, v.ID, v.Value))
		}
		var linkWeights []*weight.Info
		for _, v := range w.weights {
			linkWeights = append(linkWeights, weight.NewInfo(v.ID, v.LinkID, v.Value))
		}
		var info *weights.Info
		if w.weightsInfo != nil {
			info = weights.NewInfo(w.weightsInfo.Name, linkWeights, offsets)
		} else {
			info = weights.NewInfo("", linkWeights, offsets)
		}
		wholeWeightsInfo = append(wholeWeightsInfo, info)
	}

	structureInfo := structure.NewInfo(
		info.structure.Name,
		neurons,
		layers,
		links,
		wholeWeightsInfo)

	return *model.NewInfo(info.model.Name, structureInfo)
}
