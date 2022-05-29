package interactors

import (
	"context"
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
	Register(ctx context.Context, info user.Info) (string, error)
	Get(ctx context.Context, id string) (user.Info, error)
	Find(ctx context.Context, filter UserInfoFilter) ([]user.Info, error)
	Update(ctx context.Context, info user.Info) error
	Block(ctx context.Context, userId string, until time.Time) error
	Delete(ctx context.Context, userId string) error

	GetUserRegistrationStat(ctx context.Context, from, to time.Time) ([]*userstat.Info, error)
	GetUserEditStat(ctx context.Context, from, to time.Time) ([]*userstat.Info, error)
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
	Add(ctx context.Context, info model.Info) (string, error)
	Get(ctx context.Context, modelID string) (*model.Info, error)
	Find(ctx context.Context, filter ModelInfoFilter) ([]*model.Info, error)
	Delete(ctx context.Context, userID, modelID string) error

	AddStructureWeights(ctx context.Context, ownerID string, modelID string, info sw.Info) error
	GetStructureWeights(ctx context.Context, weightsId string) (*sw.Info, error)
	FindStructureWeights(ctx context.Context, filter ModelWeightsInfoFilter) ([]*sw.Info, error)
	UpdateStructureWeights(ctx context.Context, ownerID, modelID string, info sw.Info) error
	DeleteStructureWeights(ctx context.Context, ownerID, weightsID string) error

	GetModelLoadStat(ctx context.Context, from, to time.Time) ([]*modelstat.Info, error)
	GetModelEditStat(ctx context.Context, from, to time.Time) ([]*modelstat.Info, error)
	GetWeightsLoadStat(ctx context.Context, from, to time.Time) ([]*weightsstat.Info, error)
	GetWeightsEditStat(ctx context.Context, from, to time.Time) ([]*weightsstat.Info, error)
}
