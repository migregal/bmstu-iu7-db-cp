//go:build unit
// +build unit

package model

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
	modelId := ""
	s.mockedModelInfo.On("Delete", mock.Anything).Return(nil)

	err := s.interactor.Delete(s.ctx, "", modelId)

	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedModelInfo.AssertExpectations(s.T()))
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}
