package modelinfo

import (
	"fmt"

	"neural_storage/cube/core/entities/model"
	dbmodel "neural_storage/database/core/entities/model"
	dbneuron "neural_storage/database/core/entities/neuron"
	dblink "neural_storage/database/core/entities/neuron/link"
	dbstructure "neural_storage/database/core/entities/structure"
	dblayer "neural_storage/database/core/entities/structure/layer"
	"neural_storage/database/core/services/interactor/database"
)

func (r *Repository) Add(info model.Info) (string, error) {
	data := toDBEntity(info)
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return "", err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	id, err := r.createModelInfo(database.Interactor{DB: tx}, data.model)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if data.structure == nil {
		tx.Rollback()
		return "", fmt.Errorf("missing model structure info")
	}

	data.structure.ModelID = id
	structureId, err := r.createStructInfo(database.Interactor{DB: tx}, *data.structure)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = r.createLayersInfo(database.Interactor{DB: tx}, structureId, data.layers)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = r.createNeuronsInfo(database.Interactor{DB: tx}, data.layers, data.neurons)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	neuronMap := map[int]string{}
	for i := range data.neurons {
		neuronMap[data.neurons[i].ID] = data.neurons[i].GetID()
	}

	err = r.createLinksInfo(database.Interactor{DB: tx}, structureId, neuronMap, data.links)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	linkMap := map[int]string{}
	for i := range data.links {
		linkMap[data.links[i].ID] = data.links[i].GetID()
	}

	err = r.createWeightsInfo(database.Interactor{DB: tx}, structureId, neuronMap, linkMap, data.weights)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return id, tx.Commit().Error
}

func (r *Repository) createModelInfo(tx database.Interactor, info dbmodel.Model) (string, error) {
	err := tx.Create(&info).Error
	return info.GetID(), err
}

func (r *Repository) createStructInfo(tx database.Interactor, info dbstructure.Structure) (string, error) {
	m := dbstructure.Structure{ID: info.GetID(), ModelID: info.GetModelID(), Name: info.GetName()}
	err := tx.Create(&m).Error

	if err != nil {
		return m.GetID(), fmt.Errorf("add struct info: %w", err)
	}
	return m.GetID(), nil
}

func (r *Repository) createLayersInfo(tx database.Interactor, structueID string, info []dblayer.Layer) error {
	for i := range info {
		info[i].StructureID = structueID
	}
	if err := tx.Create(&info).Error; err != nil {
		return fmt.Errorf("add struct info: %w", err)
	}
	return nil
}

func (r *Repository) createNeuronsInfo(tx database.Interactor, layers []dblayer.Layer, info []dbneuron.Neuron) error {
	layerMap := map[int]string{}
	for i := range layers {
		layerMap[layers[i].ID] = layers[i].GetID()
	}

	for j := range info {
		info[j].InnerLayerID = layerMap[info[j].LayerID]
	}

	return tx.CreateInBatches(&info, 3000).Error
}

func (r *Repository) createLinksInfo(
	tx database.Interactor,
	structureID string,
	neurons map[int]string,
	info []dblink.Link,
) error {
	for i := range info {
		info[i].InnerFrom = neurons[info[i].From]
		info[i].InnerTo = neurons[info[i].To]
	}

	return tx.Create(&info).Error
}

func (r *Repository) createWeightsInfo(
	tx database.Interactor,
	structureID string,
	neurons map[int]string,
	links map[int]string,
	info []accumulatedWeightInfo,
) error {
	for _, v := range info {
		if v.weightsInfo == nil {
			return fmt.Errorf("missing weights info data")
		}
		v.weightsInfo.StructureID = structureID
		if err := tx.Create(v.weightsInfo).Error; err != nil {
			return fmt.Errorf("create model weights info: %w", err)
		}
		for _, o := range v.offsets {
			o.InnerWeights = v.weightsInfo.GetID()
			o.InnerNeuron = neurons[o.Neuron]
			if err := tx.Create(&o).Error; err != nil {
				return fmt.Errorf("create model offsets: %w", err)
			}
		}

		for _, w := range v.weights {
			w.InnerWeightsID = v.weightsInfo.GetID()
			w.InnerLinkID = links[w.LinkID]
			if err := tx.Create(&w).Error; err != nil {
				return fmt.Errorf("create model weights: %w", err)
			}
		}
	}

	return nil
}
