package model

import (
	"fmt"
	"neural_storage/cube/core/entities/model/modelstat"
	"time"
)

func (i *Interactor) GetModelLoadStat(from, to time.Time) ([]*modelstat.Info, error) {
		if from.Before(to) {
		return nil, fmt.Errorf("invalid date period")
	}

	return i.modelInfo.GetAddStat(from, to)
}
