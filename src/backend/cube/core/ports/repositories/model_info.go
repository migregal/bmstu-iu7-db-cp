//go:generate mockery --name=ModelInfoRepository --outpkg=mock --output=../../../../database/adapters/repositories/mock/ --filename=model_info_repository.go --structname=ModelInfoRepository

package repositories

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/model/modelstat"
	"neural_storage/cube/core/entities/structure"
	"time"
)

type ModelInfoRepository interface {
	Add(info model.Info) (string, error)
	Update(info model.Info) error
	Get(modelId string) (*model.Info, error)
	Find(filter ModelInfoFilter) ([]*model.Info, error)
	GetStructure(modelId string) (*structure.Info, error)
	Delete(info model.Info) error

	GetAddStat(from, to time.Time) ([]*modelstat.Info, error)
	GetUpdateStat(from, to time.Time) ([]*modelstat.Info, error)
}

type ModelInfoFilter struct {
	Owners []string
	Ids    []string
	Names  []string
	Offset int
	Limit  int
}
