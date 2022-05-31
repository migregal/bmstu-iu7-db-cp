package modelinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/model"
	dbmodel "neural_storage/database/core/entities/model"
	"neural_storage/database/core/entities/neuron"
	"neural_storage/database/core/entities/neuron/link"
	"neural_storage/database/core/entities/neuron/offset"
	"neural_storage/database/core/entities/structure"
	"neural_storage/database/core/entities/structure/layer"
	"neural_storage/database/core/entities/structure/weight"
	"neural_storage/database/core/entities/structure/weights"
)

func (r *Repository) Get(id string) (*model.Info, error) {
	var err error
	dbInfo := accumulatedModelInfo{}

	dbInfo.model, err = r.getModelInfo(id)
	if err != nil {
		return nil, err
	}

	structure, err := r.getStructInfo(id)
	if err != nil {
		return nil, err
	}
	dbInfo.structure = &structure

	dbInfo.layers, err = r.getLayersInfo(structure.GetID())
	if err != nil {
		return nil, err
	}

	dbInfo.neurons, err = r.getNeuronsInfo(dbInfo.layers)
	if err != nil {
		return nil, err
	}

	dbInfo.links, err = r.getNeuronLinksInfo(dbInfo.neurons)
	if err != nil {
		return nil, err
	}

	weights, err := r.getWeightsInfo(structure.GetID())
	if err != nil {
		return nil, err
	}

	dbInfo.weights, err = r.getDetailsWeightsInfo(weights)
	if err != nil {
		return nil, err
	}

	res := fromDBEntity(dbInfo)
	return &res, nil
}

func (r *Repository) getModelInfo(id string) (dbmodel.Model, error) {
	var modelInfo dbmodel.Model
	err := r.db.Where("id = ?", id).First(&modelInfo).Error
	if err != nil {
		return dbmodel.Model{}, fmt.Errorf("model get error: %w", err)
	}
	return modelInfo, nil
}

func (r *Repository) getStructInfo(id string) (structure.Structure, error) {
	var modelStruct structure.Structure
	err := r.db.Where("model_id = ?", id).First(&modelStruct).Error
	if err != nil {
		return structure.Structure{}, fmt.Errorf("strucutre get error: %w", err)
	}
	return modelStruct, nil
}

func (r *Repository) getLayersInfo(id string) ([]layer.Layer, error) {
	var structLayers []layer.Layer
	err := r.db.Where("structure_id = ?", id).Find(&structLayers).Error
	if err != nil {
		return nil, fmt.Errorf("strucutre layers get error: %w", err)
	}
	return structLayers, nil
}

func (r *Repository) getNeuronsInfo(layers []layer.Layer) ([]neuron.Neuron, error) {
	ids := []string{}
	for _, v := range layers {
		ids = append(ids, v.GetID())
	}

	var neurons []neuron.Neuron
	err := r.db.Find(&neurons, "layer_id in ?", ids).Error
	if err != nil {
		return nil, fmt.Errorf("neurons get error: %w", err)
	}
	return neurons, nil
}

func (r *Repository) getNeuronLinksInfo(neurons []neuron.Neuron) ([]link.Link, error) {
	ids := []string{}
	for _, v := range neurons {
		ids = append(ids, v.GetID())
	}

	var links []link.Link
	err := r.db.Where("from_id in ?", ids).Find(&links).Error
	if err != nil {
		return nil, fmt.Errorf("neuron links get error: %w", err)
	}
	return links, nil
}

func (r *Repository) getWeightsInfo(structID string) ([]weights.Weights, error) {
	var weights []weights.Weights
	err := r.db.Find(&weights, "structure_id = ?", structID).Error
	if err != nil {
		return nil, fmt.Errorf("weights info get error: %w", err)
	}
	return weights, nil
}

func (r *Repository) getDetailsWeightsInfo(weightsInfo []weights.Weights) ([]accumulatedWeightInfo, error) {
	var weightInfo []accumulatedWeightInfo
	for _, v := range weightsInfo {
		var offsets []offset.Offset
		err := r.db.Find(&offsets, "weights_info_id = ?", v.GetID()).Error
		if err != nil {
			return nil, fmt.Errorf("neuron offsets get error: %w", err)
		}

		var weight []weight.Weight
		err = r.db.Find(&weight, "weights_info_id = ?", v.GetID()).Error
		if err != nil {
			return nil, fmt.Errorf("neuron links weights get error: %w", err)
		}

		weightInfo = append(weightInfo,
			accumulatedWeightInfo{
				weightsInfo: &weights.Weights{
					InnerID:     v.GetID(),
					Name:        v.GetName(),
					StructureID: v.GetStructureID(),
				},
				offsets: offsets, weights: weight})
	}
	return weightInfo, nil
}
