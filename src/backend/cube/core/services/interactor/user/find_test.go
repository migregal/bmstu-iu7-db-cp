//go:build unit
// +build unit

package user

import (
	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/ports/interactors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FindSuite struct {
	TestSuite
}

func (s *FindSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *FindSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *FindSuite) TestFind() {
	filter := interactors.UserInfoFilter{}
	expected := []user.Info{
		*user.NewInfo(nil, nil, nil, nil, nil, 0, nil),
	}

	s.mockedRepo.On("Find", mock.Anything).Return(expected, nil)
	info, err := s.interactor.Find(filter)

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, expected)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestFindSuite(t *testing.T) {
	suite.Run(t, new(FindSuite))
}
