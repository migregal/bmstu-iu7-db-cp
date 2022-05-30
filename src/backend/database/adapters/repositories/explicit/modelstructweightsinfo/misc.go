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
	for i, w := range info {
		temp := accumulatedWeightInfo{}
		temp.weightsInfo = &dbweights.Weights{
			ID:          i,
			InnerID:     w.ID(),
			Name:        w.Name(),
			StructureID: structureID,
		}
		for _, v := range w.Weights() {
			temp.weights = append(temp.weights, dbweight.Weight{
				ID:             v.ID(),
				LinkID:         v.LinkID(),
				WeightsID:      i,
				InnerWeightsID: w.ID(),
				Value:          v.Weight(),
			})
		}

		for _, o := range w.Offsets() {
			temp.offsets = append(temp.offsets, dboffset.Offset{
				Neuron:       o.ID(),
				Weights:      i,
				InnerWeights: w.ID(),
				Offset:       o.Offset(),
			})
		}

		weights = append(weights, temp)
	}
	return weights
}

func fromDBEntity(info accumulatedWeightInfo) *weights.Info {
	var offsets []*offset.Info
	for i, v := range info.offsets {
		offsets = append(offsets, offset.NewInfo(i, v.Neuron, v.GetValue()))
	}
	var linkWeights []*weight.Info
	for i, v := range info.weights {
		linkWeights = append(linkWeights,
			weight.NewInfo(i, v.LinkID, v.GetValue()))
	}
	var wInfo = weights.NewInfo(
		info.weightsInfo.GetID(),
		info.weightsInfo.GetName(),
		linkWeights,
		offsets,
	)

	return wInfo
}
