package model

import (
	"context"
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/core/ports/repositories"
	"neural_storage/pkg/logger"
)

func (i *Interactor) Find(ctx context.Context, filter interactors.ModelInfoFilter) ([]*model.Info, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"filter": filter}).Info("model find called")

	lg.Info("attempt to find model info")
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
