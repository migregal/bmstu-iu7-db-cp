//go:generate mockery --name=UserInfoRepository --outpkg=mock --output=../../../../database/adapters/repositories/mock/ --filename=user_info_repository.go --structname=UserInfoRepository
package repositories

import (
	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/entities/user/userstat"
	"time"
)

type UserInfoRepository interface {
	Add(user.Info) (string, error)
	Get(id string) (user.Info, error)
	Find(filter UserInfoFilter) ([]user.Info, error)
	Update(user.Info) error
	Delete(user.Info) error

	GetAddStat(from, to time.Time) ([]*userstat.Info, error)
	GetUpdateStat(from, to time.Time) ([]*userstat.Info, error)
}

type UserInfoFilter struct {
	UserIds   []string
	Usernames []string
	Emails    []string
	Offset    int
	Limit     int
}
