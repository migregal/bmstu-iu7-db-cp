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
	id := "test"
	expected := *user.NewInfo(&id, nil, nil, nil, nil, 0, nil)

	s.SqlMock.
		ExpectExec(`^UPDATE "users_info" SET`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := s.repo.Update(expected)

	require.NoError(s.T(), err)
}

func TestUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateSuite))
}
