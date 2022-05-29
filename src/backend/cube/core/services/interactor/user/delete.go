package user

import (
	"context"
	"fmt"
	"neural_storage/cube/core/entities/user"
	"neural_storage/pkg/logger"
	"time"
)

func (i *Interactor) Delete(ctx context.Context, userId string) error {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"id": userId}).Info("user delete called")
	info := user.NewInfo(userId, "", "", "", "", 0, time.Time{})
	valid := i.validator.ValidateUserInfo(info)
	if !valid {
		return fmt.Errorf("invalid user info")
	}

	lg.Info("attempt to delete user")
	return i.userInfo.Delete(*info)
}
