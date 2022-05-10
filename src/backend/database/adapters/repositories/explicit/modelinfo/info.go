package modelinfo

import (
	"neural_storage/database/core/ports/config"
	"neural_storage/database/core/services/interactor/database"
)

type Repository struct {
	db database.Interactor
}

func NewRepository(conf config.UserInfoRepositoryConfig) (Repository, error) {
	db, err := database.New(conf.ConnParams())
	if err != nil {
		return Repository{}, err
	}

	return Repository{db: db}, nil
}
