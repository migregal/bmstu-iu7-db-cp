package modelstructweightsinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/database/core/services/interactor/database"
)

func (r *Repository) Update(info weights.Info) error {
	data := toDBEntity("", []weights.Info{info})

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := r.updateModelWeightsTransact(database.Interactor{DB: tx}, data); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) updateModelWeightsTransact(tx database.Interactor, info []accumulatedWeightInfo) error {
	for _, v := range info {
		if err := tx.Where("id = ?", v.weightsInfo.GetID()).Updates(&v.weightsInfo).Error; err != nil {
			return fmt.Errorf("model weights info update: %w", err)
		}
		for _, w := range v.weights {
			if err := tx.Where("id = ?", w.GetID()).Updates(&w).Error; err != nil {
				return fmt.Errorf("model weights update: %w", err)
			}
		}
		for _, o := range v.offsets {
			if err := tx.Where("id = ?", o.GetID()).Updates(&o).Error; err != nil {
				return fmt.Errorf("model offsets update: %w", err)
			}
		}
	}

	return nil
}
