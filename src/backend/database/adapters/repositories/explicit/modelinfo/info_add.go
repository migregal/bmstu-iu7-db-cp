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

	id, err := r.createModelInfo(data.model)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if data.structure == nil {
		tx.Rollback()
		return "", fmt.Errorf("missing model structure info")
	}

	structureId, err := r.createStructInfo(*data.structure)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = r.createLayersInfo(data.layers)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = r.createNeuronsInfo(data.neurons)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = r.createLinksInfo(structureId, data.links)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = r.createWeightsInfoTransact(database.Interactor{DB: tx}, data.weights)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return id, tx.Commit().Error
}

func (r *Repository) createModelInfo(info dbmodel.Model) (string, error) {
	m := dbmodel.Model{ID: info.ID, Name: info.Name}
	err := r.db.Create(&m).Error
	return m.ID, err
}

func (r *Repository) createStructInfo(info dbstructure.Structure) (string, error) {
	m := dbstructure.Structure{ID: info.ID, Name: info.Name}
	err := r.db.Create(&m).Error

	if err != nil {
		return m.ID, fmt.Errorf("add struct info: %w", err)
	}
	return m.ID, nil
}

func (r *Repository) createLayersInfo(info []dblayer.Layer) error {
	var layers []dblayer.Layer
	for _, v := range info {
		layers = append(layers, dblayer.Layer{StructID: v.StructID, LimitFunc: v.LimitFunc, ActivationFunc: v.ActivationFunc})
	}
	if err := r.db.Create(&layers).Error; err != nil {
		return fmt.Errorf("add struct info: %w", err)
	}
	return nil
}

func (r *Repository) createNeuronsInfo(info []dbneuron.Neuron) error {
	var neurons []dbneuron.Neuron
	for _, v := range info {
		neurons = append(neurons, dbneuron.Neuron{NeuronID: v.NeuronID, LayerID: v.LayerID})
	}
	return r.db.Create(&neurons).Error
}

func (r *Repository) createLinksInfo(structureID string, info []dblink.Link) error {
	var links []dblink.Link
	for _, v := range info {
		links = append(links,
			dblink.Link{
				StructureID: structureID,
				LinkID:      v.LinkID,
				FromID:      v.FromID,
				ToID:        v.ToID,
			})
	}
	return r.db.Create(&links).Error
}

func (r *Repository) createWeightsInfoTransact(tx database.Interactor, info []accumulatedWeightInfo) error {
	for _, v := range info {
		if v.weightsInfo == nil {
			return fmt.Errorf("missing weights info data")
		}
		if err := tx.Create(v.weightsInfo).Error; err != nil {
			return fmt.Errorf("create model weights info: %w", err)
		}
		for _, o := range v.offsets {
			if err := tx.Create(&o).Error; err != nil {
				return fmt.Errorf("create model offsets: %w", err)
			}
		}

		for _, w := range v.weights {
			if err := tx.Create(&w).Error; err != nil {
				return fmt.Errorf("create model weights: %w", err)
			}
		}
	}

	return nil
}
