package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

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
	config_path := os.Getenv("CONFIG_PATH")
	if config_path == "" {
		log.Fatal("config path is not set")
	}
	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		log.Fatal("config file does not exist")
	}
	var cfg Config
	if err := cleanenv.ReadConfig(config_path, &cfg); err != nil {
		log.Fatal("can't read config file")

	}
	return &cfg
}
