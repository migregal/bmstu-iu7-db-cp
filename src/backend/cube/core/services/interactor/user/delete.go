package user

import (
	"fmt"
	"neural_storage/cube/core/entities/user"
)

func (i *Interactor) Delete(userId string) error {
	info := user.NewInfo(&userId, nil, nil, nil, nil)
	valid := i.validator.ValidateUserInfo(info)
	if !valid {
		return fmt.Errorf("invalid user info")
	}

	return i.userInfo.Delete(*info)
}
