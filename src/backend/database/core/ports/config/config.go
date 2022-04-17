//go:generate mockery --name=UserInfoRepositoryConfig --outpkg=mock --output=../../../../config/adapters/repositories/mock/ --filename=user_info_repository_config.go --structname=UserInfoRepositoryConfig
//go:generate mockery --name=ModelInfoRepositoryConfig --outpkg=mock --output=../../../../config/adapters/repositories/mock/ --filename=model_info_repository_config.go --structname=ModelInfoRepositoryConfig
//go:generate mockery --name=ModelStructureWeightsInfoRepositoryConfig --outpkg=mock --output=../../../../config/adapters/repositories/mock/ --filename=model_structure_weights_info_repository_config.go --structname=ModelStructureWeightsInfoRepositoryConfig

package config

type UserInfoRepositoryConfig interface {
	IsMocked() bool
}

type ModelInfoRepositoryConfig interface {
	IsMocked() bool
}

type ModelStructureWeightsInfoRepositoryConfig interface {
	IsMocked() bool
}
