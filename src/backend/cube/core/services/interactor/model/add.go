package model

import "neural_storage/cube/core/entities/model"

func (i *Interactor) Add(info model.Info) error {
	if err := i.validator.ValidateModelInfo(&info); err != nil {
		return nil
	}

	return i.modelInfo.Add(info)
}
