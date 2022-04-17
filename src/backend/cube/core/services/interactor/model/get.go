package model

import "neural_storage/cube/core/entities/model"

func (i *Interactor) Get(modelId string) (*model.Info, error) {
	return i.modelInfo.Get(modelId)
}
