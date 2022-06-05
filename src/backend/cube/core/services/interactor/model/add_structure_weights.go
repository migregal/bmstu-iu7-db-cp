package model

import (
	"context"
	"neural_storage/cube/core/entities/model"
	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/pkg/logger"
)

func (i *Interactor) AddStructureWeights(ctx context.Context, ownerID, modelID string, info sw.Info) (string, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"owner": ownerID, "struct": modelID}).Info("model struct weights add called")

	lg.Info("attempt to get struct info")
	structure, err := i.modelInfo.GetStructure(modelID)
	if err != nil {
		lg.Error("failed to get struct info")
		return "", err
	}

	structure.SetWeights([]*sw.Info{&info})

	if err := i.validator.ValidateModelInfo(model.NewInfo("", ownerID, "", structure)); err != nil {
		lg.Error("invlaid model info")
		return "", err
	}

	lg.Info("attempt to add struct weights info")

	if ids, err := i.weightsInfo.Add(structure.ID(), []sw.Info{info}); err != nil {
		return "", err
	} else if len(ids) > 0 {
		return ids[0], nil
	}

	return "", nil
}
