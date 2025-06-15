package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Kafka       Kafka  `yaml:"kafka"`
	AppName     string `yaml:"app_name"`
	HttpPort    string `yaml:"http_port"`
	GrpcPort    string `yaml:"grpc_port"`
	MasterDBURL string `yaml:"master_db_url" env:"MASTER_DB_URL"`
	SyncDBURL   string `yaml:"sync_db_url" env:"SYNC_DB_URL"`
}
type Kafka struct {
	Addr     string   `yaml:"addr" env:"KAFKA_ADDR"`
	Topic    Topic    `yaml:"topic"`
	Producer Producer `yaml:"producer"`
}

type Topic struct {
	Name              string `yaml:"name"`
	NumPartitions     int    `yaml:"num_partitions"`
	ReplicationFactor int    `yaml:"replication_factor"`
	RetentionsMS      int    `yaml:"retention_ms"`
}

type Producer struct {
	RequiredAcks    int `yaml:"required_acks"`
	MaxOpenRequests int `yaml:"max_open_requests"`
	MaxRetries      int `yaml:"max_retries"`
	RetryBackoff    int `yaml:"retry_backoff"`
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

	return &cfg
}
