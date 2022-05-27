package interactorscfg

import (
	"neural_storage/config/adapters/repositories/repositoriescfg"
	"neural_storage/config/adapters/validator/validatorcfg"
	"neural_storage/cube/core/ports/config"
	port "neural_storage/database/core/ports/config"
)

type UserInfoConfig struct {
	RepoConf  repositoriescfg.UserInfoConfig
	Validator validatorcfg.Config
}

func (c *UserInfoConfig) RepoConfig() port.UserInfoRepositoryConfig {
	return &c.RepoConf
}

func (c *UserInfoConfig) ValidatorConfig() config.ValidatorConfig {
	return &c.Validator
}

func (c *UserInfoConfig) NormalizerConfig() config.NormalizerConfig {
	return nil
}
