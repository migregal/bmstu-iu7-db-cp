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
	id := "weights 1"
	info := weights.NewInfo(
		id,
		"awesome weights",
		[]*weight.Info{weight.NewInfo(1, 1, 0.1)},
		[]*offset.Info{offset.NewInfo(1, 1, 0.5)},
	)

	s.SqlMock.ExpectBegin()
	s.SqlMock.ExpectExec(`^UPDATE "weights_info" SET .* WHERE id = .*$`).WillReturnResult(sqlmock.NewResult(0, 1))
	s.SqlMock.ExpectExec(`^UPDATE "link_weights" SET .* WHERE id = .*$`).WillReturnResult(sqlmock.NewResult(0, 1))
	s.SqlMock.ExpectExec(`^UPDATE "neuron_offsets" SET .* WHERE id = .*$`).WillReturnResult(sqlmock.NewResult(0, 1))
	s.SqlMock.ExpectCommit()

	err := s.repo.Update(*info)

	require.NoError(s.T(), err)
}

func TestUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateSuite))
}
