//go:build unit
// +build unit

package normalizer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	normalizer "neural_storage/config/adapters/normalizer/mock"

	"neural_storage/cube/core/entities/user"
)

type UserInfoSuite struct {
	suite.Suite

	conf *normalizer.NormalizerConfig
}

func (s *UserInfoSuite) SetupTest() {
	s.conf = &normalizer.NormalizerConfig{}
}

func (s *UserInfoSuite) TearDownTest() {
}

func (s *UserInfoSuite) TestValidateUserInfo() {
	pwd := "something good"
	expected := "ad35ea933f9868ba42c74e6f6c053a870d6ff78d2a722161cfb0c6809f2e4d3b"

	info := user.NewInfo("", "", "", "", pwd, 0, time.Time{})

	n := NewNormalizer(s.conf)

	i, err := n.NormalizeUserInfo(*info)

	require.NoError(s.T(), err)
	require.Equal(s.T(), i.Pwd(), expected)
}

func TestUserInfoSuite(t *testing.T) {
	suite.Run(t, new(UserInfoSuite))
}
