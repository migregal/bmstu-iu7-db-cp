package repositoriescfg

import "neural_storage/database/core/services/interactor/database"

type ModelInfoConfig struct {
	DBAdapter string
	DBParams  database.Params
}

func (c *ModelInfoConfig) IsMocked() bool {
	return false
}

func (c *ModelInfoConfig) Adapter() string {
	return c.DBAdapter
}

func (c *ModelInfoConfig) ConnParams() database.Params {
	return c.DBParams
}

type ModelStructureWeightsInfoRepositoryConfig struct {
	DBAdapter string
	DBParams  database.Params
}

func (c *ModelStructureWeightsInfoRepositoryConfig) IsMocked() bool {
	return false
}

func (c *ModelStructureWeightsInfoRepositoryConfig) Adapter() string {
	return c.DBAdapter
}

func (c *ModelStructureWeightsInfoRepositoryConfig) ConnParams() database.Params {
	return c.DBParams
}
