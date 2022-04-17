package user

import (
	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/ports/repositories"
)

type Filter struct {
	Ids       []string
	Usernames []string
	Emails    []string
	Limit     int
}

func (i *Interactor) Find(filter Filter) ([]*user.Info, error) {
	return i.userInfo.Find(
		repositories.UserInfoFilter{
			UserIds:   filter.Ids,
			Usernames: filter.Usernames,
			Emails:    filter.Emails,
			Limit:     filter.Limit,
		},
	)
}
