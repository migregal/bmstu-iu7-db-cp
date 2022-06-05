package modelstructweightsinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/database/core/entities/neuron"
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

	neurons, err := r.getNeuronsInfoByOffsets(dbInfo.offsets)
	if err != nil {
		return nil, err
	}

	neuronsMap := map[string]int{}
	for i := range neurons {
		neuronsMap[neurons[i].InnerID] = neurons[i].ID
	}

	links, err := r.getNeuronLinksInfo(neurons)
	if err != nil {
		return nil, err
	}
	neurons = nil

	linksMap := map[string]int{}
	for i := range links {
		linksMap[links[i].InnerID] = links[i].ID
	}
	links = nil

	return fromDBEntityStructured(neuronsMap, linksMap, dbInfo), nil
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
	err := r.db.Find(&offsets, "weights_info_id = ?", info.GetID()).Error
	if err != nil {
		return accumulatedWeightInfo{}, fmt.Errorf("neuron offsets get error: %w", err)
	}

	var weight []weight.Weight
	err = r.db.Find(&weight, "weights_info_id = ?", info.GetID()).Error
	if err != nil {
		return accumulatedWeightInfo{}, fmt.Errorf("neuron links weights get error: %w", err)
	}

	return accumulatedWeightInfo{
			weightsInfo: &dbweights.Weights{
				InnerID:     info.GetID(),
				Name:        info.GetName(),
				StructureID: info.GetStructureID(),
			},
			offsets: offsets,
			weights: weight,
		},
		nil
}

func (r *Repository) getNeuronsInfoByOffsets(offsets []offset.Offset) ([]neuron.Neuron, error) {
	ids := []string{}
	for _, v := range offsets {
		ids = append(ids, v.GetNeuronID())
	}

	var neurons []neuron.Neuron
	err := r.db.Find(&neurons, "id in ?", ids).Error
	if err != nil {
		return nil, fmt.Errorf("neurons get error: %w", err)
	}
	return neurons, nil
}
