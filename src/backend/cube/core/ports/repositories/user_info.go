//go:generate mockery --name=UserInfoRepository --outpkg=mock --output=../../../../database/adapters/repositories/mock/ --filename=user_info_repository.go --structname=UserInfoRepository
package repositories

import "neural_storage/cube/core/entities/user"

type UserInfoRepository interface {
	Add(user.Info) error
	Get(user.Info) (user.Info, error)
	Find(filter UserInfoFilter) ([]*user.Info, error)
	Update(user.Info) error
	Delete(user.Info) error
}

type UserInfoFilter struct {
	UserIds []string
	Usernames []string
	Emails []string
	Limit int
}
