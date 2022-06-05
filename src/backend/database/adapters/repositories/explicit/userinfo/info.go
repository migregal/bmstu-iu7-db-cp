package userinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/entities/user/userstat"
	"neural_storage/cube/core/ports/repositories"
	"neural_storage/database/core/entities/user_info"
	"neural_storage/database/core/entities/user_info/userinfostat"
	"neural_storage/database/core/ports/config"
	"neural_storage/database/core/services/interactor/database"
	"time"
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

func (r *Repository) Add(info user.Info) (string, error) {
	data := toDBEntity(info)

	err := r.db.Create(&data).Error
	if err == nil {
		info.SetId(data.ID)
	}

	return data.ID, err
}

func (r *Repository) Get(id string) (user.Info, error) {
	var res user_info.UserInfo

	err := r.db.Where("id = ?", id).First(&res).Error
	if err != nil {
		return user.Info{}, err
	}

	return fromDBEntity(res), nil
}

func (r *Repository) Find(filter repositories.UserInfoFilter) ([]user.Info, error) {
	query := r.db.DB

	if len(filter.UserIds) > 0 {
		query = query.Where("id in ?", filter.UserIds)
	}
	if len(filter.Usernames) > 0 {
		query = query.Where("username in ?", filter.Usernames)
	}
	if len(filter.Emails) > 0 {
		query = query.Where("email in ?", filter.Emails)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}

	var res []user_info.UserInfo

	err := query.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return fromDBEntities(res), nil
}

func (r *Repository) Update(info user.Info) error {
	data := toDBEntity(info)
	return r.db.DB.Updates(&data).Error
}

func (r *Repository) Delete(info user.Info) error {
	data := toDBEntity(info)
	return r.db.Delete(&data).Error
}

func (r *Repository) GetAddStat(from, to time.Time) ([]*userstat.Info, error) {
	query := r.db.DB
	if !from.IsZero() {
		query = query.Where("created_at > ?", from)
	}
	if !to.IsZero() {
		query = query.Where("created_at < ?", to)
	}

	var info []userinfostat.Info
	if err := query.Find(&info).Error; err != nil {
		return nil, fmt.Errorf("user get error: %w", err)
	}

	var res []*userstat.Info
	for _, v := range info {
		res = append(res, userstat.New(v.GetID(), v.GetCreatedAt(), v.GetUpdatedAt()))
	}
	return res, nil
}

func (r *Repository) GetUpdateStat(from, to time.Time) ([]*userstat.Info, error) {
	query := r.db.DB
	if !from.IsZero() {
		query = query.Where("updated_at > ?", from)
	}
	if !to.IsZero() {
		query = query.Where("updated_at < ?", to)
	}

	var info []userinfostat.Info
	if err := query.Find(&info).Error; err != nil {
		return nil, fmt.Errorf("user get error: %w", err)
	}

	var res []*userstat.Info
	for _, v := range info {
		res = append(res, userstat.New(v.GetID(), v.GetCreatedAt(), v.GetUpdatedAt()))
	}
	return res, nil
}
