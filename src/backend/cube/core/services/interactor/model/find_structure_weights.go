package model

import (
	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/core/ports/repositories"
)

type StructureWeightsFilter struct {
	Ids   []string
	Limit int
}

func (i *Interactor) FindStructureWeights(filter StructureWeightsFilter) ([]*sw.Info, error) {
	return i.weightsInfo.Find(
		repositories.StructWeightsInfoFilter{
			Ids:   filter.Ids,
			Limit: filter.Limit,
		})
}
