package model

import (
	"fmt"
	sw "neural_storage/cube/core/entities/structure/weights"
)

func (i *Interactor) UpdateStructureWeights(ownerID, modelId string, info sw.Info) error {
	model, err := i.modelInfo.Get(modelId)
	if err != nil {
		return err
	}

	if ownerID != "" && model.OwnerID() != ownerID {
		return fmt.Errorf("permission denied")
	}

	model.Structure().SetWeights([]*sw.Info{&info})

	if err := i.validator.ValidateModelInfo(model); err != nil {
		return nil
	}

	return i.weightsInfo.Update(info)
}
