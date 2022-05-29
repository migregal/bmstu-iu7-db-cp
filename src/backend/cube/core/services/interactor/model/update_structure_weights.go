package model

import (
	"context"
	"fmt"
	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/pkg/logger"
)

func (i *Interactor) UpdateStructureWeights(ctx context.Context, ownerID, modelId string, info sw.Info) error {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"owner": ownerID, "model": modelId}).Info("weights update called")

	lg.Info("attempt to get model info")
	model, err := i.modelInfo.Get(modelId)
	if err != nil {
		lg.Error("failed to get model info")
		return err
	}

	if ownerID != "" && model.OwnerID() != ownerID {
		lg.Error("permission denied")
		return fmt.Errorf("permission denied")
	}

	model.Structure().SetWeights([]*sw.Info{&info})

	lg.Error("invalid result model info")
	if err := i.validator.ValidateModelInfo(model); err != nil {
		return err
	}

	lg.Info("attempt to update weights info")
	return i.weightsInfo.Update(info)
}
