package repositories

import (
	"neural_storage/cube/core/ports/repositories"
	"neural_storage/database/adapters/repositories/mock"
	"neural_storage/database/core/ports/config"
)

func NewUserInfoAdapter(conf config.UserInfoRepositoryConfig) repositories.UserInfoRepository {
	if conf.IsMocked() {
		return &mock.UserInfoRepository{}
	}
	return nil
}

func NewModelInfoAdapter(conf config.ModelInfoRepositoryConfig) repositories.ModelInfoRepository {
	if conf.IsMocked() {
		return &mock.ModelInfoRepository{}
	}
	return nil
}

func NewModelStructureWeightsInfoAdapter(
	conf config.ModelStructureWeightsInfoRepositoryConfig,
) repositories.ModelStructWeightsInfoRepository {
	if conf.IsMocked() {
		return &mock.ModelStructWeightsInfoRepository{}
	}
	return nil
}
