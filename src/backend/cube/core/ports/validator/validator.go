//go:generate mockery --name=Validator --outpkg=mock --output=../../../adapters/validator/mock/ --filename=validator.go --structname=Validator

package validator

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/user"
)

type Validator interface {
	ValidateUserInfo(info *user.Info) bool

	ValidateModelInfo(info *model.Info) error
}
