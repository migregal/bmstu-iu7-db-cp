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
	"neural_storage/cube/core/ports/repositories"
	dbmodel "neural_storage/database/core/entities/model"
	dbneuron "neural_storage/database/core/entities/neuron"
	dblink "neural_storage/database/core/entities/neuron/link"
	dboffset "neural_storage/database/core/entities/neuron/offset"
	dbstructure "neural_storage/database/core/entities/structure"
	dblayer "neural_storage/database/core/entities/structure/layer"
	dbweight "neural_storage/database/core/entities/structure/weight"
	dbweights "neural_storage/database/core/entities/structure/weights"
	"neural_storage/database/test/mock/utils"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FindSuite struct {
	TestSuite
}

func (s *FindSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *FindSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *FindSuite) TestFind() {
	name := "test"
	info := *model.NewInfo(
		name,
		"",
		"",
		structure.NewInfo(
			"",
			"awesome struct",
			[]*neuron.Info{neuron.NewInfo(0, 0)},
			[]*layer.Info{layer.NewInfo(0, "alpha", "beta")},
			[]*link.Info{link.NewInfo(0, 0, 0)},
			[]*weights.Info{
				weights.NewInfo(
					"weights 1",
					name,
					[]*weight.Info{weight.NewInfo(0, 0, 0.1)},
					[]*offset.Info{offset.NewInfo(0, 0, 0.5)},
				),
			},
		))

	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "models" WHERE id in .* LIMIT 10$`).
		WillReturnRows(utils.MockRows(dbmodel.Model{ID: name, Name: info.Name()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "models" WHERE id = .* ORDER BY .* LIMIT 1$`).
		WillReturnRows(utils.MockRows(dbmodel.Model{ID: name, Name: info.Name()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "structures" WHERE model_id = .* ORDER BY .* LIMIT 1$`).
		WillReturnRows(utils.MockRows(dbstructure.Structure{
			ID:   info.Structure().ID(),
			Name: info.Structure().Name()}))
	s.SqlMock.
		ExpectQuery(`SELECT \* FROM "layers" WHERE structure_id = .*`).
		WillReturnRows(utils.MockRows(dblayer.Layer{
			ID:             info.Structure().Layers()[0].ID(),
			StructureID:    info.Structure().ID(),
			LimitFunc:      info.Structure().Layers()[0].LimitFunc(),
			ActivationFunc: info.Structure().Layers()[0].ActivationFunc()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "neurons" WHERE layer_id in .*$`).
		WillReturnRows(utils.MockRows(dbneuron.Neuron{
			ID:      info.Structure().Neurons()[0].ID(),
			LayerID: info.Structure().Neurons()[0].LayerID()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "neuron_links" WHERE from_id in .*$`).
		WillReturnRows(utils.MockRows(dblink.Link{
			ID:   info.Structure().Links()[0].ID(),
			From: info.Structure().Links()[0].From(),
			To:   info.Structure().Links()[0].To()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "weights_info" WHERE structure_id = .*$`).
		WillReturnRows(utils.MockRows(dbweights.Weights{
			InnerID:   info.Structure().Weights()[0].ID(),
			Name: info.Structure().Weights()[0].Name()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "neuron_offsets" WHERE weights_info_id = .*$`).
		WillReturnRows(utils.MockRows(dboffset.Offset{
			InnerWeights: info.Structure().Weights()[0].ID(),
			ID:      info.Structure().Weights()[0].Offsets()[0].ID(),
			Neuron:  info.Structure().Weights()[0].Offsets()[0].NeuronID(),
			Offset:  info.Structure().Weights()[0].Offsets()[0].Offset()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "link_weights" WHERE weights_info_id = .*$`).
		WillReturnRows(utils.MockRows(dbweight.Weight{
			InnerWeightsID: info.Structure().Weights()[0].ID(),
			ID:        info.Structure().Weights()[0].Weights()[0].ID(),
			LinkID:    info.Structure().Weights()[0].Weights()[0].LinkID(),
			Value:     info.Structure().Weights()[0].Weights()[0].Weight()}))

	res, err := s.repo.Find(repositories.ModelInfoFilter{Ids: []string{name}, Limit: 10})

	require.NoError(s.T(), err)
	require.Equal(s.T(), []*model.Info{&info}, res)
}

func TestFindSuite(t *testing.T) {
	suite.Run(t, new(FindSuite))
}
