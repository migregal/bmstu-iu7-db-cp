package repositories

import (
	"neural_storage/cube/core/ports/repositories"
	"neural_storage/database/adapters/repositories/explicit/modelinfo"
	"neural_storage/database/adapters/repositories/explicit/modelstructweightsinfo"
	"neural_storage/database/adapters/repositories/explicit/userinfo"
	"neural_storage/database/adapters/repositories/mock"
	"neural_storage/database/core/ports/config"
)

func NewUserInfoAdapter(conf config.UserInfoRepositoryConfig) repositories.UserInfoRepository {
	if conf.IsMocked() {
		return &mock.UserInfoRepository{}
	}

	switch conf.Adapter() {
	case "expl":
		repo, err := userinfo.NewRepository(conf)
		if err != nil {
			panic(err)
		}
		return &repo
	}

	panic("wrong db adapter specified")
}

func NewModelInfoAdapter(conf config.ModelInfoRepositoryConfig) repositories.ModelInfoRepository {
	if conf.IsMocked() {
		return &mock.ModelInfoRepository{}
	}
	repo, err := modelinfo.NewRepository(conf)
	if err != nil {
		panic(err)
	}
	return &repo
}

func NewModelStructureWeightsInfoAdapter(
	conf config.ModelStructureWeightsInfoRepositoryConfig,
) repositories.ModelStructWeightsInfoRepository {
	if conf.IsMocked() {
		return &mock.ModelStructWeightsInfoRepository{}
	}
	repo, err := modelstructweightsinfo.NewRepository(conf)
	if err != nil {
		panic(err)
	}
	return &repo
}
