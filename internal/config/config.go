package config

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v3"
)

var configPath = "./config/local.yaml"

type Config struct {
	Env         string     `yaml:"env"`
	StoragePath string     `yaml:"storage_path"`
	HttpServer  HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func MustReadConfigFile() *Config {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("config file does not exist: %w", err)
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg
}
