package gobot

import (
	"github.com/qianlnk/config"
)

const (
	ConfigFile = "config.yaml"
)

type Config struct {
	Tuling Tuling `yaml:"tuling"`
}

type Rebot struct {
	Name string `yaml: "Name"`
	Key  string `yaml:"key"`
}
type Tuling struct {
	URL  string           `yaml:"url"`
	Keys map[string]Rebot `yaml:"keys"`
}

func Load() Config {
	var cfg Config
	if err := config.Parse(&cfg, config.GetConfigAbsolutePath(ConfigFile)); err != nil {
		panic(err)
	}

	return cfg
}
