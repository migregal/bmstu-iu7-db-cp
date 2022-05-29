package model

import (
	"context"
	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/pkg/logger"
)

func (i *Interactor) GetStructureWeights(ctx context.Context, weightsId string) (*sw.Info, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"id": weightsId}).Info("weights get called")

	lg.Info("attempt to get weights info")
	return i.weightsInfo.Get(weightsId)
}
