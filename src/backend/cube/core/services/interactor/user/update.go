package user

import (
	"context"
	"fmt"
	"neural_storage/cube/core/entities/user"
	"neural_storage/pkg/logger"
)

func (i *Interactor) Update(ctx context.Context, info user.Info) error {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"info": info}).Info("user update get called")
	valid := i.validator.ValidateUserInfo(&info)
	if !valid {
		lg.Error("invlaid user info")
		return fmt.Errorf("invalid user info")
	}

	lg.Info("attempt to update user")
	return i.userInfo.Update(info)
}
