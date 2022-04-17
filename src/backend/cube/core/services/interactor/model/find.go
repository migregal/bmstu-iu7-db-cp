package model

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/ports/repositories"
)

type Filter struct {
	Ids   []string
	Limit int
}

func (i *Interactor) Find(filter Filter) ([]*model.Info, error) {
	return i.modelInfo.Find(
		repositories.ModelInfoFilter{
			Ids:   filter.Ids,
			Limit: filter.Limit,
		},
	)
}
