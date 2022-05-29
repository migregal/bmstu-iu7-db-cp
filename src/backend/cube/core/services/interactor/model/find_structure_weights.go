package model

import (
	"context"
	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/core/ports/repositories"
	"neural_storage/pkg/logger"
)

func (i *Interactor) FindStructureWeights(ctx context.Context, filter interactors.ModelWeightsInfoFilter) ([]*sw.Info, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"filter": filter}).Info("model weights find called")

	lg.Info("attempt to find model weights info")
	return i.weightsInfo.Find(
		repositories.StructWeightsInfoFilter{
			Structures: filter.Structures,
			Ids:        filter.IDs,
			Names:      filter.Names,
			Offset:     filter.Offset,
			Limit:      filter.Limit,
		})
}
