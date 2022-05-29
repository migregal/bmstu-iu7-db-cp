package user

import (
	"neural_storage/cube/core/ports/config"
	"neural_storage/cube/core/ports/normalizer"
	"neural_storage/cube/core/ports/repositories"
	"neural_storage/cube/core/ports/validator"
	"neural_storage/pkg/logger"

	adapters3 "neural_storage/cube/adapters/normalizer"
	adapters2 "neural_storage/cube/adapters/validator"
	adapters "neural_storage/database/adapters/repositories"
)

type Interactor struct {
	userInfo repositories.UserInfoRepository

	validator  validator.Validator
	normalizer normalizer.Normalizer

	lg *logger.Logger
}

func NewInteractor(lg *logger.Logger, conf config.UserInfoInteractorConfig) *Interactor {
	return &Interactor{
		userInfo:   adapters.NewUserInfoAdapter(conf.RepoConfig()),
		validator:  adapters2.NewValidator(conf.ValidatorConfig()),
		normalizer: adapters3.NewNormalizer(conf.NormalizerConfig()),
		lg:         lg,
	}
}
