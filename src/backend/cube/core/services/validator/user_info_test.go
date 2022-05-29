//go:build unit
// +build unit

package validator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	validator "neural_storage/config/adapters/validator/mock"
	"neural_storage/cube/core/entities/user"
)

type UserInfoSuite struct {
	suite.Suite

	conf *validator.ValidatorConfig
}

func (s *UserInfoSuite) SetupTest() {
	s.conf = &validator.ValidatorConfig{}

	s.conf.On("MinUnameLen").Return(2)
	s.conf.On("MaxUnameLen").Return(10)
	s.conf.On("MinPwdLen").Return(8)
	s.conf.On("MaxPwdLen").Return(64)
}

func (s *UserInfoSuite) TearDownTest() {
}

func (s *UserInfoSuite) TestValidateUserInfo() {
	ids := []string{"25892208-5d94-4372-b55e-6e0d4d5d3eaa"}
	usernames := []string{"username9"}
	emails := []string{"test@test.io"}
	passwrods := []string{"4z6TI}j_gP`nyp~;6<Xwxa3-n8Xfq=qp"}
	blocks := []time.Time{time.Now().Add(5 * time.Minute)}

	type fields struct {
		conf *validator.ValidatorConfig
	}
	type args struct {
		info *user.Info
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"empty info",
			fields{s.conf},
			args{user.NewInfo("", "", "", "", "", 0, time.Time{})},
			true,
		},
		{
			"id",
			fields{s.conf},
			args{user.NewInfo(ids[0], "", "", "", "", 0, time.Time{})},
			true,
		},
		{
			"username",
			fields{s.conf},

			args{user.NewInfo("", usernames[0], "", "", "", 0, time.Time{})},
			true,
		},
		{
			"email",
			fields{s.conf},
			args{user.NewInfo("", "", emails[0], "", "", 0, time.Time{})},
			true,
		},
		{
			"password",
			fields{s.conf},
			args{user.NewInfo("", "", "", "", passwrods[0], 0, time.Time{})},
			true,
		},
		{
			"blocked until",
			fields{s.conf},
			args{user.NewInfo("", "", "", "", "", 0, blocks[0])},
			true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			v := NewValidator(tt.fields.conf)
			require.Equal(t, v.ValidateUserInfo(tt.args.info), tt.want)
		})
	}
}

func TestUserInfoSuite(t *testing.T) {
	suite.Run(t, new(UserInfoSuite))
}
