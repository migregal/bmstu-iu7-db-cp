package model

import (
	"context"
	"fmt"
	"neural_storage/cube/core/entities/model/modelstat"
	"neural_storage/pkg/logger"
	"time"
)

func (i *Interactor) GetModelLoadStat(ctx context.Context, from, to time.Time) ([]*modelstat.Info, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"from": from, "to": to}).Info("model loads stat get called")
	if from.After(to) {
		lg.Error("invlaid date period")
		return nil, fmt.Errorf("invalid date period")
	}

	lg.Info("attempt to get models loads stat info")
	return i.modelInfo.GetAddStat(from, to)
}
