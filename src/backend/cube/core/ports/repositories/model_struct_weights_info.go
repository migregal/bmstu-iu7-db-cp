//go:generate mockery --name=ModelStructWeightsInfoRepository --outpkg=mock --output=../../../../database/adapters/repositories/mock/ --filename=model_struct_weights_info_repository.go --structname=ModelStructWeightsInfoRepository

package repositories

import (
	sw "neural_storage/cube/core/entities/structure/weights"
)

type ModelStructWeightsInfoRepository interface {
	Add(modelId string, info sw.Info) error
	Get(weightsId string) (*sw.Info, error)
	Find(filter StructWeightsInfoFilter) ([]*sw.Info, error)
	Update(info sw.Info) error
	Delete(weightsId string) error
}

type StructWeightsInfoFilter struct {
	Ids   []string
	Limit int
}
