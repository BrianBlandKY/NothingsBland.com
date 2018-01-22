package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	LogLevel         string `yaml:"log_level"`
	Port             string `yaml:"port"`
	RebuildTemplates bool   `yaml:"rebuild_templates"`
	EnableCaching    bool   `yaml:"enable_caching"`
	EnableLogging    bool   `yaml:"enable_logging"`
}

type Config struct {
	Env    string `yaml:"env"`
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

func defaultConfig() (config Config) {
	return Config{
		Env: "dev",
		Server: Server{
			LogLevel:         "debug",
			Port:             "8080",
			RebuildTemplates: true,
			EnableCaching:    false,
			EnableLogging:    false,
		},
	}
}

func BuildConfig(path string) (config Config) {
	if len(path) > 0 {
		config = parseConfig(path)
	}

	// How to merge two objects?
	defaultConfig := defaultConfig()

	return defaultConfig
}
