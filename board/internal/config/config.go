package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/poymanov/codemania-task-board/board/internal/config/env"
)

var appConfig *Config

type Config struct {
	Grpc   GrpcConfig
	Logger LoggerConfig
	Db     DbConfig
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

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	db, err := env.NewDbConfig()
	if err != nil {
		return err
	}

	appConfig = &Config{
		Grpc:   grpcCfg,
		Logger: loggerCfg,
		Db:     db,
	}

	return nil
}

func AppConfig() *Config {
	return appConfig
}
