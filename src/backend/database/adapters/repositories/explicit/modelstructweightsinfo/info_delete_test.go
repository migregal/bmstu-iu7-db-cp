//go:build unit
// +build unit

package modelstructweightsinfo

import (
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure/weights"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
	name := "test"
	info := weights.NewInfo(
		"",
		name,
		[]*weight.Info{weight.NewInfo(1, 1, 10)},
		[]*offset.Info{offset.NewInfo(1, 1, 0.1)},
	)

	s.SqlMock.ExpectBegin()
	s.SqlMock.ExpectExec(`^DELETE FROM "weights_info" WHERE id = .*$`).WillReturnResult(sqlmock.NewResult(0, 1))
	s.SqlMock.ExpectCommit()

	err := s.repo.Delete([]weights.Info{*info})

	require.NoError(s.T(), err)
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}
