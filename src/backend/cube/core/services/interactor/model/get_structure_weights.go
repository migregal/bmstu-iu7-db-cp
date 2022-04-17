package model

import (
	sw "neural_storage/cube/core/entities/structure/weights"
)

func (i *Interactor) GetStructureWeights(weightsId string) (*sw.Info, error) {
	return i.weightsInfo.Get(weightsId)
}
