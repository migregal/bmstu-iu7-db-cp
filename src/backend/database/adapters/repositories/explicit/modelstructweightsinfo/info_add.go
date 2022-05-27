package modelstructweightsinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/database/core/services/interactor/database"
)

func (r *Repository) Add(structID string, info []weights.Info) error {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	data := toDBEntity(structID, info)

	err := r.createWeightsInfoTransact(database.Interactor{DB: tx}, data)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) createWeightsInfoTransact(tx database.Interactor, info []accumulatedWeightInfo) error {
	for _, v := range info {
		if err := tx.Create(&v.weightsInfo).Error; err != nil {
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
