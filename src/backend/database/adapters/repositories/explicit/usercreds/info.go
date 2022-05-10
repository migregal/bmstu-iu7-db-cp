package usercreds

import (
	"neural_storage/cube/core/entities/user"
	"neural_storage/database/core/entities/user_creds"
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

func (r *Repository) Add(info user.Info) error {
	data := toDBEntity(info)

	err := r.db.Create(&data).Error
	if err == nil {
		info.SetId(&data.ID)
	}
	return err
}

func (r *Repository) Get(userID string) (user.Info, error) {
	var res user_creds.UserCreds

	err := r.db.Where("user_id = ?", userID).First(&res).Error
	if err != nil {
		return user.Info{}, err
	}

	return fromDBEntity(res), nil
}

func (r *Repository) Update(info user.Info) error {
	data := toDBEntity(info)

	return r.db.Save(&data).Error
}

func (r *Repository) Delete(info user.Info) error {
	data := toDBEntity(info)
	return r.db.Delete(&data).Error
}
