package interactorscfg

import (
	"neural_storage/config/adapters/repositories/repositoriescfg"
	"neural_storage/config/adapters/validator/validatorcfg"
	"neural_storage/cube/core/ports/config"
	port "neural_storage/database/core/ports/config"
)

type ModelInfoInteractorConfig struct {
	RepoConf   repositoriescfg.ModelInfoConfig
	SWRepoConf repositoriescfg.ModelStructureWeightsInfoRepositoryConfig
	Validator  validatorcfg.Config
}

func (c *ModelInfoInteractorConfig) ModelInfoRepoConfig() port.ModelInfoRepositoryConfig {
	return &c.RepoConf
}

func (c *ModelInfoInteractorConfig) ModelStructureWeightInfoRepoConfig() port.ModelStructureWeightsInfoRepositoryConfig {
	return &c.SWRepoConf
}

func (c *ModelInfoInteractorConfig) ValidatorConfig() config.ValidatorConfig {
	return &c.Validator
}
