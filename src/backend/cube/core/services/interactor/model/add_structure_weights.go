package model

import (
	"context"
	"neural_storage/cube/core/entities/model"
	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/pkg/logger"
)

func (i *Interactor) AddStructureWeights(ctx context.Context, ownerID, structID string, info sw.Info) error {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"owner": ownerID, "struct": structID}).Info("model struct weights add called")

	lg.Info("attempt to get struct info")
	structure, err := i.modelInfo.GetStructure(structID)
	if err != nil {
		lg.Error("failed to get struct info")
		return err
	}

	structure.SetWeights([]*sw.Info{&info})

	if err := i.validator.ValidateModelInfo(model.NewInfo(ownerID, "", structure)); err != nil {
		lg.Error("invlaid model info")
		return err
	}

	lg.Info("attempt to add struct weights info")
	return i.weightsInfo.Add(structID, []sw.Info{info})
}
