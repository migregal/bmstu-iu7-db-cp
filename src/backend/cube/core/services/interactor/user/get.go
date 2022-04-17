package user

import (
	"fmt"
	"neural_storage/cube/core/entities/user"
)

func (i *Interactor) Get(info user.Info) (user.Info, error) {
	valid := i.validator.ValidateUserInfo(&info)
	if !valid {
		return user.Info{}, fmt.Errorf("invalid user info")
	}

	return i.userInfo.Get(info)
}
