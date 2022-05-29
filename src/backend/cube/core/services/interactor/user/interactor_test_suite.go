//go:build unit
// +build unit

package user

import (
	"context"
	"io/ioutil"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"neural_storage/pkg/logger"

	interactors "neural_storage/config/adapters/interactors/mock"
	normalizer "neural_storage/config/adapters/normalizer/mock"
	repositories "neural_storage/config/adapters/repositories/mock"
	validator "neural_storage/config/adapters/validator/mock"
	validator2 "neural_storage/cube/adapters/validator/mock"

	repo "neural_storage/database/adapters/repositories/mock"
)

type TestSuite struct {
	suite.Suite

	interactor *Interactor

	mockedRepo      *repo.UserInfoRepository
	mockedValidator *validator2.Validator

	ctx context.Context
}

func (s *TestSuite) SetupTest() {
	normalizerConf := normalizer.NormalizerConfig{}

	validatorConf := validator.ValidatorConfig{}
	validatorConf.On("IsMocked").Return(true)
	validatorConf.On("MinUnameLen").Return(2)
	validatorConf.On("MaxUnameLen").Return(20)

	repoCfg := repositories.UserInfoRepositoryConfig{}
	repoCfg.On("IsMocked").Return(true)

	cfg := interactors.UserInfoInteractorConfig{}
	cfg.On("RepoConfig").Return(&repoCfg)
	cfg.On("ValidatorConfig").Return(&validatorConf)
	cfg.On("NormalizerConfig").Return(&normalizerConf)

	s.ctx = context.Background()

	lg := logger.New()
	lg.SetOutput(ioutil.Discard)

	s.interactor = NewInteractor(lg, &cfg)
	require.NotNil(s.T(), s.interactor)

	s.mockedRepo = s.interactor.userInfo.(*repo.UserInfoRepository)
	s.mockedValidator = s.interactor.validator.(*validator2.Validator)
}

func (s *TestSuite) TearDownTest() {
}
