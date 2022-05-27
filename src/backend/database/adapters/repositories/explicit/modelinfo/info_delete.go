package modelinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/model"
	dbweights "neural_storage/database/core/entities/structure/weights"
	"neural_storage/database/core/services/interactor/database"
)

func (r *Repository) Delete(info model.Info) error {
	data := toDBEntity(info)

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

	if data.structure == nil {
		if err := tx.Delete(data.model).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("model info update: %w", err)
		}
		return tx.Commit().Error
	}

	if err := r.deleteModelStructureTransact(database.Interactor{DB: tx}, data); err != nil {
		tx.Rollback()
		return err
	}

	if err := r.deleteModelWeightsTransact(database.Interactor{DB: tx}, data.weights); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) deleteModelStructureTransact(tx database.Interactor, info accumulatedModelInfo) error {
	if info.structure == nil {
		return nil
	}

	if len(info.layers) == 0 && len(info.neurons) == 0 && len(info.links) == 0 {
		if err := tx.Delete(info.structure).Error; err != nil {
			return fmt.Errorf("model structure info update: %w", err)
		}
		return nil
	}

	if len(info.layers) > 0 {
		for _, v := range info.layers {
			if err := tx.Where("id = ?", v.GetID()).Delete(&v).Error; err != nil {
				return fmt.Errorf("model layers info update: %w", err)
			}
		}
	}

	if len(info.neurons) > 0 {
		for _, v := range info.neurons {
			if err := tx.Where("id = ?", v.GetID()).Delete(&v).Error; err != nil {
				return fmt.Errorf("model neurons info update: %w", err)
			}
		}
	}

	if len(info.links) > 0 {
		for _, v := range info.links {
			if err := tx.Where("id = ?", v.GetID()).Delete(&v).Error; err != nil {
				return fmt.Errorf("model links info update: %w", err)
			}
		}
	}

	return nil
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
