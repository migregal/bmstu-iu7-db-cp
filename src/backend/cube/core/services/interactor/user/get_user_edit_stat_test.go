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

type GetEditStatSuite struct {
	TestSuite
}

func (s *GetEditStatSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetEditStatSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *GetEditStatSuite) TestGetEditStat() {
	id := ""
	expected := []*userstat.Info{userstat.New(id, time.Time{}, time.Time{})}
	s.mockedRepo.On("GetUpdateStat", mock.Anything, mock.Anything).Return(expected, nil)

	info, err := s.interactor.GetUserEditStat(s.ctx, time.Time{}, time.Now())

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, expected)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestGetEditStatSuite(t *testing.T) {
	suite.Run(t, new(GetEditStatSuite))
}
