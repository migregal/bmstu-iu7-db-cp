package user

import (
	"neural_storage/cube/core/entities/user"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetSuite struct {
	TestSuite
}

func (s *GetSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *GetSuite) TestGet() {
	s.mockedValidator.On("ValidateUserInfo", mock.Anything).Return(true)

	expected := user.NewInfo(nil, nil, nil, nil, nil)
	s.mockedRepo.On("Get", mock.Anything).Return(*expected, nil)

	info, err := s.interactor.Get(*expected)

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, *expected)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestGetSuite(t *testing.T) {
	suite.Run(t, new(GetSuite))
}
