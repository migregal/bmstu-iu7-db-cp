package modelstructweightsinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/structure/weights/weightsstat"
	dbstat "neural_storage/database/core/entities/structure/weights/weightsstat"
	"time"
)

func (r *Repository) GetAddStat(from, to time.Time) ([]*weightsstat.Info, error) {
	query := r.db.DB
	if !from.IsZero() {
		query = query.Where("created_at > ?", from)
	}
	if !to.IsZero() {
		query = query.Where("created_at < ?", to)
	}

	var info []dbstat.Info
	if err := query.Find(&info).Error; err != nil {
		return nil, fmt.Errorf("weights stat get error: %w", err)
	}

	var res []*weightsstat.Info
	for _, v := range info {
		res = append(res, weightsstat.New(v.GetID(), v.GetCreatedAt(), v.GetUpdatedAt()))
	}
	return res, nil
}
