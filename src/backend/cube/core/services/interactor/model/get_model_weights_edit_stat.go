package model

import (
	"fmt"
	"neural_storage/cube/core/entities/structure/weights/weightsstat"
	"time"
)

func (i *Interactor) GetWeightsEditStat(from, to time.Time) ([]*weightsstat.Info, error) {
	if from.Before(to) {
		return nil, fmt.Errorf("invalid date period")
	}

	return i.weightsInfo.GetUpdateStat(from, to)
}
