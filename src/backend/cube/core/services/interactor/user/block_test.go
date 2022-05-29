//go:build unit
// +build unit

package user

import (
	"time"

	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BlockSuite struct {
	TestSuite
}

func (s *BlockSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *BlockSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *BlockSuite) TestBlock() {
	s.mockedValidator.On("ValidateUserInfo", mock.Anything).Return(true)
	s.mockedRepo.On("Update", mock.Anything).Return(nil)
	err := s.interactor.Block(
		s.ctx,
		"dacc4a61-49a6-487a-afa9-eb1fc37d528c",
		time.Now().Add(60*time.Minute))

	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestBlockSuite(t *testing.T) {
	suite.Run(t, new(BlockSuite))
}
