package user

import (
	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/core/ports/repositories"
)

func (i *Interactor) Find(filter interactors.UserInfoFilter) ([]user.Info, error) {
	return i.userInfo.Find(
		repositories.UserInfoFilter{
			UserIds:   filter.Ids,
			Usernames: filter.Usernames,
			Emails:    filter.Emails,
			Limit:     filter.Limit,
			Offset:    filter.Offset,
		},
	)
}
