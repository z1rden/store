package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	AppName             string `yaml:"app_name"`
	HttpPort            string `yaml:"http_port"`
	LomsServiceGrpcHost string `yaml:"loms_service_grpc_host"`
}

func MustLoad() *Config {
	configPath := "./config/config_cart.yaml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config path does not exist: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err)
	}

	// TODO подумать, как красиво сделать с хостом
	lomsGrpcHost := os.Getenv("LOMS_SERVICE_GRPC_HOST")
	if lomsGrpcHost != "" {
		cfg.LomsServiceGrpcHost = lomsGrpcHost
	}

	return &cfg
}
