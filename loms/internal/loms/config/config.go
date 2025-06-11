package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	AppName     string `yaml:"app_name"`
	HttpPort    string `yaml:"http_port"`
	GrpcPort    string `yaml:"grpc_port"`
	MasterDBURL string `yaml:"master_db_url"`
	SyncDBURL   string `yaml:"sync_db_url"`
}

func MustLoad() *Config {
	configPath := "./config/config_loms.yaml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config path does not exist: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err)
	}

	// TODO подумать, как красиво сделать с хостом
	masterDBURL := os.Getenv("MASTER_DB_URL")
	if masterDBURL != "" {
		cfg.MasterDBURL = masterDBURL
	}

	// TODO подумать, как красиво сделать с хостом
	syncDBURL := os.Getenv("SYNC_DB_URL")
	if syncDBURL != "" {
		cfg.SyncDBURL = syncDBURL
	}

	return &cfg
}
