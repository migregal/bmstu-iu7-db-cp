package config

import (
	"neural_storage/config/adapters/interactors/interactorscfg"
	"neural_storage/config/adapters/repositories/repositoriescfg"
	"neural_storage/config/adapters/validator/validatorcfg"
	"neural_storage/cube/core/ports/config"
	"neural_storage/database/core/services/interactor/database"
)

func (c *Config) AdminUserInfo() config.UserInfoInteractorConfig {
	return c.getUserInfo(c.conf.DBConnAdmin)
}

func (c *Config) StatUserInfo() config.UserInfoInteractorConfig {
	return c.getUserInfo(c.conf.DBConnStat)
}

func (c *Config) UserInfo() config.UserInfoInteractorConfig {
	return c.getUserInfo(c.conf.DBConn)
}

func (c *Config) getUserInfo(conn dbConn) config.UserInfoInteractorConfig {
	return &interactorscfg.UserInfoConfig{
		RepoConf: repositoriescfg.UserInfoConfig{
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
