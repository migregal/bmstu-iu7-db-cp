package validator

import (
	"neural_storage/cube/adapters/validator/mock"
	"neural_storage/cube/core/ports/config"
	ports "neural_storage/cube/core/ports/validator"
	"neural_storage/cube/core/services/validator"
)

func NewValidator(conf config.ValidatorConfig) ports.Validator {
	if conf.IsMocked() {
		return &mock.Validator{}
	}
	return validator.NewValidator(conf)
}
