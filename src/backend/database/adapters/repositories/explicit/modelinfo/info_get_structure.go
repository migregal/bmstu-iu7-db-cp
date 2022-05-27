package modelinfo

import (
	"neural_storage/cube/core/entities/structure"
)

func (r *Repository) GetStructure(modelId string) (*structure.Info, error) {
	model, err := r.Get(modelId)
	if err != nil {
		return nil, err
	}
	return model.Structure(), nil
}
