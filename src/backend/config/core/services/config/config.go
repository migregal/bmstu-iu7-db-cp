package config

import (
	config2 "neural_storage/cache/core/ports/config"
	"neural_storage/cache/core/services/interactors/cache"
	"neural_storage/config/adapters/cache/cachecfg"
)

type dbConn struct {
	Adapter  string `yaml:"adapter,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
	DBName   string `yaml:"db_name,omitempty"`
	Driver   string `yaml:"driver,omitempty"`
}

type ParsebleConfig struct {
	Debug       bool   `yaml:"debug,omitempty"`
	DBConn      dbConn `yaml:"db,omitempty"`
	DBConnStat  dbConn `yaml:"statdb,omitempty"`
	DBConnAdmin dbConn `yaml:"admindb,omitempty"`
	CacheConn   struct {
		Adapter     string `yaml:"adapter,omitempty"`
		Host        string `yaml:"host,omitempty"`
		Port        string `yaml:"port,omitempty"`
		User        string `yaml:"user,omitempty"`
		Password    string `yaml:"password,omitempty"`
		ModelSpace  string `yaml:"model,omitempty"`
		WeightSpace string `yaml:"weight,omitempty"`
	} `yaml:"cache,omitempty"`
	Validation struct {
		MinUserNameLen int `yaml:"min_username_len"`
		MaxUserNameLen int `yaml:"max_username_len"`
		MinPasswordLen int `yaml:"min_pwd_len"`
		MaxPasswordLen int `yaml:"max_pwd_len"`
	} `yaml:"validation"`
	PrivKeyPath string `yaml:"priv_key_path"`
	PubKeyPath  string `yaml:"pub_key_path"`
}

type Config struct {
	conf ParsebleConfig
}

func (c *Config) Debug() bool {
	return c.conf.Debug
}

func (c *Config) Cache() config2.CacheConfig {
	return &cachecfg.CacheConfig{
		CacheAdapter: c.conf.CacheConn.Adapter,
		CacheParams: cache.Params{
			Host:     c.conf.CacheConn.Host,
			Port:     c.conf.CacheConn.Port,
			User:     c.conf.CacheConn.User,
			Password: c.conf.CacheConn.Password,
		},
	}
}

func (c *Config) PubKeyPath() string {
	return c.conf.PubKeyPath
}

func (c *Config) PrivKeyPath() string {
	return c.conf.PrivKeyPath
}
