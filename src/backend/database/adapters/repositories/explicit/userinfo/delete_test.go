//go:build unit
// +build unit

package userinfo

import (
	"neural_storage/cube/core/entities/user"
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

func (s *DeleteSuite) TestAdd() {
	id := "test"
	expected := *user.NewInfo(&id, nil, nil, nil, nil, nil)

	s.SqlMock.
		ExpectExec(`^DELETE FROM "users_info" WHERE "users_info"."user_id"`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := s.repo.Delete(expected)

	require.NoError(s.T(), err)
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}
