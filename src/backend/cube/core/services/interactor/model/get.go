package model

import (
	"context"
	"neural_storage/cube/core/entities/model"
	"neural_storage/pkg/logger"
)

func (i *Interactor) Get(ctx context.Context, modelId string) (*model.Info, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"id": modelId}).Info("model get called")

	lg.Info("attempt to get model info")
	return i.modelInfo.Get(modelId)
}
