//go:build unit
// +build unit

package user

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DeleteSuite struct {
	TestSuite
}

func (s *DeleteSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *DeleteSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *DeleteSuite) TestDelete() {
	s.mockedValidator.On("ValidateUserInfo", mock.Anything).Return(true)
	s.mockedRepo.On("Delete", mock.Anything).Return(nil)

	err := s.interactor.Delete(s.ctx, "b56ee3e1-a5ef-4138-a229-59655a3c66aa")

	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}
