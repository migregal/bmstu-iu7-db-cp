package modelinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/ports/repositories"
	dbmodel "neural_storage/database/core/entities/model"
)

func (r *Repository) Find(filter repositories.ModelInfoFilter) ([]*model.Info, error) {
	query := r.db.DB
	if len(filter.Ids) > 0 {
		query = query.Where("id in ?", filter.Ids)
	}
	if len(filter.Names) > 0 {
		query = query.Where("name in ?", filter.Names)
	}
	if len(filter.Owners) > 0 {
		query = query.Where("owner_id in ?", filter.Owners)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}

	var modelInfo []dbmodel.Model
	err := query.Find(&modelInfo).Error
	if err != nil {
		return nil, fmt.Errorf("model get error: %w", err)
	}

	dbInfo := []*model.Info{}
	if len(filter.Ids) == 0 {
		for _, v := range modelInfo {
			st, err := r.getStructInfo(v.ID)
			if err != nil {
				return nil, err
			}

			dbInfo = append(dbInfo, model.NewInfo(
				v.GetID(), v.GetOwnerID(),
				v.GetName(),
				structure.NewInfo(st.GetID(), st.GetName(), nil, nil, nil, nil)))
		}
		return dbInfo, nil
	}

	for _, v := range modelInfo {
		data, err := r.Get(v.GetID())
		if err != nil {
			return nil, err
		}

		dbInfo = append(dbInfo, data)
	}
	return dbInfo, nil
}
