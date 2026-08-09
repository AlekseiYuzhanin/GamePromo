package config

import (
	"errors"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database   DbConfig         `yaml:"database" env:"DATABASE"`
	GrpcClient GRPCClientConfig `yaml:"grpc_client_config" env:"GRPC_CLIENT_CONFIG"`
	GrpcServer GRPCServerConfig `yaml:"grpc_server_config" env:"GRPC_SERVER_CONFIG"`
}

type DbConfig struct {
	User     string `yaml:"user" env:"ITEMS_SERVICE_DB_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"ITEMS_SERVICE_DB_PASSWORD" env-default:"postgres"`
	DbName   string `yaml:"dbName" env:"ITEMS_SERVICE_DB_NAME" env-default:"postgres"`
	Host     string `yaml:"host" env:"ITEMS_SERVICE_DB_HOST" env-default:"127.0.0.1"`
	Port     string `yaml:"port" env:"ITEMS_SERVICE_DB_PORT" env-default:"5432"`
	SslMode  string `yaml:"ssl_mode" env:"ITEMS_SERVICE_SSL_MODE" env-default:"disable"`
}

type GRPCClientConfig struct {
	Ip   string `yaml:"grpc_client_ip" env:"GRPC_CLIENT_IP" env-default:"localhost"`
	Port string `yaml:"grpc_client_ip" env:"GRPC_CLIENT_PORT" env-default:"54251"`
}

type GRPCServerConfig struct {
	Ip   string `yaml:"grpc_server_ip" env:"GRPC_SERVER_IP" env-default:"localhost"`
	Port string `yaml:"grpc_server_ip" env:"GRPC_SERVER_PORT" env-default:"54252"`
}

func New(path string) *Config {
	var cfg Config
	if path != "" {
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		}
		err := cleanenv.ReadConfig(path, &cfg)
		if err != nil {
			log.Fatal(err)
		}
		return &cfg
	}
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
