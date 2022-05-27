package user

import (
	"fmt"
	"neural_storage/cube/core/entities/user/userstat"
	"time"
)

func (i *Interactor) GetUserEditStat(from, to time.Time) ([]*userstat.Info, error) {
	if from.Before(to) {
		return nil, fmt.Errorf("invalid date period")
	}

	return i.userInfo.GetUpdateStat(from, to)
}
