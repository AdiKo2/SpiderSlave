package Utils

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DataBase struct {
		Name   string `yaml:"name"`
		Server struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"server"`
		Credentials struct {
			UserName string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"credentials"`
	} `yaml:"database"`
}

func InitConfig() *Config {
	var cfg Config
	readFile(&cfg)
	return &cfg
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		panic("cannot open config")
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		panic("cannot read config")
	}
}
