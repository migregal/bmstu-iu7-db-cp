package weights

import (
	"neural_storage/cube/core/ports/cache"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat"
)

const (
	modelStorage = "model"
	weightStorage = "weight"
)

var (
	statCallAdd stat.Counter
	statFailAdd stat.Counter
	statOKAdd   stat.Counter

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
	statCallAdd = stat.NewCounter("v1", "cube_weights_call_add", "The total number of adding weights attempts")
	statFailAdd = stat.NewCounter("v1", "cube_weights_fail_add", "The total number of adding weights fails")
	statOKAdd = stat.NewCounter("v1", "cube_weights_ok_add", "The total number of added attempts")

	statCallGet = stat.NewCounter("v1", "cube_weights_call_get", "The total number of getting weights attempts")
	statFailGet = stat.NewCounter("v1", "cube_weights_fail_get", "The total number of getting weights fails")
	statOKGet = stat.NewCounter("v1", "cube_weights_ok_get", "The total number of getting weights")

	statCallDelete = stat.NewCounter("v1", "cube_weights_call_delete", "The total number of deleting weights attempts")
	statFailDelete = stat.NewCounter("v1", "cube_weights_fail_delete", "The total number of deleting weights fails attempts")
	statOKDelete = stat.NewCounter("v1", "cube_weights_ok_delete", "The total number of deleted weights")

	statCallUpdate = stat.NewCounter("v1", "cube_weights_call_update", "The total number of updating maodels attempts")
	statFailUpdate = stat.NewCounter("v1", "cube_weights_fail_update", "The total number of updating weights fails")
	statOKUpdate = stat.NewCounter("v1", "cube_weights_ok_update", "The total number of updated weights")
}

type Handler struct {
	resolver interactors.NeuralNetworkInteractor
	cache    cache.CacheInteractor

	lg *logger.Logger
}

func New(lg *logger.Logger, resolver interactors.NeuralNetworkInteractor, cache cache.CacheInteractor) Handler {
	return Handler{resolver: resolver, cache: cache, lg: lg}
}
