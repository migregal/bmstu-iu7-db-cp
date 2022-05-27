package modelinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/model"
	"neural_storage/database/core/services/interactor/database"
)

func (r *Repository) Update(info model.Info) error {
	data := toDBEntity(info)

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

	if err := tx.Where("id = ?", data.model.GetID()).Updates(data.model).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("model info update: %w", err)
	}

	if data.structure == nil {
		return tx.Commit().Error
	}

	if err := r.updateModelStructureTransact(database.Interactor{DB: tx}, data); err != nil {
		tx.Rollback()
		return err
	}

	if err := r.updateModelWeightsTransact(database.Interactor{DB: tx}, data.weights); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) updateModelStructureTransact(tx database.Interactor, info accumulatedModelInfo) error {
	if info.structure == nil {
		return nil
	}

	if err := tx.Where("id = ?", info.structure.GetID()).Updates(&info.structure).Error; err != nil {
		return fmt.Errorf("model structure info update: %w", err)
	}

	for _, v := range info.layers {
		if err := tx.Where("id = ?", v.GetID()).Updates(&v).Error; err != nil {
			return fmt.Errorf("model layers info update: %w", err)
		}
	}

	for _, v := range info.neurons {
		if err := tx.Where("id = ?", v.GetID()).Updates(&v).Error; err != nil {
			return fmt.Errorf("model neurons info update: %w", err)
		}
	}

	for _, v := range info.links {
		if err := tx.Where("id = ?", v.GetID()).Updates(&v).Error; err != nil {
			return fmt.Errorf("model links info update: %w", err)
		}
	}

	return nil
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
