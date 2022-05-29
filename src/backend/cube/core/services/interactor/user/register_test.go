//go:build unit
// +build unit

package user

import (
	"time"

	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"neural_storage/cube/core/entities/user"
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

	expected := *user.NewInfo("", "", "", "", "", 0, time.Time{})
	s.mockedRepo.On("Add", mock.Anything).Return("hehe", nil)

	id, err := s.interactor.Register(s.ctx, expected)

	require.NoError(s.T(), err)
	require.Equal(s.T(), "hehe", id)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
}
