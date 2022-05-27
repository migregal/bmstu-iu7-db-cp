package modelstructweightsinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/core/ports/repositories"
	dbweights "neural_storage/database/core/entities/structure/weights"
)

func (r *Repository) Find(filter repositories.StructWeightsInfoFilter) ([]*weights.Info, error) {
	query := r.db.DB
	if len(filter.Ids) > 0 {
		query = query.Where("id in ?", filter.Ids)
	}
	if len(filter.Structures) > 0 {
		query = query.Where("structure_id in ?", filter.Structures)
	}
	if len(filter.Names) > 0 {
		query = query.Where("name in ?", filter.Names)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}

	var weightsInfo []dbweights.Weights
	err := query.Find(&weightsInfo).Error
	if err != nil {
		return nil, fmt.Errorf("model get error: %w", err)
	}

	dbInfo := []*weights.Info{}
	for _, v := range weightsInfo {
		data, err := r.Get(v.GetID())
		if err != nil {
			return nil, err
		}

		dbInfo = append(dbInfo, data)
	}
	return dbInfo, nil
}
