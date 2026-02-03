package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/poymanov/codemania-task-board/board/internal/config/env"
)

var appConfig *Config

type Config struct {
	Grpc GrpcConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	grpcCfg, err := env.NewGrpcConfig()
	if err != nil {
		return err
	}

	appConfig = &Config{
		Grpc: grpcCfg,
	}

	return nil
}

func AppConfig() *Config {
	return appConfig
}
