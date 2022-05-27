package model

import (
	"neural_storage/cube/core/entities/model"
	sw "neural_storage/cube/core/entities/structure/weights"
)

func (i *Interactor) AddStructureWeights(ownerID, structID string, info sw.Info) error {
	structure, err := i.modelInfo.GetStructure(structID)
	if err != nil {
		return err
	}

	structure.SetWeights([]*sw.Info{&info})

	if err := i.validator.ValidateModelInfo(model.NewInfo(ownerID, "", structure)); err != nil {
		return nil
	}

	return i.weightsInfo.Add(structID, []sw.Info{info})
}
