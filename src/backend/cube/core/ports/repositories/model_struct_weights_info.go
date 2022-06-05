//go:generate mockery --name=ModelStructWeightsInfoRepository --outpkg=mock --output=../../../../database/adapters/repositories/mock/ --filename=model_struct_weights_info_repository.go --structname=ModelStructWeightsInfoRepository

package repositories

import (
	"neural_storage/cube/core/entities/structure/weights"
	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/core/entities/structure/weights/weightsstat"
	"time"
)

type ModelStructWeightsInfoRepository interface {
	Add(structure string, info []sw.Info) ([]string, error)
	Get(weightsId string) (*sw.Info, error)
	Find(filter StructWeightsInfoFilter) ([]*sw.Info, error)
	Update(info sw.Info) error
	Delete(info []weights.Info) error

	GetAddStat(from, to time.Time) ([]*weightsstat.Info, error)
	GetUpdateStat(from, to time.Time) ([]*weightsstat.Info, error)
}

type StructWeightsInfoFilter struct {
	Structures []string
	Ids        []string
	Names      []string
	Offset     int
	Limit      int
}
