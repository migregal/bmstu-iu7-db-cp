package user

import (
	"context"
	"fmt"
	"neural_storage/cube/core/entities/user"
	"neural_storage/pkg/logger"
)

func (i *Interactor) Register(ctx context.Context, info user.Info) (string, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"info": info}).Info("user registration get called")
	valid := i.validator.ValidateUserInfo(&info)
	if !valid {
		lg.Error("invlaid user info")
		return "", fmt.Errorf("invalid user info")
	}

	ninfo, err := i.normalizer.NormalizeUserInfo(info)
	if err != nil {
		lg.Errorf("user info normalization error: %v", err)
		return "", err
	}

	lg.Info("attempt to register user")
	return i.userInfo.Add(ninfo)
}
