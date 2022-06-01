package modelstructweightsinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/structure/weights"
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

	data := toDBEntity(structID, info)

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
			if err = tx.Create(&o).Error; err != nil {
				return nil, fmt.Errorf("create model offsets: %w", err)
			}
		}

		for _, w := range v.weights {
			if err = tx.Create(&w).Error; err != nil {
				return nil, fmt.Errorf("create model weights: %w", err)
			}
		}
	}

	return ids, nil
}
