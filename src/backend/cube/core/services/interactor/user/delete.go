package user

import (
	"fmt"
	"neural_storage/cube/core/entities/user"
	"time"
)

func (i *Interactor) Delete(userId string) error {
	info := user.NewInfo(&userId, nil, nil, nil, nil, 0, time.Time{})
	valid := i.validator.ValidateUserInfo(info)
	if !valid {
		return fmt.Errorf("invalid user info")
	}

	return i.userInfo.Delete(*info)
}
