package adminblock

import (
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat"
)

var (
	statCallGet stat.Counter
	statFailGet stat.Counter
	statOKGet   stat.Counter

	statCallDelete stat.Counter
	statFailDelete stat.Counter
	statOKDelete   stat.Counter

	statCallUpdate stat.Counter
	statFailUpdate stat.Counter
	statOKUpdate   stat.Counter
)

func init() {
	statCallGet = stat.NewCounter("v1", "cube_blocked_call_get", "The total number of getting blocked users info attempts")
	statFailGet = stat.NewCounter("v1", "cube_blocked_fail_get", "The total number of getting blocked users info fails")
	statOKGet = stat.NewCounter("v1", "cube_blocked_ok_get", "The total number of getting blocked users info")

	statCallDelete = stat.NewCounter("v1", "cube_blocked_call_delete", "The total number of deleting blocked users info attempts")
	statFailDelete = stat.NewCounter("v1", "cube_blocked_fail_delete", "The total number of deleting blocked users info fails attempts")
	statOKDelete = stat.NewCounter("v1", "cube_blocked_ok_delete", "The total number of deleted blocked")

	statCallUpdate = stat.NewCounter("v1", "cube_blocked_call_update", "The total number of updating blocked users info attempts")
	statFailUpdate = stat.NewCounter("v1", "cube_blocked_fail_update", "The total number of updating blocked users info fails")
	statOKUpdate = stat.NewCounter("v1", "cube_blocked_ok_update", "The total number of updated blocked users info")
}

type Handler struct {
	resolver interactors.UserInfoInteractor

	lg *logger.Logger
}

func New(lg *logger.Logger, resolver interactors.UserInfoInteractor) Handler {
	return Handler{resolver: resolver, lg: lg}
}
