package user

import (
	"context"
	"fmt"
	"neural_storage/cube/core/entities/user"
	"neural_storage/pkg/logger"
	"time"
)

func (i *Interactor) Get(ctx context.Context, id string) (user.Info, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"id": id}).Info("user get called")
	valid := i.validator.ValidateUserInfo(user.NewInfo(id, "", "", "", "", 0, time.Time{}))
	if !valid {
		lg.Error("invlaid user info")
		return user.Info{}, fmt.Errorf("invalid user info")
	}

	lg.Info("attempt to get user")
	return i.userInfo.Get(id)
}
