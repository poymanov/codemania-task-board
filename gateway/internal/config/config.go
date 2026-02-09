package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/poymanov/codemania-task-board/gateway/internal/config/env"
)

type Config struct {
	Logger LoggerConfig
	Http   HttpConfig
}

func Load(path ...string) (*Config, error) {
	err := godotenv.Load(path...)

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return nil, err
	}

	httpCfg, err := env.NewHttpConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Logger: loggerCfg,
		Http:   httpCfg,
	}, nil
}
