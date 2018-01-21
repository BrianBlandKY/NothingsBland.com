package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	LogLevel string `yaml:"log_level"`
	Port     string `yaml:"port"`
}

type Config struct {
	Env    string `yaml:"env"`
	Server Server `yaml:"server"`
}

func ParseConfig(path string) (config Config) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config = Config{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}

	return config
}
