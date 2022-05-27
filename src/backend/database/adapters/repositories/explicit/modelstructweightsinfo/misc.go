package modelstructweightsinfo

import (
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure/weights"
	dboffset "neural_storage/database/core/entities/neuron/offset"
	dbweight "neural_storage/database/core/entities/structure/weight"
	dbweights "neural_storage/database/core/entities/structure/weights"
)

type accumulatedWeightInfo struct {
	weightsInfo *dbweights.Weights
	weights     []dbweight.Weight
	offsets     []dboffset.Offset
}

func toDBEntity(structureID string, info []weights.Info) []accumulatedWeightInfo {
	var weights []accumulatedWeightInfo
	for _, w := range info {
		temp := accumulatedWeightInfo{}
		temp.weightsInfo = &dbweights.Weights{
			ID:          w.ID(),
			Name:        w.Name(),
			StructureID: structureID,
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
				Neuron:  o.ID(),
				Weights: w.ID(),
				Offset:  o.Offset(),
			})
		}

		weights = append(weights, temp)
	}
	return weights
}

func fromDBEntity(info accumulatedWeightInfo) *weights.Info {
	var offsets []*offset.Info
	for _, v := range info.offsets {
		offsets = append(offsets, offset.NewInfo(v.GetID(), v.GetNeuronID(), v.GetValue()))
	}
	var linkWeights []*weight.Info
	for _, v := range info.weights {
		linkWeights = append(linkWeights,
			weight.NewInfo(v.GetID(), v.GetLinkID(), v.GetValue()))
	}
	var wInfo = weights.NewInfo(
		info.weightsInfo.GetID(),
		info.weightsInfo.GetName(),
		linkWeights,
		offsets,
	)

	return wInfo
}
