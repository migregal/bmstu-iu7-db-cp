//go:build unit
// +build unit

package user

import (
	"time"

	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"neural_storage/cube/core/entities/user/userstat"
)

type GetLoadStatSuite struct {
	TestSuite
}

func (s *GetLoadStatSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetLoadStatSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *GetLoadStatSuite) TestGetLoadStat() {
	id := ""
	expected := []*userstat.Info{userstat.New(id, time.Time{}, time.Time{})}
	s.mockedRepo.On("GetAddStat", mock.Anything, mock.Anything).Return(expected, nil)

	info, err := s.interactor.GetUserRegistrationStat(s.ctx, time.Time{}, time.Now())

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, expected)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestGetLoadStatSuite(t *testing.T) {
	suite.Run(t, new(GetLoadStatSuite))
}
