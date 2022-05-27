package model

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/core/ports/repositories"
)

func (i *Interactor) Find(filter interactors.ModelInfoFilter) ([]*model.Info, error) {
	return i.modelInfo.Find(
		repositories.ModelInfoFilter{
			Owners: filter.Owners,
			Ids:    filter.Ids,
			Names:  filter.Names,
			Offset: filter.Offset,
			Limit:  filter.Limit,
		},
	)
}
