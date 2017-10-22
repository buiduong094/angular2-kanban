package config

import (
	"github.com/jinzhu/configor"
	"sync"
)

var once sync.Once
var config *Config

var ConfigPaths = []string{"server/config/config.yml", "server/config/default_config.yml"}

func Get() *Config {
	once.Do(func() {
		config = &Config{}
	})
	return config
}

func init() {
	newConfig := &Config{
		Host:        Get().Host,
		Port:        Get().Port,
		Environment: Get().Environment,
	}

	config = newConfig
	configor.Load(config, ConfigPaths...)
}
