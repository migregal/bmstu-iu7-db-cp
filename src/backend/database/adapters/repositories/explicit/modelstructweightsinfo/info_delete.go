package modelstructweightsinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/structure/weights"
	dbweights "neural_storage/database/core/entities/structure/weights"
	"neural_storage/database/core/services/interactor/database"
)

func (r *Repository) Delete(info []weights.Info) error {
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

	data := toDBEntity("", info)

	err := r.deleteModelWeightsTransact(database.Interactor{DB: tx}, data)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) deleteModelWeightsTransact(tx database.Interactor, info []accumulatedWeightInfo) error {
	for _, v := range info {
		if v.weightsInfo != nil {
			if err := tx.Where("id = ?", v.weightsInfo.GetID()).Delete(&v.weightsInfo).Error; err != nil {
				return fmt.Errorf("model weights infp delete: %w", err)
			}
			continue
		}

		for _, o := range v.offsets {
			if err := tx.Where("id = ?", o.GetWeightsID()).Delete(&dbweights.Weights{}).Error; err != nil {
				return fmt.Errorf("model weights info update: %w", err)
			}
		}

		for _, w := range v.weights {
			if err := tx.Where("id = ?", w.GetWeightsID()).Delete(&dbweights.Weights{}).Error; err != nil {
				return fmt.Errorf("model weights info update: %w", err)
			}
		}
	}

	return nil
}
