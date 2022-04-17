package user

import (
	"fmt"
	"neural_storage/cube/core/entities/user"
)

func (i *Interactor) Update(info user.Info) error {
	valid := i.validator.ValidateUserInfo(&info)
	if !valid {
		return fmt.Errorf("invalid user info")
	}

	return i.userInfo.Update(info)
}
