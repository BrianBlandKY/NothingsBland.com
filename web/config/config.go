package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Environment     string `yaml:"env"`
	LogLevel        string `yaml:"log_level"`
	Port            string `yaml:"port"`
	AssetsDirectory string `yaml:"assets_directory"`
}

type Config struct {
	Server Server `yaml:"server"`
}

func parseConfig(path string) (config Config) {
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

func GetConfig(path string) (config Config) {
	return parseConfig(path)
}
