package model

import "neural_storage/cube/core/entities/structure/weights"

func (i *Interactor) DeleteStructureWeights(ownerID, weightsId string) error {
	info := *weights.NewInfo(weightsId, "", nil, nil)
	return i.weightsInfo.Delete([]weights.Info{info})
}
