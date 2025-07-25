package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address" env:"HTTP_SERVER_ADDRESS" env-default:"localhost"`
	Port    string `yaml:"port" env:"HTTP_SERVER_PORT" env-default:"8080"`
}

// env-default:"production"
type Config struct {
	Env          string `yaml:"env" env:"ENV" env-required:"true"`
	Storage_path string `yaml:"storage_path" env-required:"true"`
	HttpServer   `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatalf("config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist %s", configPath)
	}
	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file %s", err.Error())
	}
	return &cfg

}
