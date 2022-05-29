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

type UpdateSuite struct {
	TestSuite
}

func (s *UpdateSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *UpdateSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *UpdateSuite) TestUpdate() {
	s.mockedValidator.On("ValidateUserInfo", mock.Anything).Return(true)

	expected := *user.NewInfo("", "", "", "", "", 0, time.Time{})
	s.mockedRepo.On("Update", mock.Anything).Return(nil)

	err := s.interactor.Update(s.ctx, expected)

	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateSuite))
}
