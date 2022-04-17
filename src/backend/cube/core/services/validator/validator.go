package validator

import (
	"neural_storage/cube/core/ports/config"
	ports "neural_storage/cube/core/ports/validator"
)

type Validator struct {
	conf config.ValidatorConfig
}

func NewValidator(conf config.ValidatorConfig) ports.Validator {
	return &Validator{conf}
}
