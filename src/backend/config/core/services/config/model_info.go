package config

import (
	"neural_storage/config/adapters/interactors/interactorscfg"
	"neural_storage/config/adapters/repositories/repositoriescfg"
	"neural_storage/config/adapters/validator/validatorcfg"
	"neural_storage/cube/core/ports/config"
	"neural_storage/database/core/services/interactor/database"
)

func (c *Config) AdminModelInfo() config.ModelInfoInteractorConfig {
	return c.getModelInfo(c.conf.DBConnAdmin)
}

func (c *Config) StatModelInfo() config.ModelInfoInteractorConfig {
	return c.getModelInfo(c.conf.DBConnStat)
}

func (c *Config) ModelInfo() config.ModelInfoInteractorConfig {
	return c.getModelInfo(c.conf.DBConn)
}

func (c *Config) getModelInfo(conn dbConn) config.ModelInfoInteractorConfig {
	return &interactorscfg.ModelInfoInteractorConfig{
		RepoConf: repositoriescfg.ModelInfoConfig{
			DBAdapter: conn.Adapter,
			DBParams: database.Params{
				Host:     conn.Host,
				Port:     conn.Port,
				User:     conn.User,
				Password: conn.Password,
				DBName:   conn.DBName,
				Driver:   conn.Driver,
			},
		},
		SWRepoConf: repositoriescfg.ModelStructureWeightsInfoRepositoryConfig{
			DBAdapter: conn.Adapter,
			DBParams: database.Params{
				Host:     conn.Host,
				Port:     conn.Port,
				User:     conn.User,
				Password: conn.Password,
				DBName:   conn.DBName,
				Driver:   conn.Driver,
			},
		},
		Validator: validatorcfg.Config{
			MinUserNameLen: c.conf.Validation.MinUserNameLen,
			MaxUserNameLen: c.conf.Validation.MaxUserNameLen,
			MinPasswordLen: c.conf.Validation.MinPasswordLen,
			MaxPasswordLen: c.conf.Validation.MaxPasswordLen,
		},
	}
}
