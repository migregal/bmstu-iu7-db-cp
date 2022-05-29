package model

import (
	"context"
	"fmt"
	"neural_storage/cube/core/entities/structure/weights/weightsstat"
	"neural_storage/pkg/logger"
	"time"
)

func (i *Interactor) GetWeightsEditStat(ctx context.Context, from, to time.Time) ([]*weightsstat.Info, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"from": from, "to": to}).Info("weights edits stat get called")
	if from.After(to) {
		lg.Error("invlaid date period")
		return nil, fmt.Errorf("invalid date period")
	}

	lg.Info("attempt to get weights edits stat info")
	return i.weightsInfo.GetUpdateStat(from, to)
}
