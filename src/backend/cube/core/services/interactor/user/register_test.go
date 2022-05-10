package user

import (
	"neural_storage/cube/core/entities/user"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RegisterSuite struct {
	TestSuite
}

func (s *RegisterSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *RegisterSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *RegisterSuite) TestRegister() {
	s.mockedValidator.On("ValidateUserInfo", mock.Anything).Return(true)

	expected := user.NewInfo(nil, nil, nil, nil, nil, nil)
	s.mockedRepo.On("Add", mock.Anything).Return("hehe", nil)

	id, err := s.interactor.Register(*expected)

	require.NoError(s.T(), err)
	require.Equal(s.T(), "hehe", id)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
}
