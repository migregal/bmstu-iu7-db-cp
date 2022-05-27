package repositoriescfg

import "neural_storage/database/core/services/interactor/database"

type UserInfoConfig struct {
	DBAdapter string
	DBParams  database.Params
}

func (c *UserInfoConfig) IsMocked() bool {
	return false
}

func (c *UserInfoConfig) Adapter() string {
	return c.DBAdapter
}

func (c *UserInfoConfig) ConnParams() database.Params {
	return c.DBParams
}
