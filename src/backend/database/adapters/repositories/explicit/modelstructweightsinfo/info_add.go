package modelstructweightsinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/database/core/entities/neuron"
	"neural_storage/database/core/entities/neuron/link"
	"neural_storage/database/core/entities/structure"
	"neural_storage/database/core/entities/structure/layer"
	"neural_storage/database/core/services/interactor/database"
)

func (r *Repository) Add(structID string, info []weights.Info) ([]string, error) {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	layers, err := r.getLayersInfo(structID)
	if err != nil {
		return nil, err
	}

	neurons, err := r.getNeuronsInfo(layers)
	if err != nil {
		return nil, err
	}
	layers = nil

	neuronsMap := map[int]string{}
	for i := range neurons {
		neuronsMap[neurons[i].ID] = neurons[i].InnerID
	}

	links, err := r.getNeuronLinksInfo(neurons)
	if err != nil {
		return nil, err
	}
	neurons = nil

	linksMap := map[int]string{}
	for i := range links {
		linksMap[links[i].ID] = links[i].InnerID
	}
	links = nil

	data := toDBEntityStructured(structID, neuronsMap, linksMap, info)

	ids, err := r.createWeightsInfoTransact(database.Interactor{DB: tx}, data)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return ids, tx.Commit().Error
}

func (r *Repository) createWeightsInfoTransact(tx database.Interactor, info []accumulatedWeightInfo) ([]string, error) {
	ids := []string{}
	for _, v := range info {
		err := tx.Create(&v.weightsInfo).Error
		if err != nil {
			return nil, fmt.Errorf("create model weights info: %w", err)
		}
		ids = append(ids, v.weightsInfo.InnerID)
		for _, o := range v.offsets {
			o.InnerWeights = v.weightsInfo.InnerID
			if err = tx.Create(&o).Error; err != nil {
				return nil, fmt.Errorf("create model offsets: %w", err)
			}
		}

		for _, w := range v.weights {
			w.InnerWeightsID = v.weightsInfo.InnerID
			if err = tx.Create(&w).Error; err != nil {
				return nil, fmt.Errorf("create model weights: %w", err)
			}
		}
	}

	return ids, nil
}

func (r *Repository) getStructInfo(id string) (structure.Structure, error) {
	var modelStruct structure.Structure
	err := r.db.Where("structure_id = ?", id).First(&modelStruct).Error
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
