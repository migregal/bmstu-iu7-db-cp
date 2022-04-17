package user

import (
	"fmt"
	"neural_storage/cube/core/entities/user"
	"time"
)

func (i *Interactor) Block(userId string, until time.Time) error {
	info := user.NewInfo(&userId, nil, nil, nil, &until)
	valid := i.validator.ValidateUserInfo(info)
	if !valid {
		return fmt.Errorf("invalid user info")
	}

	return i.userInfo.Update(*info)
}
