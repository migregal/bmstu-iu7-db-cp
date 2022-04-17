package normalizer

import "neural_storage/cube/core/entities/user"

type Normalizer interface {
	NormalizeUserInfo(user.Info) (user.Info, error)
}
