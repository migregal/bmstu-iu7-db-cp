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
	dblink "neural_storage/database/core/entities/neuron/link"
	dboffset "neural_storage/database/core/entities/neuron/offset"
	dbstructure "neural_storage/database/core/entities/structure"
	dblayer "neural_storage/database/core/entities/structure/layer"
	dblw "neural_storage/database/core/entities/structure/weight"
	dbweights "neural_storage/database/core/entities/structure/weights"
	"neural_storage/database/test/mock/utils"
	"testing"

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
		"",
		"",
		name,
		structure.NewInfo(
			"",
			"awesome struct",
			[]*neuron.Info{neuron.NewInfo(1, 1)},
			[]*layer.Info{layer.NewInfo(1, "alpha", "beta")},
			[]*link.Info{link.NewInfo(1, 1, 1)},
			[]*weights.Info{
				weights.NewInfo(
					"",
					"weights1",
					[]*weight.Info{weight.NewInfo(1, 1, 0.1)},
					[]*offset.Info{offset.NewInfo(1, 1, 0.5)},
				),
			},
		))

	s.SqlMock.ExpectBegin()
	s.SqlMock.ExpectQuery(`^INSERT INTO "models" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dbmodel.Model{ID: name}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "structures" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dbstructure.Structure{ID: "struct_id"}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "layers" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dblayer.Layer{ID: 1}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "neurons" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dbneuron.Neuron{ID: 1}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "neuron_links" .*$`).WillReturnRows(utils.MockRows(dblink.Link{ID: 1}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "weights_info" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dbweights.Weights{ID: 1}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "neuron_offsets" .*$`).WillReturnRows(utils.MockRows(dboffset.Offset{ID: 1}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "link_weights" .*$`).WillReturnRows(utils.MockRows(dblw.Weight{ID: 1}))
	s.SqlMock.ExpectCommit()

	res, err := s.repo.Add(*info)

	require.NoError(s.T(), err)
	require.Equal(s.T(), name, res)
}

func TestAddSuite(t *testing.T) {
	suite.Run(t, new(AddSuite))
}
