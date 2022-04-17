//go:generate mockery --name=ModelInfoRepository --outpkg=mock --output=../../../../database/adapters/repositories/mock/ --filename=model_info_repository.go --structname=ModelInfoRepository

package repositories

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/structure"
)

type ModelInfoRepository interface {
	Add(info model.Info) error
	Get(modelId string) (*model.Info, error)
	Find(filter ModelInfoFilter) ([]*model.Info, error)
	GetStructure(modelId string) (*structure.Info, error)
	Delete(modelId string) error
}

type ModelInfoFilter struct {
	Ids []string
	Limit int
}
