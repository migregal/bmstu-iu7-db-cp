package model

import (
	"context"
	"neural_storage/cube/core/entities/model"
	"neural_storage/pkg/logger"
)

func (i *Interactor) Delete(ctx context.Context, ownerID, modelID string) error {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"owner": ownerID, "model": modelID}).Info("model delete called")

	lg.Info("attempt to delete model info")
	return i.modelInfo.Delete(*model.NewInfo(modelID, ownerID, "", nil))
}
