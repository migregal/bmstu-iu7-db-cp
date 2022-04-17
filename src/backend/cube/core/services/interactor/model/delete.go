package model

func (i *Interactor) Delete(modelId string) error {
	return i.modelInfo.Delete(modelId)
}
