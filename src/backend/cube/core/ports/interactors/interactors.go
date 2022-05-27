package interactors

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/model/modelstat"
	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/core/entities/structure/weights/weightsstat"
	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/entities/user/userstat"
	"time"
)

type UserInfoFilter struct {
	Ids       []string
	Usernames []string
	Emails    []string
	Offset    int
	Limit     int
}

type UserInfoInteractor interface {
	Register(info user.Info) (string, error)
	Get(id string) (user.Info, error)
	Find(filter UserInfoFilter) ([]user.Info, error)
	Update(info user.Info) error
	Block(userId string, until time.Time) error
	Delete(userId string) error

	GetUserRegistrationStat(from, to time.Time) ([]*userstat.Info, error)
	GetUserEditStat(from, to time.Time) ([]*userstat.Info, error)
}

type ModelInfoFilter struct {
	Owners []string
	Ids    []string
	Names  []string
	Offset int
	Limit  int
}

type ModelWeightsInfoFilter struct {
	Structures []string
	IDs        []string
	Names      []string
	Offset     int
	Limit      int
}

type NeuralNetworkInteractor interface {
	Add(info model.Info) error
	Get(modelID string) (*model.Info, error)
	Find(filter ModelInfoFilter) ([]*model.Info, error)
	Delete(userID, modelID string) error

	AddStructureWeights(ownerID string, modelID string, info sw.Info) error
	GetStructureWeights(weightsId string) (*sw.Info, error)
	FindStructureWeights(filter ModelWeightsInfoFilter) ([]*sw.Info, error)
	UpdateStructureWeights(ownerID, modelID string, info sw.Info) error
	DeleteStructureWeights(ownerID, weightsID string) error

	GetModelLoadStat(from, to time.Time) ([]*modelstat.Info, error)
	GetModelEditStat(from, to time.Time) ([]*modelstat.Info, error)
	GetWeightsLoadStat(from, to time.Time) ([]*weightsstat.Info, error)
	GetWeightsEditStat(from, to time.Time) ([]*weightsstat.Info, error)
}
