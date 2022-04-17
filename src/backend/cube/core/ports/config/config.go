//go:generate mockery --name=UserInfoInteractorConfig --outpkg=mock --output=../../../../config/adapters/interactors/mock/ --filename=user_info_interactor_config.go --structname=UserInfoInteractorConfig
//go:generate mockery --name=ModelInfoInteractorConfig --outpkg=mock --output=../../../../config/adapters/interactors/mock/ --filename=model_info_interactor_config.go --structname=ModelInfoInteractorConfig
//go:generate mockery --name=ValidatorConfig --outpkg=mock --output=../../../../config/adapters/validator/mock/ --filename=validator_config.go --structname=ValidatorConfig
//go:generate mockery --name=NormalizerConfig --outpkg=mock --output=../../../../config/adapters/normalizer/mock/ --filename=normalizer_config.go --structname=NormalizerConfig

package config

import (
	"neural_storage/database/core/ports/config"
)

type UserInfoInteractorConfig interface {
	RepoConfig() config.UserInfoRepositoryConfig
	ValidatorConfig() ValidatorConfig
	NormalizerConfig() NormalizerConfig
}

type ModelInfoInteractorConfig interface {
	ModelInfoRepoConfig() config.ModelInfoRepositoryConfig
	ModelStructureWeightInfoRepoConfig() config.ModelStructureWeightsInfoRepositoryConfig
	ValidatorConfig() ValidatorConfig
}

type ValidatorConfig interface {
	IsMocked() bool

	MinUnameLen() int
	MaxUnameLen() int

	MinPwdLen() int
	MaxPwdLen() int
}

type NormalizerConfig interface {
}
