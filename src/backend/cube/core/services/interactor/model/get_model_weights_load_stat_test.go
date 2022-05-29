//go:build unit
// +build unit

package model

import (
	"time"

	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"neural_storage/cube/core/entities/structure/weights/weightsstat"
)

type GetLoadWeightsStatSuite struct {
	TestSuite
}

func (s *GetLoadWeightsStatSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetLoadWeightsStatSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *GetLoadWeightsStatSuite) TestGetLoadWeightsStat() {
	id := ""
	expected := []*weightsstat.Info{weightsstat.New(id, time.Time{}, time.Time{})}
	s.mockedWeightsInfo.On("GetAddStat", mock.Anything, mock.Anything).Return(expected, nil)

	info, err := s.interactor.GetWeightsLoadStat(s.ctx, time.Time{}, time.Now())

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, expected)

	require.True(s.T(), s.mockedWeightsInfo.AssertExpectations(s.T()))
}

func TestGetLoadWeightsStatSuite(t *testing.T) {
	suite.Run(t, new(GetLoadWeightsStatSuite))
}
