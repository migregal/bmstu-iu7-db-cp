package statusers

import (
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat"
)

var (
	statCallGet stat.Counter
	statFailGet stat.Counter
	statOKGet   stat.Counter
)

func init() {
	statCallGet = stat.NewCounter("v1", "cube_users_stat_call_get", "The total number of getting users_stat attempts")
	statFailGet = stat.NewCounter("v1", "cube_users_stat_fail_get", "The total number of getting users_stat fails")
	statOKGet = stat.NewCounter("v1", "cube_users_stat_ok_get", "The total number of getting users_stat")
}

type Handler struct {
	resolver interactors.UserInfoInteractor

	lg *logger.Logger
}

func New(lg *logger.Logger, resolver interactors.UserInfoInteractor) Handler {
	return Handler{resolver: resolver, lg: lg}
}
