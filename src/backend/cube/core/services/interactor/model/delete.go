package model

import "neural_storage/cube/core/entities/model"

func (i *Interactor) Delete(ownerID, modelID string) error {
	return i.modelInfo.Delete(*model.NewInfo(ownerID, modelID, nil))
}
