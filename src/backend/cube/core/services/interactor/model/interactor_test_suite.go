package model

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	interactors "neural_storage/config/adapters/interactors/mock"
	repositories "neural_storage/config/adapters/repositories/mock"
	validator "neural_storage/config/adapters/validator/mock"
	validator2 "neural_storage/cube/adapters/validator/mock"
	r "neural_storage/database/adapters/repositories/mock"
)

type TestSuite struct {
	suite.Suite

	interactor *Interactor

	mockedModelInfo   *r.ModelInfoRepository
	mockedWeightsInfo *r.ModelStructWeightsInfoRepository
	mockedValidator   *validator2.Validator
}

func (s *TestSuite) SetupTest() {
	validatorConf := validator.ValidatorConfig{}
	validatorConf.On("IsMocked").Return(true)

	modelInfoRepoCfg := repositories.ModelInfoRepositoryConfig{}
	modelInfoRepoCfg.On("IsMocked").Return(true)

	modelStructureWeightsInfoRepoCfg := repositories.ModelStructureWeightsInfoRepositoryConfig{}
	modelStructureWeightsInfoRepoCfg.On("IsMocked").Return(true)

	cfg := interactors.ModelInfoInteractorConfig{}
	cfg.On("ModelInfoRepoConfig").Return(&modelInfoRepoCfg)
	cfg.On("ModelStructureWeightInfoRepoConfig").Return(&modelStructureWeightsInfoRepoCfg)
	cfg.On("ValidatorConfig").Return(&validatorConf)

	s.interactor = NewInteractor(&cfg)
	require.NotNil(s.T(), s.interactor)

	s.mockedModelInfo = s.interactor.modelInfo.(*r.ModelInfoRepository)
	s.mockedWeightsInfo = s.interactor.weightsInfo.(*r.ModelStructWeightsInfoRepository)
	s.mockedValidator = s.interactor.validator.(*validator2.Validator)
}

func (s *TestSuite) TearDownTest() {
}
