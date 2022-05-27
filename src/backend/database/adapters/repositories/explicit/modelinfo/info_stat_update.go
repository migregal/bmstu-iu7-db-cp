package modelinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/model/modelstat"
	dbstat "neural_storage/database/core/entities/model/modelstat"
	"time"
)

func (r *Repository) GetUpdateStat(from, to time.Time) ([]*modelstat.Info, error) {
	query := r.db.DB
	if !from.IsZero() {
		query = query.Where("updated_at > ?", from)
	}
	if !to.IsZero() {
		query = query.Where("updated_at < ?", to)
	}

	var info []dbstat.Info
	if err := query.Find(&info).Error; err != nil {
		return nil, fmt.Errorf("model stat get error: %w", err)
	}

	var res []*modelstat.Info
	for _, v := range info {
		res = append(res, modelstat.New(v.GetID(), v.GetCreatedAt(), v.GetUpdatedAt()))
	}

	return res, nil
}
