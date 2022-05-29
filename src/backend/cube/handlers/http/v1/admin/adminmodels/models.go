package adminmodels

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
)

func init() {
	statCallGet = stat.NewCounter("v1", "cube_admin_models_call_get", "The total number of getting admin models attempts")
	statFailGet = stat.NewCounter("v1", "cube_admin_models_fail_get", "The total number of getting admin models fails")
	statOKGet = stat.NewCounter("v1", "cube_admin_models_ok_get", "The total number of getting admin models")

	statCallDelete = stat.NewCounter("v1", "cube_admin_models_call_delete", "The total number of deleting admin models attempts")
	statFailDelete = stat.NewCounter("v1", "cube_admin_models_fail_delete", "The total number of deleting admin models fails attempts")
	statOKDelete = stat.NewCounter("v1", "cube_admin_models_ok_delete", "The total number of deleted admin models")
}

type Handler struct {
	resolver interactors.NeuralNetworkInteractor

	lg *logger.Logger
}

func New(lg *logger.Logger, resolver interactors.NeuralNetworkInteractor) Handler {
	return Handler{resolver: resolver, lg: lg}
}
