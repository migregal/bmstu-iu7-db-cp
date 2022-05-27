package user

import (
	"fmt"
	"neural_storage/cube/core/entities/user"
	"time"
)

func (i *Interactor) Get(id string) (user.Info, error) {
	valid := i.validator.ValidateUserInfo(user.NewInfo(&id, nil, nil, nil, nil, 0, time.Time{}))
	if !valid {
		return user.Info{}, fmt.Errorf("invalid user info")
	}

	return i.userInfo.Get(id)
}
