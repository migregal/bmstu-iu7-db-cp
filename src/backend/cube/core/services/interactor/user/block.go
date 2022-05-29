package user

import (
	"context"
	"fmt"
	"neural_storage/cube/core/entities/user"
	"neural_storage/pkg/logger"
	"time"
)

func (i *Interactor) Block(ctx context.Context, userId string, until time.Time) error {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"id": userId, "until": until}).Info("user block called")
	info := user.NewInfo(userId, "", "", "", "", 0, until)
	valid := i.validator.ValidateUserInfo(info)
	if !valid {
		lg.Error("invlaid user info")
		return fmt.Errorf("invalid user info")
	}

	lg.Info("attempt to block user")
	return i.userInfo.Update(*info)
}
