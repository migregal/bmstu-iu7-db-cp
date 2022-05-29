//go:build unit
// +build unit

package model

import (
	"time"

	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"neural_storage/cube/core/entities/model/modelstat"
)

type GetEditModelStatSuite struct {
	TestSuite
}

func (s *GetEditModelStatSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetEditModelStatSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *GetEditModelStatSuite) TestGetEditModelStat() {
	id := ""
	expected := []*modelstat.Info{modelstat.New(id, time.Time{}, time.Time{})}
	s.mockedModelInfo.On("GetUpdateStat", mock.Anything, mock.Anything).Return(expected, nil)

	info, err := s.interactor.GetModelEditStat(s.ctx, time.Time{}, time.Now())

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, expected)

	require.True(s.T(), s.mockedModelInfo.AssertExpectations(s.T()))
}

func TestGetEditModelStatSuite(t *testing.T) {
	suite.Run(t, new(GetEditModelStatSuite))
}
