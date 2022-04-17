package interactors

import (
	"neural_storage/cube/core/entities/model"
	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/core/entities/user"
	"time"
)

type UserInfoInteractor interface {
	Register(info user.Info) error
	Get(info user.Info) (user.Info, error)
	Update(info user.Info) error
	Block(userId string, until time.Time) error
	Delete(userId string) error
}

type NeuralNetworkInteractor interface {
	Add(info model.Info) error
	Get(modelId string) (*model.Info, error)
	Delete(modelId string) error

	AddStructureWeights(modelId string, info sw.Info) error
	GetStructureWeights(weightsId string) (*sw.Info, error)
	UpdateStructureWeights(info sw.Info) error
	DeleteStructureWeights(weightsId string) error
}
