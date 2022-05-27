//go:build unit
// +build unit

package userinfo

import (
	"neural_storage/cube/core/entities/user"
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
	id := "test"
	expected := user.NewInfo(&id, nil, nil, nil, nil, 0, nil)
	info := user.NewInfo(&id, nil, nil, nil, nil, 0, nil)

	s.SqlMock.ExpectQuery(`^INSERT INTO "users_info"`).WillReturnRows(utils.MockRows(*expected))
	res, err := s.repo.Add(*info)

	require.NoError(s.T(), err)
	require.Equal(s.T(), *expected, *info)
	require.Equal(s.T(), *expected.ID(), res)
}

func TestAddSuite(t *testing.T) {
	suite.Run(t, new(AddSuite))
}
