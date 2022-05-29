package user

import (
	"context"
	"fmt"
	"neural_storage/cube/core/entities/user/userstat"
	"neural_storage/pkg/logger"
	"time"
)

func (i *Interactor) GetUserRegistrationStat(ctx context.Context, from, to time.Time) ([]*userstat.Info, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"from": from, "to": to}).Info("user registration stat get called")
	if from.After(to) {
		lg.Error("invlaid date period")
		return nil, fmt.Errorf("invalid date period")
	}

	lg.Info("attempt to get user registration stat")
	return i.userInfo.GetAddStat(from, to)
}
