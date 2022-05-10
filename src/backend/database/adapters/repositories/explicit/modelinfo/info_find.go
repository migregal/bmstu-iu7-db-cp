package modelinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/ports/repositories"
	dbmodel "neural_storage/database/core/entities/model"
)

func (r *Repository) Find(filter repositories.ModelInfoFilter) ([]model.Info, error) {
	if len(filter.Ids) == 0 {
		query := r.db.DB
		if filter.Limit > 0 {
			query = query.Limit(filter.Limit)
		}
		if len(filter.Ids) > 0 {
			query = query.Where("id in ?", filter.Ids)
		}
		if filter.OwnerID != "" {
			query = query.Where("owner_id = ?", filter.OwnerID)
		}
		var modelInfo []dbmodel.Model
		err := query.Find(&modelInfo).Error
		if err != nil {
			return nil, fmt.Errorf("model get error: %w", err)
		}

		for _, v := range modelInfo {
			filter.Ids = append(filter.Ids, v.ID)
		}
	}

	dbInfo := []model.Info{}
	for _, v := range filter.Ids {
		data, err := r.Get(v)
		if err != nil {
			return nil, err
		}

		dbInfo = append(dbInfo, data)
	}
	return dbInfo, nil
}
