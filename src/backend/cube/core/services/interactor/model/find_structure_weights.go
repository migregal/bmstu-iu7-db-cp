package model

import (
	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/core/ports/repositories"
)

func (i *Interactor) FindStructureWeights(filter interactors.ModelWeightsInfoFilter) ([]*sw.Info, error) {
	return i.weightsInfo.Find(
		repositories.StructWeightsInfoFilter{
			Structures: filter.Structures,
			Ids:        filter.IDs,
			Names:      filter.Names,
			Offset:     filter.Offset,
			Limit:      filter.Limit,
		})
}
