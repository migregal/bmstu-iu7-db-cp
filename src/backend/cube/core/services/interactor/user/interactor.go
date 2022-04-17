package user

import (
	"neural_storage/cube/core/ports/config"
	"neural_storage/cube/core/ports/normalizer"
	"neural_storage/cube/core/ports/repositories"
	"neural_storage/cube/core/ports/validator"

	adapters3 "neural_storage/cube/adapters/normalizer"
	adapters2 "neural_storage/cube/adapters/validator"
	adapters "neural_storage/database/adapters/repositories"
)

type Interactor struct {
	userInfo repositories.UserInfoRepository

	validator  validator.Validator
	normalizer normalizer.Normalizer
}

func NewInteractor(conf config.UserInfoInteractorConfig) *Interactor {
	return &Interactor{
		userInfo:   adapters.NewUserInfoAdapter(conf.RepoConfig()),
		validator:  adapters2.NewValidator(conf.ValidatorConfig()),
		normalizer: adapters3.NewNormalizer(conf.NormalizerConfig()),
	}
}
