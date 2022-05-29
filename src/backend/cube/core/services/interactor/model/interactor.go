package model

import (
	"neural_storage/cube/core/ports/config"
	r "neural_storage/cube/core/ports/repositories"
	r2 "neural_storage/cube/core/ports/validator"
	"neural_storage/pkg/logger"

	adapters2 "neural_storage/cube/adapters/validator"
	adapters "neural_storage/database/adapters/repositories"
)

type Interactor struct {
	modelInfo   r.ModelInfoRepository
	weightsInfo r.ModelStructWeightsInfoRepository
	validator   r2.Validator

	lg *logger.Logger
}

func NewInteractor(lg *logger.Logger, conf config.ModelInfoInteractorConfig) *Interactor {
	interactor := Interactor{
		modelInfo:   adapters.NewModelInfoAdapter(conf.ModelInfoRepoConfig()),
		weightsInfo: adapters.NewModelStructureWeightsInfoAdapter(conf.ModelStructureWeightInfoRepoConfig()),
		validator:   adapters2.NewValidator(conf.ValidatorConfig()),
		lg:          lg,
	}

	return &interactor
}
