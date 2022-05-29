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

	id := ""
	expected := user.NewInfo(id, "", "", "", "", 0, time.Time{})
	s.mockedRepo.On("Get", mock.Anything).Return(*expected, nil)

	info, err := s.interactor.Get(s.ctx, expected.ID())

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, *expected)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestGetSuite(t *testing.T) {
	suite.Run(t, new(GetSuite))
}
