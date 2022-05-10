package user

import (
	"fmt"
	"neural_storage/cube/core/entities/user"
)

func (i *Interactor) Register(info user.Info) (string, error) {
	valid := i.validator.ValidateUserInfo(&info)
	if !valid {
		return "", fmt.Errorf("invalid user info")
	}

	ninfo, err := i.normalizer.NormalizeUserInfo(info)
	if err != nil {
		return "", err
	}

	return i.userInfo.Add(ninfo)
}
