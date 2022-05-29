package modelstructweightsinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/database/core/entities/neuron/offset"
	"neural_storage/database/core/entities/structure/weight"
	dbweights "neural_storage/database/core/entities/structure/weights"
)

func (r *Repository) Get(weightsId string) (*weights.Info, error) {
	var err error

	weights, err := r.getWeightsInfo(weightsId)
	if err != nil {
		return nil, err
	}

	dbInfo, err := r.getDetailsWeightsInfo(weights)
	if err != nil {
		return nil, err
	}

	return fromDBEntity(dbInfo), nil
}

func (r *Repository) getWeightsInfo(id string) (dbweights.Weights, error) {
	var info dbweights.Weights
	err := r.db.First(&info, "id = ?", id).Error
	if err != nil {
		return dbweights.Weights{}, fmt.Errorf("weights info get error: %w", err)
	}
	return info, nil
}

func (r *Repository) getDetailsWeightsInfo(info dbweights.Weights) (accumulatedWeightInfo, error) {
	var offsets []offset.Offset
	err := r.db.Find(&offsets, "weights_id = ?", info.GetID()).Error
	if err != nil {
		return accumulatedWeightInfo{}, fmt.Errorf("neuron offsets get error: %w", err)
	}

	var weight []weight.Weight
	err = r.db.Find(&weight, "weights_id = ?", info.GetID()).Error
	if err != nil {
		return accumulatedWeightInfo{}, fmt.Errorf("neuron links weights get error: %w", err)
	}

	return accumulatedWeightInfo{
			weightsInfo: &dbweights.Weights{
				ID:          info.GetID(),
				Name:        info.GetName(),
				StructureID: info.GetStructureID(),
			},
			offsets: offsets,
			weights: weight,
		},
		nil
}
