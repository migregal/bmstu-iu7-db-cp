package model

import (
	"neural_storage/cube/core/entities/model"
	sw "neural_storage/cube/core/entities/structure/weights"
)

func (i *Interactor) AddStructureWeights(modelId string, info sw.Info) error {
	structure, err := i.modelInfo.GetStructure(modelId)
	if err != nil {
		return err
	}

	structure.SetWeights([]*sw.Info{&info})

	if err := i.validator.ValidateModelInfo(model.NewInfo("", structure)); err != nil {
		return nil
	}

	return i.weightsInfo.Add(modelId, info)
}
