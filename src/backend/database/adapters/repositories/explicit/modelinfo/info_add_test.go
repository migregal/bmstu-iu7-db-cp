//go:build unit
// +build unit

package modelinfo

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/neuron"
	"neural_storage/cube/core/entities/neuron/link"
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/entities/structure/layer"
	"neural_storage/cube/core/entities/structure/weights"
	dbmodel "neural_storage/database/core/entities/model"
	dbneuron "neural_storage/database/core/entities/neuron"
	dbstructure "neural_storage/database/core/entities/structure"
	dblayer "neural_storage/database/core/entities/structure/layer"
	dbweights "neural_storage/database/core/entities/structure/weights"
	"neural_storage/database/test/mock/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AddSuite struct {
	TestSuite
}

func (s *AddSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *AddSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *AddSuite) TestAdd() {
	name := "test"
	info := model.NewInfo(
		name,
		structure.NewInfo(
			"awesome struct",
			[]*neuron.Info{neuron.NewInfo("neuron1", "test")},
			[]*layer.Info{layer.NewInfo("test", "alpha", "beta")},
			[]*link.Info{link.NewInfo("link1", "neuron1", "neuron1")},
			[]*weights.Info{
				weights.NewInfo(
					"weights1",
					[]*weight.Info{weight.NewInfo("weights1", "w1", 0.1)},
					[]*offset.Info{offset.NewInfo("weights1", "o1", 0.5)},
				),
			},
		))

	s.SqlMock.ExpectBegin()
	s.SqlMock.ExpectQuery(`^INSERT INTO "models" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dbmodel.Model{ID: name}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "structures" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dbstructure.Structure{ID: "struct_id"}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "layers" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dblayer.Layer{ID: "layer_id"}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "neurons" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dbneuron.Neuron{ID: "some id for neuron"}))
	s.SqlMock.ExpectExec(`^INSERT INTO "links" .*$`).WillReturnResult(sqlmock.NewResult(1, 0))
	s.SqlMock.ExpectQuery(`^INSERT INTO "weights_info" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dbweights.Weights{ID: "some id for weights"}))
	s.SqlMock.ExpectExec(`^INSERT INTO "offsets" .*$`).WillReturnResult(sqlmock.NewResult(1, 0))
	s.SqlMock.ExpectExec(`^INSERT INTO "weights" .*$`).WillReturnResult(sqlmock.NewResult(1, 0))
	s.SqlMock.ExpectCommit()

	res, err := s.repo.Add(*info)

	require.NoError(s.T(), err)
	require.Equal(s.T(), name, res)
}

func TestAddSuite(t *testing.T) {
	suite.Run(t, new(AddSuite))
}
