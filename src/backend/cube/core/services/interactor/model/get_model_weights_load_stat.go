package model

import (
	"context"
	"fmt"
	"neural_storage/cube/core/entities/structure/weights/weightsstat"
	"neural_storage/pkg/logger"
	"time"
)

func (i *Interactor) GetWeightsLoadStat(ctx context.Context, from, to time.Time) ([]*weightsstat.Info, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"from": from, "to": to}).Info("weights loads stat get called")
	if from.After(to) {
		lg.Error("invlaid date period")
		return nil, fmt.Errorf("invalid date period")
	}

	lg.Info("attempt to get weights loads stat info")
	return i.weightsInfo.GetAddStat(from, to)
}
