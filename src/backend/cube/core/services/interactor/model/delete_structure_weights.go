package model

func (i *Interactor) DeleteStructureWeights(weightsId string) error {
	return i.weightsInfo.Delete(weightsId)
}
