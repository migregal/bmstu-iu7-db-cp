//go:build unit
// +build unit

package userinfo

import (
	"neural_storage/cube/core/entities/user"
	"neural_storage/database/core/entities/user_info"
	"neural_storage/database/test/mock/utils"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetSuite struct {
	TestSuite
}

func (s *GetSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *GetSuite) TestGet() {
	id := "test"
	expected := *user.NewInfo(&id, nil, nil, nil, nil, 0, nil)
	res := user_info.UserInfo{ID: id}

	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "users_info" WHERE user_id =`).
		WithArgs(id).
		WillReturnRows(utils.MockRows(res))

	info, err := s.repo.Get(id)

	require.NoError(s.T(), err)
	require.Equal(s.T(), expected, info)
}

func TestGetSuite(t *testing.T) {
	suite.Run(t, new(GetSuite))
}
