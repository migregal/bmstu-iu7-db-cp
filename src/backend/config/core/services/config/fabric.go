package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func New(confFile string) (Config, error) {
	yfile, err := ioutil.ReadFile(confFile)
	if err != nil {
		return Config{}, err
	}

	data := ParsebleConfig{}

	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		return Config{}, err
	}

	return Config{conf: data}, nil
}
