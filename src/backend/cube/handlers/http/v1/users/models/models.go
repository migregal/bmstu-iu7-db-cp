package models

import (
	"neural_storage/cube/core/ports/cache"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat"
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
	statCallAdd = stat.NewCounter("v1", "cube_models_call_add", "The total number of adding model attempts")
	statFailAdd = stat.NewCounter("v1", "cube_models_fail_add", "The total number of adding model fails")
	statOKAdd = stat.NewCounter("v1", "cube_models_ok_add", "The total number of added attempts")

	statCallGet = stat.NewCounter("v1", "cube_models_call_get", "The total number of getting model attempts")
	statFailGet = stat.NewCounter("v1", "cube_models_fail_get", "The total number of getting model fails")
	statOKGet = stat.NewCounter("v1", "cube_models_ok_get", "The total number of getting models")

	statCallDelete = stat.NewCounter("v1", "cube_models_call_delete", "The total number of deleting model attempts")
	statFailDelete = stat.NewCounter("v1", "cube_models_fail_delete", "The total number of deleting model fails attempts")
	statOKDelete = stat.NewCounter("v1", "cube_models_ok_delete", "The total number of deleted models")

	statCallUpdate = stat.NewCounter("v1", "cube_models_call_update", "The total number of updating maodels attempts")
	statFailUpdate = stat.NewCounter("v1", "cube_models_fail_update", "The total number of updating models fails")
	statOKUpdate = stat.NewCounter("v1", "cube_models_ok_update", "The total number of updated model")
}

type Handler struct {
	resolver interactors.NeuralNetworkInteractor
	cache    cache.CacheInteractor

	lg *logger.Logger
}

func New(lg *logger.Logger, resolver interactors.NeuralNetworkInteractor, cache cache.CacheInteractor) Handler {
	return Handler{resolver: resolver, cache: cache, lg: lg}
}
